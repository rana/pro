package app

import (
	"fmt"
	"image"
	"unicode"
	"unicode/utf8"
	"unsafe"
)

type (
	ReadBox struct {
		BorderBse
		Style       Style
		buf         []byte
		lns         []Ln
		bufSize     image.Point
		pos         image.Point
		viewSize    image.Point
		sizeHint    image.Point
		minSizeHint image.Point
		title       *PrtTitle
		border      *PrtBorder
	}
	Ln struct {
		Text  string
		Width int // visible rune width
	}
)

// TODO: ProBox
//  -> Table
// 	-> Taber (support Multiple Table, Histogram etc)
// 	-> ProgressBar

// TODO: Prolang
// Instr: Calculate summary statistics
// Dsk: store instrument ticks
// Dsk: store instrument summary locally
// Dsk: store strat perf in cloud? or locally?

// TODO: Prolang
//	-> Trds Fld to Tbl column mapping
//	-> Trds .showTbl()
//	-> Trds Whr<prop>
//	-> Trds Srt<prop>(Asc|Dsc)

func NewReadBox(text ...string) (r *ReadBox) {
	r = &ReadBox{}
	if len(text) != 0 {
		r.SetText(text[0])
	}
	return r
}

func (x *ReadBox) SetTitle(text string) {
	x.BorderBse.SetTitle(text)
	x.ClearSizeHints()
}
func (x *ReadBox) SetBorder(has bool) {
	x.BorderBse.SetBorder(has)
	x.ClearSizeHints()
}
func (x *ReadBox) SetTextf(format string, args ...interface{}) {
	x.SetText(fmt.Sprintf(format, args...))
}
func (x *ReadBox) SetText(text string) {
	x.buf = append(x.buf, text...)
	x.CalcLines()
	x.pos.X = 0
	x.pos.Y = x.bufSize.Y // move cursor to end
}
func (x *ReadBox) Clear() {
	x.buf = nil
	x.lns = nil
	x.bufSize = image.ZP
	x.pos = image.ZP
}
func (x *ReadBox) Text() string {
	return *(*string)(unsafe.Pointer(&x.buf))
}
func (x *ReadBox) ClearSizeHints() {
	x.sizeHint = image.ZP
	x.minSizeHint = image.ZP
}
func (x *ReadBox) SizeHint() image.Point {
	if x.sizeHint.Eq(image.ZP) {
		x.Measure()
	}
	return x.sizeHint
}
func (x *ReadBox) MinSizeHint() image.Point {
	if x.minSizeHint.Eq(image.ZP) {
		x.Measure()
	}
	return x.minSizeHint
}

