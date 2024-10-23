package plt

import (
	"sys/ana/vis"
	"sys/ana/vis/pen"
	"sys/bsc/tme"

	"golang.org/x/image/vector"
)

type (
	VrtLn struct {
		val tme.Tme
		pen pen.Pen
		plt *TmeFltPltBse
	}
	VrtLnRndrSeg struct {
		*VrtLn
		vis.Stk
		ras vector.Rasterizer
	}
)

func (x *VrtLnRndrSeg) Act() {
	x.Reset(2)
	// calculate pth
	x1 := x.plt.off.X + x.plt.x.Pxl(x.val)
	x.Pth(x1, x.plt.off.Y)
	x.Pth(x1, x.plt.off.Y+float32(x.plt.y.Height))
	// rasterize pth
	x.ras.Reset(int(x.plt.siz.Width), int(x.plt.siz.Height))
	x.Rndr(float32(x.pen.Wid), &x.ras)
}
