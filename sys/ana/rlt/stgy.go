package rlt

import (
	"sys"
	"sys/ana"
	"sys/bsc/flt"
	"sys/bsc/flts"
	"sys/bsc/tme"
	"sys/bsc/tmes"
	"sys/bsc/unt"
)

type (
	StgyStgyEvalOpn struct {
		X    *StgyStgy
		ret  []sys.Act
		tier int
	}

	// StgyStgyRxOpn struct {
	// 	X     *StgyStgy
	// 	InPkt tme.Tme
	// 	ret   []sys.Act
	// 	tier  int
	// }
	StgyStgyRxClsLim struct {
		X     *StgyStgy
		InPkt ana.TmeIdx
		ret   []sys.Act
		tier  int
	}
	StgyFtrStm struct {
		Stgy *StgyStgy
		Name string
		Tmes *tmes.Tmes
		Vals *flts.Flts
		idx  unt.Unt
	}
)

func (x *StgyStgyEvalOpn) Act() { // sys.Act interface
	x.ret = x.X.EvalOpnEnd()
}
func (x *StgyStgyEvalOpn) DecTier()       { x.tier-- }      // prv.Tx interface
func (x *StgyStgyEvalOpn) Ret() []sys.Act { return x.ret }  // prv.Tx interface
func (x *StgyStgyEvalOpn) Tier() int      { return x.tier } // prv.Tx interface

// func (x *StgyStgyRxOpn) Act() { // sys.Act interface
// 	x.ret = x.X.RxOpnEnd(x.InPkt)
// }
// func (x *StgyStgyRxOpn) DecTier()       { x.tier-- }      // prv.Tx interface
// func (x *StgyStgyRxOpn) Ret() []sys.Act { return x.ret }  // prv.Tx interface
// func (x *StgyStgyRxOpn) Tier() int      { return x.tier } // prv.Tx interface

func (x *StgyStgyRxClsLim) Act() { // sys.Act interface
	x.ret = x.X.RxClsLimEnd(x.InPkt)
}
func (x *StgyStgyRxClsLim) DecTier()       { x.tier-- }      // prv.Tx interface
func (x *StgyStgyRxClsLim) Ret() []sys.Act { return x.ret }  // prv.Tx interface
func (x *StgyStgyRxClsLim) Tier() int      { return x.tier } // prv.Tx interface

