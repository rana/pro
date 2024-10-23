package ana

import (
	"sys/bsc/flts"
	"sys/bsc/tme"
)

type (
	Tic struct {
		Tme  tme.Tme
		Bids *flts.Flts
		Asks *flts.Flts
	}
)
