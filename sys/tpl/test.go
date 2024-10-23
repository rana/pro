package tpl

import (
	"fmt"
	"strings"
	"sys"
	"sys/k"
	"unicode/utf8"
)

type (
	DirTest struct {
		DirBse
	}
)

func (x *DirSys) NewTest() (r *DirTest) {
	r = &DirTest{}
	x.Test = r
	r.Pkg = x.Pkg.New(k.Test)
	return r
}

type (
	FleTest struct {
		FleBse
		Src       *FleBse
		SliceVars Vars
	}
	FnT struct { // for test generation: expected := ...
		Lines
		RxrOnly  bool
		Empty    bool
		Manual   bool
		SkpCpy   bool
		SkpExcpy bool
		SkpTst   bool
		TstCnd   string
	}
	FnT2 struct {
		MdlFst Lines
		MdlLst Lines
	}
)

func (x *FleTest) InitTypFn() {
	if Opt.IsTest() {
		if x.Src.Typs.Ok() {
			bse := x.Src.Typ().Bse()
			for _, fn := range x.Src.PkgFns {
				if len(fn.T.Lines) != 0 {
					x.PkgFn(fn, x.Src.Typ().PkgTypTitle())
				}
			}
			// TODO: MOVE TO TypBse.TypFns?
			for _, fn := range x.Src.TypFns {
				if len(fn.T.Lines) != 0 && x.Src.Typ() != nil && !x.Src.Typ().Bse().IsAna() {
					x.TypFn(fn)
				}
			}
			if bse.IsArr() {
				if bse.IsBytWrt() && Opt.Is(ArrOpt) { // TODO: MOVE TO BETTER PLACE?
					x.BytArr()
				}
				if x.PkgFns.Ok() {
					x.Import(x.Src)
					// x.Import(_sys.Bsc.Unt)
				}
			} else {
				if bse.IsBytWrt() {
					x.Byt()
				}
				if x.PkgFns.Ok() {
					x.Import(x.Src)
				}
			}
		} else { // test file with no typ associated
			for _, fn := range x.Src.PkgFns {
				if len(fn.T.Lines) != 0 {
					x.PkgFn(fn, x.Src.Title())
				}
			}
		}
	}
}
func (x *FleTest) WriteToDisk(dir string) {
	if Opt.IsTest() {
		if x.PkgFns.Cnt() != 0 {
			x.Import("testing")
			x.Import("sys/tst")
		}
		x.FleBse.WriteToDisk(dir)
	}
}

func (x *FleTest) TestFn(name string) (r *PkgFn) {
	r = x.PkgFnf("Test%v", strings.Title(name))
	r.InPrmPtr(T, "t")
	return r
}
func (x *FleTest) TestFnf(format string, args ...interface{}) (r *PkgFn) {
	r = x.PkgFnf("Test%v", strings.Title(fmt.Sprintf(format, args...)))
	r.InPrmPtr(T, "t")
	return r
}

// InitCnst
func (x *FleTest) InitCnst() {
	if x.Src.Cnsts.MayXpr() {
		x.Import("fmt")
		x.Import(x.Src.Pkg.Pth)
		r := x.TestFnf("%vCnst", x.Src.Typ().PkgTypTitle())
		r.Add("cses := []struct{")
		r.Add("idn  string")
		r.Addf("e  %v", x.Src.Typ().Ref(x))
		r.Addf("a  %v", x.Src.Typ().Ref(x))
		r.Add("}{")
		for _, c := range x.Src.Cnsts {
			if x.Src.Typ() == c.Typ {
				if _, ok := x.Src.Typ().(*Struct); ok {
					r.Addf("{%q, %v{%v}, %v},", c.Name, x.Src.Typ().Ref(x), c.Value, c.Ref(x))
				} else {
					r.Addf("{%q, %v(%v), %v},", c.Name, x.Src.Typ().Ref(x), c.Value, c.Ref(x))
				}
			}
		}
		r.Add("}")
		r.Add("for _, cse := range cses {")
		r.Add("t.Run(fmt.Sprintf(\"%q\", cse.idn), func(t *testing.T) {")
		r.Addf("tst.%vEql(t, cse.e, cse.a)", x.Src.Typ().Title())
		r.Add("})")
		r.Add("}")
	}
}

func (x *FleTest) PkgFn(fn *PkgFn, title string) (r *PkgFn) {
	b := &strings.Builder{}
	r = x.TestFnf("%v%vPkgFn", title, fn.Title())
	if fn.T.Empty {
		r.Add(fn.T.Lines...)
		return r
	}
	r.Add("cses := []struct{")
	for _, prm := range fn.InPrms {
		r.Add(prm.TestFld(x, b))
	}
	r.Add("}{")
	WriteCses(nil, &fn.InPrms, r, _sys, b)
	r.Add("}")
	r.Add("for _, cse := range cses {")
	r.Addf("t.Run(%q, func(t *testing.T) {", fn.Title())
	fn.InPrms.TestFldAsn(r, b)
	r.Addf("actual := %v", fn.TestCall(x, b))
	r.Add(fn.T.Lines...)
	r.Addf("tst.%vEql(t, expected, actual)", fn.OutTyp().Title())
	r.Add("})")
	r.Add("}")
	return r
}

func (x *FleTest) ImportRecurse(t Typ, addedTyps map[Typ]interface{}) {
	var typs []Typ
	typs = append(typs, t)
	for len(typs) > 0 { // recursively import types and struct field typs
		curTyp := typs[len(typs)-1]
		typs = typs[:len(typs)-1]
		_, imported := addedTyps[curTyp]
		if !imported && !curTyp.Bse().IsLitSkp() {
			x.Import(curTyp)        // import typ
			addedTyps[curTyp] = nil // mark as added
			prtArr := curTyp.Bse().PrtArr()
			if prtArr != nil {
				typs = append(typs, prtArr.Arr.Alias.Elm)
			} else if elmStruct, ok := curTyp.(*Struct); ok {
				for _, fld := range elmStruct.Flds {
					if !fld.IsLitSkp() && fld.Name != "" {
						typs = append(typs, fld.Typ)
					}
				}
			}
		}
	}
}

// TypFn
func (x *FleTest) TypFn(fn *TypFn) (r *PkgFn) {
	b := &strings.Builder{}
	bse := fn.Rxr.Typ.Bse()
	r = x.TestFnf("%v%vTypFn", x.Src.Typ().PkgTypTitle(), fn.Title())
	if fn.T.Empty {
		r.Add(fn.T.Lines...)
		return r
	}
	if !bse.IsBse() && (bse.IsStruct() || (x.Src.Typ().Bse().IsArr() && x.Src.Typ().Bse().PrtArr().Arr.Elm.IsStruct())) {
		// THIS SECTION IS PROBLEMATIC TO GET RIGHT FOR ALL CASES
		// TRY ADDING atr.LitSkp FOR PROBLEM FIELDS
		var s *Struct
		if bse.IsStruct() {
			s = x.Src.Typ().(*Struct) // for ana.Instr, Perf/Trd
		} else {
			s = x.Src.Typ().Bse().PrtArr().Arr.Alias.Elm.(*Struct) // for Perfs/Trds
		}
		for _, fld := range s.Flds {
			if !fld.IsLitSkp() { // for Perf.Pth
				x.Import(fld.Typ)
			}
		}
	}
	r.Add("cses := []struct{")
	if bse.IsAna() && !bse.IsArr() && len(bse.Lits) != 0 {
		r.Add("pth string")
	}
	r.Add(fn.Rxr.TestFld(x, b))
	if !fn.T.RxrOnly {
		fn.InPrms.TestFld(x, r, b)
	}
	r.Add("}{")
	for n, _ := range bse.Vals {
		rxrVal := bse.ValAt(n)
		r.AddStrt("{")
		if bse.IsAna() && !bse.IsArr() && len(bse.Lits) != 0 { // Perf does not have Lits
			r.AddMdlf("%q,", bse.Lits[n])
		}
		r.AddMdlf("%v", rxrVal)
		if !fn.T.RxrOnly && len(fn.InPrms) != 0 {
			// addedTyps := make(map[Typ]interface{})
			for _, prm := range fn.InPrms {
				// import slice elem field types
				// ImportRecurse FOR: HstTrdStmSeg, HstTrdStmSegs
				// TRICKY TO IMPLEMENT EXACT IMPORT STATEMENT; AVOID IF POSSIBLE
				// x.ImportRecurse(prm.Typ, addedTyps)
				r.AddMdl(",")
				if prm.Vals != nil {
					r.AddMdl(prm.Vals[0])
				} else {
					val := prm.Typ.Bse().Val()
					if prm.Typ.Bse().IsAna() && fn.Rxr.Typ.Bse().IsBse() { // Stm, Cnd...
						val = sys.CpyInstr(rxrVal, val) // allow RltStm prm to have same instr for testing
					}
					r.AddMdl(val)
				}
			}
		}
		r.AddEnd("},")
	}
	r.Add("}")
	r.Add("for _, cse := range cses {")
	r.Addf("t.Run(%q, func(t *testing.T) {", fn.Title())
	if fn.Rxr.Typ.Bse().IsAna() {
		x.Import("sys/app")
		r.Add("ap := app.New(tst.Cfg)")
		r.Add("ap.Oan.CloneInstrs(tst.Instrs)")
	}
	r.Add(fn.Rxr.TestFldAsn(b))
	if !fn.T.RxrOnly {
		fn.InPrms.TestFldAsn(r, b)
	}
	if fn.T.TstCnd != "" {
		r.Addf("if %v {", fn.T.TstCnd)
	}
	if fn.T.Manual {
		r.Add(fn.T.Lines...)
	} else {
		outBse := fn.OutTyp().Bse()
		if fn.Rxr.Typ.Bse().IsArr() && !fn.T.SkpCpy {
			if !fn.T.SkpExcpy {
				r.Add("exCpy := x.Cpy()")
			}
			r.Add("axCpy := x.Cpy()")
			r.Addf("actual := %v", fn.TestCall(b, "axCpy"))
		} else {
			var cast string
			if fn.OutPrms.Fst().Tst != nil && !outBse.IsArr() {
				cast = fmt.Sprintf(".(%v)", fn.OutPrms.Fst().Tst.Ref(x))
			}
			r.Addf("actual := %v%v", fn.TestCall(b), cast)
		}
		r.Add(fn.T.Lines...)
		if fn.OutTyp() != nil && !fn.T.SkpTst {
			if fn.OutPrms.Fst().Tst != nil {
				r.Addf("%v(t, expected, actual)", fn.OutPrms.Fst().Tst.Bse().TstEql.Ref(x))
			} else {
				r.Addf("%v(t, expected, actual)", fn.OutTyp().Bse().TstEql.Ref(x))
			}
		}
	}
	if fn.T.TstCnd != "" {
		r.Add("}")
	}
	if fn.Rxr.Typ.Bse().IsAna() {
		r.Add("ap.Cls()")
	}
	r.Add("})")
	r.Add("}")
	return r
}

