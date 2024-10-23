package rlt

import (
	"strings"
	"sync"
	"sys"
	"sys/ana"
	"sys/bsc/bol"
	"sys/bsc/flt"
	"sys/bsc/flts"
	"sys/bsc/str"
	"sys/bsc/tme"
	"sys/bsc/tmes"
	"sys/bsc/unt"
)

type (
	Stm interface {
		ana.Pth
		Unsub(id uint32, slot ...uint32)
		DstToInstr() int
		Bse() *StmBse
		Sub(rx ana.TmeFltRx, id uint32, slot ...uint32)
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
		mu  sync.Mutex
		Id  uint32
		Slf Stm
		Rxs ana.TmeFltRxs
	}
	StmScp struct {
		Idx uint32
		Arr []Stm
	}
	StmRteFst struct {
		StmBse
		Side Side
	}
	StmRteLst struct {
		StmBse
		Side Side
	}
	StmRteSum struct {
		StmBse
		Side Side
	}
	StmRtePrd struct {
		StmBse
		Side Side
	}
	StmRteMin struct {
		StmBse
		Side Side
	}
	StmRteMax struct {
		StmBse
		Side Side
	}
	StmRteMid struct {
		StmBse
		Side Side
	}
	StmRteMdn struct {
		StmBse
		Side Side
	}
	StmRteSma struct {
		StmBse
		Side Side
	}
	StmRteGma struct {
		StmBse
		Side Side
	}
	StmRteWma struct {
		StmBse
		Side Side
	}
	StmRteRsi struct {
		StmBse
		Side Side
	}
	StmRteWrsi struct {
		StmBse
		Side Side
	}
	StmRteAlma struct {
		StmBse
		Side Side
	}
	StmRteVrnc struct {
		StmBse
		Side Side
	}
	StmRteStd struct {
		StmBse
		Side Side
	}
	StmRteRngFul struct {
		StmBse
		Side Side
	}
	StmRteRngLst struct {
		StmBse
		Side Side
	}
	StmRteProLst struct {
		StmBse
		Side Side
	}
	StmRteProSma struct {
		StmBse
		Side Side
	}
	StmRteProAlma struct {
		StmBse
		Side Side
	}
	StmRte1Sar struct {
		StmBse
		Side   Side
		AfInc  flt.Flt
		AfMax  flt.Flt
		IsLong bol.Bol
		Sar    flt.Flt
		Ep     flt.Flt
		Af     flt.Flt
		PrvLo  flt.Flt
		PrvHi  flt.Flt
	}
	StmRteEma struct {
		StmBse
		Side Side
		Prv  flt.Flt
	}
	StmUnaPos struct {
		StmBse
		Stm Stm
	}
	StmUnaNeg struct {
		StmBse
		Stm Stm
	}
	StmUnaInv struct {
		StmBse
		Stm Stm
	}
	StmUnaSqr struct {
		StmBse
		Stm Stm
	}
	StmUnaSqrt struct {
		StmBse
		Stm Stm
	}
	StmSclAdd struct {
		StmBse
		Stm Stm
		Scl flt.Flt
	}
	StmSclSub struct {
		StmBse
		Stm Stm
		Scl flt.Flt
	}
	StmSclMul struct {
		StmBse
		Stm Stm
		Scl flt.Flt
	}
	StmSclDiv struct {
		StmBse
		Stm Stm
		Scl flt.Flt
	}
	StmSclRem struct {
		StmBse
		Stm Stm
		Scl flt.Flt
	}
	StmSclPow struct {
		StmBse
		Stm Stm
		Scl flt.Flt
	}
	StmSclMin struct {
		StmBse
		Stm Stm
		Scl flt.Flt
	}
	StmSclMax struct {
		StmBse
		Stm Stm
		Scl flt.Flt
	}
	StmSelEql struct {
		StmBse
		Stm Stm
		Sel flt.Flt
	}
	StmSelNeq struct {
		StmBse
		Stm Stm
		Sel flt.Flt
	}
	StmSelLss struct {
		StmBse
		Stm Stm
		Sel flt.Flt
	}
	StmSelGtr struct {
		StmBse
		Stm Stm
		Sel flt.Flt
	}
	StmSelLeq struct {
		StmBse
		Stm Stm
		Sel flt.Flt
	}
	StmSelGeq struct {
		StmBse
		Stm Stm
		Sel flt.Flt
	}
	StmAggFst struct {
		StmBse
		Stm    Stm
		Tmes   *tmes.Tmes
		Vals   *flts.Flts
		Length unt.Unt
	}
	StmAggLst struct {
		StmBse
		Stm    Stm
		Tmes   *tmes.Tmes
		Vals   *flts.Flts
		Length unt.Unt
	}
	StmAggSum struct {
		StmBse
		Stm    Stm
		Tmes   *tmes.Tmes
		Vals   *flts.Flts
		Length unt.Unt
	}
	StmAggPrd struct {
		StmBse
		Stm    Stm
		Tmes   *tmes.Tmes
		Vals   *flts.Flts
		Length unt.Unt
	}
	StmAggMin struct {
		StmBse
		Stm    Stm
		Tmes   *tmes.Tmes
		Vals   *flts.Flts
		Length unt.Unt
	}
	StmAggMax struct {
		StmBse
		Stm    Stm
		Tmes   *tmes.Tmes
		Vals   *flts.Flts
		Length unt.Unt
	}
	StmAggMid struct {
		StmBse
		Stm    Stm
		Tmes   *tmes.Tmes
		Vals   *flts.Flts
		Length unt.Unt
	}
	StmAggMdn struct {
		StmBse
		Stm    Stm
		Tmes   *tmes.Tmes
		Vals   *flts.Flts
		Length unt.Unt
	}
	StmAggSma struct {
		StmBse
		Stm    Stm
		Tmes   *tmes.Tmes
		Vals   *flts.Flts
		Length unt.Unt
	}
	StmAggGma struct {
		StmBse
		Stm    Stm
		Tmes   *tmes.Tmes
		Vals   *flts.Flts
		Length unt.Unt
	}
	StmAggWma struct {
		StmBse
		Stm    Stm
		Tmes   *tmes.Tmes
		Vals   *flts.Flts
		Length unt.Unt
	}
	StmAggRsi struct {
		StmBse
		Stm    Stm
		Tmes   *tmes.Tmes
		Vals   *flts.Flts
		Length unt.Unt
	}
	StmAggWrsi struct {
		StmBse
		Stm    Stm
		Tmes   *tmes.Tmes
		Vals   *flts.Flts
		Length unt.Unt
	}
	StmAggAlma struct {
		StmBse
		Stm    Stm
		Tmes   *tmes.Tmes
		Vals   *flts.Flts
		Length unt.Unt
	}
	StmAggVrnc struct {
		StmBse
		Stm    Stm
		Tmes   *tmes.Tmes
		Vals   *flts.Flts
		Length unt.Unt
	}
	StmAggStd struct {
		StmBse
		Stm    Stm
		Tmes   *tmes.Tmes
		Vals   *flts.Flts
		Length unt.Unt
	}
	StmAggRngFul struct {
		StmBse
		Stm    Stm
		Tmes   *tmes.Tmes
		Vals   *flts.Flts
		Length unt.Unt
	}
	StmAggRngLst struct {
		StmBse
		Stm    Stm
		Tmes   *tmes.Tmes
		Vals   *flts.Flts
		Length unt.Unt
	}
	StmAggProLst struct {
		StmBse
		Stm    Stm
		Tmes   *tmes.Tmes
		Vals   *flts.Flts
		Length unt.Unt
	}
	StmAggProSma struct {
		StmBse
		Stm    Stm
		Tmes   *tmes.Tmes
		Vals   *flts.Flts
		Length unt.Unt
	}
	StmAggProAlma struct {
		StmBse
		Stm    Stm
		Tmes   *tmes.Tmes
		Vals   *flts.Flts
		Length unt.Unt
	}
	StmAggEma struct {
		StmBse
		Stm    Stm
		Prv    flt.Flt
		Tmes   *tmes.Tmes
		Vals   *flts.Flts
		Length unt.Unt
	}
	StmInrAdd struct {
		StmBse
		Stm  Stm
		Tmes *tmes.Tmes
		Vals *flts.Flts
		Off  unt.Unt
	}
	StmInrSub struct {
		StmBse
		Stm  Stm
		Tmes *tmes.Tmes
		Vals *flts.Flts
		Off  unt.Unt
	}
	StmInrMul struct {
		StmBse
		Stm  Stm
		Tmes *tmes.Tmes
		Vals *flts.Flts
		Off  unt.Unt
	}
	StmInrDiv struct {
		StmBse
		Stm  Stm
		Tmes *tmes.Tmes
		Vals *flts.Flts
		Off  unt.Unt
	}
	StmInrRem struct {
		StmBse
		Stm  Stm
		Tmes *tmes.Tmes
		Vals *flts.Flts
		Off  unt.Unt
	}
	StmInrPow struct {
		StmBse
		Stm  Stm
		Tmes *tmes.Tmes
		Vals *flts.Flts
		Off  unt.Unt
	}
	StmInrMin struct {
		StmBse
		Stm  Stm
		Tmes *tmes.Tmes
		Vals *flts.Flts
		Off  unt.Unt
	}
	StmInrMax struct {
		StmBse
		Stm  Stm
		Tmes *tmes.Tmes
		Vals *flts.Flts
		Off  unt.Unt
	}
	StmInr1Slp struct {
		StmBse
		Stm  Stm
		Tmes *tmes.Tmes
		Vals *flts.Flts
		Off  unt.Unt
	}
	StmOtrAdd struct {
		StmBse
		Stm   Stm
		Tmes  *tmes.Tmes
		Vals  *flts.Flts
		TmesA *tmes.Tmes
		ValsA *flts.Flts
		Off   unt.Unt
		A     Stm
	}
	StmOtrSub struct {
		StmBse
		Stm   Stm
		Tmes  *tmes.Tmes
		Vals  *flts.Flts
		TmesA *tmes.Tmes
		ValsA *flts.Flts
		Off   unt.Unt
		A     Stm
	}
	StmOtrMul struct {
		StmBse
		Stm   Stm
		Tmes  *tmes.Tmes
		Vals  *flts.Flts
		TmesA *tmes.Tmes
		ValsA *flts.Flts
		Off   unt.Unt
		A     Stm
	}
	StmOtrDiv struct {
		StmBse
		Stm   Stm
		Tmes  *tmes.Tmes
		Vals  *flts.Flts
		TmesA *tmes.Tmes
		ValsA *flts.Flts
		Off   unt.Unt
		A     Stm
	}
	StmOtrRem struct {
		StmBse
		Stm   Stm
		Tmes  *tmes.Tmes
		Vals  *flts.Flts
		TmesA *tmes.Tmes
		ValsA *flts.Flts
		Off   unt.Unt
		A     Stm
	}
	StmOtrPow struct {
		StmBse
		Stm   Stm
		Tmes  *tmes.Tmes
		Vals  *flts.Flts
		TmesA *tmes.Tmes
		ValsA *flts.Flts
		Off   unt.Unt
		A     Stm
	}
	StmOtrMin struct {
		StmBse
		Stm   Stm
		Tmes  *tmes.Tmes
		Vals  *flts.Flts
		TmesA *tmes.Tmes
		ValsA *flts.Flts
		Off   unt.Unt
		A     Stm
	}
	StmOtrMax struct {
		StmBse
		Stm   Stm
		Tmes  *tmes.Tmes
		Vals  *flts.Flts
		TmesA *tmes.Tmes
		ValsA *flts.Flts
		Off   unt.Unt
		A     Stm
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
func (x *StmRteFst) DstToInstr() int { return 3 }
func (x *StmRteFst) Unsub(id uint32, slot ...uint32) {
	var uSlot uint32
	if len(slot) > 0 {
		uSlot = slot[0]
	}
	x.mu.Lock()
	delete(x.Rxs, sys.Uint64(id, uSlot))
	if len(x.Rxs) == 0 {
		x.Side.Unsub(x.Id)
	}
	x.mu.Unlock()
}
func (x *StmRteFst) Rx(inPkt ana.TmeFlts) (r []sys.Act) {
	x.mu.Lock()
	if ana.Cfg.Trc.IsRltStm() {
		sys.Logf("rlt.StmRteFst(%v).Fst %p inPkt %v", x.Prm(), x, inPkt)
	}
	outPkt := ana.TmeFlt{Tme: inPkt.Tme, Flt: inPkt.Flts.Fst()}
	for _, rx := range x.Rxs {
		r = append(r, &ana.TmeFltTx{Pkt: outPkt, Rx: rx})
	}
	x.mu.Unlock()
	return r
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
func (x *StmRteLst) DstToInstr() int { return 3 }
func (x *StmRteLst) Unsub(id uint32, slot ...uint32) {
	var uSlot uint32
	if len(slot) > 0 {
		uSlot = slot[0]
	}
	x.mu.Lock()
	delete(x.Rxs, sys.Uint64(id, uSlot))
	if len(x.Rxs) == 0 {
		x.Side.Unsub(x.Id)
	}
	x.mu.Unlock()
}
func (x *StmRteLst) Rx(inPkt ana.TmeFlts) (r []sys.Act) {
	x.mu.Lock()
	if ana.Cfg.Trc.IsRltStm() {
		sys.Logf("rlt.StmRteLst(%v).Lst %p inPkt %v", x.Prm(), x, inPkt)
	}
	outPkt := ana.TmeFlt{Tme: inPkt.Tme, Flt: inPkt.Flts.Lst()}
	for _, rx := range x.Rxs {
		r = append(r, &ana.TmeFltTx{Pkt: outPkt, Rx: rx})
	}
	x.mu.Unlock()
	return r
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
func (x *StmRteSum) DstToInstr() int { return 3 }
func (x *StmRteSum) Unsub(id uint32, slot ...uint32) {
	var uSlot uint32
	if len(slot) > 0 {
		uSlot = slot[0]
	}
	x.mu.Lock()
	delete(x.Rxs, sys.Uint64(id, uSlot))
	if len(x.Rxs) == 0 {
		x.Side.Unsub(x.Id)
	}
	x.mu.Unlock()
}
func (x *StmRteSum) Rx(inPkt ana.TmeFlts) (r []sys.Act) {
	x.mu.Lock()
	if ana.Cfg.Trc.IsRltStm() {
		sys.Logf("rlt.StmRteSum(%v).Sum %p inPkt %v", x.Prm(), x, inPkt)
	}
	outPkt := ana.TmeFlt{Tme: inPkt.Tme, Flt: inPkt.Flts.Sum()}
	for _, rx := range x.Rxs {
		r = append(r, &ana.TmeFltTx{Pkt: outPkt, Rx: rx})
	}
	x.mu.Unlock()
	return r
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
func (x *StmRtePrd) DstToInstr() int { return 3 }
func (x *StmRtePrd) Unsub(id uint32, slot ...uint32) {
	var uSlot uint32
	if len(slot) > 0 {
		uSlot = slot[0]
	}
	x.mu.Lock()
	delete(x.Rxs, sys.Uint64(id, uSlot))
	if len(x.Rxs) == 0 {
		x.Side.Unsub(x.Id)
	}
	x.mu.Unlock()
}
func (x *StmRtePrd) Rx(inPkt ana.TmeFlts) (r []sys.Act) {
	x.mu.Lock()
	if ana.Cfg.Trc.IsRltStm() {
		sys.Logf("rlt.StmRtePrd(%v).Prd %p inPkt %v", x.Prm(), x, inPkt)
	}
	outPkt := ana.TmeFlt{Tme: inPkt.Tme, Flt: inPkt.Flts.Prd()}
	for _, rx := range x.Rxs {
		r = append(r, &ana.TmeFltTx{Pkt: outPkt, Rx: rx})
	}
	x.mu.Unlock()
	return r
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
func (x *StmRteMin) DstToInstr() int { return 3 }
func (x *StmRteMin) Unsub(id uint32, slot ...uint32) {
	var uSlot uint32
	if len(slot) > 0 {
		uSlot = slot[0]
	}
	x.mu.Lock()
	delete(x.Rxs, sys.Uint64(id, uSlot))
	if len(x.Rxs) == 0 {
		x.Side.Unsub(x.Id)
	}
	x.mu.Unlock()
}
func (x *StmRteMin) Rx(inPkt ana.TmeFlts) (r []sys.Act) {
	x.mu.Lock()
	if ana.Cfg.Trc.IsRltStm() {
		sys.Logf("rlt.StmRteMin(%v).Min %p inPkt %v", x.Prm(), x, inPkt)
	}
	outPkt := ana.TmeFlt{Tme: inPkt.Tme, Flt: inPkt.Flts.Min()}
	for _, rx := range x.Rxs {
		r = append(r, &ana.TmeFltTx{Pkt: outPkt, Rx: rx})
	}
	x.mu.Unlock()
	return r
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
func (x *StmRteMax) DstToInstr() int { return 3 }
func (x *StmRteMax) Unsub(id uint32, slot ...uint32) {
	var uSlot uint32
	if len(slot) > 0 {
		uSlot = slot[0]
	}
	x.mu.Lock()
	delete(x.Rxs, sys.Uint64(id, uSlot))
	if len(x.Rxs) == 0 {
		x.Side.Unsub(x.Id)
	}
	x.mu.Unlock()
}
func (x *StmRteMax) Rx(inPkt ana.TmeFlts) (r []sys.Act) {
	x.mu.Lock()
	if ana.Cfg.Trc.IsRltStm() {
		sys.Logf("rlt.StmRteMax(%v).Max %p inPkt %v", x.Prm(), x, inPkt)
	}
	outPkt := ana.TmeFlt{Tme: inPkt.Tme, Flt: inPkt.Flts.Max()}
	for _, rx := range x.Rxs {
		r = append(r, &ana.TmeFltTx{Pkt: outPkt, Rx: rx})
	}
	x.mu.Unlock()
	return r
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
func (x *StmRteMid) DstToInstr() int { return 3 }
func (x *StmRteMid) Unsub(id uint32, slot ...uint32) {
	var uSlot uint32
	if len(slot) > 0 {
		uSlot = slot[0]
	}
	x.mu.Lock()
	delete(x.Rxs, sys.Uint64(id, uSlot))
	if len(x.Rxs) == 0 {
		x.Side.Unsub(x.Id)
	}
	x.mu.Unlock()
}
func (x *StmRteMid) Rx(inPkt ana.TmeFlts) (r []sys.Act) {
	x.mu.Lock()
	if ana.Cfg.Trc.IsRltStm() {
		sys.Logf("rlt.StmRteMid(%v).Mid %p inPkt %v", x.Prm(), x, inPkt)
	}
	outPkt := ana.TmeFlt{Tme: inPkt.Tme, Flt: inPkt.Flts.Mid()}
	for _, rx := range x.Rxs {
		r = append(r, &ana.TmeFltTx{Pkt: outPkt, Rx: rx})
	}
	x.mu.Unlock()
	return r
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
func (x *StmRteMdn) DstToInstr() int { return 3 }
func (x *StmRteMdn) Unsub(id uint32, slot ...uint32) {
	var uSlot uint32
	if len(slot) > 0 {
		uSlot = slot[0]
	}
	x.mu.Lock()
	delete(x.Rxs, sys.Uint64(id, uSlot))
	if len(x.Rxs) == 0 {
		x.Side.Unsub(x.Id)
	}
	x.mu.Unlock()
}
func (x *StmRteMdn) Rx(inPkt ana.TmeFlts) (r []sys.Act) {
	x.mu.Lock()
	if ana.Cfg.Trc.IsRltStm() {
		sys.Logf("rlt.StmRteMdn(%v).Mdn %p inPkt %v", x.Prm(), x, inPkt)
	}
	outPkt := ana.TmeFlt{Tme: inPkt.Tme, Flt: inPkt.Flts.Mdn()}
	for _, rx := range x.Rxs {
		r = append(r, &ana.TmeFltTx{Pkt: outPkt, Rx: rx})
	}
	x.mu.Unlock()
	return r
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
func (x *StmRteSma) DstToInstr() int { return 3 }
func (x *StmRteSma) Unsub(id uint32, slot ...uint32) {
	var uSlot uint32
	if len(slot) > 0 {
		uSlot = slot[0]
	}
	x.mu.Lock()
	delete(x.Rxs, sys.Uint64(id, uSlot))
	if len(x.Rxs) == 0 {
		x.Side.Unsub(x.Id)
	}
	x.mu.Unlock()
}
func (x *StmRteSma) Rx(inPkt ana.TmeFlts) (r []sys.Act) {
	x.mu.Lock()
	if ana.Cfg.Trc.IsRltStm() {
		sys.Logf("rlt.StmRteSma(%v).Sma %p inPkt %v", x.Prm(), x, inPkt)
	}
	outPkt := ana.TmeFlt{Tme: inPkt.Tme, Flt: inPkt.Flts.Sma()}
	for _, rx := range x.Rxs {
		r = append(r, &ana.TmeFltTx{Pkt: outPkt, Rx: rx})
	}
	x.mu.Unlock()
	return r
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
func (x *StmRteGma) DstToInstr() int { return 3 }
func (x *StmRteGma) Unsub(id uint32, slot ...uint32) {
	var uSlot uint32
	if len(slot) > 0 {
		uSlot = slot[0]
	}
	x.mu.Lock()
	delete(x.Rxs, sys.Uint64(id, uSlot))
	if len(x.Rxs) == 0 {
		x.Side.Unsub(x.Id)
	}
	x.mu.Unlock()
}
func (x *StmRteGma) Rx(inPkt ana.TmeFlts) (r []sys.Act) {
	x.mu.Lock()
	if ana.Cfg.Trc.IsRltStm() {
		sys.Logf("rlt.StmRteGma(%v).Gma %p inPkt %v", x.Prm(), x, inPkt)
	}
	outPkt := ana.TmeFlt{Tme: inPkt.Tme, Flt: inPkt.Flts.Gma()}
	for _, rx := range x.Rxs {
		r = append(r, &ana.TmeFltTx{Pkt: outPkt, Rx: rx})
	}
	x.mu.Unlock()
	return r
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
func (x *StmRteWma) DstToInstr() int { return 3 }
func (x *StmRteWma) Unsub(id uint32, slot ...uint32) {
	var uSlot uint32
	if len(slot) > 0 {
		uSlot = slot[0]
	}
	x.mu.Lock()
	delete(x.Rxs, sys.Uint64(id, uSlot))
	if len(x.Rxs) == 0 {
		x.Side.Unsub(x.Id)
	}
	x.mu.Unlock()
}
func (x *StmRteWma) Rx(inPkt ana.TmeFlts) (r []sys.Act) {
	x.mu.Lock()
	if ana.Cfg.Trc.IsRltStm() {
		sys.Logf("rlt.StmRteWma(%v).Wma %p inPkt %v", x.Prm(), x, inPkt)
	}
	outPkt := ana.TmeFlt{Tme: inPkt.Tme, Flt: inPkt.Flts.Wma()}
	for _, rx := range x.Rxs {
		r = append(r, &ana.TmeFltTx{Pkt: outPkt, Rx: rx})
	}
	x.mu.Unlock()
	return r
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
func (x *StmRteRsi) DstToInstr() int { return 3 }
func (x *StmRteRsi) Unsub(id uint32, slot ...uint32) {
	var uSlot uint32
	if len(slot) > 0 {
		uSlot = slot[0]
	}
	x.mu.Lock()
	delete(x.Rxs, sys.Uint64(id, uSlot))
	if len(x.Rxs) == 0 {
		x.Side.Unsub(x.Id)
	}
	x.mu.Unlock()
}
func (x *StmRteRsi) Rx(inPkt ana.TmeFlts) (r []sys.Act) {
	x.mu.Lock()
	if ana.Cfg.Trc.IsRltStm() {
		sys.Logf("rlt.StmRteRsi(%v).Rsi %p inPkt %v", x.Prm(), x, inPkt)
	}
	outPkt := ana.TmeFlt{Tme: inPkt.Tme, Flt: inPkt.Flts.Rsi()}
	for _, rx := range x.Rxs {
		r = append(r, &ana.TmeFltTx{Pkt: outPkt, Rx: rx})
	}
	x.mu.Unlock()
	return r
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
func (x *StmRteWrsi) DstToInstr() int { return 3 }
func (x *StmRteWrsi) Unsub(id uint32, slot ...uint32) {
	var uSlot uint32
	if len(slot) > 0 {
		uSlot = slot[0]
	}
	x.mu.Lock()
	delete(x.Rxs, sys.Uint64(id, uSlot))
	if len(x.Rxs) == 0 {
		x.Side.Unsub(x.Id)
	}
	x.mu.Unlock()
}
func (x *StmRteWrsi) Rx(inPkt ana.TmeFlts) (r []sys.Act) {
	x.mu.Lock()
	if ana.Cfg.Trc.IsRltStm() {
		sys.Logf("rlt.StmRteWrsi(%v).Wrsi %p inPkt %v", x.Prm(), x, inPkt)
	}
	outPkt := ana.TmeFlt{Tme: inPkt.Tme, Flt: inPkt.Flts.Wrsi()}
	for _, rx := range x.Rxs {
		r = append(r, &ana.TmeFltTx{Pkt: outPkt, Rx: rx})
	}
	x.mu.Unlock()
	return r
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
func (x *StmRteAlma) DstToInstr() int { return 3 }
func (x *StmRteAlma) Unsub(id uint32, slot ...uint32) {
	var uSlot uint32
	if len(slot) > 0 {
		uSlot = slot[0]
	}
	x.mu.Lock()
	delete(x.Rxs, sys.Uint64(id, uSlot))
	if len(x.Rxs) == 0 {
		x.Side.Unsub(x.Id)
	}
	x.mu.Unlock()
}
func (x *StmRteAlma) Rx(inPkt ana.TmeFlts) (r []sys.Act) {
	x.mu.Lock()
	if ana.Cfg.Trc.IsRltStm() {
		sys.Logf("rlt.StmRteAlma(%v).Alma %p inPkt %v", x.Prm(), x, inPkt)
	}
	outPkt := ana.TmeFlt{Tme: inPkt.Tme, Flt: inPkt.Flts.Alma()}
	for _, rx := range x.Rxs {
		r = append(r, &ana.TmeFltTx{Pkt: outPkt, Rx: rx})
	}
	x.mu.Unlock()
	return r
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
func (x *StmRteVrnc) DstToInstr() int { return 3 }
func (x *StmRteVrnc) Unsub(id uint32, slot ...uint32) {
	var uSlot uint32
	if len(slot) > 0 {
		uSlot = slot[0]
	}
	x.mu.Lock()
	delete(x.Rxs, sys.Uint64(id, uSlot))
	if len(x.Rxs) == 0 {
		x.Side.Unsub(x.Id)
	}
	x.mu.Unlock()
}
func (x *StmRteVrnc) Rx(inPkt ana.TmeFlts) (r []sys.Act) {
	x.mu.Lock()
	if ana.Cfg.Trc.IsRltStm() {
		sys.Logf("rlt.StmRteVrnc(%v).Vrnc %p inPkt %v", x.Prm(), x, inPkt)
	}
	outPkt := ana.TmeFlt{Tme: inPkt.Tme, Flt: inPkt.Flts.Vrnc()}
	for _, rx := range x.Rxs {
		r = append(r, &ana.TmeFltTx{Pkt: outPkt, Rx: rx})
	}
	x.mu.Unlock()
	return r
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
func (x *StmRteStd) DstToInstr() int { return 3 }
func (x *StmRteStd) Unsub(id uint32, slot ...uint32) {
	var uSlot uint32
	if len(slot) > 0 {
		uSlot = slot[0]
	}
	x.mu.Lock()
	delete(x.Rxs, sys.Uint64(id, uSlot))
	if len(x.Rxs) == 0 {
		x.Side.Unsub(x.Id)
	}
	x.mu.Unlock()
}
func (x *StmRteStd) Rx(inPkt ana.TmeFlts) (r []sys.Act) {
	x.mu.Lock()
	if ana.Cfg.Trc.IsRltStm() {
		sys.Logf("rlt.StmRteStd(%v).Std %p inPkt %v", x.Prm(), x, inPkt)
	}
	outPkt := ana.TmeFlt{Tme: inPkt.Tme, Flt: inPkt.Flts.Std()}
	for _, rx := range x.Rxs {
		r = append(r, &ana.TmeFltTx{Pkt: outPkt, Rx: rx})
	}
	x.mu.Unlock()
	return r
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
func (x *StmRteRngFul) DstToInstr() int { return 3 }
func (x *StmRteRngFul) Unsub(id uint32, slot ...uint32) {
	var uSlot uint32
	if len(slot) > 0 {
		uSlot = slot[0]
	}
	x.mu.Lock()
	delete(x.Rxs, sys.Uint64(id, uSlot))
	if len(x.Rxs) == 0 {
		x.Side.Unsub(x.Id)
	}
	x.mu.Unlock()
}
func (x *StmRteRngFul) Rx(inPkt ana.TmeFlts) (r []sys.Act) {
	x.mu.Lock()
	if ana.Cfg.Trc.IsRltStm() {
		sys.Logf("rlt.StmRteRngFul(%v).RngFul %p inPkt %v", x.Prm(), x, inPkt)
	}
	outPkt := ana.TmeFlt{Tme: inPkt.Tme, Flt: inPkt.Flts.RngFul()}
	for _, rx := range x.Rxs {
		r = append(r, &ana.TmeFltTx{Pkt: outPkt, Rx: rx})
	}
	x.mu.Unlock()
	return r
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
func (x *StmRteRngLst) DstToInstr() int { return 3 }
func (x *StmRteRngLst) Unsub(id uint32, slot ...uint32) {
	var uSlot uint32
	if len(slot) > 0 {
		uSlot = slot[0]
	}
	x.mu.Lock()
	delete(x.Rxs, sys.Uint64(id, uSlot))
	if len(x.Rxs) == 0 {
		x.Side.Unsub(x.Id)
	}
	x.mu.Unlock()
}
func (x *StmRteRngLst) Rx(inPkt ana.TmeFlts) (r []sys.Act) {
	x.mu.Lock()
	if ana.Cfg.Trc.IsRltStm() {
		sys.Logf("rlt.StmRteRngLst(%v).RngLst %p inPkt %v", x.Prm(), x, inPkt)
	}
	outPkt := ana.TmeFlt{Tme: inPkt.Tme, Flt: inPkt.Flts.RngLst()}
	for _, rx := range x.Rxs {
		r = append(r, &ana.TmeFltTx{Pkt: outPkt, Rx: rx})
	}
	x.mu.Unlock()
	return r
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
func (x *StmRteProLst) DstToInstr() int { return 3 }
func (x *StmRteProLst) Unsub(id uint32, slot ...uint32) {
	var uSlot uint32
	if len(slot) > 0 {
		uSlot = slot[0]
	}
	x.mu.Lock()
	delete(x.Rxs, sys.Uint64(id, uSlot))
	if len(x.Rxs) == 0 {
		x.Side.Unsub(x.Id)
	}
	x.mu.Unlock()
}
func (x *StmRteProLst) Rx(inPkt ana.TmeFlts) (r []sys.Act) {
	x.mu.Lock()
	if ana.Cfg.Trc.IsRltStm() {
		sys.Logf("rlt.StmRteProLst(%v).ProLst %p inPkt %v", x.Prm(), x, inPkt)
	}
	outPkt := ana.TmeFlt{Tme: inPkt.Tme, Flt: inPkt.Flts.ProLst()}
	for _, rx := range x.Rxs {
		r = append(r, &ana.TmeFltTx{Pkt: outPkt, Rx: rx})
	}
	x.mu.Unlock()
	return r
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
func (x *StmRteProSma) DstToInstr() int { return 3 }
func (x *StmRteProSma) Unsub(id uint32, slot ...uint32) {
	var uSlot uint32
	if len(slot) > 0 {
		uSlot = slot[0]
	}
	x.mu.Lock()
	delete(x.Rxs, sys.Uint64(id, uSlot))
	if len(x.Rxs) == 0 {
		x.Side.Unsub(x.Id)
	}
	x.mu.Unlock()
}
func (x *StmRteProSma) Rx(inPkt ana.TmeFlts) (r []sys.Act) {
	x.mu.Lock()
	if ana.Cfg.Trc.IsRltStm() {
		sys.Logf("rlt.StmRteProSma(%v).ProSma %p inPkt %v", x.Prm(), x, inPkt)
	}
	outPkt := ana.TmeFlt{Tme: inPkt.Tme, Flt: inPkt.Flts.ProSma()}
	for _, rx := range x.Rxs {
		r = append(r, &ana.TmeFltTx{Pkt: outPkt, Rx: rx})
	}
	x.mu.Unlock()
	return r
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
func (x *StmRteProAlma) DstToInstr() int { return 3 }
func (x *StmRteProAlma) Unsub(id uint32, slot ...uint32) {
	var uSlot uint32
	if len(slot) > 0 {
		uSlot = slot[0]
	}
	x.mu.Lock()
	delete(x.Rxs, sys.Uint64(id, uSlot))
	if len(x.Rxs) == 0 {
		x.Side.Unsub(x.Id)
	}
	x.mu.Unlock()
}
func (x *StmRteProAlma) Rx(inPkt ana.TmeFlts) (r []sys.Act) {
	x.mu.Lock()
	if ana.Cfg.Trc.IsRltStm() {
		sys.Logf("rlt.StmRteProAlma(%v).ProAlma %p inPkt %v", x.Prm(), x, inPkt)
	}
	outPkt := ana.TmeFlt{Tme: inPkt.Tme, Flt: inPkt.Flts.ProAlma()}
	for _, rx := range x.Rxs {
		r = append(r, &ana.TmeFltTx{Pkt: outPkt, Rx: rx})
	}
	x.mu.Unlock()
	return r
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
func (x *StmRte1Sar) DstToInstr() int { return 3 }
func (x *StmRte1Sar) Unsub(id uint32, slot ...uint32) {
	var uSlot uint32
	if len(slot) > 0 {
		uSlot = slot[0]
	}
	x.mu.Lock()
	delete(x.Rxs, sys.Uint64(id, uSlot))
	if len(x.Rxs) == 0 {
		x.Side.Unsub(x.Id)
	}
	x.mu.Unlock()
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
func (x *StmRteEma) DstToInstr() int { return 3 }
func (x *StmRteEma) Unsub(id uint32, slot ...uint32) {
	var uSlot uint32
	if len(slot) > 0 {
		uSlot = slot[0]
	}
	x.mu.Lock()
	delete(x.Rxs, sys.Uint64(id, uSlot))
	if len(x.Rxs) == 0 {
		x.Side.Unsub(x.Id)
	}
	x.mu.Unlock()
}
func (x *StmRteEma) Rx(inPkt ana.TmeFlts) (r []sys.Act) {
	x.mu.Lock()
	if ana.Cfg.Trc.IsRltStm() {
		sys.Logf("rlt.StmRteEma(%v).Rx %p inPkt %v", x.Prm(), x, inPkt)
	}
	outPkt := ana.TmeFlt{Tme: inPkt.Tme}
	if x.Prv == flt.Min {
		outPkt.Flt = inPkt.Flts.Sma()
	} else {
		outPkt.Flt = x.Prv + (inPkt.Flts.Lst()-x.Prv)*(flt.Flt(2)/flt.Flt(inPkt.Flts.Cnt()+1))
	}
	x.Prv = outPkt.Flt
	for _, rx := range x.Rxs {
		r = append(r, &ana.TmeFltTx{Pkt: outPkt, Rx: rx})
	}
	x.mu.Unlock()
	return r
}
func (x *StmBse) Bse() *StmBse { return x }
func (x *StmBse) Sub(rx ana.TmeFltRx, id uint32, slot ...uint32) {
	var uSlot uint32
	if len(slot) > 0 {
		uSlot = slot[0]
	}
	x.mu.Lock()
	x.Rxs[sys.Uint64(id, uSlot)] = rx
	x.mu.Unlock()
}
func (x *StmBse) UnaPos() Stm {
	r := &StmUnaPos{}
	r.Slf = r
	r.Stm = x.Slf
	r.Id = sys.NextID()
	r.Rxs = make(ana.TmeFltRxs)
	x.Sub(r.Rx, r.Id)
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
func (x *StmUnaPos) DstToInstr() int { return x.Stm.DstToInstr() + 1 }
func (x *StmUnaPos) Unsub(id uint32, slot ...uint32) {
	var uSlot uint32
	if len(slot) > 0 {
		uSlot = slot[0]
	}
	x.mu.Lock()
	delete(x.Rxs, sys.Uint64(id, uSlot))
	if len(x.Rxs) == 0 {
		x.Stm.Unsub(x.Id)
	}
	x.mu.Unlock()
}
func (x *StmUnaPos) Rx(inPkt ana.TmeFlt) (r []sys.Act) {
	x.mu.Lock()
	if ana.Cfg.Trc.IsRltStm() {
		sys.Logf("rlt.StmUnaPos(%v).Rx %p inPkt %v", x.Prm(), x, inPkt)
	}
	outPkt := ana.TmeFlt{Tme: inPkt.Tme, Flt: inPkt.Flt.Pos()}
	for _, rx := range x.Rxs {
		r = append(r, &ana.TmeFltTx{Pkt: outPkt, Rx: rx})
	}
	x.mu.Unlock()
	return r
}
func (x *StmBse) UnaNeg() Stm {
	r := &StmUnaNeg{}
	r.Slf = r
	r.Stm = x.Slf
	r.Id = sys.NextID()
	r.Rxs = make(ana.TmeFltRxs)
	x.Sub(r.Rx, r.Id)
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
func (x *StmUnaNeg) DstToInstr() int { return x.Stm.DstToInstr() + 1 }
func (x *StmUnaNeg) Unsub(id uint32, slot ...uint32) {
	var uSlot uint32
	if len(slot) > 0 {
		uSlot = slot[0]
	}
	x.mu.Lock()
	delete(x.Rxs, sys.Uint64(id, uSlot))
	if len(x.Rxs) == 0 {
		x.Stm.Unsub(x.Id)
	}
	x.mu.Unlock()
}
func (x *StmUnaNeg) Rx(inPkt ana.TmeFlt) (r []sys.Act) {
	x.mu.Lock()
	if ana.Cfg.Trc.IsRltStm() {
		sys.Logf("rlt.StmUnaNeg(%v).Rx %p inPkt %v", x.Prm(), x, inPkt)
	}
	outPkt := ana.TmeFlt{Tme: inPkt.Tme, Flt: inPkt.Flt.Neg()}
	for _, rx := range x.Rxs {
		r = append(r, &ana.TmeFltTx{Pkt: outPkt, Rx: rx})
	}
	x.mu.Unlock()
	return r
}
func (x *StmBse) UnaInv() Stm {
	r := &StmUnaInv{}
	r.Slf = r
	r.Stm = x.Slf
	r.Id = sys.NextID()
	r.Rxs = make(ana.TmeFltRxs)
	x.Sub(r.Rx, r.Id)
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
func (x *StmUnaInv) DstToInstr() int { return x.Stm.DstToInstr() + 1 }
func (x *StmUnaInv) Unsub(id uint32, slot ...uint32) {
	var uSlot uint32
	if len(slot) > 0 {
		uSlot = slot[0]
	}
	x.mu.Lock()
	delete(x.Rxs, sys.Uint64(id, uSlot))
	if len(x.Rxs) == 0 {
		x.Stm.Unsub(x.Id)
	}
	x.mu.Unlock()
}
func (x *StmUnaInv) Rx(inPkt ana.TmeFlt) (r []sys.Act) {
	x.mu.Lock()
	if ana.Cfg.Trc.IsRltStm() {
		sys.Logf("rlt.StmUnaInv(%v).Rx %p inPkt %v", x.Prm(), x, inPkt)
	}
	outPkt := ana.TmeFlt{Tme: inPkt.Tme, Flt: inPkt.Flt.Inv()}
	for _, rx := range x.Rxs {
		r = append(r, &ana.TmeFltTx{Pkt: outPkt, Rx: rx})
	}
	x.mu.Unlock()
	return r
}
func (x *StmBse) UnaSqr() Stm {
	r := &StmUnaSqr{}
	r.Slf = r
	r.Stm = x.Slf
	r.Id = sys.NextID()
	r.Rxs = make(ana.TmeFltRxs)
	x.Sub(r.Rx, r.Id)
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
func (x *StmUnaSqr) DstToInstr() int { return x.Stm.DstToInstr() + 1 }
func (x *StmUnaSqr) Unsub(id uint32, slot ...uint32) {
	var uSlot uint32
	if len(slot) > 0 {
		uSlot = slot[0]
	}
	x.mu.Lock()
	delete(x.Rxs, sys.Uint64(id, uSlot))
	if len(x.Rxs) == 0 {
		x.Stm.Unsub(x.Id)
	}
	x.mu.Unlock()
}
func (x *StmUnaSqr) Rx(inPkt ana.TmeFlt) (r []sys.Act) {
	x.mu.Lock()
	if ana.Cfg.Trc.IsRltStm() {
		sys.Logf("rlt.StmUnaSqr(%v).Rx %p inPkt %v", x.Prm(), x, inPkt)
	}
	outPkt := ana.TmeFlt{Tme: inPkt.Tme, Flt: inPkt.Flt.Sqr()}
	for _, rx := range x.Rxs {
		r = append(r, &ana.TmeFltTx{Pkt: outPkt, Rx: rx})
	}
	x.mu.Unlock()
	return r
}
func (x *StmBse) UnaSqrt() Stm {
	r := &StmUnaSqrt{}
	r.Slf = r
	r.Stm = x.Slf
	r.Id = sys.NextID()
	r.Rxs = make(ana.TmeFltRxs)
	x.Sub(r.Rx, r.Id)
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
func (x *StmUnaSqrt) DstToInstr() int { return x.Stm.DstToInstr() + 1 }
func (x *StmUnaSqrt) Unsub(id uint32, slot ...uint32) {
	var uSlot uint32
	if len(slot) > 0 {
		uSlot = slot[0]
	}
	x.mu.Lock()
	delete(x.Rxs, sys.Uint64(id, uSlot))
	if len(x.Rxs) == 0 {
		x.Stm.Unsub(x.Id)
	}
	x.mu.Unlock()
}
func (x *StmUnaSqrt) Rx(inPkt ana.TmeFlt) (r []sys.Act) {
	x.mu.Lock()
	if ana.Cfg.Trc.IsRltStm() {
		sys.Logf("rlt.StmUnaSqrt(%v).Rx %p inPkt %v", x.Prm(), x, inPkt)
	}
	outPkt := ana.TmeFlt{Tme: inPkt.Tme, Flt: inPkt.Flt.Sqrt()}
	for _, rx := range x.Rxs {
		r = append(r, &ana.TmeFltTx{Pkt: outPkt, Rx: rx})
	}
	x.mu.Unlock()
	return r
}
func (x *StmBse) SclAdd(scl flt.Flt) Stm {
	r := &StmSclAdd{}
	r.Slf = r
	r.Stm = x.Slf
	r.Scl = scl
	r.Id = sys.NextID()
	r.Rxs = make(ana.TmeFltRxs)
	x.Sub(r.Rx, r.Id)
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
func (x *StmSclAdd) DstToInstr() int { return x.Stm.DstToInstr() + 1 }
func (x *StmSclAdd) Unsub(id uint32, slot ...uint32) {
	var uSlot uint32
	if len(slot) > 0 {
		uSlot = slot[0]
	}
	x.mu.Lock()
	delete(x.Rxs, sys.Uint64(id, uSlot))
	if len(x.Rxs) == 0 {
		x.Stm.Unsub(x.Id)
	}
	x.mu.Unlock()
}
func (x *StmSclAdd) Rx(inPkt ana.TmeFlt) (r []sys.Act) {
	x.mu.Lock()
	if ana.Cfg.Trc.IsRltStm() {
		sys.Logf("rlt.StmSclAdd(%v).Rx %p inPkt %v", x.Prm(), x, inPkt)
	}
	outPkt := ana.TmeFlt{Tme: inPkt.Tme, Flt: inPkt.Flt.Add(x.Scl)}
	for _, rx := range x.Rxs {
		r = append(r, &ana.TmeFltTx{Pkt: outPkt, Rx: rx})
	}
	x.mu.Unlock()
	return r
}
func (x *StmBse) SclSub(scl flt.Flt) Stm {
	r := &StmSclSub{}
	r.Slf = r
	r.Stm = x.Slf
	r.Scl = scl
	r.Id = sys.NextID()
	r.Rxs = make(ana.TmeFltRxs)
	x.Sub(r.Rx, r.Id)
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
func (x *StmSclSub) DstToInstr() int { return x.Stm.DstToInstr() + 1 }
func (x *StmSclSub) Unsub(id uint32, slot ...uint32) {
	var uSlot uint32
	if len(slot) > 0 {
		uSlot = slot[0]
	}
	x.mu.Lock()
	delete(x.Rxs, sys.Uint64(id, uSlot))
	if len(x.Rxs) == 0 {
		x.Stm.Unsub(x.Id)
	}
	x.mu.Unlock()
}
func (x *StmSclSub) Rx(inPkt ana.TmeFlt) (r []sys.Act) {
	x.mu.Lock()
	if ana.Cfg.Trc.IsRltStm() {
		sys.Logf("rlt.StmSclSub(%v).Rx %p inPkt %v", x.Prm(), x, inPkt)
	}
	outPkt := ana.TmeFlt{Tme: inPkt.Tme, Flt: inPkt.Flt.Sub(x.Scl)}
	for _, rx := range x.Rxs {
		r = append(r, &ana.TmeFltTx{Pkt: outPkt, Rx: rx})
	}
	x.mu.Unlock()
	return r
}
func (x *StmBse) SclMul(scl flt.Flt) Stm {
	r := &StmSclMul{}
	r.Slf = r
	r.Stm = x.Slf
	r.Scl = scl
	r.Id = sys.NextID()
	r.Rxs = make(ana.TmeFltRxs)
	x.Sub(r.Rx, r.Id)
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
func (x *StmSclMul) DstToInstr() int { return x.Stm.DstToInstr() + 1 }
func (x *StmSclMul) Unsub(id uint32, slot ...uint32) {
	var uSlot uint32
	if len(slot) > 0 {
		uSlot = slot[0]
	}
	x.mu.Lock()
	delete(x.Rxs, sys.Uint64(id, uSlot))
	if len(x.Rxs) == 0 {
		x.Stm.Unsub(x.Id)
	}
	x.mu.Unlock()
}
func (x *StmSclMul) Rx(inPkt ana.TmeFlt) (r []sys.Act) {
	x.mu.Lock()
	if ana.Cfg.Trc.IsRltStm() {
		sys.Logf("rlt.StmSclMul(%v).Rx %p inPkt %v", x.Prm(), x, inPkt)
	}
	outPkt := ana.TmeFlt{Tme: inPkt.Tme, Flt: inPkt.Flt.Mul(x.Scl)}
	for _, rx := range x.Rxs {
		r = append(r, &ana.TmeFltTx{Pkt: outPkt, Rx: rx})
	}
	x.mu.Unlock()
	return r
}
func (x *StmBse) SclDiv(scl flt.Flt) Stm {
	r := &StmSclDiv{}
	r.Slf = r
	r.Stm = x.Slf
	r.Scl = scl
	r.Id = sys.NextID()
	r.Rxs = make(ana.TmeFltRxs)
	x.Sub(r.Rx, r.Id)
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
func (x *StmSclDiv) DstToInstr() int { return x.Stm.DstToInstr() + 1 }
func (x *StmSclDiv) Unsub(id uint32, slot ...uint32) {
	var uSlot uint32
	if len(slot) > 0 {
		uSlot = slot[0]
	}
	x.mu.Lock()
	delete(x.Rxs, sys.Uint64(id, uSlot))
	if len(x.Rxs) == 0 {
		x.Stm.Unsub(x.Id)
	}
	x.mu.Unlock()
}
func (x *StmSclDiv) Rx(inPkt ana.TmeFlt) (r []sys.Act) {
	x.mu.Lock()
	if ana.Cfg.Trc.IsRltStm() {
		sys.Logf("rlt.StmSclDiv(%v).Rx %p inPkt %v", x.Prm(), x, inPkt)
	}
	outPkt := ana.TmeFlt{Tme: inPkt.Tme, Flt: inPkt.Flt.Div(x.Scl)}
	for _, rx := range x.Rxs {
		r = append(r, &ana.TmeFltTx{Pkt: outPkt, Rx: rx})
	}
	x.mu.Unlock()
	return r
}
func (x *StmBse) SclRem(scl flt.Flt) Stm {
	r := &StmSclRem{}
	r.Slf = r
	r.Stm = x.Slf
	r.Scl = scl
	r.Id = sys.NextID()
	r.Rxs = make(ana.TmeFltRxs)
	x.Sub(r.Rx, r.Id)
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
func (x *StmSclRem) DstToInstr() int { return x.Stm.DstToInstr() + 1 }
func (x *StmSclRem) Unsub(id uint32, slot ...uint32) {
	var uSlot uint32
	if len(slot) > 0 {
		uSlot = slot[0]
	}
	x.mu.Lock()
	delete(x.Rxs, sys.Uint64(id, uSlot))
	if len(x.Rxs) == 0 {
		x.Stm.Unsub(x.Id)
	}
	x.mu.Unlock()
}
func (x *StmSclRem) Rx(inPkt ana.TmeFlt) (r []sys.Act) {
	x.mu.Lock()
	if ana.Cfg.Trc.IsRltStm() {
		sys.Logf("rlt.StmSclRem(%v).Rx %p inPkt %v", x.Prm(), x, inPkt)
	}
	outPkt := ana.TmeFlt{Tme: inPkt.Tme, Flt: inPkt.Flt.Rem(x.Scl)}
	for _, rx := range x.Rxs {
		r = append(r, &ana.TmeFltTx{Pkt: outPkt, Rx: rx})
	}
	x.mu.Unlock()
	return r
}
func (x *StmBse) SclPow(scl flt.Flt) Stm {
	r := &StmSclPow{}
	r.Slf = r
	r.Stm = x.Slf
	r.Scl = scl
	r.Id = sys.NextID()
	r.Rxs = make(ana.TmeFltRxs)
	x.Sub(r.Rx, r.Id)
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
func (x *StmSclPow) DstToInstr() int { return x.Stm.DstToInstr() + 1 }
func (x *StmSclPow) Unsub(id uint32, slot ...uint32) {
	var uSlot uint32
	if len(slot) > 0 {
		uSlot = slot[0]
	}
	x.mu.Lock()
	delete(x.Rxs, sys.Uint64(id, uSlot))
	if len(x.Rxs) == 0 {
		x.Stm.Unsub(x.Id)
	}
	x.mu.Unlock()
}
func (x *StmSclPow) Rx(inPkt ana.TmeFlt) (r []sys.Act) {
	x.mu.Lock()
	if ana.Cfg.Trc.IsRltStm() {
		sys.Logf("rlt.StmSclPow(%v).Rx %p inPkt %v", x.Prm(), x, inPkt)
	}
	outPkt := ana.TmeFlt{Tme: inPkt.Tme, Flt: inPkt.Flt.Pow(x.Scl)}
	for _, rx := range x.Rxs {
		r = append(r, &ana.TmeFltTx{Pkt: outPkt, Rx: rx})
	}
	x.mu.Unlock()
	return r
}
func (x *StmBse) SclMin(scl flt.Flt) Stm {
	r := &StmSclMin{}
	r.Slf = r
	r.Stm = x.Slf
	r.Scl = scl
	r.Id = sys.NextID()
	r.Rxs = make(ana.TmeFltRxs)
	x.Sub(r.Rx, r.Id)
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
func (x *StmSclMin) DstToInstr() int { return x.Stm.DstToInstr() + 1 }
func (x *StmSclMin) Unsub(id uint32, slot ...uint32) {
	var uSlot uint32
	if len(slot) > 0 {
		uSlot = slot[0]
	}
	x.mu.Lock()
	delete(x.Rxs, sys.Uint64(id, uSlot))
	if len(x.Rxs) == 0 {
		x.Stm.Unsub(x.Id)
	}
	x.mu.Unlock()
}
func (x *StmSclMin) Rx(inPkt ana.TmeFlt) (r []sys.Act) {
	x.mu.Lock()
	if ana.Cfg.Trc.IsRltStm() {
		sys.Logf("rlt.StmSclMin(%v).Rx %p inPkt %v", x.Prm(), x, inPkt)
	}
	outPkt := ana.TmeFlt{Tme: inPkt.Tme, Flt: inPkt.Flt.Min(x.Scl)}
	for _, rx := range x.Rxs {
		r = append(r, &ana.TmeFltTx{Pkt: outPkt, Rx: rx})
	}
	x.mu.Unlock()
	return r
}
func (x *StmBse) SclMax(scl flt.Flt) Stm {
	r := &StmSclMax{}
	r.Slf = r
	r.Stm = x.Slf
	r.Scl = scl
	r.Id = sys.NextID()
	r.Rxs = make(ana.TmeFltRxs)
	x.Sub(r.Rx, r.Id)
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
func (x *StmSclMax) DstToInstr() int { return x.Stm.DstToInstr() + 1 }
func (x *StmSclMax) Unsub(id uint32, slot ...uint32) {
	var uSlot uint32
	if len(slot) > 0 {
		uSlot = slot[0]
	}
	x.mu.Lock()
	delete(x.Rxs, sys.Uint64(id, uSlot))
	if len(x.Rxs) == 0 {
		x.Stm.Unsub(x.Id)
	}
	x.mu.Unlock()
}
func (x *StmSclMax) Rx(inPkt ana.TmeFlt) (r []sys.Act) {
	x.mu.Lock()
	if ana.Cfg.Trc.IsRltStm() {
		sys.Logf("rlt.StmSclMax(%v).Rx %p inPkt %v", x.Prm(), x, inPkt)
	}
	outPkt := ana.TmeFlt{Tme: inPkt.Tme, Flt: inPkt.Flt.Max(x.Scl)}
	for _, rx := range x.Rxs {
		r = append(r, &ana.TmeFltTx{Pkt: outPkt, Rx: rx})
	}
	x.mu.Unlock()
	return r
}
func (x *StmBse) SelEql(sel flt.Flt) Stm {
	r := &StmSelEql{}
	r.Slf = r
	r.Stm = x.Slf
	r.Sel = sel
	r.Id = sys.NextID()
	r.Rxs = make(ana.TmeFltRxs)
	x.Sub(r.Rx, r.Id)
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
func (x *StmSelEql) DstToInstr() int { return x.Stm.DstToInstr() + 1 }
func (x *StmSelEql) Unsub(id uint32, slot ...uint32) {
	var uSlot uint32
	if len(slot) > 0 {
		uSlot = slot[0]
	}
	x.mu.Lock()
	delete(x.Rxs, sys.Uint64(id, uSlot))
	if len(x.Rxs) == 0 {
		x.Stm.Unsub(x.Id)
	}
	x.mu.Unlock()
}
func (x *StmSelEql) Rx(inPkt ana.TmeFlt) (r []sys.Act) {
	x.mu.Lock()
	if ana.Cfg.Trc.IsRltStm() {
		sys.Logf("rlt.StmSelEql(%v).Rx %p inPkt %v", x.Prm(), x, inPkt)
	}
	outPkt := ana.TmeFlt{Tme: inPkt.Tme, Flt: inPkt.Flt.SelEql(x.Sel)}
	for _, rx := range x.Rxs {
		r = append(r, &ana.TmeFltTx{Pkt: outPkt, Rx: rx})
	}
	x.mu.Unlock()
	return r
}
func (x *StmBse) SelNeq(sel flt.Flt) Stm {
	r := &StmSelNeq{}
	r.Slf = r
	r.Stm = x.Slf
	r.Sel = sel
	r.Id = sys.NextID()
	r.Rxs = make(ana.TmeFltRxs)
	x.Sub(r.Rx, r.Id)
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
func (x *StmSelNeq) DstToInstr() int { return x.Stm.DstToInstr() + 1 }
func (x *StmSelNeq) Unsub(id uint32, slot ...uint32) {
	var uSlot uint32
	if len(slot) > 0 {
		uSlot = slot[0]
	}
	x.mu.Lock()
	delete(x.Rxs, sys.Uint64(id, uSlot))
	if len(x.Rxs) == 0 {
		x.Stm.Unsub(x.Id)
	}
	x.mu.Unlock()
}
func (x *StmSelNeq) Rx(inPkt ana.TmeFlt) (r []sys.Act) {
	x.mu.Lock()
	if ana.Cfg.Trc.IsRltStm() {
		sys.Logf("rlt.StmSelNeq(%v).Rx %p inPkt %v", x.Prm(), x, inPkt)
	}
	outPkt := ana.TmeFlt{Tme: inPkt.Tme, Flt: inPkt.Flt.SelNeq(x.Sel)}
	for _, rx := range x.Rxs {
		r = append(r, &ana.TmeFltTx{Pkt: outPkt, Rx: rx})
	}
	x.mu.Unlock()
	return r
}
func (x *StmBse) SelLss(sel flt.Flt) Stm {
	r := &StmSelLss{}
	r.Slf = r
	r.Stm = x.Slf
	r.Sel = sel
	r.Id = sys.NextID()
	r.Rxs = make(ana.TmeFltRxs)
	x.Sub(r.Rx, r.Id)
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
func (x *StmSelLss) DstToInstr() int { return x.Stm.DstToInstr() + 1 }
func (x *StmSelLss) Unsub(id uint32, slot ...uint32) {
	var uSlot uint32
	if len(slot) > 0 {
		uSlot = slot[0]
	}
	x.mu.Lock()
	delete(x.Rxs, sys.Uint64(id, uSlot))
	if len(x.Rxs) == 0 {
		x.Stm.Unsub(x.Id)
	}
	x.mu.Unlock()
}
func (x *StmSelLss) Rx(inPkt ana.TmeFlt) (r []sys.Act) {
	x.mu.Lock()
	if ana.Cfg.Trc.IsRltStm() {
		sys.Logf("rlt.StmSelLss(%v).Rx %p inPkt %v", x.Prm(), x, inPkt)
	}
	outPkt := ana.TmeFlt{Tme: inPkt.Tme, Flt: inPkt.Flt.SelLss(x.Sel)}
	for _, rx := range x.Rxs {
		r = append(r, &ana.TmeFltTx{Pkt: outPkt, Rx: rx})
	}
	x.mu.Unlock()
	return r
}
func (x *StmBse) SelGtr(sel flt.Flt) Stm {
	r := &StmSelGtr{}
	r.Slf = r
	r.Stm = x.Slf
	r.Sel = sel
	r.Id = sys.NextID()
	r.Rxs = make(ana.TmeFltRxs)
	x.Sub(r.Rx, r.Id)
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
func (x *StmSelGtr) DstToInstr() int { return x.Stm.DstToInstr() + 1 }
func (x *StmSelGtr) Unsub(id uint32, slot ...uint32) {
	var uSlot uint32
	if len(slot) > 0 {
		uSlot = slot[0]
	}
	x.mu.Lock()
	delete(x.Rxs, sys.Uint64(id, uSlot))
	if len(x.Rxs) == 0 {
		x.Stm.Unsub(x.Id)
	}
	x.mu.Unlock()
}
func (x *StmSelGtr) Rx(inPkt ana.TmeFlt) (r []sys.Act) {
	x.mu.Lock()
	if ana.Cfg.Trc.IsRltStm() {
		sys.Logf("rlt.StmSelGtr(%v).Rx %p inPkt %v", x.Prm(), x, inPkt)
	}
	outPkt := ana.TmeFlt{Tme: inPkt.Tme, Flt: inPkt.Flt.SelGtr(x.Sel)}
	for _, rx := range x.Rxs {
		r = append(r, &ana.TmeFltTx{Pkt: outPkt, Rx: rx})
	}
	x.mu.Unlock()
	return r
}
func (x *StmBse) SelLeq(sel flt.Flt) Stm {
	r := &StmSelLeq{}
	r.Slf = r
	r.Stm = x.Slf
	r.Sel = sel
	r.Id = sys.NextID()
	r.Rxs = make(ana.TmeFltRxs)
	x.Sub(r.Rx, r.Id)
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
func (x *StmSelLeq) DstToInstr() int { return x.Stm.DstToInstr() + 1 }
func (x *StmSelLeq) Unsub(id uint32, slot ...uint32) {
	var uSlot uint32
	if len(slot) > 0 {
		uSlot = slot[0]
	}
	x.mu.Lock()
	delete(x.Rxs, sys.Uint64(id, uSlot))
	if len(x.Rxs) == 0 {
		x.Stm.Unsub(x.Id)
	}
	x.mu.Unlock()
}
func (x *StmSelLeq) Rx(inPkt ana.TmeFlt) (r []sys.Act) {
	x.mu.Lock()
	if ana.Cfg.Trc.IsRltStm() {
		sys.Logf("rlt.StmSelLeq(%v).Rx %p inPkt %v", x.Prm(), x, inPkt)
	}
	outPkt := ana.TmeFlt{Tme: inPkt.Tme, Flt: inPkt.Flt.SelLeq(x.Sel)}
	for _, rx := range x.Rxs {
		r = append(r, &ana.TmeFltTx{Pkt: outPkt, Rx: rx})
	}
	x.mu.Unlock()
	return r
}
func (x *StmBse) SelGeq(sel flt.Flt) Stm {
	r := &StmSelGeq{}
	r.Slf = r
	r.Stm = x.Slf
	r.Sel = sel
	r.Id = sys.NextID()
	r.Rxs = make(ana.TmeFltRxs)
	x.Sub(r.Rx, r.Id)
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
func (x *StmSelGeq) DstToInstr() int { return x.Stm.DstToInstr() + 1 }
func (x *StmSelGeq) Unsub(id uint32, slot ...uint32) {
	var uSlot uint32
	if len(slot) > 0 {
		uSlot = slot[0]
	}
	x.mu.Lock()
	delete(x.Rxs, sys.Uint64(id, uSlot))
	if len(x.Rxs) == 0 {
		x.Stm.Unsub(x.Id)
	}
	x.mu.Unlock()
}
func (x *StmSelGeq) Rx(inPkt ana.TmeFlt) (r []sys.Act) {
	x.mu.Lock()
	if ana.Cfg.Trc.IsRltStm() {
		sys.Logf("rlt.StmSelGeq(%v).Rx %p inPkt %v", x.Prm(), x, inPkt)
	}
	outPkt := ana.TmeFlt{Tme: inPkt.Tme, Flt: inPkt.Flt.SelGeq(x.Sel)}
	for _, rx := range x.Rxs {
		r = append(r, &ana.TmeFltTx{Pkt: outPkt, Rx: rx})
	}
	x.mu.Unlock()
	return r
}
func (x *StmBse) AggFst(length unt.Unt) Stm {
	r := &StmAggFst{}
	r.Slf = r
	r.Stm = x.Slf
	r.Length = length
	r.Id = sys.NextID()
	r.Rxs = make(ana.TmeFltRxs)
	r.Tmes = tmes.New()
	r.Vals = flts.New()
	x.Sub(r.Rx, r.Id)
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
func (x *StmAggFst) DstToInstr() int { return x.Stm.DstToInstr() + 1 }
func (x *StmAggFst) Unsub(id uint32, slot ...uint32) {
	var uSlot uint32
	if len(slot) > 0 {
		uSlot = slot[0]
	}
	x.mu.Lock()
	delete(x.Rxs, sys.Uint64(id, uSlot))
	if len(x.Rxs) == 0 {
		x.Stm.Unsub(x.Id)
	}
	x.mu.Unlock()
}
func (x *StmAggFst) Rx(inPkt ana.TmeFlt) (r []sys.Act) {
	if x.Length > 0 {
		x.mu.Lock()
		if ana.Cfg.Trc.IsRltStm() {
			sys.Logf("rlt.StmAggFst(%v).Rx %p inPkt %v", x.Prm(), x, inPkt)
		}
		x.Tmes.Que(inPkt.Tme)
		x.Vals.Que(inPkt.Flt)
		if x.Tmes.Cnt() >= x.Length {
			outPkt := ana.TmeFlt{
				Tme: x.Tmes.At(x.Length - 1),
				Flt: x.Vals.To(x.Length).Fst(),
			}
			x.Tmes.Dque()
			x.Vals.Dque()
			for _, rx := range x.Rxs {
				r = append(r, &ana.TmeFltTx{Pkt: outPkt, Rx: rx})
			}
		}
		x.mu.Unlock()
	}
	return r
}
func (x *StmBse) AggLst(length unt.Unt) Stm {
	r := &StmAggLst{}
	r.Slf = r
	r.Stm = x.Slf
	r.Length = length
	r.Id = sys.NextID()
	r.Rxs = make(ana.TmeFltRxs)
	r.Tmes = tmes.New()
	r.Vals = flts.New()
	x.Sub(r.Rx, r.Id)
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
func (x *StmAggLst) DstToInstr() int { return x.Stm.DstToInstr() + 1 }
func (x *StmAggLst) Unsub(id uint32, slot ...uint32) {
	var uSlot uint32
	if len(slot) > 0 {
		uSlot = slot[0]
	}
	x.mu.Lock()
	delete(x.Rxs, sys.Uint64(id, uSlot))
	if len(x.Rxs) == 0 {
		x.Stm.Unsub(x.Id)
	}
	x.mu.Unlock()
}
func (x *StmAggLst) Rx(inPkt ana.TmeFlt) (r []sys.Act) {
	if x.Length > 0 {
		x.mu.Lock()
		if ana.Cfg.Trc.IsRltStm() {
			sys.Logf("rlt.StmAggLst(%v).Rx %p inPkt %v", x.Prm(), x, inPkt)
		}
		x.Tmes.Que(inPkt.Tme)
		x.Vals.Que(inPkt.Flt)
		if x.Tmes.Cnt() >= x.Length {
			outPkt := ana.TmeFlt{
				Tme: x.Tmes.At(x.Length - 1),
				Flt: x.Vals.To(x.Length).Lst(),
			}
			x.Tmes.Dque()
			x.Vals.Dque()
			for _, rx := range x.Rxs {
				r = append(r, &ana.TmeFltTx{Pkt: outPkt, Rx: rx})
			}
		}
		x.mu.Unlock()
	}
	return r
}
func (x *StmBse) AggSum(length unt.Unt) Stm {
	r := &StmAggSum{}
	r.Slf = r
	r.Stm = x.Slf
	r.Length = length
	r.Id = sys.NextID()
	r.Rxs = make(ana.TmeFltRxs)
	r.Tmes = tmes.New()
	r.Vals = flts.New()
	x.Sub(r.Rx, r.Id)
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
func (x *StmAggSum) DstToInstr() int { return x.Stm.DstToInstr() + 1 }
func (x *StmAggSum) Unsub(id uint32, slot ...uint32) {
	var uSlot uint32
	if len(slot) > 0 {
		uSlot = slot[0]
	}
	x.mu.Lock()
	delete(x.Rxs, sys.Uint64(id, uSlot))
	if len(x.Rxs) == 0 {
		x.Stm.Unsub(x.Id)
	}
	x.mu.Unlock()
}
func (x *StmAggSum) Rx(inPkt ana.TmeFlt) (r []sys.Act) {
	if x.Length > 0 {
		x.mu.Lock()
		if ana.Cfg.Trc.IsRltStm() {
			sys.Logf("rlt.StmAggSum(%v).Rx %p inPkt %v", x.Prm(), x, inPkt)
		}
		x.Tmes.Que(inPkt.Tme)
		x.Vals.Que(inPkt.Flt)
		if x.Tmes.Cnt() >= x.Length {
			outPkt := ana.TmeFlt{
				Tme: x.Tmes.At(x.Length - 1),
				Flt: x.Vals.To(x.Length).Sum(),
			}
			x.Tmes.Dque()
			x.Vals.Dque()
			for _, rx := range x.Rxs {
				r = append(r, &ana.TmeFltTx{Pkt: outPkt, Rx: rx})
			}
		}
		x.mu.Unlock()
	}
	return r
}
func (x *StmBse) AggPrd(length unt.Unt) Stm {
	r := &StmAggPrd{}
	r.Slf = r
	r.Stm = x.Slf
	r.Length = length
	r.Id = sys.NextID()
	r.Rxs = make(ana.TmeFltRxs)
	r.Tmes = tmes.New()
	r.Vals = flts.New()
	x.Sub(r.Rx, r.Id)
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
func (x *StmAggPrd) DstToInstr() int { return x.Stm.DstToInstr() + 1 }
func (x *StmAggPrd) Unsub(id uint32, slot ...uint32) {
	var uSlot uint32
	if len(slot) > 0 {
		uSlot = slot[0]
	}
	x.mu.Lock()
	delete(x.Rxs, sys.Uint64(id, uSlot))
	if len(x.Rxs) == 0 {
		x.Stm.Unsub(x.Id)
	}
	x.mu.Unlock()
}
func (x *StmAggPrd) Rx(inPkt ana.TmeFlt) (r []sys.Act) {
	if x.Length > 0 {
		x.mu.Lock()
		if ana.Cfg.Trc.IsRltStm() {
			sys.Logf("rlt.StmAggPrd(%v).Rx %p inPkt %v", x.Prm(), x, inPkt)
		}
		x.Tmes.Que(inPkt.Tme)
		x.Vals.Que(inPkt.Flt)
		if x.Tmes.Cnt() >= x.Length {
			outPkt := ana.TmeFlt{
				Tme: x.Tmes.At(x.Length - 1),
				Flt: x.Vals.To(x.Length).Prd(),
			}
			x.Tmes.Dque()
			x.Vals.Dque()
			for _, rx := range x.Rxs {
				r = append(r, &ana.TmeFltTx{Pkt: outPkt, Rx: rx})
			}
		}
		x.mu.Unlock()
	}
	return r
}
func (x *StmBse) AggMin(length unt.Unt) Stm {
	r := &StmAggMin{}
	r.Slf = r
	r.Stm = x.Slf
	r.Length = length
	r.Id = sys.NextID()
	r.Rxs = make(ana.TmeFltRxs)
	r.Tmes = tmes.New()
	r.Vals = flts.New()
	x.Sub(r.Rx, r.Id)
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
func (x *StmAggMin) DstToInstr() int { return x.Stm.DstToInstr() + 1 }
func (x *StmAggMin) Unsub(id uint32, slot ...uint32) {
	var uSlot uint32
	if len(slot) > 0 {
		uSlot = slot[0]
	}
	x.mu.Lock()
	delete(x.Rxs, sys.Uint64(id, uSlot))
	if len(x.Rxs) == 0 {
		x.Stm.Unsub(x.Id)
	}
	x.mu.Unlock()
}
func (x *StmAggMin) Rx(inPkt ana.TmeFlt) (r []sys.Act) {
	if x.Length > 0 {
		x.mu.Lock()
		if ana.Cfg.Trc.IsRltStm() {
			sys.Logf("rlt.StmAggMin(%v).Rx %p inPkt %v", x.Prm(), x, inPkt)
		}
		x.Tmes.Que(inPkt.Tme)
		x.Vals.Que(inPkt.Flt)
		if x.Tmes.Cnt() >= x.Length {
			outPkt := ana.TmeFlt{
				Tme: x.Tmes.At(x.Length - 1),
				Flt: x.Vals.To(x.Length).Min(),
			}
			x.Tmes.Dque()
			x.Vals.Dque()
			for _, rx := range x.Rxs {
				r = append(r, &ana.TmeFltTx{Pkt: outPkt, Rx: rx})
			}
		}
		x.mu.Unlock()
	}
	return r
}
func (x *StmBse) AggMax(length unt.Unt) Stm {
	r := &StmAggMax{}
	r.Slf = r
	r.Stm = x.Slf
	r.Length = length
	r.Id = sys.NextID()
	r.Rxs = make(ana.TmeFltRxs)
	r.Tmes = tmes.New()
	r.Vals = flts.New()
	x.Sub(r.Rx, r.Id)
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
func (x *StmAggMax) DstToInstr() int { return x.Stm.DstToInstr() + 1 }
func (x *StmAggMax) Unsub(id uint32, slot ...uint32) {
	var uSlot uint32
	if len(slot) > 0 {
		uSlot = slot[0]
	}
	x.mu.Lock()
	delete(x.Rxs, sys.Uint64(id, uSlot))
	if len(x.Rxs) == 0 {
		x.Stm.Unsub(x.Id)
	}
	x.mu.Unlock()
}
func (x *StmAggMax) Rx(inPkt ana.TmeFlt) (r []sys.Act) {
	if x.Length > 0 {
		x.mu.Lock()
		if ana.Cfg.Trc.IsRltStm() {
			sys.Logf("rlt.StmAggMax(%v).Rx %p inPkt %v", x.Prm(), x, inPkt)
		}
		x.Tmes.Que(inPkt.Tme)
		x.Vals.Que(inPkt.Flt)
		if x.Tmes.Cnt() >= x.Length {
			outPkt := ana.TmeFlt{
				Tme: x.Tmes.At(x.Length - 1),
				Flt: x.Vals.To(x.Length).Max(),
			}
			x.Tmes.Dque()
			x.Vals.Dque()
			for _, rx := range x.Rxs {
				r = append(r, &ana.TmeFltTx{Pkt: outPkt, Rx: rx})
			}
		}
		x.mu.Unlock()
	}
	return r
}
func (x *StmBse) AggMid(length unt.Unt) Stm {
	r := &StmAggMid{}
	r.Slf = r
	r.Stm = x.Slf
	r.Length = length
	r.Id = sys.NextID()
	r.Rxs = make(ana.TmeFltRxs)
	r.Tmes = tmes.New()
	r.Vals = flts.New()
	x.Sub(r.Rx, r.Id)
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
func (x *StmAggMid) DstToInstr() int { return x.Stm.DstToInstr() + 1 }
func (x *StmAggMid) Unsub(id uint32, slot ...uint32) {
	var uSlot uint32
	if len(slot) > 0 {
		uSlot = slot[0]
	}
	x.mu.Lock()
	delete(x.Rxs, sys.Uint64(id, uSlot))
	if len(x.Rxs) == 0 {
		x.Stm.Unsub(x.Id)
	}
	x.mu.Unlock()
}
func (x *StmAggMid) Rx(inPkt ana.TmeFlt) (r []sys.Act) {
	if x.Length > 0 {
		x.mu.Lock()
		if ana.Cfg.Trc.IsRltStm() {
			sys.Logf("rlt.StmAggMid(%v).Rx %p inPkt %v", x.Prm(), x, inPkt)
		}
		x.Tmes.Que(inPkt.Tme)
		x.Vals.Que(inPkt.Flt)
		if x.Tmes.Cnt() >= x.Length {
			outPkt := ana.TmeFlt{
				Tme: x.Tmes.At(x.Length - 1),
				Flt: x.Vals.To(x.Length).Mid(),
			}
			x.Tmes.Dque()
			x.Vals.Dque()
			for _, rx := range x.Rxs {
				r = append(r, &ana.TmeFltTx{Pkt: outPkt, Rx: rx})
			}
		}
		x.mu.Unlock()
	}
	return r
}
func (x *StmBse) AggMdn(length unt.Unt) Stm {
	r := &StmAggMdn{}
	r.Slf = r
	r.Stm = x.Slf
	r.Length = length
	r.Id = sys.NextID()
	r.Rxs = make(ana.TmeFltRxs)
	r.Tmes = tmes.New()
	r.Vals = flts.New()
	x.Sub(r.Rx, r.Id)
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
func (x *StmAggMdn) DstToInstr() int { return x.Stm.DstToInstr() + 1 }
func (x *StmAggMdn) Unsub(id uint32, slot ...uint32) {
	var uSlot uint32
	if len(slot) > 0 {
		uSlot = slot[0]
	}
	x.mu.Lock()
	delete(x.Rxs, sys.Uint64(id, uSlot))
	if len(x.Rxs) == 0 {
		x.Stm.Unsub(x.Id)
	}
	x.mu.Unlock()
}
func (x *StmAggMdn) Rx(inPkt ana.TmeFlt) (r []sys.Act) {
	if x.Length > 0 {
		x.mu.Lock()
		if ana.Cfg.Trc.IsRltStm() {
			sys.Logf("rlt.StmAggMdn(%v).Rx %p inPkt %v", x.Prm(), x, inPkt)
		}
		x.Tmes.Que(inPkt.Tme)
		x.Vals.Que(inPkt.Flt)
		if x.Tmes.Cnt() >= x.Length {
			outPkt := ana.TmeFlt{
				Tme: x.Tmes.At(x.Length - 1),
				Flt: x.Vals.To(x.Length).Mdn(),
			}
			x.Tmes.Dque()
			x.Vals.Dque()
			for _, rx := range x.Rxs {
				r = append(r, &ana.TmeFltTx{Pkt: outPkt, Rx: rx})
			}
		}
		x.mu.Unlock()
	}
	return r
}
func (x *StmBse) AggSma(length unt.Unt) Stm {
	r := &StmAggSma{}
	r.Slf = r
	r.Stm = x.Slf
	r.Length = length
	r.Id = sys.NextID()
	r.Rxs = make(ana.TmeFltRxs)
	r.Tmes = tmes.New()
	r.Vals = flts.New()
	x.Sub(r.Rx, r.Id)
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
func (x *StmAggSma) DstToInstr() int { return x.Stm.DstToInstr() + 1 }
func (x *StmAggSma) Unsub(id uint32, slot ...uint32) {
	var uSlot uint32
	if len(slot) > 0 {
		uSlot = slot[0]
	}
	x.mu.Lock()
	delete(x.Rxs, sys.Uint64(id, uSlot))
	if len(x.Rxs) == 0 {
		x.Stm.Unsub(x.Id)
	}
	x.mu.Unlock()
}
func (x *StmAggSma) Rx(inPkt ana.TmeFlt) (r []sys.Act) {
	if x.Length > 0 {
		x.mu.Lock()
		if ana.Cfg.Trc.IsRltStm() {
			sys.Logf("rlt.StmAggSma(%v).Rx %p inPkt %v", x.Prm(), x, inPkt)
		}
		x.Tmes.Que(inPkt.Tme)
		x.Vals.Que(inPkt.Flt)
		if x.Tmes.Cnt() >= x.Length {
			outPkt := ana.TmeFlt{
				Tme: x.Tmes.At(x.Length - 1),
				Flt: x.Vals.To(x.Length).Sma(),
			}
			x.Tmes.Dque()
			x.Vals.Dque()
			for _, rx := range x.Rxs {
				r = append(r, &ana.TmeFltTx{Pkt: outPkt, Rx: rx})
			}
		}
		x.mu.Unlock()
	}
	return r
}
func (x *StmBse) AggGma(length unt.Unt) Stm {
	r := &StmAggGma{}
	r.Slf = r
	r.Stm = x.Slf
	r.Length = length
	r.Id = sys.NextID()
	r.Rxs = make(ana.TmeFltRxs)
	r.Tmes = tmes.New()
	r.Vals = flts.New()
	x.Sub(r.Rx, r.Id)
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
func (x *StmAggGma) DstToInstr() int { return x.Stm.DstToInstr() + 1 }
func (x *StmAggGma) Unsub(id uint32, slot ...uint32) {
	var uSlot uint32
	if len(slot) > 0 {
		uSlot = slot[0]
	}
	x.mu.Lock()
	delete(x.Rxs, sys.Uint64(id, uSlot))
	if len(x.Rxs) == 0 {
		x.Stm.Unsub(x.Id)
	}
	x.mu.Unlock()
}
func (x *StmAggGma) Rx(inPkt ana.TmeFlt) (r []sys.Act) {
	if x.Length > 0 {
		x.mu.Lock()
		if ana.Cfg.Trc.IsRltStm() {
			sys.Logf("rlt.StmAggGma(%v).Rx %p inPkt %v", x.Prm(), x, inPkt)
		}
		x.Tmes.Que(inPkt.Tme)
		x.Vals.Que(inPkt.Flt)
		if x.Tmes.Cnt() >= x.Length {
			outPkt := ana.TmeFlt{
				Tme: x.Tmes.At(x.Length - 1),
				Flt: x.Vals.To(x.Length).Gma(),
			}
			x.Tmes.Dque()
			x.Vals.Dque()
			for _, rx := range x.Rxs {
				r = append(r, &ana.TmeFltTx{Pkt: outPkt, Rx: rx})
			}
		}
		x.mu.Unlock()
	}
	return r
}
func (x *StmBse) AggWma(length unt.Unt) Stm {
	r := &StmAggWma{}
	r.Slf = r
	r.Stm = x.Slf
	r.Length = length
	r.Id = sys.NextID()
	r.Rxs = make(ana.TmeFltRxs)
	r.Tmes = tmes.New()
	r.Vals = flts.New()
	x.Sub(r.Rx, r.Id)
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
func (x *StmAggWma) DstToInstr() int { return x.Stm.DstToInstr() + 1 }
func (x *StmAggWma) Unsub(id uint32, slot ...uint32) {
	var uSlot uint32
	if len(slot) > 0 {
		uSlot = slot[0]
	}
	x.mu.Lock()
	delete(x.Rxs, sys.Uint64(id, uSlot))
	if len(x.Rxs) == 0 {
		x.Stm.Unsub(x.Id)
	}
	x.mu.Unlock()
}
func (x *StmAggWma) Rx(inPkt ana.TmeFlt) (r []sys.Act) {
	if x.Length > 0 {
		x.mu.Lock()
		if ana.Cfg.Trc.IsRltStm() {
			sys.Logf("rlt.StmAggWma(%v).Rx %p inPkt %v", x.Prm(), x, inPkt)
		}
		x.Tmes.Que(inPkt.Tme)
		x.Vals.Que(inPkt.Flt)
		if x.Tmes.Cnt() >= x.Length {
			outPkt := ana.TmeFlt{
				Tme: x.Tmes.At(x.Length - 1),
				Flt: x.Vals.To(x.Length).Wma(),
			}
			x.Tmes.Dque()
			x.Vals.Dque()
			for _, rx := range x.Rxs {
				r = append(r, &ana.TmeFltTx{Pkt: outPkt, Rx: rx})
			}
		}
		x.mu.Unlock()
	}
	return r
}
func (x *StmBse) AggRsi(length unt.Unt) Stm {
	r := &StmAggRsi{}
	r.Slf = r
	r.Stm = x.Slf
	r.Length = length
	r.Id = sys.NextID()
	r.Rxs = make(ana.TmeFltRxs)
	r.Tmes = tmes.New()
	r.Vals = flts.New()
	x.Sub(r.Rx, r.Id)
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
func (x *StmAggRsi) DstToInstr() int { return x.Stm.DstToInstr() + 1 }
func (x *StmAggRsi) Unsub(id uint32, slot ...uint32) {
	var uSlot uint32
	if len(slot) > 0 {
		uSlot = slot[0]
	}
	x.mu.Lock()
	delete(x.Rxs, sys.Uint64(id, uSlot))
	if len(x.Rxs) == 0 {
		x.Stm.Unsub(x.Id)
	}
	x.mu.Unlock()
}
func (x *StmAggRsi) Rx(inPkt ana.TmeFlt) (r []sys.Act) {
	if x.Length > 0 {
		x.mu.Lock()
		if ana.Cfg.Trc.IsRltStm() {
			sys.Logf("rlt.StmAggRsi(%v).Rx %p inPkt %v", x.Prm(), x, inPkt)
		}
		x.Tmes.Que(inPkt.Tme)
		x.Vals.Que(inPkt.Flt)
		if x.Tmes.Cnt() >= x.Length {
			outPkt := ana.TmeFlt{
				Tme: x.Tmes.At(x.Length - 1),
				Flt: x.Vals.To(x.Length).Rsi(),
			}
			x.Tmes.Dque()
			x.Vals.Dque()
			for _, rx := range x.Rxs {
				r = append(r, &ana.TmeFltTx{Pkt: outPkt, Rx: rx})
			}
		}
		x.mu.Unlock()
	}
	return r
}
func (x *StmBse) AggWrsi(length unt.Unt) Stm {
	r := &StmAggWrsi{}
	r.Slf = r
	r.Stm = x.Slf
	r.Length = length
	r.Id = sys.NextID()
	r.Rxs = make(ana.TmeFltRxs)
	r.Tmes = tmes.New()
	r.Vals = flts.New()
	x.Sub(r.Rx, r.Id)
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
func (x *StmAggWrsi) DstToInstr() int { return x.Stm.DstToInstr() + 1 }
func (x *StmAggWrsi) Unsub(id uint32, slot ...uint32) {
	var uSlot uint32
	if len(slot) > 0 {
		uSlot = slot[0]
	}
	x.mu.Lock()
	delete(x.Rxs, sys.Uint64(id, uSlot))
	if len(x.Rxs) == 0 {
		x.Stm.Unsub(x.Id)
	}
	x.mu.Unlock()
}
func (x *StmAggWrsi) Rx(inPkt ana.TmeFlt) (r []sys.Act) {
	if x.Length > 0 {
		x.mu.Lock()
		if ana.Cfg.Trc.IsRltStm() {
			sys.Logf("rlt.StmAggWrsi(%v).Rx %p inPkt %v", x.Prm(), x, inPkt)
		}
		x.Tmes.Que(inPkt.Tme)
		x.Vals.Que(inPkt.Flt)
		if x.Tmes.Cnt() >= x.Length {
			outPkt := ana.TmeFlt{
				Tme: x.Tmes.At(x.Length - 1),
				Flt: x.Vals.To(x.Length).Wrsi(),
			}
			x.Tmes.Dque()
			x.Vals.Dque()
			for _, rx := range x.Rxs {
				r = append(r, &ana.TmeFltTx{Pkt: outPkt, Rx: rx})
			}
		}
		x.mu.Unlock()
	}
	return r
}
func (x *StmBse) AggAlma(length unt.Unt) Stm {
	r := &StmAggAlma{}
	r.Slf = r
	r.Stm = x.Slf
	r.Length = length
	r.Id = sys.NextID()
	r.Rxs = make(ana.TmeFltRxs)
	r.Tmes = tmes.New()
	r.Vals = flts.New()
	x.Sub(r.Rx, r.Id)
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
func (x *StmAggAlma) DstToInstr() int { return x.Stm.DstToInstr() + 1 }
func (x *StmAggAlma) Unsub(id uint32, slot ...uint32) {
	var uSlot uint32
	if len(slot) > 0 {
		uSlot = slot[0]
	}
	x.mu.Lock()
	delete(x.Rxs, sys.Uint64(id, uSlot))
	if len(x.Rxs) == 0 {
		x.Stm.Unsub(x.Id)
	}
	x.mu.Unlock()
}
func (x *StmAggAlma) Rx(inPkt ana.TmeFlt) (r []sys.Act) {
	if x.Length > 0 {
		x.mu.Lock()
		if ana.Cfg.Trc.IsRltStm() {
			sys.Logf("rlt.StmAggAlma(%v).Rx %p inPkt %v", x.Prm(), x, inPkt)
		}
		x.Tmes.Que(inPkt.Tme)
		x.Vals.Que(inPkt.Flt)
		if x.Tmes.Cnt() >= x.Length {
			outPkt := ana.TmeFlt{
				Tme: x.Tmes.At(x.Length - 1),
				Flt: x.Vals.To(x.Length).Alma(),
			}
			x.Tmes.Dque()
			x.Vals.Dque()
			for _, rx := range x.Rxs {
				r = append(r, &ana.TmeFltTx{Pkt: outPkt, Rx: rx})
			}
		}
		x.mu.Unlock()
	}
	return r
}
func (x *StmBse) AggVrnc(length unt.Unt) Stm {
	r := &StmAggVrnc{}
	r.Slf = r
	r.Stm = x.Slf
	r.Length = length
	r.Id = sys.NextID()
	r.Rxs = make(ana.TmeFltRxs)
	r.Tmes = tmes.New()
	r.Vals = flts.New()
	x.Sub(r.Rx, r.Id)
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
func (x *StmAggVrnc) DstToInstr() int { return x.Stm.DstToInstr() + 1 }
func (x *StmAggVrnc) Unsub(id uint32, slot ...uint32) {
	var uSlot uint32
	if len(slot) > 0 {
		uSlot = slot[0]
	}
	x.mu.Lock()
	delete(x.Rxs, sys.Uint64(id, uSlot))
	if len(x.Rxs) == 0 {
		x.Stm.Unsub(x.Id)
	}
	x.mu.Unlock()
}
func (x *StmAggVrnc) Rx(inPkt ana.TmeFlt) (r []sys.Act) {
	if x.Length > 0 {
		x.mu.Lock()
		if ana.Cfg.Trc.IsRltStm() {
			sys.Logf("rlt.StmAggVrnc(%v).Rx %p inPkt %v", x.Prm(), x, inPkt)
		}
		x.Tmes.Que(inPkt.Tme)
		x.Vals.Que(inPkt.Flt)
		if x.Tmes.Cnt() >= x.Length {
			outPkt := ana.TmeFlt{
				Tme: x.Tmes.At(x.Length - 1),
				Flt: x.Vals.To(x.Length).Vrnc(),
			}
			x.Tmes.Dque()
			x.Vals.Dque()
			for _, rx := range x.Rxs {
				r = append(r, &ana.TmeFltTx{Pkt: outPkt, Rx: rx})
			}
		}
		x.mu.Unlock()
	}
	return r
}
func (x *StmBse) AggStd(length unt.Unt) Stm {
	r := &StmAggStd{}
	r.Slf = r
	r.Stm = x.Slf
	r.Length = length
	r.Id = sys.NextID()
	r.Rxs = make(ana.TmeFltRxs)
	r.Tmes = tmes.New()
	r.Vals = flts.New()
	x.Sub(r.Rx, r.Id)
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
func (x *StmAggStd) DstToInstr() int { return x.Stm.DstToInstr() + 1 }
func (x *StmAggStd) Unsub(id uint32, slot ...uint32) {
	var uSlot uint32
	if len(slot) > 0 {
		uSlot = slot[0]
	}
	x.mu.Lock()
	delete(x.Rxs, sys.Uint64(id, uSlot))
	if len(x.Rxs) == 0 {
		x.Stm.Unsub(x.Id)
	}
	x.mu.Unlock()
}
func (x *StmAggStd) Rx(inPkt ana.TmeFlt) (r []sys.Act) {
	if x.Length > 0 {
		x.mu.Lock()
		if ana.Cfg.Trc.IsRltStm() {
			sys.Logf("rlt.StmAggStd(%v).Rx %p inPkt %v", x.Prm(), x, inPkt)
		}
		x.Tmes.Que(inPkt.Tme)
		x.Vals.Que(inPkt.Flt)
		if x.Tmes.Cnt() >= x.Length {
			outPkt := ana.TmeFlt{
				Tme: x.Tmes.At(x.Length - 1),
				Flt: x.Vals.To(x.Length).Std(),
			}
			x.Tmes.Dque()
			x.Vals.Dque()
			for _, rx := range x.Rxs {
				r = append(r, &ana.TmeFltTx{Pkt: outPkt, Rx: rx})
			}
		}
		x.mu.Unlock()
	}
	return r
}
func (x *StmBse) AggRngFul(length unt.Unt) Stm {
	r := &StmAggRngFul{}
	r.Slf = r
	r.Stm = x.Slf
	r.Length = length
	r.Id = sys.NextID()
	r.Rxs = make(ana.TmeFltRxs)
	r.Tmes = tmes.New()
	r.Vals = flts.New()
	x.Sub(r.Rx, r.Id)
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
func (x *StmAggRngFul) DstToInstr() int { return x.Stm.DstToInstr() + 1 }
func (x *StmAggRngFul) Unsub(id uint32, slot ...uint32) {
	var uSlot uint32
	if len(slot) > 0 {
		uSlot = slot[0]
	}
	x.mu.Lock()
	delete(x.Rxs, sys.Uint64(id, uSlot))
	if len(x.Rxs) == 0 {
		x.Stm.Unsub(x.Id)
	}
	x.mu.Unlock()
}
func (x *StmAggRngFul) Rx(inPkt ana.TmeFlt) (r []sys.Act) {
	if x.Length > 0 {
		x.mu.Lock()
		if ana.Cfg.Trc.IsRltStm() {
			sys.Logf("rlt.StmAggRngFul(%v).Rx %p inPkt %v", x.Prm(), x, inPkt)
		}
		x.Tmes.Que(inPkt.Tme)
		x.Vals.Que(inPkt.Flt)
		if x.Tmes.Cnt() >= x.Length {
			outPkt := ana.TmeFlt{
				Tme: x.Tmes.At(x.Length - 1),
				Flt: x.Vals.To(x.Length).RngFul(),
			}
			x.Tmes.Dque()
			x.Vals.Dque()
			for _, rx := range x.Rxs {
				r = append(r, &ana.TmeFltTx{Pkt: outPkt, Rx: rx})
			}
		}
		x.mu.Unlock()
	}
	return r
}
func (x *StmBse) AggRngLst(length unt.Unt) Stm {
	r := &StmAggRngLst{}
	r.Slf = r
	r.Stm = x.Slf
	r.Length = length
	r.Id = sys.NextID()
	r.Rxs = make(ana.TmeFltRxs)
	r.Tmes = tmes.New()
	r.Vals = flts.New()
	x.Sub(r.Rx, r.Id)
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
func (x *StmAggRngLst) DstToInstr() int { return x.Stm.DstToInstr() + 1 }
func (x *StmAggRngLst) Unsub(id uint32, slot ...uint32) {
	var uSlot uint32
	if len(slot) > 0 {
		uSlot = slot[0]
	}
	x.mu.Lock()
	delete(x.Rxs, sys.Uint64(id, uSlot))
	if len(x.Rxs) == 0 {
		x.Stm.Unsub(x.Id)
	}
	x.mu.Unlock()
}
func (x *StmAggRngLst) Rx(inPkt ana.TmeFlt) (r []sys.Act) {
	if x.Length > 0 {
		x.mu.Lock()
		if ana.Cfg.Trc.IsRltStm() {
			sys.Logf("rlt.StmAggRngLst(%v).Rx %p inPkt %v", x.Prm(), x, inPkt)
		}
		x.Tmes.Que(inPkt.Tme)
		x.Vals.Que(inPkt.Flt)
		if x.Tmes.Cnt() >= x.Length {
			outPkt := ana.TmeFlt{
				Tme: x.Tmes.At(x.Length - 1),
				Flt: x.Vals.To(x.Length).RngLst(),
			}
			x.Tmes.Dque()
			x.Vals.Dque()
			for _, rx := range x.Rxs {
				r = append(r, &ana.TmeFltTx{Pkt: outPkt, Rx: rx})
			}
		}
		x.mu.Unlock()
	}
	return r
}
func (x *StmBse) AggProLst(length unt.Unt) Stm {
	r := &StmAggProLst{}
	r.Slf = r
	r.Stm = x.Slf
	r.Length = length
	r.Id = sys.NextID()
	r.Rxs = make(ana.TmeFltRxs)
	r.Tmes = tmes.New()
	r.Vals = flts.New()
	x.Sub(r.Rx, r.Id)
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
func (x *StmAggProLst) DstToInstr() int { return x.Stm.DstToInstr() + 1 }
func (x *StmAggProLst) Unsub(id uint32, slot ...uint32) {
	var uSlot uint32
	if len(slot) > 0 {
		uSlot = slot[0]
	}
	x.mu.Lock()
	delete(x.Rxs, sys.Uint64(id, uSlot))
	if len(x.Rxs) == 0 {
		x.Stm.Unsub(x.Id)
	}
	x.mu.Unlock()
}
func (x *StmAggProLst) Rx(inPkt ana.TmeFlt) (r []sys.Act) {
	if x.Length > 0 {
		x.mu.Lock()
		if ana.Cfg.Trc.IsRltStm() {
			sys.Logf("rlt.StmAggProLst(%v).Rx %p inPkt %v", x.Prm(), x, inPkt)
		}
		x.Tmes.Que(inPkt.Tme)
		x.Vals.Que(inPkt.Flt)
		if x.Tmes.Cnt() >= x.Length {
			outPkt := ana.TmeFlt{
				Tme: x.Tmes.At(x.Length - 1),
				Flt: x.Vals.To(x.Length).ProLst(),
			}
			x.Tmes.Dque()
			x.Vals.Dque()
			for _, rx := range x.Rxs {
				r = append(r, &ana.TmeFltTx{Pkt: outPkt, Rx: rx})
			}
		}
		x.mu.Unlock()
	}
	return r
}
func (x *StmBse) AggProSma(length unt.Unt) Stm {
	r := &StmAggProSma{}
	r.Slf = r
	r.Stm = x.Slf
	r.Length = length
	r.Id = sys.NextID()
	r.Rxs = make(ana.TmeFltRxs)
	r.Tmes = tmes.New()
	r.Vals = flts.New()
	x.Sub(r.Rx, r.Id)
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
func (x *StmAggProSma) DstToInstr() int { return x.Stm.DstToInstr() + 1 }
func (x *StmAggProSma) Unsub(id uint32, slot ...uint32) {
	var uSlot uint32
	if len(slot) > 0 {
		uSlot = slot[0]
	}
	x.mu.Lock()
	delete(x.Rxs, sys.Uint64(id, uSlot))
	if len(x.Rxs) == 0 {
		x.Stm.Unsub(x.Id)
	}
	x.mu.Unlock()
}
func (x *StmAggProSma) Rx(inPkt ana.TmeFlt) (r []sys.Act) {
	if x.Length > 0 {
		x.mu.Lock()
		if ana.Cfg.Trc.IsRltStm() {
			sys.Logf("rlt.StmAggProSma(%v).Rx %p inPkt %v", x.Prm(), x, inPkt)
		}
		x.Tmes.Que(inPkt.Tme)
		x.Vals.Que(inPkt.Flt)
		if x.Tmes.Cnt() >= x.Length {
			outPkt := ana.TmeFlt{
				Tme: x.Tmes.At(x.Length - 1),
				Flt: x.Vals.To(x.Length).ProSma(),
			}
			x.Tmes.Dque()
			x.Vals.Dque()
			for _, rx := range x.Rxs {
				r = append(r, &ana.TmeFltTx{Pkt: outPkt, Rx: rx})
			}
		}
		x.mu.Unlock()
	}
	return r
}
func (x *StmBse) AggProAlma(length unt.Unt) Stm {
	r := &StmAggProAlma{}
	r.Slf = r
	r.Stm = x.Slf
	r.Length = length
	r.Id = sys.NextID()
	r.Rxs = make(ana.TmeFltRxs)
	r.Tmes = tmes.New()
	r.Vals = flts.New()
	x.Sub(r.Rx, r.Id)
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
func (x *StmAggProAlma) DstToInstr() int { return x.Stm.DstToInstr() + 1 }
func (x *StmAggProAlma) Unsub(id uint32, slot ...uint32) {
	var uSlot uint32
	if len(slot) > 0 {
		uSlot = slot[0]
	}
	x.mu.Lock()
	delete(x.Rxs, sys.Uint64(id, uSlot))
	if len(x.Rxs) == 0 {
		x.Stm.Unsub(x.Id)
	}
	x.mu.Unlock()
}
func (x *StmAggProAlma) Rx(inPkt ana.TmeFlt) (r []sys.Act) {
	if x.Length > 0 {
		x.mu.Lock()
		if ana.Cfg.Trc.IsRltStm() {
			sys.Logf("rlt.StmAggProAlma(%v).Rx %p inPkt %v", x.Prm(), x, inPkt)
		}
		x.Tmes.Que(inPkt.Tme)
		x.Vals.Que(inPkt.Flt)
		if x.Tmes.Cnt() >= x.Length {
			outPkt := ana.TmeFlt{
				Tme: x.Tmes.At(x.Length - 1),
				Flt: x.Vals.To(x.Length).ProAlma(),
			}
			x.Tmes.Dque()
			x.Vals.Dque()
			for _, rx := range x.Rxs {
				r = append(r, &ana.TmeFltTx{Pkt: outPkt, Rx: rx})
			}
		}
		x.mu.Unlock()
	}
	return r
}
func (x *StmBse) AggEma(length unt.Unt) Stm {
	r := &StmAggEma{}
	r.Slf = r
	r.Stm = x.Slf
	r.Length = length
	r.Id = sys.NextID()
	r.Rxs = make(ana.TmeFltRxs)
	r.Tmes = tmes.New()
	r.Vals = flts.New()
	r.Prv = flt.Min
	x.Sub(r.Rx, r.Id)
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
func (x *StmAggEma) DstToInstr() int { return x.Stm.DstToInstr() + 1 }
func (x *StmAggEma) Unsub(id uint32, slot ...uint32) {
	var uSlot uint32
	if len(slot) > 0 {
		uSlot = slot[0]
	}
	x.mu.Lock()
	delete(x.Rxs, sys.Uint64(id, uSlot))
	if len(x.Rxs) == 0 {
		x.Stm.Unsub(x.Id)
	}
	x.mu.Unlock()
}
func (x *StmAggEma) Rx(inPkt ana.TmeFlt) (r []sys.Act) {
	if x.Length > 0 {
		x.mu.Lock()
		if ana.Cfg.Trc.IsRltStm() {
			sys.Logf("rlt.StmAggEma(%v).Rx %p inPkt %v", x.Prm(), x, inPkt)
		}
		x.Tmes.Que(inPkt.Tme)
		x.Vals.Que(inPkt.Flt)
		if x.Tmes.Cnt() >= x.Length {
			outPkt := ana.TmeFlt{
				Tme: x.Tmes.At(x.Length - 1),
			}
			vals := x.Vals.To(x.Length)
			if x.Prv == flt.Min {
				outPkt.Flt = vals.Sma()
			} else {
				outPkt.Flt = x.Prv + (vals.Lst()-x.Prv)*(flt.Flt(2)/flt.Flt(vals.Cnt()+1))
			}
			x.Prv = outPkt.Flt
			x.Tmes.Dque()
			x.Vals.Dque()
			for _, rx := range x.Rxs {
				r = append(r, &ana.TmeFltTx{Pkt: outPkt, Rx: rx})
			}
		}
		x.mu.Unlock()
	}
	return r
}
func (x *StmBse) InrAdd(off unt.Unt) Stm {
	r := &StmInrAdd{}
	r.Slf = r
	r.Stm = x.Slf
	r.Off = off
	r.Id = sys.NextID()
	r.Rxs = make(ana.TmeFltRxs)
	r.Tmes = tmes.New()
	r.Vals = flts.New()
	x.Sub(r.Rx, r.Id)
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
func (x *StmInrAdd) DstToInstr() int { return x.Stm.DstToInstr() + 1 }
func (x *StmInrAdd) Unsub(id uint32, slot ...uint32) {
	var uSlot uint32
	if len(slot) > 0 {
		uSlot = slot[0]
	}
	x.mu.Lock()
	delete(x.Rxs, sys.Uint64(id, uSlot))
	if len(x.Rxs) == 0 {
		x.Stm.Unsub(x.Id)
	}
	x.mu.Unlock()
}
func (x *StmInrAdd) Rx(inPkt ana.TmeFlt) (r []sys.Act) {
	if x.Off > 0 {
		x.mu.Lock()
		if ana.Cfg.Trc.IsRltStm() {
			sys.Logf("rlt.StmInrAdd(%v).Rx %p inPkt %v", x.Prm(), x, inPkt)
		}
		x.Tmes.Que(inPkt.Tme)
		x.Vals.Que(inPkt.Flt)
		if x.Tmes.Cnt() > x.Off {
			outPkt := ana.TmeFlt{
				Tme: x.Tmes.At(x.Off),
				Flt: x.Vals.At(x.Off).Add(x.Vals.Fst()),
			}
			x.Tmes.Dque()
			x.Vals.Dque()
			for _, rx := range x.Rxs {
				r = append(r, &ana.TmeFltTx{Pkt: outPkt, Rx: rx})
			}
		}
		x.mu.Unlock()
	}
	return r
}
func (x *StmBse) InrSub(off unt.Unt) Stm {
	r := &StmInrSub{}
	r.Slf = r
	r.Stm = x.Slf
	r.Off = off
	r.Id = sys.NextID()
	r.Rxs = make(ana.TmeFltRxs)
	r.Tmes = tmes.New()
	r.Vals = flts.New()
	x.Sub(r.Rx, r.Id)
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
func (x *StmInrSub) DstToInstr() int { return x.Stm.DstToInstr() + 1 }
func (x *StmInrSub) Unsub(id uint32, slot ...uint32) {
	var uSlot uint32
	if len(slot) > 0 {
		uSlot = slot[0]
	}
	x.mu.Lock()
	delete(x.Rxs, sys.Uint64(id, uSlot))
	if len(x.Rxs) == 0 {
		x.Stm.Unsub(x.Id)
	}
	x.mu.Unlock()
}
func (x *StmInrSub) Rx(inPkt ana.TmeFlt) (r []sys.Act) {
	if x.Off > 0 {
		x.mu.Lock()
		if ana.Cfg.Trc.IsRltStm() {
			sys.Logf("rlt.StmInrSub(%v).Rx %p inPkt %v", x.Prm(), x, inPkt)
		}
		x.Tmes.Que(inPkt.Tme)
		x.Vals.Que(inPkt.Flt)
		if x.Tmes.Cnt() > x.Off {
			outPkt := ana.TmeFlt{
				Tme: x.Tmes.At(x.Off),
				Flt: x.Vals.At(x.Off).Sub(x.Vals.Fst()),
			}
			x.Tmes.Dque()
			x.Vals.Dque()
			for _, rx := range x.Rxs {
				r = append(r, &ana.TmeFltTx{Pkt: outPkt, Rx: rx})
			}
		}
		x.mu.Unlock()
	}
	return r
}
func (x *StmBse) InrMul(off unt.Unt) Stm {
	r := &StmInrMul{}
	r.Slf = r
	r.Stm = x.Slf
	r.Off = off
	r.Id = sys.NextID()
	r.Rxs = make(ana.TmeFltRxs)
	r.Tmes = tmes.New()
	r.Vals = flts.New()
	x.Sub(r.Rx, r.Id)
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
func (x *StmInrMul) DstToInstr() int { return x.Stm.DstToInstr() + 1 }
func (x *StmInrMul) Unsub(id uint32, slot ...uint32) {
	var uSlot uint32
	if len(slot) > 0 {
		uSlot = slot[0]
	}
	x.mu.Lock()
	delete(x.Rxs, sys.Uint64(id, uSlot))
	if len(x.Rxs) == 0 {
		x.Stm.Unsub(x.Id)
	}
	x.mu.Unlock()
}
func (x *StmInrMul) Rx(inPkt ana.TmeFlt) (r []sys.Act) {
	if x.Off > 0 {
		x.mu.Lock()
		if ana.Cfg.Trc.IsRltStm() {
			sys.Logf("rlt.StmInrMul(%v).Rx %p inPkt %v", x.Prm(), x, inPkt)
		}
		x.Tmes.Que(inPkt.Tme)
		x.Vals.Que(inPkt.Flt)
		if x.Tmes.Cnt() > x.Off {
			outPkt := ana.TmeFlt{
				Tme: x.Tmes.At(x.Off),
				Flt: x.Vals.At(x.Off).Mul(x.Vals.Fst()),
			}
			x.Tmes.Dque()
			x.Vals.Dque()
			for _, rx := range x.Rxs {
				r = append(r, &ana.TmeFltTx{Pkt: outPkt, Rx: rx})
			}
		}
		x.mu.Unlock()
	}
	return r
}
func (x *StmBse) InrDiv(off unt.Unt) Stm {
	r := &StmInrDiv{}
	r.Slf = r
	r.Stm = x.Slf
	r.Off = off
	r.Id = sys.NextID()
	r.Rxs = make(ana.TmeFltRxs)
	r.Tmes = tmes.New()
	r.Vals = flts.New()
	x.Sub(r.Rx, r.Id)
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
func (x *StmInrDiv) DstToInstr() int { return x.Stm.DstToInstr() + 1 }
func (x *StmInrDiv) Unsub(id uint32, slot ...uint32) {
	var uSlot uint32
	if len(slot) > 0 {
		uSlot = slot[0]
	}
	x.mu.Lock()
	delete(x.Rxs, sys.Uint64(id, uSlot))
	if len(x.Rxs) == 0 {
		x.Stm.Unsub(x.Id)
	}
	x.mu.Unlock()
}
func (x *StmInrDiv) Rx(inPkt ana.TmeFlt) (r []sys.Act) {
	if x.Off > 0 {
		x.mu.Lock()
		if ana.Cfg.Trc.IsRltStm() {
			sys.Logf("rlt.StmInrDiv(%v).Rx %p inPkt %v", x.Prm(), x, inPkt)
		}
		x.Tmes.Que(inPkt.Tme)
		x.Vals.Que(inPkt.Flt)
		if x.Tmes.Cnt() > x.Off {
			outPkt := ana.TmeFlt{
				Tme: x.Tmes.At(x.Off),
				Flt: x.Vals.At(x.Off).Div(x.Vals.Fst()),
			}
			x.Tmes.Dque()
			x.Vals.Dque()
			for _, rx := range x.Rxs {
				r = append(r, &ana.TmeFltTx{Pkt: outPkt, Rx: rx})
			}
		}
		x.mu.Unlock()
	}
	return r
}
func (x *StmBse) InrRem(off unt.Unt) Stm {
	r := &StmInrRem{}
	r.Slf = r
	r.Stm = x.Slf
	r.Off = off
	r.Id = sys.NextID()
	r.Rxs = make(ana.TmeFltRxs)
	r.Tmes = tmes.New()
	r.Vals = flts.New()
	x.Sub(r.Rx, r.Id)
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
func (x *StmInrRem) DstToInstr() int { return x.Stm.DstToInstr() + 1 }
func (x *StmInrRem) Unsub(id uint32, slot ...uint32) {
	var uSlot uint32
	if len(slot) > 0 {
		uSlot = slot[0]
	}
	x.mu.Lock()
	delete(x.Rxs, sys.Uint64(id, uSlot))
	if len(x.Rxs) == 0 {
		x.Stm.Unsub(x.Id)
	}
	x.mu.Unlock()
}
func (x *StmInrRem) Rx(inPkt ana.TmeFlt) (r []sys.Act) {
	if x.Off > 0 {
		x.mu.Lock()
		if ana.Cfg.Trc.IsRltStm() {
			sys.Logf("rlt.StmInrRem(%v).Rx %p inPkt %v", x.Prm(), x, inPkt)
		}
		x.Tmes.Que(inPkt.Tme)
		x.Vals.Que(inPkt.Flt)
		if x.Tmes.Cnt() > x.Off {
			outPkt := ana.TmeFlt{
				Tme: x.Tmes.At(x.Off),
				Flt: x.Vals.At(x.Off).Rem(x.Vals.Fst()),
			}
			x.Tmes.Dque()
			x.Vals.Dque()
			for _, rx := range x.Rxs {
				r = append(r, &ana.TmeFltTx{Pkt: outPkt, Rx: rx})
			}
		}
		x.mu.Unlock()
	}
	return r
}
func (x *StmBse) InrPow(off unt.Unt) Stm {
	r := &StmInrPow{}
	r.Slf = r
	r.Stm = x.Slf
	r.Off = off
	r.Id = sys.NextID()
	r.Rxs = make(ana.TmeFltRxs)
	r.Tmes = tmes.New()
	r.Vals = flts.New()
	x.Sub(r.Rx, r.Id)
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
func (x *StmInrPow) DstToInstr() int { return x.Stm.DstToInstr() + 1 }
func (x *StmInrPow) Unsub(id uint32, slot ...uint32) {
	var uSlot uint32
	if len(slot) > 0 {
		uSlot = slot[0]
	}
	x.mu.Lock()
	delete(x.Rxs, sys.Uint64(id, uSlot))
	if len(x.Rxs) == 0 {
		x.Stm.Unsub(x.Id)
	}
	x.mu.Unlock()
}
func (x *StmInrPow) Rx(inPkt ana.TmeFlt) (r []sys.Act) {
	if x.Off > 0 {
		x.mu.Lock()
		if ana.Cfg.Trc.IsRltStm() {
			sys.Logf("rlt.StmInrPow(%v).Rx %p inPkt %v", x.Prm(), x, inPkt)
		}
		x.Tmes.Que(inPkt.Tme)
		x.Vals.Que(inPkt.Flt)
		if x.Tmes.Cnt() > x.Off {
			outPkt := ana.TmeFlt{
				Tme: x.Tmes.At(x.Off),
				Flt: x.Vals.At(x.Off).Pow(x.Vals.Fst()),
			}
			x.Tmes.Dque()
			x.Vals.Dque()
			for _, rx := range x.Rxs {
				r = append(r, &ana.TmeFltTx{Pkt: outPkt, Rx: rx})
			}
		}
		x.mu.Unlock()
	}
	return r
}
func (x *StmBse) InrMin(off unt.Unt) Stm {
	r := &StmInrMin{}
	r.Slf = r
	r.Stm = x.Slf
	r.Off = off
	r.Id = sys.NextID()
	r.Rxs = make(ana.TmeFltRxs)
	r.Tmes = tmes.New()
	r.Vals = flts.New()
	x.Sub(r.Rx, r.Id)
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
func (x *StmInrMin) DstToInstr() int { return x.Stm.DstToInstr() + 1 }
func (x *StmInrMin) Unsub(id uint32, slot ...uint32) {
	var uSlot uint32
	if len(slot) > 0 {
		uSlot = slot[0]
	}
	x.mu.Lock()
	delete(x.Rxs, sys.Uint64(id, uSlot))
	if len(x.Rxs) == 0 {
		x.Stm.Unsub(x.Id)
	}
	x.mu.Unlock()
}
func (x *StmInrMin) Rx(inPkt ana.TmeFlt) (r []sys.Act) {
	if x.Off > 0 {
		x.mu.Lock()
		if ana.Cfg.Trc.IsRltStm() {
			sys.Logf("rlt.StmInrMin(%v).Rx %p inPkt %v", x.Prm(), x, inPkt)
		}
		x.Tmes.Que(inPkt.Tme)
		x.Vals.Que(inPkt.Flt)
		if x.Tmes.Cnt() > x.Off {
			outPkt := ana.TmeFlt{
				Tme: x.Tmes.At(x.Off),
				Flt: x.Vals.At(x.Off).Min(x.Vals.Fst()),
			}
			x.Tmes.Dque()
			x.Vals.Dque()
			for _, rx := range x.Rxs {
				r = append(r, &ana.TmeFltTx{Pkt: outPkt, Rx: rx})
			}
		}
		x.mu.Unlock()
	}
	return r
}
func (x *StmBse) InrMax(off unt.Unt) Stm {
	r := &StmInrMax{}
	r.Slf = r
	r.Stm = x.Slf
	r.Off = off
	r.Id = sys.NextID()
	r.Rxs = make(ana.TmeFltRxs)
	r.Tmes = tmes.New()
	r.Vals = flts.New()
	x.Sub(r.Rx, r.Id)
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
func (x *StmInrMax) DstToInstr() int { return x.Stm.DstToInstr() + 1 }
func (x *StmInrMax) Unsub(id uint32, slot ...uint32) {
	var uSlot uint32
	if len(slot) > 0 {
		uSlot = slot[0]
	}
	x.mu.Lock()
	delete(x.Rxs, sys.Uint64(id, uSlot))
	if len(x.Rxs) == 0 {
		x.Stm.Unsub(x.Id)
	}
	x.mu.Unlock()
}
func (x *StmInrMax) Rx(inPkt ana.TmeFlt) (r []sys.Act) {
	if x.Off > 0 {
		x.mu.Lock()
		if ana.Cfg.Trc.IsRltStm() {
			sys.Logf("rlt.StmInrMax(%v).Rx %p inPkt %v", x.Prm(), x, inPkt)
		}
		x.Tmes.Que(inPkt.Tme)
		x.Vals.Que(inPkt.Flt)
		if x.Tmes.Cnt() > x.Off {
			outPkt := ana.TmeFlt{
				Tme: x.Tmes.At(x.Off),
				Flt: x.Vals.At(x.Off).Max(x.Vals.Fst()),
			}
			x.Tmes.Dque()
			x.Vals.Dque()
			for _, rx := range x.Rxs {
				r = append(r, &ana.TmeFltTx{Pkt: outPkt, Rx: rx})
			}
		}
		x.mu.Unlock()
	}
	return r
}
func (x *StmBse) InrSlp(off unt.Unt) Stm {
	r := &StmInr1Slp{}
	r.Slf = r
	r.Stm = x.Slf
	r.Off = off
	r.Id = sys.NextID()
	r.Rxs = make(ana.TmeFltRxs)
	r.Tmes = tmes.New()
	r.Vals = flts.New()
	x.Sub(r.Rx, r.Id)
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
func (x *StmInr1Slp) DstToInstr() int { return x.Stm.DstToInstr() + 1 }
func (x *StmInr1Slp) Unsub(id uint32, slot ...uint32) {
	var uSlot uint32
	if len(slot) > 0 {
		uSlot = slot[0]
	}
	x.mu.Lock()
	delete(x.Rxs, sys.Uint64(id, uSlot))
	if len(x.Rxs) == 0 {
		x.Stm.Unsub(x.Id)
	}
	x.mu.Unlock()
}
func (x *StmInr1Slp) Rx(inPkt ana.TmeFlt) (r []sys.Act) {
	if x.Off > 0 {
		x.mu.Lock()
		if ana.Cfg.Trc.IsRltStm() {
			sys.Logf("rlt.StmInr1Slp(%v).Rx %p inPkt %v", x.Prm(), x, inPkt)
		}
		x.Tmes.Que(inPkt.Tme)
		x.Vals.Que(inPkt.Flt)
		if x.Tmes.Cnt() > x.Off {
			rise := x.Vals.At(x.Off).Sub(x.Vals.Fst())
			run := x.Tmes.At(x.Off).Sub(x.Tmes.Fst())
			outPkt := ana.TmeFlt{
				Tme: x.Tmes.At(x.Off),
				Flt: rise / flt.Flt(run),
			}
			x.Tmes.Dque()
			x.Vals.Dque()
			for _, rx := range x.Rxs {
				r = append(r, &ana.TmeFltTx{Pkt: outPkt, Rx: rx})
			}
		}
		x.mu.Unlock()
	}
	return r
}
func (x *StmBse) OtrAdd(off unt.Unt, a Stm) Stm {
	r := &StmOtrAdd{}
	r.Slf = r
	r.Stm = x.Slf
	r.Off = off
	r.A = a
	r.Id = sys.NextID()
	r.Rxs = make(ana.TmeFltRxs)
	r.Tmes = tmes.New()
	r.Vals = flts.New()
	r.TmesA = tmes.New()
	r.ValsA = flts.New()
	x.Sub(r.Rx, r.Id)
	a.Sub(r.RxA, r.Id, SlotA)
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
func (x *StmOtrAdd) DstToInstr() int { return x.Stm.DstToInstr() + 1 }
func (x *StmOtrAdd) Unsub(id uint32, slot ...uint32) {
	var uSlot uint32
	if len(slot) > 0 {
		uSlot = slot[0]
	}
	x.mu.Lock()
	delete(x.Rxs, sys.Uint64(id, uSlot))
	if len(x.Rxs) == 0 {
		x.Stm.Unsub(x.Id)
	}
	x.mu.Unlock()
}
func (x *StmOtrAdd) Rx(inPkt ana.TmeFlt) (r []sys.Act) {
	x.mu.Lock()
	if ana.Cfg.Trc.IsRltStm() {
		sys.Logf("rlt.StmOtrAdd(%v).Rx %p inPkt %v", x.Prm(), x, inPkt)
	}
	x.Tmes.Que(inPkt.Tme)
	x.Vals.Que(inPkt.Flt)
	r = x.Tx()
	x.mu.Unlock()
	return r
}
func (x *StmOtrAdd) RxA(inPkt ana.TmeFlt) (r []sys.Act) {
	x.mu.Lock()
	if ana.Cfg.Trc.IsRltStm() {
		sys.Logf("rlt.StmOtrAdd.RxA %p inPkt %v", x, inPkt)
	}
	x.TmesA.Que(inPkt.Tme)
	x.ValsA.Que(inPkt.Flt)
	r = x.Tx()
	x.mu.Unlock()
	return r
}
func (x *StmOtrAdd) Tx() (r []sys.Act) {
	if x.Tmes.Cnt() == 0 || x.TmesA.Cnt() <= x.Off { // align
		return nil
	}
	if x.Tmes.At(0) != x.TmesA.At(0) {
		if x.Tmes.At(0) < x.TmesA.At(0) {
			for x.Tmes.Cnt() > 0 && x.Tmes.At(0) != x.TmesA.At(0) { // drain X queue until empty or equal
				x.Tmes.Dque()
				x.Vals.Dque()
			}
		} else {
			for x.TmesA.Cnt() > 0 && x.Tmes.At(0) != x.TmesA.At(0) { // drain A queue until empty or equal
				x.TmesA.Dque()
				x.ValsA.Dque()
			}
		}
	}
	if x.Tmes.Cnt() > 0 && x.TmesA.Cnt() > x.Off {
		outPkt := ana.TmeFlt{
			Tme: x.TmesA.At(x.Off),
			Flt: x.Vals.Fst().Add(x.ValsA.At(x.Off)),
		}
		x.Tmes.Dque()
		x.Vals.Dque()
		x.TmesA.Dque()
		x.ValsA.Dque()
		for _, rx := range x.Rxs {
			r = append(r, &ana.TmeFltTx{Pkt: outPkt, Rx: rx})
		}
	}
	return r
}
func (x *StmBse) OtrSub(off unt.Unt, a Stm) Stm {
	r := &StmOtrSub{}
	r.Slf = r
	r.Stm = x.Slf
	r.Off = off
	r.A = a
	r.Id = sys.NextID()
	r.Rxs = make(ana.TmeFltRxs)
	r.Tmes = tmes.New()
	r.Vals = flts.New()
	r.TmesA = tmes.New()
	r.ValsA = flts.New()
	x.Sub(r.Rx, r.Id)
	a.Sub(r.RxA, r.Id, SlotA)
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
func (x *StmOtrSub) DstToInstr() int { return x.Stm.DstToInstr() + 1 }
func (x *StmOtrSub) Unsub(id uint32, slot ...uint32) {
	var uSlot uint32
	if len(slot) > 0 {
		uSlot = slot[0]
	}
	x.mu.Lock()
	delete(x.Rxs, sys.Uint64(id, uSlot))
	if len(x.Rxs) == 0 {
		x.Stm.Unsub(x.Id)
	}
	x.mu.Unlock()
}
func (x *StmOtrSub) Rx(inPkt ana.TmeFlt) (r []sys.Act) {
	x.mu.Lock()
	if ana.Cfg.Trc.IsRltStm() {
		sys.Logf("rlt.StmOtrSub(%v).Rx %p inPkt %v", x.Prm(), x, inPkt)
	}
	x.Tmes.Que(inPkt.Tme)
	x.Vals.Que(inPkt.Flt)
	r = x.Tx()
	x.mu.Unlock()
	return r
}
func (x *StmOtrSub) RxA(inPkt ana.TmeFlt) (r []sys.Act) {
	x.mu.Lock()
	if ana.Cfg.Trc.IsRltStm() {
		sys.Logf("rlt.StmOtrSub.RxA %p inPkt %v", x, inPkt)
	}
	x.TmesA.Que(inPkt.Tme)
	x.ValsA.Que(inPkt.Flt)
	r = x.Tx()
	x.mu.Unlock()
	return r
}
func (x *StmOtrSub) Tx() (r []sys.Act) {
	if x.Tmes.Cnt() == 0 || x.TmesA.Cnt() <= x.Off { // align
		return nil
	}
	if x.Tmes.At(0) != x.TmesA.At(0) {
		if x.Tmes.At(0) < x.TmesA.At(0) {
			for x.Tmes.Cnt() > 0 && x.Tmes.At(0) != x.TmesA.At(0) { // drain X queue until empty or equal
				x.Tmes.Dque()
				x.Vals.Dque()
			}
		} else {
			for x.TmesA.Cnt() > 0 && x.Tmes.At(0) != x.TmesA.At(0) { // drain A queue until empty or equal
				x.TmesA.Dque()
				x.ValsA.Dque()
			}
		}
	}
	if x.Tmes.Cnt() > 0 && x.TmesA.Cnt() > x.Off {
		outPkt := ana.TmeFlt{
			Tme: x.TmesA.At(x.Off),
			Flt: x.Vals.Fst().Sub(x.ValsA.At(x.Off)),
		}
		x.Tmes.Dque()
		x.Vals.Dque()
		x.TmesA.Dque()
		x.ValsA.Dque()
		for _, rx := range x.Rxs {
			r = append(r, &ana.TmeFltTx{Pkt: outPkt, Rx: rx})
		}
	}
	return r
}
func (x *StmBse) OtrMul(off unt.Unt, a Stm) Stm {
	r := &StmOtrMul{}
	r.Slf = r
	r.Stm = x.Slf
	r.Off = off
	r.A = a
	r.Id = sys.NextID()
	r.Rxs = make(ana.TmeFltRxs)
	r.Tmes = tmes.New()
	r.Vals = flts.New()
	r.TmesA = tmes.New()
	r.ValsA = flts.New()
	x.Sub(r.Rx, r.Id)
	a.Sub(r.RxA, r.Id, SlotA)
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
func (x *StmOtrMul) DstToInstr() int { return x.Stm.DstToInstr() + 1 }
func (x *StmOtrMul) Unsub(id uint32, slot ...uint32) {
	var uSlot uint32
	if len(slot) > 0 {
		uSlot = slot[0]
	}
	x.mu.Lock()
	delete(x.Rxs, sys.Uint64(id, uSlot))
	if len(x.Rxs) == 0 {
		x.Stm.Unsub(x.Id)
	}
	x.mu.Unlock()
}
func (x *StmOtrMul) Rx(inPkt ana.TmeFlt) (r []sys.Act) {
	x.mu.Lock()
	if ana.Cfg.Trc.IsRltStm() {
		sys.Logf("rlt.StmOtrMul(%v).Rx %p inPkt %v", x.Prm(), x, inPkt)
	}
	x.Tmes.Que(inPkt.Tme)
	x.Vals.Que(inPkt.Flt)
	r = x.Tx()
	x.mu.Unlock()
	return r
}
func (x *StmOtrMul) RxA(inPkt ana.TmeFlt) (r []sys.Act) {
	x.mu.Lock()
	if ana.Cfg.Trc.IsRltStm() {
		sys.Logf("rlt.StmOtrMul.RxA %p inPkt %v", x, inPkt)
	}
	x.TmesA.Que(inPkt.Tme)
	x.ValsA.Que(inPkt.Flt)
	r = x.Tx()
	x.mu.Unlock()
	return r
}
func (x *StmOtrMul) Tx() (r []sys.Act) {
	if x.Tmes.Cnt() == 0 || x.TmesA.Cnt() <= x.Off { // align
		return nil
	}
	if x.Tmes.At(0) != x.TmesA.At(0) {
		if x.Tmes.At(0) < x.TmesA.At(0) {
			for x.Tmes.Cnt() > 0 && x.Tmes.At(0) != x.TmesA.At(0) { // drain X queue until empty or equal
				x.Tmes.Dque()
				x.Vals.Dque()
			}
		} else {
			for x.TmesA.Cnt() > 0 && x.Tmes.At(0) != x.TmesA.At(0) { // drain A queue until empty or equal
				x.TmesA.Dque()
				x.ValsA.Dque()
			}
		}
	}
	if x.Tmes.Cnt() > 0 && x.TmesA.Cnt() > x.Off {
		outPkt := ana.TmeFlt{
			Tme: x.TmesA.At(x.Off),
			Flt: x.Vals.Fst().Mul(x.ValsA.At(x.Off)),
		}
		x.Tmes.Dque()
		x.Vals.Dque()
		x.TmesA.Dque()
		x.ValsA.Dque()
		for _, rx := range x.Rxs {
			r = append(r, &ana.TmeFltTx{Pkt: outPkt, Rx: rx})
		}
	}
	return r
}
func (x *StmBse) OtrDiv(off unt.Unt, a Stm) Stm {
	r := &StmOtrDiv{}
	r.Slf = r
	r.Stm = x.Slf
	r.Off = off
	r.A = a
	r.Id = sys.NextID()
	r.Rxs = make(ana.TmeFltRxs)
	r.Tmes = tmes.New()
	r.Vals = flts.New()
	r.TmesA = tmes.New()
	r.ValsA = flts.New()
	x.Sub(r.Rx, r.Id)
	a.Sub(r.RxA, r.Id, SlotA)
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
func (x *StmOtrDiv) DstToInstr() int { return x.Stm.DstToInstr() + 1 }
func (x *StmOtrDiv) Unsub(id uint32, slot ...uint32) {
	var uSlot uint32
	if len(slot) > 0 {
		uSlot = slot[0]
	}
	x.mu.Lock()
	delete(x.Rxs, sys.Uint64(id, uSlot))
	if len(x.Rxs) == 0 {
		x.Stm.Unsub(x.Id)
	}
	x.mu.Unlock()
}
func (x *StmOtrDiv) Rx(inPkt ana.TmeFlt) (r []sys.Act) {
	x.mu.Lock()
	if ana.Cfg.Trc.IsRltStm() {
		sys.Logf("rlt.StmOtrDiv(%v).Rx %p inPkt %v", x.Prm(), x, inPkt)
	}
	x.Tmes.Que(inPkt.Tme)
	x.Vals.Que(inPkt.Flt)
	r = x.Tx()
	x.mu.Unlock()
	return r
}
func (x *StmOtrDiv) RxA(inPkt ana.TmeFlt) (r []sys.Act) {
	x.mu.Lock()
	if ana.Cfg.Trc.IsRltStm() {
		sys.Logf("rlt.StmOtrDiv.RxA %p inPkt %v", x, inPkt)
	}
	x.TmesA.Que(inPkt.Tme)
	x.ValsA.Que(inPkt.Flt)
	r = x.Tx()
	x.mu.Unlock()
	return r
}
func (x *StmOtrDiv) Tx() (r []sys.Act) {
	if x.Tmes.Cnt() == 0 || x.TmesA.Cnt() <= x.Off { // align
		return nil
	}
	if x.Tmes.At(0) != x.TmesA.At(0) {
		if x.Tmes.At(0) < x.TmesA.At(0) {
			for x.Tmes.Cnt() > 0 && x.Tmes.At(0) != x.TmesA.At(0) { // drain X queue until empty or equal
				x.Tmes.Dque()
				x.Vals.Dque()
			}
		} else {
			for x.TmesA.Cnt() > 0 && x.Tmes.At(0) != x.TmesA.At(0) { // drain A queue until empty or equal
				x.TmesA.Dque()
				x.ValsA.Dque()
			}
		}
	}
	if x.Tmes.Cnt() > 0 && x.TmesA.Cnt() > x.Off {
		outPkt := ana.TmeFlt{
			Tme: x.TmesA.At(x.Off),
			Flt: x.Vals.Fst().Div(x.ValsA.At(x.Off)),
		}
		x.Tmes.Dque()
		x.Vals.Dque()
		x.TmesA.Dque()
		x.ValsA.Dque()
		for _, rx := range x.Rxs {
			r = append(r, &ana.TmeFltTx{Pkt: outPkt, Rx: rx})
		}
	}
	return r
}
func (x *StmBse) OtrRem(off unt.Unt, a Stm) Stm {
	r := &StmOtrRem{}
	r.Slf = r
	r.Stm = x.Slf
	r.Off = off
	r.A = a
	r.Id = sys.NextID()
	r.Rxs = make(ana.TmeFltRxs)
	r.Tmes = tmes.New()
	r.Vals = flts.New()
	r.TmesA = tmes.New()
	r.ValsA = flts.New()
	x.Sub(r.Rx, r.Id)
	a.Sub(r.RxA, r.Id, SlotA)
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
func (x *StmOtrRem) DstToInstr() int { return x.Stm.DstToInstr() + 1 }
func (x *StmOtrRem) Unsub(id uint32, slot ...uint32) {
	var uSlot uint32
	if len(slot) > 0 {
		uSlot = slot[0]
	}
	x.mu.Lock()
	delete(x.Rxs, sys.Uint64(id, uSlot))
	if len(x.Rxs) == 0 {
		x.Stm.Unsub(x.Id)
	}
	x.mu.Unlock()
}
func (x *StmOtrRem) Rx(inPkt ana.TmeFlt) (r []sys.Act) {
	x.mu.Lock()
	if ana.Cfg.Trc.IsRltStm() {
		sys.Logf("rlt.StmOtrRem(%v).Rx %p inPkt %v", x.Prm(), x, inPkt)
	}
	x.Tmes.Que(inPkt.Tme)
	x.Vals.Que(inPkt.Flt)
	r = x.Tx()
	x.mu.Unlock()
	return r
}
func (x *StmOtrRem) RxA(inPkt ana.TmeFlt) (r []sys.Act) {
	x.mu.Lock()
	if ana.Cfg.Trc.IsRltStm() {
		sys.Logf("rlt.StmOtrRem.RxA %p inPkt %v", x, inPkt)
	}
	x.TmesA.Que(inPkt.Tme)
	x.ValsA.Que(inPkt.Flt)
	r = x.Tx()
	x.mu.Unlock()
	return r
}
func (x *StmOtrRem) Tx() (r []sys.Act) {
	if x.Tmes.Cnt() == 0 || x.TmesA.Cnt() <= x.Off { // align
		return nil
	}
	if x.Tmes.At(0) != x.TmesA.At(0) {
		if x.Tmes.At(0) < x.TmesA.At(0) {
			for x.Tmes.Cnt() > 0 && x.Tmes.At(0) != x.TmesA.At(0) { // drain X queue until empty or equal
				x.Tmes.Dque()
				x.Vals.Dque()
			}
		} else {
			for x.TmesA.Cnt() > 0 && x.Tmes.At(0) != x.TmesA.At(0) { // drain A queue until empty or equal
				x.TmesA.Dque()
				x.ValsA.Dque()
			}
		}
	}
	if x.Tmes.Cnt() > 0 && x.TmesA.Cnt() > x.Off {
		outPkt := ana.TmeFlt{
			Tme: x.TmesA.At(x.Off),
			Flt: x.Vals.Fst().Rem(x.ValsA.At(x.Off)),
		}
		x.Tmes.Dque()
		x.Vals.Dque()
		x.TmesA.Dque()
		x.ValsA.Dque()
		for _, rx := range x.Rxs {
			r = append(r, &ana.TmeFltTx{Pkt: outPkt, Rx: rx})
		}
	}
	return r
}
func (x *StmBse) OtrPow(off unt.Unt, a Stm) Stm {
	r := &StmOtrPow{}
	r.Slf = r
	r.Stm = x.Slf
	r.Off = off
	r.A = a
	r.Id = sys.NextID()
	r.Rxs = make(ana.TmeFltRxs)
	r.Tmes = tmes.New()
	r.Vals = flts.New()
	r.TmesA = tmes.New()
	r.ValsA = flts.New()
	x.Sub(r.Rx, r.Id)
	a.Sub(r.RxA, r.Id, SlotA)
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
func (x *StmOtrPow) DstToInstr() int { return x.Stm.DstToInstr() + 1 }
func (x *StmOtrPow) Unsub(id uint32, slot ...uint32) {
	var uSlot uint32
	if len(slot) > 0 {
		uSlot = slot[0]
	}
	x.mu.Lock()
	delete(x.Rxs, sys.Uint64(id, uSlot))
	if len(x.Rxs) == 0 {
		x.Stm.Unsub(x.Id)
	}
	x.mu.Unlock()
}
func (x *StmOtrPow) Rx(inPkt ana.TmeFlt) (r []sys.Act) {
	x.mu.Lock()
	if ana.Cfg.Trc.IsRltStm() {
		sys.Logf("rlt.StmOtrPow(%v).Rx %p inPkt %v", x.Prm(), x, inPkt)
	}
	x.Tmes.Que(inPkt.Tme)
	x.Vals.Que(inPkt.Flt)
	r = x.Tx()
	x.mu.Unlock()
	return r
}
func (x *StmOtrPow) RxA(inPkt ana.TmeFlt) (r []sys.Act) {
	x.mu.Lock()
	if ana.Cfg.Trc.IsRltStm() {
		sys.Logf("rlt.StmOtrPow.RxA %p inPkt %v", x, inPkt)
	}
	x.TmesA.Que(inPkt.Tme)
	x.ValsA.Que(inPkt.Flt)
	r = x.Tx()
	x.mu.Unlock()
	return r
}
func (x *StmOtrPow) Tx() (r []sys.Act) {
	if x.Tmes.Cnt() == 0 || x.TmesA.Cnt() <= x.Off { // align
		return nil
	}
	if x.Tmes.At(0) != x.TmesA.At(0) {
		if x.Tmes.At(0) < x.TmesA.At(0) {
			for x.Tmes.Cnt() > 0 && x.Tmes.At(0) != x.TmesA.At(0) { // drain X queue until empty or equal
				x.Tmes.Dque()
				x.Vals.Dque()
			}
		} else {
			for x.TmesA.Cnt() > 0 && x.Tmes.At(0) != x.TmesA.At(0) { // drain A queue until empty or equal
				x.TmesA.Dque()
				x.ValsA.Dque()
			}
		}
	}
	if x.Tmes.Cnt() > 0 && x.TmesA.Cnt() > x.Off {
		outPkt := ana.TmeFlt{
			Tme: x.TmesA.At(x.Off),
			Flt: x.Vals.Fst().Pow(x.ValsA.At(x.Off)),
		}
		x.Tmes.Dque()
		x.Vals.Dque()
		x.TmesA.Dque()
		x.ValsA.Dque()
		for _, rx := range x.Rxs {
			r = append(r, &ana.TmeFltTx{Pkt: outPkt, Rx: rx})
		}
	}
	return r
}
func (x *StmBse) OtrMin(off unt.Unt, a Stm) Stm {
	r := &StmOtrMin{}
	r.Slf = r
	r.Stm = x.Slf
	r.Off = off
	r.A = a
	r.Id = sys.NextID()
	r.Rxs = make(ana.TmeFltRxs)
	r.Tmes = tmes.New()
	r.Vals = flts.New()
	r.TmesA = tmes.New()
	r.ValsA = flts.New()
	x.Sub(r.Rx, r.Id)
	a.Sub(r.RxA, r.Id, SlotA)
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
func (x *StmOtrMin) DstToInstr() int { return x.Stm.DstToInstr() + 1 }
func (x *StmOtrMin) Unsub(id uint32, slot ...uint32) {
	var uSlot uint32
	if len(slot) > 0 {
		uSlot = slot[0]
	}
	x.mu.Lock()
	delete(x.Rxs, sys.Uint64(id, uSlot))
	if len(x.Rxs) == 0 {
		x.Stm.Unsub(x.Id)
	}
	x.mu.Unlock()
}
func (x *StmOtrMin) Rx(inPkt ana.TmeFlt) (r []sys.Act) {
	x.mu.Lock()
	if ana.Cfg.Trc.IsRltStm() {
		sys.Logf("rlt.StmOtrMin(%v).Rx %p inPkt %v", x.Prm(), x, inPkt)
	}
	x.Tmes.Que(inPkt.Tme)
	x.Vals.Que(inPkt.Flt)
	r = x.Tx()
	x.mu.Unlock()
	return r
}
func (x *StmOtrMin) RxA(inPkt ana.TmeFlt) (r []sys.Act) {
	x.mu.Lock()
	if ana.Cfg.Trc.IsRltStm() {
		sys.Logf("rlt.StmOtrMin.RxA %p inPkt %v", x, inPkt)
	}
	x.TmesA.Que(inPkt.Tme)
	x.ValsA.Que(inPkt.Flt)
	r = x.Tx()
	x.mu.Unlock()
	return r
}
func (x *StmOtrMin) Tx() (r []sys.Act) {
	if x.Tmes.Cnt() == 0 || x.TmesA.Cnt() <= x.Off { // align
		return nil
	}
	if x.Tmes.At(0) != x.TmesA.At(0) {
		if x.Tmes.At(0) < x.TmesA.At(0) {
			for x.Tmes.Cnt() > 0 && x.Tmes.At(0) != x.TmesA.At(0) { // drain X queue until empty or equal
				x.Tmes.Dque()
				x.Vals.Dque()
			}
		} else {
			for x.TmesA.Cnt() > 0 && x.Tmes.At(0) != x.TmesA.At(0) { // drain A queue until empty or equal
				x.TmesA.Dque()
				x.ValsA.Dque()
			}
		}
	}
	if x.Tmes.Cnt() > 0 && x.TmesA.Cnt() > x.Off {
		outPkt := ana.TmeFlt{
			Tme: x.TmesA.At(x.Off),
			Flt: x.Vals.Fst().Min(x.ValsA.At(x.Off)),
		}
		x.Tmes.Dque()
		x.Vals.Dque()
		x.TmesA.Dque()
		x.ValsA.Dque()
		for _, rx := range x.Rxs {
			r = append(r, &ana.TmeFltTx{Pkt: outPkt, Rx: rx})
		}
	}
	return r
}
func (x *StmBse) OtrMax(off unt.Unt, a Stm) Stm {
	r := &StmOtrMax{}
	r.Slf = r
	r.Stm = x.Slf
	r.Off = off
	r.A = a
	r.Id = sys.NextID()
	r.Rxs = make(ana.TmeFltRxs)
	r.Tmes = tmes.New()
	r.Vals = flts.New()
	r.TmesA = tmes.New()
	r.ValsA = flts.New()
	x.Sub(r.Rx, r.Id)
	a.Sub(r.RxA, r.Id, SlotA)
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
func (x *StmOtrMax) DstToInstr() int { return x.Stm.DstToInstr() + 1 }
func (x *StmOtrMax) Unsub(id uint32, slot ...uint32) {
	var uSlot uint32
	if len(slot) > 0 {
		uSlot = slot[0]
	}
	x.mu.Lock()
	delete(x.Rxs, sys.Uint64(id, uSlot))
	if len(x.Rxs) == 0 {
		x.Stm.Unsub(x.Id)
	}
	x.mu.Unlock()
}
func (x *StmOtrMax) Rx(inPkt ana.TmeFlt) (r []sys.Act) {
	x.mu.Lock()
	if ana.Cfg.Trc.IsRltStm() {
		sys.Logf("rlt.StmOtrMax(%v).Rx %p inPkt %v", x.Prm(), x, inPkt)
	}
	x.Tmes.Que(inPkt.Tme)
	x.Vals.Que(inPkt.Flt)
	r = x.Tx()
	x.mu.Unlock()
	return r
}
func (x *StmOtrMax) RxA(inPkt ana.TmeFlt) (r []sys.Act) {
	x.mu.Lock()
	if ana.Cfg.Trc.IsRltStm() {
		sys.Logf("rlt.StmOtrMax.RxA %p inPkt %v", x, inPkt)
	}
	x.TmesA.Que(inPkt.Tme)
	x.ValsA.Que(inPkt.Flt)
	r = x.Tx()
	x.mu.Unlock()
	return r
}
func (x *StmOtrMax) Tx() (r []sys.Act) {
	if x.Tmes.Cnt() == 0 || x.TmesA.Cnt() <= x.Off { // align
		return nil
	}
	if x.Tmes.At(0) != x.TmesA.At(0) {
		if x.Tmes.At(0) < x.TmesA.At(0) {
			for x.Tmes.Cnt() > 0 && x.Tmes.At(0) != x.TmesA.At(0) { // drain X queue until empty or equal
				x.Tmes.Dque()
				x.Vals.Dque()
			}
		} else {
			for x.TmesA.Cnt() > 0 && x.Tmes.At(0) != x.TmesA.At(0) { // drain A queue until empty or equal
				x.TmesA.Dque()
				x.ValsA.Dque()
			}
		}
	}
	if x.Tmes.Cnt() > 0 && x.TmesA.Cnt() > x.Off {
		outPkt := ana.TmeFlt{
			Tme: x.TmesA.At(x.Off),
			Flt: x.Vals.Fst().Max(x.ValsA.At(x.Off)),
		}
		x.Tmes.Dque()
		x.Vals.Dque()
		x.TmesA.Dque()
		x.ValsA.Dque()
		for _, rx := range x.Rxs {
			r = append(r, &ana.TmeFltTx{Pkt: outPkt, Rx: rx})
		}
	}
	return r
}
func (x *StmBse) SclEql(scl flt.Flt) Cnd {
	r := &CndSclEql{}
	r.Slf = r
	r.Stm = x.Slf
	r.Scl = scl
	r.Id = sys.NextID()
	r.Rxs = make(tme.TmeRxs)
	x.Sub(r.Rx, r.Id)
	return r
}
func (x *StmBse) SclNeq(scl flt.Flt) Cnd {
	r := &CndSclNeq{}
	r.Slf = r
	r.Stm = x.Slf
	r.Scl = scl
	r.Id = sys.NextID()
	r.Rxs = make(tme.TmeRxs)
	x.Sub(r.Rx, r.Id)
	return r
}
func (x *StmBse) SclLss(scl flt.Flt) Cnd {
	r := &CndSclLss{}
	r.Slf = r
	r.Stm = x.Slf
	r.Scl = scl
	r.Id = sys.NextID()
	r.Rxs = make(tme.TmeRxs)
	x.Sub(r.Rx, r.Id)
	return r
}
func (x *StmBse) SclGtr(scl flt.Flt) Cnd {
	r := &CndSclGtr{}
	r.Slf = r
	r.Stm = x.Slf
	r.Scl = scl
	r.Id = sys.NextID()
	r.Rxs = make(tme.TmeRxs)
	x.Sub(r.Rx, r.Id)
	return r
}
func (x *StmBse) SclLeq(scl flt.Flt) Cnd {
	r := &CndSclLeq{}
	r.Slf = r
	r.Stm = x.Slf
	r.Scl = scl
	r.Id = sys.NextID()
	r.Rxs = make(tme.TmeRxs)
	x.Sub(r.Rx, r.Id)
	return r
}
func (x *StmBse) SclGeq(scl flt.Flt) Cnd {
	r := &CndSclGeq{}
	r.Slf = r
	r.Stm = x.Slf
	r.Scl = scl
	r.Id = sys.NextID()
	r.Rxs = make(tme.TmeRxs)
	x.Sub(r.Rx, r.Id)
	return r
}
func (x *StmBse) InrEql(off unt.Unt) Cnd {
	r := &CndInrEql{}
	r.Slf = r
	r.Stm = x.Slf
	r.Off = off
	r.Id = sys.NextID()
	r.Rxs = make(tme.TmeRxs)
	r.Tmes = tmes.New()
	r.Vals = flts.New()
	x.Sub(r.Rx, r.Id)
	return r
}
func (x *StmBse) InrNeq(off unt.Unt) Cnd {
	r := &CndInrNeq{}
	r.Slf = r
	r.Stm = x.Slf
	r.Off = off
	r.Id = sys.NextID()
	r.Rxs = make(tme.TmeRxs)
	r.Tmes = tmes.New()
	r.Vals = flts.New()
	x.Sub(r.Rx, r.Id)
	return r
}
func (x *StmBse) InrLss(off unt.Unt) Cnd {
	r := &CndInrLss{}
	r.Slf = r
	r.Stm = x.Slf
	r.Off = off
	r.Id = sys.NextID()
	r.Rxs = make(tme.TmeRxs)
	r.Tmes = tmes.New()
	r.Vals = flts.New()
	x.Sub(r.Rx, r.Id)
	return r
}
func (x *StmBse) InrGtr(off unt.Unt) Cnd {
	r := &CndInrGtr{}
	r.Slf = r
	r.Stm = x.Slf
	r.Off = off
	r.Id = sys.NextID()
	r.Rxs = make(tme.TmeRxs)
	r.Tmes = tmes.New()
	r.Vals = flts.New()
	x.Sub(r.Rx, r.Id)
	return r
}
func (x *StmBse) InrLeq(off unt.Unt) Cnd {
	r := &CndInrLeq{}
	r.Slf = r
	r.Stm = x.Slf
	r.Off = off
	r.Id = sys.NextID()
	r.Rxs = make(tme.TmeRxs)
	r.Tmes = tmes.New()
	r.Vals = flts.New()
	x.Sub(r.Rx, r.Id)
	return r
}
func (x *StmBse) InrGeq(off unt.Unt) Cnd {
	r := &CndInrGeq{}
	r.Slf = r
	r.Stm = x.Slf
	r.Off = off
	r.Id = sys.NextID()
	r.Rxs = make(tme.TmeRxs)
	r.Tmes = tmes.New()
	r.Vals = flts.New()
	x.Sub(r.Rx, r.Id)
	return r
}
func (x *StmBse) OtrEql(off unt.Unt, a Stm) Cnd {
	r := &CndOtrEql{}
	r.Slf = r
	r.Stm = x.Slf
	r.Off = off
	r.A = a
	r.Id = sys.NextID()
	r.Rxs = make(tme.TmeRxs)
	r.Tmes = tmes.New()
	r.Vals = flts.New()
	r.TmesA = tmes.New()
	r.ValsA = flts.New()
	x.Sub(r.Rx, r.Id)
	a.Sub(r.RxA, r.Id, SlotA)
	return r
}
func (x *StmBse) OtrNeq(off unt.Unt, a Stm) Cnd {
	r := &CndOtrNeq{}
	r.Slf = r
	r.Stm = x.Slf
	r.Off = off
	r.A = a
	r.Id = sys.NextID()
	r.Rxs = make(tme.TmeRxs)
	r.Tmes = tmes.New()
	r.Vals = flts.New()
	r.TmesA = tmes.New()
	r.ValsA = flts.New()
	x.Sub(r.Rx, r.Id)
	a.Sub(r.RxA, r.Id, SlotA)
	return r
}
func (x *StmBse) OtrLss(off unt.Unt, a Stm) Cnd {
	r := &CndOtrLss{}
	r.Slf = r
	r.Stm = x.Slf
	r.Off = off
	r.A = a
	r.Id = sys.NextID()
	r.Rxs = make(tme.TmeRxs)
	r.Tmes = tmes.New()
	r.Vals = flts.New()
	r.TmesA = tmes.New()
	r.ValsA = flts.New()
	x.Sub(r.Rx, r.Id)
	a.Sub(r.RxA, r.Id, SlotA)
	return r
}
func (x *StmBse) OtrGtr(off unt.Unt, a Stm) Cnd {
	r := &CndOtrGtr{}
	r.Slf = r
	r.Stm = x.Slf
	r.Off = off
	r.A = a
	r.Id = sys.NextID()
	r.Rxs = make(tme.TmeRxs)
	r.Tmes = tmes.New()
	r.Vals = flts.New()
	r.TmesA = tmes.New()
	r.ValsA = flts.New()
	x.Sub(r.Rx, r.Id)
	a.Sub(r.RxA, r.Id, SlotA)
	return r
}
func (x *StmBse) OtrLeq(off unt.Unt, a Stm) Cnd {
	r := &CndOtrLeq{}
	r.Slf = r
	r.Stm = x.Slf
	r.Off = off
	r.A = a
	r.Id = sys.NextID()
	r.Rxs = make(tme.TmeRxs)
	r.Tmes = tmes.New()
	r.Vals = flts.New()
	r.TmesA = tmes.New()
	r.ValsA = flts.New()
	x.Sub(r.Rx, r.Id)
	a.Sub(r.RxA, r.Id, SlotA)
	return r
}
func (x *StmBse) OtrGeq(off unt.Unt, a Stm) Cnd {
	r := &CndOtrGeq{}
	r.Slf = r
	r.Stm = x.Slf
	r.Off = off
	r.A = a
	r.Id = sys.NextID()
	r.Rxs = make(tme.TmeRxs)
	r.Tmes = tmes.New()
	r.Vals = flts.New()
	r.TmesA = tmes.New()
	r.ValsA = flts.New()
	x.Sub(r.Rx, r.Id)
	a.Sub(r.RxA, r.Id, SlotA)
	return r
}
