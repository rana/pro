package tpl

import (
	"sys/k"
	"sys/tpl/atr"
)

type (
	FleAnaPrfm struct {
		FleBse
		PrtStructStrWrt
		PrtString
		New *PkgFn
	}
	FleAnaPrfms struct {
		FleBse
		PrtArr
		PrtArrFld
		PrtArrStrWrt
		PrtString
	}
)

func (x *DirAna) NewPrfm() (r *FleAnaPrfm) {
	r = &FleAnaPrfm{}
	x.Prfm = r
	r.Name = k.Prfm
	r.Pkg = x.Pkg
	r.StructPtr(r.Name, atr.TypAnaStruct|atr.Test)
	r.AddFle(r)
	return r
}
func (x *DirAna) NewPrfms() (r *FleAnaPrfms) {
	r = &FleAnaPrfms{}
	x.Prfms = r
	r.FleBse = *NewArr(x.Prfm, &r.PrtArr, x.Prfm.Pkg)
	r.AddFle(r)
	return r
}
func (x *FleAnaPrfm) InitFld(s *Struct) {
	s.Fld("PnlPct", _sys.Bsc.Flt).Atr = atr.Get | atr.Srt | atr.Grp | atr.Dlt | atr.TstZeroSkp
	s.Fld("ScsPct", _sys.Bsc.Flt).Atr = atr.Get | atr.Srt | atr.Grp | atr.Dlt | atr.TstZeroSkp
	s.Fld("PipPerDay", _sys.Bsc.Flt).Atr = atr.Get | atr.Dlt
	s.Fld("UsdPerDay", _sys.Bsc.Flt).Atr = atr.Get | atr.Dlt
	s.Fld("ScsPerDay", _sys.Bsc.Flt).Atr = atr.Get | atr.Dlt | atr.TstZeroSkp
	s.Fld("OpnPerDay", _sys.Bsc.Flt).Atr = atr.Get | atr.Dlt | atr.TstZeroSkp
	s.Fld("PnlUsd", _sys.Bsc.Flt).Atr = atr.Get | atr.Dlt
	s.Fld("PipAvg", _sys.Bsc.Flt).Atr = atr.Get | atr.Dlt
	s.Fld("PipMdn", _sys.Bsc.Flt).Atr = atr.Get | atr.Dlt
	s.Fld("PipMin", _sys.Bsc.Flt).Atr = atr.Get | atr.Dlt
	s.Fld("PipMax", _sys.Bsc.Flt).Atr = atr.Get | atr.Dlt
	s.Fld("PipSum", _sys.Bsc.Flt).Atr = atr.Get | atr.Dlt
	s.Fld("DurAvg", _sys.Bsc.Tme).Atr = atr.Get | atr.Dlt
	s.Fld("DurMdn", _sys.Bsc.Tme).Atr = atr.Get | atr.Dlt
	s.Fld("DurMin", _sys.Bsc.Tme).Atr = atr.Get | atr.Dlt
	s.Fld("DurMax", _sys.Bsc.Tme).Atr = atr.Get | atr.Dlt
	s.Fld("LosLimMax", _sys.Bsc.Flt).Atr = atr.Get | atr.Srt | atr.Grp
	s.Fld("DurLimMax", _sys.Bsc.Tme).Atr = atr.Get | atr.Srt
	s.Fld("DayCnt", _sys.Bsc.Unt).Atr = atr.Get
	s.Fld("TrdCnt", _sys.Bsc.Unt).Atr = atr.Get | atr.Dlt
	s.Fld("TrdPct", _sys.Bsc.Flt).Atr = atr.Get | atr.Dlt
	s.Fld("CstTotUsd", _sys.Bsc.Flt).Atr = atr.Get
	s.Fld("CstSpdUsd", _sys.Bsc.Flt).Atr = atr.Get
	s.Fld("CstComUsd", _sys.Bsc.Flt).Atr = atr.Get
	s.Fld("Pth", _sys.Bsc.Str).Atr = atr.BytLitEqlSelSkp | atr.TstSkp // TstSkp to allow hst/rlt cmp; Pth is DskKey
}
func (x *FleAnaPrfm) InitTypFn() {
	x.Dlt()
}
func (x *FleAnaPrfm) Dlt() (r *TypFn) {
	r = x.TypFn(k.Dlt)
	r.InPrm(x, "v")
	r.OutPrm(_sys.Ana.PrfmDlt, "r")
	r.Addf("r = %v{}", r.OutTyp().Adr(x))
	s := x.Typ().(*Struct)
	for _, fld := range s.Flds {
		if fld.IsDlt() {
			r.Addf("r.%[1]vA = x.%[1]v", fld.Name)
			r.Addf("r.%[1]vB = v.%[1]v", fld.Name)
			if fld.Typ == _sys.Bsc.Flt.Typ() {
				r.Addf("r.%[1]vDlt = r.%[1]vB.Sub(r.%[1]vA).Trnc(1)", fld.Name)
				// r.Addf("r.%[1]vDltPct = r.%[1]vA.Pct(r.%[1]vB).Mul(100).Trnc(1)", fld.Name)
			} else {
				r.Addf("r.%[1]vDlt = flt.Flt(r.%[1]vB).Sub(flt.Flt(r.%[1]vA))", fld.Name)
				// r.Addf("r.%[1]vDltPct = flt.Flt(r.%[1]vA).Pct(flt.Flt(r.%[1]vB)).Mul(100).Trnc(1)", fld.Name)
			}
		}
	}
	r.Add("r.PthB = v.Pth")
	r.Add("return r")
	return r
}
