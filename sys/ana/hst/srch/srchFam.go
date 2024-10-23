package srch

// type (
// 	SrchPrcpFam struct {
// 		Fam        str.Str
// 		BckSum     flt.Flt
// 		FwdSum     flt.Flt
// 		DltSum     flt.Flt
// 		BckPnlPcts *flts.Flts
// 		FwdPnlPcts *flts.Flts
// 		TstRngs    *tme.Rngs
// 		Durs       *tmes.Tmes

// 		Instr    Instr
// 		Sides    Sides
// 		SideStms SideStms
// 		OpnCnd   OpnCnd
// 		Stgy     Stgy
// 		PrfLim   flt.Flt
// 		LosLim   flt.Flt
// 		DurLim   tme.Tme
// 	}
// )

// func (x *SrchPrcpFam) String() string {
// 	return fmt.Sprintf("%v BckSum:%v FwdSum:%v DltSum:%v FwdPnlPcts:%v", x.Fam, x.BckSum, x.FwdSum, x.DltSum, x.FwdPnlPcts)
// }

// func SrchFam() {
// 	prfLims := flts.AddsLeq(10.0, 10.0, 5.0)
// 	losLims := flts.AddsLeq(10.0, 10.0, 5.0)
// 	durLims := tmes.AddsLeq(tme.M15, tme.M15, tme.M15)
// 	stmCntLim := unt.Unt(2)
// 	trimItrLim := unt.Unt(3)
// 	trimMin := unt.Unt(4)
// 	trimForgiveLim := unt.Unt(1)
// 	start := time.Now()

// 	// build srchPrcpFams defs
// 	sys.Logf("%v: START build srchPrcpFams defs", start)
// 	var srchPrcpFams []*SrchPrcpFam
// 	for _, sideStms := range SideStmss {
// 		for _, prfLim := range *prfLims {
// 			for _, losLim := range *losLims {
// 				for _, durLim := range *durLims {
// 					srchPrcpFam := &SrchPrcpFam{}
// 					srchPrcpFams = append(srchPrcpFams, srchPrcpFam)
// 					srchPrcpFam.BckPnlPcts = flts.New()
// 					srchPrcpFam.FwdPnlPcts = flts.New()
// 					srchPrcpFam.TstRngs = hst.Oan().EurUsd().Bse().Ana.MktDays
// 					srchPrcpFam.Durs = Durs
// 					srchPrcpFam.Instr = EurUsd
// 					srchPrcpFam.Sides = DurAsks
// 					srchPrcpFam.SideStms = sideStms
// 					srchPrcpFam.OpnCnd = OpnCndMacdCrsLong
// 					srchPrcpFam.Stgy = Long
// 					srchPrcpFam.PrfLim = prfLim
// 					srchPrcpFam.LosLim = losLim
// 					srchPrcpFam.DurLim = durLim
// 				}
// 			}
// 		}
// 	}
// 	sys.Logf("%v: DONE build srchPrcpFams defs", tme.Duration(time.Now().Sub(start)))

// 	// run srchPrcpFams
// 	sys.Logf("%v: START run %v srchPrcpFams", len(srchPrcpFams), time.Now())
// 	for m, s := range srchPrcpFams {
// 		sys.Logf("--- srchPrcpFam: %v of %v", m+1, len(srchPrcpFams))
// 		for n := unt.Zero; n < s.TstRngs.LstIdx(); n++ { // fwd test one day
// 			// bck tst
// 			bckRng := s.TstRngs.At(n)
// 			sys.Logf("bckRng %v", bckRng)

// 			instr := s.Instr(bckRng)
// 			sides := s.Sides(instr, s.Durs)
// 			stms := s.SideStms(sides)
// 			if s.Fam == "" {
// 				s.Fam = stms.Fst().Name()
// 			}
// 			opnCnd := s.OpnCnd(instr)
// 			stgy := s.Stgy(opnCnd, s.PrfLim, s.LosLim, s.DurLim, instr)
// 			bckPrfm := stgy.Port().Prfm()
// 			splt := bckPrfm.Port().Splt(0.0)
// 			prcp := hst.NewPrcp(*stms...)
// 			bckPrfm = prcp.Splt(splt).TuneSacfTil(0.0, stmCntLim, trimItrLim, trimMin, trimForgiveLim)
// 			s.BckPnlPcts.Push(bckPrfm.Ana().PnlPct)
// 			sys.Logf("%v BckPnlPct:%v", s.Fam, s.BckPnlPcts.Lst())
// 			// fwd tst
// 			fwdRng := s.TstRngs.At(n + 1)
// 			// prcp.Sub()
// 			fwdPrfm := bckPrfm.Port().Rng(fwdRng).Prfm()
// 			// prcp.Unsub()
// 			s.FwdPnlPcts.Push(fwdPrfm.Ana().PnlPct)
// 			sys.Logf("%v FwdPnlPct:%v", s.Fam, s.FwdPnlPcts.Lst())
// 		}
// 		s.BckSum = s.BckPnlPcts.Sum()
// 		s.FwdSum = s.FwdPnlPcts.Sum()
// 		s.DltSum = s.FwdSum - s.BckSum
// 		sys.Logf("%v FwdSum:%v BckSum:%v DltSum:%v", s.Fam, s.FwdSum, s.BckSum, s.DltSum)
// 	}
// 	sys.Logf("%v: DONE run srchPrcpFams", tme.Duration(time.Now().Sub(start)))

// 	// srt srchPrcpFams
// 	top := len(srchPrcpFams)
// 	sys.Logf("%v: START srt srchPrcpFams", time.Now())
// 	sort.Slice(srchPrcpFams, func(i, j int) bool {
// 		return srchPrcpFams[i].BckSum > srchPrcpFams[j].BckSum
// 	})
// 	sys.Logf("TOP BckSum %v", top)
// 	for n := 0; n < top; n++ {
// 		sys.Log(srchPrcpFams[n])
// 	}

// 	sort.Slice(srchPrcpFams, func(i, j int) bool {
// 		return srchPrcpFams[i].FwdSum > srchPrcpFams[j].FwdSum
// 	})
// 	sys.Logf("TOP FwdSum %v", top)
// 	for n := 0; n < top; n++ {
// 		sys.Log(srchPrcpFams[n])
// 	}

// 	sort.Slice(srchPrcpFams, func(i, j int) bool {
// 		return srchPrcpFams[i].DltSum > srchPrcpFams[j].DltSum
// 	})
// 	sys.Logf("TOP DltSum %v", top)
// 	for n := 0; n < top; n++ {
// 		sys.Log(srchPrcpFams[n])
// 	}

// 	sys.Logf("%v: DONE", tme.Duration(time.Now().Sub(start)))
// }
