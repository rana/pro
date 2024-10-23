package tpl

// import (
// 	"sys/k"
// 	"sys/tpl/atr"
// 	"sys/tpl/mod"
// )

// type (
// 	FleRltPrfm struct {
// 		FleRltBse
// 	}
// 	FleRltPrfms struct {
// 		FleBse
// 		PrtArr
// 		PrtArrStrWrt // for PrmWrt
// 		PrtString
// 	}
// )

// func (x *DirRlt) NewPrfm() (r *FleRltPrfm) {
// 	r = &FleRltPrfm{}
// 	x.Prfm = r
// 	r.Name = k.Prfm
// 	r.Pkg = x.Pkg
// 	ifc := r.Ifc(r.Name, atr.TypAnaIfc)
// 	ifc.Atr = ifc.Atr &^ atr.Test
// 	r.AddFle(r)
// 	return r
// }
// func (x *DirRlt) NewPrfms() (r *FleRltPrfms) {
// 	r = &FleRltPrfms{}
// 	x.Prfms = r
// 	r.FleBse = *NewArr(x.Prfm, &r.PrtArr, x.Prfm.Pkg)
// 	r.AddFle(r)
// 	return r
// }

// // func (x *FleRltPrfm) InitTyp(bse *TypBse) {
// // 	x.FleRltBse.InitTyp(bse)
// // 	x.Typ().Bse().TestPth = append(_sys.Ana.Rlt.Stgy.Typ().Bse().TestPth, &TestStp{
// // 		MdlFst: func(r *PkgFn) { r.Add("port := tst.RltStgyPrfmStgyPrfm(stgy)") },
// // 	})
// // }
// // func (x *FleRltPrfmFbr) InitTyp(bse *TypBse) {
// // 	x.Typ().Bse().TestPth = append(_sys.Ana.Rlt.Stgy.Fbr.TestPth, &TestStp{
// // 		MdlFst: func(r *PkgFn) {
// // 			r.Add("portFbr := tst.RltStgyFbrPrfmStgyFbrPrfm(stgyFbr)")
// // 		},
// // 	})
// // }
// func (x *FleRltPrfm) InitFld(s *Struct) {
// 	x.FleRltBse.InitFld(s)
// 	fld := x.bse.FldTyp(_sys.Ana.Prfm)
// 	fld.Atr = atr.Get
// 	fld.Mod = mod.Ptr
// }
