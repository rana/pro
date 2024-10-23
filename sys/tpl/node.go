package tpl

import (
	"fmt"
	"strings"
	"sys/k"
	"sys/tpl/atr"
)

type (
	FleNodeBse struct {
		FleBse
	}
	FleNode interface {
		NodeBse() *FleNodeBse
	}
)

func (x *FleNodeBse) NodeBse() *FleNodeBse { return x }
func (x *FleNodeBse) InitTyp(bse *TypBse) {
	x.bse = x.StructPtr(bse.Name+"Bse", atr.TypAnaBse)
	x.bse.Fle = x
	x.bse.ifc, _ = x.Typ().(*Ifc)
	x.bse.ifc.bse = x.bse
	bse.bse = x.bse
}
func (x *FleNodeBse) InitFld(s *Struct) {
	x.bse.Fld("Slf", x.Typ()).Atr = atr.Slf | atr.TstSkp
}
func (x *FleNodeBse) InitIfc(i *Ifc) {
	if i != nil { // IFC MEMS ADDED FROM NODE CALL
		// i.EmbedIfc(_sys.Ana.Pth)
		x.EmbedIfc(i, _sys.Ana.Pth.Typ().(*Ifc))
	}
}
func (x *FleNodeBse) InitTypFn() {
	x.FnBse()
}
func (x *FleNodeBse) FnBse() (r *TypFn) {
	r = x.TypFna(k.Bse, atr.None, x.bse)
	r.OutPrm(x.bse)
	r.Add("return x")
	x.MemSigFn(r) // add to interface
	return r
}

func (x *FleNodeBse) NodePkgFn(name, family string, out FleOrTyp, initPrms func(r *PkgFn)) (r *PkgFn) {
	outTyp := GetTyp(out).Bse()
	r = x.PkgFn(name)
	r.OutPrm(outTyp)
	r.Family = strings.Replace(outTyp.bse.Name, "Bse", "", 1)
	if family != "" {
		r.Family = fmt.Sprintf("%v%v", r.Family, strings.Title(family))
	}
	// r.Node = bse.Fle.Bse().StructPtrf("%v%v", atr.TypRoot, bse.ifc.Name, strings.Title(name))
	r.Node = outTyp.Fle.Bse().StructPtrf("%v%v", atr.Tst, r.Family, strings.Title(name))
	r.Node.Fle = outTyp.Fle
	r.Node.FldBse(outTyp.bse)
	initPrms(r) // initPrms: call after Node assignment to allow possible Prnt assignment

	x.NodePreamble(r.Node, r.InPrms, &r.Lines, false)

	outTyp.Fle.Bse().NodeFns(r, r.Node)
	testPth := x.Typ().Bse().TestPth
	if testPth != nil {
		testPth = testPth[:len(testPth)-1]
	}
	outTyp.Test.GenPkgFn(r, testPth)
	return r
}

