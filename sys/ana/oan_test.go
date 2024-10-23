package ana_test

// import (
// 	"sys/ana/prv"
// 	"sys/app"
// 	"sys/tst"
// 	"testing"
// )

// func TestCorOanLoadHst(t *testing.T) {
// 	tst.EnsureNet(t)
// 	ap := app.New(tst.Cfg) // use cfg to define query range
// 	eurUsd := ap.Ana.Oan().Hst().EurUsd()
// 	tst.NotNil(t, eurUsd.I.HstStm)
// 	tst.UntNotZero(t, eurUsd.I.HstStm.Cnt())
// }

// func TestDskInstrDetail(t *testing.T) {
// 	tst.EnsureDsk(t)
// 	c := tst.NewCfg() // use cfg to defines query range
// 	c.DskPth = tst.TimeDir("/home/rana/test/")
// 	ap := app.New(c)
// 	oan := ap.Ana.Oan()
// 	for _, e := range tst.Instrs {
// 		e.CalcStats()
// 		oan.DskSavInstrDetail(e)
// 		a := &ana.Instr{}
// 		a.Name = e.Name
// 		loaded := oan.DskLoadInstrDetail(a)
// 		tst.True(t, loaded, "loaded")
// 		tst.PrvInstrEql(t, e, a)
// 	}
// }
// func TestDskInstrStm(t *testing.T) {
// 	tst.EnsureDsk(t)
// 	c := tst.NewCfg() // use cfg to defines query range
// 	c.DskPth = tst.TimeDir("/home/rana/test/")
// 	ap := app.New(c)
// 	oan := ap.Ana.Oan()
// 	for _, e := range tst.Instrs {
// 		oan.DskSavInstrStm(e)
// 		a := &ana.Instr{}
// 		a.Name = e.Name
// 		loaded := oan.DskLoadInstrStm(a)
// 		tst.True(t, loaded, "loaded")
// 		tst.PrvStmEql(t, e.HstStm, a.HstStm)
// 	}
// }
