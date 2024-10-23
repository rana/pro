package tpl

type (
	PrtStructFldSet struct {
		PrtBse
	}
)

func (x *PrtStructFldSet) InitPrtTypFn() {
	x.evalFlds()
}
func (x *PrtStructFldSet) evalFlds() {
	rxr := x.f.Typ().(*Struct)
	for _, fld := range rxr.Flds {
		if fld.IsPrtFldSet() {
			x.fldSet(fld)
		}
	}
}
func (x *PrtStructFldSet) fldSet(fld *Fld) (r *TypFn) {
	r = x.f.TypFn(fld.Name)
	r.OutPrm(x.t)
	if fld.IsSlice() {
		r.InPrmVariadic(fld.Typ, "vs")
		r.Addf("x.%[1]v = append(x.%[1]v, vs...)", fld.Name)
	} else {
		r.InPrm(fld.Typ, "v")
		r.Addf("x.%v = v", fld.Name)
	}
	r.Add("return x")
	return r
}
