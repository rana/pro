package tpl

import (
	"sys/k"
	"sys/tpl/atr"
)

type (
	FleRltStgy struct {
		FleRltBse
	}
	FleRltStgys struct {
		FleBse
		PrtArr
	}
)

func (x *DirRlt) NewStgy() (r *FleRltStgy) {
	r = &FleRltStgy{}
	x.Stgy = r
	r.Name = k.Stgy
	r.Pkg = x.Pkg
	r.Ifc(r.Name, atr.TypAnaIfc)
	r.AddFle(r)
	return r
}
func (x *DirRlt) NewStgys() (r *FleRltStgys) {
	r = &FleRltStgys{}
	x.Stgys = r
	r.FleBse = *NewArr(x.Stgy, &r.PrtArr, x.Stgy.Pkg)
	r.AddFle(r)
	return r
}
func (x *FleRltStgy) InitTyp(bse *TypBse) {
	x.FleRltBse.InitTyp(bse)
	x.Typ().Bse().TestPth = append(_sys.Ana.Rlt.Cnd.Typ().Bse().TestPth, &TestStp{
		MdlFst: func(r *PkgFn) { r.Add("stgy := tst.RltCndStgyLong(cnd, 2.0, 4.0, 60*60, instr)") },
	})
}
func (x *FleRltStgy) InitFld(s *Struct) {
	x.FleRltBse.InitFld(s)
	x.bse.FldPrnt(_sys.Ana.Rlt.Cnd)
	x.FldRxs(_sys.Ana.Trd.Rxs)
	x.bse.Fld("IsLong", _sys.Bsc.Bol).Atr = atr.Get | atr.TstZeroSkp
	x.bse.Fld("PrfLim", _sys.Bsc.Flt).Atr = atr.Get
	x.bse.Fld("LosLim", _sys.Bsc.Flt).Atr = atr.Get
	x.bse.Fld("DurLim", _sys.Bsc.Tme).Atr = atr.Get
	x.bse.Fld("MinPnlPct", _sys.Bsc.Flt).Atr = atr.Get | atr.TstZeroSkp
	x.bse.Fld("Instr", _sys.Ana.Rlt.Instr).Atr = atr.Get
	x.bse.Fld("FtrStms", _sys.Ana.Rlt.Stms)
	x.bse.FldSlice("Clss", _sys.Ana.Rlt.Cnd)
	x.bse.Fld("ClsPrfLim", _sys.Bsc.Flt).Atr = atr.TstZeroSkp
	x.bse.Fld("ClsLosLim", _sys.Bsc.Flt).Atr = atr.TstZeroSkp
	x.bse.Fld("ClsTmeLim", _sys.Bsc.Tme).Atr = atr.TstZeroSkp
	x.bse.Fld("LstClsTme", _sys.Bsc.Tme).Atr = atr.TstZeroSkp
	x.bse.Fld("LstClsIdx", _sys.Bsc.Unt).Atr = atr.TstZeroSkp
	x.bse.Fld("OpnIdx", _sys.Bsc.Unt).Atr = atr.TstZeroSkp
	x.bse.Fld("Trd", _sys.Ana.Trd).Atr = atr.TstZeroSkp
	x.bse.Fld("opns", _sys.Bsc.Tme.arr).Atr = atr.TstZeroSkp
	x.bse.Fld("Key", String).Atr = atr.TstZeroSkp
	x.bse.FldSlice("stgyFtrStms", NewExt("*StgyFtrStm"))
	// x.bse.FldSlice("FtrTmes", _sys.Bsc.Tme.arr).Atr = atr.TstZeroSkp
	// x.bse.FldSlice("FtrVals", _sys.Bsc.Flt.arr).Atr = atr.TstZeroSkp
	// x.bse.Fld("Trds", _sys.Ana.Trd.arr).Atr = atr.Get | atr.TstZeroSkp
	// x.bse.Fld("port", _sys.Ana.Rlt.Port)
}

