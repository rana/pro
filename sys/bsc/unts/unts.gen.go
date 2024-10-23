package unts

import (
	"bytes"
	"encoding/binary"
	"math/rand"
	"sort"
	"strings"
	"sys/bsc/bnd"
	"sys/bsc/bol"
	"sys/bsc/unt"
	"time"
)

type (
	Unts    []unt.Unt
	UntsScp struct {
		Idx uint32
		Arr []*Unts
	}
	MinMaxSeg struct {
		bnd.Bnd
		Min  unt.Unt
		Max  unt.Unt
		Vals *Unts
	}
)

func New(vs ...unt.Unt) *Unts {
	r := Unts(vs)
	return &r
}
func Make(cap unt.Unt) *Unts {
	r := make(Unts, cap)
	return &r
}
func MakeEmp(cap unt.Unt) *Unts {
	r := make(Unts, 0, cap)
	return &r
}
func AddsLss(strt, lim, by unt.Unt) (r *Unts) {
	r = New()
	for n := strt; n < lim; n += by {
		r.Push(n)
	}
	return r
}
func AddsLeq(strt, lim, by unt.Unt) (r *Unts) {
	r = New()
	for n := strt; n <= lim; n += by {
		r.Push(n)
	}
	return r
}
func SubsGtr(strt, lim, by unt.Unt) (r *Unts) {
	r = New()
	for n := strt; n > lim; n -= by {
		r.Push(n)
	}
	return r
}
func SubsGeq(strt, lim, by unt.Unt) (r *Unts) {
	r = New()
	for n := strt; n >= lim; n -= by {
		r.Push(n)
	}
	return r
}
func MulsLss(strt, lim, by unt.Unt) (r *Unts) {
	r = New()
	for n := strt; n < lim; n *= by {
		r.Push(n)
	}
	return r
}
func MulsLeq(strt, lim, by unt.Unt) (r *Unts) {
	r = New()
	for n := strt; n <= lim; n *= by {
		r.Push(n)
	}
	return r
}
func DivsGtr(strt, lim, by unt.Unt) (r *Unts) {
	r = New()
	for n := strt; n > lim; n /= by {
		r.Push(n)
	}
	return r
}
func DivsGeq(strt, lim, by unt.Unt) (r *Unts) {
	r = New()
	for n := strt; n >= lim; n /= by {
		r.Push(n)
	}
	return r
}
func FibsLeq(lim unt.Unt) (r *Unts) {
	r = New(1, 2)
	for (*r)[len(*r)-1] < lim {
		*r = append(*r, (*r)[len(*r)-2]+(*r)[len(*r)-1])
	}
	return r
}
func (x *Unts) Ok() bol.Bol { return len(*x) != 0 }
func (x *Unts) Cnt() unt.Unt {
	if x == nil {
		return 0
	}
	return unt.Unt(len(*x))
}
func (x *Unts) Cpy() *Unts {
	r := make(Unts, len(*x))
	copy(r, *x)
	return &r
}
func (x *Unts) Clr() *Unts {
	*x = (*x)[:0]
	return x
}
func (x *Unts) Rand() *Unts {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	perm := r.Perm(len(*x))
	for i, randIdx := range perm {
		(*x)[i], (*x)[randIdx] = (*x)[randIdx], (*x)[i]
	}
	return x
}
func (x *Unts) Mrg(a ...*Unts) *Unts {
	for _, v := range a {
		*x = append(*x, *v...)
	}
	return x
}
func (x *Unts) Push(a ...unt.Unt) *Unts {
	*x = append(*x, a...)
	return x
}
func (x *Unts) Pop() (r unt.Unt) {
	if len(*x) == 0 {
		return r
	}
	r = (*x)[len(*x)-1]
	*x = (*x)[:len(*x)-1]
	return r
}
func (x *Unts) Que(vs ...unt.Unt) *Unts {
	*x = append(*x, vs...)
	return x
}
func (x *Unts) Dque() (r unt.Unt) {
	r = (*x)[0]
	if len(*x) == 1 {
		*x = (*x)[:0]
	} else {
		copy(*x, (*x)[1:])
		*x = (*x)[:len(*x)-1]
	}
	return r
}
func (x *Unts) Ins(idx unt.Unt, elm unt.Unt) *Unts {
	*x = append((*x)[:idx], append(Unts{elm}, (*x)[idx:]...)...)
	return x
}
func (x *Unts) Upd(idx unt.Unt, elm unt.Unt) *Unts {
	(*x)[idx] = elm
	return x
}
func (x *Unts) Del(idx unt.Unt) (r unt.Unt) {
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
func (x *Unts) At(idx unt.Unt) unt.Unt { return (*x)[idx] }
func (x *Unts) In(idx, lim unt.Unt) *Unts {
	r := (*x)[idx:lim]
	return &r
}
func (x *Unts) InBnd(b bnd.Bnd) *Unts {
	r := (*x)[b.Idx:b.Lim]
	return &r
}
func (x *Unts) From(idx unt.Unt) *Unts {
	var r Unts
	if idx < unt.Unt(len(*x)) {
		r = (*x)[idx:]
	} else {
		r = (*x)[:0]
	}
	return &r
}
func (x *Unts) To(lim unt.Unt) *Unts {
	if lim > unt.Unt(len(*x)) {
		return x
	}
	r := (*x)[:lim]
	return &r
}
func (x *Unts) Fst() unt.Unt    { return (*x)[0] }
func (x *Unts) Mdl() unt.Unt    { return (*x)[len(*x)/2] }
func (x *Unts) Lst() unt.Unt    { return (*x)[len(*x)-1] }
func (x *Unts) FstIdx() unt.Unt { return 0 }
func (x *Unts) MdlIdx() unt.Unt { return unt.Unt(len(*x) / 2) }
func (x *Unts) LstIdx() unt.Unt { return unt.Unt(len(*x) - 1) }
func (x *Unts) Rev() *Unts {
	for i, j := 0, len(*x)-1; i < j; i, j = i+1, j-1 {
		(*x)[i], (*x)[j] = (*x)[j], (*x)[i]
	}
	return x
}
func (x *Unts) SrchIdxEql(v unt.Unt) unt.Unt {
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
func (x *Unts) SrchIdx(v unt.Unt, near ...bol.Bol) unt.Unt {
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
func (x *Unts) Has(v unt.Unt) bol.Bol {
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
func (x *Unts) SrtAsc() *Unts {
	if x.Cnt() > 1 {
		x.SrtQuick(0, x.LstIdx(), unt.Lss, unt.Eql)
	}
	return x
}
func (x *Unts) SrtDsc() *Unts {
	if x.Cnt() > 1 {
		x.SrtQuick(0, x.LstIdx(), unt.Gtr, unt.Eql)
	}
	return x
}
func (x *Unts) SrtQuick(lo, hi unt.Unt, cmp, eql unt.Cmp) *Unts {
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
func (x *Unts) SrtIns(lo, hi unt.Unt, cmp unt.Cmp) *Unts {
	for i := lo; i <= hi; i++ {
		for j := i; j > lo && cmp((*x)[j], (*x)[j-1]); j-- {
			x.Swp(j, j-1)
		}
	}
	return x
}
func (x *Unts) SrtMdnOf3(i, j, k unt.Unt, cmp unt.Cmp) unt.Unt {
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
func (x *Unts) Swp(i, j unt.Unt) { (*x)[i], (*x)[j] = (*x)[j], (*x)[i] }
func (x *Unts) InrAdd(off unt.Unt) *Unts {
	if len(*x) < int(off) {
		r := make(Unts, 0)
		return &r
	}
	r := make(Unts, len(*x)-int(off))
	for n := 0; n < len(r); n++ {
		r[n] = (*x)[n+int(off)].Add((*x)[n])
	}
	return &r
}
func (x *Unts) InrSub(off unt.Unt) *Unts {
	if len(*x) < int(off) {
		r := make(Unts, 0)
		return &r
	}
	r := make(Unts, len(*x)-int(off))
	for n := 0; n < len(r); n++ {
		r[n] = (*x)[n+int(off)].Sub((*x)[n])
	}
	return &r
}
func (x *Unts) InrMul(off unt.Unt) *Unts {
	if len(*x) < int(off) {
		r := make(Unts, 0)
		return &r
	}
	r := make(Unts, len(*x)-int(off))
	for n := 0; n < len(r); n++ {
		r[n] = (*x)[n+int(off)].Mul((*x)[n])
	}
	return &r
}
func (x *Unts) InrDiv(off unt.Unt) *Unts {
	if len(*x) < int(off) {
		r := make(Unts, 0)
		return &r
	}
	r := make(Unts, len(*x)-int(off))
	for n := 0; n < len(r); n++ {
		r[n] = (*x)[n+int(off)].Div((*x)[n])
	}
	return &r
}
func (x *Unts) InrRem(off unt.Unt) *Unts {
	if len(*x) < int(off) {
		r := make(Unts, 0)
		return &r
	}
	r := make(Unts, len(*x)-int(off))
	for n := 0; n < len(r); n++ {
		r[n] = (*x)[n+int(off)].Rem((*x)[n])
	}
	return &r
}
func (x *Unts) InrPow(off unt.Unt) *Unts {
	if len(*x) < int(off) {
		r := make(Unts, 0)
		return &r
	}
	r := make(Unts, len(*x)-int(off))
	for n := 0; n < len(r); n++ {
		r[n] = (*x)[n+int(off)].Pow((*x)[n])
	}
	return &r
}
func (x *Unts) InrMin(off unt.Unt) *Unts {
	if len(*x) < int(off) {
		r := make(Unts, 0)
		return &r
	}
	r := make(Unts, len(*x)-int(off))
	for n := 0; n < len(r); n++ {
		r[n] = (*x)[n+int(off)].Min((*x)[n])
	}
	return &r
}
func (x *Unts) InrMax(off unt.Unt) *Unts {
	if len(*x) < int(off) {
		r := make(Unts, 0)
		return &r
	}
	r := make(Unts, len(*x)-int(off))
	for n := 0; n < len(r); n++ {
		r[n] = (*x)[n+int(off)].Max((*x)[n])
	}
	return &r
}
func (x *Unts) Sum() (r unt.Unt) {
	for n := 0; n < len(*x); n++ {
		r += (*x)[n]
	}
	return r
}
func (x *Unts) Prd() (r unt.Unt) {
	if len(*x) == 0 {
		return 0
	}
	r = (*x)[0]
	for n := 1; n < len(*x); n++ {
		r *= (*x)[n]
	}
	return r
}
func (x *Unts) Min() (r unt.Unt) {
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
func (x *Unts) Max() (r unt.Unt) {
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
func (x *Unts) MinMax() (min, max unt.Unt) {
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
func (x *Unts) Mid() (r unt.Unt) {
	if len(*x) == 0 {
		return r
	}
	if len(*x) == 1 {
		return (*x)[0]
	}
	min, max := x.MinMax()
	return min + ((max - min) / 2)
}
func (x *Unts) Mdn() (r unt.Unt) {
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
func (x *Unts) Sma() (r unt.Unt) {
	if len(*x) == 0 { // simple moving average
		return r
	}
	if len(*x) == 1 {
		return (*x)[0]
	}
	return x.Sum() / unt.Unt(len(*x))
}
func (x *Unts) Gma() (r unt.Unt) {
	if len(*x) == 0 { // geometric moving average
		return r
	}
	if len(*x) == 1 {
		return (*x)[0]
	}
	return x.Prd().Pow(unt.Unt(1)) / unt.Unt(len(*x))
}
func (x *Unts) Wma() (r unt.Unt) {
	// For example, a 5 period WMA would be calculated as:
	// WMA = (P1 * 1) + (P2 * 2) + (P3 * 3) + (P4 * 4) + (P5 * 5) / (1 + 2 + 3 + 4 + 5)
	if len(*x) == 0 { // weighted moving average
		return 0
	}
	if len(*x) == 1 {
		return (*x)[0]
	}
	var numr, dnmr unt.Unt
	for n, v := range *x {
		numr += v * unt.Unt(n+1)
		dnmr += unt.Unt(n + 1)
	}
	if dnmr == 0 {
		return 0
	}
	return numr / dnmr
}
func (x *Unts) Vrnc() (r unt.Unt) {
	if len(*x) == 0 {
		return 0
	}
	mean := x.Sma()
	meanDifSqrSum := r
	for _, v := range *x {
		meanDifSqrSum += (mean - v) * (mean - v) // calculate mean dif sqr sum
	}
	return meanDifSqrSum / unt.Unt(len(*x)) // calculate variance
}
func (x *Unts) Std() (r unt.Unt) {
	if len(*x) == 0 {
		return 0
	}
	return x.Vrnc().Sqrt()
}
func (x *Unts) Zscr() (r *Unts) {
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
func (x *Unts) ZscrInplace() (r *Unts) {
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
func (x *Unts) RngFul() (r unt.Unt) {
	if len(*x) == 0 {
		return 0
	}
	min, max := x.MinMax()
	return max - min
}
func (x *Unts) RngLst() (r unt.Unt) {
	if len(*x) == 0 {
		return 0
	}
	return x.Lst() - x.Min()
}
func (x *Unts) ProLst() (r unt.Unt) {
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
func (x *Unts) ProSma() (r unt.Unt) {
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
func (x *Unts) StrWrt(b *strings.Builder) {
	b.WriteRune('[')
	for n, v := range *x {
		if n != 0 {
			b.WriteRune(' ')
		}
		v.StrWrt(b)
	}
	b.WriteRune(']')
}
func (x *Unts) BytWrt(b *bytes.Buffer) {
	bLen := make([]byte, 4) // array length
	binary.LittleEndian.PutUint32(bLen, uint32(len(*x)))
	b.Write(bLen)
	for _, v := range *x {
		v.BytWrt(b)
	}
}
func (x *Unts) BytRed(b []byte) (idx int) {
	if len(b) >= 4 {
		*x = make(Unts, binary.LittleEndian.Uint32(b[:4])) // overwrite any previous existing
		idx = 4
		for n := 0; n < len(*x); n++ {
			(*x)[n].BytRed(b[idx : idx+unt.Size])
			idx += unt.Size
		}
	}
	return idx
}
