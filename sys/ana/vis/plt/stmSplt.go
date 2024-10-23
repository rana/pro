package plt

// import (
// 	"image"
// 	"sys"
// 	"sys/ana"
// 	"sys/ana/vis"
// 	"sys/bsc/bnd"
// 	"sys/bsc/bnds"
// 	"sys/bsc/flt"
// 	"sys/bsc/flts"
// 	"sys/bsc/str"
// 	"sys/bsc/unt"
// )

// type (
// 	StmSpltSctrRndrSeg struct {
// 		bnd.Bnd
// 		Plt  *StmSplt
// 		Sctr *Sctr
// 		Pnts []SctrRndrPnt
// 	}
// )

// func (x *StmSplt) Measure() vis.Siz { // Plt interface
// 	if len(x.sctrs) != 0 {
// 		x.Y.Min, x.Y.Max = flt.Max, flt.Min
// 		for _, cndStmSctr := range x.sctrs {
// 			min, max := cndStmSctr.Y.MinMax()
// 			if !x.Outlier && cndStmSctr.Y.Cnt() > 2 {
// 				// var b0, b1, b2, b3, b4 strings.Builder
// 				rngVals := flts.Make(cndStmSctr.Y.Cnt())
// 				for n, y := range *cndStmSctr.Y {
// 					(*rngVals)[n] = y - min
// 					// b0.WriteString(fmt.Sprintf(" %v", y))
// 					// b3.WriteString(fmt.Sprintf(" %v", (*rngVals)[n]))
// 				}
// 				cpy := rngVals.Cpy().SrtAsc()
// 				rngValsMdn := cpy.Mdl()
// 				// sys.Log("rngValsMdn != 0", rngValsMdn != 0)
// 				if rngValsMdn == 0 {
// 					for n := cpy.MdlIdx() + 1; rngValsMdn == 0 && n < cpy.Cnt(); n++ {
// 						rngValsMdn = cpy.At(n)
// 					}
// 					// sys.Log("FIXED rngValsMdn != 0", rngValsMdn != 0)
// 				}

// 				if rngValsMdn != 0 {
// 					for n := len(*rngVals) - 1; n > -1; n-- { // dsc in case of multiple deletions
// 						rngVal := (*rngVals)[n]
// 						// mu.Lock()
// 						// pcts.Push(rngVal / rngValsMdn)
// 						// mu.Unlock()
// 						// sys.Log("plt.StmSplt pct:%v", rngVal/rngValsMdn)
// 						// sys.Logf("%p rngVal:%v rngValsMdn:%v rngVal/rngValsMdn:%v OutlierLim:%v rngVal/rngValsMdn > OutlierLim:%v",
// 						// 	cndStmSctr, rngVal, rngValsMdn, rngVal/rngValsMdn, OutlierLim, rngVal/rngValsMdn > OutlierLim)
// 						if rngVal/rngValsMdn > OutlierLim { // outlier case observed: 7400%; lim is 1000%
// 							// sys.Logf("%p BEFORE: cndStmSctr.Y cnt:%v n:%v maxIdx:%v Y[n]:%v",
// 							// 	cndStmSctr, cndStmSctr.Y.Cnt(), n, cndStmSctr.Y.SrchIdxEql(max), (*cndStmSctr.Y)[n])
// 							// sys.Logf("%p BEFORE: cndStmSctr.Y %v", cndStmSctr, cndStmSctr.Y)
// 							del := cndStmSctr.Y.Del(unt.Unt(n))
// 							// sys.Logf("%p  AFTER: cndStmSctr.Y %v", cndStmSctr, cndStmSctr.Y)
// 							// sys.Logf("%p  AFTER: cndStmSctr.Y cnt:%v", cndStmSctr, cndStmSctr.Y.Cnt())

// 							sys.Logf("plt.StmSplt %p HIDING OUTLIER POINT %v %v%%", x, del, rngVal.Div(rngValsMdn).Trnc(0))
// 						}
// 						// (*mdnPros)[n] = rngVal / rngValsMdn
// 						// b4.WriteString(fmt.Sprintf(" %v", (*mdnPros)[n]))
// 					}
// 					// sys.Log("pcts", pcts.SrtDsc())

