package tpl

import (
	"fmt"
	"strings"
	"sys"
	"sys/k"
	"sys/tpl/atr"
	"sys/tpl/mod"
)

type (
	Typ interface {
		Lblr
		Bse() *TypBse
		WriteDecl(b *strings.Builder, f Fle)
		Qual(f Fle, prefixs ...string) string
		Ref(f Fle) string
		Adr(f Fle) string
		Cast(f Fle, ptr ...bool) string
		TypRefs() Typs
		PkgTypTitle() string
		PkgTypSuffix() string
		Full() string
		PrefixCamel() string
		PrefixTitle() string
	}
	Typs []Typ

	TypBse struct {
		Lbl
		mod.Mod
		atr.Atr
		Pkg  *Pkg
		Fle  Fle
		Flds FldMap

		TypFns            // for types with typfns (not ifc)
		Ifcs              // current typ implments these interfaces
		ConcreteTyps Typs // current type is ifc; these are the concrete typs which implement the interface

		Fns     TypFnMap
		Lits    []string
		Vals    []string
		LitsJsn []string
		ValsJsn []string
		Knd     *Cnst
		Size    *Cnst
		bse     *Struct
		ifc     *Ifc // current type presents itself as an interface by default; StmBse presents itself as Stm; CndBse as Cnd
		arr     *Arr // Elm->Arr
		elm     *Ifc // Fbr->Elm
		// fbr        *Ifc // Elm->Fbr
		// wve        *Ifc
		Rng        *Struct // Elm->Rng
		Scp        *Struct
		LitTrmJsn  *Struct
		LexTrmJsn  *TypFn
		PrsTrmJsn  *PkgFn
		NmeLex     *TypFn
		LitTrm     *Struct
		LitLex     *TypFn
		PrsTrm     *PkgFn
		PrsTxt     *PkgFn
		IfcXpr     *Ifc
		LitXpr     *Struct
		AsnXpr     *Struct
		AcsXpr     *Struct
		ThenXpr    *Struct
		ElseXpr    *Struct
		EachXpr    *Struct
		PllEachXpr *Struct
		IfcAct     *Ifc
		LitAct     *Struct
		AsnAct     *Struct
		AcsAct     *Struct
		ThenAct    *Struct
		ElseAct    *Struct
		EachAct    *Struct
		PllEachAct *Struct

		XprFldGets    Flds
		XprFldSetGets Flds

		ActFldGets    Flds
		ActFldSetGets Flds
		ActCnsts      Cnsts
		ActVars       Vars
		ActPkgFns     PkgFns
		ActTypFns     TypFns
		ActMemSigs    MemSigs

		TstEql      *PkgFn
		TstSliceEql *PkgFn
		TstNotZero  *PkgFn
		TstAsc      *PkgFn
		TstDsc      *PkgFn

		Test    *FleTest // TODO: NEEDED?
		TestPth []*TestStp
	}
	TestStp struct {
		Fst    func(r *PkgFn)
		MdlFst func(r *PkgFn)
		MdlLst func(r *PkgFn)
		Lst    func(r *PkgFn)
	}
	Alias struct {
		TypBse
		Elm      Typ
		AliasMod mod.Mod
		TypFns
	}
	Arr struct { // see PrtArr
		*Alias
		Elm     *TypBse // Bol, Unt, Stm...
		New     *PkgFn
		Make    *PkgFn
		MakeEmp *PkgFn
		Segs    *PkgFn // bnds
	}
	// Fbr struct { // see PrtFbr
	// 	*Ifc         // InstrFbr, InrvlFbr, SideFbr...
	// 	bse  *Struct // InstrFbrBse, InrvlFbrBse, SideFbrBse... (USED FOR NODE)
	// 	arr  *Arr    // Instrs, Inrvls, Sides...
	// }
	Wve struct { // see PrtWve
		*Ifc         // InstrWve, InrvlWve, SideWve...
		bse  *Struct // InstrWveBse, InrvlWveBse, SideWveBse... (USED FOR NODE)
		arr  *Arr    // InstrFbrs, InrvlFbrs, SideFbrs...
	}
	Struct struct {
		TypBse
		Flds
		TypFns
	}
	Structs []*Struct
	Ifc     struct {
		TypBse
		Ifcs
		MemSigs
		TstSliceVars Vars
	}
	Ifcs []*Ifc
	Map  struct {
		TypBse
		Key Typ
		Val Typ
	}
	Func struct {
		TypBse
		InPrms
		OutPrms
		Block
	}
	Ext struct {
		TypBse
		Imports
	}
)

