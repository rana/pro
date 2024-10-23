package tpl

type (
	PrtArrSer struct {
		PrtBse
		Arr     *Arr
		ArrTyp  Typ
		ElmTyp  Typ
		AddsLss *PkgFn
		AddsLeq *PkgFn
		SubsGtr *PkgFn
		SubsGeq *PkgFn
		MulsLss *PkgFn
		MulsLeq *PkgFn
		DivsGtr *PkgFn
		DivsGeq *PkgFn
	}
)

func (x *PrtArrSer) InitPrtTyp() {
	x.Arr = x.f.GetPrt((*PrtArr)(nil)).(*PrtArr).Arr
}
func (x *PrtArrSer) InitPrtPkgFn() {
	x.AddsLss = x.addsLss()
	x.AddsLeq = x.addsLeq()
	x.SubsGtr = x.subsGtr()
	x.SubsGeq = x.subsGeq()
	x.MulsLss = x.mulsLss()
	x.MulsLeq = x.mulsLeq()
	x.DivsGtr = x.divsGtr()
	x.DivsGeq = x.divsGeq()
	x.fibsLeq()
}
func (x *PrtArrSer) addsLss() (r *PkgFn) {
	r = x.f.PkgFn("AddsLss")
	r.InPrm(x.Arr.Elm, "strt").LitVal("2")
	r.InPrm(x.Arr.Elm, "lim").LitVal("10")
	r.InPrm(x.Arr.Elm, "by").LitVal("2")
	r.OutPrm(x.f, "r")
	r.Addf("r = %v()", x.Arr.New.Ref(x.f))
	r.Add("for n := strt; n < lim; n += by {")
	r.Add("r.Push(n)")
	r.Add("}")
	r.Add("return r")
	// test
	if x.f.Test != nil && Opt.IsArr() {
		r.T.Addf("expected := %v()", x.Arr.New.Ref(x.f.Test))
		r.T.Add("for n := strt; n < lim; n += by {")
		r.T.Add("expected.Push(n)")
		r.T.Add("}")
	}
	return r
}
func (x *PrtArrSer) addsLeq() (r *PkgFn) {
	r = x.f.PkgFn("AddsLeq")
	r.InPrm(x.Arr.Elm, "strt").LitVal("2")
	r.InPrm(x.Arr.Elm, "lim").LitVal("10")
	r.InPrm(x.Arr.Elm, "by").LitVal("2")
	r.OutPrm(x.f, "r")
	r.Addf("r = %v()", x.Arr.New.Ref(x.f))
	r.Add("for n := strt; n <= lim; n += by {")
	r.Add("r.Push(n)")
	r.Add("}")
	r.Add("return r")
	// test
	if x.f.Test != nil && Opt.IsArr() {
		r.T.Addf("expected := %v()", x.Arr.New.Ref(x.f.Test))
		r.T.Add("for n := strt; n <= lim; n += by {")
		r.T.Add("expected.Push(n)")
		r.T.Add("}")
	}
	return r
}
func (x *PrtArrSer) subsGtr() (r *PkgFn) {
	r = x.f.PkgFn("SubsGtr")
	r.InPrm(x.Arr.Elm, "strt").LitVal("10")
	r.InPrm(x.Arr.Elm, "lim").LitVal("2")
	r.InPrm(x.Arr.Elm, "by").LitVal("2")
	r.OutPrm(x.f, "r")
	r.Addf("r = %v()", x.Arr.New.Ref(x.f))
	r.Add("for n := strt; n > lim; n -= by {")
	r.Add("r.Push(n)")
	r.Add("}")
	r.Add("return r")
	// test
	if x.f.Test != nil && Opt.IsArr() {
		r.T.Addf("expected := %v()", x.Arr.New.Ref(x.f.Test))
		r.T.Add("for n := strt; n > lim; n -= by {")
		r.T.Add("expected.Push(n)")
		r.T.Add("}")
	}
	return r
}
func (x *PrtArrSer) subsGeq() (r *PkgFn) {
	r = x.f.PkgFn("SubsGeq")
	r.InPrm(x.Arr.Elm, "strt").LitVal("10")
	r.InPrm(x.Arr.Elm, "lim").LitVal("2")
	r.InPrm(x.Arr.Elm, "by").LitVal("2")
	r.OutPrm(x.f, "r")
	r.Addf("r = %v()", x.Arr.New.Ref(x.f))
	r.Add("for n := strt; n >= lim; n -= by {")
	r.Add("r.Push(n)")
	r.Add("}")
	r.Add("return r")
	// test
	if x.f.Test != nil && Opt.IsArr() {
		r.T.Addf("expected := %v()", x.Arr.New.Ref(x.f.Test))
		r.T.Add("for n := strt; n >= lim; n -= by {")
		r.T.Add("expected.Push(n)")
		r.T.Add("}")
	}
	return r
}
func (x *PrtArrSer) mulsLss() (r *PkgFn) {
	r = x.f.PkgFn("MulsLss")
	r.InPrm(x.Arr.Elm, "strt").LitVal("2")
	r.InPrm(x.Arr.Elm, "lim").LitVal("10")
	r.InPrm(x.Arr.Elm, "by").LitVal("2")
	r.OutPrm(x.f, "r")
	r.Addf("r = %v()", x.Arr.New.Ref(x.f))
	r.Add("for n := strt; n < lim; n *= by {")
	r.Add("r.Push(n)")
	r.Add("}")
	r.Add("return r")
	// test
	if x.f.Test != nil && Opt.IsArr() {
		r.T.Addf("expected := %v()", x.Arr.New.Ref(x.f.Test))
		r.T.Add("for n := strt; n < lim; n *= by {")
		r.T.Add("expected.Push(n)")
		r.T.Add("}")
	}
	return r
}
func (x *PrtArrSer) mulsLeq() (r *PkgFn) {
	r = x.f.PkgFn("MulsLeq")
	r.InPrm(x.Arr.Elm, "strt").LitVal("2")
	r.InPrm(x.Arr.Elm, "lim").LitVal("10")
	r.InPrm(x.Arr.Elm, "by").LitVal("2")
	r.OutPrm(x.f, "r")
	r.Addf("r = %v()", x.Arr.New.Ref(x.f))
	r.Add("for n := strt; n <= lim; n *= by {")
	r.Add("r.Push(n)")
	r.Add("}")
	r.Add("return r")
	// test
	if x.f.Test != nil && Opt.IsArr() {
		r.T.Addf("expected := %v()", x.Arr.New.Ref(x.f.Test))
		r.T.Add("for n := strt; n <= lim; n *= by {")
		r.T.Add("expected.Push(n)")
		r.T.Add("}")
	}
	return r
}
func (x *PrtArrSer) divsGtr() (r *PkgFn) {
	r = x.f.PkgFn("DivsGtr")
	r.InPrm(x.Arr.Elm, "strt").LitVal("10")
	r.InPrm(x.Arr.Elm, "lim").LitVal("2")
	r.InPrm(x.Arr.Elm, "by").LitVal("2")
	r.OutPrm(x.f, "r")
	r.Addf("r = %v()", x.Arr.New.Ref(x.f))
	r.Add("for n := strt; n > lim; n /= by {")
	r.Add("r.Push(n)")
	r.Add("}")
	r.Add("return r")
	// test
	if x.f.Test != nil && Opt.IsArr() {
		r.T.Addf("expected := %v()", x.Arr.New.Ref(x.f.Test))
		r.T.Add("for n := strt; n > lim; n /= by {")
		r.T.Add("expected.Push(n)")
		r.T.Add("}")
	}
	return r
}
func (x *PrtArrSer) divsGeq() (r *PkgFn) {
	r = x.f.PkgFn("DivsGeq")
	r.InPrm(x.Arr.Elm, "strt").LitVal("10")
	r.InPrm(x.Arr.Elm, "lim").LitVal("2")
	r.InPrm(x.Arr.Elm, "by").LitVal("2")
	r.OutPrm(x.f, "r")
	r.Addf("r = %v()", x.Arr.New.Ref(x.f))
	r.Add("for n := strt; n >= lim; n /= by {")
	r.Add("r.Push(n)")
	r.Add("}")
	r.Add("return r")
	// test
	if x.f.Test != nil && Opt.IsArr() {
		r.T.Addf("expected := %v()", x.Arr.New.Ref(x.f.Test))
		r.T.Add("for n := strt; n >= lim; n /= by {")
		r.T.Add("expected.Push(n)")
		r.T.Add("}")
	}
	return r
}
func (x *PrtArrSer) fibsLeq() (r *PkgFn) {
	r = x.f.PkgFn("FibsLeq")
	r.InPrm(x.Arr.Elm, "lim").LitVal("610")
	r.OutPrm(x.f, "r")
	r.Addf("r = %v(1, 2)", x.Arr.New.Ref(x.f))
	r.Add("for (*r)[len(*r)-1] < lim {")
	r.Add("*r = append(*r, (*r)[len(*r)-2] + (*r)[len(*r)-1])")
	r.Add("}")
	r.Add("return r")
	// test
	if x.f.Test != nil && Opt.IsArr() {
		r.T.Addf("expected := %v()", x.Arr.New.Ref(x.f.Test))
		r.T.Add("for (*expected)[len(*expected)-1] < lim {")
		r.T.Add("*expected = append(*rexpected, (*expected)[len(*expected)-2] + (*expected)[len(*expected)-1])")
		r.T.Add("}")
	}
	return r
}
