#2018-5-1/12h.asn(from) 
#2018-5-1/12h30m.asn(to)

plt.newVrtRng().asn(vrt1)
#vrt1.hrzSclVal(5m)
#vrt1.hrzRng(2018-5-1/12h 2018-5-1/12h30m)

# BOLLINGER http://stockcharts.com/school/doku.php?id=chart_school:technical_indicators:bollinger_bands
# Middle Band = 20-day simple moving average (SMA)
# Upper Band = 20-day SMA + (20-day standard deviation of price x 2) 
# Lower Band = 20-day SMA - (20-day standard deviation of price x 2)
vrt1.rng().asn(pltRng1)
#pltRng1.vrtScl(6.0)
#pltRng1.vrtSclVal(0.00001)
ana.oan().hst().gbpUsd().asn(instr1)
instr1.s1().asks().lst().asn(stm1)
pltRng1.stm(stm1 clr.white)
5.asn(smaLen)
stm1.aggSma(smaLen).asn(stm1Sma)
pltRng1.stm(stm1Sma clr.cyan700)
stm1.aggStd(smaLen).sclMul(2.0).asn(blngr)
stm1Sma.otrSub(0 blngr).asn(blngrBtm)
stm1Sma.otrAdd(0 blngr).asn(blngrTop)
pltRng1.stmBnd(blngrBtm blngrTop clr.hex("#0f0f0f") clr.deepPurpleA400)
stm1.otrLss(0 blngrTop).asn(stm1Lss)
stm1.otrGtr(0 blngrTop).asn(stm1Gtr)
stm1Lss.seq(2s stm1Gtr).asn(blngrTopCrs)
pltRng1.cnd(blngrTopCrs clr.orangeA400)

# PCTB http://stockcharts.com/school/doku.php?id=chart_school:technical_indicators:bollinger_band_perce
# %B = (Price - Lower Band)/(Upper Band - Lower Band)
stm1.otrSub(0 blngrBtm).asn(numB)
blngrTop.otrSub(0 blngrBtm).asn(denomB)
numB.otrDiv(0 denomB).asn(pctB)
pctB.aggSma(30).asn(pctBSma)
vrt1.rng().stm(pctB clr.green400).vrtScl(0.25).asn(pltPctb1)
pltPctb1.stm(pctBSma clr.cyan700)
pltPctb1.hrzBnd(0.0 1.0 clr.hex("#0f0f0f") clr.deepPurpleA400)
pltPctb1.hrzLn(0.5 clr.deepPurple700)
pltPctb1.cnd(blngrTopCrs clr.orangeA400)

#stm1.sclRem(0.0010).sclMul(instr1.pipScl()).asn(rem10)
#vrt1.rng().stm(rem10 clr.orange100).vrtScl(0.25).asn(pltRem10)
#pltRem10.hrzLn(0.0 clr.deepPurple700)
#pltRem10.cnd(blngrTopCrs clr.orangeA400)

#stm1.inrSub(1).sclMul(instr1.pipScl()).asn(stm1Sub1)
#vrt1.rng().stm(stm1Sub1 clr.green700).vrtScl(0.25).asn(pltSub1)
#pltSub1.hrzLn(0.0 clr.deepPurple700)
#pltSub1.cnd(blngrTopCrs clr.orangeA400)

# vrt1.sho()


# (revPerDay:)[-+]?[0-9]*\.?[0-9]*
blngrTopCrs.long(1.0 1.0 2h instr1).logPerf()
blngrTopCrs.long(2.0 2.0 2h instr1).logPerf()
blngrTopCrs.long(4.0 4.0 2h instr1).logPerf()
blngrTopCrs.long(6.0 6.0 2h instr1).logPerf()
blngrTopCrs.long(7.0 7.0 2h instr1).logPerf()
blngrTopCrs.long(8.0 8.0 2h instr1).logPerf()
blngrTopCrs.long(9.0 9.0 2h instr1).logPerf()
blngrTopCrs.long(10.0 10.0 2h instr1).logPerf()
blngrTopCrs.long(11.0 11.0 2h instr1).logPerf()
blngrTopCrs.long(15.0 15.0 2h instr1).logPerf()




