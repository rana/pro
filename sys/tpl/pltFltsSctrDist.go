package tpl

import (
	"sys/k"
	"sys/tpl/atr"
)

type (
	FlePltFltsSctrDist struct {
		FleBse
		PrtPlt
		PrtStructFldSet
	}
)

func (x *DirPlt) NewFltsSctrDist() (r *FlePltFltsSctrDist) {
	r = &FlePltFltsSctrDist{}
	x.FltsSctrDist = r
	r.Name = k.FltsSctrDist
	r.Pkg = x.Pkg
	r.StructPtr(r.Name, atr.TypUiStruct)
	r.AddFle(r)
	return r
}
func (x *FlePltFltsSctrDist) InitTyp(bse *TypBse) {
	x.ImplIfc(_sys.Ana.Vis.Plt.Plt.Typ().(*Ifc))
}
func (x *FlePltFltsSctrDist) InitFld(s *Struct) {
	s.FldTyp(NewExt("PltBse"))
	s.Fld("XWidth", Uint32)
	s.Fld("Y", NewExt("*FltAxisY"))
	s.FldSlice("itms", NewExt("*SctrDistItm"))
	s.FldSlice("rndrs", NewExt("*SctrDistRndrSeg"))
}
func (x *FlePltFltsSctrDist) InitTypFn() {
	x.New()
	x.Flts()
}
func (x *FlePltFltsSctrDist) New() (r *PkgFn) {
	r = x.PkgFna("NewFltsSctrDist", atr.Lng)
	r.OutPrm(x, "r")
	r.Addf("r = %v{}", x.Typ().Adr(x))
	r.Add("r.PltBse = NewPltBse(r)")
	r.Add("r.XWidth = uint32(Len) // glbl plt len")
	r.Add("r.Y = NewFltAxisY()")
	r.Add("r.mrgn = Mrgn // glbl plt mrgn")
	r.Add("return r")
	return r
}
func (x *FlePltFltsSctrDist) Flts() (r *TypFn) {
	r = x.TypFn(k.Flts)
	r.InPrm(_sys.Ana.Vis.Clr.Clr, "clr")
	r.InPrm(_sys.Bsc.Unt, "radius")
	r.InPrmVariadic(_sys.Bsc.Flt.arr, "vs")
	r.OutPrm(x)
	r.Add("if radius == 0 {")
	r.Add("radius = ShpRadius // glbl ShpRadius")
	r.Add("}")
	r.Add("for _, v := range vs {")
	r.Add("x.itms = append(x.itms , &SctrDistItm{")
	r.Add("Vals: v,")
	r.Add("ValsDist: v.CntrDist().Pro(), // USE Pro FOR X-LEN CALC")
	r.Add("clr: clr,")
	r.Add("radius: uint32(radius),")
	r.Add("plt: x,")
	r.Add("})")
	r.Add("}")
	r.Add("return x")
	return r
}
