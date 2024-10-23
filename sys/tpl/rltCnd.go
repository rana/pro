package tpl

import (
	"strings"
	"sys/k"
	"sys/tpl/atr"
)

type (
	FleRltCnd struct {
		FleRltBse
	}
	FleRltCnds struct {
		FleBse
		PrtArr
	}
)

func (x *DirRlt) NewCnd() (r *FleRltCnd) {
	r = &FleRltCnd{}
	x.Cnd = r
	r.Name = k.Cnd
	r.Pkg = x.Pkg
	r.Ifc(r.Name, atr.TypAnaIfc)
	r.AddFle(r)
	return r
}
func (x *DirRlt) NewCnds() (r *FleRltCnds) {
	r = &FleRltCnds{}
	x.Cnds = r
	r.FleBse = *NewArr(x.Cnd, &r.PrtArr, x.Cnd.Pkg)
	r.AddFle(r)
	return r
}
func (x *FleRltCnd) InitTyp(bse *TypBse) {
	x.FleRltBse.InitTyp(bse)
	x.Typ().Bse().TestPth = append(_sys.Ana.Rlt.Stm.Typ().Bse().TestPth, &TestStp{
		MdlFst: func(r *PkgFn) { r.Add("cnd := tst.RltStmCndInrGtr(stm, 1)") },
	})
}
func (x *FleRltCnd) InitFld(s *Struct) {
	x.FleRltBse.InitFld(s)
	x.FldRxs(_sys.Bsc.Tme.Rxs)
}
func (x *FleRltCnd) InitIfc(i *Ifc) {
	x.FleRltBse.InitIfc(i)
	var sig *MemSig
	sig = x.MemSiga(k.Unsub, atr.None)
	sig.InPrm(Uint32, "id")
	sig.InPrmVariadic(Uint32, "slot")
	sig = x.MemSiga(k.DstToInstr, atr.None)
	sig.OutPrm(Int)
}
func (x *FleRltCnd) InitTypFn() {
	x.FleRltBse.InitTypFn()
	x.Sub(_sys.Bsc.Tme.Rx, true)
	x.CndAnd()
	x.CndSeq()
	x.Stgy()
	// x.StgyLong()
	// x.StgyShrt()
	// x.StgyLongPll()
	// x.CndMl()
}
func (x *FleRltCnd) AppendTest(r *TypFn) {
	if x.Test != nil {
		x.Test.Import(_sys.Ana.Hst)
		x.Test.Import(_sys.Lng.Pro.Act)
		r.T2.MdlLst.Add("mnr := tst.NewCndMnr(ap)")
		r.T2.MdlLst.Add("a.Sub(mnr.Rx, mnr.Id)")
		r.T2.MdlLst.Add("tst.IntegerEql(t, 1, len(a.Rxs), \"Sub Rxs\")")
		r.T2.MdlLst.Add("var actr act.Actr")
		r.T2.MdlLst.Add("vs := actr.RunHst(a.String())")
		r.T2.MdlLst.Addf("eHst := vs[len(vs)-1].(*hst.%v)", r.Node.Title())
		r.T2.MdlLst.Add("if eHst.Tmes != nil {")
		r.T2.MdlLst.Add("mnr.StartFor(instr.Instr(), eHst.Tmes.Cnt())")
		r.T2.MdlLst.Add("tst.TmesEql(t, eHst.Tmes, mnr.Tmes, \"Tmes\")")
		r.T2.MdlLst.Add("}")
		r.T2.MdlLst.Add("a.Unsub(mnr.Id)")
		r.T2.MdlLst.Add("tst.IntegerEql(t, 0, len(a.Rxs), \"Unsub Rxs\")")
	}
}
func (x *FleRltCnd) CndAnd() (r *TypFn) {
	cnd, name := _sys.Ana.Rlt.Cnd, strings.Title(k.And)
	// node
	r = x.ElmNodeTypFn(name, "Cnd1", "", cnd, func(r *TypFn) {
		r.Node.FldPrnt(x)
		r.Node.Fld("Tmes", _sys.Bsc.Tme.arr).Atr = atr.TstZeroSkp
		r.Node.Fld("TmesA", _sys.Bsc.Tme.arr).Atr = atr.TstZeroSkp
		r.InPrm(x, "a").LitVal("tst.RltStmCndInrGtr(stm, 2)")
		cnd.AppendTest(r)
	})
	r.Addf("r.Tmes = %v()", _sys.Bsc.Tme.arr.New.Ref(x))
	r.Addf("r.TmesA = %v()", _sys.Bsc.Tme.arr.New.Ref(x))
	r.Addf("x.Sub(r.Rx, r.Id)")
	r.Add("a.Sub(r.RxA, r.Id, SlotA)")
	r.Add("return r")
	cnd.DstToInstr(r.Node)
	cnd.Unsub(r.Node, false, true, func(rr *TypFn) {
		rr.Addf("x.%v.Unsub(x.Id)", r.Node.Prnt().Name)
		rr.Add("x.A.Unsub(x.Id, SlotA)")
	})
	// rx
	rx := cnd.TypFn(k.Rx, r.Node)
	rx.InPrm(_sys.Bsc.Tme, "inPkt")
	rx.OutPrmSlice(_sys.Act, "r")
	rx.Add("x.mu.Lock()")
	rx.Add("if ana.Cfg.Trc.IsRltCnd() {")
	rx.Addf("sys.Logf(\"%v(%%v).%v %%p inPkt %%v\", x.Prm(), x, inPkt)", r.Node.Full(), rx.Name)
	rx.Add("}")
	rx.Add("x.Tmes.Que(inPkt)")
	rx.Addf("r = x.Tx()")
	rx.Add("x.mu.Unlock()")
	rx.Add("return r")
	// rxA
	rxA := cnd.TypFn(k.RxA, r.Node)
	rxA.InPrm(_sys.Bsc.Tme, "inPkt")
	rxA.OutPrmSlice(_sys.Act, "r")
	rxA.Add("x.mu.Lock()")
	rxA.Add("if ana.Cfg.Trc.IsRltCnd() {")
	rxA.Addf("sys.Logf(\"%v(%%v).%v %%p inPkt %%v\", x.Prm(), x, inPkt)", r.Node.Full(), rxA.Name)
	rxA.Add("}")
	rxA.Add("x.TmesA.Que(inPkt)")
	rxA.Addf("r = x.Tx()")
	rxA.Add("x.mu.Unlock()")
	rxA.Add("return r")
	// tx
	tx := cnd.TypFn(k.Tx, r.Node)
	tx.OutPrmSlice(_sys.Act, "r")
	tx.Add("if len(*x.Tmes) == 0 || len(*x.TmesA) == 0 { // align")
	tx.Add("return nil")
	tx.Add("}")
	tx.Add("if (*x.Tmes)[0] != (*x.TmesA)[0] {")
	tx.Add("if (*x.Tmes)[0] < (*x.TmesA)[0] {")
	tx.Add("for len(*x.Tmes) > 0 && (*x.Tmes)[0] != (*x.TmesA)[0] { // drain X queue until empty or equal")
	tx.Add("x.Tmes.Dque()")
	tx.Add("}")
	tx.Add("} else {")
	tx.Add("for len(*x.TmesA) > 0 && (*x.Tmes)[0] != (*x.TmesA)[0] { // drain A queue until empty or equal")
	tx.Add("x.TmesA.Dque()")
	tx.Add("}")
	tx.Add("}")
	tx.Add("}")
	tx.Add("if len(*x.Tmes) == 0 || len(*x.TmesA) == 0 {")
	tx.Add("return nil")
	tx.Add("}")
	// front is later time
	// Inr Rlt Calc: earlier time - later time =
	// Subtract produces a series of values which are negative values when the series decreases
	// Divide produces a series of values which are ascending when the series values grow
	tx.Add("if (*x.Tmes)[0] == (*x.TmesA)[0] {")
	tx.Add("for _, rx := range x.Rxs {")
	tx.Addf("r = append(r, %v((*x.TmesA)[0], rx))", _sys.Bsc.Tme.NewTx.Ref(x))
	tx.Add("}")
	tx.Add("}")
	tx.Add("x.Tmes.Dque()")
	tx.Add("x.TmesA.Dque()")
	tx.Add("return r")
	return r
}

