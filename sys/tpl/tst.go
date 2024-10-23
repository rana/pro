package tpl

import (
	"fmt"
	"strings"
	"sys/k"
	"sys/tpl/atr"
)

type (
	DirTst struct {
		DirBse
		Tst *FleTst
	}
	FleTst struct {
		FleBse
		NodeFuncs map[string]*TstNodeFunc
	}
	TstNodeFunc struct {
		Func *Func
		Vars
	}
)

func (x *DirSys) NewTst() (r *FleTst) {
	r = &FleTst{}
	x.Tst = r
	r.Name = k.Tst
	r.Pkg = x.Pkg.New(r.Name)
	r.NodeFuncs = make(map[string]*TstNodeFunc)
	r.AddFle(r) // NO TST TYP
	return r
}

func (x *FleTst) PkgFn(name string) (r *PkgFn) {
	r = x.FleBse.PkgFn(name)
	r.InPrmPtr(T, "t")
	return r
}
func (x *FleTst) PkgFnf(format string, args ...interface{}) (r *PkgFn) {
	return x.PkgFn(fmt.Sprintf(format, args...))
}

func (x *FleTst) GenCnst(typ Typ) {
	for _, c := range typ.Bse().Fle.Bse().Cnsts { // cnsts
		if c.MayXpr() {
			x.CnstYes(typ, c)
			x.CnstNot(typ, c)
		}
	}
}
func (x *FleTst) CnstYes(typ Typ, c *Cnst) (r *PkgFn) {
	x.Import("fmt")
	r = x.PkgFnf("%v%v", typ.PrefixTitle(), c.Name)
	r.InPrm(typ, "a")
	r.InPrmVariadic(Interface, "msgs")
	r.Addf("if !(%v == %v) {", c.Ref(x), r.InPrms[1].Name)
	r.Add("t.Helper()")
	r.Addf("t.Fatal(append(msgs, fmt.Sprintf(\"should equal %v (actual: %%v)\", %v))...)", c.Ref(x), r.InPrms[1].Name)
	r.Add("}")
	return r
}
func (x *FleTst) CnstNot(typ Typ, c *Cnst) (r *PkgFn) {
	r = x.PkgFnf("%vNot%v", typ.PrefixTitle(), c.Name)
	r.InPrm(typ, "a")
	r.InPrmVariadic(Interface, "msgs")
	r.Addf("if !(%v != %v) {", c.Ref(x), r.InPrms[1].Name)
	r.Add("t.Helper()")
	r.Addf("t.Fatal(append(msgs, \"should not equal %v\")...)", c.Ref(x))
	r.Add("}")
	return r
}
func (x *FleTst) GenTyp(typ Typ) {
	switch X := typ.(type) {
	case *Struct:
		x.StructEql(X)
		x.StructNotZero(X)
	case *Ifc:
		x.IfcEql(X)
		x.IfcNotZero(X)
	case *Arr:
		x.ArrEql(typ)
		x.ArrNotZero(typ)
	default:
		x.Eql(typ) // bsc, ext, other
		if typ.Bse().IsRel() {
			x.Rel("Lss", "<", typ)
			x.Rel("Gtr", ">", typ)
			x.Rel("Leq", "<=", typ)
			x.Rel("Geq", ">=", typ)
		}
	}
	x.SliceEql(typ) // all typs get slice check fro variadic usage
}
func (x *FleTst) StructNotZero(s *Struct) (r *PkgFn) {
	x.Import("fmt")
	r = x.PkgFnf("%vNotZero", s.PrefixTitle())
	s.TstNotZero = r
	r.InPrm(s, "a")
	r.InPrmVariadic(Interface, "msgs")
	if s.IsPtr() {
		r.Add("if a == nil {")
		r.Add("t.Helper()")
		r.Addf("t.Fatal(append(msgs, \"is nil\")...)")
		r.Add("}")
	}
	flds := s.SelFlds(func(f *Fld) bool { // avoid checking variadic slice which zero is valid
		return f.IsFstUpr() && !f.IsTstSkp() && !f.IsTstZeroSkp() && !f.IsSlice() && f.Typ.Bse().TstEql != nil
	})
	for _, f := range *flds {
		r.Addf("%vNotZero(t, a.%v, append(msgs, \"%v.%v\"))", f.Typ.PrefixTitle(), f.Name, s.Title(), f.Name)
	}
	// flds := append(Flds{}, s.Flds...)
	// for len(flds) != 0 {
	// 	f := flds.Dque()
	// 	if f.Name == "" {
	// 		fStrct, ok := f.Typ.(*Struct)
	// 		if ok {
	// 			flds.Ins(0, fStrct.Flds...)
	// 		}
	// 	} else if f.IsFstUpr() && !f.IsTstSkp() && !f.IsTstZeroSkp() && !f.IsSlice() && f.Typ.Bse().TstEql != nil {
	// 		// avoid checking variadic slice which zero is valid
	// 		r.Addf("%vNotZero(t, a.%v, append(msgs, \"%v.%v\"))", f.Typ.PrefixTitle(), f.Name, s.Title(), f.Name)
	// 	}
	// }
	return r
}
func (x *FleTst) ArrNotZero(typ Typ) (r *PkgFn) {
	x.Import("fmt")
	r = x.PkgFnf("%vNotZero", typ.PrefixTitle())
	typ.Bse().TstNotZero = r
	r.InPrm(typ, "a")
	r.InPrmVariadic(Interface, "msgs")
	r.Add("if a == nil {")
	r.Add("t.Helper()")
	r.Addf("t.Fatal(append(msgs, \"is nil\")...)")
	r.Add("}")
	r.Add("if a.Cnt() == 0 {")
	r.Add("t.Helper()")
	r.Add("t.Fatal(append(msgs, fmt.Errorf(\"cnt is zero\"))...)")
	r.Add("}")
	return r
}
func (x *FleTst) StructEql(s *Struct) (r *PkgFn) {
	r = x.PkgFnf("%vEql", s.PrefixTitle())
	s.TstEql = r
	r.InPrm(s, "e")
	r.InPrm(s, "a")
	r.InPrmVariadic(Interface, "msgs")
	x.Import("fmt")
	if s.IsPtr() {
		r.Add("if e == nil && a == nil {")
		r.Add("return")
		r.Add("}")
		r.Add("if e == nil {")
		r.Add("t.Helper()")
		r.Add("t.Fatal(append(msgs, fmt.Errorf(\"e is nil\"))...)")
		r.Add("}")
		r.Add("if a == nil {")
		r.Add("t.Helper()")
		r.Add("t.Fatal(append(msgs, fmt.Errorf(\"a is nil\"))...)")
		r.Add("}")
	}
	flds := append(Flds{}, s.Flds...)
	for len(flds) != 0 {
		f := flds.Dque()
		if f.Name == "" {
			fStrct, ok := f.Typ.(*Struct)
			if ok {
				flds.Ins(0, fStrct.Flds...)
			}
		} else if f.IsFstUpr() && !f.IsTstSkp() {
			var tst *PkgFn
			if f.IsSlice() {
				tst = f.Typ.Bse().TstSliceEql
			} else {
				tst = f.Typ.Bse().TstEql
			}
			if tst != nil {
				r.Addf("%v(t, e.%v, a.%v, append(msgs, \"%v.%v\"))", tst.Ref(x), f.Name, f.Name, s.Title(), f.Name)
			}
		}
	}
	if s.IsAna() && s != _sys.Ana.Prfm.Typ() {
		r.Add("if e.String() != a.String() {")
		r.Add("t.Helper()")
		r.Add("t.Fatal(append(msgs, fmt.Sprintf(\"should equal (expected:%v actual:%v)\", e, a))...)")
		r.Add("}")
	}
	return r
}
func (x *FleTst) ArrEql(typ Typ) (r *PkgFn) {
	arr := typ.Bse().PrtArr().Arr
	x.Import("fmt")
	r = x.PkgFnf("%vEql", typ.PrefixTitle())
	typ.Bse().TstEql = r
	r.InPrm(typ, "e")
	r.InPrm(typ, "a")
	r.InPrmVariadic(Interface, "msgs")
	r.Add("if e == nil && a == nil {")
	r.Add("return")
	r.Add("}")
	r.Add("if e == nil {")
	r.Add("t.Helper()")
	r.Add("t.Fatal(append(msgs, fmt.Errorf(\"e is nil\"))...)")
	r.Add("}")
	r.Add("if a == nil {")
	r.Add("t.Helper()")
	r.Add("t.Fatal(append(msgs, fmt.Errorf(\"a is nil\"))...)")
	r.Add("}")
	r.Add("if len(*e) != len(*a) {")
	r.Add("t.Helper()")
	r.Add("t.Fatal(append(msgs, fmt.Sprintf(\"length not equal (expected:%v actual:%v)\", len(*e), len(*a)))...)")
	r.Add("}")
	elmBse := arr.Elm
	_, isIfc := arr.Alias.Elm.(*Ifc)
	_, isStruct := arr.Alias.Elm.(*Struct)
	r.Add("for n := 0; n < len(*e); n++ {")
	if isIfc || (isStruct && arr.Elm.IsPtr()) {
		r.Add("if (*e)[n] == nil && (*a)[n] == nil {")
		r.Add("continue")
		r.Add("}")
		r.Add("if (*e)[n] == nil {")
		r.Add("t.Helper()")
		r.Add("t.Fatal(append(msgs, fmt.Sprintf(\"e elm is nil: idx %v nil\", n))...)")
		r.Add("}")
		r.Add("if (*a)[n] == nil {")
		r.Add("t.Helper()")
		r.Add("t.Fatal(append(msgs, fmt.Sprintf(\"a elm is nil: idx %v nil\", n))...)")
		r.Add("}")
	}
	r.Addf("%v(t, (*e)[n], (*a)[n], append(msgs, fmt.Sprintf(\"elm %v (idx:%%v expected:%%v actual:%%v)\", n, (*e)[n], (*a)[n])))", elmBse.TstEql.Ref(x), elmBse.Title())
	r.Add("}")
	return r
}
func (x *FleTst) SliceEql(typ Typ) (r *PkgFn) {
	bse := typ.Bse()
	x.Import("fmt")
	r = x.PkgFnf("%vSliceEql", typ.PrefixTitle())
	bse.TstSliceEql = r
	r.InPrmSlice(typ, "e")
	r.InPrmSlice(typ, "a")
	r.InPrmVariadic(Interface, "msgs")
	r.Add("if len(e) != len(a) {")
	r.Add("t.Helper()")
	r.Add("t.Fatal(append(msgs, fmt.Sprintf(\"length not equal (expected:%v actual:%v)\", len(e), len(a)))...)")
	r.Add("}")
	r.Add("for n := 0; n < len(e); n++ {")
	_, isIfc := typ.(*Ifc)
	_, isStruct := typ.(*Struct)
	if isIfc || (isStruct && bse.IsPtr()) {
		r.Add("if e[n] == nil && a[n] == nil {")
		r.Add("continue")
		r.Add("}")
		r.Add("if e[n] == nil {")
		r.Add("t.Helper()")
		r.Add("t.Fatal(append(msgs, fmt.Sprintf(\"e elm is nil: idx %v nil\", n))...)")
		r.Add("}")
		r.Add("if a[n] == nil {")
		r.Add("t.Helper()")
		r.Add("t.Fatal(append(msgs, fmt.Sprintf(\"a elm is nil: idx %v nil\", n))...)")
		r.Add("}")
	}
	r.Addf("%v(t, e[n], a[n], append(msgs, fmt.Sprintf(\"elm %v (idx:%%v expected:%%v actual:%%v)\", n, e[n], a[n])))", bse.TstEql.Ref(x), bse.Title())
	r.Add("}")
	return r
}
func (x *FleTst) IfcEql(ifc *Ifc) (r *PkgFn) {
	x.Import("fmt")
	x.Import("reflect")
	r = x.PkgFnf("%vEql", ifc.PrefixTitle())
	ifc.TstEql = r
	r.InPrm(ifc, "e")
	r.InPrm(ifc, "a")
	r.InPrmVariadic(Interface, "msgs")
	r.Add("if (e == nil || reflect.ValueOf(e).IsNil()) && (a == nil || reflect.ValueOf(a).IsNil()) {")
	r.Add("return")
	r.Add("}")
	ifcs := Ifcs{ifc}
	for len(ifcs) != 0 {
		i := ifcs.Dque()
		ifcs.Ins(0, i.Ifcs...)
		for _, sig := range i.MemSigs {
			if sig.Name != "" &&
				sig.IsFstUpr() &&
				!sig.IsTstSkp() &&
				len(sig.OutPrms) == 1 &&
				sig.OutTyp().Bse().TstEql != nil &&
				(sig.InPrms.Cnt() == 0 || (sig.InPrms.Cnt() == 1 && sig.InPrms[0].IsVariadic())) {
				r.Addf("%v(t, e.%v(), a.%v(), append(msgs, \"%v.%v()\"))", sig.OutTyp().Bse().TstEql.Ref(x), sig.Name, sig.Name, i.Title(), sig.Name)

			}
		}
	}
	return r
}
func (x *FleTst) IfcNotZero(ifc *Ifc) (r *PkgFn) {
	x.Import("fmt")
	x.Import("reflect")
	r = x.PkgFnf("%vNotZero", ifc.PrefixTitle())
	ifc.TstNotZero = r
	r.InPrm(ifc, "a")
	r.InPrmVariadic(Interface, "msgs")
	r.Add("if a == nil || reflect.ValueOf(a).IsNil() {")
	r.Add("t.Helper()")
	r.Addf("t.Fatal(append(msgs, \"is nil\")...)")
	r.Add("}")
	return r
}
func (x *FleTst) Eql(typ Typ) (r *PkgFn) {
	bse := typ.Bse()
	ext, _ := typ.(*Ext)
	x.Import("fmt")
	r = x.PkgFnf("%vEql", typ.PrefixTitle())
	bse.TstEql = r
	r.InPrm(typ, "e")
	r.InPrm(typ, "a")
	r.InPrmVariadic(Interface, "msgs")
	if bse.IsBsc() || ext != nil {
		var cnd string
		if typ == _sys.Bsc.Flt.Typ() {
			cnd = "&& !(e.IsNaN() && a.IsNaN())"
		}
		r.Addf("if e != a %v {", cnd)
		r.Add("t.Helper()")
		r.Add("t.Fatal(append(msgs, fmt.Sprintf(\"should be equal (expected:%v actual:%v)\", e, a))...)")
		r.Add("}")
	}
	return r
}
func (x *FleTst) Rel(name, op string, typ Typ) (r *PkgFn) {
	// bse := typ.Bse()
	x.Import("fmt")
	r = x.PkgFnf("%v%v", typ.PrefixTitle(), strings.Title(name))
	// bse.TstEql = r
	r.InPrm(typ, "e")
	r.InPrm(typ, "a")
	r.InPrmVariadic(Interface, "msgs")
	r.Addf("if !(e %v a) {", op)
	r.Add("t.Helper()")
	r.Addf("t.Fatal(append(msgs, fmt.Sprintf(\"should be %v (expected:%%v actual:%%v)\", e, a))...)", name)
	r.Add("}")
	return r
}

