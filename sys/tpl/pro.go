package tpl

import (
	"sys/k"
)

type (
	DirPro struct {
		DirBse
		Trm *DirTrm
		Cfg *DirCfg
		Xpr *DirXpr
		Act *DirAct
	}
)

func (x *DirLng) NewPro() (r *DirPro) {
	r = &DirPro{}
	x.Pro = r
	r.Pkg = x.Pkg.New(k.Pro)
	r.NewTrm()
	r.NewCfg()
	r.NewXpr()
	r.NewAct()
	return r
}
