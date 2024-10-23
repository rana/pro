package tpl

import (
	"strings"
	"sys/k"
	"sys/tpl/atr"
)

type (
	FleHstCnd struct {
		FleHstBse
	}
	FleHstCnds struct {
		FleBse
		PrtArr
		PrtArrStrWrt // for Fbr.PrmWrt
		PrtString
	}
)

func (x *DirHst) NewCnd() (r *FleHstCnd) {
	r = &FleHstCnd{}
	x.Cnd = r
	r.Name = k.Cnd
	r.Pkg = x.Pkg
	r.Ifc(r.Name, atr.TypAnaIfc)
	r.AddFle(r)
	return r
}
func (x *DirHst) NewCnds() (r *FleHstCnds) {
	r = &FleHstCnds{}
	x.Cnds = r
	r.FleBse = *NewArr(x.Cnd, &r.PrtArr, x.Cnd.Pkg)
	r.AddFle(r)
	return r
}

func (x *FleHstCnd) InitTyp(bse *TypBse) {
	x.FleHstBse.InitTyp(bse)
	x.Typ().Bse().TestPth = append(_sys.Ana.Hst.Stm.Typ().Bse().TestPth, &TestStp{
		MdlFst: func(r *PkgFn) { r.Add("cnd := tst.HstStmCndInrGtr(stm, 1)") },
	})
	x.seg = x.NewSeg()
}

func (x *FleHstCnd) InitFld(s *Struct) {
	x.FleHstBse.InitFld(s)
	x.bse.Fld("Tmes", _sys.Bsc.Tme.arr).Atr = atr.Get | atr.TstZeroSkp
	x.seg.Fld("Tmes", _sys.Bsc.Tme.arr)
	x.seg.Fld("Out", _sys.Bsc.Tme.arr)
}

func (x *FleHstCnd) InitTypFn() {
	x.FleHstBse.InitTypFn()
	x.Import(_sys)             // for log
	x.Import(_sys.Ana)         // for trc cfg
	x.Import(_sys.Bsc.Bnd.arr) // for bnds.Segs
	x.CndAnd()
	x.CndSeq()
	x.Stgy()
}

func (x *FleHstCnd) CndAnd() (r *TypFn) {
	cnd, name := _sys.Ana.Hst.Cnd, strings.Title(k.And)
	r = x.ElmNodeTypFn(name, "Cnd1", "", cnd, func(r *TypFn) {
		r.InPrm(x, "a").LitVal(TstCndLitVal)
		r.Node.FldPrnt(x)
	})
	// seg
	seg := cnd.NodeSeg(r.Node.Name)
	seg.Fld("TmesA", _sys.Bsc.Tme.arr)
	segAct := cnd.TypFn(k.Act, seg)
	segAct.Add("for n := x.Idx; n < x.Lim; n++ {")
	segAct.Add("if x.TmesA.Has((*x.Tmes)[n]) {")
	segAct.Add("*x.Out = append(*x.Out, (*x.Tmes)[n])")
	segAct.Add("}")
	segAct.Add("}")
	// node
	r.Add("if ana.Cfg.Trc.IsHstCnd() {")
	r.Addf("sys.Logf(\"%%p %v(%%v)\", r, r.Prm())", r.Node.Full())
	r.Add("}")
	r.Add("if x.Tmes.Cnt() == 0 || r.A.Bse().Tmes.Cnt() == 0 {")
	r.Add("return r")
	r.Add("}")
	r.Addf("r.Tmes = %v()", _sys.Bsc.Tme.arr.New.Ref(x))
	r.Addf("segBnds, acts := %v(x.Tmes.Cnt())", _sys.Bsc.Bnd.arr.Segs.Ref(x))
	r.Add("for n, segBnd := range *segBnds {")
	r.Addf("seg := %v{}", seg.Adr(x))
	r.Add("seg.Bnd = segBnd")
	r.Add("seg.Tmes = x.Tmes")
	r.Add("seg.TmesA = r.A.Bse().Tmes")
	r.Addf("seg.Out = %v()", _sys.Bsc.Tme.arr.New.Ref(x))
	r.Add("acts[n] = seg")
	r.Add("}")
	r.Add("sys.Run().Pll(acts...)")
	r.Add("for _, act := range acts {")
	r.Addf("r.Tmes.Mrg(act.(%v).Out)", seg.Ref(x))
	r.Add("}")
	r.Add("return r")

	return r
}

