package ana

import (
	"math/rand"
	"strings"
	"sys/bsc/bnd"
	"sys/bsc/bol"
	"sys/bsc/flts"
	"sys/bsc/strs"
	"sys/bsc/tmes"
	"sys/bsc/unt"
	"sys/bsc/unts"
	"time"
)

type (
	Prfms    []*Prfm
	PrfmsScp struct {
		Idx uint32
		Arr []*Prfms
	}
)

func NewPrfms(vs ...*Prfm) *Prfms {
	r := Prfms(vs)
	return &r
}
func MakePrfms(cap unt.Unt) *Prfms {
	r := make(Prfms, cap)
	return &r
}
func MakeEmpPrfms(cap unt.Unt) *Prfms {
	r := make(Prfms, 0, cap)
	return &r
}
func (x *Prfms) Ok() bol.Bol { return len(*x) != 0 }
func (x *Prfms) Cnt() unt.Unt {
	if x == nil {
		return 0
	}
	return unt.Unt(len(*x))
}
func (x *Prfms) Cpy() *Prfms {
	r := make(Prfms, len(*x))
	copy(r, *x)
	return &r
}
func (x *Prfms) Clr() *Prfms {
	*x = (*x)[:0]
	return x
}
func (x *Prfms) Rand() *Prfms {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	perm := r.Perm(len(*x))
	for i, randIdx := range perm {
		(*x)[i], (*x)[randIdx] = (*x)[randIdx], (*x)[i]
	}
	return x
}
func (x *Prfms) Mrg(a ...*Prfms) *Prfms {
	for _, v := range a {
		*x = append(*x, *v...)
	}
	return x
}
func (x *Prfms) Push(a ...*Prfm) *Prfms {
	*x = append(*x, a...)
	return x
}
func (x *Prfms) Pop() (r *Prfm) {
	if len(*x) == 0 {
		return r
	}
	r = (*x)[len(*x)-1]
	*x = (*x)[:len(*x)-1]
	return r
}
func (x *Prfms) Que(vs ...*Prfm) *Prfms {
	*x = append(*x, vs...)
	return x
}
func (x *Prfms) Dque() (r *Prfm) {
	r = (*x)[0]
	if len(*x) == 1 {
		*x = (*x)[:0]
	} else {
		copy(*x, (*x)[1:])
		*x = (*x)[:len(*x)-1]
	}
	return r
}
func (x *Prfms) Ins(idx unt.Unt, elm *Prfm) *Prfms {
	*x = append((*x)[:idx], append(Prfms{elm}, (*x)[idx:]...)...)
	return x
}
func (x *Prfms) Upd(idx unt.Unt, elm *Prfm) *Prfms {
	(*x)[idx] = elm
	return x
}
func (x *Prfms) Del(idx unt.Unt) (r *Prfm) {
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
func (x *Prfms) At(idx unt.Unt) *Prfm { return (*x)[idx] }
func (x *Prfms) In(idx, lim unt.Unt) *Prfms {
	r := (*x)[idx:lim]
	return &r
}
func (x *Prfms) InBnd(b bnd.Bnd) *Prfms {
	r := (*x)[b.Idx:b.Lim]
	return &r
}
func (x *Prfms) From(idx unt.Unt) *Prfms {
	var r Prfms
	if idx < unt.Unt(len(*x)) {
		r = (*x)[idx:]
	} else {
		r = (*x)[:0]
	}
	return &r
}
func (x *Prfms) To(lim unt.Unt) *Prfms {
	if lim > unt.Unt(len(*x)) {
		return x
	}
	r := (*x)[:lim]
	return &r
}
func (x *Prfms) Fst() *Prfm      { return (*x)[0] }
func (x *Prfms) Mdl() *Prfm      { return (*x)[len(*x)/2] }
func (x *Prfms) Lst() *Prfm      { return (*x)[len(*x)-1] }
func (x *Prfms) FstIdx() unt.Unt { return 0 }
func (x *Prfms) MdlIdx() unt.Unt { return unt.Unt(len(*x) / 2) }
func (x *Prfms) LstIdx() unt.Unt { return unt.Unt(len(*x) - 1) }
func (x *Prfms) Rev() *Prfms {
	for i, j := 0, len(*x)-1; i < j; i, j = i+1, j-1 {
		(*x)[i], (*x)[j] = (*x)[j], (*x)[i]
	}
	return x
}
func (x *Prfms) PnlPcts() (r *flts.Flts) {
	r = flts.Make(x.Cnt())
	for n, v := range *x {
		(*r)[n] = v.PnlPct
	}
	return r
}
func (x *Prfms) ScsPcts() (r *flts.Flts) {
	r = flts.Make(x.Cnt())
	for n, v := range *x {
		(*r)[n] = v.ScsPct
	}
	return r
}
func (x *Prfms) PipPerDays() (r *flts.Flts) {
	r = flts.Make(x.Cnt())
	for n, v := range *x {
		(*r)[n] = v.PipPerDay
	}
	return r
}
func (x *Prfms) UsdPerDays() (r *flts.Flts) {
	r = flts.Make(x.Cnt())
	for n, v := range *x {
		(*r)[n] = v.UsdPerDay
	}
	return r
}
func (x *Prfms) ScsPerDays() (r *flts.Flts) {
	r = flts.Make(x.Cnt())
	for n, v := range *x {
		(*r)[n] = v.ScsPerDay
	}
	return r
}
func (x *Prfms) OpnPerDays() (r *flts.Flts) {
	r = flts.Make(x.Cnt())
	for n, v := range *x {
		(*r)[n] = v.OpnPerDay
	}
	return r
}
func (x *Prfms) PnlUsds() (r *flts.Flts) {
	r = flts.Make(x.Cnt())
	for n, v := range *x {
		(*r)[n] = v.PnlUsd
	}
	return r
}
func (x *Prfms) PipAvgs() (r *flts.Flts) {
	r = flts.Make(x.Cnt())
	for n, v := range *x {
		(*r)[n] = v.PipAvg
	}
	return r
}
func (x *Prfms) PipMdns() (r *flts.Flts) {
	r = flts.Make(x.Cnt())
	for n, v := range *x {
		(*r)[n] = v.PipMdn
	}
	return r
}
func (x *Prfms) PipMins() (r *flts.Flts) {
	r = flts.Make(x.Cnt())
	for n, v := range *x {
		(*r)[n] = v.PipMin
	}
	return r
}
func (x *Prfms) PipMaxs() (r *flts.Flts) {
	r = flts.Make(x.Cnt())
	for n, v := range *x {
		(*r)[n] = v.PipMax
	}
	return r
}
func (x *Prfms) PipSums() (r *flts.Flts) {
	r = flts.Make(x.Cnt())
	for n, v := range *x {
		(*r)[n] = v.PipSum
	}
	return r
}
func (x *Prfms) DurAvgs() (r *tmes.Tmes) {
	r = tmes.Make(x.Cnt())
	for n, v := range *x {
		(*r)[n] = v.DurAvg
	}
	return r
}
func (x *Prfms) DurMdns() (r *tmes.Tmes) {
	r = tmes.Make(x.Cnt())
	for n, v := range *x {
		(*r)[n] = v.DurMdn
	}
	return r
}
func (x *Prfms) DurMins() (r *tmes.Tmes) {
	r = tmes.Make(x.Cnt())
	for n, v := range *x {
		(*r)[n] = v.DurMin
	}
	return r
}
func (x *Prfms) DurMaxs() (r *tmes.Tmes) {
	r = tmes.Make(x.Cnt())
	for n, v := range *x {
		(*r)[n] = v.DurMax
	}
	return r
}
func (x *Prfms) LosLimMaxs() (r *flts.Flts) {
	r = flts.Make(x.Cnt())
	for n, v := range *x {
		(*r)[n] = v.LosLimMax
	}
	return r
}
func (x *Prfms) DurLimMaxs() (r *tmes.Tmes) {
	r = tmes.Make(x.Cnt())
	for n, v := range *x {
		(*r)[n] = v.DurLimMax
	}
	return r
}
func (x *Prfms) DayCnts() (r *unts.Unts) {
	r = unts.Make(x.Cnt())
	for n, v := range *x {
		(*r)[n] = v.DayCnt
	}
	return r
}
func (x *Prfms) TrdCnts() (r *unts.Unts) {
	r = unts.Make(x.Cnt())
	for n, v := range *x {
		(*r)[n] = v.TrdCnt
	}
	return r
}
func (x *Prfms) TrdPcts() (r *flts.Flts) {
	r = flts.Make(x.Cnt())
	for n, v := range *x {
		(*r)[n] = v.TrdPct
	}
	return r
}
func (x *Prfms) CstTotUsds() (r *flts.Flts) {
	r = flts.Make(x.Cnt())
	for n, v := range *x {
		(*r)[n] = v.CstTotUsd
	}
	return r
}
func (x *Prfms) CstSpdUsds() (r *flts.Flts) {
	r = flts.Make(x.Cnt())
	for n, v := range *x {
		(*r)[n] = v.CstSpdUsd
	}
	return r
}
func (x *Prfms) CstComUsds() (r *flts.Flts) {
	r = flts.Make(x.Cnt())
	for n, v := range *x {
		(*r)[n] = v.CstComUsd
	}
	return r
}
func (x *Prfms) Pths() (r *strs.Strs) {
	r = strs.Make(x.Cnt())
	for n, v := range *x {
		(*r)[n] = v.Pth
	}
	return r
}
func (x *Prfms) StrWrt(b *strings.Builder) {
	b.WriteRune('[')
	for n, v := range *x {
		if n != 0 {
			b.WriteRune(' ')
		}
		v.StrWrt(b)
	}
	b.WriteRune(']')
}
func (x *Prfms) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