func (x *FleNodeBse) ElmNodeTypFn(name, family, cnj string, out FleOrTyp, initPrms func(r *TypFn), rxr ...Typ) (r *TypFn) {
	var rxrTyp Typ
	if len(rxr) != 0 {
		rxrTyp = rxr[0]
	} else {
		rxrTyp = GetTyp(x)
		if rxrTyp.Bse().bse != nil {
			rxrTyp = rxrTyp.Bse().bse
		}
	}
	outTyp := GetTyp(out).Bse()
	r = x.TypFnRxrf("%v%v", rxrTyp, strings.Title(cnj), strings.Title(name))
	r.OutPrm(outTyp)
	if outTyp.bse != nil {
		r.Family = strings.Replace(outTyp.bse.Name, "Bse", "", 1)
	}
	if family != "" {
		r.Family = fmt.Sprintf("%v%v", r.Family, strings.Title(family))
	}
	r.Cnj = strings.Title(cnj)

	// node
	r.Node = outTyp.Fle.Bse().StructPtrf("%v%v", atr.Tst|atr.Test, r.Family, strings.Title(name))
	r.Node.Fle = outTyp.Fle
	if outTyp.bse != nil {
		r.Node.FldBse(outTyp.bse)
	}
	initPrms(r) // initPrms: call after Node assignment to allow possible Prnt assignment
	if len(rxr) == 0 {
		x.MemSigFn(r) // add to interface (after prms defined)
	}
	// preamble
	x.NodePreamble(r.Node, r.InPrms, &r.Lines, true)
	// preamble: rlt
	if fr, ok := outTyp.Bse().Fle.(FleRlt); ok {
		x.Import(_sys)
		r.Add("r.Id = sys.NextID()")
		if fr.RltBse().Rxs != nil {
			r.Addf("r.Rxs = %v", fr.RltBse().Rxs.Make(x))
		}
	}
	outTyp.Bse().Fle.Bse().NodeFns(r, r.Node)

	funcIfc := _sys.Tst.NodeFuncIfc(r) // create func type
	_sys.Tst.NodeVar(r, funcIfc)
	if outTyp.Test != nil && r.Node.Prnt() != nil && r.Node.IsTest() { // initPrms may remove test, i.e., CndMl
		// sys.Logf("%v outTyp:%v Node.Prnt:%v", r.Name, outTyp.Title(), r.Node.Prnt().Typ.Title())
		// outTyp.Test.Gen(r, append(r.Node.Prnt().Typ.Bse().TestPth, r.Node.TestPth...)) // initPrms may set Node.TestPth
		var tstStps []*TestStp
		if len(r.Node.TestPth) == 0 {
			tstStps = r.Node.Prnt().Typ.Bse().TestPth
		} else {
			tstStps = append(tstStps, r.Node.Prnt().Typ.Bse().TestPth...)
			tstStps = append(tstStps, r.Node.TestPth...)
		}
		outTyp.Test.Gen(r, tstStps)
	}

	// if outTyp.fbr != nil {
	// 	// FBR NEW: ELM->FBR: ELM FN MUST HAVE PRMS TO VARIATE
	// 	mayFbr := true
	// 	if len(r.InPrms) != 0 {
	// 		for _, inPrm := range r.InPrms {
	// 			if inPrm.Typ.Bse().arr == nil {
	// 				mayFbr = false
	// 				break
	// 			}
	// 		}
	// 		if mayFbr {
	// 			x.FbrNodeTypFn(r)
	// 		}
	// 	}
	// 	// FBR CONTINUATION: FBR->FBR: RXR & OUT HAVE FBR
	// 	if mayFbr && rxrTyp.Bse().ifc != nil && rxrTyp.Bse().ifc.fbr != nil && r.Node.Prnt() != nil && r.Node.Prnt().Typ != _sys.Ana.Hst.Prv.Typ() {
	// 		x.FbrContNodeTypFn(r, rxrTyp.Bse().ifc.fbr, outTyp.fbr)
	// 	}
	// }

	return r
}

// func (x *FleNodeBse) FbrNodeTypFn(elmNodeFn *TypFn) (r *TypFn) {
// 	outFbr := elmNodeFn.OutTyp().Bse().fbr
// 	x.Import(_sys)
// 	r = x.TypFnRxrf("%vs", elmNodeFn.Rxr.Typ, elmNodeFn.Name)
// 	if len(elmNodeFn.InPrms) != 0 {
// 		r.InPrmArr(elmNodeFn.InPrms...)
// 		r.InPrms.Lst().Mod = elmNodeFn.InPrms.Lst().Mod // in case variadic
// 	}
// 	r.OutPrm(outFbr)
// 	r.Family = elmNodeFn.Family + "Fbr"
// 	r.Cnj = elmNodeFn.Cnj
// 	x.MemSigFn(r) // add to interface (after prms defined)

// 	r.Node = outFbr.Fle.Bse().StructPtrf("%v%v", atr.Tst, r.Family, r.Name)
// 	r.Node.Fle = outFbr.Fle
// 	r.Node.FldBse(outFbr.bse)
// 	r.Node.FldPrnt(x)

