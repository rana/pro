package tpl

import (
	"sys/k"
	"sys/tpl/atr"
)

type (
	DirRlt struct {
		DirBse
		// elm
		Prv   *FleRltPrv
		Instr *FleRltInstr
		Inrvl *FleRltInrvl
		Side  *FleRltSide
		Stm   *FleRltStm
		Cnd   *FleRltCnd
		Stgy  *FleRltStgy
		// Port  *FleRltPort
		// Prfm  *FleRltPrfm
		// elm arr
		Prvs   *FleRltPrvs
		Instrs *FleRltInstrs
		Inrvls *FleRltInrvls
		Sides  *FleRltSides
		Stms   *FleRltStms
		Cnds   *FleRltCnds
		Stgys  *FleRltStgys
		// Ports  *FleRltPorts
		// Prfms  *FleRltPrfms
	}
	FleRltBse struct {
		FleNodeBse
		Rxs  *Map
		Pkts *Arr
	}
	FleRlt interface {
		RltBse() *FleRltBse
	}
)

func (x *DirAna) NewRlt() (r *DirRlt) {
	r = &DirRlt{}
	x.Rlt = r
	r.Pkg = x.Pkg.New(k.Rlt)
	// elm
	r.NewPrv()
	r.NewInstr()
	r.NewInrvl()
	r.NewSide()
	r.NewStm()
	r.NewCnd()
	r.NewStgy()
	// r.NewPort()
	// r.NewPrfm()
	// elm arr
	r.NewPrvs()
	r.NewInstrs()
	r.NewInrvls()
	r.NewSides()
	r.NewStms()
	r.NewCnds()
	r.NewStgys()
	// r.NewPorts()
	// r.NewPrfms()
	return r
}
func (x *FleRltBse) RltBse() *FleRltBse { return x }
func (x *FleRltBse) InitTyp(bse *TypBse) {
	x.bse = x.StructPtr(bse.Name+"Bse", atr.TypAnaBse)
	x.bse.Fle = x
	x.bse.ifc, _ = x.Typ().(*Ifc)
	bse.bse = x.bse
}
func (x *FleRltBse) InitFld(s *Struct) {
	x.bse.Fld("mu", Mutex)
	x.bse.Fld("Id", Uint32)
	x.bse.Fld("Slf", x.Typ()).Atr = atr.Slf | atr.TstSkp
}
func (x *FleRltBse) InitIfc(i *Ifc) {
	if i != nil { // IFC MEMS ADDED FROM NODE CALL
		i.EmbedIfc(_sys.Ana.Pth)
	}
}
func (x *FleRltBse) InitTypFn() {
	x.FnBse()
}
func (x *FleRltBse) FnBse() (r *TypFn) {
	r = x.TypFn(k.Bse, x.bse)
	r.OutPrm(x.bse)
	r.Add("return x")
	x.MemSigFn(r) // add to interface
	return r
}

func (x *FleRltBse) FldRxs(rxs *Map) {
	x.bse.Fld("Rxs", rxs).Atr = atr.BytLitStrEqlBqSkp // avoid "sys/ana" import in side.gen_test.go
	x.Rxs = rxs
}

// func (x *FleRltBse) FldPkts() {
// 	prnt := x.bse.Prnt()
// 	prntFle := prnt.Typ.Bse().Fle.(FleRlt).RltBse()
// 	sys.Log("prntFle.Rxs.Val", prntFle.Rxs.Val.PkgTypTitle())
// 	x.Pkts = prntFle.Rxs.Val.Bse().arr
// 	x.bse.Fld("Pkts", x.Pkts).Atr = atr.TstZeroSkp
// }
func (x *FleRltBse) Sub(rx Typ, slot ...bool) (r *TypFn) {
	x.Import(_sys)
	if x.bse != nil {
		r = x.TypFn(k.Sub, x.bse)
	} else {
		r = x.TypFn(k.Sub)
	}
	r.InPrm(rx, "rx")
	r.InPrm(Uint32, "id")
	var slotStr string
	if len(slot) != 0 && slot[0] {
		r.InPrmVariadic(Uint32, "slot")
		slotStr = "uSlot"
		r.Addf("var %v uint32", slotStr)
		r.Add("if len(slot) > 0 {")
		r.Addf("%v = slot[0]", slotStr)
		r.Add("}")
	} else {
		slotStr = "0"
	}
	r.Add("x.mu.Lock()")
	r.Addf("x.Rxs[sys.Uint64(id, %v)] = rx", slotStr)
	r.Add("x.mu.Unlock()")
	if x.bse != nil {
		x.MemSigFn(r) // add to interface
	}
	return r
}
func (x *FleRltBse) Unsub(rxr Typ, ifc, slot bool, inr func(*TypFn)) (r *TypFn) {
	x.Import(_sys)
	r = x.TypFn(k.Unsub, rxr)
	r.InPrm(Uint32, "id")
	var slotStr string
	if slot {
		r.InPrmVariadic(Uint32, "slot")
		slotStr = "uSlot"
		r.Addf("var %v uint32", slotStr)
		r.Add("if len(slot) > 0 {")
		r.Addf("%v = slot[0]", slotStr)
		r.Add("}")
	} else {
		slotStr = "0"
	}
	r.Add("x.mu.Lock()")
	r.Addf("delete(x.Rxs, sys.Uint64(id, %v))", slotStr)
	r.Add("if len(x.Rxs) == 0 {")
	inr(r)
	r.Add("}")
	r.Add("x.mu.Unlock()")
	if ifc {
		x.MemSigFn(r) // add to interface
	}
	return r
}
func (x *FleRltBse) DstToInstr(node *Struct) (r *TypFn) {
	r = x.TypFn(k.DstToInstr, node) // USED BY STGY
	r.OutPrm(Int)
	for _, fld := range node.Flds {
		if fld.IsPrnt() {
			if fld.Typ == _sys.Ana.Rlt.Side.Typ() {
				r.Addf("return 3")
			} else {
				r.Addf("return x.%v.%v() + 1", fld.Name, r.Name)
			}
			break
		}
	}
	return r
}
