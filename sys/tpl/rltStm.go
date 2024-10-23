package tpl

import (
	"strings"
	"sys"
	"sys/k"
	"sys/ks"
	"sys/tpl/atr"
)

type (
	FleRltStm struct {
		FleRltBse
	}
	FleRltStms struct {
		FleBse
		PrtArr
		PrtArrStrWrt // for Fbr.PrmWrt (CndCnd3Ml usage)
		PrtString
	}
)

func (x *DirRlt) NewStm() (r *FleRltStm) {
	r = &FleRltStm{}
	x.Stm = r
	r.Name = k.Stm
	r.Pkg = x.Pkg
	r.Ifc(r.Name, atr.TypAnaIfc)
	r.AddFle(r)
	return r
}
func (x *DirRlt) NewStms() (r *FleRltStms) {
	r = &FleRltStms{}
	x.Stms = r
	r.FleBse = *NewArr(x.Stm, &r.PrtArr, x.Stm.Pkg)
	r.AddFle(r)
	return r
}
func (x *FleRltStm) InitTyp(bse *TypBse) {
	x.FleRltBse.InitTyp(bse)
	x.Typ().Bse().TestPth = append(_sys.Ana.Rlt.Side.Typ().Bse().TestPth, &TestStp{
		MdlFst: func(r *PkgFn) { r.Add("stm := tst.RltSideStmRteLst(side)") },
	})
}
func (x *FleRltStm) InitFld(s *Struct) {
	x.FleRltBse.InitFld(s)
	x.FldRxs(_sys.Ana.TmeFlt.Rxs)
}
func (x *FleRltStm) InitIfc(i *Ifc) {
	x.FleRltBse.InitIfc(i)
	var sig *MemSig
	sig = x.MemSiga(k.Unsub, atr.None)
	sig.InPrm(Uint32, "id")
	sig.InPrmVariadic(Uint32, "slot")
	sig = x.MemSiga(k.DstToInstr, atr.None)
	sig.OutPrm(Int)
}
func (x *FleRltStm) InitTypFn() {
	x.FleRltBse.InitTypFn()
	x.Sub(_sys.Ana.TmeFlt.Rx, true)
	for _, una := range ks.Unas {
		x.StmUna(una)
	}
	for _, scl := range ks.Scls {
		x.StmScl(scl)
	}
	for _, sel := range ks.Sels {
		x.StmSel(sel)
	}
	for _, agg := range ks.Aggs {
		x.StmAgg(agg)
	}
	x.StmAggEma()
	for _, inr := range ks.Inrs {
		x.StmInr(inr)
	}
	x.StmInrSlp()
	for _, otr := range ks.Otrs {
		x.StmOtr(otr)
	}
	for _, cndScl := range ks.CndScls {
		x.CndScl(cndScl)
	}
	for _, cndInr := range ks.CndInrs {
		x.CndInr(cndInr)
	}
	for _, cndOtr := range ks.CndOtrs {
		x.CndOtr(cndOtr)
	}
}

func (x *FleRltStm) AppendTest(r *TypFn) {
	if x.Test != nil {
		x.Test.Import(_sys.Ana.Hst)
		x.Test.Import(_sys.Lng.Pro.Act)
		r.T2.MdlLst.Add("mnr := tst.NewStmMnr(ap)")
		r.T2.MdlLst.Add("a.Sub(mnr.Rx, mnr.Id)")
		r.T2.MdlLst.Add("tst.IntegerEql(t, 1, len(a.Rxs), \"Sub Rxs\")")
		r.T2.MdlLst.Add("var actr act.Actr")
		r.T2.MdlLst.Add("vs := actr.RunHst(a.String())")
		r.T2.MdlLst.Addf("eHst := vs[len(vs)-1].(*hst.%v)", r.Node.Title())
		r.T2.MdlLst.Add("if eHst.Tmes != nil {")
		r.T2.MdlLst.Add("mnr.StartFor(instr.Instr(), eHst.Tmes.Cnt())")
		r.T2.MdlLst.Add("tst.TmesEql(t, eHst.Tmes, mnr.Tmes, \"Tmes\")")
		r.T2.MdlLst.Add("tst.FltsEql(t, eHst.Vals, mnr.Vals, \"Vals\")")
		r.T2.MdlLst.Add("}")
		r.T2.MdlLst.Add("a.Unsub(mnr.Id)")
		r.T2.MdlLst.Add("tst.IntegerEql(t, 0, len(a.Rxs), \"Unsub Rxs\")")
	}
}
func (x *FleRltStm) StmUna(name string) (r *TypFn) {
	stm, name := _sys.Ana.Rlt.Stm, strings.Title(name)
	// node
	r = x.ElmNodeTypFn(name, k.Una, k.Una, stm, func(r *TypFn) {
		r.Node.FldPrnt(x)
		stm.AppendTest(r)
	})
	r.Addf("x.Sub(r.Rx, r.Id)")
	r.Add("return r")
	stm.DstToInstr(r.Node)
	stm.Unsub(r.Node, false, true, func(rr *TypFn) {
		rr.Addf("x.%v.Unsub(x.Id)", r.Node.Prnt().Name)
	})
	// rx
	rx := stm.TypFn(k.Rx, r.Node)
	rx.InPrm(_sys.Ana.TmeFlt, "inPkt")
	rx.OutPrmSlice(_sys.Act, "r")
	rx.Add("x.mu.Lock()")
	rx.Add("if ana.Cfg.Trc.IsRltStm() {")
	rx.Addf("sys.Logf(\"%v(%%v).%v %%p inPkt %%v\", x.Prm(), x, inPkt)", r.Node.Full(), rx.Name)
	rx.Add("}")
	rx.Addf("outPkt := %v{Tme:inPkt.Tme, Flt:inPkt.Flt.%v()}", _sys.Ana.TmeFlt.Ref(x), name)
	rx.Add("for _, rx := range x.Rxs {")
	rx.Addf("r = append(r, %v{Pkt:outPkt, Rx:rx})", _sys.Ana.TmeFlt.Tx.Adr(x))
	rx.Add("}")
	rx.Add("x.mu.Unlock()")
	rx.Add("return r")
	return r
}

