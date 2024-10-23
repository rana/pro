package unt

import (
	"bytes"
	"encoding/binary"
	"math"
	"strings"
	"sys/bsc/bol"
)

const (
	Zero      = Unt(0)
	One       = Unt(1)
	Min       = Unt(0)
	Max       = Unt(1<<32 - 1)
	Size      = int(4)
	MinSegLen = Unt(64)
)

type (
	Unt    uint32
	Cmp    func(a, b Unt) bol.Bol
	UntScp struct {
		Idx uint32
		Arr []Unt
	}
)

func Eql(a, b Unt) bol.Bol { return a == b }
func Lss(a, b Unt) bol.Bol { return a < b }
func Gtr(a, b Unt) bol.Bol { return a > b }
func (x Unt) StrWrt(b *strings.Builder) {
	if x == 0 { // TODO: OPTIMIZE
		b.WriteRune('0')
	} else {
		var rs []rune
		for x != 0 {
			rs = append(rs, '0'+rune(x%10))
			x /= 10
		}
		for i, j := 0, len(rs)-1; i < j; i, j = i+1, j-1 { // reverse
			rs[i], rs[j] = rs[j], rs[i]
		}
		for n := 0; n < len(rs); n++ {
			b.WriteRune(rs[n])
		}
	}
}
func (x Unt) BytWrt(b *bytes.Buffer) {
	v := make([]byte, Size)
	binary.LittleEndian.PutUint32(v, uint32(x))
	b.Write(v)
}
func (x *Unt) BytRed(b []byte) int {
	*x = Unt(binary.LittleEndian.Uint32(b[:Size]))
	return Size
}
func (x Unt) Eql(a Unt) bol.Bol { return x == a }
func (x Unt) Neq(a Unt) bol.Bol { return x != a }
func (x Unt) Lss(a Unt) bol.Bol { return x < a }
func (x Unt) Gtr(a Unt) bol.Bol { return x > a }
func (x Unt) Leq(a Unt) bol.Bol { return x <= a }
func (x Unt) Geq(a Unt) bol.Bol { return x >= a }
func (x Unt) Add(a Unt) Unt     { return x + a }
func (x Unt) Sub(a Unt) Unt     { return x - a }
func (x Unt) Mul(a Unt) Unt     { return x * a }
func (x Unt) Div(a Unt) Unt {
	if a == 0 {
		return 0
	} else {
		return x / a
	}
}
func (x Unt) Rem(a Unt) Unt {
	if a == 0 {
		return 0
	} else {
		return Unt(math.Remainder(float64(x), float64(a)))
	}
}
func (x Unt) Pow(a Unt) Unt { return Unt(math.Pow(float64(x), float64(a))) }
func (x Unt) Sqr() Unt      { return x * x }
func (x Unt) Sqrt() Unt {
	if x <= 0 {
		return 0
	} else {
		return Unt(math.Sqrt(float64(x)))
	}
}
func (x Unt) Min(a Unt) Unt {
	if x < a {
		return x
	} else {
		return a
	}
}
func (x Unt) Max(a Unt) Unt {
	if x > a {
		return x
	} else {
		return a
	}
}
func (x Unt) MinMax(a Unt) (min, max Unt) {
	if x < a {
		return x, a
	}
	return a, x
}
func (x Unt) Mid(a Unt) Unt { return (x - a) / 2 }
func (x Unt) Avg(a Unt) Unt { return (x + a) / 2 }
func (x Unt) AvgGeo(a Unt) (r Unt) {
	r = x * a
	if r == 0 {
		return 0
	} else {
		return r.Sqrt()
	}
}
func (x Unt) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x Unt) Bytes() []byte {
	b := &bytes.Buffer{}
	x.BytWrt(b)
	return b.Bytes()
}
