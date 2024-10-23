package tst

import (
	"sync/atomic"
	"sys"
	"sys/ana"
	"sys/app"
	"sys/bsc/tme"
	"sys/bsc/unt"
)

type (
	StgyMnr struct {
		Mnr
		Trds *ana.Trds
	}
)

func NewStgyMnr(ap *app.App) (r *StgyMnr) {
	r = &StgyMnr{}
	r.Id = sys.NextID()
	r.Trds = ana.NewTrds()
	r.Ap = ap
	return r
}

func (x *StgyMnr) Rx(pkt *ana.Trd) []sys.Act {
	x.Mu.Lock()
	x.Trds.Push(pkt)
	atomic.AddUint32(&x.Cnt, 1)
	x.Mu.Unlock()
	return nil
}
func (x *StgyMnr) WaitFor(expected unt.Unt, lim ...tme.Tme) {
	if len(lim) == 0 && expected == 0 {
		x.Trds = nil
		return
	}
	x.Mnr.WaitFor(expected, lim...)
}
