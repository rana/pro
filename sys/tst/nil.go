package tst

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func Nil(t *testing.T, v interface{}, msgs ...interface{}) {
	rv := reflect.ValueOf(v)
	if !(v == nil || (rv.Kind() != reflect.Struct && rv.IsNil())) {
		t.Helper()
		t.Fatal(append(msgs, "should be nil")...)
	}
}
func NotNil(t *testing.T, v interface{}, msgs ...interface{}) {
	rv := reflect.ValueOf(v)
	if v == nil || (rv.Kind() != reflect.Struct && rv.IsNil()) {
		t.Helper()
		t.Fatal(append(msgs, "should not be nil")...)
	}
}
func TimeLss(t *testing.T, e, a time.Time, msgs ...interface{}) {
	if !(e.Before(a)) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("should be Lss (expected:%v actual:%v)", e, a))...)
	}
}
func TimeNeq(t *testing.T, e, a time.Time, msgs ...interface{}) {
	if !(e.Unix() != a.Unix()) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("should be Lss (expected:%v actual:%v)", e, a))...)
	}
}
func DurationLss(t *testing.T, e, a time.Duration, msgs ...interface{}) {
	if !(e < a) {
		t.Helper()
		t.Fatal(append(msgs, fmt.Sprintf("should be Lss (expected:%v actual:%v)", e, a))...)
	}
}
