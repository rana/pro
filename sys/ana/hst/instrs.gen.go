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
	Instrs    []Instr
	InstrsScp struct {
		Idx uint32
		Arr []*Instrs
	}
)

func NewInstrs(vs ...Instr) *Instrs {
	r := Instrs(vs)
	return &r
}
func MakeInstrs(cap unt.Unt) *Instrs {
	r := make(Instrs, cap)
	return &r
}
func MakeEmpInstrs(cap unt.Unt) *Instrs {
	r := make(Instrs, 0, cap)
	return &r
}
func (x *Instrs) Ok() bol.Bol { return len(*x) != 0 }
func (x *Instrs) Cnt() unt.Unt {
	if x == nil {
		return 0
	}
	return unt.Unt(len(*x))
}
func (x *Instrs) Cpy() *Instrs {
	r := make(Instrs, len(*x))
	copy(r, *x)
	return &r
}
func (x *Instrs) Clr() *Instrs {
	*x = (*x)[:0]
	return x
}
func (x *Instrs) Rand() *Instrs {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	perm := r.Perm(len(*x))
	for i, randIdx := range perm {
		(*x)[i], (*x)[randIdx] = (*x)[randIdx], (*x)[i]
	}
	return x
}
func (x *Instrs) Mrg(a ...*Instrs) *Instrs {
	for _, v := range a {
		*x = append(*x, *v...)
	}
	return x
}
func (x *Instrs) Push(a ...Instr) *Instrs {
	*x = append(*x, a...)
	return x
}
func (x *Instrs) Pop() (r Instr) {
	if len(*x) == 0 {
		return r
	}
	r = (*x)[len(*x)-1]
	*x = (*x)[:len(*x)-1]
	return r
}
func (x *Instrs) Que(vs ...Instr) *Instrs {
	*x = append(*x, vs...)
	return x
}
func (x *Instrs) Dque() (r Instr) {
	r = (*x)[0]
	if len(*x) == 1 {
		*x = (*x)[:0]
	} else {
		copy(*x, (*x)[1:])
		*x = (*x)[:len(*x)-1]
	}
	return r
}
func (x *Instrs) Ins(idx unt.Unt, elm Instr) *Instrs {
	*x = append((*x)[:idx], append(Instrs{elm}, (*x)[idx:]...)...)
	return x
}
func (x *Instrs) Upd(idx unt.Unt, elm Instr) *Instrs {
	(*x)[idx] = elm
	return x
}
func (x *Instrs) Del(idx unt.Unt) (r Instr) {
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
func (x *Instrs) At(idx unt.Unt) Instr { return (*x)[idx] }
func (x *Instrs) In(idx, lim unt.Unt) *Instrs {
	r := (*x)[idx:lim]
	return &r
}
func (x *Instrs) InBnd(b bnd.Bnd) *Instrs {
	r := (*x)[b.Idx:b.Lim]
	return &r
}
func (x *Instrs) From(idx unt.Unt) *Instrs {
	var r Instrs
	if idx < unt.Unt(len(*x)) {
		r = (*x)[idx:]
	} else {
		r = (*x)[:0]
	}
	return &r
}
func (x *Instrs) To(lim unt.Unt) *Instrs {
	if lim > unt.Unt(len(*x)) {
		return x
	}
	r := (*x)[:lim]
	return &r
}
func (x *Instrs) Fst() Instr      { return (*x)[0] }
func (x *Instrs) Mdl() Instr      { return (*x)[len(*x)/2] }
func (x *Instrs) Lst() Instr      { return (*x)[len(*x)-1] }
func (x *Instrs) FstIdx() unt.Unt { return 0 }
func (x *Instrs) MdlIdx() unt.Unt { return unt.Unt(len(*x) / 2) }
func (x *Instrs) LstIdx() unt.Unt { return unt.Unt(len(*x) - 1) }
func (x *Instrs) Rev() *Instrs {
	for i, j := 0, len(*x)-1; i < j; i, j = i+1, j-1 {
		(*x)[i], (*x)[j] = (*x)[j], (*x)[i]
	}
	return x
}
func (x *Instrs) StrWrt(b *strings.Builder) {
	b.WriteRune('[')
	for n, v := range *x {
		if n != 0 {
			b.WriteRune(' ')
		}
		v.StrWrt(b)
	}
	b.WriteRune(']')
}