func (x *FleTest) Byt() (r *PkgFn) {
	b := &strings.Builder{}
	r = x.TestFnf("%vByt", x.Src.Typ().PkgTypTitle())
	r.Add("cses := []struct{")
	r.Addf("e  %v", x.Src.Typ().Ref(x))
	r.Add("}{")
	WriteCses(x.Src.Typ(), nil, r, _sys, b)
	r.Add("}")
	r.Add("for _, cse := range cses {")
	r.Add("t.Run(\"Byt\", func(t *testing.T) {")
	r.Addf("b := %v{}", BufferPtr.Adr(x))
	r.Add("cse.e.BytWrt(b)")
	r.Addf("var a %v", x.Src.Typ().Full())
	r.Add("a.BytRed(b.Bytes())")
	r.Addf("%v(t, cse.e, a)", x.Src.Typ().Bse().TstEql.Ref(x))
	r.Add("})")
	r.Add("}")
	return r
}
func (x *FleTest) BytArr() (r *PkgFn) {
	x.Import("bytes")
	b := &strings.Builder{}
	r = x.TestFnf("%vByt", x.Src.Typ().PkgTypTitle())
	r.Add("cses := []struct{")
	r.Addf("e  %v", x.Src.Typ().Ref(x))
	r.Add("}{")
	WriteCses(x.Src.Typ(), nil, r, _sys, b)
	r.Add("}")
	r.Add("for _, cse := range cses {")
	r.Add("t.Run(\"Byt\", func(t *testing.T) {")
	r.Addf("b := %v{}", BufferPtr.Adr(x))
	r.Add("cse.e.BytWrt(b)")
	r.Addf("a := %v()", x.Src.NewFull())
	r.Add("a.BytRed(b.Bytes())")
	r.Addf("%v(t, cse.e, a)", x.Src.Typ().Bse().TstEql.Ref(x))
	r.Add("})")
	r.Add("}")
	return r
}

func WriteCseStruct(inPrms *InPrms, fn *PkgFn, f Fle, b *strings.Builder) {
	if inPrms != nil && len(*inPrms) != 0 {
		ch := 'a'
		for _, prm := range *inPrms {
			b.Reset()
			b.WriteRune(ch)
			b.WriteRune(' ')
			if prm.Typ.Bse().IsAna() {
				prm.Typ.Bse().FnValSig(f, b, prm)
			} else {
				b.WriteString(prm.Typ.Ref(f))
			}
			fn.Add(b.String())
			ch++
		}
		b.Reset()
	}
}
func WriteCsePrms(inPrms *InPrms, b *strings.Builder) {
	if inPrms != nil && len(*inPrms) != 0 {
		ch := 'a'
		for n := range *inPrms {
			if n != 0 {
				b.WriteRune(',')
			}
			b.WriteString("cse.")
			b.WriteRune(ch)
			ch++
		}
	}
}
func WriteCses(rxr Typ, inPrms *InPrms, fn *PkgFn, s *DirSys, b *strings.Builder) {
	var typs []Typ
	var valss [][]string
	if rxr != nil {
		typs = append(typs, rxr)
		if rxr.Bse().IsAna() {
			valss = append(valss, rxr.Bse().FnVals())
		} else {
			valss = append(valss, rxr.Bse().Vals)
		}
	}
	if inPrms != nil {
		for _, inPrm := range *inPrms {
			typs = append(typs, inPrm.Typ)
			if inPrm.Vals != nil {
				valss = append(valss, inPrm.Vals)
				// TODO: SUPPORT FnVals FOR PRM?
			} else {
				if inPrm.Typ.Bse().IsAna() {
					valss = append(valss, inPrm.Typ.Bse().FnVals())
				} else {
					valss = append(valss, inPrm.Typ.Bse().Vals)
				}
			}
		}
	}
	if len(typs) != 0 {
		for _, permIdxs := range sys.PermIdxsStrs(valss...) {
			b.Reset()
			b.WriteRune('{')
			for n, permIdx := range permIdxs {
				if n != 0 {
					b.WriteRune(',')
				}
				b.WriteString(valss[n][permIdx])
			}
			b.WriteString("},")
			fn.Add(b.String())
		}
		b.Reset()
	}
}

func (x *FleTest) LitTrm(trm Typ, lits []string, gens ...Gen) {
	x.Trm(trm.Title(), trm.Ref(x), lits, gens...)
}
func (x *FleTest) LitTrmTyp(bse *TypBse, jsn ...bool) {
	gens := make([]Gen, 2)
	switch {
	case bse.Atr.IsStr():
		gens[0] = TxtGen
		gens[1] = TxtGen | PrefixGen | SuffixGen
	case bse.Atr.IsNum() || bse.Atr.IsRng():
		gens[0] = NumGen
		gens[1] = NumGen | PrefixGen | SuffixGen
	case bse.Atr.IsStruct():
		gens[0] = TxtGen
		gens[1] = TxtGen | PrefixGen

	}
	var trm *Struct
	var lits []string
	if len(jsn) == 0 {
		trm, lits = bse.LitTrm, bse.Lits
	} else {
		trm, lits = bse.LitTrmJsn, bse.LitsJsn
	}
	x.LitTrm(trm, lits, gens...)
}

func (x *FleTest) Trm(title, typRef string, lits []string, gens ...Gen) {
	// x.Import(_sys.Bsc.Bnd.Pkg.Pth)
	title = strings.Title(title)
	var valid Gen
	if len(gens) > 0 {
		valid = gens[0]
	} else {
		valid = TxtGen | SuffixGen
	}
	x.TrmValid(title, typRef, lits, valid)
	var invalid Gen
	if len(gens) > 1 {
		invalid = gens[1]
	} else {
		invalid = TxtGen | PrefixGen | SuffixGen // default to Txt prefix/suffix
	}
	if invalid&NumGen == NumGen {
		invalid |= PrefixGen | SuffixGen // numbers get prefix/suffix
	}
	x.TrmInvalid(title, lits, invalid)
}
func (x *FleTest) TrmValid(title, typRef string, lits []string, gen Gen) (r *PkgFn) {
	x.Import("fmt")
	x.Import(_sys.Bsc.Unt)
	r = x.TestFnf("%vValid", title)
	r.Add("cses := []struct{")
	r.Add("txt  string")
	r.Add("lim  unt.Unt")
	r.Add("}{")
	for _, lit := range lits {
		lim := len(lit)
		r.Addf("{%q, %v},", lit, lim)
		if gen&SuffixGen == SuffixGen {
			for _, edge := range NonIdnEdges {
				r.Addf("{%q, %v},", lit+edge, lim)
			}
		}
	}
	r.Add("}")
	r.Add("var trmr trm.Trmr")
	r.Add("for _, cse := range cses {")
	r.Add("lbl := cse.txt")
	r.Add("if len(lbl) > 16 {")
	r.Add("lbl = lbl[:16]")
	r.Add("}")
	r.Add("t.Run(fmt.Sprintf(\"%q\", lbl), func(t *testing.T) {")
	r.Add("trmr.Reset(cse.txt)")
	r.Addf("a, ok := trmr.%v()", title)
	r.Add("tst.True(t, ok, \"Lex\")")
	r.Addf("tst.TypeEql(t, %v{}, a)", typRef)
	r.Add("tst.UntEql(t, 0, a.Idx, \"Idx\")")
	r.Add("tst.UntEql(t, cse.lim, a.Lim, \"Lim\")")
	r.Add("})")
	r.Add("}")
	return r
}
func (x *FleTest) TrmInvalid(title string, lits []string, gen Gen) (r *PkgFn) {
	var edges []string
	if gen&IdnGen == IdnGen {
		edges = NonIdnEdges
	} else if gen&TxtGen == TxtGen {
		edges = FalseEdgesTxt
	} else {
		edges = FalseEdgesNum
	}
	r = x.TestFnf("%vInvalid", title)
	r.Add("cses := []struct{")
	r.Add("txt  string")
	r.Add("}{")
	for _, lit := range lits {
		for _, edge := range edges {
			for lim := 1; lim < utf8.RuneCountInString(lit); lim++ {
				if lim > 16 { // doesn't need to be too long
					break
				}
				prt := lit[:lim]
				if gen&PrefixGen == PrefixGen {
					r.Addf("{%q},", edge+prt)
				}
				if gen&SuffixGen == SuffixGen {
					r.Addf("{%q},", prt+edge)
				}
			}
			if gen&PrefixGen == PrefixGen {
				r.Addf("{%q},", edge+lit)
			}
			if gen&SuffixGen == SuffixGen {
				r.Addf("{%q},", lit+edge)
			}
		}
	}
	r.Add("}")
	r.Add("var trmr trm.Trmr")
	r.Add("for _, cse := range cses {")
	r.Add("lbl := cse.txt")
	// r.Add("if len(lbl) > 16 {")
	// r.Add("lbl = lbl[:16]")
	// r.Add("}")
	r.Add("t.Run(fmt.Sprintf(\"%q\", lbl), func(t *testing.T) {")
	r.Add("trmr.Reset(cse.txt)")
	r.Addf("_, ok := trmr.%v()", title)
	r.Add("tst.False(t, ok, \"Lex\")")
	r.Add("})")
	r.Add("}")
	return r
}