// func (x *FleRltStgy) InitIfc(i *Ifc) {
// 	x.FleRltBse.InitIfc(i)
// 	var sig *MemSig
// 	sig = x.MemSig("RxOpn")
// 	sig.InPrm(_sys.Bsc.Tme, "inPkt")
// 	sig.OutPrmSlice(_sys.Act, "r")
// 	sig = x.MemSig("RxClsLim")
// 	sig.InPrm(_sys.Ana.TmeIdx, "inPkt")
// 	sig.OutPrmSlice(_sys.Act, "r")
// 	sig = x.MemSig("RxClsCnd")
// 	sig.InPrm(_sys.Bsc.Tme, "inPkt")
// 	sig.OutPrmSlice(_sys.Act, "r")
// }
func (x *FleRltStgy) InitTypFn() {
	x.FleRltBse.InitTypFn()
	x.Sub(_sys.Ana.Trd.Rx, true)
	x.Unsub(x.bse, true, true, func(r *TypFn) {
		r.Addf("x.%v.Unsub(x.Id)", x.bse.Prnt().Name)
		r.Add("x.Instr.Unsub(x.Id)")
		r.Add("for _, ftrStm := range *x.FtrStms {")
		r.Add("ftrStm.Unsub(x.Id)")
		r.Add("}")
		r.Add("for _, cndCls := range x.Clss {")
		r.Add("cndCls.Unsub(x.Id)")
		r.Add("}")
	})
	x.I()
	// x.Port()
}
func (x *FleRltStgy) I() (r *TypFn) {
	r = x.TypFna(k.I, atr.None, x.bse)
	r.OutPrm(_sys.Ana.Instr)
	r.Add("return x.Instr.Instr()")
	x.MemSigFn(r) // add to interface
	return r
}

// func (x *FleRltStgy) Port() (r *TypFn) { // PortStgy
// 	port, name := _sys.Ana.Rlt.Port, strings.Title(k.Port)
// 	r = x.ElmNodeTypFn(name, "Stgy", "", port, func(r *TypFn) {
// 		r.Node.Name = "PortStgy" // suffix Node for test
// 		r.Node.FldPrnt(x)
// 		if port.Test != nil {
// 			port.Test.Import(_sys.Ana.Hst)
// 			port.Test.Import(_sys.Lng.Pro.Act)
// 			r.T2.MdlLst.Add("mnr := tst.NewStgyMnr(ap)")
// 			r.T2.MdlLst.Add("a.Sub(mnr.Rx, mnr.Id)")
// 			r.T2.MdlLst.Add("tst.IntegerEql(t, 1, len(a.Rxs), \"Sub Rxs\")")
// 			r.T2.MdlLst.Add("var actr act.Actr")
// 			r.T2.MdlLst.Add("vs := actr.RunHst(a.String())")
// 			r.T2.MdlLst.Addf("eHst := vs[len(vs)-1].(*hst.%v)", r.Node.Title())
// 			r.T2.MdlLst.Add("mnr.StartFor(instr.Instr(), eHst.Trds.Cnt())")
// 			r.T2.MdlLst.Add("if eHst.Trds != nil {")
// 			r.T2.MdlLst.Add("tst.AnaPortEql(t, &eHst.Port, &a.Port, \"Port\")")
// 			r.T2.MdlLst.Add("}")
// 			r.T2.MdlLst.Add("a.Unsub(mnr.Id)")
// 			r.T2.MdlLst.Add("tst.IntegerEql(t, 0, len(a.Rxs), \"Unsub Rxs\")")
// 		}
// 	})
// 	r.Add("r.AddStgy(x.Slf)")
// 	r.Add("if ana.Cfg.Test {")
// 	r.Add("r.BalFstUsd = ana.Cfg.Hst.BalUsd")
// 	r.Add("r.BalLstUsd = ana.Cfg.Hst.BalUsd")
// 	r.Add("r.TrdPct = ana.Cfg.Hst.TrdPct")
// 	r.Add("} else {")
// 	r.Add("r.TrdPct = ana.Cfg.Rlt.TrdPct")
// 	r.Add("}")
// 	r.Add("x.port = r")
// 	r.Add("x.Cnd.Sub(x.Slf.RxOpn, x.Id)")
// 	r.Add("x.Instr.Sub(x.Slf.RxClsLim, x.Id)")
// 	r.Add("for _, cndCls := range x.Clss {")
// 	r.Add("cndCls.Sub(x.Slf.RxClsCnd, x.Id)")
// 	r.Add("}")
// 	r.Add("return r")
// 	return r
// }

