package vis

import (
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"math"
	"os"
	"sys"
	"sys/ana/vis/clr"
	"sys/ana/vis/fnt"
	"sys/err"

	xdraw "golang.org/x/image/draw"
	"golang.org/x/image/font"
	"golang.org/x/image/math/f64"
	"golang.org/x/image/math/fixed"
	"golang.org/x/image/vector"
)

const (
	Tau64   = float64(6.28318530717958647692528676655900576839433879875021)
	Tau     = float32(Tau64)
	TauHlf  = float32(Tau / 2)
	TauQtr  = float32(Tau / 4)
	Tau2Qtr = float32(2 * TauQtr)
	Tau3Qtr = float32(3 * TauQtr)
	TauOct  = float32(Tau / 8)
	Tau2Oct = float32(2 * TauOct)
	Tau3Oct = float32(3 * TauOct)
	Tau4Oct = float32(4 * TauOct)
	Tau5Oct = float32(5 * TauOct)
	Tau6Oct = float32(6 * TauOct)
	Tau7Oct = float32(7 * TauOct)
	Pi      = TauHlf
)

const (
	LenMin = uint32(10)
)

// var ( // DEBUGGING
// 	pxlFns []func()
// 	jClr   = colornames.Red
// 	kClr   = colornames.Green
// 	pClr   = colornames.Blueviolet
// )

type (
	Vis struct {
		Img    *image.RGBA
		ras    vector.Rasterizer
		stkHlf float32 // half-width of stroke
		pth    []Pnt
		ks     []Pnt
	}
)

func NewVis(s Siz) (r *Vis) {
	r = &Vis{}
	r.PthCap(512)
	r.SetStk(1)
	r.Img = image.NewRGBA(image.Rect(0, 0, int(s.Width), int(s.Height)))
	r.ras.Reset(int(s.Width), int(s.Height))
	return r
}
func (x *Vis) SubVis(p Pos, s Siz) (r *Vis) {
	r = &Vis{}
	r.PthCap(cap(x.pth))
	r.stkHlf = x.stkHlf
	r.Img = x.Img.SubImage(image.Rect(int(p.X), int(p.Y), int(s.Width), int(s.Height))).(*image.RGBA)
	r.ras.Reset(int(s.Width), int(s.Height))
	return r
}
func (x *Vis) PngSav(path string) {
	file, er := os.Create(path)
	if er != nil {
		err.Panic(er)
	}
	defer file.Close()
	er = png.Encode(file, x.Img)
	if er != nil {
		err.Panic(er)
	}
}
func (x *Vis) DrawBak(c clr.Clr) {
	draw.Draw(x.Img, x.Img.Bounds(), c.Uniform(), image.Point{}, draw.Src)
}
func (x *Vis) DrawImg(p Pos, i image.Image) {
	draw.Draw(x.Img, x.Img.Bounds(), i, image.Point{X: -int(p.X), Y: -int(p.Y)}, draw.Over)
}
func (x *Vis) DrawVis(p Pos, v *Vis) {
	x.DrawImg(p, v.Img)
}
func (x *Vis) DrawStk(p Pos, s Siz, c clr.Clr, preserve ...bool) {
	// separate PthStk method to enable large qty stk rendering if needed
	// DrawStk draws a single image per path
	x.PthStk(preserve...)
	sys.Log("Vis.DrawStk", "x.Img.Rect", x.Img.Rect)
	sys.Log("Vis.DrawStk", "img.dx", x.Img.Rect.Dx(), "img.dy", x.Img.Rect.Dy())
	sys.Log("Vis.DrawStk", "ras.dx", x.ras.Bounds().Dx(), "ras.dy", x.ras.Bounds().Dy())
	sys.Log("Vis.DrawStk", "sp.dx", s.Rectangle(p).Dx(), "sp.dy", s.Rectangle(p).Dy())
	sys.Log("Vis.DrawStk", "p", p)
	x.ras.Draw(x.Img, s.Rectangle(p), c.Uniform(), image.Point{})
	x.ras.Reset(x.ras.Bounds().Dx(), x.ras.Bounds().Dy()) // clears buffer

	// x.ras.Draw(x.Img, x.Img.Rect, c.Uniform(), image.Point{})
	// x.ras.Reset(x.Img.Rect.Dx(), x.Img.Rect.Dy()) // clears buffer
	// for _, pxlFn := range pxlFns { // DEBUGGING
	// 	pxlFn()
	// }
}

