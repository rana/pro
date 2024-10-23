package tpl

import (
	"sys/k"
	"sys/tpl/atr"
)

const (
// TODO: RECOMPUTE: CONSIDER DYNAMIC BASED ON FLDS
// TrdSize = BolSize + (7 * FltSize) + (3 * TmeSize)
)

type (
	FleAnaTrd struct {
		FleBse
		PrtPkt
		// PrtStructIdn
		PrtString
		PrtStructStrWrt
		// PrtStructBytWrt
		// PrtStructBytRed
		// PrtBytes
		// PrtLog
		// PrtIfc
	}
	FleAnaTrds struct {
		FleBse
		PrtArr
		// PrtArrIdn
		PrtArrFldSel
		PrtArrFld
		PrtArrStrWrt
		PrtString
		// PrtArrBytWrt
		// PrtBytes
		// PrtLog
		// PrtIfc
	}
)

func (x *DirAna) NewAnaTrd() (r *FleAnaTrd) {
	r = &FleAnaTrd{}
	x.Trd = r
	r.Name = k.Trd
	r.Pkg = x.Pkg
	r.StructPtr(r.Name, atr.AnaTrd)
	r.AddFle(r)
	return r
}
func (x *DirAna) NewAnaTrds() (r *FleAnaTrds) {
	r = &FleAnaTrds{}
	x.Trds = r
	r.FleBse = *NewArr(x.Trd, &r.PrtArr, x.Trd.Pkg)
	r.PrtArrStrWrt.Ln = true
	r.AddFle(r)
	return r
}

func (x *FleAnaTrd) InitFld(s *Struct) {
	s.Fld("OpnTme", _sys.Bsc.Tme)
	s.Fld("ClsTme", _sys.Bsc.Tme)
	s.Fld("OpnBid", _sys.Bsc.Flt)
	s.Fld("ClsBid", _sys.Bsc.Flt)
	s.Fld("OpnAsk", _sys.Bsc.Flt)
	s.Fld("ClsAsk", _sys.Bsc.Flt)
	s.Fld("OpnSpd", _sys.Bsc.Flt)
	s.Fld("ClsSpd", _sys.Bsc.Flt)
	s.Fld("ClsRsn", _sys.Bsc.Str).Atr = atr.EqlSkp

	s.Fld("Pip", _sys.Bsc.Flt)
	s.Fld("Dur", _sys.Bsc.Tme)
	s.Fld("IsLong", _sys.Bsc.Bol)
	s.Fld("PnlPct", _sys.Bsc.Flt)
	s.Fld("PnlPctPredict", _sys.Bsc.Flt)
	s.Fld("PnlUsd", _sys.Bsc.Flt)
	s.Fld("PnlGrsUsd", _sys.Bsc.Flt)
	s.Fld("CstComUsd", _sys.Bsc.Flt)
	s.Fld("CstClsSpdUsd", _sys.Bsc.Flt)
	s.Fld("CstOpnSpdUsd", _sys.Bsc.Flt)
	s.Fld("OpnBalUsd", _sys.Bsc.Flt)
	s.Fld("ClsBalUsd", _sys.Bsc.Flt)
	s.Fld("ClsBalUsdAct", _sys.Bsc.Flt)
	s.Fld("TrdPct", _sys.Bsc.Flt)
	s.Fld("MrgnRtio", _sys.Bsc.Flt)
	s.Fld("Units", _sys.Bsc.Flt)
	s.Fld("Instr", _sys.Bsc.Str)

	s.Fld("OpnReq", _sys.Bsc.Str).Atr = atr.TstSkp
	s.Fld("OpnRes", _sys.Bsc.Str).Atr = atr.TstSkp
	s.Fld("ClsReq", _sys.Bsc.Str).Atr = atr.TstSkp
	s.Fld("ClsRes", _sys.Bsc.Str).Atr = atr.TstSkp
	// TODO: ADD clsPrfLim?
	// TODO: ADD clsLosLim?
	// TODO: TRACK OPN PRICE SLIPPAGE FROM OPN SIGNAL TO OPN ACTUAL
	// TODO: TRACK CLS PRICE SLIPPAGE FROM CLS SIGNAL TO CLS ACTUAL
}
func (x *FleAnaTrd) InitCnst() {
	// x.CnstSize(TrdSize)
}
func (x *FleAnaTrd) InitTypFn() {
	x.OpnMid()
	x.ClsMid()
}
func (x *FleAnaTrd) OpnMid() (r *TypFn) {
	r = x.TypFn("OpnMid")
	r.OutPrm(_sys.Bsc.Flt)
	r.Add("return x.OpnBid + (x.OpnAsk - x.OpnBid) * 0.5")
	return r
}
func (x *FleAnaTrd) ClsMid() (r *TypFn) {
	r = x.TypFn("ClsMid")
	r.OutPrm(_sys.Bsc.Flt)
	r.Add("return x.ClsBid + (x.ClsAsk - x.ClsBid) * 0.5")
	return r
}
