package tpl

import (
	"fmt"
	"go/format"
	"path/filepath"
	"reflect"
	"strings"
	"sys"
	"sys/err"
	"sys/fs"
	"sys/k"
	"sys/tpl/atr"
	"sys/tpl/mod"
)

const (
	TestSuffix = "_test"
)

type (
	Fle interface {
		Lblr
		InitTyp(bse *TypBse)
		InitFld(s *Struct)
		InitIfc(i *Ifc)
		InitVals(bse *TypBse)
		InitCnst()
		InitVar()
		InitPkgFn()
		InitTypFn()
		InitTrm(bse *TypBse, trmr *FleTrmr)
		InitXpr(xprr *FleXprr)
		InitAct(actr *FleActr)
		WriteToDisk(dir string)
		Typ() Typ
		Bse() *FleBse
		Ref(f Fle) string
		Adr(f Fle) string
	}
	Fles []Fle

	FleBse struct {
		Lbl
		Prts
		Imports
		Typs
		Cnsts
		Vars
		PkgFns
		TypFns
		MemSigs
		Pkg       *Pkg
		Test      *FleTest
		ActCnsts  Cnsts
		ActPkgFns PkgFns
		ActTypFns TypFns
		bse       *Struct
		arr       *Arr
		// fbr       *Ifc
		// wve       *Ifc
	}
)

func (x *Fles) AddFle(fs ...Fle) {
	for _, f := range fs {
		*x = append(*x, f)
	}
}
func (x *Fles) WriteToDisk(dir string) {
	for _, f := range *x {
		// fmt.Println("* WriteToDisk:", f.Camel())
		f.WriteToDisk(dir)
	}
}

