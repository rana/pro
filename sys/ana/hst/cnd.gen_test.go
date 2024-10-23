package hst_test

import (
	"sys/ana/hst"
	"sys/app"
	"sys/bsc/flt"
	"sys/bsc/tme"
	"sys/bsc/tmes"
	"sys/bsc/unt"
	"sys/tst"
	"testing"
)

func TestHstStmSclEql(t *testing.T) {
	for _, instr := range tst.HstPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := hst.Oan()
			instr := instr(prv)
			inrvl := tst.HstInstrInrvlI(instr, 10)
			side := tst.HstInrvlSideBid(inrvl)
			stm := tst.HstSideStmRteLst(side)
			a := stm.SclEql(flt.Flt(1.1)).(*hst.CndSclEql)
			tst.HstCndSclEqlNotZero(t, a)
			ap.Cls()
		})
	}
}
func TestHstStmSclNeq(t *testing.T) {
	for _, instr := range tst.HstPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := hst.Oan()
			instr := instr(prv)
			inrvl := tst.HstInstrInrvlI(instr, 10)
			side := tst.HstInrvlSideBid(inrvl)
			stm := tst.HstSideStmRteLst(side)
			a := stm.SclNeq(flt.Flt(1.1)).(*hst.CndSclNeq)
			tst.HstCndSclNeqNotZero(t, a)
			ap.Cls()
		})
	}
}
func TestHstStmSclLss(t *testing.T) {
	for _, instr := range tst.HstPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := hst.Oan()
			instr := instr(prv)
			inrvl := tst.HstInstrInrvlI(instr, 10)
			side := tst.HstInrvlSideBid(inrvl)
			stm := tst.HstSideStmRteLst(side)
			a := stm.SclLss(flt.Flt(1.1)).(*hst.CndSclLss)
			tst.HstCndSclLssNotZero(t, a)
			ap.Cls()
		})
	}
}
func TestHstStmSclGtr(t *testing.T) {
	for _, instr := range tst.HstPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := hst.Oan()
			instr := instr(prv)
			inrvl := tst.HstInstrInrvlI(instr, 10)
			side := tst.HstInrvlSideBid(inrvl)
			stm := tst.HstSideStmRteLst(side)
			a := stm.SclGtr(flt.Flt(1.1)).(*hst.CndSclGtr)
			tst.HstCndSclGtrNotZero(t, a)
			ap.Cls()
		})
	}
}
func TestHstStmSclLeq(t *testing.T) {
	for _, instr := range tst.HstPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := hst.Oan()
			instr := instr(prv)
			inrvl := tst.HstInstrInrvlI(instr, 10)
			side := tst.HstInrvlSideBid(inrvl)
			stm := tst.HstSideStmRteLst(side)
			a := stm.SclLeq(flt.Flt(1.1)).(*hst.CndSclLeq)
			tst.HstCndSclLeqNotZero(t, a)
			ap.Cls()
		})
	}
}
func TestHstStmSclGeq(t *testing.T) {
	for _, instr := range tst.HstPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := hst.Oan()
			instr := instr(prv)
			inrvl := tst.HstInstrInrvlI(instr, 10)
			side := tst.HstInrvlSideBid(inrvl)
			stm := tst.HstSideStmRteLst(side)
			a := stm.SclGeq(flt.Flt(1.1)).(*hst.CndSclGeq)
			tst.HstCndSclGeqNotZero(t, a)
			ap.Cls()
		})
	}
}
func TestHstStmInrEql(t *testing.T) {
	for _, instr := range tst.HstPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := hst.Oan()
			instr := instr(prv)
			inrvl := tst.HstInstrInrvlI(instr, 10)
			side := tst.HstInrvlSideBid(inrvl)
			stm := tst.HstSideStmRteLst(side)
			a := stm.InrEql(unt.Unt(1)).(*hst.CndInrEql)
			tst.HstCndInrEqlNotZero(t, a)
			x, e := stm.Bse(), &hst.CndInrEql{}
			e.Off = unt.Unt(1)
			e.Tmes = tmes.New()
			for n := unt.Zero; n < x.Tmes.Cnt()-e.Off; n++ {
				if x.Vals.At(n + e.Off).Eql(x.Vals.At(n)) {
					*e.Tmes = append(*e.Tmes, x.Tmes.At(n+e.Off))
				}
			}
			tst.TmesEql(t, e.Tmes, a.Tmes, "CndInrEql.Tmes")
			ap.Cls()
		})
	}
}
func TestHstStmInrNeq(t *testing.T) {
	for _, instr := range tst.HstPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := hst.Oan()
			instr := instr(prv)
			inrvl := tst.HstInstrInrvlI(instr, 10)
			side := tst.HstInrvlSideBid(inrvl)
			stm := tst.HstSideStmRteLst(side)
			a := stm.InrNeq(unt.Unt(1)).(*hst.CndInrNeq)
			tst.HstCndInrNeqNotZero(t, a)
			x, e := stm.Bse(), &hst.CndInrNeq{}
			e.Off = unt.Unt(1)
			e.Tmes = tmes.New()
			for n := unt.Zero; n < x.Tmes.Cnt()-e.Off; n++ {
				if x.Vals.At(n + e.Off).Neq(x.Vals.At(n)) {
					*e.Tmes = append(*e.Tmes, x.Tmes.At(n+e.Off))
				}
			}
			tst.TmesEql(t, e.Tmes, a.Tmes, "CndInrNeq.Tmes")
			ap.Cls()
		})
	}
}
func TestHstStmInrLss(t *testing.T) {
	for _, instr := range tst.HstPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := hst.Oan()
			instr := instr(prv)
			inrvl := tst.HstInstrInrvlI(instr, 10)
			side := tst.HstInrvlSideBid(inrvl)
			stm := tst.HstSideStmRteLst(side)
			a := stm.InrLss(unt.Unt(1)).(*hst.CndInrLss)
			tst.HstCndInrLssNotZero(t, a)
			x, e := stm.Bse(), &hst.CndInrLss{}
			e.Off = unt.Unt(1)
			e.Tmes = tmes.New()
			for n := unt.Zero; n < x.Tmes.Cnt()-e.Off; n++ {
				if x.Vals.At(n + e.Off).Lss(x.Vals.At(n)) {
					*e.Tmes = append(*e.Tmes, x.Tmes.At(n+e.Off))
				}
			}
			tst.TmesEql(t, e.Tmes, a.Tmes, "CndInrLss.Tmes")
			ap.Cls()
		})
	}
}
func TestHstStmInrGtr(t *testing.T) {
	for _, instr := range tst.HstPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := hst.Oan()
			instr := instr(prv)
			inrvl := tst.HstInstrInrvlI(instr, 10)
			side := tst.HstInrvlSideBid(inrvl)
			stm := tst.HstSideStmRteLst(side)
			a := stm.InrGtr(unt.Unt(1)).(*hst.CndInrGtr)
			tst.HstCndInrGtrNotZero(t, a)
			x, e := stm.Bse(), &hst.CndInrGtr{}
			e.Off = unt.Unt(1)
			e.Tmes = tmes.New()
			for n := unt.Zero; n < x.Tmes.Cnt()-e.Off; n++ {
				if x.Vals.At(n + e.Off).Gtr(x.Vals.At(n)) {
					*e.Tmes = append(*e.Tmes, x.Tmes.At(n+e.Off))
				}
			}
			tst.TmesEql(t, e.Tmes, a.Tmes, "CndInrGtr.Tmes")
			ap.Cls()
		})
	}
}
func TestHstStmInrLeq(t *testing.T) {
	for _, instr := range tst.HstPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := hst.Oan()
			instr := instr(prv)
			inrvl := tst.HstInstrInrvlI(instr, 10)
			side := tst.HstInrvlSideBid(inrvl)
			stm := tst.HstSideStmRteLst(side)
			a := stm.InrLeq(unt.Unt(1)).(*hst.CndInrLeq)
			tst.HstCndInrLeqNotZero(t, a)
			x, e := stm.Bse(), &hst.CndInrLeq{}
			e.Off = unt.Unt(1)
			e.Tmes = tmes.New()
			for n := unt.Zero; n < x.Tmes.Cnt()-e.Off; n++ {
				if x.Vals.At(n + e.Off).Leq(x.Vals.At(n)) {
					*e.Tmes = append(*e.Tmes, x.Tmes.At(n+e.Off))
				}
			}
			tst.TmesEql(t, e.Tmes, a.Tmes, "CndInrLeq.Tmes")
			ap.Cls()
		})
	}
}
func TestHstStmInrGeq(t *testing.T) {
	for _, instr := range tst.HstPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := hst.Oan()
			instr := instr(prv)
			inrvl := tst.HstInstrInrvlI(instr, 10)
			side := tst.HstInrvlSideBid(inrvl)
			stm := tst.HstSideStmRteLst(side)
			a := stm.InrGeq(unt.Unt(1)).(*hst.CndInrGeq)
			tst.HstCndInrGeqNotZero(t, a)
			x, e := stm.Bse(), &hst.CndInrGeq{}
			e.Off = unt.Unt(1)
			e.Tmes = tmes.New()
			for n := unt.Zero; n < x.Tmes.Cnt()-e.Off; n++ {
				if x.Vals.At(n + e.Off).Geq(x.Vals.At(n)) {
					*e.Tmes = append(*e.Tmes, x.Tmes.At(n+e.Off))
				}
			}
			tst.TmesEql(t, e.Tmes, a.Tmes, "CndInrGeq.Tmes")
			ap.Cls()
		})
	}
}
func TestHstStmOtrEql(t *testing.T) {
	for _, instr := range tst.HstPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := hst.Oan()
			instr := instr(prv)
			inrvl := tst.HstInstrInrvlI(instr, 10)
			side := tst.HstInrvlSideBid(inrvl)
			stm := tst.HstSideStmRteLst(side)
			a := stm.OtrEql(unt.Unt(1), tst.HstStmStmInrAdd(tst.HstSideStmRteLst(tst.HstInrvlSideBid(tst.HstInstrInrvlI(instr, 10))), 4)).(*hst.CndOtrEql)
			tst.HstCndOtrEqlNotZero(t, a)
			ap.Cls()
		})
	}
}
func TestHstStmOtrNeq(t *testing.T) {
	for _, instr := range tst.HstPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := hst.Oan()
			instr := instr(prv)
			inrvl := tst.HstInstrInrvlI(instr, 10)
			side := tst.HstInrvlSideBid(inrvl)
			stm := tst.HstSideStmRteLst(side)
			a := stm.OtrNeq(unt.Unt(1), tst.HstStmStmInrAdd(tst.HstSideStmRteLst(tst.HstInrvlSideBid(tst.HstInstrInrvlI(instr, 10))), 4)).(*hst.CndOtrNeq)
			tst.HstCndOtrNeqNotZero(t, a)
			ap.Cls()
		})
	}
}
func TestHstStmOtrLss(t *testing.T) {
	for _, instr := range tst.HstPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := hst.Oan()
			instr := instr(prv)
			inrvl := tst.HstInstrInrvlI(instr, 10)
			side := tst.HstInrvlSideBid(inrvl)
			stm := tst.HstSideStmRteLst(side)
			a := stm.OtrLss(unt.Unt(1), tst.HstStmStmInrAdd(tst.HstSideStmRteLst(tst.HstInrvlSideBid(tst.HstInstrInrvlI(instr, 10))), 4)).(*hst.CndOtrLss)
			tst.HstCndOtrLssNotZero(t, a)
			ap.Cls()
		})
	}
}
func TestHstStmOtrGtr(t *testing.T) {
	for _, instr := range tst.HstPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := hst.Oan()
			instr := instr(prv)
			inrvl := tst.HstInstrInrvlI(instr, 10)
			side := tst.HstInrvlSideBid(inrvl)
			stm := tst.HstSideStmRteLst(side)
			a := stm.OtrGtr(unt.Unt(1), tst.HstStmStmInrAdd(tst.HstSideStmRteLst(tst.HstInrvlSideBid(tst.HstInstrInrvlI(instr, 10))), 4)).(*hst.CndOtrGtr)
			tst.HstCndOtrGtrNotZero(t, a)
			ap.Cls()
		})
	}
}
func TestHstStmOtrLeq(t *testing.T) {
	for _, instr := range tst.HstPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := hst.Oan()
			instr := instr(prv)
			inrvl := tst.HstInstrInrvlI(instr, 10)
			side := tst.HstInrvlSideBid(inrvl)
			stm := tst.HstSideStmRteLst(side)
			a := stm.OtrLeq(unt.Unt(1), tst.HstStmStmInrAdd(tst.HstSideStmRteLst(tst.HstInrvlSideBid(tst.HstInstrInrvlI(instr, 10))), 4)).(*hst.CndOtrLeq)
			tst.HstCndOtrLeqNotZero(t, a)
			ap.Cls()
		})
	}
}
func TestHstStmOtrGeq(t *testing.T) {
	for _, instr := range tst.HstPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := hst.Oan()
			instr := instr(prv)
			inrvl := tst.HstInstrInrvlI(instr, 10)
			side := tst.HstInrvlSideBid(inrvl)
			stm := tst.HstSideStmRteLst(side)
			a := stm.OtrGeq(unt.Unt(1), tst.HstStmStmInrAdd(tst.HstSideStmRteLst(tst.HstInrvlSideBid(tst.HstInstrInrvlI(instr, 10))), 4)).(*hst.CndOtrGeq)
			tst.HstCndOtrGeqNotZero(t, a)
			ap.Cls()
		})
	}
}
func TestHstCndAnd(t *testing.T) {
	for _, instr := range tst.HstPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := hst.Oan()
			instr := instr(prv)
			inrvl := tst.HstInstrInrvlI(instr, 10)
			side := tst.HstInrvlSideBid(inrvl)
			stm := tst.HstSideStmRteLst(side)
			cnd := tst.HstStmCndInrGtr(stm, 1)
			a := cnd.And(tst.HstStmCndInrGtr(tst.HstStmStmInrAdd(tst.HstSideStmRteLst(tst.HstInrvlSideBid(tst.HstInstrInrvlI(instr, 10))), 1), 2)).(*hst.CndCnd1And)
			tst.HstCndCnd1AndNotZero(t, a)
			ap.Cls()
		})
	}
}
func TestHstCndSeq(t *testing.T) {
	for _, instr := range tst.HstPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := hst.Oan()
			instr := instr(prv)
			inrvl := tst.HstInstrInrvlI(instr, 10)
			side := tst.HstInrvlSideBid(inrvl)
			stm := tst.HstSideStmRteLst(side)
			cnd := tst.HstStmCndInrGtr(stm, 1)
			a := cnd.Seq(tme.Tme(1), tst.HstStmCndInrLss(tst.HstStmStmInrAdd(tst.HstSideStmRteLst(tst.HstInrvlSideBid(tst.HstInstrInrvlI(instr, 10))), 1), 2)).(*hst.CndCnd2Seq)
			tst.HstCndCnd2SeqNotZero(t, a)
			ap.Cls()
		})
	}
}
