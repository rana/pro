package tpl

import (
	"sys/k"
	"sys/tpl/atr"
)

type (
	FlePltHrz struct {
		FleBse
		PrtStructFldSet
		PrtPltArng
		PrtPltArngNew
	}
)

func (x *DirPlt) NewHrz() (r *FlePltHrz) {
	r = &FlePltHrz{}
	x.Hrz = r
	r.Name = k.Hrz
	r.Pkg = x.Pkg
	r.StructPtr(r.Name, atr.TypUiStruct)
	r.AddFle(r)
	return r
}
