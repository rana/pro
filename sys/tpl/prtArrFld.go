package tpl

import (
	"sys"
)

type (
	PrtArrFld struct {
		PrtBse
	}
)

func (x *PrtArrFld) InitPrtTypFn() {
	arr := x.f.GetPrt((*PrtArr)(nil)).(*PrtArr).Arr
	if elm, ok := arr.Alias.Elm.(*Struct); ok {
		for _, fld := range elm.Flds {
			fldBse := fld.Typ.Bse()
			if fld.IsFstUpr() && fldBse.IsAct() && fldBse.arr != nil {
				x.Fld(fld)
			}
		}
	}
}
func (x *PrtArrFld) Fld(fld *Fld) (r *TypFn) {
	arr := fld.Typ.Bse().arr
	r = x.f.TypFn(sys.Plural(fld.Name))
	r.OutPrm(arr, "r")
	r.Addf("r = %v(x.Cnt())", arr.Make.Ref(x.f))
	r.Add("for n, v := range *x {")
	r.Addf("(*r)[n] = v.%v", fld.Name)
	r.Add("}")
	r.Add("return r")
	// test
	if x.f.Test != nil && Opt.IsArr() {
		r.T.SkpCpy = true
		r.T.Addf("expected :=  %v(x.Cnt())", arr.Make.Ref(x.f.Test))
		r.T.Add("for n, v := range *x {")
		r.T.Addf("expected.Upd(%v(n), v.%v)", _sys.Bsc.Unt.Typ().Ref(x.f.Test), fld.Name)
		r.T.Add("}")
	}
	return r
}
