package tpl

// import (
// 	"sys/k"
// 	"sys/tpl/atr"
// )

// type (
// 	FleHstFtr struct {
// 		FleHstBse
// 	}
// 	FleHstFtrs struct {
// 		FleBse
// 		PrtArr
// 		PrtArrStrWrt // for Fbr.PrmWrt
// 		PrtString
// 	}
// )

// func (x *DirHst) NewFtr() (r *FleHstFtr) {
// 	r = &FleHstFtr{}
// 	// x.Ftr = r
// 	r.Name = k.Ftr
// 	r.Pkg = x.Pkg
// 	r.Ifc(r.Name, atr.TypAnaIfc)
// 	r.AddFle(r)
// 	return r
// }
// func (x *DirHst) NewFtrs() (r *FleHstFtrs) {
// 	r = &FleHstFtrs{}
// 	// x.Ftrs = r
// 	// r.FleBse = *NewArr(x.Ftr, &r.PrtArr, x.Ftr.Pkg)
// 	r.AddFle(r)
// 	return r
// }
// func (x *FleHstFtr) InitTyp(bse *TypBse) {
// 	x.FleHstBse.InitTyp(bse)
// 	// x.Typ().Bse().TestPth = append(_sys.Ana.Hst.Stgy.Typ().Bse().TestPth, &TestStp{
// 	// 	MdlFst: func(r *PkgFn) { r.Add("ftr := tst.HstSideStmRteProLst(side)") },
// 	// })
// }
// func (x *FleHstFtr) InitFld(s *Struct) {
// 	x.FleHstBse.InitFld(s)
// 	// x.Import("unsafe")
// 	x.bse.FldSlice("Lbls", Float32) // ml labels
// 	// x.bse.Fld("Ftrs2", NewExt("*mat.Dense"))
// 	// x.Import("gonum.org/v1/gonum/mat")
// 	// x.bse.Fld("Ftrs3", NewExt("*blas32.General"))
// 	// x.Import("gonum.org/v1/gonum/blas/blas32")
// 	// x.bse.FldSlice("Ftrs", NewExt("[]float32"))  // ml feature columns
// 	x.bse.FldSlice("Ftrs", _sys.Bsc.Flt.arr) // ml feature columns
// 	x.bse.FldSlice("FtrNames", String)       // ml feature column names
// }

// func (x *FleHstFtr) InitTypFn() {
// 	x.FleHstBse.InitTypFn()
// 	// x.SavNpz()
// }

// // func (x *FleHstFtr) SavNpz() (r *TypFn) {
// // 	x.Import("path/filepath")
// // 	x.Import("sys/fs")
// // 	x.Import("gonum.org/v1/gonum/mat")
// // 	x.Import(_sys.Ana)
// // 	r = x.TypFn("SavNpz", x.bse)
// // 	r.OutPrm(x)
// // 	// r.Add("filepath := fs.WorkingDir(\"ftrs.npz\")")
// // 	// r.Add("var idx uint64")
// // 	// r.Add("mtrxData := make([]float64, len(*x.Ftrs)*len(*x.Ftrs[0]))")
// // 	r.Add("")
// // 	r.Add("")
// // 	r.Add("")
// // 	r.Add("")
// // 	r.Add("")
// // 	r.Add("")
// // 	r.Add("")
// // 	r.Add("")
// // 	r.Add("")
// // 	r.Add("")
// // 	r.Add("fs.EnsureDir(ana.Cfg.MlPth)")
// // 	r.Add("path := filepath.Join(ana.Cfg.MlPth, \"ftrs.npz\")")
// // 	r.Add("fs.Npz(path, []string{ \"y\", \"x\", \"col_names\"}, x.Lbls, x.Ftrs2, x.FtrNames)")
// // 	r.Add("return x.Slf")
// // 	x.MemSigFn(r)
// // 	return r
// // }
