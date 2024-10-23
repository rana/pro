package tpl

import (
	"sys/k"
	"sys/tpl/atr"
)

type (
	PrtRel struct {
		PrtBse
		Cmp *Func
		Eql *PkgFn
		Lss *PkgFn
		Gtr *PkgFn
	}
)

func (x *PrtRel) InitPrtTyp() {
	name := k.Cmp
	if x.t.Pkg.Name != x.t.Camel() {
		name += x.t.Title()
	}
	x.Cmp = x.f.Func(name, atr.None)
}
func (x *PrtRel) InitPrtFld() {
	x.Cmp.InPrm(x.f.Typ(), "a")
	x.Cmp.InPrm(x.f.Typ(), "b")
	x.Cmp.OutPrm(_sys.Bsc.Bol)
}
func (x *PrtRel) InitPrtPkgFn() {
	x.Eql = x.eqlPkg()
	x.Lss = x.lssPkg()
	x.Gtr = x.gtrPkg()
}
func (x *PrtRel) InitPrtTypFn() {
	x.lss()
	x.gtr()
	x.leq()
	x.geq()
	// // tst
	// if x.f.Tst != nil {
	// 	x.f.Tst.RelPrt()
	// }
}

func (x *PrtRel) eqlPkg() (r *PkgFn) {
	r = x.f.PkgFna(k.Eql, atr.None)
	r.Alias = k.Asc
	r.InPrm(x.f.Typ(), "a")
	r.InPrm(x.f.Typ(), "b")
	r.OutPrm(_sys.Bsc.Bol)
	r.Addf("return %v == %v", r.InPrms[0].Name, r.InPrms[1].Name)
	// test
	if x.f.Test != nil {
		x.f.Test.ImportFn(r)
		r.T.Addf("expected := %v(%v == %v)", r.OutTyp().Ref(x.f.Test), r.InPrms[0].Camel(), r.InPrms[1].Camel())
	}
	return r
}
func (x *PrtRel) lssPkg() (r *PkgFn) {
	r = x.f.PkgFna(k.Lss, atr.None)
	r.Alias = k.Asc
	r.InPrm(x.f.Typ(), "a")
	r.InPrm(x.f.Typ(), "b")
	r.OutPrm(_sys.Bsc.Bol)
	r.Addf("return %v < %v", r.InPrms[0].Name, r.InPrms[1].Name)
	// test
	if x.f.Test != nil {
		x.f.Test.ImportFn(r)
		r.T.Addf("expected := %v(%v < %v)", r.OutTyp().Ref(x.f.Test), r.InPrms[0].Camel(), r.InPrms[1].Camel())
	}
	return r
}
func (x *PrtRel) gtrPkg() (r *PkgFn) {
	r = x.f.PkgFna(k.Gtr, atr.None)
	r.Alias = k.Dsc
	r.InPrm(x.f.Typ(), "a")
	r.InPrm(x.f.Typ(), "b")
	r.OutPrm(_sys.Bsc.Bol)
	r.Addf("return %v > %v", r.InPrms[0].Name, r.InPrms[1].Name)
	// test
	if x.f.Test != nil {
		x.f.Test.ImportFn(r)
		r.T.Addf("expected := %v(%v > %v)", r.OutTyp().Ref(x.f.Test), r.InPrms[0].Camel(), r.InPrms[1].Camel())
	}
	return r
}
func (x *PrtRel) lss() (r *TypFn) {
	r = x.f.TypFn(k.Lss)
	r.InPrm(x.f.Typ(), "a")
	r.OutPrm(_sys.Bsc.Bol)
	r.Addf("return %v < %v", r.Rxr.Name, r.InPrms[0].Name)
	// test
	if x.f.Test != nil {
		x.f.Test.ImportFn(r)
		r.T.Addf("expected := %v(%v < %v)", r.OutTyp().Ref(x.f.Test), r.Rxr.Camel(), r.InPrms[0].Camel())
	}
	return r
}
func (x *PrtRel) gtr() (r *TypFn) {
	r = x.f.TypFn(k.Gtr)
	r.InPrm(x.f.Typ(), "a")
	r.OutPrm(_sys.Bsc.Bol)
	r.Addf("return %v > %v", r.Rxr.Name, r.InPrms[0].Name)
	// test
	if x.f.Test != nil {
		x.f.Test.ImportFn(r)
		r.T.Addf("expected := %v(%v > %v)", r.OutTyp().Ref(x.f.Test), r.Rxr.Camel(), r.InPrms[0].Camel())
	}
	return r
}
func (x *PrtRel) leq() (r *TypFn) {
	r = x.f.TypFn(k.Leq)
	r.InPrm(x.f.Typ(), "a")
	r.OutPrm(_sys.Bsc.Bol)
	r.Addf("return %v <= %v", r.Rxr.Name, r.InPrms[0].Name)
	// test
	if x.f.Test != nil {
		x.f.Test.ImportFn(r)
		r.T.Addf("expected := %v(cse.x <= cse.a)", r.OutTyp().Full())
	}
	return r
}
func (x *PrtRel) geq() (r *TypFn) {
	r = x.f.TypFn(k.Geq)
	r.InPrm(x.f.Typ(), "a")
	r.OutPrm(_sys.Bsc.Bol)
	r.Addf("return %v >= %v", r.Rxr.Name, r.InPrms[0].Name)
	// test
	if x.f.Test != nil {
		x.f.Test.ImportFn(r)
		r.T.Addf("expected := %v(cse.x >= cse.a)", r.OutTyp().Full())
	}
	return r
}
