package tpl

import (
	"sys/k"
	"sys/tpl/atr"
)

type (
	DirPlt struct {
		DirBse
		Plt      *FlePlt
		TmeAxisX *FlePltTmeAxisX
		FltAxisY *FlePltFltAxisY
		Stm      *FlePltStm
		// Stgy         *FlePltStgy
		FltsSctr     *FlePltFltsSctr
		FltsSctrDist *FlePltFltsSctrDist
		Hrz          *FlePltHrz
		Vrt          *FlePltVrt
		Dpth         *FlePltDpth
		// Prcp         *FlePltPrcp
		// StmSplt      *FlePltStmSplt
		// PrcpSplt     *FlePltPrcpSplt
	}
	FlePlt struct {
		FleBse
		PrtActFn
	}
	FlePltPlts struct {
		FleBse
		PrtArr
		// PrtArrIdn
		PrtArrFld
		PrtArrStrWrt
		PrtArrBytWrt
		PrtBytes
	}
)

func (x *DirVis) NewPlt() (r *DirPlt) { // dir
	r = &DirPlt{}
	x.Plt = r
	r.Pkg = x.Pkg.New(k.Plt)
	r.NewPlt()
	r.NewTmeAxisX()
	r.NewFltAxisY()
	r.NewStm()
	// r.NewStgy()
	r.NewFltsSctr()
	r.NewFltsSctrDist()
	r.NewHrz()
	r.NewVrt()
	r.NewDpth()
	// r.NewPrcp()
	// r.NewStmSplt()
	// r.NewPrcpSplt()
	return r
}
func (x *DirPlt) NewPlt() (r *FlePlt) { // fle
	r = &FlePlt{}
	x.Plt = r
	r.Name = k.Plt
	r.Pkg = x.Pkg
	r.Ifc(r.Name, atr.TypUi)
	r.AddFle(r)
	r.NewArr()
	return r
}
func (x *FlePlt) NewArr() (r *FlePltPlts) {
	r = &FlePltPlts{}
	r.FleBse = *NewArr(x, &r.PrtArr, x.Pkg)
	r.PrtArrStrWrt.Ln = true
	r.AddFle(r)
	return r
}
func (x *FlePlt) InitFld(s *Struct) {
	x.Import(_sys.Ana.Vis)
	x.Import("image")
	var sig *MemSig
	sig = x.MemSigaOr(k.Measure, atr.PrtActFn)
	sig.OutPrm(NewExt("vis.Siz"))
	sig = x.MemSigaOr(k.Rndr, atr.PrtActFn)
	sig = x.MemSigaOr(k.Draw, atr.PrtActFn)
	sig.InPrm(NewExt("*image.RGBA"), "img")
	sig = x.MemSig(k.Bse)
	sig.OutPrm(NewExt("*PltBse"))

	sig = x.MemSig(k.Sho)
	sig.OutPrm(x)
	sig = x.MemSig(k.Siz)
	sig.InPrm(_sys.Bsc.Unt, "w")
	sig.InPrm(_sys.Bsc.Unt, "h")
	sig.OutPrm(x)
	sig = x.MemSig(k.Scl)
	sig.InPrm(_sys.Bsc.Flt, "v")
	sig.OutPrm(x)
	sig = x.MemSig(k.HrzScl)
	sig.InPrm(_sys.Bsc.Flt, "v")
	sig.OutPrm(x)
	sig = x.MemSig(k.VrtScl)
	sig.InPrm(_sys.Bsc.Flt, "v")
	sig.OutPrm(x)
}
func (x *FlePlt) InitVar() {
	x.Import("sys/ana/vis/fnt/roboto")
	x.Import(_sys.Ana.Vis)

	x.Var("Scl", "1.0", _sys.Bsc.Flt)

	x.Var("StkWidth", "1", _sys.Bsc.Unt)
	x.Var("ShpRadius", "10", _sys.Bsc.Unt)

	x.Var("AxisPad", "10", _sys.Bsc.Unt)
	x.Var("BarPad", "10", _sys.Bsc.Unt)
	x.Var("Len", "100", _sys.Bsc.Unt)
	x.Var("Pad", "10", _sys.Bsc.Unt)
	x.Var("Mrgn", "vis.NewLenXY(2, 2, 2, 2)", NewExt("vis.LenXY")).FnCall = true

	x.Var("BakClr", "R:0x0, G:0x0, B:0x0, A:0xff", _sys.Ana.Vis.Clr.Clr) // Black

	x.Var("BrdrClr", "R:0x61, G:0x61, B:0x61, A:0xff", _sys.Ana.Vis.Clr.Clr) // Grey700
	x.Var("BrdrLen", "1", _sys.Bsc.Unt)

	x.Var("InrvlTxtLen", "50", _sys.Bsc.Unt)
	x.Var("InrvlTxtClrX", "R:0x61, G:0x61, B:0x61, A:0xff", _sys.Ana.Vis.Clr.Clr) // Grey700
	x.Var("InrvlTxtClrY", "R:0xe0, G:0xe0, B:0xe0, A:0xff", _sys.Ana.Vis.Clr.Clr) // Grey300
	x.Var("InrvlFnt", "roboto.Medium(12)", NewExt("*fnt.Fnt")).FnCall = true

	x.Var("MsgClr", "R:0x37, G:0x47, B:0x4f, A:0xff", _sys.Ana.Vis.Clr.Clr) // BlueGrey500
	x.Var("MsgFnt", "roboto.Medium(24)", NewExt("*fnt.Fnt")).FnCall = true

	x.Var("TitleClr", "R:0x9e, G:0x9e, B:0x9e, A:0xff", _sys.Ana.Vis.Clr.Clr) // Grey500
	x.Var("TitleFnt", "roboto.Medium(14)", NewExt("*fnt.Fnt")).FnCall = true

	x.Var("PrfClr", "clr.Green500", _sys.Ana.Vis.Clr.Clr).FnCall = true
	x.Var("LosClr", "clr.Red500", _sys.Ana.Vis.Clr.Clr).FnCall = true
	x.Var("PrfPen", "pen.Green500", _sys.Ana.Vis.Pen.Pen).FnCall = true
	x.Var("LosPen", "pen.Red500", _sys.Ana.Vis.Pen.Pen).FnCall = true
	x.Var("OutlierLim", "12.0", _sys.Bsc.Flt) // fltsSctr: 12, 50, 100, 1000 pct
}
