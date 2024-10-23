package fnt

import (
	"fmt"
	"image/draw"
	"math"
	"sys/ana/vis/clr"

	xdraw "golang.org/x/image/draw"
	"golang.org/x/image/font"
	"golang.org/x/image/math/f64"
	"golang.org/x/image/math/fixed"
)

func (x *Fnt) Width(txt string) (r float32) {
	x.Mu.Lock() // lock due to face truetype.GlyphBuf usage
	r = float32(font.MeasureString(x.Face, txt) >> 6)
	x.Mu.Unlock()
	return r
}
func (x *Fnt) Height() (r float32) { return float32(x.Face.Metrics().Height >> 6) }
func (x *Fnt) HeightUint32() (r uint32) {
	return uint32(math.Ceil(float64((x.Face.Metrics().Height >> 6))))
}
func (x *Fnt) Drawf(x1, y1, offX, offY float32, clr clr.Clr, img draw.Image, format string, args ...interface{}) {
	x.Draw(x1, y1, offX, offY, clr, img, fmt.Sprintf(format, args...))
}
func (x *Fnt) Draw(x1, y1, offX, offY float32, clr clr.Clr, img draw.Image, txt string) {
	x.Mu.Lock()
	d := &font.Drawer{Dst: img, Src: clr.Uniform(), Face: x.Face}
	w := float32(d.MeasureString(txt) >> 6)
	h := float32(d.Face.Metrics().Height >> 6)
	d.Dot = fixp(x1-w+(offX*w), y1+(offY*h))
	// based on Drawer.DrawString() in golang.org/x/image/font/font.go
	prevC := rune(-1)
	for _, c := range txt {
		if prevC >= 0 {
			d.Dot.X += d.Face.Kern(prevC, c)
		}
		dr, mask, maskp, advance, ok := d.Face.Glyph(d.Dot, c)
		if !ok {
			continue
		}
		sr := dr.Sub(dr.Min)
		s2d := f64.Aff3{1, 0, float64(dr.Min.X), 0, 1, float64(dr.Min.Y)}
		transformer := xdraw.BiLinear
		transformer.Transform(d.Dst, s2d, d.Src, sr, xdraw.Over, &xdraw.Options{
			SrcMask:  mask,
			SrcMaskP: maskp,
		})
		d.Dot.X += advance
		prevC = c
	}
	x.Mu.Unlock()
}

func fixp(x, y float32) fixed.Point26_6 {
	return fixed.Point26_6{
		X: fixed.Int26_6(x * 64),
		Y: fixed.Int26_6(y * 64),
	}
}
