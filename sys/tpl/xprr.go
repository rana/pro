package tpl

import (
	"fmt"
	"strings"
	"sys/k"
	"sys/tpl/atr"
)

type (
	FleXprr struct {
		FleBse
		Pkgs
		PllWaitXpr *Struct
	}
)

func (x *DirXpr) NewXprr() (r *FleXprr) {
	r = &FleXprr{}
	r.Name = k.Xprr
	r.Pkg = x.Pkg
	r.Xprr()
	r.AddFle(r)
	return r
}
func (x *FleXprr) Xprr() (r *Struct) {
	x.Import(_sys.Lng.Pro.Trm)
	r = x.StructPtr(k.Xprr, atr.Test) // Xprr is an expression parser
	r.FldExt("trm.Trmr")              // use Ext to avoid Ptr mod
	return r
}
func (x *FleXprr) InitXpr(xprr *FleXprr) {
	x.PllWaitXpr = x.pllWaitXpr()
	x.Prs()
	x.Prsf()
	x.PrsXprs()
	x.PrsXpr()
	for _, p := range x.Pkgs {
		x.PrsPkgXpr(p)
		for _, f := range p.Fles {
			for _, typ := range f.Bse().Typs {
				if _, isExt := typ.(*Ext); !isExt && typ.Bse().IsXpr() {
					// if typ.Bse().IsXpr() {
					x.PrsTypXpr(typ.Bse(), f.Bse())
				}
			}
			// fleBse := f.Bse()
			// if fleBse.Typ() != nil && fleBse.Typ().Bse().IsXpr() {
			// 	x.PrsTypXpr(fleBse.Typ().Bse(), fleBse)
			// }
		}
	}
	x.HasMem()
	x.NextLprn()
	x.NextRprn()
	x.Erf()
	x.Panicf()
}

