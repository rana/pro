package tpl

import (
	"sys/k"
	"sys/tpl/atr"
)

type (
	PrtArrAgg struct {
		PrtBse
		Arr *Arr
	}
)

func (x *PrtArrAgg) InitPrtTyp() {
	x.Arr = x.f.GetPrt((*PrtArr)(nil)).(*PrtArr).Arr
}
func (x *PrtArrAgg) InitPrtTypFn() {
	x.Sum()
	x.Prd()
	x.Min()
	x.Max()
	x.MinMax()
	x.Mid()
	x.Mdn()
	x.Sma()
	x.Gma()
	x.Wma()
	x.Vrnc()
	x.Std()
	x.Zscr()
	x.ZscrInplace()
	x.RngFul()
	x.RngLst()
	x.ProLst()
	x.ProSma()
	if x.Arr.Alias.Elm == _sys.Bsc.Flt.Typ() { // CORRRECT: ?
		x.SubPosSum()
		x.SubSumNeg()
		x.Rsi()
		x.Wrsi()
		x.Pro()
		x.Alma()
		x.ProAlma()
		x.CntrDist()
		x.ToBscAliased()
	}
}
func (x *PrtArrAgg) Sum() (r *TypFn) {
	r = x.f.TypFn(k.Sum)
	r.OutPrm(x.Arr.Elm, "r")
	r.Add("for n := 0; n < len(*x); n++ {")
	r.Add("r += (*x)[n]")
	r.Add("}")
	r.Add("return r")
	// test
	if x.f.Test != nil && Opt.IsArr() {
		r.T.SkpExcpy = true
		r.T.Addf("var expected %v", r.OutTyp().Full())
		r.T.Add("for n := 0; n < len(*x); n++ {")
		r.T.Add("expected += (*x)[n]")
		r.T.Add("}")
	}
	return r
}
func (x *PrtArrAgg) Prd() (r *TypFn) {
	r = x.f.TypFn(k.Prd)
	r.OutPrm(x.Arr.Elm, "r")
	r.Add("if len(*x) == 0 {")
	r.Add("return 0")
	r.Add("}")
	r.Add("r = (*x)[0]")
	r.Add("for n := 1; n < len(*x); n++ {")
	r.Add("r *= (*x)[n]")
	r.Add("}")
	r.Add("return r")
	// test
	if x.f.Test != nil && Opt.IsArr() {
		r.T.SkpExcpy = true
		r.T.Addf("var expected %v", r.OutTyp().Full())
		r.T.Add("if len(*x) != 0 {")
		r.T.Add("expected = (*x)[0]")
		r.T.Add("for n := 1; n < len(*x); n++ {")
		r.T.Add("expected *= (*x)[n]")
		r.T.Add("}")
		r.T.Add("}")
	}
	return r
}
func (x *PrtArrAgg) Min() (r *TypFn) {
	r = x.f.TypFn(k.Min)
	r.OutPrm(x.Arr.Elm, "r")
	r.Add("if len(*x) == 0 {")
	r.Add("return 0")
	r.Add("}")
	r.Add("r = (*x)[0]")
	r.Add("for n := 1; n < len(*x); n++ {")
	r.Add("if (*x)[n] < r {")
	r.Add("r = (*x)[n]")
	r.Add("}")
	r.Add("}")
	r.Add("return r")
	// test
	if x.f.Test != nil && Opt.IsArr() {
		r.T.SkpExcpy = true
		r.T.Addf("var expected %v", r.OutTyp().Full())
		r.T.Add("if len(*x) != 0 {")
		r.T.Add("expected = (*x)[0]")
		r.T.Add("for n := 1; n < len(*x); n++ {")
		r.T.Add("if (*x)[n] < expected {")
		r.T.Add("expected = (*x)[n]")
		r.T.Add("}")
		r.T.Add("}")
		r.T.Add("}")
	}
	return r
}
func (x *PrtArrAgg) Max() (r *TypFn) {
	r = x.f.TypFn(k.Max)
	r.OutPrm(x.Arr.Elm, "r")
	r.Add("if len(*x) == 0 {")
	r.Add("return 0")
	r.Add("}")
	r.Add("r = (*x)[0]")
	r.Add("for n := 1; n < len(*x); n++ {")
	r.Add("if (*x)[n] > r {")
	r.Add("r = (*x)[n]")
	r.Add("}")
	r.Add("}")
	r.Add("return r")
	// test
	if x.f.Test != nil && Opt.IsArr() {
		r.T.SkpExcpy = true
		r.T.Addf("var expected %v", r.OutTyp().Full())
		r.T.Add("if len(*x) != 0 {")
		r.T.Add("expected = (*x)[0]")
		r.T.Add("for n := 1; n < len(*x); n++ {")
		r.T.Add("if (*x)[n] > expected {")
		r.T.Add("expected = (*x)[n]")
		r.T.Add("}")
		r.T.Add("}")
		r.T.Add("}")
	}
	return r
}
func (x *PrtArrAgg) MinMax() (r *TypFn) {
	// x.f.Import(_sys)
	// x.f.Import("sys/bsc/bnds")
	// minMax: seg
	seg := x.f.StructPtrf("%vSeg", atr.None, k.MinMax)
	seg.FldTyp(_sys.Bsc.Bnd)
	seg.Fld("Min", x.Arr.Elm)
	seg.Fld("Max", x.Arr.Elm)
	seg.Fld("Vals", x.t)
	// minMax: seg act
	segAct := x.f.TypFn("Act", seg)
	segAct.Add("for n := x.Idx; n < x.Lim; n++ {")
	segAct.Add("if (*x.Vals)[n] < x.Min {")
	segAct.Add("x.Min = (*x.Vals)[n]")
	segAct.Add("}")
	segAct.Add("if (*x.Vals)[n] > x.Max {")
	segAct.Add("x.Max = (*x.Vals)[n]")
	segAct.Add("}")
	segAct.Add("}")

	r = x.f.TypFn(k.MinMax)
	r.OutPrm(x.Arr.Elm, "min")
	r.OutPrm(x.Arr.Elm, "max")
	r.Add("if len(*x) == 0 {")
	r.Add("return 0, 0")
	r.Add("}")
	r.Add("if len(*x) == 1 {")
	r.Add("return (*x)[0], (*x)[0]")
	r.Add("}")
	r.Add("min, max = (*x)[0], (*x)[0]")

	// r.Add("segBnds, acts := bnds.Segs(unt.Unt(len(*x)))")
	// r.Add("for n, segBnd := range *segBnds {")
	// r.Addf("acts[n] = %v{", seg.Adr(x.f))
	// r.Add("Bnd: segBnd,")
	// r.Add("Min: min,")
	// r.Add("Max: max,")
	// r.Add("Vals: x,")
	// r.Add("}")
	// r.Add("}")
	// r.Add("sys.Run().Pll(acts...) // process segments in pll")
	// r.Add("for _, act := range acts { // gather")
	// r.Addf("seg := act%v", seg.Cast(x.f, true))
	// r.Add("if seg.Min < min {")
	// r.Add("min = seg.Min")
	// r.Add("}")
	// r.Add("if seg.Max > max {")
	// r.Add("max = seg.Max")
	// r.Add("}")
	// r.Add("}")

	r.Add("for n := 1; n < len(*x); n++ {")
	r.Add("if (*x)[n] < min {")
	r.Add("min = (*x)[n]")
	r.Add("}")
	r.Add("if (*x)[n] > max {")
	r.Add("max = (*x)[n]")
	r.Add("}")
	r.Add("}")
	r.Add("return min, max")
	if x.f.Test != nil && Opt.IsArr() {
		r.T.Manual = true
		x.f.Test.Import("sys/app")
		r.T.Add("ap := app.New(tst.Cfg)")
		r.T.Add("defer ap.Cls()")
		r.T.Add("aMin, aMax := x.MinMax()")
		r.T.Add("eMin := x.Min()")
		r.T.Add("eMax := x.Max()")

		r.T.Addf("%v(t, eMin, aMin, \"Min\")", x.Arr.Elm.Bse().TstEql.Ref(x.f.Test))
		r.T.Addf("%v(t, eMax, aMax, \"Max\")", x.Arr.Elm.Bse().TstEql.Ref(x.f.Test))
	}
	return r
}
func (x *PrtArrAgg) Mid() (r *TypFn) {
	r = x.f.TypFn(k.Mid)
	r.OutPrm(x.Arr.Elm, "r")
	r.Add("if len(*x) == 0 {")
	r.Add("return r")
	r.Add("}")
	r.Add("if len(*x) == 1 {")
	r.Add("return (*x)[0]")
	r.Add("}")
	r.Add("min, max := x.MinMax()")
	r.Add("return min + ((max - min) / 2)")
	// test
	if x.f.Test != nil && Opt.IsArr() {
		r.T.Addf("var expected, min, max %v", r.OutTyp().Full())
		r.T.Add("min = exCpy.Min()")
		r.T.Add("max = exCpy.Max()")
		r.T.Add("expected = min + ((max - min) / 2)")
		r.T.Addf("%v(t, exCpy, axCpy, \"Source\")", x.f.Typ().Bse().TstEql.Ref(x.f.Test))
	}
	return r
}
func (x *PrtArrAgg) Mdn() (r *TypFn) {
	r = x.f.TypFn(k.Mdn)
	r.OutPrm(x.Arr.Elm, "r")
	r.Add("if len(*x) == 0 {")
	r.Add("return r")
	r.Add("}")
	r.Add("if len(*x) == 1 {")
	r.Add("return (*x)[0]")
	r.Add("}")
	r.Add("cpy := x.Cpy()")
	r.Add("cpy.SrtAsc()")
	r.Add("if len(*cpy)%2 == 0 {")
	r.Add("if len(*cpy) == 2 {")
	r.Add("return (cpy.At(0) + cpy.At(1)) / 2")
	r.Add("}")
	r.Add("return (cpy.Mdl() + cpy.At(cpy.MdlIdx()+1)) / 2")
	r.Add("}")
	r.Add("return cpy.Mdl()")
	// test
	if x.f.Test != nil && Opt.IsArr() {
		r.T.Addf("var expected %v", r.OutTyp().Full())
		r.T.Add("switch len(*exCpy){")
		r.T.Add("case 0:")
		r.T.Add("case 1:")
		r.T.Add("expected = (*exCpy)[0]")
		r.T.Add("default:")
		r.T.Add("cpy := exCpy.Cpy()")
		r.T.Add("cpy.SrtAsc()")
		r.T.Add("if len(*cpy)%2 == 0 {")
		r.T.Add("if len(*cpy) == 2 {")
		r.T.Add("expected = (cpy.At(0) + cpy.At(1)) / 2")
		r.T.Add("} else {")
		r.T.Add("expected = (cpy.Mdl() + cpy.At(cpy.MdlIdx()+1)) / 2")
		r.T.Add("}")
		r.T.Add("} else {")
		r.T.Add("expected = cpy.Mdl()")
		r.T.Add("}")
		r.T.Add("}")
		r.T.Addf("%v(t, exCpy, axCpy, \"Source\")", x.f.Typ().Bse().TstEql.Ref(x.f.Test))
	}
	return r
}
func (x *PrtArrAgg) Sma() (r *TypFn) {
	r = x.f.TypFn(k.Sma)
	r.OutPrm(x.Arr.Elm, "r")
	r.Add("if len(*x) == 0 { // simple moving average")
	r.Add("return r")
	r.Add("}")
	r.Add("if len(*x) == 1 {")
	r.Add("return (*x)[0]")
	r.Add("}")
	r.Addf("return x.Sum() / %v(len(*x))", r.OutTyp().Ref(x.f))
	// test
	if x.f.Test != nil && Opt.IsArr() {
		r.T.Addf("var expected %v", r.OutTyp().Full())
		r.T.Add("for _, v := range *exCpy {")
		r.T.Add("expected += v")
		r.T.Add("}")
		r.T.Add("if len(*exCpy) > 0 {")
		r.T.Addf("expected /= %v(len(*exCpy))", r.OutTyp().Ref(x.f.Test))
		r.T.Add("}")
		r.T.Addf("%v(t, exCpy, axCpy, \"Source\")", x.f.Typ().Bse().TstEql.Ref(x.f.Test))
	}
	return r
}
func (x *PrtArrAgg) Gma() (r *TypFn) {
	r = x.f.TypFn(k.Gma)
	r.OutPrm(x.Arr.Elm, "r")
	r.Add("if len(*x) == 0 { // geometric moving average")
	r.Add("return r")
	r.Add("}")
	r.Add("if len(*x) == 1 {")
	r.Add("return (*x)[0]")
	r.Add("}")
	r.Addf("return x.Prd().Pow(%v(1)) / %v(len(*x))", r.OutTyp().Ref(x.f), r.OutTyp().Ref(x.f))
	// test
	if x.f.Test != nil && Opt.IsArr() {
		r.T.Addf("var expected %v", r.OutTyp().Full())
		r.T.Add("for _, v := range *exCpy {")
		r.T.Add("expected *= v")
		r.T.Add("}")
		r.T.Add("if len(*exCpy) > 0 {")
		r.T.Addf("expected = expected.Pow(%v(1))", r.OutTyp().Ref(x.f.Test))
		r.T.Addf("expected /= %v(len(*exCpy))", r.OutTyp().Ref(x.f.Test))
		r.T.Add("}")
		r.T.Addf("%v(t, exCpy, axCpy, \"Source\")", x.f.Typ().Bse().TstEql.Ref(x.f.Test))
	}
	return r
}
func (x *PrtArrAgg) Wma() (r *TypFn) {
	r = x.f.TypFn(k.Wma)
	r.OutPrm(x.Arr.Elm, "r")
	r.Add("// For example, a 5 period WMA would be calculated as:")
	r.Add("// WMA = (P1 * 1) + (P2 * 2) + (P3 * 3) + (P4 * 4) + (P5 * 5) / (1 + 2 + 3 + 4 + 5)")
	r.Add("if len(*x) == 0 { // weighted moving average")
	r.Add("return 0")
	r.Add("}")
	r.Add("if len(*x) == 1 {")
	r.Add("return (*x)[0]")
	r.Add("}")
	r.Addf("var numr, dnmr %v", r.OutTyp().Ref(x.f))
	r.Add("for n, v := range *x {")
	r.Addf("numr += v * %v(n+1)", r.OutTyp().Ref(x.f))
	r.Addf("dnmr += %v(n + 1)", r.OutTyp().Ref(x.f))
	r.Add("}")
	r.Add("if dnmr == 0 {")
	r.Add("return 0")
	r.Add("}")
	r.Add("return numr / dnmr")
	// test
	if x.f.Test != nil && Opt.IsArr() {
		r.T.Addf("var numerator, denominator, expected %v", r.OutTyp().Full())
		r.T.Add("for n, v := range *exCpy {")
		r.T.Addf("numerator += v * %v(n+1)", r.OutTyp().Full())
		r.T.Addf("denominator += %v(n + 1)", r.OutTyp().Full())
		r.T.Add("}")
		r.T.Add("if len(*exCpy) > 0 {")
		r.T.Addf("expected = numerator / denominator")
		r.T.Add("}")
		r.T.Addf("%v(t, exCpy, axCpy, \"Source\")", x.f.Typ().Bse().TstEql.Ref(x.f.Test))
	}
	return r
}

