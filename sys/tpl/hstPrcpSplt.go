package tpl

// import (
// 	"sys/k"
// 	"sys/tpl/atr"
// )

// type (
// 	FleHstPrcpSplt struct {
// 		FleBse
// 		// PrtLog
// 		// PrtIfc
// 	}
// )

// func (x *DirHst) NewPrcpSplt() (r *FleHstPrcpSplt) {
// 	r = &FleHstPrcpSplt{}
// 	x.PrcpSplt = r
// 	r.Name = k.PrcpSplt
// 	r.Pkg = x.Pkg
// 	r.StructPtr(r.Name, atr.Typ)
// 	r.AddFle(r)
// 	return r
// }
// func (x *FleHstPrcpSplt) InitTyp(bse *TypBse) {
// 	x.Typ().Bse().TestPth = append(_sys.Ana.Hst.Prcp.Typ().Bse().TestPth, &TestStp{
// 		MdlFst: func(r *PkgFn) {
// 			r.Add("prcpSplt := prcp.Splt(port.Splt(1.0))")
// 		},
// 	})
// }
// func (x *FleHstPrcpSplt) InitFld(s *Struct) {
// 	s.Fld("Prcp", _sys.Ana.Hst.Prcp).Atr = atr.Get
// 	s.Fld("Splt", _sys.Ana.Hst.Splt).Atr = atr.Get
// 	s.Fld("StmSplts", _sys.Ana.Hst.StmSplts)
// }
// func (x *FleHstPrcpSplt) InitTypFn() {
// 	x.TuneStm()
// 	x.TuneSacf()
// 	x.TuneTil()
// 	x.TuneSacfTil()
// }
// func (x *FleHstPrcpSplt) TuneStm() (r *TypFn) {
// 	r = x.TypFn(k.TuneStm)
// 	r.InPrm(_sys.Bsc.Unt, "cntLim").LitVal("1")
// 	r.OutPrm(_sys.Ana.Hst.Stgy, "r")
// 	r.Addf("var maxCnt %v // filter for max side", _sys.Bsc.Unt.Ref(x))
// 	r.Addf("var tune %v", _sys.Ana.Hst.StmSplt.tune.Ref(x))
// 	// stms
// 	r.Add("for _, stmSplt := range *x.StmSplts {")
// 	r.Add("curTune := stmSplt.TuneStm(cntLim)")
// 	r.Add("if curTune != nil && curTune.Cnt() > maxCnt {")
// 	r.Add("tune = curTune")
// 	r.Add("maxCnt = curTune.Cnt()")
// 	r.Add("}")
// 	r.Add("}")

// 	r.Add("if tune == nil { // no more tuning available")
// 	r.Add("return x.Splt.Stgy")
// 	r.Add("}")
// 	r.Add("stgyBse := x.Splt.Stgy.Bse()")
// 	r.Add("tuneCnd := tune.Cnd()")
// 	r.Add("x.Prcp.andCnds = append(x.Prcp.andCnds, tuneCnd)")
// 	r.Add("tuneOpnCnd := stgyBse.Cnd.And(tuneCnd)")
// 	// r.Add("_, isLong := x.Splt.Stgy.(*StgyLong)")
// 	r.Add("if x.Splt.Stgy.Bse().IsLong {")
// 	r.Add("r = tuneOpnCnd.Long(stgyBse.PrfLim, stgyBse.LosLim, stgyBse.DurLim, stgyBse.Instr, stgyBse.Clss...)")
// 	r.Add("} else {")
// 	r.Add("r = tuneOpnCnd.Shrt(stgyBse.PrfLim, stgyBse.LosLim, stgyBse.DurLim, stgyBse.Instr, stgyBse.Clss...)")
// 	r.Add("}")
// 	x.Import(_sys.Ana)
// 	r.Add("if ana.Cfg.Trc.IsTune() {")
// 	r.Add("prfm0 := x.Splt.Stgy.Port().Prfm().Bse().Prfm")
// 	r.Add("prfm1 := r.Port().Prfm().Bse().Prfm")
// 	r.Addf("sys.Log(%q, prfm0.Dlt(prfm1))", r.Name)
// 	r.Add("}")
// 	r.Add("return r")
// 	x.Test.Gen(r, x.Typ().Bse().TestPth, x.Typ())
// 	return r
// }
// func (x *FleHstPrcpSplt) TuneSacf() (r *TypFn) {
// 	x.Import(_sys)
// 	r = x.TypFn(k.TuneSacf)
// 	r.InPrm(_sys.Bsc.Unt, "trimItrLim").LitVal("3")
// 	r.InPrm(_sys.Bsc.Unt, "trimMin").LitVal("4")
// 	r.InPrm(_sys.Bsc.Unt, "trimForgiveLim").LitVal("1")
// 	r.OutPrm(_sys.Ana.Hst.Stgy, "r")

