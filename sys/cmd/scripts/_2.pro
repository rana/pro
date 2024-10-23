

// [1m	2m 3m 5m 8m 13m	21m	34m	55m	89m	144m 233m 377m].asn(durs)
[55m].asn(durs)

hst.oan().audUsd().mktWeeks().asn(mktWeeks)
hst.oan().audUsd().mktDays().asn(mktDays)

hst.oan().audUsd(mktDays.at(0)).asn(instr)
instr.is(durs).ask().asn(ask)
instr.i(1m).ask().asn(m1Ask)
instr.i(15m).ask().asn(m15Ask)
m1Ask.alma(6 0.85).asn(m1Alma)
m15Ask.sma().asn(m15Sma)
m1Alma.otrLss(0 m15Sma).seq(1s m1Alma.otrGtr(0 m15Sma)).asn(opnCnd)
opnCnd.long(10.0 10.0 15m instr).asn(stgy)
stgy.port().prfm().asn(prfm)
prfm.log()


hst.newPrcp().asn(prcp)
prcp.fbr(ask.std())
prcp.splt(prfm.port().splt(0.0)).asn(prcpSplt)
plt.newPrcpSplt(prcpSplt).asn(pltPrcpSplt)
plt.newStgy().asn(pltStgy)
pltStgy.x().vis(fls)
pltStgy.stgy(stgy)
pltStgy.stm(instr.i(durs.fst()).ask().std())
plt.newHrz().asn(pltHrz)
pltHrz.plt(pltPrcpSplt)
pltHrz.plt(pltStgy)
pltHrz.sho()

// prcpSplt.tuneSacf(3 4 0).asn(stgySacf0)
// stgySacf0.port().prfm().asn(tunedPrfm1)
prcpSplt.tuneSacfTil(0.0 2 3 4 1).asn(tunedPrfm1)
tunedPrfm1.log("tunedPrfm1")

hst.newPrcp().asn(prcp1)
prcp1.fbr(ask.rsi())
prcp1.fbr(ask.std())
prcp1.splt(tunedPrfm1.port().splt(0.0)).asn(prcpSplt1)
plt.newPrcpSplt(prcpSplt1).asn(pltPrcpSplt1)
plt.newStgy().asn(pltStgy1)
pltStgy1.x().vis(fls)
pltStgy1.stgy(tunedPrfm1.port().stgys().at(0))
pltStgy1.stm(instr.i(durs.fst()).ask().std())
plt.newHrz().asn(pltHrz1)
pltHrz1.plt(pltPrcpSplt1)
pltHrz1.plt(pltStgy1)
pltHrz1.sho()

hst.newPrcp().asn(prcp2)
hst.oan().audUsd(0s-1s).asn(instr2)
instr2.i(15m).asn(instr2m15)
instr2m15.ask().asn(instr2m15Ask)
prcp2.stm(instr2m15Ask.rsi())
prcp2.stm(instr2m15Ask.std())
opnCnd.longRlng(
	10.0 10.0 15m instr
	mktWeeks
	2 1
	prcp2
	0.0 2
	3 4 1
)