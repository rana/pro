package cfg_test

import (
	"sys/ana/cfg"
	"sys/tst"
	"testing"
)

func TestPrsCfg(t *testing.T) {
	txt := `{
		trc:"ticrTx|rltInstr"
		ui: tru
		dskPth: "/home/rana/data/tst"
		bqProject: "pro-cld"
		bqDataset: "pro"

		oan:{
			bqDataset: "oan_cor_tic"
			accountID:"101-001-4093945-001"
			token:"d3e81b3ca79ebd8cf471e31b7d3a4f74-d5cee5a343849dcbb45e3fd6f21870f0"
			baseURL:"https://api-fxpractice.oanda.com"
			baseStreamURL:"https://stream-fxpractice.oanda.com"
			mktHr: 12h
			mktTrdBuf: 30m
			comPerUnitUsd: 0.00005
			gapFil: tru
			priceBnd: tru
		}
		hst:{
			rng: 2018y6n10d-2018y6n17d
			balUsd: 10000.0
			trdPct: 0.5
			stgySeqTrd: tru
		}
		rlt:{
			trdPct: 0.1
			opnScript: "ana.oan().rlt().opn()"
		}
		plt:{
			inrvlCnt: 8
		}
	}`

	c := cfg.Prs(txt)
	tst.StrNotEmpty(t, c.BqProject, "BqProject")
	tst.True(t, c.Trc.IsTicrTx(), "Rlt.Trc:ticrTx")
	tst.True(t, c.Trc.IsRltInstr(), "Rlt.Trc:rltInstr")
	// oan
	tst.StrNotEmpty(t, c.Oan.BqDataset, "BqDataset")
	tst.StrNotEmpty(t, c.Oan.AccountID, "AccountID")
	tst.StrNotEmpty(t, c.Oan.Token, "Token")
	tst.StrNotEmpty(t, c.Oan.BaseURL, "BaseURL")
	tst.StrNotEmpty(t, c.Oan.BaseStreamURL, "BaseStreamURL")
	tst.TmeNotZero(t, c.Oan.MktHr, "MktHr")
	tst.TmeNotZero(t, c.Oan.MktTrdBuf, "MktClsBuf")
	tst.FltNotZero(t, c.Oan.ComPerUnitUsd, "CommissionPerUnit")
	tst.BolTru(t, c.Oan.GapFil, "GapFil")
	tst.BolTru(t, c.Oan.PriceBnd, "PriceBnd")
	// hst
	tst.TmeRngNotZero(t, c.Hst.Rng, "Hst.Rng")
	tst.FltNotZero(t, c.Hst.BalUsd, "Hst.BalUsd")
	tst.FltNotZero(t, c.Hst.TrdPct, "Hst.TrdPct")
	tst.BolTru(t, c.Hst.StgySeqTrd, "Hst.StgySeqTrd")
	// rlt
	tst.FltNotZero(t, c.Rlt.TrdPct, "Rlt.TrdPct")
	tst.StrNotEmpty(t, c.Rlt.OpnScript, "Rlt.OpnScript")
	// plt
	tst.UntNotZero(t, c.Plt.InrvlCnt, "Plt.InrvlCnt")
}
