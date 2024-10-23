package hst_test

// func TestHstPort2(t *testing.T) {
// 	ap := app.New(tst.Cfg)
// 	ap.Oan.CloneInstrs(tst.Instrs)
// 	defer ap.Cls()

// 	hstPort := hst.NewPort(
// 		hst.Oan().EurUsd().I(5).Bid().Lst().InrLss(1).Shrt(1.0, 1.0, tme.H1, hst.Oan().EurUsd()),
// 		hst.Oan().EurUsd().I(5).Bid().Lst().InrGtr(1).Long(1.0, 1.0, tme.H1, hst.Oan().EurUsd()),
// 	)
// 	hstPrfm := hstPort.Prfm().Bse()
// 	pth := "hst.newPort(hst.oan().eurUsd().i(5s).bid().lst().inrLss(1).shrt(1.0 1.0 1h hst.oan().eurUsd()) hst.oan().eurUsd().i(5s).bid().lst().inrGtr(1).long(1.0 1.0 1h hst.oan().eurUsd()))"
// 	tst.UntNotZero(t, hstPort.Bse().Trds.Cnt(), "Trds.Cnt")
// 	tst.FltNotZero(t, hstPrfm.PipAvg, "Prfm.PipAvg")
// 	tst.StringEql(t, pth, hstPort.String(), "Pth") // Pth after Prfm
// }

// func TestHstPortActr(t *testing.T) {
// 	ap := app.New(tst.Cfg)
// 	ap.Oan.CloneInstrs(tst.Instrs)
// 	defer ap.Cls()

// 	txt := "hst.newPort(hst.oan().eurUsd().i(5s).bid().lst().inrLss(1).shrt(1.0 1.0 1h hst.oan().eurUsd()))"
// 	var actr act.Actr
// 	actr.RunIfc(txt)
// }
