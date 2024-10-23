package hst

import (
	"regexp"
	"strings"
	"sys"
	"sys/ana"
	"sys/bsc/bnds"
	"sys/bsc/flt"
	"sys/bsc/tme"
	"sys/bsc/tmes"
	"sys/bsc/unt"
	"sys/trc"
)

func (x *StgyBse) OpnClsTrd(opnTme tme.Tme) (r *ana.Trd, rsn ana.TrdRsnOpn) {
	//sys.Log("hst.StgyBse.OpnClsTrd: opnTme", opnTme)

	// EXPECT THAT THE OPN TME EXISTS
	// commission calc: https://www.oanda.com/register/docs/oc/price-sheet.pdf
	instr := x.Instr.Bse().Ana
	hstStm := instr.HstStm
	mktWeekIdx := instr.MktWeeks.SrchIdx(opnTme)
	if mktWeekIdx == instr.MktWeeks.Cnt() {
		sys.Logf("CANNOT FIND MKT WEEK %v (opnTme:%v mktRngs:%v)", instr.Name, opnTme, instr.MktWeeks)
		return nil, ana.ErrMktWeek
	}
	mktWeek := (*instr.MktWeeks)[mktWeekIdx]
	mktWeek.Min += ana.Cfg.MktTrdBuf
	mktWeek.Max -= ana.Cfg.MktTrdBuf // put buffer before mkt close; spds become extremely large at mkt cls; don't want to exit then

	if opnTme < mktWeek.Min {
		return nil, ana.NearMktOpn // avoid opening trades near market opn; large vol and spds
	}
	if opnTme > mktWeek.Max-x.DurLim {
		return nil, ana.NearMktCls // avoid opening trades near market cls; large vol and spds, and allow room for stgy completion
	}
	opnIdx := hstStm.Tmes.SrchIdxEql(opnTme) // lookup instr index for cnd time
	if opnIdx == hstStm.Tmes.Cnt() {
		sys.Logf("CANNOT FIND OPNCND TIME IN INSTRUMENT %v (opnTme:%v)", instr.Name, opnTme)
		return nil, ana.ErrOpnCnd
	}
	opnBid, opnAsk := hstStm.BidAskAt(opnIdx)
	opnSpd := instr.Spd(opnBid, opnAsk)
	if opnSpd > instr.SpdOpnLim {
		return nil, ana.SpdLrg // avoid opening trades with large spreads
	}
	// Bid: the price for me to sell at
	// Ask: the price for me to buy at
	// PrfLim: 1.1   Pip: 0.0001     1.1 × .0001 = 0.00011
	clsTmeLim := opnTme.Add(x.DurLim).Min(mktWeek.Max) // cls by time limit
	var clsPrfLim, clsLosLim flt.Flt
	if x.IsLong { // TODO: IS THIS CORRECT?
		clsPrfLim = opnAsk + (x.PrfLim * instr.Pip)
		clsLosLim = opnAsk - (x.LosLim * instr.Pip)
	} else {
		clsPrfLim = opnBid - (x.PrfLim * instr.Pip)
		clsLosLim = opnBid + (x.LosLim * instr.Pip)
	}
	// sys.Logf("clsTmeLim:%v clsPrfLim:%v clsLosLim:%v", clsTmeLim, clsPrfLim, clsLosLim)
	var clsdLim bool
	// linear search for exit by prf, los, durLim (success, fail, expire)
	for stmClsIdx := opnIdx + 1; stmClsIdx < hstStm.Tmes.Cnt(); stmClsIdx++ {
		if hstStm.Tmes.At(opnIdx) == hstStm.Tmes.At(stmClsIdx) || hstStm.Tmes.At(opnIdx)+tme.S1 == hstStm.Tmes.At(stmClsIdx) {
			continue // mirror rlt behavior; rlt cannot close lim within same second
		}
		var clsRsn ana.TrdRsnCls
		// sys.Logf("hstStm.Tmes.At(stmClsIdx):%v", hstStm.Tmes.At(stmClsIdx))
		if hstStm.Tmes.At(stmClsIdx).Gtr(clsTmeLim) {
			clsRsn = ana.Dur
		} else {
			if x.IsLong {
				bid := hstStm.BidAt(stmClsIdx) // Bid: the price for me to sell at
				if bid >= clsPrfLim {
					clsRsn = ana.Prf
				} else if bid <= clsLosLim {
					clsRsn = ana.Los
				}
			} else {
				ask := hstStm.AskAt(stmClsIdx) // Ask: the price for me to buy at
				if ask <= clsPrfLim {
					clsRsn = ana.Prf
				} else if ask >= clsLosLim {
					clsRsn = ana.Los
				}
			}
		}
		if clsRsn != ana.NoTrdRsnCls {
			r = &ana.Trd{}
			r.IsLong = x.IsLong
			r.OpnTme = opnTme
			r.OpnBid = opnBid
			r.OpnAsk = opnAsk
			r.OpnSpd = opnSpd
			r.ClsTme = hstStm.Tmes.At(stmClsIdx)
			r.ClsBid = hstStm.BidAt(stmClsIdx)
			r.ClsAsk = hstStm.AskAt(stmClsIdx)
			r.ClsRsn = clsRsn.Str()
			r.TrdPct = ana.Cfg.Hst.TrdPct
			clsdLim = true
			break
		}
	}
	var cndClsTmes []tme.Tme // search for close conditions
	for _, cnd := range x.Clss {
		if cnd.Bse().Tmes.Cnt() != 0 {
			cndBse := cnd.Bse()
			cndClsIdx := cndBse.Tmes.SrchIdx(opnTme, true)
			if cndClsIdx != unt.Max && cndClsIdx < cndBse.Tmes.Cnt() {
				if cndBse.Tmes.At(cndClsIdx) == opnTme {
					if cndClsIdx+1 < cndBse.Tmes.Cnt() && cndBse.Tmes.At(cndClsIdx+1) > opnTme {
						cndClsTmes = append(cndClsTmes, cndBse.Tmes.At(cndClsIdx+1))
					}
				} else {
					if cndBse.Tmes.At(cndClsIdx) > opnTme {
						cndClsTmes = append(cndClsTmes, cndBse.Tmes.At(cndClsIdx))
					}
				}
			}
		}
	}
	if len(cndClsTmes) != 0 {
		cndClsTme := tmes.New(cndClsTmes...).Min()
		if clsdLim {
			// r = trds.Lst()
			if cndClsTme < r.ClsTme { // cndClsTme lss than bounds exit time
				clsStmIdx := hstStm.Tmes.SrchIdxEql(cndClsTme)
				r.ClsTme = cndClsTme
				r.ClsBid = hstStm.BidAt(clsStmIdx)
				r.ClsAsk = hstStm.AskAt(clsStmIdx)
				r.ClsRsn = ana.Cnd.Str()
			}
		} else { // only cnd cls
			clsStmIdx := hstStm.Tmes.SrchIdxEql(cndClsTme)
			r = &ana.Trd{}
			r.IsLong = x.IsLong
			r.OpnTme = opnTme
			r.OpnBid = opnBid
			r.OpnAsk = opnAsk
			r.OpnSpd = opnSpd
			r.ClsTme = cndClsTme
			r.ClsBid = hstStm.BidAt(clsStmIdx)
			r.ClsAsk = hstStm.AskAt(clsStmIdx)
			r.ClsRsn = ana.Cnd.Str()
		}
	}
	if r != nil {
		r.ClsSpd = instr.Spd(r.ClsBid, r.ClsAsk)
		r.Dur = r.ClsTme.Sub(r.OpnTme)
		instr.Prv.CalcOpn(r, instr) // set trd flds
		instr.Prv.CalcCls(r, instr) // set trd flds
	}
	return r, ana.NoCls
}

