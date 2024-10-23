package tme

import (
	"math/rand"
	"strings"
	"sys/bsc/bnd"
	"sys/bsc/bol"
	"sys/bsc/unt"
	"time"
)

type (
	Rngs    []Rng
	RngsScp struct {
		Idx uint32
		Arr []*Rngs
	}
)

func NewRngs(vs ...Rng) *Rngs {
	r := Rngs(vs)
	return &r
}
func MakeRngs(cap unt.Unt) *Rngs {
	r := make(Rngs, cap)
	return &r
}
func MakeEmpRngs(cap unt.Unt) *Rngs {
	r := make(Rngs, 0, cap)
	return &r
}
func (x *Rngs) Ok() bol.Bol { return len(*x) != 0 }
func (x *Rngs) Cnt() unt.Unt {
	if x == nil {
		return 0
	}
	return unt.Unt(len(*x))
}
func (x *Rngs) Cpy() *Rngs {
	r := make(Rngs, len(*x))
	copy(r, *x)
	return &r
}
func (x *Rngs) Clr() *Rngs {
	*x = (*x)[:0]
	return x
}
func (x *Rngs) Rand() *Rngs {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	perm := r.Perm(len(*x))
	for i, randIdx := range perm {
		(*x)[i], (*x)[randIdx] = (*x)[randIdx], (*x)[i]
	}
	return x
}
func (x *Rngs) Mrg(a ...*Rngs) *Rngs {
	for _, v := range a {
		*x = append(*x, *v...)
	}
	return x
}
func (x *Rngs) Push(a ...Rng) *Rngs {
	*x = append(*x, a...)
	return x
}
func (x *Rngs) Pop() (r Rng) {
	if len(*x) == 0 {
		return r
	}
	r = (*x)[len(*x)-1]
	*x = (*x)[:len(*x)-1]
	return r
}
func (x *Rngs) Que(vs ...Rng) *Rngs {
	*x = append(*x, vs...)
	return x
}
func (x *Rngs) Dque() (r Rng) {
	r = (*x)[0]
	if len(*x) == 1 {
		*x = (*x)[:0]
	} else {
		copy(*x, (*x)[1:])
		*x = (*x)[:len(*x)-1]
	}
	return r
}
func (x *Rngs) Ins(idx unt.Unt, elm Rng) *Rngs {
	*x = append((*x)[:idx], append(Rngs{elm}, (*x)[idx:]...)...)
	return x
}
func (x *Rngs) Upd(idx unt.Unt, elm Rng) *Rngs {
	(*x)[idx] = elm
	return x
}
func (x *Rngs) Del(idx unt.Unt) (r Rng) {
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
func (x *Rngs) At(idx unt.Unt) Rng { return (*x)[idx] }
func (x *Rngs) In(idx, lim unt.Unt) *Rngs {
	r := (*x)[idx:lim]
	return &r
}
func (x *Rngs) InBnd(b bnd.Bnd) *Rngs {
	r := (*x)[b.Idx:b.Lim]
	return &r
}
func (x *Rngs) From(idx unt.Unt) *Rngs {
	var r Rngs
	if idx < unt.Unt(len(*x)) {
		r = (*x)[idx:]
	} else {
		r = (*x)[:0]
	}
	return &r
}
func (x *Rngs) To(lim unt.Unt) *Rngs {
	if lim > unt.Unt(len(*x)) {
		return x
	}
	r := (*x)[:lim]
	return &r
}
func (x *Rngs) Fst() Rng        { return (*x)[0] }
func (x *Rngs) Mdl() Rng        { return (*x)[len(*x)/2] }
func (x *Rngs) Lst() Rng        { return (*x)[len(*x)-1] }
func (x *Rngs) FstIdx() unt.Unt { return 0 }
func (x *Rngs) MdlIdx() unt.Unt { return unt.Unt(len(*x) / 2) }
func (x *Rngs) LstIdx() unt.Unt { return unt.Unt(len(*x) - 1) }
func (x *Rngs) Rev() *Rngs {
	for i, j := 0, len(*x)-1; i < j; i, j = i+1, j-1 {
		(*x)[i], (*x)[j] = (*x)[j], (*x)[i]
	}
	return x
}
func (x *Rngs) SrchIdx(v Tme) unt.Unt {
	i, j := unt.Zero, unt.Unt(len(*x))
	for i < j {
		if (*x)[(i+j)>>1].Max < v {
			i = (i+j)>>1 + 1
		} else {
			j = (i + j) >> 1
		}
	}
	return i
}
func (x *Rngs) RngMrg(fstIdx, lstIdx unt.Unt) Rng { return x.At(fstIdx).Mrg(x.At(lstIdx)) }
func (x *Rngs) StrWrt(b *strings.Builder) {
	b.WriteRune('[')
	for n, v := range *x {
		if n != 0 {
			b.WriteRune(' ')
		}
		v.StrWrt(b)
	}
	b.WriteRune(']')
}
