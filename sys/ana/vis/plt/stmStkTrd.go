package plt

import (
	"sys/ana"
	"sys/ana/hst"
	"sys/ana/vis"
	"sys/ana/vis/pen"
	"sys/bsc/flt"
	"sys/bsc/flts"
	"sys/bsc/tme"
	"sys/bsc/tmes"

	"golang.org/x/image/vector"
)

type (
	StmStkTrd struct {
		X   *tmes.Tmes
		Y   *flts.Flts
		trd *ana.Trd
		stm *hst.StmBse
		pen pen.Pen
		plt *TmeFltPltBse

		MinY    flt.Flt
		MaxY    flt.Flt
		OffY    flt.Flt
		BtmLenY flt.Flt
		TopLenY flt.Flt
	}
	StmStkTrdMeasureSeg struct {
		*StmStkTrd
		RngX tme.Tme
	}
	StmStkTrdRndrSeg struct {
		*StmStkTrd
		vis.Stk
		ras vector.Rasterizer
	}
)

func (x *StmStkTrdMeasureSeg) Act() {
	if x.stm.Tmes.Cnt() < 2 {
		// sys.Logf("StmStkTrd: TOO FEW PNTS (cnt:%v) %v", x.stm.Tmes.Cnt(), x.stm)
		return
	}
	idxFst := x.stm.Tmes.SrchIdxEql(x.trd.OpnTme)
	if idxFst >= x.stm.Tmes.Cnt() || (*x.stm.Tmes)[idxFst] != x.trd.OpnTme {
		return
	}
	idxLst := x.stm.Tmes.SrchIdxEql(x.trd.ClsTme)
	if idxLst >= x.stm.Tmes.Cnt() || (*x.stm.Tmes)[idxLst] != x.trd.ClsTme {
		return
	}
	x.X = x.stm.Tmes.In(idxFst, idxLst+1)
	x.Y = x.stm.Vals.In(idxFst, idxLst+1)
	x.MinY, x.MaxY = x.Y.MinMax()
	x.RngX = x.X.Lst() - x.X.Fst()
}

func (x *StmStkTrdRndrSeg) Act() {
	if x.X == nil {
		return
	}
	samplCnt := 1
	if x.plt.sampl && len(*x.X) > int(x.plt.siz.Width) {
		samplCnt = len(*x.X) / int(x.plt.siz.Width)
		if samplCnt < 0 {
			samplCnt = 1
		}
	}
	x.Reset(int(float32(len(*x.X)) * 1.5))
	// start fst y pnt at rng mid point for all segs
	// calculate pth
	for n := 0; n < len(*x.X); n += samplCnt {
		//float32(v-x.Min) * x.PxlPerVal
		// x.Pth(x.plt.off.X+x.plt.x.Pxl((*x.X)[n]), x.plt.off.Y+x.plt.y.Pxl((*x.Y)[n]))
		// x.Pth(
		// 	x.plt.off.X+(float32((*x.X)[n]-(*x.X)[0])*x.plt.x.PxlPerVal),
		// 	x.plt.off.Y+x.plt.y.Pxl((*x.Y)[n]))
		x.Pth(
			x.plt.off.X+(float32((*x.X)[n]-(*x.X)[0])*x.plt.x.PxlPerVal),
			x.plt.off.Y+float32(x.plt.y.Height)-(float32((*x.Y)[n]-x.MinY+x.OffY)*x.plt.y.PxlPerVal),
		)
	}
	// rasterize pth
	x.ras.Reset(int(x.plt.siz.Width), int(x.plt.siz.Height))
	x.Rndr(float32(x.StmStkTrd.pen.Wid), &x.ras)
}