// 					// sys.Logf("%p BEFORE: min:%v max:%v cnt:%v", cndStmSctr, min, max, cndStmSctr.Y.Cnt())
// 					min, max = cndStmSctr.Y.MinMax() // recalculate min/max
// 					// sys.Logf("%p  AFTER: min:%v max:%v cnt:%v Y:%v", cndStmSctr, min, max, cndStmSctr.Y.Cnt(), cndStmSctr.Y)
// 					// sys.Log("      y: ", b0.String())
// 					// sys.Log("rngVals: ", b3.String())
// 					// sys.Log("mdnPros: ", b4.String())
// 				}
// 			}
// 			x.Y.Min = x.Y.Min.Min(min)
// 			x.Y.Max = x.Y.Max.Max(max)
// 		}
// 		x.Y.Rng = x.Y.Max - x.Y.Min
// 		if x.Y.Rng != 0 {
// 			// y-axis txt
// 			x.Y.MeasureInrvls(ana.Cfg.Plt.InrvlCnt)
// 			x.Y.Rht.Width = uint32(AxisPad) + uint32(InrvlTxtLen) // constant txt len to avoid variable radiusths where cnd lns don't align
// 			// y-axis
// 			if x.Y.PxlPerVal != 0 { // define height on user-defined scl
// 				valPerPxl := (1.0 / x.Y.PxlPerVal)
// 				x.Y.Height = uint32(float32(x.Y.Rng) / valPerPxl)
// 				x.siz.Height = x.Y.Height + x.mrgn.Height() + uint32(BrdrLen)*2
// 			} else { // measure height based on plt-defined segment vals
// 				x.Y.Height = x.siz.Height - x.mrgn.Height() - uint32(BrdrLen)*2 - uint32(BarPad)*2
// 				x.Y.PxlPerVal = float32(x.Y.Height) / float32(x.Y.Rng)
// 			}
// 			// x-axis
// 			for n, cndStmSctr := range x.sctrs {
// 				x.XWidth += cndStmSctr.radius * 2
// 				if n == 0 {
// 					cndStmSctr.cntrX = cndStmSctr.radius // Plt.off.Y starts with pad
// 				} else {
// 					cndStmSctr.cntrX = x.sctrs[n-1].cntrX + x.sctrs[n-1].radius + uint32(BarPad) + cndStmSctr.radius
// 					x.XWidth += uint32(BarPad)
// 				}
// 			}
// 			x.siz.Width = x.mrgn.Width() + uint32(BrdrLen)*2 + x.Y.Rht.Width + x.XWidth + uint32(BarPad)*2
// 		} else {
// 			x.siz.Width = x.mrgn.Width() + uint32(BrdrLen)*2 + x.Y.Rht.Width + uint32(Len) + uint32(BarPad)*2
// 		}

// 	}
// 	return x.siz
// }

