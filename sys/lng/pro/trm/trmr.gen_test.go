package trm_test

import (
	"fmt"
	"sys/bsc/bnd"
	"sys/bsc/unt"
	"sys/lng/pro/trm"
	"sys/tst"
	"testing"
)

func TestSpceLitValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"\t", 1},
		{"\n", 1},
		{"\v", 1},
		{"\f", 1},
		{"\r", 1},
		{" ", 1},
		{"\u0085", 2},
		{"\u00a0", 2},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.SpceLit()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, trm.SpceLit{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestSpceLitInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_\t"},
		{"9713\t"},
		{"Z\t"},
		{"_\n"},
		{"9713\n"},
		{"Z\n"},
		{"_\v"},
		{"9713\v"},
		{"Z\v"},
		{"_\f"},
		{"9713\f"},
		{"Z\f"},
		{"_\r"},
		{"9713\r"},
		{"Z\r"},
		{"_ "},
		{"9713 "},
		{"Z "},
		{"_\u0085"},
		{"9713\u0085"},
		{"Z\u0085"},
		{"_\u00a0"},
		{"9713\u00a0"},
		{"Z\u00a0"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.SpceLit()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestCmntLitValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"//", 2},
		{"////", 4},
		{"//a", 3},
		{"//0", 3},
		{"// ", 3},
		{"// abc123", 9},
		{"//\n", 3},
		{"// comment\n", 11},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.CmntLit()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, trm.CmntLit{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestCmntLitInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_/"},
		{"_//"},
		{"9713/"},
		{"9713//"},
		{"Z/"},
		{"Z//"},
		{"_/"},
		{"_//"},
		{"_///"},
		{"_////"},
		{"9713/"},
		{"9713//"},
		{"9713///"},
		{"9713////"},
		{"Z/"},
		{"Z//"},
		{"Z///"},
		{"Z////"},
		{"_/"},
		{"_//"},
		{"_//a"},
		{"9713/"},
		{"9713//"},
		{"9713//a"},
		{"Z/"},
		{"Z//"},
		{"Z//a"},
		{"_/"},
		{"_//"},
		{"_//0"},
		{"9713/"},
		{"9713//"},
		{"9713//0"},
		{"Z/"},
		{"Z//"},
		{"Z//0"},
		{"_/"},
		{"_//"},
		{"_// "},
		{"9713/"},
		{"9713//"},
		{"9713// "},
		{"Z/"},
		{"Z//"},
		{"Z// "},
		{"_/"},
		{"_//"},
		{"_// "},
		{"_// a"},
		{"_// ab"},
		{"_// abc"},
		{"_// abc1"},
		{"_// abc12"},
		{"_// abc123"},
		{"9713/"},
		{"9713//"},
		{"9713// "},
		{"9713// a"},
		{"9713// ab"},
		{"9713// abc"},
		{"9713// abc1"},
		{"9713// abc12"},
		{"9713// abc123"},
		{"Z/"},
		{"Z//"},
		{"Z// "},
		{"Z// a"},
		{"Z// ab"},
		{"Z// abc"},
		{"Z// abc1"},
		{"Z// abc12"},
		{"Z// abc123"},
		{"_/"},
		{"_//"},
		{"_//\n"},
		{"9713/"},
		{"9713//"},
		{"9713//\n"},
		{"Z/"},
		{"Z//"},
		{"Z//\n"},
		{"_/"},
		{"_//"},
		{"_// "},
		{"_// c"},
		{"_// co"},
		{"_// com"},
		{"_// comm"},
		{"_// comme"},
		{"_// commen"},
		{"_// comment"},
		{"_// comment\n"},
		{"9713/"},
		{"9713//"},
		{"9713// "},
		{"9713// c"},
		{"9713// co"},
		{"9713// com"},
		{"9713// comm"},
		{"9713// comme"},
		{"9713// commen"},
		{"9713// comment"},
		{"9713// comment\n"},
		{"Z/"},
		{"Z//"},
		{"Z// "},
		{"Z// c"},
		{"Z// co"},
		{"Z// com"},
		{"Z// comm"},
		{"Z// comme"},
		{"Z// commen"},
		{"Z// comment"},
		{"Z// comment\n"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.CmntLit()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestIdnLitValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"_", 1},
		{"_ ", 1},
		{"_\n", 1},
		{"_.", 1},
		{"_:", 1},
		{"_,", 1},
		{"_\"", 1},
		{"_(", 1},
		{"_)", 1},
		{"_[", 1},
		{"_]", 1},
		{"_// comment", 1},
		{"_a", 2},
		{"_a ", 2},
		{"_a\n", 2},
		{"_a.", 2},
		{"_a:", 2},
		{"_a,", 2},
		{"_a\"", 2},
		{"_a(", 2},
		{"_a)", 2},
		{"_a[", 2},
		{"_a]", 2},
		{"_a// comment", 2},
		{"_0", 2},
		{"_0 ", 2},
		{"_0\n", 2},
		{"_0.", 2},
		{"_0:", 2},
		{"_0,", 2},
		{"_0\"", 2},
		{"_0(", 2},
		{"_0)", 2},
		{"_0[", 2},
		{"_0]", 2},
		{"_0// comment", 2},
		{"a", 1},
		{"a ", 1},
		{"a\n", 1},
		{"a.", 1},
		{"a:", 1},
		{"a,", 1},
		{"a\"", 1},
		{"a(", 1},
		{"a)", 1},
		{"a[", 1},
		{"a]", 1},
		{"a// comment", 1},
		{"a_", 2},
		{"a_ ", 2},
		{"a_\n", 2},
		{"a_.", 2},
		{"a_:", 2},
		{"a_,", 2},
		{"a_\"", 2},
		{"a_(", 2},
		{"a_)", 2},
		{"a_[", 2},
		{"a_]", 2},
		{"a_// comment", 2},
		{"a0", 2},
		{"a0 ", 2},
		{"a0\n", 2},
		{"a0.", 2},
		{"a0:", 2},
		{"a0,", 2},
		{"a0\"", 2},
		{"a0(", 2},
		{"a0)", 2},
		{"a0[", 2},
		{"a0]", 2},
		{"a0// comment", 2},
		{"abc123", 6},
		{"abc123 ", 6},
		{"abc123\n", 6},
		{"abc123.", 6},
		{"abc123:", 6},
		{"abc123,", 6},
		{"abc123\"", 6},
		{"abc123(", 6},
		{"abc123)", 6},
		{"abc123[", 6},
		{"abc123]", 6},
		{"abc123// comment", 6},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.IdnLit()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, trm.IdnLit{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestIdnLitInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{" _"},
		{"\n_"},
		{"._"},
		{":_"},
		{",_"},
		{"\"_"},
		{"(_"},
		{")_"},
		{"[_"},
		{"]_"},
		{"// comment_"},
		{" _"},
		{" _a"},
		{"\n_"},
		{"\n_a"},
		{"._"},
		{"._a"},
		{":_"},
		{":_a"},
		{",_"},
		{",_a"},
		{"\"_"},
		{"\"_a"},
		{"(_"},
		{"(_a"},
		{")_"},
		{")_a"},
		{"[_"},
		{"[_a"},
		{"]_"},
		{"]_a"},
		{"// comment_"},
		{"// comment_a"},
		{" _"},
		{" _0"},
		{"\n_"},
		{"\n_0"},
		{"._"},
		{"._0"},
		{":_"},
		{":_0"},
		{",_"},
		{",_0"},
		{"\"_"},
		{"\"_0"},
		{"(_"},
		{"(_0"},
		{")_"},
		{")_0"},
		{"[_"},
		{"[_0"},
		{"]_"},
		{"]_0"},
		{"// comment_"},
		{"// comment_0"},
		{" a"},
		{"\na"},
		{".a"},
		{":a"},
		{",a"},
		{"\"a"},
		{"(a"},
		{")a"},
		{"[a"},
		{"]a"},
		{"// commenta"},
		{" a"},
		{" a_"},
		{"\na"},
		{"\na_"},
		{".a"},
		{".a_"},
		{":a"},
		{":a_"},
		{",a"},
		{",a_"},
		{"\"a"},
		{"\"a_"},
		{"(a"},
		{"(a_"},
		{")a"},
		{")a_"},
		{"[a"},
		{"[a_"},
		{"]a"},
		{"]a_"},
		{"// commenta"},
		{"// commenta_"},
		{" a"},
		{" a0"},
		{"\na"},
		{"\na0"},
		{".a"},
		{".a0"},
		{":a"},
		{":a0"},
		{",a"},
		{",a0"},
		{"\"a"},
		{"\"a0"},
		{"(a"},
		{"(a0"},
		{")a"},
		{")a0"},
		{"[a"},
		{"[a0"},
		{"]a"},
		{"]a0"},
		{"// commenta"},
		{"// commenta0"},
		{" a"},
		{" ab"},
		{" abc"},
		{" abc1"},
		{" abc12"},
		{" abc123"},
		{"\na"},
		{"\nab"},
		{"\nabc"},
		{"\nabc1"},
		{"\nabc12"},
		{"\nabc123"},
		{".a"},
		{".ab"},
		{".abc"},
		{".abc1"},
		{".abc12"},
		{".abc123"},
		{":a"},
		{":ab"},
		{":abc"},
		{":abc1"},
		{":abc12"},
		{":abc123"},
		{",a"},
		{",ab"},
		{",abc"},
		{",abc1"},
		{",abc12"},
		{",abc123"},
		{"\"a"},
		{"\"ab"},
		{"\"abc"},
		{"\"abc1"},
		{"\"abc12"},
		{"\"abc123"},
		{"(a"},
		{"(ab"},
		{"(abc"},
		{"(abc1"},
		{"(abc12"},
		{"(abc123"},
		{")a"},
		{")ab"},
		{")abc"},
		{")abc1"},
		{")abc12"},
		{")abc123"},
		{"[a"},
		{"[ab"},
		{"[abc"},
		{"[abc1"},
		{"[abc12"},
		{"[abc123"},
		{"]a"},
		{"]ab"},
		{"]abc"},
		{"]abc1"},
		{"]abc12"},
		{"]abc123"},
		{"// commenta"},
		{"// commentab"},
		{"// commentabc"},
		{"// commentabc1"},
		{"// commentabc12"},
		{"// commentabc123"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.IdnLit()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestObjsLitValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"[ana.oan().hst().eurUsd().s1().bids().fst()]", 44},
		{"[ana.oan().hst().eurUsd().s1().bids().fst() ana.oan().hst().eurUsd().s1().bids().lst()]", 87},
		{"[ana.oan().hst().eurUsd().s1().bids().fst() ana.oan().hst().eurUsd().s1().bids().lst() ana.oan().hst().eurUsd().s1().bids().sum()]", 130},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.ObjsLit()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, trm.ObjsLit{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestObjsLitInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_["},
		{"_[a"},
		{"_[an"},
		{"_[ana"},
		{"_[ana."},
		{"_[ana.o"},
		{"_[ana.oa"},
		{"_[ana.oan"},
		{"_[ana.oan("},
		{"_[ana.oan()"},
		{"_[ana.oan()."},
		{"_[ana.oan().h"},
		{"_[ana.oan().hs"},
		{"_[ana.oan().hst"},
		{"_[ana.oan().hst("},
		{"_[ana.oan().hst()"},
		{"_[ana.oan().hst().eurUsd().s1().bids().fst()]"},
		{"9713["},
		{"9713[a"},
		{"9713[an"},
		{"9713[ana"},
		{"9713[ana."},
		{"9713[ana.o"},
		{"9713[ana.oa"},
		{"9713[ana.oan"},
		{"9713[ana.oan("},
		{"9713[ana.oan()"},
		{"9713[ana.oan()."},
		{"9713[ana.oan().h"},
		{"9713[ana.oan().hs"},
		{"9713[ana.oan().hst"},
		{"9713[ana.oan().hst("},
		{"9713[ana.oan().hst()"},
		{"9713[ana.oan().hst().eurUsd().s1().bids().fst()]"},
		{"Z["},
		{"Z[a"},
		{"Z[an"},
		{"Z[ana"},
		{"Z[ana."},
		{"Z[ana.o"},
		{"Z[ana.oa"},
		{"Z[ana.oan"},
		{"Z[ana.oan("},
		{"Z[ana.oan()"},
		{"Z[ana.oan()."},
		{"Z[ana.oan().h"},
		{"Z[ana.oan().hs"},
		{"Z[ana.oan().hst"},
		{"Z[ana.oan().hst("},
		{"Z[ana.oan().hst()"},
		{"Z[ana.oan().hst().eurUsd().s1().bids().fst()]"},
		{"_["},
		{"_[a"},
		{"_[an"},
		{"_[ana"},
		{"_[ana."},
		{"_[ana.o"},
		{"_[ana.oa"},
		{"_[ana.oan"},
		{"_[ana.oan("},
		{"_[ana.oan()"},
		{"_[ana.oan()."},
		{"_[ana.oan().h"},
		{"_[ana.oan().hs"},
		{"_[ana.oan().hst"},
		{"_[ana.oan().hst("},
		{"_[ana.oan().hst()"},
		{"_[ana.oan().hst().eurUsd().s1().bids().fst() ana.oan().hst().eurUsd().s1().bids().lst()]"},
		{"9713["},
		{"9713[a"},
		{"9713[an"},
		{"9713[ana"},
		{"9713[ana."},
		{"9713[ana.o"},
		{"9713[ana.oa"},
		{"9713[ana.oan"},
		{"9713[ana.oan("},
		{"9713[ana.oan()"},
		{"9713[ana.oan()."},
		{"9713[ana.oan().h"},
		{"9713[ana.oan().hs"},
		{"9713[ana.oan().hst"},
		{"9713[ana.oan().hst("},
		{"9713[ana.oan().hst()"},
		{"9713[ana.oan().hst().eurUsd().s1().bids().fst() ana.oan().hst().eurUsd().s1().bids().lst()]"},
		{"Z["},
		{"Z[a"},
		{"Z[an"},
		{"Z[ana"},
		{"Z[ana."},
		{"Z[ana.o"},
		{"Z[ana.oa"},
		{"Z[ana.oan"},
		{"Z[ana.oan("},
		{"Z[ana.oan()"},
		{"Z[ana.oan()."},
		{"Z[ana.oan().h"},
		{"Z[ana.oan().hs"},
		{"Z[ana.oan().hst"},
		{"Z[ana.oan().hst("},
		{"Z[ana.oan().hst()"},
		{"Z[ana.oan().hst().eurUsd().s1().bids().fst() ana.oan().hst().eurUsd().s1().bids().lst()]"},
		{"_["},
		{"_[a"},
		{"_[an"},
		{"_[ana"},
		{"_[ana."},
		{"_[ana.o"},
		{"_[ana.oa"},
		{"_[ana.oan"},
		{"_[ana.oan("},
		{"_[ana.oan()"},
		{"_[ana.oan()."},
		{"_[ana.oan().h"},
		{"_[ana.oan().hs"},
		{"_[ana.oan().hst"},
		{"_[ana.oan().hst("},
		{"_[ana.oan().hst()"},
		{"_[ana.oan().hst().eurUsd().s1().bids().fst() ana.oan().hst().eurUsd().s1().bids().lst() ana.oan().hst().eurUsd().s1().bids().sum()]"},
		{"9713["},
		{"9713[a"},
		{"9713[an"},
		{"9713[ana"},
		{"9713[ana."},
		{"9713[ana.o"},
		{"9713[ana.oa"},
		{"9713[ana.oan"},
		{"9713[ana.oan("},
		{"9713[ana.oan()"},
		{"9713[ana.oan()."},
		{"9713[ana.oan().h"},
		{"9713[ana.oan().hs"},
		{"9713[ana.oan().hst"},
		{"9713[ana.oan().hst("},
		{"9713[ana.oan().hst()"},
		{"9713[ana.oan().hst().eurUsd().s1().bids().fst() ana.oan().hst().eurUsd().s1().bids().lst() ana.oan().hst().eurUsd().s1().bids().sum()]"},
		{"Z["},
		{"Z[a"},
		{"Z[an"},
		{"Z[ana"},
		{"Z[ana."},
		{"Z[ana.o"},
		{"Z[ana.oa"},
		{"Z[ana.oan"},
		{"Z[ana.oan("},
		{"Z[ana.oan()"},
		{"Z[ana.oan()."},
		{"Z[ana.oan().h"},
		{"Z[ana.oan().hs"},
		{"Z[ana.oan().hst"},
		{"Z[ana.oan().hst("},
		{"Z[ana.oan().hst()"},
		{"Z[ana.oan().hst().eurUsd().s1().bids().fst() ana.oan().hst().eurUsd().s1().bids().lst() ana.oan().hst().eurUsd().s1().bids().sum()]"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.ObjsLit()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestAsnValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"asn", 3},
		{"asn ", 3},
		{"asn\n", 3},
		{"asn.", 3},
		{"asn:", 3},
		{"asn,", 3},
		{"asn\"", 3},
		{"asn(", 3},
		{"asn)", 3},
		{"asn[", 3},
		{"asn]", 3},
		{"asn// comment", 3},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Asn()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestAsnInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_a"},
		{"a_"},
		{"_as"},
		{"as_"},
		{"_asn"},
		{"asn_"},
		{"9713a"},
		{"a9713"},
		{"9713as"},
		{"as9713"},
		{"9713asn"},
		{"asn9713"},
		{"Za"},
		{"aZ"},
		{"Zas"},
		{"asZ"},
		{"Zasn"},
		{"asnZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Asn()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestEachValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"each", 4},
		{"each ", 4},
		{"each\n", 4},
		{"each.", 4},
		{"each:", 4},
		{"each,", 4},
		{"each\"", 4},
		{"each(", 4},
		{"each)", 4},
		{"each[", 4},
		{"each]", 4},
		{"each// comment", 4},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Each()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestEachInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_e"},
		{"e_"},
		{"_ea"},
		{"ea_"},
		{"_eac"},
		{"eac_"},
		{"_each"},
		{"each_"},
		{"9713e"},
		{"e9713"},
		{"9713ea"},
		{"ea9713"},
		{"9713eac"},
		{"eac9713"},
		{"9713each"},
		{"each9713"},
		{"Ze"},
		{"eZ"},
		{"Zea"},
		{"eaZ"},
		{"Zeac"},
		{"eacZ"},
		{"Zeach"},
		{"eachZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Each()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestPllEachValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"pllEach", 7},
		{"pllEach ", 7},
		{"pllEach\n", 7},
		{"pllEach.", 7},
		{"pllEach:", 7},
		{"pllEach,", 7},
		{"pllEach\"", 7},
		{"pllEach(", 7},
		{"pllEach)", 7},
		{"pllEach[", 7},
		{"pllEach]", 7},
		{"pllEach// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.PllEach()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestPllEachInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_p"},
		{"p_"},
		{"_pl"},
		{"pl_"},
		{"_pll"},
		{"pll_"},
		{"_pllE"},
		{"pllE_"},
		{"_pllEa"},
		{"pllEa_"},
		{"_pllEac"},
		{"pllEac_"},
		{"_pllEach"},
		{"pllEach_"},
		{"9713p"},
		{"p9713"},
		{"9713pl"},
		{"pl9713"},
		{"9713pll"},
		{"pll9713"},
		{"9713pllE"},
		{"pllE9713"},
		{"9713pllEa"},
		{"pllEa9713"},
		{"9713pllEac"},
		{"pllEac9713"},
		{"9713pllEach"},
		{"pllEach9713"},
		{"Zp"},
		{"pZ"},
		{"Zpl"},
		{"plZ"},
		{"Zpll"},
		{"pllZ"},
		{"ZpllE"},
		{"pllEZ"},
		{"ZpllEa"},
		{"pllEaZ"},
		{"ZpllEac"},
		{"pllEacZ"},
		{"ZpllEach"},
		{"pllEachZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.PllEach()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestPllWaitValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"pllWait", 7},
		{"pllWait ", 7},
		{"pllWait\n", 7},
		{"pllWait.", 7},
		{"pllWait:", 7},
		{"pllWait,", 7},
		{"pllWait\"", 7},
		{"pllWait(", 7},
		{"pllWait)", 7},
		{"pllWait[", 7},
		{"pllWait]", 7},
		{"pllWait// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.PllWait()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestPllWaitInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_p"},
		{"p_"},
		{"_pl"},
		{"pl_"},
		{"_pll"},
		{"pll_"},
		{"_pllW"},
		{"pllW_"},
		{"_pllWa"},
		{"pllWa_"},
		{"_pllWai"},
		{"pllWai_"},
		{"_pllWait"},
		{"pllWait_"},
		{"9713p"},
		{"p9713"},
		{"9713pl"},
		{"pl9713"},
		{"9713pll"},
		{"pll9713"},
		{"9713pllW"},
		{"pllW9713"},
		{"9713pllWa"},
		{"pllWa9713"},
		{"9713pllWai"},
		{"pllWai9713"},
		{"9713pllWait"},
		{"pllWait9713"},
		{"Zp"},
		{"pZ"},
		{"Zpl"},
		{"plZ"},
		{"Zpll"},
		{"pllZ"},
		{"ZpllW"},
		{"pllWZ"},
		{"ZpllWa"},
		{"pllWaZ"},
		{"ZpllWai"},
		{"pllWaiZ"},
		{"ZpllWait"},
		{"pllWaitZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.PllWait()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestThenValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"then", 4},
		{"then ", 4},
		{"then\n", 4},
		{"then.", 4},
		{"then:", 4},
		{"then,", 4},
		{"then\"", 4},
		{"then(", 4},
		{"then)", 4},
		{"then[", 4},
		{"then]", 4},
		{"then// comment", 4},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Then()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestThenInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_t"},
		{"t_"},
		{"_th"},
		{"th_"},
		{"_the"},
		{"the_"},
		{"_then"},
		{"then_"},
		{"9713t"},
		{"t9713"},
		{"9713th"},
		{"th9713"},
		{"9713the"},
		{"the9713"},
		{"9713then"},
		{"then9713"},
		{"Zt"},
		{"tZ"},
		{"Zth"},
		{"thZ"},
		{"Zthe"},
		{"theZ"},
		{"Zthen"},
		{"thenZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Then()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestElseValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"else", 4},
		{"else ", 4},
		{"else\n", 4},
		{"else.", 4},
		{"else:", 4},
		{"else,", 4},
		{"else\"", 4},
		{"else(", 4},
		{"else)", 4},
		{"else[", 4},
		{"else]", 4},
		{"else// comment", 4},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Else()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestElseInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_e"},
		{"e_"},
		{"_el"},
		{"el_"},
		{"_els"},
		{"els_"},
		{"_else"},
		{"else_"},
		{"9713e"},
		{"e9713"},
		{"9713el"},
		{"el9713"},
		{"9713els"},
		{"els9713"},
		{"9713else"},
		{"else9713"},
		{"Ze"},
		{"eZ"},
		{"Zel"},
		{"elZ"},
		{"Zels"},
		{"elsZ"},
		{"Zelse"},
		{"elseZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Else()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestPnlPctValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"pnlPct", 6},
		{"pnlPct ", 6},
		{"pnlPct\n", 6},
		{"pnlPct.", 6},
		{"pnlPct:", 6},
		{"pnlPct,", 6},
		{"pnlPct\"", 6},
		{"pnlPct(", 6},
		{"pnlPct)", 6},
		{"pnlPct[", 6},
		{"pnlPct]", 6},
		{"pnlPct// comment", 6},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.PnlPct()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestPnlPctInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_p"},
		{"p_"},
		{"_pn"},
		{"pn_"},
		{"_pnl"},
		{"pnl_"},
		{"_pnlP"},
		{"pnlP_"},
		{"_pnlPc"},
		{"pnlPc_"},
		{"_pnlPct"},
		{"pnlPct_"},
		{"9713p"},
		{"p9713"},
		{"9713pn"},
		{"pn9713"},
		{"9713pnl"},
		{"pnl9713"},
		{"9713pnlP"},
		{"pnlP9713"},
		{"9713pnlPc"},
		{"pnlPc9713"},
		{"9713pnlPct"},
		{"pnlPct9713"},
		{"Zp"},
		{"pZ"},
		{"Zpn"},
		{"pnZ"},
		{"Zpnl"},
		{"pnlZ"},
		{"ZpnlP"},
		{"pnlPZ"},
		{"ZpnlPc"},
		{"pnlPcZ"},
		{"ZpnlPct"},
		{"pnlPctZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.PnlPct()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestScsPctValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"scsPct", 6},
		{"scsPct ", 6},
		{"scsPct\n", 6},
		{"scsPct.", 6},
		{"scsPct:", 6},
		{"scsPct,", 6},
		{"scsPct\"", 6},
		{"scsPct(", 6},
		{"scsPct)", 6},
		{"scsPct[", 6},
		{"scsPct]", 6},
		{"scsPct// comment", 6},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.ScsPct()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestScsPctInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_s"},
		{"s_"},
		{"_sc"},
		{"sc_"},
		{"_scs"},
		{"scs_"},
		{"_scsP"},
		{"scsP_"},
		{"_scsPc"},
		{"scsPc_"},
		{"_scsPct"},
		{"scsPct_"},
		{"9713s"},
		{"s9713"},
		{"9713sc"},
		{"sc9713"},
		{"9713scs"},
		{"scs9713"},
		{"9713scsP"},
		{"scsP9713"},
		{"9713scsPc"},
		{"scsPc9713"},
		{"9713scsPct"},
		{"scsPct9713"},
		{"Zs"},
		{"sZ"},
		{"Zsc"},
		{"scZ"},
		{"Zscs"},
		{"scsZ"},
		{"ZscsP"},
		{"scsPZ"},
		{"ZscsPc"},
		{"scsPcZ"},
		{"ZscsPct"},
		{"scsPctZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.ScsPct()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestPipPerDayValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"pipPerDay", 9},
		{"pipPerDay ", 9},
		{"pipPerDay\n", 9},
		{"pipPerDay.", 9},
		{"pipPerDay:", 9},
		{"pipPerDay,", 9},
		{"pipPerDay\"", 9},
		{"pipPerDay(", 9},
		{"pipPerDay)", 9},
		{"pipPerDay[", 9},
		{"pipPerDay]", 9},
		{"pipPerDay// comment", 9},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.PipPerDay()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestPipPerDayInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_p"},
		{"p_"},
		{"_pi"},
		{"pi_"},
		{"_pip"},
		{"pip_"},
		{"_pipP"},
		{"pipP_"},
		{"_pipPe"},
		{"pipPe_"},
		{"_pipPer"},
		{"pipPer_"},
		{"_pipPerD"},
		{"pipPerD_"},
		{"_pipPerDa"},
		{"pipPerDa_"},
		{"_pipPerDay"},
		{"pipPerDay_"},
		{"9713p"},
		{"p9713"},
		{"9713pi"},
		{"pi9713"},
		{"9713pip"},
		{"pip9713"},
		{"9713pipP"},
		{"pipP9713"},
		{"9713pipPe"},
		{"pipPe9713"},
		{"9713pipPer"},
		{"pipPer9713"},
		{"9713pipPerD"},
		{"pipPerD9713"},
		{"9713pipPerDa"},
		{"pipPerDa9713"},
		{"9713pipPerDay"},
		{"pipPerDay9713"},
		{"Zp"},
		{"pZ"},
		{"Zpi"},
		{"piZ"},
		{"Zpip"},
		{"pipZ"},
		{"ZpipP"},
		{"pipPZ"},
		{"ZpipPe"},
		{"pipPeZ"},
		{"ZpipPer"},
		{"pipPerZ"},
		{"ZpipPerD"},
		{"pipPerDZ"},
		{"ZpipPerDa"},
		{"pipPerDaZ"},
		{"ZpipPerDay"},
		{"pipPerDayZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.PipPerDay()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestUsdPerDayValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"usdPerDay", 9},
		{"usdPerDay ", 9},
		{"usdPerDay\n", 9},
		{"usdPerDay.", 9},
		{"usdPerDay:", 9},
		{"usdPerDay,", 9},
		{"usdPerDay\"", 9},
		{"usdPerDay(", 9},
		{"usdPerDay)", 9},
		{"usdPerDay[", 9},
		{"usdPerDay]", 9},
		{"usdPerDay// comment", 9},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.UsdPerDay()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestUsdPerDayInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_u"},
		{"u_"},
		{"_us"},
		{"us_"},
		{"_usd"},
		{"usd_"},
		{"_usdP"},
		{"usdP_"},
		{"_usdPe"},
		{"usdPe_"},
		{"_usdPer"},
		{"usdPer_"},
		{"_usdPerD"},
		{"usdPerD_"},
		{"_usdPerDa"},
		{"usdPerDa_"},
		{"_usdPerDay"},
		{"usdPerDay_"},
		{"9713u"},
		{"u9713"},
		{"9713us"},
		{"us9713"},
		{"9713usd"},
		{"usd9713"},
		{"9713usdP"},
		{"usdP9713"},
		{"9713usdPe"},
		{"usdPe9713"},
		{"9713usdPer"},
		{"usdPer9713"},
		{"9713usdPerD"},
		{"usdPerD9713"},
		{"9713usdPerDa"},
		{"usdPerDa9713"},
		{"9713usdPerDay"},
		{"usdPerDay9713"},
		{"Zu"},
		{"uZ"},
		{"Zus"},
		{"usZ"},
		{"Zusd"},
		{"usdZ"},
		{"ZusdP"},
		{"usdPZ"},
		{"ZusdPe"},
		{"usdPeZ"},
		{"ZusdPer"},
		{"usdPerZ"},
		{"ZusdPerD"},
		{"usdPerDZ"},
		{"ZusdPerDa"},
		{"usdPerDaZ"},
		{"ZusdPerDay"},
		{"usdPerDayZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.UsdPerDay()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestScsPerDayValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"scsPerDay", 9},
		{"scsPerDay ", 9},
		{"scsPerDay\n", 9},
		{"scsPerDay.", 9},
		{"scsPerDay:", 9},
		{"scsPerDay,", 9},
		{"scsPerDay\"", 9},
		{"scsPerDay(", 9},
		{"scsPerDay)", 9},
		{"scsPerDay[", 9},
		{"scsPerDay]", 9},
		{"scsPerDay// comment", 9},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.ScsPerDay()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestScsPerDayInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_s"},
		{"s_"},
		{"_sc"},
		{"sc_"},
		{"_scs"},
		{"scs_"},
		{"_scsP"},
		{"scsP_"},
		{"_scsPe"},
		{"scsPe_"},
		{"_scsPer"},
		{"scsPer_"},
		{"_scsPerD"},
		{"scsPerD_"},
		{"_scsPerDa"},
		{"scsPerDa_"},
		{"_scsPerDay"},
		{"scsPerDay_"},
		{"9713s"},
		{"s9713"},
		{"9713sc"},
		{"sc9713"},
		{"9713scs"},
		{"scs9713"},
		{"9713scsP"},
		{"scsP9713"},
		{"9713scsPe"},
		{"scsPe9713"},
		{"9713scsPer"},
		{"scsPer9713"},
		{"9713scsPerD"},
		{"scsPerD9713"},
		{"9713scsPerDa"},
		{"scsPerDa9713"},
		{"9713scsPerDay"},
		{"scsPerDay9713"},
		{"Zs"},
		{"sZ"},
		{"Zsc"},
		{"scZ"},
		{"Zscs"},
		{"scsZ"},
		{"ZscsP"},
		{"scsPZ"},
		{"ZscsPe"},
		{"scsPeZ"},
		{"ZscsPer"},
		{"scsPerZ"},
		{"ZscsPerD"},
		{"scsPerDZ"},
		{"ZscsPerDa"},
		{"scsPerDaZ"},
		{"ZscsPerDay"},
		{"scsPerDayZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.ScsPerDay()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestOpnPerDayValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"opnPerDay", 9},
		{"opnPerDay ", 9},
		{"opnPerDay\n", 9},
		{"opnPerDay.", 9},
		{"opnPerDay:", 9},
		{"opnPerDay,", 9},
		{"opnPerDay\"", 9},
		{"opnPerDay(", 9},
		{"opnPerDay)", 9},
		{"opnPerDay[", 9},
		{"opnPerDay]", 9},
		{"opnPerDay// comment", 9},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.OpnPerDay()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestOpnPerDayInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_o"},
		{"o_"},
		{"_op"},
		{"op_"},
		{"_opn"},
		{"opn_"},
		{"_opnP"},
		{"opnP_"},
		{"_opnPe"},
		{"opnPe_"},
		{"_opnPer"},
		{"opnPer_"},
		{"_opnPerD"},
		{"opnPerD_"},
		{"_opnPerDa"},
		{"opnPerDa_"},
		{"_opnPerDay"},
		{"opnPerDay_"},
		{"9713o"},
		{"o9713"},
		{"9713op"},
		{"op9713"},
		{"9713opn"},
		{"opn9713"},
		{"9713opnP"},
		{"opnP9713"},
		{"9713opnPe"},
		{"opnPe9713"},
		{"9713opnPer"},
		{"opnPer9713"},
		{"9713opnPerD"},
		{"opnPerD9713"},
		{"9713opnPerDa"},
		{"opnPerDa9713"},
		{"9713opnPerDay"},
		{"opnPerDay9713"},
		{"Zo"},
		{"oZ"},
		{"Zop"},
		{"opZ"},
		{"Zopn"},
		{"opnZ"},
		{"ZopnP"},
		{"opnPZ"},
		{"ZopnPe"},
		{"opnPeZ"},
		{"ZopnPer"},
		{"opnPerZ"},
		{"ZopnPerD"},
		{"opnPerDZ"},
		{"ZopnPerDa"},
		{"opnPerDaZ"},
		{"ZopnPerDay"},
		{"opnPerDayZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.OpnPerDay()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestPnlUsdValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"pnlUsd", 6},
		{"pnlUsd ", 6},
		{"pnlUsd\n", 6},
		{"pnlUsd.", 6},
		{"pnlUsd:", 6},
		{"pnlUsd,", 6},
		{"pnlUsd\"", 6},
		{"pnlUsd(", 6},
		{"pnlUsd)", 6},
		{"pnlUsd[", 6},
		{"pnlUsd]", 6},
		{"pnlUsd// comment", 6},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.PnlUsd()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestPnlUsdInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_p"},
		{"p_"},
		{"_pn"},
		{"pn_"},
		{"_pnl"},
		{"pnl_"},
		{"_pnlU"},
		{"pnlU_"},
		{"_pnlUs"},
		{"pnlUs_"},
		{"_pnlUsd"},
		{"pnlUsd_"},
		{"9713p"},
		{"p9713"},
		{"9713pn"},
		{"pn9713"},
		{"9713pnl"},
		{"pnl9713"},
		{"9713pnlU"},
		{"pnlU9713"},
		{"9713pnlUs"},
		{"pnlUs9713"},
		{"9713pnlUsd"},
		{"pnlUsd9713"},
		{"Zp"},
		{"pZ"},
		{"Zpn"},
		{"pnZ"},
		{"Zpnl"},
		{"pnlZ"},
		{"ZpnlU"},
		{"pnlUZ"},
		{"ZpnlUs"},
		{"pnlUsZ"},
		{"ZpnlUsd"},
		{"pnlUsdZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.PnlUsd()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestPipAvgValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"pipAvg", 6},
		{"pipAvg ", 6},
		{"pipAvg\n", 6},
		{"pipAvg.", 6},
		{"pipAvg:", 6},
		{"pipAvg,", 6},
		{"pipAvg\"", 6},
		{"pipAvg(", 6},
		{"pipAvg)", 6},
		{"pipAvg[", 6},
		{"pipAvg]", 6},
		{"pipAvg// comment", 6},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.PipAvg()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestPipAvgInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_p"},
		{"p_"},
		{"_pi"},
		{"pi_"},
		{"_pip"},
		{"pip_"},
		{"_pipA"},
		{"pipA_"},
		{"_pipAv"},
		{"pipAv_"},
		{"_pipAvg"},
		{"pipAvg_"},
		{"9713p"},
		{"p9713"},
		{"9713pi"},
		{"pi9713"},
		{"9713pip"},
		{"pip9713"},
		{"9713pipA"},
		{"pipA9713"},
		{"9713pipAv"},
		{"pipAv9713"},
		{"9713pipAvg"},
		{"pipAvg9713"},
		{"Zp"},
		{"pZ"},
		{"Zpi"},
		{"piZ"},
		{"Zpip"},
		{"pipZ"},
		{"ZpipA"},
		{"pipAZ"},
		{"ZpipAv"},
		{"pipAvZ"},
		{"ZpipAvg"},
		{"pipAvgZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.PipAvg()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestPipMdnValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"pipMdn", 6},
		{"pipMdn ", 6},
		{"pipMdn\n", 6},
		{"pipMdn.", 6},
		{"pipMdn:", 6},
		{"pipMdn,", 6},
		{"pipMdn\"", 6},
		{"pipMdn(", 6},
		{"pipMdn)", 6},
		{"pipMdn[", 6},
		{"pipMdn]", 6},
		{"pipMdn// comment", 6},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.PipMdn()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestPipMdnInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_p"},
		{"p_"},
		{"_pi"},
		{"pi_"},
		{"_pip"},
		{"pip_"},
		{"_pipM"},
		{"pipM_"},
		{"_pipMd"},
		{"pipMd_"},
		{"_pipMdn"},
		{"pipMdn_"},
		{"9713p"},
		{"p9713"},
		{"9713pi"},
		{"pi9713"},
		{"9713pip"},
		{"pip9713"},
		{"9713pipM"},
		{"pipM9713"},
		{"9713pipMd"},
		{"pipMd9713"},
		{"9713pipMdn"},
		{"pipMdn9713"},
		{"Zp"},
		{"pZ"},
		{"Zpi"},
		{"piZ"},
		{"Zpip"},
		{"pipZ"},
		{"ZpipM"},
		{"pipMZ"},
		{"ZpipMd"},
		{"pipMdZ"},
		{"ZpipMdn"},
		{"pipMdnZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.PipMdn()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestPipMinValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"pipMin", 6},
		{"pipMin ", 6},
		{"pipMin\n", 6},
		{"pipMin.", 6},
		{"pipMin:", 6},
		{"pipMin,", 6},
		{"pipMin\"", 6},
		{"pipMin(", 6},
		{"pipMin)", 6},
		{"pipMin[", 6},
		{"pipMin]", 6},
		{"pipMin// comment", 6},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.PipMin()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestPipMinInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_p"},
		{"p_"},
		{"_pi"},
		{"pi_"},
		{"_pip"},
		{"pip_"},
		{"_pipM"},
		{"pipM_"},
		{"_pipMi"},
		{"pipMi_"},
		{"_pipMin"},
		{"pipMin_"},
		{"9713p"},
		{"p9713"},
		{"9713pi"},
		{"pi9713"},
		{"9713pip"},
		{"pip9713"},
		{"9713pipM"},
		{"pipM9713"},
		{"9713pipMi"},
		{"pipMi9713"},
		{"9713pipMin"},
		{"pipMin9713"},
		{"Zp"},
		{"pZ"},
		{"Zpi"},
		{"piZ"},
		{"Zpip"},
		{"pipZ"},
		{"ZpipM"},
		{"pipMZ"},
		{"ZpipMi"},
		{"pipMiZ"},
		{"ZpipMin"},
		{"pipMinZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.PipMin()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestPipMaxValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"pipMax", 6},
		{"pipMax ", 6},
		{"pipMax\n", 6},
		{"pipMax.", 6},
		{"pipMax:", 6},
		{"pipMax,", 6},
		{"pipMax\"", 6},
		{"pipMax(", 6},
		{"pipMax)", 6},
		{"pipMax[", 6},
		{"pipMax]", 6},
		{"pipMax// comment", 6},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.PipMax()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestPipMaxInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_p"},
		{"p_"},
		{"_pi"},
		{"pi_"},
		{"_pip"},
		{"pip_"},
		{"_pipM"},
		{"pipM_"},
		{"_pipMa"},
		{"pipMa_"},
		{"_pipMax"},
		{"pipMax_"},
		{"9713p"},
		{"p9713"},
		{"9713pi"},
		{"pi9713"},
		{"9713pip"},
		{"pip9713"},
		{"9713pipM"},
		{"pipM9713"},
		{"9713pipMa"},
		{"pipMa9713"},
		{"9713pipMax"},
		{"pipMax9713"},
		{"Zp"},
		{"pZ"},
		{"Zpi"},
		{"piZ"},
		{"Zpip"},
		{"pipZ"},
		{"ZpipM"},
		{"pipMZ"},
		{"ZpipMa"},
		{"pipMaZ"},
		{"ZpipMax"},
		{"pipMaxZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.PipMax()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestPipSumValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"pipSum", 6},
		{"pipSum ", 6},
		{"pipSum\n", 6},
		{"pipSum.", 6},
		{"pipSum:", 6},
		{"pipSum,", 6},
		{"pipSum\"", 6},
		{"pipSum(", 6},
		{"pipSum)", 6},
		{"pipSum[", 6},
		{"pipSum]", 6},
		{"pipSum// comment", 6},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.PipSum()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestPipSumInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_p"},
		{"p_"},
		{"_pi"},
		{"pi_"},
		{"_pip"},
		{"pip_"},
		{"_pipS"},
		{"pipS_"},
		{"_pipSu"},
		{"pipSu_"},
		{"_pipSum"},
		{"pipSum_"},
		{"9713p"},
		{"p9713"},
		{"9713pi"},
		{"pi9713"},
		{"9713pip"},
		{"pip9713"},
		{"9713pipS"},
		{"pipS9713"},
		{"9713pipSu"},
		{"pipSu9713"},
		{"9713pipSum"},
		{"pipSum9713"},
		{"Zp"},
		{"pZ"},
		{"Zpi"},
		{"piZ"},
		{"Zpip"},
		{"pipZ"},
		{"ZpipS"},
		{"pipSZ"},
		{"ZpipSu"},
		{"pipSuZ"},
		{"ZpipSum"},
		{"pipSumZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.PipSum()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestDurAvgValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"durAvg", 6},
		{"durAvg ", 6},
		{"durAvg\n", 6},
		{"durAvg.", 6},
		{"durAvg:", 6},
		{"durAvg,", 6},
		{"durAvg\"", 6},
		{"durAvg(", 6},
		{"durAvg)", 6},
		{"durAvg[", 6},
		{"durAvg]", 6},
		{"durAvg// comment", 6},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.DurAvg()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestDurAvgInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_d"},
		{"d_"},
		{"_du"},
		{"du_"},
		{"_dur"},
		{"dur_"},
		{"_durA"},
		{"durA_"},
		{"_durAv"},
		{"durAv_"},
		{"_durAvg"},
		{"durAvg_"},
		{"9713d"},
		{"d9713"},
		{"9713du"},
		{"du9713"},
		{"9713dur"},
		{"dur9713"},
		{"9713durA"},
		{"durA9713"},
		{"9713durAv"},
		{"durAv9713"},
		{"9713durAvg"},
		{"durAvg9713"},
		{"Zd"},
		{"dZ"},
		{"Zdu"},
		{"duZ"},
		{"Zdur"},
		{"durZ"},
		{"ZdurA"},
		{"durAZ"},
		{"ZdurAv"},
		{"durAvZ"},
		{"ZdurAvg"},
		{"durAvgZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.DurAvg()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestDurMdnValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"durMdn", 6},
		{"durMdn ", 6},
		{"durMdn\n", 6},
		{"durMdn.", 6},
		{"durMdn:", 6},
		{"durMdn,", 6},
		{"durMdn\"", 6},
		{"durMdn(", 6},
		{"durMdn)", 6},
		{"durMdn[", 6},
		{"durMdn]", 6},
		{"durMdn// comment", 6},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.DurMdn()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestDurMdnInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_d"},
		{"d_"},
		{"_du"},
		{"du_"},
		{"_dur"},
		{"dur_"},
		{"_durM"},
		{"durM_"},
		{"_durMd"},
		{"durMd_"},
		{"_durMdn"},
		{"durMdn_"},
		{"9713d"},
		{"d9713"},
		{"9713du"},
		{"du9713"},
		{"9713dur"},
		{"dur9713"},
		{"9713durM"},
		{"durM9713"},
		{"9713durMd"},
		{"durMd9713"},
		{"9713durMdn"},
		{"durMdn9713"},
		{"Zd"},
		{"dZ"},
		{"Zdu"},
		{"duZ"},
		{"Zdur"},
		{"durZ"},
		{"ZdurM"},
		{"durMZ"},
		{"ZdurMd"},
		{"durMdZ"},
		{"ZdurMdn"},
		{"durMdnZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.DurMdn()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestDurMinValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"durMin", 6},
		{"durMin ", 6},
		{"durMin\n", 6},
		{"durMin.", 6},
		{"durMin:", 6},
		{"durMin,", 6},
		{"durMin\"", 6},
		{"durMin(", 6},
		{"durMin)", 6},
		{"durMin[", 6},
		{"durMin]", 6},
		{"durMin// comment", 6},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.DurMin()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestDurMinInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_d"},
		{"d_"},
		{"_du"},
		{"du_"},
		{"_dur"},
		{"dur_"},
		{"_durM"},
		{"durM_"},
		{"_durMi"},
		{"durMi_"},
		{"_durMin"},
		{"durMin_"},
		{"9713d"},
		{"d9713"},
		{"9713du"},
		{"du9713"},
		{"9713dur"},
		{"dur9713"},
		{"9713durM"},
		{"durM9713"},
		{"9713durMi"},
		{"durMi9713"},
		{"9713durMin"},
		{"durMin9713"},
		{"Zd"},
		{"dZ"},
		{"Zdu"},
		{"duZ"},
		{"Zdur"},
		{"durZ"},
		{"ZdurM"},
		{"durMZ"},
		{"ZdurMi"},
		{"durMiZ"},
		{"ZdurMin"},
		{"durMinZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.DurMin()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestDurMaxValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"durMax", 6},
		{"durMax ", 6},
		{"durMax\n", 6},
		{"durMax.", 6},
		{"durMax:", 6},
		{"durMax,", 6},
		{"durMax\"", 6},
		{"durMax(", 6},
		{"durMax)", 6},
		{"durMax[", 6},
		{"durMax]", 6},
		{"durMax// comment", 6},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.DurMax()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestDurMaxInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_d"},
		{"d_"},
		{"_du"},
		{"du_"},
		{"_dur"},
		{"dur_"},
		{"_durM"},
		{"durM_"},
		{"_durMa"},
		{"durMa_"},
		{"_durMax"},
		{"durMax_"},
		{"9713d"},
		{"d9713"},
		{"9713du"},
		{"du9713"},
		{"9713dur"},
		{"dur9713"},
		{"9713durM"},
		{"durM9713"},
		{"9713durMa"},
		{"durMa9713"},
		{"9713durMax"},
		{"durMax9713"},
		{"Zd"},
		{"dZ"},
		{"Zdu"},
		{"duZ"},
		{"Zdur"},
		{"durZ"},
		{"ZdurM"},
		{"durMZ"},
		{"ZdurMa"},
		{"durMaZ"},
		{"ZdurMax"},
		{"durMaxZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.DurMax()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestLosLimMaxValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"losLimMax", 9},
		{"losLimMax ", 9},
		{"losLimMax\n", 9},
		{"losLimMax.", 9},
		{"losLimMax:", 9},
		{"losLimMax,", 9},
		{"losLimMax\"", 9},
		{"losLimMax(", 9},
		{"losLimMax)", 9},
		{"losLimMax[", 9},
		{"losLimMax]", 9},
		{"losLimMax// comment", 9},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.LosLimMax()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestLosLimMaxInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_l"},
		{"l_"},
		{"_lo"},
		{"lo_"},
		{"_los"},
		{"los_"},
		{"_losL"},
		{"losL_"},
		{"_losLi"},
		{"losLi_"},
		{"_losLim"},
		{"losLim_"},
		{"_losLimM"},
		{"losLimM_"},
		{"_losLimMa"},
		{"losLimMa_"},
		{"_losLimMax"},
		{"losLimMax_"},
		{"9713l"},
		{"l9713"},
		{"9713lo"},
		{"lo9713"},
		{"9713los"},
		{"los9713"},
		{"9713losL"},
		{"losL9713"},
		{"9713losLi"},
		{"losLi9713"},
		{"9713losLim"},
		{"losLim9713"},
		{"9713losLimM"},
		{"losLimM9713"},
		{"9713losLimMa"},
		{"losLimMa9713"},
		{"9713losLimMax"},
		{"losLimMax9713"},
		{"Zl"},
		{"lZ"},
		{"Zlo"},
		{"loZ"},
		{"Zlos"},
		{"losZ"},
		{"ZlosL"},
		{"losLZ"},
		{"ZlosLi"},
		{"losLiZ"},
		{"ZlosLim"},
		{"losLimZ"},
		{"ZlosLimM"},
		{"losLimMZ"},
		{"ZlosLimMa"},
		{"losLimMaZ"},
		{"ZlosLimMax"},
		{"losLimMaxZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.LosLimMax()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestDurLimMaxValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"durLimMax", 9},
		{"durLimMax ", 9},
		{"durLimMax\n", 9},
		{"durLimMax.", 9},
		{"durLimMax:", 9},
		{"durLimMax,", 9},
		{"durLimMax\"", 9},
		{"durLimMax(", 9},
		{"durLimMax)", 9},
		{"durLimMax[", 9},
		{"durLimMax]", 9},
		{"durLimMax// comment", 9},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.DurLimMax()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestDurLimMaxInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_d"},
		{"d_"},
		{"_du"},
		{"du_"},
		{"_dur"},
		{"dur_"},
		{"_durL"},
		{"durL_"},
		{"_durLi"},
		{"durLi_"},
		{"_durLim"},
		{"durLim_"},
		{"_durLimM"},
		{"durLimM_"},
		{"_durLimMa"},
		{"durLimMa_"},
		{"_durLimMax"},
		{"durLimMax_"},
		{"9713d"},
		{"d9713"},
		{"9713du"},
		{"du9713"},
		{"9713dur"},
		{"dur9713"},
		{"9713durL"},
		{"durL9713"},
		{"9713durLi"},
		{"durLi9713"},
		{"9713durLim"},
		{"durLim9713"},
		{"9713durLimM"},
		{"durLimM9713"},
		{"9713durLimMa"},
		{"durLimMa9713"},
		{"9713durLimMax"},
		{"durLimMax9713"},
		{"Zd"},
		{"dZ"},
		{"Zdu"},
		{"duZ"},
		{"Zdur"},
		{"durZ"},
		{"ZdurL"},
		{"durLZ"},
		{"ZdurLi"},
		{"durLiZ"},
		{"ZdurLim"},
		{"durLimZ"},
		{"ZdurLimM"},
		{"durLimMZ"},
		{"ZdurLimMa"},
		{"durLimMaZ"},
		{"ZdurLimMax"},
		{"durLimMaxZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.DurLimMax()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestDayCntValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"dayCnt", 6},
		{"dayCnt ", 6},
		{"dayCnt\n", 6},
		{"dayCnt.", 6},
		{"dayCnt:", 6},
		{"dayCnt,", 6},
		{"dayCnt\"", 6},
		{"dayCnt(", 6},
		{"dayCnt)", 6},
		{"dayCnt[", 6},
		{"dayCnt]", 6},
		{"dayCnt// comment", 6},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.DayCnt()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestDayCntInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_d"},
		{"d_"},
		{"_da"},
		{"da_"},
		{"_day"},
		{"day_"},
		{"_dayC"},
		{"dayC_"},
		{"_dayCn"},
		{"dayCn_"},
		{"_dayCnt"},
		{"dayCnt_"},
		{"9713d"},
		{"d9713"},
		{"9713da"},
		{"da9713"},
		{"9713day"},
		{"day9713"},
		{"9713dayC"},
		{"dayC9713"},
		{"9713dayCn"},
		{"dayCn9713"},
		{"9713dayCnt"},
		{"dayCnt9713"},
		{"Zd"},
		{"dZ"},
		{"Zda"},
		{"daZ"},
		{"Zday"},
		{"dayZ"},
		{"ZdayC"},
		{"dayCZ"},
		{"ZdayCn"},
		{"dayCnZ"},
		{"ZdayCnt"},
		{"dayCntZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.DayCnt()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestTrdCntValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"trdCnt", 6},
		{"trdCnt ", 6},
		{"trdCnt\n", 6},
		{"trdCnt.", 6},
		{"trdCnt:", 6},
		{"trdCnt,", 6},
		{"trdCnt\"", 6},
		{"trdCnt(", 6},
		{"trdCnt)", 6},
		{"trdCnt[", 6},
		{"trdCnt]", 6},
		{"trdCnt// comment", 6},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.TrdCnt()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestTrdCntInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_t"},
		{"t_"},
		{"_tr"},
		{"tr_"},
		{"_trd"},
		{"trd_"},
		{"_trdC"},
		{"trdC_"},
		{"_trdCn"},
		{"trdCn_"},
		{"_trdCnt"},
		{"trdCnt_"},
		{"9713t"},
		{"t9713"},
		{"9713tr"},
		{"tr9713"},
		{"9713trd"},
		{"trd9713"},
		{"9713trdC"},
		{"trdC9713"},
		{"9713trdCn"},
		{"trdCn9713"},
		{"9713trdCnt"},
		{"trdCnt9713"},
		{"Zt"},
		{"tZ"},
		{"Ztr"},
		{"trZ"},
		{"Ztrd"},
		{"trdZ"},
		{"ZtrdC"},
		{"trdCZ"},
		{"ZtrdCn"},
		{"trdCnZ"},
		{"ZtrdCnt"},
		{"trdCntZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.TrdCnt()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestTrdPctValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"trdPct", 6},
		{"trdPct ", 6},
		{"trdPct\n", 6},
		{"trdPct.", 6},
		{"trdPct:", 6},
		{"trdPct,", 6},
		{"trdPct\"", 6},
		{"trdPct(", 6},
		{"trdPct)", 6},
		{"trdPct[", 6},
		{"trdPct]", 6},
		{"trdPct// comment", 6},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.TrdPct()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestTrdPctInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_t"},
		{"t_"},
		{"_tr"},
		{"tr_"},
		{"_trd"},
		{"trd_"},
		{"_trdP"},
		{"trdP_"},
		{"_trdPc"},
		{"trdPc_"},
		{"_trdPct"},
		{"trdPct_"},
		{"9713t"},
		{"t9713"},
		{"9713tr"},
		{"tr9713"},
		{"9713trd"},
		{"trd9713"},
		{"9713trdP"},
		{"trdP9713"},
		{"9713trdPc"},
		{"trdPc9713"},
		{"9713trdPct"},
		{"trdPct9713"},
		{"Zt"},
		{"tZ"},
		{"Ztr"},
		{"trZ"},
		{"Ztrd"},
		{"trdZ"},
		{"ZtrdP"},
		{"trdPZ"},
		{"ZtrdPc"},
		{"trdPcZ"},
		{"ZtrdPct"},
		{"trdPctZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.TrdPct()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestCstTotUsdValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"cstTotUsd", 9},
		{"cstTotUsd ", 9},
		{"cstTotUsd\n", 9},
		{"cstTotUsd.", 9},
		{"cstTotUsd:", 9},
		{"cstTotUsd,", 9},
		{"cstTotUsd\"", 9},
		{"cstTotUsd(", 9},
		{"cstTotUsd)", 9},
		{"cstTotUsd[", 9},
		{"cstTotUsd]", 9},
		{"cstTotUsd// comment", 9},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.CstTotUsd()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestCstTotUsdInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_c"},
		{"c_"},
		{"_cs"},
		{"cs_"},
		{"_cst"},
		{"cst_"},
		{"_cstT"},
		{"cstT_"},
		{"_cstTo"},
		{"cstTo_"},
		{"_cstTot"},
		{"cstTot_"},
		{"_cstTotU"},
		{"cstTotU_"},
		{"_cstTotUs"},
		{"cstTotUs_"},
		{"_cstTotUsd"},
		{"cstTotUsd_"},
		{"9713c"},
		{"c9713"},
		{"9713cs"},
		{"cs9713"},
		{"9713cst"},
		{"cst9713"},
		{"9713cstT"},
		{"cstT9713"},
		{"9713cstTo"},
		{"cstTo9713"},
		{"9713cstTot"},
		{"cstTot9713"},
		{"9713cstTotU"},
		{"cstTotU9713"},
		{"9713cstTotUs"},
		{"cstTotUs9713"},
		{"9713cstTotUsd"},
		{"cstTotUsd9713"},
		{"Zc"},
		{"cZ"},
		{"Zcs"},
		{"csZ"},
		{"Zcst"},
		{"cstZ"},
		{"ZcstT"},
		{"cstTZ"},
		{"ZcstTo"},
		{"cstToZ"},
		{"ZcstTot"},
		{"cstTotZ"},
		{"ZcstTotU"},
		{"cstTotUZ"},
		{"ZcstTotUs"},
		{"cstTotUsZ"},
		{"ZcstTotUsd"},
		{"cstTotUsdZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.CstTotUsd()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestCstSpdUsdValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"cstSpdUsd", 9},
		{"cstSpdUsd ", 9},
		{"cstSpdUsd\n", 9},
		{"cstSpdUsd.", 9},
		{"cstSpdUsd:", 9},
		{"cstSpdUsd,", 9},
		{"cstSpdUsd\"", 9},
		{"cstSpdUsd(", 9},
		{"cstSpdUsd)", 9},
		{"cstSpdUsd[", 9},
		{"cstSpdUsd]", 9},
		{"cstSpdUsd// comment", 9},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.CstSpdUsd()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestCstSpdUsdInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_c"},
		{"c_"},
		{"_cs"},
		{"cs_"},
		{"_cst"},
		{"cst_"},
		{"_cstS"},
		{"cstS_"},
		{"_cstSp"},
		{"cstSp_"},
		{"_cstSpd"},
		{"cstSpd_"},
		{"_cstSpdU"},
		{"cstSpdU_"},
		{"_cstSpdUs"},
		{"cstSpdUs_"},
		{"_cstSpdUsd"},
		{"cstSpdUsd_"},
		{"9713c"},
		{"c9713"},
		{"9713cs"},
		{"cs9713"},
		{"9713cst"},
		{"cst9713"},
		{"9713cstS"},
		{"cstS9713"},
		{"9713cstSp"},
		{"cstSp9713"},
		{"9713cstSpd"},
		{"cstSpd9713"},
		{"9713cstSpdU"},
		{"cstSpdU9713"},
		{"9713cstSpdUs"},
		{"cstSpdUs9713"},
		{"9713cstSpdUsd"},
		{"cstSpdUsd9713"},
		{"Zc"},
		{"cZ"},
		{"Zcs"},
		{"csZ"},
		{"Zcst"},
		{"cstZ"},
		{"ZcstS"},
		{"cstSZ"},
		{"ZcstSp"},
		{"cstSpZ"},
		{"ZcstSpd"},
		{"cstSpdZ"},
		{"ZcstSpdU"},
		{"cstSpdUZ"},
		{"ZcstSpdUs"},
		{"cstSpdUsZ"},
		{"ZcstSpdUsd"},
		{"cstSpdUsdZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.CstSpdUsd()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestCstComUsdValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"cstComUsd", 9},
		{"cstComUsd ", 9},
		{"cstComUsd\n", 9},
		{"cstComUsd.", 9},
		{"cstComUsd:", 9},
		{"cstComUsd,", 9},
		{"cstComUsd\"", 9},
		{"cstComUsd(", 9},
		{"cstComUsd)", 9},
		{"cstComUsd[", 9},
		{"cstComUsd]", 9},
		{"cstComUsd// comment", 9},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.CstComUsd()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestCstComUsdInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_c"},
		{"c_"},
		{"_cs"},
		{"cs_"},
		{"_cst"},
		{"cst_"},
		{"_cstC"},
		{"cstC_"},
		{"_cstCo"},
		{"cstCo_"},
		{"_cstCom"},
		{"cstCom_"},
		{"_cstComU"},
		{"cstComU_"},
		{"_cstComUs"},
		{"cstComUs_"},
		{"_cstComUsd"},
		{"cstComUsd_"},
		{"9713c"},
		{"c9713"},
		{"9713cs"},
		{"cs9713"},
		{"9713cst"},
		{"cst9713"},
		{"9713cstC"},
		{"cstC9713"},
		{"9713cstCo"},
		{"cstCo9713"},
		{"9713cstCom"},
		{"cstCom9713"},
		{"9713cstComU"},
		{"cstComU9713"},
		{"9713cstComUs"},
		{"cstComUs9713"},
		{"9713cstComUsd"},
		{"cstComUsd9713"},
		{"Zc"},
		{"cZ"},
		{"Zcs"},
		{"csZ"},
		{"Zcst"},
		{"cstZ"},
		{"ZcstC"},
		{"cstCZ"},
		{"ZcstCo"},
		{"cstCoZ"},
		{"ZcstCom"},
		{"cstComZ"},
		{"ZcstComU"},
		{"cstComUZ"},
		{"ZcstComUs"},
		{"cstComUsZ"},
		{"ZcstComUsd"},
		{"cstComUsdZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.CstComUsd()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestPnlPctAValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"pnlPctA", 7},
		{"pnlPctA ", 7},
		{"pnlPctA\n", 7},
		{"pnlPctA.", 7},
		{"pnlPctA:", 7},
		{"pnlPctA,", 7},
		{"pnlPctA\"", 7},
		{"pnlPctA(", 7},
		{"pnlPctA)", 7},
		{"pnlPctA[", 7},
		{"pnlPctA]", 7},
		{"pnlPctA// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.PnlPctA()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestPnlPctAInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_p"},
		{"p_"},
		{"_pn"},
		{"pn_"},
		{"_pnl"},
		{"pnl_"},
		{"_pnlP"},
		{"pnlP_"},
		{"_pnlPc"},
		{"pnlPc_"},
		{"_pnlPct"},
		{"pnlPct_"},
		{"_pnlPctA"},
		{"pnlPctA_"},
		{"9713p"},
		{"p9713"},
		{"9713pn"},
		{"pn9713"},
		{"9713pnl"},
		{"pnl9713"},
		{"9713pnlP"},
		{"pnlP9713"},
		{"9713pnlPc"},
		{"pnlPc9713"},
		{"9713pnlPct"},
		{"pnlPct9713"},
		{"9713pnlPctA"},
		{"pnlPctA9713"},
		{"Zp"},
		{"pZ"},
		{"Zpn"},
		{"pnZ"},
		{"Zpnl"},
		{"pnlZ"},
		{"ZpnlP"},
		{"pnlPZ"},
		{"ZpnlPc"},
		{"pnlPcZ"},
		{"ZpnlPct"},
		{"pnlPctZ"},
		{"ZpnlPctA"},
		{"pnlPctAZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.PnlPctA()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestPnlPctBValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"pnlPctB", 7},
		{"pnlPctB ", 7},
		{"pnlPctB\n", 7},
		{"pnlPctB.", 7},
		{"pnlPctB:", 7},
		{"pnlPctB,", 7},
		{"pnlPctB\"", 7},
		{"pnlPctB(", 7},
		{"pnlPctB)", 7},
		{"pnlPctB[", 7},
		{"pnlPctB]", 7},
		{"pnlPctB// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.PnlPctB()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestPnlPctBInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_p"},
		{"p_"},
		{"_pn"},
		{"pn_"},
		{"_pnl"},
		{"pnl_"},
		{"_pnlP"},
		{"pnlP_"},
		{"_pnlPc"},
		{"pnlPc_"},
		{"_pnlPct"},
		{"pnlPct_"},
		{"_pnlPctB"},
		{"pnlPctB_"},
		{"9713p"},
		{"p9713"},
		{"9713pn"},
		{"pn9713"},
		{"9713pnl"},
		{"pnl9713"},
		{"9713pnlP"},
		{"pnlP9713"},
		{"9713pnlPc"},
		{"pnlPc9713"},
		{"9713pnlPct"},
		{"pnlPct9713"},
		{"9713pnlPctB"},
		{"pnlPctB9713"},
		{"Zp"},
		{"pZ"},
		{"Zpn"},
		{"pnZ"},
		{"Zpnl"},
		{"pnlZ"},
		{"ZpnlP"},
		{"pnlPZ"},
		{"ZpnlPc"},
		{"pnlPcZ"},
		{"ZpnlPct"},
		{"pnlPctZ"},
		{"ZpnlPctB"},
		{"pnlPctBZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.PnlPctB()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestPnlPctDltValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"pnlPctDlt", 9},
		{"pnlPctDlt ", 9},
		{"pnlPctDlt\n", 9},
		{"pnlPctDlt.", 9},
		{"pnlPctDlt:", 9},
		{"pnlPctDlt,", 9},
		{"pnlPctDlt\"", 9},
		{"pnlPctDlt(", 9},
		{"pnlPctDlt)", 9},
		{"pnlPctDlt[", 9},
		{"pnlPctDlt]", 9},
		{"pnlPctDlt// comment", 9},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.PnlPctDlt()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestPnlPctDltInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_p"},
		{"p_"},
		{"_pn"},
		{"pn_"},
		{"_pnl"},
		{"pnl_"},
		{"_pnlP"},
		{"pnlP_"},
		{"_pnlPc"},
		{"pnlPc_"},
		{"_pnlPct"},
		{"pnlPct_"},
		{"_pnlPctD"},
		{"pnlPctD_"},
		{"_pnlPctDl"},
		{"pnlPctDl_"},
		{"_pnlPctDlt"},
		{"pnlPctDlt_"},
		{"9713p"},
		{"p9713"},
		{"9713pn"},
		{"pn9713"},
		{"9713pnl"},
		{"pnl9713"},
		{"9713pnlP"},
		{"pnlP9713"},
		{"9713pnlPc"},
		{"pnlPc9713"},
		{"9713pnlPct"},
		{"pnlPct9713"},
		{"9713pnlPctD"},
		{"pnlPctD9713"},
		{"9713pnlPctDl"},
		{"pnlPctDl9713"},
		{"9713pnlPctDlt"},
		{"pnlPctDlt9713"},
		{"Zp"},
		{"pZ"},
		{"Zpn"},
		{"pnZ"},
		{"Zpnl"},
		{"pnlZ"},
		{"ZpnlP"},
		{"pnlPZ"},
		{"ZpnlPc"},
		{"pnlPcZ"},
		{"ZpnlPct"},
		{"pnlPctZ"},
		{"ZpnlPctD"},
		{"pnlPctDZ"},
		{"ZpnlPctDl"},
		{"pnlPctDlZ"},
		{"ZpnlPctDlt"},
		{"pnlPctDltZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.PnlPctDlt()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestScsPctAValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"scsPctA", 7},
		{"scsPctA ", 7},
		{"scsPctA\n", 7},
		{"scsPctA.", 7},
		{"scsPctA:", 7},
		{"scsPctA,", 7},
		{"scsPctA\"", 7},
		{"scsPctA(", 7},
		{"scsPctA)", 7},
		{"scsPctA[", 7},
		{"scsPctA]", 7},
		{"scsPctA// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.ScsPctA()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestScsPctAInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_s"},
		{"s_"},
		{"_sc"},
		{"sc_"},
		{"_scs"},
		{"scs_"},
		{"_scsP"},
		{"scsP_"},
		{"_scsPc"},
		{"scsPc_"},
		{"_scsPct"},
		{"scsPct_"},
		{"_scsPctA"},
		{"scsPctA_"},
		{"9713s"},
		{"s9713"},
		{"9713sc"},
		{"sc9713"},
		{"9713scs"},
		{"scs9713"},
		{"9713scsP"},
		{"scsP9713"},
		{"9713scsPc"},
		{"scsPc9713"},
		{"9713scsPct"},
		{"scsPct9713"},
		{"9713scsPctA"},
		{"scsPctA9713"},
		{"Zs"},
		{"sZ"},
		{"Zsc"},
		{"scZ"},
		{"Zscs"},
		{"scsZ"},
		{"ZscsP"},
		{"scsPZ"},
		{"ZscsPc"},
		{"scsPcZ"},
		{"ZscsPct"},
		{"scsPctZ"},
		{"ZscsPctA"},
		{"scsPctAZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.ScsPctA()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestScsPctBValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"scsPctB", 7},
		{"scsPctB ", 7},
		{"scsPctB\n", 7},
		{"scsPctB.", 7},
		{"scsPctB:", 7},
		{"scsPctB,", 7},
		{"scsPctB\"", 7},
		{"scsPctB(", 7},
		{"scsPctB)", 7},
		{"scsPctB[", 7},
		{"scsPctB]", 7},
		{"scsPctB// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.ScsPctB()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestScsPctBInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_s"},
		{"s_"},
		{"_sc"},
		{"sc_"},
		{"_scs"},
		{"scs_"},
		{"_scsP"},
		{"scsP_"},
		{"_scsPc"},
		{"scsPc_"},
		{"_scsPct"},
		{"scsPct_"},
		{"_scsPctB"},
		{"scsPctB_"},
		{"9713s"},
		{"s9713"},
		{"9713sc"},
		{"sc9713"},
		{"9713scs"},
		{"scs9713"},
		{"9713scsP"},
		{"scsP9713"},
		{"9713scsPc"},
		{"scsPc9713"},
		{"9713scsPct"},
		{"scsPct9713"},
		{"9713scsPctB"},
		{"scsPctB9713"},
		{"Zs"},
		{"sZ"},
		{"Zsc"},
		{"scZ"},
		{"Zscs"},
		{"scsZ"},
		{"ZscsP"},
		{"scsPZ"},
		{"ZscsPc"},
		{"scsPcZ"},
		{"ZscsPct"},
		{"scsPctZ"},
		{"ZscsPctB"},
		{"scsPctBZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.ScsPctB()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestScsPctDltValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"scsPctDlt", 9},
		{"scsPctDlt ", 9},
		{"scsPctDlt\n", 9},
		{"scsPctDlt.", 9},
		{"scsPctDlt:", 9},
		{"scsPctDlt,", 9},
		{"scsPctDlt\"", 9},
		{"scsPctDlt(", 9},
		{"scsPctDlt)", 9},
		{"scsPctDlt[", 9},
		{"scsPctDlt]", 9},
		{"scsPctDlt// comment", 9},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.ScsPctDlt()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestScsPctDltInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_s"},
		{"s_"},
		{"_sc"},
		{"sc_"},
		{"_scs"},
		{"scs_"},
		{"_scsP"},
		{"scsP_"},
		{"_scsPc"},
		{"scsPc_"},
		{"_scsPct"},
		{"scsPct_"},
		{"_scsPctD"},
		{"scsPctD_"},
		{"_scsPctDl"},
		{"scsPctDl_"},
		{"_scsPctDlt"},
		{"scsPctDlt_"},
		{"9713s"},
		{"s9713"},
		{"9713sc"},
		{"sc9713"},
		{"9713scs"},
		{"scs9713"},
		{"9713scsP"},
		{"scsP9713"},
		{"9713scsPc"},
		{"scsPc9713"},
		{"9713scsPct"},
		{"scsPct9713"},
		{"9713scsPctD"},
		{"scsPctD9713"},
		{"9713scsPctDl"},
		{"scsPctDl9713"},
		{"9713scsPctDlt"},
		{"scsPctDlt9713"},
		{"Zs"},
		{"sZ"},
		{"Zsc"},
		{"scZ"},
		{"Zscs"},
		{"scsZ"},
		{"ZscsP"},
		{"scsPZ"},
		{"ZscsPc"},
		{"scsPcZ"},
		{"ZscsPct"},
		{"scsPctZ"},
		{"ZscsPctD"},
		{"scsPctDZ"},
		{"ZscsPctDl"},
		{"scsPctDlZ"},
		{"ZscsPctDlt"},
		{"scsPctDltZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.ScsPctDlt()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestPipPerDayAValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"pipPerDayA", 10},
		{"pipPerDayA ", 10},
		{"pipPerDayA\n", 10},
		{"pipPerDayA.", 10},
		{"pipPerDayA:", 10},
		{"pipPerDayA,", 10},
		{"pipPerDayA\"", 10},
		{"pipPerDayA(", 10},
		{"pipPerDayA)", 10},
		{"pipPerDayA[", 10},
		{"pipPerDayA]", 10},
		{"pipPerDayA// comment", 10},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.PipPerDayA()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestPipPerDayAInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_p"},
		{"p_"},
		{"_pi"},
		{"pi_"},
		{"_pip"},
		{"pip_"},
		{"_pipP"},
		{"pipP_"},
		{"_pipPe"},
		{"pipPe_"},
		{"_pipPer"},
		{"pipPer_"},
		{"_pipPerD"},
		{"pipPerD_"},
		{"_pipPerDa"},
		{"pipPerDa_"},
		{"_pipPerDay"},
		{"pipPerDay_"},
		{"_pipPerDayA"},
		{"pipPerDayA_"},
		{"9713p"},
		{"p9713"},
		{"9713pi"},
		{"pi9713"},
		{"9713pip"},
		{"pip9713"},
		{"9713pipP"},
		{"pipP9713"},
		{"9713pipPe"},
		{"pipPe9713"},
		{"9713pipPer"},
		{"pipPer9713"},
		{"9713pipPerD"},
		{"pipPerD9713"},
		{"9713pipPerDa"},
		{"pipPerDa9713"},
		{"9713pipPerDay"},
		{"pipPerDay9713"},
		{"9713pipPerDayA"},
		{"pipPerDayA9713"},
		{"Zp"},
		{"pZ"},
		{"Zpi"},
		{"piZ"},
		{"Zpip"},
		{"pipZ"},
		{"ZpipP"},
		{"pipPZ"},
		{"ZpipPe"},
		{"pipPeZ"},
		{"ZpipPer"},
		{"pipPerZ"},
		{"ZpipPerD"},
		{"pipPerDZ"},
		{"ZpipPerDa"},
		{"pipPerDaZ"},
		{"ZpipPerDay"},
		{"pipPerDayZ"},
		{"ZpipPerDayA"},
		{"pipPerDayAZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.PipPerDayA()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestPipPerDayBValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"pipPerDayB", 10},
		{"pipPerDayB ", 10},
		{"pipPerDayB\n", 10},
		{"pipPerDayB.", 10},
		{"pipPerDayB:", 10},
		{"pipPerDayB,", 10},
		{"pipPerDayB\"", 10},
		{"pipPerDayB(", 10},
		{"pipPerDayB)", 10},
		{"pipPerDayB[", 10},
		{"pipPerDayB]", 10},
		{"pipPerDayB// comment", 10},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.PipPerDayB()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestPipPerDayBInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_p"},
		{"p_"},
		{"_pi"},
		{"pi_"},
		{"_pip"},
		{"pip_"},
		{"_pipP"},
		{"pipP_"},
		{"_pipPe"},
		{"pipPe_"},
		{"_pipPer"},
		{"pipPer_"},
		{"_pipPerD"},
		{"pipPerD_"},
		{"_pipPerDa"},
		{"pipPerDa_"},
		{"_pipPerDay"},
		{"pipPerDay_"},
		{"_pipPerDayB"},
		{"pipPerDayB_"},
		{"9713p"},
		{"p9713"},
		{"9713pi"},
		{"pi9713"},
		{"9713pip"},
		{"pip9713"},
		{"9713pipP"},
		{"pipP9713"},
		{"9713pipPe"},
		{"pipPe9713"},
		{"9713pipPer"},
		{"pipPer9713"},
		{"9713pipPerD"},
		{"pipPerD9713"},
		{"9713pipPerDa"},
		{"pipPerDa9713"},
		{"9713pipPerDay"},
		{"pipPerDay9713"},
		{"9713pipPerDayB"},
		{"pipPerDayB9713"},
		{"Zp"},
		{"pZ"},
		{"Zpi"},
		{"piZ"},
		{"Zpip"},
		{"pipZ"},
		{"ZpipP"},
		{"pipPZ"},
		{"ZpipPe"},
		{"pipPeZ"},
		{"ZpipPer"},
		{"pipPerZ"},
		{"ZpipPerD"},
		{"pipPerDZ"},
		{"ZpipPerDa"},
		{"pipPerDaZ"},
		{"ZpipPerDay"},
		{"pipPerDayZ"},
		{"ZpipPerDayB"},
		{"pipPerDayBZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.PipPerDayB()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestPipPerDayDltValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"pipPerDayDlt", 12},
		{"pipPerDayDlt ", 12},
		{"pipPerDayDlt\n", 12},
		{"pipPerDayDlt.", 12},
		{"pipPerDayDlt:", 12},
		{"pipPerDayDlt,", 12},
		{"pipPerDayDlt\"", 12},
		{"pipPerDayDlt(", 12},
		{"pipPerDayDlt)", 12},
		{"pipPerDayDlt[", 12},
		{"pipPerDayDlt]", 12},
		{"pipPerDayDlt// comment", 12},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.PipPerDayDlt()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestPipPerDayDltInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_p"},
		{"p_"},
		{"_pi"},
		{"pi_"},
		{"_pip"},
		{"pip_"},
		{"_pipP"},
		{"pipP_"},
		{"_pipPe"},
		{"pipPe_"},
		{"_pipPer"},
		{"pipPer_"},
		{"_pipPerD"},
		{"pipPerD_"},
		{"_pipPerDa"},
		{"pipPerDa_"},
		{"_pipPerDay"},
		{"pipPerDay_"},
		{"_pipPerDayD"},
		{"pipPerDayD_"},
		{"_pipPerDayDl"},
		{"pipPerDayDl_"},
		{"_pipPerDayDlt"},
		{"pipPerDayDlt_"},
		{"9713p"},
		{"p9713"},
		{"9713pi"},
		{"pi9713"},
		{"9713pip"},
		{"pip9713"},
		{"9713pipP"},
		{"pipP9713"},
		{"9713pipPe"},
		{"pipPe9713"},
		{"9713pipPer"},
		{"pipPer9713"},
		{"9713pipPerD"},
		{"pipPerD9713"},
		{"9713pipPerDa"},
		{"pipPerDa9713"},
		{"9713pipPerDay"},
		{"pipPerDay9713"},
		{"9713pipPerDayD"},
		{"pipPerDayD9713"},
		{"9713pipPerDayDl"},
		{"pipPerDayDl9713"},
		{"9713pipPerDayDlt"},
		{"pipPerDayDlt9713"},
		{"Zp"},
		{"pZ"},
		{"Zpi"},
		{"piZ"},
		{"Zpip"},
		{"pipZ"},
		{"ZpipP"},
		{"pipPZ"},
		{"ZpipPe"},
		{"pipPeZ"},
		{"ZpipPer"},
		{"pipPerZ"},
		{"ZpipPerD"},
		{"pipPerDZ"},
		{"ZpipPerDa"},
		{"pipPerDaZ"},
		{"ZpipPerDay"},
		{"pipPerDayZ"},
		{"ZpipPerDayD"},
		{"pipPerDayDZ"},
		{"ZpipPerDayDl"},
		{"pipPerDayDlZ"},
		{"ZpipPerDayDlt"},
		{"pipPerDayDltZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.PipPerDayDlt()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestUsdPerDayAValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"usdPerDayA", 10},
		{"usdPerDayA ", 10},
		{"usdPerDayA\n", 10},
		{"usdPerDayA.", 10},
		{"usdPerDayA:", 10},
		{"usdPerDayA,", 10},
		{"usdPerDayA\"", 10},
		{"usdPerDayA(", 10},
		{"usdPerDayA)", 10},
		{"usdPerDayA[", 10},
		{"usdPerDayA]", 10},
		{"usdPerDayA// comment", 10},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.UsdPerDayA()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestUsdPerDayAInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_u"},
		{"u_"},
		{"_us"},
		{"us_"},
		{"_usd"},
		{"usd_"},
		{"_usdP"},
		{"usdP_"},
		{"_usdPe"},
		{"usdPe_"},
		{"_usdPer"},
		{"usdPer_"},
		{"_usdPerD"},
		{"usdPerD_"},
		{"_usdPerDa"},
		{"usdPerDa_"},
		{"_usdPerDay"},
		{"usdPerDay_"},
		{"_usdPerDayA"},
		{"usdPerDayA_"},
		{"9713u"},
		{"u9713"},
		{"9713us"},
		{"us9713"},
		{"9713usd"},
		{"usd9713"},
		{"9713usdP"},
		{"usdP9713"},
		{"9713usdPe"},
		{"usdPe9713"},
		{"9713usdPer"},
		{"usdPer9713"},
		{"9713usdPerD"},
		{"usdPerD9713"},
		{"9713usdPerDa"},
		{"usdPerDa9713"},
		{"9713usdPerDay"},
		{"usdPerDay9713"},
		{"9713usdPerDayA"},
		{"usdPerDayA9713"},
		{"Zu"},
		{"uZ"},
		{"Zus"},
		{"usZ"},
		{"Zusd"},
		{"usdZ"},
		{"ZusdP"},
		{"usdPZ"},
		{"ZusdPe"},
		{"usdPeZ"},
		{"ZusdPer"},
		{"usdPerZ"},
		{"ZusdPerD"},
		{"usdPerDZ"},
		{"ZusdPerDa"},
		{"usdPerDaZ"},
		{"ZusdPerDay"},
		{"usdPerDayZ"},
		{"ZusdPerDayA"},
		{"usdPerDayAZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.UsdPerDayA()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestUsdPerDayBValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"usdPerDayB", 10},
		{"usdPerDayB ", 10},
		{"usdPerDayB\n", 10},
		{"usdPerDayB.", 10},
		{"usdPerDayB:", 10},
		{"usdPerDayB,", 10},
		{"usdPerDayB\"", 10},
		{"usdPerDayB(", 10},
		{"usdPerDayB)", 10},
		{"usdPerDayB[", 10},
		{"usdPerDayB]", 10},
		{"usdPerDayB// comment", 10},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.UsdPerDayB()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestUsdPerDayBInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_u"},
		{"u_"},
		{"_us"},
		{"us_"},
		{"_usd"},
		{"usd_"},
		{"_usdP"},
		{"usdP_"},
		{"_usdPe"},
		{"usdPe_"},
		{"_usdPer"},
		{"usdPer_"},
		{"_usdPerD"},
		{"usdPerD_"},
		{"_usdPerDa"},
		{"usdPerDa_"},
		{"_usdPerDay"},
		{"usdPerDay_"},
		{"_usdPerDayB"},
		{"usdPerDayB_"},
		{"9713u"},
		{"u9713"},
		{"9713us"},
		{"us9713"},
		{"9713usd"},
		{"usd9713"},
		{"9713usdP"},
		{"usdP9713"},
		{"9713usdPe"},
		{"usdPe9713"},
		{"9713usdPer"},
		{"usdPer9713"},
		{"9713usdPerD"},
		{"usdPerD9713"},
		{"9713usdPerDa"},
		{"usdPerDa9713"},
		{"9713usdPerDay"},
		{"usdPerDay9713"},
		{"9713usdPerDayB"},
		{"usdPerDayB9713"},
		{"Zu"},
		{"uZ"},
		{"Zus"},
		{"usZ"},
		{"Zusd"},
		{"usdZ"},
		{"ZusdP"},
		{"usdPZ"},
		{"ZusdPe"},
		{"usdPeZ"},
		{"ZusdPer"},
		{"usdPerZ"},
		{"ZusdPerD"},
		{"usdPerDZ"},
		{"ZusdPerDa"},
		{"usdPerDaZ"},
		{"ZusdPerDay"},
		{"usdPerDayZ"},
		{"ZusdPerDayB"},
		{"usdPerDayBZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.UsdPerDayB()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestUsdPerDayDltValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"usdPerDayDlt", 12},
		{"usdPerDayDlt ", 12},
		{"usdPerDayDlt\n", 12},
		{"usdPerDayDlt.", 12},
		{"usdPerDayDlt:", 12},
		{"usdPerDayDlt,", 12},
		{"usdPerDayDlt\"", 12},
		{"usdPerDayDlt(", 12},
		{"usdPerDayDlt)", 12},
		{"usdPerDayDlt[", 12},
		{"usdPerDayDlt]", 12},
		{"usdPerDayDlt// comment", 12},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.UsdPerDayDlt()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestUsdPerDayDltInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_u"},
		{"u_"},
		{"_us"},
		{"us_"},
		{"_usd"},
		{"usd_"},
		{"_usdP"},
		{"usdP_"},
		{"_usdPe"},
		{"usdPe_"},
		{"_usdPer"},
		{"usdPer_"},
		{"_usdPerD"},
		{"usdPerD_"},
		{"_usdPerDa"},
		{"usdPerDa_"},
		{"_usdPerDay"},
		{"usdPerDay_"},
		{"_usdPerDayD"},
		{"usdPerDayD_"},
		{"_usdPerDayDl"},
		{"usdPerDayDl_"},
		{"_usdPerDayDlt"},
		{"usdPerDayDlt_"},
		{"9713u"},
		{"u9713"},
		{"9713us"},
		{"us9713"},
		{"9713usd"},
		{"usd9713"},
		{"9713usdP"},
		{"usdP9713"},
		{"9713usdPe"},
		{"usdPe9713"},
		{"9713usdPer"},
		{"usdPer9713"},
		{"9713usdPerD"},
		{"usdPerD9713"},
		{"9713usdPerDa"},
		{"usdPerDa9713"},
		{"9713usdPerDay"},
		{"usdPerDay9713"},
		{"9713usdPerDayD"},
		{"usdPerDayD9713"},
		{"9713usdPerDayDl"},
		{"usdPerDayDl9713"},
		{"9713usdPerDayDlt"},
		{"usdPerDayDlt9713"},
		{"Zu"},
		{"uZ"},
		{"Zus"},
		{"usZ"},
		{"Zusd"},
		{"usdZ"},
		{"ZusdP"},
		{"usdPZ"},
		{"ZusdPe"},
		{"usdPeZ"},
		{"ZusdPer"},
		{"usdPerZ"},
		{"ZusdPerD"},
		{"usdPerDZ"},
		{"ZusdPerDa"},
		{"usdPerDaZ"},
		{"ZusdPerDay"},
		{"usdPerDayZ"},
		{"ZusdPerDayD"},
		{"usdPerDayDZ"},
		{"ZusdPerDayDl"},
		{"usdPerDayDlZ"},
		{"ZusdPerDayDlt"},
		{"usdPerDayDltZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.UsdPerDayDlt()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestScsPerDayAValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"scsPerDayA", 10},
		{"scsPerDayA ", 10},
		{"scsPerDayA\n", 10},
		{"scsPerDayA.", 10},
		{"scsPerDayA:", 10},
		{"scsPerDayA,", 10},
		{"scsPerDayA\"", 10},
		{"scsPerDayA(", 10},
		{"scsPerDayA)", 10},
		{"scsPerDayA[", 10},
		{"scsPerDayA]", 10},
		{"scsPerDayA// comment", 10},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.ScsPerDayA()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestScsPerDayAInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_s"},
		{"s_"},
		{"_sc"},
		{"sc_"},
		{"_scs"},
		{"scs_"},
		{"_scsP"},
		{"scsP_"},
		{"_scsPe"},
		{"scsPe_"},
		{"_scsPer"},
		{"scsPer_"},
		{"_scsPerD"},
		{"scsPerD_"},
		{"_scsPerDa"},
		{"scsPerDa_"},
		{"_scsPerDay"},
		{"scsPerDay_"},
		{"_scsPerDayA"},
		{"scsPerDayA_"},
		{"9713s"},
		{"s9713"},
		{"9713sc"},
		{"sc9713"},
		{"9713scs"},
		{"scs9713"},
		{"9713scsP"},
		{"scsP9713"},
		{"9713scsPe"},
		{"scsPe9713"},
		{"9713scsPer"},
		{"scsPer9713"},
		{"9713scsPerD"},
		{"scsPerD9713"},
		{"9713scsPerDa"},
		{"scsPerDa9713"},
		{"9713scsPerDay"},
		{"scsPerDay9713"},
		{"9713scsPerDayA"},
		{"scsPerDayA9713"},
		{"Zs"},
		{"sZ"},
		{"Zsc"},
		{"scZ"},
		{"Zscs"},
		{"scsZ"},
		{"ZscsP"},
		{"scsPZ"},
		{"ZscsPe"},
		{"scsPeZ"},
		{"ZscsPer"},
		{"scsPerZ"},
		{"ZscsPerD"},
		{"scsPerDZ"},
		{"ZscsPerDa"},
		{"scsPerDaZ"},
		{"ZscsPerDay"},
		{"scsPerDayZ"},
		{"ZscsPerDayA"},
		{"scsPerDayAZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.ScsPerDayA()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestScsPerDayBValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"scsPerDayB", 10},
		{"scsPerDayB ", 10},
		{"scsPerDayB\n", 10},
		{"scsPerDayB.", 10},
		{"scsPerDayB:", 10},
		{"scsPerDayB,", 10},
		{"scsPerDayB\"", 10},
		{"scsPerDayB(", 10},
		{"scsPerDayB)", 10},
		{"scsPerDayB[", 10},
		{"scsPerDayB]", 10},
		{"scsPerDayB// comment", 10},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.ScsPerDayB()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestScsPerDayBInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_s"},
		{"s_"},
		{"_sc"},
		{"sc_"},
		{"_scs"},
		{"scs_"},
		{"_scsP"},
		{"scsP_"},
		{"_scsPe"},
		{"scsPe_"},
		{"_scsPer"},
		{"scsPer_"},
		{"_scsPerD"},
		{"scsPerD_"},
		{"_scsPerDa"},
		{"scsPerDa_"},
		{"_scsPerDay"},
		{"scsPerDay_"},
		{"_scsPerDayB"},
		{"scsPerDayB_"},
		{"9713s"},
		{"s9713"},
		{"9713sc"},
		{"sc9713"},
		{"9713scs"},
		{"scs9713"},
		{"9713scsP"},
		{"scsP9713"},
		{"9713scsPe"},
		{"scsPe9713"},
		{"9713scsPer"},
		{"scsPer9713"},
		{"9713scsPerD"},
		{"scsPerD9713"},
		{"9713scsPerDa"},
		{"scsPerDa9713"},
		{"9713scsPerDay"},
		{"scsPerDay9713"},
		{"9713scsPerDayB"},
		{"scsPerDayB9713"},
		{"Zs"},
		{"sZ"},
		{"Zsc"},
		{"scZ"},
		{"Zscs"},
		{"scsZ"},
		{"ZscsP"},
		{"scsPZ"},
		{"ZscsPe"},
		{"scsPeZ"},
		{"ZscsPer"},
		{"scsPerZ"},
		{"ZscsPerD"},
		{"scsPerDZ"},
		{"ZscsPerDa"},
		{"scsPerDaZ"},
		{"ZscsPerDay"},
		{"scsPerDayZ"},
		{"ZscsPerDayB"},
		{"scsPerDayBZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.ScsPerDayB()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestScsPerDayDltValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"scsPerDayDlt", 12},
		{"scsPerDayDlt ", 12},
		{"scsPerDayDlt\n", 12},
		{"scsPerDayDlt.", 12},
		{"scsPerDayDlt:", 12},
		{"scsPerDayDlt,", 12},
		{"scsPerDayDlt\"", 12},
		{"scsPerDayDlt(", 12},
		{"scsPerDayDlt)", 12},
		{"scsPerDayDlt[", 12},
		{"scsPerDayDlt]", 12},
		{"scsPerDayDlt// comment", 12},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.ScsPerDayDlt()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestScsPerDayDltInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_s"},
		{"s_"},
		{"_sc"},
		{"sc_"},
		{"_scs"},
		{"scs_"},
		{"_scsP"},
		{"scsP_"},
		{"_scsPe"},
		{"scsPe_"},
		{"_scsPer"},
		{"scsPer_"},
		{"_scsPerD"},
		{"scsPerD_"},
		{"_scsPerDa"},
		{"scsPerDa_"},
		{"_scsPerDay"},
		{"scsPerDay_"},
		{"_scsPerDayD"},
		{"scsPerDayD_"},
		{"_scsPerDayDl"},
		{"scsPerDayDl_"},
		{"_scsPerDayDlt"},
		{"scsPerDayDlt_"},
		{"9713s"},
		{"s9713"},
		{"9713sc"},
		{"sc9713"},
		{"9713scs"},
		{"scs9713"},
		{"9713scsP"},
		{"scsP9713"},
		{"9713scsPe"},
		{"scsPe9713"},
		{"9713scsPer"},
		{"scsPer9713"},
		{"9713scsPerD"},
		{"scsPerD9713"},
		{"9713scsPerDa"},
		{"scsPerDa9713"},
		{"9713scsPerDay"},
		{"scsPerDay9713"},
		{"9713scsPerDayD"},
		{"scsPerDayD9713"},
		{"9713scsPerDayDl"},
		{"scsPerDayDl9713"},
		{"9713scsPerDayDlt"},
		{"scsPerDayDlt9713"},
		{"Zs"},
		{"sZ"},
		{"Zsc"},
		{"scZ"},
		{"Zscs"},
		{"scsZ"},
		{"ZscsP"},
		{"scsPZ"},
		{"ZscsPe"},
		{"scsPeZ"},
		{"ZscsPer"},
		{"scsPerZ"},
		{"ZscsPerD"},
		{"scsPerDZ"},
		{"ZscsPerDa"},
		{"scsPerDaZ"},
		{"ZscsPerDay"},
		{"scsPerDayZ"},
		{"ZscsPerDayD"},
		{"scsPerDayDZ"},
		{"ZscsPerDayDl"},
		{"scsPerDayDlZ"},
		{"ZscsPerDayDlt"},
		{"scsPerDayDltZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.ScsPerDayDlt()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestOpnPerDayAValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"opnPerDayA", 10},
		{"opnPerDayA ", 10},
		{"opnPerDayA\n", 10},
		{"opnPerDayA.", 10},
		{"opnPerDayA:", 10},
		{"opnPerDayA,", 10},
		{"opnPerDayA\"", 10},
		{"opnPerDayA(", 10},
		{"opnPerDayA)", 10},
		{"opnPerDayA[", 10},
		{"opnPerDayA]", 10},
		{"opnPerDayA// comment", 10},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.OpnPerDayA()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestOpnPerDayAInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_o"},
		{"o_"},
		{"_op"},
		{"op_"},
		{"_opn"},
		{"opn_"},
		{"_opnP"},
		{"opnP_"},
		{"_opnPe"},
		{"opnPe_"},
		{"_opnPer"},
		{"opnPer_"},
		{"_opnPerD"},
		{"opnPerD_"},
		{"_opnPerDa"},
		{"opnPerDa_"},
		{"_opnPerDay"},
		{"opnPerDay_"},
		{"_opnPerDayA"},
		{"opnPerDayA_"},
		{"9713o"},
		{"o9713"},
		{"9713op"},
		{"op9713"},
		{"9713opn"},
		{"opn9713"},
		{"9713opnP"},
		{"opnP9713"},
		{"9713opnPe"},
		{"opnPe9713"},
		{"9713opnPer"},
		{"opnPer9713"},
		{"9713opnPerD"},
		{"opnPerD9713"},
		{"9713opnPerDa"},
		{"opnPerDa9713"},
		{"9713opnPerDay"},
		{"opnPerDay9713"},
		{"9713opnPerDayA"},
		{"opnPerDayA9713"},
		{"Zo"},
		{"oZ"},
		{"Zop"},
		{"opZ"},
		{"Zopn"},
		{"opnZ"},
		{"ZopnP"},
		{"opnPZ"},
		{"ZopnPe"},
		{"opnPeZ"},
		{"ZopnPer"},
		{"opnPerZ"},
		{"ZopnPerD"},
		{"opnPerDZ"},
		{"ZopnPerDa"},
		{"opnPerDaZ"},
		{"ZopnPerDay"},
		{"opnPerDayZ"},
		{"ZopnPerDayA"},
		{"opnPerDayAZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.OpnPerDayA()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestOpnPerDayBValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"opnPerDayB", 10},
		{"opnPerDayB ", 10},
		{"opnPerDayB\n", 10},
		{"opnPerDayB.", 10},
		{"opnPerDayB:", 10},
		{"opnPerDayB,", 10},
		{"opnPerDayB\"", 10},
		{"opnPerDayB(", 10},
		{"opnPerDayB)", 10},
		{"opnPerDayB[", 10},
		{"opnPerDayB]", 10},
		{"opnPerDayB// comment", 10},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.OpnPerDayB()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestOpnPerDayBInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_o"},
		{"o_"},
		{"_op"},
		{"op_"},
		{"_opn"},
		{"opn_"},
		{"_opnP"},
		{"opnP_"},
		{"_opnPe"},
		{"opnPe_"},
		{"_opnPer"},
		{"opnPer_"},
		{"_opnPerD"},
		{"opnPerD_"},
		{"_opnPerDa"},
		{"opnPerDa_"},
		{"_opnPerDay"},
		{"opnPerDay_"},
		{"_opnPerDayB"},
		{"opnPerDayB_"},
		{"9713o"},
		{"o9713"},
		{"9713op"},
		{"op9713"},
		{"9713opn"},
		{"opn9713"},
		{"9713opnP"},
		{"opnP9713"},
		{"9713opnPe"},
		{"opnPe9713"},
		{"9713opnPer"},
		{"opnPer9713"},
		{"9713opnPerD"},
		{"opnPerD9713"},
		{"9713opnPerDa"},
		{"opnPerDa9713"},
		{"9713opnPerDay"},
		{"opnPerDay9713"},
		{"9713opnPerDayB"},
		{"opnPerDayB9713"},
		{"Zo"},
		{"oZ"},
		{"Zop"},
		{"opZ"},
		{"Zopn"},
		{"opnZ"},
		{"ZopnP"},
		{"opnPZ"},
		{"ZopnPe"},
		{"opnPeZ"},
		{"ZopnPer"},
		{"opnPerZ"},
		{"ZopnPerD"},
		{"opnPerDZ"},
		{"ZopnPerDa"},
		{"opnPerDaZ"},
		{"ZopnPerDay"},
		{"opnPerDayZ"},
		{"ZopnPerDayB"},
		{"opnPerDayBZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.OpnPerDayB()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestOpnPerDayDltValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"opnPerDayDlt", 12},
		{"opnPerDayDlt ", 12},
		{"opnPerDayDlt\n", 12},
		{"opnPerDayDlt.", 12},
		{"opnPerDayDlt:", 12},
		{"opnPerDayDlt,", 12},
		{"opnPerDayDlt\"", 12},
		{"opnPerDayDlt(", 12},
		{"opnPerDayDlt)", 12},
		{"opnPerDayDlt[", 12},
		{"opnPerDayDlt]", 12},
		{"opnPerDayDlt// comment", 12},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.OpnPerDayDlt()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestOpnPerDayDltInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_o"},
		{"o_"},
		{"_op"},
		{"op_"},
		{"_opn"},
		{"opn_"},
		{"_opnP"},
		{"opnP_"},
		{"_opnPe"},
		{"opnPe_"},
		{"_opnPer"},
		{"opnPer_"},
		{"_opnPerD"},
		{"opnPerD_"},
		{"_opnPerDa"},
		{"opnPerDa_"},
		{"_opnPerDay"},
		{"opnPerDay_"},
		{"_opnPerDayD"},
		{"opnPerDayD_"},
		{"_opnPerDayDl"},
		{"opnPerDayDl_"},
		{"_opnPerDayDlt"},
		{"opnPerDayDlt_"},
		{"9713o"},
		{"o9713"},
		{"9713op"},
		{"op9713"},
		{"9713opn"},
		{"opn9713"},
		{"9713opnP"},
		{"opnP9713"},
		{"9713opnPe"},
		{"opnPe9713"},
		{"9713opnPer"},
		{"opnPer9713"},
		{"9713opnPerD"},
		{"opnPerD9713"},
		{"9713opnPerDa"},
		{"opnPerDa9713"},
		{"9713opnPerDay"},
		{"opnPerDay9713"},
		{"9713opnPerDayD"},
		{"opnPerDayD9713"},
		{"9713opnPerDayDl"},
		{"opnPerDayDl9713"},
		{"9713opnPerDayDlt"},
		{"opnPerDayDlt9713"},
		{"Zo"},
		{"oZ"},
		{"Zop"},
		{"opZ"},
		{"Zopn"},
		{"opnZ"},
		{"ZopnP"},
		{"opnPZ"},
		{"ZopnPe"},
		{"opnPeZ"},
		{"ZopnPer"},
		{"opnPerZ"},
		{"ZopnPerD"},
		{"opnPerDZ"},
		{"ZopnPerDa"},
		{"opnPerDaZ"},
		{"ZopnPerDay"},
		{"opnPerDayZ"},
		{"ZopnPerDayD"},
		{"opnPerDayDZ"},
		{"ZopnPerDayDl"},
		{"opnPerDayDlZ"},
		{"ZopnPerDayDlt"},
		{"opnPerDayDltZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.OpnPerDayDlt()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestPnlUsdAValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"pnlUsdA", 7},
		{"pnlUsdA ", 7},
		{"pnlUsdA\n", 7},
		{"pnlUsdA.", 7},
		{"pnlUsdA:", 7},
		{"pnlUsdA,", 7},
		{"pnlUsdA\"", 7},
		{"pnlUsdA(", 7},
		{"pnlUsdA)", 7},
		{"pnlUsdA[", 7},
		{"pnlUsdA]", 7},
		{"pnlUsdA// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.PnlUsdA()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestPnlUsdAInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_p"},
		{"p_"},
		{"_pn"},
		{"pn_"},
		{"_pnl"},
		{"pnl_"},
		{"_pnlU"},
		{"pnlU_"},
		{"_pnlUs"},
		{"pnlUs_"},
		{"_pnlUsd"},
		{"pnlUsd_"},
		{"_pnlUsdA"},
		{"pnlUsdA_"},
		{"9713p"},
		{"p9713"},
		{"9713pn"},
		{"pn9713"},
		{"9713pnl"},
		{"pnl9713"},
		{"9713pnlU"},
		{"pnlU9713"},
		{"9713pnlUs"},
		{"pnlUs9713"},
		{"9713pnlUsd"},
		{"pnlUsd9713"},
		{"9713pnlUsdA"},
		{"pnlUsdA9713"},
		{"Zp"},
		{"pZ"},
		{"Zpn"},
		{"pnZ"},
		{"Zpnl"},
		{"pnlZ"},
		{"ZpnlU"},
		{"pnlUZ"},
		{"ZpnlUs"},
		{"pnlUsZ"},
		{"ZpnlUsd"},
		{"pnlUsdZ"},
		{"ZpnlUsdA"},
		{"pnlUsdAZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.PnlUsdA()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestPnlUsdBValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"pnlUsdB", 7},
		{"pnlUsdB ", 7},
		{"pnlUsdB\n", 7},
		{"pnlUsdB.", 7},
		{"pnlUsdB:", 7},
		{"pnlUsdB,", 7},
		{"pnlUsdB\"", 7},
		{"pnlUsdB(", 7},
		{"pnlUsdB)", 7},
		{"pnlUsdB[", 7},
		{"pnlUsdB]", 7},
		{"pnlUsdB// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.PnlUsdB()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestPnlUsdBInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_p"},
		{"p_"},
		{"_pn"},
		{"pn_"},
		{"_pnl"},
		{"pnl_"},
		{"_pnlU"},
		{"pnlU_"},
		{"_pnlUs"},
		{"pnlUs_"},
		{"_pnlUsd"},
		{"pnlUsd_"},
		{"_pnlUsdB"},
		{"pnlUsdB_"},
		{"9713p"},
		{"p9713"},
		{"9713pn"},
		{"pn9713"},
		{"9713pnl"},
		{"pnl9713"},
		{"9713pnlU"},
		{"pnlU9713"},
		{"9713pnlUs"},
		{"pnlUs9713"},
		{"9713pnlUsd"},
		{"pnlUsd9713"},
		{"9713pnlUsdB"},
		{"pnlUsdB9713"},
		{"Zp"},
		{"pZ"},
		{"Zpn"},
		{"pnZ"},
		{"Zpnl"},
		{"pnlZ"},
		{"ZpnlU"},
		{"pnlUZ"},
		{"ZpnlUs"},
		{"pnlUsZ"},
		{"ZpnlUsd"},
		{"pnlUsdZ"},
		{"ZpnlUsdB"},
		{"pnlUsdBZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.PnlUsdB()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestPnlUsdDltValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"pnlUsdDlt", 9},
		{"pnlUsdDlt ", 9},
		{"pnlUsdDlt\n", 9},
		{"pnlUsdDlt.", 9},
		{"pnlUsdDlt:", 9},
		{"pnlUsdDlt,", 9},
		{"pnlUsdDlt\"", 9},
		{"pnlUsdDlt(", 9},
		{"pnlUsdDlt)", 9},
		{"pnlUsdDlt[", 9},
		{"pnlUsdDlt]", 9},
		{"pnlUsdDlt// comment", 9},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.PnlUsdDlt()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestPnlUsdDltInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_p"},
		{"p_"},
		{"_pn"},
		{"pn_"},
		{"_pnl"},
		{"pnl_"},
		{"_pnlU"},
		{"pnlU_"},
		{"_pnlUs"},
		{"pnlUs_"},
		{"_pnlUsd"},
		{"pnlUsd_"},
		{"_pnlUsdD"},
		{"pnlUsdD_"},
		{"_pnlUsdDl"},
		{"pnlUsdDl_"},
		{"_pnlUsdDlt"},
		{"pnlUsdDlt_"},
		{"9713p"},
		{"p9713"},
		{"9713pn"},
		{"pn9713"},
		{"9713pnl"},
		{"pnl9713"},
		{"9713pnlU"},
		{"pnlU9713"},
		{"9713pnlUs"},
		{"pnlUs9713"},
		{"9713pnlUsd"},
		{"pnlUsd9713"},
		{"9713pnlUsdD"},
		{"pnlUsdD9713"},
		{"9713pnlUsdDl"},
		{"pnlUsdDl9713"},
		{"9713pnlUsdDlt"},
		{"pnlUsdDlt9713"},
		{"Zp"},
		{"pZ"},
		{"Zpn"},
		{"pnZ"},
		{"Zpnl"},
		{"pnlZ"},
		{"ZpnlU"},
		{"pnlUZ"},
		{"ZpnlUs"},
		{"pnlUsZ"},
		{"ZpnlUsd"},
		{"pnlUsdZ"},
		{"ZpnlUsdD"},
		{"pnlUsdDZ"},
		{"ZpnlUsdDl"},
		{"pnlUsdDlZ"},
		{"ZpnlUsdDlt"},
		{"pnlUsdDltZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.PnlUsdDlt()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestPipAvgAValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"pipAvgA", 7},
		{"pipAvgA ", 7},
		{"pipAvgA\n", 7},
		{"pipAvgA.", 7},
		{"pipAvgA:", 7},
		{"pipAvgA,", 7},
		{"pipAvgA\"", 7},
		{"pipAvgA(", 7},
		{"pipAvgA)", 7},
		{"pipAvgA[", 7},
		{"pipAvgA]", 7},
		{"pipAvgA// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.PipAvgA()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestPipAvgAInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_p"},
		{"p_"},
		{"_pi"},
		{"pi_"},
		{"_pip"},
		{"pip_"},
		{"_pipA"},
		{"pipA_"},
		{"_pipAv"},
		{"pipAv_"},
		{"_pipAvg"},
		{"pipAvg_"},
		{"_pipAvgA"},
		{"pipAvgA_"},
		{"9713p"},
		{"p9713"},
		{"9713pi"},
		{"pi9713"},
		{"9713pip"},
		{"pip9713"},
		{"9713pipA"},
		{"pipA9713"},
		{"9713pipAv"},
		{"pipAv9713"},
		{"9713pipAvg"},
		{"pipAvg9713"},
		{"9713pipAvgA"},
		{"pipAvgA9713"},
		{"Zp"},
		{"pZ"},
		{"Zpi"},
		{"piZ"},
		{"Zpip"},
		{"pipZ"},
		{"ZpipA"},
		{"pipAZ"},
		{"ZpipAv"},
		{"pipAvZ"},
		{"ZpipAvg"},
		{"pipAvgZ"},
		{"ZpipAvgA"},
		{"pipAvgAZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.PipAvgA()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestPipAvgBValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"pipAvgB", 7},
		{"pipAvgB ", 7},
		{"pipAvgB\n", 7},
		{"pipAvgB.", 7},
		{"pipAvgB:", 7},
		{"pipAvgB,", 7},
		{"pipAvgB\"", 7},
		{"pipAvgB(", 7},
		{"pipAvgB)", 7},
		{"pipAvgB[", 7},
		{"pipAvgB]", 7},
		{"pipAvgB// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.PipAvgB()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestPipAvgBInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_p"},
		{"p_"},
		{"_pi"},
		{"pi_"},
		{"_pip"},
		{"pip_"},
		{"_pipA"},
		{"pipA_"},
		{"_pipAv"},
		{"pipAv_"},
		{"_pipAvg"},
		{"pipAvg_"},
		{"_pipAvgB"},
		{"pipAvgB_"},
		{"9713p"},
		{"p9713"},
		{"9713pi"},
		{"pi9713"},
		{"9713pip"},
		{"pip9713"},
		{"9713pipA"},
		{"pipA9713"},
		{"9713pipAv"},
		{"pipAv9713"},
		{"9713pipAvg"},
		{"pipAvg9713"},
		{"9713pipAvgB"},
		{"pipAvgB9713"},
		{"Zp"},
		{"pZ"},
		{"Zpi"},
		{"piZ"},
		{"Zpip"},
		{"pipZ"},
		{"ZpipA"},
		{"pipAZ"},
		{"ZpipAv"},
		{"pipAvZ"},
		{"ZpipAvg"},
		{"pipAvgZ"},
		{"ZpipAvgB"},
		{"pipAvgBZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.PipAvgB()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestPipAvgDltValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"pipAvgDlt", 9},
		{"pipAvgDlt ", 9},
		{"pipAvgDlt\n", 9},
		{"pipAvgDlt.", 9},
		{"pipAvgDlt:", 9},
		{"pipAvgDlt,", 9},
		{"pipAvgDlt\"", 9},
		{"pipAvgDlt(", 9},
		{"pipAvgDlt)", 9},
		{"pipAvgDlt[", 9},
		{"pipAvgDlt]", 9},
		{"pipAvgDlt// comment", 9},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.PipAvgDlt()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestPipAvgDltInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_p"},
		{"p_"},
		{"_pi"},
		{"pi_"},
		{"_pip"},
		{"pip_"},
		{"_pipA"},
		{"pipA_"},
		{"_pipAv"},
		{"pipAv_"},
		{"_pipAvg"},
		{"pipAvg_"},
		{"_pipAvgD"},
		{"pipAvgD_"},
		{"_pipAvgDl"},
		{"pipAvgDl_"},
		{"_pipAvgDlt"},
		{"pipAvgDlt_"},
		{"9713p"},
		{"p9713"},
		{"9713pi"},
		{"pi9713"},
		{"9713pip"},
		{"pip9713"},
		{"9713pipA"},
		{"pipA9713"},
		{"9713pipAv"},
		{"pipAv9713"},
		{"9713pipAvg"},
		{"pipAvg9713"},
		{"9713pipAvgD"},
		{"pipAvgD9713"},
		{"9713pipAvgDl"},
		{"pipAvgDl9713"},
		{"9713pipAvgDlt"},
		{"pipAvgDlt9713"},
		{"Zp"},
		{"pZ"},
		{"Zpi"},
		{"piZ"},
		{"Zpip"},
		{"pipZ"},
		{"ZpipA"},
		{"pipAZ"},
		{"ZpipAv"},
		{"pipAvZ"},
		{"ZpipAvg"},
		{"pipAvgZ"},
		{"ZpipAvgD"},
		{"pipAvgDZ"},
		{"ZpipAvgDl"},
		{"pipAvgDlZ"},
		{"ZpipAvgDlt"},
		{"pipAvgDltZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.PipAvgDlt()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestPipMdnAValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"pipMdnA", 7},
		{"pipMdnA ", 7},
		{"pipMdnA\n", 7},
		{"pipMdnA.", 7},
		{"pipMdnA:", 7},
		{"pipMdnA,", 7},
		{"pipMdnA\"", 7},
		{"pipMdnA(", 7},
		{"pipMdnA)", 7},
		{"pipMdnA[", 7},
		{"pipMdnA]", 7},
		{"pipMdnA// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.PipMdnA()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestPipMdnAInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_p"},
		{"p_"},
		{"_pi"},
		{"pi_"},
		{"_pip"},
		{"pip_"},
		{"_pipM"},
		{"pipM_"},
		{"_pipMd"},
		{"pipMd_"},
		{"_pipMdn"},
		{"pipMdn_"},
		{"_pipMdnA"},
		{"pipMdnA_"},
		{"9713p"},
		{"p9713"},
		{"9713pi"},
		{"pi9713"},
		{"9713pip"},
		{"pip9713"},
		{"9713pipM"},
		{"pipM9713"},
		{"9713pipMd"},
		{"pipMd9713"},
		{"9713pipMdn"},
		{"pipMdn9713"},
		{"9713pipMdnA"},
		{"pipMdnA9713"},
		{"Zp"},
		{"pZ"},
		{"Zpi"},
		{"piZ"},
		{"Zpip"},
		{"pipZ"},
		{"ZpipM"},
		{"pipMZ"},
		{"ZpipMd"},
		{"pipMdZ"},
		{"ZpipMdn"},
		{"pipMdnZ"},
		{"ZpipMdnA"},
		{"pipMdnAZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.PipMdnA()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestPipMdnBValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"pipMdnB", 7},
		{"pipMdnB ", 7},
		{"pipMdnB\n", 7},
		{"pipMdnB.", 7},
		{"pipMdnB:", 7},
		{"pipMdnB,", 7},
		{"pipMdnB\"", 7},
		{"pipMdnB(", 7},
		{"pipMdnB)", 7},
		{"pipMdnB[", 7},
		{"pipMdnB]", 7},
		{"pipMdnB// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.PipMdnB()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestPipMdnBInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_p"},
		{"p_"},
		{"_pi"},
		{"pi_"},
		{"_pip"},
		{"pip_"},
		{"_pipM"},
		{"pipM_"},
		{"_pipMd"},
		{"pipMd_"},
		{"_pipMdn"},
		{"pipMdn_"},
		{"_pipMdnB"},
		{"pipMdnB_"},
		{"9713p"},
		{"p9713"},
		{"9713pi"},
		{"pi9713"},
		{"9713pip"},
		{"pip9713"},
		{"9713pipM"},
		{"pipM9713"},
		{"9713pipMd"},
		{"pipMd9713"},
		{"9713pipMdn"},
		{"pipMdn9713"},
		{"9713pipMdnB"},
		{"pipMdnB9713"},
		{"Zp"},
		{"pZ"},
		{"Zpi"},
		{"piZ"},
		{"Zpip"},
		{"pipZ"},
		{"ZpipM"},
		{"pipMZ"},
		{"ZpipMd"},
		{"pipMdZ"},
		{"ZpipMdn"},
		{"pipMdnZ"},
		{"ZpipMdnB"},
		{"pipMdnBZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.PipMdnB()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestPipMdnDltValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"pipMdnDlt", 9},
		{"pipMdnDlt ", 9},
		{"pipMdnDlt\n", 9},
		{"pipMdnDlt.", 9},
		{"pipMdnDlt:", 9},
		{"pipMdnDlt,", 9},
		{"pipMdnDlt\"", 9},
		{"pipMdnDlt(", 9},
		{"pipMdnDlt)", 9},
		{"pipMdnDlt[", 9},
		{"pipMdnDlt]", 9},
		{"pipMdnDlt// comment", 9},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.PipMdnDlt()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestPipMdnDltInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_p"},
		{"p_"},
		{"_pi"},
		{"pi_"},
		{"_pip"},
		{"pip_"},
		{"_pipM"},
		{"pipM_"},
		{"_pipMd"},
		{"pipMd_"},
		{"_pipMdn"},
		{"pipMdn_"},
		{"_pipMdnD"},
		{"pipMdnD_"},
		{"_pipMdnDl"},
		{"pipMdnDl_"},
		{"_pipMdnDlt"},
		{"pipMdnDlt_"},
		{"9713p"},
		{"p9713"},
		{"9713pi"},
		{"pi9713"},
		{"9713pip"},
		{"pip9713"},
		{"9713pipM"},
		{"pipM9713"},
		{"9713pipMd"},
		{"pipMd9713"},
		{"9713pipMdn"},
		{"pipMdn9713"},
		{"9713pipMdnD"},
		{"pipMdnD9713"},
		{"9713pipMdnDl"},
		{"pipMdnDl9713"},
		{"9713pipMdnDlt"},
		{"pipMdnDlt9713"},
		{"Zp"},
		{"pZ"},
		{"Zpi"},
		{"piZ"},
		{"Zpip"},
		{"pipZ"},
		{"ZpipM"},
		{"pipMZ"},
		{"ZpipMd"},
		{"pipMdZ"},
		{"ZpipMdn"},
		{"pipMdnZ"},
		{"ZpipMdnD"},
		{"pipMdnDZ"},
		{"ZpipMdnDl"},
		{"pipMdnDlZ"},
		{"ZpipMdnDlt"},
		{"pipMdnDltZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.PipMdnDlt()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestPipMinAValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"pipMinA", 7},
		{"pipMinA ", 7},
		{"pipMinA\n", 7},
		{"pipMinA.", 7},
		{"pipMinA:", 7},
		{"pipMinA,", 7},
		{"pipMinA\"", 7},
		{"pipMinA(", 7},
		{"pipMinA)", 7},
		{"pipMinA[", 7},
		{"pipMinA]", 7},
		{"pipMinA// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.PipMinA()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestPipMinAInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_p"},
		{"p_"},
		{"_pi"},
		{"pi_"},
		{"_pip"},
		{"pip_"},
		{"_pipM"},
		{"pipM_"},
		{"_pipMi"},
		{"pipMi_"},
		{"_pipMin"},
		{"pipMin_"},
		{"_pipMinA"},
		{"pipMinA_"},
		{"9713p"},
		{"p9713"},
		{"9713pi"},
		{"pi9713"},
		{"9713pip"},
		{"pip9713"},
		{"9713pipM"},
		{"pipM9713"},
		{"9713pipMi"},
		{"pipMi9713"},
		{"9713pipMin"},
		{"pipMin9713"},
		{"9713pipMinA"},
		{"pipMinA9713"},
		{"Zp"},
		{"pZ"},
		{"Zpi"},
		{"piZ"},
		{"Zpip"},
		{"pipZ"},
		{"ZpipM"},
		{"pipMZ"},
		{"ZpipMi"},
		{"pipMiZ"},
		{"ZpipMin"},
		{"pipMinZ"},
		{"ZpipMinA"},
		{"pipMinAZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.PipMinA()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestPipMinBValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"pipMinB", 7},
		{"pipMinB ", 7},
		{"pipMinB\n", 7},
		{"pipMinB.", 7},
		{"pipMinB:", 7},
		{"pipMinB,", 7},
		{"pipMinB\"", 7},
		{"pipMinB(", 7},
		{"pipMinB)", 7},
		{"pipMinB[", 7},
		{"pipMinB]", 7},
		{"pipMinB// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.PipMinB()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestPipMinBInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_p"},
		{"p_"},
		{"_pi"},
		{"pi_"},
		{"_pip"},
		{"pip_"},
		{"_pipM"},
		{"pipM_"},
		{"_pipMi"},
		{"pipMi_"},
		{"_pipMin"},
		{"pipMin_"},
		{"_pipMinB"},
		{"pipMinB_"},
		{"9713p"},
		{"p9713"},
		{"9713pi"},
		{"pi9713"},
		{"9713pip"},
		{"pip9713"},
		{"9713pipM"},
		{"pipM9713"},
		{"9713pipMi"},
		{"pipMi9713"},
		{"9713pipMin"},
		{"pipMin9713"},
		{"9713pipMinB"},
		{"pipMinB9713"},
		{"Zp"},
		{"pZ"},
		{"Zpi"},
		{"piZ"},
		{"Zpip"},
		{"pipZ"},
		{"ZpipM"},
		{"pipMZ"},
		{"ZpipMi"},
		{"pipMiZ"},
		{"ZpipMin"},
		{"pipMinZ"},
		{"ZpipMinB"},
		{"pipMinBZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.PipMinB()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestPipMinDltValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"pipMinDlt", 9},
		{"pipMinDlt ", 9},
		{"pipMinDlt\n", 9},
		{"pipMinDlt.", 9},
		{"pipMinDlt:", 9},
		{"pipMinDlt,", 9},
		{"pipMinDlt\"", 9},
		{"pipMinDlt(", 9},
		{"pipMinDlt)", 9},
		{"pipMinDlt[", 9},
		{"pipMinDlt]", 9},
		{"pipMinDlt// comment", 9},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.PipMinDlt()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestPipMinDltInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_p"},
		{"p_"},
		{"_pi"},
		{"pi_"},
		{"_pip"},
		{"pip_"},
		{"_pipM"},
		{"pipM_"},
		{"_pipMi"},
		{"pipMi_"},
		{"_pipMin"},
		{"pipMin_"},
		{"_pipMinD"},
		{"pipMinD_"},
		{"_pipMinDl"},
		{"pipMinDl_"},
		{"_pipMinDlt"},
		{"pipMinDlt_"},
		{"9713p"},
		{"p9713"},
		{"9713pi"},
		{"pi9713"},
		{"9713pip"},
		{"pip9713"},
		{"9713pipM"},
		{"pipM9713"},
		{"9713pipMi"},
		{"pipMi9713"},
		{"9713pipMin"},
		{"pipMin9713"},
		{"9713pipMinD"},
		{"pipMinD9713"},
		{"9713pipMinDl"},
		{"pipMinDl9713"},
		{"9713pipMinDlt"},
		{"pipMinDlt9713"},
		{"Zp"},
		{"pZ"},
		{"Zpi"},
		{"piZ"},
		{"Zpip"},
		{"pipZ"},
		{"ZpipM"},
		{"pipMZ"},
		{"ZpipMi"},
		{"pipMiZ"},
		{"ZpipMin"},
		{"pipMinZ"},
		{"ZpipMinD"},
		{"pipMinDZ"},
		{"ZpipMinDl"},
		{"pipMinDlZ"},
		{"ZpipMinDlt"},
		{"pipMinDltZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.PipMinDlt()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestPipMaxAValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"pipMaxA", 7},
		{"pipMaxA ", 7},
		{"pipMaxA\n", 7},
		{"pipMaxA.", 7},
		{"pipMaxA:", 7},
		{"pipMaxA,", 7},
		{"pipMaxA\"", 7},
		{"pipMaxA(", 7},
		{"pipMaxA)", 7},
		{"pipMaxA[", 7},
		{"pipMaxA]", 7},
		{"pipMaxA// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.PipMaxA()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestPipMaxAInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_p"},
		{"p_"},
		{"_pi"},
		{"pi_"},
		{"_pip"},
		{"pip_"},
		{"_pipM"},
		{"pipM_"},
		{"_pipMa"},
		{"pipMa_"},
		{"_pipMax"},
		{"pipMax_"},
		{"_pipMaxA"},
		{"pipMaxA_"},
		{"9713p"},
		{"p9713"},
		{"9713pi"},
		{"pi9713"},
		{"9713pip"},
		{"pip9713"},
		{"9713pipM"},
		{"pipM9713"},
		{"9713pipMa"},
		{"pipMa9713"},
		{"9713pipMax"},
		{"pipMax9713"},
		{"9713pipMaxA"},
		{"pipMaxA9713"},
		{"Zp"},
		{"pZ"},
		{"Zpi"},
		{"piZ"},
		{"Zpip"},
		{"pipZ"},
		{"ZpipM"},
		{"pipMZ"},
		{"ZpipMa"},
		{"pipMaZ"},
		{"ZpipMax"},
		{"pipMaxZ"},
		{"ZpipMaxA"},
		{"pipMaxAZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.PipMaxA()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestPipMaxBValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"pipMaxB", 7},
		{"pipMaxB ", 7},
		{"pipMaxB\n", 7},
		{"pipMaxB.", 7},
		{"pipMaxB:", 7},
		{"pipMaxB,", 7},
		{"pipMaxB\"", 7},
		{"pipMaxB(", 7},
		{"pipMaxB)", 7},
		{"pipMaxB[", 7},
		{"pipMaxB]", 7},
		{"pipMaxB// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.PipMaxB()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestPipMaxBInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_p"},
		{"p_"},
		{"_pi"},
		{"pi_"},
		{"_pip"},
		{"pip_"},
		{"_pipM"},
		{"pipM_"},
		{"_pipMa"},
		{"pipMa_"},
		{"_pipMax"},
		{"pipMax_"},
		{"_pipMaxB"},
		{"pipMaxB_"},
		{"9713p"},
		{"p9713"},
		{"9713pi"},
		{"pi9713"},
		{"9713pip"},
		{"pip9713"},
		{"9713pipM"},
		{"pipM9713"},
		{"9713pipMa"},
		{"pipMa9713"},
		{"9713pipMax"},
		{"pipMax9713"},
		{"9713pipMaxB"},
		{"pipMaxB9713"},
		{"Zp"},
		{"pZ"},
		{"Zpi"},
		{"piZ"},
		{"Zpip"},
		{"pipZ"},
		{"ZpipM"},
		{"pipMZ"},
		{"ZpipMa"},
		{"pipMaZ"},
		{"ZpipMax"},
		{"pipMaxZ"},
		{"ZpipMaxB"},
		{"pipMaxBZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.PipMaxB()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestPipMaxDltValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"pipMaxDlt", 9},
		{"pipMaxDlt ", 9},
		{"pipMaxDlt\n", 9},
		{"pipMaxDlt.", 9},
		{"pipMaxDlt:", 9},
		{"pipMaxDlt,", 9},
		{"pipMaxDlt\"", 9},
		{"pipMaxDlt(", 9},
		{"pipMaxDlt)", 9},
		{"pipMaxDlt[", 9},
		{"pipMaxDlt]", 9},
		{"pipMaxDlt// comment", 9},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.PipMaxDlt()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestPipMaxDltInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_p"},
		{"p_"},
		{"_pi"},
		{"pi_"},
		{"_pip"},
		{"pip_"},
		{"_pipM"},
		{"pipM_"},
		{"_pipMa"},
		{"pipMa_"},
		{"_pipMax"},
		{"pipMax_"},
		{"_pipMaxD"},
		{"pipMaxD_"},
		{"_pipMaxDl"},
		{"pipMaxDl_"},
		{"_pipMaxDlt"},
		{"pipMaxDlt_"},
		{"9713p"},
		{"p9713"},
		{"9713pi"},
		{"pi9713"},
		{"9713pip"},
		{"pip9713"},
		{"9713pipM"},
		{"pipM9713"},
		{"9713pipMa"},
		{"pipMa9713"},
		{"9713pipMax"},
		{"pipMax9713"},
		{"9713pipMaxD"},
		{"pipMaxD9713"},
		{"9713pipMaxDl"},
		{"pipMaxDl9713"},
		{"9713pipMaxDlt"},
		{"pipMaxDlt9713"},
		{"Zp"},
		{"pZ"},
		{"Zpi"},
		{"piZ"},
		{"Zpip"},
		{"pipZ"},
		{"ZpipM"},
		{"pipMZ"},
		{"ZpipMa"},
		{"pipMaZ"},
		{"ZpipMax"},
		{"pipMaxZ"},
		{"ZpipMaxD"},
		{"pipMaxDZ"},
		{"ZpipMaxDl"},
		{"pipMaxDlZ"},
		{"ZpipMaxDlt"},
		{"pipMaxDltZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.PipMaxDlt()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestPipSumAValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"pipSumA", 7},
		{"pipSumA ", 7},
		{"pipSumA\n", 7},
		{"pipSumA.", 7},
		{"pipSumA:", 7},
		{"pipSumA,", 7},
		{"pipSumA\"", 7},
		{"pipSumA(", 7},
		{"pipSumA)", 7},
		{"pipSumA[", 7},
		{"pipSumA]", 7},
		{"pipSumA// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.PipSumA()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestPipSumAInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_p"},
		{"p_"},
		{"_pi"},
		{"pi_"},
		{"_pip"},
		{"pip_"},
		{"_pipS"},
		{"pipS_"},
		{"_pipSu"},
		{"pipSu_"},
		{"_pipSum"},
		{"pipSum_"},
		{"_pipSumA"},
		{"pipSumA_"},
		{"9713p"},
		{"p9713"},
		{"9713pi"},
		{"pi9713"},
		{"9713pip"},
		{"pip9713"},
		{"9713pipS"},
		{"pipS9713"},
		{"9713pipSu"},
		{"pipSu9713"},
		{"9713pipSum"},
		{"pipSum9713"},
		{"9713pipSumA"},
		{"pipSumA9713"},
		{"Zp"},
		{"pZ"},
		{"Zpi"},
		{"piZ"},
		{"Zpip"},
		{"pipZ"},
		{"ZpipS"},
		{"pipSZ"},
		{"ZpipSu"},
		{"pipSuZ"},
		{"ZpipSum"},
		{"pipSumZ"},
		{"ZpipSumA"},
		{"pipSumAZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.PipSumA()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestPipSumBValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"pipSumB", 7},
		{"pipSumB ", 7},
		{"pipSumB\n", 7},
		{"pipSumB.", 7},
		{"pipSumB:", 7},
		{"pipSumB,", 7},
		{"pipSumB\"", 7},
		{"pipSumB(", 7},
		{"pipSumB)", 7},
		{"pipSumB[", 7},
		{"pipSumB]", 7},
		{"pipSumB// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.PipSumB()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestPipSumBInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_p"},
		{"p_"},
		{"_pi"},
		{"pi_"},
		{"_pip"},
		{"pip_"},
		{"_pipS"},
		{"pipS_"},
		{"_pipSu"},
		{"pipSu_"},
		{"_pipSum"},
		{"pipSum_"},
		{"_pipSumB"},
		{"pipSumB_"},
		{"9713p"},
		{"p9713"},
		{"9713pi"},
		{"pi9713"},
		{"9713pip"},
		{"pip9713"},
		{"9713pipS"},
		{"pipS9713"},
		{"9713pipSu"},
		{"pipSu9713"},
		{"9713pipSum"},
		{"pipSum9713"},
		{"9713pipSumB"},
		{"pipSumB9713"},
		{"Zp"},
		{"pZ"},
		{"Zpi"},
		{"piZ"},
		{"Zpip"},
		{"pipZ"},
		{"ZpipS"},
		{"pipSZ"},
		{"ZpipSu"},
		{"pipSuZ"},
		{"ZpipSum"},
		{"pipSumZ"},
		{"ZpipSumB"},
		{"pipSumBZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.PipSumB()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestPipSumDltValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"pipSumDlt", 9},
		{"pipSumDlt ", 9},
		{"pipSumDlt\n", 9},
		{"pipSumDlt.", 9},
		{"pipSumDlt:", 9},
		{"pipSumDlt,", 9},
		{"pipSumDlt\"", 9},
		{"pipSumDlt(", 9},
		{"pipSumDlt)", 9},
		{"pipSumDlt[", 9},
		{"pipSumDlt]", 9},
		{"pipSumDlt// comment", 9},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.PipSumDlt()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestPipSumDltInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_p"},
		{"p_"},
		{"_pi"},
		{"pi_"},
		{"_pip"},
		{"pip_"},
		{"_pipS"},
		{"pipS_"},
		{"_pipSu"},
		{"pipSu_"},
		{"_pipSum"},
		{"pipSum_"},
		{"_pipSumD"},
		{"pipSumD_"},
		{"_pipSumDl"},
		{"pipSumDl_"},
		{"_pipSumDlt"},
		{"pipSumDlt_"},
		{"9713p"},
		{"p9713"},
		{"9713pi"},
		{"pi9713"},
		{"9713pip"},
		{"pip9713"},
		{"9713pipS"},
		{"pipS9713"},
		{"9713pipSu"},
		{"pipSu9713"},
		{"9713pipSum"},
		{"pipSum9713"},
		{"9713pipSumD"},
		{"pipSumD9713"},
		{"9713pipSumDl"},
		{"pipSumDl9713"},
		{"9713pipSumDlt"},
		{"pipSumDlt9713"},
		{"Zp"},
		{"pZ"},
		{"Zpi"},
		{"piZ"},
		{"Zpip"},
		{"pipZ"},
		{"ZpipS"},
		{"pipSZ"},
		{"ZpipSu"},
		{"pipSuZ"},
		{"ZpipSum"},
		{"pipSumZ"},
		{"ZpipSumD"},
		{"pipSumDZ"},
		{"ZpipSumDl"},
		{"pipSumDlZ"},
		{"ZpipSumDlt"},
		{"pipSumDltZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.PipSumDlt()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestDurAvgAValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"durAvgA", 7},
		{"durAvgA ", 7},
		{"durAvgA\n", 7},
		{"durAvgA.", 7},
		{"durAvgA:", 7},
		{"durAvgA,", 7},
		{"durAvgA\"", 7},
		{"durAvgA(", 7},
		{"durAvgA)", 7},
		{"durAvgA[", 7},
		{"durAvgA]", 7},
		{"durAvgA// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.DurAvgA()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestDurAvgAInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_d"},
		{"d_"},
		{"_du"},
		{"du_"},
		{"_dur"},
		{"dur_"},
		{"_durA"},
		{"durA_"},
		{"_durAv"},
		{"durAv_"},
		{"_durAvg"},
		{"durAvg_"},
		{"_durAvgA"},
		{"durAvgA_"},
		{"9713d"},
		{"d9713"},
		{"9713du"},
		{"du9713"},
		{"9713dur"},
		{"dur9713"},
		{"9713durA"},
		{"durA9713"},
		{"9713durAv"},
		{"durAv9713"},
		{"9713durAvg"},
		{"durAvg9713"},
		{"9713durAvgA"},
		{"durAvgA9713"},
		{"Zd"},
		{"dZ"},
		{"Zdu"},
		{"duZ"},
		{"Zdur"},
		{"durZ"},
		{"ZdurA"},
		{"durAZ"},
		{"ZdurAv"},
		{"durAvZ"},
		{"ZdurAvg"},
		{"durAvgZ"},
		{"ZdurAvgA"},
		{"durAvgAZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.DurAvgA()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestDurAvgBValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"durAvgB", 7},
		{"durAvgB ", 7},
		{"durAvgB\n", 7},
		{"durAvgB.", 7},
		{"durAvgB:", 7},
		{"durAvgB,", 7},
		{"durAvgB\"", 7},
		{"durAvgB(", 7},
		{"durAvgB)", 7},
		{"durAvgB[", 7},
		{"durAvgB]", 7},
		{"durAvgB// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.DurAvgB()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestDurAvgBInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_d"},
		{"d_"},
		{"_du"},
		{"du_"},
		{"_dur"},
		{"dur_"},
		{"_durA"},
		{"durA_"},
		{"_durAv"},
		{"durAv_"},
		{"_durAvg"},
		{"durAvg_"},
		{"_durAvgB"},
		{"durAvgB_"},
		{"9713d"},
		{"d9713"},
		{"9713du"},
		{"du9713"},
		{"9713dur"},
		{"dur9713"},
		{"9713durA"},
		{"durA9713"},
		{"9713durAv"},
		{"durAv9713"},
		{"9713durAvg"},
		{"durAvg9713"},
		{"9713durAvgB"},
		{"durAvgB9713"},
		{"Zd"},
		{"dZ"},
		{"Zdu"},
		{"duZ"},
		{"Zdur"},
		{"durZ"},
		{"ZdurA"},
		{"durAZ"},
		{"ZdurAv"},
		{"durAvZ"},
		{"ZdurAvg"},
		{"durAvgZ"},
		{"ZdurAvgB"},
		{"durAvgBZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.DurAvgB()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestDurAvgDltValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"durAvgDlt", 9},
		{"durAvgDlt ", 9},
		{"durAvgDlt\n", 9},
		{"durAvgDlt.", 9},
		{"durAvgDlt:", 9},
		{"durAvgDlt,", 9},
		{"durAvgDlt\"", 9},
		{"durAvgDlt(", 9},
		{"durAvgDlt)", 9},
		{"durAvgDlt[", 9},
		{"durAvgDlt]", 9},
		{"durAvgDlt// comment", 9},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.DurAvgDlt()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestDurAvgDltInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_d"},
		{"d_"},
		{"_du"},
		{"du_"},
		{"_dur"},
		{"dur_"},
		{"_durA"},
		{"durA_"},
		{"_durAv"},
		{"durAv_"},
		{"_durAvg"},
		{"durAvg_"},
		{"_durAvgD"},
		{"durAvgD_"},
		{"_durAvgDl"},
		{"durAvgDl_"},
		{"_durAvgDlt"},
		{"durAvgDlt_"},
		{"9713d"},
		{"d9713"},
		{"9713du"},
		{"du9713"},
		{"9713dur"},
		{"dur9713"},
		{"9713durA"},
		{"durA9713"},
		{"9713durAv"},
		{"durAv9713"},
		{"9713durAvg"},
		{"durAvg9713"},
		{"9713durAvgD"},
		{"durAvgD9713"},
		{"9713durAvgDl"},
		{"durAvgDl9713"},
		{"9713durAvgDlt"},
		{"durAvgDlt9713"},
		{"Zd"},
		{"dZ"},
		{"Zdu"},
		{"duZ"},
		{"Zdur"},
		{"durZ"},
		{"ZdurA"},
		{"durAZ"},
		{"ZdurAv"},
		{"durAvZ"},
		{"ZdurAvg"},
		{"durAvgZ"},
		{"ZdurAvgD"},
		{"durAvgDZ"},
		{"ZdurAvgDl"},
		{"durAvgDlZ"},
		{"ZdurAvgDlt"},
		{"durAvgDltZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.DurAvgDlt()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestDurMdnAValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"durMdnA", 7},
		{"durMdnA ", 7},
		{"durMdnA\n", 7},
		{"durMdnA.", 7},
		{"durMdnA:", 7},
		{"durMdnA,", 7},
		{"durMdnA\"", 7},
		{"durMdnA(", 7},
		{"durMdnA)", 7},
		{"durMdnA[", 7},
		{"durMdnA]", 7},
		{"durMdnA// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.DurMdnA()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestDurMdnAInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_d"},
		{"d_"},
		{"_du"},
		{"du_"},
		{"_dur"},
		{"dur_"},
		{"_durM"},
		{"durM_"},
		{"_durMd"},
		{"durMd_"},
		{"_durMdn"},
		{"durMdn_"},
		{"_durMdnA"},
		{"durMdnA_"},
		{"9713d"},
		{"d9713"},
		{"9713du"},
		{"du9713"},
		{"9713dur"},
		{"dur9713"},
		{"9713durM"},
		{"durM9713"},
		{"9713durMd"},
		{"durMd9713"},
		{"9713durMdn"},
		{"durMdn9713"},
		{"9713durMdnA"},
		{"durMdnA9713"},
		{"Zd"},
		{"dZ"},
		{"Zdu"},
		{"duZ"},
		{"Zdur"},
		{"durZ"},
		{"ZdurM"},
		{"durMZ"},
		{"ZdurMd"},
		{"durMdZ"},
		{"ZdurMdn"},
		{"durMdnZ"},
		{"ZdurMdnA"},
		{"durMdnAZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.DurMdnA()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestDurMdnBValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"durMdnB", 7},
		{"durMdnB ", 7},
		{"durMdnB\n", 7},
		{"durMdnB.", 7},
		{"durMdnB:", 7},
		{"durMdnB,", 7},
		{"durMdnB\"", 7},
		{"durMdnB(", 7},
		{"durMdnB)", 7},
		{"durMdnB[", 7},
		{"durMdnB]", 7},
		{"durMdnB// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.DurMdnB()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestDurMdnBInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_d"},
		{"d_"},
		{"_du"},
		{"du_"},
		{"_dur"},
		{"dur_"},
		{"_durM"},
		{"durM_"},
		{"_durMd"},
		{"durMd_"},
		{"_durMdn"},
		{"durMdn_"},
		{"_durMdnB"},
		{"durMdnB_"},
		{"9713d"},
		{"d9713"},
		{"9713du"},
		{"du9713"},
		{"9713dur"},
		{"dur9713"},
		{"9713durM"},
		{"durM9713"},
		{"9713durMd"},
		{"durMd9713"},
		{"9713durMdn"},
		{"durMdn9713"},
		{"9713durMdnB"},
		{"durMdnB9713"},
		{"Zd"},
		{"dZ"},
		{"Zdu"},
		{"duZ"},
		{"Zdur"},
		{"durZ"},
		{"ZdurM"},
		{"durMZ"},
		{"ZdurMd"},
		{"durMdZ"},
		{"ZdurMdn"},
		{"durMdnZ"},
		{"ZdurMdnB"},
		{"durMdnBZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.DurMdnB()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestDurMdnDltValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"durMdnDlt", 9},
		{"durMdnDlt ", 9},
		{"durMdnDlt\n", 9},
		{"durMdnDlt.", 9},
		{"durMdnDlt:", 9},
		{"durMdnDlt,", 9},
		{"durMdnDlt\"", 9},
		{"durMdnDlt(", 9},
		{"durMdnDlt)", 9},
		{"durMdnDlt[", 9},
		{"durMdnDlt]", 9},
		{"durMdnDlt// comment", 9},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.DurMdnDlt()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestDurMdnDltInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_d"},
		{"d_"},
		{"_du"},
		{"du_"},
		{"_dur"},
		{"dur_"},
		{"_durM"},
		{"durM_"},
		{"_durMd"},
		{"durMd_"},
		{"_durMdn"},
		{"durMdn_"},
		{"_durMdnD"},
		{"durMdnD_"},
		{"_durMdnDl"},
		{"durMdnDl_"},
		{"_durMdnDlt"},
		{"durMdnDlt_"},
		{"9713d"},
		{"d9713"},
		{"9713du"},
		{"du9713"},
		{"9713dur"},
		{"dur9713"},
		{"9713durM"},
		{"durM9713"},
		{"9713durMd"},
		{"durMd9713"},
		{"9713durMdn"},
		{"durMdn9713"},
		{"9713durMdnD"},
		{"durMdnD9713"},
		{"9713durMdnDl"},
		{"durMdnDl9713"},
		{"9713durMdnDlt"},
		{"durMdnDlt9713"},
		{"Zd"},
		{"dZ"},
		{"Zdu"},
		{"duZ"},
		{"Zdur"},
		{"durZ"},
		{"ZdurM"},
		{"durMZ"},
		{"ZdurMd"},
		{"durMdZ"},
		{"ZdurMdn"},
		{"durMdnZ"},
		{"ZdurMdnD"},
		{"durMdnDZ"},
		{"ZdurMdnDl"},
		{"durMdnDlZ"},
		{"ZdurMdnDlt"},
		{"durMdnDltZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.DurMdnDlt()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestDurMinAValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"durMinA", 7},
		{"durMinA ", 7},
		{"durMinA\n", 7},
		{"durMinA.", 7},
		{"durMinA:", 7},
		{"durMinA,", 7},
		{"durMinA\"", 7},
		{"durMinA(", 7},
		{"durMinA)", 7},
		{"durMinA[", 7},
		{"durMinA]", 7},
		{"durMinA// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.DurMinA()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestDurMinAInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_d"},
		{"d_"},
		{"_du"},
		{"du_"},
		{"_dur"},
		{"dur_"},
		{"_durM"},
		{"durM_"},
		{"_durMi"},
		{"durMi_"},
		{"_durMin"},
		{"durMin_"},
		{"_durMinA"},
		{"durMinA_"},
		{"9713d"},
		{"d9713"},
		{"9713du"},
		{"du9713"},
		{"9713dur"},
		{"dur9713"},
		{"9713durM"},
		{"durM9713"},
		{"9713durMi"},
		{"durMi9713"},
		{"9713durMin"},
		{"durMin9713"},
		{"9713durMinA"},
		{"durMinA9713"},
		{"Zd"},
		{"dZ"},
		{"Zdu"},
		{"duZ"},
		{"Zdur"},
		{"durZ"},
		{"ZdurM"},
		{"durMZ"},
		{"ZdurMi"},
		{"durMiZ"},
		{"ZdurMin"},
		{"durMinZ"},
		{"ZdurMinA"},
		{"durMinAZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.DurMinA()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestDurMinBValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"durMinB", 7},
		{"durMinB ", 7},
		{"durMinB\n", 7},
		{"durMinB.", 7},
		{"durMinB:", 7},
		{"durMinB,", 7},
		{"durMinB\"", 7},
		{"durMinB(", 7},
		{"durMinB)", 7},
		{"durMinB[", 7},
		{"durMinB]", 7},
		{"durMinB// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.DurMinB()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestDurMinBInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_d"},
		{"d_"},
		{"_du"},
		{"du_"},
		{"_dur"},
		{"dur_"},
		{"_durM"},
		{"durM_"},
		{"_durMi"},
		{"durMi_"},
		{"_durMin"},
		{"durMin_"},
		{"_durMinB"},
		{"durMinB_"},
		{"9713d"},
		{"d9713"},
		{"9713du"},
		{"du9713"},
		{"9713dur"},
		{"dur9713"},
		{"9713durM"},
		{"durM9713"},
		{"9713durMi"},
		{"durMi9713"},
		{"9713durMin"},
		{"durMin9713"},
		{"9713durMinB"},
		{"durMinB9713"},
		{"Zd"},
		{"dZ"},
		{"Zdu"},
		{"duZ"},
		{"Zdur"},
		{"durZ"},
		{"ZdurM"},
		{"durMZ"},
		{"ZdurMi"},
		{"durMiZ"},
		{"ZdurMin"},
		{"durMinZ"},
		{"ZdurMinB"},
		{"durMinBZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.DurMinB()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestDurMinDltValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"durMinDlt", 9},
		{"durMinDlt ", 9},
		{"durMinDlt\n", 9},
		{"durMinDlt.", 9},
		{"durMinDlt:", 9},
		{"durMinDlt,", 9},
		{"durMinDlt\"", 9},
		{"durMinDlt(", 9},
		{"durMinDlt)", 9},
		{"durMinDlt[", 9},
		{"durMinDlt]", 9},
		{"durMinDlt// comment", 9},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.DurMinDlt()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestDurMinDltInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_d"},
		{"d_"},
		{"_du"},
		{"du_"},
		{"_dur"},
		{"dur_"},
		{"_durM"},
		{"durM_"},
		{"_durMi"},
		{"durMi_"},
		{"_durMin"},
		{"durMin_"},
		{"_durMinD"},
		{"durMinD_"},
		{"_durMinDl"},
		{"durMinDl_"},
		{"_durMinDlt"},
		{"durMinDlt_"},
		{"9713d"},
		{"d9713"},
		{"9713du"},
		{"du9713"},
		{"9713dur"},
		{"dur9713"},
		{"9713durM"},
		{"durM9713"},
		{"9713durMi"},
		{"durMi9713"},
		{"9713durMin"},
		{"durMin9713"},
		{"9713durMinD"},
		{"durMinD9713"},
		{"9713durMinDl"},
		{"durMinDl9713"},
		{"9713durMinDlt"},
		{"durMinDlt9713"},
		{"Zd"},
		{"dZ"},
		{"Zdu"},
		{"duZ"},
		{"Zdur"},
		{"durZ"},
		{"ZdurM"},
		{"durMZ"},
		{"ZdurMi"},
		{"durMiZ"},
		{"ZdurMin"},
		{"durMinZ"},
		{"ZdurMinD"},
		{"durMinDZ"},
		{"ZdurMinDl"},
		{"durMinDlZ"},
		{"ZdurMinDlt"},
		{"durMinDltZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.DurMinDlt()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestDurMaxAValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"durMaxA", 7},
		{"durMaxA ", 7},
		{"durMaxA\n", 7},
		{"durMaxA.", 7},
		{"durMaxA:", 7},
		{"durMaxA,", 7},
		{"durMaxA\"", 7},
		{"durMaxA(", 7},
		{"durMaxA)", 7},
		{"durMaxA[", 7},
		{"durMaxA]", 7},
		{"durMaxA// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.DurMaxA()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestDurMaxAInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_d"},
		{"d_"},
		{"_du"},
		{"du_"},
		{"_dur"},
		{"dur_"},
		{"_durM"},
		{"durM_"},
		{"_durMa"},
		{"durMa_"},
		{"_durMax"},
		{"durMax_"},
		{"_durMaxA"},
		{"durMaxA_"},
		{"9713d"},
		{"d9713"},
		{"9713du"},
		{"du9713"},
		{"9713dur"},
		{"dur9713"},
		{"9713durM"},
		{"durM9713"},
		{"9713durMa"},
		{"durMa9713"},
		{"9713durMax"},
		{"durMax9713"},
		{"9713durMaxA"},
		{"durMaxA9713"},
		{"Zd"},
		{"dZ"},
		{"Zdu"},
		{"duZ"},
		{"Zdur"},
		{"durZ"},
		{"ZdurM"},
		{"durMZ"},
		{"ZdurMa"},
		{"durMaZ"},
		{"ZdurMax"},
		{"durMaxZ"},
		{"ZdurMaxA"},
		{"durMaxAZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.DurMaxA()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestDurMaxBValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"durMaxB", 7},
		{"durMaxB ", 7},
		{"durMaxB\n", 7},
		{"durMaxB.", 7},
		{"durMaxB:", 7},
		{"durMaxB,", 7},
		{"durMaxB\"", 7},
		{"durMaxB(", 7},
		{"durMaxB)", 7},
		{"durMaxB[", 7},
		{"durMaxB]", 7},
		{"durMaxB// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.DurMaxB()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestDurMaxBInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_d"},
		{"d_"},
		{"_du"},
		{"du_"},
		{"_dur"},
		{"dur_"},
		{"_durM"},
		{"durM_"},
		{"_durMa"},
		{"durMa_"},
		{"_durMax"},
		{"durMax_"},
		{"_durMaxB"},
		{"durMaxB_"},
		{"9713d"},
		{"d9713"},
		{"9713du"},
		{"du9713"},
		{"9713dur"},
		{"dur9713"},
		{"9713durM"},
		{"durM9713"},
		{"9713durMa"},
		{"durMa9713"},
		{"9713durMax"},
		{"durMax9713"},
		{"9713durMaxB"},
		{"durMaxB9713"},
		{"Zd"},
		{"dZ"},
		{"Zdu"},
		{"duZ"},
		{"Zdur"},
		{"durZ"},
		{"ZdurM"},
		{"durMZ"},
		{"ZdurMa"},
		{"durMaZ"},
		{"ZdurMax"},
		{"durMaxZ"},
		{"ZdurMaxB"},
		{"durMaxBZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.DurMaxB()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestDurMaxDltValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"durMaxDlt", 9},
		{"durMaxDlt ", 9},
		{"durMaxDlt\n", 9},
		{"durMaxDlt.", 9},
		{"durMaxDlt:", 9},
		{"durMaxDlt,", 9},
		{"durMaxDlt\"", 9},
		{"durMaxDlt(", 9},
		{"durMaxDlt)", 9},
		{"durMaxDlt[", 9},
		{"durMaxDlt]", 9},
		{"durMaxDlt// comment", 9},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.DurMaxDlt()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestDurMaxDltInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_d"},
		{"d_"},
		{"_du"},
		{"du_"},
		{"_dur"},
		{"dur_"},
		{"_durM"},
		{"durM_"},
		{"_durMa"},
		{"durMa_"},
		{"_durMax"},
		{"durMax_"},
		{"_durMaxD"},
		{"durMaxD_"},
		{"_durMaxDl"},
		{"durMaxDl_"},
		{"_durMaxDlt"},
		{"durMaxDlt_"},
		{"9713d"},
		{"d9713"},
		{"9713du"},
		{"du9713"},
		{"9713dur"},
		{"dur9713"},
		{"9713durM"},
		{"durM9713"},
		{"9713durMa"},
		{"durMa9713"},
		{"9713durMax"},
		{"durMax9713"},
		{"9713durMaxD"},
		{"durMaxD9713"},
		{"9713durMaxDl"},
		{"durMaxDl9713"},
		{"9713durMaxDlt"},
		{"durMaxDlt9713"},
		{"Zd"},
		{"dZ"},
		{"Zdu"},
		{"duZ"},
		{"Zdur"},
		{"durZ"},
		{"ZdurM"},
		{"durMZ"},
		{"ZdurMa"},
		{"durMaZ"},
		{"ZdurMax"},
		{"durMaxZ"},
		{"ZdurMaxD"},
		{"durMaxDZ"},
		{"ZdurMaxDl"},
		{"durMaxDlZ"},
		{"ZdurMaxDlt"},
		{"durMaxDltZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.DurMaxDlt()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestTrdCntAValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"trdCntA", 7},
		{"trdCntA ", 7},
		{"trdCntA\n", 7},
		{"trdCntA.", 7},
		{"trdCntA:", 7},
		{"trdCntA,", 7},
		{"trdCntA\"", 7},
		{"trdCntA(", 7},
		{"trdCntA)", 7},
		{"trdCntA[", 7},
		{"trdCntA]", 7},
		{"trdCntA// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.TrdCntA()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestTrdCntAInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_t"},
		{"t_"},
		{"_tr"},
		{"tr_"},
		{"_trd"},
		{"trd_"},
		{"_trdC"},
		{"trdC_"},
		{"_trdCn"},
		{"trdCn_"},
		{"_trdCnt"},
		{"trdCnt_"},
		{"_trdCntA"},
		{"trdCntA_"},
		{"9713t"},
		{"t9713"},
		{"9713tr"},
		{"tr9713"},
		{"9713trd"},
		{"trd9713"},
		{"9713trdC"},
		{"trdC9713"},
		{"9713trdCn"},
		{"trdCn9713"},
		{"9713trdCnt"},
		{"trdCnt9713"},
		{"9713trdCntA"},
		{"trdCntA9713"},
		{"Zt"},
		{"tZ"},
		{"Ztr"},
		{"trZ"},
		{"Ztrd"},
		{"trdZ"},
		{"ZtrdC"},
		{"trdCZ"},
		{"ZtrdCn"},
		{"trdCnZ"},
		{"ZtrdCnt"},
		{"trdCntZ"},
		{"ZtrdCntA"},
		{"trdCntAZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.TrdCntA()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestTrdCntBValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"trdCntB", 7},
		{"trdCntB ", 7},
		{"trdCntB\n", 7},
		{"trdCntB.", 7},
		{"trdCntB:", 7},
		{"trdCntB,", 7},
		{"trdCntB\"", 7},
		{"trdCntB(", 7},
		{"trdCntB)", 7},
		{"trdCntB[", 7},
		{"trdCntB]", 7},
		{"trdCntB// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.TrdCntB()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestTrdCntBInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_t"},
		{"t_"},
		{"_tr"},
		{"tr_"},
		{"_trd"},
		{"trd_"},
		{"_trdC"},
		{"trdC_"},
		{"_trdCn"},
		{"trdCn_"},
		{"_trdCnt"},
		{"trdCnt_"},
		{"_trdCntB"},
		{"trdCntB_"},
		{"9713t"},
		{"t9713"},
		{"9713tr"},
		{"tr9713"},
		{"9713trd"},
		{"trd9713"},
		{"9713trdC"},
		{"trdC9713"},
		{"9713trdCn"},
		{"trdCn9713"},
		{"9713trdCnt"},
		{"trdCnt9713"},
		{"9713trdCntB"},
		{"trdCntB9713"},
		{"Zt"},
		{"tZ"},
		{"Ztr"},
		{"trZ"},
		{"Ztrd"},
		{"trdZ"},
		{"ZtrdC"},
		{"trdCZ"},
		{"ZtrdCn"},
		{"trdCnZ"},
		{"ZtrdCnt"},
		{"trdCntZ"},
		{"ZtrdCntB"},
		{"trdCntBZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.TrdCntB()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestTrdCntDltValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"trdCntDlt", 9},
		{"trdCntDlt ", 9},
		{"trdCntDlt\n", 9},
		{"trdCntDlt.", 9},
		{"trdCntDlt:", 9},
		{"trdCntDlt,", 9},
		{"trdCntDlt\"", 9},
		{"trdCntDlt(", 9},
		{"trdCntDlt)", 9},
		{"trdCntDlt[", 9},
		{"trdCntDlt]", 9},
		{"trdCntDlt// comment", 9},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.TrdCntDlt()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestTrdCntDltInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_t"},
		{"t_"},
		{"_tr"},
		{"tr_"},
		{"_trd"},
		{"trd_"},
		{"_trdC"},
		{"trdC_"},
		{"_trdCn"},
		{"trdCn_"},
		{"_trdCnt"},
		{"trdCnt_"},
		{"_trdCntD"},
		{"trdCntD_"},
		{"_trdCntDl"},
		{"trdCntDl_"},
		{"_trdCntDlt"},
		{"trdCntDlt_"},
		{"9713t"},
		{"t9713"},
		{"9713tr"},
		{"tr9713"},
		{"9713trd"},
		{"trd9713"},
		{"9713trdC"},
		{"trdC9713"},
		{"9713trdCn"},
		{"trdCn9713"},
		{"9713trdCnt"},
		{"trdCnt9713"},
		{"9713trdCntD"},
		{"trdCntD9713"},
		{"9713trdCntDl"},
		{"trdCntDl9713"},
		{"9713trdCntDlt"},
		{"trdCntDlt9713"},
		{"Zt"},
		{"tZ"},
		{"Ztr"},
		{"trZ"},
		{"Ztrd"},
		{"trdZ"},
		{"ZtrdC"},
		{"trdCZ"},
		{"ZtrdCn"},
		{"trdCnZ"},
		{"ZtrdCnt"},
		{"trdCntZ"},
		{"ZtrdCntD"},
		{"trdCntDZ"},
		{"ZtrdCntDl"},
		{"trdCntDlZ"},
		{"ZtrdCntDlt"},
		{"trdCntDltZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.TrdCntDlt()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestTrdPctAValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"trdPctA", 7},
		{"trdPctA ", 7},
		{"trdPctA\n", 7},
		{"trdPctA.", 7},
		{"trdPctA:", 7},
		{"trdPctA,", 7},
		{"trdPctA\"", 7},
		{"trdPctA(", 7},
		{"trdPctA)", 7},
		{"trdPctA[", 7},
		{"trdPctA]", 7},
		{"trdPctA// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.TrdPctA()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestTrdPctAInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_t"},
		{"t_"},
		{"_tr"},
		{"tr_"},
		{"_trd"},
		{"trd_"},
		{"_trdP"},
		{"trdP_"},
		{"_trdPc"},
		{"trdPc_"},
		{"_trdPct"},
		{"trdPct_"},
		{"_trdPctA"},
		{"trdPctA_"},
		{"9713t"},
		{"t9713"},
		{"9713tr"},
		{"tr9713"},
		{"9713trd"},
		{"trd9713"},
		{"9713trdP"},
		{"trdP9713"},
		{"9713trdPc"},
		{"trdPc9713"},
		{"9713trdPct"},
		{"trdPct9713"},
		{"9713trdPctA"},
		{"trdPctA9713"},
		{"Zt"},
		{"tZ"},
		{"Ztr"},
		{"trZ"},
		{"Ztrd"},
		{"trdZ"},
		{"ZtrdP"},
		{"trdPZ"},
		{"ZtrdPc"},
		{"trdPcZ"},
		{"ZtrdPct"},
		{"trdPctZ"},
		{"ZtrdPctA"},
		{"trdPctAZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.TrdPctA()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestTrdPctBValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"trdPctB", 7},
		{"trdPctB ", 7},
		{"trdPctB\n", 7},
		{"trdPctB.", 7},
		{"trdPctB:", 7},
		{"trdPctB,", 7},
		{"trdPctB\"", 7},
		{"trdPctB(", 7},
		{"trdPctB)", 7},
		{"trdPctB[", 7},
		{"trdPctB]", 7},
		{"trdPctB// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.TrdPctB()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestTrdPctBInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_t"},
		{"t_"},
		{"_tr"},
		{"tr_"},
		{"_trd"},
		{"trd_"},
		{"_trdP"},
		{"trdP_"},
		{"_trdPc"},
		{"trdPc_"},
		{"_trdPct"},
		{"trdPct_"},
		{"_trdPctB"},
		{"trdPctB_"},
		{"9713t"},
		{"t9713"},
		{"9713tr"},
		{"tr9713"},
		{"9713trd"},
		{"trd9713"},
		{"9713trdP"},
		{"trdP9713"},
		{"9713trdPc"},
		{"trdPc9713"},
		{"9713trdPct"},
		{"trdPct9713"},
		{"9713trdPctB"},
		{"trdPctB9713"},
		{"Zt"},
		{"tZ"},
		{"Ztr"},
		{"trZ"},
		{"Ztrd"},
		{"trdZ"},
		{"ZtrdP"},
		{"trdPZ"},
		{"ZtrdPc"},
		{"trdPcZ"},
		{"ZtrdPct"},
		{"trdPctZ"},
		{"ZtrdPctB"},
		{"trdPctBZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.TrdPctB()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestTrdPctDltValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"trdPctDlt", 9},
		{"trdPctDlt ", 9},
		{"trdPctDlt\n", 9},
		{"trdPctDlt.", 9},
		{"trdPctDlt:", 9},
		{"trdPctDlt,", 9},
		{"trdPctDlt\"", 9},
		{"trdPctDlt(", 9},
		{"trdPctDlt)", 9},
		{"trdPctDlt[", 9},
		{"trdPctDlt]", 9},
		{"trdPctDlt// comment", 9},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.TrdPctDlt()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestTrdPctDltInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_t"},
		{"t_"},
		{"_tr"},
		{"tr_"},
		{"_trd"},
		{"trd_"},
		{"_trdP"},
		{"trdP_"},
		{"_trdPc"},
		{"trdPc_"},
		{"_trdPct"},
		{"trdPct_"},
		{"_trdPctD"},
		{"trdPctD_"},
		{"_trdPctDl"},
		{"trdPctDl_"},
		{"_trdPctDlt"},
		{"trdPctDlt_"},
		{"9713t"},
		{"t9713"},
		{"9713tr"},
		{"tr9713"},
		{"9713trd"},
		{"trd9713"},
		{"9713trdP"},
		{"trdP9713"},
		{"9713trdPc"},
		{"trdPc9713"},
		{"9713trdPct"},
		{"trdPct9713"},
		{"9713trdPctD"},
		{"trdPctD9713"},
		{"9713trdPctDl"},
		{"trdPctDl9713"},
		{"9713trdPctDlt"},
		{"trdPctDlt9713"},
		{"Zt"},
		{"tZ"},
		{"Ztr"},
		{"trZ"},
		{"Ztrd"},
		{"trdZ"},
		{"ZtrdP"},
		{"trdPZ"},
		{"ZtrdPc"},
		{"trdPcZ"},
		{"ZtrdPct"},
		{"trdPctZ"},
		{"ZtrdPctD"},
		{"trdPctDZ"},
		{"ZtrdPctDl"},
		{"trdPctDlZ"},
		{"ZtrdPctDlt"},
		{"trdPctDltZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.TrdPctDlt()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestPthBValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"pthB", 4},
		{"pthB ", 4},
		{"pthB\n", 4},
		{"pthB.", 4},
		{"pthB:", 4},
		{"pthB,", 4},
		{"pthB\"", 4},
		{"pthB(", 4},
		{"pthB)", 4},
		{"pthB[", 4},
		{"pthB]", 4},
		{"pthB// comment", 4},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.PthB()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestPthBInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_p"},
		{"p_"},
		{"_pt"},
		{"pt_"},
		{"_pth"},
		{"pth_"},
		{"_pthB"},
		{"pthB_"},
		{"9713p"},
		{"p9713"},
		{"9713pt"},
		{"pt9713"},
		{"9713pth"},
		{"pth9713"},
		{"9713pthB"},
		{"pthB9713"},
		{"Zp"},
		{"pZ"},
		{"Zpt"},
		{"ptZ"},
		{"Zpth"},
		{"pthZ"},
		{"ZpthB"},
		{"pthBZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.PthB()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestClrValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"clr", 3},
		{"clr ", 3},
		{"clr\n", 3},
		{"clr.", 3},
		{"clr:", 3},
		{"clr,", 3},
		{"clr\"", 3},
		{"clr(", 3},
		{"clr)", 3},
		{"clr[", 3},
		{"clr]", 3},
		{"clr// comment", 3},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Clr()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestClrInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_c"},
		{"c_"},
		{"_cl"},
		{"cl_"},
		{"_clr"},
		{"clr_"},
		{"9713c"},
		{"c9713"},
		{"9713cl"},
		{"cl9713"},
		{"9713clr"},
		{"clr9713"},
		{"Zc"},
		{"cZ"},
		{"Zcl"},
		{"clZ"},
		{"Zclr"},
		{"clrZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Clr()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestWidValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"wid", 3},
		{"wid ", 3},
		{"wid\n", 3},
		{"wid.", 3},
		{"wid:", 3},
		{"wid,", 3},
		{"wid\"", 3},
		{"wid(", 3},
		{"wid)", 3},
		{"wid[", 3},
		{"wid]", 3},
		{"wid// comment", 3},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Wid()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestWidInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_w"},
		{"w_"},
		{"_wi"},
		{"wi_"},
		{"_wid"},
		{"wid_"},
		{"9713w"},
		{"w9713"},
		{"9713wi"},
		{"wi9713"},
		{"9713wid"},
		{"wid9713"},
		{"Zw"},
		{"wZ"},
		{"Zwi"},
		{"wiZ"},
		{"Zwid"},
		{"widZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Wid()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestMinValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"min", 3},
		{"min ", 3},
		{"min\n", 3},
		{"min.", 3},
		{"min:", 3},
		{"min,", 3},
		{"min\"", 3},
		{"min(", 3},
		{"min)", 3},
		{"min[", 3},
		{"min]", 3},
		{"min// comment", 3},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Min()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestMinInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_m"},
		{"m_"},
		{"_mi"},
		{"mi_"},
		{"_min"},
		{"min_"},
		{"9713m"},
		{"m9713"},
		{"9713mi"},
		{"mi9713"},
		{"9713min"},
		{"min9713"},
		{"Zm"},
		{"mZ"},
		{"Zmi"},
		{"miZ"},
		{"Zmin"},
		{"minZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Min()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestMaxValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"max", 3},
		{"max ", 3},
		{"max\n", 3},
		{"max.", 3},
		{"max:", 3},
		{"max,", 3},
		{"max\"", 3},
		{"max(", 3},
		{"max)", 3},
		{"max[", 3},
		{"max]", 3},
		{"max// comment", 3},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Max()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestMaxInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_m"},
		{"m_"},
		{"_ma"},
		{"ma_"},
		{"_max"},
		{"max_"},
		{"9713m"},
		{"m9713"},
		{"9713ma"},
		{"ma9713"},
		{"9713max"},
		{"max9713"},
		{"Zm"},
		{"mZ"},
		{"Zma"},
		{"maZ"},
		{"Zmax"},
		{"maxZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Max()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestEqiDstValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"eqiDst", 6},
		{"eqiDst ", 6},
		{"eqiDst\n", 6},
		{"eqiDst.", 6},
		{"eqiDst:", 6},
		{"eqiDst,", 6},
		{"eqiDst\"", 6},
		{"eqiDst(", 6},
		{"eqiDst)", 6},
		{"eqiDst[", 6},
		{"eqiDst]", 6},
		{"eqiDst// comment", 6},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.EqiDst()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestEqiDstInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_e"},
		{"e_"},
		{"_eq"},
		{"eq_"},
		{"_eqi"},
		{"eqi_"},
		{"_eqiD"},
		{"eqiD_"},
		{"_eqiDs"},
		{"eqiDs_"},
		{"_eqiDst"},
		{"eqiDst_"},
		{"9713e"},
		{"e9713"},
		{"9713eq"},
		{"eq9713"},
		{"9713eqi"},
		{"eqi9713"},
		{"9713eqiD"},
		{"eqiD9713"},
		{"9713eqiDs"},
		{"eqiDs9713"},
		{"9713eqiDst"},
		{"eqiDst9713"},
		{"Ze"},
		{"eZ"},
		{"Zeq"},
		{"eqZ"},
		{"Zeqi"},
		{"eqiZ"},
		{"ZeqiD"},
		{"eqiDZ"},
		{"ZeqiDs"},
		{"eqiDsZ"},
		{"ZeqiDst"},
		{"eqiDstZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.EqiDst()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestTitleValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"title", 5},
		{"title ", 5},
		{"title\n", 5},
		{"title.", 5},
		{"title:", 5},
		{"title,", 5},
		{"title\"", 5},
		{"title(", 5},
		{"title)", 5},
		{"title[", 5},
		{"title]", 5},
		{"title// comment", 5},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Title()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestTitleInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_t"},
		{"t_"},
		{"_ti"},
		{"ti_"},
		{"_tit"},
		{"tit_"},
		{"_titl"},
		{"titl_"},
		{"_title"},
		{"title_"},
		{"9713t"},
		{"t9713"},
		{"9713ti"},
		{"ti9713"},
		{"9713tit"},
		{"tit9713"},
		{"9713titl"},
		{"titl9713"},
		{"9713title"},
		{"title9713"},
		{"Zt"},
		{"tZ"},
		{"Zti"},
		{"tiZ"},
		{"Ztit"},
		{"titZ"},
		{"Ztitl"},
		{"titlZ"},
		{"Ztitle"},
		{"titleZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Title()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestYValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"y", 1},
		{"y ", 1},
		{"y\n", 1},
		{"y.", 1},
		{"y:", 1},
		{"y,", 1},
		{"y\"", 1},
		{"y(", 1},
		{"y)", 1},
		{"y[", 1},
		{"y]", 1},
		{"y// comment", 1},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Y()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestYInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_y"},
		{"y_"},
		{"9713y"},
		{"y9713"},
		{"Zy"},
		{"yZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Y()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestOutlierValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"outlier", 7},
		{"outlier ", 7},
		{"outlier\n", 7},
		{"outlier.", 7},
		{"outlier:", 7},
		{"outlier,", 7},
		{"outlier\"", 7},
		{"outlier(", 7},
		{"outlier)", 7},
		{"outlier[", 7},
		{"outlier]", 7},
		{"outlier// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Outlier()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestOutlierInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_o"},
		{"o_"},
		{"_ou"},
		{"ou_"},
		{"_out"},
		{"out_"},
		{"_outl"},
		{"outl_"},
		{"_outli"},
		{"outli_"},
		{"_outlie"},
		{"outlie_"},
		{"_outlier"},
		{"outlier_"},
		{"9713o"},
		{"o9713"},
		{"9713ou"},
		{"ou9713"},
		{"9713out"},
		{"out9713"},
		{"9713outl"},
		{"outl9713"},
		{"9713outli"},
		{"outli9713"},
		{"9713outlie"},
		{"outlie9713"},
		{"9713outlier"},
		{"outlier9713"},
		{"Zo"},
		{"oZ"},
		{"Zou"},
		{"ouZ"},
		{"Zout"},
		{"outZ"},
		{"Zoutl"},
		{"outlZ"},
		{"Zoutli"},
		{"outliZ"},
		{"Zoutlie"},
		{"outlieZ"},
		{"Zoutlier"},
		{"outlierZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Outlier()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestPltsValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"plts", 4},
		{"plts ", 4},
		{"plts\n", 4},
		{"plts.", 4},
		{"plts:", 4},
		{"plts,", 4},
		{"plts\"", 4},
		{"plts(", 4},
		{"plts)", 4},
		{"plts[", 4},
		{"plts]", 4},
		{"plts// comment", 4},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Plts()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestPltsInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_p"},
		{"p_"},
		{"_pl"},
		{"pl_"},
		{"_plt"},
		{"plt_"},
		{"_plts"},
		{"plts_"},
		{"9713p"},
		{"p9713"},
		{"9713pl"},
		{"pl9713"},
		{"9713plt"},
		{"plt9713"},
		{"9713plts"},
		{"plts9713"},
		{"Zp"},
		{"pZ"},
		{"Zpl"},
		{"plZ"},
		{"Zplt"},
		{"pltZ"},
		{"Zplts"},
		{"pltsZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Plts()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestZeroValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"zero", 4},
		{"zero ", 4},
		{"zero\n", 4},
		{"zero.", 4},
		{"zero:", 4},
		{"zero,", 4},
		{"zero\"", 4},
		{"zero(", 4},
		{"zero)", 4},
		{"zero[", 4},
		{"zero]", 4},
		{"zero// comment", 4},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Zero()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestZeroInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_z"},
		{"z_"},
		{"_ze"},
		{"ze_"},
		{"_zer"},
		{"zer_"},
		{"_zero"},
		{"zero_"},
		{"9713z"},
		{"z9713"},
		{"9713ze"},
		{"ze9713"},
		{"9713zer"},
		{"zer9713"},
		{"9713zero"},
		{"zero9713"},
		{"Zz"},
		{"zZ"},
		{"Zze"},
		{"zeZ"},
		{"Zzer"},
		{"zerZ"},
		{"Zzero"},
		{"zeroZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Zero()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestEmptyValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"empty", 5},
		{"empty ", 5},
		{"empty\n", 5},
		{"empty.", 5},
		{"empty:", 5},
		{"empty,", 5},
		{"empty\"", 5},
		{"empty(", 5},
		{"empty)", 5},
		{"empty[", 5},
		{"empty]", 5},
		{"empty// comment", 5},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Empty()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestEmptyInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_e"},
		{"e_"},
		{"_em"},
		{"em_"},
		{"_emp"},
		{"emp_"},
		{"_empt"},
		{"empt_"},
		{"_empty"},
		{"empty_"},
		{"9713e"},
		{"e9713"},
		{"9713em"},
		{"em9713"},
		{"9713emp"},
		{"emp9713"},
		{"9713empt"},
		{"empt9713"},
		{"9713empty"},
		{"empty9713"},
		{"Ze"},
		{"eZ"},
		{"Zem"},
		{"emZ"},
		{"Zemp"},
		{"empZ"},
		{"Zempt"},
		{"emptZ"},
		{"Zempty"},
		{"emptyZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Empty()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestFlsValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"fls", 3},
		{"fls ", 3},
		{"fls\n", 3},
		{"fls.", 3},
		{"fls:", 3},
		{"fls,", 3},
		{"fls\"", 3},
		{"fls(", 3},
		{"fls)", 3},
		{"fls[", 3},
		{"fls]", 3},
		{"fls// comment", 3},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Fls()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestFlsInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_f"},
		{"f_"},
		{"_fl"},
		{"fl_"},
		{"_fls"},
		{"fls_"},
		{"9713f"},
		{"f9713"},
		{"9713fl"},
		{"fl9713"},
		{"9713fls"},
		{"fls9713"},
		{"Zf"},
		{"fZ"},
		{"Zfl"},
		{"flZ"},
		{"Zfls"},
		{"flsZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Fls()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestTruValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"tru", 3},
		{"tru ", 3},
		{"tru\n", 3},
		{"tru.", 3},
		{"tru:", 3},
		{"tru,", 3},
		{"tru\"", 3},
		{"tru(", 3},
		{"tru)", 3},
		{"tru[", 3},
		{"tru]", 3},
		{"tru// comment", 3},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Tru()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestTruInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_t"},
		{"t_"},
		{"_tr"},
		{"tr_"},
		{"_tru"},
		{"tru_"},
		{"9713t"},
		{"t9713"},
		{"9713tr"},
		{"tr9713"},
		{"9713tru"},
		{"tru9713"},
		{"Zt"},
		{"tZ"},
		{"Ztr"},
		{"trZ"},
		{"Ztru"},
		{"truZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Tru()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestOneValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"one", 3},
		{"one ", 3},
		{"one\n", 3},
		{"one.", 3},
		{"one:", 3},
		{"one,", 3},
		{"one\"", 3},
		{"one(", 3},
		{"one)", 3},
		{"one[", 3},
		{"one]", 3},
		{"one// comment", 3},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.One()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestOneInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_o"},
		{"o_"},
		{"_on"},
		{"on_"},
		{"_one"},
		{"one_"},
		{"9713o"},
		{"o9713"},
		{"9713on"},
		{"on9713"},
		{"9713one"},
		{"one9713"},
		{"Zo"},
		{"oZ"},
		{"Zon"},
		{"onZ"},
		{"Zone"},
		{"oneZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.One()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestNegOneValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"negOne", 6},
		{"negOne ", 6},
		{"negOne\n", 6},
		{"negOne.", 6},
		{"negOne:", 6},
		{"negOne,", 6},
		{"negOne\"", 6},
		{"negOne(", 6},
		{"negOne)", 6},
		{"negOne[", 6},
		{"negOne]", 6},
		{"negOne// comment", 6},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.NegOne()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestNegOneInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_n"},
		{"n_"},
		{"_ne"},
		{"ne_"},
		{"_neg"},
		{"neg_"},
		{"_negO"},
		{"negO_"},
		{"_negOn"},
		{"negOn_"},
		{"_negOne"},
		{"negOne_"},
		{"9713n"},
		{"n9713"},
		{"9713ne"},
		{"ne9713"},
		{"9713neg"},
		{"neg9713"},
		{"9713negO"},
		{"negO9713"},
		{"9713negOn"},
		{"negOn9713"},
		{"9713negOne"},
		{"negOne9713"},
		{"Zn"},
		{"nZ"},
		{"Zne"},
		{"neZ"},
		{"Zneg"},
		{"negZ"},
		{"ZnegO"},
		{"negOZ"},
		{"ZnegOn"},
		{"negOnZ"},
		{"ZnegOne"},
		{"negOneZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.NegOne()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestHndrdValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"hndrd", 5},
		{"hndrd ", 5},
		{"hndrd\n", 5},
		{"hndrd.", 5},
		{"hndrd:", 5},
		{"hndrd,", 5},
		{"hndrd\"", 5},
		{"hndrd(", 5},
		{"hndrd)", 5},
		{"hndrd[", 5},
		{"hndrd]", 5},
		{"hndrd// comment", 5},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Hndrd()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestHndrdInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_h"},
		{"h_"},
		{"_hn"},
		{"hn_"},
		{"_hnd"},
		{"hnd_"},
		{"_hndr"},
		{"hndr_"},
		{"_hndrd"},
		{"hndrd_"},
		{"9713h"},
		{"h9713"},
		{"9713hn"},
		{"hn9713"},
		{"9713hnd"},
		{"hnd9713"},
		{"9713hndr"},
		{"hndr9713"},
		{"9713hndrd"},
		{"hndrd9713"},
		{"Zh"},
		{"hZ"},
		{"Zhn"},
		{"hnZ"},
		{"Zhnd"},
		{"hndZ"},
		{"Zhndr"},
		{"hndrZ"},
		{"Zhndrd"},
		{"hndrdZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Hndrd()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestTinyValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"tiny", 4},
		{"tiny ", 4},
		{"tiny\n", 4},
		{"tiny.", 4},
		{"tiny:", 4},
		{"tiny,", 4},
		{"tiny\"", 4},
		{"tiny(", 4},
		{"tiny)", 4},
		{"tiny[", 4},
		{"tiny]", 4},
		{"tiny// comment", 4},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Tiny()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestTinyInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_t"},
		{"t_"},
		{"_ti"},
		{"ti_"},
		{"_tin"},
		{"tin_"},
		{"_tiny"},
		{"tiny_"},
		{"9713t"},
		{"t9713"},
		{"9713ti"},
		{"ti9713"},
		{"9713tin"},
		{"tin9713"},
		{"9713tiny"},
		{"tiny9713"},
		{"Zt"},
		{"tZ"},
		{"Zti"},
		{"tiZ"},
		{"Ztin"},
		{"tinZ"},
		{"Ztiny"},
		{"tinyZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Tiny()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestSecondValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"second", 6},
		{"second ", 6},
		{"second\n", 6},
		{"second.", 6},
		{"second:", 6},
		{"second,", 6},
		{"second\"", 6},
		{"second(", 6},
		{"second)", 6},
		{"second[", 6},
		{"second]", 6},
		{"second// comment", 6},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Second()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestSecondInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_s"},
		{"s_"},
		{"_se"},
		{"se_"},
		{"_sec"},
		{"sec_"},
		{"_seco"},
		{"seco_"},
		{"_secon"},
		{"secon_"},
		{"_second"},
		{"second_"},
		{"9713s"},
		{"s9713"},
		{"9713se"},
		{"se9713"},
		{"9713sec"},
		{"sec9713"},
		{"9713seco"},
		{"seco9713"},
		{"9713secon"},
		{"secon9713"},
		{"9713second"},
		{"second9713"},
		{"Zs"},
		{"sZ"},
		{"Zse"},
		{"seZ"},
		{"Zsec"},
		{"secZ"},
		{"Zseco"},
		{"secoZ"},
		{"Zsecon"},
		{"seconZ"},
		{"Zsecond"},
		{"secondZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Second()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestMinuteValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"minute", 6},
		{"minute ", 6},
		{"minute\n", 6},
		{"minute.", 6},
		{"minute:", 6},
		{"minute,", 6},
		{"minute\"", 6},
		{"minute(", 6},
		{"minute)", 6},
		{"minute[", 6},
		{"minute]", 6},
		{"minute// comment", 6},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Minute()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestMinuteInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_m"},
		{"m_"},
		{"_mi"},
		{"mi_"},
		{"_min"},
		{"min_"},
		{"_minu"},
		{"minu_"},
		{"_minut"},
		{"minut_"},
		{"_minute"},
		{"minute_"},
		{"9713m"},
		{"m9713"},
		{"9713mi"},
		{"mi9713"},
		{"9713min"},
		{"min9713"},
		{"9713minu"},
		{"minu9713"},
		{"9713minut"},
		{"minut9713"},
		{"9713minute"},
		{"minute9713"},
		{"Zm"},
		{"mZ"},
		{"Zmi"},
		{"miZ"},
		{"Zmin"},
		{"minZ"},
		{"Zminu"},
		{"minuZ"},
		{"Zminut"},
		{"minutZ"},
		{"Zminute"},
		{"minuteZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Minute()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestHourValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"hour", 4},
		{"hour ", 4},
		{"hour\n", 4},
		{"hour.", 4},
		{"hour:", 4},
		{"hour,", 4},
		{"hour\"", 4},
		{"hour(", 4},
		{"hour)", 4},
		{"hour[", 4},
		{"hour]", 4},
		{"hour// comment", 4},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Hour()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestHourInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_h"},
		{"h_"},
		{"_ho"},
		{"ho_"},
		{"_hou"},
		{"hou_"},
		{"_hour"},
		{"hour_"},
		{"9713h"},
		{"h9713"},
		{"9713ho"},
		{"ho9713"},
		{"9713hou"},
		{"hou9713"},
		{"9713hour"},
		{"hour9713"},
		{"Zh"},
		{"hZ"},
		{"Zho"},
		{"hoZ"},
		{"Zhou"},
		{"houZ"},
		{"Zhour"},
		{"hourZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Hour()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestDayValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"day", 3},
		{"day ", 3},
		{"day\n", 3},
		{"day.", 3},
		{"day:", 3},
		{"day,", 3},
		{"day\"", 3},
		{"day(", 3},
		{"day)", 3},
		{"day[", 3},
		{"day]", 3},
		{"day// comment", 3},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Day()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestDayInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_d"},
		{"d_"},
		{"_da"},
		{"da_"},
		{"_day"},
		{"day_"},
		{"9713d"},
		{"d9713"},
		{"9713da"},
		{"da9713"},
		{"9713day"},
		{"day9713"},
		{"Zd"},
		{"dZ"},
		{"Zda"},
		{"daZ"},
		{"Zday"},
		{"dayZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Day()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestWeekValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"week", 4},
		{"week ", 4},
		{"week\n", 4},
		{"week.", 4},
		{"week:", 4},
		{"week,", 4},
		{"week\"", 4},
		{"week(", 4},
		{"week)", 4},
		{"week[", 4},
		{"week]", 4},
		{"week// comment", 4},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Week()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestWeekInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_w"},
		{"w_"},
		{"_we"},
		{"we_"},
		{"_wee"},
		{"wee_"},
		{"_week"},
		{"week_"},
		{"9713w"},
		{"w9713"},
		{"9713we"},
		{"we9713"},
		{"9713wee"},
		{"wee9713"},
		{"9713week"},
		{"week9713"},
		{"Zw"},
		{"wZ"},
		{"Zwe"},
		{"weZ"},
		{"Zwee"},
		{"weeZ"},
		{"Zweek"},
		{"weekZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Week()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestS1Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"s1", 2},
		{"s1 ", 2},
		{"s1\n", 2},
		{"s1.", 2},
		{"s1:", 2},
		{"s1,", 2},
		{"s1\"", 2},
		{"s1(", 2},
		{"s1)", 2},
		{"s1[", 2},
		{"s1]", 2},
		{"s1// comment", 2},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.S1()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestS1Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_s"},
		{"s_"},
		{"_s1"},
		{"s1_"},
		{"9713s"},
		{"s9713"},
		{"9713s1"},
		{"s19713"},
		{"Zs"},
		{"sZ"},
		{"Zs1"},
		{"s1Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.S1()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestS5Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"s5", 2},
		{"s5 ", 2},
		{"s5\n", 2},
		{"s5.", 2},
		{"s5:", 2},
		{"s5,", 2},
		{"s5\"", 2},
		{"s5(", 2},
		{"s5)", 2},
		{"s5[", 2},
		{"s5]", 2},
		{"s5// comment", 2},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.S5()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestS5Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_s"},
		{"s_"},
		{"_s5"},
		{"s5_"},
		{"9713s"},
		{"s9713"},
		{"9713s5"},
		{"s59713"},
		{"Zs"},
		{"sZ"},
		{"Zs5"},
		{"s5Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.S5()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestS10Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"s10", 3},
		{"s10 ", 3},
		{"s10\n", 3},
		{"s10.", 3},
		{"s10:", 3},
		{"s10,", 3},
		{"s10\"", 3},
		{"s10(", 3},
		{"s10)", 3},
		{"s10[", 3},
		{"s10]", 3},
		{"s10// comment", 3},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.S10()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestS10Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_s"},
		{"s_"},
		{"_s1"},
		{"s1_"},
		{"_s10"},
		{"s10_"},
		{"9713s"},
		{"s9713"},
		{"9713s1"},
		{"s19713"},
		{"9713s10"},
		{"s109713"},
		{"Zs"},
		{"sZ"},
		{"Zs1"},
		{"s1Z"},
		{"Zs10"},
		{"s10Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.S10()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestS15Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"s15", 3},
		{"s15 ", 3},
		{"s15\n", 3},
		{"s15.", 3},
		{"s15:", 3},
		{"s15,", 3},
		{"s15\"", 3},
		{"s15(", 3},
		{"s15)", 3},
		{"s15[", 3},
		{"s15]", 3},
		{"s15// comment", 3},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.S15()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestS15Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_s"},
		{"s_"},
		{"_s1"},
		{"s1_"},
		{"_s15"},
		{"s15_"},
		{"9713s"},
		{"s9713"},
		{"9713s1"},
		{"s19713"},
		{"9713s15"},
		{"s159713"},
		{"Zs"},
		{"sZ"},
		{"Zs1"},
		{"s1Z"},
		{"Zs15"},
		{"s15Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.S15()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestS20Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"s20", 3},
		{"s20 ", 3},
		{"s20\n", 3},
		{"s20.", 3},
		{"s20:", 3},
		{"s20,", 3},
		{"s20\"", 3},
		{"s20(", 3},
		{"s20)", 3},
		{"s20[", 3},
		{"s20]", 3},
		{"s20// comment", 3},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.S20()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestS20Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_s"},
		{"s_"},
		{"_s2"},
		{"s2_"},
		{"_s20"},
		{"s20_"},
		{"9713s"},
		{"s9713"},
		{"9713s2"},
		{"s29713"},
		{"9713s20"},
		{"s209713"},
		{"Zs"},
		{"sZ"},
		{"Zs2"},
		{"s2Z"},
		{"Zs20"},
		{"s20Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.S20()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestS30Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"s30", 3},
		{"s30 ", 3},
		{"s30\n", 3},
		{"s30.", 3},
		{"s30:", 3},
		{"s30,", 3},
		{"s30\"", 3},
		{"s30(", 3},
		{"s30)", 3},
		{"s30[", 3},
		{"s30]", 3},
		{"s30// comment", 3},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.S30()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestS30Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_s"},
		{"s_"},
		{"_s3"},
		{"s3_"},
		{"_s30"},
		{"s30_"},
		{"9713s"},
		{"s9713"},
		{"9713s3"},
		{"s39713"},
		{"9713s30"},
		{"s309713"},
		{"Zs"},
		{"sZ"},
		{"Zs3"},
		{"s3Z"},
		{"Zs30"},
		{"s30Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.S30()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestS40Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"s40", 3},
		{"s40 ", 3},
		{"s40\n", 3},
		{"s40.", 3},
		{"s40:", 3},
		{"s40,", 3},
		{"s40\"", 3},
		{"s40(", 3},
		{"s40)", 3},
		{"s40[", 3},
		{"s40]", 3},
		{"s40// comment", 3},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.S40()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestS40Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_s"},
		{"s_"},
		{"_s4"},
		{"s4_"},
		{"_s40"},
		{"s40_"},
		{"9713s"},
		{"s9713"},
		{"9713s4"},
		{"s49713"},
		{"9713s40"},
		{"s409713"},
		{"Zs"},
		{"sZ"},
		{"Zs4"},
		{"s4Z"},
		{"Zs40"},
		{"s40Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.S40()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestS50Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"s50", 3},
		{"s50 ", 3},
		{"s50\n", 3},
		{"s50.", 3},
		{"s50:", 3},
		{"s50,", 3},
		{"s50\"", 3},
		{"s50(", 3},
		{"s50)", 3},
		{"s50[", 3},
		{"s50]", 3},
		{"s50// comment", 3},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.S50()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestS50Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_s"},
		{"s_"},
		{"_s5"},
		{"s5_"},
		{"_s50"},
		{"s50_"},
		{"9713s"},
		{"s9713"},
		{"9713s5"},
		{"s59713"},
		{"9713s50"},
		{"s509713"},
		{"Zs"},
		{"sZ"},
		{"Zs5"},
		{"s5Z"},
		{"Zs50"},
		{"s50Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.S50()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestM1Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"m1", 2},
		{"m1 ", 2},
		{"m1\n", 2},
		{"m1.", 2},
		{"m1:", 2},
		{"m1,", 2},
		{"m1\"", 2},
		{"m1(", 2},
		{"m1)", 2},
		{"m1[", 2},
		{"m1]", 2},
		{"m1// comment", 2},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.M1()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestM1Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_m"},
		{"m_"},
		{"_m1"},
		{"m1_"},
		{"9713m"},
		{"m9713"},
		{"9713m1"},
		{"m19713"},
		{"Zm"},
		{"mZ"},
		{"Zm1"},
		{"m1Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.M1()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestM5Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"m5", 2},
		{"m5 ", 2},
		{"m5\n", 2},
		{"m5.", 2},
		{"m5:", 2},
		{"m5,", 2},
		{"m5\"", 2},
		{"m5(", 2},
		{"m5)", 2},
		{"m5[", 2},
		{"m5]", 2},
		{"m5// comment", 2},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.M5()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestM5Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_m"},
		{"m_"},
		{"_m5"},
		{"m5_"},
		{"9713m"},
		{"m9713"},
		{"9713m5"},
		{"m59713"},
		{"Zm"},
		{"mZ"},
		{"Zm5"},
		{"m5Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.M5()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestM10Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"m10", 3},
		{"m10 ", 3},
		{"m10\n", 3},
		{"m10.", 3},
		{"m10:", 3},
		{"m10,", 3},
		{"m10\"", 3},
		{"m10(", 3},
		{"m10)", 3},
		{"m10[", 3},
		{"m10]", 3},
		{"m10// comment", 3},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.M10()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestM10Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_m"},
		{"m_"},
		{"_m1"},
		{"m1_"},
		{"_m10"},
		{"m10_"},
		{"9713m"},
		{"m9713"},
		{"9713m1"},
		{"m19713"},
		{"9713m10"},
		{"m109713"},
		{"Zm"},
		{"mZ"},
		{"Zm1"},
		{"m1Z"},
		{"Zm10"},
		{"m10Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.M10()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestM15Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"m15", 3},
		{"m15 ", 3},
		{"m15\n", 3},
		{"m15.", 3},
		{"m15:", 3},
		{"m15,", 3},
		{"m15\"", 3},
		{"m15(", 3},
		{"m15)", 3},
		{"m15[", 3},
		{"m15]", 3},
		{"m15// comment", 3},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.M15()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestM15Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_m"},
		{"m_"},
		{"_m1"},
		{"m1_"},
		{"_m15"},
		{"m15_"},
		{"9713m"},
		{"m9713"},
		{"9713m1"},
		{"m19713"},
		{"9713m15"},
		{"m159713"},
		{"Zm"},
		{"mZ"},
		{"Zm1"},
		{"m1Z"},
		{"Zm15"},
		{"m15Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.M15()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestM20Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"m20", 3},
		{"m20 ", 3},
		{"m20\n", 3},
		{"m20.", 3},
		{"m20:", 3},
		{"m20,", 3},
		{"m20\"", 3},
		{"m20(", 3},
		{"m20)", 3},
		{"m20[", 3},
		{"m20]", 3},
		{"m20// comment", 3},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.M20()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestM20Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_m"},
		{"m_"},
		{"_m2"},
		{"m2_"},
		{"_m20"},
		{"m20_"},
		{"9713m"},
		{"m9713"},
		{"9713m2"},
		{"m29713"},
		{"9713m20"},
		{"m209713"},
		{"Zm"},
		{"mZ"},
		{"Zm2"},
		{"m2Z"},
		{"Zm20"},
		{"m20Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.M20()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestM30Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"m30", 3},
		{"m30 ", 3},
		{"m30\n", 3},
		{"m30.", 3},
		{"m30:", 3},
		{"m30,", 3},
		{"m30\"", 3},
		{"m30(", 3},
		{"m30)", 3},
		{"m30[", 3},
		{"m30]", 3},
		{"m30// comment", 3},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.M30()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestM30Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_m"},
		{"m_"},
		{"_m3"},
		{"m3_"},
		{"_m30"},
		{"m30_"},
		{"9713m"},
		{"m9713"},
		{"9713m3"},
		{"m39713"},
		{"9713m30"},
		{"m309713"},
		{"Zm"},
		{"mZ"},
		{"Zm3"},
		{"m3Z"},
		{"Zm30"},
		{"m30Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.M30()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestM40Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"m40", 3},
		{"m40 ", 3},
		{"m40\n", 3},
		{"m40.", 3},
		{"m40:", 3},
		{"m40,", 3},
		{"m40\"", 3},
		{"m40(", 3},
		{"m40)", 3},
		{"m40[", 3},
		{"m40]", 3},
		{"m40// comment", 3},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.M40()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestM40Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_m"},
		{"m_"},
		{"_m4"},
		{"m4_"},
		{"_m40"},
		{"m40_"},
		{"9713m"},
		{"m9713"},
		{"9713m4"},
		{"m49713"},
		{"9713m40"},
		{"m409713"},
		{"Zm"},
		{"mZ"},
		{"Zm4"},
		{"m4Z"},
		{"Zm40"},
		{"m40Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.M40()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestM50Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"m50", 3},
		{"m50 ", 3},
		{"m50\n", 3},
		{"m50.", 3},
		{"m50:", 3},
		{"m50,", 3},
		{"m50\"", 3},
		{"m50(", 3},
		{"m50)", 3},
		{"m50[", 3},
		{"m50]", 3},
		{"m50// comment", 3},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.M50()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestM50Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_m"},
		{"m_"},
		{"_m5"},
		{"m5_"},
		{"_m50"},
		{"m50_"},
		{"9713m"},
		{"m9713"},
		{"9713m5"},
		{"m59713"},
		{"9713m50"},
		{"m509713"},
		{"Zm"},
		{"mZ"},
		{"Zm5"},
		{"m5Z"},
		{"Zm50"},
		{"m50Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.M50()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestH1Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"h1", 2},
		{"h1 ", 2},
		{"h1\n", 2},
		{"h1.", 2},
		{"h1:", 2},
		{"h1,", 2},
		{"h1\"", 2},
		{"h1(", 2},
		{"h1)", 2},
		{"h1[", 2},
		{"h1]", 2},
		{"h1// comment", 2},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.H1()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestH1Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_h"},
		{"h_"},
		{"_h1"},
		{"h1_"},
		{"9713h"},
		{"h9713"},
		{"9713h1"},
		{"h19713"},
		{"Zh"},
		{"hZ"},
		{"Zh1"},
		{"h1Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.H1()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestD1Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"d1", 2},
		{"d1 ", 2},
		{"d1\n", 2},
		{"d1.", 2},
		{"d1:", 2},
		{"d1,", 2},
		{"d1\"", 2},
		{"d1(", 2},
		{"d1)", 2},
		{"d1[", 2},
		{"d1]", 2},
		{"d1// comment", 2},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.D1()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestD1Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_d"},
		{"d_"},
		{"_d1"},
		{"d1_"},
		{"9713d"},
		{"d9713"},
		{"9713d1"},
		{"d19713"},
		{"Zd"},
		{"dZ"},
		{"Zd1"},
		{"d1Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.D1()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestResolutionValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"resolution", 10},
		{"resolution ", 10},
		{"resolution\n", 10},
		{"resolution.", 10},
		{"resolution:", 10},
		{"resolution,", 10},
		{"resolution\"", 10},
		{"resolution(", 10},
		{"resolution)", 10},
		{"resolution[", 10},
		{"resolution]", 10},
		{"resolution// comment", 10},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Resolution()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestResolutionInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_r"},
		{"r_"},
		{"_re"},
		{"re_"},
		{"_res"},
		{"res_"},
		{"_reso"},
		{"reso_"},
		{"_resol"},
		{"resol_"},
		{"_resolu"},
		{"resolu_"},
		{"_resolut"},
		{"resolut_"},
		{"_resoluti"},
		{"resoluti_"},
		{"_resolutio"},
		{"resolutio_"},
		{"_resolution"},
		{"resolution_"},
		{"9713r"},
		{"r9713"},
		{"9713re"},
		{"re9713"},
		{"9713res"},
		{"res9713"},
		{"9713reso"},
		{"reso9713"},
		{"9713resol"},
		{"resol9713"},
		{"9713resolu"},
		{"resolu9713"},
		{"9713resolut"},
		{"resolut9713"},
		{"9713resoluti"},
		{"resoluti9713"},
		{"9713resolutio"},
		{"resolutio9713"},
		{"9713resolution"},
		{"resolution9713"},
		{"Zr"},
		{"rZ"},
		{"Zre"},
		{"reZ"},
		{"Zres"},
		{"resZ"},
		{"Zreso"},
		{"resoZ"},
		{"Zresol"},
		{"resolZ"},
		{"Zresolu"},
		{"resoluZ"},
		{"Zresolut"},
		{"resolutZ"},
		{"Zresoluti"},
		{"resolutiZ"},
		{"Zresolutio"},
		{"resolutioZ"},
		{"Zresolution"},
		{"resolutionZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Resolution()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestBlackValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"black", 5},
		{"black ", 5},
		{"black\n", 5},
		{"black.", 5},
		{"black:", 5},
		{"black,", 5},
		{"black\"", 5},
		{"black(", 5},
		{"black)", 5},
		{"black[", 5},
		{"black]", 5},
		{"black// comment", 5},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Black()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestBlackInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_b"},
		{"b_"},
		{"_bl"},
		{"bl_"},
		{"_bla"},
		{"bla_"},
		{"_blac"},
		{"blac_"},
		{"_black"},
		{"black_"},
		{"9713b"},
		{"b9713"},
		{"9713bl"},
		{"bl9713"},
		{"9713bla"},
		{"bla9713"},
		{"9713blac"},
		{"blac9713"},
		{"9713black"},
		{"black9713"},
		{"Zb"},
		{"bZ"},
		{"Zbl"},
		{"blZ"},
		{"Zbla"},
		{"blaZ"},
		{"Zblac"},
		{"blacZ"},
		{"Zblack"},
		{"blackZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Black()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestWhiteValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"white", 5},
		{"white ", 5},
		{"white\n", 5},
		{"white.", 5},
		{"white:", 5},
		{"white,", 5},
		{"white\"", 5},
		{"white(", 5},
		{"white)", 5},
		{"white[", 5},
		{"white]", 5},
		{"white// comment", 5},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.White()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestWhiteInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_w"},
		{"w_"},
		{"_wh"},
		{"wh_"},
		{"_whi"},
		{"whi_"},
		{"_whit"},
		{"whit_"},
		{"_white"},
		{"white_"},
		{"9713w"},
		{"w9713"},
		{"9713wh"},
		{"wh9713"},
		{"9713whi"},
		{"whi9713"},
		{"9713whit"},
		{"whit9713"},
		{"9713white"},
		{"white9713"},
		{"Zw"},
		{"wZ"},
		{"Zwh"},
		{"whZ"},
		{"Zwhi"},
		{"whiZ"},
		{"Zwhit"},
		{"whitZ"},
		{"Zwhite"},
		{"whiteZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.White()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestRed50Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"red50", 5},
		{"red50 ", 5},
		{"red50\n", 5},
		{"red50.", 5},
		{"red50:", 5},
		{"red50,", 5},
		{"red50\"", 5},
		{"red50(", 5},
		{"red50)", 5},
		{"red50[", 5},
		{"red50]", 5},
		{"red50// comment", 5},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Red50()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestRed50Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_r"},
		{"r_"},
		{"_re"},
		{"re_"},
		{"_red"},
		{"red_"},
		{"_red5"},
		{"red5_"},
		{"_red50"},
		{"red50_"},
		{"9713r"},
		{"r9713"},
		{"9713re"},
		{"re9713"},
		{"9713red"},
		{"red9713"},
		{"9713red5"},
		{"red59713"},
		{"9713red50"},
		{"red509713"},
		{"Zr"},
		{"rZ"},
		{"Zre"},
		{"reZ"},
		{"Zred"},
		{"redZ"},
		{"Zred5"},
		{"red5Z"},
		{"Zred50"},
		{"red50Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Red50()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestRed100Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"red100", 6},
		{"red100 ", 6},
		{"red100\n", 6},
		{"red100.", 6},
		{"red100:", 6},
		{"red100,", 6},
		{"red100\"", 6},
		{"red100(", 6},
		{"red100)", 6},
		{"red100[", 6},
		{"red100]", 6},
		{"red100// comment", 6},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Red100()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestRed100Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_r"},
		{"r_"},
		{"_re"},
		{"re_"},
		{"_red"},
		{"red_"},
		{"_red1"},
		{"red1_"},
		{"_red10"},
		{"red10_"},
		{"_red100"},
		{"red100_"},
		{"9713r"},
		{"r9713"},
		{"9713re"},
		{"re9713"},
		{"9713red"},
		{"red9713"},
		{"9713red1"},
		{"red19713"},
		{"9713red10"},
		{"red109713"},
		{"9713red100"},
		{"red1009713"},
		{"Zr"},
		{"rZ"},
		{"Zre"},
		{"reZ"},
		{"Zred"},
		{"redZ"},
		{"Zred1"},
		{"red1Z"},
		{"Zred10"},
		{"red10Z"},
		{"Zred100"},
		{"red100Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Red100()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestRed200Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"red200", 6},
		{"red200 ", 6},
		{"red200\n", 6},
		{"red200.", 6},
		{"red200:", 6},
		{"red200,", 6},
		{"red200\"", 6},
		{"red200(", 6},
		{"red200)", 6},
		{"red200[", 6},
		{"red200]", 6},
		{"red200// comment", 6},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Red200()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestRed200Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_r"},
		{"r_"},
		{"_re"},
		{"re_"},
		{"_red"},
		{"red_"},
		{"_red2"},
		{"red2_"},
		{"_red20"},
		{"red20_"},
		{"_red200"},
		{"red200_"},
		{"9713r"},
		{"r9713"},
		{"9713re"},
		{"re9713"},
		{"9713red"},
		{"red9713"},
		{"9713red2"},
		{"red29713"},
		{"9713red20"},
		{"red209713"},
		{"9713red200"},
		{"red2009713"},
		{"Zr"},
		{"rZ"},
		{"Zre"},
		{"reZ"},
		{"Zred"},
		{"redZ"},
		{"Zred2"},
		{"red2Z"},
		{"Zred20"},
		{"red20Z"},
		{"Zred200"},
		{"red200Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Red200()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestRed300Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"red300", 6},
		{"red300 ", 6},
		{"red300\n", 6},
		{"red300.", 6},
		{"red300:", 6},
		{"red300,", 6},
		{"red300\"", 6},
		{"red300(", 6},
		{"red300)", 6},
		{"red300[", 6},
		{"red300]", 6},
		{"red300// comment", 6},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Red300()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestRed300Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_r"},
		{"r_"},
		{"_re"},
		{"re_"},
		{"_red"},
		{"red_"},
		{"_red3"},
		{"red3_"},
		{"_red30"},
		{"red30_"},
		{"_red300"},
		{"red300_"},
		{"9713r"},
		{"r9713"},
		{"9713re"},
		{"re9713"},
		{"9713red"},
		{"red9713"},
		{"9713red3"},
		{"red39713"},
		{"9713red30"},
		{"red309713"},
		{"9713red300"},
		{"red3009713"},
		{"Zr"},
		{"rZ"},
		{"Zre"},
		{"reZ"},
		{"Zred"},
		{"redZ"},
		{"Zred3"},
		{"red3Z"},
		{"Zred30"},
		{"red30Z"},
		{"Zred300"},
		{"red300Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Red300()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestRed400Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"red400", 6},
		{"red400 ", 6},
		{"red400\n", 6},
		{"red400.", 6},
		{"red400:", 6},
		{"red400,", 6},
		{"red400\"", 6},
		{"red400(", 6},
		{"red400)", 6},
		{"red400[", 6},
		{"red400]", 6},
		{"red400// comment", 6},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Red400()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestRed400Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_r"},
		{"r_"},
		{"_re"},
		{"re_"},
		{"_red"},
		{"red_"},
		{"_red4"},
		{"red4_"},
		{"_red40"},
		{"red40_"},
		{"_red400"},
		{"red400_"},
		{"9713r"},
		{"r9713"},
		{"9713re"},
		{"re9713"},
		{"9713red"},
		{"red9713"},
		{"9713red4"},
		{"red49713"},
		{"9713red40"},
		{"red409713"},
		{"9713red400"},
		{"red4009713"},
		{"Zr"},
		{"rZ"},
		{"Zre"},
		{"reZ"},
		{"Zred"},
		{"redZ"},
		{"Zred4"},
		{"red4Z"},
		{"Zred40"},
		{"red40Z"},
		{"Zred400"},
		{"red400Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Red400()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestRed500Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"red500", 6},
		{"red500 ", 6},
		{"red500\n", 6},
		{"red500.", 6},
		{"red500:", 6},
		{"red500,", 6},
		{"red500\"", 6},
		{"red500(", 6},
		{"red500)", 6},
		{"red500[", 6},
		{"red500]", 6},
		{"red500// comment", 6},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Red500()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestRed500Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_r"},
		{"r_"},
		{"_re"},
		{"re_"},
		{"_red"},
		{"red_"},
		{"_red5"},
		{"red5_"},
		{"_red50"},
		{"red50_"},
		{"_red500"},
		{"red500_"},
		{"9713r"},
		{"r9713"},
		{"9713re"},
		{"re9713"},
		{"9713red"},
		{"red9713"},
		{"9713red5"},
		{"red59713"},
		{"9713red50"},
		{"red509713"},
		{"9713red500"},
		{"red5009713"},
		{"Zr"},
		{"rZ"},
		{"Zre"},
		{"reZ"},
		{"Zred"},
		{"redZ"},
		{"Zred5"},
		{"red5Z"},
		{"Zred50"},
		{"red50Z"},
		{"Zred500"},
		{"red500Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Red500()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestRed600Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"red600", 6},
		{"red600 ", 6},
		{"red600\n", 6},
		{"red600.", 6},
		{"red600:", 6},
		{"red600,", 6},
		{"red600\"", 6},
		{"red600(", 6},
		{"red600)", 6},
		{"red600[", 6},
		{"red600]", 6},
		{"red600// comment", 6},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Red600()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestRed600Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_r"},
		{"r_"},
		{"_re"},
		{"re_"},
		{"_red"},
		{"red_"},
		{"_red6"},
		{"red6_"},
		{"_red60"},
		{"red60_"},
		{"_red600"},
		{"red600_"},
		{"9713r"},
		{"r9713"},
		{"9713re"},
		{"re9713"},
		{"9713red"},
		{"red9713"},
		{"9713red6"},
		{"red69713"},
		{"9713red60"},
		{"red609713"},
		{"9713red600"},
		{"red6009713"},
		{"Zr"},
		{"rZ"},
		{"Zre"},
		{"reZ"},
		{"Zred"},
		{"redZ"},
		{"Zred6"},
		{"red6Z"},
		{"Zred60"},
		{"red60Z"},
		{"Zred600"},
		{"red600Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Red600()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestRed700Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"red700", 6},
		{"red700 ", 6},
		{"red700\n", 6},
		{"red700.", 6},
		{"red700:", 6},
		{"red700,", 6},
		{"red700\"", 6},
		{"red700(", 6},
		{"red700)", 6},
		{"red700[", 6},
		{"red700]", 6},
		{"red700// comment", 6},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Red700()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestRed700Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_r"},
		{"r_"},
		{"_re"},
		{"re_"},
		{"_red"},
		{"red_"},
		{"_red7"},
		{"red7_"},
		{"_red70"},
		{"red70_"},
		{"_red700"},
		{"red700_"},
		{"9713r"},
		{"r9713"},
		{"9713re"},
		{"re9713"},
		{"9713red"},
		{"red9713"},
		{"9713red7"},
		{"red79713"},
		{"9713red70"},
		{"red709713"},
		{"9713red700"},
		{"red7009713"},
		{"Zr"},
		{"rZ"},
		{"Zre"},
		{"reZ"},
		{"Zred"},
		{"redZ"},
		{"Zred7"},
		{"red7Z"},
		{"Zred70"},
		{"red70Z"},
		{"Zred700"},
		{"red700Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Red700()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestRed800Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"red800", 6},
		{"red800 ", 6},
		{"red800\n", 6},
		{"red800.", 6},
		{"red800:", 6},
		{"red800,", 6},
		{"red800\"", 6},
		{"red800(", 6},
		{"red800)", 6},
		{"red800[", 6},
		{"red800]", 6},
		{"red800// comment", 6},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Red800()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestRed800Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_r"},
		{"r_"},
		{"_re"},
		{"re_"},
		{"_red"},
		{"red_"},
		{"_red8"},
		{"red8_"},
		{"_red80"},
		{"red80_"},
		{"_red800"},
		{"red800_"},
		{"9713r"},
		{"r9713"},
		{"9713re"},
		{"re9713"},
		{"9713red"},
		{"red9713"},
		{"9713red8"},
		{"red89713"},
		{"9713red80"},
		{"red809713"},
		{"9713red800"},
		{"red8009713"},
		{"Zr"},
		{"rZ"},
		{"Zre"},
		{"reZ"},
		{"Zred"},
		{"redZ"},
		{"Zred8"},
		{"red8Z"},
		{"Zred80"},
		{"red80Z"},
		{"Zred800"},
		{"red800Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Red800()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestRed900Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"red900", 6},
		{"red900 ", 6},
		{"red900\n", 6},
		{"red900.", 6},
		{"red900:", 6},
		{"red900,", 6},
		{"red900\"", 6},
		{"red900(", 6},
		{"red900)", 6},
		{"red900[", 6},
		{"red900]", 6},
		{"red900// comment", 6},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Red900()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestRed900Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_r"},
		{"r_"},
		{"_re"},
		{"re_"},
		{"_red"},
		{"red_"},
		{"_red9"},
		{"red9_"},
		{"_red90"},
		{"red90_"},
		{"_red900"},
		{"red900_"},
		{"9713r"},
		{"r9713"},
		{"9713re"},
		{"re9713"},
		{"9713red"},
		{"red9713"},
		{"9713red9"},
		{"red99713"},
		{"9713red90"},
		{"red909713"},
		{"9713red900"},
		{"red9009713"},
		{"Zr"},
		{"rZ"},
		{"Zre"},
		{"reZ"},
		{"Zred"},
		{"redZ"},
		{"Zred9"},
		{"red9Z"},
		{"Zred90"},
		{"red90Z"},
		{"Zred900"},
		{"red900Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Red900()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestRedA100Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"redA100", 7},
		{"redA100 ", 7},
		{"redA100\n", 7},
		{"redA100.", 7},
		{"redA100:", 7},
		{"redA100,", 7},
		{"redA100\"", 7},
		{"redA100(", 7},
		{"redA100)", 7},
		{"redA100[", 7},
		{"redA100]", 7},
		{"redA100// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.RedA100()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestRedA100Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_r"},
		{"r_"},
		{"_re"},
		{"re_"},
		{"_red"},
		{"red_"},
		{"_redA"},
		{"redA_"},
		{"_redA1"},
		{"redA1_"},
		{"_redA10"},
		{"redA10_"},
		{"_redA100"},
		{"redA100_"},
		{"9713r"},
		{"r9713"},
		{"9713re"},
		{"re9713"},
		{"9713red"},
		{"red9713"},
		{"9713redA"},
		{"redA9713"},
		{"9713redA1"},
		{"redA19713"},
		{"9713redA10"},
		{"redA109713"},
		{"9713redA100"},
		{"redA1009713"},
		{"Zr"},
		{"rZ"},
		{"Zre"},
		{"reZ"},
		{"Zred"},
		{"redZ"},
		{"ZredA"},
		{"redAZ"},
		{"ZredA1"},
		{"redA1Z"},
		{"ZredA10"},
		{"redA10Z"},
		{"ZredA100"},
		{"redA100Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.RedA100()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestRedA200Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"redA200", 7},
		{"redA200 ", 7},
		{"redA200\n", 7},
		{"redA200.", 7},
		{"redA200:", 7},
		{"redA200,", 7},
		{"redA200\"", 7},
		{"redA200(", 7},
		{"redA200)", 7},
		{"redA200[", 7},
		{"redA200]", 7},
		{"redA200// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.RedA200()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestRedA200Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_r"},
		{"r_"},
		{"_re"},
		{"re_"},
		{"_red"},
		{"red_"},
		{"_redA"},
		{"redA_"},
		{"_redA2"},
		{"redA2_"},
		{"_redA20"},
		{"redA20_"},
		{"_redA200"},
		{"redA200_"},
		{"9713r"},
		{"r9713"},
		{"9713re"},
		{"re9713"},
		{"9713red"},
		{"red9713"},
		{"9713redA"},
		{"redA9713"},
		{"9713redA2"},
		{"redA29713"},
		{"9713redA20"},
		{"redA209713"},
		{"9713redA200"},
		{"redA2009713"},
		{"Zr"},
		{"rZ"},
		{"Zre"},
		{"reZ"},
		{"Zred"},
		{"redZ"},
		{"ZredA"},
		{"redAZ"},
		{"ZredA2"},
		{"redA2Z"},
		{"ZredA20"},
		{"redA20Z"},
		{"ZredA200"},
		{"redA200Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.RedA200()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestRedA400Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"redA400", 7},
		{"redA400 ", 7},
		{"redA400\n", 7},
		{"redA400.", 7},
		{"redA400:", 7},
		{"redA400,", 7},
		{"redA400\"", 7},
		{"redA400(", 7},
		{"redA400)", 7},
		{"redA400[", 7},
		{"redA400]", 7},
		{"redA400// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.RedA400()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestRedA400Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_r"},
		{"r_"},
		{"_re"},
		{"re_"},
		{"_red"},
		{"red_"},
		{"_redA"},
		{"redA_"},
		{"_redA4"},
		{"redA4_"},
		{"_redA40"},
		{"redA40_"},
		{"_redA400"},
		{"redA400_"},
		{"9713r"},
		{"r9713"},
		{"9713re"},
		{"re9713"},
		{"9713red"},
		{"red9713"},
		{"9713redA"},
		{"redA9713"},
		{"9713redA4"},
		{"redA49713"},
		{"9713redA40"},
		{"redA409713"},
		{"9713redA400"},
		{"redA4009713"},
		{"Zr"},
		{"rZ"},
		{"Zre"},
		{"reZ"},
		{"Zred"},
		{"redZ"},
		{"ZredA"},
		{"redAZ"},
		{"ZredA4"},
		{"redA4Z"},
		{"ZredA40"},
		{"redA40Z"},
		{"ZredA400"},
		{"redA400Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.RedA400()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestRedA700Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"redA700", 7},
		{"redA700 ", 7},
		{"redA700\n", 7},
		{"redA700.", 7},
		{"redA700:", 7},
		{"redA700,", 7},
		{"redA700\"", 7},
		{"redA700(", 7},
		{"redA700)", 7},
		{"redA700[", 7},
		{"redA700]", 7},
		{"redA700// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.RedA700()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestRedA700Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_r"},
		{"r_"},
		{"_re"},
		{"re_"},
		{"_red"},
		{"red_"},
		{"_redA"},
		{"redA_"},
		{"_redA7"},
		{"redA7_"},
		{"_redA70"},
		{"redA70_"},
		{"_redA700"},
		{"redA700_"},
		{"9713r"},
		{"r9713"},
		{"9713re"},
		{"re9713"},
		{"9713red"},
		{"red9713"},
		{"9713redA"},
		{"redA9713"},
		{"9713redA7"},
		{"redA79713"},
		{"9713redA70"},
		{"redA709713"},
		{"9713redA700"},
		{"redA7009713"},
		{"Zr"},
		{"rZ"},
		{"Zre"},
		{"reZ"},
		{"Zred"},
		{"redZ"},
		{"ZredA"},
		{"redAZ"},
		{"ZredA7"},
		{"redA7Z"},
		{"ZredA70"},
		{"redA70Z"},
		{"ZredA700"},
		{"redA700Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.RedA700()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestPink50Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"pink50", 6},
		{"pink50 ", 6},
		{"pink50\n", 6},
		{"pink50.", 6},
		{"pink50:", 6},
		{"pink50,", 6},
		{"pink50\"", 6},
		{"pink50(", 6},
		{"pink50)", 6},
		{"pink50[", 6},
		{"pink50]", 6},
		{"pink50// comment", 6},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Pink50()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestPink50Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_p"},
		{"p_"},
		{"_pi"},
		{"pi_"},
		{"_pin"},
		{"pin_"},
		{"_pink"},
		{"pink_"},
		{"_pink5"},
		{"pink5_"},
		{"_pink50"},
		{"pink50_"},
		{"9713p"},
		{"p9713"},
		{"9713pi"},
		{"pi9713"},
		{"9713pin"},
		{"pin9713"},
		{"9713pink"},
		{"pink9713"},
		{"9713pink5"},
		{"pink59713"},
		{"9713pink50"},
		{"pink509713"},
		{"Zp"},
		{"pZ"},
		{"Zpi"},
		{"piZ"},
		{"Zpin"},
		{"pinZ"},
		{"Zpink"},
		{"pinkZ"},
		{"Zpink5"},
		{"pink5Z"},
		{"Zpink50"},
		{"pink50Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Pink50()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestPink100Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"pink100", 7},
		{"pink100 ", 7},
		{"pink100\n", 7},
		{"pink100.", 7},
		{"pink100:", 7},
		{"pink100,", 7},
		{"pink100\"", 7},
		{"pink100(", 7},
		{"pink100)", 7},
		{"pink100[", 7},
		{"pink100]", 7},
		{"pink100// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Pink100()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestPink100Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_p"},
		{"p_"},
		{"_pi"},
		{"pi_"},
		{"_pin"},
		{"pin_"},
		{"_pink"},
		{"pink_"},
		{"_pink1"},
		{"pink1_"},
		{"_pink10"},
		{"pink10_"},
		{"_pink100"},
		{"pink100_"},
		{"9713p"},
		{"p9713"},
		{"9713pi"},
		{"pi9713"},
		{"9713pin"},
		{"pin9713"},
		{"9713pink"},
		{"pink9713"},
		{"9713pink1"},
		{"pink19713"},
		{"9713pink10"},
		{"pink109713"},
		{"9713pink100"},
		{"pink1009713"},
		{"Zp"},
		{"pZ"},
		{"Zpi"},
		{"piZ"},
		{"Zpin"},
		{"pinZ"},
		{"Zpink"},
		{"pinkZ"},
		{"Zpink1"},
		{"pink1Z"},
		{"Zpink10"},
		{"pink10Z"},
		{"Zpink100"},
		{"pink100Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Pink100()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestPink200Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"pink200", 7},
		{"pink200 ", 7},
		{"pink200\n", 7},
		{"pink200.", 7},
		{"pink200:", 7},
		{"pink200,", 7},
		{"pink200\"", 7},
		{"pink200(", 7},
		{"pink200)", 7},
		{"pink200[", 7},
		{"pink200]", 7},
		{"pink200// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Pink200()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestPink200Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_p"},
		{"p_"},
		{"_pi"},
		{"pi_"},
		{"_pin"},
		{"pin_"},
		{"_pink"},
		{"pink_"},
		{"_pink2"},
		{"pink2_"},
		{"_pink20"},
		{"pink20_"},
		{"_pink200"},
		{"pink200_"},
		{"9713p"},
		{"p9713"},
		{"9713pi"},
		{"pi9713"},
		{"9713pin"},
		{"pin9713"},
		{"9713pink"},
		{"pink9713"},
		{"9713pink2"},
		{"pink29713"},
		{"9713pink20"},
		{"pink209713"},
		{"9713pink200"},
		{"pink2009713"},
		{"Zp"},
		{"pZ"},
		{"Zpi"},
		{"piZ"},
		{"Zpin"},
		{"pinZ"},
		{"Zpink"},
		{"pinkZ"},
		{"Zpink2"},
		{"pink2Z"},
		{"Zpink20"},
		{"pink20Z"},
		{"Zpink200"},
		{"pink200Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Pink200()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestPink300Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"pink300", 7},
		{"pink300 ", 7},
		{"pink300\n", 7},
		{"pink300.", 7},
		{"pink300:", 7},
		{"pink300,", 7},
		{"pink300\"", 7},
		{"pink300(", 7},
		{"pink300)", 7},
		{"pink300[", 7},
		{"pink300]", 7},
		{"pink300// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Pink300()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestPink300Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_p"},
		{"p_"},
		{"_pi"},
		{"pi_"},
		{"_pin"},
		{"pin_"},
		{"_pink"},
		{"pink_"},
		{"_pink3"},
		{"pink3_"},
		{"_pink30"},
		{"pink30_"},
		{"_pink300"},
		{"pink300_"},
		{"9713p"},
		{"p9713"},
		{"9713pi"},
		{"pi9713"},
		{"9713pin"},
		{"pin9713"},
		{"9713pink"},
		{"pink9713"},
		{"9713pink3"},
		{"pink39713"},
		{"9713pink30"},
		{"pink309713"},
		{"9713pink300"},
		{"pink3009713"},
		{"Zp"},
		{"pZ"},
		{"Zpi"},
		{"piZ"},
		{"Zpin"},
		{"pinZ"},
		{"Zpink"},
		{"pinkZ"},
		{"Zpink3"},
		{"pink3Z"},
		{"Zpink30"},
		{"pink30Z"},
		{"Zpink300"},
		{"pink300Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Pink300()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestPink400Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"pink400", 7},
		{"pink400 ", 7},
		{"pink400\n", 7},
		{"pink400.", 7},
		{"pink400:", 7},
		{"pink400,", 7},
		{"pink400\"", 7},
		{"pink400(", 7},
		{"pink400)", 7},
		{"pink400[", 7},
		{"pink400]", 7},
		{"pink400// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Pink400()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestPink400Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_p"},
		{"p_"},
		{"_pi"},
		{"pi_"},
		{"_pin"},
		{"pin_"},
		{"_pink"},
		{"pink_"},
		{"_pink4"},
		{"pink4_"},
		{"_pink40"},
		{"pink40_"},
		{"_pink400"},
		{"pink400_"},
		{"9713p"},
		{"p9713"},
		{"9713pi"},
		{"pi9713"},
		{"9713pin"},
		{"pin9713"},
		{"9713pink"},
		{"pink9713"},
		{"9713pink4"},
		{"pink49713"},
		{"9713pink40"},
		{"pink409713"},
		{"9713pink400"},
		{"pink4009713"},
		{"Zp"},
		{"pZ"},
		{"Zpi"},
		{"piZ"},
		{"Zpin"},
		{"pinZ"},
		{"Zpink"},
		{"pinkZ"},
		{"Zpink4"},
		{"pink4Z"},
		{"Zpink40"},
		{"pink40Z"},
		{"Zpink400"},
		{"pink400Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Pink400()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestPink500Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"pink500", 7},
		{"pink500 ", 7},
		{"pink500\n", 7},
		{"pink500.", 7},
		{"pink500:", 7},
		{"pink500,", 7},
		{"pink500\"", 7},
		{"pink500(", 7},
		{"pink500)", 7},
		{"pink500[", 7},
		{"pink500]", 7},
		{"pink500// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Pink500()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestPink500Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_p"},
		{"p_"},
		{"_pi"},
		{"pi_"},
		{"_pin"},
		{"pin_"},
		{"_pink"},
		{"pink_"},
		{"_pink5"},
		{"pink5_"},
		{"_pink50"},
		{"pink50_"},
		{"_pink500"},
		{"pink500_"},
		{"9713p"},
		{"p9713"},
		{"9713pi"},
		{"pi9713"},
		{"9713pin"},
		{"pin9713"},
		{"9713pink"},
		{"pink9713"},
		{"9713pink5"},
		{"pink59713"},
		{"9713pink50"},
		{"pink509713"},
		{"9713pink500"},
		{"pink5009713"},
		{"Zp"},
		{"pZ"},
		{"Zpi"},
		{"piZ"},
		{"Zpin"},
		{"pinZ"},
		{"Zpink"},
		{"pinkZ"},
		{"Zpink5"},
		{"pink5Z"},
		{"Zpink50"},
		{"pink50Z"},
		{"Zpink500"},
		{"pink500Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Pink500()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestPink600Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"pink600", 7},
		{"pink600 ", 7},
		{"pink600\n", 7},
		{"pink600.", 7},
		{"pink600:", 7},
		{"pink600,", 7},
		{"pink600\"", 7},
		{"pink600(", 7},
		{"pink600)", 7},
		{"pink600[", 7},
		{"pink600]", 7},
		{"pink600// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Pink600()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestPink600Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_p"},
		{"p_"},
		{"_pi"},
		{"pi_"},
		{"_pin"},
		{"pin_"},
		{"_pink"},
		{"pink_"},
		{"_pink6"},
		{"pink6_"},
		{"_pink60"},
		{"pink60_"},
		{"_pink600"},
		{"pink600_"},
		{"9713p"},
		{"p9713"},
		{"9713pi"},
		{"pi9713"},
		{"9713pin"},
		{"pin9713"},
		{"9713pink"},
		{"pink9713"},
		{"9713pink6"},
		{"pink69713"},
		{"9713pink60"},
		{"pink609713"},
		{"9713pink600"},
		{"pink6009713"},
		{"Zp"},
		{"pZ"},
		{"Zpi"},
		{"piZ"},
		{"Zpin"},
		{"pinZ"},
		{"Zpink"},
		{"pinkZ"},
		{"Zpink6"},
		{"pink6Z"},
		{"Zpink60"},
		{"pink60Z"},
		{"Zpink600"},
		{"pink600Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Pink600()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestPink700Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"pink700", 7},
		{"pink700 ", 7},
		{"pink700\n", 7},
		{"pink700.", 7},
		{"pink700:", 7},
		{"pink700,", 7},
		{"pink700\"", 7},
		{"pink700(", 7},
		{"pink700)", 7},
		{"pink700[", 7},
		{"pink700]", 7},
		{"pink700// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Pink700()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestPink700Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_p"},
		{"p_"},
		{"_pi"},
		{"pi_"},
		{"_pin"},
		{"pin_"},
		{"_pink"},
		{"pink_"},
		{"_pink7"},
		{"pink7_"},
		{"_pink70"},
		{"pink70_"},
		{"_pink700"},
		{"pink700_"},
		{"9713p"},
		{"p9713"},
		{"9713pi"},
		{"pi9713"},
		{"9713pin"},
		{"pin9713"},
		{"9713pink"},
		{"pink9713"},
		{"9713pink7"},
		{"pink79713"},
		{"9713pink70"},
		{"pink709713"},
		{"9713pink700"},
		{"pink7009713"},
		{"Zp"},
		{"pZ"},
		{"Zpi"},
		{"piZ"},
		{"Zpin"},
		{"pinZ"},
		{"Zpink"},
		{"pinkZ"},
		{"Zpink7"},
		{"pink7Z"},
		{"Zpink70"},
		{"pink70Z"},
		{"Zpink700"},
		{"pink700Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Pink700()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestPink800Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"pink800", 7},
		{"pink800 ", 7},
		{"pink800\n", 7},
		{"pink800.", 7},
		{"pink800:", 7},
		{"pink800,", 7},
		{"pink800\"", 7},
		{"pink800(", 7},
		{"pink800)", 7},
		{"pink800[", 7},
		{"pink800]", 7},
		{"pink800// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Pink800()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestPink800Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_p"},
		{"p_"},
		{"_pi"},
		{"pi_"},
		{"_pin"},
		{"pin_"},
		{"_pink"},
		{"pink_"},
		{"_pink8"},
		{"pink8_"},
		{"_pink80"},
		{"pink80_"},
		{"_pink800"},
		{"pink800_"},
		{"9713p"},
		{"p9713"},
		{"9713pi"},
		{"pi9713"},
		{"9713pin"},
		{"pin9713"},
		{"9713pink"},
		{"pink9713"},
		{"9713pink8"},
		{"pink89713"},
		{"9713pink80"},
		{"pink809713"},
		{"9713pink800"},
		{"pink8009713"},
		{"Zp"},
		{"pZ"},
		{"Zpi"},
		{"piZ"},
		{"Zpin"},
		{"pinZ"},
		{"Zpink"},
		{"pinkZ"},
		{"Zpink8"},
		{"pink8Z"},
		{"Zpink80"},
		{"pink80Z"},
		{"Zpink800"},
		{"pink800Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Pink800()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestPink900Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"pink900", 7},
		{"pink900 ", 7},
		{"pink900\n", 7},
		{"pink900.", 7},
		{"pink900:", 7},
		{"pink900,", 7},
		{"pink900\"", 7},
		{"pink900(", 7},
		{"pink900)", 7},
		{"pink900[", 7},
		{"pink900]", 7},
		{"pink900// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Pink900()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestPink900Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_p"},
		{"p_"},
		{"_pi"},
		{"pi_"},
		{"_pin"},
		{"pin_"},
		{"_pink"},
		{"pink_"},
		{"_pink9"},
		{"pink9_"},
		{"_pink90"},
		{"pink90_"},
		{"_pink900"},
		{"pink900_"},
		{"9713p"},
		{"p9713"},
		{"9713pi"},
		{"pi9713"},
		{"9713pin"},
		{"pin9713"},
		{"9713pink"},
		{"pink9713"},
		{"9713pink9"},
		{"pink99713"},
		{"9713pink90"},
		{"pink909713"},
		{"9713pink900"},
		{"pink9009713"},
		{"Zp"},
		{"pZ"},
		{"Zpi"},
		{"piZ"},
		{"Zpin"},
		{"pinZ"},
		{"Zpink"},
		{"pinkZ"},
		{"Zpink9"},
		{"pink9Z"},
		{"Zpink90"},
		{"pink90Z"},
		{"Zpink900"},
		{"pink900Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Pink900()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestPinkA100Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"pinkA100", 8},
		{"pinkA100 ", 8},
		{"pinkA100\n", 8},
		{"pinkA100.", 8},
		{"pinkA100:", 8},
		{"pinkA100,", 8},
		{"pinkA100\"", 8},
		{"pinkA100(", 8},
		{"pinkA100)", 8},
		{"pinkA100[", 8},
		{"pinkA100]", 8},
		{"pinkA100// comment", 8},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.PinkA100()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestPinkA100Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_p"},
		{"p_"},
		{"_pi"},
		{"pi_"},
		{"_pin"},
		{"pin_"},
		{"_pink"},
		{"pink_"},
		{"_pinkA"},
		{"pinkA_"},
		{"_pinkA1"},
		{"pinkA1_"},
		{"_pinkA10"},
		{"pinkA10_"},
		{"_pinkA100"},
		{"pinkA100_"},
		{"9713p"},
		{"p9713"},
		{"9713pi"},
		{"pi9713"},
		{"9713pin"},
		{"pin9713"},
		{"9713pink"},
		{"pink9713"},
		{"9713pinkA"},
		{"pinkA9713"},
		{"9713pinkA1"},
		{"pinkA19713"},
		{"9713pinkA10"},
		{"pinkA109713"},
		{"9713pinkA100"},
		{"pinkA1009713"},
		{"Zp"},
		{"pZ"},
		{"Zpi"},
		{"piZ"},
		{"Zpin"},
		{"pinZ"},
		{"Zpink"},
		{"pinkZ"},
		{"ZpinkA"},
		{"pinkAZ"},
		{"ZpinkA1"},
		{"pinkA1Z"},
		{"ZpinkA10"},
		{"pinkA10Z"},
		{"ZpinkA100"},
		{"pinkA100Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.PinkA100()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestPinkA200Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"pinkA200", 8},
		{"pinkA200 ", 8},
		{"pinkA200\n", 8},
		{"pinkA200.", 8},
		{"pinkA200:", 8},
		{"pinkA200,", 8},
		{"pinkA200\"", 8},
		{"pinkA200(", 8},
		{"pinkA200)", 8},
		{"pinkA200[", 8},
		{"pinkA200]", 8},
		{"pinkA200// comment", 8},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.PinkA200()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestPinkA200Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_p"},
		{"p_"},
		{"_pi"},
		{"pi_"},
		{"_pin"},
		{"pin_"},
		{"_pink"},
		{"pink_"},
		{"_pinkA"},
		{"pinkA_"},
		{"_pinkA2"},
		{"pinkA2_"},
		{"_pinkA20"},
		{"pinkA20_"},
		{"_pinkA200"},
		{"pinkA200_"},
		{"9713p"},
		{"p9713"},
		{"9713pi"},
		{"pi9713"},
		{"9713pin"},
		{"pin9713"},
		{"9713pink"},
		{"pink9713"},
		{"9713pinkA"},
		{"pinkA9713"},
		{"9713pinkA2"},
		{"pinkA29713"},
		{"9713pinkA20"},
		{"pinkA209713"},
		{"9713pinkA200"},
		{"pinkA2009713"},
		{"Zp"},
		{"pZ"},
		{"Zpi"},
		{"piZ"},
		{"Zpin"},
		{"pinZ"},
		{"Zpink"},
		{"pinkZ"},
		{"ZpinkA"},
		{"pinkAZ"},
		{"ZpinkA2"},
		{"pinkA2Z"},
		{"ZpinkA20"},
		{"pinkA20Z"},
		{"ZpinkA200"},
		{"pinkA200Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.PinkA200()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestPinkA400Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"pinkA400", 8},
		{"pinkA400 ", 8},
		{"pinkA400\n", 8},
		{"pinkA400.", 8},
		{"pinkA400:", 8},
		{"pinkA400,", 8},
		{"pinkA400\"", 8},
		{"pinkA400(", 8},
		{"pinkA400)", 8},
		{"pinkA400[", 8},
		{"pinkA400]", 8},
		{"pinkA400// comment", 8},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.PinkA400()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestPinkA400Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_p"},
		{"p_"},
		{"_pi"},
		{"pi_"},
		{"_pin"},
		{"pin_"},
		{"_pink"},
		{"pink_"},
		{"_pinkA"},
		{"pinkA_"},
		{"_pinkA4"},
		{"pinkA4_"},
		{"_pinkA40"},
		{"pinkA40_"},
		{"_pinkA400"},
		{"pinkA400_"},
		{"9713p"},
		{"p9713"},
		{"9713pi"},
		{"pi9713"},
		{"9713pin"},
		{"pin9713"},
		{"9713pink"},
		{"pink9713"},
		{"9713pinkA"},
		{"pinkA9713"},
		{"9713pinkA4"},
		{"pinkA49713"},
		{"9713pinkA40"},
		{"pinkA409713"},
		{"9713pinkA400"},
		{"pinkA4009713"},
		{"Zp"},
		{"pZ"},
		{"Zpi"},
		{"piZ"},
		{"Zpin"},
		{"pinZ"},
		{"Zpink"},
		{"pinkZ"},
		{"ZpinkA"},
		{"pinkAZ"},
		{"ZpinkA4"},
		{"pinkA4Z"},
		{"ZpinkA40"},
		{"pinkA40Z"},
		{"ZpinkA400"},
		{"pinkA400Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.PinkA400()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestPinkA700Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"pinkA700", 8},
		{"pinkA700 ", 8},
		{"pinkA700\n", 8},
		{"pinkA700.", 8},
		{"pinkA700:", 8},
		{"pinkA700,", 8},
		{"pinkA700\"", 8},
		{"pinkA700(", 8},
		{"pinkA700)", 8},
		{"pinkA700[", 8},
		{"pinkA700]", 8},
		{"pinkA700// comment", 8},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.PinkA700()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestPinkA700Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_p"},
		{"p_"},
		{"_pi"},
		{"pi_"},
		{"_pin"},
		{"pin_"},
		{"_pink"},
		{"pink_"},
		{"_pinkA"},
		{"pinkA_"},
		{"_pinkA7"},
		{"pinkA7_"},
		{"_pinkA70"},
		{"pinkA70_"},
		{"_pinkA700"},
		{"pinkA700_"},
		{"9713p"},
		{"p9713"},
		{"9713pi"},
		{"pi9713"},
		{"9713pin"},
		{"pin9713"},
		{"9713pink"},
		{"pink9713"},
		{"9713pinkA"},
		{"pinkA9713"},
		{"9713pinkA7"},
		{"pinkA79713"},
		{"9713pinkA70"},
		{"pinkA709713"},
		{"9713pinkA700"},
		{"pinkA7009713"},
		{"Zp"},
		{"pZ"},
		{"Zpi"},
		{"piZ"},
		{"Zpin"},
		{"pinZ"},
		{"Zpink"},
		{"pinkZ"},
		{"ZpinkA"},
		{"pinkAZ"},
		{"ZpinkA7"},
		{"pinkA7Z"},
		{"ZpinkA70"},
		{"pinkA70Z"},
		{"ZpinkA700"},
		{"pinkA700Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.PinkA700()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestPurple50Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"purple50", 8},
		{"purple50 ", 8},
		{"purple50\n", 8},
		{"purple50.", 8},
		{"purple50:", 8},
		{"purple50,", 8},
		{"purple50\"", 8},
		{"purple50(", 8},
		{"purple50)", 8},
		{"purple50[", 8},
		{"purple50]", 8},
		{"purple50// comment", 8},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Purple50()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestPurple50Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_p"},
		{"p_"},
		{"_pu"},
		{"pu_"},
		{"_pur"},
		{"pur_"},
		{"_purp"},
		{"purp_"},
		{"_purpl"},
		{"purpl_"},
		{"_purple"},
		{"purple_"},
		{"_purple5"},
		{"purple5_"},
		{"_purple50"},
		{"purple50_"},
		{"9713p"},
		{"p9713"},
		{"9713pu"},
		{"pu9713"},
		{"9713pur"},
		{"pur9713"},
		{"9713purp"},
		{"purp9713"},
		{"9713purpl"},
		{"purpl9713"},
		{"9713purple"},
		{"purple9713"},
		{"9713purple5"},
		{"purple59713"},
		{"9713purple50"},
		{"purple509713"},
		{"Zp"},
		{"pZ"},
		{"Zpu"},
		{"puZ"},
		{"Zpur"},
		{"purZ"},
		{"Zpurp"},
		{"purpZ"},
		{"Zpurpl"},
		{"purplZ"},
		{"Zpurple"},
		{"purpleZ"},
		{"Zpurple5"},
		{"purple5Z"},
		{"Zpurple50"},
		{"purple50Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Purple50()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestPurple100Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"purple100", 9},
		{"purple100 ", 9},
		{"purple100\n", 9},
		{"purple100.", 9},
		{"purple100:", 9},
		{"purple100,", 9},
		{"purple100\"", 9},
		{"purple100(", 9},
		{"purple100)", 9},
		{"purple100[", 9},
		{"purple100]", 9},
		{"purple100// comment", 9},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Purple100()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestPurple100Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_p"},
		{"p_"},
		{"_pu"},
		{"pu_"},
		{"_pur"},
		{"pur_"},
		{"_purp"},
		{"purp_"},
		{"_purpl"},
		{"purpl_"},
		{"_purple"},
		{"purple_"},
		{"_purple1"},
		{"purple1_"},
		{"_purple10"},
		{"purple10_"},
		{"_purple100"},
		{"purple100_"},
		{"9713p"},
		{"p9713"},
		{"9713pu"},
		{"pu9713"},
		{"9713pur"},
		{"pur9713"},
		{"9713purp"},
		{"purp9713"},
		{"9713purpl"},
		{"purpl9713"},
		{"9713purple"},
		{"purple9713"},
		{"9713purple1"},
		{"purple19713"},
		{"9713purple10"},
		{"purple109713"},
		{"9713purple100"},
		{"purple1009713"},
		{"Zp"},
		{"pZ"},
		{"Zpu"},
		{"puZ"},
		{"Zpur"},
		{"purZ"},
		{"Zpurp"},
		{"purpZ"},
		{"Zpurpl"},
		{"purplZ"},
		{"Zpurple"},
		{"purpleZ"},
		{"Zpurple1"},
		{"purple1Z"},
		{"Zpurple10"},
		{"purple10Z"},
		{"Zpurple100"},
		{"purple100Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Purple100()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestPurple200Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"purple200", 9},
		{"purple200 ", 9},
		{"purple200\n", 9},
		{"purple200.", 9},
		{"purple200:", 9},
		{"purple200,", 9},
		{"purple200\"", 9},
		{"purple200(", 9},
		{"purple200)", 9},
		{"purple200[", 9},
		{"purple200]", 9},
		{"purple200// comment", 9},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Purple200()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestPurple200Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_p"},
		{"p_"},
		{"_pu"},
		{"pu_"},
		{"_pur"},
		{"pur_"},
		{"_purp"},
		{"purp_"},
		{"_purpl"},
		{"purpl_"},
		{"_purple"},
		{"purple_"},
		{"_purple2"},
		{"purple2_"},
		{"_purple20"},
		{"purple20_"},
		{"_purple200"},
		{"purple200_"},
		{"9713p"},
		{"p9713"},
		{"9713pu"},
		{"pu9713"},
		{"9713pur"},
		{"pur9713"},
		{"9713purp"},
		{"purp9713"},
		{"9713purpl"},
		{"purpl9713"},
		{"9713purple"},
		{"purple9713"},
		{"9713purple2"},
		{"purple29713"},
		{"9713purple20"},
		{"purple209713"},
		{"9713purple200"},
		{"purple2009713"},
		{"Zp"},
		{"pZ"},
		{"Zpu"},
		{"puZ"},
		{"Zpur"},
		{"purZ"},
		{"Zpurp"},
		{"purpZ"},
		{"Zpurpl"},
		{"purplZ"},
		{"Zpurple"},
		{"purpleZ"},
		{"Zpurple2"},
		{"purple2Z"},
		{"Zpurple20"},
		{"purple20Z"},
		{"Zpurple200"},
		{"purple200Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Purple200()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestPurple300Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"purple300", 9},
		{"purple300 ", 9},
		{"purple300\n", 9},
		{"purple300.", 9},
		{"purple300:", 9},
		{"purple300,", 9},
		{"purple300\"", 9},
		{"purple300(", 9},
		{"purple300)", 9},
		{"purple300[", 9},
		{"purple300]", 9},
		{"purple300// comment", 9},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Purple300()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestPurple300Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_p"},
		{"p_"},
		{"_pu"},
		{"pu_"},
		{"_pur"},
		{"pur_"},
		{"_purp"},
		{"purp_"},
		{"_purpl"},
		{"purpl_"},
		{"_purple"},
		{"purple_"},
		{"_purple3"},
		{"purple3_"},
		{"_purple30"},
		{"purple30_"},
		{"_purple300"},
		{"purple300_"},
		{"9713p"},
		{"p9713"},
		{"9713pu"},
		{"pu9713"},
		{"9713pur"},
		{"pur9713"},
		{"9713purp"},
		{"purp9713"},
		{"9713purpl"},
		{"purpl9713"},
		{"9713purple"},
		{"purple9713"},
		{"9713purple3"},
		{"purple39713"},
		{"9713purple30"},
		{"purple309713"},
		{"9713purple300"},
		{"purple3009713"},
		{"Zp"},
		{"pZ"},
		{"Zpu"},
		{"puZ"},
		{"Zpur"},
		{"purZ"},
		{"Zpurp"},
		{"purpZ"},
		{"Zpurpl"},
		{"purplZ"},
		{"Zpurple"},
		{"purpleZ"},
		{"Zpurple3"},
		{"purple3Z"},
		{"Zpurple30"},
		{"purple30Z"},
		{"Zpurple300"},
		{"purple300Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Purple300()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestPurple400Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"purple400", 9},
		{"purple400 ", 9},
		{"purple400\n", 9},
		{"purple400.", 9},
		{"purple400:", 9},
		{"purple400,", 9},
		{"purple400\"", 9},
		{"purple400(", 9},
		{"purple400)", 9},
		{"purple400[", 9},
		{"purple400]", 9},
		{"purple400// comment", 9},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Purple400()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestPurple400Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_p"},
		{"p_"},
		{"_pu"},
		{"pu_"},
		{"_pur"},
		{"pur_"},
		{"_purp"},
		{"purp_"},
		{"_purpl"},
		{"purpl_"},
		{"_purple"},
		{"purple_"},
		{"_purple4"},
		{"purple4_"},
		{"_purple40"},
		{"purple40_"},
		{"_purple400"},
		{"purple400_"},
		{"9713p"},
		{"p9713"},
		{"9713pu"},
		{"pu9713"},
		{"9713pur"},
		{"pur9713"},
		{"9713purp"},
		{"purp9713"},
		{"9713purpl"},
		{"purpl9713"},
		{"9713purple"},
		{"purple9713"},
		{"9713purple4"},
		{"purple49713"},
		{"9713purple40"},
		{"purple409713"},
		{"9713purple400"},
		{"purple4009713"},
		{"Zp"},
		{"pZ"},
		{"Zpu"},
		{"puZ"},
		{"Zpur"},
		{"purZ"},
		{"Zpurp"},
		{"purpZ"},
		{"Zpurpl"},
		{"purplZ"},
		{"Zpurple"},
		{"purpleZ"},
		{"Zpurple4"},
		{"purple4Z"},
		{"Zpurple40"},
		{"purple40Z"},
		{"Zpurple400"},
		{"purple400Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Purple400()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestPurple500Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"purple500", 9},
		{"purple500 ", 9},
		{"purple500\n", 9},
		{"purple500.", 9},
		{"purple500:", 9},
		{"purple500,", 9},
		{"purple500\"", 9},
		{"purple500(", 9},
		{"purple500)", 9},
		{"purple500[", 9},
		{"purple500]", 9},
		{"purple500// comment", 9},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Purple500()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestPurple500Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_p"},
		{"p_"},
		{"_pu"},
		{"pu_"},
		{"_pur"},
		{"pur_"},
		{"_purp"},
		{"purp_"},
		{"_purpl"},
		{"purpl_"},
		{"_purple"},
		{"purple_"},
		{"_purple5"},
		{"purple5_"},
		{"_purple50"},
		{"purple50_"},
		{"_purple500"},
		{"purple500_"},
		{"9713p"},
		{"p9713"},
		{"9713pu"},
		{"pu9713"},
		{"9713pur"},
		{"pur9713"},
		{"9713purp"},
		{"purp9713"},
		{"9713purpl"},
		{"purpl9713"},
		{"9713purple"},
		{"purple9713"},
		{"9713purple5"},
		{"purple59713"},
		{"9713purple50"},
		{"purple509713"},
		{"9713purple500"},
		{"purple5009713"},
		{"Zp"},
		{"pZ"},
		{"Zpu"},
		{"puZ"},
		{"Zpur"},
		{"purZ"},
		{"Zpurp"},
		{"purpZ"},
		{"Zpurpl"},
		{"purplZ"},
		{"Zpurple"},
		{"purpleZ"},
		{"Zpurple5"},
		{"purple5Z"},
		{"Zpurple50"},
		{"purple50Z"},
		{"Zpurple500"},
		{"purple500Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Purple500()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestPurple600Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"purple600", 9},
		{"purple600 ", 9},
		{"purple600\n", 9},
		{"purple600.", 9},
		{"purple600:", 9},
		{"purple600,", 9},
		{"purple600\"", 9},
		{"purple600(", 9},
		{"purple600)", 9},
		{"purple600[", 9},
		{"purple600]", 9},
		{"purple600// comment", 9},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Purple600()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestPurple600Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_p"},
		{"p_"},
		{"_pu"},
		{"pu_"},
		{"_pur"},
		{"pur_"},
		{"_purp"},
		{"purp_"},
		{"_purpl"},
		{"purpl_"},
		{"_purple"},
		{"purple_"},
		{"_purple6"},
		{"purple6_"},
		{"_purple60"},
		{"purple60_"},
		{"_purple600"},
		{"purple600_"},
		{"9713p"},
		{"p9713"},
		{"9713pu"},
		{"pu9713"},
		{"9713pur"},
		{"pur9713"},
		{"9713purp"},
		{"purp9713"},
		{"9713purpl"},
		{"purpl9713"},
		{"9713purple"},
		{"purple9713"},
		{"9713purple6"},
		{"purple69713"},
		{"9713purple60"},
		{"purple609713"},
		{"9713purple600"},
		{"purple6009713"},
		{"Zp"},
		{"pZ"},
		{"Zpu"},
		{"puZ"},
		{"Zpur"},
		{"purZ"},
		{"Zpurp"},
		{"purpZ"},
		{"Zpurpl"},
		{"purplZ"},
		{"Zpurple"},
		{"purpleZ"},
		{"Zpurple6"},
		{"purple6Z"},
		{"Zpurple60"},
		{"purple60Z"},
		{"Zpurple600"},
		{"purple600Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Purple600()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestPurple700Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"purple700", 9},
		{"purple700 ", 9},
		{"purple700\n", 9},
		{"purple700.", 9},
		{"purple700:", 9},
		{"purple700,", 9},
		{"purple700\"", 9},
		{"purple700(", 9},
		{"purple700)", 9},
		{"purple700[", 9},
		{"purple700]", 9},
		{"purple700// comment", 9},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Purple700()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestPurple700Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_p"},
		{"p_"},
		{"_pu"},
		{"pu_"},
		{"_pur"},
		{"pur_"},
		{"_purp"},
		{"purp_"},
		{"_purpl"},
		{"purpl_"},
		{"_purple"},
		{"purple_"},
		{"_purple7"},
		{"purple7_"},
		{"_purple70"},
		{"purple70_"},
		{"_purple700"},
		{"purple700_"},
		{"9713p"},
		{"p9713"},
		{"9713pu"},
		{"pu9713"},
		{"9713pur"},
		{"pur9713"},
		{"9713purp"},
		{"purp9713"},
		{"9713purpl"},
		{"purpl9713"},
		{"9713purple"},
		{"purple9713"},
		{"9713purple7"},
		{"purple79713"},
		{"9713purple70"},
		{"purple709713"},
		{"9713purple700"},
		{"purple7009713"},
		{"Zp"},
		{"pZ"},
		{"Zpu"},
		{"puZ"},
		{"Zpur"},
		{"purZ"},
		{"Zpurp"},
		{"purpZ"},
		{"Zpurpl"},
		{"purplZ"},
		{"Zpurple"},
		{"purpleZ"},
		{"Zpurple7"},
		{"purple7Z"},
		{"Zpurple70"},
		{"purple70Z"},
		{"Zpurple700"},
		{"purple700Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Purple700()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestPurple800Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"purple800", 9},
		{"purple800 ", 9},
		{"purple800\n", 9},
		{"purple800.", 9},
		{"purple800:", 9},
		{"purple800,", 9},
		{"purple800\"", 9},
		{"purple800(", 9},
		{"purple800)", 9},
		{"purple800[", 9},
		{"purple800]", 9},
		{"purple800// comment", 9},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Purple800()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestPurple800Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_p"},
		{"p_"},
		{"_pu"},
		{"pu_"},
		{"_pur"},
		{"pur_"},
		{"_purp"},
		{"purp_"},
		{"_purpl"},
		{"purpl_"},
		{"_purple"},
		{"purple_"},
		{"_purple8"},
		{"purple8_"},
		{"_purple80"},
		{"purple80_"},
		{"_purple800"},
		{"purple800_"},
		{"9713p"},
		{"p9713"},
		{"9713pu"},
		{"pu9713"},
		{"9713pur"},
		{"pur9713"},
		{"9713purp"},
		{"purp9713"},
		{"9713purpl"},
		{"purpl9713"},
		{"9713purple"},
		{"purple9713"},
		{"9713purple8"},
		{"purple89713"},
		{"9713purple80"},
		{"purple809713"},
		{"9713purple800"},
		{"purple8009713"},
		{"Zp"},
		{"pZ"},
		{"Zpu"},
		{"puZ"},
		{"Zpur"},
		{"purZ"},
		{"Zpurp"},
		{"purpZ"},
		{"Zpurpl"},
		{"purplZ"},
		{"Zpurple"},
		{"purpleZ"},
		{"Zpurple8"},
		{"purple8Z"},
		{"Zpurple80"},
		{"purple80Z"},
		{"Zpurple800"},
		{"purple800Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Purple800()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestPurple900Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"purple900", 9},
		{"purple900 ", 9},
		{"purple900\n", 9},
		{"purple900.", 9},
		{"purple900:", 9},
		{"purple900,", 9},
		{"purple900\"", 9},
		{"purple900(", 9},
		{"purple900)", 9},
		{"purple900[", 9},
		{"purple900]", 9},
		{"purple900// comment", 9},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Purple900()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestPurple900Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_p"},
		{"p_"},
		{"_pu"},
		{"pu_"},
		{"_pur"},
		{"pur_"},
		{"_purp"},
		{"purp_"},
		{"_purpl"},
		{"purpl_"},
		{"_purple"},
		{"purple_"},
		{"_purple9"},
		{"purple9_"},
		{"_purple90"},
		{"purple90_"},
		{"_purple900"},
		{"purple900_"},
		{"9713p"},
		{"p9713"},
		{"9713pu"},
		{"pu9713"},
		{"9713pur"},
		{"pur9713"},
		{"9713purp"},
		{"purp9713"},
		{"9713purpl"},
		{"purpl9713"},
		{"9713purple"},
		{"purple9713"},
		{"9713purple9"},
		{"purple99713"},
		{"9713purple90"},
		{"purple909713"},
		{"9713purple900"},
		{"purple9009713"},
		{"Zp"},
		{"pZ"},
		{"Zpu"},
		{"puZ"},
		{"Zpur"},
		{"purZ"},
		{"Zpurp"},
		{"purpZ"},
		{"Zpurpl"},
		{"purplZ"},
		{"Zpurple"},
		{"purpleZ"},
		{"Zpurple9"},
		{"purple9Z"},
		{"Zpurple90"},
		{"purple90Z"},
		{"Zpurple900"},
		{"purple900Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Purple900()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestPurpleA100Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"purpleA100", 10},
		{"purpleA100 ", 10},
		{"purpleA100\n", 10},
		{"purpleA100.", 10},
		{"purpleA100:", 10},
		{"purpleA100,", 10},
		{"purpleA100\"", 10},
		{"purpleA100(", 10},
		{"purpleA100)", 10},
		{"purpleA100[", 10},
		{"purpleA100]", 10},
		{"purpleA100// comment", 10},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.PurpleA100()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestPurpleA100Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_p"},
		{"p_"},
		{"_pu"},
		{"pu_"},
		{"_pur"},
		{"pur_"},
		{"_purp"},
		{"purp_"},
		{"_purpl"},
		{"purpl_"},
		{"_purple"},
		{"purple_"},
		{"_purpleA"},
		{"purpleA_"},
		{"_purpleA1"},
		{"purpleA1_"},
		{"_purpleA10"},
		{"purpleA10_"},
		{"_purpleA100"},
		{"purpleA100_"},
		{"9713p"},
		{"p9713"},
		{"9713pu"},
		{"pu9713"},
		{"9713pur"},
		{"pur9713"},
		{"9713purp"},
		{"purp9713"},
		{"9713purpl"},
		{"purpl9713"},
		{"9713purple"},
		{"purple9713"},
		{"9713purpleA"},
		{"purpleA9713"},
		{"9713purpleA1"},
		{"purpleA19713"},
		{"9713purpleA10"},
		{"purpleA109713"},
		{"9713purpleA100"},
		{"purpleA1009713"},
		{"Zp"},
		{"pZ"},
		{"Zpu"},
		{"puZ"},
		{"Zpur"},
		{"purZ"},
		{"Zpurp"},
		{"purpZ"},
		{"Zpurpl"},
		{"purplZ"},
		{"Zpurple"},
		{"purpleZ"},
		{"ZpurpleA"},
		{"purpleAZ"},
		{"ZpurpleA1"},
		{"purpleA1Z"},
		{"ZpurpleA10"},
		{"purpleA10Z"},
		{"ZpurpleA100"},
		{"purpleA100Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.PurpleA100()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestPurpleA200Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"purpleA200", 10},
		{"purpleA200 ", 10},
		{"purpleA200\n", 10},
		{"purpleA200.", 10},
		{"purpleA200:", 10},
		{"purpleA200,", 10},
		{"purpleA200\"", 10},
		{"purpleA200(", 10},
		{"purpleA200)", 10},
		{"purpleA200[", 10},
		{"purpleA200]", 10},
		{"purpleA200// comment", 10},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.PurpleA200()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestPurpleA200Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_p"},
		{"p_"},
		{"_pu"},
		{"pu_"},
		{"_pur"},
		{"pur_"},
		{"_purp"},
		{"purp_"},
		{"_purpl"},
		{"purpl_"},
		{"_purple"},
		{"purple_"},
		{"_purpleA"},
		{"purpleA_"},
		{"_purpleA2"},
		{"purpleA2_"},
		{"_purpleA20"},
		{"purpleA20_"},
		{"_purpleA200"},
		{"purpleA200_"},
		{"9713p"},
		{"p9713"},
		{"9713pu"},
		{"pu9713"},
		{"9713pur"},
		{"pur9713"},
		{"9713purp"},
		{"purp9713"},
		{"9713purpl"},
		{"purpl9713"},
		{"9713purple"},
		{"purple9713"},
		{"9713purpleA"},
		{"purpleA9713"},
		{"9713purpleA2"},
		{"purpleA29713"},
		{"9713purpleA20"},
		{"purpleA209713"},
		{"9713purpleA200"},
		{"purpleA2009713"},
		{"Zp"},
		{"pZ"},
		{"Zpu"},
		{"puZ"},
		{"Zpur"},
		{"purZ"},
		{"Zpurp"},
		{"purpZ"},
		{"Zpurpl"},
		{"purplZ"},
		{"Zpurple"},
		{"purpleZ"},
		{"ZpurpleA"},
		{"purpleAZ"},
		{"ZpurpleA2"},
		{"purpleA2Z"},
		{"ZpurpleA20"},
		{"purpleA20Z"},
		{"ZpurpleA200"},
		{"purpleA200Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.PurpleA200()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestPurpleA400Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"purpleA400", 10},
		{"purpleA400 ", 10},
		{"purpleA400\n", 10},
		{"purpleA400.", 10},
		{"purpleA400:", 10},
		{"purpleA400,", 10},
		{"purpleA400\"", 10},
		{"purpleA400(", 10},
		{"purpleA400)", 10},
		{"purpleA400[", 10},
		{"purpleA400]", 10},
		{"purpleA400// comment", 10},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.PurpleA400()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestPurpleA400Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_p"},
		{"p_"},
		{"_pu"},
		{"pu_"},
		{"_pur"},
		{"pur_"},
		{"_purp"},
		{"purp_"},
		{"_purpl"},
		{"purpl_"},
		{"_purple"},
		{"purple_"},
		{"_purpleA"},
		{"purpleA_"},
		{"_purpleA4"},
		{"purpleA4_"},
		{"_purpleA40"},
		{"purpleA40_"},
		{"_purpleA400"},
		{"purpleA400_"},
		{"9713p"},
		{"p9713"},
		{"9713pu"},
		{"pu9713"},
		{"9713pur"},
		{"pur9713"},
		{"9713purp"},
		{"purp9713"},
		{"9713purpl"},
		{"purpl9713"},
		{"9713purple"},
		{"purple9713"},
		{"9713purpleA"},
		{"purpleA9713"},
		{"9713purpleA4"},
		{"purpleA49713"},
		{"9713purpleA40"},
		{"purpleA409713"},
		{"9713purpleA400"},
		{"purpleA4009713"},
		{"Zp"},
		{"pZ"},
		{"Zpu"},
		{"puZ"},
		{"Zpur"},
		{"purZ"},
		{"Zpurp"},
		{"purpZ"},
		{"Zpurpl"},
		{"purplZ"},
		{"Zpurple"},
		{"purpleZ"},
		{"ZpurpleA"},
		{"purpleAZ"},
		{"ZpurpleA4"},
		{"purpleA4Z"},
		{"ZpurpleA40"},
		{"purpleA40Z"},
		{"ZpurpleA400"},
		{"purpleA400Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.PurpleA400()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestPurpleA700Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"purpleA700", 10},
		{"purpleA700 ", 10},
		{"purpleA700\n", 10},
		{"purpleA700.", 10},
		{"purpleA700:", 10},
		{"purpleA700,", 10},
		{"purpleA700\"", 10},
		{"purpleA700(", 10},
		{"purpleA700)", 10},
		{"purpleA700[", 10},
		{"purpleA700]", 10},
		{"purpleA700// comment", 10},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.PurpleA700()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestPurpleA700Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_p"},
		{"p_"},
		{"_pu"},
		{"pu_"},
		{"_pur"},
		{"pur_"},
		{"_purp"},
		{"purp_"},
		{"_purpl"},
		{"purpl_"},
		{"_purple"},
		{"purple_"},
		{"_purpleA"},
		{"purpleA_"},
		{"_purpleA7"},
		{"purpleA7_"},
		{"_purpleA70"},
		{"purpleA70_"},
		{"_purpleA700"},
		{"purpleA700_"},
		{"9713p"},
		{"p9713"},
		{"9713pu"},
		{"pu9713"},
		{"9713pur"},
		{"pur9713"},
		{"9713purp"},
		{"purp9713"},
		{"9713purpl"},
		{"purpl9713"},
		{"9713purple"},
		{"purple9713"},
		{"9713purpleA"},
		{"purpleA9713"},
		{"9713purpleA7"},
		{"purpleA79713"},
		{"9713purpleA70"},
		{"purpleA709713"},
		{"9713purpleA700"},
		{"purpleA7009713"},
		{"Zp"},
		{"pZ"},
		{"Zpu"},
		{"puZ"},
		{"Zpur"},
		{"purZ"},
		{"Zpurp"},
		{"purpZ"},
		{"Zpurpl"},
		{"purplZ"},
		{"Zpurple"},
		{"purpleZ"},
		{"ZpurpleA"},
		{"purpleAZ"},
		{"ZpurpleA7"},
		{"purpleA7Z"},
		{"ZpurpleA70"},
		{"purpleA70Z"},
		{"ZpurpleA700"},
		{"purpleA700Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.PurpleA700()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestDeepPurple50Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"deepPurple50", 12},
		{"deepPurple50 ", 12},
		{"deepPurple50\n", 12},
		{"deepPurple50.", 12},
		{"deepPurple50:", 12},
		{"deepPurple50,", 12},
		{"deepPurple50\"", 12},
		{"deepPurple50(", 12},
		{"deepPurple50)", 12},
		{"deepPurple50[", 12},
		{"deepPurple50]", 12},
		{"deepPurple50// comment", 12},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.DeepPurple50()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestDeepPurple50Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_d"},
		{"d_"},
		{"_de"},
		{"de_"},
		{"_dee"},
		{"dee_"},
		{"_deep"},
		{"deep_"},
		{"_deepP"},
		{"deepP_"},
		{"_deepPu"},
		{"deepPu_"},
		{"_deepPur"},
		{"deepPur_"},
		{"_deepPurp"},
		{"deepPurp_"},
		{"_deepPurpl"},
		{"deepPurpl_"},
		{"_deepPurple"},
		{"deepPurple_"},
		{"_deepPurple5"},
		{"deepPurple5_"},
		{"_deepPurple50"},
		{"deepPurple50_"},
		{"9713d"},
		{"d9713"},
		{"9713de"},
		{"de9713"},
		{"9713dee"},
		{"dee9713"},
		{"9713deep"},
		{"deep9713"},
		{"9713deepP"},
		{"deepP9713"},
		{"9713deepPu"},
		{"deepPu9713"},
		{"9713deepPur"},
		{"deepPur9713"},
		{"9713deepPurp"},
		{"deepPurp9713"},
		{"9713deepPurpl"},
		{"deepPurpl9713"},
		{"9713deepPurple"},
		{"deepPurple9713"},
		{"9713deepPurple5"},
		{"deepPurple59713"},
		{"9713deepPurple50"},
		{"deepPurple509713"},
		{"Zd"},
		{"dZ"},
		{"Zde"},
		{"deZ"},
		{"Zdee"},
		{"deeZ"},
		{"Zdeep"},
		{"deepZ"},
		{"ZdeepP"},
		{"deepPZ"},
		{"ZdeepPu"},
		{"deepPuZ"},
		{"ZdeepPur"},
		{"deepPurZ"},
		{"ZdeepPurp"},
		{"deepPurpZ"},
		{"ZdeepPurpl"},
		{"deepPurplZ"},
		{"ZdeepPurple"},
		{"deepPurpleZ"},
		{"ZdeepPurple5"},
		{"deepPurple5Z"},
		{"ZdeepPurple50"},
		{"deepPurple50Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.DeepPurple50()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestDeepPurple100Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"deepPurple100", 13},
		{"deepPurple100 ", 13},
		{"deepPurple100\n", 13},
		{"deepPurple100.", 13},
		{"deepPurple100:", 13},
		{"deepPurple100,", 13},
		{"deepPurple100\"", 13},
		{"deepPurple100(", 13},
		{"deepPurple100)", 13},
		{"deepPurple100[", 13},
		{"deepPurple100]", 13},
		{"deepPurple100// comment", 13},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.DeepPurple100()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestDeepPurple100Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_d"},
		{"d_"},
		{"_de"},
		{"de_"},
		{"_dee"},
		{"dee_"},
		{"_deep"},
		{"deep_"},
		{"_deepP"},
		{"deepP_"},
		{"_deepPu"},
		{"deepPu_"},
		{"_deepPur"},
		{"deepPur_"},
		{"_deepPurp"},
		{"deepPurp_"},
		{"_deepPurpl"},
		{"deepPurpl_"},
		{"_deepPurple"},
		{"deepPurple_"},
		{"_deepPurple1"},
		{"deepPurple1_"},
		{"_deepPurple10"},
		{"deepPurple10_"},
		{"_deepPurple100"},
		{"deepPurple100_"},
		{"9713d"},
		{"d9713"},
		{"9713de"},
		{"de9713"},
		{"9713dee"},
		{"dee9713"},
		{"9713deep"},
		{"deep9713"},
		{"9713deepP"},
		{"deepP9713"},
		{"9713deepPu"},
		{"deepPu9713"},
		{"9713deepPur"},
		{"deepPur9713"},
		{"9713deepPurp"},
		{"deepPurp9713"},
		{"9713deepPurpl"},
		{"deepPurpl9713"},
		{"9713deepPurple"},
		{"deepPurple9713"},
		{"9713deepPurple1"},
		{"deepPurple19713"},
		{"9713deepPurple10"},
		{"deepPurple109713"},
		{"9713deepPurple100"},
		{"deepPurple1009713"},
		{"Zd"},
		{"dZ"},
		{"Zde"},
		{"deZ"},
		{"Zdee"},
		{"deeZ"},
		{"Zdeep"},
		{"deepZ"},
		{"ZdeepP"},
		{"deepPZ"},
		{"ZdeepPu"},
		{"deepPuZ"},
		{"ZdeepPur"},
		{"deepPurZ"},
		{"ZdeepPurp"},
		{"deepPurpZ"},
		{"ZdeepPurpl"},
		{"deepPurplZ"},
		{"ZdeepPurple"},
		{"deepPurpleZ"},
		{"ZdeepPurple1"},
		{"deepPurple1Z"},
		{"ZdeepPurple10"},
		{"deepPurple10Z"},
		{"ZdeepPurple100"},
		{"deepPurple100Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.DeepPurple100()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestDeepPurple200Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"deepPurple200", 13},
		{"deepPurple200 ", 13},
		{"deepPurple200\n", 13},
		{"deepPurple200.", 13},
		{"deepPurple200:", 13},
		{"deepPurple200,", 13},
		{"deepPurple200\"", 13},
		{"deepPurple200(", 13},
		{"deepPurple200)", 13},
		{"deepPurple200[", 13},
		{"deepPurple200]", 13},
		{"deepPurple200// comment", 13},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.DeepPurple200()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestDeepPurple200Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_d"},
		{"d_"},
		{"_de"},
		{"de_"},
		{"_dee"},
		{"dee_"},
		{"_deep"},
		{"deep_"},
		{"_deepP"},
		{"deepP_"},
		{"_deepPu"},
		{"deepPu_"},
		{"_deepPur"},
		{"deepPur_"},
		{"_deepPurp"},
		{"deepPurp_"},
		{"_deepPurpl"},
		{"deepPurpl_"},
		{"_deepPurple"},
		{"deepPurple_"},
		{"_deepPurple2"},
		{"deepPurple2_"},
		{"_deepPurple20"},
		{"deepPurple20_"},
		{"_deepPurple200"},
		{"deepPurple200_"},
		{"9713d"},
		{"d9713"},
		{"9713de"},
		{"de9713"},
		{"9713dee"},
		{"dee9713"},
		{"9713deep"},
		{"deep9713"},
		{"9713deepP"},
		{"deepP9713"},
		{"9713deepPu"},
		{"deepPu9713"},
		{"9713deepPur"},
		{"deepPur9713"},
		{"9713deepPurp"},
		{"deepPurp9713"},
		{"9713deepPurpl"},
		{"deepPurpl9713"},
		{"9713deepPurple"},
		{"deepPurple9713"},
		{"9713deepPurple2"},
		{"deepPurple29713"},
		{"9713deepPurple20"},
		{"deepPurple209713"},
		{"9713deepPurple200"},
		{"deepPurple2009713"},
		{"Zd"},
		{"dZ"},
		{"Zde"},
		{"deZ"},
		{"Zdee"},
		{"deeZ"},
		{"Zdeep"},
		{"deepZ"},
		{"ZdeepP"},
		{"deepPZ"},
		{"ZdeepPu"},
		{"deepPuZ"},
		{"ZdeepPur"},
		{"deepPurZ"},
		{"ZdeepPurp"},
		{"deepPurpZ"},
		{"ZdeepPurpl"},
		{"deepPurplZ"},
		{"ZdeepPurple"},
		{"deepPurpleZ"},
		{"ZdeepPurple2"},
		{"deepPurple2Z"},
		{"ZdeepPurple20"},
		{"deepPurple20Z"},
		{"ZdeepPurple200"},
		{"deepPurple200Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.DeepPurple200()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestDeepPurple300Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"deepPurple300", 13},
		{"deepPurple300 ", 13},
		{"deepPurple300\n", 13},
		{"deepPurple300.", 13},
		{"deepPurple300:", 13},
		{"deepPurple300,", 13},
		{"deepPurple300\"", 13},
		{"deepPurple300(", 13},
		{"deepPurple300)", 13},
		{"deepPurple300[", 13},
		{"deepPurple300]", 13},
		{"deepPurple300// comment", 13},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.DeepPurple300()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestDeepPurple300Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_d"},
		{"d_"},
		{"_de"},
		{"de_"},
		{"_dee"},
		{"dee_"},
		{"_deep"},
		{"deep_"},
		{"_deepP"},
		{"deepP_"},
		{"_deepPu"},
		{"deepPu_"},
		{"_deepPur"},
		{"deepPur_"},
		{"_deepPurp"},
		{"deepPurp_"},
		{"_deepPurpl"},
		{"deepPurpl_"},
		{"_deepPurple"},
		{"deepPurple_"},
		{"_deepPurple3"},
		{"deepPurple3_"},
		{"_deepPurple30"},
		{"deepPurple30_"},
		{"_deepPurple300"},
		{"deepPurple300_"},
		{"9713d"},
		{"d9713"},
		{"9713de"},
		{"de9713"},
		{"9713dee"},
		{"dee9713"},
		{"9713deep"},
		{"deep9713"},
		{"9713deepP"},
		{"deepP9713"},
		{"9713deepPu"},
		{"deepPu9713"},
		{"9713deepPur"},
		{"deepPur9713"},
		{"9713deepPurp"},
		{"deepPurp9713"},
		{"9713deepPurpl"},
		{"deepPurpl9713"},
		{"9713deepPurple"},
		{"deepPurple9713"},
		{"9713deepPurple3"},
		{"deepPurple39713"},
		{"9713deepPurple30"},
		{"deepPurple309713"},
		{"9713deepPurple300"},
		{"deepPurple3009713"},
		{"Zd"},
		{"dZ"},
		{"Zde"},
		{"deZ"},
		{"Zdee"},
		{"deeZ"},
		{"Zdeep"},
		{"deepZ"},
		{"ZdeepP"},
		{"deepPZ"},
		{"ZdeepPu"},
		{"deepPuZ"},
		{"ZdeepPur"},
		{"deepPurZ"},
		{"ZdeepPurp"},
		{"deepPurpZ"},
		{"ZdeepPurpl"},
		{"deepPurplZ"},
		{"ZdeepPurple"},
		{"deepPurpleZ"},
		{"ZdeepPurple3"},
		{"deepPurple3Z"},
		{"ZdeepPurple30"},
		{"deepPurple30Z"},
		{"ZdeepPurple300"},
		{"deepPurple300Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.DeepPurple300()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestDeepPurple400Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"deepPurple400", 13},
		{"deepPurple400 ", 13},
		{"deepPurple400\n", 13},
		{"deepPurple400.", 13},
		{"deepPurple400:", 13},
		{"deepPurple400,", 13},
		{"deepPurple400\"", 13},
		{"deepPurple400(", 13},
		{"deepPurple400)", 13},
		{"deepPurple400[", 13},
		{"deepPurple400]", 13},
		{"deepPurple400// comment", 13},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.DeepPurple400()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestDeepPurple400Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_d"},
		{"d_"},
		{"_de"},
		{"de_"},
		{"_dee"},
		{"dee_"},
		{"_deep"},
		{"deep_"},
		{"_deepP"},
		{"deepP_"},
		{"_deepPu"},
		{"deepPu_"},
		{"_deepPur"},
		{"deepPur_"},
		{"_deepPurp"},
		{"deepPurp_"},
		{"_deepPurpl"},
		{"deepPurpl_"},
		{"_deepPurple"},
		{"deepPurple_"},
		{"_deepPurple4"},
		{"deepPurple4_"},
		{"_deepPurple40"},
		{"deepPurple40_"},
		{"_deepPurple400"},
		{"deepPurple400_"},
		{"9713d"},
		{"d9713"},
		{"9713de"},
		{"de9713"},
		{"9713dee"},
		{"dee9713"},
		{"9713deep"},
		{"deep9713"},
		{"9713deepP"},
		{"deepP9713"},
		{"9713deepPu"},
		{"deepPu9713"},
		{"9713deepPur"},
		{"deepPur9713"},
		{"9713deepPurp"},
		{"deepPurp9713"},
		{"9713deepPurpl"},
		{"deepPurpl9713"},
		{"9713deepPurple"},
		{"deepPurple9713"},
		{"9713deepPurple4"},
		{"deepPurple49713"},
		{"9713deepPurple40"},
		{"deepPurple409713"},
		{"9713deepPurple400"},
		{"deepPurple4009713"},
		{"Zd"},
		{"dZ"},
		{"Zde"},
		{"deZ"},
		{"Zdee"},
		{"deeZ"},
		{"Zdeep"},
		{"deepZ"},
		{"ZdeepP"},
		{"deepPZ"},
		{"ZdeepPu"},
		{"deepPuZ"},
		{"ZdeepPur"},
		{"deepPurZ"},
		{"ZdeepPurp"},
		{"deepPurpZ"},
		{"ZdeepPurpl"},
		{"deepPurplZ"},
		{"ZdeepPurple"},
		{"deepPurpleZ"},
		{"ZdeepPurple4"},
		{"deepPurple4Z"},
		{"ZdeepPurple40"},
		{"deepPurple40Z"},
		{"ZdeepPurple400"},
		{"deepPurple400Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.DeepPurple400()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestDeepPurple500Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"deepPurple500", 13},
		{"deepPurple500 ", 13},
		{"deepPurple500\n", 13},
		{"deepPurple500.", 13},
		{"deepPurple500:", 13},
		{"deepPurple500,", 13},
		{"deepPurple500\"", 13},
		{"deepPurple500(", 13},
		{"deepPurple500)", 13},
		{"deepPurple500[", 13},
		{"deepPurple500]", 13},
		{"deepPurple500// comment", 13},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.DeepPurple500()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestDeepPurple500Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_d"},
		{"d_"},
		{"_de"},
		{"de_"},
		{"_dee"},
		{"dee_"},
		{"_deep"},
		{"deep_"},
		{"_deepP"},
		{"deepP_"},
		{"_deepPu"},
		{"deepPu_"},
		{"_deepPur"},
		{"deepPur_"},
		{"_deepPurp"},
		{"deepPurp_"},
		{"_deepPurpl"},
		{"deepPurpl_"},
		{"_deepPurple"},
		{"deepPurple_"},
		{"_deepPurple5"},
		{"deepPurple5_"},
		{"_deepPurple50"},
		{"deepPurple50_"},
		{"_deepPurple500"},
		{"deepPurple500_"},
		{"9713d"},
		{"d9713"},
		{"9713de"},
		{"de9713"},
		{"9713dee"},
		{"dee9713"},
		{"9713deep"},
		{"deep9713"},
		{"9713deepP"},
		{"deepP9713"},
		{"9713deepPu"},
		{"deepPu9713"},
		{"9713deepPur"},
		{"deepPur9713"},
		{"9713deepPurp"},
		{"deepPurp9713"},
		{"9713deepPurpl"},
		{"deepPurpl9713"},
		{"9713deepPurple"},
		{"deepPurple9713"},
		{"9713deepPurple5"},
		{"deepPurple59713"},
		{"9713deepPurple50"},
		{"deepPurple509713"},
		{"9713deepPurple500"},
		{"deepPurple5009713"},
		{"Zd"},
		{"dZ"},
		{"Zde"},
		{"deZ"},
		{"Zdee"},
		{"deeZ"},
		{"Zdeep"},
		{"deepZ"},
		{"ZdeepP"},
		{"deepPZ"},
		{"ZdeepPu"},
		{"deepPuZ"},
		{"ZdeepPur"},
		{"deepPurZ"},
		{"ZdeepPurp"},
		{"deepPurpZ"},
		{"ZdeepPurpl"},
		{"deepPurplZ"},
		{"ZdeepPurple"},
		{"deepPurpleZ"},
		{"ZdeepPurple5"},
		{"deepPurple5Z"},
		{"ZdeepPurple50"},
		{"deepPurple50Z"},
		{"ZdeepPurple500"},
		{"deepPurple500Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.DeepPurple500()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestDeepPurple600Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"deepPurple600", 13},
		{"deepPurple600 ", 13},
		{"deepPurple600\n", 13},
		{"deepPurple600.", 13},
		{"deepPurple600:", 13},
		{"deepPurple600,", 13},
		{"deepPurple600\"", 13},
		{"deepPurple600(", 13},
		{"deepPurple600)", 13},
		{"deepPurple600[", 13},
		{"deepPurple600]", 13},
		{"deepPurple600// comment", 13},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.DeepPurple600()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestDeepPurple600Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_d"},
		{"d_"},
		{"_de"},
		{"de_"},
		{"_dee"},
		{"dee_"},
		{"_deep"},
		{"deep_"},
		{"_deepP"},
		{"deepP_"},
		{"_deepPu"},
		{"deepPu_"},
		{"_deepPur"},
		{"deepPur_"},
		{"_deepPurp"},
		{"deepPurp_"},
		{"_deepPurpl"},
		{"deepPurpl_"},
		{"_deepPurple"},
		{"deepPurple_"},
		{"_deepPurple6"},
		{"deepPurple6_"},
		{"_deepPurple60"},
		{"deepPurple60_"},
		{"_deepPurple600"},
		{"deepPurple600_"},
		{"9713d"},
		{"d9713"},
		{"9713de"},
		{"de9713"},
		{"9713dee"},
		{"dee9713"},
		{"9713deep"},
		{"deep9713"},
		{"9713deepP"},
		{"deepP9713"},
		{"9713deepPu"},
		{"deepPu9713"},
		{"9713deepPur"},
		{"deepPur9713"},
		{"9713deepPurp"},
		{"deepPurp9713"},
		{"9713deepPurpl"},
		{"deepPurpl9713"},
		{"9713deepPurple"},
		{"deepPurple9713"},
		{"9713deepPurple6"},
		{"deepPurple69713"},
		{"9713deepPurple60"},
		{"deepPurple609713"},
		{"9713deepPurple600"},
		{"deepPurple6009713"},
		{"Zd"},
		{"dZ"},
		{"Zde"},
		{"deZ"},
		{"Zdee"},
		{"deeZ"},
		{"Zdeep"},
		{"deepZ"},
		{"ZdeepP"},
		{"deepPZ"},
		{"ZdeepPu"},
		{"deepPuZ"},
		{"ZdeepPur"},
		{"deepPurZ"},
		{"ZdeepPurp"},
		{"deepPurpZ"},
		{"ZdeepPurpl"},
		{"deepPurplZ"},
		{"ZdeepPurple"},
		{"deepPurpleZ"},
		{"ZdeepPurple6"},
		{"deepPurple6Z"},
		{"ZdeepPurple60"},
		{"deepPurple60Z"},
		{"ZdeepPurple600"},
		{"deepPurple600Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.DeepPurple600()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestDeepPurple700Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"deepPurple700", 13},
		{"deepPurple700 ", 13},
		{"deepPurple700\n", 13},
		{"deepPurple700.", 13},
		{"deepPurple700:", 13},
		{"deepPurple700,", 13},
		{"deepPurple700\"", 13},
		{"deepPurple700(", 13},
		{"deepPurple700)", 13},
		{"deepPurple700[", 13},
		{"deepPurple700]", 13},
		{"deepPurple700// comment", 13},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.DeepPurple700()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestDeepPurple700Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_d"},
		{"d_"},
		{"_de"},
		{"de_"},
		{"_dee"},
		{"dee_"},
		{"_deep"},
		{"deep_"},
		{"_deepP"},
		{"deepP_"},
		{"_deepPu"},
		{"deepPu_"},
		{"_deepPur"},
		{"deepPur_"},
		{"_deepPurp"},
		{"deepPurp_"},
		{"_deepPurpl"},
		{"deepPurpl_"},
		{"_deepPurple"},
		{"deepPurple_"},
		{"_deepPurple7"},
		{"deepPurple7_"},
		{"_deepPurple70"},
		{"deepPurple70_"},
		{"_deepPurple700"},
		{"deepPurple700_"},
		{"9713d"},
		{"d9713"},
		{"9713de"},
		{"de9713"},
		{"9713dee"},
		{"dee9713"},
		{"9713deep"},
		{"deep9713"},
		{"9713deepP"},
		{"deepP9713"},
		{"9713deepPu"},
		{"deepPu9713"},
		{"9713deepPur"},
		{"deepPur9713"},
		{"9713deepPurp"},
		{"deepPurp9713"},
		{"9713deepPurpl"},
		{"deepPurpl9713"},
		{"9713deepPurple"},
		{"deepPurple9713"},
		{"9713deepPurple7"},
		{"deepPurple79713"},
		{"9713deepPurple70"},
		{"deepPurple709713"},
		{"9713deepPurple700"},
		{"deepPurple7009713"},
		{"Zd"},
		{"dZ"},
		{"Zde"},
		{"deZ"},
		{"Zdee"},
		{"deeZ"},
		{"Zdeep"},
		{"deepZ"},
		{"ZdeepP"},
		{"deepPZ"},
		{"ZdeepPu"},
		{"deepPuZ"},
		{"ZdeepPur"},
		{"deepPurZ"},
		{"ZdeepPurp"},
		{"deepPurpZ"},
		{"ZdeepPurpl"},
		{"deepPurplZ"},
		{"ZdeepPurple"},
		{"deepPurpleZ"},
		{"ZdeepPurple7"},
		{"deepPurple7Z"},
		{"ZdeepPurple70"},
		{"deepPurple70Z"},
		{"ZdeepPurple700"},
		{"deepPurple700Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.DeepPurple700()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestDeepPurple800Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"deepPurple800", 13},
		{"deepPurple800 ", 13},
		{"deepPurple800\n", 13},
		{"deepPurple800.", 13},
		{"deepPurple800:", 13},
		{"deepPurple800,", 13},
		{"deepPurple800\"", 13},
		{"deepPurple800(", 13},
		{"deepPurple800)", 13},
		{"deepPurple800[", 13},
		{"deepPurple800]", 13},
		{"deepPurple800// comment", 13},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.DeepPurple800()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestDeepPurple800Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_d"},
		{"d_"},
		{"_de"},
		{"de_"},
		{"_dee"},
		{"dee_"},
		{"_deep"},
		{"deep_"},
		{"_deepP"},
		{"deepP_"},
		{"_deepPu"},
		{"deepPu_"},
		{"_deepPur"},
		{"deepPur_"},
		{"_deepPurp"},
		{"deepPurp_"},
		{"_deepPurpl"},
		{"deepPurpl_"},
		{"_deepPurple"},
		{"deepPurple_"},
		{"_deepPurple8"},
		{"deepPurple8_"},
		{"_deepPurple80"},
		{"deepPurple80_"},
		{"_deepPurple800"},
		{"deepPurple800_"},
		{"9713d"},
		{"d9713"},
		{"9713de"},
		{"de9713"},
		{"9713dee"},
		{"dee9713"},
		{"9713deep"},
		{"deep9713"},
		{"9713deepP"},
		{"deepP9713"},
		{"9713deepPu"},
		{"deepPu9713"},
		{"9713deepPur"},
		{"deepPur9713"},
		{"9713deepPurp"},
		{"deepPurp9713"},
		{"9713deepPurpl"},
		{"deepPurpl9713"},
		{"9713deepPurple"},
		{"deepPurple9713"},
		{"9713deepPurple8"},
		{"deepPurple89713"},
		{"9713deepPurple80"},
		{"deepPurple809713"},
		{"9713deepPurple800"},
		{"deepPurple8009713"},
		{"Zd"},
		{"dZ"},
		{"Zde"},
		{"deZ"},
		{"Zdee"},
		{"deeZ"},
		{"Zdeep"},
		{"deepZ"},
		{"ZdeepP"},
		{"deepPZ"},
		{"ZdeepPu"},
		{"deepPuZ"},
		{"ZdeepPur"},
		{"deepPurZ"},
		{"ZdeepPurp"},
		{"deepPurpZ"},
		{"ZdeepPurpl"},
		{"deepPurplZ"},
		{"ZdeepPurple"},
		{"deepPurpleZ"},
		{"ZdeepPurple8"},
		{"deepPurple8Z"},
		{"ZdeepPurple80"},
		{"deepPurple80Z"},
		{"ZdeepPurple800"},
		{"deepPurple800Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.DeepPurple800()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestDeepPurple900Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"deepPurple900", 13},
		{"deepPurple900 ", 13},
		{"deepPurple900\n", 13},
		{"deepPurple900.", 13},
		{"deepPurple900:", 13},
		{"deepPurple900,", 13},
		{"deepPurple900\"", 13},
		{"deepPurple900(", 13},
		{"deepPurple900)", 13},
		{"deepPurple900[", 13},
		{"deepPurple900]", 13},
		{"deepPurple900// comment", 13},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.DeepPurple900()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestDeepPurple900Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_d"},
		{"d_"},
		{"_de"},
		{"de_"},
		{"_dee"},
		{"dee_"},
		{"_deep"},
		{"deep_"},
		{"_deepP"},
		{"deepP_"},
		{"_deepPu"},
		{"deepPu_"},
		{"_deepPur"},
		{"deepPur_"},
		{"_deepPurp"},
		{"deepPurp_"},
		{"_deepPurpl"},
		{"deepPurpl_"},
		{"_deepPurple"},
		{"deepPurple_"},
		{"_deepPurple9"},
		{"deepPurple9_"},
		{"_deepPurple90"},
		{"deepPurple90_"},
		{"_deepPurple900"},
		{"deepPurple900_"},
		{"9713d"},
		{"d9713"},
		{"9713de"},
		{"de9713"},
		{"9713dee"},
		{"dee9713"},
		{"9713deep"},
		{"deep9713"},
		{"9713deepP"},
		{"deepP9713"},
		{"9713deepPu"},
		{"deepPu9713"},
		{"9713deepPur"},
		{"deepPur9713"},
		{"9713deepPurp"},
		{"deepPurp9713"},
		{"9713deepPurpl"},
		{"deepPurpl9713"},
		{"9713deepPurple"},
		{"deepPurple9713"},
		{"9713deepPurple9"},
		{"deepPurple99713"},
		{"9713deepPurple90"},
		{"deepPurple909713"},
		{"9713deepPurple900"},
		{"deepPurple9009713"},
		{"Zd"},
		{"dZ"},
		{"Zde"},
		{"deZ"},
		{"Zdee"},
		{"deeZ"},
		{"Zdeep"},
		{"deepZ"},
		{"ZdeepP"},
		{"deepPZ"},
		{"ZdeepPu"},
		{"deepPuZ"},
		{"ZdeepPur"},
		{"deepPurZ"},
		{"ZdeepPurp"},
		{"deepPurpZ"},
		{"ZdeepPurpl"},
		{"deepPurplZ"},
		{"ZdeepPurple"},
		{"deepPurpleZ"},
		{"ZdeepPurple9"},
		{"deepPurple9Z"},
		{"ZdeepPurple90"},
		{"deepPurple90Z"},
		{"ZdeepPurple900"},
		{"deepPurple900Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.DeepPurple900()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestDeepPurpleA100Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"deepPurpleA100", 14},
		{"deepPurpleA100 ", 14},
		{"deepPurpleA100\n", 14},
		{"deepPurpleA100.", 14},
		{"deepPurpleA100:", 14},
		{"deepPurpleA100,", 14},
		{"deepPurpleA100\"", 14},
		{"deepPurpleA100(", 14},
		{"deepPurpleA100)", 14},
		{"deepPurpleA100[", 14},
		{"deepPurpleA100]", 14},
		{"deepPurpleA100// comment", 14},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.DeepPurpleA100()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestDeepPurpleA100Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_d"},
		{"d_"},
		{"_de"},
		{"de_"},
		{"_dee"},
		{"dee_"},
		{"_deep"},
		{"deep_"},
		{"_deepP"},
		{"deepP_"},
		{"_deepPu"},
		{"deepPu_"},
		{"_deepPur"},
		{"deepPur_"},
		{"_deepPurp"},
		{"deepPurp_"},
		{"_deepPurpl"},
		{"deepPurpl_"},
		{"_deepPurple"},
		{"deepPurple_"},
		{"_deepPurpleA"},
		{"deepPurpleA_"},
		{"_deepPurpleA1"},
		{"deepPurpleA1_"},
		{"_deepPurpleA10"},
		{"deepPurpleA10_"},
		{"_deepPurpleA100"},
		{"deepPurpleA100_"},
		{"9713d"},
		{"d9713"},
		{"9713de"},
		{"de9713"},
		{"9713dee"},
		{"dee9713"},
		{"9713deep"},
		{"deep9713"},
		{"9713deepP"},
		{"deepP9713"},
		{"9713deepPu"},
		{"deepPu9713"},
		{"9713deepPur"},
		{"deepPur9713"},
		{"9713deepPurp"},
		{"deepPurp9713"},
		{"9713deepPurpl"},
		{"deepPurpl9713"},
		{"9713deepPurple"},
		{"deepPurple9713"},
		{"9713deepPurpleA"},
		{"deepPurpleA9713"},
		{"9713deepPurpleA1"},
		{"deepPurpleA19713"},
		{"9713deepPurpleA10"},
		{"deepPurpleA109713"},
		{"9713deepPurpleA100"},
		{"deepPurpleA1009713"},
		{"Zd"},
		{"dZ"},
		{"Zde"},
		{"deZ"},
		{"Zdee"},
		{"deeZ"},
		{"Zdeep"},
		{"deepZ"},
		{"ZdeepP"},
		{"deepPZ"},
		{"ZdeepPu"},
		{"deepPuZ"},
		{"ZdeepPur"},
		{"deepPurZ"},
		{"ZdeepPurp"},
		{"deepPurpZ"},
		{"ZdeepPurpl"},
		{"deepPurplZ"},
		{"ZdeepPurple"},
		{"deepPurpleZ"},
		{"ZdeepPurpleA"},
		{"deepPurpleAZ"},
		{"ZdeepPurpleA1"},
		{"deepPurpleA1Z"},
		{"ZdeepPurpleA10"},
		{"deepPurpleA10Z"},
		{"ZdeepPurpleA100"},
		{"deepPurpleA100Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.DeepPurpleA100()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestDeepPurpleA200Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"deepPurpleA200", 14},
		{"deepPurpleA200 ", 14},
		{"deepPurpleA200\n", 14},
		{"deepPurpleA200.", 14},
		{"deepPurpleA200:", 14},
		{"deepPurpleA200,", 14},
		{"deepPurpleA200\"", 14},
		{"deepPurpleA200(", 14},
		{"deepPurpleA200)", 14},
		{"deepPurpleA200[", 14},
		{"deepPurpleA200]", 14},
		{"deepPurpleA200// comment", 14},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.DeepPurpleA200()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestDeepPurpleA200Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_d"},
		{"d_"},
		{"_de"},
		{"de_"},
		{"_dee"},
		{"dee_"},
		{"_deep"},
		{"deep_"},
		{"_deepP"},
		{"deepP_"},
		{"_deepPu"},
		{"deepPu_"},
		{"_deepPur"},
		{"deepPur_"},
		{"_deepPurp"},
		{"deepPurp_"},
		{"_deepPurpl"},
		{"deepPurpl_"},
		{"_deepPurple"},
		{"deepPurple_"},
		{"_deepPurpleA"},
		{"deepPurpleA_"},
		{"_deepPurpleA2"},
		{"deepPurpleA2_"},
		{"_deepPurpleA20"},
		{"deepPurpleA20_"},
		{"_deepPurpleA200"},
		{"deepPurpleA200_"},
		{"9713d"},
		{"d9713"},
		{"9713de"},
		{"de9713"},
		{"9713dee"},
		{"dee9713"},
		{"9713deep"},
		{"deep9713"},
		{"9713deepP"},
		{"deepP9713"},
		{"9713deepPu"},
		{"deepPu9713"},
		{"9713deepPur"},
		{"deepPur9713"},
		{"9713deepPurp"},
		{"deepPurp9713"},
		{"9713deepPurpl"},
		{"deepPurpl9713"},
		{"9713deepPurple"},
		{"deepPurple9713"},
		{"9713deepPurpleA"},
		{"deepPurpleA9713"},
		{"9713deepPurpleA2"},
		{"deepPurpleA29713"},
		{"9713deepPurpleA20"},
		{"deepPurpleA209713"},
		{"9713deepPurpleA200"},
		{"deepPurpleA2009713"},
		{"Zd"},
		{"dZ"},
		{"Zde"},
		{"deZ"},
		{"Zdee"},
		{"deeZ"},
		{"Zdeep"},
		{"deepZ"},
		{"ZdeepP"},
		{"deepPZ"},
		{"ZdeepPu"},
		{"deepPuZ"},
		{"ZdeepPur"},
		{"deepPurZ"},
		{"ZdeepPurp"},
		{"deepPurpZ"},
		{"ZdeepPurpl"},
		{"deepPurplZ"},
		{"ZdeepPurple"},
		{"deepPurpleZ"},
		{"ZdeepPurpleA"},
		{"deepPurpleAZ"},
		{"ZdeepPurpleA2"},
		{"deepPurpleA2Z"},
		{"ZdeepPurpleA20"},
		{"deepPurpleA20Z"},
		{"ZdeepPurpleA200"},
		{"deepPurpleA200Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.DeepPurpleA200()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestDeepPurpleA400Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"deepPurpleA400", 14},
		{"deepPurpleA400 ", 14},
		{"deepPurpleA400\n", 14},
		{"deepPurpleA400.", 14},
		{"deepPurpleA400:", 14},
		{"deepPurpleA400,", 14},
		{"deepPurpleA400\"", 14},
		{"deepPurpleA400(", 14},
		{"deepPurpleA400)", 14},
		{"deepPurpleA400[", 14},
		{"deepPurpleA400]", 14},
		{"deepPurpleA400// comment", 14},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.DeepPurpleA400()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestDeepPurpleA400Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_d"},
		{"d_"},
		{"_de"},
		{"de_"},
		{"_dee"},
		{"dee_"},
		{"_deep"},
		{"deep_"},
		{"_deepP"},
		{"deepP_"},
		{"_deepPu"},
		{"deepPu_"},
		{"_deepPur"},
		{"deepPur_"},
		{"_deepPurp"},
		{"deepPurp_"},
		{"_deepPurpl"},
		{"deepPurpl_"},
		{"_deepPurple"},
		{"deepPurple_"},
		{"_deepPurpleA"},
		{"deepPurpleA_"},
		{"_deepPurpleA4"},
		{"deepPurpleA4_"},
		{"_deepPurpleA40"},
		{"deepPurpleA40_"},
		{"_deepPurpleA400"},
		{"deepPurpleA400_"},
		{"9713d"},
		{"d9713"},
		{"9713de"},
		{"de9713"},
		{"9713dee"},
		{"dee9713"},
		{"9713deep"},
		{"deep9713"},
		{"9713deepP"},
		{"deepP9713"},
		{"9713deepPu"},
		{"deepPu9713"},
		{"9713deepPur"},
		{"deepPur9713"},
		{"9713deepPurp"},
		{"deepPurp9713"},
		{"9713deepPurpl"},
		{"deepPurpl9713"},
		{"9713deepPurple"},
		{"deepPurple9713"},
		{"9713deepPurpleA"},
		{"deepPurpleA9713"},
		{"9713deepPurpleA4"},
		{"deepPurpleA49713"},
		{"9713deepPurpleA40"},
		{"deepPurpleA409713"},
		{"9713deepPurpleA400"},
		{"deepPurpleA4009713"},
		{"Zd"},
		{"dZ"},
		{"Zde"},
		{"deZ"},
		{"Zdee"},
		{"deeZ"},
		{"Zdeep"},
		{"deepZ"},
		{"ZdeepP"},
		{"deepPZ"},
		{"ZdeepPu"},
		{"deepPuZ"},
		{"ZdeepPur"},
		{"deepPurZ"},
		{"ZdeepPurp"},
		{"deepPurpZ"},
		{"ZdeepPurpl"},
		{"deepPurplZ"},
		{"ZdeepPurple"},
		{"deepPurpleZ"},
		{"ZdeepPurpleA"},
		{"deepPurpleAZ"},
		{"ZdeepPurpleA4"},
		{"deepPurpleA4Z"},
		{"ZdeepPurpleA40"},
		{"deepPurpleA40Z"},
		{"ZdeepPurpleA400"},
		{"deepPurpleA400Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.DeepPurpleA400()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestDeepPurpleA700Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"deepPurpleA700", 14},
		{"deepPurpleA700 ", 14},
		{"deepPurpleA700\n", 14},
		{"deepPurpleA700.", 14},
		{"deepPurpleA700:", 14},
		{"deepPurpleA700,", 14},
		{"deepPurpleA700\"", 14},
		{"deepPurpleA700(", 14},
		{"deepPurpleA700)", 14},
		{"deepPurpleA700[", 14},
		{"deepPurpleA700]", 14},
		{"deepPurpleA700// comment", 14},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.DeepPurpleA700()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestDeepPurpleA700Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_d"},
		{"d_"},
		{"_de"},
		{"de_"},
		{"_dee"},
		{"dee_"},
		{"_deep"},
		{"deep_"},
		{"_deepP"},
		{"deepP_"},
		{"_deepPu"},
		{"deepPu_"},
		{"_deepPur"},
		{"deepPur_"},
		{"_deepPurp"},
		{"deepPurp_"},
		{"_deepPurpl"},
		{"deepPurpl_"},
		{"_deepPurple"},
		{"deepPurple_"},
		{"_deepPurpleA"},
		{"deepPurpleA_"},
		{"_deepPurpleA7"},
		{"deepPurpleA7_"},
		{"_deepPurpleA70"},
		{"deepPurpleA70_"},
		{"_deepPurpleA700"},
		{"deepPurpleA700_"},
		{"9713d"},
		{"d9713"},
		{"9713de"},
		{"de9713"},
		{"9713dee"},
		{"dee9713"},
		{"9713deep"},
		{"deep9713"},
		{"9713deepP"},
		{"deepP9713"},
		{"9713deepPu"},
		{"deepPu9713"},
		{"9713deepPur"},
		{"deepPur9713"},
		{"9713deepPurp"},
		{"deepPurp9713"},
		{"9713deepPurpl"},
		{"deepPurpl9713"},
		{"9713deepPurple"},
		{"deepPurple9713"},
		{"9713deepPurpleA"},
		{"deepPurpleA9713"},
		{"9713deepPurpleA7"},
		{"deepPurpleA79713"},
		{"9713deepPurpleA70"},
		{"deepPurpleA709713"},
		{"9713deepPurpleA700"},
		{"deepPurpleA7009713"},
		{"Zd"},
		{"dZ"},
		{"Zde"},
		{"deZ"},
		{"Zdee"},
		{"deeZ"},
		{"Zdeep"},
		{"deepZ"},
		{"ZdeepP"},
		{"deepPZ"},
		{"ZdeepPu"},
		{"deepPuZ"},
		{"ZdeepPur"},
		{"deepPurZ"},
		{"ZdeepPurp"},
		{"deepPurpZ"},
		{"ZdeepPurpl"},
		{"deepPurplZ"},
		{"ZdeepPurple"},
		{"deepPurpleZ"},
		{"ZdeepPurpleA"},
		{"deepPurpleAZ"},
		{"ZdeepPurpleA7"},
		{"deepPurpleA7Z"},
		{"ZdeepPurpleA70"},
		{"deepPurpleA70Z"},
		{"ZdeepPurpleA700"},
		{"deepPurpleA700Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.DeepPurpleA700()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestIndigo50Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"indigo50", 8},
		{"indigo50 ", 8},
		{"indigo50\n", 8},
		{"indigo50.", 8},
		{"indigo50:", 8},
		{"indigo50,", 8},
		{"indigo50\"", 8},
		{"indigo50(", 8},
		{"indigo50)", 8},
		{"indigo50[", 8},
		{"indigo50]", 8},
		{"indigo50// comment", 8},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Indigo50()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestIndigo50Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_i"},
		{"i_"},
		{"_in"},
		{"in_"},
		{"_ind"},
		{"ind_"},
		{"_indi"},
		{"indi_"},
		{"_indig"},
		{"indig_"},
		{"_indigo"},
		{"indigo_"},
		{"_indigo5"},
		{"indigo5_"},
		{"_indigo50"},
		{"indigo50_"},
		{"9713i"},
		{"i9713"},
		{"9713in"},
		{"in9713"},
		{"9713ind"},
		{"ind9713"},
		{"9713indi"},
		{"indi9713"},
		{"9713indig"},
		{"indig9713"},
		{"9713indigo"},
		{"indigo9713"},
		{"9713indigo5"},
		{"indigo59713"},
		{"9713indigo50"},
		{"indigo509713"},
		{"Zi"},
		{"iZ"},
		{"Zin"},
		{"inZ"},
		{"Zind"},
		{"indZ"},
		{"Zindi"},
		{"indiZ"},
		{"Zindig"},
		{"indigZ"},
		{"Zindigo"},
		{"indigoZ"},
		{"Zindigo5"},
		{"indigo5Z"},
		{"Zindigo50"},
		{"indigo50Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Indigo50()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestIndigo100Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"indigo100", 9},
		{"indigo100 ", 9},
		{"indigo100\n", 9},
		{"indigo100.", 9},
		{"indigo100:", 9},
		{"indigo100,", 9},
		{"indigo100\"", 9},
		{"indigo100(", 9},
		{"indigo100)", 9},
		{"indigo100[", 9},
		{"indigo100]", 9},
		{"indigo100// comment", 9},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Indigo100()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestIndigo100Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_i"},
		{"i_"},
		{"_in"},
		{"in_"},
		{"_ind"},
		{"ind_"},
		{"_indi"},
		{"indi_"},
		{"_indig"},
		{"indig_"},
		{"_indigo"},
		{"indigo_"},
		{"_indigo1"},
		{"indigo1_"},
		{"_indigo10"},
		{"indigo10_"},
		{"_indigo100"},
		{"indigo100_"},
		{"9713i"},
		{"i9713"},
		{"9713in"},
		{"in9713"},
		{"9713ind"},
		{"ind9713"},
		{"9713indi"},
		{"indi9713"},
		{"9713indig"},
		{"indig9713"},
		{"9713indigo"},
		{"indigo9713"},
		{"9713indigo1"},
		{"indigo19713"},
		{"9713indigo10"},
		{"indigo109713"},
		{"9713indigo100"},
		{"indigo1009713"},
		{"Zi"},
		{"iZ"},
		{"Zin"},
		{"inZ"},
		{"Zind"},
		{"indZ"},
		{"Zindi"},
		{"indiZ"},
		{"Zindig"},
		{"indigZ"},
		{"Zindigo"},
		{"indigoZ"},
		{"Zindigo1"},
		{"indigo1Z"},
		{"Zindigo10"},
		{"indigo10Z"},
		{"Zindigo100"},
		{"indigo100Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Indigo100()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestIndigo200Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"indigo200", 9},
		{"indigo200 ", 9},
		{"indigo200\n", 9},
		{"indigo200.", 9},
		{"indigo200:", 9},
		{"indigo200,", 9},
		{"indigo200\"", 9},
		{"indigo200(", 9},
		{"indigo200)", 9},
		{"indigo200[", 9},
		{"indigo200]", 9},
		{"indigo200// comment", 9},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Indigo200()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestIndigo200Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_i"},
		{"i_"},
		{"_in"},
		{"in_"},
		{"_ind"},
		{"ind_"},
		{"_indi"},
		{"indi_"},
		{"_indig"},
		{"indig_"},
		{"_indigo"},
		{"indigo_"},
		{"_indigo2"},
		{"indigo2_"},
		{"_indigo20"},
		{"indigo20_"},
		{"_indigo200"},
		{"indigo200_"},
		{"9713i"},
		{"i9713"},
		{"9713in"},
		{"in9713"},
		{"9713ind"},
		{"ind9713"},
		{"9713indi"},
		{"indi9713"},
		{"9713indig"},
		{"indig9713"},
		{"9713indigo"},
		{"indigo9713"},
		{"9713indigo2"},
		{"indigo29713"},
		{"9713indigo20"},
		{"indigo209713"},
		{"9713indigo200"},
		{"indigo2009713"},
		{"Zi"},
		{"iZ"},
		{"Zin"},
		{"inZ"},
		{"Zind"},
		{"indZ"},
		{"Zindi"},
		{"indiZ"},
		{"Zindig"},
		{"indigZ"},
		{"Zindigo"},
		{"indigoZ"},
		{"Zindigo2"},
		{"indigo2Z"},
		{"Zindigo20"},
		{"indigo20Z"},
		{"Zindigo200"},
		{"indigo200Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Indigo200()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestIndigo300Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"indigo300", 9},
		{"indigo300 ", 9},
		{"indigo300\n", 9},
		{"indigo300.", 9},
		{"indigo300:", 9},
		{"indigo300,", 9},
		{"indigo300\"", 9},
		{"indigo300(", 9},
		{"indigo300)", 9},
		{"indigo300[", 9},
		{"indigo300]", 9},
		{"indigo300// comment", 9},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Indigo300()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestIndigo300Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_i"},
		{"i_"},
		{"_in"},
		{"in_"},
		{"_ind"},
		{"ind_"},
		{"_indi"},
		{"indi_"},
		{"_indig"},
		{"indig_"},
		{"_indigo"},
		{"indigo_"},
		{"_indigo3"},
		{"indigo3_"},
		{"_indigo30"},
		{"indigo30_"},
		{"_indigo300"},
		{"indigo300_"},
		{"9713i"},
		{"i9713"},
		{"9713in"},
		{"in9713"},
		{"9713ind"},
		{"ind9713"},
		{"9713indi"},
		{"indi9713"},
		{"9713indig"},
		{"indig9713"},
		{"9713indigo"},
		{"indigo9713"},
		{"9713indigo3"},
		{"indigo39713"},
		{"9713indigo30"},
		{"indigo309713"},
		{"9713indigo300"},
		{"indigo3009713"},
		{"Zi"},
		{"iZ"},
		{"Zin"},
		{"inZ"},
		{"Zind"},
		{"indZ"},
		{"Zindi"},
		{"indiZ"},
		{"Zindig"},
		{"indigZ"},
		{"Zindigo"},
		{"indigoZ"},
		{"Zindigo3"},
		{"indigo3Z"},
		{"Zindigo30"},
		{"indigo30Z"},
		{"Zindigo300"},
		{"indigo300Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Indigo300()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestIndigo400Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"indigo400", 9},
		{"indigo400 ", 9},
		{"indigo400\n", 9},
		{"indigo400.", 9},
		{"indigo400:", 9},
		{"indigo400,", 9},
		{"indigo400\"", 9},
		{"indigo400(", 9},
		{"indigo400)", 9},
		{"indigo400[", 9},
		{"indigo400]", 9},
		{"indigo400// comment", 9},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Indigo400()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestIndigo400Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_i"},
		{"i_"},
		{"_in"},
		{"in_"},
		{"_ind"},
		{"ind_"},
		{"_indi"},
		{"indi_"},
		{"_indig"},
		{"indig_"},
		{"_indigo"},
		{"indigo_"},
		{"_indigo4"},
		{"indigo4_"},
		{"_indigo40"},
		{"indigo40_"},
		{"_indigo400"},
		{"indigo400_"},
		{"9713i"},
		{"i9713"},
		{"9713in"},
		{"in9713"},
		{"9713ind"},
		{"ind9713"},
		{"9713indi"},
		{"indi9713"},
		{"9713indig"},
		{"indig9713"},
		{"9713indigo"},
		{"indigo9713"},
		{"9713indigo4"},
		{"indigo49713"},
		{"9713indigo40"},
		{"indigo409713"},
		{"9713indigo400"},
		{"indigo4009713"},
		{"Zi"},
		{"iZ"},
		{"Zin"},
		{"inZ"},
		{"Zind"},
		{"indZ"},
		{"Zindi"},
		{"indiZ"},
		{"Zindig"},
		{"indigZ"},
		{"Zindigo"},
		{"indigoZ"},
		{"Zindigo4"},
		{"indigo4Z"},
		{"Zindigo40"},
		{"indigo40Z"},
		{"Zindigo400"},
		{"indigo400Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Indigo400()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestIndigo500Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"indigo500", 9},
		{"indigo500 ", 9},
		{"indigo500\n", 9},
		{"indigo500.", 9},
		{"indigo500:", 9},
		{"indigo500,", 9},
		{"indigo500\"", 9},
		{"indigo500(", 9},
		{"indigo500)", 9},
		{"indigo500[", 9},
		{"indigo500]", 9},
		{"indigo500// comment", 9},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Indigo500()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestIndigo500Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_i"},
		{"i_"},
		{"_in"},
		{"in_"},
		{"_ind"},
		{"ind_"},
		{"_indi"},
		{"indi_"},
		{"_indig"},
		{"indig_"},
		{"_indigo"},
		{"indigo_"},
		{"_indigo5"},
		{"indigo5_"},
		{"_indigo50"},
		{"indigo50_"},
		{"_indigo500"},
		{"indigo500_"},
		{"9713i"},
		{"i9713"},
		{"9713in"},
		{"in9713"},
		{"9713ind"},
		{"ind9713"},
		{"9713indi"},
		{"indi9713"},
		{"9713indig"},
		{"indig9713"},
		{"9713indigo"},
		{"indigo9713"},
		{"9713indigo5"},
		{"indigo59713"},
		{"9713indigo50"},
		{"indigo509713"},
		{"9713indigo500"},
		{"indigo5009713"},
		{"Zi"},
		{"iZ"},
		{"Zin"},
		{"inZ"},
		{"Zind"},
		{"indZ"},
		{"Zindi"},
		{"indiZ"},
		{"Zindig"},
		{"indigZ"},
		{"Zindigo"},
		{"indigoZ"},
		{"Zindigo5"},
		{"indigo5Z"},
		{"Zindigo50"},
		{"indigo50Z"},
		{"Zindigo500"},
		{"indigo500Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Indigo500()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestIndigo600Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"indigo600", 9},
		{"indigo600 ", 9},
		{"indigo600\n", 9},
		{"indigo600.", 9},
		{"indigo600:", 9},
		{"indigo600,", 9},
		{"indigo600\"", 9},
		{"indigo600(", 9},
		{"indigo600)", 9},
		{"indigo600[", 9},
		{"indigo600]", 9},
		{"indigo600// comment", 9},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Indigo600()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestIndigo600Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_i"},
		{"i_"},
		{"_in"},
		{"in_"},
		{"_ind"},
		{"ind_"},
		{"_indi"},
		{"indi_"},
		{"_indig"},
		{"indig_"},
		{"_indigo"},
		{"indigo_"},
		{"_indigo6"},
		{"indigo6_"},
		{"_indigo60"},
		{"indigo60_"},
		{"_indigo600"},
		{"indigo600_"},
		{"9713i"},
		{"i9713"},
		{"9713in"},
		{"in9713"},
		{"9713ind"},
		{"ind9713"},
		{"9713indi"},
		{"indi9713"},
		{"9713indig"},
		{"indig9713"},
		{"9713indigo"},
		{"indigo9713"},
		{"9713indigo6"},
		{"indigo69713"},
		{"9713indigo60"},
		{"indigo609713"},
		{"9713indigo600"},
		{"indigo6009713"},
		{"Zi"},
		{"iZ"},
		{"Zin"},
		{"inZ"},
		{"Zind"},
		{"indZ"},
		{"Zindi"},
		{"indiZ"},
		{"Zindig"},
		{"indigZ"},
		{"Zindigo"},
		{"indigoZ"},
		{"Zindigo6"},
		{"indigo6Z"},
		{"Zindigo60"},
		{"indigo60Z"},
		{"Zindigo600"},
		{"indigo600Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Indigo600()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestIndigo700Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"indigo700", 9},
		{"indigo700 ", 9},
		{"indigo700\n", 9},
		{"indigo700.", 9},
		{"indigo700:", 9},
		{"indigo700,", 9},
		{"indigo700\"", 9},
		{"indigo700(", 9},
		{"indigo700)", 9},
		{"indigo700[", 9},
		{"indigo700]", 9},
		{"indigo700// comment", 9},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Indigo700()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestIndigo700Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_i"},
		{"i_"},
		{"_in"},
		{"in_"},
		{"_ind"},
		{"ind_"},
		{"_indi"},
		{"indi_"},
		{"_indig"},
		{"indig_"},
		{"_indigo"},
		{"indigo_"},
		{"_indigo7"},
		{"indigo7_"},
		{"_indigo70"},
		{"indigo70_"},
		{"_indigo700"},
		{"indigo700_"},
		{"9713i"},
		{"i9713"},
		{"9713in"},
		{"in9713"},
		{"9713ind"},
		{"ind9713"},
		{"9713indi"},
		{"indi9713"},
		{"9713indig"},
		{"indig9713"},
		{"9713indigo"},
		{"indigo9713"},
		{"9713indigo7"},
		{"indigo79713"},
		{"9713indigo70"},
		{"indigo709713"},
		{"9713indigo700"},
		{"indigo7009713"},
		{"Zi"},
		{"iZ"},
		{"Zin"},
		{"inZ"},
		{"Zind"},
		{"indZ"},
		{"Zindi"},
		{"indiZ"},
		{"Zindig"},
		{"indigZ"},
		{"Zindigo"},
		{"indigoZ"},
		{"Zindigo7"},
		{"indigo7Z"},
		{"Zindigo70"},
		{"indigo70Z"},
		{"Zindigo700"},
		{"indigo700Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Indigo700()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestIndigo800Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"indigo800", 9},
		{"indigo800 ", 9},
		{"indigo800\n", 9},
		{"indigo800.", 9},
		{"indigo800:", 9},
		{"indigo800,", 9},
		{"indigo800\"", 9},
		{"indigo800(", 9},
		{"indigo800)", 9},
		{"indigo800[", 9},
		{"indigo800]", 9},
		{"indigo800// comment", 9},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Indigo800()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestIndigo800Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_i"},
		{"i_"},
		{"_in"},
		{"in_"},
		{"_ind"},
		{"ind_"},
		{"_indi"},
		{"indi_"},
		{"_indig"},
		{"indig_"},
		{"_indigo"},
		{"indigo_"},
		{"_indigo8"},
		{"indigo8_"},
		{"_indigo80"},
		{"indigo80_"},
		{"_indigo800"},
		{"indigo800_"},
		{"9713i"},
		{"i9713"},
		{"9713in"},
		{"in9713"},
		{"9713ind"},
		{"ind9713"},
		{"9713indi"},
		{"indi9713"},
		{"9713indig"},
		{"indig9713"},
		{"9713indigo"},
		{"indigo9713"},
		{"9713indigo8"},
		{"indigo89713"},
		{"9713indigo80"},
		{"indigo809713"},
		{"9713indigo800"},
		{"indigo8009713"},
		{"Zi"},
		{"iZ"},
		{"Zin"},
		{"inZ"},
		{"Zind"},
		{"indZ"},
		{"Zindi"},
		{"indiZ"},
		{"Zindig"},
		{"indigZ"},
		{"Zindigo"},
		{"indigoZ"},
		{"Zindigo8"},
		{"indigo8Z"},
		{"Zindigo80"},
		{"indigo80Z"},
		{"Zindigo800"},
		{"indigo800Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Indigo800()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestIndigo900Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"indigo900", 9},
		{"indigo900 ", 9},
		{"indigo900\n", 9},
		{"indigo900.", 9},
		{"indigo900:", 9},
		{"indigo900,", 9},
		{"indigo900\"", 9},
		{"indigo900(", 9},
		{"indigo900)", 9},
		{"indigo900[", 9},
		{"indigo900]", 9},
		{"indigo900// comment", 9},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Indigo900()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestIndigo900Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_i"},
		{"i_"},
		{"_in"},
		{"in_"},
		{"_ind"},
		{"ind_"},
		{"_indi"},
		{"indi_"},
		{"_indig"},
		{"indig_"},
		{"_indigo"},
		{"indigo_"},
		{"_indigo9"},
		{"indigo9_"},
		{"_indigo90"},
		{"indigo90_"},
		{"_indigo900"},
		{"indigo900_"},
		{"9713i"},
		{"i9713"},
		{"9713in"},
		{"in9713"},
		{"9713ind"},
		{"ind9713"},
		{"9713indi"},
		{"indi9713"},
		{"9713indig"},
		{"indig9713"},
		{"9713indigo"},
		{"indigo9713"},
		{"9713indigo9"},
		{"indigo99713"},
		{"9713indigo90"},
		{"indigo909713"},
		{"9713indigo900"},
		{"indigo9009713"},
		{"Zi"},
		{"iZ"},
		{"Zin"},
		{"inZ"},
		{"Zind"},
		{"indZ"},
		{"Zindi"},
		{"indiZ"},
		{"Zindig"},
		{"indigZ"},
		{"Zindigo"},
		{"indigoZ"},
		{"Zindigo9"},
		{"indigo9Z"},
		{"Zindigo90"},
		{"indigo90Z"},
		{"Zindigo900"},
		{"indigo900Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Indigo900()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestIndigoA100Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"indigoA100", 10},
		{"indigoA100 ", 10},
		{"indigoA100\n", 10},
		{"indigoA100.", 10},
		{"indigoA100:", 10},
		{"indigoA100,", 10},
		{"indigoA100\"", 10},
		{"indigoA100(", 10},
		{"indigoA100)", 10},
		{"indigoA100[", 10},
		{"indigoA100]", 10},
		{"indigoA100// comment", 10},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.IndigoA100()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestIndigoA100Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_i"},
		{"i_"},
		{"_in"},
		{"in_"},
		{"_ind"},
		{"ind_"},
		{"_indi"},
		{"indi_"},
		{"_indig"},
		{"indig_"},
		{"_indigo"},
		{"indigo_"},
		{"_indigoA"},
		{"indigoA_"},
		{"_indigoA1"},
		{"indigoA1_"},
		{"_indigoA10"},
		{"indigoA10_"},
		{"_indigoA100"},
		{"indigoA100_"},
		{"9713i"},
		{"i9713"},
		{"9713in"},
		{"in9713"},
		{"9713ind"},
		{"ind9713"},
		{"9713indi"},
		{"indi9713"},
		{"9713indig"},
		{"indig9713"},
		{"9713indigo"},
		{"indigo9713"},
		{"9713indigoA"},
		{"indigoA9713"},
		{"9713indigoA1"},
		{"indigoA19713"},
		{"9713indigoA10"},
		{"indigoA109713"},
		{"9713indigoA100"},
		{"indigoA1009713"},
		{"Zi"},
		{"iZ"},
		{"Zin"},
		{"inZ"},
		{"Zind"},
		{"indZ"},
		{"Zindi"},
		{"indiZ"},
		{"Zindig"},
		{"indigZ"},
		{"Zindigo"},
		{"indigoZ"},
		{"ZindigoA"},
		{"indigoAZ"},
		{"ZindigoA1"},
		{"indigoA1Z"},
		{"ZindigoA10"},
		{"indigoA10Z"},
		{"ZindigoA100"},
		{"indigoA100Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.IndigoA100()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestIndigoA200Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"indigoA200", 10},
		{"indigoA200 ", 10},
		{"indigoA200\n", 10},
		{"indigoA200.", 10},
		{"indigoA200:", 10},
		{"indigoA200,", 10},
		{"indigoA200\"", 10},
		{"indigoA200(", 10},
		{"indigoA200)", 10},
		{"indigoA200[", 10},
		{"indigoA200]", 10},
		{"indigoA200// comment", 10},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.IndigoA200()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestIndigoA200Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_i"},
		{"i_"},
		{"_in"},
		{"in_"},
		{"_ind"},
		{"ind_"},
		{"_indi"},
		{"indi_"},
		{"_indig"},
		{"indig_"},
		{"_indigo"},
		{"indigo_"},
		{"_indigoA"},
		{"indigoA_"},
		{"_indigoA2"},
		{"indigoA2_"},
		{"_indigoA20"},
		{"indigoA20_"},
		{"_indigoA200"},
		{"indigoA200_"},
		{"9713i"},
		{"i9713"},
		{"9713in"},
		{"in9713"},
		{"9713ind"},
		{"ind9713"},
		{"9713indi"},
		{"indi9713"},
		{"9713indig"},
		{"indig9713"},
		{"9713indigo"},
		{"indigo9713"},
		{"9713indigoA"},
		{"indigoA9713"},
		{"9713indigoA2"},
		{"indigoA29713"},
		{"9713indigoA20"},
		{"indigoA209713"},
		{"9713indigoA200"},
		{"indigoA2009713"},
		{"Zi"},
		{"iZ"},
		{"Zin"},
		{"inZ"},
		{"Zind"},
		{"indZ"},
		{"Zindi"},
		{"indiZ"},
		{"Zindig"},
		{"indigZ"},
		{"Zindigo"},
		{"indigoZ"},
		{"ZindigoA"},
		{"indigoAZ"},
		{"ZindigoA2"},
		{"indigoA2Z"},
		{"ZindigoA20"},
		{"indigoA20Z"},
		{"ZindigoA200"},
		{"indigoA200Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.IndigoA200()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestIndigoA400Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"indigoA400", 10},
		{"indigoA400 ", 10},
		{"indigoA400\n", 10},
		{"indigoA400.", 10},
		{"indigoA400:", 10},
		{"indigoA400,", 10},
		{"indigoA400\"", 10},
		{"indigoA400(", 10},
		{"indigoA400)", 10},
		{"indigoA400[", 10},
		{"indigoA400]", 10},
		{"indigoA400// comment", 10},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.IndigoA400()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestIndigoA400Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_i"},
		{"i_"},
		{"_in"},
		{"in_"},
		{"_ind"},
		{"ind_"},
		{"_indi"},
		{"indi_"},
		{"_indig"},
		{"indig_"},
		{"_indigo"},
		{"indigo_"},
		{"_indigoA"},
		{"indigoA_"},
		{"_indigoA4"},
		{"indigoA4_"},
		{"_indigoA40"},
		{"indigoA40_"},
		{"_indigoA400"},
		{"indigoA400_"},
		{"9713i"},
		{"i9713"},
		{"9713in"},
		{"in9713"},
		{"9713ind"},
		{"ind9713"},
		{"9713indi"},
		{"indi9713"},
		{"9713indig"},
		{"indig9713"},
		{"9713indigo"},
		{"indigo9713"},
		{"9713indigoA"},
		{"indigoA9713"},
		{"9713indigoA4"},
		{"indigoA49713"},
		{"9713indigoA40"},
		{"indigoA409713"},
		{"9713indigoA400"},
		{"indigoA4009713"},
		{"Zi"},
		{"iZ"},
		{"Zin"},
		{"inZ"},
		{"Zind"},
		{"indZ"},
		{"Zindi"},
		{"indiZ"},
		{"Zindig"},
		{"indigZ"},
		{"Zindigo"},
		{"indigoZ"},
		{"ZindigoA"},
		{"indigoAZ"},
		{"ZindigoA4"},
		{"indigoA4Z"},
		{"ZindigoA40"},
		{"indigoA40Z"},
		{"ZindigoA400"},
		{"indigoA400Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.IndigoA400()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestIndigoA700Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"indigoA700", 10},
		{"indigoA700 ", 10},
		{"indigoA700\n", 10},
		{"indigoA700.", 10},
		{"indigoA700:", 10},
		{"indigoA700,", 10},
		{"indigoA700\"", 10},
		{"indigoA700(", 10},
		{"indigoA700)", 10},
		{"indigoA700[", 10},
		{"indigoA700]", 10},
		{"indigoA700// comment", 10},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.IndigoA700()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestIndigoA700Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_i"},
		{"i_"},
		{"_in"},
		{"in_"},
		{"_ind"},
		{"ind_"},
		{"_indi"},
		{"indi_"},
		{"_indig"},
		{"indig_"},
		{"_indigo"},
		{"indigo_"},
		{"_indigoA"},
		{"indigoA_"},
		{"_indigoA7"},
		{"indigoA7_"},
		{"_indigoA70"},
		{"indigoA70_"},
		{"_indigoA700"},
		{"indigoA700_"},
		{"9713i"},
		{"i9713"},
		{"9713in"},
		{"in9713"},
		{"9713ind"},
		{"ind9713"},
		{"9713indi"},
		{"indi9713"},
		{"9713indig"},
		{"indig9713"},
		{"9713indigo"},
		{"indigo9713"},
		{"9713indigoA"},
		{"indigoA9713"},
		{"9713indigoA7"},
		{"indigoA79713"},
		{"9713indigoA70"},
		{"indigoA709713"},
		{"9713indigoA700"},
		{"indigoA7009713"},
		{"Zi"},
		{"iZ"},
		{"Zin"},
		{"inZ"},
		{"Zind"},
		{"indZ"},
		{"Zindi"},
		{"indiZ"},
		{"Zindig"},
		{"indigZ"},
		{"Zindigo"},
		{"indigoZ"},
		{"ZindigoA"},
		{"indigoAZ"},
		{"ZindigoA7"},
		{"indigoA7Z"},
		{"ZindigoA70"},
		{"indigoA70Z"},
		{"ZindigoA700"},
		{"indigoA700Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.IndigoA700()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestBlue50Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"blue50", 6},
		{"blue50 ", 6},
		{"blue50\n", 6},
		{"blue50.", 6},
		{"blue50:", 6},
		{"blue50,", 6},
		{"blue50\"", 6},
		{"blue50(", 6},
		{"blue50)", 6},
		{"blue50[", 6},
		{"blue50]", 6},
		{"blue50// comment", 6},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Blue50()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestBlue50Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_b"},
		{"b_"},
		{"_bl"},
		{"bl_"},
		{"_blu"},
		{"blu_"},
		{"_blue"},
		{"blue_"},
		{"_blue5"},
		{"blue5_"},
		{"_blue50"},
		{"blue50_"},
		{"9713b"},
		{"b9713"},
		{"9713bl"},
		{"bl9713"},
		{"9713blu"},
		{"blu9713"},
		{"9713blue"},
		{"blue9713"},
		{"9713blue5"},
		{"blue59713"},
		{"9713blue50"},
		{"blue509713"},
		{"Zb"},
		{"bZ"},
		{"Zbl"},
		{"blZ"},
		{"Zblu"},
		{"bluZ"},
		{"Zblue"},
		{"blueZ"},
		{"Zblue5"},
		{"blue5Z"},
		{"Zblue50"},
		{"blue50Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Blue50()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestBlue100Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"blue100", 7},
		{"blue100 ", 7},
		{"blue100\n", 7},
		{"blue100.", 7},
		{"blue100:", 7},
		{"blue100,", 7},
		{"blue100\"", 7},
		{"blue100(", 7},
		{"blue100)", 7},
		{"blue100[", 7},
		{"blue100]", 7},
		{"blue100// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Blue100()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestBlue100Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_b"},
		{"b_"},
		{"_bl"},
		{"bl_"},
		{"_blu"},
		{"blu_"},
		{"_blue"},
		{"blue_"},
		{"_blue1"},
		{"blue1_"},
		{"_blue10"},
		{"blue10_"},
		{"_blue100"},
		{"blue100_"},
		{"9713b"},
		{"b9713"},
		{"9713bl"},
		{"bl9713"},
		{"9713blu"},
		{"blu9713"},
		{"9713blue"},
		{"blue9713"},
		{"9713blue1"},
		{"blue19713"},
		{"9713blue10"},
		{"blue109713"},
		{"9713blue100"},
		{"blue1009713"},
		{"Zb"},
		{"bZ"},
		{"Zbl"},
		{"blZ"},
		{"Zblu"},
		{"bluZ"},
		{"Zblue"},
		{"blueZ"},
		{"Zblue1"},
		{"blue1Z"},
		{"Zblue10"},
		{"blue10Z"},
		{"Zblue100"},
		{"blue100Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Blue100()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestBlue200Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"blue200", 7},
		{"blue200 ", 7},
		{"blue200\n", 7},
		{"blue200.", 7},
		{"blue200:", 7},
		{"blue200,", 7},
		{"blue200\"", 7},
		{"blue200(", 7},
		{"blue200)", 7},
		{"blue200[", 7},
		{"blue200]", 7},
		{"blue200// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Blue200()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestBlue200Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_b"},
		{"b_"},
		{"_bl"},
		{"bl_"},
		{"_blu"},
		{"blu_"},
		{"_blue"},
		{"blue_"},
		{"_blue2"},
		{"blue2_"},
		{"_blue20"},
		{"blue20_"},
		{"_blue200"},
		{"blue200_"},
		{"9713b"},
		{"b9713"},
		{"9713bl"},
		{"bl9713"},
		{"9713blu"},
		{"blu9713"},
		{"9713blue"},
		{"blue9713"},
		{"9713blue2"},
		{"blue29713"},
		{"9713blue20"},
		{"blue209713"},
		{"9713blue200"},
		{"blue2009713"},
		{"Zb"},
		{"bZ"},
		{"Zbl"},
		{"blZ"},
		{"Zblu"},
		{"bluZ"},
		{"Zblue"},
		{"blueZ"},
		{"Zblue2"},
		{"blue2Z"},
		{"Zblue20"},
		{"blue20Z"},
		{"Zblue200"},
		{"blue200Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Blue200()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestBlue300Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"blue300", 7},
		{"blue300 ", 7},
		{"blue300\n", 7},
		{"blue300.", 7},
		{"blue300:", 7},
		{"blue300,", 7},
		{"blue300\"", 7},
		{"blue300(", 7},
		{"blue300)", 7},
		{"blue300[", 7},
		{"blue300]", 7},
		{"blue300// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Blue300()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestBlue300Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_b"},
		{"b_"},
		{"_bl"},
		{"bl_"},
		{"_blu"},
		{"blu_"},
		{"_blue"},
		{"blue_"},
		{"_blue3"},
		{"blue3_"},
		{"_blue30"},
		{"blue30_"},
		{"_blue300"},
		{"blue300_"},
		{"9713b"},
		{"b9713"},
		{"9713bl"},
		{"bl9713"},
		{"9713blu"},
		{"blu9713"},
		{"9713blue"},
		{"blue9713"},
		{"9713blue3"},
		{"blue39713"},
		{"9713blue30"},
		{"blue309713"},
		{"9713blue300"},
		{"blue3009713"},
		{"Zb"},
		{"bZ"},
		{"Zbl"},
		{"blZ"},
		{"Zblu"},
		{"bluZ"},
		{"Zblue"},
		{"blueZ"},
		{"Zblue3"},
		{"blue3Z"},
		{"Zblue30"},
		{"blue30Z"},
		{"Zblue300"},
		{"blue300Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Blue300()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestBlue400Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"blue400", 7},
		{"blue400 ", 7},
		{"blue400\n", 7},
		{"blue400.", 7},
		{"blue400:", 7},
		{"blue400,", 7},
		{"blue400\"", 7},
		{"blue400(", 7},
		{"blue400)", 7},
		{"blue400[", 7},
		{"blue400]", 7},
		{"blue400// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Blue400()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestBlue400Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_b"},
		{"b_"},
		{"_bl"},
		{"bl_"},
		{"_blu"},
		{"blu_"},
		{"_blue"},
		{"blue_"},
		{"_blue4"},
		{"blue4_"},
		{"_blue40"},
		{"blue40_"},
		{"_blue400"},
		{"blue400_"},
		{"9713b"},
		{"b9713"},
		{"9713bl"},
		{"bl9713"},
		{"9713blu"},
		{"blu9713"},
		{"9713blue"},
		{"blue9713"},
		{"9713blue4"},
		{"blue49713"},
		{"9713blue40"},
		{"blue409713"},
		{"9713blue400"},
		{"blue4009713"},
		{"Zb"},
		{"bZ"},
		{"Zbl"},
		{"blZ"},
		{"Zblu"},
		{"bluZ"},
		{"Zblue"},
		{"blueZ"},
		{"Zblue4"},
		{"blue4Z"},
		{"Zblue40"},
		{"blue40Z"},
		{"Zblue400"},
		{"blue400Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Blue400()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestBlue500Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"blue500", 7},
		{"blue500 ", 7},
		{"blue500\n", 7},
		{"blue500.", 7},
		{"blue500:", 7},
		{"blue500,", 7},
		{"blue500\"", 7},
		{"blue500(", 7},
		{"blue500)", 7},
		{"blue500[", 7},
		{"blue500]", 7},
		{"blue500// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Blue500()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestBlue500Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_b"},
		{"b_"},
		{"_bl"},
		{"bl_"},
		{"_blu"},
		{"blu_"},
		{"_blue"},
		{"blue_"},
		{"_blue5"},
		{"blue5_"},
		{"_blue50"},
		{"blue50_"},
		{"_blue500"},
		{"blue500_"},
		{"9713b"},
		{"b9713"},
		{"9713bl"},
		{"bl9713"},
		{"9713blu"},
		{"blu9713"},
		{"9713blue"},
		{"blue9713"},
		{"9713blue5"},
		{"blue59713"},
		{"9713blue50"},
		{"blue509713"},
		{"9713blue500"},
		{"blue5009713"},
		{"Zb"},
		{"bZ"},
		{"Zbl"},
		{"blZ"},
		{"Zblu"},
		{"bluZ"},
		{"Zblue"},
		{"blueZ"},
		{"Zblue5"},
		{"blue5Z"},
		{"Zblue50"},
		{"blue50Z"},
		{"Zblue500"},
		{"blue500Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Blue500()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestBlue600Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"blue600", 7},
		{"blue600 ", 7},
		{"blue600\n", 7},
		{"blue600.", 7},
		{"blue600:", 7},
		{"blue600,", 7},
		{"blue600\"", 7},
		{"blue600(", 7},
		{"blue600)", 7},
		{"blue600[", 7},
		{"blue600]", 7},
		{"blue600// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Blue600()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestBlue600Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_b"},
		{"b_"},
		{"_bl"},
		{"bl_"},
		{"_blu"},
		{"blu_"},
		{"_blue"},
		{"blue_"},
		{"_blue6"},
		{"blue6_"},
		{"_blue60"},
		{"blue60_"},
		{"_blue600"},
		{"blue600_"},
		{"9713b"},
		{"b9713"},
		{"9713bl"},
		{"bl9713"},
		{"9713blu"},
		{"blu9713"},
		{"9713blue"},
		{"blue9713"},
		{"9713blue6"},
		{"blue69713"},
		{"9713blue60"},
		{"blue609713"},
		{"9713blue600"},
		{"blue6009713"},
		{"Zb"},
		{"bZ"},
		{"Zbl"},
		{"blZ"},
		{"Zblu"},
		{"bluZ"},
		{"Zblue"},
		{"blueZ"},
		{"Zblue6"},
		{"blue6Z"},
		{"Zblue60"},
		{"blue60Z"},
		{"Zblue600"},
		{"blue600Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Blue600()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestBlue700Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"blue700", 7},
		{"blue700 ", 7},
		{"blue700\n", 7},
		{"blue700.", 7},
		{"blue700:", 7},
		{"blue700,", 7},
		{"blue700\"", 7},
		{"blue700(", 7},
		{"blue700)", 7},
		{"blue700[", 7},
		{"blue700]", 7},
		{"blue700// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Blue700()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestBlue700Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_b"},
		{"b_"},
		{"_bl"},
		{"bl_"},
		{"_blu"},
		{"blu_"},
		{"_blue"},
		{"blue_"},
		{"_blue7"},
		{"blue7_"},
		{"_blue70"},
		{"blue70_"},
		{"_blue700"},
		{"blue700_"},
		{"9713b"},
		{"b9713"},
		{"9713bl"},
		{"bl9713"},
		{"9713blu"},
		{"blu9713"},
		{"9713blue"},
		{"blue9713"},
		{"9713blue7"},
		{"blue79713"},
		{"9713blue70"},
		{"blue709713"},
		{"9713blue700"},
		{"blue7009713"},
		{"Zb"},
		{"bZ"},
		{"Zbl"},
		{"blZ"},
		{"Zblu"},
		{"bluZ"},
		{"Zblue"},
		{"blueZ"},
		{"Zblue7"},
		{"blue7Z"},
		{"Zblue70"},
		{"blue70Z"},
		{"Zblue700"},
		{"blue700Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Blue700()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestBlue800Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"blue800", 7},
		{"blue800 ", 7},
		{"blue800\n", 7},
		{"blue800.", 7},
		{"blue800:", 7},
		{"blue800,", 7},
		{"blue800\"", 7},
		{"blue800(", 7},
		{"blue800)", 7},
		{"blue800[", 7},
		{"blue800]", 7},
		{"blue800// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Blue800()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestBlue800Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_b"},
		{"b_"},
		{"_bl"},
		{"bl_"},
		{"_blu"},
		{"blu_"},
		{"_blue"},
		{"blue_"},
		{"_blue8"},
		{"blue8_"},
		{"_blue80"},
		{"blue80_"},
		{"_blue800"},
		{"blue800_"},
		{"9713b"},
		{"b9713"},
		{"9713bl"},
		{"bl9713"},
		{"9713blu"},
		{"blu9713"},
		{"9713blue"},
		{"blue9713"},
		{"9713blue8"},
		{"blue89713"},
		{"9713blue80"},
		{"blue809713"},
		{"9713blue800"},
		{"blue8009713"},
		{"Zb"},
		{"bZ"},
		{"Zbl"},
		{"blZ"},
		{"Zblu"},
		{"bluZ"},
		{"Zblue"},
		{"blueZ"},
		{"Zblue8"},
		{"blue8Z"},
		{"Zblue80"},
		{"blue80Z"},
		{"Zblue800"},
		{"blue800Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Blue800()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestBlue900Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"blue900", 7},
		{"blue900 ", 7},
		{"blue900\n", 7},
		{"blue900.", 7},
		{"blue900:", 7},
		{"blue900,", 7},
		{"blue900\"", 7},
		{"blue900(", 7},
		{"blue900)", 7},
		{"blue900[", 7},
		{"blue900]", 7},
		{"blue900// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Blue900()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestBlue900Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_b"},
		{"b_"},
		{"_bl"},
		{"bl_"},
		{"_blu"},
		{"blu_"},
		{"_blue"},
		{"blue_"},
		{"_blue9"},
		{"blue9_"},
		{"_blue90"},
		{"blue90_"},
		{"_blue900"},
		{"blue900_"},
		{"9713b"},
		{"b9713"},
		{"9713bl"},
		{"bl9713"},
		{"9713blu"},
		{"blu9713"},
		{"9713blue"},
		{"blue9713"},
		{"9713blue9"},
		{"blue99713"},
		{"9713blue90"},
		{"blue909713"},
		{"9713blue900"},
		{"blue9009713"},
		{"Zb"},
		{"bZ"},
		{"Zbl"},
		{"blZ"},
		{"Zblu"},
		{"bluZ"},
		{"Zblue"},
		{"blueZ"},
		{"Zblue9"},
		{"blue9Z"},
		{"Zblue90"},
		{"blue90Z"},
		{"Zblue900"},
		{"blue900Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Blue900()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestBlueA100Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"blueA100", 8},
		{"blueA100 ", 8},
		{"blueA100\n", 8},
		{"blueA100.", 8},
		{"blueA100:", 8},
		{"blueA100,", 8},
		{"blueA100\"", 8},
		{"blueA100(", 8},
		{"blueA100)", 8},
		{"blueA100[", 8},
		{"blueA100]", 8},
		{"blueA100// comment", 8},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.BlueA100()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestBlueA100Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_b"},
		{"b_"},
		{"_bl"},
		{"bl_"},
		{"_blu"},
		{"blu_"},
		{"_blue"},
		{"blue_"},
		{"_blueA"},
		{"blueA_"},
		{"_blueA1"},
		{"blueA1_"},
		{"_blueA10"},
		{"blueA10_"},
		{"_blueA100"},
		{"blueA100_"},
		{"9713b"},
		{"b9713"},
		{"9713bl"},
		{"bl9713"},
		{"9713blu"},
		{"blu9713"},
		{"9713blue"},
		{"blue9713"},
		{"9713blueA"},
		{"blueA9713"},
		{"9713blueA1"},
		{"blueA19713"},
		{"9713blueA10"},
		{"blueA109713"},
		{"9713blueA100"},
		{"blueA1009713"},
		{"Zb"},
		{"bZ"},
		{"Zbl"},
		{"blZ"},
		{"Zblu"},
		{"bluZ"},
		{"Zblue"},
		{"blueZ"},
		{"ZblueA"},
		{"blueAZ"},
		{"ZblueA1"},
		{"blueA1Z"},
		{"ZblueA10"},
		{"blueA10Z"},
		{"ZblueA100"},
		{"blueA100Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.BlueA100()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestBlueA200Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"blueA200", 8},
		{"blueA200 ", 8},
		{"blueA200\n", 8},
		{"blueA200.", 8},
		{"blueA200:", 8},
		{"blueA200,", 8},
		{"blueA200\"", 8},
		{"blueA200(", 8},
		{"blueA200)", 8},
		{"blueA200[", 8},
		{"blueA200]", 8},
		{"blueA200// comment", 8},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.BlueA200()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestBlueA200Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_b"},
		{"b_"},
		{"_bl"},
		{"bl_"},
		{"_blu"},
		{"blu_"},
		{"_blue"},
		{"blue_"},
		{"_blueA"},
		{"blueA_"},
		{"_blueA2"},
		{"blueA2_"},
		{"_blueA20"},
		{"blueA20_"},
		{"_blueA200"},
		{"blueA200_"},
		{"9713b"},
		{"b9713"},
		{"9713bl"},
		{"bl9713"},
		{"9713blu"},
		{"blu9713"},
		{"9713blue"},
		{"blue9713"},
		{"9713blueA"},
		{"blueA9713"},
		{"9713blueA2"},
		{"blueA29713"},
		{"9713blueA20"},
		{"blueA209713"},
		{"9713blueA200"},
		{"blueA2009713"},
		{"Zb"},
		{"bZ"},
		{"Zbl"},
		{"blZ"},
		{"Zblu"},
		{"bluZ"},
		{"Zblue"},
		{"blueZ"},
		{"ZblueA"},
		{"blueAZ"},
		{"ZblueA2"},
		{"blueA2Z"},
		{"ZblueA20"},
		{"blueA20Z"},
		{"ZblueA200"},
		{"blueA200Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.BlueA200()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestBlueA400Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"blueA400", 8},
		{"blueA400 ", 8},
		{"blueA400\n", 8},
		{"blueA400.", 8},
		{"blueA400:", 8},
		{"blueA400,", 8},
		{"blueA400\"", 8},
		{"blueA400(", 8},
		{"blueA400)", 8},
		{"blueA400[", 8},
		{"blueA400]", 8},
		{"blueA400// comment", 8},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.BlueA400()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestBlueA400Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_b"},
		{"b_"},
		{"_bl"},
		{"bl_"},
		{"_blu"},
		{"blu_"},
		{"_blue"},
		{"blue_"},
		{"_blueA"},
		{"blueA_"},
		{"_blueA4"},
		{"blueA4_"},
		{"_blueA40"},
		{"blueA40_"},
		{"_blueA400"},
		{"blueA400_"},
		{"9713b"},
		{"b9713"},
		{"9713bl"},
		{"bl9713"},
		{"9713blu"},
		{"blu9713"},
		{"9713blue"},
		{"blue9713"},
		{"9713blueA"},
		{"blueA9713"},
		{"9713blueA4"},
		{"blueA49713"},
		{"9713blueA40"},
		{"blueA409713"},
		{"9713blueA400"},
		{"blueA4009713"},
		{"Zb"},
		{"bZ"},
		{"Zbl"},
		{"blZ"},
		{"Zblu"},
		{"bluZ"},
		{"Zblue"},
		{"blueZ"},
		{"ZblueA"},
		{"blueAZ"},
		{"ZblueA4"},
		{"blueA4Z"},
		{"ZblueA40"},
		{"blueA40Z"},
		{"ZblueA400"},
		{"blueA400Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.BlueA400()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestBlueA700Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"blueA700", 8},
		{"blueA700 ", 8},
		{"blueA700\n", 8},
		{"blueA700.", 8},
		{"blueA700:", 8},
		{"blueA700,", 8},
		{"blueA700\"", 8},
		{"blueA700(", 8},
		{"blueA700)", 8},
		{"blueA700[", 8},
		{"blueA700]", 8},
		{"blueA700// comment", 8},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.BlueA700()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestBlueA700Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_b"},
		{"b_"},
		{"_bl"},
		{"bl_"},
		{"_blu"},
		{"blu_"},
		{"_blue"},
		{"blue_"},
		{"_blueA"},
		{"blueA_"},
		{"_blueA7"},
		{"blueA7_"},
		{"_blueA70"},
		{"blueA70_"},
		{"_blueA700"},
		{"blueA700_"},
		{"9713b"},
		{"b9713"},
		{"9713bl"},
		{"bl9713"},
		{"9713blu"},
		{"blu9713"},
		{"9713blue"},
		{"blue9713"},
		{"9713blueA"},
		{"blueA9713"},
		{"9713blueA7"},
		{"blueA79713"},
		{"9713blueA70"},
		{"blueA709713"},
		{"9713blueA700"},
		{"blueA7009713"},
		{"Zb"},
		{"bZ"},
		{"Zbl"},
		{"blZ"},
		{"Zblu"},
		{"bluZ"},
		{"Zblue"},
		{"blueZ"},
		{"ZblueA"},
		{"blueAZ"},
		{"ZblueA7"},
		{"blueA7Z"},
		{"ZblueA70"},
		{"blueA70Z"},
		{"ZblueA700"},
		{"blueA700Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.BlueA700()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestLightBlue50Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"lightBlue50", 11},
		{"lightBlue50 ", 11},
		{"lightBlue50\n", 11},
		{"lightBlue50.", 11},
		{"lightBlue50:", 11},
		{"lightBlue50,", 11},
		{"lightBlue50\"", 11},
		{"lightBlue50(", 11},
		{"lightBlue50)", 11},
		{"lightBlue50[", 11},
		{"lightBlue50]", 11},
		{"lightBlue50// comment", 11},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.LightBlue50()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestLightBlue50Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_l"},
		{"l_"},
		{"_li"},
		{"li_"},
		{"_lig"},
		{"lig_"},
		{"_ligh"},
		{"ligh_"},
		{"_light"},
		{"light_"},
		{"_lightB"},
		{"lightB_"},
		{"_lightBl"},
		{"lightBl_"},
		{"_lightBlu"},
		{"lightBlu_"},
		{"_lightBlue"},
		{"lightBlue_"},
		{"_lightBlue5"},
		{"lightBlue5_"},
		{"_lightBlue50"},
		{"lightBlue50_"},
		{"9713l"},
		{"l9713"},
		{"9713li"},
		{"li9713"},
		{"9713lig"},
		{"lig9713"},
		{"9713ligh"},
		{"ligh9713"},
		{"9713light"},
		{"light9713"},
		{"9713lightB"},
		{"lightB9713"},
		{"9713lightBl"},
		{"lightBl9713"},
		{"9713lightBlu"},
		{"lightBlu9713"},
		{"9713lightBlue"},
		{"lightBlue9713"},
		{"9713lightBlue5"},
		{"lightBlue59713"},
		{"9713lightBlue50"},
		{"lightBlue509713"},
		{"Zl"},
		{"lZ"},
		{"Zli"},
		{"liZ"},
		{"Zlig"},
		{"ligZ"},
		{"Zligh"},
		{"lighZ"},
		{"Zlight"},
		{"lightZ"},
		{"ZlightB"},
		{"lightBZ"},
		{"ZlightBl"},
		{"lightBlZ"},
		{"ZlightBlu"},
		{"lightBluZ"},
		{"ZlightBlue"},
		{"lightBlueZ"},
		{"ZlightBlue5"},
		{"lightBlue5Z"},
		{"ZlightBlue50"},
		{"lightBlue50Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.LightBlue50()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestLightBlue100Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"lightBlue100", 12},
		{"lightBlue100 ", 12},
		{"lightBlue100\n", 12},
		{"lightBlue100.", 12},
		{"lightBlue100:", 12},
		{"lightBlue100,", 12},
		{"lightBlue100\"", 12},
		{"lightBlue100(", 12},
		{"lightBlue100)", 12},
		{"lightBlue100[", 12},
		{"lightBlue100]", 12},
		{"lightBlue100// comment", 12},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.LightBlue100()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestLightBlue100Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_l"},
		{"l_"},
		{"_li"},
		{"li_"},
		{"_lig"},
		{"lig_"},
		{"_ligh"},
		{"ligh_"},
		{"_light"},
		{"light_"},
		{"_lightB"},
		{"lightB_"},
		{"_lightBl"},
		{"lightBl_"},
		{"_lightBlu"},
		{"lightBlu_"},
		{"_lightBlue"},
		{"lightBlue_"},
		{"_lightBlue1"},
		{"lightBlue1_"},
		{"_lightBlue10"},
		{"lightBlue10_"},
		{"_lightBlue100"},
		{"lightBlue100_"},
		{"9713l"},
		{"l9713"},
		{"9713li"},
		{"li9713"},
		{"9713lig"},
		{"lig9713"},
		{"9713ligh"},
		{"ligh9713"},
		{"9713light"},
		{"light9713"},
		{"9713lightB"},
		{"lightB9713"},
		{"9713lightBl"},
		{"lightBl9713"},
		{"9713lightBlu"},
		{"lightBlu9713"},
		{"9713lightBlue"},
		{"lightBlue9713"},
		{"9713lightBlue1"},
		{"lightBlue19713"},
		{"9713lightBlue10"},
		{"lightBlue109713"},
		{"9713lightBlue100"},
		{"lightBlue1009713"},
		{"Zl"},
		{"lZ"},
		{"Zli"},
		{"liZ"},
		{"Zlig"},
		{"ligZ"},
		{"Zligh"},
		{"lighZ"},
		{"Zlight"},
		{"lightZ"},
		{"ZlightB"},
		{"lightBZ"},
		{"ZlightBl"},
		{"lightBlZ"},
		{"ZlightBlu"},
		{"lightBluZ"},
		{"ZlightBlue"},
		{"lightBlueZ"},
		{"ZlightBlue1"},
		{"lightBlue1Z"},
		{"ZlightBlue10"},
		{"lightBlue10Z"},
		{"ZlightBlue100"},
		{"lightBlue100Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.LightBlue100()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestLightBlue200Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"lightBlue200", 12},
		{"lightBlue200 ", 12},
		{"lightBlue200\n", 12},
		{"lightBlue200.", 12},
		{"lightBlue200:", 12},
		{"lightBlue200,", 12},
		{"lightBlue200\"", 12},
		{"lightBlue200(", 12},
		{"lightBlue200)", 12},
		{"lightBlue200[", 12},
		{"lightBlue200]", 12},
		{"lightBlue200// comment", 12},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.LightBlue200()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestLightBlue200Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_l"},
		{"l_"},
		{"_li"},
		{"li_"},
		{"_lig"},
		{"lig_"},
		{"_ligh"},
		{"ligh_"},
		{"_light"},
		{"light_"},
		{"_lightB"},
		{"lightB_"},
		{"_lightBl"},
		{"lightBl_"},
		{"_lightBlu"},
		{"lightBlu_"},
		{"_lightBlue"},
		{"lightBlue_"},
		{"_lightBlue2"},
		{"lightBlue2_"},
		{"_lightBlue20"},
		{"lightBlue20_"},
		{"_lightBlue200"},
		{"lightBlue200_"},
		{"9713l"},
		{"l9713"},
		{"9713li"},
		{"li9713"},
		{"9713lig"},
		{"lig9713"},
		{"9713ligh"},
		{"ligh9713"},
		{"9713light"},
		{"light9713"},
		{"9713lightB"},
		{"lightB9713"},
		{"9713lightBl"},
		{"lightBl9713"},
		{"9713lightBlu"},
		{"lightBlu9713"},
		{"9713lightBlue"},
		{"lightBlue9713"},
		{"9713lightBlue2"},
		{"lightBlue29713"},
		{"9713lightBlue20"},
		{"lightBlue209713"},
		{"9713lightBlue200"},
		{"lightBlue2009713"},
		{"Zl"},
		{"lZ"},
		{"Zli"},
		{"liZ"},
		{"Zlig"},
		{"ligZ"},
		{"Zligh"},
		{"lighZ"},
		{"Zlight"},
		{"lightZ"},
		{"ZlightB"},
		{"lightBZ"},
		{"ZlightBl"},
		{"lightBlZ"},
		{"ZlightBlu"},
		{"lightBluZ"},
		{"ZlightBlue"},
		{"lightBlueZ"},
		{"ZlightBlue2"},
		{"lightBlue2Z"},
		{"ZlightBlue20"},
		{"lightBlue20Z"},
		{"ZlightBlue200"},
		{"lightBlue200Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.LightBlue200()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestLightBlue300Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"lightBlue300", 12},
		{"lightBlue300 ", 12},
		{"lightBlue300\n", 12},
		{"lightBlue300.", 12},
		{"lightBlue300:", 12},
		{"lightBlue300,", 12},
		{"lightBlue300\"", 12},
		{"lightBlue300(", 12},
		{"lightBlue300)", 12},
		{"lightBlue300[", 12},
		{"lightBlue300]", 12},
		{"lightBlue300// comment", 12},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.LightBlue300()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestLightBlue300Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_l"},
		{"l_"},
		{"_li"},
		{"li_"},
		{"_lig"},
		{"lig_"},
		{"_ligh"},
		{"ligh_"},
		{"_light"},
		{"light_"},
		{"_lightB"},
		{"lightB_"},
		{"_lightBl"},
		{"lightBl_"},
		{"_lightBlu"},
		{"lightBlu_"},
		{"_lightBlue"},
		{"lightBlue_"},
		{"_lightBlue3"},
		{"lightBlue3_"},
		{"_lightBlue30"},
		{"lightBlue30_"},
		{"_lightBlue300"},
		{"lightBlue300_"},
		{"9713l"},
		{"l9713"},
		{"9713li"},
		{"li9713"},
		{"9713lig"},
		{"lig9713"},
		{"9713ligh"},
		{"ligh9713"},
		{"9713light"},
		{"light9713"},
		{"9713lightB"},
		{"lightB9713"},
		{"9713lightBl"},
		{"lightBl9713"},
		{"9713lightBlu"},
		{"lightBlu9713"},
		{"9713lightBlue"},
		{"lightBlue9713"},
		{"9713lightBlue3"},
		{"lightBlue39713"},
		{"9713lightBlue30"},
		{"lightBlue309713"},
		{"9713lightBlue300"},
		{"lightBlue3009713"},
		{"Zl"},
		{"lZ"},
		{"Zli"},
		{"liZ"},
		{"Zlig"},
		{"ligZ"},
		{"Zligh"},
		{"lighZ"},
		{"Zlight"},
		{"lightZ"},
		{"ZlightB"},
		{"lightBZ"},
		{"ZlightBl"},
		{"lightBlZ"},
		{"ZlightBlu"},
		{"lightBluZ"},
		{"ZlightBlue"},
		{"lightBlueZ"},
		{"ZlightBlue3"},
		{"lightBlue3Z"},
		{"ZlightBlue30"},
		{"lightBlue30Z"},
		{"ZlightBlue300"},
		{"lightBlue300Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.LightBlue300()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestLightBlue400Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"lightBlue400", 12},
		{"lightBlue400 ", 12},
		{"lightBlue400\n", 12},
		{"lightBlue400.", 12},
		{"lightBlue400:", 12},
		{"lightBlue400,", 12},
		{"lightBlue400\"", 12},
		{"lightBlue400(", 12},
		{"lightBlue400)", 12},
		{"lightBlue400[", 12},
		{"lightBlue400]", 12},
		{"lightBlue400// comment", 12},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.LightBlue400()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestLightBlue400Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_l"},
		{"l_"},
		{"_li"},
		{"li_"},
		{"_lig"},
		{"lig_"},
		{"_ligh"},
		{"ligh_"},
		{"_light"},
		{"light_"},
		{"_lightB"},
		{"lightB_"},
		{"_lightBl"},
		{"lightBl_"},
		{"_lightBlu"},
		{"lightBlu_"},
		{"_lightBlue"},
		{"lightBlue_"},
		{"_lightBlue4"},
		{"lightBlue4_"},
		{"_lightBlue40"},
		{"lightBlue40_"},
		{"_lightBlue400"},
		{"lightBlue400_"},
		{"9713l"},
		{"l9713"},
		{"9713li"},
		{"li9713"},
		{"9713lig"},
		{"lig9713"},
		{"9713ligh"},
		{"ligh9713"},
		{"9713light"},
		{"light9713"},
		{"9713lightB"},
		{"lightB9713"},
		{"9713lightBl"},
		{"lightBl9713"},
		{"9713lightBlu"},
		{"lightBlu9713"},
		{"9713lightBlue"},
		{"lightBlue9713"},
		{"9713lightBlue4"},
		{"lightBlue49713"},
		{"9713lightBlue40"},
		{"lightBlue409713"},
		{"9713lightBlue400"},
		{"lightBlue4009713"},
		{"Zl"},
		{"lZ"},
		{"Zli"},
		{"liZ"},
		{"Zlig"},
		{"ligZ"},
		{"Zligh"},
		{"lighZ"},
		{"Zlight"},
		{"lightZ"},
		{"ZlightB"},
		{"lightBZ"},
		{"ZlightBl"},
		{"lightBlZ"},
		{"ZlightBlu"},
		{"lightBluZ"},
		{"ZlightBlue"},
		{"lightBlueZ"},
		{"ZlightBlue4"},
		{"lightBlue4Z"},
		{"ZlightBlue40"},
		{"lightBlue40Z"},
		{"ZlightBlue400"},
		{"lightBlue400Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.LightBlue400()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestLightBlue500Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"lightBlue500", 12},
		{"lightBlue500 ", 12},
		{"lightBlue500\n", 12},
		{"lightBlue500.", 12},
		{"lightBlue500:", 12},
		{"lightBlue500,", 12},
		{"lightBlue500\"", 12},
		{"lightBlue500(", 12},
		{"lightBlue500)", 12},
		{"lightBlue500[", 12},
		{"lightBlue500]", 12},
		{"lightBlue500// comment", 12},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.LightBlue500()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestLightBlue500Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_l"},
		{"l_"},
		{"_li"},
		{"li_"},
		{"_lig"},
		{"lig_"},
		{"_ligh"},
		{"ligh_"},
		{"_light"},
		{"light_"},
		{"_lightB"},
		{"lightB_"},
		{"_lightBl"},
		{"lightBl_"},
		{"_lightBlu"},
		{"lightBlu_"},
		{"_lightBlue"},
		{"lightBlue_"},
		{"_lightBlue5"},
		{"lightBlue5_"},
		{"_lightBlue50"},
		{"lightBlue50_"},
		{"_lightBlue500"},
		{"lightBlue500_"},
		{"9713l"},
		{"l9713"},
		{"9713li"},
		{"li9713"},
		{"9713lig"},
		{"lig9713"},
		{"9713ligh"},
		{"ligh9713"},
		{"9713light"},
		{"light9713"},
		{"9713lightB"},
		{"lightB9713"},
		{"9713lightBl"},
		{"lightBl9713"},
		{"9713lightBlu"},
		{"lightBlu9713"},
		{"9713lightBlue"},
		{"lightBlue9713"},
		{"9713lightBlue5"},
		{"lightBlue59713"},
		{"9713lightBlue50"},
		{"lightBlue509713"},
		{"9713lightBlue500"},
		{"lightBlue5009713"},
		{"Zl"},
		{"lZ"},
		{"Zli"},
		{"liZ"},
		{"Zlig"},
		{"ligZ"},
		{"Zligh"},
		{"lighZ"},
		{"Zlight"},
		{"lightZ"},
		{"ZlightB"},
		{"lightBZ"},
		{"ZlightBl"},
		{"lightBlZ"},
		{"ZlightBlu"},
		{"lightBluZ"},
		{"ZlightBlue"},
		{"lightBlueZ"},
		{"ZlightBlue5"},
		{"lightBlue5Z"},
		{"ZlightBlue50"},
		{"lightBlue50Z"},
		{"ZlightBlue500"},
		{"lightBlue500Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.LightBlue500()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestLightBlue600Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"lightBlue600", 12},
		{"lightBlue600 ", 12},
		{"lightBlue600\n", 12},
		{"lightBlue600.", 12},
		{"lightBlue600:", 12},
		{"lightBlue600,", 12},
		{"lightBlue600\"", 12},
		{"lightBlue600(", 12},
		{"lightBlue600)", 12},
		{"lightBlue600[", 12},
		{"lightBlue600]", 12},
		{"lightBlue600// comment", 12},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.LightBlue600()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestLightBlue600Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_l"},
		{"l_"},
		{"_li"},
		{"li_"},
		{"_lig"},
		{"lig_"},
		{"_ligh"},
		{"ligh_"},
		{"_light"},
		{"light_"},
		{"_lightB"},
		{"lightB_"},
		{"_lightBl"},
		{"lightBl_"},
		{"_lightBlu"},
		{"lightBlu_"},
		{"_lightBlue"},
		{"lightBlue_"},
		{"_lightBlue6"},
		{"lightBlue6_"},
		{"_lightBlue60"},
		{"lightBlue60_"},
		{"_lightBlue600"},
		{"lightBlue600_"},
		{"9713l"},
		{"l9713"},
		{"9713li"},
		{"li9713"},
		{"9713lig"},
		{"lig9713"},
		{"9713ligh"},
		{"ligh9713"},
		{"9713light"},
		{"light9713"},
		{"9713lightB"},
		{"lightB9713"},
		{"9713lightBl"},
		{"lightBl9713"},
		{"9713lightBlu"},
		{"lightBlu9713"},
		{"9713lightBlue"},
		{"lightBlue9713"},
		{"9713lightBlue6"},
		{"lightBlue69713"},
		{"9713lightBlue60"},
		{"lightBlue609713"},
		{"9713lightBlue600"},
		{"lightBlue6009713"},
		{"Zl"},
		{"lZ"},
		{"Zli"},
		{"liZ"},
		{"Zlig"},
		{"ligZ"},
		{"Zligh"},
		{"lighZ"},
		{"Zlight"},
		{"lightZ"},
		{"ZlightB"},
		{"lightBZ"},
		{"ZlightBl"},
		{"lightBlZ"},
		{"ZlightBlu"},
		{"lightBluZ"},
		{"ZlightBlue"},
		{"lightBlueZ"},
		{"ZlightBlue6"},
		{"lightBlue6Z"},
		{"ZlightBlue60"},
		{"lightBlue60Z"},
		{"ZlightBlue600"},
		{"lightBlue600Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.LightBlue600()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestLightBlue700Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"lightBlue700", 12},
		{"lightBlue700 ", 12},
		{"lightBlue700\n", 12},
		{"lightBlue700.", 12},
		{"lightBlue700:", 12},
		{"lightBlue700,", 12},
		{"lightBlue700\"", 12},
		{"lightBlue700(", 12},
		{"lightBlue700)", 12},
		{"lightBlue700[", 12},
		{"lightBlue700]", 12},
		{"lightBlue700// comment", 12},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.LightBlue700()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestLightBlue700Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_l"},
		{"l_"},
		{"_li"},
		{"li_"},
		{"_lig"},
		{"lig_"},
		{"_ligh"},
		{"ligh_"},
		{"_light"},
		{"light_"},
		{"_lightB"},
		{"lightB_"},
		{"_lightBl"},
		{"lightBl_"},
		{"_lightBlu"},
		{"lightBlu_"},
		{"_lightBlue"},
		{"lightBlue_"},
		{"_lightBlue7"},
		{"lightBlue7_"},
		{"_lightBlue70"},
		{"lightBlue70_"},
		{"_lightBlue700"},
		{"lightBlue700_"},
		{"9713l"},
		{"l9713"},
		{"9713li"},
		{"li9713"},
		{"9713lig"},
		{"lig9713"},
		{"9713ligh"},
		{"ligh9713"},
		{"9713light"},
		{"light9713"},
		{"9713lightB"},
		{"lightB9713"},
		{"9713lightBl"},
		{"lightBl9713"},
		{"9713lightBlu"},
		{"lightBlu9713"},
		{"9713lightBlue"},
		{"lightBlue9713"},
		{"9713lightBlue7"},
		{"lightBlue79713"},
		{"9713lightBlue70"},
		{"lightBlue709713"},
		{"9713lightBlue700"},
		{"lightBlue7009713"},
		{"Zl"},
		{"lZ"},
		{"Zli"},
		{"liZ"},
		{"Zlig"},
		{"ligZ"},
		{"Zligh"},
		{"lighZ"},
		{"Zlight"},
		{"lightZ"},
		{"ZlightB"},
		{"lightBZ"},
		{"ZlightBl"},
		{"lightBlZ"},
		{"ZlightBlu"},
		{"lightBluZ"},
		{"ZlightBlue"},
		{"lightBlueZ"},
		{"ZlightBlue7"},
		{"lightBlue7Z"},
		{"ZlightBlue70"},
		{"lightBlue70Z"},
		{"ZlightBlue700"},
		{"lightBlue700Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.LightBlue700()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestLightBlue800Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"lightBlue800", 12},
		{"lightBlue800 ", 12},
		{"lightBlue800\n", 12},
		{"lightBlue800.", 12},
		{"lightBlue800:", 12},
		{"lightBlue800,", 12},
		{"lightBlue800\"", 12},
		{"lightBlue800(", 12},
		{"lightBlue800)", 12},
		{"lightBlue800[", 12},
		{"lightBlue800]", 12},
		{"lightBlue800// comment", 12},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.LightBlue800()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestLightBlue800Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_l"},
		{"l_"},
		{"_li"},
		{"li_"},
		{"_lig"},
		{"lig_"},
		{"_ligh"},
		{"ligh_"},
		{"_light"},
		{"light_"},
		{"_lightB"},
		{"lightB_"},
		{"_lightBl"},
		{"lightBl_"},
		{"_lightBlu"},
		{"lightBlu_"},
		{"_lightBlue"},
		{"lightBlue_"},
		{"_lightBlue8"},
		{"lightBlue8_"},
		{"_lightBlue80"},
		{"lightBlue80_"},
		{"_lightBlue800"},
		{"lightBlue800_"},
		{"9713l"},
		{"l9713"},
		{"9713li"},
		{"li9713"},
		{"9713lig"},
		{"lig9713"},
		{"9713ligh"},
		{"ligh9713"},
		{"9713light"},
		{"light9713"},
		{"9713lightB"},
		{"lightB9713"},
		{"9713lightBl"},
		{"lightBl9713"},
		{"9713lightBlu"},
		{"lightBlu9713"},
		{"9713lightBlue"},
		{"lightBlue9713"},
		{"9713lightBlue8"},
		{"lightBlue89713"},
		{"9713lightBlue80"},
		{"lightBlue809713"},
		{"9713lightBlue800"},
		{"lightBlue8009713"},
		{"Zl"},
		{"lZ"},
		{"Zli"},
		{"liZ"},
		{"Zlig"},
		{"ligZ"},
		{"Zligh"},
		{"lighZ"},
		{"Zlight"},
		{"lightZ"},
		{"ZlightB"},
		{"lightBZ"},
		{"ZlightBl"},
		{"lightBlZ"},
		{"ZlightBlu"},
		{"lightBluZ"},
		{"ZlightBlue"},
		{"lightBlueZ"},
		{"ZlightBlue8"},
		{"lightBlue8Z"},
		{"ZlightBlue80"},
		{"lightBlue80Z"},
		{"ZlightBlue800"},
		{"lightBlue800Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.LightBlue800()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestLightBlue900Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"lightBlue900", 12},
		{"lightBlue900 ", 12},
		{"lightBlue900\n", 12},
		{"lightBlue900.", 12},
		{"lightBlue900:", 12},
		{"lightBlue900,", 12},
		{"lightBlue900\"", 12},
		{"lightBlue900(", 12},
		{"lightBlue900)", 12},
		{"lightBlue900[", 12},
		{"lightBlue900]", 12},
		{"lightBlue900// comment", 12},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.LightBlue900()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestLightBlue900Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_l"},
		{"l_"},
		{"_li"},
		{"li_"},
		{"_lig"},
		{"lig_"},
		{"_ligh"},
		{"ligh_"},
		{"_light"},
		{"light_"},
		{"_lightB"},
		{"lightB_"},
		{"_lightBl"},
		{"lightBl_"},
		{"_lightBlu"},
		{"lightBlu_"},
		{"_lightBlue"},
		{"lightBlue_"},
		{"_lightBlue9"},
		{"lightBlue9_"},
		{"_lightBlue90"},
		{"lightBlue90_"},
		{"_lightBlue900"},
		{"lightBlue900_"},
		{"9713l"},
		{"l9713"},
		{"9713li"},
		{"li9713"},
		{"9713lig"},
		{"lig9713"},
		{"9713ligh"},
		{"ligh9713"},
		{"9713light"},
		{"light9713"},
		{"9713lightB"},
		{"lightB9713"},
		{"9713lightBl"},
		{"lightBl9713"},
		{"9713lightBlu"},
		{"lightBlu9713"},
		{"9713lightBlue"},
		{"lightBlue9713"},
		{"9713lightBlue9"},
		{"lightBlue99713"},
		{"9713lightBlue90"},
		{"lightBlue909713"},
		{"9713lightBlue900"},
		{"lightBlue9009713"},
		{"Zl"},
		{"lZ"},
		{"Zli"},
		{"liZ"},
		{"Zlig"},
		{"ligZ"},
		{"Zligh"},
		{"lighZ"},
		{"Zlight"},
		{"lightZ"},
		{"ZlightB"},
		{"lightBZ"},
		{"ZlightBl"},
		{"lightBlZ"},
		{"ZlightBlu"},
		{"lightBluZ"},
		{"ZlightBlue"},
		{"lightBlueZ"},
		{"ZlightBlue9"},
		{"lightBlue9Z"},
		{"ZlightBlue90"},
		{"lightBlue90Z"},
		{"ZlightBlue900"},
		{"lightBlue900Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.LightBlue900()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestLightBlueA100Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"lightBlueA100", 13},
		{"lightBlueA100 ", 13},
		{"lightBlueA100\n", 13},
		{"lightBlueA100.", 13},
		{"lightBlueA100:", 13},
		{"lightBlueA100,", 13},
		{"lightBlueA100\"", 13},
		{"lightBlueA100(", 13},
		{"lightBlueA100)", 13},
		{"lightBlueA100[", 13},
		{"lightBlueA100]", 13},
		{"lightBlueA100// comment", 13},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.LightBlueA100()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestLightBlueA100Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_l"},
		{"l_"},
		{"_li"},
		{"li_"},
		{"_lig"},
		{"lig_"},
		{"_ligh"},
		{"ligh_"},
		{"_light"},
		{"light_"},
		{"_lightB"},
		{"lightB_"},
		{"_lightBl"},
		{"lightBl_"},
		{"_lightBlu"},
		{"lightBlu_"},
		{"_lightBlue"},
		{"lightBlue_"},
		{"_lightBlueA"},
		{"lightBlueA_"},
		{"_lightBlueA1"},
		{"lightBlueA1_"},
		{"_lightBlueA10"},
		{"lightBlueA10_"},
		{"_lightBlueA100"},
		{"lightBlueA100_"},
		{"9713l"},
		{"l9713"},
		{"9713li"},
		{"li9713"},
		{"9713lig"},
		{"lig9713"},
		{"9713ligh"},
		{"ligh9713"},
		{"9713light"},
		{"light9713"},
		{"9713lightB"},
		{"lightB9713"},
		{"9713lightBl"},
		{"lightBl9713"},
		{"9713lightBlu"},
		{"lightBlu9713"},
		{"9713lightBlue"},
		{"lightBlue9713"},
		{"9713lightBlueA"},
		{"lightBlueA9713"},
		{"9713lightBlueA1"},
		{"lightBlueA19713"},
		{"9713lightBlueA10"},
		{"lightBlueA109713"},
		{"9713lightBlueA100"},
		{"lightBlueA1009713"},
		{"Zl"},
		{"lZ"},
		{"Zli"},
		{"liZ"},
		{"Zlig"},
		{"ligZ"},
		{"Zligh"},
		{"lighZ"},
		{"Zlight"},
		{"lightZ"},
		{"ZlightB"},
		{"lightBZ"},
		{"ZlightBl"},
		{"lightBlZ"},
		{"ZlightBlu"},
		{"lightBluZ"},
		{"ZlightBlue"},
		{"lightBlueZ"},
		{"ZlightBlueA"},
		{"lightBlueAZ"},
		{"ZlightBlueA1"},
		{"lightBlueA1Z"},
		{"ZlightBlueA10"},
		{"lightBlueA10Z"},
		{"ZlightBlueA100"},
		{"lightBlueA100Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.LightBlueA100()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestLightBlueA200Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"lightBlueA200", 13},
		{"lightBlueA200 ", 13},
		{"lightBlueA200\n", 13},
		{"lightBlueA200.", 13},
		{"lightBlueA200:", 13},
		{"lightBlueA200,", 13},
		{"lightBlueA200\"", 13},
		{"lightBlueA200(", 13},
		{"lightBlueA200)", 13},
		{"lightBlueA200[", 13},
		{"lightBlueA200]", 13},
		{"lightBlueA200// comment", 13},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.LightBlueA200()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestLightBlueA200Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_l"},
		{"l_"},
		{"_li"},
		{"li_"},
		{"_lig"},
		{"lig_"},
		{"_ligh"},
		{"ligh_"},
		{"_light"},
		{"light_"},
		{"_lightB"},
		{"lightB_"},
		{"_lightBl"},
		{"lightBl_"},
		{"_lightBlu"},
		{"lightBlu_"},
		{"_lightBlue"},
		{"lightBlue_"},
		{"_lightBlueA"},
		{"lightBlueA_"},
		{"_lightBlueA2"},
		{"lightBlueA2_"},
		{"_lightBlueA20"},
		{"lightBlueA20_"},
		{"_lightBlueA200"},
		{"lightBlueA200_"},
		{"9713l"},
		{"l9713"},
		{"9713li"},
		{"li9713"},
		{"9713lig"},
		{"lig9713"},
		{"9713ligh"},
		{"ligh9713"},
		{"9713light"},
		{"light9713"},
		{"9713lightB"},
		{"lightB9713"},
		{"9713lightBl"},
		{"lightBl9713"},
		{"9713lightBlu"},
		{"lightBlu9713"},
		{"9713lightBlue"},
		{"lightBlue9713"},
		{"9713lightBlueA"},
		{"lightBlueA9713"},
		{"9713lightBlueA2"},
		{"lightBlueA29713"},
		{"9713lightBlueA20"},
		{"lightBlueA209713"},
		{"9713lightBlueA200"},
		{"lightBlueA2009713"},
		{"Zl"},
		{"lZ"},
		{"Zli"},
		{"liZ"},
		{"Zlig"},
		{"ligZ"},
		{"Zligh"},
		{"lighZ"},
		{"Zlight"},
		{"lightZ"},
		{"ZlightB"},
		{"lightBZ"},
		{"ZlightBl"},
		{"lightBlZ"},
		{"ZlightBlu"},
		{"lightBluZ"},
		{"ZlightBlue"},
		{"lightBlueZ"},
		{"ZlightBlueA"},
		{"lightBlueAZ"},
		{"ZlightBlueA2"},
		{"lightBlueA2Z"},
		{"ZlightBlueA20"},
		{"lightBlueA20Z"},
		{"ZlightBlueA200"},
		{"lightBlueA200Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.LightBlueA200()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestLightBlueA400Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"lightBlueA400", 13},
		{"lightBlueA400 ", 13},
		{"lightBlueA400\n", 13},
		{"lightBlueA400.", 13},
		{"lightBlueA400:", 13},
		{"lightBlueA400,", 13},
		{"lightBlueA400\"", 13},
		{"lightBlueA400(", 13},
		{"lightBlueA400)", 13},
		{"lightBlueA400[", 13},
		{"lightBlueA400]", 13},
		{"lightBlueA400// comment", 13},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.LightBlueA400()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestLightBlueA400Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_l"},
		{"l_"},
		{"_li"},
		{"li_"},
		{"_lig"},
		{"lig_"},
		{"_ligh"},
		{"ligh_"},
		{"_light"},
		{"light_"},
		{"_lightB"},
		{"lightB_"},
		{"_lightBl"},
		{"lightBl_"},
		{"_lightBlu"},
		{"lightBlu_"},
		{"_lightBlue"},
		{"lightBlue_"},
		{"_lightBlueA"},
		{"lightBlueA_"},
		{"_lightBlueA4"},
		{"lightBlueA4_"},
		{"_lightBlueA40"},
		{"lightBlueA40_"},
		{"_lightBlueA400"},
		{"lightBlueA400_"},
		{"9713l"},
		{"l9713"},
		{"9713li"},
		{"li9713"},
		{"9713lig"},
		{"lig9713"},
		{"9713ligh"},
		{"ligh9713"},
		{"9713light"},
		{"light9713"},
		{"9713lightB"},
		{"lightB9713"},
		{"9713lightBl"},
		{"lightBl9713"},
		{"9713lightBlu"},
		{"lightBlu9713"},
		{"9713lightBlue"},
		{"lightBlue9713"},
		{"9713lightBlueA"},
		{"lightBlueA9713"},
		{"9713lightBlueA4"},
		{"lightBlueA49713"},
		{"9713lightBlueA40"},
		{"lightBlueA409713"},
		{"9713lightBlueA400"},
		{"lightBlueA4009713"},
		{"Zl"},
		{"lZ"},
		{"Zli"},
		{"liZ"},
		{"Zlig"},
		{"ligZ"},
		{"Zligh"},
		{"lighZ"},
		{"Zlight"},
		{"lightZ"},
		{"ZlightB"},
		{"lightBZ"},
		{"ZlightBl"},
		{"lightBlZ"},
		{"ZlightBlu"},
		{"lightBluZ"},
		{"ZlightBlue"},
		{"lightBlueZ"},
		{"ZlightBlueA"},
		{"lightBlueAZ"},
		{"ZlightBlueA4"},
		{"lightBlueA4Z"},
		{"ZlightBlueA40"},
		{"lightBlueA40Z"},
		{"ZlightBlueA400"},
		{"lightBlueA400Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.LightBlueA400()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestLightBlueA700Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"lightBlueA700", 13},
		{"lightBlueA700 ", 13},
		{"lightBlueA700\n", 13},
		{"lightBlueA700.", 13},
		{"lightBlueA700:", 13},
		{"lightBlueA700,", 13},
		{"lightBlueA700\"", 13},
		{"lightBlueA700(", 13},
		{"lightBlueA700)", 13},
		{"lightBlueA700[", 13},
		{"lightBlueA700]", 13},
		{"lightBlueA700// comment", 13},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.LightBlueA700()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestLightBlueA700Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_l"},
		{"l_"},
		{"_li"},
		{"li_"},
		{"_lig"},
		{"lig_"},
		{"_ligh"},
		{"ligh_"},
		{"_light"},
		{"light_"},
		{"_lightB"},
		{"lightB_"},
		{"_lightBl"},
		{"lightBl_"},
		{"_lightBlu"},
		{"lightBlu_"},
		{"_lightBlue"},
		{"lightBlue_"},
		{"_lightBlueA"},
		{"lightBlueA_"},
		{"_lightBlueA7"},
		{"lightBlueA7_"},
		{"_lightBlueA70"},
		{"lightBlueA70_"},
		{"_lightBlueA700"},
		{"lightBlueA700_"},
		{"9713l"},
		{"l9713"},
		{"9713li"},
		{"li9713"},
		{"9713lig"},
		{"lig9713"},
		{"9713ligh"},
		{"ligh9713"},
		{"9713light"},
		{"light9713"},
		{"9713lightB"},
		{"lightB9713"},
		{"9713lightBl"},
		{"lightBl9713"},
		{"9713lightBlu"},
		{"lightBlu9713"},
		{"9713lightBlue"},
		{"lightBlue9713"},
		{"9713lightBlueA"},
		{"lightBlueA9713"},
		{"9713lightBlueA7"},
		{"lightBlueA79713"},
		{"9713lightBlueA70"},
		{"lightBlueA709713"},
		{"9713lightBlueA700"},
		{"lightBlueA7009713"},
		{"Zl"},
		{"lZ"},
		{"Zli"},
		{"liZ"},
		{"Zlig"},
		{"ligZ"},
		{"Zligh"},
		{"lighZ"},
		{"Zlight"},
		{"lightZ"},
		{"ZlightB"},
		{"lightBZ"},
		{"ZlightBl"},
		{"lightBlZ"},
		{"ZlightBlu"},
		{"lightBluZ"},
		{"ZlightBlue"},
		{"lightBlueZ"},
		{"ZlightBlueA"},
		{"lightBlueAZ"},
		{"ZlightBlueA7"},
		{"lightBlueA7Z"},
		{"ZlightBlueA70"},
		{"lightBlueA70Z"},
		{"ZlightBlueA700"},
		{"lightBlueA700Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.LightBlueA700()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestCyan50Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"cyan50", 6},
		{"cyan50 ", 6},
		{"cyan50\n", 6},
		{"cyan50.", 6},
		{"cyan50:", 6},
		{"cyan50,", 6},
		{"cyan50\"", 6},
		{"cyan50(", 6},
		{"cyan50)", 6},
		{"cyan50[", 6},
		{"cyan50]", 6},
		{"cyan50// comment", 6},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Cyan50()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestCyan50Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_c"},
		{"c_"},
		{"_cy"},
		{"cy_"},
		{"_cya"},
		{"cya_"},
		{"_cyan"},
		{"cyan_"},
		{"_cyan5"},
		{"cyan5_"},
		{"_cyan50"},
		{"cyan50_"},
		{"9713c"},
		{"c9713"},
		{"9713cy"},
		{"cy9713"},
		{"9713cya"},
		{"cya9713"},
		{"9713cyan"},
		{"cyan9713"},
		{"9713cyan5"},
		{"cyan59713"},
		{"9713cyan50"},
		{"cyan509713"},
		{"Zc"},
		{"cZ"},
		{"Zcy"},
		{"cyZ"},
		{"Zcya"},
		{"cyaZ"},
		{"Zcyan"},
		{"cyanZ"},
		{"Zcyan5"},
		{"cyan5Z"},
		{"Zcyan50"},
		{"cyan50Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Cyan50()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestCyan100Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"cyan100", 7},
		{"cyan100 ", 7},
		{"cyan100\n", 7},
		{"cyan100.", 7},
		{"cyan100:", 7},
		{"cyan100,", 7},
		{"cyan100\"", 7},
		{"cyan100(", 7},
		{"cyan100)", 7},
		{"cyan100[", 7},
		{"cyan100]", 7},
		{"cyan100// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Cyan100()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestCyan100Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_c"},
		{"c_"},
		{"_cy"},
		{"cy_"},
		{"_cya"},
		{"cya_"},
		{"_cyan"},
		{"cyan_"},
		{"_cyan1"},
		{"cyan1_"},
		{"_cyan10"},
		{"cyan10_"},
		{"_cyan100"},
		{"cyan100_"},
		{"9713c"},
		{"c9713"},
		{"9713cy"},
		{"cy9713"},
		{"9713cya"},
		{"cya9713"},
		{"9713cyan"},
		{"cyan9713"},
		{"9713cyan1"},
		{"cyan19713"},
		{"9713cyan10"},
		{"cyan109713"},
		{"9713cyan100"},
		{"cyan1009713"},
		{"Zc"},
		{"cZ"},
		{"Zcy"},
		{"cyZ"},
		{"Zcya"},
		{"cyaZ"},
		{"Zcyan"},
		{"cyanZ"},
		{"Zcyan1"},
		{"cyan1Z"},
		{"Zcyan10"},
		{"cyan10Z"},
		{"Zcyan100"},
		{"cyan100Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Cyan100()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestCyan200Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"cyan200", 7},
		{"cyan200 ", 7},
		{"cyan200\n", 7},
		{"cyan200.", 7},
		{"cyan200:", 7},
		{"cyan200,", 7},
		{"cyan200\"", 7},
		{"cyan200(", 7},
		{"cyan200)", 7},
		{"cyan200[", 7},
		{"cyan200]", 7},
		{"cyan200// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Cyan200()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestCyan200Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_c"},
		{"c_"},
		{"_cy"},
		{"cy_"},
		{"_cya"},
		{"cya_"},
		{"_cyan"},
		{"cyan_"},
		{"_cyan2"},
		{"cyan2_"},
		{"_cyan20"},
		{"cyan20_"},
		{"_cyan200"},
		{"cyan200_"},
		{"9713c"},
		{"c9713"},
		{"9713cy"},
		{"cy9713"},
		{"9713cya"},
		{"cya9713"},
		{"9713cyan"},
		{"cyan9713"},
		{"9713cyan2"},
		{"cyan29713"},
		{"9713cyan20"},
		{"cyan209713"},
		{"9713cyan200"},
		{"cyan2009713"},
		{"Zc"},
		{"cZ"},
		{"Zcy"},
		{"cyZ"},
		{"Zcya"},
		{"cyaZ"},
		{"Zcyan"},
		{"cyanZ"},
		{"Zcyan2"},
		{"cyan2Z"},
		{"Zcyan20"},
		{"cyan20Z"},
		{"Zcyan200"},
		{"cyan200Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Cyan200()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestCyan300Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"cyan300", 7},
		{"cyan300 ", 7},
		{"cyan300\n", 7},
		{"cyan300.", 7},
		{"cyan300:", 7},
		{"cyan300,", 7},
		{"cyan300\"", 7},
		{"cyan300(", 7},
		{"cyan300)", 7},
		{"cyan300[", 7},
		{"cyan300]", 7},
		{"cyan300// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Cyan300()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestCyan300Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_c"},
		{"c_"},
		{"_cy"},
		{"cy_"},
		{"_cya"},
		{"cya_"},
		{"_cyan"},
		{"cyan_"},
		{"_cyan3"},
		{"cyan3_"},
		{"_cyan30"},
		{"cyan30_"},
		{"_cyan300"},
		{"cyan300_"},
		{"9713c"},
		{"c9713"},
		{"9713cy"},
		{"cy9713"},
		{"9713cya"},
		{"cya9713"},
		{"9713cyan"},
		{"cyan9713"},
		{"9713cyan3"},
		{"cyan39713"},
		{"9713cyan30"},
		{"cyan309713"},
		{"9713cyan300"},
		{"cyan3009713"},
		{"Zc"},
		{"cZ"},
		{"Zcy"},
		{"cyZ"},
		{"Zcya"},
		{"cyaZ"},
		{"Zcyan"},
		{"cyanZ"},
		{"Zcyan3"},
		{"cyan3Z"},
		{"Zcyan30"},
		{"cyan30Z"},
		{"Zcyan300"},
		{"cyan300Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Cyan300()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestCyan400Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"cyan400", 7},
		{"cyan400 ", 7},
		{"cyan400\n", 7},
		{"cyan400.", 7},
		{"cyan400:", 7},
		{"cyan400,", 7},
		{"cyan400\"", 7},
		{"cyan400(", 7},
		{"cyan400)", 7},
		{"cyan400[", 7},
		{"cyan400]", 7},
		{"cyan400// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Cyan400()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestCyan400Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_c"},
		{"c_"},
		{"_cy"},
		{"cy_"},
		{"_cya"},
		{"cya_"},
		{"_cyan"},
		{"cyan_"},
		{"_cyan4"},
		{"cyan4_"},
		{"_cyan40"},
		{"cyan40_"},
		{"_cyan400"},
		{"cyan400_"},
		{"9713c"},
		{"c9713"},
		{"9713cy"},
		{"cy9713"},
		{"9713cya"},
		{"cya9713"},
		{"9713cyan"},
		{"cyan9713"},
		{"9713cyan4"},
		{"cyan49713"},
		{"9713cyan40"},
		{"cyan409713"},
		{"9713cyan400"},
		{"cyan4009713"},
		{"Zc"},
		{"cZ"},
		{"Zcy"},
		{"cyZ"},
		{"Zcya"},
		{"cyaZ"},
		{"Zcyan"},
		{"cyanZ"},
		{"Zcyan4"},
		{"cyan4Z"},
		{"Zcyan40"},
		{"cyan40Z"},
		{"Zcyan400"},
		{"cyan400Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Cyan400()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestCyan500Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"cyan500", 7},
		{"cyan500 ", 7},
		{"cyan500\n", 7},
		{"cyan500.", 7},
		{"cyan500:", 7},
		{"cyan500,", 7},
		{"cyan500\"", 7},
		{"cyan500(", 7},
		{"cyan500)", 7},
		{"cyan500[", 7},
		{"cyan500]", 7},
		{"cyan500// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Cyan500()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestCyan500Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_c"},
		{"c_"},
		{"_cy"},
		{"cy_"},
		{"_cya"},
		{"cya_"},
		{"_cyan"},
		{"cyan_"},
		{"_cyan5"},
		{"cyan5_"},
		{"_cyan50"},
		{"cyan50_"},
		{"_cyan500"},
		{"cyan500_"},
		{"9713c"},
		{"c9713"},
		{"9713cy"},
		{"cy9713"},
		{"9713cya"},
		{"cya9713"},
		{"9713cyan"},
		{"cyan9713"},
		{"9713cyan5"},
		{"cyan59713"},
		{"9713cyan50"},
		{"cyan509713"},
		{"9713cyan500"},
		{"cyan5009713"},
		{"Zc"},
		{"cZ"},
		{"Zcy"},
		{"cyZ"},
		{"Zcya"},
		{"cyaZ"},
		{"Zcyan"},
		{"cyanZ"},
		{"Zcyan5"},
		{"cyan5Z"},
		{"Zcyan50"},
		{"cyan50Z"},
		{"Zcyan500"},
		{"cyan500Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Cyan500()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestCyan600Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"cyan600", 7},
		{"cyan600 ", 7},
		{"cyan600\n", 7},
		{"cyan600.", 7},
		{"cyan600:", 7},
		{"cyan600,", 7},
		{"cyan600\"", 7},
		{"cyan600(", 7},
		{"cyan600)", 7},
		{"cyan600[", 7},
		{"cyan600]", 7},
		{"cyan600// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Cyan600()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestCyan600Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_c"},
		{"c_"},
		{"_cy"},
		{"cy_"},
		{"_cya"},
		{"cya_"},
		{"_cyan"},
		{"cyan_"},
		{"_cyan6"},
		{"cyan6_"},
		{"_cyan60"},
		{"cyan60_"},
		{"_cyan600"},
		{"cyan600_"},
		{"9713c"},
		{"c9713"},
		{"9713cy"},
		{"cy9713"},
		{"9713cya"},
		{"cya9713"},
		{"9713cyan"},
		{"cyan9713"},
		{"9713cyan6"},
		{"cyan69713"},
		{"9713cyan60"},
		{"cyan609713"},
		{"9713cyan600"},
		{"cyan6009713"},
		{"Zc"},
		{"cZ"},
		{"Zcy"},
		{"cyZ"},
		{"Zcya"},
		{"cyaZ"},
		{"Zcyan"},
		{"cyanZ"},
		{"Zcyan6"},
		{"cyan6Z"},
		{"Zcyan60"},
		{"cyan60Z"},
		{"Zcyan600"},
		{"cyan600Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Cyan600()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestCyan700Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"cyan700", 7},
		{"cyan700 ", 7},
		{"cyan700\n", 7},
		{"cyan700.", 7},
		{"cyan700:", 7},
		{"cyan700,", 7},
		{"cyan700\"", 7},
		{"cyan700(", 7},
		{"cyan700)", 7},
		{"cyan700[", 7},
		{"cyan700]", 7},
		{"cyan700// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Cyan700()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestCyan700Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_c"},
		{"c_"},
		{"_cy"},
		{"cy_"},
		{"_cya"},
		{"cya_"},
		{"_cyan"},
		{"cyan_"},
		{"_cyan7"},
		{"cyan7_"},
		{"_cyan70"},
		{"cyan70_"},
		{"_cyan700"},
		{"cyan700_"},
		{"9713c"},
		{"c9713"},
		{"9713cy"},
		{"cy9713"},
		{"9713cya"},
		{"cya9713"},
		{"9713cyan"},
		{"cyan9713"},
		{"9713cyan7"},
		{"cyan79713"},
		{"9713cyan70"},
		{"cyan709713"},
		{"9713cyan700"},
		{"cyan7009713"},
		{"Zc"},
		{"cZ"},
		{"Zcy"},
		{"cyZ"},
		{"Zcya"},
		{"cyaZ"},
		{"Zcyan"},
		{"cyanZ"},
		{"Zcyan7"},
		{"cyan7Z"},
		{"Zcyan70"},
		{"cyan70Z"},
		{"Zcyan700"},
		{"cyan700Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Cyan700()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestCyan800Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"cyan800", 7},
		{"cyan800 ", 7},
		{"cyan800\n", 7},
		{"cyan800.", 7},
		{"cyan800:", 7},
		{"cyan800,", 7},
		{"cyan800\"", 7},
		{"cyan800(", 7},
		{"cyan800)", 7},
		{"cyan800[", 7},
		{"cyan800]", 7},
		{"cyan800// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Cyan800()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestCyan800Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_c"},
		{"c_"},
		{"_cy"},
		{"cy_"},
		{"_cya"},
		{"cya_"},
		{"_cyan"},
		{"cyan_"},
		{"_cyan8"},
		{"cyan8_"},
		{"_cyan80"},
		{"cyan80_"},
		{"_cyan800"},
		{"cyan800_"},
		{"9713c"},
		{"c9713"},
		{"9713cy"},
		{"cy9713"},
		{"9713cya"},
		{"cya9713"},
		{"9713cyan"},
		{"cyan9713"},
		{"9713cyan8"},
		{"cyan89713"},
		{"9713cyan80"},
		{"cyan809713"},
		{"9713cyan800"},
		{"cyan8009713"},
		{"Zc"},
		{"cZ"},
		{"Zcy"},
		{"cyZ"},
		{"Zcya"},
		{"cyaZ"},
		{"Zcyan"},
		{"cyanZ"},
		{"Zcyan8"},
		{"cyan8Z"},
		{"Zcyan80"},
		{"cyan80Z"},
		{"Zcyan800"},
		{"cyan800Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Cyan800()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestCyan900Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"cyan900", 7},
		{"cyan900 ", 7},
		{"cyan900\n", 7},
		{"cyan900.", 7},
		{"cyan900:", 7},
		{"cyan900,", 7},
		{"cyan900\"", 7},
		{"cyan900(", 7},
		{"cyan900)", 7},
		{"cyan900[", 7},
		{"cyan900]", 7},
		{"cyan900// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Cyan900()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestCyan900Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_c"},
		{"c_"},
		{"_cy"},
		{"cy_"},
		{"_cya"},
		{"cya_"},
		{"_cyan"},
		{"cyan_"},
		{"_cyan9"},
		{"cyan9_"},
		{"_cyan90"},
		{"cyan90_"},
		{"_cyan900"},
		{"cyan900_"},
		{"9713c"},
		{"c9713"},
		{"9713cy"},
		{"cy9713"},
		{"9713cya"},
		{"cya9713"},
		{"9713cyan"},
		{"cyan9713"},
		{"9713cyan9"},
		{"cyan99713"},
		{"9713cyan90"},
		{"cyan909713"},
		{"9713cyan900"},
		{"cyan9009713"},
		{"Zc"},
		{"cZ"},
		{"Zcy"},
		{"cyZ"},
		{"Zcya"},
		{"cyaZ"},
		{"Zcyan"},
		{"cyanZ"},
		{"Zcyan9"},
		{"cyan9Z"},
		{"Zcyan90"},
		{"cyan90Z"},
		{"Zcyan900"},
		{"cyan900Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Cyan900()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestCyanA100Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"cyanA100", 8},
		{"cyanA100 ", 8},
		{"cyanA100\n", 8},
		{"cyanA100.", 8},
		{"cyanA100:", 8},
		{"cyanA100,", 8},
		{"cyanA100\"", 8},
		{"cyanA100(", 8},
		{"cyanA100)", 8},
		{"cyanA100[", 8},
		{"cyanA100]", 8},
		{"cyanA100// comment", 8},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.CyanA100()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestCyanA100Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_c"},
		{"c_"},
		{"_cy"},
		{"cy_"},
		{"_cya"},
		{"cya_"},
		{"_cyan"},
		{"cyan_"},
		{"_cyanA"},
		{"cyanA_"},
		{"_cyanA1"},
		{"cyanA1_"},
		{"_cyanA10"},
		{"cyanA10_"},
		{"_cyanA100"},
		{"cyanA100_"},
		{"9713c"},
		{"c9713"},
		{"9713cy"},
		{"cy9713"},
		{"9713cya"},
		{"cya9713"},
		{"9713cyan"},
		{"cyan9713"},
		{"9713cyanA"},
		{"cyanA9713"},
		{"9713cyanA1"},
		{"cyanA19713"},
		{"9713cyanA10"},
		{"cyanA109713"},
		{"9713cyanA100"},
		{"cyanA1009713"},
		{"Zc"},
		{"cZ"},
		{"Zcy"},
		{"cyZ"},
		{"Zcya"},
		{"cyaZ"},
		{"Zcyan"},
		{"cyanZ"},
		{"ZcyanA"},
		{"cyanAZ"},
		{"ZcyanA1"},
		{"cyanA1Z"},
		{"ZcyanA10"},
		{"cyanA10Z"},
		{"ZcyanA100"},
		{"cyanA100Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.CyanA100()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestCyanA200Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"cyanA200", 8},
		{"cyanA200 ", 8},
		{"cyanA200\n", 8},
		{"cyanA200.", 8},
		{"cyanA200:", 8},
		{"cyanA200,", 8},
		{"cyanA200\"", 8},
		{"cyanA200(", 8},
		{"cyanA200)", 8},
		{"cyanA200[", 8},
		{"cyanA200]", 8},
		{"cyanA200// comment", 8},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.CyanA200()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestCyanA200Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_c"},
		{"c_"},
		{"_cy"},
		{"cy_"},
		{"_cya"},
		{"cya_"},
		{"_cyan"},
		{"cyan_"},
		{"_cyanA"},
		{"cyanA_"},
		{"_cyanA2"},
		{"cyanA2_"},
		{"_cyanA20"},
		{"cyanA20_"},
		{"_cyanA200"},
		{"cyanA200_"},
		{"9713c"},
		{"c9713"},
		{"9713cy"},
		{"cy9713"},
		{"9713cya"},
		{"cya9713"},
		{"9713cyan"},
		{"cyan9713"},
		{"9713cyanA"},
		{"cyanA9713"},
		{"9713cyanA2"},
		{"cyanA29713"},
		{"9713cyanA20"},
		{"cyanA209713"},
		{"9713cyanA200"},
		{"cyanA2009713"},
		{"Zc"},
		{"cZ"},
		{"Zcy"},
		{"cyZ"},
		{"Zcya"},
		{"cyaZ"},
		{"Zcyan"},
		{"cyanZ"},
		{"ZcyanA"},
		{"cyanAZ"},
		{"ZcyanA2"},
		{"cyanA2Z"},
		{"ZcyanA20"},
		{"cyanA20Z"},
		{"ZcyanA200"},
		{"cyanA200Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.CyanA200()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestCyanA400Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"cyanA400", 8},
		{"cyanA400 ", 8},
		{"cyanA400\n", 8},
		{"cyanA400.", 8},
		{"cyanA400:", 8},
		{"cyanA400,", 8},
		{"cyanA400\"", 8},
		{"cyanA400(", 8},
		{"cyanA400)", 8},
		{"cyanA400[", 8},
		{"cyanA400]", 8},
		{"cyanA400// comment", 8},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.CyanA400()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestCyanA400Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_c"},
		{"c_"},
		{"_cy"},
		{"cy_"},
		{"_cya"},
		{"cya_"},
		{"_cyan"},
		{"cyan_"},
		{"_cyanA"},
		{"cyanA_"},
		{"_cyanA4"},
		{"cyanA4_"},
		{"_cyanA40"},
		{"cyanA40_"},
		{"_cyanA400"},
		{"cyanA400_"},
		{"9713c"},
		{"c9713"},
		{"9713cy"},
		{"cy9713"},
		{"9713cya"},
		{"cya9713"},
		{"9713cyan"},
		{"cyan9713"},
		{"9713cyanA"},
		{"cyanA9713"},
		{"9713cyanA4"},
		{"cyanA49713"},
		{"9713cyanA40"},
		{"cyanA409713"},
		{"9713cyanA400"},
		{"cyanA4009713"},
		{"Zc"},
		{"cZ"},
		{"Zcy"},
		{"cyZ"},
		{"Zcya"},
		{"cyaZ"},
		{"Zcyan"},
		{"cyanZ"},
		{"ZcyanA"},
		{"cyanAZ"},
		{"ZcyanA4"},
		{"cyanA4Z"},
		{"ZcyanA40"},
		{"cyanA40Z"},
		{"ZcyanA400"},
		{"cyanA400Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.CyanA400()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestCyanA700Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"cyanA700", 8},
		{"cyanA700 ", 8},
		{"cyanA700\n", 8},
		{"cyanA700.", 8},
		{"cyanA700:", 8},
		{"cyanA700,", 8},
		{"cyanA700\"", 8},
		{"cyanA700(", 8},
		{"cyanA700)", 8},
		{"cyanA700[", 8},
		{"cyanA700]", 8},
		{"cyanA700// comment", 8},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.CyanA700()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestCyanA700Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_c"},
		{"c_"},
		{"_cy"},
		{"cy_"},
		{"_cya"},
		{"cya_"},
		{"_cyan"},
		{"cyan_"},
		{"_cyanA"},
		{"cyanA_"},
		{"_cyanA7"},
		{"cyanA7_"},
		{"_cyanA70"},
		{"cyanA70_"},
		{"_cyanA700"},
		{"cyanA700_"},
		{"9713c"},
		{"c9713"},
		{"9713cy"},
		{"cy9713"},
		{"9713cya"},
		{"cya9713"},
		{"9713cyan"},
		{"cyan9713"},
		{"9713cyanA"},
		{"cyanA9713"},
		{"9713cyanA7"},
		{"cyanA79713"},
		{"9713cyanA70"},
		{"cyanA709713"},
		{"9713cyanA700"},
		{"cyanA7009713"},
		{"Zc"},
		{"cZ"},
		{"Zcy"},
		{"cyZ"},
		{"Zcya"},
		{"cyaZ"},
		{"Zcyan"},
		{"cyanZ"},
		{"ZcyanA"},
		{"cyanAZ"},
		{"ZcyanA7"},
		{"cyanA7Z"},
		{"ZcyanA70"},
		{"cyanA70Z"},
		{"ZcyanA700"},
		{"cyanA700Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.CyanA700()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestTeal50Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"teal50", 6},
		{"teal50 ", 6},
		{"teal50\n", 6},
		{"teal50.", 6},
		{"teal50:", 6},
		{"teal50,", 6},
		{"teal50\"", 6},
		{"teal50(", 6},
		{"teal50)", 6},
		{"teal50[", 6},
		{"teal50]", 6},
		{"teal50// comment", 6},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Teal50()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestTeal50Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_t"},
		{"t_"},
		{"_te"},
		{"te_"},
		{"_tea"},
		{"tea_"},
		{"_teal"},
		{"teal_"},
		{"_teal5"},
		{"teal5_"},
		{"_teal50"},
		{"teal50_"},
		{"9713t"},
		{"t9713"},
		{"9713te"},
		{"te9713"},
		{"9713tea"},
		{"tea9713"},
		{"9713teal"},
		{"teal9713"},
		{"9713teal5"},
		{"teal59713"},
		{"9713teal50"},
		{"teal509713"},
		{"Zt"},
		{"tZ"},
		{"Zte"},
		{"teZ"},
		{"Ztea"},
		{"teaZ"},
		{"Zteal"},
		{"tealZ"},
		{"Zteal5"},
		{"teal5Z"},
		{"Zteal50"},
		{"teal50Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Teal50()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestTeal100Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"teal100", 7},
		{"teal100 ", 7},
		{"teal100\n", 7},
		{"teal100.", 7},
		{"teal100:", 7},
		{"teal100,", 7},
		{"teal100\"", 7},
		{"teal100(", 7},
		{"teal100)", 7},
		{"teal100[", 7},
		{"teal100]", 7},
		{"teal100// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Teal100()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestTeal100Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_t"},
		{"t_"},
		{"_te"},
		{"te_"},
		{"_tea"},
		{"tea_"},
		{"_teal"},
		{"teal_"},
		{"_teal1"},
		{"teal1_"},
		{"_teal10"},
		{"teal10_"},
		{"_teal100"},
		{"teal100_"},
		{"9713t"},
		{"t9713"},
		{"9713te"},
		{"te9713"},
		{"9713tea"},
		{"tea9713"},
		{"9713teal"},
		{"teal9713"},
		{"9713teal1"},
		{"teal19713"},
		{"9713teal10"},
		{"teal109713"},
		{"9713teal100"},
		{"teal1009713"},
		{"Zt"},
		{"tZ"},
		{"Zte"},
		{"teZ"},
		{"Ztea"},
		{"teaZ"},
		{"Zteal"},
		{"tealZ"},
		{"Zteal1"},
		{"teal1Z"},
		{"Zteal10"},
		{"teal10Z"},
		{"Zteal100"},
		{"teal100Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Teal100()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestTeal200Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"teal200", 7},
		{"teal200 ", 7},
		{"teal200\n", 7},
		{"teal200.", 7},
		{"teal200:", 7},
		{"teal200,", 7},
		{"teal200\"", 7},
		{"teal200(", 7},
		{"teal200)", 7},
		{"teal200[", 7},
		{"teal200]", 7},
		{"teal200// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Teal200()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestTeal200Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_t"},
		{"t_"},
		{"_te"},
		{"te_"},
		{"_tea"},
		{"tea_"},
		{"_teal"},
		{"teal_"},
		{"_teal2"},
		{"teal2_"},
		{"_teal20"},
		{"teal20_"},
		{"_teal200"},
		{"teal200_"},
		{"9713t"},
		{"t9713"},
		{"9713te"},
		{"te9713"},
		{"9713tea"},
		{"tea9713"},
		{"9713teal"},
		{"teal9713"},
		{"9713teal2"},
		{"teal29713"},
		{"9713teal20"},
		{"teal209713"},
		{"9713teal200"},
		{"teal2009713"},
		{"Zt"},
		{"tZ"},
		{"Zte"},
		{"teZ"},
		{"Ztea"},
		{"teaZ"},
		{"Zteal"},
		{"tealZ"},
		{"Zteal2"},
		{"teal2Z"},
		{"Zteal20"},
		{"teal20Z"},
		{"Zteal200"},
		{"teal200Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Teal200()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestTeal300Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"teal300", 7},
		{"teal300 ", 7},
		{"teal300\n", 7},
		{"teal300.", 7},
		{"teal300:", 7},
		{"teal300,", 7},
		{"teal300\"", 7},
		{"teal300(", 7},
		{"teal300)", 7},
		{"teal300[", 7},
		{"teal300]", 7},
		{"teal300// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Teal300()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestTeal300Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_t"},
		{"t_"},
		{"_te"},
		{"te_"},
		{"_tea"},
		{"tea_"},
		{"_teal"},
		{"teal_"},
		{"_teal3"},
		{"teal3_"},
		{"_teal30"},
		{"teal30_"},
		{"_teal300"},
		{"teal300_"},
		{"9713t"},
		{"t9713"},
		{"9713te"},
		{"te9713"},
		{"9713tea"},
		{"tea9713"},
		{"9713teal"},
		{"teal9713"},
		{"9713teal3"},
		{"teal39713"},
		{"9713teal30"},
		{"teal309713"},
		{"9713teal300"},
		{"teal3009713"},
		{"Zt"},
		{"tZ"},
		{"Zte"},
		{"teZ"},
		{"Ztea"},
		{"teaZ"},
		{"Zteal"},
		{"tealZ"},
		{"Zteal3"},
		{"teal3Z"},
		{"Zteal30"},
		{"teal30Z"},
		{"Zteal300"},
		{"teal300Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Teal300()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestTeal400Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"teal400", 7},
		{"teal400 ", 7},
		{"teal400\n", 7},
		{"teal400.", 7},
		{"teal400:", 7},
		{"teal400,", 7},
		{"teal400\"", 7},
		{"teal400(", 7},
		{"teal400)", 7},
		{"teal400[", 7},
		{"teal400]", 7},
		{"teal400// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Teal400()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestTeal400Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_t"},
		{"t_"},
		{"_te"},
		{"te_"},
		{"_tea"},
		{"tea_"},
		{"_teal"},
		{"teal_"},
		{"_teal4"},
		{"teal4_"},
		{"_teal40"},
		{"teal40_"},
		{"_teal400"},
		{"teal400_"},
		{"9713t"},
		{"t9713"},
		{"9713te"},
		{"te9713"},
		{"9713tea"},
		{"tea9713"},
		{"9713teal"},
		{"teal9713"},
		{"9713teal4"},
		{"teal49713"},
		{"9713teal40"},
		{"teal409713"},
		{"9713teal400"},
		{"teal4009713"},
		{"Zt"},
		{"tZ"},
		{"Zte"},
		{"teZ"},
		{"Ztea"},
		{"teaZ"},
		{"Zteal"},
		{"tealZ"},
		{"Zteal4"},
		{"teal4Z"},
		{"Zteal40"},
		{"teal40Z"},
		{"Zteal400"},
		{"teal400Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Teal400()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestTeal500Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"teal500", 7},
		{"teal500 ", 7},
		{"teal500\n", 7},
		{"teal500.", 7},
		{"teal500:", 7},
		{"teal500,", 7},
		{"teal500\"", 7},
		{"teal500(", 7},
		{"teal500)", 7},
		{"teal500[", 7},
		{"teal500]", 7},
		{"teal500// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Teal500()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestTeal500Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_t"},
		{"t_"},
		{"_te"},
		{"te_"},
		{"_tea"},
		{"tea_"},
		{"_teal"},
		{"teal_"},
		{"_teal5"},
		{"teal5_"},
		{"_teal50"},
		{"teal50_"},
		{"_teal500"},
		{"teal500_"},
		{"9713t"},
		{"t9713"},
		{"9713te"},
		{"te9713"},
		{"9713tea"},
		{"tea9713"},
		{"9713teal"},
		{"teal9713"},
		{"9713teal5"},
		{"teal59713"},
		{"9713teal50"},
		{"teal509713"},
		{"9713teal500"},
		{"teal5009713"},
		{"Zt"},
		{"tZ"},
		{"Zte"},
		{"teZ"},
		{"Ztea"},
		{"teaZ"},
		{"Zteal"},
		{"tealZ"},
		{"Zteal5"},
		{"teal5Z"},
		{"Zteal50"},
		{"teal50Z"},
		{"Zteal500"},
		{"teal500Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Teal500()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestTeal600Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"teal600", 7},
		{"teal600 ", 7},
		{"teal600\n", 7},
		{"teal600.", 7},
		{"teal600:", 7},
		{"teal600,", 7},
		{"teal600\"", 7},
		{"teal600(", 7},
		{"teal600)", 7},
		{"teal600[", 7},
		{"teal600]", 7},
		{"teal600// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Teal600()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestTeal600Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_t"},
		{"t_"},
		{"_te"},
		{"te_"},
		{"_tea"},
		{"tea_"},
		{"_teal"},
		{"teal_"},
		{"_teal6"},
		{"teal6_"},
		{"_teal60"},
		{"teal60_"},
		{"_teal600"},
		{"teal600_"},
		{"9713t"},
		{"t9713"},
		{"9713te"},
		{"te9713"},
		{"9713tea"},
		{"tea9713"},
		{"9713teal"},
		{"teal9713"},
		{"9713teal6"},
		{"teal69713"},
		{"9713teal60"},
		{"teal609713"},
		{"9713teal600"},
		{"teal6009713"},
		{"Zt"},
		{"tZ"},
		{"Zte"},
		{"teZ"},
		{"Ztea"},
		{"teaZ"},
		{"Zteal"},
		{"tealZ"},
		{"Zteal6"},
		{"teal6Z"},
		{"Zteal60"},
		{"teal60Z"},
		{"Zteal600"},
		{"teal600Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Teal600()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestTeal700Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"teal700", 7},
		{"teal700 ", 7},
		{"teal700\n", 7},
		{"teal700.", 7},
		{"teal700:", 7},
		{"teal700,", 7},
		{"teal700\"", 7},
		{"teal700(", 7},
		{"teal700)", 7},
		{"teal700[", 7},
		{"teal700]", 7},
		{"teal700// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Teal700()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestTeal700Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_t"},
		{"t_"},
		{"_te"},
		{"te_"},
		{"_tea"},
		{"tea_"},
		{"_teal"},
		{"teal_"},
		{"_teal7"},
		{"teal7_"},
		{"_teal70"},
		{"teal70_"},
		{"_teal700"},
		{"teal700_"},
		{"9713t"},
		{"t9713"},
		{"9713te"},
		{"te9713"},
		{"9713tea"},
		{"tea9713"},
		{"9713teal"},
		{"teal9713"},
		{"9713teal7"},
		{"teal79713"},
		{"9713teal70"},
		{"teal709713"},
		{"9713teal700"},
		{"teal7009713"},
		{"Zt"},
		{"tZ"},
		{"Zte"},
		{"teZ"},
		{"Ztea"},
		{"teaZ"},
		{"Zteal"},
		{"tealZ"},
		{"Zteal7"},
		{"teal7Z"},
		{"Zteal70"},
		{"teal70Z"},
		{"Zteal700"},
		{"teal700Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Teal700()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestTeal800Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"teal800", 7},
		{"teal800 ", 7},
		{"teal800\n", 7},
		{"teal800.", 7},
		{"teal800:", 7},
		{"teal800,", 7},
		{"teal800\"", 7},
		{"teal800(", 7},
		{"teal800)", 7},
		{"teal800[", 7},
		{"teal800]", 7},
		{"teal800// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Teal800()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestTeal800Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_t"},
		{"t_"},
		{"_te"},
		{"te_"},
		{"_tea"},
		{"tea_"},
		{"_teal"},
		{"teal_"},
		{"_teal8"},
		{"teal8_"},
		{"_teal80"},
		{"teal80_"},
		{"_teal800"},
		{"teal800_"},
		{"9713t"},
		{"t9713"},
		{"9713te"},
		{"te9713"},
		{"9713tea"},
		{"tea9713"},
		{"9713teal"},
		{"teal9713"},
		{"9713teal8"},
		{"teal89713"},
		{"9713teal80"},
		{"teal809713"},
		{"9713teal800"},
		{"teal8009713"},
		{"Zt"},
		{"tZ"},
		{"Zte"},
		{"teZ"},
		{"Ztea"},
		{"teaZ"},
		{"Zteal"},
		{"tealZ"},
		{"Zteal8"},
		{"teal8Z"},
		{"Zteal80"},
		{"teal80Z"},
		{"Zteal800"},
		{"teal800Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Teal800()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestTeal900Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"teal900", 7},
		{"teal900 ", 7},
		{"teal900\n", 7},
		{"teal900.", 7},
		{"teal900:", 7},
		{"teal900,", 7},
		{"teal900\"", 7},
		{"teal900(", 7},
		{"teal900)", 7},
		{"teal900[", 7},
		{"teal900]", 7},
		{"teal900// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Teal900()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestTeal900Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_t"},
		{"t_"},
		{"_te"},
		{"te_"},
		{"_tea"},
		{"tea_"},
		{"_teal"},
		{"teal_"},
		{"_teal9"},
		{"teal9_"},
		{"_teal90"},
		{"teal90_"},
		{"_teal900"},
		{"teal900_"},
		{"9713t"},
		{"t9713"},
		{"9713te"},
		{"te9713"},
		{"9713tea"},
		{"tea9713"},
		{"9713teal"},
		{"teal9713"},
		{"9713teal9"},
		{"teal99713"},
		{"9713teal90"},
		{"teal909713"},
		{"9713teal900"},
		{"teal9009713"},
		{"Zt"},
		{"tZ"},
		{"Zte"},
		{"teZ"},
		{"Ztea"},
		{"teaZ"},
		{"Zteal"},
		{"tealZ"},
		{"Zteal9"},
		{"teal9Z"},
		{"Zteal90"},
		{"teal90Z"},
		{"Zteal900"},
		{"teal900Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Teal900()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestTealA100Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"tealA100", 8},
		{"tealA100 ", 8},
		{"tealA100\n", 8},
		{"tealA100.", 8},
		{"tealA100:", 8},
		{"tealA100,", 8},
		{"tealA100\"", 8},
		{"tealA100(", 8},
		{"tealA100)", 8},
		{"tealA100[", 8},
		{"tealA100]", 8},
		{"tealA100// comment", 8},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.TealA100()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestTealA100Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_t"},
		{"t_"},
		{"_te"},
		{"te_"},
		{"_tea"},
		{"tea_"},
		{"_teal"},
		{"teal_"},
		{"_tealA"},
		{"tealA_"},
		{"_tealA1"},
		{"tealA1_"},
		{"_tealA10"},
		{"tealA10_"},
		{"_tealA100"},
		{"tealA100_"},
		{"9713t"},
		{"t9713"},
		{"9713te"},
		{"te9713"},
		{"9713tea"},
		{"tea9713"},
		{"9713teal"},
		{"teal9713"},
		{"9713tealA"},
		{"tealA9713"},
		{"9713tealA1"},
		{"tealA19713"},
		{"9713tealA10"},
		{"tealA109713"},
		{"9713tealA100"},
		{"tealA1009713"},
		{"Zt"},
		{"tZ"},
		{"Zte"},
		{"teZ"},
		{"Ztea"},
		{"teaZ"},
		{"Zteal"},
		{"tealZ"},
		{"ZtealA"},
		{"tealAZ"},
		{"ZtealA1"},
		{"tealA1Z"},
		{"ZtealA10"},
		{"tealA10Z"},
		{"ZtealA100"},
		{"tealA100Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.TealA100()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestTealA200Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"tealA200", 8},
		{"tealA200 ", 8},
		{"tealA200\n", 8},
		{"tealA200.", 8},
		{"tealA200:", 8},
		{"tealA200,", 8},
		{"tealA200\"", 8},
		{"tealA200(", 8},
		{"tealA200)", 8},
		{"tealA200[", 8},
		{"tealA200]", 8},
		{"tealA200// comment", 8},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.TealA200()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestTealA200Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_t"},
		{"t_"},
		{"_te"},
		{"te_"},
		{"_tea"},
		{"tea_"},
		{"_teal"},
		{"teal_"},
		{"_tealA"},
		{"tealA_"},
		{"_tealA2"},
		{"tealA2_"},
		{"_tealA20"},
		{"tealA20_"},
		{"_tealA200"},
		{"tealA200_"},
		{"9713t"},
		{"t9713"},
		{"9713te"},
		{"te9713"},
		{"9713tea"},
		{"tea9713"},
		{"9713teal"},
		{"teal9713"},
		{"9713tealA"},
		{"tealA9713"},
		{"9713tealA2"},
		{"tealA29713"},
		{"9713tealA20"},
		{"tealA209713"},
		{"9713tealA200"},
		{"tealA2009713"},
		{"Zt"},
		{"tZ"},
		{"Zte"},
		{"teZ"},
		{"Ztea"},
		{"teaZ"},
		{"Zteal"},
		{"tealZ"},
		{"ZtealA"},
		{"tealAZ"},
		{"ZtealA2"},
		{"tealA2Z"},
		{"ZtealA20"},
		{"tealA20Z"},
		{"ZtealA200"},
		{"tealA200Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.TealA200()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestTealA400Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"tealA400", 8},
		{"tealA400 ", 8},
		{"tealA400\n", 8},
		{"tealA400.", 8},
		{"tealA400:", 8},
		{"tealA400,", 8},
		{"tealA400\"", 8},
		{"tealA400(", 8},
		{"tealA400)", 8},
		{"tealA400[", 8},
		{"tealA400]", 8},
		{"tealA400// comment", 8},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.TealA400()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestTealA400Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_t"},
		{"t_"},
		{"_te"},
		{"te_"},
		{"_tea"},
		{"tea_"},
		{"_teal"},
		{"teal_"},
		{"_tealA"},
		{"tealA_"},
		{"_tealA4"},
		{"tealA4_"},
		{"_tealA40"},
		{"tealA40_"},
		{"_tealA400"},
		{"tealA400_"},
		{"9713t"},
		{"t9713"},
		{"9713te"},
		{"te9713"},
		{"9713tea"},
		{"tea9713"},
		{"9713teal"},
		{"teal9713"},
		{"9713tealA"},
		{"tealA9713"},
		{"9713tealA4"},
		{"tealA49713"},
		{"9713tealA40"},
		{"tealA409713"},
		{"9713tealA400"},
		{"tealA4009713"},
		{"Zt"},
		{"tZ"},
		{"Zte"},
		{"teZ"},
		{"Ztea"},
		{"teaZ"},
		{"Zteal"},
		{"tealZ"},
		{"ZtealA"},
		{"tealAZ"},
		{"ZtealA4"},
		{"tealA4Z"},
		{"ZtealA40"},
		{"tealA40Z"},
		{"ZtealA400"},
		{"tealA400Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.TealA400()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestTealA700Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"tealA700", 8},
		{"tealA700 ", 8},
		{"tealA700\n", 8},
		{"tealA700.", 8},
		{"tealA700:", 8},
		{"tealA700,", 8},
		{"tealA700\"", 8},
		{"tealA700(", 8},
		{"tealA700)", 8},
		{"tealA700[", 8},
		{"tealA700]", 8},
		{"tealA700// comment", 8},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.TealA700()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestTealA700Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_t"},
		{"t_"},
		{"_te"},
		{"te_"},
		{"_tea"},
		{"tea_"},
		{"_teal"},
		{"teal_"},
		{"_tealA"},
		{"tealA_"},
		{"_tealA7"},
		{"tealA7_"},
		{"_tealA70"},
		{"tealA70_"},
		{"_tealA700"},
		{"tealA700_"},
		{"9713t"},
		{"t9713"},
		{"9713te"},
		{"te9713"},
		{"9713tea"},
		{"tea9713"},
		{"9713teal"},
		{"teal9713"},
		{"9713tealA"},
		{"tealA9713"},
		{"9713tealA7"},
		{"tealA79713"},
		{"9713tealA70"},
		{"tealA709713"},
		{"9713tealA700"},
		{"tealA7009713"},
		{"Zt"},
		{"tZ"},
		{"Zte"},
		{"teZ"},
		{"Ztea"},
		{"teaZ"},
		{"Zteal"},
		{"tealZ"},
		{"ZtealA"},
		{"tealAZ"},
		{"ZtealA7"},
		{"tealA7Z"},
		{"ZtealA70"},
		{"tealA70Z"},
		{"ZtealA700"},
		{"tealA700Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.TealA700()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestGreen50Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"green50", 7},
		{"green50 ", 7},
		{"green50\n", 7},
		{"green50.", 7},
		{"green50:", 7},
		{"green50,", 7},
		{"green50\"", 7},
		{"green50(", 7},
		{"green50)", 7},
		{"green50[", 7},
		{"green50]", 7},
		{"green50// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Green50()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestGreen50Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_g"},
		{"g_"},
		{"_gr"},
		{"gr_"},
		{"_gre"},
		{"gre_"},
		{"_gree"},
		{"gree_"},
		{"_green"},
		{"green_"},
		{"_green5"},
		{"green5_"},
		{"_green50"},
		{"green50_"},
		{"9713g"},
		{"g9713"},
		{"9713gr"},
		{"gr9713"},
		{"9713gre"},
		{"gre9713"},
		{"9713gree"},
		{"gree9713"},
		{"9713green"},
		{"green9713"},
		{"9713green5"},
		{"green59713"},
		{"9713green50"},
		{"green509713"},
		{"Zg"},
		{"gZ"},
		{"Zgr"},
		{"grZ"},
		{"Zgre"},
		{"greZ"},
		{"Zgree"},
		{"greeZ"},
		{"Zgreen"},
		{"greenZ"},
		{"Zgreen5"},
		{"green5Z"},
		{"Zgreen50"},
		{"green50Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Green50()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestGreen100Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"green100", 8},
		{"green100 ", 8},
		{"green100\n", 8},
		{"green100.", 8},
		{"green100:", 8},
		{"green100,", 8},
		{"green100\"", 8},
		{"green100(", 8},
		{"green100)", 8},
		{"green100[", 8},
		{"green100]", 8},
		{"green100// comment", 8},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Green100()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestGreen100Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_g"},
		{"g_"},
		{"_gr"},
		{"gr_"},
		{"_gre"},
		{"gre_"},
		{"_gree"},
		{"gree_"},
		{"_green"},
		{"green_"},
		{"_green1"},
		{"green1_"},
		{"_green10"},
		{"green10_"},
		{"_green100"},
		{"green100_"},
		{"9713g"},
		{"g9713"},
		{"9713gr"},
		{"gr9713"},
		{"9713gre"},
		{"gre9713"},
		{"9713gree"},
		{"gree9713"},
		{"9713green"},
		{"green9713"},
		{"9713green1"},
		{"green19713"},
		{"9713green10"},
		{"green109713"},
		{"9713green100"},
		{"green1009713"},
		{"Zg"},
		{"gZ"},
		{"Zgr"},
		{"grZ"},
		{"Zgre"},
		{"greZ"},
		{"Zgree"},
		{"greeZ"},
		{"Zgreen"},
		{"greenZ"},
		{"Zgreen1"},
		{"green1Z"},
		{"Zgreen10"},
		{"green10Z"},
		{"Zgreen100"},
		{"green100Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Green100()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestGreen200Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"green200", 8},
		{"green200 ", 8},
		{"green200\n", 8},
		{"green200.", 8},
		{"green200:", 8},
		{"green200,", 8},
		{"green200\"", 8},
		{"green200(", 8},
		{"green200)", 8},
		{"green200[", 8},
		{"green200]", 8},
		{"green200// comment", 8},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Green200()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestGreen200Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_g"},
		{"g_"},
		{"_gr"},
		{"gr_"},
		{"_gre"},
		{"gre_"},
		{"_gree"},
		{"gree_"},
		{"_green"},
		{"green_"},
		{"_green2"},
		{"green2_"},
		{"_green20"},
		{"green20_"},
		{"_green200"},
		{"green200_"},
		{"9713g"},
		{"g9713"},
		{"9713gr"},
		{"gr9713"},
		{"9713gre"},
		{"gre9713"},
		{"9713gree"},
		{"gree9713"},
		{"9713green"},
		{"green9713"},
		{"9713green2"},
		{"green29713"},
		{"9713green20"},
		{"green209713"},
		{"9713green200"},
		{"green2009713"},
		{"Zg"},
		{"gZ"},
		{"Zgr"},
		{"grZ"},
		{"Zgre"},
		{"greZ"},
		{"Zgree"},
		{"greeZ"},
		{"Zgreen"},
		{"greenZ"},
		{"Zgreen2"},
		{"green2Z"},
		{"Zgreen20"},
		{"green20Z"},
		{"Zgreen200"},
		{"green200Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Green200()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestGreen300Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"green300", 8},
		{"green300 ", 8},
		{"green300\n", 8},
		{"green300.", 8},
		{"green300:", 8},
		{"green300,", 8},
		{"green300\"", 8},
		{"green300(", 8},
		{"green300)", 8},
		{"green300[", 8},
		{"green300]", 8},
		{"green300// comment", 8},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Green300()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestGreen300Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_g"},
		{"g_"},
		{"_gr"},
		{"gr_"},
		{"_gre"},
		{"gre_"},
		{"_gree"},
		{"gree_"},
		{"_green"},
		{"green_"},
		{"_green3"},
		{"green3_"},
		{"_green30"},
		{"green30_"},
		{"_green300"},
		{"green300_"},
		{"9713g"},
		{"g9713"},
		{"9713gr"},
		{"gr9713"},
		{"9713gre"},
		{"gre9713"},
		{"9713gree"},
		{"gree9713"},
		{"9713green"},
		{"green9713"},
		{"9713green3"},
		{"green39713"},
		{"9713green30"},
		{"green309713"},
		{"9713green300"},
		{"green3009713"},
		{"Zg"},
		{"gZ"},
		{"Zgr"},
		{"grZ"},
		{"Zgre"},
		{"greZ"},
		{"Zgree"},
		{"greeZ"},
		{"Zgreen"},
		{"greenZ"},
		{"Zgreen3"},
		{"green3Z"},
		{"Zgreen30"},
		{"green30Z"},
		{"Zgreen300"},
		{"green300Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Green300()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestGreen400Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"green400", 8},
		{"green400 ", 8},
		{"green400\n", 8},
		{"green400.", 8},
		{"green400:", 8},
		{"green400,", 8},
		{"green400\"", 8},
		{"green400(", 8},
		{"green400)", 8},
		{"green400[", 8},
		{"green400]", 8},
		{"green400// comment", 8},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Green400()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestGreen400Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_g"},
		{"g_"},
		{"_gr"},
		{"gr_"},
		{"_gre"},
		{"gre_"},
		{"_gree"},
		{"gree_"},
		{"_green"},
		{"green_"},
		{"_green4"},
		{"green4_"},
		{"_green40"},
		{"green40_"},
		{"_green400"},
		{"green400_"},
		{"9713g"},
		{"g9713"},
		{"9713gr"},
		{"gr9713"},
		{"9713gre"},
		{"gre9713"},
		{"9713gree"},
		{"gree9713"},
		{"9713green"},
		{"green9713"},
		{"9713green4"},
		{"green49713"},
		{"9713green40"},
		{"green409713"},
		{"9713green400"},
		{"green4009713"},
		{"Zg"},
		{"gZ"},
		{"Zgr"},
		{"grZ"},
		{"Zgre"},
		{"greZ"},
		{"Zgree"},
		{"greeZ"},
		{"Zgreen"},
		{"greenZ"},
		{"Zgreen4"},
		{"green4Z"},
		{"Zgreen40"},
		{"green40Z"},
		{"Zgreen400"},
		{"green400Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Green400()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestGreen500Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"green500", 8},
		{"green500 ", 8},
		{"green500\n", 8},
		{"green500.", 8},
		{"green500:", 8},
		{"green500,", 8},
		{"green500\"", 8},
		{"green500(", 8},
		{"green500)", 8},
		{"green500[", 8},
		{"green500]", 8},
		{"green500// comment", 8},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Green500()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestGreen500Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_g"},
		{"g_"},
		{"_gr"},
		{"gr_"},
		{"_gre"},
		{"gre_"},
		{"_gree"},
		{"gree_"},
		{"_green"},
		{"green_"},
		{"_green5"},
		{"green5_"},
		{"_green50"},
		{"green50_"},
		{"_green500"},
		{"green500_"},
		{"9713g"},
		{"g9713"},
		{"9713gr"},
		{"gr9713"},
		{"9713gre"},
		{"gre9713"},
		{"9713gree"},
		{"gree9713"},
		{"9713green"},
		{"green9713"},
		{"9713green5"},
		{"green59713"},
		{"9713green50"},
		{"green509713"},
		{"9713green500"},
		{"green5009713"},
		{"Zg"},
		{"gZ"},
		{"Zgr"},
		{"grZ"},
		{"Zgre"},
		{"greZ"},
		{"Zgree"},
		{"greeZ"},
		{"Zgreen"},
		{"greenZ"},
		{"Zgreen5"},
		{"green5Z"},
		{"Zgreen50"},
		{"green50Z"},
		{"Zgreen500"},
		{"green500Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Green500()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestGreen600Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"green600", 8},
		{"green600 ", 8},
		{"green600\n", 8},
		{"green600.", 8},
		{"green600:", 8},
		{"green600,", 8},
		{"green600\"", 8},
		{"green600(", 8},
		{"green600)", 8},
		{"green600[", 8},
		{"green600]", 8},
		{"green600// comment", 8},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Green600()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestGreen600Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_g"},
		{"g_"},
		{"_gr"},
		{"gr_"},
		{"_gre"},
		{"gre_"},
		{"_gree"},
		{"gree_"},
		{"_green"},
		{"green_"},
		{"_green6"},
		{"green6_"},
		{"_green60"},
		{"green60_"},
		{"_green600"},
		{"green600_"},
		{"9713g"},
		{"g9713"},
		{"9713gr"},
		{"gr9713"},
		{"9713gre"},
		{"gre9713"},
		{"9713gree"},
		{"gree9713"},
		{"9713green"},
		{"green9713"},
		{"9713green6"},
		{"green69713"},
		{"9713green60"},
		{"green609713"},
		{"9713green600"},
		{"green6009713"},
		{"Zg"},
		{"gZ"},
		{"Zgr"},
		{"grZ"},
		{"Zgre"},
		{"greZ"},
		{"Zgree"},
		{"greeZ"},
		{"Zgreen"},
		{"greenZ"},
		{"Zgreen6"},
		{"green6Z"},
		{"Zgreen60"},
		{"green60Z"},
		{"Zgreen600"},
		{"green600Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Green600()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestGreen700Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"green700", 8},
		{"green700 ", 8},
		{"green700\n", 8},
		{"green700.", 8},
		{"green700:", 8},
		{"green700,", 8},
		{"green700\"", 8},
		{"green700(", 8},
		{"green700)", 8},
		{"green700[", 8},
		{"green700]", 8},
		{"green700// comment", 8},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Green700()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestGreen700Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_g"},
		{"g_"},
		{"_gr"},
		{"gr_"},
		{"_gre"},
		{"gre_"},
		{"_gree"},
		{"gree_"},
		{"_green"},
		{"green_"},
		{"_green7"},
		{"green7_"},
		{"_green70"},
		{"green70_"},
		{"_green700"},
		{"green700_"},
		{"9713g"},
		{"g9713"},
		{"9713gr"},
		{"gr9713"},
		{"9713gre"},
		{"gre9713"},
		{"9713gree"},
		{"gree9713"},
		{"9713green"},
		{"green9713"},
		{"9713green7"},
		{"green79713"},
		{"9713green70"},
		{"green709713"},
		{"9713green700"},
		{"green7009713"},
		{"Zg"},
		{"gZ"},
		{"Zgr"},
		{"grZ"},
		{"Zgre"},
		{"greZ"},
		{"Zgree"},
		{"greeZ"},
		{"Zgreen"},
		{"greenZ"},
		{"Zgreen7"},
		{"green7Z"},
		{"Zgreen70"},
		{"green70Z"},
		{"Zgreen700"},
		{"green700Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Green700()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestGreen800Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"green800", 8},
		{"green800 ", 8},
		{"green800\n", 8},
		{"green800.", 8},
		{"green800:", 8},
		{"green800,", 8},
		{"green800\"", 8},
		{"green800(", 8},
		{"green800)", 8},
		{"green800[", 8},
		{"green800]", 8},
		{"green800// comment", 8},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Green800()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestGreen800Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_g"},
		{"g_"},
		{"_gr"},
		{"gr_"},
		{"_gre"},
		{"gre_"},
		{"_gree"},
		{"gree_"},
		{"_green"},
		{"green_"},
		{"_green8"},
		{"green8_"},
		{"_green80"},
		{"green80_"},
		{"_green800"},
		{"green800_"},
		{"9713g"},
		{"g9713"},
		{"9713gr"},
		{"gr9713"},
		{"9713gre"},
		{"gre9713"},
		{"9713gree"},
		{"gree9713"},
		{"9713green"},
		{"green9713"},
		{"9713green8"},
		{"green89713"},
		{"9713green80"},
		{"green809713"},
		{"9713green800"},
		{"green8009713"},
		{"Zg"},
		{"gZ"},
		{"Zgr"},
		{"grZ"},
		{"Zgre"},
		{"greZ"},
		{"Zgree"},
		{"greeZ"},
		{"Zgreen"},
		{"greenZ"},
		{"Zgreen8"},
		{"green8Z"},
		{"Zgreen80"},
		{"green80Z"},
		{"Zgreen800"},
		{"green800Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Green800()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestGreen900Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"green900", 8},
		{"green900 ", 8},
		{"green900\n", 8},
		{"green900.", 8},
		{"green900:", 8},
		{"green900,", 8},
		{"green900\"", 8},
		{"green900(", 8},
		{"green900)", 8},
		{"green900[", 8},
		{"green900]", 8},
		{"green900// comment", 8},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Green900()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestGreen900Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_g"},
		{"g_"},
		{"_gr"},
		{"gr_"},
		{"_gre"},
		{"gre_"},
		{"_gree"},
		{"gree_"},
		{"_green"},
		{"green_"},
		{"_green9"},
		{"green9_"},
		{"_green90"},
		{"green90_"},
		{"_green900"},
		{"green900_"},
		{"9713g"},
		{"g9713"},
		{"9713gr"},
		{"gr9713"},
		{"9713gre"},
		{"gre9713"},
		{"9713gree"},
		{"gree9713"},
		{"9713green"},
		{"green9713"},
		{"9713green9"},
		{"green99713"},
		{"9713green90"},
		{"green909713"},
		{"9713green900"},
		{"green9009713"},
		{"Zg"},
		{"gZ"},
		{"Zgr"},
		{"grZ"},
		{"Zgre"},
		{"greZ"},
		{"Zgree"},
		{"greeZ"},
		{"Zgreen"},
		{"greenZ"},
		{"Zgreen9"},
		{"green9Z"},
		{"Zgreen90"},
		{"green90Z"},
		{"Zgreen900"},
		{"green900Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Green900()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestGreenA100Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"greenA100", 9},
		{"greenA100 ", 9},
		{"greenA100\n", 9},
		{"greenA100.", 9},
		{"greenA100:", 9},
		{"greenA100,", 9},
		{"greenA100\"", 9},
		{"greenA100(", 9},
		{"greenA100)", 9},
		{"greenA100[", 9},
		{"greenA100]", 9},
		{"greenA100// comment", 9},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.GreenA100()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestGreenA100Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_g"},
		{"g_"},
		{"_gr"},
		{"gr_"},
		{"_gre"},
		{"gre_"},
		{"_gree"},
		{"gree_"},
		{"_green"},
		{"green_"},
		{"_greenA"},
		{"greenA_"},
		{"_greenA1"},
		{"greenA1_"},
		{"_greenA10"},
		{"greenA10_"},
		{"_greenA100"},
		{"greenA100_"},
		{"9713g"},
		{"g9713"},
		{"9713gr"},
		{"gr9713"},
		{"9713gre"},
		{"gre9713"},
		{"9713gree"},
		{"gree9713"},
		{"9713green"},
		{"green9713"},
		{"9713greenA"},
		{"greenA9713"},
		{"9713greenA1"},
		{"greenA19713"},
		{"9713greenA10"},
		{"greenA109713"},
		{"9713greenA100"},
		{"greenA1009713"},
		{"Zg"},
		{"gZ"},
		{"Zgr"},
		{"grZ"},
		{"Zgre"},
		{"greZ"},
		{"Zgree"},
		{"greeZ"},
		{"Zgreen"},
		{"greenZ"},
		{"ZgreenA"},
		{"greenAZ"},
		{"ZgreenA1"},
		{"greenA1Z"},
		{"ZgreenA10"},
		{"greenA10Z"},
		{"ZgreenA100"},
		{"greenA100Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.GreenA100()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestGreenA200Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"greenA200", 9},
		{"greenA200 ", 9},
		{"greenA200\n", 9},
		{"greenA200.", 9},
		{"greenA200:", 9},
		{"greenA200,", 9},
		{"greenA200\"", 9},
		{"greenA200(", 9},
		{"greenA200)", 9},
		{"greenA200[", 9},
		{"greenA200]", 9},
		{"greenA200// comment", 9},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.GreenA200()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestGreenA200Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_g"},
		{"g_"},
		{"_gr"},
		{"gr_"},
		{"_gre"},
		{"gre_"},
		{"_gree"},
		{"gree_"},
		{"_green"},
		{"green_"},
		{"_greenA"},
		{"greenA_"},
		{"_greenA2"},
		{"greenA2_"},
		{"_greenA20"},
		{"greenA20_"},
		{"_greenA200"},
		{"greenA200_"},
		{"9713g"},
		{"g9713"},
		{"9713gr"},
		{"gr9713"},
		{"9713gre"},
		{"gre9713"},
		{"9713gree"},
		{"gree9713"},
		{"9713green"},
		{"green9713"},
		{"9713greenA"},
		{"greenA9713"},
		{"9713greenA2"},
		{"greenA29713"},
		{"9713greenA20"},
		{"greenA209713"},
		{"9713greenA200"},
		{"greenA2009713"},
		{"Zg"},
		{"gZ"},
		{"Zgr"},
		{"grZ"},
		{"Zgre"},
		{"greZ"},
		{"Zgree"},
		{"greeZ"},
		{"Zgreen"},
		{"greenZ"},
		{"ZgreenA"},
		{"greenAZ"},
		{"ZgreenA2"},
		{"greenA2Z"},
		{"ZgreenA20"},
		{"greenA20Z"},
		{"ZgreenA200"},
		{"greenA200Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.GreenA200()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestGreenA400Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"greenA400", 9},
		{"greenA400 ", 9},
		{"greenA400\n", 9},
		{"greenA400.", 9},
		{"greenA400:", 9},
		{"greenA400,", 9},
		{"greenA400\"", 9},
		{"greenA400(", 9},
		{"greenA400)", 9},
		{"greenA400[", 9},
		{"greenA400]", 9},
		{"greenA400// comment", 9},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.GreenA400()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestGreenA400Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_g"},
		{"g_"},
		{"_gr"},
		{"gr_"},
		{"_gre"},
		{"gre_"},
		{"_gree"},
		{"gree_"},
		{"_green"},
		{"green_"},
		{"_greenA"},
		{"greenA_"},
		{"_greenA4"},
		{"greenA4_"},
		{"_greenA40"},
		{"greenA40_"},
		{"_greenA400"},
		{"greenA400_"},
		{"9713g"},
		{"g9713"},
		{"9713gr"},
		{"gr9713"},
		{"9713gre"},
		{"gre9713"},
		{"9713gree"},
		{"gree9713"},
		{"9713green"},
		{"green9713"},
		{"9713greenA"},
		{"greenA9713"},
		{"9713greenA4"},
		{"greenA49713"},
		{"9713greenA40"},
		{"greenA409713"},
		{"9713greenA400"},
		{"greenA4009713"},
		{"Zg"},
		{"gZ"},
		{"Zgr"},
		{"grZ"},
		{"Zgre"},
		{"greZ"},
		{"Zgree"},
		{"greeZ"},
		{"Zgreen"},
		{"greenZ"},
		{"ZgreenA"},
		{"greenAZ"},
		{"ZgreenA4"},
		{"greenA4Z"},
		{"ZgreenA40"},
		{"greenA40Z"},
		{"ZgreenA400"},
		{"greenA400Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.GreenA400()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestGreenA700Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"greenA700", 9},
		{"greenA700 ", 9},
		{"greenA700\n", 9},
		{"greenA700.", 9},
		{"greenA700:", 9},
		{"greenA700,", 9},
		{"greenA700\"", 9},
		{"greenA700(", 9},
		{"greenA700)", 9},
		{"greenA700[", 9},
		{"greenA700]", 9},
		{"greenA700// comment", 9},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.GreenA700()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestGreenA700Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_g"},
		{"g_"},
		{"_gr"},
		{"gr_"},
		{"_gre"},
		{"gre_"},
		{"_gree"},
		{"gree_"},
		{"_green"},
		{"green_"},
		{"_greenA"},
		{"greenA_"},
		{"_greenA7"},
		{"greenA7_"},
		{"_greenA70"},
		{"greenA70_"},
		{"_greenA700"},
		{"greenA700_"},
		{"9713g"},
		{"g9713"},
		{"9713gr"},
		{"gr9713"},
		{"9713gre"},
		{"gre9713"},
		{"9713gree"},
		{"gree9713"},
		{"9713green"},
		{"green9713"},
		{"9713greenA"},
		{"greenA9713"},
		{"9713greenA7"},
		{"greenA79713"},
		{"9713greenA70"},
		{"greenA709713"},
		{"9713greenA700"},
		{"greenA7009713"},
		{"Zg"},
		{"gZ"},
		{"Zgr"},
		{"grZ"},
		{"Zgre"},
		{"greZ"},
		{"Zgree"},
		{"greeZ"},
		{"Zgreen"},
		{"greenZ"},
		{"ZgreenA"},
		{"greenAZ"},
		{"ZgreenA7"},
		{"greenA7Z"},
		{"ZgreenA70"},
		{"greenA70Z"},
		{"ZgreenA700"},
		{"greenA700Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.GreenA700()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestLightGreen50Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"lightGreen50", 12},
		{"lightGreen50 ", 12},
		{"lightGreen50\n", 12},
		{"lightGreen50.", 12},
		{"lightGreen50:", 12},
		{"lightGreen50,", 12},
		{"lightGreen50\"", 12},
		{"lightGreen50(", 12},
		{"lightGreen50)", 12},
		{"lightGreen50[", 12},
		{"lightGreen50]", 12},
		{"lightGreen50// comment", 12},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.LightGreen50()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestLightGreen50Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_l"},
		{"l_"},
		{"_li"},
		{"li_"},
		{"_lig"},
		{"lig_"},
		{"_ligh"},
		{"ligh_"},
		{"_light"},
		{"light_"},
		{"_lightG"},
		{"lightG_"},
		{"_lightGr"},
		{"lightGr_"},
		{"_lightGre"},
		{"lightGre_"},
		{"_lightGree"},
		{"lightGree_"},
		{"_lightGreen"},
		{"lightGreen_"},
		{"_lightGreen5"},
		{"lightGreen5_"},
		{"_lightGreen50"},
		{"lightGreen50_"},
		{"9713l"},
		{"l9713"},
		{"9713li"},
		{"li9713"},
		{"9713lig"},
		{"lig9713"},
		{"9713ligh"},
		{"ligh9713"},
		{"9713light"},
		{"light9713"},
		{"9713lightG"},
		{"lightG9713"},
		{"9713lightGr"},
		{"lightGr9713"},
		{"9713lightGre"},
		{"lightGre9713"},
		{"9713lightGree"},
		{"lightGree9713"},
		{"9713lightGreen"},
		{"lightGreen9713"},
		{"9713lightGreen5"},
		{"lightGreen59713"},
		{"9713lightGreen50"},
		{"lightGreen509713"},
		{"Zl"},
		{"lZ"},
		{"Zli"},
		{"liZ"},
		{"Zlig"},
		{"ligZ"},
		{"Zligh"},
		{"lighZ"},
		{"Zlight"},
		{"lightZ"},
		{"ZlightG"},
		{"lightGZ"},
		{"ZlightGr"},
		{"lightGrZ"},
		{"ZlightGre"},
		{"lightGreZ"},
		{"ZlightGree"},
		{"lightGreeZ"},
		{"ZlightGreen"},
		{"lightGreenZ"},
		{"ZlightGreen5"},
		{"lightGreen5Z"},
		{"ZlightGreen50"},
		{"lightGreen50Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.LightGreen50()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestLightGreen100Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"lightGreen100", 13},
		{"lightGreen100 ", 13},
		{"lightGreen100\n", 13},
		{"lightGreen100.", 13},
		{"lightGreen100:", 13},
		{"lightGreen100,", 13},
		{"lightGreen100\"", 13},
		{"lightGreen100(", 13},
		{"lightGreen100)", 13},
		{"lightGreen100[", 13},
		{"lightGreen100]", 13},
		{"lightGreen100// comment", 13},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.LightGreen100()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestLightGreen100Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_l"},
		{"l_"},
		{"_li"},
		{"li_"},
		{"_lig"},
		{"lig_"},
		{"_ligh"},
		{"ligh_"},
		{"_light"},
		{"light_"},
		{"_lightG"},
		{"lightG_"},
		{"_lightGr"},
		{"lightGr_"},
		{"_lightGre"},
		{"lightGre_"},
		{"_lightGree"},
		{"lightGree_"},
		{"_lightGreen"},
		{"lightGreen_"},
		{"_lightGreen1"},
		{"lightGreen1_"},
		{"_lightGreen10"},
		{"lightGreen10_"},
		{"_lightGreen100"},
		{"lightGreen100_"},
		{"9713l"},
		{"l9713"},
		{"9713li"},
		{"li9713"},
		{"9713lig"},
		{"lig9713"},
		{"9713ligh"},
		{"ligh9713"},
		{"9713light"},
		{"light9713"},
		{"9713lightG"},
		{"lightG9713"},
		{"9713lightGr"},
		{"lightGr9713"},
		{"9713lightGre"},
		{"lightGre9713"},
		{"9713lightGree"},
		{"lightGree9713"},
		{"9713lightGreen"},
		{"lightGreen9713"},
		{"9713lightGreen1"},
		{"lightGreen19713"},
		{"9713lightGreen10"},
		{"lightGreen109713"},
		{"9713lightGreen100"},
		{"lightGreen1009713"},
		{"Zl"},
		{"lZ"},
		{"Zli"},
		{"liZ"},
		{"Zlig"},
		{"ligZ"},
		{"Zligh"},
		{"lighZ"},
		{"Zlight"},
		{"lightZ"},
		{"ZlightG"},
		{"lightGZ"},
		{"ZlightGr"},
		{"lightGrZ"},
		{"ZlightGre"},
		{"lightGreZ"},
		{"ZlightGree"},
		{"lightGreeZ"},
		{"ZlightGreen"},
		{"lightGreenZ"},
		{"ZlightGreen1"},
		{"lightGreen1Z"},
		{"ZlightGreen10"},
		{"lightGreen10Z"},
		{"ZlightGreen100"},
		{"lightGreen100Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.LightGreen100()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestLightGreen200Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"lightGreen200", 13},
		{"lightGreen200 ", 13},
		{"lightGreen200\n", 13},
		{"lightGreen200.", 13},
		{"lightGreen200:", 13},
		{"lightGreen200,", 13},
		{"lightGreen200\"", 13},
		{"lightGreen200(", 13},
		{"lightGreen200)", 13},
		{"lightGreen200[", 13},
		{"lightGreen200]", 13},
		{"lightGreen200// comment", 13},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.LightGreen200()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestLightGreen200Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_l"},
		{"l_"},
		{"_li"},
		{"li_"},
		{"_lig"},
		{"lig_"},
		{"_ligh"},
		{"ligh_"},
		{"_light"},
		{"light_"},
		{"_lightG"},
		{"lightG_"},
		{"_lightGr"},
		{"lightGr_"},
		{"_lightGre"},
		{"lightGre_"},
		{"_lightGree"},
		{"lightGree_"},
		{"_lightGreen"},
		{"lightGreen_"},
		{"_lightGreen2"},
		{"lightGreen2_"},
		{"_lightGreen20"},
		{"lightGreen20_"},
		{"_lightGreen200"},
		{"lightGreen200_"},
		{"9713l"},
		{"l9713"},
		{"9713li"},
		{"li9713"},
		{"9713lig"},
		{"lig9713"},
		{"9713ligh"},
		{"ligh9713"},
		{"9713light"},
		{"light9713"},
		{"9713lightG"},
		{"lightG9713"},
		{"9713lightGr"},
		{"lightGr9713"},
		{"9713lightGre"},
		{"lightGre9713"},
		{"9713lightGree"},
		{"lightGree9713"},
		{"9713lightGreen"},
		{"lightGreen9713"},
		{"9713lightGreen2"},
		{"lightGreen29713"},
		{"9713lightGreen20"},
		{"lightGreen209713"},
		{"9713lightGreen200"},
		{"lightGreen2009713"},
		{"Zl"},
		{"lZ"},
		{"Zli"},
		{"liZ"},
		{"Zlig"},
		{"ligZ"},
		{"Zligh"},
		{"lighZ"},
		{"Zlight"},
		{"lightZ"},
		{"ZlightG"},
		{"lightGZ"},
		{"ZlightGr"},
		{"lightGrZ"},
		{"ZlightGre"},
		{"lightGreZ"},
		{"ZlightGree"},
		{"lightGreeZ"},
		{"ZlightGreen"},
		{"lightGreenZ"},
		{"ZlightGreen2"},
		{"lightGreen2Z"},
		{"ZlightGreen20"},
		{"lightGreen20Z"},
		{"ZlightGreen200"},
		{"lightGreen200Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.LightGreen200()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestLightGreen300Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"lightGreen300", 13},
		{"lightGreen300 ", 13},
		{"lightGreen300\n", 13},
		{"lightGreen300.", 13},
		{"lightGreen300:", 13},
		{"lightGreen300,", 13},
		{"lightGreen300\"", 13},
		{"lightGreen300(", 13},
		{"lightGreen300)", 13},
		{"lightGreen300[", 13},
		{"lightGreen300]", 13},
		{"lightGreen300// comment", 13},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.LightGreen300()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestLightGreen300Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_l"},
		{"l_"},
		{"_li"},
		{"li_"},
		{"_lig"},
		{"lig_"},
		{"_ligh"},
		{"ligh_"},
		{"_light"},
		{"light_"},
		{"_lightG"},
		{"lightG_"},
		{"_lightGr"},
		{"lightGr_"},
		{"_lightGre"},
		{"lightGre_"},
		{"_lightGree"},
		{"lightGree_"},
		{"_lightGreen"},
		{"lightGreen_"},
		{"_lightGreen3"},
		{"lightGreen3_"},
		{"_lightGreen30"},
		{"lightGreen30_"},
		{"_lightGreen300"},
		{"lightGreen300_"},
		{"9713l"},
		{"l9713"},
		{"9713li"},
		{"li9713"},
		{"9713lig"},
		{"lig9713"},
		{"9713ligh"},
		{"ligh9713"},
		{"9713light"},
		{"light9713"},
		{"9713lightG"},
		{"lightG9713"},
		{"9713lightGr"},
		{"lightGr9713"},
		{"9713lightGre"},
		{"lightGre9713"},
		{"9713lightGree"},
		{"lightGree9713"},
		{"9713lightGreen"},
		{"lightGreen9713"},
		{"9713lightGreen3"},
		{"lightGreen39713"},
		{"9713lightGreen30"},
		{"lightGreen309713"},
		{"9713lightGreen300"},
		{"lightGreen3009713"},
		{"Zl"},
		{"lZ"},
		{"Zli"},
		{"liZ"},
		{"Zlig"},
		{"ligZ"},
		{"Zligh"},
		{"lighZ"},
		{"Zlight"},
		{"lightZ"},
		{"ZlightG"},
		{"lightGZ"},
		{"ZlightGr"},
		{"lightGrZ"},
		{"ZlightGre"},
		{"lightGreZ"},
		{"ZlightGree"},
		{"lightGreeZ"},
		{"ZlightGreen"},
		{"lightGreenZ"},
		{"ZlightGreen3"},
		{"lightGreen3Z"},
		{"ZlightGreen30"},
		{"lightGreen30Z"},
		{"ZlightGreen300"},
		{"lightGreen300Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.LightGreen300()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestLightGreen400Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"lightGreen400", 13},
		{"lightGreen400 ", 13},
		{"lightGreen400\n", 13},
		{"lightGreen400.", 13},
		{"lightGreen400:", 13},
		{"lightGreen400,", 13},
		{"lightGreen400\"", 13},
		{"lightGreen400(", 13},
		{"lightGreen400)", 13},
		{"lightGreen400[", 13},
		{"lightGreen400]", 13},
		{"lightGreen400// comment", 13},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.LightGreen400()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestLightGreen400Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_l"},
		{"l_"},
		{"_li"},
		{"li_"},
		{"_lig"},
		{"lig_"},
		{"_ligh"},
		{"ligh_"},
		{"_light"},
		{"light_"},
		{"_lightG"},
		{"lightG_"},
		{"_lightGr"},
		{"lightGr_"},
		{"_lightGre"},
		{"lightGre_"},
		{"_lightGree"},
		{"lightGree_"},
		{"_lightGreen"},
		{"lightGreen_"},
		{"_lightGreen4"},
		{"lightGreen4_"},
		{"_lightGreen40"},
		{"lightGreen40_"},
		{"_lightGreen400"},
		{"lightGreen400_"},
		{"9713l"},
		{"l9713"},
		{"9713li"},
		{"li9713"},
		{"9713lig"},
		{"lig9713"},
		{"9713ligh"},
		{"ligh9713"},
		{"9713light"},
		{"light9713"},
		{"9713lightG"},
		{"lightG9713"},
		{"9713lightGr"},
		{"lightGr9713"},
		{"9713lightGre"},
		{"lightGre9713"},
		{"9713lightGree"},
		{"lightGree9713"},
		{"9713lightGreen"},
		{"lightGreen9713"},
		{"9713lightGreen4"},
		{"lightGreen49713"},
		{"9713lightGreen40"},
		{"lightGreen409713"},
		{"9713lightGreen400"},
		{"lightGreen4009713"},
		{"Zl"},
		{"lZ"},
		{"Zli"},
		{"liZ"},
		{"Zlig"},
		{"ligZ"},
		{"Zligh"},
		{"lighZ"},
		{"Zlight"},
		{"lightZ"},
		{"ZlightG"},
		{"lightGZ"},
		{"ZlightGr"},
		{"lightGrZ"},
		{"ZlightGre"},
		{"lightGreZ"},
		{"ZlightGree"},
		{"lightGreeZ"},
		{"ZlightGreen"},
		{"lightGreenZ"},
		{"ZlightGreen4"},
		{"lightGreen4Z"},
		{"ZlightGreen40"},
		{"lightGreen40Z"},
		{"ZlightGreen400"},
		{"lightGreen400Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.LightGreen400()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestLightGreen500Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"lightGreen500", 13},
		{"lightGreen500 ", 13},
		{"lightGreen500\n", 13},
		{"lightGreen500.", 13},
		{"lightGreen500:", 13},
		{"lightGreen500,", 13},
		{"lightGreen500\"", 13},
		{"lightGreen500(", 13},
		{"lightGreen500)", 13},
		{"lightGreen500[", 13},
		{"lightGreen500]", 13},
		{"lightGreen500// comment", 13},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.LightGreen500()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestLightGreen500Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_l"},
		{"l_"},
		{"_li"},
		{"li_"},
		{"_lig"},
		{"lig_"},
		{"_ligh"},
		{"ligh_"},
		{"_light"},
		{"light_"},
		{"_lightG"},
		{"lightG_"},
		{"_lightGr"},
		{"lightGr_"},
		{"_lightGre"},
		{"lightGre_"},
		{"_lightGree"},
		{"lightGree_"},
		{"_lightGreen"},
		{"lightGreen_"},
		{"_lightGreen5"},
		{"lightGreen5_"},
		{"_lightGreen50"},
		{"lightGreen50_"},
		{"_lightGreen500"},
		{"lightGreen500_"},
		{"9713l"},
		{"l9713"},
		{"9713li"},
		{"li9713"},
		{"9713lig"},
		{"lig9713"},
		{"9713ligh"},
		{"ligh9713"},
		{"9713light"},
		{"light9713"},
		{"9713lightG"},
		{"lightG9713"},
		{"9713lightGr"},
		{"lightGr9713"},
		{"9713lightGre"},
		{"lightGre9713"},
		{"9713lightGree"},
		{"lightGree9713"},
		{"9713lightGreen"},
		{"lightGreen9713"},
		{"9713lightGreen5"},
		{"lightGreen59713"},
		{"9713lightGreen50"},
		{"lightGreen509713"},
		{"9713lightGreen500"},
		{"lightGreen5009713"},
		{"Zl"},
		{"lZ"},
		{"Zli"},
		{"liZ"},
		{"Zlig"},
		{"ligZ"},
		{"Zligh"},
		{"lighZ"},
		{"Zlight"},
		{"lightZ"},
		{"ZlightG"},
		{"lightGZ"},
		{"ZlightGr"},
		{"lightGrZ"},
		{"ZlightGre"},
		{"lightGreZ"},
		{"ZlightGree"},
		{"lightGreeZ"},
		{"ZlightGreen"},
		{"lightGreenZ"},
		{"ZlightGreen5"},
		{"lightGreen5Z"},
		{"ZlightGreen50"},
		{"lightGreen50Z"},
		{"ZlightGreen500"},
		{"lightGreen500Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.LightGreen500()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestLightGreen600Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"lightGreen600", 13},
		{"lightGreen600 ", 13},
		{"lightGreen600\n", 13},
		{"lightGreen600.", 13},
		{"lightGreen600:", 13},
		{"lightGreen600,", 13},
		{"lightGreen600\"", 13},
		{"lightGreen600(", 13},
		{"lightGreen600)", 13},
		{"lightGreen600[", 13},
		{"lightGreen600]", 13},
		{"lightGreen600// comment", 13},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.LightGreen600()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestLightGreen600Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_l"},
		{"l_"},
		{"_li"},
		{"li_"},
		{"_lig"},
		{"lig_"},
		{"_ligh"},
		{"ligh_"},
		{"_light"},
		{"light_"},
		{"_lightG"},
		{"lightG_"},
		{"_lightGr"},
		{"lightGr_"},
		{"_lightGre"},
		{"lightGre_"},
		{"_lightGree"},
		{"lightGree_"},
		{"_lightGreen"},
		{"lightGreen_"},
		{"_lightGreen6"},
		{"lightGreen6_"},
		{"_lightGreen60"},
		{"lightGreen60_"},
		{"_lightGreen600"},
		{"lightGreen600_"},
		{"9713l"},
		{"l9713"},
		{"9713li"},
		{"li9713"},
		{"9713lig"},
		{"lig9713"},
		{"9713ligh"},
		{"ligh9713"},
		{"9713light"},
		{"light9713"},
		{"9713lightG"},
		{"lightG9713"},
		{"9713lightGr"},
		{"lightGr9713"},
		{"9713lightGre"},
		{"lightGre9713"},
		{"9713lightGree"},
		{"lightGree9713"},
		{"9713lightGreen"},
		{"lightGreen9713"},
		{"9713lightGreen6"},
		{"lightGreen69713"},
		{"9713lightGreen60"},
		{"lightGreen609713"},
		{"9713lightGreen600"},
		{"lightGreen6009713"},
		{"Zl"},
		{"lZ"},
		{"Zli"},
		{"liZ"},
		{"Zlig"},
		{"ligZ"},
		{"Zligh"},
		{"lighZ"},
		{"Zlight"},
		{"lightZ"},
		{"ZlightG"},
		{"lightGZ"},
		{"ZlightGr"},
		{"lightGrZ"},
		{"ZlightGre"},
		{"lightGreZ"},
		{"ZlightGree"},
		{"lightGreeZ"},
		{"ZlightGreen"},
		{"lightGreenZ"},
		{"ZlightGreen6"},
		{"lightGreen6Z"},
		{"ZlightGreen60"},
		{"lightGreen60Z"},
		{"ZlightGreen600"},
		{"lightGreen600Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.LightGreen600()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestLightGreen700Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"lightGreen700", 13},
		{"lightGreen700 ", 13},
		{"lightGreen700\n", 13},
		{"lightGreen700.", 13},
		{"lightGreen700:", 13},
		{"lightGreen700,", 13},
		{"lightGreen700\"", 13},
		{"lightGreen700(", 13},
		{"lightGreen700)", 13},
		{"lightGreen700[", 13},
		{"lightGreen700]", 13},
		{"lightGreen700// comment", 13},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.LightGreen700()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestLightGreen700Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_l"},
		{"l_"},
		{"_li"},
		{"li_"},
		{"_lig"},
		{"lig_"},
		{"_ligh"},
		{"ligh_"},
		{"_light"},
		{"light_"},
		{"_lightG"},
		{"lightG_"},
		{"_lightGr"},
		{"lightGr_"},
		{"_lightGre"},
		{"lightGre_"},
		{"_lightGree"},
		{"lightGree_"},
		{"_lightGreen"},
		{"lightGreen_"},
		{"_lightGreen7"},
		{"lightGreen7_"},
		{"_lightGreen70"},
		{"lightGreen70_"},
		{"_lightGreen700"},
		{"lightGreen700_"},
		{"9713l"},
		{"l9713"},
		{"9713li"},
		{"li9713"},
		{"9713lig"},
		{"lig9713"},
		{"9713ligh"},
		{"ligh9713"},
		{"9713light"},
		{"light9713"},
		{"9713lightG"},
		{"lightG9713"},
		{"9713lightGr"},
		{"lightGr9713"},
		{"9713lightGre"},
		{"lightGre9713"},
		{"9713lightGree"},
		{"lightGree9713"},
		{"9713lightGreen"},
		{"lightGreen9713"},
		{"9713lightGreen7"},
		{"lightGreen79713"},
		{"9713lightGreen70"},
		{"lightGreen709713"},
		{"9713lightGreen700"},
		{"lightGreen7009713"},
		{"Zl"},
		{"lZ"},
		{"Zli"},
		{"liZ"},
		{"Zlig"},
		{"ligZ"},
		{"Zligh"},
		{"lighZ"},
		{"Zlight"},
		{"lightZ"},
		{"ZlightG"},
		{"lightGZ"},
		{"ZlightGr"},
		{"lightGrZ"},
		{"ZlightGre"},
		{"lightGreZ"},
		{"ZlightGree"},
		{"lightGreeZ"},
		{"ZlightGreen"},
		{"lightGreenZ"},
		{"ZlightGreen7"},
		{"lightGreen7Z"},
		{"ZlightGreen70"},
		{"lightGreen70Z"},
		{"ZlightGreen700"},
		{"lightGreen700Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.LightGreen700()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestLightGreen800Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"lightGreen800", 13},
		{"lightGreen800 ", 13},
		{"lightGreen800\n", 13},
		{"lightGreen800.", 13},
		{"lightGreen800:", 13},
		{"lightGreen800,", 13},
		{"lightGreen800\"", 13},
		{"lightGreen800(", 13},
		{"lightGreen800)", 13},
		{"lightGreen800[", 13},
		{"lightGreen800]", 13},
		{"lightGreen800// comment", 13},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.LightGreen800()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestLightGreen800Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_l"},
		{"l_"},
		{"_li"},
		{"li_"},
		{"_lig"},
		{"lig_"},
		{"_ligh"},
		{"ligh_"},
		{"_light"},
		{"light_"},
		{"_lightG"},
		{"lightG_"},
		{"_lightGr"},
		{"lightGr_"},
		{"_lightGre"},
		{"lightGre_"},
		{"_lightGree"},
		{"lightGree_"},
		{"_lightGreen"},
		{"lightGreen_"},
		{"_lightGreen8"},
		{"lightGreen8_"},
		{"_lightGreen80"},
		{"lightGreen80_"},
		{"_lightGreen800"},
		{"lightGreen800_"},
		{"9713l"},
		{"l9713"},
		{"9713li"},
		{"li9713"},
		{"9713lig"},
		{"lig9713"},
		{"9713ligh"},
		{"ligh9713"},
		{"9713light"},
		{"light9713"},
		{"9713lightG"},
		{"lightG9713"},
		{"9713lightGr"},
		{"lightGr9713"},
		{"9713lightGre"},
		{"lightGre9713"},
		{"9713lightGree"},
		{"lightGree9713"},
		{"9713lightGreen"},
		{"lightGreen9713"},
		{"9713lightGreen8"},
		{"lightGreen89713"},
		{"9713lightGreen80"},
		{"lightGreen809713"},
		{"9713lightGreen800"},
		{"lightGreen8009713"},
		{"Zl"},
		{"lZ"},
		{"Zli"},
		{"liZ"},
		{"Zlig"},
		{"ligZ"},
		{"Zligh"},
		{"lighZ"},
		{"Zlight"},
		{"lightZ"},
		{"ZlightG"},
		{"lightGZ"},
		{"ZlightGr"},
		{"lightGrZ"},
		{"ZlightGre"},
		{"lightGreZ"},
		{"ZlightGree"},
		{"lightGreeZ"},
		{"ZlightGreen"},
		{"lightGreenZ"},
		{"ZlightGreen8"},
		{"lightGreen8Z"},
		{"ZlightGreen80"},
		{"lightGreen80Z"},
		{"ZlightGreen800"},
		{"lightGreen800Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.LightGreen800()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestLightGreen900Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"lightGreen900", 13},
		{"lightGreen900 ", 13},
		{"lightGreen900\n", 13},
		{"lightGreen900.", 13},
		{"lightGreen900:", 13},
		{"lightGreen900,", 13},
		{"lightGreen900\"", 13},
		{"lightGreen900(", 13},
		{"lightGreen900)", 13},
		{"lightGreen900[", 13},
		{"lightGreen900]", 13},
		{"lightGreen900// comment", 13},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.LightGreen900()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestLightGreen900Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_l"},
		{"l_"},
		{"_li"},
		{"li_"},
		{"_lig"},
		{"lig_"},
		{"_ligh"},
		{"ligh_"},
		{"_light"},
		{"light_"},
		{"_lightG"},
		{"lightG_"},
		{"_lightGr"},
		{"lightGr_"},
		{"_lightGre"},
		{"lightGre_"},
		{"_lightGree"},
		{"lightGree_"},
		{"_lightGreen"},
		{"lightGreen_"},
		{"_lightGreen9"},
		{"lightGreen9_"},
		{"_lightGreen90"},
		{"lightGreen90_"},
		{"_lightGreen900"},
		{"lightGreen900_"},
		{"9713l"},
		{"l9713"},
		{"9713li"},
		{"li9713"},
		{"9713lig"},
		{"lig9713"},
		{"9713ligh"},
		{"ligh9713"},
		{"9713light"},
		{"light9713"},
		{"9713lightG"},
		{"lightG9713"},
		{"9713lightGr"},
		{"lightGr9713"},
		{"9713lightGre"},
		{"lightGre9713"},
		{"9713lightGree"},
		{"lightGree9713"},
		{"9713lightGreen"},
		{"lightGreen9713"},
		{"9713lightGreen9"},
		{"lightGreen99713"},
		{"9713lightGreen90"},
		{"lightGreen909713"},
		{"9713lightGreen900"},
		{"lightGreen9009713"},
		{"Zl"},
		{"lZ"},
		{"Zli"},
		{"liZ"},
		{"Zlig"},
		{"ligZ"},
		{"Zligh"},
		{"lighZ"},
		{"Zlight"},
		{"lightZ"},
		{"ZlightG"},
		{"lightGZ"},
		{"ZlightGr"},
		{"lightGrZ"},
		{"ZlightGre"},
		{"lightGreZ"},
		{"ZlightGree"},
		{"lightGreeZ"},
		{"ZlightGreen"},
		{"lightGreenZ"},
		{"ZlightGreen9"},
		{"lightGreen9Z"},
		{"ZlightGreen90"},
		{"lightGreen90Z"},
		{"ZlightGreen900"},
		{"lightGreen900Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.LightGreen900()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestLightGreenA100Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"lightGreenA100", 14},
		{"lightGreenA100 ", 14},
		{"lightGreenA100\n", 14},
		{"lightGreenA100.", 14},
		{"lightGreenA100:", 14},
		{"lightGreenA100,", 14},
		{"lightGreenA100\"", 14},
		{"lightGreenA100(", 14},
		{"lightGreenA100)", 14},
		{"lightGreenA100[", 14},
		{"lightGreenA100]", 14},
		{"lightGreenA100// comment", 14},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.LightGreenA100()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestLightGreenA100Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_l"},
		{"l_"},
		{"_li"},
		{"li_"},
		{"_lig"},
		{"lig_"},
		{"_ligh"},
		{"ligh_"},
		{"_light"},
		{"light_"},
		{"_lightG"},
		{"lightG_"},
		{"_lightGr"},
		{"lightGr_"},
		{"_lightGre"},
		{"lightGre_"},
		{"_lightGree"},
		{"lightGree_"},
		{"_lightGreen"},
		{"lightGreen_"},
		{"_lightGreenA"},
		{"lightGreenA_"},
		{"_lightGreenA1"},
		{"lightGreenA1_"},
		{"_lightGreenA10"},
		{"lightGreenA10_"},
		{"_lightGreenA100"},
		{"lightGreenA100_"},
		{"9713l"},
		{"l9713"},
		{"9713li"},
		{"li9713"},
		{"9713lig"},
		{"lig9713"},
		{"9713ligh"},
		{"ligh9713"},
		{"9713light"},
		{"light9713"},
		{"9713lightG"},
		{"lightG9713"},
		{"9713lightGr"},
		{"lightGr9713"},
		{"9713lightGre"},
		{"lightGre9713"},
		{"9713lightGree"},
		{"lightGree9713"},
		{"9713lightGreen"},
		{"lightGreen9713"},
		{"9713lightGreenA"},
		{"lightGreenA9713"},
		{"9713lightGreenA1"},
		{"lightGreenA19713"},
		{"9713lightGreenA10"},
		{"lightGreenA109713"},
		{"9713lightGreenA100"},
		{"lightGreenA1009713"},
		{"Zl"},
		{"lZ"},
		{"Zli"},
		{"liZ"},
		{"Zlig"},
		{"ligZ"},
		{"Zligh"},
		{"lighZ"},
		{"Zlight"},
		{"lightZ"},
		{"ZlightG"},
		{"lightGZ"},
		{"ZlightGr"},
		{"lightGrZ"},
		{"ZlightGre"},
		{"lightGreZ"},
		{"ZlightGree"},
		{"lightGreeZ"},
		{"ZlightGreen"},
		{"lightGreenZ"},
		{"ZlightGreenA"},
		{"lightGreenAZ"},
		{"ZlightGreenA1"},
		{"lightGreenA1Z"},
		{"ZlightGreenA10"},
		{"lightGreenA10Z"},
		{"ZlightGreenA100"},
		{"lightGreenA100Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.LightGreenA100()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestLightGreenA200Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"lightGreenA200", 14},
		{"lightGreenA200 ", 14},
		{"lightGreenA200\n", 14},
		{"lightGreenA200.", 14},
		{"lightGreenA200:", 14},
		{"lightGreenA200,", 14},
		{"lightGreenA200\"", 14},
		{"lightGreenA200(", 14},
		{"lightGreenA200)", 14},
		{"lightGreenA200[", 14},
		{"lightGreenA200]", 14},
		{"lightGreenA200// comment", 14},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.LightGreenA200()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestLightGreenA200Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_l"},
		{"l_"},
		{"_li"},
		{"li_"},
		{"_lig"},
		{"lig_"},
		{"_ligh"},
		{"ligh_"},
		{"_light"},
		{"light_"},
		{"_lightG"},
		{"lightG_"},
		{"_lightGr"},
		{"lightGr_"},
		{"_lightGre"},
		{"lightGre_"},
		{"_lightGree"},
		{"lightGree_"},
		{"_lightGreen"},
		{"lightGreen_"},
		{"_lightGreenA"},
		{"lightGreenA_"},
		{"_lightGreenA2"},
		{"lightGreenA2_"},
		{"_lightGreenA20"},
		{"lightGreenA20_"},
		{"_lightGreenA200"},
		{"lightGreenA200_"},
		{"9713l"},
		{"l9713"},
		{"9713li"},
		{"li9713"},
		{"9713lig"},
		{"lig9713"},
		{"9713ligh"},
		{"ligh9713"},
		{"9713light"},
		{"light9713"},
		{"9713lightG"},
		{"lightG9713"},
		{"9713lightGr"},
		{"lightGr9713"},
		{"9713lightGre"},
		{"lightGre9713"},
		{"9713lightGree"},
		{"lightGree9713"},
		{"9713lightGreen"},
		{"lightGreen9713"},
		{"9713lightGreenA"},
		{"lightGreenA9713"},
		{"9713lightGreenA2"},
		{"lightGreenA29713"},
		{"9713lightGreenA20"},
		{"lightGreenA209713"},
		{"9713lightGreenA200"},
		{"lightGreenA2009713"},
		{"Zl"},
		{"lZ"},
		{"Zli"},
		{"liZ"},
		{"Zlig"},
		{"ligZ"},
		{"Zligh"},
		{"lighZ"},
		{"Zlight"},
		{"lightZ"},
		{"ZlightG"},
		{"lightGZ"},
		{"ZlightGr"},
		{"lightGrZ"},
		{"ZlightGre"},
		{"lightGreZ"},
		{"ZlightGree"},
		{"lightGreeZ"},
		{"ZlightGreen"},
		{"lightGreenZ"},
		{"ZlightGreenA"},
		{"lightGreenAZ"},
		{"ZlightGreenA2"},
		{"lightGreenA2Z"},
		{"ZlightGreenA20"},
		{"lightGreenA20Z"},
		{"ZlightGreenA200"},
		{"lightGreenA200Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.LightGreenA200()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestLightGreenA400Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"lightGreenA400", 14},
		{"lightGreenA400 ", 14},
		{"lightGreenA400\n", 14},
		{"lightGreenA400.", 14},
		{"lightGreenA400:", 14},
		{"lightGreenA400,", 14},
		{"lightGreenA400\"", 14},
		{"lightGreenA400(", 14},
		{"lightGreenA400)", 14},
		{"lightGreenA400[", 14},
		{"lightGreenA400]", 14},
		{"lightGreenA400// comment", 14},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.LightGreenA400()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestLightGreenA400Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_l"},
		{"l_"},
		{"_li"},
		{"li_"},
		{"_lig"},
		{"lig_"},
		{"_ligh"},
		{"ligh_"},
		{"_light"},
		{"light_"},
		{"_lightG"},
		{"lightG_"},
		{"_lightGr"},
		{"lightGr_"},
		{"_lightGre"},
		{"lightGre_"},
		{"_lightGree"},
		{"lightGree_"},
		{"_lightGreen"},
		{"lightGreen_"},
		{"_lightGreenA"},
		{"lightGreenA_"},
		{"_lightGreenA4"},
		{"lightGreenA4_"},
		{"_lightGreenA40"},
		{"lightGreenA40_"},
		{"_lightGreenA400"},
		{"lightGreenA400_"},
		{"9713l"},
		{"l9713"},
		{"9713li"},
		{"li9713"},
		{"9713lig"},
		{"lig9713"},
		{"9713ligh"},
		{"ligh9713"},
		{"9713light"},
		{"light9713"},
		{"9713lightG"},
		{"lightG9713"},
		{"9713lightGr"},
		{"lightGr9713"},
		{"9713lightGre"},
		{"lightGre9713"},
		{"9713lightGree"},
		{"lightGree9713"},
		{"9713lightGreen"},
		{"lightGreen9713"},
		{"9713lightGreenA"},
		{"lightGreenA9713"},
		{"9713lightGreenA4"},
		{"lightGreenA49713"},
		{"9713lightGreenA40"},
		{"lightGreenA409713"},
		{"9713lightGreenA400"},
		{"lightGreenA4009713"},
		{"Zl"},
		{"lZ"},
		{"Zli"},
		{"liZ"},
		{"Zlig"},
		{"ligZ"},
		{"Zligh"},
		{"lighZ"},
		{"Zlight"},
		{"lightZ"},
		{"ZlightG"},
		{"lightGZ"},
		{"ZlightGr"},
		{"lightGrZ"},
		{"ZlightGre"},
		{"lightGreZ"},
		{"ZlightGree"},
		{"lightGreeZ"},
		{"ZlightGreen"},
		{"lightGreenZ"},
		{"ZlightGreenA"},
		{"lightGreenAZ"},
		{"ZlightGreenA4"},
		{"lightGreenA4Z"},
		{"ZlightGreenA40"},
		{"lightGreenA40Z"},
		{"ZlightGreenA400"},
		{"lightGreenA400Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.LightGreenA400()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestLightGreenA700Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"lightGreenA700", 14},
		{"lightGreenA700 ", 14},
		{"lightGreenA700\n", 14},
		{"lightGreenA700.", 14},
		{"lightGreenA700:", 14},
		{"lightGreenA700,", 14},
		{"lightGreenA700\"", 14},
		{"lightGreenA700(", 14},
		{"lightGreenA700)", 14},
		{"lightGreenA700[", 14},
		{"lightGreenA700]", 14},
		{"lightGreenA700// comment", 14},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.LightGreenA700()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestLightGreenA700Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_l"},
		{"l_"},
		{"_li"},
		{"li_"},
		{"_lig"},
		{"lig_"},
		{"_ligh"},
		{"ligh_"},
		{"_light"},
		{"light_"},
		{"_lightG"},
		{"lightG_"},
		{"_lightGr"},
		{"lightGr_"},
		{"_lightGre"},
		{"lightGre_"},
		{"_lightGree"},
		{"lightGree_"},
		{"_lightGreen"},
		{"lightGreen_"},
		{"_lightGreenA"},
		{"lightGreenA_"},
		{"_lightGreenA7"},
		{"lightGreenA7_"},
		{"_lightGreenA70"},
		{"lightGreenA70_"},
		{"_lightGreenA700"},
		{"lightGreenA700_"},
		{"9713l"},
		{"l9713"},
		{"9713li"},
		{"li9713"},
		{"9713lig"},
		{"lig9713"},
		{"9713ligh"},
		{"ligh9713"},
		{"9713light"},
		{"light9713"},
		{"9713lightG"},
		{"lightG9713"},
		{"9713lightGr"},
		{"lightGr9713"},
		{"9713lightGre"},
		{"lightGre9713"},
		{"9713lightGree"},
		{"lightGree9713"},
		{"9713lightGreen"},
		{"lightGreen9713"},
		{"9713lightGreenA"},
		{"lightGreenA9713"},
		{"9713lightGreenA7"},
		{"lightGreenA79713"},
		{"9713lightGreenA70"},
		{"lightGreenA709713"},
		{"9713lightGreenA700"},
		{"lightGreenA7009713"},
		{"Zl"},
		{"lZ"},
		{"Zli"},
		{"liZ"},
		{"Zlig"},
		{"ligZ"},
		{"Zligh"},
		{"lighZ"},
		{"Zlight"},
		{"lightZ"},
		{"ZlightG"},
		{"lightGZ"},
		{"ZlightGr"},
		{"lightGrZ"},
		{"ZlightGre"},
		{"lightGreZ"},
		{"ZlightGree"},
		{"lightGreeZ"},
		{"ZlightGreen"},
		{"lightGreenZ"},
		{"ZlightGreenA"},
		{"lightGreenAZ"},
		{"ZlightGreenA7"},
		{"lightGreenA7Z"},
		{"ZlightGreenA70"},
		{"lightGreenA70Z"},
		{"ZlightGreenA700"},
		{"lightGreenA700Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.LightGreenA700()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestLime50Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"lime50", 6},
		{"lime50 ", 6},
		{"lime50\n", 6},
		{"lime50.", 6},
		{"lime50:", 6},
		{"lime50,", 6},
		{"lime50\"", 6},
		{"lime50(", 6},
		{"lime50)", 6},
		{"lime50[", 6},
		{"lime50]", 6},
		{"lime50// comment", 6},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Lime50()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestLime50Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_l"},
		{"l_"},
		{"_li"},
		{"li_"},
		{"_lim"},
		{"lim_"},
		{"_lime"},
		{"lime_"},
		{"_lime5"},
		{"lime5_"},
		{"_lime50"},
		{"lime50_"},
		{"9713l"},
		{"l9713"},
		{"9713li"},
		{"li9713"},
		{"9713lim"},
		{"lim9713"},
		{"9713lime"},
		{"lime9713"},
		{"9713lime5"},
		{"lime59713"},
		{"9713lime50"},
		{"lime509713"},
		{"Zl"},
		{"lZ"},
		{"Zli"},
		{"liZ"},
		{"Zlim"},
		{"limZ"},
		{"Zlime"},
		{"limeZ"},
		{"Zlime5"},
		{"lime5Z"},
		{"Zlime50"},
		{"lime50Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Lime50()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestLime100Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"lime100", 7},
		{"lime100 ", 7},
		{"lime100\n", 7},
		{"lime100.", 7},
		{"lime100:", 7},
		{"lime100,", 7},
		{"lime100\"", 7},
		{"lime100(", 7},
		{"lime100)", 7},
		{"lime100[", 7},
		{"lime100]", 7},
		{"lime100// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Lime100()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestLime100Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_l"},
		{"l_"},
		{"_li"},
		{"li_"},
		{"_lim"},
		{"lim_"},
		{"_lime"},
		{"lime_"},
		{"_lime1"},
		{"lime1_"},
		{"_lime10"},
		{"lime10_"},
		{"_lime100"},
		{"lime100_"},
		{"9713l"},
		{"l9713"},
		{"9713li"},
		{"li9713"},
		{"9713lim"},
		{"lim9713"},
		{"9713lime"},
		{"lime9713"},
		{"9713lime1"},
		{"lime19713"},
		{"9713lime10"},
		{"lime109713"},
		{"9713lime100"},
		{"lime1009713"},
		{"Zl"},
		{"lZ"},
		{"Zli"},
		{"liZ"},
		{"Zlim"},
		{"limZ"},
		{"Zlime"},
		{"limeZ"},
		{"Zlime1"},
		{"lime1Z"},
		{"Zlime10"},
		{"lime10Z"},
		{"Zlime100"},
		{"lime100Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Lime100()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestLime200Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"lime200", 7},
		{"lime200 ", 7},
		{"lime200\n", 7},
		{"lime200.", 7},
		{"lime200:", 7},
		{"lime200,", 7},
		{"lime200\"", 7},
		{"lime200(", 7},
		{"lime200)", 7},
		{"lime200[", 7},
		{"lime200]", 7},
		{"lime200// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Lime200()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestLime200Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_l"},
		{"l_"},
		{"_li"},
		{"li_"},
		{"_lim"},
		{"lim_"},
		{"_lime"},
		{"lime_"},
		{"_lime2"},
		{"lime2_"},
		{"_lime20"},
		{"lime20_"},
		{"_lime200"},
		{"lime200_"},
		{"9713l"},
		{"l9713"},
		{"9713li"},
		{"li9713"},
		{"9713lim"},
		{"lim9713"},
		{"9713lime"},
		{"lime9713"},
		{"9713lime2"},
		{"lime29713"},
		{"9713lime20"},
		{"lime209713"},
		{"9713lime200"},
		{"lime2009713"},
		{"Zl"},
		{"lZ"},
		{"Zli"},
		{"liZ"},
		{"Zlim"},
		{"limZ"},
		{"Zlime"},
		{"limeZ"},
		{"Zlime2"},
		{"lime2Z"},
		{"Zlime20"},
		{"lime20Z"},
		{"Zlime200"},
		{"lime200Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Lime200()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestLime300Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"lime300", 7},
		{"lime300 ", 7},
		{"lime300\n", 7},
		{"lime300.", 7},
		{"lime300:", 7},
		{"lime300,", 7},
		{"lime300\"", 7},
		{"lime300(", 7},
		{"lime300)", 7},
		{"lime300[", 7},
		{"lime300]", 7},
		{"lime300// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Lime300()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestLime300Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_l"},
		{"l_"},
		{"_li"},
		{"li_"},
		{"_lim"},
		{"lim_"},
		{"_lime"},
		{"lime_"},
		{"_lime3"},
		{"lime3_"},
		{"_lime30"},
		{"lime30_"},
		{"_lime300"},
		{"lime300_"},
		{"9713l"},
		{"l9713"},
		{"9713li"},
		{"li9713"},
		{"9713lim"},
		{"lim9713"},
		{"9713lime"},
		{"lime9713"},
		{"9713lime3"},
		{"lime39713"},
		{"9713lime30"},
		{"lime309713"},
		{"9713lime300"},
		{"lime3009713"},
		{"Zl"},
		{"lZ"},
		{"Zli"},
		{"liZ"},
		{"Zlim"},
		{"limZ"},
		{"Zlime"},
		{"limeZ"},
		{"Zlime3"},
		{"lime3Z"},
		{"Zlime30"},
		{"lime30Z"},
		{"Zlime300"},
		{"lime300Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Lime300()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestLime400Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"lime400", 7},
		{"lime400 ", 7},
		{"lime400\n", 7},
		{"lime400.", 7},
		{"lime400:", 7},
		{"lime400,", 7},
		{"lime400\"", 7},
		{"lime400(", 7},
		{"lime400)", 7},
		{"lime400[", 7},
		{"lime400]", 7},
		{"lime400// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Lime400()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestLime400Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_l"},
		{"l_"},
		{"_li"},
		{"li_"},
		{"_lim"},
		{"lim_"},
		{"_lime"},
		{"lime_"},
		{"_lime4"},
		{"lime4_"},
		{"_lime40"},
		{"lime40_"},
		{"_lime400"},
		{"lime400_"},
		{"9713l"},
		{"l9713"},
		{"9713li"},
		{"li9713"},
		{"9713lim"},
		{"lim9713"},
		{"9713lime"},
		{"lime9713"},
		{"9713lime4"},
		{"lime49713"},
		{"9713lime40"},
		{"lime409713"},
		{"9713lime400"},
		{"lime4009713"},
		{"Zl"},
		{"lZ"},
		{"Zli"},
		{"liZ"},
		{"Zlim"},
		{"limZ"},
		{"Zlime"},
		{"limeZ"},
		{"Zlime4"},
		{"lime4Z"},
		{"Zlime40"},
		{"lime40Z"},
		{"Zlime400"},
		{"lime400Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Lime400()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestLime500Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"lime500", 7},
		{"lime500 ", 7},
		{"lime500\n", 7},
		{"lime500.", 7},
		{"lime500:", 7},
		{"lime500,", 7},
		{"lime500\"", 7},
		{"lime500(", 7},
		{"lime500)", 7},
		{"lime500[", 7},
		{"lime500]", 7},
		{"lime500// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Lime500()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestLime500Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_l"},
		{"l_"},
		{"_li"},
		{"li_"},
		{"_lim"},
		{"lim_"},
		{"_lime"},
		{"lime_"},
		{"_lime5"},
		{"lime5_"},
		{"_lime50"},
		{"lime50_"},
		{"_lime500"},
		{"lime500_"},
		{"9713l"},
		{"l9713"},
		{"9713li"},
		{"li9713"},
		{"9713lim"},
		{"lim9713"},
		{"9713lime"},
		{"lime9713"},
		{"9713lime5"},
		{"lime59713"},
		{"9713lime50"},
		{"lime509713"},
		{"9713lime500"},
		{"lime5009713"},
		{"Zl"},
		{"lZ"},
		{"Zli"},
		{"liZ"},
		{"Zlim"},
		{"limZ"},
		{"Zlime"},
		{"limeZ"},
		{"Zlime5"},
		{"lime5Z"},
		{"Zlime50"},
		{"lime50Z"},
		{"Zlime500"},
		{"lime500Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Lime500()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestLime600Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"lime600", 7},
		{"lime600 ", 7},
		{"lime600\n", 7},
		{"lime600.", 7},
		{"lime600:", 7},
		{"lime600,", 7},
		{"lime600\"", 7},
		{"lime600(", 7},
		{"lime600)", 7},
		{"lime600[", 7},
		{"lime600]", 7},
		{"lime600// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Lime600()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestLime600Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_l"},
		{"l_"},
		{"_li"},
		{"li_"},
		{"_lim"},
		{"lim_"},
		{"_lime"},
		{"lime_"},
		{"_lime6"},
		{"lime6_"},
		{"_lime60"},
		{"lime60_"},
		{"_lime600"},
		{"lime600_"},
		{"9713l"},
		{"l9713"},
		{"9713li"},
		{"li9713"},
		{"9713lim"},
		{"lim9713"},
		{"9713lime"},
		{"lime9713"},
		{"9713lime6"},
		{"lime69713"},
		{"9713lime60"},
		{"lime609713"},
		{"9713lime600"},
		{"lime6009713"},
		{"Zl"},
		{"lZ"},
		{"Zli"},
		{"liZ"},
		{"Zlim"},
		{"limZ"},
		{"Zlime"},
		{"limeZ"},
		{"Zlime6"},
		{"lime6Z"},
		{"Zlime60"},
		{"lime60Z"},
		{"Zlime600"},
		{"lime600Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Lime600()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestLime700Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"lime700", 7},
		{"lime700 ", 7},
		{"lime700\n", 7},
		{"lime700.", 7},
		{"lime700:", 7},
		{"lime700,", 7},
		{"lime700\"", 7},
		{"lime700(", 7},
		{"lime700)", 7},
		{"lime700[", 7},
		{"lime700]", 7},
		{"lime700// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Lime700()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestLime700Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_l"},
		{"l_"},
		{"_li"},
		{"li_"},
		{"_lim"},
		{"lim_"},
		{"_lime"},
		{"lime_"},
		{"_lime7"},
		{"lime7_"},
		{"_lime70"},
		{"lime70_"},
		{"_lime700"},
		{"lime700_"},
		{"9713l"},
		{"l9713"},
		{"9713li"},
		{"li9713"},
		{"9713lim"},
		{"lim9713"},
		{"9713lime"},
		{"lime9713"},
		{"9713lime7"},
		{"lime79713"},
		{"9713lime70"},
		{"lime709713"},
		{"9713lime700"},
		{"lime7009713"},
		{"Zl"},
		{"lZ"},
		{"Zli"},
		{"liZ"},
		{"Zlim"},
		{"limZ"},
		{"Zlime"},
		{"limeZ"},
		{"Zlime7"},
		{"lime7Z"},
		{"Zlime70"},
		{"lime70Z"},
		{"Zlime700"},
		{"lime700Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Lime700()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestLime800Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"lime800", 7},
		{"lime800 ", 7},
		{"lime800\n", 7},
		{"lime800.", 7},
		{"lime800:", 7},
		{"lime800,", 7},
		{"lime800\"", 7},
		{"lime800(", 7},
		{"lime800)", 7},
		{"lime800[", 7},
		{"lime800]", 7},
		{"lime800// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Lime800()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestLime800Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_l"},
		{"l_"},
		{"_li"},
		{"li_"},
		{"_lim"},
		{"lim_"},
		{"_lime"},
		{"lime_"},
		{"_lime8"},
		{"lime8_"},
		{"_lime80"},
		{"lime80_"},
		{"_lime800"},
		{"lime800_"},
		{"9713l"},
		{"l9713"},
		{"9713li"},
		{"li9713"},
		{"9713lim"},
		{"lim9713"},
		{"9713lime"},
		{"lime9713"},
		{"9713lime8"},
		{"lime89713"},
		{"9713lime80"},
		{"lime809713"},
		{"9713lime800"},
		{"lime8009713"},
		{"Zl"},
		{"lZ"},
		{"Zli"},
		{"liZ"},
		{"Zlim"},
		{"limZ"},
		{"Zlime"},
		{"limeZ"},
		{"Zlime8"},
		{"lime8Z"},
		{"Zlime80"},
		{"lime80Z"},
		{"Zlime800"},
		{"lime800Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Lime800()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestLime900Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"lime900", 7},
		{"lime900 ", 7},
		{"lime900\n", 7},
		{"lime900.", 7},
		{"lime900:", 7},
		{"lime900,", 7},
		{"lime900\"", 7},
		{"lime900(", 7},
		{"lime900)", 7},
		{"lime900[", 7},
		{"lime900]", 7},
		{"lime900// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Lime900()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestLime900Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_l"},
		{"l_"},
		{"_li"},
		{"li_"},
		{"_lim"},
		{"lim_"},
		{"_lime"},
		{"lime_"},
		{"_lime9"},
		{"lime9_"},
		{"_lime90"},
		{"lime90_"},
		{"_lime900"},
		{"lime900_"},
		{"9713l"},
		{"l9713"},
		{"9713li"},
		{"li9713"},
		{"9713lim"},
		{"lim9713"},
		{"9713lime"},
		{"lime9713"},
		{"9713lime9"},
		{"lime99713"},
		{"9713lime90"},
		{"lime909713"},
		{"9713lime900"},
		{"lime9009713"},
		{"Zl"},
		{"lZ"},
		{"Zli"},
		{"liZ"},
		{"Zlim"},
		{"limZ"},
		{"Zlime"},
		{"limeZ"},
		{"Zlime9"},
		{"lime9Z"},
		{"Zlime90"},
		{"lime90Z"},
		{"Zlime900"},
		{"lime900Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Lime900()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestLimeA100Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"limeA100", 8},
		{"limeA100 ", 8},
		{"limeA100\n", 8},
		{"limeA100.", 8},
		{"limeA100:", 8},
		{"limeA100,", 8},
		{"limeA100\"", 8},
		{"limeA100(", 8},
		{"limeA100)", 8},
		{"limeA100[", 8},
		{"limeA100]", 8},
		{"limeA100// comment", 8},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.LimeA100()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestLimeA100Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_l"},
		{"l_"},
		{"_li"},
		{"li_"},
		{"_lim"},
		{"lim_"},
		{"_lime"},
		{"lime_"},
		{"_limeA"},
		{"limeA_"},
		{"_limeA1"},
		{"limeA1_"},
		{"_limeA10"},
		{"limeA10_"},
		{"_limeA100"},
		{"limeA100_"},
		{"9713l"},
		{"l9713"},
		{"9713li"},
		{"li9713"},
		{"9713lim"},
		{"lim9713"},
		{"9713lime"},
		{"lime9713"},
		{"9713limeA"},
		{"limeA9713"},
		{"9713limeA1"},
		{"limeA19713"},
		{"9713limeA10"},
		{"limeA109713"},
		{"9713limeA100"},
		{"limeA1009713"},
		{"Zl"},
		{"lZ"},
		{"Zli"},
		{"liZ"},
		{"Zlim"},
		{"limZ"},
		{"Zlime"},
		{"limeZ"},
		{"ZlimeA"},
		{"limeAZ"},
		{"ZlimeA1"},
		{"limeA1Z"},
		{"ZlimeA10"},
		{"limeA10Z"},
		{"ZlimeA100"},
		{"limeA100Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.LimeA100()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestLimeA200Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"limeA200", 8},
		{"limeA200 ", 8},
		{"limeA200\n", 8},
		{"limeA200.", 8},
		{"limeA200:", 8},
		{"limeA200,", 8},
		{"limeA200\"", 8},
		{"limeA200(", 8},
		{"limeA200)", 8},
		{"limeA200[", 8},
		{"limeA200]", 8},
		{"limeA200// comment", 8},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.LimeA200()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestLimeA200Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_l"},
		{"l_"},
		{"_li"},
		{"li_"},
		{"_lim"},
		{"lim_"},
		{"_lime"},
		{"lime_"},
		{"_limeA"},
		{"limeA_"},
		{"_limeA2"},
		{"limeA2_"},
		{"_limeA20"},
		{"limeA20_"},
		{"_limeA200"},
		{"limeA200_"},
		{"9713l"},
		{"l9713"},
		{"9713li"},
		{"li9713"},
		{"9713lim"},
		{"lim9713"},
		{"9713lime"},
		{"lime9713"},
		{"9713limeA"},
		{"limeA9713"},
		{"9713limeA2"},
		{"limeA29713"},
		{"9713limeA20"},
		{"limeA209713"},
		{"9713limeA200"},
		{"limeA2009713"},
		{"Zl"},
		{"lZ"},
		{"Zli"},
		{"liZ"},
		{"Zlim"},
		{"limZ"},
		{"Zlime"},
		{"limeZ"},
		{"ZlimeA"},
		{"limeAZ"},
		{"ZlimeA2"},
		{"limeA2Z"},
		{"ZlimeA20"},
		{"limeA20Z"},
		{"ZlimeA200"},
		{"limeA200Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.LimeA200()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestLimeA400Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"limeA400", 8},
		{"limeA400 ", 8},
		{"limeA400\n", 8},
		{"limeA400.", 8},
		{"limeA400:", 8},
		{"limeA400,", 8},
		{"limeA400\"", 8},
		{"limeA400(", 8},
		{"limeA400)", 8},
		{"limeA400[", 8},
		{"limeA400]", 8},
		{"limeA400// comment", 8},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.LimeA400()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestLimeA400Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_l"},
		{"l_"},
		{"_li"},
		{"li_"},
		{"_lim"},
		{"lim_"},
		{"_lime"},
		{"lime_"},
		{"_limeA"},
		{"limeA_"},
		{"_limeA4"},
		{"limeA4_"},
		{"_limeA40"},
		{"limeA40_"},
		{"_limeA400"},
		{"limeA400_"},
		{"9713l"},
		{"l9713"},
		{"9713li"},
		{"li9713"},
		{"9713lim"},
		{"lim9713"},
		{"9713lime"},
		{"lime9713"},
		{"9713limeA"},
		{"limeA9713"},
		{"9713limeA4"},
		{"limeA49713"},
		{"9713limeA40"},
		{"limeA409713"},
		{"9713limeA400"},
		{"limeA4009713"},
		{"Zl"},
		{"lZ"},
		{"Zli"},
		{"liZ"},
		{"Zlim"},
		{"limZ"},
		{"Zlime"},
		{"limeZ"},
		{"ZlimeA"},
		{"limeAZ"},
		{"ZlimeA4"},
		{"limeA4Z"},
		{"ZlimeA40"},
		{"limeA40Z"},
		{"ZlimeA400"},
		{"limeA400Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.LimeA400()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestLimeA700Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"limeA700", 8},
		{"limeA700 ", 8},
		{"limeA700\n", 8},
		{"limeA700.", 8},
		{"limeA700:", 8},
		{"limeA700,", 8},
		{"limeA700\"", 8},
		{"limeA700(", 8},
		{"limeA700)", 8},
		{"limeA700[", 8},
		{"limeA700]", 8},
		{"limeA700// comment", 8},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.LimeA700()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestLimeA700Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_l"},
		{"l_"},
		{"_li"},
		{"li_"},
		{"_lim"},
		{"lim_"},
		{"_lime"},
		{"lime_"},
		{"_limeA"},
		{"limeA_"},
		{"_limeA7"},
		{"limeA7_"},
		{"_limeA70"},
		{"limeA70_"},
		{"_limeA700"},
		{"limeA700_"},
		{"9713l"},
		{"l9713"},
		{"9713li"},
		{"li9713"},
		{"9713lim"},
		{"lim9713"},
		{"9713lime"},
		{"lime9713"},
		{"9713limeA"},
		{"limeA9713"},
		{"9713limeA7"},
		{"limeA79713"},
		{"9713limeA70"},
		{"limeA709713"},
		{"9713limeA700"},
		{"limeA7009713"},
		{"Zl"},
		{"lZ"},
		{"Zli"},
		{"liZ"},
		{"Zlim"},
		{"limZ"},
		{"Zlime"},
		{"limeZ"},
		{"ZlimeA"},
		{"limeAZ"},
		{"ZlimeA7"},
		{"limeA7Z"},
		{"ZlimeA70"},
		{"limeA70Z"},
		{"ZlimeA700"},
		{"limeA700Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.LimeA700()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestYellow50Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"yellow50", 8},
		{"yellow50 ", 8},
		{"yellow50\n", 8},
		{"yellow50.", 8},
		{"yellow50:", 8},
		{"yellow50,", 8},
		{"yellow50\"", 8},
		{"yellow50(", 8},
		{"yellow50)", 8},
		{"yellow50[", 8},
		{"yellow50]", 8},
		{"yellow50// comment", 8},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Yellow50()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestYellow50Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_y"},
		{"y_"},
		{"_ye"},
		{"ye_"},
		{"_yel"},
		{"yel_"},
		{"_yell"},
		{"yell_"},
		{"_yello"},
		{"yello_"},
		{"_yellow"},
		{"yellow_"},
		{"_yellow5"},
		{"yellow5_"},
		{"_yellow50"},
		{"yellow50_"},
		{"9713y"},
		{"y9713"},
		{"9713ye"},
		{"ye9713"},
		{"9713yel"},
		{"yel9713"},
		{"9713yell"},
		{"yell9713"},
		{"9713yello"},
		{"yello9713"},
		{"9713yellow"},
		{"yellow9713"},
		{"9713yellow5"},
		{"yellow59713"},
		{"9713yellow50"},
		{"yellow509713"},
		{"Zy"},
		{"yZ"},
		{"Zye"},
		{"yeZ"},
		{"Zyel"},
		{"yelZ"},
		{"Zyell"},
		{"yellZ"},
		{"Zyello"},
		{"yelloZ"},
		{"Zyellow"},
		{"yellowZ"},
		{"Zyellow5"},
		{"yellow5Z"},
		{"Zyellow50"},
		{"yellow50Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Yellow50()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestYellow100Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"yellow100", 9},
		{"yellow100 ", 9},
		{"yellow100\n", 9},
		{"yellow100.", 9},
		{"yellow100:", 9},
		{"yellow100,", 9},
		{"yellow100\"", 9},
		{"yellow100(", 9},
		{"yellow100)", 9},
		{"yellow100[", 9},
		{"yellow100]", 9},
		{"yellow100// comment", 9},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Yellow100()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestYellow100Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_y"},
		{"y_"},
		{"_ye"},
		{"ye_"},
		{"_yel"},
		{"yel_"},
		{"_yell"},
		{"yell_"},
		{"_yello"},
		{"yello_"},
		{"_yellow"},
		{"yellow_"},
		{"_yellow1"},
		{"yellow1_"},
		{"_yellow10"},
		{"yellow10_"},
		{"_yellow100"},
		{"yellow100_"},
		{"9713y"},
		{"y9713"},
		{"9713ye"},
		{"ye9713"},
		{"9713yel"},
		{"yel9713"},
		{"9713yell"},
		{"yell9713"},
		{"9713yello"},
		{"yello9713"},
		{"9713yellow"},
		{"yellow9713"},
		{"9713yellow1"},
		{"yellow19713"},
		{"9713yellow10"},
		{"yellow109713"},
		{"9713yellow100"},
		{"yellow1009713"},
		{"Zy"},
		{"yZ"},
		{"Zye"},
		{"yeZ"},
		{"Zyel"},
		{"yelZ"},
		{"Zyell"},
		{"yellZ"},
		{"Zyello"},
		{"yelloZ"},
		{"Zyellow"},
		{"yellowZ"},
		{"Zyellow1"},
		{"yellow1Z"},
		{"Zyellow10"},
		{"yellow10Z"},
		{"Zyellow100"},
		{"yellow100Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Yellow100()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestYellow200Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"yellow200", 9},
		{"yellow200 ", 9},
		{"yellow200\n", 9},
		{"yellow200.", 9},
		{"yellow200:", 9},
		{"yellow200,", 9},
		{"yellow200\"", 9},
		{"yellow200(", 9},
		{"yellow200)", 9},
		{"yellow200[", 9},
		{"yellow200]", 9},
		{"yellow200// comment", 9},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Yellow200()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestYellow200Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_y"},
		{"y_"},
		{"_ye"},
		{"ye_"},
		{"_yel"},
		{"yel_"},
		{"_yell"},
		{"yell_"},
		{"_yello"},
		{"yello_"},
		{"_yellow"},
		{"yellow_"},
		{"_yellow2"},
		{"yellow2_"},
		{"_yellow20"},
		{"yellow20_"},
		{"_yellow200"},
		{"yellow200_"},
		{"9713y"},
		{"y9713"},
		{"9713ye"},
		{"ye9713"},
		{"9713yel"},
		{"yel9713"},
		{"9713yell"},
		{"yell9713"},
		{"9713yello"},
		{"yello9713"},
		{"9713yellow"},
		{"yellow9713"},
		{"9713yellow2"},
		{"yellow29713"},
		{"9713yellow20"},
		{"yellow209713"},
		{"9713yellow200"},
		{"yellow2009713"},
		{"Zy"},
		{"yZ"},
		{"Zye"},
		{"yeZ"},
		{"Zyel"},
		{"yelZ"},
		{"Zyell"},
		{"yellZ"},
		{"Zyello"},
		{"yelloZ"},
		{"Zyellow"},
		{"yellowZ"},
		{"Zyellow2"},
		{"yellow2Z"},
		{"Zyellow20"},
		{"yellow20Z"},
		{"Zyellow200"},
		{"yellow200Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Yellow200()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestYellow300Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"yellow300", 9},
		{"yellow300 ", 9},
		{"yellow300\n", 9},
		{"yellow300.", 9},
		{"yellow300:", 9},
		{"yellow300,", 9},
		{"yellow300\"", 9},
		{"yellow300(", 9},
		{"yellow300)", 9},
		{"yellow300[", 9},
		{"yellow300]", 9},
		{"yellow300// comment", 9},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Yellow300()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestYellow300Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_y"},
		{"y_"},
		{"_ye"},
		{"ye_"},
		{"_yel"},
		{"yel_"},
		{"_yell"},
		{"yell_"},
		{"_yello"},
		{"yello_"},
		{"_yellow"},
		{"yellow_"},
		{"_yellow3"},
		{"yellow3_"},
		{"_yellow30"},
		{"yellow30_"},
		{"_yellow300"},
		{"yellow300_"},
		{"9713y"},
		{"y9713"},
		{"9713ye"},
		{"ye9713"},
		{"9713yel"},
		{"yel9713"},
		{"9713yell"},
		{"yell9713"},
		{"9713yello"},
		{"yello9713"},
		{"9713yellow"},
		{"yellow9713"},
		{"9713yellow3"},
		{"yellow39713"},
		{"9713yellow30"},
		{"yellow309713"},
		{"9713yellow300"},
		{"yellow3009713"},
		{"Zy"},
		{"yZ"},
		{"Zye"},
		{"yeZ"},
		{"Zyel"},
		{"yelZ"},
		{"Zyell"},
		{"yellZ"},
		{"Zyello"},
		{"yelloZ"},
		{"Zyellow"},
		{"yellowZ"},
		{"Zyellow3"},
		{"yellow3Z"},
		{"Zyellow30"},
		{"yellow30Z"},
		{"Zyellow300"},
		{"yellow300Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Yellow300()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestYellow400Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"yellow400", 9},
		{"yellow400 ", 9},
		{"yellow400\n", 9},
		{"yellow400.", 9},
		{"yellow400:", 9},
		{"yellow400,", 9},
		{"yellow400\"", 9},
		{"yellow400(", 9},
		{"yellow400)", 9},
		{"yellow400[", 9},
		{"yellow400]", 9},
		{"yellow400// comment", 9},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Yellow400()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestYellow400Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_y"},
		{"y_"},
		{"_ye"},
		{"ye_"},
		{"_yel"},
		{"yel_"},
		{"_yell"},
		{"yell_"},
		{"_yello"},
		{"yello_"},
		{"_yellow"},
		{"yellow_"},
		{"_yellow4"},
		{"yellow4_"},
		{"_yellow40"},
		{"yellow40_"},
		{"_yellow400"},
		{"yellow400_"},
		{"9713y"},
		{"y9713"},
		{"9713ye"},
		{"ye9713"},
		{"9713yel"},
		{"yel9713"},
		{"9713yell"},
		{"yell9713"},
		{"9713yello"},
		{"yello9713"},
		{"9713yellow"},
		{"yellow9713"},
		{"9713yellow4"},
		{"yellow49713"},
		{"9713yellow40"},
		{"yellow409713"},
		{"9713yellow400"},
		{"yellow4009713"},
		{"Zy"},
		{"yZ"},
		{"Zye"},
		{"yeZ"},
		{"Zyel"},
		{"yelZ"},
		{"Zyell"},
		{"yellZ"},
		{"Zyello"},
		{"yelloZ"},
		{"Zyellow"},
		{"yellowZ"},
		{"Zyellow4"},
		{"yellow4Z"},
		{"Zyellow40"},
		{"yellow40Z"},
		{"Zyellow400"},
		{"yellow400Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Yellow400()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestYellow500Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"yellow500", 9},
		{"yellow500 ", 9},
		{"yellow500\n", 9},
		{"yellow500.", 9},
		{"yellow500:", 9},
		{"yellow500,", 9},
		{"yellow500\"", 9},
		{"yellow500(", 9},
		{"yellow500)", 9},
		{"yellow500[", 9},
		{"yellow500]", 9},
		{"yellow500// comment", 9},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Yellow500()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestYellow500Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_y"},
		{"y_"},
		{"_ye"},
		{"ye_"},
		{"_yel"},
		{"yel_"},
		{"_yell"},
		{"yell_"},
		{"_yello"},
		{"yello_"},
		{"_yellow"},
		{"yellow_"},
		{"_yellow5"},
		{"yellow5_"},
		{"_yellow50"},
		{"yellow50_"},
		{"_yellow500"},
		{"yellow500_"},
		{"9713y"},
		{"y9713"},
		{"9713ye"},
		{"ye9713"},
		{"9713yel"},
		{"yel9713"},
		{"9713yell"},
		{"yell9713"},
		{"9713yello"},
		{"yello9713"},
		{"9713yellow"},
		{"yellow9713"},
		{"9713yellow5"},
		{"yellow59713"},
		{"9713yellow50"},
		{"yellow509713"},
		{"9713yellow500"},
		{"yellow5009713"},
		{"Zy"},
		{"yZ"},
		{"Zye"},
		{"yeZ"},
		{"Zyel"},
		{"yelZ"},
		{"Zyell"},
		{"yellZ"},
		{"Zyello"},
		{"yelloZ"},
		{"Zyellow"},
		{"yellowZ"},
		{"Zyellow5"},
		{"yellow5Z"},
		{"Zyellow50"},
		{"yellow50Z"},
		{"Zyellow500"},
		{"yellow500Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Yellow500()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestYellow600Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"yellow600", 9},
		{"yellow600 ", 9},
		{"yellow600\n", 9},
		{"yellow600.", 9},
		{"yellow600:", 9},
		{"yellow600,", 9},
		{"yellow600\"", 9},
		{"yellow600(", 9},
		{"yellow600)", 9},
		{"yellow600[", 9},
		{"yellow600]", 9},
		{"yellow600// comment", 9},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Yellow600()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestYellow600Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_y"},
		{"y_"},
		{"_ye"},
		{"ye_"},
		{"_yel"},
		{"yel_"},
		{"_yell"},
		{"yell_"},
		{"_yello"},
		{"yello_"},
		{"_yellow"},
		{"yellow_"},
		{"_yellow6"},
		{"yellow6_"},
		{"_yellow60"},
		{"yellow60_"},
		{"_yellow600"},
		{"yellow600_"},
		{"9713y"},
		{"y9713"},
		{"9713ye"},
		{"ye9713"},
		{"9713yel"},
		{"yel9713"},
		{"9713yell"},
		{"yell9713"},
		{"9713yello"},
		{"yello9713"},
		{"9713yellow"},
		{"yellow9713"},
		{"9713yellow6"},
		{"yellow69713"},
		{"9713yellow60"},
		{"yellow609713"},
		{"9713yellow600"},
		{"yellow6009713"},
		{"Zy"},
		{"yZ"},
		{"Zye"},
		{"yeZ"},
		{"Zyel"},
		{"yelZ"},
		{"Zyell"},
		{"yellZ"},
		{"Zyello"},
		{"yelloZ"},
		{"Zyellow"},
		{"yellowZ"},
		{"Zyellow6"},
		{"yellow6Z"},
		{"Zyellow60"},
		{"yellow60Z"},
		{"Zyellow600"},
		{"yellow600Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Yellow600()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestYellow700Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"yellow700", 9},
		{"yellow700 ", 9},
		{"yellow700\n", 9},
		{"yellow700.", 9},
		{"yellow700:", 9},
		{"yellow700,", 9},
		{"yellow700\"", 9},
		{"yellow700(", 9},
		{"yellow700)", 9},
		{"yellow700[", 9},
		{"yellow700]", 9},
		{"yellow700// comment", 9},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Yellow700()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestYellow700Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_y"},
		{"y_"},
		{"_ye"},
		{"ye_"},
		{"_yel"},
		{"yel_"},
		{"_yell"},
		{"yell_"},
		{"_yello"},
		{"yello_"},
		{"_yellow"},
		{"yellow_"},
		{"_yellow7"},
		{"yellow7_"},
		{"_yellow70"},
		{"yellow70_"},
		{"_yellow700"},
		{"yellow700_"},
		{"9713y"},
		{"y9713"},
		{"9713ye"},
		{"ye9713"},
		{"9713yel"},
		{"yel9713"},
		{"9713yell"},
		{"yell9713"},
		{"9713yello"},
		{"yello9713"},
		{"9713yellow"},
		{"yellow9713"},
		{"9713yellow7"},
		{"yellow79713"},
		{"9713yellow70"},
		{"yellow709713"},
		{"9713yellow700"},
		{"yellow7009713"},
		{"Zy"},
		{"yZ"},
		{"Zye"},
		{"yeZ"},
		{"Zyel"},
		{"yelZ"},
		{"Zyell"},
		{"yellZ"},
		{"Zyello"},
		{"yelloZ"},
		{"Zyellow"},
		{"yellowZ"},
		{"Zyellow7"},
		{"yellow7Z"},
		{"Zyellow70"},
		{"yellow70Z"},
		{"Zyellow700"},
		{"yellow700Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Yellow700()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestYellow800Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"yellow800", 9},
		{"yellow800 ", 9},
		{"yellow800\n", 9},
		{"yellow800.", 9},
		{"yellow800:", 9},
		{"yellow800,", 9},
		{"yellow800\"", 9},
		{"yellow800(", 9},
		{"yellow800)", 9},
		{"yellow800[", 9},
		{"yellow800]", 9},
		{"yellow800// comment", 9},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Yellow800()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestYellow800Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_y"},
		{"y_"},
		{"_ye"},
		{"ye_"},
		{"_yel"},
		{"yel_"},
		{"_yell"},
		{"yell_"},
		{"_yello"},
		{"yello_"},
		{"_yellow"},
		{"yellow_"},
		{"_yellow8"},
		{"yellow8_"},
		{"_yellow80"},
		{"yellow80_"},
		{"_yellow800"},
		{"yellow800_"},
		{"9713y"},
		{"y9713"},
		{"9713ye"},
		{"ye9713"},
		{"9713yel"},
		{"yel9713"},
		{"9713yell"},
		{"yell9713"},
		{"9713yello"},
		{"yello9713"},
		{"9713yellow"},
		{"yellow9713"},
		{"9713yellow8"},
		{"yellow89713"},
		{"9713yellow80"},
		{"yellow809713"},
		{"9713yellow800"},
		{"yellow8009713"},
		{"Zy"},
		{"yZ"},
		{"Zye"},
		{"yeZ"},
		{"Zyel"},
		{"yelZ"},
		{"Zyell"},
		{"yellZ"},
		{"Zyello"},
		{"yelloZ"},
		{"Zyellow"},
		{"yellowZ"},
		{"Zyellow8"},
		{"yellow8Z"},
		{"Zyellow80"},
		{"yellow80Z"},
		{"Zyellow800"},
		{"yellow800Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Yellow800()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestYellow900Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"yellow900", 9},
		{"yellow900 ", 9},
		{"yellow900\n", 9},
		{"yellow900.", 9},
		{"yellow900:", 9},
		{"yellow900,", 9},
		{"yellow900\"", 9},
		{"yellow900(", 9},
		{"yellow900)", 9},
		{"yellow900[", 9},
		{"yellow900]", 9},
		{"yellow900// comment", 9},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Yellow900()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestYellow900Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_y"},
		{"y_"},
		{"_ye"},
		{"ye_"},
		{"_yel"},
		{"yel_"},
		{"_yell"},
		{"yell_"},
		{"_yello"},
		{"yello_"},
		{"_yellow"},
		{"yellow_"},
		{"_yellow9"},
		{"yellow9_"},
		{"_yellow90"},
		{"yellow90_"},
		{"_yellow900"},
		{"yellow900_"},
		{"9713y"},
		{"y9713"},
		{"9713ye"},
		{"ye9713"},
		{"9713yel"},
		{"yel9713"},
		{"9713yell"},
		{"yell9713"},
		{"9713yello"},
		{"yello9713"},
		{"9713yellow"},
		{"yellow9713"},
		{"9713yellow9"},
		{"yellow99713"},
		{"9713yellow90"},
		{"yellow909713"},
		{"9713yellow900"},
		{"yellow9009713"},
		{"Zy"},
		{"yZ"},
		{"Zye"},
		{"yeZ"},
		{"Zyel"},
		{"yelZ"},
		{"Zyell"},
		{"yellZ"},
		{"Zyello"},
		{"yelloZ"},
		{"Zyellow"},
		{"yellowZ"},
		{"Zyellow9"},
		{"yellow9Z"},
		{"Zyellow90"},
		{"yellow90Z"},
		{"Zyellow900"},
		{"yellow900Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Yellow900()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestYellowA100Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"yellowA100", 10},
		{"yellowA100 ", 10},
		{"yellowA100\n", 10},
		{"yellowA100.", 10},
		{"yellowA100:", 10},
		{"yellowA100,", 10},
		{"yellowA100\"", 10},
		{"yellowA100(", 10},
		{"yellowA100)", 10},
		{"yellowA100[", 10},
		{"yellowA100]", 10},
		{"yellowA100// comment", 10},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.YellowA100()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestYellowA100Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_y"},
		{"y_"},
		{"_ye"},
		{"ye_"},
		{"_yel"},
		{"yel_"},
		{"_yell"},
		{"yell_"},
		{"_yello"},
		{"yello_"},
		{"_yellow"},
		{"yellow_"},
		{"_yellowA"},
		{"yellowA_"},
		{"_yellowA1"},
		{"yellowA1_"},
		{"_yellowA10"},
		{"yellowA10_"},
		{"_yellowA100"},
		{"yellowA100_"},
		{"9713y"},
		{"y9713"},
		{"9713ye"},
		{"ye9713"},
		{"9713yel"},
		{"yel9713"},
		{"9713yell"},
		{"yell9713"},
		{"9713yello"},
		{"yello9713"},
		{"9713yellow"},
		{"yellow9713"},
		{"9713yellowA"},
		{"yellowA9713"},
		{"9713yellowA1"},
		{"yellowA19713"},
		{"9713yellowA10"},
		{"yellowA109713"},
		{"9713yellowA100"},
		{"yellowA1009713"},
		{"Zy"},
		{"yZ"},
		{"Zye"},
		{"yeZ"},
		{"Zyel"},
		{"yelZ"},
		{"Zyell"},
		{"yellZ"},
		{"Zyello"},
		{"yelloZ"},
		{"Zyellow"},
		{"yellowZ"},
		{"ZyellowA"},
		{"yellowAZ"},
		{"ZyellowA1"},
		{"yellowA1Z"},
		{"ZyellowA10"},
		{"yellowA10Z"},
		{"ZyellowA100"},
		{"yellowA100Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.YellowA100()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestYellowA200Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"yellowA200", 10},
		{"yellowA200 ", 10},
		{"yellowA200\n", 10},
		{"yellowA200.", 10},
		{"yellowA200:", 10},
		{"yellowA200,", 10},
		{"yellowA200\"", 10},
		{"yellowA200(", 10},
		{"yellowA200)", 10},
		{"yellowA200[", 10},
		{"yellowA200]", 10},
		{"yellowA200// comment", 10},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.YellowA200()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestYellowA200Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_y"},
		{"y_"},
		{"_ye"},
		{"ye_"},
		{"_yel"},
		{"yel_"},
		{"_yell"},
		{"yell_"},
		{"_yello"},
		{"yello_"},
		{"_yellow"},
		{"yellow_"},
		{"_yellowA"},
		{"yellowA_"},
		{"_yellowA2"},
		{"yellowA2_"},
		{"_yellowA20"},
		{"yellowA20_"},
		{"_yellowA200"},
		{"yellowA200_"},
		{"9713y"},
		{"y9713"},
		{"9713ye"},
		{"ye9713"},
		{"9713yel"},
		{"yel9713"},
		{"9713yell"},
		{"yell9713"},
		{"9713yello"},
		{"yello9713"},
		{"9713yellow"},
		{"yellow9713"},
		{"9713yellowA"},
		{"yellowA9713"},
		{"9713yellowA2"},
		{"yellowA29713"},
		{"9713yellowA20"},
		{"yellowA209713"},
		{"9713yellowA200"},
		{"yellowA2009713"},
		{"Zy"},
		{"yZ"},
		{"Zye"},
		{"yeZ"},
		{"Zyel"},
		{"yelZ"},
		{"Zyell"},
		{"yellZ"},
		{"Zyello"},
		{"yelloZ"},
		{"Zyellow"},
		{"yellowZ"},
		{"ZyellowA"},
		{"yellowAZ"},
		{"ZyellowA2"},
		{"yellowA2Z"},
		{"ZyellowA20"},
		{"yellowA20Z"},
		{"ZyellowA200"},
		{"yellowA200Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.YellowA200()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestYellowA400Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"yellowA400", 10},
		{"yellowA400 ", 10},
		{"yellowA400\n", 10},
		{"yellowA400.", 10},
		{"yellowA400:", 10},
		{"yellowA400,", 10},
		{"yellowA400\"", 10},
		{"yellowA400(", 10},
		{"yellowA400)", 10},
		{"yellowA400[", 10},
		{"yellowA400]", 10},
		{"yellowA400// comment", 10},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.YellowA400()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestYellowA400Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_y"},
		{"y_"},
		{"_ye"},
		{"ye_"},
		{"_yel"},
		{"yel_"},
		{"_yell"},
		{"yell_"},
		{"_yello"},
		{"yello_"},
		{"_yellow"},
		{"yellow_"},
		{"_yellowA"},
		{"yellowA_"},
		{"_yellowA4"},
		{"yellowA4_"},
		{"_yellowA40"},
		{"yellowA40_"},
		{"_yellowA400"},
		{"yellowA400_"},
		{"9713y"},
		{"y9713"},
		{"9713ye"},
		{"ye9713"},
		{"9713yel"},
		{"yel9713"},
		{"9713yell"},
		{"yell9713"},
		{"9713yello"},
		{"yello9713"},
		{"9713yellow"},
		{"yellow9713"},
		{"9713yellowA"},
		{"yellowA9713"},
		{"9713yellowA4"},
		{"yellowA49713"},
		{"9713yellowA40"},
		{"yellowA409713"},
		{"9713yellowA400"},
		{"yellowA4009713"},
		{"Zy"},
		{"yZ"},
		{"Zye"},
		{"yeZ"},
		{"Zyel"},
		{"yelZ"},
		{"Zyell"},
		{"yellZ"},
		{"Zyello"},
		{"yelloZ"},
		{"Zyellow"},
		{"yellowZ"},
		{"ZyellowA"},
		{"yellowAZ"},
		{"ZyellowA4"},
		{"yellowA4Z"},
		{"ZyellowA40"},
		{"yellowA40Z"},
		{"ZyellowA400"},
		{"yellowA400Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.YellowA400()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestYellowA700Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"yellowA700", 10},
		{"yellowA700 ", 10},
		{"yellowA700\n", 10},
		{"yellowA700.", 10},
		{"yellowA700:", 10},
		{"yellowA700,", 10},
		{"yellowA700\"", 10},
		{"yellowA700(", 10},
		{"yellowA700)", 10},
		{"yellowA700[", 10},
		{"yellowA700]", 10},
		{"yellowA700// comment", 10},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.YellowA700()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestYellowA700Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_y"},
		{"y_"},
		{"_ye"},
		{"ye_"},
		{"_yel"},
		{"yel_"},
		{"_yell"},
		{"yell_"},
		{"_yello"},
		{"yello_"},
		{"_yellow"},
		{"yellow_"},
		{"_yellowA"},
		{"yellowA_"},
		{"_yellowA7"},
		{"yellowA7_"},
		{"_yellowA70"},
		{"yellowA70_"},
		{"_yellowA700"},
		{"yellowA700_"},
		{"9713y"},
		{"y9713"},
		{"9713ye"},
		{"ye9713"},
		{"9713yel"},
		{"yel9713"},
		{"9713yell"},
		{"yell9713"},
		{"9713yello"},
		{"yello9713"},
		{"9713yellow"},
		{"yellow9713"},
		{"9713yellowA"},
		{"yellowA9713"},
		{"9713yellowA7"},
		{"yellowA79713"},
		{"9713yellowA70"},
		{"yellowA709713"},
		{"9713yellowA700"},
		{"yellowA7009713"},
		{"Zy"},
		{"yZ"},
		{"Zye"},
		{"yeZ"},
		{"Zyel"},
		{"yelZ"},
		{"Zyell"},
		{"yellZ"},
		{"Zyello"},
		{"yelloZ"},
		{"Zyellow"},
		{"yellowZ"},
		{"ZyellowA"},
		{"yellowAZ"},
		{"ZyellowA7"},
		{"yellowA7Z"},
		{"ZyellowA70"},
		{"yellowA70Z"},
		{"ZyellowA700"},
		{"yellowA700Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.YellowA700()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestAmber50Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"amber50", 7},
		{"amber50 ", 7},
		{"amber50\n", 7},
		{"amber50.", 7},
		{"amber50:", 7},
		{"amber50,", 7},
		{"amber50\"", 7},
		{"amber50(", 7},
		{"amber50)", 7},
		{"amber50[", 7},
		{"amber50]", 7},
		{"amber50// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Amber50()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestAmber50Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_a"},
		{"a_"},
		{"_am"},
		{"am_"},
		{"_amb"},
		{"amb_"},
		{"_ambe"},
		{"ambe_"},
		{"_amber"},
		{"amber_"},
		{"_amber5"},
		{"amber5_"},
		{"_amber50"},
		{"amber50_"},
		{"9713a"},
		{"a9713"},
		{"9713am"},
		{"am9713"},
		{"9713amb"},
		{"amb9713"},
		{"9713ambe"},
		{"ambe9713"},
		{"9713amber"},
		{"amber9713"},
		{"9713amber5"},
		{"amber59713"},
		{"9713amber50"},
		{"amber509713"},
		{"Za"},
		{"aZ"},
		{"Zam"},
		{"amZ"},
		{"Zamb"},
		{"ambZ"},
		{"Zambe"},
		{"ambeZ"},
		{"Zamber"},
		{"amberZ"},
		{"Zamber5"},
		{"amber5Z"},
		{"Zamber50"},
		{"amber50Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Amber50()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestAmber100Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"amber100", 8},
		{"amber100 ", 8},
		{"amber100\n", 8},
		{"amber100.", 8},
		{"amber100:", 8},
		{"amber100,", 8},
		{"amber100\"", 8},
		{"amber100(", 8},
		{"amber100)", 8},
		{"amber100[", 8},
		{"amber100]", 8},
		{"amber100// comment", 8},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Amber100()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestAmber100Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_a"},
		{"a_"},
		{"_am"},
		{"am_"},
		{"_amb"},
		{"amb_"},
		{"_ambe"},
		{"ambe_"},
		{"_amber"},
		{"amber_"},
		{"_amber1"},
		{"amber1_"},
		{"_amber10"},
		{"amber10_"},
		{"_amber100"},
		{"amber100_"},
		{"9713a"},
		{"a9713"},
		{"9713am"},
		{"am9713"},
		{"9713amb"},
		{"amb9713"},
		{"9713ambe"},
		{"ambe9713"},
		{"9713amber"},
		{"amber9713"},
		{"9713amber1"},
		{"amber19713"},
		{"9713amber10"},
		{"amber109713"},
		{"9713amber100"},
		{"amber1009713"},
		{"Za"},
		{"aZ"},
		{"Zam"},
		{"amZ"},
		{"Zamb"},
		{"ambZ"},
		{"Zambe"},
		{"ambeZ"},
		{"Zamber"},
		{"amberZ"},
		{"Zamber1"},
		{"amber1Z"},
		{"Zamber10"},
		{"amber10Z"},
		{"Zamber100"},
		{"amber100Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Amber100()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestAmber200Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"amber200", 8},
		{"amber200 ", 8},
		{"amber200\n", 8},
		{"amber200.", 8},
		{"amber200:", 8},
		{"amber200,", 8},
		{"amber200\"", 8},
		{"amber200(", 8},
		{"amber200)", 8},
		{"amber200[", 8},
		{"amber200]", 8},
		{"amber200// comment", 8},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Amber200()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestAmber200Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_a"},
		{"a_"},
		{"_am"},
		{"am_"},
		{"_amb"},
		{"amb_"},
		{"_ambe"},
		{"ambe_"},
		{"_amber"},
		{"amber_"},
		{"_amber2"},
		{"amber2_"},
		{"_amber20"},
		{"amber20_"},
		{"_amber200"},
		{"amber200_"},
		{"9713a"},
		{"a9713"},
		{"9713am"},
		{"am9713"},
		{"9713amb"},
		{"amb9713"},
		{"9713ambe"},
		{"ambe9713"},
		{"9713amber"},
		{"amber9713"},
		{"9713amber2"},
		{"amber29713"},
		{"9713amber20"},
		{"amber209713"},
		{"9713amber200"},
		{"amber2009713"},
		{"Za"},
		{"aZ"},
		{"Zam"},
		{"amZ"},
		{"Zamb"},
		{"ambZ"},
		{"Zambe"},
		{"ambeZ"},
		{"Zamber"},
		{"amberZ"},
		{"Zamber2"},
		{"amber2Z"},
		{"Zamber20"},
		{"amber20Z"},
		{"Zamber200"},
		{"amber200Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Amber200()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestAmber300Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"amber300", 8},
		{"amber300 ", 8},
		{"amber300\n", 8},
		{"amber300.", 8},
		{"amber300:", 8},
		{"amber300,", 8},
		{"amber300\"", 8},
		{"amber300(", 8},
		{"amber300)", 8},
		{"amber300[", 8},
		{"amber300]", 8},
		{"amber300// comment", 8},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Amber300()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestAmber300Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_a"},
		{"a_"},
		{"_am"},
		{"am_"},
		{"_amb"},
		{"amb_"},
		{"_ambe"},
		{"ambe_"},
		{"_amber"},
		{"amber_"},
		{"_amber3"},
		{"amber3_"},
		{"_amber30"},
		{"amber30_"},
		{"_amber300"},
		{"amber300_"},
		{"9713a"},
		{"a9713"},
		{"9713am"},
		{"am9713"},
		{"9713amb"},
		{"amb9713"},
		{"9713ambe"},
		{"ambe9713"},
		{"9713amber"},
		{"amber9713"},
		{"9713amber3"},
		{"amber39713"},
		{"9713amber30"},
		{"amber309713"},
		{"9713amber300"},
		{"amber3009713"},
		{"Za"},
		{"aZ"},
		{"Zam"},
		{"amZ"},
		{"Zamb"},
		{"ambZ"},
		{"Zambe"},
		{"ambeZ"},
		{"Zamber"},
		{"amberZ"},
		{"Zamber3"},
		{"amber3Z"},
		{"Zamber30"},
		{"amber30Z"},
		{"Zamber300"},
		{"amber300Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Amber300()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestAmber400Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"amber400", 8},
		{"amber400 ", 8},
		{"amber400\n", 8},
		{"amber400.", 8},
		{"amber400:", 8},
		{"amber400,", 8},
		{"amber400\"", 8},
		{"amber400(", 8},
		{"amber400)", 8},
		{"amber400[", 8},
		{"amber400]", 8},
		{"amber400// comment", 8},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Amber400()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestAmber400Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_a"},
		{"a_"},
		{"_am"},
		{"am_"},
		{"_amb"},
		{"amb_"},
		{"_ambe"},
		{"ambe_"},
		{"_amber"},
		{"amber_"},
		{"_amber4"},
		{"amber4_"},
		{"_amber40"},
		{"amber40_"},
		{"_amber400"},
		{"amber400_"},
		{"9713a"},
		{"a9713"},
		{"9713am"},
		{"am9713"},
		{"9713amb"},
		{"amb9713"},
		{"9713ambe"},
		{"ambe9713"},
		{"9713amber"},
		{"amber9713"},
		{"9713amber4"},
		{"amber49713"},
		{"9713amber40"},
		{"amber409713"},
		{"9713amber400"},
		{"amber4009713"},
		{"Za"},
		{"aZ"},
		{"Zam"},
		{"amZ"},
		{"Zamb"},
		{"ambZ"},
		{"Zambe"},
		{"ambeZ"},
		{"Zamber"},
		{"amberZ"},
		{"Zamber4"},
		{"amber4Z"},
		{"Zamber40"},
		{"amber40Z"},
		{"Zamber400"},
		{"amber400Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Amber400()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestAmber500Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"amber500", 8},
		{"amber500 ", 8},
		{"amber500\n", 8},
		{"amber500.", 8},
		{"amber500:", 8},
		{"amber500,", 8},
		{"amber500\"", 8},
		{"amber500(", 8},
		{"amber500)", 8},
		{"amber500[", 8},
		{"amber500]", 8},
		{"amber500// comment", 8},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Amber500()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestAmber500Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_a"},
		{"a_"},
		{"_am"},
		{"am_"},
		{"_amb"},
		{"amb_"},
		{"_ambe"},
		{"ambe_"},
		{"_amber"},
		{"amber_"},
		{"_amber5"},
		{"amber5_"},
		{"_amber50"},
		{"amber50_"},
		{"_amber500"},
		{"amber500_"},
		{"9713a"},
		{"a9713"},
		{"9713am"},
		{"am9713"},
		{"9713amb"},
		{"amb9713"},
		{"9713ambe"},
		{"ambe9713"},
		{"9713amber"},
		{"amber9713"},
		{"9713amber5"},
		{"amber59713"},
		{"9713amber50"},
		{"amber509713"},
		{"9713amber500"},
		{"amber5009713"},
		{"Za"},
		{"aZ"},
		{"Zam"},
		{"amZ"},
		{"Zamb"},
		{"ambZ"},
		{"Zambe"},
		{"ambeZ"},
		{"Zamber"},
		{"amberZ"},
		{"Zamber5"},
		{"amber5Z"},
		{"Zamber50"},
		{"amber50Z"},
		{"Zamber500"},
		{"amber500Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Amber500()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestAmber600Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"amber600", 8},
		{"amber600 ", 8},
		{"amber600\n", 8},
		{"amber600.", 8},
		{"amber600:", 8},
		{"amber600,", 8},
		{"amber600\"", 8},
		{"amber600(", 8},
		{"amber600)", 8},
		{"amber600[", 8},
		{"amber600]", 8},
		{"amber600// comment", 8},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Amber600()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestAmber600Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_a"},
		{"a_"},
		{"_am"},
		{"am_"},
		{"_amb"},
		{"amb_"},
		{"_ambe"},
		{"ambe_"},
		{"_amber"},
		{"amber_"},
		{"_amber6"},
		{"amber6_"},
		{"_amber60"},
		{"amber60_"},
		{"_amber600"},
		{"amber600_"},
		{"9713a"},
		{"a9713"},
		{"9713am"},
		{"am9713"},
		{"9713amb"},
		{"amb9713"},
		{"9713ambe"},
		{"ambe9713"},
		{"9713amber"},
		{"amber9713"},
		{"9713amber6"},
		{"amber69713"},
		{"9713amber60"},
		{"amber609713"},
		{"9713amber600"},
		{"amber6009713"},
		{"Za"},
		{"aZ"},
		{"Zam"},
		{"amZ"},
		{"Zamb"},
		{"ambZ"},
		{"Zambe"},
		{"ambeZ"},
		{"Zamber"},
		{"amberZ"},
		{"Zamber6"},
		{"amber6Z"},
		{"Zamber60"},
		{"amber60Z"},
		{"Zamber600"},
		{"amber600Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Amber600()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestAmber700Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"amber700", 8},
		{"amber700 ", 8},
		{"amber700\n", 8},
		{"amber700.", 8},
		{"amber700:", 8},
		{"amber700,", 8},
		{"amber700\"", 8},
		{"amber700(", 8},
		{"amber700)", 8},
		{"amber700[", 8},
		{"amber700]", 8},
		{"amber700// comment", 8},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Amber700()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestAmber700Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_a"},
		{"a_"},
		{"_am"},
		{"am_"},
		{"_amb"},
		{"amb_"},
		{"_ambe"},
		{"ambe_"},
		{"_amber"},
		{"amber_"},
		{"_amber7"},
		{"amber7_"},
		{"_amber70"},
		{"amber70_"},
		{"_amber700"},
		{"amber700_"},
		{"9713a"},
		{"a9713"},
		{"9713am"},
		{"am9713"},
		{"9713amb"},
		{"amb9713"},
		{"9713ambe"},
		{"ambe9713"},
		{"9713amber"},
		{"amber9713"},
		{"9713amber7"},
		{"amber79713"},
		{"9713amber70"},
		{"amber709713"},
		{"9713amber700"},
		{"amber7009713"},
		{"Za"},
		{"aZ"},
		{"Zam"},
		{"amZ"},
		{"Zamb"},
		{"ambZ"},
		{"Zambe"},
		{"ambeZ"},
		{"Zamber"},
		{"amberZ"},
		{"Zamber7"},
		{"amber7Z"},
		{"Zamber70"},
		{"amber70Z"},
		{"Zamber700"},
		{"amber700Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Amber700()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestAmber800Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"amber800", 8},
		{"amber800 ", 8},
		{"amber800\n", 8},
		{"amber800.", 8},
		{"amber800:", 8},
		{"amber800,", 8},
		{"amber800\"", 8},
		{"amber800(", 8},
		{"amber800)", 8},
		{"amber800[", 8},
		{"amber800]", 8},
		{"amber800// comment", 8},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Amber800()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestAmber800Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_a"},
		{"a_"},
		{"_am"},
		{"am_"},
		{"_amb"},
		{"amb_"},
		{"_ambe"},
		{"ambe_"},
		{"_amber"},
		{"amber_"},
		{"_amber8"},
		{"amber8_"},
		{"_amber80"},
		{"amber80_"},
		{"_amber800"},
		{"amber800_"},
		{"9713a"},
		{"a9713"},
		{"9713am"},
		{"am9713"},
		{"9713amb"},
		{"amb9713"},
		{"9713ambe"},
		{"ambe9713"},
		{"9713amber"},
		{"amber9713"},
		{"9713amber8"},
		{"amber89713"},
		{"9713amber80"},
		{"amber809713"},
		{"9713amber800"},
		{"amber8009713"},
		{"Za"},
		{"aZ"},
		{"Zam"},
		{"amZ"},
		{"Zamb"},
		{"ambZ"},
		{"Zambe"},
		{"ambeZ"},
		{"Zamber"},
		{"amberZ"},
		{"Zamber8"},
		{"amber8Z"},
		{"Zamber80"},
		{"amber80Z"},
		{"Zamber800"},
		{"amber800Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Amber800()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestAmber900Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"amber900", 8},
		{"amber900 ", 8},
		{"amber900\n", 8},
		{"amber900.", 8},
		{"amber900:", 8},
		{"amber900,", 8},
		{"amber900\"", 8},
		{"amber900(", 8},
		{"amber900)", 8},
		{"amber900[", 8},
		{"amber900]", 8},
		{"amber900// comment", 8},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Amber900()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestAmber900Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_a"},
		{"a_"},
		{"_am"},
		{"am_"},
		{"_amb"},
		{"amb_"},
		{"_ambe"},
		{"ambe_"},
		{"_amber"},
		{"amber_"},
		{"_amber9"},
		{"amber9_"},
		{"_amber90"},
		{"amber90_"},
		{"_amber900"},
		{"amber900_"},
		{"9713a"},
		{"a9713"},
		{"9713am"},
		{"am9713"},
		{"9713amb"},
		{"amb9713"},
		{"9713ambe"},
		{"ambe9713"},
		{"9713amber"},
		{"amber9713"},
		{"9713amber9"},
		{"amber99713"},
		{"9713amber90"},
		{"amber909713"},
		{"9713amber900"},
		{"amber9009713"},
		{"Za"},
		{"aZ"},
		{"Zam"},
		{"amZ"},
		{"Zamb"},
		{"ambZ"},
		{"Zambe"},
		{"ambeZ"},
		{"Zamber"},
		{"amberZ"},
		{"Zamber9"},
		{"amber9Z"},
		{"Zamber90"},
		{"amber90Z"},
		{"Zamber900"},
		{"amber900Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Amber900()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestAmberA100Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"amberA100", 9},
		{"amberA100 ", 9},
		{"amberA100\n", 9},
		{"amberA100.", 9},
		{"amberA100:", 9},
		{"amberA100,", 9},
		{"amberA100\"", 9},
		{"amberA100(", 9},
		{"amberA100)", 9},
		{"amberA100[", 9},
		{"amberA100]", 9},
		{"amberA100// comment", 9},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.AmberA100()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestAmberA100Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_a"},
		{"a_"},
		{"_am"},
		{"am_"},
		{"_amb"},
		{"amb_"},
		{"_ambe"},
		{"ambe_"},
		{"_amber"},
		{"amber_"},
		{"_amberA"},
		{"amberA_"},
		{"_amberA1"},
		{"amberA1_"},
		{"_amberA10"},
		{"amberA10_"},
		{"_amberA100"},
		{"amberA100_"},
		{"9713a"},
		{"a9713"},
		{"9713am"},
		{"am9713"},
		{"9713amb"},
		{"amb9713"},
		{"9713ambe"},
		{"ambe9713"},
		{"9713amber"},
		{"amber9713"},
		{"9713amberA"},
		{"amberA9713"},
		{"9713amberA1"},
		{"amberA19713"},
		{"9713amberA10"},
		{"amberA109713"},
		{"9713amberA100"},
		{"amberA1009713"},
		{"Za"},
		{"aZ"},
		{"Zam"},
		{"amZ"},
		{"Zamb"},
		{"ambZ"},
		{"Zambe"},
		{"ambeZ"},
		{"Zamber"},
		{"amberZ"},
		{"ZamberA"},
		{"amberAZ"},
		{"ZamberA1"},
		{"amberA1Z"},
		{"ZamberA10"},
		{"amberA10Z"},
		{"ZamberA100"},
		{"amberA100Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.AmberA100()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestAmberA200Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"amberA200", 9},
		{"amberA200 ", 9},
		{"amberA200\n", 9},
		{"amberA200.", 9},
		{"amberA200:", 9},
		{"amberA200,", 9},
		{"amberA200\"", 9},
		{"amberA200(", 9},
		{"amberA200)", 9},
		{"amberA200[", 9},
		{"amberA200]", 9},
		{"amberA200// comment", 9},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.AmberA200()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestAmberA200Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_a"},
		{"a_"},
		{"_am"},
		{"am_"},
		{"_amb"},
		{"amb_"},
		{"_ambe"},
		{"ambe_"},
		{"_amber"},
		{"amber_"},
		{"_amberA"},
		{"amberA_"},
		{"_amberA2"},
		{"amberA2_"},
		{"_amberA20"},
		{"amberA20_"},
		{"_amberA200"},
		{"amberA200_"},
		{"9713a"},
		{"a9713"},
		{"9713am"},
		{"am9713"},
		{"9713amb"},
		{"amb9713"},
		{"9713ambe"},
		{"ambe9713"},
		{"9713amber"},
		{"amber9713"},
		{"9713amberA"},
		{"amberA9713"},
		{"9713amberA2"},
		{"amberA29713"},
		{"9713amberA20"},
		{"amberA209713"},
		{"9713amberA200"},
		{"amberA2009713"},
		{"Za"},
		{"aZ"},
		{"Zam"},
		{"amZ"},
		{"Zamb"},
		{"ambZ"},
		{"Zambe"},
		{"ambeZ"},
		{"Zamber"},
		{"amberZ"},
		{"ZamberA"},
		{"amberAZ"},
		{"ZamberA2"},
		{"amberA2Z"},
		{"ZamberA20"},
		{"amberA20Z"},
		{"ZamberA200"},
		{"amberA200Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.AmberA200()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestAmberA400Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"amberA400", 9},
		{"amberA400 ", 9},
		{"amberA400\n", 9},
		{"amberA400.", 9},
		{"amberA400:", 9},
		{"amberA400,", 9},
		{"amberA400\"", 9},
		{"amberA400(", 9},
		{"amberA400)", 9},
		{"amberA400[", 9},
		{"amberA400]", 9},
		{"amberA400// comment", 9},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.AmberA400()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestAmberA400Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_a"},
		{"a_"},
		{"_am"},
		{"am_"},
		{"_amb"},
		{"amb_"},
		{"_ambe"},
		{"ambe_"},
		{"_amber"},
		{"amber_"},
		{"_amberA"},
		{"amberA_"},
		{"_amberA4"},
		{"amberA4_"},
		{"_amberA40"},
		{"amberA40_"},
		{"_amberA400"},
		{"amberA400_"},
		{"9713a"},
		{"a9713"},
		{"9713am"},
		{"am9713"},
		{"9713amb"},
		{"amb9713"},
		{"9713ambe"},
		{"ambe9713"},
		{"9713amber"},
		{"amber9713"},
		{"9713amberA"},
		{"amberA9713"},
		{"9713amberA4"},
		{"amberA49713"},
		{"9713amberA40"},
		{"amberA409713"},
		{"9713amberA400"},
		{"amberA4009713"},
		{"Za"},
		{"aZ"},
		{"Zam"},
		{"amZ"},
		{"Zamb"},
		{"ambZ"},
		{"Zambe"},
		{"ambeZ"},
		{"Zamber"},
		{"amberZ"},
		{"ZamberA"},
		{"amberAZ"},
		{"ZamberA4"},
		{"amberA4Z"},
		{"ZamberA40"},
		{"amberA40Z"},
		{"ZamberA400"},
		{"amberA400Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.AmberA400()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestAmberA700Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"amberA700", 9},
		{"amberA700 ", 9},
		{"amberA700\n", 9},
		{"amberA700.", 9},
		{"amberA700:", 9},
		{"amberA700,", 9},
		{"amberA700\"", 9},
		{"amberA700(", 9},
		{"amberA700)", 9},
		{"amberA700[", 9},
		{"amberA700]", 9},
		{"amberA700// comment", 9},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.AmberA700()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestAmberA700Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_a"},
		{"a_"},
		{"_am"},
		{"am_"},
		{"_amb"},
		{"amb_"},
		{"_ambe"},
		{"ambe_"},
		{"_amber"},
		{"amber_"},
		{"_amberA"},
		{"amberA_"},
		{"_amberA7"},
		{"amberA7_"},
		{"_amberA70"},
		{"amberA70_"},
		{"_amberA700"},
		{"amberA700_"},
		{"9713a"},
		{"a9713"},
		{"9713am"},
		{"am9713"},
		{"9713amb"},
		{"amb9713"},
		{"9713ambe"},
		{"ambe9713"},
		{"9713amber"},
		{"amber9713"},
		{"9713amberA"},
		{"amberA9713"},
		{"9713amberA7"},
		{"amberA79713"},
		{"9713amberA70"},
		{"amberA709713"},
		{"9713amberA700"},
		{"amberA7009713"},
		{"Za"},
		{"aZ"},
		{"Zam"},
		{"amZ"},
		{"Zamb"},
		{"ambZ"},
		{"Zambe"},
		{"ambeZ"},
		{"Zamber"},
		{"amberZ"},
		{"ZamberA"},
		{"amberAZ"},
		{"ZamberA7"},
		{"amberA7Z"},
		{"ZamberA70"},
		{"amberA70Z"},
		{"ZamberA700"},
		{"amberA700Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.AmberA700()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestOrange50Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"orange50", 8},
		{"orange50 ", 8},
		{"orange50\n", 8},
		{"orange50.", 8},
		{"orange50:", 8},
		{"orange50,", 8},
		{"orange50\"", 8},
		{"orange50(", 8},
		{"orange50)", 8},
		{"orange50[", 8},
		{"orange50]", 8},
		{"orange50// comment", 8},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Orange50()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestOrange50Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_o"},
		{"o_"},
		{"_or"},
		{"or_"},
		{"_ora"},
		{"ora_"},
		{"_oran"},
		{"oran_"},
		{"_orang"},
		{"orang_"},
		{"_orange"},
		{"orange_"},
		{"_orange5"},
		{"orange5_"},
		{"_orange50"},
		{"orange50_"},
		{"9713o"},
		{"o9713"},
		{"9713or"},
		{"or9713"},
		{"9713ora"},
		{"ora9713"},
		{"9713oran"},
		{"oran9713"},
		{"9713orang"},
		{"orang9713"},
		{"9713orange"},
		{"orange9713"},
		{"9713orange5"},
		{"orange59713"},
		{"9713orange50"},
		{"orange509713"},
		{"Zo"},
		{"oZ"},
		{"Zor"},
		{"orZ"},
		{"Zora"},
		{"oraZ"},
		{"Zoran"},
		{"oranZ"},
		{"Zorang"},
		{"orangZ"},
		{"Zorange"},
		{"orangeZ"},
		{"Zorange5"},
		{"orange5Z"},
		{"Zorange50"},
		{"orange50Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Orange50()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestOrange100Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"orange100", 9},
		{"orange100 ", 9},
		{"orange100\n", 9},
		{"orange100.", 9},
		{"orange100:", 9},
		{"orange100,", 9},
		{"orange100\"", 9},
		{"orange100(", 9},
		{"orange100)", 9},
		{"orange100[", 9},
		{"orange100]", 9},
		{"orange100// comment", 9},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Orange100()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestOrange100Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_o"},
		{"o_"},
		{"_or"},
		{"or_"},
		{"_ora"},
		{"ora_"},
		{"_oran"},
		{"oran_"},
		{"_orang"},
		{"orang_"},
		{"_orange"},
		{"orange_"},
		{"_orange1"},
		{"orange1_"},
		{"_orange10"},
		{"orange10_"},
		{"_orange100"},
		{"orange100_"},
		{"9713o"},
		{"o9713"},
		{"9713or"},
		{"or9713"},
		{"9713ora"},
		{"ora9713"},
		{"9713oran"},
		{"oran9713"},
		{"9713orang"},
		{"orang9713"},
		{"9713orange"},
		{"orange9713"},
		{"9713orange1"},
		{"orange19713"},
		{"9713orange10"},
		{"orange109713"},
		{"9713orange100"},
		{"orange1009713"},
		{"Zo"},
		{"oZ"},
		{"Zor"},
		{"orZ"},
		{"Zora"},
		{"oraZ"},
		{"Zoran"},
		{"oranZ"},
		{"Zorang"},
		{"orangZ"},
		{"Zorange"},
		{"orangeZ"},
		{"Zorange1"},
		{"orange1Z"},
		{"Zorange10"},
		{"orange10Z"},
		{"Zorange100"},
		{"orange100Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Orange100()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestOrange200Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"orange200", 9},
		{"orange200 ", 9},
		{"orange200\n", 9},
		{"orange200.", 9},
		{"orange200:", 9},
		{"orange200,", 9},
		{"orange200\"", 9},
		{"orange200(", 9},
		{"orange200)", 9},
		{"orange200[", 9},
		{"orange200]", 9},
		{"orange200// comment", 9},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Orange200()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestOrange200Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_o"},
		{"o_"},
		{"_or"},
		{"or_"},
		{"_ora"},
		{"ora_"},
		{"_oran"},
		{"oran_"},
		{"_orang"},
		{"orang_"},
		{"_orange"},
		{"orange_"},
		{"_orange2"},
		{"orange2_"},
		{"_orange20"},
		{"orange20_"},
		{"_orange200"},
		{"orange200_"},
		{"9713o"},
		{"o9713"},
		{"9713or"},
		{"or9713"},
		{"9713ora"},
		{"ora9713"},
		{"9713oran"},
		{"oran9713"},
		{"9713orang"},
		{"orang9713"},
		{"9713orange"},
		{"orange9713"},
		{"9713orange2"},
		{"orange29713"},
		{"9713orange20"},
		{"orange209713"},
		{"9713orange200"},
		{"orange2009713"},
		{"Zo"},
		{"oZ"},
		{"Zor"},
		{"orZ"},
		{"Zora"},
		{"oraZ"},
		{"Zoran"},
		{"oranZ"},
		{"Zorang"},
		{"orangZ"},
		{"Zorange"},
		{"orangeZ"},
		{"Zorange2"},
		{"orange2Z"},
		{"Zorange20"},
		{"orange20Z"},
		{"Zorange200"},
		{"orange200Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Orange200()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestOrange300Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"orange300", 9},
		{"orange300 ", 9},
		{"orange300\n", 9},
		{"orange300.", 9},
		{"orange300:", 9},
		{"orange300,", 9},
		{"orange300\"", 9},
		{"orange300(", 9},
		{"orange300)", 9},
		{"orange300[", 9},
		{"orange300]", 9},
		{"orange300// comment", 9},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Orange300()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestOrange300Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_o"},
		{"o_"},
		{"_or"},
		{"or_"},
		{"_ora"},
		{"ora_"},
		{"_oran"},
		{"oran_"},
		{"_orang"},
		{"orang_"},
		{"_orange"},
		{"orange_"},
		{"_orange3"},
		{"orange3_"},
		{"_orange30"},
		{"orange30_"},
		{"_orange300"},
		{"orange300_"},
		{"9713o"},
		{"o9713"},
		{"9713or"},
		{"or9713"},
		{"9713ora"},
		{"ora9713"},
		{"9713oran"},
		{"oran9713"},
		{"9713orang"},
		{"orang9713"},
		{"9713orange"},
		{"orange9713"},
		{"9713orange3"},
		{"orange39713"},
		{"9713orange30"},
		{"orange309713"},
		{"9713orange300"},
		{"orange3009713"},
		{"Zo"},
		{"oZ"},
		{"Zor"},
		{"orZ"},
		{"Zora"},
		{"oraZ"},
		{"Zoran"},
		{"oranZ"},
		{"Zorang"},
		{"orangZ"},
		{"Zorange"},
		{"orangeZ"},
		{"Zorange3"},
		{"orange3Z"},
		{"Zorange30"},
		{"orange30Z"},
		{"Zorange300"},
		{"orange300Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Orange300()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestOrange400Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"orange400", 9},
		{"orange400 ", 9},
		{"orange400\n", 9},
		{"orange400.", 9},
		{"orange400:", 9},
		{"orange400,", 9},
		{"orange400\"", 9},
		{"orange400(", 9},
		{"orange400)", 9},
		{"orange400[", 9},
		{"orange400]", 9},
		{"orange400// comment", 9},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Orange400()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestOrange400Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_o"},
		{"o_"},
		{"_or"},
		{"or_"},
		{"_ora"},
		{"ora_"},
		{"_oran"},
		{"oran_"},
		{"_orang"},
		{"orang_"},
		{"_orange"},
		{"orange_"},
		{"_orange4"},
		{"orange4_"},
		{"_orange40"},
		{"orange40_"},
		{"_orange400"},
		{"orange400_"},
		{"9713o"},
		{"o9713"},
		{"9713or"},
		{"or9713"},
		{"9713ora"},
		{"ora9713"},
		{"9713oran"},
		{"oran9713"},
		{"9713orang"},
		{"orang9713"},
		{"9713orange"},
		{"orange9713"},
		{"9713orange4"},
		{"orange49713"},
		{"9713orange40"},
		{"orange409713"},
		{"9713orange400"},
		{"orange4009713"},
		{"Zo"},
		{"oZ"},
		{"Zor"},
		{"orZ"},
		{"Zora"},
		{"oraZ"},
		{"Zoran"},
		{"oranZ"},
		{"Zorang"},
		{"orangZ"},
		{"Zorange"},
		{"orangeZ"},
		{"Zorange4"},
		{"orange4Z"},
		{"Zorange40"},
		{"orange40Z"},
		{"Zorange400"},
		{"orange400Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Orange400()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestOrange500Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"orange500", 9},
		{"orange500 ", 9},
		{"orange500\n", 9},
		{"orange500.", 9},
		{"orange500:", 9},
		{"orange500,", 9},
		{"orange500\"", 9},
		{"orange500(", 9},
		{"orange500)", 9},
		{"orange500[", 9},
		{"orange500]", 9},
		{"orange500// comment", 9},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Orange500()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestOrange500Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_o"},
		{"o_"},
		{"_or"},
		{"or_"},
		{"_ora"},
		{"ora_"},
		{"_oran"},
		{"oran_"},
		{"_orang"},
		{"orang_"},
		{"_orange"},
		{"orange_"},
		{"_orange5"},
		{"orange5_"},
		{"_orange50"},
		{"orange50_"},
		{"_orange500"},
		{"orange500_"},
		{"9713o"},
		{"o9713"},
		{"9713or"},
		{"or9713"},
		{"9713ora"},
		{"ora9713"},
		{"9713oran"},
		{"oran9713"},
		{"9713orang"},
		{"orang9713"},
		{"9713orange"},
		{"orange9713"},
		{"9713orange5"},
		{"orange59713"},
		{"9713orange50"},
		{"orange509713"},
		{"9713orange500"},
		{"orange5009713"},
		{"Zo"},
		{"oZ"},
		{"Zor"},
		{"orZ"},
		{"Zora"},
		{"oraZ"},
		{"Zoran"},
		{"oranZ"},
		{"Zorang"},
		{"orangZ"},
		{"Zorange"},
		{"orangeZ"},
		{"Zorange5"},
		{"orange5Z"},
		{"Zorange50"},
		{"orange50Z"},
		{"Zorange500"},
		{"orange500Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Orange500()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestOrange600Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"orange600", 9},
		{"orange600 ", 9},
		{"orange600\n", 9},
		{"orange600.", 9},
		{"orange600:", 9},
		{"orange600,", 9},
		{"orange600\"", 9},
		{"orange600(", 9},
		{"orange600)", 9},
		{"orange600[", 9},
		{"orange600]", 9},
		{"orange600// comment", 9},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Orange600()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestOrange600Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_o"},
		{"o_"},
		{"_or"},
		{"or_"},
		{"_ora"},
		{"ora_"},
		{"_oran"},
		{"oran_"},
		{"_orang"},
		{"orang_"},
		{"_orange"},
		{"orange_"},
		{"_orange6"},
		{"orange6_"},
		{"_orange60"},
		{"orange60_"},
		{"_orange600"},
		{"orange600_"},
		{"9713o"},
		{"o9713"},
		{"9713or"},
		{"or9713"},
		{"9713ora"},
		{"ora9713"},
		{"9713oran"},
		{"oran9713"},
		{"9713orang"},
		{"orang9713"},
		{"9713orange"},
		{"orange9713"},
		{"9713orange6"},
		{"orange69713"},
		{"9713orange60"},
		{"orange609713"},
		{"9713orange600"},
		{"orange6009713"},
		{"Zo"},
		{"oZ"},
		{"Zor"},
		{"orZ"},
		{"Zora"},
		{"oraZ"},
		{"Zoran"},
		{"oranZ"},
		{"Zorang"},
		{"orangZ"},
		{"Zorange"},
		{"orangeZ"},
		{"Zorange6"},
		{"orange6Z"},
		{"Zorange60"},
		{"orange60Z"},
		{"Zorange600"},
		{"orange600Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Orange600()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestOrange700Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"orange700", 9},
		{"orange700 ", 9},
		{"orange700\n", 9},
		{"orange700.", 9},
		{"orange700:", 9},
		{"orange700,", 9},
		{"orange700\"", 9},
		{"orange700(", 9},
		{"orange700)", 9},
		{"orange700[", 9},
		{"orange700]", 9},
		{"orange700// comment", 9},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Orange700()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestOrange700Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_o"},
		{"o_"},
		{"_or"},
		{"or_"},
		{"_ora"},
		{"ora_"},
		{"_oran"},
		{"oran_"},
		{"_orang"},
		{"orang_"},
		{"_orange"},
		{"orange_"},
		{"_orange7"},
		{"orange7_"},
		{"_orange70"},
		{"orange70_"},
		{"_orange700"},
		{"orange700_"},
		{"9713o"},
		{"o9713"},
		{"9713or"},
		{"or9713"},
		{"9713ora"},
		{"ora9713"},
		{"9713oran"},
		{"oran9713"},
		{"9713orang"},
		{"orang9713"},
		{"9713orange"},
		{"orange9713"},
		{"9713orange7"},
		{"orange79713"},
		{"9713orange70"},
		{"orange709713"},
		{"9713orange700"},
		{"orange7009713"},
		{"Zo"},
		{"oZ"},
		{"Zor"},
		{"orZ"},
		{"Zora"},
		{"oraZ"},
		{"Zoran"},
		{"oranZ"},
		{"Zorang"},
		{"orangZ"},
		{"Zorange"},
		{"orangeZ"},
		{"Zorange7"},
		{"orange7Z"},
		{"Zorange70"},
		{"orange70Z"},
		{"Zorange700"},
		{"orange700Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Orange700()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestOrange800Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"orange800", 9},
		{"orange800 ", 9},
		{"orange800\n", 9},
		{"orange800.", 9},
		{"orange800:", 9},
		{"orange800,", 9},
		{"orange800\"", 9},
		{"orange800(", 9},
		{"orange800)", 9},
		{"orange800[", 9},
		{"orange800]", 9},
		{"orange800// comment", 9},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Orange800()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestOrange800Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_o"},
		{"o_"},
		{"_or"},
		{"or_"},
		{"_ora"},
		{"ora_"},
		{"_oran"},
		{"oran_"},
		{"_orang"},
		{"orang_"},
		{"_orange"},
		{"orange_"},
		{"_orange8"},
		{"orange8_"},
		{"_orange80"},
		{"orange80_"},
		{"_orange800"},
		{"orange800_"},
		{"9713o"},
		{"o9713"},
		{"9713or"},
		{"or9713"},
		{"9713ora"},
		{"ora9713"},
		{"9713oran"},
		{"oran9713"},
		{"9713orang"},
		{"orang9713"},
		{"9713orange"},
		{"orange9713"},
		{"9713orange8"},
		{"orange89713"},
		{"9713orange80"},
		{"orange809713"},
		{"9713orange800"},
		{"orange8009713"},
		{"Zo"},
		{"oZ"},
		{"Zor"},
		{"orZ"},
		{"Zora"},
		{"oraZ"},
		{"Zoran"},
		{"oranZ"},
		{"Zorang"},
		{"orangZ"},
		{"Zorange"},
		{"orangeZ"},
		{"Zorange8"},
		{"orange8Z"},
		{"Zorange80"},
		{"orange80Z"},
		{"Zorange800"},
		{"orange800Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Orange800()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestOrange900Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"orange900", 9},
		{"orange900 ", 9},
		{"orange900\n", 9},
		{"orange900.", 9},
		{"orange900:", 9},
		{"orange900,", 9},
		{"orange900\"", 9},
		{"orange900(", 9},
		{"orange900)", 9},
		{"orange900[", 9},
		{"orange900]", 9},
		{"orange900// comment", 9},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Orange900()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestOrange900Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_o"},
		{"o_"},
		{"_or"},
		{"or_"},
		{"_ora"},
		{"ora_"},
		{"_oran"},
		{"oran_"},
		{"_orang"},
		{"orang_"},
		{"_orange"},
		{"orange_"},
		{"_orange9"},
		{"orange9_"},
		{"_orange90"},
		{"orange90_"},
		{"_orange900"},
		{"orange900_"},
		{"9713o"},
		{"o9713"},
		{"9713or"},
		{"or9713"},
		{"9713ora"},
		{"ora9713"},
		{"9713oran"},
		{"oran9713"},
		{"9713orang"},
		{"orang9713"},
		{"9713orange"},
		{"orange9713"},
		{"9713orange9"},
		{"orange99713"},
		{"9713orange90"},
		{"orange909713"},
		{"9713orange900"},
		{"orange9009713"},
		{"Zo"},
		{"oZ"},
		{"Zor"},
		{"orZ"},
		{"Zora"},
		{"oraZ"},
		{"Zoran"},
		{"oranZ"},
		{"Zorang"},
		{"orangZ"},
		{"Zorange"},
		{"orangeZ"},
		{"Zorange9"},
		{"orange9Z"},
		{"Zorange90"},
		{"orange90Z"},
		{"Zorange900"},
		{"orange900Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Orange900()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestOrangeA100Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"orangeA100", 10},
		{"orangeA100 ", 10},
		{"orangeA100\n", 10},
		{"orangeA100.", 10},
		{"orangeA100:", 10},
		{"orangeA100,", 10},
		{"orangeA100\"", 10},
		{"orangeA100(", 10},
		{"orangeA100)", 10},
		{"orangeA100[", 10},
		{"orangeA100]", 10},
		{"orangeA100// comment", 10},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.OrangeA100()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestOrangeA100Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_o"},
		{"o_"},
		{"_or"},
		{"or_"},
		{"_ora"},
		{"ora_"},
		{"_oran"},
		{"oran_"},
		{"_orang"},
		{"orang_"},
		{"_orange"},
		{"orange_"},
		{"_orangeA"},
		{"orangeA_"},
		{"_orangeA1"},
		{"orangeA1_"},
		{"_orangeA10"},
		{"orangeA10_"},
		{"_orangeA100"},
		{"orangeA100_"},
		{"9713o"},
		{"o9713"},
		{"9713or"},
		{"or9713"},
		{"9713ora"},
		{"ora9713"},
		{"9713oran"},
		{"oran9713"},
		{"9713orang"},
		{"orang9713"},
		{"9713orange"},
		{"orange9713"},
		{"9713orangeA"},
		{"orangeA9713"},
		{"9713orangeA1"},
		{"orangeA19713"},
		{"9713orangeA10"},
		{"orangeA109713"},
		{"9713orangeA100"},
		{"orangeA1009713"},
		{"Zo"},
		{"oZ"},
		{"Zor"},
		{"orZ"},
		{"Zora"},
		{"oraZ"},
		{"Zoran"},
		{"oranZ"},
		{"Zorang"},
		{"orangZ"},
		{"Zorange"},
		{"orangeZ"},
		{"ZorangeA"},
		{"orangeAZ"},
		{"ZorangeA1"},
		{"orangeA1Z"},
		{"ZorangeA10"},
		{"orangeA10Z"},
		{"ZorangeA100"},
		{"orangeA100Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.OrangeA100()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestOrangeA200Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"orangeA200", 10},
		{"orangeA200 ", 10},
		{"orangeA200\n", 10},
		{"orangeA200.", 10},
		{"orangeA200:", 10},
		{"orangeA200,", 10},
		{"orangeA200\"", 10},
		{"orangeA200(", 10},
		{"orangeA200)", 10},
		{"orangeA200[", 10},
		{"orangeA200]", 10},
		{"orangeA200// comment", 10},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.OrangeA200()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestOrangeA200Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_o"},
		{"o_"},
		{"_or"},
		{"or_"},
		{"_ora"},
		{"ora_"},
		{"_oran"},
		{"oran_"},
		{"_orang"},
		{"orang_"},
		{"_orange"},
		{"orange_"},
		{"_orangeA"},
		{"orangeA_"},
		{"_orangeA2"},
		{"orangeA2_"},
		{"_orangeA20"},
		{"orangeA20_"},
		{"_orangeA200"},
		{"orangeA200_"},
		{"9713o"},
		{"o9713"},
		{"9713or"},
		{"or9713"},
		{"9713ora"},
		{"ora9713"},
		{"9713oran"},
		{"oran9713"},
		{"9713orang"},
		{"orang9713"},
		{"9713orange"},
		{"orange9713"},
		{"9713orangeA"},
		{"orangeA9713"},
		{"9713orangeA2"},
		{"orangeA29713"},
		{"9713orangeA20"},
		{"orangeA209713"},
		{"9713orangeA200"},
		{"orangeA2009713"},
		{"Zo"},
		{"oZ"},
		{"Zor"},
		{"orZ"},
		{"Zora"},
		{"oraZ"},
		{"Zoran"},
		{"oranZ"},
		{"Zorang"},
		{"orangZ"},
		{"Zorange"},
		{"orangeZ"},
		{"ZorangeA"},
		{"orangeAZ"},
		{"ZorangeA2"},
		{"orangeA2Z"},
		{"ZorangeA20"},
		{"orangeA20Z"},
		{"ZorangeA200"},
		{"orangeA200Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.OrangeA200()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestOrangeA400Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"orangeA400", 10},
		{"orangeA400 ", 10},
		{"orangeA400\n", 10},
		{"orangeA400.", 10},
		{"orangeA400:", 10},
		{"orangeA400,", 10},
		{"orangeA400\"", 10},
		{"orangeA400(", 10},
		{"orangeA400)", 10},
		{"orangeA400[", 10},
		{"orangeA400]", 10},
		{"orangeA400// comment", 10},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.OrangeA400()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestOrangeA400Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_o"},
		{"o_"},
		{"_or"},
		{"or_"},
		{"_ora"},
		{"ora_"},
		{"_oran"},
		{"oran_"},
		{"_orang"},
		{"orang_"},
		{"_orange"},
		{"orange_"},
		{"_orangeA"},
		{"orangeA_"},
		{"_orangeA4"},
		{"orangeA4_"},
		{"_orangeA40"},
		{"orangeA40_"},
		{"_orangeA400"},
		{"orangeA400_"},
		{"9713o"},
		{"o9713"},
		{"9713or"},
		{"or9713"},
		{"9713ora"},
		{"ora9713"},
		{"9713oran"},
		{"oran9713"},
		{"9713orang"},
		{"orang9713"},
		{"9713orange"},
		{"orange9713"},
		{"9713orangeA"},
		{"orangeA9713"},
		{"9713orangeA4"},
		{"orangeA49713"},
		{"9713orangeA40"},
		{"orangeA409713"},
		{"9713orangeA400"},
		{"orangeA4009713"},
		{"Zo"},
		{"oZ"},
		{"Zor"},
		{"orZ"},
		{"Zora"},
		{"oraZ"},
		{"Zoran"},
		{"oranZ"},
		{"Zorang"},
		{"orangZ"},
		{"Zorange"},
		{"orangeZ"},
		{"ZorangeA"},
		{"orangeAZ"},
		{"ZorangeA4"},
		{"orangeA4Z"},
		{"ZorangeA40"},
		{"orangeA40Z"},
		{"ZorangeA400"},
		{"orangeA400Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.OrangeA400()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestOrangeA700Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"orangeA700", 10},
		{"orangeA700 ", 10},
		{"orangeA700\n", 10},
		{"orangeA700.", 10},
		{"orangeA700:", 10},
		{"orangeA700,", 10},
		{"orangeA700\"", 10},
		{"orangeA700(", 10},
		{"orangeA700)", 10},
		{"orangeA700[", 10},
		{"orangeA700]", 10},
		{"orangeA700// comment", 10},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.OrangeA700()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestOrangeA700Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_o"},
		{"o_"},
		{"_or"},
		{"or_"},
		{"_ora"},
		{"ora_"},
		{"_oran"},
		{"oran_"},
		{"_orang"},
		{"orang_"},
		{"_orange"},
		{"orange_"},
		{"_orangeA"},
		{"orangeA_"},
		{"_orangeA7"},
		{"orangeA7_"},
		{"_orangeA70"},
		{"orangeA70_"},
		{"_orangeA700"},
		{"orangeA700_"},
		{"9713o"},
		{"o9713"},
		{"9713or"},
		{"or9713"},
		{"9713ora"},
		{"ora9713"},
		{"9713oran"},
		{"oran9713"},
		{"9713orang"},
		{"orang9713"},
		{"9713orange"},
		{"orange9713"},
		{"9713orangeA"},
		{"orangeA9713"},
		{"9713orangeA7"},
		{"orangeA79713"},
		{"9713orangeA70"},
		{"orangeA709713"},
		{"9713orangeA700"},
		{"orangeA7009713"},
		{"Zo"},
		{"oZ"},
		{"Zor"},
		{"orZ"},
		{"Zora"},
		{"oraZ"},
		{"Zoran"},
		{"oranZ"},
		{"Zorang"},
		{"orangZ"},
		{"Zorange"},
		{"orangeZ"},
		{"ZorangeA"},
		{"orangeAZ"},
		{"ZorangeA7"},
		{"orangeA7Z"},
		{"ZorangeA70"},
		{"orangeA70Z"},
		{"ZorangeA700"},
		{"orangeA700Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.OrangeA700()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestDeepOrange50Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"deepOrange50", 12},
		{"deepOrange50 ", 12},
		{"deepOrange50\n", 12},
		{"deepOrange50.", 12},
		{"deepOrange50:", 12},
		{"deepOrange50,", 12},
		{"deepOrange50\"", 12},
		{"deepOrange50(", 12},
		{"deepOrange50)", 12},
		{"deepOrange50[", 12},
		{"deepOrange50]", 12},
		{"deepOrange50// comment", 12},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.DeepOrange50()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestDeepOrange50Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_d"},
		{"d_"},
		{"_de"},
		{"de_"},
		{"_dee"},
		{"dee_"},
		{"_deep"},
		{"deep_"},
		{"_deepO"},
		{"deepO_"},
		{"_deepOr"},
		{"deepOr_"},
		{"_deepOra"},
		{"deepOra_"},
		{"_deepOran"},
		{"deepOran_"},
		{"_deepOrang"},
		{"deepOrang_"},
		{"_deepOrange"},
		{"deepOrange_"},
		{"_deepOrange5"},
		{"deepOrange5_"},
		{"_deepOrange50"},
		{"deepOrange50_"},
		{"9713d"},
		{"d9713"},
		{"9713de"},
		{"de9713"},
		{"9713dee"},
		{"dee9713"},
		{"9713deep"},
		{"deep9713"},
		{"9713deepO"},
		{"deepO9713"},
		{"9713deepOr"},
		{"deepOr9713"},
		{"9713deepOra"},
		{"deepOra9713"},
		{"9713deepOran"},
		{"deepOran9713"},
		{"9713deepOrang"},
		{"deepOrang9713"},
		{"9713deepOrange"},
		{"deepOrange9713"},
		{"9713deepOrange5"},
		{"deepOrange59713"},
		{"9713deepOrange50"},
		{"deepOrange509713"},
		{"Zd"},
		{"dZ"},
		{"Zde"},
		{"deZ"},
		{"Zdee"},
		{"deeZ"},
		{"Zdeep"},
		{"deepZ"},
		{"ZdeepO"},
		{"deepOZ"},
		{"ZdeepOr"},
		{"deepOrZ"},
		{"ZdeepOra"},
		{"deepOraZ"},
		{"ZdeepOran"},
		{"deepOranZ"},
		{"ZdeepOrang"},
		{"deepOrangZ"},
		{"ZdeepOrange"},
		{"deepOrangeZ"},
		{"ZdeepOrange5"},
		{"deepOrange5Z"},
		{"ZdeepOrange50"},
		{"deepOrange50Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.DeepOrange50()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestDeepOrange100Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"deepOrange100", 13},
		{"deepOrange100 ", 13},
		{"deepOrange100\n", 13},
		{"deepOrange100.", 13},
		{"deepOrange100:", 13},
		{"deepOrange100,", 13},
		{"deepOrange100\"", 13},
		{"deepOrange100(", 13},
		{"deepOrange100)", 13},
		{"deepOrange100[", 13},
		{"deepOrange100]", 13},
		{"deepOrange100// comment", 13},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.DeepOrange100()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestDeepOrange100Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_d"},
		{"d_"},
		{"_de"},
		{"de_"},
		{"_dee"},
		{"dee_"},
		{"_deep"},
		{"deep_"},
		{"_deepO"},
		{"deepO_"},
		{"_deepOr"},
		{"deepOr_"},
		{"_deepOra"},
		{"deepOra_"},
		{"_deepOran"},
		{"deepOran_"},
		{"_deepOrang"},
		{"deepOrang_"},
		{"_deepOrange"},
		{"deepOrange_"},
		{"_deepOrange1"},
		{"deepOrange1_"},
		{"_deepOrange10"},
		{"deepOrange10_"},
		{"_deepOrange100"},
		{"deepOrange100_"},
		{"9713d"},
		{"d9713"},
		{"9713de"},
		{"de9713"},
		{"9713dee"},
		{"dee9713"},
		{"9713deep"},
		{"deep9713"},
		{"9713deepO"},
		{"deepO9713"},
		{"9713deepOr"},
		{"deepOr9713"},
		{"9713deepOra"},
		{"deepOra9713"},
		{"9713deepOran"},
		{"deepOran9713"},
		{"9713deepOrang"},
		{"deepOrang9713"},
		{"9713deepOrange"},
		{"deepOrange9713"},
		{"9713deepOrange1"},
		{"deepOrange19713"},
		{"9713deepOrange10"},
		{"deepOrange109713"},
		{"9713deepOrange100"},
		{"deepOrange1009713"},
		{"Zd"},
		{"dZ"},
		{"Zde"},
		{"deZ"},
		{"Zdee"},
		{"deeZ"},
		{"Zdeep"},
		{"deepZ"},
		{"ZdeepO"},
		{"deepOZ"},
		{"ZdeepOr"},
		{"deepOrZ"},
		{"ZdeepOra"},
		{"deepOraZ"},
		{"ZdeepOran"},
		{"deepOranZ"},
		{"ZdeepOrang"},
		{"deepOrangZ"},
		{"ZdeepOrange"},
		{"deepOrangeZ"},
		{"ZdeepOrange1"},
		{"deepOrange1Z"},
		{"ZdeepOrange10"},
		{"deepOrange10Z"},
		{"ZdeepOrange100"},
		{"deepOrange100Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.DeepOrange100()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestDeepOrange200Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"deepOrange200", 13},
		{"deepOrange200 ", 13},
		{"deepOrange200\n", 13},
		{"deepOrange200.", 13},
		{"deepOrange200:", 13},
		{"deepOrange200,", 13},
		{"deepOrange200\"", 13},
		{"deepOrange200(", 13},
		{"deepOrange200)", 13},
		{"deepOrange200[", 13},
		{"deepOrange200]", 13},
		{"deepOrange200// comment", 13},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.DeepOrange200()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestDeepOrange200Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_d"},
		{"d_"},
		{"_de"},
		{"de_"},
		{"_dee"},
		{"dee_"},
		{"_deep"},
		{"deep_"},
		{"_deepO"},
		{"deepO_"},
		{"_deepOr"},
		{"deepOr_"},
		{"_deepOra"},
		{"deepOra_"},
		{"_deepOran"},
		{"deepOran_"},
		{"_deepOrang"},
		{"deepOrang_"},
		{"_deepOrange"},
		{"deepOrange_"},
		{"_deepOrange2"},
		{"deepOrange2_"},
		{"_deepOrange20"},
		{"deepOrange20_"},
		{"_deepOrange200"},
		{"deepOrange200_"},
		{"9713d"},
		{"d9713"},
		{"9713de"},
		{"de9713"},
		{"9713dee"},
		{"dee9713"},
		{"9713deep"},
		{"deep9713"},
		{"9713deepO"},
		{"deepO9713"},
		{"9713deepOr"},
		{"deepOr9713"},
		{"9713deepOra"},
		{"deepOra9713"},
		{"9713deepOran"},
		{"deepOran9713"},
		{"9713deepOrang"},
		{"deepOrang9713"},
		{"9713deepOrange"},
		{"deepOrange9713"},
		{"9713deepOrange2"},
		{"deepOrange29713"},
		{"9713deepOrange20"},
		{"deepOrange209713"},
		{"9713deepOrange200"},
		{"deepOrange2009713"},
		{"Zd"},
		{"dZ"},
		{"Zde"},
		{"deZ"},
		{"Zdee"},
		{"deeZ"},
		{"Zdeep"},
		{"deepZ"},
		{"ZdeepO"},
		{"deepOZ"},
		{"ZdeepOr"},
		{"deepOrZ"},
		{"ZdeepOra"},
		{"deepOraZ"},
		{"ZdeepOran"},
		{"deepOranZ"},
		{"ZdeepOrang"},
		{"deepOrangZ"},
		{"ZdeepOrange"},
		{"deepOrangeZ"},
		{"ZdeepOrange2"},
		{"deepOrange2Z"},
		{"ZdeepOrange20"},
		{"deepOrange20Z"},
		{"ZdeepOrange200"},
		{"deepOrange200Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.DeepOrange200()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestDeepOrange300Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"deepOrange300", 13},
		{"deepOrange300 ", 13},
		{"deepOrange300\n", 13},
		{"deepOrange300.", 13},
		{"deepOrange300:", 13},
		{"deepOrange300,", 13},
		{"deepOrange300\"", 13},
		{"deepOrange300(", 13},
		{"deepOrange300)", 13},
		{"deepOrange300[", 13},
		{"deepOrange300]", 13},
		{"deepOrange300// comment", 13},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.DeepOrange300()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestDeepOrange300Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_d"},
		{"d_"},
		{"_de"},
		{"de_"},
		{"_dee"},
		{"dee_"},
		{"_deep"},
		{"deep_"},
		{"_deepO"},
		{"deepO_"},
		{"_deepOr"},
		{"deepOr_"},
		{"_deepOra"},
		{"deepOra_"},
		{"_deepOran"},
		{"deepOran_"},
		{"_deepOrang"},
		{"deepOrang_"},
		{"_deepOrange"},
		{"deepOrange_"},
		{"_deepOrange3"},
		{"deepOrange3_"},
		{"_deepOrange30"},
		{"deepOrange30_"},
		{"_deepOrange300"},
		{"deepOrange300_"},
		{"9713d"},
		{"d9713"},
		{"9713de"},
		{"de9713"},
		{"9713dee"},
		{"dee9713"},
		{"9713deep"},
		{"deep9713"},
		{"9713deepO"},
		{"deepO9713"},
		{"9713deepOr"},
		{"deepOr9713"},
		{"9713deepOra"},
		{"deepOra9713"},
		{"9713deepOran"},
		{"deepOran9713"},
		{"9713deepOrang"},
		{"deepOrang9713"},
		{"9713deepOrange"},
		{"deepOrange9713"},
		{"9713deepOrange3"},
		{"deepOrange39713"},
		{"9713deepOrange30"},
		{"deepOrange309713"},
		{"9713deepOrange300"},
		{"deepOrange3009713"},
		{"Zd"},
		{"dZ"},
		{"Zde"},
		{"deZ"},
		{"Zdee"},
		{"deeZ"},
		{"Zdeep"},
		{"deepZ"},
		{"ZdeepO"},
		{"deepOZ"},
		{"ZdeepOr"},
		{"deepOrZ"},
		{"ZdeepOra"},
		{"deepOraZ"},
		{"ZdeepOran"},
		{"deepOranZ"},
		{"ZdeepOrang"},
		{"deepOrangZ"},
		{"ZdeepOrange"},
		{"deepOrangeZ"},
		{"ZdeepOrange3"},
		{"deepOrange3Z"},
		{"ZdeepOrange30"},
		{"deepOrange30Z"},
		{"ZdeepOrange300"},
		{"deepOrange300Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.DeepOrange300()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestDeepOrange400Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"deepOrange400", 13},
		{"deepOrange400 ", 13},
		{"deepOrange400\n", 13},
		{"deepOrange400.", 13},
		{"deepOrange400:", 13},
		{"deepOrange400,", 13},
		{"deepOrange400\"", 13},
		{"deepOrange400(", 13},
		{"deepOrange400)", 13},
		{"deepOrange400[", 13},
		{"deepOrange400]", 13},
		{"deepOrange400// comment", 13},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.DeepOrange400()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestDeepOrange400Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_d"},
		{"d_"},
		{"_de"},
		{"de_"},
		{"_dee"},
		{"dee_"},
		{"_deep"},
		{"deep_"},
		{"_deepO"},
		{"deepO_"},
		{"_deepOr"},
		{"deepOr_"},
		{"_deepOra"},
		{"deepOra_"},
		{"_deepOran"},
		{"deepOran_"},
		{"_deepOrang"},
		{"deepOrang_"},
		{"_deepOrange"},
		{"deepOrange_"},
		{"_deepOrange4"},
		{"deepOrange4_"},
		{"_deepOrange40"},
		{"deepOrange40_"},
		{"_deepOrange400"},
		{"deepOrange400_"},
		{"9713d"},
		{"d9713"},
		{"9713de"},
		{"de9713"},
		{"9713dee"},
		{"dee9713"},
		{"9713deep"},
		{"deep9713"},
		{"9713deepO"},
		{"deepO9713"},
		{"9713deepOr"},
		{"deepOr9713"},
		{"9713deepOra"},
		{"deepOra9713"},
		{"9713deepOran"},
		{"deepOran9713"},
		{"9713deepOrang"},
		{"deepOrang9713"},
		{"9713deepOrange"},
		{"deepOrange9713"},
		{"9713deepOrange4"},
		{"deepOrange49713"},
		{"9713deepOrange40"},
		{"deepOrange409713"},
		{"9713deepOrange400"},
		{"deepOrange4009713"},
		{"Zd"},
		{"dZ"},
		{"Zde"},
		{"deZ"},
		{"Zdee"},
		{"deeZ"},
		{"Zdeep"},
		{"deepZ"},
		{"ZdeepO"},
		{"deepOZ"},
		{"ZdeepOr"},
		{"deepOrZ"},
		{"ZdeepOra"},
		{"deepOraZ"},
		{"ZdeepOran"},
		{"deepOranZ"},
		{"ZdeepOrang"},
		{"deepOrangZ"},
		{"ZdeepOrange"},
		{"deepOrangeZ"},
		{"ZdeepOrange4"},
		{"deepOrange4Z"},
		{"ZdeepOrange40"},
		{"deepOrange40Z"},
		{"ZdeepOrange400"},
		{"deepOrange400Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.DeepOrange400()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestDeepOrange500Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"deepOrange500", 13},
		{"deepOrange500 ", 13},
		{"deepOrange500\n", 13},
		{"deepOrange500.", 13},
		{"deepOrange500:", 13},
		{"deepOrange500,", 13},
		{"deepOrange500\"", 13},
		{"deepOrange500(", 13},
		{"deepOrange500)", 13},
		{"deepOrange500[", 13},
		{"deepOrange500]", 13},
		{"deepOrange500// comment", 13},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.DeepOrange500()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestDeepOrange500Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_d"},
		{"d_"},
		{"_de"},
		{"de_"},
		{"_dee"},
		{"dee_"},
		{"_deep"},
		{"deep_"},
		{"_deepO"},
		{"deepO_"},
		{"_deepOr"},
		{"deepOr_"},
		{"_deepOra"},
		{"deepOra_"},
		{"_deepOran"},
		{"deepOran_"},
		{"_deepOrang"},
		{"deepOrang_"},
		{"_deepOrange"},
		{"deepOrange_"},
		{"_deepOrange5"},
		{"deepOrange5_"},
		{"_deepOrange50"},
		{"deepOrange50_"},
		{"_deepOrange500"},
		{"deepOrange500_"},
		{"9713d"},
		{"d9713"},
		{"9713de"},
		{"de9713"},
		{"9713dee"},
		{"dee9713"},
		{"9713deep"},
		{"deep9713"},
		{"9713deepO"},
		{"deepO9713"},
		{"9713deepOr"},
		{"deepOr9713"},
		{"9713deepOra"},
		{"deepOra9713"},
		{"9713deepOran"},
		{"deepOran9713"},
		{"9713deepOrang"},
		{"deepOrang9713"},
		{"9713deepOrange"},
		{"deepOrange9713"},
		{"9713deepOrange5"},
		{"deepOrange59713"},
		{"9713deepOrange50"},
		{"deepOrange509713"},
		{"9713deepOrange500"},
		{"deepOrange5009713"},
		{"Zd"},
		{"dZ"},
		{"Zde"},
		{"deZ"},
		{"Zdee"},
		{"deeZ"},
		{"Zdeep"},
		{"deepZ"},
		{"ZdeepO"},
		{"deepOZ"},
		{"ZdeepOr"},
		{"deepOrZ"},
		{"ZdeepOra"},
		{"deepOraZ"},
		{"ZdeepOran"},
		{"deepOranZ"},
		{"ZdeepOrang"},
		{"deepOrangZ"},
		{"ZdeepOrange"},
		{"deepOrangeZ"},
		{"ZdeepOrange5"},
		{"deepOrange5Z"},
		{"ZdeepOrange50"},
		{"deepOrange50Z"},
		{"ZdeepOrange500"},
		{"deepOrange500Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.DeepOrange500()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestDeepOrange600Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"deepOrange600", 13},
		{"deepOrange600 ", 13},
		{"deepOrange600\n", 13},
		{"deepOrange600.", 13},
		{"deepOrange600:", 13},
		{"deepOrange600,", 13},
		{"deepOrange600\"", 13},
		{"deepOrange600(", 13},
		{"deepOrange600)", 13},
		{"deepOrange600[", 13},
		{"deepOrange600]", 13},
		{"deepOrange600// comment", 13},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.DeepOrange600()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestDeepOrange600Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_d"},
		{"d_"},
		{"_de"},
		{"de_"},
		{"_dee"},
		{"dee_"},
		{"_deep"},
		{"deep_"},
		{"_deepO"},
		{"deepO_"},
		{"_deepOr"},
		{"deepOr_"},
		{"_deepOra"},
		{"deepOra_"},
		{"_deepOran"},
		{"deepOran_"},
		{"_deepOrang"},
		{"deepOrang_"},
		{"_deepOrange"},
		{"deepOrange_"},
		{"_deepOrange6"},
		{"deepOrange6_"},
		{"_deepOrange60"},
		{"deepOrange60_"},
		{"_deepOrange600"},
		{"deepOrange600_"},
		{"9713d"},
		{"d9713"},
		{"9713de"},
		{"de9713"},
		{"9713dee"},
		{"dee9713"},
		{"9713deep"},
		{"deep9713"},
		{"9713deepO"},
		{"deepO9713"},
		{"9713deepOr"},
		{"deepOr9713"},
		{"9713deepOra"},
		{"deepOra9713"},
		{"9713deepOran"},
		{"deepOran9713"},
		{"9713deepOrang"},
		{"deepOrang9713"},
		{"9713deepOrange"},
		{"deepOrange9713"},
		{"9713deepOrange6"},
		{"deepOrange69713"},
		{"9713deepOrange60"},
		{"deepOrange609713"},
		{"9713deepOrange600"},
		{"deepOrange6009713"},
		{"Zd"},
		{"dZ"},
		{"Zde"},
		{"deZ"},
		{"Zdee"},
		{"deeZ"},
		{"Zdeep"},
		{"deepZ"},
		{"ZdeepO"},
		{"deepOZ"},
		{"ZdeepOr"},
		{"deepOrZ"},
		{"ZdeepOra"},
		{"deepOraZ"},
		{"ZdeepOran"},
		{"deepOranZ"},
		{"ZdeepOrang"},
		{"deepOrangZ"},
		{"ZdeepOrange"},
		{"deepOrangeZ"},
		{"ZdeepOrange6"},
		{"deepOrange6Z"},
		{"ZdeepOrange60"},
		{"deepOrange60Z"},
		{"ZdeepOrange600"},
		{"deepOrange600Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.DeepOrange600()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestDeepOrange700Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"deepOrange700", 13},
		{"deepOrange700 ", 13},
		{"deepOrange700\n", 13},
		{"deepOrange700.", 13},
		{"deepOrange700:", 13},
		{"deepOrange700,", 13},
		{"deepOrange700\"", 13},
		{"deepOrange700(", 13},
		{"deepOrange700)", 13},
		{"deepOrange700[", 13},
		{"deepOrange700]", 13},
		{"deepOrange700// comment", 13},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.DeepOrange700()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestDeepOrange700Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_d"},
		{"d_"},
		{"_de"},
		{"de_"},
		{"_dee"},
		{"dee_"},
		{"_deep"},
		{"deep_"},
		{"_deepO"},
		{"deepO_"},
		{"_deepOr"},
		{"deepOr_"},
		{"_deepOra"},
		{"deepOra_"},
		{"_deepOran"},
		{"deepOran_"},
		{"_deepOrang"},
		{"deepOrang_"},
		{"_deepOrange"},
		{"deepOrange_"},
		{"_deepOrange7"},
		{"deepOrange7_"},
		{"_deepOrange70"},
		{"deepOrange70_"},
		{"_deepOrange700"},
		{"deepOrange700_"},
		{"9713d"},
		{"d9713"},
		{"9713de"},
		{"de9713"},
		{"9713dee"},
		{"dee9713"},
		{"9713deep"},
		{"deep9713"},
		{"9713deepO"},
		{"deepO9713"},
		{"9713deepOr"},
		{"deepOr9713"},
		{"9713deepOra"},
		{"deepOra9713"},
		{"9713deepOran"},
		{"deepOran9713"},
		{"9713deepOrang"},
		{"deepOrang9713"},
		{"9713deepOrange"},
		{"deepOrange9713"},
		{"9713deepOrange7"},
		{"deepOrange79713"},
		{"9713deepOrange70"},
		{"deepOrange709713"},
		{"9713deepOrange700"},
		{"deepOrange7009713"},
		{"Zd"},
		{"dZ"},
		{"Zde"},
		{"deZ"},
		{"Zdee"},
		{"deeZ"},
		{"Zdeep"},
		{"deepZ"},
		{"ZdeepO"},
		{"deepOZ"},
		{"ZdeepOr"},
		{"deepOrZ"},
		{"ZdeepOra"},
		{"deepOraZ"},
		{"ZdeepOran"},
		{"deepOranZ"},
		{"ZdeepOrang"},
		{"deepOrangZ"},
		{"ZdeepOrange"},
		{"deepOrangeZ"},
		{"ZdeepOrange7"},
		{"deepOrange7Z"},
		{"ZdeepOrange70"},
		{"deepOrange70Z"},
		{"ZdeepOrange700"},
		{"deepOrange700Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.DeepOrange700()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestDeepOrange800Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"deepOrange800", 13},
		{"deepOrange800 ", 13},
		{"deepOrange800\n", 13},
		{"deepOrange800.", 13},
		{"deepOrange800:", 13},
		{"deepOrange800,", 13},
		{"deepOrange800\"", 13},
		{"deepOrange800(", 13},
		{"deepOrange800)", 13},
		{"deepOrange800[", 13},
		{"deepOrange800]", 13},
		{"deepOrange800// comment", 13},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.DeepOrange800()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestDeepOrange800Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_d"},
		{"d_"},
		{"_de"},
		{"de_"},
		{"_dee"},
		{"dee_"},
		{"_deep"},
		{"deep_"},
		{"_deepO"},
		{"deepO_"},
		{"_deepOr"},
		{"deepOr_"},
		{"_deepOra"},
		{"deepOra_"},
		{"_deepOran"},
		{"deepOran_"},
		{"_deepOrang"},
		{"deepOrang_"},
		{"_deepOrange"},
		{"deepOrange_"},
		{"_deepOrange8"},
		{"deepOrange8_"},
		{"_deepOrange80"},
		{"deepOrange80_"},
		{"_deepOrange800"},
		{"deepOrange800_"},
		{"9713d"},
		{"d9713"},
		{"9713de"},
		{"de9713"},
		{"9713dee"},
		{"dee9713"},
		{"9713deep"},
		{"deep9713"},
		{"9713deepO"},
		{"deepO9713"},
		{"9713deepOr"},
		{"deepOr9713"},
		{"9713deepOra"},
		{"deepOra9713"},
		{"9713deepOran"},
		{"deepOran9713"},
		{"9713deepOrang"},
		{"deepOrang9713"},
		{"9713deepOrange"},
		{"deepOrange9713"},
		{"9713deepOrange8"},
		{"deepOrange89713"},
		{"9713deepOrange80"},
		{"deepOrange809713"},
		{"9713deepOrange800"},
		{"deepOrange8009713"},
		{"Zd"},
		{"dZ"},
		{"Zde"},
		{"deZ"},
		{"Zdee"},
		{"deeZ"},
		{"Zdeep"},
		{"deepZ"},
		{"ZdeepO"},
		{"deepOZ"},
		{"ZdeepOr"},
		{"deepOrZ"},
		{"ZdeepOra"},
		{"deepOraZ"},
		{"ZdeepOran"},
		{"deepOranZ"},
		{"ZdeepOrang"},
		{"deepOrangZ"},
		{"ZdeepOrange"},
		{"deepOrangeZ"},
		{"ZdeepOrange8"},
		{"deepOrange8Z"},
		{"ZdeepOrange80"},
		{"deepOrange80Z"},
		{"ZdeepOrange800"},
		{"deepOrange800Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.DeepOrange800()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestDeepOrange900Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"deepOrange900", 13},
		{"deepOrange900 ", 13},
		{"deepOrange900\n", 13},
		{"deepOrange900.", 13},
		{"deepOrange900:", 13},
		{"deepOrange900,", 13},
		{"deepOrange900\"", 13},
		{"deepOrange900(", 13},
		{"deepOrange900)", 13},
		{"deepOrange900[", 13},
		{"deepOrange900]", 13},
		{"deepOrange900// comment", 13},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.DeepOrange900()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestDeepOrange900Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_d"},
		{"d_"},
		{"_de"},
		{"de_"},
		{"_dee"},
		{"dee_"},
		{"_deep"},
		{"deep_"},
		{"_deepO"},
		{"deepO_"},
		{"_deepOr"},
		{"deepOr_"},
		{"_deepOra"},
		{"deepOra_"},
		{"_deepOran"},
		{"deepOran_"},
		{"_deepOrang"},
		{"deepOrang_"},
		{"_deepOrange"},
		{"deepOrange_"},
		{"_deepOrange9"},
		{"deepOrange9_"},
		{"_deepOrange90"},
		{"deepOrange90_"},
		{"_deepOrange900"},
		{"deepOrange900_"},
		{"9713d"},
		{"d9713"},
		{"9713de"},
		{"de9713"},
		{"9713dee"},
		{"dee9713"},
		{"9713deep"},
		{"deep9713"},
		{"9713deepO"},
		{"deepO9713"},
		{"9713deepOr"},
		{"deepOr9713"},
		{"9713deepOra"},
		{"deepOra9713"},
		{"9713deepOran"},
		{"deepOran9713"},
		{"9713deepOrang"},
		{"deepOrang9713"},
		{"9713deepOrange"},
		{"deepOrange9713"},
		{"9713deepOrange9"},
		{"deepOrange99713"},
		{"9713deepOrange90"},
		{"deepOrange909713"},
		{"9713deepOrange900"},
		{"deepOrange9009713"},
		{"Zd"},
		{"dZ"},
		{"Zde"},
		{"deZ"},
		{"Zdee"},
		{"deeZ"},
		{"Zdeep"},
		{"deepZ"},
		{"ZdeepO"},
		{"deepOZ"},
		{"ZdeepOr"},
		{"deepOrZ"},
		{"ZdeepOra"},
		{"deepOraZ"},
		{"ZdeepOran"},
		{"deepOranZ"},
		{"ZdeepOrang"},
		{"deepOrangZ"},
		{"ZdeepOrange"},
		{"deepOrangeZ"},
		{"ZdeepOrange9"},
		{"deepOrange9Z"},
		{"ZdeepOrange90"},
		{"deepOrange90Z"},
		{"ZdeepOrange900"},
		{"deepOrange900Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.DeepOrange900()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestDeepOrangeA100Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"deepOrangeA100", 14},
		{"deepOrangeA100 ", 14},
		{"deepOrangeA100\n", 14},
		{"deepOrangeA100.", 14},
		{"deepOrangeA100:", 14},
		{"deepOrangeA100,", 14},
		{"deepOrangeA100\"", 14},
		{"deepOrangeA100(", 14},
		{"deepOrangeA100)", 14},
		{"deepOrangeA100[", 14},
		{"deepOrangeA100]", 14},
		{"deepOrangeA100// comment", 14},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.DeepOrangeA100()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestDeepOrangeA100Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_d"},
		{"d_"},
		{"_de"},
		{"de_"},
		{"_dee"},
		{"dee_"},
		{"_deep"},
		{"deep_"},
		{"_deepO"},
		{"deepO_"},
		{"_deepOr"},
		{"deepOr_"},
		{"_deepOra"},
		{"deepOra_"},
		{"_deepOran"},
		{"deepOran_"},
		{"_deepOrang"},
		{"deepOrang_"},
		{"_deepOrange"},
		{"deepOrange_"},
		{"_deepOrangeA"},
		{"deepOrangeA_"},
		{"_deepOrangeA1"},
		{"deepOrangeA1_"},
		{"_deepOrangeA10"},
		{"deepOrangeA10_"},
		{"_deepOrangeA100"},
		{"deepOrangeA100_"},
		{"9713d"},
		{"d9713"},
		{"9713de"},
		{"de9713"},
		{"9713dee"},
		{"dee9713"},
		{"9713deep"},
		{"deep9713"},
		{"9713deepO"},
		{"deepO9713"},
		{"9713deepOr"},
		{"deepOr9713"},
		{"9713deepOra"},
		{"deepOra9713"},
		{"9713deepOran"},
		{"deepOran9713"},
		{"9713deepOrang"},
		{"deepOrang9713"},
		{"9713deepOrange"},
		{"deepOrange9713"},
		{"9713deepOrangeA"},
		{"deepOrangeA9713"},
		{"9713deepOrangeA1"},
		{"deepOrangeA19713"},
		{"9713deepOrangeA10"},
		{"deepOrangeA109713"},
		{"9713deepOrangeA100"},
		{"deepOrangeA1009713"},
		{"Zd"},
		{"dZ"},
		{"Zde"},
		{"deZ"},
		{"Zdee"},
		{"deeZ"},
		{"Zdeep"},
		{"deepZ"},
		{"ZdeepO"},
		{"deepOZ"},
		{"ZdeepOr"},
		{"deepOrZ"},
		{"ZdeepOra"},
		{"deepOraZ"},
		{"ZdeepOran"},
		{"deepOranZ"},
		{"ZdeepOrang"},
		{"deepOrangZ"},
		{"ZdeepOrange"},
		{"deepOrangeZ"},
		{"ZdeepOrangeA"},
		{"deepOrangeAZ"},
		{"ZdeepOrangeA1"},
		{"deepOrangeA1Z"},
		{"ZdeepOrangeA10"},
		{"deepOrangeA10Z"},
		{"ZdeepOrangeA100"},
		{"deepOrangeA100Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.DeepOrangeA100()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestDeepOrangeA200Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"deepOrangeA200", 14},
		{"deepOrangeA200 ", 14},
		{"deepOrangeA200\n", 14},
		{"deepOrangeA200.", 14},
		{"deepOrangeA200:", 14},
		{"deepOrangeA200,", 14},
		{"deepOrangeA200\"", 14},
		{"deepOrangeA200(", 14},
		{"deepOrangeA200)", 14},
		{"deepOrangeA200[", 14},
		{"deepOrangeA200]", 14},
		{"deepOrangeA200// comment", 14},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.DeepOrangeA200()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestDeepOrangeA200Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_d"},
		{"d_"},
		{"_de"},
		{"de_"},
		{"_dee"},
		{"dee_"},
		{"_deep"},
		{"deep_"},
		{"_deepO"},
		{"deepO_"},
		{"_deepOr"},
		{"deepOr_"},
		{"_deepOra"},
		{"deepOra_"},
		{"_deepOran"},
		{"deepOran_"},
		{"_deepOrang"},
		{"deepOrang_"},
		{"_deepOrange"},
		{"deepOrange_"},
		{"_deepOrangeA"},
		{"deepOrangeA_"},
		{"_deepOrangeA2"},
		{"deepOrangeA2_"},
		{"_deepOrangeA20"},
		{"deepOrangeA20_"},
		{"_deepOrangeA200"},
		{"deepOrangeA200_"},
		{"9713d"},
		{"d9713"},
		{"9713de"},
		{"de9713"},
		{"9713dee"},
		{"dee9713"},
		{"9713deep"},
		{"deep9713"},
		{"9713deepO"},
		{"deepO9713"},
		{"9713deepOr"},
		{"deepOr9713"},
		{"9713deepOra"},
		{"deepOra9713"},
		{"9713deepOran"},
		{"deepOran9713"},
		{"9713deepOrang"},
		{"deepOrang9713"},
		{"9713deepOrange"},
		{"deepOrange9713"},
		{"9713deepOrangeA"},
		{"deepOrangeA9713"},
		{"9713deepOrangeA2"},
		{"deepOrangeA29713"},
		{"9713deepOrangeA20"},
		{"deepOrangeA209713"},
		{"9713deepOrangeA200"},
		{"deepOrangeA2009713"},
		{"Zd"},
		{"dZ"},
		{"Zde"},
		{"deZ"},
		{"Zdee"},
		{"deeZ"},
		{"Zdeep"},
		{"deepZ"},
		{"ZdeepO"},
		{"deepOZ"},
		{"ZdeepOr"},
		{"deepOrZ"},
		{"ZdeepOra"},
		{"deepOraZ"},
		{"ZdeepOran"},
		{"deepOranZ"},
		{"ZdeepOrang"},
		{"deepOrangZ"},
		{"ZdeepOrange"},
		{"deepOrangeZ"},
		{"ZdeepOrangeA"},
		{"deepOrangeAZ"},
		{"ZdeepOrangeA2"},
		{"deepOrangeA2Z"},
		{"ZdeepOrangeA20"},
		{"deepOrangeA20Z"},
		{"ZdeepOrangeA200"},
		{"deepOrangeA200Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.DeepOrangeA200()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestDeepOrangeA400Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"deepOrangeA400", 14},
		{"deepOrangeA400 ", 14},
		{"deepOrangeA400\n", 14},
		{"deepOrangeA400.", 14},
		{"deepOrangeA400:", 14},
		{"deepOrangeA400,", 14},
		{"deepOrangeA400\"", 14},
		{"deepOrangeA400(", 14},
		{"deepOrangeA400)", 14},
		{"deepOrangeA400[", 14},
		{"deepOrangeA400]", 14},
		{"deepOrangeA400// comment", 14},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.DeepOrangeA400()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestDeepOrangeA400Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_d"},
		{"d_"},
		{"_de"},
		{"de_"},
		{"_dee"},
		{"dee_"},
		{"_deep"},
		{"deep_"},
		{"_deepO"},
		{"deepO_"},
		{"_deepOr"},
		{"deepOr_"},
		{"_deepOra"},
		{"deepOra_"},
		{"_deepOran"},
		{"deepOran_"},
		{"_deepOrang"},
		{"deepOrang_"},
		{"_deepOrange"},
		{"deepOrange_"},
		{"_deepOrangeA"},
		{"deepOrangeA_"},
		{"_deepOrangeA4"},
		{"deepOrangeA4_"},
		{"_deepOrangeA40"},
		{"deepOrangeA40_"},
		{"_deepOrangeA400"},
		{"deepOrangeA400_"},
		{"9713d"},
		{"d9713"},
		{"9713de"},
		{"de9713"},
		{"9713dee"},
		{"dee9713"},
		{"9713deep"},
		{"deep9713"},
		{"9713deepO"},
		{"deepO9713"},
		{"9713deepOr"},
		{"deepOr9713"},
		{"9713deepOra"},
		{"deepOra9713"},
		{"9713deepOran"},
		{"deepOran9713"},
		{"9713deepOrang"},
		{"deepOrang9713"},
		{"9713deepOrange"},
		{"deepOrange9713"},
		{"9713deepOrangeA"},
		{"deepOrangeA9713"},
		{"9713deepOrangeA4"},
		{"deepOrangeA49713"},
		{"9713deepOrangeA40"},
		{"deepOrangeA409713"},
		{"9713deepOrangeA400"},
		{"deepOrangeA4009713"},
		{"Zd"},
		{"dZ"},
		{"Zde"},
		{"deZ"},
		{"Zdee"},
		{"deeZ"},
		{"Zdeep"},
		{"deepZ"},
		{"ZdeepO"},
		{"deepOZ"},
		{"ZdeepOr"},
		{"deepOrZ"},
		{"ZdeepOra"},
		{"deepOraZ"},
		{"ZdeepOran"},
		{"deepOranZ"},
		{"ZdeepOrang"},
		{"deepOrangZ"},
		{"ZdeepOrange"},
		{"deepOrangeZ"},
		{"ZdeepOrangeA"},
		{"deepOrangeAZ"},
		{"ZdeepOrangeA4"},
		{"deepOrangeA4Z"},
		{"ZdeepOrangeA40"},
		{"deepOrangeA40Z"},
		{"ZdeepOrangeA400"},
		{"deepOrangeA400Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.DeepOrangeA400()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestDeepOrangeA700Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"deepOrangeA700", 14},
		{"deepOrangeA700 ", 14},
		{"deepOrangeA700\n", 14},
		{"deepOrangeA700.", 14},
		{"deepOrangeA700:", 14},
		{"deepOrangeA700,", 14},
		{"deepOrangeA700\"", 14},
		{"deepOrangeA700(", 14},
		{"deepOrangeA700)", 14},
		{"deepOrangeA700[", 14},
		{"deepOrangeA700]", 14},
		{"deepOrangeA700// comment", 14},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.DeepOrangeA700()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestDeepOrangeA700Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_d"},
		{"d_"},
		{"_de"},
		{"de_"},
		{"_dee"},
		{"dee_"},
		{"_deep"},
		{"deep_"},
		{"_deepO"},
		{"deepO_"},
		{"_deepOr"},
		{"deepOr_"},
		{"_deepOra"},
		{"deepOra_"},
		{"_deepOran"},
		{"deepOran_"},
		{"_deepOrang"},
		{"deepOrang_"},
		{"_deepOrange"},
		{"deepOrange_"},
		{"_deepOrangeA"},
		{"deepOrangeA_"},
		{"_deepOrangeA7"},
		{"deepOrangeA7_"},
		{"_deepOrangeA70"},
		{"deepOrangeA70_"},
		{"_deepOrangeA700"},
		{"deepOrangeA700_"},
		{"9713d"},
		{"d9713"},
		{"9713de"},
		{"de9713"},
		{"9713dee"},
		{"dee9713"},
		{"9713deep"},
		{"deep9713"},
		{"9713deepO"},
		{"deepO9713"},
		{"9713deepOr"},
		{"deepOr9713"},
		{"9713deepOra"},
		{"deepOra9713"},
		{"9713deepOran"},
		{"deepOran9713"},
		{"9713deepOrang"},
		{"deepOrang9713"},
		{"9713deepOrange"},
		{"deepOrange9713"},
		{"9713deepOrangeA"},
		{"deepOrangeA9713"},
		{"9713deepOrangeA7"},
		{"deepOrangeA79713"},
		{"9713deepOrangeA70"},
		{"deepOrangeA709713"},
		{"9713deepOrangeA700"},
		{"deepOrangeA7009713"},
		{"Zd"},
		{"dZ"},
		{"Zde"},
		{"deZ"},
		{"Zdee"},
		{"deeZ"},
		{"Zdeep"},
		{"deepZ"},
		{"ZdeepO"},
		{"deepOZ"},
		{"ZdeepOr"},
		{"deepOrZ"},
		{"ZdeepOra"},
		{"deepOraZ"},
		{"ZdeepOran"},
		{"deepOranZ"},
		{"ZdeepOrang"},
		{"deepOrangZ"},
		{"ZdeepOrange"},
		{"deepOrangeZ"},
		{"ZdeepOrangeA"},
		{"deepOrangeAZ"},
		{"ZdeepOrangeA7"},
		{"deepOrangeA7Z"},
		{"ZdeepOrangeA70"},
		{"deepOrangeA70Z"},
		{"ZdeepOrangeA700"},
		{"deepOrangeA700Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.DeepOrangeA700()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestBrown50Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"brown50", 7},
		{"brown50 ", 7},
		{"brown50\n", 7},
		{"brown50.", 7},
		{"brown50:", 7},
		{"brown50,", 7},
		{"brown50\"", 7},
		{"brown50(", 7},
		{"brown50)", 7},
		{"brown50[", 7},
		{"brown50]", 7},
		{"brown50// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Brown50()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestBrown50Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_b"},
		{"b_"},
		{"_br"},
		{"br_"},
		{"_bro"},
		{"bro_"},
		{"_brow"},
		{"brow_"},
		{"_brown"},
		{"brown_"},
		{"_brown5"},
		{"brown5_"},
		{"_brown50"},
		{"brown50_"},
		{"9713b"},
		{"b9713"},
		{"9713br"},
		{"br9713"},
		{"9713bro"},
		{"bro9713"},
		{"9713brow"},
		{"brow9713"},
		{"9713brown"},
		{"brown9713"},
		{"9713brown5"},
		{"brown59713"},
		{"9713brown50"},
		{"brown509713"},
		{"Zb"},
		{"bZ"},
		{"Zbr"},
		{"brZ"},
		{"Zbro"},
		{"broZ"},
		{"Zbrow"},
		{"browZ"},
		{"Zbrown"},
		{"brownZ"},
		{"Zbrown5"},
		{"brown5Z"},
		{"Zbrown50"},
		{"brown50Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Brown50()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestBrown100Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"brown100", 8},
		{"brown100 ", 8},
		{"brown100\n", 8},
		{"brown100.", 8},
		{"brown100:", 8},
		{"brown100,", 8},
		{"brown100\"", 8},
		{"brown100(", 8},
		{"brown100)", 8},
		{"brown100[", 8},
		{"brown100]", 8},
		{"brown100// comment", 8},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Brown100()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestBrown100Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_b"},
		{"b_"},
		{"_br"},
		{"br_"},
		{"_bro"},
		{"bro_"},
		{"_brow"},
		{"brow_"},
		{"_brown"},
		{"brown_"},
		{"_brown1"},
		{"brown1_"},
		{"_brown10"},
		{"brown10_"},
		{"_brown100"},
		{"brown100_"},
		{"9713b"},
		{"b9713"},
		{"9713br"},
		{"br9713"},
		{"9713bro"},
		{"bro9713"},
		{"9713brow"},
		{"brow9713"},
		{"9713brown"},
		{"brown9713"},
		{"9713brown1"},
		{"brown19713"},
		{"9713brown10"},
		{"brown109713"},
		{"9713brown100"},
		{"brown1009713"},
		{"Zb"},
		{"bZ"},
		{"Zbr"},
		{"brZ"},
		{"Zbro"},
		{"broZ"},
		{"Zbrow"},
		{"browZ"},
		{"Zbrown"},
		{"brownZ"},
		{"Zbrown1"},
		{"brown1Z"},
		{"Zbrown10"},
		{"brown10Z"},
		{"Zbrown100"},
		{"brown100Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Brown100()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestBrown200Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"brown200", 8},
		{"brown200 ", 8},
		{"brown200\n", 8},
		{"brown200.", 8},
		{"brown200:", 8},
		{"brown200,", 8},
		{"brown200\"", 8},
		{"brown200(", 8},
		{"brown200)", 8},
		{"brown200[", 8},
		{"brown200]", 8},
		{"brown200// comment", 8},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Brown200()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestBrown200Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_b"},
		{"b_"},
		{"_br"},
		{"br_"},
		{"_bro"},
		{"bro_"},
		{"_brow"},
		{"brow_"},
		{"_brown"},
		{"brown_"},
		{"_brown2"},
		{"brown2_"},
		{"_brown20"},
		{"brown20_"},
		{"_brown200"},
		{"brown200_"},
		{"9713b"},
		{"b9713"},
		{"9713br"},
		{"br9713"},
		{"9713bro"},
		{"bro9713"},
		{"9713brow"},
		{"brow9713"},
		{"9713brown"},
		{"brown9713"},
		{"9713brown2"},
		{"brown29713"},
		{"9713brown20"},
		{"brown209713"},
		{"9713brown200"},
		{"brown2009713"},
		{"Zb"},
		{"bZ"},
		{"Zbr"},
		{"brZ"},
		{"Zbro"},
		{"broZ"},
		{"Zbrow"},
		{"browZ"},
		{"Zbrown"},
		{"brownZ"},
		{"Zbrown2"},
		{"brown2Z"},
		{"Zbrown20"},
		{"brown20Z"},
		{"Zbrown200"},
		{"brown200Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Brown200()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestBrown300Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"brown300", 8},
		{"brown300 ", 8},
		{"brown300\n", 8},
		{"brown300.", 8},
		{"brown300:", 8},
		{"brown300,", 8},
		{"brown300\"", 8},
		{"brown300(", 8},
		{"brown300)", 8},
		{"brown300[", 8},
		{"brown300]", 8},
		{"brown300// comment", 8},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Brown300()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestBrown300Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_b"},
		{"b_"},
		{"_br"},
		{"br_"},
		{"_bro"},
		{"bro_"},
		{"_brow"},
		{"brow_"},
		{"_brown"},
		{"brown_"},
		{"_brown3"},
		{"brown3_"},
		{"_brown30"},
		{"brown30_"},
		{"_brown300"},
		{"brown300_"},
		{"9713b"},
		{"b9713"},
		{"9713br"},
		{"br9713"},
		{"9713bro"},
		{"bro9713"},
		{"9713brow"},
		{"brow9713"},
		{"9713brown"},
		{"brown9713"},
		{"9713brown3"},
		{"brown39713"},
		{"9713brown30"},
		{"brown309713"},
		{"9713brown300"},
		{"brown3009713"},
		{"Zb"},
		{"bZ"},
		{"Zbr"},
		{"brZ"},
		{"Zbro"},
		{"broZ"},
		{"Zbrow"},
		{"browZ"},
		{"Zbrown"},
		{"brownZ"},
		{"Zbrown3"},
		{"brown3Z"},
		{"Zbrown30"},
		{"brown30Z"},
		{"Zbrown300"},
		{"brown300Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Brown300()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestBrown400Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"brown400", 8},
		{"brown400 ", 8},
		{"brown400\n", 8},
		{"brown400.", 8},
		{"brown400:", 8},
		{"brown400,", 8},
		{"brown400\"", 8},
		{"brown400(", 8},
		{"brown400)", 8},
		{"brown400[", 8},
		{"brown400]", 8},
		{"brown400// comment", 8},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Brown400()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestBrown400Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_b"},
		{"b_"},
		{"_br"},
		{"br_"},
		{"_bro"},
		{"bro_"},
		{"_brow"},
		{"brow_"},
		{"_brown"},
		{"brown_"},
		{"_brown4"},
		{"brown4_"},
		{"_brown40"},
		{"brown40_"},
		{"_brown400"},
		{"brown400_"},
		{"9713b"},
		{"b9713"},
		{"9713br"},
		{"br9713"},
		{"9713bro"},
		{"bro9713"},
		{"9713brow"},
		{"brow9713"},
		{"9713brown"},
		{"brown9713"},
		{"9713brown4"},
		{"brown49713"},
		{"9713brown40"},
		{"brown409713"},
		{"9713brown400"},
		{"brown4009713"},
		{"Zb"},
		{"bZ"},
		{"Zbr"},
		{"brZ"},
		{"Zbro"},
		{"broZ"},
		{"Zbrow"},
		{"browZ"},
		{"Zbrown"},
		{"brownZ"},
		{"Zbrown4"},
		{"brown4Z"},
		{"Zbrown40"},
		{"brown40Z"},
		{"Zbrown400"},
		{"brown400Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Brown400()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestBrown500Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"brown500", 8},
		{"brown500 ", 8},
		{"brown500\n", 8},
		{"brown500.", 8},
		{"brown500:", 8},
		{"brown500,", 8},
		{"brown500\"", 8},
		{"brown500(", 8},
		{"brown500)", 8},
		{"brown500[", 8},
		{"brown500]", 8},
		{"brown500// comment", 8},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Brown500()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestBrown500Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_b"},
		{"b_"},
		{"_br"},
		{"br_"},
		{"_bro"},
		{"bro_"},
		{"_brow"},
		{"brow_"},
		{"_brown"},
		{"brown_"},
		{"_brown5"},
		{"brown5_"},
		{"_brown50"},
		{"brown50_"},
		{"_brown500"},
		{"brown500_"},
		{"9713b"},
		{"b9713"},
		{"9713br"},
		{"br9713"},
		{"9713bro"},
		{"bro9713"},
		{"9713brow"},
		{"brow9713"},
		{"9713brown"},
		{"brown9713"},
		{"9713brown5"},
		{"brown59713"},
		{"9713brown50"},
		{"brown509713"},
		{"9713brown500"},
		{"brown5009713"},
		{"Zb"},
		{"bZ"},
		{"Zbr"},
		{"brZ"},
		{"Zbro"},
		{"broZ"},
		{"Zbrow"},
		{"browZ"},
		{"Zbrown"},
		{"brownZ"},
		{"Zbrown5"},
		{"brown5Z"},
		{"Zbrown50"},
		{"brown50Z"},
		{"Zbrown500"},
		{"brown500Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Brown500()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestBrown600Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"brown600", 8},
		{"brown600 ", 8},
		{"brown600\n", 8},
		{"brown600.", 8},
		{"brown600:", 8},
		{"brown600,", 8},
		{"brown600\"", 8},
		{"brown600(", 8},
		{"brown600)", 8},
		{"brown600[", 8},
		{"brown600]", 8},
		{"brown600// comment", 8},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Brown600()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestBrown600Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_b"},
		{"b_"},
		{"_br"},
		{"br_"},
		{"_bro"},
		{"bro_"},
		{"_brow"},
		{"brow_"},
		{"_brown"},
		{"brown_"},
		{"_brown6"},
		{"brown6_"},
		{"_brown60"},
		{"brown60_"},
		{"_brown600"},
		{"brown600_"},
		{"9713b"},
		{"b9713"},
		{"9713br"},
		{"br9713"},
		{"9713bro"},
		{"bro9713"},
		{"9713brow"},
		{"brow9713"},
		{"9713brown"},
		{"brown9713"},
		{"9713brown6"},
		{"brown69713"},
		{"9713brown60"},
		{"brown609713"},
		{"9713brown600"},
		{"brown6009713"},
		{"Zb"},
		{"bZ"},
		{"Zbr"},
		{"brZ"},
		{"Zbro"},
		{"broZ"},
		{"Zbrow"},
		{"browZ"},
		{"Zbrown"},
		{"brownZ"},
		{"Zbrown6"},
		{"brown6Z"},
		{"Zbrown60"},
		{"brown60Z"},
		{"Zbrown600"},
		{"brown600Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Brown600()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestBrown700Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"brown700", 8},
		{"brown700 ", 8},
		{"brown700\n", 8},
		{"brown700.", 8},
		{"brown700:", 8},
		{"brown700,", 8},
		{"brown700\"", 8},
		{"brown700(", 8},
		{"brown700)", 8},
		{"brown700[", 8},
		{"brown700]", 8},
		{"brown700// comment", 8},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Brown700()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestBrown700Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_b"},
		{"b_"},
		{"_br"},
		{"br_"},
		{"_bro"},
		{"bro_"},
		{"_brow"},
		{"brow_"},
		{"_brown"},
		{"brown_"},
		{"_brown7"},
		{"brown7_"},
		{"_brown70"},
		{"brown70_"},
		{"_brown700"},
		{"brown700_"},
		{"9713b"},
		{"b9713"},
		{"9713br"},
		{"br9713"},
		{"9713bro"},
		{"bro9713"},
		{"9713brow"},
		{"brow9713"},
		{"9713brown"},
		{"brown9713"},
		{"9713brown7"},
		{"brown79713"},
		{"9713brown70"},
		{"brown709713"},
		{"9713brown700"},
		{"brown7009713"},
		{"Zb"},
		{"bZ"},
		{"Zbr"},
		{"brZ"},
		{"Zbro"},
		{"broZ"},
		{"Zbrow"},
		{"browZ"},
		{"Zbrown"},
		{"brownZ"},
		{"Zbrown7"},
		{"brown7Z"},
		{"Zbrown70"},
		{"brown70Z"},
		{"Zbrown700"},
		{"brown700Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Brown700()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestBrown800Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"brown800", 8},
		{"brown800 ", 8},
		{"brown800\n", 8},
		{"brown800.", 8},
		{"brown800:", 8},
		{"brown800,", 8},
		{"brown800\"", 8},
		{"brown800(", 8},
		{"brown800)", 8},
		{"brown800[", 8},
		{"brown800]", 8},
		{"brown800// comment", 8},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Brown800()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestBrown800Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_b"},
		{"b_"},
		{"_br"},
		{"br_"},
		{"_bro"},
		{"bro_"},
		{"_brow"},
		{"brow_"},
		{"_brown"},
		{"brown_"},
		{"_brown8"},
		{"brown8_"},
		{"_brown80"},
		{"brown80_"},
		{"_brown800"},
		{"brown800_"},
		{"9713b"},
		{"b9713"},
		{"9713br"},
		{"br9713"},
		{"9713bro"},
		{"bro9713"},
		{"9713brow"},
		{"brow9713"},
		{"9713brown"},
		{"brown9713"},
		{"9713brown8"},
		{"brown89713"},
		{"9713brown80"},
		{"brown809713"},
		{"9713brown800"},
		{"brown8009713"},
		{"Zb"},
		{"bZ"},
		{"Zbr"},
		{"brZ"},
		{"Zbro"},
		{"broZ"},
		{"Zbrow"},
		{"browZ"},
		{"Zbrown"},
		{"brownZ"},
		{"Zbrown8"},
		{"brown8Z"},
		{"Zbrown80"},
		{"brown80Z"},
		{"Zbrown800"},
		{"brown800Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Brown800()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestBrown900Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"brown900", 8},
		{"brown900 ", 8},
		{"brown900\n", 8},
		{"brown900.", 8},
		{"brown900:", 8},
		{"brown900,", 8},
		{"brown900\"", 8},
		{"brown900(", 8},
		{"brown900)", 8},
		{"brown900[", 8},
		{"brown900]", 8},
		{"brown900// comment", 8},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Brown900()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestBrown900Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_b"},
		{"b_"},
		{"_br"},
		{"br_"},
		{"_bro"},
		{"bro_"},
		{"_brow"},
		{"brow_"},
		{"_brown"},
		{"brown_"},
		{"_brown9"},
		{"brown9_"},
		{"_brown90"},
		{"brown90_"},
		{"_brown900"},
		{"brown900_"},
		{"9713b"},
		{"b9713"},
		{"9713br"},
		{"br9713"},
		{"9713bro"},
		{"bro9713"},
		{"9713brow"},
		{"brow9713"},
		{"9713brown"},
		{"brown9713"},
		{"9713brown9"},
		{"brown99713"},
		{"9713brown90"},
		{"brown909713"},
		{"9713brown900"},
		{"brown9009713"},
		{"Zb"},
		{"bZ"},
		{"Zbr"},
		{"brZ"},
		{"Zbro"},
		{"broZ"},
		{"Zbrow"},
		{"browZ"},
		{"Zbrown"},
		{"brownZ"},
		{"Zbrown9"},
		{"brown9Z"},
		{"Zbrown90"},
		{"brown90Z"},
		{"Zbrown900"},
		{"brown900Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Brown900()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestGrey50Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"grey50", 6},
		{"grey50 ", 6},
		{"grey50\n", 6},
		{"grey50.", 6},
		{"grey50:", 6},
		{"grey50,", 6},
		{"grey50\"", 6},
		{"grey50(", 6},
		{"grey50)", 6},
		{"grey50[", 6},
		{"grey50]", 6},
		{"grey50// comment", 6},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Grey50()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestGrey50Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_g"},
		{"g_"},
		{"_gr"},
		{"gr_"},
		{"_gre"},
		{"gre_"},
		{"_grey"},
		{"grey_"},
		{"_grey5"},
		{"grey5_"},
		{"_grey50"},
		{"grey50_"},
		{"9713g"},
		{"g9713"},
		{"9713gr"},
		{"gr9713"},
		{"9713gre"},
		{"gre9713"},
		{"9713grey"},
		{"grey9713"},
		{"9713grey5"},
		{"grey59713"},
		{"9713grey50"},
		{"grey509713"},
		{"Zg"},
		{"gZ"},
		{"Zgr"},
		{"grZ"},
		{"Zgre"},
		{"greZ"},
		{"Zgrey"},
		{"greyZ"},
		{"Zgrey5"},
		{"grey5Z"},
		{"Zgrey50"},
		{"grey50Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Grey50()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestGrey100Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"grey100", 7},
		{"grey100 ", 7},
		{"grey100\n", 7},
		{"grey100.", 7},
		{"grey100:", 7},
		{"grey100,", 7},
		{"grey100\"", 7},
		{"grey100(", 7},
		{"grey100)", 7},
		{"grey100[", 7},
		{"grey100]", 7},
		{"grey100// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Grey100()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestGrey100Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_g"},
		{"g_"},
		{"_gr"},
		{"gr_"},
		{"_gre"},
		{"gre_"},
		{"_grey"},
		{"grey_"},
		{"_grey1"},
		{"grey1_"},
		{"_grey10"},
		{"grey10_"},
		{"_grey100"},
		{"grey100_"},
		{"9713g"},
		{"g9713"},
		{"9713gr"},
		{"gr9713"},
		{"9713gre"},
		{"gre9713"},
		{"9713grey"},
		{"grey9713"},
		{"9713grey1"},
		{"grey19713"},
		{"9713grey10"},
		{"grey109713"},
		{"9713grey100"},
		{"grey1009713"},
		{"Zg"},
		{"gZ"},
		{"Zgr"},
		{"grZ"},
		{"Zgre"},
		{"greZ"},
		{"Zgrey"},
		{"greyZ"},
		{"Zgrey1"},
		{"grey1Z"},
		{"Zgrey10"},
		{"grey10Z"},
		{"Zgrey100"},
		{"grey100Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Grey100()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestGrey200Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"grey200", 7},
		{"grey200 ", 7},
		{"grey200\n", 7},
		{"grey200.", 7},
		{"grey200:", 7},
		{"grey200,", 7},
		{"grey200\"", 7},
		{"grey200(", 7},
		{"grey200)", 7},
		{"grey200[", 7},
		{"grey200]", 7},
		{"grey200// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Grey200()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestGrey200Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_g"},
		{"g_"},
		{"_gr"},
		{"gr_"},
		{"_gre"},
		{"gre_"},
		{"_grey"},
		{"grey_"},
		{"_grey2"},
		{"grey2_"},
		{"_grey20"},
		{"grey20_"},
		{"_grey200"},
		{"grey200_"},
		{"9713g"},
		{"g9713"},
		{"9713gr"},
		{"gr9713"},
		{"9713gre"},
		{"gre9713"},
		{"9713grey"},
		{"grey9713"},
		{"9713grey2"},
		{"grey29713"},
		{"9713grey20"},
		{"grey209713"},
		{"9713grey200"},
		{"grey2009713"},
		{"Zg"},
		{"gZ"},
		{"Zgr"},
		{"grZ"},
		{"Zgre"},
		{"greZ"},
		{"Zgrey"},
		{"greyZ"},
		{"Zgrey2"},
		{"grey2Z"},
		{"Zgrey20"},
		{"grey20Z"},
		{"Zgrey200"},
		{"grey200Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Grey200()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestGrey300Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"grey300", 7},
		{"grey300 ", 7},
		{"grey300\n", 7},
		{"grey300.", 7},
		{"grey300:", 7},
		{"grey300,", 7},
		{"grey300\"", 7},
		{"grey300(", 7},
		{"grey300)", 7},
		{"grey300[", 7},
		{"grey300]", 7},
		{"grey300// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Grey300()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestGrey300Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_g"},
		{"g_"},
		{"_gr"},
		{"gr_"},
		{"_gre"},
		{"gre_"},
		{"_grey"},
		{"grey_"},
		{"_grey3"},
		{"grey3_"},
		{"_grey30"},
		{"grey30_"},
		{"_grey300"},
		{"grey300_"},
		{"9713g"},
		{"g9713"},
		{"9713gr"},
		{"gr9713"},
		{"9713gre"},
		{"gre9713"},
		{"9713grey"},
		{"grey9713"},
		{"9713grey3"},
		{"grey39713"},
		{"9713grey30"},
		{"grey309713"},
		{"9713grey300"},
		{"grey3009713"},
		{"Zg"},
		{"gZ"},
		{"Zgr"},
		{"grZ"},
		{"Zgre"},
		{"greZ"},
		{"Zgrey"},
		{"greyZ"},
		{"Zgrey3"},
		{"grey3Z"},
		{"Zgrey30"},
		{"grey30Z"},
		{"Zgrey300"},
		{"grey300Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Grey300()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestGrey400Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"grey400", 7},
		{"grey400 ", 7},
		{"grey400\n", 7},
		{"grey400.", 7},
		{"grey400:", 7},
		{"grey400,", 7},
		{"grey400\"", 7},
		{"grey400(", 7},
		{"grey400)", 7},
		{"grey400[", 7},
		{"grey400]", 7},
		{"grey400// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Grey400()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestGrey400Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_g"},
		{"g_"},
		{"_gr"},
		{"gr_"},
		{"_gre"},
		{"gre_"},
		{"_grey"},
		{"grey_"},
		{"_grey4"},
		{"grey4_"},
		{"_grey40"},
		{"grey40_"},
		{"_grey400"},
		{"grey400_"},
		{"9713g"},
		{"g9713"},
		{"9713gr"},
		{"gr9713"},
		{"9713gre"},
		{"gre9713"},
		{"9713grey"},
		{"grey9713"},
		{"9713grey4"},
		{"grey49713"},
		{"9713grey40"},
		{"grey409713"},
		{"9713grey400"},
		{"grey4009713"},
		{"Zg"},
		{"gZ"},
		{"Zgr"},
		{"grZ"},
		{"Zgre"},
		{"greZ"},
		{"Zgrey"},
		{"greyZ"},
		{"Zgrey4"},
		{"grey4Z"},
		{"Zgrey40"},
		{"grey40Z"},
		{"Zgrey400"},
		{"grey400Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Grey400()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestGrey500Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"grey500", 7},
		{"grey500 ", 7},
		{"grey500\n", 7},
		{"grey500.", 7},
		{"grey500:", 7},
		{"grey500,", 7},
		{"grey500\"", 7},
		{"grey500(", 7},
		{"grey500)", 7},
		{"grey500[", 7},
		{"grey500]", 7},
		{"grey500// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Grey500()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestGrey500Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_g"},
		{"g_"},
		{"_gr"},
		{"gr_"},
		{"_gre"},
		{"gre_"},
		{"_grey"},
		{"grey_"},
		{"_grey5"},
		{"grey5_"},
		{"_grey50"},
		{"grey50_"},
		{"_grey500"},
		{"grey500_"},
		{"9713g"},
		{"g9713"},
		{"9713gr"},
		{"gr9713"},
		{"9713gre"},
		{"gre9713"},
		{"9713grey"},
		{"grey9713"},
		{"9713grey5"},
		{"grey59713"},
		{"9713grey50"},
		{"grey509713"},
		{"9713grey500"},
		{"grey5009713"},
		{"Zg"},
		{"gZ"},
		{"Zgr"},
		{"grZ"},
		{"Zgre"},
		{"greZ"},
		{"Zgrey"},
		{"greyZ"},
		{"Zgrey5"},
		{"grey5Z"},
		{"Zgrey50"},
		{"grey50Z"},
		{"Zgrey500"},
		{"grey500Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Grey500()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestGrey600Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"grey600", 7},
		{"grey600 ", 7},
		{"grey600\n", 7},
		{"grey600.", 7},
		{"grey600:", 7},
		{"grey600,", 7},
		{"grey600\"", 7},
		{"grey600(", 7},
		{"grey600)", 7},
		{"grey600[", 7},
		{"grey600]", 7},
		{"grey600// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Grey600()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestGrey600Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_g"},
		{"g_"},
		{"_gr"},
		{"gr_"},
		{"_gre"},
		{"gre_"},
		{"_grey"},
		{"grey_"},
		{"_grey6"},
		{"grey6_"},
		{"_grey60"},
		{"grey60_"},
		{"_grey600"},
		{"grey600_"},
		{"9713g"},
		{"g9713"},
		{"9713gr"},
		{"gr9713"},
		{"9713gre"},
		{"gre9713"},
		{"9713grey"},
		{"grey9713"},
		{"9713grey6"},
		{"grey69713"},
		{"9713grey60"},
		{"grey609713"},
		{"9713grey600"},
		{"grey6009713"},
		{"Zg"},
		{"gZ"},
		{"Zgr"},
		{"grZ"},
		{"Zgre"},
		{"greZ"},
		{"Zgrey"},
		{"greyZ"},
		{"Zgrey6"},
		{"grey6Z"},
		{"Zgrey60"},
		{"grey60Z"},
		{"Zgrey600"},
		{"grey600Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Grey600()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestGrey700Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"grey700", 7},
		{"grey700 ", 7},
		{"grey700\n", 7},
		{"grey700.", 7},
		{"grey700:", 7},
		{"grey700,", 7},
		{"grey700\"", 7},
		{"grey700(", 7},
		{"grey700)", 7},
		{"grey700[", 7},
		{"grey700]", 7},
		{"grey700// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Grey700()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestGrey700Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_g"},
		{"g_"},
		{"_gr"},
		{"gr_"},
		{"_gre"},
		{"gre_"},
		{"_grey"},
		{"grey_"},
		{"_grey7"},
		{"grey7_"},
		{"_grey70"},
		{"grey70_"},
		{"_grey700"},
		{"grey700_"},
		{"9713g"},
		{"g9713"},
		{"9713gr"},
		{"gr9713"},
		{"9713gre"},
		{"gre9713"},
		{"9713grey"},
		{"grey9713"},
		{"9713grey7"},
		{"grey79713"},
		{"9713grey70"},
		{"grey709713"},
		{"9713grey700"},
		{"grey7009713"},
		{"Zg"},
		{"gZ"},
		{"Zgr"},
		{"grZ"},
		{"Zgre"},
		{"greZ"},
		{"Zgrey"},
		{"greyZ"},
		{"Zgrey7"},
		{"grey7Z"},
		{"Zgrey70"},
		{"grey70Z"},
		{"Zgrey700"},
		{"grey700Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Grey700()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestGrey800Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"grey800", 7},
		{"grey800 ", 7},
		{"grey800\n", 7},
		{"grey800.", 7},
		{"grey800:", 7},
		{"grey800,", 7},
		{"grey800\"", 7},
		{"grey800(", 7},
		{"grey800)", 7},
		{"grey800[", 7},
		{"grey800]", 7},
		{"grey800// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Grey800()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestGrey800Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_g"},
		{"g_"},
		{"_gr"},
		{"gr_"},
		{"_gre"},
		{"gre_"},
		{"_grey"},
		{"grey_"},
		{"_grey8"},
		{"grey8_"},
		{"_grey80"},
		{"grey80_"},
		{"_grey800"},
		{"grey800_"},
		{"9713g"},
		{"g9713"},
		{"9713gr"},
		{"gr9713"},
		{"9713gre"},
		{"gre9713"},
		{"9713grey"},
		{"grey9713"},
		{"9713grey8"},
		{"grey89713"},
		{"9713grey80"},
		{"grey809713"},
		{"9713grey800"},
		{"grey8009713"},
		{"Zg"},
		{"gZ"},
		{"Zgr"},
		{"grZ"},
		{"Zgre"},
		{"greZ"},
		{"Zgrey"},
		{"greyZ"},
		{"Zgrey8"},
		{"grey8Z"},
		{"Zgrey80"},
		{"grey80Z"},
		{"Zgrey800"},
		{"grey800Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Grey800()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestGrey900Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"grey900", 7},
		{"grey900 ", 7},
		{"grey900\n", 7},
		{"grey900.", 7},
		{"grey900:", 7},
		{"grey900,", 7},
		{"grey900\"", 7},
		{"grey900(", 7},
		{"grey900)", 7},
		{"grey900[", 7},
		{"grey900]", 7},
		{"grey900// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Grey900()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestGrey900Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_g"},
		{"g_"},
		{"_gr"},
		{"gr_"},
		{"_gre"},
		{"gre_"},
		{"_grey"},
		{"grey_"},
		{"_grey9"},
		{"grey9_"},
		{"_grey90"},
		{"grey90_"},
		{"_grey900"},
		{"grey900_"},
		{"9713g"},
		{"g9713"},
		{"9713gr"},
		{"gr9713"},
		{"9713gre"},
		{"gre9713"},
		{"9713grey"},
		{"grey9713"},
		{"9713grey9"},
		{"grey99713"},
		{"9713grey90"},
		{"grey909713"},
		{"9713grey900"},
		{"grey9009713"},
		{"Zg"},
		{"gZ"},
		{"Zgr"},
		{"grZ"},
		{"Zgre"},
		{"greZ"},
		{"Zgrey"},
		{"greyZ"},
		{"Zgrey9"},
		{"grey9Z"},
		{"Zgrey90"},
		{"grey90Z"},
		{"Zgrey900"},
		{"grey900Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Grey900()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestBlueGrey50Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"blueGrey50", 10},
		{"blueGrey50 ", 10},
		{"blueGrey50\n", 10},
		{"blueGrey50.", 10},
		{"blueGrey50:", 10},
		{"blueGrey50,", 10},
		{"blueGrey50\"", 10},
		{"blueGrey50(", 10},
		{"blueGrey50)", 10},
		{"blueGrey50[", 10},
		{"blueGrey50]", 10},
		{"blueGrey50// comment", 10},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.BlueGrey50()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestBlueGrey50Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_b"},
		{"b_"},
		{"_bl"},
		{"bl_"},
		{"_blu"},
		{"blu_"},
		{"_blue"},
		{"blue_"},
		{"_blueG"},
		{"blueG_"},
		{"_blueGr"},
		{"blueGr_"},
		{"_blueGre"},
		{"blueGre_"},
		{"_blueGrey"},
		{"blueGrey_"},
		{"_blueGrey5"},
		{"blueGrey5_"},
		{"_blueGrey50"},
		{"blueGrey50_"},
		{"9713b"},
		{"b9713"},
		{"9713bl"},
		{"bl9713"},
		{"9713blu"},
		{"blu9713"},
		{"9713blue"},
		{"blue9713"},
		{"9713blueG"},
		{"blueG9713"},
		{"9713blueGr"},
		{"blueGr9713"},
		{"9713blueGre"},
		{"blueGre9713"},
		{"9713blueGrey"},
		{"blueGrey9713"},
		{"9713blueGrey5"},
		{"blueGrey59713"},
		{"9713blueGrey50"},
		{"blueGrey509713"},
		{"Zb"},
		{"bZ"},
		{"Zbl"},
		{"blZ"},
		{"Zblu"},
		{"bluZ"},
		{"Zblue"},
		{"blueZ"},
		{"ZblueG"},
		{"blueGZ"},
		{"ZblueGr"},
		{"blueGrZ"},
		{"ZblueGre"},
		{"blueGreZ"},
		{"ZblueGrey"},
		{"blueGreyZ"},
		{"ZblueGrey5"},
		{"blueGrey5Z"},
		{"ZblueGrey50"},
		{"blueGrey50Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.BlueGrey50()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestBlueGrey100Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"blueGrey100", 11},
		{"blueGrey100 ", 11},
		{"blueGrey100\n", 11},
		{"blueGrey100.", 11},
		{"blueGrey100:", 11},
		{"blueGrey100,", 11},
		{"blueGrey100\"", 11},
		{"blueGrey100(", 11},
		{"blueGrey100)", 11},
		{"blueGrey100[", 11},
		{"blueGrey100]", 11},
		{"blueGrey100// comment", 11},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.BlueGrey100()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestBlueGrey100Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_b"},
		{"b_"},
		{"_bl"},
		{"bl_"},
		{"_blu"},
		{"blu_"},
		{"_blue"},
		{"blue_"},
		{"_blueG"},
		{"blueG_"},
		{"_blueGr"},
		{"blueGr_"},
		{"_blueGre"},
		{"blueGre_"},
		{"_blueGrey"},
		{"blueGrey_"},
		{"_blueGrey1"},
		{"blueGrey1_"},
		{"_blueGrey10"},
		{"blueGrey10_"},
		{"_blueGrey100"},
		{"blueGrey100_"},
		{"9713b"},
		{"b9713"},
		{"9713bl"},
		{"bl9713"},
		{"9713blu"},
		{"blu9713"},
		{"9713blue"},
		{"blue9713"},
		{"9713blueG"},
		{"blueG9713"},
		{"9713blueGr"},
		{"blueGr9713"},
		{"9713blueGre"},
		{"blueGre9713"},
		{"9713blueGrey"},
		{"blueGrey9713"},
		{"9713blueGrey1"},
		{"blueGrey19713"},
		{"9713blueGrey10"},
		{"blueGrey109713"},
		{"9713blueGrey100"},
		{"blueGrey1009713"},
		{"Zb"},
		{"bZ"},
		{"Zbl"},
		{"blZ"},
		{"Zblu"},
		{"bluZ"},
		{"Zblue"},
		{"blueZ"},
		{"ZblueG"},
		{"blueGZ"},
		{"ZblueGr"},
		{"blueGrZ"},
		{"ZblueGre"},
		{"blueGreZ"},
		{"ZblueGrey"},
		{"blueGreyZ"},
		{"ZblueGrey1"},
		{"blueGrey1Z"},
		{"ZblueGrey10"},
		{"blueGrey10Z"},
		{"ZblueGrey100"},
		{"blueGrey100Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.BlueGrey100()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestBlueGrey200Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"blueGrey200", 11},
		{"blueGrey200 ", 11},
		{"blueGrey200\n", 11},
		{"blueGrey200.", 11},
		{"blueGrey200:", 11},
		{"blueGrey200,", 11},
		{"blueGrey200\"", 11},
		{"blueGrey200(", 11},
		{"blueGrey200)", 11},
		{"blueGrey200[", 11},
		{"blueGrey200]", 11},
		{"blueGrey200// comment", 11},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.BlueGrey200()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestBlueGrey200Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_b"},
		{"b_"},
		{"_bl"},
		{"bl_"},
		{"_blu"},
		{"blu_"},
		{"_blue"},
		{"blue_"},
		{"_blueG"},
		{"blueG_"},
		{"_blueGr"},
		{"blueGr_"},
		{"_blueGre"},
		{"blueGre_"},
		{"_blueGrey"},
		{"blueGrey_"},
		{"_blueGrey2"},
		{"blueGrey2_"},
		{"_blueGrey20"},
		{"blueGrey20_"},
		{"_blueGrey200"},
		{"blueGrey200_"},
		{"9713b"},
		{"b9713"},
		{"9713bl"},
		{"bl9713"},
		{"9713blu"},
		{"blu9713"},
		{"9713blue"},
		{"blue9713"},
		{"9713blueG"},
		{"blueG9713"},
		{"9713blueGr"},
		{"blueGr9713"},
		{"9713blueGre"},
		{"blueGre9713"},
		{"9713blueGrey"},
		{"blueGrey9713"},
		{"9713blueGrey2"},
		{"blueGrey29713"},
		{"9713blueGrey20"},
		{"blueGrey209713"},
		{"9713blueGrey200"},
		{"blueGrey2009713"},
		{"Zb"},
		{"bZ"},
		{"Zbl"},
		{"blZ"},
		{"Zblu"},
		{"bluZ"},
		{"Zblue"},
		{"blueZ"},
		{"ZblueG"},
		{"blueGZ"},
		{"ZblueGr"},
		{"blueGrZ"},
		{"ZblueGre"},
		{"blueGreZ"},
		{"ZblueGrey"},
		{"blueGreyZ"},
		{"ZblueGrey2"},
		{"blueGrey2Z"},
		{"ZblueGrey20"},
		{"blueGrey20Z"},
		{"ZblueGrey200"},
		{"blueGrey200Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.BlueGrey200()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestBlueGrey300Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"blueGrey300", 11},
		{"blueGrey300 ", 11},
		{"blueGrey300\n", 11},
		{"blueGrey300.", 11},
		{"blueGrey300:", 11},
		{"blueGrey300,", 11},
		{"blueGrey300\"", 11},
		{"blueGrey300(", 11},
		{"blueGrey300)", 11},
		{"blueGrey300[", 11},
		{"blueGrey300]", 11},
		{"blueGrey300// comment", 11},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.BlueGrey300()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestBlueGrey300Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_b"},
		{"b_"},
		{"_bl"},
		{"bl_"},
		{"_blu"},
		{"blu_"},
		{"_blue"},
		{"blue_"},
		{"_blueG"},
		{"blueG_"},
		{"_blueGr"},
		{"blueGr_"},
		{"_blueGre"},
		{"blueGre_"},
		{"_blueGrey"},
		{"blueGrey_"},
		{"_blueGrey3"},
		{"blueGrey3_"},
		{"_blueGrey30"},
		{"blueGrey30_"},
		{"_blueGrey300"},
		{"blueGrey300_"},
		{"9713b"},
		{"b9713"},
		{"9713bl"},
		{"bl9713"},
		{"9713blu"},
		{"blu9713"},
		{"9713blue"},
		{"blue9713"},
		{"9713blueG"},
		{"blueG9713"},
		{"9713blueGr"},
		{"blueGr9713"},
		{"9713blueGre"},
		{"blueGre9713"},
		{"9713blueGrey"},
		{"blueGrey9713"},
		{"9713blueGrey3"},
		{"blueGrey39713"},
		{"9713blueGrey30"},
		{"blueGrey309713"},
		{"9713blueGrey300"},
		{"blueGrey3009713"},
		{"Zb"},
		{"bZ"},
		{"Zbl"},
		{"blZ"},
		{"Zblu"},
		{"bluZ"},
		{"Zblue"},
		{"blueZ"},
		{"ZblueG"},
		{"blueGZ"},
		{"ZblueGr"},
		{"blueGrZ"},
		{"ZblueGre"},
		{"blueGreZ"},
		{"ZblueGrey"},
		{"blueGreyZ"},
		{"ZblueGrey3"},
		{"blueGrey3Z"},
		{"ZblueGrey30"},
		{"blueGrey30Z"},
		{"ZblueGrey300"},
		{"blueGrey300Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.BlueGrey300()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestBlueGrey400Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"blueGrey400", 11},
		{"blueGrey400 ", 11},
		{"blueGrey400\n", 11},
		{"blueGrey400.", 11},
		{"blueGrey400:", 11},
		{"blueGrey400,", 11},
		{"blueGrey400\"", 11},
		{"blueGrey400(", 11},
		{"blueGrey400)", 11},
		{"blueGrey400[", 11},
		{"blueGrey400]", 11},
		{"blueGrey400// comment", 11},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.BlueGrey400()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestBlueGrey400Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_b"},
		{"b_"},
		{"_bl"},
		{"bl_"},
		{"_blu"},
		{"blu_"},
		{"_blue"},
		{"blue_"},
		{"_blueG"},
		{"blueG_"},
		{"_blueGr"},
		{"blueGr_"},
		{"_blueGre"},
		{"blueGre_"},
		{"_blueGrey"},
		{"blueGrey_"},
		{"_blueGrey4"},
		{"blueGrey4_"},
		{"_blueGrey40"},
		{"blueGrey40_"},
		{"_blueGrey400"},
		{"blueGrey400_"},
		{"9713b"},
		{"b9713"},
		{"9713bl"},
		{"bl9713"},
		{"9713blu"},
		{"blu9713"},
		{"9713blue"},
		{"blue9713"},
		{"9713blueG"},
		{"blueG9713"},
		{"9713blueGr"},
		{"blueGr9713"},
		{"9713blueGre"},
		{"blueGre9713"},
		{"9713blueGrey"},
		{"blueGrey9713"},
		{"9713blueGrey4"},
		{"blueGrey49713"},
		{"9713blueGrey40"},
		{"blueGrey409713"},
		{"9713blueGrey400"},
		{"blueGrey4009713"},
		{"Zb"},
		{"bZ"},
		{"Zbl"},
		{"blZ"},
		{"Zblu"},
		{"bluZ"},
		{"Zblue"},
		{"blueZ"},
		{"ZblueG"},
		{"blueGZ"},
		{"ZblueGr"},
		{"blueGrZ"},
		{"ZblueGre"},
		{"blueGreZ"},
		{"ZblueGrey"},
		{"blueGreyZ"},
		{"ZblueGrey4"},
		{"blueGrey4Z"},
		{"ZblueGrey40"},
		{"blueGrey40Z"},
		{"ZblueGrey400"},
		{"blueGrey400Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.BlueGrey400()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestBlueGrey500Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"blueGrey500", 11},
		{"blueGrey500 ", 11},
		{"blueGrey500\n", 11},
		{"blueGrey500.", 11},
		{"blueGrey500:", 11},
		{"blueGrey500,", 11},
		{"blueGrey500\"", 11},
		{"blueGrey500(", 11},
		{"blueGrey500)", 11},
		{"blueGrey500[", 11},
		{"blueGrey500]", 11},
		{"blueGrey500// comment", 11},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.BlueGrey500()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestBlueGrey500Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_b"},
		{"b_"},
		{"_bl"},
		{"bl_"},
		{"_blu"},
		{"blu_"},
		{"_blue"},
		{"blue_"},
		{"_blueG"},
		{"blueG_"},
		{"_blueGr"},
		{"blueGr_"},
		{"_blueGre"},
		{"blueGre_"},
		{"_blueGrey"},
		{"blueGrey_"},
		{"_blueGrey5"},
		{"blueGrey5_"},
		{"_blueGrey50"},
		{"blueGrey50_"},
		{"_blueGrey500"},
		{"blueGrey500_"},
		{"9713b"},
		{"b9713"},
		{"9713bl"},
		{"bl9713"},
		{"9713blu"},
		{"blu9713"},
		{"9713blue"},
		{"blue9713"},
		{"9713blueG"},
		{"blueG9713"},
		{"9713blueGr"},
		{"blueGr9713"},
		{"9713blueGre"},
		{"blueGre9713"},
		{"9713blueGrey"},
		{"blueGrey9713"},
		{"9713blueGrey5"},
		{"blueGrey59713"},
		{"9713blueGrey50"},
		{"blueGrey509713"},
		{"9713blueGrey500"},
		{"blueGrey5009713"},
		{"Zb"},
		{"bZ"},
		{"Zbl"},
		{"blZ"},
		{"Zblu"},
		{"bluZ"},
		{"Zblue"},
		{"blueZ"},
		{"ZblueG"},
		{"blueGZ"},
		{"ZblueGr"},
		{"blueGrZ"},
		{"ZblueGre"},
		{"blueGreZ"},
		{"ZblueGrey"},
		{"blueGreyZ"},
		{"ZblueGrey5"},
		{"blueGrey5Z"},
		{"ZblueGrey50"},
		{"blueGrey50Z"},
		{"ZblueGrey500"},
		{"blueGrey500Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.BlueGrey500()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestBlueGrey600Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"blueGrey600", 11},
		{"blueGrey600 ", 11},
		{"blueGrey600\n", 11},
		{"blueGrey600.", 11},
		{"blueGrey600:", 11},
		{"blueGrey600,", 11},
		{"blueGrey600\"", 11},
		{"blueGrey600(", 11},
		{"blueGrey600)", 11},
		{"blueGrey600[", 11},
		{"blueGrey600]", 11},
		{"blueGrey600// comment", 11},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.BlueGrey600()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestBlueGrey600Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_b"},
		{"b_"},
		{"_bl"},
		{"bl_"},
		{"_blu"},
		{"blu_"},
		{"_blue"},
		{"blue_"},
		{"_blueG"},
		{"blueG_"},
		{"_blueGr"},
		{"blueGr_"},
		{"_blueGre"},
		{"blueGre_"},
		{"_blueGrey"},
		{"blueGrey_"},
		{"_blueGrey6"},
		{"blueGrey6_"},
		{"_blueGrey60"},
		{"blueGrey60_"},
		{"_blueGrey600"},
		{"blueGrey600_"},
		{"9713b"},
		{"b9713"},
		{"9713bl"},
		{"bl9713"},
		{"9713blu"},
		{"blu9713"},
		{"9713blue"},
		{"blue9713"},
		{"9713blueG"},
		{"blueG9713"},
		{"9713blueGr"},
		{"blueGr9713"},
		{"9713blueGre"},
		{"blueGre9713"},
		{"9713blueGrey"},
		{"blueGrey9713"},
		{"9713blueGrey6"},
		{"blueGrey69713"},
		{"9713blueGrey60"},
		{"blueGrey609713"},
		{"9713blueGrey600"},
		{"blueGrey6009713"},
		{"Zb"},
		{"bZ"},
		{"Zbl"},
		{"blZ"},
		{"Zblu"},
		{"bluZ"},
		{"Zblue"},
		{"blueZ"},
		{"ZblueG"},
		{"blueGZ"},
		{"ZblueGr"},
		{"blueGrZ"},
		{"ZblueGre"},
		{"blueGreZ"},
		{"ZblueGrey"},
		{"blueGreyZ"},
		{"ZblueGrey6"},
		{"blueGrey6Z"},
		{"ZblueGrey60"},
		{"blueGrey60Z"},
		{"ZblueGrey600"},
		{"blueGrey600Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.BlueGrey600()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestBlueGrey700Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"blueGrey700", 11},
		{"blueGrey700 ", 11},
		{"blueGrey700\n", 11},
		{"blueGrey700.", 11},
		{"blueGrey700:", 11},
		{"blueGrey700,", 11},
		{"blueGrey700\"", 11},
		{"blueGrey700(", 11},
		{"blueGrey700)", 11},
		{"blueGrey700[", 11},
		{"blueGrey700]", 11},
		{"blueGrey700// comment", 11},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.BlueGrey700()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestBlueGrey700Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_b"},
		{"b_"},
		{"_bl"},
		{"bl_"},
		{"_blu"},
		{"blu_"},
		{"_blue"},
		{"blue_"},
		{"_blueG"},
		{"blueG_"},
		{"_blueGr"},
		{"blueGr_"},
		{"_blueGre"},
		{"blueGre_"},
		{"_blueGrey"},
		{"blueGrey_"},
		{"_blueGrey7"},
		{"blueGrey7_"},
		{"_blueGrey70"},
		{"blueGrey70_"},
		{"_blueGrey700"},
		{"blueGrey700_"},
		{"9713b"},
		{"b9713"},
		{"9713bl"},
		{"bl9713"},
		{"9713blu"},
		{"blu9713"},
		{"9713blue"},
		{"blue9713"},
		{"9713blueG"},
		{"blueG9713"},
		{"9713blueGr"},
		{"blueGr9713"},
		{"9713blueGre"},
		{"blueGre9713"},
		{"9713blueGrey"},
		{"blueGrey9713"},
		{"9713blueGrey7"},
		{"blueGrey79713"},
		{"9713blueGrey70"},
		{"blueGrey709713"},
		{"9713blueGrey700"},
		{"blueGrey7009713"},
		{"Zb"},
		{"bZ"},
		{"Zbl"},
		{"blZ"},
		{"Zblu"},
		{"bluZ"},
		{"Zblue"},
		{"blueZ"},
		{"ZblueG"},
		{"blueGZ"},
		{"ZblueGr"},
		{"blueGrZ"},
		{"ZblueGre"},
		{"blueGreZ"},
		{"ZblueGrey"},
		{"blueGreyZ"},
		{"ZblueGrey7"},
		{"blueGrey7Z"},
		{"ZblueGrey70"},
		{"blueGrey70Z"},
		{"ZblueGrey700"},
		{"blueGrey700Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.BlueGrey700()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestBlueGrey800Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"blueGrey800", 11},
		{"blueGrey800 ", 11},
		{"blueGrey800\n", 11},
		{"blueGrey800.", 11},
		{"blueGrey800:", 11},
		{"blueGrey800,", 11},
		{"blueGrey800\"", 11},
		{"blueGrey800(", 11},
		{"blueGrey800)", 11},
		{"blueGrey800[", 11},
		{"blueGrey800]", 11},
		{"blueGrey800// comment", 11},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.BlueGrey800()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestBlueGrey800Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_b"},
		{"b_"},
		{"_bl"},
		{"bl_"},
		{"_blu"},
		{"blu_"},
		{"_blue"},
		{"blue_"},
		{"_blueG"},
		{"blueG_"},
		{"_blueGr"},
		{"blueGr_"},
		{"_blueGre"},
		{"blueGre_"},
		{"_blueGrey"},
		{"blueGrey_"},
		{"_blueGrey8"},
		{"blueGrey8_"},
		{"_blueGrey80"},
		{"blueGrey80_"},
		{"_blueGrey800"},
		{"blueGrey800_"},
		{"9713b"},
		{"b9713"},
		{"9713bl"},
		{"bl9713"},
		{"9713blu"},
		{"blu9713"},
		{"9713blue"},
		{"blue9713"},
		{"9713blueG"},
		{"blueG9713"},
		{"9713blueGr"},
		{"blueGr9713"},
		{"9713blueGre"},
		{"blueGre9713"},
		{"9713blueGrey"},
		{"blueGrey9713"},
		{"9713blueGrey8"},
		{"blueGrey89713"},
		{"9713blueGrey80"},
		{"blueGrey809713"},
		{"9713blueGrey800"},
		{"blueGrey8009713"},
		{"Zb"},
		{"bZ"},
		{"Zbl"},
		{"blZ"},
		{"Zblu"},
		{"bluZ"},
		{"Zblue"},
		{"blueZ"},
		{"ZblueG"},
		{"blueGZ"},
		{"ZblueGr"},
		{"blueGrZ"},
		{"ZblueGre"},
		{"blueGreZ"},
		{"ZblueGrey"},
		{"blueGreyZ"},
		{"ZblueGrey8"},
		{"blueGrey8Z"},
		{"ZblueGrey80"},
		{"blueGrey80Z"},
		{"ZblueGrey800"},
		{"blueGrey800Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.BlueGrey800()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestBlueGrey900Valid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"blueGrey900", 11},
		{"blueGrey900 ", 11},
		{"blueGrey900\n", 11},
		{"blueGrey900.", 11},
		{"blueGrey900:", 11},
		{"blueGrey900,", 11},
		{"blueGrey900\"", 11},
		{"blueGrey900(", 11},
		{"blueGrey900)", 11},
		{"blueGrey900[", 11},
		{"blueGrey900]", 11},
		{"blueGrey900// comment", 11},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.BlueGrey900()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestBlueGrey900Invalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_b"},
		{"b_"},
		{"_bl"},
		{"bl_"},
		{"_blu"},
		{"blu_"},
		{"_blue"},
		{"blue_"},
		{"_blueG"},
		{"blueG_"},
		{"_blueGr"},
		{"blueGr_"},
		{"_blueGre"},
		{"blueGre_"},
		{"_blueGrey"},
		{"blueGrey_"},
		{"_blueGrey9"},
		{"blueGrey9_"},
		{"_blueGrey90"},
		{"blueGrey90_"},
		{"_blueGrey900"},
		{"blueGrey900_"},
		{"9713b"},
		{"b9713"},
		{"9713bl"},
		{"bl9713"},
		{"9713blu"},
		{"blu9713"},
		{"9713blue"},
		{"blue9713"},
		{"9713blueG"},
		{"blueG9713"},
		{"9713blueGr"},
		{"blueGr9713"},
		{"9713blueGre"},
		{"blueGre9713"},
		{"9713blueGrey"},
		{"blueGrey9713"},
		{"9713blueGrey9"},
		{"blueGrey99713"},
		{"9713blueGrey90"},
		{"blueGrey909713"},
		{"9713blueGrey900"},
		{"blueGrey9009713"},
		{"Zb"},
		{"bZ"},
		{"Zbl"},
		{"blZ"},
		{"Zblu"},
		{"bluZ"},
		{"Zblue"},
		{"blueZ"},
		{"ZblueG"},
		{"blueGZ"},
		{"ZblueGr"},
		{"blueGrZ"},
		{"ZblueGre"},
		{"blueGreZ"},
		{"ZblueGrey"},
		{"blueGreyZ"},
		{"ZblueGrey9"},
		{"blueGrey9Z"},
		{"ZblueGrey90"},
		{"blueGrey90Z"},
		{"ZblueGrey900"},
		{"blueGrey900Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.BlueGrey900()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestSclValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"scl", 3},
		{"scl ", 3},
		{"scl\n", 3},
		{"scl.", 3},
		{"scl:", 3},
		{"scl,", 3},
		{"scl\"", 3},
		{"scl(", 3},
		{"scl)", 3},
		{"scl[", 3},
		{"scl]", 3},
		{"scl// comment", 3},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Scl()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestSclInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_s"},
		{"s_"},
		{"_sc"},
		{"sc_"},
		{"_scl"},
		{"scl_"},
		{"9713s"},
		{"s9713"},
		{"9713sc"},
		{"sc9713"},
		{"9713scl"},
		{"scl9713"},
		{"Zs"},
		{"sZ"},
		{"Zsc"},
		{"scZ"},
		{"Zscl"},
		{"sclZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Scl()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestStkWidthValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"stkWidth", 8},
		{"stkWidth ", 8},
		{"stkWidth\n", 8},
		{"stkWidth.", 8},
		{"stkWidth:", 8},
		{"stkWidth,", 8},
		{"stkWidth\"", 8},
		{"stkWidth(", 8},
		{"stkWidth)", 8},
		{"stkWidth[", 8},
		{"stkWidth]", 8},
		{"stkWidth// comment", 8},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.StkWidth()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestStkWidthInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_s"},
		{"s_"},
		{"_st"},
		{"st_"},
		{"_stk"},
		{"stk_"},
		{"_stkW"},
		{"stkW_"},
		{"_stkWi"},
		{"stkWi_"},
		{"_stkWid"},
		{"stkWid_"},
		{"_stkWidt"},
		{"stkWidt_"},
		{"_stkWidth"},
		{"stkWidth_"},
		{"9713s"},
		{"s9713"},
		{"9713st"},
		{"st9713"},
		{"9713stk"},
		{"stk9713"},
		{"9713stkW"},
		{"stkW9713"},
		{"9713stkWi"},
		{"stkWi9713"},
		{"9713stkWid"},
		{"stkWid9713"},
		{"9713stkWidt"},
		{"stkWidt9713"},
		{"9713stkWidth"},
		{"stkWidth9713"},
		{"Zs"},
		{"sZ"},
		{"Zst"},
		{"stZ"},
		{"Zstk"},
		{"stkZ"},
		{"ZstkW"},
		{"stkWZ"},
		{"ZstkWi"},
		{"stkWiZ"},
		{"ZstkWid"},
		{"stkWidZ"},
		{"ZstkWidt"},
		{"stkWidtZ"},
		{"ZstkWidth"},
		{"stkWidthZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.StkWidth()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestShpRadiusValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"shpRadius", 9},
		{"shpRadius ", 9},
		{"shpRadius\n", 9},
		{"shpRadius.", 9},
		{"shpRadius:", 9},
		{"shpRadius,", 9},
		{"shpRadius\"", 9},
		{"shpRadius(", 9},
		{"shpRadius)", 9},
		{"shpRadius[", 9},
		{"shpRadius]", 9},
		{"shpRadius// comment", 9},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.ShpRadius()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestShpRadiusInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_s"},
		{"s_"},
		{"_sh"},
		{"sh_"},
		{"_shp"},
		{"shp_"},
		{"_shpR"},
		{"shpR_"},
		{"_shpRa"},
		{"shpRa_"},
		{"_shpRad"},
		{"shpRad_"},
		{"_shpRadi"},
		{"shpRadi_"},
		{"_shpRadiu"},
		{"shpRadiu_"},
		{"_shpRadius"},
		{"shpRadius_"},
		{"9713s"},
		{"s9713"},
		{"9713sh"},
		{"sh9713"},
		{"9713shp"},
		{"shp9713"},
		{"9713shpR"},
		{"shpR9713"},
		{"9713shpRa"},
		{"shpRa9713"},
		{"9713shpRad"},
		{"shpRad9713"},
		{"9713shpRadi"},
		{"shpRadi9713"},
		{"9713shpRadiu"},
		{"shpRadiu9713"},
		{"9713shpRadius"},
		{"shpRadius9713"},
		{"Zs"},
		{"sZ"},
		{"Zsh"},
		{"shZ"},
		{"Zshp"},
		{"shpZ"},
		{"ZshpR"},
		{"shpRZ"},
		{"ZshpRa"},
		{"shpRaZ"},
		{"ZshpRad"},
		{"shpRadZ"},
		{"ZshpRadi"},
		{"shpRadiZ"},
		{"ZshpRadiu"},
		{"shpRadiuZ"},
		{"ZshpRadius"},
		{"shpRadiusZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.ShpRadius()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestAxisPadValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"axisPad", 7},
		{"axisPad ", 7},
		{"axisPad\n", 7},
		{"axisPad.", 7},
		{"axisPad:", 7},
		{"axisPad,", 7},
		{"axisPad\"", 7},
		{"axisPad(", 7},
		{"axisPad)", 7},
		{"axisPad[", 7},
		{"axisPad]", 7},
		{"axisPad// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.AxisPad()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestAxisPadInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_a"},
		{"a_"},
		{"_ax"},
		{"ax_"},
		{"_axi"},
		{"axi_"},
		{"_axis"},
		{"axis_"},
		{"_axisP"},
		{"axisP_"},
		{"_axisPa"},
		{"axisPa_"},
		{"_axisPad"},
		{"axisPad_"},
		{"9713a"},
		{"a9713"},
		{"9713ax"},
		{"ax9713"},
		{"9713axi"},
		{"axi9713"},
		{"9713axis"},
		{"axis9713"},
		{"9713axisP"},
		{"axisP9713"},
		{"9713axisPa"},
		{"axisPa9713"},
		{"9713axisPad"},
		{"axisPad9713"},
		{"Za"},
		{"aZ"},
		{"Zax"},
		{"axZ"},
		{"Zaxi"},
		{"axiZ"},
		{"Zaxis"},
		{"axisZ"},
		{"ZaxisP"},
		{"axisPZ"},
		{"ZaxisPa"},
		{"axisPaZ"},
		{"ZaxisPad"},
		{"axisPadZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.AxisPad()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestBarPadValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"barPad", 6},
		{"barPad ", 6},
		{"barPad\n", 6},
		{"barPad.", 6},
		{"barPad:", 6},
		{"barPad,", 6},
		{"barPad\"", 6},
		{"barPad(", 6},
		{"barPad)", 6},
		{"barPad[", 6},
		{"barPad]", 6},
		{"barPad// comment", 6},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.BarPad()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestBarPadInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_b"},
		{"b_"},
		{"_ba"},
		{"ba_"},
		{"_bar"},
		{"bar_"},
		{"_barP"},
		{"barP_"},
		{"_barPa"},
		{"barPa_"},
		{"_barPad"},
		{"barPad_"},
		{"9713b"},
		{"b9713"},
		{"9713ba"},
		{"ba9713"},
		{"9713bar"},
		{"bar9713"},
		{"9713barP"},
		{"barP9713"},
		{"9713barPa"},
		{"barPa9713"},
		{"9713barPad"},
		{"barPad9713"},
		{"Zb"},
		{"bZ"},
		{"Zba"},
		{"baZ"},
		{"Zbar"},
		{"barZ"},
		{"ZbarP"},
		{"barPZ"},
		{"ZbarPa"},
		{"barPaZ"},
		{"ZbarPad"},
		{"barPadZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.BarPad()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestLenValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"len", 3},
		{"len ", 3},
		{"len\n", 3},
		{"len.", 3},
		{"len:", 3},
		{"len,", 3},
		{"len\"", 3},
		{"len(", 3},
		{"len)", 3},
		{"len[", 3},
		{"len]", 3},
		{"len// comment", 3},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Len()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestLenInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_l"},
		{"l_"},
		{"_le"},
		{"le_"},
		{"_len"},
		{"len_"},
		{"9713l"},
		{"l9713"},
		{"9713le"},
		{"le9713"},
		{"9713len"},
		{"len9713"},
		{"Zl"},
		{"lZ"},
		{"Zle"},
		{"leZ"},
		{"Zlen"},
		{"lenZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Len()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestPadValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"pad", 3},
		{"pad ", 3},
		{"pad\n", 3},
		{"pad.", 3},
		{"pad:", 3},
		{"pad,", 3},
		{"pad\"", 3},
		{"pad(", 3},
		{"pad)", 3},
		{"pad[", 3},
		{"pad]", 3},
		{"pad// comment", 3},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Pad()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestPadInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_p"},
		{"p_"},
		{"_pa"},
		{"pa_"},
		{"_pad"},
		{"pad_"},
		{"9713p"},
		{"p9713"},
		{"9713pa"},
		{"pa9713"},
		{"9713pad"},
		{"pad9713"},
		{"Zp"},
		{"pZ"},
		{"Zpa"},
		{"paZ"},
		{"Zpad"},
		{"padZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Pad()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestBakClrValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"bakClr", 6},
		{"bakClr ", 6},
		{"bakClr\n", 6},
		{"bakClr.", 6},
		{"bakClr:", 6},
		{"bakClr,", 6},
		{"bakClr\"", 6},
		{"bakClr(", 6},
		{"bakClr)", 6},
		{"bakClr[", 6},
		{"bakClr]", 6},
		{"bakClr// comment", 6},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.BakClr()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestBakClrInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_b"},
		{"b_"},
		{"_ba"},
		{"ba_"},
		{"_bak"},
		{"bak_"},
		{"_bakC"},
		{"bakC_"},
		{"_bakCl"},
		{"bakCl_"},
		{"_bakClr"},
		{"bakClr_"},
		{"9713b"},
		{"b9713"},
		{"9713ba"},
		{"ba9713"},
		{"9713bak"},
		{"bak9713"},
		{"9713bakC"},
		{"bakC9713"},
		{"9713bakCl"},
		{"bakCl9713"},
		{"9713bakClr"},
		{"bakClr9713"},
		{"Zb"},
		{"bZ"},
		{"Zba"},
		{"baZ"},
		{"Zbak"},
		{"bakZ"},
		{"ZbakC"},
		{"bakCZ"},
		{"ZbakCl"},
		{"bakClZ"},
		{"ZbakClr"},
		{"bakClrZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.BakClr()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestBrdrClrValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"brdrClr", 7},
		{"brdrClr ", 7},
		{"brdrClr\n", 7},
		{"brdrClr.", 7},
		{"brdrClr:", 7},
		{"brdrClr,", 7},
		{"brdrClr\"", 7},
		{"brdrClr(", 7},
		{"brdrClr)", 7},
		{"brdrClr[", 7},
		{"brdrClr]", 7},
		{"brdrClr// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.BrdrClr()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestBrdrClrInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_b"},
		{"b_"},
		{"_br"},
		{"br_"},
		{"_brd"},
		{"brd_"},
		{"_brdr"},
		{"brdr_"},
		{"_brdrC"},
		{"brdrC_"},
		{"_brdrCl"},
		{"brdrCl_"},
		{"_brdrClr"},
		{"brdrClr_"},
		{"9713b"},
		{"b9713"},
		{"9713br"},
		{"br9713"},
		{"9713brd"},
		{"brd9713"},
		{"9713brdr"},
		{"brdr9713"},
		{"9713brdrC"},
		{"brdrC9713"},
		{"9713brdrCl"},
		{"brdrCl9713"},
		{"9713brdrClr"},
		{"brdrClr9713"},
		{"Zb"},
		{"bZ"},
		{"Zbr"},
		{"brZ"},
		{"Zbrd"},
		{"brdZ"},
		{"Zbrdr"},
		{"brdrZ"},
		{"ZbrdrC"},
		{"brdrCZ"},
		{"ZbrdrCl"},
		{"brdrClZ"},
		{"ZbrdrClr"},
		{"brdrClrZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.BrdrClr()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestBrdrLenValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"brdrLen", 7},
		{"brdrLen ", 7},
		{"brdrLen\n", 7},
		{"brdrLen.", 7},
		{"brdrLen:", 7},
		{"brdrLen,", 7},
		{"brdrLen\"", 7},
		{"brdrLen(", 7},
		{"brdrLen)", 7},
		{"brdrLen[", 7},
		{"brdrLen]", 7},
		{"brdrLen// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.BrdrLen()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestBrdrLenInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_b"},
		{"b_"},
		{"_br"},
		{"br_"},
		{"_brd"},
		{"brd_"},
		{"_brdr"},
		{"brdr_"},
		{"_brdrL"},
		{"brdrL_"},
		{"_brdrLe"},
		{"brdrLe_"},
		{"_brdrLen"},
		{"brdrLen_"},
		{"9713b"},
		{"b9713"},
		{"9713br"},
		{"br9713"},
		{"9713brd"},
		{"brd9713"},
		{"9713brdr"},
		{"brdr9713"},
		{"9713brdrL"},
		{"brdrL9713"},
		{"9713brdrLe"},
		{"brdrLe9713"},
		{"9713brdrLen"},
		{"brdrLen9713"},
		{"Zb"},
		{"bZ"},
		{"Zbr"},
		{"brZ"},
		{"Zbrd"},
		{"brdZ"},
		{"Zbrdr"},
		{"brdrZ"},
		{"ZbrdrL"},
		{"brdrLZ"},
		{"ZbrdrLe"},
		{"brdrLeZ"},
		{"ZbrdrLen"},
		{"brdrLenZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.BrdrLen()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestInrvlTxtLenValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"inrvlTxtLen", 11},
		{"inrvlTxtLen ", 11},
		{"inrvlTxtLen\n", 11},
		{"inrvlTxtLen.", 11},
		{"inrvlTxtLen:", 11},
		{"inrvlTxtLen,", 11},
		{"inrvlTxtLen\"", 11},
		{"inrvlTxtLen(", 11},
		{"inrvlTxtLen)", 11},
		{"inrvlTxtLen[", 11},
		{"inrvlTxtLen]", 11},
		{"inrvlTxtLen// comment", 11},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.InrvlTxtLen()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestInrvlTxtLenInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_i"},
		{"i_"},
		{"_in"},
		{"in_"},
		{"_inr"},
		{"inr_"},
		{"_inrv"},
		{"inrv_"},
		{"_inrvl"},
		{"inrvl_"},
		{"_inrvlT"},
		{"inrvlT_"},
		{"_inrvlTx"},
		{"inrvlTx_"},
		{"_inrvlTxt"},
		{"inrvlTxt_"},
		{"_inrvlTxtL"},
		{"inrvlTxtL_"},
		{"_inrvlTxtLe"},
		{"inrvlTxtLe_"},
		{"_inrvlTxtLen"},
		{"inrvlTxtLen_"},
		{"9713i"},
		{"i9713"},
		{"9713in"},
		{"in9713"},
		{"9713inr"},
		{"inr9713"},
		{"9713inrv"},
		{"inrv9713"},
		{"9713inrvl"},
		{"inrvl9713"},
		{"9713inrvlT"},
		{"inrvlT9713"},
		{"9713inrvlTx"},
		{"inrvlTx9713"},
		{"9713inrvlTxt"},
		{"inrvlTxt9713"},
		{"9713inrvlTxtL"},
		{"inrvlTxtL9713"},
		{"9713inrvlTxtLe"},
		{"inrvlTxtLe9713"},
		{"9713inrvlTxtLen"},
		{"inrvlTxtLen9713"},
		{"Zi"},
		{"iZ"},
		{"Zin"},
		{"inZ"},
		{"Zinr"},
		{"inrZ"},
		{"Zinrv"},
		{"inrvZ"},
		{"Zinrvl"},
		{"inrvlZ"},
		{"ZinrvlT"},
		{"inrvlTZ"},
		{"ZinrvlTx"},
		{"inrvlTxZ"},
		{"ZinrvlTxt"},
		{"inrvlTxtZ"},
		{"ZinrvlTxtL"},
		{"inrvlTxtLZ"},
		{"ZinrvlTxtLe"},
		{"inrvlTxtLeZ"},
		{"ZinrvlTxtLen"},
		{"inrvlTxtLenZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.InrvlTxtLen()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestInrvlTxtClrXValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"inrvlTxtClrX", 12},
		{"inrvlTxtClrX ", 12},
		{"inrvlTxtClrX\n", 12},
		{"inrvlTxtClrX.", 12},
		{"inrvlTxtClrX:", 12},
		{"inrvlTxtClrX,", 12},
		{"inrvlTxtClrX\"", 12},
		{"inrvlTxtClrX(", 12},
		{"inrvlTxtClrX)", 12},
		{"inrvlTxtClrX[", 12},
		{"inrvlTxtClrX]", 12},
		{"inrvlTxtClrX// comment", 12},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.InrvlTxtClrX()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestInrvlTxtClrXInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_i"},
		{"i_"},
		{"_in"},
		{"in_"},
		{"_inr"},
		{"inr_"},
		{"_inrv"},
		{"inrv_"},
		{"_inrvl"},
		{"inrvl_"},
		{"_inrvlT"},
		{"inrvlT_"},
		{"_inrvlTx"},
		{"inrvlTx_"},
		{"_inrvlTxt"},
		{"inrvlTxt_"},
		{"_inrvlTxtC"},
		{"inrvlTxtC_"},
		{"_inrvlTxtCl"},
		{"inrvlTxtCl_"},
		{"_inrvlTxtClr"},
		{"inrvlTxtClr_"},
		{"_inrvlTxtClrX"},
		{"inrvlTxtClrX_"},
		{"9713i"},
		{"i9713"},
		{"9713in"},
		{"in9713"},
		{"9713inr"},
		{"inr9713"},
		{"9713inrv"},
		{"inrv9713"},
		{"9713inrvl"},
		{"inrvl9713"},
		{"9713inrvlT"},
		{"inrvlT9713"},
		{"9713inrvlTx"},
		{"inrvlTx9713"},
		{"9713inrvlTxt"},
		{"inrvlTxt9713"},
		{"9713inrvlTxtC"},
		{"inrvlTxtC9713"},
		{"9713inrvlTxtCl"},
		{"inrvlTxtCl9713"},
		{"9713inrvlTxtClr"},
		{"inrvlTxtClr9713"},
		{"9713inrvlTxtClrX"},
		{"inrvlTxtClrX9713"},
		{"Zi"},
		{"iZ"},
		{"Zin"},
		{"inZ"},
		{"Zinr"},
		{"inrZ"},
		{"Zinrv"},
		{"inrvZ"},
		{"Zinrvl"},
		{"inrvlZ"},
		{"ZinrvlT"},
		{"inrvlTZ"},
		{"ZinrvlTx"},
		{"inrvlTxZ"},
		{"ZinrvlTxt"},
		{"inrvlTxtZ"},
		{"ZinrvlTxtC"},
		{"inrvlTxtCZ"},
		{"ZinrvlTxtCl"},
		{"inrvlTxtClZ"},
		{"ZinrvlTxtClr"},
		{"inrvlTxtClrZ"},
		{"ZinrvlTxtClrX"},
		{"inrvlTxtClrXZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.InrvlTxtClrX()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestInrvlTxtClrYValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"inrvlTxtClrY", 12},
		{"inrvlTxtClrY ", 12},
		{"inrvlTxtClrY\n", 12},
		{"inrvlTxtClrY.", 12},
		{"inrvlTxtClrY:", 12},
		{"inrvlTxtClrY,", 12},
		{"inrvlTxtClrY\"", 12},
		{"inrvlTxtClrY(", 12},
		{"inrvlTxtClrY)", 12},
		{"inrvlTxtClrY[", 12},
		{"inrvlTxtClrY]", 12},
		{"inrvlTxtClrY// comment", 12},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.InrvlTxtClrY()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestInrvlTxtClrYInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_i"},
		{"i_"},
		{"_in"},
		{"in_"},
		{"_inr"},
		{"inr_"},
		{"_inrv"},
		{"inrv_"},
		{"_inrvl"},
		{"inrvl_"},
		{"_inrvlT"},
		{"inrvlT_"},
		{"_inrvlTx"},
		{"inrvlTx_"},
		{"_inrvlTxt"},
		{"inrvlTxt_"},
		{"_inrvlTxtC"},
		{"inrvlTxtC_"},
		{"_inrvlTxtCl"},
		{"inrvlTxtCl_"},
		{"_inrvlTxtClr"},
		{"inrvlTxtClr_"},
		{"_inrvlTxtClrY"},
		{"inrvlTxtClrY_"},
		{"9713i"},
		{"i9713"},
		{"9713in"},
		{"in9713"},
		{"9713inr"},
		{"inr9713"},
		{"9713inrv"},
		{"inrv9713"},
		{"9713inrvl"},
		{"inrvl9713"},
		{"9713inrvlT"},
		{"inrvlT9713"},
		{"9713inrvlTx"},
		{"inrvlTx9713"},
		{"9713inrvlTxt"},
		{"inrvlTxt9713"},
		{"9713inrvlTxtC"},
		{"inrvlTxtC9713"},
		{"9713inrvlTxtCl"},
		{"inrvlTxtCl9713"},
		{"9713inrvlTxtClr"},
		{"inrvlTxtClr9713"},
		{"9713inrvlTxtClrY"},
		{"inrvlTxtClrY9713"},
		{"Zi"},
		{"iZ"},
		{"Zin"},
		{"inZ"},
		{"Zinr"},
		{"inrZ"},
		{"Zinrv"},
		{"inrvZ"},
		{"Zinrvl"},
		{"inrvlZ"},
		{"ZinrvlT"},
		{"inrvlTZ"},
		{"ZinrvlTx"},
		{"inrvlTxZ"},
		{"ZinrvlTxt"},
		{"inrvlTxtZ"},
		{"ZinrvlTxtC"},
		{"inrvlTxtCZ"},
		{"ZinrvlTxtCl"},
		{"inrvlTxtClZ"},
		{"ZinrvlTxtClr"},
		{"inrvlTxtClrZ"},
		{"ZinrvlTxtClrY"},
		{"inrvlTxtClrYZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.InrvlTxtClrY()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestMsgClrValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"msgClr", 6},
		{"msgClr ", 6},
		{"msgClr\n", 6},
		{"msgClr.", 6},
		{"msgClr:", 6},
		{"msgClr,", 6},
		{"msgClr\"", 6},
		{"msgClr(", 6},
		{"msgClr)", 6},
		{"msgClr[", 6},
		{"msgClr]", 6},
		{"msgClr// comment", 6},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.MsgClr()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestMsgClrInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_m"},
		{"m_"},
		{"_ms"},
		{"ms_"},
		{"_msg"},
		{"msg_"},
		{"_msgC"},
		{"msgC_"},
		{"_msgCl"},
		{"msgCl_"},
		{"_msgClr"},
		{"msgClr_"},
		{"9713m"},
		{"m9713"},
		{"9713ms"},
		{"ms9713"},
		{"9713msg"},
		{"msg9713"},
		{"9713msgC"},
		{"msgC9713"},
		{"9713msgCl"},
		{"msgCl9713"},
		{"9713msgClr"},
		{"msgClr9713"},
		{"Zm"},
		{"mZ"},
		{"Zms"},
		{"msZ"},
		{"Zmsg"},
		{"msgZ"},
		{"ZmsgC"},
		{"msgCZ"},
		{"ZmsgCl"},
		{"msgClZ"},
		{"ZmsgClr"},
		{"msgClrZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.MsgClr()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestTitleClrValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"titleClr", 8},
		{"titleClr ", 8},
		{"titleClr\n", 8},
		{"titleClr.", 8},
		{"titleClr:", 8},
		{"titleClr,", 8},
		{"titleClr\"", 8},
		{"titleClr(", 8},
		{"titleClr)", 8},
		{"titleClr[", 8},
		{"titleClr]", 8},
		{"titleClr// comment", 8},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.TitleClr()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestTitleClrInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_t"},
		{"t_"},
		{"_ti"},
		{"ti_"},
		{"_tit"},
		{"tit_"},
		{"_titl"},
		{"titl_"},
		{"_title"},
		{"title_"},
		{"_titleC"},
		{"titleC_"},
		{"_titleCl"},
		{"titleCl_"},
		{"_titleClr"},
		{"titleClr_"},
		{"9713t"},
		{"t9713"},
		{"9713ti"},
		{"ti9713"},
		{"9713tit"},
		{"tit9713"},
		{"9713titl"},
		{"titl9713"},
		{"9713title"},
		{"title9713"},
		{"9713titleC"},
		{"titleC9713"},
		{"9713titleCl"},
		{"titleCl9713"},
		{"9713titleClr"},
		{"titleClr9713"},
		{"Zt"},
		{"tZ"},
		{"Zti"},
		{"tiZ"},
		{"Ztit"},
		{"titZ"},
		{"Ztitl"},
		{"titlZ"},
		{"Ztitle"},
		{"titleZ"},
		{"ZtitleC"},
		{"titleCZ"},
		{"ZtitleCl"},
		{"titleClZ"},
		{"ZtitleClr"},
		{"titleClrZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.TitleClr()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestPrfClrValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"prfClr", 6},
		{"prfClr ", 6},
		{"prfClr\n", 6},
		{"prfClr.", 6},
		{"prfClr:", 6},
		{"prfClr,", 6},
		{"prfClr\"", 6},
		{"prfClr(", 6},
		{"prfClr)", 6},
		{"prfClr[", 6},
		{"prfClr]", 6},
		{"prfClr// comment", 6},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.PrfClr()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestPrfClrInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_p"},
		{"p_"},
		{"_pr"},
		{"pr_"},
		{"_prf"},
		{"prf_"},
		{"_prfC"},
		{"prfC_"},
		{"_prfCl"},
		{"prfCl_"},
		{"_prfClr"},
		{"prfClr_"},
		{"9713p"},
		{"p9713"},
		{"9713pr"},
		{"pr9713"},
		{"9713prf"},
		{"prf9713"},
		{"9713prfC"},
		{"prfC9713"},
		{"9713prfCl"},
		{"prfCl9713"},
		{"9713prfClr"},
		{"prfClr9713"},
		{"Zp"},
		{"pZ"},
		{"Zpr"},
		{"prZ"},
		{"Zprf"},
		{"prfZ"},
		{"ZprfC"},
		{"prfCZ"},
		{"ZprfCl"},
		{"prfClZ"},
		{"ZprfClr"},
		{"prfClrZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.PrfClr()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestLosClrValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"losClr", 6},
		{"losClr ", 6},
		{"losClr\n", 6},
		{"losClr.", 6},
		{"losClr:", 6},
		{"losClr,", 6},
		{"losClr\"", 6},
		{"losClr(", 6},
		{"losClr)", 6},
		{"losClr[", 6},
		{"losClr]", 6},
		{"losClr// comment", 6},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.LosClr()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestLosClrInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_l"},
		{"l_"},
		{"_lo"},
		{"lo_"},
		{"_los"},
		{"los_"},
		{"_losC"},
		{"losC_"},
		{"_losCl"},
		{"losCl_"},
		{"_losClr"},
		{"losClr_"},
		{"9713l"},
		{"l9713"},
		{"9713lo"},
		{"lo9713"},
		{"9713los"},
		{"los9713"},
		{"9713losC"},
		{"losC9713"},
		{"9713losCl"},
		{"losCl9713"},
		{"9713losClr"},
		{"losClr9713"},
		{"Zl"},
		{"lZ"},
		{"Zlo"},
		{"loZ"},
		{"Zlos"},
		{"losZ"},
		{"ZlosC"},
		{"losCZ"},
		{"ZlosCl"},
		{"losClZ"},
		{"ZlosClr"},
		{"losClrZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.LosClr()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestPrfPenValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"prfPen", 6},
		{"prfPen ", 6},
		{"prfPen\n", 6},
		{"prfPen.", 6},
		{"prfPen:", 6},
		{"prfPen,", 6},
		{"prfPen\"", 6},
		{"prfPen(", 6},
		{"prfPen)", 6},
		{"prfPen[", 6},
		{"prfPen]", 6},
		{"prfPen// comment", 6},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.PrfPen()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestPrfPenInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_p"},
		{"p_"},
		{"_pr"},
		{"pr_"},
		{"_prf"},
		{"prf_"},
		{"_prfP"},
		{"prfP_"},
		{"_prfPe"},
		{"prfPe_"},
		{"_prfPen"},
		{"prfPen_"},
		{"9713p"},
		{"p9713"},
		{"9713pr"},
		{"pr9713"},
		{"9713prf"},
		{"prf9713"},
		{"9713prfP"},
		{"prfP9713"},
		{"9713prfPe"},
		{"prfPe9713"},
		{"9713prfPen"},
		{"prfPen9713"},
		{"Zp"},
		{"pZ"},
		{"Zpr"},
		{"prZ"},
		{"Zprf"},
		{"prfZ"},
		{"ZprfP"},
		{"prfPZ"},
		{"ZprfPe"},
		{"prfPeZ"},
		{"ZprfPen"},
		{"prfPenZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.PrfPen()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestLosPenValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"losPen", 6},
		{"losPen ", 6},
		{"losPen\n", 6},
		{"losPen.", 6},
		{"losPen:", 6},
		{"losPen,", 6},
		{"losPen\"", 6},
		{"losPen(", 6},
		{"losPen)", 6},
		{"losPen[", 6},
		{"losPen]", 6},
		{"losPen// comment", 6},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.LosPen()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestLosPenInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_l"},
		{"l_"},
		{"_lo"},
		{"lo_"},
		{"_los"},
		{"los_"},
		{"_losP"},
		{"losP_"},
		{"_losPe"},
		{"losPe_"},
		{"_losPen"},
		{"losPen_"},
		{"9713l"},
		{"l9713"},
		{"9713lo"},
		{"lo9713"},
		{"9713los"},
		{"los9713"},
		{"9713losP"},
		{"losP9713"},
		{"9713losPe"},
		{"losPe9713"},
		{"9713losPen"},
		{"losPen9713"},
		{"Zl"},
		{"lZ"},
		{"Zlo"},
		{"loZ"},
		{"Zlos"},
		{"losZ"},
		{"ZlosP"},
		{"losPZ"},
		{"ZlosPe"},
		{"losPeZ"},
		{"ZlosPen"},
		{"losPenZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.LosPen()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestOutlierLimValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"outlierLim", 10},
		{"outlierLim ", 10},
		{"outlierLim\n", 10},
		{"outlierLim.", 10},
		{"outlierLim:", 10},
		{"outlierLim,", 10},
		{"outlierLim\"", 10},
		{"outlierLim(", 10},
		{"outlierLim)", 10},
		{"outlierLim[", 10},
		{"outlierLim]", 10},
		{"outlierLim// comment", 10},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.OutlierLim()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestOutlierLimInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_o"},
		{"o_"},
		{"_ou"},
		{"ou_"},
		{"_out"},
		{"out_"},
		{"_outl"},
		{"outl_"},
		{"_outli"},
		{"outli_"},
		{"_outlie"},
		{"outlie_"},
		{"_outlier"},
		{"outlier_"},
		{"_outlierL"},
		{"outlierL_"},
		{"_outlierLi"},
		{"outlierLi_"},
		{"_outlierLim"},
		{"outlierLim_"},
		{"9713o"},
		{"o9713"},
		{"9713ou"},
		{"ou9713"},
		{"9713out"},
		{"out9713"},
		{"9713outl"},
		{"outl9713"},
		{"9713outli"},
		{"outli9713"},
		{"9713outlie"},
		{"outlie9713"},
		{"9713outlier"},
		{"outlier9713"},
		{"9713outlierL"},
		{"outlierL9713"},
		{"9713outlierLi"},
		{"outlierLi9713"},
		{"9713outlierLim"},
		{"outlierLim9713"},
		{"Zo"},
		{"oZ"},
		{"Zou"},
		{"ouZ"},
		{"Zout"},
		{"outZ"},
		{"Zoutl"},
		{"outlZ"},
		{"Zoutli"},
		{"outliZ"},
		{"Zoutlie"},
		{"outlieZ"},
		{"Zoutlier"},
		{"outlierZ"},
		{"ZoutlierL"},
		{"outlierLZ"},
		{"ZoutlierLi"},
		{"outlierLiZ"},
		{"ZoutlierLim"},
		{"outlierLimZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.OutlierLim()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestIfoValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"ifo", 3},
		{"ifo ", 3},
		{"ifo\n", 3},
		{"ifo.", 3},
		{"ifo:", 3},
		{"ifo,", 3},
		{"ifo\"", 3},
		{"ifo(", 3},
		{"ifo)", 3},
		{"ifo[", 3},
		{"ifo]", 3},
		{"ifo// comment", 3},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Ifo()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestIfoInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_i"},
		{"i_"},
		{"_if"},
		{"if_"},
		{"_ifo"},
		{"ifo_"},
		{"9713i"},
		{"i9713"},
		{"9713if"},
		{"if9713"},
		{"9713ifo"},
		{"ifo9713"},
		{"Zi"},
		{"iZ"},
		{"Zif"},
		{"ifZ"},
		{"Zifo"},
		{"ifoZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Ifo()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestIfofValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"ifof", 4},
		{"ifof ", 4},
		{"ifof\n", 4},
		{"ifof.", 4},
		{"ifof:", 4},
		{"ifof,", 4},
		{"ifof\"", 4},
		{"ifof(", 4},
		{"ifof)", 4},
		{"ifof[", 4},
		{"ifof]", 4},
		{"ifof// comment", 4},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Ifof()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestIfofInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_i"},
		{"i_"},
		{"_if"},
		{"if_"},
		{"_ifo"},
		{"ifo_"},
		{"_ifof"},
		{"ifof_"},
		{"9713i"},
		{"i9713"},
		{"9713if"},
		{"if9713"},
		{"9713ifo"},
		{"ifo9713"},
		{"9713ifof"},
		{"ifof9713"},
		{"Zi"},
		{"iZ"},
		{"Zif"},
		{"ifZ"},
		{"Zifo"},
		{"ifoZ"},
		{"Zifof"},
		{"ifofZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Ifof()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestFmtValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"fmt", 3},
		{"fmt ", 3},
		{"fmt\n", 3},
		{"fmt.", 3},
		{"fmt:", 3},
		{"fmt,", 3},
		{"fmt\"", 3},
		{"fmt(", 3},
		{"fmt)", 3},
		{"fmt[", 3},
		{"fmt]", 3},
		{"fmt// comment", 3},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Fmt()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestFmtInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_f"},
		{"f_"},
		{"_fm"},
		{"fm_"},
		{"_fmt"},
		{"fmt_"},
		{"9713f"},
		{"f9713"},
		{"9713fm"},
		{"fm9713"},
		{"9713fmt"},
		{"fmt9713"},
		{"Zf"},
		{"fZ"},
		{"Zfm"},
		{"fmZ"},
		{"Zfmt"},
		{"fmtZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Fmt()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestNowValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"now", 3},
		{"now ", 3},
		{"now\n", 3},
		{"now.", 3},
		{"now:", 3},
		{"now,", 3},
		{"now\"", 3},
		{"now(", 3},
		{"now)", 3},
		{"now[", 3},
		{"now]", 3},
		{"now// comment", 3},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Now()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestNowInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_n"},
		{"n_"},
		{"_no"},
		{"no_"},
		{"_now"},
		{"now_"},
		{"9713n"},
		{"n9713"},
		{"9713no"},
		{"no9713"},
		{"9713now"},
		{"now9713"},
		{"Zn"},
		{"nZ"},
		{"Zno"},
		{"noZ"},
		{"Znow"},
		{"nowZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Now()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestNewRngValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"newRng", 6},
		{"newRng ", 6},
		{"newRng\n", 6},
		{"newRng.", 6},
		{"newRng:", 6},
		{"newRng,", 6},
		{"newRng\"", 6},
		{"newRng(", 6},
		{"newRng)", 6},
		{"newRng[", 6},
		{"newRng]", 6},
		{"newRng// comment", 6},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.NewRng()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestNewRngInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_n"},
		{"n_"},
		{"_ne"},
		{"ne_"},
		{"_new"},
		{"new_"},
		{"_newR"},
		{"newR_"},
		{"_newRn"},
		{"newRn_"},
		{"_newRng"},
		{"newRng_"},
		{"9713n"},
		{"n9713"},
		{"9713ne"},
		{"ne9713"},
		{"9713new"},
		{"new9713"},
		{"9713newR"},
		{"newR9713"},
		{"9713newRn"},
		{"newRn9713"},
		{"9713newRng"},
		{"newRng9713"},
		{"Zn"},
		{"nZ"},
		{"Zne"},
		{"neZ"},
		{"Znew"},
		{"newZ"},
		{"ZnewR"},
		{"newRZ"},
		{"ZnewRn"},
		{"newRnZ"},
		{"ZnewRng"},
		{"newRngZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.NewRng()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestNewRngArndValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"newRngArnd", 10},
		{"newRngArnd ", 10},
		{"newRngArnd\n", 10},
		{"newRngArnd.", 10},
		{"newRngArnd:", 10},
		{"newRngArnd,", 10},
		{"newRngArnd\"", 10},
		{"newRngArnd(", 10},
		{"newRngArnd)", 10},
		{"newRngArnd[", 10},
		{"newRngArnd]", 10},
		{"newRngArnd// comment", 10},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.NewRngArnd()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestNewRngArndInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_n"},
		{"n_"},
		{"_ne"},
		{"ne_"},
		{"_new"},
		{"new_"},
		{"_newR"},
		{"newR_"},
		{"_newRn"},
		{"newRn_"},
		{"_newRng"},
		{"newRng_"},
		{"_newRngA"},
		{"newRngA_"},
		{"_newRngAr"},
		{"newRngAr_"},
		{"_newRngArn"},
		{"newRngArn_"},
		{"_newRngArnd"},
		{"newRngArnd_"},
		{"9713n"},
		{"n9713"},
		{"9713ne"},
		{"ne9713"},
		{"9713new"},
		{"new9713"},
		{"9713newR"},
		{"newR9713"},
		{"9713newRn"},
		{"newRn9713"},
		{"9713newRng"},
		{"newRng9713"},
		{"9713newRngA"},
		{"newRngA9713"},
		{"9713newRngAr"},
		{"newRngAr9713"},
		{"9713newRngArn"},
		{"newRngArn9713"},
		{"9713newRngArnd"},
		{"newRngArnd9713"},
		{"Zn"},
		{"nZ"},
		{"Zne"},
		{"neZ"},
		{"Znew"},
		{"newZ"},
		{"ZnewR"},
		{"newRZ"},
		{"ZnewRn"},
		{"newRnZ"},
		{"ZnewRng"},
		{"newRngZ"},
		{"ZnewRngA"},
		{"newRngAZ"},
		{"ZnewRngAr"},
		{"newRngArZ"},
		{"ZnewRngArn"},
		{"newRngArnZ"},
		{"ZnewRngArnd"},
		{"newRngArndZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.NewRngArnd()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestNewRngFulValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"newRngFul", 9},
		{"newRngFul ", 9},
		{"newRngFul\n", 9},
		{"newRngFul.", 9},
		{"newRngFul:", 9},
		{"newRngFul,", 9},
		{"newRngFul\"", 9},
		{"newRngFul(", 9},
		{"newRngFul)", 9},
		{"newRngFul[", 9},
		{"newRngFul]", 9},
		{"newRngFul// comment", 9},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.NewRngFul()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestNewRngFulInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_n"},
		{"n_"},
		{"_ne"},
		{"ne_"},
		{"_new"},
		{"new_"},
		{"_newR"},
		{"newR_"},
		{"_newRn"},
		{"newRn_"},
		{"_newRng"},
		{"newRng_"},
		{"_newRngF"},
		{"newRngF_"},
		{"_newRngFu"},
		{"newRngFu_"},
		{"_newRngFul"},
		{"newRngFul_"},
		{"9713n"},
		{"n9713"},
		{"9713ne"},
		{"ne9713"},
		{"9713new"},
		{"new9713"},
		{"9713newR"},
		{"newR9713"},
		{"9713newRn"},
		{"newRn9713"},
		{"9713newRng"},
		{"newRng9713"},
		{"9713newRngF"},
		{"newRngF9713"},
		{"9713newRngFu"},
		{"newRngFu9713"},
		{"9713newRngFul"},
		{"newRngFul9713"},
		{"Zn"},
		{"nZ"},
		{"Zne"},
		{"neZ"},
		{"Znew"},
		{"newZ"},
		{"ZnewR"},
		{"newRZ"},
		{"ZnewRn"},
		{"newRnZ"},
		{"ZnewRng"},
		{"newRngZ"},
		{"ZnewRngF"},
		{"newRngFZ"},
		{"ZnewRngFu"},
		{"newRngFuZ"},
		{"ZnewRngFul"},
		{"newRngFulZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.NewRngFul()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestNewValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"new", 3},
		{"new ", 3},
		{"new\n", 3},
		{"new.", 3},
		{"new:", 3},
		{"new,", 3},
		{"new\"", 3},
		{"new(", 3},
		{"new)", 3},
		{"new[", 3},
		{"new]", 3},
		{"new// comment", 3},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.New()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestNewInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_n"},
		{"n_"},
		{"_ne"},
		{"ne_"},
		{"_new"},
		{"new_"},
		{"9713n"},
		{"n9713"},
		{"9713ne"},
		{"ne9713"},
		{"9713new"},
		{"new9713"},
		{"Zn"},
		{"nZ"},
		{"Zne"},
		{"neZ"},
		{"Znew"},
		{"newZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.New()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestMakeValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"make", 4},
		{"make ", 4},
		{"make\n", 4},
		{"make.", 4},
		{"make:", 4},
		{"make,", 4},
		{"make\"", 4},
		{"make(", 4},
		{"make)", 4},
		{"make[", 4},
		{"make]", 4},
		{"make// comment", 4},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Make()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestMakeInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_m"},
		{"m_"},
		{"_ma"},
		{"ma_"},
		{"_mak"},
		{"mak_"},
		{"_make"},
		{"make_"},
		{"9713m"},
		{"m9713"},
		{"9713ma"},
		{"ma9713"},
		{"9713mak"},
		{"mak9713"},
		{"9713make"},
		{"make9713"},
		{"Zm"},
		{"mZ"},
		{"Zma"},
		{"maZ"},
		{"Zmak"},
		{"makZ"},
		{"Zmake"},
		{"makeZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Make()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestMakeEmpValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"makeEmp", 7},
		{"makeEmp ", 7},
		{"makeEmp\n", 7},
		{"makeEmp.", 7},
		{"makeEmp:", 7},
		{"makeEmp,", 7},
		{"makeEmp\"", 7},
		{"makeEmp(", 7},
		{"makeEmp)", 7},
		{"makeEmp[", 7},
		{"makeEmp]", 7},
		{"makeEmp// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.MakeEmp()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestMakeEmpInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_m"},
		{"m_"},
		{"_ma"},
		{"ma_"},
		{"_mak"},
		{"mak_"},
		{"_make"},
		{"make_"},
		{"_makeE"},
		{"makeE_"},
		{"_makeEm"},
		{"makeEm_"},
		{"_makeEmp"},
		{"makeEmp_"},
		{"9713m"},
		{"m9713"},
		{"9713ma"},
		{"ma9713"},
		{"9713mak"},
		{"mak9713"},
		{"9713make"},
		{"make9713"},
		{"9713makeE"},
		{"makeE9713"},
		{"9713makeEm"},
		{"makeEm9713"},
		{"9713makeEmp"},
		{"makeEmp9713"},
		{"Zm"},
		{"mZ"},
		{"Zma"},
		{"maZ"},
		{"Zmak"},
		{"makZ"},
		{"Zmake"},
		{"makeZ"},
		{"ZmakeE"},
		{"makeEZ"},
		{"ZmakeEm"},
		{"makeEmZ"},
		{"ZmakeEmp"},
		{"makeEmpZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.MakeEmp()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestAddsLssValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"addsLss", 7},
		{"addsLss ", 7},
		{"addsLss\n", 7},
		{"addsLss.", 7},
		{"addsLss:", 7},
		{"addsLss,", 7},
		{"addsLss\"", 7},
		{"addsLss(", 7},
		{"addsLss)", 7},
		{"addsLss[", 7},
		{"addsLss]", 7},
		{"addsLss// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.AddsLss()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestAddsLssInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_a"},
		{"a_"},
		{"_ad"},
		{"ad_"},
		{"_add"},
		{"add_"},
		{"_adds"},
		{"adds_"},
		{"_addsL"},
		{"addsL_"},
		{"_addsLs"},
		{"addsLs_"},
		{"_addsLss"},
		{"addsLss_"},
		{"9713a"},
		{"a9713"},
		{"9713ad"},
		{"ad9713"},
		{"9713add"},
		{"add9713"},
		{"9713adds"},
		{"adds9713"},
		{"9713addsL"},
		{"addsL9713"},
		{"9713addsLs"},
		{"addsLs9713"},
		{"9713addsLss"},
		{"addsLss9713"},
		{"Za"},
		{"aZ"},
		{"Zad"},
		{"adZ"},
		{"Zadd"},
		{"addZ"},
		{"Zadds"},
		{"addsZ"},
		{"ZaddsL"},
		{"addsLZ"},
		{"ZaddsLs"},
		{"addsLsZ"},
		{"ZaddsLss"},
		{"addsLssZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.AddsLss()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestAddsLeqValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"addsLeq", 7},
		{"addsLeq ", 7},
		{"addsLeq\n", 7},
		{"addsLeq.", 7},
		{"addsLeq:", 7},
		{"addsLeq,", 7},
		{"addsLeq\"", 7},
		{"addsLeq(", 7},
		{"addsLeq)", 7},
		{"addsLeq[", 7},
		{"addsLeq]", 7},
		{"addsLeq// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.AddsLeq()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestAddsLeqInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_a"},
		{"a_"},
		{"_ad"},
		{"ad_"},
		{"_add"},
		{"add_"},
		{"_adds"},
		{"adds_"},
		{"_addsL"},
		{"addsL_"},
		{"_addsLe"},
		{"addsLe_"},
		{"_addsLeq"},
		{"addsLeq_"},
		{"9713a"},
		{"a9713"},
		{"9713ad"},
		{"ad9713"},
		{"9713add"},
		{"add9713"},
		{"9713adds"},
		{"adds9713"},
		{"9713addsL"},
		{"addsL9713"},
		{"9713addsLe"},
		{"addsLe9713"},
		{"9713addsLeq"},
		{"addsLeq9713"},
		{"Za"},
		{"aZ"},
		{"Zad"},
		{"adZ"},
		{"Zadd"},
		{"addZ"},
		{"Zadds"},
		{"addsZ"},
		{"ZaddsL"},
		{"addsLZ"},
		{"ZaddsLe"},
		{"addsLeZ"},
		{"ZaddsLeq"},
		{"addsLeqZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.AddsLeq()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestSubsGtrValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"subsGtr", 7},
		{"subsGtr ", 7},
		{"subsGtr\n", 7},
		{"subsGtr.", 7},
		{"subsGtr:", 7},
		{"subsGtr,", 7},
		{"subsGtr\"", 7},
		{"subsGtr(", 7},
		{"subsGtr)", 7},
		{"subsGtr[", 7},
		{"subsGtr]", 7},
		{"subsGtr// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.SubsGtr()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestSubsGtrInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_s"},
		{"s_"},
		{"_su"},
		{"su_"},
		{"_sub"},
		{"sub_"},
		{"_subs"},
		{"subs_"},
		{"_subsG"},
		{"subsG_"},
		{"_subsGt"},
		{"subsGt_"},
		{"_subsGtr"},
		{"subsGtr_"},
		{"9713s"},
		{"s9713"},
		{"9713su"},
		{"su9713"},
		{"9713sub"},
		{"sub9713"},
		{"9713subs"},
		{"subs9713"},
		{"9713subsG"},
		{"subsG9713"},
		{"9713subsGt"},
		{"subsGt9713"},
		{"9713subsGtr"},
		{"subsGtr9713"},
		{"Zs"},
		{"sZ"},
		{"Zsu"},
		{"suZ"},
		{"Zsub"},
		{"subZ"},
		{"Zsubs"},
		{"subsZ"},
		{"ZsubsG"},
		{"subsGZ"},
		{"ZsubsGt"},
		{"subsGtZ"},
		{"ZsubsGtr"},
		{"subsGtrZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.SubsGtr()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestSubsGeqValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"subsGeq", 7},
		{"subsGeq ", 7},
		{"subsGeq\n", 7},
		{"subsGeq.", 7},
		{"subsGeq:", 7},
		{"subsGeq,", 7},
		{"subsGeq\"", 7},
		{"subsGeq(", 7},
		{"subsGeq)", 7},
		{"subsGeq[", 7},
		{"subsGeq]", 7},
		{"subsGeq// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.SubsGeq()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestSubsGeqInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_s"},
		{"s_"},
		{"_su"},
		{"su_"},
		{"_sub"},
		{"sub_"},
		{"_subs"},
		{"subs_"},
		{"_subsG"},
		{"subsG_"},
		{"_subsGe"},
		{"subsGe_"},
		{"_subsGeq"},
		{"subsGeq_"},
		{"9713s"},
		{"s9713"},
		{"9713su"},
		{"su9713"},
		{"9713sub"},
		{"sub9713"},
		{"9713subs"},
		{"subs9713"},
		{"9713subsG"},
		{"subsG9713"},
		{"9713subsGe"},
		{"subsGe9713"},
		{"9713subsGeq"},
		{"subsGeq9713"},
		{"Zs"},
		{"sZ"},
		{"Zsu"},
		{"suZ"},
		{"Zsub"},
		{"subZ"},
		{"Zsubs"},
		{"subsZ"},
		{"ZsubsG"},
		{"subsGZ"},
		{"ZsubsGe"},
		{"subsGeZ"},
		{"ZsubsGeq"},
		{"subsGeqZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.SubsGeq()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestMulsLssValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"mulsLss", 7},
		{"mulsLss ", 7},
		{"mulsLss\n", 7},
		{"mulsLss.", 7},
		{"mulsLss:", 7},
		{"mulsLss,", 7},
		{"mulsLss\"", 7},
		{"mulsLss(", 7},
		{"mulsLss)", 7},
		{"mulsLss[", 7},
		{"mulsLss]", 7},
		{"mulsLss// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.MulsLss()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestMulsLssInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_m"},
		{"m_"},
		{"_mu"},
		{"mu_"},
		{"_mul"},
		{"mul_"},
		{"_muls"},
		{"muls_"},
		{"_mulsL"},
		{"mulsL_"},
		{"_mulsLs"},
		{"mulsLs_"},
		{"_mulsLss"},
		{"mulsLss_"},
		{"9713m"},
		{"m9713"},
		{"9713mu"},
		{"mu9713"},
		{"9713mul"},
		{"mul9713"},
		{"9713muls"},
		{"muls9713"},
		{"9713mulsL"},
		{"mulsL9713"},
		{"9713mulsLs"},
		{"mulsLs9713"},
		{"9713mulsLss"},
		{"mulsLss9713"},
		{"Zm"},
		{"mZ"},
		{"Zmu"},
		{"muZ"},
		{"Zmul"},
		{"mulZ"},
		{"Zmuls"},
		{"mulsZ"},
		{"ZmulsL"},
		{"mulsLZ"},
		{"ZmulsLs"},
		{"mulsLsZ"},
		{"ZmulsLss"},
		{"mulsLssZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.MulsLss()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestMulsLeqValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"mulsLeq", 7},
		{"mulsLeq ", 7},
		{"mulsLeq\n", 7},
		{"mulsLeq.", 7},
		{"mulsLeq:", 7},
		{"mulsLeq,", 7},
		{"mulsLeq\"", 7},
		{"mulsLeq(", 7},
		{"mulsLeq)", 7},
		{"mulsLeq[", 7},
		{"mulsLeq]", 7},
		{"mulsLeq// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.MulsLeq()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestMulsLeqInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_m"},
		{"m_"},
		{"_mu"},
		{"mu_"},
		{"_mul"},
		{"mul_"},
		{"_muls"},
		{"muls_"},
		{"_mulsL"},
		{"mulsL_"},
		{"_mulsLe"},
		{"mulsLe_"},
		{"_mulsLeq"},
		{"mulsLeq_"},
		{"9713m"},
		{"m9713"},
		{"9713mu"},
		{"mu9713"},
		{"9713mul"},
		{"mul9713"},
		{"9713muls"},
		{"muls9713"},
		{"9713mulsL"},
		{"mulsL9713"},
		{"9713mulsLe"},
		{"mulsLe9713"},
		{"9713mulsLeq"},
		{"mulsLeq9713"},
		{"Zm"},
		{"mZ"},
		{"Zmu"},
		{"muZ"},
		{"Zmul"},
		{"mulZ"},
		{"Zmuls"},
		{"mulsZ"},
		{"ZmulsL"},
		{"mulsLZ"},
		{"ZmulsLe"},
		{"mulsLeZ"},
		{"ZmulsLeq"},
		{"mulsLeqZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.MulsLeq()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestDivsGtrValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"divsGtr", 7},
		{"divsGtr ", 7},
		{"divsGtr\n", 7},
		{"divsGtr.", 7},
		{"divsGtr:", 7},
		{"divsGtr,", 7},
		{"divsGtr\"", 7},
		{"divsGtr(", 7},
		{"divsGtr)", 7},
		{"divsGtr[", 7},
		{"divsGtr]", 7},
		{"divsGtr// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.DivsGtr()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestDivsGtrInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_d"},
		{"d_"},
		{"_di"},
		{"di_"},
		{"_div"},
		{"div_"},
		{"_divs"},
		{"divs_"},
		{"_divsG"},
		{"divsG_"},
		{"_divsGt"},
		{"divsGt_"},
		{"_divsGtr"},
		{"divsGtr_"},
		{"9713d"},
		{"d9713"},
		{"9713di"},
		{"di9713"},
		{"9713div"},
		{"div9713"},
		{"9713divs"},
		{"divs9713"},
		{"9713divsG"},
		{"divsG9713"},
		{"9713divsGt"},
		{"divsGt9713"},
		{"9713divsGtr"},
		{"divsGtr9713"},
		{"Zd"},
		{"dZ"},
		{"Zdi"},
		{"diZ"},
		{"Zdiv"},
		{"divZ"},
		{"Zdivs"},
		{"divsZ"},
		{"ZdivsG"},
		{"divsGZ"},
		{"ZdivsGt"},
		{"divsGtZ"},
		{"ZdivsGtr"},
		{"divsGtrZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.DivsGtr()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestDivsGeqValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"divsGeq", 7},
		{"divsGeq ", 7},
		{"divsGeq\n", 7},
		{"divsGeq.", 7},
		{"divsGeq:", 7},
		{"divsGeq,", 7},
		{"divsGeq\"", 7},
		{"divsGeq(", 7},
		{"divsGeq)", 7},
		{"divsGeq[", 7},
		{"divsGeq]", 7},
		{"divsGeq// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.DivsGeq()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestDivsGeqInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_d"},
		{"d_"},
		{"_di"},
		{"di_"},
		{"_div"},
		{"div_"},
		{"_divs"},
		{"divs_"},
		{"_divsG"},
		{"divsG_"},
		{"_divsGe"},
		{"divsGe_"},
		{"_divsGeq"},
		{"divsGeq_"},
		{"9713d"},
		{"d9713"},
		{"9713di"},
		{"di9713"},
		{"9713div"},
		{"div9713"},
		{"9713divs"},
		{"divs9713"},
		{"9713divsG"},
		{"divsG9713"},
		{"9713divsGe"},
		{"divsGe9713"},
		{"9713divsGeq"},
		{"divsGeq9713"},
		{"Zd"},
		{"dZ"},
		{"Zdi"},
		{"diZ"},
		{"Zdiv"},
		{"divZ"},
		{"Zdivs"},
		{"divsZ"},
		{"ZdivsG"},
		{"divsGZ"},
		{"ZdivsGe"},
		{"divsGeZ"},
		{"ZdivsGeq"},
		{"divsGeqZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.DivsGeq()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestFibsLeqValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"fibsLeq", 7},
		{"fibsLeq ", 7},
		{"fibsLeq\n", 7},
		{"fibsLeq.", 7},
		{"fibsLeq:", 7},
		{"fibsLeq,", 7},
		{"fibsLeq\"", 7},
		{"fibsLeq(", 7},
		{"fibsLeq)", 7},
		{"fibsLeq[", 7},
		{"fibsLeq]", 7},
		{"fibsLeq// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.FibsLeq()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestFibsLeqInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_f"},
		{"f_"},
		{"_fi"},
		{"fi_"},
		{"_fib"},
		{"fib_"},
		{"_fibs"},
		{"fibs_"},
		{"_fibsL"},
		{"fibsL_"},
		{"_fibsLe"},
		{"fibsLe_"},
		{"_fibsLeq"},
		{"fibsLeq_"},
		{"9713f"},
		{"f9713"},
		{"9713fi"},
		{"fi9713"},
		{"9713fib"},
		{"fib9713"},
		{"9713fibs"},
		{"fibs9713"},
		{"9713fibsL"},
		{"fibsL9713"},
		{"9713fibsLe"},
		{"fibsLe9713"},
		{"9713fibsLeq"},
		{"fibsLeq9713"},
		{"Zf"},
		{"fZ"},
		{"Zfi"},
		{"fiZ"},
		{"Zfib"},
		{"fibZ"},
		{"Zfibs"},
		{"fibsZ"},
		{"ZfibsL"},
		{"fibsLZ"},
		{"ZfibsLe"},
		{"fibsLeZ"},
		{"ZfibsLeq"},
		{"fibsLeqZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.FibsLeq()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestNewRngsValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"newRngs", 7},
		{"newRngs ", 7},
		{"newRngs\n", 7},
		{"newRngs.", 7},
		{"newRngs:", 7},
		{"newRngs,", 7},
		{"newRngs\"", 7},
		{"newRngs(", 7},
		{"newRngs)", 7},
		{"newRngs[", 7},
		{"newRngs]", 7},
		{"newRngs// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.NewRngs()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestNewRngsInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_n"},
		{"n_"},
		{"_ne"},
		{"ne_"},
		{"_new"},
		{"new_"},
		{"_newR"},
		{"newR_"},
		{"_newRn"},
		{"newRn_"},
		{"_newRng"},
		{"newRng_"},
		{"_newRngs"},
		{"newRngs_"},
		{"9713n"},
		{"n9713"},
		{"9713ne"},
		{"ne9713"},
		{"9713new"},
		{"new9713"},
		{"9713newR"},
		{"newR9713"},
		{"9713newRn"},
		{"newRn9713"},
		{"9713newRng"},
		{"newRng9713"},
		{"9713newRngs"},
		{"newRngs9713"},
		{"Zn"},
		{"nZ"},
		{"Zne"},
		{"neZ"},
		{"Znew"},
		{"newZ"},
		{"ZnewR"},
		{"newRZ"},
		{"ZnewRn"},
		{"newRnZ"},
		{"ZnewRng"},
		{"newRngZ"},
		{"ZnewRngs"},
		{"newRngsZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.NewRngs()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestMakeRngsValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"makeRngs", 8},
		{"makeRngs ", 8},
		{"makeRngs\n", 8},
		{"makeRngs.", 8},
		{"makeRngs:", 8},
		{"makeRngs,", 8},
		{"makeRngs\"", 8},
		{"makeRngs(", 8},
		{"makeRngs)", 8},
		{"makeRngs[", 8},
		{"makeRngs]", 8},
		{"makeRngs// comment", 8},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.MakeRngs()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestMakeRngsInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_m"},
		{"m_"},
		{"_ma"},
		{"ma_"},
		{"_mak"},
		{"mak_"},
		{"_make"},
		{"make_"},
		{"_makeR"},
		{"makeR_"},
		{"_makeRn"},
		{"makeRn_"},
		{"_makeRng"},
		{"makeRng_"},
		{"_makeRngs"},
		{"makeRngs_"},
		{"9713m"},
		{"m9713"},
		{"9713ma"},
		{"ma9713"},
		{"9713mak"},
		{"mak9713"},
		{"9713make"},
		{"make9713"},
		{"9713makeR"},
		{"makeR9713"},
		{"9713makeRn"},
		{"makeRn9713"},
		{"9713makeRng"},
		{"makeRng9713"},
		{"9713makeRngs"},
		{"makeRngs9713"},
		{"Zm"},
		{"mZ"},
		{"Zma"},
		{"maZ"},
		{"Zmak"},
		{"makZ"},
		{"Zmake"},
		{"makeZ"},
		{"ZmakeR"},
		{"makeRZ"},
		{"ZmakeRn"},
		{"makeRnZ"},
		{"ZmakeRng"},
		{"makeRngZ"},
		{"ZmakeRngs"},
		{"makeRngsZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.MakeRngs()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestMakeEmpRngsValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"makeEmpRngs", 11},
		{"makeEmpRngs ", 11},
		{"makeEmpRngs\n", 11},
		{"makeEmpRngs.", 11},
		{"makeEmpRngs:", 11},
		{"makeEmpRngs,", 11},
		{"makeEmpRngs\"", 11},
		{"makeEmpRngs(", 11},
		{"makeEmpRngs)", 11},
		{"makeEmpRngs[", 11},
		{"makeEmpRngs]", 11},
		{"makeEmpRngs// comment", 11},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.MakeEmpRngs()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestMakeEmpRngsInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_m"},
		{"m_"},
		{"_ma"},
		{"ma_"},
		{"_mak"},
		{"mak_"},
		{"_make"},
		{"make_"},
		{"_makeE"},
		{"makeE_"},
		{"_makeEm"},
		{"makeEm_"},
		{"_makeEmp"},
		{"makeEmp_"},
		{"_makeEmpR"},
		{"makeEmpR_"},
		{"_makeEmpRn"},
		{"makeEmpRn_"},
		{"_makeEmpRng"},
		{"makeEmpRng_"},
		{"_makeEmpRngs"},
		{"makeEmpRngs_"},
		{"9713m"},
		{"m9713"},
		{"9713ma"},
		{"ma9713"},
		{"9713mak"},
		{"mak9713"},
		{"9713make"},
		{"make9713"},
		{"9713makeE"},
		{"makeE9713"},
		{"9713makeEm"},
		{"makeEm9713"},
		{"9713makeEmp"},
		{"makeEmp9713"},
		{"9713makeEmpR"},
		{"makeEmpR9713"},
		{"9713makeEmpRn"},
		{"makeEmpRn9713"},
		{"9713makeEmpRng"},
		{"makeEmpRng9713"},
		{"9713makeEmpRngs"},
		{"makeEmpRngs9713"},
		{"Zm"},
		{"mZ"},
		{"Zma"},
		{"maZ"},
		{"Zmak"},
		{"makZ"},
		{"Zmake"},
		{"makeZ"},
		{"ZmakeE"},
		{"makeEZ"},
		{"ZmakeEm"},
		{"makeEmZ"},
		{"ZmakeEmp"},
		{"makeEmpZ"},
		{"ZmakeEmpR"},
		{"makeEmpRZ"},
		{"ZmakeEmpRn"},
		{"makeEmpRnZ"},
		{"ZmakeEmpRng"},
		{"makeEmpRngZ"},
		{"ZmakeEmpRngs"},
		{"makeEmpRngsZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.MakeEmpRngs()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestNewTrdsValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"newTrds", 7},
		{"newTrds ", 7},
		{"newTrds\n", 7},
		{"newTrds.", 7},
		{"newTrds:", 7},
		{"newTrds,", 7},
		{"newTrds\"", 7},
		{"newTrds(", 7},
		{"newTrds)", 7},
		{"newTrds[", 7},
		{"newTrds]", 7},
		{"newTrds// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.NewTrds()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestNewTrdsInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_n"},
		{"n_"},
		{"_ne"},
		{"ne_"},
		{"_new"},
		{"new_"},
		{"_newT"},
		{"newT_"},
		{"_newTr"},
		{"newTr_"},
		{"_newTrd"},
		{"newTrd_"},
		{"_newTrds"},
		{"newTrds_"},
		{"9713n"},
		{"n9713"},
		{"9713ne"},
		{"ne9713"},
		{"9713new"},
		{"new9713"},
		{"9713newT"},
		{"newT9713"},
		{"9713newTr"},
		{"newTr9713"},
		{"9713newTrd"},
		{"newTrd9713"},
		{"9713newTrds"},
		{"newTrds9713"},
		{"Zn"},
		{"nZ"},
		{"Zne"},
		{"neZ"},
		{"Znew"},
		{"newZ"},
		{"ZnewT"},
		{"newTZ"},
		{"ZnewTr"},
		{"newTrZ"},
		{"ZnewTrd"},
		{"newTrdZ"},
		{"ZnewTrds"},
		{"newTrdsZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.NewTrds()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestMakeTrdsValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"makeTrds", 8},
		{"makeTrds ", 8},
		{"makeTrds\n", 8},
		{"makeTrds.", 8},
		{"makeTrds:", 8},
		{"makeTrds,", 8},
		{"makeTrds\"", 8},
		{"makeTrds(", 8},
		{"makeTrds)", 8},
		{"makeTrds[", 8},
		{"makeTrds]", 8},
		{"makeTrds// comment", 8},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.MakeTrds()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestMakeTrdsInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_m"},
		{"m_"},
		{"_ma"},
		{"ma_"},
		{"_mak"},
		{"mak_"},
		{"_make"},
		{"make_"},
		{"_makeT"},
		{"makeT_"},
		{"_makeTr"},
		{"makeTr_"},
		{"_makeTrd"},
		{"makeTrd_"},
		{"_makeTrds"},
		{"makeTrds_"},
		{"9713m"},
		{"m9713"},
		{"9713ma"},
		{"ma9713"},
		{"9713mak"},
		{"mak9713"},
		{"9713make"},
		{"make9713"},
		{"9713makeT"},
		{"makeT9713"},
		{"9713makeTr"},
		{"makeTr9713"},
		{"9713makeTrd"},
		{"makeTrd9713"},
		{"9713makeTrds"},
		{"makeTrds9713"},
		{"Zm"},
		{"mZ"},
		{"Zma"},
		{"maZ"},
		{"Zmak"},
		{"makZ"},
		{"Zmake"},
		{"makeZ"},
		{"ZmakeT"},
		{"makeTZ"},
		{"ZmakeTr"},
		{"makeTrZ"},
		{"ZmakeTrd"},
		{"makeTrdZ"},
		{"ZmakeTrds"},
		{"makeTrdsZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.MakeTrds()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestMakeEmpTrdsValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"makeEmpTrds", 11},
		{"makeEmpTrds ", 11},
		{"makeEmpTrds\n", 11},
		{"makeEmpTrds.", 11},
		{"makeEmpTrds:", 11},
		{"makeEmpTrds,", 11},
		{"makeEmpTrds\"", 11},
		{"makeEmpTrds(", 11},
		{"makeEmpTrds)", 11},
		{"makeEmpTrds[", 11},
		{"makeEmpTrds]", 11},
		{"makeEmpTrds// comment", 11},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.MakeEmpTrds()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestMakeEmpTrdsInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_m"},
		{"m_"},
		{"_ma"},
		{"ma_"},
		{"_mak"},
		{"mak_"},
		{"_make"},
		{"make_"},
		{"_makeE"},
		{"makeE_"},
		{"_makeEm"},
		{"makeEm_"},
		{"_makeEmp"},
		{"makeEmp_"},
		{"_makeEmpT"},
		{"makeEmpT_"},
		{"_makeEmpTr"},
		{"makeEmpTr_"},
		{"_makeEmpTrd"},
		{"makeEmpTrd_"},
		{"_makeEmpTrds"},
		{"makeEmpTrds_"},
		{"9713m"},
		{"m9713"},
		{"9713ma"},
		{"ma9713"},
		{"9713mak"},
		{"mak9713"},
		{"9713make"},
		{"make9713"},
		{"9713makeE"},
		{"makeE9713"},
		{"9713makeEm"},
		{"makeEm9713"},
		{"9713makeEmp"},
		{"makeEmp9713"},
		{"9713makeEmpT"},
		{"makeEmpT9713"},
		{"9713makeEmpTr"},
		{"makeEmpTr9713"},
		{"9713makeEmpTrd"},
		{"makeEmpTrd9713"},
		{"9713makeEmpTrds"},
		{"makeEmpTrds9713"},
		{"Zm"},
		{"mZ"},
		{"Zma"},
		{"maZ"},
		{"Zmak"},
		{"makZ"},
		{"Zmake"},
		{"makeZ"},
		{"ZmakeE"},
		{"makeEZ"},
		{"ZmakeEm"},
		{"makeEmZ"},
		{"ZmakeEmp"},
		{"makeEmpZ"},
		{"ZmakeEmpT"},
		{"makeEmpTZ"},
		{"ZmakeEmpTr"},
		{"makeEmpTrZ"},
		{"ZmakeEmpTrd"},
		{"makeEmpTrdZ"},
		{"ZmakeEmpTrds"},
		{"makeEmpTrdsZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.MakeEmpTrds()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestNewPrfmsValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"newPrfms", 8},
		{"newPrfms ", 8},
		{"newPrfms\n", 8},
		{"newPrfms.", 8},
		{"newPrfms:", 8},
		{"newPrfms,", 8},
		{"newPrfms\"", 8},
		{"newPrfms(", 8},
		{"newPrfms)", 8},
		{"newPrfms[", 8},
		{"newPrfms]", 8},
		{"newPrfms// comment", 8},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.NewPrfms()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestNewPrfmsInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_n"},
		{"n_"},
		{"_ne"},
		{"ne_"},
		{"_new"},
		{"new_"},
		{"_newP"},
		{"newP_"},
		{"_newPr"},
		{"newPr_"},
		{"_newPrf"},
		{"newPrf_"},
		{"_newPrfm"},
		{"newPrfm_"},
		{"_newPrfms"},
		{"newPrfms_"},
		{"9713n"},
		{"n9713"},
		{"9713ne"},
		{"ne9713"},
		{"9713new"},
		{"new9713"},
		{"9713newP"},
		{"newP9713"},
		{"9713newPr"},
		{"newPr9713"},
		{"9713newPrf"},
		{"newPrf9713"},
		{"9713newPrfm"},
		{"newPrfm9713"},
		{"9713newPrfms"},
		{"newPrfms9713"},
		{"Zn"},
		{"nZ"},
		{"Zne"},
		{"neZ"},
		{"Znew"},
		{"newZ"},
		{"ZnewP"},
		{"newPZ"},
		{"ZnewPr"},
		{"newPrZ"},
		{"ZnewPrf"},
		{"newPrfZ"},
		{"ZnewPrfm"},
		{"newPrfmZ"},
		{"ZnewPrfms"},
		{"newPrfmsZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.NewPrfms()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestMakePrfmsValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"makePrfms", 9},
		{"makePrfms ", 9},
		{"makePrfms\n", 9},
		{"makePrfms.", 9},
		{"makePrfms:", 9},
		{"makePrfms,", 9},
		{"makePrfms\"", 9},
		{"makePrfms(", 9},
		{"makePrfms)", 9},
		{"makePrfms[", 9},
		{"makePrfms]", 9},
		{"makePrfms// comment", 9},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.MakePrfms()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestMakePrfmsInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_m"},
		{"m_"},
		{"_ma"},
		{"ma_"},
		{"_mak"},
		{"mak_"},
		{"_make"},
		{"make_"},
		{"_makeP"},
		{"makeP_"},
		{"_makePr"},
		{"makePr_"},
		{"_makePrf"},
		{"makePrf_"},
		{"_makePrfm"},
		{"makePrfm_"},
		{"_makePrfms"},
		{"makePrfms_"},
		{"9713m"},
		{"m9713"},
		{"9713ma"},
		{"ma9713"},
		{"9713mak"},
		{"mak9713"},
		{"9713make"},
		{"make9713"},
		{"9713makeP"},
		{"makeP9713"},
		{"9713makePr"},
		{"makePr9713"},
		{"9713makePrf"},
		{"makePrf9713"},
		{"9713makePrfm"},
		{"makePrfm9713"},
		{"9713makePrfms"},
		{"makePrfms9713"},
		{"Zm"},
		{"mZ"},
		{"Zma"},
		{"maZ"},
		{"Zmak"},
		{"makZ"},
		{"Zmake"},
		{"makeZ"},
		{"ZmakeP"},
		{"makePZ"},
		{"ZmakePr"},
		{"makePrZ"},
		{"ZmakePrf"},
		{"makePrfZ"},
		{"ZmakePrfm"},
		{"makePrfmZ"},
		{"ZmakePrfms"},
		{"makePrfmsZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.MakePrfms()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestMakeEmpPrfmsValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"makeEmpPrfms", 12},
		{"makeEmpPrfms ", 12},
		{"makeEmpPrfms\n", 12},
		{"makeEmpPrfms.", 12},
		{"makeEmpPrfms:", 12},
		{"makeEmpPrfms,", 12},
		{"makeEmpPrfms\"", 12},
		{"makeEmpPrfms(", 12},
		{"makeEmpPrfms)", 12},
		{"makeEmpPrfms[", 12},
		{"makeEmpPrfms]", 12},
		{"makeEmpPrfms// comment", 12},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.MakeEmpPrfms()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestMakeEmpPrfmsInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_m"},
		{"m_"},
		{"_ma"},
		{"ma_"},
		{"_mak"},
		{"mak_"},
		{"_make"},
		{"make_"},
		{"_makeE"},
		{"makeE_"},
		{"_makeEm"},
		{"makeEm_"},
		{"_makeEmp"},
		{"makeEmp_"},
		{"_makeEmpP"},
		{"makeEmpP_"},
		{"_makeEmpPr"},
		{"makeEmpPr_"},
		{"_makeEmpPrf"},
		{"makeEmpPrf_"},
		{"_makeEmpPrfm"},
		{"makeEmpPrfm_"},
		{"_makeEmpPrfms"},
		{"makeEmpPrfms_"},
		{"9713m"},
		{"m9713"},
		{"9713ma"},
		{"ma9713"},
		{"9713mak"},
		{"mak9713"},
		{"9713make"},
		{"make9713"},
		{"9713makeE"},
		{"makeE9713"},
		{"9713makeEm"},
		{"makeEm9713"},
		{"9713makeEmp"},
		{"makeEmp9713"},
		{"9713makeEmpP"},
		{"makeEmpP9713"},
		{"9713makeEmpPr"},
		{"makeEmpPr9713"},
		{"9713makeEmpPrf"},
		{"makeEmpPrf9713"},
		{"9713makeEmpPrfm"},
		{"makeEmpPrfm9713"},
		{"9713makeEmpPrfms"},
		{"makeEmpPrfms9713"},
		{"Zm"},
		{"mZ"},
		{"Zma"},
		{"maZ"},
		{"Zmak"},
		{"makZ"},
		{"Zmake"},
		{"makeZ"},
		{"ZmakeE"},
		{"makeEZ"},
		{"ZmakeEm"},
		{"makeEmZ"},
		{"ZmakeEmp"},
		{"makeEmpZ"},
		{"ZmakeEmpP"},
		{"makeEmpPZ"},
		{"ZmakeEmpPr"},
		{"makeEmpPrZ"},
		{"ZmakeEmpPrf"},
		{"makeEmpPrfZ"},
		{"ZmakeEmpPrfm"},
		{"makeEmpPrfmZ"},
		{"ZmakeEmpPrfms"},
		{"makeEmpPrfmsZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.MakeEmpPrfms()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestOanValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"oan", 3},
		{"oan ", 3},
		{"oan\n", 3},
		{"oan.", 3},
		{"oan:", 3},
		{"oan,", 3},
		{"oan\"", 3},
		{"oan(", 3},
		{"oan)", 3},
		{"oan[", 3},
		{"oan]", 3},
		{"oan// comment", 3},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Oan()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestOanInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_o"},
		{"o_"},
		{"_oa"},
		{"oa_"},
		{"_oan"},
		{"oan_"},
		{"9713o"},
		{"o9713"},
		{"9713oa"},
		{"oa9713"},
		{"9713oan"},
		{"oan9713"},
		{"Zo"},
		{"oZ"},
		{"Zoa"},
		{"oaZ"},
		{"Zoan"},
		{"oanZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Oan()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestNewPrvsValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"newPrvs", 7},
		{"newPrvs ", 7},
		{"newPrvs\n", 7},
		{"newPrvs.", 7},
		{"newPrvs:", 7},
		{"newPrvs,", 7},
		{"newPrvs\"", 7},
		{"newPrvs(", 7},
		{"newPrvs)", 7},
		{"newPrvs[", 7},
		{"newPrvs]", 7},
		{"newPrvs// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.NewPrvs()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestNewPrvsInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_n"},
		{"n_"},
		{"_ne"},
		{"ne_"},
		{"_new"},
		{"new_"},
		{"_newP"},
		{"newP_"},
		{"_newPr"},
		{"newPr_"},
		{"_newPrv"},
		{"newPrv_"},
		{"_newPrvs"},
		{"newPrvs_"},
		{"9713n"},
		{"n9713"},
		{"9713ne"},
		{"ne9713"},
		{"9713new"},
		{"new9713"},
		{"9713newP"},
		{"newP9713"},
		{"9713newPr"},
		{"newPr9713"},
		{"9713newPrv"},
		{"newPrv9713"},
		{"9713newPrvs"},
		{"newPrvs9713"},
		{"Zn"},
		{"nZ"},
		{"Zne"},
		{"neZ"},
		{"Znew"},
		{"newZ"},
		{"ZnewP"},
		{"newPZ"},
		{"ZnewPr"},
		{"newPrZ"},
		{"ZnewPrv"},
		{"newPrvZ"},
		{"ZnewPrvs"},
		{"newPrvsZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.NewPrvs()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestMakePrvsValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"makePrvs", 8},
		{"makePrvs ", 8},
		{"makePrvs\n", 8},
		{"makePrvs.", 8},
		{"makePrvs:", 8},
		{"makePrvs,", 8},
		{"makePrvs\"", 8},
		{"makePrvs(", 8},
		{"makePrvs)", 8},
		{"makePrvs[", 8},
		{"makePrvs]", 8},
		{"makePrvs// comment", 8},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.MakePrvs()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestMakePrvsInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_m"},
		{"m_"},
		{"_ma"},
		{"ma_"},
		{"_mak"},
		{"mak_"},
		{"_make"},
		{"make_"},
		{"_makeP"},
		{"makeP_"},
		{"_makePr"},
		{"makePr_"},
		{"_makePrv"},
		{"makePrv_"},
		{"_makePrvs"},
		{"makePrvs_"},
		{"9713m"},
		{"m9713"},
		{"9713ma"},
		{"ma9713"},
		{"9713mak"},
		{"mak9713"},
		{"9713make"},
		{"make9713"},
		{"9713makeP"},
		{"makeP9713"},
		{"9713makePr"},
		{"makePr9713"},
		{"9713makePrv"},
		{"makePrv9713"},
		{"9713makePrvs"},
		{"makePrvs9713"},
		{"Zm"},
		{"mZ"},
		{"Zma"},
		{"maZ"},
		{"Zmak"},
		{"makZ"},
		{"Zmake"},
		{"makeZ"},
		{"ZmakeP"},
		{"makePZ"},
		{"ZmakePr"},
		{"makePrZ"},
		{"ZmakePrv"},
		{"makePrvZ"},
		{"ZmakePrvs"},
		{"makePrvsZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.MakePrvs()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestMakeEmpPrvsValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"makeEmpPrvs", 11},
		{"makeEmpPrvs ", 11},
		{"makeEmpPrvs\n", 11},
		{"makeEmpPrvs.", 11},
		{"makeEmpPrvs:", 11},
		{"makeEmpPrvs,", 11},
		{"makeEmpPrvs\"", 11},
		{"makeEmpPrvs(", 11},
		{"makeEmpPrvs)", 11},
		{"makeEmpPrvs[", 11},
		{"makeEmpPrvs]", 11},
		{"makeEmpPrvs// comment", 11},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.MakeEmpPrvs()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestMakeEmpPrvsInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_m"},
		{"m_"},
		{"_ma"},
		{"ma_"},
		{"_mak"},
		{"mak_"},
		{"_make"},
		{"make_"},
		{"_makeE"},
		{"makeE_"},
		{"_makeEm"},
		{"makeEm_"},
		{"_makeEmp"},
		{"makeEmp_"},
		{"_makeEmpP"},
		{"makeEmpP_"},
		{"_makeEmpPr"},
		{"makeEmpPr_"},
		{"_makeEmpPrv"},
		{"makeEmpPrv_"},
		{"_makeEmpPrvs"},
		{"makeEmpPrvs_"},
		{"9713m"},
		{"m9713"},
		{"9713ma"},
		{"ma9713"},
		{"9713mak"},
		{"mak9713"},
		{"9713make"},
		{"make9713"},
		{"9713makeE"},
		{"makeE9713"},
		{"9713makeEm"},
		{"makeEm9713"},
		{"9713makeEmp"},
		{"makeEmp9713"},
		{"9713makeEmpP"},
		{"makeEmpP9713"},
		{"9713makeEmpPr"},
		{"makeEmpPr9713"},
		{"9713makeEmpPrv"},
		{"makeEmpPrv9713"},
		{"9713makeEmpPrvs"},
		{"makeEmpPrvs9713"},
		{"Zm"},
		{"mZ"},
		{"Zma"},
		{"maZ"},
		{"Zmak"},
		{"makZ"},
		{"Zmake"},
		{"makeZ"},
		{"ZmakeE"},
		{"makeEZ"},
		{"ZmakeEm"},
		{"makeEmZ"},
		{"ZmakeEmp"},
		{"makeEmpZ"},
		{"ZmakeEmpP"},
		{"makeEmpPZ"},
		{"ZmakeEmpPr"},
		{"makeEmpPrZ"},
		{"ZmakeEmpPrv"},
		{"makeEmpPrvZ"},
		{"ZmakeEmpPrvs"},
		{"makeEmpPrvsZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.MakeEmpPrvs()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestNewInstrsValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"newInstrs", 9},
		{"newInstrs ", 9},
		{"newInstrs\n", 9},
		{"newInstrs.", 9},
		{"newInstrs:", 9},
		{"newInstrs,", 9},
		{"newInstrs\"", 9},
		{"newInstrs(", 9},
		{"newInstrs)", 9},
		{"newInstrs[", 9},
		{"newInstrs]", 9},
		{"newInstrs// comment", 9},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.NewInstrs()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestNewInstrsInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_n"},
		{"n_"},
		{"_ne"},
		{"ne_"},
		{"_new"},
		{"new_"},
		{"_newI"},
		{"newI_"},
		{"_newIn"},
		{"newIn_"},
		{"_newIns"},
		{"newIns_"},
		{"_newInst"},
		{"newInst_"},
		{"_newInstr"},
		{"newInstr_"},
		{"_newInstrs"},
		{"newInstrs_"},
		{"9713n"},
		{"n9713"},
		{"9713ne"},
		{"ne9713"},
		{"9713new"},
		{"new9713"},
		{"9713newI"},
		{"newI9713"},
		{"9713newIn"},
		{"newIn9713"},
		{"9713newIns"},
		{"newIns9713"},
		{"9713newInst"},
		{"newInst9713"},
		{"9713newInstr"},
		{"newInstr9713"},
		{"9713newInstrs"},
		{"newInstrs9713"},
		{"Zn"},
		{"nZ"},
		{"Zne"},
		{"neZ"},
		{"Znew"},
		{"newZ"},
		{"ZnewI"},
		{"newIZ"},
		{"ZnewIn"},
		{"newInZ"},
		{"ZnewIns"},
		{"newInsZ"},
		{"ZnewInst"},
		{"newInstZ"},
		{"ZnewInstr"},
		{"newInstrZ"},
		{"ZnewInstrs"},
		{"newInstrsZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.NewInstrs()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestMakeInstrsValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"makeInstrs", 10},
		{"makeInstrs ", 10},
		{"makeInstrs\n", 10},
		{"makeInstrs.", 10},
		{"makeInstrs:", 10},
		{"makeInstrs,", 10},
		{"makeInstrs\"", 10},
		{"makeInstrs(", 10},
		{"makeInstrs)", 10},
		{"makeInstrs[", 10},
		{"makeInstrs]", 10},
		{"makeInstrs// comment", 10},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.MakeInstrs()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestMakeInstrsInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_m"},
		{"m_"},
		{"_ma"},
		{"ma_"},
		{"_mak"},
		{"mak_"},
		{"_make"},
		{"make_"},
		{"_makeI"},
		{"makeI_"},
		{"_makeIn"},
		{"makeIn_"},
		{"_makeIns"},
		{"makeIns_"},
		{"_makeInst"},
		{"makeInst_"},
		{"_makeInstr"},
		{"makeInstr_"},
		{"_makeInstrs"},
		{"makeInstrs_"},
		{"9713m"},
		{"m9713"},
		{"9713ma"},
		{"ma9713"},
		{"9713mak"},
		{"mak9713"},
		{"9713make"},
		{"make9713"},
		{"9713makeI"},
		{"makeI9713"},
		{"9713makeIn"},
		{"makeIn9713"},
		{"9713makeIns"},
		{"makeIns9713"},
		{"9713makeInst"},
		{"makeInst9713"},
		{"9713makeInstr"},
		{"makeInstr9713"},
		{"9713makeInstrs"},
		{"makeInstrs9713"},
		{"Zm"},
		{"mZ"},
		{"Zma"},
		{"maZ"},
		{"Zmak"},
		{"makZ"},
		{"Zmake"},
		{"makeZ"},
		{"ZmakeI"},
		{"makeIZ"},
		{"ZmakeIn"},
		{"makeInZ"},
		{"ZmakeIns"},
		{"makeInsZ"},
		{"ZmakeInst"},
		{"makeInstZ"},
		{"ZmakeInstr"},
		{"makeInstrZ"},
		{"ZmakeInstrs"},
		{"makeInstrsZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.MakeInstrs()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestMakeEmpInstrsValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"makeEmpInstrs", 13},
		{"makeEmpInstrs ", 13},
		{"makeEmpInstrs\n", 13},
		{"makeEmpInstrs.", 13},
		{"makeEmpInstrs:", 13},
		{"makeEmpInstrs,", 13},
		{"makeEmpInstrs\"", 13},
		{"makeEmpInstrs(", 13},
		{"makeEmpInstrs)", 13},
		{"makeEmpInstrs[", 13},
		{"makeEmpInstrs]", 13},
		{"makeEmpInstrs// comment", 13},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.MakeEmpInstrs()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestMakeEmpInstrsInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_m"},
		{"m_"},
		{"_ma"},
		{"ma_"},
		{"_mak"},
		{"mak_"},
		{"_make"},
		{"make_"},
		{"_makeE"},
		{"makeE_"},
		{"_makeEm"},
		{"makeEm_"},
		{"_makeEmp"},
		{"makeEmp_"},
		{"_makeEmpI"},
		{"makeEmpI_"},
		{"_makeEmpIn"},
		{"makeEmpIn_"},
		{"_makeEmpIns"},
		{"makeEmpIns_"},
		{"_makeEmpInst"},
		{"makeEmpInst_"},
		{"_makeEmpInstr"},
		{"makeEmpInstr_"},
		{"_makeEmpInstrs"},
		{"makeEmpInstrs_"},
		{"9713m"},
		{"m9713"},
		{"9713ma"},
		{"ma9713"},
		{"9713mak"},
		{"mak9713"},
		{"9713make"},
		{"make9713"},
		{"9713makeE"},
		{"makeE9713"},
		{"9713makeEm"},
		{"makeEm9713"},
		{"9713makeEmp"},
		{"makeEmp9713"},
		{"9713makeEmpI"},
		{"makeEmpI9713"},
		{"9713makeEmpIn"},
		{"makeEmpIn9713"},
		{"9713makeEmpIns"},
		{"makeEmpIns9713"},
		{"9713makeEmpInst"},
		{"makeEmpInst9713"},
		{"9713makeEmpInstr"},
		{"makeEmpInstr9713"},
		{"9713makeEmpInstrs"},
		{"makeEmpInstrs9713"},
		{"Zm"},
		{"mZ"},
		{"Zma"},
		{"maZ"},
		{"Zmak"},
		{"makZ"},
		{"Zmake"},
		{"makeZ"},
		{"ZmakeE"},
		{"makeEZ"},
		{"ZmakeEm"},
		{"makeEmZ"},
		{"ZmakeEmp"},
		{"makeEmpZ"},
		{"ZmakeEmpI"},
		{"makeEmpIZ"},
		{"ZmakeEmpIn"},
		{"makeEmpInZ"},
		{"ZmakeEmpIns"},
		{"makeEmpInsZ"},
		{"ZmakeEmpInst"},
		{"makeEmpInstZ"},
		{"ZmakeEmpInstr"},
		{"makeEmpInstrZ"},
		{"ZmakeEmpInstrs"},
		{"makeEmpInstrsZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.MakeEmpInstrs()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestNewInrvlsValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"newInrvls", 9},
		{"newInrvls ", 9},
		{"newInrvls\n", 9},
		{"newInrvls.", 9},
		{"newInrvls:", 9},
		{"newInrvls,", 9},
		{"newInrvls\"", 9},
		{"newInrvls(", 9},
		{"newInrvls)", 9},
		{"newInrvls[", 9},
		{"newInrvls]", 9},
		{"newInrvls// comment", 9},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.NewInrvls()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestNewInrvlsInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_n"},
		{"n_"},
		{"_ne"},
		{"ne_"},
		{"_new"},
		{"new_"},
		{"_newI"},
		{"newI_"},
		{"_newIn"},
		{"newIn_"},
		{"_newInr"},
		{"newInr_"},
		{"_newInrv"},
		{"newInrv_"},
		{"_newInrvl"},
		{"newInrvl_"},
		{"_newInrvls"},
		{"newInrvls_"},
		{"9713n"},
		{"n9713"},
		{"9713ne"},
		{"ne9713"},
		{"9713new"},
		{"new9713"},
		{"9713newI"},
		{"newI9713"},
		{"9713newIn"},
		{"newIn9713"},
		{"9713newInr"},
		{"newInr9713"},
		{"9713newInrv"},
		{"newInrv9713"},
		{"9713newInrvl"},
		{"newInrvl9713"},
		{"9713newInrvls"},
		{"newInrvls9713"},
		{"Zn"},
		{"nZ"},
		{"Zne"},
		{"neZ"},
		{"Znew"},
		{"newZ"},
		{"ZnewI"},
		{"newIZ"},
		{"ZnewIn"},
		{"newInZ"},
		{"ZnewInr"},
		{"newInrZ"},
		{"ZnewInrv"},
		{"newInrvZ"},
		{"ZnewInrvl"},
		{"newInrvlZ"},
		{"ZnewInrvls"},
		{"newInrvlsZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.NewInrvls()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestMakeInrvlsValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"makeInrvls", 10},
		{"makeInrvls ", 10},
		{"makeInrvls\n", 10},
		{"makeInrvls.", 10},
		{"makeInrvls:", 10},
		{"makeInrvls,", 10},
		{"makeInrvls\"", 10},
		{"makeInrvls(", 10},
		{"makeInrvls)", 10},
		{"makeInrvls[", 10},
		{"makeInrvls]", 10},
		{"makeInrvls// comment", 10},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.MakeInrvls()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestMakeInrvlsInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_m"},
		{"m_"},
		{"_ma"},
		{"ma_"},
		{"_mak"},
		{"mak_"},
		{"_make"},
		{"make_"},
		{"_makeI"},
		{"makeI_"},
		{"_makeIn"},
		{"makeIn_"},
		{"_makeInr"},
		{"makeInr_"},
		{"_makeInrv"},
		{"makeInrv_"},
		{"_makeInrvl"},
		{"makeInrvl_"},
		{"_makeInrvls"},
		{"makeInrvls_"},
		{"9713m"},
		{"m9713"},
		{"9713ma"},
		{"ma9713"},
		{"9713mak"},
		{"mak9713"},
		{"9713make"},
		{"make9713"},
		{"9713makeI"},
		{"makeI9713"},
		{"9713makeIn"},
		{"makeIn9713"},
		{"9713makeInr"},
		{"makeInr9713"},
		{"9713makeInrv"},
		{"makeInrv9713"},
		{"9713makeInrvl"},
		{"makeInrvl9713"},
		{"9713makeInrvls"},
		{"makeInrvls9713"},
		{"Zm"},
		{"mZ"},
		{"Zma"},
		{"maZ"},
		{"Zmak"},
		{"makZ"},
		{"Zmake"},
		{"makeZ"},
		{"ZmakeI"},
		{"makeIZ"},
		{"ZmakeIn"},
		{"makeInZ"},
		{"ZmakeInr"},
		{"makeInrZ"},
		{"ZmakeInrv"},
		{"makeInrvZ"},
		{"ZmakeInrvl"},
		{"makeInrvlZ"},
		{"ZmakeInrvls"},
		{"makeInrvlsZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.MakeInrvls()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestMakeEmpInrvlsValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"makeEmpInrvls", 13},
		{"makeEmpInrvls ", 13},
		{"makeEmpInrvls\n", 13},
		{"makeEmpInrvls.", 13},
		{"makeEmpInrvls:", 13},
		{"makeEmpInrvls,", 13},
		{"makeEmpInrvls\"", 13},
		{"makeEmpInrvls(", 13},
		{"makeEmpInrvls)", 13},
		{"makeEmpInrvls[", 13},
		{"makeEmpInrvls]", 13},
		{"makeEmpInrvls// comment", 13},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.MakeEmpInrvls()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestMakeEmpInrvlsInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_m"},
		{"m_"},
		{"_ma"},
		{"ma_"},
		{"_mak"},
		{"mak_"},
		{"_make"},
		{"make_"},
		{"_makeE"},
		{"makeE_"},
		{"_makeEm"},
		{"makeEm_"},
		{"_makeEmp"},
		{"makeEmp_"},
		{"_makeEmpI"},
		{"makeEmpI_"},
		{"_makeEmpIn"},
		{"makeEmpIn_"},
		{"_makeEmpInr"},
		{"makeEmpInr_"},
		{"_makeEmpInrv"},
		{"makeEmpInrv_"},
		{"_makeEmpInrvl"},
		{"makeEmpInrvl_"},
		{"_makeEmpInrvls"},
		{"makeEmpInrvls_"},
		{"9713m"},
		{"m9713"},
		{"9713ma"},
		{"ma9713"},
		{"9713mak"},
		{"mak9713"},
		{"9713make"},
		{"make9713"},
		{"9713makeE"},
		{"makeE9713"},
		{"9713makeEm"},
		{"makeEm9713"},
		{"9713makeEmp"},
		{"makeEmp9713"},
		{"9713makeEmpI"},
		{"makeEmpI9713"},
		{"9713makeEmpIn"},
		{"makeEmpIn9713"},
		{"9713makeEmpInr"},
		{"makeEmpInr9713"},
		{"9713makeEmpInrv"},
		{"makeEmpInrv9713"},
		{"9713makeEmpInrvl"},
		{"makeEmpInrvl9713"},
		{"9713makeEmpInrvls"},
		{"makeEmpInrvls9713"},
		{"Zm"},
		{"mZ"},
		{"Zma"},
		{"maZ"},
		{"Zmak"},
		{"makZ"},
		{"Zmake"},
		{"makeZ"},
		{"ZmakeE"},
		{"makeEZ"},
		{"ZmakeEm"},
		{"makeEmZ"},
		{"ZmakeEmp"},
		{"makeEmpZ"},
		{"ZmakeEmpI"},
		{"makeEmpIZ"},
		{"ZmakeEmpIn"},
		{"makeEmpInZ"},
		{"ZmakeEmpInr"},
		{"makeEmpInrZ"},
		{"ZmakeEmpInrv"},
		{"makeEmpInrvZ"},
		{"ZmakeEmpInrvl"},
		{"makeEmpInrvlZ"},
		{"ZmakeEmpInrvls"},
		{"makeEmpInrvlsZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.MakeEmpInrvls()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestNewSidesValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"newSides", 8},
		{"newSides ", 8},
		{"newSides\n", 8},
		{"newSides.", 8},
		{"newSides:", 8},
		{"newSides,", 8},
		{"newSides\"", 8},
		{"newSides(", 8},
		{"newSides)", 8},
		{"newSides[", 8},
		{"newSides]", 8},
		{"newSides// comment", 8},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.NewSides()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestNewSidesInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_n"},
		{"n_"},
		{"_ne"},
		{"ne_"},
		{"_new"},
		{"new_"},
		{"_newS"},
		{"newS_"},
		{"_newSi"},
		{"newSi_"},
		{"_newSid"},
		{"newSid_"},
		{"_newSide"},
		{"newSide_"},
		{"_newSides"},
		{"newSides_"},
		{"9713n"},
		{"n9713"},
		{"9713ne"},
		{"ne9713"},
		{"9713new"},
		{"new9713"},
		{"9713newS"},
		{"newS9713"},
		{"9713newSi"},
		{"newSi9713"},
		{"9713newSid"},
		{"newSid9713"},
		{"9713newSide"},
		{"newSide9713"},
		{"9713newSides"},
		{"newSides9713"},
		{"Zn"},
		{"nZ"},
		{"Zne"},
		{"neZ"},
		{"Znew"},
		{"newZ"},
		{"ZnewS"},
		{"newSZ"},
		{"ZnewSi"},
		{"newSiZ"},
		{"ZnewSid"},
		{"newSidZ"},
		{"ZnewSide"},
		{"newSideZ"},
		{"ZnewSides"},
		{"newSidesZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.NewSides()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestMakeSidesValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"makeSides", 9},
		{"makeSides ", 9},
		{"makeSides\n", 9},
		{"makeSides.", 9},
		{"makeSides:", 9},
		{"makeSides,", 9},
		{"makeSides\"", 9},
		{"makeSides(", 9},
		{"makeSides)", 9},
		{"makeSides[", 9},
		{"makeSides]", 9},
		{"makeSides// comment", 9},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.MakeSides()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestMakeSidesInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_m"},
		{"m_"},
		{"_ma"},
		{"ma_"},
		{"_mak"},
		{"mak_"},
		{"_make"},
		{"make_"},
		{"_makeS"},
		{"makeS_"},
		{"_makeSi"},
		{"makeSi_"},
		{"_makeSid"},
		{"makeSid_"},
		{"_makeSide"},
		{"makeSide_"},
		{"_makeSides"},
		{"makeSides_"},
		{"9713m"},
		{"m9713"},
		{"9713ma"},
		{"ma9713"},
		{"9713mak"},
		{"mak9713"},
		{"9713make"},
		{"make9713"},
		{"9713makeS"},
		{"makeS9713"},
		{"9713makeSi"},
		{"makeSi9713"},
		{"9713makeSid"},
		{"makeSid9713"},
		{"9713makeSide"},
		{"makeSide9713"},
		{"9713makeSides"},
		{"makeSides9713"},
		{"Zm"},
		{"mZ"},
		{"Zma"},
		{"maZ"},
		{"Zmak"},
		{"makZ"},
		{"Zmake"},
		{"makeZ"},
		{"ZmakeS"},
		{"makeSZ"},
		{"ZmakeSi"},
		{"makeSiZ"},
		{"ZmakeSid"},
		{"makeSidZ"},
		{"ZmakeSide"},
		{"makeSideZ"},
		{"ZmakeSides"},
		{"makeSidesZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.MakeSides()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestMakeEmpSidesValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"makeEmpSides", 12},
		{"makeEmpSides ", 12},
		{"makeEmpSides\n", 12},
		{"makeEmpSides.", 12},
		{"makeEmpSides:", 12},
		{"makeEmpSides,", 12},
		{"makeEmpSides\"", 12},
		{"makeEmpSides(", 12},
		{"makeEmpSides)", 12},
		{"makeEmpSides[", 12},
		{"makeEmpSides]", 12},
		{"makeEmpSides// comment", 12},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.MakeEmpSides()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestMakeEmpSidesInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_m"},
		{"m_"},
		{"_ma"},
		{"ma_"},
		{"_mak"},
		{"mak_"},
		{"_make"},
		{"make_"},
		{"_makeE"},
		{"makeE_"},
		{"_makeEm"},
		{"makeEm_"},
		{"_makeEmp"},
		{"makeEmp_"},
		{"_makeEmpS"},
		{"makeEmpS_"},
		{"_makeEmpSi"},
		{"makeEmpSi_"},
		{"_makeEmpSid"},
		{"makeEmpSid_"},
		{"_makeEmpSide"},
		{"makeEmpSide_"},
		{"_makeEmpSides"},
		{"makeEmpSides_"},
		{"9713m"},
		{"m9713"},
		{"9713ma"},
		{"ma9713"},
		{"9713mak"},
		{"mak9713"},
		{"9713make"},
		{"make9713"},
		{"9713makeE"},
		{"makeE9713"},
		{"9713makeEm"},
		{"makeEm9713"},
		{"9713makeEmp"},
		{"makeEmp9713"},
		{"9713makeEmpS"},
		{"makeEmpS9713"},
		{"9713makeEmpSi"},
		{"makeEmpSi9713"},
		{"9713makeEmpSid"},
		{"makeEmpSid9713"},
		{"9713makeEmpSide"},
		{"makeEmpSide9713"},
		{"9713makeEmpSides"},
		{"makeEmpSides9713"},
		{"Zm"},
		{"mZ"},
		{"Zma"},
		{"maZ"},
		{"Zmak"},
		{"makZ"},
		{"Zmake"},
		{"makeZ"},
		{"ZmakeE"},
		{"makeEZ"},
		{"ZmakeEm"},
		{"makeEmZ"},
		{"ZmakeEmp"},
		{"makeEmpZ"},
		{"ZmakeEmpS"},
		{"makeEmpSZ"},
		{"ZmakeEmpSi"},
		{"makeEmpSiZ"},
		{"ZmakeEmpSid"},
		{"makeEmpSidZ"},
		{"ZmakeEmpSide"},
		{"makeEmpSideZ"},
		{"ZmakeEmpSides"},
		{"makeEmpSidesZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.MakeEmpSides()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestNewStmsValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"newStms", 7},
		{"newStms ", 7},
		{"newStms\n", 7},
		{"newStms.", 7},
		{"newStms:", 7},
		{"newStms,", 7},
		{"newStms\"", 7},
		{"newStms(", 7},
		{"newStms)", 7},
		{"newStms[", 7},
		{"newStms]", 7},
		{"newStms// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.NewStms()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestNewStmsInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_n"},
		{"n_"},
		{"_ne"},
		{"ne_"},
		{"_new"},
		{"new_"},
		{"_newS"},
		{"newS_"},
		{"_newSt"},
		{"newSt_"},
		{"_newStm"},
		{"newStm_"},
		{"_newStms"},
		{"newStms_"},
		{"9713n"},
		{"n9713"},
		{"9713ne"},
		{"ne9713"},
		{"9713new"},
		{"new9713"},
		{"9713newS"},
		{"newS9713"},
		{"9713newSt"},
		{"newSt9713"},
		{"9713newStm"},
		{"newStm9713"},
		{"9713newStms"},
		{"newStms9713"},
		{"Zn"},
		{"nZ"},
		{"Zne"},
		{"neZ"},
		{"Znew"},
		{"newZ"},
		{"ZnewS"},
		{"newSZ"},
		{"ZnewSt"},
		{"newStZ"},
		{"ZnewStm"},
		{"newStmZ"},
		{"ZnewStms"},
		{"newStmsZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.NewStms()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestMakeStmsValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"makeStms", 8},
		{"makeStms ", 8},
		{"makeStms\n", 8},
		{"makeStms.", 8},
		{"makeStms:", 8},
		{"makeStms,", 8},
		{"makeStms\"", 8},
		{"makeStms(", 8},
		{"makeStms)", 8},
		{"makeStms[", 8},
		{"makeStms]", 8},
		{"makeStms// comment", 8},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.MakeStms()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestMakeStmsInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_m"},
		{"m_"},
		{"_ma"},
		{"ma_"},
		{"_mak"},
		{"mak_"},
		{"_make"},
		{"make_"},
		{"_makeS"},
		{"makeS_"},
		{"_makeSt"},
		{"makeSt_"},
		{"_makeStm"},
		{"makeStm_"},
		{"_makeStms"},
		{"makeStms_"},
		{"9713m"},
		{"m9713"},
		{"9713ma"},
		{"ma9713"},
		{"9713mak"},
		{"mak9713"},
		{"9713make"},
		{"make9713"},
		{"9713makeS"},
		{"makeS9713"},
		{"9713makeSt"},
		{"makeSt9713"},
		{"9713makeStm"},
		{"makeStm9713"},
		{"9713makeStms"},
		{"makeStms9713"},
		{"Zm"},
		{"mZ"},
		{"Zma"},
		{"maZ"},
		{"Zmak"},
		{"makZ"},
		{"Zmake"},
		{"makeZ"},
		{"ZmakeS"},
		{"makeSZ"},
		{"ZmakeSt"},
		{"makeStZ"},
		{"ZmakeStm"},
		{"makeStmZ"},
		{"ZmakeStms"},
		{"makeStmsZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.MakeStms()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestMakeEmpStmsValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"makeEmpStms", 11},
		{"makeEmpStms ", 11},
		{"makeEmpStms\n", 11},
		{"makeEmpStms.", 11},
		{"makeEmpStms:", 11},
		{"makeEmpStms,", 11},
		{"makeEmpStms\"", 11},
		{"makeEmpStms(", 11},
		{"makeEmpStms)", 11},
		{"makeEmpStms[", 11},
		{"makeEmpStms]", 11},
		{"makeEmpStms// comment", 11},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.MakeEmpStms()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestMakeEmpStmsInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_m"},
		{"m_"},
		{"_ma"},
		{"ma_"},
		{"_mak"},
		{"mak_"},
		{"_make"},
		{"make_"},
		{"_makeE"},
		{"makeE_"},
		{"_makeEm"},
		{"makeEm_"},
		{"_makeEmp"},
		{"makeEmp_"},
		{"_makeEmpS"},
		{"makeEmpS_"},
		{"_makeEmpSt"},
		{"makeEmpSt_"},
		{"_makeEmpStm"},
		{"makeEmpStm_"},
		{"_makeEmpStms"},
		{"makeEmpStms_"},
		{"9713m"},
		{"m9713"},
		{"9713ma"},
		{"ma9713"},
		{"9713mak"},
		{"mak9713"},
		{"9713make"},
		{"make9713"},
		{"9713makeE"},
		{"makeE9713"},
		{"9713makeEm"},
		{"makeEm9713"},
		{"9713makeEmp"},
		{"makeEmp9713"},
		{"9713makeEmpS"},
		{"makeEmpS9713"},
		{"9713makeEmpSt"},
		{"makeEmpSt9713"},
		{"9713makeEmpStm"},
		{"makeEmpStm9713"},
		{"9713makeEmpStms"},
		{"makeEmpStms9713"},
		{"Zm"},
		{"mZ"},
		{"Zma"},
		{"maZ"},
		{"Zmak"},
		{"makZ"},
		{"Zmake"},
		{"makeZ"},
		{"ZmakeE"},
		{"makeEZ"},
		{"ZmakeEm"},
		{"makeEmZ"},
		{"ZmakeEmp"},
		{"makeEmpZ"},
		{"ZmakeEmpS"},
		{"makeEmpSZ"},
		{"ZmakeEmpSt"},
		{"makeEmpStZ"},
		{"ZmakeEmpStm"},
		{"makeEmpStmZ"},
		{"ZmakeEmpStms"},
		{"makeEmpStmsZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.MakeEmpStms()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestNewCndsValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"newCnds", 7},
		{"newCnds ", 7},
		{"newCnds\n", 7},
		{"newCnds.", 7},
		{"newCnds:", 7},
		{"newCnds,", 7},
		{"newCnds\"", 7},
		{"newCnds(", 7},
		{"newCnds)", 7},
		{"newCnds[", 7},
		{"newCnds]", 7},
		{"newCnds// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.NewCnds()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestNewCndsInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_n"},
		{"n_"},
		{"_ne"},
		{"ne_"},
		{"_new"},
		{"new_"},
		{"_newC"},
		{"newC_"},
		{"_newCn"},
		{"newCn_"},
		{"_newCnd"},
		{"newCnd_"},
		{"_newCnds"},
		{"newCnds_"},
		{"9713n"},
		{"n9713"},
		{"9713ne"},
		{"ne9713"},
		{"9713new"},
		{"new9713"},
		{"9713newC"},
		{"newC9713"},
		{"9713newCn"},
		{"newCn9713"},
		{"9713newCnd"},
		{"newCnd9713"},
		{"9713newCnds"},
		{"newCnds9713"},
		{"Zn"},
		{"nZ"},
		{"Zne"},
		{"neZ"},
		{"Znew"},
		{"newZ"},
		{"ZnewC"},
		{"newCZ"},
		{"ZnewCn"},
		{"newCnZ"},
		{"ZnewCnd"},
		{"newCndZ"},
		{"ZnewCnds"},
		{"newCndsZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.NewCnds()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestMakeCndsValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"makeCnds", 8},
		{"makeCnds ", 8},
		{"makeCnds\n", 8},
		{"makeCnds.", 8},
		{"makeCnds:", 8},
		{"makeCnds,", 8},
		{"makeCnds\"", 8},
		{"makeCnds(", 8},
		{"makeCnds)", 8},
		{"makeCnds[", 8},
		{"makeCnds]", 8},
		{"makeCnds// comment", 8},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.MakeCnds()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestMakeCndsInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_m"},
		{"m_"},
		{"_ma"},
		{"ma_"},
		{"_mak"},
		{"mak_"},
		{"_make"},
		{"make_"},
		{"_makeC"},
		{"makeC_"},
		{"_makeCn"},
		{"makeCn_"},
		{"_makeCnd"},
		{"makeCnd_"},
		{"_makeCnds"},
		{"makeCnds_"},
		{"9713m"},
		{"m9713"},
		{"9713ma"},
		{"ma9713"},
		{"9713mak"},
		{"mak9713"},
		{"9713make"},
		{"make9713"},
		{"9713makeC"},
		{"makeC9713"},
		{"9713makeCn"},
		{"makeCn9713"},
		{"9713makeCnd"},
		{"makeCnd9713"},
		{"9713makeCnds"},
		{"makeCnds9713"},
		{"Zm"},
		{"mZ"},
		{"Zma"},
		{"maZ"},
		{"Zmak"},
		{"makZ"},
		{"Zmake"},
		{"makeZ"},
		{"ZmakeC"},
		{"makeCZ"},
		{"ZmakeCn"},
		{"makeCnZ"},
		{"ZmakeCnd"},
		{"makeCndZ"},
		{"ZmakeCnds"},
		{"makeCndsZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.MakeCnds()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestMakeEmpCndsValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"makeEmpCnds", 11},
		{"makeEmpCnds ", 11},
		{"makeEmpCnds\n", 11},
		{"makeEmpCnds.", 11},
		{"makeEmpCnds:", 11},
		{"makeEmpCnds,", 11},
		{"makeEmpCnds\"", 11},
		{"makeEmpCnds(", 11},
		{"makeEmpCnds)", 11},
		{"makeEmpCnds[", 11},
		{"makeEmpCnds]", 11},
		{"makeEmpCnds// comment", 11},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.MakeEmpCnds()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestMakeEmpCndsInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_m"},
		{"m_"},
		{"_ma"},
		{"ma_"},
		{"_mak"},
		{"mak_"},
		{"_make"},
		{"make_"},
		{"_makeE"},
		{"makeE_"},
		{"_makeEm"},
		{"makeEm_"},
		{"_makeEmp"},
		{"makeEmp_"},
		{"_makeEmpC"},
		{"makeEmpC_"},
		{"_makeEmpCn"},
		{"makeEmpCn_"},
		{"_makeEmpCnd"},
		{"makeEmpCnd_"},
		{"_makeEmpCnds"},
		{"makeEmpCnds_"},
		{"9713m"},
		{"m9713"},
		{"9713ma"},
		{"ma9713"},
		{"9713mak"},
		{"mak9713"},
		{"9713make"},
		{"make9713"},
		{"9713makeE"},
		{"makeE9713"},
		{"9713makeEm"},
		{"makeEm9713"},
		{"9713makeEmp"},
		{"makeEmp9713"},
		{"9713makeEmpC"},
		{"makeEmpC9713"},
		{"9713makeEmpCn"},
		{"makeEmpCn9713"},
		{"9713makeEmpCnd"},
		{"makeEmpCnd9713"},
		{"9713makeEmpCnds"},
		{"makeEmpCnds9713"},
		{"Zm"},
		{"mZ"},
		{"Zma"},
		{"maZ"},
		{"Zmak"},
		{"makZ"},
		{"Zmake"},
		{"makeZ"},
		{"ZmakeE"},
		{"makeEZ"},
		{"ZmakeEm"},
		{"makeEmZ"},
		{"ZmakeEmp"},
		{"makeEmpZ"},
		{"ZmakeEmpC"},
		{"makeEmpCZ"},
		{"ZmakeEmpCn"},
		{"makeEmpCnZ"},
		{"ZmakeEmpCnd"},
		{"makeEmpCndZ"},
		{"ZmakeEmpCnds"},
		{"makeEmpCndsZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.MakeEmpCnds()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestNewStgysValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"newStgys", 8},
		{"newStgys ", 8},
		{"newStgys\n", 8},
		{"newStgys.", 8},
		{"newStgys:", 8},
		{"newStgys,", 8},
		{"newStgys\"", 8},
		{"newStgys(", 8},
		{"newStgys)", 8},
		{"newStgys[", 8},
		{"newStgys]", 8},
		{"newStgys// comment", 8},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.NewStgys()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestNewStgysInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_n"},
		{"n_"},
		{"_ne"},
		{"ne_"},
		{"_new"},
		{"new_"},
		{"_newS"},
		{"newS_"},
		{"_newSt"},
		{"newSt_"},
		{"_newStg"},
		{"newStg_"},
		{"_newStgy"},
		{"newStgy_"},
		{"_newStgys"},
		{"newStgys_"},
		{"9713n"},
		{"n9713"},
		{"9713ne"},
		{"ne9713"},
		{"9713new"},
		{"new9713"},
		{"9713newS"},
		{"newS9713"},
		{"9713newSt"},
		{"newSt9713"},
		{"9713newStg"},
		{"newStg9713"},
		{"9713newStgy"},
		{"newStgy9713"},
		{"9713newStgys"},
		{"newStgys9713"},
		{"Zn"},
		{"nZ"},
		{"Zne"},
		{"neZ"},
		{"Znew"},
		{"newZ"},
		{"ZnewS"},
		{"newSZ"},
		{"ZnewSt"},
		{"newStZ"},
		{"ZnewStg"},
		{"newStgZ"},
		{"ZnewStgy"},
		{"newStgyZ"},
		{"ZnewStgys"},
		{"newStgysZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.NewStgys()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestMakeStgysValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"makeStgys", 9},
		{"makeStgys ", 9},
		{"makeStgys\n", 9},
		{"makeStgys.", 9},
		{"makeStgys:", 9},
		{"makeStgys,", 9},
		{"makeStgys\"", 9},
		{"makeStgys(", 9},
		{"makeStgys)", 9},
		{"makeStgys[", 9},
		{"makeStgys]", 9},
		{"makeStgys// comment", 9},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.MakeStgys()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestMakeStgysInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_m"},
		{"m_"},
		{"_ma"},
		{"ma_"},
		{"_mak"},
		{"mak_"},
		{"_make"},
		{"make_"},
		{"_makeS"},
		{"makeS_"},
		{"_makeSt"},
		{"makeSt_"},
		{"_makeStg"},
		{"makeStg_"},
		{"_makeStgy"},
		{"makeStgy_"},
		{"_makeStgys"},
		{"makeStgys_"},
		{"9713m"},
		{"m9713"},
		{"9713ma"},
		{"ma9713"},
		{"9713mak"},
		{"mak9713"},
		{"9713make"},
		{"make9713"},
		{"9713makeS"},
		{"makeS9713"},
		{"9713makeSt"},
		{"makeSt9713"},
		{"9713makeStg"},
		{"makeStg9713"},
		{"9713makeStgy"},
		{"makeStgy9713"},
		{"9713makeStgys"},
		{"makeStgys9713"},
		{"Zm"},
		{"mZ"},
		{"Zma"},
		{"maZ"},
		{"Zmak"},
		{"makZ"},
		{"Zmake"},
		{"makeZ"},
		{"ZmakeS"},
		{"makeSZ"},
		{"ZmakeSt"},
		{"makeStZ"},
		{"ZmakeStg"},
		{"makeStgZ"},
		{"ZmakeStgy"},
		{"makeStgyZ"},
		{"ZmakeStgys"},
		{"makeStgysZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.MakeStgys()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestMakeEmpStgysValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"makeEmpStgys", 12},
		{"makeEmpStgys ", 12},
		{"makeEmpStgys\n", 12},
		{"makeEmpStgys.", 12},
		{"makeEmpStgys:", 12},
		{"makeEmpStgys,", 12},
		{"makeEmpStgys\"", 12},
		{"makeEmpStgys(", 12},
		{"makeEmpStgys)", 12},
		{"makeEmpStgys[", 12},
		{"makeEmpStgys]", 12},
		{"makeEmpStgys// comment", 12},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.MakeEmpStgys()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestMakeEmpStgysInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_m"},
		{"m_"},
		{"_ma"},
		{"ma_"},
		{"_mak"},
		{"mak_"},
		{"_make"},
		{"make_"},
		{"_makeE"},
		{"makeE_"},
		{"_makeEm"},
		{"makeEm_"},
		{"_makeEmp"},
		{"makeEmp_"},
		{"_makeEmpS"},
		{"makeEmpS_"},
		{"_makeEmpSt"},
		{"makeEmpSt_"},
		{"_makeEmpStg"},
		{"makeEmpStg_"},
		{"_makeEmpStgy"},
		{"makeEmpStgy_"},
		{"_makeEmpStgys"},
		{"makeEmpStgys_"},
		{"9713m"},
		{"m9713"},
		{"9713ma"},
		{"ma9713"},
		{"9713mak"},
		{"mak9713"},
		{"9713make"},
		{"make9713"},
		{"9713makeE"},
		{"makeE9713"},
		{"9713makeEm"},
		{"makeEm9713"},
		{"9713makeEmp"},
		{"makeEmp9713"},
		{"9713makeEmpS"},
		{"makeEmpS9713"},
		{"9713makeEmpSt"},
		{"makeEmpSt9713"},
		{"9713makeEmpStg"},
		{"makeEmpStg9713"},
		{"9713makeEmpStgy"},
		{"makeEmpStgy9713"},
		{"9713makeEmpStgys"},
		{"makeEmpStgys9713"},
		{"Zm"},
		{"mZ"},
		{"Zma"},
		{"maZ"},
		{"Zmak"},
		{"makZ"},
		{"Zmake"},
		{"makeZ"},
		{"ZmakeE"},
		{"makeEZ"},
		{"ZmakeEm"},
		{"makeEmZ"},
		{"ZmakeEmp"},
		{"makeEmpZ"},
		{"ZmakeEmpS"},
		{"makeEmpSZ"},
		{"ZmakeEmpSt"},
		{"makeEmpStZ"},
		{"ZmakeEmpStg"},
		{"makeEmpStgZ"},
		{"ZmakeEmpStgy"},
		{"makeEmpStgyZ"},
		{"ZmakeEmpStgys"},
		{"makeEmpStgysZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.MakeEmpStgys()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestRgbaValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"rgba", 4},
		{"rgba ", 4},
		{"rgba\n", 4},
		{"rgba.", 4},
		{"rgba:", 4},
		{"rgba,", 4},
		{"rgba\"", 4},
		{"rgba(", 4},
		{"rgba)", 4},
		{"rgba[", 4},
		{"rgba]", 4},
		{"rgba// comment", 4},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Rgba()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestRgbaInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_r"},
		{"r_"},
		{"_rg"},
		{"rg_"},
		{"_rgb"},
		{"rgb_"},
		{"_rgba"},
		{"rgba_"},
		{"9713r"},
		{"r9713"},
		{"9713rg"},
		{"rg9713"},
		{"9713rgb"},
		{"rgb9713"},
		{"9713rgba"},
		{"rgba9713"},
		{"Zr"},
		{"rZ"},
		{"Zrg"},
		{"rgZ"},
		{"Zrgb"},
		{"rgbZ"},
		{"Zrgba"},
		{"rgbaZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Rgba()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestRgbValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"rgb", 3},
		{"rgb ", 3},
		{"rgb\n", 3},
		{"rgb.", 3},
		{"rgb:", 3},
		{"rgb,", 3},
		{"rgb\"", 3},
		{"rgb(", 3},
		{"rgb)", 3},
		{"rgb[", 3},
		{"rgb]", 3},
		{"rgb// comment", 3},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Rgb()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestRgbInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_r"},
		{"r_"},
		{"_rg"},
		{"rg_"},
		{"_rgb"},
		{"rgb_"},
		{"9713r"},
		{"r9713"},
		{"9713rg"},
		{"rg9713"},
		{"9713rgb"},
		{"rgb9713"},
		{"Zr"},
		{"rZ"},
		{"Zrg"},
		{"rgZ"},
		{"Zrgb"},
		{"rgbZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Rgb()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestHexValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"hex", 3},
		{"hex ", 3},
		{"hex\n", 3},
		{"hex.", 3},
		{"hex:", 3},
		{"hex,", 3},
		{"hex\"", 3},
		{"hex(", 3},
		{"hex)", 3},
		{"hex[", 3},
		{"hex]", 3},
		{"hex// comment", 3},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Hex()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestHexInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_h"},
		{"h_"},
		{"_he"},
		{"he_"},
		{"_hex"},
		{"hex_"},
		{"9713h"},
		{"h9713"},
		{"9713he"},
		{"he9713"},
		{"9713hex"},
		{"hex9713"},
		{"Zh"},
		{"hZ"},
		{"Zhe"},
		{"heZ"},
		{"Zhex"},
		{"hexZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Hex()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestNewPensValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"newPens", 7},
		{"newPens ", 7},
		{"newPens\n", 7},
		{"newPens.", 7},
		{"newPens:", 7},
		{"newPens,", 7},
		{"newPens\"", 7},
		{"newPens(", 7},
		{"newPens)", 7},
		{"newPens[", 7},
		{"newPens]", 7},
		{"newPens// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.NewPens()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestNewPensInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_n"},
		{"n_"},
		{"_ne"},
		{"ne_"},
		{"_new"},
		{"new_"},
		{"_newP"},
		{"newP_"},
		{"_newPe"},
		{"newPe_"},
		{"_newPen"},
		{"newPen_"},
		{"_newPens"},
		{"newPens_"},
		{"9713n"},
		{"n9713"},
		{"9713ne"},
		{"ne9713"},
		{"9713new"},
		{"new9713"},
		{"9713newP"},
		{"newP9713"},
		{"9713newPe"},
		{"newPe9713"},
		{"9713newPen"},
		{"newPen9713"},
		{"9713newPens"},
		{"newPens9713"},
		{"Zn"},
		{"nZ"},
		{"Zne"},
		{"neZ"},
		{"Znew"},
		{"newZ"},
		{"ZnewP"},
		{"newPZ"},
		{"ZnewPe"},
		{"newPeZ"},
		{"ZnewPen"},
		{"newPenZ"},
		{"ZnewPens"},
		{"newPensZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.NewPens()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestMakePensValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"makePens", 8},
		{"makePens ", 8},
		{"makePens\n", 8},
		{"makePens.", 8},
		{"makePens:", 8},
		{"makePens,", 8},
		{"makePens\"", 8},
		{"makePens(", 8},
		{"makePens)", 8},
		{"makePens[", 8},
		{"makePens]", 8},
		{"makePens// comment", 8},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.MakePens()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestMakePensInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_m"},
		{"m_"},
		{"_ma"},
		{"ma_"},
		{"_mak"},
		{"mak_"},
		{"_make"},
		{"make_"},
		{"_makeP"},
		{"makeP_"},
		{"_makePe"},
		{"makePe_"},
		{"_makePen"},
		{"makePen_"},
		{"_makePens"},
		{"makePens_"},
		{"9713m"},
		{"m9713"},
		{"9713ma"},
		{"ma9713"},
		{"9713mak"},
		{"mak9713"},
		{"9713make"},
		{"make9713"},
		{"9713makeP"},
		{"makeP9713"},
		{"9713makePe"},
		{"makePe9713"},
		{"9713makePen"},
		{"makePen9713"},
		{"9713makePens"},
		{"makePens9713"},
		{"Zm"},
		{"mZ"},
		{"Zma"},
		{"maZ"},
		{"Zmak"},
		{"makZ"},
		{"Zmake"},
		{"makeZ"},
		{"ZmakeP"},
		{"makePZ"},
		{"ZmakePe"},
		{"makePeZ"},
		{"ZmakePen"},
		{"makePenZ"},
		{"ZmakePens"},
		{"makePensZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.MakePens()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestMakeEmpPensValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"makeEmpPens", 11},
		{"makeEmpPens ", 11},
		{"makeEmpPens\n", 11},
		{"makeEmpPens.", 11},
		{"makeEmpPens:", 11},
		{"makeEmpPens,", 11},
		{"makeEmpPens\"", 11},
		{"makeEmpPens(", 11},
		{"makeEmpPens)", 11},
		{"makeEmpPens[", 11},
		{"makeEmpPens]", 11},
		{"makeEmpPens// comment", 11},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.MakeEmpPens()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestMakeEmpPensInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_m"},
		{"m_"},
		{"_ma"},
		{"ma_"},
		{"_mak"},
		{"mak_"},
		{"_make"},
		{"make_"},
		{"_makeE"},
		{"makeE_"},
		{"_makeEm"},
		{"makeEm_"},
		{"_makeEmp"},
		{"makeEmp_"},
		{"_makeEmpP"},
		{"makeEmpP_"},
		{"_makeEmpPe"},
		{"makeEmpPe_"},
		{"_makeEmpPen"},
		{"makeEmpPen_"},
		{"_makeEmpPens"},
		{"makeEmpPens_"},
		{"9713m"},
		{"m9713"},
		{"9713ma"},
		{"ma9713"},
		{"9713mak"},
		{"mak9713"},
		{"9713make"},
		{"make9713"},
		{"9713makeE"},
		{"makeE9713"},
		{"9713makeEm"},
		{"makeEm9713"},
		{"9713makeEmp"},
		{"makeEmp9713"},
		{"9713makeEmpP"},
		{"makeEmpP9713"},
		{"9713makeEmpPe"},
		{"makeEmpPe9713"},
		{"9713makeEmpPen"},
		{"makeEmpPen9713"},
		{"9713makeEmpPens"},
		{"makeEmpPens9713"},
		{"Zm"},
		{"mZ"},
		{"Zma"},
		{"maZ"},
		{"Zmak"},
		{"makZ"},
		{"Zmake"},
		{"makeZ"},
		{"ZmakeE"},
		{"makeEZ"},
		{"ZmakeEm"},
		{"makeEmZ"},
		{"ZmakeEmp"},
		{"makeEmpZ"},
		{"ZmakeEmpP"},
		{"makeEmpPZ"},
		{"ZmakeEmpPe"},
		{"makeEmpPeZ"},
		{"ZmakeEmpPen"},
		{"makeEmpPenZ"},
		{"ZmakeEmpPens"},
		{"makeEmpPensZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.MakeEmpPens()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestNewPltsValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"newPlts", 7},
		{"newPlts ", 7},
		{"newPlts\n", 7},
		{"newPlts.", 7},
		{"newPlts:", 7},
		{"newPlts,", 7},
		{"newPlts\"", 7},
		{"newPlts(", 7},
		{"newPlts)", 7},
		{"newPlts[", 7},
		{"newPlts]", 7},
		{"newPlts// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.NewPlts()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestNewPltsInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_n"},
		{"n_"},
		{"_ne"},
		{"ne_"},
		{"_new"},
		{"new_"},
		{"_newP"},
		{"newP_"},
		{"_newPl"},
		{"newPl_"},
		{"_newPlt"},
		{"newPlt_"},
		{"_newPlts"},
		{"newPlts_"},
		{"9713n"},
		{"n9713"},
		{"9713ne"},
		{"ne9713"},
		{"9713new"},
		{"new9713"},
		{"9713newP"},
		{"newP9713"},
		{"9713newPl"},
		{"newPl9713"},
		{"9713newPlt"},
		{"newPlt9713"},
		{"9713newPlts"},
		{"newPlts9713"},
		{"Zn"},
		{"nZ"},
		{"Zne"},
		{"neZ"},
		{"Znew"},
		{"newZ"},
		{"ZnewP"},
		{"newPZ"},
		{"ZnewPl"},
		{"newPlZ"},
		{"ZnewPlt"},
		{"newPltZ"},
		{"ZnewPlts"},
		{"newPltsZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.NewPlts()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestMakePltsValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"makePlts", 8},
		{"makePlts ", 8},
		{"makePlts\n", 8},
		{"makePlts.", 8},
		{"makePlts:", 8},
		{"makePlts,", 8},
		{"makePlts\"", 8},
		{"makePlts(", 8},
		{"makePlts)", 8},
		{"makePlts[", 8},
		{"makePlts]", 8},
		{"makePlts// comment", 8},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.MakePlts()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestMakePltsInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_m"},
		{"m_"},
		{"_ma"},
		{"ma_"},
		{"_mak"},
		{"mak_"},
		{"_make"},
		{"make_"},
		{"_makeP"},
		{"makeP_"},
		{"_makePl"},
		{"makePl_"},
		{"_makePlt"},
		{"makePlt_"},
		{"_makePlts"},
		{"makePlts_"},
		{"9713m"},
		{"m9713"},
		{"9713ma"},
		{"ma9713"},
		{"9713mak"},
		{"mak9713"},
		{"9713make"},
		{"make9713"},
		{"9713makeP"},
		{"makeP9713"},
		{"9713makePl"},
		{"makePl9713"},
		{"9713makePlt"},
		{"makePlt9713"},
		{"9713makePlts"},
		{"makePlts9713"},
		{"Zm"},
		{"mZ"},
		{"Zma"},
		{"maZ"},
		{"Zmak"},
		{"makZ"},
		{"Zmake"},
		{"makeZ"},
		{"ZmakeP"},
		{"makePZ"},
		{"ZmakePl"},
		{"makePlZ"},
		{"ZmakePlt"},
		{"makePltZ"},
		{"ZmakePlts"},
		{"makePltsZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.MakePlts()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestMakeEmpPltsValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"makeEmpPlts", 11},
		{"makeEmpPlts ", 11},
		{"makeEmpPlts\n", 11},
		{"makeEmpPlts.", 11},
		{"makeEmpPlts:", 11},
		{"makeEmpPlts,", 11},
		{"makeEmpPlts\"", 11},
		{"makeEmpPlts(", 11},
		{"makeEmpPlts)", 11},
		{"makeEmpPlts[", 11},
		{"makeEmpPlts]", 11},
		{"makeEmpPlts// comment", 11},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.MakeEmpPlts()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestMakeEmpPltsInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_m"},
		{"m_"},
		{"_ma"},
		{"ma_"},
		{"_mak"},
		{"mak_"},
		{"_make"},
		{"make_"},
		{"_makeE"},
		{"makeE_"},
		{"_makeEm"},
		{"makeEm_"},
		{"_makeEmp"},
		{"makeEmp_"},
		{"_makeEmpP"},
		{"makeEmpP_"},
		{"_makeEmpPl"},
		{"makeEmpPl_"},
		{"_makeEmpPlt"},
		{"makeEmpPlt_"},
		{"_makeEmpPlts"},
		{"makeEmpPlts_"},
		{"9713m"},
		{"m9713"},
		{"9713ma"},
		{"ma9713"},
		{"9713mak"},
		{"mak9713"},
		{"9713make"},
		{"make9713"},
		{"9713makeE"},
		{"makeE9713"},
		{"9713makeEm"},
		{"makeEm9713"},
		{"9713makeEmp"},
		{"makeEmp9713"},
		{"9713makeEmpP"},
		{"makeEmpP9713"},
		{"9713makeEmpPl"},
		{"makeEmpPl9713"},
		{"9713makeEmpPlt"},
		{"makeEmpPlt9713"},
		{"9713makeEmpPlts"},
		{"makeEmpPlts9713"},
		{"Zm"},
		{"mZ"},
		{"Zma"},
		{"maZ"},
		{"Zmak"},
		{"makZ"},
		{"Zmake"},
		{"makeZ"},
		{"ZmakeE"},
		{"makeEZ"},
		{"ZmakeEm"},
		{"makeEmZ"},
		{"ZmakeEmp"},
		{"makeEmpZ"},
		{"ZmakeEmpP"},
		{"makeEmpPZ"},
		{"ZmakeEmpPl"},
		{"makeEmpPlZ"},
		{"ZmakeEmpPlt"},
		{"makeEmpPltZ"},
		{"ZmakeEmpPlts"},
		{"makeEmpPltsZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.MakeEmpPlts()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestNewStmValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"newStm", 6},
		{"newStm ", 6},
		{"newStm\n", 6},
		{"newStm.", 6},
		{"newStm:", 6},
		{"newStm,", 6},
		{"newStm\"", 6},
		{"newStm(", 6},
		{"newStm)", 6},
		{"newStm[", 6},
		{"newStm]", 6},
		{"newStm// comment", 6},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.NewStm()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestNewStmInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_n"},
		{"n_"},
		{"_ne"},
		{"ne_"},
		{"_new"},
		{"new_"},
		{"_newS"},
		{"newS_"},
		{"_newSt"},
		{"newSt_"},
		{"_newStm"},
		{"newStm_"},
		{"9713n"},
		{"n9713"},
		{"9713ne"},
		{"ne9713"},
		{"9713new"},
		{"new9713"},
		{"9713newS"},
		{"newS9713"},
		{"9713newSt"},
		{"newSt9713"},
		{"9713newStm"},
		{"newStm9713"},
		{"Zn"},
		{"nZ"},
		{"Zne"},
		{"neZ"},
		{"Znew"},
		{"newZ"},
		{"ZnewS"},
		{"newSZ"},
		{"ZnewSt"},
		{"newStZ"},
		{"ZnewStm"},
		{"newStmZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.NewStm()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestNewFltsSctrValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"newFltsSctr", 11},
		{"newFltsSctr ", 11},
		{"newFltsSctr\n", 11},
		{"newFltsSctr.", 11},
		{"newFltsSctr:", 11},
		{"newFltsSctr,", 11},
		{"newFltsSctr\"", 11},
		{"newFltsSctr(", 11},
		{"newFltsSctr)", 11},
		{"newFltsSctr[", 11},
		{"newFltsSctr]", 11},
		{"newFltsSctr// comment", 11},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.NewFltsSctr()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestNewFltsSctrInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_n"},
		{"n_"},
		{"_ne"},
		{"ne_"},
		{"_new"},
		{"new_"},
		{"_newF"},
		{"newF_"},
		{"_newFl"},
		{"newFl_"},
		{"_newFlt"},
		{"newFlt_"},
		{"_newFlts"},
		{"newFlts_"},
		{"_newFltsS"},
		{"newFltsS_"},
		{"_newFltsSc"},
		{"newFltsSc_"},
		{"_newFltsSct"},
		{"newFltsSct_"},
		{"_newFltsSctr"},
		{"newFltsSctr_"},
		{"9713n"},
		{"n9713"},
		{"9713ne"},
		{"ne9713"},
		{"9713new"},
		{"new9713"},
		{"9713newF"},
		{"newF9713"},
		{"9713newFl"},
		{"newFl9713"},
		{"9713newFlt"},
		{"newFlt9713"},
		{"9713newFlts"},
		{"newFlts9713"},
		{"9713newFltsS"},
		{"newFltsS9713"},
		{"9713newFltsSc"},
		{"newFltsSc9713"},
		{"9713newFltsSct"},
		{"newFltsSct9713"},
		{"9713newFltsSctr"},
		{"newFltsSctr9713"},
		{"Zn"},
		{"nZ"},
		{"Zne"},
		{"neZ"},
		{"Znew"},
		{"newZ"},
		{"ZnewF"},
		{"newFZ"},
		{"ZnewFl"},
		{"newFlZ"},
		{"ZnewFlt"},
		{"newFltZ"},
		{"ZnewFlts"},
		{"newFltsZ"},
		{"ZnewFltsS"},
		{"newFltsSZ"},
		{"ZnewFltsSc"},
		{"newFltsScZ"},
		{"ZnewFltsSct"},
		{"newFltsSctZ"},
		{"ZnewFltsSctr"},
		{"newFltsSctrZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.NewFltsSctr()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestNewFltsSctrDistValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"newFltsSctrDist", 15},
		{"newFltsSctrDist ", 15},
		{"newFltsSctrDist\n", 15},
		{"newFltsSctrDist.", 15},
		{"newFltsSctrDist:", 15},
		{"newFltsSctrDist,", 15},
		{"newFltsSctrDist\"", 15},
		{"newFltsSctrDist(", 15},
		{"newFltsSctrDist)", 15},
		{"newFltsSctrDist[", 15},
		{"newFltsSctrDist]", 15},
		{"newFltsSctrDist// comment", 15},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.NewFltsSctrDist()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestNewFltsSctrDistInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_n"},
		{"n_"},
		{"_ne"},
		{"ne_"},
		{"_new"},
		{"new_"},
		{"_newF"},
		{"newF_"},
		{"_newFl"},
		{"newFl_"},
		{"_newFlt"},
		{"newFlt_"},
		{"_newFlts"},
		{"newFlts_"},
		{"_newFltsS"},
		{"newFltsS_"},
		{"_newFltsSc"},
		{"newFltsSc_"},
		{"_newFltsSct"},
		{"newFltsSct_"},
		{"_newFltsSctr"},
		{"newFltsSctr_"},
		{"_newFltsSctrD"},
		{"newFltsSctrD_"},
		{"_newFltsSctrDi"},
		{"newFltsSctrDi_"},
		{"_newFltsSctrDis"},
		{"newFltsSctrDis_"},
		{"_newFltsSctrDist"},
		{"newFltsSctrDist_"},
		{"9713n"},
		{"n9713"},
		{"9713ne"},
		{"ne9713"},
		{"9713new"},
		{"new9713"},
		{"9713newF"},
		{"newF9713"},
		{"9713newFl"},
		{"newFl9713"},
		{"9713newFlt"},
		{"newFlt9713"},
		{"9713newFlts"},
		{"newFlts9713"},
		{"9713newFltsS"},
		{"newFltsS9713"},
		{"9713newFltsSc"},
		{"newFltsSc9713"},
		{"9713newFltsSct"},
		{"newFltsSct9713"},
		{"9713newFltsSctr"},
		{"newFltsSctr9713"},
		{"9713newFltsSctrD"},
		{"newFltsSctrD9713"},
		{"9713newFltsSctrDi"},
		{"newFltsSctrDi9713"},
		{"9713newFltsSctrDis"},
		{"newFltsSctrDis9713"},
		{"9713newFltsSctrDist"},
		{"newFltsSctrDist9713"},
		{"Zn"},
		{"nZ"},
		{"Zne"},
		{"neZ"},
		{"Znew"},
		{"newZ"},
		{"ZnewF"},
		{"newFZ"},
		{"ZnewFl"},
		{"newFlZ"},
		{"ZnewFlt"},
		{"newFltZ"},
		{"ZnewFlts"},
		{"newFltsZ"},
		{"ZnewFltsS"},
		{"newFltsSZ"},
		{"ZnewFltsSc"},
		{"newFltsScZ"},
		{"ZnewFltsSct"},
		{"newFltsSctZ"},
		{"ZnewFltsSctr"},
		{"newFltsSctrZ"},
		{"ZnewFltsSctrD"},
		{"newFltsSctrDZ"},
		{"ZnewFltsSctrDi"},
		{"newFltsSctrDiZ"},
		{"ZnewFltsSctrDis"},
		{"newFltsSctrDisZ"},
		{"ZnewFltsSctrDist"},
		{"newFltsSctrDistZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.NewFltsSctrDist()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestNewHrzValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"newHrz", 6},
		{"newHrz ", 6},
		{"newHrz\n", 6},
		{"newHrz.", 6},
		{"newHrz:", 6},
		{"newHrz,", 6},
		{"newHrz\"", 6},
		{"newHrz(", 6},
		{"newHrz)", 6},
		{"newHrz[", 6},
		{"newHrz]", 6},
		{"newHrz// comment", 6},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.NewHrz()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestNewHrzInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_n"},
		{"n_"},
		{"_ne"},
		{"ne_"},
		{"_new"},
		{"new_"},
		{"_newH"},
		{"newH_"},
		{"_newHr"},
		{"newHr_"},
		{"_newHrz"},
		{"newHrz_"},
		{"9713n"},
		{"n9713"},
		{"9713ne"},
		{"ne9713"},
		{"9713new"},
		{"new9713"},
		{"9713newH"},
		{"newH9713"},
		{"9713newHr"},
		{"newHr9713"},
		{"9713newHrz"},
		{"newHrz9713"},
		{"Zn"},
		{"nZ"},
		{"Zne"},
		{"neZ"},
		{"Znew"},
		{"newZ"},
		{"ZnewH"},
		{"newHZ"},
		{"ZnewHr"},
		{"newHrZ"},
		{"ZnewHrz"},
		{"newHrzZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.NewHrz()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestNewVrtValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"newVrt", 6},
		{"newVrt ", 6},
		{"newVrt\n", 6},
		{"newVrt.", 6},
		{"newVrt:", 6},
		{"newVrt,", 6},
		{"newVrt\"", 6},
		{"newVrt(", 6},
		{"newVrt)", 6},
		{"newVrt[", 6},
		{"newVrt]", 6},
		{"newVrt// comment", 6},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.NewVrt()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestNewVrtInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_n"},
		{"n_"},
		{"_ne"},
		{"ne_"},
		{"_new"},
		{"new_"},
		{"_newV"},
		{"newV_"},
		{"_newVr"},
		{"newVr_"},
		{"_newVrt"},
		{"newVrt_"},
		{"9713n"},
		{"n9713"},
		{"9713ne"},
		{"ne9713"},
		{"9713new"},
		{"new9713"},
		{"9713newV"},
		{"newV9713"},
		{"9713newVr"},
		{"newVr9713"},
		{"9713newVrt"},
		{"newVrt9713"},
		{"Zn"},
		{"nZ"},
		{"Zne"},
		{"neZ"},
		{"Znew"},
		{"newZ"},
		{"ZnewV"},
		{"newVZ"},
		{"ZnewVr"},
		{"newVrZ"},
		{"ZnewVrt"},
		{"newVrtZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.NewVrt()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestNewDpthValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"newDpth", 7},
		{"newDpth ", 7},
		{"newDpth\n", 7},
		{"newDpth.", 7},
		{"newDpth:", 7},
		{"newDpth,", 7},
		{"newDpth\"", 7},
		{"newDpth(", 7},
		{"newDpth)", 7},
		{"newDpth[", 7},
		{"newDpth]", 7},
		{"newDpth// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.NewDpth()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestNewDpthInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_n"},
		{"n_"},
		{"_ne"},
		{"ne_"},
		{"_new"},
		{"new_"},
		{"_newD"},
		{"newD_"},
		{"_newDp"},
		{"newDp_"},
		{"_newDpt"},
		{"newDpt_"},
		{"_newDpth"},
		{"newDpth_"},
		{"9713n"},
		{"n9713"},
		{"9713ne"},
		{"ne9713"},
		{"9713new"},
		{"new9713"},
		{"9713newD"},
		{"newD9713"},
		{"9713newDp"},
		{"newDp9713"},
		{"9713newDpt"},
		{"newDpt9713"},
		{"9713newDpth"},
		{"newDpth9713"},
		{"Zn"},
		{"nZ"},
		{"Zne"},
		{"neZ"},
		{"Znew"},
		{"newZ"},
		{"ZnewD"},
		{"newDZ"},
		{"ZnewDp"},
		{"newDpZ"},
		{"ZnewDpt"},
		{"newDptZ"},
		{"ZnewDpth"},
		{"newDpthZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.NewDpth()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestNewMuValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"newMu", 5},
		{"newMu ", 5},
		{"newMu\n", 5},
		{"newMu.", 5},
		{"newMu:", 5},
		{"newMu,", 5},
		{"newMu\"", 5},
		{"newMu(", 5},
		{"newMu)", 5},
		{"newMu[", 5},
		{"newMu]", 5},
		{"newMu// comment", 5},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.NewMu()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestNewMuInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_n"},
		{"n_"},
		{"_ne"},
		{"ne_"},
		{"_new"},
		{"new_"},
		{"_newM"},
		{"newM_"},
		{"_newMu"},
		{"newMu_"},
		{"9713n"},
		{"n9713"},
		{"9713ne"},
		{"ne9713"},
		{"9713new"},
		{"new9713"},
		{"9713newM"},
		{"newM9713"},
		{"9713newMu"},
		{"newMu9713"},
		{"Zn"},
		{"nZ"},
		{"Zne"},
		{"neZ"},
		{"Znew"},
		{"newZ"},
		{"ZnewM"},
		{"newMZ"},
		{"ZnewMu"},
		{"newMuZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.NewMu()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestLowerValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"lower", 5},
		{"lower ", 5},
		{"lower\n", 5},
		{"lower.", 5},
		{"lower:", 5},
		{"lower,", 5},
		{"lower\"", 5},
		{"lower(", 5},
		{"lower)", 5},
		{"lower[", 5},
		{"lower]", 5},
		{"lower// comment", 5},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Lower()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestLowerInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_l"},
		{"l_"},
		{"_lo"},
		{"lo_"},
		{"_low"},
		{"low_"},
		{"_lowe"},
		{"lowe_"},
		{"_lower"},
		{"lower_"},
		{"9713l"},
		{"l9713"},
		{"9713lo"},
		{"lo9713"},
		{"9713low"},
		{"low9713"},
		{"9713lowe"},
		{"lowe9713"},
		{"9713lower"},
		{"lower9713"},
		{"Zl"},
		{"lZ"},
		{"Zlo"},
		{"loZ"},
		{"Zlow"},
		{"lowZ"},
		{"Zlowe"},
		{"loweZ"},
		{"Zlower"},
		{"lowerZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Lower()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestUpperValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"upper", 5},
		{"upper ", 5},
		{"upper\n", 5},
		{"upper.", 5},
		{"upper:", 5},
		{"upper,", 5},
		{"upper\"", 5},
		{"upper(", 5},
		{"upper)", 5},
		{"upper[", 5},
		{"upper]", 5},
		{"upper// comment", 5},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Upper()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestUpperInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_u"},
		{"u_"},
		{"_up"},
		{"up_"},
		{"_upp"},
		{"upp_"},
		{"_uppe"},
		{"uppe_"},
		{"_upper"},
		{"upper_"},
		{"9713u"},
		{"u9713"},
		{"9713up"},
		{"up9713"},
		{"9713upp"},
		{"upp9713"},
		{"9713uppe"},
		{"uppe9713"},
		{"9713upper"},
		{"upper9713"},
		{"Zu"},
		{"uZ"},
		{"Zup"},
		{"upZ"},
		{"Zupp"},
		{"uppZ"},
		{"Zuppe"},
		{"uppeZ"},
		{"Zupper"},
		{"upperZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Upper()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestEqlValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"eql", 3},
		{"eql ", 3},
		{"eql\n", 3},
		{"eql.", 3},
		{"eql:", 3},
		{"eql,", 3},
		{"eql\"", 3},
		{"eql(", 3},
		{"eql)", 3},
		{"eql[", 3},
		{"eql]", 3},
		{"eql// comment", 3},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Eql()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestEqlInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_e"},
		{"e_"},
		{"_eq"},
		{"eq_"},
		{"_eql"},
		{"eql_"},
		{"9713e"},
		{"e9713"},
		{"9713eq"},
		{"eq9713"},
		{"9713eql"},
		{"eql9713"},
		{"Ze"},
		{"eZ"},
		{"Zeq"},
		{"eqZ"},
		{"Zeql"},
		{"eqlZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Eql()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestNeqValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"neq", 3},
		{"neq ", 3},
		{"neq\n", 3},
		{"neq.", 3},
		{"neq:", 3},
		{"neq,", 3},
		{"neq\"", 3},
		{"neq(", 3},
		{"neq)", 3},
		{"neq[", 3},
		{"neq]", 3},
		{"neq// comment", 3},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Neq()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestNeqInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_n"},
		{"n_"},
		{"_ne"},
		{"ne_"},
		{"_neq"},
		{"neq_"},
		{"9713n"},
		{"n9713"},
		{"9713ne"},
		{"ne9713"},
		{"9713neq"},
		{"neq9713"},
		{"Zn"},
		{"nZ"},
		{"Zne"},
		{"neZ"},
		{"Zneq"},
		{"neqZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Neq()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestLssValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"lss", 3},
		{"lss ", 3},
		{"lss\n", 3},
		{"lss.", 3},
		{"lss:", 3},
		{"lss,", 3},
		{"lss\"", 3},
		{"lss(", 3},
		{"lss)", 3},
		{"lss[", 3},
		{"lss]", 3},
		{"lss// comment", 3},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Lss()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestLssInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_l"},
		{"l_"},
		{"_ls"},
		{"ls_"},
		{"_lss"},
		{"lss_"},
		{"9713l"},
		{"l9713"},
		{"9713ls"},
		{"ls9713"},
		{"9713lss"},
		{"lss9713"},
		{"Zl"},
		{"lZ"},
		{"Zls"},
		{"lsZ"},
		{"Zlss"},
		{"lssZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Lss()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestGtrValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"gtr", 3},
		{"gtr ", 3},
		{"gtr\n", 3},
		{"gtr.", 3},
		{"gtr:", 3},
		{"gtr,", 3},
		{"gtr\"", 3},
		{"gtr(", 3},
		{"gtr)", 3},
		{"gtr[", 3},
		{"gtr]", 3},
		{"gtr// comment", 3},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Gtr()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestGtrInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_g"},
		{"g_"},
		{"_gt"},
		{"gt_"},
		{"_gtr"},
		{"gtr_"},
		{"9713g"},
		{"g9713"},
		{"9713gt"},
		{"gt9713"},
		{"9713gtr"},
		{"gtr9713"},
		{"Zg"},
		{"gZ"},
		{"Zgt"},
		{"gtZ"},
		{"Zgtr"},
		{"gtrZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Gtr()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestLeqValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"leq", 3},
		{"leq ", 3},
		{"leq\n", 3},
		{"leq.", 3},
		{"leq:", 3},
		{"leq,", 3},
		{"leq\"", 3},
		{"leq(", 3},
		{"leq)", 3},
		{"leq[", 3},
		{"leq]", 3},
		{"leq// comment", 3},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Leq()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestLeqInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_l"},
		{"l_"},
		{"_le"},
		{"le_"},
		{"_leq"},
		{"leq_"},
		{"9713l"},
		{"l9713"},
		{"9713le"},
		{"le9713"},
		{"9713leq"},
		{"leq9713"},
		{"Zl"},
		{"lZ"},
		{"Zle"},
		{"leZ"},
		{"Zleq"},
		{"leqZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Leq()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestGeqValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"geq", 3},
		{"geq ", 3},
		{"geq\n", 3},
		{"geq.", 3},
		{"geq:", 3},
		{"geq,", 3},
		{"geq\"", 3},
		{"geq(", 3},
		{"geq)", 3},
		{"geq[", 3},
		{"geq]", 3},
		{"geq// comment", 3},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Geq()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestGeqInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_g"},
		{"g_"},
		{"_ge"},
		{"ge_"},
		{"_geq"},
		{"geq_"},
		{"9713g"},
		{"g9713"},
		{"9713ge"},
		{"ge9713"},
		{"9713geq"},
		{"geq9713"},
		{"Zg"},
		{"gZ"},
		{"Zge"},
		{"geZ"},
		{"Zgeq"},
		{"geqZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Geq()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestNotValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"not", 3},
		{"not ", 3},
		{"not\n", 3},
		{"not.", 3},
		{"not:", 3},
		{"not,", 3},
		{"not\"", 3},
		{"not(", 3},
		{"not)", 3},
		{"not[", 3},
		{"not]", 3},
		{"not// comment", 3},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Not()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestNotInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_n"},
		{"n_"},
		{"_no"},
		{"no_"},
		{"_not"},
		{"not_"},
		{"9713n"},
		{"n9713"},
		{"9713no"},
		{"no9713"},
		{"9713not"},
		{"not9713"},
		{"Zn"},
		{"nZ"},
		{"Zno"},
		{"noZ"},
		{"Znot"},
		{"notZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Not()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestTrncValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"trnc", 4},
		{"trnc ", 4},
		{"trnc\n", 4},
		{"trnc.", 4},
		{"trnc:", 4},
		{"trnc,", 4},
		{"trnc\"", 4},
		{"trnc(", 4},
		{"trnc)", 4},
		{"trnc[", 4},
		{"trnc]", 4},
		{"trnc// comment", 4},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Trnc()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestTrncInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_t"},
		{"t_"},
		{"_tr"},
		{"tr_"},
		{"_trn"},
		{"trn_"},
		{"_trnc"},
		{"trnc_"},
		{"9713t"},
		{"t9713"},
		{"9713tr"},
		{"tr9713"},
		{"9713trn"},
		{"trn9713"},
		{"9713trnc"},
		{"trnc9713"},
		{"Zt"},
		{"tZ"},
		{"Ztr"},
		{"trZ"},
		{"Ztrn"},
		{"trnZ"},
		{"Ztrnc"},
		{"trncZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Trnc()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestIsNaNValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"isNaN", 5},
		{"isNaN ", 5},
		{"isNaN\n", 5},
		{"isNaN.", 5},
		{"isNaN:", 5},
		{"isNaN,", 5},
		{"isNaN\"", 5},
		{"isNaN(", 5},
		{"isNaN)", 5},
		{"isNaN[", 5},
		{"isNaN]", 5},
		{"isNaN// comment", 5},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.IsNaN()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestIsNaNInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_i"},
		{"i_"},
		{"_is"},
		{"is_"},
		{"_isN"},
		{"isN_"},
		{"_isNa"},
		{"isNa_"},
		{"_isNaN"},
		{"isNaN_"},
		{"9713i"},
		{"i9713"},
		{"9713is"},
		{"is9713"},
		{"9713isN"},
		{"isN9713"},
		{"9713isNa"},
		{"isNa9713"},
		{"9713isNaN"},
		{"isNaN9713"},
		{"Zi"},
		{"iZ"},
		{"Zis"},
		{"isZ"},
		{"ZisN"},
		{"isNZ"},
		{"ZisNa"},
		{"isNaZ"},
		{"ZisNaN"},
		{"isNaNZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.IsNaN()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestIsInfPosValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"isInfPos", 8},
		{"isInfPos ", 8},
		{"isInfPos\n", 8},
		{"isInfPos.", 8},
		{"isInfPos:", 8},
		{"isInfPos,", 8},
		{"isInfPos\"", 8},
		{"isInfPos(", 8},
		{"isInfPos)", 8},
		{"isInfPos[", 8},
		{"isInfPos]", 8},
		{"isInfPos// comment", 8},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.IsInfPos()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestIsInfPosInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_i"},
		{"i_"},
		{"_is"},
		{"is_"},
		{"_isI"},
		{"isI_"},
		{"_isIn"},
		{"isIn_"},
		{"_isInf"},
		{"isInf_"},
		{"_isInfP"},
		{"isInfP_"},
		{"_isInfPo"},
		{"isInfPo_"},
		{"_isInfPos"},
		{"isInfPos_"},
		{"9713i"},
		{"i9713"},
		{"9713is"},
		{"is9713"},
		{"9713isI"},
		{"isI9713"},
		{"9713isIn"},
		{"isIn9713"},
		{"9713isInf"},
		{"isInf9713"},
		{"9713isInfP"},
		{"isInfP9713"},
		{"9713isInfPo"},
		{"isInfPo9713"},
		{"9713isInfPos"},
		{"isInfPos9713"},
		{"Zi"},
		{"iZ"},
		{"Zis"},
		{"isZ"},
		{"ZisI"},
		{"isIZ"},
		{"ZisIn"},
		{"isInZ"},
		{"ZisInf"},
		{"isInfZ"},
		{"ZisInfP"},
		{"isInfPZ"},
		{"ZisInfPo"},
		{"isInfPoZ"},
		{"ZisInfPos"},
		{"isInfPosZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.IsInfPos()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestIsInfNegValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"isInfNeg", 8},
		{"isInfNeg ", 8},
		{"isInfNeg\n", 8},
		{"isInfNeg.", 8},
		{"isInfNeg:", 8},
		{"isInfNeg,", 8},
		{"isInfNeg\"", 8},
		{"isInfNeg(", 8},
		{"isInfNeg)", 8},
		{"isInfNeg[", 8},
		{"isInfNeg]", 8},
		{"isInfNeg// comment", 8},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.IsInfNeg()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestIsInfNegInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_i"},
		{"i_"},
		{"_is"},
		{"is_"},
		{"_isI"},
		{"isI_"},
		{"_isIn"},
		{"isIn_"},
		{"_isInf"},
		{"isInf_"},
		{"_isInfN"},
		{"isInfN_"},
		{"_isInfNe"},
		{"isInfNe_"},
		{"_isInfNeg"},
		{"isInfNeg_"},
		{"9713i"},
		{"i9713"},
		{"9713is"},
		{"is9713"},
		{"9713isI"},
		{"isI9713"},
		{"9713isIn"},
		{"isIn9713"},
		{"9713isInf"},
		{"isInf9713"},
		{"9713isInfN"},
		{"isInfN9713"},
		{"9713isInfNe"},
		{"isInfNe9713"},
		{"9713isInfNeg"},
		{"isInfNeg9713"},
		{"Zi"},
		{"iZ"},
		{"Zis"},
		{"isZ"},
		{"ZisI"},
		{"isIZ"},
		{"ZisIn"},
		{"isInZ"},
		{"ZisInf"},
		{"isInfZ"},
		{"ZisInfN"},
		{"isInfNZ"},
		{"ZisInfNe"},
		{"isInfNeZ"},
		{"ZisInfNeg"},
		{"isInfNegZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.IsInfNeg()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestIsValidValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"isValid", 7},
		{"isValid ", 7},
		{"isValid\n", 7},
		{"isValid.", 7},
		{"isValid:", 7},
		{"isValid,", 7},
		{"isValid\"", 7},
		{"isValid(", 7},
		{"isValid)", 7},
		{"isValid[", 7},
		{"isValid]", 7},
		{"isValid// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.IsValid()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestIsValidInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_i"},
		{"i_"},
		{"_is"},
		{"is_"},
		{"_isV"},
		{"isV_"},
		{"_isVa"},
		{"isVa_"},
		{"_isVal"},
		{"isVal_"},
		{"_isVali"},
		{"isVali_"},
		{"_isValid"},
		{"isValid_"},
		{"9713i"},
		{"i9713"},
		{"9713is"},
		{"is9713"},
		{"9713isV"},
		{"isV9713"},
		{"9713isVa"},
		{"isVa9713"},
		{"9713isVal"},
		{"isVal9713"},
		{"9713isVali"},
		{"isVali9713"},
		{"9713isValid"},
		{"isValid9713"},
		{"Zi"},
		{"iZ"},
		{"Zis"},
		{"isZ"},
		{"ZisV"},
		{"isVZ"},
		{"ZisVa"},
		{"isVaZ"},
		{"ZisVal"},
		{"isValZ"},
		{"ZisVali"},
		{"isValiZ"},
		{"ZisValid"},
		{"isValidZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.IsValid()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestPctValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"pct", 3},
		{"pct ", 3},
		{"pct\n", 3},
		{"pct.", 3},
		{"pct:", 3},
		{"pct,", 3},
		{"pct\"", 3},
		{"pct(", 3},
		{"pct)", 3},
		{"pct[", 3},
		{"pct]", 3},
		{"pct// comment", 3},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Pct()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestPctInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_p"},
		{"p_"},
		{"_pc"},
		{"pc_"},
		{"_pct"},
		{"pct_"},
		{"9713p"},
		{"p9713"},
		{"9713pc"},
		{"pc9713"},
		{"9713pct"},
		{"pct9713"},
		{"Zp"},
		{"pZ"},
		{"Zpc"},
		{"pcZ"},
		{"Zpct"},
		{"pctZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Pct()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestPosValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"pos", 3},
		{"pos ", 3},
		{"pos\n", 3},
		{"pos.", 3},
		{"pos:", 3},
		{"pos,", 3},
		{"pos\"", 3},
		{"pos(", 3},
		{"pos)", 3},
		{"pos[", 3},
		{"pos]", 3},
		{"pos// comment", 3},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Pos()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestPosInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_p"},
		{"p_"},
		{"_po"},
		{"po_"},
		{"_pos"},
		{"pos_"},
		{"9713p"},
		{"p9713"},
		{"9713po"},
		{"po9713"},
		{"9713pos"},
		{"pos9713"},
		{"Zp"},
		{"pZ"},
		{"Zpo"},
		{"poZ"},
		{"Zpos"},
		{"posZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Pos()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestNegValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"neg", 3},
		{"neg ", 3},
		{"neg\n", 3},
		{"neg.", 3},
		{"neg:", 3},
		{"neg,", 3},
		{"neg\"", 3},
		{"neg(", 3},
		{"neg)", 3},
		{"neg[", 3},
		{"neg]", 3},
		{"neg// comment", 3},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Neg()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestNegInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_n"},
		{"n_"},
		{"_ne"},
		{"ne_"},
		{"_neg"},
		{"neg_"},
		{"9713n"},
		{"n9713"},
		{"9713ne"},
		{"ne9713"},
		{"9713neg"},
		{"neg9713"},
		{"Zn"},
		{"nZ"},
		{"Zne"},
		{"neZ"},
		{"Zneg"},
		{"negZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Neg()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestInvValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"inv", 3},
		{"inv ", 3},
		{"inv\n", 3},
		{"inv.", 3},
		{"inv:", 3},
		{"inv,", 3},
		{"inv\"", 3},
		{"inv(", 3},
		{"inv)", 3},
		{"inv[", 3},
		{"inv]", 3},
		{"inv// comment", 3},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Inv()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestInvInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_i"},
		{"i_"},
		{"_in"},
		{"in_"},
		{"_inv"},
		{"inv_"},
		{"9713i"},
		{"i9713"},
		{"9713in"},
		{"in9713"},
		{"9713inv"},
		{"inv9713"},
		{"Zi"},
		{"iZ"},
		{"Zin"},
		{"inZ"},
		{"Zinv"},
		{"invZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Inv()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestAddValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"add", 3},
		{"add ", 3},
		{"add\n", 3},
		{"add.", 3},
		{"add:", 3},
		{"add,", 3},
		{"add\"", 3},
		{"add(", 3},
		{"add)", 3},
		{"add[", 3},
		{"add]", 3},
		{"add// comment", 3},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Add()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestAddInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_a"},
		{"a_"},
		{"_ad"},
		{"ad_"},
		{"_add"},
		{"add_"},
		{"9713a"},
		{"a9713"},
		{"9713ad"},
		{"ad9713"},
		{"9713add"},
		{"add9713"},
		{"Za"},
		{"aZ"},
		{"Zad"},
		{"adZ"},
		{"Zadd"},
		{"addZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Add()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestSubValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"sub", 3},
		{"sub ", 3},
		{"sub\n", 3},
		{"sub.", 3},
		{"sub:", 3},
		{"sub,", 3},
		{"sub\"", 3},
		{"sub(", 3},
		{"sub)", 3},
		{"sub[", 3},
		{"sub]", 3},
		{"sub// comment", 3},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Sub()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestSubInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_s"},
		{"s_"},
		{"_su"},
		{"su_"},
		{"_sub"},
		{"sub_"},
		{"9713s"},
		{"s9713"},
		{"9713su"},
		{"su9713"},
		{"9713sub"},
		{"sub9713"},
		{"Zs"},
		{"sZ"},
		{"Zsu"},
		{"suZ"},
		{"Zsub"},
		{"subZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Sub()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestMulValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"mul", 3},
		{"mul ", 3},
		{"mul\n", 3},
		{"mul.", 3},
		{"mul:", 3},
		{"mul,", 3},
		{"mul\"", 3},
		{"mul(", 3},
		{"mul)", 3},
		{"mul[", 3},
		{"mul]", 3},
		{"mul// comment", 3},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Mul()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestMulInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_m"},
		{"m_"},
		{"_mu"},
		{"mu_"},
		{"_mul"},
		{"mul_"},
		{"9713m"},
		{"m9713"},
		{"9713mu"},
		{"mu9713"},
		{"9713mul"},
		{"mul9713"},
		{"Zm"},
		{"mZ"},
		{"Zmu"},
		{"muZ"},
		{"Zmul"},
		{"mulZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Mul()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestDivValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"div", 3},
		{"div ", 3},
		{"div\n", 3},
		{"div.", 3},
		{"div:", 3},
		{"div,", 3},
		{"div\"", 3},
		{"div(", 3},
		{"div)", 3},
		{"div[", 3},
		{"div]", 3},
		{"div// comment", 3},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Div()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestDivInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_d"},
		{"d_"},
		{"_di"},
		{"di_"},
		{"_div"},
		{"div_"},
		{"9713d"},
		{"d9713"},
		{"9713di"},
		{"di9713"},
		{"9713div"},
		{"div9713"},
		{"Zd"},
		{"dZ"},
		{"Zdi"},
		{"diZ"},
		{"Zdiv"},
		{"divZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Div()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestRemValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"rem", 3},
		{"rem ", 3},
		{"rem\n", 3},
		{"rem.", 3},
		{"rem:", 3},
		{"rem,", 3},
		{"rem\"", 3},
		{"rem(", 3},
		{"rem)", 3},
		{"rem[", 3},
		{"rem]", 3},
		{"rem// comment", 3},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Rem()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestRemInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_r"},
		{"r_"},
		{"_re"},
		{"re_"},
		{"_rem"},
		{"rem_"},
		{"9713r"},
		{"r9713"},
		{"9713re"},
		{"re9713"},
		{"9713rem"},
		{"rem9713"},
		{"Zr"},
		{"rZ"},
		{"Zre"},
		{"reZ"},
		{"Zrem"},
		{"remZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Rem()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestPowValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"pow", 3},
		{"pow ", 3},
		{"pow\n", 3},
		{"pow.", 3},
		{"pow:", 3},
		{"pow,", 3},
		{"pow\"", 3},
		{"pow(", 3},
		{"pow)", 3},
		{"pow[", 3},
		{"pow]", 3},
		{"pow// comment", 3},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Pow()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestPowInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_p"},
		{"p_"},
		{"_po"},
		{"po_"},
		{"_pow"},
		{"pow_"},
		{"9713p"},
		{"p9713"},
		{"9713po"},
		{"po9713"},
		{"9713pow"},
		{"pow9713"},
		{"Zp"},
		{"pZ"},
		{"Zpo"},
		{"poZ"},
		{"Zpow"},
		{"powZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Pow()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestSqrValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"sqr", 3},
		{"sqr ", 3},
		{"sqr\n", 3},
		{"sqr.", 3},
		{"sqr:", 3},
		{"sqr,", 3},
		{"sqr\"", 3},
		{"sqr(", 3},
		{"sqr)", 3},
		{"sqr[", 3},
		{"sqr]", 3},
		{"sqr// comment", 3},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Sqr()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestSqrInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_s"},
		{"s_"},
		{"_sq"},
		{"sq_"},
		{"_sqr"},
		{"sqr_"},
		{"9713s"},
		{"s9713"},
		{"9713sq"},
		{"sq9713"},
		{"9713sqr"},
		{"sqr9713"},
		{"Zs"},
		{"sZ"},
		{"Zsq"},
		{"sqZ"},
		{"Zsqr"},
		{"sqrZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Sqr()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestSqrtValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"sqrt", 4},
		{"sqrt ", 4},
		{"sqrt\n", 4},
		{"sqrt.", 4},
		{"sqrt:", 4},
		{"sqrt,", 4},
		{"sqrt\"", 4},
		{"sqrt(", 4},
		{"sqrt)", 4},
		{"sqrt[", 4},
		{"sqrt]", 4},
		{"sqrt// comment", 4},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Sqrt()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestSqrtInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_s"},
		{"s_"},
		{"_sq"},
		{"sq_"},
		{"_sqr"},
		{"sqr_"},
		{"_sqrt"},
		{"sqrt_"},
		{"9713s"},
		{"s9713"},
		{"9713sq"},
		{"sq9713"},
		{"9713sqr"},
		{"sqr9713"},
		{"9713sqrt"},
		{"sqrt9713"},
		{"Zs"},
		{"sZ"},
		{"Zsq"},
		{"sqZ"},
		{"Zsqr"},
		{"sqrZ"},
		{"Zsqrt"},
		{"sqrtZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Sqrt()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestMidValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"mid", 3},
		{"mid ", 3},
		{"mid\n", 3},
		{"mid.", 3},
		{"mid:", 3},
		{"mid,", 3},
		{"mid\"", 3},
		{"mid(", 3},
		{"mid)", 3},
		{"mid[", 3},
		{"mid]", 3},
		{"mid// comment", 3},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Mid()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestMidInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_m"},
		{"m_"},
		{"_mi"},
		{"mi_"},
		{"_mid"},
		{"mid_"},
		{"9713m"},
		{"m9713"},
		{"9713mi"},
		{"mi9713"},
		{"9713mid"},
		{"mid9713"},
		{"Zm"},
		{"mZ"},
		{"Zmi"},
		{"miZ"},
		{"Zmid"},
		{"midZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Mid()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestAvgValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"avg", 3},
		{"avg ", 3},
		{"avg\n", 3},
		{"avg.", 3},
		{"avg:", 3},
		{"avg,", 3},
		{"avg\"", 3},
		{"avg(", 3},
		{"avg)", 3},
		{"avg[", 3},
		{"avg]", 3},
		{"avg// comment", 3},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Avg()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestAvgInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_a"},
		{"a_"},
		{"_av"},
		{"av_"},
		{"_avg"},
		{"avg_"},
		{"9713a"},
		{"a9713"},
		{"9713av"},
		{"av9713"},
		{"9713avg"},
		{"avg9713"},
		{"Za"},
		{"aZ"},
		{"Zav"},
		{"avZ"},
		{"Zavg"},
		{"avgZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Avg()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestAvgGeoValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"avgGeo", 6},
		{"avgGeo ", 6},
		{"avgGeo\n", 6},
		{"avgGeo.", 6},
		{"avgGeo:", 6},
		{"avgGeo,", 6},
		{"avgGeo\"", 6},
		{"avgGeo(", 6},
		{"avgGeo)", 6},
		{"avgGeo[", 6},
		{"avgGeo]", 6},
		{"avgGeo// comment", 6},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.AvgGeo()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestAvgGeoInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_a"},
		{"a_"},
		{"_av"},
		{"av_"},
		{"_avg"},
		{"avg_"},
		{"_avgG"},
		{"avgG_"},
		{"_avgGe"},
		{"avgGe_"},
		{"_avgGeo"},
		{"avgGeo_"},
		{"9713a"},
		{"a9713"},
		{"9713av"},
		{"av9713"},
		{"9713avg"},
		{"avg9713"},
		{"9713avgG"},
		{"avgG9713"},
		{"9713avgGe"},
		{"avgGe9713"},
		{"9713avgGeo"},
		{"avgGeo9713"},
		{"Za"},
		{"aZ"},
		{"Zav"},
		{"avZ"},
		{"Zavg"},
		{"avgZ"},
		{"ZavgG"},
		{"avgGZ"},
		{"ZavgGe"},
		{"avgGeZ"},
		{"ZavgGeo"},
		{"avgGeoZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.AvgGeo()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestSelEqlValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"selEql", 6},
		{"selEql ", 6},
		{"selEql\n", 6},
		{"selEql.", 6},
		{"selEql:", 6},
		{"selEql,", 6},
		{"selEql\"", 6},
		{"selEql(", 6},
		{"selEql)", 6},
		{"selEql[", 6},
		{"selEql]", 6},
		{"selEql// comment", 6},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.SelEql()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestSelEqlInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_s"},
		{"s_"},
		{"_se"},
		{"se_"},
		{"_sel"},
		{"sel_"},
		{"_selE"},
		{"selE_"},
		{"_selEq"},
		{"selEq_"},
		{"_selEql"},
		{"selEql_"},
		{"9713s"},
		{"s9713"},
		{"9713se"},
		{"se9713"},
		{"9713sel"},
		{"sel9713"},
		{"9713selE"},
		{"selE9713"},
		{"9713selEq"},
		{"selEq9713"},
		{"9713selEql"},
		{"selEql9713"},
		{"Zs"},
		{"sZ"},
		{"Zse"},
		{"seZ"},
		{"Zsel"},
		{"selZ"},
		{"ZselE"},
		{"selEZ"},
		{"ZselEq"},
		{"selEqZ"},
		{"ZselEql"},
		{"selEqlZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.SelEql()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestSelNeqValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"selNeq", 6},
		{"selNeq ", 6},
		{"selNeq\n", 6},
		{"selNeq.", 6},
		{"selNeq:", 6},
		{"selNeq,", 6},
		{"selNeq\"", 6},
		{"selNeq(", 6},
		{"selNeq)", 6},
		{"selNeq[", 6},
		{"selNeq]", 6},
		{"selNeq// comment", 6},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.SelNeq()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestSelNeqInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_s"},
		{"s_"},
		{"_se"},
		{"se_"},
		{"_sel"},
		{"sel_"},
		{"_selN"},
		{"selN_"},
		{"_selNe"},
		{"selNe_"},
		{"_selNeq"},
		{"selNeq_"},
		{"9713s"},
		{"s9713"},
		{"9713se"},
		{"se9713"},
		{"9713sel"},
		{"sel9713"},
		{"9713selN"},
		{"selN9713"},
		{"9713selNe"},
		{"selNe9713"},
		{"9713selNeq"},
		{"selNeq9713"},
		{"Zs"},
		{"sZ"},
		{"Zse"},
		{"seZ"},
		{"Zsel"},
		{"selZ"},
		{"ZselN"},
		{"selNZ"},
		{"ZselNe"},
		{"selNeZ"},
		{"ZselNeq"},
		{"selNeqZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.SelNeq()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestSelLssValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"selLss", 6},
		{"selLss ", 6},
		{"selLss\n", 6},
		{"selLss.", 6},
		{"selLss:", 6},
		{"selLss,", 6},
		{"selLss\"", 6},
		{"selLss(", 6},
		{"selLss)", 6},
		{"selLss[", 6},
		{"selLss]", 6},
		{"selLss// comment", 6},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.SelLss()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestSelLssInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_s"},
		{"s_"},
		{"_se"},
		{"se_"},
		{"_sel"},
		{"sel_"},
		{"_selL"},
		{"selL_"},
		{"_selLs"},
		{"selLs_"},
		{"_selLss"},
		{"selLss_"},
		{"9713s"},
		{"s9713"},
		{"9713se"},
		{"se9713"},
		{"9713sel"},
		{"sel9713"},
		{"9713selL"},
		{"selL9713"},
		{"9713selLs"},
		{"selLs9713"},
		{"9713selLss"},
		{"selLss9713"},
		{"Zs"},
		{"sZ"},
		{"Zse"},
		{"seZ"},
		{"Zsel"},
		{"selZ"},
		{"ZselL"},
		{"selLZ"},
		{"ZselLs"},
		{"selLsZ"},
		{"ZselLss"},
		{"selLssZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.SelLss()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestSelGtrValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"selGtr", 6},
		{"selGtr ", 6},
		{"selGtr\n", 6},
		{"selGtr.", 6},
		{"selGtr:", 6},
		{"selGtr,", 6},
		{"selGtr\"", 6},
		{"selGtr(", 6},
		{"selGtr)", 6},
		{"selGtr[", 6},
		{"selGtr]", 6},
		{"selGtr// comment", 6},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.SelGtr()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestSelGtrInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_s"},
		{"s_"},
		{"_se"},
		{"se_"},
		{"_sel"},
		{"sel_"},
		{"_selG"},
		{"selG_"},
		{"_selGt"},
		{"selGt_"},
		{"_selGtr"},
		{"selGtr_"},
		{"9713s"},
		{"s9713"},
		{"9713se"},
		{"se9713"},
		{"9713sel"},
		{"sel9713"},
		{"9713selG"},
		{"selG9713"},
		{"9713selGt"},
		{"selGt9713"},
		{"9713selGtr"},
		{"selGtr9713"},
		{"Zs"},
		{"sZ"},
		{"Zse"},
		{"seZ"},
		{"Zsel"},
		{"selZ"},
		{"ZselG"},
		{"selGZ"},
		{"ZselGt"},
		{"selGtZ"},
		{"ZselGtr"},
		{"selGtrZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.SelGtr()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestSelLeqValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"selLeq", 6},
		{"selLeq ", 6},
		{"selLeq\n", 6},
		{"selLeq.", 6},
		{"selLeq:", 6},
		{"selLeq,", 6},
		{"selLeq\"", 6},
		{"selLeq(", 6},
		{"selLeq)", 6},
		{"selLeq[", 6},
		{"selLeq]", 6},
		{"selLeq// comment", 6},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.SelLeq()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestSelLeqInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_s"},
		{"s_"},
		{"_se"},
		{"se_"},
		{"_sel"},
		{"sel_"},
		{"_selL"},
		{"selL_"},
		{"_selLe"},
		{"selLe_"},
		{"_selLeq"},
		{"selLeq_"},
		{"9713s"},
		{"s9713"},
		{"9713se"},
		{"se9713"},
		{"9713sel"},
		{"sel9713"},
		{"9713selL"},
		{"selL9713"},
		{"9713selLe"},
		{"selLe9713"},
		{"9713selLeq"},
		{"selLeq9713"},
		{"Zs"},
		{"sZ"},
		{"Zse"},
		{"seZ"},
		{"Zsel"},
		{"selZ"},
		{"ZselL"},
		{"selLZ"},
		{"ZselLe"},
		{"selLeZ"},
		{"ZselLeq"},
		{"selLeqZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.SelLeq()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestSelGeqValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"selGeq", 6},
		{"selGeq ", 6},
		{"selGeq\n", 6},
		{"selGeq.", 6},
		{"selGeq:", 6},
		{"selGeq,", 6},
		{"selGeq\"", 6},
		{"selGeq(", 6},
		{"selGeq)", 6},
		{"selGeq[", 6},
		{"selGeq]", 6},
		{"selGeq// comment", 6},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.SelGeq()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestSelGeqInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_s"},
		{"s_"},
		{"_se"},
		{"se_"},
		{"_sel"},
		{"sel_"},
		{"_selG"},
		{"selG_"},
		{"_selGe"},
		{"selGe_"},
		{"_selGeq"},
		{"selGeq_"},
		{"9713s"},
		{"s9713"},
		{"9713se"},
		{"se9713"},
		{"9713sel"},
		{"sel9713"},
		{"9713selG"},
		{"selG9713"},
		{"9713selGe"},
		{"selGe9713"},
		{"9713selGeq"},
		{"selGeq9713"},
		{"Zs"},
		{"sZ"},
		{"Zse"},
		{"seZ"},
		{"Zsel"},
		{"selZ"},
		{"ZselG"},
		{"selGZ"},
		{"ZselGe"},
		{"selGeZ"},
		{"ZselGeq"},
		{"selGeqZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.SelGeq()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestWeekdayCntValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"weekdayCnt", 10},
		{"weekdayCnt ", 10},
		{"weekdayCnt\n", 10},
		{"weekdayCnt.", 10},
		{"weekdayCnt:", 10},
		{"weekdayCnt,", 10},
		{"weekdayCnt\"", 10},
		{"weekdayCnt(", 10},
		{"weekdayCnt)", 10},
		{"weekdayCnt[", 10},
		{"weekdayCnt]", 10},
		{"weekdayCnt// comment", 10},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.WeekdayCnt()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestWeekdayCntInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_w"},
		{"w_"},
		{"_we"},
		{"we_"},
		{"_wee"},
		{"wee_"},
		{"_week"},
		{"week_"},
		{"_weekd"},
		{"weekd_"},
		{"_weekda"},
		{"weekda_"},
		{"_weekday"},
		{"weekday_"},
		{"_weekdayC"},
		{"weekdayC_"},
		{"_weekdayCn"},
		{"weekdayCn_"},
		{"_weekdayCnt"},
		{"weekdayCnt_"},
		{"9713w"},
		{"w9713"},
		{"9713we"},
		{"we9713"},
		{"9713wee"},
		{"wee9713"},
		{"9713week"},
		{"week9713"},
		{"9713weekd"},
		{"weekd9713"},
		{"9713weekda"},
		{"weekda9713"},
		{"9713weekday"},
		{"weekday9713"},
		{"9713weekdayC"},
		{"weekdayC9713"},
		{"9713weekdayCn"},
		{"weekdayCn9713"},
		{"9713weekdayCnt"},
		{"weekdayCnt9713"},
		{"Zw"},
		{"wZ"},
		{"Zwe"},
		{"weZ"},
		{"Zwee"},
		{"weeZ"},
		{"Zweek"},
		{"weekZ"},
		{"Zweekd"},
		{"weekdZ"},
		{"Zweekda"},
		{"weekdaZ"},
		{"Zweekday"},
		{"weekdayZ"},
		{"ZweekdayC"},
		{"weekdayCZ"},
		{"ZweekdayCn"},
		{"weekdayCnZ"},
		{"ZweekdayCnt"},
		{"weekdayCntZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.WeekdayCnt()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestDteValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"dte", 3},
		{"dte ", 3},
		{"dte\n", 3},
		{"dte.", 3},
		{"dte:", 3},
		{"dte,", 3},
		{"dte\"", 3},
		{"dte(", 3},
		{"dte)", 3},
		{"dte[", 3},
		{"dte]", 3},
		{"dte// comment", 3},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Dte()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestDteInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_d"},
		{"d_"},
		{"_dt"},
		{"dt_"},
		{"_dte"},
		{"dte_"},
		{"9713d"},
		{"d9713"},
		{"9713dt"},
		{"dt9713"},
		{"9713dte"},
		{"dte9713"},
		{"Zd"},
		{"dZ"},
		{"Zdt"},
		{"dtZ"},
		{"Zdte"},
		{"dteZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Dte()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestToSundayValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"toSunday", 8},
		{"toSunday ", 8},
		{"toSunday\n", 8},
		{"toSunday.", 8},
		{"toSunday:", 8},
		{"toSunday,", 8},
		{"toSunday\"", 8},
		{"toSunday(", 8},
		{"toSunday)", 8},
		{"toSunday[", 8},
		{"toSunday]", 8},
		{"toSunday// comment", 8},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.ToSunday()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestToSundayInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_t"},
		{"t_"},
		{"_to"},
		{"to_"},
		{"_toS"},
		{"toS_"},
		{"_toSu"},
		{"toSu_"},
		{"_toSun"},
		{"toSun_"},
		{"_toSund"},
		{"toSund_"},
		{"_toSunda"},
		{"toSunda_"},
		{"_toSunday"},
		{"toSunday_"},
		{"9713t"},
		{"t9713"},
		{"9713to"},
		{"to9713"},
		{"9713toS"},
		{"toS9713"},
		{"9713toSu"},
		{"toSu9713"},
		{"9713toSun"},
		{"toSun9713"},
		{"9713toSund"},
		{"toSund9713"},
		{"9713toSunda"},
		{"toSunda9713"},
		{"9713toSunday"},
		{"toSunday9713"},
		{"Zt"},
		{"tZ"},
		{"Zto"},
		{"toZ"},
		{"ZtoS"},
		{"toSZ"},
		{"ZtoSu"},
		{"toSuZ"},
		{"ZtoSun"},
		{"toSunZ"},
		{"ZtoSund"},
		{"toSundZ"},
		{"ZtoSunda"},
		{"toSundaZ"},
		{"ZtoSunday"},
		{"toSundayZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.ToSunday()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestToMondayValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"toMonday", 8},
		{"toMonday ", 8},
		{"toMonday\n", 8},
		{"toMonday.", 8},
		{"toMonday:", 8},
		{"toMonday,", 8},
		{"toMonday\"", 8},
		{"toMonday(", 8},
		{"toMonday)", 8},
		{"toMonday[", 8},
		{"toMonday]", 8},
		{"toMonday// comment", 8},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.ToMonday()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestToMondayInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_t"},
		{"t_"},
		{"_to"},
		{"to_"},
		{"_toM"},
		{"toM_"},
		{"_toMo"},
		{"toMo_"},
		{"_toMon"},
		{"toMon_"},
		{"_toMond"},
		{"toMond_"},
		{"_toMonda"},
		{"toMonda_"},
		{"_toMonday"},
		{"toMonday_"},
		{"9713t"},
		{"t9713"},
		{"9713to"},
		{"to9713"},
		{"9713toM"},
		{"toM9713"},
		{"9713toMo"},
		{"toMo9713"},
		{"9713toMon"},
		{"toMon9713"},
		{"9713toMond"},
		{"toMond9713"},
		{"9713toMonda"},
		{"toMonda9713"},
		{"9713toMonday"},
		{"toMonday9713"},
		{"Zt"},
		{"tZ"},
		{"Zto"},
		{"toZ"},
		{"ZtoM"},
		{"toMZ"},
		{"ZtoMo"},
		{"toMoZ"},
		{"ZtoMon"},
		{"toMonZ"},
		{"ZtoMond"},
		{"toMondZ"},
		{"ZtoMonda"},
		{"toMondaZ"},
		{"ZtoMonday"},
		{"toMondayZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.ToMonday()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestToTuesdayValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"toTuesday", 9},
		{"toTuesday ", 9},
		{"toTuesday\n", 9},
		{"toTuesday.", 9},
		{"toTuesday:", 9},
		{"toTuesday,", 9},
		{"toTuesday\"", 9},
		{"toTuesday(", 9},
		{"toTuesday)", 9},
		{"toTuesday[", 9},
		{"toTuesday]", 9},
		{"toTuesday// comment", 9},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.ToTuesday()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestToTuesdayInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_t"},
		{"t_"},
		{"_to"},
		{"to_"},
		{"_toT"},
		{"toT_"},
		{"_toTu"},
		{"toTu_"},
		{"_toTue"},
		{"toTue_"},
		{"_toTues"},
		{"toTues_"},
		{"_toTuesd"},
		{"toTuesd_"},
		{"_toTuesda"},
		{"toTuesda_"},
		{"_toTuesday"},
		{"toTuesday_"},
		{"9713t"},
		{"t9713"},
		{"9713to"},
		{"to9713"},
		{"9713toT"},
		{"toT9713"},
		{"9713toTu"},
		{"toTu9713"},
		{"9713toTue"},
		{"toTue9713"},
		{"9713toTues"},
		{"toTues9713"},
		{"9713toTuesd"},
		{"toTuesd9713"},
		{"9713toTuesda"},
		{"toTuesda9713"},
		{"9713toTuesday"},
		{"toTuesday9713"},
		{"Zt"},
		{"tZ"},
		{"Zto"},
		{"toZ"},
		{"ZtoT"},
		{"toTZ"},
		{"ZtoTu"},
		{"toTuZ"},
		{"ZtoTue"},
		{"toTueZ"},
		{"ZtoTues"},
		{"toTuesZ"},
		{"ZtoTuesd"},
		{"toTuesdZ"},
		{"ZtoTuesda"},
		{"toTuesdaZ"},
		{"ZtoTuesday"},
		{"toTuesdayZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.ToTuesday()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestToWednesdayValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"toWednesday", 11},
		{"toWednesday ", 11},
		{"toWednesday\n", 11},
		{"toWednesday.", 11},
		{"toWednesday:", 11},
		{"toWednesday,", 11},
		{"toWednesday\"", 11},
		{"toWednesday(", 11},
		{"toWednesday)", 11},
		{"toWednesday[", 11},
		{"toWednesday]", 11},
		{"toWednesday// comment", 11},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.ToWednesday()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestToWednesdayInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_t"},
		{"t_"},
		{"_to"},
		{"to_"},
		{"_toW"},
		{"toW_"},
		{"_toWe"},
		{"toWe_"},
		{"_toWed"},
		{"toWed_"},
		{"_toWedn"},
		{"toWedn_"},
		{"_toWedne"},
		{"toWedne_"},
		{"_toWednes"},
		{"toWednes_"},
		{"_toWednesd"},
		{"toWednesd_"},
		{"_toWednesda"},
		{"toWednesda_"},
		{"_toWednesday"},
		{"toWednesday_"},
		{"9713t"},
		{"t9713"},
		{"9713to"},
		{"to9713"},
		{"9713toW"},
		{"toW9713"},
		{"9713toWe"},
		{"toWe9713"},
		{"9713toWed"},
		{"toWed9713"},
		{"9713toWedn"},
		{"toWedn9713"},
		{"9713toWedne"},
		{"toWedne9713"},
		{"9713toWednes"},
		{"toWednes9713"},
		{"9713toWednesd"},
		{"toWednesd9713"},
		{"9713toWednesda"},
		{"toWednesda9713"},
		{"9713toWednesday"},
		{"toWednesday9713"},
		{"Zt"},
		{"tZ"},
		{"Zto"},
		{"toZ"},
		{"ZtoW"},
		{"toWZ"},
		{"ZtoWe"},
		{"toWeZ"},
		{"ZtoWed"},
		{"toWedZ"},
		{"ZtoWedn"},
		{"toWednZ"},
		{"ZtoWedne"},
		{"toWedneZ"},
		{"ZtoWednes"},
		{"toWednesZ"},
		{"ZtoWednesd"},
		{"toWednesdZ"},
		{"ZtoWednesda"},
		{"toWednesdaZ"},
		{"ZtoWednesday"},
		{"toWednesdayZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.ToWednesday()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestToThursdayValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"toThursday", 10},
		{"toThursday ", 10},
		{"toThursday\n", 10},
		{"toThursday.", 10},
		{"toThursday:", 10},
		{"toThursday,", 10},
		{"toThursday\"", 10},
		{"toThursday(", 10},
		{"toThursday)", 10},
		{"toThursday[", 10},
		{"toThursday]", 10},
		{"toThursday// comment", 10},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.ToThursday()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestToThursdayInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_t"},
		{"t_"},
		{"_to"},
		{"to_"},
		{"_toT"},
		{"toT_"},
		{"_toTh"},
		{"toTh_"},
		{"_toThu"},
		{"toThu_"},
		{"_toThur"},
		{"toThur_"},
		{"_toThurs"},
		{"toThurs_"},
		{"_toThursd"},
		{"toThursd_"},
		{"_toThursda"},
		{"toThursda_"},
		{"_toThursday"},
		{"toThursday_"},
		{"9713t"},
		{"t9713"},
		{"9713to"},
		{"to9713"},
		{"9713toT"},
		{"toT9713"},
		{"9713toTh"},
		{"toTh9713"},
		{"9713toThu"},
		{"toThu9713"},
		{"9713toThur"},
		{"toThur9713"},
		{"9713toThurs"},
		{"toThurs9713"},
		{"9713toThursd"},
		{"toThursd9713"},
		{"9713toThursda"},
		{"toThursda9713"},
		{"9713toThursday"},
		{"toThursday9713"},
		{"Zt"},
		{"tZ"},
		{"Zto"},
		{"toZ"},
		{"ZtoT"},
		{"toTZ"},
		{"ZtoTh"},
		{"toThZ"},
		{"ZtoThu"},
		{"toThuZ"},
		{"ZtoThur"},
		{"toThurZ"},
		{"ZtoThurs"},
		{"toThursZ"},
		{"ZtoThursd"},
		{"toThursdZ"},
		{"ZtoThursda"},
		{"toThursdaZ"},
		{"ZtoThursday"},
		{"toThursdayZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.ToThursday()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestToFridayValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"toFriday", 8},
		{"toFriday ", 8},
		{"toFriday\n", 8},
		{"toFriday.", 8},
		{"toFriday:", 8},
		{"toFriday,", 8},
		{"toFriday\"", 8},
		{"toFriday(", 8},
		{"toFriday)", 8},
		{"toFriday[", 8},
		{"toFriday]", 8},
		{"toFriday// comment", 8},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.ToFriday()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestToFridayInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_t"},
		{"t_"},
		{"_to"},
		{"to_"},
		{"_toF"},
		{"toF_"},
		{"_toFr"},
		{"toFr_"},
		{"_toFri"},
		{"toFri_"},
		{"_toFrid"},
		{"toFrid_"},
		{"_toFrida"},
		{"toFrida_"},
		{"_toFriday"},
		{"toFriday_"},
		{"9713t"},
		{"t9713"},
		{"9713to"},
		{"to9713"},
		{"9713toF"},
		{"toF9713"},
		{"9713toFr"},
		{"toFr9713"},
		{"9713toFri"},
		{"toFri9713"},
		{"9713toFrid"},
		{"toFrid9713"},
		{"9713toFrida"},
		{"toFrida9713"},
		{"9713toFriday"},
		{"toFriday9713"},
		{"Zt"},
		{"tZ"},
		{"Zto"},
		{"toZ"},
		{"ZtoF"},
		{"toFZ"},
		{"ZtoFr"},
		{"toFrZ"},
		{"ZtoFri"},
		{"toFriZ"},
		{"ZtoFrid"},
		{"toFridZ"},
		{"ZtoFrida"},
		{"toFridaZ"},
		{"ZtoFriday"},
		{"toFridayZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.ToFriday()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestToSaturdayValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"toSaturday", 10},
		{"toSaturday ", 10},
		{"toSaturday\n", 10},
		{"toSaturday.", 10},
		{"toSaturday:", 10},
		{"toSaturday,", 10},
		{"toSaturday\"", 10},
		{"toSaturday(", 10},
		{"toSaturday)", 10},
		{"toSaturday[", 10},
		{"toSaturday]", 10},
		{"toSaturday// comment", 10},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.ToSaturday()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestToSaturdayInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_t"},
		{"t_"},
		{"_to"},
		{"to_"},
		{"_toS"},
		{"toS_"},
		{"_toSa"},
		{"toSa_"},
		{"_toSat"},
		{"toSat_"},
		{"_toSatu"},
		{"toSatu_"},
		{"_toSatur"},
		{"toSatur_"},
		{"_toSaturd"},
		{"toSaturd_"},
		{"_toSaturda"},
		{"toSaturda_"},
		{"_toSaturday"},
		{"toSaturday_"},
		{"9713t"},
		{"t9713"},
		{"9713to"},
		{"to9713"},
		{"9713toS"},
		{"toS9713"},
		{"9713toSa"},
		{"toSa9713"},
		{"9713toSat"},
		{"toSat9713"},
		{"9713toSatu"},
		{"toSatu9713"},
		{"9713toSatur"},
		{"toSatur9713"},
		{"9713toSaturd"},
		{"toSaturd9713"},
		{"9713toSaturda"},
		{"toSaturda9713"},
		{"9713toSaturday"},
		{"toSaturday9713"},
		{"Zt"},
		{"tZ"},
		{"Zto"},
		{"toZ"},
		{"ZtoS"},
		{"toSZ"},
		{"ZtoSa"},
		{"toSaZ"},
		{"ZtoSat"},
		{"toSatZ"},
		{"ZtoSatu"},
		{"toSatuZ"},
		{"ZtoSatur"},
		{"toSaturZ"},
		{"ZtoSaturd"},
		{"toSaturdZ"},
		{"ZtoSaturda"},
		{"toSaturdaZ"},
		{"ZtoSaturday"},
		{"toSaturdayZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.ToSaturday()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestIsSundayValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"isSunday", 8},
		{"isSunday ", 8},
		{"isSunday\n", 8},
		{"isSunday.", 8},
		{"isSunday:", 8},
		{"isSunday,", 8},
		{"isSunday\"", 8},
		{"isSunday(", 8},
		{"isSunday)", 8},
		{"isSunday[", 8},
		{"isSunday]", 8},
		{"isSunday// comment", 8},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.IsSunday()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestIsSundayInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_i"},
		{"i_"},
		{"_is"},
		{"is_"},
		{"_isS"},
		{"isS_"},
		{"_isSu"},
		{"isSu_"},
		{"_isSun"},
		{"isSun_"},
		{"_isSund"},
		{"isSund_"},
		{"_isSunda"},
		{"isSunda_"},
		{"_isSunday"},
		{"isSunday_"},
		{"9713i"},
		{"i9713"},
		{"9713is"},
		{"is9713"},
		{"9713isS"},
		{"isS9713"},
		{"9713isSu"},
		{"isSu9713"},
		{"9713isSun"},
		{"isSun9713"},
		{"9713isSund"},
		{"isSund9713"},
		{"9713isSunda"},
		{"isSunda9713"},
		{"9713isSunday"},
		{"isSunday9713"},
		{"Zi"},
		{"iZ"},
		{"Zis"},
		{"isZ"},
		{"ZisS"},
		{"isSZ"},
		{"ZisSu"},
		{"isSuZ"},
		{"ZisSun"},
		{"isSunZ"},
		{"ZisSund"},
		{"isSundZ"},
		{"ZisSunda"},
		{"isSundaZ"},
		{"ZisSunday"},
		{"isSundayZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.IsSunday()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestIsMondayValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"isMonday", 8},
		{"isMonday ", 8},
		{"isMonday\n", 8},
		{"isMonday.", 8},
		{"isMonday:", 8},
		{"isMonday,", 8},
		{"isMonday\"", 8},
		{"isMonday(", 8},
		{"isMonday)", 8},
		{"isMonday[", 8},
		{"isMonday]", 8},
		{"isMonday// comment", 8},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.IsMonday()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestIsMondayInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_i"},
		{"i_"},
		{"_is"},
		{"is_"},
		{"_isM"},
		{"isM_"},
		{"_isMo"},
		{"isMo_"},
		{"_isMon"},
		{"isMon_"},
		{"_isMond"},
		{"isMond_"},
		{"_isMonda"},
		{"isMonda_"},
		{"_isMonday"},
		{"isMonday_"},
		{"9713i"},
		{"i9713"},
		{"9713is"},
		{"is9713"},
		{"9713isM"},
		{"isM9713"},
		{"9713isMo"},
		{"isMo9713"},
		{"9713isMon"},
		{"isMon9713"},
		{"9713isMond"},
		{"isMond9713"},
		{"9713isMonda"},
		{"isMonda9713"},
		{"9713isMonday"},
		{"isMonday9713"},
		{"Zi"},
		{"iZ"},
		{"Zis"},
		{"isZ"},
		{"ZisM"},
		{"isMZ"},
		{"ZisMo"},
		{"isMoZ"},
		{"ZisMon"},
		{"isMonZ"},
		{"ZisMond"},
		{"isMondZ"},
		{"ZisMonda"},
		{"isMondaZ"},
		{"ZisMonday"},
		{"isMondayZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.IsMonday()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestIsTuesdayValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"isTuesday", 9},
		{"isTuesday ", 9},
		{"isTuesday\n", 9},
		{"isTuesday.", 9},
		{"isTuesday:", 9},
		{"isTuesday,", 9},
		{"isTuesday\"", 9},
		{"isTuesday(", 9},
		{"isTuesday)", 9},
		{"isTuesday[", 9},
		{"isTuesday]", 9},
		{"isTuesday// comment", 9},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.IsTuesday()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestIsTuesdayInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_i"},
		{"i_"},
		{"_is"},
		{"is_"},
		{"_isT"},
		{"isT_"},
		{"_isTu"},
		{"isTu_"},
		{"_isTue"},
		{"isTue_"},
		{"_isTues"},
		{"isTues_"},
		{"_isTuesd"},
		{"isTuesd_"},
		{"_isTuesda"},
		{"isTuesda_"},
		{"_isTuesday"},
		{"isTuesday_"},
		{"9713i"},
		{"i9713"},
		{"9713is"},
		{"is9713"},
		{"9713isT"},
		{"isT9713"},
		{"9713isTu"},
		{"isTu9713"},
		{"9713isTue"},
		{"isTue9713"},
		{"9713isTues"},
		{"isTues9713"},
		{"9713isTuesd"},
		{"isTuesd9713"},
		{"9713isTuesda"},
		{"isTuesda9713"},
		{"9713isTuesday"},
		{"isTuesday9713"},
		{"Zi"},
		{"iZ"},
		{"Zis"},
		{"isZ"},
		{"ZisT"},
		{"isTZ"},
		{"ZisTu"},
		{"isTuZ"},
		{"ZisTue"},
		{"isTueZ"},
		{"ZisTues"},
		{"isTuesZ"},
		{"ZisTuesd"},
		{"isTuesdZ"},
		{"ZisTuesda"},
		{"isTuesdaZ"},
		{"ZisTuesday"},
		{"isTuesdayZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.IsTuesday()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestIsWednesdayValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"isWednesday", 11},
		{"isWednesday ", 11},
		{"isWednesday\n", 11},
		{"isWednesday.", 11},
		{"isWednesday:", 11},
		{"isWednesday,", 11},
		{"isWednesday\"", 11},
		{"isWednesday(", 11},
		{"isWednesday)", 11},
		{"isWednesday[", 11},
		{"isWednesday]", 11},
		{"isWednesday// comment", 11},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.IsWednesday()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestIsWednesdayInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_i"},
		{"i_"},
		{"_is"},
		{"is_"},
		{"_isW"},
		{"isW_"},
		{"_isWe"},
		{"isWe_"},
		{"_isWed"},
		{"isWed_"},
		{"_isWedn"},
		{"isWedn_"},
		{"_isWedne"},
		{"isWedne_"},
		{"_isWednes"},
		{"isWednes_"},
		{"_isWednesd"},
		{"isWednesd_"},
		{"_isWednesda"},
		{"isWednesda_"},
		{"_isWednesday"},
		{"isWednesday_"},
		{"9713i"},
		{"i9713"},
		{"9713is"},
		{"is9713"},
		{"9713isW"},
		{"isW9713"},
		{"9713isWe"},
		{"isWe9713"},
		{"9713isWed"},
		{"isWed9713"},
		{"9713isWedn"},
		{"isWedn9713"},
		{"9713isWedne"},
		{"isWedne9713"},
		{"9713isWednes"},
		{"isWednes9713"},
		{"9713isWednesd"},
		{"isWednesd9713"},
		{"9713isWednesda"},
		{"isWednesda9713"},
		{"9713isWednesday"},
		{"isWednesday9713"},
		{"Zi"},
		{"iZ"},
		{"Zis"},
		{"isZ"},
		{"ZisW"},
		{"isWZ"},
		{"ZisWe"},
		{"isWeZ"},
		{"ZisWed"},
		{"isWedZ"},
		{"ZisWedn"},
		{"isWednZ"},
		{"ZisWedne"},
		{"isWedneZ"},
		{"ZisWednes"},
		{"isWednesZ"},
		{"ZisWednesd"},
		{"isWednesdZ"},
		{"ZisWednesda"},
		{"isWednesdaZ"},
		{"ZisWednesday"},
		{"isWednesdayZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.IsWednesday()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestIsThursdayValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"isThursday", 10},
		{"isThursday ", 10},
		{"isThursday\n", 10},
		{"isThursday.", 10},
		{"isThursday:", 10},
		{"isThursday,", 10},
		{"isThursday\"", 10},
		{"isThursday(", 10},
		{"isThursday)", 10},
		{"isThursday[", 10},
		{"isThursday]", 10},
		{"isThursday// comment", 10},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.IsThursday()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestIsThursdayInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_i"},
		{"i_"},
		{"_is"},
		{"is_"},
		{"_isT"},
		{"isT_"},
		{"_isTh"},
		{"isTh_"},
		{"_isThu"},
		{"isThu_"},
		{"_isThur"},
		{"isThur_"},
		{"_isThurs"},
		{"isThurs_"},
		{"_isThursd"},
		{"isThursd_"},
		{"_isThursda"},
		{"isThursda_"},
		{"_isThursday"},
		{"isThursday_"},
		{"9713i"},
		{"i9713"},
		{"9713is"},
		{"is9713"},
		{"9713isT"},
		{"isT9713"},
		{"9713isTh"},
		{"isTh9713"},
		{"9713isThu"},
		{"isThu9713"},
		{"9713isThur"},
		{"isThur9713"},
		{"9713isThurs"},
		{"isThurs9713"},
		{"9713isThursd"},
		{"isThursd9713"},
		{"9713isThursda"},
		{"isThursda9713"},
		{"9713isThursday"},
		{"isThursday9713"},
		{"Zi"},
		{"iZ"},
		{"Zis"},
		{"isZ"},
		{"ZisT"},
		{"isTZ"},
		{"ZisTh"},
		{"isThZ"},
		{"ZisThu"},
		{"isThuZ"},
		{"ZisThur"},
		{"isThurZ"},
		{"ZisThurs"},
		{"isThursZ"},
		{"ZisThursd"},
		{"isThursdZ"},
		{"ZisThursda"},
		{"isThursdaZ"},
		{"ZisThursday"},
		{"isThursdayZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.IsThursday()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestIsFridayValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"isFriday", 8},
		{"isFriday ", 8},
		{"isFriday\n", 8},
		{"isFriday.", 8},
		{"isFriday:", 8},
		{"isFriday,", 8},
		{"isFriday\"", 8},
		{"isFriday(", 8},
		{"isFriday)", 8},
		{"isFriday[", 8},
		{"isFriday]", 8},
		{"isFriday// comment", 8},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.IsFriday()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestIsFridayInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_i"},
		{"i_"},
		{"_is"},
		{"is_"},
		{"_isF"},
		{"isF_"},
		{"_isFr"},
		{"isFr_"},
		{"_isFri"},
		{"isFri_"},
		{"_isFrid"},
		{"isFrid_"},
		{"_isFrida"},
		{"isFrida_"},
		{"_isFriday"},
		{"isFriday_"},
		{"9713i"},
		{"i9713"},
		{"9713is"},
		{"is9713"},
		{"9713isF"},
		{"isF9713"},
		{"9713isFr"},
		{"isFr9713"},
		{"9713isFri"},
		{"isFri9713"},
		{"9713isFrid"},
		{"isFrid9713"},
		{"9713isFrida"},
		{"isFrida9713"},
		{"9713isFriday"},
		{"isFriday9713"},
		{"Zi"},
		{"iZ"},
		{"Zis"},
		{"isZ"},
		{"ZisF"},
		{"isFZ"},
		{"ZisFr"},
		{"isFrZ"},
		{"ZisFri"},
		{"isFriZ"},
		{"ZisFrid"},
		{"isFridZ"},
		{"ZisFrida"},
		{"isFridaZ"},
		{"ZisFriday"},
		{"isFridayZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.IsFriday()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestIsSaturdayValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"isSaturday", 10},
		{"isSaturday ", 10},
		{"isSaturday\n", 10},
		{"isSaturday.", 10},
		{"isSaturday:", 10},
		{"isSaturday,", 10},
		{"isSaturday\"", 10},
		{"isSaturday(", 10},
		{"isSaturday)", 10},
		{"isSaturday[", 10},
		{"isSaturday]", 10},
		{"isSaturday// comment", 10},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.IsSaturday()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestIsSaturdayInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_i"},
		{"i_"},
		{"_is"},
		{"is_"},
		{"_isS"},
		{"isS_"},
		{"_isSa"},
		{"isSa_"},
		{"_isSat"},
		{"isSat_"},
		{"_isSatu"},
		{"isSatu_"},
		{"_isSatur"},
		{"isSatur_"},
		{"_isSaturd"},
		{"isSaturd_"},
		{"_isSaturda"},
		{"isSaturda_"},
		{"_isSaturday"},
		{"isSaturday_"},
		{"9713i"},
		{"i9713"},
		{"9713is"},
		{"is9713"},
		{"9713isS"},
		{"isS9713"},
		{"9713isSa"},
		{"isSa9713"},
		{"9713isSat"},
		{"isSat9713"},
		{"9713isSatu"},
		{"isSatu9713"},
		{"9713isSatur"},
		{"isSatur9713"},
		{"9713isSaturd"},
		{"isSaturd9713"},
		{"9713isSaturda"},
		{"isSaturda9713"},
		{"9713isSaturday"},
		{"isSaturday9713"},
		{"Zi"},
		{"iZ"},
		{"Zis"},
		{"isZ"},
		{"ZisS"},
		{"isSZ"},
		{"ZisSa"},
		{"isSaZ"},
		{"ZisSat"},
		{"isSatZ"},
		{"ZisSatu"},
		{"isSatuZ"},
		{"ZisSatur"},
		{"isSaturZ"},
		{"ZisSaturd"},
		{"isSaturdZ"},
		{"ZisSaturda"},
		{"isSaturdaZ"},
		{"ZisSaturday"},
		{"isSaturdayZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.IsSaturday()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestCntValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"cnt", 3},
		{"cnt ", 3},
		{"cnt\n", 3},
		{"cnt.", 3},
		{"cnt:", 3},
		{"cnt,", 3},
		{"cnt\"", 3},
		{"cnt(", 3},
		{"cnt)", 3},
		{"cnt[", 3},
		{"cnt]", 3},
		{"cnt// comment", 3},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Cnt()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestCntInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_c"},
		{"c_"},
		{"_cn"},
		{"cn_"},
		{"_cnt"},
		{"cnt_"},
		{"9713c"},
		{"c9713"},
		{"9713cn"},
		{"cn9713"},
		{"9713cnt"},
		{"cnt9713"},
		{"Zc"},
		{"cZ"},
		{"Zcn"},
		{"cnZ"},
		{"Zcnt"},
		{"cntZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Cnt()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestLstIdxValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"lstIdx", 6},
		{"lstIdx ", 6},
		{"lstIdx\n", 6},
		{"lstIdx.", 6},
		{"lstIdx:", 6},
		{"lstIdx,", 6},
		{"lstIdx\"", 6},
		{"lstIdx(", 6},
		{"lstIdx)", 6},
		{"lstIdx[", 6},
		{"lstIdx]", 6},
		{"lstIdx// comment", 6},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.LstIdx()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestLstIdxInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_l"},
		{"l_"},
		{"_ls"},
		{"ls_"},
		{"_lst"},
		{"lst_"},
		{"_lstI"},
		{"lstI_"},
		{"_lstId"},
		{"lstId_"},
		{"_lstIdx"},
		{"lstIdx_"},
		{"9713l"},
		{"l9713"},
		{"9713ls"},
		{"ls9713"},
		{"9713lst"},
		{"lst9713"},
		{"9713lstI"},
		{"lstI9713"},
		{"9713lstId"},
		{"lstId9713"},
		{"9713lstIdx"},
		{"lstIdx9713"},
		{"Zl"},
		{"lZ"},
		{"Zls"},
		{"lsZ"},
		{"Zlst"},
		{"lstZ"},
		{"ZlstI"},
		{"lstIZ"},
		{"ZlstId"},
		{"lstIdZ"},
		{"ZlstIdx"},
		{"lstIdxZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.LstIdx()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestEnsureValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"ensure", 6},
		{"ensure ", 6},
		{"ensure\n", 6},
		{"ensure.", 6},
		{"ensure:", 6},
		{"ensure,", 6},
		{"ensure\"", 6},
		{"ensure(", 6},
		{"ensure)", 6},
		{"ensure[", 6},
		{"ensure]", 6},
		{"ensure// comment", 6},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Ensure()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestEnsureInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_e"},
		{"e_"},
		{"_en"},
		{"en_"},
		{"_ens"},
		{"ens_"},
		{"_ensu"},
		{"ensu_"},
		{"_ensur"},
		{"ensur_"},
		{"_ensure"},
		{"ensure_"},
		{"9713e"},
		{"e9713"},
		{"9713en"},
		{"en9713"},
		{"9713ens"},
		{"ens9713"},
		{"9713ensu"},
		{"ensu9713"},
		{"9713ensur"},
		{"ensur9713"},
		{"9713ensure"},
		{"ensure9713"},
		{"Ze"},
		{"eZ"},
		{"Zen"},
		{"enZ"},
		{"Zens"},
		{"ensZ"},
		{"Zensu"},
		{"ensuZ"},
		{"Zensur"},
		{"ensurZ"},
		{"Zensure"},
		{"ensureZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Ensure()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestMinSubValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"minSub", 6},
		{"minSub ", 6},
		{"minSub\n", 6},
		{"minSub.", 6},
		{"minSub:", 6},
		{"minSub,", 6},
		{"minSub\"", 6},
		{"minSub(", 6},
		{"minSub)", 6},
		{"minSub[", 6},
		{"minSub]", 6},
		{"minSub// comment", 6},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.MinSub()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestMinSubInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_m"},
		{"m_"},
		{"_mi"},
		{"mi_"},
		{"_min"},
		{"min_"},
		{"_minS"},
		{"minS_"},
		{"_minSu"},
		{"minSu_"},
		{"_minSub"},
		{"minSub_"},
		{"9713m"},
		{"m9713"},
		{"9713mi"},
		{"mi9713"},
		{"9713min"},
		{"min9713"},
		{"9713minS"},
		{"minS9713"},
		{"9713minSu"},
		{"minSu9713"},
		{"9713minSub"},
		{"minSub9713"},
		{"Zm"},
		{"mZ"},
		{"Zmi"},
		{"miZ"},
		{"Zmin"},
		{"minZ"},
		{"ZminS"},
		{"minSZ"},
		{"ZminSu"},
		{"minSuZ"},
		{"ZminSub"},
		{"minSubZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.MinSub()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestMaxAddValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"maxAdd", 6},
		{"maxAdd ", 6},
		{"maxAdd\n", 6},
		{"maxAdd.", 6},
		{"maxAdd:", 6},
		{"maxAdd,", 6},
		{"maxAdd\"", 6},
		{"maxAdd(", 6},
		{"maxAdd)", 6},
		{"maxAdd[", 6},
		{"maxAdd]", 6},
		{"maxAdd// comment", 6},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.MaxAdd()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestMaxAddInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_m"},
		{"m_"},
		{"_ma"},
		{"ma_"},
		{"_max"},
		{"max_"},
		{"_maxA"},
		{"maxA_"},
		{"_maxAd"},
		{"maxAd_"},
		{"_maxAdd"},
		{"maxAdd_"},
		{"9713m"},
		{"m9713"},
		{"9713ma"},
		{"ma9713"},
		{"9713max"},
		{"max9713"},
		{"9713maxA"},
		{"maxA9713"},
		{"9713maxAd"},
		{"maxAd9713"},
		{"9713maxAdd"},
		{"maxAdd9713"},
		{"Zm"},
		{"mZ"},
		{"Zma"},
		{"maZ"},
		{"Zmax"},
		{"maxZ"},
		{"ZmaxA"},
		{"maxAZ"},
		{"ZmaxAd"},
		{"maxAdZ"},
		{"ZmaxAdd"},
		{"maxAddZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.MaxAdd()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestMrgValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"mrg", 3},
		{"mrg ", 3},
		{"mrg\n", 3},
		{"mrg.", 3},
		{"mrg:", 3},
		{"mrg,", 3},
		{"mrg\"", 3},
		{"mrg(", 3},
		{"mrg)", 3},
		{"mrg[", 3},
		{"mrg]", 3},
		{"mrg// comment", 3},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Mrg()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestMrgInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_m"},
		{"m_"},
		{"_mr"},
		{"mr_"},
		{"_mrg"},
		{"mrg_"},
		{"9713m"},
		{"m9713"},
		{"9713mr"},
		{"mr9713"},
		{"9713mrg"},
		{"mrg9713"},
		{"Zm"},
		{"mZ"},
		{"Zmr"},
		{"mrZ"},
		{"Zmrg"},
		{"mrgZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Mrg()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestCpyValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"cpy", 3},
		{"cpy ", 3},
		{"cpy\n", 3},
		{"cpy.", 3},
		{"cpy:", 3},
		{"cpy,", 3},
		{"cpy\"", 3},
		{"cpy(", 3},
		{"cpy)", 3},
		{"cpy[", 3},
		{"cpy]", 3},
		{"cpy// comment", 3},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Cpy()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestCpyInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_c"},
		{"c_"},
		{"_cp"},
		{"cp_"},
		{"_cpy"},
		{"cpy_"},
		{"9713c"},
		{"c9713"},
		{"9713cp"},
		{"cp9713"},
		{"9713cpy"},
		{"cpy9713"},
		{"Zc"},
		{"cZ"},
		{"Zcp"},
		{"cpZ"},
		{"Zcpy"},
		{"cpyZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Cpy()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestRandValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"rand", 4},
		{"rand ", 4},
		{"rand\n", 4},
		{"rand.", 4},
		{"rand:", 4},
		{"rand,", 4},
		{"rand\"", 4},
		{"rand(", 4},
		{"rand)", 4},
		{"rand[", 4},
		{"rand]", 4},
		{"rand// comment", 4},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Rand()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestRandInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_r"},
		{"r_"},
		{"_ra"},
		{"ra_"},
		{"_ran"},
		{"ran_"},
		{"_rand"},
		{"rand_"},
		{"9713r"},
		{"r9713"},
		{"9713ra"},
		{"ra9713"},
		{"9713ran"},
		{"ran9713"},
		{"9713rand"},
		{"rand9713"},
		{"Zr"},
		{"rZ"},
		{"Zra"},
		{"raZ"},
		{"Zran"},
		{"ranZ"},
		{"Zrand"},
		{"randZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Rand()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestPushValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"push", 4},
		{"push ", 4},
		{"push\n", 4},
		{"push.", 4},
		{"push:", 4},
		{"push,", 4},
		{"push\"", 4},
		{"push(", 4},
		{"push)", 4},
		{"push[", 4},
		{"push]", 4},
		{"push// comment", 4},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Push()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestPushInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_p"},
		{"p_"},
		{"_pu"},
		{"pu_"},
		{"_pus"},
		{"pus_"},
		{"_push"},
		{"push_"},
		{"9713p"},
		{"p9713"},
		{"9713pu"},
		{"pu9713"},
		{"9713pus"},
		{"pus9713"},
		{"9713push"},
		{"push9713"},
		{"Zp"},
		{"pZ"},
		{"Zpu"},
		{"puZ"},
		{"Zpus"},
		{"pusZ"},
		{"Zpush"},
		{"pushZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Push()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestPopValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"pop", 3},
		{"pop ", 3},
		{"pop\n", 3},
		{"pop.", 3},
		{"pop:", 3},
		{"pop,", 3},
		{"pop\"", 3},
		{"pop(", 3},
		{"pop)", 3},
		{"pop[", 3},
		{"pop]", 3},
		{"pop// comment", 3},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Pop()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestPopInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_p"},
		{"p_"},
		{"_po"},
		{"po_"},
		{"_pop"},
		{"pop_"},
		{"9713p"},
		{"p9713"},
		{"9713po"},
		{"po9713"},
		{"9713pop"},
		{"pop9713"},
		{"Zp"},
		{"pZ"},
		{"Zpo"},
		{"poZ"},
		{"Zpop"},
		{"popZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Pop()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestQueValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"que", 3},
		{"que ", 3},
		{"que\n", 3},
		{"que.", 3},
		{"que:", 3},
		{"que,", 3},
		{"que\"", 3},
		{"que(", 3},
		{"que)", 3},
		{"que[", 3},
		{"que]", 3},
		{"que// comment", 3},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Que()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestQueInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_q"},
		{"q_"},
		{"_qu"},
		{"qu_"},
		{"_que"},
		{"que_"},
		{"9713q"},
		{"q9713"},
		{"9713qu"},
		{"qu9713"},
		{"9713que"},
		{"que9713"},
		{"Zq"},
		{"qZ"},
		{"Zqu"},
		{"quZ"},
		{"Zque"},
		{"queZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Que()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestDqueValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"dque", 4},
		{"dque ", 4},
		{"dque\n", 4},
		{"dque.", 4},
		{"dque:", 4},
		{"dque,", 4},
		{"dque\"", 4},
		{"dque(", 4},
		{"dque)", 4},
		{"dque[", 4},
		{"dque]", 4},
		{"dque// comment", 4},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Dque()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestDqueInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_d"},
		{"d_"},
		{"_dq"},
		{"dq_"},
		{"_dqu"},
		{"dqu_"},
		{"_dque"},
		{"dque_"},
		{"9713d"},
		{"d9713"},
		{"9713dq"},
		{"dq9713"},
		{"9713dqu"},
		{"dqu9713"},
		{"9713dque"},
		{"dque9713"},
		{"Zd"},
		{"dZ"},
		{"Zdq"},
		{"dqZ"},
		{"Zdqu"},
		{"dquZ"},
		{"Zdque"},
		{"dqueZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Dque()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestInsValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"ins", 3},
		{"ins ", 3},
		{"ins\n", 3},
		{"ins.", 3},
		{"ins:", 3},
		{"ins,", 3},
		{"ins\"", 3},
		{"ins(", 3},
		{"ins)", 3},
		{"ins[", 3},
		{"ins]", 3},
		{"ins// comment", 3},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Ins()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestInsInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_i"},
		{"i_"},
		{"_in"},
		{"in_"},
		{"_ins"},
		{"ins_"},
		{"9713i"},
		{"i9713"},
		{"9713in"},
		{"in9713"},
		{"9713ins"},
		{"ins9713"},
		{"Zi"},
		{"iZ"},
		{"Zin"},
		{"inZ"},
		{"Zins"},
		{"insZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Ins()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestUpdValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"upd", 3},
		{"upd ", 3},
		{"upd\n", 3},
		{"upd.", 3},
		{"upd:", 3},
		{"upd,", 3},
		{"upd\"", 3},
		{"upd(", 3},
		{"upd)", 3},
		{"upd[", 3},
		{"upd]", 3},
		{"upd// comment", 3},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Upd()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestUpdInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_u"},
		{"u_"},
		{"_up"},
		{"up_"},
		{"_upd"},
		{"upd_"},
		{"9713u"},
		{"u9713"},
		{"9713up"},
		{"up9713"},
		{"9713upd"},
		{"upd9713"},
		{"Zu"},
		{"uZ"},
		{"Zup"},
		{"upZ"},
		{"Zupd"},
		{"updZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Upd()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestDelValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"del", 3},
		{"del ", 3},
		{"del\n", 3},
		{"del.", 3},
		{"del:", 3},
		{"del,", 3},
		{"del\"", 3},
		{"del(", 3},
		{"del)", 3},
		{"del[", 3},
		{"del]", 3},
		{"del// comment", 3},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Del()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestDelInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_d"},
		{"d_"},
		{"_de"},
		{"de_"},
		{"_del"},
		{"del_"},
		{"9713d"},
		{"d9713"},
		{"9713de"},
		{"de9713"},
		{"9713del"},
		{"del9713"},
		{"Zd"},
		{"dZ"},
		{"Zde"},
		{"deZ"},
		{"Zdel"},
		{"delZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Del()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestAtValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"at", 2},
		{"at ", 2},
		{"at\n", 2},
		{"at.", 2},
		{"at:", 2},
		{"at,", 2},
		{"at\"", 2},
		{"at(", 2},
		{"at)", 2},
		{"at[", 2},
		{"at]", 2},
		{"at// comment", 2},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.At()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestAtInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_a"},
		{"a_"},
		{"_at"},
		{"at_"},
		{"9713a"},
		{"a9713"},
		{"9713at"},
		{"at9713"},
		{"Za"},
		{"aZ"},
		{"Zat"},
		{"atZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.At()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestInValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"in", 2},
		{"in ", 2},
		{"in\n", 2},
		{"in.", 2},
		{"in:", 2},
		{"in,", 2},
		{"in\"", 2},
		{"in(", 2},
		{"in)", 2},
		{"in[", 2},
		{"in]", 2},
		{"in// comment", 2},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.In()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestInInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_i"},
		{"i_"},
		{"_in"},
		{"in_"},
		{"9713i"},
		{"i9713"},
		{"9713in"},
		{"in9713"},
		{"Zi"},
		{"iZ"},
		{"Zin"},
		{"inZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.In()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestInBndValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"inBnd", 5},
		{"inBnd ", 5},
		{"inBnd\n", 5},
		{"inBnd.", 5},
		{"inBnd:", 5},
		{"inBnd,", 5},
		{"inBnd\"", 5},
		{"inBnd(", 5},
		{"inBnd)", 5},
		{"inBnd[", 5},
		{"inBnd]", 5},
		{"inBnd// comment", 5},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.InBnd()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestInBndInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_i"},
		{"i_"},
		{"_in"},
		{"in_"},
		{"_inB"},
		{"inB_"},
		{"_inBn"},
		{"inBn_"},
		{"_inBnd"},
		{"inBnd_"},
		{"9713i"},
		{"i9713"},
		{"9713in"},
		{"in9713"},
		{"9713inB"},
		{"inB9713"},
		{"9713inBn"},
		{"inBn9713"},
		{"9713inBnd"},
		{"inBnd9713"},
		{"Zi"},
		{"iZ"},
		{"Zin"},
		{"inZ"},
		{"ZinB"},
		{"inBZ"},
		{"ZinBn"},
		{"inBnZ"},
		{"ZinBnd"},
		{"inBndZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.InBnd()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestFromValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"from", 4},
		{"from ", 4},
		{"from\n", 4},
		{"from.", 4},
		{"from:", 4},
		{"from,", 4},
		{"from\"", 4},
		{"from(", 4},
		{"from)", 4},
		{"from[", 4},
		{"from]", 4},
		{"from// comment", 4},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.From()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestFromInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_f"},
		{"f_"},
		{"_fr"},
		{"fr_"},
		{"_fro"},
		{"fro_"},
		{"_from"},
		{"from_"},
		{"9713f"},
		{"f9713"},
		{"9713fr"},
		{"fr9713"},
		{"9713fro"},
		{"fro9713"},
		{"9713from"},
		{"from9713"},
		{"Zf"},
		{"fZ"},
		{"Zfr"},
		{"frZ"},
		{"Zfro"},
		{"froZ"},
		{"Zfrom"},
		{"fromZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.From()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestToValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"to", 2},
		{"to ", 2},
		{"to\n", 2},
		{"to.", 2},
		{"to:", 2},
		{"to,", 2},
		{"to\"", 2},
		{"to(", 2},
		{"to)", 2},
		{"to[", 2},
		{"to]", 2},
		{"to// comment", 2},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.To()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestToInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_t"},
		{"t_"},
		{"_to"},
		{"to_"},
		{"9713t"},
		{"t9713"},
		{"9713to"},
		{"to9713"},
		{"Zt"},
		{"tZ"},
		{"Zto"},
		{"toZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.To()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestFstValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"fst", 3},
		{"fst ", 3},
		{"fst\n", 3},
		{"fst.", 3},
		{"fst:", 3},
		{"fst,", 3},
		{"fst\"", 3},
		{"fst(", 3},
		{"fst)", 3},
		{"fst[", 3},
		{"fst]", 3},
		{"fst// comment", 3},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Fst()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestFstInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_f"},
		{"f_"},
		{"_fs"},
		{"fs_"},
		{"_fst"},
		{"fst_"},
		{"9713f"},
		{"f9713"},
		{"9713fs"},
		{"fs9713"},
		{"9713fst"},
		{"fst9713"},
		{"Zf"},
		{"fZ"},
		{"Zfs"},
		{"fsZ"},
		{"Zfst"},
		{"fstZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Fst()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestMdlValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"mdl", 3},
		{"mdl ", 3},
		{"mdl\n", 3},
		{"mdl.", 3},
		{"mdl:", 3},
		{"mdl,", 3},
		{"mdl\"", 3},
		{"mdl(", 3},
		{"mdl)", 3},
		{"mdl[", 3},
		{"mdl]", 3},
		{"mdl// comment", 3},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Mdl()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestMdlInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_m"},
		{"m_"},
		{"_md"},
		{"md_"},
		{"_mdl"},
		{"mdl_"},
		{"9713m"},
		{"m9713"},
		{"9713md"},
		{"md9713"},
		{"9713mdl"},
		{"mdl9713"},
		{"Zm"},
		{"mZ"},
		{"Zmd"},
		{"mdZ"},
		{"Zmdl"},
		{"mdlZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Mdl()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestLstValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"lst", 3},
		{"lst ", 3},
		{"lst\n", 3},
		{"lst.", 3},
		{"lst:", 3},
		{"lst,", 3},
		{"lst\"", 3},
		{"lst(", 3},
		{"lst)", 3},
		{"lst[", 3},
		{"lst]", 3},
		{"lst// comment", 3},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Lst()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestLstInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_l"},
		{"l_"},
		{"_ls"},
		{"ls_"},
		{"_lst"},
		{"lst_"},
		{"9713l"},
		{"l9713"},
		{"9713ls"},
		{"ls9713"},
		{"9713lst"},
		{"lst9713"},
		{"Zl"},
		{"lZ"},
		{"Zls"},
		{"lsZ"},
		{"Zlst"},
		{"lstZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Lst()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestFstIdxValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"fstIdx", 6},
		{"fstIdx ", 6},
		{"fstIdx\n", 6},
		{"fstIdx.", 6},
		{"fstIdx:", 6},
		{"fstIdx,", 6},
		{"fstIdx\"", 6},
		{"fstIdx(", 6},
		{"fstIdx)", 6},
		{"fstIdx[", 6},
		{"fstIdx]", 6},
		{"fstIdx// comment", 6},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.FstIdx()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestFstIdxInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_f"},
		{"f_"},
		{"_fs"},
		{"fs_"},
		{"_fst"},
		{"fst_"},
		{"_fstI"},
		{"fstI_"},
		{"_fstId"},
		{"fstId_"},
		{"_fstIdx"},
		{"fstIdx_"},
		{"9713f"},
		{"f9713"},
		{"9713fs"},
		{"fs9713"},
		{"9713fst"},
		{"fst9713"},
		{"9713fstI"},
		{"fstI9713"},
		{"9713fstId"},
		{"fstId9713"},
		{"9713fstIdx"},
		{"fstIdx9713"},
		{"Zf"},
		{"fZ"},
		{"Zfs"},
		{"fsZ"},
		{"Zfst"},
		{"fstZ"},
		{"ZfstI"},
		{"fstIZ"},
		{"ZfstId"},
		{"fstIdZ"},
		{"ZfstIdx"},
		{"fstIdxZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.FstIdx()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestMdlIdxValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"mdlIdx", 6},
		{"mdlIdx ", 6},
		{"mdlIdx\n", 6},
		{"mdlIdx.", 6},
		{"mdlIdx:", 6},
		{"mdlIdx,", 6},
		{"mdlIdx\"", 6},
		{"mdlIdx(", 6},
		{"mdlIdx)", 6},
		{"mdlIdx[", 6},
		{"mdlIdx]", 6},
		{"mdlIdx// comment", 6},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.MdlIdx()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestMdlIdxInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_m"},
		{"m_"},
		{"_md"},
		{"md_"},
		{"_mdl"},
		{"mdl_"},
		{"_mdlI"},
		{"mdlI_"},
		{"_mdlId"},
		{"mdlId_"},
		{"_mdlIdx"},
		{"mdlIdx_"},
		{"9713m"},
		{"m9713"},
		{"9713md"},
		{"md9713"},
		{"9713mdl"},
		{"mdl9713"},
		{"9713mdlI"},
		{"mdlI9713"},
		{"9713mdlId"},
		{"mdlId9713"},
		{"9713mdlIdx"},
		{"mdlIdx9713"},
		{"Zm"},
		{"mZ"},
		{"Zmd"},
		{"mdZ"},
		{"Zmdl"},
		{"mdlZ"},
		{"ZmdlI"},
		{"mdlIZ"},
		{"ZmdlId"},
		{"mdlIdZ"},
		{"ZmdlIdx"},
		{"mdlIdxZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.MdlIdx()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestRevValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"rev", 3},
		{"rev ", 3},
		{"rev\n", 3},
		{"rev.", 3},
		{"rev:", 3},
		{"rev,", 3},
		{"rev\"", 3},
		{"rev(", 3},
		{"rev)", 3},
		{"rev[", 3},
		{"rev]", 3},
		{"rev// comment", 3},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Rev()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestRevInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_r"},
		{"r_"},
		{"_re"},
		{"re_"},
		{"_rev"},
		{"rev_"},
		{"9713r"},
		{"r9713"},
		{"9713re"},
		{"re9713"},
		{"9713rev"},
		{"rev9713"},
		{"Zr"},
		{"rZ"},
		{"Zre"},
		{"reZ"},
		{"Zrev"},
		{"revZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Rev()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestSrchIdxEqlValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"srchIdxEql", 10},
		{"srchIdxEql ", 10},
		{"srchIdxEql\n", 10},
		{"srchIdxEql.", 10},
		{"srchIdxEql:", 10},
		{"srchIdxEql,", 10},
		{"srchIdxEql\"", 10},
		{"srchIdxEql(", 10},
		{"srchIdxEql)", 10},
		{"srchIdxEql[", 10},
		{"srchIdxEql]", 10},
		{"srchIdxEql// comment", 10},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.SrchIdxEql()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestSrchIdxEqlInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_s"},
		{"s_"},
		{"_sr"},
		{"sr_"},
		{"_src"},
		{"src_"},
		{"_srch"},
		{"srch_"},
		{"_srchI"},
		{"srchI_"},
		{"_srchId"},
		{"srchId_"},
		{"_srchIdx"},
		{"srchIdx_"},
		{"_srchIdxE"},
		{"srchIdxE_"},
		{"_srchIdxEq"},
		{"srchIdxEq_"},
		{"_srchIdxEql"},
		{"srchIdxEql_"},
		{"9713s"},
		{"s9713"},
		{"9713sr"},
		{"sr9713"},
		{"9713src"},
		{"src9713"},
		{"9713srch"},
		{"srch9713"},
		{"9713srchI"},
		{"srchI9713"},
		{"9713srchId"},
		{"srchId9713"},
		{"9713srchIdx"},
		{"srchIdx9713"},
		{"9713srchIdxE"},
		{"srchIdxE9713"},
		{"9713srchIdxEq"},
		{"srchIdxEq9713"},
		{"9713srchIdxEql"},
		{"srchIdxEql9713"},
		{"Zs"},
		{"sZ"},
		{"Zsr"},
		{"srZ"},
		{"Zsrc"},
		{"srcZ"},
		{"Zsrch"},
		{"srchZ"},
		{"ZsrchI"},
		{"srchIZ"},
		{"ZsrchId"},
		{"srchIdZ"},
		{"ZsrchIdx"},
		{"srchIdxZ"},
		{"ZsrchIdxE"},
		{"srchIdxEZ"},
		{"ZsrchIdxEq"},
		{"srchIdxEqZ"},
		{"ZsrchIdxEql"},
		{"srchIdxEqlZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.SrchIdxEql()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestSrchIdxValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"srchIdx", 7},
		{"srchIdx ", 7},
		{"srchIdx\n", 7},
		{"srchIdx.", 7},
		{"srchIdx:", 7},
		{"srchIdx,", 7},
		{"srchIdx\"", 7},
		{"srchIdx(", 7},
		{"srchIdx)", 7},
		{"srchIdx[", 7},
		{"srchIdx]", 7},
		{"srchIdx// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.SrchIdx()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestSrchIdxInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_s"},
		{"s_"},
		{"_sr"},
		{"sr_"},
		{"_src"},
		{"src_"},
		{"_srch"},
		{"srch_"},
		{"_srchI"},
		{"srchI_"},
		{"_srchId"},
		{"srchId_"},
		{"_srchIdx"},
		{"srchIdx_"},
		{"9713s"},
		{"s9713"},
		{"9713sr"},
		{"sr9713"},
		{"9713src"},
		{"src9713"},
		{"9713srch"},
		{"srch9713"},
		{"9713srchI"},
		{"srchI9713"},
		{"9713srchId"},
		{"srchId9713"},
		{"9713srchIdx"},
		{"srchIdx9713"},
		{"Zs"},
		{"sZ"},
		{"Zsr"},
		{"srZ"},
		{"Zsrc"},
		{"srcZ"},
		{"Zsrch"},
		{"srchZ"},
		{"ZsrchI"},
		{"srchIZ"},
		{"ZsrchId"},
		{"srchIdZ"},
		{"ZsrchIdx"},
		{"srchIdxZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.SrchIdx()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestHasValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"has", 3},
		{"has ", 3},
		{"has\n", 3},
		{"has.", 3},
		{"has:", 3},
		{"has,", 3},
		{"has\"", 3},
		{"has(", 3},
		{"has)", 3},
		{"has[", 3},
		{"has]", 3},
		{"has// comment", 3},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Has()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestHasInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_h"},
		{"h_"},
		{"_ha"},
		{"ha_"},
		{"_has"},
		{"has_"},
		{"9713h"},
		{"h9713"},
		{"9713ha"},
		{"ha9713"},
		{"9713has"},
		{"has9713"},
		{"Zh"},
		{"hZ"},
		{"Zha"},
		{"haZ"},
		{"Zhas"},
		{"hasZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Has()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestSrtAscValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"srtAsc", 6},
		{"srtAsc ", 6},
		{"srtAsc\n", 6},
		{"srtAsc.", 6},
		{"srtAsc:", 6},
		{"srtAsc,", 6},
		{"srtAsc\"", 6},
		{"srtAsc(", 6},
		{"srtAsc)", 6},
		{"srtAsc[", 6},
		{"srtAsc]", 6},
		{"srtAsc// comment", 6},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.SrtAsc()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestSrtAscInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_s"},
		{"s_"},
		{"_sr"},
		{"sr_"},
		{"_srt"},
		{"srt_"},
		{"_srtA"},
		{"srtA_"},
		{"_srtAs"},
		{"srtAs_"},
		{"_srtAsc"},
		{"srtAsc_"},
		{"9713s"},
		{"s9713"},
		{"9713sr"},
		{"sr9713"},
		{"9713srt"},
		{"srt9713"},
		{"9713srtA"},
		{"srtA9713"},
		{"9713srtAs"},
		{"srtAs9713"},
		{"9713srtAsc"},
		{"srtAsc9713"},
		{"Zs"},
		{"sZ"},
		{"Zsr"},
		{"srZ"},
		{"Zsrt"},
		{"srtZ"},
		{"ZsrtA"},
		{"srtAZ"},
		{"ZsrtAs"},
		{"srtAsZ"},
		{"ZsrtAsc"},
		{"srtAscZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.SrtAsc()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestSrtDscValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"srtDsc", 6},
		{"srtDsc ", 6},
		{"srtDsc\n", 6},
		{"srtDsc.", 6},
		{"srtDsc:", 6},
		{"srtDsc,", 6},
		{"srtDsc\"", 6},
		{"srtDsc(", 6},
		{"srtDsc)", 6},
		{"srtDsc[", 6},
		{"srtDsc]", 6},
		{"srtDsc// comment", 6},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.SrtDsc()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestSrtDscInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_s"},
		{"s_"},
		{"_sr"},
		{"sr_"},
		{"_srt"},
		{"srt_"},
		{"_srtD"},
		{"srtD_"},
		{"_srtDs"},
		{"srtDs_"},
		{"_srtDsc"},
		{"srtDsc_"},
		{"9713s"},
		{"s9713"},
		{"9713sr"},
		{"sr9713"},
		{"9713srt"},
		{"srt9713"},
		{"9713srtD"},
		{"srtD9713"},
		{"9713srtDs"},
		{"srtDs9713"},
		{"9713srtDsc"},
		{"srtDsc9713"},
		{"Zs"},
		{"sZ"},
		{"Zsr"},
		{"srZ"},
		{"Zsrt"},
		{"srtZ"},
		{"ZsrtD"},
		{"srtDZ"},
		{"ZsrtDs"},
		{"srtDsZ"},
		{"ZsrtDsc"},
		{"srtDscZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.SrtDsc()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestUnaPosValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"unaPos", 6},
		{"unaPos ", 6},
		{"unaPos\n", 6},
		{"unaPos.", 6},
		{"unaPos:", 6},
		{"unaPos,", 6},
		{"unaPos\"", 6},
		{"unaPos(", 6},
		{"unaPos)", 6},
		{"unaPos[", 6},
		{"unaPos]", 6},
		{"unaPos// comment", 6},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.UnaPos()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestUnaPosInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_u"},
		{"u_"},
		{"_un"},
		{"un_"},
		{"_una"},
		{"una_"},
		{"_unaP"},
		{"unaP_"},
		{"_unaPo"},
		{"unaPo_"},
		{"_unaPos"},
		{"unaPos_"},
		{"9713u"},
		{"u9713"},
		{"9713un"},
		{"un9713"},
		{"9713una"},
		{"una9713"},
		{"9713unaP"},
		{"unaP9713"},
		{"9713unaPo"},
		{"unaPo9713"},
		{"9713unaPos"},
		{"unaPos9713"},
		{"Zu"},
		{"uZ"},
		{"Zun"},
		{"unZ"},
		{"Zuna"},
		{"unaZ"},
		{"ZunaP"},
		{"unaPZ"},
		{"ZunaPo"},
		{"unaPoZ"},
		{"ZunaPos"},
		{"unaPosZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.UnaPos()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestUnaNegValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"unaNeg", 6},
		{"unaNeg ", 6},
		{"unaNeg\n", 6},
		{"unaNeg.", 6},
		{"unaNeg:", 6},
		{"unaNeg,", 6},
		{"unaNeg\"", 6},
		{"unaNeg(", 6},
		{"unaNeg)", 6},
		{"unaNeg[", 6},
		{"unaNeg]", 6},
		{"unaNeg// comment", 6},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.UnaNeg()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestUnaNegInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_u"},
		{"u_"},
		{"_un"},
		{"un_"},
		{"_una"},
		{"una_"},
		{"_unaN"},
		{"unaN_"},
		{"_unaNe"},
		{"unaNe_"},
		{"_unaNeg"},
		{"unaNeg_"},
		{"9713u"},
		{"u9713"},
		{"9713un"},
		{"un9713"},
		{"9713una"},
		{"una9713"},
		{"9713unaN"},
		{"unaN9713"},
		{"9713unaNe"},
		{"unaNe9713"},
		{"9713unaNeg"},
		{"unaNeg9713"},
		{"Zu"},
		{"uZ"},
		{"Zun"},
		{"unZ"},
		{"Zuna"},
		{"unaZ"},
		{"ZunaN"},
		{"unaNZ"},
		{"ZunaNe"},
		{"unaNeZ"},
		{"ZunaNeg"},
		{"unaNegZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.UnaNeg()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestUnaInvValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"unaInv", 6},
		{"unaInv ", 6},
		{"unaInv\n", 6},
		{"unaInv.", 6},
		{"unaInv:", 6},
		{"unaInv,", 6},
		{"unaInv\"", 6},
		{"unaInv(", 6},
		{"unaInv)", 6},
		{"unaInv[", 6},
		{"unaInv]", 6},
		{"unaInv// comment", 6},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.UnaInv()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestUnaInvInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_u"},
		{"u_"},
		{"_un"},
		{"un_"},
		{"_una"},
		{"una_"},
		{"_unaI"},
		{"unaI_"},
		{"_unaIn"},
		{"unaIn_"},
		{"_unaInv"},
		{"unaInv_"},
		{"9713u"},
		{"u9713"},
		{"9713un"},
		{"un9713"},
		{"9713una"},
		{"una9713"},
		{"9713unaI"},
		{"unaI9713"},
		{"9713unaIn"},
		{"unaIn9713"},
		{"9713unaInv"},
		{"unaInv9713"},
		{"Zu"},
		{"uZ"},
		{"Zun"},
		{"unZ"},
		{"Zuna"},
		{"unaZ"},
		{"ZunaI"},
		{"unaIZ"},
		{"ZunaIn"},
		{"unaInZ"},
		{"ZunaInv"},
		{"unaInvZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.UnaInv()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestUnaSqrValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"unaSqr", 6},
		{"unaSqr ", 6},
		{"unaSqr\n", 6},
		{"unaSqr.", 6},
		{"unaSqr:", 6},
		{"unaSqr,", 6},
		{"unaSqr\"", 6},
		{"unaSqr(", 6},
		{"unaSqr)", 6},
		{"unaSqr[", 6},
		{"unaSqr]", 6},
		{"unaSqr// comment", 6},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.UnaSqr()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestUnaSqrInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_u"},
		{"u_"},
		{"_un"},
		{"un_"},
		{"_una"},
		{"una_"},
		{"_unaS"},
		{"unaS_"},
		{"_unaSq"},
		{"unaSq_"},
		{"_unaSqr"},
		{"unaSqr_"},
		{"9713u"},
		{"u9713"},
		{"9713un"},
		{"un9713"},
		{"9713una"},
		{"una9713"},
		{"9713unaS"},
		{"unaS9713"},
		{"9713unaSq"},
		{"unaSq9713"},
		{"9713unaSqr"},
		{"unaSqr9713"},
		{"Zu"},
		{"uZ"},
		{"Zun"},
		{"unZ"},
		{"Zuna"},
		{"unaZ"},
		{"ZunaS"},
		{"unaSZ"},
		{"ZunaSq"},
		{"unaSqZ"},
		{"ZunaSqr"},
		{"unaSqrZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.UnaSqr()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestUnaSqrtValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"unaSqrt", 7},
		{"unaSqrt ", 7},
		{"unaSqrt\n", 7},
		{"unaSqrt.", 7},
		{"unaSqrt:", 7},
		{"unaSqrt,", 7},
		{"unaSqrt\"", 7},
		{"unaSqrt(", 7},
		{"unaSqrt)", 7},
		{"unaSqrt[", 7},
		{"unaSqrt]", 7},
		{"unaSqrt// comment", 7},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.UnaSqrt()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestUnaSqrtInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_u"},
		{"u_"},
		{"_un"},
		{"un_"},
		{"_una"},
		{"una_"},
		{"_unaS"},
		{"unaS_"},
		{"_unaSq"},
		{"unaSq_"},
		{"_unaSqr"},
		{"unaSqr_"},
		{"_unaSqrt"},
		{"unaSqrt_"},
		{"9713u"},
		{"u9713"},
		{"9713un"},
		{"un9713"},
		{"9713una"},
		{"una9713"},
		{"9713unaS"},
		{"unaS9713"},
		{"9713unaSq"},
		{"unaSq9713"},
		{"9713unaSqr"},
		{"unaSqr9713"},
		{"9713unaSqrt"},
		{"unaSqrt9713"},
		{"Zu"},
		{"uZ"},
		{"Zun"},
		{"unZ"},
		{"Zuna"},
		{"unaZ"},
		{"ZunaS"},
		{"unaSZ"},
		{"ZunaSq"},
		{"unaSqZ"},
		{"ZunaSqr"},
		{"unaSqrZ"},
		{"ZunaSqrt"},
		{"unaSqrtZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.UnaSqrt()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestSclAddValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"sclAdd", 6},
		{"sclAdd ", 6},
		{"sclAdd\n", 6},
		{"sclAdd.", 6},
		{"sclAdd:", 6},
		{"sclAdd,", 6},
		{"sclAdd\"", 6},
		{"sclAdd(", 6},
		{"sclAdd)", 6},
		{"sclAdd[", 6},
		{"sclAdd]", 6},
		{"sclAdd// comment", 6},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.SclAdd()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestSclAddInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_s"},
		{"s_"},
		{"_sc"},
		{"sc_"},
		{"_scl"},
		{"scl_"},
		{"_sclA"},
		{"sclA_"},
		{"_sclAd"},
		{"sclAd_"},
		{"_sclAdd"},
		{"sclAdd_"},
		{"9713s"},
		{"s9713"},
		{"9713sc"},
		{"sc9713"},
		{"9713scl"},
		{"scl9713"},
		{"9713sclA"},
		{"sclA9713"},
		{"9713sclAd"},
		{"sclAd9713"},
		{"9713sclAdd"},
		{"sclAdd9713"},
		{"Zs"},
		{"sZ"},
		{"Zsc"},
		{"scZ"},
		{"Zscl"},
		{"sclZ"},
		{"ZsclA"},
		{"sclAZ"},
		{"ZsclAd"},
		{"sclAdZ"},
		{"ZsclAdd"},
		{"sclAddZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.SclAdd()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestSclSubValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"sclSub", 6},
		{"sclSub ", 6},
		{"sclSub\n", 6},
		{"sclSub.", 6},
		{"sclSub:", 6},
		{"sclSub,", 6},
		{"sclSub\"", 6},
		{"sclSub(", 6},
		{"sclSub)", 6},
		{"sclSub[", 6},
		{"sclSub]", 6},
		{"sclSub// comment", 6},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.SclSub()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestSclSubInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_s"},
		{"s_"},
		{"_sc"},
		{"sc_"},
		{"_scl"},
		{"scl_"},
		{"_sclS"},
		{"sclS_"},
		{"_sclSu"},
		{"sclSu_"},
		{"_sclSub"},
		{"sclSub_"},
		{"9713s"},
		{"s9713"},
		{"9713sc"},
		{"sc9713"},
		{"9713scl"},
		{"scl9713"},
		{"9713sclS"},
		{"sclS9713"},
		{"9713sclSu"},
		{"sclSu9713"},
		{"9713sclSub"},
		{"sclSub9713"},
		{"Zs"},
		{"sZ"},
		{"Zsc"},
		{"scZ"},
		{"Zscl"},
		{"sclZ"},
		{"ZsclS"},
		{"sclSZ"},
		{"ZsclSu"},
		{"sclSuZ"},
		{"ZsclSub"},
		{"sclSubZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.SclSub()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestSclMulValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"sclMul", 6},
		{"sclMul ", 6},
		{"sclMul\n", 6},
		{"sclMul.", 6},
		{"sclMul:", 6},
		{"sclMul,", 6},
		{"sclMul\"", 6},
		{"sclMul(", 6},
		{"sclMul)", 6},
		{"sclMul[", 6},
		{"sclMul]", 6},
		{"sclMul// comment", 6},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.SclMul()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestSclMulInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_s"},
		{"s_"},
		{"_sc"},
		{"sc_"},
		{"_scl"},
		{"scl_"},
		{"_sclM"},
		{"sclM_"},
		{"_sclMu"},
		{"sclMu_"},
		{"_sclMul"},
		{"sclMul_"},
		{"9713s"},
		{"s9713"},
		{"9713sc"},
		{"sc9713"},
		{"9713scl"},
		{"scl9713"},
		{"9713sclM"},
		{"sclM9713"},
		{"9713sclMu"},
		{"sclMu9713"},
		{"9713sclMul"},
		{"sclMul9713"},
		{"Zs"},
		{"sZ"},
		{"Zsc"},
		{"scZ"},
		{"Zscl"},
		{"sclZ"},
		{"ZsclM"},
		{"sclMZ"},
		{"ZsclMu"},
		{"sclMuZ"},
		{"ZsclMul"},
		{"sclMulZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.SclMul()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestSclDivValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"sclDiv", 6},
		{"sclDiv ", 6},
		{"sclDiv\n", 6},
		{"sclDiv.", 6},
		{"sclDiv:", 6},
		{"sclDiv,", 6},
		{"sclDiv\"", 6},
		{"sclDiv(", 6},
		{"sclDiv)", 6},
		{"sclDiv[", 6},
		{"sclDiv]", 6},
		{"sclDiv// comment", 6},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.SclDiv()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestSclDivInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_s"},
		{"s_"},
		{"_sc"},
		{"sc_"},
		{"_scl"},
		{"scl_"},
		{"_sclD"},
		{"sclD_"},
		{"_sclDi"},
		{"sclDi_"},
		{"_sclDiv"},
		{"sclDiv_"},
		{"9713s"},
		{"s9713"},
		{"9713sc"},
		{"sc9713"},
		{"9713scl"},
		{"scl9713"},
		{"9713sclD"},
		{"sclD9713"},
		{"9713sclDi"},
		{"sclDi9713"},
		{"9713sclDiv"},
		{"sclDiv9713"},
		{"Zs"},
		{"sZ"},
		{"Zsc"},
		{"scZ"},
		{"Zscl"},
		{"sclZ"},
		{"ZsclD"},
		{"sclDZ"},
		{"ZsclDi"},
		{"sclDiZ"},
		{"ZsclDiv"},
		{"sclDivZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.SclDiv()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestSclRemValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"sclRem", 6},
		{"sclRem ", 6},
		{"sclRem\n", 6},
		{"sclRem.", 6},
		{"sclRem:", 6},
		{"sclRem,", 6},
		{"sclRem\"", 6},
		{"sclRem(", 6},
		{"sclRem)", 6},
		{"sclRem[", 6},
		{"sclRem]", 6},
		{"sclRem// comment", 6},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.SclRem()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestSclRemInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_s"},
		{"s_"},
		{"_sc"},
		{"sc_"},
		{"_scl"},
		{"scl_"},
		{"_sclR"},
		{"sclR_"},
		{"_sclRe"},
		{"sclRe_"},
		{"_sclRem"},
		{"sclRem_"},
		{"9713s"},
		{"s9713"},
		{"9713sc"},
		{"sc9713"},
		{"9713scl"},
		{"scl9713"},
		{"9713sclR"},
		{"sclR9713"},
		{"9713sclRe"},
		{"sclRe9713"},
		{"9713sclRem"},
		{"sclRem9713"},
		{"Zs"},
		{"sZ"},
		{"Zsc"},
		{"scZ"},
		{"Zscl"},
		{"sclZ"},
		{"ZsclR"},
		{"sclRZ"},
		{"ZsclRe"},
		{"sclReZ"},
		{"ZsclRem"},
		{"sclRemZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.SclRem()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestSclPowValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"sclPow", 6},
		{"sclPow ", 6},
		{"sclPow\n", 6},
		{"sclPow.", 6},
		{"sclPow:", 6},
		{"sclPow,", 6},
		{"sclPow\"", 6},
		{"sclPow(", 6},
		{"sclPow)", 6},
		{"sclPow[", 6},
		{"sclPow]", 6},
		{"sclPow// comment", 6},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.SclPow()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestSclPowInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_s"},
		{"s_"},
		{"_sc"},
		{"sc_"},
		{"_scl"},
		{"scl_"},
		{"_sclP"},
		{"sclP_"},
		{"_sclPo"},
		{"sclPo_"},
		{"_sclPow"},
		{"sclPow_"},
		{"9713s"},
		{"s9713"},
		{"9713sc"},
		{"sc9713"},
		{"9713scl"},
		{"scl9713"},
		{"9713sclP"},
		{"sclP9713"},
		{"9713sclPo"},
		{"sclPo9713"},
		{"9713sclPow"},
		{"sclPow9713"},
		{"Zs"},
		{"sZ"},
		{"Zsc"},
		{"scZ"},
		{"Zscl"},
		{"sclZ"},
		{"ZsclP"},
		{"sclPZ"},
		{"ZsclPo"},
		{"sclPoZ"},
		{"ZsclPow"},
		{"sclPowZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.SclPow()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestSclMinValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"sclMin", 6},
		{"sclMin ", 6},
		{"sclMin\n", 6},
		{"sclMin.", 6},
		{"sclMin:", 6},
		{"sclMin,", 6},
		{"sclMin\"", 6},
		{"sclMin(", 6},
		{"sclMin)", 6},
		{"sclMin[", 6},
		{"sclMin]", 6},
		{"sclMin// comment", 6},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.SclMin()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestSclMinInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_s"},
		{"s_"},
		{"_sc"},
		{"sc_"},
		{"_scl"},
		{"scl_"},
		{"_sclM"},
		{"sclM_"},
		{"_sclMi"},
		{"sclMi_"},
		{"_sclMin"},
		{"sclMin_"},
		{"9713s"},
		{"s9713"},
		{"9713sc"},
		{"sc9713"},
		{"9713scl"},
		{"scl9713"},
		{"9713sclM"},
		{"sclM9713"},
		{"9713sclMi"},
		{"sclMi9713"},
		{"9713sclMin"},
		{"sclMin9713"},
		{"Zs"},
		{"sZ"},
		{"Zsc"},
		{"scZ"},
		{"Zscl"},
		{"sclZ"},
		{"ZsclM"},
		{"sclMZ"},
		{"ZsclMi"},
		{"sclMiZ"},
		{"ZsclMin"},
		{"sclMinZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.SclMin()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestSclMaxValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"sclMax", 6},
		{"sclMax ", 6},
		{"sclMax\n", 6},
		{"sclMax.", 6},
		{"sclMax:", 6},
		{"sclMax,", 6},
		{"sclMax\"", 6},
		{"sclMax(", 6},
		{"sclMax)", 6},
		{"sclMax[", 6},
		{"sclMax]", 6},
		{"sclMax// comment", 6},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.SclMax()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestSclMaxInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_s"},
		{"s_"},
		{"_sc"},
		{"sc_"},
		{"_scl"},
		{"scl_"},
		{"_sclM"},
		{"sclM_"},
		{"_sclMa"},
		{"sclMa_"},
		{"_sclMax"},
		{"sclMax_"},
		{"9713s"},
		{"s9713"},
		{"9713sc"},
		{"sc9713"},
		{"9713scl"},
		{"scl9713"},
		{"9713sclM"},
		{"sclM9713"},
		{"9713sclMa"},
		{"sclMa9713"},
		{"9713sclMax"},
		{"sclMax9713"},
		{"Zs"},
		{"sZ"},
		{"Zsc"},
		{"scZ"},
		{"Zscl"},
		{"sclZ"},
		{"ZsclM"},
		{"sclMZ"},
		{"ZsclMa"},
		{"sclMaZ"},
		{"ZsclMax"},
		{"sclMaxZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.SclMax()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestCntEqlValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"cntEql", 6},
		{"cntEql ", 6},
		{"cntEql\n", 6},
		{"cntEql.", 6},
		{"cntEql:", 6},
		{"cntEql,", 6},
		{"cntEql\"", 6},
		{"cntEql(", 6},
		{"cntEql)", 6},
		{"cntEql[", 6},
		{"cntEql]", 6},
		{"cntEql// comment", 6},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.CntEql()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestCntEqlInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_c"},
		{"c_"},
		{"_cn"},
		{"cn_"},
		{"_cnt"},
		{"cnt_"},
		{"_cntE"},
		{"cntE_"},
		{"_cntEq"},
		{"cntEq_"},
		{"_cntEql"},
		{"cntEql_"},
		{"9713c"},
		{"c9713"},
		{"9713cn"},
		{"cn9713"},
		{"9713cnt"},
		{"cnt9713"},
		{"9713cntE"},
		{"cntE9713"},
		{"9713cntEq"},
		{"cntEq9713"},
		{"9713cntEql"},
		{"cntEql9713"},
		{"Zc"},
		{"cZ"},
		{"Zcn"},
		{"cnZ"},
		{"Zcnt"},
		{"cntZ"},
		{"ZcntE"},
		{"cntEZ"},
		{"ZcntEq"},
		{"cntEqZ"},
		{"ZcntEql"},
		{"cntEqlZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.CntEql()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestCntNeqValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"cntNeq", 6},
		{"cntNeq ", 6},
		{"cntNeq\n", 6},
		{"cntNeq.", 6},
		{"cntNeq:", 6},
		{"cntNeq,", 6},
		{"cntNeq\"", 6},
		{"cntNeq(", 6},
		{"cntNeq)", 6},
		{"cntNeq[", 6},
		{"cntNeq]", 6},
		{"cntNeq// comment", 6},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.CntNeq()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestCntNeqInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_c"},
		{"c_"},
		{"_cn"},
		{"cn_"},
		{"_cnt"},
		{"cnt_"},
		{"_cntN"},
		{"cntN_"},
		{"_cntNe"},
		{"cntNe_"},
		{"_cntNeq"},
		{"cntNeq_"},
		{"9713c"},
		{"c9713"},
		{"9713cn"},
		{"cn9713"},
		{"9713cnt"},
		{"cnt9713"},
		{"9713cntN"},
		{"cntN9713"},
		{"9713cntNe"},
		{"cntNe9713"},
		{"9713cntNeq"},
		{"cntNeq9713"},
		{"Zc"},
		{"cZ"},
		{"Zcn"},
		{"cnZ"},
		{"Zcnt"},
		{"cntZ"},
		{"ZcntN"},
		{"cntNZ"},
		{"ZcntNe"},
		{"cntNeZ"},
		{"ZcntNeq"},
		{"cntNeqZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.CntNeq()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestCntLssValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"cntLss", 6},
		{"cntLss ", 6},
		{"cntLss\n", 6},
		{"cntLss.", 6},
		{"cntLss:", 6},
		{"cntLss,", 6},
		{"cntLss\"", 6},
		{"cntLss(", 6},
		{"cntLss)", 6},
		{"cntLss[", 6},
		{"cntLss]", 6},
		{"cntLss// comment", 6},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.CntLss()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestCntLssInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_c"},
		{"c_"},
		{"_cn"},
		{"cn_"},
		{"_cnt"},
		{"cnt_"},
		{"_cntL"},
		{"cntL_"},
		{"_cntLs"},
		{"cntLs_"},
		{"_cntLss"},
		{"cntLss_"},
		{"9713c"},
		{"c9713"},
		{"9713cn"},
		{"cn9713"},
		{"9713cnt"},
		{"cnt9713"},
		{"9713cntL"},
		{"cntL9713"},
		{"9713cntLs"},
		{"cntLs9713"},
		{"9713cntLss"},
		{"cntLss9713"},
		{"Zc"},
		{"cZ"},
		{"Zcn"},
		{"cnZ"},
		{"Zcnt"},
		{"cntZ"},
		{"ZcntL"},
		{"cntLZ"},
		{"ZcntLs"},
		{"cntLsZ"},
		{"ZcntLss"},
		{"cntLssZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.CntLss()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestCntGtrValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"cntGtr", 6},
		{"cntGtr ", 6},
		{"cntGtr\n", 6},
		{"cntGtr.", 6},
		{"cntGtr:", 6},
		{"cntGtr,", 6},
		{"cntGtr\"", 6},
		{"cntGtr(", 6},
		{"cntGtr)", 6},
		{"cntGtr[", 6},
		{"cntGtr]", 6},
		{"cntGtr// comment", 6},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.CntGtr()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestCntGtrInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_c"},
		{"c_"},
		{"_cn"},
		{"cn_"},
		{"_cnt"},
		{"cnt_"},
		{"_cntG"},
		{"cntG_"},
		{"_cntGt"},
		{"cntGt_"},
		{"_cntGtr"},
		{"cntGtr_"},
		{"9713c"},
		{"c9713"},
		{"9713cn"},
		{"cn9713"},
		{"9713cnt"},
		{"cnt9713"},
		{"9713cntG"},
		{"cntG9713"},
		{"9713cntGt"},
		{"cntGt9713"},
		{"9713cntGtr"},
		{"cntGtr9713"},
		{"Zc"},
		{"cZ"},
		{"Zcn"},
		{"cnZ"},
		{"Zcnt"},
		{"cntZ"},
		{"ZcntG"},
		{"cntGZ"},
		{"ZcntGt"},
		{"cntGtZ"},
		{"ZcntGtr"},
		{"cntGtrZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.CntGtr()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestCntLeqValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"cntLeq", 6},
		{"cntLeq ", 6},
		{"cntLeq\n", 6},
		{"cntLeq.", 6},
		{"cntLeq:", 6},
		{"cntLeq,", 6},
		{"cntLeq\"", 6},
		{"cntLeq(", 6},
		{"cntLeq)", 6},
		{"cntLeq[", 6},
		{"cntLeq]", 6},
		{"cntLeq// comment", 6},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.CntLeq()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestCntLeqInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_c"},
		{"c_"},
		{"_cn"},
		{"cn_"},
		{"_cnt"},
		{"cnt_"},
		{"_cntL"},
		{"cntL_"},
		{"_cntLe"},
		{"cntLe_"},
		{"_cntLeq"},
		{"cntLeq_"},
		{"9713c"},
		{"c9713"},
		{"9713cn"},
		{"cn9713"},
		{"9713cnt"},
		{"cnt9713"},
		{"9713cntL"},
		{"cntL9713"},
		{"9713cntLe"},
		{"cntLe9713"},
		{"9713cntLeq"},
		{"cntLeq9713"},
		{"Zc"},
		{"cZ"},
		{"Zcn"},
		{"cnZ"},
		{"Zcnt"},
		{"cntZ"},
		{"ZcntL"},
		{"cntLZ"},
		{"ZcntLe"},
		{"cntLeZ"},
		{"ZcntLeq"},
		{"cntLeqZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.CntLeq()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestCntGeqValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"cntGeq", 6},
		{"cntGeq ", 6},
		{"cntGeq\n", 6},
		{"cntGeq.", 6},
		{"cntGeq:", 6},
		{"cntGeq,", 6},
		{"cntGeq\"", 6},
		{"cntGeq(", 6},
		{"cntGeq)", 6},
		{"cntGeq[", 6},
		{"cntGeq]", 6},
		{"cntGeq// comment", 6},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.CntGeq()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestCntGeqInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_c"},
		{"c_"},
		{"_cn"},
		{"cn_"},
		{"_cnt"},
		{"cnt_"},
		{"_cntG"},
		{"cntG_"},
		{"_cntGe"},
		{"cntGe_"},
		{"_cntGeq"},
		{"cntGeq_"},
		{"9713c"},
		{"c9713"},
		{"9713cn"},
		{"cn9713"},
		{"9713cnt"},
		{"cnt9713"},
		{"9713cntG"},
		{"cntG9713"},
		{"9713cntGe"},
		{"cntGe9713"},
		{"9713cntGeq"},
		{"cntGeq9713"},
		{"Zc"},
		{"cZ"},
		{"Zcn"},
		{"cnZ"},
		{"Zcnt"},
		{"cntZ"},
		{"ZcntG"},
		{"cntGZ"},
		{"ZcntGe"},
		{"cntGeZ"},
		{"ZcntGeq"},
		{"cntGeqZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.CntGeq()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestInrAddValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"inrAdd", 6},
		{"inrAdd ", 6},
		{"inrAdd\n", 6},
		{"inrAdd.", 6},
		{"inrAdd:", 6},
		{"inrAdd,", 6},
		{"inrAdd\"", 6},
		{"inrAdd(", 6},
		{"inrAdd)", 6},
		{"inrAdd[", 6},
		{"inrAdd]", 6},
		{"inrAdd// comment", 6},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.InrAdd()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestInrAddInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_i"},
		{"i_"},
		{"_in"},
		{"in_"},
		{"_inr"},
		{"inr_"},
		{"_inrA"},
		{"inrA_"},
		{"_inrAd"},
		{"inrAd_"},
		{"_inrAdd"},
		{"inrAdd_"},
		{"9713i"},
		{"i9713"},
		{"9713in"},
		{"in9713"},
		{"9713inr"},
		{"inr9713"},
		{"9713inrA"},
		{"inrA9713"},
		{"9713inrAd"},
		{"inrAd9713"},
		{"9713inrAdd"},
		{"inrAdd9713"},
		{"Zi"},
		{"iZ"},
		{"Zin"},
		{"inZ"},
		{"Zinr"},
		{"inrZ"},
		{"ZinrA"},
		{"inrAZ"},
		{"ZinrAd"},
		{"inrAdZ"},
		{"ZinrAdd"},
		{"inrAddZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.InrAdd()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestInrSubValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"inrSub", 6},
		{"inrSub ", 6},
		{"inrSub\n", 6},
		{"inrSub.", 6},
		{"inrSub:", 6},
		{"inrSub,", 6},
		{"inrSub\"", 6},
		{"inrSub(", 6},
		{"inrSub)", 6},
		{"inrSub[", 6},
		{"inrSub]", 6},
		{"inrSub// comment", 6},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.InrSub()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestInrSubInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_i"},
		{"i_"},
		{"_in"},
		{"in_"},
		{"_inr"},
		{"inr_"},
		{"_inrS"},
		{"inrS_"},
		{"_inrSu"},
		{"inrSu_"},
		{"_inrSub"},
		{"inrSub_"},
		{"9713i"},
		{"i9713"},
		{"9713in"},
		{"in9713"},
		{"9713inr"},
		{"inr9713"},
		{"9713inrS"},
		{"inrS9713"},
		{"9713inrSu"},
		{"inrSu9713"},
		{"9713inrSub"},
		{"inrSub9713"},
		{"Zi"},
		{"iZ"},
		{"Zin"},
		{"inZ"},
		{"Zinr"},
		{"inrZ"},
		{"ZinrS"},
		{"inrSZ"},
		{"ZinrSu"},
		{"inrSuZ"},
		{"ZinrSub"},
		{"inrSubZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.InrSub()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestInrMulValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"inrMul", 6},
		{"inrMul ", 6},
		{"inrMul\n", 6},
		{"inrMul.", 6},
		{"inrMul:", 6},
		{"inrMul,", 6},
		{"inrMul\"", 6},
		{"inrMul(", 6},
		{"inrMul)", 6},
		{"inrMul[", 6},
		{"inrMul]", 6},
		{"inrMul// comment", 6},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.InrMul()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestInrMulInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_i"},
		{"i_"},
		{"_in"},
		{"in_"},
		{"_inr"},
		{"inr_"},
		{"_inrM"},
		{"inrM_"},
		{"_inrMu"},
		{"inrMu_"},
		{"_inrMul"},
		{"inrMul_"},
		{"9713i"},
		{"i9713"},
		{"9713in"},
		{"in9713"},
		{"9713inr"},
		{"inr9713"},
		{"9713inrM"},
		{"inrM9713"},
		{"9713inrMu"},
		{"inrMu9713"},
		{"9713inrMul"},
		{"inrMul9713"},
		{"Zi"},
		{"iZ"},
		{"Zin"},
		{"inZ"},
		{"Zinr"},
		{"inrZ"},
		{"ZinrM"},
		{"inrMZ"},
		{"ZinrMu"},
		{"inrMuZ"},
		{"ZinrMul"},
		{"inrMulZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.InrMul()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestInrDivValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"inrDiv", 6},
		{"inrDiv ", 6},
		{"inrDiv\n", 6},
		{"inrDiv.", 6},
		{"inrDiv:", 6},
		{"inrDiv,", 6},
		{"inrDiv\"", 6},
		{"inrDiv(", 6},
		{"inrDiv)", 6},
		{"inrDiv[", 6},
		{"inrDiv]", 6},
		{"inrDiv// comment", 6},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.InrDiv()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestInrDivInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_i"},
		{"i_"},
		{"_in"},
		{"in_"},
		{"_inr"},
		{"inr_"},
		{"_inrD"},
		{"inrD_"},
		{"_inrDi"},
		{"inrDi_"},
		{"_inrDiv"},
		{"inrDiv_"},
		{"9713i"},
		{"i9713"},
		{"9713in"},
		{"in9713"},
		{"9713inr"},
		{"inr9713"},
		{"9713inrD"},
		{"inrD9713"},
		{"9713inrDi"},
		{"inrDi9713"},
		{"9713inrDiv"},
		{"inrDiv9713"},
		{"Zi"},
		{"iZ"},
		{"Zin"},
		{"inZ"},
		{"Zinr"},
		{"inrZ"},
		{"ZinrD"},
		{"inrDZ"},
		{"ZinrDi"},
		{"inrDiZ"},
		{"ZinrDiv"},
		{"inrDivZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.InrDiv()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestInrRemValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"inrRem", 6},
		{"inrRem ", 6},
		{"inrRem\n", 6},
		{"inrRem.", 6},
		{"inrRem:", 6},
		{"inrRem,", 6},
		{"inrRem\"", 6},
		{"inrRem(", 6},
		{"inrRem)", 6},
		{"inrRem[", 6},
		{"inrRem]", 6},
		{"inrRem// comment", 6},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.InrRem()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestInrRemInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_i"},
		{"i_"},
		{"_in"},
		{"in_"},
		{"_inr"},
		{"inr_"},
		{"_inrR"},
		{"inrR_"},
		{"_inrRe"},
		{"inrRe_"},
		{"_inrRem"},
		{"inrRem_"},
		{"9713i"},
		{"i9713"},
		{"9713in"},
		{"in9713"},
		{"9713inr"},
		{"inr9713"},
		{"9713inrR"},
		{"inrR9713"},
		{"9713inrRe"},
		{"inrRe9713"},
		{"9713inrRem"},
		{"inrRem9713"},
		{"Zi"},
		{"iZ"},
		{"Zin"},
		{"inZ"},
		{"Zinr"},
		{"inrZ"},
		{"ZinrR"},
		{"inrRZ"},
		{"ZinrRe"},
		{"inrReZ"},
		{"ZinrRem"},
		{"inrRemZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.InrRem()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestInrPowValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"inrPow", 6},
		{"inrPow ", 6},
		{"inrPow\n", 6},
		{"inrPow.", 6},
		{"inrPow:", 6},
		{"inrPow,", 6},
		{"inrPow\"", 6},
		{"inrPow(", 6},
		{"inrPow)", 6},
		{"inrPow[", 6},
		{"inrPow]", 6},
		{"inrPow// comment", 6},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.InrPow()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestInrPowInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_i"},
		{"i_"},
		{"_in"},
		{"in_"},
		{"_inr"},
		{"inr_"},
		{"_inrP"},
		{"inrP_"},
		{"_inrPo"},
		{"inrPo_"},
		{"_inrPow"},
		{"inrPow_"},
		{"9713i"},
		{"i9713"},
		{"9713in"},
		{"in9713"},
		{"9713inr"},
		{"inr9713"},
		{"9713inrP"},
		{"inrP9713"},
		{"9713inrPo"},
		{"inrPo9713"},
		{"9713inrPow"},
		{"inrPow9713"},
		{"Zi"},
		{"iZ"},
		{"Zin"},
		{"inZ"},
		{"Zinr"},
		{"inrZ"},
		{"ZinrP"},
		{"inrPZ"},
		{"ZinrPo"},
		{"inrPoZ"},
		{"ZinrPow"},
		{"inrPowZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.InrPow()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestInrMinValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"inrMin", 6},
		{"inrMin ", 6},
		{"inrMin\n", 6},
		{"inrMin.", 6},
		{"inrMin:", 6},
		{"inrMin,", 6},
		{"inrMin\"", 6},
		{"inrMin(", 6},
		{"inrMin)", 6},
		{"inrMin[", 6},
		{"inrMin]", 6},
		{"inrMin// comment", 6},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.InrMin()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestInrMinInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_i"},
		{"i_"},
		{"_in"},
		{"in_"},
		{"_inr"},
		{"inr_"},
		{"_inrM"},
		{"inrM_"},
		{"_inrMi"},
		{"inrMi_"},
		{"_inrMin"},
		{"inrMin_"},
		{"9713i"},
		{"i9713"},
		{"9713in"},
		{"in9713"},
		{"9713inr"},
		{"inr9713"},
		{"9713inrM"},
		{"inrM9713"},
		{"9713inrMi"},
		{"inrMi9713"},
		{"9713inrMin"},
		{"inrMin9713"},
		{"Zi"},
		{"iZ"},
		{"Zin"},
		{"inZ"},
		{"Zinr"},
		{"inrZ"},
		{"ZinrM"},
		{"inrMZ"},
		{"ZinrMi"},
		{"inrMiZ"},
		{"ZinrMin"},
		{"inrMinZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.InrMin()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestInrMaxValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"inrMax", 6},
		{"inrMax ", 6},
		{"inrMax\n", 6},
		{"inrMax.", 6},
		{"inrMax:", 6},
		{"inrMax,", 6},
		{"inrMax\"", 6},
		{"inrMax(", 6},
		{"inrMax)", 6},
		{"inrMax[", 6},
		{"inrMax]", 6},
		{"inrMax// comment", 6},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.InrMax()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestInrMaxInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_i"},
		{"i_"},
		{"_in"},
		{"in_"},
		{"_inr"},
		{"inr_"},
		{"_inrM"},
		{"inrM_"},
		{"_inrMa"},
		{"inrMa_"},
		{"_inrMax"},
		{"inrMax_"},
		{"9713i"},
		{"i9713"},
		{"9713in"},
		{"in9713"},
		{"9713inr"},
		{"inr9713"},
		{"9713inrM"},
		{"inrM9713"},
		{"9713inrMa"},
		{"inrMa9713"},
		{"9713inrMax"},
		{"inrMax9713"},
		{"Zi"},
		{"iZ"},
		{"Zin"},
		{"inZ"},
		{"Zinr"},
		{"inrZ"},
		{"ZinrM"},
		{"inrMZ"},
		{"ZinrMa"},
		{"inrMaZ"},
		{"ZinrMax"},
		{"inrMaxZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.InrMax()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestSumValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"sum", 3},
		{"sum ", 3},
		{"sum\n", 3},
		{"sum.", 3},
		{"sum:", 3},
		{"sum,", 3},
		{"sum\"", 3},
		{"sum(", 3},
		{"sum)", 3},
		{"sum[", 3},
		{"sum]", 3},
		{"sum// comment", 3},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Sum()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestSumInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_s"},
		{"s_"},
		{"_su"},
		{"su_"},
		{"_sum"},
		{"sum_"},
		{"9713s"},
		{"s9713"},
		{"9713su"},
		{"su9713"},
		{"9713sum"},
		{"sum9713"},
		{"Zs"},
		{"sZ"},
		{"Zsu"},
		{"suZ"},
		{"Zsum"},
		{"sumZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Sum()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestPrdValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"prd", 3},
		{"prd ", 3},
		{"prd\n", 3},
		{"prd.", 3},
		{"prd:", 3},
		{"prd,", 3},
		{"prd\"", 3},
		{"prd(", 3},
		{"prd)", 3},
		{"prd[", 3},
		{"prd]", 3},
		{"prd// comment", 3},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Prd()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestPrdInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_p"},
		{"p_"},
		{"_pr"},
		{"pr_"},
		{"_prd"},
		{"prd_"},
		{"9713p"},
		{"p9713"},
		{"9713pr"},
		{"pr9713"},
		{"9713prd"},
		{"prd9713"},
		{"Zp"},
		{"pZ"},
		{"Zpr"},
		{"prZ"},
		{"Zprd"},
		{"prdZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Prd()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestMdnValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"mdn", 3},
		{"mdn ", 3},
		{"mdn\n", 3},
		{"mdn.", 3},
		{"mdn:", 3},
		{"mdn,", 3},
		{"mdn\"", 3},
		{"mdn(", 3},
		{"mdn)", 3},
		{"mdn[", 3},
		{"mdn]", 3},
		{"mdn// comment", 3},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Mdn()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestMdnInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_m"},
		{"m_"},
		{"_md"},
		{"md_"},
		{"_mdn"},
		{"mdn_"},
		{"9713m"},
		{"m9713"},
		{"9713md"},
		{"md9713"},
		{"9713mdn"},
		{"mdn9713"},
		{"Zm"},
		{"mZ"},
		{"Zmd"},
		{"mdZ"},
		{"Zmdn"},
		{"mdnZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Mdn()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestSmaValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"sma", 3},
		{"sma ", 3},
		{"sma\n", 3},
		{"sma.", 3},
		{"sma:", 3},
		{"sma,", 3},
		{"sma\"", 3},
		{"sma(", 3},
		{"sma)", 3},
		{"sma[", 3},
		{"sma]", 3},
		{"sma// comment", 3},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Sma()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestSmaInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_s"},
		{"s_"},
		{"_sm"},
		{"sm_"},
		{"_sma"},
		{"sma_"},
		{"9713s"},
		{"s9713"},
		{"9713sm"},
		{"sm9713"},
		{"9713sma"},
		{"sma9713"},
		{"Zs"},
		{"sZ"},
		{"Zsm"},
		{"smZ"},
		{"Zsma"},
		{"smaZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Sma()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestGmaValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"gma", 3},
		{"gma ", 3},
		{"gma\n", 3},
		{"gma.", 3},
		{"gma:", 3},
		{"gma,", 3},
		{"gma\"", 3},
		{"gma(", 3},
		{"gma)", 3},
		{"gma[", 3},
		{"gma]", 3},
		{"gma// comment", 3},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Gma()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestGmaInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_g"},
		{"g_"},
		{"_gm"},
		{"gm_"},
		{"_gma"},
		{"gma_"},
		{"9713g"},
		{"g9713"},
		{"9713gm"},
		{"gm9713"},
		{"9713gma"},
		{"gma9713"},
		{"Zg"},
		{"gZ"},
		{"Zgm"},
		{"gmZ"},
		{"Zgma"},
		{"gmaZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Gma()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestWmaValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"wma", 3},
		{"wma ", 3},
		{"wma\n", 3},
		{"wma.", 3},
		{"wma:", 3},
		{"wma,", 3},
		{"wma\"", 3},
		{"wma(", 3},
		{"wma)", 3},
		{"wma[", 3},
		{"wma]", 3},
		{"wma// comment", 3},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Wma()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestWmaInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_w"},
		{"w_"},
		{"_wm"},
		{"wm_"},
		{"_wma"},
		{"wma_"},
		{"9713w"},
		{"w9713"},
		{"9713wm"},
		{"wm9713"},
		{"9713wma"},
		{"wma9713"},
		{"Zw"},
		{"wZ"},
		{"Zwm"},
		{"wmZ"},
		{"Zwma"},
		{"wmaZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Wma()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestVrncValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"vrnc", 4},
		{"vrnc ", 4},
		{"vrnc\n", 4},
		{"vrnc.", 4},
		{"vrnc:", 4},
		{"vrnc,", 4},
		{"vrnc\"", 4},
		{"vrnc(", 4},
		{"vrnc)", 4},
		{"vrnc[", 4},
		{"vrnc]", 4},
		{"vrnc// comment", 4},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Vrnc()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestVrncInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_v"},
		{"v_"},
		{"_vr"},
		{"vr_"},
		{"_vrn"},
		{"vrn_"},
		{"_vrnc"},
		{"vrnc_"},
		{"9713v"},
		{"v9713"},
		{"9713vr"},
		{"vr9713"},
		{"9713vrn"},
		{"vrn9713"},
		{"9713vrnc"},
		{"vrnc9713"},
		{"Zv"},
		{"vZ"},
		{"Zvr"},
		{"vrZ"},
		{"Zvrn"},
		{"vrnZ"},
		{"Zvrnc"},
		{"vrncZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Vrnc()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestStdValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"std", 3},
		{"std ", 3},
		{"std\n", 3},
		{"std.", 3},
		{"std:", 3},
		{"std,", 3},
		{"std\"", 3},
		{"std(", 3},
		{"std)", 3},
		{"std[", 3},
		{"std]", 3},
		{"std// comment", 3},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Std()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestStdInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_s"},
		{"s_"},
		{"_st"},
		{"st_"},
		{"_std"},
		{"std_"},
		{"9713s"},
		{"s9713"},
		{"9713st"},
		{"st9713"},
		{"9713std"},
		{"std9713"},
		{"Zs"},
		{"sZ"},
		{"Zst"},
		{"stZ"},
		{"Zstd"},
		{"stdZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Std()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestZscrValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"zscr", 4},
		{"zscr ", 4},
		{"zscr\n", 4},
		{"zscr.", 4},
		{"zscr:", 4},
		{"zscr,", 4},
		{"zscr\"", 4},
		{"zscr(", 4},
		{"zscr)", 4},
		{"zscr[", 4},
		{"zscr]", 4},
		{"zscr// comment", 4},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Zscr()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestZscrInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_z"},
		{"z_"},
		{"_zs"},
		{"zs_"},
		{"_zsc"},
		{"zsc_"},
		{"_zscr"},
		{"zscr_"},
		{"9713z"},
		{"z9713"},
		{"9713zs"},
		{"zs9713"},
		{"9713zsc"},
		{"zsc9713"},
		{"9713zscr"},
		{"zscr9713"},
		{"Zz"},
		{"zZ"},
		{"Zzs"},
		{"zsZ"},
		{"Zzsc"},
		{"zscZ"},
		{"Zzscr"},
		{"zscrZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Zscr()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestZscrInplaceValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"zscrInplace", 11},
		{"zscrInplace ", 11},
		{"zscrInplace\n", 11},
		{"zscrInplace.", 11},
		{"zscrInplace:", 11},
		{"zscrInplace,", 11},
		{"zscrInplace\"", 11},
		{"zscrInplace(", 11},
		{"zscrInplace)", 11},
		{"zscrInplace[", 11},
		{"zscrInplace]", 11},
		{"zscrInplace// comment", 11},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.ZscrInplace()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestZscrInplaceInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_z"},
		{"z_"},
		{"_zs"},
		{"zs_"},
		{"_zsc"},
		{"zsc_"},
		{"_zscr"},
		{"zscr_"},
		{"_zscrI"},
		{"zscrI_"},
		{"_zscrIn"},
		{"zscrIn_"},
		{"_zscrInp"},
		{"zscrInp_"},
		{"_zscrInpl"},
		{"zscrInpl_"},
		{"_zscrInpla"},
		{"zscrInpla_"},
		{"_zscrInplac"},
		{"zscrInplac_"},
		{"_zscrInplace"},
		{"zscrInplace_"},
		{"9713z"},
		{"z9713"},
		{"9713zs"},
		{"zs9713"},
		{"9713zsc"},
		{"zsc9713"},
		{"9713zscr"},
		{"zscr9713"},
		{"9713zscrI"},
		{"zscrI9713"},
		{"9713zscrIn"},
		{"zscrIn9713"},
		{"9713zscrInp"},
		{"zscrInp9713"},
		{"9713zscrInpl"},
		{"zscrInpl9713"},
		{"9713zscrInpla"},
		{"zscrInpla9713"},
		{"9713zscrInplac"},
		{"zscrInplac9713"},
		{"9713zscrInplace"},
		{"zscrInplace9713"},
		{"Zz"},
		{"zZ"},
		{"Zzs"},
		{"zsZ"},
		{"Zzsc"},
		{"zscZ"},
		{"Zzscr"},
		{"zscrZ"},
		{"ZzscrI"},
		{"zscrIZ"},
		{"ZzscrIn"},
		{"zscrInZ"},
		{"ZzscrInp"},
		{"zscrInpZ"},
		{"ZzscrInpl"},
		{"zscrInplZ"},
		{"ZzscrInpla"},
		{"zscrInplaZ"},
		{"ZzscrInplac"},
		{"zscrInplacZ"},
		{"ZzscrInplace"},
		{"zscrInplaceZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.ZscrInplace()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestRngFulValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"rngFul", 6},
		{"rngFul ", 6},
		{"rngFul\n", 6},
		{"rngFul.", 6},
		{"rngFul:", 6},
		{"rngFul,", 6},
		{"rngFul\"", 6},
		{"rngFul(", 6},
		{"rngFul)", 6},
		{"rngFul[", 6},
		{"rngFul]", 6},
		{"rngFul// comment", 6},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.RngFul()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestRngFulInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_r"},
		{"r_"},
		{"_rn"},
		{"rn_"},
		{"_rng"},
		{"rng_"},
		{"_rngF"},
		{"rngF_"},
		{"_rngFu"},
		{"rngFu_"},
		{"_rngFul"},
		{"rngFul_"},
		{"9713r"},
		{"r9713"},
		{"9713rn"},
		{"rn9713"},
		{"9713rng"},
		{"rng9713"},
		{"9713rngF"},
		{"rngF9713"},
		{"9713rngFu"},
		{"rngFu9713"},
		{"9713rngFul"},
		{"rngFul9713"},
		{"Zr"},
		{"rZ"},
		{"Zrn"},
		{"rnZ"},
		{"Zrng"},
		{"rngZ"},
		{"ZrngF"},
		{"rngFZ"},
		{"ZrngFu"},
		{"rngFuZ"},
		{"ZrngFul"},
		{"rngFulZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.RngFul()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestRngLstValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"rngLst", 6},
		{"rngLst ", 6},
		{"rngLst\n", 6},
		{"rngLst.", 6},
		{"rngLst:", 6},
		{"rngLst,", 6},
		{"rngLst\"", 6},
		{"rngLst(", 6},
		{"rngLst)", 6},
		{"rngLst[", 6},
		{"rngLst]", 6},
		{"rngLst// comment", 6},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.RngLst()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestRngLstInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_r"},
		{"r_"},
		{"_rn"},
		{"rn_"},
		{"_rng"},
		{"rng_"},
		{"_rngL"},
		{"rngL_"},
		{"_rngLs"},
		{"rngLs_"},
		{"_rngLst"},
		{"rngLst_"},
		{"9713r"},
		{"r9713"},
		{"9713rn"},
		{"rn9713"},
		{"9713rng"},
		{"rng9713"},
		{"9713rngL"},
		{"rngL9713"},
		{"9713rngLs"},
		{"rngLs9713"},
		{"9713rngLst"},
		{"rngLst9713"},
		{"Zr"},
		{"rZ"},
		{"Zrn"},
		{"rnZ"},
		{"Zrng"},
		{"rngZ"},
		{"ZrngL"},
		{"rngLZ"},
		{"ZrngLs"},
		{"rngLsZ"},
		{"ZrngLst"},
		{"rngLstZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.RngLst()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestProLstValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"proLst", 6},
		{"proLst ", 6},
		{"proLst\n", 6},
		{"proLst.", 6},
		{"proLst:", 6},
		{"proLst,", 6},
		{"proLst\"", 6},
		{"proLst(", 6},
		{"proLst)", 6},
		{"proLst[", 6},
		{"proLst]", 6},
		{"proLst// comment", 6},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.ProLst()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestProLstInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_p"},
		{"p_"},
		{"_pr"},
		{"pr_"},
		{"_pro"},
		{"pro_"},
		{"_proL"},
		{"proL_"},
		{"_proLs"},
		{"proLs_"},
		{"_proLst"},
		{"proLst_"},
		{"9713p"},
		{"p9713"},
		{"9713pr"},
		{"pr9713"},
		{"9713pro"},
		{"pro9713"},
		{"9713proL"},
		{"proL9713"},
		{"9713proLs"},
		{"proLs9713"},
		{"9713proLst"},
		{"proLst9713"},
		{"Zp"},
		{"pZ"},
		{"Zpr"},
		{"prZ"},
		{"Zpro"},
		{"proZ"},
		{"ZproL"},
		{"proLZ"},
		{"ZproLs"},
		{"proLsZ"},
		{"ZproLst"},
		{"proLstZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.ProLst()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestProSmaValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"proSma", 6},
		{"proSma ", 6},
		{"proSma\n", 6},
		{"proSma.", 6},
		{"proSma:", 6},
		{"proSma,", 6},
		{"proSma\"", 6},
		{"proSma(", 6},
		{"proSma)", 6},
		{"proSma[", 6},
		{"proSma]", 6},
		{"proSma// comment", 6},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.ProSma()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestProSmaInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_p"},
		{"p_"},
		{"_pr"},
		{"pr_"},
		{"_pro"},
		{"pro_"},
		{"_proS"},
		{"proS_"},
		{"_proSm"},
		{"proSm_"},
		{"_proSma"},
		{"proSma_"},
		{"9713p"},
		{"p9713"},
		{"9713pr"},
		{"pr9713"},
		{"9713pro"},
		{"pro9713"},
		{"9713proS"},
		{"proS9713"},
		{"9713proSm"},
		{"proSm9713"},
		{"9713proSma"},
		{"proSma9713"},
		{"Zp"},
		{"pZ"},
		{"Zpr"},
		{"prZ"},
		{"Zpro"},
		{"proZ"},
		{"ZproS"},
		{"proSZ"},
		{"ZproSm"},
		{"proSmZ"},
		{"ZproSma"},
		{"proSmaZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.ProSma()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestSubSumPosValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"subSumPos", 9},
		{"subSumPos ", 9},
		{"subSumPos\n", 9},
		{"subSumPos.", 9},
		{"subSumPos:", 9},
		{"subSumPos,", 9},
		{"subSumPos\"", 9},
		{"subSumPos(", 9},
		{"subSumPos)", 9},
		{"subSumPos[", 9},
		{"subSumPos]", 9},
		{"subSumPos// comment", 9},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.SubSumPos()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestSubSumPosInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_s"},
		{"s_"},
		{"_su"},
		{"su_"},
		{"_sub"},
		{"sub_"},
		{"_subS"},
		{"subS_"},
		{"_subSu"},
		{"subSu_"},
		{"_subSum"},
		{"subSum_"},
		{"_subSumP"},
		{"subSumP_"},
		{"_subSumPo"},
		{"subSumPo_"},
		{"_subSumPos"},
		{"subSumPos_"},
		{"9713s"},
		{"s9713"},
		{"9713su"},
		{"su9713"},
		{"9713sub"},
		{"sub9713"},
		{"9713subS"},
		{"subS9713"},
		{"9713subSu"},
		{"subSu9713"},
		{"9713subSum"},
		{"subSum9713"},
		{"9713subSumP"},
		{"subSumP9713"},
		{"9713subSumPo"},
		{"subSumPo9713"},
		{"9713subSumPos"},
		{"subSumPos9713"},
		{"Zs"},
		{"sZ"},
		{"Zsu"},
		{"suZ"},
		{"Zsub"},
		{"subZ"},
		{"ZsubS"},
		{"subSZ"},
		{"ZsubSu"},
		{"subSuZ"},
		{"ZsubSum"},
		{"subSumZ"},
		{"ZsubSumP"},
		{"subSumPZ"},
		{"ZsubSumPo"},
		{"subSumPoZ"},
		{"ZsubSumPos"},
		{"subSumPosZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.SubSumPos()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestSubSumNegValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"subSumNeg", 9},
		{"subSumNeg ", 9},
		{"subSumNeg\n", 9},
		{"subSumNeg.", 9},
		{"subSumNeg:", 9},
		{"subSumNeg,", 9},
		{"subSumNeg\"", 9},
		{"subSumNeg(", 9},
		{"subSumNeg)", 9},
		{"subSumNeg[", 9},
		{"subSumNeg]", 9},
		{"subSumNeg// comment", 9},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.SubSumNeg()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestSubSumNegInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_s"},
		{"s_"},
		{"_su"},
		{"su_"},
		{"_sub"},
		{"sub_"},
		{"_subS"},
		{"subS_"},
		{"_subSu"},
		{"subSu_"},
		{"_subSum"},
		{"subSum_"},
		{"_subSumN"},
		{"subSumN_"},
		{"_subSumNe"},
		{"subSumNe_"},
		{"_subSumNeg"},
		{"subSumNeg_"},
		{"9713s"},
		{"s9713"},
		{"9713su"},
		{"su9713"},
		{"9713sub"},
		{"sub9713"},
		{"9713subS"},
		{"subS9713"},
		{"9713subSu"},
		{"subSu9713"},
		{"9713subSum"},
		{"subSum9713"},
		{"9713subSumN"},
		{"subSumN9713"},
		{"9713subSumNe"},
		{"subSumNe9713"},
		{"9713subSumNeg"},
		{"subSumNeg9713"},
		{"Zs"},
		{"sZ"},
		{"Zsu"},
		{"suZ"},
		{"Zsub"},
		{"subZ"},
		{"ZsubS"},
		{"subSZ"},
		{"ZsubSu"},
		{"subSuZ"},
		{"ZsubSum"},
		{"subSumZ"},
		{"ZsubSumN"},
		{"subSumNZ"},
		{"ZsubSumNe"},
		{"subSumNeZ"},
		{"ZsubSumNeg"},
		{"subSumNegZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.SubSumNeg()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestRsiValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"rsi", 3},
		{"rsi ", 3},
		{"rsi\n", 3},
		{"rsi.", 3},
		{"rsi:", 3},
		{"rsi,", 3},
		{"rsi\"", 3},
		{"rsi(", 3},
		{"rsi)", 3},
		{"rsi[", 3},
		{"rsi]", 3},
		{"rsi// comment", 3},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Rsi()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestRsiInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_r"},
		{"r_"},
		{"_rs"},
		{"rs_"},
		{"_rsi"},
		{"rsi_"},
		{"9713r"},
		{"r9713"},
		{"9713rs"},
		{"rs9713"},
		{"9713rsi"},
		{"rsi9713"},
		{"Zr"},
		{"rZ"},
		{"Zrs"},
		{"rsZ"},
		{"Zrsi"},
		{"rsiZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Rsi()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestWrsiValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"wrsi", 4},
		{"wrsi ", 4},
		{"wrsi\n", 4},
		{"wrsi.", 4},
		{"wrsi:", 4},
		{"wrsi,", 4},
		{"wrsi\"", 4},
		{"wrsi(", 4},
		{"wrsi)", 4},
		{"wrsi[", 4},
		{"wrsi]", 4},
		{"wrsi// comment", 4},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Wrsi()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestWrsiInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_w"},
		{"w_"},
		{"_wr"},
		{"wr_"},
		{"_wrs"},
		{"wrs_"},
		{"_wrsi"},
		{"wrsi_"},
		{"9713w"},
		{"w9713"},
		{"9713wr"},
		{"wr9713"},
		{"9713wrs"},
		{"wrs9713"},
		{"9713wrsi"},
		{"wrsi9713"},
		{"Zw"},
		{"wZ"},
		{"Zwr"},
		{"wrZ"},
		{"Zwrs"},
		{"wrsZ"},
		{"Zwrsi"},
		{"wrsiZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Wrsi()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestProValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"pro", 3},
		{"pro ", 3},
		{"pro\n", 3},
		{"pro.", 3},
		{"pro:", 3},
		{"pro,", 3},
		{"pro\"", 3},
		{"pro(", 3},
		{"pro)", 3},
		{"pro[", 3},
		{"pro]", 3},
		{"pro// comment", 3},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.Pro()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestProInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_p"},
		{"p_"},
		{"_pr"},
		{"pr_"},
		{"_pro"},
		{"pro_"},
		{"9713p"},
		{"p9713"},
		{"9713pr"},
		{"pr9713"},
		{"9713pro"},
		{"pro9713"},
		{"Zp"},
		{"pZ"},
		{"Zpr"},
		{"prZ"},
		{"Zpro"},
		{"proZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.Pro()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestAlmaValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"alma", 4},
		{"alma ", 4},
		{"alma.", 4},