// 	// cmd: for pll
// 	cmd := elmNodeFn.Fle.Bse().StructPtrf("%v%vFbrCmd", atr.None, elmNodeFn.Rxr.Typ.Title(), elmNodeFn.Title())
// 	cmd.Fld("Rxr", elmNodeFn.Rxr.Typ)
// 	for _, p := range elmNodeFn.InPrms {
// 		fld := cmd.Fld(p.Title(), p.Typ)
// 		if p.IsVariadic() {
// 			fld.Mod = mod.Slice
// 		}
// 	}
// 	if len(elmNodeFn.OutPrms) != 0 && elmNodeFn.OutTyp().Bse().arr != nil {
// 		cmd.Fld("OutIdx", _sys.Bsc.Unt)
// 		cmd.Fld("Outs", elmNodeFn.OutTyp().Bse().arr)
// 	}
// 	fnAct := x.TypFn(k.Act, cmd)
// 	fnAct.Addf("(*x.Outs)[x.OutIdx] = x.Rxr%v", elmNodeFn.CallNode())

// 	// preamble
// 	x.NodePreamble(r.Node, r.InPrms, &r.Lines, true)
// 	// body
// 	if len(r.InPrms) == 0 {
// 		r.Addf("r.%v = %v()", outFbr.elm.arr.Camel(), outFbr.elm.arr.New.Ref(x))
// 	} else {
// 		var b strings.Builder // prm len
// 		for n, p := range r.InPrms {
// 			if n != 0 {
// 				b.WriteRune('*')
// 			}
// 			if !p.IsVariadic() {
// 				b.WriteString("len(*")
// 				b.WriteString(p.Name)
// 				b.WriteString(")")
// 			} else {
// 				r.Addf("%vCnt := 1", p.Name)
// 				r.Addf("if len(%v) != 0 {", p.Name)
// 				r.Addf("%[1]vCnt = len(%[1]v)", p.Name)
// 				r.Add("}")
// 				b.WriteString(fmt.Sprintf("%vCnt", p.Name))
// 			}
// 		}
// 		r.Addf("r.%v = %v(unt.Unt(%v))", outFbr.elm.arr.Camel(), outFbr.elm.arr.Make.Ref(x), b.String())
// 	}
// 	r.Addf("outIdx, acts := %v, make([]sys.Act, len(*r.%v))", _sys.Bsc.Unt.Zero.Ref(x), outFbr.elm.arr.Camel())
// 	for n := 0; n < len(r.InPrms)-1; n++ {
// 		r.Addf("for _, %v := range *%v {", elmNodeFn.InPrms[n].Name, r.InPrms[n].Name)
// 	}

// 	wrtCmd := func(skpVaridic bool) {
// 		r.Addf("cmd := %v{}", cmd.Adr(x))
// 		r.Add("cmd.Rxr = x")
// 		for _, p := range elmNodeFn.InPrms {
// 			if p.IsVariadic() {
// 				if !skpVaridic {
// 					r.Addf("cmd.%v = *%v", p.Title(), p.Camel())
// 				}
// 			} else {
// 				r.Addf("cmd.%v = %v", p.Title(), p.Camel())
// 			}
// 		}
// 		r.Add("cmd.OutIdx = outIdx")
// 		r.Addf("cmd.Outs = r.%[1]v", outFbr.elm.arr.Camel())
// 		r.Add("acts[outIdx] = cmd")
// 		r.Add("outIdx++")
// 	}
// 	lstPrm := r.InPrms.Lst()
// 	if !lstPrm.IsVariadic() {
// 		r.Addf("for _, %v := range *%v {", elmNodeFn.InPrms.Lst().Name, lstPrm.Name)
// 		wrtCmd(false)
// 		r.Add("}")
// 	} else {
// 		r.Addf("if len(%v) == 0 {", lstPrm.Name) //no variadic count
// 		wrtCmd(true)
// 		r.Add("} else {")
// 		r.Addf("for _, %v := range %v {", elmNodeFn.InPrms.Lst().Name, lstPrm.Name)
// 		wrtCmd(false)
// 		r.Add("}")
// 		r.Add("}")
// 	}
// 	for n := 0; n < len(r.InPrms)-1; n++ {
// 		r.Add("}")
// 	}
// 	r.Add("sys.Run().Pll(acts...)")
// 	r.Add("return r")
// 	// end
// 	outFbr.bse.Fle.Bse().NodeFns(r, r.Node)

// 	funcIfc := _sys.Tst.NodeFuncIfc(r) // create func type
// 	_sys.Tst.NodeVar(r, funcIfc)
// 	outFbr.Test.Gen(r, r.Node.Prnt().Typ.Bse().TestPth)

// 	return r
// }

