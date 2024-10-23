package tpl

import (
	"fmt"
	"strings"
	"sys"
	"sys/err"
	"sys/tpl/atr"
	"sys/tpl/mod"
)

type (
	Prm struct {
		Lbl
		mod.Mod
		atr.Atr
		Typ  Typ
		Tst  Typ
		Lits []string
		Vals []string
	}
	InPrms  []*Prm
	OutPrms []*Prm
)

func (x *Prm) Cpy(a ...atr.Atr) (r *Prm) {
	r = &Prm{}
	r.Name = x.Name
	r.Mod = x.Mod
	if len(a) != 0 {
		r.Atr = a[0]
	} else {
		r.Atr = x.Atr
	}
	r.Typ = x.Typ
	r.Tst = x.Tst
	r.Lits = x.Lits
	r.Vals = x.Vals
	return r
}
func (x *Prm) LitVal(litVals ...string) *Prm {
	x.Lits = litVals
	if x.Typ.Bse().IsBsc() && !x.Typ.Bse().IsArr() {
		for _, litVal := range litVals {
			x.Vals = append(x.Vals, fmt.Sprintf("%v(%v)", x.Typ.Full(), litVal))
		}
	} else {
		x.Vals = litVals
	}
	return x
}
func (x *Prm) WritePrm(next *Prm, b *strings.Builder, f *FleBse) {

	// if x.IsLowercase() {
	// 	b.WriteString(strings.ToLower(x.Name)) // ext: interface{}
	// } else {

	// }
	b.WriteString(x.Name)
	if next == nil || x.Typ != next.Typ || x.Mod != next.Mod {
		b.WriteRune(' ')
		if x.IsVariadic() {
			b.WriteString("...")
		}
		if x.IsSlice() {
			b.WriteString("[]")
		}
		if x.IsPtr() {
			b.WriteRune('*')
		}
		// if x.Typ.Title() == "Interface" {
		// 	sys.Log("- Prm", x.Name, x.Typ.Title())
		// }
		b.WriteString(x.Typ.Ref(f))
	}
}
func (x *Prm) TestFld(f Fle, b *strings.Builder) string {
	b.Reset()
	b.WriteString(x.Camel())
	b.WriteRune(' ')
	if x.Typ.Bse().IsAna() {
		if x.Typ.Bse().ifc != nil {
			x.Typ.Bse().ifc.FnValSig(f, b)
		} else {
			x.Typ.Bse().FnValSig(f, b)
		}
	} else {
		b.WriteString(x.Typ.Ref(f))
	}
	return b.String()
}
func (x *Prm) TestFldAsn(b *strings.Builder, val ...bool) string {
	b.Reset()
	b.WriteString(x.Camel())
	b.WriteString(" := ")
	if len(val) != 0 {
		if len(x.Vals) != 0 {
			b.WriteString(x.Vals[0])
		} else {
			b.WriteString(x.Typ.Bse().Val())
		}
	} else { // cse
		b.WriteString("cse.")
		b.WriteString(x.Camel())
	}
	if x.Typ.Bse().IsAna() {
		// if x.Typ.Bse().Pkg.Name == k.Hst {
		// 	b.WriteString("(ap.Hst)")
		// } else {
		// 	b.WriteString("(ap.Rlt)")
		// }
		b.WriteString("()")
	}
	return b.String()
}

func (x *Prm) Asn() string {
	b := &strings.Builder{}
	b.WriteString(x.Name)
	b.WriteString(" := ")
	b.WriteString(x.Vals[0])
	return b.String()
}
func (x *Prm) AsnTest() string {
	b := &strings.Builder{}
	b.WriteString("e.")
	b.WriteString(strings.Title(x.Name))
	b.WriteString(" = ")
	b.WriteString(x.Vals[0])
	return b.String()
}

func (x *InPrms) Cpy(a ...atr.Atr) *InPrms {
	r := make(InPrms, len(*x))
	for n, p := range *x {
		r[n] = p.Cpy(a...)
	}
	return &r
}
func (x *InPrms) TestFld(f Fle, test *PkgFn, b *strings.Builder) {
	for _, prm := range *x {
		test.Add(prm.TestFld(f, b))
	}
	b.Reset()
}
func (x *InPrms) TestFldAsn(test *PkgFn, b *strings.Builder, val ...bool) {
	for _, prm := range *x {
		test.Add(prm.TestFldAsn(b, val...))
	}
	b.Reset()
}

func (x *InPrms) TestCall(b *strings.Builder, hst ...bool) {
	if len(*x) != 0 {
		for n, prm := range *x {
			if n != 0 {
				b.WriteRune(',')
			}
			b.WriteString(prm.Camel())
			if len(hst) != 0 && hst[0] {
				b.WriteString("Hst")
			}
		}
	}
}

