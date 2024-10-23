package tst

import (
	"fmt"
	"reflect"
	"sys/ana"
	"sys/ana/hst"
	"sys/ana/rlt"
	"sys/bsc/bnd"
	"sys/bsc/bnds"
	"sys/bsc/bol"
	"sys/bsc/bols"
	"sys/bsc/flt"
	"sys/bsc/flts"
	"sys/bsc/int"
	"sys/bsc/ints"
	"sys/bsc/str"
	"sys/bsc/strs"
	"sys/bsc/tme"
	"sys/bsc/tmes"
	"sys/bsc/unt"
	"sys/bsc/unts"
	"testing"
	"time"
)

var (
	HstPrvInstrEurUsd    = func(x hst.Prv, rng ...tme.Rng) hst.Instr { return x.EurUsd(rng...) }
	HstPrvInstrAudUsd    = func(x hst.Prv, rng ...tme.Rng) hst.Instr { return x.AudUsd(rng...) }
	HstPrvInstrNzdUsd    = func(x hst.Prv, rng ...tme.Rng) hst.Instr { return x.NzdUsd(rng...) }
	HstPrvInstrGbpUsd    = func(x hst.Prv, rng ...tme.Rng) hst.Instr { return x.GbpUsd(rng...) }
	HstInstrInrvlI       = func(x hst.Instr, dur tme.Tme) hst.Inrvl { return x.I(dur) }
	HstInrvlSideBid      = func(x hst.Inrvl) hst.Side { return x.Bid() }
	HstInrvlSideAsk      = func(x hst.Inrvl) hst.Side { return x.Ask() }
	HstSideStmRteFst     = func(x hst.Side) hst.Stm { return x.Fst() }
	HstSideStmRteLst     = func(x hst.Side) hst.Stm { return x.Lst() }
	HstSideStmRteSum     = func(x hst.Side) hst.Stm { return x.Sum() }
	HstSideStmRtePrd     = func(x hst.Side) hst.Stm { return x.Prd() }
	HstSideStmRteMin     = func(x hst.Side) hst.Stm { return x.Min() }
	HstSideStmRteMax     = func(x hst.Side) hst.Stm { return x.Max() }
	HstSideStmRteMid     = func(x hst.Side) hst.Stm { return x.Mid() }
	HstSideStmRteMdn     = func(x hst.Side) hst.Stm { return x.Mdn() }
	HstSideStmRteSma     = func(x hst.Side) hst.Stm { return x.Sma() }
	HstSideStmRteGma     = func(x hst.Side) hst.Stm { return x.Gma() }
	HstSideStmRteWma     = func(x hst.Side) hst.Stm { return x.Wma() }
	HstSideStmRteRsi     = func(x hst.Side) hst.Stm { return x.Rsi() }
	HstSideStmRteWrsi    = func(x hst.Side) hst.Stm { return x.Wrsi() }
	HstSideStmRteAlma    = func(x hst.Side) hst.Stm { return x.Alma() }
	HstSideStmRteVrnc    = func(x hst.Side) hst.Stm { return x.Vrnc() }
	HstSideStmRteStd     = func(x hst.Side) hst.Stm { return x.Std() }
	HstSideStmRteRngFul  = func(x hst.Side) hst.Stm { return x.RngFul() }
	HstSideStmRteRngLst  = func(x hst.Side) hst.Stm { return x.RngLst() }
	HstSideStmRteProLst  = func(x hst.Side) hst.Stm { return x.ProLst() }
	HstSideStmRteProSma  = func(x hst.Side) hst.Stm { return x.ProSma() }
	HstSideStmRteProAlma = func(x hst.Side) hst.Stm { return x.ProAlma() }
	HstSideStmRte1Sar    = func(x hst.Side, afInc, afMax flt.Flt) hst.Stm { return x.Sar(afInc, afMax) }
	HstSideStmRteEma     = func(x hst.Side) hst.Stm { return x.Ema() }
	HstStmStmUnaPos      = func(x hst.Stm) hst.Stm { return x.UnaPos() }
	HstStmStmUnaNeg      = func(x hst.Stm) hst.Stm { return x.UnaNeg() }
	HstStmStmUnaInv      = func(x hst.Stm) hst.Stm { return x.UnaInv() }
	HstStmStmUnaSqr      = func(x hst.Stm) hst.Stm { return x.UnaSqr() }
	HstStmStmUnaSqrt     = func(x hst.Stm) hst.Stm { return x.UnaSqrt() }
	HstStmStmSclAdd      = func(x hst.Stm, scl flt.Flt) hst.Stm { return x.SclAdd(scl) }
	HstStmStmSclSub      = func(x hst.Stm, scl flt.Flt) hst.Stm { return x.SclSub(scl) }
	HstStmStmSclMul      = func(x hst.Stm, scl flt.Flt) hst.Stm { return x.SclMul(scl) }
	HstStmStmSclDiv      = func(x hst.Stm, scl flt.Flt) hst.Stm { return x.SclDiv(scl) }
	HstStmStmSclRem      = func(x hst.Stm, scl flt.Flt) hst.Stm { return x.SclRem(scl) }
	HstStmStmSclPow      = func(x hst.Stm, scl flt.Flt) hst.Stm { return x.SclPow(scl) }
	HstStmStmSclMin      = func(x hst.Stm, scl flt.Flt) hst.Stm { return x.SclMin(scl) }
	HstStmStmSclMax      = func(x hst.Stm, scl flt.Flt) hst.Stm { return x.SclMax(scl) }
	HstStmStmSelEql      = func(x hst.Stm, sel flt.Flt) hst.Stm { return x.SelEql(sel) }
	HstStmStmSelNeq      = func(x hst.Stm, sel flt.Flt) hst.Stm { return x.SelNeq(sel) }
	HstStmStmSelLss      = func(x hst.Stm, sel flt.Flt) hst.Stm { return x.SelLss(sel) }
	HstStmStmSelGtr      = func(x hst.Stm, sel flt.Flt) hst.Stm { return x.SelGtr(sel) }
	HstStmStmSelLeq      = func(x hst.Stm, sel flt.Flt) hst.Stm { return x.SelLeq(sel) }
	HstStmStmSelGeq      = func(x hst.Stm, sel flt.Flt) hst.Stm { return x.SelGeq(sel) }
	HstStmStmAggFst      = func(x hst.Stm, length unt.Unt) hst.Stm { return x.AggFst(length) }
	HstStmStmAggLst      = func(x hst.Stm, length unt.Unt) hst.Stm { return x.AggLst(length) }
	HstStmStmAggSum      = func(x hst.Stm, length unt.Unt) hst.Stm { return x.AggSum(length) }
	HstStmStmAggPrd      = func(x hst.Stm, length unt.Unt) hst.Stm { return x.AggPrd(length) }
	HstStmStmAggMin      = func(x hst.Stm, length unt.Unt) hst.Stm { return x.AggMin(length) }
	HstStmStmAggMax      = func(x hst.Stm, length unt.Unt) hst.Stm { return x.AggMax(length) }
	HstStmStmAggMid      = func(x hst.Stm, length unt.Unt) hst.Stm { return x.AggMid(length) }
	HstStmStmAggMdn      = func(x hst.Stm, length unt.Unt) hst.Stm { return x.AggMdn(length) }
	HstStmStmAggSma      = func(x hst.Stm, length unt.Unt) hst.Stm { return x.AggSma(length) }
	HstStmStmAggGma      = func(x hst.Stm, length unt.Unt) hst.Stm { return x.AggGma(length) }
	HstStmStmAggWma      = func(x hst.Stm, length unt.Unt) hst.Stm { return x.AggWma(length) }
	HstStmStmAggRsi      = func(x hst.Stm, length unt.Unt) hst.Stm { return x.AggRsi(length) }
	HstStmStmAggWrsi     = func(x hst.Stm, length unt.Unt) hst.Stm { return x.AggWrsi(length) }
	HstStmStmAggAlma     = func(x hst.Stm, length unt.Unt) hst.Stm { return x.AggAlma(length) }
	HstStmStmAggVrnc     = func(x hst.Stm, length unt.Unt) hst.Stm { return x.AggVrnc(length) }
	HstStmStmAggStd      = func(x hst.Stm, length unt.Unt) hst.Stm { return x.AggStd(length) }
	HstStmStmAggRngFul   = func(x hst.Stm, length unt.Unt) hst.Stm { return x.AggRngFul(length) }
	HstStmStmAggRngLst   = func(x hst.Stm, length unt.Unt) hst.Stm { return x.AggRngLst(length) }
	HstStmStmAggProLst   = func(x hst.Stm, length unt.Unt) hst.Stm { return x.AggProLst(length) }
	HstStmStmAggProSma   = func(x hst.Stm, length unt.Unt) hst.Stm { return x.AggProSma(length) }
	HstStmStmAggProAlma  = func(x hst.Stm, length unt.Unt) hst.Stm { return x.AggProAlma(length) }
	HstStmStmAggEma      = func(x hst.Stm, length unt.Unt) hst.Stm { return x.AggEma(length) }
	HstStmStmInrAdd      = func(x hst.Stm, off unt.Unt) hst.Stm { return x.InrAdd(off) }
	HstStmStmInrSub      = func(x hst.Stm, off unt.Unt) hst.Stm { return x.InrSub(off) }
	HstStmStmInrMul      = func(x hst.Stm, off unt.Unt) hst.Stm { return x.InrMul(off) }
	HstStmStmInrDiv      = func(x hst.Stm, off unt.Unt) hst.Stm { return x.InrDiv(off) }
	HstStmStmInrRem      = func(x hst.Stm, off unt.Unt) hst.Stm { return x.InrRem(off) }
	HstStmStmInrPow      = func(x hst.Stm, off unt.Unt) hst.Stm { return x.InrPow(off) }
	HstStmStmInrMin      = func(x hst.Stm, off unt.Unt) hst.Stm { return x.InrMin(off) }
	HstStmStmInrMax      = func(x hst.Stm, off unt.Unt) hst.Stm { return x.InrMax(off) }
	HstStmStmInr1Slp     = func(x hst.Stm, off unt.Unt) hst.Stm { return x.InrSlp(off) }
	HstStmStmOtrAdd      = func(x hst.Stm, off unt.Unt, a hst.Stm) hst.Stm { return x.OtrAdd(off, a) }
	HstStmStmOtrSub      = func(x hst.Stm, off unt.Unt, a hst.Stm) hst.Stm { return x.OtrSub(off, a) }
	HstStmStmOtrMul      = func(x hst.Stm, off unt.Unt, a hst.Stm) hst.Stm { return x.OtrMul(off, a) }
	HstStmStmOtrDiv      = func(x hst.Stm, off unt.Unt, a hst.Stm) hst.Stm { return x.OtrDiv(off, a) }
	HstStmStmOtrRem      = func(x hst.Stm, off unt.Unt, a hst.Stm) hst.Stm { return x.OtrRem(off, a) }
	HstStmStmOtrPow      = func(x hst.Stm, off unt.Unt, a hst.Stm) hst.Stm { return x.OtrPow(off, a) }
	HstStmStmOtrMin      = func(x hst.Stm, off unt.Unt, a hst.Stm) hst.Stm { return x.OtrMin(off, a) }
	HstStmStmOtrMax      = func(x hst.Stm, off unt.Unt, a hst.Stm) hst.Stm { return x.OtrMax(off, a) }
	HstStmCndSclEql      = func(x hst.Stm, scl flt.Flt) hst.Cnd { return x.SclEql(scl) }
	HstStmCndSclNeq      = func(x hst.Stm, scl flt.Flt) hst.Cnd { return x.SclNeq(scl) }
	HstStmCndSclLss      = func(x hst.Stm, scl flt.Flt) hst.Cnd { return x.SclLss(scl) }
	HstStmCndSclGtr      = func(x hst.Stm, scl flt.Flt) hst.Cnd { return x.SclGtr(scl) }
	HstStmCndSclLeq      = func(x hst.Stm, scl flt.Flt) hst.Cnd { return x.SclLeq(scl) }
	HstStmCndSclGeq      = func(x hst.Stm, scl flt.Flt) hst.Cnd { return x.SclGeq(scl) }
	HstStmCndInrEql      = func(x hst.Stm, off unt.Unt) hst.Cnd { return x.InrEql(off) }
	HstStmCndInrNeq      = func(x hst.Stm, off unt.Unt) hst.Cnd { return x.InrNeq(off) }
	HstStmCndInrLss      = func(x hst.Stm, off unt.Unt) hst.Cnd { return x.InrLss(off) }
	HstStmCndInrGtr      = func(x hst.Stm, off unt.Unt) hst.Cnd { return x.InrGtr(off) }
	HstStmCndInrLeq      = func(x hst.Stm, off unt.Unt) hst.Cnd { return x.InrLeq(off) }
	HstStmCndInrGeq      = func(x hst.Stm, off unt.Unt) hst.Cnd { return x.InrGeq(off) }
	HstStmCndOtrEql      = func(x hst.Stm, off unt.Unt, a hst.Stm) hst.Cnd { return x.OtrEql(off, a) }
	HstStmCndOtrNeq      = func(x hst.Stm, off unt.Unt, a hst.Stm) hst.Cnd { return x.OtrNeq(off, a) }
	HstStmCndOtrLss      = func(x hst.Stm, off unt.Unt, a hst.Stm) hst.Cnd { return x.OtrLss(off, a) }
	HstStmCndOtrGtr      = func(x hst.Stm, off unt.Unt, a hst.Stm) hst.Cnd { return x.OtrGtr(off, a) }
	HstStmCndOtrLeq      = func(x hst.Stm, off unt.Unt, a hst.Stm) hst.Cnd { return x.OtrLeq(off, a) }
	HstStmCndOtrGeq      = func(x hst.Stm, off unt.Unt, a hst.Stm) hst.Cnd { return x.OtrGeq(off, a) }
	HstCndCndCnd1And     = func(x, a hst.Cnd) hst.Cnd { return x.And(a) }
	HstCndCndCnd2Seq     = func(x hst.Cnd, dur tme.Tme, a hst.Cnd) hst.Cnd { return x.Seq(dur, a) }
	HstCndStgyStgy       = func(x hst.Cnd, isLong bol.Bol, prfLim, losLim flt.Flt, durLim tme.Tme, minPnlPct flt.Flt, instr hst.Instr, ftrStms *hst.Stms, clss ...hst.Cnd) hst.Stgy {
		return x.Stgy(isLong, prfLim, losLim, durLim, minPnlPct, instr, ftrStms, clss...)
	}
	RltPrvInstrEurUsd    = func(x rlt.Prv, rng ...tme.Rng) rlt.Instr { return x.EurUsd(rng...) }
	RltPrvInstrAudUsd    = func(x rlt.Prv, rng ...tme.Rng) rlt.Instr { return x.AudUsd(rng...) }
	RltPrvInstrNzdUsd    = func(x rlt.Prv, rng ...tme.Rng) rlt.Instr { return x.NzdUsd(rng...) }
	RltPrvInstrGbpUsd    = func(x rlt.Prv, rng ...tme.Rng) rlt.Instr { return x.GbpUsd(rng...) }
	RltInstrInrvlI       = func(x rlt.Instr, dur tme.Tme) rlt.Inrvl { return x.I(dur) }
	RltInrvlSideBid      = func(x rlt.Inrvl) rlt.Side { return x.Bid() }
	RltInrvlSideAsk      = func(x rlt.Inrvl) rlt.Side { return x.Ask() }
	RltSideStmRteFst     = func(x rlt.Side) rlt.Stm { return x.Fst() }
	RltSideStmRteLst     = func(x rlt.Side) rlt.Stm { return x.Lst() }
	RltSideStmRteSum     = func(x rlt.Side) rlt.Stm { return x.Sum() }
	RltSideStmRtePrd     = func(x rlt.Side) rlt.Stm { return x.Prd() }
	RltSideStmRteMin     = func(x rlt.Side) rlt.Stm { return x.Min() }
	RltSideStmRteMax     = func(x rlt.Side) rlt.Stm { return x.Max() }
	RltSideStmRteMid     = func(x rlt.Side) rlt.Stm { return x.Mid() }
	RltSideStmRteMdn     = func(x rlt.Side) rlt.Stm { return x.Mdn() }
	RltSideStmRteSma     = func(x rlt.Side) rlt.Stm { return x.Sma() }
	RltSideStmRteGma     = func(x rlt.Side) rlt.Stm { return x.Gma() }
	RltSideStmRteWma     = func(x rlt.Side) rlt.Stm { return x.Wma() }
	RltSideStmRteRsi     = func(x rlt.Side) rlt.Stm { return x.Rsi() }
	RltSideStmRteWrsi    = func(x rlt.Side) rlt.Stm { return x.Wrsi() }
	RltSideStmRteAlma    = func(x rlt.Side) rlt.Stm { return x.Alma() }
	RltSideStmRteVrnc    = func(x rlt.Side) rlt.Stm { return x.Vrnc() }
	RltSideStmRteStd     = func(x rlt.Side) rlt.Stm { return x.Std() }
	RltSideStmRteRngFul  = func(x rlt.Side) rlt.Stm { return x.RngFul() }
	RltSideStmRteRngLst  = func(x rlt.Side) rlt.Stm { return x.RngLst() }
	RltSideStmRteProLst  = func(x rlt.Side) rlt.Stm { return x.ProLst() }
	RltSideStmRteProSma  = func(x rlt.Side) rlt.Stm { return x.ProSma() }
	RltSideStmRteProAlma = func(x rlt.Side) rlt.Stm { return x.ProAlma() }
	RltSideStmRte1Sar    = func(x rlt.Side, afInc, afMax flt.Flt) rlt.Stm { return x.Sar(afInc, afMax) }
	RltSideStmRteEma     = func(x rlt.Side) rlt.Stm { return x.Ema() }
	RltStmStmUnaPos      = func(x rlt.Stm) rlt.Stm { return x.UnaPos() }
	RltStmStmUnaNeg      = func(x rlt.Stm) rlt.Stm { return x.UnaNeg() }
	RltStmStmUnaInv      = func(x rlt.Stm) rlt.Stm { return x.UnaInv() }
	RltStmStmUnaSqr      = func(x rlt.Stm) rlt.Stm { return x.UnaSqr() }
	RltStmStmUnaSqrt     = func(x rlt.Stm) rlt.Stm { return x.UnaSqrt() }
	RltStmStmSclAdd      = func(x rlt.Stm, scl flt.Flt) rlt.Stm { return x.SclAdd(scl) }
	RltStmStmSclSub      = func(x rlt.Stm, scl flt.Flt) rlt.Stm { return x.SclSub(scl) }
	RltStmStmSclMul      = func(x rlt.Stm, scl flt.Flt) rlt.Stm { return x.SclMul(scl) }
	RltStmStmSclDiv      = func(x rlt.Stm, scl flt.Flt) rlt.Stm { return x.SclDiv(scl) }
	RltStmStmSclRem      = func(x rlt.Stm, scl flt.Flt) rlt.Stm { return x.SclRem(scl) }
	RltStmStmSclPow      = func(x rlt.Stm, scl flt.Flt) rlt.Stm { return x.SclPow(scl) }
	RltStmStmSclMin      = func(x rlt.Stm, scl flt.Flt) rlt.Stm { return x.SclMin(scl) }
	RltStmStmSclMax      = func(x rlt.Stm, scl flt.Flt) rlt.Stm { return x.SclMax(scl) }
	RltStmStmSelEql      = func(x rlt.Stm, sel flt.Flt) rlt.Stm { return x.SelEql(sel) }
	RltStmStmSelNeq      = func(x rlt.Stm, sel flt.Flt) rlt.Stm { return x.SelNeq(sel) }
	RltStmStmSelLss      = func(x rlt.Stm, sel flt.Flt) rlt.Stm { return x.SelLss(sel) }
	RltStmStmSelGtr      = func(x rlt.Stm, sel flt.Flt) rlt.Stm { return x.SelGtr(sel) }
	RltStmStmSelLeq      = func(x rlt.Stm, sel flt.Flt) rlt.Stm { return x.SelLeq(sel) }
	RltStmStmSelGeq      = func(x rlt.Stm, sel flt.Flt) rlt.Stm { return x.SelGeq(sel) }
	RltStmStmAggFst      = func(x rlt.Stm, length unt.Unt) rlt.Stm { return x.AggFst(length) }
	RltStmStmAggLst      = func(x rlt.Stm, length unt.Unt) rlt.Stm { return x.AggLst(length) }
	RltStmStmAggSum      = func(x rlt.Stm, length unt.Unt) rlt.Stm { return x.AggSum(length) }
	RltStmStmAggPrd      = func(x rlt.Stm, length unt.Unt) rlt.Stm { return x.AggPrd(length) }
	RltStmStmAggMin      = func(x rlt.Stm, length unt.Unt) rlt.Stm { return x.AggMin(length) }
	RltStmStmAggMax      = func(x rlt.Stm, length unt.Unt) rlt.Stm { return x.AggMax(length) }
	RltStmStmAggMid      = func(x rlt.Stm, length unt.Unt) rlt.Stm { return x.AggMid(length) }
	RltStmStmAggMdn      = func(x rlt.Stm, length unt.Unt) rlt.Stm { return x.AggMdn(length) }
	RltStmStmAggSma      = func(x rlt.Stm, length unt.Unt) rlt.Stm { return x.AggSma(length) }
	RltStmStmAggGma      = func(x rlt.Stm, length unt.Unt) rlt.Stm { return x.AggGma(length) }
	RltStmStmAggWma      = func(x rlt.Stm, length unt.Unt) rlt.Stm { return x.AggWma(length) }
	RltStmStmAggRsi      = func(x rlt.Stm, length unt.Unt) rlt.Stm { return x.AggRsi(length) }
	RltStmStmAggWrsi     = func(x rlt.Stm, length unt.Unt) rlt.Stm { return x.AggWrsi(length) }
	RltStmStmAggAlma     = func(x rlt.Stm, length unt.Unt) rlt.Stm { return x.AggAlma(length) }
	RltStmStmAggVrnc     = func(x rlt.Stm, length unt.Unt) rlt.Stm { return x.AggVrnc(length) }
	RltStmStmAggStd      = func(x rlt.Stm, length unt.Unt) rlt.Stm { return x.AggStd(length) }
	RltStmStmAggRngFul   = func(x rlt.Stm, length unt.Unt) rlt.Stm { return x.AggRngFul(length) }
	RltStmStmAggRngLst   = func(x rlt.Stm, length unt.Unt) rlt.Stm { return x.AggRngLst(length) }
	RltStmStmAggProLst   = func(x rlt.Stm, length unt.Unt) rlt.Stm { return x.AggProLst(length) }
	RltStmStmAggProSma   = func(x rlt.Stm, length unt.Unt) rlt.Stm { return x.AggProSma(length) }
	RltStmStmAggProAlma  = func(x rlt.Stm, length unt.Unt) rlt.Stm { return x.AggProAlma(length) }
	RltStmStmAggEma      = func(x rlt.Stm, length unt.Unt) rlt.Stm { return x.AggEma(length) }
	RltStmStmInrAdd      = func(x rlt.Stm, off unt.Unt) rlt.Stm { return x.InrAdd(off) }
	RltStmStmInrSub      = func(x rlt.Stm, off unt.Unt) rlt.Stm { return x.InrSub(off) }
	RltStmStmInrMul      = func(x rlt.Stm, off unt.Unt) rlt.Stm { return x.InrMul(off) }
	RltStmStmInrDiv      = func(x rlt.Stm, off unt.Unt) rlt.Stm { return x.InrDiv(off) }
	RltStmStmInrRem      = func(x rlt.Stm, off unt.Unt) rlt.Stm { return x.InrRem(off) }
	RltStmStmInrPow      = func(x rlt.Stm, off unt.Unt) rlt.Stm { return x.InrPow(off) }
	RltStmStmInrMin      = func(x rlt.Stm, off unt.Unt) rlt.Stm { return x.InrMin(off) }
	RltStmStmInrMax      = func(x rlt.Stm, off unt.Unt) rlt.Stm { return x.InrMax(off) }
	RltStmStmInr1Slp     = func(x rlt.Stm, off unt.Unt) rlt.Stm { return x.InrSlp(off) }
	RltStmStmOtrAdd      = func(x rlt.Stm, off unt.Unt, a rlt.Stm) rlt.Stm { return x.OtrAdd(off, a) }
	RltStmStmOtrSub      = func(x rlt.Stm, off unt.Unt, a rlt.Stm) rlt.Stm { return x.OtrSub(off, a) }
	RltStmStmOtrMul      = func(x rlt.Stm, off unt.Unt, a rlt.Stm) rlt.Stm { return x.OtrMul(off, a) }
	RltStmStmOtrDiv      = func(x rlt.Stm, off unt.Unt, a rlt.Stm) rlt.Stm { return x.OtrDiv(off, a) }
	RltStmStmOtrRem      = func(x rlt.Stm, off unt.Unt, a rlt.Stm) rlt.Stm { return x.OtrRem(off, a) }
	RltStmStmOtrPow      = func(x rlt.Stm, off unt.Unt, a rlt.Stm) rlt.Stm { return x.OtrPow(off, a) }
	RltStmStmOtrMin      = func(x rlt.Stm, off unt.Unt, a rlt.Stm) rlt.Stm { return x.OtrMin(off, a) }
	RltStmStmOtrMax      = func(x rlt.Stm, off unt.Unt, a rlt.Stm) rlt.Stm { return x.OtrMax(off, a) }
	RltStmCndSclEql      = func(x rlt.Stm, scl flt.Flt) rlt.Cnd { return x.SclEql(scl) }
	RltStmCndSclNeq      = func(x rlt.Stm, scl flt.Flt) rlt.Cnd { return x.SclNeq(scl) }
	RltStmCndSclLss      = func(x rlt.Stm, scl flt.Flt) rlt.Cnd { return x.SclLss(scl) }
	RltStmCndSclGtr      = func(x rlt.Stm, scl flt.Flt) rlt.Cnd { return x.SclGtr(scl) }
	RltStmCndSclLeq      = func(x rlt.Stm, scl flt.Flt) rlt.Cnd { return x.SclLeq(scl) }
	RltStmCndSclGeq      = func(x rlt.Stm, scl flt.Flt) rlt.Cnd { return x.SclGeq(scl) }
	RltStmCndInrEql      = func(x rlt.Stm, off unt.Unt) rlt.Cnd { return x.InrEql(off) }
	RltStmCndInrNeq      = func(x rlt.Stm, off unt.Unt) rlt.Cnd { return x.InrNeq(off) }
	RltStmCndInrLss      = func(x rlt.Stm, off unt.Unt) rlt.Cnd { return x.InrLss(off) }
	RltStmCndInrGtr      = func(x rlt.Stm, off unt.Unt) rlt.Cnd { return x.InrGtr(off) }
	RltStmCndInrLeq      = func(x rlt.Stm, off unt.Unt) rlt.Cnd { return x.InrLeq(off) }
	RltStmCndInrGeq      = func(x rlt.Stm, off unt.Unt) rlt.Cnd { return x.InrGeq(off) }
	RltStmCndOtrEql      = func(x rlt.Stm, off unt.Unt, a rlt.Stm) rlt.Cnd { return x.OtrEql(off, a) }
	RltStmCndOtrNeq      = func(x rlt.Stm, off unt.Unt, a rlt.Stm) rlt.Cnd { return x.OtrNeq(off, a) }
	RltStmCndOtrLss      = func(x rlt.Stm, off unt.Unt, a rlt.Stm) rlt.Cnd { return x.OtrLss(off, a) }
	RltStmCndOtrGtr      = func(x rlt.Stm, off unt.Unt, a rlt.Stm) rlt.Cnd { return x.OtrGtr(off, a) }
	RltStmCndOtrLeq      = func(x rlt.Stm, off unt.Unt, a rlt.Stm) rlt.Cnd { return x.OtrLeq(off, a) }
	RltStmCndOtrGeq      = func(x rlt.Stm, off unt.Unt, a rlt.Stm) rlt.Cnd { return x.OtrGeq(off, a) }
	RltCndCndCnd1And     = func(x, a rlt.Cnd) rlt.Cnd { return x.And(a) }
	RltCndCndCnd2Seq     = func(x rlt.Cnd, dur tme.Tme, a rlt.Cnd) rlt.Cnd { return x.Seq(dur, a) }
	RltCndStgyStgy       = func(x rlt.Cnd, isLong bol.Bol, prfLim, losLim flt.Flt, durLim tme.Tme, minPnlPct flt.Flt, instr rlt.Instr, ftrStms *rlt.Stms, clss ...rlt.Cnd) rlt.Stgy {
		return x.Stgy(isLong, prfLim, losLim, durLim, minPnlPct, instr, ftrStms, clss...)
	}
	HstStmCndOtrs   = []HstStmCndOtr{HstStmCndOtrEql, HstStmCndOtrNeq, HstStmCndOtrLss, HstStmCndOtrGtr, HstStmCndOtrLeq, HstStmCndOtrGeq}
	RltInrvlSides   = []RltInrvlSide{RltInrvlSideBid, RltInrvlSideAsk}
	RltCndCndCnd2s  = []RltCndCndCnd2{RltCndCndCnd2Seq}
	HstStmStmInr1s  = []HstStmStmInr1{HstStmStmInr1Slp}
	RltStmStmSels   = []RltStmStmSel{RltStmStmSelEql, RltStmStmSelNeq, RltStmStmSelLss, RltStmStmSelGtr, RltStmStmSelLeq, RltStmStmSelGeq}
	RltStmCndInrs   = []RltStmCndInr{RltStmCndInrEql, RltStmCndInrNeq, RltStmCndInrLss, RltStmCndInrGtr, RltStmCndInrLeq, RltStmCndInrGeq}
	HstCndStgys     = []HstCndStgy{HstCndStgyStgy}
	HstCndCndCnd2s  = []HstCndCndCnd2{HstCndCndCnd2Seq}
	RltStmStmUnas   = []RltStmStmUna{RltStmStmUnaPos, RltStmStmUnaNeg, RltStmStmUnaInv, RltStmStmUnaSqr, RltStmStmUnaSqrt}
	RltStmCndScls   = []RltStmCndScl{RltStmCndSclEql, RltStmCndSclNeq, RltStmCndSclLss, RltStmCndSclGtr, RltStmCndSclLeq, RltStmCndSclGeq}
	RltCndCndCnd1s  = []RltCndCndCnd1{RltCndCndCnd1And}
	HstStmStmInrs   = []HstStmStmInr{HstStmStmInrAdd, HstStmStmInrSub, HstStmStmInrMul, HstStmStmInrDiv, HstStmStmInrRem, HstStmStmInrPow, HstStmStmInrMin, HstStmStmInrMax}
	HstSideStmRtes  = []HstSideStmRte{HstSideStmRteFst, HstSideStmRteLst, HstSideStmRteSum, HstSideStmRtePrd, HstSideStmRteMin, HstSideStmRteMax, HstSideStmRteMid, HstSideStmRteMdn, HstSideStmRteSma, HstSideStmRteGma, HstSideStmRteWma, HstSideStmRteRsi, HstSideStmRteWrsi, HstSideStmRteAlma, HstSideStmRteVrnc, HstSideStmRteStd, HstSideStmRteRngFul, HstSideStmRteRngLst, HstSideStmRteProLst, HstSideStmRteProSma, HstSideStmRteProAlma, HstSideStmRteEma}
	RltStmStmInrs   = []RltStmStmInr{RltStmStmInrAdd, RltStmStmInrSub, RltStmStmInrMul, RltStmStmInrDiv, RltStmStmInrRem, RltStmStmInrPow, RltStmStmInrMin, RltStmStmInrMax}
	RltStmStmInr1s  = []RltStmStmInr1{RltStmStmInr1Slp}
	HstPrvInstrs    = []HstPrvInstr{HstPrvInstrEurUsd, HstPrvInstrAudUsd, HstPrvInstrNzdUsd, HstPrvInstrGbpUsd}
	HstCndCndCnd1s  = []HstCndCndCnd1{HstCndCndCnd1And}
	RltPrvInstrs    = []RltPrvInstr{RltPrvInstrEurUsd, RltPrvInstrAudUsd, RltPrvInstrNzdUsd, RltPrvInstrGbpUsd}
	RltInstrInrvls  = []RltInstrInrvl{RltInstrInrvlI}
	RltCndStgys     = []RltCndStgy{RltCndStgyStgy}
	HstStmStmSels   = []HstStmStmSel{HstStmStmSelEql, HstStmStmSelNeq, HstStmStmSelLss, HstStmStmSelGtr, HstStmStmSelLeq, HstStmStmSelGeq}
	HstStmStmAggs   = []HstStmStmAgg{HstStmStmAggFst, HstStmStmAggLst, HstStmStmAggSum, HstStmStmAggPrd, HstStmStmAggMin, HstStmStmAggMax, HstStmStmAggMid, HstStmStmAggMdn, HstStmStmAggSma, HstStmStmAggGma, HstStmStmAggWma, HstStmStmAggRsi, HstStmStmAggWrsi, HstStmStmAggAlma, HstStmStmAggVrnc, HstStmStmAggStd, HstStmStmAggRngFul, HstStmStmAggRngLst, HstStmStmAggProLst, HstStmStmAggProSma, HstStmStmAggProAlma, HstStmStmAggEma}
	HstStmCndScls   = []HstStmCndScl{HstStmCndSclEql, HstStmCndSclNeq, HstStmCndSclLss, HstStmCndSclGtr, HstStmCndSclLeq, HstStmCndSclGeq}
	RltStmStmAggs   = []RltStmStmAgg{RltStmStmAggFst, RltStmStmAggLst, RltStmStmAggSum, RltStmStmAggPrd, RltStmStmAggMin, RltStmStmAggMax, RltStmStmAggMid, RltStmStmAggMdn, RltStmStmAggSma, RltStmStmAggGma, RltStmStmAggWma, RltStmStmAggRsi, RltStmStmAggWrsi, RltStmStmAggAlma, RltStmStmAggVrnc, RltStmStmAggStd, RltStmStmAggRngFul, RltStmStmAggRngLst, RltStmStmAggProLst, RltStmStmAggProSma, RltStmStmAggProAlma, RltStmStmAggEma}
	RltStmCndOtrs   = []RltStmCndOtr{RltStmCndOtrEql, RltStmCndOtrNeq, RltStmCndOtrLss, RltStmCndOtrGtr, RltStmCndOtrLeq, RltStmCndOtrGeq}
	HstStmStmUnas   = []HstStmStmUna{HstStmStmUnaPos, HstStmStmUnaNeg, HstStmStmUnaInv, HstStmStmUnaSqr, HstStmStmUnaSqrt}
	HstInrvlSides   = []HstInrvlSide{HstInrvlSideBid, HstInrvlSideAsk}
	HstSideStmRte1s = []HstSideStmRte1{HstSideStmRte1Sar}
	HstStmStmScls   = []HstStmStmScl{HstStmStmSclAdd, HstStmStmSclSub, HstStmStmSclMul, HstStmStmSclDiv, HstStmStmSclRem, HstStmStmSclPow, HstStmStmSclMin, HstStmStmSclMax}
	HstStmStmOtrs   = []HstStmStmOtr{HstStmStmOtrAdd, HstStmStmOtrSub, HstStmStmOtrMul, HstStmStmOtrDiv, HstStmStmOtrRem, HstStmStmOtrPow, HstStmStmOtrMin, HstStmStmOtrMax}
	HstStmCndInrs   = []HstStmCndInr{HstStmCndInrEql, HstStmCndInrNeq, HstStmCndInrLss, HstStmCndInrGtr, HstStmCndInrLeq, HstStmCndInrGeq}
	RltSideStmRtes  = []RltSideStmRte{RltSideStmRteFst, RltSideStmRteLst, RltSideStmRteSum, RltSideStmRtePrd, RltSideStmRteMin, RltSideStmRteMax, RltSideStmRteMid, RltSideStmRteMdn, RltSideStmRteSma, RltSideStmRteGma, RltSideStmRteWma, RltSideStmRteRsi, RltSideStmRteWrsi, RltSideStmRteAlma, RltSideStmRteVrnc, RltSideStmRteStd, RltSideStmRteRngFul, RltSideStmRteRngLst, RltSideStmRteProLst, RltSideStmRteProSma, RltSideStmRteProAlma, RltSideStmRteEma}
	RltSideStmRte1s = []RltSideStmRte1{RltSideStmRte1Sar}
	HstInstrInrvls  = []HstInstrInrvl{HstInstrInrvlI}
	RltStmStmOtrs   = []RltStmStmOtr{RltStmStmOtrAdd, RltStmStmOtrSub, RltStmStmOtrMul, RltStmStmOtrDiv, RltStmStmOtrRem, RltStmStmOtrPow, RltStmStmOtrMin, RltStmStmOtrMax}
	RltStmStmScls   = []RltStmStmScl{RltStmStmSclAdd, RltStmStmSclSub, RltStmStmSclMul, RltStmStmSclDiv, RltStmStmSclRem, RltStmStmSclPow, RltStmStmSclMin, RltStmStmSclMax}
)

