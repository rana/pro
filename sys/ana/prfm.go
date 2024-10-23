package ana

import (
	"sys/bsc/flt"
	"sys/bsc/str"
	"sys/bsc/tme"
	"sys/bsc/unt"
)

func CalcPrfm(trds *Trds, dayCnt unt.Unt, losLimMax flt.Flt, durLimMax tme.Tme, pth string) (r *Prfm) {
	r = &Prfm{}
	if trds.Cnt() == 0 {
		return r
	}
	durs := trds.Durs()
	pips := trds.Pips()
	r.TrdPct = trds.Fst().TrdPct
	r.TrdCnt = trds.Cnt()
	r.DayCnt = dayCnt
	r.LosLimMax = losLimMax
	r.DurLimMax = durLimMax
	r.Pth = str.Str(pth)
	r.PipAvg = pips.Sma().Trnc(2)
	r.PipMdn = pips.Mdn().Trnc(1)
	r.PipMin = pips.Min().Trnc(1)
	r.PipMax = pips.Max().Trnc(1)
	r.PipSum = pips.Sum().Trnc(1)
	r.DurAvg = durs.Sma()
	r.DurMdn = durs.Mdn()
	r.DurMin = durs.Min()
	r.DurMax = durs.Max()
	dayCntf, trdCntf, scsCnt := flt.Flt(dayCnt), flt.Flt(r.TrdCnt), pips.CntGtr(0.0) // CORRECT DEF OF SCS?
	r.OpnPerDay = trdCntf.Div(dayCntf).Trnc(1)
	r.PipPerDay = trds.Pips().Sum().Div(dayCntf).Trnc(1)
	r.ScsPerDay = scsCnt.Div(dayCntf).Trnc(1)
	r.ScsPct = scsCnt.Div(trdCntf).Mul(100).Trnc(1)
	for _, trd := range *trds {
		r.CstComUsd += trd.CstComUsd
		r.CstSpdUsd += trd.CstOpnSpdUsd + trd.CstClsSpdUsd
	}
	r.PnlPct = trds.PnlPcts().Sum()
	r.PnlUsd = trds.PnlUsds().Sum()
	r.UsdPerDay = r.PnlUsd.Div(dayCntf).Trnc(0)
	r.CstTotUsd = r.CstSpdUsd.Add(r.CstComUsd).Trnc(0)
	r.CstSpdUsd = r.CstSpdUsd.Trnc(0)
	r.CstComUsd = r.CstComUsd.Trnc(0)
	return r
}
