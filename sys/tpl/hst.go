package tpl

import (
	"sys/k"
	"sys/tpl/atr"
)

type (
	DirHst struct {
		DirBse
		// elm
		Prv   *FleHstPrv
		Instr *FleHstInstr
		Inrvl *FleHstInrvl
		Side  *FleHstSide
		Stm   *FleHstStm
		Cnd   *FleHstCnd
		Stgy  *FleHstStgy
		// Port  *FleHstPort
		// Prfm *FleHstPrfm
		// Splt *FleHstSplt
		// elm arr
		Prvs   *FleHstPrvs
		Instrs *FleHstInstrs
		Inrvls *FleHstInrvls
		Sides  *FleHstSides
		Stms   *FleHstStms
		Cnds   *FleHstCnds
		Stgys  *FleHstStgys
		// Ports  *FleHstPorts
		// Prfms *FleHstPrfms
		// // fbr
		// PrvFbr   *FleHstPrvFbr
		// InstrFbr *FleHstInstrFbr
		// InrvlFbr *FleHstInrvlFbr
		// SideFbr  *FleHstSideFbr
		// StmFbr   *FleHstStmFbr
		// CndFbr   *FleHstCndFbr
		// StgyFbr  *FleHstStgyFbr
		// PortFbr  *FleHstPortFbr
		// PrfmFbr  *FleHstPrfmFbr
		// // fbr arr
		// PrvFbrs   *FleHstPrvFbrs
		// InstrFbrs *FleHstInstrFbrs
		// InrvlFbrs *FleHstInrvlFbrs
		// SideFbrs  *FleHstSideFbrs
		// StmFbrs   *FleHstStmFbrs
		// CndFbrs   *FleHstCndFbrs
		// StgyFbrs  *FleHstStgyFbrs
		// PortFbrs  *FleHstPortFbrs
		// PrfmFbrs  *FleHstPrfmFbrs
		// // wve
		// PrvWve   *FleHstPrvWve
		// InstrWve *FleHstInstrWve
		// InrvlWve *FleHstInrvlWve
		// SideWve  *FleHstSideWve
		// StmWve   *FleHstStmWve
		// CndWve   *FleHstCndWve
		// StgyWve  *FleHstStgyWve
		// PortWve  *FleHstPortWve
		// PrfmWve  *FleHstPrfmWve
		// // // wve arr
		// // PrvWves   *FleHstPrvWves
		// // InstrWves *FleHstInstrWves
		// // InrvlWves *FleHstInrvlWves
		// // SideWves  *FleHstSideWves
		// StmWves *FleHstStmWves
		// // CndWves   *FleHstCndWves
		// // StgyWves  *FleHstStgyWves
		// // PortWves  *FleHstPortWves
		// // PrfmWves  *FleHstPrfmWves
		// //
		// // PrfmGrp *FleHstPrfmGrp
		// //

		// Prcp     *FleHstPrcp
		// PrcpSplt *FleHstPrcpSplt
		// StmSplt  *FleHstStmSplt
		// StmSplts *FleHstStmSplts

		// StmFbrSplt  *FleHstStmFbrSplt
		// StmFbrSplts *FleHstStmFbrSplts
		// StmWveSplt  *FleHstStmWveSplt
		// StmWveSplts *FleHstStmWveSplts
		// Ftr  *FleHstFtr
		// Ftrs *FleHstFtrs
	}
	FleHstBse struct {
		FleNodeBse
		seg *Struct
	}
)

func (x *DirAna) NewHst() (r *DirHst) {
	r = &DirHst{}
	x.Hst = r
	r.Pkg = x.Pkg.New(k.Hst)
	// elm
	r.NewPrv()
	r.NewInstr()
	r.NewInrvl()
	r.NewSide()
	r.NewStm()
	r.NewCnd()
	r.NewStgy()
	// r.NewPort()
	// r.NewPrfm()
	// r.NewSplt()
	// elm arr
	r.NewPrvs()
	r.NewInstrs()
	r.NewInrvls()
	r.NewSides()
	r.NewStms()
	r.NewCnds()
	r.NewStgys()
	// r.NewPorts()
	// r.NewPrfms()
	// // fbr
	// r.NewPrvFbr()
	// r.NewInstrFbr()
	// r.NewInrvlFbr()
	// r.NewSideFbr()
	// r.NewStmFbr()
	// r.NewCndFbr()
	// r.NewStgyFbr()
	// r.NewPortFbr()
	// r.NewPrfmFbr()
	// // fbr arr
	// r.NewPrvFbrs()
	// r.NewInstrFbrs()
	// r.NewInrvlFbrs()
	// r.NewSideFbrs()
	// r.NewStmFbrs()
	// r.NewCndFbrs()
	// r.NewStgyFbrs()
	// r.NewPortFbrs()
	// r.NewPrfmFbrs()
	// // // wve
	// r.NewPrvWve()
	// r.NewInstrWve()
	// r.NewInrvlWve()
	// r.NewSideWve()
	// r.NewStmWve()
	// r.NewCndWve()
	// r.NewStgyWve()
	// r.NewPortWve()
	// r.NewPrfmWve()
	// // // wve arr
	// // r.NewPrvWves()
	// // r.NewInstrWves()
	// // r.NewInrvlWves()
	// // r.NewSideWves()
	// r.NewStmWves()
	// // r.NewCndWves()
	// // r.NewStgyWves()
	// // r.NewPortWves()
	// // r.NewPrfmWves()
	// //
	// // r.NewPrfmGrp()
	// //

	// r.NewPrcp()
	// r.NewPrcpSplt()
	// r.NewStmSplt()
	// r.NewStmSplts()

	// r.NewStmFbrSplt()
	// r.NewStmFbrSplts()
	// r.NewStmWveSplt()
	// r.NewStmWveSplts()
	// r.NewFtr()
	// r.NewFtrs()
	return r
}

func (x *FleHstBse) NewSeg(suffix ...string) (r *Struct) {
	var s string
	if len(suffix) != 0 {
		s = suffix[0]
	}
	r = x.Structf("%v%vSeg", atr.None, x.Typ().Title(), s)
	r.FldTyp(_sys.Bsc.Bnd)
	return r
}
func (x *FleHstBse) NodeSeg(name string) (r *Struct) {
	r = x.StructPtrf("%vSeg", atr.None, name)
	r.FldTyp(x.seg)
	return r
}
