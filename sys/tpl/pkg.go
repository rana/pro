package tpl

import (
	"fmt"
	"path"
	"strings"
)

type (
	Pkg struct {
		Lbl
		Fles
		Cnsts  CnstMap
		Vars   VarMap
		PkgFns PkgFnMap
		Pth    string
		Trm    *TypFn
		Xpr    *Struct
		Act    *Struct
	}
	Pkgs []*Pkg
)

func (x *Pkgs) Ok() bool { return len(*x) != 0 }
func (x *Pkgs) HasPkg(v *Pkg) bool {
	for _, pkg := range *x {
		if v == pkg {
			return true
		}
	}
	return false
}
func (x *Pkgs) AddPkg(vs ...*Pkg) { *x = append(*x, vs...) }
func NewPkg(pth string, test ...bool) (r *Pkg) {
	r = &Pkg{}
	r.Cnsts = make(CnstMap)
	r.Vars = make(VarMap)
	r.PkgFns = make(PkgFnMap)
	r.Pth = strings.ToLower(pth)
	idx := strings.LastIndex(pth, "/")
	if idx > -1 {
		r.Name = pth[idx+1:]
	} else {
		r.Name = pth
	}
	if len(test) > 0 && test[0] {
		r.Name += TestSuffix
	}
	return r
}
func (x *Pkg) New(name string) (r *Pkg) {
	return NewPkg(fmt.Sprintf("%v/%v", x.Pth, name))
}
func (x *Pkg) NewFromPrnt(name string) (r *Pkg) {
	return NewPkg(fmt.Sprintf("%v/%v", path.Dir(x.Pth), name))
}
func (x *Pkg) ParntName() string {
	idx := strings.LastIndex(x.Pth, "/")
	if idx < 0 {
		return ""
	}
	return x.Pth[idx+1:]
}
func (x *Pkg) Ref(f Fle) string {
	for _, i := range f.Bse().Imports {
		if x.Pth == i.Pth && i.Alias != "" {
			return i.Alias
		}
	}
	return x.Name
}
func (x *Pkg) String() string              { return x.Name }
func (x *Pkg) NewTest() (r *Pkg)           { return NewPkg(x.Pth, true) }
func (x *Pkg) IsTest() bool                { return strings.HasSuffix(x.Name, TestSuffix) }
func (x *Pkg) WritePkg(b *strings.Builder) { b.WriteString(fmt.Sprintf("package %v\n\n", x.Name)) }
func (x *Pkg) AddFle(f Fle) {
	x.Fles = append(x.Fles, f)
}