func (x *PrtArrAgg) SubPosSum() (r *TypFn) {
	r = x.f.TypFn(k.SubSumPos)
	r.OutPrm(x.Arr.Elm, "r")
	r.Add("// positive sum of subtracted elements")
	r.Add("if len(*x) <= 1 {")
	r.Add("return 0")
	r.Add("}")

	r.Add("// skp eql is only issue for calculation on tic streams")
	r.Add("var prv, cur int")
	r.Add("for ; prv < len(*x)-1; prv = cur {")
	r.Add("cur = prv + 1")
	r.Add("if (*x)[cur] == (*x)[prv] { // skp eql due to gap fil; may alter result if valid, non-gap eql present")
	r.Add("continue")
	r.Add("}")

	r.Add("if (*x)[cur]-(*x)[prv] > 0 {")
	r.Add("r += (*x)[cur] - (*x)[prv]")
	r.Add("}")

	r.Add("}")

	r.Add("return r")
	// test
	if x.f.Test != nil && Opt.IsArr() {
		r.T.SkpExcpy = true
		r.T.Addf("var expected %v", r.OutTyp().Full())
		r.T.Add("if len(*x) > 1 {")

		r.T.Add("var prv, cur int")
		r.T.Add("for ; prv < len(*x)-1; prv = cur {")
		r.T.Add("cur = prv + 1")
		r.T.Add("if (*x)[cur] == (*x)[prv] { // skp eql due to gap fil; may alter result if valid, non-gap eql present")
		r.T.Add("continue")
		r.T.Add("}")

		r.T.Add("if (*x)[cur]-(*x)[prv] > 0 {")
		r.T.Add("expected += (*x)[cur] - (*x)[prv]")
		r.T.Add("}")

		r.T.Add("}")

		r.T.Add("}")
	}
	return r
}

