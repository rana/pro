package plt

import (
	"sys"
	"sys/ana/vis/clr"

	"golang.org/x/image/vector"
)

type (
	StmBnd struct {
		btm    *StmStk
		top    *StmStk
		filClr clr.Clr
		filRas vector.Rasterizer
	}
	StmBndMeasureSeg struct {
		*StmBnd
		btmSeg *StmStkMeasureSeg
		topSeg *StmStkMeasureSeg
	}
	StmBndRndrSeg struct {
		*StmBnd
		btmSeg *StmStkRndrSeg
		topSeg *StmStkRndrSeg
		ras    vector.Rasterizer
	}
)

func (x *StmBndMeasureSeg) Act() {
	x.btmSeg = &StmStkMeasureSeg{StmStk: x.btm}
	x.topSeg = &StmStkMeasureSeg{StmStk: x.top}
	acts := []sys.Act{x.btmSeg, x.topSeg}
	sys.Run().Pll(acts...)
}
func (x *StmBndRndrSeg) Act() {
	x.btmSeg = &StmStkRndrSeg{StmStk: x.btm}
	x.topSeg = &StmStkRndrSeg{StmStk: x.top}
	acts := []sys.Act{x.btmSeg, x.topSeg}
	sys.Run().Pll(acts...)
	// rasterize area points
	btmPth := x.btmSeg.GetPth()
	topPth := x.topSeg.GetPth()
	if len(btmPth) >= 2 && len(topPth) >= 2 {
		x.ras.Reset(int(x.btm.plt.siz.Width), int(x.btm.plt.siz.Height))
		x.ras.MoveTo(btmPth[0].X, btmPth[0].Y)
		for n := 1; n < len(btmPth); n++ {
			x.ras.LineTo(btmPth[n].X, btmPth[n].Y)
		}
		for n := len(topPth) - 1; n > -1; n-- {
			x.ras.LineTo(topPth[n].X, topPth[n].Y)
		}
		x.ras.ClosePath()
	}
}
