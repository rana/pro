package tpl

import (
	"sys/err"
	"sys/k"
)

type (
	PrtStructKeyWrt struct {
		PrtBse
		S      *Struct
		KeyWrt *TypFn
		Key    *TypFn
	}
)

func (x *PrtStructKeyWrt) InitPrtTyp() {
	x.S = x.f.Typ().(*Struct)
	if x.S == nil {
		err.Panicf("'%v' is not *Struct", x.f.Typ().Full())
	}
}
func (x *PrtStructKeyWrt) InitPrtTypFn() {
	x.KeyWrt = x.keyWrt()
	x.Key = x.key()
}
func (x *PrtStructKeyWrt) keyWrt() (r *TypFn) {
	r = x.f.TypFn(k.KeyWrt)
	r.InPrm(BufferPtr, "b")
	if len(x.S.Flds) == 0 {
		r.Add("panic(\"no flds\")")
		r.T.Empty = true
	} else {
		for _, fld := range x.S.Flds {
			if !fld.IsBytSkp() && fld.IsKey() {
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
func (x *PrtStructKeyWrt) key() (r *TypFn) {
	r = x.f.TypFn(k.Key)
	r.OutPrmSlice(Byte)
	r.Addf("b := %v{}", BufferPtr.Adr(x.f))
	r.Add("x.KeyWrt(b)")
	r.Add("return b.Bytes()")
	return r
}
