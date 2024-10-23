package tpl

import (
	"fmt"
	"strings"
	"sys/k"
	"sys/tpl/atr"
	"sys/tpl/mod"
)

type (
	FleActr struct {
		FleBse
		Pkgs
		PllWaitAct *Struct
	}
)

func (x *DirAct) NewActr() (r *FleActr) {
	r = &FleActr{}
	r.Name = k.Actr
	r.Pkg = x.Pkg
	r.StructPtr(r.Name, atr.Test) // Actr is an expression parser
	r.AddFle(r)
	return r
}
func (x *FleActr) InitFld(s *Struct) {
	x.Import(_sys.Lng.Pro.Xpr)
	s.FldExt("xpr.Xprr") // use Ext to avoid Ptr mod
}
func (x *FleActr) InitPkgFn() {
	x.RunPkgFn()
}
func (x *FleActr) RunPkgFn() (r *PkgFn) {
	r = x.PkgFn(k.Run)
	r.InPrm(String, "txt")
	r.Addf("var actr %v", x.Title())
	r.Add("actr.Run(txt)")
	return r
}
func (x *FleActr) InitAct(actr *FleActr) {
	x.Import(_sys.Lng.Pro.Trm.Prs)
	x.ifc()
	x.PllWaitAct = x.pllWaitAct()
	x.Cmpl()
	x.Cmplf()
	x.Run()
	x.Runf()
	x.RunIfc()
	x.RunIfcf()
	x.RunRlt()
	x.RunHst()
	x.Reduce()
	x.CmplActs()
	x.CmplAct()
	for _, f := range _sys.Fles {
		for _, typ := range f.Bse().Typs {
			if _, isExt := typ.(*Ext); !isExt && typ.Bse().IfcAct != nil {
				x.CmplTypAct(typ)
			}
		}
	}
}
func (x *FleActr) ifc() (r *Struct) {
	r = x.Struct("IfcIfc", atr.None)
	r.Fld("X", _sys.Lng.Pro.Act.Ifc)
	act := x.TypFn(k.Act, r)
	act.Addf("x.Ifc()")
	fn := x.TypFn(k.Ifc, r)
	fn.OutPrm(Interface)
	fn.Add("return x.X.Ifc()")
	return r
}
func (x *FleActr) NewAct(name string, ret Typ) (r *Struct) {
	r = x.Struct(name, atr.None)
	act := x.TypFn(k.Act, r)
	act.Addf("x.%v()", ret.PkgTypTitle())
	fn := x.TypFn(k.Ifc, r)
	fn.OutPrm(Interface)
	fn.Addf("return x.%v()", ret.PkgTypTitle())
	for _, ifc := range ret.Bse().Ifcs {
		fn := x.TypFn(ifc.PkgTypTitle(), r)
		fn.OutPrm(ifc)
		fn.Addf("return x.%v()", ret.PkgTypTitle())
	}
	return r
}
func (x *FleActr) NewActf(format string, ret Typ, args ...interface{}) (r *Struct) {
	return x.NewAct(fmt.Sprintf(format, args...), ret)
}
func (x *FleActr) TypIfc(typ Typ) (r *Ifc) {
	if typ == Interface {
		typ.Bse().IfcAct = _sys.Lng.Pro.Act.Ifc
		return typ.Bse().IfcAct
	}
	var sig *MemSig
	r = x.Ifcf("%vAct", atr.None, typ.PkgTypTitle())
	typ.Bse().IfcAct = r
	x.MemSiga(k.Act, atr.None, r)
	sig = x.MemSiga(k.Ifc, atr.None, r)
	sig.OutPrm(Interface)
	sig = x.MemSiga(typ.PkgTypTitle(), atr.None, r) // PkgTypTitle needed to distinguish hst/rlt Stm
	sig.OutPrm(typ)
	for _, ifc := range typ.Bse().Ifcs {
		sig = x.MemSiga(ifc.PkgTypTitle(), atr.None, r)
		sig.OutPrm(ifc)
	}
	return r
}
func (x *FleActr) LitAct(ret Typ) (r *Struct) {
	bse := ret.Bse()
	r = x.NewActf("%vLit", ret, ret.PkgTypTitle())
	bse.LitAct = r
	r.Fld(strings.Title(k.Trm), ret.Bse().LitTrm)
	r.Fld(strings.Title(k.Txt), String)
	act := x.TypFn(ret.PkgTypTitle(), r)
	act.OutPrm(ret)
	act.Addf("return %v(x.Trm, x.Txt)", bse.PrsTrm.Ref(x))
	if x.Test != nil && Opt.IsTest() && bse.IsTestAct() {
		x.Test.ActLit(ret)
	}
	return r
}
func (x *FleActr) AsnAct(ret Typ) (r *Struct) {
	bse := ret.Bse()
	r = x.NewActf("%vAsn", ret, ret.PkgTypTitle())
	bse.AsnAct = r
	r.FldTyp(ret.Bse().Scp)
	r.Fld("X", bse.IfcAct)
	act := x.TypFn(ret.PkgTypTitle(), r)
	act.OutPrm(ret)
	act.Addf("x.Arr[x.Idx] = x.X.%v()", bse.PkgTypTitle())
	act.Add("return x.Arr[x.Idx]")
	if x.Test != nil && Opt.IsTest() && bse.IsTestAct() {
		x.Test.ActAsn(ret)
	}
	return r
}
func (x *FleActr) AcsAct(ret Typ) (r *Struct) {
	bse := ret.Bse()
	r = x.NewActf("%vAcs", ret, ret.PkgTypTitle())
	bse.AcsAct = r
	r.FldTyp(ret.Bse().Scp)
	act := x.TypFn(ret.PkgTypTitle(), r)
	act.OutPrm(ret)
	act.Add("return x.Arr[x.Idx]")
	if x.Test != nil && Opt.IsTest() && bse.IsTestAct() {
		x.Test.ActAcs(ret)
	}
	return r
}
func (x *FleActr) ThenAct(ret Typ) (r *Struct) {
	bse := ret.Bse()
	r = x.NewActf("%vThen", ret, ret.PkgTypTitle())
	bse.ThenAct = r
	r.Fld("X", bse.IfcAct)
	r.FldSlice("Acts", _sys.Lng.Pro.Act.Ifc)
	act := x.TypFn(ret.PkgTypTitle(), r)
	act.OutPrm(ret)
	act.Addf("v := x.X.%v()", bse.PkgTypTitle())
	act.Add("if v {")
	act.Add("for _, a := range x.Acts {")
	act.Add("a.Act()")
	act.Add("}")
	act.Add("}")
	act.Add("return v")
	if x.Test != nil && Opt.IsTest() && bse.IsTestAct() {
		x.Test.ActThen(ret)
	}
	return r
}
func (x *FleActr) ElseAct(ret Typ) (r *Struct) {
	bse := ret.Bse()
	r = x.NewActf("%vElse", ret, ret.PkgTypTitle())
	bse.ElseAct = r
	r.Fld("X", bse.IfcAct)
	r.FldSlice("Acts", _sys.Lng.Pro.Act.Ifc)
	act := x.TypFn(ret.PkgTypTitle(), r)
	act.OutPrm(ret)
	act.Addf("v := x.X.%v()", bse.PkgTypTitle())
	act.Add("if !v {")
	act.Add("for _, a := range x.Acts {")
	act.Add("a.Act()")
	act.Add("}")
	act.Add("}")
	act.Add("return v")
	if x.Test != nil && Opt.IsTest() && bse.IsTestAct() {
		x.Test.ActElse(ret)
	}
	return r
}
func (x *FleActr) EachAct(typ Typ, arr *Arr) (r *Struct) {
	bse := typ.Bse()
	r = x.NewActf("%vEach", arr, typ.PkgTypTitle())
	// sys.Log("FleActr.EachAct", r.Name)
	bse.EachAct = r
	r.FldTyp(arr.Elm.Scp)
	r.Fld("X", bse.IfcAct)
	r.FldSlice("Acts", _sys.Lng.Pro.Act.Ifc)
	act := x.TypFn(arr.PkgTypTitle(), r)
	act.OutPrm(arr)
	act.Addf("vs := x.X.%v()", bse.PkgTypTitle())
	act.Add("if vs != nil {")
	act.Add("for _, v := range *vs {")
	act.Add("x.Arr[x.Idx] = v // set cur elm to scp")
	act.Add("for _, a := range x.Acts {")
	act.Add("a.Act()")
	act.Add("}")
	act.Add("}")
	act.Add("}")
	act.Add("return vs")
	if x.Test != nil && Opt.IsTest() && bse.IsTestAct() {
		x.Test.ActEach(arr)
	}
	return r
}
func (x *FleActr) PllEachAct(ret *Arr) (r *Struct) {
	r = x.NewActf("%vPllEach", ret, ret.PkgTypTitle())
	ret.PllEachAct = r
	r.Fld("X", ret.IfcAct)
	r.Fld("Idn", _sys.Bsc.Bnd)
	r.Fld("Txt", String)
	r.Fld("ActScpPrnt", _sys.Lng.Pro.Act.Scp)
	r.Fld("XprScp", _sys.Lng.Pro.Xpr.Scp)
	r.FldSlice("Xprs", _sys.Lng.Pro.Xpr.Ifc)

	// seg
	seg := x.StructPtrf("%vSeg", atr.None, r.Name)
	seg.Fld("Val", ret.Elm)
	seg.Fld("Idn", _sys.Bsc.Bnd)
	seg.Fld("Txt", String)
	seg.Fld("ActScpPrnt", _sys.Lng.Pro.Act.Scp)
	seg.Fld("XprScp", _sys.Lng.Pro.Xpr.Scp)
	seg.FldSlice("Xprs", _sys.Lng.Pro.Xpr.Ifc)
	segAct := x.TypFn(k.Act, seg)
	segAct.Add("var actr Actr")
	segAct.Add("actr.Reset(x.Txt)")
	segAct.Add("scp := NewScp(x.XprScp, x.ActScpPrnt)")
	segAct.Addf("elmScp := scp.%v(x.Txt[x.Idn.Idx:x.Idn.Lim])", ret.Elm.PkgTypTitle())
	segAct.Add("acts := actr.Acts(scp, x.Xprs...)")
	segAct.Add("elmScp.Arr[elmScp.Idx] = x.Val // set cur elm to scp")

	segAct.Add("for _, a := range acts {")
	segAct.Add("a.Act()")
	segAct.Add("}")

	act := x.TypFn(ret.PkgTypTitle(), r)
	act.OutPrm(ret)
	act.Addf("vs := x.X.%v()", ret.PkgTypTitle())
	act.Add("segs := make([]sys.Act, len(*vs))")
	act.Add("for n, v := range *vs {")
	act.Addf("segs[n] = %v{", seg.Adr(x))
	act.Add("Val: v,")
	act.Add("Idn: x.Idn,")
	act.Add("Txt: x.Txt,")
	act.Add("ActScpPrnt: x.ActScpPrnt,")
	act.Add("XprScp: x.XprScp,")
	act.Add("Xprs: x.Xprs,")
	act.Add("}")
	act.Add("}")
	act.Add("sys.Run().Pll(segs...) // run segs in pll")
	act.Add("return vs")
	if x.Test != nil && Opt.IsTest() && ret.IsTestAct() {
		x.Test.ActPllEach(ret)
	}
	return r
}

