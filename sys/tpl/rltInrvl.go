package tpl

import (
	"strings"
	"sys/k"
	"sys/ks"
	"sys/tpl/atr"
)

type (
	FleRltInrvl struct {
		FleRltBse
	}
	FleRltInrvls struct {
		FleBse
		PrtArr
	}
)

func (x *DirRlt) NewInrvl() (r *FleRltInrvl) {
	r = &FleRltInrvl{}
	x.Inrvl = r
	r.Name = k.Inrvl
	r.Pkg = x.Pkg
	r.Ifc(r.Name, atr.TypAnaIfc)
	r.AddFle(r)
	return r
}
func (x *DirRlt) NewInrvls() (r *FleRltInrvls) {
	r = &FleRltInrvls{}
	x.Inrvls = r
	r.FleBse = *NewArr(x.Inrvl, &r.PrtArr, x.Inrvl.Pkg)
	r.AddFle(r)
	return r
}
func (x *FleRltInrvl) InitTyp(bse *TypBse) {
	x.FleRltBse.InitTyp(bse)
	x.Typ().Bse().TestPth = append(_sys.Ana.Rlt.Instr.Typ().Bse().TestPth, &TestStp{
		MdlFst: func(r *PkgFn) { r.Add("inrvl := tst.RltInstrInrvlI(instr, 10)") },
	})
}
func (x *FleRltInrvl) InitFld(s *Struct) {
	x.FleRltBse.InitFld(s)
	x.bse.FldPrnt(_sys.Ana.Rlt.Instr)
	x.FldRxs(_sys.Bsc.Bnd.Rxs)
	x.Pkts = _sys.Ana.TmeIdx.arr
	x.bse.Fld("Pkts", x.Pkts).Atr = atr.TstZeroSkp
}
func (x *FleRltInrvl) InitTypFn() {
	x.FleRltBse.InitTypFn()
	x.Sub(_sys.Bsc.Bnd.Rx)
	x.Unsub(x.bse, true, true, func(r *TypFn) {
		r.Addf("x.%v.Unsub(x.Id)", x.bse.Prnt().Name)
	})
	for _, side := range ks.Sides {
		x.Side(side)
	}
}

// rolling inrvl
// interval receives price packets and heartbeat packets
// interval sends price packets only
// process each time interval its own processing tier
// wait for entire graph completion of each interval before processing the next interval
// first pkt must be a price pkt
func (x *FleRltInrvl) Rx(nodeFn *TypFn) (r *TypFn) {
	x.Import(_sys.Bsc.Tme)
	x.Import("sys/err")
	r = x.TypFn(k.Rx, nodeFn.Node)
	r.InPrm(_sys.Ana.TmeIdx, "inPkt")
	r.OutPrmSlice(_sys.Act, "r")
	r.Add("x.mu.Lock()")
	r.Add("defer x.mu.Unlock()")
	r.Add("x.Pkts.Que(inPkt)")
	r.Add("if ana.Cfg.Trc.IsRltInrvl() {")
	r.Addf("sys.Logf(\"%v(%%v).%v %%p inPkt %%v\", x.Prm(), x, inPkt)", nodeFn.Node.Full(), r.Name)
	r.Add("}")
	r.Add("for x.Pkts.Cnt() != 0 && x.Pkts.Fst().Idx == unt.Max { // drain leading hearbeat tics")
	r.Add("x.Pkts.Dque()")
	r.Add("}")
	r.Add("if x.Pkts.Cnt() < 2 { // all starting heartbeats drained, or fst tic")
	r.Add("return nil")
	r.Add("}")
	r.Add("start := x.Pkts.At(0)")
	r.Add("endTme := start.Tme + x.Dur")
	r.Add("endLim := start.Idx + 1")
	r.Add("for m := unt.One; m < x.Pkts.Cnt(); m++ {")
	r.Add("cur := x.Pkts.At(m)")
	r.Add("if cur.Tme < endTme {")
	r.Addf("if cur.Idx != %v {", _sys.Bsc.Unt.Max.Ref(x))
	r.Add("endLim = cur.Idx + 1 // set end price idx")
	r.Add("}")
	r.Add("} else { // heartbeat or price advance the inrvl")
	r.Addf("outPkt := %v{Idx: start.Idx, Lim: endLim}", _sys.Bsc.Bnd.Ref(x))
	r.Add("ts := x.Instr.Instr().RltStm.Tmes")
	r.Addf("if ts.At(outPkt.Lim-1).Sub(ts.At(outPkt.Idx)) >= x.Dur {")
	r.Add("err.Panicf(\"INVALID %v INRVL CALC (calc:%v outPkt:%v)\", x.Dur, ts.At(outPkt.Lim-1).Sub(ts.At(outPkt.Idx)), outPkt)")
	r.Add("}")
	r.Add("for _, rx := range x.Rxs {")
	r.Addf("r = append(r, %v{Pkt:outPkt, Rx:rx})", _sys.Bsc.Bnd.Tx.Adr(x))
	r.Add("}")
	r.Add("x.Pkts.Dque() // deque fst non-heartbeat for rolling inrvl behavior")
	r.Add("break")
	r.Add("}")
	r.Add("}")
	r.Add("return r")

	return r
}
func (x *FleRltInrvl) Side(name string) (r *TypFn) {
	side, name := _sys.Ana.Rlt.Side, strings.Title(name)
	// node
	r = x.ElmNodeTypFn(name, "", "", side, func(r *TypFn) {
		if test := side.Test; test != nil {
			test.Import(_sys.Ana.Hst)
			test.Import(_sys.Lng.Pro.Act)
			r.T2.MdlLst.Add("mnr := tst.NewSideMnr(ap)")
			r.T2.MdlLst.Add("a.Sub(mnr.Rx, mnr.Id)")
			r.T2.MdlLst.Add("tst.IntegerEql(t, 1, len(a.Rxs), \"Sub Rxs\")")
			r.T2.MdlLst.Add("var actr act.Actr")
			r.T2.MdlLst.Add("vs := actr.RunHst(a.String())")
			r.T2.MdlLst.Addf("eHst := vs[len(vs)-1].(*hst.%v)", r.Node.Title())
			r.T2.MdlLst.Add("if eHst.ValBnds != nil {")
			r.T2.MdlLst.Add("mnr.StartFor(instr.Instr(), eHst.ValBnds.Cnt())")
			r.T2.MdlLst.Add("tst.IntegerEql(t, len(*eHst.ValBnds), len(mnr.Fltss), \"ValBnds Cnt\")")
			r.T2.MdlLst.Add("for n, valBnd := range *eHst.ValBnds {")
			r.T2.MdlLst.Add("tst.FltsEql(t, eHst.Vals.In(valBnd.Idx, valBnd.Lim), mnr.Fltss[n], n, \"Mnr Vals\")")
			r.T2.MdlLst.Add("}")
			r.T2.MdlLst.Add("}")
			r.T2.MdlLst.Add("a.Unsub(mnr.Id)")
			r.T2.MdlLst.Add("tst.IntegerEql(t, 0, len(a.Rxs), \"Unsub Rxs\")")
		}
	})
	r.Add("x.Sub(r.Rx, r.Id)")
	r.Add("return r")
	side.Rx(r)
	return r
}
