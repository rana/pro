package plt_test

import (
	"sys/ana/hst"
	"sys/ana/vis/clr"
	"sys/ana/vis/pen"
	"sys/ana/vis/plt"
	"sys/app"
	"sys/bsc/tme"
	"sys/tst"
	"testing"
)

func TestPltStmStm1(t *testing.T) {
	ap := app.New(tst.Cfg)
	ap.Oan.CloneInstrs(tst.Instrs)
	defer ap.Cls()

	pltStm := plt.NewStm()
	pltStm.Stm(pen.LimeA400, hst.Oan().GbpUsd().I(1).Ask().Lst())
	pltStm.X().Vis(true)

	tst.DrawPlt(pltStm)
}

func TestPltStmStm3(t *testing.T) {
	ap := app.New(tst.Cfg)
	ap.Oan.CloneInstrs(tst.Instrs)
	defer ap.Cls()

	pltStm := plt.NewStm()
	pltStm.X().Vis(true)
	pltStm.Stm(pen.LimeA400, hst.Oan().EurUsd().I(1).Ask().Lst())
	pltStm.Stm(pen.IndigoA400, hst.Oan().AudUsd().I(1).Ask().Lst())
	pltStm.Stm(pen.RedA400, hst.Oan().NzdUsd().I(1).Ask().Lst())

	tst.DrawPlt(pltStm)
}

func TestPltStmWrsi(t *testing.T) {
	ap := app.New(tst.Cfg)
	ap.Oan.CloneInstrs(tst.Instrs)
	defer ap.Cls()

	ask := hst.Oan().EurUsd().I(1).Ask()
	pltStm := plt.NewStm()
	pltStm.X().Vis(true)
	pltStm.Stm(pen.LimeA400, ask.Wrsi())

	tst.DrawPlt(pltStm)
}

func TestPltStmEmaMacd(t *testing.T) {
	ap := app.New(tst.Cfg)
	ap.Oan.CloneInstrs(tst.Instrs)
	defer ap.Cls()

	// https://tradingsim.com/blog/macd/
	// MACD line: 12-period EMA - 26 period EMA (qck)
	// trigger line: 9-day EMA of MACD line (slw)
	pltStm := plt.NewStm()
	ema12 := hst.Oan().EurUsd().I(tme.S30 * 12).Ask().Ema()
	ema26 := hst.Oan().EurUsd().I(tme.S30 * 26).Ask().Ema()
	macd := ema12.OtrSub(0, ema26)
	macdTrg := macd.AggEma(9)

	pltStm.Stm(pen.IndigoA400, macd)
	pltStm.Stm(pen.RedA400, macdTrg)

	tst.DrawPlt(pltStm)
}

func TestPltStmStmBnd(t *testing.T) {
	ap := app.New(tst.Cfg)
	ap.Oan.CloneInstrs(tst.Instrs)
	defer ap.Cls()

	eurUsd := hst.Oan().EurUsd().I(1).Ask().Lst()
	sma20 := eurUsd.AggSma(20)
	blngr := eurUsd.AggStd(20).SclMul(2.0)
	blngrBtm := sma20.OtrSub(0, blngr)
	blngrTop := sma20.OtrAdd(0, blngr)

	pltStm := plt.NewStm()
	pltStm.StmBnd(clr.Hex("#0f0f0f"), pen.DeepPurpleA400, blngrBtm, blngrTop)

	tst.DrawPlt(pltStm)
}

func TestPltStmStmBndWithStm(t *testing.T) {
	ap := app.New(tst.Cfg)
	ap.Oan.CloneInstrs(tst.Instrs)
	defer ap.Cls()

	eurUsd := hst.Oan().EurUsd().I(1).Ask().Lst()
	sma20 := eurUsd.AggSma(20)
	blngr := eurUsd.AggStd(20).SclMul(2.0)
	blngrBtm := sma20.OtrSub(0, blngr)
	blngrTop := sma20.OtrAdd(0, blngr)

	pltStm := plt.NewStm()
	pltStm.StmBnd(clr.Purple900.Opa(.1), pen.DeepPurpleA400, blngrBtm, blngrTop)
	pltStm.Stm(pen.White, eurUsd)

	tst.DrawPlt(pltStm)
}

