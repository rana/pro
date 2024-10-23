package tpl

// import (
// 	"sys/k"
// 	"sys/tpl/atr"
// )

// type (
// 	FlePltStmSplt struct {
// 		FleBse
// 		PrtPlt
// 		PrtStructFldSet
// 	}
// )

// func (x *DirPlt) NewStmSplt() (r *FlePltStmSplt) {
// 	r = &FlePltStmSplt{}
// 	x.StmSplt = r
// 	r.Name = k.StmSplt
// 	r.Pkg = x.Pkg
// 	r.StructPtr(r.Name, atr.TypUiStruct)
// 	r.AddFle(r)
// 	return r
// }
// func (x *FlePltStmSplt) InitTyp(bse *TypBse) {
// 	x.ImplIfc(_sys.Ana.Vis.Plt.Plt.Typ().(*Ifc))
// }
// func (x *FlePltStmSplt) InitFld(s *Struct) {
// 	s.FldTyp(NewExt("PltBse"))
// 	s.Fld("XWidth", Uint32)
// 	s.Fld("Y", _sys.Ana.Vis.Plt.FltAxisY).Atr = atr.Get
// 	s.Fld("Title", _sys.Bsc.Str).Atr = atr.SetGet
// 	s.Fld("Outlier", _sys.Bsc.Bol).Atr = atr.SetGet
// 	s.FldSlice("sctrs", NewExt("*Sctr"))
// 	s.FldSlice("sctrRndrSegs", NewExt("*StmSpltSctrRndrSeg"))
// }
// func (x *FlePltStmSplt) InitTypFn() {
// 	x.New()
// 	x.StmSplt()
// 	// x.PrfLos()
// }
// func (x *FlePltStmSplt) New() (r *PkgFn) {
// 	r = x.PkgFna("NewStmSplt", atr.Lng)
// 	r.InPrmVariadic(_sys.Ana.Hst.StmSplt, "stmSplts")
// 	r.OutPrm(x, "r")
// 	r.Addf("r = %v{}", x.Typ().Adr(x))
// 	r.Add("r.PltBse = NewPltBse(r)")
// 	r.Add("r.Y = NewFltAxisY()")
// 	r.Add("r.mrgn = Mrgn // glbl mrgn")
// 	r.Add("r.StmSplt(stmSplts...)")
// 	r.Add("return r")
// 	return r
// }
// func (x *FlePltStmSplt) StmSplt() (r *TypFn) {
// 	// x.Import("fmt")
// 	r = x.TypFn(k.StmSplt)
// 	r.InPrmVariadic(_sys.Ana.Hst.StmSplt, "stmSplts")
// 	r.OutPrm(x)
// 	r.Add("for _, stmSplt := range stmSplts {")
// 	r.Add("x.sctrs = append(x.sctrs , &Sctr{")
// 	r.Add("Y: stmSplt.Top,")
// 	r.Add("clr: PrfClr,")
// 	r.Add("radius: uint32(ShpRadius), // glbl ShpRadius")
// 	r.Add("})")
// 	r.Add("x.sctrs = append(x.sctrs , &Sctr{")
// 	r.Add("Y: stmSplt.Btm,")
// 	r.Add("clr: LosClr,")
// 	r.Add("radius: uint32(ShpRadius), // glbl ShpRadius")
// 	r.Add("})")
// 	r.Add("}")
// 	r.Add("//x.Title = str.Str(fmt.Sprintf(\"%v %v\", v.Stm.Name(), v.Stm.Prm()))")
// 	r.Add("return x")
// 	return r
// }
// func (x *FlePltStmSplt) PrfLos() (r *TypFn) {
// 	r = x.TypFn(k.PrfLos)
// 	r.InPrm(_sys.Bsc.Tme.arr, "prfs")
// 	r.InPrm(_sys.Bsc.Tme.arr, "loss")
// 	r.InPrmVariadic(_sys.Ana.Hst.Stm, "stms")
// 	r.OutPrm(x)
// 	r.Add("for _, stm := range stms {")
// 	r.Add("x.Flts(PrfClr, stm.At(prfs))")
// 	r.Add("x.Flts(LosClr, stm.At(loss))")
// 	r.Add("}")
// 	r.Add("return x")
// 	return r
// }
