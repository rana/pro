package tpl

import "sys/k"

type (
	PrtArrRel struct {
		PrtBse
		Arr       *Arr
		ElmRel    *PrtRel
		SrchIdx   *TypFn
		Has       *TypFn
		SrtAsc    *TypFn
		SrtDsc    *TypFn
		SrtQuick  *TypFn
		SrtIns    *TypFn
		SrtMdnOf3 *TypFn
		Swp       *TypFn
	}
)

func (x *PrtArrRel) InitPrtTyp() {
	x.Arr = x.f.GetPrt((*PrtArr)(nil)).(*PrtArr).Arr
	x.ElmRel = x.Arr.Elm.Fle.Bse().GetPrt((*PrtRel)(nil)).(*PrtRel)
}
func (x *PrtArrRel) InitPrtTypFn() {
	// TODO: REVISE WITH PrtArrRng
	// // tst
	// if x.f.Tst != nil {
	// 	x.f.Tst.ArrRelPrt()
	// }
	x.srchIdxEql()
	x.SrchIdx = x.srchIdx()
	x.Has = x.has()
	// x.SrtAsc = x.srtAsc()
	// x.SrtDsc = x.srtDsc()
	// x.SrtQuick = x.srtQuick()
	// x.SrtIns = x.srtIns()
	// x.SrtMdnOf3 = x.srtMdnOf3()
	// x.Swp = x.swp()
}

func (x *PrtArrRel) srchIdxEql() (r *TypFn) {
	r = x.f.TypFn("SrchIdxEql")
	r.InPrm(x.Arr.Elm, "v")
	r.OutPrm(_sys.Bsc.Unt)
	r.Add("i, j := unt.Zero, unt.Unt(len(*x))")
	r.Add("for i < j {")
	r.Add("if (*x)[(i+j)>>1] < v {")
	r.Add("i = (i+j)>>1 + 1")
	r.Add("} else {")
	r.Add("j = (i+j)>>1")
	r.Add("}")
	r.Add("}")
	r.Add("return i")
	// test
	if x.f.Test != nil && Opt.IsArr() {
		// x.f.Test.Import("sort")
		// x.f.Test.ImportAlias(_sys.Bsc.Int)
		r.T.RxrOnly = true
		r.T.Manual = true
		r.T.Add("// TODO: MORE THOROUGH TESTING?")
		r.T.Addf("axCpy := x.Cpy().SrtAsc() // %v expects srt order", r.Name)
		r.T.Add("for _, v := range *axCpy { // positive cases")
		// r.T.Add("expected := unt.Unt(sort.Search(len(*axCpy), func(i int) bool { return (*axCpy)[i] >= v }))")
		r.T.Add("expected := axCpy.SrchIdx(v)")
		r.T.Addf("actual := axCpy.%v(v)", r.Name)
		r.T.Addf("%v(t, expected, actual, \"Idx search for \", v)", _sys.Bsc.Unt.Typ().Bse().TstEql.Ref(x.f.Test))
		r.T.Add("}")
	}
	return r
}