func (x *StgyFtrStm) Rx(inPkt ana.TmeFlt) (r []sys.Act) {
	if ana.Cfg.Trc.IsRltStgy() {
		sys.Logf("rlt.StgyStgy StgyFtrStm.Rx %p inPkt %v %v", x, inPkt, x.Name)
	}
	x.Stgy.mu.Lock()
	x.Tmes.Push(inPkt.Tme)
	x.Vals.Push(inPkt.Flt)
	x.Stgy.mu.Unlock()
	r = append(r, &StgyStgyEvalOpn{
		X:    x.Stgy,
		tier: x.Stgy.Cnd.DstToInstr() + 1,
	})
	return r
}
func (x *StgyStgy) RxOpn(inPkt tme.Tme) (r []sys.Act) {
	// RxOpnEnd allows cnd in same graph phase to cls a trd
	// and for RxOpnEnd to then open a new trade within the same 1s
	if ana.Cfg.Trc.IsRltStgy() {
		sys.Logf("rlt.StgyStgy.RxOpn %p inPkt %v", x, inPkt)
	}
	if x.Instr.Bse().Prv.MayTrd() { // IMPORTANT TO AVOID BUG WHERE OPNING NEW TRD WHILE EXISTING ONE ALREADY OPN
		x.mu.Lock()
		x.opns.Push(inPkt)
		x.mu.Unlock()
		r = append(r, &StgyStgyEvalOpn{
			X:    x,
			tier: x.Cnd.DstToInstr() + 1,
		})
	}
	return r
}
func (x *StgyStgy) EvalOpnEnd() (r []sys.Act) {
	// if ana.Cfg.Trc.IsRltStgy() {
	// 	sys.Logf("rlt.StgyStgy.EvalOpnEnd %p", x)
	// }
	x.mu.Lock()
	defer x.mu.Unlock()
	if len(*x.opns) == 0 {
		// FTR STM MAY PRODUCE EVALUATION WITHOUT OPN TME
		return r
	}
	opnTme := x.opns.Fst()
	// ENSURE EACH FTR STM HAS PRODUCED A VAL FOR THE CUR OPN TME
	ok := true
	for _, stgyFtrStm := range x.stgyFtrStms {
		if len(*stgyFtrStm.Tmes) == 0 {
			ok = false
			break
		}
		stgyFtrStm.idx = stgyFtrStm.Tmes.SrchIdxEql(opnTme)
		if stgyFtrStm.idx == unt.Unt(len(*stgyFtrStm.Tmes)) ||
			opnTme != (*stgyFtrStm.Tmes)[stgyFtrStm.idx] {
			ok = false
			break
		}
	}
	if !ok {
		// WAITING FOR ALL FTR STMS TO BE RECEIVED FOR CUR OPN TME
		return nil
	}
	if ana.Cfg.Trc.IsRltStgy() {
		sys.Logf("rlt.StgyStgy.EvalOpnEnd %p opnTme %v", x, opnTme)
	}
	if x.Instr.Bse().Prv.MayTrd() && x.Trd == nil && opnTme.Geq(x.LstClsTme) {
		i := x.Instr.Instr()
		opnIdx := i.RltStm.Tmes.SrchIdxEql(opnTme)
		if opnIdx < x.LstClsIdx {
			opnIdx = x.LstClsIdx
		}
		mktWeekMin := ana.MktWeekMin(opnTme) + ana.Cfg.MktTrdBuf
		mktWeekMax := ana.MktWeekMax(opnTme) - ana.Cfg.MktTrdBuf
		if opnTme < mktWeekMin || opnTme > mktWeekMax-x.DurLim {
			return nil // avoid opening trades near market close time
		}
		// FILTER THROUGH ML PREDICTION
		xftrs := make([]float32, len(x.stgyFtrStms))
		for n, stgyFtrStm := range x.stgyFtrStms {
			xftrs[n] = float32((*stgyFtrStm.Vals)[stgyFtrStm.idx])
		}
		// DRAIN PRIOR TO PREDICT INCASE JULIA CRASH IN PROD
		// DRAIN FTR STMS THROUGH & INCLUDING CURRENT VAL
		for _, stgyFtrStm := range x.stgyFtrStms {
			stgyFtrStm.Tmes = stgyFtrStm.Tmes.From(stgyFtrStm.idx + 1)
			stgyFtrStm.Vals = stgyFtrStm.Vals.From(stgyFtrStm.idx + 1)
		}
		x.opns.Dque()

		// sys.Logf("- x.Key:%v", x.Key)
		// sys.Logf("- xftrs:%v", xftrs)
		// sys.Lrnr().(*ml.Lrnr).Opnd = false
		pnlPctPredict := flt.Flt(sys.Lrnr().Predict(x.Key, xftrs))
		// pnlPctPredict := flt.Flt(x.lrnr.Predict(x.Key, xftrs))
		if ana.Cfg.Trc.IsRltStgy() {
			sys.Logf("rlt.StgyStgy.EvalOpnEnd opnTme:%v pnlPctPredict:%v > x.MinPnlPct:%v = %v", opnTme, pnlPctPredict, x.MinPnlPct, pnlPctPredict > x.MinPnlPct)
		}
		if pnlPctPredict > x.MinPnlPct {
			x.ClsTmeLim = opnTme.Add(x.DurLim).Min(mktWeekMax) // avoid holding postion after market close
			// Bid: the price for me to sell at
			// Ask: the price for me to buy at
			opnBid, opnAsk := i.RltStm.BidAskAt(opnIdx)
			opnSpd := i.Spd(opnBid, opnAsk)
			if opnSpd > i.SpdOpnLim {
				return nil // avoid opening trades with large spreads
			}
			// PrfLim: 1.1   Pip: 0.0001     1.1 × .0001 = 0.00011
			if x.IsLong {
				x.ClsPrfLim = opnAsk + (x.PrfLim * i.Pip)
				x.ClsLosLim = opnAsk - (x.LosLim * i.Pip)
			} else {
				x.ClsPrfLim = opnBid - (x.PrfLim * i.Pip)
				x.ClsLosLim = opnBid + (x.LosLim * i.Pip)
			}
			t := &ana.Trd{OpnTme: opnTme, OpnBid: opnBid, OpnAsk: opnAsk, OpnSpd: opnSpd}
			t.IsLong = x.IsLong
			t.PnlPctPredict = pnlPctPredict
			t.TrdPct = ana.Cfg.Rlt.TrdPct
			if ana.Cfg.Test { // USE HST TRDPCT WHILE TESTING (FOR PARITY TESTING)
				t.TrdPct = ana.Cfg.Hst.TrdPct
			}
			ok, rsn := i.Prv.OpnTrd(t, i) // blocks until network completion
			if !ok {
				if ana.Cfg.Trc.IsRltStgy() {
					sys.Logf("rlt.StgyStgy.RxOpnEnd %p OPN FAIL %v", x, rsn)
				}
			}
			x.Trd = t
			x.OpnIdx = opnIdx
			if ana.Cfg.Trc.IsRltStgy() {
				sys.Logf("rlt.StgyStgy.RxOpnEnd %p OPN %v", x, t)
			}
		}
	}
	return nil
}
func (x *StgyStgy) RxClsLim(inPkt ana.TmeIdx) (r []sys.Act) {
	// RxClsLim inPkt is i pkt 1s larger than cnds
	// place ClsLimEnd at end of tier processing
	// so that it is called last
	r = append(r, &StgyStgyRxClsLim{
		X:     x,
		InPkt: inPkt,
		tier:  x.Cnd.DstToInstr() + 1,
	})
	return r
}
func (x *StgyStgy) RxClsLimEnd(inPkt ana.TmeIdx) (r []sys.Act) {
	// heartbeat pkt has Idx == unt.max
	if inPkt.Idx != unt.Max && x.Trd != nil {
		x.mu.Lock() // cls by iument by profit, loss or expiration
		defer x.mu.Unlock()
		i := x.Instr.Instr()
		if x.Trd != nil {
			var clsRsn ana.TrdRsnCls
			if inPkt.Tme.Gtr(x.ClsTmeLim) {
				clsRsn = ana.Dur
			} else {
				if x.IsLong {
					bid := i.RltStm.BidAt(inPkt.Idx) // Bid: the price for me to sell at
					if bid >= x.ClsPrfLim {
						clsRsn = ana.Prf
					} else if bid <= x.ClsLosLim {
						clsRsn = ana.Los
					}
				} else {
					ask := i.RltStm.AskAt(inPkt.Idx) // Ask: the price for me to buy at
					if ask <= x.ClsPrfLim {
						clsRsn = ana.Prf
					} else if ask >= x.ClsLosLim {
						clsRsn = ana.Los
					}
				}
			}
			if clsRsn != ana.NoTrdRsnCls && i.Prv.ClsTrd(x.Trd, i) {
				x.Trd.ClsTme = inPkt.Tme // set before ClsTrd, ClsTrd will use to calculate results
				x.Trd.ClsBid = i.RltStm.BidAt(inPkt.Idx)
				x.Trd.ClsAsk = i.RltStm.AskAt(inPkt.Idx)
				x.Trd.ClsSpd = i.Spd(x.Trd.ClsBid, x.Trd.ClsAsk)
				x.Trd.ClsRsn = clsRsn.Str()
				x.Trd.Dur = x.Trd.ClsTme.Sub(x.Trd.OpnTme)
				x.LstClsTme = inPkt.Tme
				x.LstClsIdx = inPkt.Idx
				i.Prv.CalcCls(x.Trd, i)
				if ana.Cfg.Trc.IsRltStgy() {
					sys.Logf("rlt.StgyStgy.RxClsLimEnd %p CLS %v", x, x.Trd)
				}
				for _, rx := range x.Rxs {
					r = append(r, ana.NewTrdTx(x.Trd, rx))
				}
				x.Trd = nil
			}
		}
	}
	return r
}