func (x *FleRltStm) StmScl(name string) (r *TypFn) {
	stm, name := _sys.Ana.Rlt.Stm, strings.Title(name)
	// node
	r = x.ElmNodeTypFn(name, k.Scl, k.Scl, stm, func(r *TypFn) {
		r.Node.FldPrnt(x)
		r.InPrm(_sys.Bsc.Flt, "scl").LitVal("1.1")
		stm.AppendTest(r)
	})
	r.Addf("x.Sub(r.Rx, r.Id)")
	r.Add("return r")
	stm.DstToInstr(r.Node)
	stm.Unsub(r.Node, false, true, func(rr *TypFn) {
		rr.Addf("x.%v.Unsub(x.Id)", r.Node.Prnt().Name)
	})
	// rx
	rx := stm.TypFn(k.Rx, r.Node)
	rx.InPrm(_sys.Ana.TmeFlt, "inPkt")
	rx.OutPrmSlice(_sys.Act, "r")
	rx.Add("x.mu.Lock()")
	rx.Add("if ana.Cfg.Trc.IsRltStm() {")
	rx.Addf("sys.Logf(\"%v(%%v).%v %%p inPkt %%v\", x.Prm(), x, inPkt)", r.Node.Full(), rx.Name)
	rx.Add("}")
	rx.Addf("outPkt := %v{Tme:inPkt.Tme, Flt:inPkt.Flt.%v(x.Scl)}", _sys.Ana.TmeFlt.Ref(x), name)
	rx.Add("for _, rx := range x.Rxs {")
	rx.Addf("r = append(r, %v{Pkt:outPkt, Rx:rx})", _sys.Ana.TmeFlt.Tx.Adr(x))
	rx.Add("}")
	rx.Add("x.mu.Unlock()")
	rx.Add("return r")
	return r
}

func (x *FleRltStm) StmSel(name string) (r *TypFn) {
	stm, name := _sys.Ana.Rlt.Stm, strings.Title(name)
	// node
	r = x.ElmNodeTypFn(name, k.Sel, k.Sel, stm, func(r *TypFn) {
		r.Node.FldPrnt(x)
		r.InPrm(_sys.Bsc.Flt, "sel").LitVal("1.3171") // 1.3171 common in tst data
		stm.AppendTest(r)
	})
	r.Addf("x.Sub(r.Rx, r.Id)")
	r.Add("return r")
	stm.DstToInstr(r.Node)
	stm.Unsub(r.Node, false, true, func(rr *TypFn) {
		rr.Addf("x.%v.Unsub(x.Id)", r.Node.Prnt().Name)
	})
	// rx
	rx := x.TypFn(k.Rx, r.Node)
	rx.InPrm(_sys.Ana.TmeFlt, "inPkt")
	rx.OutPrmSlice(_sys.Act, "r")
	rx.Add("x.mu.Lock()")
	rx.Add("if ana.Cfg.Trc.IsRltStm() {")
	rx.Addf("sys.Logf(\"%v(%%v).%v %%p inPkt %%v\", x.Prm(), x, inPkt)", r.Node.Full(), rx.Name)
	rx.Add("}")
	rx.Addf("outPkt := %v{Tme:inPkt.Tme, Flt:inPkt.Flt.%v(x.Sel)}", _sys.Ana.TmeFlt.Ref(x), strings.Title(sys.CnjSel(name)))
	rx.Add("for _, rx := range x.Rxs {")
	rx.Addf("r = append(r, %v{Pkt:outPkt, Rx:rx})", _sys.Ana.TmeFlt.Tx.Adr(x))
	rx.Add("}")
	rx.Add("x.mu.Unlock()")
	rx.Add("return r")
	return r
}

