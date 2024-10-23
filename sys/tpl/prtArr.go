package tpl

import (
	"fmt"
	"sys/k"
	"sys/tpl/atr"
)

type (
	PrtArr struct {
		PrtBse
		Arr *Arr
	}
)

func NewArr(elmFle Fle, p *PrtArr, pkg ...*Pkg) (r *FleBse) {
	r = &FleBse{}
	r.Name = fmt.Sprintf("%vs", elmFle.Typ().Camel())
	if len(pkg) != 0 {
		r.Pkg = pkg[0]
	} else {
		r.Pkg = elmFle.Bse().Pkg.NewFromPrnt(r.Name)
	}
	p.Arr = r.Typs.Arr(r.Name, elmFle, r.Pkg)
	return r
}

func (x *PrtArr) InitPrtTyp() {
	if x.Arr.Elm.Pkg.Name == k.Int {
		x.f.ImportAlias(x.Arr.Elm)
	}
}
func (x *PrtArr) InitPrtTrm(trmr *FleTrmr) {
	bse := x.f.Typ().Bse()
	if bse.IsTrm() {
		if bse.IsLit() {
			trmr.LexArrLit(x.Arr)
			_sys.Lng.Pro.Trm.Prs.PrsArr(x.Arr)
			if x.t.IsCfg() {
				x.f.InitPrsCfg()
			}
			x.f.InitTrm(bse, trmr)
		}
	}
}

func (x *PrtArr) InitPrtPkgFn() {
	x.Arr.New = x.new()
	x.Arr.Make = x.make()
	x.Arr.MakeEmp = x.makeEmp()
	// sys.Log("-", x.t.Name, "IsTestXpr", x.Arr.IsTestXpr())
	if !x.Arr.IsTestXpr() {
		for _, fn := range x.f.PkgFns {
			fn.Atr = fn.Atr &^ atr.TestXpr
		}
	}
	if !x.Arr.IsTestAct() {
		for _, fn := range x.f.PkgFns {
			fn.Atr = fn.Atr &^ atr.TestAct
		}
	}
}
func (x *PrtArr) new() (r *PkgFn) {
	r = x.f.PkgFn("New", true)
	r.InPrmVariadic(x.Arr.Elm, "vs")
	r.OutPrm(x.f.Typ())
	r.Addf("r := %v(vs)", x.t.Title())
	r.Add("return &r")
	return r
}
func (x *PrtArr) make() (r *PkgFn) {
	r = x.f.PkgFn("Make", true)
	r.InPrm(_sys.Bsc.Unt, "cap")
	r.OutPrm(x.f.Typ())
	r.Addf("r := make(%v, cap)", x.t.Title())
	r.Add("return &r")
	return r
}
func (x *PrtArr) makeEmp() (r *PkgFn) {
	r = x.f.PkgFn("MakeEmp", true)
	r.InPrm(_sys.Bsc.Unt, "cap")
	r.OutPrm(x.f.Typ())
	r.Addf("r := make(%v, 0, cap)", x.t.Title())
	r.Add("return &r")
	return r
}
func (x *PrtArr) InitPrtTypFn() {
	x.ok()
	x.cnt()
	x.cpy()
	x.clr()
	x.rand()
	x.mrg()
	x.push()
	x.pop()
	x.que()
	x.dque()
	x.ins()
	x.upd()
	x.del()
	x.at()
	x.in()
	x.inBnd()
	x.from()
	x.to()
	x.fst()
	x.mdl()
	x.lst()
	x.fstIdx()
	x.mdlIdx()
	x.lstIdx()
	x.rev()
	if !x.Arr.IsTestXpr() {
		for _, fn := range x.t.TypFns {
			fn.Atr = fn.Atr &^ atr.TestXpr
		}
	}
	if !x.Arr.IsTestAct() {
		for _, fn := range x.t.TypFns {
			fn.Atr = fn.Atr &^ atr.TestAct
		}
	}
}

