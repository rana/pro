package ana

import (
	"math/rand"
	"strings"
	"sys/bsc/bnd"
	"sys/bsc/bol"
	"sys/bsc/bols"
	"sys/bsc/flt"
	"sys/bsc/flts"
	"sys/bsc/str"
	"sys/bsc/strs"
	"sys/bsc/tme"
	"sys/bsc/tmes"
	"sys/bsc/unt"
	"time"
)

type (
	Trds    []*Trd
	TrdsScp struct {
		Idx uint32
		Arr []*Trds
	}
)

func NewTrds(vs ...*Trd) *Trds {
	r := Trds(vs)
	return &r
}
func MakeTrds(cap unt.Unt) *Trds {
	r := make(Trds, cap)
	return &r
}
func MakeEmpTrds(cap unt.Unt) *Trds {
	r := make(Trds, 0, cap)
	return &r
}
func (x *Trds) Ok() bol.Bol { return len(*x) != 0 }
func (x *Trds) Cnt() unt.Unt {
	if x == nil {
		return 0
	}
	return unt.Unt(len(*x))
}
func (x *Trds) Cpy() *Trds {
	r := make(Trds, len(*x))
	copy(r, *x)
	return &r
}
func (x *Trds) Clr() *Trds {
	*x = (*x)[:0]
	return x
}
func (x *Trds) Rand() *Trds {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	perm := r.Perm(len(*x))
	for i, randIdx := range perm {
		(*x)[i], (*x)[randIdx] = (*x)[randIdx], (*x)[i]
	}
	return x
}
func (x *Trds) Mrg(a ...*Trds) *Trds {
	for _, v := range a {
		*x = append(*x, *v...)
	}
	return x
}
func (x *Trds) Push(a ...*Trd) *Trds {
	*x = append(*x, a...)
	return x
}
func (x *Trds) Pop() (r *Trd) {
	if len(*x) == 0 {
		return r
	}
	r = (*x)[len(*x)-1]
	*x = (*x)[:len(*x)-1]
	return r
}
func (x *Trds) Que(vs ...*Trd) *Trds {
	*x = append(*x, vs...)
	return x
}
func (x *Trds) Dque() (r *Trd) {
	r = (*x)[0]
	if len(*x) == 1 {
		*x = (*x)[:0]
	} else {
		copy(*x, (*x)[1:])
		*x = (*x)[:len(*x)-1]
	}
	return r
}
func (x *Trds) Ins(idx unt.Unt, elm *Trd) *Trds {
	*x = append((*x)[:idx], append(Trds{elm}, (*x)[idx:]...)...)
	return x
}
func (x *Trds) Upd(idx unt.Unt, elm *Trd) *Trds {
	(*x)[idx] = elm
	return x
}
func (x *Trds) Del(idx unt.Unt) (r *Trd) {
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
func (x *Trds) At(idx unt.Unt) *Trd { return (*x)[idx] }
func (x *Trds) In(idx, lim unt.Unt) *Trds {
	r := (*x)[idx:lim]
	return &r
}
func (x *Trds) InBnd(b bnd.Bnd) *Trds {
	r := (*x)[b.Idx:b.Lim]
	return &r
}
func (x *Trds) From(idx unt.Unt) *Trds {
	var r Trds
	if idx < unt.Unt(len(*x)) {
		r = (*x)[idx:]
	} else {
		r = (*x)[:0]
	}
	return &r
}
func (x *Trds) To(lim unt.Unt) *Trds {
	if lim > unt.Unt(len(*x)) {
		return x
	}
	r := (*x)[:lim]
	return &r
}
func (x *Trds) Fst() *Trd       { return (*x)[0] }
func (x *Trds) Mdl() *Trd       { return (*x)[len(*x)/2] }
func (x *Trds) Lst() *Trd       { return (*x)[len(*x)-1] }
func (x *Trds) FstIdx() unt.Unt { return 0 }
func (x *Trds) MdlIdx() unt.Unt { return unt.Unt(len(*x) / 2) }
func (x *Trds) LstIdx() unt.Unt { return unt.Unt(len(*x) - 1) }
func (x *Trds) Rev() *Trds {
	for i, j := 0, len(*x)-1; i < j; i, j = i+1, j-1 {
		(*x)[i], (*x)[j] = (*x)[j], (*x)[i]
	}
	return x
}
func (x *Trds) SelClsResEql(v str.Str) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.ClsRes.Eql(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelClsResNeq(v str.Str) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.ClsRes.Neq(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelClsResLss(v str.Str) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.ClsRes.Lss(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelClsResGtr(v str.Str) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.ClsRes.Gtr(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelClsResLeq(v str.Str) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.ClsRes.Leq(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelClsResGeq(v str.Str) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.ClsRes.Geq(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelClsResSplt(v str.Str) (btm, top *Trds) {
	btm = NewTrds()
	top = NewTrds()
	for _, cur := range *x {
		if cur.ClsRes > v {
			*top = append(*top, cur)
		} else {
			*btm = append(*btm, cur)
		}
	}
	return btm, top
}
func (x *Trds) SelClsReqEql(v str.Str) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.ClsReq.Eql(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelClsReqNeq(v str.Str) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.ClsReq.Neq(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelClsReqLss(v str.Str) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.ClsReq.Lss(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelClsReqGtr(v str.Str) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.ClsReq.Gtr(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelClsReqLeq(v str.Str) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.ClsReq.Leq(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelClsReqGeq(v str.Str) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.ClsReq.Geq(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelClsReqSplt(v str.Str) (btm, top *Trds) {
	btm = NewTrds()
	top = NewTrds()
	for _, cur := range *x {
		if cur.ClsReq > v {
			*top = append(*top, cur)
		} else {
			*btm = append(*btm, cur)
		}
	}
	return btm, top
}
func (x *Trds) SelOpnResEql(v str.Str) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.OpnRes.Eql(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelOpnResNeq(v str.Str) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.OpnRes.Neq(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelOpnResLss(v str.Str) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.OpnRes.Lss(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelOpnResGtr(v str.Str) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.OpnRes.Gtr(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelOpnResLeq(v str.Str) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.OpnRes.Leq(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelOpnResGeq(v str.Str) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.OpnRes.Geq(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelOpnResSplt(v str.Str) (btm, top *Trds) {
	btm = NewTrds()
	top = NewTrds()
	for _, cur := range *x {
		if cur.OpnRes > v {
			*top = append(*top, cur)
		} else {
			*btm = append(*btm, cur)
		}
	}
	return btm, top
}
func (x *Trds) SelOpnReqEql(v str.Str) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.OpnReq.Eql(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelOpnReqNeq(v str.Str) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.OpnReq.Neq(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelOpnReqLss(v str.Str) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.OpnReq.Lss(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelOpnReqGtr(v str.Str) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.OpnReq.Gtr(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelOpnReqLeq(v str.Str) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.OpnReq.Leq(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelOpnReqGeq(v str.Str) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.OpnReq.Geq(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelOpnReqSplt(v str.Str) (btm, top *Trds) {
	btm = NewTrds()
	top = NewTrds()
	for _, cur := range *x {
		if cur.OpnReq > v {
			*top = append(*top, cur)
		} else {
			*btm = append(*btm, cur)
		}
	}
	return btm, top
}
func (x *Trds) SelInstrEql(v str.Str) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.Instr.Eql(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelInstrNeq(v str.Str) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.Instr.Neq(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelInstrLss(v str.Str) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.Instr.Lss(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelInstrGtr(v str.Str) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.Instr.Gtr(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelInstrLeq(v str.Str) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.Instr.Leq(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelInstrGeq(v str.Str) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.Instr.Geq(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelInstrSplt(v str.Str) (btm, top *Trds) {
	btm = NewTrds()
	top = NewTrds()
	for _, cur := range *x {
		if cur.Instr > v {
			*top = append(*top, cur)
		} else {
			*btm = append(*btm, cur)
		}
	}
	return btm, top
}
func (x *Trds) SelUnitsEql(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.Units.Eql(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelUnitsNeq(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.Units.Neq(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelUnitsLss(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.Units.Lss(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelUnitsGtr(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.Units.Gtr(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelUnitsLeq(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.Units.Leq(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelUnitsGeq(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.Units.Geq(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelUnitsSplt(v flt.Flt) (btm, top *Trds) {
	btm = NewTrds()
	top = NewTrds()
	for _, cur := range *x {
		if cur.Units > v {
			*top = append(*top, cur)
		} else {
			*btm = append(*btm, cur)
		}
	}
	return btm, top
}
func (x *Trds) SelMrgnRtioEql(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.MrgnRtio.Eql(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelMrgnRtioNeq(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.MrgnRtio.Neq(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelMrgnRtioLss(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.MrgnRtio.Lss(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelMrgnRtioGtr(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.MrgnRtio.Gtr(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelMrgnRtioLeq(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.MrgnRtio.Leq(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelMrgnRtioGeq(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.MrgnRtio.Geq(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelMrgnRtioSplt(v flt.Flt) (btm, top *Trds) {
	btm = NewTrds()
	top = NewTrds()
	for _, cur := range *x {
		if cur.MrgnRtio > v {
			*top = append(*top, cur)
		} else {
			*btm = append(*btm, cur)
		}
	}
	return btm, top
}
func (x *Trds) SelTrdPctEql(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.TrdPct.Eql(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelTrdPctNeq(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.TrdPct.Neq(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelTrdPctLss(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.TrdPct.Lss(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelTrdPctGtr(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.TrdPct.Gtr(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelTrdPctLeq(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.TrdPct.Leq(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelTrdPctGeq(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.TrdPct.Geq(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelTrdPctSplt(v flt.Flt) (btm, top *Trds) {
	btm = NewTrds()
	top = NewTrds()
	for _, cur := range *x {
		if cur.TrdPct > v {
			*top = append(*top, cur)
		} else {
			*btm = append(*btm, cur)
		}
	}
	return btm, top
}
func (x *Trds) SelClsBalUsdActEql(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.ClsBalUsdAct.Eql(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelClsBalUsdActNeq(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.ClsBalUsdAct.Neq(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelClsBalUsdActLss(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.ClsBalUsdAct.Lss(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelClsBalUsdActGtr(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.ClsBalUsdAct.Gtr(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelClsBalUsdActLeq(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.ClsBalUsdAct.Leq(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelClsBalUsdActGeq(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.ClsBalUsdAct.Geq(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelClsBalUsdActSplt(v flt.Flt) (btm, top *Trds) {
	btm = NewTrds()
	top = NewTrds()
	for _, cur := range *x {
		if cur.ClsBalUsdAct > v {
			*top = append(*top, cur)
		} else {
			*btm = append(*btm, cur)
		}
	}
	return btm, top
}
func (x *Trds) SelClsBalUsdEql(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.ClsBalUsd.Eql(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelClsBalUsdNeq(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.ClsBalUsd.Neq(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelClsBalUsdLss(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.ClsBalUsd.Lss(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelClsBalUsdGtr(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.ClsBalUsd.Gtr(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelClsBalUsdLeq(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.ClsBalUsd.Leq(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelClsBalUsdGeq(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.ClsBalUsd.Geq(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelClsBalUsdSplt(v flt.Flt) (btm, top *Trds) {
	btm = NewTrds()
	top = NewTrds()
	for _, cur := range *x {
		if cur.ClsBalUsd > v {
			*top = append(*top, cur)
		} else {
			*btm = append(*btm, cur)
		}
	}
	return btm, top
}
func (x *Trds) SelOpnBalUsdEql(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.OpnBalUsd.Eql(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelOpnBalUsdNeq(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.OpnBalUsd.Neq(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelOpnBalUsdLss(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.OpnBalUsd.Lss(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelOpnBalUsdGtr(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.OpnBalUsd.Gtr(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelOpnBalUsdLeq(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.OpnBalUsd.Leq(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelOpnBalUsdGeq(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.OpnBalUsd.Geq(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelOpnBalUsdSplt(v flt.Flt) (btm, top *Trds) {
	btm = NewTrds()
	top = NewTrds()
	for _, cur := range *x {
		if cur.OpnBalUsd > v {
			*top = append(*top, cur)
		} else {
			*btm = append(*btm, cur)
		}
	}
	return btm, top
}
func (x *Trds) SelCstOpnSpdUsdEql(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.CstOpnSpdUsd.Eql(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelCstOpnSpdUsdNeq(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.CstOpnSpdUsd.Neq(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelCstOpnSpdUsdLss(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.CstOpnSpdUsd.Lss(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelCstOpnSpdUsdGtr(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.CstOpnSpdUsd.Gtr(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelCstOpnSpdUsdLeq(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.CstOpnSpdUsd.Leq(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelCstOpnSpdUsdGeq(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.CstOpnSpdUsd.Geq(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelCstOpnSpdUsdSplt(v flt.Flt) (btm, top *Trds) {
	btm = NewTrds()
	top = NewTrds()
	for _, cur := range *x {
		if cur.CstOpnSpdUsd > v {
			*top = append(*top, cur)
		} else {
			*btm = append(*btm, cur)
		}
	}
	return btm, top
}
func (x *Trds) SelCstClsSpdUsdEql(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.CstClsSpdUsd.Eql(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelCstClsSpdUsdNeq(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.CstClsSpdUsd.Neq(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelCstClsSpdUsdLss(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.CstClsSpdUsd.Lss(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelCstClsSpdUsdGtr(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.CstClsSpdUsd.Gtr(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelCstClsSpdUsdLeq(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.CstClsSpdUsd.Leq(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelCstClsSpdUsdGeq(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.CstClsSpdUsd.Geq(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelCstClsSpdUsdSplt(v flt.Flt) (btm, top *Trds) {
	btm = NewTrds()
	top = NewTrds()
	for _, cur := range *x {
		if cur.CstClsSpdUsd > v {
			*top = append(*top, cur)
		} else {
			*btm = append(*btm, cur)
		}
	}
	return btm, top
}
func (x *Trds) SelCstComUsdEql(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.CstComUsd.Eql(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelCstComUsdNeq(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.CstComUsd.Neq(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelCstComUsdLss(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.CstComUsd.Lss(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelCstComUsdGtr(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.CstComUsd.Gtr(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelCstComUsdLeq(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.CstComUsd.Leq(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelCstComUsdGeq(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.CstComUsd.Geq(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelCstComUsdSplt(v flt.Flt) (btm, top *Trds) {
	btm = NewTrds()
	top = NewTrds()
	for _, cur := range *x {
		if cur.CstComUsd > v {
			*top = append(*top, cur)
		} else {
			*btm = append(*btm, cur)
		}
	}
	return btm, top
}
func (x *Trds) SelPnlGrsUsdEql(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.PnlGrsUsd.Eql(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelPnlGrsUsdNeq(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.PnlGrsUsd.Neq(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelPnlGrsUsdLss(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.PnlGrsUsd.Lss(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelPnlGrsUsdGtr(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.PnlGrsUsd.Gtr(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelPnlGrsUsdLeq(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.PnlGrsUsd.Leq(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelPnlGrsUsdGeq(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.PnlGrsUsd.Geq(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelPnlGrsUsdSplt(v flt.Flt) (btm, top *Trds) {
	btm = NewTrds()
	top = NewTrds()
	for _, cur := range *x {
		if cur.PnlGrsUsd > v {
			*top = append(*top, cur)
		} else {
			*btm = append(*btm, cur)
		}
	}
	return btm, top
}
func (x *Trds) SelPnlUsdEql(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.PnlUsd.Eql(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelPnlUsdNeq(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.PnlUsd.Neq(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelPnlUsdLss(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.PnlUsd.Lss(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelPnlUsdGtr(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.PnlUsd.Gtr(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelPnlUsdLeq(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.PnlUsd.Leq(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelPnlUsdGeq(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.PnlUsd.Geq(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelPnlUsdSplt(v flt.Flt) (btm, top *Trds) {
	btm = NewTrds()
	top = NewTrds()
	for _, cur := range *x {
		if cur.PnlUsd > v {
			*top = append(*top, cur)
		} else {
			*btm = append(*btm, cur)
		}
	}
	return btm, top
}
func (x *Trds) SelPnlPctPredictEql(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.PnlPctPredict.Eql(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelPnlPctPredictNeq(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.PnlPctPredict.Neq(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelPnlPctPredictLss(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.PnlPctPredict.Lss(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelPnlPctPredictGtr(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.PnlPctPredict.Gtr(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelPnlPctPredictLeq(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.PnlPctPredict.Leq(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelPnlPctPredictGeq(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.PnlPctPredict.Geq(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelPnlPctPredictSplt(v flt.Flt) (btm, top *Trds) {
	btm = NewTrds()
	top = NewTrds()
	for _, cur := range *x {
		if cur.PnlPctPredict > v {
			*top = append(*top, cur)
		} else {
			*btm = append(*btm, cur)
		}
	}
	return btm, top
}
func (x *Trds) SelPnlPctEql(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.PnlPct.Eql(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelPnlPctNeq(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.PnlPct.Neq(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelPnlPctLss(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.PnlPct.Lss(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelPnlPctGtr(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.PnlPct.Gtr(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelPnlPctLeq(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.PnlPct.Leq(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelPnlPctGeq(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.PnlPct.Geq(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelPnlPctSplt(v flt.Flt) (btm, top *Trds) {
	btm = NewTrds()
	top = NewTrds()
	for _, cur := range *x {
		if cur.PnlPct > v {
			*top = append(*top, cur)
		} else {
			*btm = append(*btm, cur)
		}
	}
	return btm, top
}
func (x *Trds) SelIsLongEql(v bol.Bol) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.IsLong.Eql(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelIsLongNeq(v bol.Bol) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.IsLong.Neq(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelDurEql(v tme.Tme) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.Dur.Eql(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelDurNeq(v tme.Tme) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.Dur.Neq(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelDurLss(v tme.Tme) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.Dur.Lss(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelDurGtr(v tme.Tme) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.Dur.Gtr(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelDurLeq(v tme.Tme) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.Dur.Leq(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelDurGeq(v tme.Tme) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.Dur.Geq(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelDurSplt(v tme.Tme) (btm, top *Trds) {
	btm = NewTrds()
	top = NewTrds()
	for _, cur := range *x {
		if cur.Dur > v {
			*top = append(*top, cur)
		} else {
			*btm = append(*btm, cur)
		}
	}
	return btm, top
}
func (x *Trds) SelPipEql(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.Pip.Eql(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelPipNeq(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.Pip.Neq(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelPipLss(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.Pip.Lss(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelPipGtr(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.Pip.Gtr(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelPipLeq(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.Pip.Leq(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelPipGeq(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.Pip.Geq(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelPipSplt(v flt.Flt) (btm, top *Trds) {
	btm = NewTrds()
	top = NewTrds()
	for _, cur := range *x {
		if cur.Pip > v {
			*top = append(*top, cur)
		} else {
			*btm = append(*btm, cur)
		}
	}
	return btm, top
}
func (x *Trds) SelClsRsnEql(v str.Str) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.ClsRsn.Eql(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelClsRsnNeq(v str.Str) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.ClsRsn.Neq(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelClsRsnLss(v str.Str) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.ClsRsn.Lss(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelClsRsnGtr(v str.Str) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.ClsRsn.Gtr(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelClsRsnLeq(v str.Str) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.ClsRsn.Leq(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelClsRsnGeq(v str.Str) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.ClsRsn.Geq(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelClsRsnSplt(v str.Str) (btm, top *Trds) {
	btm = NewTrds()
	top = NewTrds()
	for _, cur := range *x {
		if cur.ClsRsn > v {
			*top = append(*top, cur)
		} else {
			*btm = append(*btm, cur)
		}
	}
	return btm, top
}
func (x *Trds) SelClsSpdEql(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.ClsSpd.Eql(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelClsSpdNeq(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.ClsSpd.Neq(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelClsSpdLss(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.ClsSpd.Lss(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelClsSpdGtr(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.ClsSpd.Gtr(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelClsSpdLeq(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.ClsSpd.Leq(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelClsSpdGeq(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.ClsSpd.Geq(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelClsSpdSplt(v flt.Flt) (btm, top *Trds) {
	btm = NewTrds()
	top = NewTrds()
	for _, cur := range *x {
		if cur.ClsSpd > v {
			*top = append(*top, cur)
		} else {
			*btm = append(*btm, cur)
		}
	}
	return btm, top
}
func (x *Trds) SelOpnSpdEql(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.OpnSpd.Eql(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelOpnSpdNeq(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.OpnSpd.Neq(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelOpnSpdLss(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.OpnSpd.Lss(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelOpnSpdGtr(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.OpnSpd.Gtr(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelOpnSpdLeq(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.OpnSpd.Leq(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelOpnSpdGeq(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.OpnSpd.Geq(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelOpnSpdSplt(v flt.Flt) (btm, top *Trds) {
	btm = NewTrds()
	top = NewTrds()
	for _, cur := range *x {
		if cur.OpnSpd > v {
			*top = append(*top, cur)
		} else {
			*btm = append(*btm, cur)
		}
	}
	return btm, top
}
func (x *Trds) SelClsAskEql(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.ClsAsk.Eql(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelClsAskNeq(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.ClsAsk.Neq(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelClsAskLss(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.ClsAsk.Lss(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelClsAskGtr(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.ClsAsk.Gtr(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelClsAskLeq(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.ClsAsk.Leq(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelClsAskGeq(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.ClsAsk.Geq(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelClsAskSplt(v flt.Flt) (btm, top *Trds) {
	btm = NewTrds()
	top = NewTrds()
	for _, cur := range *x {
		if cur.ClsAsk > v {
			*top = append(*top, cur)
		} else {
			*btm = append(*btm, cur)
		}
	}
	return btm, top
}
func (x *Trds) SelOpnAskEql(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.OpnAsk.Eql(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelOpnAskNeq(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.OpnAsk.Neq(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelOpnAskLss(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.OpnAsk.Lss(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelOpnAskGtr(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.OpnAsk.Gtr(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelOpnAskLeq(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.OpnAsk.Leq(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelOpnAskGeq(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.OpnAsk.Geq(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelOpnAskSplt(v flt.Flt) (btm, top *Trds) {
	btm = NewTrds()
	top = NewTrds()
	for _, cur := range *x {
		if cur.OpnAsk > v {
			*top = append(*top, cur)
		} else {
			*btm = append(*btm, cur)
		}
	}
	return btm, top
}
func (x *Trds) SelClsBidEql(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.ClsBid.Eql(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelClsBidNeq(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.ClsBid.Neq(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelClsBidLss(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.ClsBid.Lss(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelClsBidGtr(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.ClsBid.Gtr(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelClsBidLeq(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.ClsBid.Leq(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelClsBidGeq(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.ClsBid.Geq(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelClsBidSplt(v flt.Flt) (btm, top *Trds) {
	btm = NewTrds()
	top = NewTrds()
	for _, cur := range *x {
		if cur.ClsBid > v {
			*top = append(*top, cur)
		} else {
			*btm = append(*btm, cur)
		}
	}
	return btm, top
}
func (x *Trds) SelOpnBidEql(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.OpnBid.Eql(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelOpnBidNeq(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.OpnBid.Neq(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelOpnBidLss(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.OpnBid.Lss(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelOpnBidGtr(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.OpnBid.Gtr(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelOpnBidLeq(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.OpnBid.Leq(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelOpnBidGeq(v flt.Flt) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.OpnBid.Geq(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelOpnBidSplt(v flt.Flt) (btm, top *Trds) {
	btm = NewTrds()
	top = NewTrds()
	for _, cur := range *x {
		if cur.OpnBid > v {
			*top = append(*top, cur)
		} else {
			*btm = append(*btm, cur)
		}
	}
	return btm, top
}
func (x *Trds) SelClsTmeEql(v tme.Tme) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.ClsTme.Eql(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelClsTmeNeq(v tme.Tme) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.ClsTme.Neq(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelClsTmeLss(v tme.Tme) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.ClsTme.Lss(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelClsTmeGtr(v tme.Tme) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.ClsTme.Gtr(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelClsTmeLeq(v tme.Tme) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.ClsTme.Leq(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelClsTmeGeq(v tme.Tme) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.ClsTme.Geq(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelClsTmeSplt(v tme.Tme) (btm, top *Trds) {
	btm = NewTrds()
	top = NewTrds()
	for _, cur := range *x {
		if cur.ClsTme > v {
			*top = append(*top, cur)
		} else {
			*btm = append(*btm, cur)
		}
	}
	return btm, top
}
func (x *Trds) SelOpnTmeEql(v tme.Tme) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.OpnTme.Eql(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelOpnTmeNeq(v tme.Tme) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.OpnTme.Neq(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelOpnTmeLss(v tme.Tme) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.OpnTme.Lss(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelOpnTmeGtr(v tme.Tme) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.OpnTme.Gtr(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelOpnTmeLeq(v tme.Tme) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.OpnTme.Leq(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelOpnTmeGeq(v tme.Tme) (r *Trds) {
	r = NewTrds()
	for _, cur := range *x {
		if cur.OpnTme.Geq(v) {
			*r = append(*r, cur)
		}
	}
	return r
}
func (x *Trds) SelOpnTmeSplt(v tme.Tme) (btm, top *Trds) {
	btm = NewTrds()
	top = NewTrds()
	for _, cur := range *x {
		if cur.OpnTme > v {
			*top = append(*top, cur)
		} else {
			*btm = append(*btm, cur)
		}
	}
	return btm, top
}
func (x *Trds) OpnTmes() (r *tmes.Tmes) {
	r = tmes.Make(x.Cnt())
	for n, v := range *x {
		(*r)[n] = v.OpnTme
	}
	return r
}
func (x *Trds) ClsTmes() (r *tmes.Tmes) {
	r = tmes.Make(x.Cnt())
	for n, v := range *x {
		(*r)[n] = v.ClsTme
	}
	return r
}
func (x *Trds) OpnBids() (r *flts.Flts) {
	r = flts.Make(x.Cnt())
	for n, v := range *x {
		(*r)[n] = v.OpnBid
	}
	return r
}
func (x *Trds) ClsBids() (r *flts.Flts) {
	r = flts.Make(x.Cnt())
	for n, v := range *x {
		(*r)[n] = v.ClsBid
	}
	return r
}
func (x *Trds) OpnAsks() (r *flts.Flts) {
	r = flts.Make(x.Cnt())
	for n, v := range *x {
		(*r)[n] = v.OpnAsk
	}
	return r
}
func (x *Trds) ClsAsks() (r *flts.Flts) {
	r = flts.Make(x.Cnt())
	for n, v := range *x {
		(*r)[n] = v.ClsAsk
	}
	return r
}
func (x *Trds) OpnSpds() (r *flts.Flts) {
	r = flts.Make(x.Cnt())
	for n, v := range *x {
		(*r)[n] = v.OpnSpd
	}
	return r
}
func (x *Trds) ClsSpds() (r *flts.Flts) {
	r = flts.Make(x.Cnt())
	for n, v := range *x {
		(*r)[n] = v.ClsSpd
	}
	return r
}
func (x *Trds) ClsRsns() (r *strs.Strs) {
	r = strs.Make(x.Cnt())
	for n, v := range *x {
		(*r)[n] = v.ClsRsn
	}
	return r
}
func (x *Trds) Pips() (r *flts.Flts) {
	r = flts.Make(x.Cnt())
	for n, v := range *x {
		(*r)[n] = v.Pip
	}
	return r
}
func (x *Trds) Durs() (r *tmes.Tmes) {
	r = tmes.Make(x.Cnt())
	for n, v := range *x {
		(*r)[n] = v.Dur
	}
	return r
}
func (x *Trds) IsLongs() (r *bols.Bols) {
	r = bols.Make(x.Cnt())
	for n, v := range *x {
		(*r)[n] = v.IsLong
	}
	return r
}
func (x *Trds) PnlPcts() (r *flts.Flts) {
	r = flts.Make(x.Cnt())
	for n, v := range *x {
		(*r)[n] = v.PnlPct
	}
	return r
}
func (x *Trds) PnlPctPredicts() (r *flts.Flts) {
	r = flts.Make(x.Cnt())
	for n, v := range *x {
		(*r)[n] = v.PnlPctPredict
	}
	return r
}
func (x *Trds) PnlUsds() (r *flts.Flts) {
	r = flts.Make(x.Cnt())
	for n, v := range *x {
		(*r)[n] = v.PnlUsd
	}
	return r
}
func (x *Trds) PnlGrsUsds() (r *flts.Flts) {
	r = flts.Make(x.Cnt())
	for n, v := range *x {
		(*r)[n] = v.PnlGrsUsd
	}
	return r
}
func (x *Trds) CstComUsds() (r *flts.Flts) {
	r = flts.Make(x.Cnt())
	for n, v := range *x {
		(*r)[n] = v.CstComUsd
	}
	return r
}
func (x *Trds) CstClsSpdUsds() (r *flts.Flts) {
	r = flts.Make(x.Cnt())
	for n, v := range *x {
		(*r)[n] = v.CstClsSpdUsd
	}
	return r
}
func (x *Trds) CstOpnSpdUsds() (r *flts.Flts) {
	r = flts.Make(x.Cnt())
	for n, v := range *x {
		(*r)[n] = v.CstOpnSpdUsd
	}
	return r
}
func (x *Trds) OpnBalUsds() (r *flts.Flts) {
	r = flts.Make(x.Cnt())
	for n, v := range *x {
		(*r)[n] = v.OpnBalUsd
	}
	return r
}
func (x *Trds) ClsBalUsds() (r *flts.Flts) {
	r = flts.Make(x.Cnt())
	for n, v := range *x {
		(*r)[n] = v.ClsBalUsd
	}
	return r
}
func (x *Trds) ClsBalUsdActs() (r *flts.Flts) {
	r = flts.Make(x.Cnt())
	for n, v := range *x {
		(*r)[n] = v.ClsBalUsdAct
	}
	return r
}
func (x *Trds) TrdPcts() (r *flts.Flts) {
	r = flts.Make(x.Cnt())
	for n, v := range *x {
		(*r)[n] = v.TrdPct
	}
	return r
}
func (x *Trds) MrgnRtios() (r *flts.Flts) {
	r = flts.Make(x.Cnt())
	for n, v := range *x {
		(*r)[n] = v.MrgnRtio
	}
	return r
}
func (x *Trds) Unitss() (r *flts.Flts) {
	r = flts.Make(x.Cnt())
	for n, v := range *x {
		(*r)[n] = v.Units
	}
	return r
}
func (x *Trds) Instrs() (r *strs.Strs) {
	r = strs.Make(x.Cnt())
	for n, v := range *x {
		(*r)[n] = v.Instr
	}
	return r
}
func (x *Trds) OpnReqs() (r *strs.Strs) {
	r = strs.Make(x.Cnt())
	for n, v := range *x {
		(*r)[n] = v.OpnReq
	}
	return r
}
func (x *Trds) OpnRess() (r *strs.Strs) {
	r = strs.Make(x.Cnt())
	for n, v := range *x {
		(*r)[n] = v.OpnRes
	}
	return r
}
func (x *Trds) ClsReqs() (r *strs.Strs) {
	r = strs.Make(x.Cnt())
	for n, v := range *x {
		(*r)[n] = v.ClsReq
	}
	return r
}
func (x *Trds) ClsRess() (r *strs.Strs) {
	r = strs.Make(x.Cnt())
	for n, v := range *x {
		(*r)[n] = v.ClsRes
	}
	return r
}
func (x *Trds) StrWrt(b *strings.Builder) {
	b.WriteRune('[')
	b.WriteRune('\n')
	for n, v := range *x {
		v.StrWrt(b)
		if n != len(*x)-1 {
			b.WriteRune('\n')
		}
	}
	b.WriteRune('\n')
	b.WriteRune(']')
}
func (x *Trds) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
