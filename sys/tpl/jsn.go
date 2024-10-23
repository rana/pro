package tpl

import "sys/k"

type (
	DirJsn struct {
		DirBse
		Trm  *DirJsnTrm
		Jsnr *FleJsnr
	}
)

func (x *DirLng) NewJsn() (r *DirJsn) {
	r = &DirJsn{}
	x.Jsn = r
	r.Pkg = x.Pkg.New(k.Jsn)
	r.NewJsnTrm()
	r.NewJsnr()
	return r
}
