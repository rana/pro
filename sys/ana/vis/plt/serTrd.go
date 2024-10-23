package plt

// import (
// 	"sys/ana/hst"
// 	"sys/ana"
// 	"sys/ana/vis"
// 	"sys/bsc/bnd"
// 	"sys/bsc/flts"
// 	"sys/bsc/tmes"
// )

// type (
// 	SerTrd struct {
// 		PltBse
// 		// Sers    []*SerTrd
// 		TrdsStm *hst.TrdsStm
// 		X       *TmeAxisX
// 		Y       *FltAxisY
// 		inr     vis.Rct
// 	}
// 	SerTrdSeg struct {
// 		bnd.Bnd
// 		*vis.Vis
// 		plt *SerTrd
// 	}
// 	SerTrd struct {
// 		Trd *prv.Trd
// 		X   *tmes.Tmes
// 		Y   *flts.Flts
// 	}
// )

// func NewSerTrd(trdsStm *hst.TrdsStm) (r *SerTrd) {
// 	r = &SerTrd{}
// 	r.TrdsStm = trdsStm
// 	r.X = &TmeAxisX{}
// 	r.Y = NewFltAxisY()
// 	return r
// }
// func (x *SerTrd) Measure() vis.Size {
// 	//x.X.Len = x.crt.inr.Width()
// 	//x.Y.Len = x.crt.inr.Height()
// 	if x.TrdsStm != nil && len(*x.TrdsStm.TrdStmSegs) != 0 {
// 		trdStmSegs := *x.TrdsStm.TrdStmSegs
// 		// x-axis rng
// 		x.X.Rng = trdStmSegs[0].Tmes.Lst() - trdStmSegs[0].Tmes.Fst() // shared max range
// 		for n := 1; n < len(trdStmSegs); n++ {
// 			if trdStmSegs[n].Tmes.Lst()-trdStmSegs[n].Tmes.Fst() > x.X.Rng {
// 				x.X.Rng = trdStmSegs[n].Tmes.Lst() - trdStmSegs[n].Tmes.Fst()
// 			}
// 		}
// 		// y-axis rng
// 		x.Y.Min = trdStmSegs[0].Vals.Min()
// 		x.Y.Max = trdStmSegs[0].Vals.Max()
// 		for n := 1; n < len(trdStmSegs); n++ {
// 			if trdStmSegs[n].Vals.Min() < x.Y.Min {
// 				x.Y.Min = trdStmSegs[n].Vals.Min()
// 			}
// 			if trdStmSegs[n].Vals.Max() > x.Y.Max {
// 				x.Y.Max = trdStmSegs[n].Vals.Max()
// 			}
// 		}
// 		x.Y.Rng = x.Y.Max.Sub(x.Y.Min)

// 		// sys.Log("x.Y.MeasureTxt(x.Y.Min.String())", x.Y.MeasureTxt(x.Y.Min.String()))
// 		// y-axis txt
// 		x.Y.MeasureInrvls(4)
// 		x.Y.Width = max(x.Y.MeasureTxt(x.Y.Min.String()), x.Y.MeasureTxt(x.Y.Max.String()))
// 		x.X.Len -= x.Y.Width + x.Y.Padding + x.Y.Margin

// 		// x.inr.Lft = x.crt.inr.Lft
// 		// x.inr.Rht = x.inr.Lft + x.X.Len + 1 // +1 for rht vrt ln
// 		// x.inr.Top = x.crt.inr.Top
// 		// x.inr.Btm = x.inr.Top + x.Y.Len

// 		x.X.PxPerVal = x.X.Len / float32(x.X.Rng)
// 		x.Y.PxPerVal = x.Y.Len / float32(x.Y.Rng)
// 	}
// 	return vis.Size{Width: x.X.Len, Height: x.Y.Len}
// }

// func (x *SerTrdSeg) Act() { // Seg.Draw
// 	// sys.Log("SerTrdSeg.Act", "Bnd", x.Bnd)
// 	x.SetRGBA(1, 1, 1, 1)