func (x *FleRltStm) StmAgg(name string) (r *TypFn) {
	stm, name := _sys.Ana.Rlt.Stm, strings.Title(name)
	// node
	r = x.ElmNodeTypFn(name, k.Agg, k.Agg, stm, func(r *TypFn) {
		r.Node.FldPrnt(x)
		r.Node.Fld("Tmes", _sys.Bsc.Tme.arr).Atr = atr.TstZeroSkp
		r.Node.Fld("Vals", _sys.Bsc.Flt.arr).Atr = atr.TstZeroSkp
		r.InPrm(_sys.Bsc.Unt, "length").LitVal("2") // length
		stm.AppendTest(r)
	})
	r.Addf("r.Tmes = %v()", _sys.Bsc.Tme.arr.New.Ref(x))
	r.Addf("r.Vals = %v()", _sys.Bsc.Flt.arr.New.Ref(x))
	r.Addf("x.Sub(r.Rx, r.Id)")
	r.Add("return r")
	stm.DstToInstr(r.Node)
	stm.Unsub(r.Node, false, true, func(rr *TypFn) {
		rr.Addf("x.%v.Unsub(x.Id)", r.Node.Prnt().Name)
	})
	// rx
	rx := x.TypFn(k.Rx, r.Node)
	rx.InPrm(_sys.Ana.TmeFlt, "inPkt")
	rx.OutPrmSlice(_sys.Act, "r")
	rx.Add("if x.Length > 0 {")
	rx.Add("x.mu.Lock()")
	rx.Add("if ana.Cfg.Trc.IsRltStm() {")
	rx.Addf("sys.Logf(\"%v(%%v).%v %%p inPkt %%v\", x.Prm(), x, inPkt)", r.Node.Full(), rx.Name)
	rx.Add("}")
	rx.Add("x.Tmes.Que(inPkt.Tme)")
	rx.Add("x.Vals.Que(inPkt.Flt)")
	rx.Add("if x.Tmes.Cnt() >= x.Length {")
	rx.Addf("outPkt := %v{", _sys.Ana.TmeFlt.Ref(x))
	rx.Add("Tme: x.Tmes.At(x.Length - 1),")
	rx.Addf("Flt: x.Vals.To(x.Length).%v(),", name)
	rx.Add("}")

	rx.Add("x.Tmes.Dque()")
	rx.Add("x.Vals.Dque()")
	rx.Add("for _, rx := range x.Rxs {")
	rx.Addf("r = append(r, %v{Pkt:outPkt, Rx:rx})", _sys.Ana.TmeFlt.Tx.Adr(x))
	rx.Add("}")
	rx.Add("}")
	rx.Add("x.mu.Unlock()")
	rx.Add("}")
	rx.Add("return r")
	return r
}

