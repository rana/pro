package tpl

import (
	"sys/k"
)

type (
	PrtIdn struct {
		PrtBse
		Eql *TypFn
		Neq *TypFn
	}
)

func (x *PrtIdn) InitPrtTypFn() {
	// // tst
	// if x.f.Tst != nil {
	// 	x.f.Tst.IdnPrt()
	// }
	x.Eql = x.eql()
	x.Neq = x.neq()
}
func (x *PrtIdn) eql() (r *TypFn) {
	r = x.f.TypFn(k.Eql)
	r.InPrm(x.f.Typ(), "a")
	r.OutPrm(_sys.Bsc.Bol)
	r.Addf("return %v == %v", r.Rxr.Name, r.InPrms[0].Name)
	// test
	if x.f.Test != nil {
		x.f.Test.ImportFn(r)
		r.T.Addf("expected := %v(%v == %v)", r.OutTyp().Ref(x.f.Test), r.Rxr.Camel(), r.InPrms[0].Camel())
	}
	return r
}
func (x *PrtIdn) neq() (r *TypFn) {
	r = x.f.TypFn(k.Neq)
	r.InPrm(x.f.Typ(), "a")
	r.OutPrm(_sys.Bsc.Bol)
	r.Addf("return %v != %v", r.Rxr.Name, r.InPrms[0].Name)
	// test
	if x.f.Test != nil {
		x.f.Test.ImportFn(r)
		r.T.Addf("expected := %v(%v != %v)", r.OutTyp().Ref(x.f.Test), r.Rxr.Camel(), r.InPrms[0].Camel())
	}
	return r
}
