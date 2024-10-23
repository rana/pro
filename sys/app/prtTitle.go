package app

import (
	"image"
	"unicode/utf8"
)

type (
	PrtTitle struct {
		txt   string
		Style Style
		size  image.Point
	}
)

func NewPrtTitle(txt ...string) (r *PrtTitle) {
	r = &PrtTitle{}
	if len(txt) != 0 {
		r.SetTitle(txt[0])
	}
	r.ApplyStyle()
	return r
}
func (x *PrtTitle) SetTitle(txt string) {
	x.txt = txt
	x.size = image.Point{utf8.RuneCountInString(x.txt), 1}
}
func (x *PrtTitle) ApplyStyle(isFocused ...bool) {
	if len(isFocused) != 0 && isFocused[0] {
		x.Style = _ui.Theme().Style(PrtTitleFocusedKey)
	} else {
		x.Style = _ui.Theme().Style(PrtTitleKey)
	}
}
