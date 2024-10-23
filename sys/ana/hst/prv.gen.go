package hst

import (
	"strings"
	"sys"
	"sys/ana"
	"sys/bsc/str"
	"sys/bsc/tme"
	"sys/bsc/unt"
	"sys/k"
)

type (
	Prv interface {
		Name() str.Str
		PrmWrt(b *strings.Builder)
		Prm() string
		StrWrt(b *strings.Builder)
		String() string
		Instr(name str.Str) (r *ana.Instr)
		LoadHst(i *ana.Instr)
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
	b.WriteString("hst.oan(")
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
	if ana.Cfg.Trc.IsHstInstr() {
		sys.Logf("%p hst.InstrEurUsd(%v)", r, r.Prm())
	}
	r.Ana = r.Prv.Instr(k.EurUsdName)
	if r.Ana.HstStm == nil {
		r.Prv.LoadHst(r.Ana)
	}
	if len(r.Rng) == 0 {
		r.TmeBnd.Lim = r.Ana.HstStm.Tmes.Cnt()
	} else {
		r.TmeBnd = r.Ana.HstStm.Tmes.Bnd(r.Rng[0])
		if !r.TmeBnd.IsValid() { // 0s-0s for LongRlng
			r.TmeBnd.Idx = unt.Max
		}
	}
	return r
}
func (x *PrvBse) AudUsd(rng ...tme.Rng) Instr {
	r := &InstrAudUsd{}
	r.Slf = r
	r.Prv = x.Slf
	r.Rng = rng
	if ana.Cfg.Trc.IsHstInstr() {
		sys.Logf("%p hst.InstrAudUsd(%v)", r, r.Prm())
	}
	r.Ana = r.Prv.Instr(k.AudUsdName)
	if r.Ana.HstStm == nil {
		r.Prv.LoadHst(r.Ana)
	}
	if len(r.Rng) == 0 {
		r.TmeBnd.Lim = r.Ana.HstStm.Tmes.Cnt()
	} else {
		r.TmeBnd = r.Ana.HstStm.Tmes.Bnd(r.Rng[0])
		if !r.TmeBnd.IsValid() { // 0s-0s for LongRlng
			r.TmeBnd.Idx = unt.Max
		}
	}
	return r
}
func (x *PrvBse) NzdUsd(rng ...tme.Rng) Instr {
	r := &InstrNzdUsd{}
	r.Slf = r
	r.Prv = x.Slf
	r.Rng = rng
	if ana.Cfg.Trc.IsHstInstr() {
		sys.Logf("%p hst.InstrNzdUsd(%v)", r, r.Prm())
	}
	r.Ana = r.Prv.Instr(k.NzdUsdName)
	if r.Ana.HstStm == nil {
		r.Prv.LoadHst(r.Ana)
	}
	if len(r.Rng) == 0 {
		r.TmeBnd.Lim = r.Ana.HstStm.Tmes.Cnt()
	} else {
		r.TmeBnd = r.Ana.HstStm.Tmes.Bnd(r.Rng[0])
		if !r.TmeBnd.IsValid() { // 0s-0s for LongRlng
			r.TmeBnd.Idx = unt.Max
		}
	}
	return r
}
func (x *PrvBse) GbpUsd(rng ...tme.Rng) Instr {
	r := &InstrGbpUsd{}
	r.Slf = r
	r.Prv = x.Slf
	r.Rng = rng
	if ana.Cfg.Trc.IsHstInstr() {
		sys.Logf("%p hst.InstrGbpUsd(%v)", r, r.Prm())
	}
	r.Ana = r.Prv.Instr(k.GbpUsdName)
	if r.Ana.HstStm == nil {
		r.Prv.LoadHst(r.Ana)
	}
	if len(r.Rng) == 0 {
		r.TmeBnd.Lim = r.Ana.HstStm.Tmes.Cnt()
	} else {
		r.TmeBnd = r.Ana.HstStm.Tmes.Bnd(r.Rng[0])
		if !r.TmeBnd.IsValid() { // 0s-0s for LongRlng
			r.TmeBnd.Idx = unt.Max
		}
	}
	return r
}
