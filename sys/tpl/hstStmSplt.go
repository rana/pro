package tpl

// import (
// 	"sys/k"
// 	"sys/tpl/atr"
// )

// type (
// 	FleHstStmSplt struct {
// 		FleBse
// 		// PrtLog
// 		// PrtIfc
// 		lss  *Struct
// 		gtr  *Struct
// 		tune *Ifc
// 	}
// 	FleHstStmSplts struct {
// 		FleBse
// 		PrtArr
// 		// PrtLog
// 		// PrtIfc
// 	}
// )

// func (x *DirHst) NewStmSplt() (r *FleHstStmSplt) {
// 	r = &FleHstStmSplt{}
// 	x.StmSplt = r
// 	r.Name = k.StmSplt
// 	r.Pkg = x.Pkg
// 	r.StructPtr(r.Name, atr.Typ)
// 	r.AddFle(r)
// 	return r
// }
// func (x *DirHst) NewStmSplts() (r *FleHstStmSplts) {
// 	r = &FleHstStmSplts{}
// 	x.StmSplts = r
// 	r.FleBse = *NewArr(x.StmSplt, &r.PrtArr, x.StmSplt.Pkg)
// 	r.AddFle(r)
// 	return r
// }
// func (x *FleHstStmSplt) InitTyp(bse *TypBse) {
// 	x.lss = x.StructPtrf("%vLssTune", atr.None, x.Typ().Title())
// 	x.gtr = x.StructPtrf("%vGtrTune", atr.None, x.Typ().Title())
// 	x.tune = x.Ifcf("%vTune", atr.None, x.Typ().Title())
// }
// func (x *FleHstStmSplt) InitFld(s *Struct) {
// 	s.Fld("Stm", _sys.Ana.Hst.Stm).Atr = atr.Get
// 	s.Fld("Btm", _sys.Bsc.Flt.arr).Atr = atr.Get
// 	s.Fld("Top", _sys.Bsc.Flt.arr).Atr = atr.Get
// 	x.lss.Fld("stm", _sys.Ana.Hst.Stm)
// 	x.lss.Fld("min", _sys.Bsc.Flt)
// 	x.lss.Fld("cnt", _sys.Bsc.Unt)
// 	x.gtr.Fld("stm", _sys.Ana.Hst.Stm)
// 	x.gtr.Fld("max", _sys.Bsc.Flt)
// 	x.gtr.Fld("cnt", _sys.Bsc.Unt)
// }
// func (x *FleHstStmSplt) InitIfc(i *Ifc) {
// 	var sig *MemSig
// 	sig = x.MemSiga(k.Cnt, atr.None, x.tune)
// 	sig.OutPrm(_sys.Bsc.Unt)
// 	sig = x.MemSiga(k.Cnd, atr.None, x.tune)
// 	sig.OutPrm(_sys.Ana.Hst.Cnd)
// }
// func (x *FleHstStmSplt) InitTypFn() {
// 	x.TuneStm()
// 	x.Cnt(x.lss)
// 	x.Cnt(x.gtr)
// 	x.LssCnd()
// 	x.GtrCnd()
// 	// x.TopSrt()
// 	x.TuneSacf()
// }
// func (x *FleHstStmSplt) TuneStm() (r *TypFn) {
// 	r = x.TypFn(k.TuneStm)
// 	r.InPrm(_sys.Bsc.Unt, "cntLim")
// 	r.OutPrm(x.tune)
// 	r.Add("var btmLssCnt, btmGtrCnt unt.Unt")
// 	r.Add("minLim, maxLim := x.Top.MinMax()")
// 	r.Add("for _, v := range *x.Btm {")
// 	r.Add("if v < minLim {")
// 	r.Add("btmLssCnt++")
// 	r.Add("} else if v > maxLim {")
// 	r.Add("btmGtrCnt++")
// 	r.Add("}")
// 	r.Add("}")
// 	r.Add("if btmLssCnt <= cntLim && btmGtrCnt <= cntLim {")
// 	r.Add("return nil")
// 	r.Add("}")
// 	r.Add("if btmGtrCnt > btmLssCnt {")
// 	r.Addf("r := %v{}", x.gtr.Adr(x))
// 	r.Add("r.stm = x.Stm")
// 	r.Add("r.max = maxLim")
// 	r.Add("r.cnt = btmGtrCnt")
// 	r.Add("return r")
// 	r.Add("}")
// 	r.Addf("r := %v{}", x.lss.Adr(x))
// 	r.Add("r.stm = x.Stm")
// 	r.Add("r.min = minLim")
// 	r.Add("r.cnt = btmLssCnt")
// 	r.Add("return r")
// 	return r
// }

