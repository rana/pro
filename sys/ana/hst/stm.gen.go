package hst

import (
	"math"
	"strings"
	"sys"
	"sys/ana"
	"sys/bsc/bnd"
	"sys/bsc/bnds"
	"sys/bsc/flt"
	"sys/bsc/flts"
	"sys/bsc/str"
	"sys/bsc/tmes"
	"sys/bsc/unt"
	"sys/err"
)

type (
	Stm interface {
		Name() str.Str
		PrmWrt(b *strings.Builder)
		Prm() string
		StrWrt(b *strings.Builder)
		String() string
		Bse() *StmBse
		At(ts *tmes.Tmes) (r *flts.Flts)
		Atf(ts *tmes.Tmes) (r []float32)
		UnaPos() Stm
		UnaNeg() Stm
		UnaInv() Stm
		UnaSqr() Stm
		UnaSqrt() Stm
		SclAdd(scl flt.Flt) Stm
		SclSub(scl flt.Flt) Stm
		SclMul(scl flt.Flt) Stm
		SclDiv(scl flt.Flt) Stm
		SclRem(scl flt.Flt) Stm
		SclPow(scl flt.Flt) Stm
		SclMin(scl flt.Flt) Stm
		SclMax(scl flt.Flt) Stm
		SelEql(sel flt.Flt) Stm
		SelNeq(sel flt.Flt) Stm
		SelLss(sel flt.Flt) Stm
		SelGtr(sel flt.Flt) Stm
		SelLeq(sel flt.Flt) Stm
		SelGeq(sel flt.Flt) Stm
		AggFst(length unt.Unt) Stm
		AggLst(length unt.Unt) Stm
		AggSum(length unt.Unt) Stm
		AggPrd(length unt.Unt) Stm
		AggMin(length unt.Unt) Stm
		AggMax(length unt.Unt) Stm
		AggMid(length unt.Unt) Stm
		AggMdn(length unt.Unt) Stm
		AggSma(length unt.Unt) Stm
		AggGma(length unt.Unt) Stm
		AggWma(length unt.Unt) Stm
		AggRsi(length unt.Unt) Stm
		AggWrsi(length unt.Unt) Stm
		AggAlma(length unt.Unt) Stm
		AggVrnc(length unt.Unt) Stm
		AggStd(length unt.Unt) Stm
		AggRngFul(length unt.Unt) Stm
		AggRngLst(length unt.Unt) Stm
		AggProLst(length unt.Unt) Stm
		AggProSma(length unt.Unt) Stm
		AggProAlma(length unt.Unt) Stm
		AggEma(length unt.Unt) Stm
		InrAdd(off unt.Unt) Stm
		InrSub(off unt.Unt) Stm
		InrMul(off unt.Unt) Stm
		InrDiv(off unt.Unt) Stm
		InrRem(off unt.Unt) Stm
		InrPow(off unt.Unt) Stm
		InrMin(off unt.Unt) Stm
		InrMax(off unt.Unt) Stm
		InrSlp(off unt.Unt) Stm
		OtrAdd(off unt.Unt, a Stm) Stm
		OtrSub(off unt.Unt, a Stm) Stm
		OtrMul(off unt.Unt, a Stm) Stm
		OtrDiv(off unt.Unt, a Stm) Stm
		OtrRem(off unt.Unt, a Stm) Stm
		OtrPow(off unt.Unt, a Stm) Stm
		OtrMin(off unt.Unt, a Stm) Stm
		OtrMax(off unt.Unt, a Stm) Stm
		SclEql(scl flt.Flt) Cnd
		SclNeq(scl flt.Flt) Cnd
		SclLss(scl flt.Flt) Cnd
		SclGtr(scl flt.Flt) Cnd
		SclLeq(scl flt.Flt) Cnd
		SclGeq(scl flt.Flt) Cnd
		InrEql(off unt.Unt) Cnd
		InrNeq(off unt.Unt) Cnd
		InrLss(off unt.Unt) Cnd
		InrGtr(off unt.Unt) Cnd
		InrLeq(off unt.Unt) Cnd
		InrGeq(off unt.Unt) Cnd
		OtrEql(off unt.Unt, a Stm) Cnd
		OtrNeq(off unt.Unt, a Stm) Cnd
		OtrLss(off unt.Unt, a Stm) Cnd
		OtrGtr(off unt.Unt, a Stm) Cnd
		OtrLeq(off unt.Unt, a Stm) Cnd
		OtrGeq(off unt.Unt, a Stm) Cnd
	}
	StmBse struct {
		Slf  Stm
		Tmes *tmes.Tmes
		Vals *flts.Flts
	}
	StmSeg struct {
		bnd.Bnd
		Vals *flts.Flts
		Out  *flts.Flts
	}
	StmScp struct {
		Idx uint32
		Arr []Stm
	}
	StmRteFst struct {
		StmBse
		Side Side
	}
	StmRteFstSeg struct {
		StmSeg
		ValBnds *bnds.Bnds
	}
	StmRteLst struct {
		StmBse
		Side Side
	}
	StmRteLstSeg struct {
		StmSeg
		ValBnds *bnds.Bnds
	}
	StmRteSum struct {
		StmBse
		Side Side
	}
	StmRteSumSeg struct {
		StmSeg
		ValBnds *bnds.Bnds
	}
	StmRtePrd struct {
		StmBse
		Side Side
	}
	StmRtePrdSeg struct {
		StmSeg
		ValBnds *bnds.Bnds
	}
	StmRteMin struct {
		StmBse
		Side Side
	}
	StmRteMinSeg struct {
		StmSeg
		ValBnds *bnds.Bnds
	}
	StmRteMax struct {
		StmBse
		Side Side
	}
	StmRteMaxSeg struct {
		StmSeg
		ValBnds *bnds.Bnds
	}
	StmRteMid struct {
		StmBse
		Side Side
	}
	StmRteMidSeg struct {
		StmSeg
		ValBnds *bnds.Bnds
	}
	StmRteMdn struct {
		StmBse
		Side Side
	}
	StmRteMdnSeg struct {
		StmSeg
		ValBnds *bnds.Bnds
	}
	StmRteSma struct {
		StmBse
		Side Side
	}
	StmRteSmaSeg struct {
		StmSeg
		ValBnds *bnds.Bnds
	}
	StmRteGma struct {
		StmBse
		Side Side
	}
	StmRteGmaSeg struct {
		StmSeg
		ValBnds *bnds.Bnds
	}
	StmRteWma struct {
		StmBse
		Side Side
	}
	StmRteWmaSeg struct {
		StmSeg
		ValBnds *bnds.Bnds
	}
	StmRteRsi struct {
		StmBse
		Side Side
	}
	StmRteRsiSeg struct {
		StmSeg
		ValBnds *bnds.Bnds
	}
	StmRteWrsi struct {
		StmBse
		Side Side
	}
	StmRteWrsiSeg struct {
		StmSeg
		ValBnds *bnds.Bnds
	}
	StmRteAlma struct {
		StmBse
		Side Side
	}
	StmRteAlmaSeg struct {
		StmSeg
		ValBnds *bnds.Bnds
	}
	StmRteVrnc struct {
		StmBse
		Side Side
	}
	StmRteVrncSeg struct {
		StmSeg
		ValBnds *bnds.Bnds
	}
	StmRteStd struct {
		StmBse
		Side Side
	}
	StmRteStdSeg struct {
		StmSeg
		ValBnds *bnds.Bnds
	}
	StmRteRngFul struct {
		StmBse
		Side Side
	}
	StmRteRngFulSeg struct {
		StmSeg
		ValBnds *bnds.Bnds
	}
	StmRteRngLst struct {
		StmBse
		Side Side
	}
	StmRteRngLstSeg struct {
		StmSeg
		ValBnds *bnds.Bnds
	}
	StmRteProLst struct {
		StmBse
		Side Side
	}
	StmRteProLstSeg struct {
		StmSeg
		ValBnds *bnds.Bnds
	}
	StmRteProSma struct {
		StmBse
		Side Side
	}
	StmRteProSmaSeg struct {
		StmSeg
		ValBnds *bnds.Bnds
	}
	StmRteProAlma struct {
		StmBse
		Side Side
	}
	StmRteProAlmaSeg struct {
		StmSeg
		ValBnds *bnds.Bnds
	}
	StmRte1Sar struct {
		StmBse
		Side  Side
		AfInc flt.Flt
		AfMax flt.Flt
	}
	StmRteEma struct {
		StmBse
		Side Side
	}
	StmUnaPos struct {
		StmBse
		Stm Stm
	}
	StmUnaPosSeg struct {
		StmSeg
	}
	StmUnaNeg struct {
		StmBse
		Stm Stm
	}
	StmUnaNegSeg struct {
		StmSeg
	}
	StmUnaInv struct {
		StmBse
		Stm Stm
	}
	StmUnaInvSeg struct {
		StmSeg
	}
	StmUnaSqr struct {
		StmBse
		Stm Stm
	}
	StmUnaSqrSeg struct {
		StmSeg
	}
	StmUnaSqrt struct {
		StmBse
		Stm Stm
	}
	StmUnaSqrtSeg struct {
		StmSeg
	}
	StmSclAdd struct {
		StmBse
		Stm Stm
		Scl flt.Flt
	}
	StmSclAddSeg struct {
		StmSeg
		Scl flt.Flt
	}
	StmSclSub struct {
		StmBse
		Stm Stm
		Scl flt.Flt
	}
	StmSclSubSeg struct {
		StmSeg
		Scl flt.Flt
	}
	StmSclMul struct {
		StmBse
		Stm Stm
		Scl flt.Flt
	}
	StmSclMulSeg struct {
		StmSeg
		Scl flt.Flt
	}
	StmSclDiv struct {
		StmBse
		Stm Stm
		Scl flt.Flt
	}
	StmSclDivSeg struct {
		StmSeg
		Scl flt.Flt
	}
	StmSclRem struct {
		StmBse
		Stm Stm
		Scl flt.Flt
	}
	StmSclRemSeg struct {
		StmSeg
		Scl flt.Flt
	}
	StmSclPow struct {
		StmBse
		Stm Stm
		Scl flt.Flt
	}
	StmSclPowSeg struct {
		StmSeg
		Scl flt.Flt
	}
	StmSclMin struct {
		StmBse
		Stm Stm
		Scl flt.Flt
	}
	StmSclMinSeg struct {
		StmSeg
		Scl flt.Flt
	}
	StmSclMax struct {
		StmBse
		Stm Stm
		Scl flt.Flt
	}
	StmSclMaxSeg struct {
		StmSeg
		Scl flt.Flt
	}
	StmSelEql struct {
		StmBse
		Stm Stm
		Sel flt.Flt
	}
	StmSelEqlSeg struct {
		StmSeg
		Sel flt.Flt
	}
	StmSelNeq struct {
		StmBse
		Stm Stm
		Sel flt.Flt
	}
	StmSelNeqSeg struct {
		StmSeg
		Sel flt.Flt
	}
	StmSelLss struct {
		StmBse
		Stm Stm
		Sel flt.Flt
	}
	StmSelLssSeg struct {
		StmSeg
		Sel flt.Flt
	}
	StmSelGtr struct {
		StmBse
		Stm Stm
		Sel flt.Flt
	}
	StmSelGtrSeg struct {
		StmSeg
		Sel flt.Flt
	}
	StmSelLeq struct {
		StmBse
		Stm Stm
		Sel flt.Flt
	}
	StmSelLeqSeg struct {
		StmSeg
		Sel flt.Flt
	}
	StmSelGeq struct {
		StmBse
		Stm Stm
		Sel flt.Flt
	}
	StmSelGeqSeg struct {
		StmSeg
		Sel flt.Flt
	}
	StmAggFst struct {
		StmBse
		Stm    Stm
		Length unt.Unt
	}
	StmAggFstSeg struct {
		StmSeg
		Length unt.Unt
	}
	StmAggLst struct {
		StmBse
		Stm    Stm
		Length unt.Unt
	}
	StmAggLstSeg struct {
		StmSeg
		Length unt.Unt
	}
	StmAggSum struct {
		StmBse
		Stm    Stm
		Length unt.Unt
	}
	StmAggSumSeg struct {
		StmSeg
		Length unt.Unt
	}
	StmAggPrd struct {
		StmBse
		Stm    Stm
		Length unt.Unt
	}
	StmAggPrdSeg struct {
		StmSeg
		Length unt.Unt
	}
	StmAggMin struct {
		StmBse
		Stm    Stm
		Length unt.Unt
	}
	StmAggMinSeg struct {
		StmSeg
		Length unt.Unt
	}
	StmAggMax struct {
		StmBse
		Stm    Stm
		Length unt.Unt
	}
	StmAggMaxSeg struct {
		StmSeg
		Length unt.Unt
	}
	StmAggMid struct {
		StmBse
		Stm    Stm
		Length unt.Unt
	}
	StmAggMidSeg struct {
		StmSeg
		Length unt.Unt
	}
	StmAggMdn struct {
		StmBse
		Stm    Stm
		Length unt.Unt
	}
	StmAggMdnSeg struct {
		StmSeg
		Length unt.Unt
	}
	StmAggSma struct {
		StmBse
		Stm    Stm
		Length unt.Unt
	}
	StmAggSmaSeg struct {
		StmSeg
		Length unt.Unt
	}
	StmAggGma struct {
		StmBse
		Stm    Stm
		Length unt.Unt
	}
	StmAggGmaSeg struct {
		StmSeg
		Length unt.Unt
	}
	StmAggWma struct {
		StmBse
		Stm    Stm
		Length unt.Unt
	}
	StmAggWmaSeg struct {
		StmSeg
		Length unt.Unt
	}
	StmAggRsi struct {
		StmBse
		Stm    Stm
		Length unt.Unt
	}
	StmAggRsiSeg struct {
		StmSeg
		Length unt.Unt
	}
	StmAggWrsi struct {
		StmBse
		Stm    Stm
		Length unt.Unt
	}
	StmAggWrsiSeg struct {
		StmSeg
		Length unt.Unt
	}
	StmAggAlma struct {
		StmBse
		Stm    Stm
		Length unt.Unt
	}
	StmAggAlmaSeg struct {
		StmSeg
		Length unt.Unt
	}
	StmAggVrnc struct {
		StmBse
		Stm    Stm
		Length unt.Unt
	}
	StmAggVrncSeg struct {
		StmSeg
		Length unt.Unt
	}
	StmAggStd struct {
		StmBse
		Stm    Stm
		Length unt.Unt
	}
	StmAggStdSeg struct {
		StmSeg
		Length unt.Unt
	}
	StmAggRngFul struct {
		StmBse
		Stm    Stm
		Length unt.Unt
	}
	StmAggRngFulSeg struct {
		StmSeg
		Length unt.Unt
	}
	StmAggRngLst struct {
		StmBse
		Stm    Stm
		Length unt.Unt
	}
	StmAggRngLstSeg struct {
		StmSeg
		Length unt.Unt
	}
	StmAggProLst struct {
		StmBse
		Stm    Stm
		Length unt.Unt
	}
	StmAggProLstSeg struct {
		StmSeg
		Length unt.Unt
	}
	StmAggProSma struct {
		StmBse
		Stm    Stm
		Length unt.Unt
	}
	StmAggProSmaSeg struct {
		StmSeg
		Length unt.Unt
	}
	StmAggProAlma struct {
		StmBse
		Stm    Stm
		Length unt.Unt
	}
	StmAggProAlmaSeg struct {
		StmSeg
		Length unt.Unt
	}
	StmAggEma struct {
		StmBse
		Stm    Stm
		Length unt.Unt
	}
	StmInrAdd struct {
		StmBse
		Stm Stm
		Off unt.Unt
	}
	StmInrAddSeg struct {
		StmSeg
		Off unt.Unt
	}
	StmInrSub struct {
		StmBse
		Stm Stm
		Off unt.Unt
	}
	StmInrSubSeg struct {
		StmSeg
		Off unt.Unt
	}
	StmInrMul struct {
		StmBse
		Stm Stm
		Off unt.Unt
	}
	StmInrMulSeg struct {
		StmSeg
		Off unt.Unt
	}
	StmInrDiv struct {
		StmBse
		Stm Stm
		Off unt.Unt
	}
	StmInrDivSeg struct {
		StmSeg
		Off unt.Unt
	}
	StmInrRem struct {
		StmBse
		Stm Stm
		Off unt.Unt
	}
	StmInrRemSeg struct {
		StmSeg
		Off unt.Unt
	}
	StmInrPow struct {
		StmBse
		Stm Stm
		Off unt.Unt
	}
	StmInrPowSeg struct {
		StmSeg
		Off unt.Unt
	}
	StmInrMin struct {
		StmBse
		Stm Stm
		Off unt.Unt
	}
	StmInrMinSeg struct {
		StmSeg
		Off unt.Unt
	}
	StmInrMax struct {
		StmBse
		Stm Stm
		Off unt.Unt
	}
	StmInrMaxSeg struct {
		StmSeg
		Off unt.Unt
	}
	StmInr1Slp struct {
		StmBse
		Stm Stm
		Off unt.Unt
	}
	StmInr1SlpSeg struct {
		StmSeg
		Off  unt.Unt
		Tmes *tmes.Tmes
	}
	StmOtrAdd struct {
		StmBse
		Stm Stm
		Off unt.Unt
		A   Stm
	}
	StmOtrAddSeg struct {
		StmSeg
		Off   unt.Unt
		ValsA *flts.Flts
	}
	StmOtrSub struct {
		StmBse
		Stm Stm
		Off unt.Unt
		A   Stm
	}
	StmOtrSubSeg struct {
		StmSeg
		Off   unt.Unt
		ValsA *flts.Flts
	}
	StmOtrMul struct {
		StmBse
		Stm Stm
		Off unt.Unt
		A   Stm
	}
	StmOtrMulSeg struct {
		StmSeg
		Off   unt.Unt
		ValsA *flts.Flts
	}
	StmOtrDiv struct {
		StmBse
		Stm Stm
		Off unt.Unt
		A   Stm
	}
	StmOtrDivSeg struct {
		StmSeg
		Off   unt.Unt
		ValsA *flts.Flts
	}
	StmOtrRem struct {
		StmBse
		Stm Stm
		Off unt.Unt
		A   Stm
	}
	StmOtrRemSeg struct {
		StmSeg
		Off   unt.Unt
		ValsA *flts.Flts
	}
	StmOtrPow struct {
		StmBse
		Stm Stm
		Off unt.Unt
		A   Stm
	}
	StmOtrPowSeg struct {
		StmSeg
		Off   unt.Unt
		ValsA *flts.Flts
	}
	StmOtrMin struct {
		StmBse
		Stm Stm
		Off unt.Unt
		A   Stm
	}
	StmOtrMinSeg struct {
		StmSeg
		Off   unt.Unt
		ValsA *flts.Flts
	}
	StmOtrMax struct {
		StmBse
		Stm Stm
		Off unt.Unt
		A   Stm
	}
	StmOtrMaxSeg struct {
		StmSeg
		Off   unt.Unt
		ValsA *flts.Flts
	}
)