func (x *FleRltStm) StmAggEma() (r *TypFn) {
	stm, name := _sys.Ana.Rlt.Stm, strings.Title(k.Ema)
	// node
	r = x.ElmNodeTypFn(name, k.Agg, k.Agg, stm, func(r *TypFn) {
		r.Node.FldPrnt(x)
		r.Node.Fld("Prv", _sys.Bsc.Flt).Atr = atr.TstZeroSkp
		r.Node.Fld("Tmes", _sys.Bsc.Tme.arr).Atr = atr.TstZeroSkp
		r.Node.Fld("Vals", _sys.Bsc.Flt.arr).Atr = atr.TstZeroSkp
		r.InPrm(_sys.Bsc.Unt, "length").LitVal("9") // length
		stm.AppendTest(r)
	})
	r.Addf("r.Tmes = %v()", _sys.Bsc.Tme.arr.New.Ref(x))
	r.Addf("r.Vals = %v()", _sys.Bsc.Flt.arr.New.Ref(x))
	r.Add("r.Prv = flt.Min")
	r.Addf("x.Sub(r.Rx, r.Id)")
	r.Add("return r")
	stm.DstToInstr(r.Node)
	stm.Unsub(r.Node, false, true, func(rr *TypFn) {
		rr.Addf("x.%v.Unsub(x.Id)", r.Node.Prnt().Name)
	})
	// rx
	// x.Import("math")
	rx := x.TypFn(k.Rx, r.Node)
	rx.InPrm(_sys.Ana.TmeFlt, "inPkt")
	rx.OutPrmSlice(_sys.Act, "r")
	rx.Add("if x.Length > 0 {")
	rx.Add("x.mu.Lock()")
	rx.Add("if ana.Cfg.Trc.IsRltStm() {")
	rx.Addf("sys.Logf(\"%v(%%v).%v %%p inPkt %%v\", x.Prm(), x, inPkt)", r.Node.Full(), rx.Name)
	rx.Add("}")
	rx.Add("x.Tmes.Que(inPkt.Tme)")
	rx.Add("x.Vals.Que(inPkt.Flt)")

	rx.Add("if x.Tmes.Cnt() >= x.Length {")

	rx.Addf("outPkt := %v{", _sys.Ana.TmeFlt.Ref(x))
	rx.Add("Tme: x.Tmes.At(x.Length - 1),")
	rx.Add("}")
	rx.Add("vals := x.Vals.To(x.Length)")
	rx.Add("if x.Prv == flt.Min {")
	rx.Add("outPkt.Flt = vals.Sma()")
	rx.Add("} else {")
	rx.Addf("outPkt.Flt = x.Prv + (vals.Lst() - x.Prv) * (flt.Flt(2) / flt.Flt(vals.Cnt() + 1))")
	rx.Add("}")
	rx.Add("x.Prv = outPkt.Flt")

	rx.Add("x.Tmes.Dque()")
	rx.Add("x.Vals.Dque()")
	rx.Add("for _, rx := range x.Rxs {")
	rx.Addf("r = append(r, %v{Pkt:outPkt, Rx:rx})", _sys.Ana.TmeFlt.Tx.Adr(x))
	rx.Add("}")
	rx.Add("}")
	rx.Add("x.mu.Unlock()")
	rx.Add("}")
	rx.Add("return r")
	return r
}

func (x *FleRltStm) StmInr(name string) (r *TypFn) {
	stm, name := _sys.Ana.Rlt.Stm, strings.Title(name)
	// node
	r = x.ElmNodeTypFn(name, k.Inr, k.Inr, stm, func(r *TypFn) {
		r.Node.FldPrnt(x)
		r.Node.Fld("Tmes", _sys.Bsc.Tme.arr).Atr = atr.TstZeroSkp
		r.Node.Fld("Vals", _sys.Bsc.Flt.arr).Atr = atr.TstZeroSkp
		r.InPrm(_sys.Bsc.Unt, "off").LitVal("1")
		stm.AppendTest(r)
	})
	r.Addf("r.Tmes = %v()", _sys.Bsc.Tme.arr.New.Ref(x))
	r.Addf("r.Vals = %v()", _sys.Bsc.Flt.arr.New.Ref(x))
	r.Addf("x.Sub(r.Rx, r.Id)")
	r.Add("return r")
	stm.DstToInstr(r.Node)
	stm.Unsub(r.Node, false, true, func(rr *TypFn) {
		rr.Addf("x.%v.Unsub(x.Id)", r.Node.Prnt().Name)
	})
	// rx
	rx := x.TypFn(k.Rx, r.Node)
	rx.InPrm(_sys.Ana.TmeFlt, "inPkt")
	rx.OutPrmSlice(_sys.Act, "r")
	rx.Add("if x.Off > 0 {")
	rx.Add("x.mu.Lock()")
	rx.Add("if ana.Cfg.Trc.IsRltStm() {")
	rx.Addf("sys.Logf(\"%v(%%v).%v %%p inPkt %%v\", x.Prm(), x, inPkt)", r.Node.Full(), rx.Name)
	rx.Add("}")
	rx.Add("x.Tmes.Que(inPkt.Tme)")
	rx.Add("x.Vals.Que(inPkt.Flt)")

	rx.Add("if x.Tmes.Cnt() > x.Off {")
	rx.Addf("outPkt := %v{", _sys.Ana.TmeFlt.Ref(x))
	rx.Add("Tme: x.Tmes.At(x.Off),")
	rx.Addf("Flt: x.Vals.At(x.Off).%v(x.Vals.Fst()),", name)
	rx.Add("}")

	rx.Add("x.Tmes.Dque()")
	rx.Add("x.Vals.Dque()")
	rx.Add("for _, rx := range x.Rxs {")
	rx.Addf("r = append(r, %v{Pkt:outPkt, Rx:rx})", _sys.Ana.TmeFlt.Tx.Adr(x))
	rx.Add("}")
	rx.Add("}")
	rx.Add("x.mu.Unlock()")
	rx.Add("}")
	rx.Add("return r")
	return r
}

