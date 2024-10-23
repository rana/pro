package tpl

import (
	"strings"
	"sys"
	"sys/k"
	"sys/ks"
	"sys/tpl/atr"
)

type (
	FleRltPrv struct {
		FleRltBse
		// PrtLog
		// PrtIfc
	}
	FleRltPrvs struct {
		FleBse
		PrtArr
		PrtArrStrWrt // for PrmWrt
		PrtString
		// PrtLog
		// PrtIfc
	}
)

func (x *DirRlt) NewPrv() (r *FleRltPrv) {
	r = &FleRltPrv{}
	x.Prv = r
	r.Name = k.Prv
	r.Pkg = x.Pkg
	r.Ifc(r.Name, atr.TypAnaIfc)
	r.AddFle(r)
	return r
}
func (x *DirRlt) NewPrvs() (r *FleRltPrvs) {
	r = &FleRltPrvs{}
	x.Prvs = r
	r.FleBse = *NewArr(x.Prv, &r.PrtArr, x.Prv.Pkg)
	r.AddFle(r)
	return r
}
func (x *FleRltPrv) InitTyp(bse *TypBse) {
	x.FleRltBse.InitTyp(bse)
	x.Typ().Bse().TestPth = []*TestStp{&TestStp{
		MdlFst: func(r *PkgFn) { r.Add("prv := rlt.Oan()") },
	}}
}
func (x *FleRltPrv) InitVals(bse *TypBse) {
	bse.Lits = sys.Vs("rlt.oan()")
	bse.Vals = sys.Vs("rlt.Oan()")
}
func (x *FleRltPrv) InitFld(s *Struct) {
	// x.FleRltBse.InitFld(s) // DO NOT CALL BASE
	x.bse.Fld("Slf", x.Typ()).Atr = atr.Slf | atr.TstSkp
}
func (x *FleRltPrv) InitIfc(i *Ifc) {
	i.EmbedIfc(_sys.Ana.Pth)
	var sig *MemSig

	sig = x.MemSig(k.Instr)
	sig.InPrm(_sys.Bsc.Str, "name")
	sig.OutPrm(_sys.Ana.Instr, "r")

	sig = x.MemSig(k.Sub)
	sig.InPrm(_sys.Ana.Instr, "i")

	sig = x.MemSig(k.Unsub)
	sig.InPrm(_sys.Ana.Instr, "i")

	sig = x.MemSig(k.MayTrd)
	sig.OutPrm(_sys.Bsc.Bol)

	sig = x.MemSig(k.OpnTrd)
	sig.InPrm(_sys.Ana.Trd, "t")
	sig.InPrm(_sys.Ana.Instr, "i")
	sig.OutPrm(_sys.Bsc.Bol, "ok")
	sig.OutPrm(_sys.Ana.TrdRsnOpn, "rsn")

	sig = x.MemSig(k.ClsTrd)
	sig.InPrm(_sys.Ana.Trd, "t")
	sig.InPrm(_sys.Ana.Instr, "i")
	sig.OutPrm(_sys.Bsc.Bol, "ok")
}
func (x *FleRltPrv) InitPkgFn() {
	x.Oan()
}
func (x *FleRltPrv) Oan() (r *PkgFn) {
	x.Import(_sys.Ana)
	r = x.NodePkgFn(k.Oan, "", x, func(r *PkgFn) {
		r.Node.FldTyp(NewExt("*ana.Oan"))
	})
	r.Add("r.Oan = ana.PrvOan")
	r.Add("if !ana.Cfg.Test {")
	r.Add("r.Oan.AcntRefresh()") // TODO: MOVE? WHERE?
	r.Add("}")
	r.Add("return r")
	return r
}
func (x *FleRltPrv) InitTypFn() {
	x.FleRltBse.InitTypFn()
	for _, instr := range ks.Instrs {
		x.Instr(instr)
	}
	// x.Opn()
	// x.Instrs()
}

