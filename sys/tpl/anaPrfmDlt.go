package tpl

import (
	"sys/k"
	"sys/tpl/atr"
)

type (
	FleAnaPrfmDlt struct {
		FleBse
		PrtStructStrWrt
		PrtString
		PrtStructBytWrt
		PrtStructBytRed
		PrtBytes
		// PrtLog
		// PrtIfc
		New *PkgFn
	}
)

func (x *DirAna) NewPrfmDlt() (r *FleAnaPrfmDlt) {
	r = &FleAnaPrfmDlt{}
	x.PrfmDlt = r
	r.Name = k.PrfmDlt
	r.Pkg = x.Pkg
	r.StructPtr(r.Name, atr.LngScp)
	r.AddFle(r)
	return r
}

func (x *FleAnaPrfmDlt) InitFld(s *Struct) {
	prfm := _sys.Ana.Prfm.Typ().(*Struct)
	for _, fld := range prfm.Flds {
		if fld.IsDlt() {
			s.Fldf("%vA", fld.Typ, fld.Name).Atr = atr.Get
			s.Fldf("%vB", fld.Typ, fld.Name).Atr = atr.Get
			s.Fldf("%vDlt", _sys.Bsc.Flt, fld.Name).Atr = atr.Get // use flt in case unt goes neg
			// s.Fldf("%vDltPct", _sys.Bsc.Flt, fld.Name) // use flt for neg pct
		}
	}
	s.Fld("PthB", _sys.Bsc.Str).Atr = atr.Get
}
