package tpl

import (
	"fmt"
	"strings"
	"sys"
	"sys/tpl/atr"
)

type (
	Cnst struct {
		Lbl
		atr.Atr
		Pkg   *Pkg
		Typ   Typ
		Value string
		Trm   *TypFn
		Xpr   *Struct
		Act   *Struct
	}
	Cnsts   []*Cnst
	CnstMap map[string]*Cnst
)

func (x *Cnsts) Ok() bool            { return len(*x) != 0 }
func (x *Cnsts) Cnt() int            { return len(*x) }
func (x *Cnsts) AddCnst(vs ...*Cnst) { *x = append(*x, vs...) }
func (x *Cnsts) Cnst(name, value string, typ Typ, pkg *Pkg, a atr.Atr) (r *Cnst) {
	r = &Cnst{}
	r.Name = strings.Title(name)
	r.Pkg = pkg
	r.Typ = typ
	r.Value = value
	r.Atr = a
	x.AddCnst(r)
	return r
}
func (x *Cnst) Full() string { return fmt.Sprintf("%v.%v", x.Pkg.Name, x.Name) }
func (x *Cnst) Ref(f Fle, camel ...bool) string {
	b := &strings.Builder{}
	bse := f.Bse()
	if bse.Pkg != x.Pkg {
		f.Bse().Import(x.Pkg)
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
				b.WriteString(fmt.Sprintf("%v.%v", imp.Alias, name))
			} else {
				b.WriteString(x.Pkg.Name)
				b.WriteRune('.')
				b.WriteString(name)
			}
		}
	} else {
		b.WriteString(x.Name)
	}
	return b.String()
}
func (x *Cnst) MayXpr() bool {
	if !x.IsXpr() || !x.Typ.Bse().IsXpr() {
		return false
	}
	_, ok := x.Typ.(*Ext)
	return !ok && x.Typ.Bse().IsXpr()
}

func (x *Cnsts) TypRefs() (r Typs) {
	for _, f := range *x {
		if f.Typ != nil {
			r = append(r, f.Typ)
		}
	}
	return r
}
func (x *Cnsts) MayXpr() bool { // at least one
	for _, v := range *x {
		if v.MayXpr() {
			return true
		}
	}
	return false
}

func (x *Cnsts) WriteCnstDecls(b *strings.Builder, f *FleBse, hasIota ...bool) {
	if len(*x) > 0 {
		b.WriteString("const (\n")

		hasIota := len(*x) > 0 && (*x)[0].Value == ""
		if hasIota {
			for n, c := range *x {
				b.WriteString(c.Name)
				if n == 0 && c.IsIota() {
					b.WriteRune(' ')
					b.WriteString(c.Typ.Ref(f))
					b.WriteString(" = ")
					if c.IsFlg() {
						b.WriteString("1<<")
					}
					b.WriteString("iota")
				} else if c.Value != "" {
					b.WriteRune(' ')
					b.WriteString(c.Typ.Ref(f))
					b.WriteString(" = ")
					b.WriteString(c.Value)
				}
				// if n == 0 && c.Value == "" {
				// 	b.WriteString(c.Name)
				// 	b.WriteRune(' ')
				// 	b.WriteString(c.Typ.Ref(f))
				// 	b.WriteString(" = iota")
				// } else {
				// 	b.WriteString(c.Name)
				// }

				b.WriteRune('\n')
			}
		} else {
			for _, c := range *x {
				b.WriteString(c.Name)
				b.WriteString(" = ")
				b.WriteString(c.Typ.Ref(f))
				b.WriteString("(")
				b.WriteString(c.Value)
				b.WriteString(")")
				b.WriteString("\n")
			}
		}
		b.WriteString(")\n\n")
	}
}
