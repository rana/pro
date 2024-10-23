package rlt

import (
	"strings"
	"sync"
	"sys"
	"sys/ana"
	"sys/bsc/bol"
	"sys/bsc/flt"
	"sys/bsc/str"
	"sys/bsc/tme"
	"sys/bsc/tmes"
	"sys/bsc/unt"
)

type (
	Stgy interface {
		ana.Pth
		Bse() *StgyBse
		Sub(rx ana.TrdRx, id uint32, slot ...uint32)
		Unsub(id uint32, slot ...uint32)
		I() *ana.Instr
	}
	StgyBse struct {
		mu          sync.Mutex
		Id          uint32
		Slf         Stgy
		Cnd         Cnd
		Rxs         ana.TrdRxs
		IsLong      bol.Bol
		PrfLim      flt.Flt
		LosLim      flt.Flt
		DurLim      tme.Tme
		MinPnlPct   flt.Flt
		Instr       Instr
		FtrStms     *Stms
		Clss        []Cnd
		ClsPrfLim   flt.Flt
		ClsLosLim   flt.Flt
		ClsTmeLim   tme.Tme
		LstClsTme   tme.Tme
		LstClsIdx   unt.Unt
		OpnIdx      unt.Unt
		Trd         *ana.Trd
		opns        *tmes.Tmes
		Key         string
		stgyFtrStms []*StgyFtrStm
	}
	StgyScp struct {
		Idx uint32
		Arr []Stgy
	}
	StgyStgy struct {
		StgyBse
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
func (x *StgyBse) Bse() *StgyBse { return x }
func (x *StgyBse) Sub(rx ana.TrdRx, id uint32, slot ...uint32) {
	var uSlot uint32
	if len(slot) > 0 {
		uSlot = slot[0]
	}
	x.mu.Lock()
	x.Rxs[sys.Uint64(id, uSlot)] = rx
	x.mu.Unlock()
}
func (x *StgyBse) Unsub(id uint32, slot ...uint32) {
	var uSlot uint32
	if len(slot) > 0 {
		uSlot = slot[0]
	}
	x.mu.Lock()
	delete(x.Rxs, sys.Uint64(id, uSlot))
	if len(x.Rxs) == 0 {
		x.Cnd.Unsub(x.Id)
		x.Instr.Unsub(x.Id)
		for _, ftrStm := range *x.FtrStms {
			ftrStm.Unsub(x.Id)
		}
		for _, cndCls := range x.Clss {
			cndCls.Unsub(x.Id)
		}
	}
	x.mu.Unlock()
}
func (x *StgyBse) I() *ana.Instr { return x.Instr.Instr() }