func (x *FleRltStm) StmInrSlp() (r *TypFn) {
	stm, name := _sys.Ana.Rlt.Stm, strings.Title(k.Slp)
	// node
	r = x.ElmNodeTypFn(name, "Inr1", k.Inr, stm, func(r *TypFn) {
		r.Node.FldPrnt(x)
		r.Node.Fld("Tmes", _sys.Bsc.Tme.arr).Atr = atr.TstZeroSkp
		r.Node.Fld("Vals", _sys.Bsc.Flt.arr).Atr = atr.TstZeroSkp
		r.InPrm(_sys.Bsc.Unt, "off").LitVal("1")
		stm.AppendTest(r)
	})
	r.Addf("r.Tmes = %v()", _sys.Bsc.Tme.arr.New.Ref(x))
	r.Addf("r.Vals = %v()", _sys.Bsc.Flt.arr.New.Ref(x))
	r.Addf("x.Sub(r.Rx, r.Id)")
	r.Add("return r")
	stm.DstToInstr(r.Node)
	stm.Unsub(r.Node, false, true, func(rr *TypFn) {
		rr.Addf("x.%v.Unsub(x.Id)", r.Node.Prnt().Name)
	})
	// rx
	rx := x.TypFn(k.Rx, r.Node)
	rx.InPrm(_sys.Ana.TmeFlt, "inPkt")
	rx.OutPrmSlice(_sys.Act, "r")
	rx.Add("if x.Off > 0 {")
	rx.Add("x.mu.Lock()")
	rx.Add("if ana.Cfg.Trc.IsRltStm() {")
	rx.Addf("sys.Logf(\"%v(%%v).%v %%p inPkt %%v\", x.Prm(), x, inPkt)", r.Node.Full(), rx.Name)
	rx.Add("}")
	rx.Add("x.Tmes.Que(inPkt.Tme)")
	rx.Add("x.Vals.Que(inPkt.Flt)")

	rx.Add("if x.Tmes.Cnt() > x.Off {")
	rx.Add("rise := x.Vals.At(x.Off).Sub(x.Vals.Fst())")
	rx.Add("run := x.Tmes.At(x.Off).Sub(x.Tmes.Fst())")
	rx.Addf("outPkt := %v{", _sys.Ana.TmeFlt.Ref(x))
	rx.Add("Tme: x.Tmes.At(x.Off),")
	rx.Addf("Flt: rise/%v(run),", _sys.Bsc.Flt.Ref(x))
	rx.Add("}")

	rx.Add("x.Tmes.Dque()")
	rx.Add("x.Vals.Dque()")
	rx.Add("for _, rx := range x.Rxs {")
	rx.Addf("r = append(r, %v{Pkt:outPkt, Rx:rx})", _sys.Ana.TmeFlt.Tx.Adr(x))
	rx.Add("}")
	rx.Add("}")
	rx.Add("x.mu.Unlock()")
	rx.Add("}")
	rx.Add("return r")
	return r
}

