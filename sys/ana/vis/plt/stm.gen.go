package plt

import (
	"sys/ana/hst"
	"sys/ana/vis/clr"
	"sys/ana/vis/pen"
	"sys/bsc/flt"
	"sys/bsc/str"
	"sys/bsc/tme"
	"sys/bsc/unt"
)

type (
	Stm struct {
		TmeFltPltBse
		Title       str.Str
		stmStks     []*StmStk
		stmBnds     []*StmBnd
		cndStks     []*CndStk
		hrzLns      []*HrzLn
		vrtLns      []*VrtLn
		hrzBnds     []*HrzBnd
		vrtBnds     []*VrtBnd
		stmStkRndrs []*StmStkRndrSeg
		stmBndRndrs []*StmBndRndrSeg
		cndStkRndrs []*CndStkRndrSeg
		hrzLnRndrs  []*HrzLnRndrSeg
		vrtLnRndrs  []*VrtLnRndrSeg
		hrzBndRndrs []*HrzBndRndrSeg
		vrtBndRndrs []*VrtBndRndrSeg
	}
	StmScp struct {
		Idx uint32
		Arr []*Stm
	}
)

func NewStm() (r *Stm) {
	r = &Stm{}
	r.PltBse = NewPltBse(r)
	r.x = NewTmeAxisX()
	r.y = NewFltAxisY()
	r.sampl = true
	r.mrgn = Mrgn // glbl mrgn
	return r
}
func (x *Stm) X() *TmeAxisX { return x.x }
func (x *Stm) Y() *FltAxisY { return x.y }
func (x *Stm) Stm(pen pen.Pen, stms ...hst.Stm) *Stm {
	for _, stm := range stms {
		x.stmStks = append(x.stmStks, &StmStk{
			stm: stm.Bse(),
			pen: pen,
			plt: &x.TmeFltPltBse,
		})
	}
	return x
}
func (x *Stm) StmBnd(fil clr.Clr, stk pen.Pen, btm, top hst.Stm) *Stm {
	x.stmBnds = append(x.stmBnds, &StmBnd{
		btm: &StmStk{
			stm: btm.Bse(),
			pen: stk,
			plt: &x.TmeFltPltBse,
		},
		top: &StmStk{
			stm: top.Bse(),
			pen: stk,
			plt: &x.TmeFltPltBse,
		},
		filClr: fil,
	})
	return x
}
func (x *Stm) Cnd(pen pen.Pen, cnds ...hst.Cnd) *Stm {
	for _, cnd := range cnds {
		x.cndStks = append(x.cndStks, &CndStk{
			cnd: cnd.Bse(),
			pen: pen,
			plt: &x.TmeFltPltBse,
		})
	}
	return x
}
func (x *Stm) HrzLn(pen pen.Pen, ys ...flt.Flt) *Stm {
	for _, val := range ys {
		x.hrzLns = append(x.hrzLns, &HrzLn{
			val: val,
			pen: pen,
			plt: &x.TmeFltPltBse,
		})
	}
	return x
}
func (x *Stm) VrtLn(pen pen.Pen, xs ...tme.Tme) *Stm {
	for _, val := range xs {
		x.vrtLns = append(x.vrtLns, &VrtLn{
			val: val,
			pen: pen,
			plt: &x.TmeFltPltBse,
		})
	}
	return x
}
func (x *Stm) HrzBnd(fil clr.Clr, stk pen.Pen, btm, top flt.Flt) *Stm {
	x.hrzBnds = append(x.hrzBnds, &HrzBnd{
		btm: &HrzLn{
			val: btm,
			pen: stk,
			plt: &x.TmeFltPltBse,
		},
		top: &HrzLn{
			val: top,
			pen: stk,
			plt: &x.TmeFltPltBse,
		},
		filClr: fil,
	})
	return x
}
func (x *Stm) VrtBnd(fil clr.Clr, stk pen.Pen, lft, rht tme.Tme) *Stm {
	x.vrtBnds = append(x.vrtBnds, &VrtBnd{
		lft: &VrtLn{
			val: lft,
			pen: stk,
			plt: &x.TmeFltPltBse,
		},
		rht: &VrtLn{
			val: rht,
			pen: stk,
			plt: &x.TmeFltPltBse,
		},
		filClr: fil,
	})
	return x
}
func (x *Stm) HrzSclVal(val tme.Tme) *Stm {
	x.X().PxlPerVal = 1.0 / float32(val)
	return x
}
func (x *Stm) VrtSclVal(val flt.Flt) *Stm {
	x.Y().PxlPerVal = 1.0 / float32(val)
	return x
}
func (x *Stm) Sho() Plt             { return x.PltBse.Sho() }
func (x *Stm) Siz(w, h unt.Unt) Plt { return x.PltBse.Siz(w, h) }
func (x *Stm) Scl(v flt.Flt) Plt    { return x.PltBse.Scl(v) }
func (x *Stm) HrzScl(v flt.Flt) Plt { return x.PltBse.HrzScl(v) }
func (x *Stm) VrtScl(v flt.Flt) Plt { return x.PltBse.VrtScl(v) }
