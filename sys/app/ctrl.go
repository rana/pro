package app

import "image"

type (
	// Ctrl defines common operations on widgets.
	Ctrl interface {
		Bse() *CtrlBse
		Draw(p *Painter)
		MinSizeHint() image.Point
		Size() image.Point
		SizeHint() image.Point
		SizePolicy() (SizePolicy, SizePolicy)
		Resize(size image.Point)
		OnKeyEvent(ev KeyEvent)
		Focus(prev ...bool) bool
		Unfocus()
		FocusPrev() bool
		FocusNext() bool
		IsChildFocused(child Ctrl) bool
	}
	CtrlBse struct {
		Prnt        Ctrl
		size        image.Point
		sizePolicyX SizePolicy
		sizePolicyY SizePolicy
	}
)

func (x *CtrlBse) Bse() *CtrlBse                  { return x }
func (x *CtrlBse) Draw(p *Painter)                {}
func (x *CtrlBse) FocusPrev() bool                { return false }
func (x *CtrlBse) FocusNext() bool                { return false }
func (x *CtrlBse) Focus(prev ...bool) bool        { return false }
func (x *CtrlBse) Unfocus()                       {}
func (x *CtrlBse) IsChildFocused(child Ctrl) bool { return false }

// MinSizeHint returns the size below which the widget cannot shrink.
func (x *CtrlBse) MinSizeHint() image.Point { return image.Point{1, 1} }

// Size returns the current size of the widget.
func (x *CtrlBse) Size() image.Point { return x.size }

// SizeHint returns the size hint of the widget.
func (x *CtrlBse) SizeHint() image.Point { return image.ZP }

// SetSizePolicy sets the size policy for horizontal and vertical directions.
func (x *CtrlBse) SetSizePolicy(h, v SizePolicy) {
	x.sizePolicyX = h
	x.sizePolicyY = v
}
func (x *CtrlBse) SetSizePolicyX(h SizePolicy) { x.sizePolicyX = h }
func (x *CtrlBse) SetSizePolicyY(v SizePolicy) { x.sizePolicyY = v }

// SizePolicy returns the current size policy.
func (x *CtrlBse) SizePolicy() (SizePolicy, SizePolicy) { return x.sizePolicyX, x.sizePolicyY }

// Resize sets the size of the widget.
func (x *CtrlBse) Resize(size image.Point) { x.size = size }

// OnKeyEvent is an empty operation to fulfill the Ctrl interface.
func (x *CtrlBse) OnKeyEvent(ev KeyEvent) {}
