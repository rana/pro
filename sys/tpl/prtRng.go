package tpl

import (
	"sys/k"
	"sys/tpl/atr"
	"sys/tpl/mod"
)

type (
	PrtRng struct {
		PrtBse
		Elm     *FleBse
		New     *PkgFn
		NewArnd *PkgFn
		NewFul  *PkgFn
	}
)

// func NewRng(elm *FleBse, pkg ...*Pkg) (r *FleBse) {
// 	elmBse := elm.Typ().Bse()
// 	r = &FleBse{}
// 	r.Name = k.Rng
// 	r.Pkg = elmBse.Pkg
// 	elmBse.Rng = r.Structf(r.Name, elmBse.Atr|atr.Rng)
// 	elmBse.Rng.Atr = elmBse.Rng.Atr &^ atr.BytWrt
// 	elmBse.Rng.Atr |= atr.XprActTestSkp | atr.Test | atr.Tst
// 	elmBse.Rng.Fld("Min", elmBse)
// 	elmBse.Rng.Fld("Max", elmBse)
// 	return r
// }
func (x *PrtRng) InitPrtTyp() {
	// x.t.Atr |= atr.Rng | atr.Lit | atr.Test | atr.Tst | atr.XprActTestSkp
	// x.t.Atr = x.t.Atr &^ atr.Bsc
	// x.t.Atr = x.t.Atr &^ atr.BytWrt
}
func (x *PrtRng) InitPrtFld() {
	s := x.f.Typ().(*Struct)
	s.Fld("Min", x.Elm)
	s.Fld("Max", x.Elm)
}
func (x *PrtRng) InitPrtPkgFn() {
	x.New = x.new()
	x.NewArnd = x.newArnd()
	x.NewFul = x.newFul()
}
func (x *PrtRng) new() (r *PkgFn) {
	r = x.f.PkgFn(k.New, true)
	r.InPrm(x.Elm, "min")
	r.InPrm(x.Elm, "max")
	r.OutPrm(x.f.Typ(), "r")
	r.Atr = atr.Lng
	r.Add("r.Min = min")
	r.Add("r.Max = max")
	r.Add("return r")
	return r
}
func (x *PrtRng) newArnd() (r *PkgFn) {
	r = x.f.PkgFnf("New%vArnd", x.t.Title())
	r.InPrm(x.Elm, "cntr")
	r.InPrm(x.Elm, "radius")
	r.OutPrm(x.f.Typ(), "r")
	r.Atr = atr.Lng
	r.Add("r.Min = cntr - radius")
	r.Add("r.Max = cntr + radius")
	r.Add("return r")
	return r
}
func (x *PrtRng) newFul() (r *PkgFn) {
	r = x.f.PkgFnf("New%vFul", x.t.Title())
	r.OutPrm(x.f.Typ(), "r")
	r.Atr = atr.Lng
	r.Add("r.Min = Min")
	r.Add("r.Max = Max")
	r.Add("return r")
	return r
}
func (x *PrtRng) InitPrtTypFn() {
	// // tst
	// if x.f.Tst != nil {
	// 	x.f.Tst.RngPrt()
	// }
	x.Len()
	x.IsValid()
	x.Ensure()
	x.MinSub()
	x.MaxAdd()
	x.Mrg()
	// x.By()
	x.StrWrt()
	x.BytWrt()
	x.BytRed()
}
func (x *PrtRng) Len() (r *TypFn) {
	r = x.f.TypFn(k.Len)
	r.OutPrm(x.Elm)
	r.Add("return x.Max - x.Min")
	return r
}
func (x *PrtRng) IsValid() (r *TypFn) {
	r = x.f.TypFn(k.IsValid)
	r.OutPrm(_sys.Bsc.Bol)
	r.Add("return x.Min < x.Max")
	return r
}
func (x *PrtRng) Ensure() (r *TypFn) {
	r = x.f.TypFn(k.Ensure)
	r.OutPrm(x.t)
	r.Add("if x.Min > x.Max {")
	r.Add("x.Min, x.Max = x.Max, x.Min // swp")
	r.Add("}")
	r.Add("return x")
	return r
}
func (x *PrtRng) MinSub() (r *TypFn) {
	r = x.f.TypFn("MinSub")
	r.InPrm(x.Elm, "v")
	r.OutPrm(x.t)
	r.Add("x.Min -= v")
	r.Add("return x")
	return r
}
func (x *PrtRng) MaxAdd() (r *TypFn) {
	r = x.f.TypFn("MaxAdd")
	r.InPrm(x.Elm, "v")
	r.OutPrm(x.t)
	r.Add("x.Max += v")
	r.Add("return x")
	return r
}
func (x *PrtRng) Mrg() (r *TypFn) {
	r = x.f.TypFn(k.Mrg)
	r.InPrm(x.t, "v")
	r.OutPrm(x.t)
	r.Add("x.Min = x.Min.Min(v.Min)")
	r.Add("x.Max = x.Max.Max(v.Max)")
	r.Add("return x")
	return r
}

// func (x *PrtRng) By() (r *TypFn) {
// 	r = x.f.TypFn("By")
// 	r.InPrm(x.Elm, "inc")
// 	r.OutPrm(x.Elm.Arr, "r")
// 	r.Addf("r = %v()", x.Elm.Arr.New.Ref(x.f))
// 	r.Add("for v := x.Min; v <= x.Max; v += inc {")
// 	r.Add("*r = append(*r, v)")
// 	r.Add("}")
// 	r.Add("return r")
// 	return r
// }
func (x *PrtRng) StrWrt() (r *TypFn) {
	r = x.f.TypFn(k.StrWrt)
	r.InPrm(BuilderPtr, "b")
	r.Add("x.Min.StrWrt(b)")
	r.Add("b.WriteRune('-')")
	r.Add("x.Max.StrWrt(b)")
	return r
}
func (x *PrtRng) BytWrt() (r *TypFn) {
	r = x.f.TypFn(k.BytWrt)
	r.InPrm(BufferPtr, "b")
	r.Add("x.Min.BytWrt(b)")
	r.Add("x.Max.BytWrt(b)")
	return r
}
func (x *PrtRng) BytRed() (r *TypFn) {
	r = x.f.TypFn(k.BytRed)
	r.Rxr.Mod = mod.Ptr
	r.InPrmSlice(Byte, "b")
	r.OutPrm(Int)
	r.Add("idx := 0")
	r.Add("idx += x.Min.BytRed(b[idx:])")
	r.Add("x.Max.BytRed(b[idx:])")
	r.Add("return Size")
	return r
}
