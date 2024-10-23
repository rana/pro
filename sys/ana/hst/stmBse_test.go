package hst_test

// import (
// 	"sys/ana/vis/pen"
// 	"sys/ana/vis/plt"
// 	"sys/app"
// 	"sys/tst"
// 	"testing"
// )

// func TestHstStmBseAggAlmaTypFn(t *testing.T) {
// 	ap := app.New(tst.Cfg)
// 	ap.Oan.CloneInstrs(tst.Instrs)
// 	defer ap.Cls()

// 	pltRng := plt.NewStm()
// 	pltRng.Stm(pen.Purple400, ap.HstOan.EurUsd().S(5).Asks().Lst())
// 	pltRng.Stm(pen.Indigo400, ap.HstOan.EurUsd().S(5).Asks().Lst().AggSma(20))
// 	pltRng.Stm(pen.Orange700, ap.HstOan.EurUsd().S(5).Asks().Lst().AggAlma(9, 6, 0.85))

// 	// imgOpn := true
// 	// tst.ImgOpn = &imgOpn

// 	tst.DrawPlt(pltRng)
// }

// func TestHstStmBseAggSarTypFn(t *testing.T) {
// 	ap := app.New(tst.Cfg)
// 	ap.Oan.CloneInstrs(tst.Instrs)
// 	defer ap.Cls()

// 	pltRng := plt.NewStm()
// 	pltRng.Stm(pen.Purple700, ap.HstOan.EurUsd().S(5).Asks().Lst())
// 	pltRng.Stm(pen.White, ap.HstOan.EurUsd().S(5).Asks().Sar(0.02, 0.2))

// 	// imgOpn := true
// 	// tst.ImgOpn = &imgOpn

// 	tst.DrawPlt(pltRng)
// }

// func TestHstStmBseInrSlpTypFn(t *testing.T) {
// 	ap := app.New(tst.Cfg)
// 	ap.Oan.CloneInstrs(tst.Instrs)
// 	defer ap.Cls()

// 	eurUsd := ap.HstOan.EurUsd()
// 	lst := eurUsd.S(5).Asks().Lst()
// 	sma20 := lst.AggSma(20)
// 	slp := sma20.SclMul(100000).InrSlp(2)
// 	slpSma20 := slp.AggAlma(9, 6, 0.85) //.AggSma(20)
// 	// sys.Log("slp.Bse().Vals", slp.Bse().Vals)

// 	pltRng0 := plt.NewStm()
// 	pltRng0.Stm(pen.Purple400, lst)
// 	pltRng0.Stm(pen.Orange700, sma20)
// 	pltRng0.VrtScl(0.5)

// 	pltRng1 := plt.NewStm()
// 	pltRng1.Stm(pen.Indigo400, slp)
// 	pltRng1.Stm(pen.Lime400, slpSma20)
// 	pltRng1.VrtScl(0.5)
// 	pltRng1.HrzLn(pen.Grey400, 1, 0)

// 	pltVrt := plt.NewVrt(pltRng0, pltRng1)

// 	// imgOpn := true
// 	// tst.ImgOpn = &imgOpn

// 	tst.DrawPlt(pltVrt)
// }
