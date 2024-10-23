package tpl

import "sys/k"

type (
	DirTrm struct {
		DirBse
		Prs  *FlePrs
		Trmr *FleTrmr
	}
)

func (x *DirPro) NewTrm() (r *DirTrm) {
	r = &DirTrm{}
	x.Trm = r
	r.Pkg = x.Pkg.New(k.Trm)
	r.Trmr = r.NewTrmr()
	r.Prs = r.NewPrs()
	return r
}