func (x *FleRltStm) StmOtr(name string) (r *TypFn) {
	stm, name := _sys.Ana.Rlt.Stm, strings.Title(name)
	// node
	r = x.ElmNodeTypFn(name, k.Otr, k.Otr, stm, func(r *TypFn) {
		r.Node.FldPrnt(x)
		r.Node.Fld("Tmes", _sys.Bsc.Tme.arr).Atr = atr.TstZeroSkp
		r.Node.Fld("Vals", _sys.Bsc.Flt.arr).Atr = atr.TstZeroSkp
		r.Node.Fld("TmesA", _sys.Bsc.Tme.arr).Atr = atr.TstZeroSkp
		r.Node.Fld("ValsA", _sys.Bsc.Flt.arr).Atr = atr.TstZeroSkp
		r.InPrm(_sys.Bsc.Unt, "off").LitVal("1")
		r.InPrm(x, "a").LitVal("tst.RltStmStmInrAdd(stm, 2)")
		stm.AppendTest(r)
	})
	r.Addf("r.Tmes = %v()", _sys.Bsc.Tme.arr.New.Ref(x))
	r.Addf("r.Vals = %v()", _sys.Bsc.Flt.arr.New.Ref(x))
	r.Addf("r.TmesA = %v()", _sys.Bsc.Tme.arr.New.Ref(x))
	r.Addf("r.ValsA = %v()", _sys.Bsc.Flt.arr.New.Ref(x))
	r.Addf("x.Sub(r.Rx, r.Id)")
	r.Add("a.Sub(r.RxA, r.Id, SlotA)")
	r.Add("return r")
	stm.DstToInstr(r.Node)
	stm.Unsub(r.Node, false, true, func(rr *TypFn) {
		rr.Addf("x.%v.Unsub(x.Id)", r.Node.Prnt().Name)
	})
	// rx
	rx := x.TypFn(k.Rx, r.Node)
	rx.InPrm(_sys.Ana.TmeFlt, "inPkt")
	rx.OutPrmSlice(_sys.Act, "r")
	rx.Add("x.mu.Lock()")
	rx.Add("if ana.Cfg.Trc.IsRltStm() {")
	rx.Addf("sys.Logf(\"%v(%%v).%v %%p inPkt %%v\", x.Prm(), x, inPkt)", r.Node.Full(), rx.Name)
	rx.Add("}")
	rx.Add("x.Tmes.Que(inPkt.Tme)")
	rx.Add("x.Vals.Que(inPkt.Flt)")
	rx.Addf("r = x.Tx()")
	rx.Add("x.mu.Unlock()")
	rx.Add("return r")
	// rxA
	rxA := x.TypFn(k.RxA, r.Node)
	rxA.InPrm(_sys.Ana.TmeFlt, "inPkt")
	rxA.OutPrmSlice(_sys.Act, "r")
	rxA.Add("x.mu.Lock()")
	rxA.Add("if ana.Cfg.Trc.IsRltStm() {")
	rxA.Addf("sys.Logf(\"%v.%v %%p inPkt %%v\", x, inPkt)", r.Node.Full(), rxA.Name)
	rxA.Add("}")
	rxA.Add("x.TmesA.Que(inPkt.Tme)")
	rxA.Add("x.ValsA.Que(inPkt.Flt)")
	rxA.Addf("r = x.Tx()")
	rxA.Add("x.mu.Unlock()")
	rxA.Add("return r")
	// tx
	tx := x.TypFn(k.Tx, r.Node)
	tx.OutPrmSlice(_sys.Act, "r")
	tx.Add("if x.Tmes.Cnt() == 0 || x.TmesA.Cnt() <= x.Off { // align")
	tx.Add("return nil")
	tx.Add("}")
	tx.Add("if x.Tmes.At(0) != x.TmesA.At(0) {")
	tx.Add("if x.Tmes.At(0) < x.TmesA.At(0) {")
	tx.Add("for x.Tmes.Cnt() > 0 && x.Tmes.At(0) != x.TmesA.At(0) { // drain X queue until empty or equal")
	tx.Add("x.Tmes.Dque()")
	tx.Add("x.Vals.Dque()")
	tx.Add("}")
	tx.Add("} else {")
	tx.Add("for x.TmesA.Cnt() > 0 && x.Tmes.At(0) != x.TmesA.At(0) { // drain A queue until empty or equal")
	tx.Add("x.TmesA.Dque()")
	tx.Add("x.ValsA.Dque()")
	tx.Add("}")
	tx.Add("}")
	tx.Add("}")
	tx.Add("if x.Tmes.Cnt() > 0 && x.TmesA.Cnt() > x.Off {")
	tx.Addf("outPkt := %v{", _sys.Ana.TmeFlt.Ref(x))
	tx.Add("Tme: x.TmesA.At(x.Off),")
	tx.Addf("Flt: x.Vals.Fst().%v(x.ValsA.At(x.Off)),", name)
	tx.Add("}")
	tx.Add("x.Tmes.Dque()")
	tx.Add("x.Vals.Dque()")
	tx.Add("x.TmesA.Dque()")
	tx.Add("x.ValsA.Dque()")
	tx.Add("for _, rx := range x.Rxs {")
	tx.Addf("r = append(r, %v{Pkt:outPkt, Rx:rx})", _sys.Ana.TmeFlt.Tx.Adr(x))
	tx.Add("}")
	tx.Add("}")
	tx.Add("return r")
	return r
}

