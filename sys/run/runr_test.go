package run_test

import (
	"fmt"
	"runtime"
	"sys"
	"sys/run"
	"sys/tst"
	"testing"
	"time"
)

type (
	Step struct {
		Wait     time.Duration
		CallTime time.Time
	}
)

func (x *Step) Act() {
	<-time.NewTimer(x.Wait).C
	x.CallTime = time.Now()
}

func GenSteps(wait time.Duration) (r []sys.Act) {
	for n := 0; n < runtime.NumCPU(); n++ {
		r = append(r, &Step{Wait: wait})
	}
	return r
}

func TestSeq(t *testing.T) {
	wait := time.Microsecond
	steps := GenSteps(wait)
	runr := run.NewRunr()
	tst.NotNil(t, runr)
	start := time.Now()
	runr.Seq(steps...)
	tst.DurationLss(t, wait*time.Duration(len(steps)), time.Now().Sub(start))
	if len(steps) > 0 {
		for n := 1; n < len(steps); n++ { // ensure each called in order
			tst.TimeLss(t, steps[n-1].(*Step).CallTime, steps[n].(*Step).CallTime, fmt.Sprintf("step %v", n-1))
		}
	}
}

func TestPll(t *testing.T) {
	wait := time.Microsecond
	steps := GenSteps(wait)
	runr := run.NewRunr()
	tst.NotNil(t, runr)
	runr.Pll(steps...)
	for n := 0; n < len(steps); n++ { // ensure each called
		tst.TimeNeq(t, time.Time{}, steps[n].(*Step).CallTime)
	}
}
