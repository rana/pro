package ana

import (
	"strings"
	"sys"
	"sys/bsc/bol"
	"sys/bsc/flt"
	"sys/bsc/str"
	"sys/bsc/tme"
)

type (
	Trd struct {
		OpnTme        tme.Tme
		ClsTme        tme.Tme
		OpnBid        flt.Flt
		ClsBid        flt.Flt
		OpnAsk        flt.Flt
		ClsAsk        flt.Flt
		OpnSpd        flt.Flt
		ClsSpd        flt.Flt
		ClsRsn        str.Str
		Pip           flt.Flt
		Dur           tme.Tme
		IsLong        bol.Bol
		PnlPct        flt.Flt
		PnlPctPredict flt.Flt
		PnlUsd        flt.Flt
		PnlGrsUsd     flt.Flt
		CstComUsd     flt.Flt
		CstClsSpdUsd  flt.Flt
		CstOpnSpdUsd  flt.Flt
		OpnBalUsd     flt.Flt
		ClsBalUsd     flt.Flt
		ClsBalUsdAct  flt.Flt
		TrdPct        flt.Flt
		MrgnRtio      flt.Flt
		Units         flt.Flt
		Instr         str.Str
		OpnReq        str.Str
		OpnRes        str.Str
		ClsReq        str.Str
		ClsRes        str.Str
	}
	TrdRx  func(pkt *Trd) []sys.Act
	TrdRxs map[uint64]TrdRx
	TrdTx  struct {
		Pkt  *Trd
		Rx   TrdRx
		ret  []sys.Act
		tier int
	}
	TrdTxr interface {
		Sub(rx TrdRx, id uint32, slot ...uint32)
		Unsub(id uint32, slot ...uint32)
	}
	TrdScp struct {
		Idx uint32
		Arr []*Trd
	}
)

func NewTrdTx(pkt *Trd, rx TrdRx, tier ...int) (r *TrdTx) {
	r = &TrdTx{}
	r.Pkt = pkt
	r.Rx = rx
	if len(tier) > 0 {
		r.tier = tier[0]
	}
	return r
}
func (x *Trd) OpnMid() flt.Flt  { return x.OpnBid + (x.OpnAsk-x.OpnBid)*0.5 }
func (x *Trd) ClsMid() flt.Flt  { return x.ClsBid + (x.ClsAsk-x.ClsBid)*0.5 }
func (x *TrdTx) Act()           { x.ret = x.Rx(x.Pkt) }
func (x *TrdTx) Ret() []sys.Act { return x.ret }
func (x *TrdTx) Tier() int      { return x.tier }
func (x *TrdTx) DecTier()       { x.tier-- }
func (x *Trd) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *Trd) StrWrt(b *strings.Builder) string {
	b.WriteString("ana.trd(")
	b.WriteString("opnTme:")
	x.OpnTme.StrWrt(b)
	b.WriteString(" clsTme:")
	x.ClsTme.StrWrt(b)
	b.WriteString(" opnBid:")
	x.OpnBid.StrWrt(b)
	b.WriteString(" clsBid:")
	x.ClsBid.StrWrt(b)
	b.WriteString(" opnAsk:")
	x.OpnAsk.StrWrt(b)
	b.WriteString(" clsAsk:")
	x.ClsAsk.StrWrt(b)
	b.WriteString(" opnSpd:")
	x.OpnSpd.StrWrt(b)
	b.WriteString(" clsSpd:")
	x.ClsSpd.StrWrt(b)
	b.WriteString(" clsRsn:")
	x.ClsRsn.StrWrt(b)
	b.WriteString(" pip:")
	x.Pip.StrWrt(b)
	b.WriteString(" dur:")
	x.Dur.StrWrt(b)
	b.WriteString(" isLong:")
	x.IsLong.StrWrt(b)
	b.WriteString(" pnlPct:")
	x.PnlPct.StrWrt(b)
	b.WriteString(" pnlPctPredict:")
	x.PnlPctPredict.StrWrt(b)
	b.WriteString(" pnlUsd:")
	x.PnlUsd.StrWrt(b)
	b.WriteString(" pnlGrsUsd:")
	x.PnlGrsUsd.StrWrt(b)
	b.WriteString(" cstComUsd:")
	x.CstComUsd.StrWrt(b)
	b.WriteString(" cstClsSpdUsd:")
	x.CstClsSpdUsd.StrWrt(b)
	b.WriteString(" cstOpnSpdUsd:")
	x.CstOpnSpdUsd.StrWrt(b)
	b.WriteString(" opnBalUsd:")
	x.OpnBalUsd.StrWrt(b)
	b.WriteString(" clsBalUsd:")
	x.ClsBalUsd.StrWrt(b)
	b.WriteString(" clsBalUsdAct:")
	x.ClsBalUsdAct.StrWrt(b)
	b.WriteString(" trdPct:")
	x.TrdPct.StrWrt(b)
	b.WriteString(" mrgnRtio:")
	x.MrgnRtio.StrWrt(b)
	b.WriteString(" units:")
	x.Units.StrWrt(b)
	b.WriteString(" instr:")
	x.Instr.StrWrt(b)
	b.WriteString(" opnReq:")
	x.OpnReq.StrWrt(b)
	b.WriteString(" opnRes:")
	x.OpnRes.StrWrt(b)
	b.WriteString(" clsReq:")
	x.ClsReq.StrWrt(b)
	b.WriteString(" clsRes:")
	x.ClsRes.StrWrt(b)
	b.WriteRune(')')
	return b.String()
}
