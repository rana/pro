package vis

import (
	"fmt"
	"math"
)

type (
	// https://github.com/processing/p5.js/blob/master/src/math/p5.Vector.js
	Pnt struct {
		X float32
		Y float32
	}
	PntOp func(a, b Pnt) Pnt
)

func Add(p, q Pnt) (r Pnt) {
	r.X = p.X + q.X
	r.Y = p.Y + q.Y
	return r
}
func Sub(p, q Pnt) (r Pnt) {
	r.X = p.X - q.X
	r.Y = p.Y - q.Y
	return r
}

func (p Pnt) Add(q Pnt) (r Pnt) {
	r.X = p.X + q.X
	r.Y = p.Y + q.Y
	return r
}
func (p Pnt) Sub(q Pnt) (r Pnt) {
	r.X = p.X - q.X
	r.Y = p.Y - q.Y
	return r
}
func (p Pnt) Mul(k float32) (r Pnt) {
	r.X = p.X * k
	r.Y = p.Y * k
	return r
}
func (p Pnt) Div(k float32) (r Pnt) {
	if k == 0 {
		return r
	}
	r.X = p.X / k
	r.Y = p.Y / k
	return r
}
func (p Pnt) Abs() Pnt {
	if p.X < 0 {
		p.X = -p.X
	}
	if p.Y < 0 {
		p.Y = -p.Y
	}
	return p
}

func (p Pnt) Mag() float32 { // magnitude
	return float32(math.Sqrt(float64(p.X*p.X + p.Y*p.Y)))
}
func (p Pnt) Nrm() (r Pnt) { // normalize
	m := p.Mag()
	if m > 0 {
		return p.Div(m)
	}
	return r
}
func (p Pnt) Dir() (r Pnt) { // direction (same as normalize)
	m := p.Mag()
	if m > 0 {
		return p.Div(m)
	}
	return r
}
func (p Pnt) Dist(q Pnt) float32 {
	dX := q.X - p.X
	dY := q.Y - p.Y
	return float32(math.Sqrt(float64(dX*dX + dY*dY)))
}
func (p Pnt) Perp() (r Pnt) { // perpendicular
	// http://johnblackburne.blogspot.com/2012/02/perp-dot-product.html
	// rotation through a quarter turn
	p.X, p.Y = -p.Y, p.X
	return p
}
func (p Pnt) PerpDot(q Pnt) float32 { // perpendicular dot prodcut
	// http://johnblackburne.blogspot.com/2012/02/perp-dot-product.html
	return -p.Y*q.X + p.X*q.Y
}
func (p Pnt) Dot(q Pnt) float32 { // dot product
	return p.X*q.X + p.Y*q.Y
}
func (p Pnt) Cross(q Pnt) float32 { // cross product
	return p.X*q.Y - p.Y*q.Y
}
func (p Pnt) Heading() float32 { // heading in radians
	return float32(math.Atan2(float64(p.Y), float64(p.X)))
}

// Angl returns the radian angle between two vectors.
func (p Pnt) Angl(q Pnt) float32 { // radian angle between p and q
	dot := p.X*q.X + p.Y*q.Y
	mag := float32(math.Sqrt(float64(p.X*p.X + p.Y*p.Y)))
	dotmagmag := dot / (mag * mag)
	// Mathematically speaking: the dotmagmag variable will be between -1 and 1
	// inclusive. Practically though it could be slightly outside this range due
	// to floating-point rounding issues. This can make Math.acos return NaN.
	//
	// Solution: we'll clamp the value to the -1,1 range
	return float32(math.Acos(math.Min(1, math.Max(-1, float64(dotmagmag)))))
}

// AnglTau returns the unit circle radian angle between two vectors.
func (p Pnt) AnglTau(q Pnt) (angl float32, dir Pnt) {
	// direction (normalize)
	// dirAB := x.b.Sub(x.a).Dir()
	dir.X = p.X - q.X
	dir.Y = p.Y - q.Y
	mag := float32(math.Sqrt(float64(dir.X*dir.X + dir.Y*dir.Y)))
	if mag == 0 {
		return 0, dir
	}
	dir.X /= mag
	dir.Y /= mag
	switch {
	case dir.X == -1: // HRZ: LFT
		return TauHlf, dir
	case dir.X == 1: // HRZ: RHT
		return 0, dir
	case dir.Y == -1: // VRT: TOP
		return TauQtr, dir
	case dir.Y == 1: // VRT: BTM
		return Tau3Qtr, dir
	case dir.X > 0 && dir.Y > 0: // "QUAD I" (upper right)
		return TauQtr * dir.Y, dir
	case dir.X < 0 && dir.Y > 0: // "QUAD II" (upper left)
		return TauQtr + (TauQtr * dir.Y), dir
	case dir.X < 0 && dir.Y < 0: // "QUAD III" (lower left)
		return TauHlf + (TauQtr * -dir.Y), dir
	case dir.X > 0 && dir.Y < 0: // "QUAD IIII" (lower right)
		return Tau3Qtr + (TauQtr * -dir.Y), dir
	}
	return 0, dir
}

