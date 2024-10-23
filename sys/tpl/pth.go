package tpl

import (
	"fmt"
	"strings"
	"sys"
	"sys/k"
	"sys/ks"
)

func GenAnaLits(typ Typ) (r []string) { return GenAna(typ, false) }
func GenAnaVals(typ Typ) (r []string) { return GenAna(typ, true) }
func GenAna(typ Typ, isSys bool) (r []string) {
	bse := typ.Bse()
	b := &strings.Builder{}
	if bse.IsArr() {
		return GenArr(typ.(*Arr), isSys, b)
	}
	switch bse.Camel() {
	case k.Hst, k.Rlt:
		return HstOrRltPth(bse, isSys, b)
	case k.Oan:
		return PrvPth(bse, isSys, b)
	case k.Instr:
		return InstrPths(bse, isSys, b)
	case k.Inrvl:
		return InrvlPths(bse, isSys, b)
	case k.Side:
		return SidePths(bse, isSys, b)
	case k.Stm:
		r = StmRtePths(bse, isSys, b)
		if len(r) > 3 {
			r = r[:3] // reduce for test speed
		}
		return r
	case k.StmBse: //, k.Stm, k.Stms:
		return StmBsePths(bse, isSys, b)
	case k.Stms:
		return StmRtePths(bse, isSys, b)
	case k.StmRte:
		return StmRtePths(bse, isSys, b)
	case k.StmUna:
		return StmUnaPths(bse, isSys, b)
	case k.StmScl:
		return StmSclPths(bse, isSys, b)
	case k.StmSel:
		return StmSelPths(bse, isSys, b)
	case k.StmAgg:
		return StmAggPths(bse, isSys, b)
	case k.StmInr:
		return StmInrPths(bse, isSys, b)
	case k.StmOtr:
		return StmOtrPths(bse, isSys, b)
	case k.CndBse, k.Cnd, k.Cnds, k.CndAnd, k.CndSeq:
		return CndBsePths(bse, isSys, b)
	case k.CndInr:
		return CndInrPths(bse, isSys, b)
	case k.CndOtr:
		return CndOtrPths(bse, isSys, b)
	case k.Stgy, k.StgyBse, k.StgyLong, k.StgyShrt:
		return StgyBsePths(bse, isSys, b)
	case k.StmFbr:
		return []string{"hst.Oan().EurUsd().Is(tmes.New(2, 4, 8)).Bid().Sma()"}
	case k.CndFbr:
		return []string{"hst.Oan().EurUsd().Is(tmes.New(2, 4, 8)).Bid().Sma().InrGtr(1)"}
	}
	// sys.Logf("pth gen: '%v' not implemented (%v)", bse.Camel(), bse.PkgTypTitle())
	// panic(fmt.Sprintf("pth gen: '%v' not implemented (%v)", bse.Camel(), bse.PkgTypTitle()))
	return nil // allow nil for HstStmAggAlma
}
func GenArr(arr *Arr, isSys bool, b *strings.Builder) (r []string) {
	if b == nil {
		b = &strings.Builder{}
	}
	var fst, mdl, lst string
	if isSys {
		// FnOpn(bse, b)
		// b.WriteString(bse.Arr().New.Full())
		b.WriteString(arr.Pkg.Lower()) // manual 'New' to enable InitVals before Fns
		b.WriteRune('.')
		b.WriteString(arr.NewTitle())
		b.WriteRune('(')
		fst = b.String()
		mdl = ","
		lst = ")"
		// lst = ")}"
		b.Reset()
	} else {
		fst, mdl, lst = "[", " ", "]"
	}
	inr := func(vs ...string) string {
		b.Reset()
		b.WriteString(fst)
		for n, v := range vs {
			if n != 0 {
				b.WriteString(mdl)
			}
			b.WriteString(fmt.Sprintf("%v", v))
		}
		b.WriteString(lst)
		return b.String()
	}
	var elms []string
	if isSys {
		elms = arr.Alias.Elm.Bse().Vals
	} else {
		elms = arr.Alias.Elm.Bse().Lits
	}
	if len(elms) > 3 {
		r = append(r, inr(elms[:3]...))
	}
	r = append(r, inr(elms...))
	if arr.IsTstLrg() {
		minSegLen := 64             // flts pll processing
		for len(elms) < minSegLen { // USE ALL UNIQUE; FOR TestFn.TestStmsRltStm
			elms = append(elms, elms...)
		}
		r = append(r, inr(elms...))
	}
	r = append(r, inr()) // put empty at end so first arr may be used by fbr
	return r
}
func GenValsStruct(s *Struct) (r []string) {
	if s == nil {
		return nil
	}
	b := &strings.Builder{}
	wrtCnt := 0
	for _, fld := range s.Flds {
		if !fld.IsLitSkp() && fld.Name != "" {
			if wrtCnt != 0 {
				b.WriteString(", ")
			}
			b.WriteString(fld.Name)
			b.WriteRune(':')
			b.WriteString(fld.Typ.Bse().Val())
			wrtCnt++
		}
	}
	name := s.Full()
	if s.IsPtr() {
		name = "&" + name
	}
	return sys.VsStruct(name, b.String())
}
func GenLitsStruct(s *Struct) (r []string) {
	if s == nil {
		return nil
	}
	b := &strings.Builder{}
	wrtCnt := 0
	for _, fld := range s.Flds {
		if !fld.IsLitSkp() && fld.Name != "" {
			if wrtCnt != 0 {
				b.WriteRune(' ')
			}
			b.WriteString(fld.Camel())
			b.WriteRune(':')
			b.WriteString(fld.Typ.Bse().Lit())
			wrtCnt++
		}
	}
	name := fmt.Sprintf("%v.%v", s.Pkg.Name, s.Camel())
	return sys.VsCtor(name, b.String())
}

