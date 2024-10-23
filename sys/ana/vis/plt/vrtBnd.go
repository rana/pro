package plt

import (
	"sys"
	"sys/ana/vis/clr"

	"golang.org/x/image/vector"
)

type (
	VrtBnd struct {
		lft    *VrtLn
		rht    *VrtLn
		filClr clr.Clr
		filRas vector.Rasterizer
	}
	VrtBndRndrSeg struct {
		*VrtBnd
		lftSeg *VrtLnRndrSeg
		rhtSeg *VrtLnRndrSeg
		ras    vector.Rasterizer
	}
)

func (x *VrtBndRndrSeg) Act() {
	x.lftSeg = &VrtLnRndrSeg{VrtLn: x.lft}
	x.rhtSeg = &VrtLnRndrSeg{VrtLn: x.rht}
	acts := []sys.Act{x.lftSeg, x.rhtSeg}
	sys.Run().Pll(acts...)
	// rasterize area points
	lftPth := x.lftSeg.GetPth()
	rhtPth := x.rhtSeg.GetPth()
	if len(lftPth) >= 2 && len(rhtPth) >= 2 {
		x.ras.Reset(int(x.lft.plt.siz.Width), int(x.lft.plt.siz.Height))
		x.ras.MoveTo(lftPth[0].X, lftPth[0].Y)
		x.ras.LineTo(lftPth[1].X, lftPth[1].Y)
		x.ras.LineTo(rhtPth[1].X, rhtPth[1].Y)
		x.ras.LineTo(rhtPth[0].X, rhtPth[0].Y)
		x.ras.ClosePath()
	}
}
