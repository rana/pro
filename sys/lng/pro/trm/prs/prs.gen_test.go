package prs_test

import (
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
	"sys/lng/pro/trm"
	"sys/lng/pro/trm/prs"
	"sys/tst"
	"testing"
)

func TestPrsStrTrm(t *testing.T) {
	cses := []struct {
		lit string
		val str.Str
	}{
		{"\"\"", ""},
		{"\"xYz\"", "xYz"},
		{"\"a\"", "a"},
		{"\"efg HIJ jKl\"", "efg HIJ jKl"},
	}
	for _, cse := range cses {
		t.Run("StrTrm", func(t *testing.T) {
			var trmr trm.Trmr
			trmr.Reset(cse.lit)
			strLit, ok := trmr.StrLit()
			tst.True(t, ok)
			a := prs.StrTrm(strLit, cse.lit)
			tst.StrEql(t, cse.val, a)
		})
	}
}
func TestPrsStr(t *testing.T) {
	cses := []struct {
		lit string
		val str.Str
	}{
		{"\"\"", ""},
		{"\"xYz\"", "xYz"},
		{"\"a\"", "a"},
		{"\"efg HIJ jKl\"", "efg HIJ jKl"},
	}
	for _, cse := range cses {
		t.Run("Str", func(t *testing.T) {
			a := prs.Str(cse.lit)
			tst.StrEql(t, cse.val, a)
		})
	}
}
func TestPrsBolTrm(t *testing.T) {
	cses := []struct {
		lit string
		val bol.Bol
	}{
		{"fls", false},
		{"tru", true},
	}
	for _, cse := range cses {
		t.Run("BolTrm", func(t *testing.T) {
			var trmr trm.Trmr
			trmr.Reset(cse.lit)
			bolLit, ok := trmr.BolLit()
			tst.True(t, ok)
			a := prs.BolTrm(bolLit, cse.lit)
			tst.BolEql(t, cse.val, a)
		})
	}
}
func TestPrsBol(t *testing.T) {
	cses := []struct {
		lit string
		val bol.Bol
	}{
		{"fls", false},
		{"tru", true},
	}
	for _, cse := range cses {
		t.Run("Bol", func(t *testing.T) {
			a := prs.Bol(cse.lit)
			tst.BolEql(t, cse.val, a)
		})
	}
}
func TestPrsFltTrm(t *testing.T) {
	cses := []struct {
		lit string
		val flt.Flt
	}{
		{"0.0", 0.0},
		{"1.1", 1.1},
		{"3.0", 3.0},
		{"3.0", 3.0},
		{"3.0", 3.0},
		{"99999.99", 99999.99},
		{"-1.1", -1.1},
		{"-99999.99", -99999.99},
	}
	for _, cse := range cses {
		t.Run("FltTrm", func(t *testing.T) {
			var trmr trm.Trmr
			trmr.Reset(cse.lit)
			fltLit, ok := trmr.FltLit()
			tst.True(t, ok)
			a := prs.FltTrm(fltLit, cse.lit)
			tst.FltEql(t, cse.val, a)
		})
	}
}
func TestPrsFlt(t *testing.T) {
	cses := []struct {
		lit string
		val flt.Flt
	}{
		{"0.0", 0.0},
		{"1.1", 1.1},
		{"3.0", 3.0},
		{"3.0", 3.0},
		{"3.0", 3.0},
		{"99999.99", 99999.99},
		{"-1.1", -1.1},
		{"-99999.99", -99999.99},
	}
	for _, cse := range cses {
		t.Run("Flt", func(t *testing.T) {
			a := prs.Flt(cse.lit)
			tst.FltEql(t, cse.val, a)
		})
	}
}
func TestPrsUntTrm(t *testing.T) {
	cses := []struct {
		lit string
		val unt.Unt
	}{
		{"0", 0},
		{"1", 1},
		{"1000", 1000},
		{"10", 10},
	}
	for _, cse := range cses {
		t.Run("UntTrm", func(t *testing.T) {
			var trmr trm.Trmr
			trmr.Reset(cse.lit)
			untLit, ok := trmr.UntLit()
			tst.True(t, ok)
			a := prs.UntTrm(untLit, cse.lit)
			tst.UntEql(t, cse.val, a)
		})
	}
}
func TestPrsUnt(t *testing.T) {
	cses := []struct {
		lit string
		val unt.Unt
	}{
		{"0", 0},
		{"1", 1},
		{"1000", 1000},
		{"10", 10},
	}
	for _, cse := range cses {
		t.Run("Unt", func(t *testing.T) {
			a := prs.Unt(cse.lit)
			tst.UntEql(t, cse.val, a)
		})
	}
}
func TestPrsIntTrm(t *testing.T) {
	cses := []struct {
		lit string
		val int.Int
	}{
		{"+0", 0},
		{"+10", 10},
		{"+1000", 1000},
		{"-10", -10},
		{"-1000", -1000},
	}
	for _, cse := range cses {
		t.Run("IntTrm", func(t *testing.T) {
			var trmr trm.Trmr
			trmr.Reset(cse.lit)
			intLit, ok := trmr.IntLit()
			tst.True(t, ok)
			a := prs.IntTrm(intLit, cse.lit)
			tst.IntEql(t, cse.val, a)
		})
	}
}
func TestPrsInt(t *testing.T) {
	cses := []struct {
		lit string
		val int.Int
	}{
		{"+0", 0},
		{"+10", 10},
		{"+1000", 1000},
		{"-10", -10},
		{"-1000", -1000},
	}
	for _, cse := range cses {
		t.Run("Int", func(t *testing.T) {
			a := prs.Int(cse.lit)
			tst.IntEql(t, cse.val, a)
		})
	}
}
func TestPrsTmeTrm(t *testing.T) {
	cses := []struct {
		lit string
		val tme.Tme
	}{
		{"0s", 0},
		{"1s", 1},
		{"10s", 10},
		{"-1s", -1},
		{"-10m", -10 * 60},
		{"1w2d3h4m5s", 788645},
		{"-1w2d3h4m5s", -788645},
		{"2000y1n2d3h4m5s", 946782245},
		{"2000y2d3h4m5s", 946782245},
		{"2000y3h4m5s", 946695845},
	}
	for _, cse := range cses {
		t.Run("TmeTrm", func(t *testing.T) {
			var trmr trm.Trmr
			trmr.Reset(cse.lit)
			tmeLit, ok := trmr.TmeLit()
			tst.True(t, ok)
			a := prs.TmeTrm(tmeLit, cse.lit)
			tst.TmeEql(t, cse.val, a)
		})
	}
}
func TestPrsTme(t *testing.T) {
	cses := []struct {
		lit string
		val tme.Tme
	}{
		{"0s", 0},
		{"1s", 1},
		{"10s", 10},
		{"-1s", -1},
		{"-10m", -10 * 60},
		{"1w2d3h4m5s", 788645},
		{"-1w2d3h4m5s", -788645},
		{"2000y1n2d3h4m5s", 946782245},
		{"2000y2d3h4m5s", 946782245},
		{"2000y3h4m5s", 946695845},
	}
	for _, cse := range cses {
		t.Run("Tme", func(t *testing.T) {
			a := prs.Tme(cse.lit)
			tst.TmeEql(t, cse.val, a)
		})
	}
}
func TestPrsBndTrm(t *testing.T) {
	cses := []struct {
		lit string
		val bnd.Bnd
	}{
		{"0-0", bnd.Bnd{Idx: 0, Lim: 0}},
		{"0-1", bnd.Bnd{Idx: 0, Lim: 1}},
		{"0-1000", bnd.Bnd{Idx: 0, Lim: 1000}},
		{"999-1000", bnd.Bnd{Idx: 999, Lim: 1000}},
		{"1-0", bnd.Bnd{Idx: 1, Lim: 0}},
	}
	for _, cse := range cses {
		t.Run("BndTrm", func(t *testing.T) {
			var trmr trm.Trmr
			trmr.Reset(cse.lit)
			bndLit, ok := trmr.BndLit()
			tst.True(t, ok)
			a := prs.BndTrm(bndLit, cse.lit)
			tst.BndEql(t, cse.val, a)
		})
	}
}
func TestPrsBnd(t *testing.T) {
	cses := []struct {
		lit string
		val bnd.Bnd
	}{
		{"0-0", bnd.Bnd{Idx: 0, Lim: 0}},
		{"0-1", bnd.Bnd{Idx: 0, Lim: 1}},
		{"0-1000", bnd.Bnd{Idx: 0, Lim: 1000}},
		{"999-1000", bnd.Bnd{Idx: 999, Lim: 1000}},
		{"1-0", bnd.Bnd{Idx: 1, Lim: 0}},
	}
	for _, cse := range cses {
		t.Run("Bnd", func(t *testing.T) {
			a := prs.Bnd(cse.lit)
			tst.BndEql(t, cse.val, a)
		})
	}
}
func TestPrsFltRngTrm(t *testing.T) {
	cses := []struct {
		lit string
		val flt.Rng
	}{
		{"0.0-0.0", flt.Rng{Min: 0, Max: 0}},
		{"0.0-1.0", flt.Rng{Min: 0, Max: 1}},
		{"-3.0--4.0", flt.Rng{Min: -3, Max: -4}},
		{"-999.0-1000.0", flt.Rng{Min: -999, Max: 1000}},
		{"1.0-0.0", flt.Rng{Min: 1, Max: 0}},
	}
	for _, cse := range cses {
		t.Run("FltRngTrm", func(t *testing.T) {
			var trmr trm.Trmr
			trmr.Reset(cse.lit)
			fltRngLit, ok := trmr.FltRngLit()
			tst.True(t, ok)
			a := prs.FltRngTrm(fltRngLit, cse.lit)
			tst.FltRngEql(t, cse.val, a)
		})
	}
}
func TestPrsFltRng(t *testing.T) {
	cses := []struct {
		lit string
		val flt.Rng
	}{
		{"0.0-0.0", flt.Rng{Min: 0, Max: 0}},
		{"0.0-1.0", flt.Rng{Min: 0, Max: 1}},
		{"-3.0--4.0", flt.Rng{Min: -3, Max: -4}},
		{"-999.0-1000.0", flt.Rng{Min: -999, Max: 1000}},
		{"1.0-0.0", flt.Rng{Min: 1, Max: 0}},
	}
	for _, cse := range cses {
		t.Run("FltRng", func(t *testing.T) {
			a := prs.FltRng(cse.lit)
			tst.FltRngEql(t, cse.val, a)
		})
	}
}
func TestPrsTmeRngTrm(t *testing.T) {
	cses := []struct {
		lit string
		val tme.Rng
	}{
		{"-10s--1s", tme.Rng{Min: -10, Max: -1}},
		{"0s-1s", tme.Rng{Min: 0, Max: 1}},
		{"2s-4s", tme.Rng{Min: 2, Max: 4}},
		{"6s-10s", tme.Rng{Min: 6, Max: 10}},
		{"50s-60s", tme.Rng{Min: 50, Max: 60}},
	}
	for _, cse := range cses {
		t.Run("TmeRngTrm", func(t *testing.T) {
			var trmr trm.Trmr
			trmr.Reset(cse.lit)
			tmeRngLit, ok := trmr.TmeRngLit()
			tst.True(t, ok)
			a := prs.TmeRngTrm(tmeRngLit, cse.lit)
			tst.TmeRngEql(t, cse.val, a)
		})
	}
}
func TestPrsTmeRng(t *testing.T) {
	cses := []struct {
		lit string
		val tme.Rng
	}{
		{"-10s--1s", tme.Rng{Min: -10, Max: -1}},
		{"0s-1s", tme.Rng{Min: 0, Max: 1}},
		{"2s-4s", tme.Rng{Min: 2, Max: 4}},
		{"6s-10s", tme.Rng{Min: 6, Max: 10}},
		{"50s-60s", tme.Rng{Min: 50, Max: 60}},
	}
	for _, cse := range cses {
		t.Run("TmeRng", func(t *testing.T) {
			a := prs.TmeRng(cse.lit)
			tst.TmeRngEql(t, cse.val, a)
		})
	}
}
func TestPrsStrsTrm(t *testing.T) {
	cses := []struct {
		lit string
		val *strs.Strs
	}{
		{"[\"\" \"xYz\" \"a\"]", strs.New("", "xYz", "a")},
		{"[\"\" \"xYz\" \"a\" \"efg HIJ jKl\"]", strs.New("", "xYz", "a", "efg HIJ jKl")},
	}
	for _, cse := range cses {
		t.Run("StrsTrm", func(t *testing.T) {
			var trmr trm.Trmr
			trmr.Reset(cse.lit)
			strsLit, ok := trmr.StrsLit()
			tst.True(t, ok)
			a := prs.StrsTrm(strsLit, cse.lit)
			tst.StrsEql(t, cse.val, a)
		})
	}
}
func TestPrsStrs(t *testing.T) {
	cses := []struct {
		lit string
		val *strs.Strs
	}{
		{"[\"\" \"xYz\" \"a\"]", strs.New("", "xYz", "a")},
		{"[\"\" \"xYz\" \"a\" \"efg HIJ jKl\"]", strs.New("", "xYz", "a", "efg HIJ jKl")},
	}
	for _, cse := range cses {
		t.Run("Strs", func(t *testing.T) {
			a := prs.Strs(cse.lit)
			tst.StrsEql(t, cse.val, a)
		})
	}
}
func TestPrsBolsTrm(t *testing.T) {
	cses := []struct {
		lit string
		val *bols.Bols
	}{
		{"[fls tru]", bols.New(false, true)},
	}
	for _, cse := range cses {
		t.Run("BolsTrm", func(t *testing.T) {
			var trmr trm.Trmr
			trmr.Reset(cse.lit)
			bolsLit, ok := trmr.BolsLit()
			tst.True(t, ok)
			a := prs.BolsTrm(bolsLit, cse.lit)
			tst.BolsEql(t, cse.val, a)
		})
	}
}
func TestPrsBols(t *testing.T) {
	cses := []struct {
		lit string
		val *bols.Bols
	}{
		{"[fls tru]", bols.New(false, true)},
	}
	for _, cse := range cses {
		t.Run("Bols", func(t *testing.T) {
			a := prs.Bols(cse.lit)
			tst.BolsEql(t, cse.val, a)
		})
	}
}
func TestPrsFltsTrm(t *testing.T) {
	cses := []struct {
		lit string
		val *flts.Flts
	}{
		{"[0.0 1.1 3.0]", flts.New(0.0, 1.1, 3.0)},
		{"[0.0 1.1 3.0 3.0 3.0 99999.99 -1.1 -99999.99]", flts.New(0.0, 1.1, 3.0, 3.0, 3.0, 99999.99, -1.1, -99999.99)},
		{"[0.0 1.1 3.0 3.0 3.0 99999.99 -1.1 -99999.99 0.0 1.1 3.0 3.0 3.0 99999.99 -1.1 -99999.99 0.0 1.1 3.0 3.0 3.0 99999.99 -1.1 -99999.99 0.0 1.1 3.0 3.0 3.0 99999.99 -1.1 -99999.99 0.0 1.1 3.0 3.0 3.0 99999.99 -1.1 -99999.99 0.0 1.1 3.0 3.0 3.0 99999.99 -1.1 -99999.99 0.0 1.1 3.0 3.0 3.0 99999.99 -1.1 -99999.99 0.0 1.1 3.0 3.0 3.0 99999.99 -1.1 -99999.99]", flts.New(0.0, 1.1, 3.0, 3.0, 3.0, 99999.99, -1.1, -99999.99, 0.0, 1.1, 3.0, 3.0, 3.0, 99999.99, -1.1, -99999.99, 0.0, 1.1, 3.0, 3.0, 3.0, 99999.99, -1.1, -99999.99, 0.0, 1.1, 3.0, 3.0, 3.0, 99999.99, -1.1, -99999.99, 0.0, 1.1, 3.0, 3.0, 3.0, 99999.99, -1.1, -99999.99, 0.0, 1.1, 3.0, 3.0, 3.0, 99999.99, -1.1, -99999.99, 0.0, 1.1, 3.0, 3.0, 3.0, 99999.99, -1.1, -99999.99, 0.0, 1.1, 3.0, 3.0, 3.0, 99999.99, -1.1, -99999.99)},
	}
	for _, cse := range cses {
		t.Run("FltsTrm", func(t *testing.T) {
			var trmr trm.Trmr
			trmr.Reset(cse.lit)
			fltsLit, ok := trmr.FltsLit()
			tst.True(t, ok)
			a := prs.FltsTrm(fltsLit, cse.lit)
			tst.FltsEql(t, cse.val, a)
		})
	}
}
func TestPrsFlts(t *testing.T) {
	cses := []struct {
		lit string
		val *flts.Flts
	}{
		{"[0.0 1.1 3.0]", flts.New(0.0, 1.1, 3.0)},
		{"[0.0 1.1 3.0 3.0 3.0 99999.99 -1.1 -99999.99]", flts.New(0.0, 1.1, 3.0, 3.0, 3.0, 99999.99, -1.1, -99999.99)},
		{"[0.0 1.1 3.0 3.0 3.0 99999.99 -1.1 -99999.99 0.0 1.1 3.0 3.0 3.0 99999.99 -1.1 -99999.99 0.0 1.1 3.0 3.0 3.0 99999.99 -1.1 -99999.99 0.0 1.1 3.0 3.0 3.0 99999.99 -1.1 -99999.99 0.0 1.1 3.0 3.0 3.0 99999.99 -1.1 -99999.99 0.0 1.1 3.0 3.0 3.0 99999.99 -1.1 -99999.99 0.0 1.1 3.0 3.0 3.0 99999.99 -1.1 -99999.99 0.0 1.1 3.0 3.0 3.0 99999.99 -1.1 -99999.99]", flts.New(0.0, 1.1, 3.0, 3.0, 3.0, 99999.99, -1.1, -99999.99, 0.0, 1.1, 3.0, 3.0, 3.0, 99999.99, -1.1, -99999.99, 0.0, 1.1, 3.0, 3.0, 3.0, 99999.99, -1.1, -99999.99, 0.0, 1.1, 3.0, 3.0, 3.0, 99999.99, -1.1, -99999.99, 0.0, 1.1, 3.0, 3.0, 3.0, 99999.99, -1.1, -99999.99, 0.0, 1.1, 3.0, 3.0, 3.0, 99999.99, -1.1, -99999.99, 0.0, 1.1, 3.0, 3.0, 3.0, 99999.99, -1.1, -99999.99, 0.0, 1.1, 3.0, 3.0, 3.0, 99999.99, -1.1, -99999.99)},
	}
	for _, cse := range cses {
		t.Run("Flts", func(t *testing.T) {
			a := prs.Flts(cse.lit)
			tst.FltsEql(t, cse.val, a)
		})
	}
}
func TestPrsUntsTrm(t *testing.T) {
	cses := []struct {
		lit string
		val *unts.Unts
	}{
		{"[0 1 1000]", unts.New(0, 1, 1000)},
		{"[0 1 1000 10]", unts.New(0, 1, 1000, 10)},
	}
	for _, cse := range cses {
		t.Run("UntsTrm", func(t *testing.T) {
			var trmr trm.Trmr
			trmr.Reset(cse.lit)
			untsLit, ok := trmr.UntsLit()
			tst.True(t, ok)
			a := prs.UntsTrm(untsLit, cse.lit)
			tst.UntsEql(t, cse.val, a)
		})
	}
}
func TestPrsUnts(t *testing.T) {
	cses := []struct {
		lit string
		val *unts.Unts
	}{
		{"[0 1 1000]", unts.New(0, 1, 1000)},
		{"[0 1 1000 10]", unts.New(0, 1, 1000, 10)},
	}
	for _, cse := range cses {
		t.Run("Unts", func(t *testing.T) {
			a := prs.Unts(cse.lit)
			tst.UntsEql(t, cse.val, a)
		})
	}
}
func TestPrsIntsTrm(t *testing.T) {
	cses := []struct {
		lit string
		val *ints.Ints
	}{
		{"[+0 +10 +1000]", ints.New(0, 10, 1000)},
		{"[+0 +10 +1000 -10 -1000]", ints.New(0, 10, 1000, -10, -1000)},
	}
	for _, cse := range cses {
		t.Run("IntsTrm", func(t *testing.T) {
			var trmr trm.Trmr
			trmr.Reset(cse.lit)
			intsLit, ok := trmr.IntsLit()
			tst.True(t, ok)
			a := prs.IntsTrm(intsLit, cse.lit)
			tst.IntsEql(t, cse.val, a)
		})
	}
}
func TestPrsInts(t *testing.T) {
	cses := []struct {
		lit string
		val *ints.Ints
	}{
		{"[+0 +10 +1000]", ints.New(0, 10, 1000)},
		{"[+0 +10 +1000 -10 -1000]", ints.New(0, 10, 1000, -10, -1000)},
	}
	for _, cse := range cses {
		t.Run("Ints", func(t *testing.T) {
			a := prs.Ints(cse.lit)
			tst.IntsEql(t, cse.val, a)
		})
	}
}
func TestPrsTmesTrm(t *testing.T) {
	cses := []struct {
		lit string
		val *tmes.Tmes
	}{
		{"[0s 1s 10s]", tmes.New(0, 1, 10)},
		{"[0s 1s 10s -1s -10m 1w2d3h4m5s -1w2d3h4m5s 2000y1n2d3h4m5s 2000y2d3h4m5s 2000y3h4m5s]", tmes.New(0, 1, 10, -1, -10*60, 788645, -788645, 946782245, 946782245, 946695845)},
	}
	for _, cse := range cses {
		t.Run("TmesTrm", func(t *testing.T) {
			var trmr trm.Trmr
			trmr.Reset(cse.lit)
			tmesLit, ok := trmr.TmesLit()
			tst.True(t, ok)
			a := prs.TmesTrm(tmesLit, cse.lit)
			tst.TmesEql(t, cse.val, a)
		})
	}
}
func TestPrsTmes(t *testing.T) {
	cses := []struct {
		lit string
		val *tmes.Tmes
	}{
		{"[0s 1s 10s]", tmes.New(0, 1, 10)},
		{"[0s 1s 10s -1s -10m 1w2d3h4m5s -1w2d3h4m5s 2000y1n2d3h4m5s 2000y2d3h4m5s 2000y3h4m5s]", tmes.New(0, 1, 10, -1, -10*60, 788645, -788645, 946782245, 946782245, 946695845)},
	}
	for _, cse := range cses {
		t.Run("Tmes", func(t *testing.T) {
			a := prs.Tmes(cse.lit)
			tst.TmesEql(t, cse.val, a)
		})
	}
}
func TestPrsBndsTrm(t *testing.T) {
	cses := []struct {
		lit string
		val *bnds.Bnds
	}{
		{"[0-0 0-1 0-1000]", bnds.New(bnd.Bnd{Idx: 0, Lim: 0}, bnd.Bnd{Idx: 0, Lim: 1}, bnd.Bnd{Idx: 0, Lim: 1000})},
		{"[0-0 0-1 0-1000 999-1000 1-0]", bnds.New(bnd.Bnd{Idx: 0, Lim: 0}, bnd.Bnd{Idx: 0, Lim: 1}, bnd.Bnd{Idx: 0, Lim: 1000}, bnd.Bnd{Idx: 999, Lim: 1000}, bnd.Bnd{Idx: 1, Lim: 0})},
	}
	for _, cse := range cses {
		t.Run("BndsTrm", func(t *testing.T) {
			var trmr trm.Trmr
			trmr.Reset(cse.lit)
			bndsLit, ok := trmr.BndsLit()
			tst.True(t, ok)
			a := prs.BndsTrm(bndsLit, cse.lit)
			tst.BndsEql(t, cse.val, a)
		})
	}
}
func TestPrsBnds(t *testing.T) {
	cses := []struct {
		lit string
		val *bnds.Bnds
	}{
		{"[0-0 0-1 0-1000]", bnds.New(bnd.Bnd{Idx: 0, Lim: 0}, bnd.Bnd{Idx: 0, Lim: 1}, bnd.Bnd{Idx: 0, Lim: 1000})},
		{"[0-0 0-1 0-1000 999-1000 1-0]", bnds.New(bnd.Bnd{Idx: 0, Lim: 0}, bnd.Bnd{Idx: 0, Lim: 1}, bnd.Bnd{Idx: 0, Lim: 1000}, bnd.Bnd{Idx: 999, Lim: 1000}, bnd.Bnd{Idx: 1, Lim: 0})},
	}
	for _, cse := range cses {
		t.Run("Bnds", func(t *testing.T) {
			a := prs.Bnds(cse.lit)
			tst.BndsEql(t, cse.val, a)
		})
	}
}
func TestPrsTmeRngsTrm(t *testing.T) {
	cses := []struct {
		lit string
		val *tme.Rngs
	}{
		{"[-10s--1s 0s-1s 2s-4s]", tme.NewRngs(tme.Rng{Min: -10, Max: -1}, tme.Rng{Min: 0, Max: 1}, tme.Rng{Min: 2, Max: 4})},
		{"[-10s--1s 0s-1s 2s-4s 6s-10s 50s-60s]", tme.NewRngs(tme.Rng{Min: -10, Max: -1}, tme.Rng{Min: 0, Max: 1}, tme.Rng{Min: 2, Max: 4}, tme.Rng{Min: 6, Max: 10}, tme.Rng{Min: 50, Max: 60})},
	}
	for _, cse := range cses {
		t.Run("TmeRngsTrm", func(t *testing.T) {
			var trmr trm.Trmr
			trmr.Reset(cse.lit)
			tmeRngsLit, ok := trmr.TmeRngsLit()
			tst.True(t, ok)
			a := prs.TmeRngsTrm(tmeRngsLit, cse.lit)
			tst.TmeRngsEql(t, cse.val, a)
		})
	}
}
func TestPrsTmeRngs(t *testing.T) {
	cses := []struct {
		lit string
		val *tme.Rngs
	}{
		{"[-10s--1s 0s-1s 2s-4s]", tme.NewRngs(tme.Rng{Min: -10, Max: -1}, tme.Rng{Min: 0, Max: 1}, tme.Rng{Min: 2, Max: 4})},
		{"[-10s--1s 0s-1s 2s-4s 6s-10s 50s-60s]", tme.NewRngs(tme.Rng{Min: -10, Max: -1}, tme.Rng{Min: 0, Max: 1}, tme.Rng{Min: 2, Max: 4}, tme.Rng{Min: 6, Max: 10}, tme.Rng{Min: 50, Max: 60})},
	}
	for _, cse := range cses {
		t.Run("TmeRngs", func(t *testing.T) {
			a := prs.TmeRngs(cse.lit)
			tst.TmeRngsEql(t, cse.val, a)
		})
	}
}
