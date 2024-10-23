package jsn_test

import (
	"sys/bsc/bol"
	"sys/bsc/flt"
	"sys/bsc/int"
	"sys/bsc/str"
	"sys/lng/jsn"
	"sys/tst"
	"testing"
)

func TestJsnStrValid(t *testing.T) {
	cses := []struct {
		e   str.Str
		pth []string
		txt string
	}{
		{"", []string{"key"}, "{  \t\r\n  \"key\"  \t\r\n  :  \t\r\n  \"\"  \t\r\n  }"},
		{"", []string{"key"}, "{ \"no1\":\"no\", \"no2\":[\"no\" \"no\"],  \t\r\n  \"key\"  \t\r\n  :  \t\r\n  \"\",  \t\r\n   \"no3\":\"no\" }"},
		{"", []string{"k2", "k1", "key"}, "{  \t\r\n  \"k2\"  \t\r\n  :  \t\r\n  {  \t\r\n  \"k1\"  \t\r\n  :  \t\r\n  {  \t\r\n  \"key\"  \t\r\n  :  \t\r\n  \"\"  \t\r\n  }  \t\r\n  }  \t\r\n  }"},
		{"", []string{"k2", "k1", "key"}, "{ \"no1\":\"no\" \"no2\":[\"no\" \"no\"],  \t\r\n  \"k2\"  \t\r\n  :  \t\r\n  { \"no3\":\"no\",  \t\r\n  \"no4\":[\"no\" \"no\"], \"k1\"  \t\r\n  :  \t\r\n  { \"no5\":\"no\",  \"no6\":[\"no\" \"no\"],  \t\r\n  \"key\"  \t\r\n  :  \t\r\n  \"\"  \t\r\n  }  \t\r\n  }  \t\r\n  }"},
		{"xYz", []string{"key"}, "{  \t\r\n  \"key\"  \t\r\n  :  \t\r\n  \"xYz\"  \t\r\n  }"},
		{"xYz", []string{"key"}, "{ \"no1\":\"no\", \"no2\":[\"no\" \"no\"],  \t\r\n  \"key\"  \t\r\n  :  \t\r\n  \"xYz\",  \t\r\n   \"no3\":\"no\" }"},
		{"xYz", []string{"k2", "k1", "key"}, "{  \t\r\n  \"k2\"  \t\r\n  :  \t\r\n  {  \t\r\n  \"k1\"  \t\r\n  :  \t\r\n  {  \t\r\n  \"key\"  \t\r\n  :  \t\r\n  \"xYz\"  \t\r\n  }  \t\r\n  }  \t\r\n  }"},
		{"xYz", []string{"k2", "k1", "key"}, "{ \"no1\":\"no\" \"no2\":[\"no\" \"no\"],  \t\r\n  \"k2\"  \t\r\n  :  \t\r\n  { \"no3\":\"no\",  \t\r\n  \"no4\":[\"no\" \"no\"], \"k1\"  \t\r\n  :  \t\r\n  { \"no5\":\"no\",  \"no6\":[\"no\" \"no\"],  \t\r\n  \"key\"  \t\r\n  :  \t\r\n  \"xYz\"  \t\r\n  }  \t\r\n  }  \t\r\n  }"},
		{"a", []string{"key"}, "{  \t\r\n  \"key\"  \t\r\n  :  \t\r\n  \"a\"  \t\r\n  }"},
		{"a", []string{"key"}, "{ \"no1\":\"no\", \"no2\":[\"no\" \"no\"],  \t\r\n  \"key\"  \t\r\n  :  \t\r\n  \"a\",  \t\r\n   \"no3\":\"no\" }"},
		{"a", []string{"k2", "k1", "key"}, "{  \t\r\n  \"k2\"  \t\r\n  :  \t\r\n  {  \t\r\n  \"k1\"  \t\r\n  :  \t\r\n  {  \t\r\n  \"key\"  \t\r\n  :  \t\r\n  \"a\"  \t\r\n  }  \t\r\n  }  \t\r\n  }"},
		{"a", []string{"k2", "k1", "key"}, "{ \"no1\":\"no\" \"no2\":[\"no\" \"no\"],  \t\r\n  \"k2\"  \t\r\n  :  \t\r\n  { \"no3\":\"no\",  \t\r\n  \"no4\":[\"no\" \"no\"], \"k1\"  \t\r\n  :  \t\r\n  { \"no5\":\"no\",  \"no6\":[\"no\" \"no\"],  \t\r\n  \"key\"  \t\r\n  :  \t\r\n  \"a\"  \t\r\n  }  \t\r\n  }  \t\r\n  }"},
		{"efg HIJ jKl", []string{"key"}, "{  \t\r\n  \"key\"  \t\r\n  :  \t\r\n  \"efg HIJ jKl\"  \t\r\n  }"},
		{"efg HIJ jKl", []string{"key"}, "{ \"no1\":\"no\", \"no2\":[\"no\" \"no\"],  \t\r\n  \"key\"  \t\r\n  :  \t\r\n  \"efg HIJ jKl\",  \t\r\n   \"no3\":\"no\" }"},
		{"efg HIJ jKl", []string{"k2", "k1", "key"}, "{  \t\r\n  \"k2\"  \t\r\n  :  \t\r\n  {  \t\r\n  \"k1\"  \t\r\n  :  \t\r\n  {  \t\r\n  \"key\"  \t\r\n  :  \t\r\n  \"efg HIJ jKl\"  \t\r\n  }  \t\r\n  }  \t\r\n  }"},
		{"efg HIJ jKl", []string{"k2", "k1", "key"}, "{ \"no1\":\"no\" \"no2\":[\"no\" \"no\"],  \t\r\n  \"k2\"  \t\r\n  :  \t\r\n  { \"no3\":\"no\",  \t\r\n  \"no4\":[\"no\" \"no\"], \"k1\"  \t\r\n  :  \t\r\n  { \"no5\":\"no\",  \"no6\":[\"no\" \"no\"],  \t\r\n  \"key\"  \t\r\n  :  \t\r\n  \"efg HIJ jKl\"  \t\r\n  }  \t\r\n  }  \t\r\n  }"},
	}
	for _, cse := range cses {
		t.Run("Str", func(t *testing.T) {
			var j jsn.Jsnr
			j.Reset(cse.txt)
			a := j.Str(cse.pth...)
			tst.StrEql(t, cse.e, a)
		})
	}
}
func TestJsnBolValid(t *testing.T) {
	cses := []struct {
		e   bol.Bol
		pth []string
		txt string
	}{
		{false, []string{"key"}, "{  \t\r\n  \"key\"  \t\r\n  :  \t\r\n  false  \t\r\n  }"},
		{false, []string{"key"}, "{ \"no1\":\"no\", \"no2\":[\"no\" \"no\"],  \t\r\n  \"key\"  \t\r\n  :  \t\r\n  false,  \t\r\n   \"no3\":\"no\" }"},
		{false, []string{"k2", "k1", "key"}, "{  \t\r\n  \"k2\"  \t\r\n  :  \t\r\n  {  \t\r\n  \"k1\"  \t\r\n  :  \t\r\n  {  \t\r\n  \"key\"  \t\r\n  :  \t\r\n  false  \t\r\n  }  \t\r\n  }  \t\r\n  }"},
		{false, []string{"k2", "k1", "key"}, "{ \"no1\":\"no\" \"no2\":[\"no\" \"no\"],  \t\r\n  \"k2\"  \t\r\n  :  \t\r\n  { \"no3\":\"no\",  \t\r\n  \"no4\":[\"no\" \"no\"], \"k1\"  \t\r\n  :  \t\r\n  { \"no5\":\"no\",  \"no6\":[\"no\" \"no\"],  \t\r\n  \"key\"  \t\r\n  :  \t\r\n  false  \t\r\n  }  \t\r\n  }  \t\r\n  }"},
		{true, []string{"key"}, "{  \t\r\n  \"key\"  \t\r\n  :  \t\r\n  true  \t\r\n  }"},
		{true, []string{"key"}, "{ \"no1\":\"no\", \"no2\":[\"no\" \"no\"],  \t\r\n  \"key\"  \t\r\n  :  \t\r\n  true,  \t\r\n   \"no3\":\"no\" }"},
		{true, []string{"k2", "k1", "key"}, "{  \t\r\n  \"k2\"  \t\r\n  :  \t\r\n  {  \t\r\n  \"k1\"  \t\r\n  :  \t\r\n  {  \t\r\n  \"key\"  \t\r\n  :  \t\r\n  true  \t\r\n  }  \t\r\n  }  \t\r\n  }"},
		{true, []string{"k2", "k1", "key"}, "{ \"no1\":\"no\" \"no2\":[\"no\" \"no\"],  \t\r\n  \"k2\"  \t\r\n  :  \t\r\n  { \"no3\":\"no\",  \t\r\n  \"no4\":[\"no\" \"no\"], \"k1\"  \t\r\n  :  \t\r\n  { \"no5\":\"no\",  \"no6\":[\"no\" \"no\"],  \t\r\n  \"key\"  \t\r\n  :  \t\r\n  true  \t\r\n  }  \t\r\n  }  \t\r\n  }"},
	}
	for _, cse := range cses {
		t.Run("Bol", func(t *testing.T) {
			var j jsn.Jsnr
			j.Reset(cse.txt)
			a := j.Bol(cse.pth...)
			tst.BolEql(t, cse.e, a)
		})
	}
}
func TestJsnFltValid(t *testing.T) {
	cses := []struct {
		e   flt.Flt
		pth []string
		txt string
	}{
		{0.0, []string{"key"}, "{  \t\r\n  \"key\"  \t\r\n  :  \t\r\n  0.0  \t\r\n  }"},
		{0.0, []string{"key"}, "{ \"no1\":\"no\", \"no2\":[\"no\" \"no\"],  \t\r\n  \"key\"  \t\r\n  :  \t\r\n  0.0,  \t\r\n   \"no3\":\"no\" }"},
		{0.0, []string{"k2", "k1", "key"}, "{  \t\r\n  \"k2\"  \t\r\n  :  \t\r\n  {  \t\r\n  \"k1\"  \t\r\n  :  \t\r\n  {  \t\r\n  \"key\"  \t\r\n  :  \t\r\n  0.0  \t\r\n  }  \t\r\n  }  \t\r\n  }"},
		{0.0, []string{"k2", "k1", "key"}, "{ \"no1\":\"no\" \"no2\":[\"no\" \"no\"],  \t\r\n  \"k2\"  \t\r\n  :  \t\r\n  { \"no3\":\"no\",  \t\r\n  \"no4\":[\"no\" \"no\"], \"k1\"  \t\r\n  :  \t\r\n  { \"no5\":\"no\",  \"no6\":[\"no\" \"no\"],  \t\r\n  \"key\"  \t\r\n  :  \t\r\n  0.0  \t\r\n  }  \t\r\n  }  \t\r\n  }"},
		{1.1, []string{"key"}, "{  \t\r\n  \"key\"  \t\r\n  :  \t\r\n  1.1  \t\r\n  }"},
		{1.1, []string{"key"}, "{ \"no1\":\"no\", \"no2\":[\"no\" \"no\"],  \t\r\n  \"key\"  \t\r\n  :  \t\r\n  1.1,  \t\r\n   \"no3\":\"no\" }"},
		{1.1, []string{"k2", "k1", "key"}, "{  \t\r\n  \"k2\"  \t\r\n  :  \t\r\n  {  \t\r\n  \"k1\"  \t\r\n  :  \t\r\n  {  \t\r\n  \"key\"  \t\r\n  :  \t\r\n  1.1  \t\r\n  }  \t\r\n  }  \t\r\n  }"},
		{1.1, []string{"k2", "k1", "key"}, "{ \"no1\":\"no\" \"no2\":[\"no\" \"no\"],  \t\r\n  \"k2\"  \t\r\n  :  \t\r\n  { \"no3\":\"no\",  \t\r\n  \"no4\":[\"no\" \"no\"], \"k1\"  \t\r\n  :  \t\r\n  { \"no5\":\"no\",  \"no6\":[\"no\" \"no\"],  \t\r\n  \"key\"  \t\r\n  :  \t\r\n  1.1  \t\r\n  }  \t\r\n  }  \t\r\n  }"},
		{3.0, []string{"key"}, "{  \t\r\n  \"key\"  \t\r\n  :  \t\r\n  3.0  \t\r\n  }"},
		{3.0, []string{"key"}, "{ \"no1\":\"no\", \"no2\":[\"no\" \"no\"],  \t\r\n  \"key\"  \t\r\n  :  \t\r\n  3.0,  \t\r\n   \"no3\":\"no\" }"},
		{3.0, []string{"k2", "k1", "key"}, "{  \t\r\n  \"k2\"  \t\r\n  :  \t\r\n  {  \t\r\n  \"k1\"  \t\r\n  :  \t\r\n  {  \t\r\n  \"key\"  \t\r\n  :  \t\r\n  3.0  \t\r\n  }  \t\r\n  }  \t\r\n  }"},
		{3.0, []string{"k2", "k1", "key"}, "{ \"no1\":\"no\" \"no2\":[\"no\" \"no\"],  \t\r\n  \"k2\"  \t\r\n  :  \t\r\n  { \"no3\":\"no\",  \t\r\n  \"no4\":[\"no\" \"no\"], \"k1\"  \t\r\n  :  \t\r\n  { \"no5\":\"no\",  \"no6\":[\"no\" \"no\"],  \t\r\n  \"key\"  \t\r\n  :  \t\r\n  3.0  \t\r\n  }  \t\r\n  }  \t\r\n  }"},
		{3.0, []string{"key"}, "{  \t\r\n  \"key\"  \t\r\n  :  \t\r\n  3.0  \t\r\n  }"},
		{3.0, []string{"key"}, "{ \"no1\":\"no\", \"no2\":[\"no\" \"no\"],  \t\r\n  \"key\"  \t\r\n  :  \t\r\n  3.0,  \t\r\n   \"no3\":\"no\" }"},
		{3.0, []string{"k2", "k1", "key"}, "{  \t\r\n  \"k2\"  \t\r\n  :  \t\r\n  {  \t\r\n  \"k1\"  \t\r\n  :  \t\r\n  {  \t\r\n  \"key\"  \t\r\n  :  \t\r\n  3.0  \t\r\n  }  \t\r\n  }  \t\r\n  }"},
		{3.0, []string{"k2", "k1", "key"}, "{ \"no1\":\"no\" \"no2\":[\"no\" \"no\"],  \t\r\n  \"k2\"  \t\r\n  :  \t\r\n  { \"no3\":\"no\",  \t\r\n  \"no4\":[\"no\" \"no\"], \"k1\"  \t\r\n  :  \t\r\n  { \"no5\":\"no\",  \"no6\":[\"no\" \"no\"],  \t\r\n  \"key\"  \t\r\n  :  \t\r\n  3.0  \t\r\n  }  \t\r\n  }  \t\r\n  }"},
		{3.0, []string{"key"}, "{  \t\r\n  \"key\"  \t\r\n  :  \t\r\n  3.0  \t\r\n  }"},
		{3.0, []string{"key"}, "{ \"no1\":\"no\", \"no2\":[\"no\" \"no\"],  \t\r\n  \"key\"  \t\r\n  :  \t\r\n  3.0,  \t\r\n   \"no3\":\"no\" }"},
		{3.0, []string{"k2", "k1", "key"}, "{  \t\r\n  \"k2\"  \t\r\n  :  \t\r\n  {  \t\r\n  \"k1\"  \t\r\n  :  \t\r\n  {  \t\r\n  \"key\"  \t\r\n  :  \t\r\n  3.0  \t\r\n  }  \t\r\n  }  \t\r\n  }"},
		{3.0, []string{"k2", "k1", "key"}, "{ \"no1\":\"no\" \"no2\":[\"no\" \"no\"],  \t\r\n  \"k2\"  \t\r\n  :  \t\r\n  { \"no3\":\"no\",  \t\r\n  \"no4\":[\"no\" \"no\"], \"k1\"  \t\r\n  :  \t\r\n  { \"no5\":\"no\",  \"no6\":[\"no\" \"no\"],  \t\r\n  \"key\"  \t\r\n  :  \t\r\n  3.0  \t\r\n  }  \t\r\n  }  \t\r\n  }"},
		{99999.99, []string{"key"}, "{  \t\r\n  \"key\"  \t\r\n  :  \t\r\n  99999.99  \t\r\n  }"},
		{99999.99, []string{"key"}, "{ \"no1\":\"no\", \"no2\":[\"no\" \"no\"],  \t\r\n  \"key\"  \t\r\n  :  \t\r\n  99999.99,  \t\r\n   \"no3\":\"no\" }"},
		{99999.99, []string{"k2", "k1", "key"}, "{  \t\r\n  \"k2\"  \t\r\n  :  \t\r\n  {  \t\r\n  \"k1\"  \t\r\n  :  \t\r\n  {  \t\r\n  \"key\"  \t\r\n  :  \t\r\n  99999.99  \t\r\n  }  \t\r\n  }  \t\r\n  }"},
		{99999.99, []string{"k2", "k1", "key"}, "{ \"no1\":\"no\" \"no2\":[\"no\" \"no\"],  \t\r\n  \"k2\"  \t\r\n  :  \t\r\n  { \"no3\":\"no\",  \t\r\n  \"no4\":[\"no\" \"no\"], \"k1\"  \t\r\n  :  \t\r\n  { \"no5\":\"no\",  \"no6\":[\"no\" \"no\"],  \t\r\n  \"key\"  \t\r\n  :  \t\r\n  99999.99  \t\r\n  }  \t\r\n  }  \t\r\n  }"},
		{-1.1, []string{"key"}, "{  \t\r\n  \"key\"  \t\r\n  :  \t\r\n  -1.1  \t\r\n  }"},
		{-1.1, []string{"key"}, "{ \"no1\":\"no\", \"no2\":[\"no\" \"no\"],  \t\r\n  \"key\"  \t\r\n  :  \t\r\n  -1.1,  \t\r\n   \"no3\":\"no\" }"},
		{-1.1, []string{"k2", "k1", "key"}, "{  \t\r\n  \"k2\"  \t\r\n  :  \t\r\n  {  \t\r\n  \"k1\"  \t\r\n  :  \t\r\n  {  \t\r\n  \"key\"  \t\r\n  :  \t\r\n  -1.1  \t\r\n  }  \t\r\n  }  \t\r\n  }"},
		{-1.1, []string{"k2", "k1", "key"}, "{ \"no1\":\"no\" \"no2\":[\"no\" \"no\"],  \t\r\n  \"k2\"  \t\r\n  :  \t\r\n  { \"no3\":\"no\",  \t\r\n  \"no4\":[\"no\" \"no\"], \"k1\"  \t\r\n  :  \t\r\n  { \"no5\":\"no\",  \"no6\":[\"no\" \"no\"],  \t\r\n  \"key\"  \t\r\n  :  \t\r\n  -1.1  \t\r\n  }  \t\r\n  }  \t\r\n  }"},
		{-99999.99, []string{"key"}, "{  \t\r\n  \"key\"  \t\r\n  :  \t\r\n  -99999.99  \t\r\n  }"},
		{-99999.99, []string{"key"}, "{ \"no1\":\"no\", \"no2\":[\"no\" \"no\"],  \t\r\n  \"key\"  \t\r\n  :  \t\r\n  -99999.99,  \t\r\n   \"no3\":\"no\" }"},
		{-99999.99, []string{"k2", "k1", "key"}, "{  \t\r\n  \"k2\"  \t\r\n  :  \t\r\n  {  \t\r\n  \"k1\"  \t\r\n  :  \t\r\n  {  \t\r\n  \"key\"  \t\r\n  :  \t\r\n  -99999.99  \t\r\n  }  \t\r\n  }  \t\r\n  }"},
		{-99999.99, []string{"k2", "k1", "key"}, "{ \"no1\":\"no\" \"no2\":[\"no\" \"no\"],  \t\r\n  \"k2\"  \t\r\n  :  \t\r\n  { \"no3\":\"no\",  \t\r\n  \"no4\":[\"no\" \"no\"], \"k1\"  \t\r\n  :  \t\r\n  { \"no5\":\"no\",  \"no6\":[\"no\" \"no\"],  \t\r\n  \"key\"  \t\r\n  :  \t\r\n  -99999.99  \t\r\n  }  \t\r\n  }  \t\r\n  }"},
	}
	for _, cse := range cses {
		t.Run("Flt", func(t *testing.T) {
			var j jsn.Jsnr
			j.Reset(cse.txt)
			a := j.Flt(cse.pth...)
			tst.FltEql(t, cse.e, a)
		})
	}
}
func TestJsnIntValid(t *testing.T) {
	cses := []struct {
		e   int.Int
		pth []string
		txt string
	}{
		{0, []string{"key"}, "{  \t\r\n  \"key\"  \t\r\n  :  \t\r\n  0  \t\r\n  }"},
		{0, []string{"key"}, "{ \"no1\":\"no\", \"no2\":[\"no\" \"no\"],  \t\r\n  \"key\"  \t\r\n  :  \t\r\n  0,  \t\r\n   \"no3\":\"no\" }"},
		{0, []string{"k2", "k1", "key"}, "{  \t\r\n  \"k2\"  \t\r\n  :  \t\r\n  {  \t\r\n  \"k1\"  \t\r\n  :  \t\r\n  {  \t\r\n  \"key\"  \t\r\n  :  \t\r\n  0  \t\r\n  }  \t\r\n  }  \t\r\n  }"},
		{0, []string{"k2", "k1", "key"}, "{ \"no1\":\"no\" \"no2\":[\"no\" \"no\"],  \t\r\n  \"k2\"  \t\r\n  :  \t\r\n  { \"no3\":\"no\",  \t\r\n  \"no4\":[\"no\" \"no\"], \"k1\"  \t\r\n  :  \t\r\n  { \"no5\":\"no\",  \"no6\":[\"no\" \"no\"],  \t\r\n  \"key\"  \t\r\n  :  \t\r\n  0  \t\r\n  }  \t\r\n  }  \t\r\n  }"},
		{10, []string{"key"}, "{  \t\r\n  \"key\"  \t\r\n  :  \t\r\n  10  \t\r\n  }"},
		{10, []string{"key"}, "{ \"no1\":\"no\", \"no2\":[\"no\" \"no\"],  \t\r\n  \"key\"  \t\r\n  :  \t\r\n  10,  \t\r\n   \"no3\":\"no\" }"},
		{10, []string{"k2", "k1", "key"}, "{  \t\r\n  \"k2\"  \t\r\n  :  \t\r\n  {  \t\r\n  \"k1\"  \t\r\n  :  \t\r\n  {  \t\r\n  \"key\"  \t\r\n  :  \t\r\n  10  \t\r\n  }  \t\r\n  }  \t\r\n  }"},
		{10, []string{"k2", "k1", "key"}, "{ \"no1\":\"no\" \"no2\":[\"no\" \"no\"],  \t\r\n  \"k2\"  \t\r\n  :  \t\r\n  { \"no3\":\"no\",  \t\r\n  \"no4\":[\"no\" \"no\"], \"k1\"  \t\r\n  :  \t\r\n  { \"no5\":\"no\",  \"no6\":[\"no\" \"no\"],  \t\r\n  \"key\"  \t\r\n  :  \t\r\n  10  \t\r\n  }  \t\r\n  }  \t\r\n  }"},
		{1000, []string{"key"}, "{  \t\r\n  \"key\"  \t\r\n  :  \t\r\n  1000  \t\r\n  }"},
		{1000, []string{"key"}, "{ \"no1\":\"no\", \"no2\":[\"no\" \"no\"],  \t\r\n  \"key\"  \t\r\n  :  \t\r\n  1000,  \t\r\n   \"no3\":\"no\" }"},
		{1000, []string{"k2", "k1", "key"}, "{  \t\r\n  \"k2\"  \t\r\n  :  \t\r\n  {  \t\r\n  \"k1\"  \t\r\n  :  \t\r\n  {  \t\r\n  \"key\"  \t\r\n  :  \t\r\n  1000  \t\r\n  }  \t\r\n  }  \t\r\n  }"},
		{1000, []string{"k2", "k1", "key"}, "{ \"no1\":\"no\" \"no2\":[\"no\" \"no\"],  \t\r\n  \"k2\"  \t\r\n  :  \t\r\n  { \"no3\":\"no\",  \t\r\n  \"no4\":[\"no\" \"no\"], \"k1\"  \t\r\n  :  \t\r\n  { \"no5\":\"no\",  \"no6\":[\"no\" \"no\"],  \t\r\n  \"key\"  \t\r\n  :  \t\r\n  1000  \t\r\n  }  \t\r\n  }  \t\r\n  }"},
		{-10, []string{"key"}, "{  \t\r\n  \"key\"  \t\r\n  :  \t\r\n  -10  \t\r\n  }"},
		{-10, []string{"key"}, "{ \"no1\":\"no\", \"no2\":[\"no\" \"no\"],  \t\r\n  \"key\"  \t\r\n  :  \t\r\n  -10,  \t\r\n   \"no3\":\"no\" }"},
		{-10, []string{"k2", "k1", "key"}, "{  \t\r\n  \"k2\"  \t\r\n  :  \t\r\n  {  \t\r\n  \"k1\"  \t\r\n  :  \t\r\n  {  \t\r\n  \"key\"  \t\r\n  :  \t\r\n  -10  \t\r\n  }  \t\r\n  }  \t\r\n  }"},
		{-10, []string{"k2", "k1", "key"}, "{ \"no1\":\"no\" \"no2\":[\"no\" \"no\"],  \t\r\n  \"k2\"  \t\r\n  :  \t\r\n  { \"no3\":\"no\",  \t\r\n  \"no4\":[\"no\" \"no\"], \"k1\"  \t\r\n  :  \t\r\n  { \"no5\":\"no\",  \"no6\":[\"no\" \"no\"],  \t\r\n  \"key\"  \t\r\n  :  \t\r\n  -10  \t\r\n  }  \t\r\n  }  \t\r\n  }"},
		{-1000, []string{"key"}, "{  \t\r\n  \"key\"  \t\r\n  :  \t\r\n  -1000  \t\r\n  }"},
		{-1000, []string{"key"}, "{ \"no1\":\"no\", \"no2\":[\"no\" \"no\"],  \t\r\n  \"key\"  \t\r\n  :  \t\r\n  -1000,  \t\r\n   \"no3\":\"no\" }"},
		{-1000, []string{"k2", "k1", "key"}, "{  \t\r\n  \"k2\"  \t\r\n  :  \t\r\n  {  \t\r\n  \"k1\"  \t\r\n  :  \t\r\n  {  \t\r\n  \"key\"  \t\r\n  :  \t\r\n  -1000  \t\r\n  }  \t\r\n  }  \t\r\n  }"},
		{-1000, []string{"k2", "k1", "key"}, "{ \"no1\":\"no\" \"no2\":[\"no\" \"no\"],  \t\r\n  \"k2\"  \t\r\n  :  \t\r\n  { \"no3\":\"no\",  \t\r\n  \"no4\":[\"no\" \"no\"], \"k1\"  \t\r\n  :  \t\r\n  { \"no5\":\"no\",  \"no6\":[\"no\" \"no\"],  \t\r\n  \"key\"  \t\r\n  :  \t\r\n  -1000  \t\r\n  }  \t\r\n  }  \t\r\n  }"},
	}
	for _, cse := range cses {
		t.Run("Int", func(t *testing.T) {
			var j jsn.Jsnr
			j.Reset(cse.txt)
			a := j.Int(cse.pth...)
			tst.IntEql(t, cse.e, a)
		})
	}
}
