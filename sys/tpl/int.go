package tpl

import (
	"sys"
	"sys/k"
	"sys/tpl/atr"
	"sys/tpl/mod"
)

const (
	IntSize = 4
)

type (
	FleInt struct {
		FleBse
		PrtIdn
		PrtRel
		PrtSgn
		PrtAri
		PrtString
		PrtBytes
		// PrtLog
		// PrtIfc
	}
	FleInts struct {
		FleBse
		PrtArr
		// PrtArrIdn
		PrtArrRel
		PrtArrSrt
		PrtArrStrWrt
		PrtArrBytWrt
		// PrtLog
		// PrtIfc
	}
)

func (x *DirBsc) NewInt() (r *FleInt) {
	r = &FleInt{}
	x.Int = r
	r.Name = k.Int
	r.Pkg = x.Pkg.New(r.Name)
	r.Alias(r.Name, Int32, atr.TypInt)
	r.AddFle(r)
	return r
}
func (x *FleInt) NewArr() (r *FleInts) {
	r = &FleInts{}
	r.FleBse = *NewArr(x, &r.PrtArr)
	r.AddFle(r)
	return r
}
func (x *FleInt) InitVals(bse *TypBse) {
	bse.Lits = sys.Vs("+0", "+10", "+1000", "-10", "-1000")
	bse.Vals = sys.Vs("0", "10", "1000", "-10", "-1000")
	bse.LitsJsn = bse.Vals
	bse.ValsJsn = bse.Vals
}
func (x *FleInt) InitCnst() {
	x.Cnst(k.Zero, "0")
	x.Cnst(k.One, "1")
	x.Cnst(k.NegOne, "-1")
	x.Cnst(k.Min, "-1 << 31")
	x.Cnst(k.Max, "1<<31-1")
	x.CnstSize(IntSize)
}
func (x *FleInt) InitTypFn() {
	x.strWrt()
	x.bytWrt()
	x.BytRed()
}
func (x *FleInt) InitTrm(bse *TypBse, trmr *FleTrmr) {
	x.InitLexLit(func(r *TypFn, t *FleTrmr) {
		r.Add("r.Idx = x.Idx")
		r.Add("if x.Ch != '+' && x.Ch != '-' {")
		r.Add("return r, false")
		r.Add("}")
		r.Add("x.NextRune()")
		r.Add("if !unicode.IsDigit(x.Ch) { // ch must be digit")
		r.Add("return r, false")
		r.Add("}")
		r.Add("x.NextRune()")
		r.Add("for !x.End && unicode.IsDigit(x.Ch) {")
		r.Add("x.NextRune()")
		r.Add("}")
		r.Add("if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {")
		r.Add("return r, false")
		r.Add("}")
		r.Add("r.Lim = x.Idx")
		r.Add("return r, r.Idx != r.Lim")
	})
	x.InitPrsLit(func(r *PkgFn, f *FlePrs) {
		r.Add("lit := txt[trm.Idx+1:trm.Lim]")
		r.Addf("mag := %v(1)", x.Typ().Ref(f))
		r.Add("for n := len(lit) - 1; n > -1; n-- {")
		r.Addf("r += mag * %v(lit[n]-'0')", x.Typ().Ref(f))
		r.Addf("mag *= %v(10)", x.Typ().Ref(f))
		r.Add("}")
		r.Add("if txt[trm.Idx:trm.Idx+1] == \"-\" {")
		r.Add("r = -r")
		r.Add("}")
		r.Add("return r")
	})
	x.InitPrsCfg()
	x.InitJsnLexLit(func(r *TypFn, t *FleJsnTrmr) {
		r.Add("r.Idx = x.Idx")
		r.Add("if x.Ch == '-' { // optional minus")
		r.Add("x.NextRune()")
		r.Add("if !unicode.IsDigit(x.Ch) { // ch must be digit")
		r.Add("return r, false")
		r.Add("}")
		r.Add("x.NextRune()")
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
		r.Add("var lit string")
		r.Add("if txt[trm.Idx:trm.Idx+1] == \"-\" {")
		r.Add("lit = txt[trm.Idx+1:trm.Lim]")
		r.Add("} else {")
		r.Add("lit = txt[trm.Idx:trm.Lim]")
		r.Add("}")
		r.Addf("mag := %v(1)", x.Typ().Ref(f))
		r.Add("for n := len(lit) - 1; n > -1; n-- {")
		r.Addf("r += mag * %v(lit[n]-'0')", x.Typ().Ref(f))
		r.Addf("mag *= %v(10)", x.Typ().Ref(f))
		r.Add("}")
		r.Add("if txt[trm.Idx:trm.Idx+1] == \"-\" {")
		r.Add("r = -r")
		r.Add("}")
		r.Add("return r")
	})
	x.InitJsnPrs()
}

func (x *FleInt) strWrt() (r *TypFn) {
	r = x.TypFn("StrWrt")
	r.InPrm(BuilderPtr, "b")
	r.Add("if x == 0 {")
	r.Add("b.WriteString(\"+0\")")
	r.Add("} else {")
	r.Add("var rs []rune")
	r.Add("if x < 0 {")
	r.Add("b.WriteRune('-')")
	r.Add("for x != 0 { // separate branch to support math.MinInt32")
	r.Add("rs = append(rs, '0'+rune(-(x%10)))")
	r.Add("x /= 10")
	r.Add("}")
	r.Add("} else {")
	r.Add("b.WriteRune('+')")
	r.Add("for x != 0 {")
	r.Add("rs = append(rs, '0'+rune(x%10))")
	r.Add("x /= 10")
	r.Add("}")
	r.Add("}")
	r.Add("for i, j := 0, len(rs)-1; i < j; i, j = i+1, j-1 { // reverse")
	r.Add("rs[i], rs[j] = rs[j], rs[i]")
	r.Add("}")
	r.Add("for n := 0; n < len(rs); n++ {")
	r.Add("b.WriteRune(rs[n])")
	r.Add("}")
	r.Add("}")
	return r
}
func (x *FleInt) bytWrt() (r *TypFn) {
	x.Import("encoding/binary")
	x.Import("unsafe")
	r = x.TypFn("BytWrt")
	r.InPrm(BufferPtr, "b")
	r.Add("v := make([]byte, Size)")
	r.Add("binary.LittleEndian.PutUint32(v, *(*uint32)(unsafe.Pointer(&x)))")
	r.Add("b.Write(v)")
	return r
}
func (x *FleInt) BytRed() (r *TypFn) {
	x.Import("unsafe")
	r = x.TypFn("BytRed")
	r.Rxr.Mod = mod.Ptr
	r.InPrmSlice(Byte, "b")
	r.OutPrm(Int)
	r.Add("bits := binary.LittleEndian.Uint32(b[:Size])")
	r.Add("*x = *(*Int)(unsafe.Pointer(&bits))")
	r.Add("return Size")
	return r
}
