package rlt_test

import (
	"sys/ana/rlt"
	"sys/app"
	"sys/tst"
	"testing"
)

func TestRltOan(t *testing.T) {
	t.Run("", func(t *testing.T) {
		ap := app.New(tst.Cfg)
		ap.Oan.CloneInstrs(tst.Instrs)
		a := rlt.Oan().(*rlt.PrvOan)
		tst.RltPrvOanNotZero(t, a)
		ap.Cls()
	})
}