func (x *InPrms) CallWrt(b *strings.Builder) {
	if len(*x) != 0 {
		for n, prm := range *x {
			if n != 0 {
				b.WriteRune(',')
			}
			b.WriteString(prm.Name)
			if prm.IsVariadic() {
				b.WriteString("...")
			}
		}
	}
}
func (x *InPrms) CallLitWrt(b *strings.Builder) {
	if len(*x) != 0 {
		for n, prm := range *x {
			if n != 0 && !prm.IsVariadic() {
				b.WriteRune(',')
			}
			var lit string
			if len(prm.Lits) == 0 {
				if !prm.IsVariadic() {
					if len(prm.Typ.Bse().Lits) != 0 {
						lit = prm.Typ.Bse().Lits[0]
					} else {
						lit = "<NONE>"
					}
				}
			} else {
				lit = prm.Lits[0]
			}
			b.WriteString(lit)
		}
	}
}
func (x *InPrms) CallValWrt(b *strings.Builder, f Fle) {
	if len(*x) != 0 {
		for n, prm := range *x {
			if n != 0 && !prm.IsVariadic() {
				b.WriteRune(',')
			}
			var val string
			if len(prm.Vals) == 0 {
				if !prm.IsVariadic() {
					if len(prm.Typ.Bse().Vals) != 0 {
						val = prm.Typ.Bse().Vals[0]
						f.Bse().Import(prm.Typ)
					} else {
						val = "<NONE>"
					}
				}
			} else {
				val = prm.Vals[0]
				f.Bse().Import(prm.Typ)
			}
			b.WriteString(val)
		}
	}
}
func (x *InPrms) CallNodeWrt(b *strings.Builder) {
	if len(*x) != 0 {
		for n, prm := range *x {
			if n != 0 {
				b.WriteRune(',')
			}
			b.WriteString("x.")
			b.WriteString(strings.Title(prm.Name))
			if prm.IsVariadic() {
				b.WriteString("...")
			}
		}
	}
}

func (x *InPrms) Ok() bool { return len(*x) != 0 }
func (x *InPrms) Cnt() int { return len(*x) }

func (x *InPrms) Fst() *Prm   { return (*x)[0] }
func (x *InPrms) Lst() *Prm   { return (*x)[len(*x)-1] }
func (x *InPrms) LstIdx() int { return len(*x) - 1 }

func (x *InPrms) AddInPrm(vs ...*Prm) { *x = append(*x, vs...) }

func (x *InPrms) InPrmMod(typ FleOrTyp, m mod.Mod, name ...string) (r *Prm) {
	r = &Prm{}
	if len(name) > 0 && name[0] != "" {
		r.Name = sys.Trim(name[0])
	}
	r.Typ = GetTyp(typ)
	r.Mod = m
	x.AddInPrm(r)
	return r
}
func (x *InPrms) InPrm(typ FleOrTyp, name ...string) (r *Prm) {
	return x.InPrmMod(typ, mod.None, name...)
}
func (x *InPrms) InPrmPtr(typ FleOrTyp, name ...string) (r *Prm) {
	return x.InPrmMod(typ, mod.Ptr, name...)
}
func (x *InPrms) InPrmSlice(typ FleOrTyp, name ...string) (r *Prm) {
	return x.InPrmMod(typ, mod.Slice, name...)
}
func (x *InPrms) InPrmVariadic(typ FleOrTyp, name ...string) (r *Prm) {
	return x.InPrmMod(typ, mod.Variadic, name...)
}
func (x *InPrms) InPrmArr(vs ...*Prm) {
	for _, v := range vs {
		if v.Typ.Bse().arr == nil {
			panic(fmt.Sprintf("prm '%v:%v' does not have an associated array", v.Name, v.Typ.Title()))
		}
		x.InPrm(v.Typ.Bse().arr, v.Name+"s")
	}
}

