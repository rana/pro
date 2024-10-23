package cfg_test

import (
	"sys/bsc/bnd"
	"sys/bsc/bnds"
	"sys/bsc/bol"
	"sys/bsc/bols"
	"sys/bsc/flt"
	"sys/bsc/flts"
	"sys/bsc/int"
	"sys/bsc/ints"
	"sys/bsc/str"
	"sys/bsc/strs"
	"sys/bsc/tme"
	"sys/bsc/tmes"
	"sys/bsc/unt"
	"sys/bsc/unts"
	"sys/lng/pro/cfg"
	"sys/tst"
	"testing"
)

func TestCfgStrValid(t *testing.T) {
	cses := []struct {
		e   str.Str
		pth []string
		txt string
	}{
		{"", []string{"key"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n \"\"  \t\r\n // comment\n }"},
		{"", []string{"key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n \"\"  \t\r\n // comment\n  no3:\"no\" }"},
		{"", []string{"k2", "k1", "key"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n \"\"  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{"", []string{"k2", "k1", "key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n \"\"  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{"xYz", []string{"key"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n \"xYz\"  \t\r\n // comment\n }"},
		{"xYz", []string{"key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n \"xYz\"  \t\r\n // comment\n  no3:\"no\" }"},
		{"xYz", []string{"k2", "k1", "key"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n \"xYz\"  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{"xYz", []string{"k2", "k1", "key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n \"xYz\"  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{"a", []string{"key"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n \"a\"  \t\r\n // comment\n }"},
		{"a", []string{"key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n \"a\"  \t\r\n // comment\n  no3:\"no\" }"},
		{"a", []string{"k2", "k1", "key"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n \"a\"  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{"a", []string{"k2", "k1", "key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n \"a\"  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{"efg HIJ jKl", []string{"key"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n \"efg HIJ jKl\"  \t\r\n // comment\n }"},
		{"efg HIJ jKl", []string{"key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n \"efg HIJ jKl\"  \t\r\n // comment\n  no3:\"no\" }"},
		{"efg HIJ jKl", []string{"k2", "k1", "key"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n \"efg HIJ jKl\"  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{"efg HIJ jKl", []string{"k2", "k1", "key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n \"efg HIJ jKl\"  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
	}
	for _, cse := range cses {
		t.Run("Str", func(t *testing.T) {
			var c cfg.Cfgr
			c.Reset(cse.txt)
			a := c.Str(cse.pth...)
			tst.StrEql(t, cse.e, a)
		})
	}
}
func TestCfgStrInvalid(t *testing.T) {
	cses := []struct {
		e   str.Str
		pth []string
		txt string
	}{
		{"", []string{}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n \"\"  \t\r\n // comment\n }"},
		{"", []string{"wrong"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n \"\"  \t\r\n // comment\n }"},
		{"", []string{"wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n \"\"  \t\r\n // comment\n  no3:\"no\" }"},
		{"", []string{"k2", "k1", "wrong"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n \"\"  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{"", []string{"k2", "k1", "wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n \"\"  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{"xYz", []string{}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n \"xYz\"  \t\r\n // comment\n }"},
		{"xYz", []string{"wrong"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n \"xYz\"  \t\r\n // comment\n }"},
		{"xYz", []string{"wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n \"xYz\"  \t\r\n // comment\n  no3:\"no\" }"},
		{"xYz", []string{"k2", "k1", "wrong"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n \"xYz\"  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{"xYz", []string{"k2", "k1", "wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n \"xYz\"  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{"a", []string{}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n \"a\"  \t\r\n // comment\n }"},
		{"a", []string{"wrong"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n \"a\"  \t\r\n // comment\n }"},
		{"a", []string{"wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n \"a\"  \t\r\n // comment\n  no3:\"no\" }"},
		{"a", []string{"k2", "k1", "wrong"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n \"a\"  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{"a", []string{"k2", "k1", "wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n \"a\"  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{"efg HIJ jKl", []string{}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n \"efg HIJ jKl\"  \t\r\n // comment\n }"},
		{"efg HIJ jKl", []string{"wrong"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n \"efg HIJ jKl\"  \t\r\n // comment\n }"},
		{"efg HIJ jKl", []string{"wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n \"efg HIJ jKl\"  \t\r\n // comment\n  no3:\"no\" }"},
		{"efg HIJ jKl", []string{"k2", "k1", "wrong"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n \"efg HIJ jKl\"  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{"efg HIJ jKl", []string{"k2", "k1", "wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n \"efg HIJ jKl\"  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
	}
	var c cfg.Cfgr
	for _, cse := range cses {
		t.Run("Str", func(t *testing.T) {
			c.Reset(cse.txt)
			tst.Panic(t, func() { c.Str(cse.pth...) })
		})
	}
}
func TestCfgBolValid(t *testing.T) {
	cses := []struct {
		e   bol.Bol
		pth []string
		txt string
	}{
		{false, []string{"key"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n fls  \t\r\n // comment\n }"},
		{false, []string{"key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n fls  \t\r\n // comment\n  no3:\"no\" }"},
		{false, []string{"k2", "k1", "key"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n fls  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{false, []string{"k2", "k1", "key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n fls  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{true, []string{"key"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n tru  \t\r\n // comment\n }"},
		{true, []string{"key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n tru  \t\r\n // comment\n  no3:\"no\" }"},
		{true, []string{"k2", "k1", "key"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n tru  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{true, []string{"k2", "k1", "key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n tru  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
	}
	for _, cse := range cses {
		t.Run("Bol", func(t *testing.T) {
			var c cfg.Cfgr
			c.Reset(cse.txt)
			a := c.Bol(cse.pth...)
			tst.BolEql(t, cse.e, a)
		})
	}
}
func TestCfgBolInvalid(t *testing.T) {
	cses := []struct {
		e   bol.Bol
		pth []string
		txt string
	}{
		{false, []string{}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n fls  \t\r\n // comment\n }"},
		{false, []string{"wrong"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n fls  \t\r\n // comment\n }"},
		{false, []string{"wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n fls  \t\r\n // comment\n  no3:\"no\" }"},
		{false, []string{"k2", "k1", "wrong"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n fls  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{false, []string{"k2", "k1", "wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n fls  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{true, []string{}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n tru  \t\r\n // comment\n }"},
		{true, []string{"wrong"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n tru  \t\r\n // comment\n }"},
		{true, []string{"wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n tru  \t\r\n // comment\n  no3:\"no\" }"},
		{true, []string{"k2", "k1", "wrong"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n tru  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{true, []string{"k2", "k1", "wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n tru  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
	}
	var c cfg.Cfgr
	for _, cse := range cses {
		t.Run("Bol", func(t *testing.T) {
			c.Reset(cse.txt)
			tst.Panic(t, func() { c.Bol(cse.pth...) })
		})
	}
}
func TestCfgFltValid(t *testing.T) {
	cses := []struct {
		e   flt.Flt
		pth []string
		txt string
	}{
		{0.0, []string{"key"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 0.0  \t\r\n // comment\n }"},
		{0.0, []string{"key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 0.0  \t\r\n // comment\n  no3:\"no\" }"},
		{0.0, []string{"k2", "k1", "key"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 0.0  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{0.0, []string{"k2", "k1", "key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 0.0  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{1.1, []string{"key"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 1.1  \t\r\n // comment\n }"},
		{1.1, []string{"key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 1.1  \t\r\n // comment\n  no3:\"no\" }"},
		{1.1, []string{"k2", "k1", "key"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 1.1  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{1.1, []string{"k2", "k1", "key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 1.1  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{3.0, []string{"key"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 3.0  \t\r\n // comment\n }"},
		{3.0, []string{"key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 3.0  \t\r\n // comment\n  no3:\"no\" }"},
		{3.0, []string{"k2", "k1", "key"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 3.0  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{3.0, []string{"k2", "k1", "key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 3.0  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{3.0, []string{"key"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 3.0  \t\r\n // comment\n }"},
		{3.0, []string{"key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 3.0  \t\r\n // comment\n  no3:\"no\" }"},
		{3.0, []string{"k2", "k1", "key"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 3.0  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{3.0, []string{"k2", "k1", "key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 3.0  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{3.0, []string{"key"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 3.0  \t\r\n // comment\n }"},
		{3.0, []string{"key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 3.0  \t\r\n // comment\n  no3:\"no\" }"},
		{3.0, []string{"k2", "k1", "key"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 3.0  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{3.0, []string{"k2", "k1", "key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 3.0  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{99999.99, []string{"key"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 99999.99  \t\r\n // comment\n }"},
		{99999.99, []string{"key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 99999.99  \t\r\n // comment\n  no3:\"no\" }"},
		{99999.99, []string{"k2", "k1", "key"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 99999.99  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{99999.99, []string{"k2", "k1", "key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 99999.99  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{-1.1, []string{"key"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n -1.1  \t\r\n // comment\n }"},
		{-1.1, []string{"key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n -1.1  \t\r\n // comment\n  no3:\"no\" }"},
		{-1.1, []string{"k2", "k1", "key"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n -1.1  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{-1.1, []string{"k2", "k1", "key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n -1.1  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{-99999.99, []string{"key"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n -99999.99  \t\r\n // comment\n }"},
		{-99999.99, []string{"key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n -99999.99  \t\r\n // comment\n  no3:\"no\" }"},
		{-99999.99, []string{"k2", "k1", "key"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n -99999.99  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{-99999.99, []string{"k2", "k1", "key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n -99999.99  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
	}
	for _, cse := range cses {
		t.Run("Flt", func(t *testing.T) {
			var c cfg.Cfgr
			c.Reset(cse.txt)
			a := c.Flt(cse.pth...)
			tst.FltEql(t, cse.e, a)
		})
	}
}
func TestCfgFltInvalid(t *testing.T) {
	cses := []struct {
		e   flt.Flt
		pth []string
		txt string
	}{
		{0.0, []string{}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 0.0  \t\r\n // comment\n }"},
		{0.0, []string{"wrong"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 0.0  \t\r\n // comment\n }"},
		{0.0, []string{"wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 0.0  \t\r\n // comment\n  no3:\"no\" }"},
		{0.0, []string{"k2", "k1", "wrong"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 0.0  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{0.0, []string{"k2", "k1", "wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 0.0  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{1.1, []string{}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 1.1  \t\r\n // comment\n }"},
		{1.1, []string{"wrong"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 1.1  \t\r\n // comment\n }"},
		{1.1, []string{"wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 1.1  \t\r\n // comment\n  no3:\"no\" }"},
		{1.1, []string{"k2", "k1", "wrong"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 1.1  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{1.1, []string{"k2", "k1", "wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 1.1  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{3.0, []string{}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 3.0  \t\r\n // comment\n }"},
		{3.0, []string{"wrong"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 3.0  \t\r\n // comment\n }"},
		{3.0, []string{"wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 3.0  \t\r\n // comment\n  no3:\"no\" }"},
		{3.0, []string{"k2", "k1", "wrong"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 3.0  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{3.0, []string{"k2", "k1", "wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 3.0  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{3.0, []string{}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 3.0  \t\r\n // comment\n }"},
		{3.0, []string{"wrong"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 3.0  \t\r\n // comment\n }"},
		{3.0, []string{"wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 3.0  \t\r\n // comment\n  no3:\"no\" }"},
		{3.0, []string{"k2", "k1", "wrong"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 3.0  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{3.0, []string{"k2", "k1", "wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 3.0  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{3.0, []string{}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 3.0  \t\r\n // comment\n }"},
		{3.0, []string{"wrong"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 3.0  \t\r\n // comment\n }"},
		{3.0, []string{"wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 3.0  \t\r\n // comment\n  no3:\"no\" }"},
		{3.0, []string{"k2", "k1", "wrong"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 3.0  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{3.0, []string{"k2", "k1", "wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 3.0  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{99999.99, []string{}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 99999.99  \t\r\n // comment\n }"},
		{99999.99, []string{"wrong"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 99999.99  \t\r\n // comment\n }"},
		{99999.99, []string{"wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 99999.99  \t\r\n // comment\n  no3:\"no\" }"},
		{99999.99, []string{"k2", "k1", "wrong"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 99999.99  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{99999.99, []string{"k2", "k1", "wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 99999.99  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{-1.1, []string{}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n -1.1  \t\r\n // comment\n }"},
		{-1.1, []string{"wrong"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n -1.1  \t\r\n // comment\n }"},
		{-1.1, []string{"wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n -1.1  \t\r\n // comment\n  no3:\"no\" }"},
		{-1.1, []string{"k2", "k1", "wrong"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n -1.1  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{-1.1, []string{"k2", "k1", "wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n -1.1  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{-99999.99, []string{}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n -99999.99  \t\r\n // comment\n }"},
		{-99999.99, []string{"wrong"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n -99999.99  \t\r\n // comment\n }"},
		{-99999.99, []string{"wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n -99999.99  \t\r\n // comment\n  no3:\"no\" }"},
		{-99999.99, []string{"k2", "k1", "wrong"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n -99999.99  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{-99999.99, []string{"k2", "k1", "wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n -99999.99  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
	}
	var c cfg.Cfgr
	for _, cse := range cses {
		t.Run("Flt", func(t *testing.T) {
			c.Reset(cse.txt)
			tst.Panic(t, func() { c.Flt(cse.pth...) })
		})
	}
}
func TestCfgUntValid(t *testing.T) {
	cses := []struct {
		e   unt.Unt
		pth []string
		txt string
	}{
		{0, []string{"key"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 0  \t\r\n // comment\n }"},
		{0, []string{"key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 0  \t\r\n // comment\n  no3:\"no\" }"},
		{0, []string{"k2", "k1", "key"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 0  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{0, []string{"k2", "k1", "key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 0  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{1, []string{"key"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 1  \t\r\n // comment\n }"},
		{1, []string{"key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 1  \t\r\n // comment\n  no3:\"no\" }"},
		{1, []string{"k2", "k1", "key"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 1  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{1, []string{"k2", "k1", "key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 1  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{1000, []string{"key"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 1000  \t\r\n // comment\n }"},
		{1000, []string{"key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 1000  \t\r\n // comment\n  no3:\"no\" }"},
		{1000, []string{"k2", "k1", "key"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 1000  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{1000, []string{"k2", "k1", "key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 1000  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{10, []string{"key"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 10  \t\r\n // comment\n }"},
		{10, []string{"key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 10  \t\r\n // comment\n  no3:\"no\" }"},
		{10, []string{"k2", "k1", "key"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 10  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{10, []string{"k2", "k1", "key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 10  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
	}
	for _, cse := range cses {
		t.Run("Unt", func(t *testing.T) {
			var c cfg.Cfgr
			c.Reset(cse.txt)
			a := c.Unt(cse.pth...)
			tst.UntEql(t, cse.e, a)
		})
	}
}
func TestCfgUntInvalid(t *testing.T) {
	cses := []struct {
		e   unt.Unt
		pth []string
		txt string
	}{
		{0, []string{}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 0  \t\r\n // comment\n }"},
		{0, []string{"wrong"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 0  \t\r\n // comment\n }"},
		{0, []string{"wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 0  \t\r\n // comment\n  no3:\"no\" }"},
		{0, []string{"k2", "k1", "wrong"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 0  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{0, []string{"k2", "k1", "wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 0  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{1, []string{}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 1  \t\r\n // comment\n }"},
		{1, []string{"wrong"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 1  \t\r\n // comment\n }"},
		{1, []string{"wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 1  \t\r\n // comment\n  no3:\"no\" }"},
		{1, []string{"k2", "k1", "wrong"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 1  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{1, []string{"k2", "k1", "wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 1  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{1000, []string{}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 1000  \t\r\n // comment\n }"},
		{1000, []string{"wrong"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 1000  \t\r\n // comment\n }"},
		{1000, []string{"wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 1000  \t\r\n // comment\n  no3:\"no\" }"},
		{1000, []string{"k2", "k1", "wrong"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 1000  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{1000, []string{"k2", "k1", "wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 1000  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{10, []string{}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 10  \t\r\n // comment\n }"},
		{10, []string{"wrong"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 10  \t\r\n // comment\n }"},
		{10, []string{"wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 10  \t\r\n // comment\n  no3:\"no\" }"},
		{10, []string{"k2", "k1", "wrong"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 10  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{10, []string{"k2", "k1", "wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 10  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
	}
	var c cfg.Cfgr
	for _, cse := range cses {
		t.Run("Unt", func(t *testing.T) {
			c.Reset(cse.txt)
			tst.Panic(t, func() { c.Unt(cse.pth...) })
		})
	}
}
func TestCfgIntValid(t *testing.T) {
	cses := []struct {
		e   int.Int
		pth []string
		txt string
	}{
		{0, []string{"key"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n +0  \t\r\n // comment\n }"},
		{0, []string{"key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n +0  \t\r\n // comment\n  no3:\"no\" }"},
		{0, []string{"k2", "k1", "key"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n +0  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{0, []string{"k2", "k1", "key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n +0  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{10, []string{"key"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n +10  \t\r\n // comment\n }"},
		{10, []string{"key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n +10  \t\r\n // comment\n  no3:\"no\" }"},
		{10, []string{"k2", "k1", "key"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n +10  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{10, []string{"k2", "k1", "key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n +10  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{1000, []string{"key"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n +1000  \t\r\n // comment\n }"},
		{1000, []string{"key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n +1000  \t\r\n // comment\n  no3:\"no\" }"},
		{1000, []string{"k2", "k1", "key"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n +1000  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{1000, []string{"k2", "k1", "key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n +1000  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{-10, []string{"key"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n -10  \t\r\n // comment\n }"},
		{-10, []string{"key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n -10  \t\r\n // comment\n  no3:\"no\" }"},
		{-10, []string{"k2", "k1", "key"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n -10  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{-10, []string{"k2", "k1", "key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n -10  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{-1000, []string{"key"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n -1000  \t\r\n // comment\n }"},
		{-1000, []string{"key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n -1000  \t\r\n // comment\n  no3:\"no\" }"},
		{-1000, []string{"k2", "k1", "key"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n -1000  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{-1000, []string{"k2", "k1", "key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n -1000  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
	}
	for _, cse := range cses {
		t.Run("Int", func(t *testing.T) {
			var c cfg.Cfgr
			c.Reset(cse.txt)
			a := c.Int(cse.pth...)
			tst.IntEql(t, cse.e, a)
		})
	}
}
func TestCfgIntInvalid(t *testing.T) {
	cses := []struct {
		e   int.Int
		pth []string
		txt string
	}{
		{0, []string{}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n +0  \t\r\n // comment\n }"},
		{0, []string{"wrong"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n +0  \t\r\n // comment\n }"},
		{0, []string{"wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n +0  \t\r\n // comment\n  no3:\"no\" }"},
		{0, []string{"k2", "k1", "wrong"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n +0  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{0, []string{"k2", "k1", "wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n +0  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{10, []string{}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n +10  \t\r\n // comment\n }"},
		{10, []string{"wrong"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n +10  \t\r\n // comment\n }"},
		{10, []string{"wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n +10  \t\r\n // comment\n  no3:\"no\" }"},
		{10, []string{"k2", "k1", "wrong"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n +10  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{10, []string{"k2", "k1", "wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n +10  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{1000, []string{}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n +1000  \t\r\n // comment\n }"},
		{1000, []string{"wrong"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n +1000  \t\r\n // comment\n }"},
		{1000, []string{"wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n +1000  \t\r\n // comment\n  no3:\"no\" }"},
		{1000, []string{"k2", "k1", "wrong"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n +1000  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{1000, []string{"k2", "k1", "wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n +1000  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{-10, []string{}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n -10  \t\r\n // comment\n }"},
		{-10, []string{"wrong"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n -10  \t\r\n // comment\n }"},
		{-10, []string{"wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n -10  \t\r\n // comment\n  no3:\"no\" }"},
		{-10, []string{"k2", "k1", "wrong"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n -10  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{-10, []string{"k2", "k1", "wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n -10  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{-1000, []string{}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n -1000  \t\r\n // comment\n }"},
		{-1000, []string{"wrong"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n -1000  \t\r\n // comment\n }"},
		{-1000, []string{"wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n -1000  \t\r\n // comment\n  no3:\"no\" }"},
		{-1000, []string{"k2", "k1", "wrong"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n -1000  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{-1000, []string{"k2", "k1", "wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n -1000  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
	}
	var c cfg.Cfgr
	for _, cse := range cses {
		t.Run("Int", func(t *testing.T) {
			c.Reset(cse.txt)
			tst.Panic(t, func() { c.Int(cse.pth...) })
		})
	}
}
func TestCfgTmeValid(t *testing.T) {
	cses := []struct {
		e   tme.Tme
		pth []string
		txt string
	}{
		{0, []string{"key"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 0s  \t\r\n // comment\n }"},
		{0, []string{"key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 0s  \t\r\n // comment\n  no3:\"no\" }"},
		{0, []string{"k2", "k1", "key"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 0s  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{0, []string{"k2", "k1", "key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 0s  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{1, []string{"key"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 1s  \t\r\n // comment\n }"},
		{1, []string{"key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 1s  \t\r\n // comment\n  no3:\"no\" }"},
		{1, []string{"k2", "k1", "key"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 1s  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{1, []string{"k2", "k1", "key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 1s  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{10, []string{"key"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 10s  \t\r\n // comment\n }"},
		{10, []string{"key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 10s  \t\r\n // comment\n  no3:\"no\" }"},
		{10, []string{"k2", "k1", "key"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 10s  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{10, []string{"k2", "k1", "key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 10s  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{-1, []string{"key"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n -1s  \t\r\n // comment\n }"},
		{-1, []string{"key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n -1s  \t\r\n // comment\n  no3:\"no\" }"},
		{-1, []string{"k2", "k1", "key"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n -1s  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{-1, []string{"k2", "k1", "key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n -1s  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{-10 * 60, []string{"key"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n -10m  \t\r\n // comment\n }"},
		{-10 * 60, []string{"key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n -10m  \t\r\n // comment\n  no3:\"no\" }"},
		{-10 * 60, []string{"k2", "k1", "key"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n -10m  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{-10 * 60, []string{"k2", "k1", "key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n -10m  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{788645, []string{"key"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 1w2d3h4m5s  \t\r\n // comment\n }"},
		{788645, []string{"key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 1w2d3h4m5s  \t\r\n // comment\n  no3:\"no\" }"},
		{788645, []string{"k2", "k1", "key"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 1w2d3h4m5s  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{788645, []string{"k2", "k1", "key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 1w2d3h4m5s  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{-788645, []string{"key"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n -1w2d3h4m5s  \t\r\n // comment\n }"},
		{-788645, []string{"key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n -1w2d3h4m5s  \t\r\n // comment\n  no3:\"no\" }"},
		{-788645, []string{"k2", "k1", "key"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n -1w2d3h4m5s  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{-788645, []string{"k2", "k1", "key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n -1w2d3h4m5s  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{946782245, []string{"key"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 2000y1n2d3h4m5s  \t\r\n // comment\n }"},
		{946782245, []string{"key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 2000y1n2d3h4m5s  \t\r\n // comment\n  no3:\"no\" }"},
		{946782245, []string{"k2", "k1", "key"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 2000y1n2d3h4m5s  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{946782245, []string{"k2", "k1", "key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 2000y1n2d3h4m5s  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{946782245, []string{"key"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 2000y2d3h4m5s  \t\r\n // comment\n }"},
		{946782245, []string{"key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 2000y2d3h4m5s  \t\r\n // comment\n  no3:\"no\" }"},
		{946782245, []string{"k2", "k1", "key"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 2000y2d3h4m5s  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{946782245, []string{"k2", "k1", "key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 2000y2d3h4m5s  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{946695845, []string{"key"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 2000y3h4m5s  \t\r\n // comment\n }"},
		{946695845, []string{"key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 2000y3h4m5s  \t\r\n // comment\n  no3:\"no\" }"},
		{946695845, []string{"k2", "k1", "key"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 2000y3h4m5s  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{946695845, []string{"k2", "k1", "key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 2000y3h4m5s  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
	}
	for _, cse := range cses {
		t.Run("Tme", func(t *testing.T) {
			var c cfg.Cfgr
			c.Reset(cse.txt)
			a := c.Tme(cse.pth...)
			tst.TmeEql(t, cse.e, a)
		})
	}
}
func TestCfgTmeInvalid(t *testing.T) {
	cses := []struct {
		e   tme.Tme
		pth []string
		txt string
	}{
		{0, []string{}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 0s  \t\r\n // comment\n }"},
		{0, []string{"wrong"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 0s  \t\r\n // comment\n }"},
		{0, []string{"wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 0s  \t\r\n // comment\n  no3:\"no\" }"},
		{0, []string{"k2", "k1", "wrong"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 0s  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{0, []string{"k2", "k1", "wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 0s  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{1, []string{}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 1s  \t\r\n // comment\n }"},
		{1, []string{"wrong"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 1s  \t\r\n // comment\n }"},
		{1, []string{"wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 1s  \t\r\n // comment\n  no3:\"no\" }"},
		{1, []string{"k2", "k1", "wrong"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 1s  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{1, []string{"k2", "k1", "wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 1s  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{10, []string{}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 10s  \t\r\n // comment\n }"},
		{10, []string{"wrong"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 10s  \t\r\n // comment\n }"},
		{10, []string{"wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 10s  \t\r\n // comment\n  no3:\"no\" }"},
		{10, []string{"k2", "k1", "wrong"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 10s  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{10, []string{"k2", "k1", "wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 10s  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{-1, []string{}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n -1s  \t\r\n // comment\n }"},
		{-1, []string{"wrong"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n -1s  \t\r\n // comment\n }"},
		{-1, []string{"wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n -1s  \t\r\n // comment\n  no3:\"no\" }"},
		{-1, []string{"k2", "k1", "wrong"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n -1s  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{-1, []string{"k2", "k1", "wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n -1s  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{-10 * 60, []string{}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n -10m  \t\r\n // comment\n }"},
		{-10 * 60, []string{"wrong"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n -10m  \t\r\n // comment\n }"},
		{-10 * 60, []string{"wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n -10m  \t\r\n // comment\n  no3:\"no\" }"},
		{-10 * 60, []string{"k2", "k1", "wrong"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n -10m  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{-10 * 60, []string{"k2", "k1", "wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n -10m  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{788645, []string{}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 1w2d3h4m5s  \t\r\n // comment\n }"},
		{788645, []string{"wrong"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 1w2d3h4m5s  \t\r\n // comment\n }"},
		{788645, []string{"wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 1w2d3h4m5s  \t\r\n // comment\n  no3:\"no\" }"},
		{788645, []string{"k2", "k1", "wrong"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 1w2d3h4m5s  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{788645, []string{"k2", "k1", "wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 1w2d3h4m5s  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{-788645, []string{}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n -1w2d3h4m5s  \t\r\n // comment\n }"},
		{-788645, []string{"wrong"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n -1w2d3h4m5s  \t\r\n // comment\n }"},
		{-788645, []string{"wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n -1w2d3h4m5s  \t\r\n // comment\n  no3:\"no\" }"},
		{-788645, []string{"k2", "k1", "wrong"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n -1w2d3h4m5s  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{-788645, []string{"k2", "k1", "wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n -1w2d3h4m5s  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{946782245, []string{}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 2000y1n2d3h4m5s  \t\r\n // comment\n }"},
		{946782245, []string{"wrong"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 2000y1n2d3h4m5s  \t\r\n // comment\n }"},
		{946782245, []string{"wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 2000y1n2d3h4m5s  \t\r\n // comment\n  no3:\"no\" }"},
		{946782245, []string{"k2", "k1", "wrong"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 2000y1n2d3h4m5s  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{946782245, []string{"k2", "k1", "wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 2000y1n2d3h4m5s  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{946782245, []string{}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 2000y2d3h4m5s  \t\r\n // comment\n }"},
		{946782245, []string{"wrong"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 2000y2d3h4m5s  \t\r\n // comment\n }"},
		{946782245, []string{"wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 2000y2d3h4m5s  \t\r\n // comment\n  no3:\"no\" }"},
		{946782245, []string{"k2", "k1", "wrong"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 2000y2d3h4m5s  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{946782245, []string{"k2", "k1", "wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 2000y2d3h4m5s  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{946695845, []string{}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 2000y3h4m5s  \t\r\n // comment\n }"},
		{946695845, []string{"wrong"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 2000y3h4m5s  \t\r\n // comment\n }"},
		{946695845, []string{"wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 2000y3h4m5s  \t\r\n // comment\n  no3:\"no\" }"},
		{946695845, []string{"k2", "k1", "wrong"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 2000y3h4m5s  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{946695845, []string{"k2", "k1", "wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 2000y3h4m5s  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
	}
	var c cfg.Cfgr
	for _, cse := range cses {
		t.Run("Tme", func(t *testing.T) {
			c.Reset(cse.txt)
			tst.Panic(t, func() { c.Tme(cse.pth...) })
		})
	}
}
func TestCfgBndValid(t *testing.T) {
	cses := []struct {
		e   bnd.Bnd
		pth []string
		txt string
	}{
		{bnd.Bnd{Idx: 0, Lim: 0}, []string{"key"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 0-0  \t\r\n // comment\n }"},
		{bnd.Bnd{Idx: 0, Lim: 0}, []string{"key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 0-0  \t\r\n // comment\n  no3:\"no\" }"},
		{bnd.Bnd{Idx: 0, Lim: 0}, []string{"k2", "k1", "key"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 0-0  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{bnd.Bnd{Idx: 0, Lim: 0}, []string{"k2", "k1", "key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 0-0  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{bnd.Bnd{Idx: 0, Lim: 1}, []string{"key"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 0-1  \t\r\n // comment\n }"},
		{bnd.Bnd{Idx: 0, Lim: 1}, []string{"key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 0-1  \t\r\n // comment\n  no3:\"no\" }"},
		{bnd.Bnd{Idx: 0, Lim: 1}, []string{"k2", "k1", "key"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 0-1  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{bnd.Bnd{Idx: 0, Lim: 1}, []string{"k2", "k1", "key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 0-1  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{bnd.Bnd{Idx: 0, Lim: 1000}, []string{"key"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 0-1000  \t\r\n // comment\n }"},
		{bnd.Bnd{Idx: 0, Lim: 1000}, []string{"key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 0-1000  \t\r\n // comment\n  no3:\"no\" }"},
		{bnd.Bnd{Idx: 0, Lim: 1000}, []string{"k2", "k1", "key"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 0-1000  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{bnd.Bnd{Idx: 0, Lim: 1000}, []string{"k2", "k1", "key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 0-1000  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{bnd.Bnd{Idx: 999, Lim: 1000}, []string{"key"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 999-1000  \t\r\n // comment\n }"},
		{bnd.Bnd{Idx: 999, Lim: 1000}, []string{"key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 999-1000  \t\r\n // comment\n  no3:\"no\" }"},
		{bnd.Bnd{Idx: 999, Lim: 1000}, []string{"k2", "k1", "key"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 999-1000  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{bnd.Bnd{Idx: 999, Lim: 1000}, []string{"k2", "k1", "key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 999-1000  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{bnd.Bnd{Idx: 1, Lim: 0}, []string{"key"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 1-0  \t\r\n // comment\n }"},
		{bnd.Bnd{Idx: 1, Lim: 0}, []string{"key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 1-0  \t\r\n // comment\n  no3:\"no\" }"},
		{bnd.Bnd{Idx: 1, Lim: 0}, []string{"k2", "k1", "key"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 1-0  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{bnd.Bnd{Idx: 1, Lim: 0}, []string{"k2", "k1", "key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 1-0  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
	}
	for _, cse := range cses {
		t.Run("Bnd", func(t *testing.T) {
			var c cfg.Cfgr
			c.Reset(cse.txt)
			a := c.Bnd(cse.pth...)
			tst.BndEql(t, cse.e, a)
		})
	}
}
func TestCfgBndInvalid(t *testing.T) {
	cses := []struct {
		e   bnd.Bnd
		pth []string
		txt string
	}{
		{bnd.Bnd{Idx: 0, Lim: 0}, []string{}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 0-0  \t\r\n // comment\n }"},
		{bnd.Bnd{Idx: 0, Lim: 0}, []string{"wrong"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 0-0  \t\r\n // comment\n }"},
		{bnd.Bnd{Idx: 0, Lim: 0}, []string{"wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 0-0  \t\r\n // comment\n  no3:\"no\" }"},
		{bnd.Bnd{Idx: 0, Lim: 0}, []string{"k2", "k1", "wrong"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 0-0  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{bnd.Bnd{Idx: 0, Lim: 0}, []string{"k2", "k1", "wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 0-0  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{bnd.Bnd{Idx: 0, Lim: 1}, []string{}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 0-1  \t\r\n // comment\n }"},
		{bnd.Bnd{Idx: 0, Lim: 1}, []string{"wrong"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 0-1  \t\r\n // comment\n }"},
		{bnd.Bnd{Idx: 0, Lim: 1}, []string{"wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 0-1  \t\r\n // comment\n  no3:\"no\" }"},
		{bnd.Bnd{Idx: 0, Lim: 1}, []string{"k2", "k1", "wrong"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 0-1  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{bnd.Bnd{Idx: 0, Lim: 1}, []string{"k2", "k1", "wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 0-1  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{bnd.Bnd{Idx: 0, Lim: 1000}, []string{}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 0-1000  \t\r\n // comment\n }"},
		{bnd.Bnd{Idx: 0, Lim: 1000}, []string{"wrong"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 0-1000  \t\r\n // comment\n }"},
		{bnd.Bnd{Idx: 0, Lim: 1000}, []string{"wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 0-1000  \t\r\n // comment\n  no3:\"no\" }"},
		{bnd.Bnd{Idx: 0, Lim: 1000}, []string{"k2", "k1", "wrong"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 0-1000  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{bnd.Bnd{Idx: 0, Lim: 1000}, []string{"k2", "k1", "wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 0-1000  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{bnd.Bnd{Idx: 999, Lim: 1000}, []string{}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 999-1000  \t\r\n // comment\n }"},
		{bnd.Bnd{Idx: 999, Lim: 1000}, []string{"wrong"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 999-1000  \t\r\n // comment\n }"},
		{bnd.Bnd{Idx: 999, Lim: 1000}, []string{"wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 999-1000  \t\r\n // comment\n  no3:\"no\" }"},
		{bnd.Bnd{Idx: 999, Lim: 1000}, []string{"k2", "k1", "wrong"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 999-1000  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{bnd.Bnd{Idx: 999, Lim: 1000}, []string{"k2", "k1", "wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 999-1000  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{bnd.Bnd{Idx: 1, Lim: 0}, []string{}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 1-0  \t\r\n // comment\n }"},
		{bnd.Bnd{Idx: 1, Lim: 0}, []string{"wrong"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 1-0  \t\r\n // comment\n }"},
		{bnd.Bnd{Idx: 1, Lim: 0}, []string{"wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 1-0  \t\r\n // comment\n  no3:\"no\" }"},
		{bnd.Bnd{Idx: 1, Lim: 0}, []string{"k2", "k1", "wrong"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 1-0  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{bnd.Bnd{Idx: 1, Lim: 0}, []string{"k2", "k1", "wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 1-0  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
	}
	var c cfg.Cfgr
	for _, cse := range cses {
		t.Run("Bnd", func(t *testing.T) {
			c.Reset(cse.txt)
			tst.Panic(t, func() { c.Bnd(cse.pth...) })
		})
	}
}
func TestCfgFltRngValid(t *testing.T) {
	cses := []struct {
		e   flt.Rng
		pth []string
		txt string
	}{
		{flt.Rng{Min: 0, Max: 0}, []string{"key"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 0.0-0.0  \t\r\n // comment\n }"},
		{flt.Rng{Min: 0, Max: 0}, []string{"key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 0.0-0.0  \t\r\n // comment\n  no3:\"no\" }"},
		{flt.Rng{Min: 0, Max: 0}, []string{"k2", "k1", "key"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 0.0-0.0  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{flt.Rng{Min: 0, Max: 0}, []string{"k2", "k1", "key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 0.0-0.0  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{flt.Rng{Min: 0, Max: 1}, []string{"key"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 0.0-1.0  \t\r\n // comment\n }"},
		{flt.Rng{Min: 0, Max: 1}, []string{"key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 0.0-1.0  \t\r\n // comment\n  no3:\"no\" }"},
		{flt.Rng{Min: 0, Max: 1}, []string{"k2", "k1", "key"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 0.0-1.0  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{flt.Rng{Min: 0, Max: 1}, []string{"k2", "k1", "key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 0.0-1.0  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{flt.Rng{Min: -3, Max: -4}, []string{"key"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n -3.0--4.0  \t\r\n // comment\n }"},
		{flt.Rng{Min: -3, Max: -4}, []string{"key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n -3.0--4.0  \t\r\n // comment\n  no3:\"no\" }"},
		{flt.Rng{Min: -3, Max: -4}, []string{"k2", "k1", "key"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n -3.0--4.0  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{flt.Rng{Min: -3, Max: -4}, []string{"k2", "k1", "key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n -3.0--4.0  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{flt.Rng{Min: -999, Max: 1000}, []string{"key"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n -999.0-1000.0  \t\r\n // comment\n }"},
		{flt.Rng{Min: -999, Max: 1000}, []string{"key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n -999.0-1000.0  \t\r\n // comment\n  no3:\"no\" }"},
		{flt.Rng{Min: -999, Max: 1000}, []string{"k2", "k1", "key"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n -999.0-1000.0  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{flt.Rng{Min: -999, Max: 1000}, []string{"k2", "k1", "key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n -999.0-1000.0  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{flt.Rng{Min: 1, Max: 0}, []string{"key"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 1.0-0.0  \t\r\n // comment\n }"},
		{flt.Rng{Min: 1, Max: 0}, []string{"key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 1.0-0.0  \t\r\n // comment\n  no3:\"no\" }"},
		{flt.Rng{Min: 1, Max: 0}, []string{"k2", "k1", "key"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 1.0-0.0  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{flt.Rng{Min: 1, Max: 0}, []string{"k2", "k1", "key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 1.0-0.0  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
	}
	for _, cse := range cses {
		t.Run("FltRng", func(t *testing.T) {
			var c cfg.Cfgr
			c.Reset(cse.txt)
			a := c.FltRng(cse.pth...)
			tst.FltRngEql(t, cse.e, a)
		})
	}
}
func TestCfgFltRngInvalid(t *testing.T) {
	cses := []struct {
		e   flt.Rng
		pth []string
		txt string
	}{
		{flt.Rng{Min: 0, Max: 0}, []string{}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 0.0-0.0  \t\r\n // comment\n }"},
		{flt.Rng{Min: 0, Max: 0}, []string{"wrong"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 0.0-0.0  \t\r\n // comment\n }"},
		{flt.Rng{Min: 0, Max: 0}, []string{"wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 0.0-0.0  \t\r\n // comment\n  no3:\"no\" }"},
		{flt.Rng{Min: 0, Max: 0}, []string{"k2", "k1", "wrong"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 0.0-0.0  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{flt.Rng{Min: 0, Max: 0}, []string{"k2", "k1", "wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 0.0-0.0  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{flt.Rng{Min: 0, Max: 1}, []string{}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 0.0-1.0  \t\r\n // comment\n }"},
		{flt.Rng{Min: 0, Max: 1}, []string{"wrong"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 0.0-1.0  \t\r\n // comment\n }"},
		{flt.Rng{Min: 0, Max: 1}, []string{"wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 0.0-1.0  \t\r\n // comment\n  no3:\"no\" }"},
		{flt.Rng{Min: 0, Max: 1}, []string{"k2", "k1", "wrong"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 0.0-1.0  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{flt.Rng{Min: 0, Max: 1}, []string{"k2", "k1", "wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 0.0-1.0  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{flt.Rng{Min: -3, Max: -4}, []string{}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n -3.0--4.0  \t\r\n // comment\n }"},
		{flt.Rng{Min: -3, Max: -4}, []string{"wrong"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n -3.0--4.0  \t\r\n // comment\n }"},
		{flt.Rng{Min: -3, Max: -4}, []string{"wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n -3.0--4.0  \t\r\n // comment\n  no3:\"no\" }"},
		{flt.Rng{Min: -3, Max: -4}, []string{"k2", "k1", "wrong"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n -3.0--4.0  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{flt.Rng{Min: -3, Max: -4}, []string{"k2", "k1", "wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n -3.0--4.0  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{flt.Rng{Min: -999, Max: 1000}, []string{}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n -999.0-1000.0  \t\r\n // comment\n }"},
		{flt.Rng{Min: -999, Max: 1000}, []string{"wrong"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n -999.0-1000.0  \t\r\n // comment\n }"},
		{flt.Rng{Min: -999, Max: 1000}, []string{"wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n -999.0-1000.0  \t\r\n // comment\n  no3:\"no\" }"},
		{flt.Rng{Min: -999, Max: 1000}, []string{"k2", "k1", "wrong"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n -999.0-1000.0  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{flt.Rng{Min: -999, Max: 1000}, []string{"k2", "k1", "wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n -999.0-1000.0  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{flt.Rng{Min: 1, Max: 0}, []string{}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 1.0-0.0  \t\r\n // comment\n }"},
		{flt.Rng{Min: 1, Max: 0}, []string{"wrong"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 1.0-0.0  \t\r\n // comment\n }"},
		{flt.Rng{Min: 1, Max: 0}, []string{"wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 1.0-0.0  \t\r\n // comment\n  no3:\"no\" }"},
		{flt.Rng{Min: 1, Max: 0}, []string{"k2", "k1", "wrong"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 1.0-0.0  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{flt.Rng{Min: 1, Max: 0}, []string{"k2", "k1", "wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 1.0-0.0  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
	}
	var c cfg.Cfgr
	for _, cse := range cses {
		t.Run("FltRng", func(t *testing.T) {
			c.Reset(cse.txt)
			tst.Panic(t, func() { c.FltRng(cse.pth...) })
		})
	}
}
func TestCfgTmeRngValid(t *testing.T) {
	cses := []struct {
		e   tme.Rng
		pth []string
		txt string
	}{
		{tme.Rng{Min: -10, Max: -1}, []string{"key"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n -10s--1s  \t\r\n // comment\n }"},
		{tme.Rng{Min: -10, Max: -1}, []string{"key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n -10s--1s  \t\r\n // comment\n  no3:\"no\" }"},
		{tme.Rng{Min: -10, Max: -1}, []string{"k2", "k1", "key"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n -10s--1s  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{tme.Rng{Min: -10, Max: -1}, []string{"k2", "k1", "key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n -10s--1s  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{tme.Rng{Min: 0, Max: 1}, []string{"key"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 0s-1s  \t\r\n // comment\n }"},
		{tme.Rng{Min: 0, Max: 1}, []string{"key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 0s-1s  \t\r\n // comment\n  no3:\"no\" }"},
		{tme.Rng{Min: 0, Max: 1}, []string{"k2", "k1", "key"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 0s-1s  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{tme.Rng{Min: 0, Max: 1}, []string{"k2", "k1", "key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 0s-1s  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{tme.Rng{Min: 2, Max: 4}, []string{"key"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 2s-4s  \t\r\n // comment\n }"},
		{tme.Rng{Min: 2, Max: 4}, []string{"key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 2s-4s  \t\r\n // comment\n  no3:\"no\" }"},
		{tme.Rng{Min: 2, Max: 4}, []string{"k2", "k1", "key"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 2s-4s  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{tme.Rng{Min: 2, Max: 4}, []string{"k2", "k1", "key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 2s-4s  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{tme.Rng{Min: 6, Max: 10}, []string{"key"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 6s-10s  \t\r\n // comment\n }"},
		{tme.Rng{Min: 6, Max: 10}, []string{"key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 6s-10s  \t\r\n // comment\n  no3:\"no\" }"},
		{tme.Rng{Min: 6, Max: 10}, []string{"k2", "k1", "key"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 6s-10s  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{tme.Rng{Min: 6, Max: 10}, []string{"k2", "k1", "key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 6s-10s  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{tme.Rng{Min: 50, Max: 60}, []string{"key"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 50s-60s  \t\r\n // comment\n }"},
		{tme.Rng{Min: 50, Max: 60}, []string{"key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 50s-60s  \t\r\n // comment\n  no3:\"no\" }"},
		{tme.Rng{Min: 50, Max: 60}, []string{"k2", "k1", "key"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 50s-60s  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{tme.Rng{Min: 50, Max: 60}, []string{"k2", "k1", "key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 50s-60s  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
	}
	for _, cse := range cses {
		t.Run("TmeRng", func(t *testing.T) {
			var c cfg.Cfgr
			c.Reset(cse.txt)
			a := c.TmeRng(cse.pth...)
			tst.TmeRngEql(t, cse.e, a)
		})
	}
}
func TestCfgTmeRngInvalid(t *testing.T) {
	cses := []struct {
		e   tme.Rng
		pth []string
		txt string
	}{
		{tme.Rng{Min: -10, Max: -1}, []string{}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n -10s--1s  \t\r\n // comment\n }"},
		{tme.Rng{Min: -10, Max: -1}, []string{"wrong"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n -10s--1s  \t\r\n // comment\n }"},
		{tme.Rng{Min: -10, Max: -1}, []string{"wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n -10s--1s  \t\r\n // comment\n  no3:\"no\" }"},
		{tme.Rng{Min: -10, Max: -1}, []string{"k2", "k1", "wrong"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n -10s--1s  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{tme.Rng{Min: -10, Max: -1}, []string{"k2", "k1", "wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n -10s--1s  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{tme.Rng{Min: 0, Max: 1}, []string{}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 0s-1s  \t\r\n // comment\n }"},
		{tme.Rng{Min: 0, Max: 1}, []string{"wrong"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 0s-1s  \t\r\n // comment\n }"},
		{tme.Rng{Min: 0, Max: 1}, []string{"wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 0s-1s  \t\r\n // comment\n  no3:\"no\" }"},
		{tme.Rng{Min: 0, Max: 1}, []string{"k2", "k1", "wrong"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 0s-1s  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{tme.Rng{Min: 0, Max: 1}, []string{"k2", "k1", "wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 0s-1s  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{tme.Rng{Min: 2, Max: 4}, []string{}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 2s-4s  \t\r\n // comment\n }"},
		{tme.Rng{Min: 2, Max: 4}, []string{"wrong"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 2s-4s  \t\r\n // comment\n }"},
		{tme.Rng{Min: 2, Max: 4}, []string{"wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 2s-4s  \t\r\n // comment\n  no3:\"no\" }"},
		{tme.Rng{Min: 2, Max: 4}, []string{"k2", "k1", "wrong"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 2s-4s  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{tme.Rng{Min: 2, Max: 4}, []string{"k2", "k1", "wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 2s-4s  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{tme.Rng{Min: 6, Max: 10}, []string{}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 6s-10s  \t\r\n // comment\n }"},
		{tme.Rng{Min: 6, Max: 10}, []string{"wrong"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 6s-10s  \t\r\n // comment\n }"},
		{tme.Rng{Min: 6, Max: 10}, []string{"wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 6s-10s  \t\r\n // comment\n  no3:\"no\" }"},
		{tme.Rng{Min: 6, Max: 10}, []string{"k2", "k1", "wrong"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 6s-10s  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{tme.Rng{Min: 6, Max: 10}, []string{"k2", "k1", "wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 6s-10s  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{tme.Rng{Min: 50, Max: 60}, []string{}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 50s-60s  \t\r\n // comment\n }"},
		{tme.Rng{Min: 50, Max: 60}, []string{"wrong"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 50s-60s  \t\r\n // comment\n }"},
		{tme.Rng{Min: 50, Max: 60}, []string{"wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 50s-60s  \t\r\n // comment\n  no3:\"no\" }"},
		{tme.Rng{Min: 50, Max: 60}, []string{"k2", "k1", "wrong"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 50s-60s  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{tme.Rng{Min: 50, Max: 60}, []string{"k2", "k1", "wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n 50s-60s  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
	}
	var c cfg.Cfgr
	for _, cse := range cses {
		t.Run("TmeRng", func(t *testing.T) {
			c.Reset(cse.txt)
			tst.Panic(t, func() { c.TmeRng(cse.pth...) })
		})
	}
}
func TestCfgStrsValid(t *testing.T) {
	cses := []struct {
		e   *strs.Strs
		pth []string
		txt string
	}{
		{strs.New("", "xYz", "a"), []string{"key"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n [\"\" \"xYz\" \"a\"]  \t\r\n // comment\n }"},
		{strs.New("", "xYz", "a"), []string{"key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n [\"\" \"xYz\" \"a\"]  \t\r\n // comment\n  no3:\"no\" }"},
		{strs.New("", "xYz", "a"), []string{"k2", "k1", "key"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n [\"\" \"xYz\" \"a\"]  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{strs.New("", "xYz", "a"), []string{"k2", "k1", "key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n [\"\" \"xYz\" \"a\"]  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{strs.New("", "xYz", "a", "efg HIJ jKl"), []string{"key"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n [\"\" \"xYz\" \"a\" \"efg HIJ jKl\"]  \t\r\n // comment\n }"},
		{strs.New("", "xYz", "a", "efg HIJ jKl"), []string{"key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n [\"\" \"xYz\" \"a\" \"efg HIJ jKl\"]  \t\r\n // comment\n  no3:\"no\" }"},
		{strs.New("", "xYz", "a", "efg HIJ jKl"), []string{"k2", "k1", "key"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n [\"\" \"xYz\" \"a\" \"efg HIJ jKl\"]  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{strs.New("", "xYz", "a", "efg HIJ jKl"), []string{"k2", "k1", "key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n [\"\" \"xYz\" \"a\" \"efg HIJ jKl\"]  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
	}
	for _, cse := range cses {
		t.Run("Strs", func(t *testing.T) {
			var c cfg.Cfgr
			c.Reset(cse.txt)
			a := c.Strs(cse.pth...)
			tst.StrsEql(t, cse.e, a)
		})
	}
}
func TestCfgStrsInvalid(t *testing.T) {
	cses := []struct {
		e   *strs.Strs
		pth []string
		txt string
	}{
		{strs.New("", "xYz", "a", "efg HIJ jKl"), []string{}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n [\"\" \"xYz\" \"a\" \"efg HIJ jKl\"]  \t\r\n // comment\n }"},
		{strs.New("", "xYz", "a", "efg HIJ jKl"), []string{"wrong"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n [\"\" \"xYz\" \"a\" \"efg HIJ jKl\"]  \t\r\n // comment\n }"},
		{strs.New("", "xYz", "a", "efg HIJ jKl"), []string{"wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n [\"\" \"xYz\" \"a\" \"efg HIJ jKl\"]  \t\r\n // comment\n  no3:\"no\" }"},
		{strs.New("", "xYz", "a", "efg HIJ jKl"), []string{"k2", "k1", "wrong"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n [\"\" \"xYz\" \"a\" \"efg HIJ jKl\"]  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{strs.New("", "xYz", "a", "efg HIJ jKl"), []string{"k2", "k1", "wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n [\"\" \"xYz\" \"a\" \"efg HIJ jKl\"]  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{strs.New(), []string{}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n []  \t\r\n // comment\n }"},
		{strs.New(), []string{"wrong"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n []  \t\r\n // comment\n }"},
		{strs.New(), []string{"wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n []  \t\r\n // comment\n  no3:\"no\" }"},
		{strs.New(), []string{"k2", "k1", "wrong"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n []  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{strs.New(), []string{"k2", "k1", "wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n []  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
	}
	var c cfg.Cfgr
	for _, cse := range cses {
		t.Run("Strs", func(t *testing.T) {
			c.Reset(cse.txt)
			tst.Panic(t, func() { c.Strs(cse.pth...) })
		})
	}
}
func TestCfgBolsValid(t *testing.T) {
	cses := []struct {
		e   *bols.Bols
		pth []string
		txt string
	}{
		{bols.New(false, true), []string{"key"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n [fls tru]  \t\r\n // comment\n }"},
		{bols.New(false, true), []string{"key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n [fls tru]  \t\r\n // comment\n  no3:\"no\" }"},
		{bols.New(false, true), []string{"k2", "k1", "key"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n [fls tru]  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{bols.New(false, true), []string{"k2", "k1", "key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n [fls tru]  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
	}
	for _, cse := range cses {
		t.Run("Bols", func(t *testing.T) {
			var c cfg.Cfgr
			c.Reset(cse.txt)
			a := c.Bols(cse.pth...)
			tst.BolsEql(t, cse.e, a)
		})
	}
}
func TestCfgBolsInvalid(t *testing.T) {
	cses := []struct {
		e   *bols.Bols
		pth []string
		txt string
	}{
		{bols.New(), []string{}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n []  \t\r\n // comment\n }"},
		{bols.New(), []string{"wrong"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n []  \t\r\n // comment\n }"},
		{bols.New(), []string{"wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n []  \t\r\n // comment\n  no3:\"no\" }"},
		{bols.New(), []string{"k2", "k1", "wrong"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n []  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{bols.New(), []string{"k2", "k1", "wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n []  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
	}
	var c cfg.Cfgr
	for _, cse := range cses {
		t.Run("Bols", func(t *testing.T) {
			c.Reset(cse.txt)
			tst.Panic(t, func() { c.Bols(cse.pth...) })
		})
	}
}
func TestCfgFltsValid(t *testing.T) {
	cses := []struct {
		e   *flts.Flts
		pth []string
		txt string
	}{
		{flts.New(0.0, 1.1, 3.0), []string{"key"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n [0.0 1.1 3.0]  \t\r\n // comment\n }"},
		{flts.New(0.0, 1.1, 3.0), []string{"key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n [0.0 1.1 3.0]  \t\r\n // comment\n  no3:\"no\" }"},
		{flts.New(0.0, 1.1, 3.0), []string{"k2", "k1", "key"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n [0.0 1.1 3.0]  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{flts.New(0.0, 1.1, 3.0), []string{"k2", "k1", "key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n [0.0 1.1 3.0]  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{flts.New(0.0, 1.1, 3.0, 3.0, 3.0, 99999.99, -1.1, -99999.99), []string{"key"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n [0.0 1.1 3.0 3.0 3.0 99999.99 -1.1 -99999.99]  \t\r\n // comment\n }"},
		{flts.New(0.0, 1.1, 3.0, 3.0, 3.0, 99999.99, -1.1, -99999.99), []string{"key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n [0.0 1.1 3.0 3.0 3.0 99999.99 -1.1 -99999.99]  \t\r\n // comment\n  no3:\"no\" }"},
		{flts.New(0.0, 1.1, 3.0, 3.0, 3.0, 99999.99, -1.1, -99999.99), []string{"k2", "k1", "key"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n [0.0 1.1 3.0 3.0 3.0 99999.99 -1.1 -99999.99]  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{flts.New(0.0, 1.1, 3.0, 3.0, 3.0, 99999.99, -1.1, -99999.99), []string{"k2", "k1", "key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n [0.0 1.1 3.0 3.0 3.0 99999.99 -1.1 -99999.99]  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{flts.New(0.0, 1.1, 3.0, 3.0, 3.0, 99999.99, -1.1, -99999.99, 0.0, 1.1, 3.0, 3.0, 3.0, 99999.99, -1.1, -99999.99, 0.0, 1.1, 3.0, 3.0, 3.0, 99999.99, -1.1, -99999.99, 0.0, 1.1, 3.0, 3.0, 3.0, 99999.99, -1.1, -99999.99, 0.0, 1.1, 3.0, 3.0, 3.0, 99999.99, -1.1, -99999.99, 0.0, 1.1, 3.0, 3.0, 3.0, 99999.99, -1.1, -99999.99, 0.0, 1.1, 3.0, 3.0, 3.0, 99999.99, -1.1, -99999.99, 0.0, 1.1, 3.0, 3.0, 3.0, 99999.99, -1.1, -99999.99), []string{"key"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n [0.0 1.1 3.0 3.0 3.0 99999.99 -1.1 -99999.99 0.0 1.1 3.0 3.0 3.0 99999.99 -1.1 -99999.99 0.0 1.1 3.0 3.0 3.0 99999.99 -1.1 -99999.99 0.0 1.1 3.0 3.0 3.0 99999.99 -1.1 -99999.99 0.0 1.1 3.0 3.0 3.0 99999.99 -1.1 -99999.99 0.0 1.1 3.0 3.0 3.0 99999.99 -1.1 -99999.99 0.0 1.1 3.0 3.0 3.0 99999.99 -1.1 -99999.99 0.0 1.1 3.0 3.0 3.0 99999.99 -1.1 -99999.99]  \t\r\n // comment\n }"},
		{flts.New(0.0, 1.1, 3.0, 3.0, 3.0, 99999.99, -1.1, -99999.99, 0.0, 1.1, 3.0, 3.0, 3.0, 99999.99, -1.1, -99999.99, 0.0, 1.1, 3.0, 3.0, 3.0, 99999.99, -1.1, -99999.99, 0.0, 1.1, 3.0, 3.0, 3.0, 99999.99, -1.1, -99999.99, 0.0, 1.1, 3.0, 3.0, 3.0, 99999.99, -1.1, -99999.99, 0.0, 1.1, 3.0, 3.0, 3.0, 99999.99, -1.1, -99999.99, 0.0, 1.1, 3.0, 3.0, 3.0, 99999.99, -1.1, -99999.99, 0.0, 1.1, 3.0, 3.0, 3.0, 99999.99, -1.1, -99999.99), []string{"key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n [0.0 1.1 3.0 3.0 3.0 99999.99 -1.1 -99999.99 0.0 1.1 3.0 3.0 3.0 99999.99 -1.1 -99999.99 0.0 1.1 3.0 3.0 3.0 99999.99 -1.1 -99999.99 0.0 1.1 3.0 3.0 3.0 99999.99 -1.1 -99999.99 0.0 1.1 3.0 3.0 3.0 99999.99 -1.1 -99999.99 0.0 1.1 3.0 3.0 3.0 99999.99 -1.1 -99999.99 0.0 1.1 3.0 3.0 3.0 99999.99 -1.1 -99999.99 0.0 1.1 3.0 3.0 3.0 99999.99 -1.1 -99999.99]  \t\r\n // comment\n  no3:\"no\" }"},
		{flts.New(0.0, 1.1, 3.0, 3.0, 3.0, 99999.99, -1.1, -99999.99, 0.0, 1.1, 3.0, 3.0, 3.0, 99999.99, -1.1, -99999.99, 0.0, 1.1, 3.0, 3.0, 3.0, 99999.99, -1.1, -99999.99, 0.0, 1.1, 3.0, 3.0, 3.0, 99999.99, -1.1, -99999.99, 0.0, 1.1, 3.0, 3.0, 3.0, 99999.99, -1.1, -99999.99, 0.0, 1.1, 3.0, 3.0, 3.0, 99999.99, -1.1, -99999.99, 0.0, 1.1, 3.0, 3.0, 3.0, 99999.99, -1.1, -99999.99, 0.0, 1.1, 3.0, 3.0, 3.0, 99999.99, -1.1, -99999.99), []string{"k2", "k1", "key"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n [0.0 1.1 3.0 3.0 3.0 99999.99 -1.1 -99999.99 0.0 1.1 3.0 3.0 3.0 99999.99 -1.1 -99999.99 0.0 1.1 3.0 3.0 3.0 99999.99 -1.1 -99999.99 0.0 1.1 3.0 3.0 3.0 99999.99 -1.1 -99999.99 0.0 1.1 3.0 3.0 3.0 99999.99 -1.1 -99999.99 0.0 1.1 3.0 3.0 3.0 99999.99 -1.1 -99999.99 0.0 1.1 3.0 3.0 3.0 99999.99 -1.1 -99999.99 0.0 1.1 3.0 3.0 3.0 99999.99 -1.1 -99999.99]  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{flts.New(0.0, 1.1, 3.0, 3.0, 3.0, 99999.99, -1.1, -99999.99, 0.0, 1.1, 3.0, 3.0, 3.0, 99999.99, -1.1, -99999.99, 0.0, 1.1, 3.0, 3.0, 3.0, 99999.99, -1.1, -99999.99, 0.0, 1.1, 3.0, 3.0, 3.0, 99999.99, -1.1, -99999.99, 0.0, 1.1, 3.0, 3.0, 3.0, 99999.99, -1.1, -99999.99, 0.0, 1.1, 3.0, 3.0, 3.0, 99999.99, -1.1, -99999.99, 0.0, 1.1, 3.0, 3.0, 3.0, 99999.99, -1.1, -99999.99, 0.0, 1.1, 3.0, 3.0, 3.0, 99999.99, -1.1, -99999.99), []string{"k2", "k1", "key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n [0.0 1.1 3.0 3.0 3.0 99999.99 -1.1 -99999.99 0.0 1.1 3.0 3.0 3.0 99999.99 -1.1 -99999.99 0.0 1.1 3.0 3.0 3.0 99999.99 -1.1 -99999.99 0.0 1.1 3.0 3.0 3.0 99999.99 -1.1 -99999.99 0.0 1.1 3.0 3.0 3.0 99999.99 -1.1 -99999.99 0.0 1.1 3.0 3.0 3.0 99999.99 -1.1 -99999.99 0.0 1.1 3.0 3.0 3.0 99999.99 -1.1 -99999.99 0.0 1.1 3.0 3.0 3.0 99999.99 -1.1 -99999.99]  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
	}
	for _, cse := range cses {
		t.Run("Flts", func(t *testing.T) {
			var c cfg.Cfgr
			c.Reset(cse.txt)
			a := c.Flts(cse.pth...)
			tst.FltsEql(t, cse.e, a)
		})
	}
}
func TestCfgFltsInvalid(t *testing.T) {
	cses := []struct {
		e   *flts.Flts
		pth []string
		txt string
	}{
		{flts.New(0.0, 1.1, 3.0, 3.0, 3.0, 99999.99, -1.1, -99999.99), []string{}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n [0.0 1.1 3.0 3.0 3.0 99999.99 -1.1 -99999.99]  \t\r\n // comment\n }"},
		{flts.New(0.0, 1.1, 3.0, 3.0, 3.0, 99999.99, -1.1, -99999.99), []string{"wrong"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n [0.0 1.1 3.0 3.0 3.0 99999.99 -1.1 -99999.99]  \t\r\n // comment\n }"},
		{flts.New(0.0, 1.1, 3.0, 3.0, 3.0, 99999.99, -1.1, -99999.99), []string{"wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n [0.0 1.1 3.0 3.0 3.0 99999.99 -1.1 -99999.99]  \t\r\n // comment\n  no3:\"no\" }"},
		{flts.New(0.0, 1.1, 3.0, 3.0, 3.0, 99999.99, -1.1, -99999.99), []string{"k2", "k1", "wrong"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n [0.0 1.1 3.0 3.0 3.0 99999.99 -1.1 -99999.99]  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{flts.New(0.0, 1.1, 3.0, 3.0, 3.0, 99999.99, -1.1, -99999.99), []string{"k2", "k1", "wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n [0.0 1.1 3.0 3.0 3.0 99999.99 -1.1 -99999.99]  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{flts.New(0.0, 1.1, 3.0, 3.0, 3.0, 99999.99, -1.1, -99999.99, 0.0, 1.1, 3.0, 3.0, 3.0, 99999.99, -1.1, -99999.99, 0.0, 1.1, 3.0, 3.0, 3.0, 99999.99, -1.1, -99999.99, 0.0, 1.1, 3.0, 3.0, 3.0, 99999.99, -1.1, -99999.99, 0.0, 1.1, 3.0, 3.0, 3.0, 99999.99, -1.1, -99999.99, 0.0, 1.1, 3.0, 3.0, 3.0, 99999.99, -1.1, -99999.99, 0.0, 1.1, 3.0, 3.0, 3.0, 99999.99, -1.1, -99999.99, 0.0, 1.1, 3.0, 3.0, 3.0, 99999.99, -1.1, -99999.99), []string{}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n [0.0 1.1 3.0 3.0 3.0 99999.99 -1.1 -99999.99 0.0 1.1 3.0 3.0 3.0 99999.99 -1.1 -99999.99 0.0 1.1 3.0 3.0 3.0 99999.99 -1.1 -99999.99 0.0 1.1 3.0 3.0 3.0 99999.99 -1.1 -99999.99 0.0 1.1 3.0 3.0 3.0 99999.99 -1.1 -99999.99 0.0 1.1 3.0 3.0 3.0 99999.99 -1.1 -99999.99 0.0 1.1 3.0 3.0 3.0 99999.99 -1.1 -99999.99 0.0 1.1 3.0 3.0 3.0 99999.99 -1.1 -99999.99]  \t\r\n // comment\n }"},
		{flts.New(0.0, 1.1, 3.0, 3.0, 3.0, 99999.99, -1.1, -99999.99, 0.0, 1.1, 3.0, 3.0, 3.0, 99999.99, -1.1, -99999.99, 0.0, 1.1, 3.0, 3.0, 3.0, 99999.99, -1.1, -99999.99, 0.0, 1.1, 3.0, 3.0, 3.0, 99999.99, -1.1, -99999.99, 0.0, 1.1, 3.0, 3.0, 3.0, 99999.99, -1.1, -99999.99, 0.0, 1.1, 3.0, 3.0, 3.0, 99999.99, -1.1, -99999.99, 0.0, 1.1, 3.0, 3.0, 3.0, 99999.99, -1.1, -99999.99, 0.0, 1.1, 3.0, 3.0, 3.0, 99999.99, -1.1, -99999.99), []string{"wrong"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n [0.0 1.1 3.0 3.0 3.0 99999.99 -1.1 -99999.99 0.0 1.1 3.0 3.0 3.0 99999.99 -1.1 -99999.99 0.0 1.1 3.0 3.0 3.0 99999.99 -1.1 -99999.99 0.0 1.1 3.0 3.0 3.0 99999.99 -1.1 -99999.99 0.0 1.1 3.0 3.0 3.0 99999.99 -1.1 -99999.99 0.0 1.1 3.0 3.0 3.0 99999.99 -1.1 -99999.99 0.0 1.1 3.0 3.0 3.0 99999.99 -1.1 -99999.99 0.0 1.1 3.0 3.0 3.0 99999.99 -1.1 -99999.99]  \t\r\n // comment\n }"},
		{flts.New(0.0, 1.1, 3.0, 3.0, 3.0, 99999.99, -1.1, -99999.99, 0.0, 1.1, 3.0, 3.0, 3.0, 99999.99, -1.1, -99999.99, 0.0, 1.1, 3.0, 3.0, 3.0, 99999.99, -1.1, -99999.99, 0.0, 1.1, 3.0, 3.0, 3.0, 99999.99, -1.1, -99999.99, 0.0, 1.1, 3.0, 3.0, 3.0, 99999.99, -1.1, -99999.99, 0.0, 1.1, 3.0, 3.0, 3.0, 99999.99, -1.1, -99999.99, 0.0, 1.1, 3.0, 3.0, 3.0, 99999.99, -1.1, -99999.99, 0.0, 1.1, 3.0, 3.0, 3.0, 99999.99, -1.1, -99999.99), []string{"wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n [0.0 1.1 3.0 3.0 3.0 99999.99 -1.1 -99999.99 0.0 1.1 3.0 3.0 3.0 99999.99 -1.1 -99999.99 0.0 1.1 3.0 3.0 3.0 99999.99 -1.1 -99999.99 0.0 1.1 3.0 3.0 3.0 99999.99 -1.1 -99999.99 0.0 1.1 3.0 3.0 3.0 99999.99 -1.1 -99999.99 0.0 1.1 3.0 3.0 3.0 99999.99 -1.1 -99999.99 0.0 1.1 3.0 3.0 3.0 99999.99 -1.1 -99999.99 0.0 1.1 3.0 3.0 3.0 99999.99 -1.1 -99999.99]  \t\r\n // comment\n  no3:\"no\" }"},
		{flts.New(0.0, 1.1, 3.0, 3.0, 3.0, 99999.99, -1.1, -99999.99, 0.0, 1.1, 3.0, 3.0, 3.0, 99999.99, -1.1, -99999.99, 0.0, 1.1, 3.0, 3.0, 3.0, 99999.99, -1.1, -99999.99, 0.0, 1.1, 3.0, 3.0, 3.0, 99999.99, -1.1, -99999.99, 0.0, 1.1, 3.0, 3.0, 3.0, 99999.99, -1.1, -99999.99, 0.0, 1.1, 3.0, 3.0, 3.0, 99999.99, -1.1, -99999.99, 0.0, 1.1, 3.0, 3.0, 3.0, 99999.99, -1.1, -99999.99, 0.0, 1.1, 3.0, 3.0, 3.0, 99999.99, -1.1, -99999.99), []string{"k2", "k1", "wrong"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n [0.0 1.1 3.0 3.0 3.0 99999.99 -1.1 -99999.99 0.0 1.1 3.0 3.0 3.0 99999.99 -1.1 -99999.99 0.0 1.1 3.0 3.0 3.0 99999.99 -1.1 -99999.99 0.0 1.1 3.0 3.0 3.0 99999.99 -1.1 -99999.99 0.0 1.1 3.0 3.0 3.0 99999.99 -1.1 -99999.99 0.0 1.1 3.0 3.0 3.0 99999.99 -1.1 -99999.99 0.0 1.1 3.0 3.0 3.0 99999.99 -1.1 -99999.99 0.0 1.1 3.0 3.0 3.0 99999.99 -1.1 -99999.99]  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{flts.New(0.0, 1.1, 3.0, 3.0, 3.0, 99999.99, -1.1, -99999.99, 0.0, 1.1, 3.0, 3.0, 3.0, 99999.99, -1.1, -99999.99, 0.0, 1.1, 3.0, 3.0, 3.0, 99999.99, -1.1, -99999.99, 0.0, 1.1, 3.0, 3.0, 3.0, 99999.99, -1.1, -99999.99, 0.0, 1.1, 3.0, 3.0, 3.0, 99999.99, -1.1, -99999.99, 0.0, 1.1, 3.0, 3.0, 3.0, 99999.99, -1.1, -99999.99, 0.0, 1.1, 3.0, 3.0, 3.0, 99999.99, -1.1, -99999.99, 0.0, 1.1, 3.0, 3.0, 3.0, 99999.99, -1.1, -99999.99), []string{"k2", "k1", "wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n [0.0 1.1 3.0 3.0 3.0 99999.99 -1.1 -99999.99 0.0 1.1 3.0 3.0 3.0 99999.99 -1.1 -99999.99 0.0 1.1 3.0 3.0 3.0 99999.99 -1.1 -99999.99 0.0 1.1 3.0 3.0 3.0 99999.99 -1.1 -99999.99 0.0 1.1 3.0 3.0 3.0 99999.99 -1.1 -99999.99 0.0 1.1 3.0 3.0 3.0 99999.99 -1.1 -99999.99 0.0 1.1 3.0 3.0 3.0 99999.99 -1.1 -99999.99 0.0 1.1 3.0 3.0 3.0 99999.99 -1.1 -99999.99]  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{flts.New(), []string{}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n []  \t\r\n // comment\n }"},
		{flts.New(), []string{"wrong"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n []  \t\r\n // comment\n }"},
		{flts.New(), []string{"wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n []  \t\r\n // comment\n  no3:\"no\" }"},
		{flts.New(), []string{"k2", "k1", "wrong"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n []  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{flts.New(), []string{"k2", "k1", "wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n []  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
	}
	var c cfg.Cfgr
	for _, cse := range cses {
		t.Run("Flts", func(t *testing.T) {
			c.Reset(cse.txt)
			tst.Panic(t, func() { c.Flts(cse.pth...) })
		})
	}
}
func TestCfgUntsValid(t *testing.T) {
	cses := []struct {
		e   *unts.Unts
		pth []string
		txt string
	}{
		{unts.New(0, 1, 1000), []string{"key"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n [0 1 1000]  \t\r\n // comment\n }"},
		{unts.New(0, 1, 1000), []string{"key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n [0 1 1000]  \t\r\n // comment\n  no3:\"no\" }"},
		{unts.New(0, 1, 1000), []string{"k2", "k1", "key"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n [0 1 1000]  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{unts.New(0, 1, 1000), []string{"k2", "k1", "key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n [0 1 1000]  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{unts.New(0, 1, 1000, 10), []string{"key"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n [0 1 1000 10]  \t\r\n // comment\n }"},
		{unts.New(0, 1, 1000, 10), []string{"key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n [0 1 1000 10]  \t\r\n // comment\n  no3:\"no\" }"},
		{unts.New(0, 1, 1000, 10), []string{"k2", "k1", "key"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n [0 1 1000 10]  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{unts.New(0, 1, 1000, 10), []string{"k2", "k1", "key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n [0 1 1000 10]  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
	}
	for _, cse := range cses {
		t.Run("Unts", func(t *testing.T) {
			var c cfg.Cfgr
			c.Reset(cse.txt)
			a := c.Unts(cse.pth...)
			tst.UntsEql(t, cse.e, a)
		})
	}
}
func TestCfgUntsInvalid(t *testing.T) {
	cses := []struct {
		e   *unts.Unts
		pth []string
		txt string
	}{
		{unts.New(0, 1, 1000, 10), []string{}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n [0 1 1000 10]  \t\r\n // comment\n }"},
		{unts.New(0, 1, 1000, 10), []string{"wrong"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n [0 1 1000 10]  \t\r\n // comment\n }"},
		{unts.New(0, 1, 1000, 10), []string{"wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n [0 1 1000 10]  \t\r\n // comment\n  no3:\"no\" }"},
		{unts.New(0, 1, 1000, 10), []string{"k2", "k1", "wrong"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n [0 1 1000 10]  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{unts.New(0, 1, 1000, 10), []string{"k2", "k1", "wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n [0 1 1000 10]  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{unts.New(), []string{}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n []  \t\r\n // comment\n }"},
		{unts.New(), []string{"wrong"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n []  \t\r\n // comment\n }"},
		{unts.New(), []string{"wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n []  \t\r\n // comment\n  no3:\"no\" }"},
		{unts.New(), []string{"k2", "k1", "wrong"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n []  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{unts.New(), []string{"k2", "k1", "wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n []  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
	}
	var c cfg.Cfgr
	for _, cse := range cses {
		t.Run("Unts", func(t *testing.T) {
			c.Reset(cse.txt)
			tst.Panic(t, func() { c.Unts(cse.pth...) })
		})
	}
}
func TestCfgIntsValid(t *testing.T) {
	cses := []struct {
		e   *ints.Ints
		pth []string
		txt string
	}{
		{ints.New(0, 10, 1000), []string{"key"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n [+0 +10 +1000]  \t\r\n // comment\n }"},
		{ints.New(0, 10, 1000), []string{"key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n [+0 +10 +1000]  \t\r\n // comment\n  no3:\"no\" }"},
		{ints.New(0, 10, 1000), []string{"k2", "k1", "key"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n [+0 +10 +1000]  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{ints.New(0, 10, 1000), []string{"k2", "k1", "key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n [+0 +10 +1000]  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{ints.New(0, 10, 1000, -10, -1000), []string{"key"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n [+0 +10 +1000 -10 -1000]  \t\r\n // comment\n }"},
		{ints.New(0, 10, 1000, -10, -1000), []string{"key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n [+0 +10 +1000 -10 -1000]  \t\r\n // comment\n  no3:\"no\" }"},
		{ints.New(0, 10, 1000, -10, -1000), []string{"k2", "k1", "key"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n [+0 +10 +1000 -10 -1000]  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{ints.New(0, 10, 1000, -10, -1000), []string{"k2", "k1", "key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n [+0 +10 +1000 -10 -1000]  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
	}
	for _, cse := range cses {
		t.Run("Ints", func(t *testing.T) {
			var c cfg.Cfgr
			c.Reset(cse.txt)
			a := c.Ints(cse.pth...)
			tst.IntsEql(t, cse.e, a)
		})
	}
}
func TestCfgIntsInvalid(t *testing.T) {
	cses := []struct {
		e   *ints.Ints
		pth []string
		txt string
	}{
		{ints.New(0, 10, 1000, -10, -1000), []string{}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n [+0 +10 +1000 -10 -1000]  \t\r\n // comment\n }"},
		{ints.New(0, 10, 1000, -10, -1000), []string{"wrong"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n [+0 +10 +1000 -10 -1000]  \t\r\n // comment\n }"},
		{ints.New(0, 10, 1000, -10, -1000), []string{"wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n [+0 +10 +1000 -10 -1000]  \t\r\n // comment\n  no3:\"no\" }"},
		{ints.New(0, 10, 1000, -10, -1000), []string{"k2", "k1", "wrong"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n [+0 +10 +1000 -10 -1000]  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{ints.New(0, 10, 1000, -10, -1000), []string{"k2", "k1", "wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n [+0 +10 +1000 -10 -1000]  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{ints.New(), []string{}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n []  \t\r\n // comment\n }"},
		{ints.New(), []string{"wrong"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n []  \t\r\n // comment\n }"},
		{ints.New(), []string{"wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n []  \t\r\n // comment\n  no3:\"no\" }"},
		{ints.New(), []string{"k2", "k1", "wrong"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n []  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{ints.New(), []string{"k2", "k1", "wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n []  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
	}
	var c cfg.Cfgr
	for _, cse := range cses {
		t.Run("Ints", func(t *testing.T) {
			c.Reset(cse.txt)
			tst.Panic(t, func() { c.Ints(cse.pth...) })
		})
	}
}
func TestCfgTmesValid(t *testing.T) {
	cses := []struct {
		e   *tmes.Tmes
		pth []string
		txt string
	}{
		{tmes.New(0, 1, 10), []string{"key"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n [0s 1s 10s]  \t\r\n // comment\n }"},
		{tmes.New(0, 1, 10), []string{"key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n [0s 1s 10s]  \t\r\n // comment\n  no3:\"no\" }"},
		{tmes.New(0, 1, 10), []string{"k2", "k1", "key"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n [0s 1s 10s]  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{tmes.New(0, 1, 10), []string{"k2", "k1", "key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n [0s 1s 10s]  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{tmes.New(0, 1, 10, -1, -10*60, 788645, -788645, 946782245, 946782245, 946695845), []string{"key"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n [0s 1s 10s -1s -10m 1w2d3h4m5s -1w2d3h4m5s 2000y1n2d3h4m5s 2000y2d3h4m5s 2000y3h4m5s]  \t\r\n // comment\n }"},
		{tmes.New(0, 1, 10, -1, -10*60, 788645, -788645, 946782245, 946782245, 946695845), []string{"key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n [0s 1s 10s -1s -10m 1w2d3h4m5s -1w2d3h4m5s 2000y1n2d3h4m5s 2000y2d3h4m5s 2000y3h4m5s]  \t\r\n // comment\n  no3:\"no\" }"},
		{tmes.New(0, 1, 10, -1, -10*60, 788645, -788645, 946782245, 946782245, 946695845), []string{"k2", "k1", "key"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n [0s 1s 10s -1s -10m 1w2d3h4m5s -1w2d3h4m5s 2000y1n2d3h4m5s 2000y2d3h4m5s 2000y3h4m5s]  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{tmes.New(0, 1, 10, -1, -10*60, 788645, -788645, 946782245, 946782245, 946695845), []string{"k2", "k1", "key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n [0s 1s 10s -1s -10m 1w2d3h4m5s -1w2d3h4m5s 2000y1n2d3h4m5s 2000y2d3h4m5s 2000y3h4m5s]  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
	}
	for _, cse := range cses {
		t.Run("Tmes", func(t *testing.T) {
			var c cfg.Cfgr
			c.Reset(cse.txt)
			a := c.Tmes(cse.pth...)
			tst.TmesEql(t, cse.e, a)
		})
	}
}
func TestCfgTmesInvalid(t *testing.T) {
	cses := []struct {
		e   *tmes.Tmes
		pth []string
		txt string
	}{
		{tmes.New(0, 1, 10, -1, -10*60, 788645, -788645, 946782245, 946782245, 946695845), []string{}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n [0s 1s 10s -1s -10m 1w2d3h4m5s -1w2d3h4m5s 2000y1n2d3h4m5s 2000y2d3h4m5s 2000y3h4m5s]  \t\r\n // comment\n }"},
		{tmes.New(0, 1, 10, -1, -10*60, 788645, -788645, 946782245, 946782245, 946695845), []string{"wrong"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n [0s 1s 10s -1s -10m 1w2d3h4m5s -1w2d3h4m5s 2000y1n2d3h4m5s 2000y2d3h4m5s 2000y3h4m5s]  \t\r\n // comment\n }"},
		{tmes.New(0, 1, 10, -1, -10*60, 788645, -788645, 946782245, 946782245, 946695845), []string{"wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n [0s 1s 10s -1s -10m 1w2d3h4m5s -1w2d3h4m5s 2000y1n2d3h4m5s 2000y2d3h4m5s 2000y3h4m5s]  \t\r\n // comment\n  no3:\"no\" }"},
		{tmes.New(0, 1, 10, -1, -10*60, 788645, -788645, 946782245, 946782245, 946695845), []string{"k2", "k1", "wrong"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n [0s 1s 10s -1s -10m 1w2d3h4m5s -1w2d3h4m5s 2000y1n2d3h4m5s 2000y2d3h4m5s 2000y3h4m5s]  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{tmes.New(0, 1, 10, -1, -10*60, 788645, -788645, 946782245, 946782245, 946695845), []string{"k2", "k1", "wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n [0s 1s 10s -1s -10m 1w2d3h4m5s -1w2d3h4m5s 2000y1n2d3h4m5s 2000y2d3h4m5s 2000y3h4m5s]  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{tmes.New(), []string{}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n []  \t\r\n // comment\n }"},
		{tmes.New(), []string{"wrong"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n []  \t\r\n // comment\n }"},
		{tmes.New(), []string{"wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n []  \t\r\n // comment\n  no3:\"no\" }"},
		{tmes.New(), []string{"k2", "k1", "wrong"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n []  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{tmes.New(), []string{"k2", "k1", "wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n []  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
	}
	var c cfg.Cfgr
	for _, cse := range cses {
		t.Run("Tmes", func(t *testing.T) {
			c.Reset(cse.txt)
			tst.Panic(t, func() { c.Tmes(cse.pth...) })
		})
	}
}
func TestCfgBndsValid(t *testing.T) {
	cses := []struct {
		e   *bnds.Bnds
		pth []string
		txt string
	}{
		{bnds.New(bnd.Bnd{Idx: 0, Lim: 0}, bnd.Bnd{Idx: 0, Lim: 1}, bnd.Bnd{Idx: 0, Lim: 1000}), []string{"key"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n [0-0 0-1 0-1000]  \t\r\n // comment\n }"},
		{bnds.New(bnd.Bnd{Idx: 0, Lim: 0}, bnd.Bnd{Idx: 0, Lim: 1}, bnd.Bnd{Idx: 0, Lim: 1000}), []string{"key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n [0-0 0-1 0-1000]  \t\r\n // comment\n  no3:\"no\" }"},
		{bnds.New(bnd.Bnd{Idx: 0, Lim: 0}, bnd.Bnd{Idx: 0, Lim: 1}, bnd.Bnd{Idx: 0, Lim: 1000}), []string{"k2", "k1", "key"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n [0-0 0-1 0-1000]  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{bnds.New(bnd.Bnd{Idx: 0, Lim: 0}, bnd.Bnd{Idx: 0, Lim: 1}, bnd.Bnd{Idx: 0, Lim: 1000}), []string{"k2", "k1", "key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n [0-0 0-1 0-1000]  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{bnds.New(bnd.Bnd{Idx: 0, Lim: 0}, bnd.Bnd{Idx: 0, Lim: 1}, bnd.Bnd{Idx: 0, Lim: 1000}, bnd.Bnd{Idx: 999, Lim: 1000}, bnd.Bnd{Idx: 1, Lim: 0}), []string{"key"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n [0-0 0-1 0-1000 999-1000 1-0]  \t\r\n // comment\n }"},
		{bnds.New(bnd.Bnd{Idx: 0, Lim: 0}, bnd.Bnd{Idx: 0, Lim: 1}, bnd.Bnd{Idx: 0, Lim: 1000}, bnd.Bnd{Idx: 999, Lim: 1000}, bnd.Bnd{Idx: 1, Lim: 0}), []string{"key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n [0-0 0-1 0-1000 999-1000 1-0]  \t\r\n // comment\n  no3:\"no\" }"},
		{bnds.New(bnd.Bnd{Idx: 0, Lim: 0}, bnd.Bnd{Idx: 0, Lim: 1}, bnd.Bnd{Idx: 0, Lim: 1000}, bnd.Bnd{Idx: 999, Lim: 1000}, bnd.Bnd{Idx: 1, Lim: 0}), []string{"k2", "k1", "key"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n [0-0 0-1 0-1000 999-1000 1-0]  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{bnds.New(bnd.Bnd{Idx: 0, Lim: 0}, bnd.Bnd{Idx: 0, Lim: 1}, bnd.Bnd{Idx: 0, Lim: 1000}, bnd.Bnd{Idx: 999, Lim: 1000}, bnd.Bnd{Idx: 1, Lim: 0}), []string{"k2", "k1", "key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n [0-0 0-1 0-1000 999-1000 1-0]  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
	}
	for _, cse := range cses {
		t.Run("Bnds", func(t *testing.T) {
			var c cfg.Cfgr
			c.Reset(cse.txt)
			a := c.Bnds(cse.pth...)
			tst.BndsEql(t, cse.e, a)
		})
	}
}
func TestCfgBndsInvalid(t *testing.T) {
	cses := []struct {
		e   *bnds.Bnds
		pth []string
		txt string
	}{
		{bnds.New(bnd.Bnd{Idx: 0, Lim: 0}, bnd.Bnd{Idx: 0, Lim: 1}, bnd.Bnd{Idx: 0, Lim: 1000}, bnd.Bnd{Idx: 999, Lim: 1000}, bnd.Bnd{Idx: 1, Lim: 0}), []string{}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n [0-0 0-1 0-1000 999-1000 1-0]  \t\r\n // comment\n }"},
		{bnds.New(bnd.Bnd{Idx: 0, Lim: 0}, bnd.Bnd{Idx: 0, Lim: 1}, bnd.Bnd{Idx: 0, Lim: 1000}, bnd.Bnd{Idx: 999, Lim: 1000}, bnd.Bnd{Idx: 1, Lim: 0}), []string{"wrong"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n [0-0 0-1 0-1000 999-1000 1-0]  \t\r\n // comment\n }"},
		{bnds.New(bnd.Bnd{Idx: 0, Lim: 0}, bnd.Bnd{Idx: 0, Lim: 1}, bnd.Bnd{Idx: 0, Lim: 1000}, bnd.Bnd{Idx: 999, Lim: 1000}, bnd.Bnd{Idx: 1, Lim: 0}), []string{"wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n [0-0 0-1 0-1000 999-1000 1-0]  \t\r\n // comment\n  no3:\"no\" }"},
		{bnds.New(bnd.Bnd{Idx: 0, Lim: 0}, bnd.Bnd{Idx: 0, Lim: 1}, bnd.Bnd{Idx: 0, Lim: 1000}, bnd.Bnd{Idx: 999, Lim: 1000}, bnd.Bnd{Idx: 1, Lim: 0}), []string{"k2", "k1", "wrong"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n [0-0 0-1 0-1000 999-1000 1-0]  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{bnds.New(bnd.Bnd{Idx: 0, Lim: 0}, bnd.Bnd{Idx: 0, Lim: 1}, bnd.Bnd{Idx: 0, Lim: 1000}, bnd.Bnd{Idx: 999, Lim: 1000}, bnd.Bnd{Idx: 1, Lim: 0}), []string{"k2", "k1", "wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n [0-0 0-1 0-1000 999-1000 1-0]  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{bnds.New(), []string{}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n []  \t\r\n // comment\n }"},
		{bnds.New(), []string{"wrong"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n []  \t\r\n // comment\n }"},
		{bnds.New(), []string{"wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n []  \t\r\n // comment\n  no3:\"no\" }"},
		{bnds.New(), []string{"k2", "k1", "wrong"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n []  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{bnds.New(), []string{"k2", "k1", "wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n []  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
	}
	var c cfg.Cfgr
	for _, cse := range cses {
		t.Run("Bnds", func(t *testing.T) {
			c.Reset(cse.txt)
			tst.Panic(t, func() { c.Bnds(cse.pth...) })
		})
	}
}
func TestCfgTmeRngsValid(t *testing.T) {
	cses := []struct {
		e   *tme.Rngs
		pth []string
		txt string
	}{
		{tme.NewRngs(tme.Rng{Min: -10, Max: -1}, tme.Rng{Min: 0, Max: 1}, tme.Rng{Min: 2, Max: 4}), []string{"key"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n [-10s--1s 0s-1s 2s-4s]  \t\r\n // comment\n }"},
		{tme.NewRngs(tme.Rng{Min: -10, Max: -1}, tme.Rng{Min: 0, Max: 1}, tme.Rng{Min: 2, Max: 4}), []string{"key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n [-10s--1s 0s-1s 2s-4s]  \t\r\n // comment\n  no3:\"no\" }"},
		{tme.NewRngs(tme.Rng{Min: -10, Max: -1}, tme.Rng{Min: 0, Max: 1}, tme.Rng{Min: 2, Max: 4}), []string{"k2", "k1", "key"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n [-10s--1s 0s-1s 2s-4s]  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{tme.NewRngs(tme.Rng{Min: -10, Max: -1}, tme.Rng{Min: 0, Max: 1}, tme.Rng{Min: 2, Max: 4}), []string{"k2", "k1", "key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n [-10s--1s 0s-1s 2s-4s]  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{tme.NewRngs(tme.Rng{Min: -10, Max: -1}, tme.Rng{Min: 0, Max: 1}, tme.Rng{Min: 2, Max: 4}, tme.Rng{Min: 6, Max: 10}, tme.Rng{Min: 50, Max: 60}), []string{"key"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n [-10s--1s 0s-1s 2s-4s 6s-10s 50s-60s]  \t\r\n // comment\n }"},
		{tme.NewRngs(tme.Rng{Min: -10, Max: -1}, tme.Rng{Min: 0, Max: 1}, tme.Rng{Min: 2, Max: 4}, tme.Rng{Min: 6, Max: 10}, tme.Rng{Min: 50, Max: 60}), []string{"key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n [-10s--1s 0s-1s 2s-4s 6s-10s 50s-60s]  \t\r\n // comment\n  no3:\"no\" }"},
		{tme.NewRngs(tme.Rng{Min: -10, Max: -1}, tme.Rng{Min: 0, Max: 1}, tme.Rng{Min: 2, Max: 4}, tme.Rng{Min: 6, Max: 10}, tme.Rng{Min: 50, Max: 60}), []string{"k2", "k1", "key"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n [-10s--1s 0s-1s 2s-4s 6s-10s 50s-60s]  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{tme.NewRngs(tme.Rng{Min: -10, Max: -1}, tme.Rng{Min: 0, Max: 1}, tme.Rng{Min: 2, Max: 4}, tme.Rng{Min: 6, Max: 10}, tme.Rng{Min: 50, Max: 60}), []string{"k2", "k1", "key"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n [-10s--1s 0s-1s 2s-4s 6s-10s 50s-60s]  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
	}
	for _, cse := range cses {
		t.Run("TmeRngs", func(t *testing.T) {
			var c cfg.Cfgr
			c.Reset(cse.txt)
			a := c.TmeRngs(cse.pth...)
			tst.TmeRngsEql(t, cse.e, a)
		})
	}
}
func TestCfgTmeRngsInvalid(t *testing.T) {
	cses := []struct {
		e   *tme.Rngs
		pth []string
		txt string
	}{
		{tme.NewRngs(tme.Rng{Min: -10, Max: -1}, tme.Rng{Min: 0, Max: 1}, tme.Rng{Min: 2, Max: 4}, tme.Rng{Min: 6, Max: 10}, tme.Rng{Min: 50, Max: 60}), []string{}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n [-10s--1s 0s-1s 2s-4s 6s-10s 50s-60s]  \t\r\n // comment\n }"},
		{tme.NewRngs(tme.Rng{Min: -10, Max: -1}, tme.Rng{Min: 0, Max: 1}, tme.Rng{Min: 2, Max: 4}, tme.Rng{Min: 6, Max: 10}, tme.Rng{Min: 50, Max: 60}), []string{"wrong"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n [-10s--1s 0s-1s 2s-4s 6s-10s 50s-60s]  \t\r\n // comment\n }"},
		{tme.NewRngs(tme.Rng{Min: -10, Max: -1}, tme.Rng{Min: 0, Max: 1}, tme.Rng{Min: 2, Max: 4}, tme.Rng{Min: 6, Max: 10}, tme.Rng{Min: 50, Max: 60}), []string{"wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n [-10s--1s 0s-1s 2s-4s 6s-10s 50s-60s]  \t\r\n // comment\n  no3:\"no\" }"},
		{tme.NewRngs(tme.Rng{Min: -10, Max: -1}, tme.Rng{Min: 0, Max: 1}, tme.Rng{Min: 2, Max: 4}, tme.Rng{Min: 6, Max: 10}, tme.Rng{Min: 50, Max: 60}), []string{"k2", "k1", "wrong"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n [-10s--1s 0s-1s 2s-4s 6s-10s 50s-60s]  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{tme.NewRngs(tme.Rng{Min: -10, Max: -1}, tme.Rng{Min: 0, Max: 1}, tme.Rng{Min: 2, Max: 4}, tme.Rng{Min: 6, Max: 10}, tme.Rng{Min: 50, Max: 60}), []string{"k2", "k1", "wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n [-10s--1s 0s-1s 2s-4s 6s-10s 50s-60s]  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{tme.NewRngs(), []string{}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n []  \t\r\n // comment\n }"},
		{tme.NewRngs(), []string{"wrong"}, "{  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n []  \t\r\n // comment\n }"},
		{tme.NewRngs(), []string{"wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n []  \t\r\n // comment\n  no3:\"no\" }"},
		{tme.NewRngs(), []string{"k2", "k1", "wrong"}, "{  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n k1  \t\r\n // comment\n :  \t\r\n // comment\n {  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n []  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
		{tme.NewRngs(), []string{"k2", "k1", "wrong"}, "{ no1:\"no\" no2:[\"no\" \"no\"]  \t\r\n // comment\n k2  \t\r\n // comment\n :  \t\r\n // comment\n { no3:\"no\"  \t\r\n // comment\n no4:[\"no\" \"no\"] k1  \t\r\n // comment\n :  \t\r\n // comment\n { no5:\"no\"  no6:[\"no\" \"no\"]  \t\r\n // comment\n key  \t\r\n // comment\n :  \t\r\n // comment\n []  \t\r\n // comment\n }  \t\r\n // comment\n }  \t\r\n // comment\n }"},
	}
	var c cfg.Cfgr
	for _, cse := range cses {
		t.Run("TmeRngs", func(t *testing.T) {
			c.Reset(cse.txt)
			tst.Panic(t, func() { c.TmeRngs(cse.pth...) })
		})
	}
}
