package tst

import (
	"fmt"
	"reflect"
	"testing"
)

// TypeEql tests whether the types of two values are equal.
func TypeEql(t *testing.T, e, a interface{}, msgs ...interface{}) {
	if reflect.TypeOf(e) != reflect.TypeOf(a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("should be equal (expected:%v actual:%v)", reflect.TypeOf(e).Name(), reflect.TypeOf(a).Name()))...)
	}
}
