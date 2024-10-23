package plt

// import (
// 	"image"
// 	"sys"
// 	"sys/ana"
// 	"sys/ana/vis"
// 	"sys/bsc/flt"
// 	"sys/bsc/str"
// 	"sys/bsc/tme"
// )

// func (x *Stgy) Measure() vis.Siz { // Plt interface
// 	trds := x.stgy.Port().Bse().Trds // calc trds
// 	if trds.Cnt() == 0 {
// 		return x.siz
// 	}
// 	for _, trd := range *trds {
// 		if trd.PnlPct > 0 {
// 			if x.Pos {
// 				x.stmStkTrds = append(x.stmStkTrds, &StmStkTrd{
// 					trd: trd,
// 					stm: x.stm,
// 					pen: PrfPen,
// 					plt: &x.TmeFltPltBse,
// 				})
// 			}
// 		} else {
// 			if x.Neg {
// 				x.stmStkTrds = append(x.stmStkTrds, &StmStkTrd{
// 					trd: trd,
// 					stm: x.stm,
// 					pen: LosPen,
// 					plt: &x.TmeFltPltBse,
// 				})
// 			}
// 		}

// 	}

// 	acts := make([]sys.Act, 0, len(x.stmStkTrds))
// 	// stmStkTrds
// 	var stmStkActs []*StmStkTrdMeasureSeg
// 	if len(x.stmStkTrds) != 0 {
// 		stmStkActs = make([]*StmStkTrdMeasureSeg, len(x.stmStkTrds))
// 		for n, stmStk := range x.stmStkTrds {
// 			stmStkActs[n] = &StmStkTrdMeasureSeg{StmStkTrd: stmStk}
// 			acts = append(acts, stmStkActs[n])
// 		}
// 	}
// 	sys.Run().Pll(acts...) // measure all elements in pll
// 	if len(stmStkActs) != 0 {
// 		// fmt.Println("x.x.Min", x.x.Min)
// 		// fmt.Println("tme.Max", tme.Max)
// 		// x.x.Min = tme.Max
// 		// x.x.Max = tme.Min
// 		x.y.Min = flt.Max
// 		x.y.Max = flt.Min
// 		x.x.Rng = 0
// 		x.y.Rng = 0
// 		maxBtmLenY := flt.Min
// 		maxTopLenY := flt.Min
// 		for _, seg := range stmStkActs {
// 			x.x.Rng = x.x.Rng.Max(seg.RngX)
// 			if seg.Y != nil && len(*seg.Y) != 0 {
// 				if seg.MinY < x.y.Min {
// 					x.y.Min = seg.MinY
// 				}
// 				if seg.MaxY > x.y.Max {
// 					x.y.Max = seg.MaxY
// 				}
// 				seg.BtmLenY = seg.Y.Fst() - seg.MinY
// 				seg.TopLenY = seg.MaxY - seg.Y.Fst()
// 				maxBtmLenY = maxBtmLenY.Max(seg.Y.Fst() - seg.MinY)
// 				maxTopLenY = maxTopLenY.Max(seg.MaxY - seg.Y.Fst())
// 			}
// 		}
// 		for _, seg := range stmStkActs {
// 			seg.OffY = maxBtmLenY - seg.BtmLenY
// 		}
// 		x.x.Min = tme.Zero
// 		x.x.Max = x.x.Rng
// 		x.y.Rng = maxBtmLenY + maxTopLenY
// 		x.y.Min = -maxBtmLenY
// 		x.y.Max = x.y.Rng - maxBtmLenY

// 		// x.y.Rng *= 2
// 		// x.y.Rng = x.y.Rng.Max(seg.MaxY - seg.MinY) // find max eng

