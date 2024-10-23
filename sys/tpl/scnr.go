package tpl

import (
	"sys/k"
	"sys/tpl/atr"
)

type (
	FleScnr struct {
		FleBse
		Scn   *DirScn
		EndCh *Cnst
	}
)

func (x *DirScn) NewScnr() (r *FleScnr) {
	r = &FleScnr{Scn: x}
	r.Name = k.Scnr
	r.Pkg = x.Pkg
	r.Scnr()
	r.AddFle(r)
	return r
}
func (x *FleScnr) Scnr() (r *Struct) {
	r = x.StructPtr(k.Scnr, atr.Test) // Scnr is a rune scanner.
	r.FldTyp(x.Scn.Scn.Typ())
	r.Fld("Txt", String) // the next position
	return r
}
func (x *FleScnr) InitCnst() {
	x.Cnst("BOM", "0xFEFF", Rune) // BOM is a rune representing a byte order mark.
	x.EndCh = x.Cnst("EndCh", "-1", Rune)
}
func (x *FleScnr) InitTypFn() {
	x.Reset()
	x.Resetf()
	x.NextRune()
	x.PeekRune()
	x.SkpSet()
	x.PrintScn()
}
func (x *FleScnr) Reset() (r *TypFn) {
	r = x.TypFn("Reset") // Reset sets the scanner to the starting position.
	r.InPrm(String, "txt")
	r.Add("x.Txt = txt")
	r.Add("x.Ch = ' '")
	r.Add("x.Size = 0")
	r.Add("x.Idx = 0")
	r.Add("x.Ln = 1")
	r.Add("x.Col = 0")
	r.Add("x.End = false")
	r.Add("x.NextRune()")
	r.Add("if x.Ch == BOM {")
	r.Add("x.NextRune()")
	r.Add("}")
	// test
	r.T.Empty = true
	r.T.Addf("var a %v", x.Typ().Full())
	r.T.Add("a.Reset(\"\")")
	r.T.Add("tst.True(t, a.End)")
	r.T.Addf("tst.RuneEql(t, %v, a.Ch)", x.EndCh.Ref(x.Test))
	r.T.Add("tst.UntZero(t, a.Idx)")
	r.T.Add("tst.UntOne(t, a.Ln)")
	r.T.Add("tst.UntZero(t, a.Col)")
	r.T.Add("tst.IntegerZero(t, a.Size)")
	return r
}
func (x *FleScnr) Resetf() (r *TypFn) {
	r = x.TypFn("Resetf")
	r.InPrm(String, "format")
	r.InPrmVariadic(Interface, "args")
	r.Add("x.Reset(fmt.Sprintf(format, args...))")
	return r
}
func (x *FleScnr) NextRune() (r *TypFn) {
	x.Import("unicode/utf8")
	x.ImportTyp(_sys.Bsc.Unt.Typ())
	r = x.TypFn("NextRune") // NextRune decodes the next rune and position. True is returned if there are more runes to decode.
	r.OutPrm(Bool)
	r.Addf("x.Idx += %v(x.Size)", _sys.Bsc.Unt.Typ().Full())
	r.Addf("x.End = x.Idx >= %v(len(x.Txt))", _sys.Bsc.Unt.Typ().Full())
	r.Add("if x.End {")
	r.Add("x.Ch = EndCh // necessary for char check beyond lim to fail")
	r.Add("} else {")
	r.Add("x.Ch, x.Size = utf8.DecodeRuneInString(x.Txt[x.Idx:])")
	r.Add("if x.Ch == '\\n' {")
	r.Add("x.Ln++")
	r.Add("x.Col = 0")
	r.Add("} else {")
	r.Add("x.Col++")
	r.Add("}")
	r.Add("}")
	r.Add("return x.End")
	// test
	r.T.Empty = true
	r.T.Addf("var a %v", x.Typ().Full())
	r.T.Add("a.Reset(\"abc\")")
	r.T.Add("tst.False(t, a.End)")
	r.T.Add("tst.RuneEql(t, 'a', a.Ch)")
	r.T.Add("tst.UntEql(t, 0, a.Idx)")
	r.T.Add("tst.UntEql(t, 1, a.Ln)")
	r.T.Add("tst.UntEql(t, 1, a.Col)")
	r.T.Add("tst.IntegerEql(t, 1, a.Size)")
	r.T.Add("")
	r.T.Add("a.NextRune()")
	r.T.Add("tst.False(t, a.End)")
	r.T.Add("tst.RuneEql(t, 'b', a.Ch)")
	r.T.Add("tst.UntEql(t, 1, a.Idx)")
	r.T.Add("tst.UntEql(t, 1, a.Ln)")
	r.T.Add("tst.UntEql(t, 2, a.Col)")
	r.T.Add("tst.IntegerEql(t, 1, a.Size)")
	r.T.Add("")
	r.T.Add("a.NextRune()")
	r.T.Add("tst.False(t, a.End)")
	r.T.Add("tst.RuneEql(t, 'c', a.Ch)")
	r.T.Add("tst.UntEql(t, 2, a.Idx)")
	r.T.Add("tst.UntEql(t, 1, a.Ln)")
	r.T.Add("tst.UntEql(t, 3, a.Col)")
	r.T.Add("tst.IntegerEql(t, 1, a.Size)")
	r.T.Add("")
	r.T.Add("a.NextRune()")
	r.T.Add("tst.True(t, a.End)")
	r.T.Addf("tst.RuneEql(t, %v, a.Ch)", x.EndCh.Ref(x.Test))
	return r
}
func (x *FleScnr) PeekRune() (r *TypFn) {
	r = x.TypFn("PeekRune") // PeekRune returns the next rune without moving the Scnr position.
	r.OutPrm(Rune, "r")
	r.Add("scn := x.Scn")
	r.Add("x.NextRune()")
	r.Add("r = x.Ch")
	r.Add("x.Scn = scn")
	r.Add("return r")
	// test
	r.T.Empty = true
	r.T.Addf("var a %v", x.Typ().Full())
	r.T.Add("a.Reset(\"\")")
	r.T.Addf("tst.RuneEql(t, %v, a.%v())", x.EndCh.Ref(x.Test), r.Name)
	r.T.Addf("tst.RuneEql(t, %v, a.Ch)", x.EndCh.Ref(x.Test))
	r.T.Add("tst.UntEql(t, 0, a.Idx)")
	r.T.Add("tst.UntEql(t, 1, a.Ln)")
	r.T.Add("tst.UntEql(t, 0, a.Col)")
	r.T.Add("tst.IntegerEql(t, 0, a.Size)")
	r.T.Add("tst.True(t, a.End)")
	r.T.Add("")
	r.T.Add("a.Reset(\"a\")")
	r.T.Addf("tst.RuneEql(t, 'a', a.Ch)")
	r.T.Add("tst.UntEql(t, 0, a.Idx)")
	r.T.Add("tst.UntEql(t, 1, a.Ln)")
	r.T.Add("tst.UntEql(t, 1, a.Col)")
	r.T.Add("tst.IntegerEql(t, 1, a.Size)")
	r.T.Add("tst.False(t, a.End)")
	r.T.Add("")
	r.T.Addf("tst.RuneEql(t, %v, a.%v())", x.EndCh.Ref(x.Test), r.Name)
	r.T.Addf("tst.RuneEql(t, 'a', a.Ch)")
	r.T.Add("tst.UntEql(t, 0, a.Idx)")
	r.T.Add("tst.UntEql(t, 1, a.Ln)")
	r.T.Add("tst.UntEql(t, 1, a.Col)")
	r.T.Add("tst.IntegerEql(t, 1, a.Size)")
	r.T.Add("tst.False(t, a.End)")
	return r
}
func (x *FleScnr) SkpSet() (r *TypFn) {
	x.Import("sys/err")
	r = x.TypFn("SkpSet")
	r.InPrm(Rune, "opn")
	r.InPrm(Rune, "cls")
	r.Add("if x.Ch != opn { // must start with opn rune")
	r.Addf("err.Panicf(\"%v: missing open rune '%%v'\", string(opn))", x.Title())
	r.Add("}")
	r.Add("x.NextRune() // skp opn")
	r.Add("if x.Ch != cls {")
	r.Add("depth := 0")
	r.Add("for !x.End {")
	r.Add("x.NextRune()")
	r.Add("if x.Ch == opn {")
	r.Add("depth++")
	r.Add("} else if x.Ch == cls {")
	r.Add("if depth == 0 {")
	r.Add("break")
	r.Add("} else {")
	r.Add("depth--")
	r.Add("}")
	r.Add("}")
	r.Add("}")
	r.Add("if x.Ch != cls { // must end with cls rune")
	r.Addf("err.Panicf(\"%v: missing close rune '%%v'\", string(cls))", x.Title())
	r.Add("}")
	r.Add("}")
	r.Add("x.NextRune() // skp cls")
	return r
}
func (x *FleScnr) PrintScn() (r *TypFn) {
	x.Import("fmt")
	r = x.TypFn("PrintScn")
	r.Add("fmt.Printf(\"Ch:'%v' %v Size:%v Idx:%v Ln:%v Col:%v End:%v \\n\", string(x.Ch), x.Ch, x.Size, x.Idx, x.Ln, x.Col, x.End)")
	return r
}