func (x *FleRltStm) CndScl(name string) (r *TypFn) {
	cnd, name := _sys.Ana.Rlt.Cnd, strings.Title(name)
	// node
	r = x.ElmNodeTypFn(name, k.Scl, k.Scl, cnd, func(r *TypFn) {
		r.Node.FldPrnt(x)
		r.InPrm(_sys.Bsc.Flt, "scl").LitVal("1.1")
		cnd.AppendTest(r)
	})
	r.Addf("x.Sub(r.Rx, r.Id)")
	r.Add("return r")
	cnd.DstToInstr(r.Node)
	cnd.Unsub(r.Node, false, true, func(rr *TypFn) {
		rr.Addf("x.%v.Unsub(x.Id)", r.Node.Prnt().Name)
	})
	// rx
	rx := cnd.TypFn(k.Rx, r.Node)
	rx.InPrm(_sys.Ana.TmeFlt, "inPkt")
	rx.OutPrmSlice(_sys.Act, "r")
	rx.Add("x.mu.Lock()")
	rx.Add("if ana.Cfg.Trc.IsRltCnd() {")
	rx.Addf("sys.Logf(\"%v(%%v).%v %%p inPkt %%v\", x.Prm(), x, inPkt)", r.Node.Full(), rx.Name)
	rx.Add("}")
	rx.Addf("if inPkt.Flt.%v(x.Scl) {", name)
	rx.Add("for _, rx := range x.Rxs {")
	rx.Addf("r = append(r, %v{Pkt:inPkt.Tme, Rx:rx})", _sys.Bsc.Tme.Tx.Adr(x))
	rx.Add("}")
	rx.Add("}")
	rx.Add("x.mu.Unlock()")
	rx.Add("return r")
	return r
}

func (x *FleRltStm) CndInr(name string) (r *TypFn) {
	cnd, name := _sys.Ana.Rlt.Cnd, strings.Title(name)
	// node
	r = x.ElmNodeTypFn(name, k.Inr, k.Inr, cnd, func(r *TypFn) {
		r.Node.FldPrnt(x)
		r.Node.Fld("Tmes", _sys.Bsc.Tme.arr).Atr = atr.TstZeroSkp
		r.Node.Fld("Vals", _sys.Bsc.Flt.arr).Atr = atr.TstZeroSkp
		r.InPrm(_sys.Bsc.Unt, "off").LitVal("1")
		cnd.AppendTest(r)
	})
	r.Addf("r.Tmes = %v()", _sys.Bsc.Tme.arr.New.Ref(x))
	r.Addf("r.Vals = %v()", _sys.Bsc.Flt.arr.New.Ref(x))
	r.Addf("x.Sub(r.Rx, r.Id)")
	r.Add("return r")
	cnd.DstToInstr(r.Node)
	cnd.Unsub(r.Node, false, true, func(rr *TypFn) {
		rr.Addf("x.%v.Unsub(x.Id)", r.Node.Prnt().Name)
	})
	// rx
	rx := cnd.TypFn(k.Rx, r.Node)
	rx.InPrm(_sys.Ana.TmeFlt, "inPkt")
	rx.OutPrmSlice(_sys.Act, "r")
	rx.Add("if x.Off > 0 {")
	rx.Add("x.mu.Lock()")
	rx.Add("if ana.Cfg.Trc.IsRltCnd() {")
	rx.Addf("sys.Logf(\"%v(%%v).%v %%p inPkt %%v\", x.Prm(), x, inPkt)", r.Node.Full(), rx.Name)
	rx.Add("}")
	rx.Add("x.Tmes.Que(inPkt.Tme)")
	rx.Add("x.Vals.Que(inPkt.Flt)")
	rx.Add("if x.Tmes.Cnt() > x.Off {")
	rx.Addf("if x.Vals.At(x.Off).%v(x.Vals.Fst()) {", name) // TODO: INLINE
	rx.Add("for _, rx := range x.Rxs {")
	rx.Addf("r = append(r, %v{Pkt:(*x.Tmes)[x.Off], Rx:rx})", _sys.Bsc.Tme.Tx.Adr(x))
	rx.Add("}")
	rx.Add("}")
	rx.Add("x.Tmes.Dque()")
	rx.Add("x.Vals.Dque()")
	rx.Add("}")
	rx.Add("x.mu.Unlock()")
	rx.Add("}")
	rx.Add("return r")
	return r
}

