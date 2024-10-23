tme.now().log("START")

[1m 2m 4m 8m 16m 32m 64m 128m 256m 512m].asn(dursA)
[1m 5m 10m 15m 20m 30m 50m 60m 100m 120m 200m 240m 400m 480m].asn(dursB)
[1m 5m 10m 20m 50m 100m 200m 400m].asn(dursBa)
[1m	2m 3m 5m 8m 13m	21m	34m	55m	89m	144m 233m 377m 610m].asn(dursC) // 1h29m|2h24m|3h53m|6h17m|10h10m
[1m	2m 3m 5m 8m 13m	21m	34m	42m 55m 68m	89m 110m 144m 178m 233m 288m 377m 466m 610m].asn(dursCa)
[1m	2m 3m 5m 8m 13m	21m	34m	55m	89m	144m 233m 377m 610m 987m].asn(dursCb)
[1m 2m 3m 5m 8m 13m	21m	34m	55m	89m	144m 233m].asn(dursCc)
tmes.addsLss(5m 400m 5m).asn(dursE)


dursCc.asn(durs)
// unts.fibsLeq(610).asn(offs)
// unts.fibsLeq(610).from(1).asn(lens) // start at 2
// tmes.addsLeq(1m 480m 1m).asn(durs)
unts.addsLeq(1 200 2).asn(offs)
unts.addsLeq(1 200 2).asn(lens)

// [1m 5m 15m 30m 1h 4h].asn(durs)
// unts.fibsLeq(55).asn(offs)
// unts.fibsLeq(55).from(1).asn(lens)
flts.addsLeq(10.0 40.0 5.0).asn(prfLims)
flts.addsLeq(10.0 40.0 5.0).asn(losLims)
tmes.addsLeq(4h 4h 15m).asn(durLims)

hst.oan().eurUsd().mktWeeks().asn(mktWeeks)
hst.oan().eurUsd().mktDays().asn(mktDays)
hst.oan().eurUsd().mktHrs().asn(mktHrs)
mktDays.asn(mktRngs)
mktRngs.cnt().log("mkt rng cnt")

