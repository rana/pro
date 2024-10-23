package tst

import (
	"testing"
)

func TestNotNilA(t *testing.T) {
	NotNil(t, t)
}

func TestNotNilB(t *testing.T) {
	var b interface{}
	b = t
	NotNil(t, b)
}