func (x *InPrms) TypRefs() (r Typs) {
	for _, prm := range *x {
		if prm.Typ != nil {
			r = append(r, prm.Typ)
		}
	}
	return r
}
func (x *InPrms) MayXpr() bool {
	for _, prm := range *x {
		// if prm.Typ == Interface {
		// 	return true
		// }
		// if _, isExt := prm.Typ.(*Ext); isExt || !prm.Typ.Bse().IsXpr() {
		// 	return false
		// }
		if !prm.Typ.Bse().IsXpr() {
			return false
		}
	}
	return true
}
func (x *InPrms) Lits() string { // for testing
	b := &strings.Builder{}
	if x != nil {
		for n, prm := range *x {
			if n != 0 {
				b.WriteRune(' ')
			}
			if prm.Lits != nil {
				b.WriteString(prm.Lits[0])
			} else {
				b.WriteString(prm.Typ.Bse().Lit())
			}
		}
	}
	return b.String()
}
func (x *InPrms) Vals(call ...bool) string { // for testing
	b := &strings.Builder{}
	if len(*x) != 0 {
		for n, prm := range *x {
			if n != 0 {
				b.WriteRune(',')
			}
			if prm.Vals != nil {
				b.WriteString(prm.Vals[0])
			} else {
				b.WriteString(prm.Typ.Bse().Val(call...))
			}
		}
	}
	return b.String()
}
func (x *InPrms) Vals2(call ...bool) string { // for testing
	b := &strings.Builder{}
	if len(*x) != 0 {
		for n, prm := range *x {
			if n != 0 {
				b.WriteRune(',')
			}
			b.WriteString(prm.Typ.Bse().Val(call...))
		}
	}
	return b.String()
}
func (x *InPrms) ValVs() string {
	b := &strings.Builder{}
	if x != nil {
		for n, _ := range *x {
			if n != 0 {
				b.WriteRune(' ')
			}
			b.WriteString("%v")
		}
	}
	return b.String()
}
func (x *InPrms) Names() string {
	b := &strings.Builder{}
	if x != nil {
		for n, prm := range *x {
			if n != 0 {
				b.WriteRune(',')
			}
			b.WriteString(prm.Camel())
		}
	}
	return b.String()
}

func (x *InPrms) WriteInPrms(b *strings.Builder, f *FleBse) {
	b.WriteString("(")
	if x != nil {
		for n, prm := range *x {
			if prm.Name == "" {
				err.Panic("in prm: parameter name is empty")
			}
			if n != 0 {
				b.WriteString(",")
			}
			var next *Prm
			if n+1 < len(*x) {
				next = (*x)[n+1]
			}
			prm.WritePrm(next, b, f)
		}
	}
	b.WriteString(")")
}

func (x *OutPrms) Ok() bool { return len(*x) != 0 }
func (x *OutPrms) Cnt() int { return len(*x) }

func (x *OutPrms) Fst() *Prm            { return (*x)[0] }
func (x *OutPrms) Lst() *Prm            { return (*x)[len(*x)-1] }
func (x *OutPrms) AddOutPrm(vs ...*Prm) { *x = append(*x, vs...) }
func (x *OutPrms) OutPrmMod(typ FleOrTyp, m mod.Mod, name ...string) (r *Prm) {
	r = &Prm{}
	if len(name) > 0 && name[0] != "" {
		r.Name = sys.Trim(name[0])
	}
	r.Typ = GetTyp(typ)
	r.Mod = m
	x.AddOutPrm(r)
	return r
}

func (x *OutPrms) AddOutPrms(typs ...Typ) (r []*Prm) {
	r = make([]*Prm, len(typs))
	for n, typ := range typs {
		r[n] = x.OutPrm(typ)
	}
	return r
}
func (x *OutPrms) OutPrm(typ FleOrTyp, name ...string) (r *Prm) {
	return x.OutPrmMod(typ, mod.None, name...)
}
func (x *OutPrms) OutPrmPtr(typ FleOrTyp, name ...string) (r *Prm) {
	return x.OutPrmMod(typ, mod.Ptr, name...)
}
func (x *OutPrms) OutPrmSlice(typ FleOrTyp, name ...string) (r *Prm) {
	return x.OutPrmMod(typ, mod.Slice, name...)
}
func (x *OutPrms) OutPrmVariadic(typ FleOrTyp, name ...string) (r *Prm) {
	return x.OutPrmMod(typ, mod.None, name...)
}
func (x *OutPrms) OutTyp() Typ {
	if len(*x) != 0 {
		return (*x)[0].Typ
	}
	return nil
}
func (x *OutPrms) TypRefs() (r Typs) {
	for _, prm := range *x {
		if prm.Typ != nil {
			r = append(r, prm.Typ)
		}
	}
	return r
}
func (x *OutPrms) MayXpr() bool {
	if len(*x) != 1 {
		return false
	}
	for _, prm := range *x {
		// if prm.Typ == Interface {
		// 	return true
		// }
		// if _, isExt := prm.Typ.(*Ext); isExt || !prm.Typ.Bse().IsXpr() {
		// 	return false
		// }
		if !prm.Typ.Bse().IsXpr() {
			return false
		}
	}
	return true
}
func (x *OutPrms) WriteOutPrms(b *strings.Builder, f *FleBse) {
	if len(*x) == 0 {
		return
	}
	if len(*x) == 1 && x.Fst().Name == "" {
		x.Fst().WritePrm(nil, b, f)
		return
	}
	b.WriteString("(")
	for n, prm := range *x {
		if n != 0 {
			b.WriteString(",")
		}
		var next *Prm
		if n+1 < len(*x) {
			next = (*x)[n+1]
		}
		prm.WritePrm(next, b, f)
	}
	b.WriteString(")")
}
