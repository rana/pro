package tpl

import (
	"sys/k"
	"sys/tpl/atr"
)

type (
	FleAnaPth struct {
		FleBse
	}
)

func (x *DirAna) NewPth() (r *FleAnaPth) {
	r = &FleAnaPth{}
	x.Pth = r
	r.Name = k.Pth
	r.Pkg = x.Pkg
	r.Ifc(r.Name, atr.None)
	r.AddFle(r)
	return r
}
func (x *FleAnaPth) InitFld(s *Struct) {
	var sig *MemSig
	sig = x.MemSig(k.Name)
	sig.OutPrm(_sys.Bsc.Str)
	sig = x.MemSig(k.PrmWrt)
	sig.InPrm(BuilderPtr, "b")
	sig = x.MemSig(k.Prm)
	sig.OutPrm(String)
	// sig = x.MemSig(k.Ttl)
	// sig.OutPrm(_sys.Bsc.Str)
	sig = x.MemSig(k.StrWrt)
	sig.InPrm(BuilderPtr, "b")
	sig = x.MemSig(k.String)
	sig.OutPrm(String)
}
