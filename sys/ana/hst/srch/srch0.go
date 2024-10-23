package srch

// import (
// 	"sys/ana/hst"
// 	"sys/bsc/flt"
// 	"sys/bsc/flts"
// 	"sys/bsc/tme"
// )

// type (
// 	PrcpStm struct {
// 		Dur     tme.Tme
// 		PnlPct  flt.Flt
// 		FwdPrfm hst.Prfm
// 		Stm     hst.Stm
// 	}
// 	PrcpStmDurGrp struct {
// 		Dur      tme.Tme
// 		PrcpStms PrcpStms
// 	}
// 	PrcpStmDurGrps map[tme.Tme]*PrcpStmDurGrp
// 	PrcpStms       []*PrcpStm
// 	PrcpDur        struct {
// 		Dur        tme.Tme
// 		FwdPnlPcts *flts.Flts
// 	}
// )

// func Srch0() {
// 	durs := tmes.New(
// 		tme.M1, tme.M1*2, tme.M1*3,
// 		tme.M1*5, tme.M1*8, tme.M1*13,
// 		tme.M1*21, tme.M1*34, tme.M1*55,
// 		tme.M1*89, tme.M1*144, tme.M1*233, tme.M1*377,
// 	)
// 	prfLims := flts.AddsLeq(10.0, 10.0, 5.0)
// 	losLims := flts.AddsLeq(10.0, 10.0, 5.0)
// 	durLims := tmes.AddsLeq(tme.M15, tme.M15, tme.M15)
// 	mktDays := hst.Oan().EurUsd().Bse().Ana.MktDays

// 	bckTrnLen := 2
// 	stmCntLim := unt.Unt(2)
// 	trimItrLim := unt.Unt(3)
// 	trimMin := unt.Unt(4)
// 	trimForgiveLim := unt.Unt(1)
// 	sys.Logf("stmCntLim:%v bckTrnLen:%v trimItrLim:%v", stmCntLim, bckTrnLen, trimItrLim)
// 	sys.Logf("trimMin:%v trimForgiveLim:%v", trimMin, trimForgiveLim)
// 	prcpStms := make([]*PrcpStm, 0, 512)
// 	sesStart := time.Now()
// 	for _, dur := range *durs {
// 		durStart := time.Now()
// 		prcpDur := &PrcpDur{Dur: dur, FwdPnlPcts: flts.New()}
// 		for dayIdx := 0; dayIdx < len(*mktDays)-1-bckTrnLen; dayIdx++ {
// 			// mktDay := (*mktDays)[idx]
// 			trnStart := time.Now()
// 			prcpStm := &PrcpStm{Dur: dur}
// 			prcpStms = append(prcpStms, prcpStm)

