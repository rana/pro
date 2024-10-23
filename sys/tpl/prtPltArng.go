package tpl

import (
	"sys/k"
	"sys/tpl/atr"
)

type (
	PrtPltArng struct {
		PrtBse
		New *PkgFn
	}
)

func (x *PrtPltArng) InitPrtTyp() {
	x.f.ImplIfc(_sys.Ana.Vis.Plt.Plt.Typ().(*Ifc))
}
func (x *PrtPltArng) InitPrtFld() {
	s := x.f.Typ().(*Struct)
	s.FldTyp(NewExt("PltBse"))
	s.Fld("Plts", _sys.Ana.Vis.Plt.Plt.arr).Atr = atr.Get
}
func (x *PrtPltArng) InitPrtTypFn() {
	x.Plt()
	x.Sho()
	x.Siz()
	x.Scl()
	x.HrzScl()
	x.VrtScl()
}
func (x *PrtPltArng) Plt() (r *TypFn) {
	r = x.f.TypFn(k.Plt)
	r.InPrmVariadic(_sys.Ana.Vis.Plt.Plt, "vs")
	r.OutPrm(x.t)
	r.Add("x.Plts.Push(vs...)")
	r.Add("return x")
	return r
}
func (x *PrtPltArng) Sho() (r *TypFn) {
	r = x.f.TypFn(k.Sho)
	r.OutPrm(_sys.Ana.Vis.Plt.Plt) // for Plt interface
	r.Add("return x.PltBse.Sho()")
	return r
}
func (x *PrtPltArng) Siz() (r *TypFn) {
	r = x.f.TypFn(k.Siz)
	r.InPrm(_sys.Bsc.Unt, "w")
	r.InPrm(_sys.Bsc.Unt, "h")
	r.OutPrm(_sys.Ana.Vis.Plt.Plt)
	r.Add("for _, plt := range *x.Plts {")
	r.Add("plt.Siz(w, h)")
	r.Add("}")
	r.Add("return x")
	return r
}
func (x *PrtPltArng) Scl() (r *TypFn) {
	r = x.f.TypFn(k.Scl)
	r.InPrm(_sys.Bsc.Flt, "v")
	r.OutPrm(_sys.Ana.Vis.Plt.Plt)
	r.Add("for _, plt := range *x.Plts {")
	r.Add("plt.Scl(v)")
	r.Add("}")
	r.Add("return x")
	return r
}
func (x *PrtPltArng) HrzScl() (r *TypFn) {
	r = x.f.TypFn(k.HrzScl)
	r.InPrm(_sys.Bsc.Flt, "v")
	r.OutPrm(_sys.Ana.Vis.Plt.Plt)
	r.Add("for _, plt := range *x.Plts {")
	r.Add("plt.HrzScl(v)")
	r.Add("}")
	r.Add("return x")
	return r
}
func (x *PrtPltArng) VrtScl() (r *TypFn) {
	r = x.f.TypFn(k.VrtScl)
	r.InPrm(_sys.Bsc.Flt, "v")
	r.OutPrm(_sys.Ana.Vis.Plt.Plt)
	r.Add("for _, plt := range *x.Plts {")
	r.Add("plt.VrtScl(v)")
	r.Add("}")
	r.Add("return x")
	return r
}
