package tpl

import (
	"sys/k"
	"sys/tpl/atr"
)

type (
	FleHstStgy struct {
		FleHstBse
	}
	FleHstStgys struct {
		FleBse
		PrtArr
		PrtArrStrWrt // for Fbr.PrmWrt
		PrtString
	}
)

func (x *DirHst) NewStgy() (r *FleHstStgy) {
	r = &FleHstStgy{}
	x.Stgy = r
	r.Name = k.Stgy
	r.Pkg = x.Pkg
	r.Ifc(r.Name, atr.TypAnaIfc)
	r.AddFle(r)
	return r
}
func (x *DirHst) NewStgys() (r *FleHstStgys) {
	r = &FleHstStgys{}
	x.Stgys = r
	r.FleBse = *NewArr(x.Stgy, &r.PrtArr, x.Stgy.Pkg)
	r.AddFle(r)
	return r
}

func (x *FleHstStgy) InitTyp(bse *TypBse) {
	x.FleHstBse.InitTyp(bse)
	x.Typ().Bse().TestPth = append(_sys.Ana.Hst.Cnd.Typ().Bse().TestPth, &TestStp{
		MdlFst: func(r *PkgFn) { r.Add("stgy := tst.HstCndStgyLong(cnd, 2.0, 4.0, 60*60, instr)") },
	})
	x.seg = x.NewSeg()
}

func (x *FleHstStgy) InitFld(s *Struct) {
	x.FleHstBse.InitFld(s)
	x.bse.FldPrnt(_sys.Ana.Hst.Cnd) //.Atr = atr.Get
	x.bse.Fld("IsLong", _sys.Bsc.Bol).Atr = atr.Get | atr.TstZeroSkp
	x.bse.Fld("PrfLim", _sys.Bsc.Flt).Atr = atr.Get
	x.bse.Fld("LosLim", _sys.Bsc.Flt).Atr = atr.Get
	x.bse.Fld("DurLim", _sys.Bsc.Tme).Atr = atr.Get
	x.bse.Fld("MinPnlPct", _sys.Bsc.Flt).Atr = atr.Get | atr.TstZeroSkp
	x.bse.Fld("Instr", _sys.Ana.Hst.Instr).Atr = atr.Get
	x.bse.FldSlice("Clss", _sys.Ana.Hst.Cnd)
	x.bse.Fld("Trds", _sys.Ana.Trd.arr).Atr = atr.Get | atr.TstZeroSkp
	// x.bse.Fld("port", _sys.Ana.Hst.Port)
}
func (x *FleHstStgy) InitTypFn() {
	x.FleHstBse.InitTypFn()
	// x.Port()
	// x.MlFtr()
}

// func (x *FleHstStgy) Port() (r *TypFn) { // PortStgy
// 	port, name := _sys.Ana.Hst.Port, strings.Title(k.Port)
// 	r = x.ElmNodeTypFn(name, "Stgy", "", port, func(r *TypFn) {
// 		r.Node.Name = "PortStgy"
// 		r.Node.FldPrnt(x)
// 	})
// 	r.Add("if ana.Cfg.Trc.IsHstPort() {")
// 	r.Addf("sys.Logf(\"%%p %v(%%v)\", r, r.Prm())", r.Node.Full())
// 	r.Add("}")
// 	r.Add("r.BalFstUsd = ana.Cfg.Hst.BalUsd")
// 	r.Add("r.BalLstUsd = ana.Cfg.Hst.BalUsd")
// 	r.Add("r.TrdPct = ana.Cfg.Hst.TrdPct")
// 	r.Add("r.PortBse.AddStgy(x.Slf)")
// 	r.Add("if x.Cnd.Bse().Tmes.Cnt() == 0 {")
// 	r.Add("return r")
// 	r.Add("}")
// 	r.Add("cndTmes := x.Cnd.Bse().Tmes")
// 	r.Add("nxtAvailOpnTme := tme.Min")
// 	r.Add("var opnTme, prvOpnTme tme.Tme")
// 	r.Add("for {")
// 	r.Add("opnTme = tme.Max")
// 	r.Add("idx := cndTmes.SrchIdxEql(nxtAvailOpnTme)")
// 	r.Add("if int(idx) < len(*cndTmes) && (*cndTmes)[idx] < opnTme {")
// 	r.Add("opnTme = (*cndTmes)[idx]")
// 	r.Add("}")
// 	r.Add("if opnTme == prvOpnTme {")
// 	r.Add("nxtAvailOpnTme += tme.Resolution")
// 	r.Add("continue")
// 	r.Add("}")
// 	r.Add("if opnTme == tme.Max {")
// 	r.Add("break")
// 	r.Add("}")
// 	r.Add("trd, rsnFail := x.OpnClsTrd(opnTme, r.BalLstUsd, r.TrdPct)")
// 	r.Add("if trd != nil { // may fail to close due to near mkt opn, near mkt cls, spd lim exceeded")
// 	r.Add("r.CalcOpn(trd, x.Instr.Bse().Ana) // set trd flds")
// 	r.Add("r.CalcCls(trd, x.Instr.Bse().Ana, true) // set trd flds")
// 	r.Add("r.Trds.Push(trd)")
// 	r.Add("x.Trds.Push(trd)")
// 	r.Add("nxtAvailOpnTme = trd.ClsTme")
// 	r.Add("} else {")
// 	r.Add("if rsnFail == ana.NoCls {")
// 	r.Add("break // exit last opn fail to mirror rlt behavior; single ana.NoCls expected")
// 	r.Add("}")
// 	r.Add("nxtAvailOpnTme = opnTme // trd failed to cls; advance to next open time")
// 	r.Add("}")
// 	r.Add("prvOpnTme = opnTme")
// 	r.Add("}")
// 	r.Add("return r")
// 	return r
// }

// func (x *FleHstStgy) MlFtr() (r *TypFn) { // PortStgy
// 	ftr := _sys.Ana.Hst.Ftr
// 	r = x.ElmNodeTypFn("MlFtr", "", "", ftr, func(r *TypFn) {
// 		r.Node.FldPrnt(x)
// 		r.Node.Atr = r.Node.Atr &^ atr.Test // manual test in ml_test.go
// 		r.InPrm(_sys.Ana.Hst.Stms, "stms")
// 		r.InPrm(_sys.Ana.Hst.Stms, "stmsToNorm")
// 	}, x.bse)
// 	r.Add("if stms.Cnt() == 0 && stmsToNorm.Cnt() == 0 {")
// 	r.Add("return r")
// 	r.Add("}")
// 	r.Add("r.Calc()")
// 	r.Add("return r")
// 	x.MemSigFn(r)
// 	return r
// }