// Typs
func (x *Typs) Ok() bool { return len(*x) != 0 }
func (x *Typs) Cnt() int { return len(*x) }
func (x *Typs) Typ() Typ {
	if len(*x) == 0 {
		return nil
	}
	return (*x)[0]
}

func (x *Typs) PrtArr() *PrtArr {
	if len(*x) == 0 {
		return nil
	}
	return (*x)[0].Bse().PrtArr()
}
func (x *Typs) AddTyp(vs ...Typ) { *x = append(*x, vs...) }

func (x *Typs) Arr(name string, elmFle Fle, pkg *Pkg) (r *Arr) {
	r = &Arr{}
	x.AddTyp(r) // add before alias creation so that it is the first typ in slice for fle.Typ() call
	r.Elm = elmFle.Typ().Bse()
	r.Alias = &Alias{}
	r.Alias.TypBse = *NewTypBse()
	r.Alias.Name = strings.Title(name)
	r.Alias.Pkg = pkg
	r.Alias.Elm = elmFle.Typ() // must be typ (non TypBse)
	r.Alias.AliasMod = mod.Slice
	r.Alias.Mod = mod.Ptr
	r.Alias.Atr = r.Elm.Atr | atr.Arr
	r.Alias.Atr = r.Alias.Atr &^ atr.Bsc
	r.Alias.Atr = r.Alias.Atr &^ atr.Struct
	r.Alias.Atr = r.Alias.Atr &^ atr.Test // NO TESTING (FOR NOW)
	// r.Alias.Atr = r.Alias.Atr &^ atr.TestXpr // NO TESTING (FOR NOW)
	// r.Alias.Atr = r.Alias.Atr &^ atr.TestAct // NO TESTING (FOR NOW)
	elmFle.Typ().Bse().arr = r
	elmFle.Bse().arr = r
	return r
}

// func (x *Typs) Fbr(elmFle, fbrFle *FleBse) (r *Ifc) {
// 	r = &Ifc{}
// 	x.AddTyp(r) // add before ifc,bse creation so that it is the first typ in slice for fle.Typ() call
// 	r.TypBse = *NewTypBse()
// 	r.Name = strings.Title(fmt.Sprintf("%vFbr", elmFle.Typ().Camel()))
// 	r.Pkg = elmFle.Pkg
// 	r.Atr = atr.TypAnaIfc | atr.Fbr
// 	r.elm = elmFle.Typ().(*Ifc) // Prv, Instr, Inrvl...
// 	r.elm.fbr = r
// 	elmFle.fbr = r
// 	fbrFle.Pkg = r.Pkg
// 	fbrFle.Name = r.Camel()
// 	return r
// }
// func (x *Typs) Wve(elmFle, wveFle *FleBse) (r *Ifc) {
// 	r = &Ifc{}
// 	x.AddTyp(r) // add before ifc,bse creation so that it is the first typ in slice for fle.Typ() call
// 	r.TypBse = *NewTypBse()
// 	r.Name = strings.Title(fmt.Sprintf("%vWve", elmFle.Typ().Camel()))
// 	r.Pkg = elmFle.Pkg
// 	r.Atr = atr.TypAnaIfc
// 	r.elm = elmFle.fbr
// 	r.elm.wve = r
// 	r.elm.elm.wve = r
// 	elmFle.wve = r
// 	wveFle.Pkg = r.Pkg
// 	wveFle.Name = r.Camel()
// 	return r
// }

func (x *Typs) Alias(name string, alias Typ, pkg *Pkg, aliasMod mod.Mod, a atr.Atr) (r *Alias) {
	r = &Alias{}
	r.TypBse = *NewTypBse()
	r.Name = strings.Title(name)
	r.Pkg = pkg
	r.Elm = alias
	r.AliasMod = aliasMod
	r.Atr = a
	x.AddTyp(r)
	return r
}
func (x *Typs) Struct(name string, pkg *Pkg, m mod.Mod, a atr.Atr) (r *Struct) {
	r = &Struct{}
	r.TypBse = *NewTypBse()
	r.Name = strings.Title(name)
	r.Pkg = pkg
	r.Mod = m
	r.Atr = a
	x.AddTyp(r)
	return r
}
func (x *Typs) Ifc(name string, pkg *Pkg, a atr.Atr) (r *Ifc) {
	r = &Ifc{}
	r.TypBse = *NewTypBse()
	r.Name = strings.Title(name)
	r.Pkg = pkg
	r.Atr = a
	x.AddTyp(r)
	return r
}