func (x *FleActr) pllWaitAct() (r *Struct) {
	x.Import(_sys)
	r = x.Struct(k.PllWait, atr.None)
	r.Fld("Txt", String)
	r.Fld("ActScpPrnt", _sys.Lng.Pro.Act.Scp)
	r.Fld("XprScp", _sys.Lng.Pro.Xpr.Scp)
	r.FldSlice("Xprs", _sys.Lng.Pro.Xpr.Ifc)
	r.Fld("Wg", WaitGroupPtr)

	// seg
	seg := x.StructPtrf("%vSeg", atr.None, r.Name)
	seg.Fld("Txt", String)
	seg.Fld("ActScpPrnt", _sys.Lng.Pro.Act.Scp)
	seg.Fld("XprScp", _sys.Lng.Pro.Xpr.Scp)
	seg.Fld("Xpr", _sys.Lng.Pro.Xpr.Ifc)
	seg.Fld("Wg", WaitGroupPtr)
	segAct := x.TypFn(k.Act, seg)
	segAct.Add("defer x.Wg.Done()")
	segAct.Add("var actr Actr")
	segAct.Add("actr.Reset(x.Txt)")
	segAct.Add("scp := NewScp(x.XprScp, x.ActScpPrnt)")
	segAct.Add("act := actr.Act(scp, x.Xpr)")
	segAct.Add("act.Act()")

	act := x.TypFn(k.Act, r)
	act.Add("wg := &sync.WaitGroup{}")
	act.Add("segs := make([]sys.Act, len(x.Xprs))")
	act.Add("for n, xpr := range x.Xprs {")
	act.Addf("segs[n] = %v{", seg.Adr(x))
	act.Add("Txt: x.Txt,")
	act.Add("ActScpPrnt: x.ActScpPrnt,")
	act.Add("XprScp: x.XprScp,")
	act.Add("Xpr: xpr,")
	act.Add("Wg: wg,")
	act.Add("}")
	act.Add("}")
	act.Add("wg.Add(len(x.Xprs))")
	act.Add("sys.Run().Pll(segs...) // run segs in pll")
	act.Add("wg.Wait()")

	ifcFn := x.TypFn(k.Ifc, r) // for act interface compliance
	ifcFn.OutPrm(Interface)
	ifcFn.Add("return nil")

	return r
}

