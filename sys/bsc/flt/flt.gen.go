package flt

import (
	"bytes"
	"encoding/binary"
	"math"
	"strconv"
	"strings"
	"sys/bsc/bol"
	"sys/bsc/unt"
	"sys/err"
	"unsafe"
)

const (
	Zero   = Flt(0.0)
	One    = Flt(1.0)
	NegOne = Flt(-1.0)
	Hndrd  = Flt(100.0)
	Min    = Flt(-3.40282346638528859811704183484516925440e+38)
	Max    = Flt(3.40282346638528859811704183484516925440e+38)
	Tiny   = Flt(1.401298464324817070923729583289916131280e-45)
	Size   = int(4)
)

type (
	Flt    float32
	Cmp    func(a, b Flt) bol.Bol
	FltScp struct {
		Idx uint32
		Arr []Flt
	}
)

func Eql(a, b Flt) bol.Bol      { return a == b }
func Lss(a, b Flt) bol.Bol      { return a < b }
func Gtr(a, b Flt) bol.Bol      { return a > b }
func (x Flt) Eql(a Flt) bol.Bol { return x == a || x.IsNaN() && a.IsNaN() }
func (x Flt) Neq(a Flt) bol.Bol { return x != a }
func (x Flt) Trnc(precision unt.Unt) Flt {
	if x != x { // IsNaN
		return x
	}
	s := x.String()
	idx := strings.Index(s, ".")
	if len(s)-idx-1 <= int(precision) {
		return x
	}
	v, er := strconv.ParseFloat(s[:idx+1+int(precision)], 32)
	if er != nil {
		err.Panicf("Flt: failed to parse (txt:%q err:%q s:%v)", s[:idx+1+int(precision)], er, s)
	}
	return Flt(v)
}
func (x Flt) IsNaN() bol.Bol {
	// From /usr/local/go/src/math/bits.go
	return x != x
}
func (x Flt) IsInfPos() bol.Bol {
	// From /usr/local/go/src/math/bits.go
	return x > Max
}
func (x Flt) IsInfNeg() bol.Bol {
	// From /usr/local/go/src/math/bits.go
	return x < Min
}
func (x Flt) IsValid() bol.Bol {
	// From /usr/local/go/src/math/bits.go
	return x == x && x >= Min && x <= Max
}
func (x Flt) Pct(v Flt) (r Flt) {
	if v-x == 0 {
		return 0
	}
	if x == 0 {
		return 1
	}
	if x < 0 {
		return ((v - x) / -x)
	}
	return ((v - x) / x)
}
func (x Flt) StrWrt(b *strings.Builder) {
	s := strconv.FormatFloat(float64(x), byte('f'), -1, 32)
	b.WriteString(s)
	if x.IsValid() && strings.LastIndex(s, ".") < 0 {
		b.WriteString(".0")
	}
}
func (x Flt) BytWrt(b *bytes.Buffer) {
	v := make([]byte, Size)
	binary.LittleEndian.PutUint32(v, *(*uint32)(unsafe.Pointer(&x)))
	b.Write(v)
}
func (x *Flt) BytRed(b []byte) int {
	bits := binary.LittleEndian.Uint32(b[:4])
	*x = *(*Flt)(unsafe.Pointer(&bits))
	return Size
}
func (x Flt) Lss(a Flt) bol.Bol { return x < a }
func (x Flt) Gtr(a Flt) bol.Bol { return x > a }
func (x Flt) Leq(a Flt) bol.Bol { return x <= a }
func (x Flt) Geq(a Flt) bol.Bol { return x >= a }
func (x Flt) Pos() Flt {
	if x < 0 {
		return -x
	} else {
		return x
	}
}
func (x Flt) Neg() Flt {
	if x > 0 {
		return -x
	} else {
		return x
	}
}
func (x Flt) Inv() Flt      { return -x }
func (x Flt) Add(a Flt) Flt { return x + a }
func (x Flt) Sub(a Flt) Flt { return x - a }
func (x Flt) Mul(a Flt) Flt { return x * a }
func (x Flt) Div(a Flt) Flt {
	if a == 0 {
		return 0
	} else {
		return x / a
	}
}
func (x Flt) Rem(a Flt) Flt {
	if a == 0 {
		return 0
	} else {
		return Flt(math.Remainder(float64(x), float64(a)))
	}
}
func (x Flt) Pow(a Flt) Flt { return Flt(math.Pow(float64(x), float64(a))) }
func (x Flt) Sqr() Flt      { return x * x }
func (x Flt) Sqrt() Flt {
	if x <= 0 {
		return 0
	} else {
		return Flt(math.Sqrt(float64(x)))
	}
}
func (x Flt) Min(a Flt) Flt {
	if x < a {
		return x
	} else {
		return a
	}
}
func (x Flt) Max(a Flt) Flt {
	if x > a {
		return x
	} else {
		return a
	}
}
func (x Flt) MinMax(a Flt) (min, max Flt) {
	if x < a {
		return x, a
	}
	return a, x
}
func (x Flt) Mid(a Flt) Flt { return (x - a) / 2 }
func (x Flt) Avg(a Flt) Flt { return (x + a) / 2 }
func (x Flt) AvgGeo(a Flt) (r Flt) {
	r = x * a
	if r == 0 {
		return 0
	} else {
		return r.Sqrt()
	}
}
func (x Flt) SelEql(a Flt) Flt {
	if x == a {
		return x
	} else {
		return 0
	}
}
func (x Flt) SelNeq(a Flt) Flt {
	if x != a {
		return x
	} else {
		return 0
	}
}
func (x Flt) SelLss(a Flt) Flt {
	if x < a {
		return x
	} else {
		return 0
	}
}
func (x Flt) SelGtr(a Flt) Flt {
	if x > a {
		return x
	} else {
		return 0
	}
}
func (x Flt) SelLeq(a Flt) Flt {
	if x <= a {
		return x
	} else {
		return 0
	}
}
func (x Flt) SelGeq(a Flt) Flt {
	if x >= a {
		return x
	} else {
		return 0
	}
}
func (x Flt) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x Flt) Bytes() []byte {
	b := &bytes.Buffer{}
	x.BytWrt(b)
	return b.Bytes()
}