func (x *Typs) Map(name string, key, val FleOrTyp, pkg *Pkg, a atr.Atr) (r *Map) {
	r = &Map{}
	r.TypBse = *NewTypBse()
	r.Name = strings.Title(name)
	r.Pkg = pkg
	r.Key = GetTyp(key)
	r.Val = GetTyp(val)
	r.Atr = a
	x.AddTyp(r)
	return r
}
func (x *Typs) Func(name string, pkg *Pkg, a atr.Atr) (r *Func) {
	r = &Func{}
	r.TypBse = *NewTypBse()
	r.Name = strings.Title(name)
	r.Pkg = pkg
	r.Atr = a
	x.AddTyp(r)
	return r
}
func (x *Typs) Ext(name string, pkg ...*Pkg) (r *Ext) {
	r = NewExt(name, pkg...)
	x.AddTyp(r)
	return r
}
func (x *Typs) Extf(format string, args ...interface{}) (r *Ext) {
	return x.Ext(fmt.Sprintf(format, args...))
}
func (x *Typs) WriteTypDecls(b *strings.Builder, f Fle) {
	if len(*x) > 0 {
		b.WriteString("type (\n")
		for _, t := range *x {
			t.WriteDecl(b, f)
		}
		b.WriteString(")\n\n")
	}
}
func (x *Typs) WriteTypEmbeds(b *strings.Builder, f Fle) {
	if len(*x) > 0 {
		for _, t := range *x {
			b.WriteString(t.Ref(f))
			b.WriteRune('\n')
		}
	}
}

