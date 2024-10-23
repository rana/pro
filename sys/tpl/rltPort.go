package tpl

// import (
// 	"strings"
// 	"sys/k"
// 	"sys/tpl/atr"
// )

// type (
// 	FleRltPort struct {
// 		FleRltBse
// 	}
// 	FleRltPorts struct {
// 		FleBse
// 		PrtArr
// 	}
// )

// func (x *DirRlt) NewPort() (r *FleRltPort) {
// 	r = &FleRltPort{}
// 	x.Port = r
// 	r.Name = k.Port
// 	r.Pkg = x.Pkg
// 	r.Ifc(r.Name, atr.TypAnaIfc)
// 	r.AddFle(r)
// 	return r
// }
// func (x *DirRlt) NewPorts() (r *FleRltPorts) {
// 	r = &FleRltPorts{}
// 	x.Ports = r
// 	r.FleBse = *NewArr(x.Port, &r.PrtArr, x.Port.Pkg)
// 	r.AddFle(r)
// 	return r
// }
// func (x *FleRltPort) InitCnst() {
// 	x.Cnst("PortTrdCls", "0", Uint32)
// 	x.Cnst("PortTrdOpn", "1", Uint32)
// }

// func (x *FleRltPort) InitTyp(bse *TypBse) {
// 	x.FleRltBse.InitTyp(bse)
// 	x.Typ().Bse().TestPth = append(_sys.Ana.Rlt.Stgy.Typ().Bse().TestPth, &TestStp{
// 		MdlFst: func(r *PkgFn) { r.Add("port := tst.RltStgyPortStgyPort(stgy)") },
// 	})
// }
// func (x *FleRltPort) InitFld(s *Struct) {
// 	x.bse.FldTyp(_sys.Ana.Port)
// 	x.FleRltBse.InitFld(s)
// 	x.FldRxs(_sys.Ana.Trd.Rxs)
// 	x.bse.Fld("stgys", _sys.Ana.Rlt.Stgy.arr).Atr = atr.BytLitStrEqlBqSelSkp | atr.Get | atr.Arr
// 	x.bse.Fld("TrdState", Uint32).Atr = atr.TstZeroSkp
// }

// // func (x *FleRltPort) InitIfc(i *Ifc) {
// // 	x.FleRltBse.InitIfc(i)
// // 	var sig *MemSig
// // 	sig = x.MemSig(k.Prfm)
// // 	sig.OutPrm(_sys.Ana.Prfm, "r")
// // }

// // TODO: WITH HST PARITY
// // func (x *FleRltPort) InitPkgFn() {
// // 	x.new()
// // }
// // func (x *FleRltPort) new() (r *PkgFn) {
// // 	r = x.PkgFn("NewPort")
// // 	r.InPrmVariadic(_sys.Ana.Rlt.Stgy, "vs")
// // 	r.OutPrm(x, "r")
// // 	r.Addf("r = %v{}", x.Adr(x))
// // 	r.Add("r.Stgys = NewStgys()")
// // 	r.Add("r.Stgy(vs...)")
// // 	r.Addf("r.Trds = %v()", _sys.Ana.Trd.Arr.New.Ref(x))
// // 	r.Addf("r.Rxs = %v", _sys.Ana.Trd.Rxs.Make(x))
// // 	r.Add("if ana.Cfg.Test {")
// // 	r.Add("r.BalFstUsd = ana.Cfg.Hst.BalUsd")
// // 	r.Add("r.BalLstUsd = ana.Cfg.Hst.BalUsd")
// // 	r.Add("r.TrdPct = ana.Cfg.Hst.TrdPct")
// // 	r.Add("} else {")
// // 	r.Add("r.TrdPct = ana.Cfg.Rlt.TrdPct")
// // 	r.Add("}")
// // 	r.Add("return r")
// // 	return r
// // }
// func (x *FleRltPort) InitTypFn() {
// 	x.FleRltBse.InitTypFn()
// 	x.Stgys()
// 	x.AddStgy()
// 	x.Prfm()
// 	x.Sub(_sys.Ana.Trd.Rx, true)
// 	x.Unsub()
// 	x.MayTrd()
// 	x.OpnTrd()
// 	x.ClsTrd()
// 	x.CalcCls()
// }

