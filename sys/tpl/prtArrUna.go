package tpl

import (
	"strings"
	"sys"
	"sys/ks"
)

type (
	PrtArrUna struct {
		PrtBse
		Arr *Arr
	}
)

func (x *PrtArrUna) InitPrtTyp() {
	x.Arr = x.f.GetPrt((*PrtArr)(nil)).(*PrtArr).Arr
}
func (x *PrtArrUna) InitPrtTypFn() {
	for _, una := range ks.Unas {
		x.op(una)
	}
}
func (x *PrtArrUna) op(name string) (r *TypFn) {
	name = strings.Title(name)
	r = x.f.TypFn(sys.CnjUna(name))
	r.OutPrm(x.f.Typ())
	r.Addf("r := make(%v, len(*x))", x.t.Title())
	r.Add("for n, v := range *x {")
	r.Addf("r[n] = v.%v()", name)
	r.Add("}")
	r.Add("return &r")
	// test
	if x.f.Test != nil && Opt.IsArr() {
		r.T.SkpCpy = true
		r.T.Addf("expected :=  %v(x.Cnt())", x.Arr.Make.Ref(x.f.Test))
		r.T.Add("for n, v := range *x {")
		r.T.Addf("expected.Upd(%v(n), v.%v())", _sys.Bsc.Unt.Typ().Ref(x.f.Test), name)
		r.T.Add("}")
	}
	return r
}
