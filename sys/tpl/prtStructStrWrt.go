package tpl

import (
	"sys/err"
	"sys/k"
)

type (
	PrtStructStrWrt struct {
		PrtBse
		S      *Struct
		StrWrt *TypFn
	}
)

func (x *PrtStructStrWrt) InitPrtTyp() {
	x.S = x.f.Typ().(*Struct)
	if x.S == nil {
		err.Panicf("'%v' is not *Struct", x.f.Typ().Full())
	}
}
func (x *PrtStructStrWrt) InitPrtTypFn() {
	x.StrWrt = x.strWrt()
}
func (x *PrtStructStrWrt) strWrt() (r *TypFn) {
	r = x.f.TypFn(k.StrWrt)
	r.InPrm(BuilderPtr, "b")
	r.OutPrm(String)
	if len(x.S.Flds) == 0 {
		r.Add("panic(\"no flds\")")
		r.T.Empty = true
	} else {
		r.Addf("b.WriteString(\"%v.%v(\")", x.t.Pkg.Name, x.t.Camel())
		prefix := ""
		for _, fld := range x.S.Flds {
			if !fld.IsStrSkp() {
				r.Addf("b.WriteString(\"%v%v:\")", prefix, fld.Camel())
				r.Addf("x.%v.StrWrt(b)", fld.Name)
				if prefix == "" {
					prefix = " "
				}
			}
		}
	}
	r.Add("b.WriteRune(')')")
	r.Add("return b.String()")
	return r
}
