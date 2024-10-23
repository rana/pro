package srch

// type (
// 	SrchPrcpReq0 struct {
// 		Instr  Instr
// 		Sides  Sides
// 		OpnCnd OpnCnd
// 		Stgy   Stgy
// 		PrfLim flt.Flt
// 		LosLim flt.Flt
// 		DurLim tme.Tme

// 		DurSideStms  []DurSideStm
// 		SrchPrcpRess []SrchPrcpRes0
// 	}
// 	SrchPrcpRes0 struct {
// 		Name       string
// 		Prcp       *hst.Prcp
// 		BckSum     flt.Flt
// 		FwdSum     flt.Flt
// 		DltSum     flt.Flt
// 		BckPnlPcts *flts.Flts
// 		FwdPnlPcts *flts.Flts
// 	}
// )

// func (x SrchPrcpRes0) String() string {
// 	return fmt.Sprintf("%v BckSum:%v FwdSum:%v DltSum:%v FwdPnlPcts:%v", x.Name, x.BckSum, x.FwdSum, x.DltSum, x.FwdPnlPcts)
// }

// func SrchPerms0() {
// 	prfLims := flts.AddsLeq(10.0, 10.0, 5.0)
// 	losLims := flts.AddsLeq(10.0, 10.0, 5.0)
// 	durLims := tmes.AddsLeq(tme.M15, tme.M15, tme.M15)
// 	stmCntLim := unt.Unt(2)
// 	trimItrLim := unt.Unt(3)
// 	trimMin := unt.Unt(4)
// 	trimForgiveLim := unt.Unt(1)
// 	tstRngs := hst.Oan().EurUsd().Bse().Ana.MktDays

// 	start := time.Now()
// 	sys.Logf("%v: START SrchPrcp", start)
// 	// build srchPrcpReqs defs
// 	sys.Logf("%v: START build srchPrcpReqs defs", start)
// 	var srchPrcpReqs []*SrchPrcpReq0
// 	for _, prfLim := range *prfLims {
// 		for _, losLim := range *losLims {
// 			for _, durLim := range *durLims {
// 				srchPrcpReq := &SrchPrcpReq0{}
// 				srchPrcpReqs = append(srchPrcpReqs, srchPrcpReq)
// 				srchPrcpReq.Instr = EurUsd
// 				srchPrcpReq.Sides = DurAsks
// 				srchPrcpReq.OpnCnd = OpnCndAlmaSmaCrsUp520
// 				srchPrcpReq.Stgy = Long
// 				srchPrcpReq.PrfLim = prfLim
// 				srchPrcpReq.LosLim = losLim
// 				srchPrcpReq.DurLim = durLim
// 			}
// 		}
// 	}
// 	sys.Logf("%v: DONE build srchPrcpReqs defs", tme.Duration(time.Now().Sub(start)))

// 	// run srchPrcpReqs
// 	sys.Logf("%v: START run %v srchPrcpReqs", len(srchPrcpReqs), time.Now())
// 	for m, s := range srchPrcpReqs {
// 		sys.Logf("--- srchPrcpReq: %v of %v", m+1, len(srchPrcpReqs))
// 		n := len(*Durs) * len(VSideStms)
// 		k := 2
// 		p := NewPerm(n, k)
// 		sys.Logf("GEN PERMS DEFS n:%v k:%v NumPerms:%v", n, k, p.Cnt)
// 		startPerms := time.Now()
// 		s.DurSideStms = make([]DurSideStm, p.Cnt)
// 		var idx int
// 		for _, dur := range *Durs {
// 			for _, sideStm := range VSideStms {
// 				s.DurSideStms[idx].Dur = dur
// 				s.DurSideStms[idx].SideStm = sideStm
// 				idx++
// 			}
// 		}
// 		sys.Logf("%v: DONE: GEN PERMS DEFS", tme.Duration(time.Now().Sub(startPerms)))
// 		s.SrchPrcpRess = make([]SrchPrcpRes0, len(s.DurSideStms))
// 		for n := 0; n < len(s.SrchPrcpRess); n++ {
// 			s.SrchPrcpRess[n].BckPnlPcts = flts.New()
// 			s.SrchPrcpRess[n].FwdPnlPcts = flts.New()
// 		}

// 		for n := unt.Zero; n < tstRngs.LstIdx(); n++ { // fwd test one day
// 			p.Reset() // reset permutation for bck tst

// 			// bck tst
// 			bckRng := tstRngs.At(n)
// 			sys.Logf("bckRng %v", bckRng)

