package plt

import (
	"sys/bsc/flt"
	"sys/bsc/unt"
)

type (
	Vrt struct {
		PltBse
		Plts *Plts
	}
	VrtScp struct {
		Idx uint32
		Arr []*Vrt
	}
)

func NewVrt(vs ...Plt) (r *Vrt) {
	r = &Vrt{}
	r.slf = r
	r.Plts = NewPlts()
	r.Plt(vs...)
	return r
}
func (x *Vrt) Plt(vs ...Plt) *Vrt {
	x.Plts.Push(vs...)
	return x
}
func (x *Vrt) Sho() Plt { return x.PltBse.Sho() }
func (x *Vrt) Siz(w, h unt.Unt) Plt {
	for _, plt := range *x.Plts {
		plt.Siz(w, h)
	}
	return x
}
func (x *Vrt) Scl(v flt.Flt) Plt {
	for _, plt := range *x.Plts {
		plt.Scl(v)
	}
	return x
}
func (x *Vrt) HrzScl(v flt.Flt) Plt {
	for _, plt := range *x.Plts {
		plt.HrzScl(v)
	}
	return x
}
func (x *Vrt) VrtScl(v flt.Flt) Plt {
	for _, plt := range *x.Plts {
		plt.VrtScl(v)
	}
	return x
}