func (x *FleTest) PrsTrm(fn *PkgFn, typ, litTyp Typ, jsn ...bool) (r *PkgFn) {
	bse := typ.Bse()
	b := &strings.Builder{}
	r = x.TestFnf("Prs%v", fn.Title())
	if fn.T.Empty {
		r.Add(fn.T.Lines...)
		return r
	}
	r.Add("cses := []struct{")
	r.Add("lit string")
	r.Addf("val %v", typ.Ref(x))
	r.Add("}{")
	var lits, vals []string
	if len(jsn) == 0 {
		lits, vals = bse.LitsNonEmp(), bse.Vals
	} else {
		lits, vals = bse.LitsJsn, bse.ValsJsn
	}
	_, isStruct := typ.(*Struct)
	alias, isAlias := typ.(*Alias)
	var isSlice bool
	if isAlias {
		isSlice = alias.AliasMod.IsSlice()
	}
	for n := 0; n < len(lits); n++ {
		if isSlice && n == 0 { // skip first empty slice
			continue
		}
		if isStruct || isSlice {
			r.Addf("{%q, %v},", lits[n], vals[n])
		} else {
			r.Addf("{%q, %v},", lits[n], vals[n])
		}
	}
	r.Add("}")
	r.Add("for _, cse := range cses {")
	r.Addf("t.Run(%q, func(t *testing.T) {", fn.Title())
	WriteCsePrms(&fn.InPrms, b)
	r.Add("var trmr trm.Trmr")
	r.Add("trmr.Reset(cse.lit)")
	r.Addf("%v, ok := trmr.%v()", litTyp.Camel(), litTyp.Title())
	r.Add("tst.True(t, ok)")
	r.Addf("a := %v(%v, cse.lit)", fn.Ref(x), litTyp.Camel())
	r.Addf("%v(t, cse.val, a)", fn.OutTyp().Bse().TstEql.Ref(x))
	r.Add("})")
	r.Add("}")
	return r
}
func (x *FleTest) PrsTxt(fn *PkgFn, typ Typ, jsn ...bool) (r *PkgFn) {
	bse := typ.Bse()
	b := &strings.Builder{}
	r = x.TestFnf("Prs%v", fn.Title())
	if fn.T.Empty {
		r.Add(fn.T.Lines...)
		return r
	}
	r.Add("cses := []struct{")
	r.Add("lit string")
	r.Addf("val %v", typ.Ref(x))
	r.Add("}{")
	var lits, vals []string
	if len(jsn) == 0 {
		lits, vals = bse.LitsNonEmp(), bse.Vals
	} else {
		lits, vals = bse.LitsJsn, bse.ValsJsn
	}
	_, isStruct := typ.(*Struct)
	alias, isAlias := typ.(*Alias)
	var isSlice bool
	if isAlias {
		isSlice = alias.AliasMod.IsSlice()
	}
	for n := 0; n < len(lits); n++ {
		if isSlice && n == 0 { // skip first empty slice
			continue
		}
		if isStruct || isSlice {
			r.Addf("{%q, %v},", lits[n], vals[n])
		} else {
			r.Addf("{%q, %v},", lits[n], vals[n])
		}
	}
	r.Add("}")
	r.Add("for _, cse := range cses {")
	r.Addf("t.Run(%q, func(t *testing.T) {", fn.Title())
	WriteCsePrms(&fn.InPrms, b)
	r.Addf("a := %v(cse.lit)", fn.Ref(x))
	r.Addf("%v(t, cse.val, a)", fn.OutTyp().Bse().TstEql.Ref(x))
	r.Add("})")
	r.Add("}")
	return r
}

func (x *FleTest) Cfg(fn *TypFn, typ Typ) {
	x.CfgValid(fn, typ)
	x.CfgInvalid(fn, typ)
}
func (x *FleTest) CfgValid(fn *TypFn, typ Typ) (r *PkgFn) {
	bse := typ.Bse()
	x.Import(bse.Pkg.Pth)
	r = x.TestFnf("Cfg%vValid", fn.Title())
	r.Add("cses := []struct{")
	r.Addf("e %v", typ.Ref(x))
	r.Add("pth []string")
	r.Add("txt string")
	r.Add("}{")
	spce, cmnt := " \t\r\n ", "// comment\n"
	for n := 0; n < len(bse.Lits); n++ {
		if bse.Lits[n] == "[]" { // skip empty arr
			continue
		}
		var txt string
		// single flat
		txt = fmt.Sprintf("{ %[2]v%[3]v key %[2]v%[3]v : %[2]v%[3]v %[1]v %[2]v%[3]v }", bse.Lits[n], spce, cmnt)
		r.Addf("{%v, []string{\"key\"}, %q},", bse.Vals[n], txt)
		// multi flat
		txt = fmt.Sprintf("{ no1:\"no\" no2:[\"no\" \"no\"] %[2]v%[3]v key %[2]v%[3]v : %[2]v%[3]v %[1]v %[2]v%[3]v  no3:\"no\" }", bse.Lits[n], spce, cmnt)
		r.Addf("{%v, []string{\"key\"}, %q},", bse.Vals[n], txt)
		// single depth
		txt = fmt.Sprintf("{ %[2]v%[3]v k2 %[2]v%[3]v : %[2]v%[3]v { %[2]v%[3]v k1 %[2]v%[3]v : %[2]v%[3]v { %[2]v%[3]v key %[2]v%[3]v : %[2]v%[3]v %[1]v %[2]v%[3]v } %[2]v%[3]v } %[2]v%[3]v }", bse.Lits[n], spce, cmnt)
		r.Addf("{%v, []string{\"k2\", \"k1\", \"key\"}, %q},", bse.Vals[n], txt)
		// multi depth
		txt = fmt.Sprintf("{ no1:\"no\" no2:[\"no\" \"no\"] %[2]v%[3]v k2 %[2]v%[3]v : %[2]v%[3]v { no3:\"no\" %[2]v%[3]v no4:[\"no\" \"no\"] k1 %[2]v%[3]v : %[2]v%[3]v { no5:\"no\"  no6:[\"no\" \"no\"] %[2]v%[3]v key %[2]v%[3]v : %[2]v%[3]v %[1]v %[2]v%[3]v } %[2]v%[3]v } %[2]v%[3]v }", bse.Lits[n], spce, cmnt)
		r.Addf("{%v, []string{\"k2\", \"k1\", \"key\"}, %q},", bse.Vals[n], txt)
	}
	r.Add("}")
	r.Add("for _, cse := range cses {")
	r.Addf("t.Run(%q, func(t *testing.T) {", fn.Title())
	r.Add("var c cfg.Cfgr")
	r.Add("c.Reset(cse.txt)")
	r.Addf("a := c.%v(cse.pth...)", fn.Name)
	r.Addf("%v(t, cse.e, a)", bse.TstEql.Ref(x))
	r.Add("})")
	r.Add("}")
	return r
}
func (x *FleTest) CfgInvalid(fn *TypFn, typ Typ) (r *PkgFn) {
	bse := typ.Bse()
	x.Import(bse.Pkg.Pth)
	r = x.TestFnf("Cfg%vInvalid", fn.Title())
	r.Add("cses := []struct{")
	r.Addf("e %v", typ.Ref(x))
	r.Add("pth []string")
	r.Add("txt string")
	r.Add("}{")
	spce, cmnt := " \t\r\n ", "// comment\n"
	for n := 0; n < len(bse.Lits); n++ {
		if n == 0 && bse.IsArr() { // skip empty arr
			continue
		}
		var txt string
		// single flat, no key
		txt = fmt.Sprintf("{ %[2]v%[3]v key %[2]v%[3]v : %[2]v%[3]v %[1]v %[2]v%[3]v }", bse.Lits[n], spce, cmnt)
		r.Addf("{%v, []string{}, %q},", bse.Vals[n], txt)
		// single flat, wrong key
		txt = fmt.Sprintf("{ %[2]v%[3]v key %[2]v%[3]v : %[2]v%[3]v %[1]v %[2]v%[3]v }", bse.Lits[n], spce, cmnt)
		r.Addf("{%v, []string{\"wrong\"}, %q},", bse.Vals[n], txt)
		// multi flat, wrong key
		txt = fmt.Sprintf("{ no1:\"no\" no2:[\"no\" \"no\"] %[2]v%[3]v key %[2]v%[3]v : %[2]v%[3]v %[1]v %[2]v%[3]v  no3:\"no\" }", bse.Lits[n], spce, cmnt)
		r.Addf("{%v, []string{\"wrong\"}, %q},", bse.Vals[n], txt)
		// single depth, wrong key
		txt = fmt.Sprintf("{ %[2]v%[3]v k2 %[2]v%[3]v : %[2]v%[3]v { %[2]v%[3]v k1 %[2]v%[3]v : %[2]v%[3]v { %[2]v%[3]v key %[2]v%[3]v : %[2]v%[3]v %[1]v %[2]v%[3]v } %[2]v%[3]v } %[2]v%[3]v }", bse.Lits[n], spce, cmnt)
		r.Addf("{%v, []string{\"k2\", \"k1\", \"wrong\"}, %q},", bse.Vals[n], txt)
		// multi depth, wrong key
		txt = fmt.Sprintf("{ no1:\"no\" no2:[\"no\" \"no\"] %[2]v%[3]v k2 %[2]v%[3]v : %[2]v%[3]v { no3:\"no\" %[2]v%[3]v no4:[\"no\" \"no\"] k1 %[2]v%[3]v : %[2]v%[3]v { no5:\"no\"  no6:[\"no\" \"no\"] %[2]v%[3]v key %[2]v%[3]v : %[2]v%[3]v %[1]v %[2]v%[3]v } %[2]v%[3]v } %[2]v%[3]v }", bse.Lits[n], spce, cmnt)
		r.Addf("{%v, []string{\"k2\", \"k1\", \"wrong\"}, %q},", bse.Vals[n], txt)
	}
	r.Add("}")
	r.Add("var c cfg.Cfgr")
	r.Add("for _, cse := range cses {")
	r.Addf("t.Run(%q, func(t *testing.T) {", fn.Title())
	r.Add("c.Reset(cse.txt)")
	r.Addf("tst.Panic(t, func() { c.%v(cse.pth...) })", fn.Name)
	r.Add("})")
	r.Add("}")
	return r
}

