package tpl

import (
	"fmt"
	"sys/tpl/atr"
	"sys/tpl/mod"
)

type (
	PrtArrFldGrp struct {
		PrtBse
		Arr    *Arr
		Elm    *Struct
		Ifc    *Ifc
		ElmRel *PrtStructRel
	}
)

func (x *PrtArrFldGrp) InitPrtTyp() {
	x.Arr = x.f.GetPrt((*PrtArr)(nil)).(*PrtArr).Arr
	if s, ok := x.Arr.Alias.Elm.(*Struct); ok {
		x.Elm = s
		x.ElmRel = x.Elm.Fle.Bse().GetPrt((*PrtStructRel)(nil)).(*PrtStructRel)
	} else if ifc, ok := x.Arr.Alias.Elm.(*Ifc); ok && ifc.bse != nil {
		x.Ifc = ifc
		x.Elm = ifc.bse
		x.ElmRel = x.Elm.Fle.Bse().GetPrt((*PrtStructRel)(nil)).(*PrtStructRel)
	}
}

func (x *PrtArrFldGrp) InitPrtTypFn() {
	if x.Elm != nil {
		var bse string
		if x.Ifc != nil {
			bse = ".Bse()"
		}
		flds := x.Elm.SelFlds(func(f *Fld) bool {
			return f.IsFstUpr() &&
				f.IsGrp() &&
				f.Typ.Bse().IsIdn() &&
				f.Typ.Bse().IsAct()
		})
		kvsMaps := make(map[string]*Map)
		// kvsSlices := make(map[string]*Alias)
		for _, fld := range *flds {
			kvsName := fmt.Sprintf("%vBy%v", x.t.Title(), fld.Typ.Title())
			kvsMap, ok := kvsMaps[kvsName]
			if !ok {
				kvsMap = x.f.Map(kvsName, fld.Typ, x.t, atr.LngScp)
				kvsMap.Mod = mod.Ptr
				kvsMaps[kvsName] = kvsMap
				// kvsSlices[kvsName] = x.f.AliasSlice(kvsName+"s", kvsMap, atr.TypNoTest)
			}
			// kvsSlice := kvsSlices[kvsName]
			// sys.Log("PrtArrFldGrp.InitPrtTypFn", kvsMap, kvsSlice)
			r := x.f.TypFnf("GrpBy%v", fld.Name)
			r.Atr = atr.Lng // no lng test
			r.OutPrm(kvsMap)
			r.Addf("r := %v", kvsMap.Make(x.f))
			r.Add("for _, cur := range *x {")
			r.Addf("if grpArr, ok := r[cur%v.%v]; ok {", bse, fld.Name)
			r.Add("grpArr.Push(cur)")
			r.Add("} else {")
			r.Addf("r[cur%v.%v] = %v(cur)", bse, fld.Name, x.Arr.New.Ref(x.f))
			r.Add("}")
			r.Add("}")
			r.Add("return &r")
			if fld.Typ.Bse().IsRel() {
				// TODO: SRTASC, SRTDSC ?
			}
		}
	}
}