func (x *Fles) Init() {
	for _, f := range *x { // assign fle to each typ
		for _, t := range f.Bse().Typs {
			t.Bse().Fle = f
		}
	}
	for _, f := range *x { // RegPrts
		f.Bse().RegPrts(f)
	}

	for _, f := range *x { // typs
		f.Bse().Prts.InitPrtTyp()
		var bse *TypBse
		if f.Typ() != nil {
			bse = f.Typ().Bse()
		}
		f.InitTyp(bse)
	}
	for _, f := range *x { // typs: scp
		for _, t := range f.Bse().Typs {
			if t.Bse().IsScp() {
				f.Bse().Scp(t)
			}
		}
	}
	for _, f := range *x { // flds
		f.Bse().Prts.InitPrtFld()
		var s *Struct
		if f.Typ() != nil {
			if cur, ok := f.Typ().(*Struct); ok {
				s = cur
			}
		}
		f.InitFld(s)
	}
	for _, f := range *x { // ifcs
		var i *Ifc
		if f.Typ() != nil {
			if cur, ok := f.Typ().(*Ifc); ok {
				i = cur
			}
		}
		f.InitIfc(i)
		f.Bse().Prts.InitPrtIfc()
	}
	for _, f := range *x { // cnsts
		f.InitCnst()
		f.Bse().Prts.InitPrtCnst() // PLACE AFTER f.InitCnst() FOR ENUM OPERATION ON DEFINED CONSTS
		if f.Typ() != nil {
			_sys.Tst.GenCnst(f.Typ())
		}
	}
	for _, ext := range ExtTsts { // tst ext typs; BEFORE typs TO SUPPORT IFC ana.Pth
		_sys.Tst.GenTyp(ext)
	}
	for _, f := range *x { // tst typs
		if f.Typ() != nil && f.Typ().Bse().IsTst() {
			_sys.Tst.GenTyp(f.Typ())
		}
	}
	for _, f := range *x { // vals
		// sys.Log("-", f.Bse().Name)
		var bse *TypBse
		if f.Typ() != nil {
			bse = f.Typ().Bse()
		}
		f.InitVals(bse)
	}
	for _, f := range *x { // vars
		f.Bse().Prts.InitPrtVar()
		f.InitVar()
	}
	for _, f := range *x { // prt: pkg fns
		f.Bse().Prts.InitPrtPkgFn()
	}
	for _, f := range *x { // pkg fns
		f.InitPkgFn()
	}
	for _, f := range *x { // typ fns
		f.InitTypFn()
		f.Bse().Prts.InitPrtTypFn()
	}
	_sys.Tst.GenNodeVarSlices()

	// TRMR: separate sections allow elm, arr, fbr inits spearately
	trmr := _sys.Lng.Pro.Trm.Trmr
	for _, f := range *x { // trm: lit
		for _, typ := range f.Bse().Typs {
			if bse := typ.Bse(); bse.IsXpr() && bse.IsLit() && bse.LitTrm != nil {
				f.Bse().InitCtor()
			}
		}
	}
	for _, f := range *x { // trm: typ-specific
		for _, typ := range f.Bse().Typs {
			if bse := typ.Bse(); bse.IsXpr() {
				if s, ok := typ.(*Struct); ok { // struct
					// sys.Logf(" -- %v ", s.Full())
					for _, fld := range s.Flds {
						if fld.IsGet() || fld.IsSetGet() {
							fld.Trm = trmr.LexTrm(fld.Name)
						}
					}
				}
			}
		}
	}
	for _, f := range *x { // trm: cnst
		for _, cnst := range f.Bse().Cnsts {
			// sys.Logf(" --- fle:%v  cnst:%v  isTrm:%v", f.Title(), cnst.Title(), cnst.IsXpr())
			if cnst.MayXpr() {
				cnst.Trm = trmr.LexTrm(cnst.Name)
			}
		}
	}
	for _, f := range *x { // trm: var
		for _, Var := range f.Bse().Vars {
			if Var.MayXpr() {
				Var.Trm = trmr.LexTrm(Var.Name)
			}
		}
	}
	for _, f := range *x { // trm: pkgFn
		for _, fn := range f.Bse().PkgFns {
			if fn.MayXpr() {
				fn.Trm = trmr.LexTrm(fn.Name)
			}
		}
	}
	for _, f := range *x { // trm: typFn
		for _, typ := range f.Bse().Typs {
			if typ.Bse().MayXpr() {
				// sys.Log("typ", typ.Bse().Name)
				for _, fn := range typ.Bse().TypFns {
					// sys.Log("fn", fn.Name)
					if fn.MayXpr() {
						fn.Trm = trmr.LexTrm(fn.Name)
					}
				}
			}
		}
	}
	for _, f := range *x { // trm: memSig
		for _, typ := range f.Bse().Typs {
			if ifc, ok := typ.(*Ifc); ok {
				if ifc.MayXpr() {
					for _, fn := range ifc.MemSigs {
						if fn.MayXpr() {
							fn.Trm = trmr.LexTrm(fn.Name)
						}
					}
				}
			}
		}
	}
	for _, f := range *x { // trm: pkg
		for _, typ := range f.Bse().Typs {
			if bse := typ.Bse(); bse.IsTrm() && !trmr.HasPkg(bse.Pkg) {
				bse.Pkg.Trm = trmr.LexTrm(bse.Pkg.Name)
				trmr.Pkgs.AddPkg(bse.Pkg)
			}
		}
	}
	for _, f := range *x { // trm: overrides
		for _, typ := range f.Bse().Typs {
			if bse := typ.Bse(); bse.IsTrm() {
				f.InitTrm(bse, trmr)
			}
		}
	}
	for _, f := range *x { // trm: prts
		f.Bse().Prts.InitPrtTrm(trmr)
	}

	// XPRR: separate sections allow elm, arr, fbr inits spearately; Xprr.TypFnStm calls stm and stms
	xprr := _sys.Lng.Pro.Xpr.Xprr
	for _, f := range *x { // xpr: prts
		f.Bse().Prts.InitPrtXpr(xprr)
	}
	for _, f := range *x { // xpr: ifc
		for _, typ := range f.Bse().Typs {
			if bse := typ.Bse(); bse.IsXpr() {
				xprr.TypIfc(typ)
			}
		}
	}
	for _, f := range *x { // xpr: lit
		for _, typ := range f.Bse().Typs {
			if bse := typ.Bse(); bse.IsXpr() && bse.IsLit() && bse.LitTrm != nil {
				xprr.LitXpr(typ)
			}
		}
	}
	for _, f := range *x { // xpr: scp
		for _, typ := range f.Bse().Typs {
			if bse := typ.Bse(); bse.IsXpr() && bse.Scp != nil {
				xprr.AsnXpr(typ)
				xprr.AcsXpr(typ)
				if arr, ok := typ.(*Arr); ok {
					xprr.EachXpr(bse, arr)
					xprr.PllEachXpr(arr)
				}
			}
		}
	}
	for _, f := range *x { // xpr: typ-specific
		for _, typ := range f.Bse().Typs {
			if bse := typ.Bse(); bse.IsXpr() {
				if typ == _sys.Bsc.Bol.Typ() { // bol
					xprr.ThenXpr(typ)
					xprr.ElseXpr(typ)
				}
				if s, ok := typ.(*Struct); ok { // struct
					for _, fld := range s.Flds {
						if fld.IsGet() {
							xprr.FldGetXpr(s, fld)
						}
						if fld.IsSetGet() {
							xprr.FldSetGetXpr(s, fld)
						}
					}
				}
			}
		}
	}
	for _, f := range *x { // xpr: cnst
		for _, cnst := range f.Bse().Cnsts {
			if cnst.MayXpr() {
				xprr.CnstXpr(cnst)
			}
		}
	}
	for _, f := range *x { // xpr: var
		for _, Var := range f.Bse().Vars {
			if Var.MayXpr() {
				xprr.VarXpr(Var)
			}
		}
	}
	for _, f := range *x { // xpr: pkgFn
		for _, fn := range f.Bse().PkgFns {
			if fn.MayXpr() {
				xprr.PkgFnXpr(fn)
			}
		}
	}
	for _, f := range *x { // xpr: typFn
		for _, typ := range f.Bse().Typs {
			if typ.Bse().MayXpr() {
				for _, fn := range typ.Bse().TypFns {
					if fn.MayXpr() {
						xprr.TypFnXpr(fn)
					}
				}
			}
		}
	}
	for _, f := range *x { // xpr: memSig
		for _, typ := range f.Bse().Typs {
			if ifc, ok := typ.(*Ifc); ok {
				if ifc.MayXpr() {
					for _, fn := range ifc.MemSigs {
						if fn.MayXpr() {
							xprr.MemSigXpr(fn)
						}
					}
				}
			}
		}
	}
	for _, f := range *x { // xpr: pkg
		for _, typ := range f.Bse().Typs {
			if bse := typ.Bse(); bse.IsXpr() && !xprr.HasPkg(bse.Pkg) {
				xprr.Pkgs.AddPkg(bse.Pkg)
			}
		}
	}
	for _, f := range *x { // xpr: overrides (Xprr uses)
		f.InitXpr(xprr)
	}

	// ACTR: separate sections allow elm, arr, fbr inits spearately; Actr.TypFnStm calls stm and stms act.HstStmOtrAdds
	actr := _sys.Lng.Pro.Act.Actr
	for _, f := range *x { // act: prts
		f.Bse().Prts.InitPrtAct(actr)
	}
	for _, f := range *x { // act: ifc
		for _, typ := range f.Bse().Typs {
			if bse := typ.Bse(); bse.IsAct() {
				actr.TypIfc(typ)
			}
		}
	}
	for _, f := range *x { // act: lit
		for _, typ := range f.Bse().Typs {
			if bse := typ.Bse(); bse.IsAct() && bse.IsLit() && bse.LitXpr != nil {
				actr.LitAct(typ)
			}
		}
	}
	for _, f := range *x { // act: scp
		for _, typ := range f.Bse().Typs {
			if bse := typ.Bse(); bse.IsAct() && bse.Scp != nil {
				actr.AsnAct(typ)
				actr.AcsAct(typ)
				if arr, ok := typ.(*Arr); ok {
					actr.EachAct(typ, arr)
					actr.PllEachAct(arr)
				}
			}
		}
	}
	for _, f := range *x { // act: typ-specific
		for _, typ := range f.Bse().Typs {
			if bse := typ.Bse(); bse.IsAct() {
				if typ == _sys.Bsc.Bol.Typ() { // bol
					actr.ThenAct(typ)
					actr.ElseAct(typ)
				}
				if s, ok := typ.(*Struct); ok { // struct
					for _, fld := range s.Flds {
						if fld.IsGet() {
							actr.FldGetAct(s, fld)
						}
						if fld.IsSetGet() {
							actr.FldSetGetAct(s, fld)
						}
					}
				}
			}
		}
	}
	for _, f := range *x { // act: cnst
		for _, cnst := range f.Bse().Cnsts {
			if cnst.Xpr != nil {
				actr.CnstAct(cnst)
			}
		}
	}
	for _, f := range *x { // act: var
		for _, Var := range f.Bse().Vars {
			if Var.Xpr != nil {
				actr.VarAct(Var)
			}
		}
	}
	for _, f := range *x { // act: pkgFn
		// sys.Log("-", f.Bse().Name)
		for _, fn := range f.Bse().PkgFns {
			if fn.Xpr != nil {
				// sys.Log("  ", fn.PkgTitle())
				actr.PkgFnAct(fn)
			}
		}
	}
	for _, f := range *x { // act: typFn
		for _, typ := range f.Bse().Typs {
			if typ.Bse().IfcXpr != nil {
				for _, fn := range typ.Bse().TypFns {
					if fn.Xpr != nil {
						actr.TypFnAct(fn)
					}
				}
			}
		}
	}
	for _, f := range *x { // act: memSig
		for _, typ := range f.Bse().Typs {
			if ifc, ok := typ.(*Ifc); ok {
				if ifc.IfcXpr != nil {
					for _, fn := range ifc.MemSigs {
						if fn.Xpr != nil {
							actr.MemSigAct(fn)
						}
					}
				}
			}
		}
	}
	for _, f := range *x { // act: pkg
		for _, typ := range f.Bse().Typs {
			if bse := typ.Bse(); bse.IsAct() && !actr.HasPkg(bse.Pkg) {
				actr.Pkgs.AddPkg(bse.Pkg)
			}
		}
	}
	for _, f := range *x { // act: overrides
		f.InitAct(actr)
	}
}
func (x *FleBse) RegPrts(f Fle) { // f interface required; not FleBse
	if x.Typ() != nil {
		rPrtType := reflect.TypeOf((*Prt)(nil)).Elem()
		rvFleIfc := reflect.ValueOf(f)
		rvFle := rvFleIfc.Elem()
		rtTyp := rvFle.Type()
		for n := 0; n < rtTyp.NumField(); n++ {
			rtFld := rtTyp.Field(n)
			rvFld := rvFle.Field(n)
			if rtFld.Anonymous && rvFld.CanAddr() {
				fva := rvFld.Addr()
				// fmt.Printf("%v: %d: %s -> %t\n", rtTyp.Name(), n, rtFld.Name, fva.Type().Implements(rPrtType))
				if fva.Type().Implements(rPrtType) {
					rRegPrt := fva.MethodByName("RegPrt")
					rRegPrt.Call([]reflect.Value{fva, rvFleIfc})
				}
			}
		}
	}
}
func (x *FleBse) InitTyp(bse *TypBse) {}
func (x *FleBse) InitFld(s *Struct)   {}
func (x *FleBse) InitIfc(i *Ifc)      {}
func (x *FleBse) InitVals(bse *TypBse) {
	if bse != nil && bse.Vals == nil && !bse.IsUi() {
		switch {
		case bse.IsAna():
			bse.Lits = GenAna(x.Typ(), false) // pass most derived typ; not bse
			bse.Vals = GenAna(x.Typ(), true)  // pass most derived typ; not bse
		case bse.IsArr():
			bse.Lits = GenArr(x.Typ().(*Arr), false, nil)
			bse.Vals = GenArr(x.Typ().(*Arr), true, nil)
		case bse.IsStruct():
			bse.Lits = GenLitsStruct(x.Typ().(*Struct))
			bse.Vals = GenValsStruct(x.Typ().(*Struct))
		}
	}
}
func (x *FleBse) InitCnst()  {}
func (x *FleBse) InitVar()   {}
func (x *FleBse) InitPkgFn() {}
func (x *FleBse) InitTypFn() {}
func (x *FleBse) InitJsnLexLit(lex func(r *TypFn, f *FleJsnTrmr), litTrm ...*Struct) {
	trmr := _sys.Lng.Jsn.Trm.Trmr
	bse := x.Typ().Bse()
	if len(litTrm) == 0 {
		bse.LitTrmJsn = trmr.StructLit(bse.Title())
	} else {
		bse.LitTrmJsn = litTrm[0]
	}
	bse.LexTrmJsn = trmr.MemLit(bse.LitTrmJsn)
	lex(bse.LexTrmJsn, trmr)
	trmr.Test.LitTrmTyp(bse, true)
}
func (x *FleBse) InitJsnPrsLit(prs func(r *PkgFn, f *FleJsnPrs)) {
	_sys.Lng.Jsn.Trm.Prs.PrsTypTrm(x.Typ(), prs)
}
func (x *FleBse) InitJsnPrs() {
	_sys.Lng.Jsn.Jsnr.PrsJsn(x)
}

