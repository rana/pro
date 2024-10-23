package ana

import (
	"sys"
	"sys/bsc/tme"
	"sys/bsc/unt"
)

type (
	TmeIdx struct {
		Tme tme.Tme
		Idx unt.Unt
	}
	TmeIdxRx  func(pkt TmeIdx) []sys.Act
	TmeIdxRxs map[uint64]TmeIdxRx
	TmeIdxTx  struct {
		Pkt  TmeIdx
		Rx   TmeIdxRx
		ret  []sys.Act
		tier int
	}
	TmeIdxTxr interface {
		Sub(rx TmeIdxRx, id uint32, slot ...uint32)
		Unsub(id uint32, slot ...uint32)
	}
)

func NewTmeIdxTx(pkt TmeIdx, rx TmeIdxRx, tier ...int) (r *TmeIdxTx) {
	r = &TmeIdxTx{}
	r.Pkt = pkt
	r.Rx = rx
	if len(tier) > 0 {
		r.tier = tier[0]
	}
	return r
}
func (x *TmeIdxTx) Act()           { x.ret = x.Rx(x.Pkt) }
func (x *TmeIdxTx) Ret() []sys.Act { return x.ret }
func (x *TmeIdxTx) Tier() int      { return x.tier }
func (x *TmeIdxTx) DecTier()       { x.tier-- }
