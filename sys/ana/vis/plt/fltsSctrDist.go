package plt

import (
	"image"
	"sys"
	"sys/ana"
	"sys/ana/vis"
	"sys/ana/vis/clr"
	"sys/bsc/bnd"
	"sys/bsc/bnds"
	"sys/bsc/flt"
	"sys/bsc/flts"
	"sys/bsc/unt"

	"golang.org/x/image/vector"
)

type (
	SctrDistItm struct {
		Vals     *flts.Flts
		ValsDist *flts.Flts
		clr      clr.Clr
		radius   uint32
		plt      *FltsSctrDist
	}
	SctrDistRndrSeg struct {
		bnd.Bnd
		plt  *FltsSctrDist
		itm  *SctrDistItm
		pnts []SctrDistRndrPnt
	}
	SctrDistRndrPnt struct {
		ras    vector.Rasterizer
		x1     float32
		x2     float32
		y1     float32
		radius float32
		clr    clr.Clr
	}
)

func (x *FltsSctrDist) Measure() vis.Siz { // Plt interface
	x.Y.Min, x.Y.Max = flt.Max, flt.Min
	for _, v := range x.itms {
		min, max := v.Vals.MinMax() // MinMax IS PLL
		x.Y.Min = x.Y.Min.Min(min)
		x.Y.Max = x.Y.Max.Max(max)
	}
	if x.Y.Min == flt.Max || x.Y.Max == flt.Min {
		x.Y.Rng = 0
	} else {
		x.Y.Rng = x.Y.Max - x.Y.Min
	}
	if x.Y.Rng != 0 {
		// y-axis txt
		x.Y.MeasureInrvls(ana.Cfg.Plt.InrvlCnt)
		x.Y.Rht.Width = uint32(AxisPad) + uint32(InrvlTxtLen) // constant txt len to avoid variable radiusths where cnd lns don't align
		// y-axis
		if x.Y.PxlPerVal != 0 { // define height on user-defined scl
			valPerPxl := (1.0 / x.Y.PxlPerVal)
			x.Y.Height = uint32(float32(x.Y.Rng) / valPerPxl)
			x.siz.Height = x.Y.Height + x.mrgn.Height() + uint32(BrdrLen)*2
		} else { // measure height based on plt-defined segment vals
			x.Y.Height = x.siz.Height - x.mrgn.Height() - uint32(BrdrLen)*2 - uint32(BarPad)*2
			x.Y.PxlPerVal = float32(x.Y.Height) / float32(x.Y.Rng)
		}

	}
	// x-axis
	x.siz.Width = x.mrgn.Width() + uint32(BrdrLen)*2 + x.Y.Rht.Width + x.XWidth + uint32(BarPad)*2
	return x.siz
}
func (x *SctrDistRndrSeg) Act() {
	var pntCnt unt.Unt
	for m := x.Idx; m < x.Lim; m++ { // count pnts within current seg
		pntCnt += x.plt.itms[m].Vals.Cnt()
	}

	if pntCnt != 0 {
		x.pnts = make([]SctrDistRndrPnt, pntCnt) // ensure alloc for draw
		acts := make([]sys.Act, pntCnt)
		var pntIdx int
		for m := x.Idx; m < x.Lim; m++ { // generate a rendr for each point within seg
			itm := x.plt.itms[m]
			if itm.Vals != nil {

				// sys.Log("         len(x.pnts)", len(x.pnts))
				// sys.Log("      itm.Vals.Cnt()", itm.Vals.Cnt())
				// sys.Log("x.itm.ValsDist.Cnt()", x.itm.ValsDist.Cnt())
				for valIdx, val := range *itm.Vals {
					// sys.Log("     pntIdx", pntIdx)
					// sys.Log("x.pnts == nil", x.pnts == nil)
					// sys.Log("x.plt == nil", x.plt == nil)
					// sys.Log("x.itm == nil", x.itm == nil)
					// sys.Log("x.itm.ValsDist == nil", x.itm.ValsDist == nil)
					x.pnts[pntIdx].ras.Reset(int(x.plt.siz.Width), int(x.plt.siz.Height))
					x.pnts[pntIdx].x1 = x.plt.off.X + float32(x.plt.itms[m].radius)
					x.pnts[pntIdx].x2 = x.pnts[pntIdx].x1 + (float32(x.plt.XWidth) * float32((*x.itm.ValsDist)[valIdx])) - float32(x.plt.itms[m].radius)
					x.pnts[pntIdx].y1 = x.plt.off.Y + x.plt.Y.Pxl(val)
					x.pnts[pntIdx].radius = float32(x.plt.itms[m].radius)
					x.pnts[pntIdx].clr = x.plt.itms[m].clr
					acts[pntIdx] = &x.pnts[pntIdx]
					pntIdx++
				}
			}
		}
		sys.Run().Pll(acts...)
	}
}
func (x *SctrDistRndrPnt) Act() {
	vis.RndrRctHrz(x.x1, x.x2, x.y1, x.radius, &x.ras)
}