func (x *FleRltCnd) CndSeq() (r *TypFn) {
	cnd, name := _sys.Ana.Rlt.Cnd, strings.Title(k.Seq)
	// node
	r = x.ElmNodeTypFn(name, "Cnd2", "", cnd, func(r *TypFn) {
		r.Node.FldPrnt(x)
		r.Node.Fld("Tmes", _sys.Bsc.Tme.arr).Atr = atr.TstZeroSkp
		r.Node.Fld("TmesA", _sys.Bsc.Tme.arr).Atr = atr.TstZeroSkp
		r.InPrm(_sys.Bsc.Tme, "dur").LitVal("1")
		r.InPrm(x, "a").LitVal("tst.RltStmCndInrLss(stm, 2)")
		cnd.AppendTest(r)
	})
	r.Addf("r.Tmes = %v()", _sys.Bsc.Tme.arr.New.Ref(x))
	r.Addf("r.TmesA = %v()", _sys.Bsc.Tme.arr.New.Ref(x))
	r.Addf("x.Sub(r.Rx, r.Id)")
	r.Add("a.Sub(r.RxA, r.Id, SlotA)")
	r.Add("return r")
	cnd.DstToInstr(r.Node)
	cnd.Unsub(r.Node, false, true, func(rr *TypFn) {
		rr.Addf("x.%v.Unsub(x.Id)", r.Node.Prnt().Name)
		rr.Add("x.A.Unsub(x.Id, SlotA)")
	})
	// rx
	rx := cnd.TypFn(k.Rx, r.Node)
	rx.InPrm(_sys.Bsc.Tme, "inPkt")
	rx.OutPrmSlice(_sys.Act, "r")
	rx.Add("x.mu.Lock()")
	rx.Add("if ana.Cfg.Trc.IsRltCnd() {")
	rx.Addf("sys.Logf(\"%v(%%v).%v %%p inPkt %%v\", x.Prm(), x, inPkt)", r.Node.Full(), rx.Name)
	rx.Add("}")
	rx.Add("x.Tmes.Que(inPkt)")
	rx.Addf("r = x.Tx()")
	rx.Add("x.mu.Unlock()")
	rx.Add("return r")
	// rxA
	rxA := cnd.TypFn(k.RxA, r.Node)
	rxA.InPrm(_sys.Bsc.Tme, "inPkt")
	rxA.OutPrmSlice(_sys.Act, "r")
	rxA.Add("x.mu.Lock()")
	rxA.Add("if ana.Cfg.Trc.IsRltCnd() {")
	rxA.Addf("sys.Logf(\"%v(%%v).%v %%p inPkt %%v\", x.Prm(), x, inPkt)", r.Node.Full(), rxA.Name)
	rxA.Add("}")
	rxA.Add("x.TmesA.Que(inPkt)")
	rxA.Addf("r = x.Tx()")
	rxA.Add("x.mu.Unlock()")
	rxA.Add("return r")
	// tx
	tx := cnd.TypFn(k.Tx, r.Node)
	tx.OutPrmSlice(_sys.Act, "r")
	tx.Add("if len(*x.Tmes) == 0 || len(*x.TmesA) == 0 {")
	tx.Add("return nil")
	tx.Add("}")
	tx.Add("if (*x.Tmes)[0]+x.Dur != (*x.TmesA)[0] { // align; X & A must have GapFil")
	tx.Add("if (*x.Tmes)[0]+x.Dur < (*x.TmesA)[0] {")
	tx.Add("for len(*x.Tmes) > 1 && (*x.Tmes)[0]+x.Dur < (*x.TmesA)[0] { // drain X until empty or equal")
	tx.Add("x.Tmes.Dque()")
	tx.Add("}")
	tx.Add("} else if (*x.Tmes)[0]+x.Dur > (*x.TmesA)[0] {")
	tx.Add("for len(*x.TmesA) > 1 && (*x.Tmes)[0]+x.Dur > (*x.TmesA)[0] { // drain A until empty or equal")
	tx.Add("x.TmesA.Dque()")
	tx.Add("}")
	tx.Add("}")
	tx.Add("}")
	// front is later time
	// Inr Rlt Calc: earlier time - later time =
	// Subtract produces a series of values which are negative values when the series decreases
	// Divide produces a series of values which are ascending when the series values grow
	tx.Add("if (*x.Tmes)[0]+x.Dur == (*x.TmesA)[0] {")
	tx.Add("for _, rx := range x.Rxs {")
	tx.Addf("r = append(r, %v((*x.TmesA)[0], rx))", _sys.Bsc.Tme.NewTx.Ref(x))
	tx.Add("}")
	tx.Add("x.Tmes.Dque()")
	tx.Add("x.TmesA.Dque()")
	tx.Add("}")
	tx.Add("return r")
	return r
}