func HstOrRltPth(bse *TypBse, isSys bool, b *strings.Builder) (r []string) {
	r = append(r, Itr(bse, b, func() {
		b.WriteString(bse.Pkg.Lower()) // hst, rlt
	}))
	return r
}
func PrvPth(bse *TypBse, isSys bool, b *strings.Builder) (r []string) {
	r = append(r, Itr(bse, b, func() {
		PrvB(k.Oan, isSys, bse, b)
	}))
	return r
}
func InstrPths(bse *TypBse, isSys bool, b *strings.Builder) (r []string) {
	for _, instr := range ks.Instrs {
		r = append(r, Itr(bse, b, func() {
			InstrB(k.Oan, instr, isSys, bse, b)
		}))
	}
	return r
}
func InrvlPths(bse *TypBse, isSys bool, b *strings.Builder) (r []string) {
	for _, instr := range ks.Instrs {
		for _, inrvl := range Inrvls {
			r = append(r, Itr(bse, b, func() {
				InrvlB(k.Oan, instr, inrvl, isSys, bse, b)
			}))
		}
	}
	return r
}
func SidePths(bse *TypBse, isSys bool, b *strings.Builder) (r []string) {
	for _, instr := range ks.Instrs {
		for _, inrvl := range Inrvls {
			for _, side := range Sides {
				r = append(r, Itr(bse, b, func() {
					SideB(k.Oan, instr, inrvl, side, isSys, bse, b)
				}))
			}
		}
	}
	return r
}
func StmRtePths(bse *TypBse, isSys bool, b *strings.Builder) (r []string) {
	for _, instr := range ks.Instrs {
		for _, inrvl := range Inrvls {
			for _, side := range Sides {
				for _, stmRte := range ks.StmRtes {
					r = append(r, Itr(bse, b, func() {
						StmRteB(k.Oan, instr, inrvl, side, stmRte, isSys, bse, b)
					}))
				}
			}
		}
	}
	return r
}
func StmBsePths(bse *TypBse, isSys bool, b *strings.Builder) (r []string) {
	for _, instr := range ks.Instrs {
		for _, inrvl := range Inrvls {
			for _, side := range Sides {
				for _, stmRte := range StmRtes {
					r = append(r, Itr(bse, b, func() {
						StmRteB(k.Oan, instr, inrvl, side, stmRte, isSys, bse, b)
					}))
				}
			}
		}
	}
	return r
}
func StmUnaPths(bse *TypBse, isSys bool, b *strings.Builder) (r []string) {
	for _, instr := range ks.Instrs {
		for _, inrvl := range Inrvls {
			for _, side := range Sides {
				for _, stmRte := range StmRtes {
					for _, stmUna := range ks.StmUnas {
						r = append(r, Itr(bse, b, func() {
							StmUnaB(k.Oan, instr, inrvl, side, stmRte, stmUna, isSys, bse, b)
						}))
					}
				}
			}
		}
	}
	return r
}
func StmSclPths(bse *TypBse, isSys bool, b *strings.Builder) (r []string) {
	for _, instr := range ks.Instrs {
		for _, inrvl := range Inrvls {
			for _, side := range Sides {
				for _, stmRte := range StmRtes {
					for _, stmScl := range ks.StmScls {
						r = append(r, Itr(bse, b, func() {
							StmSclB(k.Oan, instr, inrvl, side, stmRte, stmScl, "4.0", isSys, bse, b)
						}))
					}
				}
			}
		}
	}
	return r
}
func StmSelPths(bse *TypBse, isSys bool, b *strings.Builder) (r []string) {
	for _, instr := range ks.Instrs {
		for _, inrvl := range Inrvls {
			for _, side := range Sides {
				for _, stmRte := range StmRtes {
					for _, stmSel := range ks.StmSels {
						r = append(r, Itr(bse, b, func() {
							StmSelB(k.Oan, instr, inrvl, side, stmRte, stmSel, "1.3171", isSys, bse, b) // 1.3171 common in tst data
						}))
					}
				}
			}
		}
	}
	return r
}
func StmAggPths(bse *TypBse, isSys bool, b *strings.Builder) (r []string) {
	for _, instr := range ks.Instrs {
		for _, inrvl := range Inrvls {
			for _, side := range Sides {
				for _, stmRte := range StmRtes {
					for _, stmAgg := range ks.StmAggs {
						for _, length := range []string{"10", "100", "1000"} {
							r = append(r, Itr(bse, b, func() {
								StmAggB(k.Oan, instr, inrvl, side, stmRte, stmAgg, length, isSys, bse, b)
							}))
						}
					}
				}
			}
		}
	}
	return r
}
func StmInrPths(bse *TypBse, isSys bool, b *strings.Builder) (r []string) {
	for _, instr := range ks.Instrs {
		for _, inrvl := range Inrvls {
			for _, side := range Sides {
				for _, stmRte := range StmRtes {
					for _, stmInr := range ks.StmInrs {
						for _, off := range []string{"10", "100", "1000"} {
							r = append(r, Itr(bse, b, func() {
								StmInrB(k.Oan, instr, inrvl, side, stmRte, stmInr, off, isSys, bse, b)
							}))
						}
					}
				}
			}
		}
	}
	return r
}
func StmOtrPths(bse *TypBse, isSys bool, b *strings.Builder) (r []string) {
	StmRteB(k.Oan, ks.Instrs[0], Inrvls[0], ks.Sides[0], ks.StmRtes[0], isSys, bse, b)
	stmA := b.String()
	for _, instr := range ks.Instrs[:1] {
		for _, inrvl := range Inrvls {
			for _, side := range Sides {
				for _, stmRte := range StmRtes {
					for _, stmOtr := range ks.StmOtrs {
						for _, off := range []string{"0", "2"} {
							r = append(r, Itr(bse, b, func() {
								StmOtrB(k.Oan, instr, inrvl, side, stmRte, stmOtr, off, stmA, isSys, bse, b)
							}))
						}
						if !*Long {
							return r
						}
					}
				}
			}
		}
	}
	return r
}
func CndBsePths(bse *TypBse, isSys bool, b *strings.Builder) (r []string) {
	for _, instr := range ks.Instrs {
		for _, inrvl := range Inrvls {
			for _, side := range Sides {
				for _, stmRte := range StmRtes {
					r = append(r, Itr(bse, b, func() {
						CndInrB(k.Oan, instr, inrvl, side, stmRte, ks.CndInrs[2], "1", isSys, bse, b)
					}))
				}
			}
		}
	}
	return r
}
func CndInrPths(bse *TypBse, isSys bool, b *strings.Builder) (r []string) {
	for _, instr := range ks.Instrs {
		for _, inrvl := range Inrvls {
			for _, side := range Sides {
				for _, stmRte := range StmRtes {
					for _, cndInr := range ks.CndInrs {
						for _, off := range []string{"10", "100", "1000"} {
							r = append(r, Itr(bse, b, func() {
								CndInrB(k.Oan, instr, inrvl, side, stmRte, cndInr, off, isSys, bse, b)
							}))
						}
					}
				}
			}
		}
	}
	return r
}
func CndOtrPths(bse *TypBse, isSys bool, b *strings.Builder) (r []string) {
	StmRteB(k.Oan, ks.Instrs[0], Inrvls[0], ks.Sides[0], ks.StmRtes[0], isSys, bse, b)
	stmA := b.String()
	for _, instr := range ks.Instrs[:1] {
		for _, inrvl := range Inrvls {
			for _, side := range Sides {
				for _, stmRte := range StmRtes {
					for _, cndOtr := range ks.CndOtrs {
						for _, off := range []string{"0", "2"} {
							r = append(r, Itr(bse, b, func() {
								CndOtrB(k.Oan, instr, inrvl, side, stmRte, cndOtr, off, stmA, isSys, bse, b)
							}))
						}
						if !*Long {
							return r
						}
					}
				}
			}
		}
	}
	return r
}
func StgyBsePths(bse *TypBse, isSys bool, b *strings.Builder) (r []string) {
	var wnd string
	if isSys {
		wnd = "tme.Tme(3 * 60 * 60)"
	} else {
		wnd = "3h"
	}
	for _, instr := range ks.Instrs {
		b2 := &strings.Builder{}
		InstrB(k.Oan, instr, isSys, bse, b2)
		instr2 := b2.String()
		for _, inrvl := range Inrvls {
			for _, side := range Sides {
				for _, stmRte := range StmRtes {
					for _, stgy := range ks.Stgys {
						r = append(r, Itr(bse, b, func() {
							StgyB(k.Oan, instr, inrvl, side, stmRte, ks.CndInrs[2], "1", stgy, "1.0", "1.0", wnd, instr2, isSys, bse, b)
						}))
					}
				}
			}
		}
	}
	return r
}

