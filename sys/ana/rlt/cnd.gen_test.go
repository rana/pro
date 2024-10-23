package rlt_test

import (
	"sys/ana/hst"
	"sys/ana/rlt"
	"sys/app"
	"sys/bsc/flt"
	"sys/bsc/tme"
	"sys/bsc/unt"
	"sys/lng/pro/act"
	"sys/tst"
	"testing"
)

func TestRltStmSclEql(t *testing.T) {
	for _, instr := range tst.RltPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := rlt.Oan()
			instr := instr(prv)
			inrvl := tst.RltInstrInrvlI(instr, 10)
			side := tst.RltInrvlSideBid(inrvl)
			stm := tst.RltSideStmRteLst(side)
			a := stm.SclEql(flt.Flt(1.1)).(*rlt.CndSclEql)
			tst.RltCndSclEqlNotZero(t, a)
			mnr := tst.NewCndMnr(ap)
			a.Sub(mnr.Rx, mnr.Id)
			tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
			var actr act.Actr
			vs := actr.RunHst(a.String())
			eHst := vs[len(vs)-1].(*hst.CndSclEql)
			if eHst.Tmes != nil {
				mnr.StartFor(instr.Instr(), eHst.Tmes.Cnt())
				tst.TmesEql(t, eHst.Tmes, mnr.Tmes, "Tmes")
			}
			a.Unsub(mnr.Id)
			tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
			ap.Cls()
		})
	}
}
func TestRltStmSclNeq(t *testing.T) {
	for _, instr := range tst.RltPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := rlt.Oan()
			instr := instr(prv)
			inrvl := tst.RltInstrInrvlI(instr, 10)
			side := tst.RltInrvlSideBid(inrvl)
			stm := tst.RltSideStmRteLst(side)
			a := stm.SclNeq(flt.Flt(1.1)).(*rlt.CndSclNeq)
			tst.RltCndSclNeqNotZero(t, a)
			mnr := tst.NewCndMnr(ap)
			a.Sub(mnr.Rx, mnr.Id)
			tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
			var actr act.Actr
			vs := actr.RunHst(a.String())
			eHst := vs[len(vs)-1].(*hst.CndSclNeq)
			if eHst.Tmes != nil {
				mnr.StartFor(instr.Instr(), eHst.Tmes.Cnt())
				tst.TmesEql(t, eHst.Tmes, mnr.Tmes, "Tmes")
			}
			a.Unsub(mnr.Id)
			tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
			ap.Cls()
		})
	}
}
func TestRltStmSclLss(t *testing.T) {
	for _, instr := range tst.RltPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := rlt.Oan()
			instr := instr(prv)
			inrvl := tst.RltInstrInrvlI(instr, 10)
			side := tst.RltInrvlSideBid(inrvl)
			stm := tst.RltSideStmRteLst(side)
			a := stm.SclLss(flt.Flt(1.1)).(*rlt.CndSclLss)
			tst.RltCndSclLssNotZero(t, a)
			mnr := tst.NewCndMnr(ap)
			a.Sub(mnr.Rx, mnr.Id)
			tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
			var actr act.Actr
			vs := actr.RunHst(a.String())
			eHst := vs[len(vs)-1].(*hst.CndSclLss)
			if eHst.Tmes != nil {
				mnr.StartFor(instr.Instr(), eHst.Tmes.Cnt())
				tst.TmesEql(t, eHst.Tmes, mnr.Tmes, "Tmes")
			}
			a.Unsub(mnr.Id)
			tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
			ap.Cls()
		})
	}
}
func TestRltStmSclGtr(t *testing.T) {
	for _, instr := range tst.RltPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := rlt.Oan()
			instr := instr(prv)
			inrvl := tst.RltInstrInrvlI(instr, 10)
			side := tst.RltInrvlSideBid(inrvl)
			stm := tst.RltSideStmRteLst(side)
			a := stm.SclGtr(flt.Flt(1.1)).(*rlt.CndSclGtr)
			tst.RltCndSclGtrNotZero(t, a)
			mnr := tst.NewCndMnr(ap)
			a.Sub(mnr.Rx, mnr.Id)
			tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
			var actr act.Actr
			vs := actr.RunHst(a.String())
			eHst := vs[len(vs)-1].(*hst.CndSclGtr)
			if eHst.Tmes != nil {
				mnr.StartFor(instr.Instr(), eHst.Tmes.Cnt())
				tst.TmesEql(t, eHst.Tmes, mnr.Tmes, "Tmes")
			}
			a.Unsub(mnr.Id)
			tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
			ap.Cls()
		})
	}
}
func TestRltStmSclLeq(t *testing.T) {
	for _, instr := range tst.RltPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := rlt.Oan()
			instr := instr(prv)
			inrvl := tst.RltInstrInrvlI(instr, 10)
			side := tst.RltInrvlSideBid(inrvl)
			stm := tst.RltSideStmRteLst(side)
			a := stm.SclLeq(flt.Flt(1.1)).(*rlt.CndSclLeq)
			tst.RltCndSclLeqNotZero(t, a)
			mnr := tst.NewCndMnr(ap)
			a.Sub(mnr.Rx, mnr.Id)
			tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
			var actr act.Actr
			vs := actr.RunHst(a.String())
			eHst := vs[len(vs)-1].(*hst.CndSclLeq)
			if eHst.Tmes != nil {
				mnr.StartFor(instr.Instr(), eHst.Tmes.Cnt())
				tst.TmesEql(t, eHst.Tmes, mnr.Tmes, "Tmes")
			}
			a.Unsub(mnr.Id)
			tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
			ap.Cls()
		})
	}
}
func TestRltStmSclGeq(t *testing.T) {
	for _, instr := range tst.RltPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := rlt.Oan()
			instr := instr(prv)
			inrvl := tst.RltInstrInrvlI(instr, 10)
			side := tst.RltInrvlSideBid(inrvl)
			stm := tst.RltSideStmRteLst(side)
			a := stm.SclGeq(flt.Flt(1.1)).(*rlt.CndSclGeq)
			tst.RltCndSclGeqNotZero(t, a)
			mnr := tst.NewCndMnr(ap)
			a.Sub(mnr.Rx, mnr.Id)
			tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
			var actr act.Actr
			vs := actr.RunHst(a.String())
			eHst := vs[len(vs)-1].(*hst.CndSclGeq)
			if eHst.Tmes != nil {
				mnr.StartFor(instr.Instr(), eHst.Tmes.Cnt())
				tst.TmesEql(t, eHst.Tmes, mnr.Tmes, "Tmes")
			}
			a.Unsub(mnr.Id)
			tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
			ap.Cls()
		})
	}
}
func TestRltStmInrEql(t *testing.T) {
	for _, instr := range tst.RltPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := rlt.Oan()
			instr := instr(prv)
			inrvl := tst.RltInstrInrvlI(instr, 10)
			side := tst.RltInrvlSideBid(inrvl)
			stm := tst.RltSideStmRteLst(side)
			a := stm.InrEql(unt.Unt(1)).(*rlt.CndInrEql)
			tst.RltCndInrEqlNotZero(t, a)
			mnr := tst.NewCndMnr(ap)
			a.Sub(mnr.Rx, mnr.Id)
			tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
			var actr act.Actr
			vs := actr.RunHst(a.String())
			eHst := vs[len(vs)-1].(*hst.CndInrEql)
			if eHst.Tmes != nil {
				mnr.StartFor(instr.Instr(), eHst.Tmes.Cnt())
				tst.TmesEql(t, eHst.Tmes, mnr.Tmes, "Tmes")
			}
			a.Unsub(mnr.Id)
			tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
			ap.Cls()
		})
	}
}
func TestRltStmInrNeq(t *testing.T) {
	for _, instr := range tst.RltPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := rlt.Oan()
			instr := instr(prv)
			inrvl := tst.RltInstrInrvlI(instr, 10)
			side := tst.RltInrvlSideBid(inrvl)
			stm := tst.RltSideStmRteLst(side)
			a := stm.InrNeq(unt.Unt(1)).(*rlt.CndInrNeq)
			tst.RltCndInrNeqNotZero(t, a)
			mnr := tst.NewCndMnr(ap)
			a.Sub(mnr.Rx, mnr.Id)
			tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
			var actr act.Actr
			vs := actr.RunHst(a.String())
			eHst := vs[len(vs)-1].(*hst.CndInrNeq)
			if eHst.Tmes != nil {
				mnr.StartFor(instr.Instr(), eHst.Tmes.Cnt())
				tst.TmesEql(t, eHst.Tmes, mnr.Tmes, "Tmes")
			}
			a.Unsub(mnr.Id)
			tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
			ap.Cls()
		})
	}
}
func TestRltStmInrLss(t *testing.T) {
	for _, instr := range tst.RltPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := rlt.Oan()
			instr := instr(prv)
			inrvl := tst.RltInstrInrvlI(instr, 10)
			side := tst.RltInrvlSideBid(inrvl)
			stm := tst.RltSideStmRteLst(side)
			a := stm.InrLss(unt.Unt(1)).(*rlt.CndInrLss)
			tst.RltCndInrLssNotZero(t, a)
			mnr := tst.NewCndMnr(ap)
			a.Sub(mnr.Rx, mnr.Id)
			tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
			var actr act.Actr
			vs := actr.RunHst(a.String())
			eHst := vs[len(vs)-1].(*hst.CndInrLss)
			if eHst.Tmes != nil {
				mnr.StartFor(instr.Instr(), eHst.Tmes.Cnt())
				tst.TmesEql(t, eHst.Tmes, mnr.Tmes, "Tmes")
			}
			a.Unsub(mnr.Id)
			tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
			ap.Cls()
		})
	}
}
func TestRltStmInrGtr(t *testing.T) {
	for _, instr := range tst.RltPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := rlt.Oan()
			instr := instr(prv)
			inrvl := tst.RltInstrInrvlI(instr, 10)
			side := tst.RltInrvlSideBid(inrvl)
			stm := tst.RltSideStmRteLst(side)
			a := stm.InrGtr(unt.Unt(1)).(*rlt.CndInrGtr)
			tst.RltCndInrGtrNotZero(t, a)
			mnr := tst.NewCndMnr(ap)
			a.Sub(mnr.Rx, mnr.Id)
			tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
			var actr act.Actr
			vs := actr.RunHst(a.String())
			eHst := vs[len(vs)-1].(*hst.CndInrGtr)
			if eHst.Tmes != nil {
				mnr.StartFor(instr.Instr(), eHst.Tmes.Cnt())
				tst.TmesEql(t, eHst.Tmes, mnr.Tmes, "Tmes")
			}
			a.Unsub(mnr.Id)
			tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
			ap.Cls()
		})
	}
}
func TestRltStmInrLeq(t *testing.T) {
	for _, instr := range tst.RltPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := rlt.Oan()
			instr := instr(prv)
			inrvl := tst.RltInstrInrvlI(instr, 10)
			side := tst.RltInrvlSideBid(inrvl)
			stm := tst.RltSideStmRteLst(side)
			a := stm.InrLeq(unt.Unt(1)).(*rlt.CndInrLeq)
			tst.RltCndInrLeqNotZero(t, a)
			mnr := tst.NewCndMnr(ap)
			a.Sub(mnr.Rx, mnr.Id)
			tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
			var actr act.Actr
			vs := actr.RunHst(a.String())
			eHst := vs[len(vs)-1].(*hst.CndInrLeq)
			if eHst.Tmes != nil {
				mnr.StartFor(instr.Instr(), eHst.Tmes.Cnt())
				tst.TmesEql(t, eHst.Tmes, mnr.Tmes, "Tmes")
			}
			a.Unsub(mnr.Id)
			tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
			ap.Cls()
		})
	}
}
func TestRltStmInrGeq(t *testing.T) {
	for _, instr := range tst.RltPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := rlt.Oan()
			instr := instr(prv)
			inrvl := tst.RltInstrInrvlI(instr, 10)
			side := tst.RltInrvlSideBid(inrvl)
			stm := tst.RltSideStmRteLst(side)
			a := stm.InrGeq(unt.Unt(1)).(*rlt.CndInrGeq)
			tst.RltCndInrGeqNotZero(t, a)
			mnr := tst.NewCndMnr(ap)
			a.Sub(mnr.Rx, mnr.Id)
			tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
			var actr act.Actr
			vs := actr.RunHst(a.String())
			eHst := vs[len(vs)-1].(*hst.CndInrGeq)
			if eHst.Tmes != nil {
				mnr.StartFor(instr.Instr(), eHst.Tmes.Cnt())
				tst.TmesEql(t, eHst.Tmes, mnr.Tmes, "Tmes")
			}
			a.Unsub(mnr.Id)
			tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
			ap.Cls()
		})
	}
}
func TestRltStmOtrEql(t *testing.T) {
	for _, instr := range tst.RltPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := rlt.Oan()
			instr := instr(prv)
			inrvl := tst.RltInstrInrvlI(instr, 10)
			side := tst.RltInrvlSideBid(inrvl)
			stm := tst.RltSideStmRteLst(side)
			a := stm.OtrEql(unt.Unt(1), tst.RltStmStmInrAdd(stm, 2)).(*rlt.CndOtrEql)
			tst.RltCndOtrEqlNotZero(t, a)
			mnr := tst.NewCndMnr(ap)
			a.Sub(mnr.Rx, mnr.Id)
			tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
			var actr act.Actr
			vs := actr.RunHst(a.String())
			eHst := vs[len(vs)-1].(*hst.CndOtrEql)
			if eHst.Tmes != nil {
				mnr.StartFor(instr.Instr(), eHst.Tmes.Cnt())
				tst.TmesEql(t, eHst.Tmes, mnr.Tmes, "Tmes")
			}
			a.Unsub(mnr.Id)
			tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
			ap.Cls()
		})
	}
}
func TestRltStmOtrNeq(t *testing.T) {
	for _, instr := range tst.RltPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := rlt.Oan()
			instr := instr(prv)
			inrvl := tst.RltInstrInrvlI(instr, 10)
			side := tst.RltInrvlSideBid(inrvl)
			stm := tst.RltSideStmRteLst(side)
			a := stm.OtrNeq(unt.Unt(1), tst.RltStmStmInrAdd(stm, 2)).(*rlt.CndOtrNeq)
			tst.RltCndOtrNeqNotZero(t, a)
			mnr := tst.NewCndMnr(ap)
			a.Sub(mnr.Rx, mnr.Id)
			tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
			var actr act.Actr
			vs := actr.RunHst(a.String())
			eHst := vs[len(vs)-1].(*hst.CndOtrNeq)
			if eHst.Tmes != nil {
				mnr.StartFor(instr.Instr(), eHst.Tmes.Cnt())
				tst.TmesEql(t, eHst.Tmes, mnr.Tmes, "Tmes")
			}
			a.Unsub(mnr.Id)
			tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
			ap.Cls()
		})
	}
}
func TestRltStmOtrLss(t *testing.T) {
	for _, instr := range tst.RltPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := rlt.Oan()
			instr := instr(prv)
			inrvl := tst.RltInstrInrvlI(instr, 10)
			side := tst.RltInrvlSideBid(inrvl)
			stm := tst.RltSideStmRteLst(side)
			a := stm.OtrLss(unt.Unt(1), tst.RltStmStmInrAdd(stm, 2)).(*rlt.CndOtrLss)
			tst.RltCndOtrLssNotZero(t, a)
			mnr := tst.NewCndMnr(ap)
			a.Sub(mnr.Rx, mnr.Id)
			tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
			var actr act.Actr
			vs := actr.RunHst(a.String())
			eHst := vs[len(vs)-1].(*hst.CndOtrLss)
			if eHst.Tmes != nil {
				mnr.StartFor(instr.Instr(), eHst.Tmes.Cnt())
				tst.TmesEql(t, eHst.Tmes, mnr.Tmes, "Tmes")
			}
			a.Unsub(mnr.Id)
			tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
			ap.Cls()
		})
	}
}
func TestRltStmOtrGtr(t *testing.T) {
	for _, instr := range tst.RltPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := rlt.Oan()
			instr := instr(prv)
			inrvl := tst.RltInstrInrvlI(instr, 10)
			side := tst.RltInrvlSideBid(inrvl)
			stm := tst.RltSideStmRteLst(side)
			a := stm.OtrGtr(unt.Unt(1), tst.RltStmStmInrAdd(stm, 2)).(*rlt.CndOtrGtr)
			tst.RltCndOtrGtrNotZero(t, a)
			mnr := tst.NewCndMnr(ap)
			a.Sub(mnr.Rx, mnr.Id)
			tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
			var actr act.Actr
			vs := actr.RunHst(a.String())
			eHst := vs[len(vs)-1].(*hst.CndOtrGtr)
			if eHst.Tmes != nil {
				mnr.StartFor(instr.Instr(), eHst.Tmes.Cnt())
				tst.TmesEql(t, eHst.Tmes, mnr.Tmes, "Tmes")
			}
			a.Unsub(mnr.Id)
			tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
			ap.Cls()
		})
	}
}
func TestRltStmOtrLeq(t *testing.T) {
	for _, instr := range tst.RltPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := rlt.Oan()
			instr := instr(prv)
			inrvl := tst.RltInstrInrvlI(instr, 10)
			side := tst.RltInrvlSideBid(inrvl)
			stm := tst.RltSideStmRteLst(side)
			a := stm.OtrLeq(unt.Unt(1), tst.RltStmStmInrAdd(stm, 2)).(*rlt.CndOtrLeq)
			tst.RltCndOtrLeqNotZero(t, a)
			mnr := tst.NewCndMnr(ap)
			a.Sub(mnr.Rx, mnr.Id)
			tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
			var actr act.Actr
			vs := actr.RunHst(a.String())
			eHst := vs[len(vs)-1].(*hst.CndOtrLeq)
			if eHst.Tmes != nil {
				mnr.StartFor(instr.Instr(), eHst.Tmes.Cnt())
				tst.TmesEql(t, eHst.Tmes, mnr.Tmes, "Tmes")
			}
			a.Unsub(mnr.Id)
			tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
			ap.Cls()
		})
	}
}
func TestRltStmOtrGeq(t *testing.T) {
	for _, instr := range tst.RltPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := rlt.Oan()
			instr := instr(prv)
			inrvl := tst.RltInstrInrvlI(instr, 10)
			side := tst.RltInrvlSideBid(inrvl)
			stm := tst.RltSideStmRteLst(side)
			a := stm.OtrGeq(unt.Unt(1), tst.RltStmStmInrAdd(stm, 2)).(*rlt.CndOtrGeq)
			tst.RltCndOtrGeqNotZero(t, a)
			mnr := tst.NewCndMnr(ap)
			a.Sub(mnr.Rx, mnr.Id)
			tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
			var actr act.Actr
			vs := actr.RunHst(a.String())
			eHst := vs[len(vs)-1].(*hst.CndOtrGeq)
			if eHst.Tmes != nil {
				mnr.StartFor(instr.Instr(), eHst.Tmes.Cnt())
				tst.TmesEql(t, eHst.Tmes, mnr.Tmes, "Tmes")
			}
			a.Unsub(mnr.Id)
			tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
			ap.Cls()
		})
	}
}
func TestRltCndAnd(t *testing.T) {
	for _, instr := range tst.RltPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := rlt.Oan()
			instr := instr(prv)
			inrvl := tst.RltInstrInrvlI(instr, 10)
			side := tst.RltInrvlSideBid(inrvl)
			stm := tst.RltSideStmRteLst(side)
			cnd := tst.RltStmCndInrGtr(stm, 1)
			a := cnd.And(tst.RltStmCndInrGtr(stm, 2)).(*rlt.CndCnd1And)
			tst.RltCndCnd1AndNotZero(t, a)
			mnr := tst.NewCndMnr(ap)
			a.Sub(mnr.Rx, mnr.Id)
			tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
			var actr act.Actr
			vs := actr.RunHst(a.String())
			eHst := vs[len(vs)-1].(*hst.CndCnd1And)
			if eHst.Tmes != nil {
				mnr.StartFor(instr.Instr(), eHst.Tmes.Cnt())
				tst.TmesEql(t, eHst.Tmes, mnr.Tmes, "Tmes")
			}
			a.Unsub(mnr.Id)
			tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
			ap.Cls()
		})
	}
}
func TestRltCndSeq(t *testing.T) {
	for _, instr := range tst.RltPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := rlt.Oan()
			instr := instr(prv)
			inrvl := tst.RltInstrInrvlI(instr, 10)
			side := tst.RltInrvlSideBid(inrvl)
			stm := tst.RltSideStmRteLst(side)
			cnd := tst.RltStmCndInrGtr(stm, 1)
			a := cnd.Seq(tme.Tme(1), tst.RltStmCndInrLss(stm, 2)).(*rlt.CndCnd2Seq)
			tst.RltCndCnd2SeqNotZero(t, a)
			mnr := tst.NewCndMnr(ap)
			a.Sub(mnr.Rx, mnr.Id)
			tst.IntegerEql(t, 1, len(a.Rxs), "Sub Rxs")
			var actr act.Actr
			vs := actr.RunHst(a.String())
			eHst := vs[len(vs)-1].(*hst.CndCnd2Seq)
			if eHst.Tmes != nil {
				mnr.StartFor(instr.Instr(), eHst.Tmes.Cnt())
				tst.TmesEql(t, eHst.Tmes, mnr.Tmes, "Tmes")
			}
			a.Unsub(mnr.Id)
			tst.IntegerEql(t, 0, len(a.Rxs), "Unsub Rxs")
			ap.Cls()
		})
	}
}
