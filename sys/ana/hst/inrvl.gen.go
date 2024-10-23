package hst

import (
	"strings"
	"sys"
	"sys/ana"
	"sys/bsc/bnd"
	"sys/bsc/bnds"
	"sys/bsc/str"
	"sys/bsc/tme"
	"sys/bsc/tmes"
	"sys/err"
)

type (
	Inrvl interface {
		Name() str.Str
		PrmWrt(b *strings.Builder)
		Prm() string
		StrWrt(b *strings.Builder)
		String() string
		Bse() *InrvlBse
		Bid() Side
		Ask() Side
	}
	InrvlBse struct {
		Slf     Inrvl
		Instr   Instr
		Tmes    *tmes.Tmes
		TmeBnds *bnds.Bnds
	}
	InrvlSeg struct {
		bnd.Bnd
		TmeBnds *bnds.Bnds
		Out     *bnds.Bnds
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
func (x *InrvlBse) Bse() *InrvlBse { return x }
func (x *InrvlBse) Bid() Side {
	r := &SideBid{}
	r.Slf = r
	r.Inrvl = x.Slf
	if ana.Cfg.Trc.IsHstSide() {
		sys.Logf("%p hst.SideBid(%v)", r, r.Prm())
	}
	if x.Tmes.Cnt() == 0 {
		return r
	}
	if len(*x.Tmes) != len(*x.TmeBnds) {
		err.Panicf("length unequal (Tmes:%v TmeBnds:%v)", len(*x.Tmes), len(*x.TmeBnds))
	}
	r.Vals = x.Instr.Bse().Ana.HstStm.Bids
	r.ValBnds = bnds.Make(x.Tmes.Cnt())
	segBnds, acts := bnds.Segs(x.Tmes.Cnt())
	for n, segBnd := range *segBnds {
		seg := &SideBidSeg{}
		seg.Idx = segBnd.Idx
		seg.Lim = segBnd.Lim
		seg.Stm = x.Instr.Bse().Ana.HstStm
		seg.TmeBnds = x.TmeBnds
		seg.Out = r.ValBnds
		acts[n] = seg
	}
	sys.Run().Pll(acts...)
	return r
}
func (x *SideBidSeg) Act() {
	for n := x.Idx; n < x.Lim; n++ {
		(*x.Out)[n] = x.Stm.BidBndByTmeBnd((*x.TmeBnds)[n])
	}
}
func (x *InrvlBse) Ask() Side {
	r := &SideAsk{}
	r.Slf = r
	r.Inrvl = x.Slf
	if ana.Cfg.Trc.IsHstSide() {
		sys.Logf("%p hst.SideAsk(%v)", r, r.Prm())
	}
	if x.Tmes.Cnt() == 0 {
		return r
	}
	if len(*x.Tmes) != len(*x.TmeBnds) {
		err.Panicf("length unequal (Tmes:%v TmeBnds:%v)", len(*x.Tmes), len(*x.TmeBnds))
	}
	r.Vals = x.Instr.Bse().Ana.HstStm.Asks
	r.ValBnds = bnds.Make(x.Tmes.Cnt())
	segBnds, acts := bnds.Segs(x.Tmes.Cnt())
	for n, segBnd := range *segBnds {
		seg := &SideAskSeg{}
		seg.Idx = segBnd.Idx
		seg.Lim = segBnd.Lim
		seg.Stm = x.Instr.Bse().Ana.HstStm
		seg.TmeBnds = x.TmeBnds
		seg.Out = r.ValBnds
		acts[n] = seg
	}
	sys.Run().Pll(acts...)
	return r
}
func (x *SideAskSeg) Act() {
	for n := x.Idx; n < x.Lim; n++ {
		(*x.Out)[n] = x.Stm.AskBndByTmeBnd((*x.TmeBnds)[n])
	}
}