func (x *FleBse) InitLexLit(lex func(r *TypFn, f *FleTrmr), litTrm ...*Struct) {
	trmr, bse := _sys.Lng.Pro.Trm.Trmr, x.Typ().Bse()
	if len(litTrm) == 0 {
		bse.LitTrm = trmr.StructLit(bse.Title())
	} else {
		bse.LitTrm = litTrm[0]
	}
	bse.LitLex = trmr.MemLit(bse.LitTrm)
	lex(bse.LitLex, trmr)
	trmr.Test.LitTrmTyp(bse)
	x.InitTrm(bse, trmr)
}

func (x *FleBse) InitCtor() {
	trmr, bse := _sys.Lng.Pro.Trm.Trmr, x.Typ().Bse()
	s := x.Typ().(*Struct)
	trmr.InitCtor(s)
	x.InitPrsLit(func(r *PkgFn, f *FlePrs) {
		if s.IsPtr() {
			r.Addf("r = %v{}", s.Adr(f))
		}
		for _, fld := range s.Flds {
			fldBse := fld.Typ.Bse()
			if fld.IsTrm() { // exported fld with lit
				r.Addf("if trm.%v.Bnd.Cnt() != 0 {", fld.Name)
				r.Addf("r.%v = %v(trm.%v, txt)", fld.Name, fldBse.PrsTrm.Ref(f), fld.Name)
				r.Add("}")
			}
		}
		r.Add("return r")
	})
	x.InitTrm(bse, trmr)
}
func (x *FleBse) InitPrsLit(prs func(r *PkgFn, f *FlePrs)) {
	_sys.Lng.Pro.Trm.Prs.PrsTypTrm(x.Typ(), prs)
}
func (x *FleBse) InitPrsCfg() {
	_sys.Lng.Pro.Cfg.Cfgr.PrsCfg(x)
}

