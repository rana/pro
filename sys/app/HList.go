package app

import (
	"image"
	"unicode/utf8"
)

type (
	HList struct {
		List
	}
)

func NewHList() (r *HList) {
	r = &HList{}
	r.Padding = NewLength(1, 1, 0, 0)
	r.measure = r.Measure
	return r
}

func (x *HList) Measure() {
	for _, item := range x.items {
		item.Width = utf8.RuneCountInString(item.Itm.GetName()) + x.Padding.Width()
		x.sizeHint.X += item.Width
	}
	x.minSizeHint.X = x.sizeHint.X
	x.sizeHint.Y = 1
	if x.title != nil { // title min length
		if x.title.size.X > x.minSizeHint.X { // ensure full title visible
			x.minSizeHint.X = x.title.size.X
		}
		x.sizeHint.Y += x.title.size.Y
	}
	if x.border != nil {
		x.sizeHint.X += 2 // left & right
		x.minSizeHint.X += 2
		if x.title != nil {
			x.sizeHint.Y += 1 // bottom
		} else {
			x.sizeHint.Y += 2 // bottom & top
		}
	}
	x.minSizeHint.Y = x.sizeHint.Y
}

func (x *HList) Draw(p *Painter) {
	size := x.Size()
	var offX, offY int
	if x.border != nil { // border
		p.WithStyle(x.border.Style, func(p *Painter) {
			p.DrawRect(0, 0, size.X, size.Y)
		})
		offX++
	}
	if x.title != nil { // title
		p.WithStyle(x.title.Style, func(p *Painter) {
			p.WithMask(image.Rect(0, 0, size.X-2, 1), func(p *Painter) {
				p.FillRect(offX, offY, x.title.size.X-offX, x.title.size.Y)
				p.DrawText(offX, offY, x.title.txt)
			})
		})
	}
	if x.border != nil || x.title != nil {
		offY++
	}
	for _, item := range x.items { // items
		curX := offX
		p.WithStyle(item.Style, func(p *Painter) {
			p.FillRect(curX, offY, item.Width, 1)
			p.DrawText(curX+x.Padding.Left, offY, item.Itm.GetName())
		})
		offX += item.Width
	}
}
func (x *HList) OnKeyEvent(e KeyEvent) {
	switch e.Key {
	case KeyRight:
		x.Next()
	case KeyLeft:
		x.Prev()
	case KeyEnter:
		x.RaiseActivate()
	}
}
