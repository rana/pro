package tpl

import (
	"strings"
	"sys/k"
	"sys/tpl/atr"
)

type (
	FleHstInstr struct {
		FleHstBse
	}
	FleHstInstrs struct {
		FleBse
		PrtArr
		PrtArrStrWrt // for PrmWrt
	}
)

func (x *DirHst) NewInstr() (r *FleHstInstr) {
	r = &FleHstInstr{}
	x.Instr = r
	r.Name = k.Instr
	r.Pkg = x.Pkg
	r.Ifc(r.Name, atr.TypAnaIfc)
	r.AddFle(r)
	return r
}
func (x *DirHst) NewInstrs() (r *FleHstInstrs) {
	r = &FleHstInstrs{}
	x.Instrs = r
	r.FleBse = *NewArr(x.Instr, &r.PrtArr, x.Instr.Pkg)
	r.AddFle(r)
	return r
}

func (x *FleHstInstr) InitTyp(bse *TypBse) {
	x.FleHstBse.InitTyp(bse)
	x.Typ().Bse().TestPth = append(_sys.Ana.Hst.Prv.Typ().Bse().TestPth, &TestStp{
		Fst: func(r *PkgFn) {
			r.Add("for _, instr := range tst.HstPrvInstrs {")
		},
		MdlFst: func(r *PkgFn) { r.Add("instr := instr(prv)") },
		Lst:    func(r *PkgFn) { r.Add("}") },
	})

}

func (x *FleHstInstr) InitFld(s *Struct) {
	x.FleHstBse.InitFld(s)
	x.bse.FldPrnt(_sys.Ana.Hst.Prv)
	x.bse.Fld("Ana", _sys.Ana.Instr)
	x.bse.Fld("TmeBnd", _sys.Bsc.Bnd)
	x.bse.FldSlice("Rng", _sys.Bsc.TmeRng)
}
func (x *FleHstInstr) InitTypFn() {
	x.FleHstBse.InitTypFn()
	x.I()
}
func (x *FleHstInstr) I() (r *TypFn) {
	inrvl, name := _sys.Ana.Hst.Inrvl, strings.Title(k.I)
	x.Import(_sys)
	x.Import(_sys.Bsc.Unt)
	x.Import(_sys.Bsc.Tme.arr)
	x.Import(_sys.Bsc.Bnd)
	x.Import(_sys.Bsc.Bnd.arr)
	r = x.ElmNodeTypFn(name, "", "", inrvl, func(r *TypFn) {
		r.InPrm(_sys.Bsc.Tme, "dur").LitVal("10")
	})
	r.Add("if ana.Cfg.Trc.IsHstInrvl() {")
	r.Addf("sys.Logf(\"%%p %v(%%v)\", r, r.Prm())", r.Node.Full())
	r.Add("}")
	r.Add("if r.Dur < 1 {")
	r.Add("return r")
	r.Add("}")
	r.Add("ts := x.Ana.HstStm.Tmes")
	r.Add("r.Tmes = tmes.Make(ts.Cnt()).Clr()")
	r.Add("r.TmeBnds = bnds.Make(ts.Cnt()).Clr()")
	r.Add("if !x.TmeBnd.IsValid() {")
	r.Add("if x.TmeBnd.Idx == unt.Max { // 0s-0s for LongRlng")
	r.Add("return r")
	r.Add("}")
	r.Add("} else {")
	r.Add("ts = ts.InBnd(x.TmeBnd) // instr time range qualifier .EurUsd(rng)")
	r.Add("}")
	r.Add("for s := 0; s < len(*ts); s++ { // rolling inrvl // EXPECT NO EQUAL START TMES DUE TO INTAKE PROCESSING")
	r.Add("idx := ts.SrchIdxEql((*ts)[s] + r.Dur)")
	r.Add("if idx == unt.Unt(len(*ts)) && (s == len(*ts)-1 || (*ts)[s+1]-(*ts)[s] < ana.MktSessionGap) {")
	r.Add("if (*ts)[s]+r.Dur < ts.Lst() {")
	r.Addf("sys.Logf(\"%v.%v(%%v): MISSING INRVL END %%v\", r.Dur, (*ts)[s]+r.Dur)", x.Typ().Full(), r.Name)
	r.Add("}")
	r.Add("break")
	r.Add("}")
	r.Add("r.Tmes.Push((*ts)[idx-1]) // USE idx-1 TO MATCH RLT BEHAVIOR")
	r.Add("r.TmeBnds.Push(bnd.Bnd{Idx: x.TmeBnd.Idx + unt.Unt(s), Lim: x.TmeBnd.Idx + unt.Unt(idx)})")
	r.Add("}")
	r.Add("if len(*r.Tmes) == 0 {")
	r.Add("r.Tmes = nil")
	r.Add("r.TmeBnds = nil")
	r.Add("}")
	r.Add("return r")
	return r
}
