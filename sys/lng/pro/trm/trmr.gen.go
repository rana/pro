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
	SpceLit struct {
		bnd.Bnd
	}
	CmntLit struct {
		bnd.Bnd
	}
	IdnLit struct {
		bnd.Bnd
	}
	ObjsLit struct {
		bnd.Bnd
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
	UntLit struct {
		bnd.Bnd
	}
	IntLit struct {
		bnd.Bnd
	}
	TmePrtLit struct {
		bnd.Bnd
	}
	TmeLit struct {
		bnd.Bnd
		Year   TmePrtLit
		Month  TmePrtLit
		Week   TmePrtLit
		Day    TmePrtLit
		Hour   TmePrtLit
		Minute TmePrtLit
		Second TmePrtLit
	}
	BndLit struct {
		bnd.Bnd
		IdxTrm UntLit
		LimTrm UntLit
	}
	FltRngLit struct {
		bnd.Bnd
		MinTrm FltLit
		MaxTrm FltLit
	}
	TmeRngLit struct {
		bnd.Bnd
		MinTrm TmeLit
		MaxTrm TmeLit
	}
	StrsLit struct {
		bnd.Bnd
		Elms []StrLit
	}
	BolsLit struct {
		bnd.Bnd
		Elms []BolLit
	}
	FltsLit struct {
		bnd.Bnd
		Elms []FltLit
	}
	UntsLit struct {
		bnd.Bnd
		Elms []UntLit
	}
	IntsLit struct {
		bnd.Bnd
		Elms []IntLit
	}
	TmesLit struct {
		bnd.Bnd
		Elms []TmeLit
	}
	BndsLit struct {
		bnd.Bnd
		Elms []BndLit
	}
	TmeRngsLit struct {
		bnd.Bnd
		Elms []TmeRngLit
	}
)

