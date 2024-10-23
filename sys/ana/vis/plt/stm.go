package plt

import (
	"image"
	"sys"
	"sys/ana"
	"sys/ana/vis"
	"sys/bsc/flt"
	"sys/bsc/str"
	"sys/bsc/tme"
	"sys/bsc/tmes"
)

func (x *Stm) Measure() vis.Siz { // Plt interface
	acts := make([]sys.Act, 0, len(x.stmStks)+len(x.stmBnds))
	// stmStks
	var stmStkActs []*StmStkMeasureSeg
	if len(x.stmStks) != 0 {
		stmStkActs = make([]*StmStkMeasureSeg, len(x.stmStks))
		for n, stmStk := range x.stmStks {
			stmStkActs[n] = &StmStkMeasureSeg{StmStk: stmStk}
			acts = append(acts, stmStkActs[n])
		}
	}
	// stmBnds
	var stmBndActs []*StmBndMeasureSeg
	if len(x.stmBnds) != 0 {
		stmBndActs = make([]*StmBndMeasureSeg, len(x.stmBnds))
		for n, stmBnd := range x.stmBnds {
			stmBndActs[n] = &StmBndMeasureSeg{StmBnd: stmBnd}
			acts = append(acts, stmBndActs[n])
		}
	}
	sys.Run().Pll(acts...) // measure all elements in pll
	if len(stmStkActs) != 0 || len(stmBndActs) != 0 {
		// fmt.Println("x.x.Min", x.x.Min)
		// fmt.Println("tme.Max", tme.Max)
		x.x.Min = tme.Max
		x.x.Max = tme.Min
		calcMinY := x.y.Min == flt.Max
		if calcMinY { // min may be set by user
			x.y.Min = flt.Max
		}
		calcMaxY := x.y.Max == flt.Min
		if calcMaxY {
			x.y.Max = flt.Min
		}
		for _, stmBndAct := range stmBndActs { // place bnd stmStks in slice for min/max calculation
			stmStkActs = append(stmStkActs, stmBndAct.btmSeg, stmBndAct.topSeg)
		}
		for _, seg := range stmStkActs {
			if seg.X == nil || len(*seg.X) < 2 {
				// var cnt int
				// if seg.X != nil {
				// 	cnt = len(*seg.X)
				// }
				// sys.Logf("Stm/StmStk: TOO FEW PNTS TO CALCULATE MIN/MAX (cnt:%v) %v", cnt, seg.stm)
				continue
			}
			var tmes *tmes.Tmes
			if x.x.PxlPerVal != 0 { // user-defined scl
				tmes = seg.stm.Tmes // use all stm tmes
			} else {
				tmes = seg.X // use plt-defined segment tmes
			}
			if tmes.Fst() < x.x.Min {
				x.x.Min = tmes.Fst()
			}
			if tmes.Lst() > x.x.Max {
				x.x.Max = tmes.Lst()
			}
			// var vals *flts.Flts
			// if x.y.PxlPerVal != 0 { // user-defined scl
			// 	vals = seg.stm.Vals // use all stm vals
			// } else {
			// 	vals = seg.Y // use plt-defined segment vals
			// }
			if calcMinY && seg.MinY < x.y.Min {
				x.y.Min = seg.MinY
			}
			if calcMaxY && seg.MaxY > x.y.Max {
				x.y.Max = seg.MaxY
			}
		}
		if x.x.Min == tme.Max || x.x.Max == tme.Min {
			x.x.Rng = 0
		} else {
			x.x.Rng = x.x.Max - x.x.Min
		}
		if x.y.Min == flt.Max || x.y.Max == flt.Min {
			x.y.Rng = 0
		} else {
			if x.y.EqiDst != flt.Max && x.y.EqiDst >= x.y.Min && x.y.EqiDst <= x.y.Max {
				// make min/max equi-distant from the specified point
				if x.y.EqiDst.Sub(x.y.Min) < x.y.Max.Sub(x.y.EqiDst) {
					x.y.Min = x.y.EqiDst - x.y.Max.Sub(x.y.EqiDst)
				} else {
					x.y.Max = x.y.EqiDst + x.y.EqiDst.Sub(x.y.Min)
				}
			}
			x.y.Rng = x.y.Max - x.y.Min
		}
		if x.x.Rng != 0 && x.y.Rng != 0 {
			// x-axis txt
			if x.x.vis {
				x.x.MeasureInrvls(ana.Cfg.Plt.InrvlCnt)
				x.x.Btm.Height = uint32(AxisPad) + InrvlFnt.HeightUint32()
			}
			// y-axis txt
			x.y.MeasureInrvls(ana.Cfg.Plt.InrvlCnt)
			x.y.Rht.Width = uint32(AxisPad) + uint32(InrvlTxtLen) // constant txt len to avoid variable widths where cnd lns don't align
			// y-axis
			if x.y.PxlPerVal != 0 { // define height on user-defined scl
				valPerPxl := (1.0 / x.y.PxlPerVal)
				x.y.Height = uint32(float32(x.y.Rng) / valPerPxl)
				x.siz.Height = x.y.Height + x.mrgn.Height() + uint32(BrdrLen)*2 + x.x.Btm.Height
			} else { // measure height based on plt-defined segment vals
				x.y.Height = x.siz.Height - x.mrgn.Height() - uint32(BrdrLen)*2 - x.x.Btm.Height
				x.y.PxlPerVal = float32(x.y.Height) / float32(x.y.Rng)
			}
			// x-axis
			if x.x.PxlPerVal != 0 { // define width on user-defined scl
				valPerPxl := (1.0 / x.x.PxlPerVal)
				x.x.Width = uint32(float32(x.x.Rng) / valPerPxl)
				x.siz.Width = x.x.Width + x.mrgn.Width() + uint32(BrdrLen)*2 + x.y.Rht.Width
			} else { // measure width based on plt-defined segment tmes
				x.x.Width = x.siz.Width - x.mrgn.Width() - uint32(BrdrLen)*2 - x.y.Rht.Width
				x.x.PxlPerVal = float32(x.x.Width) / float32(x.x.Rng)
			}
		}
	}
	// sys.Log("Stm.Measure", "x.pos", x.pos, "x.siz", x.siz)
	return x.siz
}
func (x *Stm) Rndr() {
	if x.x.Rng <= 0 || x.y.Rng <= 0 {
		sys.Logf("%p Stm.Rndr INVALID (X.Rng:%v Y.Rng:%v X.Min:%v X.Max:%v Y.Min:%v Y.Max:%v) \n", x, x.x.Rng.DurString(), x.y.Rng, x.x.Min, x.x.Max, x.y.Min, x.y.Max)
		return
	}
	// sys.Log("Stm.Rndr", "x.pos", x.pos, "x.siz", x.siz)
	// brdr (for pxl pnts written directly to img)
	// x.pos known only after Measure call
	x.brdr.Lft = x.pos.X + x.mrgn.Lft + uint32(BrdrLen)
	x.brdr.Rht = x.brdr.Lft + x.x.Width + uint32(BrdrLen)
	x.brdr.Top = x.pos.Y + x.mrgn.Top + uint32(BrdrLen)
	x.brdr.Btm = x.brdr.Top + x.y.Height + uint32(BrdrLen)
	// off (for stk/fil pnts rasterized within brdr)
	x.off.X = float32(x.brdr.Lft - x.pos.X + uint32(BrdrLen))
	x.off.Y = float32(x.brdr.Top - x.pos.Y + uint32(BrdrLen))
	// sys.Log("Stm.Rndr", "x.brdr", x.brdr)
	// sys.Log("Stm.Rndr", "x.off", x.off)
	// sys.Log("Stm.Rndr", "len(x.stmStks)", len(x.stmStks))
	// sys.Log("Stm.Rndr", "len(x.stmBnds)", len(x.stmBnds))
	// sys.Log("Stm.Rndr", "len(x.cndStks)", len(x.cndStks))
	// sys.Log("Stm.Rndr", "len(x.hrzLns)", len(x.hrzLns))
	// sys.Log("Stm.Rndr", "len(x.vrtLns)", len(x.vrtLns))
	// sys.Log("Stm.Rndr", "len(x.hrzBnds)", len(x.hrzBnds))
	// sys.Log("Stm.Rndr", "len(x.vrtBnds)", len(x.vrtBnds))

	acts := make([]sys.Act, 0,
		len(x.stmStks)+
			len(x.stmBnds)+
			len(x.cndStks)+
			len(x.hrzLns)+
			len(x.vrtLns)+
			len(x.hrzBnds)+
			len(x.vrtBnds))
	// stmStks
	if len(x.stmStks) != 0 {
		x.stmStkRndrs = make([]*StmStkRndrSeg, len(x.stmStks))
		for n, stmStk := range x.stmStks {
			x.stmStkRndrs[n] = &StmStkRndrSeg{StmStk: stmStk}
			acts = append(acts, x.stmStkRndrs[n])
		}
	}
	// stmBnds
	if len(x.stmBnds) != 0 {
		x.stmBndRndrs = make([]*StmBndRndrSeg, len(x.stmBnds))
		for n, stmBnd := range x.stmBnds {
			x.stmBndRndrs[n] = &StmBndRndrSeg{StmBnd: stmBnd}
			acts = append(acts, x.stmBndRndrs[n])
		}
	}
	// cndStks
	if len(x.cndStks) != 0 {
		x.cndStkRndrs = make([]*CndStkRndrSeg, len(x.cndStks))
		for n, cndStk := range x.cndStks {
			x.cndStkRndrs[n] = &CndStkRndrSeg{CndStk: cndStk}
			acts = append(acts, x.cndStkRndrs[n])
		}
	}
	// hrzLns
	if len(x.hrzLns) != 0 {
		x.hrzLnRndrs = make([]*HrzLnRndrSeg, len(x.hrzLns))
		for n, hrzLn := range x.hrzLns {
			x.hrzLnRndrs[n] = &HrzLnRndrSeg{HrzLn: hrzLn}
			acts = append(acts, x.hrzLnRndrs[n])
		}
	}
	// vrtLns
	if len(x.vrtLns) != 0 {
		x.vrtLnRndrs = make([]*VrtLnRndrSeg, len(x.vrtLns))
		for n, vrtLn := range x.vrtLns {
			x.vrtLnRndrs[n] = &VrtLnRndrSeg{VrtLn: vrtLn}
			acts = append(acts, x.vrtLnRndrs[n])
		}
	}
	// hrzBnds
	if len(x.hrzBnds) != 0 {
		x.hrzBndRndrs = make([]*HrzBndRndrSeg, len(x.hrzBnds))
		for n, hrzBnd := range x.hrzBnds {
			x.hrzBndRndrs[n] = &HrzBndRndrSeg{HrzBnd: hrzBnd}
			acts = append(acts, x.hrzBndRndrs[n])
		}
	}
	// vrtBnds
	if len(x.vrtBnds) != 0 {
		x.vrtBndRndrs = make([]*VrtBndRndrSeg, len(x.vrtBnds))
		for n, vrtBnd := range x.vrtBnds {
			x.vrtBndRndrs[n] = &VrtBndRndrSeg{VrtBnd: vrtBnd}
			acts = append(acts, x.vrtBndRndrs[n])
		}
	}
	sys.Run().Pll(acts...) // render all elements in pll
	// for _, a := range acts {
	// 	a.Act()
	// }
}
func (x *Stm) Draw(img *image.RGBA) {
	if x.x.Rng <= 0 || x.y.Rng <= 0 {
		x1 := float32(x.pos.X) + float32(x.siz.Width)*.5
		y1 := float32(x.pos.Y) + float32(x.siz.Height)*.5
		MsgFnt.Draw(x1, y1, .5, .5, MsgClr, img, TxtNoData)
		return
	}
	rct := vis.Rect(x.pos, x.siz) // draw order is z-order
	// sys.Log("Stm.Draw", "rct", rct)
	// sys.Log("Stm.Draw", "len(x.stmStkRndrs)", len(x.stmStkRndrs))
	// sys.Log("Stm.Draw", "len(x.stmBndRndrs)", len(x.stmBndRndrs))
	// sys.Log("Stm.Draw", "len(x.cndStkRndrs)", len(x.cndStkRndrs))
	// sys.Log("Stm.Draw", "len(x.hrzLnRndrs)", len(x.hrzLnRndrs))
	// sys.Log("Stm.Draw", "len(x.vrtLnRndrs)", len(x.vrtLnRndrs))
	// sys.Log("Stm.Draw", "len(x.hrzBndRndrs)", len(x.hrzBndRndrs))
	// sys.Log("Stm.Draw", "len(x.vrtBndRndrs)", len(x.vrtBndRndrs))

	// hrzBnds
	if len(x.hrzBndRndrs) != 0 {
		for _, seg := range x.hrzBndRndrs {
			if seg.ras.Size() == image.ZP {
				continue
			}
			seg.ras.Draw(img, rct, seg.filClr.Uniform(), image.Point{})
			seg.btmSeg.Stk.Draw(seg.btmSeg.pen.Clr, rct, img, &seg.btmSeg.ras)
			seg.topSeg.Stk.Draw(seg.topSeg.pen.Clr, rct, img, &seg.topSeg.ras)
		}
	}
	// vrtBnds
	if len(x.vrtBndRndrs) != 0 {
		for _, seg := range x.vrtBndRndrs {
			if seg.ras.Size() == image.ZP {
				continue
			}
			seg.ras.Draw(img, rct, seg.filClr.Uniform(), image.Point{})
			seg.lftSeg.Stk.Draw(seg.lftSeg.pen.Clr, rct, img, &seg.lftSeg.ras)
			seg.rhtSeg.Stk.Draw(seg.rhtSeg.pen.Clr, rct, img, &seg.rhtSeg.ras)
		}
	}
	// stmBnds
	if len(x.stmBndRndrs) != 0 {
		for _, seg := range x.stmBndRndrs {
			if seg.ras.Size() == image.ZP {
				continue
			}
			seg.ras.Draw(img, rct, seg.filClr.Uniform(), image.Point{})
			seg.btmSeg.Stk.Draw(seg.btmSeg.pen.Clr, rct, img, &seg.btmSeg.ras)
			seg.topSeg.Stk.Draw(seg.topSeg.pen.Clr, rct, img, &seg.topSeg.ras)
		}
	}
	// hrzLns
	if len(x.hrzLnRndrs) != 0 {
		for _, seg := range x.hrzLnRndrs {
			if seg.ras.Size() == image.ZP {
				continue
			}
			seg.Draw(seg.pen.Clr, rct, img, &seg.ras)
		}
	}
	// vrtLns
	if len(x.vrtLnRndrs) != 0 {
		for _, seg := range x.vrtLnRndrs {
			if seg.ras.Size() == image.ZP {
				continue
			}
			seg.Draw(seg.pen.Clr, rct, img, &seg.ras)
		}
	}
	// cndStks
	if len(x.cndStkRndrs) != 0 {
		for _, seg := range x.cndStkRndrs {
			if seg.ras.Size() == image.ZP { // cnd.Tmes may exist, but not for the rng timeframe; therefore use ras.Size()
				continue
			}
			// x1, y1 := seg.ras.Pen()
			// sys.Log("Stm.Draw", "cndStkRndrs: rasPen", x1, y1)
			// sys.Log("Stm.Draw", "cndStkRndrs: seg.ras.Size()", seg.ras.Size())
			// sys.Log("Stm.Draw", "seg.cnd.Tmes.Cnt()", seg.cnd.Tmes.Cnt())
			seg.Draw(seg.pen.Clr, rct, img, &seg.ras)
		}
	}
	// stmStks
	if len(x.stmStkRndrs) != 0 {
		for _, seg := range x.stmStkRndrs {
			if seg.ras.Size() == image.ZP {
				continue
			}
			seg.Draw(seg.pen.Clr, rct, img, &seg.ras)
		}
	}

	// brdr
	vis.PxlRct(x.brdr, BrdrClr.Color(), img)

	// x-axis txt
	if x.x.vis {
		y1 := float32(x.brdr.Btm + uint32(BrdrLen+AxisPad))
		InrvlFnt.Draw(float32(x.brdr.Lft+uint32(BrdrLen)), y1, 1, 1, InrvlTxtClrX, img, x.x.Min.String())
		for _, inrvl := range x.x.Inrvls[1 : len(x.x.Inrvls)-1] {
			InrvlFnt.Draw(float32(x.brdr.Lft)+x.x.Pxl(inrvl), y1, .5, 1, InrvlTxtClrX, img, inrvl.String())
		}
		InrvlFnt.Draw(float32(x.brdr.Rht-uint32(BrdrLen)), y1, 0, 1, InrvlTxtClrX, img, x.x.Max.String())
	}
	// y-axis txt
	x1 := float32(x.brdr.Rht + uint32(BrdrLen+AxisPad))
	InrvlFnt.Draw(x1, float32(x.brdr.Top+uint32(BrdrLen)), 1, 1, InrvlTxtClrY, img, x.y.Max.Trnc(AxisTrnc).String())
	for _, inrvl := range x.y.Inrvls[1 : len(x.y.Inrvls)-1] {
		InrvlFnt.Draw(x1, float32(x.brdr.Top)+x.y.Pxl(inrvl), 1, .5, InrvlTxtClrY, img, inrvl.Trnc(AxisTrnc).String())
	}
	InrvlFnt.Draw(x1, float32(x.brdr.Btm-uint32(BrdrLen)-4), 1, 0, InrvlTxtClrY, img, x.y.Min.Trnc(AxisTrnc).String())

	// plt title
	if x.Title != str.Empty {
		x1 := float32(x.brdr.Lft+uint32(BrdrLen)) + float32(x.x.Width)*.5 //- TitleFnt.Width(x.Title.Unquo())*.5
		y1 := float32(x.brdr.Btm-uint32(BrdrLen)-uint32(AxisPad)) - TitleFnt.Height()*.5
		TitleFnt.Draw(x1, y1, 0.5, 0.5, TitleClr, img, x.Title.Unquo())
	}
}