// 	srtCmd := x.StructPtrf("StmSpltSrtCmd", atr.None)
// 	srtCmd.Fld("StmSplt", _sys.Ana.Hst.StmSplt)
// 	srtCmdAct := x.TypFn(k.Act, srtCmd)
// 	srtCmdAct.Add("x.StmSplt.Top.SrtAsc()")

// 	r.Add("var acts []sys.Act")
// 	r.Add("for _, stmSplt := range *x.StmSplts {")
// 	r.Add("acts = append(acts, &StmSpltSrtCmd{StmSplt: stmSplt})")
// 	r.Add("}")
// 	r.Add("sys.Run().Pll(acts...) // srt stm splts")

// 	cmd := x.StructPtrf("StmSpltTuneSacfCmd", atr.None)
// 	cmd.Fld("StmSplt", _sys.Ana.Hst.StmSplt)
// 	cmd.Fld("Tune", _sys.Ana.Hst.StmSplt.tune)
// 	cmd.Fld("TrimItrLim", _sys.Bsc.Unt)
// 	cmd.Fld("TrimMin", _sys.Bsc.Unt)
// 	cmd.Fld("TrimForgiveCnt", _sys.Bsc.Unt)
// 	cmdAct := x.TypFn(k.Act, cmd)
// 	cmdAct.Add("x.Tune = x.StmSplt.TuneSacf(x.TrimItrLim, x.TrimMin, x.TrimForgiveCnt)")

// 	r.Add("acts = acts[:0]")
// 	r.Add("for _, stmSplt := range *x.StmSplts {")
// 	r.Add("acts = append(acts, &StmSpltTuneSacfCmd{")
// 	r.Add("StmSplt: stmSplt,")
// 	r.Add("TrimItrLim: trimItrLim,")
// 	r.Add("TrimMin: trimMin,")
// 	r.Add("TrimForgiveCnt: trimForgiveLim})")
// 	r.Add("}")
// 	r.Add("sys.Run().Pll(acts...) // tune individual stm splts")

// 	r.Addf("var maxCnt %v // find tune with max cnt", _sys.Bsc.Unt.Ref(x))
// 	r.Addf("var tune %v", _sys.Ana.Hst.StmSplt.tune.Ref(x))

// 	r.Addf("for n := 0; n < len(acts); n++ {")
// 	r.Addf("c := acts[n].(*StmSpltTuneSacfCmd)")
// 	r.Addf("if c.Tune != nil && c.Tune.Cnt() > maxCnt {")
// 	r.Addf("tune = c.Tune")
// 	r.Addf("maxCnt = c.Tune.Cnt()")
// 	r.Addf("}")
// 	r.Addf("}")
// 	r.Addf("")

