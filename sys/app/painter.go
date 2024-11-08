package app

import (
	"image"
)

type (
	// Painter provides operations to paint on a surface.
	Painter struct {
		theme      *Theme
		surface    Surface       // Surface to paint on.
		style      Style         // Current brush.
		transforms []image.Point // Transform stack
		mask       image.Rectangle
	}
)

// NewPainter returns a new instance of Painter.
func NewPainter(s Surface, theme *Theme) *Painter {
	// logger.Printf("NewPainter: %v", theme == nil)
	// fmt.Println("NewPainter: v", theme == nil)
	return &Painter{
		theme:   theme,
		surface: s,
		style:   theme.Style("normal"),
		mask: image.Rectangle{
			Min: image.ZP,
			Max: s.Size(),
		},
	}
}

// Translate pushes a new translation transform to the stack.
func (p *Painter) Translate(x, y int) {
	p.transforms = append(p.transforms, image.Point{x, y})
}

// Restore pops the latest transform from the stack.
func (p *Painter) Restore() {
	if len(p.transforms) > 0 {
		p.transforms = p.transforms[:len(p.transforms)-1]
	}
}

// Begin prepares the surface for painting.
func (p *Painter) Begin() {
	p.surface.Begin()
}

// End finalizes any painting that has been made.
func (p *Painter) End() {
	p.surface.End()
}

// Repaint clears the surface, draws the scene and flushes it.
func (p *Painter) Repaint(w Ctrl) {
	p.mask = image.Rectangle{
		Min: image.ZP,
		Max: p.surface.Size(),
	}
	p.surface.HideCursor()
	p.Begin()
	w.Resize(p.surface.Size())
	w.Draw(p)
	p.End()
}

// DrawCursor draws the cursor at the given position.
func (p *Painter) DrawCursor(x, y int) {
	wp := p.mapLocalToWorld(image.Point{x, y})
	p.surface.SetCursor(wp.X, wp.Y)
}

// DrawRune paints a rune at the given coordinate.
func (p *Painter) DrawRune(x, y int, r rune) {
	wp := p.mapLocalToWorld(image.Point{x, y})
	if (p.mask.Min.X <= wp.X) && (wp.X < p.mask.Max.X) && (p.mask.Min.Y <= wp.Y) && (wp.Y < p.mask.Max.Y) {
		p.surface.SetCell(wp.X, wp.Y, r, p.style)
	}
}

// DrawText paints a string starting at the given coordinate.
func (p *Painter) DrawText(x, y int, text string) {
	for _, r := range text {
		p.DrawRune(x, y, r)
		// x += runeWidth(r)
		x++
	}
}

// DrawHorizontalLine paints a horizontal line using box characters.
func (p *Painter) DrawHorizontalLine(x1, x2, y int) {
	for x := x1; x < x2; x++ {
		p.DrawRune(x, y, '─')
	}
}

// DrawVerticalLine paints a vertical line using box characters.
func (p *Painter) DrawVerticalLine(x, y1, y2 int) {
	for y := y1; y < y2; y++ {
		p.DrawRune(x, y, '│')
	}
}

// DrawRect paints a rectangle using box characters.
func (p *Painter) DrawRect(x, y, w, h int) {
	for j := 0; j < h; j++ {
		for i := 0; i < w; i++ {
			m := i + x
			n := j + y

			switch {
			case i == 0 && j == 0:
				p.DrawRune(m, n, '┌')
			case i == w-1 && j == 0:
				p.DrawRune(m, n, '┐')
			case i == 0 && j == h-1:
				p.DrawRune(m, n, '└')
			case i == w-1 && j == h-1:
				p.DrawRune(m, n, '┘')
			case i == 0 || i == w-1:
				p.DrawRune(m, n, '│')
			case j == 0 || j == h-1:
				p.DrawRune(m, n, '─')
			}
		}
	}
}

// FillRect clears a rectangular area with whitespace.
func (p *Painter) FillRect(x, y, w, h int) {
	for j := 0; j < h; j++ {
		for i := 0; i < w; i++ {
			p.DrawRune(i+x, j+y, ' ')
		}
	}
}

// SetStyle sets the style used when painting.
func (p *Painter) SetStyle(s Style) {
	p.style = s
}

// WithStyleKey executes the provided function with the named Style applied on top of the current one.
func (p *Painter) WithStyleKey(key string, fn func(*Painter)) {
	p.WithStyle(p.theme.Style(key), fn)
}

// WithStyle executes the provided function with the named Style applied on top of the current one.
func (p *Painter) WithStyle(s Style, fn func(*Painter)) {
	prev := p.style
	new := prev.Merge(s)
	p.SetStyle(new)
	fn(p)
	p.SetStyle(prev)
}

// WithMask masks a painter to restrict painting within the given rectangle.
func (p *Painter) WithMask(r image.Rectangle, fn func(*Painter)) {
	tmp := p.mask
	defer func() { p.mask = tmp }()

	p.mask = p.mask.Intersect(image.Rectangle{
		Min: p.mapLocalToWorld(r.Min),
		Max: p.mapLocalToWorld(r.Max),
	})

	fn(p)
}

func (p *Painter) mapLocalToWorld(point image.Point) image.Point {
	var offset image.Point
	for _, s := range p.transforms {
		offset = offset.Add(s)
	}
	return point.Add(offset)
}
