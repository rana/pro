package rlt

import (
	"strings"
	"sys"
	"sys/ana"
	"sys/bsc/bol"
	"sys/bsc/str"
	"sys/bsc/tme"
	"sys/k"
)

type (
	Prv interface {
		ana.Pth
		Instr(name str.Str) (r *ana.Instr)
		Sub(i *ana.Instr)
		Unsub(i *ana.Instr)
		MayTrd() bol.Bol
		OpnTrd(t *ana.Trd, i *ana.Instr) (ok bol.Bol, rsn ana.TrdRsnOpn)
		ClsTrd(t *ana.Trd, i *ana.Instr) (ok bol.Bol)
		Bse() *PrvBse
		EurUsd(rng ...tme.Rng) Instr
		AudUsd(rng ...tme.Rng) Instr
		NzdUsd(rng ...tme.Rng) Instr
		GbpUsd(rng ...tme.Rng) Instr
	}
	PrvBse struct {
		Slf Prv
	}
	PrvScp struct {
		Idx uint32
		Arr []Prv
	}
	PrvOan struct {
		PrvBse
		*ana.Oan
	}
)

func Oan() Prv {
	r := &PrvOan{}
	r.Slf = r
	r.Oan = ana.PrvOan
	if !ana.Cfg.Test {
		r.Oan.AcntRefresh()
	}
	return r
}
func (x *PrvOan) Name() str.Str             { return str.Str("Oan") }
func (x *PrvOan) PrmWrt(b *strings.Builder) {}
func (x *PrvOan) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *PrvOan) StrWrt(b *strings.Builder) {
	b.WriteString("rlt.oan(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *PrvOan) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *PrvBse) Bse() *PrvBse { return x }
func (x *PrvBse) EurUsd(rng ...tme.Rng) Instr {
	r := &InstrEurUsd{}
	r.Slf = r
	r.Prv = x.Slf
	r.Rng = rng
	r.Id = sys.NextID()
	r.Rxs = make(ana.TmeIdxRxs)
	r.Ana = x.Slf.Instr(k.EurUsdName)
	r.Ana.Sub(r.Rx, r.Id)
	return r
}
func (x *PrvBse) AudUsd(rng ...tme.Rng) Instr {
	r := &InstrAudUsd{}
	r.Slf = r
	r.Prv = x.Slf
	r.Rng = rng
	r.Id = sys.NextID()
	r.Rxs = make(ana.TmeIdxRxs)
	r.Ana = x.Slf.Instr(k.AudUsdName)
	r.Ana.Sub(r.Rx, r.Id)
	return r
}
func (x *PrvBse) NzdUsd(rng ...tme.Rng) Instr {
	r := &InstrNzdUsd{}
	r.Slf = r
	r.Prv = x.Slf
	r.Rng = rng
	r.Id = sys.NextID()
	r.Rxs = make(ana.TmeIdxRxs)
	r.Ana = x.Slf.Instr(k.NzdUsdName)
	r.Ana.Sub(r.Rx, r.Id)
	return r
}
func (x *PrvBse) GbpUsd(rng ...tme.Rng) Instr {
	r := &InstrGbpUsd{}
	r.Slf = r
	r.Prv = x.Slf
	r.Rng = rng
	r.Id = sys.NextID()
	r.Rxs = make(ana.TmeIdxRxs)
	r.Ana = x.Slf.Instr(k.GbpUsdName)
	r.Ana.Sub(r.Rx, r.Id)
	return r
}
