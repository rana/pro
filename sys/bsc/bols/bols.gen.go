package bols

import (
	"bytes"
	"encoding/binary"
	"math/rand"
	"strings"
	"sys/bsc/bnd"
	"sys/bsc/bol"
	"sys/bsc/unt"
	"time"
)

type (
	Bols    []bol.Bol
	BolsScp struct {
		Idx uint32
		Arr []*Bols
	}
)

func New(vs ...bol.Bol) *Bols {
	r := Bols(vs)
	return &r
}
func Make(cap unt.Unt) *Bols {
	r := make(Bols, cap)
	return &r
}
func MakeEmp(cap unt.Unt) *Bols {
	r := make(Bols, 0, cap)
	return &r
}
func (x *Bols) Ok() bol.Bol { return len(*x) != 0 }
func (x *Bols) Cnt() unt.Unt {
	if x == nil {
		return 0
	}
	return unt.Unt(len(*x))
}
func (x *Bols) Cpy() *Bols {
	r := make(Bols, len(*x))
	copy(r, *x)
	return &r
}
func (x *Bols) Clr() *Bols {
	*x = (*x)[:0]
	return x
}
func (x *Bols) Rand() *Bols {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	perm := r.Perm(len(*x))
	for i, randIdx := range perm {
		(*x)[i], (*x)[randIdx] = (*x)[randIdx], (*x)[i]
	}
	return x
}
func (x *Bols) Mrg(a ...*Bols) *Bols {
	for _, v := range a {
		*x = append(*x, *v...)
	}
	return x
}
func (x *Bols) Push(a ...bol.Bol) *Bols {
	*x = append(*x, a...)
	return x
}
func (x *Bols) Pop() (r bol.Bol) {
	if len(*x) == 0 {
		return r
	}
	r = (*x)[len(*x)-1]
	*x = (*x)[:len(*x)-1]
	return r
}
func (x *Bols) Que(vs ...bol.Bol) *Bols {
	*x = append(*x, vs...)
	return x
}
func (x *Bols) Dque() (r bol.Bol) {
	r = (*x)[0]
	if len(*x) == 1 {
		*x = (*x)[:0]
	} else {
		copy(*x, (*x)[1:])
		*x = (*x)[:len(*x)-1]
	}
	return r
}
func (x *Bols) Ins(idx unt.Unt, elm bol.Bol) *Bols {
	*x = append((*x)[:idx], append(Bols{elm}, (*x)[idx:]...)...)
	return x
}
func (x *Bols) Upd(idx unt.Unt, elm bol.Bol) *Bols {
	(*x)[idx] = elm
	return x
}
func (x *Bols) Del(idx unt.Unt) (r bol.Bol) {
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
func (x *Bols) At(idx unt.Unt) bol.Bol { return (*x)[idx] }
func (x *Bols) In(idx, lim unt.Unt) *Bols {
	r := (*x)[idx:lim]
	return &r
}
func (x *Bols) InBnd(b bnd.Bnd) *Bols {
	r := (*x)[b.Idx:b.Lim]
	return &r
}
func (x *Bols) From(idx unt.Unt) *Bols {
	var r Bols
	if idx < unt.Unt(len(*x)) {
		r = (*x)[idx:]
	} else {
		r = (*x)[:0]
	}
	return &r
}
func (x *Bols) To(lim unt.Unt) *Bols {
	if lim > unt.Unt(len(*x)) {
		return x
	}
	r := (*x)[:lim]
	return &r
}
func (x *Bols) Fst() bol.Bol    { return (*x)[0] }
func (x *Bols) Mdl() bol.Bol    { return (*x)[len(*x)/2] }
func (x *Bols) Lst() bol.Bol    { return (*x)[len(*x)-1] }
func (x *Bols) FstIdx() unt.Unt { return 0 }
func (x *Bols) MdlIdx() unt.Unt { return unt.Unt(len(*x) / 2) }
func (x *Bols) LstIdx() unt.Unt { return unt.Unt(len(*x) - 1) }
func (x *Bols) Rev() *Bols {
	for i, j := 0, len(*x)-1; i < j; i, j = i+1, j-1 {
		(*x)[i], (*x)[j] = (*x)[j], (*x)[i]
	}
	return x
}
func (x *Bols) StrWrt(b *strings.Builder) {
	b.WriteRune('[')
	for n, v := range *x {
		if n != 0 {
			b.WriteRune(' ')
		}
		v.StrWrt(b)
	}
	b.WriteRune(']')
}
func (x *Bols) BytWrt(b *bytes.Buffer) {
	bLen := make([]byte, 4) // array length
	binary.LittleEndian.PutUint32(bLen, uint32(len(*x)))
	b.Write(bLen)
	for _, v := range *x {
		v.BytWrt(b)
	}
}
func (x *Bols) BytRed(b []byte) (idx int) {
	if len(b) >= 4 {
		*x = make(Bols, binary.LittleEndian.Uint32(b[:4])) // overwrite any previous existing
		idx = 4
		for n := 0; n < len(*x); n++ {
			(*x)[n].BytRed(b[idx : idx+bol.Size])
			idx += bol.Size
		}
	}
	return idx
}
