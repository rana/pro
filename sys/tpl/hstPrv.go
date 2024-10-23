package tpl

import (
	"strings"
	"sys"
	"sys/k"
	"sys/ks"
	"sys/tpl/atr"
)

type (
	FleHstPrv struct {
		FleHstBse
	}
	FleHstPrvs struct {
		FleBse
		PrtArr
		PrtArrStrWrt // for PrmWrt
		PrtString
	}
)

func (x *DirHst) NewPrv() (r *FleHstPrv) {
	r = &FleHstPrv{}
	x.Prv = r
	r.Name = k.Prv
	r.Pkg = x.Pkg
	r.Ifc(r.Name, atr.TypAnaIfc)
	r.AddFle(r)
	return r
}
func (x *DirHst) NewPrvs() (r *FleHstPrvs) {
	r = &FleHstPrvs{}
	x.Prvs = r
	r.FleBse = *NewArr(x.Prv, &r.PrtArr, x.Prv.Pkg)
	r.AddFle(r)
	return r
}

func (x *FleHstPrv) InitTyp(bse *TypBse) {
	x.FleHstBse.InitTyp(bse)
	x.Typ().Bse().TestPth = []*TestStp{&TestStp{
		MdlFst: func(r *PkgFn) { r.Add("prv := hst.Oan()") },
	}}
}
func (x *FleHstPrv) InitVals(bse *TypBse) {
	bse.Lits = sys.Vs("hst.oan()")
	bse.Vals = sys.Vs("hst.Oan()")
}
func (x *FleHstPrv) InitIfc(i *Ifc) {
	x.FleHstBse.InitIfc(i)
	var sig *MemSig
	sig = x.MemSig(k.Instr)
	sig.InPrm(_sys.Bsc.Str, "name")
	sig.OutPrm(_sys.Ana.Instr, "r")
	sig = x.MemSig(k.LoadHst)
	sig.InPrm(_sys.Ana.Instr, "i")
}

func (x *FleHstPrv) InitPkgFn() {
	x.Oan()
}
func (x *FleHstPrv) Oan() (r *PkgFn) {
	x.Import(_sys.Ana)
	r = x.NodePkgFn(k.Oan, "", x, func(r *PkgFn) {
		r.Node.FldTyp(NewExt("*ana.Oan"))
	})
	r.Add("r.Oan = ana.PrvOan")
	r.Add("return r")
	return r
}
func (x *FleHstPrv) InitTypFn() {
	x.FleHstBse.InitTypFn()
	for _, instr := range ks.Instrs {
		x.Instr(instr)
	}
	// x.Instrs()
}

func (x *FleHstPrv) Instr(name string) (r *TypFn) {
	instr, name := _sys.Ana.Hst.Instr, strings.Title(name)
	x.Import("sys/k")
	x.Import(_sys)
	r = x.ElmNodeTypFn(name, "", "", instr, func(r *TypFn) {
		r.InPrmVariadic(_sys.Bsc.TmeRng, "rng").Atr = atr.FldSkp
	})
	r.Add("if ana.Cfg.Trc.IsHstInstr() {")
	r.Addf("sys.Logf(\"%%p %v(%%v)\", r, r.Prm())", r.Node.Full())
	r.Add("}")
	r.Addf("r.Ana = r.Prv.Instr(k.%vName)", name)
	r.Add("if r.Ana.HstStm == nil {")
	r.Add("r.Prv.LoadHst(r.Ana)")
	r.Add("}")
	r.Add("if len(r.Rng) == 0 {")
	r.Add("r.TmeBnd.Lim = r.Ana.HstStm.Tmes.Cnt()")
	r.Add("} else {")
	r.Add("r.TmeBnd = r.Ana.HstStm.Tmes.Bnd(r.Rng[0])")
	r.Add("if !r.TmeBnd.IsValid() { // 0s-0s for LongRlng")
	r.Addf("r.TmeBnd.Idx = %v", _sys.Bsc.Unt.Max.Ref(x))
	r.Add("}")
	r.Add("}")
	r.Add("return r")
	return r
}

// TODO: KEEP?
// func (x *FleHstPrvOan) Instrs() (r *TypFn) {
// 	r = x.TypFn(k.Instrs)
// 	r.InPrmVariadic(_sys.Bsc.TmeRng, "rng")
// 	r.OutPrm(_sys.Ana.Hst.Instrs, "r")
// 	r.Addf("r = %v()", _sys.Ana.Hst.Instrs.New.Ref(x))
// 	for _, instr := range ks.Instrs {
// 		r.Addf("r.Push(x.%v(rng...))", strings.Title(instr))
// 	}
// 	r.Add("return r")
// 	return r
// }