// 			instr := s.Instr(bckRng)
// 			opnCnd := s.OpnCnd(instr)
// 			stgy := s.Stgy(opnCnd, s.PrfLim, s.LosLim, s.DurLim, instr)
// 			bckPrfm := stgy.Port().Prfm()
// 			splt := bckPrfm.Port().Splt(0.0)

// 			sides := s.Sides(instr, Durs)
// 			durSides := make(map[tme.Tme]hst.Side)
// 			for n, dur := range *Durs {
// 				durSides[dur] = (*sides)[n]
// 			}
// 			var sb strings.Builder
// 			for p.Next() {
// 				stms := hst.NewStms()
// 				for _, idx := range p.Perm() {
// 					durSideStm := s.DurSideStms[idx]
// 					stm := durSideStm.SideStm(durSides[durSideStm.Dur])
// 					stms.Push(stm)
// 				}
// 				prcp := hst.NewPrcp(*stms...)
// 				bckPrfm = prcp.Splt(splt).TuneSacfTil(0.0, stmCntLim, trimItrLim, trimMin, trimForgiveLim)

// 				sb.Reset()
// 				for n, stm := range *stms {
// 					if n != 0 {
// 						sb.WriteRune('-')
// 					}
// 					sb.WriteString(stm.Name().Unquo())
// 				}
// 				s.SrchPrcpRess[p.Idx].Name = sb.String()
// 				s.SrchPrcpRess[p.Idx].BckPnlPcts.Push(bckPrfm.Ana().PnlPct)
// 				sys.Logf("%v BckPnlPct:%v %v", p.Prgrs(), s.SrchPrcpRess[p.Idx].BckPnlPcts.Lst(), s.SrchPrcpRess[p.Idx])

// 				// fwd tst
// 				fwdRng := tstRngs.At(n + 1)
// 				// prcp.Sub()
// 				fwdPrfm := bckPrfm.Port().Rng(fwdRng).Prfm()
// 				// prcp.Unsub()
// 				s.SrchPrcpRess[p.Idx].FwdPnlPcts.Push(fwdPrfm.Ana().PnlPct)
// 				sys.Logf("%v FwdPnlPct:%v %v", p.Prgrs(), s.SrchPrcpRess[p.Idx].FwdPnlPcts.Lst(), s.SrchPrcpRess[p.Idx])
// 			}
// 		}
// 		for n := 0; n < len(s.SrchPrcpRess); n++ {
// 			s.SrchPrcpRess[n].BckSum = s.SrchPrcpRess[n].BckPnlPcts.Sum()
// 			s.SrchPrcpRess[n].FwdSum = s.SrchPrcpRess[n].FwdPnlPcts.Sum()
// 			s.SrchPrcpRess[n].DltSum = s.SrchPrcpRess[n].FwdSum - s.SrchPrcpRess[n].BckSum
// 			//sys.Logf("%v FwdSum:%v BckSum:%v DltSum:%v", s.Req, s.FwdSum, s.BckSum, s.DltSum)
// 		}

// 		sys.Logf("%v: START srt srchPrcpFams", time.Now())
// 		sort.Slice(s.SrchPrcpRess, func(i, j int) bool {
// 			return s.SrchPrcpRess[i].BckSum > s.SrchPrcpRess[j].BckSum
// 		})
// 		sys.Logf("TOP BckSum %v", len(s.SrchPrcpRess))
// 		for n := 0; n < len(s.SrchPrcpRess); n++ {
// 			sys.Log(s.SrchPrcpRess[n])
// 		}

// 		sort.Slice(s.SrchPrcpRess, func(i, j int) bool {
// 			return s.SrchPrcpRess[i].FwdSum > s.SrchPrcpRess[j].FwdSum
// 		})
// 		sys.Logf("TOP FwdSum %v", len(s.SrchPrcpRess))
// 		for n := 0; n < len(s.SrchPrcpRess); n++ {
// 			sys.Log(s.SrchPrcpRess[n])
// 		}

// 		sort.Slice(s.SrchPrcpRess, func(i, j int) bool {
// 			return s.SrchPrcpRess[i].DltSum > s.SrchPrcpRess[j].DltSum
// 		})
// 		sys.Logf("TOP DltSum %v", len(s.SrchPrcpRess))
// 		for n := 0; n < len(s.SrchPrcpRess); n++ {
// 			sys.Log(s.SrchPrcpRess[n])
// 		}

// 	}
// 	sys.Logf("%v: DONE run srchPrcpReqs", tme.Duration(time.Now().Sub(start)))

// 	sys.Logf("%v: DONE", tme.Duration(time.Now().Sub(start)))
// }
