package plt

import (
	"sys/ana/vis/clr"
	"sys/bsc/flt"
	"sys/bsc/flts"
	"sys/bsc/unt"
)

type (
	FltsSctrDist struct {
		PltBse
		XWidth uint32
		Y      *FltAxisY
		itms   []*SctrDistItm
		rndrs  []*SctrDistRndrSeg
	}
	FltsSctrDistScp struct {
		Idx uint32
		Arr []*FltsSctrDist
	}
)

func NewFltsSctrDist() (r *FltsSctrDist) {
	r = &FltsSctrDist{}
	r.PltBse = NewPltBse(r)
	r.XWidth = uint32(Len) // glbl plt len
	r.Y = NewFltAxisY()
	r.mrgn = Mrgn // glbl plt mrgn
	return r
}
func (x *FltsSctrDist) Flts(clr clr.Clr, radius unt.Unt, vs ...*flts.Flts) *FltsSctrDist {
	if radius == 0 {
		radius = ShpRadius // glbl ShpRadius
	}
	for _, v := range vs {
		x.itms = append(x.itms, &SctrDistItm{
			Vals:     v,
			ValsDist: v.CntrDist().Pro(), // USE Pro FOR X-LEN CALC
			clr:      clr,
			radius:   uint32(radius),
			plt:      x,
		})
	}
	return x
}
func (x *FltsSctrDist) Sho() Plt             { return x.PltBse.Sho() }
func (x *FltsSctrDist) Siz(w, h unt.Unt) Plt { return x.PltBse.Siz(w, h) }
func (x *FltsSctrDist) Scl(v flt.Flt) Plt    { return x.PltBse.Scl(v) }
func (x *FltsSctrDist) HrzScl(v flt.Flt) Plt { return x.PltBse.HrzScl(v) }
func (x *FltsSctrDist) VrtScl(v flt.Flt) Plt { return x.PltBse.VrtScl(v) }
