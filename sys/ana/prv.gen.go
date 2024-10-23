package ana

import (
	"sys/bsc/bol"
	"sys/bsc/flt"
	"sys/bsc/str"
)

type (
	Prv interface {
		AcntRefresh() (balance flt.Flt)
		Instr(name str.Str) *Instr
		Sub(i *Instr)
		Unsub(i *Instr)
		LoadHst(i *Instr)
		OpnTrd(trd *Trd, i *Instr) (ok bol.Bol, rsn TrdRsnOpn)
		ClsTrd(trd *Trd, i *Instr) (ok bol.Bol)
		CalcOpn(trd *Trd, i *Instr)
		CalcCls(trd *Trd, i *Instr)
	}
)
