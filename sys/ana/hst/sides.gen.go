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
	Sides    []Side
	SidesScp struct {
		Idx uint32
		Arr []*Sides
	}
)

func NewSides(vs ...Side) *Sides {
	r := Sides(vs)
	return &r
}
func MakeSides(cap unt.Unt) *Sides {
	r := make(Sides, cap)
	return &r
}
func MakeEmpSides(cap unt.Unt) *Sides {
	r := make(Sides, 0, cap)
	return &r
}
func (x *Sides) Ok() bol.Bol { return len(*x) != 0 }
func (x *Sides) Cnt() unt.Unt {
	if x == nil {
		return 0
	}
	return unt.Unt(len(*x))
}
func (x *Sides) Cpy() *Sides {
	r := make(Sides, len(*x))
	copy(r, *x)
	return &r
}
func (x *Sides) Clr() *Sides {
	*x = (*x)[:0]
	return x
}
func (x *Sides) Rand() *Sides {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	perm := r.Perm(len(*x))
	for i, randIdx := range perm {
		(*x)[i], (*x)[randIdx] = (*x)[randIdx], (*x)[i]
	}
	return x
}
func (x *Sides) Mrg(a ...*Sides) *Sides {
	for _, v := range a {
		*x = append(*x, *v...)
	}
	return x
}
func (x *Sides) Push(a ...Side) *Sides {
	*x = append(*x, a...)
	return x
}
func (x *Sides) Pop() (r Side) {
	if len(*x) == 0 {
		return r
	}
	r = (*x)[len(*x)-1]
	*x = (*x)[:len(*x)-1]
	return r
}
func (x *Sides) Que(vs ...Side) *Sides {
	*x = append(*x, vs...)
	return x
}
func (x *Sides) Dque() (r Side) {
	r = (*x)[0]
	if len(*x) == 1 {
		*x = (*x)[:0]
	} else {
		copy(*x, (*x)[1:])
		*x = (*x)[:len(*x)-1]
	}
	return r
}
func (x *Sides) Ins(idx unt.Unt, elm Side) *Sides {
	*x = append((*x)[:idx], append(Sides{elm}, (*x)[idx:]...)...)
	return x
}
func (x *Sides) Upd(idx unt.Unt, elm Side) *Sides {
	(*x)[idx] = elm
	return x
}
func (x *Sides) Del(idx unt.Unt) (r Side) {
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
func (x *Sides) At(idx unt.Unt) Side { return (*x)[idx] }
func (x *Sides) In(idx, lim unt.Unt) *Sides {
	r := (*x)[idx:lim]
	return &r
}
func (x *Sides) InBnd(b bnd.Bnd) *Sides {
	r := (*x)[b.Idx:b.Lim]
	return &r
}
func (x *Sides) From(idx unt.Unt) *Sides {
	var r Sides
	if idx < unt.Unt(len(*x)) {
		r = (*x)[idx:]
	} else {
		r = (*x)[:0]
	}
	return &r
}
func (x *Sides) To(lim unt.Unt) *Sides {
	if lim > unt.Unt(len(*x)) {
		return x
	}
	r := (*x)[:lim]
	return &r
}
func (x *Sides) Fst() Side       { return (*x)[0] }
func (x *Sides) Mdl() Side       { return (*x)[len(*x)/2] }
func (x *Sides) Lst() Side       { return (*x)[len(*x)-1] }
func (x *Sides) FstIdx() unt.Unt { return 0 }
func (x *Sides) MdlIdx() unt.Unt { return unt.Unt(len(*x) / 2) }
func (x *Sides) LstIdx() unt.Unt { return unt.Unt(len(*x) - 1) }
func (x *Sides) Rev() *Sides {
	for i, j := 0, len(*x)-1; i < j; i, j = i+1, j-1 {
		(*x)[i], (*x)[j] = (*x)[j], (*x)[i]
	}
	return x
}
func (x *Sides) StrWrt(b *strings.Builder) {
	b.WriteRune('[')
	for n, v := range *x {
		if n != 0 {
			b.WriteRune(' ')
		}
		v.StrWrt(b)
	}
	b.WriteRune(']')
}
func (x *Sides) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
