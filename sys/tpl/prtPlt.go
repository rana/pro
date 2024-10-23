package tpl

import (
	"sys/k"
)

type (
	PrtPlt struct {
		PrtBse
	}
)

func (x *PrtPlt) InitPrtTypFn() {
	x.Sho()
	x.Siz()
	x.Scl()
	x.HrzScl()
	x.VrtScl()
	// x.ToPlts()
}
func (x *PrtPlt) Sho() (r *TypFn) {
	r = x.f.TypFn(k.Sho)
	r.OutPrm(_sys.Ana.Vis.Plt.Plt) // for Plt interface
	r.Add("return x.PltBse.Sho()")
	return r
}
func (x *PrtPlt) Siz() (r *TypFn) {
	r = x.f.TypFn(k.Siz)
	r.InPrm(_sys.Bsc.Unt, "w")
	r.InPrm(_sys.Bsc.Unt, "h")
	r.OutPrm(_sys.Ana.Vis.Plt.Plt)
	r.Add("return x.PltBse.Siz(w, h)")
	return r
}
func (x *PrtPlt) Scl() (r *TypFn) {
	r = x.f.TypFn(k.Scl)
	r.InPrm(_sys.Bsc.Flt, "v")
	r.OutPrm(_sys.Ana.Vis.Plt.Plt)
	r.Add("return x.PltBse.Scl(v)")
	return r
}
func (x *PrtPlt) HrzScl() (r *TypFn) {
	r = x.f.TypFn(k.HrzScl)
	r.InPrm(_sys.Bsc.Flt, "v")
	r.OutPrm(_sys.Ana.Vis.Plt.Plt)
	r.Add("return x.PltBse.HrzScl(v)")
	return r
}
func (x *PrtPlt) VrtScl() (r *TypFn) {
	r = x.f.TypFn(k.VrtScl)
	r.InPrm(_sys.Bsc.Flt, "v")
	r.OutPrm(_sys.Ana.Vis.Plt.Plt)
	r.Add("return x.PltBse.VrtScl(v)")
	return r
}

// func (x *PrtPlt) ToPlts() (r *PkgFn) {
// 	r = x.f.PkgFnf("%vsToPlts", x.t.Title())
// 	r.SkpXpr = true
// 	r.InPrmVariadic(x.t, "vs")
// 	r.OutPrmSlice(_sys.Ana.Vis.Plt.Plt, "r")
// 	r.Addf("r = make([]%v, len(vs))", _sys.Ana.Vis.Plt.Plt.Ref(x.f))
// 	r.Add("for n, v := range vs {")
// 	r.Add("r[n] = v")
// 	r.Add("}")
// 	r.Add("return r")
// 	return r
// }
