package tpl

import (
	"sys/k"
	"sys/tpl/atr"
)

type (
	PrtPkt struct {
		PrtBse
		Rx      *Func
		Rxs     *Map
		Tx      *Struct
		Txr     *Ifc
		NewTx   *PkgFn
		Act     *TypFn
		Ret     *TypFn
		Tier    *TypFn
		DecTier *TypFn
	}
)

func (x *PrtPkt) InitPrtTyp() {
	x.Rx = x.f.Funcf("%vRx", atr.None, x.t.Title())
	x.Rxs = x.f.Mapf("%vRxs", Uint64, x.Rx, atr.None, x.t.Title())
	x.Tx = x.f.StructPtrf("%vTx", atr.None, x.t.Title())
	x.Txr = x.f.Ifcf("%vTxr", atr.None, x.t.Title())
}
func (x *PrtPkt) InitPrtFld() {
	// Rx
	x.Rx.InPrm(x.f.Typ(), "pkt")
	x.Rx.OutPrmSlice(_sys.Act.Typ())
	// Tx
	x.Tx.Fld("Pkt", x.t)
	x.Tx.Fld("Rx", x.Rx)
	x.Tx.FldSlice("ret", _sys.Act.Typ())
	x.Tx.Fld("tier", Int)
	// Txr
	var sig *MemSig
	sig = x.f.MemSig(k.Sub, x.Txr)
	sig.InPrm(x.Rx, "rx")
	sig.InPrm(Uint32, "id")
	sig.InPrmVariadic(Uint32, "slot")
	sig = x.f.MemSig(k.Unsub, x.Txr)
	sig.InPrm(Uint32, "id")
	sig.InPrmVariadic(Uint32, "slot")
}
func (x *PrtPkt) InitPrtPkgFn() {
	x.NewTx = x.newTx()
}
func (x *PrtPkt) InitPrtTypFn() {
	x.Act = x.act()
	x.Ret = x.ret()
	x.Tier = x.tier()
	x.DecTier = x.decTier()
}
func (x *PrtPkt) newTx() (r *PkgFn) {
	r = x.f.PkgFnf("New%v", x.Tx.Title())
	x.NewTx = r
	r.InPrm(x.f.Typ(), "pkt")
	r.InPrm(x.Rx, "rx")
	r.InPrmVariadic(Int, "tier")
	r.OutPrm(x.Tx, "r")
	r.Addf("r = %v{}", r.OutTyp().Adr(x.f))
	r.Add("r.Pkt = pkt")
	r.Add("r.Rx = rx")
	r.Add("if len(tier) > 0 {")
	r.Add("r.tier = tier[0]")
	r.Add("}")
	r.Add("return r")
	return r
}
func (x *PrtPkt) act() (r *TypFn) {
	r = x.f.TypFn("Act", x.Tx)
	r.Add("x.ret = x.Rx(x.Pkt)")
	return r
}
func (x *PrtPkt) ret() (r *TypFn) {
	r = x.f.TypFn("Ret", x.Tx)
	r.OutPrmSlice(_sys.Act.Typ())
	r.Add("return x.ret")
	return r
}
func (x *PrtPkt) tier() (r *TypFn) {
	r = x.f.TypFn("Tier", x.Tx)
	r.OutPrm(Int)
	r.Add("return x.tier")
	return r
}
func (x *PrtPkt) decTier() (r *TypFn) {
	r = x.f.TypFn("DecTier", x.Tx)
	r.Add("x.tier--")
	return r
}
