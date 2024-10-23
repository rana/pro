package srch

import (
	"sys/ana/hst"
	"sys/bsc/flt"
	"sys/bsc/tme"
	"sys/bsc/tmes"
)

type (
	Instr    func(tme.Rng) hst.Instr
	Side     func(instr hst.Instr, dur tme.Tme) hst.Side
	Sides    func(instr hst.Instr, durs *tmes.Tmes) *hst.Sides
	SideStm  func(hst.Side) hst.Stm
	SideStms func(*hst.Sides) *hst.Stms
	OpnCnd   func(hst.Instr) hst.Cnd
	Stgy     func(opnCnd hst.Cnd, prfLim, losLim flt.Flt, durLim tme.Tme, instr hst.Instr, clss ...hst.Cnd) hst.Stgy

	DurSideStm struct {
		Dur     tme.Tme
		SideStm SideStm
	}
)

// OPNCNDS
func OpnCndMacdCrsLong(instr hst.Instr) hst.Cnd {
	h1AskLst := instr.I(tme.H1).Ask().Lst()
	qckStm := h1AskLst.AggEma(12).OtrSub(0, h1AskLst.AggEma(26))
	slwStm := qckStm.AggEma(9)
	opnCnd := slwStm.SclGtr(0.0).And(qckStm.OtrLss(0, slwStm).Seq(tme.S1, qckStm.OtrGtr(0, slwStm)))
	return opnCnd
}
func OpnCndAlmaSmaCrsUp520(instr hst.Instr) hst.Cnd {
	qckStm := instr.I(tme.M5).Ask().Alma()
	slwStm := instr.I(tme.M20).Ask().Sma()
	opnCnd := qckStm.OtrLss(0, slwStm).Seq(tme.S1, qckStm.OtrGtr(0, slwStm))
	return opnCnd
}
func OpnCndAlmaSmaCrsUp50200(instr hst.Instr) hst.Cnd {
	qckStm := instr.I(tme.M1 * 50).Ask().Alma()
	slwStm := instr.I(tme.M1 * 200).Ask().Sma()
	opnCnd := qckStm.OtrLss(0, slwStm).Seq(tme.S1, qckStm.OtrGtr(0, slwStm))
	return opnCnd
}
func OpnCndAlmaSmaCrsUpH1H4(instr hst.Instr) hst.Cnd {
	qckStm := instr.I(tme.H1).Ask().Alma()
	slwStm := instr.I(tme.H1 * 4).Ask().Sma()
	opnCnd := qckStm.OtrLss(0, slwStm).Seq(tme.S1, qckStm.OtrGtr(0, slwStm))
	return opnCnd
}
func OpnCndBlngrLwrCrsUp(instr hst.Instr, qckDur, sloDur tme.Tme) hst.Cnd {
	qckStm := instr.I(qckDur).Ask().Lst()

	m20Ask := instr.I(sloDur).Ask()
	m20Sma := m20Ask.Sma()
	m20Std := m20Ask.Std().SclMul(2.0)
	blngrLwr := m20Sma.OtrSub(0, m20Std)

	opnCnd := qckStm.OtrLss(0, blngrLwr).Seq(tme.S1, qckStm.OtrGtr(0, blngrLwr))
	return opnCnd
}
func OpnCndBlngrUprCrsUp(instr hst.Instr, qckDur, sloDur tme.Tme) hst.Cnd {
	qckStm := instr.I(qckDur).Ask().Lst()

	m20Ask := instr.I(sloDur).Ask()
	m20Sma := m20Ask.Sma()
	m20Std := m20Ask.Std().SclMul(2.0)
	blngrLwr := m20Sma.OtrAdd(0, m20Std)

	opnCnd := qckStm.OtrLss(0, blngrLwr).Seq(tme.S1, qckStm.OtrGtr(0, blngrLwr))
	return opnCnd
}
func OpnCndBlngrLwrCrsDwn(instr hst.Instr, qckDur, sloDur tme.Tme) hst.Cnd {
	qckStm := instr.I(qckDur).Bid().Lst()

	m20Ask := instr.I(sloDur).Bid()
	m20Sma := m20Ask.Sma()
	m20Std := m20Ask.Std().SclMul(2.0)
	blngrLwr := m20Sma.OtrSub(0, m20Std)

	opnCnd := qckStm.OtrGtr(0, blngrLwr).Seq(tme.S1, qckStm.OtrLss(0, blngrLwr))
	return opnCnd
}
func OpnCndLstSmaCrsUp15(instr hst.Instr) hst.Cnd {
	qckStm := instr.I(tme.M1).Ask().Lst()
	slwStm := instr.I(tme.M5).Ask().Sma()
	opnCnd := qckStm.OtrLss(0, slwStm).Seq(tme.S1, qckStm.OtrGtr(0, slwStm))
	return opnCnd
}