// TypBase
func NewTypBse() (r *TypBse) {
	r = &TypBse{}
	r.Flds = make(FldMap)
	r.Fns = make(TypFnMap)
	// r.Sigs = make(MemSigMap)
	return r
}
func (x *TypBse) Bse() *TypBse      { return x }
func (x *TypBse) NewTitle() string  { return x.UniqueTitle("New") }
func (x *TypBse) MakeTitle() string { return x.UniqueTitle("Make") }
func (x *TypBse) UniqueTitle(name string) string {
	name = strings.Title(name)
	if x.Lower() != x.Pkg.Lower() {
		// NewStms, MakeStms, ...
		return fmt.Sprintf("%v%v", name, x.Title())
	}
	return name
}
func (x *TypBse) PkgTypTitle() string {
	if x.Pkg == nil {
		return ""
	}
	return fmt.Sprintf("%v%v", x.Pkg.Title(), x.Name)
}
func (x *TypBse) PkgTypSuffix() string {
	if x.Pkg == nil {
		return ""
	}
	var suffix string
	if x.Pkg.Lower() != x.Lower() {
		suffix = x.Title()
	}
	return fmt.Sprintf("%v%v", x.Pkg.Title(), suffix)
}
func (x *TypBse) PkgTypCamel() string {
	if x.Pkg == nil {
		return ""
	}
	return fmt.Sprintf("%v%v", x.Pkg.Camel(), x.Name)
}
func (x *TypBse) Full() string {
	if x.Pkg == nil {
		return ""
	}
	return fmt.Sprintf("%v.%v", x.Pkg.Lower(), x.Name)
}
func (x *TypBse) PrefixCamel() string { return sys.Camel(x.PrefixTitle()) }
func (x *TypBse) PrefixTitle() string {
	if x.Pkg == nil || x.Pkg.Title() == x.Title() {
		return x.Title()
	} // ensure pkg name prefixs if pkg name is not equal to the typ name
	return fmt.Sprintf("%v%v", x.Pkg.Title(), x.Title())
}
func (x *TypBse) PrefixTyp(name string) string {
	if x.Pkg == nil || x.Pkg.Title() == x.Title() {
		return name
	}
	return fmt.Sprintf("%v%v", x.Title(), strings.Title(name))
}
func (x *TypBse) Qual(f Fle, prefixs ...string) string {
	fleBse := f.Bse()
	b := &strings.Builder{}
	for _, prefix := range prefixs {
		b.WriteString(prefix)
	}
	// sys.Log("- Qual", x.Name, x.IsLowercase(), x.IsCurlyBracket())
	if x.IsLowercase() && x.IsCurlyBracket() { // ext: interface{}
		b.WriteString(strings.ToLower(x.Name))
		b.WriteString("{}")
		return b.String()
	}
	if x.Pkg == nil || x.Pkg == fleBse.Pkg {
		b.WriteString(x.Name)
		return b.String()
	}
	fleBse.Import(x)
	if len(fleBse.Imports) == 0 {
		b.WriteString(x.Full())
		return b.String()
	}
	pkgPth := x.Pkg.Pth
	var imp *Import
	for _, cur := range fleBse.Imports {
		if pkgPth == cur.Pth && cur.Alias != "" {
			imp = cur
			break
		}
	}
	if imp != nil {
		b.WriteString(imp.Alias)
		b.WriteRune('.')
		b.WriteString(x.Name)
	} else {
		b.WriteString(x.Full())
	}
	return b.String()
}
func (x *TypBse) Adr(f Fle) string {
	var prefixs []string
	if x.IsPtr() {
		prefixs = append(prefixs, "&")
	}
	return x.Qual(f, prefixs...)
}
func (x *TypBse) Ref(f Fle) string {
	var prefixs []string
	if x.IsSlice() {
		prefixs = append(prefixs, "[]")
	}
	if x.IsPtr() {
		prefixs = append(prefixs, "*")
	}
	return x.Qual(f, prefixs...)
}
func (x *TypBse) WriteDecl(b *strings.Builder, f Fle) {}
func (x *TypBse) TypRefs() Typs                       { return nil }
func (x *TypBse) String() string                      { return x.Name }
func (x *TypBse) Prm(name string) (r *Prm) {
	r = &Prm{}
	r.Name = name
	r.Typ = x
	return r
}
func (x *TypBse) Make(f Fle) string { return fmt.Sprintf("%v(%v)", k.Make, x.Qual(f)) }
func (x *TypBse) Cast(f Fle, ptr ...bool) string {
	if x.Pkg != nil && x.Pkg != f.Bse().Pkg {
		f.Bse().Import(x)
	}
	b := &strings.Builder{}
	b.WriteString(".(")
	if len(ptr) != 0 && ptr[0] {
		b.WriteRune('*')
	}
	b.WriteString(x.Qual(f))
	b.WriteRune(')')
	return b.String()
}
func (x *TypBse) Fn(name string) (r *TypFn) {
	name = strings.Title(name)
	r, ok := x.Fns[name]
	if !ok {
		for _, fld := range x.Flds {
			if fld.Name == "" { // "" implies embedded typ
				r = fld.Typ.Bse().Fn(name)
				if r != nil {
					break
				}
			}
		}
	}
	return r
}
func (x *TypBse) Lit() string { // for testing

	if len(x.Lits) == 0 {
		sys.Log("*", x.Name)
		panic("<NONE>")
		// return "<NONE>"
	}
	if len(x.Lits) == 1 { // ana sngl
		return x.Lits[0]
	}
	return x.Lits[1] // non-zero or non-empty value
}
func (x *TypBse) Val(call ...bool) string { // for testing
	if len(x.Vals) == 0 {
		panic("<NONE>")
		// return "<NONE>"
	}
	var idx int
	if len(x.Vals) > 1 {
		idx = 1 // index 1 is non-empty array or more interesting
	}
	return x.ValAt(idx, call...)
}
func (x *TypBse) ValAt(idx int, call ...bool) string {
	if x.IsBsc() {
		return fmt.Sprintf("%v(%v)", x.Full(), x.Vals[idx])
	}
	if x.IsAna() {
		b := &strings.Builder{}
		x.FnVal(idx, b)
		if len(call) != 0 && call[0] {
			// if x.Pkg.Name == k.Hst {
			// 	b.WriteString("(ap.Hst)")
			// } else {
			// 	b.WriteString("(ap.Rlt)")
			// }
			b.WriteString("()")
		}
		return b.String()
	}
	return x.Vals[idx]
}
func (x *TypBse) FnValSig(f Fle, b *strings.Builder, prm ...*Prm) string { // for testing
	// if x.Pkg.Name == k.Hst {
	// 	b.WriteString("func(hst *hst.Hst) ")
	// } else {
	// 	b.WriteString("func(rlt *rlt.Rlt) ")
	// }
	b.WriteString("func() ")
	if len(prm) != 0 && prm[0].IsVariadic() {
		b.WriteString("[]")
	}
	if x.ifc != nil {
		b.WriteString(x.ifc.Ref(f))
	} else {
		b.WriteString(x.Ref(f))
	}
	return b.String()
}
func (x *TypBse) FnVal(idx int, bs ...*strings.Builder) string { // for testing
	var b *strings.Builder
	if len(bs) != 0 {
		b = bs[0]
	} else {
		b = &strings.Builder{}
	}
	// if x.Pkg.Name == k.Hst {
	// 	b.WriteString("func(hst *hst.Hst) ")
	// } else {
	// 	b.WriteString("func(rlt *rlt.Rlt) ")
	// }
	// b.WriteString("func(ana *ana.Ana) ")
	b.WriteString("func() ")
	if x.ifc != nil {
		b.WriteString(x.ifc.Full())
	} else {
		if x.IsPtr() {
			b.WriteString("*")
		}
		b.WriteString(x.Full())
	}
	b.WriteString(" { return ")
	b.WriteString(x.Vals[idx])
	b.WriteString(" }")
	return b.String()
}
func (x *TypBse) FnVals() (r []string) {
	b := &strings.Builder{}
	for n := 0; n < len(x.Vals); n++ {
		b.Reset()
		r = append(r, x.FnVal(n, b))
	}
	return r
}

