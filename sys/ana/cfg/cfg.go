package cfg

import (
	"sys/bsc/bol"
	"sys/bsc/flt"
	"sys/bsc/str"
	"sys/bsc/tme"
	"sys/bsc/unt"
	"sys/fs"
	procfg "sys/lng/pro/cfg"
	"sys/trc"
)

const (
	Filename = "sys.cfg"
)

type (
	Cfg struct {
		Trc       trc.Opt
		Test      bol.Bol
		Ui        bol.Bol
		Wd        string
		DskPth    string
		BqProject str.Str
		Oan
		Hst Hst
		Rlt Rlt
		Plt Plt
		Ml  Ml
	}
	Oan struct {
		AccountID     str.Str // TODO: ACCOUNT TO DB? OR ARRAY?
		Token         str.Str
		BaseURL       str.Str
		BaseStreamURL str.Str
		MktHr         tme.Tme
		MktTrdBuf     tme.Tme
		ComPerUnitUsd flt.Flt
		BqDataset     str.Str
		GapFil        bol.Bol
		PriceBnd      bol.Bol
	}
	Hst struct {
		Rng        tme.Rng
		BalUsd     flt.Flt
		TrdPct     flt.Flt
		StgySeqTrd bol.Bol
		MktWeeks   []tme.Rng // generated: mkt rngs based on rng
	}
	Rlt struct {
		TrdPct    flt.Flt
		OpnScript str.Str
	}
	Plt struct {
		InrvlCnt unt.Unt
	}
	Ml struct {
		File string
	}
)

func Load(filename string) *Cfg { return Prs(fs.LoadText(filename)) }
func Prs(txt string) (r *Cfg) {
	var c procfg.Cfgr
	c.Reset(txt)
	const hst = "hst"
	const rlt = "rlt"
	const plt = "plt"
	const oan = "oan"
	r = &Cfg{}
	r.Trc = trc.PrsOpt(c.Str("trc").Unquo())
	r.Ui = c.Bol("ui")
	r.DskPth = c.Str("dskPth").Unquo()
	r.BqProject = c.Str("bqProject")
	// oan
	r.BqDataset = c.Str(oan, "bqDataset")
	r.Oan.BqDataset = c.Str(oan, "bqDataset")
	r.Oan.AccountID = c.Str(oan, "accountID")
	r.Oan.Token = c.Str(oan, "token")
	r.Oan.BaseURL = c.Str(oan, "baseURL")
	r.Oan.BaseStreamURL = c.Str(oan, "baseStreamURL")
	r.Oan.MktHr = c.Tme(oan, "mktHr")
	r.Oan.MktTrdBuf = c.Tme(oan, "mktTrdBuf")
	r.Oan.ComPerUnitUsd = c.Flt(oan, "comPerUnitUsd")
	r.Oan.GapFil = c.Bol(oan, "gapFil")
	r.Oan.PriceBnd = c.Bol(oan, "priceBnd")
	// hst
	r.Hst.Rng = c.TmeRng(hst, "rng")
	r.Hst.BalUsd = c.Flt(hst, "balUsd")
	r.Hst.TrdPct = c.Flt(hst, "trdPct")
	r.Hst.StgySeqTrd = c.Bol(hst, "stgySeqTrd")
	// rlt
	r.Rlt.TrdPct = c.Flt(rlt, "trdPct")
	r.Rlt.OpnScript = c.Str(rlt, "opnScript")
	// plt
	r.Plt.InrvlCnt = c.Unt(plt, "inrvlCnt")
	return r
}
