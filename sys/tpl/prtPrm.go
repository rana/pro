package tpl

import "sys/k"

type (
	PrtPrm struct {
		PrtBse
	}
)

func (x *PrtPrm) InitPrtTypFn() {
	x.string()
}
func (x *PrtPrm) string() (r *TypFn) {
	r = x.f.TypFn(k.Prm)
	r.OutPrm(String)
	r.Addf("b := %v{}", BuilderPtr.Adr(x.f))
	r.Add("x.PrmWrt(b)")
	r.Add("return b.String()")
	return r
}
