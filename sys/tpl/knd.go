package tpl

import (
	"sys/k"
	"sys/tpl/atr"
)

type (
	FleKnd struct {
		FleBse
		PrtEnum
	}
)

func (x *DirXpr) NewKnd() (r *FleKnd) {
	r = &FleKnd{}
	x.Knd = r
	r.Name = k.Knd
	r.Pkg = x.Pkg.New(r.Name)
	r.Alias(k.Knd, Byte, atr.None)
	r.AddFle(r)
	return r
}

func (x *FleKnd) InitCnst() {
	for _, f := range _sys.Fles {
		if f.Typ() != nil && f.Typ().Bse().IsXpr() {
			f.Typ().Bse().Knd = x.Cnst(f.Typ().PkgTypTitle(), "")
		}
	}
}

// func (x *FleKnd) Key() (r *TypFn) {
// 	r = x.TypFn("Key")
// 	r.InPrmSlice(Byte, "b")
// 	r.OutPrmSlice(Byte)
// 	r.Lines.Add("return append([]byte{byte(x)}, b...)")
// 	return r
// }
// func (x *FleKnd) KeyPrefix() (r *TypFn) {
// 	r = x.TypFn("KeyPrefix")
// 	r.InPrm(Byte, "prefix")
// 	r.InPrmSlice(Byte, "b")
// 	r.OutPrmSlice(Byte)
// 	r.Lines.Add("return append([]byte{byte(x), prefix}, b...)")
// 	return r
// }