func (x *PrtArrAgg) SubSumNeg() (r *TypFn) {
	r = x.f.TypFn(k.SubSumNeg)
	r.OutPrm(x.Arr.Elm, "r")
	r.Add("// negative sum of subtracted elements")
	r.Add("if len(*x) <= 1 {")
	r.Add("return 0")
	r.Add("}")

	r.Add("// skp eql is only issue for calculation on tic streams")
	r.Add("var prv, cur int")
	r.Add("for ; prv < len(*x)-1; prv = cur {")
	r.Add("cur = prv + 1")
	r.Add("if (*x)[cur] == (*x)[prv] { // skp eql due to gap fil; may alter result if valid, non-gap eql present")
	r.Add("continue")
	r.Add("}")

	r.Add("if (*x)[cur]-(*x)[prv] < 0 {")
	r.Add("r += (*x)[cur] - (*x)[prv]")
	r.Add("}")

	r.Add("}")

	r.Add("return r")
	// test
	if x.f.Test != nil && Opt.IsArr() {
		r.T.SkpExcpy = true
		r.T.Addf("var expected %v", r.OutTyp().Full())
		r.T.Add("if len(*x) > 1 {")

		r.T.Add("var prv, cur int")
		r.T.Add("for ; prv < len(*x)-1; prv = cur {")
		r.T.Add("cur = prv + 1")
		r.T.Add("if (*x)[cur] == (*x)[prv] { // skp eql due to gap fil; may alter result if valid, non-gap eql present")
		r.T.Add("continue")
		r.T.Add("}")

		r.T.Add("if (*x)[cur]-(*x)[prv] < 0 {")
		r.T.Add("expected += (*x)[cur] - (*x)[prv]")
		r.T.Add("}")

		r.T.Add("}")

		r.T.Add("}")
	}
	return r
}

