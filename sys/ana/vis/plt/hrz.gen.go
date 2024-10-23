package plt

import (
	"sys/bsc/flt"
	"sys/bsc/unt"
)

type (
	Hrz struct {
		PltBse
		Plts *Plts
	}
	HrzScp struct {
		Idx uint32
		Arr []*Hrz
	}
)

func NewHrz(vs ...Plt) (r *Hrz) {
	r = &Hrz{}
	r.slf = r
	r.Plts = NewPlts()
	r.Plt(vs...)
	return r
}
func (x *Hrz) Plt(vs ...Plt) *Hrz {
	x.Plts.Push(vs...)
	return x
}
func (x *Hrz) Sho() Plt { return x.PltBse.Sho() }
func (x *Hrz) Siz(w, h unt.Unt) Plt {
	for _, plt := range *x.Plts {
		plt.Siz(w, h)
	}
	return x
}
func (x *Hrz) Scl(v flt.Flt) Plt {
	for _, plt := range *x.Plts {
		plt.Scl(v)
	}
	return x
}
func (x *Hrz) HrzScl(v flt.Flt) Plt {
	for _, plt := range *x.Plts {
		plt.HrzScl(v)
	}
	return x
}
func (x *Hrz) VrtScl(v flt.Flt) Plt {
	for _, plt := range *x.Plts {
		plt.VrtScl(v)
	}
	return x
}
