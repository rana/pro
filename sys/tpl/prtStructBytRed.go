package tpl

import (
	"sys/k"
)

type (
	PrtStructBytRed struct {
		PrtBse
		S *Struct
	}
)

func (x *PrtStructBytRed) InitPrtTyp() {
	x.S = x.f.Typ().(*Struct)
}
func (x *PrtStructBytRed) InitPrtTypFn() {
	x.bytRed()
}
func (x *PrtStructBytRed) bytRed() (r *TypFn) {
	r = x.f.TypFn(k.BytRed)
	r.InPrmSlice(Byte, "b")
	r.OutPrm(Int, "idx")
	if len(x.S.Flds) == 0 {
		r.Add("panic(\"no flds\")")
		r.T.Empty = true
	} else {
		var suffix string
		for _, fld := range x.S.Flds {
			if arr, ok := fld.Typ.(*Arr); ok && !fld.IsBytSkp() {
				r.Addf("x.%v = %v()", fld.Name, arr.New.Ref(x.f))
			}
		}
		for _, fld := range x.S.Flds {
			if !fld.IsBytSkp() && !fld.IsKey() {
				if fld.Name == "" { // embedded type
					if fld.IsEmbeddedBytWrt() {
						r.Addf("idx += x.%v.BytRed(b%v)", fld.Typ.Title(), suffix)
					}
				} else {
					r.Addf("idx += x.%v.BytRed(b%v)", fld.Name, suffix)
				}
				if suffix == "" {
					suffix = "[idx:]"
				}
			}
		}
	}
	r.Add("return idx")
	return r
}
