package tpl

import (
	"fmt"
	"strings"
	"sys/k"
	"sys/tpl/atr"
)

type (
	PrtEnum struct {
		PrtBse
	}
)

func (x *PrtEnum) InitPrtCnst() {
	if len(x.f.Cnsts) != 0 {
		c := x.f.Cnstf("No%v", x.t.Title())
		c.Atr = atr.Iota
		x.f.Cnsts = x.f.Cnsts[:len(x.f.Cnsts)-1] // cnst was added at end
		x.f.Cnsts = append([]*Cnst{c}, x.f.Cnsts...)
	}
}
func (x *PrtEnum) InitPrtVar() {
	// x.Var("StkWidth", "1", _sys.Bsc.Unt)
	// x.Var("Red50", "0xff, 0xeb, 0xee, 0xff")
	x.TypName()
}
func (x *PrtEnum) TypName() {
	var b strings.Builder
	b.WriteString("map[")
	b.WriteString(x.t.Title())
	b.WriteString("]string{\n")
	for _, c := range x.f.Cnsts {
		b.WriteString(c.Name)
		b.WriteString(": \"")
		b.WriteString(c.Camel())
		b.WriteString("\",\n")
	}
	b.WriteString("}")
	v := x.f.Var(x.TypNameVarName(), b.String())
	v.Name = v.Camel()
	v.FnCall = true
}
func (x *PrtEnum) TypNameVarName() string {
	return fmt.Sprintf("%vNames", x.t.Camel())
}
func (x *PrtEnum) InitPrtTypFn() {
	x.string()
}
func (x *PrtEnum) string() (r *TypFn) {
	r = x.f.TypFn(k.String)
	r.OutPrm(String)
	r.Addf("return %v[x]", x.TypNameVarName())
	return r
}
