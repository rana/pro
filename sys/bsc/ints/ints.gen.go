package ints

import (
	"bytes"
	"encoding/binary"
	"math/rand"
	"sort"
	"strings"
	"sys/bsc/bnd"
	"sys/bsc/bol"
	bscint "sys/bsc/int"
	"sys/bsc/unt"
	"time"
)

type (
	Ints    []bscint.Int
	IntsScp struct {
		Idx uint32
		Arr []*Ints
	}
)

func New(vs ...bscint.Int) *Ints {
	r := Ints(vs)
	return &r
}
func Make(cap unt.Unt) *Ints {
	r := make(Ints, cap)
	return &r
}
func MakeEmp(cap unt.Unt) *Ints {
	r := make(Ints, 0, cap)
	return &r
}
func (x *Ints) Ok() bol.Bol { return len(*x) != 0 }
func (x *Ints) Cnt() unt.Unt {
	if x == nil {
		return 0
	}
	return unt.Unt(len(*x))
}
func (x *Ints) Cpy() *Ints {
	r := make(Ints, len(*x))
	copy(r, *x)
	return &r
}
func (x *Ints) Clr() *Ints {
	*x = (*x)[:0]
	return x
}
func (x *Ints) Rand() *Ints {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	perm := r.Perm(len(*x))
	for i, randIdx := range perm {
		(*x)[i], (*x)[randIdx] = (*x)[randIdx], (*x)[i]
	}
	return x
}
func (x *Ints) Mrg(a ...*Ints) *Ints {
	for _, v := range a {
		*x = append(*x, *v...)
	}
	return x
}
func (x *Ints) Push(a ...bscint.Int) *Ints {
	*x = append(*x, a...)
	return x
}
func (x *Ints) Pop() (r bscint.Int) {
	if len(*x) == 0 {
		return r
	}
	r = (*x)[len(*x)-1]
	*x = (*x)[:len(*x)-1]
	return r
}
func (x *Ints) Que(vs ...bscint.Int) *Ints {
	*x = append(*x, vs...)
	return x
}
func (x *Ints) Dque() (r bscint.Int) {
	r = (*x)[0]
	if len(*x) == 1 {
		*x = (*x)[:0]
	} else {
		copy(*x, (*x)[1:])
		*x = (*x)[:len(*x)-1]
	}
	return r
}
func (x *Ints) Ins(idx unt.Unt, elm bscint.Int) *Ints {
	*x = append((*x)[:idx], append(Ints{elm}, (*x)[idx:]...)...)
	return x
}
func (x *Ints) Upd(idx unt.Unt, elm bscint.Int) *Ints {
	(*x)[idx] = elm
	return x
}
func (x *Ints) Del(idx unt.Unt) (r bscint.Int) {
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
func (x *Ints) At(idx unt.Unt) bscint.Int { return (*x)[idx] }
func (x *Ints) In(idx, lim unt.Unt) *Ints {
	r := (*x)[idx:lim]
	return &r
}
func (x *Ints) InBnd(b bnd.Bnd) *Ints {
	r := (*x)[b.Idx:b.Lim]
	return &r
}
func (x *Ints) From(idx unt.Unt) *Ints {
	var r Ints
	if idx < unt.Unt(len(*x)) {
		r = (*x)[idx:]
	} else {
		r = (*x)[:0]
	}
	return &r
}
func (x *Ints) To(lim unt.Unt) *Ints {
	if lim > unt.Unt(len(*x)) {
		return x
	}
	r := (*x)[:lim]
	return &r
}
func (x *Ints) Fst() bscint.Int { return (*x)[0] }
func (x *Ints) Mdl() bscint.Int { return (*x)[len(*x)/2] }
func (x *Ints) Lst() bscint.Int { return (*x)[len(*x)-1] }
func (x *Ints) FstIdx() unt.Unt { return 0 }
func (x *Ints) MdlIdx() unt.Unt { return unt.Unt(len(*x) / 2) }
func (x *Ints) LstIdx() unt.Unt { return unt.Unt(len(*x) - 1) }
func (x *Ints) Rev() *Ints {
	for i, j := 0, len(*x)-1; i < j; i, j = i+1, j-1 {
		(*x)[i], (*x)[j] = (*x)[j], (*x)[i]
	}
	return x
}
func (x *Ints) SrchIdxEql(v bscint.Int) unt.Unt {
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
func (x *Ints) SrchIdx(v bscint.Int, near ...bol.Bol) unt.Unt {
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
func (x *Ints) Has(v bscint.Int) bol.Bol {
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
func (x *Ints) SrtAsc() *Ints {
	if x.Cnt() > 1 {
		x.SrtQuick(0, x.LstIdx(), bscint.Lss, bscint.Eql)
	}
	return x
}
func (x *Ints) SrtDsc() *Ints {
	if x.Cnt() > 1 {
		x.SrtQuick(0, x.LstIdx(), bscint.Gtr, bscint.Eql)
	}
	return x
}
func (x *Ints) SrtQuick(lo, hi unt.Unt, cmp, eql bscint.Cmp) *Ints {
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
func (x *Ints) SrtIns(lo, hi unt.Unt, cmp bscint.Cmp) *Ints {
	for i := lo; i <= hi; i++ {
		for j := i; j > lo && cmp((*x)[j], (*x)[j-1]); j-- {
			x.Swp(j, j-1)
		}
	}
	return x
}
func (x *Ints) SrtMdnOf3(i, j, k unt.Unt, cmp bscint.Cmp) unt.Unt {
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
func (x *Ints) Swp(i, j unt.Unt) { (*x)[i], (*x)[j] = (*x)[j], (*x)[i] }
func (x *Ints) StrWrt(b *strings.Builder) {
	b.WriteRune('[')
	for n, v := range *x {
		if n != 0 {
			b.WriteRune(' ')
		}
		v.StrWrt(b)
	}
	b.WriteRune(']')
}
func (x *Ints) BytWrt(b *bytes.Buffer) {
	bLen := make([]byte, 4) // array length
	binary.LittleEndian.PutUint32(bLen, uint32(len(*x)))
	b.Write(bLen)
	for _, v := range *x {
		v.BytWrt(b)
	}
}
func (x *Ints) BytRed(b []byte) (idx int) {
	if len(b) >= 4 {
		*x = make(Ints, binary.LittleEndian.Uint32(b[:4])) // overwrite any previous existing
		idx = 4
		for n := 0; n < len(*x); n++ {
			(*x)[n].BytRed(b[idx : idx+bscint.Size])
			idx += bscint.Size
		}
	}
	return idx
}
