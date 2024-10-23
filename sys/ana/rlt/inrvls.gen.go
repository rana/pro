package rlt

import (
	"math/rand"
	"sys/bsc/bnd"
	"sys/bsc/bol"
	"sys/bsc/unt"
	"time"
)

type (
	Inrvls    []Inrvl
	InrvlsScp struct {
		Idx uint32
		Arr []*Inrvls
	}
)

func NewInrvls(vs ...Inrvl) *Inrvls {
	r := Inrvls(vs)
	return &r
}
func MakeInrvls(cap unt.Unt) *Inrvls {
	r := make(Inrvls, cap)
	return &r
}
func MakeEmpInrvls(cap unt.Unt) *Inrvls {
	r := make(Inrvls, 0, cap)
	return &r
}
func (x *Inrvls) Ok() bol.Bol { return len(*x) != 0 }
func (x *Inrvls) Cnt() unt.Unt {
	if x == nil {
		return 0
	}
	return unt.Unt(len(*x))
}
func (x *Inrvls) Cpy() *Inrvls {
	r := make(Inrvls, len(*x))
	copy(r, *x)
	return &r
}
func (x *Inrvls) Clr() *Inrvls {
	*x = (*x)[:0]
	return x
}
func (x *Inrvls) Rand() *Inrvls {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	perm := r.Perm(len(*x))
	for i, randIdx := range perm {
		(*x)[i], (*x)[randIdx] = (*x)[randIdx], (*x)[i]
	}
	return x
}
func (x *Inrvls) Mrg(a ...*Inrvls) *Inrvls {
	for _, v := range a {
		*x = append(*x, *v...)
	}
	return x
}
func (x *Inrvls) Push(a ...Inrvl) *Inrvls {
	*x = append(*x, a...)
	return x
}
func (x *Inrvls) Pop() (r Inrvl) {
	if len(*x) == 0 {
		return r
	}
	r = (*x)[len(*x)-1]
	*x = (*x)[:len(*x)-1]
	return r
}
func (x *Inrvls) Que(vs ...Inrvl) *Inrvls {
	*x = append(*x, vs...)
	return x
}
func (x *Inrvls) Dque() (r Inrvl) {
	r = (*x)[0]
	if len(*x) == 1 {
		*x = (*x)[:0]
	} else {
		copy(*x, (*x)[1:])
		*x = (*x)[:len(*x)-1]
	}
	return r
}
func (x *Inrvls) Ins(idx unt.Unt, elm Inrvl) *Inrvls {
	*x = append((*x)[:idx], append(Inrvls{elm}, (*x)[idx:]...)...)
	return x
}
func (x *Inrvls) Upd(idx unt.Unt, elm Inrvl) *Inrvls {
	(*x)[idx] = elm
	return x
}
func (x *Inrvls) Del(idx unt.Unt) (r Inrvl) {
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
func (x *Inrvls) At(idx unt.Unt) Inrvl { return (*x)[idx] }
func (x *Inrvls) In(idx, lim unt.Unt) *Inrvls {
	r := (*x)[idx:lim]
	return &r
}
func (x *Inrvls) InBnd(b bnd.Bnd) *Inrvls {
	r := (*x)[b.Idx:b.Lim]
	return &r
}
func (x *Inrvls) From(idx unt.Unt) *Inrvls {
	var r Inrvls
	if idx < unt.Unt(len(*x)) {
		r = (*x)[idx:]
	} else {
		r = (*x)[:0]
	}
	return &r
}
func (x *Inrvls) To(lim unt.Unt) *Inrvls {
	if lim > unt.Unt(len(*x)) {
		return x
	}
	r := (*x)[:lim]
	return &r
}
func (x *Inrvls) Fst() Inrvl      { return (*x)[0] }
func (x *Inrvls) Mdl() Inrvl      { return (*x)[len(*x)/2] }
func (x *Inrvls) Lst() Inrvl      { return (*x)[len(*x)-1] }
func (x *Inrvls) FstIdx() unt.Unt { return 0 }
func (x *Inrvls) MdlIdx() unt.Unt { return unt.Unt(len(*x) / 2) }
func (x *Inrvls) LstIdx() unt.Unt { return unt.Unt(len(*x) - 1) }
func (x *Inrvls) Rev() *Inrvls {
	for i, j := 0, len(*x)-1; i < j; i, j = i+1, j-1 {
		(*x)[i], (*x)[j] = (*x)[j], (*x)[i]
	}
	return x
}