func (x *FleTest) Jsn(fn *TypFn, typ Typ) {
	x.JsnValid(fn, typ)
	// x.JsnInvalid(fn, typ, tst)
}
func (x *FleTest) JsnValid(fn *TypFn, typ Typ) (r *PkgFn) {
	bse := typ.Bse()
	x.ImportTyp(typ)
	r = x.TestFnf("Jsn%vValid", fn.Title())
	r.Add("cses := []struct{")
	r.Addf("e %v", typ.Ref(x))
	r.Add("pth []string")
	r.Add("txt string")
	r.Add("}{")
	spce := " \t\r\n "
	lits, vals := bse.LitsJsn, bse.ValsJsn
	for n := 0; n < len(lits); n++ {
		if n == 0 && bse.IsArr() { // skip empty arr
			continue
		}
		var txt string
		// single flat
		txt = fmt.Sprintf("{ %[2]v \"key\" %[2]v : %[2]v %[1]v %[2]v }", lits[n], spce)
		r.Addf("{%v, []string{\"key\"}, %q},", vals[n], txt)
		// multi flat
		txt = fmt.Sprintf("{ \"no1\":\"no\", \"no2\":[\"no\" \"no\"], %[2]v \"key\" %[2]v : %[2]v %[1]v, %[2]v  \"no3\":\"no\" }", lits[n], spce)
		r.Addf("{%v, []string{\"key\"}, %q},", vals[n], txt)
		// single depth
		txt = fmt.Sprintf("{ %[2]v \"k2\" %[2]v : %[2]v { %[2]v \"k1\" %[2]v : %[2]v { %[2]v \"key\" %[2]v : %[2]v %[1]v %[2]v } %[2]v } %[2]v }", lits[n], spce)
		r.Addf("{%v, []string{\"k2\", \"k1\", \"key\"}, %q},", vals[n], txt)
		// multi depth
		txt = fmt.Sprintf("{ \"no1\":\"no\" \"no2\":[\"no\" \"no\"], %[2]v \"k2\" %[2]v : %[2]v { \"no3\":\"no\", %[2]v \"no4\":[\"no\" \"no\"], \"k1\" %[2]v : %[2]v { \"no5\":\"no\",  \"no6\":[\"no\" \"no\"], %[2]v \"key\" %[2]v : %[2]v %[1]v %[2]v } %[2]v } %[2]v }", lits[n], spce)
		r.Addf("{%v, []string{\"k2\", \"k1\", \"key\"}, %q},", vals[n], txt)
	}
	r.Add("}")
	r.Add("for _, cse := range cses {")
	r.Addf("t.Run(%q, func(t *testing.T) {", fn.Title())
	r.Add("var j jsn.Jsnr")
	r.Add("j.Reset(cse.txt)")
	r.Addf("a := j.%v(cse.pth...)", fn.Name)
	r.Addf("%v(t, cse.e, a)", bse.TstEql.Ref(x))
	r.Add("})")
	r.Add("}")
	return r
}
func (x *FleTest) JsnInvalid(fn *TypFn, typ Typ, tst *FleTst) (r *PkgFn) {
	bse := typ.Bse()
	x.ImportTyp(typ)
	r = x.TestFnf("Jsn%vInvalid", fn.Title())
	r.Add("cses := []struct{")
	r.Addf("e %v", typ.Ref(x))
	r.Add("pth []string")
	r.Add("txt string")
	r.Add("}{")
	spce := " \t\r\n "
	lits, vals := bse.LitsJsn, bse.ValsJsn
	for n := 0; n < len(lits); n++ {
		if n == 0 && bse.IsArr() { // skip empty arr
			continue
		}
		var txt string
		// single flat, no key
		txt = fmt.Sprintf("{ %[2]v \"key\" %[2]v : %[2]v %[1]v %[2]v }", lits[n], spce)
		r.Addf("{%v, []string{}, %q},", vals[n], txt)
		// single flat, wrong key
		txt = fmt.Sprintf("{ %[2]v \"key\" %[2]v : %[2]v %[1]v %[2]v }", lits[n], spce)
		r.Addf("{%v, []string{\"wrong\"}, %q},", vals[n], txt)
		// multi flat, wrong key
		txt = fmt.Sprintf("{ \"no1\":\"no\", \"no2\":[\"no\" \"no\"], %[2]v \"key\" %[2]v : %[2]v %[1]v, %[2]v  \"no3\":\"no\" }", lits[n], spce)
		r.Addf("{%v, []string{\"wrong\"}, %q},", vals[n], txt)
		// single depth, wrong key
		txt = fmt.Sprintf("{ %[2]v \"k2\" %[2]v : %[2]v { %[2]v \"k1\" %[2]v : %[2]v { %[2]v \"key\" %[2]v : %[2]v %[1]v %[2]v } %[2]v } %[2]v }", lits[n], spce)
		r.Addf("{%v, []string{\"k2\", \"k1\", \"wrong\"}, %q},", vals[n], txt)
		// multi depth, wrong key
		txt = fmt.Sprintf("{ \"no1\":\"no\" \"no2\":[\"no\" \"no\"], %[2]v \"k2\" %[2]v : %[2]v { \"no3\":\"no\", %[2]v \"no4\":[\"no\" \"no\"], \"k1\" %[2]v : %[2]v { \"no5\":\"no\",  \"no6\":[\"no\" \"no\"], %[2]v \"key\" %[2]v : %[2]v %[1]v %[2]v } %[2]v } %[2]v }", lits[n], spce)
		r.Addf("{%v, []string{\"k2\", \"k1\", \"wrong\"}, %q},", vals[n], txt)
	}
	r.Add("}")
	r.Add("var j jsn.Jsnr")
	r.Add("for _, cse := range cses {")
	r.Addf("t.Run(%q, func(t *testing.T) {", fn.Title())
	r.Add("j.Reset(cse.txt)")
	r.Addf("tst.Panic(t, func() { j.%v(cse.pth...) })", fn.Name)
	r.Add("})")
	r.Add("}")
	return r
}

func (x *FleTest) XprLit(typ Typ) (r *PkgFn) {
	bse := typ.Bse()
	x.Import("fmt")
	x.Import(_sys.Lng.Pro.Xpr)
	x.Import(_sys.Bsc.Unt)
	r = x.TestFnf("TestXpr%vLit", typ.PkgTypTitle())
	r.Add("cses := []struct{")
	r.Add("lit string")
	r.Add("}{")
	for n, lit := range bse.Lits {
		if n == 0 && bse.IsArr() { // skip empty arr
			continue
		}
		r.Addf("{%q},", lit)
	}
	r.Add("}")
	r.Add("var xprr xpr.Xprr")
	r.Add("for _, cse := range cses {")
	r.Add("t.Run(fmt.Sprintf(\"%q\", cse.lit), func(t *testing.T) {")
	r.Add("_, xprs := xprr.Prs(cse.lit)")
	r.Add("tst.IntegerEql(t, 1, len(xprs), cse.lit, \"xprs count\")")
	r.Addf("a, ok := xprs[0].(*%v)", bse.LitXpr.Ref(x))
	r.Add("tst.True(t, ok, \"cast\")")
	r.Add("tst.NotNil(t, a, \"cast result\")")
	r.Add("tst.UntEql(t, unt.Zero, a.Trm.Idx, cse.lit, \"Idx\")")
	r.Add("tst.UntEql(t, unt.Unt(len(cse.lit)), a.Trm.Lim, cse.lit, \"Lim\")")
	r.Add("tst.StringEql(t, cse.lit, xprr.Txt[a.Trm.Idx:a.Trm.Lim], cse.lit, \"Lit\")")
	r.Add("})")
	r.Add("}")
	return r
}