// // STGYS
// func Long(opnCnd hst.Cnd, prfLim, losLim flt.Flt, durLim tme.Tme, instr hst.Instr, clss ...hst.Cnd) hst.Stgy {
// 	return opnCnd.Long(prfLim, losLim, durLim, instr, clss...)
// }
// func Shrt(opnCnd hst.Cnd, prfLim, losLim flt.Flt, durLim tme.Tme, instr hst.Instr, clss ...hst.Cnd) hst.Stgy {
// 	return opnCnd.Shrt(prfLim, losLim, durLim, instr, clss...)
// }
// func LongPll(opnCnd hst.Cnd, prfLim, losLim flt.Flt, durLim tme.Tme, instr hst.Instr, clss ...hst.Cnd) hst.Stgy {
// 	return opnCnd.LongPll(prfLim, losLim, durLim, instr, clss...)
// }

// INSTRS
func EurUsd(bckRng tme.Rng) hst.Instr { return hst.Oan().EurUsd(bckRng) }
func AudUsd(bckRng tme.Rng) hst.Instr { return hst.Oan().AudUsd(bckRng) }
func NzdUsd(bckRng tme.Rng) hst.Instr { return hst.Oan().NzdUsd(bckRng) }
func GbpUsd(bckRng tme.Rng) hst.Instr { return hst.Oan().GbpUsd(bckRng) }

// // SIDE
// func Ask(instr hst.Instr, dur tme.Tme) hst.Side { return instr.I(dur).Ask() }
// func Bid(instr hst.Instr, dur tme.Tme) hst.Side { return instr.I(dur).Ask() }

// // SIDES
// func DurAsks(instr hst.Instr, durs *tmes.Tmes) (r *hst.Sides) {
// 	r = hst.NewSides()
// 	for _, dur := range *durs {
// 		r.Push(instr.I(dur).Ask())
// 	}
// 	return r
// }

// // SIDESTM
// func Vrnc(side hst.Side) hst.Stm {
// 	return side.Vrnc()
// }
// func Rsi(side hst.Side) hst.Stm {
// 	return side.Rsi()
// }
// func Wrsi(side hst.Side) hst.Stm {
// 	return side.Wrsi()
// }
// func Ema(side hst.Side) hst.Stm {
// 	return side.Ema()
// }
// func Wma(side hst.Side) hst.Stm {
// 	return side.Wma()
// }
// func Sma(side hst.Side) hst.Stm {
// 	return side.Sma()
// }
// func Alma(side hst.Side) hst.Stm {
// 	return side.Alma()
// }
// func RngFul(side hst.Side) hst.Stm {
// 	return side.RngFul()
// }
// func RngLst(side hst.Side) hst.Stm {
// 	return side.RngLst()
// }
// func Min(side hst.Side) hst.Stm {
// 	return side.Min()
// }
// func Max(side hst.Side) hst.Stm {
// 	return side.Max()
// }
// func Lst(side hst.Side) hst.Stm {
// 	return side.Lst()
// }
// func ProLst(side hst.Side) hst.Stm {
// 	return side.ProLst()
// }
// func ProSma(side hst.Side) hst.Stm {
// 	return side.ProSma()
// }
// func ProAlma(side hst.Side) hst.Stm {
// 	return side.ProAlma()
// }

