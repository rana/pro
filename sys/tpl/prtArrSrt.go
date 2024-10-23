package tpl

import (
	"strings"
	"sys/k"
	"sys/tpl/atr"
)

type (
	PrtArrSrt struct {
		PrtBse
	}
)

func (x *PrtArrSrt) InitPrtTypFn() {
	arr := x.t.PrtArr().Arr
	elmRel := arr.Elm.Fle.Bse().GetPrt((*PrtRel)(nil)).(*PrtRel)

	// // tst
	// if x.f.Tst != nil {
	// 	x.f.Tst.ArrRelPrt(elmRel.Lss, elmRel.Gtr, "")
	// }
	// tst
	x.t.TstAsc = _sys.Tst.Srt(x.f.Typ(), elmRel.Lss)
	x.t.TstDsc = _sys.Tst.Srt(x.f.Typ(), elmRel.Gtr)
	x.SrtAsc(elmRel.Lss, elmRel.Eql, "")
	x.SrtDsc(elmRel.Gtr, elmRel.Eql, "")
	x.SrtQuick(elmRel.Cmp)
	x.SrtIns(elmRel.Cmp, elmRel.Lss, elmRel.Gtr)
	x.SrtMdnOf3(elmRel.Cmp, elmRel.Lss)
	x.Swp()
}
func (x *PrtArrSrt) SrtAsc(lss, eql *PkgFn, suffix string) (r *TypFn) {
	r = x.f.TypFnf("%v%v", k.SrtAsc, strings.Title(suffix))
	r.OutPrm(x.f.Typ())
	r.Add("if x.Cnt() > 1 {")
	r.Addf("x.SrtQuick(0, x.LstIdx(), %v, %v)", lss.Ref(x.f), eql.Ref(x.f))
	r.Add("}")
	r.Add("return x")
	// test
	if x.f.Test != nil && Opt.IsArr() {
		r.T.Manual = true
		r.T.Addf("%v(t, x.%v())", x.t.TstAsc.Ref(x.f.Test), r.Name)
	}
	return r
}
func (x *PrtArrSrt) SrtDsc(gtr, eql *PkgFn, suffix string) (r *TypFn) {
	r = x.f.TypFnf("%v%v", k.SrtDsc, strings.Title(suffix))
	r.OutPrm(x.f.Typ())
	r.Add("if x.Cnt() > 1 {")
	r.Addf("x.SrtQuick(0, x.LstIdx(), %v, %v)", gtr.Ref(x.f), eql.Ref(x.f))
	r.Add("}")
	r.Add("return x")
	// test
	if x.f.Test != nil && Opt.IsArr() {
		r.T.SkpTst = true
		r.T.SkpExcpy = true
		r.T.Addf("%v(t, actual)", x.t.TstDsc.Ref(x.f.Test))
		r.T.Addf("%v(t, axCpy)", x.t.TstDsc.Ref(x.f.Test))
	}
	return r
}
func (x *PrtArrSrt) SrtQuick(cmp *Func) (r *TypFn) {
	r = x.f.TypFna(k.SrtQuick, atr.None) // Quick sort
	r.InPrm(_sys.Bsc.Unt, "lo")
	r.InPrm(_sys.Bsc.Unt, "hi")
	r.InPrm(cmp, "cmp")
	r.InPrm(cmp, "eql")
	r.OutPrm(x.f.Typ())
	r.Add("n := hi - lo + 1")
	r.Add("if (n <= 8) { // cutoff to insertion sort")
	r.Add("return x.SrtIns(lo, hi, cmp)")
	r.Add("}")
	r.Add("if (n <= 40) { // use median-of-3 as partitioning element")
	r.Add("mdn := x.SrtMdnOf3(lo, lo + n/2, hi, cmp)")
	r.Add("x.Swp(mdn, lo)")
	r.Add("} else { // use Tukey ninther as partitioning element")
	r.Add("eps := n/8")
	r.Add("mid := lo + n/2")
	r.Add("m1 := x.SrtMdnOf3(lo, lo + eps, lo + eps + eps, cmp)")
	r.Add("m2 := x.SrtMdnOf3(mid - eps, mid, mid + eps, cmp)")
	r.Add("m3 := x.SrtMdnOf3(hi - eps - eps, hi - eps, hi, cmp)")
	r.Add("ninther := x.SrtMdnOf3(m1, m2, m3, cmp)")
	r.Add("x.Swp(ninther, lo)")
	r.Add("}")
	r.Add("i, j := lo, hi+1 // Bentley-McIlroy 3-way partitioning")
	r.Add("p, q := lo, hi+1")
	r.Add("v := (*x)[lo]")
	r.Add("for {")
	r.Add("i++")
	r.Add("for cmp((*x)[i], v) {")
	r.Add("if i == hi {")
	r.Add("break")
	r.Add("}")
	r.Add("i++")
	r.Add("}")
	r.Add("if j != 0 {")
	r.Add("j--")
	r.Add("}")
	r.Add("for cmp(v, (*x)[j]) {")
	r.Add("if j == lo {")
	r.Add("break")
	r.Add("}")
	r.Add("if j != 0 {")
	r.Add("j--")
	r.Add("}")
	r.Add("}")
	// r.Add("if i == j && (*x)[i].Eql(v) { // pointers cross")
	r.Add("if i == j && eql((*x)[i], v) { // pointers cross")
	r.Add("p++")
	r.Add("x.Swp(p, i)")
	r.Add("}")
	r.Add("if i >= j {")
	r.Add("break")
	r.Add("}")
	r.Add("x.Swp(i, j)")
	// r.Add("if (*x)[i].Eql(v) {")
	r.Add("if eql((*x)[i], v) {")
	r.Add("p++")
	r.Add("x.Swp(p, i)")
	r.Add("}")
	// r.Add("if (*x)[j].Eql(v) {")
	r.Add("if eql((*x)[j], v) {")
	r.Add("q--")
	r.Add("x.Swp(q, j)")
	r.Add("}")
	r.Add("}")
	r.Add("i = j + 1")
	r.Add("for k := lo; k <= p; k++ {")
	r.Add("x.Swp(k, j)")
	r.Add("if j != 0 {")
	r.Add("j--")
	r.Add("}")
	r.Add("}")
	r.Add("for k := hi; k >= q; k-- {")
	r.Add("x.Swp(k, i)")
	r.Add("i++")
	r.Add("}")
	r.Add("x.SrtQuick(lo, j, cmp, eql)")
	r.Add("x.SrtQuick(i, hi, cmp, eql)")
	r.Add("return x")
	return r
}
func (x *PrtArrSrt) SrtIns(cmp *Func, lss, gtr *PkgFn) (r *TypFn) {
	r = x.f.TypFna(k.SrtIns, atr.None) // Insertion sort
	r.InPrm(_sys.Bsc.Unt, "lo")
	r.InPrm(_sys.Bsc.Unt, "hi")
	r.InPrm(cmp, "cmp")
	r.OutPrm(x.f.Typ())
	r.Add("for i := lo; i <= hi; i++ {")
	r.Add("for j := i; j > lo && cmp((*x)[j], (*x)[j-1]); j-- {")
	r.Add("x.Swp(j, j-1)")
	r.Add("}")
	r.Add("}")
	r.Add("return x")
	// test
	if x.f.Test != nil && Opt.IsArr() && lss != nil && gtr != nil {
		r.T.RxrOnly = true
		r.T.Manual = true
		r.T.TstCnd = "x.Cnt() > 2"
		r.T.Add("axCpyAsc := x.Cpy()")
		r.T.Addf("a := axCpyAsc.%v(axCpyAsc.FstIdx(), axCpyAsc.LstIdx(), %v)", r.Name, lss.Ref(x.f.Test))
		r.T.Addf("%v(t, a)", x.t.TstAsc.Ref(x.f.Test))
		r.T.Addf("%v(t, axCpyAsc)", x.t.TstAsc.Ref(x.f.Test))
		r.T.Add("axCpyDsc := x.Cpy()")
		r.T.Addf("a = axCpyDsc.%v(axCpyDsc.FstIdx(), axCpyDsc.LstIdx(), %v)", r.Name, gtr.Ref(x.f.Test))
		r.T.Addf("%v(t, a)", x.t.TstDsc.Ref(x.f.Test))
		r.T.Addf("%v(t, axCpyDsc)", x.t.TstDsc.Ref(x.f.Test))
	}
	return r
}
func (x *PrtArrSrt) SrtMdnOf3(cmp *Func, lss *PkgFn) (r *TypFn) {
	r = x.f.TypFna(k.SrtMdnOf3, atr.None)
	r.InPrm(_sys.Bsc.Unt, "i")
	r.InPrm(_sys.Bsc.Unt, "j")
	r.InPrm(_sys.Bsc.Unt, "k")
	r.InPrm(cmp, "cmp")
	r.OutPrm(_sys.Bsc.Unt)
	r.Add("if cmp((*x)[i], (*x)[j]) {")
	r.Add("if cmp((*x)[j], (*x)[k]) {")
	r.Add("return j")
	r.Add("}")
	r.Add("if cmp((*x)[i], (*x)[k]) {")
	r.Add("return k")
	r.Add("}")
	r.Add("return i")
	r.Add("}")
	r.Add("if cmp((*x)[k], (*x)[j]) {")
	r.Add("return j")
	r.Add("}")
	r.Add("if cmp((*x)[k], (*x)[i]) {")
	r.Add("return k")
	r.Add("}")
	r.Add("return i")
	// test
	if x.f.Test != nil && Opt.IsArr() && lss != nil {
		r.T.RxrOnly = true
		r.T.Manual = true
		r.T.TstCnd = "x.Cnt() > 2"
		r.T.Add("cpy := x.Cpy().Rand()")
		r.T.Addf("a := cpy.%v(cpy.FstIdx(), cpy.MdlIdx(), cpy.LstIdx(), %v)", r.Name, lss.Ref(x.f.Test))
		r.T.Add("switch {")
		r.T.Add("case a == cpy.FstIdx():")
		r.T.Add("tst.True(t, a < cpy.MdlIdx() || a > cpy.MdlIdx())")
		r.T.Add("tst.True(t, a < cpy.LstIdx() || a > cpy.LstIdx())")
		r.T.Add("case a == cpy.MdlIdx():")
		r.T.Add("tst.True(t, a < cpy.FstIdx() || a > cpy.FstIdx())")
		r.T.Add("tst.True(t, a < cpy.LstIdx() || a > cpy.LstIdx())")
		r.T.Add("case a == cpy.LstIdx():")
		r.T.Add("tst.True(t, a < cpy.FstIdx() || a > cpy.FstIdx())")
		r.T.Add("tst.True(t, a < cpy.MdlIdx() || a > cpy.MdlIdx())")
		r.T.Add("}")
	}
	return r
}
func (x *PrtArrSrt) Swp() (r *TypFn) {
	r = x.f.TypFna(k.Swp, atr.None)
	r.InPrm(_sys.Bsc.Unt, "i")
	r.InPrm(_sys.Bsc.Unt, "j")
	r.Add("(*x)[i], (*x)[j] = (*x)[j], (*x)[i]")
	// test
	if x.f.Test != nil && Opt.IsArr() {
		r.T.Manual = true
		r.T.RxrOnly = true
		r.T.TstCnd = "x.Cnt() > 1"
		r.T.Add("expected := x.Cpy()")
		r.T.Add("axCpy := x.Cpy()")
		r.T.Add("i, j := x.FstIdx(), x.LstIdx()")
		r.T.Add("(*expected)[i], (*expected)[j] = (*expected)[j], (*expected)[i]")
		r.T.Add("axCpy.Swp(i, j)")
		r.T.Addf("%v(t, expected, axCpy)", x.t.TstEql.Ref(x.f.Test))
	}
	return r
}