func (x *PrtArr) ok() (r *TypFn) {
	r = x.f.TypFna("Ok", atr.None)
	r.OutPrm(_sys.Bsc.Bol)
	r.Add("return len(*x) != 0")
	// test
	if x.f.Test != nil && Opt.IsArr() {
		r.T.Addf("expected := %v(len(*exCpy) != 0)", r.OutTyp().Ref(x.f.Test))
	}
	return r
}
func (x *PrtArr) cnt() (r *TypFn) {
	r = x.f.TypFn("Cnt")
	r.OutPrm(_sys.Bsc.Unt)
	r.Add("if x == nil {")
	r.Add("return 0")
	r.Add("}")
	r.Addf("return %v(len(*x))", _sys.Bsc.Unt.Typ().Ref(x.f))
	// test
	if x.f.Test != nil && Opt.IsArr() {
		r.T.Addf("expected := %v(len(*exCpy))", _sys.Bsc.Unt.Typ().Ref(x.f.Test))
		if x.t.IsAna() {
			r.T.Addf("tst.%vEql(t, exCpy, axCpy)", r.Rxr.Typ.PrefixTitle())
		} else {
			r.T.Addf("tst.%vEql(t, exCpy, axCpy)", r.Rxr.Typ.PrefixTitle())
		}
	}
	return r
}
func (x *PrtArr) cpy() (r *TypFn) {
	r = x.f.TypFn("Cpy")
	r.OutPrm(x.f.Typ())
	r.Addf("r := make(%v, len(*x))", x.t.Title())
	r.Add("copy(r, *x)")
	r.Add("return &r")
	// test
	if x.f.Test != nil && Opt.IsArr() {
		r.T.Manual = true // manual since auto-gen uses Cpy()
		r.T.Addf("axCpy := make(%v, len(*x))", x.t.Full())
		r.T.Add("copy(axCpy, *x)")
		r.T.Add("actual := axCpy.Cpy()")
		r.T.Addf("exCpy := make(%v, len(*x))", x.t.Full())
		r.T.Add("copy(exCpy, *x)")
		r.T.Add("expected := &exCpy")
		if x.t.IsAna() {
			r.T.Addf("tst.%vEql(t, expected, actual)", r.OutTyp().PrefixTitle())
		} else {
			r.T.Addf("tst.%vEql(t, expected, actual)", r.OutTyp().PrefixTitle())
		}
	}
	return r
}
func (x *PrtArr) clr() (r *TypFn) {
	r = x.f.TypFn("Clr")
	r.OutPrm(x.f.Typ())
	r.Add("*x = (*x)[:0]")
	r.Add("return x")
	// test
	if x.f.Test != nil && Opt.IsArr() {
		r.T.Add("*exCpy = (*exCpy)[:0]")
		r.T.Add("expected := exCpy")
		r.T.Add("tst.IntegerZero(t, len(*actual))")
	}
	return r
}
func (x *PrtArr) rand() (r *TypFn) { // randomize element positions
	x.f.Import("math/rand")
	x.f.Import("time")
	r = x.f.TypFna(k.Rand, atr.Lng)
	r.OutPrm(x.f.Typ())
	r.Add("r := rand.New(rand.NewSource(time.Now().Unix()))")
	r.Add("perm := r.Perm(len(*x))")
	r.Add("for i, randIdx := range perm {")
	r.Add("(*x)[i], (*x)[randIdx] = (*x)[randIdx], (*x)[i]")
	r.Add("}")
	r.Add("return x")
	// test
	if x.f.Test != nil && Opt.IsArr() {
	}
	return r
}
func (x *PrtArr) mrg() (r *TypFn) {
	r = x.f.TypFn("Mrg")
	r.InPrmVariadic(x.f.Typ(), "a")
	r.OutPrm(x.f.Typ())
	r.Addf("for _, v := range %v {", r.InPrms[0].Name)
	r.Add("*x = append(*x, *v...)")
	r.Add("}")
	r.Add("return x")
	// test
	if x.f.Test != nil && Opt.IsArr() {
		r.T.Add("*exCpy = append(*exCpy, *a...)")
		r.T.Add("expected := exCpy")
	}
	return r
}
func (x *PrtArr) push() (r *TypFn) {
	r = x.f.TypFn("Push")
	r.InPrmVariadic(x.Arr.Elm, "a")
	r.OutPrm(x.f.Typ())
	r.Addf("*x = append(*x, %v...)", r.InPrms[0].Name)
	r.Add("return x")
	// test
	if x.f.Test != nil && Opt.IsArr() {
		r.T.Add("*exCpy = append(*exCpy, a)")
		r.T.Add("expected := exCpy")
	}
	return r
}
func (x *PrtArr) pop() (r *TypFn) {
	r = x.f.TypFn("Pop")
	r.OutPrm(x.Arr.Elm, "r")
	r.Add("if len(*x) == 0 {")
	r.Add("return r")
	r.Add("}")
	r.Add("r = (*x)[len(*x)-1]")
	r.Add("*x = (*x)[:len(*x)-1]")
	r.Add("return r")
	// test
	if x.f.Test != nil && Opt.IsArr() {
		r.T.Addf("var expected %v", r.OutTyp().Ref(x.f.Test))
		r.T.Add("if len(*exCpy) != 0 {")
		r.T.Add("expected = (*exCpy)[len(*exCpy)-1]")
		r.T.Add("*exCpy = (*exCpy)[:len(*exCpy)-1]")
		r.T.Add("}")
	}
	return r
}
func (x *PrtArr) que() (r *TypFn) {
	r = x.f.TypFn("Que")
	r.InPrmVariadic(x.Arr.Elm, "vs")
	r.OutPrm(x.f.Typ())
	r.Add("*x = append(*x, vs...)")
	r.Add("return x")
	// test
	if x.f.Test != nil && Opt.IsArr() {
		r.T.Addf("*exCpy = append(*exCpy, %v)", r.InPrms[0].Camel())
		r.T.Add("expected := exCpy")
	}
	return r
}
func (x *PrtArr) dque() (r *TypFn) {
	r = x.f.TypFn("Dque")
	r.OutPrm(x.Arr.Elm, "r")
	r.Add("r = (*x)[0]")
	r.Add("if len(*x) == 1 {")
	r.Add("*x = (*x)[:0]")
	r.Add("} else {")
	r.Add("copy(*x, (*x)[1:])")
	r.Add("*x = (*x)[:len(*x)-1]")
	r.Add("}")
	r.Add("return r")
	// test
	if x.f.Test != nil && Opt.IsArr() {
		r.T.TstCnd = "x.Cnt() != 0"
		r.T.Add("expected := (*exCpy)[0]")
		r.T.Add("if len(*exCpy) == 1 {")
		r.T.Add("*exCpy = (*exCpy)[:0]")
		r.T.Add("} else {")
		r.T.Add("copy(*exCpy, (*exCpy)[1:])")
		r.T.Add("*exCpy = (*exCpy)[:len(*exCpy)-1]")
		r.T.Add("}")
	}
	return r
}
func (x *PrtArr) ins() (r *TypFn) {
	r = x.f.TypFn("Ins")
	r.InPrm(_sys.Bsc.Unt, "idx").LitVal("0")
	r.InPrm(x.Arr.Elm, "elm")
	r.OutPrm(x.f.Typ())
	r.Addf("*x = append((*x)[:idx], append(%v{elm}, (*x)[idx:]...)...)", x.t.Title())
	r.Add("return x")
	// test
	if x.f.Test != nil && Opt.IsArr() {
		r.T.TstCnd = "idx <= x.Cnt()"
		r.T.Addf("v := append((*exCpy)[:idx], append(%v{elm}, (*exCpy)[idx:]...)...)", x.t.Full())
		r.T.Add("expected := &v")
	}
	return r
}
func (x *PrtArr) upd() (r *TypFn) {
	r = x.f.TypFn("Upd")
	r.InPrm(_sys.Bsc.Unt, "idx").LitVal("0")
	r.InPrm(x.Arr.Elm, "elm")
	r.OutPrm(x.f.Typ())
	r.Add("(*x)[idx] = elm")
	r.Add("return x")
	// test
	if x.f.Test != nil && Opt.IsArr() {
		r.T.TstCnd = "idx < x.Cnt()"
		r.T.Addf("(*exCpy)[idx] = elm")
		r.T.Add("expected := exCpy")
	}
	return r
}
func (x *PrtArr) del() (r *TypFn) {
	r = x.f.TypFn("Del")
	r.InPrm(_sys.Bsc.Unt, "idx").LitVal("0")
	r.OutPrm(x.Arr.Elm, "r")
	r.Add("r = (*x)[idx]")
	r.Add("if idx == 0 && len(*x) == 1 {")
	r.Add("*x = (*x)[:0]")
	r.Addf("} else if idx == %v(len(*x)-1) {", _sys.Bsc.Unt.Typ().Ref(x.f))
	r.Add("*x = (*x)[:idx]")
	r.Add("} else {")
	r.Add("*x = append((*x)[:idx], (*x)[idx+1:]...)")
	r.Add("}")
	r.Add("return r")
	// test
	if x.f.Test != nil && Opt.IsArr() {
		r.T.TstCnd = "idx < x.Cnt()"
		r.T.Add("expected := (*exCpy)[idx]")
		r.T.Add("if idx == 0 && len(*exCpy) == 1 {")
		r.T.Add("*exCpy = (*exCpy)[:0]")
		r.T.Addf("} else if idx == %v(len(*exCpy)-1) {", _sys.Bsc.Unt.Typ().Ref(x.f.Test))
		r.T.Add("*exCpy = (*exCpy)[:idx]")
		r.T.Add("} else {")
		r.T.Add("*exCpy = append((*exCpy)[:idx], (*exCpy)[idx+1:]...)")
		r.T.Add("}")
	}
	return r
}
func (x *PrtArr) at() (r *TypFn) {
	r = x.f.TypFn("At")
	r.InPrm(_sys.Bsc.Unt, "idx").LitVal("0")
	r.OutPrm(x.Arr.Elm)
	r.Add("return (*x)[idx]")
	// test
	if x.f.Test != nil && Opt.IsArr() {
		r.T.TstCnd = "idx < x.Cnt()"
		r.T.Add("expected := (*exCpy)[idx]")
		if x.t.IsAna() {
			r.T.Addf("tst.%vEql(t, exCpy, axCpy)", r.Rxr.Typ.PrefixTitle())
		} else {
			r.T.Addf("tst.%vEql(t, exCpy, axCpy)", r.Rxr.Typ.PrefixTitle())
		}
	}
	return r
}
func (x *PrtArr) in() (r *TypFn) {
	r = x.f.TypFn("In")
	r.InPrm(_sys.Bsc.Unt, "idx")
	r.InPrm(_sys.Bsc.Unt, "lim")
	r.OutPrm(x.f.Typ())
	r.Add("r := (*x)[idx:lim]")
	r.Add("return &r")
	// test
	if x.f.Test != nil && Opt.IsArr() {
		r.T.TstCnd = "idx < x.Cnt() && lim <= x.Cnt() && idx <= lim"
		r.T.Add("vs := (*exCpy)[idx:lim]")
		r.T.Add("expected := &vs")
		if x.t.IsAna() {
			r.T.Addf("tst.%vEql(t, exCpy, axCpy)", r.Rxr.Typ.PrefixTitle())
		} else {
			r.T.Addf("tst.%vEql(t, exCpy, axCpy)", r.Rxr.Typ.PrefixTitle())
		}
	}
	return r
}
func (x *PrtArr) inBnd() (r *TypFn) {
	r = x.f.TypFn("InBnd")
	r.InPrm(_sys.Bsc.Bnd, "b")
	r.OutPrm(x.f.Typ())
	r.Add("r := (*x)[b.Idx:b.Lim]")
	r.Add("return &r")
	// test
	if x.f.Test != nil && Opt.IsArr() {
		x.f.Test.ImportFn(r)
		r.T.TstCnd = "b.IsValid() && b.Idx < x.Cnt() && b.Lim < x.Cnt() && b.Idx < b.Lim"
		r.T.Add("vs := (*exCpy)[b.Idx:b.Lim]")
		r.T.Add("expected := &vs")
		if x.t.IsAna() {
			r.T.Addf("tst.%vEql(t, exCpy, axCpy)", r.Rxr.Typ.PrefixTitle())
		} else {
			r.T.Addf("tst.%vEql(t, exCpy, axCpy)", r.Rxr.Typ.PrefixTitle())
		}
	}
	return r
}
func (x *PrtArr) from() (r *TypFn) {
	r = x.f.TypFn("From")
	r.InPrm(_sys.Bsc.Unt, "idx")
	r.OutPrm(x.f.Typ())
	r.Addf("var r %v", x.t.Title())
	r.Add("if idx < unt.Unt(len(*x)) {")
	r.Add("r = (*x)[idx:]")
	r.Add("} else {")
	r.Add("r = (*x)[:0]")
	r.Add("}")
	r.Add("return &r")
	// test
	if x.f.Test != nil && Opt.IsArr() {
		r.T.TstCnd = "idx < x.Cnt()"
		r.T.Addf("var vs %v", r.OutTyp().Full())
		r.T.Add("if idx < unt.Unt(len(*exCpy)) {")
		r.T.Add("vs = (*exCpy)[idx:]")
		r.T.Add("} else {")
		r.T.Add("vs = (*exCpy)[:0]")
		r.T.Add("}")
		r.T.Add("expected := &vs")
	}
	return r
}
func (x *PrtArr) to() (r *TypFn) {
	r = x.f.TypFn("To")
	r.InPrm(_sys.Bsc.Unt, "lim")
	r.OutPrm(x.f.Typ())
	r.Add("if lim > unt.Unt(len(*x)) {")
	r.Add("return x")
	r.Add("}")
	r.Add("r := (*x)[:lim]")
	r.Add("return &r")
	// test
	if x.f.Test != nil && Opt.IsArr() {
		r.T.TstCnd = "lim <= x.Cnt()"
		r.T.Add("vs := (*exCpy)[:lim]")
		r.T.Add("expected := &vs")
	}
	return r
}
func (x *PrtArr) fst() (r *TypFn) {
	r = x.f.TypFn("Fst")
	r.OutPrm(x.Arr.Elm)
	r.Add("return (*x)[0]")
	// test
	if x.f.Test != nil && Opt.IsArr() {
		r.T.TstCnd = "x.Cnt() != 0"
		r.T.Add("expected := (*exCpy)[0]")
	}
	return r
}
func (x *PrtArr) mdl() (r *TypFn) {
	r = x.f.TypFn("Mdl")
	r.OutPrm(x.Arr.Elm)
	r.Add("return (*x)[len(*x)/2]")
	// test
	if x.f.Test != nil && Opt.IsArr() {
		r.T.TstCnd = "x.Cnt() != 0"
		r.T.Add("expected := (*exCpy)[len(*exCpy)/2]")
	}
	return r
}
func (x *PrtArr) lst() (r *TypFn) {
	r = x.f.TypFn("Lst")
	r.OutPrm(x.Arr.Elm)
	r.Add("return (*x)[len(*x)-1]")
	// test
	if x.f.Test != nil && Opt.IsArr() {
		r.T.TstCnd = "x.Cnt() != 0"
		r.T.Add("expected := (*exCpy)[len(*exCpy)-1]")
	}
	return r
}
func (x *PrtArr) fstIdx() (r *TypFn) {
	r = x.f.TypFn("FstIdx")
	r.OutPrm(_sys.Bsc.Unt)
	r.Addf("return 0")
	// test
	if x.f.Test != nil && Opt.IsArr() {
		r.T.SkpCpy = true
		r.T.Add("expected := unt.Zero")
	}
	return r
}
func (x *PrtArr) mdlIdx() (r *TypFn) {
	r = x.f.TypFn("MdlIdx")
	r.OutPrm(_sys.Bsc.Unt)
	r.Addf("return %v(len(*x)/2)", _sys.Bsc.Unt.Typ().Ref(x.f))
	// test
	if x.f.Test != nil && Opt.IsArr() {
		r.T.TstCnd = "x.Cnt() != 0"
		r.T.Addf("expected := %v(len(*exCpy)/2)", r.OutTyp().Full())
	}
	return r
}
func (x *PrtArr) lstIdx() (r *TypFn) {
	r = x.f.TypFn("LstIdx")
	r.OutPrm(_sys.Bsc.Unt)
	r.Addf("return %v(len(*x)-1)", _sys.Bsc.Unt.Typ().Ref(x.f))
	// test
	if x.f.Test != nil && Opt.IsArr() {
		r.T.TstCnd = "x.Cnt() != 0"
		r.T.Addf("expected := %v(len(*exCpy)-1)", r.OutTyp().Full())
	}
	return r
}

