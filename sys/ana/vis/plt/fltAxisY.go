package plt

import (
	"sys/bsc/flt"
	"sys/bsc/unt"
)

type (
	// FltAxisY struct {
	// 	Height    uint32 // pixel height of values (not including any side length, border or margin)
	// 	PxlPerVal float32
	// 	Min       flt.Flt
	// 	Max       flt.Flt
	// 	Rng       flt.Flt
	// 	Inrvls    []flt.Flt
	// 	Lns       []flt.Flt
	// 	Rht       SideFltAxisY
	// }
	SideFltAxisY struct {
		Width uint32 // total width of side (does not include border)
	}
)

func NewFltAxisY() (r *FltAxisY) {
	r = &FltAxisY{}
	r.Min = flt.Max
	r.Max = flt.Min
	r.EqiDst = flt.Max
	return r
}

func (x *FltAxisY) Pxl(v flt.Flt) float32 {
	// USE x.Height - TO FLIP Y; DOMAIN VALUE STARTS AT BTM
	return float32(x.Height) - (float32(v-x.Min) * x.PxlPerVal)
}
func (x *FltAxisY) MeasureInrvls(cnt unt.Unt) {
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