func (x *FleXprr) NewXpr(name string, ret ...Typ) (r *Struct) {
	r = x.Struct(name, atr.None)
	x.TypFn(k.Xpr, r)
	if len(ret) != 0 {
		x.TypFn(ret[0].PkgTypTitle()+"Xpr", r) // PkgTypTitle needed to distinguish hst/rlt Stm etc

		for _, ifc := range ret[0].Bse().Ifcs {
			x.TypFn(ifc.PkgTypTitle()+"Xpr", r)
		}

		// if ret[0] != Interface {
		// 	x.TypFn("SysInterfaceXpr", r)
		// }
	}
	return r
}
func (x *FleXprr) NewXprf(format string, ret Typ, args ...interface{}) (r *Struct) {
	return x.NewXpr(fmt.Sprintf(format, args...), ret)
}
func (x *FleXprr) TypIfc(typ Typ) (r *Ifc) {
	if typ == Interface {
		typ.Bse().IfcXpr = _sys.Lng.Pro.Xpr.Ifc
		return typ.Bse().IfcXpr
	}
	r = x.Ifcf("%vXpr", atr.None, typ.PkgTypTitle())
	typ.Bse().IfcXpr = r
	x.MemSiga(k.Xpr, atr.None, r)
	x.MemSigRxrf("%v%v", r, typ.PkgTypTitle(), strings.Title(k.Xpr)) // PkgTypTitle needed to distinguish hst/rlt Stm
	for _, ifc := range typ.Bse().Ifcs {
		x.MemSigRxrf("%v%v", r, ifc.PkgTypTitle(), strings.Title(k.Xpr))
	}
	return r
}
func (x *FleXprr) LitXpr(ret Typ) (r *Struct) {
	bse := ret.Bse()
	r = x.NewXprf("%vLit", ret, ret.PkgTypTitle())
	ret.Bse().LitXpr = r
	r.Fld(strings.Title(k.Trm), bse.LitTrm)
	if x.Test != nil && Opt.IsTest() && bse.IsTestXpr() {
		x.Test.XprLit(ret)
	}
	return r
}
func (x *FleXprr) AsnXpr(ret Typ) (r *Struct) {
	bse := ret.Bse()
	r = x.NewXprf("%vAsn", ret, ret.PkgTypTitle())
	ret.Bse().AsnXpr = r
	r.Fld(strings.Title(k.Trm), _sys.Bsc.Bnd)
	r.Fld("X", ret.Bse().IfcXpr)
	r.Fld(strings.Title(k.Idn), _sys.Bsc.Bnd)
	r.Fld("Depth", Int)
	if x.Test != nil && Opt.IsTest() && bse.IsTestXpr() {
		x.Test.XprAsn(ret)
	}
	return r
}
func (x *FleXprr) AcsXpr(ret Typ) (r *Struct) {
	bse := ret.Bse()
	r = x.NewXprf("%vAcs", ret, ret.PkgTypTitle())
	ret.Bse().AcsXpr = r
	r.Fld(strings.Title(k.Trm), _sys.Bsc.Bnd)
	if x.Test != nil && Opt.IsTest() && bse.IsTestXpr() {
		x.Test.XprAcs(ret)
	}
	return r
}
func (x *FleXprr) ThenXpr(ret Typ) (r *Struct) {
	bse := ret.Bse()
	r = x.NewXprf("%vThen", ret, ret.PkgTypTitle())
	ret.Bse().ThenXpr = r
	r.Fld(strings.Title(k.Trm), _sys.Bsc.Bnd)
	r.Fld("X", ret.Bse().IfcXpr)
	r.FldSlice("Xprs", _sys.Lng.Pro.Xpr.Ifc)
	r.Fld("Scp", _sys.Lng.Pro.Xpr.Scp)
	if x.Test != nil && Opt.IsTest() && bse.IsTestXpr() {
		x.Test.XprThen(ret)
	}
	return r
}
func (x *FleXprr) ElseXpr(ret Typ) (r *Struct) {
	bse := ret.Bse()
	r = x.NewXprf("%vElse", ret, ret.PkgTypTitle())
	ret.Bse().ElseXpr = r
	r.Fld(strings.Title(k.Trm), _sys.Bsc.Bnd)
	r.Fld("X", ret.Bse().IfcXpr)
	r.FldSlice("Xprs", _sys.Lng.Pro.Xpr.Ifc)
	r.Fld("Scp", _sys.Lng.Pro.Xpr.Scp)
	if x.Test != nil && Opt.IsTest() && bse.IsTestXpr() {
		x.Test.XprElse(ret)
	}
	return r
}
func (x *FleXprr) EachXpr(bse *TypBse, arr *Arr) (r *Struct) {
	r = x.NewXprf("%vEach", arr, bse.PkgTypTitle())
	// sys.Log("FleXprr.EachXpr", r.Name)
	bse.EachXpr = r
	r.Fld(strings.Title(k.Trm), _sys.Bsc.Bnd)
	r.Fld("X", bse.IfcXpr)
	r.Fld("Idn", _sys.Bsc.Bnd)
	r.FldSlice("Xprs", _sys.Lng.Pro.Xpr.Ifc)
	r.Fld("Scp", _sys.Lng.Pro.Xpr.Scp)
	if x.Test != nil && Opt.IsTest() && bse.IsTestXpr() {
		x.Test.XprEach(arr)
	}
	return r
}
func (x *FleXprr) PllEachXpr(ret *Arr) (r *Struct) {
	r = x.NewXprf("%vPllEach", ret, ret.PkgTypTitle())
	ret.Bse().PllEachXpr = r
	r.Fld(strings.Title(k.Trm), _sys.Bsc.Bnd)
	r.Fld("X", ret.Bse().IfcXpr)
	r.Fld("Idn", _sys.Bsc.Bnd)
	r.FldSlice("Xprs", _sys.Lng.Pro.Xpr.Ifc)
	r.Fld("Scp", _sys.Lng.Pro.Xpr.Scp)
	if x.Test != nil && Opt.IsTest() && ret.IsTestXpr() {
		x.Test.XprPllEach(ret)
	}
	return r
}
func (x *FleXprr) pllWaitXpr() (r *Struct) {
	r = x.NewXpr(k.PllWait)
	r.Fld(strings.Title(k.Trm), _sys.Bsc.Bnd)
	r.FldSlice("Xprs", _sys.Lng.Pro.Xpr.Ifc)
	r.Fld("Scp", _sys.Lng.Pro.Xpr.Scp)
	return r
}
func (x *FleXprr) FldGetXpr(s *Struct, fld *Fld) (r *Struct) {
	s.XprFldGets = append(s.XprFldGets, fld)
	r = x.NewXprf("%v%vGet", fld.Typ, s.PkgTypTitle(), fld.Title())
	fld.GetXpr = r
	r.Fld(strings.Title(k.Trm), _sys.Bsc.Bnd)
	r.Fld("X", s.IfcXpr)
	if x.Test != nil && Opt.IsTest() && s.IsTestXpr() && fld.IsTestXpr() {
		x.Test.XprFldGet(s, fld)
	}
	return r
}
func (x *FleXprr) FldSetGetXpr(s *Struct, fld *Fld) (r *Struct) {
	s.XprFldSetGets = append(s.XprFldSetGets, fld)
	r = x.NewXprf("%v%vSetGet", fld.Typ, s.PkgTypTitle(), fld.Title())
	fld.SetGetXpr = r
	r.Fld(strings.Title(k.Trm), _sys.Bsc.Bnd)
	r.Fld("X", s.IfcXpr)
	r.Fldf("I0", fld.Typ.Bse().IfcXpr)
	if x.Test != nil && Opt.IsTest() && s.IsTestXpr() && fld.IsTestXpr() {
		x.Test.XprFldSetGet(s, fld)
	}
	return r
}
func (x *FleXprr) CnstXpr(c *Cnst) (r *Struct) {
	bse := c.Typ.Bse()
	r = x.NewXprf("%v%v", c.Typ, c.Typ.Bse().Pkg.Title(), c.Title())
	c.Xpr = r
	r.Fld(strings.Title(k.Trm), _sys.Bsc.Bnd)
	if x.Test != nil && Opt.IsTest() && bse.IsTestXpr() && c.IsTestXpr() {
		x.Test.XprCnst(c)
	}
	return r
}
func (x *FleXprr) VarXpr(v *Var) (r *Struct) {
	bse := v.Typ.Bse()
	r = x.NewXprf("%v%v", v.Typ, v.Typ.Bse().Pkg.Title(), v.Title())
	v.Xpr = r
	r.Fld(strings.Title(k.Trm), _sys.Bsc.Bnd)
	if x.Test != nil && Opt.IsTest() && bse.IsTestXpr() && v.IsTestXpr() {
		x.Test.XprVar(v)
	}
	return r
}
func (x *FleXprr) PkgFnXpr(fn *PkgFn) (r *Struct) {
	// sys.Log("-", fn.Name)
	bse := fn.OutTyp().Bse()
	r = x.NewXprf("%v%v", fn.OutTyp(), fn.Pkg.Title(), fn.Title())
	fn.Xpr = r
	r.Fld(strings.Title(k.Trm), _sys.Bsc.Bnd)
	for n, inPrm := range fn.InPrms {
		if n == fn.InPrms.LstIdx() && inPrm.IsVariadic() {
			r.FldSlicef("I%v", inPrm.Typ.Bse().IfcXpr, n)
		} else {
			r.Fldf("I%v", inPrm.Typ.Bse().IfcXpr, n)
		}
	}
	if x.Test != nil && Opt.IsTest() && bse.IsTestXpr() && fn.IsTestXpr() {
		x.Test.XprPkgFn(fn)
	}
	return r
}
func (x *FleXprr) TypFnXpr(fn *TypFn) (r *Struct) {
	rxrBse := fn.Rxr.Typ.Bse()
	r = x.NewXprf("%v%v", fn.OutTyp(), fn.Rxr.Typ.PkgTypTitle(), fn.Title())
	fn.Xpr = r
	r.Fld(strings.Title(k.Trm), _sys.Bsc.Bnd)
	r.Fld("X", rxrBse.IfcXpr)
	for n, inPrm := range fn.InPrms {
		if n == fn.InPrms.LstIdx() && inPrm.IsVariadic() {
			r.FldSlicef("I%v", inPrm.Typ.Bse().IfcXpr, n)
		} else {
			r.Fldf("I%v", inPrm.Typ.Bse().IfcXpr, n)
		}
	}
	if x.Test != nil && Opt.IsTest() && rxrBse.IsTestXpr() && fn.IsTestXpr() {
		x.Test.XprTypFn(fn)
	}
	return r
}
func (x *FleXprr) MemSigXpr(fn *MemSig) (r *Struct) {
	rxrBse := fn.Rxr.Bse()
	outBse := fn.OutTyp().Bse()
	r = x.NewXprf("%v%v", fn.OutTyp(), fn.Rxr.PkgTypTitle(), fn.Title())
	fn.Xpr = r
	r.Fld(strings.Title(k.Trm), _sys.Bsc.Bnd)
	r.Fld("X", rxrBse.IfcXpr)
	for n, inPrm := range fn.InPrms {
		if n == fn.InPrms.LstIdx() && inPrm.IsVariadic() {
			r.FldSlicef("I%v", inPrm.Typ.Bse().IfcXpr, n)
		} else {
			r.Fldf("I%v", inPrm.Typ.Bse().IfcXpr, n)
		}
	}
	if x.Test != nil && Opt.IsTest() && outBse.IsTestXpr() && fn.IsTestXpr() {
		x.Test.XprMemSig(fn)
	}
	return r
}

