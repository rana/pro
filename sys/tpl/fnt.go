package tpl

import (
	"sys/k"
	"sys/tpl/atr"
)

type (
	DirFnt struct {
		DirBse
		Fnt *FleFnt
	}
	FleFnt struct {
		FleBse
	}
)

func (x *DirVis) NewFnt() (r *DirFnt) { // dir
	r = &DirFnt{}
	x.Fnt = r
	r.Pkg = x.Pkg.New(k.Fnt)
	r.NewFnt()
	return r
}
func (x *DirFnt) NewFnt() (r *FleFnt) { // fle
	r = &FleFnt{}
	x.Fnt = r
	r.Name = k.Fnt
	r.Pkg = x.Pkg
	r.StructPtr(r.Name, atr.TypUiStruct)
	r.AddFle(r)
	return r
}

func (x *FleFnt) InitFld(s *Struct) {
	x.Import("sync")
	x.Import("golang.org/x/image/font")
	s.Fld("Face", NewExt("font.Face"))
	s.Fld("Mu", NewExt("sync.Mutex"))
}