func (x *TypBse) PrtAri() *PrtAri {
	if x.Fle == nil {
		return nil
	}
	prt := x.Fle.Bse().GetPrt((*PrtAri)(nil))
	if prt == nil {
		return nil
	}
	return prt.(*PrtAri)
}
func (x *TypBse) PrtArr() *PrtArr {
	if x.Fle == nil {
		return nil
	}
	prt := x.Fle.Bse().GetPrt((*PrtArr)(nil))
	if prt == nil {
		return nil
	}
	return prt.(*PrtArr)
}
func (x *TypBse) PrtRng() *PrtRng {
	if x.Fle == nil {
		return nil
	}
	prt := x.Fle.Bse().GetPrt((*PrtRng)(nil))
	if prt == nil {
		return nil
	}
	return prt.(*PrtRng)
}

func (x *TypBse) PrtStructBytWrt() *PrtStructBytWrt {
	if x.Fle == nil {
		return nil
	}
	prt := x.Fle.Bse().GetPrt((*PrtStructBytWrt)(nil))
	if prt == nil {
		return nil
	}
	return prt.(*PrtStructBytWrt)
}
func (x *TypBse) PrtArrBytWrt() *PrtArrBytWrt {
	if x.Fle == nil {
		return nil
	}
	prt := x.Fle.Bse().GetPrt((*PrtArrBytWrt)(nil))
	if prt == nil {
		return nil
	}
	return prt.(*PrtArrBytWrt)
}
func (x *TypBse) PrtArrRel() *PrtArrRel {
	if x.Fle == nil {
		return nil
	}
	prt := x.Fle.Bse().GetPrt((*PrtArrRel)(nil))
	if prt == nil {
		return nil
	}
	return prt.(*PrtArrRel)
}
func (x *TypBse) ArrStrWrt() *PrtArrStrWrt {
	if x.Fle == nil {
		return nil
	}
	prt := x.Fle.Bse().GetPrt((*PrtArrStrWrt)(nil))
	if prt == nil {
		return nil
	}
	return prt.(*PrtArrStrWrt)
}
func (x *TypBse) BytWrt() *PrtBytes { return x.Fle.Bse().GetPrt((*PrtBytes)(nil)).(*PrtBytes) }
func (x *TypBse) Idn() *PrtIdn      { return x.Fle.Bse().GetPrt((*PrtIdn)(nil)).(*PrtIdn) }
func (x *TypBse) IdnStruct() *PrtStructIdn {
	return x.Fle.Bse().GetPrt((*PrtStructIdn)(nil)).(*PrtStructIdn)
}
func (x *TypBse) Pkt() *PrtPkt       { return x.Fle.Bse().GetPrt((*PrtPkt)(nil)).(*PrtPkt) }
func (x *TypBse) Sel() *PrtSel       { return x.Fle.Bse().GetPrt((*PrtSel)(nil)).(*PrtSel) }
func (x *TypBse) Sgn() *PrtSgn       { return x.Fle.Bse().GetPrt((*PrtSgn)(nil)).(*PrtSgn) }
func (x *TypBse) StrWrt() *PrtString { return x.Fle.Bse().GetPrt((*PrtString)(nil)).(*PrtString) }
func (x *TypBse) Agg() *PrtArrAgg    { return x.Fle.Bse().GetPrt((*PrtArrAgg)(nil)).(*PrtArrAgg) }