func (x *FleXprr) Prs() (r *TypFn) {
	r = x.TypFn(k.Prs)
	r.InPrm(String, "txt")
	r.OutPrm(_sys.Lng.Pro.Xpr.Scp, "scp")
	r.OutPrmSlice(_sys.Lng.Pro.Xpr.Ifc, "xprs")
	r.Add("x.Reset(txt)")
	r.Add("scp = NewScp()")
	r.Add("return scp, x.Xprs(scp)")
	return r
}
func (x *FleXprr) Prsf() (r *TypFn) {
	x.Import("fmt")
	r = x.TypFn("Prsf")
	r.InPrm(String, "format")
	r.InPrmVariadic(Interface, "args")
	r.OutPrm(_sys.Lng.Pro.Xpr.Scp, "scp")
	r.OutPrmSlice(_sys.Lng.Pro.Xpr.Ifc, "xprs")
	r.Add("return x.Prs(fmt.Sprintf(format, args...))")
	return r
}
func (x *FleXprr) PrsXprs() (r *TypFn) {
	r = x.TypFn(k.Xprs)
	r.InPrm(_sys.Lng.Pro.Xpr.Scp, "scp")
	r.OutPrmSlice(_sys.Lng.Pro.Xpr.Ifc, "r")
	r.Add("for !x.End {")
	r.Add("r = append(r, x.Xpr(scp))")
	r.Add("x.SkpSpceCmnt()")
	r.Add("}")
	r.Add("return r")
	return r
}
func (x *FleXprr) PrsXpr() (r *TypFn) {
	x.Import(_sys.Lng.Pro.Xpr.Knd)
	r = x.TypFn(k.Xpr)
	r.InPrm(_sys.Lng.Pro.Xpr.Scp, "scp")
	r.InPrmVariadic(Bool, "retNil")
	r.OutPrm(_sys.Lng.Pro.Xpr.Ifc, "r")
	r.Add("x.SkpSpceCmnt()")
	r.Add("var ok bool")
	r.Add("scn := x.Scn // lit: bsc")
	for _, p := range x.Pkgs { // lit: bsc
		// sys.Log("-", p)
		for _, f := range p.Fles {
			if f.Typ() != nil {
				bse := f.Typ().Bse()
				if !bse.IsArr() && bse.LitTrm != nil {
					r.Addf("%vTrm, ok := x.%v()", bse.LitTrm.Camel(), bse.LitTrm.Name)
					r.Add("if ok {")
					r.Addf("return x.%vXpr(scp, &%v{Trm: %vTrm})", f.Typ().PkgTypTitle(), bse.LitXpr.Ref(x), bse.LitTrm.Camel())
					r.Add("}")
					r.Add("x.Scn = scn")
				}
			}
		}
	}
	// pllWait
	r.Add("pllWait, ok := x.PllWait() // pllWait")
	r.Add("if ok {")
	r.Add("x.NextLprn()")
	r.Add("x.SkpSpceCmnt()")
	r.Add("pllWaitScp := NewScp(scp)")
	r.Addf("var xprs []%v", _sys.Lng.Pro.Xpr.Ifc.Ref(x))
	r.Add("for {")
	r.Add("cur := x.Xpr(pllWaitScp, true)")
	r.Add("if cur == nil {")
	r.Add("break")
	r.Add("}")
	r.Add("xprs = append(xprs, cur)")
	r.Add("}")
	r.Add("x.NextRprn()")
	r.Addf("return &%v{Trm: pllWait, Xprs: xprs, Scp: pllWaitScp}", x.PllWaitXpr.Ref(x))
	r.Add("}")
	r.Add("x.Scn = scn")
	if x.Pkgs.Ok() {
		abcPkgs := make([][]*Pkg, 26) // gather fns by fst ch
		for _, p := range x.Pkgs {
			idx := p.Lower()[0] - 'a'
			abcPkgs[idx] = append(abcPkgs[idx], p)
		}
		r.Add("switch x.Ch {")
		r.Add("case '[': // arr: lit")
		var arrs []*PrtArr
		for _, p := range x.Pkgs {
			for _, f := range p.Fles {
				if f.Typ() != nil {
					bse := f.Typ().Bse()
					if bse.IsAct() && bse.IsArr() {
						if bse.LitTrm == nil {
							arrs = append(arrs, bse.PrtArr())
						} else { // arr: lit
							r.Addf("%vTrm, ok := x.%v()", bse.LitTrm.Camel(), bse.LitTrm.Name)
							r.Add("if ok {")
							r.Addf("return x.%vXpr(scp, &%v{Trm: %vTrm})", f.Typ().PkgTypTitle(), bse.LitXpr.Ref(x), bse.LitTrm.Camel())
							r.Add("}")
							r.Add("x.Scn = scn")
						}
					}
				}
			}
		}
		if len(arrs) != 0 { // arr: obj
			trm := _sys.Lng.Pro.Trm.Trmr.Objs
			r.Addf("%vTrm, ok := x.%v() // arr: obj", trm.Camel(), trm.Title())
			r.Add("if ok {")
			r.Add("x.Scn = scn")
			r.Add("x.NextRune() // skp lsqr")
			r.Add("var xprs []Xpr")
			r.Add("for !x.End && x.Ch != ']' {")
			r.Add("xprs = append(xprs, x.Xpr(scp))")
			r.Add("x.SkpSpceCmnt()")
			r.Add("}")
			r.Add("x.NextRune() // skp rsqr")
			r.Add("if len(xprs) == 0 {")
			r.Add("x.Panicf(\"empty array\")")
			r.Add("}")
			r.Add("switch xprs[0].(type) {")
			for _, arr := range arrs {
				elmXpr := arr.Arr.Elm.IfcXpr
				r.Addf("case %v:", elmXpr.Ref(x))
				r.Addf("i0 := make([]%v, len(xprs))", elmXpr.Ref(x))
				r.Add("for n, xpr := range xprs {")
				r.Addf("elm, ok := xpr.(%v)", elmXpr.Ref(x))
				r.Add("if !ok {")
				r.Addf("x.Panicf(\"inconsistent array (0:%v %%v:%%v)\", n, xpr)", elmXpr.Ref(x))
				r.Add("}")
				r.Add("i0[n] = elm")
				r.Add("}")
				r.Addf("return x.%vXpr(scp, &%v{Trm: %vTrm.Bnd, I0: i0})", arr.Fle.Typ().PkgTypTitle(), arr.Arr.New.Xpr.Adr(x), trm.Camel())

			}
			r.Add("}")
			r.Add("}")
			r.Add("x.Scn = scn")
		}
		for m, curPkgs := range abcPkgs {
			if len(curPkgs) > 0 {
				r.Addf("case '%v':", string(rune(m)+'a'))
				for _, p := range curPkgs {
					r.Addf("r, ok = x.%vPkgXpr(scp)", p.Title())
					r.Add("if ok {")
					r.Add("return r")
					r.Add("}")
					r.Add("x.Scn = scn")
				}
			}
		}
		r.Add("}")
	}
	r.Addf("idn, ok := x.IdnLit() // acs")
	r.Add("if ok {")
	r.Add("xprKnd, exists := scp.Knd(x.Txt[idn.Idx:idn.Lim])")
	r.Add("if !exists {")
	r.Add("x.Panicf(\"acs xpr: idn '%v' not declared\", x.Txt[idn.Idx:idn.Lim])")
	r.Add("}")
	r.Add("switch xprKnd {")
	for _, p := range x.Pkgs { // acs
		for _, f := range p.Fles {
			if f.Typ() != nil && f.Typ().Bse().AcsXpr != nil {
				bse := f.Typ().Bse()
				r.Addf("case %v:", bse.Knd.Ref(x))
				r.Addf("return x.%vXpr(scp, &%v{Trm: idn.Bnd})", bse.PkgTypTitle(), bse.AcsXpr.Ref(x))
			}
		}
	}
	r.Add("}")
	r.Add("}")
	r.Add("x.Scn = scn // rewind")
	r.Add("if len(retNil) > 0 && retNil[0] {")
	r.Add("return nil")
	r.Add("}")
	r.Addf("panic(x.Erf(\"%v: no expression found\"))", r.Name)
	return r
}