func (x *FleTest) XprAsn(typ Typ) (r *PkgFn) {
	bse := typ.Bse()
	x.Import(_sys.Lng.Pro.Xpr)
	x.Import(_sys.Bsc.Unt)
	r = x.TestFnf("Xpr%vAsn", typ.PkgTypTitle())
	r.Add("cses := []struct{")
	r.Add("lit string")
	r.Add("}{")
	for n, lit := range bse.Lits {
		if n == 0 && bse.IsArr() { // skip empty arr
			continue
		}
		r.Addf("{%q},", lit)
	}
	r.Add("}")
	r.Add("var xprr xpr.Xprr")
	r.Add("mem := \"asn\"")
	r.Add("for _, cse := range cses {")
	r.Add("t.Run(fmt.Sprintf(\"%q\", cse.lit), func(t *testing.T) {")
	r.Add("_, xprs := xprr.Prsf(\"%v.asn(a)\", cse.lit)")
	r.Add("tst.IntegerEql(t, 1, len(xprs), cse.lit, \"xprs count\")")
	r.Addf("a, ok := xprs[0].(*%v)", bse.AsnXpr.Ref(x))
	r.Add("tst.True(t, ok, \"cast\")")
	r.Add("tst.NotNil(t, a, \"cast result\")")
	r.Add("tst.UntEql(t, unt.Unt(len(cse.lit)+1), a.Trm.Idx, cse.lit, \"Idx\")")
	r.Add("tst.UntEql(t, unt.Unt(len(cse.lit)+1+len(mem)), a.Trm.Lim, cse.lit, \"Lim\")")
	r.Add("tst.StringEql(t, mem, xprr.Txt[a.Trm.Idx:a.Trm.Lim], cse.lit, \"Lit\")")
	r.Add("})")
	r.Add("}")
	return r
}
func (x *FleTest) XprAcs(typ Typ) (r *PkgFn) {
	bse := typ.Bse()
	x.Import(_sys.Lng.Pro.Xpr)
	x.Import(_sys.Bsc.Unt)
	r = x.TestFnf("Xpr%vAcs", typ.PkgTypTitle())
	r.Add("cses := []struct{")
	r.Add("lit string")
	r.Add("}{")
	for n, lit := range bse.Lits {
		if n == 0 && bse.IsArr() { // skip empty arr
			continue
		}
		r.Addf("{%q},", lit)
	}
	r.Add("}")
	r.Add("var xprr xpr.Xprr")
	r.Add("idnA := \"a\"")
	r.Add("for _, cse := range cses {")
	r.Add("t.Run(fmt.Sprintf(\"%q\", cse.lit), func(t *testing.T) {")
	r.Add("src := fmt.Sprintf(\"%v.asn(%v) %v\", cse.lit, idnA, idnA)")
	r.Add("_, xprs := xprr.Prs(src)")
	r.Add("tst.IntegerEql(t, 2, len(xprs), src, \"xprs count\")")
	r.Addf("a, ok := xprs[1].(*%v)", bse.AcsXpr.Ref(x))
	r.Add("tst.True(t, ok, \"cast\")")
	r.Add("tst.NotNil(t, a, \"cast result\")")
	r.Add("tst.UntEql(t, unt.Unt(len(src)-1), a.Trm.Idx, src, \"Idx\")")
	r.Add("tst.UntEql(t, unt.Unt(len(src)), a.Trm.Lim, src, \"Lim\")")
	r.Add("tst.StringEql(t, idnA, xprr.Txt[a.Trm.Idx:a.Trm.Lim], src, \"Lit\")")
	r.Add("})")
	r.Add("}")
	return r
}
func (x *FleTest) XprThen(typ Typ) (r *PkgFn) {
	bse := typ.Bse()
	x.Import(_sys.Lng.Pro.Xpr)
	x.Import(_sys.Bsc.Unt)
	r = x.TestFnf("Xpr%vThen", typ.PkgTypTitle())
	r.Add("cses := []struct{")
	r.Add("lit string")
	r.Add("}{")
	for n, lit := range bse.Lits {
		if n == 0 && bse.IsArr() { // skip empty arr
			continue
		}
		r.Addf("{%q},", lit)
	}
	r.Add("}")
	r.Add("var xprr xpr.Xprr")
	r.Add("mem := \"then\"")
	r.Add("for _, cse := range cses {")
	r.Add("t.Run(fmt.Sprintf(\"%q\", cse.lit), func(t *testing.T) {")
	r.Add("src := fmt.Sprintf(\"%v.then(0.asn(a) 1.asn(b) 2.asn(c))\", cse.lit)")
	r.Add("_, xprs := xprr.Prs(src)")
	r.Add("tst.IntegerEql(t, 1, len(xprs), src, \"xprs count\")")
	r.Addf("a, ok := xprs[0].(*%v)", bse.ThenXpr.Ref(x))
	r.Add("tst.True(t, ok, \"cast\")")
	r.Add("tst.NotNil(t, a, \"cast result\")")
	r.Add("tst.UntEql(t, unt.Unt(len(cse.lit)+1), a.Trm.Idx, cse.lit, \"Idx\")")
	r.Add("tst.UntEql(t, unt.Unt(len(cse.lit)+1+len(mem)), a.Trm.Lim, cse.lit, \"Lim\")")
	r.Add("tst.IntegerEql(t, 3, len(a.Xprs), \"Xprs\")")
	r.Add("})")
	r.Add("}")
	return r
}
func (x *FleTest) XprElse(typ Typ) (r *PkgFn) {
	bse := typ.Bse()
	x.Import(_sys.Lng.Pro.Xpr)
	x.Import(_sys.Bsc.Unt)
	r = x.TestFnf("Xpr%vElse", typ.PkgTypTitle())
	r.Add("cses := []struct{")
	r.Add("lit string")
	r.Add("}{")
	for n, lit := range bse.Lits {
		if n == 0 && bse.IsArr() { // skip empty arr
			continue
		}
		r.Addf("{%q},", lit)
	}
	r.Add("}")
	r.Add("var xprr xpr.Xprr")
	r.Add("mem := \"else\"")
	r.Add("for _, cse := range cses {")
	r.Add("t.Run(fmt.Sprintf(\"%q\", cse.lit), func(t *testing.T) {")
	r.Add("src := fmt.Sprintf(\"%v.else(0.asn(a) 1.asn(b) 2.asn(c))\", cse.lit)")
	r.Add("_, xprs := xprr.Prs(src)")
	r.Add("tst.IntegerEql(t, 1, len(xprs), src, \"xprs count\")")
	r.Addf("a, ok := xprs[0].(*%v)", bse.ElseXpr.Ref(x))
	r.Add("tst.True(t, ok, \"cast\")")
	r.Add("tst.NotNil(t, a, \"cast result\")")
	r.Add("tst.UntEql(t, unt.Unt(len(cse.lit)+1), a.Trm.Idx, cse.lit, \"Idx\")")
	r.Add("tst.UntEql(t, unt.Unt(len(cse.lit)+1+len(mem)), a.Trm.Lim, cse.lit, \"Lim\")")
	r.Add("tst.IntegerEql(t, 3, len(a.Xprs), \"Xprs\")")
	r.Add("})")
	r.Add("}")
	return r
}
func (x *FleTest) XprEach(typ Typ) (r *PkgFn) {
	bse := typ.Bse()
	x.Import(_sys.Lng.Pro.Xpr)
	x.Import(_sys.Bsc.Unt)
	r = x.TestFnf("Xpr%vEach", typ.PkgTypTitle())
	r.Add("cses := []struct{")
	r.Add("lit string")
	r.Add("}{")
	for n, lit := range bse.Lits {
		if n == 0 && bse.IsArr() { // skip empty arr
			continue
		}
		r.Addf("{%q},", lit)
	}
	r.Add("}")
	r.Add("var xprr xpr.Xprr")
	r.Add("mem, idnCur := \"each\", \"cur\"")
	r.Add("for _, cse := range cses {")
	r.Add("t.Run(fmt.Sprintf(\"%q\", cse.lit), func(t *testing.T) {")
	r.Add("src := fmt.Sprintf(\"%v.each(%v cur.asn(a) cur.asn(b) cur.asn(c))\", cse.lit, idnCur)")
	r.Add("_, xprs := xprr.Prs(src)")
	r.Add("tst.IntegerEql(t, 1, len(xprs), src, \"xprs count\")")
	r.Addf("a, ok := xprs[0].(*%v)", bse.EachXpr.Ref(x))
	r.Add("tst.True(t, ok, \"cast\")")
	r.Add("tst.NotNil(t, a, \"cast result\")")
	r.Add("tst.UntEql(t, unt.Unt(len(cse.lit)+1), a.Trm.Idx, cse.lit, \"Idx\")")
	r.Add("tst.UntEql(t, unt.Unt(len(cse.lit)+1+len(mem)), a.Trm.Lim, cse.lit, \"Lim\")")
	r.Add("tst.StringEql(t, idnCur, xprr.Txt[a.Idn.Idx:a.Idn.Lim], src, \"Lit Idn\")")
	r.Add("tst.IntegerEql(t, 3, len(a.Xprs), \"Xprs\")")
	r.Add("})")
	r.Add("}")
	return r
}
func (x *FleTest) XprPllEach(typ Typ) (r *PkgFn) {
	bse := typ.Bse()
	x.Import(_sys.Lng.Pro.Xpr)
	x.Import(_sys.Bsc.Unt)
	r = x.TestFnf("Xpr%vPllEach", typ.PkgTypTitle())
	r.Add("cses := []struct{")
	r.Add("lit string")
	r.Add("}{")
	for n, lit := range bse.Lits {
		if n == 0 && bse.IsArr() { // skip empty arr
			continue
		}
		r.Addf("{%q},", lit)
	}
	r.Add("}")
	r.Add("var xprr xpr.Xprr")
	r.Add("mem, idnCur := \"pllEach\", \"cur\"")
	r.Add("for _, cse := range cses {")
	r.Add("t.Run(fmt.Sprintf(\"%q\", cse.lit), func(t *testing.T) {")
	r.Add("src := fmt.Sprintf(\"%v.pllEach(%v cur.asn(a) cur.asn(b) cur.asn(c))\", cse.lit, idnCur)")
	r.Add("_, xprs := xprr.Prs(src)")
	r.Add("tst.IntegerEql(t, 1, len(xprs), src, \"xprs count\")")
	r.Addf("a, ok := xprs[0].(*%v)", bse.PllEachXpr.Ref(x))
	r.Add("tst.True(t, ok, \"cast\")")
	r.Add("tst.NotNil(t, a, \"cast result\")")
	r.Add("tst.UntEql(t, unt.Unt(len(cse.lit)+1), a.Trm.Idx, cse.lit, \"Idx\")")
	r.Add("tst.UntEql(t, unt.Unt(len(cse.lit)+1+len(mem)), a.Trm.Lim, cse.lit, \"Lim\")")
	r.Add("tst.StringEql(t, idnCur, xprr.Txt[a.Idn.Idx:a.Idn.Lim], src, \"Lit Idn\")")
	r.Add("tst.IntegerEql(t, 3, len(a.Xprs), \"Xprs\")")
	r.Add("})")
	r.Add("}")
	return r
}
func (x *FleTest) XprFldGet(s *Struct, v *Fld) (r *PkgFn) {
	x.Import(_sys.Lng.Pro.Xpr)
	r = x.TestFnf("Xpr%v%vGet", s.PkgTypTitle(), v.Title())
	r.Addf("x, mem := %q, %q", s.Lit(), v.Camel())
	r.Add("var xprr xpr.Xprr")
	r.Add("_, xprs := xprr.Prsf(\"%v.%v()\", x, mem)")
	r.Add("tst.IntegerEql(t, 1, len(xprs), \"xprs count\")")
	r.Addf("a, ok := xprs[0].(*%v)", v.GetXpr.Ref(x))
	r.Add("tst.True(t, ok, \"cast\")")
	r.Add("tst.NotNil(t, a, \"cast result\")")
	r.Add("tst.UntEql(t, unt.Unt(len(x)+1), a.Trm.Idx, \"Idx\")")
	r.Add("tst.UntEql(t, unt.Unt(len(x)+1+len(mem)), a.Trm.Lim, \"Lim\")")
	r.Add("tst.StringEql(t, mem, xprr.Txt[a.Trm.Idx : a.Trm.Lim], \"Lit\")")
	return r
}
func (x *FleTest) XprFldSetGet(s *Struct, fld *Fld) (r *PkgFn) {
	x.Import(_sys.Lng.Pro.Xpr)
	r = x.TestFnf("Xpr%v%vSetGet", s.PkgTypTitle(), fld.Title())
	r.Addf("x, mem := %q, %q", s.Lit(), fld.Camel())
	if fld.Typ == _sys.Bsc.Str.Typ() {
		r.Addf("val := %q", fld.Typ.Bse().Lit())
	} else {
		r.Addf("val := %v", fld.Typ.Bse().Lit())
	}
	r.Add("var xprr xpr.Xprr")
	r.Add("_, xprs := xprr.Prsf(\"%v.%v(%v)\", x, mem, val)")
	r.Add("tst.IntegerEql(t, 1, len(xprs), \"xprs count\")")
	r.Addf("a, ok := xprs[0].(*%v)", fld.SetGetXpr.Ref(x))
	r.Add("tst.True(t, ok, \"cast\")")
	r.Add("tst.NotNil(t, a, \"cast result\")")
	r.Add("tst.UntEql(t, unt.Unt(len(x)+1), a.Trm.Idx, \"Idx\")")
	r.Add("tst.UntEql(t, unt.Unt(len(x)+1+len(mem)), a.Trm.Lim, \"Lim\")")
	r.Add("tst.StringEql(t, mem, xprr.Txt[a.Trm.Idx : a.Trm.Lim], \"Lit\")")
	return r
}
func (x *FleTest) XprCnst(c *Cnst) (r *PkgFn) {
	x.Import(_sys.Lng.Pro.Xpr)
	r = x.TestFnf("Xpr%v%vCnst", c.Pkg.Title(), c.Title())
	r.Addf("pkg, mem := %q, %q", c.Pkg.Lower(), c.Camel())
	r.Add("var xprr xpr.Xprr")
	r.Add("_, xprs := xprr.Prsf(\"%v.%v\", pkg, mem)")
	r.Add("tst.IntegerEql(t, 1, len(xprs), \"xprs count\")")
	r.Addf("a, ok := xprs[0].(*%v)", c.Xpr.Ref(x))
	r.Add("tst.True(t, ok, \"cast\")")
	r.Add("tst.NotNil(t, a, \"cast result\")")
	r.Add("tst.UntEql(t, unt.Unt(len(pkg)+1), a.Trm.Idx, \"Idx\")")
	r.Add("tst.UntEql(t, unt.Unt(len(pkg)+1+len(mem)), a.Trm.Lim, \"Lim\")")
	r.Add("tst.StringEql(t, mem, xprr.Txt[a.Trm.Idx : a.Trm.Lim], \"Lit\")")
	return r
}
func (x *FleTest) XprVar(v *Var) (r *PkgFn) {
	x.Import(_sys.Lng.Pro.Xpr)
	r = x.TestFnf("Xpr%v%vVar", v.Pkg.Title(), v.Title())
	r.Addf("pkg, mem := %q, %q", v.Pkg.Lower(), v.Camel())
	r.Add("var xprr xpr.Xprr")
	r.Add("_, xprs := xprr.Prsf(\"%v.%v\", pkg, mem)")
	r.Add("tst.IntegerEql(t, 1, len(xprs), \"xprs count\")")
	r.Addf("a, ok := xprs[0].(*%v)", v.Xpr.Ref(x))
	r.Add("tst.True(t, ok, \"cast\")")
	r.Add("tst.NotNil(t, a, \"cast result\")")
	r.Add("tst.UntEql(t, unt.Unt(len(pkg)+1), a.Trm.Idx, \"Idx\")")
	r.Add("tst.UntEql(t, unt.Unt(len(pkg)+1+len(mem)), a.Trm.Lim, \"Lim\")")
	r.Add("tst.StringEql(t, mem, xprr.Txt[a.Trm.Idx : a.Trm.Lim], \"Lit\")")
	return r
}
func (x *FleTest) XprPkgFn(fn *PkgFn) (r *PkgFn) {
	r = x.TestFnf("Xpr%v%vPkgFn", fn.Pkg.Title(), fn.Title())
	r.Addf("pkg, mem := %q, %q", fn.Pkg.Lower(), fn.Camel())
	r.Add("var xprr xpr.Xprr")
	if !fn.InPrms.Ok() { // no in prms
		r.Add("_, xprs := xprr.Prs(fmt.Sprintf(\"%v.%v()\", pkg, mem))")
	} else { // generate in prm lits
		r.Addf("b := %v{}", BuilderPtr.Adr(x))
		for n, inPrm := range fn.InPrms {
			if n != 0 {
				r.Add("b.WriteRune(' ')")
			}
			r.Addf("b.WriteString(%q)", inPrm.Typ.Bse().Lit()) // idx 1 to avoid any empty array
		}
		r.Add("_, xprs := xprr.Prs(fmt.Sprintf(\"%v.%v(%v)\", pkg, mem, b.String()))")
	}
	r.Add("tst.IntegerEql(t, 1, len(xprs), \"xprs count\")")
	r.Addf("a, ok := xprs[0].(*%v)", fn.Xpr.Ref(x))
	r.Add("tst.True(t, ok, \"cast\")")
	r.Add("tst.NotNil(t, a, \"cast result\")")
	r.Add("tst.UntEql(t, unt.Unt(len(pkg)+1), a.Trm.Idx, \"Idx\")")
	r.Add("tst.UntEql(t, unt.Unt(len(pkg)+1+len(mem)), a.Trm.Lim, \"Lim\")")
	r.Add("tst.StringEql(t, mem, xprr.Txt[a.Trm.Idx : a.Trm.Lim], \"Lit\")")
	return r
}
func (x *FleTest) XprTypFn(fn *TypFn) (r *PkgFn) {
	r = x.TestFnf("Xpr%v%vTypFn", fn.Rxr.Typ.PkgTypTitle(), fn.Name)
	b := &strings.Builder{}
	for n, prm := range fn.InPrms {
		if n != 0 {
			b.WriteRune(' ')
		}
		b.WriteString(prm.Typ.Bse().Lit())
	}
	r.Addf("x, mem, prms := %q, %q, %q", fn.Rxr.Typ.Bse().Lit(), fn.Camel(), b.String())
	r.Add("var xprr xpr.Xprr")
	r.Add("_, xprs := xprr.Prsf(\"%v.%v(%v)\", x, mem, prms)")
	r.Add("tst.IntegerEql(t, 1, len(xprs), \"xprs count\")")
	r.Addf("a, ok := xprs[0].(*%v)", fn.Xpr.Ref(x))
	r.Add("tst.True(t, ok, \"cast\")")
	r.Add("tst.NotNil(t, a, \"cast result\")")
	r.Add("tst.UntEql(t, unt.Unt(len(x)+1), a.Trm.Idx, \"Idx\")")
	r.Add("tst.UntEql(t, unt.Unt(len(x)+1+len(mem)), a.Trm.Lim, \"Lim\")")
	r.Add("tst.StringEql(t, mem, xprr.Txt[a.Trm.Idx : a.Trm.Lim], \"Lit\")")
	return r
}
func (x *FleTest) XprMemSig(fn *MemSig) (r *PkgFn) {
	r = x.TestFnf("Xpr%v%vMemSig", fn.Rxr.PkgTypTitle(), fn.Name)
	b := &strings.Builder{}
	for n, prm := range fn.InPrms {
		if n != 0 {
			b.WriteRune(' ')
		}
		b.WriteString(prm.Typ.Bse().Lit())
	}
	r.Addf("x, mem, prms := %q, %q, %q", fn.Rxr.Bse().Lit(), fn.Camel(), b.String())
	r.Add("var xprr xpr.Xprr")
	r.Add("_, xprs := xprr.Prsf(\"%v.%v(%v)\", x, mem, prms)")
	r.Add("tst.IntegerEql(t, 1, len(xprs), \"xprs count\")")
	r.Addf("a, ok := xprs[0].(*%v)", fn.Xpr.Ref(x))
	r.Add("tst.True(t, ok, \"cast\")")
	r.Add("tst.NotNil(t, a, \"cast result\")")
	r.Add("tst.UntEql(t, unt.Unt(len(x)+1), a.Trm.Idx, \"Idx\")")
	r.Add("tst.UntEql(t, unt.Unt(len(x)+1+len(mem)), a.Trm.Lim, \"Lim\")")
	r.Add("tst.StringEql(t, mem, xprr.Txt[a.Trm.Idx : a.Trm.Lim], \"Lit\")")
	return r
}

