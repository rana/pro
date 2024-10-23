package pro_test

// func TestScriptLive(t *testing.T) {
// 	cfg := cfg.Load("/home/rana/go/src/sys/cmd/sys.cfg")
// 	cfg.Ui = false
// 	ap := app.New(cfg)
// 	defer ap.Cls()

// 	txt := `
// tme.now().asn(start)
// log.ifo("START")

// [1m	2m 3m 5m 8m 13m	21m	34m	55m	89m	144m 233m 377m].asn(durs) // 1h29m|2h24m|3h53m|6h17m|10h10m
// flts.addsLeq(4.0 4.0 5.0).asn(prfLims)
// flts.addsLeq(4.0 4.0 5.0).asn(losLims)
// tmes.addsLeq(4m 4m 15m).asn(durLims)

// hst.oan().eurUsd(0s-0s).mktWeeks().asn(mktWeeks)
// hst.oan().eurUsd(0s-0s).mktDays().asn(mktDays)
// hst.oan().eurUsd(0s-0s).mktHrs().asn(mktHrs)
// mktDays.asn(mktRngs).cnt()

// 2.asn(stmCntLim)
// 2.asn(bckTrnLen)
// 3.asn(trimItrLim)
// 4.asn(trimMin)
// 1.asn(trimForgiveLim)
// // log.ifof(
// // 	"stmCntLim:%v bckTrnLen:%v trimItrLim:%v trimMin:%v trimForgiveLim:%v"
// // 	stmCntLim bckTrnLen trimItrLim trimMin trimForgiveLim
// // )

// durs.from(0).to(1).asn(durs)
// mktRngs.from(0).to(bckTrnLen.add(2)).asn(mktRngs)
// durs.each(dur
// 	// log.ifo("--- dur" dur)
// 	flts.new().asn(fwdPnlPcts)
// 	unts.addsLss(0 mktRngs.lstIdx().sub(bckTrnLen) 1).asn(idxs)
// 	idxs.each(idx
// 		tru.asn(isFstStgy)
// 		sys.newMu().asn(prfmsMu)
// 		hst.newPrfms().asn(bckPfmsA)
// 		hst.newPrfms().asn(bckPfmsB)
// 		mktRngs.cnt()
// 		mktRngs.at(idx.add(bckTrnLen)).asn(fwdRng)
// 		mktRngs.from(idx).to(bckTrnLen).each(bckRng
// 			// log.ifof("bckRng:%v bckPfmsACnt:%v" bckRng bckPfmsA.cnt())

// 			hst.oan().eurUsd(bckRng).asn(instr)
// 			instr.i(1m).ask().alma(6 0.85).asn(m1Alma)
// 			instr.i(15m).ask().sma().asn(m15Sma)
// 			instr.i(dur).ask().alma(6 0.85).asn(alma)
// 			instr.i(dur).ask().std().asn(stm)
// 			hst.newPrcp().asn(prcp)
// 			prcp.stm(stm)

// 			isFstStgy.then( // find fst period stgy
// 				m1Alma.otrLss(0 m15Sma).seq(1s m1Alma.otrGtr(0 m15Sma)).asn(opnCnd)
// 				// log.ifo(" ~ TUNING CND" opnCnd)
// 				opnCnd.longs(prfLims losLims durLims [instr]).stgys().rev().asn(stgys)
// 				stgys.pllEach(stgy
// 					stgy.port().prfm().asn(prfm)
// 					prfm.port().splt(0.0).asn(splt)
// 					// prcp.splt(splt).tuneTil(0.0 stmCntLim).asn(tunedPrfm)
// 					prcp.splt(splt).tuneSacfTil(0.0 stmCntLim
// 						trimItrLim trimMin trimForgiveLim
// 					).asn(tunedPrfm)
// 					// log.ifo("tunedPrfm.ana().pnlPct()" tunedPrfm.ana().pnlPct())
// 					prfmsMu.lck()
// 					bckPfmsA.push(tunedPrfm)
// 					prfmsMu.ulck()

