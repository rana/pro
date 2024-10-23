package tpl

import (
	"sys"
	"sys/k"
	"sys/tpl/atr"
)

type (
	FleScpAct struct {
		FleBse
	}
)

func (x *DirAct) NewScp() (r *FleScpAct) {
	r = &FleScpAct{}
	r.Name = k.Scp
	r.Pkg = x.Pkg
	r.StructPtr(k.Scp, atr.None)
	r.AddFle(r)
	return r
}
func (x *FleScpAct) InitFld(s *Struct) {
	s.Fld("Xpr", _sys.Lng.Pro.Xpr.Scp)
	s.Fld("Prnt", s)
	s.Fld("IdxRet", Uint32)
	for _, f := range _sys.Fles {
		if f.Typ() != nil && f.Typ().Bse().Scp != nil {
			s.FldSlice(sys.Camel(f.Typ().PkgTypTitle()), f)
		}
	}
}
func (x *FleScpAct) InitAct(actr *FleActr) {
	x.Var("FnRetIdn", "\"!fnRetIdn!\"", String)
	x.NewScp()
	for _, f := range _sys.Fles {
		if f.Typ() != nil && f.Typ().Bse().Scp != nil {
			x.ScpAcsGet(f.Typ())
		}
	}
}
func (x *FleScpAct) NewScp() (r *PkgFn) {
	x.Import(_sys.Lng.Pro.Xpr.Knd)
	r = x.PkgFn("NewScp")
	r.InPrm(_sys.Lng.Pro.Xpr.Scp, "xprScp")
	r.InPrmVariadic(x, "prnt")
	r.OutPrm(x, "r")
	r.Addf("r = &%v{}", x.Typ().Bse().Name)
	r.Add("r.Xpr = xprScp")
	r.Add("if len(prnt) != 0 {")
	r.Add("r.Prnt = prnt[0]")
	r.Add("}")
	r.Add("r.IdxRet = xprScp.Vars[FnRetIdn].Idx")
	r.Add("for xprKnd, cnt := range xprScp.Cnts { // allocate slices of exact count")
	r.Add("switch xprKnd {")
	for _, f := range _sys.Fles {
		if f.Typ() != nil && f.Typ().Bse().Scp != nil {
			r.Addf("case %v:", f.Typ().Bse().Knd.Ref(x))
			r.Addf("r.%v = make([]%v, cnt)", sys.Camel(f.Typ().PkgTypTitle()), f.Typ().Ref(x))
		}
	}
	r.Add("}")
	r.Add("}")
	r.Add("return r")
	return r
}

func (x *FleScpAct) ScpAcsGet(typ Typ) (r *TypFn) {
	x.Import("fmt")
	r = x.TypFnf("%v", typ.PkgTypTitle())
	r.InPrm(String, "idn")
	r.OutPrm(typ.Bse().Scp, "r")
	r.Add("cur := x")
	r.Add("for cur != nil {")
	r.Add("kndIdx, exists := cur.Xpr.Vars[idn]")
	r.Add("if exists {")
	r.Add("r.Idx = kndIdx.Idx")
	r.Addf("r.Arr = cur.%v", sys.Camel(typ.PkgTypTitle()))
	r.Add("return r")
	r.Add("}")
	r.Add("cur = cur.Prnt")
	r.Add("}")
	r.Add("panic(fmt.Sprintf(\"scope missing idn '%v'\", idn))")
	return r
}
