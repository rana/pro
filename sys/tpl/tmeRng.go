package tpl

import (
	"sys"
	"sys/k"
	"sys/tpl/atr"
)

type (
	FleTmeRng struct {
		FleBse
		// PrtStructIdn
		PrtRng
		PrtString
		PrtBytes
		// PrtLog
		// PrtIfc
		Elm *FleTme
	}
	FleTmeRngs struct {
		FleBse
		PrtArr
		// PrtArrIdn
		PrtArrRng
		// PrtLog
		// PrtIfc
		PrtArrStrWrt
	}
)

func (x *DirBsc) NewTmeRng() (r *FleTmeRng) {
	r = &FleTmeRng{}
	x.TmeRng = r
	r.Name = k.Rng
	r.Elm = x.Tme
	r.PrtRng.Elm = r.Elm.Bse()
	r.Pkg = r.Elm.Pkg
	r.Structf(r.Name, atr.TypRng)
	r.AddFle(r)
	return r
}
func (x *FleTmeRng) NewArr() (r *FleTmeRngs) {
	r = &FleTmeRngs{}
	r.FleBse = *NewArr(x, &r.PrtArr, x.Pkg)
	r.AddFle(r)
	return r
}
func (x *FleTmeRng) InitVals(bse *TypBse) {
	// bse.Lits = sys.Vs("0s-0s", "0s-1s", "-3s--4s", "-999s-1000s", "1s-0s")
	// bse.Vals = sys.VsStruct(x.Typ().Full(), "Min:0, Max:0", "Min:0, Max:1", "Min:-3, Max:-4", "Min:-999, Max:1000", "Min:1, Max:0")
	// USE SIMPLE SRTED VALS TO SUPPORT RNGS TESTING
	bse.Lits = sys.Vs("-10s--1s", "0s-1s", "2s-4s", "6s-10s", "50s-60s")
	bse.Vals = sys.VsStruct(x.Typ().Full(), "Min:-10, Max:-1", "Min:0, Max:1", "Min:2, Max:4", "Min:6, Max:10", "Min:50, Max:60")
}
func (x *FleTmeRng) InitTrm(bse *TypBse, trmr *FleTrmr) {
	x.InitLexLit(func(r *TypFn, t *FleTrmr) {
		r.Add("r.Idx = x.Idx")

		r.Addf("r.MinTrm, ok = x.%v(true)", _sys.Bsc.Tme.Typ().Bse().LitTrm.Name)
		r.Add("if !ok {")
		r.Add("return r, false")
		r.Add("}")

		r.Add("if x.Ch != '-' {")
		r.Add("return r, false")
		r.Add("}")
		r.Add("x.NextRune()")

		r.Addf("r.MaxTrm, ok = x.%v()", _sys.Bsc.Tme.Typ().Bse().LitTrm.Name)
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
