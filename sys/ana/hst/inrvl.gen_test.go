package hst_test

import (
	"sys/ana/hst"
	"sys/app"
	"sys/bsc/tme"
	"sys/tst"
	"testing"
)

func TestHstInstrI(t *testing.T) {
	for _, instr := range tst.HstPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := hst.Oan()
			instr := instr(prv)
			a := instr.I(tme.Tme(10)).(*hst.InrvlI)
			tst.HstInrvlINotZero(t, a)
			ap.Cls()
		})
	}
}