// 			bckRng := (*mktDays)[dayIdx]
// 			sys.Logf("TRN: START:  dayIdx:%v  bckRng:%v", dayIdx, bckRng)
// 			instr := hst.Oan().EurUsd(bckRng)
// 			prcp := hst.NewPrcp()
// 			stm := instr.I(dur).Ask().Vrnc()
// 			prcp.Stm(stm)
// 			prcpStm.Stm = stm
// 			// qckStm := instr.I(tme.M5).Ask().Alma(6, 0.85)
// 			// slwStm := instr.I(tme.M20).Ask().Sma()
// 			// opnCnd := qckStm.OtrLss(0, slwStm).Seq(tme.S1, qckStm.OtrGtr(0, slwStm))
// 			// MACD CRS
// 			h1AskLst := instr.I(tme.H1).Ask().Lst()
// 			qckStm := h1AskLst.AggEma(12).OtrSub(0, h1AskLst.AggEma(26))
// 			slwStm := qckStm.AggEma(9)
// 			opnCnd := slwStm.SclGtr(0.0).And(qckStm.OtrLss(0, slwStm).Seq(tme.S1, qckStm.OtrGtr(0, slwStm)))
// 			for _, prfLim := range *prfLims {
// 				for _, losLim := range *losLims {
// 					for _, durLim := range *durLims {
// 						// BCK STGY
// 						stgy := opnCnd.Long(prfLim, losLim, durLim, instr)
// 						// sys.Log("stgy:", stgy)
// 						prfm := stgy.Port().Prfm()
// 						splt := prfm.Port().Splt(0.0)
// 						prfm = prcp.Splt(splt).TuneSacfTil(
// 							0.0, stmCntLim, trimItrLim, trimMin, trimForgiveLim,
// 						)
// 						sys.Logf("bck fst prfm:%v", prfm.Ana())
// 						// BCK TUNE
// 						lim := dayIdx + bckTrnLen
// 						for dayTunIdx := dayIdx + 1; dayTunIdx < lim; dayTunIdx++ {
// 							// prcp.Sub()
// 							bckTunRng := (*mktDays)[dayTunIdx]
// 							splt := prfm.Port().Rng(bckTunRng).Splt(0.0)
// 							prfm = prcp.Splt(splt).TuneSacfTil(
// 								0.0, stmCntLim, trimItrLim, trimMin, trimForgiveLim,
// 							)
// 							sys.Logf("bck tune prfm:%v", prfm.Ana())
// 							// prcp.Unsub()
// 						}
// 						// FWD TST
// 						fwdIdx := dayIdx + bckTrnLen
// 						sys.Log(" ### bckRng ###", mktDays.RngMrg(unt.Unt(dayIdx), unt.Unt(fwdIdx-1)))
// 						fwdRng := (*mktDays)[fwdIdx]
// 						sys.Log(" ### fwdRng ###", fwdRng)
// 						// prcp.Sub()
// 						prfm = prfm.Port().Rng(fwdRng).Prfm()
// 						sys.Logf("fwd prfm:%v", prfm.Ana())
// 						// prcp.Unsub()

// 						// if len(*bckPfmsA) == 0 {
// 						// 	// fwdPnlPcts.Push(0)
// 						// 	prcpDur.FwdPnlPcts.Push(0)
// 						// 	prcpStm.PnlPct = flt.Min
// 						// 	sys.Log("NONE")
// 						// } else {
// 						// 	bckPrfm := bckPfmsA.SrtDscPnlPct().Fst()
// 						// 	sys.Log("bckPrfm", bckPrfm.Ana())
// 						// 	fwdPrfm := bckPrfm.Port().Rng(fwdRng).Prfm()
// 						// 	sys.Log("fwdPrfm", fwdPrfm.Ana())
// 						// 	prcpDur.FwdPnlPcts.Push(fwdPrfm.Ana().PnlPct)
// 						// 	prcpStm.FwdPrfm = fwdPrfm
// 						// 	prcpStm.PnlPct = fwdPrfm.Ana().PnlPct
// 						// }
// 					}
// 				}
// 			}

// 			// fwdIdx := dayIdx + bckTrnLen
// 			// for bckIdx := dayIdx; bckIdx < fwdIdx; bckIdx++ {
// 			// 	bckRng := (*mktDays)[bckIdx]
// 			// 	sys.Logf("TRN: START: bckIdx:%v bckRng:%v len(*bckPfmsA):%v", bckIdx, bckRng, len(*bckPfmsA))
// 			// 	instr := hst.Oan().EurUsd(bckRng)
// 			// 	prcp := hst.NewPrcp()
// 			// 	stm := instr.I(dur).Ask().Vrnc()
// 			// 	prcp.Stm(stm)
// 			// 	prcpStm.Stm = stm

// 			// 	if isFstStgy { // find fst period stgy
// 			// 		// qckStm := instr.I(tme.M5).Ask().Sma()
// 			// 		// slwStm := instr.I(tme.M20).Ask().Sma()
// 			// 		// opnCnd := qckStm.OtrLss(0, slwStm).Seq(tme.S1, qckStm.OtrGtr(0, slwStm))
// 			// 		// MACD CRS
// 			// 		h1AskLst := instr.I(tme.H1).Ask().Lst()
// 			// 		qckStm := h1AskLst.AggEma(12).OtrSub(0, h1AskLst.AggEma(26))
// 			// 		slwStm := qckStm.AggEma(9)
// 			// 		opnCnd := slwStm.SclGtr(0.0).And(qckStm.OtrLss(0, slwStm).Seq(tme.S1, qckStm.OtrGtr(0, slwStm)))

