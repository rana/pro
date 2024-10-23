package tpl

// import (
// 	"sys/k"
// 	"sys/tpl/atr"
// )

// type (
// 	FlePltPrcpSplt struct {
// 		FleBse
// 		PrtStructFldSet
// 		PrtPltArng
// 	}
// )

// func (x *DirPlt) NewPrcpSplt() (r *FlePltPrcpSplt) {
// 	r = &FlePltPrcpSplt{}
// 	x.PrcpSplt = r
// 	r.Name = k.PrcpSplt
// 	r.Pkg = x.Pkg
// 	r.StructPtr(r.Name, atr.TypUiStruct)
// 	r.AddFle(r)
// 	return r
// }
// func (x *FlePltPrcpSplt) InitTyp(bse *TypBse) {

// }
// func (x *FlePltPrcpSplt) InitFld(s *Struct) {
// 	s.Fld("PrcpSplt", _sys.Ana.Hst.PrcpSplt).Atr = atr.Get
// }

// func (x *FlePltPrcpSplt) InitPkgFn() {
// 	x.New()
// }
// func (x *FlePltPrcpSplt) New() (r *PkgFn) {
// 	r = x.PkgFna("NewPrcpSplt", atr.Lng)
// 	r.InPrm(_sys.Ana.Hst.PrcpSplt, "prcpSplt")
// 	r.OutPrm(x, "r")
// 	r.Addf("r = %v{}", x.Typ().Adr(x))
// 	r.Add("r.slf = r")
// 	r.Add("r.Plts = NewPlts()")
// 	r.Add("r.PrcpSplt = prcpSplt")
// 	r.Add("r.Plt(NewStmSplt(*prcpSplt.StmSplts...))") // place plt creation here so that scling works
// 	r.Add("return r")
// 	return r
// }
// func (x *FlePltPrcpSplt) InitTypFn() {

// }
