package tpl

import (
	"sys/k"
	"sys/tpl/atr"
)

type (
	FleHstTrdStmSeg struct {
		FleBse
		// PrtStructIdn
		PrtStructStrWrt
		PrtString
	}
	FleHstTrdStmSegs struct {
		FleBse
		PrtArr
		// PrtArrIdn
	}
)

func (x *DirHst) NewTrdStmSeg() (r *FleHstTrdStmSeg) {
	r = &FleHstTrdStmSeg{}
	// x.TrdStmSeg = r
	r.Name = k.TrdStmSeg
	r.Pkg = x.Pkg
	r.StructPtr(r.Name, atr.HstTrdStmSeg)
	r.AddFle(r)
	r.NewTrdStmSegs()
	return r
}
func (x *FleHstTrdStmSeg) NewTrdStmSegs() (r *FleHstTrdStmSegs) {
	r = &FleHstTrdStmSegs{}
	r.FleBse = *NewArr(x, &r.PrtArr, x.Pkg)
	r.AddFle(r)
	return r
}

func (x *FleHstTrdStmSeg) InitFld(s *Struct) {
	s.Fld("Trd", _sys.Ana.Trd)
	s.Fld("Tmes", _sys.Bsc.Tme.arr)
	s.Fld("Vals", _sys.Bsc.Flt.arr)
}