// func (x *Vis) Width() uint32  { return uint32(x.Img.Rect.Dx()) }
// func (x *Vis) Height() uint32 { return uint32(x.Img.Rect.Dy()) }
// func (x *Vis) Siz() (r Siz) {
// 	r.Width = uint32(x.Img.Rect.Dy())
// 	r.Height = uint32(x.Img.Rect.Dy())
// 	return r
// }
func (x *Vis) SetStk(stk float32) { x.stkHlf = stk * 0.5 }
func (x *Vis) PthCap(size int) {
	x.pth = make([]Pnt, 0, size)
	x.ks = make([]Pnt, 0, size)
}
func (x *Vis) PthClr() {
	x.pth = x.pth[:0]
	x.ks = x.ks[:0]
}
func (x *Vis) PthLen() int        { return len(x.pth) }
func (x *Vis) Pth(x1, y1 float32) { x.pth = append(x.pth, Pnt{X: x1, Y: y1}) }

func (x *Vis) PthStk(preserve ...bool) {
	if len(x.pth) >= 2 {
		// --- 1ST PNT; 1ST SEG ---
		cur := NewStkSeg(x.stkHlf, x.pth[0], x.pth[1])
		x.ras.MoveTo(cur.P1.J.X, cur.P1.J.Y)
		x.ks = append(x.ks, cur.P1.K)
		prv := cur
		// CUR := cur                       // DEBUGGING
		// pxlFns = append(pxlFns, func() { // DEBUGGING
		// 	x.PxlPnt(CUR.P1.Cntr, pClr)
		// 	x.PxlPnt(CUR.P1.J, jClr)
		// 	x.PxlPnt(CUR.P1.K, kClr)
		// 	x.PxlPnt(CUR.P2.J, jClr)
		// 	x.PxlPnt(CUR.P2.K, kClr)
		// })
		// --- MDL PNTS ---
		for n := 1; n < len(x.pth); n++ {
			cur = NewStkSeg(x.stkHlf, x.pth[n-1], x.pth[n])

			if prv.Dir.X == cur.Dir.X && prv.Dir.Y == cur.Dir.Y { // co-linear segments
				x.ras.LineTo(cur.P1.J.X, cur.P1.J.Y)
				x.ks = append(x.ks, cur.P1.K)
				// CUR := cur                       // DEBUGGING
				// pxlFns = append(pxlFns, func() { // DEBUGGING
				// 	x.PxlPnt(CUR.P1.J, jClr)
				// 	x.PxlPnt(CUR.P1.K, kClr)
				// })
			} else if n != 1 {
				if prv.P1.Cntr.Dist(cur.P2.Cntr) < prv.P1.K.Dist(cur.P2.Cntr) {
					// sys.Log("--- P")
					// JOIN: J-SIDE RAY INTERSECT
					jI := Intersect(prv.P1.J, cur.P2.J, prv.Dir, cur.Dir.RotHlf())
					x.ras.LineTo(jI.X, jI.Y)
					// pxlFns = append(pxlFns, func() { // DEBUGGING
					// 	x.PxlPnt(jI, jClr)
					// })
					// JOIN: K-SIDE DOUBLE PNT
					x.ks = append(x.ks, prv.P2.K)
					x.ks = append(x.ks, cur.P1.K)
				} else {
					// sys.Log("--- Q")
					// JOIN: K-SIDE RAY INTERSECT
					kI := Intersect(prv.P1.K, cur.P2.K, prv.Dir, cur.Dir.RotHlf())
					x.ks = append(x.ks, kI)
					// pxlFns = append(pxlFns, func() { // DEBUGGING
					// 	x.PxlPnt(kI, kClr)
					// })
					// JOIN: J-SIDE DOUBLE PNT
					x.ras.LineTo(prv.P2.J.X, prv.P2.J.Y)
					x.ras.LineTo(cur.P1.J.X, cur.P1.J.Y)
				}
			}
			prv = cur
			// CUR := cur                       // DEBUGGING
			// pxlFns = append(pxlFns, func() { // DEBUGGING
			// 	x.PxlPnt(CUR.P1.Cntr, pClr)
			// 	x.PxlPnt(CUR.P1.J, jClr)
			// 	x.PxlPnt(CUR.P1.K, kClr)
			// 	x.PxlPnt(CUR.P2.Cntr, pClr)
			// 	x.PxlPnt(CUR.P2.J, jClr)
			// 	x.PxlPnt(CUR.P2.K, kClr)
			// })
		}
		// --- LST PNT; LST SEG ---
		x.ras.LineTo(prv.P2.J.X, prv.P2.J.Y)
		x.ks = append(x.ks, prv.P2.K)
		for n := len(x.ks) - 1; n > -1; n-- { // draw k-pnts
			x.ras.LineTo(x.ks[n].X, x.ks[n].Y)
		}
		x.ras.ClosePath()
	}
	if len(preserve) == 0 || !preserve[0] {
		x.PthClr()
	}
	//image.NewUniform(c)
}

