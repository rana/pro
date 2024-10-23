package tpl

import (
	"strings"
	"sys"
	"sys/k"
)

type (
	PrtArrCnt struct {
		PrtBse
	}
)

func (x *PrtArrCnt) InitPrtTypFn() {
	x.op(k.Eql)
	x.op(k.Neq)
	x.op(k.Lss)
	x.op(k.Gtr)
	x.op(k.Leq)
	x.op(k.Geq)
}
func (x *PrtArrCnt) TypFnCnt(name string) (r *TypFn) { return x.f.TypFn(sys.CnjCnt(name)) }
func (x *PrtArrCnt) op(name string) (r *TypFn) {
	r = x.TypFnCnt(name)
	r.InPrm(_sys.Bsc.Flt, "pnt")
	r.OutPrm(_sys.Bsc.Flt, "r")
	r.Add("for _, v := range *x {")
	r.Addf("if v.%v(pnt) {", strings.Title(name))
	r.Add("r++")
	r.Add("}")
	r.Add("}")
	r.Add("return r")
	// test
	if x.f.Test != nil && Opt.IsArr() {
		r.T.SkpCpy = true
		r.T.Addf("var expected %v", r.OutTyp().Ref(x.f.Test))
		r.T.Add("for _, v := range *x {")
		r.T.Addf("if v.%v(pnt) {", strings.Title(name))
		r.T.Add("expected++")
		r.T.Add("}")
		r.T.Add("}")
	}
	return r
}
