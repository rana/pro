package tpl

import (
	"strings"
	"sys/k"
	"sys/tpl/atr"
)

type (
	FleRltInstr struct {
		FleRltBse
	}
	FleRltInstrs struct {
		FleBse
		PrtArr
	}
)

func (x *DirRlt) NewInstr() (r *FleRltInstr) {
	r = &FleRltInstr{}
	x.Instr = r
	r.Name = k.Instr
	r.Pkg = x.Pkg
	r.Ifc(r.Name, atr.TypAnaIfc)
	r.AddFle(r)
	return r
}
func (x *DirRlt) NewInstrs() (r *FleRltInstrs) {
	r = &FleRltInstrs{}
	x.Instrs = r
	r.FleBse = *NewArr(x.Instr, &r.PrtArr, x.Instr.Pkg)
	r.AddFle(r)
	return r
}
func (x *FleRltInstr) InitTyp(bse *TypBse) {
	x.FleRltBse.InitTyp(bse)
	x.Typ().Bse().TestPth = append(_sys.Ana.Rlt.Prv.Typ().Bse().TestPth, &TestStp{
		Fst: func(r *PkgFn) {
			r.Add("for _, instr := range tst.RltPrvInstrs {")
		},
		MdlFst: func(r *PkgFn) { r.Add("instr := instr(prv)") },
		Lst:    func(r *PkgFn) { r.Add("}") },
	})
}
func (x *FleRltInstr) InitFld(s *Struct) {
	x.FleRltBse.InitFld(s)
	x.bse.FldPrnt(_sys.Ana.Rlt.Prv)
	x.bse.Fld("Ana", _sys.Ana.Instr)
	x.FldRxs(_sys.Ana.TmeIdx.Rxs)
}
func (x *FleRltInstr) InitTypFn() {
	x.FleRltBse.InitTypFn()
	x.Instr()
	x.Sub(_sys.Ana.TmeIdx.Rx)
	x.Unsub(x.bse, true, true, func(r *TypFn) {
		r.Add("x.Ana.Unsub(x.Id)")
	})
	x.I()
}
func (x *FleRltInstr) Instr() (r *TypFn) {
	r = x.TypFn(k.Instr, x.bse)
	r.OutPrm(_sys.Ana.Instr)
	r.Add("return x.Ana")
	x.MemSigFn(r) // add to interface
	return r
}
func (x *FleRltInstr) I() (r *TypFn) {
	inrvl, name := _sys.Ana.Rlt.Inrvl, strings.Title(k.I)
	// node
	r = x.ElmNodeTypFn(name, "", "", inrvl, func(r *TypFn) {
		r.InPrm(_sys.Bsc.Tme, "dur").LitVal("10")
		if test := inrvl.Test; test != nil {
			test.Import(_sys.Ana.Hst)
			test.Import(_sys.Lng.Pro.Act)
			r.T2.MdlLst.Add("tst.NotNil(t, a.Pkts, \"Pkts\")")
			r.T2.MdlLst.Add("mnr := tst.NewInrvlMnr(ap)")
			r.T2.MdlLst.Add("a.Sub(mnr.Rx, mnr.Id)")
			r.T2.MdlLst.Add("tst.IntegerEql(t, 1, len(a.Rxs), \"Sub Rxs\")")
			r.T2.MdlLst.Add("var actr act.Actr")
			r.T2.MdlLst.Add("vs := actr.RunHst(a.String())")
			r.T2.MdlLst.Addf("eHst := vs[len(vs)-1].(*hst.%v)", r.Node.Title())
			r.T2.MdlLst.Add("if eHst.TmeBnds != nil {")
			r.T2.MdlLst.Add("mnr.StartFor(instr.Instr(), eHst.TmeBnds.Cnt())")
			r.T2.MdlLst.Add("tst.BndsEql(t, eHst.TmeBnds, mnr.Bnds, \"TmeBnds\")")
			r.T2.MdlLst.Add("}")
			r.T2.MdlLst.Add("a.Unsub(mnr.Id)")
			r.T2.MdlLst.Add("tst.IntegerEql(t, 0, len(a.Rxs), \"Unsub Rxs\")")
		}
	})
	r.Addf("r.Pkts = %v()", inrvl.Pkts.New.Ref(x))
	r.Add("x.Sub(r.Rx, r.Id)")
	r.Add("x.Ana.RltSubsMu.Lock()")
	r.Add("x.Ana.RltInrvlMax = x.Ana.RltInrvlMax.Max(dur) // used to load hst data on fst opn")
	r.Add("x.Ana.RltSubsMu.Unlock()")
	r.Add("return r")
	inrvl.Rx(r)
	return r
}
