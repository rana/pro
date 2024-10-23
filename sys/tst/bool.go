package tst

import (
	"testing"
)

func False(t *testing.T, v bool, msgs ...interface{}) {
	if v {
		t.Helper()
		t.Fatal(append(msgs, "should be false")...)
	}
}
func True(t *testing.T, v bool, msgs ...interface{}) {
	if !v {
		t.Helper()
		t.Fatal(append(msgs, "should be true")...)
	}
}
