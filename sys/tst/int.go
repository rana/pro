package tst

import (
	"fmt"
	"testing"
)

// IntegerEql tests whether a int value is equal to another int value.
func IntegerEql(t *testing.T, e, a int, msgs ...interface{}) {
	if e != a {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("should be equal (expected:%v actual:%v)", e, a))...)
	}
}
func IntegerZero(t *testing.T, a int, msgs ...interface{}) {
	if a != 0 {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("integer not zero (actual:%v)", a))...)
	}
}
func Uint32NotZero(t *testing.T, a uint32, msgs ...interface{}) {
	if a == 0 {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("integer not zero (actual:%v)", a))...)
	}
}
