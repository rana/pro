package tpl

import (
	"sys/k"
	"sys/tpl/atr"
)

type (
	FleHstTrdsStm struct {
		FleHstBse
	}
)

func (x *DirHst) NewTrdsStm() (r *FleHstTrdsStm) {
	r = &FleHstTrdsStm{}
	// x.TrdsStm = r
	r.Name = k.TrdsStm
	r.Pkg = x.Pkg
	r.StructPtr(r.Name, atr.HstTrdsStm)
	r.AddFle(r)
	return r
}

func (x *FleHstTrdsStm) InitTyp(bse *TypBse) {

}

// func (x *FleHstTrdsStm) InitFld(s *Struct) {
// 	// s.FldTyp(_sys.Ana.PthBse).Atr = atr.BytLitStrEqlBqSkp
// 	s.Fld("Margin", _sys.Bsc.Unt)
// 	s.Fld("GapMax", _sys.Bsc.Tme)
// 	s.Fld("Stm", _sys.Ana.Hst.Stm).Atr = atr.BytLitStrEqlBqSkp
// 	s.Fld("TrdStmSegs", _sys.Ana.Hst.TrdStmSegs)
// }
func (x *FleHstTrdsStm) InitTypFn() {
	x.CrsInvSum()
	x.CrsInvSumKey()
}
func (x *FleHstTrdsStm) CrsInvSum() (r *TypFn) {
	x.Import(_sys)
	x.Import("sys/trc")
	// seg
	seg := x.StructPtrf("%vSeg", atr.None, k.CrsInvSum)
	seg.FldTyp(_sys.Bsc.Bnd)
	seg.Fld("RngFul", _sys.Bsc.Flt)
	seg.Fld("Vals", _sys.Bsc.Flt.arr)
	seg.Fld("Out", _sys.Bsc.Flt.arr)
	// seg act
	segAct := x.TypFn("Act", seg)
	segAct.Add("for m := x.Idx; m < x.Lim; m++ {")
	segAct.Add("for n := unt.Zero; n < x.Vals.Cnt(); n++ {")
	segAct.Add("if m != n && x.Vals.At(n).Neq(flt.Max) {")
	segAct.Add("// out[m] += 1 - (abs(m-n)/rngFul)")
	segAct.Add("x.Out.Upd(m, x.Out.At(m)+flt.One.Sub(x.Vals.At(m).Sub(x.Vals.At(n)).Pos().Div(x.RngFul)))")
	segAct.Add("}")
	segAct.Add("}")
	segAct.Add("}")
	// CrsInvSum: TypFn
	r = x.TypFna(k.CrsInvSum, atr.Lng)
	r.InPrm(_sys.Bsc.Unt, "idx")
	r.OutPrm(_sys.Bsc.Flt.arr, "r")
	r.Add("trcr := trc.New(x.String())")
	r.Add("defer trcr.End()")
	r.Add("if sys.HasDsk() { // load from lcl dsk")
	r.Add("if val := sys.Dsk().LoadTrdsStmCrsInvSum(x.CrsInvSumKey(idx)); val != nil {")
	r.Add("r = flts.New()")
	r.Add("r.BytRed(val)")
	r.Add("return r")
	r.Add("}")
	r.Add("}")
	r.Add("trcr = trc.New(\"CrsInvSum.TrdStmSegs\")")
	r.Addf("r = %v(x.TrdStmSegs.Cnt())", _sys.Bsc.Flt.arr.Make.Ref(x))
	r.Addf("vals := %v(x.TrdStmSegs.Cnt())", _sys.Bsc.Flt.arr.Make.Ref(x))
	r.Add("for n := unt.Zero; n < x.TrdStmSegs.Cnt(); n++ {")
	r.Add("trdStmSeg := x.TrdStmSegs.At(n)")
	r.Add("if idx < trdStmSeg.Vals.Cnt() {")
	r.Add("vals.Upd(n, trdStmSeg.Vals.At(idx)) // populate cross-section vals")
	r.Add("} else {")
	r.Add("vals.Upd(n, flt.Max) // some segs may not be full")
	r.Add("trcr.Cnt++")
	r.Add("}")
	r.Add("}")
	r.Add("trcr.End()")
	r.Add("rngFul := vals.RngFul()")
	r.Addf("segBnds, acts := %v(x.TrdStmSegs.Cnt())", _sys.Bsc.Bnd.arr.Segs.Ref(x))
	r.Add("for n, segBnd := range *segBnds {")
	r.Addf("acts[n] = %v{", seg.Adr(x))
	r.Add("Bnd: segBnd,")
	r.Add("RngFul: rngFul,")
	r.Add("Vals: vals,")
	r.Add("Out: r,")
	r.Add("}")
	r.Add("}")
	r.Add("trcr = trc.New(\"CrsInvSum.Process\")")
	r.Add("sys.Run().Pll(acts...)")
	r.Add("trcr.End()")
	r.Add("if sys.HasDsk() { // sav to lcl dsk")
	r.Add("trcr = trc.New(\"CrsInvSum.SavDsk\")")
	r.Add("sys.Dsk().SavTrdsStmCrsInvSum(x.CrsInvSumKey(idx), r.Bytes())")
	r.Add("trcr.End()")
	r.Add("}")
	r.Add("return r")
	return r
}

func (x *FleHstTrdsStm) CrsInvSumKey() (r *TypFn) {
	x.Import("bytes")
	r = x.TypFna(k.CrsInvSum+"Key", atr.None)
	r.InPrm(_sys.Bsc.Unt, "idx")
	r.OutPrmSlice(Byte)
	r.Add("b := &bytes.Buffer{}")
	r.Add("b.WriteString(x.String())")
	r.Add("idx.BytWrt(b)")
	r.Add("return b.Bytes()")
	return r
}
