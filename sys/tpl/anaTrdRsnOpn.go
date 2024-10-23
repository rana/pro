package tpl

import (
	"sys/k"
	"sys/tpl/atr"
)

type (
	FleAnaTrdRsnOpn struct {
		FleBse
		PrtEnum
	}
)

func (x *DirAna) NewTrdRsnOpn() (r *FleAnaTrdRsnOpn) {
	r = &FleAnaTrdRsnOpn{}
	x.TrdRsnOpn = r
	r.Name = k.TrdRsnOpn
	r.Pkg = x.Pkg
	r.Alias(r.Name, Byte, atr.None)
	r.AddFle(r)
	return r
}
func (x *FleAnaTrdRsnOpn) InitCnst() {
	x.Cnst("ErrMktWeek", "", x)
	x.Cnst("ErrOpnCnd", "", x)
	x.Cnst("NearMktOpn", "", x)
	x.Cnst("NearMktCls", "", x)
	x.Cnst("SpdLrg", "", x)    // HST OR RLT
	x.Cnst("NoCls", "", x)     // HST OPN HAS NO CLS
	x.Cnst("PrvReject", "", x) // RLT RPOVIDER REJECTS OPEN
	x.Cnst("InTrd", "", x)     // RLT IS ALREADY IN TRD
	x.Cnst("FilHstGap", "", x) // RLT IS FILLING WITH HST DATA
	x.Cnst("NoCapital", "", x) // RLT INSUFFICIENT CAPITAL TO OPN TRD
	x.Cnst("PrvErr", "", x)    // GENERAL PROVIDER ERROR
}
