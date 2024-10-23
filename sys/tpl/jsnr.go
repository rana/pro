package tpl

import (
	"sys/k"
	"sys/tpl/atr"
)

type (
	FleJsnr struct {
		FleBse
	}
)

func (x *DirJsn) NewJsnr() (r *FleJsnr) {
	r = &FleJsnr{}
	x.Jsnr = r
	r.Name = k.Jsnr
	r.Pkg = x.Pkg
	r.StructPtr(r.Name, atr.Test) // Jsnr is a cfg term lexer
	r.AddFle(r)
	return r
}
func (x *FleJsnr) InitFld(s *Struct) {
	x.Import(_sys.Lng.Jsn.Trm)
	s.FldExt("trm.Trmr") // use Ext to avoid Ptr mod
}
func (x *FleJsnr) InitTypFn() {
	x.Test.Import(x)
	x.Import(_sys)
	x.Import(_sys.Lng.Jsn.Trm.Prs)
	x.SrchKey()
	x.SkpVal()
	x.Arr()
	x.ElmObj()
	x.ArrObjs()
	x.StrTmeLayout()
	x.StrTme()
	x.StrFlt()
	x.StrUnt()
	x.Unt()
}

func (x *FleJsnr) SrchKey() (r *TypFn) {
	r = x.TypFn("SrchKey")
	r.InPrmVariadic(String, "path")
	r.OutPrm(Bool)
	r.Add("if len(path) > 0 {")
	r.Add("x.SkpSpce()")
	r.Add("if x.Ch != '{' { // must start with lcrl")
	r.Add("return false")
	r.Add("}")
	r.Add("x.NextRune() // skip lcrl")
	r.Add("x.SkpSpce()")
	r.Add("key, ok := x.StrLit()")
	r.Add("for !x.End && ok {")
	r.Add("x.SkpSpce()")
	r.Add("if x.Ch != ':' { // cln must follow key")
	r.Add("return false")
	r.Add("}")
	r.Add("x.NextRune() // skip cln")
	r.Add("if path[0] == x.Txt[key.Idx+1:key.Lim-1] {")
	r.Add("x.SkpSpce()")
	r.Add("if len(path) == 1 { // found final key")
	r.Add("return true")
	r.Add("}")
	r.Add("return x.SrchKey(path[1:]...) // look for inr key")
	r.Add("}")
	r.Add("// continue search for key at current depth")
	r.Add("if !x.SkpVal() { // skip over current value; may have nested values")
	r.Add("return false")
	r.Add("}")
	r.Add("key, ok = x.StrLit()")
	r.Add("}")
	r.Add("}")
	r.Add("return false")
	return r
}
func (x *FleJsnr) SkpVal() (r *TypFn) {
	r = x.TypFn("SkpVal")
	r.OutPrm(Bool)
	r.Add("x.SkpSpce()")
	r.Add("switch {")
	r.Add("case x.Ch == '\"': // skip over str typ")
	r.Add("_, ok := x.StrLit()")
	r.Add("if !ok {")
	r.Add("return false")
	r.Add("}")
	r.Add("case x.Ch == '[': // skip over arr typs")
	r.Add("x.SkpSet('[', ']')")
	r.Add("case x.Ch == '{': // skip over obj")
	r.Add("x.SkpSet('{', '}')")
	r.Add("default: // skip over all non-str types")
	r.Add("for !x.End && x.Ch != ',' && x.Ch != '}' {")
	r.Add("x.SkpSpce()")
	r.Add("x.NextRune()")
	r.Add("}")
	r.Add("}")
	r.Add("x.SkpSpce()")
	r.Add("if x.Ch == ',' {")
	r.Add("x.NextRune()")
	r.Add("x.SkpSpce()")
	r.Add("}")
	r.Add("return true")
	return r
}