// 	drawCnt := (x.Lim - x.Idx) / 10 // TODO: PLACE DIVISOR (10) IN CFG
// 	if drawCnt <= 0 {
// 		drawCnt = 1
// 	}
// 	// sys.Log("SerTrdSeg.Act", "x.Lim - x.Idx", x.Lim-x.Idx, "drawCnt", drawCnt)
// 	for n := x.Idx; n < x.Lim; n++ {
// 		// ser := x.plt.Sers[n]
// 		trdStm := (*x.plt.TrdsStm.TrdStmSegs)[n]
// 		if trdStm.Tmes.Cnt() < 2 {
// 			continue
// 		}
// 		// x.DrawTxt("abc 100,100", 100, 100)
// 		// a := x.plt.crt.inr
// 		// a.Lft = 0
// 		// a.Rht = x.plt.Width()
// 		// a.Top = x.plt.Height()
// 		// a.Btm = 0
// 		// x.StkRct(a)
// 		// x.StkPthOpn(0, 100, 0, 100)
// 		// x.StkPthTo(150, 200)
// 		// x.StkPthTo(250, 260)
// 		// x.StkPthTo(400, 400)
// 		// x.StkPthCls()

// 		// offX := x.plt.X.Len - (float32(trdStm.Tmes.Lst()-trdStm.Tmes.Fst()) * x.plt.X.PxPerVal) // right align
// 		// x.StkPthOpn(
// 		// 	offX,
// 		// 	offX+float32(trdStm.Tmes.At(1)-trdStm.Tmes.Fst())*x.plt.X.PxPerVal,
// 		// 	x.plt.Y.Px(trdStm.Vals.At(0)),
// 		// 	x.plt.Y.Px(trdStm.Vals.At(1)))
// 		// for n := unt.Unt(2); n < trdStm.Tmes.Cnt(); n++ {
// 		// 	x.StkPthTo(offX+float32(trdStm.Tmes.At(n)-trdStm.Tmes.Fst())*x.plt.X.PxPerVal, x.plt.Y.Px(trdStm.Vals.At(n)))
// 		// }
// 		// x.StkPthCls()
// 		// if n%drawCnt == 0 { // improves additive layering and remove artifacts with high series count
// 		// 	x.Draw()
// 		// }
// 	}
// 	// sys.Log(">># PthLen", x.PthLen())
// 	x.Draw() // draw paths to segement layer
// }
// func (x *SerTrd) Draw() {
// 	// // sys.Log("SerTrd.Draw: START")
// 	// // start := time.Now()
// 	// if len(*x.TrdsStm.TrdStmSegs) != 0 { // draw series
// 	// 	x.crt.vis.SetRGBA(1, 1, 1, 1)
// 	// 	segBnds := bnds.Segs(unt.Unt(len(*x.TrdsStm.TrdStmSegs)))
// 	// 	acts := make([]sys.Act, segBnds.Cnt())
// 	// 	for n, segBnd := range *segBnds {
// 	// 		acts[n] = &SerTrdSeg{
// 	// 			Bnd: segBnd,
// 	// 			plt: x,
// 	// 			Vis: vis.NewVis(x.crt.inr.Width(), x.crt.inr.Height()),
// 	// 		}
// 	// 	}
// 	// 	sys.Run().Pll(acts...)     // process segments in pll
// 	// 	for _, act := range acts { // merge segment images
// 	// 		// x.crt.vis.DrawVis(x.crt.margin.Left, x.crt.margin.Top, act.(*SerTrdSeg).Vis)
// 	// 		x.crt.vis.DrawVis(0, 0, act.(*SerTrdSeg).Vis)
// 	// 	}
// 	// 	// // draw border
// 	// 	// x.crt.vis.PxlRct(x.crt.inr, colornames.Gray)
// 	// 	// draw y-axis lft
// 	// 	x.crt.vis.PxlRct(x.inr, colornames.Gray)
// 	// 	// draw y-axis txt
// 	// 	x.crt.vis.Txt(x.inr.Rht+x.Y.Padding, x.inr.Top, 1, 1, x.Y.Font, x.Y.Max.Trnc(5).String())
// 	// 	for _, inrvl := range x.Y.Inrvls {
// 	// 		x.crt.vis.Txt(x.inr.Rht+x.Y.Padding, x.Y.Px(inrvl), 1, .5, x.Y.Font, inrvl.Trnc(5).String())
// 	// 	}
// 	// 	x.crt.vis.Txt(x.inr.Rht+x.Y.Padding, x.inr.Btm-4, 1, 0, x.Y.Font, x.Y.Min.Trnc(5).String())
// 	// 	// sys.Log("x.inr", x.inr)
// 	// }
// 	// // sys.Log("SerTrd.Draw: END: ellapsed:", time.Now().Sub(start), "series count:", len(*x.TrdsStm.TrdStmSegs))
// }
// func (x *SerTrd) Bse() *PltBse    { return &x.PltBse }
