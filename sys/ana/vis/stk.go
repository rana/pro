package vis

import (
	"image"
	"math"
	"sys/ana/vis/clr"

	"golang.org/x/image/vector"
)

type (
	Stk struct {
		pth []Pnt
		ks  []Pnt
	}
	StkPnt struct {
		Cntr Pnt
		J    Pnt
		K    Pnt
	}
	StkSeg struct {
		Dir Pnt
		P1  StkPnt
		P2  StkPnt
	}
)

func NewStk() (r *Stk) {
	r = &Stk{}
	r.pth = make([]Pnt, 0, 512)
	r.ks = make([]Pnt, 0, 512)
	return r
}
func (x *Stk) Reset(cap int) {
	x.pth = make([]Pnt, 0, cap)
	x.ks = make([]Pnt, 0, cap)
}
func (x *Stk) PthClr() {
	x.pth = x.pth[:0]
	x.ks = x.ks[:0]
}
func (x *Stk) Pth(x1, y1 float32) { x.pth = append(x.pth, Pnt{X: x1, Y: y1}) }
func (x *Stk) GetPth() []Pnt      { return x.pth }

// Rndr rasterizes path points.
func (x *Stk) Rndr(width float32, ras *vector.Rasterizer) {
	if len(x.pth) >= 2 {
		// for _, pnt := range x.pth {
		// 	sys.Logf("v.Pth(%v, %v)\n", pnt.X, pnt.Y)
		// }

		// ras.Reset(int(s.Width), int(s.Height))
		stkHlf := width * .5
		// --- 1ST PNT; 1ST SEG ---
		cur := NewStkSeg(stkHlf, x.pth[0], x.pth[1])
		ras.MoveTo(cur.P1.J.X, cur.P1.J.Y)
		x.ks = append(x.ks, cur.P1.K)
		prv := cur
		// sys.Log("fst", cur)
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
			cur = NewStkSeg(stkHlf, x.pth[n-1], x.pth[n])
			// sys.Log("cur", cur)

			if prv.Dir.X == cur.Dir.X && prv.Dir.Y == cur.Dir.Y { // co-linear segments
				ras.LineTo(cur.P1.J.X, cur.P1.J.Y)
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
					// sys.Log("prv.P1.J", prv.P1.J, prv)
					// sys.Log("cur.P2.J", cur.P2.J, cur)
					jI := Intersect(prv.P1.J, cur.P2.J, prv.Dir, cur.Dir.RotHlf())
					if jI != jI { // IsNaN
						ras.LineTo(prv.P1.J.X, prv.P1.J.Y)
					} else {
						ras.LineTo(jI.X, jI.Y)
						// pxlFns = append(pxlFns, func() { // DEBUGGING
						// 	x.PxlPnt(jI, jClr)
						// })
					}

					// JOIN: K-SIDE DOUBLE PNT
					x.ks = append(x.ks, prv.P2.K)
					x.ks = append(x.ks, cur.P1.K)
				} else {
					// sys.Log("--- Q")
					// JOIN: K-SIDE RAY INTERSECT
					// sys.Log("prv.P1.K", prv.P1.K)
					// sys.Log("prv.P2.K", prv.P2.K)
					kI := Intersect(prv.P1.K, cur.P2.K, prv.Dir, cur.Dir.RotHlf())
					if kI != kI { // IsNaN
						ras.LineTo(prv.P1.K.X, prv.P1.K.Y)
					} else {
						x.ks = append(x.ks, kI)
						// pxlFns = append(pxlFns, func() { // DEBUGGING
						// 	x.PxlPnt(kI, kClr)
						// })
					}
					// JOIN: J-SIDE DOUBLE PNT
					ras.LineTo(prv.P2.J.X, prv.P2.J.Y)
					ras.LineTo(cur.P1.J.X, cur.P1.J.Y)
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
		// sys.Log(x.Ras.Pen())
		// sys.Log("LST: prv.P2.J", prv.P2.J)
		ras.LineTo(prv.P2.J.X, prv.P2.J.Y)
		x.ks = append(x.ks, prv.P2.K)
		for n := len(x.ks) - 1; n > -1; n-- { // draw k-pnts
			ras.LineTo(x.ks[n].X, x.ks[n].Y)
		}
		ras.ClosePath()
	}
	// DISABLE AUTO-PTH CLR TO ENABLE StmBnd
	// x.PthClr()
}

func (x *Stk) RndrDraw(width float32, c clr.Clr, img *image.RGBA) {
	// RndrDraw is mainly for testing
	ras := &vector.Rasterizer{}
	ras.Reset(img.Rect.Dx(), img.Rect.Dy())
	x.Rndr(width, ras)
	ras.Draw(img, img.Rect, c.Uniform(), image.Point{})
	// for _, pxlFn := range pxlFns { // DEBUGGING
	// 	pxlFn()
	// }
}
func (x *Stk) Draw(c clr.Clr, r image.Rectangle, img *image.RGBA, ras *vector.Rasterizer) {
	ras.Draw(img, r, c.Uniform(), image.Point{})
	// for _, pxlFn := range pxlFns { // DEBUGGING
	// 	pxlFn()
	// }
}

func NewStkSeg(stkHlf float32, p1, p2 Pnt) (r StkSeg) {
	var angl float32
	angl, r.Dir = p2.AnglTau(p1)
	anglStk := stkAngl(angl, r.Dir)
	c, s := float32(math.Cos(float64(anglStk))), float32(math.Sin(float64(anglStk)))
	j, k := Pnt{Y: stkHlf}, Pnt{Y: -stkHlf}
	r.P1 = NewStkPnt(c, s, p1, j, k)
	r.P2 = NewStkPnt(c, s, p2, j, k)
	return r
}
func NewStkPnt(c, s float32, cntr, j, k Pnt) (r StkPnt) { // rotate point at origin then move
	// rot
	r.J.X = c*j.X + -s*j.Y
	r.J.Y = s*j.X + c*j.Y
	r.K.X = c*k.X + -s*k.Y
	r.K.Y = s*k.X + c*k.Y
	// mov
	r.J.X += cntr.X
	r.J.Y += cntr.Y
	r.K.X += cntr.X
	r.K.Y += cntr.Y
	// center
	r.Cntr = cntr
	return r
}
func stkAngl(angl float32, dir Pnt) (stkAngl float32) {
	switch {
	case dir.X == 1: // HRZ: RHT
		// no-op; stkAngl = 0
	case dir.X == -1: // HRZ: LFT
		stkAngl = TauHlf
	case dir.Y == -1: // VRT: TOP
		// stkAngl = TauQtr
		stkAngl = Tau3Qtr
	case dir.Y == 1: // VRT: BTM
		// stkAngl = Tau3Qtr
		stkAngl = TauQtr
	case dir.X > 0 && dir.Y < 0: // "QUAD I" (upper right)
		stkAngl = TauQtr * dir.Y // (original)
	case dir.X < 0 && dir.Y < 0: // "QUAD II" (upper left)
		stkAngl = Tau3Qtr + (TauQtr * dir.X) // +TauHlf for j,k swp
	case dir.X < 0 && dir.Y > 0: // "QUAD III" (lower left)
		stkAngl = TauHlf + (TauQtr * -dir.Y)
	case dir.X > 0 && dir.Y > 0: // "QUAD IIII" (lower right)
		stkAngl = Tau3Qtr + (TauQtr * -dir.X) + TauHlf // +TauHlf for j,k swp
	}
	// sys.Log("Vis.stkAngl", "angl", angl, "stkAngl", stkAngl, "dir", dir.Desc())
	return float32(math.Remainder(float64(stkAngl), Tau64))
}
