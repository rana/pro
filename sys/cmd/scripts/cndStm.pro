# 2018-5-21  2018-5-21/1h

plt.newVrtRng().asn(vrt1)
vrt1.rng().asn(pltRng1)

[20].asn(aggLens)

hst.oan().eurUsd(2018-5-21/12h  2018-5-21/13h).s1().asks().asn(eurUsdBse)
eurUsdBse.lst().asn(eurUsd)
eurUsdBse.sar(0.005 0.2).asn(sar)
sar.otrGtr(0 eurUsd).asn(sarTop)
sar.otrLss(0 eurUsd).asn(sarBtm)
sarTop.seq(1s sarBtm).asn(sarDwn)
sarBtm.seq(1s sarTop).asn(sarUp)

pltRng1.stm(eurUsd clr.purple400 1)
pltRng1.stm(sar clr.white)

aggLens.each(aggLen
	eurUsd.aggAlma(aggLen 6 0.85).asn(alma)
	eurUsd.aggSma(aggLen).asn(sma)
	eurUsd.aggStd(aggLen).asn(std)
	std.sclMul(2.0).asn(blngr)
	sma.otrSub(0 blngr).asn(blngrBtm)
	sma.otrAdd(0 blngr).asn(blngrTop)

	pltRng1.stm(alma clr.orange700)
	pltRng1.stm(sma clr.indigo700)
	
	pltRng1.stm(blngrBtm clr.lime700)
	pltRng1.stm(blngrTop clr.lime700)
	pltRng1.cnd(sarUp clr.teal700)
)

pltRng1.scl(5.0)
vrt1.sho()