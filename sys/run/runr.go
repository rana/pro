package run

import (
	"runtime"
	"sync"
	"sys"
	"sys/err"
)

var (
	// Val is an instance of Runr.
	Val *Runr
)

type (
	// Runr runs actions.
	Runr struct {
		on  *Runs
		off *Runs
	}
)

// NewRunr creates an action runner.
func NewRunr() (r *Runr) {
	r = &Runr{
		on:  NewRuns(),
		off: NewRuns(),
	}
	return r
}

// Seq sequentially runs actions on a background go routine. Seq waits for each action to complete.
func (x *Runr) Seq(acts ...sys.Act) {
	if len(acts) == 0 {
		return
	}
	rn := x.Get()
	wg := &sync.WaitGroup{}
	wg.Add(1)
	rn.ActsC <- &Seq{Acts: acts, WaitGroup: wg}
	wg.Wait()
	if rn.Err != nil {
		err.Panic(rn.Err)
	}
}

// Pll runs actions in parallel on background go routines. Pll waits for all actions to complete or panic before continuing.
func (x *Runr) Pll(acts ...sys.Act) {
	if len(acts) == 0 {
		return
	}
	var rns []*Run
	wg := &sync.WaitGroup{}
	wg.Add(len(acts))
	if len(acts) <= runtime.NumCPU() { // LEQ CPU CNT
		rns = make([]*Run, len(acts))
		for n, a := range acts {
			rns[n] = x.Get()
			rns[n].ActsC <- &Grp{Itm: a, WaitGroup: wg}
		}
	} else { // GTR CPU CNT
		// SEGMENT ACTS TO NUM OF CPU SO NOT TO OVERWHLEM MEMORY
		// POSSIBLE TO HAVE THOUSANDS OF PLL
		elmCnt := uint32(len(acts))
		segCnt := uint32(runtime.NumCPU())
		segLen := elmCnt / segCnt
		idx := uint32(0)
		lim := segCnt
		for n := uint32(0); n < segCnt; n++ {
			if n > 0 {
				idx = lim
				lim += segLen
			}
			if idx >= elmCnt {
				break
			}
			if n == segCnt-1 || lim > elmCnt {
				lim = elmCnt
			}
			run := x.Get()
			rns = append(rns, run)
			// sys.Logf("Runr.Pll idx:%v lim:%v acts:%v", idx, lim, len(acts))
			run.ActsC <- &Seg{
				Idx:       idx,
				Lim:       lim,
				Acts:      acts,
				WaitGroup: wg,
			}
		}
	}
	wg.Wait()
	for _, rn := range rns {
		if rn.Err != nil {
			switch e := rn.Err.(type) {
			case *err.Err:
				err.Panic(e.Full())
			default:
				err.Panic(rn.Err)
			}
		}
	}
}

func (x *Runr) PllFn(vs ...func()) {
	fns := make([]sys.Act, len(vs))
	for n, fn := range vs {
		fns[n] = &Fn{Fn: fn}
	}
	x.Pll(fns...)
}

// Get returns a run.
func (x *Runr) Get() (r *Run) {
	r = x.off.Pop()
	if r == nil {
		r = NewRun(x)
	}
	x.on.Push(r)
	return r
}

// Put puts a run.
func (x *Runr) Put(v *Run) {
	// sys.Logf("run.Runr.Put %p", v)
	x.on.Rem(v)
	x.off.Push(v)
}

func (x *Runr) Cls() {
	x.off.mu.Lock()
	x.on.mu.Lock()
	x.off.Cls()
	x.on.Cls()
	x.off.mu.Unlock()
	x.on.mu.Unlock()
}
