package tpl

import (
	"sys/k"
	"sys/tpl/atr"
)

type (
	DirXpr struct {
		DirBse
		Knd  *FleKnd
		Scp  *FleScpXpr
		Ifc  *Ifc
		Xprr *FleXprr
	}
)

func (x *DirPro) NewXpr() (r *DirXpr) {
	r = &DirXpr{}
	x.Xpr = r
	r.Pkg = x.Pkg.New(k.Xpr)
	r.NewKnd()
	r.Scp = r.NewScp()
	r.Xprr = r.NewXprr()
	r.Ifc = r.NewIfc()
	return r
}
func (x *DirXpr) NewIfc() (r *Ifc) {
	r = x.Xprr.Ifc(k.Xpr, atr.None)
	x.Xprr.MemSig(k.Xpr, r)
	return r
}
