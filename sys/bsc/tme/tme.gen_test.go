package tme_test

import (
	"bytes"
	"fmt"
	"math"
	"sys/bsc/bol"
	"sys/bsc/tme"
	"sys/tst"
	"testing"
)

func TestTmeTmeCnst(t *testing.T) {
	cses := []struct {
		idn string
		e   tme.Tme
		a   tme.Tme
	}{
		{"Zero", tme.Tme(0), tme.Zero},
		{"One", tme.Tme(1), tme.One},
		{"NegOne", tme.Tme(-1), tme.NegOne},
		{"Min", tme.Tme(-1 << 31), tme.Min},
		{"Max", tme.Tme(1<<31 - 1), tme.Max},
		{"Second", tme.Tme(1), tme.Second},
		{"Minute", tme.Tme(60), tme.Minute},
		{"Hour", tme.Tme(60 * 60), tme.Hour},
		{"Day", tme.Tme(24 * 60 * 60), tme.Day},
		{"Week", tme.Tme(7 * 24 * 60 * 60), tme.Week},
		{"S1", tme.Tme(1), tme.S1},
		{"S5", tme.Tme(5), tme.S5},
		{"S10", tme.Tme(10), tme.S10},
		{"S15", tme.Tme(15), tme.S15},
		{"S20", tme.Tme(20), tme.S20},
		{"S30", tme.Tme(30), tme.S30},
		{"S40", tme.Tme(40), tme.S40},
		{"S50", tme.Tme(50), tme.S50},
		{"M1", tme.Tme(1 * 60), tme.M1},
		{"M5", tme.Tme(5 * 60), tme.M5},
		{"M10", tme.Tme(10 * 60), tme.M10},
		{"M15", tme.Tme(15 * 60), tme.M15},
		{"M20", tme.Tme(20 * 60), tme.M20},
		{"M30", tme.Tme(30 * 60), tme.M30},
		{"M40", tme.Tme(40 * 60), tme.M40},
		{"M50", tme.Tme(50 * 60), tme.M50},
		{"H1", tme.Tme(1 * 60 * 60), tme.H1},
		{"D1", tme.Tme(1 * 60 * 60 * 24), tme.D1},
		{"DurStrLim", tme.Tme(1 * 60 * 60 * 24 * 365 * 10), tme.DurStrLim},
		{"Resolution", tme.Tme(1), tme.Resolution},
	}
	for _, cse := range cses {
		t.Run(fmt.Sprintf("%q", cse.idn), func(t *testing.T) {
			tst.TmeEql(t, cse.e, cse.a)
		})
	}
}
func TestTmeTmeEqlPkgFn(t *testing.T) {
	cses := []struct {
		a tme.Tme
		b tme.Tme
	}{
		{0, 0},
		{0, 1},
		{0, 10},
		{0, -1},
		{0, -10 * 60},
		{0, 788645},
		{0, -788645},
		{0, 946782245},
		{0, 946782245},
		{0, 946695845},
		{1, 0},
		{1, 1},
		{1, 10},
		{1, -1},
		{1, -10 * 60},
		{1, 788645},
		{1, -788645},
		{1, 946782245},
		{1, 946782245},
		{1, 946695845},
		{10, 0},
		{10, 1},
		{10, 10},
		{10, -1},
		{10, -10 * 60},
		{10, 788645},
		{10, -788645},
		{10, 946782245},
		{10, 946782245},
		{10, 946695845},
		{-1, 0},
		{-1, 1},
		{-1, 10},
		{-1, -1},
		{-1, -10 * 60},
		{-1, 788645},
		{-1, -788645},
		{-1, 946782245},
		{-1, 946782245},
		{-1, 946695845},
		{-10 * 60, 0},
		{-10 * 60, 1},
		{-10 * 60, 10},
		{-10 * 60, -1},
		{-10 * 60, -10 * 60},
		{-10 * 60, 788645},
		{-10 * 60, -788645},
		{-10 * 60, 946782245},
		{-10 * 60, 946782245},
		{-10 * 60, 946695845},
		{788645, 0},
		{788645, 1},
		{788645, 10},
		{788645, -1},
		{788645, -10 * 60},
		{788645, 788645},
		{788645, -788645},
		{788645, 946782245},
		{788645, 946782245},
		{788645, 946695845},
		{-788645, 0},
		{-788645, 1},
		{-788645, 10},
		{-788645, -1},
		{-788645, -10 * 60},
		{-788645, 788645},
		{-788645, -788645},
		{-788645, 946782245},
		{-788645, 946782245},
		{-788645, 946695845},
		{946782245, 0},
		{946782245, 1},
		{946782245, 10},
		{946782245, -1},
		{946782245, -10 * 60},
		{946782245, 788645},
		{946782245, -788645},
		{946782245, 946782245},
		{946782245, 946782245},
		{946782245, 946695845},
		{946782245, 0},
		{946782245, 1},
		{946782245, 10},
		{946782245, -1},
		{946782245, -10 * 60},
		{946782245, 788645},
		{946782245, -788645},
		{946782245, 946782245},
		{946782245, 946782245},
		{946782245, 946695845},
		{946695845, 0},
		{946695845, 1},
		{946695845, 10},
		{946695845, -1},
		{946695845, -10 * 60},
		{946695845, 788645},
		{946695845, -788645},
		{946695845, 946782245},
		{946695845, 946782245},
		{946695845, 946695845},
	}
	for _, cse := range cses {
		t.Run("Eql", func(t *testing.T) {
			a := cse.a
			b := cse.b
			actual := tme.Eql(a, b)
			expected := bol.Bol(a == b)
			tst.BolEql(t, expected, actual)
		})
	}
}
func TestTmeTmeLssPkgFn(t *testing.T) {
	cses := []struct {
		a tme.Tme
		b tme.Tme
	}{
		{0, 0},
		{0, 1},
		{0, 10},
		{0, -1},
		{0, -10 * 60},
		{0, 788645},
		{0, -788645},
		{0, 946782245},
		{0, 946782245},
		{0, 946695845},
		{1, 0},
		{1, 1},
		{1, 10},
		{1, -1},
		{1, -10 * 60},
		{1, 788645},
		{1, -788645},
		{1, 946782245},
		{1, 946782245},
		{1, 946695845},
		{10, 0},
		{10, 1},
		{10, 10},
		{10, -1},
		{10, -10 * 60},
		{10, 788645},
		{10, -788645},
		{10, 946782245},
		{10, 946782245},
		{10, 946695845},
		{-1, 0},
		{-1, 1},
		{-1, 10},
		{-1, -1},
		{-1, -10 * 60},
		{-1, 788645},
		{-1, -788645},
		{-1, 946782245},
		{-1, 946782245},
		{-1, 946695845},
		{-10 * 60, 0},
		{-10 * 60, 1},
		{-10 * 60, 10},
		{-10 * 60, -1},
		{-10 * 60, -10 * 60},
		{-10 * 60, 788645},
		{-10 * 60, -788645},
		{-10 * 60, 946782245},
		{-10 * 60, 946782245},
		{-10 * 60, 946695845},
		{788645, 0},
		{788645, 1},
		{788645, 10},
		{788645, -1},
		{788645, -10 * 60},
		{788645, 788645},
		{788645, -788645},
		{788645, 946782245},
		{788645, 946782245},
		{788645, 946695845},
		{-788645, 0},
		{-788645, 1},
		{-788645, 10},
		{-788645, -1},
		{-788645, -10 * 60},
		{-788645, 788645},
		{-788645, -788645},
		{-788645, 946782245},
		{-788645, 946782245},
		{-788645, 946695845},
		{946782245, 0},
		{946782245, 1},
		{946782245, 10},
		{946782245, -1},
		{946782245, -10 * 60},
		{946782245, 788645},
		{946782245, -788645},
		{946782245, 946782245},
		{946782245, 946782245},
		{946782245, 946695845},
		{946782245, 0},
		{946782245, 1},
		{946782245, 10},
		{946782245, -1},
		{946782245, -10 * 60},
		{946782245, 788645},
		{946782245, -788645},
		{946782245, 946782245},
		{946782245, 946782245},
		{946782245, 946695845},
		{946695845, 0},
		{946695845, 1},
		{946695845, 10},
		{946695845, -1},
		{946695845, -10 * 60},
		{946695845, 788645},
		{946695845, -788645},
		{946695845, 946782245},
		{946695845, 946782245},
		{946695845, 946695845},
	}
	for _, cse := range cses {
		t.Run("Lss", func(t *testing.T) {
			a := cse.a
			b := cse.b
			actual := tme.Lss(a, b)
			expected := bol.Bol(a < b)
			tst.BolEql(t, expected, actual)
		})
	}
}
func TestTmeTmeGtrPkgFn(t *testing.T) {
	cses := []struct {
		a tme.Tme
		b tme.Tme
	}{
		{0, 0},
		{0, 1},
		{0, 10},
		{0, -1},
		{0, -10 * 60},
		{0, 788645},
		{0, -788645},
		{0, 946782245},
		{0, 946782245},
		{0, 946695845},
		{1, 0},
		{1, 1},
		{1, 10},
		{1, -1},
		{1, -10 * 60},
		{1, 788645},
		{1, -788645},
		{1, 946782245},
		{1, 946782245},
		{1, 946695845},
		{10, 0},
		{10, 1},
		{10, 10},
		{10, -1},
		{10, -10 * 60},
		{10, 788645},
		{10, -788645},
		{10, 946782245},
		{10, 946782245},
		{10, 946695845},
		{-1, 0},
		{-1, 1},
		{-1, 10},
		{-1, -1},
		{-1, -10 * 60},
		{-1, 788645},
		{-1, -788645},
		{-1, 946782245},
		{-1, 946782245},
		{-1, 946695845},
		{-10 * 60, 0},
		{-10 * 60, 1},
		{-10 * 60, 10},
		{-10 * 60, -1},
		{-10 * 60, -10 * 60},
		{-10 * 60, 788645},
		{-10 * 60, -788645},
		{-10 * 60, 946782245},
		{-10 * 60, 946782245},
		{-10 * 60, 946695845},
		{788645, 0},
		{788645, 1},
		{788645, 10},
		{788645, -1},
		{788645, -10 * 60},
		{788645, 788645},
		{788645, -788645},
		{788645, 946782245},
		{788645, 946782245},
		{788645, 946695845},
		{-788645, 0},
		{-788645, 1},
		{-788645, 10},
		{-788645, -1},
		{-788645, -10 * 60},
		{-788645, 788645},
		{-788645, -788645},
		{-788645, 946782245},
		{-788645, 946782245},
		{-788645, 946695845},
		{946782245, 0},
		{946782245, 1},
		{946782245, 10},
		{946782245, -1},
		{946782245, -10 * 60},
		{946782245, 788645},
		{946782245, -788645},
		{946782245, 946782245},
		{946782245, 946782245},
		{946782245, 946695845},
		{946782245, 0},
		{946782245, 1},
		{946782245, 10},
		{946782245, -1},
		{946782245, -10 * 60},
		{946782245, 788645},
		{946782245, -788645},
		{946782245, 946782245},
		{946782245, 946782245},
		{946782245, 946695845},
		{946695845, 0},
		{946695845, 1},
		{946695845, 10},
		{946695845, -1},
		{946695845, -10 * 60},
		{946695845, 788645},
		{946695845, -788645},
		{946695845, 946782245},
		{946695845, 946782245},
		{946695845, 946695845},
	}
	for _, cse := range cses {
		t.Run("Gtr", func(t *testing.T) {
			a := cse.a
			b := cse.b
			actual := tme.Gtr(a, b)
			expected := bol.Bol(a > b)
			tst.BolEql(t, expected, actual)
		})
	}
}
func TestTmeTmeEqlTypFn(t *testing.T) {
	cses := []struct {
		x tme.Tme
		a tme.Tme
	}{
		{tme.Tme(0), tme.Tme(1)},
		{tme.Tme(1), tme.Tme(1)},
		{tme.Tme(10), tme.Tme(1)},
		{tme.Tme(-1), tme.Tme(1)},
		{tme.Tme(-10 * 60), tme.Tme(1)},
		{tme.Tme(788645), tme.Tme(1)},
		{tme.Tme(-788645), tme.Tme(1)},
		{tme.Tme(946782245), tme.Tme(1)},
		{tme.Tme(946782245), tme.Tme(1)},
		{tme.Tme(946695845), tme.Tme(1)},
	}
	for _, cse := range cses {
		t.Run("Eql", func(t *testing.T) {
			x := cse.x
			a := cse.a
			actual := x.Eql(a)
			expected := bol.Bol(x == a)
			tst.BolEql(t, expected, actual)
		})
	}
}
func TestTmeTmeNeqTypFn(t *testing.T) {
	cses := []struct {
		x tme.Tme
		a tme.Tme
	}{
		{tme.Tme(0), tme.Tme(1)},
		{tme.Tme(1), tme.Tme(1)},
		{tme.Tme(10), tme.Tme(1)},
		{tme.Tme(-1), tme.Tme(1)},
		{tme.Tme(-10 * 60), tme.Tme(1)},
		{tme.Tme(788645), tme.Tme(1)},
		{tme.Tme(-788645), tme.Tme(1)},
		{tme.Tme(946782245), tme.Tme(1)},
		{tme.Tme(946782245), tme.Tme(1)},
		{tme.Tme(946695845), tme.Tme(1)},
	}
	for _, cse := range cses {
		t.Run("Neq", func(t *testing.T) {
			x := cse.x
			a := cse.a
			actual := x.Neq(a)
			expected := bol.Bol(x != a)
			tst.BolEql(t, expected, actual)
		})
	}
}
func TestTmeTmeLssTypFn(t *testing.T) {
	cses := []struct {
		x tme.Tme
		a tme.Tme
	}{
		{tme.Tme(0), tme.Tme(1)},
		{tme.Tme(1), tme.Tme(1)},
		{tme.Tme(10), tme.Tme(1)},
		{tme.Tme(-1), tme.Tme(1)},
		{tme.Tme(-10 * 60), tme.Tme(1)},
		{tme.Tme(788645), tme.Tme(1)},
		{tme.Tme(-788645), tme.Tme(1)},
		{tme.Tme(946782245), tme.Tme(1)},
		{tme.Tme(946782245), tme.Tme(1)},
		{tme.Tme(946695845), tme.Tme(1)},
	}
	for _, cse := range cses {
		t.Run("Lss", func(t *testing.T) {
			x := cse.x
			a := cse.a
			actual := x.Lss(a)
			expected := bol.Bol(x < a)
			tst.BolEql(t, expected, actual)
		})
	}
}
func TestTmeTmeGtrTypFn(t *testing.T) {
	cses := []struct {
		x tme.Tme
		a tme.Tme
	}{
		{tme.Tme(0), tme.Tme(1)},
		{tme.Tme(1), tme.Tme(1)},
		{tme.Tme(10), tme.Tme(1)},
		{tme.Tme(-1), tme.Tme(1)},
		{tme.Tme(-10 * 60), tme.Tme(1)},
		{tme.Tme(788645), tme.Tme(1)},
		{tme.Tme(-788645), tme.Tme(1)},
		{tme.Tme(946782245), tme.Tme(1)},
		{tme.Tme(946782245), tme.Tme(1)},
		{tme.Tme(946695845), tme.Tme(1)},
	}
	for _, cse := range cses {
		t.Run("Gtr", func(t *testing.T) {
			x := cse.x
			a := cse.a
			actual := x.Gtr(a)
			expected := bol.Bol(x > a)
			tst.BolEql(t, expected, actual)
		})
	}
}
func TestTmeTmeLeqTypFn(t *testing.T) {
	cses := []struct {
		x tme.Tme
		a tme.Tme
	}{
		{tme.Tme(0), tme.Tme(1)},
		{tme.Tme(1), tme.Tme(1)},
		{tme.Tme(10), tme.Tme(1)},
		{tme.Tme(-1), tme.Tme(1)},
		{tme.Tme(-10 * 60), tme.Tme(1)},
		{tme.Tme(788645), tme.Tme(1)},
		{tme.Tme(-788645), tme.Tme(1)},
		{tme.Tme(946782245), tme.Tme(1)},
		{tme.Tme(946782245), tme.Tme(1)},
		{tme.Tme(946695845), tme.Tme(1)},
	}
	for _, cse := range cses {
		t.Run("Leq", func(t *testing.T) {
			x := cse.x
			a := cse.a
			actual := x.Leq(a)
			expected := bol.Bol(cse.x <= cse.a)
			tst.BolEql(t, expected, actual)
		})
	}
}
func TestTmeTmeGeqTypFn(t *testing.T) {
	cses := []struct {
		x tme.Tme
		a tme.Tme
	}{
		{tme.Tme(0), tme.Tme(1)},
		{tme.Tme(1), tme.Tme(1)},
		{tme.Tme(10), tme.Tme(1)},
		{tme.Tme(-1), tme.Tme(1)},
		{tme.Tme(-10 * 60), tme.Tme(1)},
		{tme.Tme(788645), tme.Tme(1)},
		{tme.Tme(-788645), tme.Tme(1)},
		{tme.Tme(946782245), tme.Tme(1)},
		{tme.Tme(946782245), tme.Tme(1)},
		{tme.Tme(946695845), tme.Tme(1)},
	}
	for _, cse := range cses {
		t.Run("Geq", func(t *testing.T) {
			x := cse.x
			a := cse.a
			actual := x.Geq(a)
			expected := bol.Bol(cse.x >= cse.a)
			tst.BolEql(t, expected, actual)
		})
	}
}
func TestTmeTmePosTypFn(t *testing.T) {
	cses := []struct {
		x tme.Tme
	}{
		{tme.Tme(0)},
		{tme.Tme(1)},
		{tme.Tme(10)},
		{tme.Tme(-1)},
		{tme.Tme(-10 * 60)},
		{tme.Tme(788645)},
		{tme.Tme(-788645)},
		{tme.Tme(946782245)},
		{tme.Tme(946782245)},
		{tme.Tme(946695845)},
	}
	for _, cse := range cses {
		t.Run("Pos", func(t *testing.T) {
			x := cse.x
			actual := x.Pos()
			expected := cse.x
			if expected < 0 {
				expected = -expected
			}
			tst.TmeEql(t, expected, actual)
		})
	}
}
func TestTmeTmeNegTypFn(t *testing.T) {
	cses := []struct {
		x tme.Tme
	}{
		{tme.Tme(0)},
		{tme.Tme(1)},
		{tme.Tme(10)},
		{tme.Tme(-1)},
		{tme.Tme(-10 * 60)},
		{tme.Tme(788645)},
		{tme.Tme(-788645)},
		{tme.Tme(946782245)},
		{tme.Tme(946782245)},
		{tme.Tme(946695845)},
	}
	for _, cse := range cses {
		t.Run("Neg", func(t *testing.T) {
			x := cse.x
			actual := x.Neg()
			expected := cse.x
			if expected > 0 {
				expected = -expected
			}
			tst.TmeEql(t, expected, actual)
		})
	}
}
func TestTmeTmeInvTypFn(t *testing.T) {
	cses := []struct {
		x tme.Tme
	}{
		{tme.Tme(0)},
		{tme.Tme(1)},
		{tme.Tme(10)},
		{tme.Tme(-1)},
		{tme.Tme(-10 * 60)},
		{tme.Tme(788645)},
		{tme.Tme(-788645)},
		{tme.Tme(946782245)},
		{tme.Tme(946782245)},
		{tme.Tme(946695845)},
	}
	for _, cse := range cses {
		t.Run("Inv", func(t *testing.T) {
			x := cse.x
			actual := x.Inv()
			expected := -cse.x
			tst.TmeEql(t, expected, actual)
		})
	}
}
func TestTmeTmeAddTypFn(t *testing.T) {
	cses := []struct {
		x tme.Tme
		a tme.Tme
	}{
		{tme.Tme(0), tme.Tme(1)},
		{tme.Tme(1), tme.Tme(1)},
		{tme.Tme(10), tme.Tme(1)},
		{tme.Tme(-1), tme.Tme(1)},
		{tme.Tme(-10 * 60), tme.Tme(1)},
		{tme.Tme(788645), tme.Tme(1)},
		{tme.Tme(-788645), tme.Tme(1)},
		{tme.Tme(946782245), tme.Tme(1)},
		{tme.Tme(946782245), tme.Tme(1)},
		{tme.Tme(946695845), tme.Tme(1)},
	}
	for _, cse := range cses {
		t.Run("Add", func(t *testing.T) {
			x := cse.x
			a := cse.a
			actual := x.Add(a)
			expected := x + a
			tst.TmeEql(t, expected, actual)
		})
	}
}
func TestTmeTmeSubTypFn(t *testing.T) {
	cses := []struct {
		x tme.Tme
		a tme.Tme
	}{
		{tme.Tme(0), tme.Tme(1)},
		{tme.Tme(1), tme.Tme(1)},
		{tme.Tme(10), tme.Tme(1)},
		{tme.Tme(-1), tme.Tme(1)},
		{tme.Tme(-10 * 60), tme.Tme(1)},
		{tme.Tme(788645), tme.Tme(1)},
		{tme.Tme(-788645), tme.Tme(1)},
		{tme.Tme(946782245), tme.Tme(1)},
		{tme.Tme(946782245), tme.Tme(1)},
		{tme.Tme(946695845), tme.Tme(1)},
	}
	for _, cse := range cses {
		t.Run("Sub", func(t *testing.T) {
			x := cse.x
			a := cse.a
			actual := x.Sub(a)
			expected := x - a
			tst.TmeEql(t, expected, actual)
		})
	}
}
func TestTmeTmeMulTypFn(t *testing.T) {
	cses := []struct {
		x tme.Tme
		a tme.Tme
	}{
		{tme.Tme(0), tme.Tme(1)},
		{tme.Tme(1), tme.Tme(1)},
		{tme.Tme(10), tme.Tme(1)},
		{tme.Tme(-1), tme.Tme(1)},
		{tme.Tme(-10 * 60), tme.Tme(1)},
		{tme.Tme(788645), tme.Tme(1)},
		{tme.Tme(-788645), tme.Tme(1)},
		{tme.Tme(946782245), tme.Tme(1)},
		{tme.Tme(946782245), tme.Tme(1)},
		{tme.Tme(946695845), tme.Tme(1)},
	}
	for _, cse := range cses {
		t.Run("Mul", func(t *testing.T) {
			x := cse.x
			a := cse.a
			actual := x.Mul(a)
			expected := x * a
			tst.TmeEql(t, expected, actual)
		})
	}
}
func TestTmeTmeDivTypFn(t *testing.T) {
	cses := []struct {
		x tme.Tme
		a tme.Tme
	}{
		{tme.Tme(0), tme.Tme(1)},
		{tme.Tme(1), tme.Tme(1)},
		{tme.Tme(10), tme.Tme(1)},
		{tme.Tme(-1), tme.Tme(1)},
		{tme.Tme(-10 * 60), tme.Tme(1)},
		{tme.Tme(788645), tme.Tme(1)},
		{tme.Tme(-788645), tme.Tme(1)},
		{tme.Tme(946782245), tme.Tme(1)},
		{tme.Tme(946782245), tme.Tme(1)},
		{tme.Tme(946695845), tme.Tme(1)},
	}
	for _, cse := range cses {
		t.Run("Div", func(t *testing.T) {
			x := cse.x
			a := cse.a
			actual := x.Div(a)
			var expected tme.Tme
			if a != 0 {
				expected = x / a
			}
			tst.TmeEql(t, expected, actual)
		})
	}
}
func TestTmeTmeRemTypFn(t *testing.T) {
	cses := []struct {
		x tme.Tme
		a tme.Tme
	}{
		{tme.Tme(0), tme.Tme(1)},
		{tme.Tme(1), tme.Tme(1)},
		{tme.Tme(10), tme.Tme(1)},
		{tme.Tme(-1), tme.Tme(1)},
		{tme.Tme(-10 * 60), tme.Tme(1)},
		{tme.Tme(788645), tme.Tme(1)},
		{tme.Tme(-788645), tme.Tme(1)},
		{tme.Tme(946782245), tme.Tme(1)},
		{tme.Tme(946782245), tme.Tme(1)},
		{tme.Tme(946695845), tme.Tme(1)},
	}
	for _, cse := range cses {
		t.Run("Rem", func(t *testing.T) {
			x := cse.x
			a := cse.a
			actual := x.Rem(a)
			var expected tme.Tme
			if a != 0 {
				expected = tme.Tme(math.Remainder(float64(x), float64(a)))
			}
			tst.TmeEql(t, expected, actual)
		})
	}
}
func TestTmeTmePowTypFn(t *testing.T) {
	cses := []struct {
		x tme.Tme
		a tme.Tme
	}{
		{tme.Tme(0), tme.Tme(1)},
		{tme.Tme(1), tme.Tme(1)},
		{tme.Tme(10), tme.Tme(1)},
		{tme.Tme(-1), tme.Tme(1)},
		{tme.Tme(-10 * 60), tme.Tme(1)},
		{tme.Tme(788645), tme.Tme(1)},
		{tme.Tme(-788645), tme.Tme(1)},
		{tme.Tme(946782245), tme.Tme(1)},
		{tme.Tme(946782245), tme.Tme(1)},
		{tme.Tme(946695845), tme.Tme(1)},
	}
	for _, cse := range cses {
		t.Run("Pow", func(t *testing.T) {
			x := cse.x
			a := cse.a
			actual := x.Pow(a)
			expected := tme.Tme(math.Pow(float64(x), float64(a)))
			tst.TmeEql(t, expected, actual)
		})
	}
}
func TestTmeTmeSqrTypFn(t *testing.T) {
	cses := []struct {
		x tme.Tme
	}{
		{tme.Tme(0)},
		{tme.Tme(1)},
		{tme.Tme(10)},
		{tme.Tme(-1)},
		{tme.Tme(-10 * 60)},
		{tme.Tme(788645)},
		{tme.Tme(-788645)},
		{tme.Tme(946782245)},
		{tme.Tme(946782245)},
		{tme.Tme(946695845)},
	}
	for _, cse := range cses {
		t.Run("Sqr", func(t *testing.T) {
			x := cse.x
			actual := x.Sqr()
			expected := x * x
			tst.TmeEql(t, expected, actual)
		})
	}
}
func TestTmeTmeSqrtTypFn(t *testing.T) {
	cses := []struct {
		x tme.Tme
	}{
		{tme.Tme(0)},
		{tme.Tme(1)},
		{tme.Tme(10)},
		{tme.Tme(-1)},
		{tme.Tme(-10 * 60)},
		{tme.Tme(788645)},
		{tme.Tme(-788645)},
		{tme.Tme(946782245)},
		{tme.Tme(946782245)},
		{tme.Tme(946695845)},
	}
	for _, cse := range cses {
		t.Run("Sqrt", func(t *testing.T) {
			x := cse.x
			actual := x.Sqrt()
			var expected tme.Tme
			if x > 0 {
				expected = tme.Tme(math.Sqrt(float64(x)))
			}
			tst.TmeEql(t, expected, actual)
		})
	}
}
func TestTmeTmeMinTypFn(t *testing.T) {
	cses := []struct {
		x tme.Tme
		a tme.Tme
	}{
		{tme.Tme(0), tme.Tme(1)},
		{tme.Tme(1), tme.Tme(1)},
		{tme.Tme(10), tme.Tme(1)},
		{tme.Tme(-1), tme.Tme(1)},
		{tme.Tme(-10 * 60), tme.Tme(1)},
		{tme.Tme(788645), tme.Tme(1)},
		{tme.Tme(-788645), tme.Tme(1)},
		{tme.Tme(946782245), tme.Tme(1)},
		{tme.Tme(946782245), tme.Tme(1)},
		{tme.Tme(946695845), tme.Tme(1)},
	}
	for _, cse := range cses {
		t.Run("Min", func(t *testing.T) {
			x := cse.x
			a := cse.a
			actual := x.Min(a)
			expected := x
			if a < expected {
				expected = a
			}
			tst.TmeEql(t, expected, actual)
		})
	}
}
func TestTmeTmeMaxTypFn(t *testing.T) {
	cses := []struct {
		x tme.Tme
		a tme.Tme
	}{
		{tme.Tme(0), tme.Tme(1)},
		{tme.Tme(1), tme.Tme(1)},
		{tme.Tme(10), tme.Tme(1)},
		{tme.Tme(-1), tme.Tme(1)},
		{tme.Tme(-10 * 60), tme.Tme(1)},
		{tme.Tme(788645), tme.Tme(1)},
		{tme.Tme(-788645), tme.Tme(1)},
		{tme.Tme(946782245), tme.Tme(1)},
		{tme.Tme(946782245), tme.Tme(1)},
		{tme.Tme(946695845), tme.Tme(1)},
	}
	for _, cse := range cses {
		t.Run("Max", func(t *testing.T) {
			x := cse.x
			a := cse.a
			actual := x.Max(a)
			expected := x
			if a > expected {
				expected = a
			}
			tst.TmeEql(t, expected, actual)
		})
	}
}
func TestTmeTmeMinMaxTypFn(t *testing.T) {
	cses := []struct {
		x tme.Tme
		a tme.Tme
	}{
		{tme.Tme(0), tme.Tme(1)},
		{tme.Tme(1), tme.Tme(1)},
		{tme.Tme(10), tme.Tme(1)},
		{tme.Tme(-1), tme.Tme(1)},
		{tme.Tme(-10 * 60), tme.Tme(1)},
		{tme.Tme(788645), tme.Tme(1)},
		{tme.Tme(-788645), tme.Tme(1)},
		{tme.Tme(946782245), tme.Tme(1)},
		{tme.Tme(946782245), tme.Tme(1)},
		{tme.Tme(946695845), tme.Tme(1)},
	}
	for _, cse := range cses {
		t.Run("MinMax", func(t *testing.T) {
			x := cse.x
			a := cse.a
			aMin, aMax := x.MinMax(a)
			var eMin, eMax tme.Tme
			if x < a {
				eMin, eMax = x, a
			} else {
				eMin, eMax = a, x
			}
			tst.TmeEql(t, eMin, aMin, "Min")
			tst.TmeEql(t, eMax, aMax, "Max")
		})
	}
}
func TestTmeTmeMidTypFn(t *testing.T) {
	cses := []struct {
		x tme.Tme
		a tme.Tme
	}{
		{tme.Tme(0), tme.Tme(1)},
		{tme.Tme(1), tme.Tme(1)},
		{tme.Tme(10), tme.Tme(1)},
		{tme.Tme(-1), tme.Tme(1)},
		{tme.Tme(-10 * 60), tme.Tme(1)},
		{tme.Tme(788645), tme.Tme(1)},
		{tme.Tme(-788645), tme.Tme(1)},
		{tme.Tme(946782245), tme.Tme(1)},
		{tme.Tme(946782245), tme.Tme(1)},
		{tme.Tme(946695845), tme.Tme(1)},
	}
	for _, cse := range cses {
		t.Run("Mid", func(t *testing.T) {
			x := cse.x
			a := cse.a
			actual := x.Mid(a)
			expected := (x - a) / 2
			tst.TmeEql(t, expected, actual)
		})
	}
}
func TestTmeTmeAvgTypFn(t *testing.T) {
	cses := []struct {
		x tme.Tme
		a tme.Tme
	}{
		{tme.Tme(0), tme.Tme(1)},
		{tme.Tme(1), tme.Tme(1)},
		{tme.Tme(10), tme.Tme(1)},
		{tme.Tme(-1), tme.Tme(1)},
		{tme.Tme(-10 * 60), tme.Tme(1)},
		{tme.Tme(788645), tme.Tme(1)},
		{tme.Tme(-788645), tme.Tme(1)},
		{tme.Tme(946782245), tme.Tme(1)},
		{tme.Tme(946782245), tme.Tme(1)},
		{tme.Tme(946695845), tme.Tme(1)},
	}
	for _, cse := range cses {
		t.Run("Avg", func(t *testing.T) {
			x := cse.x
			a := cse.a
			actual := x.Avg(a)
			expected := (x + a) / 2
			tst.TmeEql(t, expected, actual)
		})
	}
}
func TestTmeTmeAvgGeoTypFn(t *testing.T) {
	cses := []struct {
		x tme.Tme
		a tme.Tme
	}{
		{tme.Tme(0), tme.Tme(1)},
		{tme.Tme(1), tme.Tme(1)},
		{tme.Tme(10), tme.Tme(1)},
		{tme.Tme(-1), tme.Tme(1)},
		{tme.Tme(-10 * 60), tme.Tme(1)},
		{tme.Tme(788645), tme.Tme(1)},
		{tme.Tme(-788645), tme.Tme(1)},
		{tme.Tme(946782245), tme.Tme(1)},
		{tme.Tme(946782245), tme.Tme(1)},
		{tme.Tme(946695845), tme.Tme(1)},
	}
	for _, cse := range cses {
		t.Run("AvgGeo", func(t *testing.T) {
			x := cse.x
			a := cse.a
			actual := x.AvgGeo(a)
			expected := x * a
			if expected != 0 {
				expected = expected.Sqrt()
			}
			tst.TmeEql(t, expected, actual)
		})
	}
}
func TestTmeTmeByt(t *testing.T) {
	cses := []struct {
		e tme.Tme
	}{
		{0},
		{1},
		{10},
		{-1},
		{-10 * 60},
		{788645},
		{-788645},
		{946782245},
		{946782245},
		{946695845},
	}
	for _, cse := range cses {
		t.Run("Byt", func(t *testing.T) {
			b := &bytes.Buffer{}
			cse.e.BytWrt(b)
			var a tme.Tme
			a.BytRed(b.Bytes())
			tst.TmeEql(t, cse.e, a)
		})
	}
}
