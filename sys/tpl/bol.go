package tpl

import (
	"sys"
	"sys/k"
	"sys/tpl/atr"
	"sys/tpl/mod"
)

const (
	BolSize = 1
)

type (
	FleBol struct {
		FleBse
		PrtIdn
		PrtString
		PrtBytes
	}
	FleBols struct {
		FleBse
		PrtArr
		PrtArrStrWrt
		PrtArrBytWrt
	}
)

func (x *DirBsc) NewBol() (r *FleBol) {
	r = &FleBol{}
	x.Bol = r
	r.Name = k.Bol
	r.Pkg = x.Pkg.New(r.Name)
	r.Alias(r.Name, Bool, atr.TypBol)
	r.AddFle(r)
	return r
}
func (x *FleBol) NewArr() (r *FleBols) {
	r = &FleBols{}
	r.FleBse = *NewArr(x, &r.PrtArr)
	r.AddFle(r)
	return r
}
func (x *FleBol) InitVals(bse *TypBse) {
	bse.Lits = sys.Vs(k.Fls, k.Tru)
	bse.Vals = sys.Vs(k.False, k.True)
	bse.LitsJsn = sys.Vs(k.False, k.True)
	bse.ValsJsn = bse.Vals
}
func (x *FleBol) InitCnst() {
	x.Cnst(k.Zero, k.False)
	x.Cnst(k.Fls, k.False)
	x.Cnst(k.Tru, k.True)
	x.CnstSize(BolSize)
}
func (x *FleBol) InitTypFn() {
	x.Not()
	x.strWrt()
	x.bytWrt()
	x.BytRed()
}
func (x *FleBol) InitTrm(bse *TypBse, trmr *FleTrmr) {
	x.InitLexLit(func(r *TypFn, t *FleTrmr) {
		r.Add("scn := x.Scn")
		t.CallLit(r, t.LexTrm(k.Fls), bse.LitTrm)
		t.CallLit(r, t.LexTrm(k.Tru), bse.LitTrm)
		r.Add("return r, false")
	})
	x.InitPrsLit(func(r *PkgFn, f *FlePrs) {
		f.Import("sys/k")
		r.Addf("return txt[trm.Idx:trm.Lim] == k.Tru")
	})
	x.InitPrsCfg()
	x.InitJsnLexLit(func(r *TypFn, t *FleJsnTrmr) {
		r.Add("scn := x.Scn")
		t.CallLit(r, t.LexTrm(k.False), bse.LitTrm)
		t.CallLit(r, t.LexTrm(k.True), bse.LitTrm)
		r.Add("return r, false")
	})
	x.InitJsnPrsLit(func(r *PkgFn, f *FleJsnPrs) {
		f.Import("sys/k")
		r.Addf("return txt[trm.Idx:trm.Lim] == k.True")
	})
	x.InitJsnPrs()
}
func (x *FleBol) Not() (r *TypFn) {
	r = x.TypFn(k.Not)
	r.OutPrm(x)
	r.Add("return !x")
	// test
	r.T.Add("expected := !cse.x")
	return r
}
func (x *FleBol) strWrt() (r *TypFn) {
	x.Import("sys/k")
	r = x.TypFn(k.StrWrt)
	r.InPrm(BuilderPtr, "b")
	r.Add("if x {")
	r.Add("b.WriteString(k.Tru)")
	r.Add("} else {")
	r.Add("b.WriteString(k.Fls)")
	r.Add("}")
	return r
}
func (x *FleBol) bytWrt() (r *TypFn) {
	r = x.TypFn(k.BytWrt)
	r.InPrm(BufferPtr, "b")
	r.Add("if x {")
	r.Add("b.WriteByte(1)")
	r.Add("} else {")
	r.Add("b.WriteByte(0)")
	r.Add("}")
	return r
}
func (x *FleBol) BytRed() (r *TypFn) {
	r = x.TypFn(k.BytRed)
	r.Rxr.Mod = mod.Ptr
	r.InPrmSlice(Byte, "b")
	r.OutPrm(Int)
	r.Add("*x = Bol(b[0] == 1)")
	r.Add("return Size")
	return r
}
