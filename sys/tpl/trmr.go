package tpl

import (
	"fmt"
	"strings"
	"sys"
	"sys/k"
	"sys/tpl/atr"
	"unicode/utf8"
)

type (
	FleTrmr struct {
		FleBse
		Pkgs
		Lexs     map[string]*TypFn
		Spce     *Struct
		Cmnt     *Struct
		Idn      *Struct
		DatePart *Struct
		TmePrt   *Struct
		Date     *Struct
		Time     *Struct
		Objs     *Struct
	}
)

func (x *DirTrm) NewTrmr() (r *FleTrmr) {
	r = &FleTrmr{}
	r.Name = k.Trmr
	r.Pkg = x.Pkg
	r.Trmr()
	r.AddFle(r)
	r.Lexs = make(map[string]*TypFn)
	return r
}
func (x *FleTrmr) Trmr() (r *Struct) {
	x.Import(_sys.Lng.Scn)
	r = x.StructPtr(k.Trmr, atr.Test) // Trmr is a term lexer.
	r.FldExt("scn.Scnr")              // use Ext to avoid Ptr mod
	return r
}
func (x *FleTrmr) InitTypFn() {
	x.LexSpce()
	x.LexCmnt()
	x.SkpSpceCmnt()
	x.LexIdn()
	x.LexObjs()
	x.Prefixs()
	x.LexTrm(k.Asn)
	x.LexTrm(k.Each)
	x.LexTrm(k.PllEach)
	x.LexTrm(k.PllWait)
	x.LexTrm(k.Then)
	x.LexTrm(k.Else)
}
func (x *FleTrmr) StructLit(name string) (r *Struct) {
	r = x.Structf("%vLit", atr.None, name)
	r.FldTyp(_sys.Bsc.Bnd)
	return r
}
func (x *FleTrmr) StructLitf(format string, args ...interface{}) (r *Struct) {
	r = x.Structf("%vLit", atr.None, fmt.Sprintf(format, args...))
	r.FldTyp(_sys.Bsc.Bnd)
	return r
}
func (x *FleTrmr) StructArrLit(name string, elmLit *Struct) (r *Struct) {
	r = x.Structf("%vLit", atr.None, name)
	r.FldTyp(_sys.Bsc.Bnd)
	r.FldSlice("Elms", elmLit)
	return r
}
func (x *FleTrmr) MemLit(lit *Struct) (r *TypFn) {
	r = x.TypFn(lit.Name)
	r.OutPrm(lit, "r")
	r.OutPrm(Bool, "ok")
	return r
}
func (x *FleTrmr) LexTrm(trm string) (r *TypFn) {
	trm = sys.Camel(trm)
	if r, exists := x.Lexs[trm]; exists {
		return r
	}
	r = x.TypFn(trm)
	x.Lexs[trm] = r
	r.OutPrm(_sys.Bsc.Bnd, "r")
	r.OutPrm(Bool, "ok")
	ch, size, idx := rune(0), 0, 0
	for idx < len(trm) {
		ch, size = utf8.DecodeRuneInString(trm[idx:])
		r.Addf("if x.Ch != %q {", ch)
		r.Add("return r, false")
		r.Add("}")
		if idx == 0 {
			r.Add("r.Idx = x.Idx")
		}
		r.Add("x.NextRune()")
		idx += size
	}
	r.Add("if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {")
	r.Add("return r, false")
	r.Add("}")
	r.Add("r.Lim = x.Idx")
	r.Add("return r, true")
	if x.Test != nil && Opt.IsTest() {
		x.Test.Import(_sys.Bsc.Bnd)
		x.Test.Trm(trm, _sys.Bsc.Bnd.Typ().Ref(x), []string{trm})
	}
	return r
}
func (x *FleTrmr) CallLit(fn, trm *TypFn, lit *Struct) {
	fn.Addf("%v, ok := x.%v()", trm.Camel(), trm.Name)
	fn.Add("if ok {")
	fn.Addf("return %v{Bnd: %v}, true", lit.Title(), trm.Camel())
	fn.Add("}")
	fn.Add("x.Scn = scn")
}
func (x *FleTrmr) CallTrm(fn, trm *TypFn) {
	fn.Addf("%v, ok := x.%v()", trm.Camel(), trm.Name)
	fn.Add("if ok {")
	fn.Addf("return %v, true", trm.Camel())
	fn.Add("}")
	fn.Add("x.Scn = scn")
}

