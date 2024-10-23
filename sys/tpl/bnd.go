package tpl

import (
	"sys"
	"sys/k"
	"sys/tpl/atr"
	"sys/tpl/mod"
)

const (
	BndSize = 2 * UntSize
)

type (
	FleBnd struct {
		FleBse
		// PrtStructIdn
		PrtString
		PrtBytes
		PrtPkt
		// PrtLog
		// PrtIfc
	}
	FleBnds struct {
		FleBse
		PrtArr
		// PrtArrIdn
		PrtArrStrWrt
		PrtString
		PrtArrBytWrt
		// PrtLog
		// PrtIfc
	}
)

func (x *DirBsc) NewBnd() (r *FleBnd) {
	r = &FleBnd{}
	x.Bnd = r
	r.Name = k.Bnd
	r.Pkg = x.Pkg.New(r.Name)
	r.Struct(r.Name, atr.TypBnd)
	r.AddFle(r)
	return r
}
func (x *FleBnd) NewArr() (r *FleBnds) {
	r = &FleBnds{}
	r.FleBse = *NewArr(x, &r.PrtArr)
	r.AddFle(r)
	return r
}
func (x *FleBnd) InitFld(s *Struct) {
	s.Fld("Idx", _sys.Bsc.Unt).Atr = atr.TstSkp
	s.Fld("Lim", _sys.Bsc.Unt)
}
func (x *FleBnd) InitVals(bse *TypBse) {
	bse.Lits = sys.Vs("0-0", "0-1", "0-1000", "999-1000", "1-0")
	bse.Vals = sys.VsStruct(x.Typ().Full(), "Idx:0, Lim:0", "Idx:0, Lim:1", "Idx:0, Lim:1000", "Idx:999, Lim:1000", "Idx:1, Lim:0")
}
func (x *FleBnd) InitCnst() {
	x.CnstSize(BndSize)
}
func (x *FleBnd) InitTypFn() {
	x.Cnt()
	x.Len()
	x.LstIdx()
	x.IsValid()
	x.strWrt()
	x.bytWrt()
	x.BytRed()
}
func (x *FleBnd) InitTrm(bse *TypBse, trmr *FleTrmr) {
	x.InitLexLit(func(r *TypFn, t *FleTrmr) {
		r.Add("r.Idx = x.Idx")
		r.Add("for unicode.IsDigit(x.Ch) {")
		r.Add("x.NextRune()")
		r.Add("}")

		r.Add("if x.Ch != '-' {")
		r.Add("return r, false")
		r.Add("}")
		r.Add("r.IdxTrm.Idx = r.Idx")
		r.Add("r.IdxTrm.Lim = x.Idx")
		r.Add("x.NextRune()")
		r.Add("r.LimTrm.Idx = x.Idx")

		r.Add("for unicode.IsDigit(x.Ch) {")
		r.Add("x.NextRune()")
		r.Add("}")
		r.Add("if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {")
		r.Add("return r, false")
		r.Add("}")

		r.Add("r.LimTrm.Lim = x.Idx")
		r.Add("r.Lim = x.Idx")
		r.Add("return r, true")
	}, trmr.StructBnd())
	x.InitPrsLit(func(r *PkgFn, f *FlePrs) {
		untBse := _sys.Bsc.Unt.Typ().Bse()
		r.Addf("r.Idx = %v(trm.IdxTrm, txt)", untBse.PrsTrm.Ref(f))
		r.Addf("r.Lim = %v(trm.LimTrm, txt)", untBse.PrsTrm.Ref(f))
		r.Add("return r")
	})
	x.InitPrsCfg()
}
func (x *FleBnd) Cnt() (r *TypFn) {
	r = x.TypFn("Cnt")
	r.OutPrm(_sys.Bsc.Unt)
	r.Add("return x.Lim - x.Idx")
	// test
	if x.Test != nil && Opt.IsTest() {
		r.T.Add("expected := cse.x.Lim - cse.x.Idx")
	}
	return r
}
func (x *FleBnd) Len() (r *TypFn) {
	r = x.TypFn(k.Len)
	r.OutPrm(_sys.Bsc.Unt)
	r.Add("return x.Lim-x.Idx")
	// test
	if x.Test != nil && Opt.IsTest() {
		x.Test.ImportFn(r)
		r.T.Add("expected := cse.x.Lim-cse.x.Idx")
	}
	return r
}
func (x *FleBnd) LstIdx() (r *TypFn) {
	r = x.TypFn("LstIdx")
	r.OutPrm(_sys.Bsc.Unt)
	r.Add("if x.Lim == 0 {")
	r.Add("return 0")
	r.Add("}")
	r.Add("return x.Lim-1")
	// test
	if x.Test != nil && Opt.IsTest() {
		x.Test.ImportFn(r)
		r.T.Addf("var expected %v", r.OutTyp().Full())
		r.T.Add("if cse.x.Lim > 0 {")
		r.T.Add("expected = cse.x.Lim-1")
		r.T.Add("}")
	}
	return r
}
func (x *FleBnd) IsValid() (r *TypFn) {
	r = x.TypFn("IsValid")
	r.OutPrm(_sys.Bsc.Bol)
	r.Add("return x.Idx < x.Lim") // IMPORTANT: USE < FOR hst.Instr
	// test
	if x.Test != nil && Opt.IsTest() {
		x.Test.ImportFn(r)
		r.T.Addf("expected := %v(cse.x.Idx < cse.x.Lim)", r.OutTyp().Full()) // test
	}
	return r
}
func (x *FleBnd) strWrt() (r *TypFn) {
	r = x.TypFn("StrWrt")
	r.InPrm(BuilderPtr, "b")
	r.Add("x.Idx.StrWrt(b)")
	r.Add("b.WriteRune('-')")
	r.Add("x.Lim.StrWrt(b)")
	return r
}
func (x *FleBnd) bytWrt() (r *TypFn) {
	r = x.TypFn("BytWrt")
	r.InPrm(BufferPtr, "b")
	r.Add("x.Idx.BytWrt(b)")
	r.Add("x.Lim.BytWrt(b)")
	return r
}
func (x *FleBnd) BytRed() (r *TypFn) {
	r = x.TypFn("BytRed")
	r.Rxr.Mod = mod.Ptr
	r.InPrmSlice(Byte, "b")
	r.OutPrm(Int)
	r.Add("idx := 0")
	r.Add("idx += x.Idx.BytRed(b[idx:])")
	r.Add("x.Lim.BytRed(b[idx:])")
	r.Add("return Size")
	return r
}

