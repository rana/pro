package tpl

import (
	"fmt"
	"strings"
	"sys/k"
	"sys/ks"
	"sys/tpl/atr"
)

type (
	DirPen struct {
		DirBse
		Pen *FlePen
	}
	FlePen struct {
		FleBse
	}
	FlePens struct {
		FleBse
		PrtArr
	}
)

func (x *DirVis) NewPen() (r *DirPen) { // dir
	r = &DirPen{}
	x.Pen = r
	r.Pkg = x.Pkg.New(k.Pen)
	r.NewPen()
	r.Pen.NewArr()
	return r
}
func (x *DirPen) NewPen() (r *FlePen) { // fle
	r = &FlePen{}
	x.Pen = r
	r.Name = k.Pen
	r.Pkg = x.Pkg
	r.Struct(r.Name, atr.TypUiStruct)
	r.AddFle(r)
	return r
}
func (x *FlePen) NewArr() (r *FlePens) {
	r = &FlePens{}
	r.FleBse = *NewArr(x, &r.PrtArr, x.Pkg)
	r.AddFle(r)
	return r
}
func (x *FlePen) InitFld(s *Struct) {
	s.Fld("Clr", _sys.Ana.Vis.Clr.Clr).Atr = atr.SetGet
	s.Fld("Wid", _sys.Bsc.Unt).Atr = atr.SetGet
}
func (x *FlePen) InitPkgFn() {
	x.New()
	x.Rgba()
	x.Rgb()
	x.Hex()
}
func (x *FlePen) New() (r *PkgFn) {
	r = x.PkgFna(k.New, atr.Lng)
	r.InPrm(_sys.Ana.Vis.Clr.Clr, "clr")
	r.InPrmVariadic(_sys.Bsc.Unt, "wid")
	r.OutPrm(x, "p")
	r.Add("p.Clr = clr")
	r.Add("if len(wid) != 0 {")
	r.Add("p.Wid = wid[0]")
	r.Add("} else {")
	r.Add("p.Wid = 1")
	r.Add("}")
	r.Add("return p")
	return r
}
func (x *FlePen) Rgba() (r *PkgFn) {
	r = x.PkgFna(k.Rgba, atr.Lng)
	r.InPrm(_sys.Bsc.Flt, "r")
	r.InPrm(_sys.Bsc.Flt, "g")
	r.InPrm(_sys.Bsc.Flt, "b")
	r.InPrm(_sys.Bsc.Flt, "a")
	r.InPrmVariadic(_sys.Bsc.Unt, "wid")
	r.OutPrm(x, "p")
	r.Add("p.Clr = clr.Rgba(r, g, b, a)")
	r.Add("if len(wid) != 0 {")
	r.Add("p.Wid = wid[0]")
	r.Add("} else {")
	r.Add("p.Wid = 1")
	r.Add("}")
	r.Add("return p")
	return r
}
func (x *FlePen) Rgb() (r *PkgFn) {
	r = x.PkgFna(k.Rgb, atr.Lng)
	r.InPrm(_sys.Bsc.Flt, "r")
	r.InPrm(_sys.Bsc.Flt, "g")
	r.InPrm(_sys.Bsc.Flt, "b")
	r.InPrmVariadic(_sys.Bsc.Unt, "wid")
	r.OutPrm(x, "p")
	r.Add("p.Clr = clr.Rgb(r, g, b)")
	r.Add("if len(wid) != 0 {")
	r.Add("p.Wid = wid[0]")
	r.Add("} else {")
	r.Add("p.Wid = 1")
	r.Add("}")
	r.Add("return p")
	return r
}
func (x *FlePen) Hex() (r *PkgFn) {
	r = x.PkgFna(k.Hex, atr.Lng)
	r.InPrm(_sys.Bsc.Str, "txt")
	r.InPrmVariadic(_sys.Bsc.Unt, "wid")
	r.OutPrm(x, "p")
	r.Add("p.Clr = clr.Hex(txt)")
	r.Add("if len(wid) != 0 {")
	r.Add("p.Wid = wid[0]")
	r.Add("} else {")
	r.Add("p.Wid = 1")
	r.Add("}")
	r.Add("return p")
	return r
}
func (x *FlePen) InitTypFn() {
	x.Opa()
	x.Inv()
}
func (x *FlePen) Opa() (r *TypFn) {
	r = x.TypFn(k.Opa)
	r.InPrm(_sys.Bsc.Flt, "pct")
	r.OutPrm(x)
	r.Add("x.Clr = x.Clr.Opa(pct)")
	r.Add("return x")
	return r
}
func (x *FlePen) Inv() (r *TypFn) {
	r = x.TypFn(k.Inv)
	r.OutPrm(x)
	r.Add("x.Clr = x.Clr.Inv()")
	r.Add("return x")
	return r
}
func (x *FlePen) InitVar() {
	wid := 1
	x.Var("Black", wid)
	x.Var("White", wid)
	for n, clr := range ks.Clrs {
		for _, num := range ks.ClrNums {
			name := fmt.Sprintf("%v%v", strings.Title(clr), num)
			x.Var(name, wid)
		}
		if n < len(ks.ClrAs) {
			for _, numA := range ks.ClrNumAs {
				nameA := fmt.Sprintf("%vA%v", strings.Title(clr), numA)
				x.Var(nameA, wid)
			}
		}
	}
}
func (x *FlePen) Var(name string, wid int) {
	x.FleBse.Var(name, fmt.Sprintf("Clr: clr.%v, Wid: %v", name, wid))
}