type (
	HstPrvInstr    func(x hst.Prv, rng ...tme.Rng) hst.Instr
	HstInstrInrvl  func(x hst.Instr, dur tme.Tme) hst.Inrvl
	HstInrvlSide   func(x hst.Inrvl) hst.Side
	HstSideStmRte  func(x hst.Side) hst.Stm
	HstSideStmRte1 func(x hst.Side, afInc, afMax flt.Flt) hst.Stm
	HstStmStmUna   func(x hst.Stm) hst.Stm
	HstStmStmScl   func(x hst.Stm, scl flt.Flt) hst.Stm
	HstStmStmSel   func(x hst.Stm, sel flt.Flt) hst.Stm
	HstStmStmAgg   func(x hst.Stm, length unt.Unt) hst.Stm
	HstStmStmInr   func(x hst.Stm, off unt.Unt) hst.Stm
	HstStmStmInr1  func(x hst.Stm, off unt.Unt) hst.Stm
	HstStmStmOtr   func(x hst.Stm, off unt.Unt, a hst.Stm) hst.Stm
	HstStmCndScl   func(x hst.Stm, scl flt.Flt) hst.Cnd
	HstStmCndInr   func(x hst.Stm, off unt.Unt) hst.Cnd
	HstStmCndOtr   func(x hst.Stm, off unt.Unt, a hst.Stm) hst.Cnd
	HstCndCndCnd1  func(x, a hst.Cnd) hst.Cnd
	HstCndCndCnd2  func(x hst.Cnd, dur tme.Tme, a hst.Cnd) hst.Cnd
	HstCndStgy     func(x hst.Cnd, isLong bol.Bol, prfLim, losLim flt.Flt, durLim tme.Tme, minPnlPct flt.Flt, instr hst.Instr, ftrStms *hst.Stms, clss ...hst.Cnd) hst.Stgy
	RltPrvInstr    func(x rlt.Prv, rng ...tme.Rng) rlt.Instr
	RltInstrInrvl  func(x rlt.Instr, dur tme.Tme) rlt.Inrvl
	RltInrvlSide   func(x rlt.Inrvl) rlt.Side
	RltSideStmRte  func(x rlt.Side) rlt.Stm
	RltSideStmRte1 func(x rlt.Side, afInc, afMax flt.Flt) rlt.Stm
	RltStmStmUna   func(x rlt.Stm) rlt.Stm
	RltStmStmScl   func(x rlt.Stm, scl flt.Flt) rlt.Stm
	RltStmStmSel   func(x rlt.Stm, sel flt.Flt) rlt.Stm
	RltStmStmAgg   func(x rlt.Stm, length unt.Unt) rlt.Stm
	RltStmStmInr   func(x rlt.Stm, off unt.Unt) rlt.Stm
	RltStmStmInr1  func(x rlt.Stm, off unt.Unt) rlt.Stm
	RltStmStmOtr   func(x rlt.Stm, off unt.Unt, a rlt.Stm) rlt.Stm
	RltStmCndScl   func(x rlt.Stm, scl flt.Flt) rlt.Cnd
	RltStmCndInr   func(x rlt.Stm, off unt.Unt) rlt.Cnd
	RltStmCndOtr   func(x rlt.Stm, off unt.Unt, a rlt.Stm) rlt.Cnd
	RltCndCndCnd1  func(x, a rlt.Cnd) rlt.Cnd
	RltCndCndCnd2  func(x rlt.Cnd, dur tme.Tme, a rlt.Cnd) rlt.Cnd
	RltCndStgy     func(x rlt.Cnd, isLong bol.Bol, prfLim, losLim flt.Flt, durLim tme.Tme, minPnlPct flt.Flt, instr rlt.Instr, ftrStms *rlt.Stms, clss ...rlt.Cnd) rlt.Stgy
)

