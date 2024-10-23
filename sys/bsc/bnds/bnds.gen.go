package bnds

import (
	"bytes"
	"encoding/binary"
	"math/rand"
	"runtime"
	"strings"
	"sys"
	"sys/bsc/bnd"
	"sys/bsc/bol"
	"sys/bsc/unt"
	"time"
)

type (
	Bnds    []bnd.Bnd
	BndsScp struct {
		Idx uint32
		Arr []*Bnds
	}
)

func New(vs ...bnd.Bnd) *Bnds {
	r := Bnds(vs)
	return &r
}
func Make(cap unt.Unt) *Bnds {
	r := make(Bnds, cap)
	return &r
}
func MakeEmp(cap unt.Unt) *Bnds {
	r := make(Bnds, 0, cap)
	return &r
}
func Segs(elmCnt unt.Unt) (r *Bnds, acts []sys.Act) {
	r = New()
	segCnt := unt.Unt(runtime.NumCPU())
	segLen := elmCnt.Div(segCnt)
	if segLen.Lss(unt.MinSegLen) {
		segLen = elmCnt
		segCnt = unt.One
	}
	idx, lim := unt.Zero, segLen
	for n := unt.Zero; n < segCnt; n++ {
		if n > 0 {
			idx = lim
			lim += segLen
		}
		if idx >= elmCnt {
			break
		}
		if n == segCnt-1 || lim > elmCnt {
			lim = elmCnt
		}
		r.Push(bnd.Bnd{Idx: idx, Lim: lim})
	}
	return r, make([]sys.Act, len(*r))
}
func (x *Bnds) Ok() bol.Bol { return len(*x) != 0 }
func (x *Bnds) Cnt() unt.Unt {
	if x == nil {
		return 0
	}
	return unt.Unt(len(*x))
}
func (x *Bnds) Cpy() *Bnds {
	r := make(Bnds, len(*x))
	copy(r, *x)
	return &r
}
func (x *Bnds) Clr() *Bnds {
	*x = (*x)[:0]
	return x
}
func (x *Bnds) Rand() *Bnds {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	perm := r.Perm(len(*x))
	for i, randIdx := range perm {
		(*x)[i], (*x)[randIdx] = (*x)[randIdx], (*x)[i]
	}
	return x
}
func (x *Bnds) Mrg(a ...*Bnds) *Bnds {
	for _, v := range a {
		*x = append(*x, *v...)
	}
	return x
}
func (x *Bnds) Push(a ...bnd.Bnd) *Bnds {
	*x = append(*x, a...)
	return x
}
func (x *Bnds) Pop() (r bnd.Bnd) {
	if len(*x) == 0 {
		return r
	}
	r = (*x)[len(*x)-1]
	*x = (*x)[:len(*x)-1]
	return r
}
func (x *Bnds) Que(vs ...bnd.Bnd) *Bnds {
	*x = append(*x, vs...)
	return x
}
func (x *Bnds) Dque() (r bnd.Bnd) {
	r = (*x)[0]
	if len(*x) == 1 {
		*x = (*x)[:0]
	} else {
		copy(*x, (*x)[1:])
		*x = (*x)[:len(*x)-1]
	}
	return r
}
func (x *Bnds) Ins(idx unt.Unt, elm bnd.Bnd) *Bnds {
	*x = append((*x)[:idx], append(Bnds{elm}, (*x)[idx:]...)...)
	return x
}
func (x *Bnds) Upd(idx unt.Unt, elm bnd.Bnd) *Bnds {
	(*x)[idx] = elm
	return x
}
func (x *Bnds) Del(idx unt.Unt) (r bnd.Bnd) {
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
func (x *Bnds) At(idx unt.Unt) bnd.Bnd { return (*x)[idx] }
func (x *Bnds) In(idx, lim unt.Unt) *Bnds {
	r := (*x)[idx:lim]
	return &r
}
func (x *Bnds) InBnd(b bnd.Bnd) *Bnds {
	r := (*x)[b.Idx:b.Lim]
	return &r
}
func (x *Bnds) From(idx unt.Unt) *Bnds {
	var r Bnds
	if idx < unt.Unt(len(*x)) {
		r = (*x)[idx:]
	} else {
		r = (*x)[:0]
	}
	return &r
}
func (x *Bnds) To(lim unt.Unt) *Bnds {
	if lim > unt.Unt(len(*x)) {
		return x
	}
	r := (*x)[:lim]
	return &r
}
func (x *Bnds) Fst() bnd.Bnd    { return (*x)[0] }
func (x *Bnds) Mdl() bnd.Bnd    { return (*x)[len(*x)/2] }
func (x *Bnds) Lst() bnd.Bnd    { return (*x)[len(*x)-1] }
func (x *Bnds) FstIdx() unt.Unt { return 0 }
func (x *Bnds) MdlIdx() unt.Unt { return unt.Unt(len(*x) / 2) }
func (x *Bnds) LstIdx() unt.Unt { return unt.Unt(len(*x) - 1) }
func (x *Bnds) Rev() *Bnds {
	for i, j := 0, len(*x)-1; i < j; i, j = i+1, j-1 {
		(*x)[i], (*x)[j] = (*x)[j], (*x)[i]
	}
	return x
}
func (x *Bnds) StrWrt(b *strings.Builder) {
	b.WriteRune('[')
	for n, v := range *x {
		if n != 0 {
			b.WriteRune(' ')
		}
		v.StrWrt(b)
	}
	b.WriteRune(']')
}
func (x *Bnds) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *Bnds) BytWrt(b *bytes.Buffer) {
	bLen := make([]byte, 4) // array length
	binary.LittleEndian.PutUint32(bLen, uint32(len(*x)))
	b.Write(bLen)
	for _, v := range *x {
		v.BytWrt(b)
	}
}
func (x *Bnds) BytRed(b []byte) (idx int) {
	if len(b) >= 4 {
		*x = make(Bnds, binary.LittleEndian.Uint32(b[:4])) // overwrite any previous existing
		idx = 4
		for n := 0; n < len(*x); n++ {
			(*x)[n].BytRed(b[idx : idx+bnd.Size])
			idx += bnd.Size
		}
	}
	return idx
}
