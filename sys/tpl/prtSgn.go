package tpl

import (
	"sys/k"
)

type (
	PrtSgn struct {
		PrtBse
		Pos *TypFn
		Neg *TypFn
		Inv *TypFn
	}
)

func (x *PrtSgn) InitPrtTypFn() {
	x.Pos = x.pos()
	x.Neg = x.neg()
	x.Inv = x.inv()
}
func (x *PrtSgn) pos() (r *TypFn) {
	r = x.f.TypFn(k.Pos)
	r.OutPrm(x.f.Typ())
	r.Addf("if %v < 0 {", r.Rxr.Name)
	r.Addf("return -%v", r.Rxr.Name)
	r.Add("} else {")
	r.Addf("return %v", r.Rxr.Name)
	r.Add("}")
	// test
	if x.f.Test != nil {
		r.T.Add("expected := cse.x")
		r.T.Add("if expected < 0 {")
		r.T.Add("expected = -expected")
		r.T.Add("}")
	}
	return r
}
func (x *PrtSgn) neg() (r *TypFn) {
	r = x.f.TypFn(k.Neg)
	r.OutPrm(x.f.Typ())
	r.Addf("if %v > 0 {", r.Rxr.Name)
	r.Addf("return -%v", r.Rxr.Name)
	r.Add("} else {")
	r.Addf("return %v", r.Rxr.Name)
	r.Add("}")
	// test
	if x.f.Test != nil {
		r.T.Add("expected := cse.x")
		r.T.Add("if expected > 0 {")
		r.T.Add("expected = -expected")
		r.T.Add("}")
	}
	return r
}
func (x *PrtSgn) inv() (r *TypFn) {
	r = x.f.TypFn(k.Inv)
	r.OutPrm(x.f.Typ())
	r.Addf("return -%v", r.Rxr.Name)
	// test
	if x.f.Test != nil {
		r.T.Addf("expected := -cse.x")
	}
	return r
}

// func (x *PrtSgn) sgn() (r *TypFn) {
// 	r = x.f.TypFn(k.Sgn)
// 	r.InPrm(x.t, "v")
// 	r.OutPrm(x.f.Typ())

// 	r.Addf("return -%v", r.Rxr.Name)
// 	// test
// 	if x.f.Test != nil {
// 		r.T.Addf("expected := -cse.x")
// 	}
// 	return r
// }
