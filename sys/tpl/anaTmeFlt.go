package tpl

import (
	"sys"
	"sys/k"
	"sys/tpl/atr"
)

type (
	FleAnaTmeFlt struct {
		FleBse
		// PrtStructIdn
		PrtPkt
	}
)

func (x *DirAna) NewAnaTmeFlt() (r *FleAnaTmeFlt) {
	r = &FleAnaTmeFlt{}
	x.TmeFlt = r
	r.Name = k.TmeFlt
	r.Pkg = x.Pkg
	r.Struct(r.Name, atr.AnaTmeFlt)
	r.AddFle(r)
	return r
}
func (x *FleAnaTmeFlt) InitVals(bse *TypBse) {
	bse.Vals = sys.VsStruct(x.Typ().Full(), "Tme:0, Flt:0", "Tme:10, Flt:10", "Tme:100, Flt:100")
}
func (x *FleAnaTmeFlt) InitFld(s *Struct) {
	s.Fld("Tme", _sys.Bsc.Tme)
	s.Fld("Flt", _sys.Bsc.Flt)
}
