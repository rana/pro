package tpl

import (
	"sys/k"
)

type (
	PrtArrRng struct {
		PrtBse
		Arr *Arr
		Rng *PrtRng
	}
)

func (x *PrtArrRng) InitPrtTyp() {

}
func (x *PrtArrRng) InitPrtTypFn() {
	x.Arr = x.f.GetPrt((*PrtArr)(nil)).(*PrtArr).Arr
	x.Rng = x.Arr.Elm.Fle.Bse().GetPrt((*PrtRng)(nil)).(*PrtRng)
	x.SrchIdx()
	x.RngMrg()
}
func (x *PrtArrRng) SrchIdx() (r *TypFn) {
	r = x.f.TypFn(k.SrchIdx)
	r.InPrm(x.Rng.Elm.Typ().Bse(), "v")
	r.OutPrm(_sys.Bsc.Unt)
	r.Add("i, j := unt.Zero, unt.Unt(len(*x))")
	r.Add("for i < j {")
	r.Add("if (*x)[(i+j)>>1].Max < v {")
	r.Add("i = (i+j)>>1 + 1")
	r.Add("} else {")
	r.Add("j = (i+j)>>1")
	r.Add("}")
	r.Add("}")
	r.Add("return i")
	return r
}
func (x *PrtArrRng) RngMrg() (r *TypFn) {
	r = x.f.TypFn("RngMrg")
	r.InPrm(_sys.Bsc.Unt, "fstIdx")
	r.InPrm(_sys.Bsc.Unt, "lstIdx")
	r.OutPrm(x.Rng.t)
	r.Add("return x.At(fstIdx).Mrg(x.At(lstIdx))")
	return r
}