func (x *FleBse) InitTrm(bse *TypBse, trmr *FleTrmr) {}
func (x *FleBse) InitXpr(xprr *FleXprr)              {}
func (x *FleBse) InitAct(actr *FleActr)              {}
func (x *FleBse) FleTest(name ...string) (r *FleTest) {
	r = &FleTest{}
	r.Pkg = x.Pkg.NewTest()
	r.Src = x
	if len(name) != 0 {
		r.Name = name[0]
	} else {
		r.Name = x.Typ().Camel()
	}
	_sys.AddFle(r)
	return r
}
func (x *FleBse) FleTestTyp(typ Typ) (r *FleTest) {
	r = &FleTest{}
	r.Pkg = x.Pkg.NewTest()
	r.Src = x
	r.Name = typ.Camel()
	_sys.AddFle(r)
	return r
}
func (x *FleBse) AddFle(f Fle) {
	// if x.Typ() != nil { // define tst before typ for init order; tst.asc available to typ during test writing
	// 	if x.Typ().Bse().IsTst() {
	// 		x.Tst = x.FleTst()
	// 	}
	// }
	_sys.AddFle(f)
	if x.Typ() != nil {
		if x.Typ().Bse().IsTest() {
			// fmt.Println("-", x.Typ().Title())
			x.Test = x.FleTest()
			x.Typ().Bse().Test = x.Test // for node
		}

	}
	if f.Bse().Pkg != nil {
		f.Bse().Pkg.AddFle(f)
	}
}
func (x *FleBse) Bse() *FleBse     { return x }
func (x *FleBse) Ref(f Fle) string { return x.Typ().Ref(f) }
func (x *FleBse) Adr(f Fle) string { return x.Typ().Adr(f) }