func TestPltVrt(t *testing.T) {
	ap := app.New(tst.Cfg)
	ap.Oan.CloneInstrs(tst.Instrs)
	defer ap.Cls()

	pltStm1 := plt.NewStm()
	pltStm1.Stm(pen.LimeA400, hst.Oan().EurUsd().I(1).Ask().Lst())
	pltStm2 := plt.NewStm()
	pltStm2.Stm(pen.IndigoA400, hst.Oan().AudUsd().I(1).Ask().Lst())
	pltStm3 := plt.NewStm()
	pltStm3.Stm(pen.RedA400, hst.Oan().NzdUsd().I(1).Ask().Lst())
	pltVrt := plt.NewVrt(pltStm1, pltStm2, pltStm3)

	tst.DrawPlt(pltVrt)
}

func TestPltHrz(t *testing.T) {
	ap := app.New(tst.Cfg)
	ap.Oan.CloneInstrs(tst.Instrs)
	defer ap.Cls()

	pltStm1 := plt.NewStm()
	pltStm1.HrzScl(0.33)
	pltStm1.Stm(pen.LimeA400, hst.Oan().EurUsd().I(1).Ask().Lst())
	pltStm2 := plt.NewStm()
	pltStm2.HrzScl(0.33)
	pltStm2.Stm(pen.IndigoA400, hst.Oan().AudUsd().I(1).Ask().Lst())
	pltStm3 := plt.NewStm()
	pltStm3.HrzScl(0.33)
	pltStm3.Stm(pen.RedA400, hst.Oan().NzdUsd().I(1).Ask().Lst())
	pltHrz := plt.NewHrz(pltStm1, pltStm2, pltStm3)

	tst.DrawPlt(pltHrz)
}

func TestPltStmHrzLnWithStm(t *testing.T) {
	ap := app.New(tst.Cfg)
	ap.Oan.CloneInstrs(tst.Instrs)
	defer ap.Cls()

	// must have a stm or bnd to display ln
	stm := hst.Oan().EurUsd().I(1).Ask().Lst()
	pltStm := plt.NewStm()
	pltStm.Stm(pen.LimeA400, stm)
	pltStm.HrzLn(pen.Grey700, 1, stm.Bse().Vals.Mid())

	tst.DrawPlt(pltStm)
}

func TestPltStmHrzLnWithBnd(t *testing.T) {
	ap := app.New(tst.Cfg)
	ap.Oan.CloneInstrs(tst.Instrs)
	defer ap.Cls()

	// must have a stm or bnd to display ln
	eurUsd := hst.Oan().EurUsd().I(1).Ask().Lst()
	sma20 := eurUsd.AggSma(20)
	blngr := eurUsd.AggStd(20).SclMul(2.0)
	blngrBtm := sma20.OtrSub(0, blngr)
	blngrTop := sma20.OtrAdd(0, blngr)

	pltStm := plt.NewStm()
	pltStm.StmBnd(clr.Purple900.Opa(.1), pen.DeepPurpleA400, blngrBtm, blngrTop)
	pltStm.HrzLn(pen.Grey700, eurUsd.Bse().Vals.Mid())

	tst.DrawPlt(pltStm)
}

func TestPltStmVrtLnWithStm(t *testing.T) {
	ap := app.New(tst.Cfg)
	ap.Oan.CloneInstrs(tst.Instrs)
	defer ap.Cls()

	// must have a stm or bnd to display ln
	stm := hst.Oan().EurUsd().I(1).Ask().Lst()
	pltStm := plt.NewStm()
	pltStm.Stm(pen.LimeA400, stm)
	pltStm.VrtLn(pen.Grey700, stm.Bse().Tmes.Mid())

	tst.DrawPlt(pltStm)
}

