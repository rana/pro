package prs_test

import (
	"sys/bsc/bol"
	"sys/bsc/flt"
	"sys/bsc/int"
	"sys/bsc/str"
	"sys/lng/jsn/trm"
	"sys/lng/jsn/trm/prs"
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
func TestPrsStrTxt(t *testing.T) {
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
		t.Run("StrTxt", func(t *testing.T) {
			a := prs.StrTxt(cse.lit)
			tst.StrEql(t, cse.val, a)
		})
	}
}
func TestPrsBolTrm(t *testing.T) {
	cses := []struct {
		lit string
		val bol.Bol
	}{
		{"false", false},
		{"true", true},
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
func TestPrsBolTxt(t *testing.T) {
	cses := []struct {
		lit string
		val bol.Bol
	}{
		{"false", false},
		{"true", true},
	}
	for _, cse := range cses {
		t.Run("BolTxt", func(t *testing.T) {
			a := prs.BolTxt(cse.lit)
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
func TestPrsFltTxt(t *testing.T) {
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
		t.Run("FltTxt", func(t *testing.T) {
			a := prs.FltTxt(cse.lit)
			tst.FltEql(t, cse.val, a)
		})
	}
}
func TestPrsIntTrm(t *testing.T) {
	cses := []struct {
		lit string
		val int.Int
	}{
		{"0", 0},
		{"10", 10},
		{"1000", 1000},
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
func TestPrsIntTxt(t *testing.T) {
	cses := []struct {
		lit string
		val int.Int
	}{
		{"0", 0},
		{"10", 10},
		{"1000", 1000},
		{"-10", -10},
		{"-1000", -1000},
	}
	for _, cse := range cses {
		t.Run("IntTxt", func(t *testing.T) {
			a := prs.IntTxt(cse.lit)
			tst.IntEql(t, cse.val, a)
		})
	}
}