func (x *PrtArrAgg) Rsi() (r *TypFn) {
	r = x.f.TypFn(k.Rsi)
	r.OutPrm(x.Arr.Elm, "r")
	r.Add("// relative strength index")
	r.Add("// RS = Average Gain / Average Loss")
	r.Add("//               100")
	r.Add("// RSI = 100 - --------")
	r.Add("//              1 + RS")
	r.Add("// NOTE: This impl has a scale of 0.0 to 1.0")
	r.Add("//       and returns 0.5 as a default value")

	r.Add("if len(*x) <= 1 {")
	r.Add("return .5 // return mdl")
	r.Add("}")

	r.Add("// skp eql is only issue for calculation on tic streams")
	r.Add("var neg, pos flt.Flt")
	r.Add("var prv, cur int")
	r.Add("for ; prv < len(*x)-1; prv = cur {")
	r.Add("cur = prv + 1")
	r.Add("if (*x)[cur] == (*x)[prv] { // skp eql due to gap fil; may alter result if valid, non-gap eql present")
	r.Add("continue")
	r.Add("}")

	r.Add("if (*x)[cur]-(*x)[prv] < 0 {")
	r.Add("neg -= (*x)[cur] - (*x)[prv]")
	r.Add("} else {")
	r.Add("pos += (*x)[cur] - (*x)[prv]")
	r.Add("}")

	r.Add("}")

	r.Add("if neg == 0 && pos == 0 {")
	r.Add("return .5 // return mdl")
	r.Add("}")
	r.Add("return pos / (pos + neg) // rng is 0 to 1")
	// test
	if x.f.Test != nil && Opt.IsArr() {
		r.T.SkpExcpy = true
		r.T.Addf("var expected %v", r.OutTyp().Full())
		r.T.Add("if len(*x) <= 1 {")
		r.T.Add("expected = .5")
		r.T.Add("} else {")

		r.T.Add("var neg, pos flt.Flt")
		r.T.Add("var prv, cur int")
		r.T.Add("for ; prv < len(*x)-1; prv = cur {")
		r.T.Add("cur = prv + 1")
		r.T.Add("if (*x)[cur] == (*x)[prv] { // skp eql due to gap fil; may alter result if valid, non-gap eql present")
		r.T.Add("continue")
		r.T.Add("}")

		r.T.Add("if (*x)[cur]-(*x)[prv] < 0 {")
		r.T.Add("neg -= (*x)[cur] - (*x)[prv]")
		r.T.Add("} else {")
		r.T.Add("pos += (*x)[cur] - (*x)[prv]")
		r.T.Add("}")

		r.T.Add("}")

		r.T.Add("if neg == 0 && pos == 0 {")
		r.T.Add("expected = .5 // return mdl")
		r.T.Add("} else {")
		r.T.Add("expected = pos / (pos + neg) // rng is 0 to 1")
		r.T.Add("}")
		r.T.Add("}")
	}
	return r
}

