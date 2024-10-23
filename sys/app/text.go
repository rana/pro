package app

import (
	"fmt"
	"image"
	"unicode/utf8"
)

type (
	Text struct {
		CtrlBse
		txt     string
		runeCnt int
		Padding LengthX
		Style   Style
	}
)

func NewText(txt ...string) (r *Text) {
	r = &Text{}
	if len(txt) != 0 {
		r.SetTxt(txt[0])
	}
	return r
}
func (x *Text) SetTxt(txt string) {
	x.txt = txt
	x.runeCnt = utf8.RuneCountInString(x.txt)
}
func (x *Text) SetTxtf(format string, args ...interface{}) {
	x.SetTxt(fmt.Sprintf(format, args...))
}
func (x *Text) SizeHint() image.Point {
	return image.Point{x.runeCnt + x.Padding.Width(), 1}
}
func (x *Text) Draw(p *Painter) {
	size := x.Size()
	p.WithStyle(x.Style, func(p *Painter) {
		p.FillRect(0, 0, size.X, 1)
		p.DrawText(x.Padding.Left, 0, x.txt)
	})
}