func (x *StgyStgy) Fit() {
	trcr0 := trc.New("StgyStgy.Fit")
	defer trcr0.End()
	cndTmes := x.Cnd.Bse().Tmes

	// CALC PLL TRDS; ALLOW OVERLAPPING TRD TMES FOR INCREASED ML MODEL EXAMPLES
	trcr := trc.New("StgyStgy.Fit: PLL TRDS")
	pllTrds := ana.NewTrds()
	segBnds, acts := bnds.Segs(cndTmes.Cnt())
	for n, segBnd := range *segBnds {
		seg := &StgyStgySeg{}
		seg.Bnd = segBnd
		seg.Stgy = x.Bse()
		seg.Tmes = cndTmes
		seg.Out = ana.NewTrds()
		acts[n] = seg
		// break // TODO: REMOVE
	}
	sys.Run().Pll(acts...)
	for _, act := range acts {
		pllTrds.Mrg(act.(*StgyStgySeg).Out)
	}
	trcr.End()
	if len(*pllTrds) == 0 {
		// NO ML TRAINING OR ML KEY
		return
	}
	// ML FEATURES: GATHER (PLL TRD)
	trcr = trc.New("StgyStgy.Fit: GATHER (PLL TRD)")
	rltKey := x.String()
	rltKey = strings.Replace(rltKey, "hst.", "rlt.", -1) // CHANGE TO RLT FOR RLT LOOKUP
	ftrNames := make([]string, len(*x.FtrStms))
	ftrs := make([][]float32, len(*x.FtrStms))
	lbls := make([]float32, pllTrds.Cnt())
	pllTrdTmes := tmes.Make(pllTrds.Cnt())
	for n, trd := range *pllTrds {
		(*pllTrdTmes)[n] = trd.OpnTme // gather trade times
		lbls[n] = float32(trd.PnlPct) // gather ml feature labels
	}
	for n := 0; n < len(*x.FtrStms); n++ {
		ftrNames[n] = (*x.FtrStms)[n].String()    // gather ml feature column names
		ftrs[n] = (*x.FtrStms)[n].Atf(pllTrdTmes) // gather ml feature values
	}
	trcr.End()
	// ML MODEL: FIT (PLL TRD)
	trcr = trc.New("StgyStgy.Fit: FIT (PLL TRD)")
	fitBits := sys.Lrnr().Fit(rltKey, ftrNames, ftrs, lbls)
	trcr.End()
	if len(fitBits) == 0 {
		// NOT ENOUGH DATA TO PRODUCE MODEL
		return
	}
	if sys.HasDsk() { // may be nil during testing
		// ML MODEL: DSK SAV (PLL TRD) FOR RLT RETRIEVAL
		trcr = trc.New("StgyStgy.Fit: DSK SAV (PLL TRD)")
		sys.Lrnr().SaveNetToDsk(rltKey, fitBits)
		trcr.End()
	}
	// if ana.Cfg.Hst.StgySeqTrd { // SEQ TRD ONLY NECESSARY FOR HST/RLT PARITY TESTING
	// 	x.CalcSeqTrds(key)
	// }
}