func (x *FleRltStm) CndOtr(name string) (r *TypFn) {
	cnd, name := _sys.Ana.Rlt.Cnd, strings.Title(name)
	// node
	r = x.ElmNodeTypFn(name, k.Otr, k.Otr, cnd, func(r *TypFn) {
		r.Node.FldPrnt(x)
		r.Node.Fld("Tmes", _sys.Bsc.Tme.arr).Atr = atr.TstZeroSkp
		r.Node.Fld("Vals", _sys.Bsc.Flt.arr).Atr = atr.TstZeroSkp
		r.Node.Fld("TmesA", _sys.Bsc.Tme.arr).Atr = atr.TstZeroSkp
		r.Node.Fld("ValsA", _sys.Bsc.Flt.arr).Atr = atr.TstZeroSkp
		r.InPrm(_sys.Bsc.Unt, "off").LitVal("1")
		r.InPrm(x, "a").LitVal("tst.RltStmStmInrAdd(stm, 2)")
		cnd.AppendTest(r)
	})
	r.Addf("r.Tmes = %v()", _sys.Bsc.Tme.arr.New.Ref(x))
	r.Addf("r.Vals = %v()", _sys.Bsc.Flt.arr.New.Ref(x))
	r.Addf("r.TmesA = %v()", _sys.Bsc.Tme.arr.New.Ref(x))
	r.Addf("r.ValsA = %v()", _sys.Bsc.Flt.arr.New.Ref(x))
	r.Addf("x.Sub(r.Rx, r.Id)")
	r.Add("a.Sub(r.RxA, r.Id, SlotA)")
	r.Add("return r")
	cnd.DstToInstr(r.Node)
	cnd.Unsub(r.Node, false, true, func(rr *TypFn) {
		rr.Addf("x.%v.Unsub(x.Id)", r.Node.Prnt().Name)
	})
	// rx
	rx := cnd.TypFn(k.Rx, r.Node)
	rx.InPrm(_sys.Ana.TmeFlt, "inPkt")
	rx.OutPrmSlice(_sys.Act, "r")
	rx.Add("x.mu.Lock()")
	rx.Add("if ana.Cfg.Trc.IsRltCnd() {")
	rx.Addf("sys.Logf(\"%v(%%v).%v %%p inPkt %%v\", x.Prm(), x, inPkt)", r.Node.Full(), rx.Name)
	rx.Add("}")
	rx.Add("x.Tmes.Que(inPkt.Tme)")
	rx.Add("x.Vals.Que(inPkt.Flt)")
	rx.Addf("r = x.Tx()")
	rx.Add("x.mu.Unlock()")
	rx.Add("return r")
	// rxA
	rxA := cnd.TypFn(k.RxA, r.Node)
	rxA.InPrm(_sys.Ana.TmeFlt, "inPkt")
	rxA.OutPrmSlice(_sys.Act, "r")
	rxA.Add("x.mu.Lock()")
	rxA.Add("if ana.Cfg.Trc.IsRltCnd() {")
	rxA.Addf("sys.Logf(\"%v(%%v).%v %%p inPkt %%v\", x.Prm(), x, inPkt)", r.Node.Full(), rxA.Name)
	rxA.Add("}")
	rxA.Add("x.TmesA.Que(inPkt.Tme)")
	rxA.Add("x.ValsA.Que(inPkt.Flt)")
	rxA.Addf("r = x.Tx()")
	rxA.Add("x.mu.Unlock()")
	rxA.Add("return r")
	// tx
	tx := cnd.TypFn(k.Tx, r.Node)
	tx.OutPrmSlice(_sys.Act, "r")
	tx.Add("if x.Tmes.Cnt() == 0 || x.TmesA.Cnt() <= x.Off { // align")
	tx.Add("return nil")
	tx.Add("}")
	tx.Add("if x.Tmes.At(0) != x.TmesA.At(0) {")
	tx.Add("if x.Tmes.At(0) < x.TmesA.At(0) {")
	tx.Add("for x.Tmes.Cnt() > 0 && x.Tmes.At(0) != x.TmesA.At(0) { // drain X queue until empty or equal")
	tx.Add("x.Tmes.Dque()")
	tx.Add("x.Vals.Dque()")
	tx.Add("}")
	tx.Add("} else {")
	tx.Add("for x.TmesA.Cnt() > 0 && x.Tmes.At(0) != x.TmesA.At(0) { // drain A queue until empty or equal")
	tx.Add("x.TmesA.Dque()")
	tx.Add("x.ValsA.Dque()")
	tx.Add("}")
	tx.Add("}")
	tx.Add("}")
	tx.Add("if x.Tmes.Cnt() > 0 && x.TmesA.Cnt() > x.Off {")
	tx.Addf("if x.Vals.Fst().%v(x.ValsA.At(x.Off)) {", name)
	tx.Add("for _, rx := range x.Rxs {")
	tx.Addf("r = append(r, %v{Pkt:(*x.TmesA)[x.Off], Rx:rx})", _sys.Bsc.Tme.Tx.Adr(x))
	tx.Add("}")
	tx.Add("}")
	tx.Add("x.Tmes.Dque()")
	tx.Add("x.Vals.Dque()")
	tx.Add("x.TmesA.Dque()")
	tx.Add("x.ValsA.Dque()")
	tx.Add("}")
	tx.Add("return r")
	return r
}
