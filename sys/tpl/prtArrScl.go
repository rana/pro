package tpl

import (
	"strings"
	"sys"
	"sys/ks"
)

type (
	PrtArrScl struct {
		PrtBse
		Arr *Arr
	}
)

func (x *PrtArrScl) InitPrtTyp() {
	x.Arr = x.f.GetPrt((*PrtArr)(nil)).(*PrtArr).Arr
}
func (x *PrtArrScl) InitPrtTypFn() {
	for _, scl := range ks.Scls {
		x.op(scl)
	}
}
func (x *PrtArrScl) op(name string) (r *TypFn) {
	name = strings.Title(name)
	r = x.f.TypFn(sys.CnjScl(name))
	r.InPrm(_sys.Bsc.Flt, "scl")
	r.OutPrm(x.f.Typ())
	r.Addf("r := make(%v, len(*x))", x.t.Title())
	r.Add("for n, v := range *x {")
	r.Addf("r[n] = v.%v(scl)", name) // TODO: INLINE
	r.Add("}")
	r.Add("return &r")
	// test
	if x.f.Test != nil && Opt.IsArr() {
		r.T.SkpCpy = true
		r.T.Addf("expected :=  %v(x.Cnt())", x.Arr.Make.Ref(x.f.Test))
		r.T.Add("for n, v := range *x {")
		r.T.Addf("expected.Upd(%v(n), v.%v(scl))", _sys.Bsc.Unt.Typ().Ref(x.f.Test), name)
		r.T.Add("}")
	}
	return r
}
