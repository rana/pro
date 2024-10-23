package plt

import (
	"image"
	"sys"
	"sys/ana/vis"
	"sys/ana/vis/clr"
	"sys/ana/vis/fnt/roboto"
	"sys/ana/vis/pen"
	"sys/bsc/flt"
	"sys/bsc/unt"
)

var (
	Scl          = flt.Flt(1.0)
	StkWidth     = unt.Unt(1)
	ShpRadius    = unt.Unt(10)
	AxisPad      = unt.Unt(10)
	BarPad       = unt.Unt(10)
	Len          = unt.Unt(100)
	Pad          = unt.Unt(10)
	Mrgn         = vis.NewLenXY(2, 2, 2, 2)
	BakClr       = clr.Clr{R: 0x0, G: 0x0, B: 0x0, A: 0xff}
	BrdrClr      = clr.Clr{R: 0x61, G: 0x61, B: 0x61, A: 0xff}
	BrdrLen      = unt.Unt(1)
	InrvlTxtLen  = unt.Unt(50)
	InrvlTxtClrX = clr.Clr{R: 0x61, G: 0x61, B: 0x61, A: 0xff}
	InrvlTxtClrY = clr.Clr{R: 0xe0, G: 0xe0, B: 0xe0, A: 0xff}
	InrvlFnt     = roboto.Medium(12)
	MsgClr       = clr.Clr{R: 0x37, G: 0x47, B: 0x4f, A: 0xff}
	MsgFnt       = roboto.Medium(24)
	TitleClr     = clr.Clr{R: 0x9e, G: 0x9e, B: 0x9e, A: 0xff}
	TitleFnt     = roboto.Medium(14)
	PrfClr       = clr.Green500
	LosClr       = clr.Red500
	PrfPen       = pen.Green500
	LosPen       = pen.Red500
	OutlierLim   = flt.Flt(12.0)
)

type (
	Plt interface {
		Measure() vis.Siz
		Rndr()
		Draw(img *image.RGBA)
		Bse() *PltBse
		Sho() Plt
		Siz(w, h unt.Unt) Plt
		Scl(v flt.Flt) Plt
		HrzScl(v flt.Flt) Plt
		VrtScl(v flt.Flt) Plt
	}
	PltScp struct {
		Idx uint32
		Arr []Plt
	}
	PltMeasureSeg struct {
		Plt Plt
		Out vis.Siz
	}
	PltRndrSeg struct {
		Plt Plt
	}
	PltDrawSeg struct {
		Plt Plt
		Img *image.RGBA
	}
)

func PltMeasurePll(vs ...Plt) (r []vis.Siz) {
	if len(vs) != 0 {
		acts := make([]sys.Act, len(vs))
		for n, v := range vs {
			acts[n] = &PltMeasureSeg{Plt: v}
		}
		sys.Run().Pll(acts...)
		r = make([]vis.Siz, len(vs))
		for n, act := range acts {
			r[n] = act.(*PltMeasureSeg).Out
		}
	}
	return r
}
func PltRndrPll(vs ...Plt) {
	if len(vs) != 0 {
		acts := make([]sys.Act, len(vs))
		for n, v := range vs {
			acts[n] = &PltRndrSeg{Plt: v}
		}
		sys.Run().Pll(acts...)
	}
}
func (x *PltMeasureSeg) Act() { x.Out = x.Plt.Measure() }
func (x *PltRndrSeg) Act()    { x.Plt.Rndr() }
func (x *PltDrawSeg) Act()    { x.Plt.Draw(x.Img) }
