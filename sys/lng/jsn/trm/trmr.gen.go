package trm

import (
	"sys/bsc/bnd"
	"sys/lng/scn"
	"unicode"
)

type (
	Trmr struct {
		scn.Scnr
	}
	StrLit struct {
		bnd.Bnd
	}
	BolLit struct {
		bnd.Bnd
	}
	FltLit struct {
		bnd.Bnd
	}
	IntLit struct {
		bnd.Bnd
	}
)

func (x *Trmr) SkpSpce() {
	for unicode.IsSpace(x.Ch) {
		x.NextRune()
	}
}
func (x *Trmr) StrLit() (r StrLit, ok bool) {
	r.Idx = x.Idx
	if x.Ch != '"' {
		return r, false
	}
	x.NextRune()
	for !x.End && x.Ch != '"' {
		if x.Ch == '\\' { // escape sequence: start
			x.NextRune()
			switch x.Ch {
			case 'a', 'b', 'f', 'n', 'r', 't', 'v', '\\', '"':
			default:
				return r, false // unknown escape sequence
			}
		}
		x.NextRune()
	}
	if x.End || x.Ch != '"' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) BolLit() (r BolLit, ok bool) {
	scn := x.Scn
	falseLit, ok := x.FalseLit()
	if ok {
		return BolLit{Bnd: falseLit}, true
	}
	x.Scn = scn
	trueLit, ok := x.TrueLit()
	if ok {
		return BolLit{Bnd: trueLit}, true
	}
	x.Scn = scn
	return r, false
}
func (x *Trmr) FalseLit() (r bnd.Bnd, ok bool) {
	if x.Ch != 'f' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) TrueLit() (r bnd.Bnd, ok bool) {
	if x.Ch != 't' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'u' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) FltLit() (r FltLit, ok bool) {
	// TODO: NaN, +Inf, -Inf
	r.Idx = x.Idx
	if x.Ch == '-' { // optional minus
		x.NextRune()
		if !unicode.IsDigit(x.Ch) { // next ch must be digit
			return r, false
		}
		x.NextRune()
	}
	for !x.End && unicode.IsDigit(x.Ch) {
		x.NextRune()
	}
	if x.Ch != '.' {
		return r, false
	}
	x.NextRune()
	if !unicode.IsDigit(x.Ch) { // next ch must be digit
		return r, false
	}
	for !x.End && unicode.IsDigit(x.Ch) {
		x.NextRune()
	}
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, r.Idx != r.Lim
}
func (x *Trmr) IntLit() (r IntLit, ok bool) {
	r.Idx = x.Idx
	if x.Ch == '-' { // optional minus
		x.NextRune()
		if !unicode.IsDigit(x.Ch) { // ch must be digit
			return r, false
		}
		x.NextRune()
	}
	for !x.End && unicode.IsDigit(x.Ch) {
		x.NextRune()
	}
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, r.Idx != r.Lim
}
