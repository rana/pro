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
	SideMnr struct {
		Mnr
		Tmes  *tmes.Tmes
		Fltss []*flts.Flts
	}
)

func NewSideMnr(ap *app.App) (r *SideMnr) {
	r = &SideMnr{}
	r.Id = sys.NextID()
	r.Tmes = tmes.New()
	r.Ap = ap
	return r
}
func (x *SideMnr) Rx(pkt ana.TmeFlts) []sys.Act {
	x.Mu.Lock()
	x.Tmes.Push(pkt.Tme)
	x.Fltss = append(x.Fltss, pkt.Flts)
	atomic.AddUint32(&x.Cnt, 1)
	x.Mu.Unlock()
	return nil
}
func (x *SideMnr) WaitFor(expected unt.Unt, lim ...tme.Tme) {
	if len(lim) == 0 && expected == 0 {
		x.Tmes = nil
		x.Fltss = nil
		return
	}
	x.Mnr.WaitFor(expected, lim...)
}

// func (x *SideMnr) HstVals(valBnds *bnds.Bnds, sideVals *flts.Flts) (r *flts.Flts) {
// 	r = flts.New()
// 	for _, bnd := range *valBnds {
// 		r.Mrg(sideVals.In(bnd.Idx, bnd.Lim))
// 	}
// 	return r
// }
