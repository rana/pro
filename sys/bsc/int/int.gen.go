package int

import (
	"bytes"
	"encoding/binary"
	"math"
	"strings"
	"sys/bsc/bol"
	"unsafe"
)

const (
	Zero   = Int(0)
	One    = Int(1)
	NegOne = Int(-1)
	Min    = Int(-1 << 31)
	Max    = Int(1<<31 - 1)
	Size   = int(4)
)

type (
	Int    int32
	Cmp    func(a, b Int) bol.Bol
	IntScp struct {
		Idx uint32
		Arr []Int
	}
)

func Eql(a, b Int) bol.Bol { return a == b }
func Lss(a, b Int) bol.Bol { return a < b }
func Gtr(a, b Int) bol.Bol { return a > b }
func (x Int) StrWrt(b *strings.Builder) {
	if x == 0 {
		b.WriteString("+0")
	} else {
		var rs []rune
		if x < 0 {
			b.WriteRune('-')
			for x != 0 { // separate branch to support math.MinInt32
				rs = append(rs, '0'+rune(-(x%10)))
				x /= 10
			}
		} else {
			b.WriteRune('+')
			for x != 0 {
				rs = append(rs, '0'+rune(x%10))
				x /= 10
			}
		}
		for i, j := 0, len(rs)-1; i < j; i, j = i+1, j-1 { // reverse
			rs[i], rs[j] = rs[j], rs[i]
		}
		for n := 0; n < len(rs); n++ {
			b.WriteRune(rs[n])
		}
	}
}
func (x Int) BytWrt(b *bytes.Buffer) {
	v := make([]byte, Size)
	binary.LittleEndian.PutUint32(v, *(*uint32)(unsafe.Pointer(&x)))
	b.Write(v)
}
func (x *Int) BytRed(b []byte) int {
	bits := binary.LittleEndian.Uint32(b[:Size])
	*x = *(*Int)(unsafe.Pointer(&bits))
	return Size
}
func (x Int) Eql(a Int) bol.Bol { return x == a }
func (x Int) Neq(a Int) bol.Bol { return x != a }
func (x Int) Lss(a Int) bol.Bol { return x < a }
func (x Int) Gtr(a Int) bol.Bol { return x > a }
func (x Int) Leq(a Int) bol.Bol { return x <= a }
func (x Int) Geq(a Int) bol.Bol { return x >= a }
func (x Int) Pos() Int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}
func (x Int) Neg() Int {
	if x > 0 {
		return -x
	} else {
		return x
	}
}
func (x Int) Inv() Int      { return -x }
func (x Int) Add(a Int) Int { return x + a }
func (x Int) Sub(a Int) Int { return x - a }
func (x Int) Mul(a Int) Int { return x * a }
func (x Int) Div(a Int) Int {
	if a == 0 {
		return 0
	} else {
		return x / a
	}
}
func (x Int) Rem(a Int) Int {
	if a == 0 {
		return 0
	} else {
		return Int(math.Remainder(float64(x), float64(a)))
	}
}
func (x Int) Pow(a Int) Int { return Int(math.Pow(float64(x), float64(a))) }
func (x Int) Sqr() Int      { return x * x }
func (x Int) Sqrt() Int {
	if x <= 0 {
		return 0
	} else {
		return Int(math.Sqrt(float64(x)))
	}
}
func (x Int) Min(a Int) Int {
	if x < a {
		return x
	} else {
		return a
	}
}
func (x Int) Max(a Int) Int {
	if x > a {
		return x
	} else {
		return a
	}
}
func (x Int) MinMax(a Int) (min, max Int) {
	if x < a {
		return x, a
	}
	return a, x
}
func (x Int) Mid(a Int) Int { return (x - a) / 2 }
func (x Int) Avg(a Int) Int { return (x + a) / 2 }
func (x Int) AvgGeo(a Int) (r Int) {
	r = x * a
	if r == 0 {
		return 0
	} else {
		return r.Sqrt()
	}
}
func (x Int) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x Int) Bytes() []byte {
	b := &bytes.Buffer{}
	x.BytWrt(b)
	return b.Bytes()
}
