package app

type (
	PrtBorder struct {
		Style Style
	}
)

func NewPrtBorder(isFocused ...bool) (r *PrtBorder) {
	r = &PrtBorder{}
	r.ApplyStyle(isFocused...)
	return r
}
func (x *PrtBorder) ApplyStyle(isFocused ...bool) {
	if len(isFocused) != 0 && isFocused[0] {
		x.Style = _ui.Theme().Style(PrtBorderFocusedKey)
	} else {
		x.Style = _ui.Theme().Style(PrtBorderKey)
	}
}
func (x *PrtBorder) Width() int  { return 2 }
func (x *PrtBorder) Height() int { return 2 }
