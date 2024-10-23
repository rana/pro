package int_test

import (
	"bytes"
	"fmt"
	"math"
	"sys/bsc/bol"
	"sys/bsc/int"
	"sys/tst"
	"testing"
)

func TestIntIntCnst(t *testing.T) {
	cses := []struct {
		idn string
		e   int.Int
		a   int.Int
	}{
		{"Zero", int.Int(0), int.Zero},
		{"One", int.Int(1), int.One},
		{"NegOne", int.Int(-1), int.NegOne},
		{"Min", int.Int(-1 << 31), int.Min},
		{"Max", int.Int(1<<31 - 1), int.Max},
	}
	for _, cse := range cses {
		t.Run(fmt.Sprintf("%q", cse.idn), func(t *testing.T) {
			tst.IntEql(t, cse.e, cse.a)
		})
	}
}
func TestIntIntEqlPkgFn(t *testing.T) {
	cses := []struct {
		a int.Int
		b int.Int
	}{
		{0, 0},
		{0, 10},
		{0, 1000},
		{0, -10},
		{0, -1000},
		{10, 0},
		{10, 10},
		{10, 1000},
		{10, -10},
		{10, -1000},
		{1000, 0},
		{1000, 10},
		{1000, 1000},
		{1000, -10},
		{1000, -1000},
		{-10, 0},
		{-10, 10},
		{-10, 1000},
		{-10, -10},
		{-10, -1000},
		{-1000, 0},
		{-1000, 10},
		{-1000, 1000},
		{-1000, -10},
		{-1000, -1000},
	}
	for _, cse := range cses {
		t.Run("Eql", func(t *testing.T) {
			a := cse.a
			b := cse.b
			actual := int.Eql(a, b)
			expected := bol.Bol(a == b)
			tst.BolEql(t, expected, actual)
		})
	}
}
func TestIntIntLssPkgFn(t *testing.T) {
	cses := []struct {
		a int.Int
		b int.Int
	}{
		{0, 0},
		{0, 10},
		{0, 1000},
		{0, -10},
		{0, -1000},
		{10, 0},
		{10, 10},
		{10, 1000},
		{10, -10},
		{10, -1000},
		{1000, 0},
		{1000, 10},
		{1000, 1000},
		{1000, -10},
		{1000, -1000},
		{-10, 0},
		{-10, 10},
		{-10, 1000},
		{-10, -10},
		{-10, -1000},
		{-1000, 0},
		{-1000, 10},
		{-1000, 1000},
		{-1000, -10},
		{-1000, -1000},
	}
	for _, cse := range cses {
		t.Run("Lss", func(t *testing.T) {
			a := cse.a
			b := cse.b
			actual := int.Lss(a, b)
			expected := bol.Bol(a < b)
			tst.BolEql(t, expected, actual)
		})
	}
}
func TestIntIntGtrPkgFn(t *testing.T) {
	cses := []struct {
		a int.Int
		b int.Int
	}{
		{0, 0},
		{0, 10},
		{0, 1000},
		{0, -10},
		{0, -1000},
		{10, 0},
		{10, 10},
		{10, 1000},
		{10, -10},
		{10, -1000},
		{1000, 0},
		{1000, 10},
		{1000, 1000},
		{1000, -10},
		{1000, -1000},
		{-10, 0},
		{-10, 10},
		{-10, 1000},
		{-10, -10},
		{-10, -1000},
		{-1000, 0},
		{-1000, 10},
		{-1000, 1000},
		{-1000, -10},
		{-1000, -1000},
	}
	for _, cse := range cses {
		t.Run("Gtr", func(t *testing.T) {
			a := cse.a
			b := cse.b
			actual := int.Gtr(a, b)
			expected := bol.Bol(a > b)
			tst.BolEql(t, expected, actual)
		})
	}
}
func TestIntIntEqlTypFn(t *testing.T) {
	cses := []struct {
		x int.Int
		a int.Int
	}{
		{int.Int(0), int.Int(10)},
		{int.Int(10), int.Int(10)},
		{int.Int(1000), int.Int(10)},
		{int.Int(-10), int.Int(10)},
		{int.Int(-1000), int.Int(10)},
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
func TestIntIntNeqTypFn(t *testing.T) {
	cses := []struct {
		x int.Int
		a int.Int
	}{
		{int.Int(0), int.Int(10)},
		{int.Int(10), int.Int(10)},
		{int.Int(1000), int.Int(10)},
		{int.Int(-10), int.Int(10)},
		{int.Int(-1000), int.Int(10)},
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
func TestIntIntLssTypFn(t *testing.T) {
	cses := []struct {
		x int.Int
		a int.Int
	}{
		{int.Int(0), int.Int(10)},
		{int.Int(10), int.Int(10)},
		{int.Int(1000), int.Int(10)},
		{int.Int(-10), int.Int(10)},
		{int.Int(-1000), int.Int(10)},
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
func TestIntIntGtrTypFn(t *testing.T) {
	cses := []struct {
		x int.Int
		a int.Int
	}{
		{int.Int(0), int.Int(10)},
		{int.Int(10), int.Int(10)},
		{int.Int(1000), int.Int(10)},
		{int.Int(-10), int.Int(10)},
		{int.Int(-1000), int.Int(10)},
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
func TestIntIntLeqTypFn(t *testing.T) {
	cses := []struct {
		x int.Int
		a int.Int
	}{
		{int.Int(0), int.Int(10)},
		{int.Int(10), int.Int(10)},
		{int.Int(1000), int.Int(10)},
		{int.Int(-10), int.Int(10)},
		{int.Int(-1000), int.Int(10)},
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
func TestIntIntGeqTypFn(t *testing.T) {
	cses := []struct {
		x int.Int
		a int.Int
	}{
		{int.Int(0), int.Int(10)},
		{int.Int(10), int.Int(10)},
		{int.Int(1000), int.Int(10)},
		{int.Int(-10), int.Int(10)},
		{int.Int(-1000), int.Int(10)},
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
func TestIntIntPosTypFn(t *testing.T) {
	cses := []struct {
		x int.Int
	}{
		{int.Int(0)},
		{int.Int(10)},
		{int.Int(1000)},
		{int.Int(-10)},
		{int.Int(-1000)},
	}
	for _, cse := range cses {
		t.Run("Pos", func(t *testing.T) {
			x := cse.x
			actual := x.Pos()
			expected := cse.x
			if expected < 0 {
				expected = -expected
			}
			tst.IntEql(t, expected, actual)
		})
	}
}
func TestIntIntNegTypFn(t *testing.T) {
	cses := []struct {
		x int.Int
	}{
		{int.Int(0)},
		{int.Int(10)},
		{int.Int(1000)},
		{int.Int(-10)},
		{int.Int(-1000)},
	}
	for _, cse := range cses {
		t.Run("Neg", func(t *testing.T) {
			x := cse.x
			actual := x.Neg()
			expected := cse.x
			if expected > 0 {
				expected = -expected
			}
			tst.IntEql(t, expected, actual)
		})
	}
}
func TestIntIntInvTypFn(t *testing.T) {
	cses := []struct {
		x int.Int
	}{
		{int.Int(0)},
		{int.Int(10)},
		{int.Int(1000)},
		{int.Int(-10)},
		{int.Int(-1000)},
	}
	for _, cse := range cses {
		t.Run("Inv", func(t *testing.T) {
			x := cse.x
			actual := x.Inv()
			expected := -cse.x
			tst.IntEql(t, expected, actual)
		})
	}
}
func TestIntIntAddTypFn(t *testing.T) {
	cses := []struct {
		x int.Int
		a int.Int
	}{
		{int.Int(0), int.Int(10)},
		{int.Int(10), int.Int(10)},
		{int.Int(1000), int.Int(10)},
		{int.Int(-10), int.Int(10)},
		{int.Int(-1000), int.Int(10)},
	}
	for _, cse := range cses {
		t.Run("Add", func(t *testing.T) {
			x := cse.x
			a := cse.a
			actual := x.Add(a)
			expected := x + a
			tst.IntEql(t, expected, actual)
		})
	}
}
func TestIntIntSubTypFn(t *testing.T) {
	cses := []struct {
		x int.Int
		a int.Int
	}{
		{int.Int(0), int.Int(10)},
		{int.Int(10), int.Int(10)},
		{int.Int(1000), int.Int(10)},
		{int.Int(-10), int.Int(10)},
		{int.Int(-1000), int.Int(10)},
	}
	for _, cse := range cses {
		t.Run("Sub", func(t *testing.T) {
			x := cse.x
			a := cse.a
			actual := x.Sub(a)
			expected := x - a
			tst.IntEql(t, expected, actual)
		})
	}
}
func TestIntIntMulTypFn(t *testing.T) {
	cses := []struct {
		x int.Int
		a int.Int
	}{
		{int.Int(0), int.Int(10)},
		{int.Int(10), int.Int(10)},
		{int.Int(1000), int.Int(10)},
		{int.Int(-10), int.Int(10)},
		{int.Int(-1000), int.Int(10)},
	}
	for _, cse := range cses {
		t.Run("Mul", func(t *testing.T) {
			x := cse.x
			a := cse.a
			actual := x.Mul(a)
			expected := x * a
			tst.IntEql(t, expected, actual)
		})
	}
}
func TestIntIntDivTypFn(t *testing.T) {
	cses := []struct {
		x int.Int
		a int.Int
	}{
		{int.Int(0), int.Int(10)},
		{int.Int(10), int.Int(10)},
		{int.Int(1000), int.Int(10)},
		{int.Int(-10), int.Int(10)},
		{int.Int(-1000), int.Int(10)},
	}
	for _, cse := range cses {
		t.Run("Div", func(t *testing.T) {
			x := cse.x
			a := cse.a
			actual := x.Div(a)
			var expected int.Int
			if a != 0 {
				expected = x / a
			}
			tst.IntEql(t, expected, actual)
		})
	}
}
func TestIntIntRemTypFn(t *testing.T) {
	cses := []struct {
		x int.Int
		a int.Int
	}{
		{int.Int(0), int.Int(10)},
		{int.Int(10), int.Int(10)},
		{int.Int(1000), int.Int(10)},
		{int.Int(-10), int.Int(10)},
		{int.Int(-1000), int.Int(10)},
	}
	for _, cse := range cses {
		t.Run("Rem", func(t *testing.T) {
			x := cse.x
			a := cse.a
			actual := x.Rem(a)
			var expected int.Int
			if a != 0 {
				expected = int.Int(math.Remainder(float64(x), float64(a)))
			}
			tst.IntEql(t, expected, actual)
		})
	}
}
func TestIntIntPowTypFn(t *testing.T) {
	cses := []struct {
		x int.Int
		a int.Int
	}{
		{int.Int(0), int.Int(10)},
		{int.Int(10), int.Int(10)},
		{int.Int(1000), int.Int(10)},
		{int.Int(-10), int.Int(10)},
		{int.Int(-1000), int.Int(10)},
	}
	for _, cse := range cses {
		t.Run("Pow", func(t *testing.T) {
			x := cse.x
			a := cse.a
			actual := x.Pow(a)
			expected := int.Int(math.Pow(float64(x), float64(a)))
			tst.IntEql(t, expected, actual)
		})
	}
}
func TestIntIntSqrTypFn(t *testing.T) {
	cses := []struct {
		x int.Int
	}{
		{int.Int(0)},
		{int.Int(10)},
		{int.Int(1000)},
		{int.Int(-10)},
		{int.Int(-1000)},
	}
	for _, cse := range cses {
		t.Run("Sqr", func(t *testing.T) {
			x := cse.x
			actual := x.Sqr()
			expected := x * x
			tst.IntEql(t, expected, actual)
		})
	}
}
func TestIntIntSqrtTypFn(t *testing.T) {
	cses := []struct {
		x int.Int
	}{
		{int.Int(0)},
		{int.Int(10)},
		{int.Int(1000)},
		{int.Int(-10)},
		{int.Int(-1000)},
	}
	for _, cse := range cses {
		t.Run("Sqrt", func(t *testing.T) {
			x := cse.x
			actual := x.Sqrt()
			var expected int.Int
			if x > 0 {
				expected = int.Int(math.Sqrt(float64(x)))
			}
			tst.IntEql(t, expected, actual)
		})
	}
}
func TestIntIntMinTypFn(t *testing.T) {
	cses := []struct {
		x int.Int
		a int.Int
	}{
		{int.Int(0), int.Int(10)},
		{int.Int(10), int.Int(10)},
		{int.Int(1000), int.Int(10)},
		{int.Int(-10), int.Int(10)},
		{int.Int(-1000), int.Int(10)},
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
			tst.IntEql(t, expected, actual)
		})
	}
}
func TestIntIntMaxTypFn(t *testing.T) {
	cses := []struct {
		x int.Int
		a int.Int
	}{
		{int.Int(0), int.Int(10)},
		{int.Int(10), int.Int(10)},
		{int.Int(1000), int.Int(10)},
		{int.Int(-10), int.Int(10)},
		{int.Int(-1000), int.Int(10)},
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
			tst.IntEql(t, expected, actual)
		})
	}
}
func TestIntIntMinMaxTypFn(t *testing.T) {
	cses := []struct {
		x int.Int
		a int.Int
	}{
		{int.Int(0), int.Int(10)},
		{int.Int(10), int.Int(10)},
		{int.Int(1000), int.Int(10)},
		{int.Int(-10), int.Int(10)},
		{int.Int(-1000), int.Int(10)},
	}
	for _, cse := range cses {
		t.Run("MinMax", func(t *testing.T) {
			x := cse.x
			a := cse.a
			aMin, aMax := x.MinMax(a)
			var eMin, eMax int.Int
			if x < a {
				eMin, eMax = x, a
			} else {
				eMin, eMax = a, x
			}
			tst.IntEql(t, eMin, aMin, "Min")
			tst.IntEql(t, eMax, aMax, "Max")
		})
	}
}
func TestIntIntMidTypFn(t *testing.T) {
	cses := []struct {
		x int.Int
		a int.Int
	}{
		{int.Int(0), int.Int(10)},
		{int.Int(10), int.Int(10)},
		{int.Int(1000), int.Int(10)},
		{int.Int(-10), int.Int(10)},
		{int.Int(-1000), int.Int(10)},
	}
	for _, cse := range cses {
		t.Run("Mid", func(t *testing.T) {
			x := cse.x
			a := cse.a
			actual := x.Mid(a)
			expected := (x - a) / 2
			tst.IntEql(t, expected, actual)
		})
	}
}
func TestIntIntAvgTypFn(t *testing.T) {
	cses := []struct {
		x int.Int
		a int.Int
	}{
		{int.Int(0), int.Int(10)},
		{int.Int(10), int.Int(10)},
		{int.Int(1000), int.Int(10)},
		{int.Int(-10), int.Int(10)},
		{int.Int(-1000), int.Int(10)},
	}
	for _, cse := range cses {
		t.Run("Avg", func(t *testing.T) {
			x := cse.x
			a := cse.a
			actual := x.Avg(a)
			expected := (x + a) / 2
			tst.IntEql(t, expected, actual)
		})
	}
}
func TestIntIntAvgGeoTypFn(t *testing.T) {
	cses := []struct {
		x int.Int
		a int.Int
	}{
		{int.Int(0), int.Int(10)},
		{int.Int(10), int.Int(10)},
		{int.Int(1000), int.Int(10)},
		{int.Int(-10), int.Int(10)},
		{int.Int(-1000), int.Int(10)},
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
			tst.IntEql(t, expected, actual)
		})
	}
}
func TestIntIntByt(t *testing.T) {
	cses := []struct {
		e int.Int
	}{
		{0},
		{10},
		{1000},
		{-10},
		{-1000},
	}
	for _, cse := range cses {
		t.Run("Byt", func(t *testing.T) {
			b := &bytes.Buffer{}
			cse.e.BytWrt(b)
			var a int.Int
			a.BytRed(b.Bytes())
			tst.IntEql(t, cse.e, a)
		})
	}
}
