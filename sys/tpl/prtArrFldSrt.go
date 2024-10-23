package tpl

import (
	"strings"
)

type (
	PrtArrFldSrt struct {
		PrtArrSrt
		Arr    *Arr
		Elm    *Struct
		Ifc    *Ifc
		ElmRel *PrtStructRel
	}
)

func (x *PrtArrFldSrt) InitPrtTyp() {
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
func (x *PrtArrFldSrt) InitPrtTypFn() {
	if x.Elm != nil {
		// var bse string
		// if x.Ifc != nil {
		// 	bse = ".Bse()"
		// }
		flds0 := NewFlds(x.Elm.Flds...)
		flds := NewFlds()
		for flds0.Cnt() != 0 {
			fld := flds0.Pop()
			if fld.Name == "" { // embedded typ
				if s, ok := fld.Typ.(*Struct); ok {
					flds0.Push(s.Flds...) // add embedded typ fields
				}
				continue
			}
			fldBse := fld.Typ.Bse()
			if fld.IsFstUpr() && fld.IsSrt() && fldBse.IsRel() && fldBse.IsAct() {
				flds.Push(fld)
			}
		}
		for _, fld1 := range *flds {
			eql := x.ElmRel.Eqls[fld1.Name]
			lss := x.ElmRel.Lsss[fld1.Name]
			gtr := x.ElmRel.Gtrs[fld1.Name]
			// tst
			// if x.f.Tst != nil {
			// 	x.f.Tst.ArrRelPrt(lss, gtr, fld1.Name)
			// }
			_sys.Tst.Srt(x.f.Typ(), lss, fld1)
			_sys.Tst.Srt(x.f.Typ(), gtr, fld1)
			srtAsc := x.SrtAsc(lss, eql, fld1.Name)
			srtDsc := x.SrtDsc(gtr, eql, fld1.Name)
			srts := []*TypFn{srtAsc, srtDsc}
			for _, fld2 := range *flds {
				if fld1 != fld2 {
					op2s := []*PkgFn{x.ElmRel.Lsss[fld2.Name], x.ElmRel.Gtrs[fld2.Name]}
					for _, srt := range srts {
						for _, op2 := range op2s {
							x.Srt2(srt, fld1, fld2, op2, x.ElmRel)
							// for _, fld3 := range *flds { // SO SLOW
							// 	if fld1 != fld3 && fld2 != fld3 && fld3.IsSrt() {
							// 		op3s := []*PkgFn{elmStructRel.Lsss[fld3.Name], elmStructRel.Gtrs[fld3.Name]}
							// 		for _, op3 := range op3s {
							// 			x.Srt3(srt, fld1, fld2, fld3, op2, op3, elmStructRel)
							// 		}
							// 	}
							// }
						}
					}

				}
			}
		}

		x.SrtQuick(x.ElmRel.Cmp)
		x.SrtIns(x.ElmRel.Cmp, nil, nil)
		x.SrtMdnOf3(x.ElmRel.Cmp, nil)
		x.Swp()
	}
}

func (x *PrtArrFldSrt) Srt2(srt1 *TypFn, fld1, fld2 *Fld, op2 *PkgFn, p *PrtStructRel) (r *TypFn) {
	r = x.f.TypFnf("%v%v%v", srt1.Name, strings.Title(op2.Alias), strings.Title(fld2.Name))
	r.OutPrm(x.f.Typ())
	r.Add("if len(*x) > 1 {")
	r.Addf("x.%v()", srt1.Name)
	r.Add("var strt int")
	r.Add("for cur := strt + 1; cur < len(*x); cur++ {")
	r.Addf("if !%v((*x)[cur-1], (*x)[cur]) {", p.Eqls[fld1.Name].Ref(x.f))
	r.Add("if cur-strt > 1 {")
	r.Addf("x.SrtQuick(unt.Unt(strt), unt.Unt(cur-1), %v, %v)", op2.Ref(x.f), p.Eqls[fld2.Name].Ref(x.f))
	r.Add("}")
	r.Add("strt = cur")
	r.Add("}")
	r.Add("}")
	r.Add("}")
	r.Add("return x")
	return r
}

// func (x *PrtArrFldSrt) Srt3(srt1 *TypFn, fld1, fld2, fld3 *Fld, op2, op3 *PkgFn, p *PrtStructRel) (r *TypFn) {
// 	r = x.f.TypFnf("%v%v%v%v%v", srt1.Name, op2.Alias, fld2.Name, op3.Alias, fld3.Name)
// 	r.OutPrm(x.f.Typ())
// 	r.Add("if len(*x) > 1 {")
// 	r.Addf("x.%v()", srt1.Name)
// 	r.Add("var strt int")
// 	r.Add("for cur2 := strt + 1; cur2 < len(*x); cur2++ {")
// 	r.Addf("if !%v((*x)[cur2-1], (*x)[cur2]) {", p.Eqls[fld1.Name].Ref(x.f))
// 	r.Add("if cur2-strt > 1 {")
// 	r.Addf("x.SrtQuick(unt.Unt(strt), unt.Unt(cur2-1), %v)", op2.Ref(x.f))
// 	r.Add("}")
// 	r.Add("strt = cur2")
// 	r.Add("}")
// 	r.Add("}")
// 	r.Add("}")
// 	r.Add("return x")
// 	return r
// }
