package ana_test

// func TestAnaStmInAll(t *testing.T) {
// 	ap := app.New(tst.Cfg)
// 	defer ap.Cls()
// 	ap.Oan.CloneInstrs(tst.Instrs)
// 	for _, i := range ap.Oan.Instrs {
// 		e := i.HstStm
// 		t.Run(i.Name.Unquo(), func(t *testing.T) {
// 			a := e.In(0, e.Cnt())
// 			// a.Validate()
// 			for n := unt.Zero; n < a.Cnt(); n++ {
// 				if e.Tmes.At(n).Neq(a.Tmes.At(n)) {
// 					t.Fatalf("Tmes (n:%v e:%v a:%v)", n, e.Tmes.At(n), a.Tmes.At(n))
// 				}
// 				if e.BidLims.At(n).Neq(a.BidLims.At(n)) {
// 					t.Fatalf("BidLims (n:%v e:%v a:%v)", n, e.BidLims.At(n), a.BidLims.At(n))
// 				}
// 				if e.AskLims.At(n).Neq(a.AskLims.At(n)) {
// 					t.Fatalf("AskLims (n:%v e:%v a:%v)", n, e.AskLims.At(n), a.AskLims.At(n))
// 				}
// 				if e.BidsAt(n).Neq(a.BidsAt(n)) {
// 					t.Fatalf("BidsAt (n:%v e:%v a:%v)", n, e.BidsAt(n), a.BidsAt(n))
// 				}
// 				if e.AsksAt(n).Neq(a.AsksAt(n)) {
// 					t.Fatalf("AsksAt (n:%v e:%v a:%v)", n, e.AsksAt(n), a.AsksAt(n))
// 				}
// 			}
// 		})
// 	}
// }

// func TestAnaStmInMdlPct40(t *testing.T) {
// 	ap := app.New(tst.Cfg)
// 	defer ap.Cls()
// 	ap.Oan.CloneInstrs(tst.Instrs)
// 	for _, i := range ap.Oan.Instrs {
// 		e := i.HstStm
// 		lenPct30 := unt.Unt(flt.Flt(e.Cnt()) * .3)
// 		fstIdx := lenPct30
// 		lstIdx := e.Cnt() - lenPct30
// 		t.Run(i.Name.Unquo(), func(t *testing.T) {
// 			a := e.In(fstIdx, lstIdx)
// 			// a.Validate()
// 			bidLimDlta := e.BidLims.At(fstIdx - 1)
// 			askLimDlta := e.AskLims.At(fstIdx - 1)
// 			for n := unt.Zero; n < a.Cnt(); n++ {
// 				if e.Tmes.At(fstIdx + n).Neq(a.Tmes.At(n)) {
// 					t.Fatalf("Tmes (n:%v e:%v a:%v)", n, e.Tmes.At(fstIdx+n), a.Tmes.At(n))
// 				}
// 				if e.BidLims.At(fstIdx + n).Sub(bidLimDlta).Neq(a.BidLims.At(n)) {
// 					t.Fatalf("BidLims (n:%v bidLimDlta:%v e:%v a:%v)", n, bidLimDlta, e.BidLims.At(fstIdx+n).Sub(bidLimDlta), a.BidLims.At(n))
// 				}
// 				if e.AskLims.At(fstIdx + n).Sub(askLimDlta).Neq(a.AskLims.At(n)) {
// 					t.Fatalf("AskLims (n:%v askLimDlta:%v e:%v a:%v)", n, askLimDlta, e.AskLims.At(fstIdx+n).Sub(askLimDlta), a.AskLims.At(n))
// 				}
// 				if e.BidsAt(fstIdx + n).Neq(a.BidsAt(n)) {
// 					t.Fatalf("BidsAt (n:%v e:%v a:%v)", n, e.BidsAt(fstIdx+n), a.BidsAt(n))
// 				}
// 				if e.AsksAt(fstIdx + n).Neq(a.AsksAt(n)) {
// 					t.Fatalf("AsksAt (n:%v e:%v a:%v)", n, e.AsksAt(fstIdx+n), a.AsksAt(n))
// 				}
// 			}
// 		})
// 	}
// }

