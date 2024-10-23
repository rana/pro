package srch

import (
	"sys"
	"sys/ana"
	"sys/ana/hst"
	"sys/bsc/flt"
	"sys/bsc/tme"
	"sys/bsc/tmes"
	"sys/bsc/unt"
	"sys/trc"
)

func SrchMl0() {
	sys.Log("- start")
	trcr := trc.New("srch")
	defer trcr.End()

	prfms := ana.NewPrfms()

	// mktRngs := hst.Oan().EurUsd().Bse().Ana.MktWeeks
	// fitRng := mktRngs.RngMrg(0, mktRngs.LstIdx()-1)
	// vldRng := mktRngs.RngMrg(mktRngs.LstIdx()-1, mktRngs.LstIdx())

	mktRngs := hst.Oan().EurUsd().Bse().Ana.MktDays
	// fitRng := mktRngs.RngMrg(0, 4)
	// vldRng := mktRngs.RngMrg(5, 6)

	mktRngsCnt := mktRngs.Cnt() - 6
	for n := unt.Zero; n < mktRngsCnt; n++ {
		fitRng := mktRngs.RngMrg(n, n+4)
		vldRng := mktRngs.RngMrg(n+5, n+6)

		sys.Log("-  fitRng", fitRng)
		sys.Log("-  vldRng", vldRng)
		prfLim := flt.Flt(20.0)
		losLim := flt.Flt(20.0)
		durLim := tme.H1
		minPnlPct := flt.Flt(0.0)
		Instr := EurUsd
		instr := Instr(fitRng)
		durs := tmes.New(1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 600, 977, 1577)
		ftrStms := hst.NewStms()
		for _, dur := range *durs {
			inrvlBid := instr.I(dur).Bid()
			lst := inrvlBid.Lst()
			min := inrvlBid.Min()
			max := inrvlBid.Max()
			rsi := inrvlBid.Rsi()
			ftrStms.Push(lst)
			ftrStms.Push(min)
			ftrStms.Push(max)
			ftrStms.Push(rsi)
			// ftrStms.Push(inrvlBid.Std())
			// ftrStms.Push(inrvlBid.RngFul())
			// ftrStms.Push(inrvlBid.RngLst())
			ftrStms.Push(max.InrSub(1))
			ftrStms.Push(rsi.InrSub(1))
		}
		opnCnd := OpnCndBlngrUprCrsUp(instr, tme.M50, tme.M1*400)
		stgy := opnCnd.Stgy(true, prfLim, losLim, durLim, minPnlPct, instr, ftrStms).(*hst.StgyStgy)
		stgy.Fit()
		prfm := stgy.Prfm(vldRng)
		sys.Log(prfm)
		prfms.Push(prfm)
	}

	pnlPctSum := prfms.PnlPcts().Sum()
	sys.Logf("--- prfms: pnlPctSum:%v cnt:%v", pnlPctSum, prfms.Cnt())
	sys.Log(prfms)
	sys.Logf("--- prfms: pnlPctSum:%v cnt:%v", pnlPctSum, prfms.Cnt())

	sys.Log("- end")
}
