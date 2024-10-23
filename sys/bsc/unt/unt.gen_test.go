package unt_test

import (
	"bytes"
	"fmt"
	"math"
	"sys/bsc/bol"
	"sys/bsc/unt"
	"sys/tst"
	"testing"
)

func TestUntUntCnst(t *testing.T) {
	cses := []struct {
		idn string
		e   unt.Unt
		a   unt.Unt
	}{
		{"Zero", unt.Unt(0), unt.Zero},
		{"One", unt.Unt(1), unt.One},
		{"Min", unt.Unt(0), unt.Min},
		{"Max", unt.Unt(1<<32 - 1), unt.Max},
		{"MinSegLen", unt.Unt(64), unt.MinSegLen},
	}
	for _, cse := range cses {
		t.Run(fmt.Sprintf("%q", cse.idn), func(t *testing.T) {
			tst.UntEql(t, cse.e, cse.a)
		})
	}
}
func TestUntUntEqlPkgFn(t *testing.T) {
	cses := []struct {
		a unt.Unt
		b unt.Unt
	}{
		{0, 0},
		{0, 1},
		{0, 1000},
		{0, 10},
		{1, 0},
		{1, 1},
		{1, 1000},
		{1, 10},
		{1000, 0},
		{1000, 1},
		{1000, 1000},
		{1000, 10},
		{10, 0},
		{10, 1},
		{10, 1000},
		{10, 10},
	}
	for _, cse := range cses {
		t.Run("Eql", func(t *testing.T) {
			a := cse.a
			b := cse.b
			actual := unt.Eql(a, b)
			expected := bol.Bol(a == b)
			tst.BolEql(t, expected, actual)
		})
	}
}
func TestUntUntLssPkgFn(t *testing.T) {
	cses := []struct {
		a unt.Unt
		b unt.Unt
	}{
		{0, 0},
		{0, 1},
		{0, 1000},
		{0, 10},
		{1, 0},
		{1, 1},
		{1, 1000},
		{1, 10},
		{1000, 0},
		{1000, 1},
		{1000, 1000},
		{1000, 10},
		{10, 0},
		{10, 1},
		{10, 1000},
		{10, 10},
	}
	for _, cse := range cses {
		t.Run("Lss", func(t *testing.T) {
			a := cse.a
			b := cse.b
			actual := unt.Lss(a, b)
			expected := bol.Bol(a < b)
			tst.BolEql(t, expected, actual)
		})
	}
}
func TestUntUntGtrPkgFn(t *testing.T) {
	cses := []struct {
		a unt.Unt
		b unt.Unt
	}{
		{0, 0},
		{0, 1},
		{0, 1000},
		{0, 10},
		{1, 0},
		{1, 1},
		{1, 1000},
		{1, 10},
		{1000, 0},
		{1000, 1},
		{1000, 1000},
		{1000, 10},
		{10, 0},
		{10, 1},
		{10, 1000},
		{10, 10},
	}
	for _, cse := range cses {
		t.Run("Gtr", func(t *testing.T) {
			a := cse.a
			b := cse.b
			actual := unt.Gtr(a, b)
			expected := bol.Bol(a > b)
			tst.BolEql(t, expected, actual)
		})
	}
}
func TestUntUntEqlTypFn(t *testing.T) {
	cses := []struct {
		x unt.Unt
		a unt.Unt
	}{
		{unt.Unt(0), unt.Unt(1)},
		{unt.Unt(1), unt.Unt(1)},
		{unt.Unt(1000), unt.Unt(1)},
		{unt.Unt(10), unt.Unt(1)},
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
func TestUntUntNeqTypFn(t *testing.T) {
	cses := []struct {
		x unt.Unt
		a unt.Unt
	}{
		{unt.Unt(0), unt.Unt(1)},
		{unt.Unt(1), unt.Unt(1)},
		{unt.Unt(1000), unt.Unt(1)},
		{unt.Unt(10), unt.Unt(1)},
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
func TestUntUntLssTypFn(t *testing.T) {
	cses := []struct {
		x unt.Unt
		a unt.Unt
	}{
		{unt.Unt(0), unt.Unt(1)},
		{unt.Unt(1), unt.Unt(1)},
		{unt.Unt(1000), unt.Unt(1)},
		{unt.Unt(10), unt.Unt(1)},
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
func TestUntUntGtrTypFn(t *testing.T) {
	cses := []struct {
		x unt.Unt
		a unt.Unt
	}{
		{unt.Unt(0), unt.Unt(1)},
		{unt.Unt(1), unt.Unt(1)},
		{unt.Unt(1000), unt.Unt(1)},
		{unt.Unt(10), unt.Unt(1)},
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
func TestUntUntLeqTypFn(t *testing.T) {
	cses := []struct {
		x unt.Unt
		a unt.Unt
	}{
		{unt.Unt(0), unt.Unt(1)},
		{unt.Unt(1), unt.Unt(1)},
		{unt.Unt(1000), unt.Unt(1)},
		{unt.Unt(10), unt.Unt(1)},
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
func TestUntUntGeqTypFn(t *testing.T) {
	cses := []struct {
		x unt.Unt
		a unt.Unt
	}{
		{unt.Unt(0), unt.Unt(1)},
		{unt.Unt(1), unt.Unt(1)},
		{unt.Unt(1000), unt.Unt(1)},
		{unt.Unt(10), unt.Unt(1)},
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
func TestUntUntAddTypFn(t *testing.T) {
	cses := []struct {
		x unt.Unt
		a unt.Unt
	}{
		{unt.Unt(0), unt.Unt(1)},
		{unt.Unt(1), unt.Unt(1)},
		{unt.Unt(1000), unt.Unt(1)},
		{unt.Unt(10), unt.Unt(1)},
	}
	for _, cse := range cses {
		t.Run("Add", func(t *testing.T) {
			x := cse.x
			a := cse.a
			actual := x.Add(a)
			expected := x + a
			tst.UntEql(t, expected, actual)
		})
	}
}
func TestUntUntSubTypFn(t *testing.T) {
	cses := []struct {
		x unt.Unt
		a unt.Unt
	}{
		{unt.Unt(0), unt.Unt(1)},
		{unt.Unt(1), unt.Unt(1)},
		{unt.Unt(1000), unt.Unt(1)},
		{unt.Unt(10), unt.Unt(1)},
	}
	for _, cse := range cses {
		t.Run("Sub", func(t *testing.T) {
			x := cse.x
			a := cse.a
			actual := x.Sub(a)
			expected := x - a
			tst.UntEql(t, expected, actual)
		})
	}
}
func TestUntUntMulTypFn(t *testing.T) {
	cses := []struct {
		x unt.Unt
		a unt.Unt
	}{
		{unt.Unt(0), unt.Unt(1)},
		{unt.Unt(1), unt.Unt(1)},
		{unt.Unt(1000), unt.Unt(1)},
		{unt.Unt(10), unt.Unt(1)},
	}
	for _, cse := range cses {
		t.Run("Mul", func(t *testing.T) {
			x := cse.x
			a := cse.a
			actual := x.Mul(a)
			expected := x * a
			tst.UntEql(t, expected, actual)
		})
	}
}
func TestUntUntDivTypFn(t *testing.T) {
	cses := []struct {
		x unt.Unt
		a unt.Unt
	}{
		{unt.Unt(0), unt.Unt(1)},
		{unt.Unt(1), unt.Unt(1)},
		{unt.Unt(1000), unt.Unt(1)},
		{unt.Unt(10), unt.Unt(1)},
	}
	for _, cse := range cses {
		t.Run("Div", func(t *testing.T) {
			x := cse.x
			a := cse.a
			actual := x.Div(a)
			var expected unt.Unt
			if a != 0 {
				expected = x / a
			}
			tst.UntEql(t, expected, actual)
		})
	}
}
func TestUntUntRemTypFn(t *testing.T) {
	cses := []struct {
		x unt.Unt
		a unt.Unt
	}{
		{unt.Unt(0), unt.Unt(1)},
		{unt.Unt(1), unt.Unt(1)},
		{unt.Unt(1000), unt.Unt(1)},
		{unt.Unt(10), unt.Unt(1)},
	}
	for _, cse := range cses {
		t.Run("Rem", func(t *testing.T) {
			x := cse.x
			a := cse.a
			actual := x.Rem(a)
			var expected unt.Unt
			if a != 0 {
				expected = unt.Unt(math.Remainder(float64(x), float64(a)))
			}
			tst.UntEql(t, expected, actual)
		})
	}
}
func TestUntUntPowTypFn(t *testing.T) {
	cses := []struct {
		x unt.Unt
		a unt.Unt
	}{
		{unt.Unt(0), unt.Unt(1)},
		{unt.Unt(1), unt.Unt(1)},
		{unt.Unt(1000), unt.Unt(1)},
		{unt.Unt(10), unt.Unt(1)},
	}
	for _, cse := range cses {
		t.Run("Pow", func(t *testing.T) {
			x := cse.x
			a := cse.a
			actual := x.Pow(a)
			expected := unt.Unt(math.Pow(float64(x), float64(a)))
			tst.UntEql(t, expected, actual)
		})
	}
}
func TestUntUntSqrTypFn(t *testing.T) {
	cses := []struct {
		x unt.Unt
	}{
		{unt.Unt(0)},
		{unt.Unt(1)},
		{unt.Unt(1000)},
		{unt.Unt(10)},
	}
	for _, cse := range cses {
		t.Run("Sqr", func(t *testing.T) {
			x := cse.x
			actual := x.Sqr()
			expected := x * x
			tst.UntEql(t, expected, actual)
		})
	}
}
func TestUntUntSqrtTypFn(t *testing.T) {
	cses := []struct {
		x unt.Unt
	}{
		{unt.Unt(0)},
		{unt.Unt(1)},
		{unt.Unt(1000)},
		{unt.Unt(10)},
	}
	for _, cse := range cses {
		t.Run("Sqrt", func(t *testing.T) {
			x := cse.x
			actual := x.Sqrt()
			var expected unt.Unt
			if x > 0 {
				expected = unt.Unt(math.Sqrt(float64(x)))
			}
			tst.UntEql(t, expected, actual)
		})
	}
}
func TestUntUntMinTypFn(t *testing.T) {
	cses := []struct {
		x unt.Unt
		a unt.Unt
	}{
		{unt.Unt(0), unt.Unt(1)},
		{unt.Unt(1), unt.Unt(1)},
		{unt.Unt(1000), unt.Unt(1)},
		{unt.Unt(10), unt.Unt(1)},
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
			tst.UntEql(t, expected, actual)
		})
	}
}
func TestUntUntMaxTypFn(t *testing.T) {
	cses := []struct {
		x unt.Unt
		a unt.Unt
	}{
		{unt.Unt(0), unt.Unt(1)},
		{unt.Unt(1), unt.Unt(1)},
		{unt.Unt(1000), unt.Unt(1)},
		{unt.Unt(10), unt.Unt(1)},
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
			tst.UntEql(t, expected, actual)
		})
	}
}
func TestUntUntMinMaxTypFn(t *testing.T) {
	cses := []struct {
		x unt.Unt
		a unt.Unt
	}{
		{unt.Unt(0), unt.Unt(1)},
		{unt.Unt(1), unt.Unt(1)},
		{unt.Unt(1000), unt.Unt(1)},
		{unt.Unt(10), unt.Unt(1)},
	}
	for _, cse := range cses {
		t.Run("MinMax", func(t *testing.T) {
			x := cse.x
			a := cse.a
			aMin, aMax := x.MinMax(a)
			var eMin, eMax unt.Unt
			if x < a {
				eMin, eMax = x, a
			} else {
				eMin, eMax = a, x
			}
			tst.UntEql(t, eMin, aMin, "Min")
			tst.UntEql(t, eMax, aMax, "Max")
		})
	}
}
func TestUntUntMidTypFn(t *testing.T) {
	cses := []struct {
		x unt.Unt
		a unt.Unt
	}{
		{unt.Unt(0), unt.Unt(1)},
		{unt.Unt(1), unt.Unt(1)},
		{unt.Unt(1000), unt.Unt(1)},
		{unt.Unt(10), unt.Unt(1)},
	}
	for _, cse := range cses {
		t.Run("Mid", func(t *testing.T) {
			x := cse.x
			a := cse.a
			actual := x.Mid(a)
			expected := (x - a) / 2
			tst.UntEql(t, expected, actual)
		})
	}
}
func TestUntUntAvgTypFn(t *testing.T) {
	cses := []struct {
		x unt.Unt
		a unt.Unt
	}{
		{unt.Unt(0), unt.Unt(1)},
		{unt.Unt(1), unt.Unt(1)},
		{unt.Unt(1000), unt.Unt(1)},
		{unt.Unt(10), unt.Unt(1)},
	}
	for _, cse := range cses {
		t.Run("Avg", func(t *testing.T) {
			x := cse.x
			a := cse.a
			actual := x.Avg(a)
			expected := (x + a) / 2
			tst.UntEql(t, expected, actual)
		})
	}
}
func TestUntUntAvgGeoTypFn(t *testing.T) {
	cses := []struct {
		x unt.Unt
		a unt.Unt
	}{
		{unt.Unt(0), unt.Unt(1)},
		{unt.Unt(1), unt.Unt(1)},
		{unt.Unt(1000), unt.Unt(1)},
		{unt.Unt(10), unt.Unt(1)},
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
			tst.UntEql(t, expected, actual)
		})
	}
}
func TestUntUntByt(t *testing.T) {
	cses := []struct {
		e unt.Unt
	}{
		{0},
		{1},
		{1000},
		{10},
	}
	for _, cse := range cses {
		t.Run("Byt", func(t *testing.T) {
			b := &bytes.Buffer{}
			cse.e.BytWrt(b)
			var a unt.Unt
			a.BytRed(b.Bytes())
			tst.UntEql(t, cse.e, a)
		})
	}
}
