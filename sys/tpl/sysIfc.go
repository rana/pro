package tpl

import (
	"sys/k"
	"sys/tpl/atr"
)

type (
	FleSysIfc struct {
		FleBse
	}
)

// for (x *Actr) RunIfc(txt string)
func (x *DirSys) NewIfc() (r *FleSysIfc) {
	r = &FleSysIfc{}
	x.Ifc = r
	r.Name = k.Ifc
	r.Pkg = x.Pkg
	r.Ifc(r.Name, atr.None) //atr.LngScp&^atr.Ifc)
	r.AddFle(r)
	return r
}
func (x *FleSysIfc) InitTypFn() {
	// // IFC SIG WRITTEN BY r.Ifc() CALL
	var sig *MemSig
	sig = x.MemSig(k.Ifc)
	// sig.OutPrm(x)
	sig.OutPrm(Interface)
}