// var (
// 	VSideStms = []SideStm{
// 		Vrnc, Rsi, Wrsi,
// 		Ema, Wma, Sma, Alma,
// 		RngFul, RngLst,
// 		Min, Max, Lst,
// 		ProLst, ProSma, ProAlma,
// 	}
// )

// func Vrncs(sides *hst.Sides) (r *hst.Stms) {
// 	r = hst.NewStms()
// 	for _, side := range *sides {
// 		r.Push(side.Vrnc())
// 	}
// 	return r
// }
// func Rsis(sides *hst.Sides) (r *hst.Stms) {
// 	r = hst.NewStms()
// 	for _, side := range *sides {
// 		r.Push(side.Rsi())
// 	}
// 	return r
// }
// func Wrsis(sides *hst.Sides) (r *hst.Stms) {
// 	r = hst.NewStms()
// 	for _, side := range *sides {
// 		r.Push(side.Wrsi())
// 	}
// 	return r
// }
// func Emas(sides *hst.Sides) (r *hst.Stms) {
// 	r = hst.NewStms()
// 	for _, side := range *sides {
// 		r.Push(side.Ema())
// 	}
// 	return r
// }
// func Wmas(sides *hst.Sides) (r *hst.Stms) {
// 	r = hst.NewStms()
// 	for _, side := range *sides {
// 		r.Push(side.Wma())
// 	}
// 	return r
// }
// func Smas(sides *hst.Sides) (r *hst.Stms) {
// 	r = hst.NewStms()
// 	for _, side := range *sides {
// 		r.Push(side.Sma())
// 	}
// 	return r
// }
// func Almas(sides *hst.Sides) (r *hst.Stms) {
// 	r = hst.NewStms()
// 	for _, side := range *sides {
// 		r.Push(side.Alma())
// 	}
// 	return r
// }
// func RngFuls(sides *hst.Sides) (r *hst.Stms) {
// 	r = hst.NewStms()
// 	for _, side := range *sides {
// 		r.Push(side.RngFul())
// 	}
// 	return r
// }
// func RngLsts(sides *hst.Sides) (r *hst.Stms) {
// 	r = hst.NewStms()
// 	for _, side := range *sides {
// 		r.Push(side.RngLst())
// 	}
// 	return r
// }
// func Mins(sides *hst.Sides) (r *hst.Stms) {
// 	r = hst.NewStms()
// 	for _, side := range *sides {
// 		r.Push(side.Min())
// 	}
// 	return r
// }
// func Maxs(sides *hst.Sides) (r *hst.Stms) {
// 	r = hst.NewStms()
// 	for _, side := range *sides {
// 		r.Push(side.Max())
// 	}
// 	return r
// }
// func Lsts(sides *hst.Sides) (r *hst.Stms) {
// 	r = hst.NewStms()
// 	for _, side := range *sides {
// 		r.Push(side.Lst())
// 	}
// 	return r
// }
// func ProLsts(sides *hst.Sides) (r *hst.Stms) {
// 	r = hst.NewStms()
// 	for _, side := range *sides {
// 		r.Push(side.ProLst())
// 	}
// 	return r
// }
// func ProSmas(sides *hst.Sides) (r *hst.Stms) {
// 	r = hst.NewStms()
// 	for _, side := range *sides {
// 		r.Push(side.ProSma())
// 	}
// 	return r
// }
// func ProAlmas(sides *hst.Sides) (r *hst.Stms) {
// 	r = hst.NewStms()
// 	for _, side := range *sides {
// 		r.Push(side.ProAlma())
// 	}
// 	return r
// }

// var (
// 	SideStmss = []SideStms{
// 		Vrncs, Rsis, Wrsis,
// 		Emas, Wmas, Smas, Almas,
// 		RngFuls, RngLsts,
// 		Mins, Maxs, Lsts,
// 		ProLsts, ProSmas, ProAlmas,
// 	}
// )

// var (
// 	Durs = tmes.New(
// 		tme.M1, tme.M1*2, tme.M1*3,
// 		tme.M1*5, tme.M1*8, tme.M1*13,
// 		tme.M1*21, tme.M1*34, tme.M1*55,
// 		tme.M1*89, tme.M1*144, tme.M1*233,
// 		tme.M1*377,
// 	)
// )