func (x *FleActr) FldGetAct(s *Struct, fld *Fld) (r *Struct) {
	outBse := fld.Typ.Bse()
	outBse.ActFldGets.AddFld(fld)
	r = x.NewActf("%v%vGet", fld.Typ, s.PkgTypTitle(), strings.Title(fld.Name))
	fld.GetAct = r
	r.Fld("X", s.IfcAct)
	act := x.TypFn(fld.Typ.PkgTypTitle(), r)
	act.OutPrm(fld.Typ)
	act.Addf("return x.X.%v().%v", s.PkgTypTitle(), fld.Title())
	if x.Test != nil && Opt.IsTest() && s.IsTestAct() && fld.IsTestAct() {
		x.Test.ActFldGet(s, fld)
	}
	return r
}
func (x *FleActr) FldSetGetAct(s *Struct, fld *Fld) (r *Struct) {
	outBse := fld.Typ.Bse()
	outBse.ActFldSetGets.AddFld(fld)
	r = x.NewActf("%v%vSetGet", fld.Typ, s.PkgTypTitle(), strings.Title(fld.Name))
	fld.SetGetAct = r
	r.Fld("X", s.IfcAct)
	r.Fldf("I0", fld.Typ.Bse().IfcAct)
	act := x.TypFn(fld.Typ.PkgTypTitle(), r)
	act.OutPrm(fld.Typ)
	act.Addf("v := x.X.%v()", s.PkgTypTitle())
	act.Add("if x.I0 != nil {")
	act.Addf("v.%v = x.I0.%v()", fld.Title(), fld.Typ.PkgTypTitle())
	act.Add("}")
	act.Addf("return v.%v", fld.Title())
	if x.Test != nil && Opt.IsTest() && s.IsTestAct() && fld.IsTestAct() {
		x.Test.ActFldSetGet(s, fld)
	}
	return r
}
func (x *FleActr) CnstAct(c *Cnst) (r *Struct) {
	bse := c.Typ.Bse()
	bse.ActCnsts.AddCnst(c)
	r = x.NewActf("%v%v", c.Typ, bse.Pkg.Title(), strings.Title(c.Name))
	c.Act = r
	act := x.TypFn(bse.PkgTypTitle(), r)
	act.OutPrm(c.Typ)
	act.Addf("return %v", c.Ref(x))
	if x.Test != nil && Opt.IsTest() && bse.IsTestAct() && c.IsTestAct() {
		x.Test.ActCnst(c)
	}
	return r
}
func (x *FleActr) VarAct(v *Var) (r *Struct) {
	bse := v.Typ.Bse()
	bse.ActVars.AddVar(v)
	r = x.NewActf("%v%v", v.Typ, bse.Pkg.Title(), strings.Title(v.Name))
	v.Act = r
	act := x.TypFn(bse.PkgTypTitle(), r)
	act.OutPrm(v.Typ)
	act.Addf("return %v", v.Ref(x))
	if x.Test != nil && Opt.IsTest() && bse.IsTestAct() && v.IsTestAct() {
		x.Test.ActVar(v)
	}
	return r
}
func (x *FleActr) PkgFnAct(fn *PkgFn) (r *Struct) {
	// sys.Log("~~~", fn.PkgTitle())
	bse := fn.OutTyp().Bse()
	bse.ActPkgFns.AddPkgFn(fn)
	r = x.NewActf("%v%v", fn.OutTyp(), bse.Pkg.Title(), strings.Title(fn.Name))
	fn.Act = r
	for n, prm := range fn.InPrms {
		if prm.IsVariadic() {
			r.FldSlice(fmt.Sprintf("I%v", n), prm.Typ.Bse().IfcAct)
		} else {
			r.Fld(fmt.Sprintf("I%v", n), prm.Typ.Bse().IfcAct)
		}
	}
	act := x.TypFn(bse.PkgTypTitle(), r)
	act.OutPrm(fn.OutTyp())
	act.Addf("return %v(%v)", fn.Ref(x), x.WriteCallInPrms(act, &fn.FnBse))
	// if fn.Name == "NewPrfms" {
	// 	sys.Log("  FleActr.PkgFnAct", fn.Pkg.Name, fn.Name)
	// }
	if x.Test != nil && Opt.IsTest() && fn.IsTestAct() {
		x.Test.ActPkgFn(fn)
	}
	return r
}
func (x *FleActr) TypFnAct(fn *TypFn) (r *Struct) {
	rxrBse := fn.Rxr.Typ.Bse()
	outBse := fn.OutTyp().Bse()
	outBse.ActTypFns.AddTypFn(fn)
	r = x.NewActf("%v%v", fn.OutTyp(), fn.Rxr.Typ.PkgTypTitle(), strings.Title(fn.Name))
	fn.Act = r
	r.Fld("X", rxrBse.IfcAct)
	for n, prm := range fn.InPrms {
		if prm.IsVariadic() {
			r.FldSlice(fmt.Sprintf("I%v", n), prm.Typ.Bse().IfcAct)
		} else {
			r.Fld(fmt.Sprintf("I%v", n), prm.Typ.Bse().IfcAct)
		}
	}
	act := x.TypFn(outBse.PkgTypTitle(), r)
	act.OutPrm(fn.OutTyp())
	act.Addf("return x.X.%v().%v(%v)", fn.Rxr.Typ.PkgTypTitle(), fn.Title(), x.WriteCallInPrms(act, &fn.FnBse))
	if x.Test != nil && Opt.IsTest() && fn.Rxr.IsTestAct() && fn.IsTestAct() {
		// sys.Log("+++", fn.Rxr.Typ.Title(), fn.Title())
		x.Test.ActTypFn(fn)
	}
	return r
}
func (x *FleActr) MemSigAct(fn *MemSig) (r *Struct) {
	outBse := fn.OutTyp().Bse()
	outBse.ActMemSigs.AddSig(fn)
	r = x.NewActf("%v%v", fn.OutTyp(), fn.Rxr.PkgTypTitle(), strings.Title(fn.Name))
	fn.Act = r
	r.Fld("X", fn.Rxr.Bse().IfcAct)
	for n, prm := range fn.InPrms {
		fld := r.Fld(fmt.Sprintf("I%v", n), prm.Typ.Bse().IfcAct)
		if prm.IsVariadic() {
			fld.Mod = mod.Slice
		}
	}
	act := x.TypFn(outBse.PkgTypTitle(), r)
	act.OutPrm(fn.OutTyp())
	act.Addf("return x.X.%v().%v(%v)", fn.Rxr.PkgTypTitle(), fn.Title(), x.WriteCallInPrms(act, &fn.FnBse))
	if x.Test != nil && Opt.IsTest() && fn.Rxr.IsTestAct() && fn.IsTestAct() {
		x.Test.ActMemSig(fn)
	}
	return r
}
func (x *FleActr) WriteCallInPrms(act *TypFn, fn *FnBse) string {
	b := &strings.Builder{}
	if fn.InPrms.Ok() {
		lstIdx := fn.InPrms.LstIdx()
		if fn.InPrms.Lst().IsVariadic() {
			prm := fn.InPrms.Lst()
			act.Addf("var i%v []%v", lstIdx, prm.Typ.Ref(x))
			act.Addf("for _, cur := range x.I%v {", lstIdx)
			if prm.Typ == Interface {
				act.Addf("i%v = append(i%v, cur.Ifc())", lstIdx, lstIdx)
			} else {
				act.Addf("i%v = append(i%v, cur.%v())", lstIdx, lstIdx, prm.Typ.PkgTypTitle())
			}
			act.Addf("}")
		}
		for n, prm := range fn.InPrms {
			if n != 0 {
				b.WriteRune(',')
			}
			if n == lstIdx && prm.IsVariadic() {
				b.WriteString(fmt.Sprintf("i%v...", lstIdx))
			} else {
				if prm.Typ == Interface {
					b.WriteString(fmt.Sprintf("x.I%v.Ifc()", n))
				} else {
					b.WriteString(fmt.Sprintf("x.I%v.%v()", n, prm.Typ.PkgTypTitle()))
				}
			}
		}
	}
	return b.String()
}