// TODO: IMPLEMENT/TEST WHEN HAVE TIME
// // TODO: Mdls, FstsPct, MdlsPct LstsPct
// func (x *Arr) Fsts() (r *TypFn) {
// 	r = x.f.TypFn("Fsts")
// 	r.InPrm(_sys.Bsc.Unt, "cnt")
// 	r.OutPrm(x.f.Typ())
// 	r.Addf("if cnt > %v(len(*x)) {", _sys.Bsc.Unt.Typ().Ref(x.Fle))
// 	r.Addf("cnt = %v(len(*x))", _sys.Bsc.Unt.Typ().Ref(x.Fle))
// 	r.Add("}")
// 	r.Addf("r := make(%v, cnt)", x.t.Title())
// 	r.Add("copy(r, *x)")
// 	r.Add("return &r")
// 	return r
// }
// func (x *Arr) Lsts() (r *TypFn) {
// 	r = x.f.TypFn("Lsts")
// 	r.InPrm(_sys.Bsc.Unt, "cnt")
// 	r.OutPrm(x.f.Typ())
// 	r.Addf("if cnt > %v(len(*x)) {", _sys.Bsc.Unt.Typ().Ref(x.Fle))
// 	r.Addf("cnt = %v(len(*x))", _sys.Bsc.Unt.Typ().Ref(x.Fle))
// 	r.Add("}")
// 	r.Addf("r := make(%v, cnt)", x.t.Title())
// 	r.Addf("copy(r, (*x)[%v(len((*x)))-cnt:])", _sys.Bsc.Unt.Typ().Ref(x.Fle))
// 	r.Add("return &r")
// 	return r
// }

