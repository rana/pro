package clr

import (
	"image"
	"image/color"
)

func (x *Clr) Uniform() *image.Uniform {
	return &image.Uniform{color.RGBA(*x)}
}
func (x *Clr) Color() color.Color {
	return color.RGBA(*x)
}
