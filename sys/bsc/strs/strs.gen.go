package strs

import (
	"bytes"
	"encoding/binary"
	"math/rand"
	"sort"
	"strings"
	"sys/bsc/bnd"
	"sys/bsc/bol"
	"sys/bsc/str"
	"sys/bsc/unt"
	"time"
)

type (
	Strs    []str.Str
	StrsScp struct {
		Idx uint32
		Arr []*Strs
	}
)

func New(vs ...str.Str) *Strs {
	r := Strs(vs)
	return &r
}
func Make(cap unt.Unt) *Strs {
	r := make(Strs, cap)
	return &r
}
func MakeEmp(cap unt.Unt) *Strs {
	r := make(Strs, 0, cap)
	return &r
}
func (x *Strs) BytWrt(b *bytes.Buffer) {
	bLen := make([]byte, 4) // array length
	binary.LittleEndian.PutUint32(bLen, uint32(len(*x)))
	b.Write(bLen)
	for _, v := range *x {
		binary.LittleEndian.PutUint32(bLen, uint32(len(v)))
		b.Write(bLen)            // string length
		b.WriteString(string(v)) // current string
	}
}
func (x *Strs) BytRed(b []byte) (r int) {
	if len(b) >= 4 {
		cnt := int(binary.LittleEndian.Uint32(b[:4]))
		idx := 4
		for n := 0; n < cnt; n++ {
			vLen := int(binary.LittleEndian.Uint32(b[idx : idx+4]))
			*x = append(*x, str.Str(string(b[idx+4:idx+4+vLen])))
			idx += 4 + vLen
			r += vLen
		}
	}
	return r
}
func (x *Strs) Ok() bol.Bol { return len(*x) != 0 }
func (x *Strs) Cnt() unt.Unt {
	if x == nil {
		return 0
	}
	return unt.Unt(len(*x))
}
func (x *Strs) Cpy() *Strs {
	r := make(Strs, len(*x))
	copy(r, *x)
	return &r
}
func (x *Strs) Clr() *Strs {
	*x = (*x)[:0]
	return x
}
func (x *Strs) Rand() *Strs {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	perm := r.Perm(len(*x))
	for i, randIdx := range perm {
		(*x)[i], (*x)[randIdx] = (*x)[randIdx], (*x)[i]
	}
	return x
}
func (x *Strs) Mrg(a ...*Strs) *Strs {
	for _, v := range a {
		*x = append(*x, *v...)
	}
	return x
}
func (x *Strs) Push(a ...str.Str) *Strs {
	*x = append(*x, a...)
	return x
}
func (x *Strs) Pop() (r str.Str) {
	if len(*x) == 0 {
		return r
	}
	r = (*x)[len(*x)-1]
	*x = (*x)[:len(*x)-1]
	return r
}
func (x *Strs) Que(vs ...str.Str) *Strs {
	*x = append(*x, vs...)
	return x
}
func (x *Strs) Dque() (r str.Str) {
	r = (*x)[0]
	if len(*x) == 1 {
		*x = (*x)[:0]
	} else {
		copy(*x, (*x)[1:])
		*x = (*x)[:len(*x)-1]
	}
	return r
}
func (x *Strs) Ins(idx unt.Unt, elm str.Str) *Strs {
	*x = append((*x)[:idx], append(Strs{elm}, (*x)[idx:]...)...)
	return x
}
func (x *Strs) Upd(idx unt.Unt, elm str.Str) *Strs {
	(*x)[idx] = elm
	return x
}
func (x *Strs) Del(idx unt.Unt) (r str.Str) {
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
func (x *Strs) At(idx unt.Unt) str.Str { return (*x)[idx] }
func (x *Strs) In(idx, lim unt.Unt) *Strs {
	r := (*x)[idx:lim]
	return &r
}
func (x *Strs) InBnd(b bnd.Bnd) *Strs {
	r := (*x)[b.Idx:b.Lim]
	return &r
}
func (x *Strs) From(idx unt.Unt) *Strs {
	var r Strs
	if idx < unt.Unt(len(*x)) {
		r = (*x)[idx:]
	} else {
		r = (*x)[:0]
	}
	return &r
}
func (x *Strs) To(lim unt.Unt) *Strs {
	if lim > unt.Unt(len(*x)) {
		return x
	}
	r := (*x)[:lim]
	return &r
}
func (x *Strs) Fst() str.Str    { return (*x)[0] }
func (x *Strs) Mdl() str.Str    { return (*x)[len(*x)/2] }
func (x *Strs) Lst() str.Str    { return (*x)[len(*x)-1] }
func (x *Strs) FstIdx() unt.Unt { return 0 }
func (x *Strs) MdlIdx() unt.Unt { return unt.Unt(len(*x) / 2) }
func (x *Strs) LstIdx() unt.Unt { return unt.Unt(len(*x) - 1) }
func (x *Strs) Rev() *Strs {
	for i, j := 0, len(*x)-1; i < j; i, j = i+1, j-1 {
		(*x)[i], (*x)[j] = (*x)[j], (*x)[i]
	}
	return x
}
func (x *Strs) SrchIdxEql(v str.Str) unt.Unt {
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
func (x *Strs) SrchIdx(v str.Str, near ...bol.Bol) unt.Unt {
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
func (x *Strs) Has(v str.Str) bol.Bol {
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
func (x *Strs) SrtAsc() *Strs {
	if x.Cnt() > 1 {
		x.SrtQuick(0, x.LstIdx(), str.Lss, str.Eql)
	}
	return x
}
func (x *Strs) SrtDsc() *Strs {
	if x.Cnt() > 1 {
		x.SrtQuick(0, x.LstIdx(), str.Gtr, str.Eql)
	}
	return x
}
func (x *Strs) SrtQuick(lo, hi unt.Unt, cmp, eql str.Cmp) *Strs {
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
func (x *Strs) SrtIns(lo, hi unt.Unt, cmp str.Cmp) *Strs {
	for i := lo; i <= hi; i++ {
		for j := i; j > lo && cmp((*x)[j], (*x)[j-1]); j-- {
			x.Swp(j, j-1)
		}
	}
	return x
}
func (x *Strs) SrtMdnOf3(i, j, k unt.Unt, cmp str.Cmp) unt.Unt {
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
func (x *Strs) Swp(i, j unt.Unt) { (*x)[i], (*x)[j] = (*x)[j], (*x)[i] }
func (x *Strs) StrWrt(b *strings.Builder) {
	b.WriteRune('[')
	for n, v := range *x {
		if n != 0 {
			b.WriteRune(' ')
		}
		v.StrWrt(b)
	}
	b.WriteRune(']')
}
