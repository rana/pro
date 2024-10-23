package vis

import (
	"image"
	"image/color"
)

func PxlPnt(p Pnt, c color.Color, img *image.RGBA) {
	img.Set(int(p.X), int(p.Y), c)
}
func PxlLnVrt(x1, y1, y2 uint32, c color.Color, img *image.RGBA) {
	if y1 != y2 {
		X := int(x1)
		min, max := minMaxI(int(y1), int(y2))
		for Y := min; Y <= max; Y++ {
			img.Set(X, Y, c)
		}
	}
}
func PxlLnHrz(x1, x2, y1 uint32, c color.Color, img *image.RGBA) {
	if x1 != x2 {
		Y := int(y1)
		min, max := minMaxI(int(x1), int(x2))
		for X := min; X <= max; X++ {
			img.Set(X, Y, c)
		}
	}
}

func PxlRct(r Rct, c color.Color, img *image.RGBA) {
	// rht, btm := r.Rht-1, r.Btm-1
	rht, btm := r.Rht, r.Btm
	PxlLnVrt(r.Lft, btm, r.Top, c, img)
	PxlLnVrt(rht, btm, r.Top, c, img)
	PxlLnHrz(r.Lft, rht, r.Top, c, img)
	PxlLnHrz(r.Lft, rht, btm, c, img)
}

// func (x *Vis) PthPxl(c color.Color) { // pxl at each pnt
// 	for _, p := range x.pth {
// 		x.PxlPnt(p, c)
// 	}
// }