// func (x *FleRltStgy) RxOpn(nodeFn *TypFn) (r *TypFn) {
// 	node := nodeFn.Node
// 	r = x.TypFna("RxOpn", atr.None, node)
// 	r.InPrm(_sys.Bsc.Tme, "inPkt")
// 	r.OutPrmSlice(_sys.Act, "r")
// 	r.Add("// RxOpnEnd allows cnd in same graph phase to cls a trd")
// 	r.Add("// and for RxOpnEnd to then open a new trade within the same 1s")
// 	r.Add("if ana.Cfg.Trc.IsRltStgy() {")
// 	r.Addf("sys.Logf(\"%v.%v %%p inPkt %%v\", x, inPkt)", node.Full(), r.Name)
// 	r.Add("}")
// 	r.Add("if x.port.MayTrd() { // IMPORTANT TO AVOID BUG WHERE OPNING NEW TRD WHILE EXISTING ONE ALREADY OPN")
// 	r.Addf("r = append(r, &%vRxOpn{", node.Name)
// 	r.Add("X:     x,")
// 	r.Add("InPkt: inPkt,")
// 	r.Add("tier:  x.Cnd.DstToInstr() + 1,")
// 	r.Add("})")
// 	r.Add("}")
// 	r.Add("return r")
// 	return r
// }
// func (x *FleRltStgy) RxOpnEnd(nodeFn *TypFn) (r *TypFn) {
// 	node, isLong := nodeFn.Node, nodeFn.Lower() == k.Long
// 	r = x.TypFna("RxOpnEnd", atr.None, node)
// 	r.InPrm(_sys.Bsc.Tme, "opnTme") // change to inPkt?
// 	r.OutPrmSlice(_sys.Act, "r")
// 	r.Add("if ana.Cfg.Trc.IsRltStgy() {")
// 	r.Addf("sys.Logf(\"%v.%v %%p inPkt %%v\", x, opnTme)", node.Full(), r.Name)
// 	r.Add("}")
// 	r.Add("if x.port.MayTrd() && x.Trd == nil && opnTme.Geq(x.LstClsTme) {")
// 	r.Add("x.mu.Lock()")
// 	r.Add("defer x.mu.Unlock()")
// 	r.Add("if x.port.MayTrd() && x.Trd == nil {")
// 	r.Add("i := x.Instr.Instr()")
// 	r.Add("opnIdx := i.RltStm.Tmes.SrchIdxEql(opnTme)")
// 	r.Add("if opnIdx < x.LstClsIdx {")
// 	r.Add("opnIdx = x.LstClsIdx")
// 	r.Add("}")
// 	r.Add("mktWeekMin := ana.MktWeekMin(opnTme) + ana.Cfg.MktTrdBuf")
// 	r.Add("mktWeekMax := ana.MktWeekMax(opnTme) - ana.Cfg.MktTrdBuf")
// 	r.Add("if opnTme < mktWeekMin || opnTme > mktWeekMax-x.DurLim {")
// 	r.Add("return nil // avoid opening trades near market close time")
// 	r.Add("}")
// 	r.Add("x.ClsTmeLim = opnTme.Add(x.DurLim).Min(mktWeekMax) // avoid holding postion after market close")
// 	r.Add("// Bid: the price for me to sell at")
// 	r.Add("// Ask: the price for me to buy at")
// 	r.Add("opnBid, opnAsk := i.RltStm.BidAskAt(opnIdx)")
// 	r.Add("opnSpd := i.Spd(opnBid, opnAsk)")
// 	r.Add("if opnSpd > i.SpdOpnLim {")
// 	r.Add("return nil // avoid opening trades with large spreads")
// 	r.Add("}")
// 	r.Add("// PrfLim: 1.1   Pip: 0.0001     1.1 × .0001 = 0.00011")
// 	var long string
// 	if isLong {
// 		long = "IsLong: true,"
// 		r.Add("x.ClsPrfLim = opnAsk + (x.PrfLim * i.Pip)")
// 		r.Add("x.ClsLosLim = opnAsk - (x.LosLim * i.Pip)")
// 	} else {
// 		r.Add("x.ClsPrfLim = opnBid - (x.PrfLim * i.Pip)")
// 		r.Add("x.ClsLosLim = opnBid + (x.LosLim * i.Pip)")
// 	}
// 	r.Addf("t := %v{%v OpnTme: opnTme, OpnBid: opnBid, OpnAsk: opnAsk, OpnSpd: opnSpd}", _sys.Ana.Trd.Adr(x), long)
// 	r.Add("ok, rsn := x.port.OpnTrd(t, i) // blocks until network completion")
// 	r.Add("if !ok {")
// 	r.Add("if ana.Cfg.Trc.IsRltStgy() {")
// 	r.Addf("sys.Logf(\"%v.%v %%p OPN FAIL %%v\", x, rsn)", node.Full(), r.Name)
// 	r.Add("}")
// 	r.Add("}")
// 	r.Add("x.Trd = t")
// 	r.Add("x.OpnIdx = opnIdx")
// 	r.Add("if ana.Cfg.Trc.IsRltStgy() {")
// 	r.Addf("sys.Logf(\"%v.%v %%p OPN %%v\", x, t)", node.Full(), r.Name)
// 	r.Add("}")
// 	r.Add("}")
// 	r.Add("}")
// 	r.Add("return nil")
// 	return r
// }