// func (x *FleNodeBse) FbrContNodeTypFn(elmNodeFn *TypFn, rxrFbr, outFbr *Ifc) (r *TypFn) {
// 	// sys.Logf("-- FbrContNodeTypFn %v %v -> %v", fbrNodeFn.Name, rxrFbr.Name, outFbr.Name)
// 	rxrFbr.Fle.Bse().Import(_sys)
// 	r = rxrFbr.Fle.Bse().TypFn(elmNodeFn.Name, rxrFbr.bse)
// 	var fbrPrm *Prm
// 	for _, p := range *elmNodeFn.InPrms.Cpy(atr.None) { // COPY PRMS AND SET ATR; FOR StgyLong/Shrt REMOVE atr.FldSkp

// 		if elmNodeFn.Rxr.Typ.Bse().ifc != nil && elmNodeFn.Rxr.Typ.Bse().ifc == p.Typ && !p.IsVariadic() {
// 			fbrPrm = r.InPrm(rxrFbr, p.Name)
// 			fbrPrm.Atr = p.Atr
// 			fbrPrm.Mod = p.Mod
// 		} else {
// 			r.AddInPrm(p)
// 		}
// 	}
// 	// r.AddInPrm(*elmNodeFn.InPrms.Cpy(atr.None)...)
// 	r.OutPrm(outFbr)
// 	r.Family = elmNodeFn.Family + "Fbr"
// 	r.Cnj = elmNodeFn.Cnj
// 	rxrFbr.Fle.Bse().MemSigFn(r) // add to interface (after prms defined)

// 	r.Node = outFbr.bse.Fle.Bse().StructPtrf("%v%v", atr.Tst, r.Family, r.Name)
// 	r.Node.Fle = outFbr.Fle
// 	r.Node.FldBse(outFbr.bse)
// 	r.Node.FldPrnt(rxrFbr)

// 	// cmd: for pll
// 	cmd := rxrFbr.Fle.Bse().StructPtrf("%v%vCmd", atr.None, r.Rxr.Typ.Title(), elmNodeFn.Title())
// 	cmd.Fld("Rxr", elmNodeFn.Rxr.Typ)
// 	for _, p := range elmNodeFn.InPrms {
// 		fld := cmd.Fld(p.Title(), p.Typ)
// 		if p.IsVariadic() {
// 			fld.Mod = mod.Slice
// 		}
// 	}
// 	cmd.Fld("OutIdx", _sys.Bsc.Unt)
// 	cmd.Fld("Outs", outFbr.elm.arr)
// 	fnAct := rxrFbr.Fle.Bse().TypFn(k.Act, cmd)
// 	fnAct.Addf("(*x.Outs)[x.OutIdx] = x.Rxr%v", elmNodeFn.CallNode())

// 	// preamble
// 	x.NodePreamble(r.Node, r.InPrms, &r.Lines, true)
// 	// body

