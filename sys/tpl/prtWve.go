package tpl

// type (
// 	PrtWve struct {
// 		PrtBse
// 		Wve *Ifc
// 	}
// )

// func NewWve(elmFle, fbrFle Fle, p *PrtWve, pkg ...*Pkg) (r *FleBse) {
// 	r = &FleBse{}
// 	r.Name = fmt.Sprintf("%vWve", fbrFle.Typ().Bse().elm.Camel())
// 	if len(pkg) != 0 {
// 		r.Pkg = pkg[0]
// 	} else {
// 		r.Pkg = fbrFle.Bse().Pkg.NewFromParnt(r.Name)
// 	}
// 	p.Wve = r.Typs.Wve(r.Name, elmFle, fbrFle, r.Pkg)
// 	return r
// }
// func (x *PrtWve) InitPrtFld() {
// 	x.Wve.bse.Fldf("Slf", x.Wve)
// 	x.Wve.bse.Fldf(x.Wve.elm.arr.Camel(), x.Wve.elm.arr) //.Atr = atr.TstZeroSkp
// }

// func (x *PrtWve) InitPrtTypFn() {
// 	x.bse()
// 	x.arr()
// }
// func (x *PrtWve) bse() (r *TypFn) {
// 	r = x.f.TypFn(k.Bse, x.Wve.bse) // USE MTHD FOR IFC
// 	r.OutPrm(x.Wve.bse)
// 	r.Add("return x")
// 	x.f.MemSigFn(r) // add to interface
// 	return r
// }
// func (x *PrtWve) arr() (r *TypFn) {
// 	r = x.f.TypFn(x.Wve.elm.arr.Title(), x.Wve.bse)
// 	r.OutPrm(x.Wve.elm.arr)
// 	r.Addf("return x.%v", x.Wve.elm.arr.Camel())
// 	x.f.MemSigFn(r) // add to interface
// 	return r
// }
