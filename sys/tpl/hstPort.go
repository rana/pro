package tpl

// import (
// 	"sys/k"
// 	"sys/tpl/atr"
// )

// type (
// 	FleHstPort struct {
// 		FleHstBse
// 	}
// 	FleHstPorts struct {
// 		FleBse
// 		PrtArr
// 		PrtArrStrWrt // for PrmWrt
// 		PrtString
// 	}
// )

// func (x *DirHst) NewPort() (r *FleHstPort) {
// 	r = &FleHstPort{}
// 	// x.Port = r
// 	r.Name = k.Port
// 	r.Pkg = x.Pkg
// 	r.Ifc(r.Name, atr.TypAnaIfc)
// 	r.AddFle(r)
// 	return r
// }
// func (x *DirHst) NewPorts() (r *FleHstPorts) {
// 	r = &FleHstPorts{}
// 	// x.Ports = r
// 	// r.FleBse = *NewArr(x.Port, &r.PrtArr, x.Port.Pkg)
// 	r.AddFle(r)
// 	return r
// }

// func (x *FleHstPort) InitTyp(bse *TypBse) {
// 	x.FleHstBse.InitTyp(bse)
// 	x.Typ().Bse().TestPth = append(_sys.Ana.Hst.Stgy.Typ().Bse().TestPth, &TestStp{
// 		MdlFst: func(r *PkgFn) { r.Add("port := tst.HstStgyPortStgyPort(stgy)") },
// 	})
// }

// func (x *FleHstPort) InitFld(s *Struct) {
// 	x.FleHstBse.InitFld(s)
// 	x.bse.FldTyp(_sys.Ana.Port).Atr = atr.Get
// 	x.bse.Fld("stgys", _sys.Ana.Hst.Stgy.arr).Atr = atr.BytLitStrEqlBqSelSkp | atr.Get | atr.Arr
// }
// func (x *FleHstPort) InitPkgFn() {
// 	x.New()
// }
// func (x *FleHstPort) New() (r *PkgFn) {
// 	r = x.NodePkgFn("NewPort", "", x, func(r *PkgFn) {
// 		r.Node.Name = "PortMulti"
// 		r.InPrmVariadic(_sys.Ana.Hst.Stgy, "vs").LitVal("stgy")
// 	})
// 	r.Atr = atr.Lng
// 	r.Add("r.BalFstUsd = ana.Cfg.Hst.BalUsd")
// 	r.Add("r.BalLstUsd = ana.Cfg.Hst.BalUsd")
// 	r.Add("r.TrdPct = ana.Cfg.Hst.TrdPct")
// 	r.Add("r.AddStgy(vs...)")
// 	r.Add("return r")
// 	return r
// }
// func (x *FleHstPort) InitTypFn() {
// 	x.FleHstBse.InitTypFn()
// 	x.Stgys()
// 	x.AddStgy()
// 	// x.Prfm()
// 	// x.Splt()
// 	x.Rng()
// }
// func (x *FleHstPort) Stgys() (r *TypFn) {
// 	r = x.TypFna(k.Stgys, atr.Lng, x.bse)
// 	r.OutPrm(_sys.Ana.Hst.Stgy.arr)
// 	r.Add("return x.stgys")
// 	x.MemSigFn(r)
// 	return r
// }
// func (x *FleHstPort) AddStgy() (r *TypFn) {
// 	r = x.TypFna("AddStgy", atr.Lng, x.bse)
// 	r.InPrmVariadic(_sys.Ana.Hst.Stgy, "vs")
// 	r.OutPrm(x)
// 	r.Add("for _, v := range vs {")
// 	r.Add("v.Bse().port = x.Slf")
// 	r.Add("}")
// 	r.Add("*x.stgys = append(*x.stgys, vs...)")
// 	r.Add("return x.Slf")
// 	x.MemSigFn(r)
// 	return r
// }

