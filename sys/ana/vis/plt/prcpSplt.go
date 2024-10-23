package plt

// import (
// 	"image"
// 	"sys"
// 	"sys/ana/vis"
// )

// func (x *PrcpSplt) Measure() vis.Siz { // Plt interface
// 	pltSizs := PltMeasurePll(*x.Plts...)
// 	var pos vis.Pos
// 	for n, pltSiz := range pltSizs { // vertical stack layout
// 		x.siz.MaxWidth(pltSiz.Width)
// 		x.siz.Height += pltSiz.Height
// 		(*x.Plts)[n].Bse().pos = pos
// 		pos.Y += pltSiz.Height
// 	}
// 	return x.siz
// }

// func (x *PrcpSplt) Rndr() {
// 	for _, plt := range *x.Plts {
// 		plt.Bse().pos.X += x.pos.X
// 		plt.Bse().pos.Y += x.pos.Y
// 	}
// 	PltRndrPll(*x.Plts...)
// }
// func (x *PrcpSplt) Draw(img *image.RGBA) {
// 	if len(*x.Plts) != 0 {
// 		acts := make([]sys.Act, len(*x.Plts))
// 		for n, plt := range *x.Plts {
// 			acts[n] = &PltDrawSeg{Plt: plt, Img: img}
// 		}
// 		sys.Run().Pll(acts...)
// 	}
// }
