package vis

import (
	"math"
)

type Mtrx struct {
	XX, YX, XY, YY, X0, Y0 float32
}

func NewMtrx() (r Mtrx) { // identity
	r.XX = 1
	r.YY = 1
	return r
}
func NewMtrxPnt(p Pnt) (r Mtrx) { // identity + pnt
	r.XX = 1
	r.YY = 1
	r.X0 = p.X
	r.Y0 = p.Y
	return r
}
func NewMtrxAngl(angl float32) (r Mtrx) { // angle
	c, s := float32(math.Cos(float64(angl))), float32(math.Sin(float64(angl)))
	// fmt.Println("angle", angl, "c", c, "s", c)
	r.XX = c
	r.YX = s
	r.XY = -s
	r.YY = c
	return r
}
func NewMtrxAnglAt(angl float32, cntr Pnt) (r Mtrx) { // identity + angle
	c, s := float32(math.Cos(float64(angl))), float32(math.Sin(float64(angl)))
	r.XX = c
	r.YX = s
	r.XY = -s
	r.YY = c
	return r
}

func (m *Mtrx) Mul(b Mtrx) {
	xx := m.XX*b.XX + m.YX*b.XY
	yx := m.XX*b.YX + m.YX*b.YY
	xy := m.XY*b.XX + m.YY*b.XY
	yy := m.XY*b.YX + m.YY*b.YY
	x0 := m.X0*b.XX + m.Y0*b.XY + b.X0
	y0 := m.X0*b.YX + m.Y0*b.YY + b.Y0
	m.XX = xx
	m.YX = yx
	m.XY = xy
	m.YY = yy
	m.X0 = x0
	m.Y0 = y0
}
func (m *Mtrx) Pnt(p Pnt) (r Pnt) {
	r.X = m.XX*p.X + m.XY*p.Y + m.X0
	r.Y = m.YX*p.X + m.YY*p.Y + m.Y0
	return
}

func Mov(by, p1, p2 Pnt) (r1, r2 Pnt) {
	r1.X = p1.X + by.X
	r1.Y = p1.Y + by.Y
	r2.X = p2.X + by.X
	r2.Y = p2.Y + by.Y
	return r1, r2
}
func (x *Mtrx) Mov(p Pnt) {
	x.X0 += p.X
	x.Y0 += p.Y
}

func Intersect(a, c, dirAB, dirBC180 Pnt) (r Pnt) {
	// ray-intersect
	//x.bj = x.aj.Add(x.dirAB.Mul(cj.Sub(x.aj).PerpDot(dirBC180) / x.dirAB.PerpDot(dirBC180)))
	// perpendicular dot product
	// -p.Y*q.X + p.X*q.Y
	var lenAC Pnt
	lenAC.X = c.X - a.X
	lenAC.Y = c.Y - a.Y
	numScl := (-lenAC.Y*dirBC180.X + lenAC.X*dirBC180.Y) / (-dirAB.Y*dirBC180.X + dirAB.X*dirBC180.Y)
	r.X = a.X + dirAB.X*numScl
	r.Y = a.Y + dirAB.Y*numScl
	return r
}

func RotMov(angl float32, mov, p1 Pnt) (r1 Pnt) { // rotate point at origin then move
	// rot
	c, s := float32(math.Cos(float64(angl))), float32(math.Sin(float64(angl)))
	r1.X = c*p1.X + -s*p1.Y
	r1.Y = s*p1.X + c*p1.Y
	// mov
	r1.X += mov.X
	r1.Y += mov.Y
	return r1
}