func (x *FltsSctrDist) Rndr() {
	if x.Y.Rng <= 0 { // no data
		//sys.Logf("%p Rng.Rndr INVALID (Y.Rng:%v Y.Min:%v Y.Max:%v) \n", x, x.Y.Rng, x.Y.Min, x.Y.Max)
		return
	}

	// sys.Log("FltsSctrDist.Rndr", "x.pos", x.pos, "x.siz", x.siz)
	// brdr (for pxl pnts written directly to img)
	// x.pos known only after Measure call
	x.brdr.Lft = x.pos.X + x.mrgn.Lft + uint32(BrdrLen)
	x.brdr.Rht = x.brdr.Lft + x.XWidth + uint32(BrdrLen) + uint32(BarPad)*2
	x.brdr.Top = x.pos.Y + x.mrgn.Top + uint32(BrdrLen)
	x.brdr.Btm = x.brdr.Top + x.Y.Height + uint32(BrdrLen) + uint32(BarPad)*2
	// off (for stk/fil pnts rasterized within brdr)
	x.off.X = float32(x.brdr.Lft - x.pos.X + uint32(BrdrLen) + uint32(BarPad))
	x.off.Y = float32(x.brdr.Top - x.pos.Y + uint32(BrdrLen) + uint32(BarPad))
	// sys.Log("FltsSctrDist.Rndr", "x.brdr", x.brdr)
	// sys.Log("FltsSctrDist.Rndr", "x.off", x.off)
	// sys.Log("FltsSctrDist.Rndr", "len(x.itms)", len(x.itms))

	// itms
	if len(x.itms) != 0 {
		segBnds, acts := bnds.Segs(unt.Unt(len(x.itms)))
		x.rndrs = make([]*SctrDistRndrSeg, len(*segBnds))
		for n, segBnd := range *segBnds {
			x.rndrs[n] = &SctrDistRndrSeg{
				Bnd: segBnd,
				itm: x.itms[n],
				plt: x,
			}
			acts[n] = x.rndrs[n]
		}
		sys.Run().Pll(acts...) // render all elements in pll
	}
}
func (x *FltsSctrDist) Draw(img *image.RGBA) {
	if x.Y.Rng <= 0 {
		x1 := float32(x.pos.X) + float32(x.siz.Width)*.5
		y1 := float32(x.pos.Y) + float32(x.siz.Height)*.5
		MsgFnt.Draw(x1, y1, .5, .5, MsgClr, img, TxtNoData)
		return
	}
	rct := vis.Rect(x.pos, x.siz) // draw order is z-order
	// // shp
	// if x.shp.ras.Size() != image.ZP {
	// 	x.shp.ras.Draw(img, rct, x.shp.clr.Uniform(), image.Point{})
	// }
	for _, rndr := range x.rndrs {
		if rndr.pnts != nil {
			for _, pnt := range rndr.pnts {
				if pnt.ras.Size() != image.ZP {
					pnt.ras.Draw(img, rct, pnt.clr.Uniform(), image.Point{})
				}
			}
		}
	}
	// brdr
	vis.PxlRct(x.brdr, BrdrClr.Color(), img)
	// y-axis txt
	x1 := float32(x.brdr.Rht + uint32(BrdrLen+AxisPad))
	// InrvlFnt.Draw(x1, float32(x.brdr.Top+uint32(BarPad)+uint32(BrdrLen)), 1, 1, InrvlTxtClrY, img, x.Y.Max.Trnc(AxisTrnc).String())
	for _, inrvl := range x.Y.Inrvls {
		InrvlFnt.Draw(x1, float32(x.brdr.Top+uint32(BarPad))+x.Y.Pxl(inrvl), 1, .5, InrvlTxtClrY, img, inrvl.Trnc(AxisTrnc).String())
	}
	// InrvlFnt.Draw(x1, float32(x.brdr.Btm-uint32(BarPad)-uint32(BrdrLen)-4), 1, 0, InrvlTxtClrY, img, x.Y.Min.Trnc(AxisTrnc).String())
}
