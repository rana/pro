package tpl

import (
	"sys/k"
	"sys/tpl/atr"
)

type (
	FlePltFltsSctr struct {
		FleBse
		PrtPlt
		PrtStructFldSet
	}
)

func (x *DirPlt) NewFltsSctr() (r *FlePltFltsSctr) {
	r = &FlePltFltsSctr{}
	x.FltsSctr = r
	r.Name = k.FltsSctr
	r.Pkg = x.Pkg
	r.StructPtr(r.Name, atr.TypUiStruct)
	r.AddFle(r)
	return r
}
func (x *FlePltFltsSctr) InitTyp(bse *TypBse) {
	x.ImplIfc(_sys.Ana.Vis.Plt.Plt.Typ().(*Ifc))
}
func (x *FlePltFltsSctr) InitFld(s *Struct) {
	s.FldTyp(NewExt("PltBse"))
	s.Fld("XWidth", Uint32)
	s.Fld("Y", _sys.Ana.Vis.Plt.FltAxisY).Atr = atr.Get
	s.Fld("Title", _sys.Bsc.Str).Atr = atr.SetGet
	s.Fld("Outlier", _sys.Bsc.Bol).Atr = atr.SetGet
	s.FldSlice("sctrs", NewExt("*Sctr"))
	s.FldSlice("sctrRndrSegs", NewExt("*SctrRndrSeg"))
}
func (x *FlePltFltsSctr) InitTypFn() {
	x.New()
	x.Flts()
	x.PrfLos()
}
func (x *FlePltFltsSctr) New() (r *PkgFn) {
	r = x.PkgFna("NewFltsSctr", atr.Lng)
	r.OutPrm(x, "r")
	r.Addf("r = %v{}", x.Typ().Adr(x))
	r.Add("r.PltBse = NewPltBse(r)")
	r.Add("r.Y = NewFltAxisY()")
	r.Add("r.mrgn = Mrgn // glbl mrgn")
	r.Add("return r")
	return r
}
func (x *FlePltFltsSctr) Flts() (r *TypFn) {
	r = x.TypFn(k.Flts)
	r.InPrm(_sys.Ana.Vis.Clr.Clr, "clr")
	// r.InPrm(_sys.Bsc.Unt, "radius")
	r.InPrmVariadic(_sys.Bsc.Flt.arr, "vs")
	r.OutPrm(x)
	// r.Add("if radius == 0 {")
	// r.Add("radius = ShpRadius // glbl ShpRadius")
	// r.Add("}")
	r.Add("for _, v := range vs {")
	r.Add("x.sctrs = append(x.sctrs , &Sctr{")
	r.Add("Y: v,")
	r.Add("clr: clr,")
	r.Add("radius: uint32(ShpRadius), // glbl ShpRadius")
	r.Add("})")
	r.Add("}")
	r.Add("return x")
	return r
}
func (x *FlePltFltsSctr) PrfLos() (r *TypFn) {
	r = x.TypFn(k.PrfLos)
	r.InPrm(_sys.Bsc.Tme.arr, "prfs")
	r.InPrm(_sys.Bsc.Tme.arr, "loss")
	r.InPrmVariadic(_sys.Ana.Hst.Stm, "stms")
	r.OutPrm(x)
	r.Add("for _, stm := range stms {")
	r.Add("x.Flts(PrfClr, stm.At(prfs))")
	r.Add("x.Flts(LosClr, stm.At(loss))")
	r.Add("}")
	r.Add("return x")
	return r
}
