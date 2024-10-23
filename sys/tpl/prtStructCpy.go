package tpl

import (
	"sys/err"
	"sys/k"
)

type (
	PrtStructCpy struct {
		PrtBse
		S   *Struct
		Cpy *TypFn
	}
)

func (x *PrtStructCpy) InitPrtTyp() {
	x.S = x.f.Typ().(*Struct)
	if x.S == nil {
		err.Panicf("'%v' is not *Struct", x.f.Typ().Full())
	}
}
func (x *PrtStructCpy) InitPrtTypFn() {
	x.Cpy = x.cpy()
}
func (x *PrtStructCpy) cpy() (r *TypFn) {
	r = x.f.TypFn(k.Cpy)
	r.OutPrm(x.f, "r")
	if len(x.S.Flds) == 0 {
		r.Add("panic(\"no flds\")")
		r.T.Empty = true
	} else {
		r.Addf("r = %v{}", r.OutTyp().Adr(x.f))
		for _, fld := range x.S.Flds {
			if !fld.IsStrSkp() { // TODO: ADD CPY ATR?
				r.Addf("r.%[1]v = x.%[1]v", fld.Name)
			}
		}
	}
	r.Add("return r")
	return r
}
