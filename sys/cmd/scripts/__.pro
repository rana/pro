hst.oan().eurUsd(0s-0s).mktHrs().asn(mktHrs)
hst.oan().eurUsd(0s-0s).mktDays().asn(mktDays)
hst.oan().eurUsd(mktHrs.at(1)).asn(instr)
// instr.i(1m).ask().lst().asn(lst)
// lst.aggEma(12).asn(ema12)
// lst.aggEma(26).asn(ema26)
instr.i(12m).ask().ema().asn(ema12)
instr.i(26m).ask().ema().asn(ema26)
ema12.otrSub(0 ema26).asn(macd)
macd.aggEma(9).asn(macdTrg)

plt.newStm().asn(pltStm)
pltStm.stm(pen.indigoA400 macd)
pltStm.stm(pen.redA400 macdTrg)
pltStm.x().vis(tru)
pltStm.sho()