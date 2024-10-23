package tpl

import "sys/k"

type (
	PrtElmArr struct {
		PrtBse
	}
)

func (x *PrtElmArr) InitPrtFld() {
	x.t.bse.Fldf(x.t.elm.arr.Camel(), x.t.elm.arr) //.Atr = atr.TstZeroSkp
}
func (x *PrtElmArr) InitPrtTypFn() {
	x.ElmArr()
	x.Rev()
}
func (x *PrtElmArr) ElmArr() (r *TypFn) {
	r = x.f.TypFn(x.t.elm.arr.Title(), x.t.bse)
	r.OutPrm(x.t.elm.arr)
	r.Addf("return x.%v", x.t.elm.arr.Camel())
	x.f.MemSigFn(r) // add to interface
	return r
}
func (x *PrtElmArr) Rev() (r *TypFn) {
	r = x.f.TypFn(k.Rev, x.t.bse)
	r.OutPrm(x.t)
	r.Addf("x.%v.Rev()", x.t.elm.arr.Camel())
	r.Add("return x.Slf")
	x.f.MemSigFn(r) // add to interface
	return r
}
