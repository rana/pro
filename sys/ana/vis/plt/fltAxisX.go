package plt

import (
	"sys/bsc/flt"
	"sys/bsc/unt"
)

type (
	FltAxisX struct {
		Width     uint32 // pixel width of values (not including any side length, border or margin)
		PxlPerVal float32
		Min       flt.Flt
		Max       flt.Flt
		Rng       flt.Flt
		Inrvls    []flt.Flt
		Lns       []flt.Flt
		Btm       SideFltAxisX
	}
	SideFltAxisX struct {
		Height uint32 // total height of side (does not include border)
	}
)

func NewFltAxisX() (r *FltAxisX) {
	r = &FltAxisX{}
	return r
}

func (x *FltAxisX) Pxl(v flt.Flt) float32 {
	return float32(v-x.Min) * x.PxlPerVal
}
func (x *FltAxisX) MeasureInrvls(cnt unt.Unt) {
	if x.Min < x.Max {
		x.Inrvls = append(x.Inrvls, x.Min)
		if cnt >= 2 {
			inc := x.Rng / flt.Flt(cnt-1)
			cnt -= 2
			for n := unt.Zero; n < cnt; n++ {
				x.Inrvls = append(x.Inrvls, x.Min+(flt.Flt(n+1)*inc))
			}
		}
		x.Inrvls = append(x.Inrvls, x.Max)
	}
}
