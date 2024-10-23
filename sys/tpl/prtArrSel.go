package tpl

import (
	"sys/k"
)

type (
	PrtArrSel struct {
		PrtBse
		Arr *Arr
	}
)

func (x *PrtArrSel) InitPrtTyp() {
	x.Arr = x.f.GetPrt((*PrtArr)(nil)).(*PrtArr).Arr
}
func (x *PrtArrSel) InitPrtTypFn() {
	x.op(_sys.Bsc.Flt.PrtSel.SelEql.Name)
	x.op(_sys.Bsc.Flt.PrtSel.SelNeq.Name)
	x.op(_sys.Bsc.Flt.PrtSel.SelLss.Name)
	x.op(_sys.Bsc.Flt.PrtSel.SelGtr.Name)
	x.op(_sys.Bsc.Flt.PrtSel.SelLeq.Name)
	x.op(_sys.Bsc.Flt.PrtSel.SelGeq.Name)
	x.Splt()
}
func (x *PrtArrSel) op(name string) (r *TypFn) {
	r = x.f.TypFn(name)
	r.InPrm(_sys.Bsc.Flt, "sel")
	r.OutPrm(x.f.Typ(), "r")
	r.Addf("r = %v(x.Cnt())", x.Arr.Make.Ref(x.f))
	r.Add("for n, v := range *x {")
	// r.Addf("r.Upd(%v(n), v.%v(sel))", _sys.Bsc.Unt.Typ().Ref(x.f), name)
	r.Addf("(*r)[n] = v.%v(sel)", name)
	r.Add("}")
	r.Add("return r")
	// test
	if x.f.Test != nil && Opt.IsArr() {
		r.T.SkpCpy = true
		r.T.Addf("expected :=  %v(x.Cnt())", x.Arr.Make.Ref(x.f.Test))
		r.T.Add("for n, v := range *x {")
		r.T.Addf("expected.Upd(%v(n), v.%v(sel))", _sys.Bsc.Unt.Typ().Ref(x.f.Test), name)
		r.T.Add("}")
	}
	return r
}
func (x *PrtArrSel) Splt() (r *TypFn) {
	r = x.f.TypFn(k.Splt) // Splt(0.0) (btm, top *Flts)
	r.InPrm(x.Arr.Alias.Elm, "v")
	r.OutPrm(x.t, "btm")
	r.OutPrm(x.t, "top")
	r.Addf("btm = %v()", x.Arr.New.Ref(x.f))
	r.Addf("top = %v()", x.Arr.New.Ref(x.f))
	r.Add("for _, cur := range *x {")
	r.Add("if cur > v {")
	r.Add("*top = append(*top, cur)")
	r.Add("} else {")
	r.Add("*btm = append(*btm, cur)")
	r.Add("}")
	r.Add("}")
	r.Add("return btm, top")
	return r
}
