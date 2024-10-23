package rlt_test

import (
	"testing"
)

func TestRltCndStgy(t *testing.T) {
	// for _, instr := range tst.RltPrvInstrs {
	// 	t.Run("", func(t *testing.T) {
	// 		ap := app.New(tst.Cfg)
	// 		ap.Oan.CloneInstrs(tst.Instrs)
	// 		prv := rlt.Oan()
	// 		instr := instr(prv)
	// 		inrvl := tst.RltInstrInrvlI(instr, 10)
	// 		side := tst.RltInrvlSideBid(inrvl)
	// 		stm := tst.RltSideStmRteLst(side)
	// 		cnd := tst.RltStmCndInrGtr(stm, 1)
	// 		a := cnd.Seq(tme.Tme(1), tst.RltStmCndInrLss(stm, 2)).(*rlt.CndCnd2Seq)
	// 		tst.RltCndCnd2SeqNotZero(t, a)
	// 		mnr := tst.NewCndMnr(ap)
	// 		a.Sub(mnr.Rx, mnr.Id)
	// 		tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
	// 		var actr act.Actr
	// 		vs := actr.RunHst(a.String())
	// 		eHst := vs[len(vs)-1].(*hst.CndCnd2Seq)
	// 		if eHst.Tmes != nil {
	// 			mnr.StartFor(instr.Instr(), eHst.Tmes.Cnt())
	// 			tst.TmesEql(t, eHst.Tmes, mnr.Tmes, "Tmes")
	// 		}
	// 		a.Unsub(mnr.Id)
	// 		tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
	// 		ap.Cls()
	// 	})
	// }
}