// func TestAnaStmTo(t *testing.T) {
// 	ap := app.New(tst.Cfg)
// 	defer ap.Cls()
// 	ap.Oan.CloneInstrs(tst.Instrs)
// 	for _, i := range ap.Oan.Instrs {
// 		v := i.HstStm
// 		lim := unt.Unt(flt.Flt(v.Cnt()) * .3)
// 		t.Run(i.Name.Unquo(), func(t *testing.T) {
// 			e := &ana.Stm{}
// 			e.Tmes = v.Tmes.To(lim)
// 			e.BidLims = v.BidLims.To(lim)
// 			e.AskLims = v.AskLims.To(lim)
// 			e.Bids = v.Bids.To(v.BidLims.Lst())
// 			e.Asks = v.Asks.To(v.AskLims.Lst())
// 			a := v.To(lim)
// 			//a.Validate()
// 			for n := unt.Zero; n < a.Cnt(); n++ {
// 				if e.Tmes.At(n).Neq(a.Tmes.At(n)) {
// 					t.Fatalf("Tmes (n:%v e:%v a:%v)", n, e.Tmes.At(n), a.Tmes.At(n))
// 				}
// 				if e.BidLims.At(n).Neq(a.BidLims.At(n)) {
// 					t.Fatalf("BidLims (n:%v e:%v a:%v)", n, e.BidLims.At(n), a.BidLims.At(n))
// 				}
// 				if e.AskLims.At(n).Neq(a.AskLims.At(n)) {
// 					t.Fatalf("AskLims (n:%v e:%v a:%v)", n, e.AskLims.At(n), a.AskLims.At(n))
// 				}
// 				if e.BidsAt(n).Neq(a.BidsAt(n)) {
// 					t.Fatalf("BidsAt (n:%v e:%v a:%v)", n, e.BidsAt(n), a.BidsAt(n))
// 				}
// 				if e.AsksAt(n).Neq(a.AsksAt(n)) {
// 					t.Fatalf("AsksAt (n:%v e:%v a:%v)", n, e.AsksAt(n), a.AsksAt(n))
// 				}
// 			}
// 		})
// 	}
// }

// func TestAnaStmQual(t *testing.T) {
// 	cfg := cfg.Load("/home/rana/go/src/sys/cmd/sys.cfg")
// 	cfg.Ui = false
// 	ap := app.New(cfg)
// 	defer ap.Cls()
// 	sys.Log("--- START")

// 	oan := hst.Oan()

// 	f := func(start, end tme.Tme) {
// 		sys.Logf("--- start:%v end:%v", start, end)
// 		eurUsdQual := oan.EurUsd(start, end)
// 		inrvlQual := eurUsdQual.S(1)
// 		bidsQual := inrvlQual.Bids()
// 		lstQual := bidsQual.Lst().Bse()
// 		sys.Log("lstQual.Tmes", lstQual.Tmes)
// 		sys.Log("lstQual.Vals", lstQual.Vals)
// 	}

// 	f(prs.TmeTxt("2018-5-29"), prs.TmeTxt("2018-5-29/20s"))
// 	f(prs.TmeTxt("2018-5-30"), prs.TmeTxt("2018-5-30/20s"))

// 	eurUsd := oan.EurUsd()
// 	hstStm := eurUsd.I.HstStm
// 	lim := unt.Unt(8)
// 	tmes := hstStm.Tmes.To(lim)
// 	bidLims := hstStm.BidLims.To(lim)
// 	bidLimsLst := bidLims.Lst()
// 	bids := hstStm.Bids.To(bidLimsLst)
// 	sys.Log("         eurUsd.I.HstStm.Tmes", tmes)
// 	sys.Log("      eurUsd.I.HstStm.BidLims", bidLims)
// 	sys.Log("eurUsd.I.HstStm.BidLims.Lst()", bidLimsLst)
// 	sys.Log("    hstStm.Bids.To(bidLimsLst", bids)

// 	sys.Log("--- END")
// }
