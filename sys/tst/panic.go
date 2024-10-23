package tst

import (
	"sys/err"
	"testing"
)

// Panic tests whether the value is nil.
func Panic(t *testing.T, f func(), msgs ...interface{}) {
	var er error
	func() {
		defer func() {
			if v := recover(); v != nil {
				er = err.New(er)
			}
		}()
		f()
	}()
	if er == nil {
		t.Helper()
		t.Fatal(append(msgs, "should panic")...)
	}
}
