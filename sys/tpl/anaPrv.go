package tpl

import (
	"sys/k"
	"sys/tpl/atr"
)

type (
	FleAnaPrv struct {
		FleBse
	}
)

func (x *DirAna) NewPrv() (r *FleAnaPrv) {
	r = &FleAnaPrv{}
	x.Prv = r
	r.Name = k.Prv
	r.Pkg = x.Pkg
	r.Ifc(r.Name, atr.None)
	r.AddFle(r)
	return r
}
func (x *FleAnaPrv) InitFld(s *Struct) {
	var sig *MemSig

	sig = x.MemSig(k.AcntRefresh)
	sig.OutPrm(_sys.Bsc.Flt, "balance")

	sig = x.MemSig(k.Instr)
	sig.InPrm(_sys.Bsc.Str, "name")
	sig.OutPrm(_sys.Ana.Instr)

	sig = x.MemSig(k.Sub)
	sig.InPrm(_sys.Ana.Instr, "i")

	sig = x.MemSig(k.Unsub)
	sig.InPrm(_sys.Ana.Instr, "i")

	sig = x.MemSig(k.LoadHst)
	sig.InPrm(_sys.Ana.Instr, "i")

	sig = x.MemSig(k.OpnTrd)
	sig.InPrm(_sys.Ana.Trd, "trd")
	sig.InPrm(_sys.Ana.Instr, "i")
	sig.OutPrm(_sys.Bsc.Bol, "ok")
	sig.OutPrm(_sys.Ana.TrdRsnOpn, "rsn")

	sig = x.MemSig(k.ClsTrd)
	sig.InPrm(_sys.Ana.Trd, "trd")
	sig.InPrm(_sys.Ana.Instr, "i")
	sig.OutPrm(_sys.Bsc.Bol, "ok")

	sig = x.MemSig(k.CalcOpn)
	sig.InPrm(_sys.Ana.Trd, "trd")
	sig.InPrm(_sys.Ana.Instr, "i")

	sig = x.MemSig(k.CalcCls)
	sig.InPrm(_sys.Ana.Trd, "trd")
	sig.InPrm(_sys.Ana.Instr, "i")
}
