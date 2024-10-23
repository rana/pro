tme.now().log("START:")
// plt.scl(1.6)


[1m 2m 4m 8m 16m 32m 64m 128m 256m 512m].asn(dursA)
[5m 10m 15m 20m 30m 50m 60m 100m 120m 200m 240m 400m 480m].asn(dursB)
[1m	2m 3m 5m 8m 13m	21m	34m	55m	89m	144m 233m 377m 610m].asn(dursC) // 1h29m|2h24m|3h53m|6h17m|10h10m
tmes.addsLss(5m 400m 5m).asn(dursE)
unts.fibsLeq(610).asn(offsA)
unts.fibsLeq(610).from(1).asn(lensA) // start at 2

dursC.asn(durs)
offsA.asn(offs)
lensA.asn(lens)
flts.addsLeq(10.0 40.0 5.0).asn(prfLims)
flts.addsLeq(10.0 40.0 5.0).asn(losLims)
// tmes.addsLeq(15m 4h 15m).asn(durLims)
tmes.addsLeq(4h 4h 15m).asn(durLims)

// tru.asn(quick)
// quick.then(
// 	"a".log()
// 	durs.cnt().log("durs")
// 	durs.to(3).asn(durs)
// 	"b".log()
// 	offs.to(3).asn(offs)
// 	"c".log()
// 	lens.to(3).asn(lens)
// 	"d".log()
// 	flts.addsLeq(30.0 15.0 1.0).asn(prfLims)

// 	flts.addsLeq(30.0 15.0 1.0).asn(losLims)
// 	tmes.addsLeq(4h 4h 15m).asn(durLims)
// )



hst.oan().eurUsd().mktWeeks().asn(mktWeeks)
hst.oan().eurUsd().mktDays().asn(mktDays)
mktDays.cnt().log("mkt day cnt")