2.asn(stmCntLim).log("stmCntLim")
2.asn(len).log("len")
unts.addsLss(0 mktRngs.lstIdx().sub(len) 1).asn(idxs)
idxs.each(idx
	tru.asn(fstStgy)
	sys.newMu().asn(prfmsMu)
	hst.newPrfms().asn(prfmsA)
	hst.newPrfms().asn(prfmsB)

	idx.log("idx")
	mktRngs.cnt().log("mktRngs cnt")
	mktRngs.at(idx.add(len)).asn(fwdRng)
	fwdRng.log("fwdRng")
	mktRngs.from(idx).to(len).each(bckRng
		bckRng.log("bckRng")
		prfmsA.cnt().log("prfmsA cnt")
		hst.newPrcp().asn(prcp)

		hst.oan().eurUsd(bckRng).asn(instr)
		instr.i(1s).ask().lst().asn(s1Lst)
		instr.is(durs).ask().asn(ask).log("INRVL")

		instr.i(1m).ask().asn(m1Ask)
		instr.i(15m).ask().asn(m15Ask)
		m1Ask.alma(6 0.85).asn(m1Alma)
		m15Ask.sma().asn(m15Sma)


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
		// tme.now().log("STM: RSI,STD,WRSI2,RNGFUL,PROLST ORTSUB 2")
		// rsi.ortSub(2).asn(rsiOrtSub2)
		// std.ortSub(2).asn(stdOrtSub2)
		// wrsi2.ortSub(2).asn(wrsi2OrtSub2)
		// rngFul.ortSub(2).asn(rngFulOrtSub2)
		// proLst.ortSub(2).asn(proLstOrtSub2)
		// prcp.fbr(rsiOrtSub2 stdOrtSub2 wrsi2OrtSub2 rngFulOrtSub2 proLstOrtSub2)
		// tme.now().log("STM: RSI,STD,WRSI2,RNGFUL,PROLST ORTSUB 3")
		// rsi.ortSub(3).asn(rsiOrtSub3)
		// std.ortSub(3).asn(stdOrtSub3)
		// wrsi2.ortSub(3).asn(wrsi2OrtSub3)
		// rngFul.ortSub(3).asn(rngFulOrtSub3)
		// proLst.ortSub(3).asn(proLstOrtSub3)
		// prcp.fbr(rsiOrtSub3 stdOrtSub3 wrsi2OrtSub3 rngFulOrtSub3 proLstOrtSub3)
		// tme.now().log("STM: RSI,STD,WRSI2,RNGFUL,PROLST ORTDIV 1")
		// rsi.ortDiv(1).asn(rsiOrtDiv1)
		// std.ortDiv(1).asn(stdOrtDiv1)
		// wrsi2.ortDiv(1).asn(wrsi2OrtDiv1)
		// rngFul.ortDiv(1).asn(rngFulOrtDiv1)
		// proLst.ortDiv(1).asn(proLstOrtDiv1)
		// prcp.fbr(rsiOrtDiv1 stdOrtDiv1 wrsi2OrtDiv1 rngFulOrtDiv1 proLstOrtDiv1)
		// tme.now().log("STM: RSI,STD,WRSI2,RNGFUL,PROLST ORTDIV 2")
		// rsi.ortDiv(2).asn(rsiOrtDiv2)
		// std.ortDiv(2).asn(stdOrtDiv2)
		// wrsi2.ortDiv(2).asn(wrsi2OrtDiv2)
		// rngFul.ortDiv(2).asn(rngFulOrtDiv2)
		// proLst.ortDiv(2).asn(proLstOrtDiv2)
		// prcp.fbr(rsiOrtDiv2 stdOrtDiv2 wrsi2OrtDiv2 rngFulOrtDiv2 proLstOrtDiv2)
		// tme.now().log("STM: RSI,STD,WRSI2,RNGFUL,PROLST ORTDIV 3")
		// rsi.ortDiv(3).asn(rsiOrtDiv3)
		// std.ortDiv(3).asn(stdOrtDiv3)
		// wrsi2.ortDiv(3).asn(wrsi2OrtDiv3)
		// rngFul.ortDiv(3).asn(rngFulOrtDiv3)
		// proLst.ortDiv(3).asn(proLstOrtDiv3)
		// prcp.fbr(rsiOrtDiv3 stdOrtDiv3 wrsi2OrtDiv3 rngFulOrtDiv3 proLstOrtDiv3)

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
		// tme.now().log("STM: SMA,WMA,ALMA ORTSUB 2")
		// sma.ortSub(2).asn(smaOrtSub2)
		// wma.ortSub(2).asn(wmaOrtSub2)
		// alma.ortSub(2).asn(almaOrtSub2)
		// prcp.fbr(smaOrtSub2 wmaOrtSub2 almaOrtSub2)
		// tme.now().log("STM: SMA,WMA,ALMA ORTSUB 3")
		// sma.ortSub(3).asn(smaOrtSub3)
		// wma.ortSub(3).asn(wmaOrtSub3)
		// alma.ortSub(3).asn(almaOrtSub3)
		// prcp.fbr(smaOrtSub3 wmaOrtSub3 almaOrtSub3)
		// tme.now().log("STM: SMA,WMA,ALMA ORTDIV 1")
		// sma.ortDiv(1).asn(smaOrtDiv1)
		// wma.ortDiv(1).asn(wmaOrtDiv1)
		// alma.ortDiv(1).asn(almaOrtDiv1)
		// prcp.fbr(smaOrtDiv1 wmaOrtDiv1 almaOrtDiv1)
		// tme.now().log("STM: SMA,WMA,ALMA ORTDIV 2")
		// sma.ortDiv(2).asn(smaOrtDiv2)
		// wma.ortDiv(2).asn(wmaOrtDiv2)
		// alma.ortDiv(2).asn(almaOrtDiv2)
		// prcp.fbr(smaOrtDiv2 wmaOrtDiv2 almaOrtDiv2)
		// tme.now().log("STM: SMA,WMA,ALMA ORTDIV 3")
		// sma.ortDiv(3).asn(smaOrtDiv3)
		// wma.ortDiv(3).asn(wmaOrtDiv3)
		// alma.ortDiv(3).asn(almaOrtDiv3)
		// prcp.fbr(smaOrtDiv3 wmaOrtDiv3 almaOrtDiv3)

		fstStgy.then( // find fst period stgy
			// tme.now().log("CND: S1LST CRS SMA")
			// s1Lst.otrLsss([0] sma.stms()).seq(1s s1Lst.otrGtrs([0] sma.stms())).cnds().rev().asn(opnCnds0)
			// tme.now().log("CND: ALMA CRS SMA")
			// alma.otrLss(0 sma).seq(1s alma.otrGtr(0 sma)).cnds().rev().asn(opnCnds1)
			// tme.now().log("CND: SMA SUPPORT")
			// sma.inrLss(1).seq(1s sma.inrGtr(1)).cnds().rev().asn(opnCnds2)
			// tme.now().log("CND: WMA SUPPORT")
			// wma.inrLss(1).seq(1s wma.inrGtr(1)).cnds().rev().asn(opnCnds3)
			// tme.now().log("CND: ALMA SUPPORT")
			// alma.inrLss(1).seq(1s alma.inrGtr(1)).cnds().rev().asn(opnCnds4)
			// hst.newCnds().asn(opnCnds).mrg(opnCnds0 opnCnds2 opnCnds3)

			tme.now().log("CND: CRS")
			// 0.asn(lhsIdx)
			// lhsIdx.add(1).asn(rhsIdx)
			// alma.stms().at(lhsIdx).asn(lhsStm)
			// blngrBtm.stms().from(rhsIdx).asn(rhsStms)
			// lhsStm.otrLsss([0] rhsStms).seq(1s lhsStm.otrGtrs([0] rhsStms)).cnds().rev().asn(opnCnds)
			m1Alma.otrLss(0 m15Sma).seq(1s m1Alma.otrGtr(0 m15Sma)).asn(opnCnd)
			hst.newCnds(opnCnd).asn(opnCnds)

			tme.now().log("STGYS")
			opnCnds.cnt().log("OPN CND CNT")
			opnCnds.each(opnCnd
				opnCnd.log(" ~ TUNING CND")
				opnCnd.longs(prfLims losLims durLims [instr]).stgys().rev().asn(stgys)
				stgys.pllEach(stgy
					stgy.port().prfm().asn(prfm)
					prfm.port().splt(0.0).asn(splt)
					// prcp.splt(splt).tuneTilScsPct(95.0 0.0 2).asn(tunedPrfm)
					prcp.splt(splt).tuneTil(0.0 stmCntLim).asn(tunedPrfm)
					tunedPrfm.ana().pnlPct().gtr(5.0).then(
						prfmsMu.lck()
						prfmsA.push(tunedPrfm)
						prfmsMu.ulck()
					)
				)
			)
			fls.asn(fstStgy 2)
		).else( // tune existing stgy
			prfmsA.pllEach(prfmA
				prfmA.port().splt(0.0).asn(splt)
				// prcp.splt(splt).tuneTilScsPct(95.0 0.0 2).asn(tunedPrfmB)
				prcp.splt(splt).tuneTil(0.0 stmCntLim).asn(tunedPrfmB)
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

	mktRngs.rngMrg(idx idx.add(len).sub(1)).log("### bck rng ###")
	fwdRng.log("### fwd rng ###")
	prfmsA.cnt().log("fwd prfmsA cnt")
	prfmsA.srtDscPnlPct().to(10).each(prfmA // validation rng
	// prfmsA.srtAscLosLimMax().to(10).each(prfmA // validation rng
		prfmA.port().rng(fwdRng).prfm().asn(fwdPrfm)
		prfmA.log("PRFM BCK")
		fwdPrfm.log("PRFM FWD")
	)
)
"DONE".log()