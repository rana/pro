package tpl

// import (
// 	"sys/k"
// 	"sys/tpl/atr"
// 	"sys/tpl/mod"
// )

// type (
// 	FleHstPrfm struct {
// 		FleHstBse
// 		PrtStructRel
// 	}
// 	FleHstPrfms struct {
// 		FleBse
// 		PrtArr
// 		PrtArrStrWrt // for PrmWrt
// 		PrtString
// 		PrtArrFldSel
// 		PrtArrFldSrt
// 		PrtArrFldGrp
// 	}
// )

// func (x *DirHst) NewPrfm() (r *FleHstPrfm) {
// 	r = &FleHstPrfm{}
// 	x.Prfm = r
// 	r.Name = k.Prfm
// 	r.Pkg = x.Pkg
// 	r.Ifc(r.Name, atr.TypAnaIfc)
// 	r.AddFle(r)
// 	return r
// }
// func (x *DirHst) NewPrfms() (r *FleHstPrfms) {
// 	r = &FleHstPrfms{}
// 	x.Prfms = r
// 	r.FleBse = *NewArr(x.Prfm, &r.PrtArr, x.Prfm.Pkg)
// 	r.AddFle(r)
// 	return r
// }

// func (x *FleHstPrfm) InitFld(s *Struct) {
// 	x.FleHstBse.InitFld(s)
// 	fld := x.bse.FldTyp(_sys.Ana.Prfm)
// 	fld.Atr = atr.Get
// 	fld.Mod = mod.Ptr
// 	fld = x.bse.Fld("port", _sys.Ana.Hst.Port)
// 	fld.Atr = atr.Prnt
// }
// func (x *FleHstPrfm) InitTypFn() {
// 	x.FleHstBse.InitTypFn()
// 	x.Port()
// 	x.Dlt()
// 	x.Ana()
// }
// func (x *FleHstPrfm) Port() (r *TypFn) {
// 	r = x.TypFna(k.Port, atr.Lng, x.bse)
// 	r.OutPrm(_sys.Ana.Hst.Port)
// 	r.Add("return x.port")
// 	x.MemSigFn(r)
// 	return r
// }
// func (x *FleHstPrfm) Dlt() (r *TypFn) {
// 	r = x.TypFna(k.Dlt, atr.Lng, x.bse)
// 	r.InPrm(_sys.Ana.Hst.Prfm, "v")
// 	r.OutPrm(_sys.Ana.PrfmDlt)
// 	r.Add("return x.Prfm.Dlt(v.Bse().Prfm)")
// 	x.MemSigFn(r)
// 	return r
// }
// func (x *FleHstPrfm) Ana() (r *TypFn) {
// 	r = x.TypFna("Ana", atr.Lng, x.bse)
// 	r.OutPrm(_sys.Ana.Prfm)
// 	r.Add("return x.Prfm")
// 	x.MemSigFn(r)
// 	return r
// }
