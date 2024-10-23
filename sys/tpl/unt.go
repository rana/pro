package tpl

import (
	"sys"
	"sys/k"
	"sys/tpl/atr"
	"sys/tpl/mod"
)

const (
	UntSize = 4
)

type (
	FleUnt struct {
		FleBse
		PrtIdn
		PrtRel
		PrtAri
		PrtString
		PrtBytes
		// PrtLog
		// PrtIfc
		Zero      *Cnst
		One       *Cnst
		Min       *Cnst
		Max       *Cnst
		Size      *Cnst
		MinSegLen *Cnst
	}
	FleUnts struct {
		FleBse
		PrtArr
		// PrtArrIdn
		PrtArrRel
		PrtArrSrt
		PrtArrSer
		PrtArrInr
		PrtArrAgg
		PrtArrStrWrt
		PrtArrBytWrt
		// PrtLog
		// PrtIfc
	}
)

func (x *DirBsc) NewUnt() (r *FleUnt) {
	r = &FleUnt{}
	x.Unt = r
	r.Name = k.Unt
	r.Pkg = x.Pkg.New(r.Name)
	r.Alias(r.Name, Uint32, atr.TypUnt)
	r.AddFle(r)
	return r
}
func (x *FleUnt) NewArr() (r *FleUnts) {
	r = &FleUnts{}
	r.FleBse = *NewArr(x, &r.PrtArr)
	r.AddFle(r)
	return r
}
func (x *FleUnt) InitVals(bse *TypBse) {
	bse.Lits = sys.Vs("0", "1", "1000", "10")
	bse.Vals = bse.Lits
}
func (x *FleUnt) InitCnst() {
	x.Zero = x.Cnst(k.Zero, "0")
	x.One = x.Cnst(k.One, "1")
	x.Min = x.Cnst(k.Min, "0")
	x.Max = x.Cnst(k.Max, "1<<32-1")
	x.Size = x.CnstSize(UntSize)
	x.MinSegLen = x.Cnst("MinSegLen", "64")
	x.MinSegLen.Atr = atr.None
}
func (x *FleUnt) InitTypFn() {
	x.strWrt()
	x.bytWrt()
	x.BytRed()
}
func (x *FleUnt) InitTrm(bse *TypBse, trmr *FleTrmr) {
	x.InitLexLit(func(r *TypFn, t *FleTrmr) {
		r.Add("r.Idx = x.Idx")
		r.Add("for unicode.IsDigit(x.Ch) {")
		r.Add("x.NextRune()")
		r.Add("}")
		r.Add("if x.Ch == '-' || unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' { // check dash for tme/bnd")
		r.Add("return r, false")
		r.Add("}")
		r.Add("r.Lim = x.Idx")
		r.Add("return r, r.Idx != r.Lim")
	})
	x.InitPrsLit(func(r *PkgFn, f *FlePrs) {
		r.Add("lit := txt[trm.Idx:trm.Lim]")
		r.Addf("mag := %v(1)", x.Typ().Full())
		r.Add("for n := len(lit) - 1; n > -1; n-- {")
		r.Addf("r += mag * %v(lit[n]-'0')", x.Typ().Full())
		r.Addf("mag *= %v(10)", x.Typ().Full())
		r.Add("}")
		r.Add("return r")
	})
	x.InitPrsCfg()
}

func (x *FleUnt) strWrt() (r *TypFn) {
	r = x.TypFn("StrWrt")
	r.InPrm(BuilderPtr, "b")
	r.Add("if x == 0 { // TODO: OPTIMIZE")
	r.Add("b.WriteRune('0')")
	r.Add("} else {")
	r.Add("var rs []rune")
	r.Add("for x != 0 {")
	r.Add("rs = append(rs, '0'+rune(x%10))")
	r.Add("x /= 10")
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
func (x *FleUnt) bytWrt() (r *TypFn) {
	x.Import("encoding/binary")
	r = x.TypFn("BytWrt")
	r.InPrm(BufferPtr, "b")
	r.Add("v := make([]byte, Size)")
	r.Add("binary.LittleEndian.PutUint32(v, uint32(x))")
	r.Add("b.Write(v)")
	return r
}
func (x *FleUnt) BytRed() (r *TypFn) {
	r = x.TypFn("BytRed")
	r.Rxr.Mod = mod.Ptr
	r.InPrmSlice(Byte, "b")
	r.OutPrm(Int)
	r.Add("*x = Unt(binary.LittleEndian.Uint32(b[:Size]))")
	r.Add("return Size")
	return r
}
