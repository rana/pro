package plt

import (
	"sys/bsc/tme"
	"sys/bsc/unt"
)

type (
	// TmeAxisX struct {
	// 	Width     uint32 // pixel width of values (not including any margin, border or y-axis width)
	// 	PxlPerVal float32
	// 	Min       tme.Tme
	// 	Max       tme.Tme
	// 	Rng       tme.Tme
	// 	Inrvls    []tme.Tme
	// 	Lns       []tme.Tme
	// 	Btm       SideTmeAxisX
	// }
	SideTmeAxisX struct {
		Height uint32 // total height of side (does not include border)
	}
)

func NewTmeAxisX() (r *TmeAxisX) {
	r = &TmeAxisX{}
	return r
}
func (x *TmeAxisX) Pxl(v tme.Tme) float32 {
	return float32(v-x.Min) * x.PxlPerVal
}
func (x *TmeAxisX) MeasureInrvls(cnt unt.Unt) {
	if x.Min < x.Max {
		x.Inrvls = append(x.Inrvls, x.Min)
		if cnt >= 2 {
			inc := x.Rng / tme.Tme(cnt-1)
			cnt -= 2
			for n := unt.Zero; n < cnt; n++ {
				x.Inrvls = append(x.Inrvls, x.Min+(tme.Tme(n+1)*inc))
			}
		}
		x.Inrvls = append(x.Inrvls, x.Max)
	}
}
