package bol_test

import (
	"bytes"
	"fmt"
	"sys/bsc/bol"
	"sys/tst"
	"testing"
)

func TestBolBolCnst(t *testing.T) {
	cses := []struct {
		idn string
		e   bol.Bol
		a   bol.Bol
	}{
		{"Zero", bol.Bol(false), bol.Zero},
		{"Fls", bol.Bol(false), bol.Fls},
		{"Tru", bol.Bol(true), bol.Tru},
	}
	for _, cse := range cses {
		t.Run(fmt.Sprintf("%q", cse.idn), func(t *testing.T) {
			tst.BolEql(t, cse.e, cse.a)
		})
	}
}
func TestBolBolNotTypFn(t *testing.T) {
	cses := []struct {
		x bol.Bol
	}{
		{bol.Bol(false)},
		{bol.Bol(true)},
	}
	for _, cse := range cses {
		t.Run("Not", func(t *testing.T) {
			x := cse.x
			actual := x.Not()
			expected := !cse.x
			tst.BolEql(t, expected, actual)
		})
	}
}
func TestBolBolEqlTypFn(t *testing.T) {
	cses := []struct {
		x bol.Bol
		a bol.Bol
	}{
		{bol.Bol(false), bol.Bol(true)},
		{bol.Bol(true), bol.Bol(true)},
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
func TestBolBolNeqTypFn(t *testing.T) {
	cses := []struct {
		x bol.Bol
		a bol.Bol
	}{
		{bol.Bol(false), bol.Bol(true)},
		{bol.Bol(true), bol.Bol(true)},
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
func TestBolBolByt(t *testing.T) {
	cses := []struct {
		e bol.Bol
	}{
		{false},
		{true},
	}
	for _, cse := range cses {
		t.Run("Byt", func(t *testing.T) {
			b := &bytes.Buffer{}
			cse.e.BytWrt(b)
			var a bol.Bol
			a.BytRed(b.Bytes())
			tst.BolEql(t, cse.e, a)
		})
	}
}
