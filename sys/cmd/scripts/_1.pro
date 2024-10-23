tme.now().log("START:")

[1m 5m 10m 20m 50m 100m 200m 400m].asn(dursB)
[1m	2m 3m 5m 8m 13m	21m	34m	55m	89m	144m 233m 377m].asn(dursC) // 1h29m|2h24m|3h53m|6h17m|10h10m

dursC.asn(durs)
unts.addsLeq(1 200 2).asn(offs)
unts.addsLeq(1 200 2).asn(lens)
flts.addsLeq(10.0 40.0 5.0).asn(prfLims)
flts.addsLeq(5.0 5.0 5.0).asn(losLims)
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
		instr.i(2m).ask().asn(m2Ask)
		instr.i(3m).ask().asn(m3Ask)
		instr.i(5m).ask().asn(m5Ask)
		instr.i(8m).ask().asn(m8Ask)
		instr.i(13m).ask().asn(m13Ask)
		instr.i(21m).ask().asn(m21Ask)


		// prcp.fbr(
		// 	ask.std().ortSub(2).rev()
		// )
		// prcp.fbr(
		// 	ask.std()
		// )



		prcp.fbr(
			ask.wrsi2()
		)
		prcp.fbr(
			ask.rsi()
		)
		prcp.fbr(
			ask.std()
		)

		// ask.std().asn(std)
		// prcp.fbr(
		// 	ask.rsi()
		// )
		// prcp.fbr(
		// 	std
		// )
		// prcp.fbr(
		// 	std.aggStd(13)
		// )

		fstStgy.then( // find fst period stgy
			tme.now().log("CND: CRS")
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
	prfmsA.srtDscPnlPct().to(4).each(prfmA // validation rng
	// prfmsA.srtAscLosLimMax().to(10).each(prfmA // validation rng
		prfmA.port().rng(fwdRng).prfm().asn(fwdPrfm)
		prfmA.log("PRFM BCK")
		fwdPrfm.log("PRFM FWD")
	)
)
"DONE".log()