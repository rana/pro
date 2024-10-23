package plt

import (
	"image"
	"sys"
	"sys/ana/vis"
)

func (x *Dpth) Measure() vis.Siz { // Plt interface
	pltSizs := PltMeasurePll(*x.Plts...)
	for _, pltSiz := range pltSizs { // vertical stack layout
		x.siz.MaxWidth(pltSiz.Width)
		x.siz.MaxHeight(pltSiz.Height)
	}
	return x.siz
}
func (x *Dpth) Rndr() {
	for _, plt := range *x.Plts {
		plt.Bse().pos.X += x.pos.X
		plt.Bse().pos.Y += x.pos.Y
	}
	PltRndrPll(*x.Plts...)
}
func (x *Dpth) Draw(img *image.RGBA) { // Plt interface
	if len(*x.Plts) != 0 {
		acts := make([]sys.Act, len(*x.Plts))
		for n, plt := range *x.Plts {
			acts[n] = &PltDrawSeg{Plt: plt, Img: img}
		}
		sys.Run().Pll(acts...)
	}
}