func TestPltStmVrtLnWithBnd(t *testing.T) {
	ap := app.New(tst.Cfg)
	ap.Oan.CloneInstrs(tst.Instrs)
	defer ap.Cls()

	// must have a stm or bnd to display ln
	eurUsd := hst.Oan().EurUsd().I(1).Ask().Lst()
	sma20 := eurUsd.AggSma(20)
	blngr := eurUsd.AggStd(20).SclMul(2.0)
	blngrBtm := sma20.OtrSub(0, blngr)
	blngrTop := sma20.OtrAdd(0, blngr)

	pltStm := plt.NewStm()
	pltStm.StmBnd(clr.Purple900.Opa(.1), pen.DeepPurpleA400, blngrBtm, blngrTop)
	pltStm.VrtLn(pen.Grey700, eurUsd.Bse().Tmes.Mid())

	tst.DrawPlt(pltStm)
}

func TestPltStmHrzBndWithStm(t *testing.T) {
	ap := app.New(tst.Cfg)
	ap.Oan.CloneInstrs(tst.Instrs)
	defer ap.Cls()

	// must have a stm or bnd to display ln
	stm := hst.Oan().EurUsd().I(1).Ask().Lst()
	stmBse := stm.Bse()
	oct := stmBse.Vals.RngFul() / 8
	valA := stmBse.Vals.Mid() - oct
	valB := stmBse.Vals.Mid() + oct

	pltStm := plt.NewStm()
	pltStm.Stm(pen.LimeA400, stm)
	pltStm.HrzBnd(clr.Purple900.Opa(.1), pen.DeepPurpleA400, valA, valB)

	tst.DrawPlt(pltStm)
}

func TestPltStmHrzBndWithBnd(t *testing.T) {
	ap := app.New(tst.Cfg)
	ap.Oan.CloneInstrs(tst.Instrs)
	defer ap.Cls()

	// must have a stm or bnd to display ln
	eurUsd := hst.Oan().EurUsd().I(1).Ask().Lst()
	sma20 := eurUsd.AggSma(20)
	blngr := eurUsd.AggStd(20).SclMul(2.0)
	blngrBtm := sma20.OtrSub(0, blngr)
	blngrTop := sma20.OtrAdd(0, blngr)
	stmBse := eurUsd.Bse()
	oct := stmBse.Vals.RngFul() / 8
	valA := stmBse.Vals.Mid() - oct
	valB := stmBse.Vals.Mid() + oct

	pltStm := plt.NewStm()
	pltStm.StmBnd(clr.LimeA400.Opa(.1), pen.LimeA400, blngrBtm, blngrTop)
	pltStm.HrzBnd(clr.Purple900.Opa(.1), pen.DeepPurpleA400, valA, valB)

	tst.DrawPlt(pltStm)
}

func TestPltStmVrtBndWithStm(t *testing.T) {
	ap := app.New(tst.Cfg)
	ap.Oan.CloneInstrs(tst.Instrs)
	defer ap.Cls()

	// must have a stm or bnd to display ln
	stm := hst.Oan().EurUsd().I(1).Ask().Lst()
	stmBse := stm.Bse()
	oct := stmBse.Tmes.RngFul() / 8
	valA := stmBse.Tmes.Mid() - oct
	valB := stmBse.Tmes.Mid() + oct

	pltStm := plt.NewStm()
	pltStm.Stm(pen.LimeA400, stm)
	pltStm.VrtBnd(clr.Purple900.Opa(.1), pen.DeepPurpleA400, valA, valB)

	tst.DrawPlt(pltStm)
}