func (x *FleTrmr) LexSpce() (r *TypFn) {
	x.Spce = x.StructLit(k.Spce)
	x.Import("unicode")
	r = x.MemLit(x.Spce)
	r.Add("if !unicode.IsSpace(x.Ch) {")
	r.Add("return r, false")
	r.Add("}")
	r.Add("r.Idx = x.Idx")
	r.Add("for !x.NextRune() && unicode.IsSpace(x.Ch) {")
	r.Add("}")
	r.Add("r.Lim = x.Idx")
	r.Add("return r, true")
	if x.Test != nil && Opt.IsTest() {
		x.Test.LitTrm(x.Spce, x.LitSpce(), TxtGen, TxtGen|PrefixGen)
	}
	return r
}
func (x *FleTrmr) LitSpce() (r []string) {
	// /usr/local/go/src/unicode/graphic.go
	//	'\t', '\\n', '\v', '\f', '\r', ' ', U+0085 (NEL), U+00A0 (NBSP).
	r = []string{"\t", "\n", "\v", "\f", "\r", " ", string(rune(0x85)), string(rune(0xA0))}
	return r
}

func (x *FleTrmr) LexCmnt() (r *TypFn) {
	x.Cmnt = x.StructLit(k.Cmnt)
	r = x.MemLit(x.Cmnt)
	// r.Add("if x.Ch != '//' {")
	r.Add("if x.Ch != '/' || (x.Ch == '/' && x.PeekRune() != '/') {")
	r.Add("return r, false")
	r.Add("}")
	r.Add("r.Idx = x.Idx")
	r.Add("for !x.NextRune() && x.Ch != '\\n' {")
	r.Add("}")
	r.Add("if x.Ch == '\\n' {")
	r.Add("x.NextRune()")
	r.Add("}")
	r.Add("r.Lim = x.Idx")
	r.Add("return r, true")
	if x.Test != nil && Opt.IsTest() {
		x.Test.LitTrm(x.Cmnt, x.LitCmnt(), TxtGen, TxtGen|PrefixGen)
	}
	return r
}
func (x *FleTrmr) LitCmnt() (r []string) {
	r = []string{"//", "////", "//a", "//0", "// ", "// abc123", "//\n", "// comment\n"}
	return r
}

func (x *FleTrmr) SkpSpceCmnt() (r *TypFn) {
	r = x.TypFn("SkpSpceCmnt")
	r.Add("// skip all consecutive spce and cmnt")
	r.Add("scn, spceOk, cmntOk := x.Scn, true, true")
	r.Add("for spceOk || cmntOk {")
	r.Addf("_, spceOk = x.%v()", x.Spce.Ref(x))
	r.Add("if spceOk {")
	r.Add("scn = x.Scn")
	r.Add("} else {")
	r.Add("x.Scn = scn")
	r.Add("}")
	r.Addf("_, cmntOk = x.%v()", x.Cmnt.Ref(x))
	r.Add("if cmntOk {")
	r.Add("scn = x.Scn")
	r.Add("} else {")
	r.Add("x.Scn = scn")
	r.Add("}")
	r.Add("}")
	return r
}
func (x *FleTrmr) LexIdn() (r *TypFn) {
	x.Idn = x.StructLit(k.Idn)
	r = x.MemLit(x.Idn)
	r.Add("if !(unicode.IsLetter(x.Ch) || x.Ch == '_') {")
	r.Add("return r, false")
	r.Add("}")
	r.Add("r.Idx = x.Idx")
	r.Add("x.NextRune()")
	r.Add("for unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {")
	r.Add("x.NextRune()")
	r.Add("}")
	r.Add("r.Lim = x.Idx")
	r.Add("return r, true")
	if x.Test != nil && Opt.IsTest() {
		x.Test.LitTrm(x.Idn, x.LitIdn(), TxtGen|SuffixGen, IdnGen|PrefixGen)
	}
	return r
}
func (x *FleTrmr) LitIdn() (r []string) {
	r = []string{"_", "_a", "_0", "a", "a_", "a0", "abc123"}
	return r
}

