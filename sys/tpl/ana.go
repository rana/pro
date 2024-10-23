package tpl

import (
	"sys/k"
)

type (
	DirAna struct {
		DirBse
		Pth       *FleAnaPth
		Prv       *FleAnaPrv
		TrdRsnOpn *FleAnaTrdRsnOpn
		TrdRsnCls *FleAnaTrdRsnCls
		TmeIdx    *FleAnaTmeIdx
		TmeFlt    *FleAnaTmeFlt
		TmeFlts   *FleAnaTmeFlts
		Tic       *FleAnaTic
		Stm       *FleAnaStm
		Instr     *FleAnaInstr
		Trd       *FleAnaTrd
		Trds      *FleAnaTrds
		Prfm      *FleAnaPrfm
		Prfms     *FleAnaPrfms
		PrfmDlt   *FleAnaPrfmDlt
		Port      *FleAnaPort

		Hst *DirHst
		Rlt *DirRlt
		Vis *DirVis
	}
)

func (x *DirSys) NewAna() (r *DirAna) {
	r = &DirAna{}
	x.Ana = r
	r.Pkg = x.Pkg.New(k.Ana)
	r.NewPth()
	r.NewPrv()
	r.NewAnaTmeIdx()
	r.NewAnaTmeFlt()
	r.NewAnaTmeFlts()
	r.NewAnaTic()
	r.NewAnaStm()
	r.NewAnaInstr()
	r.NewAnaTrd()
	r.NewAnaTrds()
	r.NewTrdRsnOpn()
	r.NewTrdRsnCls()
	r.NewPrfm()
	r.NewPrfms()
	r.NewPrfmDlt()
	r.NewPort()

	if Opt.IsHst() { // dirs
		r.NewHst()
		if Opt.IsRlt() { // rlt test requires hst
			r.NewRlt()
		}
		if Opt.IsVis() {
			r.NewVis()
		}
	}
	return r
}
