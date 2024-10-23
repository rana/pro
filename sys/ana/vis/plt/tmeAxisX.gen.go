package plt

import (
	"sys/bsc/bol"
	"sys/bsc/tme"
)

type (
	TmeAxisX struct {
		Width     uint32
		PxlPerVal float32
		Min       tme.Tme
		Max       tme.Tme
		Rng       tme.Tme
		Inrvls    []tme.Tme
		Lns       []tme.Tme
		Btm       SideTmeAxisX
		vis       bol.Bol
	}
)

func (x *TmeAxisX) Vis(v bol.Bol) *TmeAxisX {
	x.vis = v
	return x
}
