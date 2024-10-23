package tpl

import (
	"strings"
	"sys/k"
	"sys/ks"
)

type (
	PrtArrFldSel struct {
		PrtBse
		Arr *Arr
		Elm *Struct
		Ifc *Ifc
	}
)

func (x *PrtArrFldSel) InitPrtTyp() {
	x.Arr = x.f.GetPrt((*PrtArr)(nil)).(*PrtArr).Arr
	if s, ok := x.Arr.Alias.Elm.(*Struct); ok {
		x.Elm = s
	} else if ifc, ok := x.Arr.Alias.Elm.(*Ifc); ok && ifc.bse != nil {
		x.Elm = ifc.bse
		x.Ifc = ifc
	}
}
func (x *PrtArrFldSel) InitPrtTypFn() {
	if x.Elm != nil {
		var bse string
		if x.Ifc != nil {
			bse = ".Bse()"
		}
		flds := NewFlds(x.Elm.Flds...)
		for flds.Cnt() != 0 {
			fld := flds.Pop()
			if fld.IsSlf() || fld.IsPrnt() {
				continue
			}
			if fld.Name == "" { // embedded typ
				if s, ok := fld.Typ.(*Struct); ok {
					flds.Push(s.Flds...) // add embedded typ fields
				}
				continue
			}
			fldBse := fld.Typ.Bse()
			if fld.IsFstUpr() && !fld.IsSelSkp() && fldBse.IsAct() {
				if fldBse.IsIdn() { // supports bol.Bol
					for _, op := range ks.Idns {
						x.Fld(fld, op, bse)
					}
				}
				if fldBse.IsRel() {
					for _, op := range ks.Rels {
						x.Fld(fld, op, bse)
					}
					x.Splt(fld, bse)
				}
			}
		}
	}
}
func (x *PrtArrFldSel) Fld(fld *Fld, op, bse string) (r *TypFn) {
	op = strings.Title(op)
	r = x.f.TypFnf("Sel%v%v", fld.Title(), op) // SelPipGtr(), SelIsLongEql()
	r.InPrm(fld.Typ, "v")
	r.OutPrm(x.t, "r")
	r.Addf("r = %v()", x.Arr.New.Ref(x.f))
	r.Add("for _, cur := range *x {")
	r.Addf("if cur%v.%v.%v(v) {", bse, fld.Name, op)
	r.Add("*r = append(*r, cur)")
	r.Add("}")
	r.Add("}")
	r.Add("return r")
	// test
	if x.f.Test != nil && Opt.IsArr() {
		r.T.SkpCpy = true
		r.T.Addf("expected := %v()", x.Arr.New.Ref(x.f.Test))
		r.T.Add("for _, cur := range *x {")
		r.T.Addf("if cur%v.%v.%v(v) {", bse, fld.Name, op)
		r.T.Add("expected.Push(cur)")
		r.T.Add("}")
		r.T.Add("}")
	}
	return r
}
func (x *PrtArrFldSel) Splt(fld *Fld, bse string) (r *TypFn) {
	r = x.f.TypFnf("Sel%v%v", fld.Title(), strings.Title(k.Splt)) // SelPipSplt()
	r.InPrm(fld.Typ, "v")
	r.OutPrm(x.t, "btm")
	r.OutPrm(x.t, "top")
	r.Addf("btm = %v()", x.Arr.New.Ref(x.f))
	r.Addf("top = %v()", x.Arr.New.Ref(x.f))
	r.Add("for _, cur := range *x {")
	r.Addf("if cur%v.%v > v {", bse, fld.Name)
	r.Add("*top = append(*top, cur)")
	r.Add("} else {")
	r.Add("*btm = append(*btm, cur)")
	r.Add("}")
	r.Add("}")
	r.Add("return btm, top")
	return r
}
