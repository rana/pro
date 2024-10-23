package rlt

import (
	"strings"
	"sync"
	"sys"
	"sys/ana"
	"sys/bsc/bnd"
	"sys/bsc/str"
	"sys/bsc/tme"
	"sys/bsc/unt"
	"sys/err"
)

type (
	Inrvl interface {
		ana.Pth
		Bse() *InrvlBse
		Sub(rx bnd.BndRx, id uint32)
		Unsub(id uint32, slot ...uint32)
		Bid() Side
		Ask() Side
	}
	InrvlBse struct {
		mu    sync.Mutex
		Id    uint32
		Slf   Inrvl
		Instr Instr
		Rxs   bnd.BndRxs
		Pkts  *ana.TmeIdxs
	}
	InrvlScp struct {
		Idx uint32
		Arr []Inrvl
	}
	InrvlI struct {
		InrvlBse
		Dur tme.Tme
	}
)

func (x *InrvlI) Name() str.Str             { return str.Str("I") }
func (x *InrvlI) PrmWrt(b *strings.Builder) { x.Dur.StrWrt(b) }
func (x *InrvlI) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *InrvlI) StrWrt(b *strings.Builder) {
	x.Instr.StrWrt(b)
	b.WriteString(".i(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *InrvlI) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *InrvlI) Rx(inPkt ana.TmeIdx) (r []sys.Act) {
	x.mu.Lock()
	defer x.mu.Unlock()
	x.Pkts.Que(inPkt)
	if ana.Cfg.Trc.IsRltInrvl() {
		sys.Logf("rlt.InrvlI(%v).Rx %p inPkt %v", x.Prm(), x, inPkt)
	}
	for x.Pkts.Cnt() != 0 && x.Pkts.Fst().Idx == unt.Max { // drain leading hearbeat tics
		x.Pkts.Dque()
	}
	if x.Pkts.Cnt() < 2 { // all starting heartbeats drained, or fst tic
		return nil
	}
	start := x.Pkts.At(0)
	endTme := start.Tme + x.Dur
	endLim := start.Idx + 1
	for m := unt.One; m < x.Pkts.Cnt(); m++ {
		cur := x.Pkts.At(m)
		if cur.Tme < endTme {
			if cur.Idx != unt.Max {
				endLim = cur.Idx + 1 // set end price idx
			}
		} else { // heartbeat or price advance the inrvl
			outPkt := bnd.Bnd{Idx: start.Idx, Lim: endLim}
			ts := x.Instr.Instr().RltStm.Tmes
			if ts.At(outPkt.Lim-1).Sub(ts.At(outPkt.Idx)) >= x.Dur {
				err.Panicf("INVALID %v INRVL CALC (calc:%v outPkt:%v)", x.Dur, ts.At(outPkt.Lim-1).Sub(ts.At(outPkt.Idx)), outPkt)
			}
			for _, rx := range x.Rxs {
				r = append(r, &bnd.BndTx{Pkt: outPkt, Rx: rx})
			}
			x.Pkts.Dque() // deque fst non-heartbeat for rolling inrvl behavior
			break
		}
	}
	return r
}
func (x *InrvlBse) Bse() *InrvlBse { return x }
func (x *InrvlBse) Sub(rx bnd.BndRx, id uint32) {
	x.mu.Lock()
	x.Rxs[sys.Uint64(id, 0)] = rx
	x.mu.Unlock()
}
func (x *InrvlBse) Unsub(id uint32, slot ...uint32) {
	var uSlot uint32
	if len(slot) > 0 {
		uSlot = slot[0]
	}
	x.mu.Lock()
	delete(x.Rxs, sys.Uint64(id, uSlot))
	if len(x.Rxs) == 0 {
		x.Instr.Unsub(x.Id)
	}
	x.mu.Unlock()
}
func (x *InrvlBse) Bid() Side {
	r := &SideBid{}
	r.Slf = r
	r.Inrvl = x.Slf
	r.Id = sys.NextID()
	r.Rxs = make(ana.TmeFltsRxs)
	x.Sub(r.Rx, r.Id)
	return r
}
func (x *InrvlBse) Ask() Side {
	r := &SideAsk{}
	r.Slf = r
	r.Inrvl = x.Slf
	r.Id = sys.NextID()
	r.Rxs = make(ana.TmeFltsRxs)
	x.Sub(r.Rx, r.Id)
	return r
}