func Itr(bse *TypBse, b *strings.Builder, f func()) string {
	b.Reset()
	f()
	return b.String()
}
func PkgB(bse *TypBse, b *strings.Builder) {
	b.WriteString(bse.Pkg.Lower()) // hst, rlt
}
func PrvB(prv string, isSys bool, bse *TypBse, b *strings.Builder) {
	PkgB(bse, b)
	b.WriteRune('.') // hst.Oan()
	if isSys {
		prv = strings.Title(prv)
	}
	b.WriteString(prv)
	b.WriteString("()")
}
func InstrB(prv, instr string, isSys bool, bse *TypBse, b *strings.Builder) {
	PrvB(prv, isSys, bse, b)
	b.WriteRune('.') // hst.Oan().EurUsd()
	if isSys {
		instr = strings.Title(instr)
	}
	b.WriteString(instr)
	b.WriteString("()")
}
func InrvlB(prv, instr, inrvl string, isSys bool, bse *TypBse, b *strings.Builder) {
	InstrB(prv, instr, isSys, bse, b)
	b.WriteRune('.') // hst.Oan().EurUsd().I(5)
	name := k.I
	if isSys {
		name = strings.Title(name)
	}
	b.WriteString(name)
	b.WriteString("(")
	b.WriteString(inrvl)
	b.WriteString(")")
}
func SideB(prv, instr, inrvl, side string, isSys bool, bse *TypBse, b *strings.Builder) {
	InrvlB(prv, instr, inrvl, isSys, bse, b)
	b.WriteRune('.') // hst.Oan().EurUsd().I(5).Bid()
	if isSys {
		side = strings.Title(side)
	}
	b.WriteString(side)
	b.WriteString("()")
}
func StmRteB(prv, instr, inrvl, side, stmRte string, isSys bool, bse *TypBse, b *strings.Builder) {
	SideB(prv, instr, inrvl, side, isSys, bse, b)
	b.WriteRune('.') // hst.Oan().EurUsd().I(5).Bid().Lst()
	if isSys {
		stmRte = strings.Title(stmRte)
	}
	b.WriteString(stmRte)
	b.WriteString("()")
}
func StmUnaB(prv, instr, inrvl, side, stmRte, stmUna string, isSys bool, bse *TypBse, b *strings.Builder) {
	StmRteB(prv, instr, inrvl, side, stmRte, isSys, bse, b)
	b.WriteRune('.') // hst.Oan().EurUsd().I(5).Bid().Lst().Pos()
	if isSys {
		stmUna = strings.Title(stmUna)
	}
	b.WriteString(stmUna)
	b.WriteString("()")
}
func StmSclB(prv, instr, inrvl, side, stmRte, stmScl, scl string, isSys bool, bse *TypBse, b *strings.Builder) {
	StmRteB(prv, instr, inrvl, side, stmRte, isSys, bse, b)
	b.WriteRune('.') // hst.Oan().EurUsd().I(5).Bid().Lst().SclMul(10000.0)
	if isSys {
		stmScl = strings.Title(stmScl)
	}
	b.WriteString(stmScl)
	b.WriteRune('(')
	b.WriteString(scl)
	b.WriteRune(')')
}
func StmSelB(prv, instr, inrvl, side, stmRte, stmSel, sel string, isSys bool, bse *TypBse, b *strings.Builder) {
	StmRteB(prv, instr, inrvl, side, stmRte, isSys, bse, b)
	b.WriteRune('.') // hst.Oan().EurUsd().I(5).Bid().Lst().SelGtr(0.0)
	if isSys {
		stmSel = strings.Title(stmSel)
	}
	b.WriteString(stmSel)
	b.WriteRune('(')
	b.WriteString(sel)
	b.WriteRune(')')
}
func StmAggB(prv, instr, inrvl, side, stmRte, stmAgg, len string, isSys bool, bse *TypBse, b *strings.Builder) {
	StmRteB(prv, instr, inrvl, side, stmRte, isSys, bse, b)
	b.WriteRune('.') // hst.Oan().EurUsd().I(5).Bid().Lst().AggSum(10)
	if isSys {
		stmAgg = strings.Title(stmAgg)
	}
	b.WriteString(stmAgg)
	b.WriteRune('(')
	b.WriteString(len)
	b.WriteRune(')')
}
func StmInrB(prv, instr, inrvl, side, stmRte, stmInr, off string, isSys bool, bse *TypBse, b *strings.Builder) {
	StmRteB(prv, instr, inrvl, side, stmRte, isSys, bse, b)
	b.WriteRune('.') // hst.Oan().EurUsd().I(5).Bid().Lst().InrSub(1)
	if isSys {
		stmInr = strings.Title(stmInr)
	}
	b.WriteString(stmInr)
	b.WriteRune('(')
	b.WriteString(off)
	b.WriteRune(')')
}
func StmOtrB(prv, instr, inrvl, side, stmRte, stmOtr, off, stmA string, isSys bool, bse *TypBse, b *strings.Builder) {
	StmRteB(prv, instr, inrvl, side, stmRte, isSys, bse, b)
	b.WriteRune('.') // hst.Oan().EurUsd().I(5).Bid().Lst().OtrSub(0, hst.Oan().EurUsd().I(5).Bid().Lst())
	delim := ' '
	if isSys {
		stmOtr = strings.Title(stmOtr)
		delim = ','
	}
	b.WriteString(stmOtr)
	b.WriteRune('(')
	b.WriteString(off)
	b.WriteRune(delim)
	b.WriteString(stmA)
	b.WriteRune(')')
}
func CndInrB(prv, instr, inrvl, side, stmRte, cndInr, off string, isSys bool, bse *TypBse, b *strings.Builder) {
	cndInr = sys.CnjInr(cndInr)
	StmRteB(prv, instr, inrvl, side, stmRte, isSys, bse, b)
	b.WriteRune('.') // hst.Oan().EurUsd().I(5).Bid().Lst().InrEql(1)
	if isSys {
		cndInr = strings.Title(cndInr)
	}
	b.WriteString(cndInr)
	b.WriteRune('(')
	b.WriteString(off)
	b.WriteRune(')')
}
func CndOtrB(prv, instr, inrvl, side, stmRte, cndOtr, off, stmA string, isSys bool, bse *TypBse, b *strings.Builder) {
	StmRteB(prv, instr, inrvl, side, stmRte, isSys, bse, b)
	b.WriteRune('.') // hst.Oan().EurUsd().I(5).Bid().Lst().OtrEql(0, hst.Oan().EurUsd().I(5).Bid().Lst())
	delim := ' '
	if isSys {
		cndOtr = strings.Title(cndOtr)
		delim = ','
	}
	b.WriteString(cndOtr)
	b.WriteRune('(')
	b.WriteString(off)
	b.WriteRune(delim)
	b.WriteString(stmA)
	b.WriteRune(')')
}
func StgyB(prv, instr, inrvl, side, stmRte, cndInr, off, stgy, prfLim, losLim, wnd, instr2 string, isSys bool, bse *TypBse, b *strings.Builder) {
	CndInrB(prv, instr, inrvl, side, stmRte, cndInr, off, isSys, bse, b)
	b.WriteRune('.') // hst.Oan().EurUsd().I(5).Bid().Lst().InrEql(1).Long(1.0, 1.0, 1h, hst.Oan().EurUsd())
	delim := ' '
	if isSys {
		stgy = strings.Title(stgy)
		delim = ','
	}
	b.WriteString(stgy)
	b.WriteRune('(')
	b.WriteString(prfLim)
	b.WriteRune(delim)
	b.WriteString(losLim)
	b.WriteRune(delim)
	b.WriteString(wnd)
	b.WriteRune(delim)
	b.WriteString(instr2)
	b.WriteRune(')')
}

var (
	Inrvls  = []string{"5"}
	Sides   = ks.Sides[:1]
	StmRtes = ks.StmRtes[1:2] // .lst()
)
