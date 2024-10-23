package plt

import (
	"image"
	"sys"
	"sys/ana/vis"
)

func (x *Hrz) Measure() vis.Siz { // Plt interface
	pltSizs := PltMeasurePll(*x.Plts...)
	var pos vis.Pos
	for n, pltSiz := range pltSizs { // horizontal stack layout
		x.siz.MaxHeight(pltSiz.Height)
		x.siz.Width += pltSiz.Width
		(*x.Plts)[n].Bse().pos = pos
		pos.X += pltSiz.Width
	}
	return x.siz
}
func (x *Hrz) Rndr() {
	for _, plt := range *x.Plts {
		plt.Bse().pos.X += x.pos.X
		plt.Bse().pos.Y += x.pos.Y
	}
	PltRndrPll(*x.Plts...)
}
func (x *Hrz) Draw(img *image.RGBA) { // Plt interface
	if len(*x.Plts) != 0 {
		acts := make([]sys.Act, len(*x.Plts))
		for n, plt := range *x.Plts {
			acts[n] = &PltDrawSeg{Plt: plt, Img: img}
		}
		sys.Run().Pll(acts...)
	}
}
