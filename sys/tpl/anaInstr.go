package tpl

import (
	"sys/k"
	"sys/tpl/atr"
)

type (
	FleAnaInstr struct {
		FleBse
		PrtStructCpy
		PrtStructStrWrt
		PrtString
		PrtStructBytWrt
		PrtStructBytRed
		PrtBytes
		// PrtLog
		// PrtIfc
	}
)

func (x *DirAna) NewAnaInstr() (r *FleAnaInstr) {
	r = &FleAnaInstr{}
	x.Instr = r
	r.Name = k.Instr
	r.Pkg = x.Pkg
	r.StructPtr(r.Name, atr.AnaInstr)
	r.AddFle(r)
	return r
}
func (x *FleAnaInstr) InitFld(s *Struct) {
	s.Fld("Name", _sys.Bsc.Str).Atr = atr.BytLitSkp

	s.Fld("Pip", _sys.Bsc.Flt)
	s.Fld("MrgnRtio", _sys.Bsc.Flt)
	s.Fld("SpdMin", _sys.Bsc.Flt)
	s.Fld("SpdMax", _sys.Bsc.Flt)
	s.Fld("SpdMdn", _sys.Bsc.Flt)
	s.Fld("SpdAvg", _sys.Bsc.Flt)
	s.Fld("SpdStd", _sys.Bsc.Flt)
	s.Fld("SpdOpnLim", _sys.Bsc.Flt)
	s.Fld("Fst", _sys.Bsc.Tme)
	s.Fld("Lst", _sys.Bsc.Tme)
	s.Fld("TmeCnt", _sys.Bsc.Unt)
	s.Fld("DayCnt", _sys.Bsc.Unt)
	s.Fld("DisplayPrecision", _sys.Bsc.Unt)
	s.Fld("TradeUnitsPrecision", _sys.Bsc.Unt)
	s.Fld("MinTrdSize", _sys.Bsc.Unt)
	s.Fld("MaxTrailingStopDistance", _sys.Bsc.Flt)
	s.Fld("MinTrailingStopDistance", _sys.Bsc.Flt)
	s.Fld("MaxPositionSize", _sys.Bsc.Unt)
	s.Fld("MaxOrderUnits", _sys.Bsc.Unt)
	s.Fld("Typ", _sys.Bsc.Str).Atr = atr.LitSkp

	s.Fld("Prv", _sys.Ana.Prv).Atr = atr.BytLitStrEqlBqTstSkp
	s.Fld("HstStm", _sys.Ana.Stm).Atr = atr.BytLitStrEqlBqTstSkp
	s.Fld("HstMu", RWMutex).Atr = atr.BytLitStrEqlBqTstSkp
	s.Fld("RltStm", _sys.Ana.Stm).Atr = atr.BytLitStrEqlBqTstSkp
	s.Fld("RltSubs", _sys.Ana.TmeIdx.Rxs).Atr = atr.BytLitStrEqlBqTstSkp
	s.Fld("RltSubsMu", Mutex).Atr = atr.BytLitStrEqlBqTstSkp
	s.Fld("RltLstPktTme", _sys.Bsc.Tme).Atr = atr.BytLitStrEqlBqTstSkp
	s.Fld("RltInrvlMax", _sys.Bsc.Tme).Atr = atr.BytLitStrEqlBqTstSkp
	s.Fld("MktWeeks", _sys.Bsc.TmeRng.arr).Atr = atr.BytLitStrEqlBqTstSkp
	s.Fld("MktDays", _sys.Bsc.TmeRng.arr).Atr = atr.BytLitStrEqlBqTstSkp
	s.Fld("MktHrs", _sys.Bsc.TmeRng.arr).Atr = atr.BytLitStrEqlBqTstSkp
}