func (x *FleTest) ActLit(typ Typ) (r *PkgFn) {
	bse := typ.Bse()
	r = x.TestFnf("Act%vLit", bse.PkgTypTitle())
	x.Import(_sys.Lng.Pro.Act)
	if bse.IsArr() || bse.IsBnd() {
		x.Import(bse)
	}
	r.Add("var actr act.Actr")
	r.Addf("expected := %v", bse.Val())
	r.Addf("acts := actr.Cmpl(%q)", bse.Lit())
	r.Add("tst.IntegerEql(t, 1, len(acts), \"acts count\")")
	r.Addf("act, ok := acts[0].(%v)", bse.LitAct.Ref(x))
	r.Add("tst.True(t, ok, \"cast\")")
	r.Addf("%v(t, expected, act.%v(), \"act\")", bse.TstEql.Ref(x), bse.PkgTypTitle())
	return r
}
func (x *FleTest) ActAsn(typ Typ) (r *PkgFn) {
	bse := typ.Bse()
	r = x.TestFnf("Act%vAsn", bse.PkgTypTitle())
	src := fmt.Sprintf("%v.asn(a)", bse.Lit())
	if !bse.IsAna() {
		r.Add("var actr act.Actr")
		r.Addf("expected := %v", bse.Val(true))
		r.Addf("acts := actr.Cmpl(%q)", src)
	} else {
		x.Import("sys/app")
		r.Add("ap := app.New(tst.Cfg)")
		r.Add("ap.Oan.CloneInstrs(tst.Instrs)")
		r.Add("defer ap.Cls()")
		r.Addf("expected := %v", bse.Val(true))
		r.Addf("acts := ap.Actr.Cmpl(%q)", src)
	}
	r.Add("tst.IntegerEql(t, 1, len(acts), \"acts count\")")
	r.Addf("act, ok := acts[0].(%v)", bse.AsnAct.Ref(x))
	r.Add("tst.True(t, ok, \"cast\")")
	r.Addf("%v(t, expected, act.%v(), \"act\")", bse.TstEql.Ref(x), bse.PkgTypTitle())
	return r
}
func (x *FleTest) ActAcs(typ Typ) (r *PkgFn) {
	bse := typ.Bse()
	r = x.TestFnf("Act%vAcs", bse.PkgTypTitle())
	src := fmt.Sprintf("%v.asn(a) a", bse.Lit())
	if !bse.IsAna() {
		r.Add("var actr act.Actr")
		r.Addf("expected := %v", bse.Val(true))
		r.Addf("acts := actr.Cmpl(%q)", src)
	} else {
		x.Import("sys/app")
		r.Add("ap := app.New(tst.Cfg)")
		r.Add("ap.Oan.CloneInstrs(tst.Instrs)")
		r.Add("defer ap.Cls()")
		r.Addf("expected := %v", bse.Val(true))
		r.Addf("acts := ap.Actr.Cmpl(%q)", src)
	}
	r.Add("tst.IntegerEql(t, 2, len(acts), \"acts count\")")
	r.Add("acts[0].Act()")
	r.Addf("act, ok := acts[1].(%v)", bse.AcsAct.Ref(x))
	r.Add("tst.True(t, ok, \"cast\")")
	r.Addf("%v(t, expected, act.%v(), \"act\")", bse.TstEql.Ref(x), bse.PkgTypTitle())
	return r
}
func (x *FleTest) ActThen(typ Typ) (r *PkgFn) {
	bse := typ.Bse()
	r = x.TestFnf("Act%vThen", bse.PkgTypTitle())
	src := fmt.Sprintf("tru.then(0.asn(a) 1.asn(b) 2.asn(x))")
	r.Add("var actr act.Actr")
	r.Addf("actr.Run(%q)", src)
	// TODO: MORE TESTING?
	return r
}
func (x *FleTest) ActElse(typ Typ) (r *PkgFn) {
	bse := typ.Bse()
	r = x.TestFnf("Act%vElse", bse.PkgTypTitle())
	src := fmt.Sprintf("fls.else(0.asn(a) 1.asn(b) 2.asn(x))")
	r.Add("var actr act.Actr")
	r.Addf("actr.Run(%q)", src)
	// TODO: MORE TESTING?
	return r
}
func (x *FleTest) ActFldGet(s *Struct, fld *Fld) (r *PkgFn) {
	r = x.TestFnf("Act%v%vGet", s.PkgTypTitle(), fld.Title())
	src := fmt.Sprintf("%v.%v()", s.Lit(), fld.Camel())
	r.Add("var actr act.Actr")
	r.Addf("actr.Run(%q)", src)
	// TODO: MORE TESTING?
	return r
}
func (x *FleTest) ActFldSetGet(s *Struct, fld *Fld) (r *PkgFn) {
	r = x.TestFnf("Act%v%vSetGet", s.PkgTypTitle(), fld.Title())
	src := fmt.Sprintf("%v.%v(%v)", s.Lit(), fld.Camel(), fld.Typ.Bse().Lit())
	r.Add("var actr act.Actr")
	r.Addf("actr.Run(%q)", src)
	// TODO: MORE TESTING?
	return r
}
func (x *FleTest) ActEach(typ Typ) (r *PkgFn) {
	bse := typ.Bse()
	r = x.TestFnf("Act%vEach", bse.PkgTypTitle())
	// TODO: ADD "b" ref
	src := fmt.Sprintf("33.asn(cnt) 0 0.asn(dd) 0.asn(d) 44.asn(lst) %v.each(cur  cur.asn(a)  cnt.add(1).asn(cnt)  lst.asn(b)) cnt", bse.Lit())
	if !bse.IsAna() {
		r.Add("var actr act.Actr")
		r.Addf("expected := %v", bse.Val(true))
		r.Addf("acts := actr.Cmpl(%q)", src)
	} else {
		x.Import("sys/app")
		r.Add("ap := app.New(tst.Cfg)")
		r.Add("ap.Oan.CloneInstrs(tst.Instrs)")
		r.Add("defer ap.Cls()")
		r.Addf("expected := %v", bse.Val(true))
		r.Addf("acts := ap.Actr.Cmpl(%q)", src)
	}
	// TODO: ADD NESTED EACH TEST
	r.Add("tst.IntegerEql(t, 7, len(acts), \"acts count\")")
	r.Addf("acts[0].Act()")
	r.Addf("acts[1].Act()")
	r.Addf("acts[2].Act()")
	r.Addf("acts[3].Act()")
	r.Addf("acts[4].Act()")
	r.Addf("a5, ok := acts[5].(%v)", bse.EachAct.Ref(x))
	r.Add("tst.True(t, ok, \"cast\")")
	r.Addf("%v(t, expected, a5.%v(), \"each act\")", bse.TstEql.Ref(x), bse.PkgTypTitle())
	r.Add("tst.IntegerEql(t, 3, len(a5.Acts), \"inner acts count\")")

	untBse := _sys.Bsc.Unt.Typ().Bse()
	r.Addf("a5CntAsn, ok := a5.Acts[1].(%v)", untBse.AsnAct.Ref(x))
	r.Add("tst.True(t, ok, \"cast\")")
	r.Addf("%v(t, expected.Cnt(), a5CntAsn.Arr[a5CntAsn.Idx], \"a5CntAsn: inner scp inc\")", untBse.TstEql.Ref(x))

	r.Addf("a2BAsn, ok := a5.Acts[2].(%v)", untBse.AsnAct.Ref(x))
	r.Add("tst.True(t, ok, \"cast\")")
	r.Addf("%v(t, unt.Unt(44), a2BAsn.Arr[a2BAsn.Idx], \"a2BAsn: inner scp asn\")", untBse.TstEql.Ref(x))

	r.Addf("a6, ok := acts[6].(%v)", untBse.AcsAct.Ref(x))
	r.Add("tst.True(t, ok, \"cast\")")
	r.Addf("%v(t, 33, a6.%v(), \"outer scp cnt\")", untBse.TstEql.Ref(x), untBse.PkgTypTitle())
	return r
}

