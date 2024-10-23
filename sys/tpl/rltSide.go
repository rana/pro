package tpl

import (
	"strings"
	"sys/k"
	"sys/ks"
	"sys/tpl/atr"
)

type (
	FleRltSide struct {
		FleRltBse
	}
	FleRltSides struct {
		FleBse
		PrtArr
	}
)

func (x *DirRlt) NewSide() (r *FleRltSide) {
	r = &FleRltSide{}
	x.Side = r
	r.Name = k.Side
	r.Pkg = x.Pkg
	r.Ifc(r.Name, atr.TypAnaIfc)
	r.AddFle(r)
	return r
}
func (x *DirRlt) NewSides() (r *FleRltSides) {
	r = &FleRltSides{}
	x.Sides = r
	r.FleBse = *NewArr(x.Side, &r.PrtArr, x.Side.Pkg)
	r.AddFle(r)
	return r
}
func (x *FleRltSide) InitTyp(bse *TypBse) {
	x.FleRltBse.InitTyp(bse)
	x.Typ().Bse().TestPth = append(_sys.Ana.Rlt.Inrvl.Typ().Bse().TestPth, &TestStp{
		MdlFst: func(r *PkgFn) { r.Add("side := tst.RltInrvlSideBid(inrvl)") },
	})
}
func (x *FleRltSide) InitFld(s *Struct) {
	x.FleRltBse.InitFld(s)
	x.bse.FldPrnt(_sys.Ana.Rlt.Inrvl)
	x.FldRxs(_sys.Ana.TmeFlts.Rxs)
}

func (x *FleRltSide) InitTypFn() {
	x.FleRltBse.InitTypFn()
	x.Sub(_sys.Ana.TmeFlts.Rx)
	x.Unsub(x.bse, true, true, func(r *TypFn) {
		r.Addf("x.%v.Unsub(x.Id)", x.bse.Prnt().Name)
	})
	for _, stmRte := range ks.StmRtes {
		x.StmRte(stmRte)
	}
	x.Sar()
	x.Ema()
}

func (x *FleRltSide) Rx(nodeFn *TypFn) (r *TypFn) {
	r = x.TypFn(k.Rx, nodeFn.Node)
	r.InPrm(_sys.Bsc.Bnd, "inPkt")
	r.OutPrmSlice(_sys.Act, "r")
	r.Add("x.mu.Lock() // translate time range to val range")
	r.Add("if ana.Cfg.Trc.IsRltSide() {")
	r.Addf("sys.Logf(\"%v(%%v).%v %%p inPkt %%v\", x.Prm(), x, inPkt)", nodeFn.Node.Full(), r.Name)
	r.Add("}")
	r.Add("stm := x.Inrvl.Bse().Instr.Instr().RltStm")
	r.Addf("outPkt := %v{", _sys.Ana.TmeFlts.Ref(x))
	r.Add("Tme: stm.Tmes.At(inPkt.LstIdx()),")
	r.Addf("Flts: stm.%vsByTmeBnd(inPkt),", nodeFn.Title())
	r.Add("}")
	r.Add("for _, rx := range x.Rxs {")
	r.Addf("r = append(r, %v{Pkt:outPkt, Rx:rx})", _sys.Ana.TmeFlts.Tx.Adr(x))
	r.Add("}")
	r.Add("x.mu.Unlock()")
	r.Add("return r")
	return r
}
func (x *FleRltSide) StmRte(name string) (r *TypFn) {
	stm, name := _sys.Ana.Rlt.Stm, strings.Title(name)
	r = x.ElmNodeTypFn(name, k.Rte, "", stm, func(r *TypFn) {
		r.Node.FldPrnt(x)
		stm.AppendTest(r)
	})
	r.Add("x.Sub(r.Rx, r.Id)")
	r.Add("return r")
	stm.DstToInstr(r.Node)
	stm.Unsub(r.Node, false, true, func(rr *TypFn) {
		rr.Addf("x.%v.Unsub(x.Id)", r.Node.Prnt().Name)
	})
	// rx
	rx := stm.TypFn(k.Rx, r.Node)
	rx.InPrm(_sys.Ana.TmeFlts, "inPkt")
	rx.OutPrmSlice(_sys.Act, "r")
	rx.Add("x.mu.Lock()")
	rx.Add("if ana.Cfg.Trc.IsRltStm() {")
	rx.Addf("sys.Logf(\"%v(%%v).%v %%p inPkt %%v\", x.Prm(), x, inPkt)", r.Node.Full(), r.Name)
	rx.Add("}")
	rx.Addf("outPkt := %v{Tme:inPkt.Tme, Flt:inPkt.Flts%v}", _sys.Ana.TmeFlt.Ref(x), r.CallNode())
	rx.Add("for _, rx := range x.Rxs {")
	rx.Addf("r = append(r, %v{Pkt:outPkt, Rx:rx})", _sys.Ana.TmeFlt.Tx.Adr(x))
	rx.Add("}")
	rx.Add("x.mu.Unlock()")
	rx.Add("return r")
	return r
}
func (x *FleRltSide) Sar() (r *TypFn) {
	stm, name := _sys.Ana.Rlt.Stm, strings.Title(k.Sar)
	// node
	r = x.ElmNodeTypFn(name, "Rte1", "", stm, func(r *TypFn) {
		r.Node.FldPrnt(x)
		r.InPrm(_sys.Bsc.Flt, "afInc").LitVal("0.02")
		r.InPrm(_sys.Bsc.Flt, "afMax").LitVal("0.2")
		stm.AppendTest(r)
	})
	r.Node.Fld("IsLong", _sys.Bsc.Bol)
	r.Node.Fld("Sar", _sys.Bsc.Flt)
	r.Node.Fld("Ep", _sys.Bsc.Flt)
	r.Node.Fld("Af", _sys.Bsc.Flt)
	r.Node.Fld("PrvLo", _sys.Bsc.Flt)
	r.Node.Fld("PrvHi", _sys.Bsc.Flt)
	r.Addf("x.Sub(r.Rx, r.Id)")
	r.Add("return r")
	stm.DstToInstr(r.Node)
	stm.Unsub(r.Node, false, true, func(rr *TypFn) {
		rr.Addf("x.%v.Unsub(x.Id)", r.Node.Prnt().Name)
	})
	// RX IS MANUALLY CODED
	return r
}

