package flt_test

import (
	"bytes"
	"fmt"
	"math"
	"strconv"
	"strings"
	"sys/bsc/bol"
	"sys/bsc/flt"
	"sys/bsc/unt"
	"sys/tst"
	"testing"
)

func TestFltFltCnst(t *testing.T) {
	cses := []struct {
		idn string
		e   flt.Flt
		a   flt.Flt
	}{
		{"Zero", flt.Flt(0.0), flt.Zero},
		{"One", flt.Flt(1.0), flt.One},
		{"NegOne", flt.Flt(-1.0), flt.NegOne},
		{"Hndrd", flt.Flt(100.0), flt.Hndrd},
		{"Min", flt.Flt(-3.40282346638528859811704183484516925440e+38), flt.Min},
		{"Max", flt.Flt(3.40282346638528859811704183484516925440e+38), flt.Max},
		{"Tiny", flt.Flt(1.401298464324817070923729583289916131280e-45), flt.Tiny},
	}
	for _, cse := range cses {
		t.Run(fmt.Sprintf("%q", cse.idn), func(t *testing.T) {
			tst.FltEql(t, cse.e, cse.a)
		})
	}
}
func TestFltFltEqlPkgFn(t *testing.T) {
	cses := []struct {
		a flt.Flt
		b flt.Flt
	}{
		{0.0, 0.0},
		{0.0, 1.1},
		{0.0, 3.0},
		{0.0, 3.0},
		{0.0, 3.0},
		{0.0, 99999.99},
		{0.0, -1.1},
		{0.0, -99999.99},
		{1.1, 0.0},
		{1.1, 1.1},
		{1.1, 3.0},
		{1.1, 3.0},
		{1.1, 3.0},
		{1.1, 99999.99},
		{1.1, -1.1},
		{1.1, -99999.99},
		{3.0, 0.0},
		{3.0, 1.1},
		{3.0, 3.0},
		{3.0, 3.0},
		{3.0, 3.0},
		{3.0, 99999.99},
		{3.0, -1.1},
		{3.0, -99999.99},
		{3.0, 0.0},
		{3.0, 1.1},
		{3.0, 3.0},
		{3.0, 3.0},
		{3.0, 3.0},
		{3.0, 99999.99},
		{3.0, -1.1},
		{3.0, -99999.99},
		{3.0, 0.0},
		{3.0, 1.1},
		{3.0, 3.0},
		{3.0, 3.0},
		{3.0, 3.0},
		{3.0, 99999.99},
		{3.0, -1.1},
		{3.0, -99999.99},
		{99999.99, 0.0},
		{99999.99, 1.1},
		{99999.99, 3.0},
		{99999.99, 3.0},
		{99999.99, 3.0},
		{99999.99, 99999.99},
		{99999.99, -1.1},
		{99999.99, -99999.99},
		{-1.1, 0.0},
		{-1.1, 1.1},
		{-1.1, 3.0},
		{-1.1, 3.0},
		{-1.1, 3.0},
		{-1.1, 99999.99},
		{-1.1, -1.1},
		{-1.1, -99999.99},
		{-99999.99, 0.0},
		{-99999.99, 1.1},
		{-99999.99, 3.0},
		{-99999.99, 3.0},
		{-99999.99, 3.0},
		{-99999.99, 99999.99},
		{-99999.99, -1.1},
		{-99999.99, -99999.99},
	}
	for _, cse := range cses {
		t.Run("Eql", func(t *testing.T) {
			a := cse.a
			b := cse.b
			actual := flt.Eql(a, b)
			expected := bol.Bol(a == b)
			tst.BolEql(t, expected, actual)
		})
	}
}
func TestFltFltLssPkgFn(t *testing.T) {
	cses := []struct {
		a flt.Flt
		b flt.Flt
	}{
		{0.0, 0.0},
		{0.0, 1.1},
		{0.0, 3.0},
		{0.0, 3.0},
		{0.0, 3.0},
		{0.0, 99999.99},
		{0.0, -1.1},
		{0.0, -99999.99},
		{1.1, 0.0},
		{1.1, 1.1},
		{1.1, 3.0},
		{1.1, 3.0},
		{1.1, 3.0},
		{1.1, 99999.99},
		{1.1, -1.1},
		{1.1, -99999.99},
		{3.0, 0.0},
		{3.0, 1.1},
		{3.0, 3.0},
		{3.0, 3.0},
		{3.0, 3.0},
		{3.0, 99999.99},
		{3.0, -1.1},
		{3.0, -99999.99},
		{3.0, 0.0},
		{3.0, 1.1},
		{3.0, 3.0},
		{3.0, 3.0},
		{3.0, 3.0},
		{3.0, 99999.99},
		{3.0, -1.1},
		{3.0, -99999.99},
		{3.0, 0.0},
		{3.0, 1.1},
		{3.0, 3.0},
		{3.0, 3.0},
		{3.0, 3.0},
		{3.0, 99999.99},
		{3.0, -1.1},
		{3.0, -99999.99},
		{99999.99, 0.0},
		{99999.99, 1.1},
		{99999.99, 3.0},
		{99999.99, 3.0},
		{99999.99, 3.0},
		{99999.99, 99999.99},
		{99999.99, -1.1},
		{99999.99, -99999.99},
		{-1.1, 0.0},
		{-1.1, 1.1},
		{-1.1, 3.0},
		{-1.1, 3.0},
		{-1.1, 3.0},
		{-1.1, 99999.99},
		{-1.1, -1.1},
		{-1.1, -99999.99},
		{-99999.99, 0.0},
		{-99999.99, 1.1},
		{-99999.99, 3.0},
		{-99999.99, 3.0},
		{-99999.99, 3.0},
		{-99999.99, 99999.99},
		{-99999.99, -1.1},
		{-99999.99, -99999.99},
	}
	for _, cse := range cses {
		t.Run("Lss", func(t *testing.T) {
			a := cse.a
			b := cse.b
			actual := flt.Lss(a, b)
			expected := bol.Bol(a < b)
			tst.BolEql(t, expected, actual)
		})
	}
}
func TestFltFltGtrPkgFn(t *testing.T) {
	cses := []struct {
		a flt.Flt
		b flt.Flt
	}{
		{0.0, 0.0},
		{0.0, 1.1},
		{0.0, 3.0},
		{0.0, 3.0},
		{0.0, 3.0},
		{0.0, 99999.99},
		{0.0, -1.1},
		{0.0, -99999.99},
		{1.1, 0.0},
		{1.1, 1.1},
		{1.1, 3.0},
		{1.1, 3.0},
		{1.1, 3.0},
		{1.1, 99999.99},
		{1.1, -1.1},
		{1.1, -99999.99},
		{3.0, 0.0},
		{3.0, 1.1},
		{3.0, 3.0},
		{3.0, 3.0},
		{3.0, 3.0},
		{3.0, 99999.99},
		{3.0, -1.1},
		{3.0, -99999.99},
		{3.0, 0.0},
		{3.0, 1.1},
		{3.0, 3.0},
		{3.0, 3.0},
		{3.0, 3.0},
		{3.0, 99999.99},
		{3.0, -1.1},
		{3.0, -99999.99},
		{3.0, 0.0},
		{3.0, 1.1},
		{3.0, 3.0},
		{3.0, 3.0},
		{3.0, 3.0},
		{3.0, 99999.99},
		{3.0, -1.1},
		{3.0, -99999.99},
		{99999.99, 0.0},
		{99999.99, 1.1},
		{99999.99, 3.0},
		{99999.99, 3.0},
		{99999.99, 3.0},
		{99999.99, 99999.99},
		{99999.99, -1.1},
		{99999.99, -99999.99},
		{-1.1, 0.0},
		{-1.1, 1.1},
		{-1.1, 3.0},
		{-1.1, 3.0},
		{-1.1, 3.0},
		{-1.1, 99999.99},
		{-1.1, -1.1},
		{-1.1, -99999.99},
		{-99999.99, 0.0},
		{-99999.99, 1.1},
		{-99999.99, 3.0},
		{-99999.99, 3.0},
		{-99999.99, 3.0},
		{-99999.99, 99999.99},
		{-99999.99, -1.1},
		{-99999.99, -99999.99},
	}
	for _, cse := range cses {
		t.Run("Gtr", func(t *testing.T) {
			a := cse.a
			b := cse.b
			actual := flt.Gtr(a, b)
			expected := bol.Bol(a > b)
			tst.BolEql(t, expected, actual)
		})
	}
}
func TestFltFltEqlTypFn(t *testing.T) {
	cses := []struct {
		x flt.Flt
		a flt.Flt
	}{
		{flt.Flt(0.0), flt.Flt(1.1)},
		{flt.Flt(1.1), flt.Flt(1.1)},
		{flt.Flt(3.0), flt.Flt(1.1)},
		{flt.Flt(3.0), flt.Flt(1.1)},
		{flt.Flt(3.0), flt.Flt(1.1)},
		{flt.Flt(99999.99), flt.Flt(1.1)},
		{flt.Flt(-1.1), flt.Flt(1.1)},
		{flt.Flt(-99999.99), flt.Flt(1.1)},
	}
	for _, cse := range cses {
		t.Run("Eql", func(t *testing.T) {
			x := cse.x
			a := cse.a
			actual := x.Eql(a)
			expected := bol.Bol(cse.x == cse.a || cse.x.IsNaN() && cse.a.IsNaN())
			tst.BolEql(t, expected, actual)
		})
	}
}
func TestFltFltNeqTypFn(t *testing.T) {
	cses := []struct {
		x flt.Flt
		a flt.Flt
	}{
		{flt.Flt(0.0), flt.Flt(1.1)},
		{flt.Flt(1.1), flt.Flt(1.1)},
		{flt.Flt(3.0), flt.Flt(1.1)},
		{flt.Flt(3.0), flt.Flt(1.1)},
		{flt.Flt(3.0), flt.Flt(1.1)},
		{flt.Flt(99999.99), flt.Flt(1.1)},
		{flt.Flt(-1.1), flt.Flt(1.1)},
		{flt.Flt(-99999.99), flt.Flt(1.1)},
	}
	for _, cse := range cses {
		t.Run("Neq", func(t *testing.T) {
			x := cse.x
			a := cse.a
			actual := x.Neq(a)
			expected := bol.Bol(cse.x != cse.a)
			tst.BolEql(t, expected, actual)
		})
	}
}
func TestFltFltTrncTypFn(t *testing.T) {
	cses := []struct {
		x         flt.Flt
		precision unt.Unt
	}{
		{flt.Flt(0.0), unt.Unt(1)},
		{flt.Flt(1.1), unt.Unt(1)},
		{flt.Flt(3.0), unt.Unt(1)},
		{flt.Flt(3.0), unt.Unt(1)},
		{flt.Flt(3.0), unt.Unt(1)},
		{flt.Flt(99999.99), unt.Unt(1)},
		{flt.Flt(-1.1), unt.Unt(1)},
		{flt.Flt(-99999.99), unt.Unt(1)},
	}
	for _, cse := range cses {
		t.Run("Trnc", func(t *testing.T) {
			x := cse.x
			precision := cse.precision
			actual := x.Trnc(precision)
			expected := cse.x
			s := x.String()
			idx := strings.Index(s, ".")
			if len(s)-idx-1 > int(precision) {
				v, _ := strconv.ParseFloat(s[:idx+1+int(precision)], 32)
				expected = flt.Flt(v)
			}
			tst.FltEql(t, expected, actual)
		})
	}
}
func TestFltFltIsNaNTypFn(t *testing.T) {
	cses := []struct {
		x flt.Flt
	}{
		{flt.Flt(0.0)},
		{flt.Flt(1.1)},
		{flt.Flt(3.0)},
		{flt.Flt(3.0)},
		{flt.Flt(3.0)},
		{flt.Flt(99999.99)},
		{flt.Flt(-1.1)},
		{flt.Flt(-99999.99)},
	}
	for _, cse := range cses {
		t.Run("IsNaN", func(t *testing.T) {
			x := cse.x
			actual := x.IsNaN()
			expected := bol.Bol(cse.x != cse.x)
			tst.BolEql(t, expected, actual)
		})
	}
}
func TestFltFltIsInfPosTypFn(t *testing.T) {
	cses := []struct {
		x flt.Flt
	}{
		{flt.Flt(0.0)},
		{flt.Flt(1.1)},
		{flt.Flt(3.0)},
		{flt.Flt(3.0)},
		{flt.Flt(3.0)},
		{flt.Flt(99999.99)},
		{flt.Flt(-1.1)},
		{flt.Flt(-99999.99)},
	}
	for _, cse := range cses {
		t.Run("IsInfPos", func(t *testing.T) {
			x := cse.x
			actual := x.IsInfPos()
			expected := bol.Bol(cse.x > flt.Max)
			tst.BolEql(t, expected, actual)
		})
	}
}
func TestFltFltIsInfNegTypFn(t *testing.T) {
	cses := []struct {
		x flt.Flt
	}{
		{flt.Flt(0.0)},
		{flt.Flt(1.1)},
		{flt.Flt(3.0)},
		{flt.Flt(3.0)},
		{flt.Flt(3.0)},
		{flt.Flt(99999.99)},
		{flt.Flt(-1.1)},
		{flt.Flt(-99999.99)},
	}
	for _, cse := range cses {
		t.Run("IsInfNeg", func(t *testing.T) {
			x := cse.x
			actual := x.IsInfNeg()
			expected := bol.Bol(cse.x < flt.Min)
			tst.BolEql(t, expected, actual)
		})
	}
}
func TestFltFltIsValidTypFn(t *testing.T) {
	cses := []struct {
		x flt.Flt
	}{
		{flt.Flt(0.0)},
		{flt.Flt(1.1)},
		{flt.Flt(3.0)},
		{flt.Flt(3.0)},
		{flt.Flt(3.0)},
		{flt.Flt(99999.99)},
		{flt.Flt(-1.1)},
		{flt.Flt(-99999.99)},
	}
	for _, cse := range cses {
		t.Run("IsValid", func(t *testing.T) {
			x := cse.x
			actual := x.IsValid()
			expected := bol.Bol(cse.x == cse.x && cse.x >= flt.Min && cse.x <= flt.Max)
			tst.BolEql(t, expected, actual)
		})
	}
}
func TestFltFltLssTypFn(t *testing.T) {
	cses := []struct {
		x flt.Flt
		a flt.Flt
	}{
		{flt.Flt(0.0), flt.Flt(1.1)},
		{flt.Flt(1.1), flt.Flt(1.1)},
		{flt.Flt(3.0), flt.Flt(1.1)},
		{flt.Flt(3.0), flt.Flt(1.1)},
		{flt.Flt(3.0), flt.Flt(1.1)},
		{flt.Flt(99999.99), flt.Flt(1.1)},
		{flt.Flt(-1.1), flt.Flt(1.1)},
		{flt.Flt(-99999.99), flt.Flt(1.1)},
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
func TestFltFltGtrTypFn(t *testing.T) {
	cses := []struct {
		x flt.Flt
		a flt.Flt
	}{
		{flt.Flt(0.0), flt.Flt(1.1)},
		{flt.Flt(1.1), flt.Flt(1.1)},
		{flt.Flt(3.0), flt.Flt(1.1)},
		{flt.Flt(3.0), flt.Flt(1.1)},
		{flt.Flt(3.0), flt.Flt(1.1)},
		{flt.Flt(99999.99), flt.Flt(1.1)},
		{flt.Flt(-1.1), flt.Flt(1.1)},
		{flt.Flt(-99999.99), flt.Flt(1.1)},
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
func TestFltFltLeqTypFn(t *testing.T) {
	cses := []struct {
		x flt.Flt
		a flt.Flt
	}{
		{flt.Flt(0.0), flt.Flt(1.1)},
		{flt.Flt(1.1), flt.Flt(1.1)},
		{flt.Flt(3.0), flt.Flt(1.1)},
		{flt.Flt(3.0), flt.Flt(1.1)},
		{flt.Flt(3.0), flt.Flt(1.1)},
		{flt.Flt(99999.99), flt.Flt(1.1)},
		{flt.Flt(-1.1), flt.Flt(1.1)},
		{flt.Flt(-99999.99), flt.Flt(1.1)},
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
func TestFltFltGeqTypFn(t *testing.T) {
	cses := []struct {
		x flt.Flt
		a flt.Flt
	}{
		{flt.Flt(0.0), flt.Flt(1.1)},
		{flt.Flt(1.1), flt.Flt(1.1)},
		{flt.Flt(3.0), flt.Flt(1.1)},
		{flt.Flt(3.0), flt.Flt(1.1)},
		{flt.Flt(3.0), flt.Flt(1.1)},
		{flt.Flt(99999.99), flt.Flt(1.1)},
		{flt.Flt(-1.1), flt.Flt(1.1)},
		{flt.Flt(-99999.99), flt.Flt(1.1)},
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
func TestFltFltPosTypFn(t *testing.T) {
	cses := []struct {
		x flt.Flt
	}{
		{flt.Flt(0.0)},
		{flt.Flt(1.1)},
		{flt.Flt(3.0)},
		{flt.Flt(3.0)},
		{flt.Flt(3.0)},
		{flt.Flt(99999.99)},
		{flt.Flt(-1.1)},
		{flt.Flt(-99999.99)},
	}
	for _, cse := range cses {
		t.Run("Pos", func(t *testing.T) {
			x := cse.x
			actual := x.Pos()
			expected := cse.x
			if expected < 0 {
				expected = -expected
			}
			tst.FltEql(t, expected, actual)
		})
	}
}
func TestFltFltNegTypFn(t *testing.T) {
	cses := []struct {
		x flt.Flt
	}{
		{flt.Flt(0.0)},
		{flt.Flt(1.1)},
		{flt.Flt(3.0)},
		{flt.Flt(3.0)},
		{flt.Flt(3.0)},
		{flt.Flt(99999.99)},
		{flt.Flt(-1.1)},
		{flt.Flt(-99999.99)},
	}
	for _, cse := range cses {
		t.Run("Neg", func(t *testing.T) {
			x := cse.x
			actual := x.Neg()
			expected := cse.x
			if expected > 0 {
				expected = -expected
			}
			tst.FltEql(t, expected, actual)
		})
	}
}
func TestFltFltInvTypFn(t *testing.T) {
	cses := []struct {
		x flt.Flt
	}{
		{flt.Flt(0.0)},
		{flt.Flt(1.1)},
		{flt.Flt(3.0)},
		{flt.Flt(3.0)},
		{flt.Flt(3.0)},
		{flt.Flt(99999.99)},
		{flt.Flt(-1.1)},
		{flt.Flt(-99999.99)},
	}
	for _, cse := range cses {
		t.Run("Inv", func(t *testing.T) {
			x := cse.x
			actual := x.Inv()
			expected := -cse.x
			tst.FltEql(t, expected, actual)
		})
	}
}
func TestFltFltAddTypFn(t *testing.T) {
	cses := []struct {
		x flt.Flt
		a flt.Flt
	}{
		{flt.Flt(0.0), flt.Flt(1.1)},
		{flt.Flt(1.1), flt.Flt(1.1)},
		{flt.Flt(3.0), flt.Flt(1.1)},
		{flt.Flt(3.0), flt.Flt(1.1)},
		{flt.Flt(3.0), flt.Flt(1.1)},
		{flt.Flt(99999.99), flt.Flt(1.1)},
		{flt.Flt(-1.1), flt.Flt(1.1)},
		{flt.Flt(-99999.99), flt.Flt(1.1)},
	}
	for _, cse := range cses {
		t.Run("Add", func(t *testing.T) {
			x := cse.x
			a := cse.a
			actual := x.Add(a)
			expected := x + a
			tst.FltEql(t, expected, actual)
		})
	}
}
func TestFltFltSubTypFn(t *testing.T) {
	cses := []struct {
		x flt.Flt
		a flt.Flt
	}{
		{flt.Flt(0.0), flt.Flt(1.1)},
		{flt.Flt(1.1), flt.Flt(1.1)},
		{flt.Flt(3.0), flt.Flt(1.1)},
		{flt.Flt(3.0), flt.Flt(1.1)},
		{flt.Flt(3.0), flt.Flt(1.1)},
		{flt.Flt(99999.99), flt.Flt(1.1)},
		{flt.Flt(-1.1), flt.Flt(1.1)},
		{flt.Flt(-99999.99), flt.Flt(1.1)},
	}
	for _, cse := range cses {
		t.Run("Sub", func(t *testing.T) {
			x := cse.x
			a := cse.a
			actual := x.Sub(a)
			expected := x - a
			tst.FltEql(t, expected, actual)
		})
	}
}
func TestFltFltMulTypFn(t *testing.T) {
	cses := []struct {
		x flt.Flt
		a flt.Flt
	}{
		{flt.Flt(0.0), flt.Flt(1.1)},
		{flt.Flt(1.1), flt.Flt(1.1)},
		{flt.Flt(3.0), flt.Flt(1.1)},
		{flt.Flt(3.0), flt.Flt(1.1)},
		{flt.Flt(3.0), flt.Flt(1.1)},
		{flt.Flt(99999.99), flt.Flt(1.1)},
		{flt.Flt(-1.1), flt.Flt(1.1)},
		{flt.Flt(-99999.99), flt.Flt(1.1)},
	}
	for _, cse := range cses {
		t.Run("Mul", func(t *testing.T) {
			x := cse.x
			a := cse.a
			actual := x.Mul(a)
			expected := x * a
			tst.FltEql(t, expected, actual)
		})
	}
}
func TestFltFltDivTypFn(t *testing.T) {
	cses := []struct {
		x flt.Flt
		a flt.Flt
	}{
		{flt.Flt(0.0), flt.Flt(1.1)},
		{flt.Flt(1.1), flt.Flt(1.1)},
		{flt.Flt(3.0), flt.Flt(1.1)},
		{flt.Flt(3.0), flt.Flt(1.1)},
		{flt.Flt(3.0), flt.Flt(1.1)},
		{flt.Flt(99999.99), flt.Flt(1.1)},
		{flt.Flt(-1.1), flt.Flt(1.1)},
		{flt.Flt(-99999.99), flt.Flt(1.1)},
	}
	for _, cse := range cses {
		t.Run("Div", func(t *testing.T) {
			x := cse.x
			a := cse.a
			actual := x.Div(a)
			var expected flt.Flt
			if a != 0 {
				expected = x / a
			}
			tst.FltEql(t, expected, actual)
		})
	}
}
func TestFltFltRemTypFn(t *testing.T) {
	cses := []struct {
		x flt.Flt
		a flt.Flt
	}{
		{flt.Flt(0.0), flt.Flt(1.1)},
		{flt.Flt(1.1), flt.Flt(1.1)},
		{flt.Flt(3.0), flt.Flt(1.1)},
		{flt.Flt(3.0), flt.Flt(1.1)},
		{flt.Flt(3.0), flt.Flt(1.1)},
		{flt.Flt(99999.99), flt.Flt(1.1)},
		{flt.Flt(-1.1), flt.Flt(1.1)},
		{flt.Flt(-99999.99), flt.Flt(1.1)},
	}
	for _, cse := range cses {
		t.Run("Rem", func(t *testing.T) {
			x := cse.x
			a := cse.a
			actual := x.Rem(a)
			var expected flt.Flt
			if a != 0 {
				expected = flt.Flt(math.Remainder(float64(x), float64(a)))
			}
			tst.FltEql(t, expected, actual)
		})
	}
}
func TestFltFltPowTypFn(t *testing.T) {
	cses := []struct {
		x flt.Flt
		a flt.Flt
	}{
		{flt.Flt(0.0), flt.Flt(1.1)},
		{flt.Flt(1.1), flt.Flt(1.1)},
		{flt.Flt(3.0), flt.Flt(1.1)},
		{flt.Flt(3.0), flt.Flt(1.1)},
		{flt.Flt(3.0), flt.Flt(1.1)},
		{flt.Flt(99999.99), flt.Flt(1.1)},
		{flt.Flt(-1.1), flt.Flt(1.1)},
		{flt.Flt(-99999.99), flt.Flt(1.1)},
	}
	for _, cse := range cses {
		t.Run("Pow", func(t *testing.T) {
			x := cse.x
			a := cse.a
			actual := x.Pow(a)
			expected := flt.Flt(math.Pow(float64(x), float64(a)))
			tst.FltEql(t, expected, actual)
		})
	}
}
func TestFltFltSqrTypFn(t *testing.T) {
	cses := []struct {
		x flt.Flt
	}{
		{flt.Flt(0.0)},
		{flt.Flt(1.1)},
		{flt.Flt(3.0)},
		{flt.Flt(3.0)},
		{flt.Flt(3.0)},
		{flt.Flt(99999.99)},
		{flt.Flt(-1.1)},
		{flt.Flt(-99999.99)},
	}
	for _, cse := range cses {
		t.Run("Sqr", func(t *testing.T) {
			x := cse.x
			actual := x.Sqr()
			expected := x * x
			tst.FltEql(t, expected, actual)
		})
	}
}
func TestFltFltSqrtTypFn(t *testing.T) {
	cses := []struct {
		x flt.Flt
	}{
		{flt.Flt(0.0)},
		{flt.Flt(1.1)},
		{flt.Flt(3.0)},
		{flt.Flt(3.0)},
		{flt.Flt(3.0)},
		{flt.Flt(99999.99)},
		{flt.Flt(-1.1)},
		{flt.Flt(-99999.99)},
	}
	for _, cse := range cses {
		t.Run("Sqrt", func(t *testing.T) {
			x := cse.x
			actual := x.Sqrt()
			var expected flt.Flt
			if x > 0 {
				expected = flt.Flt(math.Sqrt(float64(x)))
			}
			tst.FltEql(t, expected, actual)
		})
	}
}
func TestFltFltMinTypFn(t *testing.T) {
	cses := []struct {
		x flt.Flt
		a flt.Flt
	}{
		{flt.Flt(0.0), flt.Flt(1.1)},
		{flt.Flt(1.1), flt.Flt(1.1)},
		{flt.Flt(3.0), flt.Flt(1.1)},
		{flt.Flt(3.0), flt.Flt(1.1)},
		{flt.Flt(3.0), flt.Flt(1.1)},
		{flt.Flt(99999.99), flt.Flt(1.1)},
		{flt.Flt(-1.1), flt.Flt(1.1)},
		{flt.Flt(-99999.99), flt.Flt(1.1)},
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
			tst.FltEql(t, expected, actual)
		})
	}
}
func TestFltFltMaxTypFn(t *testing.T) {
	cses := []struct {
		x flt.Flt
		a flt.Flt
	}{
		{flt.Flt(0.0), flt.Flt(1.1)},
		{flt.Flt(1.1), flt.Flt(1.1)},
		{flt.Flt(3.0), flt.Flt(1.1)},
		{flt.Flt(3.0), flt.Flt(1.1)},
		{flt.Flt(3.0), flt.Flt(1.1)},
		{flt.Flt(99999.99), flt.Flt(1.1)},
		{flt.Flt(-1.1), flt.Flt(1.1)},
		{flt.Flt(-99999.99), flt.Flt(1.1)},
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
			tst.FltEql(t, expected, actual)
		})
	}
}
func TestFltFltMinMaxTypFn(t *testing.T) {
	cses := []struct {
		x flt.Flt
		a flt.Flt
	}{
		{flt.Flt(0.0), flt.Flt(1.1)},
		{flt.Flt(1.1), flt.Flt(1.1)},
		{flt.Flt(3.0), flt.Flt(1.1)},
		{flt.Flt(3.0), flt.Flt(1.1)},
		{flt.Flt(3.0), flt.Flt(1.1)},
		{flt.Flt(99999.99), flt.Flt(1.1)},
		{flt.Flt(-1.1), flt.Flt(1.1)},
		{flt.Flt(-99999.99), flt.Flt(1.1)},
	}
	for _, cse := range cses {
		t.Run("MinMax", func(t *testing.T) {
			x := cse.x
			a := cse.a
			aMin, aMax := x.MinMax(a)
			var eMin, eMax flt.Flt
			if x < a {
				eMin, eMax = x, a
			} else {
				eMin, eMax = a, x
			}
			tst.FltEql(t, eMin, aMin, "Min")
			tst.FltEql(t, eMax, aMax, "Max")
		})
	}
}
func TestFltFltMidTypFn(t *testing.T) {
	cses := []struct {
		x flt.Flt
		a flt.Flt
	}{
		{flt.Flt(0.0), flt.Flt(1.1)},
		{flt.Flt(1.1), flt.Flt(1.1)},
		{flt.Flt(3.0), flt.Flt(1.1)},
		{flt.Flt(3.0), flt.Flt(1.1)},
		{flt.Flt(3.0), flt.Flt(1.1)},
		{flt.Flt(99999.99), flt.Flt(1.1)},
		{flt.Flt(-1.1), flt.Flt(1.1)},
		{flt.Flt(-99999.99), flt.Flt(1.1)},
	}
	for _, cse := range cses {
		t.Run("Mid", func(t *testing.T) {
			x := cse.x
			a := cse.a
			actual := x.Mid(a)
			expected := (x - a) / 2
			tst.FltEql(t, expected, actual)
		})
	}
}
func TestFltFltAvgTypFn(t *testing.T) {
	cses := []struct {
		x flt.Flt
		a flt.Flt
	}{
		{flt.Flt(0.0), flt.Flt(1.1)},
		{flt.Flt(1.1), flt.Flt(1.1)},
		{flt.Flt(3.0), flt.Flt(1.1)},
		{flt.Flt(3.0), flt.Flt(1.1)},
		{flt.Flt(3.0), flt.Flt(1.1)},
		{flt.Flt(99999.99), flt.Flt(1.1)},
		{flt.Flt(-1.1), flt.Flt(1.1)},
		{flt.Flt(-99999.99), flt.Flt(1.1)},
	}
	for _, cse := range cses {
		t.Run("Avg", func(t *testing.T) {
			x := cse.x
			a := cse.a
			actual := x.Avg(a)
			expected := (x + a) / 2
			tst.FltEql(t, expected, actual)
		})
	}
}
func TestFltFltAvgGeoTypFn(t *testing.T) {
	cses := []struct {
		x flt.Flt
		a flt.Flt
	}{
		{flt.Flt(0.0), flt.Flt(1.1)},
		{flt.Flt(1.1), flt.Flt(1.1)},
		{flt.Flt(3.0), flt.Flt(1.1)},
		{flt.Flt(3.0), flt.Flt(1.1)},
		{flt.Flt(3.0), flt.Flt(1.1)},
		{flt.Flt(99999.99), flt.Flt(1.1)},
		{flt.Flt(-1.1), flt.Flt(1.1)},
		{flt.Flt(-99999.99), flt.Flt(1.1)},
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
			tst.FltEql(t, expected, actual)
		})
	}
}
func TestFltFltSelEqlTypFn(t *testing.T) {
	cses := []struct {
		x flt.Flt
		a flt.Flt
	}{
		{flt.Flt(0.0), flt.Flt(1.1)},
		{flt.Flt(1.1), flt.Flt(1.1)},
		{flt.Flt(3.0), flt.Flt(1.1)},
		{flt.Flt(3.0), flt.Flt(1.1)},
		{flt.Flt(3.0), flt.Flt(1.1)},
		{flt.Flt(99999.99), flt.Flt(1.1)},
		{flt.Flt(-1.1), flt.Flt(1.1)},
		{flt.Flt(-99999.99), flt.Flt(1.1)},
	}
	for _, cse := range cses {
		t.Run("SelEql", func(t *testing.T) {
			x := cse.x
			a := cse.a
			actual := x.SelEql(a)
			var expected flt.Flt
			if x == a {
				expected = x
			}
			tst.FltEql(t, expected, actual)
		})
	}
}
func TestFltFltSelNeqTypFn(t *testing.T) {
	cses := []struct {
		x flt.Flt
		a flt.Flt
	}{
		{flt.Flt(0.0), flt.Flt(1.1)},
		{flt.Flt(1.1), flt.Flt(1.1)},
		{flt.Flt(3.0), flt.Flt(1.1)},
		{flt.Flt(3.0), flt.Flt(1.1)},
		{flt.Flt(3.0), flt.Flt(1.1)},
		{flt.Flt(99999.99), flt.Flt(1.1)},
		{flt.Flt(-1.1), flt.Flt(1.1)},
		{flt.Flt(-99999.99), flt.Flt(1.1)},
	}
	for _, cse := range cses {
		t.Run("SelNeq", func(t *testing.T) {
			x := cse.x
			a := cse.a
			actual := x.SelNeq(a)
			var expected flt.Flt
			if x != a {
				expected = x
			}
			tst.FltEql(t, expected, actual)
		})
	}
}
func TestFltFltSelLssTypFn(t *testing.T) {
	cses := []struct {
		x flt.Flt
		a flt.Flt
	}{
		{flt.Flt(0.0), flt.Flt(1.1)},
		{flt.Flt(1.1), flt.Flt(1.1)},
		{flt.Flt(3.0), flt.Flt(1.1)},
		{flt.Flt(3.0), flt.Flt(1.1)},
		{flt.Flt(3.0), flt.Flt(1.1)},
		{flt.Flt(99999.99), flt.Flt(1.1)},
		{flt.Flt(-1.1), flt.Flt(1.1)},
		{flt.Flt(-99999.99), flt.Flt(1.1)},
	}
	for _, cse := range cses {
		t.Run("SelLss", func(t *testing.T) {
			x := cse.x
			a := cse.a
			actual := x.SelLss(a)
			var expected flt.Flt
			if x < a {
				expected = x
			}
			tst.FltEql(t, expected, actual)
		})
	}
}
func TestFltFltSelGtrTypFn(t *testing.T) {
	cses := []struct {
		x flt.Flt
		a flt.Flt
	}{
		{flt.Flt(0.0), flt.Flt(1.1)},
		{flt.Flt(1.1), flt.Flt(1.1)},
		{flt.Flt(3.0), flt.Flt(1.1)},
		{flt.Flt(3.0), flt.Flt(1.1)},
		{flt.Flt(3.0), flt.Flt(1.1)},
		{flt.Flt(99999.99), flt.Flt(1.1)},
		{flt.Flt(-1.1), flt.Flt(1.1)},
		{flt.Flt(-99999.99), flt.Flt(1.1)},
	}
	for _, cse := range cses {
		t.Run("SelGtr", func(t *testing.T) {
			x := cse.x
			a := cse.a
			actual := x.SelGtr(a)
			var expected flt.Flt
			if x > a {
				expected = x
			}
			tst.FltEql(t, expected, actual)
		})
	}
}
func TestFltFltSelLeqTypFn(t *testing.T) {
	cses := []struct {
		x flt.Flt
		a flt.Flt
	}{
		{flt.Flt(0.0), flt.Flt(1.1)},
		{flt.Flt(1.1), flt.Flt(1.1)},
		{flt.Flt(3.0), flt.Flt(1.1)},
		{flt.Flt(3.0), flt.Flt(1.1)},
		{flt.Flt(3.0), flt.Flt(1.1)},
		{flt.Flt(99999.99), flt.Flt(1.1)},
		{flt.Flt(-1.1), flt.Flt(1.1)},
		{flt.Flt(-99999.99), flt.Flt(1.1)},
	}
	for _, cse := range cses {
		t.Run("SelLeq", func(t *testing.T) {
			x := cse.x
			a := cse.a
			actual := x.SelLeq(a)
			var expected flt.Flt
			if x <= a {
				expected = x
			}
			tst.FltEql(t, expected, actual)
		})
	}
}
func TestFltFltSelGeqTypFn(t *testing.T) {
	cses := []struct {
		x flt.Flt
		a flt.Flt
	}{
		{flt.Flt(0.0), flt.Flt(1.1)},
		{flt.Flt(1.1), flt.Flt(1.1)},
		{flt.Flt(3.0), flt.Flt(1.1)},
		{flt.Flt(3.0), flt.Flt(1.1)},
		{flt.Flt(3.0), flt.Flt(1.1)},
		{flt.Flt(99999.99), flt.Flt(1.1)},
		{flt.Flt(-1.1), flt.Flt(1.1)},
		{flt.Flt(-99999.99), flt.Flt(1.1)},
	}
	for _, cse := range cses {
		t.Run("SelGeq", func(t *testing.T) {
			x := cse.x
			a := cse.a
			actual := x.SelGeq(a)
			var expected flt.Flt
			if x >= a {
				expected = x
			}
			tst.FltEql(t, expected, actual)
		})
	}
}
func TestFltFltByt(t *testing.T) {
	cses := []struct {
		e flt.Flt
	}{
		{0.0},
		{1.1},
		{3.0},
		{3.0},
		{3.0},
		{99999.99},
		{-1.1},
		{-99999.99},
	}
	for _, cse := range cses {
		t.Run("Byt", func(t *testing.T) {
			b := &bytes.Buffer{}
			cse.e.BytWrt(b)
			var a flt.Flt
			a.BytRed(b.Bytes())
			tst.FltEql(t, cse.e, a)
		})
	}
}
