package tpl

import (
	"fmt"
	"strings"
	"sys/k"
)

type (
	PrtStructIdn struct {
		PrtBse
		Eql *TypFn
		Neq *TypFn
	}
)

func (x *PrtStructIdn) InitPrtTypFn() {
	x.Eql = x.op(k.Eql)
	x.Neq = x.op(k.Neq)
	// // tst
	// if x.f.Tst != nil {
	// 	x.f.Tst.IdnPrt()
	// }
}
func (x *PrtStructIdn) op(name string) (r *TypFn) {
	rxr := x.f.Typ().(*Struct)
	r = x.f.TypFn(name)
	r.InPrm(x.f.Typ(), "a")
	r.OutPrm(_sys.Bsc.Bol)
	if len(rxr.Flds) == 0 {
		r.Add("panic(\"no flds\")")
		r.T.Empty = true
	} else {
		a := &strings.Builder{}
		expected := &strings.Builder{}
		a.WriteString("return ")
		expected.WriteString("expected := ")
		var fst string
		hasFld := false
		for _, fld := range rxr.Flds {
			if !fld.IsEqlSkp() {
				if hasFld {
					if k.Eql == strings.ToLower(name) {
						fst = "&&"
					} else {
						fst = "||"
					}
				}
				a.WriteString(fmt.Sprintf(" %v %v.%v.%v(%v.%v)", fst, r.Rxr.Name, fld.Name, r.Name, r.InPrms[0].Name, fld.Name))
				expected.WriteString(fmt.Sprintf(" %v cse.x.%v.%v(cse.a.%v)", fst, fld.Name, r.Name, fld.Name))
				hasFld = true
			}
		}
		r.Add(a.String())
		// test
		if x.f.Test != nil && !x.t.IsUi() {
			r.T.Add(expected.String())
		}
	}
	return r
}
