package tpl

import "sys/k"

type (
	PrtAri struct {
		PrtBse
		Add    *TypFn
		Sub    *TypFn
		Mul    *TypFn
		Div    *TypFn
		Rem    *TypFn
		Pow    *TypFn
		Sqr    *TypFn
		Sqrt   *TypFn
		Min    *TypFn
		Max    *TypFn
		MinMax *TypFn
		Mid    *TypFn
		Avg    *TypFn
		AvgGeo *TypFn
	}
)

func (x *PrtAri) InitPrtTypFn() {
	x.Add = x.add()
	x.Sub = x.sub()
	x.Mul = x.mul()
	x.Div = x.div()
	x.Rem = x.rem()
	x.Pow = x.pow()
	x.Sqr = x.sqr()
	x.Sqrt = x.sqrt()
	x.Min = x.min()
	x.Max = x.max()
	x.MinMax = x.minMax()
	x.Mid = x.mid()
	x.Avg = x.avg()
	x.AvgGeo = x.avgGeo()
}

func (x *PrtAri) add() (r *TypFn) {
	r = x.f.TypFn(k.Add)
	r.InPrm(x.f.Typ(), "a")
	r.OutPrm(x.f.Typ())
	r.Addf("return %v + %v", r.Rxr.Name, r.InPrms[0].Name)
	// test
	if x.f.Test != nil {
		r.T.Addf("expected := %v + %v", r.Rxr.Name, r.InPrms[0].Name)
	}
	return r
}
func (x *PrtAri) sub() (r *TypFn) {
	r = x.f.TypFn(k.Sub)
	r.InPrm(x.f.Typ(), "a")
	r.OutPrm(x.f.Typ())
	r.Addf("return %v - %v", r.Rxr.Name, r.InPrms[0].Name)
	// test
	if x.f.Test != nil {
		r.T.Addf("expected := %v - %v", r.Rxr.Name, r.InPrms[0].Name)
	}
	return r
}
func (x *PrtAri) mul() (r *TypFn) {
	r = x.f.TypFn(k.Mul)
	r.InPrm(x.f.Typ(), "a")
	r.OutPrm(x.f.Typ())
	r.Addf("return %v * %v", r.Rxr.Name, r.InPrms[0].Name)
	// test
	if x.f.Test != nil {
		r.T.Addf("expected := %v * %v", r.Rxr.Name, r.InPrms[0].Name)
	}
	return r
}
func (x *PrtAri) div() (r *TypFn) {
	r = x.f.TypFn(k.Div)
	r.InPrm(x.f.Typ(), "a")
	r.OutPrm(x.f.Typ())
	r.Addf("if %v == 0 {", r.InPrms[0].Name)
	r.Add("return 0")
	r.Add("} else {")
	r.Addf("return %v / %v", r.Rxr.Name, r.InPrms[0].Name)
	r.Add("}")
	// test
	if x.f.Test != nil {
		r.T.Addf("var expected %v", r.InPrms[0].Typ.Full())
		r.T.Addf("if %v != 0 {", r.InPrms[0].Name)
		r.T.Addf("expected = %v / %v", r.Rxr.Name, r.InPrms[0].Name)
		r.T.Add("}")
	}
	return r
}
func (x *PrtAri) rem() (r *TypFn) {
	r = x.f.TypFn(k.Rem)
	r.InPrm(x.f.Typ(), "a")
	r.OutPrm(x.f.Typ())
	r.CastOut = true
	r.Import("math")
	r.Addf("if %v == 0 {", r.InPrms[0].Name)
	r.Add("return 0")
	r.Add("} else {")
	r.Addf("return %v(math.Remainder(float64(%v), float64(%v)))", r.Rxr.Typ.Title(), r.Rxr.Name, r.InPrms[0].Name)
	r.Add("}")
	// test
	if x.f.Test != nil {
		x.f.Test.Import("math")
		x.f.Test.ImportFn(r)
		r.T.Addf("var expected %v", r.InPrms[0].Typ.Full())
		r.T.Addf("if %v != 0 {", r.InPrms[0].Name)
		r.T.Addf("expected = %v(math.Remainder(float64(%v), float64(%v)))", r.OutTyp().Full(), r.Rxr.Name, r.InPrms[0].Name)
		r.T.Add("}")
	}
	return r
}
func (x *PrtAri) pow() (r *TypFn) {
	x.f.Import("math")
	r = x.f.TypFn(k.Pow)
	r.CastOut = true
	r.Import("math")
	r.InPrm(x.f.Typ(), "a")
	r.OutPrm(x.f.Typ())
	r.Addf("return %v(math.Pow(float64(%v), float64(%v)))", r.Rxr.Typ.Title(), r.Rxr.Name, r.InPrms[0].Name)
	// test
	if x.f.Test != nil {
		x.f.Test.Import("math")
		x.f.Test.ImportFn(r)
		r.T.Addf("expected := %v(math.Pow(float64(%v), float64(%v)))", r.OutTyp().Full(), r.Rxr.Name, r.InPrms[0].Name)
	}
	return r
}
func (x *PrtAri) sqr() (r *TypFn) {
	r = x.f.TypFn(k.Sqr)
	r.OutPrm(x.f.Typ())
	r.Addf("return %v * %v", r.Rxr.Name, r.Rxr.Name)
	// test
	if x.f.Test != nil {
		r.T.Addf("expected := %v * %v", r.Rxr.Name, r.Rxr.Name)
	}
	return r
}
func (x *PrtAri) sqrt() (r *TypFn) {
	r = x.f.TypFn(k.Sqrt)
	r.OutPrm(x.f.Typ())
	r.CastOut = true
	r.Import("math")
	r.Addf("if x <= 0 {")
	r.Addf("return 0")
	r.Addf("} else {")
	r.Addf("return %v(math.Sqrt(float64(%v)))", r.OutTyp().Title(), r.Rxr.Name)
	r.Addf("}")
	// test
	if x.f.Test != nil {
		r.T.Addf("var expected %v", r.OutTyp().Full())
		r.T.Addf("if %v > 0 {", r.Rxr.Name)
		r.T.Addf("expected = %v(math.Sqrt(float64(%v)))", r.OutTyp().Full(), r.Rxr.Name)
		r.T.Add("}")
	}
	return r
}
func (x *PrtAri) min() (r *TypFn) {
	r = x.f.TypFn(k.Min)
	r.InPrm(x.f.Typ(), "a")
	r.OutPrm(x.f.Typ())
	r.Addf("if %v < %v {", r.Rxr.Name, r.InPrms[0].Name)
	r.Addf("return %v", r.Rxr.Name)
	r.Addf("} else {")
	r.Addf("return %v", r.InPrms[0].Name)
	r.Addf("}")
	// test
	if x.f.Test != nil {
		r.T.Addf("expected := %v", r.Rxr.Name)
		r.T.Addf("if %v < expected {", r.InPrms[0].Name)
		r.T.Addf("expected = %v", r.InPrms[0].Name)
		r.T.Addf("}")
	}
	return r
}
func (x *PrtAri) max() (r *TypFn) {
	r = x.f.TypFn(k.Max)
	r.InPrm(x.f.Typ(), "a")
	r.OutPrm(x.f.Typ())
	r.Addf("if %v > %v {", r.Rxr.Name, r.InPrms[0].Name)
	r.Addf("return %v", r.Rxr.Name)
	r.Addf("} else {")
	r.Addf("return %v", r.InPrms[0].Name)
	r.Addf("}")
	// test
	if x.f.Test != nil {
		r.T.Addf("expected := %v", r.Rxr.Name)
		r.T.Addf("if %v > expected {", r.InPrms[0].Name)
		r.T.Addf("expected = %v", r.InPrms[0].Name)
		r.T.Addf("}")
	}
	return r
}
func (x *PrtAri) minMax() (r *TypFn) {
	r = x.f.TypFn(k.MinMax)
	r.InPrm(x.f.Typ(), "a")
	r.OutPrm(x.f.Typ(), "min")
	r.OutPrm(x.f.Typ(), "max")
	r.Addf("if %v < %v {", r.Rxr.Name, r.InPrms[0].Name)
	r.Addf("return %v, %v", r.Rxr.Name, r.InPrms[0].Name)
	r.Addf("}")
	r.Addf("return %v, %v", r.InPrms[0].Name, r.Rxr.Name)
	// test
	if x.f.Test != nil {
		r.T.Manual = true
		r.T.Addf("aMin, aMax := %v.%v(%v)", r.Rxr.Name, r.Name, r.InPrms[0].Name)
		r.T.Addf("var eMin, eMax %v", r.OutTyp().Ref(x.f.Test))
		r.T.Addf("if %v < %v {", r.Rxr.Name, r.InPrms[0].Name)
		r.T.Addf("eMin, eMax = %v, %v", r.Rxr.Name, r.InPrms[0].Name)
		r.T.Addf("} else {")
		r.T.Addf("eMin, eMax = %v, %v", r.InPrms[0].Name, r.Rxr.Name)
		r.T.Addf("}")
		r.T.Addf("%v(t, eMin, aMin, \"Min\")", r.OutTyp().Bse().TstEql.Ref(x.f.Test))
		r.T.Addf("%v(t, eMax, aMax, \"Max\")", r.OutTyp().Bse().TstEql.Ref(x.f.Test))
	}
	return r
}
func (x *PrtAri) mid() (r *TypFn) {
	r = x.f.TypFn(k.Mid)
	r.InPrm(x.f.Typ(), "a")
	r.OutPrm(x.f.Typ())
	r.Addf("return (%v - %v) / 2", r.Rxr.Name, r.InPrms[0].Name)
	// test
	if x.f.Test != nil {
		r.T.Addf("expected := (%v - %v) / 2", r.Rxr.Name, r.InPrms[0].Name)
	}
	return r
}
func (x *PrtAri) avg() (r *TypFn) {
	r = x.f.TypFn(k.Avg)
	r.InPrm(x.f.Typ(), "a")
	r.OutPrm(x.f.Typ())
	r.Addf("return (%v + %v) / 2", r.Rxr.Name, r.InPrms[0].Name)
	// test
	if x.f.Test != nil {
		r.T.Addf("expected := (%v + %v) / 2", r.Rxr.Name, r.InPrms[0].Name)
	}
	return r
}
func (x *PrtAri) avgGeo() (r *TypFn) {
	r = x.f.TypFn(k.AvgGeo)
	r.InPrm(x.f.Typ(), "a")
	r.OutPrm(x.f.Typ(), "r")
	r.Addf("r = %v * %v", r.Rxr.Name, r.InPrms[0].Name)
	r.Add("if r == 0 {")
	r.Add("return 0")
	r.Add("} else {")
	r.Add("return r.Sqrt()")
	r.Add("}")
	// test
	if x.f.Test != nil {
		r.T.Addf("expected := %v * %v", r.Rxr.Name, r.InPrms[0].Name)
		r.T.Add("if expected != 0 {")
		r.T.Add("expected = expected.Sqrt()")
		r.T.Add("}")
	}
	return r
}
