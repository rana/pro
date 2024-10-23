package tpl

import (
	"sys/k"
	"sys/tpl/atr"
)

type (
	FleSysIdn struct {
		FleBse
	}
)

func (x *DirSys) NewIdn() (r *FleSysIdn) {
	r = &FleSysIdn{}
	x.Idn = r
	r.Name = k.Idn
	r.Pkg = x.Pkg
	r.Alias(r.Name, String, atr.TypIdn)
	r.AddFle(r)
	return r
}
