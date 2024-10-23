package tpl

import (
	"sys/k"
	"sys/tpl/atr"
)

type (
	FlePltTmeAxisX struct {
		FleBse
		PrtStructFldSet
	}
)

func (x *DirPlt) NewTmeAxisX() (r *FlePltTmeAxisX) {
	r = &FlePltTmeAxisX{}
	x.TmeAxisX = r
	r.Name = k.TmeAxisX
	r.Pkg = x.Pkg
	r.StructPtr(r.Name, atr.TypUiStructSupport)
	r.AddFle(r)
	return r
}
func (x *FlePltTmeAxisX) InitFld(s *Struct) {
	s.Fld("Width", Uint32)
	s.Fld("PxlPerVal", Float32)
	s.Fld("Min", _sys.Bsc.Tme)
	s.Fld("Max", _sys.Bsc.Tme)
	s.Fld("Rng", _sys.Bsc.Tme)
	s.FldSlice("Inrvls", _sys.Bsc.Tme)
	s.FldSlice("Lns", _sys.Bsc.Tme)
	s.Fld("Btm", NewExt("SideTmeAxisX"))
	s.Fld("vis", _sys.Bsc.Bol).Atr = atr.PrtFldSet
}