// 	if fbrPrm == nil { // no fbr prm
// 		r.Addf("r.%v = %v(x.%v.Cnt())", outFbr.elm.arr.Camel(), outFbr.elm.arr.Make.Ref(rxrFbr.Fle), rxrFbr.elm.arr.Camel())
// 		r.Addf("outIdx, acts := %v, make([]sys.Act, len(*r.%v))", _sys.Bsc.Unt.Zero.Ref(rxrFbr.Fle), outFbr.elm.arr.Camel())
// 		r.Addf("for _, %v := range *x.%v {", rxrFbr.elm.arr.Elm.Camel(), rxrFbr.elm.arr.Camel()) // fbr elm
// 		r.Addf("cmd := %v{}", cmd.Adr(x))
// 		r.Addf("cmd.Rxr = %v.Bse()", rxrFbr.elm.arr.Elm.Camel())
// 		for _, p := range elmNodeFn.InPrms {
// 			r.Addf("cmd.%v = %v", p.Title(), p.Camel())
// 		}
// 		r.Add("cmd.OutIdx = outIdx")
// 		r.Addf("cmd.Outs = r.%[1]v", outFbr.elm.arr.Camel())
// 		r.Add("acts[outIdx] = cmd")
// 		r.Add("outIdx++")
// 		r.Add("}") // end: fbr elm
// 	} else { // fbr prm
// 		r.Addf("lim := x.%v.Cnt().Min(%v.Bse().%v.Cnt())", rxrFbr.elm.arr.Camel(), fbrPrm.Name, rxrFbr.elm.arr.Camel())
// 		r.Addf("r.%v = %v(lim)", outFbr.elm.arr.Camel(), outFbr.elm.arr.Make.Ref(rxrFbr.Fle))
// 		r.Addf("outIdx, acts := %v, make([]sys.Act, lim)", _sys.Bsc.Unt.Zero.Ref(rxrFbr.Fle))
// 		r.Addf("for n := unt.Zero; n < lim; n++ {")
// 		r.Addf("cmd := %v{}", cmd.Adr(x))
// 		r.Addf("cmd.Rxr = (*x.%v)[n].Bse()", rxrFbr.elm.arr.Camel())
// 		for _, p := range elmNodeFn.InPrms {
// 			if p.Name == fbrPrm.Name {
// 				r.Addf("cmd.%v = (*%v.Bse().%v)[n]", p.Title(), fbrPrm.Name, rxrFbr.elm.arr.Camel())
// 			} else {
// 				r.Addf("cmd.%v = %v", p.Title(), p.Camel())
// 			}
// 		}
// 		r.Add("cmd.OutIdx = outIdx")
// 		r.Addf("cmd.Outs = r.%[1]v", outFbr.elm.arr.Camel())
// 		r.Add("acts[outIdx] = cmd")
// 		r.Add("outIdx++")
// 		r.Addf("}")
// 	}
// 	r.Add("sys.Run().Pll(acts...)")
// 	r.Add("return r")
// 	// end
// 	outFbr.bse.Fle.Bse().NodeFns(r, r.Node)

// 	funcIfc := _sys.Tst.NodeFuncIfc(r) // create func type
// 	_sys.Tst.NodeVar(r, funcIfc)
// 	outFbr.Test.Gen(r, r.Node.Prnt().Typ.Bse().TestPth)

// 	if outFbr.wve != nil {
// 		// WVE NEW: FBR->WVE: FBR FN MUST HAVE PRMS TO VARIATE
// 		if len(r.InPrms) != 0 {
// 			mayWve := true
// 			for _, inPrm := range r.InPrms {
// 				if inPrm.Typ.Bse().arr == nil {
// 					mayWve = false
// 					break
// 				}
// 			}
// 			if mayWve {
// 				outFbr.Fle.(FleNode).NodeBse().WveNodeTypFn(r, rxrFbr, outFbr)
// 			}
// 		}
// 	}

// 	return r
// }

// func (x *FleNodeBse) WveNodeTypFn(fbrNodeFn *TypFn, rxrFbr, outFbr *Ifc) (r *TypFn) {
// 	outWve := fbrNodeFn.OutTyp().Bse().wve
// 	rxrFbr.Fle.Bse().Import(_sys)
// 	r = rxrFbr.Fle.Bse().TypFnRxrf("%vs", fbrNodeFn.Rxr.Typ, fbrNodeFn.Name)
// 	if len(fbrNodeFn.InPrms) != 0 {
// 		r.InPrmArr(fbrNodeFn.InPrms...)
// 		r.InPrms.Lst().Mod = fbrNodeFn.InPrms.Lst().Mod // in case variadic
// 	}
// 	r.OutPrm(outWve)
// 	r.Family = fbrNodeFn.Family + "Wve"
// 	r.Cnj = fbrNodeFn.Cnj
// 	rxrFbr.Fle.Bse().MemSigFn(r) // add to interface (after prms defined)

// 	r.Node = outWve.Fle.Bse().StructPtrf("%v%v", atr.Tst, r.Family, r.Name)
// 	r.Node.Fle = outWve.Fle
// 	r.Node.FldBse(outWve.bse)
// 	r.Node.FldPrnt(rxrFbr)