func (x *PrtArrAgg) Wrsi() (r *TypFn) {
	r = x.f.TypFn(k.Wrsi)
	r.OutPrm(x.Arr.Elm, "r")

	r.Add("if len(*x) <= 1 {")
	r.Add("return .5 // return mdl")
	r.Add("}")
	r.Add("var neg, pos flt.Flt")

	r.Add("for n := 1; n < len(*x); n++ {")
	r.Add("if (*x)[n]-(*x)[n-1] < 0 {")
	r.Add("neg -= ((*x)[n] - (*x)[n-1]) * flt.Flt(n+1)")
	r.Add("} else {")
	r.Add("pos += ((*x)[n] - (*x)[n-1]) * flt.Flt(n+1)")
	r.Add("}")
	r.Add("}")

	r.Add("if neg == 0 && pos == 0 {")
	r.Add("return .5 // return mdl")
	r.Add("}")
	r.Add("return pos / (pos + neg) // rng is 0 to 1")

	return r
}

func (x *PrtArrAgg) Vrnc() (r *TypFn) {
	r = x.f.TypFn(k.Vrnc)
	r.OutPrm(x.Arr.Elm, "r")
	r.Add("if len(*x) == 0 {")
	r.Add("return 0")
	r.Add("}")
	r.Add("mean := x.Sma()")
	r.Add("meanDifSqrSum := r")
	r.Add("for _, v := range *x {")
	r.Add("meanDifSqrSum += (mean - v) * (mean - v) // calculate mean dif sqr sum")
	r.Add("}")
	r.Addf("return meanDifSqrSum / %v(len(*x)) // calculate variance", r.OutTyp().Full())
	// test
	if x.f.Test != nil && Opt.IsArr() {
		r.T.Addf("var expected %v", r.OutTyp().Full())
		r.T.Add("if len(*exCpy) != 0 {")
		r.T.Add("mean := exCpy.Sma()")
		r.T.Add("meanDifSqrSum := expected")
		r.T.Add("for _, v := range *exCpy {")
		r.T.Add("meanDif := mean - v")
		r.T.Add("meanDifSqrSum += meanDif * meanDif // calculate mean dif sqr sum")
		r.T.Add("}")
		r.T.Addf("expected = meanDifSqrSum / %v(len(*exCpy)) // calculate variance", r.OutTyp().Full())
		r.T.Add("}")
		r.T.Addf("%v(t, exCpy, axCpy, \"Source\")", x.f.Typ().Bse().TstEql.Ref(x.f.Test))
	}
	return r
}
func (x *PrtArrAgg) Std() (r *TypFn) {
	r = x.f.TypFn(k.Std)
	r.OutPrm(x.Arr.Elm, "r")
	r.Add("if len(*x) == 0 {")
	r.Add("return 0")
	r.Add("}")
	r.Add("return x.Vrnc().Sqrt()")
	// test
	if x.f.Test != nil && Opt.IsArr() {
		r.T.Addf("var expected %v", r.OutTyp().Full())
		r.T.Add("if len(*exCpy) != 0 {")
		r.T.Add("expected = exCpy.Vrnc().Sqrt()")
		r.T.Add("}")
		r.T.Addf("%v(t, exCpy, axCpy, \"Source\")", x.f.Typ().Bse().TstEql.Ref(x.f.Test))
	}
	return r
}
func (x *PrtArrAgg) Zscr() (r *TypFn) {
	r = x.f.TypFn(k.Zscr)
	r.OutPrm(x.t, "r")
	r.Add("if len(*x) == 0 {")
	r.Add("return New()")
	r.Add("}")
	r.Add("r = Make(x.Cnt())")
	r.Add("mean := x.Sma()")
	r.Add("std := x.Std()")
	r.Add("for n, v := range *x {")
	r.Add("(*r)[n] = (v - mean) / std")
	r.Add("}")
	r.Add("return r")
	return r
}
func (x *PrtArrAgg) ZscrInplace() (r *TypFn) {
	r = x.f.TypFn("ZscrInplace")
	r.OutPrm(x.t, "r")
	r.Add("if len(*x) == 0 {")
	r.Add("return x")
	r.Add("}")
	r.Add("mean := x.Sma()")
	r.Add("std := x.Std()")
	r.Add("for n, v := range *x {")
	r.Add("(*x)[n] = (v - mean) / std")
	r.Add("}")
	r.Add("return x")
	return r
}
func (x *PrtArrAgg) RngFul() (r *TypFn) {
	r = x.f.TypFn(k.RngFul)
	r.OutPrm(x.Arr.Elm, "r")
	r.Add("if len(*x) == 0 {")
	r.Add("return 0")
	r.Add("}")
	r.Add("min, max := x.MinMax()")
	r.Add("return max - min")
	// test
	if x.f.Test != nil && Opt.IsArr() {
		r.T.Addf("var expected %v", r.OutTyp().Full())
		r.T.Add("if len(*exCpy) != 0 {")
		r.T.Add("min, max := exCpy.MinMax()")
		r.T.Add("expected = max - min")
		r.T.Add("}")
		r.T.Addf("%v(t, exCpy, axCpy, \"Source\")", x.f.Typ().Bse().TstEql.Ref(x.f.Test))
	}
	return r
}
func (x *PrtArrAgg) RngLst() (r *TypFn) {
	r = x.f.TypFn(k.RngLst)
	r.OutPrm(x.Arr.Elm, "r")
	r.Add("if len(*x) == 0 {")
	r.Add("return 0")
	r.Add("}")
	r.Add("return x.Lst() - x.Min()")
	// test
	if x.f.Test != nil && Opt.IsArr() {
		r.T.Addf("var expected %v", r.OutTyp().Full())
		r.T.Add("if len(*exCpy) != 0 {")
		r.T.Add("expected = exCpy.Lst() - exCpy.Min()")
		r.T.Add("}")
		r.T.Addf("%v(t, exCpy, axCpy, \"Source\")", x.f.Typ().Bse().TstEql.Ref(x.f.Test))
	}
	return r
}
func (x *PrtArrAgg) ProLst() (r *TypFn) {
	r = x.f.TypFn(k.ProLst)
	r.OutPrm(x.Arr.Elm, "r")
	r.Add("if len(*x) == 0 {")
	r.Add("return 0")
	r.Add("}")
	r.Add("min, max := x.MinMax()")
	r.Add("rngFul := max - min")
	r.Add("if rngFul == 0 {")
	r.Add("return r")
	r.Add("}")
	r.Add("return (x.Lst() - min) / rngFul")
	// test
	if x.f.Test != nil && Opt.IsArr() {
		r.T.Addf("var expected %v", r.OutTyp().Full())
		r.T.Add("if len(*exCpy) != 0 {")
		r.T.Add("min, max := exCpy.MinMax()")
		r.T.Add("rngFul := max - min")
		r.T.Add("if rngFul != 0 {")
		r.T.Add("expected = (exCpy.Lst() - min) / rngFul")
		r.T.Add("}")
		r.T.Add("}")
		r.T.Addf("%v(t, exCpy, axCpy, \"Source\")", x.f.Typ().Bse().TstEql.Ref(x.f.Test))
	}
	return r
}
func (x *PrtArrAgg) ProSma() (r *TypFn) {
	r = x.f.TypFn(k.ProSma)
	r.OutPrm(x.Arr.Elm, "r")
	r.Add("if len(*x) == 0 {")
	r.Add("return 0")
	r.Add("}")
	r.Add("min, max := x.MinMax()")
	r.Add("rngFul := max - min")
	r.Add("if rngFul == 0 {")
	r.Add("return r")
	r.Add("}")
	r.Add("return (x.Sma() - min) / rngFul")
	// test
	if x.f.Test != nil && Opt.IsArr() {
		r.T.Addf("var expected %v", r.OutTyp().Full())
		r.T.Add("if len(*exCpy) != 0 {")
		r.T.Add("min, max := exCpy.MinMax()")
		r.T.Add("rngFul := max - min")
		r.T.Add("if rngFul != 0 {")
		r.T.Add("expected = (exCpy.Sma() - min) / rngFul")
		r.T.Add("}")
		r.T.Add("}")
		r.T.Addf("%v(t, exCpy, axCpy, \"Source\")", x.f.Typ().Bse().TstEql.Ref(x.f.Test))
	}
	return r
}
func (x *PrtArrAgg) ProAlma() (r *TypFn) {
	r = x.f.TypFn(k.ProAlma)
	r.OutPrm(x.Arr.Elm, "r")
	r.Add("if len(*x) == 0 {")
	r.Add("return 0")
	r.Add("}")
	r.Add("min, max := x.MinMax()")
	r.Add("rngFul := max - min")
	r.Add("if rngFul == 0 {")
	r.Add("return r")
	r.Add("}")
	r.Add("return (x.Alma() - min) / rngFul")
	// test
	if x.f.Test != nil && Opt.IsArr() {
		r.T.Addf("var expected %v", r.OutTyp().Full())
		r.T.Add("if len(*exCpy) != 0 {")
		r.T.Add("min, max := exCpy.MinMax()")
		r.T.Add("rngFul := max - min")
		r.T.Add("if rngFul != 0 {")
		r.T.Add("expected = (exCpy.Alma() - min) / rngFul")
		r.T.Add("}")
		r.T.Add("}")
		r.T.Addf("%v(t, exCpy, axCpy, \"Source\")", x.f.Typ().Bse().TstEql.Ref(x.f.Test))
	}
	return r
}
func (x *PrtArrAgg) Pro() (r *TypFn) {
	r = x.f.TypFn(k.Pro)
	r.OutPrm(x.t, "r")
	r.Add("if len(*x) == 0 {")
	r.Addf("return %v(0) // IMPORTANT FOR PLT TO RETURN EMPTY (NOT NIL)", _sys.Bsc.Flt.arr.Make.Ref(x.f))
	r.Add("}")
	r.Add("min, max := x.MinMax()")
	r.Add("rngFul := max - min")
	r.Add("if rngFul == 0 {")
	r.Addf("return %v(x.Cnt()) // IMPORTANT FOR PLT TO RETURN EMPTY (NOT NIL)", _sys.Bsc.Flt.arr.Make.Ref(x.f))
	r.Add("}")
	r.Addf("pros := make(%v, len(*x))", x.t.Title())
	r.Add("for n := 0; n < len(pros); n++ {")
	r.Add("pros[n] = ((*x)[n]-min)/rngFul")
	r.Add("}")
	r.Add("return &pros")
	// test
	if x.f.Test != nil && Opt.IsArr() {
		r.T.Manual = true
		x.f.Test.Import("sys/app")
		r.T.Add("ap := app.New(tst.Cfg)")
		r.T.Add("defer ap.Cls()")
		r.T.Addf("expected := %v(x.Cnt())", _sys.Bsc.Flt.arr.Make.Ref(x.f.Test))

		r.T.Add("if len(*x) != 0 {")
		r.T.Add("min, max := x.MinMax()")
		r.T.Add("rngFul := max - min")
		r.T.Add("if rngFul != 0 {")
		r.T.Add("for n := 0; n < len(*expected); n++ {")
		r.T.Add("(*expected)[n] = ((*x)[n]-min)/rngFul")
		r.T.Add("}")
		r.T.Add("}")
		r.T.Add("}")

		r.T.Addf("actual := x.%v()", r.Name)
		// r.T.Add("if expected.Cnt() == 0 {")
		// r.T.Add("expected = nil")
		// r.T.Add("}")
		r.T.Addf("%v(t, expected, actual)", x.f.Typ().Bse().TstEql.Ref(x.f.Test))
	}
	return r
}

