package tst

import (
	"sync/atomic"
	"sys"
	"sys/app"
	"sys/bsc/tme"
	"sys/bsc/tmes"
	"sys/bsc/unt"
)

type (
	CndMnr struct {
		Mnr
		Tmes *tmes.Tmes
	}
)

func NewCndMnr(ap *app.App) (r *CndMnr) {
	r = &CndMnr{}
	r.Id = sys.NextID()
	r.Tmes = tmes.New()
	r.Ap = ap
	return r
}
func (x *CndMnr) Rx(pkt tme.Tme) []sys.Act {
	x.Mu.Lock()
	x.Tmes.Push(pkt)
	atomic.AddUint32(&x.Cnt, 1)
	x.Mu.Unlock()
	return nil
}

func (x *CndMnr) WaitFor(expected unt.Unt, lim ...tme.Tme) {
	if len(lim) == 0 && expected == 0 {
		x.Tmes = nil
		return
	}
	x.Mnr.WaitFor(expected, lim...)
}
