package sys

import (
	"sync"
)

type (
	Mu struct {
		sync.Mutex
	}
	MuScp struct {
		Idx uint32
		Arr []*Mu
	}
)

func NewMu() *Mu { return &Mu{} }
func (x *Mu) Lck() *Mu {
	x.Mutex.Lock()
	return x
}
func (x *Mu) Ulck() *Mu {
	x.Mutex.Unlock()
	return x
}
