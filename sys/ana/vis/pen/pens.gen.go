package pen

import (
	"math/rand"
	"sys/bsc/bnd"
	"sys/bsc/bol"
	"sys/bsc/unt"
	"time"
)

type (
	Pens    []Pen
	PensScp struct {
		Idx uint32
		Arr []*Pens
	}
)

func NewPens(vs ...Pen) *Pens {
	r := Pens(vs)
	return &r
}
func MakePens(cap unt.Unt) *Pens {
	r := make(Pens, cap)
	return &r
}
func MakeEmpPens(cap unt.Unt) *Pens {
	r := make(Pens, 0, cap)
	return &r
}
func (x *Pens) Ok() bol.Bol { return len(*x) != 0 }
func (x *Pens) Cnt() unt.Unt {
	if x == nil {
		return 0
	}
	return unt.Unt(len(*x))
}
func (x *Pens) Cpy() *Pens {
	r := make(Pens, len(*x))
	copy(r, *x)
	return &r
}
func (x *Pens) Clr() *Pens {
	*x = (*x)[:0]
	return x
}
func (x *Pens) Rand() *Pens {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	perm := r.Perm(len(*x))
	for i, randIdx := range perm {
		(*x)[i], (*x)[randIdx] = (*x)[randIdx], (*x)[i]
	}
	return x
}
func (x *Pens) Mrg(a ...*Pens) *Pens {
	for _, v := range a {
		*x = append(*x, *v...)
	}
	return x
}
func (x *Pens) Push(a ...Pen) *Pens {
	*x = append(*x, a...)
	return x
}
func (x *Pens) Pop() (r Pen) {
	if len(*x) == 0 {
		return r
	}
	r = (*x)[len(*x)-1]
	*x = (*x)[:len(*x)-1]
	return r
}
func (x *Pens) Que(vs ...Pen) *Pens {
	*x = append(*x, vs...)
	return x
}
func (x *Pens) Dque() (r Pen) {
	r = (*x)[0]
	if len(*x) == 1 {
		*x = (*x)[:0]
	} else {
		copy(*x, (*x)[1:])
		*x = (*x)[:len(*x)-1]
	}
	return r
}
func (x *Pens) Ins(idx unt.Unt, elm Pen) *Pens {
	*x = append((*x)[:idx], append(Pens{elm}, (*x)[idx:]...)...)
	return x
}
func (x *Pens) Upd(idx unt.Unt, elm Pen) *Pens {
	(*x)[idx] = elm
	return x
}
func (x *Pens) Del(idx unt.Unt) (r Pen) {
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
func (x *Pens) At(idx unt.Unt) Pen { return (*x)[idx] }
func (x *Pens) In(idx, lim unt.Unt) *Pens {
	r := (*x)[idx:lim]
	return &r
}
func (x *Pens) InBnd(b bnd.Bnd) *Pens {
	r := (*x)[b.Idx:b.Lim]
	return &r
}
func (x *Pens) From(idx unt.Unt) *Pens {
	var r Pens
	if idx < unt.Unt(len(*x)) {
		r = (*x)[idx:]
	} else {
		r = (*x)[:0]
	}
	return &r
}
func (x *Pens) To(lim unt.Unt) *Pens {
	if lim > unt.Unt(len(*x)) {
		return x
	}
	r := (*x)[:lim]
	return &r
}
func (x *Pens) Fst() Pen        { return (*x)[0] }
func (x *Pens) Mdl() Pen        { return (*x)[len(*x)/2] }
func (x *Pens) Lst() Pen        { return (*x)[len(*x)-1] }
func (x *Pens) FstIdx() unt.Unt { return 0 }
func (x *Pens) MdlIdx() unt.Unt { return unt.Unt(len(*x) / 2) }
func (x *Pens) LstIdx() unt.Unt { return unt.Unt(len(*x) - 1) }
func (x *Pens) Rev() *Pens {
	for i, j := 0, len(*x)-1; i < j; i, j = i+1, j-1 {
		(*x)[i], (*x)[j] = (*x)[j], (*x)[i]
	}
	return x
}