// func (x *StmSpltSctrRndrSeg) Act() {
// 	var pntCnt unt.Unt
// 	// sys.Logf("SctrRndrSeg x.Idx:%v x.Lim:%v len(x.Plt.sctrs):%v", x.Idx, x.Lim, len(x.Plt.sctrs))
// 	for m := x.Idx; m < x.Lim; m++ { // count pnts within current seg
// 		pntCnt += x.Plt.sctrs[m].Y.Cnt()
// 	}
// 	x.Pnts = make([]SctrRndrPnt, pntCnt) // ensure alloc for draw
// 	if pntCnt != 0 {
// 		pntActs := make([]sys.Act, pntCnt)
// 		var pntIdx int
// 		for m := x.Idx; m < x.Lim; m++ { // generate a rendr for each point within seg
// 			cndStmSctr := x.Plt.sctrs[m]
// 			if cndStmSctr.Y != nil {
// 				// sys.Logf("SctrRndrSeg cndStmSctr.Y:%v", cndStmSctr.Y.Cnt())
// 				for _, val := range *cndStmSctr.Y {
// 					x.Pnts[pntIdx].ras.Reset(int(x.Plt.siz.Width), int(x.Plt.siz.Height))
// 					x.Pnts[pntIdx].cntr.X = x.Plt.off.X + float32(x.Plt.sctrs[m].cntrX)
// 					x.Pnts[pntIdx].cntr.Y = x.Plt.off.Y + x.Plt.Y.Pxl(val)
// 					x.Pnts[pntIdx].radius = float32(x.Plt.sctrs[m].radius)
// 					x.Pnts[pntIdx].clr = x.Plt.sctrs[m].clr
// 					pntActs[pntIdx] = &x.Pnts[pntIdx]
// 					// sys.Logf("SctrRndrSeg val:%v, cntr.Y:%v", val, x.Plt.Y.Pxl(val))
// 					pntIdx++
// 				}
// 			}
// 		}
// 		// sys.Logf("SctrRndrSeg pntCnt:%v pntIdx:%v", pntCnt, pntIdx)
// 		sys.Run().Pll(pntActs...)
// 	}
// }
// func (x *StmSplt) Rndr() {
// 	if x.Y.Rng <= 0 { // no data
// 		// sys.Logf("%p Rng.Rndr INVALID (Y.Rng:%v Y.Min:%v Y.Max:%v) \n", x, x.Y.Rng, x.Y.Min, x.Y.Max)
// 		return
// 	}
// 	// sys.Log("Rng.Rndr", "x.pos", x.pos, "x.siz", x.siz)
// 	// brdr (for pxl pnts written directly to img)
// 	// x.pos known only after Measure call
// 	x.brdr.Lft = x.pos.X + x.mrgn.Lft + uint32(BrdrLen)
// 	x.brdr.Rht = x.brdr.Lft + x.XWidth + uint32(BrdrLen) + uint32(BarPad)*2
// 	x.brdr.Top = x.pos.Y + x.mrgn.Top + uint32(BrdrLen)
// 	x.brdr.Btm = x.brdr.Top + x.Y.Height + uint32(BrdrLen) + uint32(BarPad)*2
// 	// off (for stk/fil pnts rasterized within brdr)
// 	x.off.X = float32(x.brdr.Lft - x.pos.X + uint32(BrdrLen) + uint32(BarPad))
// 	x.off.Y = float32(x.brdr.Top - x.pos.Y + uint32(BrdrLen) + uint32(BarPad))
// 	// draw sqrs for each pnt; single rasterizer for all shapes
// 	segBnds, acts := bnds.Segs(unt.Unt(len(x.sctrs)))
// 	x.sctrRndrSegs = make([]*StmSpltSctrRndrSeg, len(acts))
// 	// sys.Log("len(x.sctrRndrSegs)", len(x.sctrRndrSegs))
// 	for n, segBnd := range *segBnds {
// 		x.sctrRndrSegs[n] = &StmSpltSctrRndrSeg{
// 			Bnd: segBnd,
// 			Plt: x,
// 		}
// 		acts[n] = x.sctrRndrSegs[n]
// 	}
// 	sys.Run().Pll(acts...)
// }
// func (x *StmSplt) Draw(img *image.RGBA) {
// 	if x.Y.Rng <= 0 {
// 		x1 := float32(x.pos.X) + float32(x.siz.Width)*.5
// 		y1 := float32(x.pos.Y) + float32(x.siz.Height)*.5
// 		MsgFnt.Draw(x1, y1, .5, .5, MsgClr, img, TxtNoData)
// 		return
// 	}
// 	rct := vis.Rect(x.pos, x.siz) // draw order is z-order
// 	// // shp
// 	// if x.shp.ras.Size() != image.ZP {
// 	// 	x.shp.ras.Draw(img, rct, x.shp.clr.Uniform(), image.Point{})
// 	// }
// 	for _, cndStmSctrRndrSeg := range x.sctrRndrSegs {
// 		for _, pnt := range cndStmSctrRndrSeg.Pnts {
// 			if pnt.ras.Size() != image.ZP {
// 				pnt.ras.Draw(img, rct, pnt.clr.Uniform(), image.Point{})
// 			}
// 		}
// 	}
// 	// brdr
// 	vis.PxlRct(x.brdr, BrdrClr.Color(), img)
// 	// y-axis txt
// 	x1 := float32(x.brdr.Rht + uint32(BrdrLen+AxisPad))
// 	// InrvlFnt.Draw(x1, float32(x.brdr.Top+uint32(BarPad)+uint32(BrdrLen)), 1, 1, InrvlTxtClrY, img, x.Y.Max.Trnc(AxisTrnc).String())
// 	for _, inrvl := range x.Y.Inrvls {
// 		InrvlFnt.Draw(x1, float32(x.brdr.Top+uint32(BarPad))+x.Y.Pxl(inrvl), 1, .5, InrvlTxtClrY, img, inrvl.Trnc(AxisTrnc).String())
// 	}
// 	// InrvlFnt.Draw(x1, float32(x.brdr.Btm-uint32(BarPad)-uint32(BrdrLen)-4), 1, 0, InrvlTxtClrY, img, x.Y.Min.Trnc(AxisTrnc).String())

// 	// plt title
// 	if x.Title != str.Empty {
// 		x1 := float32(x.brdr.Lft+uint32(BrdrLen)) + float32(x.XWidth)*.5
// 		y1 := float32(x.brdr.Btm-uint32(BrdrLen)-uint32(AxisPad)) - TitleFnt.Height()*.5
// 		TitleFnt.Draw(x1, y1, 0.5, 0.5, TitleClr, img, x.Title.Unquo())
// 	}
// }
