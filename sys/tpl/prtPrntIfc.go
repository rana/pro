package tpl

import "sys/tpl/atr"

type (
	PrtPrntIfc struct {
		PrtBse
	}
)

func (x *PrtPrntIfc) InitPrtIfc() {
	var sig *MemSig
	ifc := x.f.Ifcf("I%v", atr.None, x.t.Title())
	sig = x.f.MemSigRxrf("%vSet", ifc, x.t.Title())
	sig.InPrm(x.t, "v")
	sig = x.f.MemSigRxrf("%vGet", ifc, x.t.Title())
	sig.OutPrm(x.t)
}
