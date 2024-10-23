package run

import (
	"sync"
	"sys"
	"sys/err"
)

type (
	// Act is an action interface.
	// Act interface {
	// 	Act()
	// }
	Acts []sys.Act

	// Waitr is a WaitGroup interface.
	Waitr interface {
		Done()
	}
	// Seq is a sequence action.
	Seq struct {
		Acts []sys.Act
		*sync.WaitGroup
	}
	// Grp is a wait group action.
	Grp struct {
		Itm sys.Act
		*sync.WaitGroup
	}
	// Seg is a segment action.
	Seg struct {
		Idx  uint32
		Lim  uint32
		Acts []sys.Act
		*sync.WaitGroup
	}
	Fn struct {
		Fn func()
	}
)

// Act performs a sequence of actions.
func (x *Seq) Act() {
	for _, a := range x.Acts {
		a.Act()
	}
	x.Done()
}
func (x *Seq) Ifc() interface{} { return x }

// Act performs an action as part of a wait group.
func (x *Grp) Act() {
	x.Itm.Act()
	x.Done()
}
func (x *Grp) Ifc() interface{} { return x }

// Act performs a segment of actions as part of a wait group.
func (x *Seg) Act() {
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
	// sys.Logf("Act.Seg idx:%v lim:%v acts:%v", x.Idx, x.Lim, len(x.Acts))
	for n := x.Idx; n < x.Lim; n++ {
		// sys.Logf("Act.Seg n:%v acts:%v x.Acts[n] == nil:%v", n, len(x.Acts), x.Acts[n] == nil)
		x.Acts[n].Act()
		x.Done()
	}
}
func (x *Seg) Ifc() interface{} { return x }

func (x *Fn) Act() {
	x.Fn()
}
func (x *Fn) Ifc() interface{} { return x }
func (x *Fn) DecTier()         {}             // prv.Tx interface
func (x *Fn) Ret() []sys.Act   { return nil } // prv.Tx interface
func (x *Fn) Tier() int        { return 0 }   // prv.Tx interface

func NewActs() *Acts {
	var r Acts
	return &r
}
func (x *Acts) Push(vs ...sys.Act) {
	for _, v := range vs {
		*x = append(*x, v)
	}
}
func (x *Acts) Del(idx int) {
	if idx == 0 && len(*x) == 1 {
		*x = (*x)[:0]
	} else if idx == len(*x)-1 {
		*x = (*x)[:idx]
	} else {
		*x = append((*x)[:idx], (*x)[idx+1:]...)
	}
}
func (x *Acts) Elm(idx int) sys.Act { return (*x)[idx] }
func (x *Acts) Clr()                { *x = (*x)[:0] }
func (x *Acts) Cnt() int            { return len(*x) }