func (x *Trmr) SpceLit() (r SpceLit, ok bool) {
	if !unicode.IsSpace(x.Ch) {
		return r, false
	}
	r.Idx = x.Idx
	for !x.NextRune() && unicode.IsSpace(x.Ch) {
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) CmntLit() (r CmntLit, ok bool) {
	if x.Ch != '/' || (x.Ch == '/' && x.PeekRune() != '/') {
		return r, false
	}
	r.Idx = x.Idx
	for !x.NextRune() && x.Ch != '\n' {
	}
	if x.Ch == '\n' {
		x.NextRune()
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SkpSpceCmnt() {
	// skip all consecutive spce and cmnt
	scn, spceOk, cmntOk := x.Scn, true, true
	for spceOk || cmntOk {
		_, spceOk = x.SpceLit()
		if spceOk {
			scn = x.Scn
		} else {
			x.Scn = scn
		}
		_, cmntOk = x.CmntLit()
		if cmntOk {
			scn = x.Scn
		} else {
			x.Scn = scn
		}
	}
}
func (x *Trmr) IdnLit() (r IdnLit, ok bool) {
	if !(unicode.IsLetter(x.Ch) || x.Ch == '_') {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	for unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		x.NextRune()
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) ObjsLit() (r ObjsLit, ok bool) {
	if x.Ch != '[' {
		return r, false
	}
	r.Idx = x.Idx
	x.SkpSet('[', ']')
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Prefixs(idn string) (r []bnd.Bnd) {
	for !x.End {
		x.SkpSpceCmnt()
		if x.Ch != rune(idn[0]) {
			x.NextRune()
			continue
		}
		curBnd := bnd.Bnd{Idx: x.Idx, Lim: x.Idx + 1}
		for n := 1; n < len(idn); n++ {
			x.NextRune()
			if x.Ch == rune(idn[n]) {
				curBnd.Lim++
				continue
			}
			break
		}
		if int(curBnd.Lim-curBnd.Idx) == len(idn) {
			x.NextRune()
			x.SkpSpceCmnt()
			if x.Ch == '.' {
				x.NextRune()
				x.SkpSpceCmnt()
				_, ok := x.IdnLit()
				if ok {
					x.SkpSpceCmnt()
					x.SkpSet('(', ')')
					curBnd.Lim = x.Idx
					r = append(r, curBnd)
				}
			}
		}
	}
	return r
}
func (x *Trmr) Asn() (r bnd.Bnd, ok bool) {
	if x.Ch != 'a' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Each() (r bnd.Bnd, ok bool) {
	if x.Ch != 'e' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'c' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'h' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) PllEach() (r bnd.Bnd, ok bool) {
	if x.Ch != 'p' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'E' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'c' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'h' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) PllWait() (r bnd.Bnd, ok bool) {
	if x.Ch != 'p' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'W' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Then() (r bnd.Bnd, ok bool) {
	if x.Ch != 't' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'h' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Else() (r bnd.Bnd, ok bool) {
	if x.Ch != 'e' {
		return r, false
	}
	r.Idx = x.Idx
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
func (x *Trmr) PnlPct() (r bnd.Bnd, ok bool) {
	if x.Ch != 'p' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'c' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) ScsPct() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'c' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'c' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) PipPerDay() (r bnd.Bnd, ok bool) {
	if x.Ch != 'p' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'D' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'y' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) UsdPerDay() (r bnd.Bnd, ok bool) {
	if x.Ch != 'u' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'D' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'y' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) ScsPerDay() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'c' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'D' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'y' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) OpnPerDay() (r bnd.Bnd, ok bool) {
	if x.Ch != 'o' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'D' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'y' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) PnlUsd() (r bnd.Bnd, ok bool) {
	if x.Ch != 'p' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'U' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) PipAvg() (r bnd.Bnd, ok bool) {
	if x.Ch != 'p' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'v' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) PipMdn() (r bnd.Bnd, ok bool) {
	if x.Ch != 'p' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'M' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) PipMin() (r bnd.Bnd, ok bool) {
	if x.Ch != 'p' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'M' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) PipMax() (r bnd.Bnd, ok bool) {
	if x.Ch != 'p' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'M' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'x' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) PipSum() (r bnd.Bnd, ok bool) {
	if x.Ch != 'p' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'S' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'u' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) DurAvg() (r bnd.Bnd, ok bool) {
	if x.Ch != 'd' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'u' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'v' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) DurMdn() (r bnd.Bnd, ok bool) {
	if x.Ch != 'd' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'u' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'M' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) DurMin() (r bnd.Bnd, ok bool) {
	if x.Ch != 'd' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'u' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'M' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) DurMax() (r bnd.Bnd, ok bool) {
	if x.Ch != 'd' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'u' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'M' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'x' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) LosLimMax() (r bnd.Bnd, ok bool) {
	if x.Ch != 'l' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'o' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'L' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'M' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'x' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) DurLimMax() (r bnd.Bnd, ok bool) {
	if x.Ch != 'd' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'u' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'L' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'M' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'x' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) DayCnt() (r bnd.Bnd, ok bool) {
	if x.Ch != 'd' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'y' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'C' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) TrdCnt() (r bnd.Bnd, ok bool) {
	if x.Ch != 't' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'C' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) TrdPct() (r bnd.Bnd, ok bool) {
	if x.Ch != 't' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'c' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) CstTotUsd() (r bnd.Bnd, ok bool) {
	if x.Ch != 'c' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'T' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'o' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'U' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) CstSpdUsd() (r bnd.Bnd, ok bool) {
	if x.Ch != 'c' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'S' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'U' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) CstComUsd() (r bnd.Bnd, ok bool) {
	if x.Ch != 'c' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'C' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'o' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'U' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) PnlPctA() (r bnd.Bnd, ok bool) {
	if x.Ch != 'p' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'c' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) PnlPctB() (r bnd.Bnd, ok bool) {
	if x.Ch != 'p' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'c' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'B' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) PnlPctDlt() (r bnd.Bnd, ok bool) {
	if x.Ch != 'p' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'c' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'D' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) ScsPctA() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'c' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'c' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) ScsPctB() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'c' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'c' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'B' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) ScsPctDlt() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'c' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'c' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'D' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) PipPerDayA() (r bnd.Bnd, ok bool) {
	if x.Ch != 'p' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'D' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'y' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) PipPerDayB() (r bnd.Bnd, ok bool) {
	if x.Ch != 'p' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'D' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'y' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'B' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) PipPerDayDlt() (r bnd.Bnd, ok bool) {
	if x.Ch != 'p' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'D' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'y' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'D' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) UsdPerDayA() (r bnd.Bnd, ok bool) {
	if x.Ch != 'u' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'D' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'y' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) UsdPerDayB() (r bnd.Bnd, ok bool) {
	if x.Ch != 'u' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'D' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'y' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'B' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) UsdPerDayDlt() (r bnd.Bnd, ok bool) {
	if x.Ch != 'u' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'D' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'y' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'D' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) ScsPerDayA() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'c' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'D' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'y' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) ScsPerDayB() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'c' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'D' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'y' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'B' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) ScsPerDayDlt() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'c' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'D' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'y' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'D' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) OpnPerDayA() (r bnd.Bnd, ok bool) {
	if x.Ch != 'o' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'D' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'y' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) OpnPerDayB() (r bnd.Bnd, ok bool) {
	if x.Ch != 'o' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'D' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'y' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'B' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) OpnPerDayDlt() (r bnd.Bnd, ok bool) {
	if x.Ch != 'o' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'D' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'y' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'D' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) PnlUsdA() (r bnd.Bnd, ok bool) {
	if x.Ch != 'p' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'U' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) PnlUsdB() (r bnd.Bnd, ok bool) {
	if x.Ch != 'p' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'U' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'B' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) PnlUsdDlt() (r bnd.Bnd, ok bool) {
	if x.Ch != 'p' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'U' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'D' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) PipAvgA() (r bnd.Bnd, ok bool) {
	if x.Ch != 'p' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'v' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) PipAvgB() (r bnd.Bnd, ok bool) {
	if x.Ch != 'p' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'v' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'B' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) PipAvgDlt() (r bnd.Bnd, ok bool) {
	if x.Ch != 'p' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'v' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'D' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) PipMdnA() (r bnd.Bnd, ok bool) {
	if x.Ch != 'p' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'M' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) PipMdnB() (r bnd.Bnd, ok bool) {
	if x.Ch != 'p' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'M' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'B' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) PipMdnDlt() (r bnd.Bnd, ok bool) {
	if x.Ch != 'p' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'M' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'D' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) PipMinA() (r bnd.Bnd, ok bool) {
	if x.Ch != 'p' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'M' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) PipMinB() (r bnd.Bnd, ok bool) {
	if x.Ch != 'p' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'M' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'B' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) PipMinDlt() (r bnd.Bnd, ok bool) {
	if x.Ch != 'p' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'M' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'D' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) PipMaxA() (r bnd.Bnd, ok bool) {
	if x.Ch != 'p' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'M' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'x' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) PipMaxB() (r bnd.Bnd, ok bool) {
	if x.Ch != 'p' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'M' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'x' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'B' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) PipMaxDlt() (r bnd.Bnd, ok bool) {
	if x.Ch != 'p' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'M' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'x' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'D' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) PipSumA() (r bnd.Bnd, ok bool) {
	if x.Ch != 'p' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'S' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'u' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) PipSumB() (r bnd.Bnd, ok bool) {
	if x.Ch != 'p' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'S' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'u' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'B' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) PipSumDlt() (r bnd.Bnd, ok bool) {
	if x.Ch != 'p' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'S' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'u' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'D' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) DurAvgA() (r bnd.Bnd, ok bool) {
	if x.Ch != 'd' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'u' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'v' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) DurAvgB() (r bnd.Bnd, ok bool) {
	if x.Ch != 'd' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'u' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'v' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'B' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) DurAvgDlt() (r bnd.Bnd, ok bool) {
	if x.Ch != 'd' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'u' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'v' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'D' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) DurMdnA() (r bnd.Bnd, ok bool) {
	if x.Ch != 'd' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'u' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'M' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) DurMdnB() (r bnd.Bnd, ok bool) {
	if x.Ch != 'd' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'u' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'M' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'B' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) DurMdnDlt() (r bnd.Bnd, ok bool) {
	if x.Ch != 'd' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'u' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'M' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'D' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) DurMinA() (r bnd.Bnd, ok bool) {
	if x.Ch != 'd' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'u' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'M' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) DurMinB() (r bnd.Bnd, ok bool) {
	if x.Ch != 'd' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'u' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'M' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'B' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) DurMinDlt() (r bnd.Bnd, ok bool) {
	if x.Ch != 'd' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'u' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'M' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'D' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) DurMaxA() (r bnd.Bnd, ok bool) {
	if x.Ch != 'd' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'u' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'M' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'x' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) DurMaxB() (r bnd.Bnd, ok bool) {
	if x.Ch != 'd' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'u' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'M' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'x' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'B' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) DurMaxDlt() (r bnd.Bnd, ok bool) {
	if x.Ch != 'd' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'u' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'M' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'x' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'D' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) TrdCntA() (r bnd.Bnd, ok bool) {
	if x.Ch != 't' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'C' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) TrdCntB() (r bnd.Bnd, ok bool) {
	if x.Ch != 't' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'C' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'B' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) TrdCntDlt() (r bnd.Bnd, ok bool) {
	if x.Ch != 't' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'C' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'D' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) TrdPctA() (r bnd.Bnd, ok bool) {
	if x.Ch != 't' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'c' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) TrdPctB() (r bnd.Bnd, ok bool) {
	if x.Ch != 't' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'c' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'B' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) TrdPctDlt() (r bnd.Bnd, ok bool) {
	if x.Ch != 't' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'c' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'D' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) PthB() (r bnd.Bnd, ok bool) {
	if x.Ch != 'p' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'h' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'B' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Clr() (r bnd.Bnd, ok bool) {
	if x.Ch != 'c' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Wid() (r bnd.Bnd, ok bool) {
	if x.Ch != 'w' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Min() (r bnd.Bnd, ok bool) {
	if x.Ch != 'm' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Max() (r bnd.Bnd, ok bool) {
	if x.Ch != 'm' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'x' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) EqiDst() (r bnd.Bnd, ok bool) {
	if x.Ch != 'e' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'D' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Title() (r bnd.Bnd, ok bool) {
	if x.Ch != 't' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
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
func (x *Trmr) Y() (r bnd.Bnd, ok bool) {
	if x.Ch != 'y' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Outlier() (r bnd.Bnd, ok bool) {
	if x.Ch != 'o' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'u' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Plts() (r bnd.Bnd, ok bool) {
	if x.Ch != 'p' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Zero() (r bnd.Bnd, ok bool) {
	if x.Ch != 'z' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'o' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Empty() (r bnd.Bnd, ok bool) {
	if x.Ch != 'e' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'y' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Fls() (r bnd.Bnd, ok bool) {
	if x.Ch != 'f' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Tru() (r bnd.Bnd, ok bool) {
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
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) One() (r bnd.Bnd, ok bool) {
	if x.Ch != 'o' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'n' {
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
func (x *Trmr) NegOne() (r bnd.Bnd, ok bool) {
	if x.Ch != 'n' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'O' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
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
func (x *Trmr) Hndrd() (r bnd.Bnd, ok bool) {
	if x.Ch != 'h' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Tiny() (r bnd.Bnd, ok bool) {
	if x.Ch != 't' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'y' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Second() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'c' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'o' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Minute() (r bnd.Bnd, ok bool) {
	if x.Ch != 'm' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'u' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
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
func (x *Trmr) Hour() (r bnd.Bnd, ok bool) {
	if x.Ch != 'h' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'o' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'u' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Day() (r bnd.Bnd, ok bool) {
	if x.Ch != 'd' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'y' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Week() (r bnd.Bnd, ok bool) {
	if x.Ch != 'w' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'k' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) S1() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != '1' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) S5() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != '5' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) S10() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != '1' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) S15() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != '1' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '5' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) S20() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != '2' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) S30() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != '3' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) S40() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != '4' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) S50() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != '5' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) M1() (r bnd.Bnd, ok bool) {
	if x.Ch != 'm' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != '1' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) M5() (r bnd.Bnd, ok bool) {
	if x.Ch != 'm' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != '5' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) M10() (r bnd.Bnd, ok bool) {
	if x.Ch != 'm' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != '1' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) M15() (r bnd.Bnd, ok bool) {
	if x.Ch != 'm' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != '1' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '5' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) M20() (r bnd.Bnd, ok bool) {
	if x.Ch != 'm' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != '2' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) M30() (r bnd.Bnd, ok bool) {
	if x.Ch != 'm' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != '3' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) M40() (r bnd.Bnd, ok bool) {
	if x.Ch != 'm' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != '4' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) M50() (r bnd.Bnd, ok bool) {
	if x.Ch != 'm' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != '5' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) H1() (r bnd.Bnd, ok bool) {
	if x.Ch != 'h' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != '1' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) D1() (r bnd.Bnd, ok bool) {
	if x.Ch != 'd' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != '1' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Resolution() (r bnd.Bnd, ok bool) {
	if x.Ch != 'r' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'o' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'u' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'o' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Black() (r bnd.Bnd, ok bool) {
	if x.Ch != 'b' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'c' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'k' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) White() (r bnd.Bnd, ok bool) {
	if x.Ch != 'w' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'h' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
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
func (x *Trmr) Red50() (r bnd.Bnd, ok bool) {
	if x.Ch != 'r' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '5' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Red100() (r bnd.Bnd, ok bool) {
	if x.Ch != 'r' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '1' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Red200() (r bnd.Bnd, ok bool) {
	if x.Ch != 'r' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '2' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Red300() (r bnd.Bnd, ok bool) {
	if x.Ch != 'r' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '3' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Red400() (r bnd.Bnd, ok bool) {
	if x.Ch != 'r' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '4' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Red500() (r bnd.Bnd, ok bool) {
	if x.Ch != 'r' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '5' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Red600() (r bnd.Bnd, ok bool) {
	if x.Ch != 'r' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '6' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Red700() (r bnd.Bnd, ok bool) {
	if x.Ch != 'r' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '7' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Red800() (r bnd.Bnd, ok bool) {
	if x.Ch != 'r' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '8' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Red900() (r bnd.Bnd, ok bool) {
	if x.Ch != 'r' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '9' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) RedA100() (r bnd.Bnd, ok bool) {
	if x.Ch != 'r' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '1' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) RedA200() (r bnd.Bnd, ok bool) {
	if x.Ch != 'r' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '2' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) RedA400() (r bnd.Bnd, ok bool) {
	if x.Ch != 'r' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '4' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) RedA700() (r bnd.Bnd, ok bool) {
	if x.Ch != 'r' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '7' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Pink50() (r bnd.Bnd, ok bool) {
	if x.Ch != 'p' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'k' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '5' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Pink100() (r bnd.Bnd, ok bool) {
	if x.Ch != 'p' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'k' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '1' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Pink200() (r bnd.Bnd, ok bool) {
	if x.Ch != 'p' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'k' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '2' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Pink300() (r bnd.Bnd, ok bool) {
	if x.Ch != 'p' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'k' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '3' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Pink400() (r bnd.Bnd, ok bool) {
	if x.Ch != 'p' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'k' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '4' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Pink500() (r bnd.Bnd, ok bool) {
	if x.Ch != 'p' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'k' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '5' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Pink600() (r bnd.Bnd, ok bool) {
	if x.Ch != 'p' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'k' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '6' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Pink700() (r bnd.Bnd, ok bool) {
	if x.Ch != 'p' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'k' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '7' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Pink800() (r bnd.Bnd, ok bool) {
	if x.Ch != 'p' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'k' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '8' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Pink900() (r bnd.Bnd, ok bool) {
	if x.Ch != 'p' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'k' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '9' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) PinkA100() (r bnd.Bnd, ok bool) {
	if x.Ch != 'p' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'k' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '1' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) PinkA200() (r bnd.Bnd, ok bool) {
	if x.Ch != 'p' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'k' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '2' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) PinkA400() (r bnd.Bnd, ok bool) {
	if x.Ch != 'p' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'k' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '4' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) PinkA700() (r bnd.Bnd, ok bool) {
	if x.Ch != 'p' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'k' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '7' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Purple50() (r bnd.Bnd, ok bool) {
	if x.Ch != 'p' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'u' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '5' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Purple100() (r bnd.Bnd, ok bool) {
	if x.Ch != 'p' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'u' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '1' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Purple200() (r bnd.Bnd, ok bool) {
	if x.Ch != 'p' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'u' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '2' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Purple300() (r bnd.Bnd, ok bool) {
	if x.Ch != 'p' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'u' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '3' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Purple400() (r bnd.Bnd, ok bool) {
	if x.Ch != 'p' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'u' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '4' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Purple500() (r bnd.Bnd, ok bool) {
	if x.Ch != 'p' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'u' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '5' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Purple600() (r bnd.Bnd, ok bool) {
	if x.Ch != 'p' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'u' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '6' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Purple700() (r bnd.Bnd, ok bool) {
	if x.Ch != 'p' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'u' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '7' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Purple800() (r bnd.Bnd, ok bool) {
	if x.Ch != 'p' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'u' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '8' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Purple900() (r bnd.Bnd, ok bool) {
	if x.Ch != 'p' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'u' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '9' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) PurpleA100() (r bnd.Bnd, ok bool) {
	if x.Ch != 'p' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'u' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '1' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) PurpleA200() (r bnd.Bnd, ok bool) {
	if x.Ch != 'p' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'u' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '2' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) PurpleA400() (r bnd.Bnd, ok bool) {
	if x.Ch != 'p' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'u' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '4' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) PurpleA700() (r bnd.Bnd, ok bool) {
	if x.Ch != 'p' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'u' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '7' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) DeepPurple50() (r bnd.Bnd, ok bool) {
	if x.Ch != 'd' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'u' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '5' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) DeepPurple100() (r bnd.Bnd, ok bool) {
	if x.Ch != 'd' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'u' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '1' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) DeepPurple200() (r bnd.Bnd, ok bool) {
	if x.Ch != 'd' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'u' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '2' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) DeepPurple300() (r bnd.Bnd, ok bool) {
	if x.Ch != 'd' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'u' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '3' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) DeepPurple400() (r bnd.Bnd, ok bool) {
	if x.Ch != 'd' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'u' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '4' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) DeepPurple500() (r bnd.Bnd, ok bool) {
	if x.Ch != 'd' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'u' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '5' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) DeepPurple600() (r bnd.Bnd, ok bool) {
	if x.Ch != 'd' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'u' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '6' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) DeepPurple700() (r bnd.Bnd, ok bool) {
	if x.Ch != 'd' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'u' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '7' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) DeepPurple800() (r bnd.Bnd, ok bool) {
	if x.Ch != 'd' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'u' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '8' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) DeepPurple900() (r bnd.Bnd, ok bool) {
	if x.Ch != 'd' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'u' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '9' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) DeepPurpleA100() (r bnd.Bnd, ok bool) {
	if x.Ch != 'd' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'u' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '1' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) DeepPurpleA200() (r bnd.Bnd, ok bool) {
	if x.Ch != 'd' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'u' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '2' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) DeepPurpleA400() (r bnd.Bnd, ok bool) {
	if x.Ch != 'd' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'u' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '4' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) DeepPurpleA700() (r bnd.Bnd, ok bool) {
	if x.Ch != 'd' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'u' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '7' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Indigo50() (r bnd.Bnd, ok bool) {
	if x.Ch != 'i' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'o' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '5' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Indigo100() (r bnd.Bnd, ok bool) {
	if x.Ch != 'i' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'o' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '1' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Indigo200() (r bnd.Bnd, ok bool) {
	if x.Ch != 'i' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'o' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '2' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Indigo300() (r bnd.Bnd, ok bool) {
	if x.Ch != 'i' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'o' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '3' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Indigo400() (r bnd.Bnd, ok bool) {
	if x.Ch != 'i' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'o' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '4' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Indigo500() (r bnd.Bnd, ok bool) {
	if x.Ch != 'i' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'o' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '5' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Indigo600() (r bnd.Bnd, ok bool) {
	if x.Ch != 'i' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'o' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '6' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Indigo700() (r bnd.Bnd, ok bool) {
	if x.Ch != 'i' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'o' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '7' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Indigo800() (r bnd.Bnd, ok bool) {
	if x.Ch != 'i' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'o' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '8' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Indigo900() (r bnd.Bnd, ok bool) {
	if x.Ch != 'i' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'o' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '9' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) IndigoA100() (r bnd.Bnd, ok bool) {
	if x.Ch != 'i' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'o' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '1' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) IndigoA200() (r bnd.Bnd, ok bool) {
	if x.Ch != 'i' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'o' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '2' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) IndigoA400() (r bnd.Bnd, ok bool) {
	if x.Ch != 'i' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'o' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '4' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) IndigoA700() (r bnd.Bnd, ok bool) {
	if x.Ch != 'i' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'o' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '7' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Blue50() (r bnd.Bnd, ok bool) {
	if x.Ch != 'b' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'l' {
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
	if x.Ch != '5' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Blue100() (r bnd.Bnd, ok bool) {
	if x.Ch != 'b' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'l' {
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
	if x.Ch != '1' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Blue200() (r bnd.Bnd, ok bool) {
	if x.Ch != 'b' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'l' {
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
	if x.Ch != '2' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Blue300() (r bnd.Bnd, ok bool) {
	if x.Ch != 'b' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'l' {
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
	if x.Ch != '3' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Blue400() (r bnd.Bnd, ok bool) {
	if x.Ch != 'b' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'l' {
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
	if x.Ch != '4' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Blue500() (r bnd.Bnd, ok bool) {
	if x.Ch != 'b' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'l' {
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
	if x.Ch != '5' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Blue600() (r bnd.Bnd, ok bool) {
	if x.Ch != 'b' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'l' {
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
	if x.Ch != '6' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Blue700() (r bnd.Bnd, ok bool) {
	if x.Ch != 'b' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'l' {
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
	if x.Ch != '7' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Blue800() (r bnd.Bnd, ok bool) {
	if x.Ch != 'b' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'l' {
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
	if x.Ch != '8' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Blue900() (r bnd.Bnd, ok bool) {
	if x.Ch != 'b' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'l' {
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
	if x.Ch != '9' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) BlueA100() (r bnd.Bnd, ok bool) {
	if x.Ch != 'b' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'l' {
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
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '1' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) BlueA200() (r bnd.Bnd, ok bool) {
	if x.Ch != 'b' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'l' {
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
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '2' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) BlueA400() (r bnd.Bnd, ok bool) {
	if x.Ch != 'b' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'l' {
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
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '4' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) BlueA700() (r bnd.Bnd, ok bool) {
	if x.Ch != 'b' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'l' {
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
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '7' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) LightBlue50() (r bnd.Bnd, ok bool) {
	if x.Ch != 'l' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'h' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'B' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
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
	if x.Ch != '5' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) LightBlue100() (r bnd.Bnd, ok bool) {
	if x.Ch != 'l' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'h' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'B' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
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
	if x.Ch != '1' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) LightBlue200() (r bnd.Bnd, ok bool) {
	if x.Ch != 'l' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'h' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'B' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
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
	if x.Ch != '2' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) LightBlue300() (r bnd.Bnd, ok bool) {
	if x.Ch != 'l' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'h' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'B' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
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
	if x.Ch != '3' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) LightBlue400() (r bnd.Bnd, ok bool) {
	if x.Ch != 'l' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'h' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'B' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
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
	if x.Ch != '4' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) LightBlue500() (r bnd.Bnd, ok bool) {
	if x.Ch != 'l' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'h' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'B' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
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
	if x.Ch != '5' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) LightBlue600() (r bnd.Bnd, ok bool) {
	if x.Ch != 'l' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'h' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'B' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
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
	if x.Ch != '6' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) LightBlue700() (r bnd.Bnd, ok bool) {
	if x.Ch != 'l' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'h' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'B' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
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
	if x.Ch != '7' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) LightBlue800() (r bnd.Bnd, ok bool) {
	if x.Ch != 'l' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'h' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'B' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
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
	if x.Ch != '8' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) LightBlue900() (r bnd.Bnd, ok bool) {
	if x.Ch != 'l' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'h' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'B' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
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
	if x.Ch != '9' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) LightBlueA100() (r bnd.Bnd, ok bool) {
	if x.Ch != 'l' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'h' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'B' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
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
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '1' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) LightBlueA200() (r bnd.Bnd, ok bool) {
	if x.Ch != 'l' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'h' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'B' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
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
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '2' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) LightBlueA400() (r bnd.Bnd, ok bool) {
	if x.Ch != 'l' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'h' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'B' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
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
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '4' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) LightBlueA700() (r bnd.Bnd, ok bool) {
	if x.Ch != 'l' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'h' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'B' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
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
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '7' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Cyan50() (r bnd.Bnd, ok bool) {
	if x.Ch != 'c' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'y' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '5' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Cyan100() (r bnd.Bnd, ok bool) {
	if x.Ch != 'c' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'y' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '1' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Cyan200() (r bnd.Bnd, ok bool) {
	if x.Ch != 'c' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'y' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '2' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Cyan300() (r bnd.Bnd, ok bool) {
	if x.Ch != 'c' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'y' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '3' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Cyan400() (r bnd.Bnd, ok bool) {
	if x.Ch != 'c' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'y' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '4' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Cyan500() (r bnd.Bnd, ok bool) {
	if x.Ch != 'c' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'y' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '5' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Cyan600() (r bnd.Bnd, ok bool) {
	if x.Ch != 'c' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'y' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '6' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Cyan700() (r bnd.Bnd, ok bool) {
	if x.Ch != 'c' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'y' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '7' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Cyan800() (r bnd.Bnd, ok bool) {
	if x.Ch != 'c' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'y' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '8' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Cyan900() (r bnd.Bnd, ok bool) {
	if x.Ch != 'c' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'y' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '9' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) CyanA100() (r bnd.Bnd, ok bool) {
	if x.Ch != 'c' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'y' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '1' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) CyanA200() (r bnd.Bnd, ok bool) {
	if x.Ch != 'c' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'y' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '2' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) CyanA400() (r bnd.Bnd, ok bool) {
	if x.Ch != 'c' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'y' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '4' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) CyanA700() (r bnd.Bnd, ok bool) {
	if x.Ch != 'c' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'y' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '7' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Teal50() (r bnd.Bnd, ok bool) {
	if x.Ch != 't' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '5' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Teal100() (r bnd.Bnd, ok bool) {
	if x.Ch != 't' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '1' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Teal200() (r bnd.Bnd, ok bool) {
	if x.Ch != 't' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '2' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Teal300() (r bnd.Bnd, ok bool) {
	if x.Ch != 't' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '3' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Teal400() (r bnd.Bnd, ok bool) {
	if x.Ch != 't' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '4' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Teal500() (r bnd.Bnd, ok bool) {
	if x.Ch != 't' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '5' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Teal600() (r bnd.Bnd, ok bool) {
	if x.Ch != 't' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '6' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Teal700() (r bnd.Bnd, ok bool) {
	if x.Ch != 't' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '7' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Teal800() (r bnd.Bnd, ok bool) {
	if x.Ch != 't' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '8' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Teal900() (r bnd.Bnd, ok bool) {
	if x.Ch != 't' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '9' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) TealA100() (r bnd.Bnd, ok bool) {
	if x.Ch != 't' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '1' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) TealA200() (r bnd.Bnd, ok bool) {
	if x.Ch != 't' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '2' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) TealA400() (r bnd.Bnd, ok bool) {
	if x.Ch != 't' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '4' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) TealA700() (r bnd.Bnd, ok bool) {
	if x.Ch != 't' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '7' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Green50() (r bnd.Bnd, ok bool) {
	if x.Ch != 'g' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '5' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Green100() (r bnd.Bnd, ok bool) {
	if x.Ch != 'g' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '1' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Green200() (r bnd.Bnd, ok bool) {
	if x.Ch != 'g' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '2' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Green300() (r bnd.Bnd, ok bool) {
	if x.Ch != 'g' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '3' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Green400() (r bnd.Bnd, ok bool) {
	if x.Ch != 'g' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '4' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Green500() (r bnd.Bnd, ok bool) {
	if x.Ch != 'g' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '5' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Green600() (r bnd.Bnd, ok bool) {
	if x.Ch != 'g' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '6' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Green700() (r bnd.Bnd, ok bool) {
	if x.Ch != 'g' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '7' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Green800() (r bnd.Bnd, ok bool) {
	if x.Ch != 'g' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '8' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Green900() (r bnd.Bnd, ok bool) {
	if x.Ch != 'g' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '9' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) GreenA100() (r bnd.Bnd, ok bool) {
	if x.Ch != 'g' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '1' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) GreenA200() (r bnd.Bnd, ok bool) {
	if x.Ch != 'g' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '2' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) GreenA400() (r bnd.Bnd, ok bool) {
	if x.Ch != 'g' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '4' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) GreenA700() (r bnd.Bnd, ok bool) {
	if x.Ch != 'g' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '7' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) LightGreen50() (r bnd.Bnd, ok bool) {
	if x.Ch != 'l' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'h' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'G' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '5' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) LightGreen100() (r bnd.Bnd, ok bool) {
	if x.Ch != 'l' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'h' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'G' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '1' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) LightGreen200() (r bnd.Bnd, ok bool) {
	if x.Ch != 'l' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'h' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'G' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '2' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) LightGreen300() (r bnd.Bnd, ok bool) {
	if x.Ch != 'l' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'h' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'G' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '3' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) LightGreen400() (r bnd.Bnd, ok bool) {
	if x.Ch != 'l' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'h' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'G' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '4' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) LightGreen500() (r bnd.Bnd, ok bool) {
	if x.Ch != 'l' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'h' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'G' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '5' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) LightGreen600() (r bnd.Bnd, ok bool) {
	if x.Ch != 'l' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'h' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'G' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '6' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) LightGreen700() (r bnd.Bnd, ok bool) {
	if x.Ch != 'l' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'h' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'G' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '7' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) LightGreen800() (r bnd.Bnd, ok bool) {
	if x.Ch != 'l' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'h' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'G' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '8' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) LightGreen900() (r bnd.Bnd, ok bool) {
	if x.Ch != 'l' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'h' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'G' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '9' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) LightGreenA100() (r bnd.Bnd, ok bool) {
	if x.Ch != 'l' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'h' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'G' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '1' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) LightGreenA200() (r bnd.Bnd, ok bool) {
	if x.Ch != 'l' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'h' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'G' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '2' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) LightGreenA400() (r bnd.Bnd, ok bool) {
	if x.Ch != 'l' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'h' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'G' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '4' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) LightGreenA700() (r bnd.Bnd, ok bool) {
	if x.Ch != 'l' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'h' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'G' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '7' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Lime50() (r bnd.Bnd, ok bool) {
	if x.Ch != 'l' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '5' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Lime100() (r bnd.Bnd, ok bool) {
	if x.Ch != 'l' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '1' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Lime200() (r bnd.Bnd, ok bool) {
	if x.Ch != 'l' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '2' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Lime300() (r bnd.Bnd, ok bool) {
	if x.Ch != 'l' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '3' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Lime400() (r bnd.Bnd, ok bool) {
	if x.Ch != 'l' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '4' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Lime500() (r bnd.Bnd, ok bool) {
	if x.Ch != 'l' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '5' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Lime600() (r bnd.Bnd, ok bool) {
	if x.Ch != 'l' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '6' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Lime700() (r bnd.Bnd, ok bool) {
	if x.Ch != 'l' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '7' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Lime800() (r bnd.Bnd, ok bool) {
	if x.Ch != 'l' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '8' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Lime900() (r bnd.Bnd, ok bool) {
	if x.Ch != 'l' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '9' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) LimeA100() (r bnd.Bnd, ok bool) {
	if x.Ch != 'l' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '1' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) LimeA200() (r bnd.Bnd, ok bool) {
	if x.Ch != 'l' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '2' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) LimeA400() (r bnd.Bnd, ok bool) {
	if x.Ch != 'l' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '4' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) LimeA700() (r bnd.Bnd, ok bool) {
	if x.Ch != 'l' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '7' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Yellow50() (r bnd.Bnd, ok bool) {
	if x.Ch != 'y' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'o' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'w' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '5' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Yellow100() (r bnd.Bnd, ok bool) {
	if x.Ch != 'y' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'o' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'w' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '1' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Yellow200() (r bnd.Bnd, ok bool) {
	if x.Ch != 'y' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'o' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'w' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '2' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Yellow300() (r bnd.Bnd, ok bool) {
	if x.Ch != 'y' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'o' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'w' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '3' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Yellow400() (r bnd.Bnd, ok bool) {
	if x.Ch != 'y' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'o' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'w' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '4' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Yellow500() (r bnd.Bnd, ok bool) {
	if x.Ch != 'y' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'o' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'w' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '5' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Yellow600() (r bnd.Bnd, ok bool) {
	if x.Ch != 'y' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'o' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'w' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '6' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Yellow700() (r bnd.Bnd, ok bool) {
	if x.Ch != 'y' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'o' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'w' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '7' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Yellow800() (r bnd.Bnd, ok bool) {
	if x.Ch != 'y' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'o' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'w' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '8' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Yellow900() (r bnd.Bnd, ok bool) {
	if x.Ch != 'y' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'o' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'w' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '9' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) YellowA100() (r bnd.Bnd, ok bool) {
	if x.Ch != 'y' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'o' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'w' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '1' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) YellowA200() (r bnd.Bnd, ok bool) {
	if x.Ch != 'y' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'o' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'w' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '2' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) YellowA400() (r bnd.Bnd, ok bool) {
	if x.Ch != 'y' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'o' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'w' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '4' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) YellowA700() (r bnd.Bnd, ok bool) {
	if x.Ch != 'y' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'o' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'w' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '7' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Amber50() (r bnd.Bnd, ok bool) {
	if x.Ch != 'a' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'b' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '5' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Amber100() (r bnd.Bnd, ok bool) {
	if x.Ch != 'a' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'b' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '1' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Amber200() (r bnd.Bnd, ok bool) {
	if x.Ch != 'a' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'b' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '2' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Amber300() (r bnd.Bnd, ok bool) {
	if x.Ch != 'a' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'b' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '3' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Amber400() (r bnd.Bnd, ok bool) {
	if x.Ch != 'a' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'b' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '4' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Amber500() (r bnd.Bnd, ok bool) {
	if x.Ch != 'a' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'b' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '5' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Amber600() (r bnd.Bnd, ok bool) {
	if x.Ch != 'a' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'b' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '6' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Amber700() (r bnd.Bnd, ok bool) {
	if x.Ch != 'a' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'b' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '7' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Amber800() (r bnd.Bnd, ok bool) {
	if x.Ch != 'a' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'b' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '8' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Amber900() (r bnd.Bnd, ok bool) {
	if x.Ch != 'a' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'b' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '9' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) AmberA100() (r bnd.Bnd, ok bool) {
	if x.Ch != 'a' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'b' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '1' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) AmberA200() (r bnd.Bnd, ok bool) {
	if x.Ch != 'a' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'b' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '2' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) AmberA400() (r bnd.Bnd, ok bool) {
	if x.Ch != 'a' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'b' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '4' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) AmberA700() (r bnd.Bnd, ok bool) {
	if x.Ch != 'a' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'b' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '7' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Orange50() (r bnd.Bnd, ok bool) {
	if x.Ch != 'o' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '5' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Orange100() (r bnd.Bnd, ok bool) {
	if x.Ch != 'o' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '1' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Orange200() (r bnd.Bnd, ok bool) {
	if x.Ch != 'o' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '2' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Orange300() (r bnd.Bnd, ok bool) {
	if x.Ch != 'o' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '3' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Orange400() (r bnd.Bnd, ok bool) {
	if x.Ch != 'o' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '4' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Orange500() (r bnd.Bnd, ok bool) {
	if x.Ch != 'o' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '5' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Orange600() (r bnd.Bnd, ok bool) {
	if x.Ch != 'o' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '6' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Orange700() (r bnd.Bnd, ok bool) {
	if x.Ch != 'o' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '7' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Orange800() (r bnd.Bnd, ok bool) {
	if x.Ch != 'o' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '8' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Orange900() (r bnd.Bnd, ok bool) {
	if x.Ch != 'o' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '9' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) OrangeA100() (r bnd.Bnd, ok bool) {
	if x.Ch != 'o' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '1' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) OrangeA200() (r bnd.Bnd, ok bool) {
	if x.Ch != 'o' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '2' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) OrangeA400() (r bnd.Bnd, ok bool) {
	if x.Ch != 'o' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '4' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) OrangeA700() (r bnd.Bnd, ok bool) {
	if x.Ch != 'o' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '7' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) DeepOrange50() (r bnd.Bnd, ok bool) {
	if x.Ch != 'd' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'O' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '5' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) DeepOrange100() (r bnd.Bnd, ok bool) {
	if x.Ch != 'd' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'O' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '1' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) DeepOrange200() (r bnd.Bnd, ok bool) {
	if x.Ch != 'd' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'O' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '2' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) DeepOrange300() (r bnd.Bnd, ok bool) {
	if x.Ch != 'd' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'O' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '3' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) DeepOrange400() (r bnd.Bnd, ok bool) {
	if x.Ch != 'd' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'O' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '4' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) DeepOrange500() (r bnd.Bnd, ok bool) {
	if x.Ch != 'd' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'O' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '5' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) DeepOrange600() (r bnd.Bnd, ok bool) {
	if x.Ch != 'd' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'O' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '6' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) DeepOrange700() (r bnd.Bnd, ok bool) {
	if x.Ch != 'd' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'O' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '7' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) DeepOrange800() (r bnd.Bnd, ok bool) {
	if x.Ch != 'd' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'O' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '8' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) DeepOrange900() (r bnd.Bnd, ok bool) {
	if x.Ch != 'd' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'O' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '9' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) DeepOrangeA100() (r bnd.Bnd, ok bool) {
	if x.Ch != 'd' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'O' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '1' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) DeepOrangeA200() (r bnd.Bnd, ok bool) {
	if x.Ch != 'd' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'O' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '2' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) DeepOrangeA400() (r bnd.Bnd, ok bool) {
	if x.Ch != 'd' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'O' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '4' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) DeepOrangeA700() (r bnd.Bnd, ok bool) {
	if x.Ch != 'd' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'O' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '7' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Brown50() (r bnd.Bnd, ok bool) {
	if x.Ch != 'b' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'o' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'w' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '5' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Brown100() (r bnd.Bnd, ok bool) {
	if x.Ch != 'b' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'o' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'w' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '1' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Brown200() (r bnd.Bnd, ok bool) {
	if x.Ch != 'b' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'o' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'w' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '2' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Brown300() (r bnd.Bnd, ok bool) {
	if x.Ch != 'b' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'o' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'w' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '3' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Brown400() (r bnd.Bnd, ok bool) {
	if x.Ch != 'b' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'o' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'w' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '4' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Brown500() (r bnd.Bnd, ok bool) {
	if x.Ch != 'b' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'o' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'w' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '5' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Brown600() (r bnd.Bnd, ok bool) {
	if x.Ch != 'b' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'o' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'w' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '6' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Brown700() (r bnd.Bnd, ok bool) {
	if x.Ch != 'b' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'o' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'w' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '7' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Brown800() (r bnd.Bnd, ok bool) {
	if x.Ch != 'b' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'o' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'w' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '8' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Brown900() (r bnd.Bnd, ok bool) {
	if x.Ch != 'b' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'o' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'w' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '9' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Grey50() (r bnd.Bnd, ok bool) {
	if x.Ch != 'g' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'y' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '5' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Grey100() (r bnd.Bnd, ok bool) {
	if x.Ch != 'g' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'y' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '1' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Grey200() (r bnd.Bnd, ok bool) {
	if x.Ch != 'g' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'y' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '2' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Grey300() (r bnd.Bnd, ok bool) {
	if x.Ch != 'g' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'y' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '3' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Grey400() (r bnd.Bnd, ok bool) {
	if x.Ch != 'g' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'y' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '4' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Grey500() (r bnd.Bnd, ok bool) {
	if x.Ch != 'g' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'y' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '5' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Grey600() (r bnd.Bnd, ok bool) {
	if x.Ch != 'g' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'y' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '6' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Grey700() (r bnd.Bnd, ok bool) {
	if x.Ch != 'g' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'y' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '7' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Grey800() (r bnd.Bnd, ok bool) {
	if x.Ch != 'g' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'y' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '8' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Grey900() (r bnd.Bnd, ok bool) {
	if x.Ch != 'g' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'y' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '9' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) BlueGrey50() (r bnd.Bnd, ok bool) {
	if x.Ch != 'b' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'l' {
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
	if x.Ch != 'G' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'y' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '5' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) BlueGrey100() (r bnd.Bnd, ok bool) {
	if x.Ch != 'b' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'l' {
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
	if x.Ch != 'G' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'y' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '1' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) BlueGrey200() (r bnd.Bnd, ok bool) {
	if x.Ch != 'b' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'l' {
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
	if x.Ch != 'G' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'y' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '2' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) BlueGrey300() (r bnd.Bnd, ok bool) {
	if x.Ch != 'b' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'l' {
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
	if x.Ch != 'G' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'y' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '3' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) BlueGrey400() (r bnd.Bnd, ok bool) {
	if x.Ch != 'b' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'l' {
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
	if x.Ch != 'G' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'y' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '4' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) BlueGrey500() (r bnd.Bnd, ok bool) {
	if x.Ch != 'b' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'l' {
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
	if x.Ch != 'G' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'y' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '5' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) BlueGrey600() (r bnd.Bnd, ok bool) {
	if x.Ch != 'b' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'l' {
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
	if x.Ch != 'G' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'y' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '6' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) BlueGrey700() (r bnd.Bnd, ok bool) {
	if x.Ch != 'b' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'l' {
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
	if x.Ch != 'G' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'y' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '7' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) BlueGrey800() (r bnd.Bnd, ok bool) {
	if x.Ch != 'b' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'l' {
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
	if x.Ch != 'G' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'y' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '8' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) BlueGrey900() (r bnd.Bnd, ok bool) {
	if x.Ch != 'b' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'l' {
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
	if x.Ch != 'G' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'y' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '9' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if x.Ch != '0' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Scl() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'c' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) StkWidth() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'k' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'W' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'h' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) ShpRadius() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'h' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'R' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'u' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) AxisPad() (r bnd.Bnd, ok bool) {
	if x.Ch != 'a' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'x' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) BarPad() (r bnd.Bnd, ok bool) {
	if x.Ch != 'b' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Len() (r bnd.Bnd, ok bool) {
	if x.Ch != 'l' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Pad() (r bnd.Bnd, ok bool) {
	if x.Ch != 'p' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) BakClr() (r bnd.Bnd, ok bool) {
	if x.Ch != 'b' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'k' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'C' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) BrdrClr() (r bnd.Bnd, ok bool) {
	if x.Ch != 'b' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'C' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) BrdrLen() (r bnd.Bnd, ok bool) {
	if x.Ch != 'b' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'L' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) InrvlTxtLen() (r bnd.Bnd, ok bool) {
	if x.Ch != 'i' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'v' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'T' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'x' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'L' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) InrvlTxtClrX() (r bnd.Bnd, ok bool) {
	if x.Ch != 'i' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'v' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'T' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'x' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'C' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'X' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) InrvlTxtClrY() (r bnd.Bnd, ok bool) {
	if x.Ch != 'i' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'v' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'T' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'x' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'C' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'Y' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) MsgClr() (r bnd.Bnd, ok bool) {
	if x.Ch != 'm' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'C' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) TitleClr() (r bnd.Bnd, ok bool) {
	if x.Ch != 't' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'C' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) PrfClr() (r bnd.Bnd, ok bool) {
	if x.Ch != 'p' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'f' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'C' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) LosClr() (r bnd.Bnd, ok bool) {
	if x.Ch != 'l' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'o' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'C' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) PrfPen() (r bnd.Bnd, ok bool) {
	if x.Ch != 'p' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'f' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) LosPen() (r bnd.Bnd, ok bool) {
	if x.Ch != 'l' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'o' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) OutlierLim() (r bnd.Bnd, ok bool) {
	if x.Ch != 'o' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'u' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'L' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Ifo() (r bnd.Bnd, ok bool) {
	if x.Ch != 'i' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'f' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'o' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Ifof() (r bnd.Bnd, ok bool) {
	if x.Ch != 'i' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'f' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'o' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'f' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Fmt() (r bnd.Bnd, ok bool) {
	if x.Ch != 'f' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Now() (r bnd.Bnd, ok bool) {
	if x.Ch != 'n' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'o' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'w' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) NewRng() (r bnd.Bnd, ok bool) {
	if x.Ch != 'n' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'w' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'R' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) NewRngArnd() (r bnd.Bnd, ok bool) {
	if x.Ch != 'n' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'w' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'R' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) NewRngFul() (r bnd.Bnd, ok bool) {
	if x.Ch != 'n' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'w' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'R' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'F' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'u' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) New() (r bnd.Bnd, ok bool) {
	if x.Ch != 'n' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'w' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Make() (r bnd.Bnd, ok bool) {
	if x.Ch != 'm' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'k' {
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
func (x *Trmr) MakeEmp() (r bnd.Bnd, ok bool) {
	if x.Ch != 'm' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'k' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'E' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) AddsLss() (r bnd.Bnd, ok bool) {
	if x.Ch != 'a' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'L' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) AddsLeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 'a' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'L' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SubsGtr() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'u' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'b' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'G' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SubsGeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'u' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'b' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'G' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) MulsLss() (r bnd.Bnd, ok bool) {
	if x.Ch != 'm' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'u' {
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
	if x.Ch != 'L' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) MulsLeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 'm' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'u' {
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
	if x.Ch != 'L' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) DivsGtr() (r bnd.Bnd, ok bool) {
	if x.Ch != 'd' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'v' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'G' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) DivsGeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 'd' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'v' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'G' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) FibsLeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 'f' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'b' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'L' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) NewRngs() (r bnd.Bnd, ok bool) {
	if x.Ch != 'n' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'w' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'R' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) MakeRngs() (r bnd.Bnd, ok bool) {
	if x.Ch != 'm' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'k' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'R' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) MakeEmpRngs() (r bnd.Bnd, ok bool) {
	if x.Ch != 'm' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'k' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'E' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'R' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) NewTrds() (r bnd.Bnd, ok bool) {
	if x.Ch != 'n' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'w' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'T' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) MakeTrds() (r bnd.Bnd, ok bool) {
	if x.Ch != 'm' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'k' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'T' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) MakeEmpTrds() (r bnd.Bnd, ok bool) {
	if x.Ch != 'm' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'k' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'E' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'T' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) NewPrfms() (r bnd.Bnd, ok bool) {
	if x.Ch != 'n' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'w' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'f' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) MakePrfms() (r bnd.Bnd, ok bool) {
	if x.Ch != 'm' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'k' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'f' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) MakeEmpPrfms() (r bnd.Bnd, ok bool) {
	if x.Ch != 'm' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'k' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'E' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'f' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Oan() (r bnd.Bnd, ok bool) {
	if x.Ch != 'o' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) NewPrvs() (r bnd.Bnd, ok bool) {
	if x.Ch != 'n' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'w' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'v' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) MakePrvs() (r bnd.Bnd, ok bool) {
	if x.Ch != 'm' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'k' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'v' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) MakeEmpPrvs() (r bnd.Bnd, ok bool) {
	if x.Ch != 'm' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'k' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'E' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'v' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) NewInstrs() (r bnd.Bnd, ok bool) {
	if x.Ch != 'n' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'w' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'I' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) MakeInstrs() (r bnd.Bnd, ok bool) {
	if x.Ch != 'm' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'k' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'I' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) MakeEmpInstrs() (r bnd.Bnd, ok bool) {
	if x.Ch != 'm' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'k' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'E' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'I' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) NewInrvls() (r bnd.Bnd, ok bool) {
	if x.Ch != 'n' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'w' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'I' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'v' {
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
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) MakeInrvls() (r bnd.Bnd, ok bool) {
	if x.Ch != 'm' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'k' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'I' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'v' {
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
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) MakeEmpInrvls() (r bnd.Bnd, ok bool) {
	if x.Ch != 'm' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'k' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'E' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'I' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'v' {
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
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) NewSides() (r bnd.Bnd, ok bool) {
	if x.Ch != 'n' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'w' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'S' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) MakeSides() (r bnd.Bnd, ok bool) {
	if x.Ch != 'm' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'k' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'S' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) MakeEmpSides() (r bnd.Bnd, ok bool) {
	if x.Ch != 'm' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'k' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'E' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'S' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) NewStms() (r bnd.Bnd, ok bool) {
	if x.Ch != 'n' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'w' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'S' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) MakeStms() (r bnd.Bnd, ok bool) {
	if x.Ch != 'm' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'k' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'S' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) MakeEmpStms() (r bnd.Bnd, ok bool) {
	if x.Ch != 'm' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'k' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'E' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'S' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) NewCnds() (r bnd.Bnd, ok bool) {
	if x.Ch != 'n' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'w' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'C' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) MakeCnds() (r bnd.Bnd, ok bool) {
	if x.Ch != 'm' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'k' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'C' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) MakeEmpCnds() (r bnd.Bnd, ok bool) {
	if x.Ch != 'm' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'k' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'E' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'C' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) NewStgys() (r bnd.Bnd, ok bool) {
	if x.Ch != 'n' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'w' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'S' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'y' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) MakeStgys() (r bnd.Bnd, ok bool) {
	if x.Ch != 'm' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'k' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'S' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'y' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) MakeEmpStgys() (r bnd.Bnd, ok bool) {
	if x.Ch != 'm' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'k' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'E' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'S' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'y' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Rgba() (r bnd.Bnd, ok bool) {
	if x.Ch != 'r' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'b' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Rgb() (r bnd.Bnd, ok bool) {
	if x.Ch != 'r' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'b' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Hex() (r bnd.Bnd, ok bool) {
	if x.Ch != 'h' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'x' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) NewPens() (r bnd.Bnd, ok bool) {
	if x.Ch != 'n' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'w' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) MakePens() (r bnd.Bnd, ok bool) {
	if x.Ch != 'm' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'k' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) MakeEmpPens() (r bnd.Bnd, ok bool) {
	if x.Ch != 'm' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'k' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'E' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) NewPlts() (r bnd.Bnd, ok bool) {
	if x.Ch != 'n' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'w' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) MakePlts() (r bnd.Bnd, ok bool) {
	if x.Ch != 'm' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'k' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) MakeEmpPlts() (r bnd.Bnd, ok bool) {
	if x.Ch != 'm' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'k' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'E' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) NewStm() (r bnd.Bnd, ok bool) {
	if x.Ch != 'n' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'w' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'S' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) NewFltsSctr() (r bnd.Bnd, ok bool) {
	if x.Ch != 'n' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'w' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'F' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'S' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'c' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) NewFltsSctrDist() (r bnd.Bnd, ok bool) {
	if x.Ch != 'n' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'w' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'F' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'S' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'c' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'D' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) NewHrz() (r bnd.Bnd, ok bool) {
	if x.Ch != 'n' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'w' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'H' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'z' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) NewVrt() (r bnd.Bnd, ok bool) {
	if x.Ch != 'n' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'w' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'V' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) NewDpth() (r bnd.Bnd, ok bool) {
	if x.Ch != 'n' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'w' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'D' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'h' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) NewMu() (r bnd.Bnd, ok bool) {
	if x.Ch != 'n' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'w' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'M' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'u' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Lower() (r bnd.Bnd, ok bool) {
	if x.Ch != 'l' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'o' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'w' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Upper() (r bnd.Bnd, ok bool) {
	if x.Ch != 'u' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Eql() (r bnd.Bnd, ok bool) {
	if x.Ch != 'e' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Neq() (r bnd.Bnd, ok bool) {
	if x.Ch != 'n' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Lss() (r bnd.Bnd, ok bool) {
	if x.Ch != 'l' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Gtr() (r bnd.Bnd, ok bool) {
	if x.Ch != 'g' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Leq() (r bnd.Bnd, ok bool) {
	if x.Ch != 'l' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Geq() (r bnd.Bnd, ok bool) {
	if x.Ch != 'g' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Not() (r bnd.Bnd, ok bool) {
	if x.Ch != 'n' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'o' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Trnc() (r bnd.Bnd, ok bool) {
	if x.Ch != 't' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'c' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) IsNaN() (r bnd.Bnd, ok bool) {
	if x.Ch != 'i' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'N' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'N' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) IsInfPos() (r bnd.Bnd, ok bool) {
	if x.Ch != 'i' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'I' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'f' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'o' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) IsInfNeg() (r bnd.Bnd, ok bool) {
	if x.Ch != 'i' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'I' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'f' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'N' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) IsValid() (r bnd.Bnd, ok bool) {
	if x.Ch != 'i' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'V' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Pct() (r bnd.Bnd, ok bool) {
	if x.Ch != 'p' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'c' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Pos() (r bnd.Bnd, ok bool) {
	if x.Ch != 'p' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'o' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Neg() (r bnd.Bnd, ok bool) {
	if x.Ch != 'n' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Inv() (r bnd.Bnd, ok bool) {
	if x.Ch != 'i' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'v' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Add() (r bnd.Bnd, ok bool) {
	if x.Ch != 'a' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Sub() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'u' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'b' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Mul() (r bnd.Bnd, ok bool) {
	if x.Ch != 'm' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'u' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Div() (r bnd.Bnd, ok bool) {
	if x.Ch != 'd' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'v' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Rem() (r bnd.Bnd, ok bool) {
	if x.Ch != 'r' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Pow() (r bnd.Bnd, ok bool) {
	if x.Ch != 'p' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'o' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'w' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Sqr() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Sqrt() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Mid() (r bnd.Bnd, ok bool) {
	if x.Ch != 'm' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Avg() (r bnd.Bnd, ok bool) {
	if x.Ch != 'a' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'v' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) AvgGeo() (r bnd.Bnd, ok bool) {
	if x.Ch != 'a' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'v' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'G' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'o' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelEql() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'E' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelNeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'N' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelLss() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'L' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelGtr() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'G' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelLeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'L' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelGeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'G' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) WeekdayCnt() (r bnd.Bnd, ok bool) {
	if x.Ch != 'w' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'k' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'y' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'C' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Dte() (r bnd.Bnd, ok bool) {
	if x.Ch != 'd' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 't' {
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
func (x *Trmr) ToSunday() (r bnd.Bnd, ok bool) {
	if x.Ch != 't' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'o' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'S' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'u' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'y' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) ToMonday() (r bnd.Bnd, ok bool) {
	if x.Ch != 't' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'o' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'M' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'o' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'y' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) ToTuesday() (r bnd.Bnd, ok bool) {
	if x.Ch != 't' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'o' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'T' {
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
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'y' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) ToWednesday() (r bnd.Bnd, ok bool) {
	if x.Ch != 't' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'o' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'W' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'y' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) ToThursday() (r bnd.Bnd, ok bool) {
	if x.Ch != 't' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'o' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'T' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'h' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'u' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'y' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) ToFriday() (r bnd.Bnd, ok bool) {
	if x.Ch != 't' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'o' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'F' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'y' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) ToSaturday() (r bnd.Bnd, ok bool) {
	if x.Ch != 't' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'o' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'S' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'u' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'y' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) IsSunday() (r bnd.Bnd, ok bool) {
	if x.Ch != 'i' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'S' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'u' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'y' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) IsMonday() (r bnd.Bnd, ok bool) {
	if x.Ch != 'i' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'M' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'o' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'y' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) IsTuesday() (r bnd.Bnd, ok bool) {
	if x.Ch != 'i' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'T' {
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
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'y' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) IsWednesday() (r bnd.Bnd, ok bool) {
	if x.Ch != 'i' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'W' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'y' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) IsThursday() (r bnd.Bnd, ok bool) {
	if x.Ch != 'i' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'T' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'h' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'u' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'y' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) IsFriday() (r bnd.Bnd, ok bool) {
	if x.Ch != 'i' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'F' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'y' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) IsSaturday() (r bnd.Bnd, ok bool) {
	if x.Ch != 'i' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'S' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'u' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'y' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Cnt() (r bnd.Bnd, ok bool) {
	if x.Ch != 'c' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) LstIdx() (r bnd.Bnd, ok bool) {
	if x.Ch != 'l' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'I' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'x' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Ensure() (r bnd.Bnd, ok bool) {
	if x.Ch != 'e' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'u' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
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
func (x *Trmr) MinSub() (r bnd.Bnd, ok bool) {
	if x.Ch != 'm' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'S' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'u' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'b' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) MaxAdd() (r bnd.Bnd, ok bool) {
	if x.Ch != 'm' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'x' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Mrg() (r bnd.Bnd, ok bool) {
	if x.Ch != 'm' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Cpy() (r bnd.Bnd, ok bool) {
	if x.Ch != 'c' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'y' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Rand() (r bnd.Bnd, ok bool) {
	if x.Ch != 'r' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Push() (r bnd.Bnd, ok bool) {
	if x.Ch != 'p' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'u' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'h' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Pop() (r bnd.Bnd, ok bool) {
	if x.Ch != 'p' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'o' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Que() (r bnd.Bnd, ok bool) {
	if x.Ch != 'q' {
		return r, false
	}
	r.Idx = x.Idx
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
func (x *Trmr) Dque() (r bnd.Bnd, ok bool) {
	if x.Ch != 'd' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'q' {
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
func (x *Trmr) Ins() (r bnd.Bnd, ok bool) {
	if x.Ch != 'i' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Upd() (r bnd.Bnd, ok bool) {
	if x.Ch != 'u' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Del() (r bnd.Bnd, ok bool) {
	if x.Ch != 'd' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) At() (r bnd.Bnd, ok bool) {
	if x.Ch != 'a' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) In() (r bnd.Bnd, ok bool) {
	if x.Ch != 'i' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) InBnd() (r bnd.Bnd, ok bool) {
	if x.Ch != 'i' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'B' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) From() (r bnd.Bnd, ok bool) {
	if x.Ch != 'f' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'o' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) To() (r bnd.Bnd, ok bool) {
	if x.Ch != 't' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'o' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Fst() (r bnd.Bnd, ok bool) {
	if x.Ch != 'f' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Mdl() (r bnd.Bnd, ok bool) {
	if x.Ch != 'm' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Lst() (r bnd.Bnd, ok bool) {
	if x.Ch != 'l' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) FstIdx() (r bnd.Bnd, ok bool) {
	if x.Ch != 'f' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'I' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'x' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) MdlIdx() (r bnd.Bnd, ok bool) {
	if x.Ch != 'm' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'I' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'x' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Rev() (r bnd.Bnd, ok bool) {
	if x.Ch != 'r' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'v' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SrchIdxEql() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'c' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'h' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'I' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'x' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'E' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SrchIdx() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'c' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'h' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'I' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'x' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Has() (r bnd.Bnd, ok bool) {
	if x.Ch != 'h' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SrtAsc() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'c' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SrtDsc() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'D' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'c' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) UnaPos() (r bnd.Bnd, ok bool) {
	if x.Ch != 'u' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'o' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) UnaNeg() (r bnd.Bnd, ok bool) {
	if x.Ch != 'u' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'N' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) UnaInv() (r bnd.Bnd, ok bool) {
	if x.Ch != 'u' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'I' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'v' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) UnaSqr() (r bnd.Bnd, ok bool) {
	if x.Ch != 'u' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'S' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) UnaSqrt() (r bnd.Bnd, ok bool) {
	if x.Ch != 'u' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'S' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SclAdd() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'c' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SclSub() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'c' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'S' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'u' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'b' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SclMul() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'c' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'M' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'u' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SclDiv() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'c' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'D' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'v' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SclRem() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'c' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'R' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SclPow() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'c' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'o' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'w' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SclMin() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'c' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'M' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SclMax() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'c' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'M' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'x' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) CntEql() (r bnd.Bnd, ok bool) {
	if x.Ch != 'c' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'E' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) CntNeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 'c' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'N' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) CntLss() (r bnd.Bnd, ok bool) {
	if x.Ch != 'c' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'L' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) CntGtr() (r bnd.Bnd, ok bool) {
	if x.Ch != 'c' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'G' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) CntLeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 'c' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'L' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) CntGeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 'c' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'G' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) InrAdd() (r bnd.Bnd, ok bool) {
	if x.Ch != 'i' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) InrSub() (r bnd.Bnd, ok bool) {
	if x.Ch != 'i' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'S' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'u' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'b' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) InrMul() (r bnd.Bnd, ok bool) {
	if x.Ch != 'i' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'M' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'u' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) InrDiv() (r bnd.Bnd, ok bool) {
	if x.Ch != 'i' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'D' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'v' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) InrRem() (r bnd.Bnd, ok bool) {
	if x.Ch != 'i' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'R' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) InrPow() (r bnd.Bnd, ok bool) {
	if x.Ch != 'i' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'o' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'w' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) InrMin() (r bnd.Bnd, ok bool) {
	if x.Ch != 'i' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'M' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) InrMax() (r bnd.Bnd, ok bool) {
	if x.Ch != 'i' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'M' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'x' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Sum() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'u' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Prd() (r bnd.Bnd, ok bool) {
	if x.Ch != 'p' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Mdn() (r bnd.Bnd, ok bool) {
	if x.Ch != 'm' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Sma() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Gma() (r bnd.Bnd, ok bool) {
	if x.Ch != 'g' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Wma() (r bnd.Bnd, ok bool) {
	if x.Ch != 'w' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Vrnc() (r bnd.Bnd, ok bool) {
	if x.Ch != 'v' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'c' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Std() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Zscr() (r bnd.Bnd, ok bool) {
	if x.Ch != 'z' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'c' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) ZscrInplace() (r bnd.Bnd, ok bool) {
	if x.Ch != 'z' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'c' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'I' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'c' {
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
func (x *Trmr) RngFul() (r bnd.Bnd, ok bool) {
	if x.Ch != 'r' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'F' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'u' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) RngLst() (r bnd.Bnd, ok bool) {
	if x.Ch != 'r' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'L' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) ProLst() (r bnd.Bnd, ok bool) {
	if x.Ch != 'p' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'o' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'L' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) ProSma() (r bnd.Bnd, ok bool) {
	if x.Ch != 'p' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'o' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'S' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SubSumPos() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'u' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'b' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'S' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'u' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'o' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SubSumNeg() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'u' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'b' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'S' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'u' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'N' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Rsi() (r bnd.Bnd, ok bool) {
	if x.Ch != 'r' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Wrsi() (r bnd.Bnd, ok bool) {
	if x.Ch != 'w' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Pro() (r bnd.Bnd, ok bool) {
	if x.Ch != 'p' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'o' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Alma() (r bnd.Bnd, ok bool) {
	if x.Ch != 'a' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) ProAlma() (r bnd.Bnd, ok bool) {
	if x.Ch != 'p' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'o' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) CntrDist() (r bnd.Bnd, ok bool) {
	if x.Ch != 'c' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'D' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Bnd() (r bnd.Bnd, ok bool) {
	if x.Ch != 'b' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) RngMrg() (r bnd.Bnd, ok bool) {
	if x.Ch != 'r' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'M' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) OpnMid() (r bnd.Bnd, ok bool) {
	if x.Ch != 'o' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'M' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) ClsMid() (r bnd.Bnd, ok bool) {
	if x.Ch != 'c' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'M' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelClsResEql() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'C' {
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
	if x.Ch != 'R' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'E' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelClsResNeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'C' {
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
	if x.Ch != 'R' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'N' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelClsResLss() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'C' {
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
	if x.Ch != 'R' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'L' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelClsResGtr() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'C' {
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
	if x.Ch != 'R' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'G' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelClsResLeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'C' {
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
	if x.Ch != 'R' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'L' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelClsResGeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'C' {
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
	if x.Ch != 'R' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'G' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelClsReqEql() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'C' {
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
	if x.Ch != 'R' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'E' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelClsReqNeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'C' {
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
	if x.Ch != 'R' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'N' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelClsReqLss() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'C' {
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
	if x.Ch != 'R' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'L' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelClsReqGtr() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'C' {
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
	if x.Ch != 'R' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'G' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelClsReqLeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'C' {
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
	if x.Ch != 'R' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'L' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelClsReqGeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'C' {
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
	if x.Ch != 'R' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'G' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelOpnResEql() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'O' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'R' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'E' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelOpnResNeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'O' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'R' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'N' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelOpnResLss() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'O' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'R' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'L' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelOpnResGtr() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'O' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'R' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'G' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelOpnResLeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'O' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'R' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'L' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelOpnResGeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'O' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'R' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'G' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelOpnReqEql() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'O' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'R' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'E' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelOpnReqNeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'O' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'R' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'N' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelOpnReqLss() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'O' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'R' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'L' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelOpnReqGtr() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'O' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'R' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'G' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelOpnReqLeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'O' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'R' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'L' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelOpnReqGeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'O' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'R' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'G' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelInstrEql() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'I' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'E' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelInstrNeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'I' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'N' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelInstrLss() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'I' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'L' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelInstrGtr() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'I' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'G' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelInstrLeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'I' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'L' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelInstrGeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'I' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'G' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelUnitsEql() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'U' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'E' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelUnitsNeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'U' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'N' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelUnitsLss() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'U' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'L' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelUnitsGtr() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'U' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'G' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelUnitsLeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'U' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'L' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelUnitsGeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'U' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'G' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelMrgnRtioEql() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'M' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'R' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'o' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'E' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelMrgnRtioNeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'M' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'R' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'o' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'N' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelMrgnRtioLss() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'M' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'R' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'o' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'L' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelMrgnRtioGtr() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'M' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'R' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'o' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'G' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelMrgnRtioLeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'M' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'R' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'o' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'L' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelMrgnRtioGeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'M' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'R' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'o' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'G' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelTrdPctEql() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'T' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'c' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'E' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelTrdPctNeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'T' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'c' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'N' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelTrdPctLss() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'T' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'c' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'L' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelTrdPctGtr() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'T' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'c' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'G' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelTrdPctLeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'T' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'c' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'L' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelTrdPctGeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'T' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'c' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'G' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelClsBalUsdActEql() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'C' {
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
	if x.Ch != 'B' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'U' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'c' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'E' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelClsBalUsdActNeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'C' {
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
	if x.Ch != 'B' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'U' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'c' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'N' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelClsBalUsdActLss() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'C' {
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
	if x.Ch != 'B' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'U' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'c' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'L' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelClsBalUsdActGtr() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'C' {
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
	if x.Ch != 'B' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'U' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'c' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'G' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelClsBalUsdActLeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'C' {
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
	if x.Ch != 'B' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'U' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'c' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'L' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelClsBalUsdActGeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'C' {
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
	if x.Ch != 'B' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'U' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'c' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'G' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelClsBalUsdEql() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'C' {
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
	if x.Ch != 'B' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'U' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'E' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelClsBalUsdNeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'C' {
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
	if x.Ch != 'B' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'U' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'N' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelClsBalUsdLss() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'C' {
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
	if x.Ch != 'B' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'U' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'L' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelClsBalUsdGtr() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'C' {
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
	if x.Ch != 'B' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'U' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'G' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelClsBalUsdLeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'C' {
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
	if x.Ch != 'B' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'U' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'L' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelClsBalUsdGeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'C' {
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
	if x.Ch != 'B' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'U' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'G' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelOpnBalUsdEql() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'O' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'B' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'U' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'E' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelOpnBalUsdNeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'O' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'B' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'U' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'N' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelOpnBalUsdLss() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'O' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'B' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'U' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'L' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelOpnBalUsdGtr() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'O' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'B' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'U' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'G' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelOpnBalUsdLeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'O' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'B' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'U' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'L' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelOpnBalUsdGeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'O' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'B' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'U' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'G' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelCstOpnSpdUsdEql() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'C' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'O' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'S' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'U' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'E' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelCstOpnSpdUsdNeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'C' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'O' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'S' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'U' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'N' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelCstOpnSpdUsdLss() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'C' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'O' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'S' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'U' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'L' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelCstOpnSpdUsdGtr() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'C' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'O' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'S' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'U' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'G' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelCstOpnSpdUsdLeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'C' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'O' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'S' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'U' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'L' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelCstOpnSpdUsdGeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'C' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'O' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'S' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'U' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'G' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelCstClsSpdUsdEql() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'C' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'C' {
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
	if x.Ch != 'S' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'U' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'E' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelCstClsSpdUsdNeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'C' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'C' {
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
	if x.Ch != 'S' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'U' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'N' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelCstClsSpdUsdLss() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'C' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'C' {
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
	if x.Ch != 'S' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'U' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'L' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelCstClsSpdUsdGtr() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'C' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'C' {
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
	if x.Ch != 'S' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'U' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'G' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelCstClsSpdUsdLeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'C' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'C' {
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
	if x.Ch != 'S' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'U' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'L' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelCstClsSpdUsdGeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'C' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'C' {
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
	if x.Ch != 'S' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'U' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'G' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelCstComUsdEql() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'C' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'C' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'o' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'U' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'E' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelCstComUsdNeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'C' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'C' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'o' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'U' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'N' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelCstComUsdLss() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'C' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'C' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'o' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'U' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'L' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelCstComUsdGtr() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'C' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'C' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'o' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'U' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'G' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelCstComUsdLeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'C' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'C' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'o' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'U' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'L' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelCstComUsdGeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'C' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'C' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'o' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'U' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'G' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelPnlGrsUsdEql() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'G' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'U' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'E' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelPnlGrsUsdNeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'G' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'U' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'N' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelPnlGrsUsdLss() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'G' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'U' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'L' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelPnlGrsUsdGtr() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'G' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'U' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'G' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelPnlGrsUsdLeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'G' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'U' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'L' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelPnlGrsUsdGeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'G' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'U' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'G' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelPnlUsdEql() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'U' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'E' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelPnlUsdNeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'U' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'N' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelPnlUsdLss() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'U' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'L' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelPnlUsdGtr() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'U' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'G' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelPnlUsdLeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'U' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'L' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelPnlUsdGeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'U' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'G' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelPnlPctPredictEql() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'c' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'c' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'E' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelPnlPctPredictNeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'c' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'c' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'N' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelPnlPctPredictLss() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'c' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'c' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'L' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelPnlPctPredictGtr() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'c' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'c' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'G' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelPnlPctPredictLeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'c' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'c' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'L' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelPnlPctPredictGeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'c' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'c' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'G' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelPnlPctEql() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'c' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'E' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelPnlPctNeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'c' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'N' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelPnlPctLss() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'c' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'L' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelPnlPctGtr() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'c' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'G' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelPnlPctLeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'c' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'L' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelPnlPctGeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'c' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'G' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelIsLongEql() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'I' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'L' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'o' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'E' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelIsLongNeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'I' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'L' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'o' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'N' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelDurEql() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'D' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'u' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'E' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelDurNeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'D' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'u' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'N' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelDurLss() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'D' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'u' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'L' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelDurGtr() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'D' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'u' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'G' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelDurLeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'D' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'u' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'L' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelDurGeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'D' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'u' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'G' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelPipEql() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'E' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelPipNeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'N' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelPipLss() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'L' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelPipGtr() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'G' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelPipLeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'L' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelPipGeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'G' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelClsRsnEql() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'C' {
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
	if x.Ch != 'R' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'E' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelClsRsnNeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'C' {
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
	if x.Ch != 'R' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'N' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelClsRsnLss() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'C' {
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
	if x.Ch != 'R' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'L' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelClsRsnGtr() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'C' {
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
	if x.Ch != 'R' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'G' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelClsRsnLeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'C' {
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
	if x.Ch != 'R' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'L' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelClsRsnGeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'C' {
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
	if x.Ch != 'R' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'G' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelClsSpdEql() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'C' {
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
	if x.Ch != 'S' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'E' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelClsSpdNeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'C' {
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
	if x.Ch != 'S' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'N' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelClsSpdLss() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'C' {
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
	if x.Ch != 'S' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'L' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelClsSpdGtr() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'C' {
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
	if x.Ch != 'S' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'G' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelClsSpdLeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'C' {
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
	if x.Ch != 'S' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'L' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelClsSpdGeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'C' {
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
	if x.Ch != 'S' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'G' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelOpnSpdEql() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'O' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'S' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'E' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelOpnSpdNeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'O' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'S' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'N' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelOpnSpdLss() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'O' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'S' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'L' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelOpnSpdGtr() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'O' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'S' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'G' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelOpnSpdLeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'O' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'S' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'L' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelOpnSpdGeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'O' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'S' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'G' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelClsAskEql() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'C' {
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
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'k' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'E' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelClsAskNeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'C' {
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
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'k' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'N' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelClsAskLss() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'C' {
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
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'k' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'L' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelClsAskGtr() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'C' {
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
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'k' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'G' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelClsAskLeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'C' {
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
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'k' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'L' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelClsAskGeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'C' {
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
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'k' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'G' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelOpnAskEql() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'O' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'k' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'E' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelOpnAskNeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'O' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'k' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'N' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelOpnAskLss() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'O' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'k' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'L' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelOpnAskGtr() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'O' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'k' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'G' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelOpnAskLeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'O' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'k' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'L' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelOpnAskGeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'O' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'k' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'G' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelClsBidEql() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'C' {
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
	if x.Ch != 'B' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'E' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelClsBidNeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'C' {
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
	if x.Ch != 'B' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'N' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelClsBidLss() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'C' {
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
	if x.Ch != 'B' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'L' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelClsBidGtr() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'C' {
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
	if x.Ch != 'B' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'G' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelClsBidLeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'C' {
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
	if x.Ch != 'B' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'L' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelClsBidGeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'C' {
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
	if x.Ch != 'B' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'G' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelOpnBidEql() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'O' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'B' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'E' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelOpnBidNeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'O' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'B' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'N' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelOpnBidLss() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'O' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'B' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'L' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelOpnBidGtr() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'O' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'B' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'G' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelOpnBidLeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'O' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'B' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'L' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelOpnBidGeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'O' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'B' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'G' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelClsTmeEql() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'C' {
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
	if x.Ch != 'T' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'E' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelClsTmeNeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'C' {
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
	if x.Ch != 'T' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'N' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelClsTmeLss() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'C' {
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
	if x.Ch != 'T' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'L' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelClsTmeGtr() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'C' {
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
	if x.Ch != 'T' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'G' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelClsTmeLeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'C' {
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
	if x.Ch != 'T' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'L' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelClsTmeGeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'C' {
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
	if x.Ch != 'T' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'G' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelOpnTmeEql() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'O' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'T' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'E' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelOpnTmeNeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'O' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'T' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'N' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelOpnTmeLss() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'O' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'T' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'L' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelOpnTmeGtr() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'O' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'T' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'G' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelOpnTmeLeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'O' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'T' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'L' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SelOpnTmeGeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'O' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'T' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'G' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) OpnTmes() (r bnd.Bnd, ok bool) {
	if x.Ch != 'o' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'T' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) ClsTmes() (r bnd.Bnd, ok bool) {
	if x.Ch != 'c' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'T' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) OpnBids() (r bnd.Bnd, ok bool) {
	if x.Ch != 'o' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'B' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) ClsBids() (r bnd.Bnd, ok bool) {
	if x.Ch != 'c' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'B' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) OpnAsks() (r bnd.Bnd, ok bool) {
	if x.Ch != 'o' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'k' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) ClsAsks() (r bnd.Bnd, ok bool) {
	if x.Ch != 'c' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'k' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) OpnSpds() (r bnd.Bnd, ok bool) {
	if x.Ch != 'o' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'S' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) ClsSpds() (r bnd.Bnd, ok bool) {
	if x.Ch != 'c' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'S' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) ClsRsns() (r bnd.Bnd, ok bool) {
	if x.Ch != 'c' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'R' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Pips() (r bnd.Bnd, ok bool) {
	if x.Ch != 'p' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Durs() (r bnd.Bnd, ok bool) {
	if x.Ch != 'd' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'u' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) IsLongs() (r bnd.Bnd, ok bool) {
	if x.Ch != 'i' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'L' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'o' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) PnlPcts() (r bnd.Bnd, ok bool) {
	if x.Ch != 'p' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'c' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) PnlPctPredicts() (r bnd.Bnd, ok bool) {
	if x.Ch != 'p' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'c' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'c' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) PnlUsds() (r bnd.Bnd, ok bool) {
	if x.Ch != 'p' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'U' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) PnlGrsUsds() (r bnd.Bnd, ok bool) {
	if x.Ch != 'p' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'G' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'U' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) CstComUsds() (r bnd.Bnd, ok bool) {
	if x.Ch != 'c' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'C' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'o' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'U' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) CstClsSpdUsds() (r bnd.Bnd, ok bool) {
	if x.Ch != 'c' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'C' {
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
	if x.Ch != 'S' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'U' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) CstOpnSpdUsds() (r bnd.Bnd, ok bool) {
	if x.Ch != 'c' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'O' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'S' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'U' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) OpnBalUsds() (r bnd.Bnd, ok bool) {
	if x.Ch != 'o' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'B' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'U' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) ClsBalUsds() (r bnd.Bnd, ok bool) {
	if x.Ch != 'c' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'B' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'U' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) ClsBalUsdActs() (r bnd.Bnd, ok bool) {
	if x.Ch != 'c' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'B' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'U' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'c' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) TrdPcts() (r bnd.Bnd, ok bool) {
	if x.Ch != 't' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'c' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) MrgnRtios() (r bnd.Bnd, ok bool) {
	if x.Ch != 'm' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'R' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'o' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Unitss() (r bnd.Bnd, ok bool) {
	if x.Ch != 'u' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Instrs() (r bnd.Bnd, ok bool) {
	if x.Ch != 'i' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) OpnReqs() (r bnd.Bnd, ok bool) {
	if x.Ch != 'o' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'R' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) OpnRess() (r bnd.Bnd, ok bool) {
	if x.Ch != 'o' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'R' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) ClsReqs() (r bnd.Bnd, ok bool) {
	if x.Ch != 'c' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'R' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) ClsRess() (r bnd.Bnd, ok bool) {
	if x.Ch != 'c' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'R' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Dlt() (r bnd.Bnd, ok bool) {
	if x.Ch != 'd' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) ScsPcts() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'c' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'c' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) PipPerDays() (r bnd.Bnd, ok bool) {
	if x.Ch != 'p' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'D' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'y' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) UsdPerDays() (r bnd.Bnd, ok bool) {
	if x.Ch != 'u' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'D' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'y' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) ScsPerDays() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'c' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'D' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'y' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) OpnPerDays() (r bnd.Bnd, ok bool) {
	if x.Ch != 'o' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'D' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'y' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) PipAvgs() (r bnd.Bnd, ok bool) {
	if x.Ch != 'p' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'v' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) PipMdns() (r bnd.Bnd, ok bool) {
	if x.Ch != 'p' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'M' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) PipMins() (r bnd.Bnd, ok bool) {
	if x.Ch != 'p' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'M' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) PipMaxs() (r bnd.Bnd, ok bool) {
	if x.Ch != 'p' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'M' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'x' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) PipSums() (r bnd.Bnd, ok bool) {
	if x.Ch != 'p' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'S' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'u' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) DurAvgs() (r bnd.Bnd, ok bool) {
	if x.Ch != 'd' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'u' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'v' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) DurMdns() (r bnd.Bnd, ok bool) {
	if x.Ch != 'd' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'u' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'M' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) DurMins() (r bnd.Bnd, ok bool) {
	if x.Ch != 'd' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'u' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'M' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) DurMaxs() (r bnd.Bnd, ok bool) {
	if x.Ch != 'd' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'u' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'M' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'x' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) LosLimMaxs() (r bnd.Bnd, ok bool) {
	if x.Ch != 'l' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'o' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'L' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'M' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'x' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) DurLimMaxs() (r bnd.Bnd, ok bool) {
	if x.Ch != 'd' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'u' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'L' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'M' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'x' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) DayCnts() (r bnd.Bnd, ok bool) {
	if x.Ch != 'd' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'y' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'C' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) TrdCnts() (r bnd.Bnd, ok bool) {
	if x.Ch != 't' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'C' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) CstTotUsds() (r bnd.Bnd, ok bool) {
	if x.Ch != 'c' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'T' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'o' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'U' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) CstSpdUsds() (r bnd.Bnd, ok bool) {
	if x.Ch != 'c' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'S' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'U' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Pths() (r bnd.Bnd, ok bool) {
	if x.Ch != 'p' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'h' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Opa() (r bnd.Bnd, ok bool) {
	if x.Ch != 'o' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Vis() (r bnd.Bnd, ok bool) {
	if x.Ch != 'v' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) X() (r bnd.Bnd, ok bool) {
	if x.Ch != 'x' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Stm() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) StmBnd() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'B' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Cnd() (r bnd.Bnd, ok bool) {
	if x.Ch != 'c' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) HrzLn() (r bnd.Bnd, ok bool) {
	if x.Ch != 'h' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'z' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'L' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) VrtLn() (r bnd.Bnd, ok bool) {
	if x.Ch != 'v' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'L' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) HrzBnd() (r bnd.Bnd, ok bool) {
	if x.Ch != 'h' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'z' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'B' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) VrtBnd() (r bnd.Bnd, ok bool) {
	if x.Ch != 'v' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'B' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) HrzSclVal() (r bnd.Bnd, ok bool) {
	if x.Ch != 'h' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'z' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'S' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'c' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'V' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) VrtSclVal() (r bnd.Bnd, ok bool) {
	if x.Ch != 'v' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'S' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'c' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'V' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Sho() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'h' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'o' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Siz() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'z' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) HrzScl() (r bnd.Bnd, ok bool) {
	if x.Ch != 'h' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'z' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'S' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'c' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) VrtScl() (r bnd.Bnd, ok bool) {
	if x.Ch != 'v' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'S' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'c' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Flts() (r bnd.Bnd, ok bool) {
	if x.Ch != 'f' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) PrfLos() (r bnd.Bnd, ok bool) {
	if x.Ch != 'p' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'f' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'L' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'o' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Plt() (r bnd.Bnd, ok bool) {
	if x.Ch != 'p' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Lck() (r bnd.Bnd, ok bool) {
	if x.Ch != 'l' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'c' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'k' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Ulck() (r bnd.Bnd, ok bool) {
	if x.Ch != 'u' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'c' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'k' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Name() (r bnd.Bnd, ok bool) {
	if x.Ch != 'n' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'm' {
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
func (x *Trmr) EurUsd() (r bnd.Bnd, ok bool) {
	if x.Ch != 'e' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'u' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'U' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) AudUsd() (r bnd.Bnd, ok bool) {
	if x.Ch != 'a' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'u' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'U' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) NzdUsd() (r bnd.Bnd, ok bool) {
	if x.Ch != 'n' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'z' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'U' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) GbpUsd() (r bnd.Bnd, ok bool) {
	if x.Ch != 'g' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'b' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'U' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) I() (r bnd.Bnd, ok bool) {
	if x.Ch != 'i' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Bid() (r bnd.Bnd, ok bool) {
	if x.Ch != 'b' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Ask() (r bnd.Bnd, ok bool) {
	if x.Ch != 'a' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'k' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Sar() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Ema() (r bnd.Bnd, ok bool) {
	if x.Ch != 'e' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) AggFst() (r bnd.Bnd, ok bool) {
	if x.Ch != 'a' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'F' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) AggLst() (r bnd.Bnd, ok bool) {
	if x.Ch != 'a' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'L' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) AggSum() (r bnd.Bnd, ok bool) {
	if x.Ch != 'a' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'S' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'u' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) AggPrd() (r bnd.Bnd, ok bool) {
	if x.Ch != 'a' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) AggMin() (r bnd.Bnd, ok bool) {
	if x.Ch != 'a' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'M' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) AggMax() (r bnd.Bnd, ok bool) {
	if x.Ch != 'a' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'M' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'x' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) AggMid() (r bnd.Bnd, ok bool) {
	if x.Ch != 'a' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'M' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) AggMdn() (r bnd.Bnd, ok bool) {
	if x.Ch != 'a' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'M' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) AggSma() (r bnd.Bnd, ok bool) {
	if x.Ch != 'a' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'S' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) AggGma() (r bnd.Bnd, ok bool) {
	if x.Ch != 'a' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'G' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) AggWma() (r bnd.Bnd, ok bool) {
	if x.Ch != 'a' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'W' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) AggRsi() (r bnd.Bnd, ok bool) {
	if x.Ch != 'a' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'R' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) AggWrsi() (r bnd.Bnd, ok bool) {
	if x.Ch != 'a' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'W' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) AggAlma() (r bnd.Bnd, ok bool) {
	if x.Ch != 'a' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) AggVrnc() (r bnd.Bnd, ok bool) {
	if x.Ch != 'a' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'V' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'c' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) AggStd() (r bnd.Bnd, ok bool) {
	if x.Ch != 'a' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'S' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) AggRngFul() (r bnd.Bnd, ok bool) {
	if x.Ch != 'a' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'R' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'F' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'u' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) AggRngLst() (r bnd.Bnd, ok bool) {
	if x.Ch != 'a' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'R' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'L' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) AggProLst() (r bnd.Bnd, ok bool) {
	if x.Ch != 'a' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'o' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'L' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) AggProSma() (r bnd.Bnd, ok bool) {
	if x.Ch != 'a' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'o' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'S' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) AggProAlma() (r bnd.Bnd, ok bool) {
	if x.Ch != 'a' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'o' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) AggEma() (r bnd.Bnd, ok bool) {
	if x.Ch != 'a' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'E' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) InrSlp() (r bnd.Bnd, ok bool) {
	if x.Ch != 'i' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'S' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'p' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) OtrAdd() (r bnd.Bnd, ok bool) {
	if x.Ch != 'o' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'A' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) OtrSub() (r bnd.Bnd, ok bool) {
	if x.Ch != 'o' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'S' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'u' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'b' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) OtrMul() (r bnd.Bnd, ok bool) {
	if x.Ch != 'o' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'M' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'u' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) OtrDiv() (r bnd.Bnd, ok bool) {
	if x.Ch != 'o' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'D' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'v' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) OtrRem() (r bnd.Bnd, ok bool) {
	if x.Ch != 'o' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'R' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) OtrPow() (r bnd.Bnd, ok bool) {
	if x.Ch != 'o' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'P' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'o' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'w' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) OtrMin() (r bnd.Bnd, ok bool) {
	if x.Ch != 'o' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'M' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'i' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) OtrMax() (r bnd.Bnd, ok bool) {
	if x.Ch != 'o' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'M' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'x' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SclEql() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'c' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'E' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SclNeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'c' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'N' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SclLss() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'c' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'L' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SclGtr() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'c' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'G' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SclLeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'c' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'L' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) SclGeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'c' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'G' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) InrEql() (r bnd.Bnd, ok bool) {
	if x.Ch != 'i' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'E' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) InrNeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 'i' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'N' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) InrLss() (r bnd.Bnd, ok bool) {
	if x.Ch != 'i' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'L' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) InrGtr() (r bnd.Bnd, ok bool) {
	if x.Ch != 'i' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'G' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) InrLeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 'i' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'L' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) InrGeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 'i' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'G' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) OtrEql() (r bnd.Bnd, ok bool) {
	if x.Ch != 'o' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'E' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) OtrNeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 'o' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'N' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) OtrLss() (r bnd.Bnd, ok bool) {
	if x.Ch != 'o' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'L' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) OtrGtr() (r bnd.Bnd, ok bool) {
	if x.Ch != 'o' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'G' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) OtrLeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 'o' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'L' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) OtrGeq() (r bnd.Bnd, ok bool) {
	if x.Ch != 'o' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'G' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) And() (r bnd.Bnd, ok bool) {
	if x.Ch != 'a' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Seq() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'q' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Stgy() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'y' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) MayTrd() (r bnd.Bnd, ok bool) {
	if x.Ch != 'm' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'y' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'T' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Log() (r bnd.Bnd, ok bool) {
	if x.Ch != 'l' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'o' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'g' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Str() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Bol() (r bnd.Bnd, ok bool) {
	if x.Ch != 'b' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'o' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Flt() (r bnd.Bnd, ok bool) {
	if x.Ch != 'f' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Unt() (r bnd.Bnd, ok bool) {
	if x.Ch != 'u' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Int() (r bnd.Bnd, ok bool) {
	if x.Ch != 'i' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Tme() (r bnd.Bnd, ok bool) {
	if x.Ch != 't' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'm' {
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
func (x *Trmr) Strs() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'r' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Bols() (r bnd.Bnd, ok bool) {
	if x.Ch != 'b' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'o' {
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
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Unts() (r bnd.Bnd, ok bool) {
	if x.Ch != 'u' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Ints() (r bnd.Bnd, ok bool) {
	if x.Ch != 'i' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Tmes() (r bnd.Bnd, ok bool) {
	if x.Ch != 't' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'm' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Bnds() (r bnd.Bnd, ok bool) {
	if x.Ch != 'b' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'd' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Ana() (r bnd.Bnd, ok bool) {
	if x.Ch != 'a' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'a' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Hst() (r bnd.Bnd, ok bool) {
	if x.Ch != 'h' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Rlt() (r bnd.Bnd, ok bool) {
	if x.Ch != 'r' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'l' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Fnt() (r bnd.Bnd, ok bool) {
	if x.Ch != 'f' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 't' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Pen() (r bnd.Bnd, ok bool) {
	if x.Ch != 'p' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'e' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 'n' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) Sys() (r bnd.Bnd, ok bool) {
	if x.Ch != 's' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	if x.Ch != 'y' {
		return r, false
	}
	x.NextRune()
	if x.Ch != 's' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
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
	fls, ok := x.Fls()
	if ok {
		return BolLit{Bnd: fls}, true
	}
	x.Scn = scn
	tru, ok := x.Tru()
	if ok {
		return BolLit{Bnd: tru}, true
	}
	x.Scn = scn
	return r, false
}
func (x *Trmr) FltLit(skpDsh ...bool) (r FltLit, ok bool) {
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
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' || (len(skpDsh) == 0 && x.Ch == '-') {
		return r, false
	}
	r.Lim = x.Idx
	return r, r.Idx != r.Lim
}
func (x *Trmr) UntLit() (r UntLit, ok bool) {
	r.Idx = x.Idx
	for unicode.IsDigit(x.Ch) {
		x.NextRune()
	}
	if x.Ch == '-' || unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' { // check dash for tme/bnd
		return r, false
	}
	r.Lim = x.Idx
	return r, r.Idx != r.Lim
}
func (x *Trmr) IntLit() (r IntLit, ok bool) {
	r.Idx = x.Idx
	if x.Ch != '+' && x.Ch != '-' {
		return r, false
	}
	x.NextRune()
	if !unicode.IsDigit(x.Ch) { // ch must be digit
		return r, false
	}
	x.NextRune()
	for !x.End && unicode.IsDigit(x.Ch) {
		x.NextRune()
	}
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, r.Idx != r.Lim
}
func (x *Trmr) TmePrtLit(suffix rune) (r TmePrtLit, ok bool) {
	r.Idx = x.Idx
	if x.Ch == '-' { // optional minus
		x.NextRune()
		if !unicode.IsDigit(x.Ch) { // next ch must be digit
			return r, false
		}
		x.NextRune()
	}
	for unicode.IsDigit(x.Ch) {
		x.NextRune()
	}
	if r.Idx == x.Idx || (x.Idx-r.Idx == 1 && x.Txt[r.Idx] == '-') {
		return r, false
	}
	if x.Ch != suffix { // 1s, -3s, 1m, -3m, 1h, -3h, 1d, -3d, 1w, -3w
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) TmeLit(skpDsh ...bool) (r TmeLit, ok bool) {
	r.Idx = x.Idx
	scn := x.Scn
	r.Year, ok = x.TmePrtLit('y')
	if ok {
		scn = x.Scn
	} else {
		x.Scn = scn
	}
	r.Month, ok = x.TmePrtLit('n')
	if ok {
		scn = x.Scn
	} else {
		x.Scn = scn
	}
	r.Week, ok = x.TmePrtLit('w')
	if ok {
		scn = x.Scn
	} else {
		x.Scn = scn
	}
	r.Day, ok = x.TmePrtLit('d')
	if ok {
		scn = x.Scn
	} else {
		x.Scn = scn
	}
	r.Hour, ok = x.TmePrtLit('h')
	if ok {
		scn = x.Scn
	} else {
		x.Scn = scn
	}
	r.Minute, ok = x.TmePrtLit('m')
	if ok {
		scn = x.Scn
	} else {
		x.Scn = scn
	}
	r.Second, ok = x.TmePrtLit('s')
	if ok {
		scn = x.Scn
	} else {
		x.Scn = scn
	}
	if r.Second.Lim == 0 && r.Minute.Lim == 0 && r.Hour.Lim == 0 && r.Day.Lim == 0 && r.Week.Lim == 0 && r.Month.Lim == 0 && r.Year.Lim == 0 {
		return r, false
	}
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' || (len(skpDsh) == 0 && x.Ch == '-') {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) BndLit() (r BndLit, ok bool) {
	r.Idx = x.Idx
	for unicode.IsDigit(x.Ch) {
		x.NextRune()
	}
	if x.Ch != '-' {
		return r, false
	}
	r.IdxTrm.Idx = r.Idx
	r.IdxTrm.Lim = x.Idx
	x.NextRune()
	r.LimTrm.Idx = x.Idx
	for unicode.IsDigit(x.Ch) {
		x.NextRune()
	}
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.LimTrm.Lim = x.Idx
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) FltRngLit() (r FltRngLit, ok bool) {
	r.Idx = x.Idx
	r.MinTrm, ok = x.FltLit(true)
	if !ok {
		return r, false
	}
	if x.Ch != '-' {
		return r, false
	}
	x.NextRune()
	r.MaxTrm, ok = x.FltLit()
	if !ok {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) TmeRngLit() (r TmeRngLit, ok bool) {
	r.Idx = x.Idx
	r.MinTrm, ok = x.TmeLit(true)
	if !ok {
		return r, false
	}
	if x.Ch != '-' {
		return r, false
	}
	x.NextRune()
	r.MaxTrm, ok = x.TmeLit()
	if !ok {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) StrsLit() (r StrsLit, ok bool) {
	if x.Ch != '[' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	scn := x.Scn
	for {
		x.SkpSpceCmnt()
		scn = x.Scn
		elm, ok := x.StrLit()
		if !ok {
			x.Scn = scn
			break
		}
		r.Elms = append(r.Elms, elm)
	}
	if len(r.Elms) == 0 {
		return r, false
	}
	x.SkpSpceCmnt()
	if x.End || x.Ch != ']' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) BolsLit() (r BolsLit, ok bool) {
	if x.Ch != '[' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	scn := x.Scn
	for {
		x.SkpSpceCmnt()
		scn = x.Scn
		elm, ok := x.BolLit()
		if !ok {
			x.Scn = scn
			break
		}
		r.Elms = append(r.Elms, elm)
	}
	if len(r.Elms) == 0 {
		return r, false
	}
	x.SkpSpceCmnt()
	if x.End || x.Ch != ']' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) FltsLit() (r FltsLit, ok bool) {
	if x.Ch != '[' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	scn := x.Scn
	for {
		x.SkpSpceCmnt()
		scn = x.Scn
		elm, ok := x.FltLit()
		if !ok {
			x.Scn = scn
			break
		}
		r.Elms = append(r.Elms, elm)
	}
	if len(r.Elms) == 0 {
		return r, false
	}
	x.SkpSpceCmnt()
	if x.End || x.Ch != ']' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) UntsLit() (r UntsLit, ok bool) {
	if x.Ch != '[' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	scn := x.Scn
	for {
		x.SkpSpceCmnt()
		scn = x.Scn
		elm, ok := x.UntLit()
		if !ok {
			x.Scn = scn
			break
		}
		r.Elms = append(r.Elms, elm)
	}
	if len(r.Elms) == 0 {
		return r, false
	}
	x.SkpSpceCmnt()
	if x.End || x.Ch != ']' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) IntsLit() (r IntsLit, ok bool) {
	if x.Ch != '[' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	scn := x.Scn
	for {
		x.SkpSpceCmnt()
		scn = x.Scn
		elm, ok := x.IntLit()
		if !ok {
			x.Scn = scn
			break
		}
		r.Elms = append(r.Elms, elm)
	}
	if len(r.Elms) == 0 {
		return r, false
	}
	x.SkpSpceCmnt()
	if x.End || x.Ch != ']' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) TmesLit() (r TmesLit, ok bool) {
	if x.Ch != '[' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	scn := x.Scn
	for {
		x.SkpSpceCmnt()
		scn = x.Scn
		elm, ok := x.TmeLit()
		if !ok {
			x.Scn = scn
			break
		}
		r.Elms = append(r.Elms, elm)
	}
	if len(r.Elms) == 0 {
		return r, false
	}
	x.SkpSpceCmnt()
	if x.End || x.Ch != ']' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) BndsLit() (r BndsLit, ok bool) {
	if x.Ch != '[' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	scn := x.Scn
	for {
		x.SkpSpceCmnt()
		scn = x.Scn
		elm, ok := x.BndLit()
		if !ok {
			x.Scn = scn
			break
		}
		r.Elms = append(r.Elms, elm)
	}
	if len(r.Elms) == 0 {
		return r, false
	}
	x.SkpSpceCmnt()
	if x.End || x.Ch != ']' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
func (x *Trmr) TmeRngsLit() (r TmeRngsLit, ok bool) {
	if x.Ch != '[' {
		return r, false
	}
	r.Idx = x.Idx
	x.NextRune()
	scn := x.Scn
	for {
		x.SkpSpceCmnt()
		scn = x.Scn
		elm, ok := x.TmeRngLit()
		if !ok {
			x.Scn = scn
			break
		}
		r.Elms = append(r.Elms, elm)
	}
	if len(r.Elms) == 0 {
		return r, false
	}
	x.SkpSpceCmnt()
	if x.End || x.Ch != ']' {
		return r, false
	}
	x.NextRune()
	if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {
		return r, false
	}
	r.Lim = x.Idx
	return r, true
}
