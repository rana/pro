package tpl

import (
	"strings"
	"sys/k"
	"sys/ks"
	"sys/tpl/atr"
)

type (
	FleHstInrvl struct {
		FleHstBse
	}
	FleHstInrvls struct {
		FleBse
		PrtArr
		PrtArrStrWrt // for Fbr.PrmWrt
	}
)

func (x *DirHst) NewInrvl() (r *FleHstInrvl) {
	r = &FleHstInrvl{}
	x.Inrvl = r
	r.Name = k.Inrvl
	r.Pkg = x.Pkg
	r.Ifc(r.Name, atr.TypAnaIfc)
	r.AddFle(r)
	return r
}
func (x *DirHst) NewInrvls() (r *FleHstInrvls) {
	r = &FleHstInrvls{}
	x.Inrvls = r
	r.FleBse = *NewArr(x.Inrvl, &r.PrtArr, x.Inrvl.Pkg)
	r.AddFle(r)
	return r
}

func (x *FleHstInrvl) InitTyp(bse *TypBse) {
	x.FleHstBse.InitTyp(bse)
	x.Typ().Bse().TestPth = append(_sys.Ana.Hst.Instr.Typ().Bse().TestPth, &TestStp{
		MdlFst: func(r *PkgFn) { r.Add("inrvl := tst.HstInstrInrvlI(instr, 10)") },
	})
	x.seg = x.NewSeg()
}

func (x *FleHstInrvl) InitFld(s *Struct) {
	x.FleHstBse.InitFld(s)
	x.bse.FldPrnt(_sys.Ana.Hst.Instr)
	x.bse.Fld("Tmes", _sys.Bsc.Tme.arr).Atr = atr.TstZeroSkp
	x.bse.Fld("TmeBnds", _sys.Bsc.Bnd.arr).Atr = atr.TstZeroSkp
	x.seg.Fld("TmeBnds", _sys.Bsc.Bnd.arr)
	x.seg.Fld("Out", _sys.Bsc.Bnd.arr)
}
func (x *FleHstInrvl) InitTypFn() {
	x.FleHstBse.InitTypFn()
	for _, side := range ks.Sides {
		x.Side(side)
	}
}

func (x *FleHstInrvl) Side(name string) (r *TypFn) {
	side, name := _sys.Ana.Hst.Side, strings.Title(name)
	x.Import("sys/err")
	x.Import(_sys)
	x.Import(_sys.Ana)
	x.Import(_sys.Bsc.Bnd)
	r = x.ElmNodeTypFn(name, "", "", side, func(r *TypFn) {})
	// seg
	seg := side.NodeSeg(r.Node.Name)
	segAct := x.TypFn(k.Act, seg)
	segAct.Add("for n := x.Idx; n < x.Lim; n++ {")
	segAct.Addf("(*x.Out)[n] = x.Stm.%vBndByTmeBnd((*x.TmeBnds)[n])", name)
	segAct.Addf("}")
	// node
	r.Add("if ana.Cfg.Trc.IsHstSide() {")
	r.Addf("sys.Logf(\"%%p %v(%%v)\", r, r.Prm())", r.Node.Full())
	r.Add("}")
	r.Add("if x.Tmes.Cnt() == 0 {")
	r.Add("return r")
	r.Add("}")
	r.Add("if len(*x.Tmes) != len(*x.TmeBnds) {")
	r.Add("err.Panicf(\"length unequal (Tmes:%v TmeBnds:%v)\", len(*x.Tmes), len(*x.TmeBnds))")
	r.Add("}")
	r.Addf("r.Vals = x.Instr.Bse().Ana.HstStm.%vs", name)
	r.Addf("r.ValBnds = %v(x.Tmes.Cnt())", _sys.Bsc.Bnd.arr.Make.Ref(x))
	r.Addf("segBnds, acts := %v(x.Tmes.Cnt())", _sys.Bsc.Bnd.arr.Segs.Ref(x))
	r.Add("for n, segBnd := range *segBnds {")
	r.Addf("seg := %v{}", seg.Adr(x))
	r.Add("seg.Idx = segBnd.Idx")
	r.Add("seg.Lim = segBnd.Lim")
	r.Add("seg.Stm = x.Instr.Bse().Ana.HstStm")
	r.Add("seg.TmeBnds = x.TmeBnds")
	r.Add("seg.Out = r.ValBnds")
	r.Add("acts[n] = seg")
	r.Add("}")
	r.Add("sys.Run().Pll(acts...)")
	r.Add("return r")
	return r
}