func (x *PrtArrRel) srchIdx() (r *TypFn) {
	x.f.Import("sort")
	r = x.f.TypFn("SrchIdx")
	r.InPrm(x.Arr.Elm, "v")
	r.InPrmVariadic(_sys.Bsc.Bol, "near")
	r.OutPrm(_sys.Bsc.Unt)
	r.Add("if len(*x) == 0 {")
	r.Add("return unt.Max")
	r.Add("}")
	r.Add("if len(near) > 0 {")
	r.Add("if v <= (*x)[0] { // lwr bnd")
	r.Add("return 0")
	r.Add("} else if v >= (*x)[len(*x)-1] { // upr bnd")
	r.Add("return unt.Unt(len(*x) - 1)")
	r.Add("}")
	r.Add("}")
	r.Add("idx := sort.Search(len(*x), func(i int) bool { return (*x)[i] >= v })")
	r.Add("if idx < len(*x) && (len(near) > 0 || (*x)[idx] == v) { // near does not require exact match; default requires exact match")
	r.Add("return unt.Unt(idx)")
	r.Add("}")
	r.Add("return unt.Max")
	// test
	if x.f.Test != nil && Opt.IsArr() {
		r.T.RxrOnly = true
		r.T.Manual = true
		r.T.Add("// TODO: MORE THOROUGH TESTING")
		r.T.Addf("axCpy := x.Cpy().SrtAsc() // %v expects srt order", r.Name)
		r.T.Add("for _, v := range *axCpy { // positive cases")
		r.T.Addf("if idx := axCpy.%v(v); idx == unt.Max {", r.Name)
		r.T.Add("t.Fatalf(\"expected to find value (v:%v)\", v)")
		r.T.Add("}")
		r.T.Add("}")
	}
	return r
}
func (x *PrtArrRel) has() (r *TypFn) {
	r = x.f.TypFn(k.Has)
	r.InPrm(x.Arr.Elm, "v")
	r.OutPrm(_sys.Bsc.Bol)
	r.Add("i, j := unt.Zero, unt.Unt(len(*x))")
	r.Add("for i < j {")
	r.Add("if (*x)[(i+j)>>1] < v {")
	r.Add("i = (i+j)>>1 + 1")
	r.Add("} else {")
	r.Add("j = (i+j)>>1")
	r.Add("}")
	r.Add("}")
	r.Add("return i != unt.Unt(len(*x)) && (*x)[i] == v")
	// test
	if x.f.Test != nil && Opt.IsArr() {
		r.T.RxrOnly = true
		r.T.Manual = true
		r.T.Addf("axCpy := x.Cpy().SrtAsc() // %v expects srt order", r.Name)
		r.T.Add("for _, v := range *axCpy { // positive cases")
		r.T.Addf("if !axCpy.%v(v) {", r.Name)
		r.T.Add("t.Fatalf(\"expected to find value (v:%v)\", v)")
		r.T.Add("}")
		r.T.Add("}")
	}
	return r
}

