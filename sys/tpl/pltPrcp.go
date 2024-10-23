package tpl

// import (
// 	"sys/k"
// 	"sys/tpl/atr"
// )

// type (
// 	FlePltPrcp struct {
// 		FleBse
// 		PrtStructFldSet
// 		PrtPltArng
// 	}
// )

// func (x *DirPlt) NewPrcp() (r *FlePltPrcp) {
// 	r = &FlePltPrcp{}
// 	x.Prcp = r
// 	r.Name = k.Prcp
// 	r.Pkg = x.Pkg
// 	r.StructPtr(r.Name, atr.TypUiStruct)
// 	r.AddFle(r)
// 	return r
// }
// func (x *FlePltPrcp) InitTyp(bse *TypBse) {

// }
// func (x *FlePltPrcp) InitFld(s *Struct) {
// 	s.Fld("Prcp", _sys.Ana.Hst.Prcp).Atr = atr.Get
// 	s.FldSlice("Penss", _sys.Ana.Vis.Pen.Pen.arr)
// }

// func (x *FlePltPrcp) InitPkgFn() {
// 	x.New()
// }
// func (x *FlePltPrcp) New() (r *PkgFn) {
// 	r = x.PkgFna("NewPrcp", atr.Lng)
// 	r.InPrm(_sys.Ana.Hst.Prcp, "prcp")
// 	r.OutPrm(x, "r")
// 	r.Addf("r = %v{}", x.Typ().Adr(x))
// 	r.Addf("r.slf = r")
// 	r.Add("r.Plts = NewPlts()")
// 	r.Add("r.Prcp = prcp")
// 	// r.Add("for m, stmFbrs := range r.Prcp.StmFbrss {")
// 	// r.Add("hrz := NewHrz()")
// 	// r.Add("r.Plt(hrz)")
// 	// r.Add("for _, stmFbr := range *stmFbrs {")
// 	// r.Add("stm := NewStm()")
// 	// r.Add("hrz.Plt(stm)")
// 	// r.Add("stm.StmFbr(r.Penss[m], stmFbr)")
// 	// r.Add("}")
// 	// r.Add("}")
// 	r.Add("return r")
// 	return r
// }
// func (x *FlePltPrcp) InitTypFn() {
// 	x.Pens()
// }
// func (x *FlePltPrcp) Pens() (r *TypFn) {
// 	r = x.TypFn("Pens")
// 	r.InPrm(_sys.Ana.Vis.Pen.Pen.arr, "v")
// 	r.OutPrm(x)
// 	r.Add("x.Penss = append(x.Penss, v)")
// 	r.Add("return x")
// 	return r
// }
