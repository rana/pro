package ana

import (
	"strings"
	"sys/bsc/flt"
	"sys/bsc/str"
	"sys/bsc/tme"
	"sys/bsc/unt"
)

type (
	Prfm struct {
		PnlPct    flt.Flt
		ScsPct    flt.Flt
		PipPerDay flt.Flt
		UsdPerDay flt.Flt
		ScsPerDay flt.Flt
		OpnPerDay flt.Flt
		PnlUsd    flt.Flt
		PipAvg    flt.Flt
		PipMdn    flt.Flt
		PipMin    flt.Flt
		PipMax    flt.Flt
		PipSum    flt.Flt
		DurAvg    tme.Tme
		DurMdn    tme.Tme
		DurMin    tme.Tme
		DurMax    tme.Tme
		LosLimMax flt.Flt
		DurLimMax tme.Tme
		DayCnt    unt.Unt
		TrdCnt    unt.Unt
		TrdPct    flt.Flt
		CstTotUsd flt.Flt
		CstSpdUsd flt.Flt
		CstComUsd flt.Flt
		Pth       str.Str
	}
	PrfmScp struct {
		Idx uint32
		Arr []*Prfm
	}
)

func (x *Prfm) Dlt(v *Prfm) (r *PrfmDlt) {
	r = &PrfmDlt{}
	r.PnlPctA = x.PnlPct
	r.PnlPctB = v.PnlPct
	r.PnlPctDlt = r.PnlPctB.Sub(r.PnlPctA).Trnc(1)
	r.ScsPctA = x.ScsPct
	r.ScsPctB = v.ScsPct
	r.ScsPctDlt = r.ScsPctB.Sub(r.ScsPctA).Trnc(1)
	r.PipPerDayA = x.PipPerDay
	r.PipPerDayB = v.PipPerDay
	r.PipPerDayDlt = r.PipPerDayB.Sub(r.PipPerDayA).Trnc(1)
	r.UsdPerDayA = x.UsdPerDay
	r.UsdPerDayB = v.UsdPerDay
	r.UsdPerDayDlt = r.UsdPerDayB.Sub(r.UsdPerDayA).Trnc(1)
	r.ScsPerDayA = x.ScsPerDay
	r.ScsPerDayB = v.ScsPerDay
	r.ScsPerDayDlt = r.ScsPerDayB.Sub(r.ScsPerDayA).Trnc(1)
	r.OpnPerDayA = x.OpnPerDay
	r.OpnPerDayB = v.OpnPerDay
	r.OpnPerDayDlt = r.OpnPerDayB.Sub(r.OpnPerDayA).Trnc(1)
	r.PnlUsdA = x.PnlUsd
	r.PnlUsdB = v.PnlUsd
	r.PnlUsdDlt = r.PnlUsdB.Sub(r.PnlUsdA).Trnc(1)
	r.PipAvgA = x.PipAvg
	r.PipAvgB = v.PipAvg
	r.PipAvgDlt = r.PipAvgB.Sub(r.PipAvgA).Trnc(1)
	r.PipMdnA = x.PipMdn
	r.PipMdnB = v.PipMdn
	r.PipMdnDlt = r.PipMdnB.Sub(r.PipMdnA).Trnc(1)
	r.PipMinA = x.PipMin
	r.PipMinB = v.PipMin
	r.PipMinDlt = r.PipMinB.Sub(r.PipMinA).Trnc(1)
	r.PipMaxA = x.PipMax
	r.PipMaxB = v.PipMax
	r.PipMaxDlt = r.PipMaxB.Sub(r.PipMaxA).Trnc(1)
	r.PipSumA = x.PipSum
	r.PipSumB = v.PipSum
	r.PipSumDlt = r.PipSumB.Sub(r.PipSumA).Trnc(1)
	r.DurAvgA = x.DurAvg
	r.DurAvgB = v.DurAvg
	r.DurAvgDlt = flt.Flt(r.DurAvgB).Sub(flt.Flt(r.DurAvgA))
	r.DurMdnA = x.DurMdn
	r.DurMdnB = v.DurMdn
	r.DurMdnDlt = flt.Flt(r.DurMdnB).Sub(flt.Flt(r.DurMdnA))
	r.DurMinA = x.DurMin
	r.DurMinB = v.DurMin
	r.DurMinDlt = flt.Flt(r.DurMinB).Sub(flt.Flt(r.DurMinA))
	r.DurMaxA = x.DurMax
	r.DurMaxB = v.DurMax
	r.DurMaxDlt = flt.Flt(r.DurMaxB).Sub(flt.Flt(r.DurMaxA))
	r.TrdCntA = x.TrdCnt
	r.TrdCntB = v.TrdCnt
	r.TrdCntDlt = flt.Flt(r.TrdCntB).Sub(flt.Flt(r.TrdCntA))
	r.TrdPctA = x.TrdPct
	r.TrdPctB = v.TrdPct
	r.TrdPctDlt = r.TrdPctB.Sub(r.TrdPctA).Trnc(1)
	r.PthB = v.Pth
	return r
}
func (x *Prfm) StrWrt(b *strings.Builder) string {
	b.WriteString("ana.prfm(")
	b.WriteString("pnlPct:")
	x.PnlPct.StrWrt(b)
	b.WriteString(" scsPct:")
	x.ScsPct.StrWrt(b)
	b.WriteString(" pipPerDay:")
	x.PipPerDay.StrWrt(b)
	b.WriteString(" usdPerDay:")
	x.UsdPerDay.StrWrt(b)
	b.WriteString(" scsPerDay:")
	x.ScsPerDay.StrWrt(b)
	b.WriteString(" opnPerDay:")
	x.OpnPerDay.StrWrt(b)
	b.WriteString(" pnlUsd:")
	x.PnlUsd.StrWrt(b)
	b.WriteString(" pipAvg:")
	x.PipAvg.StrWrt(b)
	b.WriteString(" pipMdn:")
	x.PipMdn.StrWrt(b)
	b.WriteString(" pipMin:")
	x.PipMin.StrWrt(b)
	b.WriteString(" pipMax:")
	x.PipMax.StrWrt(b)
	b.WriteString(" pipSum:")
	x.PipSum.StrWrt(b)
	b.WriteString(" durAvg:")
	x.DurAvg.StrWrt(b)
	b.WriteString(" durMdn:")
	x.DurMdn.StrWrt(b)
	b.WriteString(" durMin:")
	x.DurMin.StrWrt(b)
	b.WriteString(" durMax:")
	x.DurMax.StrWrt(b)
	b.WriteString(" losLimMax:")
	x.LosLimMax.StrWrt(b)
	b.WriteString(" durLimMax:")
	x.DurLimMax.StrWrt(b)
	b.WriteString(" dayCnt:")
	x.DayCnt.StrWrt(b)
	b.WriteString(" trdCnt:")
	x.TrdCnt.StrWrt(b)
	b.WriteString(" trdPct:")
	x.TrdPct.StrWrt(b)
	b.WriteString(" cstTotUsd:")
	x.CstTotUsd.StrWrt(b)
	b.WriteString(" cstSpdUsd:")
	x.CstSpdUsd.StrWrt(b)
	b.WriteString(" cstComUsd:")
	x.CstComUsd.StrWrt(b)
	b.WriteString(" pth:")
	x.Pth.StrWrt(b)
	b.WriteRune(')')
	return b.String()
}
func (x *Prfm) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
