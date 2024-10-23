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
	Cnd interface {
		ana.Pth
		Unsub(id uint32, slot ...uint32)
		DstToInstr() int
		Bse() *CndBse
		Sub(rx tme.TmeRx, id uint32, slot ...uint32)
		And(a Cnd) Cnd
		Seq(dur tme.Tme, a Cnd) Cnd
		Stgy(isLong bol.Bol, prfLim, losLim flt.Flt, durLim tme.Tme, minPnlPct flt.Flt, instr Instr, ftrStms *Stms, clss ...Cnd) Stgy
	}
	CndBse struct {
		mu  sync.Mutex
		Id  uint32
		Slf Cnd
		Rxs tme.TmeRxs
	}
	CndScp struct {
		Idx uint32
		Arr []Cnd
	}
	CndSclEql struct {
		CndBse
		Stm Stm
		Scl flt.Flt
	}
	CndSclNeq struct {
		CndBse
		Stm Stm
		Scl flt.Flt
	}
	CndSclLss struct {
		CndBse
		Stm Stm
		Scl flt.Flt
	}
	CndSclGtr struct {
		CndBse
		Stm Stm
		Scl flt.Flt
	}
	CndSclLeq struct {
		CndBse
		Stm Stm
		Scl flt.Flt
	}
	CndSclGeq struct {
		CndBse
		Stm Stm
		Scl flt.Flt
	}
	CndInrEql struct {
		CndBse
		Stm  Stm
		Tmes *tmes.Tmes
		Vals *flts.Flts
		Off  unt.Unt
	}
	CndInrNeq struct {
		CndBse
		Stm  Stm
		Tmes *tmes.Tmes
		Vals *flts.Flts
		Off  unt.Unt
	}
	CndInrLss struct {
		CndBse
		Stm  Stm
		Tmes *tmes.Tmes
		Vals *flts.Flts
		Off  unt.Unt
	}
	CndInrGtr struct {
		CndBse
		Stm  Stm
		Tmes *tmes.Tmes
		Vals *flts.Flts
		Off  unt.Unt
	}
	CndInrLeq struct {
		CndBse
		Stm  Stm
		Tmes *tmes.Tmes
		Vals *flts.Flts
		Off  unt.Unt
	}
	CndInrGeq struct {
		CndBse
		Stm  Stm
		Tmes *tmes.Tmes
		Vals *flts.Flts
		Off  unt.Unt
	}
	CndOtrEql struct {
		CndBse
		Stm   Stm
		Tmes  *tmes.Tmes
		Vals  *flts.Flts
		TmesA *tmes.Tmes
		ValsA *flts.Flts
		Off   unt.Unt
		A     Stm
	}
	CndOtrNeq struct {
		CndBse
		Stm   Stm
		Tmes  *tmes.Tmes
		Vals  *flts.Flts
		TmesA *tmes.Tmes
		ValsA *flts.Flts
		Off   unt.Unt
		A     Stm
	}
	CndOtrLss struct {
		CndBse
		Stm   Stm
		Tmes  *tmes.Tmes
		Vals  *flts.Flts
		TmesA *tmes.Tmes
		ValsA *flts.Flts
		Off   unt.Unt
		A     Stm
	}
	CndOtrGtr struct {
		CndBse
		Stm   Stm
		Tmes  *tmes.Tmes
		Vals  *flts.Flts
		TmesA *tmes.Tmes
		ValsA *flts.Flts
		Off   unt.Unt
		A     Stm
	}
	CndOtrLeq struct {
		CndBse
		Stm   Stm
		Tmes  *tmes.Tmes
		Vals  *flts.Flts
		TmesA *tmes.Tmes
		ValsA *flts.Flts
		Off   unt.Unt
		A     Stm
	}
	CndOtrGeq struct {
		CndBse
		Stm   Stm
		Tmes  *tmes.Tmes
		Vals  *flts.Flts
		TmesA *tmes.Tmes
		ValsA *flts.Flts
		Off   unt.Unt
		A     Stm
	}
	CndCnd1And struct {
		CndBse
		Cnd   Cnd
		Tmes  *tmes.Tmes
		TmesA *tmes.Tmes
		A     Cnd
	}
	CndCnd2Seq struct {
		CndBse
		Cnd   Cnd
		Tmes  *tmes.Tmes
		TmesA *tmes.Tmes
		Dur   tme.Tme
		A     Cnd
	}
)