func (x *PrtArr) rev() (r *TypFn) {
	r = x.f.TypFn("Rev")
	r.OutPrm(x.f.Typ())
	r.Add("for i, j := 0, len(*x)-1; i < j; i, j = i+1, j-1 { ")
	r.Add("(*x)[i], (*x)[j] = (*x)[j], (*x)[i]")
	r.Add("}")
	r.Add("return x")
	// test
	if x.f.Test != nil && Opt.IsArr() {
		r.T.SkpTst = true
		r.T.TstCnd = "x.Cnt() != 0"
		r.T.Add("for n := axCpy.FstIdx(); n.Lss(axCpy.Cnt()); n++ {")
		r.T.Add("eElm := exCpy.At(axCpy.Cnt()-n-1)")
		r.T.Add("aElm := actual.At(n)")
		r.T.Addf("%v(t, eElm, aElm, \"Idx\", n)", x.Arr.Elm.TstEql.Ref(x.f.Test))
		r.T.Add("}")
	}
	return r
}

// func (x *Arr) Gen(lits bool) (r []string) {
// 	b := strings.Builder{}
// 	var fst, mdl, lst string
// 	if lits {
// 		fst, mdl, lst = "[", " ", "]"
// 	} else {
// 		fst, mdl, lst = fmt.Sprintf("%v(", x.f.NewFull()), ",", ")"
// 	}
// 	inr := func(vs ...string) string {
// 		b.Reset()
// 		b.WriteString(fst)
// 		for n, v := range vs {
// 			if n != 0 {
// 				b.WriteString(mdl)
// 			}
// 			b.WriteString(fmt.Sprintf("%v", v))
// 		}
// 		b.WriteString(lst)
// 		return b.String()
// 	}
// 	var elms []string
// 	if lits {
// 		elms = x.Arr.Elm.Typ().Bse().Lits
// 	} else {
// 		elms = x.Arr.Elm.Typ().Bse().Vals
// 	}
// 	r = append(r, inr())
// 	r = append(r, inr(elms[:2]...)) // IMPORTANT: 2 for auto test gen
// 	r = append(r, inr(elms...))
// 	for n := 0; n < 4; n++ {
// 		elms = append(elms, elms...)
// 	}
// 	r = append(r, inr(elms...))
// 	return r
// }
