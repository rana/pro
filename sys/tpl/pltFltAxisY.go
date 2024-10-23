package tpl

import (
	"sys/k"
	"sys/tpl/atr"
)

type (
	FlePltFltAxisY struct {
		FleBse
		PrtStructFldSet
	}
)

func (x *DirPlt) NewFltAxisY() (r *FlePltFltAxisY) {
	r = &FlePltFltAxisY{}
	x.FltAxisY = r
	r.Name = k.FltAxisY
	r.Pkg = x.Pkg
	r.StructPtr(r.Name, atr.TypUiStructSupport)
	r.AddFle(r)
	return r
}
func (x *FlePltFltAxisY) InitFld(s *Struct) {
	s.Fld("Height", Uint32)
	s.Fld("PxlPerVal", Float32)
	s.Fld("Min", _sys.Bsc.Flt).Atr = atr.SetGet
	s.Fld("Max", _sys.Bsc.Flt).Atr = atr.SetGet
	s.Fld("Rng", _sys.Bsc.Flt)
	s.Fld("EqiDst", _sys.Bsc.Flt).Atr = atr.SetGet
	s.FldSlice("Inrvls", _sys.Bsc.Flt)
	s.FldSlice("Lns", _sys.Bsc.Flt)
	s.Fld("Rht", NewExt("SideFltAxisY"))
	s.Fld("vis", _sys.Bsc.Bol).Atr = atr.PrtFldSet
}