type (
	TypMem struct {
		FnBse *FnBse
		Fld   *Fld
	}
)

func (x *FleXprr) PrsTypXpr(rxr *TypBse, f *FleBse) (r *TypFn) {
	r = x.TypFnf("%vXpr", rxr.PkgTypTitle())
	r.InPrm(_sys.Lng.Pro.Xpr.Scp, "scp")
	r.InPrm(rxr.IfcXpr, "X")
	r.OutPrm(_sys.Lng.Pro.Xpr.Ifc)
	r.Add("if x.HasMem() {")
	// if rxr.AcsXpr != nil || f.TypFns.MayXpr() || f.MemSigs.MayXpr() ||
	if rxr.AcsXpr != nil || rxr.TypFns.MayXpr() || f.MemSigs.MayXpr() ||
		rxr.XprFldGets.Ok() || rxr.XprFldSetGets.Ok() {
		r.Add("scn := x.Scn")
	}
	// if f.TypFns.Ok() || f.MemSigs.Ok() || rxr.Scp != nil {
	if rxr.TypFns.Ok() || f.MemSigs.Ok() || rxr.Scp != nil {
		abcTypMems := make([][]*TypMem, 26) // gather TypMems by fst ch
		if rxr.AsnXpr != nil {              // asn
			abcTypMems['a'-'a'] = append(abcTypMems['a'-'a'], nil)
		}
		if rxr.EachXpr != nil || rxr.ElseXpr != nil { // each, else
			abcTypMems['e'-'a'] = append(abcTypMems['e'-'a'], nil)
		}
		if rxr.PllEachXpr != nil { // pllEach, pllWait
			abcTypMems['p'-'a'] = append(abcTypMems['p'-'a'], nil)
		}
		if rxr.ThenXpr != nil { // then
			abcTypMems['t'-'a'] = append(abcTypMems['t'-'a'], nil)
		}
		// for _, fn := range f.TypFns {
		for _, fn := range rxr.TypFns {
			if fn.MayXpr() {
				idx := fn.Lower()[0] - 'a'
				abcTypMems[idx] = append(abcTypMems[idx], &TypMem{FnBse: &fn.FnBse})
			}
		}
		for _, fn := range f.MemSigs {
			if fn.MayXpr() {
				idx := fn.Lower()[0] - 'a'
				abcTypMems[idx] = append(abcTypMems[idx], &TypMem{FnBse: &fn.FnBse})
			}
		}
		for _, fld := range rxr.XprFldGets {
			idx := fld.Lower()[0] - 'a'
			abcTypMems[idx] = append(abcTypMems[idx], &TypMem{Fld: fld})
		}
		for _, fld := range rxr.XprFldSetGets {
			idx := fld.Lower()[0] - 'a'
			abcTypMems[idx] = append(abcTypMems[idx], &TypMem{Fld: fld})
		}
		r.Add("switch x.Ch {")
		for m, curTypMems := range abcTypMems {
			if len(curTypMems) > 0 {
				r.Addf("case '%v':", string(rune(m)+'a'))
				for _, tm := range curTypMems {
					if tm != nil && tm.FnBse != nil { // fn
						x.WriteInPrms(r, tm.FnBse, true)
						r.Add("x.NextRprn()")
						r.Addf("return x.%vXpr(scp, cur)", tm.FnBse.OutPrms.Fst().Typ.PkgTypTitle())
						r.Add("}")
						r.Add("x.Scn = scn // rewind")
					} else if tm != nil && tm.Fld != nil { // fld
						r.Addf("%vTrm, ok := x.%v()", tm.Fld.Camel(), tm.Fld.Name) // USE 'Trm' SUFFIX FOR 'x' NAME
						r.Add("if ok {")
						r.Add("x.NextLprn()")
						if tm.Fld.IsGet() {
							r.Add("x.SkpSpceCmnt()")
							r.Add("x.NextRprn()")
							r.Addf("return x.%vXpr(scp, &%v{Trm: %vTrm, X: X})", tm.Fld.Typ.PkgTypTitle(), tm.Fld.GetXpr.Ref(x), tm.Fld.Camel())
						} else { // SetGet
							r.Add("x.SkpSpceCmnt()")
							r.Addf("var i0 %v", tm.Fld.Typ.Bse().IfcXpr.Ref(x))
							r.Add("i0Xpr := x.Xpr(scp, true)")
							r.Add("if i0Xpr != nil {")
							r.Addf("i0 = i0Xpr.(%v) // nil indicates get only", tm.Fld.Typ.Bse().IfcXpr.Ref(x))
							r.Add("}")
							r.Add("x.SkpSpceCmnt()")
							r.Add("x.NextRprn()")
							r.Addf("return x.%vXpr(scp, &%v{Trm: %vTrm, X: X, I0: i0})", tm.Fld.Typ.PkgTypTitle(), tm.Fld.SetGetXpr.Ref(x), tm.Fld.Camel())
						}
						r.Add("}")
						r.Add("x.Scn = scn // rewind")
					} else if rune(m)+'a' == 'a' { // asn
						x.Import(_sys.Lng.Pro.Trm.Prs)
						r.Add("asnTrm, ok := x.Asn()")
						r.Add("if ok {")
						r.Add("x.NextLprn()")
						r.Add("idn, ok := x.IdnLit()")
						r.Add("if !ok {")
						r.Addf("x.Panicf(\"%v asn xpr: missing Idn parameter\")", rxr.Bse().Pkg.Title())
						r.Add("}")
						r.Add("x.SkpSpceCmnt()")
						r.Add("lclScn := x.Scn")
						r.Add("var depth int")
						r.Add("depthTrm, ok := x.UntLit() // optional")
						r.Add("if ok {")
						r.Add("depth = int(prs.UntTrm(depthTrm, x.Txt))")
						r.Add("if depth == 0 {")
						r.Addf("x.Panicf(\"%v asn xpr: depth may not equal zero\")", rxr.Bse().Pkg.Title())
						r.Add("}")
						r.Add("for d := depth; scp != nil && d != 0; d-- { // recurse up scp by depth")
						r.Add("scp = scp.Prnt")
						r.Add("}")
						r.Add("if scp == nil {")
						r.Addf("x.Panicf(\"%v asn xpr: depth traversal too deep (depth:%%v)\", depth)", rxr.Bse().Pkg.Title())
						r.Add("}")
						r.Add("} else {")
						r.Add("x.Scn = lclScn")
						r.Add("}")
						r.Add("x.NextRprn()")
						r.Addf("scp.Decl(x.Txt[idn.Idx:idn.Lim], %v)", rxr.Knd.Ref(x))
						r.Addf("return x.%v(scp, &%v{Trm: asnTrm, X: X, Idn: idn.Bnd, Depth: depth})", r.Name, rxr.AsnXpr.Ref(x))
						r.Add("}")
						r.Add("x.Scn = scn")
					} else if rune(m)+'a' == 'e' { // each
						if rxr.EachXpr != nil {
							r.Add("eachTrm, ok := x.Each()")
							r.Add("if ok {")
							r.Add("x.NextLprn()")
							r.Add("idn, ok := x.IdnLit()")
							r.Add("if !ok {")
							r.Addf("x.Panicf(\"%v 'each' xpr: missing idn parameter\")", rxr.Bse().Pkg.Title())
							r.Add("}")
							r.Add("x.SkpSpceCmnt()")
							r.Add("eachScp := NewScp(scp)")
							// sys.Log("---", rxr.Title(), rxr.elm, rxr.IsArr())
							if rxr.IsArr() {
								r.Addf("eachScp.Decl(x.Txt[idn.Idx:idn.Lim], %v)", rxr.PrtArr().Arr.Elm.Knd.Ref(x))
							} else {
								r.Addf("eachScp.Decl(x.Txt[idn.Idx:idn.Lim], %v)", rxr.elm.Knd.Ref(x))
							}
							r.Addf("var xprs []%v", _sys.Lng.Pro.Xpr.Ifc.Ref(x))
							r.Add("for {")
							r.Add("cur := x.Xpr(eachScp, true)")
							r.Add("if cur == nil {")
							r.Add("break")
							r.Add("}")
							r.Add("xprs = append(xprs, cur)")
							r.Add("}")
							r.Add("x.NextRprn()")
							r.Addf("return &%v{Trm: eachTrm, X: X, Idn: idn.Bnd, Xprs: xprs, Scp: eachScp}", rxr.EachXpr.Ref(x))
							r.Add("}")
							r.Add("x.Scn = scn")
						}
						if rxr.ElseXpr != nil {
							r.Add("elseTrm, ok := x.Else()")
							r.Add("if ok {")
							r.Add("x.NextLprn()")
							r.Add("elseScp := NewScp(scp)")
							r.Addf("var xprs []%v", _sys.Lng.Pro.Xpr.Ifc.Ref(x))
							r.Add("for {")
							r.Add("cur := x.Xpr(elseScp, true)")
							r.Add("if cur == nil {")
							r.Add("break")
							r.Add("}")
							r.Add("xprs = append(xprs, cur)")
							r.Add("}")
							r.Add("x.NextRprn()")
							r.Addf("return  x.%vXpr(scp, &%v{Trm: elseTrm, X: X, Xprs: xprs, Scp: elseScp})", _sys.Bsc.Bol.Typ().PkgTypTitle(), rxr.ElseXpr.Ref(x))
							r.Add("}")
							r.Add("x.Scn = scn")
						}
					} else if rune(m)+'a' == 'p' { // pllEach
						r.Add("pllEachTrm, ok := x.PllEach()")
						r.Add("if ok {")
						r.Add("x.NextLprn()")
						r.Add("idn, ok := x.IdnLit()")
						r.Add("if !ok {")
						r.Addf("x.Panicf(\"%v 'pllEach' xpr: missing idn parameter\")", rxr.Bse().Pkg.Title())
						r.Add("}")
						r.Add("x.SkpSpceCmnt()")
						r.Add("pllEachScp := NewScp(scp)")
						r.Addf("pllEachScp.Decl(x.Txt[idn.Idx:idn.Lim], %v)", rxr.PrtArr().Arr.Elm.Knd.Ref(x))
						r.Addf("var xprs []%v", _sys.Lng.Pro.Xpr.Ifc.Ref(x))
						r.Add("for {")
						r.Add("cur := x.Xpr(pllEachScp, true)")
						r.Add("if cur == nil {")
						r.Add("break")
						r.Add("}")
						r.Add("xprs = append(xprs, cur)")
						r.Add("}")
						r.Add("x.NextRprn()")
						r.Addf("return &%v{Trm: pllEachTrm, X: X, Idn: idn.Bnd, Xprs: xprs, Scp: pllEachScp}", rxr.PllEachXpr.Ref(x))
						r.Add("}")
						r.Add("x.Scn = scn")
					} else if rune(m)+'a' == 't' { // then
						r.Add("thenTrm, ok := x.Then()")
						r.Add("if ok {")
						r.Add("x.NextLprn()")
						r.Add("thenScp := NewScp(scp)")
						r.Addf("var xprs []%v", _sys.Lng.Pro.Xpr.Ifc.Ref(x))
						r.Add("for {")
						r.Add("cur := x.Xpr(thenScp, true)")
						r.Add("if cur == nil {")
						r.Add("break")
						r.Add("}")
						r.Add("xprs = append(xprs, cur)")
						r.Add("}")
						r.Add("x.NextRprn()")
						r.Addf("return x.%vXpr(scp, &%v{Trm: thenTrm, X: X, Xprs: xprs, Scp: thenScp})", _sys.Bsc.Bol.Typ().PkgTypTitle(), rxr.ThenXpr.Ref(x))
						r.Add("}")
						r.Add("x.Scn = scn")
					}
				}
			}
		}
		r.Add("}")
	}
	r.Addf("x.Panicf(\"%v: no expression found\")", r.Name)
	r.Add("}")
	r.Add("return X")
	return r
}

