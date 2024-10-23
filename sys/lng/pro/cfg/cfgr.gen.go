package cfg

import (
	"sys"
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
	"sys/err"
	"sys/lng/pro/trm"
	"sys/lng/pro/trm/prs"
	"unicode"
)

type (
	Cfgr struct {
		trm.Trmr
	}
)

func (x *Cfgr) SrchKey(path ...string) bool {
	if len(path) > 0 {
		x.SkpSpceCmnt()
		if x.Ch != '{' { // must start with lcrl
			return false
		}
		x.SkpSpceCmnt()
		x.NextRune() // skip lcrl
		x.SkpSpceCmnt()
		idn, ok := x.IdnLit()
		for !x.End && ok {
			x.SkpSpceCmnt()
			if x.Ch != ':' { // cln must follow idn
				return false
			}
			x.NextRune() // skip cln
			if path[0] == x.Txt[idn.Idx:idn.Lim] {
				x.SkpSpceCmnt()
				if len(path) == 1 { // found final key
					return true
				}
				return x.SrchKey(path[1:]...) // look for inr key
			}
			// continue search for key at current depth
			if !x.SkpVal() { // skip over current value; may have nested values
				return false
			}
			idn, ok = x.IdnLit()
		}
	}
	return false
}
func (x *Cfgr) SkpVal() bool {
	x.SkpSpceCmnt()
	switch {
	case x.Ch == '"': // skip over str typ
		_, ok := x.StrLit()
		if !ok {
			return false
		}
	case x.Ch == '[': // skip over arr typs
		x.SkpSet('[', ']')
	case x.Ch == '{': // skip over obj
		x.SkpSet('{', '}')
	default: // skip over all non-str types
		for !x.End && !unicode.IsSpace(x.Ch) {
			x.NextRune()
		}
	}
	x.SkpSpceCmnt()
	return true
}
func (x *Cfgr) Str(path ...string) (r str.Str) {
	x.Reset(x.Txt) // enables multiple sequential calls
	if x.SrchKey(path...) {
		strLit, ok := x.StrLit()
		if !ok {
			err.Panicf("Cfgr: invalid Str (path:%q)", sys.JoinPth(path...))
		}
		return prs.StrTrm(strLit, x.Txt)
	}
	err.Panicf("Cfgr: invalid path (path:%q)", sys.JoinPth(path...))
	return r
}
func (x *Cfgr) Bol(path ...string) (r bol.Bol) {
	x.Reset(x.Txt) // enables multiple sequential calls
	if x.SrchKey(path...) {
		bolLit, ok := x.BolLit()
		if !ok {
			err.Panicf("Cfgr: invalid Bol (path:%q)", sys.JoinPth(path...))
		}
		return prs.BolTrm(bolLit, x.Txt)
	}
	err.Panicf("Cfgr: invalid path (path:%q)", sys.JoinPth(path...))
	return r
}
func (x *Cfgr) Flt(path ...string) (r flt.Flt) {
	x.Reset(x.Txt) // enables multiple sequential calls
	if x.SrchKey(path...) {
		fltLit, ok := x.FltLit()
		if !ok {
			err.Panicf("Cfgr: invalid Flt (path:%q)", sys.JoinPth(path...))
		}
		return prs.FltTrm(fltLit, x.Txt)
	}
	err.Panicf("Cfgr: invalid path (path:%q)", sys.JoinPth(path...))
	return r
}
func (x *Cfgr) Unt(path ...string) (r unt.Unt) {
	x.Reset(x.Txt) // enables multiple sequential calls
	if x.SrchKey(path...) {
		untLit, ok := x.UntLit()
		if !ok {
			err.Panicf("Cfgr: invalid Unt (path:%q)", sys.JoinPth(path...))
		}
		return prs.UntTrm(untLit, x.Txt)
	}
	err.Panicf("Cfgr: invalid path (path:%q)", sys.JoinPth(path...))
	return r
}
func (x *Cfgr) Int(path ...string) (r int.Int) {
	x.Reset(x.Txt) // enables multiple sequential calls
	if x.SrchKey(path...) {
		intLit, ok := x.IntLit()
		if !ok {
			err.Panicf("Cfgr: invalid Int (path:%q)", sys.JoinPth(path...))
		}
		return prs.IntTrm(intLit, x.Txt)
	}
	err.Panicf("Cfgr: invalid path (path:%q)", sys.JoinPth(path...))
	return r
}
func (x *Cfgr) Tme(path ...string) (r tme.Tme) {
	x.Reset(x.Txt) // enables multiple sequential calls
	if x.SrchKey(path...) {
		tmeLit, ok := x.TmeLit()
		if !ok {
			err.Panicf("Cfgr: invalid Tme (path:%q)", sys.JoinPth(path...))
		}
		return prs.TmeTrm(tmeLit, x.Txt)
	}
	err.Panicf("Cfgr: invalid path (path:%q)", sys.JoinPth(path...))
	return r
}
func (x *Cfgr) Bnd(path ...string) (r bnd.Bnd) {
	x.Reset(x.Txt) // enables multiple sequential calls
	if x.SrchKey(path...) {
		bndLit, ok := x.BndLit()
		if !ok {
			err.Panicf("Cfgr: invalid Bnd (path:%q)", sys.JoinPth(path...))
		}
		return prs.BndTrm(bndLit, x.Txt)
	}
	err.Panicf("Cfgr: invalid path (path:%q)", sys.JoinPth(path...))
	return r
}
func (x *Cfgr) FltRng(path ...string) (r flt.Rng) {
	x.Reset(x.Txt) // enables multiple sequential calls
	if x.SrchKey(path...) {
		fltRngLit, ok := x.FltRngLit()
		if !ok {
			err.Panicf("Cfgr: invalid FltRng (path:%q)", sys.JoinPth(path...))
		}
		return prs.FltRngTrm(fltRngLit, x.Txt)
	}
	err.Panicf("Cfgr: invalid path (path:%q)", sys.JoinPth(path...))
	return r
}
func (x *Cfgr) TmeRng(path ...string) (r tme.Rng) {
	x.Reset(x.Txt) // enables multiple sequential calls
	if x.SrchKey(path...) {
		tmeRngLit, ok := x.TmeRngLit()
		if !ok {
			err.Panicf("Cfgr: invalid TmeRng (path:%q)", sys.JoinPth(path...))
		}
		return prs.TmeRngTrm(tmeRngLit, x.Txt)
	}
	err.Panicf("Cfgr: invalid path (path:%q)", sys.JoinPth(path...))
	return r
}
func (x *Cfgr) Strs(path ...string) (r *strs.Strs) {
	x.Reset(x.Txt) // enables multiple sequential calls
	if x.SrchKey(path...) {
		strsLit, ok := x.StrsLit()
		if !ok {
			err.Panicf("Cfgr: invalid Strs (path:%q)", sys.JoinPth(path...))
		}
		return prs.StrsTrm(strsLit, x.Txt)
	}
	err.Panicf("Cfgr: invalid path (path:%q)", sys.JoinPth(path...))
	return r
}
func (x *Cfgr) Bols(path ...string) (r *bols.Bols) {
	x.Reset(x.Txt) // enables multiple sequential calls
	if x.SrchKey(path...) {
		bolsLit, ok := x.BolsLit()
		if !ok {
			err.Panicf("Cfgr: invalid Bols (path:%q)", sys.JoinPth(path...))
		}
		return prs.BolsTrm(bolsLit, x.Txt)
	}
	err.Panicf("Cfgr: invalid path (path:%q)", sys.JoinPth(path...))
	return r
}
func (x *Cfgr) Flts(path ...string) (r *flts.Flts) {
	x.Reset(x.Txt) // enables multiple sequential calls
	if x.SrchKey(path...) {
		fltsLit, ok := x.FltsLit()
		if !ok {
			err.Panicf("Cfgr: invalid Flts (path:%q)", sys.JoinPth(path...))
		}
		return prs.FltsTrm(fltsLit, x.Txt)
	}
	err.Panicf("Cfgr: invalid path (path:%q)", sys.JoinPth(path...))
	return r
}
func (x *Cfgr) Unts(path ...string) (r *unts.Unts) {
	x.Reset(x.Txt) // enables multiple sequential calls
	if x.SrchKey(path...) {
		untsLit, ok := x.UntsLit()
		if !ok {
			err.Panicf("Cfgr: invalid Unts (path:%q)", sys.JoinPth(path...))
		}
		return prs.UntsTrm(untsLit, x.Txt)
	}
	err.Panicf("Cfgr: invalid path (path:%q)", sys.JoinPth(path...))
	return r
}
func (x *Cfgr) Ints(path ...string) (r *ints.Ints) {
	x.Reset(x.Txt) // enables multiple sequential calls
	if x.SrchKey(path...) {
		intsLit, ok := x.IntsLit()
		if !ok {
			err.Panicf("Cfgr: invalid Ints (path:%q)", sys.JoinPth(path...))
		}
		return prs.IntsTrm(intsLit, x.Txt)
	}
	err.Panicf("Cfgr: invalid path (path:%q)", sys.JoinPth(path...))
	return r
}
func (x *Cfgr) Tmes(path ...string) (r *tmes.Tmes) {
	x.Reset(x.Txt) // enables multiple sequential calls
	if x.SrchKey(path...) {
		tmesLit, ok := x.TmesLit()
		if !ok {
			err.Panicf("Cfgr: invalid Tmes (path:%q)", sys.JoinPth(path...))
		}
		return prs.TmesTrm(tmesLit, x.Txt)
	}
	err.Panicf("Cfgr: invalid path (path:%q)", sys.JoinPth(path...))
	return r
}
func (x *Cfgr) Bnds(path ...string) (r *bnds.Bnds) {
	x.Reset(x.Txt) // enables multiple sequential calls
	if x.SrchKey(path...) {
		bndsLit, ok := x.BndsLit()
		if !ok {
			err.Panicf("Cfgr: invalid Bnds (path:%q)", sys.JoinPth(path...))
		}
		return prs.BndsTrm(bndsLit, x.Txt)
	}
	err.Panicf("Cfgr: invalid path (path:%q)", sys.JoinPth(path...))
	return r
}
func (x *Cfgr) TmeRngs(path ...string) (r *tme.Rngs) {
	x.Reset(x.Txt) // enables multiple sequential calls
	if x.SrchKey(path...) {
		tmeRngsLit, ok := x.TmeRngsLit()
		if !ok {
			err.Panicf("Cfgr: invalid TmeRngs (path:%q)", sys.JoinPth(path...))
		}
		return prs.TmeRngsTrm(tmeRngsLit, x.Txt)
	}
	err.Panicf("Cfgr: invalid path (path:%q)", sys.JoinPth(path...))
	return r
}
