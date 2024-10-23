package prs

import (
	"strconv"
	"sys/bsc/bnd"
	"sys/bsc/bnds"
	"sys/bsc/bol"
	"sys/bsc/bols"
	"sys/bsc/flt"
	"sys/bsc/flts"
	bscint "sys/bsc/int"
	"sys/bsc/ints"
	"sys/bsc/str"
	"sys/bsc/strs"
	"sys/bsc/tme"
	"sys/bsc/tmes"
	"sys/bsc/unt"
	"sys/bsc/unts"
	"sys/err"
	"sys/k"
	"sys/lng/pro/trm"
	"time"
)

func StrTrm(trm trm.StrLit, txt string) (r str.Str) { return str.Str(txt[trm.Idx+1 : trm.Lim-1]) }
func Str(txt string) (r str.Str) {
	var trmr trm.Trmr
	trmr.Reset(txt)
	strLit, ok := trmr.StrLit()
	if !ok {
		err.Panicf("Str: failed to parse (txt:%q)", txt)
	}
	return StrTrm(strLit, txt)
}
func BolTrm(trm trm.BolLit, txt string) (r bol.Bol) { return txt[trm.Idx:trm.Lim] == k.Tru }
func Bol(txt string) (r bol.Bol) {
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
func Flt(txt string) (r flt.Flt) {
	var trmr trm.Trmr
	trmr.Reset(txt)
	fltLit, ok := trmr.FltLit()
	if !ok {
		err.Panicf("Flt: failed to parse (txt:%q)", txt)
	}
	return FltTrm(fltLit, txt)
}
func UntTrm(trm trm.UntLit, txt string) (r unt.Unt) {
	lit := txt[trm.Idx:trm.Lim]
	mag := unt.Unt(1)
	for n := len(lit) - 1; n > -1; n-- {
		r += mag * unt.Unt(lit[n]-'0')
		mag *= unt.Unt(10)
	}
	return r
}
func Unt(txt string) (r unt.Unt) {
	var trmr trm.Trmr
	trmr.Reset(txt)
	untLit, ok := trmr.UntLit()
	if !ok {
		err.Panicf("Unt: failed to parse (txt:%q)", txt)
	}
	return UntTrm(untLit, txt)
}
func IntTrm(trm trm.IntLit, txt string) (r bscint.Int) {
	lit := txt[trm.Idx+1 : trm.Lim]
	mag := bscint.Int(1)
	for n := len(lit) - 1; n > -1; n-- {
		r += mag * bscint.Int(lit[n]-'0')
		mag *= bscint.Int(10)
	}
	if txt[trm.Idx:trm.Idx+1] == "-" {
		r = -r
	}
	return r
}
func Int(txt string) (r bscint.Int) {
	var trmr trm.Trmr
	trmr.Reset(txt)
	intLit, ok := trmr.IntLit()
	if !ok {
		err.Panicf("Int: failed to parse (txt:%q)", txt)
	}
	return IntTrm(intLit, txt)
}
func TmePrtTrm(trm trm.TmePrtLit, txt string) (r tme.Tme) {
	if trm.Lim == 0 {
		return 0
	}
	lit := txt[trm.Idx : trm.Lim-1] // trim suffix ch: y,n,w,d,h,m,s
	neg := lit[0] == '-'            // optional minus
	if neg {
		lit = lit[1:]
	}
	mag := tme.Tme(1)
	for n := len(lit) - 1; n > -1; n-- {
		r += mag * tme.Tme(lit[n]-'0')
		mag *= tme.Tme(10)
	}
	if neg {
		r = -r
	}
	return r
}
func TmeTrm(trm trm.TmeLit, txt string) (r tme.Tme) {
	year := TmePrtTrm(trm.Year, txt)
	if year != 0 {
		month := TmePrtTrm(trm.Month, txt)
		if month == 0 {
			month = 1
		}
		day := TmePrtTrm(trm.Day, txt)
		if day == 0 {
			day = 1
		}
		return tme.Time(time.Date(int(year), time.Month(month), int(day), 0, 0, 0, 0, time.UTC)) +
			TmePrtTrm(trm.Hour, txt)*tme.Hour +
			TmePrtTrm(trm.Minute, txt)*tme.Minute +
			TmePrtTrm(trm.Second, txt)*tme.Second
	}
	week := TmePrtTrm(trm.Week, txt) * tme.Week
	day := TmePrtTrm(trm.Day, txt) * tme.Day
	hour := TmePrtTrm(trm.Hour, txt) * tme.Hour
	minute := TmePrtTrm(trm.Minute, txt) * tme.Minute
	second := TmePrtTrm(trm.Second, txt) * tme.Second
	switch {
	case week < 0:
		day = -day
		hour = -hour
		minute = -minute
		second = -second
	case day < 0:
		hour = -hour
		minute = -minute
		second = -second
	case hour < 0:
		minute = -minute
		second = -second
	case minute < 0:
		second = -second
	}
	return week + day + hour + minute + second
}
func Tme(txt string) (r tme.Tme) {
	var trmr trm.Trmr
	trmr.Reset(txt)
	tmeLit, ok := trmr.TmeLit()
	if !ok {
		err.Panicf("Tme: failed to parse (txt:%q)", txt)
	}
	return TmeTrm(tmeLit, txt)
}
func BndTrm(trm trm.BndLit, txt string) (r bnd.Bnd) {
	r.Idx = UntTrm(trm.IdxTrm, txt)
	r.Lim = UntTrm(trm.LimTrm, txt)
	return r
}
func Bnd(txt string) (r bnd.Bnd) {
	var trmr trm.Trmr
	trmr.Reset(txt)
	bndLit, ok := trmr.BndLit()
	if !ok {
		err.Panicf("Bnd: failed to parse (txt:%q)", txt)
	}
	return BndTrm(bndLit, txt)
}
func FltRngTrm(trm trm.FltRngLit, txt string) (r flt.Rng) {
	r.Min = FltTrm(trm.MinTrm, txt)
	r.Max = FltTrm(trm.MaxTrm, txt)
	return r
}
func FltRng(txt string) (r flt.Rng) {
	var trmr trm.Trmr
	trmr.Reset(txt)
	fltRngLit, ok := trmr.FltRngLit()
	if !ok {
		err.Panicf("Rng: failed to parse (txt:%q)", txt)
	}
	return FltRngTrm(fltRngLit, txt)
}
func TmeRngTrm(trm trm.TmeRngLit, txt string) (r tme.Rng) {
	r.Min = TmeTrm(trm.MinTrm, txt)
	r.Max = TmeTrm(trm.MaxTrm, txt)
	return r
}
func TmeRng(txt string) (r tme.Rng) {
	var trmr trm.Trmr
	trmr.Reset(txt)
	tmeRngLit, ok := trmr.TmeRngLit()
	if !ok {
		err.Panicf("Rng: failed to parse (txt:%q)", txt)
	}
	return TmeRngTrm(tmeRngLit, txt)
}
func StrsTrm(trm trm.StrsLit, txt string) *strs.Strs {
	r := strs.Strs{}
	for _, elm := range trm.Elms {
		r.Push(StrTrm(elm, txt))
	}
	return &r
}
func Strs(txt string) (r *strs.Strs) {
	var trmr trm.Trmr
	trmr.Reset(txt)
	strsLit, ok := trmr.StrsLit()
	if !ok {
		err.Panicf("Strs: failed to parse (txt:%q)", txt)
	}
	return StrsTrm(strsLit, txt)
}
func BolsTrm(trm trm.BolsLit, txt string) *bols.Bols {
	r := bols.Bols{}
	for _, elm := range trm.Elms {
		r.Push(BolTrm(elm, txt))
	}
	return &r
}
func Bols(txt string) (r *bols.Bols) {
	var trmr trm.Trmr
	trmr.Reset(txt)
	bolsLit, ok := trmr.BolsLit()
	if !ok {
		err.Panicf("Bols: failed to parse (txt:%q)", txt)
	}
	return BolsTrm(bolsLit, txt)
}
func FltsTrm(trm trm.FltsLit, txt string) *flts.Flts {
	r := flts.Flts{}
	for _, elm := range trm.Elms {
		r.Push(FltTrm(elm, txt))
	}
	return &r
}
func Flts(txt string) (r *flts.Flts) {
	var trmr trm.Trmr
	trmr.Reset(txt)
	fltsLit, ok := trmr.FltsLit()
	if !ok {
		err.Panicf("Flts: failed to parse (txt:%q)", txt)
	}
	return FltsTrm(fltsLit, txt)
}
func UntsTrm(trm trm.UntsLit, txt string) *unts.Unts {
	r := unts.Unts{}
	for _, elm := range trm.Elms {
		r.Push(UntTrm(elm, txt))
	}
	return &r
}
func Unts(txt string) (r *unts.Unts) {
	var trmr trm.Trmr
	trmr.Reset(txt)
	untsLit, ok := trmr.UntsLit()
	if !ok {
		err.Panicf("Unts: failed to parse (txt:%q)", txt)
	}
	return UntsTrm(untsLit, txt)
}
func IntsTrm(trm trm.IntsLit, txt string) *ints.Ints {
	r := ints.Ints{}
	for _, elm := range trm.Elms {
		r.Push(IntTrm(elm, txt))
	}
	return &r
}
func Ints(txt string) (r *ints.Ints) {
	var trmr trm.Trmr
	trmr.Reset(txt)
	intsLit, ok := trmr.IntsLit()
	if !ok {
		err.Panicf("Ints: failed to parse (txt:%q)", txt)
	}
	return IntsTrm(intsLit, txt)
}
func TmesTrm(trm trm.TmesLit, txt string) *tmes.Tmes {
	r := tmes.Tmes{}
	for _, elm := range trm.Elms {
		r.Push(TmeTrm(elm, txt))
	}
	return &r
}
func Tmes(txt string) (r *tmes.Tmes) {
	var trmr trm.Trmr
	trmr.Reset(txt)
	tmesLit, ok := trmr.TmesLit()
	if !ok {
		err.Panicf("Tmes: failed to parse (txt:%q)", txt)
	}
	return TmesTrm(tmesLit, txt)
}
func BndsTrm(trm trm.BndsLit, txt string) *bnds.Bnds {
	r := bnds.Bnds{}
	for _, elm := range trm.Elms {
		r.Push(BndTrm(elm, txt))
	}
	return &r
}
func Bnds(txt string) (r *bnds.Bnds) {
	var trmr trm.Trmr
	trmr.Reset(txt)
	bndsLit, ok := trmr.BndsLit()
	if !ok {
		err.Panicf("Bnds: failed to parse (txt:%q)", txt)
	}
	return BndsTrm(bndsLit, txt)
}
func TmeRngsTrm(trm trm.TmeRngsLit, txt string) *tme.Rngs {
	r := tme.Rngs{}
	for _, elm := range trm.Elms {
		r.Push(TmeRngTrm(elm, txt))
	}
	return &r
}
func TmeRngs(txt string) (r *tme.Rngs) {
	var trmr trm.Trmr
	trmr.Reset(txt)
	tmeRngsLit, ok := trmr.TmeRngsLit()
	if !ok {
		err.Panicf("Rngs: failed to parse (txt:%q)", txt)
	}
	return TmeRngsTrm(tmeRngsLit, txt)
}