func (x *FleBnds) InitTypFn() {
	x.PrtArr.Arr.Segs = x.segs()
}
func (x *FleBnds) segs() (r *PkgFn) {
	x.Import(_sys)
	x.Import("runtime")
	r = x.PkgFn(k.Segs)
	r.InPrm(_sys.Bsc.Unt, "elmCnt")
	r.OutPrm(x, "r")
	r.OutPrmSlice(NewExt("sys.Act"), "acts")
	r.Add("r = New()")
	r.Add("segCnt := unt.Unt(runtime.NumCPU())")
	r.Add("segLen := elmCnt.Div(segCnt)")
	r.Addf("if segLen.Lss(%v) {", _sys.Bsc.Unt.MinSegLen.Ref(x))
	r.Add("segLen = elmCnt")
	r.Add("segCnt = unt.One")
	r.Add("}")
	r.Add("idx, lim := unt.Zero, segLen")
	r.Add("for n := unt.Zero; n < segCnt; n++ {")
	r.Add("if n > 0 {")
	r.Add("idx = lim")
	r.Add("lim += segLen")
	r.Add("}")
	r.Add("if idx >= elmCnt {")
	r.Add("break")
	r.Add("}")
	r.Add("if n == segCnt-1 || lim > elmCnt {")
	r.Add("lim = elmCnt")
	r.Add("}")
	r.Add("r.Push(bnd.Bnd{Idx: idx, Lim: lim})")
	r.Add("}")
	r.Add("return r, make([]sys.Act, len(*r))")
	return r
}