func (x *FleXprr) PrsPkgXpr(p *Pkg) (r *TypFn) {
	r = x.TypFnf("%vPkgXpr", p.Title())
	r.InPrm(_sys.Lng.Pro.Xpr.Scp, "scp")
	r.OutPrm(_sys.Lng.Pro.Xpr.Ifc, "r")
	r.OutPrm(Bool, "ok")
	r.Addf("_, ok = x.%v()", p.Trm.Name)
	r.Add("if ok {")
	hasMem := false // determine if pkg has any xpr members
	for _, f := range p.Fles {
		// for _, t := range f.Bse().Typs {
		// 	if t.Bse().IsStruct() && t.Bse().IsXpr() { // check for struct
		// 		hasMem = true
		// 	}
		// }
		if f.Bse().Cnsts.MayXpr() || f.Bse().Vars.MayXpr() || f.Bse().PkgFns.MayXpr() {
			hasMem = true
			break
		}
	}
	if hasMem {
		r.Add("if x.HasMem() {")
		r.Add("scn := x.Scn")
		abcFns := make([][]func(), 26) // gather fns by fst ch
		for _, f := range p.Fles {
			for _, c := range f.Bse().Cnsts { // pkg cnsts
				c := c // capture for closure
				if c.MayXpr() {
					idx := c.Lower()[0] - 'a'
					abcFns[idx] = append(abcFns[idx], func() {
						r.Addf("%vTrm, ok := x.%v()", c.Camel(), c.Name)
						r.Add("if ok {")
						r.Addf("return x.%vXpr(scp, &%v{Trm: %vTrm}), true", c.Typ.PkgTypTitle(), c.Xpr.Ref(x), c.Camel())
						r.Add("}")
						r.Add("x.Scn = scn // rewind")
					})
				}
			}
			for _, v := range f.Bse().Vars { // pkg vars
				v := v // capture for closure
				if v.MayXpr() {
					idx := v.Lower()[0] - 'a'
					abcFns[idx] = append(abcFns[idx], func() {
						r.Addf("%vTrm, ok := x.%v()", v.Camel(), v.Name)
						r.Add("if ok {")
						r.Addf("return x.%vXpr(scp, &%v{Trm: %vTrm}), true", v.Typ.PkgTypTitle(), v.Xpr.Ref(x), v.Camel())
						r.Add("}")
						r.Add("x.Scn = scn // rewind")
					})
				}
			}
			for _, fn := range f.Bse().PkgFns { // pkg fns
				fn := fn // capture for closure
				if fn.MayXpr() {
					idx := fn.Lower()[0] - 'a'
					abcFns[idx] = append(abcFns[idx], func() {
						fn := fn
						x.WriteInPrms(r, &fn.FnBse, false)
						r.Add("x.NextRprn()")
						r.Addf("return x.%vXpr(scp, cur), true", fn.OutTyp().PkgTypTitle())
						r.Add("}")
						r.Add("x.Scn = scn // rewind")
					})
				}
			}
		}
		r.Add("switch x.Ch {")
		for m, curFns := range abcFns {
			if len(curFns) > 0 {
				r.Addf("case '%v':", string(rune(m)+'a'))
				for _, fn := range curFns {
					fn()
				}
			}
		}
		r.Add("}")
		r.Add("}")
	}
	r.Addf("x.Panicf(\"%v: no expression found\")", r.Name)
	r.Add("}")
	r.Add("return nil, false")
	return r
}