func Rot(angl float32, p1, p2 Pnt) (r1, r2 Pnt) { // rotate point at origin
	c, s := float32(math.Cos(float64(angl))), float32(math.Sin(float64(angl)))
	m := Mtrx{XX: c, YX: s, XY: -s, YY: c}
	r1.X = m.XX*p1.X + m.XY*p1.Y // rotate p1
	r1.Y = m.YX*p1.X + m.YY*p1.Y
	r2.X = m.XX*p2.X + m.XY*p2.Y // rotate p2
	r2.Y = m.YX*p2.X + m.YY*p2.Y
	return r1, r2
}
func (m *Mtrx) Rot(angl float32) { // rotate at origin; ignore translation
	c, s := float32(math.Cos(float64(angl))), float32(math.Sin(float64(angl)))
	xx := m.XX*c + m.YX*-s
	yx := m.XX*s + m.YX*c
	xy := m.XY*c + m.YY*-s
	yy := m.XY*s + m.YY*c
	m.XX = xx
	m.YX = yx
	m.XY = xy
	m.YY = yy
}
func (m *Mtrx) RotAt(angl float32, cntr Pnt) { // rotate at origin; ignore translation
	c, s := float32(math.Cos(float64(angl))), float32(math.Sin(float64(angl)))
	xx := m.XX*c + m.YX*-s
	yx := m.XX*s + m.YX*c
	xy := m.XY*c + m.YY*-s
	yy := m.XY*s + m.YY*c

	m.XX = xx
	m.YX = yx
	m.XY = xy
	m.YY = yy
	m.X0 = cntr.X*xx + cntr.Y*xy
	m.Y0 = cntr.X*yx + cntr.Y*yy

	// a.X0*b.XX + a.Y0*b.XY + b.X0,
	// 	a.X0*b.YX + a.Y0*b.YY + b.Y0,

	// XX := c
	// YX := s
	// XY := -s
	// YY := c
	// xx := m.XX*XX + m.YX*XY
	// yx := m.XX*YX + m.YX*YY
	// xy := m.XY*XX + m.YY*XY
	// yy := m.XY*YX + m.YY*YY
	// m.XX = xx
	// m.YX = yx
	// m.XY = xy
	// m.YY = yy
}

// func (a Mtrx) Rotate(angle float32) Mtrx {
// 	return Rotate(angle).Multiply(a)
// }
func Rotate(angle float32) Mtrx {
	c := float32(math.Cos(float64(angle)))
	s := float32(math.Sin(float64(angle)))
	return Mtrx{
		c, s,
		-s, c,
		0, 0,
	}
}
func (a Mtrx) Rotate(angle float32) Mtrx {
	return Rotate(angle).Multiply(a)
}

func (a Mtrx) Multiply(b Mtrx) Mtrx {
	return Mtrx{
		a.XX*b.XX + a.YX*b.XY,
		a.XX*b.YX + a.YX*b.YY,
		a.XY*b.XX + a.YY*b.XY,
		a.XY*b.YX + a.YY*b.YY,
		a.X0*b.XX + a.Y0*b.XY + b.X0,
		a.X0*b.YX + a.Y0*b.YY + b.Y0,
	}
}

// func (a Mtrx) Scale(x, y float32) Mtrx {
// 	return Scale(x, y).Multiply(a)
// }
// func Scale(x, y float32) Mtrx {
// 	return Mtrx{
// 		x, 0,
// 		0, y,
// 		0, 0,
// 	}
// }

// func (a Mtrx) Shear(x, y float32) Mtrx {
// 	return Shear(x, y).Multiply(a)
// }
// func Shear(x, y float32) Mtrx {
// 	return Mtrx{
// 		1, y,
// 		x, 1,
// 		0, 0,
// 	}
// }

// func (a Mtrx) Translate(x, y float32) Mtrx {
// 	return Translate(x, y).Multiply(a)
// }
// func Translate(x, y float32) Mtrx {
// 	return Mtrx{
// 		1, 0,
// 		0, 1,
// 		x, y,
// 	}
// }

// func (a Mtrx) TransformVector(x, y float32) (tx, ty float32) {
// 	tx = a.XX*x + a.XY*y
// 	ty = a.YX*x + a.YY*y
// 	return
// }
// func (a Mtrx) TransformPoint(x, y float32) (tx, ty float32) {
// 	tx = a.XX*x + a.XY*y + a.X0
// 	ty = a.YX*x + a.YY*y + a.Y0
// 	return
// }
