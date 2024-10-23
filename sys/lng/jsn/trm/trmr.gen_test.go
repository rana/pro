package trm_test

import (
	"fmt"
	"sys/bsc/bnd"
	"sys/bsc/unt"
	"sys/lng/jsn/trm"
	"sys/tst"
	"testing"
)

func TestStrLitValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"\"\"", 2},
		{"\"xYz\"", 5},
		{"\"a\"", 3},
		{"\"efg HIJ jKl\"", 13},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.StrLit()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, trm.StrLit{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestStrLitInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_\""},
		{"\"_"},
		{"_\"\""},
		{"\"\"_"},
		{"9713\""},
		{"\"9713"},
		{"9713\"\""},
		{"\"\"9713"},
		{"Z\""},
		{"\"Z"},
		{"Z\"\""},
		{"\"\"Z"},
		{"_\""},
		{"\"_"},
		{"_\"x"},
		{"\"x_"},
		{"_\"xY"},
		{"\"xY_"},
		{"_\"xYz"},
		{"\"xYz_"},
		{"_\"xYz\""},
		{"\"xYz\"_"},
		{"9713\""},
		{"\"9713"},
		{"9713\"x"},
		{"\"x9713"},
		{"9713\"xY"},
		{"\"xY9713"},
		{"9713\"xYz"},
		{"\"xYz9713"},
		{"9713\"xYz\""},
		{"\"xYz\"9713"},
		{"Z\""},
		{"\"Z"},
		{"Z\"x"},
		{"\"xZ"},
		{"Z\"xY"},
		{"\"xYZ"},
		{"Z\"xYz"},
		{"\"xYzZ"},
		{"Z\"xYz\""},
		{"\"xYz\"Z"},
		{"_\""},
		{"\"_"},
		{"_\"a"},
		{"\"a_"},
		{"_\"a\""},
		{"\"a\"_"},
		{"9713\""},
		{"\"9713"},
		{"9713\"a"},
		{"\"a9713"},
		{"9713\"a\""},
		{"\"a\"9713"},
		{"Z\""},
		{"\"Z"},
		{"Z\"a"},
		{"\"aZ"},
		{"Z\"a\""},
		{"\"a\"Z"},
		{"_\""},
		{"\"_"},
		{"_\"e"},
		{"\"e_"},
		{"_\"ef"},
		{"\"ef_"},
		{"_\"efg"},
		{"\"efg_"},
		{"_\"efg "},
		{"\"efg _"},
		{"_\"efg H"},
		{"\"efg H_"},
		{"_\"efg HI"},
		{"\"efg HI_"},
		{"_\"efg HIJ"},
		{"\"efg HIJ_"},
		{"_\"efg HIJ "},
		{"\"efg HIJ _"},
		{"_\"efg HIJ j"},
		{"\"efg HIJ j_"},
		{"_\"efg HIJ jK"},
		{"\"efg HIJ jK_"},
		{"_\"efg HIJ jKl"},
		{"\"efg HIJ jKl_"},
		{"_\"efg HIJ jKl\""},
		{"\"efg HIJ jKl\"_"},
		{"9713\""},
		{"\"9713"},
		{"9713\"e"},
		{"\"e9713"},
		{"9713\"ef"},
		{"\"ef9713"},
		{"9713\"efg"},
		{"\"efg9713"},
		{"9713\"efg "},
		{"\"efg 9713"},
		{"9713\"efg H"},
		{"\"efg H9713"},
		{"9713\"efg HI"},
		{"\"efg HI9713"},
		{"9713\"efg HIJ"},
		{"\"efg HIJ9713"},
		{"9713\"efg HIJ "},
		{"\"efg HIJ 9713"},
		{"9713\"efg HIJ j"},
		{"\"efg HIJ j9713"},
		{"9713\"efg HIJ jK"},
		{"\"efg HIJ jK9713"},
		{"9713\"efg HIJ jKl"},
		{"\"efg HIJ jKl9713"},
		{"9713\"efg HIJ jKl\""},
		{"\"efg HIJ jKl\"9713"},
		{"Z\""},
		{"\"Z"},
		{"Z\"e"},
		{"\"eZ"},
		{"Z\"ef"},
		{"\"efZ"},
		{"Z\"efg"},
		{"\"efgZ"},
		{"Z\"efg "},
		{"\"efg Z"},
		{"Z\"efg H"},
		{"\"efg HZ"},
		{"Z\"efg HI"},
		{"\"efg HIZ"},
		{"Z\"efg HIJ"},
		{"\"efg HIJZ"},
		{"Z\"efg HIJ "},
		{"\"efg HIJ Z"},
		{"Z\"efg HIJ j"},
		{"\"efg HIJ jZ"},
		{"Z\"efg HIJ jK"},
		{"\"efg HIJ jKZ"},
		{"Z\"efg HIJ jKl"},
		{"\"efg HIJ jKlZ"},
		{"Z\"efg HIJ jKl\""},
		{"\"efg HIJ jKl\"Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.StrLit()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestFalseLitValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"false", 5},
		{"false ", 5},
		{"false\n", 5},
		{"false.", 5},
		{"false:", 5},
		{"false,", 5},
		{"false\"", 5},
		{"false(", 5},
		{"false)", 5},
		{"false[", 5},
		{"false]", 5},
		{"false// comment", 5},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.FalseLit()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestFalseLitInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_f"},
		{"f_"},
		{"_fa"},
		{"fa_"},
		{"_fal"},
		{"fal_"},
		{"_fals"},
		{"fals_"},
		{"_false"},
		{"false_"},
		{"9713f"},
		{"f9713"},
		{"9713fa"},
		{"fa9713"},
		{"9713fal"},
		{"fal9713"},
		{"9713fals"},
		{"fals9713"},
		{"9713false"},
		{"false9713"},
		{"Zf"},
		{"fZ"},
		{"Zfa"},
		{"faZ"},
		{"Zfal"},
		{"falZ"},
		{"Zfals"},
		{"falsZ"},
		{"Zfalse"},
		{"falseZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.FalseLit()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestTrueLitValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"true", 4},
		{"true ", 4},
		{"true\n", 4},
		{"true.", 4},
		{"true:", 4},
		{"true,", 4},
		{"true\"", 4},
		{"true(", 4},
		{"true)", 4},
		{"true[", 4},
		{"true]", 4},
		{"true// comment", 4},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.TrueLit()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, bnd.Bnd{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestTrueLitInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_t"},
		{"t_"},
		{"_tr"},
		{"tr_"},
		{"_tru"},
		{"tru_"},
		{"_true"},
		{"true_"},
		{"9713t"},
		{"t9713"},
		{"9713tr"},
		{"tr9713"},
		{"9713tru"},
		{"tru9713"},
		{"9713true"},
		{"true9713"},
		{"Zt"},
		{"tZ"},
		{"Ztr"},
		{"trZ"},
		{"Ztru"},
		{"truZ"},
		{"Ztrue"},
		{"trueZ"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.TrueLit()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestBolLitValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"false", 5},
		{"true", 4},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.BolLit()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, trm.BolLit{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestBolLitInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.BolLit()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestFltLitValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"0.0", 3},
		{"1.1", 3},
		{"3.0", 3},
		{"3.0", 3},
		{"3.0", 3},
		{"99999.99", 8},
		{"-1.1", 4},
		{"-99999.99", 9},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.FltLit()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, trm.FltLit{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestFltLitInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_0"},
		{"0_"},
		{"_0."},
		{"0._"},
		{"_0.0"},
		{"0.0_"},
		{"Z0"},
		{"0Z"},
		{"Z0."},
		{"0.Z"},
		{"Z0.0"},
		{"0.0Z"},
		{"_1"},
		{"1_"},
		{"_1."},
		{"1._"},
		{"_1.1"},
		{"1.1_"},
		{"Z1"},
		{"1Z"},
		{"Z1."},
		{"1.Z"},
		{"Z1.1"},
		{"1.1Z"},
		{"_3"},
		{"3_"},
		{"_3."},
		{"3._"},
		{"_3.0"},
		{"3.0_"},
		{"Z3"},
		{"3Z"},
		{"Z3."},
		{"3.Z"},
		{"Z3.0"},
		{"3.0Z"},
		{"_3"},
		{"3_"},
		{"_3."},
		{"3._"},
		{"_3.0"},
		{"3.0_"},
		{"Z3"},
		{"3Z"},
		{"Z3."},
		{"3.Z"},
		{"Z3.0"},
		{"3.0Z"},
		{"_3"},
		{"3_"},
		{"_3."},
		{"3._"},
		{"_3.0"},
		{"3.0_"},
		{"Z3"},
		{"3Z"},
		{"Z3."},
		{"3.Z"},
		{"Z3.0"},
		{"3.0Z"},
		{"_9"},
		{"9_"},
		{"_99"},
		{"99_"},
		{"_999"},
		{"999_"},
		{"_9999"},
		{"9999_"},
		{"_99999"},
		{"99999_"},
		{"_99999."},
		{"99999._"},
		{"_99999.9"},
		{"99999.9_"},
		{"_99999.99"},
		{"99999.99_"},
		{"Z9"},
		{"9Z"},
		{"Z99"},
		{"99Z"},
		{"Z999"},
		{"999Z"},
		{"Z9999"},
		{"9999Z"},
		{"Z99999"},
		{"99999Z"},
		{"Z99999."},
		{"99999.Z"},
		{"Z99999.9"},
		{"99999.9Z"},
		{"Z99999.99"},
		{"99999.99Z"},
		{"_-"},
		{"-_"},
		{"_-1"},
		{"-1_"},
		{"_-1."},
		{"-1._"},
		{"_-1.1"},
		{"-1.1_"},
		{"Z-"},
		{"-Z"},
		{"Z-1"},
		{"-1Z"},
		{"Z-1."},
		{"-1.Z"},
		{"Z-1.1"},
		{"-1.1Z"},
		{"_-"},
		{"-_"},
		{"_-9"},
		{"-9_"},
		{"_-99"},
		{"-99_"},
		{"_-999"},
		{"-999_"},
		{"_-9999"},
		{"-9999_"},
		{"_-99999"},
		{"-99999_"},
		{"_-99999."},
		{"-99999._"},
		{"_-99999.9"},
		{"-99999.9_"},
		{"_-99999.99"},
		{"-99999.99_"},
		{"Z-"},
		{"-Z"},
		{"Z-9"},
		{"-9Z"},
		{"Z-99"},
		{"-99Z"},
		{"Z-999"},
		{"-999Z"},
		{"Z-9999"},
		{"-9999Z"},
		{"Z-99999"},
		{"-99999Z"},
		{"Z-99999."},
		{"-99999.Z"},
		{"Z-99999.9"},
		{"-99999.9Z"},
		{"Z-99999.99"},
		{"-99999.99Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.FltLit()
			tst.False(t, ok, "Lex")
		})
	}
}
func TestIntLitValid(t *testing.T) {
	cses := []struct {
		txt string
		lim unt.Unt
	}{
		{"0", 1},
		{"10", 2},
		{"1000", 4},
		{"-10", 3},
		{"-1000", 5},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		if len(lbl) > 16 {
			lbl = lbl[:16]
		}
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			a, ok := trmr.IntLit()
			tst.True(t, ok, "Lex")
			tst.TypeEql(t, trm.IntLit{}, a)
			tst.UntEql(t, 0, a.Idx, "Idx")
			tst.UntEql(t, cse.lim, a.Lim, "Lim")
		})
	}
}
func TestIntLitInvalid(t *testing.T) {
	cses := []struct {
		txt string
	}{
		{"_0"},
		{"0_"},
		{"Z0"},
		{"0Z"},
		{"_1"},
		{"1_"},
		{"_10"},
		{"10_"},
		{"Z1"},
		{"1Z"},
		{"Z10"},
		{"10Z"},
		{"_1"},
		{"1_"},
		{"_10"},
		{"10_"},
		{"_100"},
		{"100_"},
		{"_1000"},
		{"1000_"},
		{"Z1"},
		{"1Z"},
		{"Z10"},
		{"10Z"},
		{"Z100"},
		{"100Z"},
		{"Z1000"},
		{"1000Z"},
		{"_-"},
		{"-_"},
		{"_-1"},
		{"-1_"},
		{"_-10"},
		{"-10_"},
		{"Z-"},
		{"-Z"},
		{"Z-1"},
		{"-1Z"},
		{"Z-10"},
		{"-10Z"},
		{"_-"},
		{"-_"},
		{"_-1"},
		{"-1_"},
		{"_-10"},
		{"-10_"},
		{"_-100"},
		{"-100_"},
		{"_-1000"},
		{"-1000_"},
		{"Z-"},
		{"-Z"},
		{"Z-1"},
		{"-1Z"},
		{"Z-10"},
		{"-10Z"},
		{"Z-100"},
		{"-100Z"},
		{"Z-1000"},
		{"-1000Z"},
	}
	var trmr trm.Trmr
	for _, cse := range cses {
		lbl := cse.txt
		t.Run(fmt.Sprintf("%q", lbl), func(t *testing.T) {
			trmr.Reset(cse.txt)
			_, ok := trmr.IntLit()
			tst.False(t, ok, "Lex")
		})
	}
}
