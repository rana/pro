package rlt_test

import (
	"sys/ana/hst"
	"sys/ana/rlt"
	"sys/app"
	"sys/bsc/flt"
	"sys/bsc/unt"
	"sys/lng/pro/act"
	"sys/tst"
	"testing"
)

func TestRltSideFst(t *testing.T) {
	for _, instr := range tst.RltPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := rlt.Oan()
			instr := instr(prv)
			inrvl := tst.RltInstrInrvlI(instr, 10)
			side := tst.RltInrvlSideBid(inrvl)
			a := side.Fst().(*rlt.StmRteFst)
			tst.RltStmRteFstNotZero(t, a)
			mnr := tst.NewStmMnr(ap)
			a.Sub(mnr.Rx, mnr.Id)
			tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
			var actr act.Actr
			vs := actr.RunHst(a.String())
			eHst := vs[len(vs)-1].(*hst.StmRteFst)
			if eHst.Tmes != nil {
				mnr.StartFor(instr.Instr(), eHst.Tmes.Cnt())
				tst.TmesEql(t, eHst.Tmes, mnr.Tmes, "Tmes")
				tst.FltsEql(t, eHst.Vals, mnr.Vals, "Vals")
			}
			a.Unsub(mnr.Id)
			tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
			ap.Cls()
		})
	}
}
func TestRltSideLst(t *testing.T) {
	for _, instr := range tst.RltPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := rlt.Oan()
			instr := instr(prv)
			inrvl := tst.RltInstrInrvlI(instr, 10)
			side := tst.RltInrvlSideBid(inrvl)
			a := side.Lst().(*rlt.StmRteLst)
			tst.RltStmRteLstNotZero(t, a)
			mnr := tst.NewStmMnr(ap)
			a.Sub(mnr.Rx, mnr.Id)
			tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
			var actr act.Actr
			vs := actr.RunHst(a.String())
			eHst := vs[len(vs)-1].(*hst.StmRteLst)
			if eHst.Tmes != nil {
				mnr.StartFor(instr.Instr(), eHst.Tmes.Cnt())
				tst.TmesEql(t, eHst.Tmes, mnr.Tmes, "Tmes")
				tst.FltsEql(t, eHst.Vals, mnr.Vals, "Vals")
			}
			a.Unsub(mnr.Id)
			tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
			ap.Cls()
		})
	}
}
func TestRltSideSum(t *testing.T) {
	for _, instr := range tst.RltPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := rlt.Oan()
			instr := instr(prv)
			inrvl := tst.RltInstrInrvlI(instr, 10)
			side := tst.RltInrvlSideBid(inrvl)
			a := side.Sum().(*rlt.StmRteSum)
			tst.RltStmRteSumNotZero(t, a)
			mnr := tst.NewStmMnr(ap)
			a.Sub(mnr.Rx, mnr.Id)
			tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
			var actr act.Actr
			vs := actr.RunHst(a.String())
			eHst := vs[len(vs)-1].(*hst.StmRteSum)
			if eHst.Tmes != nil {
				mnr.StartFor(instr.Instr(), eHst.Tmes.Cnt())
				tst.TmesEql(t, eHst.Tmes, mnr.Tmes, "Tmes")
				tst.FltsEql(t, eHst.Vals, mnr.Vals, "Vals")
			}
			a.Unsub(mnr.Id)
			tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
			ap.Cls()
		})
	}
}
func TestRltSidePrd(t *testing.T) {
	for _, instr := range tst.RltPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := rlt.Oan()
			instr := instr(prv)
			inrvl := tst.RltInstrInrvlI(instr, 10)
			side := tst.RltInrvlSideBid(inrvl)
			a := side.Prd().(*rlt.StmRtePrd)
			tst.RltStmRtePrdNotZero(t, a)
			mnr := tst.NewStmMnr(ap)
			a.Sub(mnr.Rx, mnr.Id)
			tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
			var actr act.Actr
			vs := actr.RunHst(a.String())
			eHst := vs[len(vs)-1].(*hst.StmRtePrd)
			if eHst.Tmes != nil {
				mnr.StartFor(instr.Instr(), eHst.Tmes.Cnt())
				tst.TmesEql(t, eHst.Tmes, mnr.Tmes, "Tmes")
				tst.FltsEql(t, eHst.Vals, mnr.Vals, "Vals")
			}
			a.Unsub(mnr.Id)
			tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
			ap.Cls()
		})
	}
}
func TestRltSideMin(t *testing.T) {
	for _, instr := range tst.RltPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := rlt.Oan()
			instr := instr(prv)
			inrvl := tst.RltInstrInrvlI(instr, 10)
			side := tst.RltInrvlSideBid(inrvl)
			a := side.Min().(*rlt.StmRteMin)
			tst.RltStmRteMinNotZero(t, a)
			mnr := tst.NewStmMnr(ap)
			a.Sub(mnr.Rx, mnr.Id)
			tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
			var actr act.Actr
			vs := actr.RunHst(a.String())
			eHst := vs[len(vs)-1].(*hst.StmRteMin)
			if eHst.Tmes != nil {
				mnr.StartFor(instr.Instr(), eHst.Tmes.Cnt())
				tst.TmesEql(t, eHst.Tmes, mnr.Tmes, "Tmes")
				tst.FltsEql(t, eHst.Vals, mnr.Vals, "Vals")
			}
			a.Unsub(mnr.Id)
			tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
			ap.Cls()
		})
	}
}
func TestRltSideMax(t *testing.T) {
	for _, instr := range tst.RltPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := rlt.Oan()
			instr := instr(prv)
			inrvl := tst.RltInstrInrvlI(instr, 10)
			side := tst.RltInrvlSideBid(inrvl)
			a := side.Max().(*rlt.StmRteMax)
			tst.RltStmRteMaxNotZero(t, a)
			mnr := tst.NewStmMnr(ap)
			a.Sub(mnr.Rx, mnr.Id)
			tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
			var actr act.Actr
			vs := actr.RunHst(a.String())
			eHst := vs[len(vs)-1].(*hst.StmRteMax)
			if eHst.Tmes != nil {
				mnr.StartFor(instr.Instr(), eHst.Tmes.Cnt())
				tst.TmesEql(t, eHst.Tmes, mnr.Tmes, "Tmes")
				tst.FltsEql(t, eHst.Vals, mnr.Vals, "Vals")
			}
			a.Unsub(mnr.Id)
			tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
			ap.Cls()
		})
	}
}
func TestRltSideMid(t *testing.T) {
	for _, instr := range tst.RltPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := rlt.Oan()
			instr := instr(prv)
			inrvl := tst.RltInstrInrvlI(instr, 10)
			side := tst.RltInrvlSideBid(inrvl)
			a := side.Mid().(*rlt.StmRteMid)
			tst.RltStmRteMidNotZero(t, a)
			mnr := tst.NewStmMnr(ap)
			a.Sub(mnr.Rx, mnr.Id)
			tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
			var actr act.Actr
			vs := actr.RunHst(a.String())
			eHst := vs[len(vs)-1].(*hst.StmRteMid)
			if eHst.Tmes != nil {
				mnr.StartFor(instr.Instr(), eHst.Tmes.Cnt())
				tst.TmesEql(t, eHst.Tmes, mnr.Tmes, "Tmes")
				tst.FltsEql(t, eHst.Vals, mnr.Vals, "Vals")
			}
			a.Unsub(mnr.Id)
			tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
			ap.Cls()
		})
	}
}
func TestRltSideMdn(t *testing.T) {
	for _, instr := range tst.RltPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := rlt.Oan()
			instr := instr(prv)
			inrvl := tst.RltInstrInrvlI(instr, 10)
			side := tst.RltInrvlSideBid(inrvl)
			a := side.Mdn().(*rlt.StmRteMdn)
			tst.RltStmRteMdnNotZero(t, a)
			mnr := tst.NewStmMnr(ap)
			a.Sub(mnr.Rx, mnr.Id)
			tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
			var actr act.Actr
			vs := actr.RunHst(a.String())
			eHst := vs[len(vs)-1].(*hst.StmRteMdn)
			if eHst.Tmes != nil {
				mnr.StartFor(instr.Instr(), eHst.Tmes.Cnt())
				tst.TmesEql(t, eHst.Tmes, mnr.Tmes, "Tmes")
				tst.FltsEql(t, eHst.Vals, mnr.Vals, "Vals")
			}
			a.Unsub(mnr.Id)
			tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
			ap.Cls()
		})
	}
}
func TestRltSideSma(t *testing.T) {
	for _, instr := range tst.RltPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := rlt.Oan()
			instr := instr(prv)
			inrvl := tst.RltInstrInrvlI(instr, 10)
			side := tst.RltInrvlSideBid(inrvl)
			a := side.Sma().(*rlt.StmRteSma)
			tst.RltStmRteSmaNotZero(t, a)
			mnr := tst.NewStmMnr(ap)
			a.Sub(mnr.Rx, mnr.Id)
			tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
			var actr act.Actr
			vs := actr.RunHst(a.String())
			eHst := vs[len(vs)-1].(*hst.StmRteSma)
			if eHst.Tmes != nil {
				mnr.StartFor(instr.Instr(), eHst.Tmes.Cnt())
				tst.TmesEql(t, eHst.Tmes, mnr.Tmes, "Tmes")
				tst.FltsEql(t, eHst.Vals, mnr.Vals, "Vals")
			}
			a.Unsub(mnr.Id)
			tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
			ap.Cls()
		})
	}
}
func TestRltSideGma(t *testing.T) {
	for _, instr := range tst.RltPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := rlt.Oan()
			instr := instr(prv)
			inrvl := tst.RltInstrInrvlI(instr, 10)
			side := tst.RltInrvlSideBid(inrvl)
			a := side.Gma().(*rlt.StmRteGma)
			tst.RltStmRteGmaNotZero(t, a)
			mnr := tst.NewStmMnr(ap)
			a.Sub(mnr.Rx, mnr.Id)
			tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
			var actr act.Actr
			vs := actr.RunHst(a.String())
			eHst := vs[len(vs)-1].(*hst.StmRteGma)
			if eHst.Tmes != nil {
				mnr.StartFor(instr.Instr(), eHst.Tmes.Cnt())
				tst.TmesEql(t, eHst.Tmes, mnr.Tmes, "Tmes")
				tst.FltsEql(t, eHst.Vals, mnr.Vals, "Vals")
			}
			a.Unsub(mnr.Id)
			tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
			ap.Cls()
		})
	}
}
func TestRltSideWma(t *testing.T) {
	for _, instr := range tst.RltPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := rlt.Oan()
			instr := instr(prv)
			inrvl := tst.RltInstrInrvlI(instr, 10)
			side := tst.RltInrvlSideBid(inrvl)
			a := side.Wma().(*rlt.StmRteWma)
			tst.RltStmRteWmaNotZero(t, a)
			mnr := tst.NewStmMnr(ap)
			a.Sub(mnr.Rx, mnr.Id)
			tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
			var actr act.Actr
			vs := actr.RunHst(a.String())
			eHst := vs[len(vs)-1].(*hst.StmRteWma)
			if eHst.Tmes != nil {
				mnr.StartFor(instr.Instr(), eHst.Tmes.Cnt())
				tst.TmesEql(t, eHst.Tmes, mnr.Tmes, "Tmes")
				tst.FltsEql(t, eHst.Vals, mnr.Vals, "Vals")
			}
			a.Unsub(mnr.Id)
			tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
			ap.Cls()
		})
	}
}
func TestRltSideRsi(t *testing.T) {
	for _, instr := range tst.RltPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := rlt.Oan()
			instr := instr(prv)
			inrvl := tst.RltInstrInrvlI(instr, 10)
			side := tst.RltInrvlSideBid(inrvl)
			a := side.Rsi().(*rlt.StmRteRsi)
			tst.RltStmRteRsiNotZero(t, a)
			mnr := tst.NewStmMnr(ap)
			a.Sub(mnr.Rx, mnr.Id)
			tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
			var actr act.Actr
			vs := actr.RunHst(a.String())
			eHst := vs[len(vs)-1].(*hst.StmRteRsi)
			if eHst.Tmes != nil {
				mnr.StartFor(instr.Instr(), eHst.Tmes.Cnt())
				tst.TmesEql(t, eHst.Tmes, mnr.Tmes, "Tmes")
				tst.FltsEql(t, eHst.Vals, mnr.Vals, "Vals")
			}
			a.Unsub(mnr.Id)
			tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
			ap.Cls()
		})
	}
}
func TestRltSideWrsi(t *testing.T) {
	for _, instr := range tst.RltPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := rlt.Oan()
			instr := instr(prv)
			inrvl := tst.RltInstrInrvlI(instr, 10)
			side := tst.RltInrvlSideBid(inrvl)
			a := side.Wrsi().(*rlt.StmRteWrsi)
			tst.RltStmRteWrsiNotZero(t, a)
			mnr := tst.NewStmMnr(ap)
			a.Sub(mnr.Rx, mnr.Id)
			tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
			var actr act.Actr
			vs := actr.RunHst(a.String())
			eHst := vs[len(vs)-1].(*hst.StmRteWrsi)
			if eHst.Tmes != nil {
				mnr.StartFor(instr.Instr(), eHst.Tmes.Cnt())
				tst.TmesEql(t, eHst.Tmes, mnr.Tmes, "Tmes")
				tst.FltsEql(t, eHst.Vals, mnr.Vals, "Vals")
			}
			a.Unsub(mnr.Id)
			tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
			ap.Cls()
		})
	}
}
func TestRltSideAlma(t *testing.T) {
	for _, instr := range tst.RltPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := rlt.Oan()
			instr := instr(prv)
			inrvl := tst.RltInstrInrvlI(instr, 10)
			side := tst.RltInrvlSideBid(inrvl)
			a := side.Alma().(*rlt.StmRteAlma)
			tst.RltStmRteAlmaNotZero(t, a)
			mnr := tst.NewStmMnr(ap)
			a.Sub(mnr.Rx, mnr.Id)
			tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
			var actr act.Actr
			vs := actr.RunHst(a.String())
			eHst := vs[len(vs)-1].(*hst.StmRteAlma)
			if eHst.Tmes != nil {
				mnr.StartFor(instr.Instr(), eHst.Tmes.Cnt())
				tst.TmesEql(t, eHst.Tmes, mnr.Tmes, "Tmes")
				tst.FltsEql(t, eHst.Vals, mnr.Vals, "Vals")
			}
			a.Unsub(mnr.Id)
			tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
			ap.Cls()
		})
	}
}
func TestRltSideVrnc(t *testing.T) {
	for _, instr := range tst.RltPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := rlt.Oan()
			instr := instr(prv)
			inrvl := tst.RltInstrInrvlI(instr, 10)
			side := tst.RltInrvlSideBid(inrvl)
			a := side.Vrnc().(*rlt.StmRteVrnc)
			tst.RltStmRteVrncNotZero(t, a)
			mnr := tst.NewStmMnr(ap)
			a.Sub(mnr.Rx, mnr.Id)
			tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
			var actr act.Actr
			vs := actr.RunHst(a.String())
			eHst := vs[len(vs)-1].(*hst.StmRteVrnc)
			if eHst.Tmes != nil {
				mnr.StartFor(instr.Instr(), eHst.Tmes.Cnt())
				tst.TmesEql(t, eHst.Tmes, mnr.Tmes, "Tmes")
				tst.FltsEql(t, eHst.Vals, mnr.Vals, "Vals")
			}
			a.Unsub(mnr.Id)
			tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
			ap.Cls()
		})
	}
}
func TestRltSideStd(t *testing.T) {
	for _, instr := range tst.RltPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := rlt.Oan()
			instr := instr(prv)
			inrvl := tst.RltInstrInrvlI(instr, 10)
			side := tst.RltInrvlSideBid(inrvl)
			a := side.Std().(*rlt.StmRteStd)
			tst.RltStmRteStdNotZero(t, a)
			mnr := tst.NewStmMnr(ap)
			a.Sub(mnr.Rx, mnr.Id)
			tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
			var actr act.Actr
			vs := actr.RunHst(a.String())
			eHst := vs[len(vs)-1].(*hst.StmRteStd)
			if eHst.Tmes != nil {
				mnr.StartFor(instr.Instr(), eHst.Tmes.Cnt())
				tst.TmesEql(t, eHst.Tmes, mnr.Tmes, "Tmes")
				tst.FltsEql(t, eHst.Vals, mnr.Vals, "Vals")
			}
			a.Unsub(mnr.Id)
			tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
			ap.Cls()
		})
	}
}
func TestRltSideRngFul(t *testing.T) {
	for _, instr := range tst.RltPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := rlt.Oan()
			instr := instr(prv)
			inrvl := tst.RltInstrInrvlI(instr, 10)
			side := tst.RltInrvlSideBid(inrvl)
			a := side.RngFul().(*rlt.StmRteRngFul)
			tst.RltStmRteRngFulNotZero(t, a)
			mnr := tst.NewStmMnr(ap)
			a.Sub(mnr.Rx, mnr.Id)
			tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
			var actr act.Actr
			vs := actr.RunHst(a.String())
			eHst := vs[len(vs)-1].(*hst.StmRteRngFul)
			if eHst.Tmes != nil {
				mnr.StartFor(instr.Instr(), eHst.Tmes.Cnt())
				tst.TmesEql(t, eHst.Tmes, mnr.Tmes, "Tmes")
				tst.FltsEql(t, eHst.Vals, mnr.Vals, "Vals")
			}
			a.Unsub(mnr.Id)
			tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
			ap.Cls()
		})
	}
}
func TestRltSideRngLst(t *testing.T) {
	for _, instr := range tst.RltPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := rlt.Oan()
			instr := instr(prv)
			inrvl := tst.RltInstrInrvlI(instr, 10)
			side := tst.RltInrvlSideBid(inrvl)
			a := side.RngLst().(*rlt.StmRteRngLst)
			tst.RltStmRteRngLstNotZero(t, a)
			mnr := tst.NewStmMnr(ap)
			a.Sub(mnr.Rx, mnr.Id)
			tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
			var actr act.Actr
			vs := actr.RunHst(a.String())
			eHst := vs[len(vs)-1].(*hst.StmRteRngLst)
			if eHst.Tmes != nil {
				mnr.StartFor(instr.Instr(), eHst.Tmes.Cnt())
				tst.TmesEql(t, eHst.Tmes, mnr.Tmes, "Tmes")
				tst.FltsEql(t, eHst.Vals, mnr.Vals, "Vals")
			}
			a.Unsub(mnr.Id)
			tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
			ap.Cls()
		})
	}
}
func TestRltSideProLst(t *testing.T) {
	for _, instr := range tst.RltPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := rlt.Oan()
			instr := instr(prv)
			inrvl := tst.RltInstrInrvlI(instr, 10)
			side := tst.RltInrvlSideBid(inrvl)
			a := side.ProLst().(*rlt.StmRteProLst)
			tst.RltStmRteProLstNotZero(t, a)
			mnr := tst.NewStmMnr(ap)
			a.Sub(mnr.Rx, mnr.Id)
			tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
			var actr act.Actr
			vs := actr.RunHst(a.String())
			eHst := vs[len(vs)-1].(*hst.StmRteProLst)
			if eHst.Tmes != nil {
				mnr.StartFor(instr.Instr(), eHst.Tmes.Cnt())
				tst.TmesEql(t, eHst.Tmes, mnr.Tmes, "Tmes")
				tst.FltsEql(t, eHst.Vals, mnr.Vals, "Vals")
			}
			a.Unsub(mnr.Id)
			tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
			ap.Cls()
		})
	}
}
func TestRltSideProSma(t *testing.T) {
	for _, instr := range tst.RltPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := rlt.Oan()
			instr := instr(prv)
			inrvl := tst.RltInstrInrvlI(instr, 10)
			side := tst.RltInrvlSideBid(inrvl)
			a := side.ProSma().(*rlt.StmRteProSma)
			tst.RltStmRteProSmaNotZero(t, a)
			mnr := tst.NewStmMnr(ap)
			a.Sub(mnr.Rx, mnr.Id)
			tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
			var actr act.Actr
			vs := actr.RunHst(a.String())
			eHst := vs[len(vs)-1].(*hst.StmRteProSma)
			if eHst.Tmes != nil {
				mnr.StartFor(instr.Instr(), eHst.Tmes.Cnt())
				tst.TmesEql(t, eHst.Tmes, mnr.Tmes, "Tmes")
				tst.FltsEql(t, eHst.Vals, mnr.Vals, "Vals")
			}
			a.Unsub(mnr.Id)
			tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
			ap.Cls()
		})
	}
}
func TestRltSideProAlma(t *testing.T) {
	for _, instr := range tst.RltPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := rlt.Oan()
			instr := instr(prv)
			inrvl := tst.RltInstrInrvlI(instr, 10)
			side := tst.RltInrvlSideBid(inrvl)
			a := side.ProAlma().(*rlt.StmRteProAlma)
			tst.RltStmRteProAlmaNotZero(t, a)
			mnr := tst.NewStmMnr(ap)
			a.Sub(mnr.Rx, mnr.Id)
			tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
			var actr act.Actr
			vs := actr.RunHst(a.String())
			eHst := vs[len(vs)-1].(*hst.StmRteProAlma)
			if eHst.Tmes != nil {
				mnr.StartFor(instr.Instr(), eHst.Tmes.Cnt())
				tst.TmesEql(t, eHst.Tmes, mnr.Tmes, "Tmes")
				tst.FltsEql(t, eHst.Vals, mnr.Vals, "Vals")
			}
			a.Unsub(mnr.Id)
			tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
			ap.Cls()
		})
	}
}
func TestRltSideSar(t *testing.T) {
	for _, instr := range tst.RltPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := rlt.Oan()
			instr := instr(prv)
			inrvl := tst.RltInstrInrvlI(instr, 10)
			side := tst.RltInrvlSideBid(inrvl)
			a := side.Sar(flt.Flt(0.02), flt.Flt(0.2)).(*rlt.StmRte1Sar)
			tst.RltStmRte1SarNotZero(t, a)
			mnr := tst.NewStmMnr(ap)
			a.Sub(mnr.Rx, mnr.Id)
			tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
			var actr act.Actr
			vs := actr.RunHst(a.String())
			eHst := vs[len(vs)-1].(*hst.StmRte1Sar)
			if eHst.Tmes != nil {
				mnr.StartFor(instr.Instr(), eHst.Tmes.Cnt())
				tst.TmesEql(t, eHst.Tmes, mnr.Tmes, "Tmes")
				tst.FltsEql(t, eHst.Vals, mnr.Vals, "Vals")
			}
			a.Unsub(mnr.Id)
			tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
			ap.Cls()
		})
	}
}
func TestRltSideEma(t *testing.T) {
	for _, instr := range tst.RltPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := rlt.Oan()
			instr := instr(prv)
			inrvl := tst.RltInstrInrvlI(instr, 10)
			side := tst.RltInrvlSideBid(inrvl)
			a := side.Ema().(*rlt.StmRteEma)
			tst.RltStmRteEmaNotZero(t, a)
			mnr := tst.NewStmMnr(ap)
			a.Sub(mnr.Rx, mnr.Id)
			tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
			var actr act.Actr
			vs := actr.RunHst(a.String())
			eHst := vs[len(vs)-1].(*hst.StmRteEma)
			if eHst.Tmes != nil {
				mnr.StartFor(instr.Instr(), eHst.Tmes.Cnt())
				tst.TmesEql(t, eHst.Tmes, mnr.Tmes, "Tmes")
				tst.FltsEql(t, eHst.Vals, mnr.Vals, "Vals")
			}
			a.Unsub(mnr.Id)
			tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
			ap.Cls()
		})
	}
}
func TestRltStmUnaPos(t *testing.T) {
	for _, instr := range tst.RltPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := rlt.Oan()
			instr := instr(prv)
			inrvl := tst.RltInstrInrvlI(instr, 10)
			side := tst.RltInrvlSideBid(inrvl)
			stm := tst.RltSideStmRteLst(side)
			a := stm.UnaPos().(*rlt.StmUnaPos)
			tst.RltStmUnaPosNotZero(t, a)
			mnr := tst.NewStmMnr(ap)
			a.Sub(mnr.Rx, mnr.Id)
			tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
			var actr act.Actr
			vs := actr.RunHst(a.String())
			eHst := vs[len(vs)-1].(*hst.StmUnaPos)
			if eHst.Tmes != nil {
				mnr.StartFor(instr.Instr(), eHst.Tmes.Cnt())
				tst.TmesEql(t, eHst.Tmes, mnr.Tmes, "Tmes")
				tst.FltsEql(t, eHst.Vals, mnr.Vals, "Vals")
			}
			a.Unsub(mnr.Id)
			tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
			ap.Cls()
		})
	}
}
func TestRltStmUnaNeg(t *testing.T) {
	for _, instr := range tst.RltPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := rlt.Oan()
			instr := instr(prv)
			inrvl := tst.RltInstrInrvlI(instr, 10)
			side := tst.RltInrvlSideBid(inrvl)
			stm := tst.RltSideStmRteLst(side)
			a := stm.UnaNeg().(*rlt.StmUnaNeg)
			tst.RltStmUnaNegNotZero(t, a)
			mnr := tst.NewStmMnr(ap)
			a.Sub(mnr.Rx, mnr.Id)
			tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
			var actr act.Actr
			vs := actr.RunHst(a.String())
			eHst := vs[len(vs)-1].(*hst.StmUnaNeg)
			if eHst.Tmes != nil {
				mnr.StartFor(instr.Instr(), eHst.Tmes.Cnt())
				tst.TmesEql(t, eHst.Tmes, mnr.Tmes, "Tmes")
				tst.FltsEql(t, eHst.Vals, mnr.Vals, "Vals")
			}
			a.Unsub(mnr.Id)
			tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
			ap.Cls()
		})
	}
}
func TestRltStmUnaInv(t *testing.T) {
	for _, instr := range tst.RltPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := rlt.Oan()
			instr := instr(prv)
			inrvl := tst.RltInstrInrvlI(instr, 10)
			side := tst.RltInrvlSideBid(inrvl)
			stm := tst.RltSideStmRteLst(side)
			a := stm.UnaInv().(*rlt.StmUnaInv)
			tst.RltStmUnaInvNotZero(t, a)
			mnr := tst.NewStmMnr(ap)
			a.Sub(mnr.Rx, mnr.Id)
			tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
			var actr act.Actr
			vs := actr.RunHst(a.String())
			eHst := vs[len(vs)-1].(*hst.StmUnaInv)
			if eHst.Tmes != nil {
				mnr.StartFor(instr.Instr(), eHst.Tmes.Cnt())
				tst.TmesEql(t, eHst.Tmes, mnr.Tmes, "Tmes")
				tst.FltsEql(t, eHst.Vals, mnr.Vals, "Vals")
			}
			a.Unsub(mnr.Id)
			tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
			ap.Cls()
		})
	}
}
func TestRltStmUnaSqr(t *testing.T) {
	for _, instr := range tst.RltPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := rlt.Oan()
			instr := instr(prv)
			inrvl := tst.RltInstrInrvlI(instr, 10)
			side := tst.RltInrvlSideBid(inrvl)
			stm := tst.RltSideStmRteLst(side)
			a := stm.UnaSqr().(*rlt.StmUnaSqr)
			tst.RltStmUnaSqrNotZero(t, a)
			mnr := tst.NewStmMnr(ap)
			a.Sub(mnr.Rx, mnr.Id)
			tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
			var actr act.Actr
			vs := actr.RunHst(a.String())
			eHst := vs[len(vs)-1].(*hst.StmUnaSqr)
			if eHst.Tmes != nil {
				mnr.StartFor(instr.Instr(), eHst.Tmes.Cnt())
				tst.TmesEql(t, eHst.Tmes, mnr.Tmes, "Tmes")
				tst.FltsEql(t, eHst.Vals, mnr.Vals, "Vals")
			}
			a.Unsub(mnr.Id)
			tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
			ap.Cls()
		})
	}
}
func TestRltStmUnaSqrt(t *testing.T) {
	for _, instr := range tst.RltPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := rlt.Oan()
			instr := instr(prv)
			inrvl := tst.RltInstrInrvlI(instr, 10)
			side := tst.RltInrvlSideBid(inrvl)
			stm := tst.RltSideStmRteLst(side)
			a := stm.UnaSqrt().(*rlt.StmUnaSqrt)
			tst.RltStmUnaSqrtNotZero(t, a)
			mnr := tst.NewStmMnr(ap)
			a.Sub(mnr.Rx, mnr.Id)
			tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
			var actr act.Actr
			vs := actr.RunHst(a.String())
			eHst := vs[len(vs)-1].(*hst.StmUnaSqrt)
			if eHst.Tmes != nil {
				mnr.StartFor(instr.Instr(), eHst.Tmes.Cnt())
				tst.TmesEql(t, eHst.Tmes, mnr.Tmes, "Tmes")
				tst.FltsEql(t, eHst.Vals, mnr.Vals, "Vals")
			}
			a.Unsub(mnr.Id)
			tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
			ap.Cls()
		})
	}
}
func TestRltStmSclAdd(t *testing.T) {
	for _, instr := range tst.RltPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := rlt.Oan()
			instr := instr(prv)
			inrvl := tst.RltInstrInrvlI(instr, 10)
			side := tst.RltInrvlSideBid(inrvl)
			stm := tst.RltSideStmRteLst(side)
			a := stm.SclAdd(flt.Flt(1.1)).(*rlt.StmSclAdd)
			tst.RltStmSclAddNotZero(t, a)
			mnr := tst.NewStmMnr(ap)
			a.Sub(mnr.Rx, mnr.Id)
			tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
			var actr act.Actr
			vs := actr.RunHst(a.String())
			eHst := vs[len(vs)-1].(*hst.StmSclAdd)
			if eHst.Tmes != nil {
				mnr.StartFor(instr.Instr(), eHst.Tmes.Cnt())
				tst.TmesEql(t, eHst.Tmes, mnr.Tmes, "Tmes")
				tst.FltsEql(t, eHst.Vals, mnr.Vals, "Vals")
			}
			a.Unsub(mnr.Id)
			tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
			ap.Cls()
		})
	}
}
func TestRltStmSclSub(t *testing.T) {
	for _, instr := range tst.RltPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := rlt.Oan()
			instr := instr(prv)
			inrvl := tst.RltInstrInrvlI(instr, 10)
			side := tst.RltInrvlSideBid(inrvl)
			stm := tst.RltSideStmRteLst(side)
			a := stm.SclSub(flt.Flt(1.1)).(*rlt.StmSclSub)
			tst.RltStmSclSubNotZero(t, a)
			mnr := tst.NewStmMnr(ap)
			a.Sub(mnr.Rx, mnr.Id)
			tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
			var actr act.Actr
			vs := actr.RunHst(a.String())
			eHst := vs[len(vs)-1].(*hst.StmSclSub)
			if eHst.Tmes != nil {
				mnr.StartFor(instr.Instr(), eHst.Tmes.Cnt())
				tst.TmesEql(t, eHst.Tmes, mnr.Tmes, "Tmes")
				tst.FltsEql(t, eHst.Vals, mnr.Vals, "Vals")
			}
			a.Unsub(mnr.Id)
			tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
			ap.Cls()
		})
	}
}
func TestRltStmSclMul(t *testing.T) {
	for _, instr := range tst.RltPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := rlt.Oan()
			instr := instr(prv)
			inrvl := tst.RltInstrInrvlI(instr, 10)
			side := tst.RltInrvlSideBid(inrvl)
			stm := tst.RltSideStmRteLst(side)
			a := stm.SclMul(flt.Flt(1.1)).(*rlt.StmSclMul)
			tst.RltStmSclMulNotZero(t, a)
			mnr := tst.NewStmMnr(ap)
			a.Sub(mnr.Rx, mnr.Id)
			tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
			var actr act.Actr
			vs := actr.RunHst(a.String())
			eHst := vs[len(vs)-1].(*hst.StmSclMul)
			if eHst.Tmes != nil {
				mnr.StartFor(instr.Instr(), eHst.Tmes.Cnt())
				tst.TmesEql(t, eHst.Tmes, mnr.Tmes, "Tmes")
				tst.FltsEql(t, eHst.Vals, mnr.Vals, "Vals")
			}
			a.Unsub(mnr.Id)
			tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
			ap.Cls()
		})
	}
}
func TestRltStmSclDiv(t *testing.T) {
	for _, instr := range tst.RltPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := rlt.Oan()
			instr := instr(prv)
			inrvl := tst.RltInstrInrvlI(instr, 10)
			side := tst.RltInrvlSideBid(inrvl)
			stm := tst.RltSideStmRteLst(side)
			a := stm.SclDiv(flt.Flt(1.1)).(*rlt.StmSclDiv)
			tst.RltStmSclDivNotZero(t, a)
			mnr := tst.NewStmMnr(ap)
			a.Sub(mnr.Rx, mnr.Id)
			tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
			var actr act.Actr
			vs := actr.RunHst(a.String())
			eHst := vs[len(vs)-1].(*hst.StmSclDiv)
			if eHst.Tmes != nil {
				mnr.StartFor(instr.Instr(), eHst.Tmes.Cnt())
				tst.TmesEql(t, eHst.Tmes, mnr.Tmes, "Tmes")
				tst.FltsEql(t, eHst.Vals, mnr.Vals, "Vals")
			}
			a.Unsub(mnr.Id)
			tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
			ap.Cls()
		})
	}
}
func TestRltStmSclRem(t *testing.T) {
	for _, instr := range tst.RltPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := rlt.Oan()
			instr := instr(prv)
			inrvl := tst.RltInstrInrvlI(instr, 10)
			side := tst.RltInrvlSideBid(inrvl)
			stm := tst.RltSideStmRteLst(side)
			a := stm.SclRem(flt.Flt(1.1)).(*rlt.StmSclRem)
			tst.RltStmSclRemNotZero(t, a)
			mnr := tst.NewStmMnr(ap)
			a.Sub(mnr.Rx, mnr.Id)
			tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
			var actr act.Actr
			vs := actr.RunHst(a.String())
			eHst := vs[len(vs)-1].(*hst.StmSclRem)
			if eHst.Tmes != nil {
				mnr.StartFor(instr.Instr(), eHst.Tmes.Cnt())
				tst.TmesEql(t, eHst.Tmes, mnr.Tmes, "Tmes")
				tst.FltsEql(t, eHst.Vals, mnr.Vals, "Vals")
			}
			a.Unsub(mnr.Id)
			tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
			ap.Cls()
		})
	}
}
func TestRltStmSclPow(t *testing.T) {
	for _, instr := range tst.RltPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := rlt.Oan()
			instr := instr(prv)
			inrvl := tst.RltInstrInrvlI(instr, 10)
			side := tst.RltInrvlSideBid(inrvl)
			stm := tst.RltSideStmRteLst(side)
			a := stm.SclPow(flt.Flt(1.1)).(*rlt.StmSclPow)
			tst.RltStmSclPowNotZero(t, a)
			mnr := tst.NewStmMnr(ap)
			a.Sub(mnr.Rx, mnr.Id)
			tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
			var actr act.Actr
			vs := actr.RunHst(a.String())
			eHst := vs[len(vs)-1].(*hst.StmSclPow)
			if eHst.Tmes != nil {
				mnr.StartFor(instr.Instr(), eHst.Tmes.Cnt())
				tst.TmesEql(t, eHst.Tmes, mnr.Tmes, "Tmes")
				tst.FltsEql(t, eHst.Vals, mnr.Vals, "Vals")
			}
			a.Unsub(mnr.Id)
			tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
			ap.Cls()
		})
	}
}
func TestRltStmSclMin(t *testing.T) {
	for _, instr := range tst.RltPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := rlt.Oan()
			instr := instr(prv)
			inrvl := tst.RltInstrInrvlI(instr, 10)
			side := tst.RltInrvlSideBid(inrvl)
			stm := tst.RltSideStmRteLst(side)
			a := stm.SclMin(flt.Flt(1.1)).(*rlt.StmSclMin)
			tst.RltStmSclMinNotZero(t, a)
			mnr := tst.NewStmMnr(ap)
			a.Sub(mnr.Rx, mnr.Id)
			tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
			var actr act.Actr
			vs := actr.RunHst(a.String())
			eHst := vs[len(vs)-1].(*hst.StmSclMin)
			if eHst.Tmes != nil {
				mnr.StartFor(instr.Instr(), eHst.Tmes.Cnt())
				tst.TmesEql(t, eHst.Tmes, mnr.Tmes, "Tmes")
				tst.FltsEql(t, eHst.Vals, mnr.Vals, "Vals")
			}
			a.Unsub(mnr.Id)
			tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
			ap.Cls()
		})
	}
}
func TestRltStmSclMax(t *testing.T) {
	for _, instr := range tst.RltPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := rlt.Oan()
			instr := instr(prv)
			inrvl := tst.RltInstrInrvlI(instr, 10)
			side := tst.RltInrvlSideBid(inrvl)
			stm := tst.RltSideStmRteLst(side)
			a := stm.SclMax(flt.Flt(1.1)).(*rlt.StmSclMax)
			tst.RltStmSclMaxNotZero(t, a)
			mnr := tst.NewStmMnr(ap)
			a.Sub(mnr.Rx, mnr.Id)
			tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
			var actr act.Actr
			vs := actr.RunHst(a.String())
			eHst := vs[len(vs)-1].(*hst.StmSclMax)
			if eHst.Tmes != nil {
				mnr.StartFor(instr.Instr(), eHst.Tmes.Cnt())
				tst.TmesEql(t, eHst.Tmes, mnr.Tmes, "Tmes")
				tst.FltsEql(t, eHst.Vals, mnr.Vals, "Vals")
			}
			a.Unsub(mnr.Id)
			tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
			ap.Cls()
		})
	}
}
func TestRltStmSelEql(t *testing.T) {
	for _, instr := range tst.RltPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := rlt.Oan()
			instr := instr(prv)
			inrvl := tst.RltInstrInrvlI(instr, 10)
			side := tst.RltInrvlSideBid(inrvl)
			stm := tst.RltSideStmRteLst(side)
			a := stm.SelEql(flt.Flt(1.3171)).(*rlt.StmSelEql)
			tst.RltStmSelEqlNotZero(t, a)
			mnr := tst.NewStmMnr(ap)
			a.Sub(mnr.Rx, mnr.Id)
			tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
			var actr act.Actr
			vs := actr.RunHst(a.String())
			eHst := vs[len(vs)-1].(*hst.StmSelEql)
			if eHst.Tmes != nil {
				mnr.StartFor(instr.Instr(), eHst.Tmes.Cnt())
				tst.TmesEql(t, eHst.Tmes, mnr.Tmes, "Tmes")
				tst.FltsEql(t, eHst.Vals, mnr.Vals, "Vals")
			}
			a.Unsub(mnr.Id)
			tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
			ap.Cls()
		})
	}
}
func TestRltStmSelNeq(t *testing.T) {
	for _, instr := range tst.RltPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := rlt.Oan()
			instr := instr(prv)
			inrvl := tst.RltInstrInrvlI(instr, 10)
			side := tst.RltInrvlSideBid(inrvl)
			stm := tst.RltSideStmRteLst(side)
			a := stm.SelNeq(flt.Flt(1.3171)).(*rlt.StmSelNeq)
			tst.RltStmSelNeqNotZero(t, a)
			mnr := tst.NewStmMnr(ap)
			a.Sub(mnr.Rx, mnr.Id)
			tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
			var actr act.Actr
			vs := actr.RunHst(a.String())
			eHst := vs[len(vs)-1].(*hst.StmSelNeq)
			if eHst.Tmes != nil {
				mnr.StartFor(instr.Instr(), eHst.Tmes.Cnt())
				tst.TmesEql(t, eHst.Tmes, mnr.Tmes, "Tmes")
				tst.FltsEql(t, eHst.Vals, mnr.Vals, "Vals")
			}
			a.Unsub(mnr.Id)
			tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
			ap.Cls()
		})
	}
}
func TestRltStmSelLss(t *testing.T) {
	for _, instr := range tst.RltPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := rlt.Oan()
			instr := instr(prv)
			inrvl := tst.RltInstrInrvlI(instr, 10)
			side := tst.RltInrvlSideBid(inrvl)
			stm := tst.RltSideStmRteLst(side)
			a := stm.SelLss(flt.Flt(1.3171)).(*rlt.StmSelLss)
			tst.RltStmSelLssNotZero(t, a)
			mnr := tst.NewStmMnr(ap)
			a.Sub(mnr.Rx, mnr.Id)
			tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
			var actr act.Actr
			vs := actr.RunHst(a.String())
			eHst := vs[len(vs)-1].(*hst.StmSelLss)
			if eHst.Tmes != nil {
				mnr.StartFor(instr.Instr(), eHst.Tmes.Cnt())
				tst.TmesEql(t, eHst.Tmes, mnr.Tmes, "Tmes")
				tst.FltsEql(t, eHst.Vals, mnr.Vals, "Vals")
			}
			a.Unsub(mnr.Id)
			tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
			ap.Cls()
		})
	}
}
func TestRltStmSelGtr(t *testing.T) {
	for _, instr := range tst.RltPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := rlt.Oan()
			instr := instr(prv)
			inrvl := tst.RltInstrInrvlI(instr, 10)
			side := tst.RltInrvlSideBid(inrvl)
			stm := tst.RltSideStmRteLst(side)
			a := stm.SelGtr(flt.Flt(1.3171)).(*rlt.StmSelGtr)
			tst.RltStmSelGtrNotZero(t, a)
			mnr := tst.NewStmMnr(ap)
			a.Sub(mnr.Rx, mnr.Id)
			tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
			var actr act.Actr
			vs := actr.RunHst(a.String())
			eHst := vs[len(vs)-1].(*hst.StmSelGtr)
			if eHst.Tmes != nil {
				mnr.StartFor(instr.Instr(), eHst.Tmes.Cnt())
				tst.TmesEql(t, eHst.Tmes, mnr.Tmes, "Tmes")
				tst.FltsEql(t, eHst.Vals, mnr.Vals, "Vals")
			}
			a.Unsub(mnr.Id)
			tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
			ap.Cls()
		})
	}
}
func TestRltStmSelLeq(t *testing.T) {
	for _, instr := range tst.RltPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := rlt.Oan()
			instr := instr(prv)
			inrvl := tst.RltInstrInrvlI(instr, 10)
			side := tst.RltInrvlSideBid(inrvl)
			stm := tst.RltSideStmRteLst(side)
			a := stm.SelLeq(flt.Flt(1.3171)).(*rlt.StmSelLeq)
			tst.RltStmSelLeqNotZero(t, a)
			mnr := tst.NewStmMnr(ap)
			a.Sub(mnr.Rx, mnr.Id)
			tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
			var actr act.Actr
			vs := actr.RunHst(a.String())
			eHst := vs[len(vs)-1].(*hst.StmSelLeq)
			if eHst.Tmes != nil {
				mnr.StartFor(instr.Instr(), eHst.Tmes.Cnt())
				tst.TmesEql(t, eHst.Tmes, mnr.Tmes, "Tmes")
				tst.FltsEql(t, eHst.Vals, mnr.Vals, "Vals")
			}
			a.Unsub(mnr.Id)
			tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
			ap.Cls()
		})
	}
}
func TestRltStmSelGeq(t *testing.T) {
	for _, instr := range tst.RltPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := rlt.Oan()
			instr := instr(prv)
			inrvl := tst.RltInstrInrvlI(instr, 10)
			side := tst.RltInrvlSideBid(inrvl)
			stm := tst.RltSideStmRteLst(side)
			a := stm.SelGeq(flt.Flt(1.3171)).(*rlt.StmSelGeq)
			tst.RltStmSelGeqNotZero(t, a)
			mnr := tst.NewStmMnr(ap)
			a.Sub(mnr.Rx, mnr.Id)
			tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
			var actr act.Actr
			vs := actr.RunHst(a.String())
			eHst := vs[len(vs)-1].(*hst.StmSelGeq)
			if eHst.Tmes != nil {
				mnr.StartFor(instr.Instr(), eHst.Tmes.Cnt())
				tst.TmesEql(t, eHst.Tmes, mnr.Tmes, "Tmes")
				tst.FltsEql(t, eHst.Vals, mnr.Vals, "Vals")
			}
			a.Unsub(mnr.Id)
			tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
			ap.Cls()
		})
	}
}
func TestRltStmAggFst(t *testing.T) {
	for _, instr := range tst.RltPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := rlt.Oan()
			instr := instr(prv)
			inrvl := tst.RltInstrInrvlI(instr, 10)
			side := tst.RltInrvlSideBid(inrvl)
			stm := tst.RltSideStmRteLst(side)
			a := stm.AggFst(unt.Unt(2)).(*rlt.StmAggFst)
			tst.RltStmAggFstNotZero(t, a)
			mnr := tst.NewStmMnr(ap)
			a.Sub(mnr.Rx, mnr.Id)
			tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
			var actr act.Actr
			vs := actr.RunHst(a.String())
			eHst := vs[len(vs)-1].(*hst.StmAggFst)
			if eHst.Tmes != nil {
				mnr.StartFor(instr.Instr(), eHst.Tmes.Cnt())
				tst.TmesEql(t, eHst.Tmes, mnr.Tmes, "Tmes")
				tst.FltsEql(t, eHst.Vals, mnr.Vals, "Vals")
			}
			a.Unsub(mnr.Id)
			tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
			ap.Cls()
		})
	}
}
func TestRltStmAggLst(t *testing.T) {
	for _, instr := range tst.RltPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := rlt.Oan()
			instr := instr(prv)
			inrvl := tst.RltInstrInrvlI(instr, 10)
			side := tst.RltInrvlSideBid(inrvl)
			stm := tst.RltSideStmRteLst(side)
			a := stm.AggLst(unt.Unt(2)).(*rlt.StmAggLst)
			tst.RltStmAggLstNotZero(t, a)
			mnr := tst.NewStmMnr(ap)
			a.Sub(mnr.Rx, mnr.Id)
			tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
			var actr act.Actr
			vs := actr.RunHst(a.String())
			eHst := vs[len(vs)-1].(*hst.StmAggLst)
			if eHst.Tmes != nil {
				mnr.StartFor(instr.Instr(), eHst.Tmes.Cnt())
				tst.TmesEql(t, eHst.Tmes, mnr.Tmes, "Tmes")
				tst.FltsEql(t, eHst.Vals, mnr.Vals, "Vals")
			}
			a.Unsub(mnr.Id)
			tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
			ap.Cls()
		})
	}
}
func TestRltStmAggSum(t *testing.T) {
	for _, instr := range tst.RltPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := rlt.Oan()
			instr := instr(prv)
			inrvl := tst.RltInstrInrvlI(instr, 10)
			side := tst.RltInrvlSideBid(inrvl)
			stm := tst.RltSideStmRteLst(side)
			a := stm.AggSum(unt.Unt(2)).(*rlt.StmAggSum)
			tst.RltStmAggSumNotZero(t, a)
			mnr := tst.NewStmMnr(ap)
			a.Sub(mnr.Rx, mnr.Id)
			tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
			var actr act.Actr
			vs := actr.RunHst(a.String())
			eHst := vs[len(vs)-1].(*hst.StmAggSum)
			if eHst.Tmes != nil {
				mnr.StartFor(instr.Instr(), eHst.Tmes.Cnt())
				tst.TmesEql(t, eHst.Tmes, mnr.Tmes, "Tmes")
				tst.FltsEql(t, eHst.Vals, mnr.Vals, "Vals")
			}
			a.Unsub(mnr.Id)
			tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
			ap.Cls()
		})
	}
}
func TestRltStmAggPrd(t *testing.T) {
	for _, instr := range tst.RltPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := rlt.Oan()
			instr := instr(prv)
			inrvl := tst.RltInstrInrvlI(instr, 10)
			side := tst.RltInrvlSideBid(inrvl)
			stm := tst.RltSideStmRteLst(side)
			a := stm.AggPrd(unt.Unt(2)).(*rlt.StmAggPrd)
			tst.RltStmAggPrdNotZero(t, a)
			mnr := tst.NewStmMnr(ap)
			a.Sub(mnr.Rx, mnr.Id)
			tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
			var actr act.Actr
			vs := actr.RunHst(a.String())
			eHst := vs[len(vs)-1].(*hst.StmAggPrd)
			if eHst.Tmes != nil {
				mnr.StartFor(instr.Instr(), eHst.Tmes.Cnt())
				tst.TmesEql(t, eHst.Tmes, mnr.Tmes, "Tmes")
				tst.FltsEql(t, eHst.Vals, mnr.Vals, "Vals")
			}
			a.Unsub(mnr.Id)
			tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
			ap.Cls()
		})
	}
}
func TestRltStmAggMin(t *testing.T) {
	for _, instr := range tst.RltPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := rlt.Oan()
			instr := instr(prv)
			inrvl := tst.RltInstrInrvlI(instr, 10)
			side := tst.RltInrvlSideBid(inrvl)
			stm := tst.RltSideStmRteLst(side)
			a := stm.AggMin(unt.Unt(2)).(*rlt.StmAggMin)
			tst.RltStmAggMinNotZero(t, a)
			mnr := tst.NewStmMnr(ap)
			a.Sub(mnr.Rx, mnr.Id)
			tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
			var actr act.Actr
			vs := actr.RunHst(a.String())
			eHst := vs[len(vs)-1].(*hst.StmAggMin)
			if eHst.Tmes != nil {
				mnr.StartFor(instr.Instr(), eHst.Tmes.Cnt())
				tst.TmesEql(t, eHst.Tmes, mnr.Tmes, "Tmes")
				tst.FltsEql(t, eHst.Vals, mnr.Vals, "Vals")
			}
			a.Unsub(mnr.Id)
			tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
			ap.Cls()
		})
	}
}
func TestRltStmAggMax(t *testing.T) {
	for _, instr := range tst.RltPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := rlt.Oan()
			instr := instr(prv)
			inrvl := tst.RltInstrInrvlI(instr, 10)
			side := tst.RltInrvlSideBid(inrvl)
			stm := tst.RltSideStmRteLst(side)
			a := stm.AggMax(unt.Unt(2)).(*rlt.StmAggMax)
			tst.RltStmAggMaxNotZero(t, a)
			mnr := tst.NewStmMnr(ap)
			a.Sub(mnr.Rx, mnr.Id)
			tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
			var actr act.Actr
			vs := actr.RunHst(a.String())
			eHst := vs[len(vs)-1].(*hst.StmAggMax)
			if eHst.Tmes != nil {
				mnr.StartFor(instr.Instr(), eHst.Tmes.Cnt())
				tst.TmesEql(t, eHst.Tmes, mnr.Tmes, "Tmes")
				tst.FltsEql(t, eHst.Vals, mnr.Vals, "Vals")
			}
			a.Unsub(mnr.Id)
			tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
			ap.Cls()
		})
	}
}
func TestRltStmAggMid(t *testing.T) {
	for _, instr := range tst.RltPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := rlt.Oan()
			instr := instr(prv)
			inrvl := tst.RltInstrInrvlI(instr, 10)
			side := tst.RltInrvlSideBid(inrvl)
			stm := tst.RltSideStmRteLst(side)
			a := stm.AggMid(unt.Unt(2)).(*rlt.StmAggMid)
			tst.RltStmAggMidNotZero(t, a)
			mnr := tst.NewStmMnr(ap)
			a.Sub(mnr.Rx, mnr.Id)
			tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
			var actr act.Actr
			vs := actr.RunHst(a.String())
			eHst := vs[len(vs)-1].(*hst.StmAggMid)
			if eHst.Tmes != nil {
				mnr.StartFor(instr.Instr(), eHst.Tmes.Cnt())
				tst.TmesEql(t, eHst.Tmes, mnr.Tmes, "Tmes")
				tst.FltsEql(t, eHst.Vals, mnr.Vals, "Vals")
			}
			a.Unsub(mnr.Id)
			tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
			ap.Cls()
		})
	}
}
func TestRltStmAggMdn(t *testing.T) {
	for _, instr := range tst.RltPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := rlt.Oan()
			instr := instr(prv)
			inrvl := tst.RltInstrInrvlI(instr, 10)
			side := tst.RltInrvlSideBid(inrvl)
			stm := tst.RltSideStmRteLst(side)
			a := stm.AggMdn(unt.Unt(2)).(*rlt.StmAggMdn)
			tst.RltStmAggMdnNotZero(t, a)
			mnr := tst.NewStmMnr(ap)
			a.Sub(mnr.Rx, mnr.Id)
			tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
			var actr act.Actr
			vs := actr.RunHst(a.String())
			eHst := vs[len(vs)-1].(*hst.StmAggMdn)
			if eHst.Tmes != nil {
				mnr.StartFor(instr.Instr(), eHst.Tmes.Cnt())
				tst.TmesEql(t, eHst.Tmes, mnr.Tmes, "Tmes")
				tst.FltsEql(t, eHst.Vals, mnr.Vals, "Vals")
			}
			a.Unsub(mnr.Id)
			tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
			ap.Cls()
		})
	}
}
func TestRltStmAggSma(t *testing.T) {
	for _, instr := range tst.RltPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := rlt.Oan()
			instr := instr(prv)
			inrvl := tst.RltInstrInrvlI(instr, 10)
			side := tst.RltInrvlSideBid(inrvl)
			stm := tst.RltSideStmRteLst(side)
			a := stm.AggSma(unt.Unt(2)).(*rlt.StmAggSma)
			tst.RltStmAggSmaNotZero(t, a)
			mnr := tst.NewStmMnr(ap)
			a.Sub(mnr.Rx, mnr.Id)
			tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
			var actr act.Actr
			vs := actr.RunHst(a.String())
			eHst := vs[len(vs)-1].(*hst.StmAggSma)
			if eHst.Tmes != nil {
				mnr.StartFor(instr.Instr(), eHst.Tmes.Cnt())
				tst.TmesEql(t, eHst.Tmes, mnr.Tmes, "Tmes")
				tst.FltsEql(t, eHst.Vals, mnr.Vals, "Vals")
			}
			a.Unsub(mnr.Id)
			tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
			ap.Cls()
		})
	}
}
func TestRltStmAggGma(t *testing.T) {
	for _, instr := range tst.RltPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := rlt.Oan()
			instr := instr(prv)
			inrvl := tst.RltInstrInrvlI(instr, 10)
			side := tst.RltInrvlSideBid(inrvl)
			stm := tst.RltSideStmRteLst(side)
			a := stm.AggGma(unt.Unt(2)).(*rlt.StmAggGma)
			tst.RltStmAggGmaNotZero(t, a)
			mnr := tst.NewStmMnr(ap)
			a.Sub(mnr.Rx, mnr.Id)
			tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
			var actr act.Actr
			vs := actr.RunHst(a.String())
			eHst := vs[len(vs)-1].(*hst.StmAggGma)
			if eHst.Tmes != nil {
				mnr.StartFor(instr.Instr(), eHst.Tmes.Cnt())
				tst.TmesEql(t, eHst.Tmes, mnr.Tmes, "Tmes")
				tst.FltsEql(t, eHst.Vals, mnr.Vals, "Vals")
			}
			a.Unsub(mnr.Id)
			tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
			ap.Cls()
		})
	}
}
func TestRltStmAggWma(t *testing.T) {
	for _, instr := range tst.RltPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := rlt.Oan()
			instr := instr(prv)
			inrvl := tst.RltInstrInrvlI(instr, 10)
			side := tst.RltInrvlSideBid(inrvl)
			stm := tst.RltSideStmRteLst(side)
			a := stm.AggWma(unt.Unt(2)).(*rlt.StmAggWma)
			tst.RltStmAggWmaNotZero(t, a)
			mnr := tst.NewStmMnr(ap)
			a.Sub(mnr.Rx, mnr.Id)
			tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
			var actr act.Actr
			vs := actr.RunHst(a.String())
			eHst := vs[len(vs)-1].(*hst.StmAggWma)
			if eHst.Tmes != nil {
				mnr.StartFor(instr.Instr(), eHst.Tmes.Cnt())
				tst.TmesEql(t, eHst.Tmes, mnr.Tmes, "Tmes")
				tst.FltsEql(t, eHst.Vals, mnr.Vals, "Vals")
			}
			a.Unsub(mnr.Id)
			tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
			ap.Cls()
		})
	}
}
func TestRltStmAggRsi(t *testing.T) {
	for _, instr := range tst.RltPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := rlt.Oan()
			instr := instr(prv)
			inrvl := tst.RltInstrInrvlI(instr, 10)
			side := tst.RltInrvlSideBid(inrvl)
			stm := tst.RltSideStmRteLst(side)
			a := stm.AggRsi(unt.Unt(2)).(*rlt.StmAggRsi)
			tst.RltStmAggRsiNotZero(t, a)
			mnr := tst.NewStmMnr(ap)
			a.Sub(mnr.Rx, mnr.Id)
			tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
			var actr act.Actr
			vs := actr.RunHst(a.String())
			eHst := vs[len(vs)-1].(*hst.StmAggRsi)
			if eHst.Tmes != nil {
				mnr.StartFor(instr.Instr(), eHst.Tmes.Cnt())
				tst.TmesEql(t, eHst.Tmes, mnr.Tmes, "Tmes")
				tst.FltsEql(t, eHst.Vals, mnr.Vals, "Vals")
			}
			a.Unsub(mnr.Id)
			tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
			ap.Cls()
		})
	}
}
func TestRltStmAggWrsi(t *testing.T) {
	for _, instr := range tst.RltPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := rlt.Oan()
			instr := instr(prv)
			inrvl := tst.RltInstrInrvlI(instr, 10)
			side := tst.RltInrvlSideBid(inrvl)
			stm := tst.RltSideStmRteLst(side)
			a := stm.AggWrsi(unt.Unt(2)).(*rlt.StmAggWrsi)
			tst.RltStmAggWrsiNotZero(t, a)
			mnr := tst.NewStmMnr(ap)
			a.Sub(mnr.Rx, mnr.Id)
			tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
			var actr act.Actr
			vs := actr.RunHst(a.String())
			eHst := vs[len(vs)-1].(*hst.StmAggWrsi)
			if eHst.Tmes != nil {
				mnr.StartFor(instr.Instr(), eHst.Tmes.Cnt())
				tst.TmesEql(t, eHst.Tmes, mnr.Tmes, "Tmes")
				tst.FltsEql(t, eHst.Vals, mnr.Vals, "Vals")
			}
			a.Unsub(mnr.Id)
			tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
			ap.Cls()
		})
	}
}
func TestRltStmAggAlma(t *testing.T) {
	for _, instr := range tst.RltPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := rlt.Oan()
			instr := instr(prv)
			inrvl := tst.RltInstrInrvlI(instr, 10)
			side := tst.RltInrvlSideBid(inrvl)
			stm := tst.RltSideStmRteLst(side)
			a := stm.AggAlma(unt.Unt(2)).(*rlt.StmAggAlma)
			tst.RltStmAggAlmaNotZero(t, a)
			mnr := tst.NewStmMnr(ap)
			a.Sub(mnr.Rx, mnr.Id)
			tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
			var actr act.Actr
			vs := actr.RunHst(a.String())
			eHst := vs[len(vs)-1].(*hst.StmAggAlma)
			if eHst.Tmes != nil {
				mnr.StartFor(instr.Instr(), eHst.Tmes.Cnt())
				tst.TmesEql(t, eHst.Tmes, mnr.Tmes, "Tmes")
				tst.FltsEql(t, eHst.Vals, mnr.Vals, "Vals")
			}
			a.Unsub(mnr.Id)
			tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
			ap.Cls()
		})
	}
}
func TestRltStmAggVrnc(t *testing.T) {
	for _, instr := range tst.RltPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := rlt.Oan()
			instr := instr(prv)
			inrvl := tst.RltInstrInrvlI(instr, 10)
			side := tst.RltInrvlSideBid(inrvl)
			stm := tst.RltSideStmRteLst(side)
			a := stm.AggVrnc(unt.Unt(2)).(*rlt.StmAggVrnc)
			tst.RltStmAggVrncNotZero(t, a)
			mnr := tst.NewStmMnr(ap)
			a.Sub(mnr.Rx, mnr.Id)
			tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
			var actr act.Actr
			vs := actr.RunHst(a.String())
			eHst := vs[len(vs)-1].(*hst.StmAggVrnc)
			if eHst.Tmes != nil {
				mnr.StartFor(instr.Instr(), eHst.Tmes.Cnt())
				tst.TmesEql(t, eHst.Tmes, mnr.Tmes, "Tmes")
				tst.FltsEql(t, eHst.Vals, mnr.Vals, "Vals")
			}
			a.Unsub(mnr.Id)
			tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
			ap.Cls()
		})
	}
}
func TestRltStmAggStd(t *testing.T) {
	for _, instr := range tst.RltPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := rlt.Oan()
			instr := instr(prv)
			inrvl := tst.RltInstrInrvlI(instr, 10)
			side := tst.RltInrvlSideBid(inrvl)
			stm := tst.RltSideStmRteLst(side)
			a := stm.AggStd(unt.Unt(2)).(*rlt.StmAggStd)
			tst.RltStmAggStdNotZero(t, a)
			mnr := tst.NewStmMnr(ap)
			a.Sub(mnr.Rx, mnr.Id)
			tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
			var actr act.Actr
			vs := actr.RunHst(a.String())
			eHst := vs[len(vs)-1].(*hst.StmAggStd)
			if eHst.Tmes != nil {
				mnr.StartFor(instr.Instr(), eHst.Tmes.Cnt())
				tst.TmesEql(t, eHst.Tmes, mnr.Tmes, "Tmes")
				tst.FltsEql(t, eHst.Vals, mnr.Vals, "Vals")
			}
			a.Unsub(mnr.Id)
			tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
			ap.Cls()
		})
	}
}
func TestRltStmAggRngFul(t *testing.T) {
	for _, instr := range tst.RltPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := rlt.Oan()
			instr := instr(prv)
			inrvl := tst.RltInstrInrvlI(instr, 10)
			side := tst.RltInrvlSideBid(inrvl)
			stm := tst.RltSideStmRteLst(side)
			a := stm.AggRngFul(unt.Unt(2)).(*rlt.StmAggRngFul)
			tst.RltStmAggRngFulNotZero(t, a)
			mnr := tst.NewStmMnr(ap)
			a.Sub(mnr.Rx, mnr.Id)
			tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
			var actr act.Actr
			vs := actr.RunHst(a.String())
			eHst := vs[len(vs)-1].(*hst.StmAggRngFul)
			if eHst.Tmes != nil {
				mnr.StartFor(instr.Instr(), eHst.Tmes.Cnt())
				tst.TmesEql(t, eHst.Tmes, mnr.Tmes, "Tmes")
				tst.FltsEql(t, eHst.Vals, mnr.Vals, "Vals")
			}
			a.Unsub(mnr.Id)
			tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
			ap.Cls()
		})
	}
}
func TestRltStmAggRngLst(t *testing.T) {
	for _, instr := range tst.RltPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := rlt.Oan()
			instr := instr(prv)
			inrvl := tst.RltInstrInrvlI(instr, 10)
			side := tst.RltInrvlSideBid(inrvl)
			stm := tst.RltSideStmRteLst(side)
			a := stm.AggRngLst(unt.Unt(2)).(*rlt.StmAggRngLst)
			tst.RltStmAggRngLstNotZero(t, a)
			mnr := tst.NewStmMnr(ap)
			a.Sub(mnr.Rx, mnr.Id)
			tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
			var actr act.Actr
			vs := actr.RunHst(a.String())
			eHst := vs[len(vs)-1].(*hst.StmAggRngLst)
			if eHst.Tmes != nil {
				mnr.StartFor(instr.Instr(), eHst.Tmes.Cnt())
				tst.TmesEql(t, eHst.Tmes, mnr.Tmes, "Tmes")
				tst.FltsEql(t, eHst.Vals, mnr.Vals, "Vals")
			}
			a.Unsub(mnr.Id)
			tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
			ap.Cls()
		})
	}
}
func TestRltStmAggProLst(t *testing.T) {
	for _, instr := range tst.RltPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := rlt.Oan()
			instr := instr(prv)
			inrvl := tst.RltInstrInrvlI(instr, 10)
			side := tst.RltInrvlSideBid(inrvl)
			stm := tst.RltSideStmRteLst(side)
			a := stm.AggProLst(unt.Unt(2)).(*rlt.StmAggProLst)
			tst.RltStmAggProLstNotZero(t, a)
			mnr := tst.NewStmMnr(ap)
			a.Sub(mnr.Rx, mnr.Id)
			tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
			var actr act.Actr
			vs := actr.RunHst(a.String())
			eHst := vs[len(vs)-1].(*hst.StmAggProLst)
			if eHst.Tmes != nil {
				mnr.StartFor(instr.Instr(), eHst.Tmes.Cnt())
				tst.TmesEql(t, eHst.Tmes, mnr.Tmes, "Tmes")
				tst.FltsEql(t, eHst.Vals, mnr.Vals, "Vals")
			}
			a.Unsub(mnr.Id)
			tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
			ap.Cls()
		})
	}
}
func TestRltStmAggProSma(t *testing.T) {
	for _, instr := range tst.RltPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := rlt.Oan()
			instr := instr(prv)
			inrvl := tst.RltInstrInrvlI(instr, 10)
			side := tst.RltInrvlSideBid(inrvl)
			stm := tst.RltSideStmRteLst(side)
			a := stm.AggProSma(unt.Unt(2)).(*rlt.StmAggProSma)
			tst.RltStmAggProSmaNotZero(t, a)
			mnr := tst.NewStmMnr(ap)
			a.Sub(mnr.Rx, mnr.Id)
			tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
			var actr act.Actr
			vs := actr.RunHst(a.String())
			eHst := vs[len(vs)-1].(*hst.StmAggProSma)
			if eHst.Tmes != nil {
				mnr.StartFor(instr.Instr(), eHst.Tmes.Cnt())
				tst.TmesEql(t, eHst.Tmes, mnr.Tmes, "Tmes")
				tst.FltsEql(t, eHst.Vals, mnr.Vals, "Vals")
			}
			a.Unsub(mnr.Id)
			tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
			ap.Cls()
		})
	}
}
func TestRltStmAggProAlma(t *testing.T) {
	for _, instr := range tst.RltPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := rlt.Oan()
			instr := instr(prv)
			inrvl := tst.RltInstrInrvlI(instr, 10)
			side := tst.RltInrvlSideBid(inrvl)
			stm := tst.RltSideStmRteLst(side)
			a := stm.AggProAlma(unt.Unt(2)).(*rlt.StmAggProAlma)
			tst.RltStmAggProAlmaNotZero(t, a)
			mnr := tst.NewStmMnr(ap)
			a.Sub(mnr.Rx, mnr.Id)
			tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
			var actr act.Actr
			vs := actr.RunHst(a.String())
			eHst := vs[len(vs)-1].(*hst.StmAggProAlma)
			if eHst.Tmes != nil {
				mnr.StartFor(instr.Instr(), eHst.Tmes.Cnt())
				tst.TmesEql(t, eHst.Tmes, mnr.Tmes, "Tmes")
				tst.FltsEql(t, eHst.Vals, mnr.Vals, "Vals")
			}
			a.Unsub(mnr.Id)
			tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
			ap.Cls()
		})
	}
}
func TestRltStmAggEma(t *testing.T) {
	for _, instr := range tst.RltPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := rlt.Oan()
			instr := instr(prv)
			inrvl := tst.RltInstrInrvlI(instr, 10)
			side := tst.RltInrvlSideBid(inrvl)
			stm := tst.RltSideStmRteLst(side)
			a := stm.AggEma(unt.Unt(9)).(*rlt.StmAggEma)
			tst.RltStmAggEmaNotZero(t, a)
			mnr := tst.NewStmMnr(ap)
			a.Sub(mnr.Rx, mnr.Id)
			tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
			var actr act.Actr
			vs := actr.RunHst(a.String())
			eHst := vs[len(vs)-1].(*hst.StmAggEma)
			if eHst.Tmes != nil {
				mnr.StartFor(instr.Instr(), eHst.Tmes.Cnt())
				tst.TmesEql(t, eHst.Tmes, mnr.Tmes, "Tmes")
				tst.FltsEql(t, eHst.Vals, mnr.Vals, "Vals")
			}
			a.Unsub(mnr.Id)
			tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
			ap.Cls()
		})
	}
}
func TestRltStmInrAdd(t *testing.T) {
	for _, instr := range tst.RltPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := rlt.Oan()
			instr := instr(prv)
			inrvl := tst.RltInstrInrvlI(instr, 10)
			side := tst.RltInrvlSideBid(inrvl)
			stm := tst.RltSideStmRteLst(side)
			a := stm.InrAdd(unt.Unt(1)).(*rlt.StmInrAdd)
			tst.RltStmInrAddNotZero(t, a)
			mnr := tst.NewStmMnr(ap)
			a.Sub(mnr.Rx, mnr.Id)
			tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
			var actr act.Actr
			vs := actr.RunHst(a.String())
			eHst := vs[len(vs)-1].(*hst.StmInrAdd)
			if eHst.Tmes != nil {
				mnr.StartFor(instr.Instr(), eHst.Tmes.Cnt())
				tst.TmesEql(t, eHst.Tmes, mnr.Tmes, "Tmes")
				tst.FltsEql(t, eHst.Vals, mnr.Vals, "Vals")
			}
			a.Unsub(mnr.Id)
			tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
			ap.Cls()
		})
	}
}
func TestRltStmInrSub(t *testing.T) {
	for _, instr := range tst.RltPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := rlt.Oan()
			instr := instr(prv)
			inrvl := tst.RltInstrInrvlI(instr, 10)
			side := tst.RltInrvlSideBid(inrvl)
			stm := tst.RltSideStmRteLst(side)
			a := stm.InrSub(unt.Unt(1)).(*rlt.StmInrSub)
			tst.RltStmInrSubNotZero(t, a)
			mnr := tst.NewStmMnr(ap)
			a.Sub(mnr.Rx, mnr.Id)
			tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
			var actr act.Actr
			vs := actr.RunHst(a.String())
			eHst := vs[len(vs)-1].(*hst.StmInrSub)
			if eHst.Tmes != nil {
				mnr.StartFor(instr.Instr(), eHst.Tmes.Cnt())
				tst.TmesEql(t, eHst.Tmes, mnr.Tmes, "Tmes")
				tst.FltsEql(t, eHst.Vals, mnr.Vals, "Vals")
			}
			a.Unsub(mnr.Id)
			tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
			ap.Cls()
		})
	}
}
func TestRltStmInrMul(t *testing.T) {
	for _, instr := range tst.RltPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := rlt.Oan()
			instr := instr(prv)
			inrvl := tst.RltInstrInrvlI(instr, 10)
			side := tst.RltInrvlSideBid(inrvl)
			stm := tst.RltSideStmRteLst(side)
			a := stm.InrMul(unt.Unt(1)).(*rlt.StmInrMul)
			tst.RltStmInrMulNotZero(t, a)
			mnr := tst.NewStmMnr(ap)
			a.Sub(mnr.Rx, mnr.Id)
			tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
			var actr act.Actr
			vs := actr.RunHst(a.String())
			eHst := vs[len(vs)-1].(*hst.StmInrMul)
			if eHst.Tmes != nil {
				mnr.StartFor(instr.Instr(), eHst.Tmes.Cnt())
				tst.TmesEql(t, eHst.Tmes, mnr.Tmes, "Tmes")
				tst.FltsEql(t, eHst.Vals, mnr.Vals, "Vals")
			}
			a.Unsub(mnr.Id)
			tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
			ap.Cls()
		})
	}
}
func TestRltStmInrDiv(t *testing.T) {
	for _, instr := range tst.RltPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := rlt.Oan()
			instr := instr(prv)
			inrvl := tst.RltInstrInrvlI(instr, 10)
			side := tst.RltInrvlSideBid(inrvl)
			stm := tst.RltSideStmRteLst(side)
			a := stm.InrDiv(unt.Unt(1)).(*rlt.StmInrDiv)
			tst.RltStmInrDivNotZero(t, a)
			mnr := tst.NewStmMnr(ap)
			a.Sub(mnr.Rx, mnr.Id)
			tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
			var actr act.Actr
			vs := actr.RunHst(a.String())
			eHst := vs[len(vs)-1].(*hst.StmInrDiv)
			if eHst.Tmes != nil {
				mnr.StartFor(instr.Instr(), eHst.Tmes.Cnt())
				tst.TmesEql(t, eHst.Tmes, mnr.Tmes, "Tmes")
				tst.FltsEql(t, eHst.Vals, mnr.Vals, "Vals")
			}
			a.Unsub(mnr.Id)
			tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
			ap.Cls()
		})
	}
}
func TestRltStmInrRem(t *testing.T) {
	for _, instr := range tst.RltPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := rlt.Oan()
			instr := instr(prv)
			inrvl := tst.RltInstrInrvlI(instr, 10)
			side := tst.RltInrvlSideBid(inrvl)
			stm := tst.RltSideStmRteLst(side)
			a := stm.InrRem(unt.Unt(1)).(*rlt.StmInrRem)
			tst.RltStmInrRemNotZero(t, a)
			mnr := tst.NewStmMnr(ap)
			a.Sub(mnr.Rx, mnr.Id)
			tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
			var actr act.Actr
			vs := actr.RunHst(a.String())
			eHst := vs[len(vs)-1].(*hst.StmInrRem)
			if eHst.Tmes != nil {
				mnr.StartFor(instr.Instr(), eHst.Tmes.Cnt())
				tst.TmesEql(t, eHst.Tmes, mnr.Tmes, "Tmes")
				tst.FltsEql(t, eHst.Vals, mnr.Vals, "Vals")
			}
			a.Unsub(mnr.Id)
			tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
			ap.Cls()
		})
	}
}
func TestRltStmInrPow(t *testing.T) {
	for _, instr := range tst.RltPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := rlt.Oan()
			instr := instr(prv)
			inrvl := tst.RltInstrInrvlI(instr, 10)
			side := tst.RltInrvlSideBid(inrvl)
			stm := tst.RltSideStmRteLst(side)
			a := stm.InrPow(unt.Unt(1)).(*rlt.StmInrPow)
			tst.RltStmInrPowNotZero(t, a)
			mnr := tst.NewStmMnr(ap)
			a.Sub(mnr.Rx, mnr.Id)
			tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
			var actr act.Actr
			vs := actr.RunHst(a.String())
			eHst := vs[len(vs)-1].(*hst.StmInrPow)
			if eHst.Tmes != nil {
				mnr.StartFor(instr.Instr(), eHst.Tmes.Cnt())
				tst.TmesEql(t, eHst.Tmes, mnr.Tmes, "Tmes")
				tst.FltsEql(t, eHst.Vals, mnr.Vals, "Vals")
			}
			a.Unsub(mnr.Id)
			tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
			ap.Cls()
		})
	}
}
func TestRltStmInrMin(t *testing.T) {
	for _, instr := range tst.RltPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := rlt.Oan()
			instr := instr(prv)
			inrvl := tst.RltInstrInrvlI(instr, 10)
			side := tst.RltInrvlSideBid(inrvl)
			stm := tst.RltSideStmRteLst(side)
			a := stm.InrMin(unt.Unt(1)).(*rlt.StmInrMin)
			tst.RltStmInrMinNotZero(t, a)
			mnr := tst.NewStmMnr(ap)
			a.Sub(mnr.Rx, mnr.Id)
			tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
			var actr act.Actr
			vs := actr.RunHst(a.String())
			eHst := vs[len(vs)-1].(*hst.StmInrMin)
			if eHst.Tmes != nil {
				mnr.StartFor(instr.Instr(), eHst.Tmes.Cnt())
				tst.TmesEql(t, eHst.Tmes, mnr.Tmes, "Tmes")
				tst.FltsEql(t, eHst.Vals, mnr.Vals, "Vals")
			}
			a.Unsub(mnr.Id)
			tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
			ap.Cls()
		})
	}
}
func TestRltStmInrMax(t *testing.T) {
	for _, instr := range tst.RltPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := rlt.Oan()
			instr := instr(prv)
			inrvl := tst.RltInstrInrvlI(instr, 10)
			side := tst.RltInrvlSideBid(inrvl)
			stm := tst.RltSideStmRteLst(side)
			a := stm.InrMax(unt.Unt(1)).(*rlt.StmInrMax)
			tst.RltStmInrMaxNotZero(t, a)
			mnr := tst.NewStmMnr(ap)
			a.Sub(mnr.Rx, mnr.Id)
			tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
			var actr act.Actr
			vs := actr.RunHst(a.String())
			eHst := vs[len(vs)-1].(*hst.StmInrMax)
			if eHst.Tmes != nil {
				mnr.StartFor(instr.Instr(), eHst.Tmes.Cnt())
				tst.TmesEql(t, eHst.Tmes, mnr.Tmes, "Tmes")
				tst.FltsEql(t, eHst.Vals, mnr.Vals, "Vals")
			}
			a.Unsub(mnr.Id)
			tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
			ap.Cls()
		})
	}
}
func TestRltStmInrSlp(t *testing.T) {
	for _, instr := range tst.RltPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := rlt.Oan()
			instr := instr(prv)
			inrvl := tst.RltInstrInrvlI(instr, 10)
			side := tst.RltInrvlSideBid(inrvl)
			stm := tst.RltSideStmRteLst(side)
			a := stm.InrSlp(unt.Unt(1)).(*rlt.StmInr1Slp)
			tst.RltStmInr1SlpNotZero(t, a)
			mnr := tst.NewStmMnr(ap)
			a.Sub(mnr.Rx, mnr.Id)
			tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
			var actr act.Actr
			vs := actr.RunHst(a.String())
			eHst := vs[len(vs)-1].(*hst.StmInr1Slp)
			if eHst.Tmes != nil {
				mnr.StartFor(instr.Instr(), eHst.Tmes.Cnt())
				tst.TmesEql(t, eHst.Tmes, mnr.Tmes, "Tmes")
				tst.FltsEql(t, eHst.Vals, mnr.Vals, "Vals")
			}
			a.Unsub(mnr.Id)
			tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
			ap.Cls()
		})
	}
}
func TestRltStmOtrAdd(t *testing.T) {
	for _, instr := range tst.RltPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := rlt.Oan()
			instr := instr(prv)
			inrvl := tst.RltInstrInrvlI(instr, 10)
			side := tst.RltInrvlSideBid(inrvl)
			stm := tst.RltSideStmRteLst(side)
			a := stm.OtrAdd(unt.Unt(1), tst.RltStmStmInrAdd(stm, 2)).(*rlt.StmOtrAdd)
			tst.RltStmOtrAddNotZero(t, a)
			mnr := tst.NewStmMnr(ap)
			a.Sub(mnr.Rx, mnr.Id)
			tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
			var actr act.Actr
			vs := actr.RunHst(a.String())
			eHst := vs[len(vs)-1].(*hst.StmOtrAdd)
			if eHst.Tmes != nil {
				mnr.StartFor(instr.Instr(), eHst.Tmes.Cnt())
				tst.TmesEql(t, eHst.Tmes, mnr.Tmes, "Tmes")
				tst.FltsEql(t, eHst.Vals, mnr.Vals, "Vals")
			}
			a.Unsub(mnr.Id)
			tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
			ap.Cls()
		})
	}
}
func TestRltStmOtrSub(t *testing.T) {
	for _, instr := range tst.RltPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := rlt.Oan()
			instr := instr(prv)
			inrvl := tst.RltInstrInrvlI(instr, 10)
			side := tst.RltInrvlSideBid(inrvl)
			stm := tst.RltSideStmRteLst(side)
			a := stm.OtrSub(unt.Unt(1), tst.RltStmStmInrAdd(stm, 2)).(*rlt.StmOtrSub)
			tst.RltStmOtrSubNotZero(t, a)
			mnr := tst.NewStmMnr(ap)
			a.Sub(mnr.Rx, mnr.Id)
			tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
			var actr act.Actr
			vs := actr.RunHst(a.String())
			eHst := vs[len(vs)-1].(*hst.StmOtrSub)
			if eHst.Tmes != nil {
				mnr.StartFor(instr.Instr(), eHst.Tmes.Cnt())
				tst.TmesEql(t, eHst.Tmes, mnr.Tmes, "Tmes")
				tst.FltsEql(t, eHst.Vals, mnr.Vals, "Vals")
			}
			a.Unsub(mnr.Id)
			tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
			ap.Cls()
		})
	}
}
func TestRltStmOtrMul(t *testing.T) {
	for _, instr := range tst.RltPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := rlt.Oan()
			instr := instr(prv)
			inrvl := tst.RltInstrInrvlI(instr, 10)
			side := tst.RltInrvlSideBid(inrvl)
			stm := tst.RltSideStmRteLst(side)
			a := stm.OtrMul(unt.Unt(1), tst.RltStmStmInrAdd(stm, 2)).(*rlt.StmOtrMul)
			tst.RltStmOtrMulNotZero(t, a)
			mnr := tst.NewStmMnr(ap)
			a.Sub(mnr.Rx, mnr.Id)
			tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
			var actr act.Actr
			vs := actr.RunHst(a.String())
			eHst := vs[len(vs)-1].(*hst.StmOtrMul)
			if eHst.Tmes != nil {
				mnr.StartFor(instr.Instr(), eHst.Tmes.Cnt())
				tst.TmesEql(t, eHst.Tmes, mnr.Tmes, "Tmes")
				tst.FltsEql(t, eHst.Vals, mnr.Vals, "Vals")
			}
			a.Unsub(mnr.Id)
			tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
			ap.Cls()
		})
	}
}
func TestRltStmOtrDiv(t *testing.T) {
	for _, instr := range tst.RltPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := rlt.Oan()
			instr := instr(prv)
			inrvl := tst.RltInstrInrvlI(instr, 10)
			side := tst.RltInrvlSideBid(inrvl)
			stm := tst.RltSideStmRteLst(side)
			a := stm.OtrDiv(unt.Unt(1), tst.RltStmStmInrAdd(stm, 2)).(*rlt.StmOtrDiv)
			tst.RltStmOtrDivNotZero(t, a)
			mnr := tst.NewStmMnr(ap)
			a.Sub(mnr.Rx, mnr.Id)
			tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
			var actr act.Actr
			vs := actr.RunHst(a.String())
			eHst := vs[len(vs)-1].(*hst.StmOtrDiv)
			if eHst.Tmes != nil {
				mnr.StartFor(instr.Instr(), eHst.Tmes.Cnt())
				tst.TmesEql(t, eHst.Tmes, mnr.Tmes, "Tmes")
				tst.FltsEql(t, eHst.Vals, mnr.Vals, "Vals")
			}
			a.Unsub(mnr.Id)
			tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
			ap.Cls()
		})
	}
}
func TestRltStmOtrRem(t *testing.T) {
	for _, instr := range tst.RltPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := rlt.Oan()
			instr := instr(prv)
			inrvl := tst.RltInstrInrvlI(instr, 10)
			side := tst.RltInrvlSideBid(inrvl)
			stm := tst.RltSideStmRteLst(side)
			a := stm.OtrRem(unt.Unt(1), tst.RltStmStmInrAdd(stm, 2)).(*rlt.StmOtrRem)
			tst.RltStmOtrRemNotZero(t, a)
			mnr := tst.NewStmMnr(ap)
			a.Sub(mnr.Rx, mnr.Id)
			tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
			var actr act.Actr
			vs := actr.RunHst(a.String())
			eHst := vs[len(vs)-1].(*hst.StmOtrRem)
			if eHst.Tmes != nil {
				mnr.StartFor(instr.Instr(), eHst.Tmes.Cnt())
				tst.TmesEql(t, eHst.Tmes, mnr.Tmes, "Tmes")
				tst.FltsEql(t, eHst.Vals, mnr.Vals, "Vals")
			}
			a.Unsub(mnr.Id)
			tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
			ap.Cls()
		})
	}
}
func TestRltStmOtrPow(t *testing.T) {
	for _, instr := range tst.RltPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := rlt.Oan()
			instr := instr(prv)
			inrvl := tst.RltInstrInrvlI(instr, 10)
			side := tst.RltInrvlSideBid(inrvl)
			stm := tst.RltSideStmRteLst(side)
			a := stm.OtrPow(unt.Unt(1), tst.RltStmStmInrAdd(stm, 2)).(*rlt.StmOtrPow)
			tst.RltStmOtrPowNotZero(t, a)
			mnr := tst.NewStmMnr(ap)
			a.Sub(mnr.Rx, mnr.Id)
			tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
			var actr act.Actr
			vs := actr.RunHst(a.String())
			eHst := vs[len(vs)-1].(*hst.StmOtrPow)
			if eHst.Tmes != nil {
				mnr.StartFor(instr.Instr(), eHst.Tmes.Cnt())
				tst.TmesEql(t, eHst.Tmes, mnr.Tmes, "Tmes")
				tst.FltsEql(t, eHst.Vals, mnr.Vals, "Vals")
			}
			a.Unsub(mnr.Id)
			tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
			ap.Cls()
		})
	}
}
func TestRltStmOtrMin(t *testing.T) {
	for _, instr := range tst.RltPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := rlt.Oan()
			instr := instr(prv)
			inrvl := tst.RltInstrInrvlI(instr, 10)
			side := tst.RltInrvlSideBid(inrvl)
			stm := tst.RltSideStmRteLst(side)
			a := stm.OtrMin(unt.Unt(1), tst.RltStmStmInrAdd(stm, 2)).(*rlt.StmOtrMin)
			tst.RltStmOtrMinNotZero(t, a)
			mnr := tst.NewStmMnr(ap)
			a.Sub(mnr.Rx, mnr.Id)
			tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
			var actr act.Actr
			vs := actr.RunHst(a.String())
			eHst := vs[len(vs)-1].(*hst.StmOtrMin)
			if eHst.Tmes != nil {
				mnr.StartFor(instr.Instr(), eHst.Tmes.Cnt())
				tst.TmesEql(t, eHst.Tmes, mnr.Tmes, "Tmes")
				tst.FltsEql(t, eHst.Vals, mnr.Vals, "Vals")
			}
			a.Unsub(mnr.Id)
			tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
			ap.Cls()
		})
	}
}
func TestRltStmOtrMax(t *testing.T) {
	for _, instr := range tst.RltPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := rlt.Oan()
			instr := instr(prv)
			inrvl := tst.RltInstrInrvlI(instr, 10)
			side := tst.RltInrvlSideBid(inrvl)
			stm := tst.RltSideStmRteLst(side)
			a := stm.OtrMax(unt.Unt(1), tst.RltStmStmInrAdd(stm, 2)).(*rlt.StmOtrMax)
			tst.RltStmOtrMaxNotZero(t, a)
			mnr := tst.NewStmMnr(ap)
			a.Sub(mnr.Rx, mnr.Id)
			tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
			var actr act.Actr
			vs := actr.RunHst(a.String())
			eHst := vs[len(vs)-1].(*hst.StmOtrMax)
			if eHst.Tmes != nil {
				mnr.StartFor(instr.Instr(), eHst.Tmes.Cnt())
				tst.TmesEql(t, eHst.Tmes, mnr.Tmes, "Tmes")
				tst.FltsEql(t, eHst.Vals, mnr.Vals, "Vals")
			}
			a.Unsub(mnr.Id)
			tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
			ap.Cls()
		})
	}
}
