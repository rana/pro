package tpl

import (
	"fmt"
	"strings"
	"sys"
	"sys/tpl/atr"
)

type (
	Var struct {
		Lbl
		atr.Atr
		Pkg    *Pkg
		Typ    Typ
		Value  string
		FnCall bool
		Trm    *TypFn
		Xpr    *Struct
		Act    *Struct
	}
	Vars   []*Var
	VarMap map[string]*Var
)

func (x *Vars) Ok() bool          { return len(*x) != 0 }
func (x *Vars) Cnt() int          { return len(*x) }
func (x *Vars) AddVar(vs ...*Var) { *x = append(*x, vs...) }
func (x *Vars) Var(name, value string, typ Typ, pkg *Pkg, a atr.Atr) (r *Var) {
	r = &Var{}
	r.Name = strings.Title(name)
	r.Pkg = pkg
	r.Typ = typ
	r.Value = value
	r.Atr = a
	x.AddVar(r)
	return r
}
func (x *Var) Full() string { return fmt.Sprintf("%v.%v", x.Pkg.Name, x.Name) }
func (x *Var) Ref(f Fle, camel ...bool) string {
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
func (x *Var) MayXpr() bool {
	if x.Typ == Interface {
		return true
	}
	if !x.IsXpr() || !x.Typ.Bse().IsXpr() {
		return false
	}
	_, ok := x.Typ.(*Ext)
	return !ok && x.Typ.Bse().IsXpr()
}

func (x *Vars) TypRefs() (r Typs) {
	for _, f := range *x {
		if f.Typ != nil {
			r = append(r, f.Typ)
		}
	}
	return r
}
func (x *Vars) MayXpr() bool { // at least one
	for _, v := range *x {
		if v.MayXpr() {
			return true
		}
	}
	return false
}

func (x *Vars) WriteVarDecls(b *strings.Builder, f *FleBse) {
	if len(*x) > 0 {
		b.WriteString("var (\n")
		for _, v := range *x {
			b.WriteString(v.Name)
			if v.IsEmpty() {
				b.WriteString(" ")
				b.WriteString(v.Typ.Ref(f))
			} else {
				if v.FnCall {
					b.WriteString(" = ")
					b.WriteString(v.Value)
				} else {
					b.WriteString(" = ")
					b.WriteString(v.Typ.Ref(f))
					if v.Typ.Bse().IsStruct() {
						b.WriteString("{")
						b.WriteString(v.Value)
						b.WriteString("}")
					} else {
						b.WriteString("(")
						b.WriteString(v.Value)
						b.WriteString(")")
					}
				}
			}
			b.WriteString("\n")
		}
		b.WriteString(")\n\n")
	}
}