1.asn(len)
unts.addsLss(0 mktWeeks.lstIdx() 1).asn(idxs)
idxs.each(idx
	tru.asn(fstDay)
	sys.newMu().asn(prfmsMu)
	hst.newPrfms().asn(prfmsA)
	hst.newPrfms().asn(prfmsB)

	idx.log("idx")
	idx.add(len).add(1).log("idx.add(len).add(1)")
	mktWeeks.cnt().log("mktWeeks cnt")
	mktWeeks.rngMrg(idx idx.add(len)).asn(bckRngTtl)
	mktWeeks.at(idx.add(len).add(1)).asn(fwdRng)
	mktWeeks.from(idx).to(len).each(bckRng
		bckRng.log("MKT RNG")
		prfmsA.cnt().log("prfmsA cnt")
		hst.newPrcp().asn(prcp)

		hst.oan().eurUsd(bckRng).asn(instr)
		instr.i(1s).ask().lst().asn(s1Lst)
		instr.is(durs).ask().asn(ask).log("INRVL")

		ask.lst().asn(lst).log("STM: LST")
		ask.rsi().asn(rsi).log("STM: RSI")
		ask.std().asn(std).log("STM: STD")
		ask.wrsi2().asn(wrsi2).log("STM: WRSI2")
		ask.rngFul().asn(rngFul).log("STM: RNG FUL")
		ask.proLst().asn(proLst).log("STM: PRO LST")
		prcp.fbr(rsi std rngFul proLst)
		tme.now().log("STM: RSI,STD,WRSI2,RNGFUL,PROLST ORTSUB 1")
		rsi.ortSub(1).asn(rsiOrtSub1)
		std.ortSub(1).asn(stdOrtSub1)
		wrsi2.ortSub(1).asn(wrsi2OrtSub1)
		rngFul.ortSub(1).asn(rngFulOrtSub1)
		proLst.ortSub(1).asn(proLstOrtSub1)
		prcp.fbr(rsiOrtSub1 stdOrtSub1 wrsi2OrtSub1 rngFulOrtSub1 proLstOrtSub1)

		ask.sma().asn(sma).log("STM: SMA")
		ask.wma().asn(wma).log("STM: WMA")
		ask.alma(6 0.85).asn(alma).log("STM: ALMA")
		prcp.fbr(sma wma alma)

		std.sclMul(2.0).asn(blngr).log("STM: BLNGR")
		sma.otrSub(0 blngr).asn(blngrBtm).log("STM: BLNGR: BTM")
		sma.otrAdd(0 blngr).asn(blngrTop).log("STM: BLNGR: TOP")
		std.sclMul(2.0).asn(blngr).log("STM: BLNGR")
		lst.otrSub(0 blngrBtm).asn(blngrPctBTop).log("STM: % BLNGR: BTM")
		blngrTop.otrSub(0 blngrBtm).asn(blngrPctBBtm).log("STM: % BLNGR: TOP")
		// %B = (Price - Lower Band)/(Upper Band - Lower Band); default: 20, 2
		blngrPctBTop.otrDiv(0 blngrPctBBtm).asn(pctB).log("STM: % BLNGR")
		prcp.fbr(pctB)

		tme.now().log("STM: S1,SMA,WMA,ALMA INRSUB 1")
		s1Lst.inrSubs(offs).asn(s1LstSub1)
		s1LstSub1.ortSub(1).asn(s1LstSub1OrtSub1)
		prcp.fbr(s1LstSub1 s1LstSub1OrtSub1)
		sma.inrSubs(offs).asn(smaSub1)
		wma.inrSubs(offs).asn(wmaSub1)
		alma.inrSubs(offs).asn(almaSub1)
		prcp.wve(smaSub1 wmaSub1 almaSub1)

		tme.now().log("STM: S1,SMA,WMA,ALMA INRSLP 1")
		s1Lst.inrSlps(offs).asn(s1LstSlp1)
		s1LstSlp1.ortSub(1).asn(s1LstSlp1OrtSub1)
		prcp.fbr(s1LstSlp1 s1LstSlp1OrtSub1)
		sma.inrSlps(offs).asn(smaSlp1)
		wma.inrSlps(offs).asn(wmaSlp1)
		alma.inrSlps(offs).asn(almaSlp1)
		prcp.wve(smaSlp1 wmaSlp1 almaSlp1)

		tme.now().log("STM: SMA,WMA,ALMA AGGSMA 2")
		sma.aggSmas(offs).asn(smaSma2)
		wma.aggSmas(offs).asn(wmaSma2)
		alma.aggSmas(offs).asn(almaSma2)
		prcp.wve(smaSma2 wmaSma2 almaSma2)

		tme.now().log("STM: SMA,WMA,ALMA AGGWMA 2")
		sma.aggWmas(offs).asn(smaWma2)
		wma.aggWmas(offs).asn(wmaWma2)
		alma.aggWmas(offs).asn(almaWma2)
		prcp.wve(smaWma2 wmaWma2 almaWma2)

		tme.now().log("STM: SMA,WMA,ALMA AGGALMA 2")
		sma.aggAlmas(offs [6] [0.85]).asn(smaAlma2)
		wma.aggAlmas(offs [6] [0.85]).asn(wmaAlma2)
		alma.aggAlmas(offs [6] [0.85]).asn(almaAlma2)
		prcp.wve(smaAlma2 wmaAlma2 almaAlma2)

		tme.now().log("STM: SMA,WMA,ALMA ORTSUB 1")
		sma.ortSub(1).asn(smaOrtSub1)
		wma.ortSub(1).asn(wmaOrtSub1)
		alma.ortSub(1).asn(almaOrtSub1)
		prcp.fbr(smaOrtSub1 wmaOrtSub1 almaOrtSub1)

		fstDay.then( // find fst day stgy
			tme.now().log("CND: S1LST CRS SMA")
			s1Lst.otrLsss([0] sma.stms()).seq(1s s1Lst.otrGtrs([0] sma.stms())).cnds().rev().asn(opnCnds0)
			// tme.now().log("CND: ALMA CRS SMA")
			// alma.otrLss(0 sma).seq(1s alma.otrGtr(0 sma)).cnds().rev().asn(opnCnds1)
			tme.now().log("CND: SMA SUPPORT")
			sma.inrLss(1).seq(1s sma.inrGtr(1)).cnds().rev().asn(opnCnds2)
			tme.now().log("CND: WMA SUPPORT")
			wma.inrLss(1).seq(1s wma.inrGtr(1)).cnds().rev().asn(opnCnds3)
			// tme.now().log("CND: ALMA SUPPORT")
			// alma.inrLss(1).seq(1s alma.inrGtr(1)).cnds().rev().asn(opnCnds4)

			hst.newCnds().asn(opnCnds).mrg(opnCnds0 opnCnds2 opnCnds3)

			tme.now().log("STGY LONGS")
			opnCnds.cnt().log("OPN CND CNT")
			opnCnds.each(opnCnd
				opnCnd.log(" ~ TUNING CND")
				opnCnd.longs(prfLims losLims durLims [instr]).stgys().rev().asn(longStgys)
				// longStgys.cnt().log("LONG STGY CNT")
				longStgys.pllEach(stgy
					stgy.port().prfm().asn(prfm)
					prfm.port().splt(0.0).asn(splt)
					prcp.splt(splt).tuneTilScsPct(95.0 0.0 2).asn(tunedPrfm)
					tunedPrfm.ana().pnlPct().gtr(5.0).then(
						prfmsMu.lck()
						prfmsA.push(tunedPrfm)
						prfmsMu.ulck()
					)
				)
			)
			fls.asn(fstDay 2)
		).else( // tune fst day stgy
			prfmsA.pllEach(prfmA
				prfmA.port().splt(0.0).asn(splt)
				prcp.splt(splt).tuneTilScsPct(95.0 0.0 2).asn(tunedPrfmB)
				prfmsMu.lck()
				prfmsB.push(tunedPrfmB)
				prfmsMu.ulck()
			)
			prfmsA.asn(prfmsTmp)
			prfmsB.asn(prfmsA 2)
			prfmsTmp.clr()
			prfmsTmp.asn(prfmsB 2)
		)
	) // end training

	bckRngTtl.log("### bck rng total ###")
	fwdRng.log("### fwd rng ###")
	prfmsA.cnt().log("fwd prfmsA cnt")
	prfmsA.srtDscPnlPct().to(10).each(prfmA // validation rng
	// prfmsA.srtAscLosLimMax().to(10).each(prfmA // validation rng
		prfmA.port().rng(fwdRng).prfm().asn(runPrfm)
		prfmA.log("PRFM BCK")
		runPrfm.log("PRFM FWD")
		prfmA.dlt(runPrfm).asn(runDlt)
		runPrfm.ana().pnlPct().gtr(0.0).then(
			runDlt.log("  ++++ FWD DLT")
		).else(
			runDlt.log("  ---- FWD DLT")
		)
	)

)



