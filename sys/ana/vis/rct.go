package vis

import (
	"image"
)

type (
	Rct struct {
		Lft uint32
		Rht uint32
		Btm uint32
		Top uint32
	}
	Rct2 struct {
		Pos
		Siz
	}
)

func (x *Rct) Set(lft, top, width, height uint32) {
	x.Lft = lft
	x.Rht = lft + width
	x.Top = top
	x.Btm = top + height
}

func (x *Rct) Width() uint32  { return x.Rht - x.Lft }
func (x *Rct) Height() uint32 { return x.Btm - x.Top }

func Rect(p Pos, s Siz) (r image.Rectangle) {
	r.Min.X = int(p.X)
	r.Min.Y = int(p.Y)
	r.Max.X = r.Min.X + int(s.Width)
	r.Max.Y = r.Min.Y + int(s.Height)
	return r
}
