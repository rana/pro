package plt

import (
	"sys/ana/hst"
	"sys/ana/vis/clr"
	"sys/bsc/bol"
	"sys/bsc/flt"
	"sys/bsc/flts"
	"sys/bsc/str"
	"sys/bsc/tmes"
	"sys/bsc/unt"
)

type (
	FltsSctr struct {
		PltBse
		XWidth       uint32
		Y            *FltAxisY
		Title        str.Str
		Outlier      bol.Bol
		sctrs        []*Sctr
		sctrRndrSegs []*SctrRndrSeg
	}
	FltsSctrScp struct {
		Idx uint32
		Arr []*FltsSctr
	}
)

func NewFltsSctr() (r *FltsSctr) {
	r = &FltsSctr{}
	r.PltBse = NewPltBse(r)
	r.Y = NewFltAxisY()
	r.mrgn = Mrgn // glbl mrgn
	return r
}
func (x *FltsSctr) Flts(clr clr.Clr, vs ...*flts.Flts) *FltsSctr {
	for _, v := range vs {
		x.sctrs = append(x.sctrs, &Sctr{
			Y:      v,
			clr:    clr,
			radius: uint32(ShpRadius), // glbl ShpRadius
		})
	}
	return x
}
func (x *FltsSctr) PrfLos(prfs, loss *tmes.Tmes, stms ...hst.Stm) *FltsSctr {
	for _, stm := range stms {
		x.Flts(PrfClr, stm.At(prfs))
		x.Flts(LosClr, stm.At(loss))
	}
	return x
}
func (x *FltsSctr) Sho() Plt             { return x.PltBse.Sho() }
func (x *FltsSctr) Siz(w, h unt.Unt) Plt { return x.PltBse.Siz(w, h) }
func (x *FltsSctr) Scl(v flt.Flt) Plt    { return x.PltBse.Scl(v) }
func (x *FltsSctr) HrzScl(v flt.Flt) Plt { return x.PltBse.HrzScl(v) }
func (x *FltsSctr) VrtScl(v flt.Flt) Plt { return x.PltBse.VrtScl(v) }
