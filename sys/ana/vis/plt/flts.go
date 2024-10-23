package plt

// import (
// 	"sys"
// 	"sys/ana/vis"
// 	"sys/bsc/flt"
// 	"sys/bsc/flts"
// 	"sys/bsc/unt"

// 	"golang.org/x/image/colornames"
// )

// func NewFlts(v *flts.Flts) (r *Flts) {
// 	r = &Flts{}
// 	r.X = &UntAxisX{}
// 	r.Y = NewFltAxisY()
// 	r.AddFlts(v)
// 	return r
// }
// func (x *Flts) AddFlts(v *flts.Flts) {
// 	x.Flts = append(x.Flts, v)
// }
// func (x *Flts) Measure() vis.Size {
// 	x.X.Len = vis.ScreenWidth  //- x.crt.margin.Width()
// 	x.Y.Len = vis.ScreenHeight //- x.crt.margin.Height()
// 	if len(x.Flts) != 0 {
// 		// x-axis
// 		x.X.Min = 0
// 		x.X.Max = x.Flts[0].Cnt()
// 		for n := unt.One; n < unt.Unt(len(x.Flts)); n++ {
// 			if x.Flts[n].Cnt() > x.X.Max {
// 				x.X.Max = x.Flts[n].Cnt()
// 			}
// 		}
// 		x.X.Rng = x.X.Max.Sub(x.X.Min)
// 		// y-axis
// 		x.Y.Min = 0
// 		x.Y.Max = x.Flts[0].Max()
// 		for n := 1; n < len(x.Flts); n++ {
// 			max := x.Flts[n].Max()
// 			if max > x.Y.Max {
// 				x.Y.Max = max
// 			}
// 		}
// 		x.Y.Rng = x.Y.Max.Sub(x.Y.Min)

// 		// y-axis txt
// 		x.Y.MeasureInrvls(4)
// 		x.Y.Width = max(x.Y.MeasureTxt(x.Y.Min.String()), x.Y.MeasureTxt(x.Y.Max.String()))
// 		x.X.Len -= x.Y.Width + x.Y.Padding + x.Y.Margin

// 		//x.inr.Lft = x.crt.inr.Lft
// 		//x.inr.Rht = x.inr.Lft + x.X.Len + 1 // +1 for rht vrt ln
// 		//x.inr.Top = x.crt.inr.Top
// 		//x.inr.Btm = x.inr.Top + x.Y.Len

// 		x.X.PxPerVal = x.X.Len / float32(x.X.Rng)
// 		x.Y.PxPerVal = x.Y.Len / float32(x.Y.Rng)

// 		sys.Log("x.inr", x.inr)
// 		// x.Vis = vis.NewVis(x.inr.Width(), x.inr.Height())
// 		// x.Vis = vis.NewVis(x.crt.inr.Width(), x.crt.inr.Height())
// 	}
// 	return vis.Size{Width: x.X.Len, Height: x.Y.Len}
// }
// func (x *Flts) Draw() {
// 	if len(x.Flts) != 0 {
// 		// x.crt.vis.StkPthOpn(0, 100, 0, 100)
// 		// x.crt.vis.StkPthTo(150, 200)
// 		// x.crt.vis.StkPthTo(250, 260)
// 		// x.crt.vis.StkPthTo(400, 400)
// 		// x.crt.vis.StkPthCls()
// 		// x.crt.vis.Draw()

// 		offX := x.inr.Lft
// 		offY := x.inr.Top
// 		for _, vs := range x.Flts {
// 			sampleCnt := 1
// 			if len(*vs) > int(x.inr.Width()) {
// 				sampleCnt = int(float32(len(*vs)) / x.inr.Width())
// 				if sampleCnt < 0 {
// 					sampleCnt = 1
// 				}
// 			}
// 			sys.Log("sampleCnt", sampleCnt)
// 			sys.Log("vs", len(*vs))
// 			sys.Log("offX", offX, "offX+x.X.PxPerVal", offX+x.X.PxPerVal)
// 			sys.Log("vs.At(0)", vs.At(0), "vs.At(1)", vs.At(1))
// 			sys.Log("x.Y.Px(vs.At(0))", x.Y.Px(vs.At(0)), "x.Y.Px(vs.At(1))", x.Y.Px(vs.At(1)))
// 			x.crt.vis.StkPthOpn(offX, offX+x.X.PxPerVal, offY+x.Y.Px(vs.At(0)), offY+x.Y.Px(vs.At(1)))
// 			for n := 2; n < len(*vs); n += sampleCnt {
// 				if (*vs)[n] == flt.Max {
// 					continue
// 				}
// 				// sys.Log("x: offX+(float32(n)*x.X.PxPerVal)", offX+(float32(n)*x.X.PxPerVal))
// 				// sys.Log("y: x.Y.Px((*vs)[n]", x.Y.Px((*vs)[n]))
// 				x.crt.vis.StkPthTo(offX+(float32(n)*x.X.PxPerVal), offY+x.Y.Px((*vs)[n]))
// 			}
// 			x.crt.vis.StkPthCls()
// 		}
// 		x.crt.vis.Draw()
// 	}

// 	// // draw border
// 	// x.crt.vis.PxlRct(x.crt.inr, colornames.Gray)
// 	// draw y-axis lft
// 	x.crt.vis.PxlRct(x.inr, colornames.Gray)
// 	// draw y-axis txt
// 	x.crt.vis.Txt(x.inr.Rht+x.Y.Padding, x.inr.Top, 1, 1, x.Y.Font, x.Y.Max.Trnc(5).String())
// 	for _, inrvl := range x.Y.Inrvls {
// 		x.crt.vis.Txt(x.inr.Rht+x.Y.Padding, x.Y.Px(inrvl), 1, .5, x.Y.Font, inrvl.Trnc(5).String())
// 	}
// 	x.crt.vis.Txt(x.inr.Rht+x.Y.Padding, x.inr.Btm-4, 1, 0, x.Y.Font, x.Y.Min.Trnc(5).String())

// 	// draw plot vis
// 	// x.crt.vis.DrawVis(x.crt.margin.Left, x.crt.margin.Top, x.Vis)
// 	// x.crt.vis.DrawVis(0, 0, x.Vis)
// }
