package tpl

import (
	"sys"
	"sys/k"
	"sys/tpl/atr"
)

type (
	FleAnaTmeIdx struct {
		FleBse
		// PrtStructIdn
		PrtPkt
	}
	FleAnaTmeIdxs struct {
		FleBse
		PrtArr
	}
)

func (x *DirAna) NewAnaTmeIdx() (r *FleAnaTmeIdx) {
	r = &FleAnaTmeIdx{}
	x.TmeIdx = r
	r.Name = k.TmeIdx
	r.Pkg = x.Pkg
	r.Struct(r.Name, atr.AnaTmeIdx)
	r.AddFle(r)
	r.NewArr()
	return r
}
func (x *FleAnaTmeIdx) NewArr() (r *FleAnaTmeIdxs) {
	r = &FleAnaTmeIdxs{}
	r.FleBse = *NewArr(x, &r.PrtArr, x.Pkg)
	r.AddFle(r)
	return r
}
func (x *FleAnaTmeIdx) InitVals(bse *TypBse) {
	bse.Vals = sys.VsStruct(x.Typ().Full(), "Tme:tme.Tme(0), Idx:0", "Tme:tme.Tme(10), Idx:1", "Tme:tme.Tme(100), Idx:2")
}
func (x *FleAnaTmeIdx) InitFld(s *Struct) {
	s.Fld("Tme", _sys.Bsc.Tme)
	s.Fld("Idx", _sys.Bsc.Unt)
}