// Typ
func (x *FleBse) AliasMod(name string, alias Typ, aliasMod mod.Mod, a atr.Atr) (r *Alias) {
	r = x.Typs.Alias(name, alias, x.Pkg, aliasMod, a)
	// if a.IsIfc() {
	// 	x.TypFnIfc(r)
	// }
	return r
}
func (x *FleBse) Alias(name string, alias Typ, a atr.Atr) (r *Alias) {
	return x.AliasMod(name, alias, mod.None, a)
}
func (x *FleBse) Aliasf(format string, alias Typ, a atr.Atr, args ...interface{}) (r *Alias) {
	return x.AliasMod(fmt.Sprintf(format, args...), alias, mod.None, a)
}
func (x *FleBse) AliasSlice(name string, alias Typ, a atr.Atr) (r *Alias) {
	r = x.AliasMod(name, alias, mod.Slice, a)
	r.Mod = mod.Ptr // mod return typ
	return r
}
func (x *FleBse) StructMod(name string, m mod.Mod, a atr.Atr) (r *Struct) {
	r = x.Typs.Struct(name, x.Pkg, m, a)
	// if a.IsIfc() {
	// 	x.TypFnIfc(r)
	// }
	return r
}
func (x *FleBse) Struct(name string, a atr.Atr) (r *Struct) {
	r = x.StructMod(name, mod.None, a)
	return r
}
func (x *FleBse) Structf(format string, a atr.Atr, args ...interface{}) (r *Struct) {
	r = x.StructMod(fmt.Sprintf(format, args...), mod.None, a)
	return r
}
func (x *FleBse) StructPtr(name string, a atr.Atr) (r *Struct) {
	r = x.StructMod(name, mod.Ptr, a)
	return r
}
func (x *FleBse) StructPtrf(format string, a atr.Atr, args ...interface{}) (r *Struct) {
	r = x.StructMod(fmt.Sprintf(format, args...), mod.Ptr, a)
	return r
}
func (x *FleBse) Ifc(name string, a atr.Atr) (r *Ifc) {
	r = x.Typs.Ifc(name, x.Pkg, a)
	// if a.IsIfc() {
	// 	x.MemSigIfc(r)
	// }
	return r
}
func (x *FleBse) Ifcf(format string, a atr.Atr, args ...interface{}) (r *Ifc) {
	return x.Ifc(fmt.Sprintf(format, args...), a)
}
func (x *FleBse) Map(name string, key, val FleOrTyp, a atr.Atr) (r *Map) {
	return x.Typs.Map(name, key, val, x.Pkg, a)
}
func (x *FleBse) Mapf(format string, key, val FleOrTyp, a atr.Atr, args ...interface{}) (r *Map) {
	return x.Map(fmt.Sprintf(format, args...), key, val, a)
}
func (x *FleBse) Func(name string, a atr.Atr) (r *Func) {
	return x.Typs.Func(name, x.Pkg, a)
}
func (x *FleBse) Funcf(format string, a atr.Atr, args ...interface{}) (r *Func) {
	return x.Func(fmt.Sprintf(format, args...), a)
}
func (x *FleBse) FuncFn(name string, fn Fn) (r *Func) {
	r = x.Func(name, atr.None)
	r.AddFn(fn)
	return r
}

