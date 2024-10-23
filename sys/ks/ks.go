package ks

import (
	"sys"
	"sys/k"
)

var (
	Idns       = sys.Vs(k.Eql, k.Neq)
	Rels       = sys.Vs(k.Lss, k.Gtr, k.Leq, k.Geq)
	Preds      = append(Idns, Rels...)
	Prvs       = sys.Vs(k.Oan)
	Doms       = sys.Vs(k.Hst, k.Rlt)
	Instrs     = sys.Vs(k.EurUsd, k.AudUsd, k.NzdUsd, k.GbpUsd)
	InstrNames = sys.Vs(k.EurUsdName, k.AudUsdName, k.NzdUsdName, k.GbpUsdName)
	Inrvls     = sys.Vs(k.S1, k.S5, k.S10, k.S15, k.S20, k.S30, k.S40, k.S50, k.M1, k.M5, k.M10, k.M15, k.M20, k.M30, k.M40, k.M50, k.H1)
	Sides      = sys.Vs(k.Bid, k.Ask)
	// Aggs: TODO: ADD CntRels To Agg, CntLss, CntLeq... and single Cnt
	Aggs = sys.Vs(k.Fst, k.Lst, k.Sum, k.Prd, k.Min, k.Max, k.Mid, k.Mdn, k.Sma, k.Gma, k.Wma, k.Rsi, k.Wrsi, k.Alma, k.Vrnc, k.Std, k.RngFul, k.RngLst, k.ProLst, k.ProSma, k.ProAlma)
	Aris = sys.Vs(k.Add, k.Sub, k.Mul, k.Div, k.Rem, k.Pow, k.Min, k.Max)

	Rtes  = Aggs
	Unas  = sys.Vs(k.Pos, k.Neg, k.Inv, k.Sqr, k.Sqrt)
	Scls  = Aris
	Sels  = Preds
	Inrs  = Aris
	Otrs  = Aris
	Stgys = sys.Vs(k.Long) //, k.Shrt)

	StmRtes = Rtes
	StmUnas = sys.CnjUnas(Unas...)
	StmScls = sys.CnjScls(Scls...)
	StmSels = sys.CnjSels(Sels...)
	StmAggs = sys.CnjAggs(Aggs...)
	StmInrs = sys.CnjInrs(Aris...)
	StmOtrs = sys.CnjOtrs(Aris...)
	CndScls = Preds //sys.CnjScls(Preds...)
	CndInrs = Preds //sys.CnjInrs(Preds...)
	CndOtrs = Preds //sys.CnjOtrs(Preds...)
	// Cnds    = sys.Vs(k.And)

	// no StmUnass (no prms to pluralize)
	StmSclss = sys.Plurals(StmScls...)
	StmSelss = sys.Plurals(StmSels...)
	StmAggss = sys.Plurals(StmAggs...)
	StmInrss = sys.Plurals(StmInrs...)
	StmOtrss = sys.Plurals(StmOtrs...)
	CndInrss = sys.Plurals(CndInrs...)
	CndOtrss = sys.Plurals(CndOtrs...)

	// TODO: SerPrt AddsLss, AddsLeg...

	Clrs = sys.Vs(
		k.Red, k.Pink, k.Purple,
		k.DeepPurple, k.Indigo, k.Blue,
		k.LightBlue, k.Cyan, k.Teal,
		k.Green, k.LightGreen, k.Lime,
		k.Yellow, k.Amber, k.Orange,
		k.DeepOrange,
		k.Brown, k.BlueGrey, k.Grey)
	ClrAs = sys.Vs(
		k.Red, k.Pink, k.Purple,
		k.DeepPurple, k.Indigo, k.Blue,
		k.LightBlue, k.Cyan, k.Teal,
		k.Green, k.LightGreen, k.Lime,
		k.Yellow, k.Amber, k.Orange,
		k.DeepOrange)

	ClrNums = sys.Vs(
		k.Clr50,
		k.Clr100, k.Clr200, k.Clr300,
		k.Clr400, k.Clr500, k.Clr600,
		k.Clr700, k.Clr800, k.Clr900,
	)
	ClrNumAs = sys.Vs(
		k.Clr100,
		k.Clr200,
		k.Clr400,
		k.Clr700,
	)
)
