package tpl

import "sys/k"

type (
	DirLng struct {
		DirBse
		Scn *DirScn
		Pro *DirPro
		Jsn *DirJsn
	}
)

func (x *DirSys) NewLng() (r *DirLng) {
	r = &DirLng{}
	x.Lng = r
	r.Pkg = x.Pkg.New(k.Lng)
	r.NewScn()
	r.NewPro()
	r.NewJsn()
	return r
}
