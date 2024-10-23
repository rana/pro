package tpl

import "sys/k"

type (
	PrtPth struct {
		PrtBse
	}
)

func (x *PrtPth) InitPrtTypFn() {
	x.Pth()
}
func (x *PrtPth) Pth() (r *TypFn) {
	r = x.f.TypFn(k.Pth)
	r.OutPrm(String)
	r.Addf("b := %v{}", BuilderPtr.Adr(x.f))
	r.Add("x.PthWrt(b)")
	r.Add("return b.String()")
	return r
}
