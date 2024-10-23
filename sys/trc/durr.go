package trc

import (
	"sys"
	"time"
)

type (
	Durr struct {
		Name string
		Fst  time.Time
	}
)

func NewDurr(name ...string) (r *Durr) {
	r = &Durr{}
	if len(name) != 0 {
		r.Name += name[0]
	} else {
		r.Name = "~~~"
	}
	r.Fst = time.Now()
	return r
}
func (x *Durr) End() {
	sys.Logf("%v Ellapsed:%v", x.Name, time.Now().Sub(x.Fst))
}
func (x *Durr) Dur() time.Duration { return time.Now().Sub(x.Fst) }
