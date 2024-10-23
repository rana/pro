package ana

import (
	"bytes"
	"sys/bsc/flts"
	"sys/bsc/tme"
	"sys/bsc/tmes"
	"sys/bsc/unt"
	"sys/bsc/unts"
)

type (
	Stm struct {
		Tmes    *tmes.Tmes
		Bids    *flts.Flts
		Asks    *flts.Flts
		BidLims *unts.Unts
		AskLims *unts.Unts
		RxIdx   unt.Unt
		RxTme   tme.Tme
	}
)

func (x *Stm) Bytes() []byte {
	b := &bytes.Buffer{}
	x.BytWrt(b)
	return b.Bytes()
}
func (x *Stm) BytWrt(b *bytes.Buffer) {
	x.Tmes.BytWrt(b)
	x.Bids.BytWrt(b)
	x.Asks.BytWrt(b)
	x.BidLims.BytWrt(b)
	x.AskLims.BytWrt(b)
}
func (x *Stm) BytRed(b []byte) (idx int) {
	x.Tmes = tmes.New()
	x.Bids = flts.New()
	x.Asks = flts.New()
	x.BidLims = unts.New()
	x.AskLims = unts.New()
	idx += x.Tmes.BytRed(b)
	idx += x.Bids.BytRed(b[idx:])
	idx += x.Asks.BytRed(b[idx:])
	idx += x.BidLims.BytRed(b[idx:])
	idx += x.AskLims.BytRed(b[idx:])
	return idx
}
