package ana

import (
	"fmt"
	"strconv"
	"strings"
	"sys"
	"sys/bsc/bol"
	"sys/bsc/flt"
	bscint "sys/bsc/int"
	"sys/bsc/str"
)

type (
	OrdReq struct {
		Type        str.Str
		Instrument  str.Str
		Units       bscint.Int
		TimeInForce str.Str
		PriceBound  flt.Flt
	}
	PosClsReq struct {
		IsBuy bol.Bol
	}
)

// {
//   "order": {
//     "price": "1.2000",
//     "timeInForce": "GTC",
//     "instrument": "EUR_CAD",
//     "units": "10000",
//     "clientExtensions": { // TODO
//       "comment": "New idea for trading",
//       "tag": "strategy_9",
//       "id": "my_order_100"
//     },
//     "type": "MARKET_IF_TOUCHED",
//     "positionFill": "DEFAULT"
//   }
// }
func (x *OrdReq) JsnWrt(i *Instr, b *strings.Builder) {
	// http://developer.oanda.com/rest-live-v20/order-df/#OrderRequest
	b.WriteString("{\"order\":{")
	b.WriteString("\"type\":")
	x.Type.StrWrt(b)
	b.WriteString(",\"instrument\":")
	x.Instrument.StrWrt(b)
	b.WriteString(",\"units\":")
	// x.Units.StrWrt(b) // TODO: JSNWRT?
	// int.Int prefixs sign, for json avoid sign prefix
	b.WriteString(fmt.Sprintf("%v", int(x.Units)))
	b.WriteString(",\"timeInForce\":")
	x.TimeInForce.StrWrt(b)
	if x.PriceBound != 0 {
		b.WriteString(",\"priceBound\":")
		// http err: 400 Bad Request {"orderRejectTransaction":{"type":"MARKET_ORDER_REJECT","rejectReason":"PRICE_BOUND_PRECISION_EXCEEDED","instrument":"EUR_USD","units":"213781","priceBound":"1.16941999999999","timeInForce":"FOK","positionFill":"DEFAULT","reason":"CLIENT_ORDER","id":"1556","userID":4093945,"accountID":"101-001-4093945-001","batchID":"1556","requestID":"24423653537521249","time":"2018-06-04T04:11:16.905713984Z"},"relatedTransactionIDs":["1556"],"lastTransactionID":"1556","errorMessage":"The price bound specified contains more precision than is allowed for the Order's instrument","errorCode":"PRICE_BOUND_PRECISION_EXCEEDED"}
		// sys.Log("******", "i.TradeUnitsPrecision", i.TradeUnitsPrecision)
		// sys.Log("******", "i.DisplayPrecision", i.DisplayPrecision)
		// x.PriceBound.Trnc(i.DisplayPrecision).StrWrt(b) // Trunc to avoid server response error regaring precision

		s := strconv.FormatFloat(float64(x.PriceBound), byte('f'), int(i.DisplayPrecision), 32)
		idx := strings.Index(s, ".")
		if len(s)-idx-1 > int(i.DisplayPrecision) {
			s = s[:idx+1+int(i.DisplayPrecision)]
			sys.Logf("*** priceBound EXTRA TRIM from:%v to:%v", strconv.FormatFloat(float64(x.PriceBound), byte('f'), int(i.DisplayPrecision), 32), s)
		}
		b.WriteString(s)
	}
	b.WriteString(",\"positionFill\": \"DEFAULT\"") // REQUIRED
	b.WriteString("}}")
}

func (x *PosClsReq) JsnWrt(b *strings.Builder) {
	// http://developer.oanda.com/rest-live-v20/order-ep/
	b.WriteString("{")
	if x.IsBuy {
		b.WriteString("\"longUnits\":\"ALL\"")
	} else {
		b.WriteString("\"shortUnits\":\"ALL\"")
	}
	b.WriteString("}")
}
