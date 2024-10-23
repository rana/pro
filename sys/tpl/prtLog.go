package tpl

// import (
// 	"sys/k"
// 	"sys/tpl/atr"
// )

// type (
// 	PrtLog struct {
// 		PrtBse
// 	}
// )

// func (x *PrtLog) InitPrtTypFn() {
// 	x.Log()
// }
// func (x *PrtLog) Log() (r *TypFn) {
// 	x.f.Import(_sys)
// 	var rxr Typ
// 	var ifc *Ifc
// 	var isRxrIfc bool
// 	var suffix string
// 	if ifc, isRxrIfc = x.f.Typ().(*Ifc); isRxrIfc {
// 		rxr = ifc.bse
// 		suffix = ".Slf"
// 	} else {
// 		rxr = x.f.Typ()
// 	}
// 	r = x.f.TypFna(k.Log, atr.Lng, rxr)
// 	// r.InPrmVariadic(_sys.Bsc.Str, "vs")
// 	r.InPrmVariadic(_sys.Ifc, "vs")
// 	r.OutPrm(x.t)
// 	r.Add("if len(vs) == 0 {")
// 	r.Addf("sys.Log(x)")
// 	r.Add("} else {")
// 	// r.Addf("sys.Log(vs[0].Unquo(), x)")
// 	r.Add("sys.LogIfc(append(vs, x)...)")
// 	r.Add("}")
// 	r.Addf("return x%v", suffix)
// 	if isRxrIfc {
// 		x.f.MemSigFn(r, ifc)
// 	}
// 	return r
// }
