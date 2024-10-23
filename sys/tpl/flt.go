package tpl

import (
	"sys"
	"sys/k"
	"sys/tpl/atr"
	"sys/tpl/mod"
)

const (
	FltSize = 4
)

type (
	FleFlt struct {
		FleBse
		PrtRel
		PrtSgn
		PrtAri
		PrtSel
		PrtString
		PrtBytes
		// PrtLog
		// PrtIfc
		Zero     *Cnst
		One      *Cnst
		NegOne   *Cnst
		Hndrd    *Cnst
		Min      *Cnst
		Max      *Cnst
		Tiny     *Cnst
		Size     *Cnst
		Trnc     *TypFn
		IsNaN    *TypFn
		IsInfPos *TypFn
		IsInfNeg *TypFn
		IsValid  *TypFn
	}
	FleFlts struct {
		FleBse
		PrtArr
		// PrtArrIdn
		PrtArrRel
		PrtArrSrt
		PrtArrSer
		PrtArrUna
		PrtArrScl
		PrtArrSel
		PrtArrCnt
		PrtArrInr
		// TODO:
		// PrtArrOtr
		PrtArrAgg
		PrtArrStrWrt
		PrtArrBytWrt
		PrtBytes
		// PrtLog
		// PrtIfc
	}
)

func (x *DirBsc) NewFlt() (r *FleFlt) {
	r = &FleFlt{}
	x.Flt = r
	r.Name = k.Flt
	r.Pkg = x.Pkg.New(r.Name)
	r.Alias(r.Name, Float32, atr.TypFlt)
	r.AddFle(r)
	return r
}
func (x *FleFlt) NewArr() (r *FleFlts) {
	r = &FleFlts{}
	r.FleBse = *NewArr(x, &r.PrtArr)
	r.AddFle(r)
	return r
}
func (x *FleFlt) InitVals(bse *TypBse) {
	bse.Lits = sys.Vs("0.0", "1.1", "3.0", "3.0", "3.0", "99999.99", "-1.1", "-99999.99") // add 3.0 repetition for gap fil simulation
	bse.Vals = bse.Lits
	bse.LitsJsn = bse.Lits
	bse.ValsJsn = bse.Vals
}
func (x *FleFlt) InitCnst() {
	x.Zero = x.Cnst(k.Zero, "0.0")
	x.One = x.Cnst(k.One, "1.0")
	x.NegOne = x.Cnst(k.NegOne, "-1.0")
	x.Hndrd = x.Cnst(k.Hndrd, "100.0")
	x.Min = x.Cnst(k.Min, "-3.40282346638528859811704183484516925440e+38")
	x.Max = x.Cnst(k.Max, "3.40282346638528859811704183484516925440e+38")
	x.Tiny = x.Cnst(k.Tiny, "1.401298464324817070923729583289916131280e-45")
	x.Size = x.CnstSize(FltSize)
}
func (x *FleFlt) InitTypFn() {
	// x.RemTypFn(k.Eql)
	x.eql()
	x.neq()
	// if x.Tst != nil {
	// 	x.Tst.IdnPrt()
	// }
	x.Trnc = x.trunc()
	x.IsNaN = x.isNaN()
	x.IsInfPos = x.isInfPos()
	x.IsInfNeg = x.isInfNeg()
	x.IsValid = x.isValid()
	x.pct()
	x.strWrt()
	x.bytWrt()
	x.BytRed()
}
func (x *FleFlt) InitTrm(bse *TypBse, trmr *FleTrmr) {
	x.InitLexLit(func(r *TypFn, t *FleTrmr) {
		r.InPrmVariadic(Bool, "skpDsh")
		r.Add("// TODO: NaN, +Inf, -Inf")
		r.Add("r.Idx = x.Idx")
		r.Add("if x.Ch == '-' { // optional minus")
		r.Add("x.NextRune()")
		r.Add("if !unicode.IsDigit(x.Ch) { // next ch must be digit")
		r.Add("return r, false")
		r.Add("}")
		r.Add("x.NextRune()")
		r.Add("}")
		r.Add("for !x.End && unicode.IsDigit(x.Ch) {")
		r.Add("x.NextRune()")
		r.Add("}")
		r.Add("if x.Ch != '.' {")
		r.Add("return r, false")
		r.Add("}")
		r.Add("x.NextRune()")
		r.Add("if !unicode.IsDigit(x.Ch) { // next ch must be digit")
		r.Add("return r, false")
		r.Add("}")
		r.Add("for !x.End && unicode.IsDigit(x.Ch) {")
		r.Add("x.NextRune()")
		r.Add("}")
		r.Add("if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' || (len(skpDsh) == 0 && x.Ch == '-') {")
		r.Add("return r, false")
		r.Add("}")
		r.Add("r.Lim = x.Idx")
		r.Add("return r, r.Idx != r.Lim")
	})
	x.InitPrsLit(func(r *PkgFn, f *FlePrs) {
		f.Import("strconv")
		f.Import("sys/err")
		r.Add("v, er := strconv.ParseFloat(txt[trm.Idx:trm.Lim], 32)")
		r.Add("if er != nil {")
		r.Addf("err.Panicf(\"%v: failed to parse (txt:%%q err:%%q)\", txt, er)", x.Typ().Title())
		r.Add("}")
		r.Add("return flt.Flt(v)")
	})
	x.InitPrsCfg()
	x.InitJsnLexLit(func(r *TypFn, t *FleJsnTrmr) {
		r.Add("// TODO: NaN, +Inf, -Inf")
		r.Add("r.Idx = x.Idx")
		r.Add("if x.Ch == '-' { // optional minus")
		r.Add("x.NextRune()")
		r.Add("if !unicode.IsDigit(x.Ch) { // next ch must be digit")
		r.Add("return r, false")
		r.Add("}")
		r.Add("x.NextRune()")
		r.Add("}")
		r.Add("for !x.End && unicode.IsDigit(x.Ch) {")
		r.Add("x.NextRune()")
		r.Add("}")
		r.Add("if x.Ch != '.' {")
		r.Add("return r, false")
		r.Add("}")
		r.Add("x.NextRune()")
		r.Add("if !unicode.IsDigit(x.Ch) { // next ch must be digit")
		r.Add("return r, false")
		r.Add("}")
		r.Add("for !x.End && unicode.IsDigit(x.Ch) {")
		r.Add("x.NextRune()")
		r.Add("}")
		r.Add("if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {")
		r.Add("return r, false")
		r.Add("}")
		r.Add("r.Lim = x.Idx")
		r.Add("return r, r.Idx != r.Lim")
	})
	x.InitJsnPrsLit(func(r *PkgFn, f *FleJsnPrs) {
		f.Import("strconv")
		f.Import("sys/err")
		r.Add("v, er := strconv.ParseFloat(txt[trm.Idx:trm.Lim], 32)")
		r.Add("if er != nil {")
		r.Addf("err.Panicf(\"%v: failed to parse (txt:%%q err:%%q)\", txt, er)", x.Typ().Title())
		r.Add("}")
		r.Add("return flt.Flt(v)")
	})
	x.InitJsnPrs()
}
func (x *FleFlt) LexLit(r *TypFn, t *FleTrmr) {
	r.Add("// TODO: NaN, +Inf, -Inf")
	r.Add("r.Idx = x.Idx")
	r.Add("if x.Ch == '-' { // optional minus")
	r.Add("x.NextRune()")
	r.Add("if !unicode.IsDigit(x.Ch) { // next ch must be digit")
	r.Add("return r, false")
	r.Add("}")
	r.Add("x.NextRune()")
	r.Add("}")
	r.Add("for !x.End && unicode.IsDigit(x.Ch) {")
	r.Add("x.NextRune()")
	r.Add("}")
	r.Add("if x.Ch != '.' {")
	r.Add("return r, false")
	r.Add("}")
	r.Add("x.NextRune()")
	r.Add("if !unicode.IsDigit(x.Ch) { // next ch must be digit")
	r.Add("return r, false")
	r.Add("}")
	r.Add("for !x.End && unicode.IsDigit(x.Ch) {")
	r.Add("x.NextRune()")
	r.Add("}")
	r.Add("if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {")
	r.Add("return r, false")
	r.Add("}")
	r.Add("r.Lim = x.Idx")
	r.Add("return r, r.Idx != r.Lim")
}
func (x *FleFlt) eql() (r *TypFn) {
	r = x.TypFn(k.Eql)
	r.InPrm(x, "a")
	r.OutPrm(_sys.Bsc.Bol)
	r.Addf("return %v == %v || %[1]v.IsNaN() && %[2]v.IsNaN()", r.Rxr.Name, r.InPrms[0].Name)
	// test
	if x.Test != nil && Opt.IsTest() {
		x.Test.ImportFn(r)
		r.T.Addf("expected := %v(cse.x == cse.a || cse.x.IsNaN() && cse.a.IsNaN())", r.OutTyp().Full())
	}
	return r
}
func (x *FleFlt) neq() (r *TypFn) {
	r = x.TypFn(k.Neq)
	r.InPrm(x, "a")
	r.OutPrm(_sys.Bsc.Bol)
	r.Addf("return %v != %v", r.Rxr.Name, r.InPrms[0].Name)
	// test
	if x.Test != nil && Opt.IsTest() {
		x.Test.ImportFn(r)
		r.T.Addf("expected := %v(cse.x != cse.a)", r.OutTyp().Full())
	}
	return r
}
func (x *FleFlt) trunc() (r *TypFn) {
	x.Import("strconv")
	x.Import("sys/err")
	r = x.TypFna(k.Trnc, atr.Lng)
	r.InPrm(_sys.Bsc.Unt, "precision")
	r.OutPrm(x)
	r.Add("if x != x { // IsNaN")
	r.Add("return x")
	r.Add("}")
	r.Add("s := x.String()")
	r.Add("idx := strings.Index(s, \".\")")
	r.Add("if len(s)-idx-1 <= int(precision) {")
	r.Add("return x")
	r.Add("}")
	r.Add("v, er := strconv.ParseFloat(s[:idx+1+int(precision)], 32)")
	r.Add("if er != nil {")
	r.Add("err.Panicf(\"Flt: failed to parse (txt:%q err:%q s:%v)\", s[:idx+1+int(precision)], er, s)")
	r.Add("}")
	r.Addf("return %v(v)", x.Title())
	// test
	if x.Test != nil && Opt.IsTest() {
		x.Test.Import("strconv")
		x.Test.Import("strings")
		r.T.Addf("expected := cse.x")
		r.T.Add("s := x.String()")
		r.T.Add("idx := strings.Index(s, \".\")")
		r.T.Add("if len(s)-idx-1 > int(precision) {")
		r.T.Add("v, _ := strconv.ParseFloat(s[:idx+1+int(precision)], 32)")
		r.T.Addf("expected = %v(v)", x.Typ().Full())
		r.T.Add("}")
	}
	return r
}
func (x *FleFlt) isNaN() (r *TypFn) {
	r = x.TypFn("IsNaN")
	r.OutPrm(_sys.Bsc.Bol)
	r.Add("// From /usr/local/go/src/math/bits.go")
	r.Add("return x != x")
	// test
	if x.Test != nil && Opt.IsTest() {
		r.T.Addf("expected := %v(cse.x != cse.x)", r.OutTyp().Full())
	}
	return r
}
func (x *FleFlt) isInfPos() (r *TypFn) {
	r = x.TypFn("IsInfPos")
	r.OutPrm(_sys.Bsc.Bol)
	r.Add("// From /usr/local/go/src/math/bits.go")
	r.Add("return x > Max")
	// test
	if x.Test != nil && Opt.IsTest() {
		r.T.Addf("expected := %v(cse.x > %v)", r.OutTyp().Full(), x.Max.Full())
	}
	return r
}
func (x *FleFlt) isInfNeg() (r *TypFn) {
	r = x.TypFn("IsInfNeg")
	r.OutPrm(_sys.Bsc.Bol)
	r.Add("// From /usr/local/go/src/math/bits.go")
	r.Add("return x < Min")
	// test
	if x.Test != nil && Opt.IsTest() {
		r.T.Addf("expected := %v(cse.x < %v)", r.OutTyp().Full(), x.Min.Full())
	}
	return r
}
func (x *FleFlt) isValid() (r *TypFn) {
	r = x.TypFn("IsValid")
	r.OutPrm(_sys.Bsc.Bol)
	r.Add("// From /usr/local/go/src/math/bits.go")
	r.Add("return x == x && x >= Min && x <= Max")
	// test
	if x.Test != nil && Opt.IsTest() {
		r.T.Addf("expected := %v(cse.x == cse.x && cse.x >= %v && cse.x <= %v)", r.OutTyp().Full(), x.Min.Full(), x.Max.Full())
	}
	return r
}
func (x *FleFlt) pct() (r *TypFn) {
	r = x.TypFn(k.Pct)
	r.InPrm(x, "v")
	r.OutPrm(x, "r")
	r.Add("if v - x == 0 {")
	r.Add("return 0")
	r.Add("}")
	r.Add("if x == 0 {")
	r.Add("return 1")
	r.Add("}")
	r.Add("if x < 0 {")
	r.Add("return ((v - x) / -x)")
	r.Add("}")
	r.Add("return ((v - x) / x)")
	return r
}
func (x *FleFlt) strWrt() (r *TypFn) {
	x.Import("strconv")
	x.Import("strings")
	r = x.TypFn("StrWrt")
	r.InPrm(BuilderPtr, "b")
	r.Add("s := strconv.FormatFloat(float64(x), byte('f'), -1, 32)")
	r.Add("b.WriteString(s)")
	r.Add("if x.IsValid() && strings.LastIndex(s, \".\") < 0 {")
	r.Add("b.WriteString(\".0\")")
	r.Add("}")
	return r
}
func (x *FleFlt) bytWrt() (r *TypFn) {
	x.Import("encoding/binary")
	r = x.TypFn("BytWrt")
	r.InPrm(BufferPtr, "b")
	r.Add("v := make([]byte, Size)")
	r.Add("binary.LittleEndian.PutUint32(v, *(*uint32)(unsafe.Pointer(&x)))")
	r.Add("b.Write(v)")
	return r
}
func (x *FleFlt) BytRed() (r *TypFn) {
	x.Import("unsafe")
	r = x.TypFn("BytRed")
	r.Rxr.Mod = mod.Ptr
	r.InPrmSlice(Byte, "b")
	r.OutPrm(Int)
	r.Add("bits := binary.LittleEndian.Uint32(b[:4])")
	r.Addf("*x = *(*%v)(unsafe.Pointer(&bits))", r.Rxr.Typ.Title())
	r.Add("return Size")
	return r
}