func (x *FleJsnr) PrsJsn(fle Fle) (r *TypFn) {
	typ, bse := fle.Typ(), fle.Typ().Bse()
	x.Import("sys/err")
	r = x.TypFn(typ.Title())
	r.InPrmVariadic(String, "path")
	r.OutPrm(typ, "r")
	r.Add("x.Reset(x.Txt) // enables multiple sequential calls")
	r.Add("if x.SrchKey(path...) {")
	r.Addf("%v, ok := x.%v()", bse.LitTrmJsn.Camel(), bse.LitTrmJsn.Title())
	r.Add("if !ok {")
	r.Addf("err.Panicf(\"Jsnr: invalid %v (path:%%q)\", sys.JoinPth(path...))", typ.Title())
	r.Add("}")
	r.Addf("return %v(%v, x.Txt)", bse.PrsTrmJsn.Ref(x), bse.LitTrmJsn.Camel())
	r.Add("}")
	r.Add("err.Panicf(\"Jsnr: invalid path (path:%q)\", sys.JoinPth(path...))")
	r.Add("return r")
	x.Test.Jsn(r, typ)
	return r
}
func (x *FleJsnr) Unt() (r *TypFn) {
	r = x.TypFn("Unt")
	r.InPrmVariadic(String, "path")
	r.OutPrm(_sys.Bsc.Unt)
	r.Add("return unt.Unt(x.Int(path...))")
	return r
}
func (x *FleJsnr) StrUnt() (r *TypFn) {
	x.Import("strconv")
	r = x.TypFn("StrUnt")
	r.InPrmVariadic(String, "path")
	r.OutPrm(_sys.Bsc.Unt)
	r.Add("txt := x.Str(path...).Unquo()")
	r.Add("v, er := strconv.ParseUint(txt, 10, 32)")
	r.Add("if er != nil {")
	r.Addf("err.Panicf(\"%v: failed to parse (txt:%%q err:%%q)\", txt, er)", x.Typ().Title())
	r.Add("}")
	r.Add("return unt.Unt(v)")
	return r
}
func (x *FleJsnr) StrFlt() (r *TypFn) {
	x.Import("strconv")
	r = x.TypFn("StrFlt")
	r.InPrmVariadic(String, "path")
	r.OutPrm(_sys.Bsc.Flt)
	r.Add("txt := x.Str(path...).Unquo()")
	r.Add("v, er := strconv.ParseFloat(txt, 32)")
	r.Add("if er != nil {")
	r.Addf("err.Panicf(\"%v: failed to parse (txt:%%q err:%%q)\", txt, er)", x.Typ().Title())
	r.Add("}")
	r.Add("return flt.Flt(v)")
	return r
}
func (x *FleJsnr) StrTmeLayout() (r *TypFn) {
	x.Import("time")
	r = x.TypFn("StrTmeLayout")
	r.InPrm(String, "layout")
	r.InPrmVariadic(String, "path")
	r.OutPrm(_sys.Bsc.Tme)
	r.Add("txt := x.Str(path...).Unquo()")
	r.Add("v, er := time.Parse(layout, txt)")
	r.Add("if er != nil {")
	r.Addf("err.Panicf(\"%v: failed to parse (txt:%%q err:%%q)\", txt, er)", x.Typ().Title())
	r.Add("}")
	r.Add("return tme.Time(v)")
	return r
}
func (x *FleJsnr) StrTme() (r *TypFn) {
	x.Import("time")
	r = x.TypFn("StrTme")
	r.InPrmVariadic(String, "path")
	r.OutPrm(_sys.Bsc.Tme)
	r.Add("txt := x.Str(path...).Unquo()")
	r.Add("v, er := time.Parse(time.RFC3339Nano, txt)")
	r.Add("if er != nil {")
	r.Addf("err.Panicf(\"%v: failed to parse (txt:%%q err:%%q)\", txt, er)", x.Typ().Title())
	r.Add("}")
	r.Add("return tme.Time(v)")
	return r
}
func (x *FleJsnr) Arr() (r *TypFn) {
	r = x.TypFn("Arr")
	r.InPrmVariadic(String, "path")
	r.OutPrm(_sys.Bsc.Bnd, "r")
	r.Add("x.Reset(x.Txt) // enables multiple sequential calls")
	r.Add("if x.SrchKey(path...) {")
	r.Add("x.SkpSpce()")
	r.Add("r.Idx = x.Idx")
	r.Add("x.SkpSet('[', ']')")
	r.Add("r.Lim = x.Idx")
	r.Add("return r")
	r.Add("}")
	r.Add("err.Panicf(\"Jsnr: invalid path (path:%q)\", sys.JoinPth(path...))")
	r.Add("return r")
	return r
}
func (x *FleJsnr) ElmObj() (r *TypFn) {
	r = x.TypFn("ElmObj")
	r.OutPrm(_sys.Bsc.Bnd, "r")
	r.Add("x.SkpSpce()")
	r.Add("r.Idx = x.Idx")
	r.Add("x.SkpSet('{', '}')")
	r.Add("r.Lim = x.Idx")
	r.Add("return r")
	return r
}
func (x *FleJsnr) ArrObjs() (r *TypFn) {
	r = x.TypFn("ArrObjs")
	r.InPrmVariadic(String, "path")
	r.OutPrmSlice(_sys.Bsc.Bnd, "r")
	r.Add("txt := x.Txt")
	r.Add("arr := x.Arr(path...)")
	r.Add("arrTxt := x.Txt[arr.Idx+1 : arr.Lim]")
	r.Add("// fmt.Println(\"arrTxt:\", arrTxt)")
	r.Add("x.Reset(arrTxt)")
	r.Add("for !x.End && x.Ch != ']' {")
	r.Add("x.SkpSpce()")
	r.Addf("var b %v", r.OutTyp().Ref(x))
	r.Add("b.Idx = x.Idx + arr.Idx + 1")
	r.Add("x.SkpSet('{', '}')")
	r.Add("b.Lim = x.Idx + arr.Idx + 1")
	r.Add("r = append(r, b)")
	r.Add("// fmt.Println(\"---\", b, arrTxt[b.Idx:b.Lim])")
	r.Add("x.SkpSpce()")
	r.Add("if x.Ch == ']' {")
	r.Add("break")
	r.Add("}")
	r.Add("x.NextRune() // skp comma")
	r.Add("}")
	r.Add("x.Reset(txt)")
	r.Add("return r")
	return r
}
