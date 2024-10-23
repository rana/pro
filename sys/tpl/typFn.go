package tpl

import (
	"fmt"
	"regexp"
	"strings"
)

type (
	TypFn struct {
		FnBse
		Imports
		Block
		Rxr     *Prm
		CastOut bool
		Node    *Struct
		Cmd     *Struct
	}
	TypFns   []*TypFn
	TypFnMap map[string]*TypFn
)

func (x *TypFns) Ok() bool          { return len(*x) != 0 }
func (x *TypFns) Cnt() int          { return len(*x) }
func (x *TypFns) AddTypFn(v *TypFn) { *x = append(*x, v) }
func (x *TypFns) RemTypFn(name string) {
	for idx, v := range *x {
		if v.Name == name {
			if idx == 0 && len(*x) == 1 {
				*x = (*x)[:0]
			} else if idx == len(*x)-1 {
				*x = (*x)[:idx]
			} else {
				*x = append((*x)[:idx], (*x)[idx+1:]...)
			}
			break
		}
	}
}
func (x *TypFns) TypRefs() (r Typs) {
	for _, fn := range *x {
		r = append(r, fn.TypRefs()...)
	}
	return r
}
func (x *TypFns) MayXpr() bool { // at least one
	for _, v := range *x {
		if v.MayXpr() {
			return true
		}
	}
	return false
}

func (x *TypFns) WriteTypFns(b *strings.Builder, f *FleBse) {
	for _, fn := range *x {
		fn.WriteTypFn(b, f)
	}
}
func (x *TypFn) WriteTypFn(b *strings.Builder, f *FleBse) {
	b.WriteString("func ")
	if x.Rxr != nil {
		b.WriteRune('(')
		x.Rxr.WritePrm(nil, b, f)
		b.WriteString(") ")
	}
	b.WriteString(x.Name)
	if len(x.InPrms) > 0 {
		x.InPrms.WriteInPrms(b, f)
	} else {
		b.WriteString("()")
	}
	if len(x.OutPrms) > 0 {
		b.WriteRune(' ')
		x.OutPrms.WriteOutPrms(b, f)
	}
	x.WriteBlock(b)
}

func (x *TypFn) TestCall(b *strings.Builder, rxrName ...string) string {
	b.Reset()
	if len(rxrName) != 0 {
		b.WriteString(rxrName[0])
	} else {
		b.WriteString(x.Rxr.Name)
	}
	b.WriteRune('.')
	if len(rxrName) == 2 {
		b.WriteString(rxrName[1])
	} else {
		b.WriteString(x.Name)
	}
	b.WriteRune('(')
	x.InPrms.TestCall(b)
	b.WriteRune(')')
	return b.String()
}
func (x *TypFn) CallWrt(b *strings.Builder, rxrName ...string) {
	if len(rxrName) != 0 {
		b.WriteString(rxrName[0])
	} else {
		b.WriteString(x.Rxr.Name)
	}
	x.FnBse.CallWrt(b)
}
func (x *TypFn) Call(rxrName ...string) string {
	var b strings.Builder
	x.CallWrt(&b, rxrName...)
	return b.String()
}
func (x *TypFn) Vals(call ...bool) string {
	b := &strings.Builder{}
	if len(x.InPrms) != 0 {
		for n, prm := range x.InPrms {
			if n != 0 {
				b.WriteRune(' ')
			}
			if prm.Vals != nil {
				b.WriteString(prm.Vals[0])
			} else {
				if x.Rxr.Typ == prm.Typ { // allow RltStm prm to have same instr for testing

				} else {
					b.WriteString(prm.Typ.Bse().Val(call...))
				}
			}
		}
	}
	return b.String()
}
func (x *TypFn) Full() string {
	return fmt.Sprintf("%v.%v", x.Rxr.Typ.Full(), x.Name)
}
func (x *TypFn) TestStrWrt(test Fle, name ...string) (r Lines) {
	test.Bse().Import("strings")
	if len(name) != 0 {
		r.Addf("%v.StrWrt = func(b %v) {", name[0], BuilderPtr.Ref(test))
	} else {
		r.Addf("expected.StrWrt = func(b %v) {", BuilderPtr.Ref(test))
	}
	r.Add("b.WriteString(cse.pth)")
	r.Add("b.WriteRune('.')")
	r.Addf("b.WriteString(%q)", x.Camel())
	r.Add("b.WriteRune('(')")
	if len(x.InPrms) != 0 {
		for n, prm := range x.InPrms {
			if n != 0 {
				r.Add("b.WriteRune(' ')")
			}
			if prm.Typ.Bse().IsAna() {
				r.Addf("b.WriteString(%v.String())", prm.Name)
			} else {
				r.Addf("%v.StrWrt(b)", prm.Name)
			}
		}
	}
	r.Add("b.WriteRune(')')")
	r.Add("}")
	return r
}

func (x *TypFn) Inline(dst *TypFn, f Fle, rxr, ret string, prms ...string) (r Lines) {
	f.Bse().Imports.AddImport(x.Imports...)
	r = x.Lines.Cpy()
	Ret := fmt.Sprintf("%v =", ret)
	regexRxr0 := regexp.MustCompile("\\bx{1}\\b")
	// regexRxr1 := regexp.MustCompile("\\(x\\){1}")
	// regexRxr2 := regexp.MustCompile("([x]){1}[^\\.]")
	for n := 0; n < len(r); n++ {

		// r[n] = strings.Replace(r[n], "x", rxr, -1) // rxr
		r[n] = regexRxr0.ReplaceAllString(r[n], rxr)
		// r[n] = regexRxr1.ReplaceAllString(r[n], rxr)
		// r[n] = regexRxr2.ReplaceAllString(r[n], rxr)
		r[n] = strings.Replace(r[n], "return", Ret, -1) // return
		for p, inPrm := range x.InPrms {
			//\ba{1}\b
			regexPrm := regexp.MustCompile(fmt.Sprintf("\\b%v{1}\\b", inPrm.Name))
			r[n] = regexPrm.ReplaceAllString(r[n], prms[p])
		}
		if x.CastOut {
			r[n] = strings.Replace(r[n], x.OutTyp().Title(), x.OutTyp().Full(), -1)
			f.Bse().Import(x.OutTyp())
		}

		dst.Add(r[n])
	}
	return r
}
func (x *TypFn) Preamble() {
	if x.Node != nil {
		x.Addf("r := %v{}", x.Node.Adr(x.Rxr.Typ.Bse().Fle))
		if x.Node.bse != nil {
			for _, f := range x.Node.bse.Flds {
				if f.IsPrnt() {
					x.Addf("r.%v = x", f.Name)
				}
			}
		}
		for _, p := range x.InPrms {
			x.Node.FldPrm(p)
			x.Addf("r.%v = %v", p.Title(), p.Camel())
		}
		//x.Node.FldPrm(r.InPrmVariadic(_sys.Bsc.TmeRng, "rng"))
		//x.Addf("r := %v{}", r.Node.Adr(x))
	}
}
func (x *TypFn) MayXpr() bool {
	return x.Rxr.Typ.Bse().IsXpr() && x.FnBse.MayXpr()
}
