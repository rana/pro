2018-5-21/7h.asn(start)   2018-5-21/10h.asn(end)

## plt 1
hst.oan().eurUsd().asn(eurUsd)
hst.oan().eurUsd(start end).s(1).bids().asn(s1)
hst.oan().eurUsd(start.sub(10m) end).m(10).bids().asn(m10)
hst.oan().eurUsd(start.sub(20m) end).m(20).bids().asn(m20)
m10.alma(6 0.85).asn(m10Alma)
m20.sma().asn(m20Sma)
m20.std().sclMul(2.0).asn(m20Blngr)
m20Sma.otrSub(0 m20Blngr).asn(m20BlngrBtm)
m20Sma.otrAdd(0 m20Blngr).asn(m20BlngrTop)

## plt 2
m20.pro().asn(m20Pro)

## plt 3
m10Alma.inrSlp(0).sclMul(eurUsd.slpScl()).asn(m10AlmaSlp)
m20Sma.inrSlp(0).sclMul(eurUsd.slpScl()).asn(m20SmaSlp)

plt.newVrt().asn(vrt1)
vrt1.plt(plt.newStm().asn(plt1))
plt1.stm(s1.lst() clr.grey800)
plt1.stm(m10Alma clr.indigo300)
plt1.stm(m20Sma clr.red300)
plt1.stm(m20BlngrBtm clr.lime700)
plt1.stm(m20BlngrTop clr.lime700)
plt1.vrtScl(0.5)
plt1.x().vis(tru)

vrt1.plt(plt.newStm().asn(plt2))
plt2.stm(m20Pro clr.orange300)
plt2.hrzLn(0.5 clr.grey500 2)
plt2.vrtScl(0.25)

vrt1.plt(plt.newStm().asn(plt3))
plt3.stm(m10AlmaSlp clr.indigo300)
plt3.stm(m20SmaSlp clr.red300)
plt3.vrtScl(0.25)

#vrt1.plt(plt.newHrz().asn(hrz1))
#hrz1.plt(plt.newCndStm().asn(sctr1))
#sctr1.cndStm(clr.indigo300 )

vrt1.sho()