func (x *FleTrmr) LexTmePrt() (r *TypFn) {
	x.TmePrt = x.StructLit("TmePrt")
	r = x.MemLit(x.TmePrt)
	r.InPrm(Rune, "suffix")
	r.Add("r.Idx = x.Idx")
	r.Add("if x.Ch == '-' { // optional minus")
	r.Add("x.NextRune()")
	r.Add("if !unicode.IsDigit(x.Ch) { // next ch must be digit")
	r.Add("return r, false")
	r.Add("}")
	r.Add("x.NextRune()")
	r.Add("}")
	r.Add("for unicode.IsDigit(x.Ch) {")
	r.Add("x.NextRune()")
	r.Add("}")
	r.Add("if r.Idx == x.Idx || (x.Idx-r.Idx == 1 && x.Txt[r.Idx] == '-') {")
	r.Add("return r, false")
	r.Add("}")
	r.Add("if x.Ch != suffix { // 1s, -3s, 1m, -3m, 1h, -3h, 1d, -3d, 1w, -3w")
	r.Add("return r, false")
	r.Add("}")
	r.Add("x.NextRune()")
	r.Add("if unicode.IsLetter(x.Ch) || x.Ch == '_' {")
	r.Add("return r, false")
	r.Add("}")
	r.Add("r.Lim = x.Idx")
	r.Add("return r, true")
	_sys.Lng.Pro.Trm.Prs.PrsTmePrtTrm()
	return r
}
func (x *FleTrmr) StructTme() (r *Struct) {
	x.LexTmePrt()
	r = x.StructLit(k.Tme)
	r.Fld(strings.Title(k.Year), x.TmePrt)
	r.Fld(strings.Title(k.Month), x.TmePrt)
	r.Fld(strings.Title(k.Week), x.TmePrt)
	r.Fld(strings.Title(k.Day), x.TmePrt)
	r.Fld(strings.Title(k.Hour), x.TmePrt)
	r.Fld(strings.Title(k.Minute), x.TmePrt)
	r.Fld(strings.Title(k.Second), x.TmePrt)
	// _sys.Lng.Pro.Trm.Prs.PrsTmeTrm(r)
	return r
}
func (x *FleTrmr) StructBnd() (r *Struct) {
	untBse := _sys.Bsc.Unt.Typ().Bse()
	r = x.StructLit(k.Bnd)
	r.Fld("IdxTrm", untBse.LitTrm)
	r.Fld("LimTrm", untBse.LitTrm)
	return r
}
func (x *FleTrmr) StructRng(elm *TypBse) (r *Struct) {
	r = x.StructLitf("%vRng", elm.Title())
	// fmt.Println("***", "elm.LitTrm == nil", elm.LitTrm == nil)
	r.Fld("MinTrm", elm.LitTrm)
	r.Fld("MaxTrm", elm.LitTrm)
	return r
}
func (x *FleTrmr) InitCtor(s *Struct) {
	s.NmeLex = x.LexTrm(s.Camel())
	s.LitTrm = x.StructLit(s.Title())
	for _, fld := range s.Flds {
		if fld.IsTrm() { // exported fld with lit
			s.LitTrm.Fld(fld.Name, fld.Typ.Bse().LitTrm)
			fld.Lex = x.LexTrm(fld.Camel())
		}
	}
	s.LitLex = x.LexStruct(s)
	if x.Test != nil && Opt.IsTest() {
		x.Test.LitTrmTyp(s.Bse())
	}
}
func (x *FleTrmr) LexStruct(s *Struct) (r *TypFn) {
	r = x.TypFnf("%vLit", s.Title())
	r.OutPrm(s.LitTrm, "r")
	r.OutPrm(Bool, "ok")
	r.Add("r.Idx = x.Idx")

	trm := fmt.Sprintf("%v.%v(", s.Pkg.Name, s.Camel())
	ch, size, idx := rune(0), 0, 0
	for idx < len(trm) {
		ch, size = utf8.DecodeRuneInString(trm[idx:])
		r.Addf("if x.Ch != %q {", ch)
		r.Add("return r, false")
		r.Add("}")
		if idx == 0 {
			r.Add("r.Idx = x.Idx")
		}
		r.Add("x.NextRune()")
		idx += size
	}

	r.Add("for !x.End && x.Ch != ')' {")
	r.Add("x.SkpSpceCmnt()")
	r.Add("scn := x.Scn")
	abcFlds := make([][]func(), 26) // gather flds by fst ch
	for _, fld := range s.Flds {
		if fld.IsTrm() {
			fld := fld // capture for closure
			idx := fld.Lex.Camel()[0] - 'a'
			abcFlds[idx] = append(abcFlds[idx], func() {
				r.Addf("if _, ok := x.%v(); ok {", fld.Lex.Title())
				r.Add("x.SkpSpceCmnt()")
				r.Add("if x.Ch != ':' {")
				r.Add("return r, false")
				r.Add("}")
				r.Add("x.NextRune()")
				r.Add("x.SkpSpceCmnt()")
				r.Addf("r.%v, ok = x.%v()", fld.Name, fld.Typ.Bse().LitLex.Title())
				r.Add("if !ok {")
				r.Add("return r, false")
				r.Add("}")
				r.Add("continue")
				r.Add("}")
				r.Add("x.Scn = scn // rewind")
			})
		}
	}
	r.Add("switch x.Ch {")
	for m, curFlds := range abcFlds {
		if len(curFlds) > 0 {
			r.Addf("case '%v':", string(rune(m)+'a'))
			for _, fn := range curFlds {
				fn()
			}
		}
	}
	r.Add("}")
	r.Add("break // no matching flds")

	r.Add("}") // endfor
	r.Add("if x.Ch != ')' {")
	r.Add("return r, false")
	r.Add("}")
	r.Add("x.NextRune()")
	r.Add("r.Lim = x.Idx")
	r.Add("return r, true")
	return r
}
func (x *FleTrmr) LexArrLit(arr *Arr) {
	arr.LitTrm = x.StructArrLit(arr.PrefixTitle(), arr.Elm.LitTrm)
	arr.LitLex = x.MemLit(arr.LitTrm)
	WriteLexArr(arr.LitLex, arr.Elm.LitTrm)
	if x.Test != nil && Opt.IsTest() {
		x.Test.LitTrm(arr.LitTrm, arr.LitsNonEmp(), None)
	}
}
func WriteLexArr(fn *TypFn, elmLit *Struct) {
	fn.Add("if x.Ch != '[' {")
	fn.Add("return r, false")
	fn.Add("}")
	fn.Add("r.Idx = x.Idx")
	fn.Add("x.NextRune()")
	fn.Add("scn := x.Scn")
	fn.Add("for {")
	fn.Add("x.SkpSpceCmnt()")
	fn.Add("scn = x.Scn")
	fn.Addf("elm, ok := x.%v()", elmLit.Title())
	fn.Add("if !ok {")
	fn.Add("x.Scn = scn")
	fn.Add("break")
	fn.Add("}")
	fn.Add("r.Elms = append(r.Elms, elm)")
	fn.Add("}")
	fn.Add("if len(r.Elms) == 0 {")
	fn.Add("return r, false")
	fn.Add("}")
	fn.Add("x.SkpSpceCmnt()")
	fn.Add("if x.End || x.Ch != ']' {")
	fn.Add("return r, false")
	fn.Add("}")
	fn.Add("x.NextRune()")
	fn.Add("if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {")
	fn.Add("return r, false")
	fn.Add("}")
	fn.Add("r.Lim = x.Idx")
	fn.Add("return r, true")
}

