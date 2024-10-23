package tpl

import (
	"sys"
	"sys/k"
)

type (
	PrtSel struct {
		PrtBse
		SelEql *TypFn
		SelNeq *TypFn
		SelLss *TypFn
		SelGtr *TypFn
		SelLeq *TypFn
		SelGeq *TypFn
	}
)

func (x *PrtSel) InitPrtTypFn() {
	x.SelEql = x.eql()
	x.SelNeq = x.neq()
	x.SelLss = x.lss()
	x.SelGtr = x.gtr()
	x.SelLeq = x.leq()
	x.SelGeq = x.geq()
}
func (x *PrtSel) TypFnSel(name string) (r *TypFn) { return x.f.TypFn(sys.CnjSel(name)) }
func (x *PrtSel) eql() (r *TypFn) {
	r = x.TypFnSel(k.Eql)
	r.InPrm(_sys.Bsc.Flt, "a")
	r.OutPrm(x.f.Typ())
	r.Addf("if %v == %v {", r.Rxr.Name, r.InPrms[0].Name)
	r.Addf("return %v", r.Rxr.Name)
	r.Add("} else {")
	r.Add("return 0")
	r.Add("}")
	// test
	if x.f.Test != nil {
		r.T.Addf("var expected %v", r.OutTyp().Full())
		r.T.Addf("if %v == %v {", r.Rxr.Name, r.InPrms[0].Name)
		r.T.Addf("expected = %v", r.Rxr.Name)
		r.T.Add("}")
	}
	return r
}
func (x *PrtSel) neq() (r *TypFn) {
	r = x.TypFnSel(k.Neq)
	r.InPrm(_sys.Bsc.Flt, "a")
	r.OutPrm(x.f.Typ())
	r.Addf("if %v != %v {", r.Rxr.Name, r.InPrms[0].Name)
	r.Addf("return %v", r.Rxr.Name)
	r.Add("} else {")
	r.Add("return 0")
	r.Add("}")
	// test
	if x.f.Test != nil {
		r.T.Addf("var expected %v", r.OutTyp().Full())
		r.T.Addf("if %v != %v {", r.Rxr.Name, r.InPrms[0].Name)
		r.T.Addf("expected = %v", r.Rxr.Name)
		r.T.Add("}")
	}
	return r
}
func (x *PrtSel) lss() (r *TypFn) {
	r = x.TypFnSel(k.Lss)
	r.InPrm(_sys.Bsc.Flt, "a")
	r.OutPrm(x.f.Typ())
	r.Addf("if %v < %v {", r.Rxr.Name, r.InPrms[0].Name)
	r.Addf("return %v", r.Rxr.Name)
	r.Add("} else {")
	r.Add("return 0")
	r.Add("}")
	// test
	if x.f.Test != nil {
		r.T.Addf("var expected %v", r.OutTyp().Full())
		r.T.Addf("if %v < %v {", r.Rxr.Name, r.InPrms[0].Name)
		r.T.Addf("expected = %v", r.Rxr.Name)
		r.T.Add("}")
	}
	return r
}
func (x *PrtSel) gtr() (r *TypFn) {
	r = x.TypFnSel(k.Gtr)
	r.InPrm(_sys.Bsc.Flt, "a")
	r.OutPrm(x.f.Typ())
	r.Addf("if %v > %v {", r.Rxr.Name, r.InPrms[0].Name)
	r.Addf("return %v", r.Rxr.Name)
	r.Add("} else {")
	r.Add("return 0")
	r.Add("}")
	// test
	if x.f.Test != nil {
		r.T.Addf("var expected %v", r.OutTyp().Full())
		r.T.Addf("if %v > %v {", r.Rxr.Name, r.InPrms[0].Name)
		r.T.Addf("expected = %v", r.Rxr.Name)
		r.T.Add("}")
	}
	return r
}
func (x *PrtSel) leq() (r *TypFn) {
	r = x.TypFnSel(k.Leq)
	r.InPrm(_sys.Bsc.Flt, "a")
	r.OutPrm(x.f.Typ())
	r.Addf("if %v <= %v {", r.Rxr.Name, r.InPrms[0].Name)
	r.Addf("return %v", r.Rxr.Name)
	r.Add("} else {")
	r.Add("return 0")
	r.Add("}")
	// test
	if x.f.Test != nil {
		r.T.Addf("var expected %v", r.OutTyp().Full())
		r.T.Addf("if %v <= %v {", r.Rxr.Name, r.InPrms[0].Name)
		r.T.Addf("expected = %v", r.Rxr.Name)
		r.T.Add("}")
	}
	return r
}
func (x *PrtSel) geq() (r *TypFn) {
	r = x.TypFnSel(k.Geq)
	r.InPrm(_sys.Bsc.Flt, "a")
	r.OutPrm(x.f.Typ())
	r.Addf("if %v >= %v {", r.Rxr.Name, r.InPrms[0].Name)
	r.Addf("return %v", r.Rxr.Name)
	r.Add("} else {")
	r.Add("return 0")
	r.Add("}")
	// test
	if x.f.Test != nil {
		r.T.Addf("var expected %v", r.OutTyp().Full())
		r.T.Addf("if %v >= %v {", r.Rxr.Name, r.InPrms[0].Name)
		r.T.Addf("expected = %v", r.Rxr.Name)
		r.T.Add("}")
	}
	return r
}