func (x *CndSclEql) Name() str.Str             { return str.Str("SclEql") }
func (x *CndSclEql) PrmWrt(b *strings.Builder) { x.Scl.StrWrt(b) }
func (x *CndSclEql) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *CndSclEql) StrWrt(b *strings.Builder) {
	x.Stm.StrWrt(b)
	b.WriteString(".sclEql(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *CndSclEql) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *CndSclEql) DstToInstr() int { return x.Stm.DstToInstr() + 1 }
func (x *CndSclEql) Unsub(id uint32, slot ...uint32) {
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
func (x *CndSclEql) Rx(inPkt ana.TmeFlt) (r []sys.Act) {
	x.mu.Lock()
	if ana.Cfg.Trc.IsRltCnd() {
		sys.Logf("rlt.CndSclEql(%v).Rx %p inPkt %v", x.Prm(), x, inPkt)
	}
	if inPkt.Flt.Eql(x.Scl) {
		for _, rx := range x.Rxs {
			r = append(r, &tme.TmeTx{Pkt: inPkt.Tme, Rx: rx})
		}
	}
	x.mu.Unlock()
	return r
}
func (x *CndSclNeq) Name() str.Str             { return str.Str("SclNeq") }
func (x *CndSclNeq) PrmWrt(b *strings.Builder) { x.Scl.StrWrt(b) }
func (x *CndSclNeq) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *CndSclNeq) StrWrt(b *strings.Builder) {
	x.Stm.StrWrt(b)
	b.WriteString(".sclNeq(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *CndSclNeq) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *CndSclNeq) DstToInstr() int { return x.Stm.DstToInstr() + 1 }
func (x *CndSclNeq) Unsub(id uint32, slot ...uint32) {
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
func (x *CndSclNeq) Rx(inPkt ana.TmeFlt) (r []sys.Act) {
	x.mu.Lock()
	if ana.Cfg.Trc.IsRltCnd() {
		sys.Logf("rlt.CndSclNeq(%v).Rx %p inPkt %v", x.Prm(), x, inPkt)
	}
	if inPkt.Flt.Neq(x.Scl) {
		for _, rx := range x.Rxs {
			r = append(r, &tme.TmeTx{Pkt: inPkt.Tme, Rx: rx})
		}
	}
	x.mu.Unlock()
	return r
}
func (x *CndSclLss) Name() str.Str             { return str.Str("SclLss") }
func (x *CndSclLss) PrmWrt(b *strings.Builder) { x.Scl.StrWrt(b) }
func (x *CndSclLss) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *CndSclLss) StrWrt(b *strings.Builder) {
	x.Stm.StrWrt(b)
	b.WriteString(".sclLss(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *CndSclLss) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *CndSclLss) DstToInstr() int { return x.Stm.DstToInstr() + 1 }
func (x *CndSclLss) Unsub(id uint32, slot ...uint32) {
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
func (x *CndSclLss) Rx(inPkt ana.TmeFlt) (r []sys.Act) {
	x.mu.Lock()
	if ana.Cfg.Trc.IsRltCnd() {
		sys.Logf("rlt.CndSclLss(%v).Rx %p inPkt %v", x.Prm(), x, inPkt)
	}
	if inPkt.Flt.Lss(x.Scl) {
		for _, rx := range x.Rxs {
			r = append(r, &tme.TmeTx{Pkt: inPkt.Tme, Rx: rx})
		}
	}
	x.mu.Unlock()
	return r
}
func (x *CndSclGtr) Name() str.Str             { return str.Str("SclGtr") }
func (x *CndSclGtr) PrmWrt(b *strings.Builder) { x.Scl.StrWrt(b) }
func (x *CndSclGtr) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *CndSclGtr) StrWrt(b *strings.Builder) {
	x.Stm.StrWrt(b)
	b.WriteString(".sclGtr(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *CndSclGtr) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *CndSclGtr) DstToInstr() int { return x.Stm.DstToInstr() + 1 }
func (x *CndSclGtr) Unsub(id uint32, slot ...uint32) {
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
func (x *CndSclGtr) Rx(inPkt ana.TmeFlt) (r []sys.Act) {
	x.mu.Lock()
	if ana.Cfg.Trc.IsRltCnd() {
		sys.Logf("rlt.CndSclGtr(%v).Rx %p inPkt %v", x.Prm(), x, inPkt)
	}
	if inPkt.Flt.Gtr(x.Scl) {
		for _, rx := range x.Rxs {
			r = append(r, &tme.TmeTx{Pkt: inPkt.Tme, Rx: rx})
		}
	}
	x.mu.Unlock()
	return r
}
func (x *CndSclLeq) Name() str.Str             { return str.Str("SclLeq") }
func (x *CndSclLeq) PrmWrt(b *strings.Builder) { x.Scl.StrWrt(b) }
func (x *CndSclLeq) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *CndSclLeq) StrWrt(b *strings.Builder) {
	x.Stm.StrWrt(b)
	b.WriteString(".sclLeq(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *CndSclLeq) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *CndSclLeq) DstToInstr() int { return x.Stm.DstToInstr() + 1 }
func (x *CndSclLeq) Unsub(id uint32, slot ...uint32) {
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
func (x *CndSclLeq) Rx(inPkt ana.TmeFlt) (r []sys.Act) {
	x.mu.Lock()
	if ana.Cfg.Trc.IsRltCnd() {
		sys.Logf("rlt.CndSclLeq(%v).Rx %p inPkt %v", x.Prm(), x, inPkt)
	}
	if inPkt.Flt.Leq(x.Scl) {
		for _, rx := range x.Rxs {
			r = append(r, &tme.TmeTx{Pkt: inPkt.Tme, Rx: rx})
		}
	}
	x.mu.Unlock()
	return r
}
func (x *CndSclGeq) Name() str.Str             { return str.Str("SclGeq") }
func (x *CndSclGeq) PrmWrt(b *strings.Builder) { x.Scl.StrWrt(b) }
func (x *CndSclGeq) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *CndSclGeq) StrWrt(b *strings.Builder) {
	x.Stm.StrWrt(b)
	b.WriteString(".sclGeq(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *CndSclGeq) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *CndSclGeq) DstToInstr() int { return x.Stm.DstToInstr() + 1 }
func (x *CndSclGeq) Unsub(id uint32, slot ...uint32) {
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
func (x *CndSclGeq) Rx(inPkt ana.TmeFlt) (r []sys.Act) {
	x.mu.Lock()
	if ana.Cfg.Trc.IsRltCnd() {
		sys.Logf("rlt.CndSclGeq(%v).Rx %p inPkt %v", x.Prm(), x, inPkt)
	}
	if inPkt.Flt.Geq(x.Scl) {
		for _, rx := range x.Rxs {
			r = append(r, &tme.TmeTx{Pkt: inPkt.Tme, Rx: rx})
		}
	}
	x.mu.Unlock()
	return r
}
func (x *CndInrEql) Name() str.Str             { return str.Str("InrEql") }
func (x *CndInrEql) PrmWrt(b *strings.Builder) { x.Off.StrWrt(b) }
func (x *CndInrEql) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *CndInrEql) StrWrt(b *strings.Builder) {
	x.Stm.StrWrt(b)
	b.WriteString(".inrEql(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *CndInrEql) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *CndInrEql) DstToInstr() int { return x.Stm.DstToInstr() + 1 }
func (x *CndInrEql) Unsub(id uint32, slot ...uint32) {
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
func (x *CndInrEql) Rx(inPkt ana.TmeFlt) (r []sys.Act) {
	if x.Off > 0 {
		x.mu.Lock()
		if ana.Cfg.Trc.IsRltCnd() {
			sys.Logf("rlt.CndInrEql(%v).Rx %p inPkt %v", x.Prm(), x, inPkt)
		}
		x.Tmes.Que(inPkt.Tme)
		x.Vals.Que(inPkt.Flt)
		if x.Tmes.Cnt() > x.Off {
			if x.Vals.At(x.Off).Eql(x.Vals.Fst()) {
				for _, rx := range x.Rxs {
					r = append(r, &tme.TmeTx{Pkt: (*x.Tmes)[x.Off], Rx: rx})
				}
			}
			x.Tmes.Dque()
			x.Vals.Dque()
		}
		x.mu.Unlock()
	}
	return r
}
func (x *CndInrNeq) Name() str.Str             { return str.Str("InrNeq") }
func (x *CndInrNeq) PrmWrt(b *strings.Builder) { x.Off.StrWrt(b) }
func (x *CndInrNeq) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *CndInrNeq) StrWrt(b *strings.Builder) {
	x.Stm.StrWrt(b)
	b.WriteString(".inrNeq(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *CndInrNeq) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *CndInrNeq) DstToInstr() int { return x.Stm.DstToInstr() + 1 }
func (x *CndInrNeq) Unsub(id uint32, slot ...uint32) {
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
func (x *CndInrNeq) Rx(inPkt ana.TmeFlt) (r []sys.Act) {
	if x.Off > 0 {
		x.mu.Lock()
		if ana.Cfg.Trc.IsRltCnd() {
			sys.Logf("rlt.CndInrNeq(%v).Rx %p inPkt %v", x.Prm(), x, inPkt)
		}
		x.Tmes.Que(inPkt.Tme)
		x.Vals.Que(inPkt.Flt)
		if x.Tmes.Cnt() > x.Off {
			if x.Vals.At(x.Off).Neq(x.Vals.Fst()) {
				for _, rx := range x.Rxs {
					r = append(r, &tme.TmeTx{Pkt: (*x.Tmes)[x.Off], Rx: rx})
				}
			}
			x.Tmes.Dque()
			x.Vals.Dque()
		}
		x.mu.Unlock()
	}
	return r
}
func (x *CndInrLss) Name() str.Str             { return str.Str("InrLss") }
func (x *CndInrLss) PrmWrt(b *strings.Builder) { x.Off.StrWrt(b) }
func (x *CndInrLss) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *CndInrLss) StrWrt(b *strings.Builder) {
	x.Stm.StrWrt(b)
	b.WriteString(".inrLss(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *CndInrLss) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *CndInrLss) DstToInstr() int { return x.Stm.DstToInstr() + 1 }
func (x *CndInrLss) Unsub(id uint32, slot ...uint32) {
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
func (x *CndInrLss) Rx(inPkt ana.TmeFlt) (r []sys.Act) {
	if x.Off > 0 {
		x.mu.Lock()
		if ana.Cfg.Trc.IsRltCnd() {
			sys.Logf("rlt.CndInrLss(%v).Rx %p inPkt %v", x.Prm(), x, inPkt)
		}
		x.Tmes.Que(inPkt.Tme)
		x.Vals.Que(inPkt.Flt)
		if x.Tmes.Cnt() > x.Off {
			if x.Vals.At(x.Off).Lss(x.Vals.Fst()) {
				for _, rx := range x.Rxs {
					r = append(r, &tme.TmeTx{Pkt: (*x.Tmes)[x.Off], Rx: rx})
				}
			}
			x.Tmes.Dque()
			x.Vals.Dque()
		}
		x.mu.Unlock()
	}
	return r
}
func (x *CndInrGtr) Name() str.Str             { return str.Str("InrGtr") }
func (x *CndInrGtr) PrmWrt(b *strings.Builder) { x.Off.StrWrt(b) }
func (x *CndInrGtr) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *CndInrGtr) StrWrt(b *strings.Builder) {
	x.Stm.StrWrt(b)
	b.WriteString(".inrGtr(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *CndInrGtr) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *CndInrGtr) DstToInstr() int { return x.Stm.DstToInstr() + 1 }
func (x *CndInrGtr) Unsub(id uint32, slot ...uint32) {
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
func (x *CndInrGtr) Rx(inPkt ana.TmeFlt) (r []sys.Act) {
	if x.Off > 0 {
		x.mu.Lock()
		if ana.Cfg.Trc.IsRltCnd() {
			sys.Logf("rlt.CndInrGtr(%v).Rx %p inPkt %v", x.Prm(), x, inPkt)
		}
		x.Tmes.Que(inPkt.Tme)
		x.Vals.Que(inPkt.Flt)
		if x.Tmes.Cnt() > x.Off {
			if x.Vals.At(x.Off).Gtr(x.Vals.Fst()) {
				for _, rx := range x.Rxs {
					r = append(r, &tme.TmeTx{Pkt: (*x.Tmes)[x.Off], Rx: rx})
				}
			}
			x.Tmes.Dque()
			x.Vals.Dque()
		}
		x.mu.Unlock()
	}
	return r
}
func (x *CndInrLeq) Name() str.Str             { return str.Str("InrLeq") }
func (x *CndInrLeq) PrmWrt(b *strings.Builder) { x.Off.StrWrt(b) }
func (x *CndInrLeq) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *CndInrLeq) StrWrt(b *strings.Builder) {
	x.Stm.StrWrt(b)
	b.WriteString(".inrLeq(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *CndInrLeq) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *CndInrLeq) DstToInstr() int { return x.Stm.DstToInstr() + 1 }
func (x *CndInrLeq) Unsub(id uint32, slot ...uint32) {
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
func (x *CndInrLeq) Rx(inPkt ana.TmeFlt) (r []sys.Act) {
	if x.Off > 0 {
		x.mu.Lock()
		if ana.Cfg.Trc.IsRltCnd() {
			sys.Logf("rlt.CndInrLeq(%v).Rx %p inPkt %v", x.Prm(), x, inPkt)
		}
		x.Tmes.Que(inPkt.Tme)
		x.Vals.Que(inPkt.Flt)
		if x.Tmes.Cnt() > x.Off {
			if x.Vals.At(x.Off).Leq(x.Vals.Fst()) {
				for _, rx := range x.Rxs {
					r = append(r, &tme.TmeTx{Pkt: (*x.Tmes)[x.Off], Rx: rx})
				}
			}
			x.Tmes.Dque()
			x.Vals.Dque()
		}
		x.mu.Unlock()
	}
	return r
}
func (x *CndInrGeq) Name() str.Str             { return str.Str("InrGeq") }
func (x *CndInrGeq) PrmWrt(b *strings.Builder) { x.Off.StrWrt(b) }
func (x *CndInrGeq) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *CndInrGeq) StrWrt(b *strings.Builder) {
	x.Stm.StrWrt(b)
	b.WriteString(".inrGeq(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *CndInrGeq) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *CndInrGeq) DstToInstr() int { return x.Stm.DstToInstr() + 1 }
func (x *CndInrGeq) Unsub(id uint32, slot ...uint32) {
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
func (x *CndInrGeq) Rx(inPkt ana.TmeFlt) (r []sys.Act) {
	if x.Off > 0 {
		x.mu.Lock()
		if ana.Cfg.Trc.IsRltCnd() {
			sys.Logf("rlt.CndInrGeq(%v).Rx %p inPkt %v", x.Prm(), x, inPkt)
		}
		x.Tmes.Que(inPkt.Tme)
		x.Vals.Que(inPkt.Flt)
		if x.Tmes.Cnt() > x.Off {
			if x.Vals.At(x.Off).Geq(x.Vals.Fst()) {
				for _, rx := range x.Rxs {
					r = append(r, &tme.TmeTx{Pkt: (*x.Tmes)[x.Off], Rx: rx})
				}
			}
			x.Tmes.Dque()
			x.Vals.Dque()
		}
		x.mu.Unlock()
	}
	return r
}
func (x *CndOtrEql) Name() str.Str { return str.Str("OtrEql") }
func (x *CndOtrEql) PrmWrt(b *strings.Builder) {
	x.Off.StrWrt(b)
	b.WriteRune(' ')
	x.A.StrWrt(b)
}
func (x *CndOtrEql) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *CndOtrEql) StrWrt(b *strings.Builder) {
	x.Stm.StrWrt(b)
	b.WriteString(".otrEql(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *CndOtrEql) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *CndOtrEql) DstToInstr() int { return x.Stm.DstToInstr() + 1 }
func (x *CndOtrEql) Unsub(id uint32, slot ...uint32) {
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
func (x *CndOtrEql) Rx(inPkt ana.TmeFlt) (r []sys.Act) {
	x.mu.Lock()
	if ana.Cfg.Trc.IsRltCnd() {
		sys.Logf("rlt.CndOtrEql(%v).Rx %p inPkt %v", x.Prm(), x, inPkt)
	}
	x.Tmes.Que(inPkt.Tme)
	x.Vals.Que(inPkt.Flt)
	r = x.Tx()
	x.mu.Unlock()
	return r
}
func (x *CndOtrEql) RxA(inPkt ana.TmeFlt) (r []sys.Act) {
	x.mu.Lock()
	if ana.Cfg.Trc.IsRltCnd() {
		sys.Logf("rlt.CndOtrEql(%v).RxA %p inPkt %v", x.Prm(), x, inPkt)
	}
	x.TmesA.Que(inPkt.Tme)
	x.ValsA.Que(inPkt.Flt)
	r = x.Tx()
	x.mu.Unlock()
	return r
}
func (x *CndOtrEql) Tx() (r []sys.Act) {
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
		if x.Vals.Fst().Eql(x.ValsA.At(x.Off)) {
			for _, rx := range x.Rxs {
				r = append(r, &tme.TmeTx{Pkt: (*x.TmesA)[x.Off], Rx: rx})
			}
		}
		x.Tmes.Dque()
		x.Vals.Dque()
		x.TmesA.Dque()
		x.ValsA.Dque()
	}
	return r
}
func (x *CndOtrNeq) Name() str.Str { return str.Str("OtrNeq") }
func (x *CndOtrNeq) PrmWrt(b *strings.Builder) {
	x.Off.StrWrt(b)
	b.WriteRune(' ')
	x.A.StrWrt(b)
}
func (x *CndOtrNeq) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *CndOtrNeq) StrWrt(b *strings.Builder) {
	x.Stm.StrWrt(b)
	b.WriteString(".otrNeq(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *CndOtrNeq) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *CndOtrNeq) DstToInstr() int { return x.Stm.DstToInstr() + 1 }
func (x *CndOtrNeq) Unsub(id uint32, slot ...uint32) {
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
func (x *CndOtrNeq) Rx(inPkt ana.TmeFlt) (r []sys.Act) {
	x.mu.Lock()
	if ana.Cfg.Trc.IsRltCnd() {
		sys.Logf("rlt.CndOtrNeq(%v).Rx %p inPkt %v", x.Prm(), x, inPkt)
	}
	x.Tmes.Que(inPkt.Tme)
	x.Vals.Que(inPkt.Flt)
	r = x.Tx()
	x.mu.Unlock()
	return r
}
func (x *CndOtrNeq) RxA(inPkt ana.TmeFlt) (r []sys.Act) {
	x.mu.Lock()
	if ana.Cfg.Trc.IsRltCnd() {
		sys.Logf("rlt.CndOtrNeq(%v).RxA %p inPkt %v", x.Prm(), x, inPkt)
	}
	x.TmesA.Que(inPkt.Tme)
	x.ValsA.Que(inPkt.Flt)
	r = x.Tx()
	x.mu.Unlock()
	return r
}
func (x *CndOtrNeq) Tx() (r []sys.Act) {
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
		if x.Vals.Fst().Neq(x.ValsA.At(x.Off)) {
			for _, rx := range x.Rxs {
				r = append(r, &tme.TmeTx{Pkt: (*x.TmesA)[x.Off], Rx: rx})
			}
		}
		x.Tmes.Dque()
		x.Vals.Dque()
		x.TmesA.Dque()
		x.ValsA.Dque()
	}
	return r
}
func (x *CndOtrLss) Name() str.Str { return str.Str("OtrLss") }
func (x *CndOtrLss) PrmWrt(b *strings.Builder) {
	x.Off.StrWrt(b)
	b.WriteRune(' ')
	x.A.StrWrt(b)
}
func (x *CndOtrLss) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *CndOtrLss) StrWrt(b *strings.Builder) {
	x.Stm.StrWrt(b)
	b.WriteString(".otrLss(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *CndOtrLss) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *CndOtrLss) DstToInstr() int { return x.Stm.DstToInstr() + 1 }
func (x *CndOtrLss) Unsub(id uint32, slot ...uint32) {
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
func (x *CndOtrLss) Rx(inPkt ana.TmeFlt) (r []sys.Act) {
	x.mu.Lock()
	if ana.Cfg.Trc.IsRltCnd() {
		sys.Logf("rlt.CndOtrLss(%v).Rx %p inPkt %v", x.Prm(), x, inPkt)
	}
	x.Tmes.Que(inPkt.Tme)
	x.Vals.Que(inPkt.Flt)
	r = x.Tx()
	x.mu.Unlock()
	return r
}
func (x *CndOtrLss) RxA(inPkt ana.TmeFlt) (r []sys.Act) {
	x.mu.Lock()
	if ana.Cfg.Trc.IsRltCnd() {
		sys.Logf("rlt.CndOtrLss(%v).RxA %p inPkt %v", x.Prm(), x, inPkt)
	}
	x.TmesA.Que(inPkt.Tme)
	x.ValsA.Que(inPkt.Flt)
	r = x.Tx()
	x.mu.Unlock()
	return r
}
func (x *CndOtrLss) Tx() (r []sys.Act) {
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
		if x.Vals.Fst().Lss(x.ValsA.At(x.Off)) {
			for _, rx := range x.Rxs {
				r = append(r, &tme.TmeTx{Pkt: (*x.TmesA)[x.Off], Rx: rx})
			}
		}
		x.Tmes.Dque()
		x.Vals.Dque()
		x.TmesA.Dque()
		x.ValsA.Dque()
	}
	return r
}
func (x *CndOtrGtr) Name() str.Str { return str.Str("OtrGtr") }
func (x *CndOtrGtr) PrmWrt(b *strings.Builder) {
	x.Off.StrWrt(b)
	b.WriteRune(' ')
	x.A.StrWrt(b)
}
func (x *CndOtrGtr) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *CndOtrGtr) StrWrt(b *strings.Builder) {
	x.Stm.StrWrt(b)
	b.WriteString(".otrGtr(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *CndOtrGtr) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *CndOtrGtr) DstToInstr() int { return x.Stm.DstToInstr() + 1 }
func (x *CndOtrGtr) Unsub(id uint32, slot ...uint32) {
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
func (x *CndOtrGtr) Rx(inPkt ana.TmeFlt) (r []sys.Act) {
	x.mu.Lock()
	if ana.Cfg.Trc.IsRltCnd() {
		sys.Logf("rlt.CndOtrGtr(%v).Rx %p inPkt %v", x.Prm(), x, inPkt)
	}
	x.Tmes.Que(inPkt.Tme)
	x.Vals.Que(inPkt.Flt)
	r = x.Tx()
	x.mu.Unlock()
	return r
}
func (x *CndOtrGtr) RxA(inPkt ana.TmeFlt) (r []sys.Act) {
	x.mu.Lock()
	if ana.Cfg.Trc.IsRltCnd() {
		sys.Logf("rlt.CndOtrGtr(%v).RxA %p inPkt %v", x.Prm(), x, inPkt)
	}
	x.TmesA.Que(inPkt.Tme)
	x.ValsA.Que(inPkt.Flt)
	r = x.Tx()
	x.mu.Unlock()
	return r
}
func (x *CndOtrGtr) Tx() (r []sys.Act) {
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
		if x.Vals.Fst().Gtr(x.ValsA.At(x.Off)) {
			for _, rx := range x.Rxs {
				r = append(r, &tme.TmeTx{Pkt: (*x.TmesA)[x.Off], Rx: rx})
			}
		}
		x.Tmes.Dque()
		x.Vals.Dque()
		x.TmesA.Dque()
		x.ValsA.Dque()
	}
	return r
}
func (x *CndOtrLeq) Name() str.Str { return str.Str("OtrLeq") }
func (x *CndOtrLeq) PrmWrt(b *strings.Builder) {
	x.Off.StrWrt(b)
	b.WriteRune(' ')
	x.A.StrWrt(b)
}
func (x *CndOtrLeq) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *CndOtrLeq) StrWrt(b *strings.Builder) {
	x.Stm.StrWrt(b)
	b.WriteString(".otrLeq(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *CndOtrLeq) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *CndOtrLeq) DstToInstr() int { return x.Stm.DstToInstr() + 1 }
func (x *CndOtrLeq) Unsub(id uint32, slot ...uint32) {
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
func (x *CndOtrLeq) Rx(inPkt ana.TmeFlt) (r []sys.Act) {
	x.mu.Lock()
	if ana.Cfg.Trc.IsRltCnd() {
		sys.Logf("rlt.CndOtrLeq(%v).Rx %p inPkt %v", x.Prm(), x, inPkt)
	}
	x.Tmes.Que(inPkt.Tme)
	x.Vals.Que(inPkt.Flt)
	r = x.Tx()
	x.mu.Unlock()
	return r
}
func (x *CndOtrLeq) RxA(inPkt ana.TmeFlt) (r []sys.Act) {
	x.mu.Lock()
	if ana.Cfg.Trc.IsRltCnd() {
		sys.Logf("rlt.CndOtrLeq(%v).RxA %p inPkt %v", x.Prm(), x, inPkt)
	}
	x.TmesA.Que(inPkt.Tme)
	x.ValsA.Que(inPkt.Flt)
	r = x.Tx()
	x.mu.Unlock()
	return r
}
func (x *CndOtrLeq) Tx() (r []sys.Act) {
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
		if x.Vals.Fst().Leq(x.ValsA.At(x.Off)) {
			for _, rx := range x.Rxs {
				r = append(r, &tme.TmeTx{Pkt: (*x.TmesA)[x.Off], Rx: rx})
			}
		}
		x.Tmes.Dque()
		x.Vals.Dque()
		x.TmesA.Dque()
		x.ValsA.Dque()
	}
	return r
}
func (x *CndOtrGeq) Name() str.Str { return str.Str("OtrGeq") }
func (x *CndOtrGeq) PrmWrt(b *strings.Builder) {
	x.Off.StrWrt(b)
	b.WriteRune(' ')
	x.A.StrWrt(b)
}
func (x *CndOtrGeq) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *CndOtrGeq) StrWrt(b *strings.Builder) {
	x.Stm.StrWrt(b)
	b.WriteString(".otrGeq(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *CndOtrGeq) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *CndOtrGeq) DstToInstr() int { return x.Stm.DstToInstr() + 1 }
func (x *CndOtrGeq) Unsub(id uint32, slot ...uint32) {
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
func (x *CndOtrGeq) Rx(inPkt ana.TmeFlt) (r []sys.Act) {
	x.mu.Lock()
	if ana.Cfg.Trc.IsRltCnd() {
		sys.Logf("rlt.CndOtrGeq(%v).Rx %p inPkt %v", x.Prm(), x, inPkt)
	}
	x.Tmes.Que(inPkt.Tme)
	x.Vals.Que(inPkt.Flt)
	r = x.Tx()
	x.mu.Unlock()
	return r
}
func (x *CndOtrGeq) RxA(inPkt ana.TmeFlt) (r []sys.Act) {
	x.mu.Lock()
	if ana.Cfg.Trc.IsRltCnd() {
		sys.Logf("rlt.CndOtrGeq(%v).RxA %p inPkt %v", x.Prm(), x, inPkt)
	}
	x.TmesA.Que(inPkt.Tme)
	x.ValsA.Que(inPkt.Flt)
	r = x.Tx()
	x.mu.Unlock()
	return r
}
func (x *CndOtrGeq) Tx() (r []sys.Act) {
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
		if x.Vals.Fst().Geq(x.ValsA.At(x.Off)) {
			for _, rx := range x.Rxs {
				r = append(r, &tme.TmeTx{Pkt: (*x.TmesA)[x.Off], Rx: rx})
			}
		}
		x.Tmes.Dque()
		x.Vals.Dque()
		x.TmesA.Dque()
		x.ValsA.Dque()
	}
	return r
}
func (x *CndBse) Bse() *CndBse { return x }
func (x *CndBse) Sub(rx tme.TmeRx, id uint32, slot ...uint32) {
	var uSlot uint32
	if len(slot) > 0 {
		uSlot = slot[0]
	}
	x.mu.Lock()
	x.Rxs[sys.Uint64(id, uSlot)] = rx
	x.mu.Unlock()
}
func (x *CndBse) And(a Cnd) Cnd {
	r := &CndCnd1And{}
	r.Slf = r
	r.Cnd = x.Slf
	r.A = a
	r.Id = sys.NextID()
	r.Rxs = make(tme.TmeRxs)
	r.Tmes = tmes.New()
	r.TmesA = tmes.New()
	x.Sub(r.Rx, r.Id)
	a.Sub(r.RxA, r.Id, SlotA)
	return r
}
func (x *CndCnd1And) Name() str.Str             { return str.Str("And") }
func (x *CndCnd1And) PrmWrt(b *strings.Builder) { x.A.StrWrt(b) }
func (x *CndCnd1And) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *CndCnd1And) StrWrt(b *strings.Builder) {
	x.Cnd.StrWrt(b)
	b.WriteString(".and(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *CndCnd1And) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *CndCnd1And) DstToInstr() int { return x.Cnd.DstToInstr() + 1 }
func (x *CndCnd1And) Unsub(id uint32, slot ...uint32) {
	var uSlot uint32
	if len(slot) > 0 {
		uSlot = slot[0]
	}
	x.mu.Lock()
	delete(x.Rxs, sys.Uint64(id, uSlot))
	if len(x.Rxs) == 0 {
		x.Cnd.Unsub(x.Id)
		x.A.Unsub(x.Id, SlotA)
	}
	x.mu.Unlock()
}
func (x *CndCnd1And) Rx(inPkt tme.Tme) (r []sys.Act) {
	x.mu.Lock()
	if ana.Cfg.Trc.IsRltCnd() {
		sys.Logf("rlt.CndCnd1And(%v).Rx %p inPkt %v", x.Prm(), x, inPkt)
	}
	x.Tmes.Que(inPkt)
	r = x.Tx()
	x.mu.Unlock()
	return r
}
func (x *CndCnd1And) RxA(inPkt tme.Tme) (r []sys.Act) {
	x.mu.Lock()
	if ana.Cfg.Trc.IsRltCnd() {
		sys.Logf("rlt.CndCnd1And(%v).RxA %p inPkt %v", x.Prm(), x, inPkt)
	}
	x.TmesA.Que(inPkt)
	r = x.Tx()
	x.mu.Unlock()
	return r
}
func (x *CndCnd1And) Tx() (r []sys.Act) {
	if len(*x.Tmes) == 0 || len(*x.TmesA) == 0 { // align
		return nil
	}
	if (*x.Tmes)[0] != (*x.TmesA)[0] {
		if (*x.Tmes)[0] < (*x.TmesA)[0] {
			for len(*x.Tmes) > 0 && (*x.Tmes)[0] != (*x.TmesA)[0] { // drain X queue until empty or equal
				x.Tmes.Dque()
			}
		} else {
			for len(*x.TmesA) > 0 && (*x.Tmes)[0] != (*x.TmesA)[0] { // drain A queue until empty or equal
				x.TmesA.Dque()
			}
		}
	}
	if len(*x.Tmes) == 0 || len(*x.TmesA) == 0 {
		return nil
	}
	if (*x.Tmes)[0] == (*x.TmesA)[0] {
		for _, rx := range x.Rxs {
			r = append(r, tme.NewTmeTx((*x.TmesA)[0], rx))
		}
	}
	x.Tmes.Dque()
	x.TmesA.Dque()
	return r
}
func (x *CndBse) Seq(dur tme.Tme, a Cnd) Cnd {
	r := &CndCnd2Seq{}
	r.Slf = r
	r.Cnd = x.Slf
	r.Dur = dur
	r.A = a
	r.Id = sys.NextID()
	r.Rxs = make(tme.TmeRxs)
	r.Tmes = tmes.New()
	r.TmesA = tmes.New()
	x.Sub(r.Rx, r.Id)
	a.Sub(r.RxA, r.Id, SlotA)
	return r
}
func (x *CndCnd2Seq) Name() str.Str { return str.Str("Seq") }
func (x *CndCnd2Seq) PrmWrt(b *strings.Builder) {
	x.Dur.StrWrt(b)
	b.WriteRune(' ')
	x.A.StrWrt(b)
}
func (x *CndCnd2Seq) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *CndCnd2Seq) StrWrt(b *strings.Builder) {
	x.Cnd.StrWrt(b)
	b.WriteString(".seq(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *CndCnd2Seq) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *CndCnd2Seq) DstToInstr() int { return x.Cnd.DstToInstr() + 1 }
func (x *CndCnd2Seq) Unsub(id uint32, slot ...uint32) {
	var uSlot uint32
	if len(slot) > 0 {
		uSlot = slot[0]
	}
	x.mu.Lock()
	delete(x.Rxs, sys.Uint64(id, uSlot))
	if len(x.Rxs) == 0 {
		x.Cnd.Unsub(x.Id)
		x.A.Unsub(x.Id, SlotA)
	}
	x.mu.Unlock()
}
func (x *CndCnd2Seq) Rx(inPkt tme.Tme) (r []sys.Act) {
	x.mu.Lock()
	if ana.Cfg.Trc.IsRltCnd() {
		sys.Logf("rlt.CndCnd2Seq(%v).Rx %p inPkt %v", x.Prm(), x, inPkt)
	}
	x.Tmes.Que(inPkt)
	r = x.Tx()
	x.mu.Unlock()
	return r
}
func (x *CndCnd2Seq) RxA(inPkt tme.Tme) (r []sys.Act) {
	x.mu.Lock()
	if ana.Cfg.Trc.IsRltCnd() {
		sys.Logf("rlt.CndCnd2Seq(%v).RxA %p inPkt %v", x.Prm(), x, inPkt)
	}
	x.TmesA.Que(inPkt)
	r = x.Tx()
	x.mu.Unlock()
	return r
}
func (x *CndCnd2Seq) Tx() (r []sys.Act) {
	if len(*x.Tmes) == 0 || len(*x.TmesA) == 0 {
		return nil
	}
	if (*x.Tmes)[0]+x.Dur != (*x.TmesA)[0] { // align; X & A must have GapFil
		if (*x.Tmes)[0]+x.Dur < (*x.TmesA)[0] {
			for len(*x.Tmes) > 1 && (*x.Tmes)[0]+x.Dur < (*x.TmesA)[0] { // drain X until empty or equal
				x.Tmes.Dque()
			}
		} else if (*x.Tmes)[0]+x.Dur > (*x.TmesA)[0] {
			for len(*x.TmesA) > 1 && (*x.Tmes)[0]+x.Dur > (*x.TmesA)[0] { // drain A until empty or equal
				x.TmesA.Dque()
			}
		}
	}
	if (*x.Tmes)[0]+x.Dur == (*x.TmesA)[0] {
		for _, rx := range x.Rxs {
			r = append(r, tme.NewTmeTx((*x.TmesA)[0], rx))
		}
		x.Tmes.Dque()
		x.TmesA.Dque()
	}
	return r
}
func (x *CndBse) Stgy(isLong bol.Bol, prfLim, losLim flt.Flt, durLim tme.Tme, minPnlPct flt.Flt, instr Instr, ftrStms *Stms, clss ...Cnd) Stgy {
	r := &StgyStgy{}
	r.Slf = r
	r.Cnd = x.Slf
	r.IsLong = isLong
	r.PrfLim = prfLim
	r.LosLim = losLim
	r.DurLim = durLim
	r.MinPnlPct = minPnlPct
	r.Instr = instr
	r.FtrStms = ftrStms
	r.Clss = clss
	r.Id = sys.NextID()
	r.Rxs = make(ana.TrdRxs)
	r.Key = r.String()
	if !ana.Cfg.Test {
		sys.Lrnr().LoadNetFromDsk(r.Key)
	}
	r.Cnd.Sub(r.RxOpn, r.Id)
	r.Instr.Sub(r.RxClsLim, r.Id)
	for _, cndCls := range r.Clss {
		cndCls.Sub(r.RxClsCnd, r.Id)
	}
	r.opns = tmes.New()
	r.stgyFtrStms = make([]*StgyFtrStm, ftrStms.Cnt())
	for n, ftrStm := range *ftrStms {
		v := &StgyFtrStm{}
		r.stgyFtrStms[n] = v
		v.Stgy = r
		v.Name = ftrStm.String()
		v.Tmes = tmes.New()
		v.Vals = flts.New()
		ftrStm.Sub(v.Rx, r.Id)
	}
	return r
}
