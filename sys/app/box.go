package app

import (
	"fmt"
	"image"
	"math"
)

// Box is a layout for placing widgets either horizontally or vertically. If
// horizontally, all widgets will have the same height. If vertically, they
// will all have the same width.
type Box struct {
	BorderBse
	ctrls []Ctrl
	// Border    bool
	Title     string
	Alignment Alignment
	focused   Ctrl
}

func NewVBox(vs ...Ctrl) (r *Box) {
	r = &Box{}
	r.Alignment = Vertical
	r.Add(vs...)
	return r
}
func NewHBox(vs ...Ctrl) (r *Box) {
	r = &Box{}
	r.Alignment = Horizontal
	r.Add(vs...)
	return r
}

func (x *Box) Focus(prev ...bool) bool {
	// logger.Printf("Box.Focus: %v x.focused:%v", x, x.focused)
	if x.focused == nil {
		if len(prev) != 0 && prev[0] {
			x.FocusPrev()
		} else {
			x.FocusNext()
		}
	}
	if x.focused != nil {
		// logger.Printf("Box.Focus: %v x.focused.Focus():%v RET:TRUE", x, x.focused.Focus())
		x.focused.Focus()
		x.ApplyStyle(true)
		return true
	}
	// logger.Printf("Box.Focus: %v RET:FALSE", x)
	x.ApplyStyle(false)
	return false
}
func (x *Box) Unfocus() {
	if x.focused != nil {
		x.focused.Unfocus()
	}
	x.ApplyStyle(false)
}
func (x *Box) FocusPrev() bool {
	if len(x.ctrls) != 0 { // try focus child
		if x.focused != nil && x.focused.FocusPrev() {
			return true
		}
		if len(x.ctrls) == 1 {
			if x.ctrls[0].Focus(true) {
				x.focused = x.ctrls[0]
				if x.focused.FocusPrev() {
					return true
				}
			}
		} else {
			focusedIdx := x.FocusedIdx() // look for new peer focus
			if x == _ui.Root {
				for n := len(x.ctrls) - 1; n > 0; n-- {
					idx := (focusedIdx + n) % len(x.ctrls)
					if x.ctrls[idx].Focus(true) {
						if x.focused != nil {
							x.focused.Unfocus()
						}
						x.focused = x.ctrls[idx]
						return true
					}
				}
			} else {
				if focusedIdx < 0 {
					focusedIdx = len(x.ctrls)
				}
				for n := focusedIdx - 1; n > -1; n-- {
					if x.ctrls[n].Focus(true) {
						if x.focused != nil {
							x.focused.Unfocus()
						}
						x.focused = x.ctrls[n]
						return true
					}
				}
			}
		}
		if x.focused != nil && len(x.ctrls) != 1 { // clear to enable parent to call FocusNext cycling
			x.focused.Unfocus()
			x.focused = nil
		}
	}
	return false
}
func (x *Box) FocusNext() bool {
	if len(x.ctrls) != 0 { // try focus child
		if x.focused != nil {
			if x.focused != nil && x.focused.FocusNext() {
				return true
			}
		}
		if len(x.ctrls) == 1 {
			if x.ctrls[0].Focus() {
				x.focused = x.ctrls[0]
				if x.focused.FocusNext() {
					return true
				}
			}
		} else {
			focusedIdx := x.FocusedIdx() // look for new peer focus
			if x == _ui.Root {
				for n := 1; n < len(x.ctrls); n++ {
					idx := (focusedIdx + n) % len(x.ctrls)
					if x.ctrls[idx].Focus() {
						if x.focused != nil {
							x.focused.Unfocus()
						}
						x.focused = x.ctrls[idx]
						return true
					}
				}
			} else {
				for n := focusedIdx + 1; n < len(x.ctrls); n++ {
					if x.ctrls[n].Focus() {
						if x.focused != nil {
							x.focused.Unfocus()
						}
						x.focused = x.ctrls[n]
						return true
					}
				}
			}
		}
		if x.focused != nil && len(x.ctrls) != 1 { // clear to enable parent to call FocusNext cycling
			x.focused.Unfocus()
			x.focused = nil
		}
	}
	return false
}

func (x *Box) Add(vs ...Ctrl) {
	for _, v := range vs {
		v.Bse().Prnt = x
		x.ctrls = append(x.ctrls, v)
	}
}
func (x *Box) Insert(i int, c Ctrl) {
	if len(x.ctrls) < i || i < 0 {
		return
	}
	x.ctrls = append(x.ctrls[:i], append([]Ctrl{c}, x.ctrls[i:]...)...)
}
func (x *Box) Remove(i int) {
	if len(x.ctrls) <= i || i < 0 {
		return
	}
	if x.focused == x.ctrls[i] {
		x.FocusNext()
	}
	x.ctrls = append(x.ctrls[:i], x.ctrls[i+1:]...)
}

func (x *Box) Draw(p *Painter) {
	p.WithStyleKey("box", func(p *Painter) {
		x.BorderBse.DrawBorder(p, func(p *Painter, viewSize image.Point) {
			var off image.Point
			for _, child := range x.ctrls {
				switch x.Alignment {
				case Horizontal:
					p.Translate(off.X, 0)
				case Vertical:
					p.Translate(0, off.Y)
				}
				p.WithMask(image.Rectangle{
					Min: image.ZP,
					Max: child.Size(),
				}, func(p *Painter) {
					child.Draw(p)
				})
				p.Restore()
				off = off.Add(child.Size())
			}
		})

	})
}

