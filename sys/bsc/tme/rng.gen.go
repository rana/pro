package tme

import (
	"bytes"
	"strings"
	"sys/bsc/bol"
)

type (
	Rng struct {
		Min Tme
		Max Tme
	}
	RngScp struct {
		Idx uint32
		Arr []Rng
	}
)

func NewRng(min, max Tme) (r Rng) {
	r.Min = min
	r.Max = max
	return r
}
func NewRngArnd(cntr, radius Tme) (r Rng) {
	r.Min = cntr - radius
	r.Max = cntr + radius
	return r
}
func NewRngFul() (r Rng) {
	r.Min = Min
	r.Max = Max
	return r
}
func (x Rng) Len() Tme         { return x.Max - x.Min }
func (x Rng) IsValid() bol.Bol { return x.Min < x.Max }
func (x Rng) Ensure() Rng {
	if x.Min > x.Max {
		x.Min, x.Max = x.Max, x.Min // swp
	}
	return x
}
func (x Rng) MinSub(v Tme) Rng {
	x.Min -= v
	return x
}
func (x Rng) MaxAdd(v Tme) Rng {
	x.Max += v
	return x
}
func (x Rng) Mrg(v Rng) Rng {
	x.Min = x.Min.Min(v.Min)
	x.Max = x.Max.Max(v.Max)
	return x
}
func (x Rng) StrWrt(b *strings.Builder) {
	x.Min.StrWrt(b)
	b.WriteRune('-')
	x.Max.StrWrt(b)
}
func (x Rng) BytWrt(b *bytes.Buffer) {
	x.Min.BytWrt(b)
	x.Max.BytWrt(b)
}
func (x *Rng) BytRed(b []byte) int {
	idx := 0
	idx += x.Min.BytRed(b[idx:])
	x.Max.BytRed(b[idx:])
	return Size
}
func (x Rng) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x Rng) Bytes() []byte {
	b := &bytes.Buffer{}
	x.BytWrt(b)
	return b.Bytes()
}
