package tpl

import (
	"sys"
	"sys/k"
)

type (
	ExtFle struct {
		FleBse
	}
)

func (x *DirSys) NewExt() (r *ExtFle) {
	r = &ExtFle{}
	x.Ext = r
	r.Name = k.Ext
	r.Pkg = x.Pkg
	r.AddExt(Interface)
	r.AddFle(r)
	return r
}
func (x *ExtFle) AddExt(e *Ext) {
	e.Pkg = x.Pkg
	x.AddTyp(e)
}
func (x *ExtFle) InitVals(bse *TypBse) {
	Interface.Lits = sys.Vs("\"xYz\"")
	Interface.Vals = Interface.Lits
}