// func (x *PrtArrRel) srtAsc() (r *TypFn) {
// 	r = x.f.TypFn("SrtAsc")
// 	r.OutPrm(x.f.Typ())
// 	r.Add("if x.Cnt() > 1 {")
// 	r.Addf("x.SrtQuick(0, x.LstIdx(), %v)", x.ElmRel.Lss.Ref(x.f))
// 	r.Add("}")
// 	r.Add("return x")
// 	// test
// 	if x.f.Test != nil && Opt.IsArr() {
// 		r.T.SkpTst = true
// 		r.T.SkpExcpy = true
// 		r.T.Addf("%v(t, actual)", x.f.Tst.Asc.Ref(x.f.Test))
// 		r.T.Addf("%v(t, axCpy)", x.f.Tst.Asc.Ref(x.f.Test))
// 	}
// 	return r
// }
// func (x *PrtArrRel) srtDsc() (r *TypFn) {
// 	r = x.f.TypFn("SrtDsc")
// 	r.OutPrm(x.f.Typ())
// 	r.Add("if x.Cnt() > 1 {")
// 	r.Addf("x.SrtQuick(0, x.LstIdx(), %v)", x.ElmRel.Gtr.Ref(x.f))
// 	r.Add("}")
// 	r.Add("return x")
// 	// test
// 	if x.f.Test != nil && Opt.IsArr() {
// 		r.T.SkpTst = true
// 		r.T.SkpExcpy = true
// 		r.T.Addf("%v(t, actual)", x.f.Tst.Dsc.Ref(x.f.Test))
// 		r.T.Addf("%v(t, axCpy)", x.f.Tst.Dsc.Ref(x.f.Test))
// 	}
// 	return r
// }
// func (x *PrtArrRel) srtQuick() (r *TypFn) {
// 	r = x.f.TypFn("SrtQuick") // Quick sort
// 	r.SkpXpr = true
// 	r.InPrm(_sys.Bsc.Unt, "lo")
// 	r.InPrm(_sys.Bsc.Unt, "hi")
// 	r.InPrm(x.ElmRel.Cmp, "cmp")
// 	r.OutPrm(x.f.Typ())
// 	r.Add("n := hi - lo + 1")
// 	r.Add("if (n <= 8) { // cutoff to insertion sort")
// 	r.Add("return x.SrtIns(lo, hi, cmp)")
// 	r.Add("}")
// 	r.Add("if (n <= 40) { // use median-of-3 as partitioning element")
// 	r.Add("mdn := x.SrtMdnOf3(lo, lo + n/2, hi, cmp)")
// 	r.Add("x.Swp(mdn, lo)")
// 	r.Add("} else { // use Tukey ninther as partitioning element")
// 	r.Add("eps := n/8")
// 	r.Add("mid := lo + n/2")
// 	r.Add("m1 := x.SrtMdnOf3(lo, lo + eps, lo + eps + eps, cmp)")
// 	r.Add("m2 := x.SrtMdnOf3(mid - eps, mid, mid + eps, cmp)")
// 	r.Add("m3 := x.SrtMdnOf3(hi - eps - eps, hi - eps, hi, cmp)")
// 	r.Add("ninther := x.SrtMdnOf3(m1, m2, m3, cmp)")
// 	r.Add("x.Swp(ninther, lo)")
// 	r.Add("}")
// 	r.Add("i, j := lo, hi+1 // Bentley-McIlroy 3-way partitioning")
// 	r.Add("p, q := lo, hi+1")
// 	r.Add("v := (*x)[lo]")
// 	r.Add("for {")
// 	r.Add("i++")
// 	r.Add("for cmp((*x)[i], v) {")
// 	r.Add("if i == hi {")
// 	r.Add("break")
// 	r.Add("}")
// 	r.Add("i++")
// 	r.Add("}")
// 	r.Add("if j != 0 {")
// 	r.Add("j--")
// 	r.Add("}")
// 	r.Add("for cmp(v, (*x)[j]) {")
// 	r.Add("if j == lo {")
// 	r.Add("break")
// 	r.Add("}")
// 	r.Add("if j != 0 {")
// 	r.Add("j--")
// 	r.Add("}")
// 	r.Add("}")
// 	r.Add("if i == j && (*x)[i].Eql(v) { // pointers cross")
// 	r.Add("p++")
// 	r.Add("x.Swp(p, i)")
// 	r.Add("}")
// 	r.Add("if i >= j {")
// 	r.Add("break")
// 	r.Add("}")
// 	r.Add("x.Swp(i, j)")
// 	r.Add("if (*x)[i].Eql(v) {")
// 	r.Add("p++")
// 	r.Add("x.Swp(p, i)")
// 	r.Add("}")
// 	r.Add("if (*x)[j].Eql(v) {")
// 	r.Add("q--")
// 	r.Add("x.Swp(q, j)")
// 	r.Add("}")
// 	r.Add("}")
// 	r.Add("i = j + 1")
// 	r.Add("for k := lo; k <= p; k++ {")
// 	r.Add("x.Swp(k, j)")
// 	r.Add("if j != 0 {")
// 	r.Add("j--")
// 	r.Add("}")
// 	r.Add("}")
// 	r.Add("for k := hi; k >= q; k-- {")
// 	r.Add("x.Swp(k, i)")
// 	r.Add("i++")
// 	r.Add("}")
// 	r.Add("x.SrtQuick(lo, j, cmp)")
// 	r.Add("x.SrtQuick(i, hi, cmp)")
// 	r.Add("return x")
// 	return r
// }
// func (x *PrtArrRel) srtIns() (r *TypFn) {
// 	r = x.f.TypFn("SrtIns") // Insertion sort
// 	r.SkpXpr = true
// 	r.InPrm(_sys.Bsc.Unt, "lo")
// 	r.InPrm(_sys.Bsc.Unt, "hi")
// 	r.InPrm(x.ElmRel.Cmp, "cmp")
// 	r.OutPrm(x.f.Typ())
// 	r.Add("for i := lo; i <= hi; i++ {")
// 	r.Add("for j := i; j > lo && cmp((*x)[j], (*x)[j-1]); j-- {")
// 	r.Add("x.Swp(j, j-1)")
// 	r.Add("}")
// 	r.Add("}")
// 	r.Add("return x")
// 	// test
// 	if x.f.Test != nil && Opt.IsArr() {
// 		r.T.RxrOnly = true
// 		r.T.Manual = true
// 		r.T.TstCnd = "x.Cnt() > 2"
// 		r.T.Add("axCpyAsc := x.Cpy()")
// 		r.T.Addf("a := axCpyAsc.%v(axCpyAsc.FstIdx(), axCpyAsc.LstIdx(), %v)", r.Name, x.ElmRel.Lss.Ref(x.f.Test))
// 		r.T.Addf("%v(t, a)", x.f.Tst.Asc.Ref(x.f.Test))
// 		r.T.Addf("%v(t, axCpyAsc)", x.f.Tst.Asc.Ref(x.f.Test))
// 		r.T.Add("axCpyDsc := x.Cpy()")
// 		r.T.Addf("a = axCpyDsc.%v(axCpyDsc.FstIdx(), axCpyDsc.LstIdx(), %v)", r.Name, x.ElmRel.Gtr.Ref(x.f.Test))
// 		r.T.Addf("%v(t, a)", x.f.Tst.Dsc.Ref(x.f.Test))
// 		r.T.Addf("%v(t, axCpyDsc)", x.f.Tst.Dsc.Ref(x.f.Test))
// 	}
// 	return r
// }
// func (x *PrtArrRel) srtMdnOf3() (r *TypFn) {
// 	r = x.f.TypFn("SrtMdnOf3")
// 	r.SkpXpr = true
// 	r.InPrm(_sys.Bsc.Unt, "i")
// 	r.InPrm(_sys.Bsc.Unt, "j")
// 	r.InPrm(_sys.Bsc.Unt, "k")
// 	r.InPrm(x.ElmRel.Cmp, "cmp")
// 	r.OutPrm(_sys.Bsc.Unt)
// 	r.Add("if cmp((*x)[i], (*x)[j]) {")
// 	r.Add("if cmp((*x)[j], (*x)[k]) {")
// 	r.Add("return j")
// 	r.Add("}")
// 	r.Add("if cmp((*x)[i], (*x)[k]) {")
// 	r.Add("return k")
// 	r.Add("}")
// 	r.Add("return i")
// 	r.Add("}")
// 	r.Add("if cmp((*x)[k], (*x)[j]) {")
// 	r.Add("return j")
// 	r.Add("}")
// 	r.Add("if cmp((*x)[k], (*x)[i]) {")
// 	r.Add("return k")
// 	r.Add("}")
// 	r.Add("return i")
// 	// test
// 	if x.f.Test != nil && Opt.IsArr() {
// 		r.T.RxrOnly = true
// 		r.T.Manual = true
// 		r.T.TstCnd = "x.Cnt() > 2"
// 		r.T.Add("cpy := x.Cpy().Rand()")
// 		r.T.Addf("a := cpy.%v(cpy.FstIdx(), cpy.MdlIdx(), cpy.LstIdx(), %v)", r.Name, x.ElmRel.Lss.Ref(x.f.Test))
// 		r.T.Add("switch {")
// 		r.T.Add("case a == cpy.FstIdx():")
// 		r.T.Add("tst.True(t, a < cpy.MdlIdx() || a > cpy.MdlIdx())")
// 		r.T.Add("tst.True(t, a < cpy.LstIdx() || a > cpy.LstIdx())")
// 		r.T.Add("case a == cpy.MdlIdx():")
// 		r.T.Add("tst.True(t, a < cpy.FstIdx() || a > cpy.FstIdx())")
// 		r.T.Add("tst.True(t, a < cpy.LstIdx() || a > cpy.LstIdx())")
// 		r.T.Add("case a == cpy.LstIdx():")
// 		r.T.Add("tst.True(t, a < cpy.FstIdx() || a > cpy.FstIdx())")
// 		r.T.Add("tst.True(t, a < cpy.MdlIdx() || a > cpy.MdlIdx())")
// 		r.T.Add("}")
// 	}
// 	return r
// }
// func (x *PrtArrRel) swp() (r *TypFn) {
// 	r = x.f.TypFn("Swp")
// 	r.SkpXpr = true
// 	r.InPrm(_sys.Bsc.Unt, "i")
// 	r.InPrm(_sys.Bsc.Unt, "j")
// 	r.Add("(*x)[i], (*x)[j] = (*x)[j], (*x)[i]")
// 	// test
// 	if x.f.Test != nil && Opt.IsArr() {
// 		r.T.Manual = true
// 		r.T.RxrOnly = true
// 		r.T.TstCnd = "x.Cnt() > 1"
// 		r.T.Add("expected := x.Cpy()")
// 		r.T.Add("axCpy := x.Cpy()")
// 		r.T.Add("i, j := x.FstIdx(), x.LstIdx()")
// 		r.T.Add("(*expected)[i], (*expected)[j] = (*expected)[j], (*expected)[i]")
// 		r.T.Add("axCpy.Swp(i, j)")
// 		r.T.Addf("%v(t, expected, axCpy)", x.f.Tst.Eql.Ref(x.f.Test))
// 	}
// 	return r
// }