func (x *FleTst) Srt(typ Typ, cmp *PkgFn, fld ...*Fld) (r *PkgFn) {
	var fldName string
	if len(fld) != 0 {
		fldName = strings.Title(fld[0].Name)
	}
	x.Import(_sys.Bsc.Unt.Pkg.Pth)
	r = x.PkgFnf("%v%v%v", typ.PrefixTitle(), strings.Title(cmp.Alias), fldName)
	r.InPrm(typ, "a")
	r.InPrmVariadic(Interface, "msgs")
	r.Add("if a == nil {")
	r.Add("t.Helper()")
	r.Addf("t.Fatal(append(msgs, \"actual is nil\")...)")
	r.Add("}")
	r.Add("for n := 1; n < len(*a); n++ {")
	r.Addf("if !%v((*a)[n-1], (*a)[n]) && (*a)[n-1] != (*a)[n] {", cmp.Ref(x))
	r.Add("t.Helper()")
	r.Addf("t.Fatal(append(msgs, fmt.Errorf(\"not in %v order (idxs %%v,%%v vals %%v,%%v)\", n-1, n, (*a)[n-1], (*a)[n]))...)", cmp.Alias)
	r.Add("}")
	r.Add("}")
	return r
}

func (x *FleTst) NodeFuncIfc(fn *TypFn) (r *TstNodeFunc) {
	// sys.Log("fn", fn.Name)
	// if fn.Name == "LongMl" {
	// 	sys.Log("-")
	// }
	var b strings.Builder
	b.WriteString(fn.Rxr.Typ.Bse().Pkg.Title())
	b.WriteString(strings.Replace(fn.Rxr.Typ.Bse().Title(), "Bse", "", 1))
	b.WriteString(fn.Family)
	// sys.Log("fn", fn.Name, "key", b.String())
	var ok bool
	r, ok = x.NodeFuncs[b.String()]
	if ok {
		return r
	}
	r = &TstNodeFunc{}
	r.Func = x.Func(b.String(), atr.None)
	if fn.Rxr.Typ.Bse().ifc != nil {
		r.Func.InPrm(fn.Rxr.Typ.Bse().ifc, fn.Rxr.Name)
	} else {
		r.Func.InPrm(fn.Rxr.Typ, fn.Rxr.Name)
	}
	// sys.Log("fn", fn.Name)
	r.Func.AddInPrm(fn.In()...)
	r.Func.AddOutPrm(fn.Out()...)
	// func: fn: String()
	fn0 := x.TypFn(k.String, r.Func)
	fn0.OutPrm(String)
	fn0.Addf("return %q", r.Func.Name)

	x.NodeFuncs[b.String()] = r
	return r
}

