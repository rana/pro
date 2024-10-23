package tpl

import (
	"sys/k"
	"sys/tpl/atr"
)

type (
	FleSysMu struct {
		FleBse
	}
)

func (x *DirSys) NewMu() (r *FleSysMu) {
	r = &FleSysMu{}
	x.Mu = r
	r.Name = k.Mu
	r.Pkg = x.Pkg
	r.StructPtr(k.Mu, atr.LngScp)
	r.AddFle(r)
	return r
}
func (x *FleSysMu) InitPkgFn() {
	x.New()
}
func (x *FleSysMu) New() (r *PkgFn) {
	r = x.PkgFna("NewMu", atr.Lng)
	r.OutPrm(x)
	r.Addf("return %v{}", r.OutTyp().Adr(x))
	return r
}
func (x *FleSysMu) InitFld(s *Struct) {
	s.FldTyp(Mutex)
}
func (x *FleSysMu) InitTypFn() {
	x.Lck()
	x.Ulck()
}
func (x *FleSysMu) Lck() (r *TypFn) {
	r = x.TypFna("Lck", atr.Lng)
	r.OutPrm(x)
	r.Add("x.Mutex.Lock()")
	r.Add("return x")
	return r
}
func (x *FleSysMu) Ulck() (r *TypFn) {
	r = x.TypFna("Ulck", atr.Lng)
	r.OutPrm(x)
	r.Add("x.Mutex.Unlock()")
	r.Add("return x")
	return r
}
