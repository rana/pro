tme.now().asn(start)
start.log("START")

[1m	2m 3m 5m 8m 13m	21m	34m	55m	89m	144m 233m 377m].asn(durs) // 1h29m|2h24m|3h53m|6h17m|10h10m
unts.addsLeq(1 200 2).asn(offs)
unts.addsLeq(1 200 2).asn(lens)
flts.addsLeq(10.0 40.0 5.0).asn(prfLims)
flts.addsLeq(10.0 40.0 5.0).asn(losLims)
tmes.addsLeq(4h 4h 15m).asn(durLims)

hst.oan().eurUsd(0s-0s).asn(instr)
instr.mktWeeks().asn(mktWeeks)
instr.mktDays().asn(mktDays)
instr.mktHrs().asn(mktHrs)
mktDays.asn(mktRngs)
mktRngs.cnt().log("mkt rng cnt")

2.asn(stmCntLim).log("stmCntLim")
2.asn(len).log("len")
hst.newPrcp().asn(prcp)

instr.is(durs).ask().asn(ask).log("ask")
// prcp.fbr(ask.wrsi2().log("wrsi2"))
prcp.fbr(ask.rsi().log("rsi"))
// prcp.fbr(ask.std().log("std"))
instr.i(1m).ask().alma(6 0.85).asn(m1Alma).log("m1Alma")
instr.i(15m).ask().sma().asn(m15Sma).log("m15Sma")
m1Alma.otrLss(0 m15Sma).seq(1s m1Alma.otrGtr(0 m15Sma)).asn(opnCnd).log("opnCnd")
2.asn(stmCntLim).log("stmCntLim")
2.asn(bckTrnLen).log("bckTrnLen")
3.asn(trimItrLim)
4.asn(trimMin)
1.asn(trimForgiveLim)
mktRngs.from(0).to(2).asn(bckRngs)
mktRngs.at(2).asn(fwdRng)
opnCnd.longRlng(
	10.0 10.0 15m instr
	bckRngs
	bckTrnLen
	prcp
	0.0 stmCntLim
	trimItrLim trimMin trimForgiveLim
).port().rng(fwdRng).prfm().log("fwdPrfm")
fwdRng.log("fwdRng")
tme.now().sub(start).log("ELLAPSED")
tme.now().log("END")