func (x *StgyStgy) CalcSeqTrds(rltKey string) {
	// CALC SEQ TRDS; SIMULATE RLT TRADING
	trcr := trc.New("StgyStgy.CalcSeqTrds: SEQ TRD")
	defer trcr.End()
	x.Trds = ana.NewTrds()
	cndTmes := x.Cnd.Bse().Tmes
	nxtAvailOpnTme := tme.Min
	var opnTme, prvOpnTme tme.Tme
	for {
		opnTme = tme.Max
		idx := cndTmes.SrchIdxEql(nxtAvailOpnTme)
		if int(idx) < len(*cndTmes) && (*cndTmes)[idx] < opnTme {
			opnTme = (*cndTmes)[idx]
		}
		if opnTme == prvOpnTme {
			nxtAvailOpnTme += tme.Resolution
			continue
		}
		if opnTme == tme.Max {
			break
		}
		// ML FEATURES: GATHER (SEQ TRD)
		xftrs := make([]float32, len(*x.FtrStms))
		for n := 0; n < len(*x.FtrStms); n++ {
			stmBse := (*x.FtrStms)[n].Bse()
			idx := stmBse.Tmes.SrchIdxEql(opnTme)
			if int(idx) < len(*stmBse.Tmes) && (*stmBse.Tmes)[idx] == opnTme {
				xftrs[n] = float32((*stmBse.Vals)[idx])
			} else {
				sys.Logf("SEQ TRD: UNABLE TO FIND FtmStm VAL (n:%v opnTme:%v ftrStm:%v)", n, opnTme, (*x.FtrStms)[n])
			}
		}
		// sys.Logf("ML MODEL: PREDICT: xftrs:%v", xftrs)
		// ML MODEL: PREDICT (SEQ TRD)
		// GENERATE SEQ TRD FOR HST.RLT PARITY CHECK
		pnlPctPredict := flt.Flt(sys.Lrnr().Predict(rltKey, xftrs))
		// sys.Logf("- opnTme:%v pnlPctPredict:%v", opnTme, pnlPctPredict)
		// sys.Lrnr().LoadNet(key, fitBits)
		// pnlPctPredict = flt.Flt(sys.Lrnr().Predict(key, xftrs))
		// sys.Logf(" *** StgyStgy.CalcSeqTrds opnTme:%v pnlPctPredict:%v", opnTme, pnlPctPredict)

		if pnlPctPredict > x.MinPnlPct {
			trd, rsnFail := x.OpnClsTrd(opnTme)
			if trd != nil { // may fail to close due to near mkt opn, near mkt cls, spd lim exceeded
				trd.PnlPctPredict = pnlPctPredict
				x.Trds.Push(trd)
				// sys.Logf("* opnTme:%v pnlPctPredict:%v actualPnlPct:%v dltPnlPct:%v", opnTme, pnlPctPredict, trd.PnlPct, trd.PnlPct-pnlPctPredict)
				nxtAvailOpnTme = trd.ClsTme
			} else {
				if rsnFail == ana.NoCls {
					break // exit last opn fail to mirror rlt behavior; single ana.NoCls expected
				}
				nxtAvailOpnTme = opnTme // trd failed to cls; advance to next open time
			}
		} else {
			nxtAvailOpnTme = opnTme // trd failed to opn (pnlPctPredict filter); advance to next open time
		}
		prvOpnTme = opnTme
	}
}

