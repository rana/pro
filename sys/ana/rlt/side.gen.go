package rlt

import (
	"strings"
	"sync"
	"sys"
	"sys/ana"
	"sys/bsc/bnd"
	"sys/bsc/flt"
	"sys/bsc/str"
)

type (
	Side interface {
		ana.Pth
		Bse() *SideBse
		Sub(rx ana.TmeFltsRx, id uint32)
		Unsub(id uint32, slot ...uint32)
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
		mu    sync.Mutex
		Id    uint32
		Slf   Side
		Inrvl Inrvl
		Rxs   ana.TmeFltsRxs
	}
	SideScp struct {
		Idx uint32
		Arr []Side
	}
	SideBid struct {
		SideBse
	}
	SideAsk struct {
		SideBse
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
func (x *SideBid) Rx(inPkt bnd.Bnd) (r []sys.Act) {
	x.mu.Lock() // translate time range to val range
	if ana.Cfg.Trc.IsRltSide() {
		sys.Logf("rlt.SideBid(%v).Rx %p inPkt %v", x.Prm(), x, inPkt)
	}
	stm := x.Inrvl.Bse().Instr.Instr().RltStm
	outPkt := ana.TmeFlts{
		Tme:  stm.Tmes.At(inPkt.LstIdx()),
		Flts: stm.BidsByTmeBnd(inPkt),
	}
	for _, rx := range x.Rxs {
		r = append(r, &ana.TmeFltsTx{Pkt: outPkt, Rx: rx})
	}
	x.mu.Unlock()
	return r
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
func (x *SideAsk) Rx(inPkt bnd.Bnd) (r []sys.Act) {
	x.mu.Lock() // translate time range to val range
	if ana.Cfg.Trc.IsRltSide() {
		sys.Logf("rlt.SideAsk(%v).Rx %p inPkt %v", x.Prm(), x, inPkt)
	}
	stm := x.Inrvl.Bse().Instr.Instr().RltStm
	outPkt := ana.TmeFlts{
		Tme:  stm.Tmes.At(inPkt.LstIdx()),
		Flts: stm.AsksByTmeBnd(inPkt),
	}
	for _, rx := range x.Rxs {
		r = append(r, &ana.TmeFltsTx{Pkt: outPkt, Rx: rx})
	}
	x.mu.Unlock()
	return r
}
func (x *SideBse) Bse() *SideBse { return x }
func (x *SideBse) Sub(rx ana.TmeFltsRx, id uint32) {
	x.mu.Lock()
	x.Rxs[sys.Uint64(id, 0)] = rx
	x.mu.Unlock()
}
func (x *SideBse) Unsub(id uint32, slot ...uint32) {
	var uSlot uint32
	if len(slot) > 0 {
		uSlot = slot[0]
	}
	x.mu.Lock()
	delete(x.Rxs, sys.Uint64(id, uSlot))
	if len(x.Rxs) == 0 {
		x.Inrvl.Unsub(x.Id)
	}
	x.mu.Unlock()
}
func (x *SideBse) Fst() Stm {
	r := &StmRteFst{}
	r.Slf = r
	r.Side = x.Slf
	r.Id = sys.NextID()
	r.Rxs = make(ana.TmeFltRxs)
	x.Sub(r.Rx, r.Id)
	return r
}
func (x *SideBse) Lst() Stm {
	r := &StmRteLst{}
	r.Slf = r
	r.Side = x.Slf
	r.Id = sys.NextID()
	r.Rxs = make(ana.TmeFltRxs)
	x.Sub(r.Rx, r.Id)
	return r
}
func (x *SideBse) Sum() Stm {
	r := &StmRteSum{}
	r.Slf = r
	r.Side = x.Slf
	r.Id = sys.NextID()
	r.Rxs = make(ana.TmeFltRxs)
	x.Sub(r.Rx, r.Id)
	return r
}
func (x *SideBse) Prd() Stm {
	r := &StmRtePrd{}
	r.Slf = r
	r.Side = x.Slf
	r.Id = sys.NextID()
	r.Rxs = make(ana.TmeFltRxs)
	x.Sub(r.Rx, r.Id)
	return r
}
func (x *SideBse) Min() Stm {
	r := &StmRteMin{}
	r.Slf = r
	r.Side = x.Slf
	r.Id = sys.NextID()
	r.Rxs = make(ana.TmeFltRxs)
	x.Sub(r.Rx, r.Id)
	return r
}
func (x *SideBse) Max() Stm {
	r := &StmRteMax{}
	r.Slf = r
	r.Side = x.Slf
	r.Id = sys.NextID()
	r.Rxs = make(ana.TmeFltRxs)
	x.Sub(r.Rx, r.Id)
	return r
}
func (x *SideBse) Mid() Stm {
	r := &StmRteMid{}
	r.Slf = r
	r.Side = x.Slf
	r.Id = sys.NextID()
	r.Rxs = make(ana.TmeFltRxs)
	x.Sub(r.Rx, r.Id)
	return r
}
func (x *SideBse) Mdn() Stm {
	r := &StmRteMdn{}
	r.Slf = r
	r.Side = x.Slf
	r.Id = sys.NextID()
	r.Rxs = make(ana.TmeFltRxs)
	x.Sub(r.Rx, r.Id)
	return r
}
func (x *SideBse) Sma() Stm {
	r := &StmRteSma{}
	r.Slf = r
	r.Side = x.Slf
	r.Id = sys.NextID()
	r.Rxs = make(ana.TmeFltRxs)
	x.Sub(r.Rx, r.Id)
	return r
}
func (x *SideBse) Gma() Stm {
	r := &StmRteGma{}
	r.Slf = r
	r.Side = x.Slf
	r.Id = sys.NextID()
	r.Rxs = make(ana.TmeFltRxs)
	x.Sub(r.Rx, r.Id)
	return r
}
func (x *SideBse) Wma() Stm {
	r := &StmRteWma{}
	r.Slf = r
	r.Side = x.Slf
	r.Id = sys.NextID()
	r.Rxs = make(ana.TmeFltRxs)
	x.Sub(r.Rx, r.Id)
	return r
}
func (x *SideBse) Rsi() Stm {
	r := &StmRteRsi{}
	r.Slf = r
	r.Side = x.Slf
	r.Id = sys.NextID()
	r.Rxs = make(ana.TmeFltRxs)
	x.Sub(r.Rx, r.Id)
	return r
}
func (x *SideBse) Wrsi() Stm {
	r := &StmRteWrsi{}
	r.Slf = r
	r.Side = x.Slf
	r.Id = sys.NextID()
	r.Rxs = make(ana.TmeFltRxs)
	x.Sub(r.Rx, r.Id)
	return r
}
func (x *SideBse) Alma() Stm {
	r := &StmRteAlma{}
	r.Slf = r
	r.Side = x.Slf
	r.Id = sys.NextID()
	r.Rxs = make(ana.TmeFltRxs)
	x.Sub(r.Rx, r.Id)
	return r
}
func (x *SideBse) Vrnc() Stm {
	r := &StmRteVrnc{}
	r.Slf = r
	r.Side = x.Slf
	r.Id = sys.NextID()
	r.Rxs = make(ana.TmeFltRxs)
	x.Sub(r.Rx, r.Id)
	return r
}
func (x *SideBse) Std() Stm {
	r := &StmRteStd{}
	r.Slf = r
	r.Side = x.Slf
	r.Id = sys.NextID()
	r.Rxs = make(ana.TmeFltRxs)
	x.Sub(r.Rx, r.Id)
	return r
}
func (x *SideBse) RngFul() Stm {
	r := &StmRteRngFul{}
	r.Slf = r
	r.Side = x.Slf
	r.Id = sys.NextID()
	r.Rxs = make(ana.TmeFltRxs)
	x.Sub(r.Rx, r.Id)
	return r
}
func (x *SideBse) RngLst() Stm {
	r := &StmRteRngLst{}
	r.Slf = r
	r.Side = x.Slf
	r.Id = sys.NextID()
	r.Rxs = make(ana.TmeFltRxs)
	x.Sub(r.Rx, r.Id)
	return r
}
func (x *SideBse) ProLst() Stm {
	r := &StmRteProLst{}
	r.Slf = r
	r.Side = x.Slf
	r.Id = sys.NextID()
	r.Rxs = make(ana.TmeFltRxs)
	x.Sub(r.Rx, r.Id)
	return r
}
func (x *SideBse) ProSma() Stm {
	r := &StmRteProSma{}
	r.Slf = r
	r.Side = x.Slf
	r.Id = sys.NextID()
	r.Rxs = make(ana.TmeFltRxs)
	x.Sub(r.Rx, r.Id)
	return r
}
func (x *SideBse) ProAlma() Stm {
	r := &StmRteProAlma{}
	r.Slf = r
	r.Side = x.Slf
	r.Id = sys.NextID()
	r.Rxs = make(ana.TmeFltRxs)
	x.Sub(r.Rx, r.Id)
	return r
}
func (x *SideBse) Sar(afInc, afMax flt.Flt) Stm {
	r := &StmRte1Sar{}
	r.Slf = r
	r.Side = x.Slf
	r.AfInc = afInc
	r.AfMax = afMax
	r.Id = sys.NextID()
	r.Rxs = make(ana.TmeFltRxs)
	x.Sub(r.Rx, r.Id)
	return r
}
func (x *SideBse) Ema() Stm {
	r := &StmRteEma{}
	r.Slf = r
	r.Side = x.Slf
	r.Id = sys.NextID()
	r.Rxs = make(ana.TmeFltRxs)
	r.Prv = flt.Min
	x.Sub(r.Rx, r.Id)
	return r
}
