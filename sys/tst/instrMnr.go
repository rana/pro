package tst

import (
	"sync/atomic"
	"sys"
	"sys/ana"
	"sys/app"
	"sys/bsc/tme"
	"sys/bsc/unt"
	"sys/bsc/unts"
)

type (
	InstrMnr struct {
		Mnr
		Idxs *unts.Unts
	}
)

func NewInstrMnr(ap *app.App) (r *InstrMnr) {
	r = &InstrMnr{}
	r.Id = sys.NextID()
	r.Idxs = unts.New()
	r.Ap = ap
	return r
}
func (x *InstrMnr) Rx(pkt ana.TmeIdx) []sys.Act {
	x.Mu.Lock()
	if pkt.Idx != unt.Max {
		x.Idxs.Push(pkt.Idx)
		atomic.AddUint32(&x.Cnt, 1)
	}
	x.Mu.Unlock()
	return nil
}
func (x *InstrMnr) WaitFor(expected unt.Unt, lim ...tme.Tme) {
	if len(lim) == 0 && expected == 0 {
		x.Idxs = nil
		return
	}
	x.Mnr.WaitFor(expected, lim...)
}

func (x *InstrMnr) Stm() (r *ana.Stm) {
	contiguous := false
	if x.Idxs.Cnt() > 1 {
		contiguous = true
		for n := unt.One; n < x.Idxs.Cnt(); n++ {
			if x.Idxs.At(n).Sub(x.Idxs.At(n - 1)).Neq(unt.One) {
				contiguous = false
				break
			}
		}
	}
	if contiguous {
		r = x.I.HstStm.In(x.Idxs.Fst(), x.Idxs.Lst()+1)
	} else {
		r = ana.NewStm() // create output stm for test comparison
		for n := unt.Zero; n < x.Idxs.Cnt(); n++ {
			r.PushTic(x.I.HstStm.Tic(x.Idxs.At(n)))
		}
	}
	return r
}
