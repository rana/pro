package hst_test

import (
	"sys/ana/hst"
	"sys/app"
	"sys/tst"
	"testing"
)

func TestHstPrvEurUsd(t *testing.T) {
	t.Run("", func(t *testing.T) {
		ap := app.New(tst.Cfg)
		ap.Oan.CloneInstrs(tst.Instrs)
		prv := hst.Oan()
		a := prv.EurUsd().(*hst.InstrEurUsd)
		tst.HstInstrEurUsdNotZero(t, a)
		ap.Cls()
	})
}
func TestHstPrvAudUsd(t *testing.T) {
	t.Run("", func(t *testing.T) {
		ap := app.New(tst.Cfg)
		ap.Oan.CloneInstrs(tst.Instrs)
		prv := hst.Oan()
		a := prv.AudUsd().(*hst.InstrAudUsd)
		tst.HstInstrAudUsdNotZero(t, a)
		ap.Cls()
	})
}
func TestHstPrvNzdUsd(t *testing.T) {
	t.Run("", func(t *testing.T) {
		ap := app.New(tst.Cfg)
		ap.Oan.CloneInstrs(tst.Instrs)
		prv := hst.Oan()
		a := prv.NzdUsd().(*hst.InstrNzdUsd)
		tst.HstInstrNzdUsdNotZero(t, a)
		ap.Cls()
	})
}
func TestHstPrvGbpUsd(t *testing.T) {
	t.Run("", func(t *testing.T) {
		ap := app.New(tst.Cfg)
		ap.Oan.CloneInstrs(tst.Instrs)
		prv := hst.Oan()
		a := prv.GbpUsd().(*hst.InstrGbpUsd)
		tst.HstInstrGbpUsdNotZero(t, a)
		ap.Cls()
	})
}