func TestPltStmVrtBndWithBnd(t *testing.T) {
	ap := app.New(tst.Cfg)
	ap.Oan.CloneInstrs(tst.Instrs)
	defer ap.Cls()

	// must have a stm or bnd to display ln
	eurUsd := hst.Oan().EurUsd().I(1).Ask().Lst()
	sma20 := eurUsd.AggSma(20)
	blngr := eurUsd.AggStd(20).SclMul(2.0)
	blngrBtm := sma20.OtrSub(0, blngr)
	blngrTop := sma20.OtrAdd(0, blngr)
	stmBse := eurUsd.Bse()
	oct := stmBse.Tmes.RngFul() / 8
	valA := stmBse.Tmes.Mid() - oct
	valB := stmBse.Tmes.Mid() + oct

	pltStm := plt.NewStm()
	pltStm.StmBnd(clr.LimeA400.Opa(.1), pen.LimeA400, blngrBtm, blngrTop)
	pltStm.VrtBnd(clr.Purple900.Opa(.1), pen.DeepPurpleA400, valA, valB)

	tst.DrawPlt(pltStm)
}

func TestPltStmCnd(t *testing.T) {
	ap := app.New(tst.Cfg)
	ap.Oan.CloneInstrs(tst.Instrs)
	defer ap.Cls()

	eurUsd := hst.Oan().EurUsd().I(1).Ask().Lst()
	sma20 := eurUsd.AggSma(20)
	blngr := eurUsd.AggStd(20).SclMul(2.0)
	blngrBtm := sma20.OtrSub(0, blngr)
	blngrTop := sma20.OtrAdd(0, blngr)

	eurUsdLss := eurUsd.OtrLss(0, blngrTop)
	eurUsdGtr := eurUsd.OtrGtr(0, blngrTop)
	blngrTopCrs := eurUsdLss.Seq(tme.S1, eurUsdGtr)

	pltStm := plt.NewStm()
	pltStm.StmBnd(clr.Hex("#0f0f0f"), pen.DeepPurpleA400, blngrBtm, blngrTop)
	pltStm.Stm(pen.White, eurUsd)
	pltStm.Cnd(pen.OrangeA400, blngrTopCrs)

	tst.DrawPlt(pltStm)
}

func TestPltFltsSctr1(t *testing.T) {
	ap := app.New(tst.Cfg)
	ap.Oan.CloneInstrs(tst.Instrs)
	defer ap.Cls()

	// mdl := hst.Oan().EurUsd().I.HstStm.Tmes.Mdl()

	eurUsd := hst.Oan().EurUsd().I(1).Ask().Lst()
	sma20 := eurUsd.AggSma(20)
	blngr := eurUsd.AggStd(20).SclMul(2.0)
	blngrBtm := sma20.OtrSub(0, blngr)
	blngrTop := sma20.OtrAdd(0, blngr)

	eurUsdLss := eurUsd.OtrLss(0, blngrTop)
	eurUsdGtr := eurUsd.OtrGtr(0, blngrTop)
	blngrTopCrs := eurUsdLss.Seq(tme.S1, eurUsdGtr)

	sma20Slp := sma20.InrSlp(1).SclMul(hst.Oan().EurUsd().Bse().Ana.PipetteScl())
	sma20SlpBlngrTopCrs := sma20Slp.At(blngrTopCrs.Bse().Tmes)

	pltVrt := plt.NewVrt()

	pltStm := plt.NewStm()
	pltStm.StmBnd(clr.Hex("#0f0f0f"), pen.DeepPurpleA400, blngrBtm, blngrTop)
	pltStm.Stm(pen.White, eurUsd)
	pltStm.Cnd(pen.OrangeA400, blngrTopCrs)
	pltStm.VrtScl(0.5)
	pltVrt.Plt(pltStm)

	pltSctr := plt.NewFltsSctr()
	pltSctr.Flts(clr.Indigo400.Opa(0.3), sma20SlpBlngrTopCrs)
	pltSctr.VrtScl(0.5)
	pltVrt.Plt(pltSctr)

	tst.DrawPlt(pltVrt)
}