// Cnst
func (x *FleBse) Cnst(name, value string, fleOrTyp ...interface{}) (r *Cnst) {
	var t Typ
	if len(fleOrTyp) != 0 {
		switch cur := fleOrTyp[0].(type) {
		case Typ:
			t = cur
		case Fle:
			t = cur.Typ()
		default:
			err.Panicf("invalid fle or type (v:%v)", fleOrTyp[0])
		}
	} else {
		t = x.Typ()
	}
	return x.Cnsts.Cnst(name, value, t, x.Pkg, atr.Lng)
}
func (x *FleBse) Cnstf(format string, args ...interface{}) (r *Cnst) {
	return x.Cnst(fmt.Sprintf(format, args...), "")
}
func (x *FleBse) CnstSize(size int) (r *Cnst) {
	r = x.Cnst("Size", fmt.Sprintf("%v", size), Int)
	r.Atr = atr.None
	x.Typ().Bse().Size = r
	return r
}

// Var
func (x *FleBse) Var(name, value string, fleOrTyp ...interface{}) (r *Var) {
	var t Typ
	if len(fleOrTyp) != 0 {
		switch cur := fleOrTyp[0].(type) {
		case Typ:
			t = cur
		case Fle:
			t = cur.Typ()
		default:
			err.Panicf("invalid fle or type (v:%v)", fleOrTyp[0])
		}
	} else {
		t = x.Typ()
	}
	return x.Vars.Var(name, value, t, x.Pkg, atr.Lng)
}

