package tpl

import (
	"sys/k"
	"sys/tpl/atr"
)

type (
	FleAnaTic struct {
		FleBse
	}
)

func (x *DirAna) NewAnaTic() (r *FleAnaTic) {
	r = &FleAnaTic{}
	x.Tic = r
	r.Name = k.Tic
	r.Pkg = x.Pkg
	r.StructPtr(r.Name, atr.None)
	r.AddFle(r)
	return r
}
func (x *FleAnaTic) InitFld(s *Struct) {
	s.Fld("Tme", _sys.Bsc.Tme)
	s.Fld("Bids", _sys.Bsc.Flt.arr)
	s.Fld("Asks", _sys.Bsc.Flt.arr)
}
