package tpl

import "sys/k"

type (
	PrtString struct {
		PrtBse
	}
)

func (x *PrtString) InitPrtTypFn() {
	x.string()
}
func (x *PrtString) string() (r *TypFn) {
	r = x.f.TypFn(k.String)
	r.OutPrm(String)
	r.Addf("b := %v{}", BuilderPtr.Adr(x.f))
	r.Add("x.StrWrt(b)")
	r.Add("return b.String()")
	return r
}
