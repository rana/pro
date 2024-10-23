tme.now().log("START:")
// plt.scl(1.6)


[1m 2m 4m 8m 16m 32m 64m 128m 256m 512m 1024m].asn(dursA)
[5m 10m 15m 20m 30m 50m 60m 100m 120m 200m 240m 400m 480m].asn(dursB)
[1m	2m 3m 5m 8m 13m	21m	34m	55m	89m	144m 233m 377m 610m].asn(dursC) // 1h29m|2h24m|3h53m|6h17m|10h10m
[1m	2m 3m 5m 8m 13m	21m	34m	42m 55m 68m	89m 110m 144m 178m 233m 288m 377m 466m 610m].asn(dursCa)
[1m	2m 3m 5m 8m 13m	21m	34m	55m	89m	144m 233m 377m 610m 987m].asn(dursCb)
tmes.addsLss(5m 400m 5m).asn(dursE)


dursC.asn(durs)
// unts.fibsLeq(610).asn(offs)
// unts.fibsLeq(610).from(1).asn(lens) // start at 2
// tmes.addsLeq(1m 480m 1m).asn(durs)
unts.addsLeq(1 200 3).asn(offs)
unts.addsLeq(1 200 3).asn(lens)

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
	sys.newMu().asn(prfmsMu)
	hst.newPrfms().asn(prfms)

	idx.log("idx")
	mktRngs.cnt().log("mktRngs cnt")
	mktRngs.at(idx.add(len)).asn(fwdRng)
	fwdRng.log("fwdRng")

	mktRngs.rngMrg(idx idx.add(len).sub(1)).asn(bckRng)
	bckRng.log("bckRng")
	prfms.cnt().log("prfms cnt")
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

	tme.now().log("CND: CRS")
	// 0.asn(lhsIdx)
	// lhsIdx.add(1).asn(rhsIdx)
	// alma.stms().at(lhsIdx).asn(lhsStm)
	// blngrBtm.stms().from(rhsIdx).asn(rhsStms)
	// lhsStm.otrLsss([0] rhsStms).seq(1s lhsStm.otrGtrs([0] rhsStms)).cnds().rev().asn(opnCnds)
	m1Alma.otrLss(0 m15Sma).seq(1s m1Alma.otrGtr(0 m15Sma)).asn(opnCnd)
	hst.newCnds(opnCnd).asn(opnCnds)

	tme.now().log("STGY LONGS")
	opnCnds.cnt().log("OPN CND CNT")
	opnCnds.each(opnCnd
		opnCnd.log(" ~ TUNING CND")
		opnCnd.longs(prfLims losLims durLims [instr]).stgys().rev().asn(longStgys)
		// longStgys.cnt().log("LONG STGY CNT")
		longStgys.pllEach(stgy
			stgy.port().prfm().asn(prfm)
			prfm.port().splt(0.0).asn(splt)
			// prcp.splt(splt).tuneTilScsPct(95.0 0.0 2).asn(tunedPrfm)
			prcp.splt(splt).tuneTil(0.0 stmCntLim).asn(tunedPrfm)
			tunedPrfm.ana().pnlPct().gtr(5.0).then(
				prfmsMu.lck()
				prfms.push(tunedPrfm)
				prfmsMu.ulck()
			)
		)
	)

	bckRng.log("### bck rng ###")
	fwdRng.log("### fwd rng ###")
	prfms.cnt().log("fwd prfms cnt")
	prfms.srtDscPnlPct().to(10).each(prfmA // validation rng
	// prfms.srtAscLosLimMax().to(10).each(prfmA // validation rng
		prfmA.port().rng(fwdRng).prfm().asn(runPrfm)
		prfmA.log("PRFM BCK")
		runPrfm.log("PRFM FWD")
		runPrfm.ana().pnlPct().gtr(0.0).then(
			" +++".log()
		).else(
			" ---".log()
		)
	)

)

"DONE".log()