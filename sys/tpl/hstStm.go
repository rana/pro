package tpl

import (
	"strings"
	"sys"
	"sys/k"
	"sys/ks"
	"sys/tpl/atr"
)

const (
	TstStmLitVal  = "tst.HstStmStmInrAdd(tst.HstSideStmRteLst(tst.HstInrvlSideBid(tst.HstInstrInrvlI(instr, 10))), 4)"
	TstCndLitVal  = "tst.HstStmCndInrGtr(tst.HstStmStmInrAdd(tst.HstSideStmRteLst(tst.HstInrvlSideBid(tst.HstInstrInrvlI(instr, 10))), 1), 2)"
	TstCndLitVal2 = "tst.HstStmCndInrLss(tst.HstStmStmInrAdd(tst.HstSideStmRteLst(tst.HstInrvlSideBid(tst.HstInstrInrvlI(instr, 10))), 1), 2)"
)

type (
	FleHstStm struct {
		FleHstBse
	}
	FleHstStms struct {
		FleBse
		PrtArr
		PrtArrStrWrt // for Fbr.PrmWrt
		PrtString
	}
)

func (x *DirHst) NewStm() (r *FleHstStm) {
	r = &FleHstStm{}
	x.Stm = r
	r.Name = k.Stm
	r.Pkg = x.Pkg
	r.Ifc(r.Name, atr.TypAnaIfc)
	r.AddFle(r)
	return r
}
func (x *DirHst) NewStms() (r *FleHstStms) {
	r = &FleHstStms{}
	x.Stms = r
	r.FleBse = *NewArr(x.Stm, &r.PrtArr, x.Stm.Pkg)
	r.AddFle(r)
	return r
}

func (x *FleHstStm) InitTyp(bse *TypBse) {
	x.FleHstBse.InitTyp(bse)
	x.Typ().Bse().TestPth = append(_sys.Ana.Hst.Side.Typ().Bse().TestPth, &TestStp{
		MdlFst: func(r *PkgFn) { r.Add("stm := tst.HstSideStmRteLst(side)") },
	})
	x.seg = x.NewSeg()
}

