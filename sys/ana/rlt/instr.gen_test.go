package rlt_test

import (
	"sys/ana/rlt"
	"sys/app"
	"sys/tst"
	"testing"
)

func TestRltPrvEurUsd(t *testing.T) {
	t.Run("", func(t *testing.T) {
		ap := app.New(tst.Cfg)
		ap.Oan.CloneInstrs(tst.Instrs)
		prv := rlt.Oan()
		a := prv.EurUsd().(*rlt.InstrEurUsd)
		tst.RltInstrEurUsdNotZero(t, a)
		mnr := tst.NewInstrMnr(ap)
		a.Sub(mnr.Rx, mnr.Id)
		tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
		mnr.StartFor(a.Ana, a.Ana.HstStm.Cnt())
		tst.AnaStmEql(t, a.Ana.HstStm, mnr.Stm(), "mnr.Stm")
		tst.AnaStmEql(t, a.Ana.HstStm, a.Ana.RltStm, "RltStm")
		a.Unsub(mnr.Id)
		tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
		ap.Cls()
	})
}
func TestRltPrvAudUsd(t *testing.T) {
	t.Run("", func(t *testing.T) {
		ap := app.New(tst.Cfg)
		ap.Oan.CloneInstrs(tst.Instrs)
		prv := rlt.Oan()
		a := prv.AudUsd().(*rlt.InstrAudUsd)
		tst.RltInstrAudUsdNotZero(t, a)
		mnr := tst.NewInstrMnr(ap)
		a.Sub(mnr.Rx, mnr.Id)
		tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
		mnr.StartFor(a.Ana, a.Ana.HstStm.Cnt())
		tst.AnaStmEql(t, a.Ana.HstStm, mnr.Stm(), "mnr.Stm")
		tst.AnaStmEql(t, a.Ana.HstStm, a.Ana.RltStm, "RltStm")
		a.Unsub(mnr.Id)
		tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
		ap.Cls()
	})
}
func TestRltPrvNzdUsd(t *testing.T) {
	t.Run("", func(t *testing.T) {
		ap := app.New(tst.Cfg)
		ap.Oan.CloneInstrs(tst.Instrs)
		prv := rlt.Oan()
		a := prv.NzdUsd().(*rlt.InstrNzdUsd)
		tst.RltInstrNzdUsdNotZero(t, a)
		mnr := tst.NewInstrMnr(ap)
		a.Sub(mnr.Rx, mnr.Id)
		tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
		mnr.StartFor(a.Ana, a.Ana.HstStm.Cnt())
		tst.AnaStmEql(t, a.Ana.HstStm, mnr.Stm(), "mnr.Stm")
		tst.AnaStmEql(t, a.Ana.HstStm, a.Ana.RltStm, "RltStm")
		a.Unsub(mnr.Id)
		tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
		ap.Cls()
	})
}
func TestRltPrvGbpUsd(t *testing.T) {
	t.Run("", func(t *testing.T) {
		ap := app.New(tst.Cfg)
		ap.Oan.CloneInstrs(tst.Instrs)
		prv := rlt.Oan()
		a := prv.GbpUsd().(*rlt.InstrGbpUsd)
		tst.RltInstrGbpUsdNotZero(t, a)
		mnr := tst.NewInstrMnr(ap)
		a.Sub(mnr.Rx, mnr.Id)
		tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
		mnr.StartFor(a.Ana, a.Ana.HstStm.Cnt())
		tst.AnaStmEql(t, a.Ana.HstStm, mnr.Stm(), "mnr.Stm")
		tst.AnaStmEql(t, a.Ana.HstStm, a.Ana.RltStm, "RltStm")
		a.Unsub(mnr.Id)
		tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
		ap.Cls()
	})
}
