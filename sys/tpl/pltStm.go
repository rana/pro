package tpl

import (
	"sys/k"
	"sys/tpl/atr"
)

type (
	FlePltStm struct {
		FleBse
		PrtPlt
		PrtStructFldSet
	}
)

func (x *DirPlt) NewStm() (r *FlePltStm) {
	r = &FlePltStm{}
	x.Stm = r
	r.Name = k.Stm
	r.Pkg = x.Pkg
	r.StructPtr(r.Name, atr.TypUiStruct)
	r.AddFle(r)
	return r
}
func (x *FlePltStm) InitTyp(bse *TypBse) {
	x.ImplIfc(_sys.Ana.Vis.Plt.Plt.Typ().(*Ifc))
}
func (x *FlePltStm) InitFld(s *Struct) {
	s.FldTyp(NewExt("TmeFltPltBse"))
	// s.Fld("X", _sys.Ana.Vis.Plt.TmeAxisX).Atr = atr.Get
	// s.Fld("Y", _sys.Ana.Vis.Plt.FltAxisY).Atr = atr.Get
	// s.Fld("Sampl", _sys.Bsc.Bol).Atr = atr.Get
	s.Fld("Title", _sys.Bsc.Str).Atr = atr.SetGet
	s.FldSlice("stmStks", NewExt("*StmStk"))
	s.FldSlice("stmBnds", NewExt("*StmBnd"))
	s.FldSlice("cndStks", NewExt("*CndStk"))
	s.FldSlice("hrzLns", NewExt("*HrzLn"))
	s.FldSlice("vrtLns", NewExt("*VrtLn"))
	s.FldSlice("hrzBnds", NewExt("*HrzBnd"))
	s.FldSlice("vrtBnds", NewExt("*VrtBnd"))
	s.FldSlice("stmStkRndrs", NewExt("*StmStkRndrSeg"))
	s.FldSlice("stmBndRndrs", NewExt("*StmBndRndrSeg"))
	s.FldSlice("cndStkRndrs", NewExt("*CndStkRndrSeg"))
	s.FldSlice("hrzLnRndrs", NewExt("*HrzLnRndrSeg"))
	s.FldSlice("vrtLnRndrs", NewExt("*VrtLnRndrSeg"))
	s.FldSlice("hrzBndRndrs", NewExt("*HrzBndRndrSeg"))
	s.FldSlice("vrtBndRndrs", NewExt("*VrtBndRndrSeg"))
}
func (x *FlePltStm) InitPkgFn() {
	x.New()
}
func (x *FlePltStm) New() (r *PkgFn) {
	r = x.PkgFna("NewStm", atr.Lng)
	r.OutPrm(x, "r")
	r.Addf("r = %v{}", x.Typ().Adr(x))
	r.Add("r.PltBse = NewPltBse(r)")
	r.Add("r.x = NewTmeAxisX()")
	r.Add("r.y = NewFltAxisY()")
	r.Add("r.sampl = true")
	r.Add("r.mrgn = Mrgn // glbl mrgn")
	r.Add("return r")
	return r
}
func (x *FlePltStm) InitTypFn() {
	x.X()
	x.Y()
	x.Stm()
	// x.StmFbr()
	x.StmBnd()
	x.Cnd()
	x.HrzLn()
	x.VrtLn()
	x.HrzBnd()
	x.VrtBnd()
	x.HrzSclVal()
	x.VrtSclVal()
}
func (x *FlePltStm) X() (r *TypFn) {
	r = x.TypFn("X")
	r.OutPrm(_sys.Ana.Vis.Plt.TmeAxisX)
	r.Add("return x.x")
	return r
}
func (x *FlePltStm) Y() (r *TypFn) {
	r = x.TypFn("Y")
	r.OutPrm(_sys.Ana.Vis.Plt.FltAxisY)
	r.Add("return x.y")
	return r
}
func (x *FlePltStm) Stm() (r *TypFn) {
	r = x.TypFn(k.Stm)
	r.InPrm(_sys.Ana.Vis.Pen.Pen, "pen")
	r.InPrmVariadic(_sys.Ana.Hst.Stm, "stms")
	r.OutPrm(x)
	r.Add("for _, stm := range stms {")
	r.Add("x.stmStks = append(x.stmStks, &StmStk{")
	r.Add("stm: stm.Bse(),")
	r.Add("pen: pen,")
	r.Add("plt: &x.TmeFltPltBse,")
	r.Add("})")
	r.Add("}")
	r.Add("return x")
	return r
}

