package tme

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
	"strconv"
	"strings"
	"sys"
	"sys/bsc/bol"
	"sys/bsc/unt"
	"time"
	"unsafe"
)

const (
	Zero       = Tme(0)
	One        = Tme(1)
	NegOne     = Tme(-1)
	Min        = Tme(-1 << 31)
	Max        = Tme(1<<31 - 1)
	Second     = Tme(1)
	Minute     = Tme(60)
	Hour       = Tme(60 * 60)
	Day        = Tme(24 * 60 * 60)
	Week       = Tme(7 * 24 * 60 * 60)
	S1         = Tme(1)
	S5         = Tme(5)
	S10        = Tme(10)
	S15        = Tme(15)
	S20        = Tme(20)
	S30        = Tme(30)
	S40        = Tme(40)
	S50        = Tme(50)
	M1         = Tme(1 * 60)
	M5         = Tme(5 * 60)
	M10        = Tme(10 * 60)
	M15        = Tme(15 * 60)
	M20        = Tme(20 * 60)
	M30        = Tme(30 * 60)
	M40        = Tme(40 * 60)
	M50        = Tme(50 * 60)
	H1         = Tme(1 * 60 * 60)
	D1         = Tme(1 * 60 * 60 * 24)
	Size       = int(4)
	DurStrLim  = Tme(1 * 60 * 60 * 24 * 365 * 10)
	Resolution = Tme(1)
)

type (
	Tme    int32
	Cmp    func(a, b Tme) bol.Bol
	TmeRx  func(pkt Tme) []sys.Act
	TmeRxs map[uint64]TmeRx
	TmeTx  struct {
		Pkt  Tme
		Rx   TmeRx
		ret  []sys.Act
		tier int
	}
	TmeTxr interface {
		Sub(rx TmeRx, id uint32, slot ...uint32)
		Unsub(id uint32, slot ...uint32)
	}
	TmeScp struct {
		Idx uint32
		Arr []Tme
	}
)