func (x *StgyStgy) RxClsCnd(inPkt tme.Tme) (r []sys.Act) {
	if ana.Cfg.Trc.IsRltStgy() {
		sys.Logf("rlt.StgyStgy.RxClsCnd %p inPkt %v", x, inPkt)
	}
	if x.Trd != nil {
		x.mu.Lock()
		defer x.mu.Unlock()
		if x.Trd != nil {
			i := x.Instr.Instr()
			if i.Prv.ClsTrd(x.Trd, i) { // blocks until network completion
				idx := i.RltStm.Tmes.SrchIdxEql(inPkt)
				if idx == i.RltStm.Tmes.Cnt() {
					x.Trd.ClsTme = inPkt
					x.LstClsTme = inPkt
					sys.Logf("Stgy.RxClsCnd MISSING RltStm TME: TRD RECORD FAULTY %p inPkt:%v", x, inPkt)
				} else {
					x.Trd.ClsTme = inPkt
					x.Trd.ClsBid = i.RltStm.BidAt(idx)
					x.Trd.ClsAsk = i.RltStm.AskAt(idx)
					x.Trd.ClsSpd = i.Spd(x.Trd.ClsBid, x.Trd.ClsAsk)
					x.Trd.ClsRsn = ana.Cnd.Str()
					x.Trd.Dur = x.Trd.ClsTme.Sub(x.Trd.OpnTme)
					x.LstClsTme = inPkt
					x.LstClsIdx = idx
					i.Prv.CalcCls(x.Trd, i)
					if ana.Cfg.Trc.IsRltStgy() {
						sys.Logf("rlt.StgyStgy.RxClsCnd %p CLS %v", x, x.Trd)
					}
				}
				for _, rx := range x.Rxs {
					r = append(r, ana.NewTrdTx(x.Trd, rx))
				}
				x.Trd = nil
			}
		}
	}
	return r
}
