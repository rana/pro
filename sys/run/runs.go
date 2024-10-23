package run

import (
	"sync"
)

type (
	// Runs is a thread-safe collection.
	Runs struct {
		itms []*Run
		mu   sync.RWMutex
	}
)

// NewRuns creates a new collection.
func NewRuns(itms ...*Run) *Runs {
	if itms == nil {
		itms = make([]*Run, 0)
	}
	return &Runs{itms: itms}
}

// Push pushes items to the back.
func (x *Runs) Push(vs ...*Run) *Runs {
	x.mu.Lock()
	x.itms = append(x.itms, vs...)
	x.mu.Unlock()
	return x
}

// Pop removes an element from the back.
func (x *Runs) Pop() (r *Run) {
	x.mu.Lock()
	if len(x.itms) > 0 {
		r = x.itms[len(x.itms)-1]
		x.itms = x.itms[:len(x.itms)-1]
		// sys.Logf("run.Runs.Pop %p", r)
	}
	x.mu.Unlock()
	return r
}

// Rem removes the specified element.
func (x *Runs) Rem(y *Run) (r *Run) {
	x.mu.Lock()
	for idx := 0; idx < len(x.itms); idx++ {
		if y == x.itms[idx] {
			r = y
			if idx == 0 && len(x.itms) == 1 {
				x.itms = x.itms[:0]
			} else if idx == len(x.itms)-1 {
				x.itms = x.itms[:idx]
			} else {
				x.itms = append(x.itms[:idx], x.itms[idx+1:]...)
			}
			break
		}
	}
	x.mu.Unlock()
	return r
}

func (x *Runs) Clr() {
	x.mu.Lock()
	x.itms = x.itms[:0]
	x.mu.Unlock()
}
func (x *Runs) Cls() { // locked from Runr
	for _, itm := range x.itms {
		close(itm.ActsC)
	}
	x.itms = x.itms[:0]
}
