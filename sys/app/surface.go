package app

import (
	"image"

	"github.com/gdamore/tcell"
)

type (
	// Surface defines a surface that can be painted on.
	Surface interface {
		SetCell(x, y int, ch rune, s Style)
		SetCursor(x, y int)
		HideCursor()
		Begin()
		End()
		Size() image.Point
	}
	tcellSurface struct {
		screen tcell.Screen
	}
)

func (s *tcellSurface) SetCell(x, y int, ch rune, style Style) {
	st := tcell.StyleDefault.Normal().
		Foreground(convertColor(style.Fg, false)).
		Background(convertColor(style.Bg, false)).
		Reverse(style.Reverse == DecorationOn).
		Bold(style.Bold == DecorationOn).
		Underline(style.Underline == DecorationOn)

	s.screen.SetContent(x, y, ch, nil, st)
}
func (s *tcellSurface) SetCursor(x, y int) {
	s.screen.ShowCursor(x, y)
}
func (s *tcellSurface) HideCursor() {
	s.screen.HideCursor()
}
func (s *tcellSurface) Begin() {
	s.screen.Clear()
}
func (s *tcellSurface) End() {
	s.screen.Show()
}
func (s *tcellSurface) Size() image.Point {
	w, h := s.screen.Size()
	return image.Point{w, h}
}