// 			// 		for _, prfLim := range *prfLims {
// 			// 			for _, losLim := range *losLims {
// 			// 				for _, durLim := range *durLims {
// 			// 					stgy := opnCnd.Long(prfLim, losLim, durLim, instr)
// 			// 					// sys.Log("stgy:", stgy)
// 			// 					prfm := stgy.Port().Prfm()
// 			// 					splt := prfm.Port().Splt(0.0)
// 			// 					tunedPrfm := prcp.Splt(splt).TuneSacfTil(
// 			// 						0.0, stmCntLim, trimItrLim, trimMin, trimForgiveLim,
// 			// 					)
// 			// 					mu.Lock()
// 			// 					bckPfmsA.Push(tunedPrfm)
// 			// 					mu.Unlock()
// 			// 					// sys.Log("tunedPrfm.Ana().PnlPct", tunedPrfm.Ana().PnlPct)
// 			// 					sys.Logf("tunedPrfm bckIdx:%v bckRng:%v trds:%v", bckIdx, bckRng, tunedPrfm.Port().Stgys().At(0).Bse().Trds)
// 			// 				}
// 			// 			}
// 			// 		}
// 			// 		isFstStgy = false
// 			// 	} else {
// 			// 		for n := 0; n < len(*bckPfmsA); n++ {
// 			// 			prfmA := (*bckPfmsA)[n]
// 			// 			splt := prfmA.Port().Rng(bckRng).Splt(0.0)
// 			// 			sys.Logf("idx:%v bckRng:%v prfmA.trds:%v", bckIdx, bckRng, prfmA.Port().Stgys().At(0).Bse().Trds.Cnt())
// 			// 			tunedPrfmB := prcp.Splt(splt).TuneSacfTil(
// 			// 				0.0, stmCntLim, trimItrLim, trimMin, trimForgiveLim,
// 			// 			)
// 			// 			mu.Lock()
// 			// 			bckPfmsB.Push(tunedPrfmB)
// 			// 			mu.Unlock()
// 			// 			sys.Logf("idx:%v tunedPrfmB.trds:%v", bckIdx, tunedPrfmB.Port().Stgys().At(0).Bse().Trds.Cnt())
// 			// 			// sys.Logf("tunedPrfmB bckIdx:%v trds:%v", bckIdx, tunedPrfmB.Port().Stgys().At(0).Bse().Trds)
// 			// 		}
// 			// 		tmpPrfms := hst.NewPrfms().Mrg(bckPfmsB)
// 			// 		bckPfmsB = bckPfmsA
// 			// 		bckPfmsB.Clr()
// 			// 		bckPfmsA = tmpPrfms
// 			// 	}
// 			// } // end bck training

// 			// // fwd test
// 			// sys.Log(" ### bckRng ###", mktDays.RngMrg(unt.Unt(dayIdx), unt.Unt(fwdIdx-1)))
// 			// fwdRng := (*mktDays)[fwdIdx]
// 			// sys.Log(" ### fwdRng ###", fwdRng, len(*bckPfmsA))
// 			// if len(*bckPfmsA) == 0 {
// 			// 	// fwdPnlPcts.Push(0)
// 			// 	prcpDur.FwdPnlPcts.Push(0)
// 			// 	prcpStm.PnlPct = flt.Min
// 			// 	sys.Log("NONE")
// 			// } else {
// 			// 	bckPrfm := bckPfmsA.SrtDscPnlPct().Fst()
// 			// 	sys.Log("bckPrfm", bckPrfm.Ana())
// 			// 	fwdPrfm := bckPrfm.Port().Rng(fwdRng).Prfm()
// 			// 	sys.Log("fwdPrfm", fwdPrfm.Ana())
// 			// 	prcpDur.FwdPnlPcts.Push(fwdPrfm.Ana().PnlPct)
// 			// 	prcpStm.FwdPrfm = fwdPrfm
// 			// 	prcpStm.PnlPct = fwdPrfm.Ana().PnlPct
// 			// }

