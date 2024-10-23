package tst

import (
	"sync/atomic"
	"sys"
	"sys/ana"
	"sys/app"
	"sys/bsc/flts"
	"sys/bsc/tme"
	"sys/bsc/tmes"
	"sys/bsc/unt"
)

type (
	StmMnr struct {
		Mnr
		Tmes *tmes.Tmes
		Vals *flts.Flts
	}
)

func NewStmMnr(ap *app.App) (r *StmMnr) {
	r = &StmMnr{}
	r.Id = sys.NextID()
	r.Tmes = tmes.New()
	r.Vals = flts.New()
	r.Ap = ap
	return r
}
func (x *StmMnr) Rx(pkt ana.TmeFlt) []sys.Act {
	x.Mu.Lock()
	x.Tmes.Push(pkt.Tme)
	x.Vals.Push(pkt.Flt)
	atomic.AddUint32(&x.Cnt, 1)
	x.Mu.Unlock()
	return nil
}
func (x *StmMnr) WaitFor(expected unt.Unt, lim ...tme.Tme) {
	if len(lim) == 0 && expected == 0 {
		x.Tmes = nil
		x.Vals = nil
		return
	}
	x.Mnr.WaitFor(expected, lim...)
}