// func (x *FleRltStgy) RxClsLim(nodeFn *TypFn) (r *TypFn) {
// 	node := nodeFn.Node
// 	r = x.TypFna("RxClsLim", atr.None, node)
// 	r.InPrm(_sys.Ana.TmeIdx, "inPkt")
// 	r.OutPrmSlice(_sys.Act, "r")
// 	r.Add("// RxClsLim inPkt is i pkt 1s larger than cnds")
// 	r.Add("// place ClsLimEnd at end of tier processing ")
// 	r.Add("// so that it is called last")
// 	// r.Add("if ana.Cfg.Trc.IsRltStgy() {")
// 	// r.Addf("sys.Logf(\"%v %%p inPkt %%v\", x, inPkt)", r.Full())
// 	// r.Add("}")
// 	r.Addf("r = append(r, &%vRxClsLim{", node.Name)
// 	r.Add("X:     x,")
// 	r.Add("InPkt: inPkt,")
// 	r.Add("tier:  x.Cnd.DstToInstr() + 1,")
// 	r.Add("})")
// 	r.Add("return r")
// 	return r
// }
// func (x *FleRltStgy) RxClsLimEnd(nodeFn *TypFn) (r *TypFn) {
// 	node, isLong := nodeFn.Node, nodeFn.Lower() == k.Long
// 	r = x.TypFna("RxClsLimEnd", atr.None, node)
// 	r.InPrm(_sys.Ana.TmeIdx, "inPkt")
// 	r.OutPrmSlice(_sys.Act, "r")
// 	// r.Add("if ana.Cfg.Trc.IsRltStgy() {")
// 	// r.Addf("sys.Logf(\"%v.%v %%p inPkt %%v\", x, inPkt)", node.Full(), r.Name)
// 	// r.Add("}")
// 	r.Add("// heartbeat pkt has Idx == unt.max")
// 	r.Addf("if inPkt.Idx != %v && x.Trd != nil {", _sys.Bsc.Unt.Max.Ref(x))
// 	r.Add("x.mu.Lock() // cls by iument by profit, loss or expiration")
// 	r.Add("defer x.mu.Unlock()")
// 	r.Add("i := x.Instr.Instr()")
// 	r.Add("if x.Trd != nil {")
// 	r.Add("var clsRsn ana.TrdRsnCls")
// 	r.Add("if inPkt.Tme.Gtr(x.ClsTmeLim) {")
// 	r.Add("clsRsn = ana.Dur")
// 	r.Add("} else {")
// 	if isLong {
// 		r.Add("bid := i.RltStm.BidAt(inPkt.Idx) // Bid: the price for me to sell at")
// 		r.Add("if bid.Geq(x.ClsPrfLim) {")
// 		r.Add("clsRsn = ana.Prf")
// 		r.Add("} else if bid.Leq(x.ClsLosLim) {")
// 		r.Add("clsRsn = ana.Los")
// 		r.Add("}")
// 	} else {
// 		r.Add("ask := i.RltStm.AskAt(inPkt.Idx) // Ask: the price for me to buy at")
// 		r.Add("if ask.Leq(x.ClsPrfLim) {")
// 		r.Add("clsRsn = ana.Prf")
// 		r.Add("} else if ask.Geq(x.ClsLosLim) {")
// 		r.Add("clsRsn = ana.Los")
// 		r.Add("}")
// 	}
// 	r.Add("}")
// 	r.Add("if clsRsn != ana.NoTrdRsnCls && x.port.ClsTrd(x.Trd, i) {")
// 	r.Add("x.Trd.ClsTme = inPkt.Tme // set before ClsTrd, ClsTrd will use to calculate results")
// 	r.Add("x.Trd.ClsBid = i.RltStm.BidAt(inPkt.Idx)")
// 	r.Add("x.Trd.ClsAsk = i.RltStm.AskAt(inPkt.Idx)")
// 	r.Add("x.Trd.ClsSpd = i.Spd(x.Trd.ClsBid, x.Trd.ClsAsk)")
// 	r.Add("x.Trd.ClsRsn = clsRsn.Str()")
// 	r.Add("x.Trd.Dur = x.Trd.ClsTme.Sub(x.Trd.OpnTme)")
// 	r.Add("x.LstClsTme = inPkt.Tme")
// 	r.Add("x.LstClsIdx = inPkt.Idx")
// 	r.Add("x.port.CalcCls(x.Trd, i, true)")
// 	r.Add("for _, rx := range x.Rxs {")
// 	r.Add("r = append(r, ana.NewTrdTx(x.Trd, rx))")
// 	r.Add("}")
// 	r.Add("if ana.Cfg.Trc.IsRltStgy() {")
// 	r.Addf("sys.Logf(\"%v.%v %%p CLS %%v\", x, x.Trd)", node.Full(), r.Name)
// 	r.Add("}")
// 	r.Add("x.Trds.Push(x.Trd)")
// 	r.Add("x.port.Bse().Trds.Push(x.Trd)")
// 	r.Add("x.Trd = nil")
// 	r.Add("}")
// 	r.Add("}")

