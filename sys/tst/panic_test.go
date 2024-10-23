package tst

import "testing"

func TestPanic(t *testing.T) {
	f := func() { panic("testing") }
	Panic(t, f)
}
