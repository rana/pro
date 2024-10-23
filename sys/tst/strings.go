package tst

import (
	"fmt"
	"testing"
)

// StringsEql tests whether a string slice is equal to another string slice.
func StringsEql(t *testing.T, e, a []string, msgs ...interface{}) {
	if e == nil {
		t.Helper()
		t.Fatal(append(msgs, "should be equal Expected nil")...)
	}
	if a == nil {
		t.Helper()
		t.Fatal(append(msgs, "should be equal. Actual nil")...)
	}
	if len(e) != len(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Errorf("should be equal. Unequal slice lengths (expected:%v actual:%v)", len(e), len(a)))...)
	}
	for n := 0; n < len(e); n++ {
		if e[n] != a[n] {
			t.Helper()
			t.Fatal(append(msgs, fmt.Errorf("should be equal. Unequal element at index %v (expected:%v, actual:%v)", n, e[n], a[n]))...)
		}
	}
}
