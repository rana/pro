package tpl

import (
	"fmt"
	"strings"
	"sys/k"
	"sys/tpl/atr"
)

type (
	PrtEnumFlg struct {
		PrtBse
	}
)

func (x *PrtEnumFlg) InitPrtCnst() {
	if len(x.f.Cnsts) != 0 {
		// c := x.f.Cnstf("No%v", x.t.Title())
		// c.Atr = atr.Iota
		// x.f.Cnsts = x.f.Cnsts[:len(x.f.Cnsts)-1] // cnst was added at end
		// x.f.Cnsts = append([]*Cnst{c}, x.f.Cnsts...)
		x.f.Cnsts[0].Atr = atr.Iota | atr.Flg
		c := x.f.Cnstf("No%v", x.t.Title())
		c.Value = "0"
	}
}
func (x *PrtEnumFlg) InitPrtVar() {
	// x.Var("StkWidth", "1", _sys.Bsc.Unt)
	// x.Var("Red50", "0xff, 0xeb, 0xee, 0xff")
	x.TypName()
	// x.NameTyp()
}

func (x *PrtEnumFlg) TypName() {
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
func (x *PrtEnumFlg) TypNameVarName() string {
	return fmt.Sprintf("%vNames", x.t.Camel())
}

// func (x *PrtEnumFlg) NameTyp() {
// 	var b strings.Builder
// 	b.WriteString("map[string]")
// 	b.WriteString(x.t.Title())
// 	b.WriteString("{\n")
// 	for _, c := range x.f.Cnsts {
// 		b.WriteString("\"")
// 		b.WriteString(c.Camel())
// 		b.WriteString("\":")
// 		b.WriteString(c.Name)
// 		b.WriteString(",\n")
// 	}
// 	b.WriteString("}")
// 	v := x.f.Var(x.NameTypVarName(), b.String())
// 	v.Name = v.Camel()
// 	v.FnCall = true
// }
// func (x *PrtEnumFlg) NameTypVarName() string {
// 	return fmt.Sprintf("name%vs", x.t.Title())
// }

func (x *PrtEnumFlg) InitPrtTypFn() {
	x.prs()
	x.string()
	x.is()
	for _, c := range x.f.Cnsts {
		x.isFlg(c)
	}
}

func (x *PrtEnumFlg) prs() (r *PkgFn) {
	x.f.Import("strings")
	r = x.f.PkgFnf("Prs%v", x.t.Title())
	r.InPrm(String, "txt")
	r.OutPrm(x.t, "r")
	r.Add("txtVals := strings.Split(txt, \"|\")")
	r.Add("for _, s := range txtVals {")
	r.Add("switch strings.TrimSpace(s) {")
	for _, c := range x.f.Cnsts {
		r.Addf("case %q:", c.Camel())
		r.Addf("r |= %v", c.Name)
	}
	r.Add("}")
	r.Add("}")
	r.Add("return r")
	return r
}
func (x *PrtEnumFlg) string() (r *TypFn) {
	r = x.f.TypFn(k.String)
	r.OutPrm(String)
	r.Addf("return %v[x]", x.TypNameVarName())
	return r
}
func (x *PrtEnumFlg) is() (r *TypFn) {
	r = x.f.TypFn(k.Is)
	r.InPrmVariadic(x.t, "vs")
	r.OutPrm(Bool, "r")
	r.Add("for _, v := range vs {")
	r.Add("if x&v != v {")
	r.Add("return false")
	r.Add("}")
	r.Add("}")
	r.Add("return true")
	return r
}
func (x *PrtEnumFlg) isFlg(c *Cnst) (r *TypFn) {
	r = x.f.TypFnf("Is%v", c.Title())
	r.OutPrm(Bool)
	r.Addf("return x&%[1]v == %[1]v", c.Name)
	return r
}
