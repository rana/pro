package ana

import (
	"sys"
	"sys/bsc/flt"
	"sys/bsc/tme"
)

type (
	TmeFlt struct {
		Tme tme.Tme
		Flt flt.Flt
	}
	TmeFltRx  func(pkt TmeFlt) []sys.Act
	TmeFltRxs map[uint64]TmeFltRx
	TmeFltTx  struct {
		Pkt  TmeFlt
		Rx   TmeFltRx
		ret  []sys.Act
		tier int
	}
	TmeFltTxr interface {
		Sub(rx TmeFltRx, id uint32, slot ...uint32)
		Unsub(id uint32, slot ...uint32)
	}
)

func NewTmeFltTx(pkt TmeFlt, rx TmeFltRx, tier ...int) (r *TmeFltTx) {
	r = &TmeFltTx{}
	r.Pkt = pkt
	r.Rx = rx
	if len(tier) > 0 {
		r.tier = tier[0]
	}
	return r
}
func (x *TmeFltTx) Act()           { x.ret = x.Rx(x.Pkt) }
func (x *TmeFltTx) Ret() []sys.Act { return x.ret }
func (x *TmeFltTx) Tier() int      { return x.tier }
func (x *TmeFltTx) DecTier()       { x.tier-- }
