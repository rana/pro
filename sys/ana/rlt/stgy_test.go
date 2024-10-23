package rlt_test

import (
	"sys"
	"sys/ana/hst"
	"sys/ana/rlt"
	"sys/app"
	"sys/lng/pro/act"
	"sys/tst"
	"testing"
)

func TestRltStgy(t *testing.T) {
	cfg := tst.NewCfg()
	cfg.Hst.StgySeqTrd = true          // StgySeqTrd NOT NEEDED WHEN TESTING FUTURE PERIOD PRFM
	cfg.DskPth = "/home/rana/data/tst" // ADD DSK TO TEST ML NET SAV

	ap := app.New(cfg, "/home/rana/go/src/sys/cmd")
	ap.Oan.CloneInstrs(tst.Instrs)
	prv := rlt.Oan()
	Instr := tst.RltPrvInstrs[len(tst.RltPrvInstrs)-1] // lst elm has most points
	instr := Instr(prv)
	inrvl := tst.RltInstrInrvlI(instr, 10)
	side := tst.RltInrvlSideBid(inrvl)
	stm := tst.RltSideStmRteLst(side)
	cnd := tst.RltStmCndInrGtr(stm, 1)
	s1Bid := instr.I(1).Bid()
	s2Bid := instr.I(2).Bid()
	s3Bid := instr.I(3).Bid()
	s5Bid := instr.I(5).Bid()
	s8Bid := instr.I(8).Bid()
	ftrStms := rlt.NewStms(
		s1Bid.Lst(), // 0.808
		s1Bid.Min(), // 0.824
		s1Bid.Max(), // 0.851

		s2Bid.Lst(), // 0.808
		s2Bid.Min(), // 0.824
		s2Bid.Max(), // 0.851

		s3Bid.Lst(), // 0.808
		s3Bid.Min(), // 0.824
		s3Bid.Max(), // 0.851

		s5Bid.Lst(), // 0.808
		s5Bid.Min(), // 0.824
		s5Bid.Max(), // 0.851

		s8Bid.Lst(), // 0.808
		s8Bid.Min(), // 0.824
		s8Bid.Max(), // 0.851
	)
	stgy := tst.RltCndStgyStgy(cnd, true, 2.0, 4.0, 60*60, 0.0, instr, ftrStms).(*rlt.StgyStgy)
	tst.RltStgyStgyNotZero(t, stgy)

	mnr := tst.NewStgyMnr(ap)
	stgy.Sub(mnr.Rx, mnr.Id)
	tst.IntegerEql(t, 1, len(stgy.Rxs), "Sub Rxs")
	var actr act.Actr
	vs := actr.RunHst(stgy.String())
	eHst := vs[len(vs)-1].(*hst.StgyStgy)
	eHst.Fit()
	eHst.CalcSeqTrds(stgy.String()) // rlt key
	// sys.Lrnr().LoadNetFromDsk(stgy.Key) // rlt key
	sys.Log("StartFor: eHst.Trds:", eHst.Trds.Cnt())
	if eHst.Trds != nil {
		mnr.StartFor(instr.Instr(), eHst.Trds.Cnt())
		tst.AnaTrdsEql(t, eHst.Trds, mnr.Trds, "Trds")
	}
	stgy.Unsub(mnr.Id)
	tst.IntegerEql(t, 0, len(stgy.Rxs), "Unsub Rxs")
	ap.Cls()
}
