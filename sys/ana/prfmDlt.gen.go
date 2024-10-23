package ana

import (
	"bytes"
	"strings"
	"sys/bsc/flt"
	"sys/bsc/str"
	"sys/bsc/tme"
	"sys/bsc/unt"
)

type (
	PrfmDlt struct {
		PnlPctA      flt.Flt
		PnlPctB      flt.Flt
		PnlPctDlt    flt.Flt
		ScsPctA      flt.Flt
		ScsPctB      flt.Flt
		ScsPctDlt    flt.Flt
		PipPerDayA   flt.Flt
		PipPerDayB   flt.Flt
		PipPerDayDlt flt.Flt
		UsdPerDayA   flt.Flt
		UsdPerDayB   flt.Flt
		UsdPerDayDlt flt.Flt
		ScsPerDayA   flt.Flt
		ScsPerDayB   flt.Flt
		ScsPerDayDlt flt.Flt
		OpnPerDayA   flt.Flt
		OpnPerDayB   flt.Flt
		OpnPerDayDlt flt.Flt
		PnlUsdA      flt.Flt
		PnlUsdB      flt.Flt
		PnlUsdDlt    flt.Flt
		PipAvgA      flt.Flt
		PipAvgB      flt.Flt
		PipAvgDlt    flt.Flt
		PipMdnA      flt.Flt
		PipMdnB      flt.Flt
		PipMdnDlt    flt.Flt
		PipMinA      flt.Flt
		PipMinB      flt.Flt
		PipMinDlt    flt.Flt
		PipMaxA      flt.Flt
		PipMaxB      flt.Flt
		PipMaxDlt    flt.Flt
		PipSumA      flt.Flt
		PipSumB      flt.Flt
		PipSumDlt    flt.Flt
		DurAvgA      tme.Tme
		DurAvgB      tme.Tme
		DurAvgDlt    flt.Flt
		DurMdnA      tme.Tme
		DurMdnB      tme.Tme
		DurMdnDlt    flt.Flt
		DurMinA      tme.Tme
		DurMinB      tme.Tme
		DurMinDlt    flt.Flt
		DurMaxA      tme.Tme
		DurMaxB      tme.Tme
		DurMaxDlt    flt.Flt
		TrdCntA      unt.Unt
		TrdCntB      unt.Unt
		TrdCntDlt    flt.Flt
		TrdPctA      flt.Flt
		TrdPctB      flt.Flt
		TrdPctDlt    flt.Flt
		PthB         str.Str
	}
	PrfmDltScp struct {
		Idx uint32
		Arr []*PrfmDlt
	}
)

