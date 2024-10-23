package str_test

import (
	"bytes"
	"fmt"
	"strings"
	"sys/bsc/bol"
	"sys/bsc/str"
	"sys/tst"
	"testing"
)

func TestStrStrCnst(t *testing.T) {
	cses := []struct {
		idn string
		e   str.Str
		a   str.Str
	}{
		{"Zero", str.Str(""), str.Zero},
		{"Empty", str.Str(""), str.Empty},
	}
	for _, cse := range cses {
		t.Run(fmt.Sprintf("%q", cse.idn), func(t *testing.T) {
			tst.StrEql(t, cse.e, cse.a)
		})
	}
}
func TestStrStrEqlPkgFn(t *testing.T) {
	cses := []struct {
		a str.Str
		b str.Str
	}{
		{"", ""},
		{"", "xYz"},
		{"", "a"},
		{"", "efg HIJ jKl"},
		{"xYz", ""},
		{"xYz", "xYz"},
		{"xYz", "a"},
		{"xYz", "efg HIJ jKl"},
		{"a", ""},
		{"a", "xYz"},
		{"a", "a"},
		{"a", "efg HIJ jKl"},
		{"efg HIJ jKl", ""},
		{"efg HIJ jKl", "xYz"},
		{"efg HIJ jKl", "a"},
		{"efg HIJ jKl", "efg HIJ jKl"},
	}
	for _, cse := range cses {
		t.Run("Eql", func(t *testing.T) {
			a := cse.a
			b := cse.b
			actual := str.Eql(a, b)
			expected := bol.Bol(a == b)
			tst.BolEql(t, expected, actual)
		})
	}
}
func TestStrStrLssPkgFn(t *testing.T) {
	cses := []struct {
		a str.Str
		b str.Str
	}{
		{"", ""},
		{"", "xYz"},
		{"", "a"},
		{"", "efg HIJ jKl"},
		{"xYz", ""},
		{"xYz", "xYz"},
		{"xYz", "a"},
		{"xYz", "efg HIJ jKl"},
		{"a", ""},
		{"a", "xYz"},
		{"a", "a"},
		{"a", "efg HIJ jKl"},
		{"efg HIJ jKl", ""},
		{"efg HIJ jKl", "xYz"},
		{"efg HIJ jKl", "a"},
		{"efg HIJ jKl", "efg HIJ jKl"},
	}
	for _, cse := range cses {
		t.Run("Lss", func(t *testing.T) {
			a := cse.a
			b := cse.b
			actual := str.Lss(a, b)
			expected := bol.Bol(a < b)
			tst.BolEql(t, expected, actual)
		})
	}
}
func TestStrStrGtrPkgFn(t *testing.T) {
	cses := []struct {
		a str.Str
		b str.Str
	}{
		{"", ""},
		{"", "xYz"},
		{"", "a"},
		{"", "efg HIJ jKl"},
		{"xYz", ""},
		{"xYz", "xYz"},
		{"xYz", "a"},
		{"xYz", "efg HIJ jKl"},
		{"a", ""},
		{"a", "xYz"},
		{"a", "a"},
		{"a", "efg HIJ jKl"},
		{"efg HIJ jKl", ""},
		{"efg HIJ jKl", "xYz"},
		{"efg HIJ jKl", "a"},
		{"efg HIJ jKl", "efg HIJ jKl"},
	}
	for _, cse := range cses {
		t.Run("Gtr", func(t *testing.T) {
			a := cse.a
			b := cse.b
			actual := str.Gtr(a, b)
			expected := bol.Bol(a > b)
			tst.BolEql(t, expected, actual)
		})
	}
}
func TestStrStrLowerTypFn(t *testing.T) {
	cses := []struct {
		x str.Str
	}{
		{str.Str("")},
		{str.Str("xYz")},
		{str.Str("a")},
		{str.Str("efg HIJ jKl")},
	}
	for _, cse := range cses {
		t.Run("Lower", func(t *testing.T) {
			x := cse.x
			actual := x.Lower()
			expected := str.Str(strings.ToLower(string(cse.x)))
			tst.StrEql(t, expected, actual)
		})
	}
}
func TestStrStrUpperTypFn(t *testing.T) {
	cses := []struct {
		x str.Str
	}{
		{str.Str("")},
		{str.Str("xYz")},
		{str.Str("a")},
		{str.Str("efg HIJ jKl")},
	}
	for _, cse := range cses {
		t.Run("Upper", func(t *testing.T) {
			x := cse.x
			actual := x.Upper()
			expected := str.Str(strings.ToUpper(string(cse.x)))
			tst.StrEql(t, expected, actual)
		})
	}
}
func TestStrStrEqlTypFn(t *testing.T) {
	cses := []struct {
		x str.Str
		a str.Str
	}{
		{str.Str(""), str.Str("xYz")},
		{str.Str("xYz"), str.Str("xYz")},
		{str.Str("a"), str.Str("xYz")},
		{str.Str("efg HIJ jKl"), str.Str("xYz")},
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
func TestStrStrNeqTypFn(t *testing.T) {
	cses := []struct {
		x str.Str
		a str.Str
	}{
		{str.Str(""), str.Str("xYz")},
		{str.Str("xYz"), str.Str("xYz")},
		{str.Str("a"), str.Str("xYz")},
		{str.Str("efg HIJ jKl"), str.Str("xYz")},
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
func TestStrStrLssTypFn(t *testing.T) {
	cses := []struct {
		x str.Str
		a str.Str
	}{
		{str.Str(""), str.Str("xYz")},
		{str.Str("xYz"), str.Str("xYz")},
		{str.Str("a"), str.Str("xYz")},
		{str.Str("efg HIJ jKl"), str.Str("xYz")},
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
func TestStrStrGtrTypFn(t *testing.T) {
	cses := []struct {
		x str.Str
		a str.Str
	}{
		{str.Str(""), str.Str("xYz")},
		{str.Str("xYz"), str.Str("xYz")},
		{str.Str("a"), str.Str("xYz")},
		{str.Str("efg HIJ jKl"), str.Str("xYz")},
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
func TestStrStrLeqTypFn(t *testing.T) {
	cses := []struct {
		x str.Str
		a str.Str
	}{
		{str.Str(""), str.Str("xYz")},
		{str.Str("xYz"), str.Str("xYz")},
		{str.Str("a"), str.Str("xYz")},
		{str.Str("efg HIJ jKl"), str.Str("xYz")},
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
func TestStrStrGeqTypFn(t *testing.T) {
	cses := []struct {
		x str.Str
		a str.Str
	}{
		{str.Str(""), str.Str("xYz")},
		{str.Str("xYz"), str.Str("xYz")},
		{str.Str("a"), str.Str("xYz")},
		{str.Str("efg HIJ jKl"), str.Str("xYz")},
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
func TestStrStrByt(t *testing.T) {
	cses := []struct {
		e str.Str
	}{
		{""},
		{"xYz"},
		{"a"},
		{"efg HIJ jKl"},
	}
	for _, cse := range cses {
		t.Run("Byt", func(t *testing.T) {
			b := &bytes.Buffer{}
			cse.e.BytWrt(b)
			var a str.Str
			a.BytRed(b.Bytes())
			tst.StrEql(t, cse.e, a)
		})
	}
}