// func (x *FleRltPort) Stgys() (r *TypFn) {
// 	r = x.TypFna(k.Stgys, atr.Lng, x.bse)
// 	r.OutPrm(_sys.Ana.Rlt.Stgy.arr)
// 	r.Add("return x.stgys")
// 	x.MemSigFn(r)
// 	return r
// }
// func (x *FleRltPort) AddStgy() (r *TypFn) {
// 	r = x.TypFna("AddStgy", atr.Lng, x.bse)
// 	r.InPrmVariadic(_sys.Ana.Rlt.Stgy, "vs")
// 	r.OutPrm(x)
// 	r.Add("for _, v := range vs {")
// 	r.Add("v.Bse().port = x.Slf")
// 	r.Add("}")
// 	r.Add("*x.stgys = append(*x.stgys, vs...)")
// 	r.Add("return x.Slf")
// 	x.MemSigFn(r)
// 	return r
// }
// func (x *FleRltPort) Prfm() (r *TypFn) {
// 	prfm, name := _sys.Ana.Rlt.Prfm, strings.Title(k.Prfm)
// 	r = x.ElmNodeTypFn(name, "", "", prfm, func(r *TypFn) {
// 		r.Node.Name = "PrfmPort"
// 		r.Node.FldPrnt(x)
// 	})
// 	r.Addf("var dayCnt %v", _sys.Bsc.Unt.Ref(x))
// 	r.Addf("losLimMax := %v", _sys.Bsc.Flt.Min.Ref(x))
// 	r.Addf("durLimMax := %v", _sys.Bsc.Tme.Min.Ref(x))
// 	r.Add("for _, stgy := range *x.stgys {")
// 	r.Add("losLimMax = losLimMax.Max(stgy.Bse().LosLim)")
// 	r.Add("durLimMax = durLimMax.Max(stgy.Bse().DurLim)")
// 	r.Add("dayCnt = dayCnt.Max(stgy.I().HstStm.Tmes.WeekdayCnt()) // for testing; same used by NewStgyMnr")
// 	r.Add("}")
// 	r.Add("r.Prfm = x.CalcPrfm(x.Trds, x.stgys.Cnt(), dayCnt, losLimMax, durLimMax, x.Slf.String())")
// 	r.Add("return r")
// 	if x.Test != nil {
// 		x.Test.Import(_sys.Ana.Hst)
// 		x.Test.Import(_sys.Lng.Pro.Act)
// 		r.T2.MdlFst.Add("mnr := tst.NewStgyMnr(ap)")
// 		r.T2.MdlFst.Add("port.Sub(mnr.Rx, mnr.Id)")
// 		r.T2.MdlFst.Add("var actr act.Actr")
// 		r.T2.MdlFst.Add("vs := actr.RunHst(port.String())")
// 		r.T2.MdlFst.Addf("eHst := vs[len(vs)-1].(hst.Port)")

// 		// r.T2.MdlFst.Addf("eHst := actr.RunHst(port.String())[0].(hst.Port)")
// 		r.T2.MdlFst.Add("mnr.StartFor(instr.Instr(), eHst.Bse().Trds.Cnt())")
// 		r.T2.MdlFst.Add("tst.AnaPrfmEql(t, eHst.Prfm().Bse().Prfm, port.Prfm().Bse().Prfm, \"Prfm\")")
// 		r.T2.MdlFst.Add("port.Unsub(mnr.Id)")
// 		x.Test.Gen(r, x.Typ().Bse().TestPth)
// 	}
// 	return r
// }

