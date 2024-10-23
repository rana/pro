package flts

import (
	"bytes"
	"encoding/binary"
	"math"
	"math/rand"
	"sort"
	"strings"
	"sys"
	"sys/bsc/bnd"
	"sys/bsc/bnds"
	"sys/bsc/bol"
	"sys/bsc/flt"
	"sys/bsc/unt"
	"time"
)

type (
	Flts    []flt.Flt
	FltsScp struct {
		Idx uint32
		Arr []*Flts
	}
	MinMaxSeg struct {
		bnd.Bnd
		Min  flt.Flt
		Max  flt.Flt
		Vals *Flts
	}
	CntrDistSeg struct {
		bnd.Bnd
		EvalZero bol.Bol
		RngFul   flt.Flt
		Vals     *Flts
		Out      *Flts
	}
)

func New(vs ...flt.Flt) *Flts {
	r := Flts(vs)
	return &r
}
func Make(cap unt.Unt) *Flts {
	r := make(Flts, cap)
	return &r
}
func MakeEmp(cap unt.Unt) *Flts {
	r := make(Flts, 0, cap)
	return &r
}
func AddsLss(strt, lim, by flt.Flt) (r *Flts) {
	r = New()
	for n := strt; n < lim; n += by {
		r.Push(n)
	}
	return r
}
func AddsLeq(strt, lim, by flt.Flt) (r *Flts) {
	r = New()
	for n := strt; n <= lim; n += by {
		r.Push(n)
	}
	return r
}
func SubsGtr(strt, lim, by flt.Flt) (r *Flts) {
	r = New()
	for n := strt; n > lim; n -= by {
		r.Push(n)
	}
	return r
}
func SubsGeq(strt, lim, by flt.Flt) (r *Flts) {
	r = New()
	for n := strt; n >= lim; n -= by {
		r.Push(n)
	}
	return r
}
func MulsLss(strt, lim, by flt.Flt) (r *Flts) {
	r = New()
	for n := strt; n < lim; n *= by {
		r.Push(n)
	}
	return r
}
func MulsLeq(strt, lim, by flt.Flt) (r *Flts) {
	r = New()
	for n := strt; n <= lim; n *= by {
		r.Push(n)
	}
	return r
}
func DivsGtr(strt, lim, by flt.Flt) (r *Flts) {
	r = New()
	for n := strt; n > lim; n /= by {
		r.Push(n)
	}
	return r
}
func DivsGeq(strt, lim, by flt.Flt) (r *Flts) {
	r = New()
	for n := strt; n >= lim; n /= by {
		r.Push(n)
	}
	return r
}
func FibsLeq(lim flt.Flt) (r *Flts) {
	r = New(1, 2)
	for (*r)[len(*r)-1] < lim {
		*r = append(*r, (*r)[len(*r)-2]+(*r)[len(*r)-1])
	}
	return r
}
func (x *Flts) Ok() bol.Bol { return len(*x) != 0 }
func (x *Flts) Cnt() unt.Unt {
	if x == nil {
		return 0
	}
	return unt.Unt(len(*x))
}
func (x *Flts) Cpy() *Flts {
	r := make(Flts, len(*x))
	copy(r, *x)
	return &r
}
func (x *Flts) Clr() *Flts {
	*x = (*x)[:0]
	return x
}
func (x *Flts) Rand() *Flts {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	perm := r.Perm(len(*x))
	for i, randIdx := range perm {
		(*x)[i], (*x)[randIdx] = (*x)[randIdx], (*x)[i]
	}
	return x
}
func (x *Flts) Mrg(a ...*Flts) *Flts {
	for _, v := range a {
		*x = append(*x, *v...)
	}
	return x
}
func (x *Flts) Push(a ...flt.Flt) *Flts {
	*x = append(*x, a...)
	return x
}
func (x *Flts) Pop() (r flt.Flt) {
	if len(*x) == 0 {
		return r
	}
	r = (*x)[len(*x)-1]
	*x = (*x)[:len(*x)-1]
	return r
}
func (x *Flts) Que(vs ...flt.Flt) *Flts {
	*x = append(*x, vs...)
	return x
}
func (x *Flts) Dque() (r flt.Flt) {
	r = (*x)[0]
	if len(*x) == 1 {
		*x = (*x)[:0]
	} else {
		copy(*x, (*x)[1:])
		*x = (*x)[:len(*x)-1]
	}
	return r
}
func (x *Flts) Ins(idx unt.Unt, elm flt.Flt) *Flts {
	*x = append((*x)[:idx], append(Flts{elm}, (*x)[idx:]...)...)
	return x
}
func (x *Flts) Upd(idx unt.Unt, elm flt.Flt) *Flts {
	(*x)[idx] = elm
	return x
}
func (x *Flts) Del(idx unt.Unt) (r flt.Flt) {
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
func (x *Flts) At(idx unt.Unt) flt.Flt { return (*x)[idx] }
func (x *Flts) In(idx, lim unt.Unt) *Flts {
	r := (*x)[idx:lim]
	return &r
}
func (x *Flts) InBnd(b bnd.Bnd) *Flts {
	r := (*x)[b.Idx:b.Lim]
	return &r
}
func (x *Flts) From(idx unt.Unt) *Flts {
	var r Flts
	if idx < unt.Unt(len(*x)) {
		r = (*x)[idx:]
	} else {
		r = (*x)[:0]
	}
	return &r
}
func (x *Flts) To(lim unt.Unt) *Flts {
	if lim > unt.Unt(len(*x)) {
		return x
	}
	r := (*x)[:lim]
	return &r
}
func (x *Flts) Fst() flt.Flt    { return (*x)[0] }
func (x *Flts) Mdl() flt.Flt    { return (*x)[len(*x)/2] }
func (x *Flts) Lst() flt.Flt    { return (*x)[len(*x)-1] }
func (x *Flts) FstIdx() unt.Unt { return 0 }
func (x *Flts) MdlIdx() unt.Unt { return unt.Unt(len(*x) / 2) }
func (x *Flts) LstIdx() unt.Unt { return unt.Unt(len(*x) - 1) }
func (x *Flts) Rev() *Flts {
	for i, j := 0, len(*x)-1; i < j; i, j = i+1, j-1 {
		(*x)[i], (*x)[j] = (*x)[j], (*x)[i]
	}
	return x
}
func (x *Flts) SrchIdxEql(v flt.Flt) unt.Unt {
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
func (x *Flts) SrchIdx(v flt.Flt, near ...bol.Bol) unt.Unt {
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
func (x *Flts) Has(v flt.Flt) bol.Bol {
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
func (x *Flts) SrtAsc() *Flts {
	if x.Cnt() > 1 {
		x.SrtQuick(0, x.LstIdx(), flt.Lss, flt.Eql)
	}
	return x
}
func (x *Flts) SrtDsc() *Flts {
	if x.Cnt() > 1 {
		x.SrtQuick(0, x.LstIdx(), flt.Gtr, flt.Eql)
	}
	return x
}
func (x *Flts) SrtQuick(lo, hi unt.Unt, cmp, eql flt.Cmp) *Flts {
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
func (x *Flts) SrtIns(lo, hi unt.Unt, cmp flt.Cmp) *Flts {
	for i := lo; i <= hi; i++ {
		for j := i; j > lo && cmp((*x)[j], (*x)[j-1]); j-- {
			x.Swp(j, j-1)
		}
	}
	return x
}
func (x *Flts) SrtMdnOf3(i, j, k unt.Unt, cmp flt.Cmp) unt.Unt {
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
func (x *Flts) Swp(i, j unt.Unt) { (*x)[i], (*x)[j] = (*x)[j], (*x)[i] }
func (x *Flts) UnaPos() *Flts {
	r := make(Flts, len(*x))
	for n, v := range *x {
		r[n] = v.Pos()
	}
	return &r
}
func (x *Flts) UnaNeg() *Flts {
	r := make(Flts, len(*x))
	for n, v := range *x {
		r[n] = v.Neg()
	}
	return &r
}
func (x *Flts) UnaInv() *Flts {
	r := make(Flts, len(*x))
	for n, v := range *x {
		r[n] = v.Inv()
	}
	return &r
}
func (x *Flts) UnaSqr() *Flts {
	r := make(Flts, len(*x))
	for n, v := range *x {
		r[n] = v.Sqr()
	}
	return &r
}
func (x *Flts) UnaSqrt() *Flts {
	r := make(Flts, len(*x))
	for n, v := range *x {
		r[n] = v.Sqrt()
	}
	return &r
}
func (x *Flts) SclAdd(scl flt.Flt) *Flts {
	r := make(Flts, len(*x))
	for n, v := range *x {
		r[n] = v.Add(scl)
	}
	return &r
}
func (x *Flts) SclSub(scl flt.Flt) *Flts {
	r := make(Flts, len(*x))
	for n, v := range *x {
		r[n] = v.Sub(scl)
	}
	return &r
}
func (x *Flts) SclMul(scl flt.Flt) *Flts {
	r := make(Flts, len(*x))
	for n, v := range *x {
		r[n] = v.Mul(scl)
	}
	return &r
}
func (x *Flts) SclDiv(scl flt.Flt) *Flts {
	r := make(Flts, len(*x))
	for n, v := range *x {
		r[n] = v.Div(scl)
	}
	return &r
}
func (x *Flts) SclRem(scl flt.Flt) *Flts {
	r := make(Flts, len(*x))
	for n, v := range *x {
		r[n] = v.Rem(scl)
	}
	return &r
}
func (x *Flts) SclPow(scl flt.Flt) *Flts {
	r := make(Flts, len(*x))
	for n, v := range *x {
		r[n] = v.Pow(scl)
	}
	return &r
}
func (x *Flts) SclMin(scl flt.Flt) *Flts {
	r := make(Flts, len(*x))
	for n, v := range *x {
		r[n] = v.Min(scl)
	}
	return &r
}
func (x *Flts) SclMax(scl flt.Flt) *Flts {
	r := make(Flts, len(*x))
	for n, v := range *x {
		r[n] = v.Max(scl)
	}
	return &r
}
func (x *Flts) SelEql(sel flt.Flt) (r *Flts) {
	r = Make(x.Cnt())
	for n, v := range *x {
		(*r)[n] = v.SelEql(sel)
	}
	return r
}
func (x *Flts) SelNeq(sel flt.Flt) (r *Flts) {
	r = Make(x.Cnt())
	for n, v := range *x {
		(*r)[n] = v.SelNeq(sel)
	}
	return r
}
func (x *Flts) SelLss(sel flt.Flt) (r *Flts) {
	r = Make(x.Cnt())
	for n, v := range *x {
		(*r)[n] = v.SelLss(sel)
	}
	return r
}
func (x *Flts) SelGtr(sel flt.Flt) (r *Flts) {
	r = Make(x.Cnt())
	for n, v := range *x {
		(*r)[n] = v.SelGtr(sel)
	}
	return r
}
func (x *Flts) SelLeq(sel flt.Flt) (r *Flts) {
	r = Make(x.Cnt())
	for n, v := range *x {
		(*r)[n] = v.SelLeq(sel)
	}
	return r
}
func (x *Flts) SelGeq(sel flt.Flt) (r *Flts) {
	r = Make(x.Cnt())
	for n, v := range *x {
		(*r)[n] = v.SelGeq(sel)
	}
	return r
}
func (x *Flts) Splt(v flt.Flt) (btm, top *Flts) {
	btm = New()
	top = New()
	for _, cur := range *x {
		if cur > v {
			*top = append(*top, cur)
		} else {
			*btm = append(*btm, cur)
		}
	}
	return btm, top
}
func (x *Flts) CntEql(pnt flt.Flt) (r flt.Flt) {
	for _, v := range *x {
		if v.Eql(pnt) {
			r++
		}
	}
	return r
}
func (x *Flts) CntNeq(pnt flt.Flt) (r flt.Flt) {
	for _, v := range *x {
		if v.Neq(pnt) {
			r++
		}
	}
	return r
}
func (x *Flts) CntLss(pnt flt.Flt) (r flt.Flt) {
	for _, v := range *x {
		if v.Lss(pnt) {
			r++
		}
	}
	return r
}
func (x *Flts) CntGtr(pnt flt.Flt) (r flt.Flt) {
	for _, v := range *x {
		if v.Gtr(pnt) {
			r++
		}
	}
	return r
}
func (x *Flts) CntLeq(pnt flt.Flt) (r flt.Flt) {
	for _, v := range *x {
		if v.Leq(pnt) {
			r++
		}
	}
	return r
}
func (x *Flts) CntGeq(pnt flt.Flt) (r flt.Flt) {
	for _, v := range *x {
		if v.Geq(pnt) {
			r++
		}
	}
	return r
}
func (x *Flts) InrAdd(off unt.Unt) *Flts {
	if len(*x) < int(off) {
		r := make(Flts, 0)
		return &r
	}
	r := make(Flts, len(*x)-int(off))
	for n := 0; n < len(r); n++ {
		r[n] = (*x)[n+int(off)].Add((*x)[n])
	}
	return &r
}
func (x *Flts) InrSub(off unt.Unt) *Flts {
	if len(*x) < int(off) {
		r := make(Flts, 0)
		return &r
	}
	r := make(Flts, len(*x)-int(off))
	for n := 0; n < len(r); n++ {
		r[n] = (*x)[n+int(off)].Sub((*x)[n])
	}
	return &r
}
func (x *Flts) InrMul(off unt.Unt) *Flts {
	if len(*x) < int(off) {
		r := make(Flts, 0)
		return &r
	}
	r := make(Flts, len(*x)-int(off))
	for n := 0; n < len(r); n++ {
		r[n] = (*x)[n+int(off)].Mul((*x)[n])
	}
	return &r
}
func (x *Flts) InrDiv(off unt.Unt) *Flts {
	if len(*x) < int(off) {
		r := make(Flts, 0)
		return &r
	}
	r := make(Flts, len(*x)-int(off))
	for n := 0; n < len(r); n++ {
		r[n] = (*x)[n+int(off)].Div((*x)[n])
	}
	return &r
}
func (x *Flts) InrRem(off unt.Unt) *Flts {
	if len(*x) < int(off) {
		r := make(Flts, 0)
		return &r
	}
	r := make(Flts, len(*x)-int(off))
	for n := 0; n < len(r); n++ {
		r[n] = (*x)[n+int(off)].Rem((*x)[n])
	}
	return &r
}
func (x *Flts) InrPow(off unt.Unt) *Flts {
	if len(*x) < int(off) {
		r := make(Flts, 0)
		return &r
	}
	r := make(Flts, len(*x)-int(off))
	for n := 0; n < len(r); n++ {
		r[n] = (*x)[n+int(off)].Pow((*x)[n])
	}
	return &r
}
func (x *Flts) InrMin(off unt.Unt) *Flts {
	if len(*x) < int(off) {
		r := make(Flts, 0)
		return &r
	}
	r := make(Flts, len(*x)-int(off))
	for n := 0; n < len(r); n++ {
		r[n] = (*x)[n+int(off)].Min((*x)[n])
	}
	return &r
}
func (x *Flts) InrMax(off unt.Unt) *Flts {
	if len(*x) < int(off) {
		r := make(Flts, 0)
		return &r
	}
	r := make(Flts, len(*x)-int(off))
	for n := 0; n < len(r); n++ {
		r[n] = (*x)[n+int(off)].Max((*x)[n])
	}
	return &r
}
func (x *Flts) Sum() (r flt.Flt) {
	for n := 0; n < len(*x); n++ {
		r += (*x)[n]
	}
	return r
}
func (x *Flts) Prd() (r flt.Flt) {
	if len(*x) == 0 {
		return 0
	}
	r = (*x)[0]
	for n := 1; n < len(*x); n++ {
		r *= (*x)[n]
	}
	return r
}
func (x *Flts) Min() (r flt.Flt) {
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
func (x *Flts) Max() (r flt.Flt) {
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
func (x *Flts) MinMax() (min, max flt.Flt) {
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
func (x *Flts) Mid() (r flt.Flt) {
	if len(*x) == 0 {
		return r
	}
	if len(*x) == 1 {
		return (*x)[0]
	}
	min, max := x.MinMax()
	return min + ((max - min) / 2)
}
func (x *Flts) Mdn() (r flt.Flt) {
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
func (x *Flts) Sma() (r flt.Flt) {
	if len(*x) == 0 { // simple moving average
		return r
	}
	if len(*x) == 1 {
		return (*x)[0]
	}
	return x.Sum() / flt.Flt(len(*x))
}
func (x *Flts) Gma() (r flt.Flt) {
	if len(*x) == 0 { // geometric moving average
		return r
	}
	if len(*x) == 1 {
		return (*x)[0]
	}
	return x.Prd().Pow(flt.Flt(1)) / flt.Flt(len(*x))
}
func (x *Flts) Wma() (r flt.Flt) {
	// For example, a 5 period WMA would be calculated as:
	// WMA = (P1 * 1) + (P2 * 2) + (P3 * 3) + (P4 * 4) + (P5 * 5) / (1 + 2 + 3 + 4 + 5)
	if len(*x) == 0 { // weighted moving average
		return 0
	}
	if len(*x) == 1 {
		return (*x)[0]
	}
	var numr, dnmr flt.Flt
	for n, v := range *x {
		numr += v * flt.Flt(n+1)
		dnmr += flt.Flt(n + 1)
	}
	if dnmr == 0 {
		return 0
	}
	return numr / dnmr
}
func (x *Flts) Vrnc() (r flt.Flt) {
	if len(*x) == 0 {
		return 0
	}
	mean := x.Sma()
	meanDifSqrSum := r
	for _, v := range *x {
		meanDifSqrSum += (mean - v) * (mean - v) // calculate mean dif sqr sum
	}
	return meanDifSqrSum / flt.Flt(len(*x)) // calculate variance
}
func (x *Flts) Std() (r flt.Flt) {
	if len(*x) == 0 {
		return 0
	}
	return x.Vrnc().Sqrt()
}
func (x *Flts) Zscr() (r *Flts) {
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
func (x *Flts) ZscrInplace() (r *Flts) {
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
func (x *Flts) RngFul() (r flt.Flt) {
	if len(*x) == 0 {
		return 0
	}
	min, max := x.MinMax()
	return max - min
}
func (x *Flts) RngLst() (r flt.Flt) {
	if len(*x) == 0 {
		return 0
	}
	return x.Lst() - x.Min()
}
func (x *Flts) ProLst() (r flt.Flt) {
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
func (x *Flts) ProSma() (r flt.Flt) {
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
func (x *Flts) SubSumPos() (r flt.Flt) {
	// positive sum of subtracted elements
	if len(*x) <= 1 {
		return 0
	}
	// skp eql is only issue for calculation on tic streams
	var prv, cur int
	for ; prv < len(*x)-1; prv = cur {
		cur = prv + 1
		if (*x)[cur] == (*x)[prv] { // skp eql due to gap fil; may alter result if valid, non-gap eql present
			continue
		}
		if (*x)[cur]-(*x)[prv] > 0 {
			r += (*x)[cur] - (*x)[prv]
		}
	}
	return r
}
func (x *Flts) SubSumNeg() (r flt.Flt) {
	// negative sum of subtracted elements
	if len(*x) <= 1 {
		return 0
	}
	// skp eql is only issue for calculation on tic streams
	var prv, cur int
	for ; prv < len(*x)-1; prv = cur {
		cur = prv + 1
		if (*x)[cur] == (*x)[prv] { // skp eql due to gap fil; may alter result if valid, non-gap eql present
			continue
		}
		if (*x)[cur]-(*x)[prv] < 0 {
			r += (*x)[cur] - (*x)[prv]
		}
	}
	return r
}
func (x *Flts) Rsi() (r flt.Flt) {
	// relative strength index
	// RS = Average Gain / Average Loss
	//               100
	// RSI = 100 - --------
	//              1 + RS
	// NOTE: This impl has a scale of 0.0 to 1.0
	//       and returns 0.5 as a default value
	if len(*x) <= 1 {
		return .5 // return mdl
	}
	// skp eql is only issue for calculation on tic streams
	var neg, pos flt.Flt
	var prv, cur int
	for ; prv < len(*x)-1; prv = cur {
		cur = prv + 1
		if (*x)[cur] == (*x)[prv] { // skp eql due to gap fil; may alter result if valid, non-gap eql present
			continue
		}
		if (*x)[cur]-(*x)[prv] < 0 {
			neg -= (*x)[cur] - (*x)[prv]
		} else {
			pos += (*x)[cur] - (*x)[prv]
		}
	}
	if neg == 0 && pos == 0 {
		return .5 // return mdl
	}
	return pos / (pos + neg) // rng is 0 to 1
}
func (x *Flts) Wrsi() (r flt.Flt) {
	if len(*x) <= 1 {
		return .5 // return mdl
	}
	var neg, pos flt.Flt
	for n := 1; n < len(*x); n++ {
		if (*x)[n]-(*x)[n-1] < 0 {
			neg -= ((*x)[n] - (*x)[n-1]) * flt.Flt(n+1)
		} else {
			pos += ((*x)[n] - (*x)[n-1]) * flt.Flt(n+1)
		}
	}
	if neg == 0 && pos == 0 {
		return .5 // return mdl
	}
	return pos / (pos + neg) // rng is 0 to 1
}
func (x *Flts) Pro() (r *Flts) {
	if len(*x) == 0 {
		return Make(0) // IMPORTANT FOR PLT TO RETURN EMPTY (NOT NIL)
	}
	min, max := x.MinMax()
	rngFul := max - min
	if rngFul == 0 {
		return Make(x.Cnt()) // IMPORTANT FOR PLT TO RETURN EMPTY (NOT NIL)
	}
	pros := make(Flts, len(*x))
	for n := 0; n < len(pros); n++ {
		pros[n] = ((*x)[n] - min) / rngFul
	}
	return &pros
}
func (x *Flts) Alma() (alma flt.Flt) {
	if len(*x) == 0 {
		return 0
	}
	// http://www.financial-hacker.com/trend-delusion-or-reality/
	// https://www.prorealcode.com/prorealtime-indicators/alma-arnaud-legoux-moving-average/
	// Window = 9
	// Sigma = 6
	// Offset = 0.85
	// var ALMA(var *Data, int Period) {
	// var m = floor(0.85*(Period-1));
	// var s = Period/6.0;
	// var alma = 0., wSum = 0.;
	// int i;
	// for (i = 0; i < Period; i++) {
	// var w = exp(-(i-m)*(i-m)/(2*s*s));
	// alma += Data[Period-1-i] * w;
	// wSum += w;
	// }
	// return alma / wSum;
	// }
	const sigma = flt.Flt(6)
	const offset = flt.Flt(0.85)
	m := offset * flt.Flt(len(*x)-1)
	s := flt.Flt(len(*x)) / sigma
	ss2 := s * s * 2
	var wSum flt.Flt
	for i := 0; i < len(*x)-1; i++ {
		im := flt.Flt(i) - m
		w := flt.Flt(math.Exp(float64(-(im * im) / ss2)))
		alma += (*x)[len(*x)-1-i] * w
		wSum += w
	}
	return alma / wSum
}
func (x *Flts) ProAlma() (r flt.Flt) {
	if len(*x) == 0 {
		return 0
	}
	min, max := x.MinMax()
	rngFul := max - min
	if rngFul == 0 {
		return r
	}
	return (x.Alma() - min) / rngFul
}
func (x *CntrDistSeg) Act() {
	for m := x.Idx; m < x.Lim; m++ {
		for n := unt.Zero; n < x.Vals.Cnt(); n++ {
			if m != n && x.Vals.At(n).Neq(flt.Max) {
				if !x.EvalZero && x.Vals.At(n).Eql(flt.Zero) {
					continue
				}
				// out[m] += 1 - (abs(m-n)/rngFul)
				x.Out.Upd(m, x.Out.At(m)+flt.One.Sub(x.Vals.At(m).Sub(x.Vals.At(n)).Pos().Div(x.RngFul)))
			}
		}
	}
}
func (x *Flts) CntrDist(evalZero ...bol.Bol) (r *Flts) {
	r = Make(x.Cnt())
	rngFul := x.RngFul()
	segBnds, acts := bnds.Segs(x.Cnt())
	for n, segBnd := range *segBnds {
		acts[n] = &CntrDistSeg{
			Bnd:      segBnd,
			EvalZero: len(evalZero) != 0 && evalZero[0],
			RngFul:   rngFul,
			Vals:     x,
			Out:      r,
		}
	}
	sys.Run().Pll(acts...)
	return r
}
func (x *Flts) Float32s() (r []float32) {
	r = make([]float32, len(*x))
	for n := 0; n < len(*x); n++ {
		r[n] = float32((*x)[n])
	}
	return r
}
func (x *Flts) StrWrt(b *strings.Builder) {
	b.WriteRune('[')
	for n, v := range *x {
		if n != 0 {
			b.WriteRune(' ')
		}
		v.StrWrt(b)
	}
	b.WriteRune(']')
}
func (x *Flts) BytWrt(b *bytes.Buffer) {
	bLen := make([]byte, 4) // array length
	binary.LittleEndian.PutUint32(bLen, uint32(len(*x)))
	b.Write(bLen)
	for _, v := range *x {
		v.BytWrt(b)
	}
}
func (x *Flts) BytRed(b []byte) (idx int) {
	if len(b) >= 4 {
		*x = make(Flts, binary.LittleEndian.Uint32(b[:4])) // overwrite any previous existing
		idx = 4
		for n := 0; n < len(*x); n++ {
			(*x)[n].BytRed(b[idx : idx+flt.Size])
			idx += flt.Size
		}
	}
	return idx
}
func (x *Flts) Bytes() []byte {
	b := &bytes.Buffer{}
	x.BytWrt(b)
	return b.Bytes()
}
