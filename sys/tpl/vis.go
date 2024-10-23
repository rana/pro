package tpl

import "sys/k"

type (
	DirVis struct {
		DirBse
		Fnt *DirFnt
		Clr *DirClr
		Pen *DirPen
		Plt *DirPlt
	}
)

func (x *DirAna) NewVis() (r *DirVis) {
	r = &DirVis{}
	x.Vis = r
	r.Pkg = x.Pkg.New(k.Vis)
	r.NewFnt()
	r.NewClr()
	r.NewPen()
	r.NewPlt()
	return r
}