// 	r.Add("}")
// 	r.Add("return r")
// 	return r
// }

// func (x *FleRltStgy) RxClsCnd(nodeFn *TypFn) (r *TypFn) {
// 	node := nodeFn.Node
// 	r = x.TypFna("RxClsCnd", atr.None, node)
// 	r.InPrm(_sys.Bsc.Tme, "inPkt")
// 	r.OutPrmSlice(_sys.Act, "r")
// 	r.Add("if ana.Cfg.Trc.IsRltStgy() {")
// 	r.Addf("sys.Logf(\"%v.%v %%p inPkt %%v\", x, inPkt)", node.Full(), r.Name)
// 	r.Add("}")
// 	r.Add("if x.Trd != nil {")
// 	r.Add("x.mu.Lock()")
// 	r.Add("defer x.mu.Unlock()")
// 	r.Add("if x.Trd != nil {")
// 	r.Add("i := x.Instr.Instr()")
// 	r.Add("if x.port.ClsTrd(x.Trd, i) { // blocks until network completion")
// 	r.Add("idx := i.RltStm.Tmes.SrchIdxEql(inPkt)")
// 	r.Add("if idx == i.RltStm.Tmes.Cnt() {")
// 	r.Add("x.Trd.ClsTme = inPkt")
// 	r.Add("x.LstClsTme = inPkt")
// 	r.Add("sys.Logf(\"Stgy.RxClsCnd MISSING RltStm TME: TRD RECORD FAULTY %p inPkt:%v\", x, inPkt)")
// 	r.Add("} else {")
// 	r.Add("x.Trd.ClsTme = inPkt")
// 	r.Add("x.Trd.ClsBid = i.RltStm.BidAt(idx)")
// 	r.Add("x.Trd.ClsAsk = i.RltStm.AskAt(idx)")
// 	r.Add("x.Trd.ClsSpd = i.Spd(x.Trd.ClsBid, x.Trd.ClsAsk)")
// 	r.Add("x.Trd.ClsRsn = ana.Cnd.Str()")
// 	r.Add("x.Trd.Dur = x.Trd.ClsTme.Sub(x.Trd.OpnTme)")
// 	r.Add("x.LstClsTme = inPkt")
// 	r.Add("x.LstClsIdx = idx")
// 	r.Add("x.port.CalcCls(x.Trd, i, true)")
// 	r.Add("if ana.Cfg.Trc.IsRltStgy() {")
// 	r.Addf("sys.Logf(\"%v.%v %%p CLS %%v\", x, x.Trd)", node.Full(), r.Name)
// 	r.Add("}")
// 	r.Add("}")
// 	r.Add("for _, rx := range x.Rxs {")
// 	r.Add("r = append(r, ana.NewTrdTx(x.Trd, rx))")
// 	r.Add("}")
// 	r.Add("x.Trds.Push(x.Trd)")
// 	r.Add("x.port.Bse().Trds.Push(x.Trd)")
// 	r.Add("x.Trd = nil")
// 	r.Add("}")
// 	r.Add("}")
// 	r.Add("}")
// 	r.Add("return r")
// 	return r
// }