func (x *FleXprr) WriteInPrms(prs *TypFn, fn *FnBse, isMem bool) {
	b := &strings.Builder{}
	prs.Addf("%vTrm, ok := x.%v()", fn.Camel(), fn.Name)
	prs.Add("if ok {")
	prs.Add("x.NextLprn()")
	if len(fn.InPrms) > 0 && fn.InPrms.Lst().IsVariadic() {
		for n, inPrm := range fn.InPrms {
			if n < fn.InPrms.LstIdx() {
				prs.Addf("i%v := x.Xpr(scp).(%v)", n, inPrm.Typ.Bse().IfcXpr.Ref(x))
			} else {
				x.Import("reflect")
				prs.Addf("var i%v []%v", n, inPrm.Typ.Bse().IfcXpr.Ref(x))
				prs.Addf("for {")
				prs.Addf("cur := x.Xpr(scp, true)")
				prs.Addf("if cur == nil {")
				prs.Addf("break")
				prs.Addf("}")
				prs.Addf("v, ok := cur.(%v)", inPrm.Typ.Bse().IfcXpr.Ref(x))
				prs.Addf("if !ok {")
				prs.Addf("x.Panicf(\"%v: %v: non %v in variadic parameters (actual:%%v)\", reflect.TypeOf(cur).Elem().Name())", prs.Name, fn.Name, inPrm.Typ.Bse().IfcXpr.Ref(x))
				prs.Addf("}")
				prs.Addf("i%v = append(i%[1]v, v)", n)
				prs.Addf("}")
			}
		}
		b.WriteString(fmt.Sprintf("cur := &%v{Trm: %vTrm", fn.Xpr.Ref(x), fn.Camel()))
		if isMem {
			b.WriteString(", X: X")
		}
		for n := range fn.InPrms {
			b.WriteString(fmt.Sprintf(", I%v: i%[1]v", n))
		}
		b.WriteString("}")
	} else {
		b.WriteString(fmt.Sprintf("cur := &%v{Trm: %vTrm", fn.Xpr.Ref(x), fn.Camel()))
		if isMem {
			b.WriteString(", X: X")
		}
		for n, inPrm := range fn.InPrms {
			b.WriteString(fmt.Sprintf(", I%v: x.Xpr(scp).(%v)", n, inPrm.Typ.Bse().IfcXpr.Ref(x)))
		}
		b.WriteString("}")
	}
	prs.Add(b.String())
}

