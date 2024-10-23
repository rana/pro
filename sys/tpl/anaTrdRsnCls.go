package tpl

import (
	"sys/k"
	"sys/tpl/atr"
)

type (
	FleAnaTrdRsnCls struct {
		FleBse
		PrtEnum
	}
)

func (x *DirAna) NewTrdRsnCls() (r *FleAnaTrdRsnCls) {
	r = &FleAnaTrdRsnCls{}
	x.TrdRsnCls = r
	r.Name = k.TrdRsnCls
	r.Pkg = x.Pkg
	r.Alias(r.Name, Byte, atr.None)
	r.AddFle(r)
	return r
}
func (x *FleAnaTrdRsnCls) InitCnst() {
	x.Cnst("Prf", "", x)
	x.Cnst("Los", "", x)
	x.Cnst("Dur", "", x)
	x.Cnst("Cnd", "", x)
}
