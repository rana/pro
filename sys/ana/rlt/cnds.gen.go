package rlt

import (
	"math/rand"
	"sys/bsc/bnd"
	"sys/bsc/bol"
	"sys/bsc/unt"
	"time"
)

type (
	Cnds    []Cnd
	CndsScp struct {
		Idx uint32
		Arr []*Cnds
	}
)

func NewCnds(vs ...Cnd) *Cnds {
	r := Cnds(vs)
	return &r
}
func MakeCnds(cap unt.Unt) *Cnds {
	r := make(Cnds, cap)
	return &r
}
func MakeEmpCnds(cap unt.Unt) *Cnds {
	r := make(Cnds, 0, cap)
	return &r
}
func (x *Cnds) Ok() bol.Bol { return len(*x) != 0 }
func (x *Cnds) Cnt() unt.Unt {
	if x == nil {
		return 0
	}
	return unt.Unt(len(*x))
}
func (x *Cnds) Cpy() *Cnds {
	r := make(Cnds, len(*x))
	copy(r, *x)
	return &r
}
func (x *Cnds) Clr() *Cnds {
	*x = (*x)[:0]
	return x
}
func (x *Cnds) Rand() *Cnds {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	perm := r.Perm(len(*x))
	for i, randIdx := range perm {
		(*x)[i], (*x)[randIdx] = (*x)[randIdx], (*x)[i]
	}
	return x
}
func (x *Cnds) Mrg(a ...*Cnds) *Cnds {
	for _, v := range a {
		*x = append(*x, *v...)
	}
	return x
}
func (x *Cnds) Push(a ...Cnd) *Cnds {
	*x = append(*x, a...)
	return x
}
func (x *Cnds) Pop() (r Cnd) {
	if len(*x) == 0 {
		return r
	}
	r = (*x)[len(*x)-1]
	*x = (*x)[:len(*x)-1]
	return r
}
func (x *Cnds) Que(vs ...Cnd) *Cnds {
	*x = append(*x, vs...)
	return x
}
func (x *Cnds) Dque() (r Cnd) {
	r = (*x)[0]
	if len(*x) == 1 {
		*x = (*x)[:0]
	} else {
		copy(*x, (*x)[1:])
		*x = (*x)[:len(*x)-1]
	}
	return r
}
func (x *Cnds) Ins(idx unt.Unt, elm Cnd) *Cnds {
	*x = append((*x)[:idx], append(Cnds{elm}, (*x)[idx:]...)...)
	return x
}
func (x *Cnds) Upd(idx unt.Unt, elm Cnd) *Cnds {
	(*x)[idx] = elm
	return x
}
func (x *Cnds) Del(idx unt.Unt) (r Cnd) {
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
func (x *Cnds) At(idx unt.Unt) Cnd { return (*x)[idx] }
func (x *Cnds) In(idx, lim unt.Unt) *Cnds {
	r := (*x)[idx:lim]
	return &r
}
func (x *Cnds) InBnd(b bnd.Bnd) *Cnds {
	r := (*x)[b.Idx:b.Lim]
	return &r
}
func (x *Cnds) From(idx unt.Unt) *Cnds {
	var r Cnds
	if idx < unt.Unt(len(*x)) {
		r = (*x)[idx:]
	} else {
		r = (*x)[:0]
	}
	return &r
}
func (x *Cnds) To(lim unt.Unt) *Cnds {
	if lim > unt.Unt(len(*x)) {
		return x
	}
	r := (*x)[:lim]
	return &r
}
func (x *Cnds) Fst() Cnd        { return (*x)[0] }
func (x *Cnds) Mdl() Cnd        { return (*x)[len(*x)/2] }
func (x *Cnds) Lst() Cnd        { return (*x)[len(*x)-1] }
func (x *Cnds) FstIdx() unt.Unt { return 0 }
func (x *Cnds) MdlIdx() unt.Unt { return unt.Unt(len(*x) / 2) }
func (x *Cnds) LstIdx() unt.Unt { return unt.Unt(len(*x) - 1) }
func (x *Cnds) Rev() *Cnds {
	for i, j := 0, len(*x)-1; i < j; i, j = i+1, j-1 {
		(*x)[i], (*x)[j] = (*x)[j], (*x)[i]
	}
	return x
}
