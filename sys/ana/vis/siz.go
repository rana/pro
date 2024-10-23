package vis

import "image"

type (
	Siz struct {
		Width  uint32
		Height uint32
	}
)

func NewSiz(w, h uint32) (r Siz) {
	if w <= LenMin {
		r.Width = LenMin
	} else {
		r.Width = w
	}
	if h <= LenMin {
		r.Height = LenMin
	} else {
		r.Height = h
	}
	return r
}

func (x Siz) Scl(pct float32) (r Siz) {
	r.Width = uint32(float32(x.Width) * pct)
	r.Height = uint32(float32(x.Height) * pct)
	if x.Width <= LenMin {
		x.Width = LenMin
	}
	if x.Height <= LenMin {
		x.Height = LenMin
	}
	return r
}
func (x Siz) HrzScl(pct float32) Siz {
	x.Width = uint32(float32(x.Width) * pct)
	if x.Width <= LenMin {
		x.Width = LenMin
	}
	return x
}
func (x Siz) VrtScl(pct float32) Siz {
	x.Height = uint32(float32(x.Height) * pct)
	if x.Height <= LenMin {
		x.Height = LenMin
	}
	return x
}
func (x *Siz) MaxWidth(w uint32) {
	if w > x.Width {
		x.Width = w
	}
}
func (x *Siz) MaxHeight(h uint32) {
	if h > x.Height {
		x.Height = h
	}
}
func (x *Siz) Reset() {
	x.Width = 0
	x.Height = 0
}
func (x Siz) Point() (r image.Point) {
	r.X = int(x.Width)
	r.Y = int(x.Height)
	return r
}
func (x Siz) Rectangle(p Pos) (r image.Rectangle) {
	r.Min.X = int(p.X)
	r.Min.Y = int(p.Y)
	r.Max.X = int(x.Width)
	r.Max.Y = int(x.Height)
	return r
}