func (x *FleTest) ActPllEach(typ Typ) (r *PkgFn) {
	bse := typ.Bse()
	r = x.TestFnf("Act%vPllEach", bse.PkgTypTitle())
	// TODO: ADD "b" ref
	src := fmt.Sprintf("33.asn(cnt) 0 0.asn(dd) 0.asn(d) 44.asn(lst) %v.pllEach(cur  cur.asn(a)  cnt.add(1).asn(cnt)  lst.asn(b)) cnt", bse.Lit())
	if !bse.IsAna() {
		r.Add("var actr act.Actr")
		r.Addf("expected := %v", bse.Val(true))
		r.Addf("acts := actr.Cmpl(%q)", src)
	} else {
		x.Import("sys/app")
		r.Add("ap := app.New(tst.Cfg)")
		r.Add("ap.Oan.CloneInstrs(tst.Instrs)")
		r.Add("defer ap.Cls()")
		r.Addf("expected := %v", bse.Val(true))
		r.Addf("acts := ap.Actr.Cmpl(%q)", src)
	}
	// TODO: ADD NESTED EACH TEST
	r.Add("tst.IntegerEql(t, 7, len(acts), \"acts count\")")
	r.Addf("acts[0].Act()")
	r.Addf("acts[1].Act()")
	r.Addf("acts[2].Act()")
	r.Addf("acts[3].Act()")
	r.Addf("acts[4].Act()")
	r.Addf("a5, ok := acts[5].(%v)", bse.PllEachAct.Ref(x))
	r.Add("tst.True(t, ok, \"cast\")")
	r.Addf("%v(t, expected, a5.%v(), \"pllEach act\")", bse.TstEql.Ref(x), bse.PkgTypTitle())

	untBse := _sys.Bsc.Unt.Typ().Bse()

	// r.Add("tst.IntegerEql(t, 3, len(a5.Acts), \"inner acts count\")")
	//
	// r.Addf("a5CntAsn, ok := a5.Acts[1].(%v)", untBse.AsnAct.Ref(x))
	// r.Add("tst.True(t, ok, \"cast\")")
	// r.Addf("%v(t, expected.Cnt(), a5CntAsn.Arr[a5CntAsn.Idx], \"a5CntAsn: inner scp inc\")", untBse.TstEql.Ref(x))
	// r.Addf("a2BAsn, ok := a5.Acts[2].(%v)", untBse.AsnAct.Ref(x))
	// r.Add("tst.True(t, ok, \"cast\")")
	// r.Addf("%v(t, unt.Unt(44), a2BAsn.Arr[a2BAsn.Idx], \"a2BAsn: inner scp asn\")", untBse.TstEql.Ref(x))

	r.Addf("a6, ok := acts[6].(%v)", untBse.AcsAct.Ref(x))
	r.Add("tst.True(t, ok, \"cast\")")
	r.Addf("%v(t, 33, a6.%v(), \"outer scp cnt\")", untBse.TstEql.Ref(x), untBse.PkgTypTitle())
	return r
}

func (x *FleTest) ActCnst(c *Cnst) (r *PkgFn) {
	x.Import(c.Pkg.Pth)
	r = x.TestFnf("Act%v%vCnst", c.Pkg.Title(), c.Title())
	src := c.Ref(x, true)
	r.Add("var actr act.Actr")
	r.Addf("expected := %v", c.Ref(x))
	r.Addf("acts := actr.Cmpl(%q)", src)
	r.Add("tst.IntegerEql(t, 1, len(acts), \"acts count\")")
	r.Addf("act, ok := acts[0].(%v)", c.Act.Ref(x))
	r.Add("tst.True(t, ok, \"cast\")")
	r.Addf("%v(t, expected, act.%v(), \"act\")", c.Typ.Bse().TstEql.Ref(x), c.Typ.PkgTypTitle())
	return r
}
func (x *FleTest) ActVar(v *Var) (r *PkgFn) {
	x.Import(v.Pkg.Pth)
	r = x.TestFnf("Act%v%vVar", v.Pkg.Title(), v.Title())
	src := v.Ref(x, true)
	r.Add("var actr act.Actr")
	r.Addf("expected := %v", v.Ref(x))
	r.Addf("acts := actr.Cmpl(%q)", src)
	r.Add("tst.IntegerEql(t, 1, len(acts), \"acts count\")")
	r.Addf("act, ok := acts[0].(%v)", v.Act.Ref(x))
	r.Add("tst.True(t, ok, \"cast\")")
	r.Addf("%v(t, expected, act.%v(), \"act\")", v.Typ.Bse().TstEql.Ref(x), v.Typ.PkgTypTitle())
	return r
}
func (x *FleTest) ActPkgFn(fn *PkgFn) (r *PkgFn) {
	// sys.Log("- FleTest.ActPkgFn", fn.Pkg.Name, fn.Name)
	outBse := fn.OutTyp().Bse()
	x.Import(outBse)
	x.Import(fn)
	r = x.TestFnf("Act%v%vPkgFn", fn.Pkg.Title(), fn.Title())
	// sys.Log("~~~ FleTest.ActPkgFn", x.Name, fn.PkgTitle())
	src := fmt.Sprintf("%v(%v)", fn.Ref(x, true), fn.InPrms.Lits())
	if !outBse.IsAna() {
		for _, p := range fn.InPrms {
			if p.Typ.Bse().IsBsc() {
				x.Import(p.Typ) // for actr.gen_test
			}
		}
		r.Add("var actr act.Actr")
		r.Addf("expected := %v(%v)", fn.Ref(x), fn.InPrms.Vals())
		r.Addf("acts := actr.Cmplf(\"%v(%v)\", %v)", fn.Ref(x, true), fn.InPrms.ValVs(), fn.InPrms.Vals(true))
	} else {
		x.Import("sys/app")
		r.Add("ap := app.New(tst.Cfg)")
		r.Add("ap.Oan.CloneInstrs(tst.Instrs)")
		r.Add("defer ap.Cls()")
		r.Addf("expected := %v(%v)", fn.Ref(x), fn.InPrms.Vals(true))
		r.Addf("acts := ap.Actr.Cmpl(%q)", src)
	}
	r.Add("tst.IntegerEql(t, 1, len(acts), \"acts count\")")
	r.Addf("act, ok := acts[0].(%v)", fn.Act.Ref(x))
	r.Add("tst.True(t, ok, \"cast\")")
	r.Addf("%v(t, expected, act.%v(), \"act\")", outBse.TstEql.Ref(x), outBse.PkgTypTitle())
	return r
}
func (x *FleTest) ActTypFn(fn *TypFn) (r *PkgFn) {
	b := &strings.Builder{}
	rxrBse, outBse := fn.Rxr.Typ.Bse(), fn.OutTyp().Bse()
	x.Import(outBse.Pkg.Pth)
	r = x.TestFnf("Act%v%vTypFn", rxrBse.PkgTypTitle(), strings.Title(fn.Name))
	src := fmt.Sprintf("%v.%v(%v)", rxrBse.Lit(), fn.Camel(), fn.InPrms.Lits())
	// UNCOMMENTED FOR FLTS sys.Pll support
	// if !rxrBse.IsAna() {
	// 	r.Add("var actr act.Actr")
	// } else {
	x.Import("sys/app")
	r.Add("ap := app.New(tst.Cfg)")
	r.Add("ap.Oan.CloneInstrs(tst.Instrs)")
	r.Add("defer ap.Cls()")
	r.Add("actr := ap.Actr")
	// }
	r.Add(fn.Rxr.TestFldAsn(b, true))
	fn.InPrms.TestFldAsn(r, b, true)
	r.Addf("expected := %v.%v(%v)", fn.Rxr.Camel(), fn.Title(), fn.InPrms.Names())
	r.Addf("acts := actr.Cmpl(%q)", src)
	r.Add("tst.IntegerEql(t, 1, len(acts), \"acts count\")")
	r.Addf("act, ok := acts[0].(%v)", fn.Act.Ref(x))
	r.Add("tst.True(t, ok, \"cast\")")
	r.Addf("%v(t, expected, act.%v(), \"act\")", outBse.TstEql.Ref(x), outBse.PkgTypTitle())
	return r
}

func (x *FleTest) ActMemSig(fn *MemSig) (r *PkgFn) {
	rxrBse, outBse := fn.Rxr.Bse(), fn.OutTyp().Bse()
	x.ImportTyp(fn.Rxr)
	r = x.TestFnf("Act%v%vMemSig", rxrBse.PkgTypTitle(), strings.Title(fn.Name))
	src := fmt.Sprintf("%v.%v(%v)", rxrBse.Lit(), fn.Camel(), fn.InPrms.Lits())
	if !rxrBse.IsAna() {
		r.Add("var actr act.Actr")
		r.Addf("expected := %v.%v(%v)", rxrBse.Val(true), fn.Title(), fn.InPrms.Vals(true))
		r.Addf("acts := actr.Cmpl(%q)", src)
	} else {
		x.Import("sys/app")
		r.Add("ap := app.New(tst.Cfg)")
		r.Add("ap.Oan.CloneInstrs(tst.Instrs)")
		r.Add("defer ap.Cls()")
		r.Addf("expected := %v.%v(%v)", rxrBse.Val(true), fn.Title(), fn.InPrms.Vals2(true))
		r.Addf("acts := ap.Actr.Cmpl(%q)", src)
	}
	r.Add("tst.IntegerEql(t, 1, len(acts), \"acts count\")")
	r.Addf("act, ok := acts[0].(%v)", fn.Act.Ref(x))
	r.Add("tst.True(t, ok, \"cast\")")
	r.Addf("%v(t, expected, act.%v(), \"act\")", outBse.TstEql.Ref(x), outBse.PkgTypTitle())
	return r
}

