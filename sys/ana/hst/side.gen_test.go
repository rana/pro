package hst_test

import (
	"sys/ana/hst"
	"sys/app"
	"sys/tst"
	"testing"
)

func TestHstInrvlBid(t *testing.T) {
	for _, instr := range tst.HstPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := hst.Oan()
			instr := instr(prv)
			inrvl := tst.HstInstrInrvlI(instr, 10)
			a := inrvl.Bid().(*hst.SideBid)
			tst.HstSideBidNotZero(t, a)
			ap.Cls()
		})
	}
}
func TestHstInrvlAsk(t *testing.T) {
	for _, instr := range tst.HstPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := hst.Oan()
			instr := instr(prv)
			inrvl := tst.HstInstrInrvlI(instr, 10)
			a := inrvl.Ask().(*hst.SideAsk)
			tst.HstSideAskNotZero(t, a)
			ap.Cls()
		})
	}
}
