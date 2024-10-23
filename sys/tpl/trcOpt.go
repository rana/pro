package tpl

import (
	"sys/k"
	"sys/tpl/atr"
)

type (
	FleTrcOpt struct {
		FleBse
		PrtEnumFlg
	}
)

func (x *DirTrc) NewTrc() (r *FleTrcOpt) {
	r = &FleTrcOpt{}
	x.Opt = r
	r.Name = k.Opt
	r.Pkg = x.Pkg
	r.Alias(r.Name, Uint32, atr.None)
	r.AddFle(r)
	return r
}
func (x *FleTrcOpt) InitCnst() {
	x.Cnst(k.Run, "")
	x.Cnst(k.Prv, "")

	x.Cnst(k.TicrRx, "")
	x.Cnst(k.TicrTx, "")

	x.Cnst(k.RltInstr, "")
	x.Cnst(k.RltInrvl, "")
	x.Cnst(k.RltSide, "")
	x.Cnst(k.RltStm, "")
	x.Cnst(k.RltCnd, "")
	x.Cnst(k.RltStgy, "")
	x.Cnst(k.RltPort, "")
	x.Cnst(k.RltPrfm, "")

	x.Cnst(k.HstInstr, "")
	x.Cnst(k.HstInrvl, "")
	x.Cnst(k.HstSide, "")
	x.Cnst(k.HstStm, "")
	x.Cnst(k.HstCnd, "")
	x.Cnst(k.HstStgy, "")
	x.Cnst(k.HstPort, "")
	x.Cnst(k.HstPrfm, "")

	x.Cnst(k.HstInstrFbr, "")
	x.Cnst(k.HstInrvlFbr, "")
	x.Cnst(k.HstSideFbr, "")
	x.Cnst(k.HstStmFbr, "")
	x.Cnst(k.HstCndFbr, "")
	x.Cnst(k.HstStgyFbr, "")
	x.Cnst(k.HstPortFbr, "")
	x.Cnst(k.HstPrfmFbr, "")

	x.Cnst(k.Tune, "")
}
func (x *FleTrcOpt) InitTypFn() {

}
func (x *FleTrcOpt) TrcOpt() (r *MemSig) {
	r = x.MemSig(k.TrcOpt)
	return r
}
