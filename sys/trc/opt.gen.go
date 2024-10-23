package trc

import (
	"strings"
)

var (
	optNames = map[Opt]string{
		Run:         "run",
		Prv:         "prv",
		TicrRx:      "ticrRx",
		TicrTx:      "ticrTx",
		RltInstr:    "rltInstr",
		RltInrvl:    "rltInrvl",
		RltSide:     "rltSide",
		RltStm:      "rltStm",
		RltCnd:      "rltCnd",
		RltStgy:     "rltStgy",
		RltPort:     "rltPort",
		RltPrfm:     "rltPrfm",
		HstInstr:    "hstInstr",
		HstInrvl:    "hstInrvl",
		HstSide:     "hstSide",
		HstStm:      "hstStm",
		HstCnd:      "hstCnd",
		HstStgy:     "hstStgy",
		HstPort:     "hstPort",
		HstPrfm:     "hstPrfm",
		HstInstrFbr: "hstInstrFbr",
		HstInrvlFbr: "hstInrvlFbr",
		HstSideFbr:  "hstSideFbr",
		HstStmFbr:   "hstStmFbr",
		HstCndFbr:   "hstCndFbr",
		HstStgyFbr:  "hstStgyFbr",
		HstPortFbr:  "hstPortFbr",
		HstPrfmFbr:  "hstPrfmFbr",
		Tune:        "tune",
		NoOpt:       "noOpt",
	}
)

const (
	Run Opt = 1 << iota
	Prv
	TicrRx
	TicrTx
	RltInstr
	RltInrvl
	RltSide
	RltStm
	RltCnd
	RltStgy
	RltPort
	RltPrfm
	HstInstr
	HstInrvl
	HstSide
	HstStm
	HstCnd
	HstStgy
	HstPort
	HstPrfm
	HstInstrFbr
	HstInrvlFbr
	HstSideFbr
	HstStmFbr
	HstCndFbr
	HstStgyFbr
	HstPortFbr
	HstPrfmFbr
	Tune
	NoOpt Opt = 0
)

type (
	Opt uint32
)

func PrsOpt(txt string) (r Opt) {
	txtVals := strings.Split(txt, "|")
	for _, s := range txtVals {
		switch strings.TrimSpace(s) {
		case "run":
			r |= Run
		case "prv":
			r |= Prv
		case "ticrRx":
			r |= TicrRx
		case "ticrTx":
			r |= TicrTx
		case "rltInstr":
			r |= RltInstr
		case "rltInrvl":
			r |= RltInrvl
		case "rltSide":
			r |= RltSide
		case "rltStm":
			r |= RltStm
		case "rltCnd":
			r |= RltCnd
		case "rltStgy":
			r |= RltStgy
		case "rltPort":
			r |= RltPort
		case "rltPrfm":
			r |= RltPrfm
		case "hstInstr":
			r |= HstInstr
		case "hstInrvl":
			r |= HstInrvl
		case "hstSide":
			r |= HstSide
		case "hstStm":
			r |= HstStm
		case "hstCnd":
			r |= HstCnd
		case "hstStgy":
			r |= HstStgy
		case "hstPort":
			r |= HstPort
		case "hstPrfm":
			r |= HstPrfm
		case "hstInstrFbr":
			r |= HstInstrFbr
		case "hstInrvlFbr":
			r |= HstInrvlFbr
		case "hstSideFbr":
			r |= HstSideFbr
		case "hstStmFbr":
			r |= HstStmFbr
		case "hstCndFbr":
			r |= HstCndFbr
		case "hstStgyFbr":
			r |= HstStgyFbr
		case "hstPortFbr":
			r |= HstPortFbr
		case "hstPrfmFbr":
			r |= HstPrfmFbr
		case "tune":
			r |= Tune
		case "noOpt":
			r |= NoOpt
		}
	}
	return r
}
func (x Opt) String() string { return optNames[x] }
func (x Opt) Is(vs ...Opt) (r bool) {
	for _, v := range vs {
		if x&v != v {
			return false
		}
	}
	return true
}
func (x Opt) IsRun() bool         { return x&Run == Run }
func (x Opt) IsPrv() bool         { return x&Prv == Prv }
func (x Opt) IsTicrRx() bool      { return x&TicrRx == TicrRx }
func (x Opt) IsTicrTx() bool      { return x&TicrTx == TicrTx }
func (x Opt) IsRltInstr() bool    { return x&RltInstr == RltInstr }
func (x Opt) IsRltInrvl() bool    { return x&RltInrvl == RltInrvl }
func (x Opt) IsRltSide() bool     { return x&RltSide == RltSide }
func (x Opt) IsRltStm() bool      { return x&RltStm == RltStm }
func (x Opt) IsRltCnd() bool      { return x&RltCnd == RltCnd }
func (x Opt) IsRltStgy() bool     { return x&RltStgy == RltStgy }
func (x Opt) IsRltPort() bool     { return x&RltPort == RltPort }
func (x Opt) IsRltPrfm() bool     { return x&RltPrfm == RltPrfm }
func (x Opt) IsHstInstr() bool    { return x&HstInstr == HstInstr }
func (x Opt) IsHstInrvl() bool    { return x&HstInrvl == HstInrvl }
func (x Opt) IsHstSide() bool     { return x&HstSide == HstSide }
func (x Opt) IsHstStm() bool      { return x&HstStm == HstStm }
func (x Opt) IsHstCnd() bool      { return x&HstCnd == HstCnd }
func (x Opt) IsHstStgy() bool     { return x&HstStgy == HstStgy }
func (x Opt) IsHstPort() bool     { return x&HstPort == HstPort }
func (x Opt) IsHstPrfm() bool     { return x&HstPrfm == HstPrfm }
func (x Opt) IsHstInstrFbr() bool { return x&HstInstrFbr == HstInstrFbr }
func (x Opt) IsHstInrvlFbr() bool { return x&HstInrvlFbr == HstInrvlFbr }
func (x Opt) IsHstSideFbr() bool  { return x&HstSideFbr == HstSideFbr }
func (x Opt) IsHstStmFbr() bool   { return x&HstStmFbr == HstStmFbr }
func (x Opt) IsHstCndFbr() bool   { return x&HstCndFbr == HstCndFbr }
func (x Opt) IsHstStgyFbr() bool  { return x&HstStgyFbr == HstStgyFbr }
func (x Opt) IsHstPortFbr() bool  { return x&HstPortFbr == HstPortFbr }
func (x Opt) IsHstPrfmFbr() bool  { return x&HstPrfmFbr == HstPrfmFbr }
func (x Opt) IsTune() bool        { return x&Tune == Tune }
func (x Opt) IsNoOpt() bool       { return x&NoOpt == NoOpt }