func (x *FleActr) Cmpl() (r *TypFn) {
	r = x.TypFn("Cmpl")
	r.InPrm(String, "txt")
	r.OutPrmSlice(_sys.Lng.Pro.Act.Ifc)
	r.Add("xprScp, xprs := x.Prs(txt)")
	r.Add("return x.Acts(NewScp(xprScp), xprs...)")
	return r
}
func (x *FleActr) Cmplf() (r *TypFn) {
	r = x.TypFn("Cmplf")
	r.InPrm(String, "format")
	r.InPrmVariadic(Interface, "args")
	r.OutPrmSlice(_sys.Lng.Pro.Act.Ifc)
	r.Add("xprScp, xprs := x.Prsf(format, args...)")
	r.Add("return x.Acts(NewScp(xprScp), xprs...)")
	return r
}
func (x *FleActr) Run() (r *TypFn) {
	r = x.TypFn("Run")
	r.InPrm(String, "txt")
	r.Add("for _, a := range x.Cmpl(txt) {")
	r.Add("a.Act()")
	r.Add("}")
	return r
}
func (x *FleActr) Runf() (r *TypFn) {
	r = x.TypFn("Runf")
	r.InPrm(String, "format")
	r.InPrmVariadic(Interface, "args")
	r.Add("for _, a := range x.Cmplf(format, args...) {")
	r.Add("a.Act()")
	r.Add("}")
	return r
}
func (x *FleActr) RunIfc() (r *TypFn) {
	r = x.TypFn("RunIfc")
	r.InPrm(String, "txt")
	r.OutPrmSlice(Interface, "r")
	r.Add("txt = x.Reduce(txt)")
	r.Add("var actr Actr // new instance for each cmpl assures lock-free access")
	r.Add("for _, a := range actr.Cmpl(txt) {")
	r.Add("r = append(r, a.(sys.Ifc).Ifc())")
	r.Add("}")
	r.Add("return r")
	return r
}
func (x *FleActr) RunIfcf() (r *TypFn) {
	r = x.TypFn("RunIfcf")
	r.InPrm(String, "format")
	r.InPrmVariadic(Interface, "args")
	r.OutPrmSlice(Interface, "r")
	r.Add("var actr Actr // new instance for each cmpl assures lock-free access")
	r.Add("return actr.RunIfc(fmt.Sprintf(format, args...))")
	return r
}
func (x *FleActr) RunRlt() (r *TypFn) {
	x.Import("strings")
	x.Import(_sys)
	r = x.TypFn("RunRlt")
	r.InPrm(String, "txt")
	r.OutPrmSlice(Interface)
	r.Add("var actr Actr // new instance for each cmpl assures lock-free access. and no over-write of on-going cmpl for rlt")
	r.Add("return actr.RunIfc(strings.Replace(txt, \"hst.\", \"rlt.\", -1))")
	return r
}
func (x *FleActr) RunHst() (r *TypFn) {
	x.Import("strings")
	r = x.TypFn("RunHst")
	r.InPrm(String, "txt")
	r.OutPrmSlice(Interface)
	r.Add("var actr Actr // new instance for each cmpl assures lock-free access. and no over-write of on-going cmpl for rlt")
	r.Add("return actr.RunIfc(strings.Replace(txt, \"rlt.\", \"hst.\", -1))")
	return r
}
func (x *FleActr) Reduce() (r *TypFn) {
	x.Import("regexp")
	x.Import("fmt")
	r = x.TypFn("Reduce")
	r.InPrm(String, "txt")
	r.OutPrm(String)
	r.Add("type Node struct {")
	r.Add("Xpr string")
	r.Add("Idn string")
	r.Add("}")
	r.Add("// FIND ALL COMMON SUB-EXPRESSIONS")
	r.Add("// ASSIGN EACH SUB-EXPRESSION TO A SINGLE VARIABLE")
	r.Add("// TO AVOID DUPLICATE HST OR RLT CALCULATIONS")
	r.Add("var trmr trm.Trmr")
	r.Add("var sb strings.Builder")
	r.Add("var idnCnt uint32")
	r.Add("cur := &Node{Idn: \"hst\"}")
	r.Add("stck := append([]*Node{}, cur)")
	r.Add("for len(stck) != 0 {")
	r.Add("cur := stck[len(stck)-1]")
	r.Add("stck = stck[:len(stck)-1]")
	r.Add("trmr.Reset(txt)")
	// r.Add("re := regexp.MustCompile(fmt.Sprintf(\"%v.[[:word:]]+[(][^)]*[)]\", cur.Idn))")
	// r.Add("xprs := re.FindAllString(txt, -1)")
	r.Add("curNodes := make(map[string]*Node)")
	r.Add("xprBnds := trmr.Prefixs(cur.Idn)")
	r.Add("xprs := make([]string, len(xprBnds))")
	r.Add("for n, xprBnd := range xprBnds {")
	r.Add("xprs[n] = txt[xprBnd.Idx:xprBnd.Lim]")
	r.Add("}")
	r.Add("for _, xpr := range xprs {")
	r.Add("if _, ok := curNodes[xpr]; !ok {")
	r.Add("idnCnt++")
	r.Add("node := &Node{Xpr: xpr, Idn: fmt.Sprintf(\"v%v\", idnCnt)}")
	r.Add("sb.WriteString(fmt.Sprintf(\"%v.asn(%v)\\n\", node.Xpr, node.Idn))")
	r.Add("reNode := regexp.MustCompile(regexp.QuoteMeta(node.Xpr))")
	r.Add("txt = reNode.ReplaceAllString(txt, node.Idn)")
	r.Add("stck = append(stck, node)")
	r.Add("curNodes[xpr] = node")
	r.Add("}")
	r.Add("}")
	r.Add("}")
	r.Add("sb.WriteString(txt)")
	r.Add("return sb.String()")
	return r
}
func (x *FleActr) CmplActs() (r *TypFn) {
	r = x.TypFn("Acts")
	r.InPrm(_sys.Lng.Pro.Act.Scp, "scp")
	r.InPrmVariadic(_sys.Lng.Pro.Xpr.Ifc, "vs")
	r.OutPrmSlice(_sys.Lng.Pro.Act.Ifc, "r")
	r.Add("for _, v := range vs {")
	r.Add("r = append(r, x.Act(scp, v))")
	r.Add("}")
	r.Add("return r")
	return r
}
func (x *FleActr) CmplAct() (r *TypFn) {
	x.Import("reflect")
	r = x.TypFn("Act")
	r.InPrm(_sys.Lng.Pro.Act.Scp, "scp")
	r.InPrm(_sys.Lng.Pro.Xpr.Ifc, "v")
	r.OutPrm(_sys.Lng.Pro.Act.Ifc)
	r.Add("switch X := v.(type) {")
	r.Addf("case *%v:", _sys.Lng.Pro.Xpr.Xprr.PllWaitXpr.Ref(x)) // pllWait
	r.Addf("return %v{Txt: x.Txt, ActScpPrnt: scp, XprScp: X.Scp, Xprs: X.Xprs}", x.PllWaitAct.Ref(x))
	for _, p := range x.Pkgs {
		for _, f := range p.Fles {
			if f.Typ() != nil && f.Typ().Bse().IsAct() {
				if _, isExt := f.Typ().(*Ext); !isExt {
					r.Addf("case %v:", f.Typ().Bse().IfcXpr.Ref(x))
					r.Addf("return x.%vAct(scp, X)", f.Typ().PkgTypTitle())
				}
			}
		}
	}
	r.Add("}")
	r.Add("panic(x.Erf(\"Act: no action found %v\", reflect.ValueOf(v).Elem().Type().Name()))")
	return r
}