func (x *Vis) PthFil(opn ...bool) {}

func (x *Vis) PntStk(x1, y1 float32) { // 1 stk sqr
	x.ras.MoveTo(x1-x.stkHlf, y1-x.stkHlf)
	x.ras.LineTo(x1+x.stkHlf, y1-x.stkHlf)
	x.ras.LineTo(x1+x.stkHlf, y1+x.stkHlf)
	x.ras.LineTo(x1-x.stkHlf, y1+x.stkHlf)
	x.ras.ClosePath()
}

func (x *Vis) Txtf(x1, y1, offX, offY float32, fnt *fnt.Fnt, clr clr.Clr, format string, args ...interface{}) {
	x.Txt(x1, y1, offX, offY, fnt, clr, fmt.Sprintf(format, args...))
}
func (x *Vis) Txt(x1, y1, offX, offY float32, fnt *fnt.Fnt, clr clr.Clr, txt string) {
	fnt.Mu.Lock()
	d := &font.Drawer{Dst: x.Img, Src: clr.Uniform(), Face: fnt.Face}
	w := float32(d.MeasureString(txt) >> 6)
	h := float32(d.Face.Metrics().Height >> 6)
	d.Dot = fixp(x1-w+(offX*w), y1+(offY*h))
	// based on Drawer.DrawString() in golang.org/x/image/font/font.go
	prevC := rune(-1)
	for _, c := range txt {
		if prevC >= 0 {
			d.Dot.X += d.Face.Kern(prevC, c)
		}
		dr, mask, maskp, advance, ok := d.Face.Glyph(d.Dot, c)
		if !ok {
			continue
		}
		sr := dr.Sub(dr.Min)
		s2d := f64.Aff3{1, 0, float64(dr.Min.X), 0, 1, float64(dr.Min.Y)}
		transformer := xdraw.BiLinear
		transformer.Transform(d.Dst, s2d, d.Src, sr, xdraw.Over, &xdraw.Options{
			SrcMask:  mask,
			SrcMaskP: maskp,
		})
		d.Dot.X += advance
		prevC = c
	}
	fnt.Mu.Unlock()
}

func minMaxI(x, y int) (int, int) {
	if x < y {
		return x, y
	}
	return y, x
}
func minMax(x, y float32) (float32, float32) {
	if x < y {
		return x, y
	}
	return y, x
}
func min(x, y float32) float32 {
	if x < y {
		return x
	}
	return y
}
func max(x, y float32) float32 {
	if x > y {
		return x
	}
	return y
}
func abs(v float32) float32 {
	if v < 0 {
		return -v
	}
	return v
}
func ceil(v float32) float32 {
	return float32(int(v) + 1)
}
func dist(x1, y1, x2, y2 float32) float32 {
	return float32(math.Hypot(float64(x1)-float64(x2), float64(y1)-float64(y2)))
}
func slope(x1, y1, x2, y2 float32) float32 {
	return (y2 - y1) / (x2 - x1)
}

func fixp(x, y float32) fixed.Point26_6 {
	return fixed.Point26_6{X: fix(x), Y: fix(y)}
}
func fix(x float32) fixed.Int26_6 {
	return fixed.Int26_6(x * 64)
}
