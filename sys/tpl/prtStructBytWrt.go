package tpl

import (
	"sys/err"
	"sys/k"
)

type (
	PrtStructBytWrt struct {
		PrtBse
		S      *Struct
		BytWrt *TypFn
	}
)

func (x *PrtStructBytWrt) InitPrtTyp() {
	x.S = x.f.Typ().(*Struct)
	if x.S == nil {
		err.Panicf("'%v' is not *Struct", x.f.Typ().Full())
	}
}
func (x *PrtStructBytWrt) InitPrtTypFn() {
	x.BytWrt = x.bytWrt()
}
func (x *PrtStructBytWrt) bytWrt() (r *TypFn) {
	r = x.f.TypFn(k.BytWrt)
	r.InPrm(BufferPtr, "b")
	if len(x.S.Flds) == 0 {
		r.Add("panic(\"no flds\")")
		r.T.Empty = true
	} else {
		for _, fld := range x.S.Flds {
			if !fld.IsBytSkp() && !fld.IsKey() {
				if fld.Name == "" { // embedded type
					if fld.IsEmbeddedBytWrt() {
						r.Addf("x.%v.BytWrt(b)", fld.Typ.Title())
					}
				} else {
					r.Addf("x.%v.BytWrt(b)", fld.Name)
				}
			}
		}
	}
	return r
}
