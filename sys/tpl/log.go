package tpl

import (
	"sys/k"
	"sys/tpl/atr"
)

type (
	DirLog struct {
		DirBse
		Log *FleLog
	}
	FleLog struct {
		FleBse
	}
)

func (x *DirSys) NewLog() (r *DirLog) { // dir
	r = &DirLog{}
	x.Log = r
	r.Pkg = x.Pkg.New(k.Log)
	r.NewLog() // IMPORTANT: CALL ORDER IS PRS ORDER
	return r
}
func (x *DirLog) NewLog() (r *FleLog) { // fle
	r = &FleLog{}
	x.Log = r
	r.Name = k.Log
	r.Pkg = x.Pkg.New(r.Name)
	r.Alias("Logr", String, atr.Lng) // only to get the PkgFn into Lng
	r.AddFle(r)
	return r
}
func (x *FleLog) InitPkgFn() {
	x.Ifo()
	x.Ifof()
}
func (x *FleLog) Ifo() (r *PkgFn) {
	x.Import(_sys)
	r = x.PkgFna(k.Ifo, atr.Lng)
	r.InPrmVariadic(Interface, "vs")
	r.OutPrm(_sys.Bsc.Str, "r")
	r.Add("sys.Log(vs...)")
	r.Add("return r")
	return r
}
func (x *FleLog) Ifof() (r *PkgFn) {
	x.Import(_sys)
	r = x.PkgFna("Ifof", atr.Lng)
	r.InPrm(_sys.Bsc.Str, "tmpl")
	r.InPrmVariadic(Interface, "vs")
	r.OutPrm(_sys.Bsc.Str, "r")
	r.Add("sys.Logf(string(tmpl), vs...)")
	r.Add("return r")
	return r
}
