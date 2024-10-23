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
	"sys/bsc/unt"
)

type (
	Instr interface {
		Name() str.Str
		PrmWrt(b *strings.Builder)
		Prm() string
		StrWrt(b *strings.Builder)
		String() string
		Bse() *InstrBse
		I(dur tme.Tme) Inrvl
	}
	InstrBse struct {
		Slf    Instr
		Prv    Prv
		Ana    *ana.Instr
		TmeBnd bnd.Bnd
		Rng    []tme.Rng
	}
	InstrScp struct {
		Idx uint32
		Arr []Instr
	}
	InstrEurUsd struct {
		InstrBse
	}
	InstrAudUsd struct {
		InstrBse
	}
	InstrNzdUsd struct {
		InstrBse
	}
	InstrGbpUsd struct {
		InstrBse
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
func (x *InstrBse) Bse() *InstrBse { return x }
func (x *InstrBse) I(dur tme.Tme) Inrvl {
	r := &InrvlI{}
	r.Slf = r
	r.Instr = x.Slf
	r.Dur = dur
	if ana.Cfg.Trc.IsHstInrvl() {
		sys.Logf("%p hst.InrvlI(%v)", r, r.Prm())
	}
	if r.Dur < 1 {
		return r
	}
	ts := x.Ana.HstStm.Tmes
	r.Tmes = tmes.Make(ts.Cnt()).Clr()
	r.TmeBnds = bnds.Make(ts.Cnt()).Clr()
	if !x.TmeBnd.IsValid() {
		if x.TmeBnd.Idx == unt.Max { // 0s-0s for LongRlng
			return r
		}
	} else {
		ts = ts.InBnd(x.TmeBnd) // instr time range qualifier .EurUsd(rng)
	}
	for s := 0; s < len(*ts); s++ { // rolling inrvl // EXPECT NO EQUAL START TMES DUE TO INTAKE PROCESSING
		idx := ts.SrchIdxEql((*ts)[s] + r.Dur)
		if idx == unt.Unt(len(*ts)) && (s == len(*ts)-1 || (*ts)[s+1]-(*ts)[s] < ana.MktSessionGap) {
			if (*ts)[s]+r.Dur < ts.Lst() {
				sys.Logf("hst.Instr.I(%v): MISSING INRVL END %v", r.Dur, (*ts)[s]+r.Dur)
			}
			break
		}
		r.Tmes.Push((*ts)[idx-1]) // USE idx-1 TO MATCH RLT BEHAVIOR
		r.TmeBnds.Push(bnd.Bnd{Idx: x.TmeBnd.Idx + unt.Unt(s), Lim: x.TmeBnd.Idx + unt.Unt(idx)})
	}
	if len(*r.Tmes) == 0 {
		r.Tmes = nil
		r.TmeBnds = nil
	}
	return r
}