// //
// func (x *FleHstStmSplt) TuneSacf() (r *TypFn) {
// 	r = x.TypFn(k.TuneSacf)
// 	r.InPrm(_sys.Bsc.Unt, "trimItrLim").LitVal("3")
// 	r.InPrm(_sys.Bsc.Unt, "trimMin").LitVal("4")
// 	r.InPrm(_sys.Bsc.Unt, "trimForgiveLim").LitVal("1")
// 	r.OutPrm(x.tune)
// 	r.Add("// skp positives looking for lrg qty of negs")
// 	r.Add("// trimForgiveLim: forgive pos skip that is less than trimMin")
// 	r.Add("// expect x.Top to be SrtAsc")
// 	r.Add("if x.Top.Cnt() < 2 {")
// 	r.Add("return nil")
// 	r.Add("}")
// 	r.Add("var btmLwrCnt, btmUprCnt, topLwrIdx, topUprIdx, trimForgiveCnt unt.Unt")
// 	r.Add("if x.Top.Cnt() <= trimItrLim {")
// 	r.Add("trimItrLim = x.Top.Cnt() - 1")
// 	r.Add("}")
// 	r.Add("for n := unt.One; n <= trimItrLim; n++ { // top lwr")
// 	r.Add("btmLwrCnt = 0")
// 	r.Add("for _, vBtm := range *x.Btm {")
// 	r.Add("if vBtm < (*x.Top)[n] {")
// 	r.Add("btmLwrCnt++")
// 	r.Add("}")
// 	r.Add("}")
// 	r.Add("if btmLwrCnt >= trimMin {")
// 	r.Add("topLwrIdx = n")
// 	r.Add("break")
// 	r.Add("}")
// 	r.Add("if trimForgiveLim != 0 && trimForgiveCnt == trimForgiveLim {")
// 	r.Add("break")
// 	r.Add("}")
// 	r.Add("trimForgiveCnt++")
// 	r.Add("}")
// 	r.Add("trimForgiveCnt = 0")
// 	r.Add("var itr unt.Unt")
// 	r.Add("for n := int(x.Top.LstIdx() - 1); n > -1; n-- { // top upr")
// 	r.Add("if itr == trimItrLim {")
// 	r.Add("break")
// 	r.Add("}")
// 	r.Add("btmUprCnt = 0")
// 	r.Add("for _, vBtm := range *x.Btm {")
// 	r.Add("if vBtm > (*x.Top)[n] {")
// 	r.Add("btmUprCnt++")
// 	r.Add("}")
// 	r.Add("}")
// 	r.Add("if btmUprCnt >= trimMin {")
// 	r.Add("topUprIdx = unt.Unt(n)")
// 	r.Add("break")
// 	r.Add("}")
// 	r.Add("if trimForgiveLim != 0 && trimForgiveCnt == trimForgiveLim {")
// 	r.Add("break")
// 	r.Add("}")
// 	r.Add("trimForgiveCnt++")
// 	r.Add("itr++")
// 	r.Add("}")
// 	r.Add("if topLwrIdx == 0 && topUprIdx == 0 {")
// 	r.Add("return nil")
// 	r.Add("}")
// 	r.Add("if topUprIdx > topLwrIdx {")
// 	r.Addf("r := %v{}", x.gtr.Adr(x))
// 	r.Add("r.stm = x.Stm")
// 	r.Add("r.max = (*x.Top)[topUprIdx]")
// 	r.Add("r.cnt = btmUprCnt")
// 	r.Add("return r")
// 	r.Add("}")
// 	r.Addf("r := %v{}", x.lss.Adr(x))
// 	r.Add("r.stm = x.Stm")
// 	r.Add("r.min = (*x.Top)[topLwrIdx]")
// 	r.Add("r.cnt = btmLwrCnt")
// 	r.Add("return r")
// 	return r
// }
// func (x *FleHstStmSplt) Cnt(side *Struct) (r *TypFn) {
// 	r = x.TypFn(k.Cnt, side)
// 	r.OutPrm(_sys.Bsc.Unt)
// 	r.Add("return x.cnt")
// 	return r
// }
// func (x *FleHstStmSplt) LssCnd() (r *TypFn) {
// 	r = x.TypFn(k.Cnd, x.lss)
// 	r.OutPrm(_sys.Ana.Hst.Cnd)
// 	r.Add("// exclude below limit")
// 	r.Add("return x.stm.SclGeq(x.min)")
// 	return r
// }
// func (x *FleHstStmSplt) GtrCnd() (r *TypFn) {
// 	r = x.TypFn(k.Cnd, x.gtr)
// 	r.OutPrm(_sys.Ana.Hst.Cnd)
// 	r.Add("// exclude above limit")
// 	r.Add("return x.stm.SclLeq(x.max)")
// 	return r
// }

// // func (x *FleHstStmSplt) TopSrt() (r *TypFn) {
// // 	r = x.TypFna("TopSrt", atr.PrtActFn)
// // 	r.OutPrm(x)
// // 	r.Add("x.Top.SrtDsc()")
// // 	r.Add("return x")
// // 	return r
// // }
