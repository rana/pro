package rlt_test

import (
	"sys/ana/hst"
	"sys/ana/rlt"
	"sys/app"
	"sys/lng/pro/act"
	"sys/tst"
	"testing"
)

func TestRltInrvlBid(t *testing.T) {
	for _, instr := range tst.RltPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := rlt.Oan()
			instr := instr(prv)
			inrvl := tst.RltInstrInrvlI(instr, 10)
			a := inrvl.Bid().(*rlt.SideBid)
			tst.RltSideBidNotZero(t, a)
			mnr := tst.NewSideMnr(ap)
			a.Sub(mnr.Rx, mnr.Id)
			tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
			var actr act.Actr
			vs := actr.RunHst(a.String())
			eHst := vs[len(vs)-1].(*hst.SideBid)
			if eHst.ValBnds != nil {
				mnr.StartFor(instr.Instr(), eHst.ValBnds.Cnt())
				tst.IntegerEql(t, len(*eHst.ValBnds), len(mnr.Fltss), "ValBnds Cnt")
				for n, valBnd := range *eHst.ValBnds {
					tst.FltsEql(t, eHst.Vals.In(valBnd.Idx, valBnd.Lim), mnr.Fltss[n], n, "Mnr Vals")
				}
			}
			a.Unsub(mnr.Id)
			tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
			ap.Cls()
		})
	}
}
func TestRltInrvlAsk(t *testing.T) {
	for _, instr := range tst.RltPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := rlt.Oan()
			instr := instr(prv)
			inrvl := tst.RltInstrInrvlI(instr, 10)
			a := inrvl.Ask().(*rlt.SideAsk)
			tst.RltSideAskNotZero(t, a)
			mnr := tst.NewSideMnr(ap)
			a.Sub(mnr.Rx, mnr.Id)
			tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
			var actr act.Actr
			vs := actr.RunHst(a.String())
			eHst := vs[len(vs)-1].(*hst.SideAsk)
			if eHst.ValBnds != nil {
				mnr.StartFor(instr.Instr(), eHst.ValBnds.Cnt())
				tst.IntegerEql(t, len(*eHst.ValBnds), len(mnr.Fltss), "ValBnds Cnt")
				for n, valBnd := range *eHst.ValBnds {
					tst.FltsEql(t, eHst.Vals.In(valBnd.Idx, valBnd.Lim), mnr.Fltss[n], n, "Mnr Vals")
				}
			}
			a.Unsub(mnr.Id)
			tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
			ap.Cls()
		})
	}
}
