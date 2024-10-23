package ana

import (
	"bytes"
	"strings"
	"sync"
	"sys/bsc/flt"
	"sys/bsc/str"
	"sys/bsc/tme"
	"sys/bsc/unt"
)

type (
	Instr struct {
		Name                    str.Str
		Pip                     flt.Flt
		MrgnRtio                flt.Flt
		SpdMin                  flt.Flt
		SpdMax                  flt.Flt
		SpdMdn                  flt.Flt
		SpdAvg                  flt.Flt
		SpdStd                  flt.Flt
		SpdOpnLim               flt.Flt
		Fst                     tme.Tme
		Lst                     tme.Tme
		TmeCnt                  unt.Unt
		DayCnt                  unt.Unt
		DisplayPrecision        unt.Unt
		TradeUnitsPrecision     unt.Unt
		MinTrdSize              unt.Unt
		MaxTrailingStopDistance flt.Flt
		MinTrailingStopDistance flt.Flt
		MaxPositionSize         unt.Unt
		MaxOrderUnits           unt.Unt
		Typ                     str.Str
		Prv                     Prv
		HstStm                  *Stm
		HstMu                   sync.RWMutex
		RltStm                  *Stm
		RltSubs                 TmeIdxRxs
		RltSubsMu               sync.Mutex
		RltLstPktTme            tme.Tme
		RltInrvlMax             tme.Tme
		MktWeeks                *tme.Rngs
		MktDays                 *tme.Rngs
		MktHrs                  *tme.Rngs
	}
)

func (x *Instr) Cpy() (r *Instr) {
	r = &Instr{}
	r.Name = x.Name
	r.Pip = x.Pip
	r.MrgnRtio = x.MrgnRtio
	r.SpdMin = x.SpdMin
	r.SpdMax = x.SpdMax
	r.SpdMdn = x.SpdMdn
	r.SpdAvg = x.SpdAvg
	r.SpdStd = x.SpdStd
	r.SpdOpnLim = x.SpdOpnLim
	r.Fst = x.Fst
	r.Lst = x.Lst
	r.TmeCnt = x.TmeCnt
	r.DayCnt = x.DayCnt
	r.DisplayPrecision = x.DisplayPrecision
	r.TradeUnitsPrecision = x.TradeUnitsPrecision
	r.MinTrdSize = x.MinTrdSize
	r.MaxTrailingStopDistance = x.MaxTrailingStopDistance
	r.MinTrailingStopDistance = x.MinTrailingStopDistance
	r.MaxPositionSize = x.MaxPositionSize
	r.MaxOrderUnits = x.MaxOrderUnits
	r.Typ = x.Typ
	return r
}
func (x *Instr) StrWrt(b *strings.Builder) string {
	b.WriteString("ana.instr(")
	b.WriteString("name:")
	x.Name.StrWrt(b)
	b.WriteString(" pip:")
	x.Pip.StrWrt(b)
	b.WriteString(" mrgnRtio:")
	x.MrgnRtio.StrWrt(b)
	b.WriteString(" spdMin:")
	x.SpdMin.StrWrt(b)
	b.WriteString(" spdMax:")
	x.SpdMax.StrWrt(b)
	b.WriteString(" spdMdn:")
	x.SpdMdn.StrWrt(b)
	b.WriteString(" spdAvg:")
	x.SpdAvg.StrWrt(b)
	b.WriteString(" spdStd:")
	x.SpdStd.StrWrt(b)
	b.WriteString(" spdOpnLim:")
	x.SpdOpnLim.StrWrt(b)
	b.WriteString(" fst:")
	x.Fst.StrWrt(b)
	b.WriteString(" lst:")
	x.Lst.StrWrt(b)
	b.WriteString(" tmeCnt:")
	x.TmeCnt.StrWrt(b)
	b.WriteString(" dayCnt:")
	x.DayCnt.StrWrt(b)
	b.WriteString(" displayPrecision:")
	x.DisplayPrecision.StrWrt(b)
	b.WriteString(" tradeUnitsPrecision:")
	x.TradeUnitsPrecision.StrWrt(b)
	b.WriteString(" minTrdSize:")
	x.MinTrdSize.StrWrt(b)
	b.WriteString(" maxTrailingStopDistance:")
	x.MaxTrailingStopDistance.StrWrt(b)
	b.WriteString(" minTrailingStopDistance:")
	x.MinTrailingStopDistance.StrWrt(b)
	b.WriteString(" maxPositionSize:")
	x.MaxPositionSize.StrWrt(b)
	b.WriteString(" maxOrderUnits:")
	x.MaxOrderUnits.StrWrt(b)
	b.WriteString(" typ:")
	x.Typ.StrWrt(b)
	b.WriteRune(')')
	return b.String()
}
func (x *Instr) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *Instr) BytWrt(b *bytes.Buffer) {
	x.Pip.BytWrt(b)
	x.MrgnRtio.BytWrt(b)
	x.SpdMin.BytWrt(b)
	x.SpdMax.BytWrt(b)
	x.SpdMdn.BytWrt(b)
	x.SpdAvg.BytWrt(b)
	x.SpdStd.BytWrt(b)
	x.SpdOpnLim.BytWrt(b)
	x.Fst.BytWrt(b)
	x.Lst.BytWrt(b)
	x.TmeCnt.BytWrt(b)
	x.DayCnt.BytWrt(b)
	x.DisplayPrecision.BytWrt(b)
	x.TradeUnitsPrecision.BytWrt(b)
	x.MinTrdSize.BytWrt(b)
	x.MaxTrailingStopDistance.BytWrt(b)
	x.MinTrailingStopDistance.BytWrt(b)
	x.MaxPositionSize.BytWrt(b)
	x.MaxOrderUnits.BytWrt(b)
	x.Typ.BytWrt(b)
}
func (x *Instr) BytRed(b []byte) (idx int) {
	idx += x.Pip.BytRed(b)
	idx += x.MrgnRtio.BytRed(b[idx:])
	idx += x.SpdMin.BytRed(b[idx:])
	idx += x.SpdMax.BytRed(b[idx:])
	idx += x.SpdMdn.BytRed(b[idx:])
	idx += x.SpdAvg.BytRed(b[idx:])
	idx += x.SpdStd.BytRed(b[idx:])
	idx += x.SpdOpnLim.BytRed(b[idx:])
	idx += x.Fst.BytRed(b[idx:])
	idx += x.Lst.BytRed(b[idx:])
	idx += x.TmeCnt.BytRed(b[idx:])
	idx += x.DayCnt.BytRed(b[idx:])
	idx += x.DisplayPrecision.BytRed(b[idx:])
	idx += x.TradeUnitsPrecision.BytRed(b[idx:])
	idx += x.MinTrdSize.BytRed(b[idx:])
	idx += x.MaxTrailingStopDistance.BytRed(b[idx:])
	idx += x.MinTrailingStopDistance.BytRed(b[idx:])
	idx += x.MaxPositionSize.BytRed(b[idx:])
	idx += x.MaxOrderUnits.BytRed(b[idx:])
	idx += x.Typ.BytRed(b[idx:])
	return idx
}
func (x *Instr) Bytes() []byte {
	b := &bytes.Buffer{}
	x.BytWrt(b)
	return b.Bytes()
}