func (x *StmRteFst) Name() str.Str             { return str.Str("Fst") }
func (x *StmRteFst) PrmWrt(b *strings.Builder) {}
func (x *StmRteFst) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *StmRteFst) StrWrt(b *strings.Builder) {
	x.Side.StrWrt(b)
	b.WriteString(".fst(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *StmRteFst) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *StmRteFstSeg) Act() {
	for n := x.Idx; n < x.Lim; n++ {
		elm := (*x.Vals)[(*x.ValBnds)[n].Idx:(*x.ValBnds)[n].Lim]
		(*x.Out)[n] = elm.Fst()
	}
}
func (x *StmRteLst) Name() str.Str             { return str.Str("Lst") }
func (x *StmRteLst) PrmWrt(b *strings.Builder) {}
func (x *StmRteLst) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *StmRteLst) StrWrt(b *strings.Builder) {
	x.Side.StrWrt(b)
	b.WriteString(".lst(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *StmRteLst) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *StmRteLstSeg) Act() {
	for n := x.Idx; n < x.Lim; n++ {
		elm := (*x.Vals)[(*x.ValBnds)[n].Idx:(*x.ValBnds)[n].Lim]
		(*x.Out)[n] = elm.Lst()
	}
}
func (x *StmRteSum) Name() str.Str             { return str.Str("Sum") }
func (x *StmRteSum) PrmWrt(b *strings.Builder) {}
func (x *StmRteSum) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *StmRteSum) StrWrt(b *strings.Builder) {
	x.Side.StrWrt(b)
	b.WriteString(".sum(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *StmRteSum) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *StmRteSumSeg) Act() {
	for n := x.Idx; n < x.Lim; n++ {
		elm := (*x.Vals)[(*x.ValBnds)[n].Idx:(*x.ValBnds)[n].Lim]
		(*x.Out)[n] = elm.Sum()
	}
}
func (x *StmRtePrd) Name() str.Str             { return str.Str("Prd") }
func (x *StmRtePrd) PrmWrt(b *strings.Builder) {}
func (x *StmRtePrd) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *StmRtePrd) StrWrt(b *strings.Builder) {
	x.Side.StrWrt(b)
	b.WriteString(".prd(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *StmRtePrd) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *StmRtePrdSeg) Act() {
	for n := x.Idx; n < x.Lim; n++ {
		elm := (*x.Vals)[(*x.ValBnds)[n].Idx:(*x.ValBnds)[n].Lim]
		(*x.Out)[n] = elm.Prd()
	}
}
func (x *StmRteMin) Name() str.Str             { return str.Str("Min") }
func (x *StmRteMin) PrmWrt(b *strings.Builder) {}
func (x *StmRteMin) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *StmRteMin) StrWrt(b *strings.Builder) {
	x.Side.StrWrt(b)
	b.WriteString(".min(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *StmRteMin) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *StmRteMinSeg) Act() {
	for n := x.Idx; n < x.Lim; n++ {
		elm := (*x.Vals)[(*x.ValBnds)[n].Idx:(*x.ValBnds)[n].Lim]
		(*x.Out)[n] = elm.Min()
	}
}
func (x *StmRteMax) Name() str.Str             { return str.Str("Max") }
func (x *StmRteMax) PrmWrt(b *strings.Builder) {}
func (x *StmRteMax) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *StmRteMax) StrWrt(b *strings.Builder) {
	x.Side.StrWrt(b)
	b.WriteString(".max(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *StmRteMax) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *StmRteMaxSeg) Act() {
	for n := x.Idx; n < x.Lim; n++ {
		elm := (*x.Vals)[(*x.ValBnds)[n].Idx:(*x.ValBnds)[n].Lim]
		(*x.Out)[n] = elm.Max()
	}
}
func (x *StmRteMid) Name() str.Str             { return str.Str("Mid") }
func (x *StmRteMid) PrmWrt(b *strings.Builder) {}
func (x *StmRteMid) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *StmRteMid) StrWrt(b *strings.Builder) {
	x.Side.StrWrt(b)
	b.WriteString(".mid(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *StmRteMid) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *StmRteMidSeg) Act() {
	for n := x.Idx; n < x.Lim; n++ {
		elm := (*x.Vals)[(*x.ValBnds)[n].Idx:(*x.ValBnds)[n].Lim]
		(*x.Out)[n] = elm.Mid()
	}
}
func (x *StmRteMdn) Name() str.Str             { return str.Str("Mdn") }
func (x *StmRteMdn) PrmWrt(b *strings.Builder) {}
func (x *StmRteMdn) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *StmRteMdn) StrWrt(b *strings.Builder) {
	x.Side.StrWrt(b)
	b.WriteString(".mdn(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *StmRteMdn) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *StmRteMdnSeg) Act() {
	for n := x.Idx; n < x.Lim; n++ {
		elm := (*x.Vals)[(*x.ValBnds)[n].Idx:(*x.ValBnds)[n].Lim]
		(*x.Out)[n] = elm.Mdn()
	}
}
func (x *StmRteSma) Name() str.Str             { return str.Str("Sma") }
func (x *StmRteSma) PrmWrt(b *strings.Builder) {}
func (x *StmRteSma) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *StmRteSma) StrWrt(b *strings.Builder) {
	x.Side.StrWrt(b)
	b.WriteString(".sma(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *StmRteSma) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *StmRteSmaSeg) Act() {
	for n := x.Idx; n < x.Lim; n++ {
		elm := (*x.Vals)[(*x.ValBnds)[n].Idx:(*x.ValBnds)[n].Lim]
		(*x.Out)[n] = elm.Sma()
	}
}
func (x *StmRteGma) Name() str.Str             { return str.Str("Gma") }
func (x *StmRteGma) PrmWrt(b *strings.Builder) {}
func (x *StmRteGma) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *StmRteGma) StrWrt(b *strings.Builder) {
	x.Side.StrWrt(b)
	b.WriteString(".gma(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *StmRteGma) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *StmRteGmaSeg) Act() {
	for n := x.Idx; n < x.Lim; n++ {
		elm := (*x.Vals)[(*x.ValBnds)[n].Idx:(*x.ValBnds)[n].Lim]
		(*x.Out)[n] = elm.Gma()
	}
}
func (x *StmRteWma) Name() str.Str             { return str.Str("Wma") }
func (x *StmRteWma) PrmWrt(b *strings.Builder) {}
func (x *StmRteWma) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *StmRteWma) StrWrt(b *strings.Builder) {
	x.Side.StrWrt(b)
	b.WriteString(".wma(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *StmRteWma) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *StmRteWmaSeg) Act() {
	for n := x.Idx; n < x.Lim; n++ {
		elm := (*x.Vals)[(*x.ValBnds)[n].Idx:(*x.ValBnds)[n].Lim]
		(*x.Out)[n] = elm.Wma()
	}
}
func (x *StmRteRsi) Name() str.Str             { return str.Str("Rsi") }
func (x *StmRteRsi) PrmWrt(b *strings.Builder) {}
func (x *StmRteRsi) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *StmRteRsi) StrWrt(b *strings.Builder) {
	x.Side.StrWrt(b)
	b.WriteString(".rsi(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *StmRteRsi) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *StmRteRsiSeg) Act() {
	for n := x.Idx; n < x.Lim; n++ {
		elm := (*x.Vals)[(*x.ValBnds)[n].Idx:(*x.ValBnds)[n].Lim]
		(*x.Out)[n] = elm.Rsi()
	}
}
func (x *StmRteWrsi) Name() str.Str             { return str.Str("Wrsi") }
func (x *StmRteWrsi) PrmWrt(b *strings.Builder) {}
func (x *StmRteWrsi) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *StmRteWrsi) StrWrt(b *strings.Builder) {
	x.Side.StrWrt(b)
	b.WriteString(".wrsi(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *StmRteWrsi) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *StmRteWrsiSeg) Act() {
	for n := x.Idx; n < x.Lim; n++ {
		elm := (*x.Vals)[(*x.ValBnds)[n].Idx:(*x.ValBnds)[n].Lim]
		(*x.Out)[n] = elm.Wrsi()
	}
}
func (x *StmRteAlma) Name() str.Str             { return str.Str("Alma") }
func (x *StmRteAlma) PrmWrt(b *strings.Builder) {}
func (x *StmRteAlma) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *StmRteAlma) StrWrt(b *strings.Builder) {
	x.Side.StrWrt(b)
	b.WriteString(".alma(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *StmRteAlma) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *StmRteAlmaSeg) Act() {
	for n := x.Idx; n < x.Lim; n++ {
		elm := (*x.Vals)[(*x.ValBnds)[n].Idx:(*x.ValBnds)[n].Lim]
		(*x.Out)[n] = elm.Alma()
	}
}
func (x *StmRteVrnc) Name() str.Str             { return str.Str("Vrnc") }
func (x *StmRteVrnc) PrmWrt(b *strings.Builder) {}
func (x *StmRteVrnc) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *StmRteVrnc) StrWrt(b *strings.Builder) {
	x.Side.StrWrt(b)
	b.WriteString(".vrnc(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *StmRteVrnc) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *StmRteVrncSeg) Act() {
	for n := x.Idx; n < x.Lim; n++ {
		elm := (*x.Vals)[(*x.ValBnds)[n].Idx:(*x.ValBnds)[n].Lim]
		(*x.Out)[n] = elm.Vrnc()
	}
}
func (x *StmRteStd) Name() str.Str             { return str.Str("Std") }
func (x *StmRteStd) PrmWrt(b *strings.Builder) {}
func (x *StmRteStd) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *StmRteStd) StrWrt(b *strings.Builder) {
	x.Side.StrWrt(b)
	b.WriteString(".std(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *StmRteStd) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *StmRteStdSeg) Act() {
	for n := x.Idx; n < x.Lim; n++ {
		elm := (*x.Vals)[(*x.ValBnds)[n].Idx:(*x.ValBnds)[n].Lim]
		(*x.Out)[n] = elm.Std()
	}
}
func (x *StmRteRngFul) Name() str.Str             { return str.Str("RngFul") }
func (x *StmRteRngFul) PrmWrt(b *strings.Builder) {}
func (x *StmRteRngFul) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *StmRteRngFul) StrWrt(b *strings.Builder) {
	x.Side.StrWrt(b)
	b.WriteString(".rngFul(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *StmRteRngFul) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *StmRteRngFulSeg) Act() {
	for n := x.Idx; n < x.Lim; n++ {
		elm := (*x.Vals)[(*x.ValBnds)[n].Idx:(*x.ValBnds)[n].Lim]
		(*x.Out)[n] = elm.RngFul()
	}
}
func (x *StmRteRngLst) Name() str.Str             { return str.Str("RngLst") }
func (x *StmRteRngLst) PrmWrt(b *strings.Builder) {}
func (x *StmRteRngLst) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *StmRteRngLst) StrWrt(b *strings.Builder) {
	x.Side.StrWrt(b)
	b.WriteString(".rngLst(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *StmRteRngLst) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *StmRteRngLstSeg) Act() {
	for n := x.Idx; n < x.Lim; n++ {
		elm := (*x.Vals)[(*x.ValBnds)[n].Idx:(*x.ValBnds)[n].Lim]
		(*x.Out)[n] = elm.RngLst()
	}
}
func (x *StmRteProLst) Name() str.Str             { return str.Str("ProLst") }
func (x *StmRteProLst) PrmWrt(b *strings.Builder) {}
func (x *StmRteProLst) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *StmRteProLst) StrWrt(b *strings.Builder) {
	x.Side.StrWrt(b)
	b.WriteString(".proLst(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *StmRteProLst) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *StmRteProLstSeg) Act() {
	for n := x.Idx; n < x.Lim; n++ {
		elm := (*x.Vals)[(*x.ValBnds)[n].Idx:(*x.ValBnds)[n].Lim]
		(*x.Out)[n] = elm.ProLst()
	}
}
func (x *StmRteProSma) Name() str.Str             { return str.Str("ProSma") }
func (x *StmRteProSma) PrmWrt(b *strings.Builder) {}
func (x *StmRteProSma) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *StmRteProSma) StrWrt(b *strings.Builder) {
	x.Side.StrWrt(b)
	b.WriteString(".proSma(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *StmRteProSma) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *StmRteProSmaSeg) Act() {
	for n := x.Idx; n < x.Lim; n++ {
		elm := (*x.Vals)[(*x.ValBnds)[n].Idx:(*x.ValBnds)[n].Lim]
		(*x.Out)[n] = elm.ProSma()
	}
}
func (x *StmRteProAlma) Name() str.Str             { return str.Str("ProAlma") }
func (x *StmRteProAlma) PrmWrt(b *strings.Builder) {}
func (x *StmRteProAlma) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *StmRteProAlma) StrWrt(b *strings.Builder) {
	x.Side.StrWrt(b)
	b.WriteString(".proAlma(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *StmRteProAlma) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *StmRteProAlmaSeg) Act() {
	for n := x.Idx; n < x.Lim; n++ {
		elm := (*x.Vals)[(*x.ValBnds)[n].Idx:(*x.ValBnds)[n].Lim]
		(*x.Out)[n] = elm.ProAlma()
	}
}
func (x *StmRte1Sar) Name() str.Str { return str.Str("Sar") }
func (x *StmRte1Sar) PrmWrt(b *strings.Builder) {
	x.AfInc.StrWrt(b)
	b.WriteRune(' ')
	x.AfMax.StrWrt(b)
}
func (x *StmRte1Sar) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *StmRte1Sar) StrWrt(b *strings.Builder) {
	x.Side.StrWrt(b)
	b.WriteString(".sar(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *StmRte1Sar) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *StmRteEma) Name() str.Str             { return str.Str("Ema") }
func (x *StmRteEma) PrmWrt(b *strings.Builder) {}
func (x *StmRteEma) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *StmRteEma) StrWrt(b *strings.Builder) {
	x.Side.StrWrt(b)
	b.WriteString(".ema(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *StmRteEma) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *StmBse) Bse() *StmBse { return x }
func (x *StmBse) At(ts *tmes.Tmes) (r *flts.Flts) {
	if x.Vals == nil || len(*x.Vals) == 0 || ts == nil || len(*ts) == 0 {
		return flts.New()
	}
	r = flts.MakeEmp(ts.Cnt())
	segBnds, acts := bnds.Segs(ts.Cnt())
	for n, segBnd := range *segBnds {
		seg := &StmAtSeg{}
		seg.Bnd = segBnd
		seg.Stm = x
		seg.AtTmes = ts
		seg.Out = flts.MakeEmp(segBnd.Cnt())
		acts[n] = seg
	}
	sys.Run().Pll(acts...)
	for _, act := range acts {
		r.Mrg(act.(*StmAtSeg).Out)
	}
	return r
}
func (x *StmBse) Atf(ts *tmes.Tmes) (r []float32) {
	if x.Vals == nil || len(*x.Vals) == 0 || ts == nil || len(*ts) == 0 {
		return r
	}
	r = make([]float32, 0, ts.Cnt())
	segBnds, acts := bnds.Segs(ts.Cnt())
	for n, segBnd := range *segBnds {
		seg := &StmAtfSeg{}
		seg.Bnd = segBnd
		seg.Stm = x
		seg.AtTmes = ts
		seg.Out = make([]float32, 0, segBnd.Cnt())
		acts[n] = seg
	}
	sys.Run().Pll(acts...)
	for _, act := range acts {
		r = append(r, act.(*StmAtfSeg).Out...)
	}
	return r
}
func (x *StmBse) UnaPos() Stm {
	r := &StmUnaPos{}
	r.Slf = r
	r.Stm = x.Slf
	if ana.Cfg.Trc.IsHstStm() {
		sys.Logf("%p hst.StmUnaPos(%v)", r, r.Prm())
	}
	if x.Vals.Cnt() == 0 {
		return r
	}
	if len(*x.Tmes) != len(*x.Vals) {
		err.Panicf("length unequal (Tmes:%v Vals:%v)", len(*x.Tmes), len(*x.Vals))
	}
	r.Tmes = x.Tmes
	r.Vals = flts.Make(x.Vals.Cnt())
	segBnds, acts := bnds.Segs(r.Vals.Cnt())
	for n, segBnd := range *segBnds {
		seg := &StmUnaPosSeg{}
		seg.Bnd = segBnd
		seg.Vals = x.Vals
		seg.Out = r.Vals
		acts[n] = seg
	}
	sys.Run().Pll(acts...)
	return r
}
func (x *StmUnaPos) Name() str.Str             { return str.Str("UnaPos") }
func (x *StmUnaPos) PrmWrt(b *strings.Builder) {}
func (x *StmUnaPos) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *StmUnaPos) StrWrt(b *strings.Builder) {
	x.Stm.StrWrt(b)
	b.WriteString(".unaPos(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *StmUnaPos) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *StmUnaPosSeg) Act() {
	for n := x.Idx; n < x.Lim; n++ {
		if (*x.Vals)[n] < 0 {
			(*x.Out)[n] = -(*x.Vals)[n]
		} else {
			(*x.Out)[n] = (*x.Vals)[n]
		}
	}
}
func (x *StmBse) UnaNeg() Stm {
	r := &StmUnaNeg{}
	r.Slf = r
	r.Stm = x.Slf
	if ana.Cfg.Trc.IsHstStm() {
		sys.Logf("%p hst.StmUnaNeg(%v)", r, r.Prm())
	}
	if x.Vals.Cnt() == 0 {
		return r
	}
	if len(*x.Tmes) != len(*x.Vals) {
		err.Panicf("length unequal (Tmes:%v Vals:%v)", len(*x.Tmes), len(*x.Vals))
	}
	r.Tmes = x.Tmes
	r.Vals = flts.Make(x.Vals.Cnt())
	segBnds, acts := bnds.Segs(r.Vals.Cnt())
	for n, segBnd := range *segBnds {
		seg := &StmUnaNegSeg{}
		seg.Bnd = segBnd
		seg.Vals = x.Vals
		seg.Out = r.Vals
		acts[n] = seg
	}
	sys.Run().Pll(acts...)
	return r
}
func (x *StmUnaNeg) Name() str.Str             { return str.Str("UnaNeg") }
func (x *StmUnaNeg) PrmWrt(b *strings.Builder) {}
func (x *StmUnaNeg) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *StmUnaNeg) StrWrt(b *strings.Builder) {
	x.Stm.StrWrt(b)
	b.WriteString(".unaNeg(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *StmUnaNeg) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *StmUnaNegSeg) Act() {
	for n := x.Idx; n < x.Lim; n++ {
		if (*x.Vals)[n] > 0 {
			(*x.Out)[n] = -(*x.Vals)[n]
		} else {
			(*x.Out)[n] = (*x.Vals)[n]
		}
	}
}
func (x *StmBse) UnaInv() Stm {
	r := &StmUnaInv{}
	r.Slf = r
	r.Stm = x.Slf
	if ana.Cfg.Trc.IsHstStm() {
		sys.Logf("%p hst.StmUnaInv(%v)", r, r.Prm())
	}
	if x.Vals.Cnt() == 0 {
		return r
	}
	if len(*x.Tmes) != len(*x.Vals) {
		err.Panicf("length unequal (Tmes:%v Vals:%v)", len(*x.Tmes), len(*x.Vals))
	}
	r.Tmes = x.Tmes
	r.Vals = flts.Make(x.Vals.Cnt())
	segBnds, acts := bnds.Segs(r.Vals.Cnt())
	for n, segBnd := range *segBnds {
		seg := &StmUnaInvSeg{}
		seg.Bnd = segBnd
		seg.Vals = x.Vals
		seg.Out = r.Vals
		acts[n] = seg
	}
	sys.Run().Pll(acts...)
	return r
}
func (x *StmUnaInv) Name() str.Str             { return str.Str("UnaInv") }
func (x *StmUnaInv) PrmWrt(b *strings.Builder) {}
func (x *StmUnaInv) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *StmUnaInv) StrWrt(b *strings.Builder) {
	x.Stm.StrWrt(b)
	b.WriteString(".unaInv(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *StmUnaInv) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *StmUnaInvSeg) Act() {
	for n := x.Idx; n < x.Lim; n++ {
		(*x.Out)[n] = -(*x.Vals)[n]
	}
}
func (x *StmBse) UnaSqr() Stm {
	r := &StmUnaSqr{}
	r.Slf = r
	r.Stm = x.Slf
	if ana.Cfg.Trc.IsHstStm() {
		sys.Logf("%p hst.StmUnaSqr(%v)", r, r.Prm())
	}
	if x.Vals.Cnt() == 0 {
		return r
	}
	if len(*x.Tmes) != len(*x.Vals) {
		err.Panicf("length unequal (Tmes:%v Vals:%v)", len(*x.Tmes), len(*x.Vals))
	}
	r.Tmes = x.Tmes
	r.Vals = flts.Make(x.Vals.Cnt())
	segBnds, acts := bnds.Segs(r.Vals.Cnt())
	for n, segBnd := range *segBnds {
		seg := &StmUnaSqrSeg{}
		seg.Bnd = segBnd
		seg.Vals = x.Vals
		seg.Out = r.Vals
		acts[n] = seg
	}
	sys.Run().Pll(acts...)
	return r
}
func (x *StmUnaSqr) Name() str.Str             { return str.Str("UnaSqr") }
func (x *StmUnaSqr) PrmWrt(b *strings.Builder) {}
func (x *StmUnaSqr) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *StmUnaSqr) StrWrt(b *strings.Builder) {
	x.Stm.StrWrt(b)
	b.WriteString(".unaSqr(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *StmUnaSqr) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *StmUnaSqrSeg) Act() {
	for n := x.Idx; n < x.Lim; n++ {
		(*x.Out)[n] = (*x.Vals)[n] * (*x.Vals)[n]
	}
}
func (x *StmBse) UnaSqrt() Stm {
	r := &StmUnaSqrt{}
	r.Slf = r
	r.Stm = x.Slf
	if ana.Cfg.Trc.IsHstStm() {
		sys.Logf("%p hst.StmUnaSqrt(%v)", r, r.Prm())
	}
	if x.Vals.Cnt() == 0 {
		return r
	}
	if len(*x.Tmes) != len(*x.Vals) {
		err.Panicf("length unequal (Tmes:%v Vals:%v)", len(*x.Tmes), len(*x.Vals))
	}
	r.Tmes = x.Tmes
	r.Vals = flts.Make(x.Vals.Cnt())
	segBnds, acts := bnds.Segs(r.Vals.Cnt())
	for n, segBnd := range *segBnds {
		seg := &StmUnaSqrtSeg{}
		seg.Bnd = segBnd
		seg.Vals = x.Vals
		seg.Out = r.Vals
		acts[n] = seg
	}
	sys.Run().Pll(acts...)
	return r
}
func (x *StmUnaSqrt) Name() str.Str             { return str.Str("UnaSqrt") }
func (x *StmUnaSqrt) PrmWrt(b *strings.Builder) {}
func (x *StmUnaSqrt) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *StmUnaSqrt) StrWrt(b *strings.Builder) {
	x.Stm.StrWrt(b)
	b.WriteString(".unaSqrt(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *StmUnaSqrt) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *StmUnaSqrtSeg) Act() {
	for n := x.Idx; n < x.Lim; n++ {
		if (*x.Vals)[n] <= 0 {
			(*x.Out)[n] = 0
		} else {
			(*x.Out)[n] = flt.Flt(math.Sqrt(float64((*x.Vals)[n])))
		}
	}
}
func (x *StmBse) SclAdd(scl flt.Flt) Stm {
	r := &StmSclAdd{}
	r.Slf = r
	r.Stm = x.Slf
	r.Scl = scl
	if ana.Cfg.Trc.IsHstStm() {
		sys.Logf("%p hst.StmSclAdd(%v)", r, r.Prm())
	}
	if x.Vals.Cnt() == 0 {
		return r
	}
	if len(*x.Tmes) != len(*x.Vals) {
		err.Panicf("length unequal (Tmes:%v Vals:%v)", len(*x.Tmes), len(*x.Vals))
	}
	r.Tmes = x.Tmes
	r.Vals = flts.Make(x.Vals.Cnt())
	segBnds, acts := bnds.Segs(r.Vals.Cnt())
	for n, segBnd := range *segBnds {
		seg := &StmSclAddSeg{}
		seg.Bnd = segBnd
		seg.Vals = x.Vals
		seg.Out = r.Vals
		seg.Scl = r.Scl
		acts[n] = seg
	}
	sys.Run().Pll(acts...)
	return r
}
func (x *StmSclAdd) Name() str.Str             { return str.Str("SclAdd") }
func (x *StmSclAdd) PrmWrt(b *strings.Builder) { x.Scl.StrWrt(b) }
func (x *StmSclAdd) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *StmSclAdd) StrWrt(b *strings.Builder) {
	x.Stm.StrWrt(b)
	b.WriteString(".sclAdd(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *StmSclAdd) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *StmSclAddSeg) Act() {
	for n := x.Idx; n < x.Lim; n++ {
		(*x.Out)[n] = (*x.Vals)[n] + x.Scl
	}
}
func (x *StmBse) SclSub(scl flt.Flt) Stm {
	r := &StmSclSub{}
	r.Slf = r
	r.Stm = x.Slf
	r.Scl = scl
	if ana.Cfg.Trc.IsHstStm() {
		sys.Logf("%p hst.StmSclSub(%v)", r, r.Prm())
	}
	if x.Vals.Cnt() == 0 {
		return r
	}
	if len(*x.Tmes) != len(*x.Vals) {
		err.Panicf("length unequal (Tmes:%v Vals:%v)", len(*x.Tmes), len(*x.Vals))
	}
	r.Tmes = x.Tmes
	r.Vals = flts.Make(x.Vals.Cnt())
	segBnds, acts := bnds.Segs(r.Vals.Cnt())
	for n, segBnd := range *segBnds {
		seg := &StmSclSubSeg{}
		seg.Bnd = segBnd
		seg.Vals = x.Vals
		seg.Out = r.Vals
		seg.Scl = r.Scl
		acts[n] = seg
	}
	sys.Run().Pll(acts...)
	return r
}
func (x *StmSclSub) Name() str.Str             { return str.Str("SclSub") }
func (x *StmSclSub) PrmWrt(b *strings.Builder) { x.Scl.StrWrt(b) }
func (x *StmSclSub) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *StmSclSub) StrWrt(b *strings.Builder) {
	x.Stm.StrWrt(b)
	b.WriteString(".sclSub(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *StmSclSub) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *StmSclSubSeg) Act() {
	for n := x.Idx; n < x.Lim; n++ {
		(*x.Out)[n] = (*x.Vals)[n] - x.Scl
	}
}
func (x *StmBse) SclMul(scl flt.Flt) Stm {
	r := &StmSclMul{}
	r.Slf = r
	r.Stm = x.Slf
	r.Scl = scl
	if ana.Cfg.Trc.IsHstStm() {
		sys.Logf("%p hst.StmSclMul(%v)", r, r.Prm())
	}
	if x.Vals.Cnt() == 0 {
		return r
	}
	if len(*x.Tmes) != len(*x.Vals) {
		err.Panicf("length unequal (Tmes:%v Vals:%v)", len(*x.Tmes), len(*x.Vals))
	}
	r.Tmes = x.Tmes
	r.Vals = flts.Make(x.Vals.Cnt())
	segBnds, acts := bnds.Segs(r.Vals.Cnt())
	for n, segBnd := range *segBnds {
		seg := &StmSclMulSeg{}
		seg.Bnd = segBnd
		seg.Vals = x.Vals
		seg.Out = r.Vals
		seg.Scl = r.Scl
		acts[n] = seg
	}
	sys.Run().Pll(acts...)
	return r
}
func (x *StmSclMul) Name() str.Str             { return str.Str("SclMul") }
func (x *StmSclMul) PrmWrt(b *strings.Builder) { x.Scl.StrWrt(b) }
func (x *StmSclMul) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *StmSclMul) StrWrt(b *strings.Builder) {
	x.Stm.StrWrt(b)
	b.WriteString(".sclMul(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *StmSclMul) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *StmSclMulSeg) Act() {
	for n := x.Idx; n < x.Lim; n++ {
		(*x.Out)[n] = (*x.Vals)[n] * x.Scl
	}
}
func (x *StmBse) SclDiv(scl flt.Flt) Stm {
	r := &StmSclDiv{}
	r.Slf = r
	r.Stm = x.Slf
	r.Scl = scl
	if ana.Cfg.Trc.IsHstStm() {
		sys.Logf("%p hst.StmSclDiv(%v)", r, r.Prm())
	}
	if x.Vals.Cnt() == 0 {
		return r
	}
	if len(*x.Tmes) != len(*x.Vals) {
		err.Panicf("length unequal (Tmes:%v Vals:%v)", len(*x.Tmes), len(*x.Vals))
	}
	r.Tmes = x.Tmes
	r.Vals = flts.Make(x.Vals.Cnt())
	segBnds, acts := bnds.Segs(r.Vals.Cnt())
	for n, segBnd := range *segBnds {
		seg := &StmSclDivSeg{}
		seg.Bnd = segBnd
		seg.Vals = x.Vals
		seg.Out = r.Vals
		seg.Scl = r.Scl
		acts[n] = seg
	}
	sys.Run().Pll(acts...)
	return r
}
func (x *StmSclDiv) Name() str.Str             { return str.Str("SclDiv") }
func (x *StmSclDiv) PrmWrt(b *strings.Builder) { x.Scl.StrWrt(b) }
func (x *StmSclDiv) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *StmSclDiv) StrWrt(b *strings.Builder) {
	x.Stm.StrWrt(b)
	b.WriteString(".sclDiv(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *StmSclDiv) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *StmSclDivSeg) Act() {
	for n := x.Idx; n < x.Lim; n++ {
		if x.Scl == 0 {
			(*x.Out)[n] = 0
		} else {
			(*x.Out)[n] = (*x.Vals)[n] / x.Scl
		}
	}
}
func (x *StmBse) SclRem(scl flt.Flt) Stm {
	r := &StmSclRem{}
	r.Slf = r
	r.Stm = x.Slf
	r.Scl = scl
	if ana.Cfg.Trc.IsHstStm() {
		sys.Logf("%p hst.StmSclRem(%v)", r, r.Prm())
	}
	if x.Vals.Cnt() == 0 {
		return r
	}
	if len(*x.Tmes) != len(*x.Vals) {
		err.Panicf("length unequal (Tmes:%v Vals:%v)", len(*x.Tmes), len(*x.Vals))
	}
	r.Tmes = x.Tmes
	r.Vals = flts.Make(x.Vals.Cnt())
	segBnds, acts := bnds.Segs(r.Vals.Cnt())
	for n, segBnd := range *segBnds {
		seg := &StmSclRemSeg{}
		seg.Bnd = segBnd
		seg.Vals = x.Vals
		seg.Out = r.Vals
		seg.Scl = r.Scl
		acts[n] = seg
	}
	sys.Run().Pll(acts...)
	return r
}
func (x *StmSclRem) Name() str.Str             { return str.Str("SclRem") }
func (x *StmSclRem) PrmWrt(b *strings.Builder) { x.Scl.StrWrt(b) }
func (x *StmSclRem) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *StmSclRem) StrWrt(b *strings.Builder) {
	x.Stm.StrWrt(b)
	b.WriteString(".sclRem(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *StmSclRem) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *StmSclRemSeg) Act() {
	for n := x.Idx; n < x.Lim; n++ {
		if x.Scl == 0 {
			(*x.Out)[n] = 0
		} else {
			(*x.Out)[n] = flt.Flt(math.Remainder(float64((*x.Vals)[n]), float64(x.Scl)))
		}
	}
}
func (x *StmBse) SclPow(scl flt.Flt) Stm {
	r := &StmSclPow{}
	r.Slf = r
	r.Stm = x.Slf
	r.Scl = scl
	if ana.Cfg.Trc.IsHstStm() {
		sys.Logf("%p hst.StmSclPow(%v)", r, r.Prm())
	}
	if x.Vals.Cnt() == 0 {
		return r
	}
	if len(*x.Tmes) != len(*x.Vals) {
		err.Panicf("length unequal (Tmes:%v Vals:%v)", len(*x.Tmes), len(*x.Vals))
	}
	r.Tmes = x.Tmes
	r.Vals = flts.Make(x.Vals.Cnt())
	segBnds, acts := bnds.Segs(r.Vals.Cnt())
	for n, segBnd := range *segBnds {
		seg := &StmSclPowSeg{}
		seg.Bnd = segBnd
		seg.Vals = x.Vals
		seg.Out = r.Vals
		seg.Scl = r.Scl
		acts[n] = seg
	}
	sys.Run().Pll(acts...)
	return r
}
func (x *StmSclPow) Name() str.Str             { return str.Str("SclPow") }
func (x *StmSclPow) PrmWrt(b *strings.Builder) { x.Scl.StrWrt(b) }
func (x *StmSclPow) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *StmSclPow) StrWrt(b *strings.Builder) {
	x.Stm.StrWrt(b)
	b.WriteString(".sclPow(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *StmSclPow) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *StmSclPowSeg) Act() {
	for n := x.Idx; n < x.Lim; n++ {
		(*x.Out)[n] = flt.Flt(math.Pow(float64((*x.Vals)[n]), float64(x.Scl)))
	}
}
func (x *StmBse) SclMin(scl flt.Flt) Stm {
	r := &StmSclMin{}
	r.Slf = r
	r.Stm = x.Slf
	r.Scl = scl
	if ana.Cfg.Trc.IsHstStm() {
		sys.Logf("%p hst.StmSclMin(%v)", r, r.Prm())
	}
	if x.Vals.Cnt() == 0 {
		return r
	}
	if len(*x.Tmes) != len(*x.Vals) {
		err.Panicf("length unequal (Tmes:%v Vals:%v)", len(*x.Tmes), len(*x.Vals))
	}
	r.Tmes = x.Tmes
	r.Vals = flts.Make(x.Vals.Cnt())
	segBnds, acts := bnds.Segs(r.Vals.Cnt())
	for n, segBnd := range *segBnds {
		seg := &StmSclMinSeg{}
		seg.Bnd = segBnd
		seg.Vals = x.Vals
		seg.Out = r.Vals
		seg.Scl = r.Scl
		acts[n] = seg
	}
	sys.Run().Pll(acts...)
	return r
}
func (x *StmSclMin) Name() str.Str             { return str.Str("SclMin") }
func (x *StmSclMin) PrmWrt(b *strings.Builder) { x.Scl.StrWrt(b) }
func (x *StmSclMin) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *StmSclMin) StrWrt(b *strings.Builder) {
	x.Stm.StrWrt(b)
	b.WriteString(".sclMin(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *StmSclMin) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *StmSclMinSeg) Act() {
	for n := x.Idx; n < x.Lim; n++ {
		if (*x.Vals)[n] < x.Scl {
			(*x.Out)[n] = (*x.Vals)[n]
		} else {
			(*x.Out)[n] = x.Scl
		}
	}
}
func (x *StmBse) SclMax(scl flt.Flt) Stm {
	r := &StmSclMax{}
	r.Slf = r
	r.Stm = x.Slf
	r.Scl = scl
	if ana.Cfg.Trc.IsHstStm() {
		sys.Logf("%p hst.StmSclMax(%v)", r, r.Prm())
	}
	if x.Vals.Cnt() == 0 {
		return r
	}
	if len(*x.Tmes) != len(*x.Vals) {
		err.Panicf("length unequal (Tmes:%v Vals:%v)", len(*x.Tmes), len(*x.Vals))
	}
	r.Tmes = x.Tmes
	r.Vals = flts.Make(x.Vals.Cnt())
	segBnds, acts := bnds.Segs(r.Vals.Cnt())
	for n, segBnd := range *segBnds {
		seg := &StmSclMaxSeg{}
		seg.Bnd = segBnd
		seg.Vals = x.Vals
		seg.Out = r.Vals
		seg.Scl = r.Scl
		acts[n] = seg
	}
	sys.Run().Pll(acts...)
	return r
}
func (x *StmSclMax) Name() str.Str             { return str.Str("SclMax") }
func (x *StmSclMax) PrmWrt(b *strings.Builder) { x.Scl.StrWrt(b) }
func (x *StmSclMax) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *StmSclMax) StrWrt(b *strings.Builder) {
	x.Stm.StrWrt(b)
	b.WriteString(".sclMax(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *StmSclMax) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *StmSclMaxSeg) Act() {
	for n := x.Idx; n < x.Lim; n++ {
		if (*x.Vals)[n] > x.Scl {
			(*x.Out)[n] = (*x.Vals)[n]
		} else {
			(*x.Out)[n] = x.Scl
		}
	}
}
func (x *StmBse) SelEql(sel flt.Flt) Stm {
	r := &StmSelEql{}
	r.Slf = r
	r.Stm = x.Slf
	r.Sel = sel
	if ana.Cfg.Trc.IsHstStm() {
		sys.Logf("%p hst.StmSelEql(%v)", r, r.Prm())
	}
	if x.Vals.Cnt() == 0 {
		return r
	}
	if len(*x.Tmes) != len(*x.Vals) {
		err.Panicf("length unequal (Tmes:%v Vals:%v)", len(*x.Tmes), len(*x.Vals))
	}
	r.Tmes = x.Tmes
	r.Vals = flts.Make(x.Vals.Cnt())
	segBnds, acts := bnds.Segs(r.Vals.Cnt())
	for n, segBnd := range *segBnds {
		seg := &StmSelEqlSeg{}
		seg.Bnd = segBnd
		seg.Vals = x.Vals
		seg.Out = r.Vals
		seg.Sel = r.Sel
		acts[n] = seg
	}
	sys.Run().Pll(acts...)
	return r
}
func (x *StmSelEql) Name() str.Str             { return str.Str("SelEql") }
func (x *StmSelEql) PrmWrt(b *strings.Builder) { x.Sel.StrWrt(b) }
func (x *StmSelEql) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *StmSelEql) StrWrt(b *strings.Builder) {
	x.Stm.StrWrt(b)
	b.WriteString(".selEql(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *StmSelEql) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *StmSelEqlSeg) Act() {
	for n := x.Idx; n < x.Lim; n++ {
		if (*x.Vals)[n] == x.Sel {
			(*x.Out)[n] = (*x.Vals)[n]
		} else {
			(*x.Out)[n] = 0
		}
	}
}
func (x *StmBse) SelNeq(sel flt.Flt) Stm {
	r := &StmSelNeq{}
	r.Slf = r
	r.Stm = x.Slf
	r.Sel = sel
	if ana.Cfg.Trc.IsHstStm() {
		sys.Logf("%p hst.StmSelNeq(%v)", r, r.Prm())
	}
	if x.Vals.Cnt() == 0 {
		return r
	}
	if len(*x.Tmes) != len(*x.Vals) {
		err.Panicf("length unequal (Tmes:%v Vals:%v)", len(*x.Tmes), len(*x.Vals))
	}
	r.Tmes = x.Tmes
	r.Vals = flts.Make(x.Vals.Cnt())
	segBnds, acts := bnds.Segs(r.Vals.Cnt())
	for n, segBnd := range *segBnds {
		seg := &StmSelNeqSeg{}
		seg.Bnd = segBnd
		seg.Vals = x.Vals
		seg.Out = r.Vals
		seg.Sel = r.Sel
		acts[n] = seg
	}
	sys.Run().Pll(acts...)
	return r
}
func (x *StmSelNeq) Name() str.Str             { return str.Str("SelNeq") }
func (x *StmSelNeq) PrmWrt(b *strings.Builder) { x.Sel.StrWrt(b) }
func (x *StmSelNeq) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *StmSelNeq) StrWrt(b *strings.Builder) {
	x.Stm.StrWrt(b)
	b.WriteString(".selNeq(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *StmSelNeq) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *StmSelNeqSeg) Act() {
	for n := x.Idx; n < x.Lim; n++ {
		if (*x.Vals)[n] != x.Sel {
			(*x.Out)[n] = (*x.Vals)[n]
		} else {
			(*x.Out)[n] = 0
		}
	}
}
func (x *StmBse) SelLss(sel flt.Flt) Stm {
	r := &StmSelLss{}
	r.Slf = r
	r.Stm = x.Slf
	r.Sel = sel
	if ana.Cfg.Trc.IsHstStm() {
		sys.Logf("%p hst.StmSelLss(%v)", r, r.Prm())
	}
	if x.Vals.Cnt() == 0 {
		return r
	}
	if len(*x.Tmes) != len(*x.Vals) {
		err.Panicf("length unequal (Tmes:%v Vals:%v)", len(*x.Tmes), len(*x.Vals))
	}
	r.Tmes = x.Tmes
	r.Vals = flts.Make(x.Vals.Cnt())
	segBnds, acts := bnds.Segs(r.Vals.Cnt())
	for n, segBnd := range *segBnds {
		seg := &StmSelLssSeg{}
		seg.Bnd = segBnd
		seg.Vals = x.Vals
		seg.Out = r.Vals
		seg.Sel = r.Sel
		acts[n] = seg
	}
	sys.Run().Pll(acts...)
	return r
}
func (x *StmSelLss) Name() str.Str             { return str.Str("SelLss") }
func (x *StmSelLss) PrmWrt(b *strings.Builder) { x.Sel.StrWrt(b) }
func (x *StmSelLss) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *StmSelLss) StrWrt(b *strings.Builder) {
	x.Stm.StrWrt(b)
	b.WriteString(".selLss(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *StmSelLss) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *StmSelLssSeg) Act() {
	for n := x.Idx; n < x.Lim; n++ {
		if (*x.Vals)[n] < x.Sel {
			(*x.Out)[n] = (*x.Vals)[n]
		} else {
			(*x.Out)[n] = 0
		}
	}
}
func (x *StmBse) SelGtr(sel flt.Flt) Stm {
	r := &StmSelGtr{}
	r.Slf = r
	r.Stm = x.Slf
	r.Sel = sel
	if ana.Cfg.Trc.IsHstStm() {
		sys.Logf("%p hst.StmSelGtr(%v)", r, r.Prm())
	}
	if x.Vals.Cnt() == 0 {
		return r
	}
	if len(*x.Tmes) != len(*x.Vals) {
		err.Panicf("length unequal (Tmes:%v Vals:%v)", len(*x.Tmes), len(*x.Vals))
	}
	r.Tmes = x.Tmes
	r.Vals = flts.Make(x.Vals.Cnt())
	segBnds, acts := bnds.Segs(r.Vals.Cnt())
	for n, segBnd := range *segBnds {
		seg := &StmSelGtrSeg{}
		seg.Bnd = segBnd
		seg.Vals = x.Vals
		seg.Out = r.Vals
		seg.Sel = r.Sel
		acts[n] = seg
	}
	sys.Run().Pll(acts...)
	return r
}
func (x *StmSelGtr) Name() str.Str             { return str.Str("SelGtr") }
func (x *StmSelGtr) PrmWrt(b *strings.Builder) { x.Sel.StrWrt(b) }
func (x *StmSelGtr) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *StmSelGtr) StrWrt(b *strings.Builder) {
	x.Stm.StrWrt(b)
	b.WriteString(".selGtr(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *StmSelGtr) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *StmSelGtrSeg) Act() {
	for n := x.Idx; n < x.Lim; n++ {
		if (*x.Vals)[n] > x.Sel {
			(*x.Out)[n] = (*x.Vals)[n]
		} else {
			(*x.Out)[n] = 0
		}
	}
}
func (x *StmBse) SelLeq(sel flt.Flt) Stm {
	r := &StmSelLeq{}
	r.Slf = r
	r.Stm = x.Slf
	r.Sel = sel
	if ana.Cfg.Trc.IsHstStm() {
		sys.Logf("%p hst.StmSelLeq(%v)", r, r.Prm())
	}
	if x.Vals.Cnt() == 0 {
		return r
	}
	if len(*x.Tmes) != len(*x.Vals) {
		err.Panicf("length unequal (Tmes:%v Vals:%v)", len(*x.Tmes), len(*x.Vals))
	}
	r.Tmes = x.Tmes
	r.Vals = flts.Make(x.Vals.Cnt())
	segBnds, acts := bnds.Segs(r.Vals.Cnt())
	for n, segBnd := range *segBnds {
		seg := &StmSelLeqSeg{}
		seg.Bnd = segBnd
		seg.Vals = x.Vals
		seg.Out = r.Vals
		seg.Sel = r.Sel
		acts[n] = seg
	}
	sys.Run().Pll(acts...)
	return r
}
func (x *StmSelLeq) Name() str.Str             { return str.Str("SelLeq") }
func (x *StmSelLeq) PrmWrt(b *strings.Builder) { x.Sel.StrWrt(b) }
func (x *StmSelLeq) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *StmSelLeq) StrWrt(b *strings.Builder) {
	x.Stm.StrWrt(b)
	b.WriteString(".selLeq(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *StmSelLeq) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *StmSelLeqSeg) Act() {
	for n := x.Idx; n < x.Lim; n++ {
		if (*x.Vals)[n] <= x.Sel {
			(*x.Out)[n] = (*x.Vals)[n]
		} else {
			(*x.Out)[n] = 0
		}
	}
}
func (x *StmBse) SelGeq(sel flt.Flt) Stm {
	r := &StmSelGeq{}
	r.Slf = r
	r.Stm = x.Slf
	r.Sel = sel
	if ana.Cfg.Trc.IsHstStm() {
		sys.Logf("%p hst.StmSelGeq(%v)", r, r.Prm())
	}
	if x.Vals.Cnt() == 0 {
		return r
	}
	if len(*x.Tmes) != len(*x.Vals) {
		err.Panicf("length unequal (Tmes:%v Vals:%v)", len(*x.Tmes), len(*x.Vals))
	}
	r.Tmes = x.Tmes
	r.Vals = flts.Make(x.Vals.Cnt())
	segBnds, acts := bnds.Segs(r.Vals.Cnt())
	for n, segBnd := range *segBnds {
		seg := &StmSelGeqSeg{}
		seg.Bnd = segBnd
		seg.Vals = x.Vals
		seg.Out = r.Vals
		seg.Sel = r.Sel
		acts[n] = seg
	}
	sys.Run().Pll(acts...)
	return r
}
func (x *StmSelGeq) Name() str.Str             { return str.Str("SelGeq") }
func (x *StmSelGeq) PrmWrt(b *strings.Builder) { x.Sel.StrWrt(b) }
func (x *StmSelGeq) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *StmSelGeq) StrWrt(b *strings.Builder) {
	x.Stm.StrWrt(b)
	b.WriteString(".selGeq(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *StmSelGeq) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *StmSelGeqSeg) Act() {
	for n := x.Idx; n < x.Lim; n++ {
		if (*x.Vals)[n] >= x.Sel {
			(*x.Out)[n] = (*x.Vals)[n]
		} else {
			(*x.Out)[n] = 0
		}
	}
}
func (x *StmBse) AggFst(length unt.Unt) Stm {
	r := &StmAggFst{}
	r.Slf = r
	r.Stm = x.Slf
	r.Length = length
	if ana.Cfg.Trc.IsHstStm() {
		sys.Logf("%p hst.StmAggFst(%v)", r, r.Prm())
	}
	if x.Vals.Cnt() == 0 || x.Vals.Cnt() < r.Length || r.Length == 0 {
		return r
	}
	if len(*x.Tmes) != len(*x.Vals) {
		err.Panicf("length unequal (Tmes:%v Vals:%v)", len(*x.Tmes), len(*x.Vals))
	}
	r.Tmes = x.Tmes.From(r.Length - 1)
	r.Vals = flts.Make(x.Vals.Cnt() - r.Length + 1)
	segBnds, acts := bnds.Segs(r.Vals.Cnt())
	for n, segBnd := range *segBnds {
		seg := &StmAggFstSeg{}
		seg.Bnd = segBnd
		seg.Vals = x.Vals
		seg.Out = r.Vals
		seg.Length = r.Length
		acts[n] = seg
	}
	sys.Run().Pll(acts...)
	return r
}
func (x *StmAggFst) Name() str.Str             { return str.Str("AggFst") }
func (x *StmAggFst) PrmWrt(b *strings.Builder) { x.Length.StrWrt(b) }
func (x *StmAggFst) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *StmAggFst) StrWrt(b *strings.Builder) {
	x.Stm.StrWrt(b)
	b.WriteString(".aggFst(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *StmAggFst) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *StmAggFstSeg) Act() {
	for n := x.Idx; n < x.Lim; n++ {
		(*x.Out)[n] = x.Vals.In(n, n+x.Length).Fst()
	}
}
func (x *StmBse) AggLst(length unt.Unt) Stm {
	r := &StmAggLst{}
	r.Slf = r
	r.Stm = x.Slf
	r.Length = length
	if ana.Cfg.Trc.IsHstStm() {
		sys.Logf("%p hst.StmAggLst(%v)", r, r.Prm())
	}
	if x.Vals.Cnt() == 0 || x.Vals.Cnt() < r.Length || r.Length == 0 {
		return r
	}
	if len(*x.Tmes) != len(*x.Vals) {
		err.Panicf("length unequal (Tmes:%v Vals:%v)", len(*x.Tmes), len(*x.Vals))
	}
	r.Tmes = x.Tmes.From(r.Length - 1)
	r.Vals = flts.Make(x.Vals.Cnt() - r.Length + 1)
	segBnds, acts := bnds.Segs(r.Vals.Cnt())
	for n, segBnd := range *segBnds {
		seg := &StmAggLstSeg{}
		seg.Bnd = segBnd
		seg.Vals = x.Vals
		seg.Out = r.Vals
		seg.Length = r.Length
		acts[n] = seg
	}
	sys.Run().Pll(acts...)
	return r
}
func (x *StmAggLst) Name() str.Str             { return str.Str("AggLst") }
func (x *StmAggLst) PrmWrt(b *strings.Builder) { x.Length.StrWrt(b) }
func (x *StmAggLst) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *StmAggLst) StrWrt(b *strings.Builder) {
	x.Stm.StrWrt(b)
	b.WriteString(".aggLst(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *StmAggLst) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *StmAggLstSeg) Act() {
	for n := x.Idx; n < x.Lim; n++ {
		(*x.Out)[n] = x.Vals.In(n, n+x.Length).Lst()
	}
}
func (x *StmBse) AggSum(length unt.Unt) Stm {
	r := &StmAggSum{}
	r.Slf = r
	r.Stm = x.Slf
	r.Length = length
	if ana.Cfg.Trc.IsHstStm() {
		sys.Logf("%p hst.StmAggSum(%v)", r, r.Prm())
	}
	if x.Vals.Cnt() == 0 || x.Vals.Cnt() < r.Length || r.Length == 0 {
		return r
	}
	if len(*x.Tmes) != len(*x.Vals) {
		err.Panicf("length unequal (Tmes:%v Vals:%v)", len(*x.Tmes), len(*x.Vals))
	}
	r.Tmes = x.Tmes.From(r.Length - 1)
	r.Vals = flts.Make(x.Vals.Cnt() - r.Length + 1)
	segBnds, acts := bnds.Segs(r.Vals.Cnt())
	for n, segBnd := range *segBnds {
		seg := &StmAggSumSeg{}
		seg.Bnd = segBnd
		seg.Vals = x.Vals
		seg.Out = r.Vals
		seg.Length = r.Length
		acts[n] = seg
	}
	sys.Run().Pll(acts...)
	return r
}
func (x *StmAggSum) Name() str.Str             { return str.Str("AggSum") }
func (x *StmAggSum) PrmWrt(b *strings.Builder) { x.Length.StrWrt(b) }
func (x *StmAggSum) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *StmAggSum) StrWrt(b *strings.Builder) {
	x.Stm.StrWrt(b)
	b.WriteString(".aggSum(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *StmAggSum) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *StmAggSumSeg) Act() {
	for n := x.Idx; n < x.Lim; n++ {
		(*x.Out)[n] = x.Vals.In(n, n+x.Length).Sum()
	}
}
func (x *StmBse) AggPrd(length unt.Unt) Stm {
	r := &StmAggPrd{}
	r.Slf = r
	r.Stm = x.Slf
	r.Length = length
	if ana.Cfg.Trc.IsHstStm() {
		sys.Logf("%p hst.StmAggPrd(%v)", r, r.Prm())
	}
	if x.Vals.Cnt() == 0 || x.Vals.Cnt() < r.Length || r.Length == 0 {
		return r
	}
	if len(*x.Tmes) != len(*x.Vals) {
		err.Panicf("length unequal (Tmes:%v Vals:%v)", len(*x.Tmes), len(*x.Vals))
	}
	r.Tmes = x.Tmes.From(r.Length - 1)
	r.Vals = flts.Make(x.Vals.Cnt() - r.Length + 1)
	segBnds, acts := bnds.Segs(r.Vals.Cnt())
	for n, segBnd := range *segBnds {
		seg := &StmAggPrdSeg{}
		seg.Bnd = segBnd
		seg.Vals = x.Vals
		seg.Out = r.Vals
		seg.Length = r.Length
		acts[n] = seg
	}
	sys.Run().Pll(acts...)
	return r
}
func (x *StmAggPrd) Name() str.Str             { return str.Str("AggPrd") }
func (x *StmAggPrd) PrmWrt(b *strings.Builder) { x.Length.StrWrt(b) }
func (x *StmAggPrd) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *StmAggPrd) StrWrt(b *strings.Builder) {
	x.Stm.StrWrt(b)
	b.WriteString(".aggPrd(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *StmAggPrd) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *StmAggPrdSeg) Act() {
	for n := x.Idx; n < x.Lim; n++ {
		(*x.Out)[n] = x.Vals.In(n, n+x.Length).Prd()
	}
}
func (x *StmBse) AggMin(length unt.Unt) Stm {
	r := &StmAggMin{}
	r.Slf = r
	r.Stm = x.Slf
	r.Length = length
	if ana.Cfg.Trc.IsHstStm() {
		sys.Logf("%p hst.StmAggMin(%v)", r, r.Prm())
	}
	if x.Vals.Cnt() == 0 || x.Vals.Cnt() < r.Length || r.Length == 0 {
		return r
	}
	if len(*x.Tmes) != len(*x.Vals) {
		err.Panicf("length unequal (Tmes:%v Vals:%v)", len(*x.Tmes), len(*x.Vals))
	}
	r.Tmes = x.Tmes.From(r.Length - 1)
	r.Vals = flts.Make(x.Vals.Cnt() - r.Length + 1)
	segBnds, acts := bnds.Segs(r.Vals.Cnt())
	for n, segBnd := range *segBnds {
		seg := &StmAggMinSeg{}
		seg.Bnd = segBnd
		seg.Vals = x.Vals
		seg.Out = r.Vals
		seg.Length = r.Length
		acts[n] = seg
	}
	sys.Run().Pll(acts...)
	return r
}
func (x *StmAggMin) Name() str.Str             { return str.Str("AggMin") }
func (x *StmAggMin) PrmWrt(b *strings.Builder) { x.Length.StrWrt(b) }
func (x *StmAggMin) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *StmAggMin) StrWrt(b *strings.Builder) {
	x.Stm.StrWrt(b)
	b.WriteString(".aggMin(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *StmAggMin) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *StmAggMinSeg) Act() {
	for n := x.Idx; n < x.Lim; n++ {
		(*x.Out)[n] = x.Vals.In(n, n+x.Length).Min()
	}
}
func (x *StmBse) AggMax(length unt.Unt) Stm {
	r := &StmAggMax{}
	r.Slf = r
	r.Stm = x.Slf
	r.Length = length
	if ana.Cfg.Trc.IsHstStm() {
		sys.Logf("%p hst.StmAggMax(%v)", r, r.Prm())
	}
	if x.Vals.Cnt() == 0 || x.Vals.Cnt() < r.Length || r.Length == 0 {
		return r
	}
	if len(*x.Tmes) != len(*x.Vals) {
		err.Panicf("length unequal (Tmes:%v Vals:%v)", len(*x.Tmes), len(*x.Vals))
	}
	r.Tmes = x.Tmes.From(r.Length - 1)
	r.Vals = flts.Make(x.Vals.Cnt() - r.Length + 1)
	segBnds, acts := bnds.Segs(r.Vals.Cnt())
	for n, segBnd := range *segBnds {
		seg := &StmAggMaxSeg{}
		seg.Bnd = segBnd
		seg.Vals = x.Vals
		seg.Out = r.Vals
		seg.Length = r.Length
		acts[n] = seg
	}
	sys.Run().Pll(acts...)
	return r
}
func (x *StmAggMax) Name() str.Str             { return str.Str("AggMax") }
func (x *StmAggMax) PrmWrt(b *strings.Builder) { x.Length.StrWrt(b) }
func (x *StmAggMax) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *StmAggMax) StrWrt(b *strings.Builder) {
	x.Stm.StrWrt(b)
	b.WriteString(".aggMax(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *StmAggMax) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *StmAggMaxSeg) Act() {
	for n := x.Idx; n < x.Lim; n++ {
		(*x.Out)[n] = x.Vals.In(n, n+x.Length).Max()
	}
}
func (x *StmBse) AggMid(length unt.Unt) Stm {
	r := &StmAggMid{}
	r.Slf = r
	r.Stm = x.Slf
	r.Length = length
	if ana.Cfg.Trc.IsHstStm() {
		sys.Logf("%p hst.StmAggMid(%v)", r, r.Prm())
	}
	if x.Vals.Cnt() == 0 || x.Vals.Cnt() < r.Length || r.Length == 0 {
		return r
	}
	if len(*x.Tmes) != len(*x.Vals) {
		err.Panicf("length unequal (Tmes:%v Vals:%v)", len(*x.Tmes), len(*x.Vals))
	}
	r.Tmes = x.Tmes.From(r.Length - 1)
	r.Vals = flts.Make(x.Vals.Cnt() - r.Length + 1)
	segBnds, acts := bnds.Segs(r.Vals.Cnt())
	for n, segBnd := range *segBnds {
		seg := &StmAggMidSeg{}
		seg.Bnd = segBnd
		seg.Vals = x.Vals
		seg.Out = r.Vals
		seg.Length = r.Length
		acts[n] = seg
	}
	sys.Run().Pll(acts...)
	return r
}
func (x *StmAggMid) Name() str.Str             { return str.Str("AggMid") }
func (x *StmAggMid) PrmWrt(b *strings.Builder) { x.Length.StrWrt(b) }
func (x *StmAggMid) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *StmAggMid) StrWrt(b *strings.Builder) {
	x.Stm.StrWrt(b)
	b.WriteString(".aggMid(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *StmAggMid) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *StmAggMidSeg) Act() {
	for n := x.Idx; n < x.Lim; n++ {
		(*x.Out)[n] = x.Vals.In(n, n+x.Length).Mid()
	}
}
func (x *StmBse) AggMdn(length unt.Unt) Stm {
	r := &StmAggMdn{}
	r.Slf = r
	r.Stm = x.Slf
	r.Length = length
	if ana.Cfg.Trc.IsHstStm() {
		sys.Logf("%p hst.StmAggMdn(%v)", r, r.Prm())
	}
	if x.Vals.Cnt() == 0 || x.Vals.Cnt() < r.Length || r.Length == 0 {
		return r
	}
	if len(*x.Tmes) != len(*x.Vals) {
		err.Panicf("length unequal (Tmes:%v Vals:%v)", len(*x.Tmes), len(*x.Vals))
	}
	r.Tmes = x.Tmes.From(r.Length - 1)
	r.Vals = flts.Make(x.Vals.Cnt() - r.Length + 1)
	segBnds, acts := bnds.Segs(r.Vals.Cnt())
	for n, segBnd := range *segBnds {
		seg := &StmAggMdnSeg{}
		seg.Bnd = segBnd
		seg.Vals = x.Vals
		seg.Out = r.Vals
		seg.Length = r.Length
		acts[n] = seg
	}
	sys.Run().Pll(acts...)
	return r
}
func (x *StmAggMdn) Name() str.Str             { return str.Str("AggMdn") }
func (x *StmAggMdn) PrmWrt(b *strings.Builder) { x.Length.StrWrt(b) }
func (x *StmAggMdn) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *StmAggMdn) StrWrt(b *strings.Builder) {
	x.Stm.StrWrt(b)
	b.WriteString(".aggMdn(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *StmAggMdn) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *StmAggMdnSeg) Act() {
	for n := x.Idx; n < x.Lim; n++ {
		(*x.Out)[n] = x.Vals.In(n, n+x.Length).Mdn()
	}
}
func (x *StmBse) AggSma(length unt.Unt) Stm {
	r := &StmAggSma{}
	r.Slf = r
	r.Stm = x.Slf
	r.Length = length
	if ana.Cfg.Trc.IsHstStm() {
		sys.Logf("%p hst.StmAggSma(%v)", r, r.Prm())
	}
	if x.Vals.Cnt() == 0 || x.Vals.Cnt() < r.Length || r.Length == 0 {
		return r
	}
	if len(*x.Tmes) != len(*x.Vals) {
		err.Panicf("length unequal (Tmes:%v Vals:%v)", len(*x.Tmes), len(*x.Vals))
	}
	r.Tmes = x.Tmes.From(r.Length - 1)
	r.Vals = flts.Make(x.Vals.Cnt() - r.Length + 1)
	segBnds, acts := bnds.Segs(r.Vals.Cnt())
	for n, segBnd := range *segBnds {
		seg := &StmAggSmaSeg{}
		seg.Bnd = segBnd
		seg.Vals = x.Vals
		seg.Out = r.Vals
		seg.Length = r.Length
		acts[n] = seg
	}
	sys.Run().Pll(acts...)
	return r
}
func (x *StmAggSma) Name() str.Str             { return str.Str("AggSma") }
func (x *StmAggSma) PrmWrt(b *strings.Builder) { x.Length.StrWrt(b) }
func (x *StmAggSma) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *StmAggSma) StrWrt(b *strings.Builder) {
	x.Stm.StrWrt(b)
	b.WriteString(".aggSma(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *StmAggSma) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *StmAggSmaSeg) Act() {
	for n := x.Idx; n < x.Lim; n++ {
		(*x.Out)[n] = x.Vals.In(n, n+x.Length).Sma()
	}
}
func (x *StmBse) AggGma(length unt.Unt) Stm {
	r := &StmAggGma{}
	r.Slf = r
	r.Stm = x.Slf
	r.Length = length
	if ana.Cfg.Trc.IsHstStm() {
		sys.Logf("%p hst.StmAggGma(%v)", r, r.Prm())
	}
	if x.Vals.Cnt() == 0 || x.Vals.Cnt() < r.Length || r.Length == 0 {
		return r
	}
	if len(*x.Tmes) != len(*x.Vals) {
		err.Panicf("length unequal (Tmes:%v Vals:%v)", len(*x.Tmes), len(*x.Vals))
	}
	r.Tmes = x.Tmes.From(r.Length - 1)
	r.Vals = flts.Make(x.Vals.Cnt() - r.Length + 1)
	segBnds, acts := bnds.Segs(r.Vals.Cnt())
	for n, segBnd := range *segBnds {
		seg := &StmAggGmaSeg{}
		seg.Bnd = segBnd
		seg.Vals = x.Vals
		seg.Out = r.Vals
		seg.Length = r.Length
		acts[n] = seg
	}
	sys.Run().Pll(acts...)
	return r
}
func (x *StmAggGma) Name() str.Str             { return str.Str("AggGma") }
func (x *StmAggGma) PrmWrt(b *strings.Builder) { x.Length.StrWrt(b) }
func (x *StmAggGma) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *StmAggGma) StrWrt(b *strings.Builder) {
	x.Stm.StrWrt(b)
	b.WriteString(".aggGma(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *StmAggGma) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *StmAggGmaSeg) Act() {
	for n := x.Idx; n < x.Lim; n++ {
		(*x.Out)[n] = x.Vals.In(n, n+x.Length).Gma()
	}
}
func (x *StmBse) AggWma(length unt.Unt) Stm {
	r := &StmAggWma{}
	r.Slf = r
	r.Stm = x.Slf
	r.Length = length
	if ana.Cfg.Trc.IsHstStm() {
		sys.Logf("%p hst.StmAggWma(%v)", r, r.Prm())
	}
	if x.Vals.Cnt() == 0 || x.Vals.Cnt() < r.Length || r.Length == 0 {
		return r
	}
	if len(*x.Tmes) != len(*x.Vals) {
		err.Panicf("length unequal (Tmes:%v Vals:%v)", len(*x.Tmes), len(*x.Vals))
	}
	r.Tmes = x.Tmes.From(r.Length - 1)
	r.Vals = flts.Make(x.Vals.Cnt() - r.Length + 1)
	segBnds, acts := bnds.Segs(r.Vals.Cnt())
	for n, segBnd := range *segBnds {
		seg := &StmAggWmaSeg{}
		seg.Bnd = segBnd
		seg.Vals = x.Vals
		seg.Out = r.Vals
		seg.Length = r.Length
		acts[n] = seg
	}
	sys.Run().Pll(acts...)
	return r
}
func (x *StmAggWma) Name() str.Str             { return str.Str("AggWma") }
func (x *StmAggWma) PrmWrt(b *strings.Builder) { x.Length.StrWrt(b) }
func (x *StmAggWma) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *StmAggWma) StrWrt(b *strings.Builder) {
	x.Stm.StrWrt(b)
	b.WriteString(".aggWma(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *StmAggWma) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *StmAggWmaSeg) Act() {
	for n := x.Idx; n < x.Lim; n++ {
		(*x.Out)[n] = x.Vals.In(n, n+x.Length).Wma()
	}
}
func (x *StmBse) AggRsi(length unt.Unt) Stm {
	r := &StmAggRsi{}
	r.Slf = r
	r.Stm = x.Slf
	r.Length = length
	if ana.Cfg.Trc.IsHstStm() {
		sys.Logf("%p hst.StmAggRsi(%v)", r, r.Prm())
	}
	if x.Vals.Cnt() == 0 || x.Vals.Cnt() < r.Length || r.Length == 0 {
		return r
	}
	if len(*x.Tmes) != len(*x.Vals) {
		err.Panicf("length unequal (Tmes:%v Vals:%v)", len(*x.Tmes), len(*x.Vals))
	}
	r.Tmes = x.Tmes.From(r.Length - 1)
	r.Vals = flts.Make(x.Vals.Cnt() - r.Length + 1)
	segBnds, acts := bnds.Segs(r.Vals.Cnt())
	for n, segBnd := range *segBnds {
		seg := &StmAggRsiSeg{}
		seg.Bnd = segBnd
		seg.Vals = x.Vals
		seg.Out = r.Vals
		seg.Length = r.Length
		acts[n] = seg
	}
	sys.Run().Pll(acts...)
	return r
}
func (x *StmAggRsi) Name() str.Str             { return str.Str("AggRsi") }
func (x *StmAggRsi) PrmWrt(b *strings.Builder) { x.Length.StrWrt(b) }
func (x *StmAggRsi) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *StmAggRsi) StrWrt(b *strings.Builder) {
	x.Stm.StrWrt(b)
	b.WriteString(".aggRsi(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *StmAggRsi) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *StmAggRsiSeg) Act() {
	for n := x.Idx; n < x.Lim; n++ {
		(*x.Out)[n] = x.Vals.In(n, n+x.Length).Rsi()
	}
}
func (x *StmBse) AggWrsi(length unt.Unt) Stm {
	r := &StmAggWrsi{}
	r.Slf = r
	r.Stm = x.Slf
	r.Length = length
	if ana.Cfg.Trc.IsHstStm() {
		sys.Logf("%p hst.StmAggWrsi(%v)", r, r.Prm())
	}
	if x.Vals.Cnt() == 0 || x.Vals.Cnt() < r.Length || r.Length == 0 {
		return r
	}
	if len(*x.Tmes) != len(*x.Vals) {
		err.Panicf("length unequal (Tmes:%v Vals:%v)", len(*x.Tmes), len(*x.Vals))
	}
	r.Tmes = x.Tmes.From(r.Length - 1)
	r.Vals = flts.Make(x.Vals.Cnt() - r.Length + 1)
	segBnds, acts := bnds.Segs(r.Vals.Cnt())
	for n, segBnd := range *segBnds {
		seg := &StmAggWrsiSeg{}
		seg.Bnd = segBnd
		seg.Vals = x.Vals
		seg.Out = r.Vals
		seg.Length = r.Length
		acts[n] = seg
	}
	sys.Run().Pll(acts...)
	return r
}
func (x *StmAggWrsi) Name() str.Str             { return str.Str("AggWrsi") }
func (x *StmAggWrsi) PrmWrt(b *strings.Builder) { x.Length.StrWrt(b) }
func (x *StmAggWrsi) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *StmAggWrsi) StrWrt(b *strings.Builder) {
	x.Stm.StrWrt(b)
	b.WriteString(".aggWrsi(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *StmAggWrsi) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *StmAggWrsiSeg) Act() {
	for n := x.Idx; n < x.Lim; n++ {
		(*x.Out)[n] = x.Vals.In(n, n+x.Length).Wrsi()
	}
}
func (x *StmBse) AggAlma(length unt.Unt) Stm {
	r := &StmAggAlma{}
	r.Slf = r
	r.Stm = x.Slf
	r.Length = length
	if ana.Cfg.Trc.IsHstStm() {
		sys.Logf("%p hst.StmAggAlma(%v)", r, r.Prm())
	}
	if x.Vals.Cnt() == 0 || x.Vals.Cnt() < r.Length || r.Length == 0 {
		return r
	}
	if len(*x.Tmes) != len(*x.Vals) {
		err.Panicf("length unequal (Tmes:%v Vals:%v)", len(*x.Tmes), len(*x.Vals))
	}
	r.Tmes = x.Tmes.From(r.Length - 1)
	r.Vals = flts.Make(x.Vals.Cnt() - r.Length + 1)
	segBnds, acts := bnds.Segs(r.Vals.Cnt())
	for n, segBnd := range *segBnds {
		seg := &StmAggAlmaSeg{}
		seg.Bnd = segBnd
		seg.Vals = x.Vals
		seg.Out = r.Vals
		seg.Length = r.Length
		acts[n] = seg
	}
	sys.Run().Pll(acts...)
	return r
}
func (x *StmAggAlma) Name() str.Str             { return str.Str("AggAlma") }
func (x *StmAggAlma) PrmWrt(b *strings.Builder) { x.Length.StrWrt(b) }
func (x *StmAggAlma) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *StmAggAlma) StrWrt(b *strings.Builder) {
	x.Stm.StrWrt(b)
	b.WriteString(".aggAlma(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *StmAggAlma) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *StmAggAlmaSeg) Act() {
	for n := x.Idx; n < x.Lim; n++ {
		(*x.Out)[n] = x.Vals.In(n, n+x.Length).Alma()
	}
}
func (x *StmBse) AggVrnc(length unt.Unt) Stm {
	r := &StmAggVrnc{}
	r.Slf = r
	r.Stm = x.Slf
	r.Length = length
	if ana.Cfg.Trc.IsHstStm() {
		sys.Logf("%p hst.StmAggVrnc(%v)", r, r.Prm())
	}
	if x.Vals.Cnt() == 0 || x.Vals.Cnt() < r.Length || r.Length == 0 {
		return r
	}
	if len(*x.Tmes) != len(*x.Vals) {
		err.Panicf("length unequal (Tmes:%v Vals:%v)", len(*x.Tmes), len(*x.Vals))
	}
	r.Tmes = x.Tmes.From(r.Length - 1)
	r.Vals = flts.Make(x.Vals.Cnt() - r.Length + 1)
	segBnds, acts := bnds.Segs(r.Vals.Cnt())
	for n, segBnd := range *segBnds {
		seg := &StmAggVrncSeg{}
		seg.Bnd = segBnd
		seg.Vals = x.Vals
		seg.Out = r.Vals
		seg.Length = r.Length
		acts[n] = seg
	}
	sys.Run().Pll(acts...)
	return r
}
func (x *StmAggVrnc) Name() str.Str             { return str.Str("AggVrnc") }
func (x *StmAggVrnc) PrmWrt(b *strings.Builder) { x.Length.StrWrt(b) }
func (x *StmAggVrnc) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *StmAggVrnc) StrWrt(b *strings.Builder) {
	x.Stm.StrWrt(b)
	b.WriteString(".aggVrnc(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *StmAggVrnc) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *StmAggVrncSeg) Act() {
	for n := x.Idx; n < x.Lim; n++ {
		(*x.Out)[n] = x.Vals.In(n, n+x.Length).Vrnc()
	}
}
func (x *StmBse) AggStd(length unt.Unt) Stm {
	r := &StmAggStd{}
	r.Slf = r
	r.Stm = x.Slf
	r.Length = length
	if ana.Cfg.Trc.IsHstStm() {
		sys.Logf("%p hst.StmAggStd(%v)", r, r.Prm())
	}
	if x.Vals.Cnt() == 0 || x.Vals.Cnt() < r.Length || r.Length == 0 {
		return r
	}
	if len(*x.Tmes) != len(*x.Vals) {
		err.Panicf("length unequal (Tmes:%v Vals:%v)", len(*x.Tmes), len(*x.Vals))
	}
	r.Tmes = x.Tmes.From(r.Length - 1)
	r.Vals = flts.Make(x.Vals.Cnt() - r.Length + 1)
	segBnds, acts := bnds.Segs(r.Vals.Cnt())
	for n, segBnd := range *segBnds {
		seg := &StmAggStdSeg{}
		seg.Bnd = segBnd
		seg.Vals = x.Vals
		seg.Out = r.Vals
		seg.Length = r.Length
		acts[n] = seg
	}
	sys.Run().Pll(acts...)
	return r
}
func (x *StmAggStd) Name() str.Str             { return str.Str("AggStd") }
func (x *StmAggStd) PrmWrt(b *strings.Builder) { x.Length.StrWrt(b) }
func (x *StmAggStd) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *StmAggStd) StrWrt(b *strings.Builder) {
	x.Stm.StrWrt(b)
	b.WriteString(".aggStd(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *StmAggStd) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *StmAggStdSeg) Act() {
	for n := x.Idx; n < x.Lim; n++ {
		(*x.Out)[n] = x.Vals.In(n, n+x.Length).Std()
	}
}
func (x *StmBse) AggRngFul(length unt.Unt) Stm {
	r := &StmAggRngFul{}
	r.Slf = r
	r.Stm = x.Slf
	r.Length = length
	if ana.Cfg.Trc.IsHstStm() {
		sys.Logf("%p hst.StmAggRngFul(%v)", r, r.Prm())
	}
	if x.Vals.Cnt() == 0 || x.Vals.Cnt() < r.Length || r.Length == 0 {
		return r
	}
	if len(*x.Tmes) != len(*x.Vals) {
		err.Panicf("length unequal (Tmes:%v Vals:%v)", len(*x.Tmes), len(*x.Vals))
	}
	r.Tmes = x.Tmes.From(r.Length - 1)
	r.Vals = flts.Make(x.Vals.Cnt() - r.Length + 1)
	segBnds, acts := bnds.Segs(r.Vals.Cnt())
	for n, segBnd := range *segBnds {
		seg := &StmAggRngFulSeg{}
		seg.Bnd = segBnd
		seg.Vals = x.Vals
		seg.Out = r.Vals
		seg.Length = r.Length
		acts[n] = seg
	}
	sys.Run().Pll(acts...)
	return r
}
func (x *StmAggRngFul) Name() str.Str             { return str.Str("AggRngFul") }
func (x *StmAggRngFul) PrmWrt(b *strings.Builder) { x.Length.StrWrt(b) }
func (x *StmAggRngFul) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *StmAggRngFul) StrWrt(b *strings.Builder) {
	x.Stm.StrWrt(b)
	b.WriteString(".aggRngFul(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *StmAggRngFul) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *StmAggRngFulSeg) Act() {
	for n := x.Idx; n < x.Lim; n++ {
		(*x.Out)[n] = x.Vals.In(n, n+x.Length).RngFul()
	}
}
func (x *StmBse) AggRngLst(length unt.Unt) Stm {
	r := &StmAggRngLst{}
	r.Slf = r
	r.Stm = x.Slf
	r.Length = length
	if ana.Cfg.Trc.IsHstStm() {
		sys.Logf("%p hst.StmAggRngLst(%v)", r, r.Prm())
	}
	if x.Vals.Cnt() == 0 || x.Vals.Cnt() < r.Length || r.Length == 0 {
		return r
	}
	if len(*x.Tmes) != len(*x.Vals) {
		err.Panicf("length unequal (Tmes:%v Vals:%v)", len(*x.Tmes), len(*x.Vals))
	}
	r.Tmes = x.Tmes.From(r.Length - 1)
	r.Vals = flts.Make(x.Vals.Cnt() - r.Length + 1)
	segBnds, acts := bnds.Segs(r.Vals.Cnt())
	for n, segBnd := range *segBnds {
		seg := &StmAggRngLstSeg{}
		seg.Bnd = segBnd
		seg.Vals = x.Vals
		seg.Out = r.Vals
		seg.Length = r.Length
		acts[n] = seg
	}
	sys.Run().Pll(acts...)
	return r
}
func (x *StmAggRngLst) Name() str.Str             { return str.Str("AggRngLst") }
func (x *StmAggRngLst) PrmWrt(b *strings.Builder) { x.Length.StrWrt(b) }
func (x *StmAggRngLst) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *StmAggRngLst) StrWrt(b *strings.Builder) {
	x.Stm.StrWrt(b)
	b.WriteString(".aggRngLst(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *StmAggRngLst) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *StmAggRngLstSeg) Act() {
	for n := x.Idx; n < x.Lim; n++ {
		(*x.Out)[n] = x.Vals.In(n, n+x.Length).RngLst()
	}
}
func (x *StmBse) AggProLst(length unt.Unt) Stm {
	r := &StmAggProLst{}
	r.Slf = r
	r.Stm = x.Slf
	r.Length = length
	if ana.Cfg.Trc.IsHstStm() {
		sys.Logf("%p hst.StmAggProLst(%v)", r, r.Prm())
	}
	if x.Vals.Cnt() == 0 || x.Vals.Cnt() < r.Length || r.Length == 0 {
		return r
	}
	if len(*x.Tmes) != len(*x.Vals) {
		err.Panicf("length unequal (Tmes:%v Vals:%v)", len(*x.Tmes), len(*x.Vals))
	}
	r.Tmes = x.Tmes.From(r.Length - 1)
	r.Vals = flts.Make(x.Vals.Cnt() - r.Length + 1)
	segBnds, acts := bnds.Segs(r.Vals.Cnt())
	for n, segBnd := range *segBnds {
		seg := &StmAggProLstSeg{}
		seg.Bnd = segBnd
		seg.Vals = x.Vals
		seg.Out = r.Vals
		seg.Length = r.Length
		acts[n] = seg
	}
	sys.Run().Pll(acts...)
	return r
}
func (x *StmAggProLst) Name() str.Str             { return str.Str("AggProLst") }
func (x *StmAggProLst) PrmWrt(b *strings.Builder) { x.Length.StrWrt(b) }
func (x *StmAggProLst) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *StmAggProLst) StrWrt(b *strings.Builder) {
	x.Stm.StrWrt(b)
	b.WriteString(".aggProLst(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *StmAggProLst) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *StmAggProLstSeg) Act() {
	for n := x.Idx; n < x.Lim; n++ {
		(*x.Out)[n] = x.Vals.In(n, n+x.Length).ProLst()
	}
}
func (x *StmBse) AggProSma(length unt.Unt) Stm {
	r := &StmAggProSma{}
	r.Slf = r
	r.Stm = x.Slf
	r.Length = length
	if ana.Cfg.Trc.IsHstStm() {
		sys.Logf("%p hst.StmAggProSma(%v)", r, r.Prm())
	}
	if x.Vals.Cnt() == 0 || x.Vals.Cnt() < r.Length || r.Length == 0 {
		return r
	}
	if len(*x.Tmes) != len(*x.Vals) {
		err.Panicf("length unequal (Tmes:%v Vals:%v)", len(*x.Tmes), len(*x.Vals))
	}
	r.Tmes = x.Tmes.From(r.Length - 1)
	r.Vals = flts.Make(x.Vals.Cnt() - r.Length + 1)
	segBnds, acts := bnds.Segs(r.Vals.Cnt())
	for n, segBnd := range *segBnds {
		seg := &StmAggProSmaSeg{}
		seg.Bnd = segBnd
		seg.Vals = x.Vals
		seg.Out = r.Vals
		seg.Length = r.Length
		acts[n] = seg
	}
	sys.Run().Pll(acts...)
	return r
}
func (x *StmAggProSma) Name() str.Str             { return str.Str("AggProSma") }
func (x *StmAggProSma) PrmWrt(b *strings.Builder) { x.Length.StrWrt(b) }
func (x *StmAggProSma) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *StmAggProSma) StrWrt(b *strings.Builder) {
	x.Stm.StrWrt(b)
	b.WriteString(".aggProSma(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *StmAggProSma) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *StmAggProSmaSeg) Act() {
	for n := x.Idx; n < x.Lim; n++ {
		(*x.Out)[n] = x.Vals.In(n, n+x.Length).ProSma()
	}
}
func (x *StmBse) AggProAlma(length unt.Unt) Stm {
	r := &StmAggProAlma{}
	r.Slf = r
	r.Stm = x.Slf
	r.Length = length
	if ana.Cfg.Trc.IsHstStm() {
		sys.Logf("%p hst.StmAggProAlma(%v)", r, r.Prm())
	}
	if x.Vals.Cnt() == 0 || x.Vals.Cnt() < r.Length || r.Length == 0 {
		return r
	}
	if len(*x.Tmes) != len(*x.Vals) {
		err.Panicf("length unequal (Tmes:%v Vals:%v)", len(*x.Tmes), len(*x.Vals))
	}
	r.Tmes = x.Tmes.From(r.Length - 1)
	r.Vals = flts.Make(x.Vals.Cnt() - r.Length + 1)
	segBnds, acts := bnds.Segs(r.Vals.Cnt())
	for n, segBnd := range *segBnds {
		seg := &StmAggProAlmaSeg{}
		seg.Bnd = segBnd
		seg.Vals = x.Vals
		seg.Out = r.Vals
		seg.Length = r.Length
		acts[n] = seg
	}
	sys.Run().Pll(acts...)
	return r
}
func (x *StmAggProAlma) Name() str.Str             { return str.Str("AggProAlma") }
func (x *StmAggProAlma) PrmWrt(b *strings.Builder) { x.Length.StrWrt(b) }
func (x *StmAggProAlma) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *StmAggProAlma) StrWrt(b *strings.Builder) {
	x.Stm.StrWrt(b)
	b.WriteString(".aggProAlma(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *StmAggProAlma) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *StmAggProAlmaSeg) Act() {
	for n := x.Idx; n < x.Lim; n++ {
		(*x.Out)[n] = x.Vals.In(n, n+x.Length).ProAlma()
	}
}
func (x *StmBse) AggEma(length unt.Unt) Stm {
	r := &StmAggEma{}
	r.Slf = r
	r.Stm = x.Slf
	r.Length = length
	if ana.Cfg.Trc.IsHstStm() {
		sys.Logf("%p hst.StmAggEma(%v)", r, r.Prm())
	}
	if x.Vals.Cnt() == 0 || x.Vals.Cnt() < r.Length || r.Length == 0 {
		return r
	}
	if len(*x.Tmes) != len(*x.Vals) {
		err.Panicf("length unequal (Tmes:%v Vals:%v)", len(*x.Tmes), len(*x.Vals))
	}
	r.Tmes = x.Tmes.From(r.Length - 1)
	r.Vals = flts.Make(x.Vals.Cnt() - r.Length + 1)
	if len(*r.Tmes) != len(*r.Vals) {
		err.Panicf("r length unequal (Tmes:%v Vals:%v)", len(*r.Tmes), len(*r.Vals))
	}
	// NON-PLL IMPL DUE TO PRV VAL CHAINING
	//    EMA CALC   https://stockcharts.com/school/doku.php?id=chart_school:technical_indicators:moving_averages
	// Initial SMA: 10-period sum / 10
	// Multiplier: (2 / (Time periods + 1) ) = (2 / (10 + 1) ) = 0.1818 (18.18%)
	// EMA: {Close - EMA(previous day)} x multiplier + EMA(previous day)
	scl := flt.Flt(2) / flt.Flt(r.Length+1)
	(*r.Vals)[0] = x.Vals.To(r.Length).Sma()
	for n := 1; n < len(*r.Vals); n++ {
		(*r.Vals)[n] = ((*x.Vals)[n+int(r.Length)-1]-(*r.Vals)[n-1])*scl + (*r.Vals)[n-1]
	}
	return r
}
func (x *StmAggEma) Name() str.Str             { return str.Str("AggEma") }
func (x *StmAggEma) PrmWrt(b *strings.Builder) { x.Length.StrWrt(b) }
func (x *StmAggEma) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *StmAggEma) StrWrt(b *strings.Builder) {
	x.Stm.StrWrt(b)
	b.WriteString(".aggEma(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *StmAggEma) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *StmBse) InrAdd(off unt.Unt) Stm {
	r := &StmInrAdd{}
	r.Slf = r
	r.Stm = x.Slf
	r.Off = off
	if ana.Cfg.Trc.IsHstStm() {
		sys.Logf("%p hst.StmInrAdd(%v)", r, r.Prm())
	}
	if x.Vals.Cnt() == 0 || x.Vals.Cnt() <= r.Off || r.Off == 0 {
		return r
	}
	if len(*x.Tmes) != len(*x.Vals) {
		err.Panicf("length unequal (Tmes:%v Vals:%v)", len(*x.Tmes), len(*x.Vals))
	}
	r.Tmes = x.Tmes.From(r.Off)
	r.Vals = flts.Make(x.Vals.Cnt() - r.Off)
	segBnds, acts := bnds.Segs(r.Vals.Cnt())
	for n, segBnd := range *segBnds {
		seg := &StmInrAddSeg{}
		seg.Bnd = segBnd
		seg.Vals = x.Vals
		seg.Out = r.Vals
		seg.Off = r.Off
		acts[n] = seg
	}
	sys.Run().Pll(acts...)
	return r
}
func (x *StmInrAdd) Name() str.Str             { return str.Str("InrAdd") }
func (x *StmInrAdd) PrmWrt(b *strings.Builder) { x.Off.StrWrt(b) }
func (x *StmInrAdd) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *StmInrAdd) StrWrt(b *strings.Builder) {
	x.Stm.StrWrt(b)
	b.WriteString(".inrAdd(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *StmInrAdd) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *StmInrAddSeg) Act() {
	for n := x.Idx; n < x.Lim; n++ {
		(*x.Out)[n] = x.Vals.At(n + x.Off).Add((*x.Vals)[n])
	}
}
func (x *StmBse) InrSub(off unt.Unt) Stm {
	r := &StmInrSub{}
	r.Slf = r
	r.Stm = x.Slf
	r.Off = off
	if ana.Cfg.Trc.IsHstStm() {
		sys.Logf("%p hst.StmInrSub(%v)", r, r.Prm())
	}
	if x.Vals.Cnt() == 0 || x.Vals.Cnt() <= r.Off || r.Off == 0 {
		return r
	}
	if len(*x.Tmes) != len(*x.Vals) {
		err.Panicf("length unequal (Tmes:%v Vals:%v)", len(*x.Tmes), len(*x.Vals))
	}
	r.Tmes = x.Tmes.From(r.Off)
	r.Vals = flts.Make(x.Vals.Cnt() - r.Off)
	segBnds, acts := bnds.Segs(r.Vals.Cnt())
	for n, segBnd := range *segBnds {
		seg := &StmInrSubSeg{}
		seg.Bnd = segBnd
		seg.Vals = x.Vals
		seg.Out = r.Vals
		seg.Off = r.Off
		acts[n] = seg
	}
	sys.Run().Pll(acts...)
	return r
}
func (x *StmInrSub) Name() str.Str             { return str.Str("InrSub") }
func (x *StmInrSub) PrmWrt(b *strings.Builder) { x.Off.StrWrt(b) }
func (x *StmInrSub) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *StmInrSub) StrWrt(b *strings.Builder) {
	x.Stm.StrWrt(b)
	b.WriteString(".inrSub(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *StmInrSub) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *StmInrSubSeg) Act() {
	for n := x.Idx; n < x.Lim; n++ {
		(*x.Out)[n] = x.Vals.At(n + x.Off).Sub((*x.Vals)[n])
	}
}
func (x *StmBse) InrMul(off unt.Unt) Stm {
	r := &StmInrMul{}
	r.Slf = r
	r.Stm = x.Slf
	r.Off = off
	if ana.Cfg.Trc.IsHstStm() {
		sys.Logf("%p hst.StmInrMul(%v)", r, r.Prm())
	}
	if x.Vals.Cnt() == 0 || x.Vals.Cnt() <= r.Off || r.Off == 0 {
		return r
	}
	if len(*x.Tmes) != len(*x.Vals) {
		err.Panicf("length unequal (Tmes:%v Vals:%v)", len(*x.Tmes), len(*x.Vals))
	}
	r.Tmes = x.Tmes.From(r.Off)
	r.Vals = flts.Make(x.Vals.Cnt() - r.Off)
	segBnds, acts := bnds.Segs(r.Vals.Cnt())
	for n, segBnd := range *segBnds {
		seg := &StmInrMulSeg{}
		seg.Bnd = segBnd
		seg.Vals = x.Vals
		seg.Out = r.Vals
		seg.Off = r.Off
		acts[n] = seg
	}
	sys.Run().Pll(acts...)
	return r
}
func (x *StmInrMul) Name() str.Str             { return str.Str("InrMul") }
func (x *StmInrMul) PrmWrt(b *strings.Builder) { x.Off.StrWrt(b) }
func (x *StmInrMul) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *StmInrMul) StrWrt(b *strings.Builder) {
	x.Stm.StrWrt(b)
	b.WriteString(".inrMul(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *StmInrMul) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *StmInrMulSeg) Act() {
	for n := x.Idx; n < x.Lim; n++ {
		(*x.Out)[n] = x.Vals.At(n + x.Off).Mul((*x.Vals)[n])
	}
}
func (x *StmBse) InrDiv(off unt.Unt) Stm {
	r := &StmInrDiv{}
	r.Slf = r
	r.Stm = x.Slf
	r.Off = off
	if ana.Cfg.Trc.IsHstStm() {
		sys.Logf("%p hst.StmInrDiv(%v)", r, r.Prm())
	}
	if x.Vals.Cnt() == 0 || x.Vals.Cnt() <= r.Off || r.Off == 0 {
		return r
	}
	if len(*x.Tmes) != len(*x.Vals) {
		err.Panicf("length unequal (Tmes:%v Vals:%v)", len(*x.Tmes), len(*x.Vals))
	}
	r.Tmes = x.Tmes.From(r.Off)
	r.Vals = flts.Make(x.Vals.Cnt() - r.Off)
	segBnds, acts := bnds.Segs(r.Vals.Cnt())
	for n, segBnd := range *segBnds {
		seg := &StmInrDivSeg{}
		seg.Bnd = segBnd
		seg.Vals = x.Vals
		seg.Out = r.Vals
		seg.Off = r.Off
		acts[n] = seg
	}
	sys.Run().Pll(acts...)
	return r
}
func (x *StmInrDiv) Name() str.Str             { return str.Str("InrDiv") }
func (x *StmInrDiv) PrmWrt(b *strings.Builder) { x.Off.StrWrt(b) }
func (x *StmInrDiv) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *StmInrDiv) StrWrt(b *strings.Builder) {
	x.Stm.StrWrt(b)
	b.WriteString(".inrDiv(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *StmInrDiv) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *StmInrDivSeg) Act() {
	for n := x.Idx; n < x.Lim; n++ {
		(*x.Out)[n] = x.Vals.At(n + x.Off).Div((*x.Vals)[n])
	}
}
func (x *StmBse) InrRem(off unt.Unt) Stm {
	r := &StmInrRem{}
	r.Slf = r
	r.Stm = x.Slf
	r.Off = off
	if ana.Cfg.Trc.IsHstStm() {
		sys.Logf("%p hst.StmInrRem(%v)", r, r.Prm())
	}
	if x.Vals.Cnt() == 0 || x.Vals.Cnt() <= r.Off || r.Off == 0 {
		return r
	}
	if len(*x.Tmes) != len(*x.Vals) {
		err.Panicf("length unequal (Tmes:%v Vals:%v)", len(*x.Tmes), len(*x.Vals))
	}
	r.Tmes = x.Tmes.From(r.Off)
	r.Vals = flts.Make(x.Vals.Cnt() - r.Off)
	segBnds, acts := bnds.Segs(r.Vals.Cnt())
	for n, segBnd := range *segBnds {
		seg := &StmInrRemSeg{}
		seg.Bnd = segBnd
		seg.Vals = x.Vals
		seg.Out = r.Vals
		seg.Off = r.Off
		acts[n] = seg
	}
	sys.Run().Pll(acts...)
	return r
}
func (x *StmInrRem) Name() str.Str             { return str.Str("InrRem") }
func (x *StmInrRem) PrmWrt(b *strings.Builder) { x.Off.StrWrt(b) }
func (x *StmInrRem) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *StmInrRem) StrWrt(b *strings.Builder) {
	x.Stm.StrWrt(b)
	b.WriteString(".inrRem(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *StmInrRem) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *StmInrRemSeg) Act() {
	for n := x.Idx; n < x.Lim; n++ {
		(*x.Out)[n] = x.Vals.At(n + x.Off).Rem((*x.Vals)[n])
	}
}
func (x *StmBse) InrPow(off unt.Unt) Stm {
	r := &StmInrPow{}
	r.Slf = r
	r.Stm = x.Slf
	r.Off = off
	if ana.Cfg.Trc.IsHstStm() {
		sys.Logf("%p hst.StmInrPow(%v)", r, r.Prm())
	}
	if x.Vals.Cnt() == 0 || x.Vals.Cnt() <= r.Off || r.Off == 0 {
		return r
	}
	if len(*x.Tmes) != len(*x.Vals) {
		err.Panicf("length unequal (Tmes:%v Vals:%v)", len(*x.Tmes), len(*x.Vals))
	}
	r.Tmes = x.Tmes.From(r.Off)
	r.Vals = flts.Make(x.Vals.Cnt() - r.Off)
	segBnds, acts := bnds.Segs(r.Vals.Cnt())
	for n, segBnd := range *segBnds {
		seg := &StmInrPowSeg{}
		seg.Bnd = segBnd
		seg.Vals = x.Vals
		seg.Out = r.Vals
		seg.Off = r.Off
		acts[n] = seg
	}
	sys.Run().Pll(acts...)
	return r
}
func (x *StmInrPow) Name() str.Str             { return str.Str("InrPow") }
func (x *StmInrPow) PrmWrt(b *strings.Builder) { x.Off.StrWrt(b) }
func (x *StmInrPow) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *StmInrPow) StrWrt(b *strings.Builder) {
	x.Stm.StrWrt(b)
	b.WriteString(".inrPow(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *StmInrPow) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *StmInrPowSeg) Act() {
	for n := x.Idx; n < x.Lim; n++ {
		(*x.Out)[n] = x.Vals.At(n + x.Off).Pow((*x.Vals)[n])
	}
}
func (x *StmBse) InrMin(off unt.Unt) Stm {
	r := &StmInrMin{}
	r.Slf = r
	r.Stm = x.Slf
	r.Off = off
	if ana.Cfg.Trc.IsHstStm() {
		sys.Logf("%p hst.StmInrMin(%v)", r, r.Prm())
	}
	if x.Vals.Cnt() == 0 || x.Vals.Cnt() <= r.Off || r.Off == 0 {
		return r
	}
	if len(*x.Tmes) != len(*x.Vals) {
		err.Panicf("length unequal (Tmes:%v Vals:%v)", len(*x.Tmes), len(*x.Vals))
	}
	r.Tmes = x.Tmes.From(r.Off)
	r.Vals = flts.Make(x.Vals.Cnt() - r.Off)
	segBnds, acts := bnds.Segs(r.Vals.Cnt())
	for n, segBnd := range *segBnds {
		seg := &StmInrMinSeg{}
		seg.Bnd = segBnd
		seg.Vals = x.Vals
		seg.Out = r.Vals
		seg.Off = r.Off
		acts[n] = seg
	}
	sys.Run().Pll(acts...)
	return r
}
func (x *StmInrMin) Name() str.Str             { return str.Str("InrMin") }
func (x *StmInrMin) PrmWrt(b *strings.Builder) { x.Off.StrWrt(b) }
func (x *StmInrMin) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *StmInrMin) StrWrt(b *strings.Builder) {
	x.Stm.StrWrt(b)
	b.WriteString(".inrMin(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *StmInrMin) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *StmInrMinSeg) Act() {
	for n := x.Idx; n < x.Lim; n++ {
		(*x.Out)[n] = x.Vals.At(n + x.Off).Min((*x.Vals)[n])
	}
}
func (x *StmBse) InrMax(off unt.Unt) Stm {
	r := &StmInrMax{}
	r.Slf = r
	r.Stm = x.Slf
	r.Off = off
	if ana.Cfg.Trc.IsHstStm() {
		sys.Logf("%p hst.StmInrMax(%v)", r, r.Prm())
	}
	if x.Vals.Cnt() == 0 || x.Vals.Cnt() <= r.Off || r.Off == 0 {
		return r
	}
	if len(*x.Tmes) != len(*x.Vals) {
		err.Panicf("length unequal (Tmes:%v Vals:%v)", len(*x.Tmes), len(*x.Vals))
	}
	r.Tmes = x.Tmes.From(r.Off)
	r.Vals = flts.Make(x.Vals.Cnt() - r.Off)
	segBnds, acts := bnds.Segs(r.Vals.Cnt())
	for n, segBnd := range *segBnds {
		seg := &StmInrMaxSeg{}
		seg.Bnd = segBnd
		seg.Vals = x.Vals
		seg.Out = r.Vals
		seg.Off = r.Off
		acts[n] = seg
	}
	sys.Run().Pll(acts...)
	return r
}
func (x *StmInrMax) Name() str.Str             { return str.Str("InrMax") }
func (x *StmInrMax) PrmWrt(b *strings.Builder) { x.Off.StrWrt(b) }
func (x *StmInrMax) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *StmInrMax) StrWrt(b *strings.Builder) {
	x.Stm.StrWrt(b)
	b.WriteString(".inrMax(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *StmInrMax) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *StmInrMaxSeg) Act() {
	for n := x.Idx; n < x.Lim; n++ {
		(*x.Out)[n] = x.Vals.At(n + x.Off).Max((*x.Vals)[n])
	}
}
func (x *StmBse) InrSlp(off unt.Unt) Stm {
	r := &StmInr1Slp{}
	r.Slf = r
	r.Stm = x.Slf
	r.Off = off
	if ana.Cfg.Trc.IsHstStm() {
		sys.Logf("%p hst.StmInr1Slp(%v)", r, r.Prm())
	}
	if x.Vals.Cnt() == 0 || x.Vals.Cnt() <= r.Off || r.Off == 0 {
		return r
	}
	if len(*x.Tmes) != len(*x.Vals) {
		err.Panicf("length unequal (Tmes:%v Vals:%v)", len(*x.Tmes), len(*x.Vals))
	}
	r.Tmes = x.Tmes.From(r.Off)
	r.Vals = flts.Make(x.Vals.Cnt() - r.Off)
	segBnds, acts := bnds.Segs(r.Vals.Cnt())
	for n, segBnd := range *segBnds {
		seg := &StmInr1SlpSeg{}
		seg.Bnd = segBnd
		seg.Tmes = x.Tmes
		seg.Vals = x.Vals
		seg.Out = r.Vals
		seg.Off = r.Off
		acts[n] = seg
	}
	sys.Run().Pll(acts...)
	return r
}
func (x *StmInr1Slp) Name() str.Str             { return str.Str("InrSlp") }
func (x *StmInr1Slp) PrmWrt(b *strings.Builder) { x.Off.StrWrt(b) }
func (x *StmInr1Slp) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *StmInr1Slp) StrWrt(b *strings.Builder) {
	x.Stm.StrWrt(b)
	b.WriteString(".inrSlp(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *StmInr1Slp) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *StmInr1SlpSeg) Act() {
	for n := x.Idx; n < x.Lim; n++ { // expects 32-bit tme with 1s resolution
		(*x.Out)[n] = ((*x.Vals)[n+x.Off] - (*x.Vals)[n]) / flt.Flt((*x.Tmes)[n+x.Off]-(*x.Tmes)[n])
	}
}
func (x *StmBse) OtrAdd(off unt.Unt, a Stm) Stm {
	r := &StmOtrAdd{}
	r.Slf = r
	r.Stm = x.Slf
	r.Off = off
	r.A = a
	if ana.Cfg.Trc.IsHstStm() {
		sys.Logf("%p hst.StmOtrAdd(%v)", r, r.Prm())
	}
	aBse := r.A.Bse()
	if x.Vals.Cnt() == 0 || aBse.Vals.Cnt() == 0 || aBse.Vals.Cnt() <= r.Off {
		return r
	}
	if len(*x.Tmes) != len(*x.Vals) {
		err.Panicf("'x' length unequal (Tmes:%v Vals:%v)", len(*x.Tmes), len(*x.Vals))
	}
	if len(*aBse.Tmes) != len(*aBse.Vals) {
		err.Panicf("'a' length unequal (Tmes:%v Vals:%v)", len(*aBse.Tmes), len(*aBse.Vals))
	}
	xBnd, aBnd := AlignStmOtr(x.Slf, r.A, r.Off)
	if aBnd.Cnt() == 0 {
		return r
	}
	r.Tmes = aBse.Tmes.InBnd(aBnd).From(r.Off)
	r.Vals = flts.Make(aBnd.Cnt() - r.Off)
	segBnds, acts := bnds.Segs(r.Vals.Cnt())
	for n, segBnd := range *segBnds {
		seg := &StmOtrAddSeg{}
		seg.Bnd = segBnd
		seg.Vals = x.Vals.InBnd(xBnd)
		seg.Out = r.Vals
		seg.Off = r.Off
		seg.ValsA = aBse.Vals.InBnd(aBnd)
		acts[n] = seg
	}
	sys.Run().Pll(acts...)
	return r
}
func (x *StmOtrAdd) Name() str.Str { return str.Str("OtrAdd") }
func (x *StmOtrAdd) PrmWrt(b *strings.Builder) {
	x.Off.StrWrt(b)
	b.WriteRune(' ')
	x.A.StrWrt(b)
}
func (x *StmOtrAdd) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *StmOtrAdd) StrWrt(b *strings.Builder) {
	x.Stm.StrWrt(b)
	b.WriteString(".otrAdd(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *StmOtrAdd) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *StmOtrAddSeg) Act() {
	for n := x.Idx; n < x.Lim; n++ {
		(*x.Out)[n] = x.Vals.At(n).Add((*x.ValsA)[n+x.Off])
	}
}
func (x *StmBse) OtrSub(off unt.Unt, a Stm) Stm {
	r := &StmOtrSub{}
	r.Slf = r
	r.Stm = x.Slf
	r.Off = off
	r.A = a
	if ana.Cfg.Trc.IsHstStm() {
		sys.Logf("%p hst.StmOtrSub(%v)", r, r.Prm())
	}
	aBse := r.A.Bse()
	if x.Vals.Cnt() == 0 || aBse.Vals.Cnt() == 0 || aBse.Vals.Cnt() <= r.Off {
		return r
	}
	if len(*x.Tmes) != len(*x.Vals) {
		err.Panicf("'x' length unequal (Tmes:%v Vals:%v)", len(*x.Tmes), len(*x.Vals))
	}
	if len(*aBse.Tmes) != len(*aBse.Vals) {
		err.Panicf("'a' length unequal (Tmes:%v Vals:%v)", len(*aBse.Tmes), len(*aBse.Vals))
	}
	xBnd, aBnd := AlignStmOtr(x.Slf, r.A, r.Off)
	if aBnd.Cnt() == 0 {
		return r
	}
	r.Tmes = aBse.Tmes.InBnd(aBnd).From(r.Off)
	r.Vals = flts.Make(aBnd.Cnt() - r.Off)
	segBnds, acts := bnds.Segs(r.Vals.Cnt())
	for n, segBnd := range *segBnds {
		seg := &StmOtrSubSeg{}
		seg.Bnd = segBnd
		seg.Vals = x.Vals.InBnd(xBnd)
		seg.Out = r.Vals
		seg.Off = r.Off
		seg.ValsA = aBse.Vals.InBnd(aBnd)
		acts[n] = seg
	}
	sys.Run().Pll(acts...)
	return r
}
func (x *StmOtrSub) Name() str.Str { return str.Str("OtrSub") }
func (x *StmOtrSub) PrmWrt(b *strings.Builder) {
	x.Off.StrWrt(b)
	b.WriteRune(' ')
	x.A.StrWrt(b)
}
func (x *StmOtrSub) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *StmOtrSub) StrWrt(b *strings.Builder) {
	x.Stm.StrWrt(b)
	b.WriteString(".otrSub(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *StmOtrSub) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *StmOtrSubSeg) Act() {
	for n := x.Idx; n < x.Lim; n++ {
		(*x.Out)[n] = x.Vals.At(n).Sub((*x.ValsA)[n+x.Off])
	}
}
func (x *StmBse) OtrMul(off unt.Unt, a Stm) Stm {
	r := &StmOtrMul{}
	r.Slf = r
	r.Stm = x.Slf
	r.Off = off
	r.A = a
	if ana.Cfg.Trc.IsHstStm() {
		sys.Logf("%p hst.StmOtrMul(%v)", r, r.Prm())
	}
	aBse := r.A.Bse()
	if x.Vals.Cnt() == 0 || aBse.Vals.Cnt() == 0 || aBse.Vals.Cnt() <= r.Off {
		return r
	}
	if len(*x.Tmes) != len(*x.Vals) {
		err.Panicf("'x' length unequal (Tmes:%v Vals:%v)", len(*x.Tmes), len(*x.Vals))
	}
	if len(*aBse.Tmes) != len(*aBse.Vals) {
		err.Panicf("'a' length unequal (Tmes:%v Vals:%v)", len(*aBse.Tmes), len(*aBse.Vals))
	}
	xBnd, aBnd := AlignStmOtr(x.Slf, r.A, r.Off)
	if aBnd.Cnt() == 0 {
		return r
	}
	r.Tmes = aBse.Tmes.InBnd(aBnd).From(r.Off)
	r.Vals = flts.Make(aBnd.Cnt() - r.Off)
	segBnds, acts := bnds.Segs(r.Vals.Cnt())
	for n, segBnd := range *segBnds {
		seg := &StmOtrMulSeg{}
		seg.Bnd = segBnd
		seg.Vals = x.Vals.InBnd(xBnd)
		seg.Out = r.Vals
		seg.Off = r.Off
		seg.ValsA = aBse.Vals.InBnd(aBnd)
		acts[n] = seg
	}
	sys.Run().Pll(acts...)
	return r
}
func (x *StmOtrMul) Name() str.Str { return str.Str("OtrMul") }
func (x *StmOtrMul) PrmWrt(b *strings.Builder) {
	x.Off.StrWrt(b)
	b.WriteRune(' ')
	x.A.StrWrt(b)
}
func (x *StmOtrMul) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *StmOtrMul) StrWrt(b *strings.Builder) {
	x.Stm.StrWrt(b)
	b.WriteString(".otrMul(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *StmOtrMul) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *StmOtrMulSeg) Act() {
	for n := x.Idx; n < x.Lim; n++ {
		(*x.Out)[n] = x.Vals.At(n).Mul((*x.ValsA)[n+x.Off])
	}
}
func (x *StmBse) OtrDiv(off unt.Unt, a Stm) Stm {
	r := &StmOtrDiv{}
	r.Slf = r
	r.Stm = x.Slf
	r.Off = off
	r.A = a
	if ana.Cfg.Trc.IsHstStm() {
		sys.Logf("%p hst.StmOtrDiv(%v)", r, r.Prm())
	}
	aBse := r.A.Bse()
	if x.Vals.Cnt() == 0 || aBse.Vals.Cnt() == 0 || aBse.Vals.Cnt() <= r.Off {
		return r
	}
	if len(*x.Tmes) != len(*x.Vals) {
		err.Panicf("'x' length unequal (Tmes:%v Vals:%v)", len(*x.Tmes), len(*x.Vals))
	}
	if len(*aBse.Tmes) != len(*aBse.Vals) {
		err.Panicf("'a' length unequal (Tmes:%v Vals:%v)", len(*aBse.Tmes), len(*aBse.Vals))
	}
	xBnd, aBnd := AlignStmOtr(x.Slf, r.A, r.Off)
	if aBnd.Cnt() == 0 {
		return r
	}
	r.Tmes = aBse.Tmes.InBnd(aBnd).From(r.Off)
	r.Vals = flts.Make(aBnd.Cnt() - r.Off)
	segBnds, acts := bnds.Segs(r.Vals.Cnt())
	for n, segBnd := range *segBnds {
		seg := &StmOtrDivSeg{}
		seg.Bnd = segBnd
		seg.Vals = x.Vals.InBnd(xBnd)
		seg.Out = r.Vals
		seg.Off = r.Off
		seg.ValsA = aBse.Vals.InBnd(aBnd)
		acts[n] = seg
	}
	sys.Run().Pll(acts...)
	return r
}
func (x *StmOtrDiv) Name() str.Str { return str.Str("OtrDiv") }
func (x *StmOtrDiv) PrmWrt(b *strings.Builder) {
	x.Off.StrWrt(b)
	b.WriteRune(' ')
	x.A.StrWrt(b)
}
func (x *StmOtrDiv) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *StmOtrDiv) StrWrt(b *strings.Builder) {
	x.Stm.StrWrt(b)
	b.WriteString(".otrDiv(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *StmOtrDiv) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *StmOtrDivSeg) Act() {
	for n := x.Idx; n < x.Lim; n++ {
		(*x.Out)[n] = x.Vals.At(n).Div((*x.ValsA)[n+x.Off])
	}
}
func (x *StmBse) OtrRem(off unt.Unt, a Stm) Stm {
	r := &StmOtrRem{}
	r.Slf = r
	r.Stm = x.Slf
	r.Off = off
	r.A = a
	if ana.Cfg.Trc.IsHstStm() {
		sys.Logf("%p hst.StmOtrRem(%v)", r, r.Prm())
	}
	aBse := r.A.Bse()
	if x.Vals.Cnt() == 0 || aBse.Vals.Cnt() == 0 || aBse.Vals.Cnt() <= r.Off {
		return r
	}
	if len(*x.Tmes) != len(*x.Vals) {
		err.Panicf("'x' length unequal (Tmes:%v Vals:%v)", len(*x.Tmes), len(*x.Vals))
	}
	if len(*aBse.Tmes) != len(*aBse.Vals) {
		err.Panicf("'a' length unequal (Tmes:%v Vals:%v)", len(*aBse.Tmes), len(*aBse.Vals))
	}
	xBnd, aBnd := AlignStmOtr(x.Slf, r.A, r.Off)
	if aBnd.Cnt() == 0 {
		return r
	}
	r.Tmes = aBse.Tmes.InBnd(aBnd).From(r.Off)
	r.Vals = flts.Make(aBnd.Cnt() - r.Off)
	segBnds, acts := bnds.Segs(r.Vals.Cnt())
	for n, segBnd := range *segBnds {
		seg := &StmOtrRemSeg{}
		seg.Bnd = segBnd
		seg.Vals = x.Vals.InBnd(xBnd)
		seg.Out = r.Vals
		seg.Off = r.Off
		seg.ValsA = aBse.Vals.InBnd(aBnd)
		acts[n] = seg
	}
	sys.Run().Pll(acts...)
	return r
}
func (x *StmOtrRem) Name() str.Str { return str.Str("OtrRem") }
func (x *StmOtrRem) PrmWrt(b *strings.Builder) {
	x.Off.StrWrt(b)
	b.WriteRune(' ')
	x.A.StrWrt(b)
}
func (x *StmOtrRem) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *StmOtrRem) StrWrt(b *strings.Builder) {
	x.Stm.StrWrt(b)
	b.WriteString(".otrRem(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *StmOtrRem) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *StmOtrRemSeg) Act() {
	for n := x.Idx; n < x.Lim; n++ {
		(*x.Out)[n] = x.Vals.At(n).Rem((*x.ValsA)[n+x.Off])
	}
}
func (x *StmBse) OtrPow(off unt.Unt, a Stm) Stm {
	r := &StmOtrPow{}
	r.Slf = r
	r.Stm = x.Slf
	r.Off = off
	r.A = a
	if ana.Cfg.Trc.IsHstStm() {
		sys.Logf("%p hst.StmOtrPow(%v)", r, r.Prm())
	}
	aBse := r.A.Bse()
	if x.Vals.Cnt() == 0 || aBse.Vals.Cnt() == 0 || aBse.Vals.Cnt() <= r.Off {
		return r
	}
	if len(*x.Tmes) != len(*x.Vals) {
		err.Panicf("'x' length unequal (Tmes:%v Vals:%v)", len(*x.Tmes), len(*x.Vals))
	}
	if len(*aBse.Tmes) != len(*aBse.Vals) {
		err.Panicf("'a' length unequal (Tmes:%v Vals:%v)", len(*aBse.Tmes), len(*aBse.Vals))
	}
	xBnd, aBnd := AlignStmOtr(x.Slf, r.A, r.Off)
	if aBnd.Cnt() == 0 {
		return r
	}
	r.Tmes = aBse.Tmes.InBnd(aBnd).From(r.Off)
	r.Vals = flts.Make(aBnd.Cnt() - r.Off)
	segBnds, acts := bnds.Segs(r.Vals.Cnt())
	for n, segBnd := range *segBnds {
		seg := &StmOtrPowSeg{}
		seg.Bnd = segBnd
		seg.Vals = x.Vals.InBnd(xBnd)
		seg.Out = r.Vals
		seg.Off = r.Off
		seg.ValsA = aBse.Vals.InBnd(aBnd)
		acts[n] = seg
	}
	sys.Run().Pll(acts...)
	return r
}
func (x *StmOtrPow) Name() str.Str { return str.Str("OtrPow") }
func (x *StmOtrPow) PrmWrt(b *strings.Builder) {
	x.Off.StrWrt(b)
	b.WriteRune(' ')
	x.A.StrWrt(b)
}
func (x *StmOtrPow) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *StmOtrPow) StrWrt(b *strings.Builder) {
	x.Stm.StrWrt(b)
	b.WriteString(".otrPow(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *StmOtrPow) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *StmOtrPowSeg) Act() {
	for n := x.Idx; n < x.Lim; n++ {
		(*x.Out)[n] = x.Vals.At(n).Pow((*x.ValsA)[n+x.Off])
	}
}
func (x *StmBse) OtrMin(off unt.Unt, a Stm) Stm {
	r := &StmOtrMin{}
	r.Slf = r
	r.Stm = x.Slf
	r.Off = off
	r.A = a
	if ana.Cfg.Trc.IsHstStm() {
		sys.Logf("%p hst.StmOtrMin(%v)", r, r.Prm())
	}
	aBse := r.A.Bse()
	if x.Vals.Cnt() == 0 || aBse.Vals.Cnt() == 0 || aBse.Vals.Cnt() <= r.Off {
		return r
	}
	if len(*x.Tmes) != len(*x.Vals) {
		err.Panicf("'x' length unequal (Tmes:%v Vals:%v)", len(*x.Tmes), len(*x.Vals))
	}
	if len(*aBse.Tmes) != len(*aBse.Vals) {
		err.Panicf("'a' length unequal (Tmes:%v Vals:%v)", len(*aBse.Tmes), len(*aBse.Vals))
	}
	xBnd, aBnd := AlignStmOtr(x.Slf, r.A, r.Off)
	if aBnd.Cnt() == 0 {
		return r
	}
	r.Tmes = aBse.Tmes.InBnd(aBnd).From(r.Off)
	r.Vals = flts.Make(aBnd.Cnt() - r.Off)
	segBnds, acts := bnds.Segs(r.Vals.Cnt())
	for n, segBnd := range *segBnds {
		seg := &StmOtrMinSeg{}
		seg.Bnd = segBnd
		seg.Vals = x.Vals.InBnd(xBnd)
		seg.Out = r.Vals
		seg.Off = r.Off
		seg.ValsA = aBse.Vals.InBnd(aBnd)
		acts[n] = seg
	}
	sys.Run().Pll(acts...)
	return r
}
func (x *StmOtrMin) Name() str.Str { return str.Str("OtrMin") }
func (x *StmOtrMin) PrmWrt(b *strings.Builder) {
	x.Off.StrWrt(b)
	b.WriteRune(' ')
	x.A.StrWrt(b)
}
func (x *StmOtrMin) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *StmOtrMin) StrWrt(b *strings.Builder) {
	x.Stm.StrWrt(b)
	b.WriteString(".otrMin(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *StmOtrMin) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *StmOtrMinSeg) Act() {
	for n := x.Idx; n < x.Lim; n++ {
		(*x.Out)[n] = x.Vals.At(n).Min((*x.ValsA)[n+x.Off])
	}
}
func (x *StmBse) OtrMax(off unt.Unt, a Stm) Stm {
	r := &StmOtrMax{}
	r.Slf = r
	r.Stm = x.Slf
	r.Off = off
	r.A = a
	if ana.Cfg.Trc.IsHstStm() {
		sys.Logf("%p hst.StmOtrMax(%v)", r, r.Prm())
	}
	aBse := r.A.Bse()
	if x.Vals.Cnt() == 0 || aBse.Vals.Cnt() == 0 || aBse.Vals.Cnt() <= r.Off {
		return r
	}
	if len(*x.Tmes) != len(*x.Vals) {
		err.Panicf("'x' length unequal (Tmes:%v Vals:%v)", len(*x.Tmes), len(*x.Vals))
	}
	if len(*aBse.Tmes) != len(*aBse.Vals) {
		err.Panicf("'a' length unequal (Tmes:%v Vals:%v)", len(*aBse.Tmes), len(*aBse.Vals))
	}
	xBnd, aBnd := AlignStmOtr(x.Slf, r.A, r.Off)
	if aBnd.Cnt() == 0 {
		return r
	}
	r.Tmes = aBse.Tmes.InBnd(aBnd).From(r.Off)
	r.Vals = flts.Make(aBnd.Cnt() - r.Off)
	segBnds, acts := bnds.Segs(r.Vals.Cnt())
	for n, segBnd := range *segBnds {
		seg := &StmOtrMaxSeg{}
		seg.Bnd = segBnd
		seg.Vals = x.Vals.InBnd(xBnd)
		seg.Out = r.Vals
		seg.Off = r.Off
		seg.ValsA = aBse.Vals.InBnd(aBnd)
		acts[n] = seg
	}
	sys.Run().Pll(acts...)
	return r
}
func (x *StmOtrMax) Name() str.Str { return str.Str("OtrMax") }
func (x *StmOtrMax) PrmWrt(b *strings.Builder) {
	x.Off.StrWrt(b)
	b.WriteRune(' ')
	x.A.StrWrt(b)
}
func (x *StmOtrMax) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *StmOtrMax) StrWrt(b *strings.Builder) {
	x.Stm.StrWrt(b)
	b.WriteString(".otrMax(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *StmOtrMax) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *StmOtrMaxSeg) Act() {
	for n := x.Idx; n < x.Lim; n++ {
		(*x.Out)[n] = x.Vals.At(n).Max((*x.ValsA)[n+x.Off])
	}
}
func (x *StmBse) SclEql(scl flt.Flt) Cnd {
	r := &CndSclEql{}
	r.Slf = r
	r.Stm = x.Slf
	r.Scl = scl
	if ana.Cfg.Trc.IsHstCnd() {
		sys.Logf("%p hst.CndSclEql(%v)", r, r.Prm())
	}
	if x.Vals.Cnt() == 0 {
		return r
	}
	if len(*x.Tmes) != len(*x.Vals) {
		err.Panicf("length unequal (Tmes:%v Vals:%v)", len(*x.Tmes), len(*x.Vals))
	}
	r.Tmes = tmes.New()
	segBnds, acts := bnds.Segs(x.Vals.Cnt())
	for n, segBnd := range *segBnds {
		seg := &CndSclEqlSeg{}
		seg.Bnd = segBnd
		seg.Tmes = x.Tmes
		seg.Vals = x.Vals
		seg.Scl = r.Scl
		seg.Out = tmes.New()
		acts[n] = seg
	}
	sys.Run().Pll(acts...)
	for _, act := range acts {
		r.Tmes.Mrg(act.(*CndSclEqlSeg).Out)
	}
	return r
}
func (x *StmBse) SclNeq(scl flt.Flt) Cnd {
	r := &CndSclNeq{}
	r.Slf = r
	r.Stm = x.Slf
	r.Scl = scl
	if ana.Cfg.Trc.IsHstCnd() {
		sys.Logf("%p hst.CndSclNeq(%v)", r, r.Prm())
	}
	if x.Vals.Cnt() == 0 {
		return r
	}
	if len(*x.Tmes) != len(*x.Vals) {
		err.Panicf("length unequal (Tmes:%v Vals:%v)", len(*x.Tmes), len(*x.Vals))
	}
	r.Tmes = tmes.New()
	segBnds, acts := bnds.Segs(x.Vals.Cnt())
	for n, segBnd := range *segBnds {
		seg := &CndSclNeqSeg{}
		seg.Bnd = segBnd
		seg.Tmes = x.Tmes
		seg.Vals = x.Vals
		seg.Scl = r.Scl
		seg.Out = tmes.New()
		acts[n] = seg
	}
	sys.Run().Pll(acts...)
	for _, act := range acts {
		r.Tmes.Mrg(act.(*CndSclNeqSeg).Out)
	}
	return r
}
func (x *StmBse) SclLss(scl flt.Flt) Cnd {
	r := &CndSclLss{}
	r.Slf = r
	r.Stm = x.Slf
	r.Scl = scl
	if ana.Cfg.Trc.IsHstCnd() {
		sys.Logf("%p hst.CndSclLss(%v)", r, r.Prm())
	}
	if x.Vals.Cnt() == 0 {
		return r
	}
	if len(*x.Tmes) != len(*x.Vals) {
		err.Panicf("length unequal (Tmes:%v Vals:%v)", len(*x.Tmes), len(*x.Vals))
	}
	r.Tmes = tmes.New()
	segBnds, acts := bnds.Segs(x.Vals.Cnt())
	for n, segBnd := range *segBnds {
		seg := &CndSclLssSeg{}
		seg.Bnd = segBnd
		seg.Tmes = x.Tmes
		seg.Vals = x.Vals
		seg.Scl = r.Scl
		seg.Out = tmes.New()
		acts[n] = seg
	}
	sys.Run().Pll(acts...)
	for _, act := range acts {
		r.Tmes.Mrg(act.(*CndSclLssSeg).Out)
	}
	return r
}
func (x *StmBse) SclGtr(scl flt.Flt) Cnd {
	r := &CndSclGtr{}
	r.Slf = r
	r.Stm = x.Slf
	r.Scl = scl
	if ana.Cfg.Trc.IsHstCnd() {
		sys.Logf("%p hst.CndSclGtr(%v)", r, r.Prm())
	}
	if x.Vals.Cnt() == 0 {
		return r
	}
	if len(*x.Tmes) != len(*x.Vals) {
		err.Panicf("length unequal (Tmes:%v Vals:%v)", len(*x.Tmes), len(*x.Vals))
	}
	r.Tmes = tmes.New()
	segBnds, acts := bnds.Segs(x.Vals.Cnt())
	for n, segBnd := range *segBnds {
		seg := &CndSclGtrSeg{}
		seg.Bnd = segBnd
		seg.Tmes = x.Tmes
		seg.Vals = x.Vals
		seg.Scl = r.Scl
		seg.Out = tmes.New()
		acts[n] = seg
	}
	sys.Run().Pll(acts...)
	for _, act := range acts {
		r.Tmes.Mrg(act.(*CndSclGtrSeg).Out)
	}
	return r
}
func (x *StmBse) SclLeq(scl flt.Flt) Cnd {
	r := &CndSclLeq{}
	r.Slf = r
	r.Stm = x.Slf
	r.Scl = scl
	if ana.Cfg.Trc.IsHstCnd() {
		sys.Logf("%p hst.CndSclLeq(%v)", r, r.Prm())
	}
	if x.Vals.Cnt() == 0 {
		return r
	}
	if len(*x.Tmes) != len(*x.Vals) {
		err.Panicf("length unequal (Tmes:%v Vals:%v)", len(*x.Tmes), len(*x.Vals))
	}
	r.Tmes = tmes.New()
	segBnds, acts := bnds.Segs(x.Vals.Cnt())
	for n, segBnd := range *segBnds {
		seg := &CndSclLeqSeg{}
		seg.Bnd = segBnd
		seg.Tmes = x.Tmes
		seg.Vals = x.Vals
		seg.Scl = r.Scl
		seg.Out = tmes.New()
		acts[n] = seg
	}
	sys.Run().Pll(acts...)
	for _, act := range acts {
		r.Tmes.Mrg(act.(*CndSclLeqSeg).Out)
	}
	return r
}
func (x *StmBse) SclGeq(scl flt.Flt) Cnd {
	r := &CndSclGeq{}
	r.Slf = r
	r.Stm = x.Slf
	r.Scl = scl
	if ana.Cfg.Trc.IsHstCnd() {
		sys.Logf("%p hst.CndSclGeq(%v)", r, r.Prm())
	}
	if x.Vals.Cnt() == 0 {
		return r
	}
	if len(*x.Tmes) != len(*x.Vals) {
		err.Panicf("length unequal (Tmes:%v Vals:%v)", len(*x.Tmes), len(*x.Vals))
	}
	r.Tmes = tmes.New()
	segBnds, acts := bnds.Segs(x.Vals.Cnt())
	for n, segBnd := range *segBnds {
		seg := &CndSclGeqSeg{}
		seg.Bnd = segBnd
		seg.Tmes = x.Tmes
		seg.Vals = x.Vals
		seg.Scl = r.Scl
		seg.Out = tmes.New()
		acts[n] = seg
	}
	sys.Run().Pll(acts...)
	for _, act := range acts {
		r.Tmes.Mrg(act.(*CndSclGeqSeg).Out)
	}
	return r
}
func (x *StmBse) InrEql(off unt.Unt) Cnd {
	r := &CndInrEql{}
	r.Slf = r
	r.Stm = x.Slf
	r.Off = off
	if ana.Cfg.Trc.IsHstCnd() {
		sys.Logf("%p hst.CndInrEql(%v)", r, r.Prm())
	}
	if len(*x.Vals) == 0 || x.Vals.Cnt() <= r.Off || r.Off == 0 {
		return r
	}
	if len(*x.Tmes) != len(*x.Vals) {
		err.Panicf("length unequal (Tmes:%v Vals:%v)", len(*x.Tmes), len(*x.Vals))
	}
	r.Tmes = tmes.New()
	segBnds, acts := bnds.Segs(x.Vals.Cnt() - r.Off)
	for n, segBnd := range *segBnds {
		seg := &CndInrEqlSeg{}
		seg.Bnd = segBnd
		seg.Tmes = x.Tmes
		seg.Vals = x.Vals
		seg.Off = r.Off
		seg.Out = tmes.New()
		acts[n] = seg
	}
	sys.Run().Pll(acts...)
	for _, act := range acts {
		r.Tmes.Mrg(act.(*CndInrEqlSeg).Out)
	}
	return r
}
func (x *StmBse) InrNeq(off unt.Unt) Cnd {
	r := &CndInrNeq{}
	r.Slf = r
	r.Stm = x.Slf
	r.Off = off
	if ana.Cfg.Trc.IsHstCnd() {
		sys.Logf("%p hst.CndInrNeq(%v)", r, r.Prm())
	}
	if len(*x.Vals) == 0 || x.Vals.Cnt() <= r.Off || r.Off == 0 {
		return r
	}
	if len(*x.Tmes) != len(*x.Vals) {
		err.Panicf("length unequal (Tmes:%v Vals:%v)", len(*x.Tmes), len(*x.Vals))
	}
	r.Tmes = tmes.New()
	segBnds, acts := bnds.Segs(x.Vals.Cnt() - r.Off)
	for n, segBnd := range *segBnds {
		seg := &CndInrNeqSeg{}
		seg.Bnd = segBnd
		seg.Tmes = x.Tmes
		seg.Vals = x.Vals
		seg.Off = r.Off
		seg.Out = tmes.New()
		acts[n] = seg
	}
	sys.Run().Pll(acts...)
	for _, act := range acts {
		r.Tmes.Mrg(act.(*CndInrNeqSeg).Out)
	}
	return r
}
func (x *StmBse) InrLss(off unt.Unt) Cnd {
	r := &CndInrLss{}
	r.Slf = r
	r.Stm = x.Slf
	r.Off = off
	if ana.Cfg.Trc.IsHstCnd() {
		sys.Logf("%p hst.CndInrLss(%v)", r, r.Prm())
	}
	if len(*x.Vals) == 0 || x.Vals.Cnt() <= r.Off || r.Off == 0 {
		return r
	}
	if len(*x.Tmes) != len(*x.Vals) {
		err.Panicf("length unequal (Tmes:%v Vals:%v)", len(*x.Tmes), len(*x.Vals))
	}
	r.Tmes = tmes.New()
	segBnds, acts := bnds.Segs(x.Vals.Cnt() - r.Off)
	for n, segBnd := range *segBnds {
		seg := &CndInrLssSeg{}
		seg.Bnd = segBnd
		seg.Tmes = x.Tmes
		seg.Vals = x.Vals
		seg.Off = r.Off
		seg.Out = tmes.New()
		acts[n] = seg
	}
	sys.Run().Pll(acts...)
	for _, act := range acts {
		r.Tmes.Mrg(act.(*CndInrLssSeg).Out)
	}
	return r
}
func (x *StmBse) InrGtr(off unt.Unt) Cnd {
	r := &CndInrGtr{}
	r.Slf = r
	r.Stm = x.Slf
	r.Off = off
	if ana.Cfg.Trc.IsHstCnd() {
		sys.Logf("%p hst.CndInrGtr(%v)", r, r.Prm())
	}
	if len(*x.Vals) == 0 || x.Vals.Cnt() <= r.Off || r.Off == 0 {
		return r
	}
	if len(*x.Tmes) != len(*x.Vals) {
		err.Panicf("length unequal (Tmes:%v Vals:%v)", len(*x.Tmes), len(*x.Vals))
	}
	r.Tmes = tmes.New()
	segBnds, acts := bnds.Segs(x.Vals.Cnt() - r.Off)
	for n, segBnd := range *segBnds {
		seg := &CndInrGtrSeg{}
		seg.Bnd = segBnd
		seg.Tmes = x.Tmes
		seg.Vals = x.Vals
		seg.Off = r.Off
		seg.Out = tmes.New()
		acts[n] = seg
	}
	sys.Run().Pll(acts...)
	for _, act := range acts {
		r.Tmes.Mrg(act.(*CndInrGtrSeg).Out)
	}
	return r
}
func (x *StmBse) InrLeq(off unt.Unt) Cnd {
	r := &CndInrLeq{}
	r.Slf = r
	r.Stm = x.Slf
	r.Off = off
	if ana.Cfg.Trc.IsHstCnd() {
		sys.Logf("%p hst.CndInrLeq(%v)", r, r.Prm())
	}
	if len(*x.Vals) == 0 || x.Vals.Cnt() <= r.Off || r.Off == 0 {
		return r
	}
	if len(*x.Tmes) != len(*x.Vals) {
		err.Panicf("length unequal (Tmes:%v Vals:%v)", len(*x.Tmes), len(*x.Vals))
	}
	r.Tmes = tmes.New()
	segBnds, acts := bnds.Segs(x.Vals.Cnt() - r.Off)
	for n, segBnd := range *segBnds {
		seg := &CndInrLeqSeg{}
		seg.Bnd = segBnd
		seg.Tmes = x.Tmes
		seg.Vals = x.Vals
		seg.Off = r.Off
		seg.Out = tmes.New()
		acts[n] = seg
	}
	sys.Run().Pll(acts...)
	for _, act := range acts {
		r.Tmes.Mrg(act.(*CndInrLeqSeg).Out)
	}
	return r
}
func (x *StmBse) InrGeq(off unt.Unt) Cnd {
	r := &CndInrGeq{}
	r.Slf = r
	r.Stm = x.Slf
	r.Off = off
	if ana.Cfg.Trc.IsHstCnd() {
		sys.Logf("%p hst.CndInrGeq(%v)", r, r.Prm())
	}
	if len(*x.Vals) == 0 || x.Vals.Cnt() <= r.Off || r.Off == 0 {
		return r
	}
	if len(*x.Tmes) != len(*x.Vals) {
		err.Panicf("length unequal (Tmes:%v Vals:%v)", len(*x.Tmes), len(*x.Vals))
	}
	r.Tmes = tmes.New()
	segBnds, acts := bnds.Segs(x.Vals.Cnt() - r.Off)
	for n, segBnd := range *segBnds {
		seg := &CndInrGeqSeg{}
		seg.Bnd = segBnd
		seg.Tmes = x.Tmes
		seg.Vals = x.Vals
		seg.Off = r.Off
		seg.Out = tmes.New()
		acts[n] = seg
	}
	sys.Run().Pll(acts...)
	for _, act := range acts {
		r.Tmes.Mrg(act.(*CndInrGeqSeg).Out)
	}
	return r
}
func (x *StmBse) OtrEql(off unt.Unt, a Stm) Cnd {
	r := &CndOtrEql{}
	r.Slf = r
	r.Stm = x.Slf
	r.Off = off
	r.A = a
	if ana.Cfg.Trc.IsHstCnd() {
		sys.Logf("%p hst.CndOtrEql(%v)", r, r.Prm())
	}
	xBse, aBse := x, r.A.Bse()
	if xBse.Vals.Cnt() == 0 || aBse.Vals.Cnt() == 0 || aBse.Vals.Cnt() <= r.Off {
		return r
	}
	if len(*xBse.Tmes) != len(*xBse.Vals) {
		err.Panicf("'x' length unequal (Tmes:%v Vals:%v)", len(*xBse.Tmes), len(*xBse.Vals))
	}
	if len(*aBse.Tmes) != len(*aBse.Vals) {
		err.Panicf("'a' length unequal (Tmes:%v Vals:%v)", len(*aBse.Tmes), len(*aBse.Vals))
	}
	xBnd, aBnd := AlignStmOtr(x.Slf, r.A, r.Off)
	if aBnd.Cnt() == 0 {
		return r
	}
	r.Tmes = tmes.New()
	segBnds, acts := bnds.Segs(aBnd.Cnt() - r.Off)
	for n, segBnd := range *segBnds {
		seg := &CndOtrEqlSeg{}
		seg.Bnd = segBnd
		seg.Vals = xBse.Vals.InBnd(xBnd)
		seg.Off = r.Off
		seg.ValsA = aBse.Vals.InBnd(aBnd)
		seg.Tmes = aBse.Tmes.InBnd(aBnd)
		seg.Out = tmes.New()
		acts[n] = seg
	}
	sys.Run().Pll(acts...)
	for _, act := range acts {
		r.Tmes.Mrg(act.(*CndOtrEqlSeg).Out)
	}
	return r
}
func (x *StmBse) OtrNeq(off unt.Unt, a Stm) Cnd {
	r := &CndOtrNeq{}
	r.Slf = r
	r.Stm = x.Slf
	r.Off = off
	r.A = a
	if ana.Cfg.Trc.IsHstCnd() {
		sys.Logf("%p hst.CndOtrNeq(%v)", r, r.Prm())
	}
	xBse, aBse := x, r.A.Bse()
	if xBse.Vals.Cnt() == 0 || aBse.Vals.Cnt() == 0 || aBse.Vals.Cnt() <= r.Off {
		return r
	}
	if len(*xBse.Tmes) != len(*xBse.Vals) {
		err.Panicf("'x' length unequal (Tmes:%v Vals:%v)", len(*xBse.Tmes), len(*xBse.Vals))
	}
	if len(*aBse.Tmes) != len(*aBse.Vals) {
		err.Panicf("'a' length unequal (Tmes:%v Vals:%v)", len(*aBse.Tmes), len(*aBse.Vals))
	}
	xBnd, aBnd := AlignStmOtr(x.Slf, r.A, r.Off)
	if aBnd.Cnt() == 0 {
		return r
	}
	r.Tmes = tmes.New()
	segBnds, acts := bnds.Segs(aBnd.Cnt() - r.Off)
	for n, segBnd := range *segBnds {
		seg := &CndOtrNeqSeg{}
		seg.Bnd = segBnd
		seg.Vals = xBse.Vals.InBnd(xBnd)
		seg.Off = r.Off
		seg.ValsA = aBse.Vals.InBnd(aBnd)
		seg.Tmes = aBse.Tmes.InBnd(aBnd)
		seg.Out = tmes.New()
		acts[n] = seg
	}
	sys.Run().Pll(acts...)
	for _, act := range acts {
		r.Tmes.Mrg(act.(*CndOtrNeqSeg).Out)
	}
	return r
}
func (x *StmBse) OtrLss(off unt.Unt, a Stm) Cnd {
	r := &CndOtrLss{}
	r.Slf = r
	r.Stm = x.Slf
	r.Off = off
	r.A = a
	if ana.Cfg.Trc.IsHstCnd() {
		sys.Logf("%p hst.CndOtrLss(%v)", r, r.Prm())
	}
	xBse, aBse := x, r.A.Bse()
	if xBse.Vals.Cnt() == 0 || aBse.Vals.Cnt() == 0 || aBse.Vals.Cnt() <= r.Off {
		return r
	}
	if len(*xBse.Tmes) != len(*xBse.Vals) {
		err.Panicf("'x' length unequal (Tmes:%v Vals:%v)", len(*xBse.Tmes), len(*xBse.Vals))
	}
	if len(*aBse.Tmes) != len(*aBse.Vals) {
		err.Panicf("'a' length unequal (Tmes:%v Vals:%v)", len(*aBse.Tmes), len(*aBse.Vals))
	}
	xBnd, aBnd := AlignStmOtr(x.Slf, r.A, r.Off)
	if aBnd.Cnt() == 0 {
		return r
	}
	r.Tmes = tmes.New()
	segBnds, acts := bnds.Segs(aBnd.Cnt() - r.Off)
	for n, segBnd := range *segBnds {
		seg := &CndOtrLssSeg{}
		seg.Bnd = segBnd
		seg.Vals = xBse.Vals.InBnd(xBnd)
		seg.Off = r.Off
		seg.ValsA = aBse.Vals.InBnd(aBnd)
		seg.Tmes = aBse.Tmes.InBnd(aBnd)
		seg.Out = tmes.New()
		acts[n] = seg
	}
	sys.Run().Pll(acts...)
	for _, act := range acts {
		r.Tmes.Mrg(act.(*CndOtrLssSeg).Out)
	}
	return r
}
func (x *StmBse) OtrGtr(off unt.Unt, a Stm) Cnd {
	r := &CndOtrGtr{}
	r.Slf = r
	r.Stm = x.Slf
	r.Off = off
	r.A = a
	if ana.Cfg.Trc.IsHstCnd() {
		sys.Logf("%p hst.CndOtrGtr(%v)", r, r.Prm())
	}
	xBse, aBse := x, r.A.Bse()
	if xBse.Vals.Cnt() == 0 || aBse.Vals.Cnt() == 0 || aBse.Vals.Cnt() <= r.Off {
		return r
	}
	if len(*xBse.Tmes) != len(*xBse.Vals) {
		err.Panicf("'x' length unequal (Tmes:%v Vals:%v)", len(*xBse.Tmes), len(*xBse.Vals))
	}
	if len(*aBse.Tmes) != len(*aBse.Vals) {
		err.Panicf("'a' length unequal (Tmes:%v Vals:%v)", len(*aBse.Tmes), len(*aBse.Vals))
	}
	xBnd, aBnd := AlignStmOtr(x.Slf, r.A, r.Off)
	if aBnd.Cnt() == 0 {
		return r
	}
	r.Tmes = tmes.New()
	segBnds, acts := bnds.Segs(aBnd.Cnt() - r.Off)
	for n, segBnd := range *segBnds {
		seg := &CndOtrGtrSeg{}
		seg.Bnd = segBnd
		seg.Vals = xBse.Vals.InBnd(xBnd)
		seg.Off = r.Off
		seg.ValsA = aBse.Vals.InBnd(aBnd)
		seg.Tmes = aBse.Tmes.InBnd(aBnd)
		seg.Out = tmes.New()
		acts[n] = seg
	}
	sys.Run().Pll(acts...)
	for _, act := range acts {
		r.Tmes.Mrg(act.(*CndOtrGtrSeg).Out)
	}
	return r
}
func (x *StmBse) OtrLeq(off unt.Unt, a Stm) Cnd {
	r := &CndOtrLeq{}
	r.Slf = r
	r.Stm = x.Slf
	r.Off = off
	r.A = a
	if ana.Cfg.Trc.IsHstCnd() {
		sys.Logf("%p hst.CndOtrLeq(%v)", r, r.Prm())
	}
	xBse, aBse := x, r.A.Bse()
	if xBse.Vals.Cnt() == 0 || aBse.Vals.Cnt() == 0 || aBse.Vals.Cnt() <= r.Off {
		return r
	}
	if len(*xBse.Tmes) != len(*xBse.Vals) {
		err.Panicf("'x' length unequal (Tmes:%v Vals:%v)", len(*xBse.Tmes), len(*xBse.Vals))
	}
	if len(*aBse.Tmes) != len(*aBse.Vals) {
		err.Panicf("'a' length unequal (Tmes:%v Vals:%v)", len(*aBse.Tmes), len(*aBse.Vals))
	}
	xBnd, aBnd := AlignStmOtr(x.Slf, r.A, r.Off)
	if aBnd.Cnt() == 0 {
		return r
	}
	r.Tmes = tmes.New()
	segBnds, acts := bnds.Segs(aBnd.Cnt() - r.Off)
	for n, segBnd := range *segBnds {
		seg := &CndOtrLeqSeg{}
		seg.Bnd = segBnd
		seg.Vals = xBse.Vals.InBnd(xBnd)
		seg.Off = r.Off
		seg.ValsA = aBse.Vals.InBnd(aBnd)
		seg.Tmes = aBse.Tmes.InBnd(aBnd)
		seg.Out = tmes.New()
		acts[n] = seg
	}
	sys.Run().Pll(acts...)
	for _, act := range acts {
		r.Tmes.Mrg(act.(*CndOtrLeqSeg).Out)
	}
	return r
}
func (x *StmBse) OtrGeq(off unt.Unt, a Stm) Cnd {
	r := &CndOtrGeq{}
	r.Slf = r
	r.Stm = x.Slf
	r.Off = off
	r.A = a
	if ana.Cfg.Trc.IsHstCnd() {
		sys.Logf("%p hst.CndOtrGeq(%v)", r, r.Prm())
	}
	xBse, aBse := x, r.A.Bse()
	if xBse.Vals.Cnt() == 0 || aBse.Vals.Cnt() == 0 || aBse.Vals.Cnt() <= r.Off {
		return r
	}
	if len(*xBse.Tmes) != len(*xBse.Vals) {
		err.Panicf("'x' length unequal (Tmes:%v Vals:%v)", len(*xBse.Tmes), len(*xBse.Vals))
	}
	if len(*aBse.Tmes) != len(*aBse.Vals) {
		err.Panicf("'a' length unequal (Tmes:%v Vals:%v)", len(*aBse.Tmes), len(*aBse.Vals))
	}
	xBnd, aBnd := AlignStmOtr(x.Slf, r.A, r.Off)
	if aBnd.Cnt() == 0 {
		return r
	}
	r.Tmes = tmes.New()
	segBnds, acts := bnds.Segs(aBnd.Cnt() - r.Off)
	for n, segBnd := range *segBnds {
		seg := &CndOtrGeqSeg{}
		seg.Bnd = segBnd
		seg.Vals = xBse.Vals.InBnd(xBnd)
		seg.Off = r.Off
		seg.ValsA = aBse.Vals.InBnd(aBnd)
		seg.Tmes = aBse.Tmes.InBnd(aBnd)
		seg.Out = tmes.New()
		acts[n] = seg
	}
	sys.Run().Pll(acts...)
	for _, act := range acts {
		r.Tmes.Mrg(act.(*CndOtrGeqSeg).Out)
	}
	return r
}
