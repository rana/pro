package ana

import (
	"math/rand"
	"sys/bsc/bnd"
	"sys/bsc/bol"
	"sys/bsc/unt"
	"time"
)

type (
	TmeIdxs []TmeIdx
)

func NewTmeIdxs(vs ...TmeIdx) *TmeIdxs {
	r := TmeIdxs(vs)
	return &r
}
func MakeTmeIdxs(cap unt.Unt) *TmeIdxs {
	r := make(TmeIdxs, cap)
	return &r
}
func MakeEmpTmeIdxs(cap unt.Unt) *TmeIdxs {
	r := make(TmeIdxs, 0, cap)
	return &r
}
func (x *TmeIdxs) Ok() bol.Bol { return len(*x) != 0 }
func (x *TmeIdxs) Cnt() unt.Unt {
	if x == nil {
		return 0
	}
	return unt.Unt(len(*x))
}
func (x *TmeIdxs) Cpy() *TmeIdxs {
	r := make(TmeIdxs, len(*x))
	copy(r, *x)
	return &r
}
func (x *TmeIdxs) Clr() *TmeIdxs {
	*x = (*x)[:0]
	return x
}
func (x *TmeIdxs) Rand() *TmeIdxs {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	perm := r.Perm(len(*x))
	for i, randIdx := range perm {
		(*x)[i], (*x)[randIdx] = (*x)[randIdx], (*x)[i]
	}
	return x
}
func (x *TmeIdxs) Mrg(a ...*TmeIdxs) *TmeIdxs {
	for _, v := range a {
		*x = append(*x, *v...)
	}
	return x
}
func (x *TmeIdxs) Push(a ...TmeIdx) *TmeIdxs {
	*x = append(*x, a...)
	return x
}
func (x *TmeIdxs) Pop() (r TmeIdx) {
	if len(*x) == 0 {
		return r
	}
	r = (*x)[len(*x)-1]
	*x = (*x)[:len(*x)-1]
	return r
}
func (x *TmeIdxs) Que(vs ...TmeIdx) *TmeIdxs {
	*x = append(*x, vs...)
	return x
}
func (x *TmeIdxs) Dque() (r TmeIdx) {
	r = (*x)[0]
	if len(*x) == 1 {
		*x = (*x)[:0]
	} else {
		copy(*x, (*x)[1:])
		*x = (*x)[:len(*x)-1]
	}
	return r
}
func (x *TmeIdxs) Ins(idx unt.Unt, elm TmeIdx) *TmeIdxs {
	*x = append((*x)[:idx], append(TmeIdxs{elm}, (*x)[idx:]...)...)
	return x
}
func (x *TmeIdxs) Upd(idx unt.Unt, elm TmeIdx) *TmeIdxs {
	(*x)[idx] = elm
	return x
}
func (x *TmeIdxs) Del(idx unt.Unt) (r TmeIdx) {
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
func (x *TmeIdxs) At(idx unt.Unt) TmeIdx { return (*x)[idx] }
func (x *TmeIdxs) In(idx, lim unt.Unt) *TmeIdxs {
	r := (*x)[idx:lim]
	return &r
}
func (x *TmeIdxs) InBnd(b bnd.Bnd) *TmeIdxs {
	r := (*x)[b.Idx:b.Lim]
	return &r
}
func (x *TmeIdxs) From(idx unt.Unt) *TmeIdxs {
	var r TmeIdxs
	if idx < unt.Unt(len(*x)) {
		r = (*x)[idx:]
	} else {
		r = (*x)[:0]
	}
	return &r
}
func (x *TmeIdxs) To(lim unt.Unt) *TmeIdxs {
	if lim > unt.Unt(len(*x)) {
		return x
	}
	r := (*x)[:lim]
	return &r
}
func (x *TmeIdxs) Fst() TmeIdx     { return (*x)[0] }
func (x *TmeIdxs) Mdl() TmeIdx     { return (*x)[len(*x)/2] }
func (x *TmeIdxs) Lst() TmeIdx     { return (*x)[len(*x)-1] }
func (x *TmeIdxs) FstIdx() unt.Unt { return 0 }
func (x *TmeIdxs) MdlIdx() unt.Unt { return unt.Unt(len(*x) / 2) }
func (x *TmeIdxs) LstIdx() unt.Unt { return unt.Unt(len(*x) - 1) }
func (x *TmeIdxs) Rev() *TmeIdxs {
	for i, j := 0, len(*x)-1; i < j; i, j = i+1, j-1 {
		(*x)[i], (*x)[j] = (*x)[j], (*x)[i]
	}
	return x
}
