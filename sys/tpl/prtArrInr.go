package tpl

import (
	"strings"
	"sys"
	"sys/ks"
)

type (
	PrtArrInr struct {
		PrtBse
		Arr *Arr
	}
)

func (x *PrtArrInr) InitPrtTyp() {
	x.Arr = x.f.GetPrt((*PrtArr)(nil)).(*PrtArr).Arr
}
func (x *PrtArrInr) InitPrtTypFn() {
	for _, inr := range ks.Inrs {
		x.op(inr)
	}
}
func (x *PrtArrInr) op(name string) (r *TypFn) {
	name = strings.Title(name)
	r = x.f.TypFn(sys.CnjInr(name))
	r.InPrm(_sys.Bsc.Unt, "off")
	r.OutPrm(x.f.Typ())
	r.Add("if len(*x) < int(off) {")
	r.Addf("r := make(%v, 0)", x.t.Title())
	r.Add("return &r")
	r.Add("}")
	r.Addf("r := make(%v, len(*x)-int(off))", x.t.Title())
	r.Add("for n := 0; n < len(r); n++ {")
	r.Addf("r[n] = (*x)[n+int(off)].%v((*x)[n])", name)
	r.Add("}")
	r.Add("return &r")
	// test
	if x.f.Test != nil && Opt.IsArr() {
		r.T.SkpCpy = true
		x.f.Import(_sys.Bsc.Unt)
		r.T.Addf("var expected %v", x.t.Ref(x.f.Test))
		r.T.Add("if x.Cnt() < off {")
		r.T.Addf("expected = %v(0)", x.Arr.Make.Ref(x.f.Test))
		r.T.Add("} else {")
		r.T.Addf("expected =  %v(x.Cnt()-off)", x.Arr.Make.Ref(x.f.Test))
		r.T.Addf("for n:= %v; n < expected.Cnt(); n++ {", _sys.Bsc.Unt.Zero.Ref(x.f.Test))
		r.T.Addf("expected.Upd(n, x.At(n+off).%v(x.At(n)))", name)
		r.T.Add("}")
		r.T.Add("}")
	}
	return r
}
