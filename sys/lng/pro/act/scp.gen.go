package act

import (
	"fmt"
	"sys"
	"sys/ana"
	"sys/ana/hst"
	"sys/ana/rlt"
	"sys/ana/vis/clr"
	"sys/ana/vis/fnt"
	"sys/ana/vis/pen"
	"sys/ana/vis/plt"
	"sys/bsc/bnd"
	"sys/bsc/bnds"
	"sys/bsc/bol"
	"sys/bsc/bols"
	"sys/bsc/flt"
	"sys/bsc/flts"
	"sys/bsc/int"
	"sys/bsc/ints"
	"sys/bsc/str"
	"sys/bsc/strs"
	"sys/bsc/tme"
	"sys/bsc/tmes"
	"sys/bsc/unt"
	"sys/bsc/unts"
	"sys/lng/pro/xpr"
	"sys/lng/pro/xpr/knd"
)

var (
	FnRetIdn = string("!fnRetIdn!")
)

type (
	Scp struct {
		Xpr             *xpr.Scp
		Prnt            *Scp
		IdxRet          uint32
		strStr          []str.Str
		bolBol          []bol.Bol
		fltFlt          []flt.Flt
		untUnt          []unt.Unt
		intInt          []int.Int
		tmeTme          []tme.Tme
		bndBnd          []bnd.Bnd
		fltRng          []flt.Rng
		tmeRng          []tme.Rng
		strsStrs        []*strs.Strs
		bolsBols        []*bols.Bols
		fltsFlts        []*flts.Flts
		untsUnts        []*unts.Unts
		intsInts        []*ints.Ints
		tmesTmes        []*tmes.Tmes
		bndsBnds        []*bnds.Bnds
		tmeRngs         []*tme.Rngs
		anaTrd          []*ana.Trd
		anaTrds         []*ana.Trds
		anaPrfm         []*ana.Prfm
		anaPrfms        []*ana.Prfms
		anaPrfmDlt      []*ana.PrfmDlt
		anaPort         []*ana.Port
		hstPrv          []hst.Prv
		hstInstr        []hst.Instr
		hstInrvl        []hst.Inrvl
		hstSide         []hst.Side
		hstStm          []hst.Stm
		hstCnd          []hst.Cnd
		hstStgy         []hst.Stgy
		hstPrvs         []*hst.Prvs
		hstInstrs       []*hst.Instrs
		hstInrvls       []*hst.Inrvls
		hstSides        []*hst.Sides
		hstStms         []*hst.Stms
		hstCnds         []*hst.Cnds
		hstStgys        []*hst.Stgys
		rltPrv          []rlt.Prv
		rltInstr        []rlt.Instr
		rltInrvl        []rlt.Inrvl
		rltSide         []rlt.Side
		rltStm          []rlt.Stm
		rltCnd          []rlt.Cnd
		rltStgy         []rlt.Stgy
		rltPrvs         []*rlt.Prvs
		rltInstrs       []*rlt.Instrs
		rltInrvls       []*rlt.Inrvls
		rltSides        []*rlt.Sides
		rltStms         []*rlt.Stms
		rltCnds         []*rlt.Cnds
		rltStgys        []*rlt.Stgys
		fntFnt          []*fnt.Fnt
		clrClr          []clr.Clr
		penPen          []pen.Pen
		penPens         []*pen.Pens
		pltPlt          []plt.Plt
		pltPlts         []*plt.Plts
		pltStm          []*plt.Stm
		pltFltsSctr     []*plt.FltsSctr
		pltFltsSctrDist []*plt.FltsSctrDist
		pltHrz          []*plt.Hrz
		pltVrt          []*plt.Vrt
		pltDpth         []*plt.Dpth
		sysMu           []*sys.Mu
	}
)

