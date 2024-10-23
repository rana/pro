package tpl

import (
	"fmt"
	"sys"
	"sys/tpl/atr"
)

type (
	PrtStructGrp struct {
		PrtBse
		Arr *Arr
		Elm *Struct
		// Ifc *Ifc
	}
)

func NewGrp(elmFle Fle, p *PrtStructGrp) (r *FleBse) {
	r = &FleBse{}
	r.Name = fmt.Sprintf("%vGrp", elmFle.Bse().Name)
	r.Pkg = elmFle.Bse().Pkg
	p.Arr = elmFle.Bse().arr
	if s, ok := p.Arr.Alias.Elm.(*Struct); ok {
		p.Elm = s
	} else if ifc, ok := p.Arr.Alias.Elm.(*Ifc); ok && ifc.bse != nil {
		// x.Ifc = ifc
		p.Elm = ifc.bse
	}
	sys.Log("PrtStructGrp.NewGrp", r.Name)
	return r
}

// func (x *PrtStructGrp) InitPrtTyp() {
// 	x.Arr = x.f.GetPrt((*PrtArr)(nil)).(*PrtArr).Arr
// 	if s, ok := x.Arr.Alias.Elm.(*Struct); ok {
// 		x.Elm = s
// 	} else if ifc, ok := x.Arr.Alias.Elm.(*Ifc); ok && ifc.bse != nil {
// 		// x.Ifc = ifc
// 		x.Elm = ifc.bse
// 	}
// }
func (x *PrtStructGrp) InitPrtTyp() {
	sys.Log("PrtStructGrp.InitPrtTyp", x.t.Name)
	sys.Log("PrtStructGrp.InitPrtTyp", x.Elm != nil)
	if x.Elm != nil {
		flds := x.Elm.SelFlds(func(f *Fld) bool {
			return f.IsFstUpr() &&
				f.IsGrp() &&
				f.Typ.Bse().IsIdn() &&
				f.Typ.Bse().IsAct()
		})
		sys.Log("PrtStructGrp.InitPrtTyp", len(*flds))
		kvsMaps := make(map[string]*Map)
		kvsSlices := make(map[string]*Alias)
		for _, fld := range *flds {
			sys.Log("PrtStructGrp.InitPrtTyp", x.t.Name, fld.Name, fld.IsIdn(), fld.IsRel())
			kvsName := fmt.Sprintf("%vBy%v", x.t.Title(), fld.Typ.Title())
			sys.Log("PrtStructGrp.InitPrtTyp", kvsName)
			kvsMap, ok := kvsMaps[kvsName]
			if !ok {
				kvsMap := x.f.Map(kvsName, fld.Typ, x.t, atr.LngScp)
				kvsMaps[kvsName] = kvsMap
				// rr := x.f.TypFn("Abc123", kvsMap)
				// rr.OutPrm(_sys.Bsc.Bol, "r")
				// rr.Add("return r")

				kvsSlices[kvsName] = x.f.AliasSlice(kvsName+"s", _sys.Bsc.Bol.Typ().Bse(), atr.LngScp)
				// rr = x.f.TypFn("Def456", kvsSlices[kvsName])
				// rr.OutPrm(_sys.Bsc.Bol, "r")
				// rr.Add("return r")
			}
			kvsSlice := kvsSlices[kvsName]
			sys.Log("PrtStructGrp.InitPrtTyp", kvsMap, kvsSlice)
			r := x.f.TypFnf("GrpBy%v", fld.Name)
			r.Atr = atr.Lng // no lng test
			r.InPrm(fld.Typ, fld.Camel())
			r.OutPrm(kvsSlice, "r")
			r.Add("return r")
		}
	}
}
