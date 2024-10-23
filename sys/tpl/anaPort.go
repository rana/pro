package tpl

import (
	"sys/k"
	"sys/tpl/atr"
)

type (
	FleAnaPort struct {
		FleBse
	}
)

func (x *DirAna) NewPort() (r *FleAnaPort) {
	r = &FleAnaPort{}
	x.Port = r
	r.Name = k.Port
	r.Pkg = x.Pkg
	r.StructPtr(r.Name, atr.TypAnaPort) // TODO: NO SCP/ACT
	r.AddFle(r)
	return r
}

func (x *FleAnaPort) InitFld(s *Struct) {
	s.Fld("BalFstUsd", _sys.Bsc.Flt)
	s.Fld("BalLstUsd", _sys.Bsc.Flt)
	s.Fld("TrdPct", _sys.Bsc.Flt)
	s.Fld("Trds", _sys.Ana.Trd.arr).Atr = atr.TstZeroSkp | atr.Arr
}
func (x *FleAnaPort) InitTypFn() {
	// MOVED TO HAND WRITTEN
	// x.CalcOpn()
	// x.CalcCls()
	// x.CalcPrfm()
}

// func (x *FleAnaPort) CalcOpn() (r *TypFn) {
// 	x.Import(_sys.Bsc.Int)
// 	r = x.TypFna(k.CalcOpn, atr.None)
// 	r.InPrm(_sys.Ana.Trd, "t")
// 	r.InPrm(_sys.Ana.Instr, "i")
// 	r.Add("t.Instr = i.Name")
// 	r.Add("t.MrgnRtio = i.MrgnRtio")
// 	r.Add("t.OpnBalUsd = x.BalLstUsd")
// 	r.Add("opnBalUsedUsd := t.OpnBalUsd * x.TrdPct // trdPct is appetite for risk")
// 	r.Add("opnBalMrgnUsd := opnBalUsedUsd * t.MrgnRtio")
// 	r.Add("if t.IsLong { // Ask: the price for me to buy at")
// 	r.Add("t.Units = flt.Flt(int.Int(opnBalMrgnUsd / t.OpnAsk))")
// 	r.Add("} else { // Bid: the price for me to sell at")
// 	r.Add("t.Units = flt.Flt(int.Int(opnBalMrgnUsd / t.OpnBid))")
// 	r.Add("}")
// 	return r
// }
// func (x *FleAnaPort) CalcCls() (r *TypFn) {
// 	r = x.TypFna(k.CalcCls, atr.None)
// 	r.InPrm(_sys.Ana.Trd, "t")
// 	r.InPrm(_sys.Ana.Instr, "i")
// 	r.Add("opnBalUsedUsd := t.OpnBalUsd * x.TrdPct")
// 	r.Add("opnBalUnusedUsd := t.OpnBalUsd - opnBalUsedUsd")
// 	r.Add("opnBalMrgnUsd := opnBalUsedUsd * t.MrgnRtio")
// 	r.Add("var clsBalMrgnUsd flt.Flt")
// 	r.Add("if t.IsLong { // Ask: the price for me to buy at")
// 	r.Add("clsBalMrgnUsd = t.ClsBid * t.Units")
// 	r.Add("t.PnlGrsUsd = clsBalMrgnUsd - opnBalMrgnUsd")
// 	r.Add("t.Pip = t.ClsBid.Sub(t.OpnAsk).Div(i.Pip).Trnc(2)")
// 	r.Add("} else { // Bid: the price for me to sell at")
// 	r.Add("clsBalMrgnUsd = t.ClsAsk * t.Units")
// 	r.Add("t.PnlGrsUsd = opnBalMrgnUsd - clsBalMrgnUsd")
// 	r.Add("t.Pip = t.OpnBid.Sub(t.ClsAsk).Div(i.Pip).Trnc(2)")
// 	r.Add("}")
// 	r.Add("clsBalUsedUsd := clsBalMrgnUsd / t.MrgnRtio")
// 	r.Add("t.CstComUsd = (Cfg.Oan.ComPerUnitUsd * 2) * t.Units     // commission cost for buy and sell")
// 	r.Add("t.CstOpnSpdUsd = ((t.OpnAsk - t.OpnBid) * .5) * t.Units // hlf spd opn cst")
// 	r.Add("t.CstClsSpdUsd = ((t.ClsAsk - t.ClsBid) * .5) * t.Units // hlf spd cls cst")
// 	r.Add("t.ClsBalUsd = opnBalUnusedUsd + clsBalUsedUsd + t.PnlGrsUsd - t.CstComUsd")
// 	r.Add("t.PnlUsd = t.ClsBalUsd.Sub(t.OpnBalUsd).Trnc(2)")
// 	r.Add("t.PnlPct = t.PnlUsd.Div(t.OpnBalUsd).Mul(100).Trnc(1)")
// 	r.Add("x.BalLstUsd = t.ClsBalUsd // update portfolio balance")
// 	return r
// }
// func (x *FleAnaPort) CalcPrfm() (r *TypFn) {
// 	r = x.TypFna(k.CalcPrfm, atr.None)
// 	r.InPrm(_sys.Ana.Trd.arr, "trds")
// 	r.InPrm(_sys.Bsc.Unt, "stgyCnt")
// 	r.InPrm(_sys.Bsc.Unt, "dayCnt")
// 	r.InPrm(_sys.Bsc.Flt, "losLimMax")
// 	r.InPrm(_sys.Bsc.Tme, "durLimMax")
// 	r.InPrm(String, "pth")
// 	r.OutPrm(_sys.Ana.Prfm, "r")
// 	r.Addf("r = %v{}", r.OutTyp().Adr(x))
// 	r.Add("durs := trds.Durs()")
// 	r.Add("pips := trds.Pips()")
// 	r.Add("r.TrdCnt = trds.Cnt()")
// 	r.Add("r.StgyCnt = stgyCnt")
// 	r.Add("r.DayCnt = dayCnt")
// 	r.Add("r.LosLimMax = losLimMax")
// 	r.Add("r.DurLimMax = durLimMax")
// 	r.Addf("r.Pth = %v(pth)", _sys.Bsc.Str.Ref(x))
// 	r.Add("r.PipAvg = pips.Sma().Trnc(1)")
// 	r.Add("r.PipMdn = pips.Mdn().Trnc(1)")
// 	r.Add("r.PipMin = pips.Min().Trnc(1)")
// 	r.Add("r.PipMax = pips.Max().Trnc(1)")
// 	r.Add("r.PipSum = pips.Sum().Trnc(1)")
// 	r.Add("r.DurAvg = durs.Sma()")
// 	r.Add("r.DurMdn = durs.Mdn()")
// 	r.Add("r.DurMin = durs.Min()")
// 	r.Add("r.DurMax = durs.Max()")
// 	r.Add("dayCntf, trdCntf, scsCnt := flt.Flt(dayCnt), flt.Flt(r.TrdCnt), pips.CntGtr(0.0) // CORRECT DEF OF SCS?")
// 	r.Add("r.OpnPerDay = trdCntf.Div(dayCntf).Trnc(1)")
// 	r.Add("r.PipPerDay = trds.Pips().Sum().Div(dayCntf).Trnc(1)")
// 	r.Add("r.ScsPerDay = scsCnt.Div(dayCntf).Trnc(1)")
// 	r.Add("r.ScsPct = scsCnt.Div(trdCntf).Mul(100).Trnc(1)")
// 	r.Add("for _, trd := range *trds {")
// 	r.Add("r.CstComUsd += trd.CstComUsd")
// 	r.Add("r.CstSpdUsd += trd.CstOpnSpdUsd + trd.CstClsSpdUsd")
// 	r.Add("}")
// 	r.Add("r.TrdPct = x.TrdPct")
// 	r.Add("r.BalFstUsd = x.BalFstUsd.Trnc(0)")
// 	r.Add("r.BalLstUsd = x.BalLstUsd.Trnc(0)")
// 	r.Add("r.PnlUsd = r.BalLstUsd.Sub(r.BalFstUsd).Trnc(0)")
// 	r.Add("r.PnlPct = r.PnlUsd.Div(r.BalFstUsd).Mul(100).Trnc(1)")
// 	r.Add("r.UsdPerDay = r.PnlUsd.Div(dayCntf).Trnc(0)")
// 	r.Add("r.CstTotUsd = r.CstSpdUsd.Add(r.CstComUsd).Trnc(0)")
// 	r.Add("r.CstSpdUsd = r.CstSpdUsd.Trnc(0)")
// 	r.Add("r.CstComUsd = r.CstComUsd.Trnc(0)")
// 	r.Add("return r")
// 	return r
// }
