package hst

import (
	"math/rand"
	"strings"
	"sys/bsc/bnd"
	"sys/bsc/bol"
	"sys/bsc/unt"
	"time"
)

type (
	Prvs    []Prv
	PrvsScp struct {
		Idx uint32
		Arr []*Prvs
	}
)

func NewPrvs(vs ...Prv) *Prvs {
	r := Prvs(vs)
	return &r
}
func MakePrvs(cap unt.Unt) *Prvs {
	r := make(Prvs, cap)
	return &r
}
func MakeEmpPrvs(cap unt.Unt) *Prvs {
	r := make(Prvs, 0, cap)
	return &r
}
func (x *Prvs) Ok() bol.Bol { return len(*x) != 0 }
func (x *Prvs) Cnt() unt.Unt {
	if x == nil {
		return 0
	}
	return unt.Unt(len(*x))
}
func (x *Prvs) Cpy() *Prvs {
	r := make(Prvs, len(*x))
	copy(r, *x)
	return &r
}
func (x *Prvs) Clr() *Prvs {
	*x = (*x)[:0]
	return x
}
func (x *Prvs) Rand() *Prvs {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	perm := r.Perm(len(*x))
	for i, randIdx := range perm {
		(*x)[i], (*x)[randIdx] = (*x)[randIdx], (*x)[i]
	}
	return x
}
func (x *Prvs) Mrg(a ...*Prvs) *Prvs {
	for _, v := range a {
		*x = append(*x, *v...)
	}
	return x
}
func (x *Prvs) Push(a ...Prv) *Prvs {
	*x = append(*x, a...)
	return x
}
func (x *Prvs) Pop() (r Prv) {
	if len(*x) == 0 {
		return r
	}
	r = (*x)[len(*x)-1]
	*x = (*x)[:len(*x)-1]
	return r
}
func (x *Prvs) Que(vs ...Prv) *Prvs {
	*x = append(*x, vs...)
	return x
}
func (x *Prvs) Dque() (r Prv) {
	r = (*x)[0]
	if len(*x) == 1 {
		*x = (*x)[:0]
	} else {
		copy(*x, (*x)[1:])
		*x = (*x)[:len(*x)-1]
	}
	return r
}
func (x *Prvs) Ins(idx unt.Unt, elm Prv) *Prvs {
	*x = append((*x)[:idx], append(Prvs{elm}, (*x)[idx:]...)...)
	return x
}
func (x *Prvs) Upd(idx unt.Unt, elm Prv) *Prvs {
	(*x)[idx] = elm
	return x
}
func (x *Prvs) Del(idx unt.Unt) (r Prv) {
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
func (x *Prvs) At(idx unt.Unt) Prv { return (*x)[idx] }
func (x *Prvs) In(idx, lim unt.Unt) *Prvs {
	r := (*x)[idx:lim]
	return &r
}
func (x *Prvs) InBnd(b bnd.Bnd) *Prvs {
	r := (*x)[b.Idx:b.Lim]
	return &r
}
func (x *Prvs) From(idx unt.Unt) *Prvs {
	var r Prvs
	if idx < unt.Unt(len(*x)) {
		r = (*x)[idx:]
	} else {
		r = (*x)[:0]
	}
	return &r
}
func (x *Prvs) To(lim unt.Unt) *Prvs {
	if lim > unt.Unt(len(*x)) {
		return x
	}
	r := (*x)[:lim]
	return &r
}
func (x *Prvs) Fst() Prv        { return (*x)[0] }
func (x *Prvs) Mdl() Prv        { return (*x)[len(*x)/2] }
func (x *Prvs) Lst() Prv        { return (*x)[len(*x)-1] }
func (x *Prvs) FstIdx() unt.Unt { return 0 }
func (x *Prvs) MdlIdx() unt.Unt { return unt.Unt(len(*x) / 2) }
func (x *Prvs) LstIdx() unt.Unt { return unt.Unt(len(*x) - 1) }
func (x *Prvs) Rev() *Prvs {
	for i, j := 0, len(*x)-1; i < j; i, j = i+1, j-1 {
		(*x)[i], (*x)[j] = (*x)[j], (*x)[i]
	}
	return x
}
func (x *Prvs) StrWrt(b *strings.Builder) {
	b.WriteRune('[')
	for n, v := range *x {
		if n != 0 {
			b.WriteRune(' ')
		}
		v.StrWrt(b)
	}
	b.WriteRune(']')
}
func (x *Prvs) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
