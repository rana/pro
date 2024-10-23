package hst

// func (x *Port) CalcTrds() {
// 	x.Trds = ana.NewTrds()
// 	x.BalFstUsd = ana.Cfg.Hst.BalUsd
// 	x.BalLstUsd = ana.Cfg.Hst.BalUsd
// 	x.TrdPct = ana.Cfg.Hst.TrdPct
// 	if len(*x.Stgys) == 0 {
// 		return
// 	}
// 	stgys := make([]*StgyBse, len(*x.Stgys))
// 	cndTmess := make([]*tmes.Tmes, len(*x.Stgys))
// 	for n, stgy := range *x.Stgys {
// 		stgys[n] = stgy.Bse()
// 		stgys[n].Trds = ana.NewTrds()
// 		cndTmess[n] = (*x.Stgys)[n].(ICnd).CndGet().Bse().Tmes
// 	}

// 	nxtAvailOpnTme := tme.Min
// 	var stgy *StgyBse
// 	var opnTme, prvOpnTme tme.Tme
// 	for {
// 		opnTme = tme.Max
// 		for n, cndTmes := range cndTmess { // find first open time from any stgy
// 			idx := cndTmes.SrchIdxEql(nxtAvailOpnTme)
// 			if int(idx) < len(*cndTmes) && (*cndTmes)[idx] < opnTme {
// 				opnTme = (*cndTmes)[idx]
// 				stgy = stgys[n]
// 			}
// 		}
// 		// sys.Logf("CalcTrds stgy %p %v", stgy, opnTme)
// 		if opnTme == prvOpnTme {
// 			// sys.Logf("CalcTrds stgy CONTINUING %v", opnTme == prvOpnTme)
// 			nxtAvailOpnTme += tme.Resolution
// 			continue
// 		}
// 		if opnTme == tme.Max {
// 			break
// 		}
// 		trd, rsnFail := stgy.OpnClsTrd(opnTme, x.BalLstUsd, x.TrdPct)
// 		if trd != nil { // may fail to close due to near mkt opn, near mkt cls, spd lim exceeded
// 			x.CalcOpn(trd, stgy.Instr.Instr()) // set trd flds
// 			x.CalcCls(trd, stgy.Instr.Instr()) // set trd flds
// 			stgy.Trds.Push(trd)
// 			x.Trds.Push(trd)
// 			nxtAvailOpnTme = trd.ClsTme // + tme.Resolution // RLT DOES NOT OPN AND CLS TRD ON THE SAME SECOND
// 			// sys.Logf("HST %v", trd)
// 		} else {
// 			if rsnFail == ana.NoCls {
// 				break // exit last opn fail to mirror rlt behavior; single ana.NoCls expected
// 			}
// 			// trd failed to cls; advance to next open time
// 			nxtAvailOpnTme = opnTme // + tme.Resolution // RLT DOES NOT OPN AND CLS TRD ON THE SAME SECOND
// 		}
// 		prvOpnTme = opnTme
// 	}
// }
