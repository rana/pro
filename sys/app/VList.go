package app

import (
	"image"
	"unicode/utf8"
)

type (
	VList struct {
		List
	}
)

func NewVList() (r *VList) {
	r = &VList{}
	r.Padding = NewLength(1, 1, 0, 0)
	r.measure = r.Measure
	return r
}

func (x *VList) Measure() {
	for n, item := range x.items {
		item.Width = utf8.RuneCountInString(item.Itm.GetName())
		if item.Width > x.sizeHint.X {
			x.sizeHint.X = item.Width
		}
		if n == 0 || item.Width < x.minSizeHint.X {
			x.minSizeHint.X = item.Width
		}
	}
	x.sizeHint.Y += len(x.items)

	borderSize, borderMinSize := x.BorderBse.MeasureBorder()
	x.sizeHint = x.sizeHint.Add(borderSize)
	x.minSizeHint = x.minSizeHint.Add(borderSize)

	if borderMinSize.X > x.minSizeHint.X {
		x.minSizeHint.X = borderMinSize.X
	}
	if borderMinSize.Y > x.minSizeHint.Y {
		x.minSizeHint.Y = borderMinSize.Y
	}

	x.sizeHint.X += x.Padding.Width()
	x.minSizeHint.X += x.Padding.Width()
}

func (x *VList) Draw(p *Painter) {
	x.BorderBse.DrawBorder(p, func(p *Painter, viewSize image.Point) {
		for y, item := range x.items { // items
			p.WithStyle(item.Style, func(p *Painter) {
				p.FillRect(0, y, viewSize.X, 1)
				p.DrawText(0+x.Padding.Left, y, item.Itm.GetName())
			})
		}
	})
}
func (x *VList) OnKeyEvent(e KeyEvent) {
	switch e.Key {
	case KeyDown:
		x.Next()
	case KeyUp:
		x.Prev()
	case KeyEnter:
		x.RaiseActivate()
	}
}