func (x *FleTrmr) LexObjs() (r *TypFn) { // non-lit arr; ana arr
	x.Objs = x.Struct("ObjsLit", atr.None)
	x.Objs.FldTyp(_sys.Bsc.Bnd)
	r = x.MemLit(x.Objs)
	r.Add("if x.Ch != '[' {")
	r.Add("return r, false")
	r.Add("}")
	r.Add("r.Idx = x.Idx")
	r.Add("x.SkpSet('[', ']')")
	r.Add("r.Lim = x.Idx")
	r.Add("return r, true")
	if x.Test != nil && Opt.IsTest() {
		x.Test.LitTrm(x.Objs, x.LitObjs(), TxtGen, TxtGen|PrefixGen)
	}
	return r
}
func (x *FleTrmr) Prefixs() (r *TypFn) { // non-lit arr; ana arr
	r = x.TypFn("Prefixs")
	r.InPrm(String, "idn")
	r.OutPrmSlice(_sys.Bsc.Bnd, "r")
	r.Add("for !x.End {")
	r.Add("x.SkpSpceCmnt()")
	r.Add("if x.Ch != rune(idn[0]) {")
	r.Add("x.NextRune()")
	r.Add("continue")
	r.Add("}")
	r.Add("curBnd := bnd.Bnd{Idx: x.Idx, Lim: x.Idx+1}")
	r.Add("for n := 1; n < len(idn); n++ {")
	r.Add("x.NextRune()")
	r.Add("if x.Ch == rune(idn[n]) {")
	r.Add("curBnd.Lim++")
	r.Add("continue")
	r.Add("}")
	r.Add("break")
	r.Add("}")
	r.Add("if int(curBnd.Lim-curBnd.Idx) == len(idn) {")
	r.Add("x.NextRune()")
	r.Add("x.SkpSpceCmnt()")
	r.Add("if x.Ch == '.' {")
	r.Add("x.NextRune()")
	r.Add("x.SkpSpceCmnt()")
	r.Add("_, ok := x.IdnLit()")
	r.Add("if ok {")
	r.Add("x.SkpSpceCmnt()")
	r.Add("x.SkpSet('(', ')')")
	r.Add("curBnd.Lim = x.Idx")
	r.Add("r = append(r, curBnd)")
	r.Add("}")
	r.Add("}")
	r.Add("}")
	r.Add("}")
	r.Add("return r")
	return r
}
func (x *FleTrmr) LitObjs() (r []string) {
	r = []string{
		"[ana.oan().hst().eurUsd().s1().bids().fst()]",
		"[ana.oan().hst().eurUsd().s1().bids().fst() ana.oan().hst().eurUsd().s1().bids().lst()]",
		"[ana.oan().hst().eurUsd().s1().bids().fst() ana.oan().hst().eurUsd().s1().bids().lst() ana.oan().hst().eurUsd().s1().bids().sum()]",
	}
	return r
}