// 			sys.Logf("TRN BCK/FWD: DONE:%v", time.Now().Sub(trnStart))
// 		} // end bck-fwd trn

// 		sys.Logf("DONE:%v DUR:%v FwdPnlPcts:%v", tme.Duration(time.Now().Sub(durStart)), prcpDur.Dur, prcpDur.FwdPnlPcts)
// 	} // end dur
// 	Permute(prcpStms)

// 	// sys.Logf(" --- fwdPnlPcts:%v", fwdPnlPcts)
// 	sys.Logf("SESSION: DONE:%v", tme.Duration(time.Now().Sub(sesStart)))
// } // end srch

// func Permute(prcpStms PrcpStms) {
// 	permuteStart := time.Now()
// 	sys.Log("Permute", len(prcpStms), permuteStart)
// 	filter := prcpStms.FilterPnlPctGtr(0)
// 	sys.Log("Permute: Filtered", len(filter))
// 	grps := filter.GrpDur()
// 	sys.Log(grps.String())
// 	sys.Logf("PERMUTE: DONE:%v", tme.Duration(time.Now().Sub(permuteStart)))
// }

// func (x *PrcpStms) Permute() {

// }
// func (x *PrcpStms) SortAsc() {
// 	sort.Slice(*x, func(i, j int) bool { return (*x)[i].PnlPct < (*x)[j].PnlPct })
// }
// func (x *PrcpStms) SortDsc() {
// 	sort.Slice(*x, func(i, j int) bool { return (*x)[i].PnlPct > (*x)[j].PnlPct })
// }
// func (x *PrcpStms) FilterPnlPctGtr(pnlPct flt.Flt) PrcpStms {
// 	sys.Log("PrcpStms.FilterPnlPctGtr: start:", len(*x))
// 	x.SortDsc()
// 	var n int
// 	for ; n < len(*x); n++ {
// 		if (*x)[n].PnlPct < 0 {
// 			break
// 		}
// 	}
// 	return (*x)[:n]
// }
// func (x *PrcpStms) GrpDur() (r PrcpStmDurGrps) {
// 	sys.Log("PrcpStms.GrpDur: start:", len(*x))
// 	r = make(PrcpStmDurGrps)
// 	for _, prcpStm := range *x {
// 		grp, ok := r[prcpStm.Dur]
// 		if !ok {
// 			grp = &PrcpStmDurGrp{Dur: prcpStm.Dur}
// 			r[prcpStm.Dur] = grp
// 		}
// 		grp.PrcpStms = append(grp.PrcpStms, prcpStm)
// 	}
// 	sys.Log("grps", len(r))
// 	return r
// }

// func (x *PrcpStmDurGrps) SortAsc() {
// 	for _, v := range *x {
// 		v.PrcpStms.SortAsc()
// 	}
// }
// func (x *PrcpStmDurGrps) SortDsc() {
// 	for _, v := range *x {
// 		v.PrcpStms.SortDsc()
// 	}
// }
// func (x *PrcpStmDurGrps) Fsts() (r PrcpStms) {
// 	for _, v := range *x {
// 		r = append(r, v.PrcpStms[0])
// 	}
// 	return r
// }
// func (x *PrcpStmDurGrps) String() string {
// 	var sb strings.Builder
// 	sb.WriteString("[")
// 	for _, v := range *x {
// 		sb.WriteString("\n")
// 		sb.WriteString(v.Dur.String())
// 		for n, prcpStm := range v.PrcpStms {
// 			sb.WriteString("\n")
// 			sb.WriteString(fmt.Sprintf("%v", n+1))
// 			sb.WriteString("\t")
// 			sb.WriteString(prcpStm.PnlPct.String())
// 			sb.WriteString("\t")
// 			sb.WriteString(prcpStm.FwdPrfm.Ana().String())
// 			sb.WriteString("\t")
// 			sb.WriteString(prcpStm.Stm.String())
// 		}
// 	}
// 	sb.WriteString("\n]")
// 	return sb.String()
// }
