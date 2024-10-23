1.0.asn(pltScl)
hst.oan().eurUsd().asn(instr)

// 2018-5-28/0h.asn(start)  2018-5-28/2h.asn(end)   

// hst.oan().eurUsd().s(1).bids().asn(s1)
// hst.oan().eurUsd().m(9).bids().asn(m9)
// hst.oan().eurUsd().m(20).bids().asn(m20)
// hst.oan().eurUsd().m(50).bids().asn(m50)

2h.asn(pltRngInc)
//instr.tmes().lst()).each(
tmes.addsLeq(instr.tmes().fst()  pltRngInc  instr.tmes().fst().add(24h)).each( 
	start 
	start.add(pltRngInc).asn(end)
	
	hst.oan().eurUsd(start end).s(1).bids().asn(s1)
	hst.oan().eurUsd(start.sub(9m) end).m(9).bids().asn(m9)
	hst.oan().eurUsd(start.sub(30m) end).m(30).bids().asn(m20)
	hst.oan().eurUsd(start.sub(60m) end).m(60).bids().asn(m50)

	// instr
	m9.alma(6 0.85).asn(m9Alma)
	m20.sma().asn(m20Sma)
	m50.sma().asn(m50Sma)
	m20.std().sclMul(2.0).asn(m20Blngr)
	m20Sma.otrSub(0 m20Blngr).asn(m20BlngrBtm)
	m20Sma.otrAdd(0 m20Blngr).asn(m20BlngrTop)

	// slp
	m9Alma.inrSlp(1).sclMul(instr.slpScl()).asn(m9AlmaSlp)
	m20Sma.inrSlp(1).sclMul(instr.slpScl()).asn(m20SmaSlp)
	m50Sma.inrSlp(1).sclMul(instr.slpScl()).asn(m50SmaSlp)

	// pro
	m9.proLst().asn(m9Pro)
	m20.proLst().asn(m20Pro)
	m50.proLst().asn(m50Pro)

	// m20Pro cnd
	0.000.asn(m20ProThsh)
	m20Pro.sclGtr(m20ProThsh).asn(m20ProTop)
	m20Pro.sclLeq(m20ProThsh).asn(m20ProBtm)
	m20ProTop.seq(1s m20ProBtm).asn(m20ProDwn)

	// m20Pro stgy
	m20ProDwn.asn(stgy0Opn)
	stgy0Opn.shrt(4.0 8.0 1h instr).asn(stgy0)//.prfm().log()
	stgy0.port().trds().selPipGeq(3.0).asn(trds0)
	trds0.cnt().log()
	trds0.opnTmes().asn(trds0Opns)
	trds0.clsTmes().asn(trds0Clss)

	// at stms
	m9AlmaSlp.aggAlma(9 6 0.85).asn(m9AlmaSlpAlma)
	m20SmaSlp.aggAlma(9 6 0.85).asn(m20SmaSlpAlma)
	m50SmaSlp.aggAlma(9 6 0.85).asn(m50SmaSlpAlma)
	m9AlmaSlpAlma.at(trds0Opns).asn(m9AlmaSlpAlmaAtOpn)
	m20SmaSlpAlma.at(trds0Opns).asn(m20SmaSlpAlmaAtOpn)
	m50SmaSlpAlma.at(trds0Opns).asn(m50SmaSlpAlmaAtOpn)
	m9AlmaSlpAlma.at(trds0Clss).asn(m9AlmaSlpAlmaAtCls)
	m20SmaSlpAlma.at(trds0Clss).asn(m20SmaSlpAlmaAtCls)
	m50SmaSlpAlma.at(trds0Clss).asn(m50SmaSlpAlmaAtCls)

	plt.newVrt().asn(vrt1)
	clr.grey800.asn(instrClr)
	clr.indigo300.asn(m9Clr)
	clr.yellow300.asn(m20AlmaClr)
	clr.pink300.asn(m50AlmaClr)
	clr.orange300.asn(m20Clr)
	clr.red300.asn(m50Clr)
	clr.deepPurple800.asn(allClr)

	vrt1.plt(plt.newStm().asn(pltInstr))
	pltInstr.stm(s1.lst() instrClr)
	pltInstr.stm(m9Alma m9Clr)
	pltInstr.stm(m20Sma m20Clr)
	pltInstr.stm(m50Sma m50Clr)
	pltInstr.stm(m20BlngrBtm clr.lime700)
	pltInstr.stm(m20BlngrTop clr.lime700)
	pltInstr.cnd(stgy0Opn clr.blueGrey900.opa(0.01))
	pltInstr.vrtScl(0.4.mul(pltScl))//.hrzScl(pltScl)
	pltInstr.x().vis(tru)

	vrt1.plt(plt.newStm().asn(pltSlp))
	pltSlp.stm(m9AlmaSlp m9Clr)
	pltSlp.stm(m20SmaSlp m20Clr)
	pltSlp.stm(m50SmaSlp m50Clr)
	pltSlp.cnd(stgy0Opn clr.blueGrey900)
	pltSlp.hrzLn(0.0 clr.grey500 1)
	pltSlp.vrtScl(0.4.mul(pltScl))//.hrzScl(pltScl)

	vrt1.plt(plt.newStm().asn(pltPro))
	pltPro.stm(m9Pro m9Clr)
	pltPro.stm(m20Pro m20Clr)
	pltPro.stm(m50Pro m50Clr)
	pltPro.hrzLn(0.5 clr.grey500 2)
	pltPro.hrzLn(0.1 clr.grey500 1)
	pltPro.hrzLn(0.9 clr.grey500 1)
	pltPro.hrzLn(0.01 clr.grey500 1)
	pltPro.vrtScl(0.2.mul(pltScl))//.hrzScl(pltScl)



	// vrt1.plt(plt.newHrz().asn(hrz1))
	// hrz1.plt(plt.newFltsSctr().asn(hrz1plt1))
	// hrz1plt1.flts(m9Clr 9 m9AlmaSlpAlmaAtOpn)
	// hrz1plt1.vrtScl(0.2.mul(pltScl))//.hrzScl(pltScl)
	// hrz1.plt(plt.newFltsSctr().asn(hrz1plt2))
	// hrz1plt2.flts(m9Clr 9 m9AlmaSlpAlmaAtCls)
	// hrz1plt2.flts(m20Clr 9 m20SmaSlpAlmaAtCls)
	// hrz1plt2.flts(m50Clr 9 m50SmaSlpAlmaAtCls)
	// hrz1plt2.vrtScl(0.2.mul(pltScl))//.hrzScl(pltScl)

	// hrz1.plt(plt.newFltsSctrDist().asn(hrz1plt3Opn))
	// hrz1plt3Opn.flts(m9Clr 2 m9AlmaSlpAlmaAtOpn)
	// hrz1plt3Opn.vrtScl(0.2.mul(pltScl))//.hrzScl(pltScl)
	// hrz1.plt(plt.newFltsSctrDist().asn(hrz1plt4Opn))
	// hrz1plt4Opn.flts(m20Clr 2 m20SmaSlpAlmaAtOpn)
	// hrz1plt4Opn.vrtScl(0.2.mul(pltScl))//.hrzScl(pltScl)
	// hrz1.plt(plt.newFltsSctrDist().asn(hrz1plt5Opn))
	// hrz1plt5Opn.flts(m50Clr 2 m50SmaSlpAlmaAtOpn)
	// hrz1plt5Opn.vrtScl(0.2.mul(pltScl))//.hrzScl(pltScl)
	// hrz1.plt(plt.newFltsSctrDist().asn(hrz1plt6Opn))
	// hrz1plt6Opn.flts(allClr 2 m9AlmaSlpAlmaAtOpn.cpy().mrg(m20SmaSlpAlmaAtOpn m50SmaSlpAlmaAtOpn))
	// hrz1plt6Opn.vrtScl(0.2.mul(pltScl))//.hrzScl(pltScl)

	// hrz1.plt(plt.newFltsSctrDist().asn(hrz1plt3Cls))
	// hrz1plt3Cls.flts(m9Clr 2 m9AlmaSlpAlmaAtCls)
	// hrz1plt3Cls.vrtScl(0.2.mul(pltScl))//.hrzScl(pltScl)
	// hrz1.plt(plt.newFltsSctrDist().asn(hrz1plt4Cls))
	// hrz1plt4Cls.flts(m20Clr 2 m20SmaSlpAlmaAtCls)
	// hrz1plt4Cls.vrtScl(0.2.mul(pltScl))//.hrzScl(pltScl)
	// hrz1.plt(plt.newFltsSctrDist().asn(hrz1plt5Cls))
	// hrz1plt5Cls.flts(m50Clr 2 m50SmaSlpAlmaAtCls)
	// hrz1plt5Cls.vrtScl(0.2.mul(pltScl))//.hrzScl(pltScl)
	// hrz1.plt(plt.newFltsSctrDist().asn(hrz1plt6Cls))
	// hrz1plt6Cls.flts(allClr 2 m9AlmaSlpAlmaAtCls.cpy().mrg(m20SmaSlpAlmaAtCls m50SmaSlpAlmaAtCls))
	// hrz1plt6Cls.vrtScl(0.2.mul(pltScl))//.hrzScl(pltScl)


	// vrt1.scl(2.0)
	vrt1.sho()

) // end plt each