func StrZero(t *testing.T, a str.Str, msgs ...interface{}) {
	if !(str.Zero == a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("should equal str.Zero (actual: %v)", a))...)
	}
}
func StrNotZero(t *testing.T, a str.Str, msgs ...interface{}) {
	if !(str.Zero != a) {
		t.Helper()
		t.Fatal(append(msgs, "should not equal str.Zero")...)
	}
}
func StrEmpty(t *testing.T, a str.Str, msgs ...interface{}) {
	if !(str.Empty == a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("should equal str.Empty (actual: %v)", a))...)
	}
}
func StrNotEmpty(t *testing.T, a str.Str, msgs ...interface{}) {
	if !(str.Empty != a) {
		t.Helper()
		t.Fatal(append(msgs, "should not equal str.Empty")...)
	}
}
func BolZero(t *testing.T, a bol.Bol, msgs ...interface{}) {
	if !(bol.Zero == a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("should equal bol.Zero (actual: %v)", a))...)
	}
}
func BolNotZero(t *testing.T, a bol.Bol, msgs ...interface{}) {
	if !(bol.Zero != a) {
		t.Helper()
		t.Fatal(append(msgs, "should not equal bol.Zero")...)
	}
}
func BolFls(t *testing.T, a bol.Bol, msgs ...interface{}) {
	if !(bol.Fls == a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("should equal bol.Fls (actual: %v)", a))...)
	}
}
func BolNotFls(t *testing.T, a bol.Bol, msgs ...interface{}) {
	if !(bol.Fls != a) {
		t.Helper()
		t.Fatal(append(msgs, "should not equal bol.Fls")...)
	}
}
func BolTru(t *testing.T, a bol.Bol, msgs ...interface{}) {
	if !(bol.Tru == a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("should equal bol.Tru (actual: %v)", a))...)
	}
}
func BolNotTru(t *testing.T, a bol.Bol, msgs ...interface{}) {
	if !(bol.Tru != a) {
		t.Helper()
		t.Fatal(append(msgs, "should not equal bol.Tru")...)
	}
}
func FltZero(t *testing.T, a flt.Flt, msgs ...interface{}) {
	if !(flt.Zero == a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("should equal flt.Zero (actual: %v)", a))...)
	}
}
func FltNotZero(t *testing.T, a flt.Flt, msgs ...interface{}) {
	if !(flt.Zero != a) {
		t.Helper()
		t.Fatal(append(msgs, "should not equal flt.Zero")...)
	}
}
func FltOne(t *testing.T, a flt.Flt, msgs ...interface{}) {
	if !(flt.One == a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("should equal flt.One (actual: %v)", a))...)
	}
}
func FltNotOne(t *testing.T, a flt.Flt, msgs ...interface{}) {
	if !(flt.One != a) {
		t.Helper()
		t.Fatal(append(msgs, "should not equal flt.One")...)
	}
}
func FltNegOne(t *testing.T, a flt.Flt, msgs ...interface{}) {
	if !(flt.NegOne == a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("should equal flt.NegOne (actual: %v)", a))...)
	}
}
func FltNotNegOne(t *testing.T, a flt.Flt, msgs ...interface{}) {
	if !(flt.NegOne != a) {
		t.Helper()
		t.Fatal(append(msgs, "should not equal flt.NegOne")...)
	}
}
func FltHndrd(t *testing.T, a flt.Flt, msgs ...interface{}) {
	if !(flt.Hndrd == a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("should equal flt.Hndrd (actual: %v)", a))...)
	}
}
func FltNotHndrd(t *testing.T, a flt.Flt, msgs ...interface{}) {
	if !(flt.Hndrd != a) {
		t.Helper()
		t.Fatal(append(msgs, "should not equal flt.Hndrd")...)
	}
}
func FltMin(t *testing.T, a flt.Flt, msgs ...interface{}) {
	if !(flt.Min == a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("should equal flt.Min (actual: %v)", a))...)
	}
}
func FltNotMin(t *testing.T, a flt.Flt, msgs ...interface{}) {
	if !(flt.Min != a) {
		t.Helper()
		t.Fatal(append(msgs, "should not equal flt.Min")...)
	}
}
func FltMax(t *testing.T, a flt.Flt, msgs ...interface{}) {
	if !(flt.Max == a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("should equal flt.Max (actual: %v)", a))...)
	}
}
func FltNotMax(t *testing.T, a flt.Flt, msgs ...interface{}) {
	if !(flt.Max != a) {
		t.Helper()
		t.Fatal(append(msgs, "should not equal flt.Max")...)
	}
}
func FltTiny(t *testing.T, a flt.Flt, msgs ...interface{}) {
	if !(flt.Tiny == a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("should equal flt.Tiny (actual: %v)", a))...)
	}
}
func FltNotTiny(t *testing.T, a flt.Flt, msgs ...interface{}) {
	if !(flt.Tiny != a) {
		t.Helper()
		t.Fatal(append(msgs, "should not equal flt.Tiny")...)
	}
}
func UntZero(t *testing.T, a unt.Unt, msgs ...interface{}) {
	if !(unt.Zero == a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("should equal unt.Zero (actual: %v)", a))...)
	}
}
func UntNotZero(t *testing.T, a unt.Unt, msgs ...interface{}) {
	if !(unt.Zero != a) {
		t.Helper()
		t.Fatal(append(msgs, "should not equal unt.Zero")...)
	}
}
func UntOne(t *testing.T, a unt.Unt, msgs ...interface{}) {
	if !(unt.One == a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("should equal unt.One (actual: %v)", a))...)
	}
}
func UntNotOne(t *testing.T, a unt.Unt, msgs ...interface{}) {
	if !(unt.One != a) {
		t.Helper()
		t.Fatal(append(msgs, "should not equal unt.One")...)
	}
}
func UntMin(t *testing.T, a unt.Unt, msgs ...interface{}) {
	if !(unt.Min == a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("should equal unt.Min (actual: %v)", a))...)
	}
}
func UntNotMin(t *testing.T, a unt.Unt, msgs ...interface{}) {
	if !(unt.Min != a) {
		t.Helper()
		t.Fatal(append(msgs, "should not equal unt.Min")...)
	}
}
func UntMax(t *testing.T, a unt.Unt, msgs ...interface{}) {
	if !(unt.Max == a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("should equal unt.Max (actual: %v)", a))...)
	}
}
func UntNotMax(t *testing.T, a unt.Unt, msgs ...interface{}) {
	if !(unt.Max != a) {
		t.Helper()
		t.Fatal(append(msgs, "should not equal unt.Max")...)
	}
}
func IntZero(t *testing.T, a int.Int, msgs ...interface{}) {
	if !(int.Zero == a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("should equal int.Zero (actual: %v)", a))...)
	}
}
func IntNotZero(t *testing.T, a int.Int, msgs ...interface{}) {
	if !(int.Zero != a) {
		t.Helper()
		t.Fatal(append(msgs, "should not equal int.Zero")...)
	}
}
func IntOne(t *testing.T, a int.Int, msgs ...interface{}) {
	if !(int.One == a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("should equal int.One (actual: %v)", a))...)
	}
}
func IntNotOne(t *testing.T, a int.Int, msgs ...interface{}) {
	if !(int.One != a) {
		t.Helper()
		t.Fatal(append(msgs, "should not equal int.One")...)
	}
}
func IntNegOne(t *testing.T, a int.Int, msgs ...interface{}) {
	if !(int.NegOne == a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("should equal int.NegOne (actual: %v)", a))...)
	}
}
func IntNotNegOne(t *testing.T, a int.Int, msgs ...interface{}) {
	if !(int.NegOne != a) {
		t.Helper()
		t.Fatal(append(msgs, "should not equal int.NegOne")...)
	}
}
func IntMin(t *testing.T, a int.Int, msgs ...interface{}) {
	if !(int.Min == a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("should equal int.Min (actual: %v)", a))...)
	}
}
func IntNotMin(t *testing.T, a int.Int, msgs ...interface{}) {
	if !(int.Min != a) {
		t.Helper()
		t.Fatal(append(msgs, "should not equal int.Min")...)
	}
}
func IntMax(t *testing.T, a int.Int, msgs ...interface{}) {
	if !(int.Max == a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("should equal int.Max (actual: %v)", a))...)
	}
}
func IntNotMax(t *testing.T, a int.Int, msgs ...interface{}) {
	if !(int.Max != a) {
		t.Helper()
		t.Fatal(append(msgs, "should not equal int.Max")...)
	}
}
func TmeZero(t *testing.T, a tme.Tme, msgs ...interface{}) {
	if !(tme.Zero == a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("should equal tme.Zero (actual: %v)", a))...)
	}
}
func TmeNotZero(t *testing.T, a tme.Tme, msgs ...interface{}) {
	if !(tme.Zero != a) {
		t.Helper()
		t.Fatal(append(msgs, "should not equal tme.Zero")...)
	}
}
func TmeOne(t *testing.T, a tme.Tme, msgs ...interface{}) {
	if !(tme.One == a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("should equal tme.One (actual: %v)", a))...)
	}
}
func TmeNotOne(t *testing.T, a tme.Tme, msgs ...interface{}) {
	if !(tme.One != a) {
		t.Helper()
		t.Fatal(append(msgs, "should not equal tme.One")...)
	}
}
func TmeNegOne(t *testing.T, a tme.Tme, msgs ...interface{}) {
	if !(tme.NegOne == a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("should equal tme.NegOne (actual: %v)", a))...)
	}
}
func TmeNotNegOne(t *testing.T, a tme.Tme, msgs ...interface{}) {
	if !(tme.NegOne != a) {
		t.Helper()
		t.Fatal(append(msgs, "should not equal tme.NegOne")...)
	}
}
func TmeMin(t *testing.T, a tme.Tme, msgs ...interface{}) {
	if !(tme.Min == a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("should equal tme.Min (actual: %v)", a))...)
	}
}
func TmeNotMin(t *testing.T, a tme.Tme, msgs ...interface{}) {
	if !(tme.Min != a) {
		t.Helper()
		t.Fatal(append(msgs, "should not equal tme.Min")...)
	}
}
func TmeMax(t *testing.T, a tme.Tme, msgs ...interface{}) {
	if !(tme.Max == a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("should equal tme.Max (actual: %v)", a))...)
	}
}
func TmeNotMax(t *testing.T, a tme.Tme, msgs ...interface{}) {
	if !(tme.Max != a) {
		t.Helper()
		t.Fatal(append(msgs, "should not equal tme.Max")...)
	}
}
func TmeSecond(t *testing.T, a tme.Tme, msgs ...interface{}) {
	if !(tme.Second == a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("should equal tme.Second (actual: %v)", a))...)
	}
}
func TmeNotSecond(t *testing.T, a tme.Tme, msgs ...interface{}) {
	if !(tme.Second != a) {
		t.Helper()
		t.Fatal(append(msgs, "should not equal tme.Second")...)
	}
}
func TmeMinute(t *testing.T, a tme.Tme, msgs ...interface{}) {
	if !(tme.Minute == a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("should equal tme.Minute (actual: %v)", a))...)
	}
}
func TmeNotMinute(t *testing.T, a tme.Tme, msgs ...interface{}) {
	if !(tme.Minute != a) {
		t.Helper()
		t.Fatal(append(msgs, "should not equal tme.Minute")...)
	}
}
func TmeHour(t *testing.T, a tme.Tme, msgs ...interface{}) {
	if !(tme.Hour == a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("should equal tme.Hour (actual: %v)", a))...)
	}
}
func TmeNotHour(t *testing.T, a tme.Tme, msgs ...interface{}) {
	if !(tme.Hour != a) {
		t.Helper()
		t.Fatal(append(msgs, "should not equal tme.Hour")...)
	}
}
func TmeDay(t *testing.T, a tme.Tme, msgs ...interface{}) {
	if !(tme.Day == a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("should equal tme.Day (actual: %v)", a))...)
	}
}
func TmeNotDay(t *testing.T, a tme.Tme, msgs ...interface{}) {
	if !(tme.Day != a) {
		t.Helper()
		t.Fatal(append(msgs, "should not equal tme.Day")...)
	}
}
func TmeWeek(t *testing.T, a tme.Tme, msgs ...interface{}) {
	if !(tme.Week == a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("should equal tme.Week (actual: %v)", a))...)
	}
}
func TmeNotWeek(t *testing.T, a tme.Tme, msgs ...interface{}) {
	if !(tme.Week != a) {
		t.Helper()
		t.Fatal(append(msgs, "should not equal tme.Week")...)
	}
}
func TmeS1(t *testing.T, a tme.Tme, msgs ...interface{}) {
	if !(tme.S1 == a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("should equal tme.S1 (actual: %v)", a))...)
	}
}
func TmeNotS1(t *testing.T, a tme.Tme, msgs ...interface{}) {
	if !(tme.S1 != a) {
		t.Helper()
		t.Fatal(append(msgs, "should not equal tme.S1")...)
	}
}
func TmeS5(t *testing.T, a tme.Tme, msgs ...interface{}) {
	if !(tme.S5 == a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("should equal tme.S5 (actual: %v)", a))...)
	}
}
func TmeNotS5(t *testing.T, a tme.Tme, msgs ...interface{}) {
	if !(tme.S5 != a) {
		t.Helper()
		t.Fatal(append(msgs, "should not equal tme.S5")...)
	}
}
func TmeS10(t *testing.T, a tme.Tme, msgs ...interface{}) {
	if !(tme.S10 == a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("should equal tme.S10 (actual: %v)", a))...)
	}
}
func TmeNotS10(t *testing.T, a tme.Tme, msgs ...interface{}) {
	if !(tme.S10 != a) {
		t.Helper()
		t.Fatal(append(msgs, "should not equal tme.S10")...)
	}
}
func TmeS15(t *testing.T, a tme.Tme, msgs ...interface{}) {
	if !(tme.S15 == a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("should equal tme.S15 (actual: %v)", a))...)
	}
}
func TmeNotS15(t *testing.T, a tme.Tme, msgs ...interface{}) {
	if !(tme.S15 != a) {
		t.Helper()
		t.Fatal(append(msgs, "should not equal tme.S15")...)
	}
}
func TmeS20(t *testing.T, a tme.Tme, msgs ...interface{}) {
	if !(tme.S20 == a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("should equal tme.S20 (actual: %v)", a))...)
	}
}
func TmeNotS20(t *testing.T, a tme.Tme, msgs ...interface{}) {
	if !(tme.S20 != a) {
		t.Helper()
		t.Fatal(append(msgs, "should not equal tme.S20")...)
	}
}
func TmeS30(t *testing.T, a tme.Tme, msgs ...interface{}) {
	if !(tme.S30 == a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("should equal tme.S30 (actual: %v)", a))...)
	}
}
func TmeNotS30(t *testing.T, a tme.Tme, msgs ...interface{}) {
	if !(tme.S30 != a) {
		t.Helper()
		t.Fatal(append(msgs, "should not equal tme.S30")...)
	}
}
func TmeS40(t *testing.T, a tme.Tme, msgs ...interface{}) {
	if !(tme.S40 == a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("should equal tme.S40 (actual: %v)", a))...)
	}
}
func TmeNotS40(t *testing.T, a tme.Tme, msgs ...interface{}) {
	if !(tme.S40 != a) {
		t.Helper()
		t.Fatal(append(msgs, "should not equal tme.S40")...)
	}
}
func TmeS50(t *testing.T, a tme.Tme, msgs ...interface{}) {
	if !(tme.S50 == a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("should equal tme.S50 (actual: %v)", a))...)
	}
}
func TmeNotS50(t *testing.T, a tme.Tme, msgs ...interface{}) {
	if !(tme.S50 != a) {
		t.Helper()
		t.Fatal(append(msgs, "should not equal tme.S50")...)
	}
}
func TmeM1(t *testing.T, a tme.Tme, msgs ...interface{}) {
	if !(tme.M1 == a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("should equal tme.M1 (actual: %v)", a))...)
	}
}
func TmeNotM1(t *testing.T, a tme.Tme, msgs ...interface{}) {
	if !(tme.M1 != a) {
		t.Helper()
		t.Fatal(append(msgs, "should not equal tme.M1")...)
	}
}
func TmeM5(t *testing.T, a tme.Tme, msgs ...interface{}) {
	if !(tme.M5 == a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("should equal tme.M5 (actual: %v)", a))...)
	}
}
func TmeNotM5(t *testing.T, a tme.Tme, msgs ...interface{}) {
	if !(tme.M5 != a) {
		t.Helper()
		t.Fatal(append(msgs, "should not equal tme.M5")...)
	}
}
func TmeM10(t *testing.T, a tme.Tme, msgs ...interface{}) {
	if !(tme.M10 == a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("should equal tme.M10 (actual: %v)", a))...)
	}
}
func TmeNotM10(t *testing.T, a tme.Tme, msgs ...interface{}) {
	if !(tme.M10 != a) {
		t.Helper()
		t.Fatal(append(msgs, "should not equal tme.M10")...)
	}
}
func TmeM15(t *testing.T, a tme.Tme, msgs ...interface{}) {
	if !(tme.M15 == a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("should equal tme.M15 (actual: %v)", a))...)
	}
}
func TmeNotM15(t *testing.T, a tme.Tme, msgs ...interface{}) {
	if !(tme.M15 != a) {
		t.Helper()
		t.Fatal(append(msgs, "should not equal tme.M15")...)
	}
}
func TmeM20(t *testing.T, a tme.Tme, msgs ...interface{}) {
	if !(tme.M20 == a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("should equal tme.M20 (actual: %v)", a))...)
	}
}
func TmeNotM20(t *testing.T, a tme.Tme, msgs ...interface{}) {
	if !(tme.M20 != a) {
		t.Helper()
		t.Fatal(append(msgs, "should not equal tme.M20")...)
	}
}
func TmeM30(t *testing.T, a tme.Tme, msgs ...interface{}) {
	if !(tme.M30 == a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("should equal tme.M30 (actual: %v)", a))...)
	}
}
func TmeNotM30(t *testing.T, a tme.Tme, msgs ...interface{}) {
	if !(tme.M30 != a) {
		t.Helper()
		t.Fatal(append(msgs, "should not equal tme.M30")...)
	}
}
func TmeM40(t *testing.T, a tme.Tme, msgs ...interface{}) {
	if !(tme.M40 == a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("should equal tme.M40 (actual: %v)", a))...)
	}
}
func TmeNotM40(t *testing.T, a tme.Tme, msgs ...interface{}) {
	if !(tme.M40 != a) {
		t.Helper()
		t.Fatal(append(msgs, "should not equal tme.M40")...)
	}
}
func TmeM50(t *testing.T, a tme.Tme, msgs ...interface{}) {
	if !(tme.M50 == a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("should equal tme.M50 (actual: %v)", a))...)
	}
}
func TmeNotM50(t *testing.T, a tme.Tme, msgs ...interface{}) {
	if !(tme.M50 != a) {
		t.Helper()
		t.Fatal(append(msgs, "should not equal tme.M50")...)
	}
}
func TmeH1(t *testing.T, a tme.Tme, msgs ...interface{}) {
	if !(tme.H1 == a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("should equal tme.H1 (actual: %v)", a))...)
	}
}
func TmeNotH1(t *testing.T, a tme.Tme, msgs ...interface{}) {
	if !(tme.H1 != a) {
		t.Helper()
		t.Fatal(append(msgs, "should not equal tme.H1")...)
	}
}
func TmeD1(t *testing.T, a tme.Tme, msgs ...interface{}) {
	if !(tme.D1 == a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("should equal tme.D1 (actual: %v)", a))...)
	}
}
func TmeNotD1(t *testing.T, a tme.Tme, msgs ...interface{}) {
	if !(tme.D1 != a) {
		t.Helper()
		t.Fatal(append(msgs, "should not equal tme.D1")...)
	}
}
func TmeResolution(t *testing.T, a tme.Tme, msgs ...interface{}) {
	if !(tme.Resolution == a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("should equal tme.Resolution (actual: %v)", a))...)
	}
}
func TmeNotResolution(t *testing.T, a tme.Tme, msgs ...interface{}) {
	if !(tme.Resolution != a) {
		t.Helper()
		t.Fatal(append(msgs, "should not equal tme.Resolution")...)
	}
}
func StringEql(t *testing.T, e, a string, msgs ...interface{}) {
	if e != a {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("should be equal (expected:%v actual:%v)", e, a))...)
	}
}
func StringSliceEql(t *testing.T, e, a []string, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		StringEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm String (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func BoolEql(t *testing.T, e, a bool, msgs ...interface{}) {
	if e != a {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("should be equal (expected:%v actual:%v)", e, a))...)
	}
}
func BoolSliceEql(t *testing.T, e, a []bool, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		BoolEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm Bool (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func ByteEql(t *testing.T, e, a byte, msgs ...interface{}) {
	if e != a {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("should be equal (expected:%v actual:%v)", e, a))...)
	}
}
func ByteSliceEql(t *testing.T, e, a []byte, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		ByteEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm Byte (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func Uint32Eql(t *testing.T, e, a uint32, msgs ...interface{}) {
	if e != a {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("should be equal (expected:%v actual:%v)", e, a))...)
	}
}
func Uint32SliceEql(t *testing.T, e, a []uint32, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		Uint32Eql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm Uint32 (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func Uint64Eql(t *testing.T, e, a uint64, msgs ...interface{}) {
	if e != a {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("should be equal (expected:%v actual:%v)", e, a))...)
	}
}
func Uint64SliceEql(t *testing.T, e, a []uint64, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		Uint64Eql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm Uint64 (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func Int32Eql(t *testing.T, e, a int32, msgs ...interface{}) {
	if e != a {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("should be equal (expected:%v actual:%v)", e, a))...)
	}
}
func Int32SliceEql(t *testing.T, e, a []int32, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		Int32Eql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm Int32 (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func Int64Eql(t *testing.T, e, a int64, msgs ...interface{}) {
	if e != a {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("should be equal (expected:%v actual:%v)", e, a))...)
	}
}
func Int64SliceEql(t *testing.T, e, a []int64, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		Int64Eql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm Int64 (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func Float32Eql(t *testing.T, e, a float32, msgs ...interface{}) {
	if e != a {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("should be equal (expected:%v actual:%v)", e, a))...)
	}
}
func Float32SliceEql(t *testing.T, e, a []float32, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		Float32Eql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm Float32 (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func Float64Eql(t *testing.T, e, a float64, msgs ...interface{}) {
	if e != a {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("should be equal (expected:%v actual:%v)", e, a))...)
	}
}
func Float64SliceEql(t *testing.T, e, a []float64, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		Float64Eql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm Float64 (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RuneEql(t *testing.T, e, a rune, msgs ...interface{}) {
	if e != a {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("should be equal (expected:%v actual:%v)", e, a))...)
	}
}
func RuneSliceEql(t *testing.T, e, a []rune, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		RuneEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm Rune (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func TimeEql(t *testing.T, e, a time.Time, msgs ...interface{}) {
	if e != a {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("should be equal (expected:%v actual:%v)", e, a))...)
	}
}
func TimeSliceEql(t *testing.T, e, a []time.Time, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		TimeEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm Time (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func TimeDurationEql(t *testing.T, e, a time.Duration, msgs ...interface{}) {
	if e != a {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("should be equal (expected:%v actual:%v)", e, a))...)
	}
}
func TimeDurationSliceEql(t *testing.T, e, a []time.Duration, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		TimeDurationEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm Duration (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func AnaOanEql(t *testing.T, e, a *ana.Oan, msgs ...interface{}) {
	if e != a {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("should be equal (expected:%v actual:%v)", e, a))...)
	}
}
func AnaOanSliceEql(t *testing.T, e, a []*ana.Oan, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		AnaOanEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm Oan (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func StrEql(t *testing.T, e, a str.Str, msgs ...interface{}) {
	if e != a {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("should be equal (expected:%v actual:%v)", e, a))...)
	}
}
func StrLss(t *testing.T, e, a str.Str, msgs ...interface{}) {
	if !(e < a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("should be Lss (expected:%v actual:%v)", e, a))...)
	}
}
func StrGtr(t *testing.T, e, a str.Str, msgs ...interface{}) {
	if !(e > a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("should be Gtr (expected:%v actual:%v)", e, a))...)
	}
}
func StrLeq(t *testing.T, e, a str.Str, msgs ...interface{}) {
	if !(e <= a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("should be Leq (expected:%v actual:%v)", e, a))...)
	}
}
func StrGeq(t *testing.T, e, a str.Str, msgs ...interface{}) {
	if !(e >= a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("should be Geq (expected:%v actual:%v)", e, a))...)
	}
}
func StrSliceEql(t *testing.T, e, a []str.Str, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		StrEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm Str (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func BolEql(t *testing.T, e, a bol.Bol, msgs ...interface{}) {
	if e != a {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("should be equal (expected:%v actual:%v)", e, a))...)
	}
}
func BolSliceEql(t *testing.T, e, a []bol.Bol, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		BolEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm Bol (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func FltEql(t *testing.T, e, a flt.Flt, msgs ...interface{}) {
	if e != a && !(e.IsNaN() && a.IsNaN()) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("should be equal (expected:%v actual:%v)", e, a))...)
	}
}
func FltLss(t *testing.T, e, a flt.Flt, msgs ...interface{}) {
	if !(e < a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("should be Lss (expected:%v actual:%v)", e, a))...)
	}
}
func FltGtr(t *testing.T, e, a flt.Flt, msgs ...interface{}) {
	if !(e > a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("should be Gtr (expected:%v actual:%v)", e, a))...)
	}
}
func FltLeq(t *testing.T, e, a flt.Flt, msgs ...interface{}) {
	if !(e <= a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("should be Leq (expected:%v actual:%v)", e, a))...)
	}
}
func FltGeq(t *testing.T, e, a flt.Flt, msgs ...interface{}) {
	if !(e >= a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("should be Geq (expected:%v actual:%v)", e, a))...)
	}
}
func FltSliceEql(t *testing.T, e, a []flt.Flt, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		FltEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm Flt (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func UntEql(t *testing.T, e, a unt.Unt, msgs ...interface{}) {
	if e != a {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("should be equal (expected:%v actual:%v)", e, a))...)
	}
}
func UntLss(t *testing.T, e, a unt.Unt, msgs ...interface{}) {
	if !(e < a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("should be Lss (expected:%v actual:%v)", e, a))...)
	}
}
func UntGtr(t *testing.T, e, a unt.Unt, msgs ...interface{}) {
	if !(e > a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("should be Gtr (expected:%v actual:%v)", e, a))...)
	}
}
func UntLeq(t *testing.T, e, a unt.Unt, msgs ...interface{}) {
	if !(e <= a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("should be Leq (expected:%v actual:%v)", e, a))...)
	}
}
func UntGeq(t *testing.T, e, a unt.Unt, msgs ...interface{}) {
	if !(e >= a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("should be Geq (expected:%v actual:%v)", e, a))...)
	}
}
func UntSliceEql(t *testing.T, e, a []unt.Unt, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		UntEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm Unt (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func IntEql(t *testing.T, e, a int.Int, msgs ...interface{}) {
	if e != a {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("should be equal (expected:%v actual:%v)", e, a))...)
	}
}
func IntLss(t *testing.T, e, a int.Int, msgs ...interface{}) {
	if !(e < a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("should be Lss (expected:%v actual:%v)", e, a))...)
	}
}
func IntGtr(t *testing.T, e, a int.Int, msgs ...interface{}) {
	if !(e > a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("should be Gtr (expected:%v actual:%v)", e, a))...)
	}
}
func IntLeq(t *testing.T, e, a int.Int, msgs ...interface{}) {
	if !(e <= a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("should be Leq (expected:%v actual:%v)", e, a))...)
	}
}
func IntGeq(t *testing.T, e, a int.Int, msgs ...interface{}) {
	if !(e >= a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("should be Geq (expected:%v actual:%v)", e, a))...)
	}
}
func IntSliceEql(t *testing.T, e, a []int.Int, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		IntEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm Int (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func TmeEql(t *testing.T, e, a tme.Tme, msgs ...interface{}) {
	if e != a {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("should be equal (expected:%v actual:%v)", e, a))...)
	}
}
func TmeLss(t *testing.T, e, a tme.Tme, msgs ...interface{}) {
	if !(e < a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("should be Lss (expected:%v actual:%v)", e, a))...)
	}
}
func TmeGtr(t *testing.T, e, a tme.Tme, msgs ...interface{}) {
	if !(e > a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("should be Gtr (expected:%v actual:%v)", e, a))...)
	}
}
func TmeLeq(t *testing.T, e, a tme.Tme, msgs ...interface{}) {
	if !(e <= a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("should be Leq (expected:%v actual:%v)", e, a))...)
	}
}
func TmeGeq(t *testing.T, e, a tme.Tme, msgs ...interface{}) {
	if !(e >= a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("should be Geq (expected:%v actual:%v)", e, a))...)
	}
}
func TmeSliceEql(t *testing.T, e, a []tme.Tme, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		TmeEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm Tme (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func BndEql(t *testing.T, e, a bnd.Bnd, msgs ...interface{}) {
	UntEql(t, e.Lim, a.Lim, append(msgs, "Bnd.Lim"))
}
func BndNotZero(t *testing.T, a bnd.Bnd, msgs ...interface{}) {
	UntNotZero(t, a.Lim, append(msgs, "Bnd.Lim"))
}
func BndSliceEql(t *testing.T, e, a []bnd.Bnd, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		BndEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm Bnd (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func FltRngEql(t *testing.T, e, a flt.Rng, msgs ...interface{}) {
	FltEql(t, e.Min, a.Min, append(msgs, "Rng.Min"))
	FltEql(t, e.Max, a.Max, append(msgs, "Rng.Max"))
}
func FltRngNotZero(t *testing.T, a flt.Rng, msgs ...interface{}) {
	FltNotZero(t, a.Min, append(msgs, "Rng.Min"))
	FltNotZero(t, a.Max, append(msgs, "Rng.Max"))
}
func FltRngSliceEql(t *testing.T, e, a []flt.Rng, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		FltRngEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm Rng (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func TmeRngEql(t *testing.T, e, a tme.Rng, msgs ...interface{}) {
	TmeEql(t, e.Min, a.Min, append(msgs, "Rng.Min"))
	TmeEql(t, e.Max, a.Max, append(msgs, "Rng.Max"))
}
func TmeRngNotZero(t *testing.T, a tme.Rng, msgs ...interface{}) {
	TmeNotZero(t, a.Min, append(msgs, "Rng.Min"))
	TmeNotZero(t, a.Max, append(msgs, "Rng.Max"))
}
func TmeRngSliceEql(t *testing.T, e, a []tme.Rng, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		TmeRngEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm Rng (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func StrsEql(t *testing.T, e, a *strs.Strs, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	if len(*e) != len(*a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(*e), len(*a)))...)
	}
	for n := 0; n < len(*e); n++ {
		StrEql(t, (*e)[n], (*a)[n], append(msgs, fmt.Sprintf("elm Str (idx:%v expected:%v actual:%v)", n, (*e)[n], (*a)[n])))
	}
}
func StrsNotZero(t *testing.T, a *strs.Strs, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	if a.Cnt() == 0 {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("cnt is zero"))...)
	}
}
func StrsSliceEql(t *testing.T, e, a []*strs.Strs, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		StrsEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm Strs (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func BolsEql(t *testing.T, e, a *bols.Bols, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	if len(*e) != len(*a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(*e), len(*a)))...)
	}
	for n := 0; n < len(*e); n++ {
		BolEql(t, (*e)[n], (*a)[n], append(msgs, fmt.Sprintf("elm Bol (idx:%v expected:%v actual:%v)", n, (*e)[n], (*a)[n])))
	}
}
func BolsNotZero(t *testing.T, a *bols.Bols, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	if a.Cnt() == 0 {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("cnt is zero"))...)
	}
}
func BolsSliceEql(t *testing.T, e, a []*bols.Bols, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		BolsEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm Bols (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func FltsEql(t *testing.T, e, a *flts.Flts, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	if len(*e) != len(*a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(*e), len(*a)))...)
	}
	for n := 0; n < len(*e); n++ {
		FltEql(t, (*e)[n], (*a)[n], append(msgs, fmt.Sprintf("elm Flt (idx:%v expected:%v actual:%v)", n, (*e)[n], (*a)[n])))
	}
}
func FltsNotZero(t *testing.T, a *flts.Flts, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	if a.Cnt() == 0 {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("cnt is zero"))...)
	}
}
func FltsSliceEql(t *testing.T, e, a []*flts.Flts, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		FltsEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm Flts (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func UntsEql(t *testing.T, e, a *unts.Unts, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	if len(*e) != len(*a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(*e), len(*a)))...)
	}
	for n := 0; n < len(*e); n++ {
		UntEql(t, (*e)[n], (*a)[n], append(msgs, fmt.Sprintf("elm Unt (idx:%v expected:%v actual:%v)", n, (*e)[n], (*a)[n])))
	}
}
func UntsNotZero(t *testing.T, a *unts.Unts, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	if a.Cnt() == 0 {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("cnt is zero"))...)
	}
}
func UntsSliceEql(t *testing.T, e, a []*unts.Unts, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		UntsEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm Unts (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func IntsEql(t *testing.T, e, a *ints.Ints, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	if len(*e) != len(*a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(*e), len(*a)))...)
	}
	for n := 0; n < len(*e); n++ {
		IntEql(t, (*e)[n], (*a)[n], append(msgs, fmt.Sprintf("elm Int (idx:%v expected:%v actual:%v)", n, (*e)[n], (*a)[n])))
	}
}
func IntsNotZero(t *testing.T, a *ints.Ints, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	if a.Cnt() == 0 {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("cnt is zero"))...)
	}
}
func IntsSliceEql(t *testing.T, e, a []*ints.Ints, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		IntsEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm Ints (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func TmesEql(t *testing.T, e, a *tmes.Tmes, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	if len(*e) != len(*a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(*e), len(*a)))...)
	}
	for n := 0; n < len(*e); n++ {
		TmeEql(t, (*e)[n], (*a)[n], append(msgs, fmt.Sprintf("elm Tme (idx:%v expected:%v actual:%v)", n, (*e)[n], (*a)[n])))
	}
}
func TmesNotZero(t *testing.T, a *tmes.Tmes, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	if a.Cnt() == 0 {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("cnt is zero"))...)
	}
}
func TmesSliceEql(t *testing.T, e, a []*tmes.Tmes, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		TmesEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm Tmes (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func BndsEql(t *testing.T, e, a *bnds.Bnds, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	if len(*e) != len(*a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(*e), len(*a)))...)
	}
	for n := 0; n < len(*e); n++ {
		BndEql(t, (*e)[n], (*a)[n], append(msgs, fmt.Sprintf("elm Bnd (idx:%v expected:%v actual:%v)", n, (*e)[n], (*a)[n])))
	}
}
func BndsNotZero(t *testing.T, a *bnds.Bnds, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	if a.Cnt() == 0 {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("cnt is zero"))...)
	}
}
func BndsSliceEql(t *testing.T, e, a []*bnds.Bnds, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		BndsEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm Bnds (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func TmeRngsEql(t *testing.T, e, a *tme.Rngs, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	if len(*e) != len(*a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(*e), len(*a)))...)
	}
	for n := 0; n < len(*e); n++ {
		TmeRngEql(t, (*e)[n], (*a)[n], append(msgs, fmt.Sprintf("elm Rng (idx:%v expected:%v actual:%v)", n, (*e)[n], (*a)[n])))
	}
}
func TmeRngsNotZero(t *testing.T, a *tme.Rngs, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	if a.Cnt() == 0 {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("cnt is zero"))...)
	}
}
func TmeRngsSliceEql(t *testing.T, e, a []*tme.Rngs, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		TmeRngsEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm Rngs (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func AnaTmeIdxEql(t *testing.T, e, a ana.TmeIdx, msgs ...interface{}) {
	TmeEql(t, e.Tme, a.Tme, append(msgs, "TmeIdx.Tme"))
	UntEql(t, e.Idx, a.Idx, append(msgs, "TmeIdx.Idx"))
}
func AnaTmeIdxNotZero(t *testing.T, a ana.TmeIdx, msgs ...interface{}) {
	TmeNotZero(t, a.Tme, append(msgs, "TmeIdx.Tme"))
	UntNotZero(t, a.Idx, append(msgs, "TmeIdx.Idx"))
}
func AnaTmeIdxSliceEql(t *testing.T, e, a []ana.TmeIdx, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		AnaTmeIdxEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm TmeIdx (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func AnaTmeIdxsEql(t *testing.T, e, a *ana.TmeIdxs, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	if len(*e) != len(*a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(*e), len(*a)))...)
	}
	for n := 0; n < len(*e); n++ {
		AnaTmeIdxEql(t, (*e)[n], (*a)[n], append(msgs, fmt.Sprintf("elm TmeIdx (idx:%v expected:%v actual:%v)", n, (*e)[n], (*a)[n])))
	}
}
func AnaTmeIdxsNotZero(t *testing.T, a *ana.TmeIdxs, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	if a.Cnt() == 0 {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("cnt is zero"))...)
	}
}
func AnaTmeIdxsSliceEql(t *testing.T, e, a []*ana.TmeIdxs, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		AnaTmeIdxsEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm TmeIdxs (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func AnaTmeFltEql(t *testing.T, e, a ana.TmeFlt, msgs ...interface{}) {
	TmeEql(t, e.Tme, a.Tme, append(msgs, "TmeFlt.Tme"))
	FltEql(t, e.Flt, a.Flt, append(msgs, "TmeFlt.Flt"))
}
func AnaTmeFltNotZero(t *testing.T, a ana.TmeFlt, msgs ...interface{}) {
	TmeNotZero(t, a.Tme, append(msgs, "TmeFlt.Tme"))
	FltNotZero(t, a.Flt, append(msgs, "TmeFlt.Flt"))
}
func AnaTmeFltSliceEql(t *testing.T, e, a []ana.TmeFlt, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		AnaTmeFltEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm TmeFlt (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func AnaTmeFltsEql(t *testing.T, e, a ana.TmeFlts, msgs ...interface{}) {
	TmeEql(t, e.Tme, a.Tme, append(msgs, "TmeFlts.Tme"))
	FltsEql(t, e.Flts, a.Flts, append(msgs, "TmeFlts.Flts"))
}
func AnaTmeFltsNotZero(t *testing.T, a ana.TmeFlts, msgs ...interface{}) {
	TmeNotZero(t, a.Tme, append(msgs, "TmeFlts.Tme"))
	FltsNotZero(t, a.Flts, append(msgs, "TmeFlts.Flts"))
}
func AnaTmeFltsSliceEql(t *testing.T, e, a []ana.TmeFlts, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		AnaTmeFltsEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm TmeFlts (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func AnaStmEql(t *testing.T, e, a *ana.Stm, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "Stm.Tmes"))
	FltsEql(t, e.Bids, a.Bids, append(msgs, "Stm.Bids"))
	FltsEql(t, e.Asks, a.Asks, append(msgs, "Stm.Asks"))
	UntsEql(t, e.BidLims, a.BidLims, append(msgs, "Stm.BidLims"))
	UntsEql(t, e.AskLims, a.AskLims, append(msgs, "Stm.AskLims"))
}
func AnaStmNotZero(t *testing.T, a *ana.Stm, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	TmesNotZero(t, a.Tmes, append(msgs, "Stm.Tmes"))
	FltsNotZero(t, a.Bids, append(msgs, "Stm.Bids"))
	FltsNotZero(t, a.Asks, append(msgs, "Stm.Asks"))
	UntsNotZero(t, a.BidLims, append(msgs, "Stm.BidLims"))
	UntsNotZero(t, a.AskLims, append(msgs, "Stm.AskLims"))
}
func AnaStmSliceEql(t *testing.T, e, a []*ana.Stm, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		AnaStmEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm Stm (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func AnaInstrEql(t *testing.T, e, a *ana.Instr, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	StrEql(t, e.Name, a.Name, append(msgs, "Instr.Name"))
	FltEql(t, e.Pip, a.Pip, append(msgs, "Instr.Pip"))
	FltEql(t, e.MrgnRtio, a.MrgnRtio, append(msgs, "Instr.MrgnRtio"))
	FltEql(t, e.SpdMin, a.SpdMin, append(msgs, "Instr.SpdMin"))
	FltEql(t, e.SpdMax, a.SpdMax, append(msgs, "Instr.SpdMax"))
	FltEql(t, e.SpdMdn, a.SpdMdn, append(msgs, "Instr.SpdMdn"))
	FltEql(t, e.SpdAvg, a.SpdAvg, append(msgs, "Instr.SpdAvg"))
	FltEql(t, e.SpdStd, a.SpdStd, append(msgs, "Instr.SpdStd"))
	FltEql(t, e.SpdOpnLim, a.SpdOpnLim, append(msgs, "Instr.SpdOpnLim"))
	TmeEql(t, e.Fst, a.Fst, append(msgs, "Instr.Fst"))
	TmeEql(t, e.Lst, a.Lst, append(msgs, "Instr.Lst"))
	UntEql(t, e.TmeCnt, a.TmeCnt, append(msgs, "Instr.TmeCnt"))
	UntEql(t, e.DayCnt, a.DayCnt, append(msgs, "Instr.DayCnt"))
	UntEql(t, e.DisplayPrecision, a.DisplayPrecision, append(msgs, "Instr.DisplayPrecision"))
	UntEql(t, e.TradeUnitsPrecision, a.TradeUnitsPrecision, append(msgs, "Instr.TradeUnitsPrecision"))
	UntEql(t, e.MinTrdSize, a.MinTrdSize, append(msgs, "Instr.MinTrdSize"))
	FltEql(t, e.MaxTrailingStopDistance, a.MaxTrailingStopDistance, append(msgs, "Instr.MaxTrailingStopDistance"))
	FltEql(t, e.MinTrailingStopDistance, a.MinTrailingStopDistance, append(msgs, "Instr.MinTrailingStopDistance"))
	UntEql(t, e.MaxPositionSize, a.MaxPositionSize, append(msgs, "Instr.MaxPositionSize"))
	UntEql(t, e.MaxOrderUnits, a.MaxOrderUnits, append(msgs, "Instr.MaxOrderUnits"))
	StrEql(t, e.Typ, a.Typ, append(msgs, "Instr.Typ"))
}
func AnaInstrNotZero(t *testing.T, a *ana.Instr, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	StrNotZero(t, a.Name, append(msgs, "Instr.Name"))
	FltNotZero(t, a.Pip, append(msgs, "Instr.Pip"))
	FltNotZero(t, a.MrgnRtio, append(msgs, "Instr.MrgnRtio"))
	FltNotZero(t, a.SpdMin, append(msgs, "Instr.SpdMin"))
	FltNotZero(t, a.SpdMax, append(msgs, "Instr.SpdMax"))
	FltNotZero(t, a.SpdMdn, append(msgs, "Instr.SpdMdn"))
	FltNotZero(t, a.SpdAvg, append(msgs, "Instr.SpdAvg"))
	FltNotZero(t, a.SpdStd, append(msgs, "Instr.SpdStd"))
	FltNotZero(t, a.SpdOpnLim, append(msgs, "Instr.SpdOpnLim"))
	TmeNotZero(t, a.Fst, append(msgs, "Instr.Fst"))
	TmeNotZero(t, a.Lst, append(msgs, "Instr.Lst"))
	UntNotZero(t, a.TmeCnt, append(msgs, "Instr.TmeCnt"))
	UntNotZero(t, a.DayCnt, append(msgs, "Instr.DayCnt"))
	UntNotZero(t, a.DisplayPrecision, append(msgs, "Instr.DisplayPrecision"))
	UntNotZero(t, a.TradeUnitsPrecision, append(msgs, "Instr.TradeUnitsPrecision"))
	UntNotZero(t, a.MinTrdSize, append(msgs, "Instr.MinTrdSize"))
	FltNotZero(t, a.MaxTrailingStopDistance, append(msgs, "Instr.MaxTrailingStopDistance"))
	FltNotZero(t, a.MinTrailingStopDistance, append(msgs, "Instr.MinTrailingStopDistance"))
	UntNotZero(t, a.MaxPositionSize, append(msgs, "Instr.MaxPositionSize"))
	UntNotZero(t, a.MaxOrderUnits, append(msgs, "Instr.MaxOrderUnits"))
	StrNotZero(t, a.Typ, append(msgs, "Instr.Typ"))
}
func AnaInstrSliceEql(t *testing.T, e, a []*ana.Instr, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		AnaInstrEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm Instr (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func AnaTrdEql(t *testing.T, e, a *ana.Trd, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	TmeEql(t, e.OpnTme, a.OpnTme, append(msgs, "Trd.OpnTme"))
	TmeEql(t, e.ClsTme, a.ClsTme, append(msgs, "Trd.ClsTme"))
	FltEql(t, e.OpnBid, a.OpnBid, append(msgs, "Trd.OpnBid"))
	FltEql(t, e.ClsBid, a.ClsBid, append(msgs, "Trd.ClsBid"))
	FltEql(t, e.OpnAsk, a.OpnAsk, append(msgs, "Trd.OpnAsk"))
	FltEql(t, e.ClsAsk, a.ClsAsk, append(msgs, "Trd.ClsAsk"))
	FltEql(t, e.OpnSpd, a.OpnSpd, append(msgs, "Trd.OpnSpd"))
	FltEql(t, e.ClsSpd, a.ClsSpd, append(msgs, "Trd.ClsSpd"))
	StrEql(t, e.ClsRsn, a.ClsRsn, append(msgs, "Trd.ClsRsn"))
	FltEql(t, e.Pip, a.Pip, append(msgs, "Trd.Pip"))
	TmeEql(t, e.Dur, a.Dur, append(msgs, "Trd.Dur"))
	BolEql(t, e.IsLong, a.IsLong, append(msgs, "Trd.IsLong"))
	FltEql(t, e.PnlPct, a.PnlPct, append(msgs, "Trd.PnlPct"))
	FltEql(t, e.PnlPctPredict, a.PnlPctPredict, append(msgs, "Trd.PnlPctPredict"))
	FltEql(t, e.PnlUsd, a.PnlUsd, append(msgs, "Trd.PnlUsd"))
	FltEql(t, e.PnlGrsUsd, a.PnlGrsUsd, append(msgs, "Trd.PnlGrsUsd"))
	FltEql(t, e.CstComUsd, a.CstComUsd, append(msgs, "Trd.CstComUsd"))
	FltEql(t, e.CstClsSpdUsd, a.CstClsSpdUsd, append(msgs, "Trd.CstClsSpdUsd"))
	FltEql(t, e.CstOpnSpdUsd, a.CstOpnSpdUsd, append(msgs, "Trd.CstOpnSpdUsd"))
	FltEql(t, e.OpnBalUsd, a.OpnBalUsd, append(msgs, "Trd.OpnBalUsd"))
	FltEql(t, e.ClsBalUsd, a.ClsBalUsd, append(msgs, "Trd.ClsBalUsd"))
	FltEql(t, e.ClsBalUsdAct, a.ClsBalUsdAct, append(msgs, "Trd.ClsBalUsdAct"))
	FltEql(t, e.TrdPct, a.TrdPct, append(msgs, "Trd.TrdPct"))
	FltEql(t, e.MrgnRtio, a.MrgnRtio, append(msgs, "Trd.MrgnRtio"))
	FltEql(t, e.Units, a.Units, append(msgs, "Trd.Units"))
	StrEql(t, e.Instr, a.Instr, append(msgs, "Trd.Instr"))
}
func AnaTrdNotZero(t *testing.T, a *ana.Trd, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	TmeNotZero(t, a.OpnTme, append(msgs, "Trd.OpnTme"))
	TmeNotZero(t, a.ClsTme, append(msgs, "Trd.ClsTme"))
	FltNotZero(t, a.OpnBid, append(msgs, "Trd.OpnBid"))
	FltNotZero(t, a.ClsBid, append(msgs, "Trd.ClsBid"))
	FltNotZero(t, a.OpnAsk, append(msgs, "Trd.OpnAsk"))
	FltNotZero(t, a.ClsAsk, append(msgs, "Trd.ClsAsk"))
	FltNotZero(t, a.OpnSpd, append(msgs, "Trd.OpnSpd"))
	FltNotZero(t, a.ClsSpd, append(msgs, "Trd.ClsSpd"))
	StrNotZero(t, a.ClsRsn, append(msgs, "Trd.ClsRsn"))
	FltNotZero(t, a.Pip, append(msgs, "Trd.Pip"))
	TmeNotZero(t, a.Dur, append(msgs, "Trd.Dur"))
	BolNotZero(t, a.IsLong, append(msgs, "Trd.IsLong"))
	FltNotZero(t, a.PnlPct, append(msgs, "Trd.PnlPct"))
	FltNotZero(t, a.PnlPctPredict, append(msgs, "Trd.PnlPctPredict"))
	FltNotZero(t, a.PnlUsd, append(msgs, "Trd.PnlUsd"))
	FltNotZero(t, a.PnlGrsUsd, append(msgs, "Trd.PnlGrsUsd"))
	FltNotZero(t, a.CstComUsd, append(msgs, "Trd.CstComUsd"))
	FltNotZero(t, a.CstClsSpdUsd, append(msgs, "Trd.CstClsSpdUsd"))
	FltNotZero(t, a.CstOpnSpdUsd, append(msgs, "Trd.CstOpnSpdUsd"))
	FltNotZero(t, a.OpnBalUsd, append(msgs, "Trd.OpnBalUsd"))
	FltNotZero(t, a.ClsBalUsd, append(msgs, "Trd.ClsBalUsd"))
	FltNotZero(t, a.ClsBalUsdAct, append(msgs, "Trd.ClsBalUsdAct"))
	FltNotZero(t, a.TrdPct, append(msgs, "Trd.TrdPct"))
	FltNotZero(t, a.MrgnRtio, append(msgs, "Trd.MrgnRtio"))
	FltNotZero(t, a.Units, append(msgs, "Trd.Units"))
	StrNotZero(t, a.Instr, append(msgs, "Trd.Instr"))
}
func AnaTrdSliceEql(t *testing.T, e, a []*ana.Trd, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		AnaTrdEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm Trd (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func AnaTrdsEql(t *testing.T, e, a *ana.Trds, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	if len(*e) != len(*a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(*e), len(*a)))...)
	}
	for n := 0; n < len(*e); n++ {
		if (*e)[n] == nil && (*a)[n] == nil {
			continue
		}
		if (*e)[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if (*a)[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		AnaTrdEql(t, (*e)[n], (*a)[n], append(msgs, fmt.Sprintf("elm Trd (idx:%v expected:%v actual:%v)", n, (*e)[n], (*a)[n])))
	}
}
func AnaTrdsNotZero(t *testing.T, a *ana.Trds, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	if a.Cnt() == 0 {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("cnt is zero"))...)
	}
}
func AnaTrdsSliceEql(t *testing.T, e, a []*ana.Trds, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		AnaTrdsEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm Trds (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func AnaPrfmEql(t *testing.T, e, a *ana.Prfm, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	FltEql(t, e.PnlPct, a.PnlPct, append(msgs, "Prfm.PnlPct"))
	FltEql(t, e.ScsPct, a.ScsPct, append(msgs, "Prfm.ScsPct"))
	FltEql(t, e.PipPerDay, a.PipPerDay, append(msgs, "Prfm.PipPerDay"))
	FltEql(t, e.UsdPerDay, a.UsdPerDay, append(msgs, "Prfm.UsdPerDay"))
	FltEql(t, e.ScsPerDay, a.ScsPerDay, append(msgs, "Prfm.ScsPerDay"))
	FltEql(t, e.OpnPerDay, a.OpnPerDay, append(msgs, "Prfm.OpnPerDay"))
	FltEql(t, e.PnlUsd, a.PnlUsd, append(msgs, "Prfm.PnlUsd"))
	FltEql(t, e.PipAvg, a.PipAvg, append(msgs, "Prfm.PipAvg"))
	FltEql(t, e.PipMdn, a.PipMdn, append(msgs, "Prfm.PipMdn"))
	FltEql(t, e.PipMin, a.PipMin, append(msgs, "Prfm.PipMin"))
	FltEql(t, e.PipMax, a.PipMax, append(msgs, "Prfm.PipMax"))
	FltEql(t, e.PipSum, a.PipSum, append(msgs, "Prfm.PipSum"))
	TmeEql(t, e.DurAvg, a.DurAvg, append(msgs, "Prfm.DurAvg"))
	TmeEql(t, e.DurMdn, a.DurMdn, append(msgs, "Prfm.DurMdn"))
	TmeEql(t, e.DurMin, a.DurMin, append(msgs, "Prfm.DurMin"))
	TmeEql(t, e.DurMax, a.DurMax, append(msgs, "Prfm.DurMax"))
	FltEql(t, e.LosLimMax, a.LosLimMax, append(msgs, "Prfm.LosLimMax"))
	TmeEql(t, e.DurLimMax, a.DurLimMax, append(msgs, "Prfm.DurLimMax"))
	UntEql(t, e.DayCnt, a.DayCnt, append(msgs, "Prfm.DayCnt"))
	UntEql(t, e.TrdCnt, a.TrdCnt, append(msgs, "Prfm.TrdCnt"))
	FltEql(t, e.TrdPct, a.TrdPct, append(msgs, "Prfm.TrdPct"))
	FltEql(t, e.CstTotUsd, a.CstTotUsd, append(msgs, "Prfm.CstTotUsd"))
	FltEql(t, e.CstSpdUsd, a.CstSpdUsd, append(msgs, "Prfm.CstSpdUsd"))
	FltEql(t, e.CstComUsd, a.CstComUsd, append(msgs, "Prfm.CstComUsd"))
}
func AnaPrfmNotZero(t *testing.T, a *ana.Prfm, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	FltNotZero(t, a.PipPerDay, append(msgs, "Prfm.PipPerDay"))
	FltNotZero(t, a.UsdPerDay, append(msgs, "Prfm.UsdPerDay"))
	FltNotZero(t, a.PnlUsd, append(msgs, "Prfm.PnlUsd"))
	FltNotZero(t, a.PipAvg, append(msgs, "Prfm.PipAvg"))
	FltNotZero(t, a.PipMdn, append(msgs, "Prfm.PipMdn"))
	FltNotZero(t, a.PipMin, append(msgs, "Prfm.PipMin"))
	FltNotZero(t, a.PipMax, append(msgs, "Prfm.PipMax"))
	FltNotZero(t, a.PipSum, append(msgs, "Prfm.PipSum"))
	TmeNotZero(t, a.DurAvg, append(msgs, "Prfm.DurAvg"))
	TmeNotZero(t, a.DurMdn, append(msgs, "Prfm.DurMdn"))
	TmeNotZero(t, a.DurMin, append(msgs, "Prfm.DurMin"))
	TmeNotZero(t, a.DurMax, append(msgs, "Prfm.DurMax"))
	FltNotZero(t, a.LosLimMax, append(msgs, "Prfm.LosLimMax"))
	TmeNotZero(t, a.DurLimMax, append(msgs, "Prfm.DurLimMax"))
	UntNotZero(t, a.DayCnt, append(msgs, "Prfm.DayCnt"))
	UntNotZero(t, a.TrdCnt, append(msgs, "Prfm.TrdCnt"))
	FltNotZero(t, a.TrdPct, append(msgs, "Prfm.TrdPct"))
	FltNotZero(t, a.CstTotUsd, append(msgs, "Prfm.CstTotUsd"))
	FltNotZero(t, a.CstSpdUsd, append(msgs, "Prfm.CstSpdUsd"))
	FltNotZero(t, a.CstComUsd, append(msgs, "Prfm.CstComUsd"))
}
func AnaPrfmSliceEql(t *testing.T, e, a []*ana.Prfm, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		AnaPrfmEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm Prfm (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func AnaPrfmsEql(t *testing.T, e, a *ana.Prfms, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	if len(*e) != len(*a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(*e), len(*a)))...)
	}
	for n := 0; n < len(*e); n++ {
		if (*e)[n] == nil && (*a)[n] == nil {
			continue
		}
		if (*e)[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if (*a)[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		AnaPrfmEql(t, (*e)[n], (*a)[n], append(msgs, fmt.Sprintf("elm Prfm (idx:%v expected:%v actual:%v)", n, (*e)[n], (*a)[n])))
	}
}
func AnaPrfmsNotZero(t *testing.T, a *ana.Prfms, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	if a.Cnt() == 0 {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("cnt is zero"))...)
	}
}
func AnaPrfmsSliceEql(t *testing.T, e, a []*ana.Prfms, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		AnaPrfmsEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm Prfms (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func AnaPortEql(t *testing.T, e, a *ana.Port, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	FltEql(t, e.BalFstUsd, a.BalFstUsd, append(msgs, "Port.BalFstUsd"))
	FltEql(t, e.BalLstUsd, a.BalLstUsd, append(msgs, "Port.BalLstUsd"))
	FltEql(t, e.TrdPct, a.TrdPct, append(msgs, "Port.TrdPct"))
	AnaTrdsEql(t, e.Trds, a.Trds, append(msgs, "Port.Trds"))
}
func AnaPortNotZero(t *testing.T, a *ana.Port, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	FltNotZero(t, a.BalFstUsd, append(msgs, "Port.BalFstUsd"))
	FltNotZero(t, a.BalLstUsd, append(msgs, "Port.BalLstUsd"))
	FltNotZero(t, a.TrdPct, append(msgs, "Port.TrdPct"))
}
func AnaPortSliceEql(t *testing.T, e, a []*ana.Port, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		AnaPortEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm Port (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstPrvEql(t *testing.T, e, a hst.Prv, msgs ...interface{}) {
	if (e == nil || reflect.ValueOf(e).IsNil()) && (a == nil || reflect.ValueOf(a).IsNil()) {
		return
	}
	StrEql(t, e.Name(), a.Name(), append(msgs, "Prv.Name()"))
	StringEql(t, e.Prm(), a.Prm(), append(msgs, "Prv.Prm()"))
	StringEql(t, e.String(), a.String(), append(msgs, "Prv.String()"))
}
func HstPrvNotZero(t *testing.T, a hst.Prv, msgs ...interface{}) {
	if a == nil || reflect.ValueOf(a).IsNil() {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
}
func HstPrvSliceEql(t *testing.T, e, a []hst.Prv, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstPrvEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm Prv (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstInstrEql(t *testing.T, e, a hst.Instr, msgs ...interface{}) {
	if (e == nil || reflect.ValueOf(e).IsNil()) && (a == nil || reflect.ValueOf(a).IsNil()) {
		return
	}
	StrEql(t, e.Name(), a.Name(), append(msgs, "Instr.Name()"))
	StringEql(t, e.Prm(), a.Prm(), append(msgs, "Instr.Prm()"))
	StringEql(t, e.String(), a.String(), append(msgs, "Instr.String()"))
}
func HstInstrNotZero(t *testing.T, a hst.Instr, msgs ...interface{}) {
	if a == nil || reflect.ValueOf(a).IsNil() {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
}
func HstInstrSliceEql(t *testing.T, e, a []hst.Instr, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstInstrEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm Instr (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstInrvlEql(t *testing.T, e, a hst.Inrvl, msgs ...interface{}) {
	if (e == nil || reflect.ValueOf(e).IsNil()) && (a == nil || reflect.ValueOf(a).IsNil()) {
		return
	}
	StrEql(t, e.Name(), a.Name(), append(msgs, "Inrvl.Name()"))
	StringEql(t, e.Prm(), a.Prm(), append(msgs, "Inrvl.Prm()"))
	StringEql(t, e.String(), a.String(), append(msgs, "Inrvl.String()"))
}
func HstInrvlNotZero(t *testing.T, a hst.Inrvl, msgs ...interface{}) {
	if a == nil || reflect.ValueOf(a).IsNil() {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
}
func HstInrvlSliceEql(t *testing.T, e, a []hst.Inrvl, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstInrvlEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm Inrvl (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstSideEql(t *testing.T, e, a hst.Side, msgs ...interface{}) {
	if (e == nil || reflect.ValueOf(e).IsNil()) && (a == nil || reflect.ValueOf(a).IsNil()) {
		return
	}
	StrEql(t, e.Name(), a.Name(), append(msgs, "Side.Name()"))
	StringEql(t, e.Prm(), a.Prm(), append(msgs, "Side.Prm()"))
	StringEql(t, e.String(), a.String(), append(msgs, "Side.String()"))
}
func HstSideNotZero(t *testing.T, a hst.Side, msgs ...interface{}) {
	if a == nil || reflect.ValueOf(a).IsNil() {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
}
func HstSideSliceEql(t *testing.T, e, a []hst.Side, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstSideEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm Side (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstStmEql(t *testing.T, e, a hst.Stm, msgs ...interface{}) {
	if (e == nil || reflect.ValueOf(e).IsNil()) && (a == nil || reflect.ValueOf(a).IsNil()) {
		return
	}
	StrEql(t, e.Name(), a.Name(), append(msgs, "Stm.Name()"))
	StringEql(t, e.Prm(), a.Prm(), append(msgs, "Stm.Prm()"))
	StringEql(t, e.String(), a.String(), append(msgs, "Stm.String()"))
}
func HstStmNotZero(t *testing.T, a hst.Stm, msgs ...interface{}) {
	if a == nil || reflect.ValueOf(a).IsNil() {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
}
func HstStmSliceEql(t *testing.T, e, a []hst.Stm, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstStmEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm Stm (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstCndEql(t *testing.T, e, a hst.Cnd, msgs ...interface{}) {
	if (e == nil || reflect.ValueOf(e).IsNil()) && (a == nil || reflect.ValueOf(a).IsNil()) {
		return
	}
	StrEql(t, e.Name(), a.Name(), append(msgs, "Cnd.Name()"))
	StringEql(t, e.Prm(), a.Prm(), append(msgs, "Cnd.Prm()"))
	StringEql(t, e.String(), a.String(), append(msgs, "Cnd.String()"))
}
func HstCndNotZero(t *testing.T, a hst.Cnd, msgs ...interface{}) {
	if a == nil || reflect.ValueOf(a).IsNil() {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
}
func HstCndSliceEql(t *testing.T, e, a []hst.Cnd, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstCndEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm Cnd (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstStgyEql(t *testing.T, e, a hst.Stgy, msgs ...interface{}) {
	if (e == nil || reflect.ValueOf(e).IsNil()) && (a == nil || reflect.ValueOf(a).IsNil()) {
		return
	}
	StrEql(t, e.Name(), a.Name(), append(msgs, "Stgy.Name()"))
	StringEql(t, e.Prm(), a.Prm(), append(msgs, "Stgy.Prm()"))
	StringEql(t, e.String(), a.String(), append(msgs, "Stgy.String()"))
}
func HstStgyNotZero(t *testing.T, a hst.Stgy, msgs ...interface{}) {
	if a == nil || reflect.ValueOf(a).IsNil() {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
}
func HstStgySliceEql(t *testing.T, e, a []hst.Stgy, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstStgyEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm Stgy (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstPrvsEql(t *testing.T, e, a *hst.Prvs, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	if len(*e) != len(*a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(*e), len(*a)))...)
	}
	for n := 0; n < len(*e); n++ {
		if (*e)[n] == nil && (*a)[n] == nil {
			continue
		}
		if (*e)[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if (*a)[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstPrvEql(t, (*e)[n], (*a)[n], append(msgs, fmt.Sprintf("elm Prv (idx:%v expected:%v actual:%v)", n, (*e)[n], (*a)[n])))
	}
}
func HstPrvsNotZero(t *testing.T, a *hst.Prvs, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	if a.Cnt() == 0 {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("cnt is zero"))...)
	}
}
func HstPrvsSliceEql(t *testing.T, e, a []*hst.Prvs, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		HstPrvsEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm Prvs (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstInstrsEql(t *testing.T, e, a *hst.Instrs, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	if len(*e) != len(*a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(*e), len(*a)))...)
	}
	for n := 0; n < len(*e); n++ {
		if (*e)[n] == nil && (*a)[n] == nil {
			continue
		}
		if (*e)[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if (*a)[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstInstrEql(t, (*e)[n], (*a)[n], append(msgs, fmt.Sprintf("elm Instr (idx:%v expected:%v actual:%v)", n, (*e)[n], (*a)[n])))
	}
}
func HstInstrsNotZero(t *testing.T, a *hst.Instrs, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	if a.Cnt() == 0 {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("cnt is zero"))...)
	}
}
func HstInstrsSliceEql(t *testing.T, e, a []*hst.Instrs, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		HstInstrsEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm Instrs (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstInrvlsEql(t *testing.T, e, a *hst.Inrvls, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	if len(*e) != len(*a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(*e), len(*a)))...)
	}
	for n := 0; n < len(*e); n++ {
		if (*e)[n] == nil && (*a)[n] == nil {
			continue
		}
		if (*e)[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if (*a)[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstInrvlEql(t, (*e)[n], (*a)[n], append(msgs, fmt.Sprintf("elm Inrvl (idx:%v expected:%v actual:%v)", n, (*e)[n], (*a)[n])))
	}
}
func HstInrvlsNotZero(t *testing.T, a *hst.Inrvls, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	if a.Cnt() == 0 {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("cnt is zero"))...)
	}
}
func HstInrvlsSliceEql(t *testing.T, e, a []*hst.Inrvls, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		HstInrvlsEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm Inrvls (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstSidesEql(t *testing.T, e, a *hst.Sides, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	if len(*e) != len(*a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(*e), len(*a)))...)
	}
	for n := 0; n < len(*e); n++ {
		if (*e)[n] == nil && (*a)[n] == nil {
			continue
		}
		if (*e)[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if (*a)[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstSideEql(t, (*e)[n], (*a)[n], append(msgs, fmt.Sprintf("elm Side (idx:%v expected:%v actual:%v)", n, (*e)[n], (*a)[n])))
	}
}
func HstSidesNotZero(t *testing.T, a *hst.Sides, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	if a.Cnt() == 0 {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("cnt is zero"))...)
	}
}
func HstSidesSliceEql(t *testing.T, e, a []*hst.Sides, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		HstSidesEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm Sides (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstStmsEql(t *testing.T, e, a *hst.Stms, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	if len(*e) != len(*a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(*e), len(*a)))...)
	}
	for n := 0; n < len(*e); n++ {
		if (*e)[n] == nil && (*a)[n] == nil {
			continue
		}
		if (*e)[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if (*a)[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstStmEql(t, (*e)[n], (*a)[n], append(msgs, fmt.Sprintf("elm Stm (idx:%v expected:%v actual:%v)", n, (*e)[n], (*a)[n])))
	}
}
func HstStmsNotZero(t *testing.T, a *hst.Stms, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	if a.Cnt() == 0 {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("cnt is zero"))...)
	}
}
func HstStmsSliceEql(t *testing.T, e, a []*hst.Stms, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		HstStmsEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm Stms (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstCndsEql(t *testing.T, e, a *hst.Cnds, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	if len(*e) != len(*a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(*e), len(*a)))...)
	}
	for n := 0; n < len(*e); n++ {
		if (*e)[n] == nil && (*a)[n] == nil {
			continue
		}
		if (*e)[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if (*a)[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstCndEql(t, (*e)[n], (*a)[n], append(msgs, fmt.Sprintf("elm Cnd (idx:%v expected:%v actual:%v)", n, (*e)[n], (*a)[n])))
	}
}
func HstCndsNotZero(t *testing.T, a *hst.Cnds, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	if a.Cnt() == 0 {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("cnt is zero"))...)
	}
}
func HstCndsSliceEql(t *testing.T, e, a []*hst.Cnds, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		HstCndsEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm Cnds (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstStgysEql(t *testing.T, e, a *hst.Stgys, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	if len(*e) != len(*a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(*e), len(*a)))...)
	}
	for n := 0; n < len(*e); n++ {
		if (*e)[n] == nil && (*a)[n] == nil {
			continue
		}
		if (*e)[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if (*a)[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstStgyEql(t, (*e)[n], (*a)[n], append(msgs, fmt.Sprintf("elm Stgy (idx:%v expected:%v actual:%v)", n, (*e)[n], (*a)[n])))
	}
}
func HstStgysNotZero(t *testing.T, a *hst.Stgys, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	if a.Cnt() == 0 {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("cnt is zero"))...)
	}
}
func HstStgysSliceEql(t *testing.T, e, a []*hst.Stgys, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		HstStgysEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm Stgys (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltPrvEql(t *testing.T, e, a rlt.Prv, msgs ...interface{}) {
	if (e == nil || reflect.ValueOf(e).IsNil()) && (a == nil || reflect.ValueOf(a).IsNil()) {
		return
	}
	BolEql(t, e.MayTrd(), a.MayTrd(), append(msgs, "Prv.MayTrd()"))
	StrEql(t, e.Name(), a.Name(), append(msgs, "Pth.Name()"))
	StringEql(t, e.Prm(), a.Prm(), append(msgs, "Pth.Prm()"))
	StringEql(t, e.String(), a.String(), append(msgs, "Pth.String()"))
}
func RltPrvNotZero(t *testing.T, a rlt.Prv, msgs ...interface{}) {
	if a == nil || reflect.ValueOf(a).IsNil() {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
}
func RltPrvSliceEql(t *testing.T, e, a []rlt.Prv, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltPrvEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm Prv (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltInstrEql(t *testing.T, e, a rlt.Instr, msgs ...interface{}) {
	if (e == nil || reflect.ValueOf(e).IsNil()) && (a == nil || reflect.ValueOf(a).IsNil()) {
		return
	}
	StrEql(t, e.Name(), a.Name(), append(msgs, "Pth.Name()"))
	StringEql(t, e.Prm(), a.Prm(), append(msgs, "Pth.Prm()"))
	StringEql(t, e.String(), a.String(), append(msgs, "Pth.String()"))
}
func RltInstrNotZero(t *testing.T, a rlt.Instr, msgs ...interface{}) {
	if a == nil || reflect.ValueOf(a).IsNil() {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
}
func RltInstrSliceEql(t *testing.T, e, a []rlt.Instr, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltInstrEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm Instr (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltInrvlEql(t *testing.T, e, a rlt.Inrvl, msgs ...interface{}) {
	if (e == nil || reflect.ValueOf(e).IsNil()) && (a == nil || reflect.ValueOf(a).IsNil()) {
		return
	}
	StrEql(t, e.Name(), a.Name(), append(msgs, "Pth.Name()"))
	StringEql(t, e.Prm(), a.Prm(), append(msgs, "Pth.Prm()"))
	StringEql(t, e.String(), a.String(), append(msgs, "Pth.String()"))
}
func RltInrvlNotZero(t *testing.T, a rlt.Inrvl, msgs ...interface{}) {
	if a == nil || reflect.ValueOf(a).IsNil() {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
}
func RltInrvlSliceEql(t *testing.T, e, a []rlt.Inrvl, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltInrvlEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm Inrvl (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltSideEql(t *testing.T, e, a rlt.Side, msgs ...interface{}) {
	if (e == nil || reflect.ValueOf(e).IsNil()) && (a == nil || reflect.ValueOf(a).IsNil()) {
		return
	}
	StrEql(t, e.Name(), a.Name(), append(msgs, "Pth.Name()"))
	StringEql(t, e.Prm(), a.Prm(), append(msgs, "Pth.Prm()"))
	StringEql(t, e.String(), a.String(), append(msgs, "Pth.String()"))
}
func RltSideNotZero(t *testing.T, a rlt.Side, msgs ...interface{}) {
	if a == nil || reflect.ValueOf(a).IsNil() {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
}
func RltSideSliceEql(t *testing.T, e, a []rlt.Side, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltSideEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm Side (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltStmEql(t *testing.T, e, a rlt.Stm, msgs ...interface{}) {
	if (e == nil || reflect.ValueOf(e).IsNil()) && (a == nil || reflect.ValueOf(a).IsNil()) {
		return
	}
	StrEql(t, e.Name(), a.Name(), append(msgs, "Pth.Name()"))
	StringEql(t, e.Prm(), a.Prm(), append(msgs, "Pth.Prm()"))
	StringEql(t, e.String(), a.String(), append(msgs, "Pth.String()"))
}
func RltStmNotZero(t *testing.T, a rlt.Stm, msgs ...interface{}) {
	if a == nil || reflect.ValueOf(a).IsNil() {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
}
func RltStmSliceEql(t *testing.T, e, a []rlt.Stm, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltStmEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm Stm (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltCndEql(t *testing.T, e, a rlt.Cnd, msgs ...interface{}) {
	if (e == nil || reflect.ValueOf(e).IsNil()) && (a == nil || reflect.ValueOf(a).IsNil()) {
		return
	}
	StrEql(t, e.Name(), a.Name(), append(msgs, "Pth.Name()"))
	StringEql(t, e.Prm(), a.Prm(), append(msgs, "Pth.Prm()"))
	StringEql(t, e.String(), a.String(), append(msgs, "Pth.String()"))
}
func RltCndNotZero(t *testing.T, a rlt.Cnd, msgs ...interface{}) {
	if a == nil || reflect.ValueOf(a).IsNil() {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
}
func RltCndSliceEql(t *testing.T, e, a []rlt.Cnd, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltCndEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm Cnd (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltStgyEql(t *testing.T, e, a rlt.Stgy, msgs ...interface{}) {
	if (e == nil || reflect.ValueOf(e).IsNil()) && (a == nil || reflect.ValueOf(a).IsNil()) {
		return
	}
	StrEql(t, e.Name(), a.Name(), append(msgs, "Pth.Name()"))
	StringEql(t, e.Prm(), a.Prm(), append(msgs, "Pth.Prm()"))
	StringEql(t, e.String(), a.String(), append(msgs, "Pth.String()"))
}
func RltStgyNotZero(t *testing.T, a rlt.Stgy, msgs ...interface{}) {
	if a == nil || reflect.ValueOf(a).IsNil() {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
}
func RltStgySliceEql(t *testing.T, e, a []rlt.Stgy, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltStgyEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm Stgy (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltPrvsEql(t *testing.T, e, a *rlt.Prvs, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	if len(*e) != len(*a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(*e), len(*a)))...)
	}
	for n := 0; n < len(*e); n++ {
		if (*e)[n] == nil && (*a)[n] == nil {
			continue
		}
		if (*e)[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if (*a)[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltPrvEql(t, (*e)[n], (*a)[n], append(msgs, fmt.Sprintf("elm Prv (idx:%v expected:%v actual:%v)", n, (*e)[n], (*a)[n])))
	}
}
func RltPrvsNotZero(t *testing.T, a *rlt.Prvs, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	if a.Cnt() == 0 {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("cnt is zero"))...)
	}
}
func RltPrvsSliceEql(t *testing.T, e, a []*rlt.Prvs, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		RltPrvsEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm Prvs (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltInstrsEql(t *testing.T, e, a *rlt.Instrs, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	if len(*e) != len(*a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(*e), len(*a)))...)
	}
	for n := 0; n < len(*e); n++ {
		if (*e)[n] == nil && (*a)[n] == nil {
			continue
		}
		if (*e)[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if (*a)[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltInstrEql(t, (*e)[n], (*a)[n], append(msgs, fmt.Sprintf("elm Instr (idx:%v expected:%v actual:%v)", n, (*e)[n], (*a)[n])))
	}
}
func RltInstrsNotZero(t *testing.T, a *rlt.Instrs, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	if a.Cnt() == 0 {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("cnt is zero"))...)
	}
}
func RltInstrsSliceEql(t *testing.T, e, a []*rlt.Instrs, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		RltInstrsEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm Instrs (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltInrvlsEql(t *testing.T, e, a *rlt.Inrvls, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	if len(*e) != len(*a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(*e), len(*a)))...)
	}
	for n := 0; n < len(*e); n++ {
		if (*e)[n] == nil && (*a)[n] == nil {
			continue
		}
		if (*e)[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if (*a)[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltInrvlEql(t, (*e)[n], (*a)[n], append(msgs, fmt.Sprintf("elm Inrvl (idx:%v expected:%v actual:%v)", n, (*e)[n], (*a)[n])))
	}
}
func RltInrvlsNotZero(t *testing.T, a *rlt.Inrvls, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	if a.Cnt() == 0 {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("cnt is zero"))...)
	}
}
func RltInrvlsSliceEql(t *testing.T, e, a []*rlt.Inrvls, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		RltInrvlsEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm Inrvls (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltSidesEql(t *testing.T, e, a *rlt.Sides, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	if len(*e) != len(*a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(*e), len(*a)))...)
	}
	for n := 0; n < len(*e); n++ {
		if (*e)[n] == nil && (*a)[n] == nil {
			continue
		}
		if (*e)[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if (*a)[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltSideEql(t, (*e)[n], (*a)[n], append(msgs, fmt.Sprintf("elm Side (idx:%v expected:%v actual:%v)", n, (*e)[n], (*a)[n])))
	}
}
func RltSidesNotZero(t *testing.T, a *rlt.Sides, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	if a.Cnt() == 0 {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("cnt is zero"))...)
	}
}
func RltSidesSliceEql(t *testing.T, e, a []*rlt.Sides, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		RltSidesEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm Sides (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltStmsEql(t *testing.T, e, a *rlt.Stms, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	if len(*e) != len(*a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(*e), len(*a)))...)
	}
	for n := 0; n < len(*e); n++ {
		if (*e)[n] == nil && (*a)[n] == nil {
			continue
		}
		if (*e)[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if (*a)[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltStmEql(t, (*e)[n], (*a)[n], append(msgs, fmt.Sprintf("elm Stm (idx:%v expected:%v actual:%v)", n, (*e)[n], (*a)[n])))
	}
}
func RltStmsNotZero(t *testing.T, a *rlt.Stms, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	if a.Cnt() == 0 {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("cnt is zero"))...)
	}
}
func RltStmsSliceEql(t *testing.T, e, a []*rlt.Stms, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		RltStmsEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm Stms (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltCndsEql(t *testing.T, e, a *rlt.Cnds, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	if len(*e) != len(*a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(*e), len(*a)))...)
	}
	for n := 0; n < len(*e); n++ {
		if (*e)[n] == nil && (*a)[n] == nil {
			continue
		}
		if (*e)[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if (*a)[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltCndEql(t, (*e)[n], (*a)[n], append(msgs, fmt.Sprintf("elm Cnd (idx:%v expected:%v actual:%v)", n, (*e)[n], (*a)[n])))
	}
}
func RltCndsNotZero(t *testing.T, a *rlt.Cnds, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	if a.Cnt() == 0 {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("cnt is zero"))...)
	}
}
func RltCndsSliceEql(t *testing.T, e, a []*rlt.Cnds, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		RltCndsEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm Cnds (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltStgysEql(t *testing.T, e, a *rlt.Stgys, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	if len(*e) != len(*a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(*e), len(*a)))...)
	}
	for n := 0; n < len(*e); n++ {
		if (*e)[n] == nil && (*a)[n] == nil {
			continue
		}
		if (*e)[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if (*a)[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltStgyEql(t, (*e)[n], (*a)[n], append(msgs, fmt.Sprintf("elm Stgy (idx:%v expected:%v actual:%v)", n, (*e)[n], (*a)[n])))
	}
}
func RltStgysNotZero(t *testing.T, a *rlt.Stgys, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	if a.Cnt() == 0 {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("cnt is zero"))...)
	}
}
func RltStgysSliceEql(t *testing.T, e, a []*rlt.Stgys, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		RltStgysEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm Stgys (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstPrvOanEql(t *testing.T, e, a *hst.PrvOan, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
}
func HstPrvOanNotZero(t *testing.T, a *hst.PrvOan, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
}
func HstPrvOanSliceEql(t *testing.T, e, a []*hst.PrvOan, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstPrvOanEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm PrvOan (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltPrvOanEql(t *testing.T, e, a *rlt.PrvOan, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
}
func RltPrvOanNotZero(t *testing.T, a *rlt.PrvOan, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
}
func RltPrvOanSliceEql(t *testing.T, e, a []*rlt.PrvOan, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltPrvOanEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm PrvOan (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func StrsAsc(t *testing.T, a *strs.Strs, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "actual is nil")...)
	}
	for n := 1; n < len(*a); n++ {
		if !str.Lss((*a)[n-1], (*a)[n]) && (*a)[n-1] != (*a)[n] {
			t.Helper()
			t.Fatal(append(msgs, fmt.Errorf("not in asc order (idxs %v,%v vals %v,%v)", n-1, n, (*a)[n-1], (*a)[n]))...)
		}
	}
}
func StrsDsc(t *testing.T, a *strs.Strs, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "actual is nil")...)
	}
	for n := 1; n < len(*a); n++ {
		if !str.Gtr((*a)[n-1], (*a)[n]) && (*a)[n-1] != (*a)[n] {
			t.Helper()
			t.Fatal(append(msgs, fmt.Errorf("not in dsc order (idxs %v,%v vals %v,%v)", n-1, n, (*a)[n-1], (*a)[n]))...)
		}
	}
}
func FltsAsc(t *testing.T, a *flts.Flts, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "actual is nil")...)
	}
	for n := 1; n < len(*a); n++ {
		if !flt.Lss((*a)[n-1], (*a)[n]) && (*a)[n-1] != (*a)[n] {
			t.Helper()
			t.Fatal(append(msgs, fmt.Errorf("not in asc order (idxs %v,%v vals %v,%v)", n-1, n, (*a)[n-1], (*a)[n]))...)
		}
	}
}
func FltsDsc(t *testing.T, a *flts.Flts, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "actual is nil")...)
	}
	for n := 1; n < len(*a); n++ {
		if !flt.Gtr((*a)[n-1], (*a)[n]) && (*a)[n-1] != (*a)[n] {
			t.Helper()
			t.Fatal(append(msgs, fmt.Errorf("not in dsc order (idxs %v,%v vals %v,%v)", n-1, n, (*a)[n-1], (*a)[n]))...)
		}
	}
}
func UntsAsc(t *testing.T, a *unts.Unts, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "actual is nil")...)
	}
	for n := 1; n < len(*a); n++ {
		if !unt.Lss((*a)[n-1], (*a)[n]) && (*a)[n-1] != (*a)[n] {
			t.Helper()
			t.Fatal(append(msgs, fmt.Errorf("not in asc order (idxs %v,%v vals %v,%v)", n-1, n, (*a)[n-1], (*a)[n]))...)
		}
	}
}
func UntsDsc(t *testing.T, a *unts.Unts, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "actual is nil")...)
	}
	for n := 1; n < len(*a); n++ {
		if !unt.Gtr((*a)[n-1], (*a)[n]) && (*a)[n-1] != (*a)[n] {
			t.Helper()
			t.Fatal(append(msgs, fmt.Errorf("not in dsc order (idxs %v,%v vals %v,%v)", n-1, n, (*a)[n-1], (*a)[n]))...)
		}
	}
}
func IntsAsc(t *testing.T, a *ints.Ints, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "actual is nil")...)
	}
	for n := 1; n < len(*a); n++ {
		if !int.Lss((*a)[n-1], (*a)[n]) && (*a)[n-1] != (*a)[n] {
			t.Helper()
			t.Fatal(append(msgs, fmt.Errorf("not in asc order (idxs %v,%v vals %v,%v)", n-1, n, (*a)[n-1], (*a)[n]))...)
		}
	}
}
func IntsDsc(t *testing.T, a *ints.Ints, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "actual is nil")...)
	}
	for n := 1; n < len(*a); n++ {
		if !int.Gtr((*a)[n-1], (*a)[n]) && (*a)[n-1] != (*a)[n] {
			t.Helper()
			t.Fatal(append(msgs, fmt.Errorf("not in dsc order (idxs %v,%v vals %v,%v)", n-1, n, (*a)[n-1], (*a)[n]))...)
		}
	}
}
func TmesAsc(t *testing.T, a *tmes.Tmes, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "actual is nil")...)
	}
	for n := 1; n < len(*a); n++ {
		if !tme.Lss((*a)[n-1], (*a)[n]) && (*a)[n-1] != (*a)[n] {
			t.Helper()
			t.Fatal(append(msgs, fmt.Errorf("not in asc order (idxs %v,%v vals %v,%v)", n-1, n, (*a)[n-1], (*a)[n]))...)
		}
	}
}
func TmesDsc(t *testing.T, a *tmes.Tmes, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "actual is nil")...)
	}
	for n := 1; n < len(*a); n++ {
		if !tme.Gtr((*a)[n-1], (*a)[n]) && (*a)[n-1] != (*a)[n] {
			t.Helper()
			t.Fatal(append(msgs, fmt.Errorf("not in dsc order (idxs %v,%v vals %v,%v)", n-1, n, (*a)[n-1], (*a)[n]))...)
		}
	}
}
func HstInstrEurUsdEql(t *testing.T, e, a *hst.InstrEurUsd, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	HstPrvEql(t, e.Prv, a.Prv, append(msgs, "InstrEurUsd.Prv"))
	AnaInstrEql(t, e.Ana, a.Ana, append(msgs, "InstrEurUsd.Ana"))
	BndEql(t, e.TmeBnd, a.TmeBnd, append(msgs, "InstrEurUsd.TmeBnd"))
	TmeRngSliceEql(t, e.Rng, a.Rng, append(msgs, "InstrEurUsd.Rng"))
}
func HstInstrEurUsdNotZero(t *testing.T, a *hst.InstrEurUsd, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	HstPrvNotZero(t, a.Prv, append(msgs, "InstrEurUsd.Prv"))
	AnaInstrNotZero(t, a.Ana, append(msgs, "InstrEurUsd.Ana"))
	BndNotZero(t, a.TmeBnd, append(msgs, "InstrEurUsd.TmeBnd"))
}
func HstInstrEurUsdSliceEql(t *testing.T, e, a []*hst.InstrEurUsd, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstInstrEurUsdEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm InstrEurUsd (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstInstrAudUsdEql(t *testing.T, e, a *hst.InstrAudUsd, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	HstPrvEql(t, e.Prv, a.Prv, append(msgs, "InstrAudUsd.Prv"))
	AnaInstrEql(t, e.Ana, a.Ana, append(msgs, "InstrAudUsd.Ana"))
	BndEql(t, e.TmeBnd, a.TmeBnd, append(msgs, "InstrAudUsd.TmeBnd"))
	TmeRngSliceEql(t, e.Rng, a.Rng, append(msgs, "InstrAudUsd.Rng"))
}
func HstInstrAudUsdNotZero(t *testing.T, a *hst.InstrAudUsd, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	HstPrvNotZero(t, a.Prv, append(msgs, "InstrAudUsd.Prv"))
	AnaInstrNotZero(t, a.Ana, append(msgs, "InstrAudUsd.Ana"))
	BndNotZero(t, a.TmeBnd, append(msgs, "InstrAudUsd.TmeBnd"))
}
func HstInstrAudUsdSliceEql(t *testing.T, e, a []*hst.InstrAudUsd, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstInstrAudUsdEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm InstrAudUsd (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstInstrNzdUsdEql(t *testing.T, e, a *hst.InstrNzdUsd, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	HstPrvEql(t, e.Prv, a.Prv, append(msgs, "InstrNzdUsd.Prv"))
	AnaInstrEql(t, e.Ana, a.Ana, append(msgs, "InstrNzdUsd.Ana"))
	BndEql(t, e.TmeBnd, a.TmeBnd, append(msgs, "InstrNzdUsd.TmeBnd"))
	TmeRngSliceEql(t, e.Rng, a.Rng, append(msgs, "InstrNzdUsd.Rng"))
}
func HstInstrNzdUsdNotZero(t *testing.T, a *hst.InstrNzdUsd, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	HstPrvNotZero(t, a.Prv, append(msgs, "InstrNzdUsd.Prv"))
	AnaInstrNotZero(t, a.Ana, append(msgs, "InstrNzdUsd.Ana"))
	BndNotZero(t, a.TmeBnd, append(msgs, "InstrNzdUsd.TmeBnd"))
}
func HstInstrNzdUsdSliceEql(t *testing.T, e, a []*hst.InstrNzdUsd, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstInstrNzdUsdEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm InstrNzdUsd (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstInstrGbpUsdEql(t *testing.T, e, a *hst.InstrGbpUsd, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	HstPrvEql(t, e.Prv, a.Prv, append(msgs, "InstrGbpUsd.Prv"))
	AnaInstrEql(t, e.Ana, a.Ana, append(msgs, "InstrGbpUsd.Ana"))
	BndEql(t, e.TmeBnd, a.TmeBnd, append(msgs, "InstrGbpUsd.TmeBnd"))
	TmeRngSliceEql(t, e.Rng, a.Rng, append(msgs, "InstrGbpUsd.Rng"))
}
func HstInstrGbpUsdNotZero(t *testing.T, a *hst.InstrGbpUsd, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	HstPrvNotZero(t, a.Prv, append(msgs, "InstrGbpUsd.Prv"))
	AnaInstrNotZero(t, a.Ana, append(msgs, "InstrGbpUsd.Ana"))
	BndNotZero(t, a.TmeBnd, append(msgs, "InstrGbpUsd.TmeBnd"))
}
func HstInstrGbpUsdSliceEql(t *testing.T, e, a []*hst.InstrGbpUsd, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstInstrGbpUsdEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm InstrGbpUsd (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstInrvlIEql(t *testing.T, e, a *hst.InrvlI, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	HstInstrEql(t, e.Instr, a.Instr, append(msgs, "InrvlI.Instr"))
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "InrvlI.Tmes"))
	BndsEql(t, e.TmeBnds, a.TmeBnds, append(msgs, "InrvlI.TmeBnds"))
	TmeEql(t, e.Dur, a.Dur, append(msgs, "InrvlI.Dur"))
}
func HstInrvlINotZero(t *testing.T, a *hst.InrvlI, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	HstInstrNotZero(t, a.Instr, append(msgs, "InrvlI.Instr"))
}
func HstInrvlISliceEql(t *testing.T, e, a []*hst.InrvlI, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstInrvlIEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm InrvlI (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstSideBidEql(t *testing.T, e, a *hst.SideBid, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	HstInrvlEql(t, e.Inrvl, a.Inrvl, append(msgs, "SideBid.Inrvl"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "SideBid.Vals"))
	BndsEql(t, e.ValBnds, a.ValBnds, append(msgs, "SideBid.ValBnds"))
}
func HstSideBidNotZero(t *testing.T, a *hst.SideBid, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	HstInrvlNotZero(t, a.Inrvl, append(msgs, "SideBid.Inrvl"))
	FltsNotZero(t, a.Vals, append(msgs, "SideBid.Vals"))
	BndsNotZero(t, a.ValBnds, append(msgs, "SideBid.ValBnds"))
}
func HstSideBidSliceEql(t *testing.T, e, a []*hst.SideBid, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstSideBidEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm SideBid (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstSideAskEql(t *testing.T, e, a *hst.SideAsk, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	HstInrvlEql(t, e.Inrvl, a.Inrvl, append(msgs, "SideAsk.Inrvl"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "SideAsk.Vals"))
	BndsEql(t, e.ValBnds, a.ValBnds, append(msgs, "SideAsk.ValBnds"))
}
func HstSideAskNotZero(t *testing.T, a *hst.SideAsk, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	HstInrvlNotZero(t, a.Inrvl, append(msgs, "SideAsk.Inrvl"))
	FltsNotZero(t, a.Vals, append(msgs, "SideAsk.Vals"))
	BndsNotZero(t, a.ValBnds, append(msgs, "SideAsk.ValBnds"))
}
func HstSideAskSliceEql(t *testing.T, e, a []*hst.SideAsk, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstSideAskEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm SideAsk (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstStmRteFstEql(t *testing.T, e, a *hst.StmRteFst, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmRteFst.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmRteFst.Vals"))
	HstSideEql(t, e.Side, a.Side, append(msgs, "StmRteFst.Side"))
}
func HstStmRteFstNotZero(t *testing.T, a *hst.StmRteFst, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	TmesNotZero(t, a.Tmes, append(msgs, "StmRteFst.Tmes"))
	FltsNotZero(t, a.Vals, append(msgs, "StmRteFst.Vals"))
	HstSideNotZero(t, a.Side, append(msgs, "StmRteFst.Side"))
}
func HstStmRteFstSliceEql(t *testing.T, e, a []*hst.StmRteFst, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstStmRteFstEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmRteFst (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstStmRteLstEql(t *testing.T, e, a *hst.StmRteLst, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmRteLst.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmRteLst.Vals"))
	HstSideEql(t, e.Side, a.Side, append(msgs, "StmRteLst.Side"))
}
func HstStmRteLstNotZero(t *testing.T, a *hst.StmRteLst, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	TmesNotZero(t, a.Tmes, append(msgs, "StmRteLst.Tmes"))
	FltsNotZero(t, a.Vals, append(msgs, "StmRteLst.Vals"))
	HstSideNotZero(t, a.Side, append(msgs, "StmRteLst.Side"))
}
func HstStmRteLstSliceEql(t *testing.T, e, a []*hst.StmRteLst, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstStmRteLstEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmRteLst (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstStmRteSumEql(t *testing.T, e, a *hst.StmRteSum, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmRteSum.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmRteSum.Vals"))
	HstSideEql(t, e.Side, a.Side, append(msgs, "StmRteSum.Side"))
}
func HstStmRteSumNotZero(t *testing.T, a *hst.StmRteSum, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	TmesNotZero(t, a.Tmes, append(msgs, "StmRteSum.Tmes"))
	FltsNotZero(t, a.Vals, append(msgs, "StmRteSum.Vals"))
	HstSideNotZero(t, a.Side, append(msgs, "StmRteSum.Side"))
}
func HstStmRteSumSliceEql(t *testing.T, e, a []*hst.StmRteSum, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstStmRteSumEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmRteSum (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstStmRtePrdEql(t *testing.T, e, a *hst.StmRtePrd, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmRtePrd.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmRtePrd.Vals"))
	HstSideEql(t, e.Side, a.Side, append(msgs, "StmRtePrd.Side"))
}
func HstStmRtePrdNotZero(t *testing.T, a *hst.StmRtePrd, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	TmesNotZero(t, a.Tmes, append(msgs, "StmRtePrd.Tmes"))
	FltsNotZero(t, a.Vals, append(msgs, "StmRtePrd.Vals"))
	HstSideNotZero(t, a.Side, append(msgs, "StmRtePrd.Side"))
}
func HstStmRtePrdSliceEql(t *testing.T, e, a []*hst.StmRtePrd, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstStmRtePrdEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmRtePrd (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstStmRteMinEql(t *testing.T, e, a *hst.StmRteMin, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmRteMin.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmRteMin.Vals"))
	HstSideEql(t, e.Side, a.Side, append(msgs, "StmRteMin.Side"))
}
func HstStmRteMinNotZero(t *testing.T, a *hst.StmRteMin, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	TmesNotZero(t, a.Tmes, append(msgs, "StmRteMin.Tmes"))
	FltsNotZero(t, a.Vals, append(msgs, "StmRteMin.Vals"))
	HstSideNotZero(t, a.Side, append(msgs, "StmRteMin.Side"))
}
func HstStmRteMinSliceEql(t *testing.T, e, a []*hst.StmRteMin, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstStmRteMinEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmRteMin (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstStmRteMaxEql(t *testing.T, e, a *hst.StmRteMax, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmRteMax.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmRteMax.Vals"))
	HstSideEql(t, e.Side, a.Side, append(msgs, "StmRteMax.Side"))
}
func HstStmRteMaxNotZero(t *testing.T, a *hst.StmRteMax, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	TmesNotZero(t, a.Tmes, append(msgs, "StmRteMax.Tmes"))
	FltsNotZero(t, a.Vals, append(msgs, "StmRteMax.Vals"))
	HstSideNotZero(t, a.Side, append(msgs, "StmRteMax.Side"))
}
func HstStmRteMaxSliceEql(t *testing.T, e, a []*hst.StmRteMax, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstStmRteMaxEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmRteMax (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstStmRteMidEql(t *testing.T, e, a *hst.StmRteMid, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmRteMid.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmRteMid.Vals"))
	HstSideEql(t, e.Side, a.Side, append(msgs, "StmRteMid.Side"))
}
func HstStmRteMidNotZero(t *testing.T, a *hst.StmRteMid, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	TmesNotZero(t, a.Tmes, append(msgs, "StmRteMid.Tmes"))
	FltsNotZero(t, a.Vals, append(msgs, "StmRteMid.Vals"))
	HstSideNotZero(t, a.Side, append(msgs, "StmRteMid.Side"))
}
func HstStmRteMidSliceEql(t *testing.T, e, a []*hst.StmRteMid, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstStmRteMidEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmRteMid (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstStmRteMdnEql(t *testing.T, e, a *hst.StmRteMdn, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmRteMdn.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmRteMdn.Vals"))
	HstSideEql(t, e.Side, a.Side, append(msgs, "StmRteMdn.Side"))
}
func HstStmRteMdnNotZero(t *testing.T, a *hst.StmRteMdn, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	TmesNotZero(t, a.Tmes, append(msgs, "StmRteMdn.Tmes"))
	FltsNotZero(t, a.Vals, append(msgs, "StmRteMdn.Vals"))
	HstSideNotZero(t, a.Side, append(msgs, "StmRteMdn.Side"))
}
func HstStmRteMdnSliceEql(t *testing.T, e, a []*hst.StmRteMdn, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstStmRteMdnEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmRteMdn (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstStmRteSmaEql(t *testing.T, e, a *hst.StmRteSma, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmRteSma.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmRteSma.Vals"))
	HstSideEql(t, e.Side, a.Side, append(msgs, "StmRteSma.Side"))
}
func HstStmRteSmaNotZero(t *testing.T, a *hst.StmRteSma, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	TmesNotZero(t, a.Tmes, append(msgs, "StmRteSma.Tmes"))
	FltsNotZero(t, a.Vals, append(msgs, "StmRteSma.Vals"))
	HstSideNotZero(t, a.Side, append(msgs, "StmRteSma.Side"))
}
func HstStmRteSmaSliceEql(t *testing.T, e, a []*hst.StmRteSma, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstStmRteSmaEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmRteSma (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstStmRteGmaEql(t *testing.T, e, a *hst.StmRteGma, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmRteGma.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmRteGma.Vals"))
	HstSideEql(t, e.Side, a.Side, append(msgs, "StmRteGma.Side"))
}
func HstStmRteGmaNotZero(t *testing.T, a *hst.StmRteGma, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	TmesNotZero(t, a.Tmes, append(msgs, "StmRteGma.Tmes"))
	FltsNotZero(t, a.Vals, append(msgs, "StmRteGma.Vals"))
	HstSideNotZero(t, a.Side, append(msgs, "StmRteGma.Side"))
}
func HstStmRteGmaSliceEql(t *testing.T, e, a []*hst.StmRteGma, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstStmRteGmaEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmRteGma (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstStmRteWmaEql(t *testing.T, e, a *hst.StmRteWma, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmRteWma.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmRteWma.Vals"))
	HstSideEql(t, e.Side, a.Side, append(msgs, "StmRteWma.Side"))
}
func HstStmRteWmaNotZero(t *testing.T, a *hst.StmRteWma, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	TmesNotZero(t, a.Tmes, append(msgs, "StmRteWma.Tmes"))
	FltsNotZero(t, a.Vals, append(msgs, "StmRteWma.Vals"))
	HstSideNotZero(t, a.Side, append(msgs, "StmRteWma.Side"))
}
func HstStmRteWmaSliceEql(t *testing.T, e, a []*hst.StmRteWma, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstStmRteWmaEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmRteWma (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstStmRteRsiEql(t *testing.T, e, a *hst.StmRteRsi, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmRteRsi.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmRteRsi.Vals"))
	HstSideEql(t, e.Side, a.Side, append(msgs, "StmRteRsi.Side"))
}
func HstStmRteRsiNotZero(t *testing.T, a *hst.StmRteRsi, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	TmesNotZero(t, a.Tmes, append(msgs, "StmRteRsi.Tmes"))
	FltsNotZero(t, a.Vals, append(msgs, "StmRteRsi.Vals"))
	HstSideNotZero(t, a.Side, append(msgs, "StmRteRsi.Side"))
}
func HstStmRteRsiSliceEql(t *testing.T, e, a []*hst.StmRteRsi, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstStmRteRsiEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmRteRsi (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstStmRteWrsiEql(t *testing.T, e, a *hst.StmRteWrsi, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmRteWrsi.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmRteWrsi.Vals"))
	HstSideEql(t, e.Side, a.Side, append(msgs, "StmRteWrsi.Side"))
}
func HstStmRteWrsiNotZero(t *testing.T, a *hst.StmRteWrsi, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	TmesNotZero(t, a.Tmes, append(msgs, "StmRteWrsi.Tmes"))
	FltsNotZero(t, a.Vals, append(msgs, "StmRteWrsi.Vals"))
	HstSideNotZero(t, a.Side, append(msgs, "StmRteWrsi.Side"))
}
func HstStmRteWrsiSliceEql(t *testing.T, e, a []*hst.StmRteWrsi, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstStmRteWrsiEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmRteWrsi (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstStmRteAlmaEql(t *testing.T, e, a *hst.StmRteAlma, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmRteAlma.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmRteAlma.Vals"))
	HstSideEql(t, e.Side, a.Side, append(msgs, "StmRteAlma.Side"))
}
func HstStmRteAlmaNotZero(t *testing.T, a *hst.StmRteAlma, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	TmesNotZero(t, a.Tmes, append(msgs, "StmRteAlma.Tmes"))
	FltsNotZero(t, a.Vals, append(msgs, "StmRteAlma.Vals"))
	HstSideNotZero(t, a.Side, append(msgs, "StmRteAlma.Side"))
}
func HstStmRteAlmaSliceEql(t *testing.T, e, a []*hst.StmRteAlma, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstStmRteAlmaEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmRteAlma (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstStmRteVrncEql(t *testing.T, e, a *hst.StmRteVrnc, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmRteVrnc.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmRteVrnc.Vals"))
	HstSideEql(t, e.Side, a.Side, append(msgs, "StmRteVrnc.Side"))
}
func HstStmRteVrncNotZero(t *testing.T, a *hst.StmRteVrnc, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	TmesNotZero(t, a.Tmes, append(msgs, "StmRteVrnc.Tmes"))
	FltsNotZero(t, a.Vals, append(msgs, "StmRteVrnc.Vals"))
	HstSideNotZero(t, a.Side, append(msgs, "StmRteVrnc.Side"))
}
func HstStmRteVrncSliceEql(t *testing.T, e, a []*hst.StmRteVrnc, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstStmRteVrncEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmRteVrnc (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstStmRteStdEql(t *testing.T, e, a *hst.StmRteStd, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmRteStd.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmRteStd.Vals"))
	HstSideEql(t, e.Side, a.Side, append(msgs, "StmRteStd.Side"))
}
func HstStmRteStdNotZero(t *testing.T, a *hst.StmRteStd, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	TmesNotZero(t, a.Tmes, append(msgs, "StmRteStd.Tmes"))
	FltsNotZero(t, a.Vals, append(msgs, "StmRteStd.Vals"))
	HstSideNotZero(t, a.Side, append(msgs, "StmRteStd.Side"))
}
func HstStmRteStdSliceEql(t *testing.T, e, a []*hst.StmRteStd, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstStmRteStdEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmRteStd (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstStmRteRngFulEql(t *testing.T, e, a *hst.StmRteRngFul, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmRteRngFul.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmRteRngFul.Vals"))
	HstSideEql(t, e.Side, a.Side, append(msgs, "StmRteRngFul.Side"))
}
func HstStmRteRngFulNotZero(t *testing.T, a *hst.StmRteRngFul, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	TmesNotZero(t, a.Tmes, append(msgs, "StmRteRngFul.Tmes"))
	FltsNotZero(t, a.Vals, append(msgs, "StmRteRngFul.Vals"))
	HstSideNotZero(t, a.Side, append(msgs, "StmRteRngFul.Side"))
}
func HstStmRteRngFulSliceEql(t *testing.T, e, a []*hst.StmRteRngFul, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstStmRteRngFulEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmRteRngFul (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstStmRteRngLstEql(t *testing.T, e, a *hst.StmRteRngLst, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmRteRngLst.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmRteRngLst.Vals"))
	HstSideEql(t, e.Side, a.Side, append(msgs, "StmRteRngLst.Side"))
}
func HstStmRteRngLstNotZero(t *testing.T, a *hst.StmRteRngLst, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	TmesNotZero(t, a.Tmes, append(msgs, "StmRteRngLst.Tmes"))
	FltsNotZero(t, a.Vals, append(msgs, "StmRteRngLst.Vals"))
	HstSideNotZero(t, a.Side, append(msgs, "StmRteRngLst.Side"))
}
func HstStmRteRngLstSliceEql(t *testing.T, e, a []*hst.StmRteRngLst, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstStmRteRngLstEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmRteRngLst (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstStmRteProLstEql(t *testing.T, e, a *hst.StmRteProLst, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmRteProLst.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmRteProLst.Vals"))
	HstSideEql(t, e.Side, a.Side, append(msgs, "StmRteProLst.Side"))
}
func HstStmRteProLstNotZero(t *testing.T, a *hst.StmRteProLst, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	TmesNotZero(t, a.Tmes, append(msgs, "StmRteProLst.Tmes"))
	FltsNotZero(t, a.Vals, append(msgs, "StmRteProLst.Vals"))
	HstSideNotZero(t, a.Side, append(msgs, "StmRteProLst.Side"))
}
func HstStmRteProLstSliceEql(t *testing.T, e, a []*hst.StmRteProLst, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstStmRteProLstEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmRteProLst (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstStmRteProSmaEql(t *testing.T, e, a *hst.StmRteProSma, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmRteProSma.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmRteProSma.Vals"))
	HstSideEql(t, e.Side, a.Side, append(msgs, "StmRteProSma.Side"))
}
func HstStmRteProSmaNotZero(t *testing.T, a *hst.StmRteProSma, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	TmesNotZero(t, a.Tmes, append(msgs, "StmRteProSma.Tmes"))
	FltsNotZero(t, a.Vals, append(msgs, "StmRteProSma.Vals"))
	HstSideNotZero(t, a.Side, append(msgs, "StmRteProSma.Side"))
}
func HstStmRteProSmaSliceEql(t *testing.T, e, a []*hst.StmRteProSma, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstStmRteProSmaEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmRteProSma (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstStmRteProAlmaEql(t *testing.T, e, a *hst.StmRteProAlma, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmRteProAlma.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmRteProAlma.Vals"))
	HstSideEql(t, e.Side, a.Side, append(msgs, "StmRteProAlma.Side"))
}
func HstStmRteProAlmaNotZero(t *testing.T, a *hst.StmRteProAlma, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	TmesNotZero(t, a.Tmes, append(msgs, "StmRteProAlma.Tmes"))
	FltsNotZero(t, a.Vals, append(msgs, "StmRteProAlma.Vals"))
	HstSideNotZero(t, a.Side, append(msgs, "StmRteProAlma.Side"))
}
func HstStmRteProAlmaSliceEql(t *testing.T, e, a []*hst.StmRteProAlma, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstStmRteProAlmaEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmRteProAlma (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstStmRte1SarEql(t *testing.T, e, a *hst.StmRte1Sar, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmRte1Sar.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmRte1Sar.Vals"))
	HstSideEql(t, e.Side, a.Side, append(msgs, "StmRte1Sar.Side"))
	FltEql(t, e.AfInc, a.AfInc, append(msgs, "StmRte1Sar.AfInc"))
	FltEql(t, e.AfMax, a.AfMax, append(msgs, "StmRte1Sar.AfMax"))
}
func HstStmRte1SarNotZero(t *testing.T, a *hst.StmRte1Sar, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	TmesNotZero(t, a.Tmes, append(msgs, "StmRte1Sar.Tmes"))
	FltsNotZero(t, a.Vals, append(msgs, "StmRte1Sar.Vals"))
	HstSideNotZero(t, a.Side, append(msgs, "StmRte1Sar.Side"))
}
func HstStmRte1SarSliceEql(t *testing.T, e, a []*hst.StmRte1Sar, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstStmRte1SarEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmRte1Sar (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstStmRteEmaEql(t *testing.T, e, a *hst.StmRteEma, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmRteEma.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmRteEma.Vals"))
	HstSideEql(t, e.Side, a.Side, append(msgs, "StmRteEma.Side"))
}
func HstStmRteEmaNotZero(t *testing.T, a *hst.StmRteEma, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	TmesNotZero(t, a.Tmes, append(msgs, "StmRteEma.Tmes"))
	FltsNotZero(t, a.Vals, append(msgs, "StmRteEma.Vals"))
	HstSideNotZero(t, a.Side, append(msgs, "StmRteEma.Side"))
}
func HstStmRteEmaSliceEql(t *testing.T, e, a []*hst.StmRteEma, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstStmRteEmaEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmRteEma (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstStmUnaPosEql(t *testing.T, e, a *hst.StmUnaPos, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmUnaPos.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmUnaPos.Vals"))
	HstStmEql(t, e.Stm, a.Stm, append(msgs, "StmUnaPos.Stm"))
}
func HstStmUnaPosNotZero(t *testing.T, a *hst.StmUnaPos, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	TmesNotZero(t, a.Tmes, append(msgs, "StmUnaPos.Tmes"))
	FltsNotZero(t, a.Vals, append(msgs, "StmUnaPos.Vals"))
	HstStmNotZero(t, a.Stm, append(msgs, "StmUnaPos.Stm"))
}
func HstStmUnaPosSliceEql(t *testing.T, e, a []*hst.StmUnaPos, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstStmUnaPosEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmUnaPos (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstStmUnaNegEql(t *testing.T, e, a *hst.StmUnaNeg, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmUnaNeg.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmUnaNeg.Vals"))
	HstStmEql(t, e.Stm, a.Stm, append(msgs, "StmUnaNeg.Stm"))
}
func HstStmUnaNegNotZero(t *testing.T, a *hst.StmUnaNeg, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	TmesNotZero(t, a.Tmes, append(msgs, "StmUnaNeg.Tmes"))
	FltsNotZero(t, a.Vals, append(msgs, "StmUnaNeg.Vals"))
	HstStmNotZero(t, a.Stm, append(msgs, "StmUnaNeg.Stm"))
}
func HstStmUnaNegSliceEql(t *testing.T, e, a []*hst.StmUnaNeg, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstStmUnaNegEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmUnaNeg (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstStmUnaInvEql(t *testing.T, e, a *hst.StmUnaInv, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmUnaInv.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmUnaInv.Vals"))
	HstStmEql(t, e.Stm, a.Stm, append(msgs, "StmUnaInv.Stm"))
}
func HstStmUnaInvNotZero(t *testing.T, a *hst.StmUnaInv, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	TmesNotZero(t, a.Tmes, append(msgs, "StmUnaInv.Tmes"))
	FltsNotZero(t, a.Vals, append(msgs, "StmUnaInv.Vals"))
	HstStmNotZero(t, a.Stm, append(msgs, "StmUnaInv.Stm"))
}
func HstStmUnaInvSliceEql(t *testing.T, e, a []*hst.StmUnaInv, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstStmUnaInvEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmUnaInv (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstStmUnaSqrEql(t *testing.T, e, a *hst.StmUnaSqr, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmUnaSqr.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmUnaSqr.Vals"))
	HstStmEql(t, e.Stm, a.Stm, append(msgs, "StmUnaSqr.Stm"))
}
func HstStmUnaSqrNotZero(t *testing.T, a *hst.StmUnaSqr, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	TmesNotZero(t, a.Tmes, append(msgs, "StmUnaSqr.Tmes"))
	FltsNotZero(t, a.Vals, append(msgs, "StmUnaSqr.Vals"))
	HstStmNotZero(t, a.Stm, append(msgs, "StmUnaSqr.Stm"))
}
func HstStmUnaSqrSliceEql(t *testing.T, e, a []*hst.StmUnaSqr, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstStmUnaSqrEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmUnaSqr (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstStmUnaSqrtEql(t *testing.T, e, a *hst.StmUnaSqrt, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmUnaSqrt.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmUnaSqrt.Vals"))
	HstStmEql(t, e.Stm, a.Stm, append(msgs, "StmUnaSqrt.Stm"))
}
func HstStmUnaSqrtNotZero(t *testing.T, a *hst.StmUnaSqrt, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	TmesNotZero(t, a.Tmes, append(msgs, "StmUnaSqrt.Tmes"))
	FltsNotZero(t, a.Vals, append(msgs, "StmUnaSqrt.Vals"))
	HstStmNotZero(t, a.Stm, append(msgs, "StmUnaSqrt.Stm"))
}
func HstStmUnaSqrtSliceEql(t *testing.T, e, a []*hst.StmUnaSqrt, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstStmUnaSqrtEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmUnaSqrt (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstStmSclAddEql(t *testing.T, e, a *hst.StmSclAdd, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmSclAdd.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmSclAdd.Vals"))
	HstStmEql(t, e.Stm, a.Stm, append(msgs, "StmSclAdd.Stm"))
	FltEql(t, e.Scl, a.Scl, append(msgs, "StmSclAdd.Scl"))
}
func HstStmSclAddNotZero(t *testing.T, a *hst.StmSclAdd, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	TmesNotZero(t, a.Tmes, append(msgs, "StmSclAdd.Tmes"))
	FltsNotZero(t, a.Vals, append(msgs, "StmSclAdd.Vals"))
	HstStmNotZero(t, a.Stm, append(msgs, "StmSclAdd.Stm"))
}
func HstStmSclAddSliceEql(t *testing.T, e, a []*hst.StmSclAdd, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstStmSclAddEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmSclAdd (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstStmSclSubEql(t *testing.T, e, a *hst.StmSclSub, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmSclSub.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmSclSub.Vals"))
	HstStmEql(t, e.Stm, a.Stm, append(msgs, "StmSclSub.Stm"))
	FltEql(t, e.Scl, a.Scl, append(msgs, "StmSclSub.Scl"))
}
func HstStmSclSubNotZero(t *testing.T, a *hst.StmSclSub, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	TmesNotZero(t, a.Tmes, append(msgs, "StmSclSub.Tmes"))
	FltsNotZero(t, a.Vals, append(msgs, "StmSclSub.Vals"))
	HstStmNotZero(t, a.Stm, append(msgs, "StmSclSub.Stm"))
}
func HstStmSclSubSliceEql(t *testing.T, e, a []*hst.StmSclSub, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstStmSclSubEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmSclSub (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstStmSclMulEql(t *testing.T, e, a *hst.StmSclMul, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmSclMul.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmSclMul.Vals"))
	HstStmEql(t, e.Stm, a.Stm, append(msgs, "StmSclMul.Stm"))
	FltEql(t, e.Scl, a.Scl, append(msgs, "StmSclMul.Scl"))
}
func HstStmSclMulNotZero(t *testing.T, a *hst.StmSclMul, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	TmesNotZero(t, a.Tmes, append(msgs, "StmSclMul.Tmes"))
	FltsNotZero(t, a.Vals, append(msgs, "StmSclMul.Vals"))
	HstStmNotZero(t, a.Stm, append(msgs, "StmSclMul.Stm"))
}
func HstStmSclMulSliceEql(t *testing.T, e, a []*hst.StmSclMul, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstStmSclMulEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmSclMul (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstStmSclDivEql(t *testing.T, e, a *hst.StmSclDiv, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmSclDiv.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmSclDiv.Vals"))
	HstStmEql(t, e.Stm, a.Stm, append(msgs, "StmSclDiv.Stm"))
	FltEql(t, e.Scl, a.Scl, append(msgs, "StmSclDiv.Scl"))
}
func HstStmSclDivNotZero(t *testing.T, a *hst.StmSclDiv, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	TmesNotZero(t, a.Tmes, append(msgs, "StmSclDiv.Tmes"))
	FltsNotZero(t, a.Vals, append(msgs, "StmSclDiv.Vals"))
	HstStmNotZero(t, a.Stm, append(msgs, "StmSclDiv.Stm"))
}
func HstStmSclDivSliceEql(t *testing.T, e, a []*hst.StmSclDiv, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstStmSclDivEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmSclDiv (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstStmSclRemEql(t *testing.T, e, a *hst.StmSclRem, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmSclRem.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmSclRem.Vals"))
	HstStmEql(t, e.Stm, a.Stm, append(msgs, "StmSclRem.Stm"))
	FltEql(t, e.Scl, a.Scl, append(msgs, "StmSclRem.Scl"))
}
func HstStmSclRemNotZero(t *testing.T, a *hst.StmSclRem, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	TmesNotZero(t, a.Tmes, append(msgs, "StmSclRem.Tmes"))
	FltsNotZero(t, a.Vals, append(msgs, "StmSclRem.Vals"))
	HstStmNotZero(t, a.Stm, append(msgs, "StmSclRem.Stm"))
}
func HstStmSclRemSliceEql(t *testing.T, e, a []*hst.StmSclRem, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstStmSclRemEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmSclRem (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstStmSclPowEql(t *testing.T, e, a *hst.StmSclPow, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmSclPow.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmSclPow.Vals"))
	HstStmEql(t, e.Stm, a.Stm, append(msgs, "StmSclPow.Stm"))
	FltEql(t, e.Scl, a.Scl, append(msgs, "StmSclPow.Scl"))
}
func HstStmSclPowNotZero(t *testing.T, a *hst.StmSclPow, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	TmesNotZero(t, a.Tmes, append(msgs, "StmSclPow.Tmes"))
	FltsNotZero(t, a.Vals, append(msgs, "StmSclPow.Vals"))
	HstStmNotZero(t, a.Stm, append(msgs, "StmSclPow.Stm"))
}
func HstStmSclPowSliceEql(t *testing.T, e, a []*hst.StmSclPow, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstStmSclPowEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmSclPow (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstStmSclMinEql(t *testing.T, e, a *hst.StmSclMin, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmSclMin.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmSclMin.Vals"))
	HstStmEql(t, e.Stm, a.Stm, append(msgs, "StmSclMin.Stm"))
	FltEql(t, e.Scl, a.Scl, append(msgs, "StmSclMin.Scl"))
}
func HstStmSclMinNotZero(t *testing.T, a *hst.StmSclMin, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	TmesNotZero(t, a.Tmes, append(msgs, "StmSclMin.Tmes"))
	FltsNotZero(t, a.Vals, append(msgs, "StmSclMin.Vals"))
	HstStmNotZero(t, a.Stm, append(msgs, "StmSclMin.Stm"))
}
func HstStmSclMinSliceEql(t *testing.T, e, a []*hst.StmSclMin, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstStmSclMinEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmSclMin (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstStmSclMaxEql(t *testing.T, e, a *hst.StmSclMax, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmSclMax.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmSclMax.Vals"))
	HstStmEql(t, e.Stm, a.Stm, append(msgs, "StmSclMax.Stm"))
	FltEql(t, e.Scl, a.Scl, append(msgs, "StmSclMax.Scl"))
}
func HstStmSclMaxNotZero(t *testing.T, a *hst.StmSclMax, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	TmesNotZero(t, a.Tmes, append(msgs, "StmSclMax.Tmes"))
	FltsNotZero(t, a.Vals, append(msgs, "StmSclMax.Vals"))
	HstStmNotZero(t, a.Stm, append(msgs, "StmSclMax.Stm"))
}
func HstStmSclMaxSliceEql(t *testing.T, e, a []*hst.StmSclMax, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstStmSclMaxEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmSclMax (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstStmSelEqlEql(t *testing.T, e, a *hst.StmSelEql, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmSelEql.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmSelEql.Vals"))
	HstStmEql(t, e.Stm, a.Stm, append(msgs, "StmSelEql.Stm"))
	FltEql(t, e.Sel, a.Sel, append(msgs, "StmSelEql.Sel"))
}
func HstStmSelEqlNotZero(t *testing.T, a *hst.StmSelEql, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	TmesNotZero(t, a.Tmes, append(msgs, "StmSelEql.Tmes"))
	FltsNotZero(t, a.Vals, append(msgs, "StmSelEql.Vals"))
	HstStmNotZero(t, a.Stm, append(msgs, "StmSelEql.Stm"))
}
func HstStmSelEqlSliceEql(t *testing.T, e, a []*hst.StmSelEql, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstStmSelEqlEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmSelEql (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstStmSelNeqEql(t *testing.T, e, a *hst.StmSelNeq, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmSelNeq.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmSelNeq.Vals"))
	HstStmEql(t, e.Stm, a.Stm, append(msgs, "StmSelNeq.Stm"))
	FltEql(t, e.Sel, a.Sel, append(msgs, "StmSelNeq.Sel"))
}
func HstStmSelNeqNotZero(t *testing.T, a *hst.StmSelNeq, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	TmesNotZero(t, a.Tmes, append(msgs, "StmSelNeq.Tmes"))
	FltsNotZero(t, a.Vals, append(msgs, "StmSelNeq.Vals"))
	HstStmNotZero(t, a.Stm, append(msgs, "StmSelNeq.Stm"))
}
func HstStmSelNeqSliceEql(t *testing.T, e, a []*hst.StmSelNeq, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstStmSelNeqEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmSelNeq (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstStmSelLssEql(t *testing.T, e, a *hst.StmSelLss, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmSelLss.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmSelLss.Vals"))
	HstStmEql(t, e.Stm, a.Stm, append(msgs, "StmSelLss.Stm"))
	FltEql(t, e.Sel, a.Sel, append(msgs, "StmSelLss.Sel"))
}
func HstStmSelLssNotZero(t *testing.T, a *hst.StmSelLss, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	TmesNotZero(t, a.Tmes, append(msgs, "StmSelLss.Tmes"))
	FltsNotZero(t, a.Vals, append(msgs, "StmSelLss.Vals"))
	HstStmNotZero(t, a.Stm, append(msgs, "StmSelLss.Stm"))
}
func HstStmSelLssSliceEql(t *testing.T, e, a []*hst.StmSelLss, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstStmSelLssEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmSelLss (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstStmSelGtrEql(t *testing.T, e, a *hst.StmSelGtr, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmSelGtr.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmSelGtr.Vals"))
	HstStmEql(t, e.Stm, a.Stm, append(msgs, "StmSelGtr.Stm"))
	FltEql(t, e.Sel, a.Sel, append(msgs, "StmSelGtr.Sel"))
}
func HstStmSelGtrNotZero(t *testing.T, a *hst.StmSelGtr, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	TmesNotZero(t, a.Tmes, append(msgs, "StmSelGtr.Tmes"))
	FltsNotZero(t, a.Vals, append(msgs, "StmSelGtr.Vals"))
	HstStmNotZero(t, a.Stm, append(msgs, "StmSelGtr.Stm"))
}
func HstStmSelGtrSliceEql(t *testing.T, e, a []*hst.StmSelGtr, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstStmSelGtrEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmSelGtr (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstStmSelLeqEql(t *testing.T, e, a *hst.StmSelLeq, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmSelLeq.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmSelLeq.Vals"))
	HstStmEql(t, e.Stm, a.Stm, append(msgs, "StmSelLeq.Stm"))
	FltEql(t, e.Sel, a.Sel, append(msgs, "StmSelLeq.Sel"))
}
func HstStmSelLeqNotZero(t *testing.T, a *hst.StmSelLeq, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	TmesNotZero(t, a.Tmes, append(msgs, "StmSelLeq.Tmes"))
	FltsNotZero(t, a.Vals, append(msgs, "StmSelLeq.Vals"))
	HstStmNotZero(t, a.Stm, append(msgs, "StmSelLeq.Stm"))
}
func HstStmSelLeqSliceEql(t *testing.T, e, a []*hst.StmSelLeq, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstStmSelLeqEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmSelLeq (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstStmSelGeqEql(t *testing.T, e, a *hst.StmSelGeq, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmSelGeq.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmSelGeq.Vals"))
	HstStmEql(t, e.Stm, a.Stm, append(msgs, "StmSelGeq.Stm"))
	FltEql(t, e.Sel, a.Sel, append(msgs, "StmSelGeq.Sel"))
}
func HstStmSelGeqNotZero(t *testing.T, a *hst.StmSelGeq, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	TmesNotZero(t, a.Tmes, append(msgs, "StmSelGeq.Tmes"))
	FltsNotZero(t, a.Vals, append(msgs, "StmSelGeq.Vals"))
	HstStmNotZero(t, a.Stm, append(msgs, "StmSelGeq.Stm"))
}
func HstStmSelGeqSliceEql(t *testing.T, e, a []*hst.StmSelGeq, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstStmSelGeqEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmSelGeq (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstStmAggFstEql(t *testing.T, e, a *hst.StmAggFst, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmAggFst.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmAggFst.Vals"))
	HstStmEql(t, e.Stm, a.Stm, append(msgs, "StmAggFst.Stm"))
	UntEql(t, e.Length, a.Length, append(msgs, "StmAggFst.Length"))
}
func HstStmAggFstNotZero(t *testing.T, a *hst.StmAggFst, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	TmesNotZero(t, a.Tmes, append(msgs, "StmAggFst.Tmes"))
	FltsNotZero(t, a.Vals, append(msgs, "StmAggFst.Vals"))
	HstStmNotZero(t, a.Stm, append(msgs, "StmAggFst.Stm"))
}
func HstStmAggFstSliceEql(t *testing.T, e, a []*hst.StmAggFst, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstStmAggFstEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmAggFst (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstStmAggLstEql(t *testing.T, e, a *hst.StmAggLst, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmAggLst.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmAggLst.Vals"))
	HstStmEql(t, e.Stm, a.Stm, append(msgs, "StmAggLst.Stm"))
	UntEql(t, e.Length, a.Length, append(msgs, "StmAggLst.Length"))
}
func HstStmAggLstNotZero(t *testing.T, a *hst.StmAggLst, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	TmesNotZero(t, a.Tmes, append(msgs, "StmAggLst.Tmes"))
	FltsNotZero(t, a.Vals, append(msgs, "StmAggLst.Vals"))
	HstStmNotZero(t, a.Stm, append(msgs, "StmAggLst.Stm"))
}
func HstStmAggLstSliceEql(t *testing.T, e, a []*hst.StmAggLst, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstStmAggLstEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmAggLst (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstStmAggSumEql(t *testing.T, e, a *hst.StmAggSum, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmAggSum.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmAggSum.Vals"))
	HstStmEql(t, e.Stm, a.Stm, append(msgs, "StmAggSum.Stm"))
	UntEql(t, e.Length, a.Length, append(msgs, "StmAggSum.Length"))
}
func HstStmAggSumNotZero(t *testing.T, a *hst.StmAggSum, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	TmesNotZero(t, a.Tmes, append(msgs, "StmAggSum.Tmes"))
	FltsNotZero(t, a.Vals, append(msgs, "StmAggSum.Vals"))
	HstStmNotZero(t, a.Stm, append(msgs, "StmAggSum.Stm"))
}
func HstStmAggSumSliceEql(t *testing.T, e, a []*hst.StmAggSum, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstStmAggSumEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmAggSum (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstStmAggPrdEql(t *testing.T, e, a *hst.StmAggPrd, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmAggPrd.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmAggPrd.Vals"))
	HstStmEql(t, e.Stm, a.Stm, append(msgs, "StmAggPrd.Stm"))
	UntEql(t, e.Length, a.Length, append(msgs, "StmAggPrd.Length"))
}
func HstStmAggPrdNotZero(t *testing.T, a *hst.StmAggPrd, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	TmesNotZero(t, a.Tmes, append(msgs, "StmAggPrd.Tmes"))
	FltsNotZero(t, a.Vals, append(msgs, "StmAggPrd.Vals"))
	HstStmNotZero(t, a.Stm, append(msgs, "StmAggPrd.Stm"))
}
func HstStmAggPrdSliceEql(t *testing.T, e, a []*hst.StmAggPrd, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstStmAggPrdEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmAggPrd (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstStmAggMinEql(t *testing.T, e, a *hst.StmAggMin, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmAggMin.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmAggMin.Vals"))
	HstStmEql(t, e.Stm, a.Stm, append(msgs, "StmAggMin.Stm"))
	UntEql(t, e.Length, a.Length, append(msgs, "StmAggMin.Length"))
}
func HstStmAggMinNotZero(t *testing.T, a *hst.StmAggMin, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	TmesNotZero(t, a.Tmes, append(msgs, "StmAggMin.Tmes"))
	FltsNotZero(t, a.Vals, append(msgs, "StmAggMin.Vals"))
	HstStmNotZero(t, a.Stm, append(msgs, "StmAggMin.Stm"))
}
func HstStmAggMinSliceEql(t *testing.T, e, a []*hst.StmAggMin, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstStmAggMinEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmAggMin (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstStmAggMaxEql(t *testing.T, e, a *hst.StmAggMax, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmAggMax.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmAggMax.Vals"))
	HstStmEql(t, e.Stm, a.Stm, append(msgs, "StmAggMax.Stm"))
	UntEql(t, e.Length, a.Length, append(msgs, "StmAggMax.Length"))
}
func HstStmAggMaxNotZero(t *testing.T, a *hst.StmAggMax, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	TmesNotZero(t, a.Tmes, append(msgs, "StmAggMax.Tmes"))
	FltsNotZero(t, a.Vals, append(msgs, "StmAggMax.Vals"))
	HstStmNotZero(t, a.Stm, append(msgs, "StmAggMax.Stm"))
}
func HstStmAggMaxSliceEql(t *testing.T, e, a []*hst.StmAggMax, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstStmAggMaxEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmAggMax (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstStmAggMidEql(t *testing.T, e, a *hst.StmAggMid, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmAggMid.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmAggMid.Vals"))
	HstStmEql(t, e.Stm, a.Stm, append(msgs, "StmAggMid.Stm"))
	UntEql(t, e.Length, a.Length, append(msgs, "StmAggMid.Length"))
}
func HstStmAggMidNotZero(t *testing.T, a *hst.StmAggMid, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	TmesNotZero(t, a.Tmes, append(msgs, "StmAggMid.Tmes"))
	FltsNotZero(t, a.Vals, append(msgs, "StmAggMid.Vals"))
	HstStmNotZero(t, a.Stm, append(msgs, "StmAggMid.Stm"))
}
func HstStmAggMidSliceEql(t *testing.T, e, a []*hst.StmAggMid, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstStmAggMidEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmAggMid (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstStmAggMdnEql(t *testing.T, e, a *hst.StmAggMdn, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmAggMdn.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmAggMdn.Vals"))
	HstStmEql(t, e.Stm, a.Stm, append(msgs, "StmAggMdn.Stm"))
	UntEql(t, e.Length, a.Length, append(msgs, "StmAggMdn.Length"))
}
func HstStmAggMdnNotZero(t *testing.T, a *hst.StmAggMdn, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	TmesNotZero(t, a.Tmes, append(msgs, "StmAggMdn.Tmes"))
	FltsNotZero(t, a.Vals, append(msgs, "StmAggMdn.Vals"))
	HstStmNotZero(t, a.Stm, append(msgs, "StmAggMdn.Stm"))
}
func HstStmAggMdnSliceEql(t *testing.T, e, a []*hst.StmAggMdn, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstStmAggMdnEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmAggMdn (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstStmAggSmaEql(t *testing.T, e, a *hst.StmAggSma, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmAggSma.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmAggSma.Vals"))
	HstStmEql(t, e.Stm, a.Stm, append(msgs, "StmAggSma.Stm"))
	UntEql(t, e.Length, a.Length, append(msgs, "StmAggSma.Length"))
}
func HstStmAggSmaNotZero(t *testing.T, a *hst.StmAggSma, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	TmesNotZero(t, a.Tmes, append(msgs, "StmAggSma.Tmes"))
	FltsNotZero(t, a.Vals, append(msgs, "StmAggSma.Vals"))
	HstStmNotZero(t, a.Stm, append(msgs, "StmAggSma.Stm"))
}
func HstStmAggSmaSliceEql(t *testing.T, e, a []*hst.StmAggSma, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstStmAggSmaEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmAggSma (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstStmAggGmaEql(t *testing.T, e, a *hst.StmAggGma, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmAggGma.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmAggGma.Vals"))
	HstStmEql(t, e.Stm, a.Stm, append(msgs, "StmAggGma.Stm"))
	UntEql(t, e.Length, a.Length, append(msgs, "StmAggGma.Length"))
}
func HstStmAggGmaNotZero(t *testing.T, a *hst.StmAggGma, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	TmesNotZero(t, a.Tmes, append(msgs, "StmAggGma.Tmes"))
	FltsNotZero(t, a.Vals, append(msgs, "StmAggGma.Vals"))
	HstStmNotZero(t, a.Stm, append(msgs, "StmAggGma.Stm"))
}
func HstStmAggGmaSliceEql(t *testing.T, e, a []*hst.StmAggGma, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstStmAggGmaEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmAggGma (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstStmAggWmaEql(t *testing.T, e, a *hst.StmAggWma, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmAggWma.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmAggWma.Vals"))
	HstStmEql(t, e.Stm, a.Stm, append(msgs, "StmAggWma.Stm"))
	UntEql(t, e.Length, a.Length, append(msgs, "StmAggWma.Length"))
}
func HstStmAggWmaNotZero(t *testing.T, a *hst.StmAggWma, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	TmesNotZero(t, a.Tmes, append(msgs, "StmAggWma.Tmes"))
	FltsNotZero(t, a.Vals, append(msgs, "StmAggWma.Vals"))
	HstStmNotZero(t, a.Stm, append(msgs, "StmAggWma.Stm"))
}
func HstStmAggWmaSliceEql(t *testing.T, e, a []*hst.StmAggWma, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstStmAggWmaEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmAggWma (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstStmAggRsiEql(t *testing.T, e, a *hst.StmAggRsi, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmAggRsi.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmAggRsi.Vals"))
	HstStmEql(t, e.Stm, a.Stm, append(msgs, "StmAggRsi.Stm"))
	UntEql(t, e.Length, a.Length, append(msgs, "StmAggRsi.Length"))
}
func HstStmAggRsiNotZero(t *testing.T, a *hst.StmAggRsi, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	TmesNotZero(t, a.Tmes, append(msgs, "StmAggRsi.Tmes"))
	FltsNotZero(t, a.Vals, append(msgs, "StmAggRsi.Vals"))
	HstStmNotZero(t, a.Stm, append(msgs, "StmAggRsi.Stm"))
}
func HstStmAggRsiSliceEql(t *testing.T, e, a []*hst.StmAggRsi, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstStmAggRsiEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmAggRsi (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstStmAggWrsiEql(t *testing.T, e, a *hst.StmAggWrsi, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmAggWrsi.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmAggWrsi.Vals"))
	HstStmEql(t, e.Stm, a.Stm, append(msgs, "StmAggWrsi.Stm"))
	UntEql(t, e.Length, a.Length, append(msgs, "StmAggWrsi.Length"))
}
func HstStmAggWrsiNotZero(t *testing.T, a *hst.StmAggWrsi, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	TmesNotZero(t, a.Tmes, append(msgs, "StmAggWrsi.Tmes"))
	FltsNotZero(t, a.Vals, append(msgs, "StmAggWrsi.Vals"))
	HstStmNotZero(t, a.Stm, append(msgs, "StmAggWrsi.Stm"))
}
func HstStmAggWrsiSliceEql(t *testing.T, e, a []*hst.StmAggWrsi, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstStmAggWrsiEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmAggWrsi (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstStmAggAlmaEql(t *testing.T, e, a *hst.StmAggAlma, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmAggAlma.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmAggAlma.Vals"))
	HstStmEql(t, e.Stm, a.Stm, append(msgs, "StmAggAlma.Stm"))
	UntEql(t, e.Length, a.Length, append(msgs, "StmAggAlma.Length"))
}
func HstStmAggAlmaNotZero(t *testing.T, a *hst.StmAggAlma, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	TmesNotZero(t, a.Tmes, append(msgs, "StmAggAlma.Tmes"))
	FltsNotZero(t, a.Vals, append(msgs, "StmAggAlma.Vals"))
	HstStmNotZero(t, a.Stm, append(msgs, "StmAggAlma.Stm"))
}
func HstStmAggAlmaSliceEql(t *testing.T, e, a []*hst.StmAggAlma, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstStmAggAlmaEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmAggAlma (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstStmAggVrncEql(t *testing.T, e, a *hst.StmAggVrnc, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmAggVrnc.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmAggVrnc.Vals"))
	HstStmEql(t, e.Stm, a.Stm, append(msgs, "StmAggVrnc.Stm"))
	UntEql(t, e.Length, a.Length, append(msgs, "StmAggVrnc.Length"))
}
func HstStmAggVrncNotZero(t *testing.T, a *hst.StmAggVrnc, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	TmesNotZero(t, a.Tmes, append(msgs, "StmAggVrnc.Tmes"))
	FltsNotZero(t, a.Vals, append(msgs, "StmAggVrnc.Vals"))
	HstStmNotZero(t, a.Stm, append(msgs, "StmAggVrnc.Stm"))
}
func HstStmAggVrncSliceEql(t *testing.T, e, a []*hst.StmAggVrnc, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstStmAggVrncEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmAggVrnc (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstStmAggStdEql(t *testing.T, e, a *hst.StmAggStd, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmAggStd.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmAggStd.Vals"))
	HstStmEql(t, e.Stm, a.Stm, append(msgs, "StmAggStd.Stm"))
	UntEql(t, e.Length, a.Length, append(msgs, "StmAggStd.Length"))
}
func HstStmAggStdNotZero(t *testing.T, a *hst.StmAggStd, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	TmesNotZero(t, a.Tmes, append(msgs, "StmAggStd.Tmes"))
	FltsNotZero(t, a.Vals, append(msgs, "StmAggStd.Vals"))
	HstStmNotZero(t, a.Stm, append(msgs, "StmAggStd.Stm"))
}
func HstStmAggStdSliceEql(t *testing.T, e, a []*hst.StmAggStd, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstStmAggStdEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmAggStd (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstStmAggRngFulEql(t *testing.T, e, a *hst.StmAggRngFul, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmAggRngFul.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmAggRngFul.Vals"))
	HstStmEql(t, e.Stm, a.Stm, append(msgs, "StmAggRngFul.Stm"))
	UntEql(t, e.Length, a.Length, append(msgs, "StmAggRngFul.Length"))
}
func HstStmAggRngFulNotZero(t *testing.T, a *hst.StmAggRngFul, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	TmesNotZero(t, a.Tmes, append(msgs, "StmAggRngFul.Tmes"))
	FltsNotZero(t, a.Vals, append(msgs, "StmAggRngFul.Vals"))
	HstStmNotZero(t, a.Stm, append(msgs, "StmAggRngFul.Stm"))
}
func HstStmAggRngFulSliceEql(t *testing.T, e, a []*hst.StmAggRngFul, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstStmAggRngFulEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmAggRngFul (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstStmAggRngLstEql(t *testing.T, e, a *hst.StmAggRngLst, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmAggRngLst.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmAggRngLst.Vals"))
	HstStmEql(t, e.Stm, a.Stm, append(msgs, "StmAggRngLst.Stm"))
	UntEql(t, e.Length, a.Length, append(msgs, "StmAggRngLst.Length"))
}
func HstStmAggRngLstNotZero(t *testing.T, a *hst.StmAggRngLst, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	TmesNotZero(t, a.Tmes, append(msgs, "StmAggRngLst.Tmes"))
	FltsNotZero(t, a.Vals, append(msgs, "StmAggRngLst.Vals"))
	HstStmNotZero(t, a.Stm, append(msgs, "StmAggRngLst.Stm"))
}
func HstStmAggRngLstSliceEql(t *testing.T, e, a []*hst.StmAggRngLst, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstStmAggRngLstEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmAggRngLst (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstStmAggProLstEql(t *testing.T, e, a *hst.StmAggProLst, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmAggProLst.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmAggProLst.Vals"))
	HstStmEql(t, e.Stm, a.Stm, append(msgs, "StmAggProLst.Stm"))
	UntEql(t, e.Length, a.Length, append(msgs, "StmAggProLst.Length"))
}
func HstStmAggProLstNotZero(t *testing.T, a *hst.StmAggProLst, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	TmesNotZero(t, a.Tmes, append(msgs, "StmAggProLst.Tmes"))
	FltsNotZero(t, a.Vals, append(msgs, "StmAggProLst.Vals"))
	HstStmNotZero(t, a.Stm, append(msgs, "StmAggProLst.Stm"))
}
func HstStmAggProLstSliceEql(t *testing.T, e, a []*hst.StmAggProLst, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstStmAggProLstEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmAggProLst (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstStmAggProSmaEql(t *testing.T, e, a *hst.StmAggProSma, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmAggProSma.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmAggProSma.Vals"))
	HstStmEql(t, e.Stm, a.Stm, append(msgs, "StmAggProSma.Stm"))
	UntEql(t, e.Length, a.Length, append(msgs, "StmAggProSma.Length"))
}
func HstStmAggProSmaNotZero(t *testing.T, a *hst.StmAggProSma, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	TmesNotZero(t, a.Tmes, append(msgs, "StmAggProSma.Tmes"))
	FltsNotZero(t, a.Vals, append(msgs, "StmAggProSma.Vals"))
	HstStmNotZero(t, a.Stm, append(msgs, "StmAggProSma.Stm"))
}
func HstStmAggProSmaSliceEql(t *testing.T, e, a []*hst.StmAggProSma, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstStmAggProSmaEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmAggProSma (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstStmAggProAlmaEql(t *testing.T, e, a *hst.StmAggProAlma, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmAggProAlma.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmAggProAlma.Vals"))
	HstStmEql(t, e.Stm, a.Stm, append(msgs, "StmAggProAlma.Stm"))
	UntEql(t, e.Length, a.Length, append(msgs, "StmAggProAlma.Length"))
}
func HstStmAggProAlmaNotZero(t *testing.T, a *hst.StmAggProAlma, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	TmesNotZero(t, a.Tmes, append(msgs, "StmAggProAlma.Tmes"))
	FltsNotZero(t, a.Vals, append(msgs, "StmAggProAlma.Vals"))
	HstStmNotZero(t, a.Stm, append(msgs, "StmAggProAlma.Stm"))
}
func HstStmAggProAlmaSliceEql(t *testing.T, e, a []*hst.StmAggProAlma, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstStmAggProAlmaEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmAggProAlma (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstStmAggEmaEql(t *testing.T, e, a *hst.StmAggEma, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmAggEma.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmAggEma.Vals"))
	HstStmEql(t, e.Stm, a.Stm, append(msgs, "StmAggEma.Stm"))
	UntEql(t, e.Length, a.Length, append(msgs, "StmAggEma.Length"))
}
func HstStmAggEmaNotZero(t *testing.T, a *hst.StmAggEma, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	TmesNotZero(t, a.Tmes, append(msgs, "StmAggEma.Tmes"))
	FltsNotZero(t, a.Vals, append(msgs, "StmAggEma.Vals"))
	HstStmNotZero(t, a.Stm, append(msgs, "StmAggEma.Stm"))
}
func HstStmAggEmaSliceEql(t *testing.T, e, a []*hst.StmAggEma, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstStmAggEmaEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmAggEma (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstStmInrAddEql(t *testing.T, e, a *hst.StmInrAdd, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmInrAdd.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmInrAdd.Vals"))
	HstStmEql(t, e.Stm, a.Stm, append(msgs, "StmInrAdd.Stm"))
	UntEql(t, e.Off, a.Off, append(msgs, "StmInrAdd.Off"))
}
func HstStmInrAddNotZero(t *testing.T, a *hst.StmInrAdd, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	TmesNotZero(t, a.Tmes, append(msgs, "StmInrAdd.Tmes"))
	FltsNotZero(t, a.Vals, append(msgs, "StmInrAdd.Vals"))
	HstStmNotZero(t, a.Stm, append(msgs, "StmInrAdd.Stm"))
}
func HstStmInrAddSliceEql(t *testing.T, e, a []*hst.StmInrAdd, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstStmInrAddEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmInrAdd (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstStmInrSubEql(t *testing.T, e, a *hst.StmInrSub, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmInrSub.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmInrSub.Vals"))
	HstStmEql(t, e.Stm, a.Stm, append(msgs, "StmInrSub.Stm"))
	UntEql(t, e.Off, a.Off, append(msgs, "StmInrSub.Off"))
}
func HstStmInrSubNotZero(t *testing.T, a *hst.StmInrSub, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	TmesNotZero(t, a.Tmes, append(msgs, "StmInrSub.Tmes"))
	FltsNotZero(t, a.Vals, append(msgs, "StmInrSub.Vals"))
	HstStmNotZero(t, a.Stm, append(msgs, "StmInrSub.Stm"))
}
func HstStmInrSubSliceEql(t *testing.T, e, a []*hst.StmInrSub, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstStmInrSubEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmInrSub (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstStmInrMulEql(t *testing.T, e, a *hst.StmInrMul, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmInrMul.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmInrMul.Vals"))
	HstStmEql(t, e.Stm, a.Stm, append(msgs, "StmInrMul.Stm"))
	UntEql(t, e.Off, a.Off, append(msgs, "StmInrMul.Off"))
}
func HstStmInrMulNotZero(t *testing.T, a *hst.StmInrMul, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	TmesNotZero(t, a.Tmes, append(msgs, "StmInrMul.Tmes"))
	FltsNotZero(t, a.Vals, append(msgs, "StmInrMul.Vals"))
	HstStmNotZero(t, a.Stm, append(msgs, "StmInrMul.Stm"))
}
func HstStmInrMulSliceEql(t *testing.T, e, a []*hst.StmInrMul, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstStmInrMulEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmInrMul (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstStmInrDivEql(t *testing.T, e, a *hst.StmInrDiv, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmInrDiv.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmInrDiv.Vals"))
	HstStmEql(t, e.Stm, a.Stm, append(msgs, "StmInrDiv.Stm"))
	UntEql(t, e.Off, a.Off, append(msgs, "StmInrDiv.Off"))
}
func HstStmInrDivNotZero(t *testing.T, a *hst.StmInrDiv, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	TmesNotZero(t, a.Tmes, append(msgs, "StmInrDiv.Tmes"))
	FltsNotZero(t, a.Vals, append(msgs, "StmInrDiv.Vals"))
	HstStmNotZero(t, a.Stm, append(msgs, "StmInrDiv.Stm"))
}
func HstStmInrDivSliceEql(t *testing.T, e, a []*hst.StmInrDiv, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstStmInrDivEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmInrDiv (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstStmInrRemEql(t *testing.T, e, a *hst.StmInrRem, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmInrRem.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmInrRem.Vals"))
	HstStmEql(t, e.Stm, a.Stm, append(msgs, "StmInrRem.Stm"))
	UntEql(t, e.Off, a.Off, append(msgs, "StmInrRem.Off"))
}
func HstStmInrRemNotZero(t *testing.T, a *hst.StmInrRem, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	TmesNotZero(t, a.Tmes, append(msgs, "StmInrRem.Tmes"))
	FltsNotZero(t, a.Vals, append(msgs, "StmInrRem.Vals"))
	HstStmNotZero(t, a.Stm, append(msgs, "StmInrRem.Stm"))
}
func HstStmInrRemSliceEql(t *testing.T, e, a []*hst.StmInrRem, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstStmInrRemEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmInrRem (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstStmInrPowEql(t *testing.T, e, a *hst.StmInrPow, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmInrPow.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmInrPow.Vals"))
	HstStmEql(t, e.Stm, a.Stm, append(msgs, "StmInrPow.Stm"))
	UntEql(t, e.Off, a.Off, append(msgs, "StmInrPow.Off"))
}
func HstStmInrPowNotZero(t *testing.T, a *hst.StmInrPow, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	TmesNotZero(t, a.Tmes, append(msgs, "StmInrPow.Tmes"))
	FltsNotZero(t, a.Vals, append(msgs, "StmInrPow.Vals"))
	HstStmNotZero(t, a.Stm, append(msgs, "StmInrPow.Stm"))
}
func HstStmInrPowSliceEql(t *testing.T, e, a []*hst.StmInrPow, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstStmInrPowEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmInrPow (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstStmInrMinEql(t *testing.T, e, a *hst.StmInrMin, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmInrMin.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmInrMin.Vals"))
	HstStmEql(t, e.Stm, a.Stm, append(msgs, "StmInrMin.Stm"))
	UntEql(t, e.Off, a.Off, append(msgs, "StmInrMin.Off"))
}
func HstStmInrMinNotZero(t *testing.T, a *hst.StmInrMin, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	TmesNotZero(t, a.Tmes, append(msgs, "StmInrMin.Tmes"))
	FltsNotZero(t, a.Vals, append(msgs, "StmInrMin.Vals"))
	HstStmNotZero(t, a.Stm, append(msgs, "StmInrMin.Stm"))
}
func HstStmInrMinSliceEql(t *testing.T, e, a []*hst.StmInrMin, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstStmInrMinEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmInrMin (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstStmInrMaxEql(t *testing.T, e, a *hst.StmInrMax, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmInrMax.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmInrMax.Vals"))
	HstStmEql(t, e.Stm, a.Stm, append(msgs, "StmInrMax.Stm"))
	UntEql(t, e.Off, a.Off, append(msgs, "StmInrMax.Off"))
}
func HstStmInrMaxNotZero(t *testing.T, a *hst.StmInrMax, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	TmesNotZero(t, a.Tmes, append(msgs, "StmInrMax.Tmes"))
	FltsNotZero(t, a.Vals, append(msgs, "StmInrMax.Vals"))
	HstStmNotZero(t, a.Stm, append(msgs, "StmInrMax.Stm"))
}
func HstStmInrMaxSliceEql(t *testing.T, e, a []*hst.StmInrMax, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstStmInrMaxEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmInrMax (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstStmInr1SlpEql(t *testing.T, e, a *hst.StmInr1Slp, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmInr1Slp.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmInr1Slp.Vals"))
	HstStmEql(t, e.Stm, a.Stm, append(msgs, "StmInr1Slp.Stm"))
	UntEql(t, e.Off, a.Off, append(msgs, "StmInr1Slp.Off"))
}
func HstStmInr1SlpNotZero(t *testing.T, a *hst.StmInr1Slp, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	TmesNotZero(t, a.Tmes, append(msgs, "StmInr1Slp.Tmes"))
	FltsNotZero(t, a.Vals, append(msgs, "StmInr1Slp.Vals"))
	HstStmNotZero(t, a.Stm, append(msgs, "StmInr1Slp.Stm"))
}
func HstStmInr1SlpSliceEql(t *testing.T, e, a []*hst.StmInr1Slp, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstStmInr1SlpEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmInr1Slp (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstStmOtrAddEql(t *testing.T, e, a *hst.StmOtrAdd, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmOtrAdd.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmOtrAdd.Vals"))
	HstStmEql(t, e.Stm, a.Stm, append(msgs, "StmOtrAdd.Stm"))
	UntEql(t, e.Off, a.Off, append(msgs, "StmOtrAdd.Off"))
	HstStmEql(t, e.A, a.A, append(msgs, "StmOtrAdd.A"))
}
func HstStmOtrAddNotZero(t *testing.T, a *hst.StmOtrAdd, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	TmesNotZero(t, a.Tmes, append(msgs, "StmOtrAdd.Tmes"))
	FltsNotZero(t, a.Vals, append(msgs, "StmOtrAdd.Vals"))
	HstStmNotZero(t, a.Stm, append(msgs, "StmOtrAdd.Stm"))
}
func HstStmOtrAddSliceEql(t *testing.T, e, a []*hst.StmOtrAdd, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstStmOtrAddEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmOtrAdd (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstStmOtrSubEql(t *testing.T, e, a *hst.StmOtrSub, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmOtrSub.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmOtrSub.Vals"))
	HstStmEql(t, e.Stm, a.Stm, append(msgs, "StmOtrSub.Stm"))
	UntEql(t, e.Off, a.Off, append(msgs, "StmOtrSub.Off"))
	HstStmEql(t, e.A, a.A, append(msgs, "StmOtrSub.A"))
}
func HstStmOtrSubNotZero(t *testing.T, a *hst.StmOtrSub, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	TmesNotZero(t, a.Tmes, append(msgs, "StmOtrSub.Tmes"))
	FltsNotZero(t, a.Vals, append(msgs, "StmOtrSub.Vals"))
	HstStmNotZero(t, a.Stm, append(msgs, "StmOtrSub.Stm"))
}
func HstStmOtrSubSliceEql(t *testing.T, e, a []*hst.StmOtrSub, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstStmOtrSubEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmOtrSub (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstStmOtrMulEql(t *testing.T, e, a *hst.StmOtrMul, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmOtrMul.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmOtrMul.Vals"))
	HstStmEql(t, e.Stm, a.Stm, append(msgs, "StmOtrMul.Stm"))
	UntEql(t, e.Off, a.Off, append(msgs, "StmOtrMul.Off"))
	HstStmEql(t, e.A, a.A, append(msgs, "StmOtrMul.A"))
}
func HstStmOtrMulNotZero(t *testing.T, a *hst.StmOtrMul, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	TmesNotZero(t, a.Tmes, append(msgs, "StmOtrMul.Tmes"))
	FltsNotZero(t, a.Vals, append(msgs, "StmOtrMul.Vals"))
	HstStmNotZero(t, a.Stm, append(msgs, "StmOtrMul.Stm"))
}
func HstStmOtrMulSliceEql(t *testing.T, e, a []*hst.StmOtrMul, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstStmOtrMulEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmOtrMul (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstStmOtrDivEql(t *testing.T, e, a *hst.StmOtrDiv, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmOtrDiv.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmOtrDiv.Vals"))
	HstStmEql(t, e.Stm, a.Stm, append(msgs, "StmOtrDiv.Stm"))
	UntEql(t, e.Off, a.Off, append(msgs, "StmOtrDiv.Off"))
	HstStmEql(t, e.A, a.A, append(msgs, "StmOtrDiv.A"))
}
func HstStmOtrDivNotZero(t *testing.T, a *hst.StmOtrDiv, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	TmesNotZero(t, a.Tmes, append(msgs, "StmOtrDiv.Tmes"))
	FltsNotZero(t, a.Vals, append(msgs, "StmOtrDiv.Vals"))
	HstStmNotZero(t, a.Stm, append(msgs, "StmOtrDiv.Stm"))
}
func HstStmOtrDivSliceEql(t *testing.T, e, a []*hst.StmOtrDiv, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstStmOtrDivEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmOtrDiv (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstStmOtrRemEql(t *testing.T, e, a *hst.StmOtrRem, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmOtrRem.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmOtrRem.Vals"))
	HstStmEql(t, e.Stm, a.Stm, append(msgs, "StmOtrRem.Stm"))
	UntEql(t, e.Off, a.Off, append(msgs, "StmOtrRem.Off"))
	HstStmEql(t, e.A, a.A, append(msgs, "StmOtrRem.A"))
}
func HstStmOtrRemNotZero(t *testing.T, a *hst.StmOtrRem, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	TmesNotZero(t, a.Tmes, append(msgs, "StmOtrRem.Tmes"))
	FltsNotZero(t, a.Vals, append(msgs, "StmOtrRem.Vals"))
	HstStmNotZero(t, a.Stm, append(msgs, "StmOtrRem.Stm"))
}
func HstStmOtrRemSliceEql(t *testing.T, e, a []*hst.StmOtrRem, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstStmOtrRemEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmOtrRem (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstStmOtrPowEql(t *testing.T, e, a *hst.StmOtrPow, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmOtrPow.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmOtrPow.Vals"))
	HstStmEql(t, e.Stm, a.Stm, append(msgs, "StmOtrPow.Stm"))
	UntEql(t, e.Off, a.Off, append(msgs, "StmOtrPow.Off"))
	HstStmEql(t, e.A, a.A, append(msgs, "StmOtrPow.A"))
}
func HstStmOtrPowNotZero(t *testing.T, a *hst.StmOtrPow, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	TmesNotZero(t, a.Tmes, append(msgs, "StmOtrPow.Tmes"))
	FltsNotZero(t, a.Vals, append(msgs, "StmOtrPow.Vals"))
	HstStmNotZero(t, a.Stm, append(msgs, "StmOtrPow.Stm"))
}
func HstStmOtrPowSliceEql(t *testing.T, e, a []*hst.StmOtrPow, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstStmOtrPowEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmOtrPow (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstStmOtrMinEql(t *testing.T, e, a *hst.StmOtrMin, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmOtrMin.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmOtrMin.Vals"))
	HstStmEql(t, e.Stm, a.Stm, append(msgs, "StmOtrMin.Stm"))
	UntEql(t, e.Off, a.Off, append(msgs, "StmOtrMin.Off"))
	HstStmEql(t, e.A, a.A, append(msgs, "StmOtrMin.A"))
}
func HstStmOtrMinNotZero(t *testing.T, a *hst.StmOtrMin, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	TmesNotZero(t, a.Tmes, append(msgs, "StmOtrMin.Tmes"))
	FltsNotZero(t, a.Vals, append(msgs, "StmOtrMin.Vals"))
	HstStmNotZero(t, a.Stm, append(msgs, "StmOtrMin.Stm"))
}
func HstStmOtrMinSliceEql(t *testing.T, e, a []*hst.StmOtrMin, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstStmOtrMinEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmOtrMin (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstStmOtrMaxEql(t *testing.T, e, a *hst.StmOtrMax, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmOtrMax.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmOtrMax.Vals"))
	HstStmEql(t, e.Stm, a.Stm, append(msgs, "StmOtrMax.Stm"))
	UntEql(t, e.Off, a.Off, append(msgs, "StmOtrMax.Off"))
	HstStmEql(t, e.A, a.A, append(msgs, "StmOtrMax.A"))
}
func HstStmOtrMaxNotZero(t *testing.T, a *hst.StmOtrMax, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	TmesNotZero(t, a.Tmes, append(msgs, "StmOtrMax.Tmes"))
	FltsNotZero(t, a.Vals, append(msgs, "StmOtrMax.Vals"))
	HstStmNotZero(t, a.Stm, append(msgs, "StmOtrMax.Stm"))
}
func HstStmOtrMaxSliceEql(t *testing.T, e, a []*hst.StmOtrMax, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstStmOtrMaxEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmOtrMax (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstCndSclEqlEql(t *testing.T, e, a *hst.CndSclEql, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "CndSclEql.Tmes"))
	HstStmEql(t, e.Stm, a.Stm, append(msgs, "CndSclEql.Stm"))
	FltEql(t, e.Scl, a.Scl, append(msgs, "CndSclEql.Scl"))
}
func HstCndSclEqlNotZero(t *testing.T, a *hst.CndSclEql, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	HstStmNotZero(t, a.Stm, append(msgs, "CndSclEql.Stm"))
}
func HstCndSclEqlSliceEql(t *testing.T, e, a []*hst.CndSclEql, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstCndSclEqlEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm CndSclEql (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstCndSclNeqEql(t *testing.T, e, a *hst.CndSclNeq, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "CndSclNeq.Tmes"))
	HstStmEql(t, e.Stm, a.Stm, append(msgs, "CndSclNeq.Stm"))
	FltEql(t, e.Scl, a.Scl, append(msgs, "CndSclNeq.Scl"))
}
func HstCndSclNeqNotZero(t *testing.T, a *hst.CndSclNeq, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	HstStmNotZero(t, a.Stm, append(msgs, "CndSclNeq.Stm"))
}
func HstCndSclNeqSliceEql(t *testing.T, e, a []*hst.CndSclNeq, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstCndSclNeqEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm CndSclNeq (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstCndSclLssEql(t *testing.T, e, a *hst.CndSclLss, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "CndSclLss.Tmes"))
	HstStmEql(t, e.Stm, a.Stm, append(msgs, "CndSclLss.Stm"))
	FltEql(t, e.Scl, a.Scl, append(msgs, "CndSclLss.Scl"))
}
func HstCndSclLssNotZero(t *testing.T, a *hst.CndSclLss, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	HstStmNotZero(t, a.Stm, append(msgs, "CndSclLss.Stm"))
}
func HstCndSclLssSliceEql(t *testing.T, e, a []*hst.CndSclLss, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstCndSclLssEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm CndSclLss (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstCndSclGtrEql(t *testing.T, e, a *hst.CndSclGtr, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "CndSclGtr.Tmes"))
	HstStmEql(t, e.Stm, a.Stm, append(msgs, "CndSclGtr.Stm"))
	FltEql(t, e.Scl, a.Scl, append(msgs, "CndSclGtr.Scl"))
}
func HstCndSclGtrNotZero(t *testing.T, a *hst.CndSclGtr, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	HstStmNotZero(t, a.Stm, append(msgs, "CndSclGtr.Stm"))
}
func HstCndSclGtrSliceEql(t *testing.T, e, a []*hst.CndSclGtr, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstCndSclGtrEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm CndSclGtr (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstCndSclLeqEql(t *testing.T, e, a *hst.CndSclLeq, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "CndSclLeq.Tmes"))
	HstStmEql(t, e.Stm, a.Stm, append(msgs, "CndSclLeq.Stm"))
	FltEql(t, e.Scl, a.Scl, append(msgs, "CndSclLeq.Scl"))
}
func HstCndSclLeqNotZero(t *testing.T, a *hst.CndSclLeq, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	HstStmNotZero(t, a.Stm, append(msgs, "CndSclLeq.Stm"))
}
func HstCndSclLeqSliceEql(t *testing.T, e, a []*hst.CndSclLeq, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstCndSclLeqEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm CndSclLeq (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstCndSclGeqEql(t *testing.T, e, a *hst.CndSclGeq, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "CndSclGeq.Tmes"))
	HstStmEql(t, e.Stm, a.Stm, append(msgs, "CndSclGeq.Stm"))
	FltEql(t, e.Scl, a.Scl, append(msgs, "CndSclGeq.Scl"))
}
func HstCndSclGeqNotZero(t *testing.T, a *hst.CndSclGeq, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	HstStmNotZero(t, a.Stm, append(msgs, "CndSclGeq.Stm"))
}
func HstCndSclGeqSliceEql(t *testing.T, e, a []*hst.CndSclGeq, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstCndSclGeqEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm CndSclGeq (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstCndInrEqlEql(t *testing.T, e, a *hst.CndInrEql, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "CndInrEql.Tmes"))
	HstStmEql(t, e.Stm, a.Stm, append(msgs, "CndInrEql.Stm"))
	UntEql(t, e.Off, a.Off, append(msgs, "CndInrEql.Off"))
}
func HstCndInrEqlNotZero(t *testing.T, a *hst.CndInrEql, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	HstStmNotZero(t, a.Stm, append(msgs, "CndInrEql.Stm"))
}
func HstCndInrEqlSliceEql(t *testing.T, e, a []*hst.CndInrEql, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstCndInrEqlEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm CndInrEql (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstCndInrNeqEql(t *testing.T, e, a *hst.CndInrNeq, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "CndInrNeq.Tmes"))
	HstStmEql(t, e.Stm, a.Stm, append(msgs, "CndInrNeq.Stm"))
	UntEql(t, e.Off, a.Off, append(msgs, "CndInrNeq.Off"))
}
func HstCndInrNeqNotZero(t *testing.T, a *hst.CndInrNeq, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	HstStmNotZero(t, a.Stm, append(msgs, "CndInrNeq.Stm"))
}
func HstCndInrNeqSliceEql(t *testing.T, e, a []*hst.CndInrNeq, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstCndInrNeqEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm CndInrNeq (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstCndInrLssEql(t *testing.T, e, a *hst.CndInrLss, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "CndInrLss.Tmes"))
	HstStmEql(t, e.Stm, a.Stm, append(msgs, "CndInrLss.Stm"))
	UntEql(t, e.Off, a.Off, append(msgs, "CndInrLss.Off"))
}
func HstCndInrLssNotZero(t *testing.T, a *hst.CndInrLss, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	HstStmNotZero(t, a.Stm, append(msgs, "CndInrLss.Stm"))
}
func HstCndInrLssSliceEql(t *testing.T, e, a []*hst.CndInrLss, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstCndInrLssEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm CndInrLss (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstCndInrGtrEql(t *testing.T, e, a *hst.CndInrGtr, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "CndInrGtr.Tmes"))
	HstStmEql(t, e.Stm, a.Stm, append(msgs, "CndInrGtr.Stm"))
	UntEql(t, e.Off, a.Off, append(msgs, "CndInrGtr.Off"))
}
func HstCndInrGtrNotZero(t *testing.T, a *hst.CndInrGtr, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	HstStmNotZero(t, a.Stm, append(msgs, "CndInrGtr.Stm"))
}
func HstCndInrGtrSliceEql(t *testing.T, e, a []*hst.CndInrGtr, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstCndInrGtrEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm CndInrGtr (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstCndInrLeqEql(t *testing.T, e, a *hst.CndInrLeq, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "CndInrLeq.Tmes"))
	HstStmEql(t, e.Stm, a.Stm, append(msgs, "CndInrLeq.Stm"))
	UntEql(t, e.Off, a.Off, append(msgs, "CndInrLeq.Off"))
}
func HstCndInrLeqNotZero(t *testing.T, a *hst.CndInrLeq, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	HstStmNotZero(t, a.Stm, append(msgs, "CndInrLeq.Stm"))
}
func HstCndInrLeqSliceEql(t *testing.T, e, a []*hst.CndInrLeq, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstCndInrLeqEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm CndInrLeq (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstCndInrGeqEql(t *testing.T, e, a *hst.CndInrGeq, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "CndInrGeq.Tmes"))
	HstStmEql(t, e.Stm, a.Stm, append(msgs, "CndInrGeq.Stm"))
	UntEql(t, e.Off, a.Off, append(msgs, "CndInrGeq.Off"))
}
func HstCndInrGeqNotZero(t *testing.T, a *hst.CndInrGeq, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	HstStmNotZero(t, a.Stm, append(msgs, "CndInrGeq.Stm"))
}
func HstCndInrGeqSliceEql(t *testing.T, e, a []*hst.CndInrGeq, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstCndInrGeqEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm CndInrGeq (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstCndOtrEqlEql(t *testing.T, e, a *hst.CndOtrEql, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "CndOtrEql.Tmes"))
	HstStmEql(t, e.Stm, a.Stm, append(msgs, "CndOtrEql.Stm"))
	UntEql(t, e.Off, a.Off, append(msgs, "CndOtrEql.Off"))
	HstStmEql(t, e.A, a.A, append(msgs, "CndOtrEql.A"))
}
func HstCndOtrEqlNotZero(t *testing.T, a *hst.CndOtrEql, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	HstStmNotZero(t, a.Stm, append(msgs, "CndOtrEql.Stm"))
}
func HstCndOtrEqlSliceEql(t *testing.T, e, a []*hst.CndOtrEql, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstCndOtrEqlEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm CndOtrEql (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstCndOtrNeqEql(t *testing.T, e, a *hst.CndOtrNeq, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "CndOtrNeq.Tmes"))
	HstStmEql(t, e.Stm, a.Stm, append(msgs, "CndOtrNeq.Stm"))
	UntEql(t, e.Off, a.Off, append(msgs, "CndOtrNeq.Off"))
	HstStmEql(t, e.A, a.A, append(msgs, "CndOtrNeq.A"))
}
func HstCndOtrNeqNotZero(t *testing.T, a *hst.CndOtrNeq, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	HstStmNotZero(t, a.Stm, append(msgs, "CndOtrNeq.Stm"))
}
func HstCndOtrNeqSliceEql(t *testing.T, e, a []*hst.CndOtrNeq, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstCndOtrNeqEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm CndOtrNeq (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstCndOtrLssEql(t *testing.T, e, a *hst.CndOtrLss, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "CndOtrLss.Tmes"))
	HstStmEql(t, e.Stm, a.Stm, append(msgs, "CndOtrLss.Stm"))
	UntEql(t, e.Off, a.Off, append(msgs, "CndOtrLss.Off"))
	HstStmEql(t, e.A, a.A, append(msgs, "CndOtrLss.A"))
}
func HstCndOtrLssNotZero(t *testing.T, a *hst.CndOtrLss, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	HstStmNotZero(t, a.Stm, append(msgs, "CndOtrLss.Stm"))
}
func HstCndOtrLssSliceEql(t *testing.T, e, a []*hst.CndOtrLss, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstCndOtrLssEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm CndOtrLss (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstCndOtrGtrEql(t *testing.T, e, a *hst.CndOtrGtr, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "CndOtrGtr.Tmes"))
	HstStmEql(t, e.Stm, a.Stm, append(msgs, "CndOtrGtr.Stm"))
	UntEql(t, e.Off, a.Off, append(msgs, "CndOtrGtr.Off"))
	HstStmEql(t, e.A, a.A, append(msgs, "CndOtrGtr.A"))
}
func HstCndOtrGtrNotZero(t *testing.T, a *hst.CndOtrGtr, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	HstStmNotZero(t, a.Stm, append(msgs, "CndOtrGtr.Stm"))
}
func HstCndOtrGtrSliceEql(t *testing.T, e, a []*hst.CndOtrGtr, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstCndOtrGtrEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm CndOtrGtr (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstCndOtrLeqEql(t *testing.T, e, a *hst.CndOtrLeq, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "CndOtrLeq.Tmes"))
	HstStmEql(t, e.Stm, a.Stm, append(msgs, "CndOtrLeq.Stm"))
	UntEql(t, e.Off, a.Off, append(msgs, "CndOtrLeq.Off"))
	HstStmEql(t, e.A, a.A, append(msgs, "CndOtrLeq.A"))
}
func HstCndOtrLeqNotZero(t *testing.T, a *hst.CndOtrLeq, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	HstStmNotZero(t, a.Stm, append(msgs, "CndOtrLeq.Stm"))
}
func HstCndOtrLeqSliceEql(t *testing.T, e, a []*hst.CndOtrLeq, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstCndOtrLeqEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm CndOtrLeq (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstCndOtrGeqEql(t *testing.T, e, a *hst.CndOtrGeq, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "CndOtrGeq.Tmes"))
	HstStmEql(t, e.Stm, a.Stm, append(msgs, "CndOtrGeq.Stm"))
	UntEql(t, e.Off, a.Off, append(msgs, "CndOtrGeq.Off"))
	HstStmEql(t, e.A, a.A, append(msgs, "CndOtrGeq.A"))
}
func HstCndOtrGeqNotZero(t *testing.T, a *hst.CndOtrGeq, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	HstStmNotZero(t, a.Stm, append(msgs, "CndOtrGeq.Stm"))
}
func HstCndOtrGeqSliceEql(t *testing.T, e, a []*hst.CndOtrGeq, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstCndOtrGeqEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm CndOtrGeq (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstCndCnd1AndEql(t *testing.T, e, a *hst.CndCnd1And, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "CndCnd1And.Tmes"))
	HstCndEql(t, e.Cnd, a.Cnd, append(msgs, "CndCnd1And.Cnd"))
	HstCndEql(t, e.A, a.A, append(msgs, "CndCnd1And.A"))
}
func HstCndCnd1AndNotZero(t *testing.T, a *hst.CndCnd1And, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	HstCndNotZero(t, a.Cnd, append(msgs, "CndCnd1And.Cnd"))
}
func HstCndCnd1AndSliceEql(t *testing.T, e, a []*hst.CndCnd1And, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstCndCnd1AndEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm CndCnd1And (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstCndCnd2SeqEql(t *testing.T, e, a *hst.CndCnd2Seq, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "CndCnd2Seq.Tmes"))
	HstCndEql(t, e.Cnd, a.Cnd, append(msgs, "CndCnd2Seq.Cnd"))
	TmeEql(t, e.Dur, a.Dur, append(msgs, "CndCnd2Seq.Dur"))
	HstCndEql(t, e.A, a.A, append(msgs, "CndCnd2Seq.A"))
}
func HstCndCnd2SeqNotZero(t *testing.T, a *hst.CndCnd2Seq, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	HstCndNotZero(t, a.Cnd, append(msgs, "CndCnd2Seq.Cnd"))
}
func HstCndCnd2SeqSliceEql(t *testing.T, e, a []*hst.CndCnd2Seq, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstCndCnd2SeqEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm CndCnd2Seq (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func HstStgyStgyEql(t *testing.T, e, a *hst.StgyStgy, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	HstCndEql(t, e.Cnd, a.Cnd, append(msgs, "StgyStgy.Cnd"))
	BolEql(t, e.IsLong, a.IsLong, append(msgs, "StgyStgy.IsLong"))
	FltEql(t, e.PrfLim, a.PrfLim, append(msgs, "StgyStgy.PrfLim"))
	FltEql(t, e.LosLim, a.LosLim, append(msgs, "StgyStgy.LosLim"))
	TmeEql(t, e.DurLim, a.DurLim, append(msgs, "StgyStgy.DurLim"))
	FltEql(t, e.MinPnlPct, a.MinPnlPct, append(msgs, "StgyStgy.MinPnlPct"))
	HstInstrEql(t, e.Instr, a.Instr, append(msgs, "StgyStgy.Instr"))
	HstCndSliceEql(t, e.Clss, a.Clss, append(msgs, "StgyStgy.Clss"))
	AnaTrdsEql(t, e.Trds, a.Trds, append(msgs, "StgyStgy.Trds"))
	HstStmsEql(t, e.FtrStms, a.FtrStms, append(msgs, "StgyStgy.FtrStms"))
}
func HstStgyStgyNotZero(t *testing.T, a *hst.StgyStgy, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	HstCndNotZero(t, a.Cnd, append(msgs, "StgyStgy.Cnd"))
	FltNotZero(t, a.PrfLim, append(msgs, "StgyStgy.PrfLim"))
	FltNotZero(t, a.LosLim, append(msgs, "StgyStgy.LosLim"))
	TmeNotZero(t, a.DurLim, append(msgs, "StgyStgy.DurLim"))
	HstInstrNotZero(t, a.Instr, append(msgs, "StgyStgy.Instr"))
}
func HstStgyStgySliceEql(t *testing.T, e, a []*hst.StgyStgy, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		HstStgyStgyEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StgyStgy (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltInstrEurUsdEql(t *testing.T, e, a *rlt.InstrEurUsd, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "InstrEurUsd.Id"))
	RltPrvEql(t, e.Prv, a.Prv, append(msgs, "InstrEurUsd.Prv"))
	AnaInstrEql(t, e.Ana, a.Ana, append(msgs, "InstrEurUsd.Ana"))
	TmeRngSliceEql(t, e.Rng, a.Rng, append(msgs, "InstrEurUsd.Rng"))
}
func RltInstrEurUsdNotZero(t *testing.T, a *rlt.InstrEurUsd, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "InstrEurUsd.Id"))
	RltPrvNotZero(t, a.Prv, append(msgs, "InstrEurUsd.Prv"))
	AnaInstrNotZero(t, a.Ana, append(msgs, "InstrEurUsd.Ana"))
}
func RltInstrEurUsdSliceEql(t *testing.T, e, a []*rlt.InstrEurUsd, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltInstrEurUsdEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm InstrEurUsd (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltInstrAudUsdEql(t *testing.T, e, a *rlt.InstrAudUsd, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "InstrAudUsd.Id"))
	RltPrvEql(t, e.Prv, a.Prv, append(msgs, "InstrAudUsd.Prv"))
	AnaInstrEql(t, e.Ana, a.Ana, append(msgs, "InstrAudUsd.Ana"))
	TmeRngSliceEql(t, e.Rng, a.Rng, append(msgs, "InstrAudUsd.Rng"))
}
func RltInstrAudUsdNotZero(t *testing.T, a *rlt.InstrAudUsd, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "InstrAudUsd.Id"))
	RltPrvNotZero(t, a.Prv, append(msgs, "InstrAudUsd.Prv"))
	AnaInstrNotZero(t, a.Ana, append(msgs, "InstrAudUsd.Ana"))
}
func RltInstrAudUsdSliceEql(t *testing.T, e, a []*rlt.InstrAudUsd, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltInstrAudUsdEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm InstrAudUsd (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltInstrNzdUsdEql(t *testing.T, e, a *rlt.InstrNzdUsd, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "InstrNzdUsd.Id"))
	RltPrvEql(t, e.Prv, a.Prv, append(msgs, "InstrNzdUsd.Prv"))
	AnaInstrEql(t, e.Ana, a.Ana, append(msgs, "InstrNzdUsd.Ana"))
	TmeRngSliceEql(t, e.Rng, a.Rng, append(msgs, "InstrNzdUsd.Rng"))
}
func RltInstrNzdUsdNotZero(t *testing.T, a *rlt.InstrNzdUsd, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "InstrNzdUsd.Id"))
	RltPrvNotZero(t, a.Prv, append(msgs, "InstrNzdUsd.Prv"))
	AnaInstrNotZero(t, a.Ana, append(msgs, "InstrNzdUsd.Ana"))
}
func RltInstrNzdUsdSliceEql(t *testing.T, e, a []*rlt.InstrNzdUsd, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltInstrNzdUsdEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm InstrNzdUsd (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltInstrGbpUsdEql(t *testing.T, e, a *rlt.InstrGbpUsd, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "InstrGbpUsd.Id"))
	RltPrvEql(t, e.Prv, a.Prv, append(msgs, "InstrGbpUsd.Prv"))
	AnaInstrEql(t, e.Ana, a.Ana, append(msgs, "InstrGbpUsd.Ana"))
	TmeRngSliceEql(t, e.Rng, a.Rng, append(msgs, "InstrGbpUsd.Rng"))
}
func RltInstrGbpUsdNotZero(t *testing.T, a *rlt.InstrGbpUsd, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "InstrGbpUsd.Id"))
	RltPrvNotZero(t, a.Prv, append(msgs, "InstrGbpUsd.Prv"))
	AnaInstrNotZero(t, a.Ana, append(msgs, "InstrGbpUsd.Ana"))
}
func RltInstrGbpUsdSliceEql(t *testing.T, e, a []*rlt.InstrGbpUsd, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltInstrGbpUsdEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm InstrGbpUsd (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltInrvlIEql(t *testing.T, e, a *rlt.InrvlI, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "InrvlI.Id"))
	RltInstrEql(t, e.Instr, a.Instr, append(msgs, "InrvlI.Instr"))
	AnaTmeIdxsEql(t, e.Pkts, a.Pkts, append(msgs, "InrvlI.Pkts"))
	TmeEql(t, e.Dur, a.Dur, append(msgs, "InrvlI.Dur"))
}
func RltInrvlINotZero(t *testing.T, a *rlt.InrvlI, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "InrvlI.Id"))
	RltInstrNotZero(t, a.Instr, append(msgs, "InrvlI.Instr"))
}
func RltInrvlISliceEql(t *testing.T, e, a []*rlt.InrvlI, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltInrvlIEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm InrvlI (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltSideBidEql(t *testing.T, e, a *rlt.SideBid, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "SideBid.Id"))
	RltInrvlEql(t, e.Inrvl, a.Inrvl, append(msgs, "SideBid.Inrvl"))
}
func RltSideBidNotZero(t *testing.T, a *rlt.SideBid, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "SideBid.Id"))
	RltInrvlNotZero(t, a.Inrvl, append(msgs, "SideBid.Inrvl"))
}
func RltSideBidSliceEql(t *testing.T, e, a []*rlt.SideBid, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltSideBidEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm SideBid (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltSideAskEql(t *testing.T, e, a *rlt.SideAsk, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "SideAsk.Id"))
	RltInrvlEql(t, e.Inrvl, a.Inrvl, append(msgs, "SideAsk.Inrvl"))
}
func RltSideAskNotZero(t *testing.T, a *rlt.SideAsk, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "SideAsk.Id"))
	RltInrvlNotZero(t, a.Inrvl, append(msgs, "SideAsk.Inrvl"))
}
func RltSideAskSliceEql(t *testing.T, e, a []*rlt.SideAsk, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltSideAskEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm SideAsk (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltStmRteFstEql(t *testing.T, e, a *rlt.StmRteFst, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "StmRteFst.Id"))
	RltSideEql(t, e.Side, a.Side, append(msgs, "StmRteFst.Side"))
}
func RltStmRteFstNotZero(t *testing.T, a *rlt.StmRteFst, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "StmRteFst.Id"))
	RltSideNotZero(t, a.Side, append(msgs, "StmRteFst.Side"))
}
func RltStmRteFstSliceEql(t *testing.T, e, a []*rlt.StmRteFst, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltStmRteFstEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmRteFst (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltStmRteLstEql(t *testing.T, e, a *rlt.StmRteLst, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "StmRteLst.Id"))
	RltSideEql(t, e.Side, a.Side, append(msgs, "StmRteLst.Side"))
}
func RltStmRteLstNotZero(t *testing.T, a *rlt.StmRteLst, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "StmRteLst.Id"))
	RltSideNotZero(t, a.Side, append(msgs, "StmRteLst.Side"))
}
func RltStmRteLstSliceEql(t *testing.T, e, a []*rlt.StmRteLst, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltStmRteLstEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmRteLst (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltStmRteSumEql(t *testing.T, e, a *rlt.StmRteSum, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "StmRteSum.Id"))
	RltSideEql(t, e.Side, a.Side, append(msgs, "StmRteSum.Side"))
}
func RltStmRteSumNotZero(t *testing.T, a *rlt.StmRteSum, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "StmRteSum.Id"))
	RltSideNotZero(t, a.Side, append(msgs, "StmRteSum.Side"))
}
func RltStmRteSumSliceEql(t *testing.T, e, a []*rlt.StmRteSum, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltStmRteSumEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmRteSum (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltStmRtePrdEql(t *testing.T, e, a *rlt.StmRtePrd, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "StmRtePrd.Id"))
	RltSideEql(t, e.Side, a.Side, append(msgs, "StmRtePrd.Side"))
}
func RltStmRtePrdNotZero(t *testing.T, a *rlt.StmRtePrd, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "StmRtePrd.Id"))
	RltSideNotZero(t, a.Side, append(msgs, "StmRtePrd.Side"))
}
func RltStmRtePrdSliceEql(t *testing.T, e, a []*rlt.StmRtePrd, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltStmRtePrdEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmRtePrd (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltStmRteMinEql(t *testing.T, e, a *rlt.StmRteMin, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "StmRteMin.Id"))
	RltSideEql(t, e.Side, a.Side, append(msgs, "StmRteMin.Side"))
}
func RltStmRteMinNotZero(t *testing.T, a *rlt.StmRteMin, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "StmRteMin.Id"))
	RltSideNotZero(t, a.Side, append(msgs, "StmRteMin.Side"))
}
func RltStmRteMinSliceEql(t *testing.T, e, a []*rlt.StmRteMin, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltStmRteMinEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmRteMin (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltStmRteMaxEql(t *testing.T, e, a *rlt.StmRteMax, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "StmRteMax.Id"))
	RltSideEql(t, e.Side, a.Side, append(msgs, "StmRteMax.Side"))
}
func RltStmRteMaxNotZero(t *testing.T, a *rlt.StmRteMax, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "StmRteMax.Id"))
	RltSideNotZero(t, a.Side, append(msgs, "StmRteMax.Side"))
}
func RltStmRteMaxSliceEql(t *testing.T, e, a []*rlt.StmRteMax, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltStmRteMaxEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmRteMax (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltStmRteMidEql(t *testing.T, e, a *rlt.StmRteMid, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "StmRteMid.Id"))
	RltSideEql(t, e.Side, a.Side, append(msgs, "StmRteMid.Side"))
}
func RltStmRteMidNotZero(t *testing.T, a *rlt.StmRteMid, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "StmRteMid.Id"))
	RltSideNotZero(t, a.Side, append(msgs, "StmRteMid.Side"))
}
func RltStmRteMidSliceEql(t *testing.T, e, a []*rlt.StmRteMid, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltStmRteMidEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmRteMid (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltStmRteMdnEql(t *testing.T, e, a *rlt.StmRteMdn, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "StmRteMdn.Id"))
	RltSideEql(t, e.Side, a.Side, append(msgs, "StmRteMdn.Side"))
}
func RltStmRteMdnNotZero(t *testing.T, a *rlt.StmRteMdn, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "StmRteMdn.Id"))
	RltSideNotZero(t, a.Side, append(msgs, "StmRteMdn.Side"))
}
func RltStmRteMdnSliceEql(t *testing.T, e, a []*rlt.StmRteMdn, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltStmRteMdnEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmRteMdn (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltStmRteSmaEql(t *testing.T, e, a *rlt.StmRteSma, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "StmRteSma.Id"))
	RltSideEql(t, e.Side, a.Side, append(msgs, "StmRteSma.Side"))
}
func RltStmRteSmaNotZero(t *testing.T, a *rlt.StmRteSma, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "StmRteSma.Id"))
	RltSideNotZero(t, a.Side, append(msgs, "StmRteSma.Side"))
}
func RltStmRteSmaSliceEql(t *testing.T, e, a []*rlt.StmRteSma, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltStmRteSmaEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmRteSma (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltStmRteGmaEql(t *testing.T, e, a *rlt.StmRteGma, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "StmRteGma.Id"))
	RltSideEql(t, e.Side, a.Side, append(msgs, "StmRteGma.Side"))
}
func RltStmRteGmaNotZero(t *testing.T, a *rlt.StmRteGma, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "StmRteGma.Id"))
	RltSideNotZero(t, a.Side, append(msgs, "StmRteGma.Side"))
}
func RltStmRteGmaSliceEql(t *testing.T, e, a []*rlt.StmRteGma, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltStmRteGmaEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmRteGma (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltStmRteWmaEql(t *testing.T, e, a *rlt.StmRteWma, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "StmRteWma.Id"))
	RltSideEql(t, e.Side, a.Side, append(msgs, "StmRteWma.Side"))
}
func RltStmRteWmaNotZero(t *testing.T, a *rlt.StmRteWma, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "StmRteWma.Id"))
	RltSideNotZero(t, a.Side, append(msgs, "StmRteWma.Side"))
}
func RltStmRteWmaSliceEql(t *testing.T, e, a []*rlt.StmRteWma, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltStmRteWmaEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmRteWma (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltStmRteRsiEql(t *testing.T, e, a *rlt.StmRteRsi, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "StmRteRsi.Id"))
	RltSideEql(t, e.Side, a.Side, append(msgs, "StmRteRsi.Side"))
}
func RltStmRteRsiNotZero(t *testing.T, a *rlt.StmRteRsi, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "StmRteRsi.Id"))
	RltSideNotZero(t, a.Side, append(msgs, "StmRteRsi.Side"))
}
func RltStmRteRsiSliceEql(t *testing.T, e, a []*rlt.StmRteRsi, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltStmRteRsiEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmRteRsi (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltStmRteWrsiEql(t *testing.T, e, a *rlt.StmRteWrsi, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "StmRteWrsi.Id"))
	RltSideEql(t, e.Side, a.Side, append(msgs, "StmRteWrsi.Side"))
}
func RltStmRteWrsiNotZero(t *testing.T, a *rlt.StmRteWrsi, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "StmRteWrsi.Id"))
	RltSideNotZero(t, a.Side, append(msgs, "StmRteWrsi.Side"))
}
func RltStmRteWrsiSliceEql(t *testing.T, e, a []*rlt.StmRteWrsi, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltStmRteWrsiEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmRteWrsi (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltStmRteAlmaEql(t *testing.T, e, a *rlt.StmRteAlma, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "StmRteAlma.Id"))
	RltSideEql(t, e.Side, a.Side, append(msgs, "StmRteAlma.Side"))
}
func RltStmRteAlmaNotZero(t *testing.T, a *rlt.StmRteAlma, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "StmRteAlma.Id"))
	RltSideNotZero(t, a.Side, append(msgs, "StmRteAlma.Side"))
}
func RltStmRteAlmaSliceEql(t *testing.T, e, a []*rlt.StmRteAlma, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltStmRteAlmaEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmRteAlma (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltStmRteVrncEql(t *testing.T, e, a *rlt.StmRteVrnc, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "StmRteVrnc.Id"))
	RltSideEql(t, e.Side, a.Side, append(msgs, "StmRteVrnc.Side"))
}
func RltStmRteVrncNotZero(t *testing.T, a *rlt.StmRteVrnc, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "StmRteVrnc.Id"))
	RltSideNotZero(t, a.Side, append(msgs, "StmRteVrnc.Side"))
}
func RltStmRteVrncSliceEql(t *testing.T, e, a []*rlt.StmRteVrnc, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltStmRteVrncEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmRteVrnc (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltStmRteStdEql(t *testing.T, e, a *rlt.StmRteStd, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "StmRteStd.Id"))
	RltSideEql(t, e.Side, a.Side, append(msgs, "StmRteStd.Side"))
}
func RltStmRteStdNotZero(t *testing.T, a *rlt.StmRteStd, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "StmRteStd.Id"))
	RltSideNotZero(t, a.Side, append(msgs, "StmRteStd.Side"))
}
func RltStmRteStdSliceEql(t *testing.T, e, a []*rlt.StmRteStd, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltStmRteStdEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmRteStd (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltStmRteRngFulEql(t *testing.T, e, a *rlt.StmRteRngFul, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "StmRteRngFul.Id"))
	RltSideEql(t, e.Side, a.Side, append(msgs, "StmRteRngFul.Side"))
}
func RltStmRteRngFulNotZero(t *testing.T, a *rlt.StmRteRngFul, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "StmRteRngFul.Id"))
	RltSideNotZero(t, a.Side, append(msgs, "StmRteRngFul.Side"))
}
func RltStmRteRngFulSliceEql(t *testing.T, e, a []*rlt.StmRteRngFul, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltStmRteRngFulEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmRteRngFul (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltStmRteRngLstEql(t *testing.T, e, a *rlt.StmRteRngLst, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "StmRteRngLst.Id"))
	RltSideEql(t, e.Side, a.Side, append(msgs, "StmRteRngLst.Side"))
}
func RltStmRteRngLstNotZero(t *testing.T, a *rlt.StmRteRngLst, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "StmRteRngLst.Id"))
	RltSideNotZero(t, a.Side, append(msgs, "StmRteRngLst.Side"))
}
func RltStmRteRngLstSliceEql(t *testing.T, e, a []*rlt.StmRteRngLst, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltStmRteRngLstEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmRteRngLst (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltStmRteProLstEql(t *testing.T, e, a *rlt.StmRteProLst, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "StmRteProLst.Id"))
	RltSideEql(t, e.Side, a.Side, append(msgs, "StmRteProLst.Side"))
}
func RltStmRteProLstNotZero(t *testing.T, a *rlt.StmRteProLst, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "StmRteProLst.Id"))
	RltSideNotZero(t, a.Side, append(msgs, "StmRteProLst.Side"))
}
func RltStmRteProLstSliceEql(t *testing.T, e, a []*rlt.StmRteProLst, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltStmRteProLstEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmRteProLst (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltStmRteProSmaEql(t *testing.T, e, a *rlt.StmRteProSma, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "StmRteProSma.Id"))
	RltSideEql(t, e.Side, a.Side, append(msgs, "StmRteProSma.Side"))
}
func RltStmRteProSmaNotZero(t *testing.T, a *rlt.StmRteProSma, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "StmRteProSma.Id"))
	RltSideNotZero(t, a.Side, append(msgs, "StmRteProSma.Side"))
}
func RltStmRteProSmaSliceEql(t *testing.T, e, a []*rlt.StmRteProSma, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltStmRteProSmaEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmRteProSma (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltStmRteProAlmaEql(t *testing.T, e, a *rlt.StmRteProAlma, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "StmRteProAlma.Id"))
	RltSideEql(t, e.Side, a.Side, append(msgs, "StmRteProAlma.Side"))
}
func RltStmRteProAlmaNotZero(t *testing.T, a *rlt.StmRteProAlma, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "StmRteProAlma.Id"))
	RltSideNotZero(t, a.Side, append(msgs, "StmRteProAlma.Side"))
}
func RltStmRteProAlmaSliceEql(t *testing.T, e, a []*rlt.StmRteProAlma, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltStmRteProAlmaEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmRteProAlma (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltStmRte1SarEql(t *testing.T, e, a *rlt.StmRte1Sar, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "StmRte1Sar.Id"))
	RltSideEql(t, e.Side, a.Side, append(msgs, "StmRte1Sar.Side"))
	FltEql(t, e.AfInc, a.AfInc, append(msgs, "StmRte1Sar.AfInc"))
	FltEql(t, e.AfMax, a.AfMax, append(msgs, "StmRte1Sar.AfMax"))
}
func RltStmRte1SarNotZero(t *testing.T, a *rlt.StmRte1Sar, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "StmRte1Sar.Id"))
	RltSideNotZero(t, a.Side, append(msgs, "StmRte1Sar.Side"))
}
func RltStmRte1SarSliceEql(t *testing.T, e, a []*rlt.StmRte1Sar, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltStmRte1SarEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmRte1Sar (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltStmRteEmaEql(t *testing.T, e, a *rlt.StmRteEma, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "StmRteEma.Id"))
	RltSideEql(t, e.Side, a.Side, append(msgs, "StmRteEma.Side"))
	FltEql(t, e.Prv, a.Prv, append(msgs, "StmRteEma.Prv"))
}
func RltStmRteEmaNotZero(t *testing.T, a *rlt.StmRteEma, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "StmRteEma.Id"))
	RltSideNotZero(t, a.Side, append(msgs, "StmRteEma.Side"))
	FltNotZero(t, a.Prv, append(msgs, "StmRteEma.Prv"))
}
func RltStmRteEmaSliceEql(t *testing.T, e, a []*rlt.StmRteEma, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltStmRteEmaEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmRteEma (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltStmUnaPosEql(t *testing.T, e, a *rlt.StmUnaPos, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "StmUnaPos.Id"))
	RltStmEql(t, e.Stm, a.Stm, append(msgs, "StmUnaPos.Stm"))
}
func RltStmUnaPosNotZero(t *testing.T, a *rlt.StmUnaPos, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "StmUnaPos.Id"))
	RltStmNotZero(t, a.Stm, append(msgs, "StmUnaPos.Stm"))
}
func RltStmUnaPosSliceEql(t *testing.T, e, a []*rlt.StmUnaPos, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltStmUnaPosEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmUnaPos (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltStmUnaNegEql(t *testing.T, e, a *rlt.StmUnaNeg, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "StmUnaNeg.Id"))
	RltStmEql(t, e.Stm, a.Stm, append(msgs, "StmUnaNeg.Stm"))
}
func RltStmUnaNegNotZero(t *testing.T, a *rlt.StmUnaNeg, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "StmUnaNeg.Id"))
	RltStmNotZero(t, a.Stm, append(msgs, "StmUnaNeg.Stm"))
}
func RltStmUnaNegSliceEql(t *testing.T, e, a []*rlt.StmUnaNeg, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltStmUnaNegEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmUnaNeg (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltStmUnaInvEql(t *testing.T, e, a *rlt.StmUnaInv, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "StmUnaInv.Id"))
	RltStmEql(t, e.Stm, a.Stm, append(msgs, "StmUnaInv.Stm"))
}
func RltStmUnaInvNotZero(t *testing.T, a *rlt.StmUnaInv, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "StmUnaInv.Id"))
	RltStmNotZero(t, a.Stm, append(msgs, "StmUnaInv.Stm"))
}
func RltStmUnaInvSliceEql(t *testing.T, e, a []*rlt.StmUnaInv, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltStmUnaInvEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmUnaInv (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltStmUnaSqrEql(t *testing.T, e, a *rlt.StmUnaSqr, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "StmUnaSqr.Id"))
	RltStmEql(t, e.Stm, a.Stm, append(msgs, "StmUnaSqr.Stm"))
}
func RltStmUnaSqrNotZero(t *testing.T, a *rlt.StmUnaSqr, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "StmUnaSqr.Id"))
	RltStmNotZero(t, a.Stm, append(msgs, "StmUnaSqr.Stm"))
}
func RltStmUnaSqrSliceEql(t *testing.T, e, a []*rlt.StmUnaSqr, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltStmUnaSqrEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmUnaSqr (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltStmUnaSqrtEql(t *testing.T, e, a *rlt.StmUnaSqrt, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "StmUnaSqrt.Id"))
	RltStmEql(t, e.Stm, a.Stm, append(msgs, "StmUnaSqrt.Stm"))
}
func RltStmUnaSqrtNotZero(t *testing.T, a *rlt.StmUnaSqrt, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "StmUnaSqrt.Id"))
	RltStmNotZero(t, a.Stm, append(msgs, "StmUnaSqrt.Stm"))
}
func RltStmUnaSqrtSliceEql(t *testing.T, e, a []*rlt.StmUnaSqrt, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltStmUnaSqrtEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmUnaSqrt (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltStmSclAddEql(t *testing.T, e, a *rlt.StmSclAdd, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "StmSclAdd.Id"))
	RltStmEql(t, e.Stm, a.Stm, append(msgs, "StmSclAdd.Stm"))
	FltEql(t, e.Scl, a.Scl, append(msgs, "StmSclAdd.Scl"))
}
func RltStmSclAddNotZero(t *testing.T, a *rlt.StmSclAdd, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "StmSclAdd.Id"))
	RltStmNotZero(t, a.Stm, append(msgs, "StmSclAdd.Stm"))
}
func RltStmSclAddSliceEql(t *testing.T, e, a []*rlt.StmSclAdd, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltStmSclAddEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmSclAdd (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltStmSclSubEql(t *testing.T, e, a *rlt.StmSclSub, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "StmSclSub.Id"))
	RltStmEql(t, e.Stm, a.Stm, append(msgs, "StmSclSub.Stm"))
	FltEql(t, e.Scl, a.Scl, append(msgs, "StmSclSub.Scl"))
}
func RltStmSclSubNotZero(t *testing.T, a *rlt.StmSclSub, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "StmSclSub.Id"))
	RltStmNotZero(t, a.Stm, append(msgs, "StmSclSub.Stm"))
}
func RltStmSclSubSliceEql(t *testing.T, e, a []*rlt.StmSclSub, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltStmSclSubEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmSclSub (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltStmSclMulEql(t *testing.T, e, a *rlt.StmSclMul, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "StmSclMul.Id"))
	RltStmEql(t, e.Stm, a.Stm, append(msgs, "StmSclMul.Stm"))
	FltEql(t, e.Scl, a.Scl, append(msgs, "StmSclMul.Scl"))
}
func RltStmSclMulNotZero(t *testing.T, a *rlt.StmSclMul, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "StmSclMul.Id"))
	RltStmNotZero(t, a.Stm, append(msgs, "StmSclMul.Stm"))
}
func RltStmSclMulSliceEql(t *testing.T, e, a []*rlt.StmSclMul, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltStmSclMulEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmSclMul (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltStmSclDivEql(t *testing.T, e, a *rlt.StmSclDiv, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "StmSclDiv.Id"))
	RltStmEql(t, e.Stm, a.Stm, append(msgs, "StmSclDiv.Stm"))
	FltEql(t, e.Scl, a.Scl, append(msgs, "StmSclDiv.Scl"))
}
func RltStmSclDivNotZero(t *testing.T, a *rlt.StmSclDiv, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "StmSclDiv.Id"))
	RltStmNotZero(t, a.Stm, append(msgs, "StmSclDiv.Stm"))
}
func RltStmSclDivSliceEql(t *testing.T, e, a []*rlt.StmSclDiv, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltStmSclDivEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmSclDiv (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltStmSclRemEql(t *testing.T, e, a *rlt.StmSclRem, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "StmSclRem.Id"))
	RltStmEql(t, e.Stm, a.Stm, append(msgs, "StmSclRem.Stm"))
	FltEql(t, e.Scl, a.Scl, append(msgs, "StmSclRem.Scl"))
}
func RltStmSclRemNotZero(t *testing.T, a *rlt.StmSclRem, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "StmSclRem.Id"))
	RltStmNotZero(t, a.Stm, append(msgs, "StmSclRem.Stm"))
}
func RltStmSclRemSliceEql(t *testing.T, e, a []*rlt.StmSclRem, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltStmSclRemEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmSclRem (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltStmSclPowEql(t *testing.T, e, a *rlt.StmSclPow, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "StmSclPow.Id"))
	RltStmEql(t, e.Stm, a.Stm, append(msgs, "StmSclPow.Stm"))
	FltEql(t, e.Scl, a.Scl, append(msgs, "StmSclPow.Scl"))
}
func RltStmSclPowNotZero(t *testing.T, a *rlt.StmSclPow, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "StmSclPow.Id"))
	RltStmNotZero(t, a.Stm, append(msgs, "StmSclPow.Stm"))
}
func RltStmSclPowSliceEql(t *testing.T, e, a []*rlt.StmSclPow, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltStmSclPowEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmSclPow (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltStmSclMinEql(t *testing.T, e, a *rlt.StmSclMin, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "StmSclMin.Id"))
	RltStmEql(t, e.Stm, a.Stm, append(msgs, "StmSclMin.Stm"))
	FltEql(t, e.Scl, a.Scl, append(msgs, "StmSclMin.Scl"))
}
func RltStmSclMinNotZero(t *testing.T, a *rlt.StmSclMin, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "StmSclMin.Id"))
	RltStmNotZero(t, a.Stm, append(msgs, "StmSclMin.Stm"))
}
func RltStmSclMinSliceEql(t *testing.T, e, a []*rlt.StmSclMin, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltStmSclMinEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmSclMin (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltStmSclMaxEql(t *testing.T, e, a *rlt.StmSclMax, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "StmSclMax.Id"))
	RltStmEql(t, e.Stm, a.Stm, append(msgs, "StmSclMax.Stm"))
	FltEql(t, e.Scl, a.Scl, append(msgs, "StmSclMax.Scl"))
}
func RltStmSclMaxNotZero(t *testing.T, a *rlt.StmSclMax, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "StmSclMax.Id"))
	RltStmNotZero(t, a.Stm, append(msgs, "StmSclMax.Stm"))
}
func RltStmSclMaxSliceEql(t *testing.T, e, a []*rlt.StmSclMax, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltStmSclMaxEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmSclMax (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltStmSelEqlEql(t *testing.T, e, a *rlt.StmSelEql, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "StmSelEql.Id"))
	RltStmEql(t, e.Stm, a.Stm, append(msgs, "StmSelEql.Stm"))
	FltEql(t, e.Sel, a.Sel, append(msgs, "StmSelEql.Sel"))
}
func RltStmSelEqlNotZero(t *testing.T, a *rlt.StmSelEql, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "StmSelEql.Id"))
	RltStmNotZero(t, a.Stm, append(msgs, "StmSelEql.Stm"))
}
func RltStmSelEqlSliceEql(t *testing.T, e, a []*rlt.StmSelEql, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltStmSelEqlEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmSelEql (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltStmSelNeqEql(t *testing.T, e, a *rlt.StmSelNeq, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "StmSelNeq.Id"))
	RltStmEql(t, e.Stm, a.Stm, append(msgs, "StmSelNeq.Stm"))
	FltEql(t, e.Sel, a.Sel, append(msgs, "StmSelNeq.Sel"))
}
func RltStmSelNeqNotZero(t *testing.T, a *rlt.StmSelNeq, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "StmSelNeq.Id"))
	RltStmNotZero(t, a.Stm, append(msgs, "StmSelNeq.Stm"))
}
func RltStmSelNeqSliceEql(t *testing.T, e, a []*rlt.StmSelNeq, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltStmSelNeqEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmSelNeq (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltStmSelLssEql(t *testing.T, e, a *rlt.StmSelLss, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "StmSelLss.Id"))
	RltStmEql(t, e.Stm, a.Stm, append(msgs, "StmSelLss.Stm"))
	FltEql(t, e.Sel, a.Sel, append(msgs, "StmSelLss.Sel"))
}
func RltStmSelLssNotZero(t *testing.T, a *rlt.StmSelLss, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "StmSelLss.Id"))
	RltStmNotZero(t, a.Stm, append(msgs, "StmSelLss.Stm"))
}
func RltStmSelLssSliceEql(t *testing.T, e, a []*rlt.StmSelLss, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltStmSelLssEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmSelLss (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltStmSelGtrEql(t *testing.T, e, a *rlt.StmSelGtr, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "StmSelGtr.Id"))
	RltStmEql(t, e.Stm, a.Stm, append(msgs, "StmSelGtr.Stm"))
	FltEql(t, e.Sel, a.Sel, append(msgs, "StmSelGtr.Sel"))
}
func RltStmSelGtrNotZero(t *testing.T, a *rlt.StmSelGtr, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "StmSelGtr.Id"))
	RltStmNotZero(t, a.Stm, append(msgs, "StmSelGtr.Stm"))
}
func RltStmSelGtrSliceEql(t *testing.T, e, a []*rlt.StmSelGtr, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltStmSelGtrEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmSelGtr (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltStmSelLeqEql(t *testing.T, e, a *rlt.StmSelLeq, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "StmSelLeq.Id"))
	RltStmEql(t, e.Stm, a.Stm, append(msgs, "StmSelLeq.Stm"))
	FltEql(t, e.Sel, a.Sel, append(msgs, "StmSelLeq.Sel"))
}
func RltStmSelLeqNotZero(t *testing.T, a *rlt.StmSelLeq, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "StmSelLeq.Id"))
	RltStmNotZero(t, a.Stm, append(msgs, "StmSelLeq.Stm"))
}
func RltStmSelLeqSliceEql(t *testing.T, e, a []*rlt.StmSelLeq, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltStmSelLeqEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmSelLeq (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltStmSelGeqEql(t *testing.T, e, a *rlt.StmSelGeq, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "StmSelGeq.Id"))
	RltStmEql(t, e.Stm, a.Stm, append(msgs, "StmSelGeq.Stm"))
	FltEql(t, e.Sel, a.Sel, append(msgs, "StmSelGeq.Sel"))
}
func RltStmSelGeqNotZero(t *testing.T, a *rlt.StmSelGeq, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "StmSelGeq.Id"))
	RltStmNotZero(t, a.Stm, append(msgs, "StmSelGeq.Stm"))
}
func RltStmSelGeqSliceEql(t *testing.T, e, a []*rlt.StmSelGeq, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltStmSelGeqEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmSelGeq (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltStmAggFstEql(t *testing.T, e, a *rlt.StmAggFst, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "StmAggFst.Id"))
	RltStmEql(t, e.Stm, a.Stm, append(msgs, "StmAggFst.Stm"))
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmAggFst.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmAggFst.Vals"))
	UntEql(t, e.Length, a.Length, append(msgs, "StmAggFst.Length"))
}
func RltStmAggFstNotZero(t *testing.T, a *rlt.StmAggFst, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "StmAggFst.Id"))
	RltStmNotZero(t, a.Stm, append(msgs, "StmAggFst.Stm"))
}
func RltStmAggFstSliceEql(t *testing.T, e, a []*rlt.StmAggFst, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltStmAggFstEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmAggFst (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltStmAggLstEql(t *testing.T, e, a *rlt.StmAggLst, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "StmAggLst.Id"))
	RltStmEql(t, e.Stm, a.Stm, append(msgs, "StmAggLst.Stm"))
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmAggLst.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmAggLst.Vals"))
	UntEql(t, e.Length, a.Length, append(msgs, "StmAggLst.Length"))
}
func RltStmAggLstNotZero(t *testing.T, a *rlt.StmAggLst, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "StmAggLst.Id"))
	RltStmNotZero(t, a.Stm, append(msgs, "StmAggLst.Stm"))
}
func RltStmAggLstSliceEql(t *testing.T, e, a []*rlt.StmAggLst, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltStmAggLstEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmAggLst (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltStmAggSumEql(t *testing.T, e, a *rlt.StmAggSum, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "StmAggSum.Id"))
	RltStmEql(t, e.Stm, a.Stm, append(msgs, "StmAggSum.Stm"))
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmAggSum.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmAggSum.Vals"))
	UntEql(t, e.Length, a.Length, append(msgs, "StmAggSum.Length"))
}
func RltStmAggSumNotZero(t *testing.T, a *rlt.StmAggSum, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "StmAggSum.Id"))
	RltStmNotZero(t, a.Stm, append(msgs, "StmAggSum.Stm"))
}
func RltStmAggSumSliceEql(t *testing.T, e, a []*rlt.StmAggSum, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltStmAggSumEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmAggSum (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltStmAggPrdEql(t *testing.T, e, a *rlt.StmAggPrd, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "StmAggPrd.Id"))
	RltStmEql(t, e.Stm, a.Stm, append(msgs, "StmAggPrd.Stm"))
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmAggPrd.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmAggPrd.Vals"))
	UntEql(t, e.Length, a.Length, append(msgs, "StmAggPrd.Length"))
}
func RltStmAggPrdNotZero(t *testing.T, a *rlt.StmAggPrd, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "StmAggPrd.Id"))
	RltStmNotZero(t, a.Stm, append(msgs, "StmAggPrd.Stm"))
}
func RltStmAggPrdSliceEql(t *testing.T, e, a []*rlt.StmAggPrd, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltStmAggPrdEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmAggPrd (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltStmAggMinEql(t *testing.T, e, a *rlt.StmAggMin, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "StmAggMin.Id"))
	RltStmEql(t, e.Stm, a.Stm, append(msgs, "StmAggMin.Stm"))
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmAggMin.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmAggMin.Vals"))
	UntEql(t, e.Length, a.Length, append(msgs, "StmAggMin.Length"))
}
func RltStmAggMinNotZero(t *testing.T, a *rlt.StmAggMin, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "StmAggMin.Id"))
	RltStmNotZero(t, a.Stm, append(msgs, "StmAggMin.Stm"))
}
func RltStmAggMinSliceEql(t *testing.T, e, a []*rlt.StmAggMin, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltStmAggMinEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmAggMin (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltStmAggMaxEql(t *testing.T, e, a *rlt.StmAggMax, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "StmAggMax.Id"))
	RltStmEql(t, e.Stm, a.Stm, append(msgs, "StmAggMax.Stm"))
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmAggMax.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmAggMax.Vals"))
	UntEql(t, e.Length, a.Length, append(msgs, "StmAggMax.Length"))
}
func RltStmAggMaxNotZero(t *testing.T, a *rlt.StmAggMax, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "StmAggMax.Id"))
	RltStmNotZero(t, a.Stm, append(msgs, "StmAggMax.Stm"))
}
func RltStmAggMaxSliceEql(t *testing.T, e, a []*rlt.StmAggMax, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltStmAggMaxEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmAggMax (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltStmAggMidEql(t *testing.T, e, a *rlt.StmAggMid, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "StmAggMid.Id"))
	RltStmEql(t, e.Stm, a.Stm, append(msgs, "StmAggMid.Stm"))
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmAggMid.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmAggMid.Vals"))
	UntEql(t, e.Length, a.Length, append(msgs, "StmAggMid.Length"))
}
func RltStmAggMidNotZero(t *testing.T, a *rlt.StmAggMid, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "StmAggMid.Id"))
	RltStmNotZero(t, a.Stm, append(msgs, "StmAggMid.Stm"))
}
func RltStmAggMidSliceEql(t *testing.T, e, a []*rlt.StmAggMid, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltStmAggMidEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmAggMid (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltStmAggMdnEql(t *testing.T, e, a *rlt.StmAggMdn, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "StmAggMdn.Id"))
	RltStmEql(t, e.Stm, a.Stm, append(msgs, "StmAggMdn.Stm"))
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmAggMdn.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmAggMdn.Vals"))
	UntEql(t, e.Length, a.Length, append(msgs, "StmAggMdn.Length"))
}
func RltStmAggMdnNotZero(t *testing.T, a *rlt.StmAggMdn, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "StmAggMdn.Id"))
	RltStmNotZero(t, a.Stm, append(msgs, "StmAggMdn.Stm"))
}
func RltStmAggMdnSliceEql(t *testing.T, e, a []*rlt.StmAggMdn, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltStmAggMdnEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmAggMdn (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltStmAggSmaEql(t *testing.T, e, a *rlt.StmAggSma, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "StmAggSma.Id"))
	RltStmEql(t, e.Stm, a.Stm, append(msgs, "StmAggSma.Stm"))
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmAggSma.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmAggSma.Vals"))
	UntEql(t, e.Length, a.Length, append(msgs, "StmAggSma.Length"))
}
func RltStmAggSmaNotZero(t *testing.T, a *rlt.StmAggSma, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "StmAggSma.Id"))
	RltStmNotZero(t, a.Stm, append(msgs, "StmAggSma.Stm"))
}
func RltStmAggSmaSliceEql(t *testing.T, e, a []*rlt.StmAggSma, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltStmAggSmaEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmAggSma (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltStmAggGmaEql(t *testing.T, e, a *rlt.StmAggGma, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "StmAggGma.Id"))
	RltStmEql(t, e.Stm, a.Stm, append(msgs, "StmAggGma.Stm"))
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmAggGma.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmAggGma.Vals"))
	UntEql(t, e.Length, a.Length, append(msgs, "StmAggGma.Length"))
}
func RltStmAggGmaNotZero(t *testing.T, a *rlt.StmAggGma, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "StmAggGma.Id"))
	RltStmNotZero(t, a.Stm, append(msgs, "StmAggGma.Stm"))
}
func RltStmAggGmaSliceEql(t *testing.T, e, a []*rlt.StmAggGma, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltStmAggGmaEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmAggGma (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltStmAggWmaEql(t *testing.T, e, a *rlt.StmAggWma, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "StmAggWma.Id"))
	RltStmEql(t, e.Stm, a.Stm, append(msgs, "StmAggWma.Stm"))
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmAggWma.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmAggWma.Vals"))
	UntEql(t, e.Length, a.Length, append(msgs, "StmAggWma.Length"))
}
func RltStmAggWmaNotZero(t *testing.T, a *rlt.StmAggWma, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "StmAggWma.Id"))
	RltStmNotZero(t, a.Stm, append(msgs, "StmAggWma.Stm"))
}
func RltStmAggWmaSliceEql(t *testing.T, e, a []*rlt.StmAggWma, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltStmAggWmaEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmAggWma (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltStmAggRsiEql(t *testing.T, e, a *rlt.StmAggRsi, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "StmAggRsi.Id"))
	RltStmEql(t, e.Stm, a.Stm, append(msgs, "StmAggRsi.Stm"))
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmAggRsi.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmAggRsi.Vals"))
	UntEql(t, e.Length, a.Length, append(msgs, "StmAggRsi.Length"))
}
func RltStmAggRsiNotZero(t *testing.T, a *rlt.StmAggRsi, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "StmAggRsi.Id"))
	RltStmNotZero(t, a.Stm, append(msgs, "StmAggRsi.Stm"))
}
func RltStmAggRsiSliceEql(t *testing.T, e, a []*rlt.StmAggRsi, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltStmAggRsiEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmAggRsi (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltStmAggWrsiEql(t *testing.T, e, a *rlt.StmAggWrsi, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "StmAggWrsi.Id"))
	RltStmEql(t, e.Stm, a.Stm, append(msgs, "StmAggWrsi.Stm"))
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmAggWrsi.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmAggWrsi.Vals"))
	UntEql(t, e.Length, a.Length, append(msgs, "StmAggWrsi.Length"))
}
func RltStmAggWrsiNotZero(t *testing.T, a *rlt.StmAggWrsi, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "StmAggWrsi.Id"))
	RltStmNotZero(t, a.Stm, append(msgs, "StmAggWrsi.Stm"))
}
func RltStmAggWrsiSliceEql(t *testing.T, e, a []*rlt.StmAggWrsi, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltStmAggWrsiEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmAggWrsi (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltStmAggAlmaEql(t *testing.T, e, a *rlt.StmAggAlma, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "StmAggAlma.Id"))
	RltStmEql(t, e.Stm, a.Stm, append(msgs, "StmAggAlma.Stm"))
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmAggAlma.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmAggAlma.Vals"))
	UntEql(t, e.Length, a.Length, append(msgs, "StmAggAlma.Length"))
}
func RltStmAggAlmaNotZero(t *testing.T, a *rlt.StmAggAlma, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "StmAggAlma.Id"))
	RltStmNotZero(t, a.Stm, append(msgs, "StmAggAlma.Stm"))
}
func RltStmAggAlmaSliceEql(t *testing.T, e, a []*rlt.StmAggAlma, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltStmAggAlmaEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmAggAlma (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltStmAggVrncEql(t *testing.T, e, a *rlt.StmAggVrnc, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "StmAggVrnc.Id"))
	RltStmEql(t, e.Stm, a.Stm, append(msgs, "StmAggVrnc.Stm"))
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmAggVrnc.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmAggVrnc.Vals"))
	UntEql(t, e.Length, a.Length, append(msgs, "StmAggVrnc.Length"))
}
func RltStmAggVrncNotZero(t *testing.T, a *rlt.StmAggVrnc, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "StmAggVrnc.Id"))
	RltStmNotZero(t, a.Stm, append(msgs, "StmAggVrnc.Stm"))
}
func RltStmAggVrncSliceEql(t *testing.T, e, a []*rlt.StmAggVrnc, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltStmAggVrncEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmAggVrnc (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltStmAggStdEql(t *testing.T, e, a *rlt.StmAggStd, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "StmAggStd.Id"))
	RltStmEql(t, e.Stm, a.Stm, append(msgs, "StmAggStd.Stm"))
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmAggStd.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmAggStd.Vals"))
	UntEql(t, e.Length, a.Length, append(msgs, "StmAggStd.Length"))
}
func RltStmAggStdNotZero(t *testing.T, a *rlt.StmAggStd, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "StmAggStd.Id"))
	RltStmNotZero(t, a.Stm, append(msgs, "StmAggStd.Stm"))
}
func RltStmAggStdSliceEql(t *testing.T, e, a []*rlt.StmAggStd, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltStmAggStdEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmAggStd (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltStmAggRngFulEql(t *testing.T, e, a *rlt.StmAggRngFul, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "StmAggRngFul.Id"))
	RltStmEql(t, e.Stm, a.Stm, append(msgs, "StmAggRngFul.Stm"))
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmAggRngFul.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmAggRngFul.Vals"))
	UntEql(t, e.Length, a.Length, append(msgs, "StmAggRngFul.Length"))
}
func RltStmAggRngFulNotZero(t *testing.T, a *rlt.StmAggRngFul, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "StmAggRngFul.Id"))
	RltStmNotZero(t, a.Stm, append(msgs, "StmAggRngFul.Stm"))
}
func RltStmAggRngFulSliceEql(t *testing.T, e, a []*rlt.StmAggRngFul, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltStmAggRngFulEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmAggRngFul (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltStmAggRngLstEql(t *testing.T, e, a *rlt.StmAggRngLst, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "StmAggRngLst.Id"))
	RltStmEql(t, e.Stm, a.Stm, append(msgs, "StmAggRngLst.Stm"))
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmAggRngLst.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmAggRngLst.Vals"))
	UntEql(t, e.Length, a.Length, append(msgs, "StmAggRngLst.Length"))
}
func RltStmAggRngLstNotZero(t *testing.T, a *rlt.StmAggRngLst, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "StmAggRngLst.Id"))
	RltStmNotZero(t, a.Stm, append(msgs, "StmAggRngLst.Stm"))
}
func RltStmAggRngLstSliceEql(t *testing.T, e, a []*rlt.StmAggRngLst, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltStmAggRngLstEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmAggRngLst (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltStmAggProLstEql(t *testing.T, e, a *rlt.StmAggProLst, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "StmAggProLst.Id"))
	RltStmEql(t, e.Stm, a.Stm, append(msgs, "StmAggProLst.Stm"))
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmAggProLst.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmAggProLst.Vals"))
	UntEql(t, e.Length, a.Length, append(msgs, "StmAggProLst.Length"))
}
func RltStmAggProLstNotZero(t *testing.T, a *rlt.StmAggProLst, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "StmAggProLst.Id"))
	RltStmNotZero(t, a.Stm, append(msgs, "StmAggProLst.Stm"))
}
func RltStmAggProLstSliceEql(t *testing.T, e, a []*rlt.StmAggProLst, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltStmAggProLstEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmAggProLst (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltStmAggProSmaEql(t *testing.T, e, a *rlt.StmAggProSma, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "StmAggProSma.Id"))
	RltStmEql(t, e.Stm, a.Stm, append(msgs, "StmAggProSma.Stm"))
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmAggProSma.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmAggProSma.Vals"))
	UntEql(t, e.Length, a.Length, append(msgs, "StmAggProSma.Length"))
}
func RltStmAggProSmaNotZero(t *testing.T, a *rlt.StmAggProSma, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "StmAggProSma.Id"))
	RltStmNotZero(t, a.Stm, append(msgs, "StmAggProSma.Stm"))
}
func RltStmAggProSmaSliceEql(t *testing.T, e, a []*rlt.StmAggProSma, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltStmAggProSmaEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmAggProSma (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltStmAggProAlmaEql(t *testing.T, e, a *rlt.StmAggProAlma, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "StmAggProAlma.Id"))
	RltStmEql(t, e.Stm, a.Stm, append(msgs, "StmAggProAlma.Stm"))
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmAggProAlma.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmAggProAlma.Vals"))
	UntEql(t, e.Length, a.Length, append(msgs, "StmAggProAlma.Length"))
}
func RltStmAggProAlmaNotZero(t *testing.T, a *rlt.StmAggProAlma, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "StmAggProAlma.Id"))
	RltStmNotZero(t, a.Stm, append(msgs, "StmAggProAlma.Stm"))
}
func RltStmAggProAlmaSliceEql(t *testing.T, e, a []*rlt.StmAggProAlma, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltStmAggProAlmaEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmAggProAlma (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltStmAggEmaEql(t *testing.T, e, a *rlt.StmAggEma, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "StmAggEma.Id"))
	RltStmEql(t, e.Stm, a.Stm, append(msgs, "StmAggEma.Stm"))
	FltEql(t, e.Prv, a.Prv, append(msgs, "StmAggEma.Prv"))
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmAggEma.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmAggEma.Vals"))
	UntEql(t, e.Length, a.Length, append(msgs, "StmAggEma.Length"))
}
func RltStmAggEmaNotZero(t *testing.T, a *rlt.StmAggEma, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "StmAggEma.Id"))
	RltStmNotZero(t, a.Stm, append(msgs, "StmAggEma.Stm"))
}
func RltStmAggEmaSliceEql(t *testing.T, e, a []*rlt.StmAggEma, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltStmAggEmaEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmAggEma (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltStmInrAddEql(t *testing.T, e, a *rlt.StmInrAdd, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "StmInrAdd.Id"))
	RltStmEql(t, e.Stm, a.Stm, append(msgs, "StmInrAdd.Stm"))
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmInrAdd.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmInrAdd.Vals"))
	UntEql(t, e.Off, a.Off, append(msgs, "StmInrAdd.Off"))
}
func RltStmInrAddNotZero(t *testing.T, a *rlt.StmInrAdd, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "StmInrAdd.Id"))
	RltStmNotZero(t, a.Stm, append(msgs, "StmInrAdd.Stm"))
}
func RltStmInrAddSliceEql(t *testing.T, e, a []*rlt.StmInrAdd, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltStmInrAddEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmInrAdd (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltStmInrSubEql(t *testing.T, e, a *rlt.StmInrSub, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "StmInrSub.Id"))
	RltStmEql(t, e.Stm, a.Stm, append(msgs, "StmInrSub.Stm"))
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmInrSub.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmInrSub.Vals"))
	UntEql(t, e.Off, a.Off, append(msgs, "StmInrSub.Off"))
}
func RltStmInrSubNotZero(t *testing.T, a *rlt.StmInrSub, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "StmInrSub.Id"))
	RltStmNotZero(t, a.Stm, append(msgs, "StmInrSub.Stm"))
}
func RltStmInrSubSliceEql(t *testing.T, e, a []*rlt.StmInrSub, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltStmInrSubEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmInrSub (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltStmInrMulEql(t *testing.T, e, a *rlt.StmInrMul, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "StmInrMul.Id"))
	RltStmEql(t, e.Stm, a.Stm, append(msgs, "StmInrMul.Stm"))
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmInrMul.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmInrMul.Vals"))
	UntEql(t, e.Off, a.Off, append(msgs, "StmInrMul.Off"))
}
func RltStmInrMulNotZero(t *testing.T, a *rlt.StmInrMul, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "StmInrMul.Id"))
	RltStmNotZero(t, a.Stm, append(msgs, "StmInrMul.Stm"))
}
func RltStmInrMulSliceEql(t *testing.T, e, a []*rlt.StmInrMul, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltStmInrMulEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmInrMul (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltStmInrDivEql(t *testing.T, e, a *rlt.StmInrDiv, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "StmInrDiv.Id"))
	RltStmEql(t, e.Stm, a.Stm, append(msgs, "StmInrDiv.Stm"))
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmInrDiv.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmInrDiv.Vals"))
	UntEql(t, e.Off, a.Off, append(msgs, "StmInrDiv.Off"))
}
func RltStmInrDivNotZero(t *testing.T, a *rlt.StmInrDiv, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "StmInrDiv.Id"))
	RltStmNotZero(t, a.Stm, append(msgs, "StmInrDiv.Stm"))
}
func RltStmInrDivSliceEql(t *testing.T, e, a []*rlt.StmInrDiv, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltStmInrDivEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmInrDiv (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltStmInrRemEql(t *testing.T, e, a *rlt.StmInrRem, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "StmInrRem.Id"))
	RltStmEql(t, e.Stm, a.Stm, append(msgs, "StmInrRem.Stm"))
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmInrRem.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmInrRem.Vals"))
	UntEql(t, e.Off, a.Off, append(msgs, "StmInrRem.Off"))
}
func RltStmInrRemNotZero(t *testing.T, a *rlt.StmInrRem, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "StmInrRem.Id"))
	RltStmNotZero(t, a.Stm, append(msgs, "StmInrRem.Stm"))
}
func RltStmInrRemSliceEql(t *testing.T, e, a []*rlt.StmInrRem, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltStmInrRemEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmInrRem (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltStmInrPowEql(t *testing.T, e, a *rlt.StmInrPow, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "StmInrPow.Id"))
	RltStmEql(t, e.Stm, a.Stm, append(msgs, "StmInrPow.Stm"))
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmInrPow.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmInrPow.Vals"))
	UntEql(t, e.Off, a.Off, append(msgs, "StmInrPow.Off"))
}
func RltStmInrPowNotZero(t *testing.T, a *rlt.StmInrPow, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "StmInrPow.Id"))
	RltStmNotZero(t, a.Stm, append(msgs, "StmInrPow.Stm"))
}
func RltStmInrPowSliceEql(t *testing.T, e, a []*rlt.StmInrPow, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltStmInrPowEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmInrPow (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltStmInrMinEql(t *testing.T, e, a *rlt.StmInrMin, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "StmInrMin.Id"))
	RltStmEql(t, e.Stm, a.Stm, append(msgs, "StmInrMin.Stm"))
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmInrMin.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmInrMin.Vals"))
	UntEql(t, e.Off, a.Off, append(msgs, "StmInrMin.Off"))
}
func RltStmInrMinNotZero(t *testing.T, a *rlt.StmInrMin, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "StmInrMin.Id"))
	RltStmNotZero(t, a.Stm, append(msgs, "StmInrMin.Stm"))
}
func RltStmInrMinSliceEql(t *testing.T, e, a []*rlt.StmInrMin, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltStmInrMinEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmInrMin (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltStmInrMaxEql(t *testing.T, e, a *rlt.StmInrMax, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "StmInrMax.Id"))
	RltStmEql(t, e.Stm, a.Stm, append(msgs, "StmInrMax.Stm"))
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmInrMax.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmInrMax.Vals"))
	UntEql(t, e.Off, a.Off, append(msgs, "StmInrMax.Off"))
}
func RltStmInrMaxNotZero(t *testing.T, a *rlt.StmInrMax, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "StmInrMax.Id"))
	RltStmNotZero(t, a.Stm, append(msgs, "StmInrMax.Stm"))
}
func RltStmInrMaxSliceEql(t *testing.T, e, a []*rlt.StmInrMax, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltStmInrMaxEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmInrMax (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltStmInr1SlpEql(t *testing.T, e, a *rlt.StmInr1Slp, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "StmInr1Slp.Id"))
	RltStmEql(t, e.Stm, a.Stm, append(msgs, "StmInr1Slp.Stm"))
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmInr1Slp.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmInr1Slp.Vals"))
	UntEql(t, e.Off, a.Off, append(msgs, "StmInr1Slp.Off"))
}
func RltStmInr1SlpNotZero(t *testing.T, a *rlt.StmInr1Slp, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "StmInr1Slp.Id"))
	RltStmNotZero(t, a.Stm, append(msgs, "StmInr1Slp.Stm"))
}
func RltStmInr1SlpSliceEql(t *testing.T, e, a []*rlt.StmInr1Slp, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltStmInr1SlpEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmInr1Slp (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltStmOtrAddEql(t *testing.T, e, a *rlt.StmOtrAdd, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "StmOtrAdd.Id"))
	RltStmEql(t, e.Stm, a.Stm, append(msgs, "StmOtrAdd.Stm"))
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmOtrAdd.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmOtrAdd.Vals"))
	TmesEql(t, e.TmesA, a.TmesA, append(msgs, "StmOtrAdd.TmesA"))
	FltsEql(t, e.ValsA, a.ValsA, append(msgs, "StmOtrAdd.ValsA"))
	UntEql(t, e.Off, a.Off, append(msgs, "StmOtrAdd.Off"))
	RltStmEql(t, e.A, a.A, append(msgs, "StmOtrAdd.A"))
}
func RltStmOtrAddNotZero(t *testing.T, a *rlt.StmOtrAdd, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "StmOtrAdd.Id"))
	RltStmNotZero(t, a.Stm, append(msgs, "StmOtrAdd.Stm"))
}
func RltStmOtrAddSliceEql(t *testing.T, e, a []*rlt.StmOtrAdd, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltStmOtrAddEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmOtrAdd (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltStmOtrSubEql(t *testing.T, e, a *rlt.StmOtrSub, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "StmOtrSub.Id"))
	RltStmEql(t, e.Stm, a.Stm, append(msgs, "StmOtrSub.Stm"))
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmOtrSub.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmOtrSub.Vals"))
	TmesEql(t, e.TmesA, a.TmesA, append(msgs, "StmOtrSub.TmesA"))
	FltsEql(t, e.ValsA, a.ValsA, append(msgs, "StmOtrSub.ValsA"))
	UntEql(t, e.Off, a.Off, append(msgs, "StmOtrSub.Off"))
	RltStmEql(t, e.A, a.A, append(msgs, "StmOtrSub.A"))
}
func RltStmOtrSubNotZero(t *testing.T, a *rlt.StmOtrSub, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "StmOtrSub.Id"))
	RltStmNotZero(t, a.Stm, append(msgs, "StmOtrSub.Stm"))
}
func RltStmOtrSubSliceEql(t *testing.T, e, a []*rlt.StmOtrSub, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltStmOtrSubEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmOtrSub (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltStmOtrMulEql(t *testing.T, e, a *rlt.StmOtrMul, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "StmOtrMul.Id"))
	RltStmEql(t, e.Stm, a.Stm, append(msgs, "StmOtrMul.Stm"))
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmOtrMul.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmOtrMul.Vals"))
	TmesEql(t, e.TmesA, a.TmesA, append(msgs, "StmOtrMul.TmesA"))
	FltsEql(t, e.ValsA, a.ValsA, append(msgs, "StmOtrMul.ValsA"))
	UntEql(t, e.Off, a.Off, append(msgs, "StmOtrMul.Off"))
	RltStmEql(t, e.A, a.A, append(msgs, "StmOtrMul.A"))
}
func RltStmOtrMulNotZero(t *testing.T, a *rlt.StmOtrMul, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "StmOtrMul.Id"))
	RltStmNotZero(t, a.Stm, append(msgs, "StmOtrMul.Stm"))
}
func RltStmOtrMulSliceEql(t *testing.T, e, a []*rlt.StmOtrMul, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltStmOtrMulEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmOtrMul (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltStmOtrDivEql(t *testing.T, e, a *rlt.StmOtrDiv, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "StmOtrDiv.Id"))
	RltStmEql(t, e.Stm, a.Stm, append(msgs, "StmOtrDiv.Stm"))
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmOtrDiv.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmOtrDiv.Vals"))
	TmesEql(t, e.TmesA, a.TmesA, append(msgs, "StmOtrDiv.TmesA"))
	FltsEql(t, e.ValsA, a.ValsA, append(msgs, "StmOtrDiv.ValsA"))
	UntEql(t, e.Off, a.Off, append(msgs, "StmOtrDiv.Off"))
	RltStmEql(t, e.A, a.A, append(msgs, "StmOtrDiv.A"))
}
func RltStmOtrDivNotZero(t *testing.T, a *rlt.StmOtrDiv, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "StmOtrDiv.Id"))
	RltStmNotZero(t, a.Stm, append(msgs, "StmOtrDiv.Stm"))
}
func RltStmOtrDivSliceEql(t *testing.T, e, a []*rlt.StmOtrDiv, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltStmOtrDivEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmOtrDiv (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltStmOtrRemEql(t *testing.T, e, a *rlt.StmOtrRem, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "StmOtrRem.Id"))
	RltStmEql(t, e.Stm, a.Stm, append(msgs, "StmOtrRem.Stm"))
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmOtrRem.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmOtrRem.Vals"))
	TmesEql(t, e.TmesA, a.TmesA, append(msgs, "StmOtrRem.TmesA"))
	FltsEql(t, e.ValsA, a.ValsA, append(msgs, "StmOtrRem.ValsA"))
	UntEql(t, e.Off, a.Off, append(msgs, "StmOtrRem.Off"))
	RltStmEql(t, e.A, a.A, append(msgs, "StmOtrRem.A"))
}
func RltStmOtrRemNotZero(t *testing.T, a *rlt.StmOtrRem, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "StmOtrRem.Id"))
	RltStmNotZero(t, a.Stm, append(msgs, "StmOtrRem.Stm"))
}
func RltStmOtrRemSliceEql(t *testing.T, e, a []*rlt.StmOtrRem, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltStmOtrRemEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmOtrRem (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltStmOtrPowEql(t *testing.T, e, a *rlt.StmOtrPow, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "StmOtrPow.Id"))
	RltStmEql(t, e.Stm, a.Stm, append(msgs, "StmOtrPow.Stm"))
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmOtrPow.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmOtrPow.Vals"))
	TmesEql(t, e.TmesA, a.TmesA, append(msgs, "StmOtrPow.TmesA"))
	FltsEql(t, e.ValsA, a.ValsA, append(msgs, "StmOtrPow.ValsA"))
	UntEql(t, e.Off, a.Off, append(msgs, "StmOtrPow.Off"))
	RltStmEql(t, e.A, a.A, append(msgs, "StmOtrPow.A"))
}
func RltStmOtrPowNotZero(t *testing.T, a *rlt.StmOtrPow, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "StmOtrPow.Id"))
	RltStmNotZero(t, a.Stm, append(msgs, "StmOtrPow.Stm"))
}
func RltStmOtrPowSliceEql(t *testing.T, e, a []*rlt.StmOtrPow, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltStmOtrPowEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmOtrPow (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltStmOtrMinEql(t *testing.T, e, a *rlt.StmOtrMin, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "StmOtrMin.Id"))
	RltStmEql(t, e.Stm, a.Stm, append(msgs, "StmOtrMin.Stm"))
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmOtrMin.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmOtrMin.Vals"))
	TmesEql(t, e.TmesA, a.TmesA, append(msgs, "StmOtrMin.TmesA"))
	FltsEql(t, e.ValsA, a.ValsA, append(msgs, "StmOtrMin.ValsA"))
	UntEql(t, e.Off, a.Off, append(msgs, "StmOtrMin.Off"))
	RltStmEql(t, e.A, a.A, append(msgs, "StmOtrMin.A"))
}
func RltStmOtrMinNotZero(t *testing.T, a *rlt.StmOtrMin, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "StmOtrMin.Id"))
	RltStmNotZero(t, a.Stm, append(msgs, "StmOtrMin.Stm"))
}
func RltStmOtrMinSliceEql(t *testing.T, e, a []*rlt.StmOtrMin, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltStmOtrMinEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmOtrMin (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltStmOtrMaxEql(t *testing.T, e, a *rlt.StmOtrMax, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "StmOtrMax.Id"))
	RltStmEql(t, e.Stm, a.Stm, append(msgs, "StmOtrMax.Stm"))
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "StmOtrMax.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "StmOtrMax.Vals"))
	TmesEql(t, e.TmesA, a.TmesA, append(msgs, "StmOtrMax.TmesA"))
	FltsEql(t, e.ValsA, a.ValsA, append(msgs, "StmOtrMax.ValsA"))
	UntEql(t, e.Off, a.Off, append(msgs, "StmOtrMax.Off"))
	RltStmEql(t, e.A, a.A, append(msgs, "StmOtrMax.A"))
}
func RltStmOtrMaxNotZero(t *testing.T, a *rlt.StmOtrMax, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "StmOtrMax.Id"))
	RltStmNotZero(t, a.Stm, append(msgs, "StmOtrMax.Stm"))
}
func RltStmOtrMaxSliceEql(t *testing.T, e, a []*rlt.StmOtrMax, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltStmOtrMaxEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StmOtrMax (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltCndSclEqlEql(t *testing.T, e, a *rlt.CndSclEql, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "CndSclEql.Id"))
	RltStmEql(t, e.Stm, a.Stm, append(msgs, "CndSclEql.Stm"))
	FltEql(t, e.Scl, a.Scl, append(msgs, "CndSclEql.Scl"))
}
func RltCndSclEqlNotZero(t *testing.T, a *rlt.CndSclEql, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "CndSclEql.Id"))
	RltStmNotZero(t, a.Stm, append(msgs, "CndSclEql.Stm"))
}
func RltCndSclEqlSliceEql(t *testing.T, e, a []*rlt.CndSclEql, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltCndSclEqlEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm CndSclEql (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltCndSclNeqEql(t *testing.T, e, a *rlt.CndSclNeq, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "CndSclNeq.Id"))
	RltStmEql(t, e.Stm, a.Stm, append(msgs, "CndSclNeq.Stm"))
	FltEql(t, e.Scl, a.Scl, append(msgs, "CndSclNeq.Scl"))
}
func RltCndSclNeqNotZero(t *testing.T, a *rlt.CndSclNeq, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "CndSclNeq.Id"))
	RltStmNotZero(t, a.Stm, append(msgs, "CndSclNeq.Stm"))
}
func RltCndSclNeqSliceEql(t *testing.T, e, a []*rlt.CndSclNeq, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltCndSclNeqEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm CndSclNeq (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltCndSclLssEql(t *testing.T, e, a *rlt.CndSclLss, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "CndSclLss.Id"))
	RltStmEql(t, e.Stm, a.Stm, append(msgs, "CndSclLss.Stm"))
	FltEql(t, e.Scl, a.Scl, append(msgs, "CndSclLss.Scl"))
}
func RltCndSclLssNotZero(t *testing.T, a *rlt.CndSclLss, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "CndSclLss.Id"))
	RltStmNotZero(t, a.Stm, append(msgs, "CndSclLss.Stm"))
}
func RltCndSclLssSliceEql(t *testing.T, e, a []*rlt.CndSclLss, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltCndSclLssEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm CndSclLss (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltCndSclGtrEql(t *testing.T, e, a *rlt.CndSclGtr, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "CndSclGtr.Id"))
	RltStmEql(t, e.Stm, a.Stm, append(msgs, "CndSclGtr.Stm"))
	FltEql(t, e.Scl, a.Scl, append(msgs, "CndSclGtr.Scl"))
}
func RltCndSclGtrNotZero(t *testing.T, a *rlt.CndSclGtr, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "CndSclGtr.Id"))
	RltStmNotZero(t, a.Stm, append(msgs, "CndSclGtr.Stm"))
}
func RltCndSclGtrSliceEql(t *testing.T, e, a []*rlt.CndSclGtr, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltCndSclGtrEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm CndSclGtr (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltCndSclLeqEql(t *testing.T, e, a *rlt.CndSclLeq, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "CndSclLeq.Id"))
	RltStmEql(t, e.Stm, a.Stm, append(msgs, "CndSclLeq.Stm"))
	FltEql(t, e.Scl, a.Scl, append(msgs, "CndSclLeq.Scl"))
}
func RltCndSclLeqNotZero(t *testing.T, a *rlt.CndSclLeq, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "CndSclLeq.Id"))
	RltStmNotZero(t, a.Stm, append(msgs, "CndSclLeq.Stm"))
}
func RltCndSclLeqSliceEql(t *testing.T, e, a []*rlt.CndSclLeq, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltCndSclLeqEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm CndSclLeq (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltCndSclGeqEql(t *testing.T, e, a *rlt.CndSclGeq, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "CndSclGeq.Id"))
	RltStmEql(t, e.Stm, a.Stm, append(msgs, "CndSclGeq.Stm"))
	FltEql(t, e.Scl, a.Scl, append(msgs, "CndSclGeq.Scl"))
}
func RltCndSclGeqNotZero(t *testing.T, a *rlt.CndSclGeq, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "CndSclGeq.Id"))
	RltStmNotZero(t, a.Stm, append(msgs, "CndSclGeq.Stm"))
}
func RltCndSclGeqSliceEql(t *testing.T, e, a []*rlt.CndSclGeq, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltCndSclGeqEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm CndSclGeq (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltCndInrEqlEql(t *testing.T, e, a *rlt.CndInrEql, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "CndInrEql.Id"))
	RltStmEql(t, e.Stm, a.Stm, append(msgs, "CndInrEql.Stm"))
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "CndInrEql.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "CndInrEql.Vals"))
	UntEql(t, e.Off, a.Off, append(msgs, "CndInrEql.Off"))
}
func RltCndInrEqlNotZero(t *testing.T, a *rlt.CndInrEql, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "CndInrEql.Id"))
	RltStmNotZero(t, a.Stm, append(msgs, "CndInrEql.Stm"))
}
func RltCndInrEqlSliceEql(t *testing.T, e, a []*rlt.CndInrEql, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltCndInrEqlEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm CndInrEql (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltCndInrNeqEql(t *testing.T, e, a *rlt.CndInrNeq, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "CndInrNeq.Id"))
	RltStmEql(t, e.Stm, a.Stm, append(msgs, "CndInrNeq.Stm"))
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "CndInrNeq.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "CndInrNeq.Vals"))
	UntEql(t, e.Off, a.Off, append(msgs, "CndInrNeq.Off"))
}
func RltCndInrNeqNotZero(t *testing.T, a *rlt.CndInrNeq, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "CndInrNeq.Id"))
	RltStmNotZero(t, a.Stm, append(msgs, "CndInrNeq.Stm"))
}
func RltCndInrNeqSliceEql(t *testing.T, e, a []*rlt.CndInrNeq, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltCndInrNeqEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm CndInrNeq (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltCndInrLssEql(t *testing.T, e, a *rlt.CndInrLss, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "CndInrLss.Id"))
	RltStmEql(t, e.Stm, a.Stm, append(msgs, "CndInrLss.Stm"))
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "CndInrLss.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "CndInrLss.Vals"))
	UntEql(t, e.Off, a.Off, append(msgs, "CndInrLss.Off"))
}
func RltCndInrLssNotZero(t *testing.T, a *rlt.CndInrLss, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "CndInrLss.Id"))
	RltStmNotZero(t, a.Stm, append(msgs, "CndInrLss.Stm"))
}
func RltCndInrLssSliceEql(t *testing.T, e, a []*rlt.CndInrLss, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltCndInrLssEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm CndInrLss (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltCndInrGtrEql(t *testing.T, e, a *rlt.CndInrGtr, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "CndInrGtr.Id"))
	RltStmEql(t, e.Stm, a.Stm, append(msgs, "CndInrGtr.Stm"))
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "CndInrGtr.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "CndInrGtr.Vals"))
	UntEql(t, e.Off, a.Off, append(msgs, "CndInrGtr.Off"))
}
func RltCndInrGtrNotZero(t *testing.T, a *rlt.CndInrGtr, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "CndInrGtr.Id"))
	RltStmNotZero(t, a.Stm, append(msgs, "CndInrGtr.Stm"))
}
func RltCndInrGtrSliceEql(t *testing.T, e, a []*rlt.CndInrGtr, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltCndInrGtrEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm CndInrGtr (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltCndInrLeqEql(t *testing.T, e, a *rlt.CndInrLeq, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "CndInrLeq.Id"))
	RltStmEql(t, e.Stm, a.Stm, append(msgs, "CndInrLeq.Stm"))
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "CndInrLeq.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "CndInrLeq.Vals"))
	UntEql(t, e.Off, a.Off, append(msgs, "CndInrLeq.Off"))
}
func RltCndInrLeqNotZero(t *testing.T, a *rlt.CndInrLeq, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "CndInrLeq.Id"))
	RltStmNotZero(t, a.Stm, append(msgs, "CndInrLeq.Stm"))
}
func RltCndInrLeqSliceEql(t *testing.T, e, a []*rlt.CndInrLeq, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltCndInrLeqEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm CndInrLeq (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltCndInrGeqEql(t *testing.T, e, a *rlt.CndInrGeq, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "CndInrGeq.Id"))
	RltStmEql(t, e.Stm, a.Stm, append(msgs, "CndInrGeq.Stm"))
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "CndInrGeq.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "CndInrGeq.Vals"))
	UntEql(t, e.Off, a.Off, append(msgs, "CndInrGeq.Off"))
}
func RltCndInrGeqNotZero(t *testing.T, a *rlt.CndInrGeq, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "CndInrGeq.Id"))
	RltStmNotZero(t, a.Stm, append(msgs, "CndInrGeq.Stm"))
}
func RltCndInrGeqSliceEql(t *testing.T, e, a []*rlt.CndInrGeq, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltCndInrGeqEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm CndInrGeq (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltCndOtrEqlEql(t *testing.T, e, a *rlt.CndOtrEql, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "CndOtrEql.Id"))
	RltStmEql(t, e.Stm, a.Stm, append(msgs, "CndOtrEql.Stm"))
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "CndOtrEql.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "CndOtrEql.Vals"))
	TmesEql(t, e.TmesA, a.TmesA, append(msgs, "CndOtrEql.TmesA"))
	FltsEql(t, e.ValsA, a.ValsA, append(msgs, "CndOtrEql.ValsA"))
	UntEql(t, e.Off, a.Off, append(msgs, "CndOtrEql.Off"))
	RltStmEql(t, e.A, a.A, append(msgs, "CndOtrEql.A"))
}
func RltCndOtrEqlNotZero(t *testing.T, a *rlt.CndOtrEql, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "CndOtrEql.Id"))
	RltStmNotZero(t, a.Stm, append(msgs, "CndOtrEql.Stm"))
}
func RltCndOtrEqlSliceEql(t *testing.T, e, a []*rlt.CndOtrEql, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltCndOtrEqlEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm CndOtrEql (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltCndOtrNeqEql(t *testing.T, e, a *rlt.CndOtrNeq, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "CndOtrNeq.Id"))
	RltStmEql(t, e.Stm, a.Stm, append(msgs, "CndOtrNeq.Stm"))
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "CndOtrNeq.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "CndOtrNeq.Vals"))
	TmesEql(t, e.TmesA, a.TmesA, append(msgs, "CndOtrNeq.TmesA"))
	FltsEql(t, e.ValsA, a.ValsA, append(msgs, "CndOtrNeq.ValsA"))
	UntEql(t, e.Off, a.Off, append(msgs, "CndOtrNeq.Off"))
	RltStmEql(t, e.A, a.A, append(msgs, "CndOtrNeq.A"))
}
func RltCndOtrNeqNotZero(t *testing.T, a *rlt.CndOtrNeq, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "CndOtrNeq.Id"))
	RltStmNotZero(t, a.Stm, append(msgs, "CndOtrNeq.Stm"))
}
func RltCndOtrNeqSliceEql(t *testing.T, e, a []*rlt.CndOtrNeq, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltCndOtrNeqEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm CndOtrNeq (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltCndOtrLssEql(t *testing.T, e, a *rlt.CndOtrLss, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "CndOtrLss.Id"))
	RltStmEql(t, e.Stm, a.Stm, append(msgs, "CndOtrLss.Stm"))
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "CndOtrLss.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "CndOtrLss.Vals"))
	TmesEql(t, e.TmesA, a.TmesA, append(msgs, "CndOtrLss.TmesA"))
	FltsEql(t, e.ValsA, a.ValsA, append(msgs, "CndOtrLss.ValsA"))
	UntEql(t, e.Off, a.Off, append(msgs, "CndOtrLss.Off"))
	RltStmEql(t, e.A, a.A, append(msgs, "CndOtrLss.A"))
}
func RltCndOtrLssNotZero(t *testing.T, a *rlt.CndOtrLss, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "CndOtrLss.Id"))
	RltStmNotZero(t, a.Stm, append(msgs, "CndOtrLss.Stm"))
}
func RltCndOtrLssSliceEql(t *testing.T, e, a []*rlt.CndOtrLss, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltCndOtrLssEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm CndOtrLss (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltCndOtrGtrEql(t *testing.T, e, a *rlt.CndOtrGtr, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "CndOtrGtr.Id"))
	RltStmEql(t, e.Stm, a.Stm, append(msgs, "CndOtrGtr.Stm"))
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "CndOtrGtr.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "CndOtrGtr.Vals"))
	TmesEql(t, e.TmesA, a.TmesA, append(msgs, "CndOtrGtr.TmesA"))
	FltsEql(t, e.ValsA, a.ValsA, append(msgs, "CndOtrGtr.ValsA"))
	UntEql(t, e.Off, a.Off, append(msgs, "CndOtrGtr.Off"))
	RltStmEql(t, e.A, a.A, append(msgs, "CndOtrGtr.A"))
}
func RltCndOtrGtrNotZero(t *testing.T, a *rlt.CndOtrGtr, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "CndOtrGtr.Id"))
	RltStmNotZero(t, a.Stm, append(msgs, "CndOtrGtr.Stm"))
}
func RltCndOtrGtrSliceEql(t *testing.T, e, a []*rlt.CndOtrGtr, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltCndOtrGtrEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm CndOtrGtr (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltCndOtrLeqEql(t *testing.T, e, a *rlt.CndOtrLeq, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "CndOtrLeq.Id"))
	RltStmEql(t, e.Stm, a.Stm, append(msgs, "CndOtrLeq.Stm"))
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "CndOtrLeq.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "CndOtrLeq.Vals"))
	TmesEql(t, e.TmesA, a.TmesA, append(msgs, "CndOtrLeq.TmesA"))
	FltsEql(t, e.ValsA, a.ValsA, append(msgs, "CndOtrLeq.ValsA"))
	UntEql(t, e.Off, a.Off, append(msgs, "CndOtrLeq.Off"))
	RltStmEql(t, e.A, a.A, append(msgs, "CndOtrLeq.A"))
}
func RltCndOtrLeqNotZero(t *testing.T, a *rlt.CndOtrLeq, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "CndOtrLeq.Id"))
	RltStmNotZero(t, a.Stm, append(msgs, "CndOtrLeq.Stm"))
}
func RltCndOtrLeqSliceEql(t *testing.T, e, a []*rlt.CndOtrLeq, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltCndOtrLeqEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm CndOtrLeq (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltCndOtrGeqEql(t *testing.T, e, a *rlt.CndOtrGeq, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "CndOtrGeq.Id"))
	RltStmEql(t, e.Stm, a.Stm, append(msgs, "CndOtrGeq.Stm"))
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "CndOtrGeq.Tmes"))
	FltsEql(t, e.Vals, a.Vals, append(msgs, "CndOtrGeq.Vals"))
	TmesEql(t, e.TmesA, a.TmesA, append(msgs, "CndOtrGeq.TmesA"))
	FltsEql(t, e.ValsA, a.ValsA, append(msgs, "CndOtrGeq.ValsA"))
	UntEql(t, e.Off, a.Off, append(msgs, "CndOtrGeq.Off"))
	RltStmEql(t, e.A, a.A, append(msgs, "CndOtrGeq.A"))
}
func RltCndOtrGeqNotZero(t *testing.T, a *rlt.CndOtrGeq, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "CndOtrGeq.Id"))
	RltStmNotZero(t, a.Stm, append(msgs, "CndOtrGeq.Stm"))
}
func RltCndOtrGeqSliceEql(t *testing.T, e, a []*rlt.CndOtrGeq, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltCndOtrGeqEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm CndOtrGeq (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltCndCnd1AndEql(t *testing.T, e, a *rlt.CndCnd1And, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "CndCnd1And.Id"))
	RltCndEql(t, e.Cnd, a.Cnd, append(msgs, "CndCnd1And.Cnd"))
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "CndCnd1And.Tmes"))
	TmesEql(t, e.TmesA, a.TmesA, append(msgs, "CndCnd1And.TmesA"))
	RltCndEql(t, e.A, a.A, append(msgs, "CndCnd1And.A"))
}
func RltCndCnd1AndNotZero(t *testing.T, a *rlt.CndCnd1And, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "CndCnd1And.Id"))
	RltCndNotZero(t, a.Cnd, append(msgs, "CndCnd1And.Cnd"))
}
func RltCndCnd1AndSliceEql(t *testing.T, e, a []*rlt.CndCnd1And, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltCndCnd1AndEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm CndCnd1And (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltCndCnd2SeqEql(t *testing.T, e, a *rlt.CndCnd2Seq, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "CndCnd2Seq.Id"))
	RltCndEql(t, e.Cnd, a.Cnd, append(msgs, "CndCnd2Seq.Cnd"))
	TmesEql(t, e.Tmes, a.Tmes, append(msgs, "CndCnd2Seq.Tmes"))
	TmesEql(t, e.TmesA, a.TmesA, append(msgs, "CndCnd2Seq.TmesA"))
	TmeEql(t, e.Dur, a.Dur, append(msgs, "CndCnd2Seq.Dur"))
	RltCndEql(t, e.A, a.A, append(msgs, "CndCnd2Seq.A"))
}
func RltCndCnd2SeqNotZero(t *testing.T, a *rlt.CndCnd2Seq, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "CndCnd2Seq.Id"))
	RltCndNotZero(t, a.Cnd, append(msgs, "CndCnd2Seq.Cnd"))
}
func RltCndCnd2SeqSliceEql(t *testing.T, e, a []*rlt.CndCnd2Seq, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltCndCnd2SeqEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm CndCnd2Seq (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func RltStgyStgyEql(t *testing.T, e, a *rlt.StgyStgy, msgs ...interface{}) {
	if e == nil && a == nil {
		return
	}
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("e is nil"))...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("a is nil"))...)
	}
	Uint32Eql(t, e.Id, a.Id, append(msgs, "StgyStgy.Id"))
	RltCndEql(t, e.Cnd, a.Cnd, append(msgs, "StgyStgy.Cnd"))
	BolEql(t, e.IsLong, a.IsLong, append(msgs, "StgyStgy.IsLong"))
	FltEql(t, e.PrfLim, a.PrfLim, append(msgs, "StgyStgy.PrfLim"))
	FltEql(t, e.LosLim, a.LosLim, append(msgs, "StgyStgy.LosLim"))
	TmeEql(t, e.DurLim, a.DurLim, append(msgs, "StgyStgy.DurLim"))
	FltEql(t, e.MinPnlPct, a.MinPnlPct, append(msgs, "StgyStgy.MinPnlPct"))
	RltInstrEql(t, e.Instr, a.Instr, append(msgs, "StgyStgy.Instr"))
	RltStmsEql(t, e.FtrStms, a.FtrStms, append(msgs, "StgyStgy.FtrStms"))
	RltCndSliceEql(t, e.Clss, a.Clss, append(msgs, "StgyStgy.Clss"))
	FltEql(t, e.ClsPrfLim, a.ClsPrfLim, append(msgs, "StgyStgy.ClsPrfLim"))
	FltEql(t, e.ClsLosLim, a.ClsLosLim, append(msgs, "StgyStgy.ClsLosLim"))
	TmeEql(t, e.ClsTmeLim, a.ClsTmeLim, append(msgs, "StgyStgy.ClsTmeLim"))
	TmeEql(t, e.LstClsTme, a.LstClsTme, append(msgs, "StgyStgy.LstClsTme"))
	UntEql(t, e.LstClsIdx, a.LstClsIdx, append(msgs, "StgyStgy.LstClsIdx"))
	UntEql(t, e.OpnIdx, a.OpnIdx, append(msgs, "StgyStgy.OpnIdx"))
	AnaTrdEql(t, e.Trd, a.Trd, append(msgs, "StgyStgy.Trd"))
	StringEql(t, e.Key, a.Key, append(msgs, "StgyStgy.Key"))
}
func RltStgyStgyNotZero(t *testing.T, a *rlt.StgyStgy, msgs ...interface{}) {
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "is nil")...)
	}
	Uint32NotZero(t, a.Id, append(msgs, "StgyStgy.Id"))
	RltCndNotZero(t, a.Cnd, append(msgs, "StgyStgy.Cnd"))
	FltNotZero(t, a.PrfLim, append(msgs, "StgyStgy.PrfLim"))
	FltNotZero(t, a.LosLim, append(msgs, "StgyStgy.LosLim"))
	TmeNotZero(t, a.DurLim, append(msgs, "StgyStgy.DurLim"))
	RltInstrNotZero(t, a.Instr, append(msgs, "StgyStgy.Instr"))
	RltStmsNotZero(t, a.FtrStms, append(msgs, "StgyStgy.FtrStms"))
}
func RltStgyStgySliceEql(t *testing.T, e, a []*rlt.StgyStgy, msgs ...interface{}) {
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("length not equal (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] == nil && a[n] == nil {
			continue
		}
		if e[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("e elm is nil: idx %v nil", n))...)
		}
		if a[n] == nil {
			t.Helper()
			t.Fatal(append(msgs, fmt.Sprintf("a elm is nil: idx %v nil", n))...)
		}
		RltStgyStgyEql(t, e[n], a[n], append(msgs, fmt.Sprintf("elm StgyStgy (idx:%v expected:%v actual:%v)", n, e[n], a[n])))
	}
}
func (x HstPrvInstr) String() string    { return "HstPrvInstr" }
func (x HstInstrInrvl) String() string  { return "HstInstrInrvl" }
func (x HstInrvlSide) String() string   { return "HstInrvlSide" }
func (x HstSideStmRte) String() string  { return "HstSideStmRte" }
func (x HstSideStmRte1) String() string { return "HstSideStmRte1" }
func (x HstStmStmUna) String() string   { return "HstStmStmUna" }
func (x HstStmStmScl) String() string   { return "HstStmStmScl" }
func (x HstStmStmSel) String() string   { return "HstStmStmSel" }
func (x HstStmStmAgg) String() string   { return "HstStmStmAgg" }
func (x HstStmStmInr) String() string   { return "HstStmStmInr" }
func (x HstStmStmInr1) String() string  { return "HstStmStmInr1" }
func (x HstStmStmOtr) String() string   { return "HstStmStmOtr" }
func (x HstStmCndScl) String() string   { return "HstStmCndScl" }
func (x HstStmCndInr) String() string   { return "HstStmCndInr" }
func (x HstStmCndOtr) String() string   { return "HstStmCndOtr" }
func (x HstCndCndCnd1) String() string  { return "HstCndCndCnd1" }
func (x HstCndCndCnd2) String() string  { return "HstCndCndCnd2" }
func (x HstCndStgy) String() string     { return "HstCndStgy" }
func (x RltPrvInstr) String() string    { return "RltPrvInstr" }
func (x RltInstrInrvl) String() string  { return "RltInstrInrvl" }
func (x RltInrvlSide) String() string   { return "RltInrvlSide" }
func (x RltSideStmRte) String() string  { return "RltSideStmRte" }
func (x RltSideStmRte1) String() string { return "RltSideStmRte1" }
func (x RltStmStmUna) String() string   { return "RltStmStmUna" }
func (x RltStmStmScl) String() string   { return "RltStmStmScl" }
func (x RltStmStmSel) String() string   { return "RltStmStmSel" }
func (x RltStmStmAgg) String() string   { return "RltStmStmAgg" }
func (x RltStmStmInr) String() string   { return "RltStmStmInr" }
func (x RltStmStmInr1) String() string  { return "RltStmStmInr1" }
func (x RltStmStmOtr) String() string   { return "RltStmStmOtr" }
func (x RltStmCndScl) String() string   { return "RltStmCndScl" }
func (x RltStmCndInr) String() string   { return "RltStmCndInr" }
func (x RltStmCndOtr) String() string   { return "RltStmCndOtr" }
func (x RltCndCndCnd1) String() string  { return "RltCndCndCnd1" }
func (x RltCndCndCnd2) String() string  { return "RltCndCndCnd2" }
func (x RltCndStgy) String() string     { return "RltCndStgy" }
