package tpl

import (
	"sys/k"
	"sys/tpl/atr"
)

type (
	FlePltStgy struct {
		FleBse
		PrtPlt
		PrtStructFldSet
	}
)

func (x *DirPlt) NewStgy() (r *FlePltStgy) {
	r = &FlePltStgy{}
	// x.Stgy = r
	r.Name = k.Stgy
	r.Pkg = x.Pkg
	r.StructPtr(r.Name, atr.TypUiStruct)
	r.AddFle(r)
	return r
}
func (x *FlePltStgy) InitTyp(bse *TypBse) {
	x.ImplIfc(_sys.Ana.Vis.Plt.Plt.Typ().(*Ifc))
}
func (x *FlePltStgy) InitFld(s *Struct) {
	s.FldTyp(NewExt("TmeFltPltBse"))
	// s.Fld("X", _sys.Ana.Vis.Plt.TmeAxisX).Atr = atr.Get
	// s.Fld("Y", _sys.Ana.Vis.Plt.FltAxisY).Atr = atr.Get
	// s.Fld("Sampl", _sys.Bsc.Bol).Atr = atr.Get
	s.Fld("Title", _sys.Bsc.Str).Atr = atr.SetGet
	s.Fld("stgy", _sys.Ana.Hst.Stgy.bse)
	s.Fld("stm", _sys.Ana.Hst.Stm.bse)
	s.Fld("Pos", _sys.Bsc.Bol).Atr = atr.SetGet
	s.Fld("Neg", _sys.Bsc.Bol).Atr = atr.SetGet
	s.FldSlice("stmStkTrds", NewExt("*StmStkTrd"))
	s.FldSlice("stmStkTrdRndrs", NewExt("*StmStkTrdRndrSeg"))
}
func (x *FlePltStgy) InitPkgFn() {
	x.New()
}
func (x *FlePltStgy) New() (r *PkgFn) {
	r = x.PkgFna("NewStgy", atr.Lng)
	r.OutPrm(x, "r")
	r.Addf("r = %v{}", x.Typ().Adr(x))
	r.Add("r.PltBse = NewPltBse(r)")
	r.Add("r.x = NewTmeAxisX()")
	r.Add("r.y = NewFltAxisY()")
	r.Add("r.sampl = true")
	r.Add("r.mrgn = Mrgn // glbl mrgn")
	r.Add("r.Pos = true")
	r.Add("r.x.Vis(true)")
	r.Add("r.Neg = true")
	r.Add("return r")
	return r
}
func (x *FlePltStgy) InitTypFn() {
	x.X()
	x.Y()
	x.Stgy()
	x.Stm()
	x.HrzSclVal()
	x.VrtSclVal()
}
func (x *FlePltStgy) X() (r *TypFn) {
	r = x.TypFn("X")
	r.OutPrm(_sys.Ana.Vis.Plt.TmeAxisX)
	r.Add("return x.x")
	return r
}
func (x *FlePltStgy) Y() (r *TypFn) {
	r = x.TypFn("Y")
	r.OutPrm(_sys.Ana.Vis.Plt.FltAxisY)
	r.Add("return x.y")
	return r
}
func (x *FlePltStgy) Stgy() (r *TypFn) {
	r = x.TypFn(k.Stgy)
	r.InPrm(_sys.Ana.Hst.Stgy, "stgy")
	r.OutPrm(x)
	r.Add("x.stgy = stgy.Bse()")
	r.Add("return x")
	return r
}
func (x *FlePltStgy) Stm() (r *TypFn) {
	r = x.TypFn(k.Stm)
	r.InPrm(_sys.Ana.Hst.Stm, "stm")
	r.OutPrm(x)
	r.Add("x.stm = stm.Bse()")
	r.Add("return x")
	return r
}
func (x *FlePltStgy) HrzSclVal() (r *TypFn) {
	r = x.TypFn(k.HrzSclVal)
	r.InPrm(_sys.Bsc.Tme, "val")
	r.OutPrm(x)
	r.Add("x.X().PxlPerVal = 1.0 / float32(val)")
	r.Add("return x")
	return r
}
func (x *FlePltStgy) VrtSclVal() (r *TypFn) {
	r = x.TypFn(k.VrtSclVal)
	r.InPrm(_sys.Bsc.Flt, "val")
	r.OutPrm(x)
	r.Add("x.Y().PxlPerVal = 1.0 / float32(val)")
	r.Add("return x")
	return r
}