func (x *FleTst) NodeVar(fn *TypFn, f *TstNodeFunc) (r *Var) {
	// TODO: SET PRNT
	// f.Func.Addf("r := %v", fn.Call())
	// f.Func.Addf("r.(%v.I%[2]v).%[2]vSet(x)", f.Func.InPrms[0].Typ.Bse().Pkg.Name, f.Func.InPrms[0].Typ.Title())
	f.Func.Addf("return %v", fn.Call())
	// if strings.HasSuffix(f.Func.Name, fn.Title()) {
	// 	r = x.Var(f.Func.Name, f.Func.Decl(x))
	// } else {
	// 	r = x.Var(f.Func.Name+fn.Title(), f.Func.Decl(x))
	// }
	fnTitle := fn.Title()
	if fn.Cnj != "" {
		fnTitle = fnTitle[len(fn.Cnj):] // decnj; UnaPos, AggSum...
	}
	r = x.Var(f.Func.Name+fnTitle, f.Func.Decl(x))
	r.FnCall = true
	f.Func.Lines = nil // clear tmp lines
	f.Vars = append(f.Vars, r)
	return r
}
func (x *FleTst) GenNodeVarSlices() {
	var b strings.Builder
	for name, nf := range x.NodeFuncs {
		b.Reset()
		b.WriteString("[]")
		b.WriteString(name)
		b.WriteRune('{')
		for n, v := range nf.Vars {
			if n != 0 {
				b.WriteRune(',')
			}
			b.WriteString(v.Name)
		}
		b.WriteRune('}')
		r := x.Var(name+"s", b.String())
		r.FnCall = true
	}
}
