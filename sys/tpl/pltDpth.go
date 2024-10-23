package tpl

import (
	"sys/k"
	"sys/tpl/atr"
)

type (
	FlePltDpth struct {
		FleBse
		PrtStructFldSet
		PrtPltArng
		PrtPltArngNew
	}
)

func (x *DirPlt) NewDpth() (r *FlePltDpth) {
	r = &FlePltDpth{}
	x.Dpth = r
	r.Name = k.Dpth
	r.Pkg = x.Pkg
	r.StructPtr(r.Name, atr.TypUiStruct)
	r.AddFle(r)
	return r
}
