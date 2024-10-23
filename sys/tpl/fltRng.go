package tpl

import (
	"sys"
	"sys/k"
	"sys/tpl/atr"
)

type (
	FleFltRng struct {
		FleBse
		// PrtStructIdn
		PrtRng
		PrtString
		PrtBytes
		// PrtLog
		// PrtIfc
		Elm *FleFlt
	}
)

func (x *DirBsc) NewFltRng() (r *FleFltRng) {
	r = &FleFltRng{}
	x.FltRng = r
	r.Name = k.Rng
	r.Elm = x.Flt
	r.PrtRng.Elm = r.Elm.Bse()
	r.Pkg = r.Elm.Pkg
	r.Structf(r.Name, atr.TypRng)
	r.AddFle(r)
	return r
}
func (x *FleFltRng) InitVals(bse *TypBse) {
	bse.Lits = sys.Vs("0.0-0.0", "0.0-1.0", "-3.0--4.0", "-999.0-1000.0", "1.0-0.0")
	bse.Vals = sys.VsStruct(x.Typ().Full(), "Min:0, Max:0", "Min:0, Max:1", "Min:-3, Max:-4", "Min:-999, Max:1000", "Min:1, Max:0")
}
func (x *FleFltRng) InitTrm(bse *TypBse, trmr *FleTrmr) {
	x.InitLexLit(func(r *TypFn, t *FleTrmr) {

		r.Add("r.Idx = x.Idx")

		r.Addf("r.MinTrm, ok = x.%v(true)", _sys.Bsc.Flt.Typ().Bse().LitTrm.Name)
		r.Add("if !ok {")
		r.Add("return r, false")
		r.Add("}")

		r.Add("if x.Ch != '-' {")
		r.Add("return r, false")
		r.Add("}")
		r.Add("x.NextRune()")

		r.Addf("r.MaxTrm, ok = x.%v()", _sys.Bsc.Flt.Typ().Bse().LitTrm.Name)
		r.Add("if !ok {")
		r.Add("return r, false")
		r.Add("}")

		r.Add("r.Lim = x.Idx")
		r.Add("return r, true")

	}, trmr.StructRng(x.Elm.Typ().Bse()))
	x.InitPrsLit(func(r *PkgFn, f *FlePrs) {
		r.Addf("r.Min = %v(trm.MinTrm, txt)", x.Elm.Typ().Bse().PrsTrm.Ref(f))
		r.Addf("r.Max = %v(trm.MaxTrm, txt)", x.Elm.Typ().Bse().PrsTrm.Ref(f))
		r.Add("return r")
	})
	x.InitPrsCfg()
}