func (x *Alias) WriteDecl(b *strings.Builder, f Fle) {
	b.WriteString(x.Name)
	b.WriteRune(' ')
	if x.AliasMod.IsSlice() {
		b.WriteString("[]")
	}
	if x.AliasMod.IsPtr() {
		b.WriteRune('*')
	}
	b.WriteString(x.Elm.Ref(f))
	b.WriteRune('\n')
}
func (x *Alias) TypRefs() (r Typs) {
	r = append(r, x.Elm)
	r = append(r, x.TypFns.TypRefs()...)
	return r
}

// Struct
func (x *Structs) AddStruct(vs ...*Struct) { *x = append(*x, vs...) }
func (x *Struct) FldMod(name string, typ FleOrTyp, m mod.Mod) (r *Fld) {
	r = &Fld{}
	r.Name = name
	r.Typ = GetTyp(typ)
	r.Struct = x
	r.Mod = m
	r.Atr = atr.LngTest
	x.Flds.AddFld(r)
	x.Bse().Flds[r.Name] = r
	return r
}
func (x *Struct) Fld(name string, typ FleOrTyp) (r *Fld) {
	return x.FldMod(name, typ, mod.None)
}
func (x *Struct) Fldf(format string, typ FleOrTyp, args ...interface{}) (r *Fld) {
	return x.FldMod(fmt.Sprintf(format, args...), typ, mod.None)
}
func (x *Struct) FldPtr(name string, typ FleOrTyp) (r *Fld) {
	return x.FldMod(name, typ, mod.Ptr)
}
func (x *Struct) FldSlice(name string, typ FleOrTyp) (r *Fld) {
	return x.FldMod(name, typ, mod.Slice)
}
func (x *Struct) FldSlicef(format string, typ FleOrTyp, args ...interface{}) (r *Fld) {
	return x.FldMod(fmt.Sprintf(format, args...), typ, mod.Slice)
}
func (x *Struct) FldSlicePtr(name string, typ FleOrTyp) (r *Fld) {
	return x.FldMod(name, typ, mod.Slice|mod.Ptr)
}
func (x *Struct) FldTyp(typ FleOrTyp) (r *Fld) {
	return x.FldMod("", typ, mod.None)
}
func (x *Struct) FldBse(typ FleOrTyp) (r *Fld) {
	r = x.FldTyp(typ)
	x.bse = r.Typ.(*Struct)
	return r
}
func (x *Struct) FldPrm(p *Prm) (r *Fld) {
	if p.IsVariadic() {
		return x.FldSlice(p.Title(), p.Typ)
	}
	return x.Fld(p.Title(), p.Typ)
}
func (x *Struct) FldPrnt(ft FleOrTyp) (r *Fld) {
	typ := GetTyp(ft)
	r = x.FldMod(typ.Title(), typ, mod.None)
	r.Atr = atr.Prnt
	// TODO: REMOVE?
	// // node: setter
	// fn1 := x.Fle.Bse().TypFnRxrf("%vSet", x, typ.Title()) // impl PrtPrntIfc
	// fn1.Atr = atr.None
	// fn1.InPrm(typ, "v")
	// fn1.Addf("x.%v = v", typ.Title())
	// // node: getter
	// fn2 := x.Fle.Bse().TypFnRxrf("%vGet", x, typ.Title()) // impl PrtPrntIfc
	// fn2.Atr = atr.None
	// fn2.OutPrm(typ)
	// fn2.Addf("return x.%v", typ.Title())
	return r
}
func (x *Struct) Prnt() *Fld {
	// USE OneFldInv TO GET TOP MOST PRNT IN CASE OF DOUBLE PRNT
	return x.OneFldInv(func(f *Fld) bool { return f.IsPrnt() })
}

