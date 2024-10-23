package tpl

import (
	"sys"
	"sys/k"
)

type (
	PrtArrIdn struct {
		PrtBse
		Arr *Arr
	}
)

func (x *PrtArrIdn) InitPrtTyp() {
	x.Arr = x.f.GetPrt((*PrtArr)(nil)).(*PrtArr).Arr
}
func (x *PrtArrIdn) InitPrtTypFn() {
	if !x.t.IsEqlSkp() {
		// // tst
		// if x.f.Tst != nil {
		// 	x.f.Tst.IdnPrt()
		// }
		x.eql()
		x.neq()
	}
}

func (x *PrtArrIdn) eql() (r *TypFn) {
	r = x.f.TypFn(k.Eql)
	r.InPrm(x.f.Typ(), "a")
	r.OutPrm(_sys.Bsc.Bol.Typ())
	r.Add("if a == nil {")
	r.Add("return false")
	r.Add("}")
	r.Add("if len(*x) != len(*a) {")
	r.Add("return false")
	r.Add("}")
	sys.Log(" -", x.t.Name)
	_, isIfc := x.Arr.Alias.Elm.(*Ifc)
	if isIfc {
		r.Add("for n := unt.Zero; n < unt.Unt(len(*x)); n++ {")
		r.Add("if x.At(n) != nil && a.At(n) != nil && x.At(n).Neq(a.At(n)) {")
		r.Add("return false")
		r.Add("} else if x.At(n) == nil && a.At(n) == nil {")
		r.Add("continue")
		r.Add("} else if x.At(n) == nil || a.At(n) == nil {")
		r.Add("return false")
		r.Add("}")
		r.Add("}")
	} else {
		r.Add("for n := unt.Zero; n < unt.Unt(len(*x)); n++ {")
		r.Add("if x.At(n).Neq(a.At(n)) {")
		r.Add("return false")
		r.Add("}")
		r.Add("}")
	}
	r.Add("return true")
	// test
	if x.f.Test != nil && Opt.IsArr() {
		x.f.Test.ImportFn(r)
		r.T.Add("expected := bol.Tru")
		r.T.Add("if a == nil {")
		r.T.Add("expected = false")
		r.T.Add("} else if exCpy.Cnt() != a.Cnt() {")
		r.T.Add("expected = false")
		r.T.Add("} else {")
		r.T.Add("for n := unt.Zero; n < exCpy.Cnt(); n++ {")
		r.T.Add("if exCpy.At(n).Neq(a.At(n)) {")
		r.T.Add("expected = false")
		r.T.Add("}")
		r.T.Add("}")
		r.T.Add("}")
	}
	return r
}
func (x *PrtArrIdn) neq() (r *TypFn) {
	r = x.f.TypFn(k.Neq)
	r.InPrm(x.f.Typ(), "a")
	r.OutPrm(_sys.Bsc.Bol.Typ())
	r.Add("return !x.Eql(a)")
	// test
	if x.f.Test != nil && Opt.IsArr() {
		r.T.Add("expected := bol.Fls")
		r.T.Add("if a == nil {")
		r.T.Add("expected = true")
		r.T.Add("} else if exCpy.Cnt() != a.Cnt() {")
		r.T.Add("expected = true")
		r.T.Add("} else {")
		r.T.Add("for n := unt.Zero; n < exCpy.Cnt(); n++ {")
		r.T.Add("if exCpy.At(n).Neq(a.At(n)) {")
		r.T.Add("expected = true")
		r.T.Add("}")
		r.T.Add("}")
		r.T.Add("}")
	}
	return r
}
