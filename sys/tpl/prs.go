package tpl

import (
	"sys/k"
)

type (
	FlePrs struct {
		FleBse
		TmePrt *PkgFn
	}
)

func (x *DirTrm) NewPrs() (r *FlePrs) {
	r = &FlePrs{}
	r.Name = k.Prs
	r.Pkg = x.Pkg.New(r.Name)
	r.AddFle(r)
	r.Test = r.FleTest(r.Name) // manual call because of no Typ associated to Fle
	r.Test.Import("testing")
	r.Test.Import("sys/tst")
	r.Test.Import(_sys.Lng.Pro.Trm)
	r.Test.Import(r)
	r.ImportAlias(_sys.Bsc.Int)
	return r
}
func (x *FlePrs) PrsTypTrm(typ Typ, prs func(r *PkgFn, f *FlePrs)) {
	bse := typ.Bse()
	x.Import(bse.Pkg.Pth)
	var r *PkgFn
	r = x.PkgFnf("%vTrm", bse.PrefixTitle())
	bse.PrsTrm = r
	r.InPrm(bse.LitTrm, "trm")
	r.InPrm(String, "txt")
	r.OutPrm(bse, "r")
	prs(r, x)
	// test
	x.Test.Import(bse.Pkg.Pth)
	x.Test.PrsTrm(r, bse, bse.LitTrm)
	x.PrsTypTxt(bse)
}
func (x *FlePrs) PrsTypTxt(bse *TypBse) (r *PkgFn) {
	x.Import("sys/err")
	r = x.PkgFnf("%v", bse.PrefixTitle())
	bse.PrsTxt = r
	r.InPrm(String, "txt")
	r.OutPrm(bse, "r")
	r.Add("var trmr trm.Trmr")
	r.Add("trmr.Reset(txt)")
	r.Addf("%v, ok := trmr.%v()", bse.LitTrm.Camel(), bse.LitTrm.Title())
	r.Add("if !ok {")
	r.Addf("err.Panicf(\"%v: failed to parse (txt:%%q)\", txt)", bse.Title())
	r.Add("}")
	r.Addf("return %v(%v, txt)", bse.PrsTrm.Ref(x), bse.LitTrm.Camel())
	// test
	x.Test.PrsTxt(r, bse)
	return r
}

func (x *FlePrs) PrsTmePrtTrm() (r *PkgFn) {
	r = x.PkgFn("TmePrtTrm")
	x.TmePrt = r
	r.InPrm(_sys.Lng.Pro.Trm.Trmr.TmePrt, "trm")
	r.InPrm(String, "txt")
	r.OutPrm(_sys.Bsc.Tme, "r")
	r.Add("if trm.Lim == 0 {")
	r.Add("return 0")
	r.Add("}")
	r.Add("lit := txt[trm.Idx : trm.Lim-1] // trim suffix ch: y,n,w,d,h,m,s")
	r.Add("neg := lit[0] == '-'        		 // optional minus")
	r.Add("if neg {")
	r.Add("lit = lit[1:]")
	r.Add("}")
	r.Addf("mag := %v(1)", r.OutTyp().Ref(x))
	r.Add("for n := len(lit) - 1; n > -1; n-- {")
	r.Addf("r += mag * %v(lit[n]-'0')", r.OutTyp().Ref(x))
	r.Addf("mag *= %v(10)", r.OutTyp().Ref(x))
	r.Add("}")
	r.Add("if neg {")
	r.Add("r = -r")
	r.Add("}")
	r.Add("return r")
	return r
}

func (x *FlePrs) PrsArr(arr *Arr) (r *PkgFn) {
	x.Import(arr.Pkg.Pth)
	r = x.PkgFnf("%vTrm", arr.PrefixTitle())
	arr.PrsTrm = r
	r.InPrm(arr.LitTrm, "trm")
	r.InPrm(String, "txt")
	r.OutPrm(arr)
	r.Addf("r := %v{}", arr.Full())
	r.Add("for _, elm := range trm.Elms {")
	r.Addf("r.Push(%v(elm, txt))", arr.Elm.PrsTrm.Ref(x))
	r.Add("}")
	r.Add("return &r")
	// test
	if Opt.Is(TestOpt, ArrOpt) {
		x.Test.Import(arr.Pkg.Pth)
		x.Test.PrsTrm(r, arr, arr.LitTrm)
	}
	x.PrsArrTxt(arr)
	return r
}
func (x *FlePrs) PrsArrTxt(arr *Arr) (r *PkgFn) {
	r = x.PkgFnf("%v", arr.PrefixTitle())
	arr.PrsTxt = r
	r.InPrm(String, "txt")
	r.OutPrm(arr, "r")
	r.Add("var trmr trm.Trmr")
	r.Add("trmr.Reset(txt)")
	r.Addf("%v, ok := trmr.%v()", arr.LitTrm.Camel(), arr.LitTrm.Title())
	r.Add("if !ok {")
	r.Addf("err.Panicf(\"%v: failed to parse (txt:%%q)\", txt)", arr.Title())
	r.Add("}")
	r.Addf("return %v(%v, txt)", arr.PrsTrm.Ref(x), arr.LitTrm.Camel())
	// test
	x.Test.PrsTxt(r, arr)
	return r
}