func TestPltFltsSctr3(t *testing.T) {
	ap := app.New(tst.Cfg)
	ap.Oan.CloneInstrs(tst.Instrs)
	defer ap.Cls()

	eurUsd := hst.Oan().EurUsd().I(1).Ask().Lst()
	sma5, sma10, sma20 := eurUsd.AggSma(5), eurUsd.AggSma(10), eurUsd.AggSma(20)
	blngr := eurUsd.AggStd(20).SclMul(2.0)
	blngrBtm := sma20.OtrSub(0, blngr)
	blngrTop := sma20.OtrAdd(0, blngr)

	eurUsdLss := eurUsd.OtrLss(0, blngrTop)
	eurUsdGtr := eurUsd.OtrGtr(0, blngrTop)
	blngrTopCrs := eurUsdLss.Seq(tme.S1, eurUsdGtr)

	sma5BlngrTopCrs := sma5.At(blngrTopCrs.Bse().Tmes)
	sma10BlngrTopCrs := sma10.At(blngrTopCrs.Bse().Tmes)
	sma20BlngrTopCrs := sma20.At(blngrTopCrs.Bse().Tmes)

	pltVrt := plt.NewVrt()

	pltStm := plt.NewStm()
	pltStm.StmBnd(clr.Hex("#0f0f0f"), pen.DeepPurpleA400, blngrBtm, blngrTop)
	pltStm.Stm(pen.White, eurUsd)
	pltStm.Cnd(pen.OrangeA400, blngrTopCrs)
	pltStm.VrtScl(0.5)
	pltVrt.Plt(pltStm)

	pltSctr := plt.NewFltsSctr()
	pltSctr.Flts(clr.Indigo400.Opa(0.1), sma5BlngrTopCrs, sma10BlngrTopCrs, sma20BlngrTopCrs)
	pltSctr.VrtScl(0.5)
	pltVrt.Plt(pltSctr)

	tst.DrawPlt(pltVrt)
}

func TestPltFltsCntrDist1(t *testing.T) {
	ap := app.New(tst.Cfg)
	ap.Oan.CloneInstrs(tst.Instrs)
	defer ap.Cls()

	eurUsd := hst.Oan().EurUsd().I(1).Ask().Lst()
	sma20 := eurUsd.AggSma(20)
	blngr := eurUsd.AggStd(20).SclMul(2.0)
	blngrBtm := sma20.OtrSub(0, blngr)
	blngrTop := sma20.OtrAdd(0, blngr)

	eurUsdLss := eurUsd.OtrLss(0, blngrTop)
	eurUsdGtr := eurUsd.OtrGtr(0, blngrTop)
	blngrTopCrs := eurUsdLss.Seq(tme.S1, eurUsdGtr)

	sma20Slp := sma20.InrSlp(1).SclMul(hst.Oan().EurUsd().Bse().Ana.PipetteScl())
	sma20SlpBlngrTopCrs := sma20Slp.At(blngrTopCrs.Bse().Tmes)

	pltVrt := plt.NewVrt()

	pltStm := plt.NewStm()
	pltStm.StmBnd(clr.Hex("#0f0f0f"), pen.DeepPurpleA400, blngrBtm, blngrTop)
	pltStm.Stm(pen.White, eurUsd)
	pltStm.Cnd(pen.OrangeA400, blngrTopCrs)
	pltStm.VrtScl(0.5)
	pltVrt.Plt(pltStm)

	pltCntrDist := plt.NewFltsSctrDist()
	pltCntrDist.Flts(clr.Indigo400.Opa(0.3), 2, sma20SlpBlngrTopCrs)
	pltCntrDist.VrtScl(0.5)
	pltVrt.Plt(pltCntrDist)

	tst.DrawPlt(pltVrt)
}

