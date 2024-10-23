package tpl

import "sys/k"

type (
	FleJsnPrs struct {
		FleBse
	}
)

func (x *DirJsnTrm) NewPrs() (r *FleJsnPrs) {
	r = &FleJsnPrs{}
	x.Prs = r
	r.Name = k.Prs
	r.Pkg = x.Pkg.New(r.Name)
	r.AddFle(r)
	r.Test = r.FleTest(r.Name) // manual call because of no Typ associated to Fle
	r.Test.Import("testing")
	r.Test.Import("sys/tst")
	r.Test.Import("sys/lng/jsn/trm")
	r.Test.Import(r)
	return r
}
func (x *FleJsnPrs) PrsTypTrm(typ Typ, prs func(r *PkgFn, f *FleJsnPrs)) {
	bse := typ.Bse()
	r := x.PkgFnf("%vTrm", bse.Title())
	bse.PrsTrmJsn = r
	r.InPrm(bse.LitTrmJsn, "trm")
	r.InPrm(String, "txt")
	r.OutPrm(bse, "r")
	prs(r, x)
	// test
	x.Test.Import(bse.Pkg)
	x.Test.PrsTrm(r, bse, bse.LitTrmJsn, true)
	x.PrsTypTxt(bse)
}
func (x *FleJsnPrs) PrsTypTxt(bse *TypBse) (r *PkgFn) {
	x.Import("sys/err")
	r = x.PkgFnf("%vTxt", bse.Title())
	bse.PrsTxt = r
	r.InPrm(String, "txt")
	r.OutPrm(bse, "r")
	r.Add("var trmr trm.Trmr")
	r.Add("trmr.Reset(txt)")
	r.Addf("%v, ok := trmr.%v()", bse.LitTrmJsn.Camel(), bse.LitTrmJsn.Title())
	r.Add("if !ok {")
	r.Addf("err.Panicf(\"%v: failed to parse (txt:%%q)\", txt)", bse.Title())
	r.Add("}")
	r.Addf("return %v(%v, txt)", bse.PrsTrmJsn.Ref(x), bse.LitTrmJsn.Camel())
	// test
	x.Test.PrsTxt(r, bse, true)
	return r
}
