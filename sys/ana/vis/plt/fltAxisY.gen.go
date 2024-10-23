package plt

import (
	"sys/bsc/bol"
	"sys/bsc/flt"
)

type (
	FltAxisY struct {
		Height    uint32
		PxlPerVal float32
		Min       flt.Flt
		Max       flt.Flt
		Rng       flt.Flt
		EqiDst    flt.Flt
		Inrvls    []flt.Flt
		Lns       []flt.Flt
		Rht       SideFltAxisY
		vis       bol.Bol
	}
)

func (x *FltAxisY) Vis(v bol.Bol) *FltAxisY {
	x.vis = v
	return x
}
