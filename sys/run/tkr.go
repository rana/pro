package run

import (
	"sys"
	"sys/err"
	"time"
)

type (
	Tkr struct {
		start  time.Time
		f      func(ellapsed time.Duration)
		exit   chan bool
		ticker *time.Ticker
	}
)

func NewTkr(d time.Duration, f func(ellapsed time.Duration)) (r *Tkr) {
	r = &Tkr{}
	r.exit = make(chan bool)
	r.ticker = time.NewTicker(d)
	r.f = f
	go r.loop()
	return r
}
func (x *Tkr) Stop() {
	x.exit <- true
	x.ticker.Stop()
}
func (x *Tkr) loop() {
	defer func() {
		if v := recover(); v != nil {
			switch t := v.(type) {
			case *err.Err:
				sys.Log(t.Full())
			default:
				sys.Log(err.New(v).Full())
			}
		}
	}()
	x.start = time.Now()
	for {
		select {
		case <-x.exit:
			return
		case now := <-x.ticker.C:
			// Log.Println("Tkr.loop", now.Sub(x.start), "x.f != nil", x.f != nil)
			if x.f != nil {
				x.f(now.Sub(x.start))
			}
		}
	}
}
