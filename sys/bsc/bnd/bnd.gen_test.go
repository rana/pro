package bnd_test

import (
	"bytes"
	"sys/bsc/bnd"
	"sys/bsc/bol"
	"sys/bsc/unt"
	"sys/tst"
	"testing"
)

func TestBndBndCntTypFn(t *testing.T) {
	cses := []struct {
		x bnd.Bnd
	}{
		{bnd.Bnd(bnd.Bnd{Idx: 0, Lim: 0})},
		{bnd.Bnd(bnd.Bnd{Idx: 0, Lim: 1})},
		{bnd.Bnd(bnd.Bnd{Idx: 0, Lim: 1000})},
		{bnd.Bnd(bnd.Bnd{Idx: 999, Lim: 1000})},
		{bnd.Bnd(bnd.Bnd{Idx: 1, Lim: 0})},
	}
	for _, cse := range cses {
		t.Run("Cnt", func(t *testing.T) {
			x := cse.x
			actual := x.Cnt()
			expected := cse.x.Lim - cse.x.Idx
			tst.UntEql(t, expected, actual)
		})
	}
}
func TestBndBndLenTypFn(t *testing.T) {
	cses := []struct {
		x bnd.Bnd
	}{
		{bnd.Bnd(bnd.Bnd{Idx: 0, Lim: 0})},
		{bnd.Bnd(bnd.Bnd{Idx: 0, Lim: 1})},
		{bnd.Bnd(bnd.Bnd{Idx: 0, Lim: 1000})},
		{bnd.Bnd(bnd.Bnd{Idx: 999, Lim: 1000})},
		{bnd.Bnd(bnd.Bnd{Idx: 1, Lim: 0})},
	}
	for _, cse := range cses {
		t.Run("Len", func(t *testing.T) {
			x := cse.x
			actual := x.Len()
			expected := cse.x.Lim - cse.x.Idx
			tst.UntEql(t, expected, actual)
		})
	}
}
func TestBndBndLstIdxTypFn(t *testing.T) {
	cses := []struct {
		x bnd.Bnd
	}{
		{bnd.Bnd(bnd.Bnd{Idx: 0, Lim: 0})},
		{bnd.Bnd(bnd.Bnd{Idx: 0, Lim: 1})},
		{bnd.Bnd(bnd.Bnd{Idx: 0, Lim: 1000})},
		{bnd.Bnd(bnd.Bnd{Idx: 999, Lim: 1000})},
		{bnd.Bnd(bnd.Bnd{Idx: 1, Lim: 0})},
	}
	for _, cse := range cses {
		t.Run("LstIdx", func(t *testing.T) {
			x := cse.x
			actual := x.LstIdx()
			var expected unt.Unt
			if cse.x.Lim > 0 {
				expected = cse.x.Lim - 1
			}
			tst.UntEql(t, expected, actual)
		})
	}
}
func TestBndBndIsValidTypFn(t *testing.T) {
	cses := []struct {
		x bnd.Bnd
	}{
		{bnd.Bnd(bnd.Bnd{Idx: 0, Lim: 0})},
		{bnd.Bnd(bnd.Bnd{Idx: 0, Lim: 1})},
		{bnd.Bnd(bnd.Bnd{Idx: 0, Lim: 1000})},
		{bnd.Bnd(bnd.Bnd{Idx: 999, Lim: 1000})},
		{bnd.Bnd(bnd.Bnd{Idx: 1, Lim: 0})},
	}
	for _, cse := range cses {
		t.Run("IsValid", func(t *testing.T) {
			x := cse.x
			actual := x.IsValid()
			expected := bol.Bol(cse.x.Idx < cse.x.Lim)
			tst.BolEql(t, expected, actual)
		})
	}
}
func TestBndBndByt(t *testing.T) {
	cses := []struct {
		e bnd.Bnd
	}{
		{bnd.Bnd{Idx: 0, Lim: 0}},
		{bnd.Bnd{Idx: 0, Lim: 1}},
		{bnd.Bnd{Idx: 0, Lim: 1000}},
		{bnd.Bnd{Idx: 999, Lim: 1000}},
		{bnd.Bnd{Idx: 1, Lim: 0}},
	}
	for _, cse := range cses {
		t.Run("Byt", func(t *testing.T) {
			b := &bytes.Buffer{}
			cse.e.BytWrt(b)
			var a bnd.Bnd
			a.BytRed(b.Bytes())
			tst.BndEql(t, cse.e, a)
		})
	}
}
