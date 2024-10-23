package plt

import (
	"sys"
	"sys/ana/hst"
	"sys/ana/vis"
	"sys/ana/vis/pen"
	"sys/bsc/unt"

	"golang.org/x/image/vector"
)

type (
	CndStk struct {
		cnd *hst.CndBse
		pen pen.Pen
		plt *TmeFltPltBse
	}
	CndStkRndrSeg struct {
		*CndStk
		vis.Stk
		ras vector.Rasterizer
	}
)

func (x *CndStkRndrSeg) Act() {
	if x.cnd.Tmes.Cnt() == 0 {
		// sys.Logf("CndStkRndrSeg: NO TMES (cnt:%v) %v", x.cnd.Tmes.Cnt(), x.cnd)
		return
	}
	// filter for cnd tmes within the visible range
	idxMin := x.cnd.Tmes.SrchIdx(x.plt.x.Min, true) // may be tme.Min
	if idxMin == unt.Max {
		sys.Logf("CndStkRndrSeg: MIN IDX LOOKUP FAILED (Min:%v) %v", x.plt.x.Min, x.cnd)
		return
	}
	idxMax := x.cnd.Tmes.SrchIdx(x.plt.x.Max, true) // may be tme.Max
	if idxMax == unt.Max {
		sys.Logf("CndStkRndrSeg: MAX IDX LOOKUP FAILED (Max:%v) %v", x.plt.x.Max, x.cnd)
		return
	}
	ts := x.cnd.Tmes.In(idxMin, idxMax+1) // visible cnd tmes
	if ts.Lst() > x.plt.x.Max {
		ts = ts.To(ts.Cnt() - 1)
	}
	if len(*ts) != 0 {
		x.Reset(2 * len(*ts))
		x.ras.Reset(int(x.plt.siz.Width), int(x.plt.siz.Height))
		for n := 0; n < len(*ts); n++ {
			// sys.Logf("%p CndStkRndrSeg.Act %v %v \n", x, n, (*ts)[n])
			// calculate pth
			x1 := x.plt.off.X + x.plt.x.Pxl((*ts)[n])
			x.PthClr()
			x.Pth(x1, x.plt.off.Y)
			x.Pth(x1, x.plt.off.Y+float32(x.plt.y.Height))
			// // rasterize pth
			x.Rndr(float32(x.pen.Wid), &x.ras)
		}
		// x.Rndr(x.wid, &x.ras)
	}

}