func (x *FleRltPrv) Instr(name string) (r *TypFn) {
	instr, name := _sys.Ana.Rlt.Instr, strings.Title(name)
	// node
	x.Import("sys/k")
	r = x.ElmNodeTypFn(name, "", "", instr, func(r *TypFn) {
		r.InPrmVariadic(_sys.Bsc.TmeRng, "rng")
		if test := instr.Test; test != nil {
			r.T2.MdlLst.Add("mnr := tst.NewInstrMnr(ap)")
			r.T2.MdlLst.Add("a.Sub(mnr.Rx, mnr.Id)")
			r.T2.MdlLst.Add("tst.IntegerEql(t, 1, len(a.Rxs), \"Sub Rxs\")")
			r.T2.MdlLst.Add("mnr.StartFor(a.Ana, a.Ana.HstStm.Cnt())")
			r.T2.MdlLst.Addf("%v(t, a.Ana.HstStm, mnr.Stm(), \"mnr.Stm\")", _sys.Ana.Stm.Typ().Bse().TstEql.Ref(test))
			r.T2.MdlLst.Addf("%v(t, a.Ana.HstStm, a.Ana.RltStm, \"RltStm\")", _sys.Ana.Stm.Typ().Bse().TstEql.Ref(test))
			r.T2.MdlLst.Add("a.Unsub(mnr.Id)")
			r.T2.MdlLst.Add("tst.IntegerEql(t, 0, len(a.Rxs), \"Unsub Rxs\")")
		}
	})
	r.Addf("r.Ana = x.Slf.Instr(k.%vName)", r.Name)
	r.Add("r.Ana.Sub(r.Rx, r.Id)")
	r.Add("return r")
	// rx
	rx := instr.TypFn(k.Rx, r.Node)
	rx.InPrm(_sys.Ana.TmeIdx, "inPkt")
	rx.OutPrmSlice(_sys.Act, "r")
	rx.Add("x.mu.Lock() // pass all through")
	rx.Add("if ana.Cfg.Trc.IsRltInstr() {")
	rx.Addf("sys.Logf(\"%v(%%v).%v %%p inPkt %%v\", x.Prm(), x, inPkt)", r.Node.Full(), rx.Name)
	rx.Add("}")
	rx.Add("for _, rx := range x.Rxs {")
	rx.Addf("r = append(r, %v{Pkt:inPkt, Rx:rx})", _sys.Ana.TmeIdx.Tx.Adr(x))
	rx.Add("}")
	rx.Add("x.mu.Unlock()")
	rx.Add("return r")
	return r
}

// func (x *FleRltPrv) Opn() (r *TypFn) {
// 	// x.Import("fmt")
// 	// x.Import("strings")
// 	// x.Import(_sys)
// 	// x.Import(_sys.Bsc.Tme)
// 	// x.Import(_sys.Bsc.Unt)
// 	r = x.TypFna(k.Opn, atr.Lng, x.bse)
// 	r.OutPrm(x)
// 	// r.Add("stgys := ana.NewStgyPrfms()")
// 	// r.Add("bqProject, bqDataset := sys.Cld().Cfg()")
// 	// r.Add("stgys.CldQryf(\"select * from `%v.%v.rlt`\", bqProject, bqDataset)")
// 	// r.Add("sys.Logf(\"--- Rlt.Opn cnt:%v now:%v\", stgys.Cnt(), tme.Now())")
// 	// r.Add("var b strings.Builder")
// 	// r.Add("b.WriteString(fmt.Sprintf(\"# rlt opn cnt:%v now:%v \\n\", stgys.Cnt(), tme.Now()))")
// 	// r.Add("for n := unt.Zero; n < stgys.Cnt(); n++ {")
// 	// r.Add("sys.Log(stgys.At(n))")
// 	// r.Add("b.WriteString(stgys.At(n).Pth.Unquo())")
// 	// r.Add("b.WriteString(\"\\n\")")
// 	// r.Add("}")
// 	// r.Add("sys.Eval().RunRlt(b.String())")
// 	// r.Add("sys.Log(\"--- Rlt.Opn: end\")")
// 	r.Add("")
// 	r.Add("")
// 	r.Add("")
// 	r.Add("return x.Slf")
// 	x.MemSigFn(r)
// 	return r
// }

// TODO: KEEP? YES
// func (x *FleRltPrvOan) Instrs() (r *TypFn) {
// 	r = x.TypFn(k.Instrs)
// 	r.OutPrm(_sys.Ana.Rlt.Instrs, "r")
// 	r.Addf("r = %v()", _sys.Ana.Rlt.Instrs.New.Ref(x))
// 	for _, instr := range ks.Instrs {
// 		r.Addf("r.Push(x.%v())", strings.Title(instr))
// 	}
// 	r.Add("return r")
// 	return r
// }
