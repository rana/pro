package tst

import (
	"testing"
)

func TestStringsEql(t *testing.T) {
	StringsEql(t,
		[]string{"0", "1", "2"},
		[]string{"0", "1", "2"},
	)
}
