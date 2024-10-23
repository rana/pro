package xpr

import (
	"sys/err"
	"sys/lng/pro/xpr/knd"
)

type (
	Scp struct {
		Prnt *Scp
		Vars map[string]KndIdx
		Cnts map[knd.Knd]uint32
	}
	KndIdx struct {
		Knd knd.Knd
		Idx uint32
	}
)

func NewScp(prnt ...*Scp) (r *Scp) {
	r = &Scp{
		Vars: make(map[string]KndIdx),
		Cnts: make(map[knd.Knd]uint32),
	}
	if len(prnt) != 0 {
		r.Prnt = prnt[0]
	}
	return r
}
func (x *Scp) Knd(idn string) (knd knd.Knd, exists bool) {
	var kndIdx KndIdx
	cur := x
	for cur != nil {
		kndIdx, exists = cur.Vars[idn]
		if exists {
			return kndIdx.Knd, exists
		}
		cur = cur.Prnt
	}
	return knd, false
}
func (x *Scp) Idx(idn string) uint32 { return x.Vars[idn].Idx }
func (x *Scp) Decl(idn string, knd knd.Knd) {
	kndIdx, exists := x.Vars[idn]
	if exists {
		if kndIdx.Knd != knd {
			err.Panicf("idn redeclared: cannot redeclare an identifier with a different type (idn:%v expected:%v actual:%v)", idn, kndIdx.Knd, knd)
		}
	}
	x.Vars[idn] = KndIdx{Knd: knd, Idx: x.Cnts[knd]}
	x.Cnts[knd]++
}
