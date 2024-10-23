package tpl

import (
	"fmt"
	"sys"
	"sys/k"
	"sys/tpl/atr"
	"unicode/utf8"
)

type (
	FleJsnTrmr struct {
		FleBse
		Pkgs
		Trm  *DirJsnTrm
		Trms map[string]interface{}
		Spce *Struct
	}
)

func (x *DirJsnTrm) NewTrmr() (r *FleJsnTrmr) {
	r = &FleJsnTrmr{Trm: x}
	x.Trmr = r
	r.Name = k.Trmr
	r.Pkg = x.Pkg
	r.Trmr()
	r.AddFle(r)
	r.Trms = make(map[string]interface{})
	return r
}
func (x *FleJsnTrmr) Trmr() (r *Struct) {
	x.Import(_sys.Lng.Scn)
	r = x.StructPtr(k.Trmr, atr.Test) // Trmr is a term lexer.
	r.FldExt("scn.Scnr")              // use Ext to avoid Ptr mod
	return r
}
func (x *FleJsnTrmr) InitTypFn() {
	x.SkpSpce()
}
func (x *FleJsnTrmr) StructLit(name string) (r *Struct) {
	r = x.Structf("%vLit", atr.None, name)
	r.FldTyp(_sys.Bsc.Bnd.Typ())
	return r
}
func (x *FleJsnTrmr) StructLitf(format string, args ...interface{}) (r *Struct) {
	r = x.Structf("%vLit", atr.None, fmt.Sprintf(format, args...))
	r.FldTyp(_sys.Bsc.Bnd.Typ())
	return r
}
func (x *FleJsnTrmr) StructArrLit(name string, elmLit *Struct) (r *Struct) {
	r = x.Structf("%vLit", atr.None, name)
	r.FldTyp(_sys.Bsc.Bnd.Typ())
	r.FldSlice("Elms", elmLit)
	return r
}
func (x *FleJsnTrmr) MemLit(lit *Struct) (r *TypFn) {
	r = x.TypFn(lit.Name)
	r.OutPrm(lit, "r")
	r.OutPrm(Bool, "ok")
	return r
}
func (x *FleJsnTrmr) LexTrm(trm string) (r *TypFn) {
	trm = sys.Camel(trm)
	if _, exists := x.Trms[trm]; !exists {
		x.Trms[trm] = trm
		title := trm
		if trm == k.False || trm == k.True {
			title += "Lit"
		}
		r = x.TypFnf(title)
		r.OutPrm(_sys.Bsc.Bnd.Typ(), "r")
		r.OutPrm(Bool, "ok")
		ch, size, idx := rune(0), 0, 0
		for idx < len(trm) {
			ch, size = utf8.DecodeRuneInString(trm[idx:])
			r.Addf("if x.Ch != %q {", ch)
			r.Add("return r, false")
			r.Add("}")
			if idx == 0 {
				r.Add("r.Idx = x.Idx")
			}
			r.Add("x.NextRune()")
			idx += size
		}
		r.Add("if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' {")
		r.Add("return r, false")
		r.Add("}")
		r.Add("r.Lim = x.Idx")
		r.Add("return r, true")
		x.Test.Import(x)
		x.Test.Import(_sys.Bsc.Bnd)
		x.Test.Trm(title, _sys.Bsc.Bnd.Typ().Ref(x), []string{trm})
	}
	return r
}
func (x *FleJsnTrmr) CallLit(fn, trm *TypFn, lit *Struct) {
	fn.Addf("%v, ok := x.%v()", trm.Camel(), trm.Name)
	fn.Add("if ok {")
	fn.Addf("return %v{Bnd: %v}, true", lit.Title(), trm.Camel())
	fn.Add("}")
	fn.Add("x.Scn = scn")
}

func (x *FleJsnTrmr) SkpSpce() (r *TypFn) {
	x.Import("unicode")
	r = x.TypFn("SkpSpce")
	r.Add("for unicode.IsSpace(x.Ch) {")
	r.Add("x.NextRune()")
	r.Add("}")
	return r
}