// PkgFn
func (x *FleBse) PkgFn(name string, ensureUnique ...bool) (r *PkgFn) {
	if len(ensureUnique) != 0 && ensureUnique[0] && x.Typs.Ok() && x.Typ().Lower() != x.Pkg.Lower() {
		name = x.Typ().Bse().UniqueTitle(name) // NewStms, MakeStms, ...
	}
	r = &PkgFn{}
	r.Name = strings.Title(name)
	r.Pkg = x.Pkg
	r.Atr = atr.LngTest
	r.Fle = x
	x.PkgFns.AddPkgFn(r)
	return r
}
func (x *FleBse) PkgFnf(format string, args ...interface{}) (r *PkgFn) {
	r = x.PkgFn(fmt.Sprintf(format, args...))
	return r
}
func (x *FleBse) PkgFna(name string, a atr.Atr, ensureUnique ...bool) (r *PkgFn) {
	r = x.PkgFn(name, ensureUnique...)
	r.Atr = a
	return r
}
func (x *FleBse) PkgFnNew(name string) (r *PkgFn) {
	r = x.PkgFnf("New%v", strings.Title(name))
	return r
}

// TypFn
func (x *FleBse) TypFn(name string, rxr ...Typ) (r *TypFn) {
	r = &TypFn{}
	r.Name = strings.Title(name)
	r.Atr = atr.LngTest
	r.Fle = x
	var t Typ
	if len(rxr) > 0 && rxr[0] != nil && !reflect.ValueOf(rxr[0]).IsNil() {
		t = rxr[0]
	} else {
		t = x.Typ()
	}
	r.Rxr = t.Bse().Prm("x")
	r.Rxr.Typ.Bse().Fns[r.Name] = r
	r.Rxr.Typ.Bse().TypFns.AddTypFn(r)
	x.TypFns.AddTypFn(r)
	return r
}
func (x *FleBse) TypFnf(format string, args ...interface{}) (r *TypFn) {
	return x.TypFn(fmt.Sprintf(format, args...))
}
func (x *FleBse) TypFnRxrf(format string, rxr Typ, args ...interface{}) (r *TypFn) {
	return x.TypFn(fmt.Sprintf(format, args...), rxr)
}
func (x *FleBse) TypFna(name string, a atr.Atr, rxr ...Typ) (r *TypFn) {
	r = x.TypFn(name, rxr...)
	r.Atr = a
	return r
}

// func (x *FleBse) TypFnIfc(rxr Typ) (r *TypFn) {
// 	r = x.TypFna(k.Ifc, atr.LngScp, rxr)
// 	r.OutPrm(_sys.Ifc)
// 	r.Add("return x")
// 	return r
// }

// MemSig
func (x *FleBse) MemSig(name string, rxr ...*Ifc) (r *MemSig) {
	r = &MemSig{}
	r.Name = strings.Title(name)
	if len(rxr) != 0 {
		r.Rxr = rxr[0]
	} else {
		r.Rxr = x.Typ().(*Ifc)
	}
	r.Atr = atr.LngTest
	r.Fle = x
	r.Rxr.MemSigs.AddSig(r)
	x.MemSigs.AddSig(r) // add to file for Xprr
	return r
}
func (x *FleBse) MemSigFn(fn Fn, rxr ...*Ifc) (r *MemSig) {
	r = x.MemSig(fn.Bse().Name, rxr...)
	for _, p := range fn.Bse().InPrms {
		r.AddInPrm(p)
	}
	for _, p := range fn.Bse().OutPrms {
		r.AddOutPrm(p)
	}
	return r
}
func (x *FleBse) MemSigf(format string, args ...interface{}) (r *MemSig) {
	return x.MemSig(fmt.Sprintf(format, args...))
}
func (x *FleBse) MemSigRxrf(format string, rxr *Ifc, args ...interface{}) (r *MemSig) {
	return x.MemSig(fmt.Sprintf(format, args...), rxr)
}
func (x *FleBse) MemSiga(name string, a atr.Atr, rxr ...*Ifc) (r *MemSig) {
	r = x.MemSig(name, rxr...)
	r.Atr = a
	return r
}
func (x *FleBse) MemSigaOr(name string, a atr.Atr, rxr ...*Ifc) (r *MemSig) {
	r = x.MemSig(name, rxr...)
	r.Atr |= a
	return r
}

// func (x *FleBse) MemSigIfc(rxr *Ifc) (r *MemSig) {
// 	r = x.MemSig(k.Ifc, rxr)
// 	r.Atr |= atr.Scp
// 	r.OutPrm(_sys.Ifc)
// 	return r
// }

