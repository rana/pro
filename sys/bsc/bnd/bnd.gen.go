package bnd

import (
	"bytes"
	"strings"
	"sys"
	"sys/bsc/bol"
	"sys/bsc/unt"
)

const (
	Size = int(8)
)

type (
	Bnd struct {
		Idx unt.Unt
		Lim unt.Unt
	}
	BndRx  func(pkt Bnd) []sys.Act
	BndRxs map[uint64]BndRx
	BndTx  struct {
		Pkt  Bnd
		Rx   BndRx
		ret  []sys.Act
		tier int
	}
	BndTxr interface {
		Sub(rx BndRx, id uint32, slot ...uint32)
		Unsub(id uint32, slot ...uint32)
	}
	BndScp struct {
		Idx uint32
		Arr []Bnd
	}
)

func NewBndTx(pkt Bnd, rx BndRx, tier ...int) (r *BndTx) {
	r = &BndTx{}
	r.Pkt = pkt
	r.Rx = rx
	if len(tier) > 0 {
		r.tier = tier[0]
	}
	return r
}
func (x Bnd) Cnt() unt.Unt { return x.Lim - x.Idx }
func (x Bnd) Len() unt.Unt { return x.Lim - x.Idx }
func (x Bnd) LstIdx() unt.Unt {
	if x.Lim == 0 {
		return 0
	}
	return x.Lim - 1
}
func (x Bnd) IsValid() bol.Bol { return x.Idx < x.Lim }
func (x Bnd) StrWrt(b *strings.Builder) {
	x.Idx.StrWrt(b)
	b.WriteRune('-')
	x.Lim.StrWrt(b)
}
func (x Bnd) BytWrt(b *bytes.Buffer) {
	x.Idx.BytWrt(b)
	x.Lim.BytWrt(b)
}
func (x *Bnd) BytRed(b []byte) int {
	idx := 0
	idx += x.Idx.BytRed(b[idx:])
	x.Lim.BytRed(b[idx:])
	return Size
}
func (x Bnd) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x Bnd) Bytes() []byte {
	b := &bytes.Buffer{}
	x.BytWrt(b)
	return b.Bytes()
}
func (x *BndTx) Act()           { x.ret = x.Rx(x.Pkt) }
func (x *BndTx) Ret() []sys.Act { return x.ret }
func (x *BndTx) Tier() int      { return x.tier }
func (x *BndTx) DecTier()       { x.tier-- }
