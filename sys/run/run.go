package run

import (
	"sys"
	"sys/err"
)

type (
	// Run is an action loop on a go routine.
	Run struct {
		ActsC chan sys.Act
		Act   sys.Act
		Err   error
		Runr  *Runr
	}
)

// NewRun creates a new Run.
func NewRun(runr *Runr) (r *Run) {
	r = &Run{
		ActsC: make(chan sys.Act),
		Runr:  runr,
	}
	// sys.Logf("run.NewRun %p", r)
	go r.loop()
	return r
}

// loop listens for actions to be run. loop runs on a separate routine.
func (x *Run) loop() {
	var cur sys.Act
	var open bool
	defer func() {
		rec := recover()
		if rec != nil {
			// sys.Logf("run.Run.loop RECOVER run:%p cur:%p", x, cur)
			x.Err = err.New(rec)
			waitr, ok := cur.(Waitr)
			if ok {
				waitr.Done()
			}
		}
	}()
	for {
		select {
		case cur, open = <-x.ActsC:
			if !open {
				return
			}
			// sys.Logf("run.Run.loop run:%p cur:%p", x, cur)
			cur.Act()
			x.Runr.Put(x)
		}
	}
}