func (x *FleRltSide) Ema() (r *TypFn) {
	stm, name := _sys.Ana.Rlt.Stm, strings.Title(k.Ema)
	r = x.ElmNodeTypFn(name, k.Rte, "", stm, func(r *TypFn) {
		r.Node.FldPrnt(x)
		r.Node.Fld("Prv", _sys.Bsc.Flt)
		stm.AppendTest(r)
	})
	r.Add("r.Prv = flt.Min")
	r.Addf("x.Sub(r.Rx, r.Id)")
	r.Add("return r")
	stm.DstToInstr(r.Node)
	stm.Unsub(r.Node, false, true, func(rr *TypFn) {
		rr.Addf("x.%v.Unsub(x.Id)", r.Node.Prnt().Name)
	})
	// rx
	rx := stm.TypFn(k.Rx, r.Node)
	rx.InPrm(_sys.Ana.TmeFlts, "inPkt")
	rx.OutPrmSlice(_sys.Act, "r")
	rx.Add("x.mu.Lock()")
	rx.Add("if ana.Cfg.Trc.IsRltStm() {")
	rx.Addf("sys.Logf(\"%v(%%v).%v %%p inPkt %%v\", x.Prm(), x, inPkt)", r.Node.Full(), rx.Name)
	rx.Add("}")
	rx.Addf("outPkt := %v{Tme:inPkt.Tme}", _sys.Ana.TmeFlt.Ref(x))
	rx.Add("if x.Prv == flt.Min {")
	rx.Add("outPkt.Flt = inPkt.Flts.Sma()")
	rx.Add("} else {")
	rx.Addf("outPkt.Flt = x.Prv + (inPkt.Flts.Lst() - x.Prv) * (flt.Flt(2) / flt.Flt(inPkt.Flts.Cnt() + 1))")
	rx.Add("}")
	rx.Add("x.Prv = outPkt.Flt")
	rx.Add("for _, rx := range x.Rxs {")
	rx.Addf("r = append(r, %v{Pkt:outPkt, Rx:rx})", _sys.Ana.TmeFlt.Tx.Adr(x))
	rx.Add("}")
	rx.Add("x.mu.Unlock()")
	rx.Add("return r")
	return r
}
