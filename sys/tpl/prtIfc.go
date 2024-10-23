package tpl

import (
	"sys/k"
	"sys/tpl/atr"
)

type (
	PrtIfc struct {
		PrtBse
	}
)

func (x *PrtIfc) InitPrtTyp() {
	x.f.ImplIfc(_sys.Ifc.Typ().(*Ifc))
}
func (x *PrtIfc) InitPrtTypFn() {
	x.ifc()
}
func (x *PrtIfc) ifc() (r *TypFn) {
	var rxr Typ
	var ifc *Ifc
	var isRxrIfc bool
	if ifc, isRxrIfc = x.f.Typ().(*Ifc); isRxrIfc {
		rxr = ifc.bse
	} else {
		rxr = x.f.Typ()
	}
	r = x.f.TypFna(k.Ifc, atr.None, rxr)
	r.OutPrm(Interface)
	if x.f.Typ() == _sys.Bsc.Str.Typ() {
		r.Add("return string(x)") // so quotes don't appear everywhere in the log
	} else {
		r.Add("return x")
	}
	if isRxrIfc {
		x.f.MemSigFn(r, ifc)
	}
	return r
}
