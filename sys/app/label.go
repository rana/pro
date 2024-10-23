package app

import (
	"image"
	"strings"
	"unicode/utf8"
	// wordwrap "github.com/mitchellh/go-wordwrap"
)

// Label is a widget to display read-only text.
type Label struct {
	CtrlBse

	text     string
	wordWrap bool

	// cache the result of SizeHint() (see #14)
	cacheSizeHint *image.Point

	styleName string
}

// NewLabel returns a new Label.
func NewLabel(text string) *Label {
	return &Label{
		text: text,
	}
}

// Resize changes the size of the Ctrl.
func (l *Label) Resize(size image.Point) {
	if l.Size() != size {
		l.cacheSizeHint = nil
	}
	l.CtrlBse.Resize(size)
}

// Draw draws the label.
func (l *Label) Draw(p *Painter) {
	lines := l.lines()

	style := "label"
	if l.styleName != "" {
		style += "." + l.styleName
	}

	p.WithStyleKey(style, func(p *Painter) {
		for i, line := range lines {
			p.DrawText(0, i, line)
		}
	})
}

// MinSizeHint returns the minimum size the widget is allowed to be.
func (l *Label) MinSizeHint() image.Point {
	return image.Point{1, 1}
}

// SizeHint returns the recommended size for the label.
func (l *Label) SizeHint() image.Point {
	if l.cacheSizeHint != nil {
		return *l.cacheSizeHint
	}
	var max int
	lines := l.lines()
	for _, line := range lines {
		if w := utf8.RuneCountInString(line); w > max {
			max = w
		}
	}
	sizeHint := image.Point{max, len(lines)}
	l.cacheSizeHint = &sizeHint
	return sizeHint
}

func (l *Label) lines() []string {
	txt := l.text
	// if l.wordWrap {
	// 	txt = wordwrap.WrapString(l.text, uint(l.Size().X))
	// }
	return strings.Split(txt, "\n")
}

// Text returns the text content of the label.
func (l *Label) Text() string {
	return l.text
}

// SetText sets the text content of the label.
func (l *Label) SetText(text string) {
	l.cacheSizeHint = nil
	l.text = text
}

// SetWordWrap sets whether text content should be wrapped.
func (l *Label) SetWordWrap(enabled bool) {
	l.wordWrap = enabled
}

// SetStyleName sets the identifier used for custom styling.
func (l *Label) SetStyleName(style string) {
	l.styleName = style
}
