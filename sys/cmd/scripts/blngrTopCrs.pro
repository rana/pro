# BOLLINGER http://stockcharts.com/school/doku.php?id=chart_school:technical_indicators:bollinger_bands
# Middle Band = 20-day simple moving average (SMA)
# Upper Band = 20-day SMA + (20-day standard deviation of price x 2) 
# Lower Band = 20-day SMA - (20-day standard deviation of price x 2)
ana.oan().hst().gbpUsd().asn(instr1)
instr1.s1().asks().lst().asn(stm1)
5.asn(smaLen)
stm1.aggSma(smaLen).asn(stm1Sma)
pltRng1.stm(stm1Sma clr.cyan700)
stm1.aggStd(smaLen).sclMul(2.0).asn(blngr)
stm1Sma.otrSub(0 blngr).asn(blngrBtm)
stm1Sma.otrAdd(0 blngr).asn(blngrTop)
stm1.otrLss(0 blngrTop).asn(stm1Lss)
stm1.otrGtr(0 blngrTop).asn(stm1Gtr)
stm1Lss.seq(2s stm1Gtr).asn(blngrTopCrs)


# PCTB http://stockcharts.com/school/doku.php?id=chart_school:technical_indicators:bollinger_band_perce
# %B = (Price - Lower Band)/(Upper Band - Lower Band)
stm1.otrSub(0 blngrBtm).asn(numB)
blngrTop.otrSub(0 blngrBtm).asn(denomB)
numB.otrDiv(0 denomB).asn(pctB)
pctB.aggSma(30).asn(pctBSma)

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



