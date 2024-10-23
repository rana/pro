package plt

import (
	"sys/ana/hst"
	"sys/ana/vis"
	"sys/ana/vis/pen"
	"sys/bsc/flt"
	"sys/bsc/flts"
	"sys/bsc/tmes"

	"golang.org/x/image/vector"
)

type (
	StmStk struct {
		X   *tmes.Tmes
		Y   *flts.Flts
		stm *hst.StmBse
		pen pen.Pen
		plt *TmeFltPltBse
	}
	StmStkMeasureSeg struct {
		*StmStk
		MinY flt.Flt
		MaxY flt.Flt
	}
	StmStkRndrSeg struct {
		*StmStk
		vis.Stk
		ras vector.Rasterizer
	}
)

func (x *StmStkMeasureSeg) Act() {
	if x.stm.Tmes.Cnt() < 2 {
		// sys.Logf("StmStk: TOO FEW PNTS (cnt:%v) %v", x.stm.Tmes.Cnt(), x.stm)
		return
	}
	// idxMin := x.stm.Tmes.SrchIdx(x.plt.X.Min, true) // may be tme.Min
	// if idxMin == unt.Max {
	// 	sys.Logf("StkStm: MIN IDX LOOKUP FAILED (Min:%v) %v", x.plt.X.Min, x.stm)
	// 	return
	// }
	// idxMax := x.stm.Tmes.SrchIdx(x.plt.X.Max, true) // may be tme.Max
	// if idxMax == unt.Max {
	// 	sys.Logf("StkStm: MAX IDX LOOKUP FAILED (Max:%v) %v", x.plt.X.Max, x.stm)
	// 	return
	// }
	// x.X = x.stm.Tmes.In(idxMin, idxMax+1)
	// x.Y = x.stm.Vals.In(idxMin, idxMax+1)
	x.X = x.stm.Tmes
	x.Y = x.stm.Vals
	x.MinY, x.MaxY = x.Y.MinMax()
}
func (x *StmStkRndrSeg) Act() {
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
	// TODO: DO NOT SKP PNTS WHICH HAVE A CND TME?
	// calculate pth
	for n := 0; n < len(*x.X); n += samplCnt {
		x.Pth(x.plt.off.X+x.plt.x.Pxl((*x.X)[n]), x.plt.off.Y+x.plt.y.Pxl((*x.Y)[n]))
	}
	// rasterize pth
	x.ras.Reset(int(x.plt.siz.Width), int(x.plt.siz.Height))
	x.Rndr(float32(x.StmStk.pen.Wid), &x.ras)
}
