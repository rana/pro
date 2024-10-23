package plt

import (
	"sys/bsc/flt"
	"sys/bsc/unt"
)

type (
	Dpth struct {
		PltBse
		Plts *Plts
	}
	DpthScp struct {
		Idx uint32
		Arr []*Dpth
	}
)

func NewDpth(vs ...Plt) (r *Dpth) {
	r = &Dpth{}
	r.slf = r
	r.Plts = NewPlts()
	r.Plt(vs...)
	return r
}
func (x *Dpth) Plt(vs ...Plt) *Dpth {
	x.Plts.Push(vs...)
	return x
}
func (x *Dpth) Sho() Plt { return x.PltBse.Sho() }
func (x *Dpth) Siz(w, h unt.Unt) Plt {
	for _, plt := range *x.Plts {
		plt.Siz(w, h)
	}
	return x
}
func (x *Dpth) Scl(v flt.Flt) Plt {
	for _, plt := range *x.Plts {
		plt.Scl(v)
	}
	return x
}
func (x *Dpth) HrzScl(v flt.Flt) Plt {
	for _, plt := range *x.Plts {
		plt.HrzScl(v)
	}
	return x
}
func (x *Dpth) VrtScl(v flt.Flt) Plt {
	for _, plt := range *x.Plts {
		plt.VrtScl(v)
	}
	return x
}
