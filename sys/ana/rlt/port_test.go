package rlt_test

// func TestRltPort2EurUsd(t *testing.T) {
// 	ap := app.New(tst.Cfg)
// 	ap.Oan.CloneInstrs(tst.Instrs)
// 	defer ap.Cls()
// 	instr := rlt.Oan().EurUsd()
// 	rltPort := rlt.NewPort(
// 		instr.I(5).Bid().Lst().InrLss(1).Shrt(1.0, 1.0, tme.H1, instr),
// 		instr.I(5).Bid().Lst().InrGtr(1).Long(1.0, 1.0, tme.H1, instr),
// 	)
// 	hstPort := ap.Actr.RunHst(rltPort.String())[0].(*hst.Port)
// 	hstPrfm := hstPort.Prfm() // call before WaitFor (for CalcTrds)
// 	mnr := tst.NewStgyMnr(ap)
// 	rltPort.Sub(mnr.Rx, mnr.Id)
// 	mnr.Start(hstPort.Stgys.At(0).I()) // send on one for all
// 	mnr.WaitFor(hstPort.Trds.Cnt())
// 	if hstPort.Trds != nil {
// 		tst.AnaTrdsEql(t, hstPort.Trds, mnr.Trds, "Trds")
// 		tst.AnaPortEql(t, &hstPort.Port, &rltPort.Port, "Port")
// 		tst.AnaPrfmEql(t, hstPrfm, rltPort.Prfm(), "Prfm")
// 	}
// }

// func TestRltPort2AudUsd(t *testing.T) {
// 	ap := app.New(tst.Cfg)
// 	ap.Oan.CloneInstrs(tst.Instrs)
// 	defer ap.Cls()
// 	instr := rlt.Oan().AudUsd()
// 	rltPort := rlt.NewPort(
// 		instr.I(5).Bid().Lst().InrLss(1).Shrt(1.0, 1.0, tme.H1, instr),
// 		instr.I(5).Bid().Lst().InrGtr(1).Long(1.0, 1.0, tme.H1, instr),
// 	)
// 	hstPort := ap.Actr.RunHst(rltPort.String())[0].(*hst.Port)
// 	hstPrfm := hstPort.Prfm() // call before WaitFor (for CalcTrds)
// 	// sys.Log("hstPort.Trds.Cnt()", hstPort.Trds.Cnt())
// 	mnr := tst.NewStgyMnr(ap)
// 	rltPort.Sub(mnr.Rx, mnr.Id)
// 	mnr.Start(hstPort.Stgys.At(0).I()) // send on one for all
// 	mnr.WaitFor(hstPort.Trds.Cnt())
// 	if hstPort.Trds != nil {
// 		tst.AnaTrdsEql(t, hstPort.Trds, mnr.Trds, "Trds")
// 		tst.AnaPortEql(t, &hstPort.Port, &rltPort.Port, "Port")
// 		tst.AnaPrfmEql(t, hstPrfm, rltPort.Prfm(), "Prfm")
// 	}
// }

// func TestRltPort2NzdUsd(t *testing.T) {
// 	ap := app.New(tst.Cfg)
// 	ap.Oan.CloneInstrs(tst.Instrs)
// 	defer ap.Cls()
// 	instr := rlt.Oan().NzdUsd()
// 	rltPort := rlt.NewPort(
// 		instr.I(5).Bid().Lst().InrLss(1).Shrt(1.0, 1.0, tme.H1, instr),
// 		instr.I(5).Bid().Lst().InrGtr(1).Long(1.0, 1.0, tme.H1, instr),
// 	)
// 	hstPort := ap.Actr.RunHst(rltPort.String())[0].(*hst.Port)
// 	hstPrfm := hstPort.Prfm() // call before WaitFor (for CalcTrds)
// 	// sys.Log("hstPort.Trds.Cnt()", hstPort.Trds.Cnt())
// 	mnr := tst.NewStgyMnr(ap)
// 	rltPort.Sub(mnr.Rx, mnr.Id)
// 	mnr.Start(hstPort.Stgys.At(0).I()) // send on one for all
// 	mnr.WaitFor(hstPort.Trds.Cnt())
// 	if hstPort.Trds != nil {
// 		tst.AnaTrdsEql(t, hstPort.Trds, mnr.Trds, "Trds")
// 		tst.AnaPortEql(t, &hstPort.Port, &rltPort.Port, "Port")
// 		tst.AnaPrfmEql(t, hstPrfm, rltPort.Prfm(), "Prfm")
// 	}
// }

// func TestRltPort2GbpUsd(t *testing.T) {
// 	ap := app.New(tst.Cfg)
// 	ap.Oan.CloneInstrs(tst.Instrs)
// 	defer ap.Cls()
// 	instr := rlt.Oan().GbpUsd()
// 	rltPort := rlt.NewPort(
// 		instr.I(5).Bid().Lst().InrLss(1).Shrt(1.0, 1.0, tme.H1, instr),
// 		instr.I(5).Bid().Lst().InrGtr(1).Long(1.0, 1.0, tme.H1, instr),
// 	)
// 	hstPort := ap.Actr.RunHst(rltPort.String())[0].(*hst.Port)
// 	hstPrfm := hstPort.Prfm() // call before WaitFor (for CalcTrds)
// 	// sys.Log("hstPort.Trds.Cnt()", hstPort.Trds.Cnt())
// 	mnr := tst.NewStgyMnr(ap)
// 	rltPort.Sub(mnr.Rx, mnr.Id)
// 	mnr.Start(hstPort.Stgys.At(0).I()) // send on one for all
// 	mnr.WaitFor(hstPort.Trds.Cnt())
// 	if hstPort.Trds != nil {
// 		tst.AnaTrdsEql(t, hstPort.Trds, mnr.Trds, "Trds")
// 		tst.AnaPortEql(t, &hstPort.Port, &rltPort.Port, "Port")
// 		tst.AnaPrfmEql(t, hstPrfm, rltPort.Prfm(), "Prfm")
// 	}
// }
