package tpl

import (
	"strings"
	"sys/k"
	"sys/ks"
	"sys/tpl/atr"
)

type (
	FleHstSide struct {
		FleHstBse
	}
	FleHstSides struct {
		FleBse
		PrtArr
		PrtArrStrWrt // for Fbr.PrmWrt
		PrtString
	}
)

func (x *DirHst) NewSide() (r *FleHstSide) {
	r = &FleHstSide{}
	x.Side = r
	r.Name = k.Side
	r.Pkg = x.Pkg
	r.Ifc(r.Name, atr.TypAnaIfc)
	r.AddFle(r)
	return r
}
func (x *DirHst) NewSides() (r *FleHstSides) {
	r = &FleHstSides{}
	x.Sides = r
	r.FleBse = *NewArr(x.Side, &r.PrtArr, x.Side.Pkg)
	r.AddFle(r)
	return r
}

func (x *FleHstSide) InitTyp(bse *TypBse) {
	x.FleHstBse.InitTyp(bse)
	x.Typ().Bse().TestPth = append(_sys.Ana.Hst.Inrvl.Typ().Bse().TestPth, &TestStp{
		MdlFst: func(r *PkgFn) { r.Add("side := tst.HstInrvlSideBid(inrvl)") },
	})
	x.seg = x.NewSeg()
}

func (x *FleHstSide) InitFld(s *Struct) {
	x.FleHstBse.InitFld(s)
	x.bse.FldPrnt(_sys.Ana.Hst.Inrvl)
	x.bse.Fld("Vals", _sys.Bsc.Flt.arr).Atr = atr.Get
	x.bse.Fld("ValBnds", _sys.Bsc.Bnd.arr)
	x.seg.Fld("Stm", _sys.Ana.Stm)
	x.seg.Fld("TmeBnds", _sys.Bsc.Bnd.arr)
	x.seg.Fld("Out", _sys.Bsc.Bnd.arr)
}
func (x *FleHstSide) InitTypFn() {
	x.FleHstBse.InitTypFn()
	x.Import("sys/err")
	x.Import(_sys)             // for log
	x.Import(_sys.Ana)         // for trc cfg
	x.Import(_sys.Bsc.Bnd.arr) // for bnds.Segs
	for _, stmRte := range ks.StmRtes {
		x.StmRte(stmRte)
	}
	x.Sar()
	x.Ema()
}
func (x *FleHstSide) StmRte(name string) (r *TypFn) {
	stm, name := _sys.Ana.Hst.Stm, strings.Title(name)
	r = x.ElmNodeTypFn(name, k.Rte, "", stm, func(r *TypFn) {
		r.Node.FldPrnt(x)
	})
	// seg
	seg := stm.NodeSeg(r.Node.Name)
	seg.Fld("ValBnds", _sys.Bsc.Bnd.arr)
	segAct := stm.TypFn(k.Act, seg)
	segAct.Add("for n := x.Idx; n < x.Lim; n++ {")
	segAct.Add("elm := (*x.Vals)[(*x.ValBnds)[n].Idx:(*x.ValBnds)[n].Lim]")
	segAct.Addf("(*x.Out)[n] = elm.%v()", name)
	segAct.Add("}")
	// node
	r.Add("if ana.Cfg.Trc.IsHstStm() {")
	r.Addf("sys.Logf(\"%%p %v(%%v)\", r, r.Prm())", r.Node.Full())
	r.Add("}")
	r.Add("if x.ValBnds.Cnt() == 0 {")
	r.Add("return r")
	r.Add("}")
	r.Add("inrvl := x.Inrvl.Bse()")
	r.Add("if inrvl.TmeBnds.Cnt() != x.ValBnds.Cnt() {")
	r.Add("err.Panicf(\"length unequal (TmeBnds:%v ValBnds:%v)\", inrvl.TmeBnds.Cnt(), x.ValBnds.Cnt())")
	r.Add("}")
	r.Add("r.Tmes = inrvl.Tmes")
	r.Add("r.Vals = flts.Make(inrvl.TmeBnds.Cnt())")
	r.Addf("segBnds, acts := %v(inrvl.TmeBnds.Cnt())", _sys.Bsc.Bnd.arr.Segs.Ref(x))
	r.Add("for n, segBnd := range *segBnds {")
	r.Addf("seg := %v{}", seg.Adr(x))
	r.Add("seg.Bnd = segBnd")
	r.Add("seg.Vals = x.Vals")
	r.Add("seg.ValBnds = x.ValBnds")
	r.Add("seg.Out = r.Vals")
	r.Add("acts[n] = seg")
	r.Add("}")
	r.Add("sys.Run().Pll(acts...)")
	r.Add("return r")
	return r
}

