package tpl

import (
	"sys/k"
	"sys/tpl/atr"
	"sys/tpl/mod"
)

type (
	DirScn struct {
		DirBse
		Scn  *FleScn
		Scnr *FleScnr
	}
	FleScn struct {
		FleBse
	}
)

func (x *DirLng) NewScn() (r *DirScn) {
	r = &DirScn{}
	x.Scn = r
	r.Pkg = x.Pkg.New(k.Scn)
	r.Scn = r.NewScn()
	r.Scnr = r.NewScnr()
	return r
}

func (x *DirScn) NewScn() (r *FleScn) {
	r = &FleScn{}
	r.Name = k.Scn
	r.Pkg = x.Pkg
	r.Scn()
	r.AddFle(r)
	return r
}
func (x *FleScn) Scn() (r *Struct) {
	r = x.Struct(k.Scn, atr.None) // Scn is a rune position within a string.
	r.Fld("Ch", Rune)             // Ch is the rune.
	r.Fld("Size", Int)            // Size is the size of the rune in bytes.
	r.Fld("Idx", _sys.Bsc.Unt)    // Idx is the rune index within the text.
	r.Fld("Ln", _sys.Bsc.Unt)     // Ln is the line number. Ln starts at one.
	r.Fld("Col", _sys.Bsc.Unt)    // Col is the column number. Col starts at one.
	r.Fld("End", Bool)            // End indicates whether the read position is at the end of the text.
	return r
}
func (x *FleScn) InitTypFn() {
	x.String()
}
func (x *FleScn) String() (r *TypFn) {
	x.Import("fmt")
	r = x.TypFn(k.String)
	r.Rxr.Mod = mod.Ptr
	r.OutPrm(String)
	r.Add("return fmt.Sprintf(\"Ch:%v:'%v' Size:%v Idx:%v Ln:%v Col:%v End:%v \\n\", x.Ch, string(x.Ch), x.Size, x.Idx, x.Ln, x.Col, x.End)")
	return r
}
