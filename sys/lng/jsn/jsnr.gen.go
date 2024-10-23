package jsn

import (
	"strconv"
	"sys"
	"sys/bsc/bnd"
	"sys/bsc/bol"
	"sys/bsc/flt"
	"sys/bsc/int"
	"sys/bsc/str"
	"sys/bsc/tme"
	"sys/bsc/unt"
	"sys/err"
	"sys/lng/jsn/trm"
	"sys/lng/jsn/trm/prs"
	"time"
)

type (
	Jsnr struct {
		trm.Trmr
	}
)

func (x *Jsnr) SrchKey(path ...string) bool {
	if len(path) > 0 {
		x.SkpSpce()
		if x.Ch != '{' { // must start with lcrl
			return false
		}
		x.NextRune() // skip lcrl
		x.SkpSpce()
		key, ok := x.StrLit()
		for !x.End && ok {
			x.SkpSpce()
			if x.Ch != ':' { // cln must follow key
				return false
			}
			x.NextRune() // skip cln
			if path[0] == x.Txt[key.Idx+1:key.Lim-1] {
				x.SkpSpce()
				if len(path) == 1 { // found final key
					return true
				}
				return x.SrchKey(path[1:]...) // look for inr key
			}
			// continue search for key at current depth
			if !x.SkpVal() { // skip over current value; may have nested values
				return false
			}
			key, ok = x.StrLit()
		}
	}
	return false
}
func (x *Jsnr) SkpVal() bool {
	x.SkpSpce()
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
		for !x.End && x.Ch != ',' && x.Ch != '}' {
			x.SkpSpce()
			x.NextRune()
		}
	}
	x.SkpSpce()
	if x.Ch == ',' {
		x.NextRune()
		x.SkpSpce()
	}
	return true
}
func (x *Jsnr) Arr(path ...string) (r bnd.Bnd) {
	x.Reset(x.Txt) // enables multiple sequential calls
	if x.SrchKey(path...) {
		x.SkpSpce()
		r.Idx = x.Idx
		x.SkpSet('[', ']')
		r.Lim = x.Idx
		return r
	}
	err.Panicf("Jsnr: invalid path (path:%q)", sys.JoinPth(path...))
	return r
}
func (x *Jsnr) ElmObj() (r bnd.Bnd) {
	x.SkpSpce()
	r.Idx = x.Idx
	x.SkpSet('{', '}')
	r.Lim = x.Idx
	return r
}
func (x *Jsnr) ArrObjs(path ...string) (r []bnd.Bnd) {
	txt := x.Txt
	arr := x.Arr(path...)
	arrTxt := x.Txt[arr.Idx+1 : arr.Lim]
	// fmt.Println("arrTxt:", arrTxt)
	x.Reset(arrTxt)
	for !x.End && x.Ch != ']' {
		x.SkpSpce()
		var b bnd.Bnd
		b.Idx = x.Idx + arr.Idx + 1
		x.SkpSet('{', '}')
		b.Lim = x.Idx + arr.Idx + 1
		r = append(r, b)
		// fmt.Println("---", b, arrTxt[b.Idx:b.Lim])
		x.SkpSpce()
		if x.Ch == ']' {
			break
		}
		x.NextRune() // skp comma
	}
	x.Reset(txt)
	return r
}
func (x *Jsnr) StrTmeLayout(layout string, path ...string) tme.Tme {
	txt := x.Str(path...).Unquo()
	v, er := time.Parse(layout, txt)
	if er != nil {
		err.Panicf("Jsnr: failed to parse (txt:%q err:%q)", txt, er)
	}
	return tme.Time(v)
}
func (x *Jsnr) StrTme(path ...string) tme.Tme {
	txt := x.Str(path...).Unquo()
	v, er := time.Parse(time.RFC3339Nano, txt)
	if er != nil {
		err.Panicf("Jsnr: failed to parse (txt:%q err:%q)", txt, er)
	}
	return tme.Time(v)
}
func (x *Jsnr) StrFlt(path ...string) flt.Flt {
	txt := x.Str(path...).Unquo()
	v, er := strconv.ParseFloat(txt, 32)
	if er != nil {
		err.Panicf("Jsnr: failed to parse (txt:%q err:%q)", txt, er)
	}
	return flt.Flt(v)
}
func (x *Jsnr) StrUnt(path ...string) unt.Unt {
	txt := x.Str(path...).Unquo()
	v, er := strconv.ParseUint(txt, 10, 32)
	if er != nil {
		err.Panicf("Jsnr: failed to parse (txt:%q err:%q)", txt, er)
	}
	return unt.Unt(v)
}
func (x *Jsnr) Unt(path ...string) unt.Unt { return unt.Unt(x.Int(path...)) }
func (x *Jsnr) Str(path ...string) (r str.Str) {
	x.Reset(x.Txt) // enables multiple sequential calls
	if x.SrchKey(path...) {
		strLit, ok := x.StrLit()
		if !ok {
			err.Panicf("Jsnr: invalid Str (path:%q)", sys.JoinPth(path...))
		}
		return prs.StrTrm(strLit, x.Txt)
	}
	err.Panicf("Jsnr: invalid path (path:%q)", sys.JoinPth(path...))
	return r
}
func (x *Jsnr) Bol(path ...string) (r bol.Bol) {
	x.Reset(x.Txt) // enables multiple sequential calls
	if x.SrchKey(path...) {
		bolLit, ok := x.BolLit()
		if !ok {
			err.Panicf("Jsnr: invalid Bol (path:%q)", sys.JoinPth(path...))
		}
		return prs.BolTrm(bolLit, x.Txt)
	}
	err.Panicf("Jsnr: invalid path (path:%q)", sys.JoinPth(path...))
	return r
}
func (x *Jsnr) Flt(path ...string) (r flt.Flt) {
	x.Reset(x.Txt) // enables multiple sequential calls
	if x.SrchKey(path...) {
		fltLit, ok := x.FltLit()
		if !ok {
			err.Panicf("Jsnr: invalid Flt (path:%q)", sys.JoinPth(path...))
		}
		return prs.FltTrm(fltLit, x.Txt)
	}
	err.Panicf("Jsnr: invalid path (path:%q)", sys.JoinPth(path...))
	return r
}
func (x *Jsnr) Int(path ...string) (r int.Int) {
	x.Reset(x.Txt) // enables multiple sequential calls
	if x.SrchKey(path...) {
		intLit, ok := x.IntLit()
		if !ok {
			err.Panicf("Jsnr: invalid Int (path:%q)", sys.JoinPth(path...))
		}
		return prs.IntTrm(intLit, x.Txt)
	}
	err.Panicf("Jsnr: invalid path (path:%q)", sys.JoinPth(path...))
	return r
}