// func (x *FlePltStm) StmFbr() (r *TypFn) {
// 	r = x.TypFn(k.StmFbr)
// 	r.InPrm(_sys.Ana.Vis.Pen.Pen.arr, "pens")
// 	r.InPrmVariadic(_sys.Ana.Hst.Stm.fbr, "stmFbrs")
// 	r.OutPrm(x)
// 	r.Add("for _, stmFbr := range stmFbrs {")
// 	r.Add("for n, stm := range *stmFbr.Stms() {")
// 	r.Add("x.stmStks = append(x.stmStks, &StmStk{")
// 	r.Add("stm: stm.Bse(),")
// 	r.Add("pen: (*pens)[n],")
// 	r.Add("plt: &x.TmeFltPltBse,")
// 	r.Add("})")
// 	r.Add("}")
// 	r.Add("}")
// 	r.Add("return x")
// 	return r
// }
func (x *FlePltStm) StmBnd() (r *TypFn) {
	r = x.TypFn(k.StmBnd)
	r.InPrm(_sys.Ana.Vis.Clr.Clr, "fil")
	r.InPrm(_sys.Ana.Vis.Pen.Pen, "stk")
	r.InPrm(_sys.Ana.Hst.Stm, "btm")
	r.InPrm(_sys.Ana.Hst.Stm, "top")
	r.OutPrm(x)
	r.Add("x.stmBnds = append(x.stmBnds, &StmBnd{")

	r.Add("btm: &StmStk{")
	r.Add("stm: btm.Bse(),")
	r.Add("pen: stk,")
	r.Add("plt: &x.TmeFltPltBse,")
	r.Add("},")

	r.Add("top: &StmStk{")
	r.Add("stm: top.Bse(),")
	r.Add("pen: stk,")
	r.Add("plt: &x.TmeFltPltBse,")
	r.Add("},")

	r.Add("filClr: fil,")
	r.Add("})")
	r.Add("return x")
	return r
}
func (x *FlePltStm) Cnd() (r *TypFn) {
	r = x.TypFn(k.Cnd)
	r.InPrm(_sys.Ana.Vis.Pen.Pen, "pen")
	r.InPrmVariadic(_sys.Ana.Hst.Cnd, "cnds")
	r.OutPrm(x)
	r.Add("for _, cnd := range cnds {")
	r.Add("x.cndStks = append(x.cndStks, &CndStk{")
	r.Add("cnd: cnd.Bse(),")
	r.Add("pen: pen,")
	r.Add("plt: &x.TmeFltPltBse,")
	r.Add("})")
	r.Add("}")
	r.Add("return x")
	return r
}
func (x *FlePltStm) HrzLn() (r *TypFn) {
	r = x.TypFn(k.HrzLn)
	r.InPrm(_sys.Ana.Vis.Pen.Pen, "pen")
	r.InPrmVariadic(_sys.Bsc.Flt, "ys")
	r.OutPrm(x)
	r.Add("for _, val := range ys {")
	r.Add("x.hrzLns = append(x.hrzLns, &HrzLn{")
	r.Add("val: val,")
	r.Add("pen: pen,")
	r.Add("plt: &x.TmeFltPltBse,")
	r.Add("})")
	r.Add("}")
	r.Add("return x")
	return r
}
func (x *FlePltStm) VrtLn() (r *TypFn) {
	r = x.TypFn(k.VrtLn)
	r.InPrm(_sys.Ana.Vis.Pen.Pen, "pen")
	r.InPrmVariadic(_sys.Bsc.Tme, "xs")
	r.OutPrm(x)
	r.Add("for _, val := range xs {")
	r.Add("x.vrtLns = append(x.vrtLns, &VrtLn{")
	r.Add("val: val,")
	r.Add("pen: pen,")
	r.Add("plt: &x.TmeFltPltBse,")
	r.Add("})")
	r.Add("}")
	r.Add("return x")
	return r
}
func (x *FlePltStm) HrzBnd() (r *TypFn) {
	r = x.TypFn(k.HrzBnd)
	r.InPrm(_sys.Ana.Vis.Clr.Clr, "fil")
	r.InPrm(_sys.Ana.Vis.Pen.Pen, "stk")
	r.InPrm(_sys.Bsc.Flt, "btm")
	r.InPrm(_sys.Bsc.Flt, "top")
	r.OutPrm(x)
	r.Add("x.hrzBnds = append(x.hrzBnds, &HrzBnd{")

	r.Add("btm: &HrzLn{")
	r.Add("val: btm,")
	r.Add("pen: stk,")
	r.Add("plt: &x.TmeFltPltBse,")
	r.Add("},")

	r.Add("top: &HrzLn{")
	r.Add("val: top,")
	r.Add("pen: stk,")
	r.Add("plt: &x.TmeFltPltBse,")
	r.Add("},")

	r.Add("filClr: fil,")
	r.Add("})")
	r.Add("return x")
	return r
}
func (x *FlePltStm) VrtBnd() (r *TypFn) {
	r = x.TypFn(k.VrtBnd)
	r.InPrm(_sys.Ana.Vis.Clr.Clr, "fil")
	r.InPrm(_sys.Ana.Vis.Pen.Pen, "stk")
	r.InPrm(_sys.Bsc.Tme, "lft")
	r.InPrm(_sys.Bsc.Tme, "rht")
	r.OutPrm(x)
	r.Add("x.vrtBnds = append(x.vrtBnds, &VrtBnd{")

	r.Add("lft: &VrtLn{")
	r.Add("val: lft,")
	r.Add("pen: stk,")
	r.Add("plt: &x.TmeFltPltBse,")
	r.Add("},")

	r.Add("rht: &VrtLn{")
	r.Add("val: rht,")
	r.Add("pen: stk,")
	r.Add("plt: &x.TmeFltPltBse,")
	r.Add("},")

	r.Add("filClr: fil,")
	r.Add("})")
	r.Add("return x")
	return r
}
func (x *FlePltStm) HrzSclVal() (r *TypFn) {
	r = x.TypFn(k.HrzSclVal)
	r.InPrm(_sys.Bsc.Tme, "val")
	r.OutPrm(x)
	r.Add("x.X().PxlPerVal = 1.0 / float32(val)")
	r.Add("return x")
	return r
}
func (x *FlePltStm) VrtSclVal() (r *TypFn) {
	r = x.TypFn(k.VrtSclVal)
	r.InPrm(_sys.Bsc.Flt, "val")
	r.OutPrm(x)
	r.Add("x.Y().PxlPerVal = 1.0 / float32(val)")
	r.Add("return x")
	return r
}