func TestPltFltsCntrDist3(t *testing.T) {
	ap := app.New(tst.Cfg)
	ap.Oan.CloneInstrs(tst.Instrs)
	defer ap.Cls()

	eurUsd := hst.Oan().EurUsd().I(1).Ask().Lst()
	sma20 := eurUsd.AggSma(20)
	blngr := eurUsd.AggStd(20).SclMul(2.0)
	blngrBtm := sma20.OtrSub(0, blngr)
	blngrTop := sma20.OtrAdd(0, blngr)

	eurUsdLss := eurUsd.OtrLss(0, blngrTop)
	eurUsdGtr := eurUsd.OtrGtr(0, blngrTop)
	blngrTopCrs := eurUsdLss.Seq(tme.S1, eurUsdGtr)

	pipetteScl := hst.Oan().EurUsd().Bse().Ana.PipetteScl()
	sma5Slp := eurUsd.AggSma(5).InrSlp(1).SclMul(pipetteScl)
	sma10Slp := eurUsd.AggSma(10).InrSlp(1).SclMul(pipetteScl)
	sma20Slp := sma20.InrSlp(1).SclMul(pipetteScl)
	sma5SlpBlngrTopCrs := sma5Slp.At(blngrTopCrs.Bse().Tmes)
	sma10SlpBlngrTopCrs := sma10Slp.At(blngrTopCrs.Bse().Tmes)
	sma20SlpBlngrTopCrs := sma20Slp.At(blngrTopCrs.Bse().Tmes)

	pltVrt := plt.NewVrt()

	pltStm := plt.NewStm()
	pltStm.StmBnd(clr.Hex("#0f0f0f"), pen.DeepPurpleA400, blngrBtm, blngrTop)
	pltStm.Stm(pen.White, eurUsd)
	pltStm.Cnd(pen.OrangeA400, blngrTopCrs)
	pltStm.VrtScl(0.5)
	pltVrt.Plt(pltStm)

	pltCntrDist := plt.NewFltsSctrDist()
	pltCntrDist.Flts(clr.Indigo400.Opa(0.3), 2, sma5SlpBlngrTopCrs, sma10SlpBlngrTopCrs, sma20SlpBlngrTopCrs)
	pltCntrDist.VrtScl(0.5)
	pltVrt.Plt(pltCntrDist)

	tst.DrawPlt(pltVrt)
}

func TestPltFltsCntrDistMrg3(t *testing.T) {
	ap := app.New(tst.Cfg)
	ap.Oan.CloneInstrs(tst.Instrs)
	defer ap.Cls()

	eurUsd := hst.Oan().EurUsd().I(1).Ask().Lst()
	sma20 := eurUsd.AggSma(20)
	blngr := eurUsd.AggStd(20).SclMul(2.0)
	blngrBtm := sma20.OtrSub(0, blngr)
	blngrTop := sma20.OtrAdd(0, blngr)

	eurUsdLss := eurUsd.OtrLss(0, blngrTop)
	eurUsdGtr := eurUsd.OtrGtr(0, blngrTop)
	blngrTopCrs := eurUsdLss.Seq(tme.S1, eurUsdGtr)

	pipetteScl := hst.Oan().EurUsd().Bse().Ana.PipetteScl()
	sma5Slp := eurUsd.AggSma(5).InrSlp(1).SclMul(pipetteScl)
	sma10Slp := eurUsd.AggSma(10).InrSlp(1).SclMul(pipetteScl)
	sma20Slp := sma20.InrSlp(1).SclMul(pipetteScl)
	sma5SlpBlngrTopCrs := sma5Slp.At(blngrTopCrs.Bse().Tmes)
	sma10SlpBlngrTopCrs := sma10Slp.At(blngrTopCrs.Bse().Tmes)
	sma20SlpBlngrTopCrs := sma20Slp.At(blngrTopCrs.Bse().Tmes)

	pltVrt := plt.NewVrt()

	pltStm := plt.NewStm()
	pltStm.StmBnd(clr.Hex("#0f0f0f"), pen.DeepPurpleA400, blngrBtm, blngrTop)
	pltStm.Stm(pen.White, eurUsd)
	pltStm.Cnd(pen.OrangeA400, blngrTopCrs)
	pltStm.VrtScl(0.5)
	pltVrt.Plt(pltStm)

	pltCntrDist := plt.NewFltsSctrDist()
	pltCntrDist.Flts(clr.Indigo400.Opa(0.3), 1, sma5SlpBlngrTopCrs.Mrg(sma10SlpBlngrTopCrs, sma20SlpBlngrTopCrs))
	pltCntrDist.VrtScl(0.5)
	pltVrt.Plt(pltCntrDist)

	tst.DrawPlt(pltVrt)
}