func (x *FleRltCnd) Stgy() (r *TypFn) {
	stgy, name := _sys.Ana.Rlt.Stgy, strings.Title(k.Stgy)
	// node
	r = x.ElmNodeTypFn(name, "", "", stgy, func(r *TypFn) {
		r.Node.Atr = r.Node.Atr &^ atr.Test // remove auto test; make manual test
		r.InPrm(_sys.Bsc.Bol, "isLong").LitVal("true").Atr = atr.FldSkp
		r.InPrm(_sys.Bsc.Flt, "prfLim").LitVal("1.1").Atr = atr.FldSkp // flds defined in bse
		r.InPrm(_sys.Bsc.Flt, "losLim").LitVal("1.1").Atr = atr.FldSkp
		r.InPrm(_sys.Bsc.Tme, "durLim").LitVal("3 * 60 * 60").Atr = atr.FldSkp
		r.InPrm(_sys.Bsc.Flt, "minPnlPct").LitVal("0.0").Atr = atr.FldSkp
		r.InPrm(_sys.Ana.Rlt.Instr, "instr").LitVal("instr").Atr = atr.FldSkp
		r.InPrm(_sys.Ana.Rlt.Stms, "ftrStms").Atr = atr.FldSkp
		r.InPrmVariadic(_sys.Ana.Rlt.Cnd, "clss").Atr = atr.FldSkp
	})

	r.Add("r.Key = r.String()")
	r.Add("if !ana.Cfg.Test {")
	r.Add("sys.Lrnr().LoadNetFromDsk(r.Key)")
	r.Add("}")
	r.Add("r.Cnd.Sub(r.RxOpn, r.Id)")
	r.Add("r.Instr.Sub(r.RxClsLim, r.Id)")
	r.Add("for _, cndCls := range r.Clss {")
	r.Add("cndCls.Sub(r.RxClsCnd, r.Id)")
	r.Add("}")
	r.Add("r.opns = tmes.New()")
	r.Add("r.stgyFtrStms = make([]*StgyFtrStm, ftrStms.Cnt())")
	r.Add("for n, ftrStm := range *ftrStms {")
	r.Add("v := &StgyFtrStm{}")
	r.Add("r.stgyFtrStms[n] = v")
	r.Add("v.Stgy = r")
	r.Add("v.Name = ftrStm.String()")
	r.Add("v.Tmes = tmes.New()")
	r.Add("v.Vals = flts.New()")
	r.Add("ftrStm.Sub(v.Rx, r.Id)")
	r.Add("}")

	r.Add("return r")
	// // rx
	// stgy.RxOpn(r)
	// stgy.RxOpnEnd(r)
	// stgy.RxClsLim(r)
	// stgy.RxClsLimEnd(r)
	// stgy.RxClsCnd(r)
	return r
}