func (x *ReadBox) CalcLines() {
	x.bufSize = image.ZP
	x.lns = nil
	if len(x.buf) != 0 {
		var chSize, lnIdx, lnLim, lnWidth int
		var ch rune
		for lnLim < len(x.buf) {
			ch, chSize = utf8.DecodeRune(x.buf[lnLim:])
			if ch == '\n' {
				x.lns = append(x.lns, Ln{
					Text:  string(x.buf[lnIdx:lnLim]),
					Width: lnWidth,
				})
				if lnWidth > x.bufSize.X {
					x.bufSize.X = lnWidth
				}
				lnIdx = lnLim + chSize
				lnWidth = 0
			} else if unicode.IsGraphic(ch) { // count visible runes
				lnWidth++
			}
			lnLim += chSize
		}
		if ch != '\n' { // last line
			x.lns = append(x.lns, Ln{
				Text:  string(x.buf[lnIdx:lnLim]),
				Width: lnWidth,
			})
		}
		x.bufSize.Y = len(x.lns)
	}
}
func (x *ReadBox) Measure() {
	x.sizeHint = x.bufSize
	if x.title != nil { // title min length
		if x.title.size.X > x.minSizeHint.X { // ensure full title visible
			x.minSizeHint.X += x.title.size.X
		}
		x.sizeHint.Y += x.title.size.Y
		x.minSizeHint.Y += x.title.size.Y
	}
	if x.border != nil {
		x.sizeHint.X += 2 // left & right
		x.minSizeHint.X += 2
		if x.title != nil {
			x.sizeHint.Y += 1 // bottom
			x.minSizeHint.Y += 1
		} else {
			x.sizeHint.Y += 2 // bottom & top
			x.minSizeHint.Y += 2
		}
	}
}
func (x *ReadBox) Draw(p *Painter) {
	x.BorderBse.DrawBorder(p, func(p *Painter, viewSize image.Point) {
		x.viewSize = viewSize // for PgUp/PgDn
		p.WithStyle(x.Style, func(p *Painter) {
			if len(x.lns) != 0 { // text
				var lnIdxStart int
				if len(x.lns) > viewSize.Y {
					lnIdxStart = x.pos.Y
					if lnIdxStart > len(x.lns)-viewSize.Y {
						lnIdxStart = len(x.lns) - viewSize.Y
					}
				}
				var lnCnt, y int
				for lnIdx := lnIdxStart; lnIdx < len(x.lns) && lnCnt < viewSize.Y; lnIdx++ {
					p.FillRect(0, y, viewSize.X, 1)
					if lnIdx > -1 && lnIdx < len(x.lns) {
						p.DrawText(0, y, x.lns[lnIdx].Text)
					}
					y++
					lnCnt++
				}
			}
			if gapY := viewSize.Y - len(x.lns); gapY > 0 { // gap
				p.FillRect(0, len(x.lns), viewSize.X, gapY)
			}
			// if _ui.Root.focused==x { // cursor
			// 	p.DrawCursor(off.X+x.pos.X, off.Y+x.pos.Y)
			// }
		})
	})
}

func (x *ReadBox) Focus(prev ...bool) bool {
	x.ApplyStyle(true)
	return true
}
func (x *ReadBox) Unfocus() { x.ApplyStyle(false) }
func (x *ReadBox) ApplyStyle(isFocused ...bool) {
	focused := len(isFocused) != 0 && isFocused[0]
	x.BorderBse.ApplyStyle(focused)
	if focused {
		x.Style = _ui.Theme().Style(TextFocusedKey)
	} else {
		x.Style = _ui.Theme().Style(TextKey)
	}
}
func (x *ReadBox) OnKeyEvent(e KeyEvent) {
	if e.Modifiers == ModCtrl {
		switch e.Key {
		case KeyHome:
			x.BufStart()
		case KeyEnd:
			x.BufEnd()
		}
	} else {
		switch e.Key {
		case KeyDown:
			x.LnNext()
		case KeyUp:
			x.LnPrev()
		case KeyPgDn:
			x.PgNext()
		case KeyPgUp:
			x.PgPrev()
		}
	}
}
func (x *ReadBox) LnPrev() {
	if x.pos.Y > 0 {
		x.pos.Y--
	}
}
func (x *ReadBox) LnNext() {
	if x.pos.Y < len(x.lns)-1 {
		x.pos.Y++
	}
}
func (x *ReadBox) PgPrev() {
	if x.pos.Y-x.viewSize.Y+1 <= 0 { // add one to view previous ln
		x.pos.Y = 0
	} else {
		x.pos.Y = x.pos.Y - x.viewSize.Y + 1
	}
}
func (x *ReadBox) PgNext() {
	if x.pos.Y+x.viewSize.Y-1 >= len(x.lns) { // sub one the view next ln
		x.pos.Y = len(x.lns) - 1
	} else {
		x.pos.Y = x.pos.Y + x.viewSize.Y - 1
	}
	// logger.Printf("%v PgNext pos:%v", x, x.pos.Y)
}
func (x *ReadBox) BufStart() { x.pos.Y = 0 }
func (x *ReadBox) BufEnd()   { x.pos.Y = len(x.lns) - 1 }
func (x *ReadBox) String() string {
	if x.title != nil {
		return fmt.Sprintf("%v:%p", x.title.txt, x)
	}
	return fmt.Sprintf("ReadBox:%p", x)
}
