package tpl

import (
	"fmt"
	"strings"
	"sys"
)

type (
	PkgFn struct {
		FnBse
		Pkg *Pkg
		Block
		B    *strings.Builder
		Node *Struct
	}
	PkgFns   []*PkgFn
	PkgFnMap map[string]*PkgFn
)

func (x *PkgFns) Ok() bool              { return len(*x) != 0 }
func (x *PkgFns) Cnt() int              { return len(*x) }
func (x *PkgFns) AddPkgFn(vs ...*PkgFn) { *x = append(*x, vs...) }
func (x *PkgFns) TypRefs() (r Typs) {
	for _, fn := range *x {
		r = append(r, fn.TypRefs()...)
	}
	return r
}
func (x *PkgFns) MayXpr() bool { // at least one
	for _, v := range *x {
		if v.MayXpr() {
			return true
		}
	}
	return false
}
func (x *PkgFns) WritePkgFns(b *strings.Builder, f *FleBse) {
	for _, fn := range *x {
		fn.WritePkgFn(b, f)
	}
}

func (x *PkgFn) Full() string {
	return fmt.Sprintf("%v.%v", x.Pkg, x.Name)
}
func (x *PkgFn) PkgTitle() string {
	return fmt.Sprintf("%v%v", x.Pkg.Title(), x.Title())
}

func (x *PkgFn) TestCall(f Fle, b *strings.Builder, camel ...bool) string {
	b.Reset()
	b.WriteString(x.Ref(f, camel...))
	b.WriteRune('(')
	x.InPrms.TestCall(b)
	b.WriteRune(')')
	return b.String()
}

func (x *PkgFn) Ref(f Fle, camel ...bool) string {
	b := &strings.Builder{}
	bse := f.Bse()
	if bse.Pkg != x.Pkg {
		f.Bse().Import(x)
		if len(bse.Imports) == 0 {
			b.WriteString(x.Full())
		} else {
			pkgPth := x.Pkg.Pth
			var imp *Import
			for _, cur := range bse.Imports {
				if pkgPth == cur.Pth && cur.Alias != "" {
					imp = cur
					break
				}
			}
			var name string
			if len(camel) != 0 && camel[0] {
				name = sys.Camel(x.Name)
			} else {
				name = x.Name
			}
			if imp != nil {
				b.WriteString(imp.Alias)
			} else {
				b.WriteString(x.Pkg.Name)
			}
			b.WriteRune('.')
			b.WriteString(name)
		}
	} else {
		b.WriteString(x.Name)
	}
	return b.String()
}
func (x *PkgFn) WritePkgFn(b *strings.Builder, f *FleBse) {
	b.WriteString("func ")
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
	x.WriteBlock(b)
}
func (x *PkgFn) AddStrt(txt string) {
	x.B = &strings.Builder{}
	x.B.WriteString(txt)
}
func (x *PkgFn) AddStrtf(format string, args ...interface{}) {
	x.AddStrt(fmt.Sprintf(format, args...))
}
func (x *PkgFn) AddMdl(txt string) {
	x.B.WriteString(txt)
}
func (x *PkgFn) AddMdlf(format string, args ...interface{}) {
	x.AddMdl(fmt.Sprintf(format, args...))
}
func (x *PkgFn) AddEnd(txt ...string) {
	if len(txt) != 0 {
		x.B.WriteString(txt[0])
	}
	x.Add(x.B.String())
	x.B.Reset()
}
func (x *PkgFn) AddEndf(format string, args ...interface{}) {
	x.AddEnd(fmt.Sprintf(format, args...))
}

func (x *PkgFn) Call(rxrName ...string) string {
	var b strings.Builder
	x.CallWrt(&b)
	return b.String()
}
func (x *PkgFn) CallLit() string {
	var b strings.Builder
	x.CallLitWrt(&b)
	return b.String()
}
func (x *PkgFn) CallNode() string {
	var b strings.Builder
	x.CallNodeWrt(&b)
	return b.String()
}

func (x *PkgFn) CallWrt(b *strings.Builder, rxrName ...string) {
	b.WriteString(x.Pkg.Name)
	x.FnBse.CallWrt(b)
}
func (x *PkgFn) CallLitWrt(b *strings.Builder) {
	b.WriteString(x.Pkg.Name)
	x.FnBse.CallLitWrt(b)
}
func (x *PkgFn) CallValWrt(b *strings.Builder, f Fle) {
	b.WriteString(x.Pkg.Name)
	x.FnBse.CallValWrt(b, f)
}
func (x *PkgFn) CallNodeWrt(b *strings.Builder) {
	b.WriteString(x.Pkg.Name)
	x.FnBse.CallNodeWrt(b)
}