// 		// if x.x.Min == tme.Max || x.x.Max == tme.Min {
// 		// 	x.x.Rng = 0
// 		// } else {
// 		// 	x.x.Rng = x.x.Max - x.x.Min
// 		// }
// 		// if x.y.Min == flt.Max || x.y.Max == flt.Min {
// 		// 	x.y.Rng = 0
// 		// } else {
// 		// 	if x.y.EqiDst != flt.Max && x.y.EqiDst >= x.y.Min && x.y.EqiDst <= x.y.Max {
// 		// 		// make min/max equi-distant from the specified point
// 		// 		if x.y.EqiDst.Sub(x.y.Min) < x.y.Max.Sub(x.y.EqiDst) {
// 		// 			x.y.Min = x.y.EqiDst - x.y.Max.Sub(x.y.EqiDst)
// 		// 		} else {
// 		// 			x.y.Max = x.y.EqiDst + x.y.EqiDst.Sub(x.y.Min)
// 		// 		}
// 		// 	}
// 		// 	x.y.Rng = x.y.Max - x.y.Min
// 		// }
// 		if x.x.Rng != 0 && x.y.Rng != 0 {
// 			// // x-axis txt
// 			if x.x.vis {
// 				x.x.MeasureInrvls(ana.Cfg.Plt.InrvlCnt)
// 				x.x.Btm.Height = uint32(AxisPad) + InrvlFnt.HeightUint32()
// 			}
// 			// y-axis txt
// 			// minY, maxY := x.y.Min, x.y.Max
// 			// x.y.Min, x.y.Max = flt.Zero, x.y.Rng // set for inrvl text
// 			x.y.MeasureInrvls(ana.Cfg.Plt.InrvlCnt)
// 			// x.y.Min, x.y.Max = minY, maxY
// 			x.y.Rht.Width = uint32(AxisPad) + uint32(InrvlTxtLen) // constant txt len to avoid variable widths where cnd lns don't align
// 			// y-axis
// 			if x.y.PxlPerVal != 0 { // define height on user-defined scl
// 				valPerPxl := (1.0 / x.y.PxlPerVal)
// 				x.y.Height = uint32(float32(x.y.Rng) / valPerPxl)
// 				x.siz.Height = x.y.Height + x.mrgn.Height() + uint32(BrdrLen)*2 + x.x.Btm.Height
// 			} else { // measure height based on plt-defined segment vals
// 				x.y.Height = x.siz.Height - x.mrgn.Height() - uint32(BrdrLen)*2 - x.x.Btm.Height
// 				x.y.PxlPerVal = float32(x.y.Height) / float32(x.y.Rng)
// 			}
// 			// x-axis
// 			if x.x.PxlPerVal != 0 { // define width on user-defined scl
// 				valPerPxl := (1.0 / x.x.PxlPerVal)
// 				x.x.Width = uint32(float32(x.x.Rng) / valPerPxl)
// 				x.siz.Width = x.x.Width + x.mrgn.Width() + uint32(BrdrLen)*2 + x.y.Rht.Width
// 			} else { // measure width based on plt-defined segment tmes
// 				x.x.Width = x.siz.Width - x.mrgn.Width() - uint32(BrdrLen)*2 - x.y.Rht.Width
// 				x.x.PxlPerVal = float32(x.x.Width) / float32(x.x.Rng)
// 			}
// 		}
// 	}
// 	// if x.Title == str.Empty {
// 	// 	x.Title = str.Str(x.x.Rng.String())
// 	// } else {
// 	// 	x.Title = str.Str(fmt.Sprintf("%v (%v)", string(x.Title), x.x.Rng.String()))
// 	// }
// 	// sys.Log("Stm.Measure", "x.pos", x.pos, "x.siz", x.siz)
// 	return x.siz
// }
// func (x *Stgy) Rndr() {
// 	if x.x.Rng <= 0 || x.y.Rng <= 0 {
// 		sys.Logf("%p Stm.Rndr INVALID (X.Rng:%v Y.Rng:%v X.Min:%v X.Max:%v Y.Min:%v Y.Max:%v) \n", x, x.x.Rng.DurString(), x.y.Rng, x.x.Min, x.x.Max, x.y.Min, x.y.Max)
// 		return
// 	}
// 	// sys.Log("Stm.Rndr", "x.pos", x.pos, "x.siz", x.siz)
// 	// brdr (for pxl pnts written directly to img)
// 	// x.pos known only after Measure call
// 	x.brdr.Lft = x.pos.X + x.mrgn.Lft + uint32(BrdrLen)
// 	x.brdr.Rht = x.brdr.Lft + x.x.Width + uint32(BrdrLen)
// 	x.brdr.Top = x.pos.Y + x.mrgn.Top + uint32(BrdrLen)
// 	x.brdr.Btm = x.brdr.Top + x.y.Height + uint32(BrdrLen)
// 	// off (for stk/fil pnts rasterized within brdr)
// 	x.off.X = float32(x.brdr.Lft - x.pos.X + uint32(BrdrLen))
// 	x.off.Y = float32(x.brdr.Top - x.pos.Y + uint32(BrdrLen))
// 	// sys.Log("Stm.Rndr", "x.brdr", x.brdr)
// 	// sys.Log("Stm.Rndr", "x.off", x.off)
// 	// sys.Log("Stm.Rndr", "len(x.stmStkTrds)", len(x.stmStkTrds))
// 	acts := make([]sys.Act, 0, len(x.stmStkTrds))
// 	// stmStkTrds
// 	if len(x.stmStkTrds) != 0 {
// 		x.stmStkTrdRndrs = make([]*StmStkTrdRndrSeg, len(x.stmStkTrds))
// 		for n, stmStkTrd := range x.stmStkTrds {
// 			x.stmStkTrdRndrs[n] = &StmStkTrdRndrSeg{StmStkTrd: stmStkTrd}
// 			acts = append(acts, x.stmStkTrdRndrs[n])
// 		}
// 	}
// 	sys.Run().Pll(acts...) // render all elements in pll
// }
// func (x *Stgy) Draw(img *image.RGBA) {
// 	if x.x.Rng <= 0 || x.y.Rng <= 0 {
// 		x1 := float32(x.pos.X) + float32(x.siz.Width)*.5
// 		y1 := float32(x.pos.Y) + float32(x.siz.Height)*.5
// 		MsgFnt.Draw(x1, y1, .5, .5, MsgClr, img, TxtNoData)
// 		return
// 	}
// 	rct := vis.Rect(x.pos, x.siz) // draw order is z-order
// 	// sys.Log("Stm.Draw", "rct", rct)
// 	// sys.Log("Stm.Draw", "len(x.stmStkTrdRndrs)", len(x.stmStkTrdRndrs))

