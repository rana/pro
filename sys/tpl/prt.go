package tpl

import "reflect"

type (
	Prt interface {
		RegPrt(p Prt, f Fle) // pass most derived for derived methods to be called
		InitPrtFle()
		InitPrtTyp()
		InitPrtFld()
		InitPrtIfc()
		InitPrtCnst()
		InitPrtTst(tst *FleTst)
		InitPrtVar()
		InitPrtPkgFn()
		InitPrtTypFn()
		InitPrtTrm(trmr *FleTrmr)
		InitPrtXpr(xprr *FleXprr)
		InitPrtAct(actr *FleActr)
	}
	PrtBse struct {
		Fle Fle
		f   *FleBse
		t   *TypBse
	}
	Prts []Prt
)

func (x *Prts) AddPrt(p Prt) { *x = append(*x, p) }
func (x *Prts) GetPrt(v interface{}) Prt {
	if v != nil {
		rPrtType := reflect.TypeOf(v).Elem()
		// fmt.Println("+ rPrtType", rPrtType.Name())
		for _, p := range *x {
			rCurPrtType := reflect.TypeOf(p).Elem()
			// fmt.Println(" > rCurPrtType", rCurPrtType.Name(), rPrtType == rCurPrtType)
			if rPrtType == rCurPrtType {
				return p
			}
		}
	}
	return nil
}
func (x *Prts) InitFle() {
	for _, p := range *x {
		p.InitPrtFle()
	}
}
func (x *Prts) InitPrtTyp() {
	for _, p := range *x {
		p.InitPrtTyp()
	}
}
func (x *Prts) InitPrtFld() {
	for _, p := range *x {
		p.InitPrtFld()
	}
}
func (x *Prts) InitPrtIfc() {
	for _, p := range *x {
		p.InitPrtIfc()
	}
}
func (x *Prts) InitPrtCnst() {
	for _, p := range *x {
		p.InitPrtCnst()
	}
}
func (x *Prts) InitPrtTst(tst *FleTst) {
	for _, p := range *x {
		p.InitPrtTst(tst)
	}
}
func (x *Prts) InitPrtVar() {
	for _, p := range *x {
		p.InitPrtVar()
	}
}
func (x *Prts) InitPrtPkgFn() {
	for _, p := range *x {
		p.InitPrtPkgFn()
	}
}
func (x *Prts) InitPrtTypFn() {
	for _, p := range *x {
		p.InitPrtTypFn()
	}
}
func (x *Prts) InitPrtTrm(trmr *FleTrmr) {
	for _, p := range *x {
		p.InitPrtTrm(trmr)
	}
}
func (x *Prts) InitPrtXpr(xprr *FleXprr) {
	for _, p := range *x {
		p.InitPrtXpr(xprr)
	}
}
func (x *Prts) InitPrtAct(actr *FleActr) {
	for _, p := range *x {
		p.InitPrtAct(actr)
	}
}
func (x *PrtBse) RegPrt(p Prt, f Fle) {
	f.Bse().AddPrt(p)
	x.Fle = f
	x.f = f.Bse()
	x.t = f.Typ().Bse()
}
func (x *PrtBse) InitPrtFle()              {}
func (x *PrtBse) InitPrtTyp()              {}
func (x *PrtBse) InitPrtFld()              {}
func (x *PrtBse) InitPrtIfc()              {}
func (x *PrtBse) InitPrtCnst()             {}
func (x *PrtBse) InitPrtTst(tst *FleTst)   {}
func (x *PrtBse) InitPrtVar()              {}
func (x *PrtBse) InitPrtPkgFn()            {}
func (x *PrtBse) InitPrtTypFn()            {}
func (x *PrtBse) InitPrtTrm(trmr *FleTrmr) {}
func (x *PrtBse) InitPrtXpr(xprr *FleXprr) {}
func (x *PrtBse) InitPrtAct(actr *FleActr) {}