func (x *PrfmDlt) StrWrt(b *strings.Builder) string {
	b.WriteString("ana.prfmDlt(")
	b.WriteString("pnlPctA:")
	x.PnlPctA.StrWrt(b)
	b.WriteString(" pnlPctB:")
	x.PnlPctB.StrWrt(b)
	b.WriteString(" pnlPctDlt:")
	x.PnlPctDlt.StrWrt(b)
	b.WriteString(" scsPctA:")
	x.ScsPctA.StrWrt(b)
	b.WriteString(" scsPctB:")
	x.ScsPctB.StrWrt(b)
	b.WriteString(" scsPctDlt:")
	x.ScsPctDlt.StrWrt(b)
	b.WriteString(" pipPerDayA:")
	x.PipPerDayA.StrWrt(b)
	b.WriteString(" pipPerDayB:")
	x.PipPerDayB.StrWrt(b)
	b.WriteString(" pipPerDayDlt:")
	x.PipPerDayDlt.StrWrt(b)
	b.WriteString(" usdPerDayA:")
	x.UsdPerDayA.StrWrt(b)
	b.WriteString(" usdPerDayB:")
	x.UsdPerDayB.StrWrt(b)
	b.WriteString(" usdPerDayDlt:")
	x.UsdPerDayDlt.StrWrt(b)
	b.WriteString(" scsPerDayA:")
	x.ScsPerDayA.StrWrt(b)
	b.WriteString(" scsPerDayB:")
	x.ScsPerDayB.StrWrt(b)
	b.WriteString(" scsPerDayDlt:")
	x.ScsPerDayDlt.StrWrt(b)
	b.WriteString(" opnPerDayA:")
	x.OpnPerDayA.StrWrt(b)
	b.WriteString(" opnPerDayB:")
	x.OpnPerDayB.StrWrt(b)
	b.WriteString(" opnPerDayDlt:")
	x.OpnPerDayDlt.StrWrt(b)
	b.WriteString(" pnlUsdA:")
	x.PnlUsdA.StrWrt(b)
	b.WriteString(" pnlUsdB:")
	x.PnlUsdB.StrWrt(b)
	b.WriteString(" pnlUsdDlt:")
	x.PnlUsdDlt.StrWrt(b)
	b.WriteString(" pipAvgA:")
	x.PipAvgA.StrWrt(b)
	b.WriteString(" pipAvgB:")
	x.PipAvgB.StrWrt(b)
	b.WriteString(" pipAvgDlt:")
	x.PipAvgDlt.StrWrt(b)
	b.WriteString(" pipMdnA:")
	x.PipMdnA.StrWrt(b)
	b.WriteString(" pipMdnB:")
	x.PipMdnB.StrWrt(b)
	b.WriteString(" pipMdnDlt:")
	x.PipMdnDlt.StrWrt(b)
	b.WriteString(" pipMinA:")
	x.PipMinA.StrWrt(b)
	b.WriteString(" pipMinB:")
	x.PipMinB.StrWrt(b)
	b.WriteString(" pipMinDlt:")
	x.PipMinDlt.StrWrt(b)
	b.WriteString(" pipMaxA:")
	x.PipMaxA.StrWrt(b)
	b.WriteString(" pipMaxB:")
	x.PipMaxB.StrWrt(b)
	b.WriteString(" pipMaxDlt:")
	x.PipMaxDlt.StrWrt(b)
	b.WriteString(" pipSumA:")
	x.PipSumA.StrWrt(b)
	b.WriteString(" pipSumB:")
	x.PipSumB.StrWrt(b)
	b.WriteString(" pipSumDlt:")
	x.PipSumDlt.StrWrt(b)
	b.WriteString(" durAvgA:")
	x.DurAvgA.StrWrt(b)
	b.WriteString(" durAvgB:")
	x.DurAvgB.StrWrt(b)
	b.WriteString(" durAvgDlt:")
	x.DurAvgDlt.StrWrt(b)
	b.WriteString(" durMdnA:")
	x.DurMdnA.StrWrt(b)
	b.WriteString(" durMdnB:")
	x.DurMdnB.StrWrt(b)
	b.WriteString(" durMdnDlt:")
	x.DurMdnDlt.StrWrt(b)
	b.WriteString(" durMinA:")
	x.DurMinA.StrWrt(b)
	b.WriteString(" durMinB:")
	x.DurMinB.StrWrt(b)
	b.WriteString(" durMinDlt:")
	x.DurMinDlt.StrWrt(b)
	b.WriteString(" durMaxA:")
	x.DurMaxA.StrWrt(b)
	b.WriteString(" durMaxB:")
	x.DurMaxB.StrWrt(b)
	b.WriteString(" durMaxDlt:")
	x.DurMaxDlt.StrWrt(b)
	b.WriteString(" trdCntA:")
	x.TrdCntA.StrWrt(b)
	b.WriteString(" trdCntB:")
	x.TrdCntB.StrWrt(b)
	b.WriteString(" trdCntDlt:")
	x.TrdCntDlt.StrWrt(b)
	b.WriteString(" trdPctA:")
	x.TrdPctA.StrWrt(b)
	b.WriteString(" trdPctB:")
	x.TrdPctB.StrWrt(b)
	b.WriteString(" trdPctDlt:")
	x.TrdPctDlt.StrWrt(b)
	b.WriteString(" pthB:")
	x.PthB.StrWrt(b)
	b.WriteRune(')')
	return b.String()
}
func (x *PrfmDlt) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *PrfmDlt) BytWrt(b *bytes.Buffer) {
	x.PnlPctA.BytWrt(b)
	x.PnlPctB.BytWrt(b)
	x.PnlPctDlt.BytWrt(b)
	x.ScsPctA.BytWrt(b)
	x.ScsPctB.BytWrt(b)
	x.ScsPctDlt.BytWrt(b)
	x.PipPerDayA.BytWrt(b)
	x.PipPerDayB.BytWrt(b)
	x.PipPerDayDlt.BytWrt(b)
	x.UsdPerDayA.BytWrt(b)
	x.UsdPerDayB.BytWrt(b)
	x.UsdPerDayDlt.BytWrt(b)
	x.ScsPerDayA.BytWrt(b)
	x.ScsPerDayB.BytWrt(b)
	x.ScsPerDayDlt.BytWrt(b)
	x.OpnPerDayA.BytWrt(b)
	x.OpnPerDayB.BytWrt(b)
	x.OpnPerDayDlt.BytWrt(b)
	x.PnlUsdA.BytWrt(b)
	x.PnlUsdB.BytWrt(b)
	x.PnlUsdDlt.BytWrt(b)
	x.PipAvgA.BytWrt(b)
	x.PipAvgB.BytWrt(b)
	x.PipAvgDlt.BytWrt(b)
	x.PipMdnA.BytWrt(b)
	x.PipMdnB.BytWrt(b)
	x.PipMdnDlt.BytWrt(b)
	x.PipMinA.BytWrt(b)
	x.PipMinB.BytWrt(b)
	x.PipMinDlt.BytWrt(b)
	x.PipMaxA.BytWrt(b)
	x.PipMaxB.BytWrt(b)
	x.PipMaxDlt.BytWrt(b)
	x.PipSumA.BytWrt(b)
	x.PipSumB.BytWrt(b)
	x.PipSumDlt.BytWrt(b)
	x.DurAvgA.BytWrt(b)
	x.DurAvgB.BytWrt(b)
	x.DurAvgDlt.BytWrt(b)
	x.DurMdnA.BytWrt(b)
	x.DurMdnB.BytWrt(b)
	x.DurMdnDlt.BytWrt(b)
	x.DurMinA.BytWrt(b)
	x.DurMinB.BytWrt(b)
	x.DurMinDlt.BytWrt(b)
	x.DurMaxA.BytWrt(b)
	x.DurMaxB.BytWrt(b)
	x.DurMaxDlt.BytWrt(b)
	x.TrdCntA.BytWrt(b)
	x.TrdCntB.BytWrt(b)
	x.TrdCntDlt.BytWrt(b)
	x.TrdPctA.BytWrt(b)
	x.TrdPctB.BytWrt(b)
	x.TrdPctDlt.BytWrt(b)
	x.PthB.BytWrt(b)
}
func (x *PrfmDlt) BytRed(b []byte) (idx int) {
	idx += x.PnlPctA.BytRed(b)
	idx += x.PnlPctB.BytRed(b[idx:])
	idx += x.PnlPctDlt.BytRed(b[idx:])
	idx += x.ScsPctA.BytRed(b[idx:])
	idx += x.ScsPctB.BytRed(b[idx:])
	idx += x.ScsPctDlt.BytRed(b[idx:])
	idx += x.PipPerDayA.BytRed(b[idx:])
	idx += x.PipPerDayB.BytRed(b[idx:])
	idx += x.PipPerDayDlt.BytRed(b[idx:])
	idx += x.UsdPerDayA.BytRed(b[idx:])
	idx += x.UsdPerDayB.BytRed(b[idx:])
	idx += x.UsdPerDayDlt.BytRed(b[idx:])
	idx += x.ScsPerDayA.BytRed(b[idx:])
	idx += x.ScsPerDayB.BytRed(b[idx:])
	idx += x.ScsPerDayDlt.BytRed(b[idx:])
	idx += x.OpnPerDayA.BytRed(b[idx:])
	idx += x.OpnPerDayB.BytRed(b[idx:])
	idx += x.OpnPerDayDlt.BytRed(b[idx:])
	idx += x.PnlUsdA.BytRed(b[idx:])
	idx += x.PnlUsdB.BytRed(b[idx:])
	idx += x.PnlUsdDlt.BytRed(b[idx:])
	idx += x.PipAvgA.BytRed(b[idx:])
	idx += x.PipAvgB.BytRed(b[idx:])
	idx += x.PipAvgDlt.BytRed(b[idx:])
	idx += x.PipMdnA.BytRed(b[idx:])
	idx += x.PipMdnB.BytRed(b[idx:])
	idx += x.PipMdnDlt.BytRed(b[idx:])
	idx += x.PipMinA.BytRed(b[idx:])
	idx += x.PipMinB.BytRed(b[idx:])
	idx += x.PipMinDlt.BytRed(b[idx:])
	idx += x.PipMaxA.BytRed(b[idx:])
	idx += x.PipMaxB.BytRed(b[idx:])
	idx += x.PipMaxDlt.BytRed(b[idx:])
	idx += x.PipSumA.BytRed(b[idx:])
	idx += x.PipSumB.BytRed(b[idx:])
	idx += x.PipSumDlt.BytRed(b[idx:])
	idx += x.DurAvgA.BytRed(b[idx:])
	idx += x.DurAvgB.BytRed(b[idx:])
	idx += x.DurAvgDlt.BytRed(b[idx:])
	idx += x.DurMdnA.BytRed(b[idx:])
	idx += x.DurMdnB.BytRed(b[idx:])
	idx += x.DurMdnDlt.BytRed(b[idx:])
	idx += x.DurMinA.BytRed(b[idx:])
	idx += x.DurMinB.BytRed(b[idx:])
	idx += x.DurMinDlt.BytRed(b[idx:])
	idx += x.DurMaxA.BytRed(b[idx:])
	idx += x.DurMaxB.BytRed(b[idx:])
	idx += x.DurMaxDlt.BytRed(b[idx:])
	idx += x.TrdCntA.BytRed(b[idx:])
	idx += x.TrdCntB.BytRed(b[idx:])
	idx += x.TrdCntDlt.BytRed(b[idx:])
	idx += x.TrdPctA.BytRed(b[idx:])
	idx += x.TrdPctB.BytRed(b[idx:])
	idx += x.TrdPctDlt.BytRed(b[idx:])
	idx += x.PthB.BytRed(b[idx:])
	return idx
}
func (x *PrfmDlt) Bytes() []byte {
	b := &bytes.Buffer{}
	x.BytWrt(b)
	return b.Bytes()
}
