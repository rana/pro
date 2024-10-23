package tpl

import (
	"fmt"
	"reflect"
	"strings"
	"sys/err"
)

type (
	Import struct {
		Pth   string
		Alias string
	}
	Imports []*Import
)

func (x *Imports) Ok() bool                { return len(*x) != 0 }
func (x *Imports) Cnt() int                { return len(*x) }
func (x *Imports) AddImport(vs ...*Import) { *x = append(*x, vs...) }
func (x *Imports) Import(v interface{}, alias ...string) {
	switch cur := v.(type) {
	case Fle:
		x.ImportPth(cur.Bse().Pkg.Pth, alias...)
	case Dir:
		x.ImportPth(cur.DirPkg().Pth, alias...)
	case *Pkg:
		x.ImportPth(cur.Pth, alias...)
	case *PkgFn:
		x.ImportPth(cur.Pkg.Pth, alias...)
	case Typ:
		// fmt.Println(">>>", cur)
		if cur.Bse().Pkg != nil { // nil for golang basic types
			x.ImportPth(cur.Bse().Pkg.Pth, alias...)
		}
	case string:
		x.ImportPth(cur, alias...)
	default:
		err.Panicf("import: unsupported (v:%v)", v)
	}
}
func (x *Imports) ImportAlias(v interface{}) {
	var pth string
	switch cur := v.(type) {
	case Fle:
		pth = cur.Bse().Pkg.Pth
	case Dir:
		pth = cur.DirPkg().Pth
	case *Pkg:
		pth = cur.Pth
	case *PkgFn:
		pth = cur.Pkg.Pth
	case Typ:
		pth = cur.Bse().Pkg.Pth
	case string:
		pth = cur
	default:
		err.Panicf("import: unsupported (v:%v)", v)
	}
	s := strings.Split(pth, "/")
	if len(s) > 1 {
		x.Import(pth, fmt.Sprintf("%v%v", s[len(s)-2], s[len(s)-1]))
	} else {
		x.Import(pth)
	}
}
func (x *Imports) ImportPth(pth string, alias ...string) {
	if pth == "" { // don't add empty pth; may be empty from Ext
		return
	}
	for _, i := range *x {
		if i.Pth == pth { // don't add duplicate pth
			return
		}
	}
	r := &Import{}
	r.Pth = pth
	if len(alias) > 0 {
		r.Alias = alias[0]
	}
	x.AddImport(r)
}
func (x *Imports) Importf(format string, args ...interface{}) {
	x.Import(fmt.Sprintf(format, args...))
}
func (x *Imports) ImportTyp(typ Typ, fle ...Fle) {
	if _, isExt := typ.(*Ext); isExt && typ.Bse().Pkg == _sys.Pkg {
		return
	}
	if typ != nil && !reflect.ValueOf(typ).IsNil() {
		if typ.Bse().Pkg == nil {
			return
		}
		if len(fle) != 0 && fle[0].Bse().Pkg != nil && fle[0].Bse().Pkg.Pth == typ.Bse().Pkg.Pth {
			return
		}
		x.Import(typ.Bse().Pkg.Pth)
	}
}
func (x *Imports) ImportPrm(prm *Prm, fle ...Fle) {
	if prm == nil {
		return
	}
	x.ImportTyp(prm.Typ, fle...)
}
func (x *Imports) ImportFn(fn Fn, fle ...Fle) {
	if fn == nil {
		return
	}
	for _, prm := range fn.In() {
		x.ImportPrm(prm, fle...)
	}
	for _, prm := range fn.Out() {
		x.ImportPrm(prm, fle...)
	}
}
func (x *Imports) WriteImports(b *strings.Builder, f *FleBse) {
	if len(*x) > 0 {
		b.WriteString("import (\n")
		for _, i := range *x {
			// if i.Pth != f.Pkg.Pth { // don't import self
			i.WriteImport(b)
			// }
		}
		b.WriteString(")\n\n")
	}
}

func (x *Import) WriteImport(b *strings.Builder) {
	if x.Alias != "" {
		b.WriteString(x.Alias)
		b.WriteRune(' ')
	}
	b.WriteRune('"')
	b.WriteString(x.Pth)
	b.WriteRune('"')
	b.WriteRune('\n')
}
