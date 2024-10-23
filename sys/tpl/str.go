package tpl

import (
	"sys"
	"sys/k"
	"sys/tpl/atr"
	"sys/tpl/mod"
)

type (
	FleStr struct {
		FleBse
		PrtIdn
		PrtRel
		PrtString
		PrtBytes
	}
	FleStrs struct {
		FleBse
		PrtArr
		PrtArrRel
		PrtArrSrt
		PrtArrStrWrt
	}
)

func (x *DirBsc) NewStr() (r *FleStr) {
	r = &FleStr{}
	x.Str = r
	r.Name = k.Str
	r.Pkg = x.Pkg.New(r.Name)
	r.Alias(r.Name, String, atr.TypStr)
	r.AddFle(r)
	return r
}
func (x *FleStr) NewArr() (r *FleStrs) {
	r = &FleStrs{}
	r.FleBse = *NewArr(x, &r.PrtArr)
	r.AddFle(r)
	return r
}
func (x *FleStr) InitVals(bse *TypBse) {
	bse.Lits = sys.Vs("\"\"", "\"xYz\"", "\"a\"", "\"efg HIJ jKl\"")
	bse.Vals = bse.Lits
	bse.LitsJsn = bse.Lits
	bse.ValsJsn = bse.Vals
}
func (x *FleStr) InitCnst() {
	x.Cnst(k.Zero, "\"\"")
	x.Cnst(k.Empty, "\"\"")
}
func (x *FleStr) InitPkgFn() {
	x.Fmt()
}

func (x *FleStr) Fmt() (r *PkgFn) {
	x.Import("fmt")
	r = x.PkgFna("Fmt", atr.Lng)
	r.InPrm(x, "tmpl")
	r.InPrmVariadic(Interface, "vs")
	r.OutPrm(x)
	r.Addf("return %v(fmt.Sprintf(string(tmpl), vs...))", x.Typ().Ref(x))
	return r
}
func (x *FleStr) InitTypFn() {
	x.lower()
	x.upper()
	x.unquo()
	x.strWrt()
	x.bytWrt()
	x.bytRed()
}
func (x *FleStr) InitTrm(bse *TypBse, trmr *FleTrmr) {
	x.InitLexLit(func(r *TypFn, f *FleTrmr) {
		r.Add("r.Idx = x.Idx")
		r.Add("if x.Ch != '\"' {")
		r.Add("return r, false")
		r.Add("}")
		r.Add("x.NextRune()")
		r.Add("for !x.End && x.Ch != '\"' {")
		r.Add("if x.Ch == '\\\\' { // escape sequence: start")
		r.Add("x.NextRune()")
		r.Add("switch x.Ch {")
		r.Add("case 'a', 'b', 'f', 'n', 'r', 't', 'v', '\\\\', '\"':")
		r.Add("default:")
		r.Add("return r, false // unknown escape sequence")
		r.Add("}")
		r.Add("}")
		r.Add("x.NextRune()")
		r.Add("}")
		r.Add("if x.End || x.Ch != '\"' {")
		r.Add("return r, false")
		r.Add("}")
		r.Add("x.NextRune()")
		r.Add("if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {")
		r.Add("return r, false")
		r.Add("}")
		r.Add("r.Lim = x.Idx")
		r.Add("return r, true")
	})
	x.InitPrsLit(func(r *PkgFn, f *FlePrs) {
		r.Addf("return %v(txt[trm.Idx+1 : trm.Lim-1])", x.Typ().Ref(f))
	})
	x.InitPrsCfg()
	x.InitJsnLexLit(func(r *TypFn, f *FleJsnTrmr) {
		r.Add("r.Idx = x.Idx")
		r.Add("if x.Ch != '\"' {")
		r.Add("return r, false")
		r.Add("}")
		r.Add("x.NextRune()")
		r.Add("for !x.End && x.Ch != '\"' {")
		r.Add("if x.Ch == '\\\\' { // escape sequence: start")
		r.Add("x.NextRune()")
		r.Add("switch x.Ch {")
		r.Add("case 'a', 'b', 'f', 'n', 'r', 't', 'v', '\\\\', '\"':")
		r.Add("default:")
		r.Add("return r, false // unknown escape sequence")
		r.Add("}")
		r.Add("}")
		r.Add("x.NextRune()")
		r.Add("}")
		r.Add("if x.End || x.Ch != '\"' {")
		r.Add("return r, false")
		r.Add("}")
		r.Add("x.NextRune()")
		r.Add("if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {")
		r.Add("return r, false")
		r.Add("}")
		r.Add("r.Lim = x.Idx")
		r.Add("return r, true")
	})
	x.InitJsnPrsLit(func(r *PkgFn, f *FleJsnPrs) {
		r.Addf("return %v(txt[trm.Idx+1 : trm.Lim-1])", x.Typ().Ref(f))
	})
	x.InitJsnPrs()
}
func (x *FleStr) strWrt() (r *TypFn) {
	r = x.TypFn("StrWrt")
	r.InPrm(BuilderPtr, "b")
	r.Add("b.WriteRune('\"')")
	r.Add("b.WriteString(string(x))")
	r.Add("b.WriteRune('\"')")
	return r
}
func (x *FleStr) bytWrt() (r *TypFn) {
	r = x.TypFn("BytWrt")
	r.InPrm(BufferPtr, "b")
	r.Add("bLen := make([]byte, 4)")
	r.Add("binary.LittleEndian.PutUint32(bLen, uint32(len(x)))")
	r.Add("b.Write(bLen)      // string length")
	r.Add("b.Write([]byte(x)) // string content")
	return r
}
func (x *FleStr) bytRed() (r *TypFn) {
	x.Import("encoding/binary")
	r = x.TypFn("BytRed")
	r.Rxr.Mod = mod.Ptr
	r.InPrmSlice(Byte, "b")
	r.OutPrm(Int)
	r.Add("sLen := binary.LittleEndian.Uint32(b[:4])")
	r.Add("if sLen > 0 {")
	r.Add("*x = Str(string(b[4:4+sLen]))")
	r.Add("}")
	r.Add("return 4 + int(sLen)")
	return r
}
func (x *FleStr) unquo() (r *TypFn) {
	r = x.TypFn("Unquo")
	r.OutPrm(String)
	r.Add("return string(x)")
	return r
}
func (x *FleStr) lower() (r *TypFn) {
	x.Import("strings")
	r = x.TypFn("Lower")
	r.OutPrm(x)
	r.Add("return Str(strings.ToLower(string(x)))")
	// test
	if x.Test != nil && Opt.IsTest() {
		x.Test.Import("strings")
		r.T.Addf("expected := %v(strings.ToLower(string(cse.x)))", r.OutTyp().Ref(x.Test))
	}
	return r
}
func (x *FleStr) upper() (r *TypFn) {
	r = x.TypFn("Upper")
	r.OutPrm(x)
	r.Add("return Str(strings.ToUpper(string(x)))")
	// test
	if x.Test != nil && Opt.IsTest() {
		x.Test.Import("strings")
		r.T.Addf("expected := %v(strings.ToUpper(string(cse.x)))", r.OutTyp().Ref(x.Test))
	}
	return r
}