// ImplIfc
func (x *FleBse) ImplIfc(i *Ifc, t ...Typ) {
	if len(t) == 0 {
		t = append(t, x.Typ())
	}
	t[0].Bse().AddIfc(i)
	i.ConcreteTyps.AddTyp(t[0])
}
func (x *FleBse) EmbedIfc(rxr, embed *Ifc) (r *Ifc) {
	// copy all sigs
	// to enable embedded signatures to be surfaced in lng
	for _, memSig := range embed.MemSigs {
		x.MemSigFn(memSig, rxr)
	}

	// rxr.EmbedIfc(embed)
	// // enable embedded signatures to be surfaced in lng
	// for _, memSig := range embed.MemSigs {
	// 	sys.Log("#", memSig.Name, "MayXpr", memSig.MayXpr(), "PrmsMayXpr", memSig.FnBse.PrmsMayXpr())
	// 	if memSig.MayXpr() { // DO NOT CHECK IF RXR IS XPRABL; AnaPth case is not xprable
	// 		x.MemSigs.AddSig(memSig) // add to file for Xprr
	// 	}
	// }
	return r
}

// Scp
func (x *FleBse) Scp(typ Typ) (r *Struct) {
	r = x.Structf("%vScp", atr.None, typ.Title())
	typ.Bse().Scp = r
	r.Fld("Idx", Uint32)
	r.FldSlice("Arr", typ)
	return r
}

func (x *FleBse) WriteToDisk(dir string) {
	x.EnsureImports()
	dir = filepath.Join(dir, x.Pkg.Pth)
	fs.EnsureDir(dir)
	var suffix string
	if x.Pkg.IsTest() {
		suffix = ".gen_test.go"
	} else {
		suffix = ".gen.go"
	}
	filename := filepath.Join(dir, x.Name+suffix)
	fmt.Printf("  - Wrt %v \n", filename)

	// write to buffer
	b := &strings.Builder{}
	x.Pkg.WritePkg(b)
	x.WriteImports(b, x)
	x.WriteVarDecls(b, x)
	x.WriteCnstDecls(b, x)
	x.WriteTypDecls(b, x)
	x.WritePkgFns(b, x)
	x.WriteTypFns(b, x)

	// write to file
	formatted, err := format.Source([]byte(b.String()))
	if err == nil {
		fs.WriteFile(filename, []byte(formatted))
	} else {
		fs.WriteFile(filename, []byte(b.String()))
	}
}
func (x *FleBse) EnsureImports() {
	for _, t := range x.Typs { // typs
		for _, tRef := range t.TypRefs() {
			x.ImportTyp(tRef, x)
		}
	}
	for _, c := range x.Cnsts { // pkg cnsts
		if c.Typ != nil {
			x.ImportTyp(c.Typ, x)
			for _, tRef := range c.Typ.TypRefs() {
				x.ImportTyp(tRef, x)
			}
		}
	}
	for _, v := range x.Vars { // pkg vars
		if v.Typ != nil {
			x.ImportTyp(v.Typ, x)
			for _, tRef := range v.Typ.TypRefs() {
				x.ImportTyp(tRef, x)
			}
		}
	}
	for _, fn := range x.PkgFns { // pkg fns
		x.ImportFn(fn, x)
	}
	for _, fn := range x.TypFns { // typ fns
		x.ImportFn(fn, x)
	}
}

func (x *FleBse) NewFull() (r string) {
	var suffix string
	if x.Typ().Lower() != x.Pkg.Lower() {
		suffix = x.Typ().Title() // NewStms, MakeStms, ...
	}
	return fmt.Sprintf("%v.New%v", x.Pkg.Lower(), suffix)
}

func (x *FleBse) Print(msgs ...interface{}) {
	if x.Typs.Ok() {
		fmt.Println(append(sys.Is(k.Ty, x.Typ().PkgTypTitle()), msgs...)...)
	} else {
		fmt.Println(append(sys.Is(k.Ty), msgs...)...)
	}
}

type (
	FleOrTyp interface{}
)

func GetTyp(v FleOrTyp) Typ {
	switch cur := v.(type) {
	case Typ:
		return cur
	case Fle:
		return cur.Typ()
	default:
		err.Panicf("invalid fle or type (v:%v)", v)
	}
	return nil
}
