package hst_test

import (
	"sys/ana/hst"
	"sys/app"
	"sys/bsc/flt"
	"sys/bsc/flts"
	"sys/bsc/unt"
	"sys/tst"
	"testing"
)

func TestHstSideFst(t *testing.T) {
	for _, instr := range tst.HstPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := hst.Oan()
			instr := instr(prv)
			inrvl := tst.HstInstrInrvlI(instr, 10)
			side := tst.HstInrvlSideBid(inrvl)
			a := side.Fst().(*hst.StmRteFst)
			tst.HstStmRteFstNotZero(t, a)
			ap.Cls()
		})
	}
}
func TestHstSideLst(t *testing.T) {
	for _, instr := range tst.HstPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := hst.Oan()
			instr := instr(prv)
			inrvl := tst.HstInstrInrvlI(instr, 10)
			side := tst.HstInrvlSideBid(inrvl)
			a := side.Lst().(*hst.StmRteLst)
			tst.HstStmRteLstNotZero(t, a)
			ap.Cls()
		})
	}
}
func TestHstSideSum(t *testing.T) {
	for _, instr := range tst.HstPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := hst.Oan()
			instr := instr(prv)
			inrvl := tst.HstInstrInrvlI(instr, 10)
			side := tst.HstInrvlSideBid(inrvl)
			a := side.Sum().(*hst.StmRteSum)
			tst.HstStmRteSumNotZero(t, a)
			ap.Cls()
		})
	}
}
func TestHstSidePrd(t *testing.T) {
	for _, instr := range tst.HstPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := hst.Oan()
			instr := instr(prv)
			inrvl := tst.HstInstrInrvlI(instr, 10)
			side := tst.HstInrvlSideBid(inrvl)
			a := side.Prd().(*hst.StmRtePrd)
			tst.HstStmRtePrdNotZero(t, a)
			ap.Cls()
		})
	}
}
func TestHstSideMin(t *testing.T) {
	for _, instr := range tst.HstPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := hst.Oan()
			instr := instr(prv)
			inrvl := tst.HstInstrInrvlI(instr, 10)
			side := tst.HstInrvlSideBid(inrvl)
			a := side.Min().(*hst.StmRteMin)
			tst.HstStmRteMinNotZero(t, a)
			ap.Cls()
		})
	}
}
func TestHstSideMax(t *testing.T) {
	for _, instr := range tst.HstPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := hst.Oan()
			instr := instr(prv)
			inrvl := tst.HstInstrInrvlI(instr, 10)
			side := tst.HstInrvlSideBid(inrvl)
			a := side.Max().(*hst.StmRteMax)
			tst.HstStmRteMaxNotZero(t, a)
			ap.Cls()
		})
	}
}
func TestHstSideMid(t *testing.T) {
	for _, instr := range tst.HstPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := hst.Oan()
			instr := instr(prv)
			inrvl := tst.HstInstrInrvlI(instr, 10)
			side := tst.HstInrvlSideBid(inrvl)
			a := side.Mid().(*hst.StmRteMid)
			tst.HstStmRteMidNotZero(t, a)
			ap.Cls()
		})
	}
}
func TestHstSideMdn(t *testing.T) {
	for _, instr := range tst.HstPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := hst.Oan()
			instr := instr(prv)
			inrvl := tst.HstInstrInrvlI(instr, 10)
			side := tst.HstInrvlSideBid(inrvl)
			a := side.Mdn().(*hst.StmRteMdn)
			tst.HstStmRteMdnNotZero(t, a)
			ap.Cls()
		})
	}
}
func TestHstSideSma(t *testing.T) {
	for _, instr := range tst.HstPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := hst.Oan()
			instr := instr(prv)
			inrvl := tst.HstInstrInrvlI(instr, 10)
			side := tst.HstInrvlSideBid(inrvl)
			a := side.Sma().(*hst.StmRteSma)
			tst.HstStmRteSmaNotZero(t, a)
			ap.Cls()
		})
	}
}
func TestHstSideGma(t *testing.T) {
	for _, instr := range tst.HstPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := hst.Oan()
			instr := instr(prv)
			inrvl := tst.HstInstrInrvlI(instr, 10)
			side := tst.HstInrvlSideBid(inrvl)
			a := side.Gma().(*hst.StmRteGma)
			tst.HstStmRteGmaNotZero(t, a)
			ap.Cls()
		})
	}
}
func TestHstSideWma(t *testing.T) {
	for _, instr := range tst.HstPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := hst.Oan()
			instr := instr(prv)
			inrvl := tst.HstInstrInrvlI(instr, 10)
			side := tst.HstInrvlSideBid(inrvl)
			a := side.Wma().(*hst.StmRteWma)
			tst.HstStmRteWmaNotZero(t, a)
			ap.Cls()
		})
	}
}
func TestHstSideRsi(t *testing.T) {
	for _, instr := range tst.HstPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := hst.Oan()
			instr := instr(prv)
			inrvl := tst.HstInstrInrvlI(instr, 10)
			side := tst.HstInrvlSideBid(inrvl)
			a := side.Rsi().(*hst.StmRteRsi)
			tst.HstStmRteRsiNotZero(t, a)
			ap.Cls()
		})
	}
}
func TestHstSideWrsi(t *testing.T) {
	for _, instr := range tst.HstPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := hst.Oan()
			instr := instr(prv)
			inrvl := tst.HstInstrInrvlI(instr, 10)
			side := tst.HstInrvlSideBid(inrvl)
			a := side.Wrsi().(*hst.StmRteWrsi)
			tst.HstStmRteWrsiNotZero(t, a)
			ap.Cls()
		})
	}
}
func TestHstSideAlma(t *testing.T) {
	for _, instr := range tst.HstPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := hst.Oan()
			instr := instr(prv)
			inrvl := tst.HstInstrInrvlI(instr, 10)
			side := tst.HstInrvlSideBid(inrvl)
			a := side.Alma().(*hst.StmRteAlma)
			tst.HstStmRteAlmaNotZero(t, a)
			ap.Cls()
		})
	}
}
func TestHstSideVrnc(t *testing.T) {
	for _, instr := range tst.HstPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := hst.Oan()
			instr := instr(prv)
			inrvl := tst.HstInstrInrvlI(instr, 10)
			side := tst.HstInrvlSideBid(inrvl)
			a := side.Vrnc().(*hst.StmRteVrnc)
			tst.HstStmRteVrncNotZero(t, a)
			ap.Cls()
		})
	}
}
func TestHstSideStd(t *testing.T) {
	for _, instr := range tst.HstPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := hst.Oan()
			instr := instr(prv)
			inrvl := tst.HstInstrInrvlI(instr, 10)
			side := tst.HstInrvlSideBid(inrvl)
			a := side.Std().(*hst.StmRteStd)
			tst.HstStmRteStdNotZero(t, a)
			ap.Cls()
		})
	}
}
func TestHstSideRngFul(t *testing.T) {
	for _, instr := range tst.HstPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := hst.Oan()
			instr := instr(prv)
			inrvl := tst.HstInstrInrvlI(instr, 10)
			side := tst.HstInrvlSideBid(inrvl)
			a := side.RngFul().(*hst.StmRteRngFul)
			tst.HstStmRteRngFulNotZero(t, a)
			ap.Cls()
		})
	}
}
func TestHstSideRngLst(t *testing.T) {
	for _, instr := range tst.HstPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := hst.Oan()
			instr := instr(prv)
			inrvl := tst.HstInstrInrvlI(instr, 10)
			side := tst.HstInrvlSideBid(inrvl)
			a := side.RngLst().(*hst.StmRteRngLst)
			tst.HstStmRteRngLstNotZero(t, a)
			ap.Cls()
		})
	}
}
func TestHstSideProLst(t *testing.T) {
	for _, instr := range tst.HstPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := hst.Oan()
			instr := instr(prv)
			inrvl := tst.HstInstrInrvlI(instr, 10)
			side := tst.HstInrvlSideBid(inrvl)
			a := side.ProLst().(*hst.StmRteProLst)
			tst.HstStmRteProLstNotZero(t, a)
			ap.Cls()
		})
	}
}
func TestHstSideProSma(t *testing.T) {
	for _, instr := range tst.HstPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := hst.Oan()
			instr := instr(prv)
			inrvl := tst.HstInstrInrvlI(instr, 10)
			side := tst.HstInrvlSideBid(inrvl)
			a := side.ProSma().(*hst.StmRteProSma)
			tst.HstStmRteProSmaNotZero(t, a)
			ap.Cls()
		})
	}
}
func TestHstSideProAlma(t *testing.T) {
	for _, instr := range tst.HstPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := hst.Oan()
			instr := instr(prv)
			inrvl := tst.HstInstrInrvlI(instr, 10)
			side := tst.HstInrvlSideBid(inrvl)
			a := side.ProAlma().(*hst.StmRteProAlma)
			tst.HstStmRteProAlmaNotZero(t, a)
			ap.Cls()
		})
	}
}
func TestHstSideSar(t *testing.T) {
	for _, instr := range tst.HstPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := hst.Oan()
			instr := instr(prv)
			inrvl := tst.HstInstrInrvlI(instr, 10)
			side := tst.HstInrvlSideBid(inrvl)
			a := side.Sar(flt.Flt(0.02), flt.Flt(0.2)).(*hst.StmRte1Sar)
			tst.HstStmRte1SarNotZero(t, a)
			ap.Cls()
		})
	}
}
func TestHstSideEma(t *testing.T) {
	for _, instr := range tst.HstPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := hst.Oan()
			instr := instr(prv)
			inrvl := tst.HstInstrInrvlI(instr, 10)
			side := tst.HstInrvlSideBid(inrvl)
			a := side.Ema().(*hst.StmRteEma)
			tst.HstStmRteEmaNotZero(t, a)
			ap.Cls()
		})
	}
}
func TestHstStmUnaPos(t *testing.T) {
	for _, instr := range tst.HstPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := hst.Oan()
			instr := instr(prv)
			inrvl := tst.HstInstrInrvlI(instr, 10)
			side := tst.HstInrvlSideBid(inrvl)
			stm := tst.HstSideStmRteLst(side)
			a := stm.UnaPos().(*hst.StmUnaPos)
			tst.HstStmUnaPosNotZero(t, a)
			ap.Cls()
		})
	}
}
func TestHstStmUnaNeg(t *testing.T) {
	for _, instr := range tst.HstPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := hst.Oan()
			instr := instr(prv)
			inrvl := tst.HstInstrInrvlI(instr, 10)
			side := tst.HstInrvlSideBid(inrvl)
			stm := tst.HstSideStmRteLst(side)
			a := stm.UnaNeg().(*hst.StmUnaNeg)
			tst.HstStmUnaNegNotZero(t, a)
			ap.Cls()
		})
	}
}
func TestHstStmUnaInv(t *testing.T) {
	for _, instr := range tst.HstPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := hst.Oan()
			instr := instr(prv)
			inrvl := tst.HstInstrInrvlI(instr, 10)
			side := tst.HstInrvlSideBid(inrvl)
			stm := tst.HstSideStmRteLst(side)
			a := stm.UnaInv().(*hst.StmUnaInv)
			tst.HstStmUnaInvNotZero(t, a)
			ap.Cls()
		})
	}
}
func TestHstStmUnaSqr(t *testing.T) {
	for _, instr := range tst.HstPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := hst.Oan()
			instr := instr(prv)
			inrvl := tst.HstInstrInrvlI(instr, 10)
			side := tst.HstInrvlSideBid(inrvl)
			stm := tst.HstSideStmRteLst(side)
			a := stm.UnaSqr().(*hst.StmUnaSqr)
			tst.HstStmUnaSqrNotZero(t, a)
			ap.Cls()
		})
	}
}
func TestHstStmUnaSqrt(t *testing.T) {
	for _, instr := range tst.HstPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := hst.Oan()
			instr := instr(prv)
			inrvl := tst.HstInstrInrvlI(instr, 10)
			side := tst.HstInrvlSideBid(inrvl)
			stm := tst.HstSideStmRteLst(side)
			a := stm.UnaSqrt().(*hst.StmUnaSqrt)
			tst.HstStmUnaSqrtNotZero(t, a)
			ap.Cls()
		})
	}
}
func TestHstStmSclAdd(t *testing.T) {
	for _, instr := range tst.HstPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := hst.Oan()
			instr := instr(prv)
			inrvl := tst.HstInstrInrvlI(instr, 10)
			side := tst.HstInrvlSideBid(inrvl)
			stm := tst.HstSideStmRteLst(side)
			a := stm.SclAdd(flt.Flt(1.1)).(*hst.StmSclAdd)
			tst.HstStmSclAddNotZero(t, a)
			ap.Cls()
		})
	}
}
func TestHstStmSclSub(t *testing.T) {
	for _, instr := range tst.HstPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := hst.Oan()
			instr := instr(prv)
			inrvl := tst.HstInstrInrvlI(instr, 10)
			side := tst.HstInrvlSideBid(inrvl)
			stm := tst.HstSideStmRteLst(side)
			a := stm.SclSub(flt.Flt(1.1)).(*hst.StmSclSub)
			tst.HstStmSclSubNotZero(t, a)
			ap.Cls()
		})
	}
}
func TestHstStmSclMul(t *testing.T) {
	for _, instr := range tst.HstPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := hst.Oan()
			instr := instr(prv)
			inrvl := tst.HstInstrInrvlI(instr, 10)
			side := tst.HstInrvlSideBid(inrvl)
			stm := tst.HstSideStmRteLst(side)
			a := stm.SclMul(flt.Flt(1.1)).(*hst.StmSclMul)
			tst.HstStmSclMulNotZero(t, a)
			ap.Cls()
		})
	}
}
func TestHstStmSclDiv(t *testing.T) {
	for _, instr := range tst.HstPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := hst.Oan()
			instr := instr(prv)
			inrvl := tst.HstInstrInrvlI(instr, 10)
			side := tst.HstInrvlSideBid(inrvl)
			stm := tst.HstSideStmRteLst(side)
			a := stm.SclDiv(flt.Flt(1.1)).(*hst.StmSclDiv)
			tst.HstStmSclDivNotZero(t, a)
			ap.Cls()
		})
	}
}
func TestHstStmSclRem(t *testing.T) {
	for _, instr := range tst.HstPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := hst.Oan()
			instr := instr(prv)
			inrvl := tst.HstInstrInrvlI(instr, 10)
			side := tst.HstInrvlSideBid(inrvl)
			stm := tst.HstSideStmRteLst(side)
			a := stm.SclRem(flt.Flt(1.1)).(*hst.StmSclRem)
			tst.HstStmSclRemNotZero(t, a)
			ap.Cls()
		})
	}
}
func TestHstStmSclPow(t *testing.T) {
	for _, instr := range tst.HstPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := hst.Oan()
			instr := instr(prv)
			inrvl := tst.HstInstrInrvlI(instr, 10)
			side := tst.HstInrvlSideBid(inrvl)
			stm := tst.HstSideStmRteLst(side)
			a := stm.SclPow(flt.Flt(1.1)).(*hst.StmSclPow)
			tst.HstStmSclPowNotZero(t, a)
			ap.Cls()
		})
	}
}
func TestHstStmSclMin(t *testing.T) {
	for _, instr := range tst.HstPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := hst.Oan()
			instr := instr(prv)
			inrvl := tst.HstInstrInrvlI(instr, 10)
			side := tst.HstInrvlSideBid(inrvl)
			stm := tst.HstSideStmRteLst(side)
			a := stm.SclMin(flt.Flt(1.1)).(*hst.StmSclMin)
			tst.HstStmSclMinNotZero(t, a)
			ap.Cls()
		})
	}
}
func TestHstStmSclMax(t *testing.T) {
	for _, instr := range tst.HstPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := hst.Oan()
			instr := instr(prv)
			inrvl := tst.HstInstrInrvlI(instr, 10)
			side := tst.HstInrvlSideBid(inrvl)
			stm := tst.HstSideStmRteLst(side)
			a := stm.SclMax(flt.Flt(1.1)).(*hst.StmSclMax)
			tst.HstStmSclMaxNotZero(t, a)
			ap.Cls()
		})
	}
}
func TestHstStmSelEql(t *testing.T) {
	for _, instr := range tst.HstPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := hst.Oan()
			instr := instr(prv)
			inrvl := tst.HstInstrInrvlI(instr, 10)
			side := tst.HstInrvlSideBid(inrvl)
			stm := tst.HstSideStmRteLst(side)
			a := stm.SelEql(flt.Flt(1.3171)).(*hst.StmSelEql)
			tst.HstStmSelEqlNotZero(t, a)
			ap.Cls()
		})
	}
}
func TestHstStmSelNeq(t *testing.T) {
	for _, instr := range tst.HstPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := hst.Oan()
			instr := instr(prv)
			inrvl := tst.HstInstrInrvlI(instr, 10)
			side := tst.HstInrvlSideBid(inrvl)
			stm := tst.HstSideStmRteLst(side)
			a := stm.SelNeq(flt.Flt(1.3171)).(*hst.StmSelNeq)
			tst.HstStmSelNeqNotZero(t, a)
			ap.Cls()
		})
	}
}
func TestHstStmSelLss(t *testing.T) {
	for _, instr := range tst.HstPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := hst.Oan()
			instr := instr(prv)
			inrvl := tst.HstInstrInrvlI(instr, 10)
			side := tst.HstInrvlSideBid(inrvl)
			stm := tst.HstSideStmRteLst(side)
			a := stm.SelLss(flt.Flt(1.3171)).(*hst.StmSelLss)
			tst.HstStmSelLssNotZero(t, a)
			ap.Cls()
		})
	}
}
func TestHstStmSelGtr(t *testing.T) {
	for _, instr := range tst.HstPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := hst.Oan()
			instr := instr(prv)
			inrvl := tst.HstInstrInrvlI(instr, 10)
			side := tst.HstInrvlSideBid(inrvl)
			stm := tst.HstSideStmRteLst(side)
			a := stm.SelGtr(flt.Flt(1.3171)).(*hst.StmSelGtr)
			tst.HstStmSelGtrNotZero(t, a)
			ap.Cls()
		})
	}
}
func TestHstStmSelLeq(t *testing.T) {
	for _, instr := range tst.HstPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := hst.Oan()
			instr := instr(prv)
			inrvl := tst.HstInstrInrvlI(instr, 10)
			side := tst.HstInrvlSideBid(inrvl)
			stm := tst.HstSideStmRteLst(side)
			a := stm.SelLeq(flt.Flt(1.3171)).(*hst.StmSelLeq)
			tst.HstStmSelLeqNotZero(t, a)
			ap.Cls()
		})
	}
}
func TestHstStmSelGeq(t *testing.T) {
	for _, instr := range tst.HstPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := hst.Oan()
			instr := instr(prv)
			inrvl := tst.HstInstrInrvlI(instr, 10)
			side := tst.HstInrvlSideBid(inrvl)
			stm := tst.HstSideStmRteLst(side)
			a := stm.SelGeq(flt.Flt(1.3171)).(*hst.StmSelGeq)
			tst.HstStmSelGeqNotZero(t, a)
			ap.Cls()
		})
	}
}
func TestHstStmAggFst(t *testing.T) {
	for _, instr := range tst.HstPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := hst.Oan()
			instr := instr(prv)
			inrvl := tst.HstInstrInrvlI(instr, 10)
			side := tst.HstInrvlSideBid(inrvl)
			stm := tst.HstSideStmRteLst(side)
			a := stm.AggFst(unt.Unt(2)).(*hst.StmAggFst)
			tst.HstStmAggFstNotZero(t, a)
			ap.Cls()
		})
	}
}
func TestHstStmAggLst(t *testing.T) {
	for _, instr := range tst.HstPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := hst.Oan()
			instr := instr(prv)
			inrvl := tst.HstInstrInrvlI(instr, 10)
			side := tst.HstInrvlSideBid(inrvl)
			stm := tst.HstSideStmRteLst(side)
			a := stm.AggLst(unt.Unt(2)).(*hst.StmAggLst)
			tst.HstStmAggLstNotZero(t, a)
			ap.Cls()
		})
	}
}
func TestHstStmAggSum(t *testing.T) {
	for _, instr := range tst.HstPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := hst.Oan()
			instr := instr(prv)
			inrvl := tst.HstInstrInrvlI(instr, 10)
			side := tst.HstInrvlSideBid(inrvl)
			stm := tst.HstSideStmRteLst(side)
			a := stm.AggSum(unt.Unt(2)).(*hst.StmAggSum)
			tst.HstStmAggSumNotZero(t, a)
			ap.Cls()
		})
	}
}
func TestHstStmAggPrd(t *testing.T) {
	for _, instr := range tst.HstPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := hst.Oan()
			instr := instr(prv)
			inrvl := tst.HstInstrInrvlI(instr, 10)
			side := tst.HstInrvlSideBid(inrvl)
			stm := tst.HstSideStmRteLst(side)
			a := stm.AggPrd(unt.Unt(2)).(*hst.StmAggPrd)
			tst.HstStmAggPrdNotZero(t, a)
			ap.Cls()
		})
	}
}
func TestHstStmAggMin(t *testing.T) {
	for _, instr := range tst.HstPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := hst.Oan()
			instr := instr(prv)
			inrvl := tst.HstInstrInrvlI(instr, 10)
			side := tst.HstInrvlSideBid(inrvl)
			stm := tst.HstSideStmRteLst(side)
			a := stm.AggMin(unt.Unt(2)).(*hst.StmAggMin)
			tst.HstStmAggMinNotZero(t, a)
			ap.Cls()
		})
	}
}
func TestHstStmAggMax(t *testing.T) {
	for _, instr := range tst.HstPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := hst.Oan()
			instr := instr(prv)
			inrvl := tst.HstInstrInrvlI(instr, 10)
			side := tst.HstInrvlSideBid(inrvl)
			stm := tst.HstSideStmRteLst(side)
			a := stm.AggMax(unt.Unt(2)).(*hst.StmAggMax)
			tst.HstStmAggMaxNotZero(t, a)
			ap.Cls()
		})
	}
}
func TestHstStmAggMid(t *testing.T) {
	for _, instr := range tst.HstPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := hst.Oan()
			instr := instr(prv)
			inrvl := tst.HstInstrInrvlI(instr, 10)
			side := tst.HstInrvlSideBid(inrvl)
			stm := tst.HstSideStmRteLst(side)
			a := stm.AggMid(unt.Unt(2)).(*hst.StmAggMid)
			tst.HstStmAggMidNotZero(t, a)
			ap.Cls()
		})
	}
}
func TestHstStmAggMdn(t *testing.T) {
	for _, instr := range tst.HstPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := hst.Oan()
			instr := instr(prv)
			inrvl := tst.HstInstrInrvlI(instr, 10)
			side := tst.HstInrvlSideBid(inrvl)
			stm := tst.HstSideStmRteLst(side)
			a := stm.AggMdn(unt.Unt(2)).(*hst.StmAggMdn)
			tst.HstStmAggMdnNotZero(t, a)
			ap.Cls()
		})
	}
}
func TestHstStmAggSma(t *testing.T) {
	for _, instr := range tst.HstPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := hst.Oan()
			instr := instr(prv)
			inrvl := tst.HstInstrInrvlI(instr, 10)
			side := tst.HstInrvlSideBid(inrvl)
			stm := tst.HstSideStmRteLst(side)
			a := stm.AggSma(unt.Unt(2)).(*hst.StmAggSma)
			tst.HstStmAggSmaNotZero(t, a)
			ap.Cls()
		})
	}
}
func TestHstStmAggGma(t *testing.T) {
	for _, instr := range tst.HstPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := hst.Oan()
			instr := instr(prv)
			inrvl := tst.HstInstrInrvlI(instr, 10)
			side := tst.HstInrvlSideBid(inrvl)
			stm := tst.HstSideStmRteLst(side)
			a := stm.AggGma(unt.Unt(2)).(*hst.StmAggGma)
			tst.HstStmAggGmaNotZero(t, a)
			ap.Cls()
		})
	}
}
func TestHstStmAggWma(t *testing.T) {
	for _, instr := range tst.HstPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := hst.Oan()
			instr := instr(prv)
			inrvl := tst.HstInstrInrvlI(instr, 10)
			side := tst.HstInrvlSideBid(inrvl)
			stm := tst.HstSideStmRteLst(side)
			a := stm.AggWma(unt.Unt(2)).(*hst.StmAggWma)
			tst.HstStmAggWmaNotZero(t, a)
			ap.Cls()
		})
	}
}
func TestHstStmAggRsi(t *testing.T) {
	for _, instr := range tst.HstPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := hst.Oan()
			instr := instr(prv)
			inrvl := tst.HstInstrInrvlI(instr, 10)
			side := tst.HstInrvlSideBid(inrvl)
			stm := tst.HstSideStmRteLst(side)
			a := stm.AggRsi(unt.Unt(2)).(*hst.StmAggRsi)
			tst.HstStmAggRsiNotZero(t, a)
			ap.Cls()
		})
	}
}
func TestHstStmAggWrsi(t *testing.T) {
	for _, instr := range tst.HstPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := hst.Oan()
			instr := instr(prv)
			inrvl := tst.HstInstrInrvlI(instr, 10)
			side := tst.HstInrvlSideBid(inrvl)
			stm := tst.HstSideStmRteLst(side)
			a := stm.AggWrsi(unt.Unt(2)).(*hst.StmAggWrsi)
			tst.HstStmAggWrsiNotZero(t, a)
			ap.Cls()
		})
	}
}
func TestHstStmAggAlma(t *testing.T) {
	for _, instr := range tst.HstPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := hst.Oan()
			instr := instr(prv)
			inrvl := tst.HstInstrInrvlI(instr, 10)
			side := tst.HstInrvlSideBid(inrvl)
			stm := tst.HstSideStmRteLst(side)
			a := stm.AggAlma(unt.Unt(2)).(*hst.StmAggAlma)
			tst.HstStmAggAlmaNotZero(t, a)
			ap.Cls()
		})
	}
}
func TestHstStmAggVrnc(t *testing.T) {
	for _, instr := range tst.HstPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := hst.Oan()
			instr := instr(prv)
			inrvl := tst.HstInstrInrvlI(instr, 10)
			side := tst.HstInrvlSideBid(inrvl)
			stm := tst.HstSideStmRteLst(side)
			a := stm.AggVrnc(unt.Unt(2)).(*hst.StmAggVrnc)
			tst.HstStmAggVrncNotZero(t, a)
			ap.Cls()
		})
	}
}
func TestHstStmAggStd(t *testing.T) {
	for _, instr := range tst.HstPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := hst.Oan()
			instr := instr(prv)
			inrvl := tst.HstInstrInrvlI(instr, 10)
			side := tst.HstInrvlSideBid(inrvl)
			stm := tst.HstSideStmRteLst(side)
			a := stm.AggStd(unt.Unt(2)).(*hst.StmAggStd)
			tst.HstStmAggStdNotZero(t, a)
			ap.Cls()
		})
	}
}
func TestHstStmAggRngFul(t *testing.T) {
	for _, instr := range tst.HstPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := hst.Oan()
			instr := instr(prv)
			inrvl := tst.HstInstrInrvlI(instr, 10)
			side := tst.HstInrvlSideBid(inrvl)
			stm := tst.HstSideStmRteLst(side)
			a := stm.AggRngFul(unt.Unt(2)).(*hst.StmAggRngFul)
			tst.HstStmAggRngFulNotZero(t, a)
			ap.Cls()
		})
	}
}
func TestHstStmAggRngLst(t *testing.T) {
	for _, instr := range tst.HstPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := hst.Oan()
			instr := instr(prv)
			inrvl := tst.HstInstrInrvlI(instr, 10)
			side := tst.HstInrvlSideBid(inrvl)
			stm := tst.HstSideStmRteLst(side)
			a := stm.AggRngLst(unt.Unt(2)).(*hst.StmAggRngLst)
			tst.HstStmAggRngLstNotZero(t, a)
			ap.Cls()
		})
	}
}
func TestHstStmAggProLst(t *testing.T) {
	for _, instr := range tst.HstPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := hst.Oan()
			instr := instr(prv)
			inrvl := tst.HstInstrInrvlI(instr, 10)
			side := tst.HstInrvlSideBid(inrvl)
			stm := tst.HstSideStmRteLst(side)
			a := stm.AggProLst(unt.Unt(2)).(*hst.StmAggProLst)
			tst.HstStmAggProLstNotZero(t, a)
			ap.Cls()
		})
	}
}
func TestHstStmAggProSma(t *testing.T) {
	for _, instr := range tst.HstPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := hst.Oan()
			instr := instr(prv)
			inrvl := tst.HstInstrInrvlI(instr, 10)
			side := tst.HstInrvlSideBid(inrvl)
			stm := tst.HstSideStmRteLst(side)
			a := stm.AggProSma(unt.Unt(2)).(*hst.StmAggProSma)
			tst.HstStmAggProSmaNotZero(t, a)
			ap.Cls()
		})
	}
}
func TestHstStmAggProAlma(t *testing.T) {
	for _, instr := range tst.HstPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := hst.Oan()
			instr := instr(prv)
			inrvl := tst.HstInstrInrvlI(instr, 10)
			side := tst.HstInrvlSideBid(inrvl)
			stm := tst.HstSideStmRteLst(side)
			a := stm.AggProAlma(unt.Unt(2)).(*hst.StmAggProAlma)
			tst.HstStmAggProAlmaNotZero(t, a)
			ap.Cls()
		})
	}
}
func TestHstStmAggEma(t *testing.T) {
	for _, instr := range tst.HstPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := hst.Oan()
			instr := instr(prv)
			inrvl := tst.HstInstrInrvlI(instr, 10)
			side := tst.HstInrvlSideBid(inrvl)
			stm := tst.HstSideStmRteLst(side)
			a := stm.AggEma(unt.Unt(2)).(*hst.StmAggEma)
			tst.HstStmAggEmaNotZero(t, a)
			ap.Cls()
		})
	}
}
func TestHstStmInrAdd(t *testing.T) {
	for _, instr := range tst.HstPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := hst.Oan()
			instr := instr(prv)
			inrvl := tst.HstInstrInrvlI(instr, 10)
			side := tst.HstInrvlSideBid(inrvl)
			stm := tst.HstSideStmRteLst(side)
			a := stm.InrAdd(unt.Unt(1)).(*hst.StmInrAdd)
			tst.HstStmInrAddNotZero(t, a)
			x, e := stm.Bse(), &hst.StmInrAdd{}
			e.Off = unt.Unt(1)
			e.Tmes = x.Tmes.From(e.Off)
			e.Vals = flts.Make(e.Tmes.Cnt())
			for n := unt.Zero; n < e.Tmes.Cnt(); n++ {
				(*e.Vals)[n] = x.Vals.At(n + e.Off).Add(x.Vals.At(n))
			}
			tst.TmesEql(t, e.Tmes, a.Tmes, "StmInrAdd.Tmes")
			tst.FltsEql(t, e.Vals, a.Vals, "StmInrAdd.Vals")
			ap.Cls()
		})
	}
}
func TestHstStmInrSub(t *testing.T) {
	for _, instr := range tst.HstPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := hst.Oan()
			instr := instr(prv)
			inrvl := tst.HstInstrInrvlI(instr, 10)
			side := tst.HstInrvlSideBid(inrvl)
			stm := tst.HstSideStmRteLst(side)
			a := stm.InrSub(unt.Unt(1)).(*hst.StmInrSub)
			tst.HstStmInrSubNotZero(t, a)
			x, e := stm.Bse(), &hst.StmInrSub{}
			e.Off = unt.Unt(1)
			e.Tmes = x.Tmes.From(e.Off)
			e.Vals = flts.Make(e.Tmes.Cnt())
			for n := unt.Zero; n < e.Tmes.Cnt(); n++ {
				(*e.Vals)[n] = x.Vals.At(n + e.Off).Sub(x.Vals.At(n))
			}
			tst.TmesEql(t, e.Tmes, a.Tmes, "StmInrSub.Tmes")
			tst.FltsEql(t, e.Vals, a.Vals, "StmInrSub.Vals")
			ap.Cls()
		})
	}
}
func TestHstStmInrMul(t *testing.T) {
	for _, instr := range tst.HstPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := hst.Oan()
			instr := instr(prv)
			inrvl := tst.HstInstrInrvlI(instr, 10)
			side := tst.HstInrvlSideBid(inrvl)
			stm := tst.HstSideStmRteLst(side)
			a := stm.InrMul(unt.Unt(1)).(*hst.StmInrMul)
			tst.HstStmInrMulNotZero(t, a)
			x, e := stm.Bse(), &hst.StmInrMul{}
			e.Off = unt.Unt(1)
			e.Tmes = x.Tmes.From(e.Off)
			e.Vals = flts.Make(e.Tmes.Cnt())
			for n := unt.Zero; n < e.Tmes.Cnt(); n++ {
				(*e.Vals)[n] = x.Vals.At(n + e.Off).Mul(x.Vals.At(n))
			}
			tst.TmesEql(t, e.Tmes, a.Tmes, "StmInrMul.Tmes")
			tst.FltsEql(t, e.Vals, a.Vals, "StmInrMul.Vals")
			ap.Cls()
		})
	}
}
func TestHstStmInrDiv(t *testing.T) {
	for _, instr := range tst.HstPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := hst.Oan()
			instr := instr(prv)
			inrvl := tst.HstInstrInrvlI(instr, 10)
			side := tst.HstInrvlSideBid(inrvl)
			stm := tst.HstSideStmRteLst(side)
			a := stm.InrDiv(unt.Unt(1)).(*hst.StmInrDiv)
			tst.HstStmInrDivNotZero(t, a)
			x, e := stm.Bse(), &hst.StmInrDiv{}
			e.Off = unt.Unt(1)
			e.Tmes = x.Tmes.From(e.Off)
			e.Vals = flts.Make(e.Tmes.Cnt())
			for n := unt.Zero; n < e.Tmes.Cnt(); n++ {
				(*e.Vals)[n] = x.Vals.At(n + e.Off).Div(x.Vals.At(n))
			}
			tst.TmesEql(t, e.Tmes, a.Tmes, "StmInrDiv.Tmes")
			tst.FltsEql(t, e.Vals, a.Vals, "StmInrDiv.Vals")
			ap.Cls()
		})
	}
}
func TestHstStmInrRem(t *testing.T) {
	for _, instr := range tst.HstPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := hst.Oan()
			instr := instr(prv)
			inrvl := tst.HstInstrInrvlI(instr, 10)
			side := tst.HstInrvlSideBid(inrvl)
			stm := tst.HstSideStmRteLst(side)
			a := stm.InrRem(unt.Unt(1)).(*hst.StmInrRem)
			tst.HstStmInrRemNotZero(t, a)
			x, e := stm.Bse(), &hst.StmInrRem{}
			e.Off = unt.Unt(1)
			e.Tmes = x.Tmes.From(e.Off)
			e.Vals = flts.Make(e.Tmes.Cnt())
			for n := unt.Zero; n < e.Tmes.Cnt(); n++ {
				(*e.Vals)[n] = x.Vals.At(n + e.Off).Rem(x.Vals.At(n))
			}
			tst.TmesEql(t, e.Tmes, a.Tmes, "StmInrRem.Tmes")
			tst.FltsEql(t, e.Vals, a.Vals, "StmInrRem.Vals")
			ap.Cls()
		})
	}
}
func TestHstStmInrPow(t *testing.T) {
	for _, instr := range tst.HstPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := hst.Oan()
			instr := instr(prv)
			inrvl := tst.HstInstrInrvlI(instr, 10)
			side := tst.HstInrvlSideBid(inrvl)
			stm := tst.HstSideStmRteLst(side)
			a := stm.InrPow(unt.Unt(1)).(*hst.StmInrPow)
			tst.HstStmInrPowNotZero(t, a)
			x, e := stm.Bse(), &hst.StmInrPow{}
			e.Off = unt.Unt(1)
			e.Tmes = x.Tmes.From(e.Off)
			e.Vals = flts.Make(e.Tmes.Cnt())
			for n := unt.Zero; n < e.Tmes.Cnt(); n++ {
				(*e.Vals)[n] = x.Vals.At(n + e.Off).Pow(x.Vals.At(n))
			}
			tst.TmesEql(t, e.Tmes, a.Tmes, "StmInrPow.Tmes")
			tst.FltsEql(t, e.Vals, a.Vals, "StmInrPow.Vals")
			ap.Cls()
		})
	}
}
func TestHstStmInrMin(t *testing.T) {
	for _, instr := range tst.HstPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := hst.Oan()
			instr := instr(prv)
			inrvl := tst.HstInstrInrvlI(instr, 10)
			side := tst.HstInrvlSideBid(inrvl)
			stm := tst.HstSideStmRteLst(side)
			a := stm.InrMin(unt.Unt(1)).(*hst.StmInrMin)
			tst.HstStmInrMinNotZero(t, a)
			x, e := stm.Bse(), &hst.StmInrMin{}
			e.Off = unt.Unt(1)
			e.Tmes = x.Tmes.From(e.Off)
			e.Vals = flts.Make(e.Tmes.Cnt())
			for n := unt.Zero; n < e.Tmes.Cnt(); n++ {
				(*e.Vals)[n] = x.Vals.At(n + e.Off).Min(x.Vals.At(n))
			}
			tst.TmesEql(t, e.Tmes, a.Tmes, "StmInrMin.Tmes")
			tst.FltsEql(t, e.Vals, a.Vals, "StmInrMin.Vals")
			ap.Cls()
		})
	}
}
func TestHstStmInrMax(t *testing.T) {
	for _, instr := range tst.HstPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := hst.Oan()
			instr := instr(prv)
			inrvl := tst.HstInstrInrvlI(instr, 10)
			side := tst.HstInrvlSideBid(inrvl)
			stm := tst.HstSideStmRteLst(side)
			a := stm.InrMax(unt.Unt(1)).(*hst.StmInrMax)
			tst.HstStmInrMaxNotZero(t, a)
			x, e := stm.Bse(), &hst.StmInrMax{}
			e.Off = unt.Unt(1)
			e.Tmes = x.Tmes.From(e.Off)
			e.Vals = flts.Make(e.Tmes.Cnt())
			for n := unt.Zero; n < e.Tmes.Cnt(); n++ {
				(*e.Vals)[n] = x.Vals.At(n + e.Off).Max(x.Vals.At(n))
			}
			tst.TmesEql(t, e.Tmes, a.Tmes, "StmInrMax.Tmes")
			tst.FltsEql(t, e.Vals, a.Vals, "StmInrMax.Vals")
			ap.Cls()
		})
	}
}
func TestHstStmInrSlp(t *testing.T) {
	for _, instr := range tst.HstPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := hst.Oan()
			instr := instr(prv)
			inrvl := tst.HstInstrInrvlI(instr, 10)
			side := tst.HstInrvlSideBid(inrvl)
			stm := tst.HstSideStmRteLst(side)
			a := stm.InrSlp(unt.Unt(1)).(*hst.StmInr1Slp)
			tst.HstStmInr1SlpNotZero(t, a)
			ap.Cls()
		})
	}
}
func TestHstStmOtrAdd(t *testing.T) {
	for _, instr := range tst.HstPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := hst.Oan()
			instr := instr(prv)
			inrvl := tst.HstInstrInrvlI(instr, 10)
			side := tst.HstInrvlSideBid(inrvl)
			stm := tst.HstSideStmRteLst(side)
			a := stm.OtrAdd(unt.Unt(1), tst.HstStmStmInrAdd(tst.HstSideStmRteLst(tst.HstInrvlSideBid(tst.HstInstrInrvlI(instr, 10))), 4)).(*hst.StmOtrAdd)
			tst.HstStmOtrAddNotZero(t, a)
			ap.Cls()
		})
	}
}
func TestHstStmOtrSub(t *testing.T) {
	for _, instr := range tst.HstPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := hst.Oan()
			instr := instr(prv)
			inrvl := tst.HstInstrInrvlI(instr, 10)
			side := tst.HstInrvlSideBid(inrvl)
			stm := tst.HstSideStmRteLst(side)
			a := stm.OtrSub(unt.Unt(1), tst.HstStmStmInrAdd(tst.HstSideStmRteLst(tst.HstInrvlSideBid(tst.HstInstrInrvlI(instr, 10))), 4)).(*hst.StmOtrSub)
			tst.HstStmOtrSubNotZero(t, a)
			ap.Cls()
		})
	}
}
func TestHstStmOtrMul(t *testing.T) {
	for _, instr := range tst.HstPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := hst.Oan()
			instr := instr(prv)
			inrvl := tst.HstInstrInrvlI(instr, 10)
			side := tst.HstInrvlSideBid(inrvl)
			stm := tst.HstSideStmRteLst(side)
			a := stm.OtrMul(unt.Unt(1), tst.HstStmStmInrAdd(tst.HstSideStmRteLst(tst.HstInrvlSideBid(tst.HstInstrInrvlI(instr, 10))), 4)).(*hst.StmOtrMul)
			tst.HstStmOtrMulNotZero(t, a)
			ap.Cls()
		})
	}
}
func TestHstStmOtrDiv(t *testing.T) {
	for _, instr := range tst.HstPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := hst.Oan()
			instr := instr(prv)
			inrvl := tst.HstInstrInrvlI(instr, 10)
			side := tst.HstInrvlSideBid(inrvl)
			stm := tst.HstSideStmRteLst(side)
			a := stm.OtrDiv(unt.Unt(1), tst.HstStmStmInrAdd(tst.HstSideStmRteLst(tst.HstInrvlSideBid(tst.HstInstrInrvlI(instr, 10))), 4)).(*hst.StmOtrDiv)
			tst.HstStmOtrDivNotZero(t, a)
			ap.Cls()
		})
	}
}
func TestHstStmOtrRem(t *testing.T) {
	for _, instr := range tst.HstPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := hst.Oan()
			instr := instr(prv)
			inrvl := tst.HstInstrInrvlI(instr, 10)
			side := tst.HstInrvlSideBid(inrvl)
			stm := tst.HstSideStmRteLst(side)
			a := stm.OtrRem(unt.Unt(1), tst.HstStmStmInrAdd(tst.HstSideStmRteLst(tst.HstInrvlSideBid(tst.HstInstrInrvlI(instr, 10))), 4)).(*hst.StmOtrRem)
			tst.HstStmOtrRemNotZero(t, a)
			ap.Cls()
		})
	}
}
func TestHstStmOtrPow(t *testing.T) {
	for _, instr := range tst.HstPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := hst.Oan()
			instr := instr(prv)
			inrvl := tst.HstInstrInrvlI(instr, 10)
			side := tst.HstInrvlSideBid(inrvl)
			stm := tst.HstSideStmRteLst(side)
			a := stm.OtrPow(unt.Unt(1), tst.HstStmStmInrAdd(tst.HstSideStmRteLst(tst.HstInrvlSideBid(tst.HstInstrInrvlI(instr, 10))), 4)).(*hst.StmOtrPow)
			tst.HstStmOtrPowNotZero(t, a)
			ap.Cls()
		})
	}
}
func TestHstStmOtrMin(t *testing.T) {
	for _, instr := range tst.HstPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := hst.Oan()
			instr := instr(prv)
			inrvl := tst.HstInstrInrvlI(instr, 10)
			side := tst.HstInrvlSideBid(inrvl)
			stm := tst.HstSideStmRteLst(side)
			a := stm.OtrMin(unt.Unt(1), tst.HstStmStmInrAdd(tst.HstSideStmRteLst(tst.HstInrvlSideBid(tst.HstInstrInrvlI(instr, 10))), 4)).(*hst.StmOtrMin)
			tst.HstStmOtrMinNotZero(t, a)
			ap.Cls()
		})
	}
}
func TestHstStmOtrMax(t *testing.T) {
	for _, instr := range tst.HstPrvInstrs {
		t.Run("", func(t *testing.T) {
			ap := app.New(tst.Cfg)
			ap.Oan.CloneInstrs(tst.Instrs)
			prv := hst.Oan()
			instr := instr(prv)
			inrvl := tst.HstInstrInrvlI(instr, 10)
			side := tst.HstInrvlSideBid(inrvl)
			stm := tst.HstSideStmRteLst(side)
			a := stm.OtrMax(unt.Unt(1), tst.HstStmStmInrAdd(tst.HstSideStmRteLst(tst.HstInrvlSideBid(tst.HstInstrInrvlI(instr, 10))), 4)).(*hst.StmOtrMax)
			tst.HstStmOtrMaxNotZero(t, a)
			ap.Cls()
		})
	}
}
