package vis

import "golang.org/x/image/vector"

func RndrRct(cntrX, cntrY, widHlf, heiHlf float32, ras *vector.Rasterizer) {
	ras.MoveTo(cntrX-widHlf, cntrY-heiHlf)
	ras.LineTo(cntrX+widHlf, cntrY-heiHlf)
	ras.LineTo(cntrX+widHlf, cntrY+heiHlf)
	ras.LineTo(cntrX-widHlf, cntrY+heiHlf)
	ras.ClosePath()
}
func RndrSqr(cntrX, cntrY, radius float32, ras *vector.Rasterizer) {
	ras.MoveTo(cntrX-radius, cntrY-radius)
	ras.LineTo(cntrX+radius, cntrY-radius)
	ras.LineTo(cntrX+radius, cntrY+radius)
	ras.LineTo(cntrX-radius, cntrY+radius)
	ras.ClosePath()
}
func RndrRctHrz(x1, x2, y1, radius float32, ras *vector.Rasterizer) {
	ras.MoveTo(x1, y1-radius)
	ras.LineTo(x1, y1+radius)
	ras.LineTo(x2, y1+radius)
	ras.LineTo(x2, y1-radius)
	ras.ClosePath()
}
