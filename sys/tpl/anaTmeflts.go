package tpl

import (
	"sys"
	"sys/k"
	"sys/tpl/atr"
)

type (
	// NOT ARR
	FleAnaTmeFlts struct {
		FleBse
		// PrtStructIdn
		PrtPkt
	}
)

func (x *DirAna) NewAnaTmeFlts() (r *FleAnaTmeFlts) {
	r = &FleAnaTmeFlts{}
	x.TmeFlts = r
	r.Name = k.TmeFlts
	r.Pkg = x.Pkg
	r.Struct(r.Name, atr.AnaTmeFlts)
	r.AddFle(r)
	return r
}
func (x *FleAnaTmeFlts) InitVals(bse *TypBse) {
	bse.Vals = sys.VsStruct(x.Typ().Full(), "Tme:0", "Tme:10", "Tme:100")
}
func (x *FleAnaTmeFlts) InitFld(s *Struct) {
	s.Fld("Tme", _sys.Bsc.Tme)
	s.Fld("Flts", _sys.Bsc.Flt.arr)
}