func (x *FleHstSide) Sar() (r *TypFn) {
	stm, name := _sys.Ana.Hst.Stm, strings.Title(k.Sar)
	r = x.ElmNodeTypFn(name, "Rte1", "", stm, func(r *TypFn) {
		r.Node.FldPrnt(x)
		r.InPrm(_sys.Bsc.Flt, "afInc").LitVal("0.02")
		r.InPrm(_sys.Bsc.Flt, "afMax").LitVal("0.2")
	})
	// node
	r.Add("if ana.Cfg.Trc.IsHstStm() {")
	r.Addf("sys.Logf(\"%%p %v(%%v)\", r, r.Prm())", r.Node.Full())
	r.Add("}")
	r.Add("if x.ValBnds.Cnt() == 0 {")
	r.Add("return r")
	r.Add("}")
	r.Add("if x.Inrvl.Bse().TmeBnds.Cnt() != x.ValBnds.Cnt() {")
	r.Add("err.Panicf(\"length unequal (TmeBnds:%v ValBnds:%v)\", x.Inrvl.Bse().TmeBnds.Cnt(), x.ValBnds.Cnt())")
	r.Add("}")
	r.Addf("r.Calc()")
	r.Add("return r")

	return r
}

func (x *FleHstSide) Ema() (r *TypFn) {
	stm, name := _sys.Ana.Hst.Stm, strings.Title(k.Ema)
	x.Import(_sys.Bsc.Unt)
	r = x.ElmNodeTypFn(name, k.Rte, "", stm, func(r *TypFn) {
		r.Node.FldPrnt(x)
	})
	// node
	r.Add("if ana.Cfg.Trc.IsHstStm() {")
	r.Addf("sys.Logf(\"%%p %v(%%v)\", r, r.Prm())", r.Node.Full())
	r.Add("}")
	r.Add("if x.ValBnds.Cnt() == 0 {")
	r.Add("return r")
	r.Add("}")
	r.Add("if x.Inrvl.Bse().TmeBnds.Cnt() != x.ValBnds.Cnt() {")
	r.Add("err.Panicf(\"length unequal (TmeBnds:%v ValBnds:%v)\", x.Inrvl.Bse().TmeBnds.Cnt(), x.ValBnds.Cnt())")
	r.Add("}")
	r.Add("r.Tmes = x.Inrvl.Bse().Tmes")
	r.Add("r.Vals = flts.Make(x.Inrvl.Bse().TmeBnds.Cnt())")
	r.Add("// NON-PLL IMPL DUE TO PRV VAL CHAINING")
	r.Add("//    EMA CALC   https://stockcharts.com/school/doku.php?id=chart_school:technical_indicators:moving_averages")
	r.Add("// Initial SMA: 10-period sum / 10")
	r.Add("// Multiplier: (2 / (Time periods + 1) ) = (2 / (10 + 1) ) = 0.1818 (18.18%)")
	r.Add("// EMA: {Close - EMA(previous day)} x multiplier + EMA(previous day)")
	r.Add("// NOTE: Each ValBnd may have different lengths; multiplier will vary from tick to tick for Side")
	r.Add("(*r.Vals)[0] = x.Vals.InBnd(x.ValBnds.At(0)).Sma()")
	r.Add("for n := 1; n < len(*x.ValBnds); n++ {")
	r.Addf("(*r.Vals)[n] = (*r.Vals)[n-1] + ((*x.Vals)[(*x.ValBnds)[n].Lim-1] - (*r.Vals)[n-1]) * (flt.Flt(2) / flt.Flt(x.ValBnds.At(unt.Unt(n)).Len() + 1))")
	r.Add("}")
	r.Add("return r")
	return r
}
