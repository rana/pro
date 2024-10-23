package tpl

import (
	"sys/k"
	"sys/tpl/atr"
)

type (
	FleSysAct struct {
		FleBse
	}
)

func (x *DirSys) NewAct() (r *FleSysAct) {
	r = &FleSysAct{}
	x.Act = r
	r.Name = k.Act
	r.Pkg = x.Pkg
	r.Ifc(r.Name, atr.None)
	r.AddFle(r)
	return r
}
func (x *FleSysAct) InitTypFn() {
	x.MemSig(k.Act)
}
