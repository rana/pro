package tpl

import "sys/k"

type (
	DirTrc struct {
		DirBse
		Opt *FleTrcOpt
	}
)

func (x *DirSys) NewOpt() (r *DirTrc) {
	r = &DirTrc{}
	x.Trc = r
	r.Pkg = x.Pkg.New(k.Trc)
	r.NewTrc()
	return r
}
