package rlt_test

import (
	"sys/ana/hst"
	"sys/ana/rlt"
	"sys/app"
	"sys/bsc/tme"
	"sys/lng/pro/act"
	"sys/tst"
	"testing"
)

func TestRltInstrI(t *testing.T) {
	for _, instr := range tst.RltPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := rlt.Oan()
			instr := instr(prv)
			a := instr.I(tme.Tme(10)).(*rlt.InrvlI)
			tst.RltInrvlINotZero(t, a)
			tst.NotNil(t, a.Pkts, "Pkts")
			mnr := tst.NewInrvlMnr(ap)
			a.Sub(mnr.Rx, mnr.Id)
			tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
			var actr act.Actr
			vs := actr.RunHst(a.String())
			eHst := vs[len(vs)-1].(*hst.InrvlI)
			if eHst.TmeBnds != nil {
				mnr.StartFor(instr.Instr(), eHst.TmeBnds.Cnt())
				tst.BndsEql(t, eHst.TmeBnds, mnr.Bnds, "TmeBnds")
			}
			a.Unsub(mnr.Id)
			tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
			ap.Cls()
		})
	}
}
