package plt

// import (
// 	"image"
// 	"sys"
// 	"sys/ana/vis"
// )

// func (x *Prcp) Measure() vis.Siz { // Plt interface
// 	// TODO: REMOVE?
// 	// // CREATE PLTSTM FOR EACH STMFBR
// 	// if len(x.Prcp.StmFbrss) != 0 {
// 	// 	for m, stmFbrs := range x.Prcp.StmFbrss {
// 	// 		hrz := NewHrz()
// 	// 		x.Plt(hrz)
// 	// 		for _, stmFbr := range *stmFbrs {
// 	// 			stm := NewStm()
// 	// 			hrz.Plt(stm)
// 	// 			stm.StmFbr(x.Penss[m], stmFbr)
// 	// 		}
// 	// 	}
// 	// }

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
// func (x *Prcp) Rndr() {
// 	for _, plt := range *x.Plts {
// 		plt.Bse().pos.X += x.pos.X
// 		plt.Bse().pos.Y += x.pos.Y
// 	}
// 	PltRndrPll(*x.Plts...)
// 	// for _, p := range *x.Plts {
// 	// 	p.Rndr()
// 	// }
// }
// func (x *Prcp) Draw(img *image.RGBA) { // Plt interface
// 	if len(*x.Plts) != 0 {
// 		acts := make([]sys.Act, len(*x.Plts))
// 		for n, plt := range *x.Plts {
// 			acts[n] = &PltDrawSeg{Plt: plt, Img: img}
// 		}
// 		sys.Run().Pll(acts...)
// 	}
// }