func (x *FleActr) CmplTypAct(typ Typ) (r *TypFn) {
	bse := typ.Bse()
	x.Import("reflect")
	r = x.TypFnf("%vAct", typ.PkgTypTitle())
	r.InPrm(_sys.Lng.Pro.Act.Scp, "scp")
	r.InPrm(bse.IfcXpr, "v")
	r.OutPrm(bse.IfcAct)
	r.Add("switch X := v.(type) {")
	b := &strings.Builder{}
	wrtCases := func(bse *TypBse) {
		if bse.LitAct != nil {
			r.Addf("case *%v:", bse.LitXpr.Ref(x))
			r.Addf("return %v{Trm: X.Trm, Txt: x.Txt}", bse.LitAct.Ref(x))
		}
		if bse.AsnAct != nil {
			r.Addf("case *%v:", bse.AsnXpr.Ref(x))
			r.Add("asnScp := scp")
			r.Add("if X.Depth != 0 {")
			r.Add("for d := X.Depth; asnScp != nil && d != 0; d-- { // recurse up scp by depth")
			r.Add("asnScp = asnScp.Prnt")
			r.Add("}")
			r.Add("}")
			r.Addf("return %v{%v: asnScp.%[3]v(x.Txt[X.Idn.Idx:X.Idn.Lim]), X: x.%[3]vAct(scp, X.X)}", bse.AsnAct.Ref(x), bse.Scp.Title(), bse.PkgTypTitle())
			r.Addf("case *%v:", bse.AcsXpr.Ref(x))
			r.Addf("return %v{%v: scp.%v(x.Txt[X.Trm.Idx:X.Trm.Lim])}", bse.AcsAct.Ref(x), bse.Scp.Title(), bse.PkgTypTitle())
			if arr, ok := typ.(*Arr); ok {
				if bse.EachAct != nil {
					r.Addf("case *%v:", bse.EachXpr.Ref(x))
					r.Add("eachScp := NewScp(X.Scp, scp)")
					r.Addf("return %v{X: x.%vAct(scp, X.X), %v: eachScp.%v(x.Txt[X.Idn.Idx:X.Idn.Lim]), Acts: x.Acts(eachScp, X.Xprs...)}",
						bse.EachAct.Ref(x), bse.PkgTypTitle(), arr.Elm.Scp.Title(), arr.Elm.PkgTypTitle())
				}
				if bse.PllEachAct != nil {
					r.Addf("case *%v:", bse.PllEachXpr.Ref(x))
					r.Addf("return %v{X: x.%vAct(scp, X.X), Txt: x.Txt, Idn: X.Idn, ActScpPrnt: scp, XprScp: X.Scp, Xprs: X.Xprs}", bse.PllEachAct.Ref(x), bse.PkgTypTitle())
				}
			}
		}
		if bse.ThenAct != nil {
			r.Addf("case *%v:", bse.ThenXpr.Ref(x))
			r.Addf("return %v{X: x.%vAct(scp, X.X), Acts: x.Acts(NewScp(X.Scp, scp), X.Xprs...)}",
				bse.ThenAct.Ref(x), bse.PkgTypTitle())
		}
		if bse.ElseAct != nil {
			r.Addf("case *%v:", bse.ElseXpr.Ref(x))
			r.Addf("return %v{X: x.%vAct(scp, X.X), Acts: x.Acts(NewScp(X.Scp, scp), X.Xprs...)}",
				bse.ElseAct.Ref(x), bse.PkgTypTitle())
		}
		for _, f := range bse.ActFldGets {
			r.Addf("case *%v:", f.GetXpr.Ref(x))
			r.Addf("return %v{X: x.%vAct(scp, X.X)}", f.GetAct.Ref(x), f.Struct.PkgTypTitle())
		}
		for _, f := range bse.ActFldSetGets {
			r.Addf("case *%v:", f.SetGetXpr.Ref(x))
			r.Add("if X.I0 == nil {")
			r.Addf("return %v{X: x.%vAct(scp, X.X)}", f.SetGetAct.Ref(x), f.Struct.PkgTypTitle())
			r.Add("} else {")
			r.Addf("return %v{X: x.%vAct(scp, X.X), I0: x.%vAct(scp, X.I0)}", f.SetGetAct.Ref(x), f.Struct.PkgTypTitle(), f.Typ.PkgTypTitle())
			r.Add("}")
		}
		for _, c := range bse.ActCnsts {
			r.Addf("case *%v:", c.Xpr.Ref(x))
			r.Addf("return %v{}", c.Act.Ref(x))
		}
		for _, v := range bse.ActVars {
			r.Addf("case *%v:", v.Xpr.Ref(x))
			r.Addf("return %v{}", v.Act.Ref(x))
		}
		for _, fn := range bse.ActPkgFns {
			r.Addf("case *%v:", fn.Xpr.Ref(x))
			r.Addf("return %v{%v}", fn.Act.Ref(x), x.WriteCmplInPrms(r, &fn.FnBse, b))
		}
		for _, fn := range bse.ActTypFns {
			r.Addf("case *%v:", fn.Xpr.Ref(x))
			r.Addf("return %v{X: x.%vAct(scp, X.X), %v}", fn.Act.Ref(x), fn.Rxr.Typ.PkgTypTitle(), x.WriteCmplInPrms(r, &fn.FnBse, b))
		}
		for _, fn := range bse.ActMemSigs {
			r.Addf("case *%v:", fn.Xpr.Ref(x))
			r.Addf("return %v{X: x.%vAct(scp, X.X), %v}", fn.Act.Ref(x), fn.Rxr.PkgTypTitle(), x.WriteCmplInPrms(r, &fn.FnBse, b))
		}
	}
	wrtCases(bse)
	if _, ok := typ.(*Ifc); ok {
		for _, concrete := range bse.ConcreteTyps {
			wrtCases(concrete.Bse())
		}
	}
	var usedX bool // check for X var usage; possible not used. Go compiler will complain if not used
	for n := 1; n < len(r.Lines); n++ {
		if strings.Index(r.Lines[n], "X.") > -1 {
			usedX = true
			break
		}
	}
	if !usedX {
		r.Lines[0] = "switch v.(type) {"
	}
	r.Add("}")
	r.Addf("panic(x.Erf(\"%v: no action found %%v\", reflect.ValueOf(v).Elem().Type().Name()))", r.Name)
	return r
}
func (x *FleActr) WriteCmplInPrms(cmpl *TypFn, fn *FnBse, b *strings.Builder) string {
	b.Reset()
	if fn.InPrms.Ok() {
		lstIdx := fn.InPrms.LstIdx()
		if fn.InPrms.Lst().IsVariadic() {
			prm := fn.InPrms.Lst()
			cmpl.Addf("var i%v []%v", lstIdx, prm.Typ.Bse().IfcAct.Ref(x))
			cmpl.Addf("for _, cur := range X.I%v {", lstIdx)
			if prm.Typ == Interface {
				cmpl.Addf("i%v = append(i%v, x.Act(scp, cur))", lstIdx, lstIdx)
			} else {
				cmpl.Addf("i%v = append(i%v, x.%vAct(scp, cur))", lstIdx, lstIdx, prm.Typ.PkgTypTitle())
			}
			cmpl.Addf("}")
		}
		for n, prm := range fn.InPrms {
			if n != 0 {
				b.WriteRune(',')
			}
			if n == lstIdx && prm.IsVariadic() {
				b.WriteString(fmt.Sprintf("I%v: i%v", lstIdx, lstIdx))
			} else {
				if prm.Typ == Interface {
					b.WriteString(fmt.Sprintf("I%v: x.Act(scp, X.I%v)", n, n))
				} else {
					b.WriteString(fmt.Sprintf("I%v: x.%vAct(scp, X.I%v)", n, prm.Typ.PkgTypTitle(), n))
				}
			}
		}
	}
	return b.String()
}

// func WriteActrInPrms(b *strings.Builder, inPrms InPrms, lines *Lines, f *act.Fle, fActr *Fle) {
// 	b.Reset()
// 	if len(inPrms) > 0 && inPrms.Lst().IsVariadic() {
// 		n := inPrms.LstIdx()
// 		prm := inPrms.Lst()
// 		flePrm := f.Lvl.Map[prm.Typ.PkgTypTitle()]
// 		lines.Addf("var i%v []%v", n, flePrm.IfcAct.Ref(fActr))
// 		lines.Addf("for _, cur := range X.I%v {", n)
// 		lines.Addf("i%v = append(i%v, x.%vAct(scp, cur))", n, n, prm.Typ.PkgTypTitle())
// 		lines.Addf("}")
// 	}
// 	for n, prm := range inPrms {
// 		if n > 0 {
// 			b.WriteString(", ")
// 		}
// 		if prm.IsVariadic() {
// 			b.WriteString(fmt.Sprintf("I%v: i%v", n, n))
// 		} else {
// 			b.WriteString(fmt.Sprintf("I%v: x.%vAct(scp, X.I%v)", n, prm.Typ.PkgTypTitle(), n))
// 		}
// 	}
// }