// 	// cmd: for pll
// 	cmd := fbrNodeFn.Fle.Bse().StructPtrf("%v%vWveCmd", atr.None, fbrNodeFn.Rxr.Typ.Title(), fbrNodeFn.Title())
// 	cmd.Fld("Rxr", fbrNodeFn.Rxr.Typ)
// 	for _, p := range fbrNodeFn.InPrms {
// 		fld := cmd.Fld(p.Title(), p.Typ)
// 		if p.IsVariadic() {
// 			fld.Mod = mod.Slice
// 		}
// 	}
// 	if len(fbrNodeFn.OutPrms) != 0 && fbrNodeFn.OutTyp().Bse().arr != nil {
// 		cmd.Fld("OutIdx", _sys.Bsc.Unt)
// 		cmd.Fld("Outs", fbrNodeFn.OutTyp().Bse().arr)
// 	}
// 	fnAct := rxrFbr.Fle.Bse().TypFn(k.Act, cmd)
// 	fnAct.Addf("(*x.Outs)[x.OutIdx] = x.Rxr%v", fbrNodeFn.CallNode())

// 	// preamble
// 	x.NodePreamble(r.Node, r.InPrms, &r.Lines, true)
// 	// body
// 	if len(r.InPrms) == 0 {
// 		r.Addf("r.%v = %v()", outWve.elm.arr.Camel(), outWve.elm.arr.New.Ref(x))
// 	} else {
// 		var b strings.Builder // prm len
// 		for n, p := range r.InPrms {
// 			if n != 0 {
// 				b.WriteRune('*')
// 			}
// 			if !p.IsVariadic() {
// 				b.WriteString("len(*")
// 				b.WriteString(p.Name)
// 				b.WriteString(")")
// 			} else {
// 				r.Addf("%vCnt := 1", p.Name)
// 				r.Addf("if len(%v) != 0 {", p.Name)
// 				r.Addf("%[1]vCnt = len(%[1]v)", p.Name)
// 				r.Add("}")
// 				b.WriteString(fmt.Sprintf("%vCnt", p.Name))
// 			}
// 		}
// 		r.Addf("r.%v = %v(unt.Unt(%v))", outWve.elm.arr.Camel(), outWve.elm.arr.Make.Ref(x), b.String())
// 	}
// 	r.Addf("outIdx, acts := %v, make([]sys.Act, len(*r.%v))", _sys.Bsc.Unt.Zero.Ref(x), outWve.elm.arr.Camel())
// 	for n := 0; n < len(r.InPrms)-1; n++ {
// 		r.Addf("for _, %v := range *%v {", fbrNodeFn.InPrms[n].Name, r.InPrms[n].Name)
// 	}

// 	wrtCmd := func(skpVaridic bool) {
// 		r.Addf("cmd := %v{}", cmd.Adr(x))
// 		r.Add("cmd.Rxr = x")
// 		for _, p := range fbrNodeFn.InPrms {
// 			if p.IsVariadic() {
// 				if !skpVaridic {
// 					r.Addf("cmd.%v = *%v", p.Title(), p.Camel())
// 				}
// 			} else {
// 				r.Addf("cmd.%v = %v", p.Title(), p.Camel())
// 			}
// 		}
// 		r.Add("cmd.OutIdx = outIdx")
// 		r.Addf("cmd.Outs = r.%[1]v", outWve.elm.arr.Camel())
// 		r.Add("acts[outIdx] = cmd")
// 		r.Add("outIdx++")
// 	}
// 	lstPrm := r.InPrms.Lst()
// 	if !lstPrm.IsVariadic() {
// 		r.Addf("for _, %v := range *%v {", fbrNodeFn.InPrms.Lst().Name, lstPrm.Name)
// 		wrtCmd(false)
// 		r.Add("}")
// 	} else {
// 		r.Addf("if len(%v) == 0 {", lstPrm.Name) //no variadic count
// 		wrtCmd(true)
// 		r.Add("} else {")
// 		r.Addf("for _, %v := range %v {", fbrNodeFn.InPrms.Lst().Name, lstPrm.Name)
// 		wrtCmd(false)
// 		r.Add("}")
// 		r.Add("}")
// 	}
// 	for n := 0; n < len(r.InPrms)-1; n++ {
// 		r.Add("}")
// 	}
// 	r.Add("sys.Run().Pll(acts...)")
// 	r.Add("return r")
// 	// end
// 	outWve.bse.Fle.Bse().NodeFns(r, r.Node)

// 	// TODO:
// 	funcIfc := _sys.Tst.NodeFuncIfc(r) // create func type
// 	_sys.Tst.NodeVar(r, funcIfc)
// 	outWve.Test.Gen(r, r.Node.Prnt().Typ.Bse().TestPth)