// func (x *FleRltCnd) StgyLong() (r *TypFn) {
// 	stgy, name := _sys.Ana.Rlt.Stgy, strings.Title(k.Long)
// 	// node
// 	r = x.ElmNodeTypFn(name, "", "", stgy, func(r *TypFn) {
// 		r.InPrm(_sys.Bsc.Flt, "prfLim").LitVal("1.1").Atr = atr.FldSkp // flds defined in bse
// 		r.InPrm(_sys.Bsc.Flt, "losLim").LitVal("1.1").Atr = atr.FldSkp
// 		r.InPrm(_sys.Bsc.Tme, "durLim").LitVal("3 * 60 * 60").Atr = atr.FldSkp
// 		r.InPrm(_sys.Ana.Rlt.Instr, "instr").LitVal("instr").Atr = atr.FldSkp
// 		r.InPrmVariadic(_sys.Ana.Rlt.Cnd, "clss").Atr = atr.FldSkp
// 	})
// 	r.Add("r.Trds = ana.NewTrds()")
// 	r.Add("return r")
// 	// rx
// 	stgy.RxOpn(r)
// 	stgy.RxOpnEnd(r)
// 	stgy.RxClsLim(r)
// 	stgy.RxClsLimEnd(r)
// 	stgy.RxClsCnd(r)
// 	return r
// }

