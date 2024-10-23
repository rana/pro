# (pipPerDay:)[0-9]*\.?[0-9]*


hst.oan().instrs(2018-5-1).asn(instrs)
unts.addsLeq(5 5 120).asn(aggLens) # 28800 = 8 hours
flts.addsLeq(1.95 0.01 2.05).asn(stds)
tmes.addsLeq(2s 1s 2s).asn(seqDurs)
flts.addsLeq(1.0 1.0 20.0).asn(pnlLims)
tmes.addsLeq(4h 1h 4h).asn(wndLims)

instrs.each(instr
	instr.s1().asks().lst().asn(stm1)
	aggLens.each(aggLen
		# BOLLINGER
		stm1.aggSma(aggLen).asn(stm1Sma)
		stm1.aggStd(aggLen).asn(stm1Std)
		stds.pllEach(std
			ana.newStgyPrfmChan(tru).asn(blngrBtmCrsUpLongC)
			ana.newStgyPrfmChan(tru).asn(blngrBtmCrsDwnShrtC)
			ana.newStgyPrfmChan(tru).asn(blngrTopCrsUpLongC)
			ana.newStgyPrfmChan(tru).asn(blngrTopCrsDwnShrtC)
			stm1Std.sclMul(std).asn(blngr)
			### BTM
			stm1Sma.otrSub(0 blngr).asn(blngrBtm)
			stm1.otrLss(0 blngrBtm).asn(stm1BtmLss)
			stm1.otrGtr(0 blngrBtm).asn(stm1BtmGtr)
			### TOP
			stm1Sma.otrAdd(0 blngr).asn(blngrTop)
			stm1.otrLss(0 blngrTop).asn(stm1TopLss)
			stm1.otrGtr(0 blngrTop).asn(stm1TopGtr)
			seqDurs.each(seqDur
				stm1BtmLss.seq(seqDur stm1BtmGtr).asn(blngrBtmCrsUp)
				stm1BtmGtr.seq(seqDur stm1BtmLss).asn(blngrBtmCrsDwn)
				stm1TopLss.seq(seqDur stm1TopGtr).asn(blngrTopCrsUp)
				stm1TopGtr.seq(seqDur stm1TopLss).asn(blngrTopCrsDwn)
				pnlLims.pllEach(prfLim
					pnlLims.pllEach(losLim
						wndLims.pllEach(wndLim
							### BTM
							blngrBtmCrsUp.long(prfLim losLim wndLim instr).asn(blngrBtmCrsUpLong)
							blngrBtmCrsUpLong.pipPerDay().gtr(0.0).then(
								#blngrBtmCrsUpLong.prfm().log()
								#blngrBtmCrsUpLong.prfm().usdPerDay().gtr(0.0).then(
									blngrBtmCrsUpLong.prfm().name("blnger-btm-crs-up-instr-long")
									blngrBtmCrsUpLongC.push(blngrBtmCrsUpLong.prfm())
								#)
							)
							blngrBtmCrsDwn.shrt(prfLim losLim wndLim instr).asn(blngrBtmCrsDwnShrt)
							blngrBtmCrsDwnShrt.pipPerDay().gtr(0.0).then(
								#blngrBtmCrsDwnShrt.prfm().log()
								#blngrBtmCrsDwnShrt.prfm().usdPerDay().gtr(0.0).then(
									blngrBtmCrsDwnShrt.prfm().name("blnger-btm-crs-dwn-instr-shrt")
									blngrBtmCrsDwnShrtC.push(blngrBtmCrsDwnShrt.prfm())
								#)
							)
							### TOP
							blngrTopCrsUp.long(prfLim losLim wndLim instr).asn(blngrTopCrsUpLong)
							blngrTopCrsUpLong.pipPerDay().gtr(0.0).then(
								#blngrTopCrsUpLong.prfm().log()
								#blngrTopCrsUpLong.prfm().usdPerDay().gtr(0.0).then(
									blngrTopCrsUpLong.prfm().name("blnger-top-crs-up-instr-long")
									blngrTopCrsUpLongC.push(blngrTopCrsUpLong.prfm())
								#)
							)
							blngrTopCrsDwn.shrt(prfLim losLim wndLim instr).asn(blngrTopCrsDwnShrt)
							blngrTopCrsDwnShrt.pipPerDay().gtr(0.0).then(
								#blngrTopCrsDwnShrt.prfm().log()
								#blngrTopCrsDwnShrt.prfm().usdPerDay().gtr(0.0).then(
									blngrTopCrsDwnShrt.prfm().name("blnger-top-crs-dwn-instr-shrt")
									blngrTopCrsDwnShrtC.push(blngrTopCrsDwnShrt.prfm())
								#)
							)
						)
					)
				)
			) # seqDurs
			blngrBtmCrsUpLongC.cldSavCls()
			blngrBtmCrsDwnShrtC.cldSavCls()
			blngrTopCrsUpLongC.cldSavCls()
			blngrTopCrsDwnShrtC.cldSavCls()
		) # stds
	) # aggLens
) # instrs