func (x *PrtArrAgg) Alma() (r *TypFn) {
	x.f.Import("math")
	r = x.f.TypFn(k.Alma)
	// r.InPrm(_sys.Bsc.Unt, "sigma").LitVal("6")
	// r.InPrm(_sys.Bsc.Flt, "offset").LitVal("0.85")
	r.OutPrm(x.Arr.Elm, "alma")
	r.Add("if len(*x) == 0 {")
	r.Add("return 0")
	r.Add("}")
	r.Add("// http://www.financial-hacker.com/trend-delusion-or-reality/")
	r.Add("// https://www.prorealcode.com/prorealtime-indicators/alma-arnaud-legoux-moving-average/")
	r.Add("// Window = 9")
	r.Add("// Sigma = 6")
	r.Add("// Offset = 0.85")
	r.Add("// var ALMA(var *Data, int Period) {")
	r.Add("// var m = floor(0.85*(Period-1));")
	r.Add("// var s = Period/6.0;")
	r.Add("// var alma = 0., wSum = 0.;")
	r.Add("// int i;")
	r.Add("// for (i = 0; i < Period; i++) {")
	r.Add("// var w = exp(-(i-m)*(i-m)/(2*s*s));")
	r.Add("// alma += Data[Period-1-i] * w;")
	r.Add("// wSum += w;")
	r.Add("// }")
	r.Add("// return alma / wSum;")
	r.Add("// }")
	r.Addf("const sigma = %v(6)", _sys.Bsc.Flt.Ref(x.f))
	r.Addf("const offset = %v(0.85)", _sys.Bsc.Flt.Ref(x.f))
	r.Add("m := offset * flt.Flt(len(*x)-1)")
	r.Add("s := flt.Flt(len(*x))/sigma")
	r.Add("ss2 := s * s * 2")
	r.Add("var wSum flt.Flt")
	r.Add("for i := 0; i < len(*x)-1; i++ {")
	r.Add("im := flt.Flt(i) - m")
	r.Add("w := flt.Flt(math.Exp(float64(-(im * im) / ss2)))")
	r.Add("alma += (*x)[len(*x)-1-i] * w")
	r.Add("wSum += w")
	r.Add("}")
	r.Add("return alma / wSum")

	return r
}
func (x *PrtArrAgg) CntrDist() (r *TypFn) {
	x.f.Import(_sys)
	x.f.Import(_sys.Bsc.Bnd.arr)
	// CntrDist: seg
	seg := x.f.StructPtrf("%vSeg", atr.None, k.CntrDist)
	seg.FldTyp(_sys.Bsc.Bnd)
	seg.Fld("EvalZero", _sys.Bsc.Bol)
	seg.Fld("RngFul", _sys.Bsc.Flt)
	seg.Fld("Vals", _sys.Bsc.Flt.arr)
	seg.Fld("Out", _sys.Bsc.Flt.arr)
	// CntrDist: seg act
	segAct := x.f.TypFn("Act", seg)
	segAct.Add("for m := x.Idx; m < x.Lim; m++ {")
	segAct.Add("for n := unt.Zero; n < x.Vals.Cnt(); n++ {")
	segAct.Add("if m != n && x.Vals.At(n).Neq(flt.Max) {")
	segAct.Add("if !x.EvalZero && x.Vals.At(n).Eql(flt.Zero) {")
	segAct.Add("continue")
	segAct.Add("}")
	segAct.Add("// out[m] += 1 - (abs(m-n)/rngFul)")
	segAct.Add("x.Out.Upd(m, x.Out.At(m)+flt.One.Sub(x.Vals.At(m).Sub(x.Vals.At(n)).Pos().Div(x.RngFul)))")
	segAct.Add("}")
	segAct.Add("}")
	segAct.Add("}")

	r = x.f.TypFn(k.CntrDist)
	r.InPrmVariadic(_sys.Bsc.Bol, "evalZero")
	r.OutPrm(x.t, "r")
	r.Addf("r = %v(x.Cnt())", _sys.Bsc.Flt.arr.Make.Ref(x.f))
	r.Add("rngFul := x.RngFul()")
	r.Add("segBnds, acts := bnds.Segs(x.Cnt())")
	r.Add("for n, segBnd := range *segBnds {")
	r.Addf("acts[n] = %v{", seg.Adr(x.f))
	r.Add("Bnd: segBnd,")
	r.Add("EvalZero: len(evalZero) != 0 && evalZero[0],")
	r.Add("RngFul: rngFul,")
	r.Add("Vals: x,")
	r.Add("Out: r,")
	r.Add("}")
	r.Add("}")
	r.Add("sys.Run().Pll(acts...)")
	r.Add("return r")
	if x.f.Test != nil && Opt.IsArr() {
		r.T.Manual = true
		x.f.Test.Import("sys/app")
		r.T.Add("ap := app.New(tst.Cfg)")
		r.T.Add("defer ap.Cls()")
		r.T.Addf("expected := %v(x.Cnt())", _sys.Bsc.Flt.arr.Make.Ref(x.f.Test))
		r.T.Add("rngFul := x.RngFul()")
		r.T.Add("for m := unt.Zero; m < x.Cnt(); m++ {")
		r.T.Add("for n := unt.Zero; n < x.Cnt(); n++ {")
		r.T.Add("if m != n && x.At(n).Neq(flt.Max) {")
		r.T.Add("if evalZero && x.At(n).Eql(flt.Zero) {")
		r.T.Add("continue")
		r.T.Add("}")
		r.T.Add("// out[m] += 1 - (abs(m-n)/rngFul)")
		r.T.Add("expected.Upd(m, expected.At(m)+flt.One.Sub(x.At(m).Sub(x.At(n)).Pos().Div(rngFul)))")
		r.T.Add("}")
		r.T.Add("}")
		r.T.Add("}")
		r.T.Addf("actual := x.%v()", r.Name)
		r.T.Addf("%v(t, expected, actual)", x.f.Typ().Bse().TstEql.Ref(x.f.Test))
	}
	return r
}

func (x *PrtArrAgg) ToBscAliased() (r *TypFn) {
	// expect to call for bsc typ only
	elmAlias := x.Arr.Alias.Elm.(*Alias).Elm.Bse()
	r = x.f.TypFnf("%vs", elmAlias.Name)
	r.OutPrmSlice(elmAlias, "r")
	r.Addf("r = make([]%v, len(*x))", elmAlias.Name)
	r.Add("for n := 0; n < len(*x); n++ {")
	r.Addf("r[n] = %v((*x)[n])", elmAlias.Name)
	r.Add("}")
	r.Add("return r")
	return r
}