func (x *FleXprr) HasMem() (r *TypFn) {
	r = x.TypFn("HasMem")
	r.OutPrm(Bool)
	r.Add("scn := x.Scn")
	r.Add("x.SkpSpceCmnt()")
	r.Add("if x.Ch == '.' {")
	r.Add("x.NextRune()")
	r.Add("x.SkpSpceCmnt()")
	r.Add("return true")
	r.Add("}")
	r.Add("x.Scn = scn")
	r.Add("return false")
	return r
}
func (x *FleXprr) NextLprn() (r *TypFn) {
	r = x.TypFn("NextLprn")
	r.Add("x.SkpSpceCmnt()")
	r.Add("if x.Ch != '(' {")
	r.Add("x.Panicf(\"xprr: missing opening parenthesis '('\")")
	r.Add("}")
	r.Add("x.NextRune()")
	r.Add("x.SkpSpceCmnt()")
	return r
}
func (x *FleXprr) NextRprn() (r *TypFn) {
	r = x.TypFn("NextRprn")
	r.Add("x.SkpSpceCmnt()")
	r.Add("if x.Ch != ')' {")
	r.Add("x.Panicf(\"xprr: missing closing parenthesis '('\")")
	r.Add("}")
	r.Add("x.NextRune()")
	r.Add("x.SkpSpceCmnt()")
	return r
}
func (x *FleXprr) Erf() (r *TypFn) {
	x.Import("sys/err")
	r = x.TypFn("Erf")
	r.InPrm(NewExt("string"), "format")
	r.InPrmVariadic(Interface, "args")
	r.OutPrmPtr(NewExt("err.Err"), "r")
	r.Add("return err.Fmt(\"%v\\nxpr pos: ln:%v col:%v ch:%q\\nxpr src: %v\", err.Fmt(format, args...), x.Ln, x.Col, x.Ch, x.Txt)")
	return r
}
func (x *FleXprr) Panicf() (r *TypFn) {
	x.Import("sys/err")
	r = x.TypFn("Panicf")
	r.InPrm(NewExt("string"), "format")
	r.InPrmVariadic(Interface, "args")
	r.Add("err.PanicXprf(\"%v\\nxpr pos: ln:%v col:%v ch:%q\\nxpr src: %v\", uint32(x.Ln), uint32(x.Col), x.Ch, err.Fmt(format, args...), x.Ln, x.Col, x.Ch, x.Txt)")
	return r
}