// opnCnds.each(opnCnd
// 	opnCnd.log(" ~ TUNING")
// 	opnCnd.longs(prfLims losLims durLims [instr]).stgys().rev().asn(longStgys)
// 	longStgys.cnt().log("LONG STGY CNT")
// 	longStgys.each(stgy
// 		stgy.port().prfm().asn(prfm)
// 		prfm.port().splt(0.0).asn(splt)
// 		prcp.splt(splt).tuneTilScsPct(95.0 0.0 2).asn(tunedPrfm)
// 		prfm.dlt(tunedPrfm).log(" * TUNED")
// 		tunedPrfm.ana().pnlPct().gtr(5.0).then(
// 			tunedPrfm.port().rng(fwdRng).prfm().asn(fwdTunedPrfm)
// 			tunedPrfm.dlt(fwdTunedPrfm).asn(fwdDlt)
// 							prfm.log("      PRFM FST")
// 				 tunedPrfm.log("    PRFM TUNED")
// 			fwdTunedPrfm.log("PRFM FWD TUNED")
// 			fwdTunedPrfm.ana().pnlPct().gtr(0.0).then(
// 				    fwdDlt.log("     + FWD DLT")
// 						fwdTunedPrfm.port().rng(fwdRng).prfm().asn(runPrfm)
// 						fwdTunedPrfm.dlt(runPrfm).asn(runDlt)
// 						runPrfm.log("      PRFM RUN")
// 				runPrfm.ana().pnlPct().gtr(0.0).then(
// 						runDlt.log("    ++ RUN DLT")
// 				).else(
// 					  runDlt.log("    -- RUN DLT")
// 				)
// 			).else(
// 				    fwdDlt.log("     - FWD DLT")
// 			)

// 			// tunedPrfm.port().rng(fwdRng).prfm().asn(fwdTunedPrfm)
// 			// tunedPrfm.dlt(fwdTunedPrfm).log(" **  FWD")
// 			// fwdTunedPrfm.ana().pnlPct().gtr(0.0).then(
// 			// 	fwdTunedPrfm.port().rng(fwdRng).prfm().asn(runPrfm)
// 			// 	fwdTunedPrfm.dlt(runPrfm).asn(runDlt)
// 			// 		        prfm.log("      PRFM FST")
// 			// 		   tunedPrfm.log("    PRFM TUNED")
// 			// 		fwdTunedPrfm.log("PRFM FWD TUNED")
// 			// 		     runPrfm.log("      PRFM RUN")
// 			// 	runPrfm.ana().pnlPct().gtr(0.0).then(
// 			// 		      runDlt.log("  ++++ RUN DLT")
// 			// 	).else(
// 			// 		      runDlt.log("  ---- RUN DLT")
// 			// 	)
// 			// )
// 		)
// 	)
// )


"DONE".log()