// 	return r
// }

func (x *FleNodeBse) NodePrms(node *Struct, inPrms InPrms, get bool) {
	for _, p := range inPrms {
		if !p.IsFldSkp() {
			nodePrm := node.FldPrm(p)
			nodePrm.Atr = atr.TstZeroSkp
			if get {
				nodePrm.Atr |= atr.Get
			}
		}
	}
}
func (x *FleNodeBse) NodePreamble(node *Struct, inPrms InPrms, r *Lines, getPrm bool) {
	x.NodePrms(node, inPrms, getPrm)
	r.Addf("r := %v{}", node.Adr(node.Fle))
	r.Add("r.Slf = r")
	if prnt := node.Prnt(); prnt != nil {
		r.Addf("r.%v = x.Slf", prnt.Name)
	}
	for _, p := range inPrms {
		r.Addf("r.%v = %v", p.Title(), p.Camel())
	}
	x.NodeArrInit(node, r)
}
func (x *FleNodeBse) NodeArrInit(node *Struct, r *Lines) {
	for _, fld := range *node.SelFlds(func(f *Fld) bool { return f.IsArr() }) {
		r.Addf("r.%v = %v()", fld.Name, fld.Typ.(*Arr).New.Ref(x))
	}
}
func (x *FleBse) NodeFns(fn Fn, node *Struct) {
	// node: fn: Name
	fn0 := x.TypFn(k.Name, node)
	fn0.OutPrm(_sys.Bsc.Str)
	fn0.Addf("return %v(%q)", _sys.Bsc.Str.Typ().Ref(x), fn.Title())
	if node.IsTst() { // tst
		_sys.Tst.GenTyp(node)
	}
	// node: fn: PrmWrt
	fn1 := x.TypFn(k.PrmWrt, node)
	fn1.InPrm(BuilderPtr, "b")
	for n, p := range fn.In() {
		if p.IsVariadic() {
			fn1.Addf("if len(x.%v) != 0 {", p.Title())
			if n != 0 {
				fn1.Add("b.WriteRune(' ')")
			}
			fn1.Addf("for n, v := range x.%v {", p.Title())
			fn1.Add("if n != 0 {")
			fn1.Add("b.WriteRune(' ')")
			fn1.Add("}")
			fn1.Addf("v.StrWrt(b)")
			fn1.Add("}")
			fn1.Add("}")
		} else {
			if n != 0 {
				fn1.Add("b.WriteRune(' ')")
			}
			fn1.Addf("x.%v.StrWrt(b)", p.Title())
		}
	}
	// node: fn: Prm
	fn2 := x.TypFn(k.Prm, node)
	fn2.OutPrm(String)
	fn2.Addf("b := %v{}", BuilderPtr.Adr(x))
	fn2.Add("x.PrmWrt(b)")
	fn2.Addf("return b.String()")
	// // node: fn: Ttl
	// fn3 := x.TypFn(k.Ttl, node)
	// fn3.OutPrm(_sys.Bsc.Str)
	// fn3.Addf("b := %v{}", BuilderPtr.Adr(x))
	// fn3.Addf("b.WriteString(\"%v(\")", fn.Upper())
	// fn3.Add("x.PrmWrt(b)")
	// fn3.Add("b.WriteRune(')')")
	// fn3.Addf("return %v(b.String())", _sys.Bsc.Str.Typ().Ref(x))
	// node: fn: StrWrt
	fn4 := x.TypFn(k.StrWrt, node)
	fn4.InPrm(BuilderPtr, "b")
	if prnt := node.Prnt(); prnt != nil {
		fn4.Addf("x.%v.StrWrt(b)", prnt.Name)
		fn4.Addf("b.WriteString(\".%v(\")", fn.Camel())
	} else {
		fn4.Addf("b.WriteString(\"%v.%v(\")", node.Pkg.Name, fn.Camel())
	}
	fn4.Add("x.PrmWrt(b)")
	fn4.Add("b.WriteRune(')')")
	// node: fn: String
	fn5 := x.TypFn(k.String, node)
	fn5.OutPrm(String)
	fn5.Addf("b := %v{}", BuilderPtr.Adr(x))
	fn5.Add("x.StrWrt(b)")
	fn5.Add("return b.String()")
}
