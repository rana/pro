package tpl

import (
	"sys/k"
)

type (
	PrtArrStrWrt struct {
		PrtBse
		StrWrt *TypFn
		Ln     bool
	}
)

func (x *PrtArrStrWrt) InitPrtTypFn() {
	if !x.t.IsStrSkp() {
		arr := x.t.PrtArr().Arr
		// if _, isStruct := arr.Alias.Elm.(*Struct); isStruct {
		if arr.Alias.Elm == _sys.Ana.Trd.Typ() {
			x.Ln = true
		}
		x.StrWrt = x.strWrt()
	}
}
func (x *PrtArrStrWrt) strWrt() (r *TypFn) {
	r = x.f.TypFn(k.StrWrt)
	r.InPrm(BuilderPtr, "b")
	r.Add("b.WriteRune('[')")
	if x.Ln {
		r.Add("b.WriteRune('\\n')")
	}
	r.Add("for n, v := range *x {")
	if x.Ln {
		r.Add("v.StrWrt(b)")
		r.Add("if n != len(*x)-1 {")
		r.Add("b.WriteRune('\\n')")
		r.Add("}")
	} else {
		r.Add("if n != 0 {")
		r.Add("b.WriteRune(' ')")
		r.Add("}")
		r.Add("v.StrWrt(b)")
	}
	r.Add("}")
	if x.Ln {
		r.Add("b.WriteRune('\\n')")
	}
	r.Add("b.WriteRune(']')")
	return r
}