func NewScp(xprScp *xpr.Scp, prnt ...*Scp) (r *Scp) {
	r = &Scp{}
	r.Xpr = xprScp
	if len(prnt) != 0 {
		r.Prnt = prnt[0]
	}
	r.IdxRet = xprScp.Vars[FnRetIdn].Idx
	for xprKnd, cnt := range xprScp.Cnts { // allocate slices of exact count
		switch xprKnd {
		case knd.StrStr:
			r.strStr = make([]str.Str, cnt)
		case knd.BolBol:
			r.bolBol = make([]bol.Bol, cnt)
		case knd.FltFlt:
			r.fltFlt = make([]flt.Flt, cnt)
		case knd.UntUnt:
			r.untUnt = make([]unt.Unt, cnt)
		case knd.IntInt:
			r.intInt = make([]int.Int, cnt)
		case knd.TmeTme:
			r.tmeTme = make([]tme.Tme, cnt)
		case knd.BndBnd:
			r.bndBnd = make([]bnd.Bnd, cnt)
		case knd.FltRng:
			r.fltRng = make([]flt.Rng, cnt)
		case knd.TmeRng:
			r.tmeRng = make([]tme.Rng, cnt)
		case knd.StrsStrs:
			r.strsStrs = make([]*strs.Strs, cnt)
		case knd.BolsBols:
			r.bolsBols = make([]*bols.Bols, cnt)
		case knd.FltsFlts:
			r.fltsFlts = make([]*flts.Flts, cnt)
		case knd.UntsUnts:
			r.untsUnts = make([]*unts.Unts, cnt)
		case knd.IntsInts:
			r.intsInts = make([]*ints.Ints, cnt)
		case knd.TmesTmes:
			r.tmesTmes = make([]*tmes.Tmes, cnt)
		case knd.BndsBnds:
			r.bndsBnds = make([]*bnds.Bnds, cnt)
		case knd.TmeRngs:
			r.tmeRngs = make([]*tme.Rngs, cnt)
		case knd.AnaTrd:
			r.anaTrd = make([]*ana.Trd, cnt)
		case knd.AnaTrds:
			r.anaTrds = make([]*ana.Trds, cnt)
		case knd.AnaPrfm:
			r.anaPrfm = make([]*ana.Prfm, cnt)
		case knd.AnaPrfms:
			r.anaPrfms = make([]*ana.Prfms, cnt)
		case knd.AnaPrfmDlt:
			r.anaPrfmDlt = make([]*ana.PrfmDlt, cnt)
		case knd.AnaPort:
			r.anaPort = make([]*ana.Port, cnt)
		case knd.HstPrv:
			r.hstPrv = make([]hst.Prv, cnt)
		case knd.HstInstr:
			r.hstInstr = make([]hst.Instr, cnt)
		case knd.HstInrvl:
			r.hstInrvl = make([]hst.Inrvl, cnt)
		case knd.HstSide:
			r.hstSide = make([]hst.Side, cnt)
		case knd.HstStm:
			r.hstStm = make([]hst.Stm, cnt)
		case knd.HstCnd:
			r.hstCnd = make([]hst.Cnd, cnt)
		case knd.HstStgy:
			r.hstStgy = make([]hst.Stgy, cnt)
		case knd.HstPrvs:
			r.hstPrvs = make([]*hst.Prvs, cnt)
		case knd.HstInstrs:
			r.hstInstrs = make([]*hst.Instrs, cnt)
		case knd.HstInrvls:
			r.hstInrvls = make([]*hst.Inrvls, cnt)
		case knd.HstSides:
			r.hstSides = make([]*hst.Sides, cnt)
		case knd.HstStms:
			r.hstStms = make([]*hst.Stms, cnt)
		case knd.HstCnds:
			r.hstCnds = make([]*hst.Cnds, cnt)
		case knd.HstStgys:
			r.hstStgys = make([]*hst.Stgys, cnt)
		case knd.RltPrv:
			r.rltPrv = make([]rlt.Prv, cnt)
		case knd.RltInstr:
			r.rltInstr = make([]rlt.Instr, cnt)
		case knd.RltInrvl:
			r.rltInrvl = make([]rlt.Inrvl, cnt)
		case knd.RltSide:
			r.rltSide = make([]rlt.Side, cnt)
		case knd.RltStm:
			r.rltStm = make([]rlt.Stm, cnt)
		case knd.RltCnd:
			r.rltCnd = make([]rlt.Cnd, cnt)
		case knd.RltStgy:
			r.rltStgy = make([]rlt.Stgy, cnt)
		case knd.RltPrvs:
			r.rltPrvs = make([]*rlt.Prvs, cnt)
		case knd.RltInstrs:
			r.rltInstrs = make([]*rlt.Instrs, cnt)
		case knd.RltInrvls:
			r.rltInrvls = make([]*rlt.Inrvls, cnt)
		case knd.RltSides:
			r.rltSides = make([]*rlt.Sides, cnt)
		case knd.RltStms:
			r.rltStms = make([]*rlt.Stms, cnt)
		case knd.RltCnds:
			r.rltCnds = make([]*rlt.Cnds, cnt)
		case knd.RltStgys:
			r.rltStgys = make([]*rlt.Stgys, cnt)
		case knd.FntFnt:
			r.fntFnt = make([]*fnt.Fnt, cnt)
		case knd.ClrClr:
			r.clrClr = make([]clr.Clr, cnt)
		case knd.PenPen:
			r.penPen = make([]pen.Pen, cnt)
		case knd.PenPens:
			r.penPens = make([]*pen.Pens, cnt)
		case knd.PltPlt:
			r.pltPlt = make([]plt.Plt, cnt)
		case knd.PltPlts:
			r.pltPlts = make([]*plt.Plts, cnt)
		case knd.PltStm:
			r.pltStm = make([]*plt.Stm, cnt)
		case knd.PltFltsSctr:
			r.pltFltsSctr = make([]*plt.FltsSctr, cnt)
		case knd.PltFltsSctrDist:
			r.pltFltsSctrDist = make([]*plt.FltsSctrDist, cnt)
		case knd.PltHrz:
			r.pltHrz = make([]*plt.Hrz, cnt)
		case knd.PltVrt:
			r.pltVrt = make([]*plt.Vrt, cnt)
		case knd.PltDpth:
			r.pltDpth = make([]*plt.Dpth, cnt)
		case knd.SysMu:
			r.sysMu = make([]*sys.Mu, cnt)
		}
	}
	return r
}
func (x *Scp) StrStr(idn string) (r str.StrScp) {
	cur := x
	for cur != nil {
		kndIdx, exists := cur.Xpr.Vars[idn]
		if exists {
			r.Idx = kndIdx.Idx
			r.Arr = cur.strStr
			return r
		}
		cur = cur.Prnt
	}
	panic(fmt.Sprintf("scope missing idn '%v'", idn))
}
func (x *Scp) BolBol(idn string) (r bol.BolScp) {
	cur := x
	for cur != nil {
		kndIdx, exists := cur.Xpr.Vars[idn]
		if exists {
			r.Idx = kndIdx.Idx
			r.Arr = cur.bolBol
			return r
		}
		cur = cur.Prnt
	}
	panic(fmt.Sprintf("scope missing idn '%v'", idn))
}
func (x *Scp) FltFlt(idn string) (r flt.FltScp) {
	cur := x
	for cur != nil {
		kndIdx, exists := cur.Xpr.Vars[idn]
		if exists {
			r.Idx = kndIdx.Idx
			r.Arr = cur.fltFlt
			return r
		}
		cur = cur.Prnt
	}
	panic(fmt.Sprintf("scope missing idn '%v'", idn))
}
func (x *Scp) UntUnt(idn string) (r unt.UntScp) {
	cur := x
	for cur != nil {
		kndIdx, exists := cur.Xpr.Vars[idn]
		if exists {
			r.Idx = kndIdx.Idx
			r.Arr = cur.untUnt
			return r
		}
		cur = cur.Prnt
	}
	panic(fmt.Sprintf("scope missing idn '%v'", idn))
}
func (x *Scp) IntInt(idn string) (r int.IntScp) {
	cur := x
	for cur != nil {
		kndIdx, exists := cur.Xpr.Vars[idn]
		if exists {
			r.Idx = kndIdx.Idx
			r.Arr = cur.intInt
			return r
		}
		cur = cur.Prnt
	}
	panic(fmt.Sprintf("scope missing idn '%v'", idn))
}
func (x *Scp) TmeTme(idn string) (r tme.TmeScp) {
	cur := x
	for cur != nil {
		kndIdx, exists := cur.Xpr.Vars[idn]
		if exists {
			r.Idx = kndIdx.Idx
			r.Arr = cur.tmeTme
			return r
		}
		cur = cur.Prnt
	}
	panic(fmt.Sprintf("scope missing idn '%v'", idn))
}
func (x *Scp) BndBnd(idn string) (r bnd.BndScp) {
	cur := x
	for cur != nil {
		kndIdx, exists := cur.Xpr.Vars[idn]
		if exists {
			r.Idx = kndIdx.Idx
			r.Arr = cur.bndBnd
			return r
		}
		cur = cur.Prnt
	}
	panic(fmt.Sprintf("scope missing idn '%v'", idn))
}
func (x *Scp) FltRng(idn string) (r flt.RngScp) {
	cur := x
	for cur != nil {
		kndIdx, exists := cur.Xpr.Vars[idn]
		if exists {
			r.Idx = kndIdx.Idx
			r.Arr = cur.fltRng
			return r
		}
		cur = cur.Prnt
	}
	panic(fmt.Sprintf("scope missing idn '%v'", idn))
}
func (x *Scp) TmeRng(idn string) (r tme.RngScp) {
	cur := x
	for cur != nil {
		kndIdx, exists := cur.Xpr.Vars[idn]
		if exists {
			r.Idx = kndIdx.Idx
			r.Arr = cur.tmeRng
			return r
		}
		cur = cur.Prnt
	}
	panic(fmt.Sprintf("scope missing idn '%v'", idn))
}
func (x *Scp) StrsStrs(idn string) (r strs.StrsScp) {
	cur := x
	for cur != nil {
		kndIdx, exists := cur.Xpr.Vars[idn]
		if exists {
			r.Idx = kndIdx.Idx
			r.Arr = cur.strsStrs
			return r
		}
		cur = cur.Prnt
	}
	panic(fmt.Sprintf("scope missing idn '%v'", idn))
}
func (x *Scp) BolsBols(idn string) (r bols.BolsScp) {
	cur := x
	for cur != nil {
		kndIdx, exists := cur.Xpr.Vars[idn]
		if exists {
			r.Idx = kndIdx.Idx
			r.Arr = cur.bolsBols
			return r
		}
		cur = cur.Prnt
	}
	panic(fmt.Sprintf("scope missing idn '%v'", idn))
}
func (x *Scp) FltsFlts(idn string) (r flts.FltsScp) {
	cur := x
	for cur != nil {
		kndIdx, exists := cur.Xpr.Vars[idn]
		if exists {
			r.Idx = kndIdx.Idx
			r.Arr = cur.fltsFlts
			return r
		}
		cur = cur.Prnt
	}
	panic(fmt.Sprintf("scope missing idn '%v'", idn))
}
func (x *Scp) UntsUnts(idn string) (r unts.UntsScp) {
	cur := x
	for cur != nil {
		kndIdx, exists := cur.Xpr.Vars[idn]
		if exists {
			r.Idx = kndIdx.Idx
			r.Arr = cur.untsUnts
			return r
		}
		cur = cur.Prnt
	}
	panic(fmt.Sprintf("scope missing idn '%v'", idn))
}
func (x *Scp) IntsInts(idn string) (r ints.IntsScp) {
	cur := x
	for cur != nil {
		kndIdx, exists := cur.Xpr.Vars[idn]
		if exists {
			r.Idx = kndIdx.Idx
			r.Arr = cur.intsInts
			return r
		}
		cur = cur.Prnt
	}
	panic(fmt.Sprintf("scope missing idn '%v'", idn))
}
func (x *Scp) TmesTmes(idn string) (r tmes.TmesScp) {
	cur := x
	for cur != nil {
		kndIdx, exists := cur.Xpr.Vars[idn]
		if exists {
			r.Idx = kndIdx.Idx
			r.Arr = cur.tmesTmes
			return r
		}
		cur = cur.Prnt
	}
	panic(fmt.Sprintf("scope missing idn '%v'", idn))
}
func (x *Scp) BndsBnds(idn string) (r bnds.BndsScp) {
	cur := x
	for cur != nil {
		kndIdx, exists := cur.Xpr.Vars[idn]
		if exists {
			r.Idx = kndIdx.Idx
			r.Arr = cur.bndsBnds
			return r
		}
		cur = cur.Prnt
	}
	panic(fmt.Sprintf("scope missing idn '%v'", idn))
}
func (x *Scp) TmeRngs(idn string) (r tme.RngsScp) {
	cur := x
	for cur != nil {
		kndIdx, exists := cur.Xpr.Vars[idn]
		if exists {
			r.Idx = kndIdx.Idx
			r.Arr = cur.tmeRngs
			return r
		}
		cur = cur.Prnt
	}
	panic(fmt.Sprintf("scope missing idn '%v'", idn))
}
func (x *Scp) AnaTrd(idn string) (r ana.TrdScp) {
	cur := x
	for cur != nil {
		kndIdx, exists := cur.Xpr.Vars[idn]
		if exists {
			r.Idx = kndIdx.Idx
			r.Arr = cur.anaTrd
			return r
		}
		cur = cur.Prnt
	}
	panic(fmt.Sprintf("scope missing idn '%v'", idn))
}
func (x *Scp) AnaTrds(idn string) (r ana.TrdsScp) {
	cur := x
	for cur != nil {
		kndIdx, exists := cur.Xpr.Vars[idn]
		if exists {
			r.Idx = kndIdx.Idx
			r.Arr = cur.anaTrds
			return r
		}
		cur = cur.Prnt
	}
	panic(fmt.Sprintf("scope missing idn '%v'", idn))
}
func (x *Scp) AnaPrfm(idn string) (r ana.PrfmScp) {
	cur := x
	for cur != nil {
		kndIdx, exists := cur.Xpr.Vars[idn]
		if exists {
			r.Idx = kndIdx.Idx
			r.Arr = cur.anaPrfm
			return r
		}
		cur = cur.Prnt
	}
	panic(fmt.Sprintf("scope missing idn '%v'", idn))
}
func (x *Scp) AnaPrfms(idn string) (r ana.PrfmsScp) {
	cur := x
	for cur != nil {
		kndIdx, exists := cur.Xpr.Vars[idn]
		if exists {
			r.Idx = kndIdx.Idx
			r.Arr = cur.anaPrfms
			return r
		}
		cur = cur.Prnt
	}
	panic(fmt.Sprintf("scope missing idn '%v'", idn))
}
func (x *Scp) AnaPrfmDlt(idn string) (r ana.PrfmDltScp) {
	cur := x
	for cur != nil {
		kndIdx, exists := cur.Xpr.Vars[idn]
		if exists {
			r.Idx = kndIdx.Idx
			r.Arr = cur.anaPrfmDlt
			return r
		}
		cur = cur.Prnt
	}
	panic(fmt.Sprintf("scope missing idn '%v'", idn))
}
func (x *Scp) AnaPort(idn string) (r ana.PortScp) {
	cur := x
	for cur != nil {
		kndIdx, exists := cur.Xpr.Vars[idn]
		if exists {
			r.Idx = kndIdx.Idx
			r.Arr = cur.anaPort
			return r
		}
		cur = cur.Prnt
	}
	panic(fmt.Sprintf("scope missing idn '%v'", idn))
}
func (x *Scp) HstPrv(idn string) (r hst.PrvScp) {
	cur := x
	for cur != nil {
		kndIdx, exists := cur.Xpr.Vars[idn]
		if exists {
			r.Idx = kndIdx.Idx
			r.Arr = cur.hstPrv
			return r
		}
		cur = cur.Prnt
	}
	panic(fmt.Sprintf("scope missing idn '%v'", idn))
}
func (x *Scp) HstInstr(idn string) (r hst.InstrScp) {
	cur := x
	for cur != nil {
		kndIdx, exists := cur.Xpr.Vars[idn]
		if exists {
			r.Idx = kndIdx.Idx
			r.Arr = cur.hstInstr
			return r
		}
		cur = cur.Prnt
	}
	panic(fmt.Sprintf("scope missing idn '%v'", idn))
}
func (x *Scp) HstInrvl(idn string) (r hst.InrvlScp) {
	cur := x
	for cur != nil {
		kndIdx, exists := cur.Xpr.Vars[idn]
		if exists {
			r.Idx = kndIdx.Idx
			r.Arr = cur.hstInrvl
			return r
		}
		cur = cur.Prnt
	}
	panic(fmt.Sprintf("scope missing idn '%v'", idn))
}
func (x *Scp) HstSide(idn string) (r hst.SideScp) {
	cur := x
	for cur != nil {
		kndIdx, exists := cur.Xpr.Vars[idn]
		if exists {
			r.Idx = kndIdx.Idx
			r.Arr = cur.hstSide
			return r
		}
		cur = cur.Prnt
	}
	panic(fmt.Sprintf("scope missing idn '%v'", idn))
}
func (x *Scp) HstStm(idn string) (r hst.StmScp) {
	cur := x
	for cur != nil {
		kndIdx, exists := cur.Xpr.Vars[idn]
		if exists {
			r.Idx = kndIdx.Idx
			r.Arr = cur.hstStm
			return r
		}
		cur = cur.Prnt
	}
	panic(fmt.Sprintf("scope missing idn '%v'", idn))
}
func (x *Scp) HstCnd(idn string) (r hst.CndScp) {
	cur := x
	for cur != nil {
		kndIdx, exists := cur.Xpr.Vars[idn]
		if exists {
			r.Idx = kndIdx.Idx
			r.Arr = cur.hstCnd
			return r
		}
		cur = cur.Prnt
	}
	panic(fmt.Sprintf("scope missing idn '%v'", idn))
}
func (x *Scp) HstStgy(idn string) (r hst.StgyScp) {
	cur := x
	for cur != nil {
		kndIdx, exists := cur.Xpr.Vars[idn]
		if exists {
			r.Idx = kndIdx.Idx
			r.Arr = cur.hstStgy
			return r
		}
		cur = cur.Prnt
	}
	panic(fmt.Sprintf("scope missing idn '%v'", idn))
}
func (x *Scp) HstPrvs(idn string) (r hst.PrvsScp) {
	cur := x
	for cur != nil {
		kndIdx, exists := cur.Xpr.Vars[idn]
		if exists {
			r.Idx = kndIdx.Idx
			r.Arr = cur.hstPrvs
			return r
		}
		cur = cur.Prnt
	}
	panic(fmt.Sprintf("scope missing idn '%v'", idn))
}
func (x *Scp) HstInstrs(idn string) (r hst.InstrsScp) {
	cur := x
	for cur != nil {
		kndIdx, exists := cur.Xpr.Vars[idn]
		if exists {
			r.Idx = kndIdx.Idx
			r.Arr = cur.hstInstrs
			return r
		}
		cur = cur.Prnt
	}
	panic(fmt.Sprintf("scope missing idn '%v'", idn))
}
func (x *Scp) HstInrvls(idn string) (r hst.InrvlsScp) {
	cur := x
	for cur != nil {
		kndIdx, exists := cur.Xpr.Vars[idn]
		if exists {
			r.Idx = kndIdx.Idx
			r.Arr = cur.hstInrvls
			return r
		}
		cur = cur.Prnt
	}
	panic(fmt.Sprintf("scope missing idn '%v'", idn))
}
func (x *Scp) HstSides(idn string) (r hst.SidesScp) {
	cur := x
	for cur != nil {
		kndIdx, exists := cur.Xpr.Vars[idn]
		if exists {
			r.Idx = kndIdx.Idx
			r.Arr = cur.hstSides
			return r
		}
		cur = cur.Prnt
	}
	panic(fmt.Sprintf("scope missing idn '%v'", idn))
}
func (x *Scp) HstStms(idn string) (r hst.StmsScp) {
	cur := x
	for cur != nil {
		kndIdx, exists := cur.Xpr.Vars[idn]
		if exists {
			r.Idx = kndIdx.Idx
			r.Arr = cur.hstStms
			return r
		}
		cur = cur.Prnt
	}
	panic(fmt.Sprintf("scope missing idn '%v'", idn))
}
func (x *Scp) HstCnds(idn string) (r hst.CndsScp) {
	cur := x
	for cur != nil {
		kndIdx, exists := cur.Xpr.Vars[idn]
		if exists {
			r.Idx = kndIdx.Idx
			r.Arr = cur.hstCnds
			return r
		}
		cur = cur.Prnt
	}
	panic(fmt.Sprintf("scope missing idn '%v'", idn))
}
func (x *Scp) HstStgys(idn string) (r hst.StgysScp) {
	cur := x
	for cur != nil {
		kndIdx, exists := cur.Xpr.Vars[idn]
		if exists {
			r.Idx = kndIdx.Idx
			r.Arr = cur.hstStgys
			return r
		}
		cur = cur.Prnt
	}
	panic(fmt.Sprintf("scope missing idn '%v'", idn))
}
func (x *Scp) RltPrv(idn string) (r rlt.PrvScp) {
	cur := x
	for cur != nil {
		kndIdx, exists := cur.Xpr.Vars[idn]
		if exists {
			r.Idx = kndIdx.Idx
			r.Arr = cur.rltPrv
			return r
		}
		cur = cur.Prnt
	}
	panic(fmt.Sprintf("scope missing idn '%v'", idn))
}
func (x *Scp) RltInstr(idn string) (r rlt.InstrScp) {
	cur := x
	for cur != nil {
		kndIdx, exists := cur.Xpr.Vars[idn]
		if exists {
			r.Idx = kndIdx.Idx
			r.Arr = cur.rltInstr
			return r
		}
		cur = cur.Prnt
	}
	panic(fmt.Sprintf("scope missing idn '%v'", idn))
}
func (x *Scp) RltInrvl(idn string) (r rlt.InrvlScp) {
	cur := x
	for cur != nil {
		kndIdx, exists := cur.Xpr.Vars[idn]
		if exists {
			r.Idx = kndIdx.Idx
			r.Arr = cur.rltInrvl
			return r
		}
		cur = cur.Prnt
	}
	panic(fmt.Sprintf("scope missing idn '%v'", idn))
}
func (x *Scp) RltSide(idn string) (r rlt.SideScp) {
	cur := x
	for cur != nil {
		kndIdx, exists := cur.Xpr.Vars[idn]
		if exists {
			r.Idx = kndIdx.Idx
			r.Arr = cur.rltSide
			return r
		}
		cur = cur.Prnt
	}
	panic(fmt.Sprintf("scope missing idn '%v'", idn))
}
func (x *Scp) RltStm(idn string) (r rlt.StmScp) {
	cur := x
	for cur != nil {
		kndIdx, exists := cur.Xpr.Vars[idn]
		if exists {
			r.Idx = kndIdx.Idx
			r.Arr = cur.rltStm
			return r
		}
		cur = cur.Prnt
	}
	panic(fmt.Sprintf("scope missing idn '%v'", idn))
}
func (x *Scp) RltCnd(idn string) (r rlt.CndScp) {
	cur := x
	for cur != nil {
		kndIdx, exists := cur.Xpr.Vars[idn]
		if exists {
			r.Idx = kndIdx.Idx
			r.Arr = cur.rltCnd
			return r
		}
		cur = cur.Prnt
	}
	panic(fmt.Sprintf("scope missing idn '%v'", idn))
}
func (x *Scp) RltStgy(idn string) (r rlt.StgyScp) {
	cur := x
	for cur != nil {
		kndIdx, exists := cur.Xpr.Vars[idn]
		if exists {
			r.Idx = kndIdx.Idx
			r.Arr = cur.rltStgy
			return r
		}
		cur = cur.Prnt
	}
	panic(fmt.Sprintf("scope missing idn '%v'", idn))
}
func (x *Scp) RltPrvs(idn string) (r rlt.PrvsScp) {
	cur := x
	for cur != nil {
		kndIdx, exists := cur.Xpr.Vars[idn]
		if exists {
			r.Idx = kndIdx.Idx
			r.Arr = cur.rltPrvs
			return r
		}
		cur = cur.Prnt
	}
	panic(fmt.Sprintf("scope missing idn '%v'", idn))
}
func (x *Scp) RltInstrs(idn string) (r rlt.InstrsScp) {
	cur := x
	for cur != nil {
		kndIdx, exists := cur.Xpr.Vars[idn]
		if exists {
			r.Idx = kndIdx.Idx
			r.Arr = cur.rltInstrs
			return r
		}
		cur = cur.Prnt
	}
	panic(fmt.Sprintf("scope missing idn '%v'", idn))
}
func (x *Scp) RltInrvls(idn string) (r rlt.InrvlsScp) {
	cur := x
	for cur != nil {
		kndIdx, exists := cur.Xpr.Vars[idn]
		if exists {
			r.Idx = kndIdx.Idx
			r.Arr = cur.rltInrvls
			return r
		}
		cur = cur.Prnt
	}
	panic(fmt.Sprintf("scope missing idn '%v'", idn))
}
func (x *Scp) RltSides(idn string) (r rlt.SidesScp) {
	cur := x
	for cur != nil {
		kndIdx, exists := cur.Xpr.Vars[idn]
		if exists {
			r.Idx = kndIdx.Idx
			r.Arr = cur.rltSides
			return r
		}
		cur = cur.Prnt
	}
	panic(fmt.Sprintf("scope missing idn '%v'", idn))
}
func (x *Scp) RltStms(idn string) (r rlt.StmsScp) {
	cur := x
	for cur != nil {
		kndIdx, exists := cur.Xpr.Vars[idn]
		if exists {
			r.Idx = kndIdx.Idx
			r.Arr = cur.rltStms
			return r
		}
		cur = cur.Prnt
	}
	panic(fmt.Sprintf("scope missing idn '%v'", idn))
}
func (x *Scp) RltCnds(idn string) (r rlt.CndsScp) {
	cur := x
	for cur != nil {
		kndIdx, exists := cur.Xpr.Vars[idn]
		if exists {
			r.Idx = kndIdx.Idx
			r.Arr = cur.rltCnds
			return r
		}
		cur = cur.Prnt
	}
	panic(fmt.Sprintf("scope missing idn '%v'", idn))
}
func (x *Scp) RltStgys(idn string) (r rlt.StgysScp) {
	cur := x
	for cur != nil {
		kndIdx, exists := cur.Xpr.Vars[idn]
		if exists {
			r.Idx = kndIdx.Idx
			r.Arr = cur.rltStgys
			return r
		}
		cur = cur.Prnt
	}
	panic(fmt.Sprintf("scope missing idn '%v'", idn))
}
func (x *Scp) FntFnt(idn string) (r fnt.FntScp) {
	cur := x
	for cur != nil {
		kndIdx, exists := cur.Xpr.Vars[idn]
		if exists {
			r.Idx = kndIdx.Idx
			r.Arr = cur.fntFnt
			return r
		}
		cur = cur.Prnt
	}
	panic(fmt.Sprintf("scope missing idn '%v'", idn))
}
func (x *Scp) ClrClr(idn string) (r clr.ClrScp) {
	cur := x
	for cur != nil {
		kndIdx, exists := cur.Xpr.Vars[idn]
		if exists {
			r.Idx = kndIdx.Idx
			r.Arr = cur.clrClr
			return r
		}
		cur = cur.Prnt
	}
	panic(fmt.Sprintf("scope missing idn '%v'", idn))
}
func (x *Scp) PenPen(idn string) (r pen.PenScp) {
	cur := x
	for cur != nil {
		kndIdx, exists := cur.Xpr.Vars[idn]
		if exists {
			r.Idx = kndIdx.Idx
			r.Arr = cur.penPen
			return r
		}
		cur = cur.Prnt
	}
	panic(fmt.Sprintf("scope missing idn '%v'", idn))
}
func (x *Scp) PenPens(idn string) (r pen.PensScp) {
	cur := x
	for cur != nil {
		kndIdx, exists := cur.Xpr.Vars[idn]
		if exists {
			r.Idx = kndIdx.Idx
			r.Arr = cur.penPens
			return r
		}
		cur = cur.Prnt
	}
	panic(fmt.Sprintf("scope missing idn '%v'", idn))
}
func (x *Scp) PltPlt(idn string) (r plt.PltScp) {
	cur := x
	for cur != nil {
		kndIdx, exists := cur.Xpr.Vars[idn]
		if exists {
			r.Idx = kndIdx.Idx
			r.Arr = cur.pltPlt
			return r
		}
		cur = cur.Prnt
	}
	panic(fmt.Sprintf("scope missing idn '%v'", idn))
}
func (x *Scp) PltPlts(idn string) (r plt.PltsScp) {
	cur := x
	for cur != nil {
		kndIdx, exists := cur.Xpr.Vars[idn]
		if exists {
			r.Idx = kndIdx.Idx
			r.Arr = cur.pltPlts
			return r
		}
		cur = cur.Prnt
	}
	panic(fmt.Sprintf("scope missing idn '%v'", idn))
}
func (x *Scp) PltStm(idn string) (r plt.StmScp) {
	cur := x
	for cur != nil {
		kndIdx, exists := cur.Xpr.Vars[idn]
		if exists {
			r.Idx = kndIdx.Idx
			r.Arr = cur.pltStm
			return r
		}
		cur = cur.Prnt
	}
	panic(fmt.Sprintf("scope missing idn '%v'", idn))
}
func (x *Scp) PltFltsSctr(idn string) (r plt.FltsSctrScp) {
	cur := x
	for cur != nil {
		kndIdx, exists := cur.Xpr.Vars[idn]
		if exists {
			r.Idx = kndIdx.Idx
			r.Arr = cur.pltFltsSctr
			return r
		}
		cur = cur.Prnt
	}
	panic(fmt.Sprintf("scope missing idn '%v'", idn))
}
func (x *Scp) PltFltsSctrDist(idn string) (r plt.FltsSctrDistScp) {
	cur := x
	for cur != nil {
		kndIdx, exists := cur.Xpr.Vars[idn]
		if exists {
			r.Idx = kndIdx.Idx
			r.Arr = cur.pltFltsSctrDist
			return r
		}
		cur = cur.Prnt
	}
	panic(fmt.Sprintf("scope missing idn '%v'", idn))
}
func (x *Scp) PltHrz(idn string) (r plt.HrzScp) {
	cur := x
	for cur != nil {
		kndIdx, exists := cur.Xpr.Vars[idn]
		if exists {
			r.Idx = kndIdx.Idx
			r.Arr = cur.pltHrz
			return r
		}
		cur = cur.Prnt
	}
	panic(fmt.Sprintf("scope missing idn '%v'", idn))
}
func (x *Scp) PltVrt(idn string) (r plt.VrtScp) {
	cur := x
	for cur != nil {
		kndIdx, exists := cur.Xpr.Vars[idn]
		if exists {
			r.Idx = kndIdx.Idx
			r.Arr = cur.pltVrt
			return r
		}
		cur = cur.Prnt
	}
	panic(fmt.Sprintf("scope missing idn '%v'", idn))
}
func (x *Scp) PltDpth(idn string) (r plt.DpthScp) {
	cur := x
	for cur != nil {
		kndIdx, exists := cur.Xpr.Vars[idn]
		if exists {
			r.Idx = kndIdx.Idx
			r.Arr = cur.pltDpth
			return r
		}
		cur = cur.Prnt
	}
	panic(fmt.Sprintf("scope missing idn '%v'", idn))
}
func (x *Scp) SysMu(idn string) (r sys.MuScp) {
	cur := x
	for cur != nil {
		kndIdx, exists := cur.Xpr.Vars[idn]
		if exists {
			r.Idx = kndIdx.Idx
			r.Arr = cur.sysMu
			return r
		}
		cur = cur.Prnt
	}
	panic(fmt.Sprintf("scope missing idn '%v'", idn))
}
