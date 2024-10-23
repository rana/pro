package tpl

import (
	"sys/k"
	"sys/tpl/atr"
)

type (
	DirAct struct {
		DirBse
		Scp  *FleScpAct
		Actr *FleActr
		Ifc  *Ifc // USING sys.Act INSTEAD
	}
)

func (x *DirPro) NewAct() (r *DirAct) {
	r = &DirAct{}
	x.Act = r
	r.Pkg = x.Pkg.New(k.Act)
	r.Scp = r.NewScp()
	r.Actr = r.NewActr()
	r.Ifc = r.NewIfc()
	return r
}
func (x *DirAct) NewIfc() (r *Ifc) {
	r = x.Actr.Ifc(k.Act, atr.None)
	x.Actr.MemSig(k.Act, r)
	sig := x.Actr.MemSig(k.Ifc, r)
	sig.OutPrm(Interface)
	return r
}
