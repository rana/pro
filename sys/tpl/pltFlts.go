package tpl

// import (
// 	"sys/k"
// 	"sys/tpl/atr"
// )

// type (
// 	FlePltFlts struct {
// 		FleBse
// 	}
// )

// func (x *DirPlt) NewPltFlts() (r *FlePltFlts) {
// 	r = &FlePltFlts{}
// 	x.PltFlts = r
// 	r.Name = k.PltFlts
// 	r.Pkg = x.Pkg
// 	r.StructPtr(r.Name, atr.TypCrt)
// 	r.AddFle(r)
// 	return r
// }

// func (x *FlePltFlts) InitFld(s *Struct) {
// 	x.Import(_sys.Ana.Vis)
// 	s.FldTyp(NewExt("PltBse"))
// 	s.FldTyp(NewExt("*vis.Vis"))
// 	s.FldSlice("Flts", _sys.Bsc.Flt.Arr).Atr = atr.BytLitStrEqlBqSkp
// 	s.Fld("X", NewExt("*UntAxisX")).Atr = atr.BytLitStrEqlBqSkp
// 	s.Fld("Y", NewExt("*AxisY")).Atr = atr.BytLitStrEqlBqSkp
// 	s.Fld("inr", NewExt("vis.Rct")).Atr = atr.BytLitStrEqlBqSkp
// }
// func (x *FlePltFlts) InitTypFn() {
// 	// x.New()
// }
