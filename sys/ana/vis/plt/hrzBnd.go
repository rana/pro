package plt

import (
	"sys"
	"sys/ana/vis/clr"

	"golang.org/x/image/vector"
)

type (
	HrzBnd struct {
		btm    *HrzLn
		top    *HrzLn
		filClr clr.Clr
		filRas vector.Rasterizer
	}
	HrzBndRndrSeg struct {
		*HrzBnd
		btmSeg *HrzLnRndrSeg
		topSeg *HrzLnRndrSeg
		ras    vector.Rasterizer
	}
)

func (x *HrzBndRndrSeg) Act() {
	x.btmSeg = &HrzLnRndrSeg{HrzLn: x.btm}
	x.topSeg = &HrzLnRndrSeg{HrzLn: x.top}
	acts := []sys.Act{x.btmSeg, x.topSeg}
	sys.Run().Pll(acts...)
	// rasterize area points
	btmPth := x.btmSeg.GetPth()
	topPth := x.topSeg.GetPth()
	if len(btmPth) >= 2 && len(topPth) >= 2 {
		x.ras.Reset(int(x.btm.plt.siz.Width), int(x.btm.plt.siz.Height))
		x.ras.MoveTo(btmPth[0].X, btmPth[0].Y)
		x.ras.LineTo(btmPth[1].X, btmPth[1].Y)
		x.ras.LineTo(topPth[1].X, topPth[1].Y)
		x.ras.LineTo(topPth[0].X, topPth[0].Y)
		x.ras.ClosePath()
	}
}
