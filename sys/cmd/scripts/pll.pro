# (revPerDay:)[0-9]*\.?[0-9]*


ana.oan().hst().instrs().asn(instrs)
unts.addsLeq(10 10 60).push(5 15).asn(aggLens)
flts.addsLeq(1.95 0.01 2.05).asn(stds)
tmes.addsLeq(2s 1s 2s).asn(seqDurs)
flts.addsLeq(1.0 1.0 20.0).asn(pnlLims)
tmes.addsLeq(2h 1h 2h).asn(wndLims)

instrs.pllEach(instr
	flt.min.asn(longMax)
	flt.min.asn(shrtMax)
	instr.s1().asks().lst().asn(stm1)
	aggLens.pllEach(aggLen
		# BOLLINGER
		stm1.aggSma(aggLen).asn(stm1Sma)
		stm1.aggStd(aggLen).asn(stm1Std)
		stds.pllEach(std
			hst.newChanStgyPrfm().asn(chan)
			stm1Std.sclMul(std).asn(blngr)
			# BTM
			stm1Sma.otrSub(0 blngr).asn(blngrBtm)
			stm1.otrLss(0 blngrBtm).asn(stm1BtmLss)
			stm1.otrGtr(0 blngrBtm).asn(stm1BtmGtr)
			# TOP
			stm1Sma.otrAdd(0 blngr).asn(blngrTop)
			stm1.otrLss(0 blngrTop).asn(stm1TopLss)
			stm1.otrGtr(0 blngrTop).asn(stm1TopGtr)
			seqDurs.pllEach(seqDur
				stm1BtmLss.seq(seqDur stm1BtmGtr).asn(blngrBtmCrsUp)
				stm1BtmGtr.seq(seqDur stm1BtmLss).asn(blngrBtmCrsDwn)
				stm1TopLss.seq(seqDur stm1TopGtr).asn(blngrTopCrsUp)
				stm1TopGtr.seq(seqDur stm1TopLss).asn(blngrTopCrsDwn)
				pnlLims.pllEach(prfLim
					pnlLims.pllEach(losLim
						wndLims.pllEach(wndLim
							# BTM
							#blngrBtmCrsDwn.long(prfLim losLim wndLim instr).asn(blngrBtmCrsDwnLong)
							#blngrBtmCrsDwnLong.prfm().revPerDay().gtr(0.0).then(
							#	blngrBtmCrsDwnLong.prfm().name("blnger-btm-crs-dwn-instr-long")
							#	chan.push(blngrBtmCrsDwnLong.prfm())
							#)
							blngrBtmCrsUp.long(prfLim losLim wndLim instr).asn(blngrBtmCrsUpLong)
							blngrBtmCrsUpLong.prfm().revPerDay().gtr(0.0).then(
								blngrBtmCrsUpLong.prfm().name("blnger-btm-crs-up-instr-long")
								chan.push(blngrBtmCrsUpLong.prfm())
							)
							blngrBtmCrsDwn.shrt(prfLim losLim wndLim instr).asn(blngrBtmCrsDwnShrt)
							blngrBtmCrsDwnShrt.prfm().revPerDay().gtr(0.0).then(
								blngrBtmCrsDwnShrt.prfm().name("blnger-btm-crs-dwn-instr-shrt")
								chan.push(blngrBtmCrsDwnShrt.prfm())
							)
							blngrBtmCrsUp.shrt(prfLim losLim wndLim instr).asn(blngrBtmCrsUpShrt)
							blngrBtmCrsUpShrt.prfm().revPerDay().gtr(0.0).then(
								blngrBtmCrsUpShrt.prfm().name("blnger-btm-crs-up-instr-shrt")
								chan.push(blngrBtmCrsUpShrt.prfm())
							)
							# TOP
							blngrTopCrsDwn.long(prfLim losLim wndLim instr).asn(blngrTopCrsDwnLong)
							blngrTopCrsDwnLong.prfm().revPerDay().gtr(0.0).then(
								blngrTopCrsDwnLong.prfm().name("blnger-top-crs-dwn-instr-long")
								chan.push(blngrTopCrsDwnLong.prfm())
							)
							blngrTopCrsUp.long(prfLim losLim wndLim instr).asn(blngrTopCrsUpLong)
							blngrTopCrsUpLong.prfm().revPerDay().gtr(0.0).then(
								blngrTopCrsUpLong.prfm().name("blnger-top-crs-up-instr-long")
								chan.push(blngrTopCrsUpLong.prfm())
							)
							blngrTopCrsDwn.shrt(prfLim losLim wndLim instr).asn(blngrTopCrsDwnShrt)
							blngrTopCrsDwnShrt.prfm().revPerDay().gtr(0.0).then(
								blngrTopCrsDwnShrt.prfm().name("blnger-top-crs-dwn-instr-shrt")
								chan.push(blngrTopCrsDwnShrt.prfm())
							)
							blngrTopCrsUp.shrt(prfLim losLim wndLim instr).asn(blngrTopCrsUpShrt)
							blngrTopCrsUpShrt.prfm().revPerDay().gtr(0.0).then(
								blngrTopCrsUpShrt.prfm().name("blnger-top-crs-up-instr-shrt")
								chan.push(blngrTopCrsUpShrt.prfm())
							)
						)
					)
				)
			) # seqDurs
			chan.stgys().cldSav()
		) # stds
	) # aggLens
) # instrs

