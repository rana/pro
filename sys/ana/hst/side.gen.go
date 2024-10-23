package hst

import (
	"strings"
	"sys"
	"sys/ana"
	"sys/bsc/bnd"
	"sys/bsc/bnds"
	"sys/bsc/flt"
	"sys/bsc/flts"
	"sys/bsc/str"
	"sys/bsc/unt"
	"sys/err"
)

type (
	Side interface {
		Name() str.Str
		PrmWrt(b *strings.Builder)
		Prm() string
		StrWrt(b *strings.Builder)
		String() string
		Bse() *SideBse
		Fst() Stm
		Lst() Stm
		Sum() Stm
		Prd() Stm
		Min() Stm
		Max() Stm
		Mid() Stm
		Mdn() Stm
		Sma() Stm
		Gma() Stm
		Wma() Stm
		Rsi() Stm
		Wrsi() Stm
		Alma() Stm
		Vrnc() Stm
		Std() Stm
		RngFul() Stm
		RngLst() Stm
		ProLst() Stm
		ProSma() Stm
		ProAlma() Stm
		Sar(afInc, afMax flt.Flt) Stm
		Ema() Stm
	}
	SideBse struct {
		Slf     Side
		Inrvl   Inrvl
		Vals    *flts.Flts
		ValBnds *bnds.Bnds
	}
	SideSeg struct {
		bnd.Bnd
		Stm     *ana.Stm
		TmeBnds *bnds.Bnds
		Out     *bnds.Bnds
	}
	SideScp struct {
		Idx uint32
		Arr []Side
	}
	SideBid struct {
		SideBse
	}
	SideBidSeg struct {
		SideSeg
	}
	SideAsk struct {
		SideBse
	}
	SideAskSeg struct {
		SideSeg
	}
)

