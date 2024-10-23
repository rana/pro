package ana

import (
	"sys"
	"sys/bsc/flts"
	"sys/bsc/tme"
)

type (
	TmeFlts struct {
		Tme  tme.Tme
		Flts *flts.Flts
	}
	TmeFltsRx  func(pkt TmeFlts) []sys.Act
	TmeFltsRxs map[uint64]TmeFltsRx
	TmeFltsTx  struct {
		Pkt  TmeFlts
		Rx   TmeFltsRx
		ret  []sys.Act
		tier int
	}
	TmeFltsTxr interface {
		Sub(rx TmeFltsRx, id uint32, slot ...uint32)
		Unsub(id uint32, slot ...uint32)
	}
)

func NewTmeFltsTx(pkt TmeFlts, rx TmeFltsRx, tier ...int) (r *TmeFltsTx) {
	r = &TmeFltsTx{}
	r.Pkt = pkt
	r.Rx = rx
	if len(tier) > 0 {
		r.tier = tier[0]
	}
	return r
}
func (x *TmeFltsTx) Act()           { x.ret = x.Rx(x.Pkt) }
func (x *TmeFltsTx) Ret() []sys.Act { return x.ret }
func (x *TmeFltsTx) Tier() int      { return x.tier }
func (x *TmeFltsTx) DecTier()       { x.tier-- }
