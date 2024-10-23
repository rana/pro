package tpl

import (
	"strings"
)

type (
	MemSig struct {
		FnBse
		Rxr *Ifc
	}
	MemSigs   []*MemSig
	MemSigMap map[string]*MemSig
)

func (x *MemSigs) Ok() bool             { return len(*x) != 0 }
func (x *MemSigs) Cnt() int             { return len(*x) }
func (x *MemSigs) AddSig(vs ...*MemSig) { *x = append(*x, vs...) }
func (x *MemSigs) TypRefs() (r Typs) {
	for _, s := range *x {
		r = append(r, s.TypRefs()...)
	}
	return r
}
func (x *MemSigs) MayXpr() bool { // at least one
	for _, v := range *x {
		if v.MayXpr() {
			return true
		}
	}
	return false
}
func (x *MemSigs) Dque() (r *MemSig) {
	r = (*x)[0]
	if len(*x) == 1 {
		*x = (*x)[:0]
	} else {
		copy(*x, (*x)[1:])
		*x = (*x)[:len(*x)-1]
	}
	return r
}

func (x *MemSigs) WriteMemSigs(b *strings.Builder, f *FleBse) {
	for _, s := range *x {
		s.WriteMemSig(b, f)
	}
}

// MemSig
// func (x *MemSig) MayXpr() bool {
// 	// return x.Rxr.Bse().IsXpr() && x.FnBse.MayXpr()
// 	return x.FnBse.MayXpr()
// }
func (x *MemSig) TypRefs() (r Typs) {
	r = append(r, x.InPrms.TypRefs()...)
	r = append(r, x.OutPrms.TypRefs()...)
	return r
}
func (x *MemSig) WriteMemSig(b *strings.Builder, f *FleBse) {
	b.WriteString(x.Name)
	if len(x.InPrms) > 0 {
		x.InPrms.WriteInPrms(b, f)
	} else {
		b.WriteString("()")
	}
	if len(x.OutPrms) > 0 {
		b.WriteRune(' ')
		x.OutPrms.WriteOutPrms(b, f)
	}
	b.WriteRune('\n')
}
