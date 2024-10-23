package rlt

import (
	"math/rand"
	"sys/bsc/bnd"
	"sys/bsc/bol"
	"sys/bsc/unt"
	"time"
)

type (
	Stgys    []Stgy
	StgysScp struct {
		Idx uint32
		Arr []*Stgys
	}
)

func NewStgys(vs ...Stgy) *Stgys {
	r := Stgys(vs)
	return &r
}
func MakeStgys(cap unt.Unt) *Stgys {
	r := make(Stgys, cap)
	return &r
}
func MakeEmpStgys(cap unt.Unt) *Stgys {
	r := make(Stgys, 0, cap)
	return &r
}
func (x *Stgys) Ok() bol.Bol { return len(*x) != 0 }
func (x *Stgys) Cnt() unt.Unt {
	if x == nil {
		return 0
	}
	return unt.Unt(len(*x))
}
func (x *Stgys) Cpy() *Stgys {
	r := make(Stgys, len(*x))
	copy(r, *x)
	return &r
}
func (x *Stgys) Clr() *Stgys {
	*x = (*x)[:0]
	return x
}
func (x *Stgys) Rand() *Stgys {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	perm := r.Perm(len(*x))
	for i, randIdx := range perm {
		(*x)[i], (*x)[randIdx] = (*x)[randIdx], (*x)[i]
	}
	return x
}
func (x *Stgys) Mrg(a ...*Stgys) *Stgys {
	for _, v := range a {
		*x = append(*x, *v...)
	}
	return x
}
func (x *Stgys) Push(a ...Stgy) *Stgys {
	*x = append(*x, a...)
	return x
}
func (x *Stgys) Pop() (r Stgy) {
	if len(*x) == 0 {
		return r
	}
	r = (*x)[len(*x)-1]
	*x = (*x)[:len(*x)-1]
	return r
}
func (x *Stgys) Que(vs ...Stgy) *Stgys {
	*x = append(*x, vs...)
	return x
}
func (x *Stgys) Dque() (r Stgy) {
	r = (*x)[0]
	if len(*x) == 1 {
		*x = (*x)[:0]
	} else {
		copy(*x, (*x)[1:])
		*x = (*x)[:len(*x)-1]
	}
	return r
}
func (x *Stgys) Ins(idx unt.Unt, elm Stgy) *Stgys {
	*x = append((*x)[:idx], append(Stgys{elm}, (*x)[idx:]...)...)
	return x
}
func (x *Stgys) Upd(idx unt.Unt, elm Stgy) *Stgys {
	(*x)[idx] = elm
	return x
}
func (x *Stgys) Del(idx unt.Unt) (r Stgy) {
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
func (x *Stgys) At(idx unt.Unt) Stgy { return (*x)[idx] }
func (x *Stgys) In(idx, lim unt.Unt) *Stgys {
	r := (*x)[idx:lim]
	return &r
}
func (x *Stgys) InBnd(b bnd.Bnd) *Stgys {
	r := (*x)[b.Idx:b.Lim]
	return &r
}
func (x *Stgys) From(idx unt.Unt) *Stgys {
	var r Stgys
	if idx < unt.Unt(len(*x)) {
		r = (*x)[idx:]
	} else {
		r = (*x)[:0]
	}
	return &r
}
func (x *Stgys) To(lim unt.Unt) *Stgys {
	if lim > unt.Unt(len(*x)) {
		return x
	}
	r := (*x)[:lim]
	return &r
}
func (x *Stgys) Fst() Stgy       { return (*x)[0] }
func (x *Stgys) Mdl() Stgy       { return (*x)[len(*x)/2] }
func (x *Stgys) Lst() Stgy       { return (*x)[len(*x)-1] }
func (x *Stgys) FstIdx() unt.Unt { return 0 }
func (x *Stgys) MdlIdx() unt.Unt { return unt.Unt(len(*x) / 2) }
func (x *Stgys) LstIdx() unt.Unt { return unt.Unt(len(*x) - 1) }
func (x *Stgys) Rev() *Stgys {
	for i, j := 0, len(*x)-1; i < j; i, j = i+1, j-1 {
		(*x)[i], (*x)[j] = (*x)[j], (*x)[i]
	}
	return x
}
