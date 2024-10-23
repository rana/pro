package tst

import (
	"sync"
	"sync/atomic"
	"sys/ana"
	"sys/app"
	"sys/bsc/str"
	"sys/bsc/tme"
	"sys/bsc/unt"
	"time"
)

type (
	Mnr struct {
		Id   uint32
		Slot uint32
		Cnt  uint32
		Mu   sync.Mutex
		Ap   *app.App
		I    *ana.Instr
	}
)

func (x *Mnr) StartFor(i *ana.Instr, expected unt.Unt, lim ...tme.Tme) {
	x.Start(i)
	x.WaitFor(expected, lim...)
}
func (x *Mnr) Start(i *ana.Instr) {
	i.RltStm = i.HstStm
	i.RltStm.RxIdx = 0
	x.I = i
	x.Ap.Ticr.OpnRx(map[str.Str]*ana.Instr{i.Name: i})
	x.Ap.Ticr.RxPktC <- 0
}
func (x *Mnr) WaitFor(expected unt.Unt, lim ...tme.Tme) {
	var esc tme.Tme
	if len(lim) > 0 {
		esc = lim[0]
	} else {
		esc = tme.Duration(*Esc)
	}
	tkr := time.NewTicker(time.Microsecond)
	start := tme.Now()
	for true {
		<-tkr.C
		if atomic.LoadUint32(&x.Cnt) >= uint32(expected) {
			break
		}
		if tme.Now().Sub(start) > esc {
			break
		}
	}
	<-time.NewTimer(time.Millisecond).C // NEEDED FOR SOME REASON
}