// func (x *FleRltCnd) StgyShrt() (r *TypFn) {
// 	stgy, name := _sys.Ana.Rlt.Stgy, strings.Title(k.Shrt)
// 	// node
// 	r = x.ElmNodeTypFn(name, "", "", stgy, func(r *TypFn) {
// 		r.InPrm(_sys.Bsc.Flt, "prfLim").LitVal("1.1").Atr = atr.FldSkp // flds defined in bse
// 		r.InPrm(_sys.Bsc.Flt, "losLim").LitVal("1.1").Atr = atr.FldSkp
// 		r.InPrm(_sys.Bsc.Tme, "durLim").LitVal("3 * 60 * 60").Atr = atr.FldSkp
// 		r.InPrm(_sys.Ana.Rlt.Instr, "instr").LitVal("instr").Atr = atr.FldSkp
// 		r.InPrmVariadic(_sys.Ana.Rlt.Cnd, "clss").Atr = atr.FldSkp
// 	})
// 	r.Add("r.Trds = ana.NewTrds()")
// 	r.Add("return r")
// 	// rx
// 	stgy.RxOpn(r)
// 	stgy.RxOpnEnd(r)
// 	stgy.RxClsLim(r)
// 	stgy.RxClsLimEnd(r)
// 	stgy.RxClsCnd(r)
// 	return r
// }

// func (x *FleRltCnd) StgyLongPll() (r *TypFn) {
// 	stgy, name := _sys.Ana.Rlt.Stgy, strings.Title(k.LongPll)
// 	// node
// 	r = x.ElmNodeTypFn(name, "", "", stgy, func(r *TypFn) {
// 		r.InPrm(_sys.Bsc.Flt, "prfLim").LitVal("1.1").Atr = atr.FldSkp // flds defined in bse
// 		r.InPrm(_sys.Bsc.Flt, "losLim").LitVal("1.1").Atr = atr.FldSkp
// 		r.InPrm(_sys.Bsc.Tme, "durLim").LitVal("3 * 60 * 60").Atr = atr.FldSkp
// 		r.InPrm(_sys.Ana.Rlt.Instr, "instr").LitVal("instr").Atr = atr.FldSkp
// 		r.InPrmVariadic(_sys.Ana.Rlt.Cnd, "clss").Atr = atr.FldSkp
// 	})
// 	r.Add("// EMPTY PLACEHOLDER FOR HST SYMMETRY")
// 	r.Add("// USED BY CND ML IN HST")
// 	r.Add("return r")

// 	rxOpn := stgy.TypFna("RxOpn", atr.None, r.Node) // for Stgy interface
// 	rxOpn.InPrm(_sys.Bsc.Tme, "inPkt")
// 	rxOpn.OutPrmSlice(_sys.Act, "r")
// 	rxOpn.Add("return r")
// 	rxClsLim := stgy.TypFna("RxClsLim", atr.None, r.Node) // for Stgy interface
// 	rxClsLim.InPrm(_sys.Ana.TmeIdx, "inPkt")
// 	rxClsLim.OutPrmSlice(_sys.Act, "r")
// 	rxClsLim.Add("return r")
// 	rxClsCnd := stgy.TypFna("RxClsCnd", atr.None, r.Node) // for Stgy interface
// 	rxClsCnd.InPrm(_sys.Bsc.Tme, "inPkt")
// 	rxClsCnd.OutPrmSlice(_sys.Act, "r")
// 	rxClsCnd.Add("return r")

// 	return r
// }

