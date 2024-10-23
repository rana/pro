package srch

// type (
// 	SrchPrcpReq1 struct {
// 		Instr  Instr
// 		Side   Side
// 		OpnCnd OpnCnd
// 		Stgy   Stgy
// 		PrfLim flt.Flt
// 		LosLim flt.Flt
// 		DurLim tme.Tme
// 	}
// 	SrchPrcp struct {
// 		Durs      *tmes.Tmes
// 		SideStms  []SideStm
// 		PermIdxss [][]uint32
// 		Ress      []*SrchPrcpRes
// 		Itr       func(itrIdx uint32, idxs []uint32)
// 	}
// 	SrchPrcpSeg struct {
// 		bnd.Bnd
// 		X *SrchPrcp
// 	}
// 	SrchPrcpRes struct {
// 		Name       string
// 		BckSum     flt.Flt
// 		FwdSum     flt.Flt
// 		DltSum     flt.Flt
// 		BckPnlPcts *flts.Flts
// 		FwdPnlPcts *flts.Flts
// 		PermIdxs   []uint32
// 	}
// 	SrchPrcpRess []*SrchPrcpRes
// )

// func SrchPerms1() {
// 	start := time.Now()

// 	tstRngs := hst.Oan().EurUsd().Bse().Ana.MktDays
// 	bckTrnLen := 2 // min 1
// 	prfLim := flt.Flt(10.0)
// 	losLim := flt.Flt(10.0)
// 	durLim := tme.M15
// 	spltPnt := flt.Zero
// 	stmCntLim := unt.Unt(2)
// 	trimItrLim := unt.Unt(3)
// 	trimMin := unt.Unt(4)
// 	trimForgiveLim := unt.Unt(1)
// 	Instr := EurUsd
// 	Side := Ask
// 	OpnCnd := OpnCndAlmaSmaCrsUp520
// 	Stgy := Long

// 	srchPrcp := NewSrchPrcp(Durs, VSideStms, 2)
// 	srchPrcp.Itr = func(itrIdx uint32, permIdxs []uint32) {
// 		srchPrcp.Ress[itrIdx] = &SrchPrcpRes{
// 			BckPnlPcts: flts.New(),
// 			FwdPnlPcts: flts.New(),
// 			PermIdxs:   permIdxs,
// 		}
// 		var prcp *hst.Prcp // for name

// 		// sys.Logf("srchPrcp.Itr len(*tstRngs):%v bckTrnLen:%v", len(*tstRngs), bckTrnLen)
// 		for rngIdx := 0; rngIdx < len(*tstRngs)-bckTrnLen; rngIdx++ { // fwd test one day
// 			// sys.Logf("itrIdx:%v rngIdx:%v bckRng:%v", itrIdx, rngIdx, (*tstRngs)[rngIdx])

// 			// BCK STGY & TUNE
// 			instr := Instr((*tstRngs)[rngIdx])

// 			opnCnd := OpnCnd(instr)
// 			stgy := Stgy(opnCnd, prfLim, losLim, durLim, instr)
// 			bckPrfm := stgy.Port().Prfm()
// 			// sys.Log("bckPrfm a.0", bckPrfm.Ana())
// 			splt := bckPrfm.Port().Splt(spltPnt)
// 			stms := hst.NewStms()
// 			for _, idx := range permIdxs {
// 				durSideStm := srchPrcp.DurSideStm(idx)
// 				stm := durSideStm.SideStm(Side(instr, durSideStm.Dur))
// 				stms.Push(stm)
// 			}
// 			prcp = hst.NewPrcp(*stms...)
// 			bckPrfm = prcp.Splt(splt).TuneSacfTil(spltPnt, stmCntLim, trimItrLim, trimMin, trimForgiveLim)
// 			// sys.Log("bckPrfm a.1", bckPrfm.Ana())

// 			// BCK TUNE
// 			for dayTunIdx := rngIdx + 1; dayTunIdx < rngIdx+bckTrnLen; dayTunIdx++ {
// 				bckRngTune := (*tstRngs)[dayTunIdx]
// 				// sys.Log("bckRngTune", bckRngTune)
// 				// port := bckPrfm.Port().Rng(bckRngTune)
// 				// sys.Log("bckPrfm b.0", port.Prfm().Ana())
// 				splt := bckPrfm.Port().Rng(bckRngTune).Splt(spltPnt)
// 				// sys.Log("SPLT CMPLT", splt.Btm.Cnt(), splt.Top.Cnt())
// 				instr = Instr(bckRngTune)
// 				stms = hst.NewStms()
// 				for _, idx := range permIdxs {
// 					durSideStm := srchPrcp.DurSideStm(idx)
// 					stm := durSideStm.SideStm(Side(instr, durSideStm.Dur))
// 					stms.Push(stm)
// 				}
// 				prcp = hst.NewPrcp(*stms...)
// 				bckPrfm = prcp.Splt(splt).TuneSacfTil(spltPnt, stmCntLim, trimItrLim, trimMin, trimForgiveLim)
// 				// sys.Log("bckPrfm b.1", bckPrfm.Ana())
// 			}
// 			// panic("DONE")

