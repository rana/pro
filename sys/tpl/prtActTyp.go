package tpl

import (
	"sys"
	"sys/k"
	"sys/tpl/atr"
)

type (
	PrtActTyp struct {
		PrtBse
		Fn  *Func
		Fns *Alias
		Act *Struct
	}
)

func (x *PrtActTyp) InitPrtTyp() {
	x.Fn = x.fn()
	x.Fns = x.fns()
	x.Act = x.act()
}
func (x *PrtActTyp) InitPrtTypFn() {
	x.actAct()
	x.fnsPll()
}

// typ
func (x *PrtActTyp) fn() (r *Func) {
	r = x.f.Func(x.f.Typ().Bse().PrefixTyp(k.Fn), atr.None)
	r.OutPrm(x.f.Typ())
	return r
}
func (x *PrtActTyp) fns() (r *Alias) {
	r = x.f.AliasSlice(sys.Plural(x.Fn.Name), x.Fn, atr.None)
	return r
}
func (x *PrtActTyp) act() (r *Struct) {
	r = x.f.StructPtr(x.f.Typ().Bse().PrefixTyp(k.Act), atr.None)
	r.Fld("Fn", x.Fn)
	r.Fld(x.f.Typ().Title(), x.t)
	return r
}

// fn
func (x *PrtActTyp) actAct() (r *TypFn) {
	r = x.f.TypFn(k.Act, x.Act)
	r.Addf("x.%v = x.Fn()", x.t.Title())
	return r
}
func (x *PrtActTyp) fnsPll() (r *TypFn) {
	x.f.Import(_sys)
	r = x.f.TypFna(k.Pll, atr.None, x.Fns)
	r.OutPrmSlice(x.f.Typ(), "r")
	r.Add("if len(*x) != 0 {")
	r.Addf("acts := make([]sys.Act, len(*x))")
	r.Add("for n, fn := range *x {")
	r.Addf("acts[n] = &%v{Fn: fn}", x.Act.Name)
	r.Add("}")
	r.Add("sys.Run().Pll(acts...)")
	r.Addf("r = make([]%v, len(*x))", x.f.Typ().Ref(x.f))
	r.Add("for n, act := range acts {")
	r.Addf("r[n] = act%v.%v", x.Act.Cast(x.f, true), x.t.Title())
	r.Add("}")
	r.Add("}")
	r.Add("return r")
	return r
}