func Eql(a, b Tme) bol.Bol { return a == b }
func Lss(a, b Tme) bol.Bol { return a < b }
func Gtr(a, b Tme) bol.Bol { return a > b }
func NewTmeTx(pkt Tme, rx TmeRx, tier ...int) (r *TmeTx) {
	r = &TmeTx{}
	r.Pkt = pkt
	r.Rx = rx
	if len(tier) > 0 {
		r.tier = tier[0]
	}
	return r
}
func Now() Tme               { return Time(time.Now()) }
func NewDte(y, n, d int) Tme { return Time(time.Date(y, time.Month(n), d, 0, 0, 0, 0, time.UTC)) }
func NewTme(h, m, s int) Tme { return (Tme(h) * Hour) + (Tme(m) * Minute) + Tme(s) }
func NewDteTme(y, n, d, h, m, s int) Tme {
	return Time(time.Date(y, time.Month(n), d, h, m, s, 0, time.UTC))
}
func Time(v time.Time) Tme         { return Tme(v.UTC().Unix()) }
func Duration(v time.Duration) Tme { return Tme(v.Seconds()) }
func (x Tme) WeekdayCnt(a Tme) (r unt.Unt) {
	min, max := x.MinMax(a)
	for cur := min; cur.Lss(max); cur = cur.Add(Day) {
		weekday := cur.Time().Weekday()
		if weekday != 6 && weekday != 7 {
			r++
		}
	}
	return r
}
func (x Tme) StrWrt(b *strings.Builder) {
	t := x.Time()
	if x >= DurStrLim { // date/time
		y, n, d := t.Date()
		b.WriteString(strconv.FormatInt(int64(y), 10))
		b.WriteRune('y')
		b.WriteString(strconv.FormatInt(int64(n), 10))
		b.WriteRune('n')
		b.WriteString(strconv.FormatInt(int64(d), 10))
		b.WriteRune('d')
		h, m, s := t.Hour(), t.Minute(), t.Second()
		if h != 0 {
			b.WriteString(strconv.FormatInt(int64(h), 10))
			b.WriteRune('h')
		}
		if m != 0 {
			b.WriteString(strconv.FormatInt(int64(m), 10))
			b.WriteRune('m')
		}
		if s != 0 {
			b.WriteString(strconv.FormatInt(int64(s), 10))
			b.WriteRune('s')
		}
	} else { // dur
		x.DurWrt(b)
	}
}
func (x Tme) BytWrt(b *bytes.Buffer) {
	v := make([]byte, Size)
	binary.LittleEndian.PutUint32(v, *(*uint32)(unsafe.Pointer(&x)))
	b.Write(v)
}
func (x *Tme) BytRed(b []byte) int {
	bits := binary.LittleEndian.Uint32(b[:Size])
	*x = *(*Tme)(unsafe.Pointer(&bits))
	return Size
}
func (x Tme) Time() time.Time { return time.Unix(int64(x), 0).UTC() }
func (x Tme) Dte() Tme {
	y, n, d := x.Time().Date()
	return NewDte(y, int(n), d)
}
func (x Tme) ToSunday() Tme {
	t := x.Time()
	y, n, d := t.Date()
	t = time.Date(y, n, d, 0, 0, 0, 0, t.Location())                    // to start of day
	t = t.Add(time.Duration(-t.Weekday()+time.Sunday) * time.Hour * 24) // to day of weeek
	return Time(t)
}
func (x Tme) ToMonday() Tme {
	t := x.Time()
	y, n, d := t.Date()
	t = time.Date(y, n, d, 0, 0, 0, 0, t.Location())                    // to start of day
	t = t.Add(time.Duration(-t.Weekday()+time.Monday) * time.Hour * 24) // to day of weeek
	return Time(t)
}
func (x Tme) ToTuesday() Tme {
	t := x.Time()
	y, n, d := t.Date()
	t = time.Date(y, n, d, 0, 0, 0, 0, t.Location())                     // to start of day
	t = t.Add(time.Duration(-t.Weekday()+time.Tuesday) * time.Hour * 24) // to day of weeek
	return Time(t)
}
func (x Tme) ToWednesday() Tme {
	t := x.Time()
	y, n, d := t.Date()
	t = time.Date(y, n, d, 0, 0, 0, 0, t.Location())                       // to start of day
	t = t.Add(time.Duration(-t.Weekday()+time.Wednesday) * time.Hour * 24) // to day of weeek
	return Time(t)
}
func (x Tme) ToThursday() Tme {
	t := x.Time()
	y, n, d := t.Date()
	t = time.Date(y, n, d, 0, 0, 0, 0, t.Location())                      // to start of day
	t = t.Add(time.Duration(-t.Weekday()+time.Thursday) * time.Hour * 24) // to day of weeek
	return Time(t)
}
func (x Tme) ToFriday() Tme {
	t := x.Time()
	y, n, d := t.Date()
	t = time.Date(y, n, d, 0, 0, 0, 0, t.Location())                    // to start of day
	t = t.Add(time.Duration(-t.Weekday()+time.Friday) * time.Hour * 24) // to day of weeek
	return Time(t)
}
func (x Tme) ToSaturday() Tme {
	t := x.Time()
	y, n, d := t.Date()
	t = time.Date(y, n, d, 0, 0, 0, 0, t.Location())                      // to start of day
	t = t.Add(time.Duration(-t.Weekday()+time.Saturday) * time.Hour * 24) // to day of weeek
	return Time(t)
}
func (x Tme) IsSunday() bol.Bol       { return x.Time().Weekday() == time.Sunday }
func (x Tme) IsMonday() bol.Bol       { return x.Time().Weekday() == time.Monday }
func (x Tme) IsTuesday() bol.Bol      { return x.Time().Weekday() == time.Tuesday }
func (x Tme) IsWednesday() bol.Bol    { return x.Time().Weekday() == time.Wednesday }
func (x Tme) IsThursday() bol.Bol     { return x.Time().Weekday() == time.Thursday }
func (x Tme) IsFriday() bol.Bol       { return x.Time().Weekday() == time.Friday }
func (x Tme) IsSaturday() bol.Bol     { return x.Time().Weekday() == time.Saturday }
func (x Tme) Duration() time.Duration { return time.Duration(x) * time.Second }
func (x Tme) DurWrt(b *strings.Builder) {
	if x == 0 {
		b.WriteString("0s")
		return
	}
	if x < 0 {
		b.WriteRune('-')
		x = -x
	}
	if x >= Week {
		v := x / Week
		b.WriteString(fmt.Sprintf("%vw", int32(v)))
		x -= v * Week
		if x == 0 {
			return
		}
	}
	if x >= Day {
		v := x / Day
		b.WriteString(fmt.Sprintf("%vd", int32(v)))
		x -= v * Day
		if x == 0 {
			return
		}
	}
	if x >= Hour {
		v := x / Hour
		b.WriteString(fmt.Sprintf("%vh", int32(v)))
		x -= v * Hour
		if x == 0 {
			return
		}
	}
	if x >= Minute {
		v := x / Minute
		b.WriteString(fmt.Sprintf("%vm", int32(v)))
		x -= v * Minute
		if x == 0 {
			return
		}
	}
	if x != 0 {
		b.WriteString(fmt.Sprintf("%vs", int32(x)))
	}
}
func (x Tme) DurString() string {
	b := &strings.Builder{}
	x.DurWrt(b)
	return b.String()
}
func (x Tme) Eql(a Tme) bol.Bol { return x == a }
func (x Tme) Neq(a Tme) bol.Bol { return x != a }
func (x Tme) Lss(a Tme) bol.Bol { return x < a }
func (x Tme) Gtr(a Tme) bol.Bol { return x > a }
func (x Tme) Leq(a Tme) bol.Bol { return x <= a }
func (x Tme) Geq(a Tme) bol.Bol { return x >= a }
func (x Tme) Pos() Tme {
	if x < 0 {
		return -x
	} else {
		return x
	}
}
func (x Tme) Neg() Tme {
	if x > 0 {
		return -x
	} else {
		return x
	}
}
func (x Tme) Inv() Tme      { return -x }
func (x Tme) Add(a Tme) Tme { return x + a }
func (x Tme) Sub(a Tme) Tme { return x - a }
func (x Tme) Mul(a Tme) Tme { return x * a }
func (x Tme) Div(a Tme) Tme {
	if a == 0 {
		return 0
	} else {
		return x / a
	}
}
func (x Tme) Rem(a Tme) Tme {
	if a == 0 {
		return 0
	} else {
		return Tme(math.Remainder(float64(x), float64(a)))
	}
}
func (x Tme) Pow(a Tme) Tme { return Tme(math.Pow(float64(x), float64(a))) }
func (x Tme) Sqr() Tme      { return x * x }
func (x Tme) Sqrt() Tme {
	if x <= 0 {
		return 0
	} else {
		return Tme(math.Sqrt(float64(x)))
	}
}
func (x Tme) Min(a Tme) Tme {
	if x < a {
		return x
	} else {
		return a
	}
}
func (x Tme) Max(a Tme) Tme {
	if x > a {
		return x
	} else {
		return a
	}
}
func (x Tme) MinMax(a Tme) (min, max Tme) {
	if x < a {
		return x, a
	}
	return a, x
}
func (x Tme) Mid(a Tme) Tme { return (x - a) / 2 }
func (x Tme) Avg(a Tme) Tme { return (x + a) / 2 }
func (x Tme) AvgGeo(a Tme) (r Tme) {
	r = x * a
	if r == 0 {
		return 0
	} else {
		return r.Sqrt()
	}
}
func (x Tme) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x Tme) Bytes() []byte {
	b := &bytes.Buffer{}
	x.BytWrt(b)
	return b.Bytes()
}
func (x *TmeTx) Act()           { x.ret = x.Rx(x.Pkt) }
func (x *TmeTx) Ret() []sys.Act { return x.ret }
func (x *TmeTx) Tier() int      { return x.tier }
func (x *TmeTx) DecTier()       { x.tier-- }