func (x *FleTest) TestSysAnaFn() (r *PkgFn, bse *TypBse) {
	x.Import("sys/app")
	return x.TestFn(x.Src.Typ().PkgTypTitle()), x.Src.Typ().Bse()
}

func (x *FleTest) GenPkgFn(fn *PkgFn, testPth []*TestStp) (r *PkgFn) {
	x.Import("sys/app")
	r = x.TestFn(fn.PkgTitle())
	for _, stp := range testPth {
		if stp.Fst != nil {
			stp.Fst(r)
		}
	}
	r.Add("t.Run(\"\", func(t *testing.T) {")
	r.Add("ap := app.New(tst.Cfg)")
	r.Add("ap.Oan.CloneInstrs(tst.Instrs)")
	for _, stp := range testPth {
		if stp.MdlFst != nil {
			stp.MdlFst(r)
		}
	}
	r.Add(fn.T2.MdlFst...)
	if fn.Node != nil {
		r.Addf("a := %v%v", fn.CallLit(), fn.Node.Cast(x, true))
		r.Addf("%v(t, a)", fn.Node.TstNotZero.Ref(x))
	} else {
		r.Addf("a := %v", fn.CallLit())
		r.Addf("%v(t, a)", fn.OutTyp().Bse().TstNotZero.Ref(x))
	}
	r.Add(fn.T2.MdlLst...)
	for _, stp := range testPth {
		if stp.MdlLst != nil {
			stp.MdlLst(r)
		}
	}
	r.Add("ap.Cls()")
	r.Add("})")
	for _, stp := range testPth {
		if stp.Lst != nil {
			stp.Lst(r)
		}
	}
	return r
}

func (x *FleTest) Gen(fn *TypFn, testPth []*TestStp, rxr ...Typ) (r *PkgFn) {
	var bse *TypBse
	if len(rxr) != 0 {
		bse = rxr[0].Bse()
	} else {
		bse = fn.Rxr.Typ.Bse()
	}
	x.Import("sys/app")
	var typ Typ
	if bse.ifc != nil {
		typ = bse.ifc
	} else {
		typ = bse
	}
	r = x.TestFnf("%v%v", typ.PkgTypTitle(), fn.Name)
	for _, stp := range testPth {
		if stp.Fst != nil {
			stp.Fst(r)
		}
	}
	r.Add("t.Run(\"\", func(t *testing.T) {")
	r.Add("ap := app.New(tst.Cfg)")
	r.Add("ap.Oan.CloneInstrs(tst.Instrs)")
	for _, stp := range testPth {
		if stp.MdlFst != nil {
			stp.MdlFst(r)
		}
	}
	r.Add(fn.T2.MdlFst...)
	if fn.Node != nil {
		r.Addf("a := %v%v%v", typ.Camel(), fn.CallVal(x), fn.Node.Cast(x, true))
		r.Addf("%v(t, a)", fn.Node.TstNotZero.Ref(x))
	} else {
		r.Addf("a := %v%v", typ.Camel(), fn.CallVal(x))
		r.Addf("%v(t, a)", fn.OutTyp().Bse().TstNotZero.Ref(x))
	}
	r.Add(fn.T2.MdlLst...)
	for _, stp := range testPth {
		if stp.MdlLst != nil {
			stp.MdlLst(r)
		}
	}
	r.Add("ap.Cls()")
	r.Add("})")
	for _, stp := range testPth {
		if stp.Lst != nil {
			stp.Lst(r)
		}
	}
	return r
}

// func (x *FleTest) Stgy(fn *TypFn) (r *PkgFn) {
// 	node, name, pkgName := fn.Node, fn.Name, strings.Title(fn.Node.Pkg.Name)
// 	x.Import("sys/app")
// 	r = x.TestFn(node.PkgTypTitle() + "__PRV")

// 	r.Addf("inrvl, side, rte, stm, cnd := tst.%[1]vInstrInrvlI, tst.%[1]vInrvlSideBid, tst.%[1]vSideStmRteLst, tst.%[1]vStmStmUnaPos, tst.%[1]vStmCndInrGtr", pkgName)

// 	r.Addf("for _, instr := range tst.%vPrvInstrs {", pkgName)

// 	r.Addf("t.Run(%q, func(t *testing.T) {", name)

// 	r.Add("ap := app.New(tst.Cfg)")
// 	r.Add("ap.Oan.CloneInstrs(tst.Instrs)")

// 	r.Addf("instr := instr(%v.Oan())", node.Pkg.Name)
// 	r.Add("stm := stm(rte(side(inrvl(instr, 10))))")
// 	r.Add("prnt := cnd(stm, 1)")
// 	r.Addf("a := prnt%v%v", fn.CallLit(), node.Cast(x, true))
// 	r.Add("tst.NotNil(t, a)") // Tmes nil is valid

// 	if node.Pkg.Name == k.Hst {
// 		r.Lines = append(r.Lines, fn.T.Lines...)
// 	} else if node.Pkg.Name == k.Rlt {
// 		r.Add("// SEE PRFM TEST FOR TRD COMPARISON")
// 	}
// 	r.Add("ap.Cls()")

// 	r.Add("})")
// 	r.Add("}")

// 	return r
// }

// func (x *FleTest) Splt(fn *TypFn) (r *PkgFn) {
// 	node, name, pkgName := fn.Node, fn.Name, strings.Title(fn.Node.Pkg.Name)
// 	x.Import("sys/app")
// 	r = x.TestFn(node.PkgTypTitle() + "__PRV")

// 	r.Addf("inrvl, side, rte, stm, cnd, stgy := tst.%[1]vInstrInrvlI, tst.%[1]vInrvlSideBid, tst.%[1]vSideStmRteLst, tst.%[1]vStmStmUnaPos, tst.%[1]vStmCndInrGtr, tst.%[1]vCndStgyLong", pkgName)
// 	r.Addf("for _, instr := range tst.%vPrvInstrs {", pkgName)

// 	r.Addf("t.Run(%q, func(t *testing.T) {", name)

// 	r.Add("ap := app.New(tst.Cfg)")
// 	r.Add("ap.Oan.CloneInstrs(tst.Instrs)")

// 	r.Addf("instr := instr(%v.Oan())", node.Pkg.Name)
// 	r.Add("stgy := stgy(cnd(stm(rte(side(inrvl(instr, 10)))), 1), 1.1, 1.1, 60*60, instr)")
// 	r.Addf("%[1]v.NewPort(stgy).Prfm()", node.Pkg.Name)

// 	r.Addf("a := stgy%v%v", fn.CallLit(), node.Cast(x, true))
// 	r.Add("tst.NotNil(t, a)")
// 	r.Add("tst.NotNil(t, a.Btm)")
// 	r.Add("tst.NotNil(t, a.Top)")

// 	if node.Pkg.Name == k.Hst {
// 	} else if node.Pkg.Name == k.Rlt {
// 	}
// 	r.Add("ap.Cls()")

// 	r.Add("})")
// 	r.Add("}")

// 	return r
// }

// func (x *FleTest) Prfm(fn *TypFn) (r *PkgFn) {
// 	rxr, name, pkgName := fn.Rxr.Typ.Bse(), fn.Name, strings.Title(fn.Rxr.Typ.Bse().Pkg.Name)
// 	x.Import("sys/app")
// 	r = x.TestFn(fn.Rxr.Typ.PkgTypTitle() + "__PRV")

// 	r.Addf("inrvl, side, rte, stm, cnd, stgy := tst.%[1]vInstrInrvlI, tst.%[1]vInrvlSideBid, tst.%[1]vSideStmRteLst, tst.%[1]vStmStmUnaPos, tst.%[1]vStmCndInrGtr, tst.%[1]vCndStgyLong", pkgName)
// 	r.Addf("for _, instr := range tst.%vPrvInstrs {", pkgName)

// 	r.Addf("t.Run(%q, func(t *testing.T) {", name)

// 	r.Add("ap := app.New(tst.Cfg)")
// 	r.Add("ap.Oan.CloneInstrs(tst.Instrs)")

// 	r.Addf("instr := instr(%v.Oan())", rxr.Pkg.Name)
// 	r.Add("stgy := stgy(cnd(stm(rte(side(inrvl(instr, 10)))), 1), 1.1, 1.1, 60*60, instr)")
// 	r.Addf("%[1]vPort := %[1]v.NewPort(stgy)", rxr.Pkg.Name)
// 	r.Addf("tst.NotNil(t, %vPort)", rxr.Pkg.Name)

// 	if rxr.Pkg.Name == k.Hst {
// 	} else if rxr.Pkg.Name == k.Rlt {
// 		x.Import(_sys.Lng.Pro.Act)
// 		x.Import(_sys.Ana.Hst)
// 		r.Add("mnr := tst.NewStgyMnr(ap)")
// 		r.Add("rltPort.Sub(mnr.Rx, mnr.Id)")
// 		r.Add("tst.IntegerEql(t, 1, len(rltPort.Rxs), \"Sub RxsCnt\")")
// 		r.Add("var actr act.Actr")
// 		r.Add("hstStgy := actr.RunHst(stgy.String())[0].(hst.Stgy)")
// 		r.Add("hstPort := hst.NewPort(hstStgy)")
// 		r.Add("hstPrfm := hstPort.Prfm() // call before WaitFor (for CalcTrds)")
// 		r.Add("mnr.Start(instr.Instr())")
// 		r.Add("mnr.WaitFor(hstPort.Trds.Cnt())")
// 		r.Add("if hstPort.Trds != nil { // cnt dlta due to rlt heartbeat")
// 		r.Add("tst.AnaTrdsEql(t, hstStgy.Bse().Trds, stgy.Bse().Trds, \"Stgy.Trds\")")
// 		r.Add("tst.AnaTrdsEql(t, hstPort.Trds, mnr.Trds, \"Port.Trds\")")
// 		r.Add("tst.AnaPortEql(t, &hstPort.Port, &rltPort.Port, \"Port\")")
// 		r.Add("tst.AnaPrfmEql(t, hstPrfm, rltPort.Prfm(), \"Prfm\")")
// 		r.Add("}")
// 		r.Add("rltPort.Unsub(mnr.Id)")
// 		r.Add("tst.IntegerEql(t, 0, len(rltPort.Rxs), \"Unsub RxsCnt\")")
// 	}
// 	r.Add("ap.Cls()")

// 	r.Add("})")
// 	r.Add("}")

// 	return r
// }

var (
	NonIdnEdges   = []string{" ", "\n", ".", ":", ",", "\"", "(", ")", "[", "]", "// comment"}
	FalseEdgesTxt = []string{"_", "9713", "Z"}
	FalseEdgesNum = []string{"_", "Z"}
)

type (
	Gen byte
)

const (
	TxtGen Gen = 1 << iota
	NumGen
	IdnGen
	PrefixGen
	SuffixGen
	None Gen = 0
)
