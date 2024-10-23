package tpl

import (
	"sys/k"
	"sys/tpl/atr"
)

type (
	FleCfgr struct {
		FleBse
		Cfg *DirCfg
	}
)

func (x *DirCfg) NewCfgr() (r *FleCfgr) {
	r = &FleCfgr{Cfg: x}
	r.Name = k.Cfgr
	r.Pkg = x.Pkg
	r.StructPtr(r.Name, atr.Test) // Cfgr is a cfg term lexer.
	r.AddFle(r)
	return r
}
func (x *FleCfgr) InitFld(s *Struct) {
	x.Import(_sys.Lng.Pro.Trm)
	x.Import(_sys.Lng.Pro.Trm.Prs)
	s.FldExt("trm.Trmr") // use Ext to avoid Ptr mod
}
func (x *FleCfgr) InitTypFn() {
	x.SrchKey()
	x.SkpVal()
	x.Test.Import(x)
}

func (x *FleCfgr) SrchKey() (r *TypFn) {
	r = x.TypFn("SrchKey")
	r.InPrmVariadic(String, "path")
	r.OutPrm(Bool)
	r.Add("if len(path) > 0 {")
	r.Add("x.SkpSpceCmnt()")
	r.Add("if x.Ch != '{' { // must start with lcrl")
	r.Add("return false")
	r.Add("}")
	r.Add("x.SkpSpceCmnt()")
	r.Add("x.NextRune() // skip lcrl")
	r.Add("x.SkpSpceCmnt()")
	r.Add("idn, ok := x.IdnLit()")
	r.Add("for !x.End && ok {")
	r.Add("x.SkpSpceCmnt()")
	r.Add("if x.Ch != ':' { // cln must follow idn")
	r.Add("return false")
	r.Add("}")
	r.Add("x.NextRune() // skip cln")
	r.Add("if path[0] == x.Txt[idn.Idx:idn.Lim] {")
	r.Add("x.SkpSpceCmnt()")
	r.Add("if len(path) == 1 { // found final key")
	r.Add("return true")
	r.Add("}")
	r.Add("return x.SrchKey(path[1:]...) // look for inr key")
	r.Add("}")
	r.Add("// continue search for key at current depth")
	r.Add("if !x.SkpVal() { // skip over current value; may have nested values")
	r.Add("return false")
	r.Add("}")
	r.Add("idn, ok = x.IdnLit()")
	r.Add("}")
	r.Add("}")
	r.Add("return false")
	return r
}
func (x *FleCfgr) SkpVal() (r *TypFn) {
	x.Import("unicode")
	r = x.TypFn("SkpVal")
	r.OutPrm(Bool)
	r.Add("x.SkpSpceCmnt()")
	r.Add("switch {")
	r.Add("case x.Ch == '\"': // skip over str typ")
	r.Add("_, ok := x.StrLit()")
	r.Add("if !ok {")
	r.Add("return false")
	r.Add("}")
	r.Add("case x.Ch == '[': // skip over arr typs")
	r.Add("x.SkpSet('[', ']')")
	r.Add("case x.Ch == '{': // skip over obj")
	r.Add("x.SkpSet('{', '}')")
	r.Add("default: // skip over all non-str types")
	r.Add("for !x.End && !unicode.IsSpace(x.Ch) {")
	r.Add("x.NextRune()")
	r.Add("}")
	r.Add("}")
	r.Add("x.SkpSpceCmnt()")
	r.Add("return true")
	return r
}

func (x *FleCfgr) PrsCfg(fle Fle) (r *TypFn) {
	typ, bse := fle.Typ(), fle.Typ().Bse()
	x.Import(_sys)
	x.Import("sys/err")
	if bse.Pkg.Title() != bse.Title() {
		r = x.TypFn(bse.PkgTypTitle())
	} else {
		r = x.TypFn(bse.Title())
	}
	r.InPrmVariadic(String, "path")
	r.OutPrm(typ, "r")
	r.Add("x.Reset(x.Txt) // enables multiple sequential calls")
	r.Add("if x.SrchKey(path...) {")
	r.Addf("%v, ok := x.%v()", bse.LitTrm.Camel(), bse.LitTrm.Title())
	r.Add("if !ok {")
	r.Addf("err.Panicf(\"Cfgr: invalid %v (path:%%q)\", sys.JoinPth(path...))", r.Name)
	r.Add("}")
	r.Addf("return %v(%v, x.Txt)", bse.PrsTrm.Ref(x), bse.LitTrm.Camel())
	r.Add("}")
	r.Add("err.Panicf(\"Cfgr: invalid path (path:%q)\", sys.JoinPth(path...))")
	r.Add("return r")
	x.Test.Cfg(r, typ)
	return r
}