// // func (x *FleHstPort) Prfm() (r *TypFn) {
// // 	prfm, name := _sys.Ana.Hst.Prfm, strings.Title(k.Prfm)
// // 	x.Import(_sys)
// // 	r = x.ElmNodeTypFn(name, "", "", prfm, func(r *TypFn) {
// // 		r.Node.Name = "PrfmPort"
// // 	})
// // 	r.Add("if ana.Cfg.Trc.IsHstPrfm() {")
// // 	r.Addf("sys.Logf(\"%%p %v(%%v)\", r, r.Prm())", r.Node.Full())
// // 	r.Add("}")
// // 	r.Addf("var dayCnt %v", _sys.Bsc.Unt.Ref(x))
// // 	r.Addf("losLimMax := %v", _sys.Bsc.Flt.Min.Ref(x))
// // 	r.Addf("durLimMax := %v", _sys.Bsc.Tme.Min.Ref(x))
// // 	r.Add("for _, stgy := range *x.stgys {")
// // 	r.Add("losLimMax = losLimMax.Max(stgy.Bse().LosLim)")
// // 	r.Add("durLimMax = durLimMax.Max(stgy.Bse().DurLim)")
// // 	r.Add("dayCnt = dayCnt.Max(stgy.Bse().Instr.Bse().Ana.HstStm.Tmes.WeekdayCnt())")
// // 	r.Add("}")
// // 	r.Add("r.port = x.Slf")
// // 	r.Add("r.Prfm = x.CalcPrfm(x.Trds, x.stgys.Cnt(), dayCnt, losLimMax, durLimMax, x.Slf.String())")
// // 	r.Add("return r")
// // 	return r
// // }

// // func (x *FleHstPort) Splt() (r *TypFn) {
// // 	r = x.TypFna(k.Splt, atr.Lng, x.bse)
// // 	r.InPrm(_sys.Bsc.Flt, "pnt").LitVal("1.0")
// // 	r.OutPrm(_sys.Ana.Hst.Splt, "r")
// // 	r.Addf("r = %v{}", r.OutTyp().Adr(x))
// // 	r.Addf("r.Stgy = x.stgys.Fst()")
// // 	r.Addf("r.Pnt = pnt")
// // 	r.Add("r.Btm, r.Top = r.Stgy.Bse().Trds.SelPnlUsdSplt(pnt)")
// // 	r.Add("return r")
// // 	x.MemSigFn(r)
// // 	x.Test.Gen(r, x.Typ().Bse().TestPth, x.Typ())
// // 	return r
// // }
// func (x *FleHstPort) Rng() (r *TypFn) {
// 	x.Import(_sys)
// 	x.Import("regexp")
// 	r = x.TypFna(k.Rng, atr.Lng, x.bse)
// 	r.InPrm(_sys.Bsc.TmeRng, "rng") //.LitVal("1.0")
// 	r.OutPrm(x, "r")

// 	r.Addf("script := x.Slf.String()")
// 	r.Add("// REPLACE RNG PRM IN ALL INSTR")
// 	r.Add("re := regexp.MustCompile(\"hst.[[:word:]]+[(][)].[[:word:]]+[(]([^)]+)\")")
// 	r.Add("matches := re.FindStringSubmatch(script)")
// 	r.Add("if len(matches) > 1 {")
// 	r.Add("for n := 1; n < len(matches); n++ {")
// 	r.Add("re = regexp.MustCompile(matches[n])")
// 	r.Add("script = re.ReplaceAllString(script, rng.String())")
// 	r.Add("}")
// 	r.Add("vs := sys.Actr().RunIfc(script)")
// 	r.Add("return vs[len(vs)-1].(Port)")
// 	r.Add("}")
// 	r.Add("return r")
// 	x.MemSigFn(r)
// 	// x.Test.Gen(r, x.Typ().Bse().TestPth, x.Typ())
// 	return r
// }

// // func (x *FleHstPort) Rng() (r *TypFn) {
// // 	r = x.TypFna(k.Rng, atr.Lng, x.bse)
// // 	r.InPrm(_sys.Bsc.TmeRng, "rng") //.LitVal("1.0")
// // 	r.OutPrm(x, "r")
// // 	r.Add("stgyBse := x.stgys.At(0).Bse()")
// // 	// r.Add("sys.Log(\"bck trds\", stgyBse.Trds)")
// // 	r.Add("stgyBse.Trds.Clr() // clear prvs rng trds")
// // 	r.Add("stgyBse.Cnd.Sub()")
// // 	r.Add("for _, cls := range stgyBse.Clss {")
// // 	r.Add("cls.Sub()")
// // 	r.Add("}")
// // 	r.Add("stgyBse.Instr.Tx(rng)")
// // 	r.Add("stgyBse.Cnd.Unsub()")
// // 	r.Add("for _, cls := range stgyBse.Clss {")
// // 	r.Add("cls.Unsub()")
// // 	r.Add("}")
// // 	r.Add("r = stgyBse.Port()")
// // 	// r.Add("sys.Log(\"fwd trds\", stgyBse.Trds)")
// // 	r.Add("return r")
// // 	x.MemSigFn(r)
// // 	return r
// // }