// 	r.Add("if tune == nil { // no more tuning available")
// 	r.Add("return x.Splt.Stgy")
// 	r.Add("}")
// 	r.Add("stgyBse := x.Splt.Stgy.Bse()")
// 	r.Add("tuneCnd := tune.Cnd()")
// 	r.Add("x.Prcp.andCnds = append(x.Prcp.andCnds, tuneCnd)")
// 	r.Add("tuneOpnCnd := stgyBse.Cnd.And(tuneCnd)")
// 	// r.Add("_, isLong := x.Splt.Stgy.(*StgyLong)")
// 	r.Add("if x.Splt.Stgy.Bse().IsLong {")
// 	r.Add("r = tuneOpnCnd.Long(stgyBse.PrfLim, stgyBse.LosLim, stgyBse.DurLim, stgyBse.Instr, stgyBse.Clss...)")
// 	r.Add("} else {")
// 	r.Add("r = tuneOpnCnd.Shrt(stgyBse.PrfLim, stgyBse.LosLim, stgyBse.DurLim, stgyBse.Instr, stgyBse.Clss...)")
// 	r.Add("}")
// 	x.Import(_sys.Ana)
// 	r.Add("if ana.Cfg.Trc.IsTune() {")
// 	r.Add("prfm0 := x.Splt.Stgy.Port().Prfm().Bse().Prfm")
// 	r.Add("prfm1 := r.Port().Prfm().Bse().Prfm")
// 	r.Addf("sys.Log(%q, prfm0.Dlt(prfm1))", r.Name)
// 	r.Add("}")
// 	r.Add("return r")
// 	x.Test.Gen(r, x.Typ().Bse().TestPth, x.Typ())
// 	return r
// }

// func (x *FleHstPrcpSplt) TuneTil() (r *TypFn) {
// 	r = x.TypFn("TuneTil")
// 	r.InPrm(_sys.Bsc.Flt, "spltPnt").LitVal("1.0")
// 	r.InPrm(_sys.Bsc.Unt, "stmCntLim").LitVal("1")
// 	r.OutPrm(_sys.Ana.Hst.Prfm)
// 	r.Add("prcpSplt := x")
// 	r.Add("prfm := prcpSplt.Splt.Stgy.Port().Prfm()")
// 	r.Add("for {")
// 	r.Add("newStgy := prcpSplt.TuneStm(stmCntLim)")
// 	r.Add("if newStgy == prcpSplt.Splt.Stgy { // tune return original stgy when done")
// 	r.Add("break")
// 	r.Add("}")
// 	r.Add("newPort := newStgy.Port()")
// 	r.Add("prcpSplt = x.Prcp.Splt(newPort.Splt(spltPnt))")
// 	r.Add("prfm = newPort.Prfm()")
// 	r.Add("}")
// 	r.Add("return prfm")
// 	x.Test.Gen(r, x.Typ().Bse().TestPth, x.Typ())
// 	return r
// }
// func (x *FleHstPrcpSplt) TuneSacfTil() (r *TypFn) {
// 	r = x.TypFn(k.TuneSacfTil)
// 	r.InPrm(_sys.Bsc.Flt, "spltPnt").LitVal("1.0")
// 	r.InPrm(_sys.Bsc.Unt, "stmCntLim").LitVal("1")
// 	r.InPrm(_sys.Bsc.Unt, "trimItrLim").LitVal("3")
// 	r.InPrm(_sys.Bsc.Unt, "trimMin").LitVal("4")
// 	r.InPrm(_sys.Bsc.Unt, "trimForgiveLim").LitVal("1")
// 	r.OutPrm(_sys.Ana.Hst.Prfm)
// 	r.Add("prcpSplt := x")
// 	r.Add("prfm := x.TuneTil(spltPnt, stmCntLim) // non-sacrifice tune fst")
// 	r.Add("for {")
// 	r.Add("newStgy := prcpSplt.TuneSacf(trimItrLim, trimMin, trimForgiveLim)")
// 	r.Add("if newStgy == prcpSplt.Splt.Stgy { // tune return original stgy when done")
// 	r.Add("break")
// 	r.Add("}")
// 	r.Add("newPort := newStgy.Port()")
// 	r.Add("prcpSplt = x.Prcp.Splt(newPort.Splt(spltPnt))")
// 	r.Add("prfm = newPort.Prfm()")
// 	r.Add("}")
// 	r.Add("return prfm")
// 	x.Test.Gen(r, x.Typ().Bse().TestPth, x.Typ())
// 	return r
// }
