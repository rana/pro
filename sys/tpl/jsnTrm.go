package tpl

import (
	"sys/k"
)

type (
	DirJsnTrm struct {
		DirBse
		Prs  *FleJsnPrs
		Trmr *FleJsnTrmr
	}
)

func (x *DirJsn) NewJsnTrm() (r *DirJsnTrm) {
	r = &DirJsnTrm{}
	x.Trm = r
	r.Pkg = x.Pkg.New(k.Trm)
	r.NewPrs()
	r.NewTrmr()
	return r
}
