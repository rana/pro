package plt

import (
	"fmt"
	"image"
	"image/draw"
	"image/gif"
	"image/png"
	"math/rand"
	"os"
	"path/filepath"
	"sys"
	"sys/ana/vis"
	"sys/bsc/bol"
	"sys/bsc/flt"
	"sys/bsc/unt"
	"sys/err"
	"time"
)

const (
	AxisTrnc = unt.Unt(4)
	LenMin   = float32(10)
)

var (
	Siz = vis.Siz{Width: 2560, Height: 1440}
)

type (
	PltBse struct {
		slf  Plt
		siz  vis.Siz
		pos  vis.Pos
		mrgn vis.LenXY
		brdr vis.Rct
		off  vis.Pnt // inner element rasterization start point
	}
	TmeFltPltBse struct {
		PltBse
		x     *TmeAxisX
		y     *FltAxisY
		sampl bol.Bol
	}
)

func NewPltBse(slf Plt) (r PltBse) {
	r.slf = slf
	r.siz = Siz // glbl Siz
	return r
}
func (x *PltBse) Bse() *PltBse { return x } // Plt interface
func (x *PltBse) ImgRndr() (r *image.RGBA) {
	x.slf.Measure() // measures siz
	x.slf.Rndr()
	r = image.NewRGBA(image.Rect(0, 0, int(x.siz.Width), int(x.siz.Height)))
	draw.Draw(r, r.Bounds(), BakClr.Uniform(), image.Point{}, draw.Src) // draw background
	x.slf.Draw(r)
	return r
}
func (x *PltBse) PngSav(path string) {
	file, er := os.Create(path)
	if er != nil {
		err.Panic(er)
	}
	defer file.Close()
	er = png.Encode(file, x.ImgRndr())
	if er != nil {
		err.Panic(er)
	}
}
func (x *PltBse) GifSav(path string) {
	file, er := os.Create(path)
	if er != nil {
		err.Panic(er)
	}
	defer file.Close()
	er = gif.Encode(file, x.ImgRndr(), &gif.Options{})
	if er != nil {
		err.Panic(er)
	}
}
func (x *PltBse) PngSavOpn(path string) {
	path = path + ".png"
	x.PngSav(path)
	sys.OpnImg(path)
}
func (x *PltBse) GifSavOpn(path string) {
	path = path + ".gif"
	x.GifSav(path)
	sys.OpnImg(path)
}
func (x *PltBse) Sho() Plt {
	t := time.Now()
	name := fmt.Sprintf("%v-%v-%v_%v-%v-%v_%v", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), rand.Int31())
	x.PngSavOpn(filepath.Join(os.TempDir(), name))
	return x.slf
}
func (x *PltBse) Siz(w, h unt.Unt) Plt {
	x.siz = vis.NewSiz(uint32(w), uint32(h))
	return x.slf
}
func (x *PltBse) Scl(v flt.Flt) Plt {
	x.siz = Siz.Scl(float32(v * Scl))
	return x.slf
}
func (x *PltBse) HrzScl(v flt.Flt) Plt {
	x.siz = Siz.HrzScl(float32(v * Scl))
	return x.slf
}
func (x *PltBse) VrtScl(v flt.Flt) Plt {
	x.siz = Siz.VrtScl(float32(v * Scl))
	return x.slf
}
