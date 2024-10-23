package hst_test

import (
	"sys/ana/hst"
	"sys/app"
	"sys/tst"
	"testing"
)

func TestHstOan(t *testing.T) {
	t.Run("", func(t *testing.T) {
		ap := app.New(tst.Cfg)
		ap.Oan.CloneInstrs(tst.Instrs)
		a := hst.Oan().(*hst.PrvOan)
		tst.HstPrvOanNotZero(t, a)
		ap.Cls()
	})
}
