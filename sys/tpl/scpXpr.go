package tpl

import (
	"sys/k"
	"sys/tpl/atr"
)

type (
	FleScpXpr struct {
		FleBse
	}
)

func (x *DirXpr) NewScp() (r *FleScpXpr) {
	r = &FleScpXpr{}
	r.Name = k.Scp
	r.Pkg = x.Pkg
	r.Scp()
	r.AddFle(r)
	return r
}
func (x *FleScpXpr) Scp() (r *Struct) {
	x.Import(_sys.Lng.Pro.Xpr.Knd)
	r = x.StructPtr(k.Scp, atr.None)
	r.Fld("Prnt", r)
	r.Fld("Vars", NewExt("map[string]KndIdx"))
	r.Fld("Cnts", NewExtf("map[%v]uint32", _sys.Lng.Pro.Xpr.Knd.Typ().Ref(x)))
	return r
}
func (x *FleScpXpr) InitXpr(xprr *FleXprr) {
	x.NewScp()
	x.KndIdx()
	x.Knd()
	x.Idx()
	x.Decl()
}
func (x *FleScpXpr) NewScp() (r *PkgFn) {
	r = x.PkgFn("NewScp")
	r.InPrmVariadic(x, "prnt")
	r.OutPrm(x, "r")
	r.Addf("r = &%v{", x.Typ().Bse().Name)
	r.Add("Vars: make(map[string]KndIdx),")
	r.Addf("Cnts: make(map[%v]uint32),", _sys.Lng.Pro.Xpr.Knd.Typ().Ref(x))
	r.Add("}")
	r.Add("if len(prnt) != 0 {")
	r.Add("r.Prnt = prnt[0]")
	r.Add("}")
	r.Add("return r")
	return r
}
func (x *FleScpXpr) KndIdx() (r *Struct) {
	r = x.Struct("KndIdx", atr.None)
	r.Fld("Knd", _sys.Lng.Pro.Xpr.Knd)
	r.Fld("Idx", Uint32)
	return r
}
func (x *FleScpXpr) Knd() (r *TypFn) {
	r = x.TypFn("Knd")
	r.InPrm(String, "idn")
	r.OutPrm(_sys.Lng.Pro.Xpr.Knd, "knd")
	r.OutPrm(Bool, "exists")
	r.Add("var kndIdx KndIdx")
	r.Add("cur := x")
	r.Add("for cur != nil {")
	r.Add("kndIdx, exists = cur.Vars[idn]")
	r.Add("if exists {")
	r.Add("return kndIdx.Knd, exists")
	r.Add("}")
	r.Add("cur = cur.Prnt")
	r.Add("}")
	r.Add("return knd, false")
	return r
}

func (x *FleScpXpr) Idx() (r *TypFn) {
	r = x.TypFn("Idx")
	r.InPrm(String, "idn")
	r.OutPrm(Uint32)
	// r.Add("cur := x")
	// r.Add("for cur != nil {")
	// r.Add("_, exists := cur.Vars[idn]")
	// r.Add("if exists {")
	// r.Add("return cur.Vars[idn].Idx")
	// r.Add("}")
	// r.Add("cur = cur.Prnt")
	// r.Add("}")
	// r.Add("return math.MaxUint32")
	r.Add("return x.Vars[idn].Idx")
	return r
}
func (x *FleScpXpr) Decl() (r *TypFn) {
	x.Import("sys/err")
	r = x.TypFn("Decl")
	r.InPrm(String, "idn")
	r.InPrm(_sys.Lng.Pro.Xpr.Knd, "knd")
	r.Add("kndIdx, exists := x.Vars[idn]")
	r.Add("if exists {")
	r.Add("if kndIdx.Knd != knd {")
	r.Add("err.Panicf(\"idn redeclared: cannot redeclare an identifier with a different type (idn:%v expected:%v actual:%v)\", idn, kndIdx.Knd, knd)")
	r.Add("}")
	r.Add("}")
	r.Add("x.Vars[idn] = KndIdx{Knd:knd, Idx:x.Cnts[knd]}")
	r.Add("x.Cnts[knd]++")
	return r
}
