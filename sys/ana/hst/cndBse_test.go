package hst_test

// import (
// 	"sys/ana/vis/pen"
// 	"sys/ana/vis/plt"
// 	"sys/app"
// 	"sys/bsc/tme"
// 	"sys/tst"
// 	"testing"
// )

// func TestHstCndBseAtTypFn(t *testing.T) {
// 	ap := app.New(tst.Cfg)
// 	ap.Oan.CloneInstrs(tst.Instrs)
// 	defer ap.Cls()

// 	i := ap.HstOan.EurUsd()
// 	rng := tme.NewRng(
// 		i.I.HstStm.Tmes.At(i.I.HstStm.Tmes.MdlIdx()/2),
// 		i.I.HstStm.Tmes.Mdl())
// 	// prs.TmeTxt("2018-5-20"),
// 	// prs.TmeTxt("2018-5-20/30m")
// 	eurUsdBse := ap.HstOan.EurUsd(rng).S(1).Asks()
// 	eurUsd := eurUsdBse.Lst()
// 	sar := eurUsdBse.Sar(0.02, 0.2)
// 	sarTop := sar.OtrGtr(0, eurUsd)
// 	sarBtm := sar.OtrLss(0, eurUsd)
// 	sarDwn := sarTop.Seq(tme.Tme(1), sarBtm)
// 	alma := eurUsd.AggAlma(20, 6, 0.85)

// 	pltRng := plt.NewStm()
// 	pltRng.Stm(pen.Purple400, eurUsd)
// 	pltRng.Stm(pen.White, sar)
// 	pltRng.Stm(pen.Orange700, alma)
// 	pltRng.Cnd(pen.Teal700, sarDwn)

// 	pltRng.Scl(4.0)

// 	// imgOpn := true
// 	// tst.ImgOpn = &imgOpn

// 	tst.DrawPlt(pltRng)
// }
