package tpl

import "sys/k"

type (
	DirBsc struct {
		DirBse
		Str    *FleStr
		Bol    *FleBol
		Flt    *FleFlt
		Unt    *FleUnt
		Int    *FleInt
		Tme    *FleTme
		Bnd    *FleBnd
		FltRng *FleFltRng
		TmeRng *FleTmeRng
	}
)

func (x *DirSys) NewBsc() (r *DirBsc) {
	r = &DirBsc{}
	x.Bsc = r
	r.Pkg = x.Pkg.New(k.Bsc)
	r.NewStr() // IMPORTANT: CALL ORDER IS PRS ORDER
	r.NewBol()
	r.NewFlt()
	r.NewUnt()
	r.NewInt()
	r.NewTme()
	r.NewBnd()
	r.NewFltRng()  // before Flts
	r.NewTmeRng()  // before Tmes
	r.Str.NewArr() // all Arr after all Elm for Actr Test generation
	r.Bol.NewArr()
	r.Flt.NewArr()
	r.Unt.NewArr()
	r.Int.NewArr()
	r.Tme.NewArr()
	r.Bnd.NewArr()
	r.TmeRng.NewArr()
	return r
}
