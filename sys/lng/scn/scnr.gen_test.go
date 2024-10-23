package scn_test

import (
	"sys/lng/scn"
	"sys/tst"
	"testing"
)

func TestScnScnrResetTypFn(t *testing.T) {
	var a scn.Scnr
	a.Reset("")
	tst.True(t, a.End)
	tst.RuneEql(t, scn.EndCh, a.Ch)
	tst.UntZero(t, a.Idx)
	tst.UntOne(t, a.Ln)
	tst.UntZero(t, a.Col)
	tst.IntegerZero(t, a.Size)
}
func TestScnScnrNextRuneTypFn(t *testing.T) {
	var a scn.Scnr
	a.Reset("abc")
	tst.False(t, a.End)
	tst.RuneEql(t, 'a', a.Ch)
	tst.UntEql(t, 0, a.Idx)
	tst.UntEql(t, 1, a.Ln)
	tst.UntEql(t, 1, a.Col)
	tst.IntegerEql(t, 1, a.Size)

	a.NextRune()
	tst.False(t, a.End)
	tst.RuneEql(t, 'b', a.Ch)
	tst.UntEql(t, 1, a.Idx)
	tst.UntEql(t, 1, a.Ln)
	tst.UntEql(t, 2, a.Col)
	tst.IntegerEql(t, 1, a.Size)

	a.NextRune()
	tst.False(t, a.End)
	tst.RuneEql(t, 'c', a.Ch)
	tst.UntEql(t, 2, a.Idx)
	tst.UntEql(t, 1, a.Ln)
	tst.UntEql(t, 3, a.Col)
	tst.IntegerEql(t, 1, a.Size)

	a.NextRune()
	tst.True(t, a.End)
	tst.RuneEql(t, scn.EndCh, a.Ch)
}
func TestScnScnrPeekRuneTypFn(t *testing.T) {
	var a scn.Scnr
	a.Reset("")
	tst.RuneEql(t, scn.EndCh, a.PeekRune())
	tst.RuneEql(t, scn.EndCh, a.Ch)
	tst.UntEql(t, 0, a.Idx)
	tst.UntEql(t, 1, a.Ln)
	tst.UntEql(t, 0, a.Col)
	tst.IntegerEql(t, 0, a.Size)
	tst.True(t, a.End)

	a.Reset("a")
	tst.RuneEql(t, 'a', a.Ch)
	tst.UntEql(t, 0, a.Idx)
	tst.UntEql(t, 1, a.Ln)
	tst.UntEql(t, 1, a.Col)
	tst.IntegerEql(t, 1, a.Size)
	tst.False(t, a.End)

	tst.RuneEql(t, scn.EndCh, a.PeekRune())
	tst.RuneEql(t, 'a', a.Ch)
	tst.UntEql(t, 0, a.Idx)
	tst.UntEql(t, 1, a.Ln)
	tst.UntEql(t, 1, a.Col)
	tst.IntegerEql(t, 1, a.Size)
	tst.False(t, a.End)
}
