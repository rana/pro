package tpl

import (
	"sys/k"
	"sys/tpl/atr"
)

type (
	PrtPltArngNew struct {
		PrtBse
		New *PkgFn
	}
)

func (x *PrtPltArngNew) InitPrtPkgFn() {
	x.New = x.new()
}
func (x *PrtPltArngNew) new() (r *PkgFn) {
	r = x.f.PkgFnf("%v%v", k.New, x.t.Name)
	r.Atr = atr.Lng
	r.InPrmVariadic(_sys.Ana.Vis.Plt.Plt, "vs")
	r.OutPrm(x.t, "r")
	r.Addf("r = %v{}", x.t.Adr(x.f))
	r.Addf("r.slf = r")
	r.Add("r.Plts = NewPlts()")
	r.Add("r.Plt(vs...)")
	r.Add("return r")
	return r
}