// func (x *FleRltPort) Unsub() (r *TypFn) {
// 	x.Import(_sys)
// 	r = x.TypFn(k.Unsub, x.bse)
// 	r.InPrm(Uint32, "id")
// 	r.Add("x.mu.Lock()")
// 	r.Add("delete(x.Rxs, sys.Uint64(id, 0))")
// 	r.Add("x.mu.Unlock()")
// 	x.MemSigFn(r) // add to interface
// 	return r
// }
// func (x *FleRltPort) MayTrd() (r *TypFn) {
// 	r = x.TypFn(k.MayTrd, x.bse)
// 	r.OutPrm(_sys.Bsc.Bol)
// 	r.Add("return atomic.LoadUint32(&x.TrdState) == PortTrdCls")
// 	x.MemSigFn(r)
// 	return r
// }
// func (x *FleRltPort) OpnTrd() (r *TypFn) {
// 	x.Import("sync/atomic")
// 	r = x.TypFn(k.OpnTrd, x.bse)
// 	r.InPrm(_sys.Ana.Trd, "t")
// 	r.InPrm(_sys.Ana.Instr, "i")
// 	r.OutPrm(_sys.Bsc.Bol)
// 	r.OutPrm(_sys.Ana.TrdRsnOpn)
// 	r.Add("if atomic.LoadUint32(&x.TrdState) == PortTrdOpn {")
// 	r.Add("return false, ana.InTrd // already in trd")
// 	r.Add("}")
// 	r.Add("x.mu.Lock()")
// 	r.Add("defer x.mu.Unlock()")
// 	r.Add("if atomic.LoadUint32(&x.TrdState) == PortTrdOpn {")
// 	r.Add("return false, ana.InTrd // already in trd")
// 	r.Add("}")
// 	r.Add("x.Port.CalcOpn(t, i) // set trd flds")
// 	r.Add("if !ana.Cfg.Test {")
// 	r.Add("if !i.Prv.OpnTrd(t, i) {")
// 	r.Add("return false, ana.PrvReject // provider rejected opening; see OpnRes fld")
// 	r.Add("}")
// 	r.Add("}")
// 	r.Add("atomic.StoreUint32(&x.TrdState, PortTrdOpn)")
// 	r.Add("return true, ana.NoTrdRsnOpn")
// 	x.MemSigFn(r)
// 	return r
// }
// func (x *FleRltPort) ClsTrd() (r *TypFn) {
// 	x.Import(_sys)
// 	r = x.TypFn(k.ClsTrd, x.bse)
// 	r.InPrm(_sys.Ana.Trd, "t")
// 	r.InPrm(_sys.Ana.Instr, "i")
// 	r.OutPrm(_sys.Bsc.Bol)
// 	r.Add("if atomic.LoadUint32(&x.TrdState) == PortTrdCls {")
// 	r.Add("return false // not in trd")
// 	r.Add("}")
// 	r.Add("x.mu.Lock()")
// 	r.Add("defer x.mu.Unlock()")
// 	r.Add("if atomic.LoadUint32(&x.TrdState) == PortTrdCls {")
// 	r.Add("return false // not in trd")
// 	r.Add("}")
// 	r.Add("if !ana.Cfg.Test {")
// 	r.Add("if !i.Prv.ClsTrd(t, i) {")
// 	r.Add("sys.Logf(\"*** LIMBO!!! *** TRD CLS FAILED %v\", t)")
// 	r.Add("return false // WHAT TO DO?")
// 	r.Add("}")
// 	r.Add("}")
// 	r.Add("for _, rx := range x.Rxs {")
// 	r.Addf("rx(t)") // used for testing
// 	r.Add("}")
// 	r.Add("atomic.StoreUint32(&x.TrdState, PortTrdCls)")
// 	r.Add("return true")
// 	x.MemSigFn(r)
// 	return r
// }
// func (x *FleRltPort) CalcCls() (r *TypFn) {
// 	r = x.TypFn(k.CalcCls, x.bse)
// 	r.InPrm(_sys.Ana.Trd, "t")
// 	r.InPrm(_sys.Ana.Instr, "i")
// 	r.InPrm(_sys.Bsc.Bol, "compound")
// 	r.Add("x.Port.CalcCls(t, i, compound) // set bse trd flds")
// 	r.Add("if !ana.Cfg.Test {")
// 	r.Add("x.BalLstUsd = i.Prv.AcntRefresh() // GET ACCURATE BALANCE")
// 	r.Add("t.ClsBalUsdAct = x.BalLstUsd")
// 	r.Add("}")
// 	x.MemSigFn(r)
// 	return r
// }