func (x *SideBid) Name() str.Str             { return str.Str("Bid") }
func (x *SideBid) PrmWrt(b *strings.Builder) {}
func (x *SideBid) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *SideBid) StrWrt(b *strings.Builder) {
	x.Inrvl.StrWrt(b)
	b.WriteString(".bid(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *SideBid) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *SideAsk) Name() str.Str             { return str.Str("Ask") }
func (x *SideAsk) PrmWrt(b *strings.Builder) {}
func (x *SideAsk) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *SideAsk) StrWrt(b *strings.Builder) {
	x.Inrvl.StrWrt(b)
	b.WriteString(".ask(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *SideAsk) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *SideBse) Bse() *SideBse { return x }
func (x *SideBse) Fst() Stm {
	r := &StmRteFst{}
	r.Slf = r
	r.Side = x.Slf
	if ana.Cfg.Trc.IsHstStm() {
		sys.Logf("%p hst.StmRteFst(%v)", r, r.Prm())
	}
	if x.ValBnds.Cnt() == 0 {
		return r
	}
	inrvl := x.Inrvl.Bse()
	if inrvl.TmeBnds.Cnt() != x.ValBnds.Cnt() {
		err.Panicf("length unequal (TmeBnds:%v ValBnds:%v)", inrvl.TmeBnds.Cnt(), x.ValBnds.Cnt())
	}
	r.Tmes = inrvl.Tmes
	r.Vals = flts.Make(inrvl.TmeBnds.Cnt())
	segBnds, acts := bnds.Segs(inrvl.TmeBnds.Cnt())
	for n, segBnd := range *segBnds {
		seg := &StmRteFstSeg{}
		seg.Bnd = segBnd
		seg.Vals = x.Vals
		seg.ValBnds = x.ValBnds
		seg.Out = r.Vals
		acts[n] = seg
	}
	sys.Run().Pll(acts...)
	return r
}
func (x *SideBse) Lst() Stm {
	r := &StmRteLst{}
	r.Slf = r
	r.Side = x.Slf
	if ana.Cfg.Trc.IsHstStm() {
		sys.Logf("%p hst.StmRteLst(%v)", r, r.Prm())
	}
	if x.ValBnds.Cnt() == 0 {
		return r
	}
	inrvl := x.Inrvl.Bse()
	if inrvl.TmeBnds.Cnt() != x.ValBnds.Cnt() {
		err.Panicf("length unequal (TmeBnds:%v ValBnds:%v)", inrvl.TmeBnds.Cnt(), x.ValBnds.Cnt())
	}
	r.Tmes = inrvl.Tmes
	r.Vals = flts.Make(inrvl.TmeBnds.Cnt())
	segBnds, acts := bnds.Segs(inrvl.TmeBnds.Cnt())
	for n, segBnd := range *segBnds {
		seg := &StmRteLstSeg{}
		seg.Bnd = segBnd
		seg.Vals = x.Vals
		seg.ValBnds = x.ValBnds
		seg.Out = r.Vals
		acts[n] = seg
	}
	sys.Run().Pll(acts...)
	return r
}
func (x *SideBse) Sum() Stm {
	r := &StmRteSum{}
	r.Slf = r
	r.Side = x.Slf
	if ana.Cfg.Trc.IsHstStm() {
		sys.Logf("%p hst.StmRteSum(%v)", r, r.Prm())
	}
	if x.ValBnds.Cnt() == 0 {
		return r
	}
	inrvl := x.Inrvl.Bse()
	if inrvl.TmeBnds.Cnt() != x.ValBnds.Cnt() {
		err.Panicf("length unequal (TmeBnds:%v ValBnds:%v)", inrvl.TmeBnds.Cnt(), x.ValBnds.Cnt())
	}
	r.Tmes = inrvl.Tmes
	r.Vals = flts.Make(inrvl.TmeBnds.Cnt())
	segBnds, acts := bnds.Segs(inrvl.TmeBnds.Cnt())
	for n, segBnd := range *segBnds {
		seg := &StmRteSumSeg{}
		seg.Bnd = segBnd
		seg.Vals = x.Vals
		seg.ValBnds = x.ValBnds
		seg.Out = r.Vals
		acts[n] = seg
	}
	sys.Run().Pll(acts...)
	return r
}
func (x *SideBse) Prd() Stm {
	r := &StmRtePrd{}
	r.Slf = r
	r.Side = x.Slf
	if ana.Cfg.Trc.IsHstStm() {
		sys.Logf("%p hst.StmRtePrd(%v)", r, r.Prm())
	}
	if x.ValBnds.Cnt() == 0 {
		return r
	}
	inrvl := x.Inrvl.Bse()
	if inrvl.TmeBnds.Cnt() != x.ValBnds.Cnt() {
		err.Panicf("length unequal (TmeBnds:%v ValBnds:%v)", inrvl.TmeBnds.Cnt(), x.ValBnds.Cnt())
	}
	r.Tmes = inrvl.Tmes
	r.Vals = flts.Make(inrvl.TmeBnds.Cnt())
	segBnds, acts := bnds.Segs(inrvl.TmeBnds.Cnt())
	for n, segBnd := range *segBnds {
		seg := &StmRtePrdSeg{}
		seg.Bnd = segBnd
		seg.Vals = x.Vals
		seg.ValBnds = x.ValBnds
		seg.Out = r.Vals
		acts[n] = seg
	}
	sys.Run().Pll(acts...)
	return r
}
func (x *SideBse) Min() Stm {
	r := &StmRteMin{}
	r.Slf = r
	r.Side = x.Slf
	if ana.Cfg.Trc.IsHstStm() {
		sys.Logf("%p hst.StmRteMin(%v)", r, r.Prm())
	}
	if x.ValBnds.Cnt() == 0 {
		return r
	}
	inrvl := x.Inrvl.Bse()
	if inrvl.TmeBnds.Cnt() != x.ValBnds.Cnt() {
		err.Panicf("length unequal (TmeBnds:%v ValBnds:%v)", inrvl.TmeBnds.Cnt(), x.ValBnds.Cnt())
	}
	r.Tmes = inrvl.Tmes
	r.Vals = flts.Make(inrvl.TmeBnds.Cnt())
	segBnds, acts := bnds.Segs(inrvl.TmeBnds.Cnt())
	for n, segBnd := range *segBnds {
		seg := &StmRteMinSeg{}
		seg.Bnd = segBnd
		seg.Vals = x.Vals
		seg.ValBnds = x.ValBnds
		seg.Out = r.Vals
		acts[n] = seg
	}
	sys.Run().Pll(acts...)
	return r
}
func (x *SideBse) Max() Stm {
	r := &StmRteMax{}
	r.Slf = r
	r.Side = x.Slf
	if ana.Cfg.Trc.IsHstStm() {
		sys.Logf("%p hst.StmRteMax(%v)", r, r.Prm())
	}
	if x.ValBnds.Cnt() == 0 {
		return r
	}
	inrvl := x.Inrvl.Bse()
	if inrvl.TmeBnds.Cnt() != x.ValBnds.Cnt() {
		err.Panicf("length unequal (TmeBnds:%v ValBnds:%v)", inrvl.TmeBnds.Cnt(), x.ValBnds.Cnt())
	}
	r.Tmes = inrvl.Tmes
	r.Vals = flts.Make(inrvl.TmeBnds.Cnt())
	segBnds, acts := bnds.Segs(inrvl.TmeBnds.Cnt())
	for n, segBnd := range *segBnds {
		seg := &StmRteMaxSeg{}
		seg.Bnd = segBnd
		seg.Vals = x.Vals
		seg.ValBnds = x.ValBnds
		seg.Out = r.Vals
		acts[n] = seg
	}
	sys.Run().Pll(acts...)
	return r
}
func (x *SideBse) Mid() Stm {
	r := &StmRteMid{}
	r.Slf = r
	r.Side = x.Slf
	if ana.Cfg.Trc.IsHstStm() {
		sys.Logf("%p hst.StmRteMid(%v)", r, r.Prm())
	}
	if x.ValBnds.Cnt() == 0 {
		return r
	}
	inrvl := x.Inrvl.Bse()
	if inrvl.TmeBnds.Cnt() != x.ValBnds.Cnt() {
		err.Panicf("length unequal (TmeBnds:%v ValBnds:%v)", inrvl.TmeBnds.Cnt(), x.ValBnds.Cnt())
	}
	r.Tmes = inrvl.Tmes
	r.Vals = flts.Make(inrvl.TmeBnds.Cnt())
	segBnds, acts := bnds.Segs(inrvl.TmeBnds.Cnt())
	for n, segBnd := range *segBnds {
		seg := &StmRteMidSeg{}
		seg.Bnd = segBnd
		seg.Vals = x.Vals
		seg.ValBnds = x.ValBnds
		seg.Out = r.Vals
		acts[n] = seg
	}
	sys.Run().Pll(acts...)
	return r
}
func (x *SideBse) Mdn() Stm {
	r := &StmRteMdn{}
	r.Slf = r
	r.Side = x.Slf
	if ana.Cfg.Trc.IsHstStm() {
		sys.Logf("%p hst.StmRteMdn(%v)", r, r.Prm())
	}
	if x.ValBnds.Cnt() == 0 {
		return r
	}
	inrvl := x.Inrvl.Bse()
	if inrvl.TmeBnds.Cnt() != x.ValBnds.Cnt() {
		err.Panicf("length unequal (TmeBnds:%v ValBnds:%v)", inrvl.TmeBnds.Cnt(), x.ValBnds.Cnt())
	}
	r.Tmes = inrvl.Tmes
	r.Vals = flts.Make(inrvl.TmeBnds.Cnt())
	segBnds, acts := bnds.Segs(inrvl.TmeBnds.Cnt())
	for n, segBnd := range *segBnds {
		seg := &StmRteMdnSeg{}
		seg.Bnd = segBnd
		seg.Vals = x.Vals
		seg.ValBnds = x.ValBnds
		seg.Out = r.Vals
		acts[n] = seg
	}
	sys.Run().Pll(acts...)
	return r
}
func (x *SideBse) Sma() Stm {
	r := &StmRteSma{}
	r.Slf = r
	r.Side = x.Slf
	if ana.Cfg.Trc.IsHstStm() {
		sys.Logf("%p hst.StmRteSma(%v)", r, r.Prm())
	}
	if x.ValBnds.Cnt() == 0 {
		return r
	}
	inrvl := x.Inrvl.Bse()
	if inrvl.TmeBnds.Cnt() != x.ValBnds.Cnt() {
		err.Panicf("length unequal (TmeBnds:%v ValBnds:%v)", inrvl.TmeBnds.Cnt(), x.ValBnds.Cnt())
	}
	r.Tmes = inrvl.Tmes
	r.Vals = flts.Make(inrvl.TmeBnds.Cnt())
	segBnds, acts := bnds.Segs(inrvl.TmeBnds.Cnt())
	for n, segBnd := range *segBnds {
		seg := &StmRteSmaSeg{}
		seg.Bnd = segBnd
		seg.Vals = x.Vals
		seg.ValBnds = x.ValBnds
		seg.Out = r.Vals
		acts[n] = seg
	}
	sys.Run().Pll(acts...)
	return r
}
func (x *SideBse) Gma() Stm {
	r := &StmRteGma{}
	r.Slf = r
	r.Side = x.Slf
	if ana.Cfg.Trc.IsHstStm() {
		sys.Logf("%p hst.StmRteGma(%v)", r, r.Prm())
	}
	if x.ValBnds.Cnt() == 0 {
		return r
	}
	inrvl := x.Inrvl.Bse()
	if inrvl.TmeBnds.Cnt() != x.ValBnds.Cnt() {
		err.Panicf("length unequal (TmeBnds:%v ValBnds:%v)", inrvl.TmeBnds.Cnt(), x.ValBnds.Cnt())
	}
	r.Tmes = inrvl.Tmes
	r.Vals = flts.Make(inrvl.TmeBnds.Cnt())
	segBnds, acts := bnds.Segs(inrvl.TmeBnds.Cnt())
	for n, segBnd := range *segBnds {
		seg := &StmRteGmaSeg{}
		seg.Bnd = segBnd
		seg.Vals = x.Vals
		seg.ValBnds = x.ValBnds
		seg.Out = r.Vals
		acts[n] = seg
	}
	sys.Run().Pll(acts...)
	return r
}
func (x *SideBse) Wma() Stm {
	r := &StmRteWma{}
	r.Slf = r
	r.Side = x.Slf
	if ana.Cfg.Trc.IsHstStm() {
		sys.Logf("%p hst.StmRteWma(%v)", r, r.Prm())
	}
	if x.ValBnds.Cnt() == 0 {
		return r
	}
	inrvl := x.Inrvl.Bse()
	if inrvl.TmeBnds.Cnt() != x.ValBnds.Cnt() {
		err.Panicf("length unequal (TmeBnds:%v ValBnds:%v)", inrvl.TmeBnds.Cnt(), x.ValBnds.Cnt())
	}
	r.Tmes = inrvl.Tmes
	r.Vals = flts.Make(inrvl.TmeBnds.Cnt())
	segBnds, acts := bnds.Segs(inrvl.TmeBnds.Cnt())
	for n, segBnd := range *segBnds {
		seg := &StmRteWmaSeg{}
		seg.Bnd = segBnd
		seg.Vals = x.Vals
		seg.ValBnds = x.ValBnds
		seg.Out = r.Vals
		acts[n] = seg
	}
	sys.Run().Pll(acts...)
	return r
}
func (x *SideBse) Rsi() Stm {
	r := &StmRteRsi{}
	r.Slf = r
	r.Side = x.Slf
	if ana.Cfg.Trc.IsHstStm() {
		sys.Logf("%p hst.StmRteRsi(%v)", r, r.Prm())
	}
	if x.ValBnds.Cnt() == 0 {
		return r
	}
	inrvl := x.Inrvl.Bse()
	if inrvl.TmeBnds.Cnt() != x.ValBnds.Cnt() {
		err.Panicf("length unequal (TmeBnds:%v ValBnds:%v)", inrvl.TmeBnds.Cnt(), x.ValBnds.Cnt())
	}
	r.Tmes = inrvl.Tmes
	r.Vals = flts.Make(inrvl.TmeBnds.Cnt())
	segBnds, acts := bnds.Segs(inrvl.TmeBnds.Cnt())
	for n, segBnd := range *segBnds {
		seg := &StmRteRsiSeg{}
		seg.Bnd = segBnd
		seg.Vals = x.Vals
		seg.ValBnds = x.ValBnds
		seg.Out = r.Vals
		acts[n] = seg
	}
	sys.Run().Pll(acts...)
	return r
}
func (x *SideBse) Wrsi() Stm {
	r := &StmRteWrsi{}
	r.Slf = r
	r.Side = x.Slf
	if ana.Cfg.Trc.IsHstStm() {
		sys.Logf("%p hst.StmRteWrsi(%v)", r, r.Prm())
	}
	if x.ValBnds.Cnt() == 0 {
		return r
	}
	inrvl := x.Inrvl.Bse()
	if inrvl.TmeBnds.Cnt() != x.ValBnds.Cnt() {
		err.Panicf("length unequal (TmeBnds:%v ValBnds:%v)", inrvl.TmeBnds.Cnt(), x.ValBnds.Cnt())
	}
	r.Tmes = inrvl.Tmes
	r.Vals = flts.Make(inrvl.TmeBnds.Cnt())
	segBnds, acts := bnds.Segs(inrvl.TmeBnds.Cnt())
	for n, segBnd := range *segBnds {
		seg := &StmRteWrsiSeg{}
		seg.Bnd = segBnd
		seg.Vals = x.Vals
		seg.ValBnds = x.ValBnds
		seg.Out = r.Vals
		acts[n] = seg
	}
	sys.Run().Pll(acts...)
	return r
}
func (x *SideBse) Alma() Stm {
	r := &StmRteAlma{}
	r.Slf = r
	r.Side = x.Slf
	if ana.Cfg.Trc.IsHstStm() {
		sys.Logf("%p hst.StmRteAlma(%v)", r, r.Prm())
	}
	if x.ValBnds.Cnt() == 0 {
		return r
	}
	inrvl := x.Inrvl.Bse()
	if inrvl.TmeBnds.Cnt() != x.ValBnds.Cnt() {
		err.Panicf("length unequal (TmeBnds:%v ValBnds:%v)", inrvl.TmeBnds.Cnt(), x.ValBnds.Cnt())
	}
	r.Tmes = inrvl.Tmes
	r.Vals = flts.Make(inrvl.TmeBnds.Cnt())
	segBnds, acts := bnds.Segs(inrvl.TmeBnds.Cnt())
	for n, segBnd := range *segBnds {
		seg := &StmRteAlmaSeg{}
		seg.Bnd = segBnd
		seg.Vals = x.Vals
		seg.ValBnds = x.ValBnds
		seg.Out = r.Vals
		acts[n] = seg
	}
	sys.Run().Pll(acts...)
	return r
}
func (x *SideBse) Vrnc() Stm {
	r := &StmRteVrnc{}
	r.Slf = r
	r.Side = x.Slf
	if ana.Cfg.Trc.IsHstStm() {
		sys.Logf("%p hst.StmRteVrnc(%v)", r, r.Prm())
	}
	if x.ValBnds.Cnt() == 0 {
		return r
	}
	inrvl := x.Inrvl.Bse()
	if inrvl.TmeBnds.Cnt() != x.ValBnds.Cnt() {
		err.Panicf("length unequal (TmeBnds:%v ValBnds:%v)", inrvl.TmeBnds.Cnt(), x.ValBnds.Cnt())
	}
	r.Tmes = inrvl.Tmes
	r.Vals = flts.Make(inrvl.TmeBnds.Cnt())
	segBnds, acts := bnds.Segs(inrvl.TmeBnds.Cnt())
	for n, segBnd := range *segBnds {
		seg := &StmRteVrncSeg{}
		seg.Bnd = segBnd
		seg.Vals = x.Vals
		seg.ValBnds = x.ValBnds
		seg.Out = r.Vals
		acts[n] = seg
	}
	sys.Run().Pll(acts...)
	return r
}
func (x *SideBse) Std() Stm {
	r := &StmRteStd{}
	r.Slf = r
	r.Side = x.Slf
	if ana.Cfg.Trc.IsHstStm() {
		sys.Logf("%p hst.StmRteStd(%v)", r, r.Prm())
	}
	if x.ValBnds.Cnt() == 0 {
		return r
	}
	inrvl := x.Inrvl.Bse()
	if inrvl.TmeBnds.Cnt() != x.ValBnds.Cnt() {
		err.Panicf("length unequal (TmeBnds:%v ValBnds:%v)", inrvl.TmeBnds.Cnt(), x.ValBnds.Cnt())
	}
	r.Tmes = inrvl.Tmes
	r.Vals = flts.Make(inrvl.TmeBnds.Cnt())
	segBnds, acts := bnds.Segs(inrvl.TmeBnds.Cnt())
	for n, segBnd := range *segBnds {
		seg := &StmRteStdSeg{}
		seg.Bnd = segBnd
		seg.Vals = x.Vals
		seg.ValBnds = x.ValBnds
		seg.Out = r.Vals
		acts[n] = seg
	}
	sys.Run().Pll(acts...)
	return r
}
func (x *SideBse) RngFul() Stm {
	r := &StmRteRngFul{}
	r.Slf = r
	r.Side = x.Slf
	if ana.Cfg.Trc.IsHstStm() {
		sys.Logf("%p hst.StmRteRngFul(%v)", r, r.Prm())
	}
	if x.ValBnds.Cnt() == 0 {
		return r
	}
	inrvl := x.Inrvl.Bse()
	if inrvl.TmeBnds.Cnt() != x.ValBnds.Cnt() {
		err.Panicf("length unequal (TmeBnds:%v ValBnds:%v)", inrvl.TmeBnds.Cnt(), x.ValBnds.Cnt())
	}
	r.Tmes = inrvl.Tmes
	r.Vals = flts.Make(inrvl.TmeBnds.Cnt())
	segBnds, acts := bnds.Segs(inrvl.TmeBnds.Cnt())
	for n, segBnd := range *segBnds {
		seg := &StmRteRngFulSeg{}
		seg.Bnd = segBnd
		seg.Vals = x.Vals
		seg.ValBnds = x.ValBnds
		seg.Out = r.Vals
		acts[n] = seg
	}
	sys.Run().Pll(acts...)
	return r
}
func (x *SideBse) RngLst() Stm {
	r := &StmRteRngLst{}
	r.Slf = r
	r.Side = x.Slf
	if ana.Cfg.Trc.IsHstStm() {
		sys.Logf("%p hst.StmRteRngLst(%v)", r, r.Prm())
	}
	if x.ValBnds.Cnt() == 0 {
		return r
	}
	inrvl := x.Inrvl.Bse()
	if inrvl.TmeBnds.Cnt() != x.ValBnds.Cnt() {
		err.Panicf("length unequal (TmeBnds:%v ValBnds:%v)", inrvl.TmeBnds.Cnt(), x.ValBnds.Cnt())
	}
	r.Tmes = inrvl.Tmes
	r.Vals = flts.Make(inrvl.TmeBnds.Cnt())
	segBnds, acts := bnds.Segs(inrvl.TmeBnds.Cnt())
	for n, segBnd := range *segBnds {
		seg := &StmRteRngLstSeg{}
		seg.Bnd = segBnd
		seg.Vals = x.Vals
		seg.ValBnds = x.ValBnds
		seg.Out = r.Vals
		acts[n] = seg
	}
	sys.Run().Pll(acts...)
	return r
}
func (x *SideBse) ProLst() Stm {
	r := &StmRteProLst{}
	r.Slf = r
	r.Side = x.Slf
	if ana.Cfg.Trc.IsHstStm() {
		sys.Logf("%p hst.StmRteProLst(%v)", r, r.Prm())
	}
	if x.ValBnds.Cnt() == 0 {
		return r
	}
	inrvl := x.Inrvl.Bse()
	if inrvl.TmeBnds.Cnt() != x.ValBnds.Cnt() {
		err.Panicf("length unequal (TmeBnds:%v ValBnds:%v)", inrvl.TmeBnds.Cnt(), x.ValBnds.Cnt())
	}
	r.Tmes = inrvl.Tmes
	r.Vals = flts.Make(inrvl.TmeBnds.Cnt())
	segBnds, acts := bnds.Segs(inrvl.TmeBnds.Cnt())
	for n, segBnd := range *segBnds {
		seg := &StmRteProLstSeg{}
		seg.Bnd = segBnd
		seg.Vals = x.Vals
		seg.ValBnds = x.ValBnds
		seg.Out = r.Vals
		acts[n] = seg
	}
	sys.Run().Pll(acts...)
	return r
}
func (x *SideBse) ProSma() Stm {
	r := &StmRteProSma{}
	r.Slf = r
	r.Side = x.Slf
	if ana.Cfg.Trc.IsHstStm() {
		sys.Logf("%p hst.StmRteProSma(%v)", r, r.Prm())
	}
	if x.ValBnds.Cnt() == 0 {
		return r
	}
	inrvl := x.Inrvl.Bse()
	if inrvl.TmeBnds.Cnt() != x.ValBnds.Cnt() {
		err.Panicf("length unequal (TmeBnds:%v ValBnds:%v)", inrvl.TmeBnds.Cnt(), x.ValBnds.Cnt())
	}
	r.Tmes = inrvl.Tmes
	r.Vals = flts.Make(inrvl.TmeBnds.Cnt())
	segBnds, acts := bnds.Segs(inrvl.TmeBnds.Cnt())
	for n, segBnd := range *segBnds {
		seg := &StmRteProSmaSeg{}
		seg.Bnd = segBnd
		seg.Vals = x.Vals
		seg.ValBnds = x.ValBnds
		seg.Out = r.Vals
		acts[n] = seg
	}
	sys.Run().Pll(acts...)
	return r
}
func (x *SideBse) ProAlma() Stm {
	r := &StmRteProAlma{}
	r.Slf = r
	r.Side = x.Slf
	if ana.Cfg.Trc.IsHstStm() {
		sys.Logf("%p hst.StmRteProAlma(%v)", r, r.Prm())
	}
	if x.ValBnds.Cnt() == 0 {
		return r
	}
	inrvl := x.Inrvl.Bse()
	if inrvl.TmeBnds.Cnt() != x.ValBnds.Cnt() {
		err.Panicf("length unequal (TmeBnds:%v ValBnds:%v)", inrvl.TmeBnds.Cnt(), x.ValBnds.Cnt())
	}
	r.Tmes = inrvl.Tmes
	r.Vals = flts.Make(inrvl.TmeBnds.Cnt())
	segBnds, acts := bnds.Segs(inrvl.TmeBnds.Cnt())
	for n, segBnd := range *segBnds {
		seg := &StmRteProAlmaSeg{}
		seg.Bnd = segBnd
		seg.Vals = x.Vals
		seg.ValBnds = x.ValBnds
		seg.Out = r.Vals
		acts[n] = seg
	}
	sys.Run().Pll(acts...)
	return r
}
func (x *SideBse) Sar(afInc, afMax flt.Flt) Stm {
	r := &StmRte1Sar{}
	r.Slf = r
	r.Side = x.Slf
	r.AfInc = afInc
	r.AfMax = afMax
	if ana.Cfg.Trc.IsHstStm() {
		sys.Logf("%p hst.StmRte1Sar(%v)", r, r.Prm())
	}
	if x.ValBnds.Cnt() == 0 {
		return r
	}
	if x.Inrvl.Bse().TmeBnds.Cnt() != x.ValBnds.Cnt() {
		err.Panicf("length unequal (TmeBnds:%v ValBnds:%v)", x.Inrvl.Bse().TmeBnds.Cnt(), x.ValBnds.Cnt())
	}
	r.Calc()
	return r
}
func (x *SideBse) Ema() Stm {
	r := &StmRteEma{}
	r.Slf = r
	r.Side = x.Slf
	if ana.Cfg.Trc.IsHstStm() {
		sys.Logf("%p hst.StmRteEma(%v)", r, r.Prm())
	}
	if x.ValBnds.Cnt() == 0 {
		return r
	}
	if x.Inrvl.Bse().TmeBnds.Cnt() != x.ValBnds.Cnt() {
		err.Panicf("length unequal (TmeBnds:%v ValBnds:%v)", x.Inrvl.Bse().TmeBnds.Cnt(), x.ValBnds.Cnt())
	}
	r.Tmes = x.Inrvl.Bse().Tmes
	r.Vals = flts.Make(x.Inrvl.Bse().TmeBnds.Cnt())
	// NON-PLL IMPL DUE TO PRV VAL CHAINING
	//    EMA CALC   https://stockcharts.com/school/doku.php?id=chart_school:technical_indicators:moving_averages
	// Initial SMA: 10-period sum / 10
	// Multiplier: (2 / (Time periods + 1) ) = (2 / (10 + 1) ) = 0.1818 (18.18%)
	// EMA: {Close - EMA(previous day)} x multiplier + EMA(previous day)
	// NOTE: Each ValBnd may have different lengths; multiplier will vary from tick to tick for Side
	(*r.Vals)[0] = x.Vals.InBnd(x.ValBnds.At(0)).Sma()
	for n := 1; n < len(*x.ValBnds); n++ {
		(*r.Vals)[n] = (*r.Vals)[n-1] + ((*x.Vals)[(*x.ValBnds)[n].Lim-1]-(*r.Vals)[n-1])*(flt.Flt(2)/flt.Flt(x.ValBnds.At(unt.Unt(n)).Len()+1))
	}
	return r
}
