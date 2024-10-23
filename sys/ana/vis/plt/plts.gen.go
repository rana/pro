package plt

import (
	"math/rand"
	"sys/bsc/bnd"
	"sys/bsc/bol"
	"sys/bsc/unt"
	"time"
)

type (
	Plts    []Plt
	PltsScp struct {
		Idx uint32
		Arr []*Plts
	}
)

func NewPlts(vs ...Plt) *Plts {
	r := Plts(vs)
	return &r
}
func MakePlts(cap unt.Unt) *Plts {
	r := make(Plts, cap)
	return &r
}
func MakeEmpPlts(cap unt.Unt) *Plts {
	r := make(Plts, 0, cap)
	return &r
}
func (x *Plts) Ok() bol.Bol { return len(*x) != 0 }
func (x *Plts) Cnt() unt.Unt {
	if x == nil {
		return 0
	}
	return unt.Unt(len(*x))
}
func (x *Plts) Cpy() *Plts {
	r := make(Plts, len(*x))
	copy(r, *x)
	return &r
}
func (x *Plts) Clr() *Plts {
	*x = (*x)[:0]
	return x
}
func (x *Plts) Rand() *Plts {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	perm := r.Perm(len(*x))
	for i, randIdx := range perm {
		(*x)[i], (*x)[randIdx] = (*x)[randIdx], (*x)[i]
	}
	return x
}
func (x *Plts) Mrg(a ...*Plts) *Plts {
	for _, v := range a {
		*x = append(*x, *v...)
	}
	return x
}
func (x *Plts) Push(a ...Plt) *Plts {
	*x = append(*x, a...)
	return x
}
func (x *Plts) Pop() (r Plt) {
	if len(*x) == 0 {
		return r
	}
	r = (*x)[len(*x)-1]
	*x = (*x)[:len(*x)-1]
	return r
}
func (x *Plts) Que(vs ...Plt) *Plts {
	*x = append(*x, vs...)
	return x
}
func (x *Plts) Dque() (r Plt) {
	r = (*x)[0]
	if len(*x) == 1 {
		*x = (*x)[:0]
	} else {
		copy(*x, (*x)[1:])
		*x = (*x)[:len(*x)-1]
	}
	return r
}
func (x *Plts) Ins(idx unt.Unt, elm Plt) *Plts {
	*x = append((*x)[:idx], append(Plts{elm}, (*x)[idx:]...)...)
	return x
}
func (x *Plts) Upd(idx unt.Unt, elm Plt) *Plts {
	(*x)[idx] = elm
	return x
}
func (x *Plts) Del(idx unt.Unt) (r Plt) {
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
func (x *Plts) At(idx unt.Unt) Plt { return (*x)[idx] }
func (x *Plts) In(idx, lim unt.Unt) *Plts {
	r := (*x)[idx:lim]
	return &r
}
func (x *Plts) InBnd(b bnd.Bnd) *Plts {
	r := (*x)[b.Idx:b.Lim]
	return &r
}
func (x *Plts) From(idx unt.Unt) *Plts {
	var r Plts
	if idx < unt.Unt(len(*x)) {
		r = (*x)[idx:]
	} else {
		r = (*x)[:0]
	}
	return &r
}
func (x *Plts) To(lim unt.Unt) *Plts {
	if lim > unt.Unt(len(*x)) {
		return x
	}
	r := (*x)[:lim]
	return &r
}
func (x *Plts) Fst() Plt        { return (*x)[0] }
func (x *Plts) Mdl() Plt        { return (*x)[len(*x)/2] }
func (x *Plts) Lst() Plt        { return (*x)[len(*x)-1] }
func (x *Plts) FstIdx() unt.Unt { return 0 }
func (x *Plts) MdlIdx() unt.Unt { return unt.Unt(len(*x) / 2) }
func (x *Plts) LstIdx() unt.Unt { return unt.Unt(len(*x) - 1) }
func (x *Plts) Rev() *Plts {
	for i, j := 0, len(*x)-1; i < j; i, j = i+1, j-1 {
		(*x)[i], (*x)[j] = (*x)[j], (*x)[i]
	}
	return x
}
