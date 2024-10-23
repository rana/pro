package tpl

import (
	"strings"
	"sys/tpl/atr"
)

type (
	Fn interface {
		Lblr
		In() InPrms
		Out() OutPrms
		CallWrt(b *strings.Builder, rxrName ...string)
		Call(rxrName ...string) string //rxrName ...string
		Bse() *FnBse
	}
	FnBse struct {
		Lbl
		InPrms
		OutPrms
		atr.Atr
		Fle    *FleBse
		Alias  string
		Family string
		Cnj    string
		T      FnT
		T2     FnT2
		Trm    *TypFn
		Xpr    *Struct
		Act    *Struct
	}
	FnBses []*FnBse
)

func (x *FnBses) Ok() bool              { return len(*x) != 0 }
func (x *FnBses) Cnt() int              { return len(*x) }
func (x *FnBses) AddTypFn(vs ...*FnBse) { *x = append(*x, vs...) }

func (x *FnBse) Bse() *FnBse { return x }
func (x *FnBse) MayXpr() bool {
	return x.IsXpr() && x.InPrms.MayXpr() && x.OutPrms.MayXpr()
}
func (x *FnBse) PrmsMayXpr() bool {
	return x.InPrms.MayXpr() && x.OutPrms.MayXpr()
}
func (x *FnBse) In() InPrms   { return x.InPrms }
func (x *FnBse) Out() OutPrms { return x.OutPrms }
func (x *FnBse) OutTyp() Typ  { return x.OutPrms[0].Typ }
func (x *FnBse) TypRefs() (r Typs) {
	r = append(r, x.InPrms.TypRefs()...)
	r = append(r, x.OutPrms.TypRefs()...)
	return r
}
func (x *FnBse) CallWrt(b *strings.Builder, rxrName ...string) {
	b.WriteRune('.')
	b.WriteString(x.Name)
	b.WriteRune('(')
	x.InPrms.CallWrt(b)
	b.WriteRune(')')
}
func (x *FnBse) CallLitWrt(b *strings.Builder) {
	b.WriteRune('.')
	b.WriteString(x.Name)
	b.WriteRune('(')
	x.InPrms.CallLitWrt(b)
	b.WriteRune(')')
}
func (x *FnBse) CallValWrt(b *strings.Builder, f Fle) {
	b.WriteRune('.')
	b.WriteString(x.Name)
	b.WriteRune('(')
	x.InPrms.CallValWrt(b, f)
	b.WriteRune(')')
}
func (x *FnBse) CallNodeWrt(b *strings.Builder) {
	b.WriteRune('.')
	b.WriteString(x.Name)
	b.WriteRune('(')
	x.InPrms.CallNodeWrt(b)
	b.WriteRune(')')
}
func (x *FnBse) Call(rxrName ...string) string {
	var b strings.Builder
	x.CallWrt(&b)
	return b.String()
}
func (x *FnBse) CallLit() string {
	var b strings.Builder
	x.CallLitWrt(&b)
	return b.String()
}
func (x *FnBse) CallVal(f Fle) string {
	var b strings.Builder
	x.CallValWrt(&b, f)
	return b.String()
}
func (x *FnBse) CallNode() string {
	var b strings.Builder
	x.CallNodeWrt(&b)
	return b.String()
}
