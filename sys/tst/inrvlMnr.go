package tst

import (
	"sync/atomic"
	"sys"
	"sys/app"
	"sys/bsc/bnd"
	"sys/bsc/bnds"
	"sys/bsc/tme"
	"sys/bsc/unt"
)

type (
	InrvlMnr struct {
		Mnr
		Bnds *bnds.Bnds
	}
)

func NewInrvlMnr(ap *app.App) (r *InrvlMnr) {
	r = &InrvlMnr{}
	r.Id = sys.NextID()
	r.Bnds = bnds.New()
	r.Ap = ap
	return r
}
func (x *InrvlMnr) Rx(pkt bnd.Bnd) []sys.Act {
	x.Mu.Lock()
	x.Bnds.Push(pkt)
	atomic.AddUint32(&x.Cnt, 1)
	x.Mu.Unlock()
	return nil
}
func (x *InrvlMnr) WaitFor(expected unt.Unt, lim ...tme.Tme) {
	if len(lim) == 0 && expected == 0 {
		x.Bnds = nil
		return
	}
	x.Mnr.WaitFor(expected, lim...)
}