// MinSizeHint returns the minimum size hint for the layout.
func (x *Box) MinSizeHint() image.Point {
	var minSize image.Point

	for _, child := range x.ctrls {
		size := child.MinSizeHint()
		if x.Alignment == Horizontal {
			minSize.X += size.X
			if size.Y > minSize.Y {
				minSize.Y = size.Y
			}
		} else {
			minSize.Y += size.Y
			if size.X > minSize.X {
				minSize.X = size.X
			}
		}
	}

	borderSize, _ := x.MeasureBorder()
	minSize.Add(borderSize)

	return minSize
}

// SizeHint returns the recommended size hint for the layout.
func (x *Box) SizeHint() image.Point {
	var sizeHint image.Point

	for _, child := range x.ctrls {
		size := child.SizeHint()
		if x.Alignment == Horizontal {
			sizeHint.X += size.X
			if size.Y > sizeHint.Y {
				sizeHint.Y = size.Y
			}
		} else {
			sizeHint.Y += size.Y
			if size.X > sizeHint.X {
				sizeHint.X = size.X
			}
		}
	}

	borderSize, _ := x.MeasureBorder()
	sizeHint.Add(borderSize)

	return sizeHint
}
func (x *Box) FocusedIdx() int {
	if x.focused != nil {
		for m := 0; m < len(x.ctrls); m++ {
			if x.ctrls[m] == x.focused {
				return m
			}
		}
	}
	return -1
}
func (x *Box) IsChildFocused(child Ctrl) bool { return x.focused == child }
func (x *Box) OnKeyEvent(e KeyEvent) {
	if x.focused != nil {
		x.focused.OnKeyEvent(e)
	}
}

// Resize recursively updates the size of the Box and all the widgets it
// contains. This is a potentially expensive operation and should be invoked
// with restraint.
//
// Resize is called by the layout engine and is not intended to be used by end
// users.
func (x *Box) Resize(size image.Point) {
	x.size = size
	borderSize, _ := x.MeasureBorder()
	inner := x.size.Sub(borderSize)
	x.layoutChildren(inner)
}

func (b *Box) layoutChildren(size image.Point) {
	space := doLayout(b.ctrls, dim(b.Alignment, size), b.Alignment)
	for i, s := range space {
		switch b.Alignment {
		case Horizontal:
			b.ctrls[i].Resize(image.Point{s, size.Y})
		case Vertical:
			b.ctrls[i].Resize(image.Point{size.X, s})
		}
	}
}

func doLayout(ws []Ctrl, space int, a Alignment) []int {
	sizes := make([]int, len(ws))

	if len(sizes) == 0 {
		return sizes
	}

	remaining := space

	// Distribute MinSizeHint
	for {
		var changed bool
		for i, sz := range sizes {
			if sz < dim(a, ws[i].MinSizeHint()) {
				sizes[i] = sz + 1
				remaining--
				if remaining <= 0 {
					goto Resize
				}
				changed = true
			}
		}
		if !changed {
			break
		}
	}

	// Distribute Minimum
	for {
		var changed bool
		for i, sz := range sizes {
			p := alignedSizePolicy(a, ws[i])
			if p == Minimum && sz < dim(a, ws[i].SizeHint()) {
				sizes[i] = sz + 1
				remaining--
				if remaining <= 0 {
					goto Resize
				}
				changed = true
			}
		}
		if !changed {
			break
		}
	}

	// Distribute Preferred
	for {
		var changed bool
		for i, sz := range sizes {
			p := alignedSizePolicy(a, ws[i])
			if (p == Preferred || p == Maximum) && sz < dim(a, ws[i].SizeHint()) {
				sizes[i] = sz + 1
				remaining--
				if remaining <= 0 {
					goto Resize
				}
				changed = true
			}
		}
		if !changed {
			break
		}
	}

	// Distribute Expanding
	for {
		var changed bool
		for i, sz := range sizes {
			p := alignedSizePolicy(a, ws[i])
			if p == Expanding {
				sizes[i] = sz + 1
				remaining--
				if remaining <= 0 {
					goto Resize
				}
				changed = true
			}
		}
		if !changed {
			break
		}
	}

	// Distribute remaining space
	for {
		min := math.MaxInt64
		for i, s := range sizes {
			p := alignedSizePolicy(a, ws[i])
			if (p == Preferred || p == Minimum) && s <= min {
				min = s
			}
		}
		var changed bool
		for i, sz := range sizes {
			if sz != min {
				continue
			}
			p := alignedSizePolicy(a, ws[i])
			if p == Preferred || p == Minimum {
				sizes[i] = sz + 1
				remaining--
				if remaining <= 0 {
					goto Resize
				}
				changed = true
			}
		}
		if !changed {
			break
		}
	}

Resize:

	return sizes
}

func dim(a Alignment, pt image.Point) int {
	if a == Horizontal {
		return pt.X
	}
	return pt.Y
}

func alignedSizePolicy(a Alignment, w Ctrl) SizePolicy {
	hpol, vpol := w.SizePolicy()
	if a == Horizontal {
		return hpol
	}
	return vpol
}

func (x *Box) String() string {
	if x.Title != "" {
		return fmt.Sprintf("%v:%p", x.Title, x)
	}
	return fmt.Sprintf("Box:%p", x)
}