func (x *StgyStgy) Prfm(rng tme.Rng) (r *ana.Prfm) {
	originalKey := x.String()

	// CREATE NEW STGY FROM EXISTING STGY AND TME RNG
	// SCRIPT: REPLACE RNG PRM IN ALL INSTR
	// ASSUMES CURRENT STGY HAS RNG DEFINED IN INSTR
	revisedKey := originalKey
	re := regexp.MustCompile("hst.[[:word:]]+[(][)].[[:word:]]+[(]([^)]+)")
	matches := re.FindStringSubmatch(revisedKey)
	if len(matches) > 1 {
		for n := 1; n < len(matches); n++ {
			re = regexp.MustCompile(matches[n])
			revisedKey = re.ReplaceAllString(revisedKey, rng.String())
		}
	}
	rltKey := strings.Replace(originalKey, "hst.", "rlt.", -1) // CHANGE TO RLT FOR RLT LOOKUP
	sys.Log("--- originalKey", originalKey)
	sys.Log("---  revisedKey", revisedKey)
	sys.Log("---      rltKey", rltKey)
	vs := sys.Actr().RunIfc(revisedKey)
	stgyAtRng := vs[len(vs)-1].(*StgyStgy)
	stgyAtRng.CalcSeqTrds(rltKey)
	sys.Log("--- stgyAtRng.Trds", stgyAtRng.Trds.Cnt())
	r = ana.CalcPrfm(stgyAtRng.Trds, rng.Min.WeekdayCnt(rng.Max), x.LosLim, x.DurLim, revisedKey)
	return r
}
