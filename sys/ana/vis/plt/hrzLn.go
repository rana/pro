package plt

import (
	"sys/ana/vis"
	"sys/ana/vis/pen"
	"sys/bsc/flt"

	"golang.org/x/image/vector"
)

type (
	HrzLn struct {
		val flt.Flt
		pen pen.Pen
		plt *TmeFltPltBse
	}
	HrzLnRndrSeg struct {
		*HrzLn
		vis.Stk
		ras vector.Rasterizer
	}
)

func (x *HrzLnRndrSeg) Act() {
	x.Reset(2)
	// calculate pth
	y1 := x.plt.off.Y + x.plt.y.Pxl(x.val)
	x.Pth(x.plt.off.X, y1)
	x.Pth(x.plt.off.X+float32(x.plt.x.Width), y1)
	// rasterize pth
	x.ras.Reset(int(x.plt.siz.Width), int(x.plt.siz.Height))
	x.Rndr(float32(x.pen.Wid), &x.ras)
}
