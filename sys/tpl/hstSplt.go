package tpl

// import (
// 	"sys/k"
// 	"sys/tpl/atr"
// )

// type (
// 	FleHstSplt struct {
// 		FleBse
// 		// PrtLog
// 		// PrtIfc
// 	}
// )

// func (x *DirHst) NewSplt() (r *FleHstSplt) {
// 	r = &FleHstSplt{}
// 	x.Splt = r
// 	r.Name = k.Splt
// 	r.Pkg = x.Pkg
// 	r.StructPtr(r.Name, atr.Typ)
// 	r.AddFle(r)
// 	return r
// }

// // func (x *FleHstSplt) InitTyp(bse *TypBse) {
// // 	x.Typ().Bse().TestPth = append(_sys.Ana.Hst.Port.Typ().Bse().TestPth, &TestStp{
// // 		MdlFst: func(r *PkgFn) { r.Add("splt := port.Splt(0.0)") },
// // 	})
// // }
// func (x *FleHstSplt) InitFld(s *Struct) {
// 	s.Fld("Stgy", _sys.Ana.Hst.Stgy).Atr = atr.Get
// 	s.Fld("Pnt", _sys.Bsc.Flt).Atr = atr.Get | atr.TstZeroSkp
// 	s.Fld("Btm", _sys.Ana.Trd.arr).Atr = atr.Get
// 	s.Fld("Top", _sys.Ana.Trd.arr).Atr = atr.Get
// }
