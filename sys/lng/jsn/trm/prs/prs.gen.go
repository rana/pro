package prs

import (
	"strconv"
	"sys/bsc/bol"
	"sys/bsc/flt"
	"sys/bsc/int"
	"sys/bsc/str"
	"sys/err"
	"sys/k"
	"sys/lng/jsn/trm"
)

func StrTrm(trm trm.StrLit, txt string) (r str.Str) { return str.Str(txt[trm.Idx+1 : trm.Lim-1]) }
func StrTxt(txt string) (r str.Str) {
	var trmr trm.Trmr
	trmr.Reset(txt)
	strLit, ok := trmr.StrLit()
	if !ok {
		err.Panicf("Str: failed to parse (txt:%q)", txt)
	}
	return StrTrm(strLit, txt)
}
func BolTrm(trm trm.BolLit, txt string) (r bol.Bol) { return txt[trm.Idx:trm.Lim] == k.True }
func BolTxt(txt string) (r bol.Bol) {
	var trmr trm.Trmr
	trmr.Reset(txt)
	bolLit, ok := trmr.BolLit()
	if !ok {
		err.Panicf("Bol: failed to parse (txt:%q)", txt)
	}
	return BolTrm(bolLit, txt)
}
func FltTrm(trm trm.FltLit, txt string) (r flt.Flt) {
	v, er := strconv.ParseFloat(txt[trm.Idx:trm.Lim], 32)
	if er != nil {
		err.Panicf("Flt: failed to parse (txt:%q err:%q)", txt, er)
	}
	return flt.Flt(v)
}
func FltTxt(txt string) (r flt.Flt) {
	var trmr trm.Trmr
	trmr.Reset(txt)
	fltLit, ok := trmr.FltLit()
	if !ok {
		err.Panicf("Flt: failed to parse (txt:%q)", txt)
	}
	return FltTrm(fltLit, txt)
}
func IntTrm(trm trm.IntLit, txt string) (r int.Int) {
	var lit string
	if txt[trm.Idx:trm.Idx+1] == "-" {
		lit = txt[trm.Idx+1 : trm.Lim]
	} else {
		lit = txt[trm.Idx:trm.Lim]
	}
	mag := int.Int(1)
	for n := len(lit) - 1; n > -1; n-- {
		r += mag * int.Int(lit[n]-'0')
		mag *= int.Int(10)
	}
	if txt[trm.Idx:trm.Idx+1] == "-" {
		r = -r
	}
	return r
}
func IntTxt(txt string) (r int.Int) {
	var trmr trm.Trmr
	trmr.Reset(txt)
	intLit, ok := trmr.IntLit()
	if !ok {
		err.Panicf("Int: failed to parse (txt:%q)", txt)
	}
	return IntTrm(intLit, txt)
}
