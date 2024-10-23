package hst

import (
	"strings"
	"sys/ana"
	"sys/bsc/bnd"
	"sys/bsc/bol"
	"sys/bsc/flt"
	"sys/bsc/str"
	"sys/bsc/tme"
	"sys/bsc/tmes"
)

type (
	Stgy interface {
		Name() str.Str
		PrmWrt(b *strings.Builder)
		Prm() string
		StrWrt(b *strings.Builder)
		String() string
		Bse() *StgyBse
	}
	StgyBse struct {
		Slf       Stgy
		Cnd       Cnd
		IsLong    bol.Bol
		PrfLim    flt.Flt
		LosLim    flt.Flt
		DurLim    tme.Tme
		MinPnlPct flt.Flt
		Instr     Instr
		Clss      []Cnd
		Trds      *ana.Trds
	}
	StgySeg struct {
		bnd.Bnd
	}
	StgyScp struct {
		Idx uint32
		Arr []Stgy
	}
	StgyStgy struct {
		StgyBse
		FtrStms *Stms
	}
	StgyStgySeg struct {
		StgySeg
		Stgy *StgyBse
		Tmes *tmes.Tmes
		Out  *ana.Trds
	}
)

func (x *StgyStgy) Name() str.Str { return str.Str("Stgy") }
func (x *StgyStgy) PrmWrt(b *strings.Builder) {
	x.IsLong.StrWrt(b)
	b.WriteRune(' ')
	x.PrfLim.StrWrt(b)
	b.WriteRune(' ')
	x.LosLim.StrWrt(b)
	b.WriteRune(' ')
	x.DurLim.StrWrt(b)
	b.WriteRune(' ')
	x.MinPnlPct.StrWrt(b)
	b.WriteRune(' ')
	x.Instr.StrWrt(b)
	b.WriteRune(' ')
	x.FtrStms.StrWrt(b)
	if len(x.Clss) != 0 {
		b.WriteRune(' ')
		for n, v := range x.Clss {
			if n != 0 {
				b.WriteRune(' ')
			}
			v.StrWrt(b)
		}
	}
}
func (x *StgyStgy) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *StgyStgy) StrWrt(b *strings.Builder) {
	x.Cnd.StrWrt(b)
	b.WriteString(".stgy(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *StgyStgy) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *StgyStgySeg) Act() {
	for n := x.Idx; n < x.Lim; n++ {
		trd, rsnFail := x.Stgy.OpnClsTrd((*x.Tmes)[n])
		if trd != nil { // may fail to close due to near mkt opn, near mkt cls, spd lim exceeded
			x.Out.Push(trd)
		} else {
			if rsnFail == ana.NoCls {
				break // exit last opn fail to mirror rlt behavior; single ana.NoCls expected
			}
		}
	}
}
func (x *StgyBse) Bse() *StgyBse { return x }