func (x *FleHstStm) InitFld(s *Struct) {
	x.FleHstBse.InitFld(s)
	x.bse.Fld("Tmes", _sys.Bsc.Tme.arr).Atr = atr.Get
	x.bse.Fld("Vals", _sys.Bsc.Flt.arr).Atr = atr.Get
	x.seg.Fld("Vals", _sys.Bsc.Flt.arr)
	x.seg.Fld("Out", _sys.Bsc.Flt.arr)
}
func (x *FleHstStm) InitTypFn() {
	x.FleHstBse.InitTypFn()
	x.Import("sys/err")
	x.Import(_sys)             // for log
	x.Import(_sys.Ana)         // for trc cfg
	x.Import(_sys.Bsc.Bnd.arr) // for bnds.Segs
	x.At()
	x.Atf()
	// x.Splt()
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
func (x *FleHstStm) At() (r *TypFn) {
	r = x.TypFn(k.At, x.bse)
	r.InPrm(_sys.Bsc.Tme.arr, "ts")
	r.OutPrm(_sys.Bsc.Flt.arr, "r")
	r.Add("if  x.Vals == nil || len(*x.Vals) == 0 || ts == nil || len(*ts) == 0 {")
	r.Addf("return %v()", _sys.Bsc.Flt.arr.New.Ref(x))
	r.Add("}")
	r.Add("r = flts.MakeEmp(ts.Cnt())")
	r.Addf("segBnds, acts := %v(ts.Cnt())", _sys.Bsc.Bnd.arr.Segs.Ref(x))
	r.Add("for n, segBnd := range *segBnds {")
	r.Add("seg := &StmAtSeg{}")
	r.Add("seg.Bnd = segBnd")
	r.Add("seg.Stm = x")
	r.Add("seg.AtTmes = ts")
	r.Add("seg.Out = flts.MakeEmp(segBnd.Cnt())")
	r.Add("acts[n] = seg")
	r.Add("}")
	r.Add("sys.Run().Pll(acts...)")
	r.Add("for _, act := range acts {")
	r.Add("r.Mrg(act.(*StmAtSeg).Out)")
	r.Add("}")
	r.Add("return r")
	x.MemSigFn(r) // add to interface
	return r
}
func (x *FleHstStm) Atf() (r *TypFn) {
	r = x.TypFn("Atf", x.bse)
	r.InPrm(_sys.Bsc.Tme.arr, "ts")
	r.OutPrmSlice(Float32, "r")
	r.Add("if  x.Vals == nil || len(*x.Vals) == 0 || ts == nil || len(*ts) == 0 {")
	r.Add("return r")
	r.Add("}")
	r.Add("r = make([]float32, 0, ts.Cnt())")
	r.Addf("segBnds, acts := %v(ts.Cnt())", _sys.Bsc.Bnd.arr.Segs.Ref(x))
	r.Add("for n, segBnd := range *segBnds {")
	r.Add("seg := &StmAtfSeg{}")
	r.Add("seg.Bnd = segBnd")
	r.Add("seg.Stm = x")
	r.Add("seg.AtTmes = ts")
	r.Add("seg.Out = make([]float32, 0, segBnd.Cnt())")
	r.Add("acts[n] = seg")
	r.Add("}")
	r.Add("sys.Run().Pll(acts...)")
	r.Add("for _, act := range acts {")
	r.Add("r = append(r, act.(*StmAtfSeg).Out...)")
	r.Add("}")
	r.Add("return r")
	x.MemSigFn(r) // add to interface
	return r
}

func (x *FleHstStm) StmUna(name string) (r *TypFn) {
	stm, name := _sys.Ana.Hst.Stm, strings.Title(name)
	r = x.ElmNodeTypFn(name, k.Una, k.Una, stm, func(r *TypFn) {
		r.Node.FldPrnt(x)
	})
	// seg
	seg := stm.NodeSeg(r.Node.Name)
	segAct := stm.TypFn(k.Act, seg)
	segAct.Add("for n := x.Idx; n < x.Lim; n++ {")
	fn := _sys.Bsc.Flt.Typ().Bse().Fn(name)
	fn.Inline(segAct, stm, "(*x.Vals)[n]", "(*x.Out)[n]")
	segAct.Add("}")
	// node
	r.Add("if ana.Cfg.Trc.IsHstStm() {")
	r.Addf("sys.Logf(\"%%p %v(%%v)\", r, r.Prm())", r.Node.Full())
	r.Add("}")
	r.Add("if x.Vals.Cnt() == 0 {")
	r.Add("return r")
	r.Add("}")
	r.Add("if len(*x.Tmes) != len(*x.Vals) {")
	r.Add("err.Panicf(\"length unequal (Tmes:%v Vals:%v)\", len(*x.Tmes), len(*x.Vals))")
	r.Add("}")
	r.Add("r.Tmes = x.Tmes")
	r.Addf("r.Vals = %v(x.Vals.Cnt())", _sys.Bsc.Flt.arr.Make.Ref(x))
	r.Addf("segBnds, acts := %v(r.Vals.Cnt())", _sys.Bsc.Bnd.arr.Segs.Ref(x))
	r.Add("for n, segBnd := range *segBnds {")
	r.Addf("seg := %v{}", seg.Adr(x))
	r.Add("seg.Bnd = segBnd")
	r.Add("seg.Vals = x.Vals")
	r.Add("seg.Out = r.Vals")
	r.Add("acts[n] = seg")
	r.Add("}")
	r.Add("sys.Run().Pll(acts...)")
	r.Add("return r")
	return r
}
func (x *FleHstStm) StmScl(name string) (r *TypFn) {
	stm, name := _sys.Ana.Hst.Stm, strings.Title(name)
	r = x.ElmNodeTypFn(name, k.Scl, k.Scl, stm, func(r *TypFn) {
		r.InPrm(_sys.Bsc.Flt, "scl").LitVal("1.1")
		r.Node.FldPrnt(x)
	})
	// seg
	seg := stm.NodeSeg(r.Node.Name)
	seg.Fld("Scl", _sys.Bsc.Flt)
	segAct := stm.TypFn(k.Act, seg)
	segAct.Add("for n := x.Idx; n < x.Lim; n++ {")
	fn := _sys.Bsc.Flt.Typ().Bse().Fn(name)
	fn.Inline(segAct, stm, "(*x.Vals)[n]", "(*x.Out)[n]", "x.Scl")
	segAct.Add("}")
	// node
	r.Add("if ana.Cfg.Trc.IsHstStm() {")
	r.Addf("sys.Logf(\"%%p %v(%%v)\", r, r.Prm())", r.Node.Full())
	r.Add("}")
	r.Add("if x.Vals.Cnt() == 0 {")
	r.Add("return r")
	r.Add("}")
	r.Add("if len(*x.Tmes) != len(*x.Vals) {")
	r.Add("err.Panicf(\"length unequal (Tmes:%v Vals:%v)\", len(*x.Tmes), len(*x.Vals))")
	r.Add("}")
	r.Add("r.Tmes = x.Tmes")
	r.Addf("r.Vals = %v(x.Vals.Cnt())", _sys.Bsc.Flt.arr.Make.Ref(x))
	r.Addf("segBnds, acts := %v(r.Vals.Cnt())", _sys.Bsc.Bnd.arr.Segs.Ref(x))
	r.Add("for n, segBnd := range *segBnds {")
	r.Addf("seg := %v{}", seg.Adr(x))
	r.Add("seg.Bnd = segBnd")
	r.Add("seg.Vals = x.Vals")
	r.Add("seg.Out = r.Vals")
	r.Add("seg.Scl = r.Scl")
	r.Add("acts[n] = seg")
	r.Add("}")
	r.Add("sys.Run().Pll(acts...)")
	r.Add("return r")
	return r
}

func (x *FleHstStm) StmSel(name string) (r *TypFn) {
	stm, name := _sys.Ana.Hst.Stm, strings.Title(name)
	r = x.ElmNodeTypFn(name, k.Sel, k.Sel, stm, func(r *TypFn) {
		r.InPrm(_sys.Bsc.Flt, "sel").LitVal("1.3171") // 1.3171 common in tst data
		r.Node.FldPrnt(x)
	})
	// seg
	seg := stm.NodeSeg(r.Node.Name)
	seg.Fld("Sel", _sys.Bsc.Flt)
	segAct := stm.TypFn(k.Act, seg)
	segAct.Add("for n := x.Idx; n < x.Lim; n++ {")
	fn := _sys.Bsc.Flt.Typ().Bse().Fn(sys.CnjSel(name))
	fn.Inline(segAct, stm, "(*x.Vals)[n]", "(*x.Out)[n]", "x.Sel")
	segAct.Add("}")
	// node
	r.Add("if ana.Cfg.Trc.IsHstStm() {")
	r.Addf("sys.Logf(\"%%p %v(%%v)\", r, r.Prm())", r.Node.Full())
	r.Add("}")
	r.Add("if x.Vals.Cnt() == 0 {")
	r.Add("return r")
	r.Add("}")
	r.Add("if len(*x.Tmes) != len(*x.Vals) {")
	r.Add("err.Panicf(\"length unequal (Tmes:%v Vals:%v)\", len(*x.Tmes), len(*x.Vals))")
	r.Add("}")
	r.Add("r.Tmes = x.Tmes")
	r.Addf("r.Vals = %v(x.Vals.Cnt())", _sys.Bsc.Flt.arr.Make.Ref(x))
	r.Addf("segBnds, acts := %v(r.Vals.Cnt())", _sys.Bsc.Bnd.arr.Segs.Ref(x))
	r.Add("for n, segBnd := range *segBnds {")
	r.Addf("seg := %v{}", seg.Adr(x))
	r.Add("seg.Bnd = segBnd")
	r.Add("seg.Vals = x.Vals")
	r.Add("seg.Out = r.Vals")
	r.Add("seg.Sel = r.Sel")
	r.Add("acts[n] = seg")
	r.Add("}")
	r.Add("sys.Run().Pll(acts...)")
	r.Add("return r")
	return r
}

func (x *FleHstStm) StmAgg(name string) (r *TypFn) {
	stm, name := _sys.Ana.Hst.Stm, strings.Title(name)
	r = x.ElmNodeTypFn(name, k.Agg, k.Agg, stm, func(r *TypFn) {
		r.InPrm(_sys.Bsc.Unt, "length").LitVal("2")
		r.Node.FldPrnt(x)
	})
	// seg
	seg := stm.NodeSeg(r.Node.Name)
	seg.Fld("Length", _sys.Bsc.Unt)
	segAct := stm.TypFn(k.Act, seg)
	segAct.Add("for n := x.Idx; n < x.Lim; n++ {")
	segAct.Addf("(*x.Out)[n] = x.Vals.In(n, n+x.Length).%v()", name)
	segAct.Add("}")
	// node
	r.Add("if ana.Cfg.Trc.IsHstStm() {")
	r.Addf("sys.Logf(\"%%p %v(%%v)\", r, r.Prm())", r.Node.Full())
	r.Add("}")
	r.Add("if x.Vals.Cnt() == 0 || x.Vals.Cnt() < r.Length || r.Length == 0 {")
	r.Add("return r")
	r.Add("}")
	r.Add("if len(*x.Tmes) != len(*x.Vals) {")
	r.Add("err.Panicf(\"length unequal (Tmes:%v Vals:%v)\", len(*x.Tmes), len(*x.Vals))")
	r.Add("}")
	r.Add("r.Tmes = x.Tmes.From(r.Length-1)")
	r.Add("r.Vals = flts.Make(x.Vals.Cnt()-r.Length+1)")
	r.Addf("segBnds, acts := %v(r.Vals.Cnt())", _sys.Bsc.Bnd.arr.Segs.Ref(x))
	r.Add("for n, segBnd := range *segBnds {")
	r.Addf("seg := %v{}", seg.Adr(x))
	r.Add("seg.Bnd = segBnd")
	r.Add("seg.Vals = x.Vals")
	r.Add("seg.Out = r.Vals")
	r.Add("seg.Length = r.Length")
	r.Add("acts[n] = seg")
	r.Add("}")
	r.Add("sys.Run().Pll(acts...)")
	r.Add("return r")
	return r
}

func (x *FleHstStm) StmAggEma() (r *TypFn) {
	stm, name := _sys.Ana.Hst.Stm, strings.Title(k.Ema)
	r = x.ElmNodeTypFn(name, k.Agg, k.Agg, stm, func(r *TypFn) {
		r.InPrm(_sys.Bsc.Unt, "length").LitVal("2")
		r.Node.FldPrnt(x)
	})
	// node
	r.Add("if ana.Cfg.Trc.IsHstStm() {")
	r.Addf("sys.Logf(\"%%p %v(%%v)\", r, r.Prm())", r.Node.Full())
	r.Add("}")
	r.Add("if x.Vals.Cnt() == 0 || x.Vals.Cnt() < r.Length || r.Length == 0  {")
	r.Add("return r")
	r.Add("}")
	r.Add("if len(*x.Tmes) != len(*x.Vals) {")
	r.Add("err.Panicf(\"length unequal (Tmes:%v Vals:%v)\", len(*x.Tmes), len(*x.Vals))")
	r.Add("}")
	r.Add("r.Tmes = x.Tmes.From(r.Length-1)")
	r.Add("r.Vals = flts.Make(x.Vals.Cnt()-r.Length+1)")
	r.Add("if len(*r.Tmes) != len(*r.Vals) {")
	r.Add("err.Panicf(\"r length unequal (Tmes:%v Vals:%v)\", len(*r.Tmes), len(*r.Vals))")
	r.Add("}")
	r.Add("// NON-PLL IMPL DUE TO PRV VAL CHAINING")
	r.Add("//    EMA CALC   https://stockcharts.com/school/doku.php?id=chart_school:technical_indicators:moving_averages")
	r.Add("// Initial SMA: 10-period sum / 10")
	r.Add("// Multiplier: (2 / (Time periods + 1) ) = (2 / (10 + 1) ) = 0.1818 (18.18%)")
	r.Add("// EMA: {Close - EMA(previous day)} x multiplier + EMA(previous day)")
	r.Add("scl := flt.Flt(2) / flt.Flt(r.Length + 1)")
	r.Add("(*r.Vals)[0] = x.Vals.To(r.Length).Sma()")
	r.Add("for n := 1; n < len(*r.Vals); n++ {")
	r.Addf("(*r.Vals)[n] = ((*x.Vals)[n+int(r.Length)-1] - (*r.Vals)[n-1]) * scl + (*r.Vals)[n-1]")
	r.Add("}")
	r.Add("return r")
	return r
}

func (x *FleHstStm) StmInr(name string) (r *TypFn) {
	stm, name := _sys.Ana.Hst.Stm, strings.Title(name)
	r = x.ElmNodeTypFn(name, k.Inr, k.Inr, stm, func(r *TypFn) {
		r.InPrm(_sys.Bsc.Unt, "off").LitVal("1")
		r.Node.FldPrnt(x)
		if test := stm.Test; test != nil {
			r.T2.MdlLst.Addf("x, e := stm.Bse(), %v{}", r.Node.Adr(test))
			r.T2.MdlLst.Addf(r.InPrms[0].AsnTest())
			r.T2.MdlLst.Addf("e.Tmes = x.Tmes.From(e.Off)")
			r.T2.MdlLst.Addf("e.Vals = %v(e.Tmes.Cnt())", _sys.Bsc.Flt.arr.Make.Ref(test))
			r.T2.MdlLst.Addf("for n := %v; n < e.Tmes.Cnt(); n++ {", _sys.Bsc.Unt.Zero.Ref(test))
			r.T2.MdlLst.Addf("(*e.Vals)[n] = x.Vals.At(n+e.Off).%v(x.Vals.At(n))", name)
			r.T2.MdlLst.Add("}")
			r.T2.MdlLst.Addf("tst.TmesEql(t, e.Tmes, a.Tmes, \"%v.Tmes\")", r.Node.Name)
			r.T2.MdlLst.Addf("tst.FltsEql(t, e.Vals, a.Vals, \"%v.Vals\")", r.Node.Name)
		}
	})
	// seg
	seg := stm.NodeSeg(r.Node.Name)
	seg.Fld("Off", _sys.Bsc.Unt)
	segAct := stm.TypFn(k.Act, seg)
	segAct.Add("for n := x.Idx; n < x.Lim; n++ {")
	segAct.Addf("(*x.Out)[n] = x.Vals.At(n+x.Off).%v((*x.Vals)[n])", name)
	segAct.Add("}")
	// node
	r.Add("if ana.Cfg.Trc.IsHstStm() {")
	r.Addf("sys.Logf(\"%%p %v(%%v)\", r, r.Prm())", r.Node.Full())
	r.Add("}")
	r.Add("if x.Vals.Cnt() == 0 || x.Vals.Cnt() <= r.Off || r.Off == 0{")
	r.Add("return r")
	r.Add("}")
	r.Add("if len(*x.Tmes) != len(*x.Vals) {")
	r.Add("err.Panicf(\"length unequal (Tmes:%v Vals:%v)\", len(*x.Tmes), len(*x.Vals))")
	r.Add("}")
	r.Add("r.Tmes = x.Tmes.From(r.Off)")
	r.Add("r.Vals = flts.Make(x.Vals.Cnt()-r.Off)")
	r.Addf("segBnds, acts := %v(r.Vals.Cnt())", _sys.Bsc.Bnd.arr.Segs.Ref(x))
	r.Add("for n, segBnd := range *segBnds {")
	r.Addf("seg := %v{}", seg.Adr(x))
	r.Add("seg.Bnd = segBnd")
	r.Add("seg.Vals = x.Vals")
	r.Add("seg.Out = r.Vals")
	r.Add("seg.Off = r.Off")
	r.Add("acts[n] = seg")
	r.Add("}")
	r.Add("sys.Run().Pll(acts...)")
	r.Add("return r")
	return r
}

func (x *FleHstStm) StmInrSlp() (r *TypFn) {
	stm, name := _sys.Ana.Hst.Stm, strings.Title(k.Slp)
	r = x.ElmNodeTypFn(name, "Inr1", k.Inr, stm, func(r *TypFn) {
		r.InPrm(_sys.Bsc.Unt, "off").LitVal("1")
		r.Node.FldPrnt(x)
	})
	// seg
	seg := stm.NodeSeg(r.Node.Name)
	seg.Fld("Off", _sys.Bsc.Unt)
	seg.Fld("Tmes", _sys.Bsc.Tme.arr)
	segAct := stm.TypFn(k.Act, seg)
	segAct.Add("for n := x.Idx; n < x.Lim; n++ { // expects 32-bit tme with 1s resolution")
	segAct.Addf("(*x.Out)[n] = ((*x.Vals)[n+x.Off] - (*x.Vals)[n]) / %v((*x.Tmes)[n+x.Off] - (*x.Tmes)[n])", _sys.Bsc.Flt.Ref(stm))
	segAct.Add("}")
	// node
	r.Add("if ana.Cfg.Trc.IsHstStm() {")
	r.Addf("sys.Logf(\"%%p %v(%%v)\", r, r.Prm())", r.Node.Full())
	r.Add("}")
	r.Add("if x.Vals.Cnt() == 0 || x.Vals.Cnt() <= r.Off || r.Off == 0{")
	r.Add("return r")
	r.Add("}")
	r.Add("if len(*x.Tmes) != len(*x.Vals) {")
	r.Add("err.Panicf(\"length unequal (Tmes:%v Vals:%v)\", len(*x.Tmes), len(*x.Vals))")
	r.Add("}")
	r.Add("r.Tmes = x.Tmes.From(r.Off)")
	r.Add("r.Vals = flts.Make(x.Vals.Cnt()-r.Off)")
	r.Addf("segBnds, acts := %v(r.Vals.Cnt())", _sys.Bsc.Bnd.arr.Segs.Ref(x))
	r.Add("for n, segBnd := range *segBnds {")
	r.Addf("seg := %v{}", seg.Adr(x))
	r.Add("seg.Bnd = segBnd")
	r.Add("seg.Tmes = x.Tmes")
	r.Add("seg.Vals = x.Vals")
	r.Add("seg.Out = r.Vals")
	r.Add("seg.Off = r.Off")
	r.Add("acts[n] = seg")
	r.Add("}")
	r.Add("sys.Run().Pll(acts...)")
	r.Add("return r")
	return r
}

func (x *FleHstStm) StmOtr(name string) (r *TypFn) {
	stm, name := _sys.Ana.Hst.Stm, strings.Title(name)
	r = x.ElmNodeTypFn(name, k.Otr, k.Otr, stm, func(r *TypFn) {
		r.InPrm(_sys.Bsc.Unt, "off").LitVal("1")
		r.InPrm(x, "a").LitVal(TstStmLitVal)
		r.Node.FldPrnt(x)
	})
	// seg
	seg := stm.NodeSeg(r.Node.Name)
	seg.Fld("Off", _sys.Bsc.Unt)
	seg.Fld("ValsA", _sys.Bsc.Flt.arr)
	segAct := stm.TypFn(k.Act, seg)
	segAct.Add("for n := x.Idx; n < x.Lim; n++ {")
	segAct.Addf("(*x.Out)[n] = x.Vals.At(n).%v((*x.ValsA)[n+x.Off])", strings.Title(name))
	segAct.Add("}")
	// node
	r.Add("if ana.Cfg.Trc.IsHstStm() {")
	r.Addf("sys.Logf(\"%%p %v(%%v)\", r, r.Prm())", r.Node.Full())
	r.Add("}")
	r.Add("aBse := r.A.Bse()")
	r.Add("if x.Vals.Cnt() == 0 || aBse.Vals.Cnt() == 0 || aBse.Vals.Cnt() <= r.Off {")
	r.Add("return r")
	r.Add("}")
	r.Add("if len(*x.Tmes) != len(*x.Vals) {")
	r.Add("err.Panicf(\"'x' length unequal (Tmes:%v Vals:%v)\", len(*x.Tmes), len(*x.Vals))")
	r.Add("}")
	r.Add("if len(*aBse.Tmes) != len(*aBse.Vals) {")
	r.Add("err.Panicf(\"'a' length unequal (Tmes:%v Vals:%v)\", len(*aBse.Tmes), len(*aBse.Vals))")
	r.Add("}")
	r.Add("xBnd, aBnd := AlignStmOtr(x.Slf, r.A, r.Off)")
	r.Add("if aBnd.Cnt() == 0 {")
	r.Add("return r")
	r.Add("}")
	r.Add("r.Tmes = aBse.Tmes.InBnd(aBnd).From(r.Off)")
	r.Add("r.Vals = flts.Make(aBnd.Cnt()-r.Off)")
	r.Addf("segBnds, acts := %v(r.Vals.Cnt())", _sys.Bsc.Bnd.arr.Segs.Ref(x))
	r.Add("for n, segBnd := range *segBnds {")
	r.Addf("seg := %v{}", seg.Adr(x))
	r.Add("seg.Bnd = segBnd")
	r.Add("seg.Vals = x.Vals.InBnd(xBnd)")
	r.Add("seg.Out = r.Vals")
	r.Add("seg.Off = r.Off")
	r.Add("seg.ValsA = aBse.Vals.InBnd(aBnd)")
	r.Add("acts[n] = seg")
	r.Add("}")
	r.Add("sys.Run().Pll(acts...)")
	r.Add("return r")
	return r
}

func (x *FleHstStm) CndScl(name string) (r *TypFn) {
	cnd, name := _sys.Ana.Hst.Cnd, strings.Title(name)
	r = x.ElmNodeTypFn(name, k.Scl, k.Scl, cnd, func(r *TypFn) {
		r.InPrm(_sys.Bsc.Flt, "scl").LitVal("1.1")
		r.Node.FldPrnt(x)
	})
	// seg
	seg := cnd.NodeSeg(r.Node.Name)
	seg.Fld("Vals", _sys.Bsc.Flt.arr)
	seg.Fld("Scl", _sys.Bsc.Flt)
	segAct := cnd.TypFn(k.Act, seg)
	segAct.Add("for n := x.Idx; n < x.Lim; n++ {")
	segAct.Addf("if x.Vals.At(n).%v(x.Scl) {", name)
	segAct.Add("*x.Out = append(*x.Out, (*x.Tmes)[n])")
	segAct.Add("}")
	segAct.Add("}")
	// node
	r.Add("if ana.Cfg.Trc.IsHstCnd() {")
	r.Addf("sys.Logf(\"%%p %v(%%v)\", r, r.Prm())", r.Node.Full())
	r.Add("}")
	r.Add("if x.Vals.Cnt() == 0 {")
	r.Add("return r")
	r.Add("}")
	r.Add("if len(*x.Tmes) != len(*x.Vals) {")
	r.Add("err.Panicf(\"length unequal (Tmes:%v Vals:%v)\", len(*x.Tmes), len(*x.Vals))")
	r.Add("}")
	r.Addf("r.Tmes = %v()", _sys.Bsc.Tme.arr.New.Ref(x))
	r.Addf("segBnds, acts := %v(x.Vals.Cnt())", _sys.Bsc.Bnd.arr.Segs.Ref(x))
	r.Add("for n, segBnd := range *segBnds {")
	r.Addf("seg := %v{}", seg.Adr(x))
	r.Add("seg.Bnd = segBnd")
	r.Add("seg.Tmes = x.Tmes")
	r.Add("seg.Vals = x.Vals")
	r.Add("seg.Scl = r.Scl")
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

func (x *FleHstStm) CndInr(name string) (r *TypFn) {
	cnd, name := _sys.Ana.Hst.Cnd, strings.Title(name)
	r = x.ElmNodeTypFn(name, k.Inr, k.Inr, cnd, func(r *TypFn) {
		r.InPrm(_sys.Bsc.Unt, "off").LitVal("1")
		r.Node.FldPrnt(x)
		if test := cnd.Test; test != nil {
			r.T2.MdlLst.Addf("x, e := stm.Bse(), %v{}", r.Node.Adr(test))
			r.T2.MdlLst.Addf(r.InPrms[0].AsnTest())
			r.T2.MdlLst.Addf("e.Tmes = %v()", _sys.Bsc.Tme.arr.New.Ref(test))
			r.T2.MdlLst.Addf("for n := %v; n < x.Tmes.Cnt()-e.Off; n++ {", _sys.Bsc.Unt.Zero.Ref(test))
			r.T2.MdlLst.Addf("if x.Vals.At(n+e.Off).%v(x.Vals.At(n)) {", name)
			r.T2.MdlLst.Add("*e.Tmes = append(*e.Tmes, x.Tmes.At(n+e.Off))")
			r.T2.MdlLst.Add("}")
			r.T2.MdlLst.Add("}")
			r.T2.MdlLst.Addf("tst.TmesEql(t, e.Tmes, a.Tmes, \"%v.Tmes\")", r.Node.Name)
		}
	})
	// seg
	seg := cnd.NodeSeg(r.Node.Name)
	seg.Fld("Vals", _sys.Bsc.Flt.arr)
	seg.Fld("Off", _sys.Bsc.Unt)
	segAct := cnd.TypFn(k.Act, seg)
	segAct.Add("for n := x.Idx; n < x.Lim; n++ {")
	segAct.Addf("if x.Vals.At(n+x.Off).%v((*x.Vals)[n]) {", name)
	segAct.Add("*x.Out = append(*x.Out, (*x.Tmes)[n+x.Off])")
	segAct.Add("}")
	segAct.Add("}")
	// node
	r.Add("if ana.Cfg.Trc.IsHstCnd() {")
	r.Addf("sys.Logf(\"%%p %v(%%v)\", r, r.Prm())", r.Node.Full())
	r.Add("}")
	r.Add("if len(*x.Vals) == 0 || x.Vals.Cnt() <= r.Off || r.Off == 0{")
	r.Add("return r")
	r.Add("}")
	r.Add("if len(*x.Tmes) != len(*x.Vals) {")
	r.Add("err.Panicf(\"length unequal (Tmes:%v Vals:%v)\", len(*x.Tmes), len(*x.Vals))")
	r.Add("}")
	r.Addf("r.Tmes = %v()", _sys.Bsc.Tme.arr.New.Ref(x))
	r.Addf("segBnds, acts := %v(x.Vals.Cnt()-r.Off)", _sys.Bsc.Bnd.arr.Segs.Ref(x))
	r.Add("for n, segBnd := range *segBnds {")
	r.Addf("seg := %v{}", seg.Adr(x))
	r.Add("seg.Bnd = segBnd")
	r.Add("seg.Tmes = x.Tmes")
	r.Add("seg.Vals = x.Vals")
	r.Add("seg.Off = r.Off")
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

func (x *FleHstStm) CndOtr(name string) (r *TypFn) {
	cnd, name := _sys.Ana.Hst.Cnd, strings.Title(name)
	r = x.ElmNodeTypFn(name, k.Otr, k.Otr, cnd, func(r *TypFn) {
		r.InPrm(_sys.Bsc.Unt, "off").LitVal("1")
		r.InPrm(x, "a").LitVal(TstStmLitVal)
		r.Node.FldPrnt(x)
	})
	// seg
	seg := cnd.NodeSeg(r.Node.Name)
	seg.Fld("Vals", _sys.Bsc.Flt.arr)
	seg.Fld("Off", _sys.Bsc.Unt)
	seg.Fld("ValsA", _sys.Bsc.Flt.arr)
	segAct := cnd.TypFn(k.Act, seg)
	segAct.Add("for n := x.Idx; n < x.Lim; n++ {")
	segAct.Addf("if x.Vals.At(n).%v((*x.ValsA)[n+x.Off]) {", name)
	segAct.Add("*x.Out = append(*x.Out, (*x.Tmes)[n+x.Off])")
	segAct.Add("}")
	segAct.Add("}")
	// node
	r.Add("if ana.Cfg.Trc.IsHstCnd() {")
	r.Addf("sys.Logf(\"%%p %v(%%v)\", r, r.Prm())", r.Node.Full())
	r.Add("}")
	r.Add("xBse, aBse := x, r.A.Bse()")
	r.Add("if xBse.Vals.Cnt() == 0 || aBse.Vals.Cnt() == 0 || aBse.Vals.Cnt() <= r.Off {")
	r.Add("return r")
	r.Add("}")
	r.Add("if len(*xBse.Tmes) != len(*xBse.Vals) {")
	r.Add("err.Panicf(\"'x' length unequal (Tmes:%v Vals:%v)\", len(*xBse.Tmes), len(*xBse.Vals))")
	r.Add("}")
	r.Add("if len(*aBse.Tmes) != len(*aBse.Vals) {")
	r.Add("err.Panicf(\"'a' length unequal (Tmes:%v Vals:%v)\", len(*aBse.Tmes), len(*aBse.Vals))")
	r.Add("}")
	r.Add("xBnd, aBnd := AlignStmOtr(x.Slf, r.A, r.Off)")
	r.Add("if aBnd.Cnt() == 0 {")
	r.Add("return r")
	r.Add("}")
	r.Addf("r.Tmes = %v()", _sys.Bsc.Tme.arr.New.Ref(x))
	r.Addf("segBnds, acts := %v(aBnd.Cnt()-r.Off)", _sys.Bsc.Bnd.arr.Segs.Ref(x))
	r.Add("for n, segBnd := range *segBnds {")
	r.Addf("seg := %v{}", seg.Adr(x))
	r.Add("seg.Bnd = segBnd")
	r.Add("seg.Vals = xBse.Vals.InBnd(xBnd)")
	r.Add("seg.Off = r.Off")
	r.Add("seg.ValsA = aBse.Vals.InBnd(aBnd)")
	r.Add("seg.Tmes = aBse.Tmes.InBnd(aBnd)")
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

// func (x *FleHstStm) Splt() (r *TypFn) { // stm splt
// 	r = x.TypFn(k.Splt, x.bse)
// 	r.Atr |= atr.PrtActFn
// 	r.InPrm(_sys.Bsc.Tme.arr, "btm")
// 	r.InPrm(_sys.Bsc.Tme.arr, "top")
// 	r.OutPrm(_sys.Ana.Hst.StmSplt, "r")
// 	r.Addf("r = %v{}", r.OutTyp().Adr(x))
// 	r.Add("r.Stm = x.Slf")
// 	r.Add("r.Btm = x.At(btm)")
// 	r.Add("r.Top = x.At(top)")
// 	r.Add("return r")
// 	x.MemSigFn(r) // add to interface
// 	return r
// }
