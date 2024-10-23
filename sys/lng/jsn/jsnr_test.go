package jsn_test

import (
	"fmt"
	"sys/bsc/bnd"
	"sys/bsc/flt"
	"sys/bsc/tme"
	"sys/bsc/unt"
	"sys/lng/jsn"
	"sys/tst"
	"testing"
)

func TestJsnStrFltValid(t *testing.T) {
	var j jsn.Jsnr
	e := flt.Flt(1.1)
	j.Resetf("{\"key\":%q}", e)
	a := j.StrFlt("key")
	tst.FltEql(t, e, a)
}

func TestJsnStrTmeValid(t *testing.T) {
	var j jsn.Jsnr
	e := tme.NewDteTme(2018, 1, 2, 3, 4, 5)
	j.Reset(" {\"key\":\"2018-01-02T03:04:05.521193208Z\"} ")
	a := j.StrTme("key")
	tst.TmeEql(t, e, a)
}
func TestJsnArrValid(t *testing.T) {
	var j jsn.Jsnr
	start := " {\"key\":"
	arr := " [0, 1, 2] "
	e := bnd.Bnd{Idx: unt.Unt(len(start) + 1), Lim: unt.Unt(len(start) + len(arr) - 1)}
	j.Resetf("%v%v}", start, arr)
	a := j.Arr("key")
	tst.BndEql(t, e, a)
}
func TestJsnElmObjValid(t *testing.T) {
	var j jsn.Jsnr
	txt := " { \"price\": \"1.16481\", \"liquidity\": 10000000 } "
	e := bnd.Bnd{Idx: unt.One, Lim: unt.Unt(len(txt) - 1)}
	j.Reset(txt)
	a := j.ElmObj()
	tst.BndEql(t, e, a)
}
func TestJsnElmObjNestedValid(t *testing.T) {
	var j jsn.Jsnr
	txt := " { \"a\" : { \"b\" : { \"price\": \"1.16481\", \"liquidity\": 10000000 } } } "
	e := bnd.Bnd{Idx: unt.One, Lim: unt.Unt(len(txt) - 1)}
	j.Reset(txt)
	a := j.ElmObj()
	tst.BndEql(t, e, a)
}
func TestJsnArrObjsValid(t *testing.T) {
	var j jsn.Jsnr
	e := []flt.Flt{1.1, 1.2}
	start := `{
		"type": "PRICE",
		"time": "2017-10-31T17:21:59.521193208Z",
		"bids": [
			`
	elm0 := fmt.Sprintf(`{
				"price": "%v",
				"liquidity": 10000000
			},`, e[0])
	elm1 := fmt.Sprintf(`{
				"price": "%v",
				"liquidity": 20000000
			}`, e[1])
	end := `],
			"asks": [
				{
					"price": "1.16485",
					"liquidity": 10000000
				}
			],
			"closeoutBid": "1.16461",
			"closeoutAsk": "1.16505",
			"status": "tradeable",
			"tradeable": true,
			"instrument": "EUR_USD"
		}`

	txt := fmt.Sprintf("%v%v%v%v", start, elm0, elm1, end)
	j.Reset(txt)
	a := j.ArrObjs("bids")
	tst.IntegerEql(t, 2, len(a))

	e0 := bnd.Bnd{Idx: unt.Unt(len(start)), Lim: unt.Unt(len(start) + len(elm0) - 1)}
	tst.BndEql(t, e0, a[0])

	e1 := bnd.Bnd{Idx: unt.Unt(len(start) + len(elm0)), Lim: unt.Unt(len(start) + len(elm0) + len(elm1))}
	tst.BndEql(t, e1, a[1])

	j.Reset(txt[a[0].Idx:a[0].Lim])
	a0 := j.StrFlt("price")
	tst.FltEql(t, e[0], a0)

	j.Reset(txt[a[1].Idx:a[1].Lim])
	a1 := j.StrFlt("price")
	tst.FltEql(t, e[1], a1)
}
