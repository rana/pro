package tmes

import (
	"bytes"
	"encoding/binary"
	"math/rand"
	"sort"
	"strings"
	"sys/bsc/bnd"
	"sys/bsc/bol"
	"sys/bsc/tme"
	"sys/bsc/unt"
	"time"
)

type (
	Tmes    []tme.Tme
	TmesScp struct {
		Idx uint32
		Arr []*Tmes
	}
	MinMaxSeg struct {
		bnd.Bnd
		Min  tme.Tme
		Max  tme.Tme
		Vals *Tmes
	}
)

func New(vs ...tme.Tme) *Tmes {
	r := Tmes(vs)
	return &r
}
func Make(cap unt.Unt) *Tmes {
	r := make(Tmes, cap)
	return &r
}
func MakeEmp(cap unt.Unt) *Tmes {
	r := make(Tmes, 0, cap)
	return &r
}
func AddsLss(strt, lim, by tme.Tme) (r *Tmes) {
	r = New()
	for n := strt; n < lim; n += by {
		r.Push(n)
	}
	return r
}
func AddsLeq(strt, lim, by tme.Tme) (r *Tmes) {
	r = New()
	for n := strt; n <= lim; n += by {
		r.Push(n)
	}
	return r
}
func SubsGtr(strt, lim, by tme.Tme) (r *Tmes) {
	r = New()
	for n := strt; n > lim; n -= by {
		r.Push(n)
	}
	return r
}
func SubsGeq(strt, lim, by tme.Tme) (r *Tmes) {
	r = New()
	for n := strt; n >= lim; n -= by {
		r.Push(n)
	}
	return r
}
func MulsLss(strt, lim, by tme.Tme) (r *Tmes) {
	r = New()
	for n := strt; n < lim; n *= by {
		r.Push(n)
	}
	return r
}
func MulsLeq(strt, lim, by tme.Tme) (r *Tmes) {
	r = New()
	for n := strt; n <= lim; n *= by {
		r.Push(n)
	}
	return r
}
func DivsGtr(strt, lim, by tme.Tme) (r *Tmes) {
	r = New()
	for n := strt; n > lim; n /= by {
		r.Push(n)
	}
	return r
}
func DivsGeq(strt, lim, by tme.Tme) (r *Tmes) {
	r = New()
	for n := strt; n >= lim; n /= by {
		r.Push(n)
	}
	return r
}
func FibsLeq(lim tme.Tme) (r *Tmes) {
	r = New(1, 2)
	for (*r)[len(*r)-1] < lim {
		*r = append(*r, (*r)[len(*r)-2]+(*r)[len(*r)-1])
	}
	return r
}
func (x *Tmes) Times() (r []time.Time) {
	r = make([]time.Time, len(*x))
	for n, t := range *x {
		r[n] = t.Time()
	}
	return r
}
func (x *Tmes) Bnd(rng tme.Rng) (r bnd.Bnd) {
	if len(*x) == 0 {
		return r
	}
	rng = rng.Ensure()
	r.Idx = x.SrchIdx(rng.Min, true)
	r.Lim = x.SrchIdx(rng.Max, true)
	if r.Idx >= x.Cnt() {
		return bnd.Bnd{}
	}
	if r.Lim > x.Cnt() {
		r.Lim = x.Cnt()
	}
	return r
}
func (x *Tmes) WeekdayCnt() (r unt.Unt) {
	if len(*x) == 0 {
		return 0
	}
	return x.Fst().WeekdayCnt((*x)[len(*x)-1])
}
func (x *Tmes) Ok() bol.Bol { return len(*x) != 0 }
func (x *Tmes) Cnt() unt.Unt {
	if x == nil {
		return 0
	}
	return unt.Unt(len(*x))
}
func (x *Tmes) Cpy() *Tmes {
	r := make(Tmes, len(*x))
	copy(r, *x)
	return &r
}
func (x *Tmes) Clr() *Tmes {
	*x = (*x)[:0]
	return x
}
func (x *Tmes) Rand() *Tmes {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	perm := r.Perm(len(*x))
	for i, randIdx := range perm {
		(*x)[i], (*x)[randIdx] = (*x)[randIdx], (*x)[i]
	}
	return x
}
func (x *Tmes) Mrg(a ...*Tmes) *Tmes {
	for _, v := range a {
		*x = append(*x, *v...)
	}
	return x
}
func (x *Tmes) Push(a ...tme.Tme) *Tmes {
	*x = append(*x, a...)
	return x
}
func (x *Tmes) Pop() (r tme.Tme) {
	if len(*x) == 0 {
		return r
	}
	r = (*x)[len(*x)-1]
	*x = (*x)[:len(*x)-1]
	return r
}
func (x *Tmes) Que(vs ...tme.Tme) *Tmes {
	*x = append(*x, vs...)
	return x
}
func (x *Tmes) Dque() (r tme.Tme) {
	r = (*x)[0]
	if len(*x) == 1 {
		*x = (*x)[:0]
	} else {
		copy(*x, (*x)[1:])
		*x = (*x)[:len(*x)-1]
	}
	return r
}
func (x *Tmes) Ins(idx unt.Unt, elm tme.Tme) *Tmes {
	*x = append((*x)[:idx], append(Tmes{elm}, (*x)[idx:]...)...)
	return x
}
func (x *Tmes) Upd(idx unt.Unt, elm tme.Tme) *Tmes {
	(*x)[idx] = elm
	return x
}
func (x *Tmes) Del(idx unt.Unt) (r tme.Tme) {
	r = (*x)[idx]
	if idx == 0 && len(*x) == 1 {
		*x = (*x)[:0]
	} else if idx == unt.Unt(len(*x)-1) {
		*x = (*x)[:idx]
	} else {
		*x = append((*x)[:idx], (*x)[idx+1:]...)
	}
	return r
}
func (x *Tmes) At(idx unt.Unt) tme.Tme { return (*x)[idx] }
func (x *Tmes) In(idx, lim unt.Unt) *Tmes {
	r := (*x)[idx:lim]
	return &r
}
func (x *Tmes) InBnd(b bnd.Bnd) *Tmes {
	r := (*x)[b.Idx:b.Lim]
	return &r
}
func (x *Tmes) From(idx unt.Unt) *Tmes {
	var r Tmes
	if idx < unt.Unt(len(*x)) {
		r = (*x)[idx:]
	} else {
		r = (*x)[:0]
	}
	return &r
}
func (x *Tmes) To(lim unt.Unt) *Tmes {
	if lim > unt.Unt(len(*x)) {
		return x
	}
	r := (*x)[:lim]
	return &r
}
func (x *Tmes) Fst() tme.Tme    { return (*x)[0] }
func (x *Tmes) Mdl() tme.Tme    { return (*x)[len(*x)/2] }
func (x *Tmes) Lst() tme.Tme    { return (*x)[len(*x)-1] }
func (x *Tmes) FstIdx() unt.Unt { return 0 }
func (x *Tmes) MdlIdx() unt.Unt { return unt.Unt(len(*x) / 2) }
func (x *Tmes) LstIdx() unt.Unt { return unt.Unt(len(*x) - 1) }
func (x *Tmes) Rev() *Tmes {
	for i, j := 0, len(*x)-1; i < j; i, j = i+1, j-1 {
		(*x)[i], (*x)[j] = (*x)[j], (*x)[i]
	}
	return x
}
func (x *Tmes) SrchIdxEql(v tme.Tme) unt.Unt {
	i, j := unt.Zero, unt.Unt(len(*x))
	for i < j {
		if (*x)[(i+j)>>1] < v {
			i = (i+j)>>1 + 1
		} else {
			j = (i + j) >> 1
		}
	}
	return i
}
func (x *Tmes) SrchIdx(v tme.Tme, near ...bol.Bol) unt.Unt {
	if len(*x) == 0 {
		return unt.Max
	}
	if len(near) > 0 {
		if v <= (*x)[0] { // lwr bnd
			return 0
		} else if v >= (*x)[len(*x)-1] { // upr bnd
			return unt.Unt(len(*x) - 1)
		}
	}
	idx := sort.Search(len(*x), func(i int) bool { return (*x)[i] >= v })
	if idx < len(*x) && (len(near) > 0 || (*x)[idx] == v) { // near does not require exact match; default requires exact match
		return unt.Unt(idx)
	}
	return unt.Max
}
func (x *Tmes) Has(v tme.Tme) bol.Bol {
	i, j := unt.Zero, unt.Unt(len(*x))
	for i < j {
		if (*x)[(i+j)>>1] < v {
			i = (i+j)>>1 + 1
		} else {
			j = (i + j) >> 1
		}
	}
	return i != unt.Unt(len(*x)) && (*x)[i] == v
}
func (x *Tmes) SrtAsc() *Tmes {
	if x.Cnt() > 1 {
		x.SrtQuick(0, x.LstIdx(), tme.Lss, tme.Eql)
	}
	return x
}
func (x *Tmes) SrtDsc() *Tmes {
	if x.Cnt() > 1 {
		x.SrtQuick(0, x.LstIdx(), tme.Gtr, tme.Eql)
	}
	return x
}
func (x *Tmes) SrtQuick(lo, hi unt.Unt, cmp, eql tme.Cmp) *Tmes {
	n := hi - lo + 1
	if n <= 8 { // cutoff to insertion sort
		return x.SrtIns(lo, hi, cmp)
	}
	if n <= 40 { // use median-of-3 as partitioning element
		mdn := x.SrtMdnOf3(lo, lo+n/2, hi, cmp)
		x.Swp(mdn, lo)
	} else { // use Tukey ninther as partitioning element
		eps := n / 8
		mid := lo + n/2
		m1 := x.SrtMdnOf3(lo, lo+eps, lo+eps+eps, cmp)
		m2 := x.SrtMdnOf3(mid-eps, mid, mid+eps, cmp)
		m3 := x.SrtMdnOf3(hi-eps-eps, hi-eps, hi, cmp)
		ninther := x.SrtMdnOf3(m1, m2, m3, cmp)
		x.Swp(ninther, lo)
	}
	i, j := lo, hi+1 // Bentley-McIlroy 3-way partitioning
	p, q := lo, hi+1
	v := (*x)[lo]
	for {
		i++
		for cmp((*x)[i], v) {
			if i == hi {
				break
			}
			i++
		}
		if j != 0 {
			j--
		}
		for cmp(v, (*x)[j]) {
			if j == lo {
				break
			}
			if j != 0 {
				j--
			}
		}
		if i == j && eql((*x)[i], v) { // pointers cross
			p++
			x.Swp(p, i)
		}
		if i >= j {
			break
		}
		x.Swp(i, j)
		if eql((*x)[i], v) {
			p++
			x.Swp(p, i)
		}
		if eql((*x)[j], v) {
			q--
			x.Swp(q, j)
		}
	}
	i = j + 1
	for k := lo; k <= p; k++ {
		x.Swp(k, j)
		if j != 0 {
			j--
		}
	}
	for k := hi; k >= q; k-- {
		x.Swp(k, i)
		i++
	}
	x.SrtQuick(lo, j, cmp, eql)
	x.SrtQuick(i, hi, cmp, eql)
	return x
}
func (x *Tmes) SrtIns(lo, hi unt.Unt, cmp tme.Cmp) *Tmes {
	for i := lo; i <= hi; i++ {
		for j := i; j > lo && cmp((*x)[j], (*x)[j-1]); j-- {
			x.Swp(j, j-1)
		}
	}
	return x
}
func (x *Tmes) SrtMdnOf3(i, j, k unt.Unt, cmp tme.Cmp) unt.Unt {
	if cmp((*x)[i], (*x)[j]) {
		if cmp((*x)[j], (*x)[k]) {
			return j
		}
		if cmp((*x)[i], (*x)[k]) {
			return k
		}
		return i
	}
	if cmp((*x)[k], (*x)[j]) {
		return j
	}
	if cmp((*x)[k], (*x)[i]) {
		return k
	}
	return i
}
func (x *Tmes) Swp(i, j unt.Unt) { (*x)[i], (*x)[j] = (*x)[j], (*x)[i] }
func (x *Tmes) InrAdd(off unt.Unt) *Tmes {
	if len(*x) < int(off) {
		r := make(Tmes, 0)
		return &r
	}
	r := make(Tmes, len(*x)-int(off))
	for n := 0; n < len(r); n++ {
		r[n] = (*x)[n+int(off)].Add((*x)[n])
	}
	return &r
}
func (x *Tmes) InrSub(off unt.Unt) *Tmes {
	if len(*x) < int(off) {
		r := make(Tmes, 0)
		return &r
	}
	r := make(Tmes, len(*x)-int(off))
	for n := 0; n < len(r); n++ {
		r[n] = (*x)[n+int(off)].Sub((*x)[n])
	}
	return &r
}
func (x *Tmes) InrMul(off unt.Unt) *Tmes {
	if len(*x) < int(off) {
		r := make(Tmes, 0)
		return &r
	}
	r := make(Tmes, len(*x)-int(off))
	for n := 0; n < len(r); n++ {
		r[n] = (*x)[n+int(off)].Mul((*x)[n])
	}
	return &r
}
func (x *Tmes) InrDiv(off unt.Unt) *Tmes {
	if len(*x) < int(off) {
		r := make(Tmes, 0)
		return &r
	}
	r := make(Tmes, len(*x)-int(off))
	for n := 0; n < len(r); n++ {
		r[n] = (*x)[n+int(off)].Div((*x)[n])
	}
	return &r
}
func (x *Tmes) InrRem(off unt.Unt) *Tmes {
	if len(*x) < int(off) {
		r := make(Tmes, 0)
		return &r
	}
	r := make(Tmes, len(*x)-int(off))
	for n := 0; n < len(r); n++ {
		r[n] = (*x)[n+int(off)].Rem((*x)[n])
	}
	return &r
}
func (x *Tmes) InrPow(off unt.Unt) *Tmes {
	if len(*x) < int(off) {
		r := make(Tmes, 0)
		return &r
	}
	r := make(Tmes, len(*x)-int(off))
	for n := 0; n < len(r); n++ {
		r[n] = (*x)[n+int(off)].Pow((*x)[n])
	}
	return &r
}
func (x *Tmes) InrMin(off unt.Unt) *Tmes {
	if len(*x) < int(off) {
		r := make(Tmes, 0)
		return &r
	}
	r := make(Tmes, len(*x)-int(off))
	for n := 0; n < len(r); n++ {
		r[n] = (*x)[n+int(off)].Min((*x)[n])
	}
	return &r
}
func (x *Tmes) InrMax(off unt.Unt) *Tmes {
	if len(*x) < int(off) {
		r := make(Tmes, 0)
		return &r
	}
	r := make(Tmes, len(*x)-int(off))
	for n := 0; n < len(r); n++ {
		r[n] = (*x)[n+int(off)].Max((*x)[n])
	}
	return &r
}
func (x *Tmes) Sum() (r tme.Tme) {
	for n := 0; n < len(*x); n++ {
		r += (*x)[n]
	}
	return r
}
func (x *Tmes) Prd() (r tme.Tme) {
	if len(*x) == 0 {
		return 0
	}
	r = (*x)[0]
	for n := 1; n < len(*x); n++ {
		r *= (*x)[n]
	}
	return r
}
func (x *Tmes) Min() (r tme.Tme) {
	if len(*x) == 0 {
		return 0
	}
	r = (*x)[0]
	for n := 1; n < len(*x); n++ {
		if (*x)[n] < r {
			r = (*x)[n]
		}
	}
	return r
}
func (x *Tmes) Max() (r tme.Tme) {
	if len(*x) == 0 {
		return 0
	}
	r = (*x)[0]
	for n := 1; n < len(*x); n++ {
		if (*x)[n] > r {
			r = (*x)[n]
		}
	}
	return r
}
func (x *MinMaxSeg) Act() {
	for n := x.Idx; n < x.Lim; n++ {
		if (*x.Vals)[n] < x.Min {
			x.Min = (*x.Vals)[n]
		}
		if (*x.Vals)[n] > x.Max {
			x.Max = (*x.Vals)[n]
		}
	}
}
func (x *Tmes) MinMax() (min, max tme.Tme) {
	if len(*x) == 0 {
		return 0, 0
	}
	if len(*x) == 1 {
		return (*x)[0], (*x)[0]
	}
	min, max = (*x)[0], (*x)[0]
	for n := 1; n < len(*x); n++ {
		if (*x)[n] < min {
			min = (*x)[n]
		}
		if (*x)[n] > max {
			max = (*x)[n]
		}
	}
	return min, max
}
func (x *Tmes) Mid() (r tme.Tme) {
	if len(*x) == 0 {
		return r
	}
	if len(*x) == 1 {
		return (*x)[0]
	}
	min, max := x.MinMax()
	return min + ((max - min) / 2)
}
func (x *Tmes) Mdn() (r tme.Tme) {
	if len(*x) == 0 {
		return r
	}
	if len(*x) == 1 {
		return (*x)[0]
	}
	cpy := x.Cpy()
	cpy.SrtAsc()
	if len(*cpy)%2 == 0 {
		if len(*cpy) == 2 {
			return (cpy.At(0) + cpy.At(1)) / 2
		}
		return (cpy.Mdl() + cpy.At(cpy.MdlIdx()+1)) / 2
	}
	return cpy.Mdl()
}
func (x *Tmes) Sma() (r tme.Tme) {
	if len(*x) == 0 { // simple moving average
		return r
	}
	if len(*x) == 1 {
		return (*x)[0]
	}
	return x.Sum() / tme.Tme(len(*x))
}
func (x *Tmes) Gma() (r tme.Tme) {
	if len(*x) == 0 { // geometric moving average
		return r
	}
	if len(*x) == 1 {
		return (*x)[0]
	}
	return x.Prd().Pow(tme.Tme(1)) / tme.Tme(len(*x))
}
func (x *Tmes) Wma() (r tme.Tme) {
	// For example, a 5 period WMA would be calculated as:
	// WMA = (P1 * 1) + (P2 * 2) + (P3 * 3) + (P4 * 4) + (P5 * 5) / (1 + 2 + 3 + 4 + 5)
	if len(*x) == 0 { // weighted moving average
		return 0
	}
	if len(*x) == 1 {
		return (*x)[0]
	}
	var numr, dnmr tme.Tme
	for n, v := range *x {
		numr += v * tme.Tme(n+1)
		dnmr += tme.Tme(n + 1)
	}
	if dnmr == 0 {
		return 0
	}
	return numr / dnmr
}
func (x *Tmes) Vrnc() (r tme.Tme) {
	if len(*x) == 0 {
		return 0
	}
	mean := x.Sma()
	meanDifSqrSum := r
	for _, v := range *x {
		meanDifSqrSum += (mean - v) * (mean - v) // calculate mean dif sqr sum
	}
	return meanDifSqrSum / tme.Tme(len(*x)) // calculate variance
}
func (x *Tmes) Std() (r tme.Tme) {
	if len(*x) == 0 {
		return 0
	}
	return x.Vrnc().Sqrt()
}
func (x *Tmes) Zscr() (r *Tmes) {
	if len(*x) == 0 {
		return New()
	}
	r = Make(x.Cnt())
	mean := x.Sma()
	std := x.Std()
	for n, v := range *x {
		(*r)[n] = (v - mean) / std
	}
	return r
}
func (x *Tmes) ZscrInplace() (r *Tmes) {
	if len(*x) == 0 {
		return x
	}
	mean := x.Sma()
	std := x.Std()
	for n, v := range *x {
		(*x)[n] = (v - mean) / std
	}
	return x
}
func (x *Tmes) RngFul() (r tme.Tme) {
	if len(*x) == 0 {
		return 0
	}
	min, max := x.MinMax()
	return max - min
}
func (x *Tmes) RngLst() (r tme.Tme) {
	if len(*x) == 0 {
		return 0
	}
	return x.Lst() - x.Min()
}
func (x *Tmes) ProLst() (r tme.Tme) {
	if len(*x) == 0 {
		return 0
	}
	min, max := x.MinMax()
	rngFul := max - min
	if rngFul == 0 {
		return r
	}
	return (x.Lst() - min) / rngFul
}
func (x *Tmes) ProSma() (r tme.Tme) {
	if len(*x) == 0 {
		return 0
	}
	min, max := x.MinMax()
	rngFul := max - min
	if rngFul == 0 {
		return r
	}
	return (x.Sma() - min) / rngFul
}
func (x *Tmes) StrWrt(b *strings.Builder) {
	b.WriteRune('[')
	for n, v := range *x {
		if n != 0 {
			b.WriteRune(' ')
		}
		v.StrWrt(b)
	}
	b.WriteRune(']')
}
func (x *Tmes) BytWrt(b *bytes.Buffer) {
	bLen := make([]byte, 4) // array length
	binary.LittleEndian.PutUint32(bLen, uint32(len(*x)))
	b.Write(bLen)
	for _, v := range *x {
		v.BytWrt(b)
	}
}
func (x *Tmes) BytRed(b []byte) (idx int) {
	if len(b) >= 4 {
		*x = make(Tmes, binary.LittleEndian.Uint32(b[:4])) // overwrite any previous existing
		idx = 4
		for n := 0; n < len(*x); n++ {
			(*x)[n].BytRed(b[idx : idx+tme.Size])
			idx += tme.Size
		}
	}
	return idx
}
