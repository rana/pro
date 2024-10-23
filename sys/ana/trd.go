package ana

import (
	"sys/bsc/flt"
)

func (x *Trd) CalcPip(instrPip flt.Flt) {
	// Bid: the price for me to sell at
	// Ask: the price for me to buy at
	if x.IsLong {
		x.Pip = x.ClsBid.Sub(x.OpnAsk).Div(instrPip).Trnc(2)
	} else {
		x.Pip = x.OpnBid.Sub(x.ClsAsk).Div(instrPip).Trnc(2)
	}
}

func (x *Trd) CalcPnl(i *Instr, opnBalUsd, trdPct flt.Flt) {
	// commission calc: https://www.oanda.com/register/docs/oc/price-sheet.pdf
	x.Instr = i.Name
	// x.OpnBalUsd = p.BalLstUsd
	// x.TrdPct = p.TrdPct
	x.OpnBalUsd = opnBalUsd
	x.TrdPct = trdPct
	x.MrgnRtio = i.MrgnRtio
	opnBalUsedUsd := x.OpnBalUsd * x.TrdPct // trdPct is appetite for risk
	opnBalUnusedUsd := x.OpnBalUsd - opnBalUsedUsd
	opnBalMrgnUsd := opnBalUsedUsd * x.MrgnRtio
	var clsBalMrgnUsd flt.Flt
	if x.IsLong { // Ask: the price for me to buy at
		x.Units = flt.Flt(int(opnBalMrgnUsd / x.OpnAsk))
		clsBalMrgnUsd = x.ClsBid * x.Units
		x.PnlGrsUsd = clsBalMrgnUsd - opnBalMrgnUsd
	} else { // Bid: the price for me to sell at
		x.Units = flt.Flt(int(opnBalMrgnUsd / x.OpnBid))
		clsBalMrgnUsd = x.ClsAsk * x.Units
		x.PnlGrsUsd = opnBalMrgnUsd - clsBalMrgnUsd
	}
	clsBalUsedUsd := clsBalMrgnUsd / x.MrgnRtio
	x.CstComUsd = (Cfg.Oan.ComPerUnitUsd * 2) * x.Units     // commission cost for buy and sell
	x.CstOpnSpdUsd = ((x.OpnAsk - x.OpnBid) * .5) * x.Units // hlf spd opn cst
	x.CstClsSpdUsd = ((x.ClsAsk - x.ClsBid) * .5) * x.Units // hlf spd cls cst
	x.ClsBalUsd = opnBalUnusedUsd + clsBalUsedUsd + x.PnlGrsUsd - x.CstComUsd
	x.PnlUsd = x.ClsBalUsd.Sub(x.OpnBalUsd).Trnc(2)
	x.PnlPct = x.PnlUsd.Div(x.OpnBalUsd).Mul(100).Trnc(1)
	// p.BalLstUsd = x.ClsBalUsd
}