// 			// FWD TST
// 			// sys.Log("fwdRng", (*tstRngs)[rngIdx+bckTrnLen])
// 			fwdPrfm := bckPrfm.Port().Rng((*tstRngs)[rngIdx+bckTrnLen]).Prfm()
// 			// sys.Log("fwdPrfm", fwdPrfm.Ana())

// 			srchPrcp.Ress[itrIdx].BckPnlPcts.Push(bckPrfm.Ana().PnlPct)
// 			srchPrcp.Ress[itrIdx].FwdPnlPcts.Push(fwdPrfm.Ana().PnlPct)
// 		} // end tstRngs
// 		srchPrcp.Ress[itrIdx].Name = prcp.Name() // use lst prcp to set name
// 		srchPrcp.Ress[itrIdx].Sum()
// 		sys.Log(" *** ", srchPrcp.Ress[itrIdx])
// 	} // end Itr

// 	segBnds, acts := bnds.Segs(unt.Unt(len(srchPrcp.PermIdxss)))
// 	srchPrcpSegs := make([]*SrchPrcpSeg, len(acts))
// 	for n, segBnd := range *segBnds {
// 		srchPrcpSegs[n] = &SrchPrcpSeg{
// 			Bnd: segBnd,
// 			X:   srchPrcp,
// 		}
// 		acts[n] = srchPrcpSegs[n]
// 	}

// 	sys.Logf("SRCH START perms:%v segBnds:%v %v", len(srchPrcp.PermIdxss), segBnds.Cnt(), start)

// 	// TODO: Pll(acts, 2)
// 	// TODO: PRGRS, ETA
// 	sys.Run().Seq(acts...)

// 	// TODO: SRT, PRNT

// 	sys.Logf("SRCH DONE ellapsed:%v %v", time.Now().Sub(start), time.Now())
// }

// func NewSrchPrcp(durs *tmes.Tmes, sideStms []SideStm, k int) (r *SrchPrcp) {
// 	r = &SrchPrcp{}
// 	r.Durs = durs
// 	r.SideStms = sideStms
// 	r.PermIdxss = AllPermIdxs(len(*r.Durs)*len(r.SideStms), k)
// 	r.Ress = make([]*SrchPrcpRes, len(r.PermIdxss))
// 	return r
// }

// func (x *SrchPrcp) DurSideStm(idx uint32) (r DurSideStm) {
// 	var n uint32
// 	for _, dur := range *x.Durs {
// 		for _, sideStm := range x.SideStms {
// 			if n == idx {
// 				r.Dur = dur
// 				r.SideStm = sideStm
// 				return r
// 			}
// 			n++
// 		}
// 	}
// 	return r
// }

// func (x *SrchPrcpSeg) Act() {
// 	for n := x.Idx; n < x.Lim; n++ {
// 		x.X.Itr(uint32(n), x.X.PermIdxss[n])
// 	}
// }

// func (x *SrchPrcpRes) Sum() {
// 	x.BckSum = x.BckPnlPcts.Sum()
// 	x.FwdSum = x.FwdPnlPcts.Sum()
// 	x.DltSum = x.FwdSum - x.BckSum
// }
// func (x *SrchPrcpRes) String() string {
// 	return fmt.Sprintf("%v BckSum:%v FwdSum:%v DltSum:%v \n BckPnlPcts:%v \n FwdPnlPcts:%v",
// 		x.Name, x.BckSum, x.FwdSum, x.DltSum, x.BckPnlPcts, x.FwdPnlPcts)
// }

// func (x SrchPrcpRess) SrtBckDsc() SrchPrcpRess {
// 	sort.Slice(x, func(i, j int) bool {
// 		return x[i].BckSum > x[j].BckSum
// 	})
// 	return x
// }
// func (x SrchPrcpRess) SrtFwdDsc() SrchPrcpRess {
// 	sort.Slice(x, func(i, j int) bool {
// 		return x[i].FwdSum > x[j].FwdSum
// 	})
// 	return x
// }
// func (x SrchPrcpRess) SrtDltDsc() SrchPrcpRess {
// 	sort.Slice(x, func(i, j int) bool {
// 		return x[i].DltSum > x[j].DltSum
// 	})
// 	return x
// }
// func (x SrchPrcpRess) Prnt() {
// 	for n := 0; n < len(x); n++ {
// 		sys.Log(x[n])
// 	}
// }
