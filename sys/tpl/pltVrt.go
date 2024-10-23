package tpl

import (
	"sys/k"
	"sys/tpl/atr"
)

type (
	FlePltVrt struct {
		FleBse
		PrtStructFldSet
		PrtPltArng
		PrtPltArngNew
	}
)

func (x *DirPlt) NewVrt() (r *FlePltVrt) {
	r = &FlePltVrt{}
	x.Vrt = r
	r.Name = k.Vrt
	r.Pkg = x.Pkg
	r.StructPtr(r.Name, atr.TypUiStruct)
	r.AddFle(r)
	return r
}