// func (x *FleRltCnd) CndMl() (r *TypFn) {
// 	cnd, name := _sys.Ana.Rlt.Cnd, strings.Title(k.Ml)
// 	// node
// 	r = x.ElmNodeTypFn(name, "Cnd3", "", cnd, func(r *TypFn) {
// 		r.Node.FldPrnt(x)
// 		r.Node.Fld("Tmes", _sys.Bsc.Tme.arr).Atr = atr.TstZeroSkp
// 		// r.Node.Fld("TmesA", _sys.Bsc.Tme.arr).Atr = atr.TstZeroSkp
// 		r.InPrm(_sys.Ana.Rlt.Stgy, "stgy")
// 		r.InPrm(_sys.Ana.Rlt.Stms, "stms")
// 		r.InPrm(_sys.Bsc.Flt, "pnlPctMin").LitVal("0.01")
// 		// cnd.AppendTest(r)
// 	})
// 	r.Addf("r.Tmes = %v()", _sys.Bsc.Tme.arr.New.Ref(x))
// 	// r.Addf("r.TmesA = %v()", _sys.Bsc.Tme.arr.New.Ref(x))
// 	r.Addf("x.Sub(r.Rx, r.Id)")
// 	// r.Add("a.Sub(r.RxA, r.Id, SlotA)")
// 	r.Add("return r")
// 	cnd.DstToInstr(r.Node)
// 	cnd.Unsub(r.Node, false, true, func(rr *TypFn) {
// 		rr.Addf("x.%v.Unsub(x.Id)", r.Node.Prnt().Name)
// 		// rr.Add("x.A.Unsub(x.Id, SlotA)")
// 	})
// 	// rx
// 	rx := cnd.TypFn(k.Rx, r.Node)
// 	rx.InPrm(_sys.Bsc.Tme, "inPkt")
// 	rx.OutPrmSlice(_sys.Act, "r")
// 	rx.Add("x.mu.Lock()")
// 	rx.Add("if ana.Cfg.Trc.IsRltCnd() {")
// 	rx.Addf("sys.Logf(\"%v(%%v).%v %%p inPkt %%v\", x.Prm(), x, inPkt)", r.Node.Full(), rx.Name)
// 	rx.Add("}")
// 	// rx.Add("x.Tmes.Que(inPkt)")
// 	// rx.Addf("r = x.Tx()")
// 	rx.Add("x.mu.Unlock()")
// 	rx.Add("return r")
// 	// // rxA
// 	// rxA := cnd.TypFn(k.RxA, r.Node)
// 	// rxA.InPrm(_sys.Bsc.Tme, "inPkt")
// 	// rxA.OutPrmSlice(_sys.Act, "r")
// 	// rxA.Add("x.mu.Lock()")
// 	// rxA.Add("if ana.Cfg.Trc.IsRltCnd() {")
// 	// rxA.Addf("sys.Logf(\"%v(%%v).%v %%p inPkt %%v\", x.Prm(), x, inPkt)", r.Node.Full(), rxA.Name)
// 	// rxA.Add("}")
// 	// rxA.Add("x.TmesA.Que(inPkt)")
// 	// rxA.Addf("r = x.Tx()")
// 	// rxA.Add("x.mu.Unlock()")
// 	// rxA.Add("return r")
// 	// // tx
// 	// tx := cnd.TypFn(k.Tx, r.Node)
// 	// tx.OutPrmSlice(_sys.Act, "r")
// 	// tx.Add("if len(*x.Tmes) == 0 || len(*x.TmesA) == 0 { // align")
// 	// tx.Add("return nil")
// 	// tx.Add("}")
// 	// tx.Add("if (*x.Tmes)[0] != (*x.TmesA)[0] {")
// 	// tx.Add("if (*x.Tmes)[0] < (*x.TmesA)[0] {")
// 	// tx.Add("for len(*x.Tmes) > 0 && (*x.Tmes)[0] != (*x.TmesA)[0] { // drain X queue until empty or equal")
// 	// tx.Add("x.Tmes.Dque()")
// 	// tx.Add("}")
// 	// tx.Add("} else {")
// 	// tx.Add("for len(*x.TmesA) > 0 && (*x.Tmes)[0] != (*x.TmesA)[0] { // drain A queue until empty or equal")
// 	// tx.Add("x.TmesA.Dque()")
// 	// tx.Add("}")
// 	// tx.Add("}")
// 	// tx.Add("}")
// 	// tx.Add("if len(*x.Tmes) == 0 || len(*x.TmesA) == 0 {")
// 	// tx.Add("return nil")
// 	// tx.Add("}")
// 	// // front is later time
// 	// // Inr Rlt Calc: earlier time - later time =
// 	// // Subtract produces a series of values which are negative values when the series decreases
// 	// // Divide produces a series of values which are ascending when the series values grow
// 	// tx.Add("if (*x.Tmes)[0] == (*x.TmesA)[0] {")
// 	// tx.Add("for _, rx := range x.Rxs {")
// 	// tx.Addf("r = append(r, %v((*x.TmesA)[0], rx))", _sys.Bsc.Tme.NewTx.Ref(x))
// 	// tx.Add("}")
// 	// tx.Add("}")
// 	// tx.Add("x.Tmes.Dque()")
// 	// tx.Add("x.TmesA.Dque()")
// 	// tx.Add("return r")
// 	return r
// }
