
2018-5-21/12h.asn(start)
2018-5-21/12h30m.asn(end)

plt.newVrtRng(start end).asn(vrt1)
vrt1.rng().asn(pltRng1)

9.asn(aggLen)

hst.oan().eurUsd(start end).s1().asks().lst().asn(eurUsd)
eurUsd.aggAlma(aggLen 6 0.85).asn(alma)
eurUsd.aggSma(aggLen).asn(sma)
eurUsd.aggStd(aggLen).asn(std)
std.sclMul(2.0).asn(blngr)
sma.otrSub(0 blngr).asn(blngrBtm)
sma.otrAdd(0 blngr).asn(blngrTop)

alma.otrLss(0 sma).asn(almaLssSma)
alma.otrGtr(0 sma).asn(almaGtrSma)
almaLssSma.seq(1s almaGtrSma).asn(almaUpSma)

pltRng1.stm(eurUsd clr.purple400 1)
pltRng1.stm(alma clr.orange700)
pltRng1.stm(sma clr.indigo700)
pltRng1.stm(blngrBtm clr.lime700)
pltRng1.stm(blngrTop clr.lime700)
pltRng1.cnd(almaUpSma clr.teal700)


pltRng1.scl(3.0)
vrt1.sho()