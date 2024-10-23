package act_test

import (
	"sys/ana/hst"
	"sys/ana/rlt"
	"sys/app"
	"sys/bsc/flt"
	"sys/bsc/flts"
	"sys/bsc/tme"
	"sys/bsc/tmes"
	"sys/bsc/unt"
	"sys/bsc/unts"
	"sys/lng/pro/act"
	"sys/tst"
	"testing"
)

func TestActTmeNowPkgFn(t *testing.T) {
	var actr act.Actr
	expected := tme.Now()
	acts := actr.Cmplf("tme.now()")
	tst.IntegerEql(t, 1, len(acts), "acts count")
	act, ok := acts[0].(act.TmeNow)
	tst.True(t, ok, "cast")
	tst.TmeEql(t, expected, act.TmeTme(), "act")
}
func TestActFltsAddsLssPkgFn(t *testing.T) {
	var actr act.Actr
	expected := flts.AddsLss(flt.Flt(2), flt.Flt(10), flt.Flt(2))
	acts := actr.Cmplf("flts.addsLss(%v %v %v)", flt.Flt(2), flt.Flt(10), flt.Flt(2))
	tst.IntegerEql(t, 1, len(acts), "acts count")
	act, ok := acts[0].(act.FltsAddsLss)
	tst.True(t, ok, "cast")
	tst.FltsEql(t, expected, act.FltsFlts(), "act")
}
func TestActFltsAddsLeqPkgFn(t *testing.T) {
	var actr act.Actr
	expected := flts.AddsLeq(flt.Flt(2), flt.Flt(10), flt.Flt(2))
	acts := actr.Cmplf("flts.addsLeq(%v %v %v)", flt.Flt(2), flt.Flt(10), flt.Flt(2))
	tst.IntegerEql(t, 1, len(acts), "acts count")
	act, ok := acts[0].(act.FltsAddsLeq)
	tst.True(t, ok, "cast")
	tst.FltsEql(t, expected, act.FltsFlts(), "act")
}
func TestActFltsSubsGtrPkgFn(t *testing.T) {
	var actr act.Actr
	expected := flts.SubsGtr(flt.Flt(10), flt.Flt(2), flt.Flt(2))
	acts := actr.Cmplf("flts.subsGtr(%v %v %v)", flt.Flt(10), flt.Flt(2), flt.Flt(2))
	tst.IntegerEql(t, 1, len(acts), "acts count")
	act, ok := acts[0].(act.FltsSubsGtr)
	tst.True(t, ok, "cast")
	tst.FltsEql(t, expected, act.FltsFlts(), "act")
}
func TestActFltsSubsGeqPkgFn(t *testing.T) {
	var actr act.Actr
	expected := flts.SubsGeq(flt.Flt(10), flt.Flt(2), flt.Flt(2))
	acts := actr.Cmplf("flts.subsGeq(%v %v %v)", flt.Flt(10), flt.Flt(2), flt.Flt(2))
	tst.IntegerEql(t, 1, len(acts), "acts count")
	act, ok := acts[0].(act.FltsSubsGeq)
	tst.True(t, ok, "cast")
	tst.FltsEql(t, expected, act.FltsFlts(), "act")
}
func TestActFltsMulsLssPkgFn(t *testing.T) {
	var actr act.Actr
	expected := flts.MulsLss(flt.Flt(2), flt.Flt(10), flt.Flt(2))
	acts := actr.Cmplf("flts.mulsLss(%v %v %v)", flt.Flt(2), flt.Flt(10), flt.Flt(2))
	tst.IntegerEql(t, 1, len(acts), "acts count")
	act, ok := acts[0].(act.FltsMulsLss)
	tst.True(t, ok, "cast")
	tst.FltsEql(t, expected, act.FltsFlts(), "act")
}
func TestActFltsMulsLeqPkgFn(t *testing.T) {
	var actr act.Actr
	expected := flts.MulsLeq(flt.Flt(2), flt.Flt(10), flt.Flt(2))
	acts := actr.Cmplf("flts.mulsLeq(%v %v %v)", flt.Flt(2), flt.Flt(10), flt.Flt(2))
	tst.IntegerEql(t, 1, len(acts), "acts count")
	act, ok := acts[0].(act.FltsMulsLeq)
	tst.True(t, ok, "cast")
	tst.FltsEql(t, expected, act.FltsFlts(), "act")
}
func TestActFltsDivsGtrPkgFn(t *testing.T) {
	var actr act.Actr
	expected := flts.DivsGtr(flt.Flt(10), flt.Flt(2), flt.Flt(2))
	acts := actr.Cmplf("flts.divsGtr(%v %v %v)", flt.Flt(10), flt.Flt(2), flt.Flt(2))
	tst.IntegerEql(t, 1, len(acts), "acts count")
	act, ok := acts[0].(act.FltsDivsGtr)
	tst.True(t, ok, "cast")
	tst.FltsEql(t, expected, act.FltsFlts(), "act")
}
func TestActFltsDivsGeqPkgFn(t *testing.T) {
	var actr act.Actr
	expected := flts.DivsGeq(flt.Flt(10), flt.Flt(2), flt.Flt(2))
	acts := actr.Cmplf("flts.divsGeq(%v %v %v)", flt.Flt(10), flt.Flt(2), flt.Flt(2))
	tst.IntegerEql(t, 1, len(acts), "acts count")
	act, ok := acts[0].(act.FltsDivsGeq)
	tst.True(t, ok, "cast")
	tst.FltsEql(t, expected, act.FltsFlts(), "act")
}
func TestActFltsFibsLeqPkgFn(t *testing.T) {
	var actr act.Actr
	expected := flts.FibsLeq(flt.Flt(610))
	acts := actr.Cmplf("flts.fibsLeq(%v)", flt.Flt(610))
	tst.IntegerEql(t, 1, len(acts), "acts count")
	act, ok := acts[0].(act.FltsFibsLeq)
	tst.True(t, ok, "cast")
	tst.FltsEql(t, expected, act.FltsFlts(), "act")
}
func TestActUntsAddsLssPkgFn(t *testing.T) {
	var actr act.Actr
	expected := unts.AddsLss(unt.Unt(2), unt.Unt(10), unt.Unt(2))
	acts := actr.Cmplf("unts.addsLss(%v %v %v)", unt.Unt(2), unt.Unt(10), unt.Unt(2))
	tst.IntegerEql(t, 1, len(acts), "acts count")
	act, ok := acts[0].(act.UntsAddsLss)
	tst.True(t, ok, "cast")
	tst.UntsEql(t, expected, act.UntsUnts(), "act")
}
func TestActUntsAddsLeqPkgFn(t *testing.T) {
	var actr act.Actr
	expected := unts.AddsLeq(unt.Unt(2), unt.Unt(10), unt.Unt(2))
	acts := actr.Cmplf("unts.addsLeq(%v %v %v)", unt.Unt(2), unt.Unt(10), unt.Unt(2))
	tst.IntegerEql(t, 1, len(acts), "acts count")
	act, ok := acts[0].(act.UntsAddsLeq)
	tst.True(t, ok, "cast")
	tst.UntsEql(t, expected, act.UntsUnts(), "act")
}
func TestActUntsSubsGtrPkgFn(t *testing.T) {
	var actr act.Actr
	expected := unts.SubsGtr(unt.Unt(10), unt.Unt(2), unt.Unt(2))
	acts := actr.Cmplf("unts.subsGtr(%v %v %v)", unt.Unt(10), unt.Unt(2), unt.Unt(2))
	tst.IntegerEql(t, 1, len(acts), "acts count")
	act, ok := acts[0].(act.UntsSubsGtr)
	tst.True(t, ok, "cast")
	tst.UntsEql(t, expected, act.UntsUnts(), "act")
}
func TestActUntsSubsGeqPkgFn(t *testing.T) {
	var actr act.Actr
	expected := unts.SubsGeq(unt.Unt(10), unt.Unt(2), unt.Unt(2))
	acts := actr.Cmplf("unts.subsGeq(%v %v %v)", unt.Unt(10), unt.Unt(2), unt.Unt(2))
	tst.IntegerEql(t, 1, len(acts), "acts count")
	act, ok := acts[0].(act.UntsSubsGeq)
	tst.True(t, ok, "cast")
	tst.UntsEql(t, expected, act.UntsUnts(), "act")
}
func TestActUntsMulsLssPkgFn(t *testing.T) {
	var actr act.Actr
	expected := unts.MulsLss(unt.Unt(2), unt.Unt(10), unt.Unt(2))
	acts := actr.Cmplf("unts.mulsLss(%v %v %v)", unt.Unt(2), unt.Unt(10), unt.Unt(2))
	tst.IntegerEql(t, 1, len(acts), "acts count")
	act, ok := acts[0].(act.UntsMulsLss)
	tst.True(t, ok, "cast")
	tst.UntsEql(t, expected, act.UntsUnts(), "act")
}
func TestActUntsMulsLeqPkgFn(t *testing.T) {
	var actr act.Actr
	expected := unts.MulsLeq(unt.Unt(2), unt.Unt(10), unt.Unt(2))
	acts := actr.Cmplf("unts.mulsLeq(%v %v %v)", unt.Unt(2), unt.Unt(10), unt.Unt(2))
	tst.IntegerEql(t, 1, len(acts), "acts count")
	act, ok := acts[0].(act.UntsMulsLeq)
	tst.True(t, ok, "cast")
	tst.UntsEql(t, expected, act.UntsUnts(), "act")
}
func TestActUntsDivsGtrPkgFn(t *testing.T) {
	var actr act.Actr
	expected := unts.DivsGtr(unt.Unt(10), unt.Unt(2), unt.Unt(2))
	acts := actr.Cmplf("unts.divsGtr(%v %v %v)", unt.Unt(10), unt.Unt(2), unt.Unt(2))
	tst.IntegerEql(t, 1, len(acts), "acts count")
	act, ok := acts[0].(act.UntsDivsGtr)
	tst.True(t, ok, "cast")
	tst.UntsEql(t, expected, act.UntsUnts(), "act")
}
func TestActUntsDivsGeqPkgFn(t *testing.T) {
	var actr act.Actr
	expected := unts.DivsGeq(unt.Unt(10), unt.Unt(2), unt.Unt(2))
	acts := actr.Cmplf("unts.divsGeq(%v %v %v)", unt.Unt(10), unt.Unt(2), unt.Unt(2))
	tst.IntegerEql(t, 1, len(acts), "acts count")
	act, ok := acts[0].(act.UntsDivsGeq)
	tst.True(t, ok, "cast")
	tst.UntsEql(t, expected, act.UntsUnts(), "act")
}
func TestActUntsFibsLeqPkgFn(t *testing.T) {
	var actr act.Actr
	expected := unts.FibsLeq(unt.Unt(610))
	acts := actr.Cmplf("unts.fibsLeq(%v)", unt.Unt(610))
	tst.IntegerEql(t, 1, len(acts), "acts count")
	act, ok := acts[0].(act.UntsFibsLeq)
	tst.True(t, ok, "cast")
	tst.UntsEql(t, expected, act.UntsUnts(), "act")
}
func TestActTmesAddsLssPkgFn(t *testing.T) {
	var actr act.Actr
	expected := tmes.AddsLss(tme.Tme(2), tme.Tme(10), tme.Tme(2))
	acts := actr.Cmplf("tmes.addsLss(%v %v %v)", tme.Tme(2), tme.Tme(10), tme.Tme(2))
	tst.IntegerEql(t, 1, len(acts), "acts count")
	act, ok := acts[0].(act.TmesAddsLss)
	tst.True(t, ok, "cast")
	tst.TmesEql(t, expected, act.TmesTmes(), "act")
}
func TestActTmesAddsLeqPkgFn(t *testing.T) {
	var actr act.Actr
	expected := tmes.AddsLeq(tme.Tme(2), tme.Tme(10), tme.Tme(2))
	acts := actr.Cmplf("tmes.addsLeq(%v %v %v)", tme.Tme(2), tme.Tme(10), tme.Tme(2))
	tst.IntegerEql(t, 1, len(acts), "acts count")
	act, ok := acts[0].(act.TmesAddsLeq)
	tst.True(t, ok, "cast")
	tst.TmesEql(t, expected, act.TmesTmes(), "act")
}
func TestActTmesSubsGtrPkgFn(t *testing.T) {
	var actr act.Actr
	expected := tmes.SubsGtr(tme.Tme(10), tme.Tme(2), tme.Tme(2))
	acts := actr.Cmplf("tmes.subsGtr(%v %v %v)", tme.Tme(10), tme.Tme(2), tme.Tme(2))
	tst.IntegerEql(t, 1, len(acts), "acts count")
	act, ok := acts[0].(act.TmesSubsGtr)
	tst.True(t, ok, "cast")
	tst.TmesEql(t, expected, act.TmesTmes(), "act")
}
func TestActTmesSubsGeqPkgFn(t *testing.T) {
	var actr act.Actr
	expected := tmes.SubsGeq(tme.Tme(10), tme.Tme(2), tme.Tme(2))
	acts := actr.Cmplf("tmes.subsGeq(%v %v %v)", tme.Tme(10), tme.Tme(2), tme.Tme(2))
	tst.IntegerEql(t, 1, len(acts), "acts count")
	act, ok := acts[0].(act.TmesSubsGeq)
	tst.True(t, ok, "cast")
	tst.TmesEql(t, expected, act.TmesTmes(), "act")
}
func TestActTmesMulsLssPkgFn(t *testing.T) {
	var actr act.Actr
	expected := tmes.MulsLss(tme.Tme(2), tme.Tme(10), tme.Tme(2))
	acts := actr.Cmplf("tmes.mulsLss(%v %v %v)", tme.Tme(2), tme.Tme(10), tme.Tme(2))
	tst.IntegerEql(t, 1, len(acts), "acts count")
	act, ok := acts[0].(act.TmesMulsLss)
	tst.True(t, ok, "cast")
	tst.TmesEql(t, expected, act.TmesTmes(), "act")
}
func TestActTmesMulsLeqPkgFn(t *testing.T) {
	var actr act.Actr
	expected := tmes.MulsLeq(tme.Tme(2), tme.Tme(10), tme.Tme(2))
	acts := actr.Cmplf("tmes.mulsLeq(%v %v %v)", tme.Tme(2), tme.Tme(10), tme.Tme(2))
	tst.IntegerEql(t, 1, len(acts), "acts count")
	act, ok := acts[0].(act.TmesMulsLeq)
	tst.True(t, ok, "cast")
	tst.TmesEql(t, expected, act.TmesTmes(), "act")
}
func TestActTmesDivsGtrPkgFn(t *testing.T) {
	var actr act.Actr
	expected := tmes.DivsGtr(tme.Tme(10), tme.Tme(2), tme.Tme(2))
	acts := actr.Cmplf("tmes.divsGtr(%v %v %v)", tme.Tme(10), tme.Tme(2), tme.Tme(2))
	tst.IntegerEql(t, 1, len(acts), "acts count")
	act, ok := acts[0].(act.TmesDivsGtr)
	tst.True(t, ok, "cast")
	tst.TmesEql(t, expected, act.TmesTmes(), "act")
}
func TestActTmesDivsGeqPkgFn(t *testing.T) {
	var actr act.Actr
	expected := tmes.DivsGeq(tme.Tme(10), tme.Tme(2), tme.Tme(2))
	acts := actr.Cmplf("tmes.divsGeq(%v %v %v)", tme.Tme(10), tme.Tme(2), tme.Tme(2))
	tst.IntegerEql(t, 1, len(acts), "acts count")
	act, ok := acts[0].(act.TmesDivsGeq)
	tst.True(t, ok, "cast")
	tst.TmesEql(t, expected, act.TmesTmes(), "act")
}
func TestActTmesFibsLeqPkgFn(t *testing.T) {
	var actr act.Actr
	expected := tmes.FibsLeq(tme.Tme(610))
	acts := actr.Cmplf("tmes.fibsLeq(%v)", tme.Tme(610))
	tst.IntegerEql(t, 1, len(acts), "acts count")
	act, ok := acts[0].(act.TmesFibsLeq)
	tst.True(t, ok, "cast")
	tst.TmesEql(t, expected, act.TmesTmes(), "act")
}
func TestActHstOanPkgFn(t *testing.T) {
	ap := app.New(tst.Cfg)
	ap.Oan.CloneInstrs(tst.Instrs)
	defer ap.Cls()
	expected := hst.Oan()
	acts := ap.Actr.Cmpl("hst.oan()")
	tst.IntegerEql(t, 1, len(acts), "acts count")
	act, ok := acts[0].(act.HstOan)
	tst.True(t, ok, "cast")
	tst.HstPrvEql(t, expected, act.HstPrv(), "act")
}
func TestActRltOanPkgFn(t *testing.T) {
	ap := app.New(tst.Cfg)
	ap.Oan.CloneInstrs(tst.Instrs)
	defer ap.Cls()
	expected := rlt.Oan()
	acts := ap.Actr.Cmpl("rlt.oan()")
	tst.IntegerEql(t, 1, len(acts), "acts count")
	act, ok := acts[0].(act.RltOan)
	tst.True(t, ok, "cast")
	tst.RltPrvEql(t, expected, act.RltPrv(), "act")
}