// func (x *Struct) Arr() *Fld { // for Fbr.Arr
// 	// USE OneFldInv TO GET TOP MOST ARR IN CASE OF DOUBLE ARR
// 	return x.OneFldInv(func(f *Fld) bool { return f.IsArr() })
// }
func (x *Struct) FldExt(name string, pkg ...*Pkg) (r *Fld) {
	return x.FldTyp(NewExt(name, pkg...))
}
func (x *Struct) WriteDecl(b *strings.Builder, f Fle) {
	b.WriteString(x.Name)
	b.WriteString(" struct {")
	x.WriteFlds(b, f.Bse())
	b.WriteRune('}')
	b.WriteRune('\n')
}
func (x *Struct) TypRefs() (r Typs) {
	r = append(r, x.Flds.TypRefs()...)
	return r
}

// Ifcs
func (x *Ifcs) Ok() bool          { return len(*x) != 0 }
func (x *Ifcs) Cnt() int          { return len(*x) }
func (x *Ifcs) AddIfc(vs ...*Ifc) { *x = append(*x, vs...) }
func (x *Ifcs) EmbedIfc(typ FleOrTyp) (r *Ifc) {
	r = GetTyp(typ).(*Ifc)
	*x = append(*x, r)
	return r
}
func (x *Ifcs) Ins(idx int, vs ...*Ifc) *Ifcs {
	*x = append((*x)[:idx], append(vs, (*x)[idx:]...)...)
	return x
}
func (x *Ifcs) Push(a ...*Ifc) *Ifcs {
	*x = append(*x, a...)
	return x
}
func (x *Ifcs) Pop() (r *Ifc) {
	if len(*x) == 0 {
		return r
	}
	r = (*x)[len(*x)-1]
	*x = (*x)[:len(*x)-1]
	return r
}
func (x *Ifcs) Dque() (r *Ifc) {
	r = (*x)[0]
	if len(*x) == 1 {
		*x = (*x)[:0]
	} else {
		copy(*x, (*x)[1:])
		*x = (*x)[:len(*x)-1]
	}
	return r
}

// Ifc
func (x *Ifc) MayXpr() bool {
	return x.IsXpr() && x.MemSigs.MayXpr()
}
func (x *Ifc) XprMemSigs() MemSigs {
	// lookf for xprbl sigs in cur or embedded
	var r, memSigs MemSigs
	for _, ifc := range x.Ifcs {
		memSigs = append(memSigs, ifc.MemSigs...)
	}
	memSigs = append(memSigs, x.MemSigs...)
	for _, sig := range memSigs {
		if sig.MayXpr() {
			r = append(r, sig)
		}
	}
	return r
}

// func (x *Ifc) MayXpr() bool {
// 	return x.IsXpr() && x.MemSigs.MayXpr()
// }
func (x *Ifc) WriteDecl(b *strings.Builder, f Fle) {
	b.WriteString(x.Name)
	b.WriteString(" interface {\n")
	if len(x.Ifcs) > 0 { // write embedded ifcs
		for _, i := range x.Ifcs {
			b.WriteString(i.Ref(f))
			b.WriteRune('\n')
		}
	}
	x.WriteMemSigs(b, f.Bse())
	b.WriteRune('}')
	b.WriteRune('\n')
}
func (x *Ifc) TypRefs() (r Typs) {
	for _, i := range x.Ifcs {
		r = append(r, i)
	}
	r = append(r, x.MemSigs.TypRefs()...)
	return r
}

func (x *Map) WriteDecl(b *strings.Builder, f Fle) {
	b.WriteString(x.Name)
	b.WriteRune(' ')
	b.WriteString(k.Map)
	b.WriteRune('[')
	b.WriteString(x.Key.Ref(f))
	b.WriteRune(']')
	b.WriteString(x.Val.Ref(f))
	b.WriteRune('\n')
}

func (x *Map) TypRefs() (r Typs) {
	r = append(r, x.Key)
	r = append(r, x.Val)
	return r
}

// Func
func (x *Func) AddFn(fn Fn) *Func {
	x.AddInPrm(fn.In()...)
	x.AddOutPrm(fn.Out()...)
	return x
}

