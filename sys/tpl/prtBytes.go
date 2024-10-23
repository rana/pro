package tpl

import (
	"sys/k"
)

type (
	PrtBytes struct {
		PrtBse
	}
)

func (x *PrtBytes) InitPrtTypFn() {
	if !x.t.IsBytSkp() {
		x.Bytes()
	}
}
func (x *PrtBytes) Bytes() (r *TypFn) {
	r = x.f.TypFn(k.Bytes)
	r.OutPrmSlice(Byte)
	r.Addf("b := %v{}", BufferPtr.Adr(x.f))
	r.Add("x.BytWrt(b)")
	r.Add("return b.Bytes()")
	return r
}
