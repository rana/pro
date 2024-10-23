package tpl

import (
	"strings"
	"sys/k"
	"sys/tpl/atr"
)

type (
	PrtStructRel struct {
		PrtBse
		Cmp  *Func
		Eqls map[string]*PkgFn
		Lsss map[string]*PkgFn
		Gtrs map[string]*PkgFn
	}
)

func (x *PrtStructRel) InitPrtTyp() {
	name := k.Cmp
	if x.t.Pkg.Name != x.t.Camel() {
		if x.t.ifc != nil {
			name = x.t.ifc.Title() + strings.Title(name)
		} else {
			name = x.t.Title() + strings.Title(name)
		}
	}
	x.Cmp = x.f.Func(name, atr.None)
	x.Eqls = make(map[string]*PkgFn)
	x.Lsss = make(map[string]*PkgFn)
	x.Gtrs = make(map[string]*PkgFn)
}
func (x *PrtStructRel) InitPrtFld() {
	x.Cmp.InPrm(x.f.Typ(), "a")
	x.Cmp.InPrm(x.f.Typ(), "b")
	x.Cmp.OutPrm(_sys.Bsc.Bol)
}
func (x *PrtStructRel) InitPrtPkgFn() {
	s, ok := x.f.Typ().(*Struct)
	if !ok {
		s = x.t.bse
	}
	var prefix, bse string
	if s.Pkg.Title() != s.Title() {
		if s.ifc != nil {
			prefix = s.ifc.Title()
			bse = ".Bse()"
		} else {
			prefix = s.Title()
		}
	}
	flds := NewFlds(s.Flds...)
	for flds.Cnt() != 0 {
		fld := flds.Pop()
		if fld.Name == "" { // embedded typ
			if s, ok := fld.Typ.(*Struct); ok {
				flds.Push(s.Flds...) // add embedded typ fields
			}
			continue
		}
		fldBse := fld.Typ.Bse()
		if fld.IsFstUpr() && fldBse.IsRel() && fldBse.IsAct() {
			x.Eqls[fld.Name] = x.CmpFn(fld, prefix, bse, "==", "Eql", "")
			x.Lsss[fld.Name] = x.CmpFn(fld, prefix, bse, "<", "Lss", "Asc")
			x.Gtrs[fld.Name] = x.CmpFn(fld, prefix, bse, ">", "Gtr", "Dsc")
		}
	}
}
func (x *PrtStructRel) CmpFn(fld *Fld, prefix, bse, op, opName, alias string) (r *PkgFn) {
	op = strings.Title(op)
	r = x.f.PkgFnf("%v%v%v", prefix, fld.Title(), opName) // PipGtr(), PipLss()
	r.InPrm(x.f.Typ(), "a")
	r.InPrm(x.f.Typ(), "b")
	r.OutPrm(_sys.Bsc.Bol)
	r.Atr = atr.None
	r.Alias = alias
	r.Addf("return a%[1]v.%[2]v %v b%[1]v.%[2]v", bse, fld.Name, op)
	return r
}