// TODO: REPAIR
// func TestPltPrcp(t *testing.T) {
// 	ap := app.New(tst.Cfg)
// 	ap.Oan.CloneInstrs(tst.Instrs)
// 	defer ap.Cls()

// 	lst := hst.Oan().EurUsd().I(5).Bid().Lst()
// 	prcp := hst.NewPrcp()
// 	prcp.Stm(lst.AggSma(4))
// 	prcp.Stm(lst.AggSma(8))
// 	prcp.Stm(lst.AggSma(16))
// 	prcp.Stm(lst.AggWma(4))
// 	prcp.Stm(lst.AggWma(8))
// 	prcp.Stm(lst.AggWma(16))

// 	pltPrcp := plt.NewPrcp(prcp)
// 	pltPrcp.Pens(pen.NewPens(pen.Indigo500, pen.LightBlue500, pen.Green500, pen.Yellow500, pen.Orange500, pen.Red500))
// 	tst.DrawPlt(pltPrcp)
// }

// func TestPltStmSplt(t *testing.T) {
// 	ap := app.New(tst.Cfg)
// 	ap.Oan.CloneInstrs(tst.Instrs)
// 	defer ap.Cls()

// 	instr := hst.Oan().EurUsd()
// 	lst := instr.I(5).Bid().Lst()
// 	sma := lst.AggSmas(unts.New(4, 8, 16))
// 	splt := lst.InrGtr(1).Long(2.0, 2.0, 60*60, instr).Port().Splt(0.0)
// 	smaSplt := sma.Splt(splt.Btm.OpnTmes(), splt.Top.OpnTmes())

// 	pltStmSplt := plt.NewStmSplt(*smaSplt.StmSplts...)
// 	tst.DrawPlt(pltStmSplt)
// }

// func TestPltPrcpSplt(t *testing.T) {
// 	ap := app.New(tst.Cfg)
// 	ap.Oan.CloneInstrs(tst.Instrs)
// 	defer ap.Cls()

// 	instr := hst.Oan().EurUsd()
// 	lst := instr.I(5).Bid().Lst()
// 	splt := lst.InrGtr(1).Long(2.0, 2.0, 60*60, instr).Port().Splt(0.0)
// 	prcp := hst.NewPrcp()
// 	prcp.Stm(lst.AggSma(4))
// 	prcp.Stm(lst.AggSma(8))
// 	prcp.Stm(lst.AggSma(16))
// 	prcp.Stm(lst.AggWma(4))
// 	prcp.Stm(lst.AggWma(8))
// 	prcp.Stm(lst.AggWma(16))
// 	prcp.Stm(lst.AggRsi(4))
// 	prcp.Stm(lst.AggRsi(8))
// 	prcp.Stm(lst.AggRsi(16))
// 	prcpSplt := prcp.Splt(splt)

// 	pltPrcpSplt := plt.NewPrcpSplt(prcpSplt)
// 	pltPrcpSplt.VrtScl(0.5)
// 	tst.DrawPlt(pltPrcpSplt)
// }

// func TestPltStmStgy(t *testing.T) {
// 	ap := app.New(tst.Cfg)
// 	ap.Oan.CloneInstrs(tst.Instrs)
// 	defer ap.Cls()

// 	instr := hst.Oan().EurUsd()
// 	s1Lst := instr.I(tme.S1).Ask().Lst()
// 	m1Sma := instr.I(tme.M1).Ask().Sma()
// 	opnCnd := s1Lst.OtrLss(0, m1Sma).Seq(tme.S1, s1Lst.OtrGtr(0, m1Sma))

// 	pltStgy := plt.NewStgy()
// 	pltStgy.Stgy(opnCnd.Long(2.0, 2.0, tme.H1, instr))
// 	pltStgy.Stm(hst.Oan().EurUsd().I(1).Ask().Lst())

// 	tst.DrawPlt(pltStgy)
// }
