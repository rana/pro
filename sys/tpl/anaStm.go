package tpl

import (
	"sys/k"
	"sys/tpl/atr"
)

type (
	FleAnaStm struct {
		FleBse
		PrtBytes
		// PrtStructIdn
		PrtStructBytWrt
		PrtStructBytRed
	}
)

func (x *DirAna) NewAnaStm() (r *FleAnaStm) {
	r = &FleAnaStm{}
	x.Stm = r
	r.Name = k.Stm
	r.Pkg = x.Pkg
	r.StructPtr(r.Name, atr.AnaStm)
	r.AddFle(r)
	return r
}
func (x *FleAnaStm) InitFld(s *Struct) {
	s.Fld("Tmes", _sys.Bsc.Tme.arr)
	s.Fld("Bids", _sys.Bsc.Flt.arr)
	s.Fld("Asks", _sys.Bsc.Flt.arr)
	s.Fld("BidLims", _sys.Bsc.Unt.arr)
	s.Fld("AskLims", _sys.Bsc.Unt.arr)
	s.Fld("RxIdx", _sys.Bsc.Unt).Atr = atr.BytLitStrEqlBqTstSkp
	s.Fld("RxTme", _sys.Bsc.Tme).Atr = atr.BytLitStrEqlBqTstSkp
}