func (x *FleHstCnd) CndSeq() (r *TypFn) {
	cnd, name := _sys.Ana.Hst.Cnd, strings.Title(k.Seq)
	r = x.ElmNodeTypFn(name, "Cnd2", "", cnd, func(r *TypFn) {
		r.InPrm(_sys.Bsc.Tme, "dur").LitVal("1")
		r.InPrm(x, "a").LitVal(TstCndLitVal2)
		r.Node.FldPrnt(x)
	})
	// seg
	seg := cnd.NodeSeg(r.Node.Name)
	seg.Fld("Dur", _sys.Bsc.Tme)
	seg.Fld("TmesA", _sys.Bsc.Tme.arr)
	segAct := cnd.TypFn(k.Act, seg)
	segAct.Add("for n := x.Idx; n < x.Lim; n++ {")
	segAct.Add("if x.TmesA.Has((*x.Tmes)[n]+x.Dur) {")
	segAct.Add("*x.Out = append(*x.Out, (*x.Tmes)[n]+x.Dur)")
	segAct.Add("}")
	segAct.Add("}")
	// node
	r.Add("if ana.Cfg.Trc.IsHstCnd() {")
	r.Addf("sys.Logf(\"%%p %v(%%v)\", r, r.Prm())", r.Node.Full())
	r.Add("}")
	r.Add("if x.Tmes.Cnt() == 0 || r.A.Bse().Tmes.Cnt() == 0 {")
	r.Add("return r")
	r.Add("}")
	r.Addf("r.Tmes = %v()", _sys.Bsc.Tme.arr.New.Ref(x))
	r.Addf("segBnds, acts := %v(x.Tmes.Cnt())", _sys.Bsc.Bnd.arr.Segs.Ref(x))
	r.Add("for n, segBnd := range *segBnds {")
	r.Addf("seg := %v{}", seg.Adr(x))
	r.Add("seg.Bnd = segBnd")
	r.Add("seg.Tmes = x.Tmes")
	r.Add("seg.TmesA = r.A.Bse().Tmes")
	r.Add("seg.Dur = r.Dur")
	r.Addf("seg.Out = %v()", _sys.Bsc.Tme.arr.New.Ref(x))
	r.Add("acts[n] = seg")
	r.Add("}")
	r.Add("sys.Run().Pll(acts...)")
	r.Add("for _, act := range acts {")
	r.Addf("r.Tmes.Mrg(act.(%v).Out)", seg.Ref(x))
	r.Add("}")
	r.Add("return r")
	return r
}

func (x *FleHstCnd) Stgy() (r *TypFn) {
	stgy, name := _sys.Ana.Hst.Stgy, strings.Title(k.Stgy)
	x.Import("sys/err")
	r = x.ElmNodeTypFn(name, "", "", stgy, func(r *TypFn) {
		r.Node.Atr = r.Node.Atr &^ atr.Test // remove auto test; make manual test
		r.InPrm(_sys.Bsc.Bol, "isLong").LitVal("true").Atr = atr.FldSkp
		r.InPrm(_sys.Bsc.Flt, "prfLim").LitVal("1.1").Atr = atr.FldSkp // flds defined in bse
		r.InPrm(_sys.Bsc.Flt, "losLim").LitVal("1.1").Atr = atr.FldSkp
		r.InPrm(_sys.Bsc.Tme, "durLim").LitVal("3 * 60 * 60").Atr = atr.FldSkp
		r.InPrm(_sys.Bsc.Flt, "minPnlPct").LitVal("0.0").Atr = atr.FldSkp
		r.InPrm(_sys.Ana.Hst.Instr, "instr").LitVal("instr").Atr = atr.FldSkp
		r.InPrm(_sys.Ana.Hst.Stms, "ftrStms")
		r.InPrmVariadic(_sys.Ana.Hst.Cnd, "clss").Atr = atr.FldSkp
	})
	// seg
	seg := stgy.NodeSeg(r.Node.Name)
	seg.Fld("Stgy", stgy.bse)
	seg.Fld("Tmes", _sys.Bsc.Tme.arr)
	seg.Fld("Out", _sys.Ana.Trds)
	segAct := stgy.TypFn(k.Act, seg)
	segAct.Add("for n := x.Idx; n < x.Lim; n++ {")
	segAct.Add("trd, rsnFail := x.Stgy.OpnClsTrd((*x.Tmes)[n])")
	segAct.Add("if trd != nil { // may fail to close due to near mkt opn, near mkt cls, spd lim exceeded")
	// segAct.Add("x.Out.CalcOpn(trd, x.Stgy.Instr.Bse().Ana)       // set trd flds")
	// segAct.Add("x.Out.CalcCls(trd, x.Stgy.Instr.Bse().Ana, false) // set trd flds")
	segAct.Add("x.Out.Push(trd)")
	segAct.Add("} else {")
	segAct.Add("if rsnFail == ana.NoCls {")
	segAct.Add("break // exit last opn fail to mirror rlt behavior; single ana.NoCls expected")
	segAct.Add("}")
	segAct.Add("}")
	segAct.Add("}")
	// node
	r.Add("if ana.Cfg.Trc.IsHstStgy() {")
	r.Addf("sys.Logf(\"%%p %v(%%v)\", r, r.Prm())", r.Node.Full())
	r.Add("}")
	r.Add("if ftrStms.Cnt() == 0 {")
	r.Add("err.Panic(\"ftrStms is empty\")")
	r.Add("}")
	r.Add("if x.Tmes.Cnt() == 0 {")
	r.Add("return r")
	r.Add("}")
	// r.Add("r.Trds = ana.NewTrds()")

	// r.Add("")
	// r.Add("r.Calc()")
	// r.Add("")
	// r.Add("")

	r.Add("return r")
	return r
}
