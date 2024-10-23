package tpl

// import (
// 	"sys"
// 	"sys/k"
// 	"sys/tpl/atr"
// )

// type (
// 	FleHstPrcp struct {
// 		FleBse
// 	}
// )

// func (x *DirHst) NewPrcp() (r *FleHstPrcp) {
// 	r = &FleHstPrcp{}
// 	x.Prcp = r
// 	r.Name = k.Prcp
// 	r.Pkg = x.Pkg
// 	r.StructPtr(r.Name, atr.Typ)
// 	r.AddFle(r)
// 	return r
// }
// func (x *FleHstPrcp) InitTyp(bse *TypBse) {
// 	x.Typ().Bse().TestPth = append(_sys.Ana.Hst.Port.Typ().Bse().TestPth, &TestStp{
// 		MdlFst: func(r *PkgFn) {
// 			r.Add("prcp := hst.NewPrcp()")
// 			r.Add("prcp.Stm(stm.AggSma(2))")
// 			r.Add("prcp.Stm(stm.AggSma(4))")
// 			r.Add("prcp.Stm(stm.AggSma(8))")
// 		},
// 	})
// }
// func (x *FleHstPrcp) InitVals(bse *TypBse) {
// 	bse.Lits = sys.Vs("hst.newPrcp()")
// 	bse.Vals = sys.Vs("hst.NewPrcp()")
// }
// func (x *FleHstPrcp) InitFld(s *Struct) {
// 	s.Fld("Stms", _sys.Ana.Hst.Stms)
// 	s.FldSlice("andCnds", _sys.Ana.Hst.Cnd)
// }
// func (x *FleHstPrcp) InitPkgFn() {
// 	x.new()
// }
// func (x *FleHstPrcp) new() (r *PkgFn) {
// 	r = x.PkgFna("NewPrcp", atr.Lng)
// 	r.InPrmVariadic(_sys.Ana.Hst.Stm, "vs")
// 	r.OutPrm(x)
// 	r.Addf("r := %v{}", r.OutTyp().Adr(x))
// 	r.Add("r.Stm(vs...)")
// 	r.Add("return r")
// 	return r
// }
// func (x *FleHstPrcp) InitTypFn() {
// 	x.Stm()
// 	x.Splt()
// 	x.StrWrt()
// }
// func (x *FleHstPrcp) Stm() (r *TypFn) {
// 	r = x.TypFn(k.Stm)
// 	r.InPrmVariadic(_sys.Ana.Hst.Stm, "stms")
// 	r.OutPrm(x)
// 	r.Add("x.Stms = NewStms()")
// 	r.Add("x.Stms.Push(stms...)")
// 	r.Add("return x")
// 	return r
// }
// func (x *FleHstPrcp) Splt() (r *TypFn) {
// 	x.Import(_sys)
// 	cmd := x.StructPtrf("%vSpltCmd", atr.None, x.Typ().Title())
// 	cmd.Fld("Stm", _sys.Ana.Hst.Stm)
// 	cmd.Fld("Btm", _sys.Bsc.Tme.arr)
// 	cmd.Fld("Top", _sys.Bsc.Tme.arr)
// 	cmd.Fld("OutIdx", _sys.Bsc.Unt)
// 	cmd.Fld("Outs", _sys.Ana.Hst.StmSplts)
// 	cmdAct := x.TypFn(k.Act, cmd)
// 	cmdAct.Add("(*x.Outs)[x.OutIdx] = x.Stm.Splt(x.Btm, x.Top)")

// 	r = x.TypFn(k.Splt)
// 	r.InPrm(_sys.Ana.Hst.Splt, "splt").LitVal("port.Splt(0.0)")
// 	r.OutPrm(_sys.Ana.Hst.PrcpSplt, "r")
// 	r.Addf("r = %v{}", r.OutTyp().Adr(x))
// 	r.Add("r.Prcp = x")
// 	r.Add("r.Splt = splt")
// 	r.Add("r.StmSplts = MakeStmSplts(x.Stms.Cnt())")
// 	r.Add("if x.Stms.Cnt() == 0 {")
// 	r.Add("return r")
// 	r.Add("}")
// 	r.Add("btm := splt.Btm.OpnTmes()")
// 	r.Add("top := splt.Top.OpnTmes()")
// 	r.Add("acts := make([]sys.Act, len(*x.Stms))")
// 	r.Add("var outIdx unt.Unt")
// 	r.Add("for _, stm := range *x.Stms {")
// 	r.Addf("cmd := %v{}", cmd.Adr(x))
// 	r.Add("cmd.Stm = stm")
// 	r.Add("cmd.Btm = btm")
// 	r.Add("cmd.Top = top")
// 	r.Add("cmd.OutIdx = outIdx")
// 	r.Add("cmd.Outs = r.StmSplts")
// 	r.Add("acts[outIdx] = cmd")
// 	r.Add("outIdx++")
// 	r.Add("}")
// 	r.Add("sys.Run().Pll(acts...)")
// 	r.Add("return r")
// 	x.Test.Gen(r, x.Typ().Bse().TestPth, x.Typ())
// 	return r
// }
// func (x *FleHstPrcp) StrWrt() (r *TypFn) {
// 	r = x.TypFn(k.StrWrt)
// 	r.InPrm(BuilderPtr, "b")
// 	r.Add("b.WriteString(\"hst.newPrcp(\")")
// 	r.Add("for n, stm := range *x.Stms {")
// 	r.Add("if n != 0 {")
// 	r.Add("b.WriteRune(' ')")
// 	r.Add("}")
// 	r.Add("stm.StrWrt(b)")
// 	r.Add("}")
// 	r.Add("b.WriteRune(')')")
// 	return r
// }
