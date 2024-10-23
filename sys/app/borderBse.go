package app

import (
	"image"
)

type (
	BorderBse struct {
		CtrlBse
		title  *PrtTitle
		border *PrtBorder
	}
)

func (x *BorderBse) SetTitle(text string) {
	if text != "" {
		x.title = NewPrtTitle(text)
	} else {
		x.title = nil
	}
}
func (x *BorderBse) SetBorder(has bool) {
	if has {
		x.border = NewPrtBorder()
	} else {
		x.border = nil
	}
}
func (x *BorderBse) MeasureBorder() (borderSize, minSize image.Point) {
	if x.title != nil { // title min length
		minSize = x.title.size // ensure full title visible
	}
	if x.border != nil {
		minSize.X += 2 // left & right
		if x.title != nil {
			minSize.Y += 1 // bottom with title
		} else {
			minSize.Y = 2 // bottom & top
		}
		borderSize.X = 2
		borderSize.Y = 2
	}
	return borderSize, minSize
}

func (x *BorderBse) DrawBorder(p *Painter, fn func(*Painter, image.Point)) {
	size := x.Size()
	viewSize := size
	var off image.Point
	if x.border != nil { // border
		p.WithStyle(x.border.Style, func(p *Painter) {
			p.DrawRect(0, 0, size.X, size.Y)
		})
		off.X++
		viewSize.X -= 2
		viewSize.Y -= 2
		p.Translate(1, 1)
		defer p.Restore()
	}
	if x.title != nil { // title
		p.WithStyle(x.title.Style, func(p *Painter) {
			var y int
			if x.border != nil {
				y = -1 // adjust for border translate
			}
			p.FillRect(0, y, x.title.size.X, x.title.size.Y)
			p.DrawText(0, y, x.title.txt)
		})
		if x.border == nil {
			viewSize.Y -= 1
		}
	}
	if x.border != nil || x.title != nil {
		off.Y++
	}
	fn(p, viewSize)
}

func (x *BorderBse) Focus(prev ...bool) bool {
	x.ApplyStyle(true)
	return true
}
func (x *BorderBse) Unfocus() { x.ApplyStyle(false) }
func (x *BorderBse) ApplyStyle(isFocused ...bool) {
	focused := len(isFocused) != 0 && isFocused[0]
	if x.title != nil {
		x.title.ApplyStyle(focused)
	}
	if x.border != nil {
		x.border.ApplyStyle(focused)
	}
}
