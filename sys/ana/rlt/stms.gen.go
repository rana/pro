package rlt

import (
	"math/rand"
	"strings"
	"sys/bsc/bnd"
	"sys/bsc/bol"
	"sys/bsc/unt"
	"time"
)

type (
	Stms    []Stm
	StmsScp struct {
		Idx uint32
		Arr []*Stms
	}
)

func NewStms(vs ...Stm) *Stms {
	r := Stms(vs)
	return &r
}
func MakeStms(cap unt.Unt) *Stms {
	r := make(Stms, cap)
	return &r
}
func MakeEmpStms(cap unt.Unt) *Stms {
	r := make(Stms, 0, cap)
	return &r
}
func (x *Stms) Ok() bol.Bol { return len(*x) != 0 }
func (x *Stms) Cnt() unt.Unt {
	if x == nil {
		return 0
	}
	return unt.Unt(len(*x))
}
func (x *Stms) Cpy() *Stms {
	r := make(Stms, len(*x))
	copy(r, *x)
	return &r
}
func (x *Stms) Clr() *Stms {
	*x = (*x)[:0]
	return x
}
func (x *Stms) Rand() *Stms {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	perm := r.Perm(len(*x))
	for i, randIdx := range perm {
		(*x)[i], (*x)[randIdx] = (*x)[randIdx], (*x)[i]
	}
	return x
}
func (x *Stms) Mrg(a ...*Stms) *Stms {
	for _, v := range a {
		*x = append(*x, *v...)
	}
	return x
}
func (x *Stms) Push(a ...Stm) *Stms {
	*x = append(*x, a...)
	return x
}
func (x *Stms) Pop() (r Stm) {
	if len(*x) == 0 {
		return r
	}
	r = (*x)[len(*x)-1]
	*x = (*x)[:len(*x)-1]
	return r
}
func (x *Stms) Que(vs ...Stm) *Stms {
	*x = append(*x, vs...)
	return x
}
func (x *Stms) Dque() (r Stm) {
	r = (*x)[0]
	if len(*x) == 1 {
		*x = (*x)[:0]
	} else {
		copy(*x, (*x)[1:])
		*x = (*x)[:len(*x)-1]
	}
	return r
}
func (x *Stms) Ins(idx unt.Unt, elm Stm) *Stms {
	*x = append((*x)[:idx], append(Stms{elm}, (*x)[idx:]...)...)
	return x
}
func (x *Stms) Upd(idx unt.Unt, elm Stm) *Stms {
	(*x)[idx] = elm
	return x
}
func (x *Stms) Del(idx unt.Unt) (r Stm) {
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
func (x *Stms) At(idx unt.Unt) Stm { return (*x)[idx] }
func (x *Stms) In(idx, lim unt.Unt) *Stms {
	r := (*x)[idx:lim]
	return &r
}
func (x *Stms) InBnd(b bnd.Bnd) *Stms {
	r := (*x)[b.Idx:b.Lim]
	return &r
}
func (x *Stms) From(idx unt.Unt) *Stms {
	var r Stms
	if idx < unt.Unt(len(*x)) {
		r = (*x)[idx:]
	} else {
		r = (*x)[:0]
	}
	return &r
}
func (x *Stms) To(lim unt.Unt) *Stms {
	if lim > unt.Unt(len(*x)) {
		return x
	}
	r := (*x)[:lim]
	return &r
}
func (x *Stms) Fst() Stm        { return (*x)[0] }
func (x *Stms) Mdl() Stm        { return (*x)[len(*x)/2] }
func (x *Stms) Lst() Stm        { return (*x)[len(*x)-1] }
func (x *Stms) FstIdx() unt.Unt { return 0 }
func (x *Stms) MdlIdx() unt.Unt { return unt.Unt(len(*x) / 2) }
func (x *Stms) LstIdx() unt.Unt { return unt.Unt(len(*x) - 1) }
func (x *Stms) Rev() *Stms {
	for i, j := 0, len(*x)-1; i < j; i, j = i+1, j-1 {
		(*x)[i], (*x)[j] = (*x)[j], (*x)[i]
	}
	return x
}
func (x *Stms) StrWrt(b *strings.Builder) {
	b.WriteRune('[')
	for n, v := range *x {
		if n != 0 {
			b.WriteRune(' ')
		}
		v.StrWrt(b)
	}
	b.WriteRune(']')
}
func (x *Stms) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
