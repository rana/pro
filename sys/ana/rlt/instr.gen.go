package rlt

import (
	"strings"
	"sync"
	"sys"
	"sys/ana"
	"sys/bsc/bnd"
	"sys/bsc/str"
	"sys/bsc/tme"
)

type (
	Instr interface {
		ana.Pth
		Bse() *InstrBse
		Instr() *ana.Instr
		Sub(rx ana.TmeIdxRx, id uint32)
		Unsub(id uint32, slot ...uint32)
		I(dur tme.Tme) Inrvl
	}
	InstrBse struct {
		mu  sync.Mutex
		Id  uint32
		Slf Instr
		Prv Prv
		Ana *ana.Instr
		Rxs ana.TmeIdxRxs
	}
	InstrScp struct {
		Idx uint32
		Arr []Instr
	}
	InstrEurUsd struct {
		InstrBse
		Rng []tme.Rng
	}
	InstrAudUsd struct {
		InstrBse
		Rng []tme.Rng
	}
	InstrNzdUsd struct {
		InstrBse
		Rng []tme.Rng
	}
	InstrGbpUsd struct {
		InstrBse
		Rng []tme.Rng
	}
)

func (x *InstrEurUsd) Name() str.Str { return str.Str("EurUsd") }
func (x *InstrEurUsd) PrmWrt(b *strings.Builder) {
	if len(x.Rng) != 0 {
		for n, v := range x.Rng {
			if n != 0 {
				b.WriteRune(' ')
			}
			v.StrWrt(b)
		}
	}
}
func (x *InstrEurUsd) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *InstrEurUsd) StrWrt(b *strings.Builder) {
	x.Prv.StrWrt(b)
	b.WriteString(".eurUsd(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *InstrEurUsd) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *InstrEurUsd) Rx(inPkt ana.TmeIdx) (r []sys.Act) {
	x.mu.Lock() // pass all through
	if ana.Cfg.Trc.IsRltInstr() {
		sys.Logf("rlt.InstrEurUsd(%v).Rx %p inPkt %v", x.Prm(), x, inPkt)
	}
	for _, rx := range x.Rxs {
		r = append(r, &ana.TmeIdxTx{Pkt: inPkt, Rx: rx})
	}
	x.mu.Unlock()
	return r
}
func (x *InstrAudUsd) Name() str.Str { return str.Str("AudUsd") }
func (x *InstrAudUsd) PrmWrt(b *strings.Builder) {
	if len(x.Rng) != 0 {
		for n, v := range x.Rng {
			if n != 0 {
				b.WriteRune(' ')
			}
			v.StrWrt(b)
		}
	}
}
func (x *InstrAudUsd) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *InstrAudUsd) StrWrt(b *strings.Builder) {
	x.Prv.StrWrt(b)
	b.WriteString(".audUsd(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *InstrAudUsd) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *InstrAudUsd) Rx(inPkt ana.TmeIdx) (r []sys.Act) {
	x.mu.Lock() // pass all through
	if ana.Cfg.Trc.IsRltInstr() {
		sys.Logf("rlt.InstrAudUsd(%v).Rx %p inPkt %v", x.Prm(), x, inPkt)
	}
	for _, rx := range x.Rxs {
		r = append(r, &ana.TmeIdxTx{Pkt: inPkt, Rx: rx})
	}
	x.mu.Unlock()
	return r
}
func (x *InstrNzdUsd) Name() str.Str { return str.Str("NzdUsd") }
func (x *InstrNzdUsd) PrmWrt(b *strings.Builder) {
	if len(x.Rng) != 0 {
		for n, v := range x.Rng {
			if n != 0 {
				b.WriteRune(' ')
			}
			v.StrWrt(b)
		}
	}
}
func (x *InstrNzdUsd) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *InstrNzdUsd) StrWrt(b *strings.Builder) {
	x.Prv.StrWrt(b)
	b.WriteString(".nzdUsd(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *InstrNzdUsd) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *InstrNzdUsd) Rx(inPkt ana.TmeIdx) (r []sys.Act) {
	x.mu.Lock() // pass all through
	if ana.Cfg.Trc.IsRltInstr() {
		sys.Logf("rlt.InstrNzdUsd(%v).Rx %p inPkt %v", x.Prm(), x, inPkt)
	}
	for _, rx := range x.Rxs {
		r = append(r, &ana.TmeIdxTx{Pkt: inPkt, Rx: rx})
	}
	x.mu.Unlock()
	return r
}
func (x *InstrGbpUsd) Name() str.Str { return str.Str("GbpUsd") }
func (x *InstrGbpUsd) PrmWrt(b *strings.Builder) {
	if len(x.Rng) != 0 {
		for n, v := range x.Rng {
			if n != 0 {
				b.WriteRune(' ')
			}
			v.StrWrt(b)
		}
	}
}
func (x *InstrGbpUsd) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *InstrGbpUsd) StrWrt(b *strings.Builder) {
	x.Prv.StrWrt(b)
	b.WriteString(".gbpUsd(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *InstrGbpUsd) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *InstrGbpUsd) Rx(inPkt ana.TmeIdx) (r []sys.Act) {
	x.mu.Lock() // pass all through
	if ana.Cfg.Trc.IsRltInstr() {
		sys.Logf("rlt.InstrGbpUsd(%v).Rx %p inPkt %v", x.Prm(), x, inPkt)
	}
	for _, rx := range x.Rxs {
		r = append(r, &ana.TmeIdxTx{Pkt: inPkt, Rx: rx})
	}
	x.mu.Unlock()
	return r
}
func (x *InstrBse) Bse() *InstrBse    { return x }
func (x *InstrBse) Instr() *ana.Instr { return x.Ana }
func (x *InstrBse) Sub(rx ana.TmeIdxRx, id uint32) {
	x.mu.Lock()
	x.Rxs[sys.Uint64(id, 0)] = rx
	x.mu.Unlock()
}
func (x *InstrBse) Unsub(id uint32, slot ...uint32) {
	var uSlot uint32
	if len(slot) > 0 {
		uSlot = slot[0]
	}
	x.mu.Lock()
	delete(x.Rxs, sys.Uint64(id, uSlot))
	if len(x.Rxs) == 0 {
		x.Ana.Unsub(x.Id)
	}
	x.mu.Unlock()
}
func (x *InstrBse) I(dur tme.Tme) Inrvl {
	r := &InrvlI{}
	r.Slf = r
	r.Instr = x.Slf
	r.Dur = dur
	r.Id = sys.NextID()
	r.Rxs = make(bnd.BndRxs)
	r.Pkts = ana.NewTmeIdxs()
	x.Sub(r.Rx, r.Id)
	x.Ana.RltSubsMu.Lock()
	x.Ana.RltInrvlMax = x.Ana.RltInrvlMax.Max(dur) // used to load hst data on fst opn
	x.Ana.RltSubsMu.Unlock()
	return r
}