// arr
func (x *FleStrs) InitTypFn() {
	x.bytWrt()
	x.bytRed()
}
func (x *FleStrs) bytWrt() (r *TypFn) {
	x.Import("encoding/binary")
	r = x.TypFn(k.BytWrt)
	r.InPrm(BufferPtr, "b")
	r.Add("bLen := make([]byte, 4) // array length")
	r.Add("binary.LittleEndian.PutUint32(bLen, uint32(len(*x)))")
	r.Add("b.Write(bLen)")
	r.Add("for _, v := range *x {")
	r.Add("binary.LittleEndian.PutUint32(bLen, uint32(len(v)))")
	r.Add("b.Write(bLen) // string length")
	r.Add("b.WriteString(string(v)) // current string")
	r.Add("}")
	return r
}
func (x *FleStrs) bytRed() (r *TypFn) {
	r = x.TypFn(k.BytRed)
	r.InPrmSlice(Byte, "b")
	r.OutPrm(Int, "r")
	r.Add("if len(b) >= 4 {")
	r.Add("cnt := int(binary.LittleEndian.Uint32(b[:4]))")
	r.Add("idx := 4")
	r.Add("for n := 0; n < cnt; n++ {")
	r.Add("vLen := int(binary.LittleEndian.Uint32(b[idx : idx+4]))")
	r.Add("*x = append(*x, str.Str(string(b[idx+4:idx+4+vLen])))")
	r.Add("idx += 4 + vLen")
	r.Add("r += vLen")
	r.Add("}")
	r.Add("}")
	r.Add("return r")
	return r
}