// 	// stmStkTrds
// 	if len(x.stmStkTrdRndrs) != 0 {
// 		for _, seg := range x.stmStkTrdRndrs {
// 			if seg.ras.Size() == image.ZP {
// 				continue
// 			}
// 			seg.Draw(seg.pen.Clr, rct, img, &seg.ras)
// 		}
// 	}

// 	// brdr
// 	vis.PxlRct(x.brdr, BrdrClr.Color(), img)

// 	if x.x.vis {
// 		y1 := float32(x.brdr.Btm + uint32(BrdrLen+AxisPad))
// 		InrvlFnt.Draw(float32(x.brdr.Lft+uint32(BrdrLen)), y1, 1, 1, InrvlTxtClrX, img, x.x.Min.String())
// 		for _, inrvl := range x.x.Inrvls[1 : len(x.x.Inrvls)-1] {
// 			InrvlFnt.Draw(float32(x.brdr.Lft)+x.x.Pxl(inrvl), y1, .5, 1, InrvlTxtClrX, img, inrvl.String())
// 		}
// 		InrvlFnt.Draw(float32(x.brdr.Rht-uint32(BrdrLen)), y1, 0, 1, InrvlTxtClrX, img, x.x.Max.String())
// 	}
// 	// y-axis txt
// 	x1 := float32(x.brdr.Rht + uint32(BrdrLen+AxisPad))
// 	InrvlFnt.Draw(x1, float32(x.brdr.Top+uint32(BrdrLen)), 1, 1, InrvlTxtClrY, img, x.y.Max.Trnc(AxisTrnc).String())
// 	for _, inrvl := range x.y.Inrvls[1 : len(x.y.Inrvls)-1] {
// 		InrvlFnt.Draw(x1, float32(x.brdr.Top)+x.y.Pxl(inrvl), 1, .5, InrvlTxtClrY, img, inrvl.Trnc(AxisTrnc).String())
// 	}
// 	InrvlFnt.Draw(x1, float32(x.brdr.Btm-uint32(BrdrLen)-4), 1, 0, InrvlTxtClrY, img, x.y.Min.Trnc(AxisTrnc).String())

// 	// plt title
// 	if x.Title != str.Empty {
// 		x1 := float32(x.brdr.Lft+uint32(BrdrLen)) + float32(x.x.Width)*.5 //- TitleFnt.Width(x.Title.Unquo())*.5
// 		y1 := float32(x.brdr.Btm-uint32(BrdrLen)-uint32(AxisPad)) - TitleFnt.Height()*.5
// 		TitleFnt.Draw(x1, y1, 0.5, 0.5, TitleClr, img, x.Title.Unquo())
// 	}
// }
