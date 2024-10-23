package hst_test

import (
	"sys"
	"sys/ana/hst"
	"sys/app"
	"sys/bsc/tme"
	"sys/tst"
	"testing"
)

func TestHstStgy(t *testing.T) {
	cfg := tst.NewCfg()
	cfg.Hst.StgySeqTrd = true          // StgySeqTrd NOT NEEDED WHEN TESTING FUTURE PERIOD PRFM
	cfg.DskPth = "/home/rana/data/tst" // ADD DSK TO TEST ML NET SAV
	ap := app.New(cfg, "/home/rana/go/src/sys/cmd")
	ap.Oan.CloneInstrs(tst.Instrs)
	prv := hst.Oan()
	// instr := Instr(prv)
	Instr := tst.HstPrvInstrs[len(tst.HstPrvInstrs)-1] // lst elm has most points
	instr0 := Instr(prv)
	ts := instr0.Bse().Ana.HstStm.Tmes
	rng := tme.NewRng(ts.Fst(), ts.Mid())
	instr := prv.GbpUsd(rng)
	inrvl := tst.HstInstrInrvlI(instr, 10)
	side := tst.HstInrvlSideBid(inrvl)
	stm := tst.HstSideStmRteLst(side)
	cnd := tst.HstStmCndInrGtr(stm, 1)
	ftrStms := hst.NewStms(
		instr.I(1).Bid().Lst(), // 0.808
		instr.I(1).Bid().Min(), // 0.824
		instr.I(1).Bid().Max(), // 0.851

		instr.I(2).Bid().Lst(), // 0.808
		instr.I(2).Bid().Min(), // 0.824
		instr.I(2).Bid().Max(), // 0.851

		instr.I(3).Bid().Lst(), // 0.808
		instr.I(3).Bid().Min(), // 0.824
		instr.I(3).Bid().Max(), // 0.851

		instr.I(5).Bid().Lst(), // 0.808
		instr.I(5).Bid().Min(), // 0.824
		instr.I(5).Bid().Max(), // 0.851

		instr.I(8).Bid().Lst(), // 0.808
		instr.I(8).Bid().Min(), // 0.824
		instr.I(8).Bid().Max(), // 0.851

		// instr.I(10).Bid().RngFul(),       // 1.054
		// instr.I(10).Bid().RngLst(),       // 1.038
		// instr.I(10).Bid().Sum(),          // 1.041
		// instr.I(10).Bid().Std(),          // 1.050
		// instr.I(10).Bid().Vrnc(),         // 1.040
		// instr.I(10).Bid().Rsi(),          // 1.032
		// instr.I(10).Bid().Wrsi(),         // 1.030
		// instr.I(10).Bid().Sma(), // 0.875
		// instr.I(10).Bid().Gma(),          // 1.039
		// instr.I(10).Bid().Wma(),  // 0.940
		// instr.I(10).Bid().Alma(), // 0.818
		// instr.I(10).Bid().Ema(),  // 0.855
		// instr.I(10).Bid().ProSma(),       // 1.029
		// instr.I(10).Bid().ProAlma(),      // 1.048
		// instr.I(10).Bid().ProLst(),       // 1.035
		// instr.I(10).Bid().Sar(0.02, 0.2), // 1.156
	)
	stgy := tst.HstCndStgyStgy(cnd, true, 2.0, 4.0, 60*60, 0.0, instr, ftrStms).(*hst.StgyStgy)
	stgy.Fit()
	stgy.CalcSeqTrds(stgy.String())
	tst.HstStgyStgyNotZero(t, stgy)
	ap.Cls()
}

func TestHstStgyPrfm(t *testing.T) {
	cfg := tst.NewCfg()
	cfg.Hst.StgySeqTrd = true          // StgySeqTrd NOT NEEDED WHEN TESTING FUTURE PERIOD PRFM
	cfg.DskPth = "/home/rana/data/tst" // ADD DSK TO TEST ML NET SAV
	ap := app.New(cfg, "/home/rana/go/src/sys/cmd")
	ap.Oan.CloneInstrs(tst.Instrs)
	prv := hst.Oan()
	// instr := Instr(prv)
	Instr := tst.HstPrvInstrs[len(tst.HstPrvInstrs)-1] // lst elm has most points
	instr0 := Instr(prv)
	ts := instr0.Bse().Ana.HstStm.Tmes
	lim := int(0.8 * float64(len(*ts)))
	limT := (*ts)[lim]

	instr := prv.GbpUsd(tme.NewRng(ts.Fst(), limT))
	inrvl := tst.HstInstrInrvlI(instr, 10)
	side := tst.HstInrvlSideBid(inrvl)
	stm := tst.HstSideStmRteLst(side)
	cnd := tst.HstStmCndInrGtr(stm, 1)
	ftrStms := hst.NewStms(
		instr.I(1).Bid().Lst(), // 0.808
		instr.I(1).Bid().Min(), // 0.824
		instr.I(1).Bid().Max(), // 0.851

		instr.I(2).Bid().Lst(), // 0.808
		instr.I(2).Bid().Min(), // 0.824
		instr.I(2).Bid().Max(), // 0.851

		instr.I(3).Bid().Lst(), // 0.808
		instr.I(3).Bid().Min(), // 0.824
		instr.I(3).Bid().Max(), // 0.851

		instr.I(5).Bid().Lst(), // 0.808
		instr.I(5).Bid().Min(), // 0.824
		instr.I(5).Bid().Max(), // 0.851

		instr.I(8).Bid().Lst(), // 0.808
		instr.I(8).Bid().Min(), // 0.824
		instr.I(8).Bid().Max(), // 0.851
	)
	stgy := tst.HstCndStgyStgy(cnd, true, 2.0, 4.0, 60*60, 0.0, instr, ftrStms).(*hst.StgyStgy)
	stgy.Fit()

	prfm := stgy.Prfm(tme.NewRng(limT, ts.Lst()))
	sys.Log(prfm)
	// tst.AnaPrfmNotZero(t, prfm)
	ap.Cls()
}