// 					plt.newStgy().asn(pltStgy0)
// 					pltStgy0.stgy(stgy)
// 					pltStgy0.stm(stm)
// 					// pltStgy0.title(str.fmt("stm %v %v" idx dur))
// 					pltStgy0.sho()
// 					plt.newVrt().asn(pltVrt)
// 					pltVrt.plt(plt.newStgy().asn(pltStgy1))
// 					pltStgy1.stgy(tunedPrfm.port().stgys().at(0))
// 					pltStgy1.stm(stm)
// 					// pltStgy1.title(str.fmt("stm tuned %v %v" idx dur))
// 					pltVrt.plt(plt.newStgy().asn(pltStgy2))
// 					pltStgy2.stgy(tunedPrfm.port().stgys().at(0))
// 					pltStgy2.stm(alma)
// 					// pltStgy2.title(str.fmt("alma tuned %v %v" idx dur))
// 					pltVrt.sho()
// 				)
// 				fls.asn(isFstStgy 2)
// 			).else( // tune existing stgy
// 				bckPfmsA.pllEach(prfmA
// 					prfmA.port().splt(0.0).asn(splt)
// 					// prcp.splt(splt).tuneTil(0.0 stmCntLim).asn(tunedPrfmB)
// 					prcp.splt(splt).tuneSacfTil(0.0 stmCntLim
// 						trimItrLim trimMin trimForgiveLim
// 					).asn(tunedPrfmB)
// 					prfmsMu.lck()
// 					bckPfmsB.push(tunedPrfmB)
// 					prfmsMu.ulck()
// 				)
// 				bckPfmsA.asn(prfmsTmp)
// 				bckPfmsB.asn(bckPfmsA 2)
// 				prfmsTmp.clr()
// 				prfmsTmp.asn(bckPfmsB 2)
// 			)
// 		) // end training

// 		// log.ifo("### bck rng ###" mktRngs.rngMrg(idx idx.add(bckTrnLen).sub(1)))
// 		// log.ifo("### fwd rng ###" fwdRng)
// 		bckPfmsA.cnt().eql(0).then(
// 			// log.ifo("NONE")
// 			fwdPnlPcts.push(0.0)
// 		).else(
// 			bckPfmsA.srtDscPnlPct().fst().asn(bckPrfm)
// 			bckPrfm.port().rng(fwdRng).prfm().asn(fwdPrfm)
// 			// log.ifo("BCK" bckPrfm)
// 			// log.ifo("FWD" fwdPrfm)
// 			fwdPnlPcts.push(fwdPrfm.ana().pnlPct())

// 			// hst.oan().eurUsd(fwdRng).asn(instr)
// 			// instr.i(dur).ask().alma(6 0.85).asn(alma)
// 			// instr.i(dur).ask().std().asn(stm)
// 			// plt.newStgy().asn(pltStgy1)
// 			// pltStgy1.stgy(fwdPrfm.port().stgys().at(0))
// 			// pltStgy1.stm(stm)
// 			// // pltStgy1.title(str.fmt("FWD STM %v %v" idx dur))
// 			// pltStgy1.sho()
// 			// plt.newStgy().asn(pltStgy2)
// 			// pltStgy2.stgy(fwdPrfm.port().stgys().at(0))
// 			// pltStgy2.stm(alma)
// 			// // pltStgy2.title(str.fmt("FWD ALMA %v %v" idx dur))
// 			// pltStgy2.sho()
// 		)
// 	)
// 	// log.ifo("fwdPnlPcts" fwdPnlPcts)
// 	// log.ifo("std" dur)
// 	// log.ifo("*** pnlPct total" fwdPnlPcts.sum().trnc(1))
// )

// // log.ifo("DONE" tme.now().sub(start))

// 	`
// 	act.Run(txt)

// 	// ap.Run() // for Rlt.OpnScript

// 	<-time.NewTimer(time.Hour).C
// }
