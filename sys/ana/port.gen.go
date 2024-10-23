package ana

import (
	"sys/bsc/flt"
)

type (
	Port struct {
		BalFstUsd flt.Flt
		BalLstUsd flt.Flt
		TrdPct    flt.Flt
		Trds      *Trds
	}
	PortScp struct {
		Idx uint32
		Arr []*Port
	}
)
