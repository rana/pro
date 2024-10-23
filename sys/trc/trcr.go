package trc

import (
	"fmt"
	"sys"
	"sys/bsc/tme"
	"time"
)

type (
	Trcr struct {
		Name string
		Fst  time.Time
		Cnt  uint64
	}
)

func New(name ...string) (r *Trcr) {
	r = &Trcr{}
	if len(name) != 0 {
		r.Name += name[0]
		sys.Log(r.Name)
	}
	r.Fst = time.Now()
	return r
}
func Newf(format string, args ...interface{}) (r *Trcr) {
	return New(fmt.Sprintf(format, args...))
}
func (x *Trcr) End(txt ...string) {
	var cntTxt, durTxt, endTxt string
	if x.Cnt != 0 {
		cntTxt = fmt.Sprintf("cnt:%v", x.Cnt)
	}
	dur := time.Now().Sub(x.Fst)
	if dur < time.Second {
		durTxt = fmt.Sprintf("%v", dur)
	} else {
		durTxt = tme.Duration(dur).String()
	}
	if len(txt) != 0 {
		endTxt = txt[0]
	}
	sys.Logf("%v %v %v %v", x.Name, durTxt, cntTxt, endTxt)
}