// RotMov returns a rotated and translated point. The rotation is around the origin.
func (p Pnt) RotMov(angl float32, mov Pnt) (r1 Pnt) { // rotate point at origin then move
	// rot
	c, s := float32(math.Cos(float64(angl))), float32(math.Sin(float64(angl)))
	r1.X = c*p.X + -s*p.Y
	r1.Y = s*p.X + c*p.Y
	// mov
	r1.X += mov.X
	r1.Y += mov.Y
	return r1
}

// RotCCW90 returns the vector p rotated counter-clockwise by 90 degrees.
//
// Note that the Y-axis grows downwards, so {1, 0}. RotCCW90 is {0, -1}.
func (p Pnt) RotCCW90() Pnt {
	p.X = -p.X
	return p
}
func (p Pnt) RotCCWQtr() Pnt {
	p.X = -p.X
	return p
}

// RotCW90 returns the vector p rotated clockwise by 90 degrees.
//
// Note that the Y-axis grows downwards, so {1, 0}. RotCW90 is {0, 1}.
func (p Pnt) RotCW90() Pnt {
	p.Y = -p.Y
	return p
}
func (p Pnt) RotCWQtr() Pnt {
	p.Y = -p.Y
	return p
}

// Rot180 returns the vector p rotated by 180 degrees.
func (p Pnt) Rot180() Pnt {
	p.X = -p.X
	p.Y = -p.Y
	return p
}
func (p Pnt) RotHlf() Pnt {
	p.X = -p.X
	p.Y = -p.Y
	return p
}

func (p Pnt) Rot(angl float32) (r Pnt) { // rotate point at origin
	c, s := float32(math.Cos(float64(angl))), float32(math.Sin(float64(angl)))
	// m := Mtrx{XX: c, YX: s, XY: -s, YY: c}
	// r.X = m.XX*p.X + m.XY*p.Y // rotate
	// r.Y = m.YX*p.X + m.YY*p.Y
	r.X = c*p.X + -s*p.Y // rotate
	r.Y = s*p.X + c*p.Y
	return r
}
func (p Pnt) RotAt(angl float32, cntr Pnt) (r Pnt) { // rotate point at origin
	c, s := float32(math.Cos(float64(angl))), float32(math.Sin(float64(angl)))
	m := Mtrx{XX: c, YX: s, XY: -s, YY: c}
	// m.X0 = cntr.X*m.XX + cntr.Y*m.XY
	// m.Y0 = cntr.X*m.YX + cntr.Y*m.YY
	// m.X0 = cntr.X
	// m.Y0 = cntr.Y
	return m.Pnt(p)
}

// RotAsVec returns the vector p rotated by the specified radian angle.
func (p Pnt) RotAsVec(radianAngle float32) Pnt {
	// https://github.com/processing/p5.js/blob/master/src/math/p5.Vector.js
	// var newHeading = this.heading() + a;
	// if (this.p5) newHeading = this.p5._toRadians(newHeading);
	// var mag = this.mag();
	// this.x = Math.cos(newHeading) * mag;
	// this.y = Math.sin(newHeading) * mag;
	// return this;
	newHeading := p.Heading() + radianAngle
	mag := p.Mag()
	p.X = float32(math.Cos(float64(newHeading))) * mag
	p.Y = float32(math.Sin(float64(newHeading))) * mag
	return p
}

func (p Pnt) Swp() Pnt {
	p.X, p.Y = p.Y, p.X
	return p
}

func (p Pnt) Desc() string {
	switch {
	case p.Y == 0 && p.X == 0:
		return fmt.Sprintf("%v EMPTY", p)
	case p.X == 1:
		return fmt.Sprintf("%v HRZ: RHT", p)
	case p.X == -1:
		return fmt.Sprintf("%v HRZ: LFT", p)
	case p.Y == -1:
		return fmt.Sprintf("%v VRT: TOP", p)
	case p.Y == 1:
		return fmt.Sprintf("%v VRT: BTM", p)
	case p.X > 0 && p.Y < 0:
		return fmt.Sprintf("%v QUAD I", p)
	case p.X < 0 && p.Y < 0:
		return fmt.Sprintf("%v QUAD II", p)
	case p.X < 0 && p.Y > 0:
		return fmt.Sprintf("%v QUAD III", p)
	case p.X > 0 && p.Y > 0:
		return fmt.Sprintf("%v QUAD IIII", p)
	}
	panic("no")
}

// switch {
// case dirBC.Y == 0: // HRZ
// case dirBC.X == 0: // VRT
// case dirBC.X > 0 && dirBC.Y < 0: // "QUAD I" (upper right)
// case dirBC.X < 0 && dirBC.Y < 0: // "QUAD II" (upper left)
// case dirBC.X < 0 && dirBC.Y > 0: // "QUAD III" (lower left)
// case dirBC.X > 0 && dirBC.Y > 0: // "QUAD IIII" (lower right)
// }

// switch {
// 	case dir.X == -1: // HRZ: LFT
// 	case dir.X == 1: // HRZ: RHT
// 	case dir.Y == -1: // VRT: TOP
// 	case dir.Y == 1: // VRT: BTM
// 	case dir.X > 0 && dir.Y < 0: // "QUAD I" (upper right)
// 	case dir.X < 0 && dir.Y < 0: // "QUAD II" (upper left)
// 	case dir.X < 0 && dir.Y > 0: // "QUAD III" (lower left)
// 	case dir.X > 0 && dir.Y > 0: // "QUAD IIII" (lower right)
// 	}
