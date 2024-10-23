package tpl

import (
	"strings"
	"sys/k"
	"sys/tpl/atr"
)

type (
	PrtActFn struct {
		PrtBse
	}
)

func (x *PrtActFn) InitPrtTyp() {

}
func (x *PrtActFn) InitPrtTypFn() {
	var fns []*FnBse
	if rxr, ok := x.f.Typ().(*Struct); ok {
		for _, fn := range rxr.TypFns {
			fns = append(fns, &fn.FnBse)
		}
	} else if rxr, ok := x.f.Typ().(*Arr); ok {
		for _, fn := range rxr.TypFns {
			fns = append(fns, &fn.FnBse)
		}
	} else if rxr, ok := x.f.Typ().(*Alias); ok {
		for _, fn := range rxr.TypFns {
			fns = append(fns, &fn.FnBse)
		}
	} else if rxr, ok := x.f.Typ().(*Ifc); ok {
		for _, fn := range rxr.MemSigs {
			fns = append(fns, &fn.FnBse)
		}
	}
	if len(fns) != 0 {
		for _, fn := range fns {
			if fn.IsPrtActFn() {
				seg := x.typSeg(fn)
				x.segAct(fn, seg)
				if len(fn.InPrms) == 0 {
					x.fnPll(fn, seg)
				}
			}
		}
	}
}

// typ
func (x *PrtActFn) typSeg(fn *FnBse) (r *Struct) {
	r = x.f.StructPtrf("%v%vSeg", atr.None, x.f.Typ().Title(), fn.Title())
	r.Fld(x.t.Title(), x.t)
	for _, prm := range fn.InPrms {
		r.Fld(prm.Title(), prm.Typ)
	}
	if len(fn.OutPrms) != 0 {
		r.Fld("Out", fn.OutTyp())
	}
	return r
}

// fn
func (x *PrtActFn) segAct(fn *FnBse, seg *Struct) (r *TypFn) {
	r = x.f.TypFn(k.Act, seg)
	var inPrms strings.Builder
	for n, prm := range fn.InPrms {
		inPrms.WriteString("x.")
		inPrms.WriteString(prm.Title())
		if n != len(fn.InPrms)-1 {
			inPrms.WriteString(",")
		}
	}
	if len(fn.OutPrms) != 0 {
		r.Addf("x.Out = x.%v.%v(%v)", x.t.Title(), fn.Name, inPrms.String())
	} else {
		r.Addf("x.%v.%v(%v)", x.t.Title(), fn.Name, inPrms.String())
	}
	return r
}
func (x *PrtActFn) fnPll(fn *FnBse, seg *Struct) (r *PkgFn) {
	x.f.Import(_sys)
	r = x.f.PkgFnf("%v%vPll", x.t.Title(), fn.Title())
	r.InPrmVariadic(x.t, "vs")
	r.Add("if len(vs) != 0 {")
	r.Addf("acts := make([]sys.Act, len(vs))")
	r.Add("for n, v := range vs {")
	r.Addf("acts[n] = &%v{%v: v}", seg.Name, x.t.Title())
	r.Add("}")
	r.Add("sys.Run().Pll(acts...)")
	if len(fn.OutPrms) != 0 {
		r.OutPrmSlice(fn.OutTyp(), "r")
		r.Addf("r = make([]%v, len(vs))", fn.OutTyp().Ref(x.f))
		r.Add("for n, act := range acts {")
		r.Addf("r[n] = act%v.Out", seg.Cast(x.f, true))
		r.Add("}")
		r.Add("}")
		r.Add("return r")
	} else {
		r.Add("}")
	}
	return r
}