func (x *Func) WriteDecl(b *strings.Builder, f Fle) {
	b.WriteString(x.Name)
	b.WriteRune(' ')
	b.WriteString(k.Func)
	x.InPrms.WriteInPrms(b, f.Bse())
	x.OutPrms.WriteOutPrms(b, f.Bse())
	b.WriteRune('\n')
}
func (x *Func) TypRefs() (r Typs) {
	r = append(r, x.InPrms.TypRefs()...)
	r = append(r, x.OutPrms.TypRefs()...)
	return r
}

func (x *Func) DeclWrt(b *strings.Builder, f Fle) string {
	b.WriteString(k.Func)
	x.InPrms.WriteInPrms(b, f.Bse())
	x.OutPrms.WriteOutPrms(b, f.Bse())
	if x.Lines.Exist() {
		x.WriteBlock(b, true)
	}
	return b.String()
}
func (x *Func) Decl(f Fle) string {
	var b strings.Builder
	x.DeclWrt(&b, f)
	return b.String()
}

func NewExtMod(name string, m mod.Mod, pkg ...*Pkg) (r *Ext) {
	r = &Ext{}
	r.Name = name
	r.Mod = m
	if len(pkg) != 0 {
		r.Pkg = pkg[0]
	}
	return r
}
func NewExt(name string, pkg ...*Pkg) (r *Ext) {
	return NewExtMod(name, mod.None, pkg...)
}
func NewExtf(format string, args ...interface{}) (r *Ext) {
	return NewExt(fmt.Sprintf(format, args...))
}
func NewExta(name string, m mod.Mod, a atr.Atr, pkg ...*Pkg) (r *Ext) {
	r = NewExtMod(name, m, pkg...)
	r.Atr = a
	return r
}

var (
	AnaPkg     = NewPkg("sys/ana")
	BytesPkg   = NewPkg("bytes")
	StringsPkg = NewPkg("strings")
	TimePkg    = NewPkg("time")
	SyncPkg    = NewPkg("sync")
	TestingPkg = NewPkg("testing")
	HttpPkg    = NewPkg("net/http")
	FontPkg    = NewPkg("golang.org/x/image/font")
	ColorPkg   = NewPkg("image/color")

	String       = NewExt("string")
	Bool         = NewExt("bool")
	Byte         = NewExt("byte")
	Uint         = NewExt("uint")
	Uint32       = NewExt("uint32")
	Uint64       = NewExt("uint64")
	Int          = NewExt("int")
	Int32        = NewExt("int32")
	Int64        = NewExt("int64")
	Float32      = NewExt("float32")
	Float64      = NewExt("float64")
	Rune         = NewExt("rune")
	Time         = NewExt("Time", TimePkg)
	Duration     = NewExt("Duration", TimePkg)
	BufferPtr    = NewExtMod("Buffer", mod.Ptr, BytesPkg)
	Builder      = NewExt("Builder", StringsPkg)
	BuilderPtr   = NewExtMod("Builder", mod.Ptr, StringsPkg)
	T            = NewExt("T", TestingPkg) // TODO: MOD
	Interface    = NewExta("Interface", mod.Lowercase|mod.CurlyBracket, atr.Ext)
	Mutex        = NewExt("Mutex", SyncPkg)
	MutexPtr     = NewExtMod("Mutex", mod.Ptr, SyncPkg)
	RWMutex      = NewExt("RWMutex", SyncPkg)
	WaitGroupPtr = NewExtMod("WaitGroup", mod.Ptr, SyncPkg)
	Error        = NewExt("error")
	Face         = NewExt("Face", FontPkg)
	AnaOan       = NewExtMod("Oan", mod.Ptr, AnaPkg) // place here to get Tst generation
)
var (
	ExtTsts = []*Ext{
		String, Bool, Byte,
		Uint32, Uint64,
		Int32, Int64,
		Float32, Float64, Rune,
		Time, Duration, AnaOan,
	}
	// ExtLngs = []*Ext{
	// 	Interface,
	// }
)

func (x *TypBse) LitsNonEmp() []string {
	if len(x.Lits) != 0 && x.Lits[len(x.Lits)-1] == "[]" {
		return x.Lits[:len(x.Lits)-1] // skp empty arr; empty arr is invalid
	}
	return x.Lits
}
