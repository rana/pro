package act

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"sync"
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
	"sys/lng/pro/trm"
	"sys/lng/pro/trm/prs"
	"sys/lng/pro/xpr"
	"sys/log/log"
)

type (
	Actr struct {
		xpr.Xprr
	}
	Act interface {
		Act()
		Ifc() interface{}
	}
	LogLogrAct interface {
		Act()
		Ifc() interface{}
		LogLogr() log.Logr
	}
	StrStrAct interface {
		Act()
		Ifc() interface{}
		StrStr() str.Str
	}
	BolBolAct interface {
		Act()
		Ifc() interface{}
		BolBol() bol.Bol
	}
	FltFltAct interface {
		Act()
		Ifc() interface{}
		FltFlt() flt.Flt
	}
	UntUntAct interface {
		Act()
		Ifc() interface{}
		UntUnt() unt.Unt
	}
	IntIntAct interface {
		Act()
		Ifc() interface{}
		IntInt() int.Int
	}
	TmeTmeAct interface {
		Act()
		Ifc() interface{}
		TmeTme() tme.Tme
	}
	BndBndAct interface {
		Act()
		Ifc() interface{}
		BndBnd() bnd.Bnd
	}
	FltRngAct interface {
		Act()
		Ifc() interface{}
		FltRng() flt.Rng
	}
	TmeRngAct interface {
		Act()
		Ifc() interface{}
		TmeRng() tme.Rng
	}
	StrsStrsAct interface {
		Act()
		Ifc() interface{}
		StrsStrs() *strs.Strs
	}
	BolsBolsAct interface {
		Act()
		Ifc() interface{}
		BolsBols() *bols.Bols
	}
	FltsFltsAct interface {
		Act()
		Ifc() interface{}
		FltsFlts() *flts.Flts
	}
	UntsUntsAct interface {
		Act()
		Ifc() interface{}
		UntsUnts() *unts.Unts
	}
	IntsIntsAct interface {
		Act()
		Ifc() interface{}
		IntsInts() *ints.Ints
	}
	TmesTmesAct interface {
		Act()
		Ifc() interface{}
		TmesTmes() *tmes.Tmes
	}
	BndsBndsAct interface {
		Act()
		Ifc() interface{}
		BndsBnds() *bnds.Bnds
	}
	TmeRngsAct interface {
		Act()
		Ifc() interface{}
		TmeRngs() *tme.Rngs
	}
	AnaTrdAct interface {
		Act()
		Ifc() interface{}
		AnaTrd() *ana.Trd
	}
	AnaTrdsAct interface {
		Act()
		Ifc() interface{}
		AnaTrds() *ana.Trds
	}
	AnaPrfmAct interface {
		Act()
		Ifc() interface{}
		AnaPrfm() *ana.Prfm
	}
	AnaPrfmsAct interface {
		Act()
		Ifc() interface{}
		AnaPrfms() *ana.Prfms
	}
	AnaPrfmDltAct interface {
		Act()
		Ifc() interface{}
		AnaPrfmDlt() *ana.PrfmDlt
	}
	AnaPortAct interface {
		Act()
		Ifc() interface{}
		AnaPort() *ana.Port
	}
	HstPrvAct interface {
		Act()
		Ifc() interface{}
		HstPrv() hst.Prv
	}
	HstInstrAct interface {
		Act()
		Ifc() interface{}
		HstInstr() hst.Instr
	}
	HstInrvlAct interface {
		Act()
		Ifc() interface{}
		HstInrvl() hst.Inrvl
	}
	HstSideAct interface {
		Act()
		Ifc() interface{}
		HstSide() hst.Side
	}
	HstStmAct interface {
		Act()
		Ifc() interface{}
		HstStm() hst.Stm
	}
	HstCndAct interface {
		Act()
		Ifc() interface{}
		HstCnd() hst.Cnd
	}
	HstStgyAct interface {
		Act()
		Ifc() interface{}
		HstStgy() hst.Stgy
	}
	HstPrvsAct interface {
		Act()
		Ifc() interface{}
		HstPrvs() *hst.Prvs
	}
	HstInstrsAct interface {
		Act()
		Ifc() interface{}
		HstInstrs() *hst.Instrs
	}
	HstInrvlsAct interface {
		Act()
		Ifc() interface{}
		HstInrvls() *hst.Inrvls
	}
	HstSidesAct interface {
		Act()
		Ifc() interface{}
		HstSides() *hst.Sides
	}
	HstStmsAct interface {
		Act()
		Ifc() interface{}
		HstStms() *hst.Stms
	}
	HstCndsAct interface {
		Act()
		Ifc() interface{}
		HstCnds() *hst.Cnds
	}
	HstStgysAct interface {
		Act()
		Ifc() interface{}
		HstStgys() *hst.Stgys
	}
	RltPrvAct interface {
		Act()
		Ifc() interface{}
		RltPrv() rlt.Prv
	}
	RltInstrAct interface {
		Act()
		Ifc() interface{}
		RltInstr() rlt.Instr
	}
	RltInrvlAct interface {
		Act()
		Ifc() interface{}
		RltInrvl() rlt.Inrvl
	}
	RltSideAct interface {
		Act()
		Ifc() interface{}
		RltSide() rlt.Side
	}
	RltStmAct interface {
		Act()
		Ifc() interface{}
		RltStm() rlt.Stm
	}
	RltCndAct interface {
		Act()
		Ifc() interface{}
		RltCnd() rlt.Cnd
	}
	RltStgyAct interface {
		Act()
		Ifc() interface{}
		RltStgy() rlt.Stgy
	}
	RltPrvsAct interface {
		Act()
		Ifc() interface{}
		RltPrvs() *rlt.Prvs
	}
	RltInstrsAct interface {
		Act()
		Ifc() interface{}
		RltInstrs() *rlt.Instrs
	}
	RltInrvlsAct interface {
		Act()
		Ifc() interface{}
		RltInrvls() *rlt.Inrvls
	}
	RltSidesAct interface {
		Act()
		Ifc() interface{}
		RltSides() *rlt.Sides
	}
	RltStmsAct interface {
		Act()
		Ifc() interface{}
		RltStms() *rlt.Stms
	}
	RltCndsAct interface {
		Act()
		Ifc() interface{}
		RltCnds() *rlt.Cnds
	}
	RltStgysAct interface {
		Act()
		Ifc() interface{}
		RltStgys() *rlt.Stgys
	}
	FntFntAct interface {
		Act()
		Ifc() interface{}
		FntFnt() *fnt.Fnt
	}
	ClrClrAct interface {
		Act()
		Ifc() interface{}
		ClrClr() clr.Clr
	}
	PenPenAct interface {
		Act()
		Ifc() interface{}
		PenPen() pen.Pen
	}
	PenPensAct interface {
		Act()
		Ifc() interface{}
		PenPens() *pen.Pens
	}
	PltPltAct interface {
		Act()
		Ifc() interface{}
		PltPlt() plt.Plt
	}
	PltPltsAct interface {
		Act()
		Ifc() interface{}
		PltPlts() *plt.Plts
	}
	PltTmeAxisXAct interface {
		Act()
		Ifc() interface{}
		PltTmeAxisX() *plt.TmeAxisX
	}
	PltFltAxisYAct interface {
		Act()
		Ifc() interface{}
		PltFltAxisY() *plt.FltAxisY
	}
	PltStmAct interface {
		Act()
		Ifc() interface{}
		PltStm() *plt.Stm
		PltPlt() plt.Plt
	}
	PltFltsSctrAct interface {
		Act()
		Ifc() interface{}
		PltFltsSctr() *plt.FltsSctr
		PltPlt() plt.Plt
	}
	PltFltsSctrDistAct interface {
		Act()
		Ifc() interface{}
		PltFltsSctrDist() *plt.FltsSctrDist
		PltPlt() plt.Plt
	}
	PltHrzAct interface {
		Act()
		Ifc() interface{}
		PltHrz() *plt.Hrz
		PltPlt() plt.Plt
	}
	PltVrtAct interface {
		Act()
		Ifc() interface{}
		PltVrt() *plt.Vrt
		PltPlt() plt.Plt
	}
	PltDpthAct interface {
		Act()
		Ifc() interface{}
		PltDpth() *plt.Dpth
		PltPlt() plt.Plt
	}
	SysMuAct interface {
		Act()
		Ifc() interface{}
		SysMu() *sys.Mu
	}
	StrStrLit struct {
		Trm trm.StrLit
		Txt string
	}
	BolBolLit struct {
		Trm trm.BolLit
		Txt string
	}
	FltFltLit struct {
		Trm trm.FltLit
		Txt string
	}
	UntUntLit struct {
		Trm trm.UntLit
		Txt string
	}
	IntIntLit struct {
		Trm trm.IntLit
		Txt string
	}
	TmeTmeLit struct {
		Trm trm.TmeLit
		Txt string
	}
	BndBndLit struct {
		Trm trm.BndLit
		Txt string
	}
	FltRngLit struct {
		Trm trm.FltRngLit
		Txt string
	}
	TmeRngLit struct {
		Trm trm.TmeRngLit
		Txt string
	}
	StrsStrsLit struct {
		Trm trm.StrsLit
		Txt string
	}
	BolsBolsLit struct {
		Trm trm.BolsLit
		Txt string
	}
	FltsFltsLit struct {
		Trm trm.FltsLit
		Txt string
	}
	UntsUntsLit struct {
		Trm trm.UntsLit
		Txt string
	}
	IntsIntsLit struct {
		Trm trm.IntsLit
		Txt string
	}
	TmesTmesLit struct {
		Trm trm.TmesLit
		Txt string
	}
	BndsBndsLit struct {
		Trm trm.BndsLit
		Txt string
	}
	TmeRngsLit struct {
		Trm trm.TmeRngsLit
		Txt string
	}
	StrStrAsn struct {
		str.StrScp
		X StrStrAct
	}
	StrStrAcs struct {
		str.StrScp
	}
	BolBolAsn struct {
		bol.BolScp
		X BolBolAct
	}
	BolBolAcs struct {
		bol.BolScp
	}
	FltFltAsn struct {
		flt.FltScp
		X FltFltAct
	}
	FltFltAcs struct {
		flt.FltScp
	}
	UntUntAsn struct {
		unt.UntScp
		X UntUntAct
	}
	UntUntAcs struct {
		unt.UntScp
	}
	IntIntAsn struct {
		int.IntScp
		X IntIntAct
	}
	IntIntAcs struct {
		int.IntScp
	}
	TmeTmeAsn struct {
		tme.TmeScp
		X TmeTmeAct
	}
	TmeTmeAcs struct {
		tme.TmeScp
	}
	BndBndAsn struct {
		bnd.BndScp
		X BndBndAct
	}
	BndBndAcs struct {
		bnd.BndScp
	}
	FltRngAsn struct {
		flt.RngScp
		X FltRngAct
	}
	FltRngAcs struct {
		flt.RngScp
	}
	TmeRngAsn struct {
		tme.RngScp
		X TmeRngAct
	}
	TmeRngAcs struct {
		tme.RngScp
	}
	StrsStrsAsn struct {
		strs.StrsScp
		X StrsStrsAct
	}
	StrsStrsAcs struct {
		strs.StrsScp
	}
	StrsStrsEach struct {
		str.StrScp
		X    StrsStrsAct
		Acts []Act
	}
	StrsStrsPllEach struct {
		X          StrsStrsAct
		Idn        bnd.Bnd
		Txt        string
		ActScpPrnt *Scp
		XprScp     *xpr.Scp
		Xprs       []xpr.Xpr
	}
	StrsStrsPllEachSeg struct {
		Val        str.Str
		Idn        bnd.Bnd
		Txt        string
		ActScpPrnt *Scp
		XprScp     *xpr.Scp
		Xprs       []xpr.Xpr
	}
	BolsBolsAsn struct {
		bols.BolsScp
		X BolsBolsAct
	}
	BolsBolsAcs struct {
		bols.BolsScp
	}
	BolsBolsEach struct {
		bol.BolScp
		X    BolsBolsAct
		Acts []Act
	}
	BolsBolsPllEach struct {
		X          BolsBolsAct
		Idn        bnd.Bnd
		Txt        string
		ActScpPrnt *Scp
		XprScp     *xpr.Scp
		Xprs       []xpr.Xpr
	}
	BolsBolsPllEachSeg struct {
		Val        bol.Bol
		Idn        bnd.Bnd
		Txt        string
		ActScpPrnt *Scp
		XprScp     *xpr.Scp
		Xprs       []xpr.Xpr
	}
	FltsFltsAsn struct {
		flts.FltsScp
		X FltsFltsAct
	}
	FltsFltsAcs struct {
		flts.FltsScp
	}
	FltsFltsEach struct {
		flt.FltScp
		X    FltsFltsAct
		Acts []Act
	}
	FltsFltsPllEach struct {
		X          FltsFltsAct
		Idn        bnd.Bnd
		Txt        string
		ActScpPrnt *Scp
		XprScp     *xpr.Scp
		Xprs       []xpr.Xpr
	}
	FltsFltsPllEachSeg struct {
		Val        flt.Flt
		Idn        bnd.Bnd
		Txt        string
		ActScpPrnt *Scp
		XprScp     *xpr.Scp
		Xprs       []xpr.Xpr
	}
	UntsUntsAsn struct {
		unts.UntsScp
		X UntsUntsAct
	}
	UntsUntsAcs struct {
		unts.UntsScp
	}
	UntsUntsEach struct {
		unt.UntScp
		X    UntsUntsAct
		Acts []Act
	}
	UntsUntsPllEach struct {
		X          UntsUntsAct
		Idn        bnd.Bnd
		Txt        string
		ActScpPrnt *Scp
		XprScp     *xpr.Scp
		Xprs       []xpr.Xpr
	}
	UntsUntsPllEachSeg struct {
		Val        unt.Unt
		Idn        bnd.Bnd
		Txt        string
		ActScpPrnt *Scp
		XprScp     *xpr.Scp
		Xprs       []xpr.Xpr
	}
	IntsIntsAsn struct {
		ints.IntsScp
		X IntsIntsAct
	}
	IntsIntsAcs struct {
		ints.IntsScp
	}
	IntsIntsEach struct {
		int.IntScp
		X    IntsIntsAct
		Acts []Act
	}
	IntsIntsPllEach struct {
		X          IntsIntsAct
		Idn        bnd.Bnd
		Txt        string
		ActScpPrnt *Scp
		XprScp     *xpr.Scp
		Xprs       []xpr.Xpr
	}
	IntsIntsPllEachSeg struct {
		Val        int.Int
		Idn        bnd.Bnd
		Txt        string
		ActScpPrnt *Scp
		XprScp     *xpr.Scp
		Xprs       []xpr.Xpr
	}
	TmesTmesAsn struct {
		tmes.TmesScp
		X TmesTmesAct
	}
	TmesTmesAcs struct {
		tmes.TmesScp
	}
	TmesTmesEach struct {
		tme.TmeScp
		X    TmesTmesAct
		Acts []Act
	}
	TmesTmesPllEach struct {
		X          TmesTmesAct
		Idn        bnd.Bnd
		Txt        string
		ActScpPrnt *Scp
		XprScp     *xpr.Scp
		Xprs       []xpr.Xpr
	}
	TmesTmesPllEachSeg struct {
		Val        tme.Tme
		Idn        bnd.Bnd
		Txt        string
		ActScpPrnt *Scp
		XprScp     *xpr.Scp
		Xprs       []xpr.Xpr
	}
	BndsBndsAsn struct {
		bnds.BndsScp
		X BndsBndsAct
	}
	BndsBndsAcs struct {
		bnds.BndsScp
	}
	BndsBndsEach struct {
		bnd.BndScp
		X    BndsBndsAct
		Acts []Act
	}
	BndsBndsPllEach struct {
		X          BndsBndsAct
		Idn        bnd.Bnd
		Txt        string
		ActScpPrnt *Scp
		XprScp     *xpr.Scp
		Xprs       []xpr.Xpr
	}
	BndsBndsPllEachSeg struct {
		Val        bnd.Bnd
		Idn        bnd.Bnd
		Txt        string
		ActScpPrnt *Scp
		XprScp     *xpr.Scp
		Xprs       []xpr.Xpr
	}
	TmeRngsAsn struct {
		tme.RngsScp
		X TmeRngsAct
	}
	TmeRngsAcs struct {
		tme.RngsScp
	}
	TmeRngsEach struct {
		tme.RngScp
		X    TmeRngsAct
		Acts []Act
	}
	TmeRngsPllEach struct {
		X          TmeRngsAct
		Idn        bnd.Bnd
		Txt        string
		ActScpPrnt *Scp
		XprScp     *xpr.Scp
		Xprs       []xpr.Xpr
	}
	TmeRngsPllEachSeg struct {
		Val        tme.Rng
		Idn        bnd.Bnd
		Txt        string
		ActScpPrnt *Scp
		XprScp     *xpr.Scp
		Xprs       []xpr.Xpr
	}
	AnaTrdAsn struct {
		ana.TrdScp
		X AnaTrdAct
	}
	AnaTrdAcs struct {
		ana.TrdScp
	}
	AnaTrdsAsn struct {
		ana.TrdsScp
		X AnaTrdsAct
	}
	AnaTrdsAcs struct {
		ana.TrdsScp
	}
	AnaTrdsEach struct {
		ana.TrdScp
		X    AnaTrdsAct
		Acts []Act
	}
	AnaTrdsPllEach struct {
		X          AnaTrdsAct
		Idn        bnd.Bnd
		Txt        string
		ActScpPrnt *Scp
		XprScp     *xpr.Scp
		Xprs       []xpr.Xpr
	}
	AnaTrdsPllEachSeg struct {
		Val        *ana.Trd
		Idn        bnd.Bnd
		Txt        string
		ActScpPrnt *Scp
		XprScp     *xpr.Scp
		Xprs       []xpr.Xpr
	}
	AnaPrfmAsn struct {
		ana.PrfmScp
		X AnaPrfmAct
	}
	AnaPrfmAcs struct {
		ana.PrfmScp
	}
	AnaPrfmsAsn struct {
		ana.PrfmsScp
		X AnaPrfmsAct
	}
	AnaPrfmsAcs struct {
		ana.PrfmsScp
	}
	AnaPrfmsEach struct {
		ana.PrfmScp
		X    AnaPrfmsAct
		Acts []Act
	}
	AnaPrfmsPllEach struct {
		X          AnaPrfmsAct
		Idn        bnd.Bnd
		Txt        string
		ActScpPrnt *Scp
		XprScp     *xpr.Scp
		Xprs       []xpr.Xpr
	}
	AnaPrfmsPllEachSeg struct {
		Val        *ana.Prfm
		Idn        bnd.Bnd
		Txt        string
		ActScpPrnt *Scp
		XprScp     *xpr.Scp
		Xprs       []xpr.Xpr
	}
	AnaPrfmDltAsn struct {
		ana.PrfmDltScp
		X AnaPrfmDltAct
	}
	AnaPrfmDltAcs struct {
		ana.PrfmDltScp
	}
	AnaPortAsn struct {
		ana.PortScp
		X AnaPortAct
	}
	AnaPortAcs struct {
		ana.PortScp
	}
	HstPrvAsn struct {
		hst.PrvScp
		X HstPrvAct
	}
	HstPrvAcs struct {
		hst.PrvScp
	}
	HstInstrAsn struct {
		hst.InstrScp
		X HstInstrAct
	}
	HstInstrAcs struct {
		hst.InstrScp
	}
	HstInrvlAsn struct {
		hst.InrvlScp
		X HstInrvlAct
	}
	HstInrvlAcs struct {
		hst.InrvlScp
	}
	HstSideAsn struct {
		hst.SideScp
		X HstSideAct
	}
	HstSideAcs struct {
		hst.SideScp
	}
	HstStmAsn struct {
		hst.StmScp
		X HstStmAct
	}
	HstStmAcs struct {
		hst.StmScp
	}
	HstCndAsn struct {
		hst.CndScp
		X HstCndAct
	}
	HstCndAcs struct {
		hst.CndScp
	}
	HstStgyAsn struct {
		hst.StgyScp
		X HstStgyAct
	}
	HstStgyAcs struct {
		hst.StgyScp
	}
	HstPrvsAsn struct {
		hst.PrvsScp
		X HstPrvsAct
	}
	HstPrvsAcs struct {
		hst.PrvsScp
	}
	HstPrvsEach struct {
		hst.PrvScp
		X    HstPrvsAct
		Acts []Act
	}
	HstPrvsPllEach struct {
		X          HstPrvsAct
		Idn        bnd.Bnd
		Txt        string
		ActScpPrnt *Scp
		XprScp     *xpr.Scp
		Xprs       []xpr.Xpr
	}
	HstPrvsPllEachSeg struct {
		Val        hst.Prv
		Idn        bnd.Bnd
		Txt        string
		ActScpPrnt *Scp
		XprScp     *xpr.Scp
		Xprs       []xpr.Xpr
	}
	HstInstrsAsn struct {
		hst.InstrsScp
		X HstInstrsAct
	}
	HstInstrsAcs struct {
		hst.InstrsScp
	}
	HstInstrsEach struct {
		hst.InstrScp
		X    HstInstrsAct
		Acts []Act
	}
	HstInstrsPllEach struct {
		X          HstInstrsAct
		Idn        bnd.Bnd
		Txt        string
		ActScpPrnt *Scp
		XprScp     *xpr.Scp
		Xprs       []xpr.Xpr
	}
	HstInstrsPllEachSeg struct {
		Val        hst.Instr
		Idn        bnd.Bnd
		Txt        string
		ActScpPrnt *Scp
		XprScp     *xpr.Scp
		Xprs       []xpr.Xpr
	}
	HstInrvlsAsn struct {
		hst.InrvlsScp
		X HstInrvlsAct
	}
	HstInrvlsAcs struct {
		hst.InrvlsScp
	}
	HstInrvlsEach struct {
		hst.InrvlScp
		X    HstInrvlsAct
		Acts []Act
	}
	HstInrvlsPllEach struct {
		X          HstInrvlsAct
		Idn        bnd.Bnd
		Txt        string
		ActScpPrnt *Scp
		XprScp     *xpr.Scp
		Xprs       []xpr.Xpr
	}
	HstInrvlsPllEachSeg struct {
		Val        hst.Inrvl
		Idn        bnd.Bnd
		Txt        string
		ActScpPrnt *Scp
		XprScp     *xpr.Scp
		Xprs       []xpr.Xpr
	}
	HstSidesAsn struct {
		hst.SidesScp
		X HstSidesAct
	}
	HstSidesAcs struct {
		hst.SidesScp
	}
	HstSidesEach struct {
		hst.SideScp
		X    HstSidesAct
		Acts []Act
	}
	HstSidesPllEach struct {
		X          HstSidesAct
		Idn        bnd.Bnd
		Txt        string
		ActScpPrnt *Scp
		XprScp     *xpr.Scp
		Xprs       []xpr.Xpr
	}
	HstSidesPllEachSeg struct {
		Val        hst.Side
		Idn        bnd.Bnd
		Txt        string
		ActScpPrnt *Scp
		XprScp     *xpr.Scp
		Xprs       []xpr.Xpr
	}
	HstStmsAsn struct {
		hst.StmsScp
		X HstStmsAct
	}
	HstStmsAcs struct {
		hst.StmsScp
	}
	HstStmsEach struct {
		hst.StmScp
		X    HstStmsAct
		Acts []Act
	}
	HstStmsPllEach struct {
		X          HstStmsAct
		Idn        bnd.Bnd
		Txt        string
		ActScpPrnt *Scp
		XprScp     *xpr.Scp
		Xprs       []xpr.Xpr
	}
	HstStmsPllEachSeg struct {
		Val        hst.Stm
		Idn        bnd.Bnd
		Txt        string
		ActScpPrnt *Scp
		XprScp     *xpr.Scp
		Xprs       []xpr.Xpr
	}
	HstCndsAsn struct {
		hst.CndsScp
		X HstCndsAct
	}
	HstCndsAcs struct {
		hst.CndsScp
	}
	HstCndsEach struct {
		hst.CndScp
		X    HstCndsAct
		Acts []Act
	}
	HstCndsPllEach struct {
		X          HstCndsAct
		Idn        bnd.Bnd
		Txt        string
		ActScpPrnt *Scp
		XprScp     *xpr.Scp
		Xprs       []xpr.Xpr
	}
	HstCndsPllEachSeg struct {
		Val        hst.Cnd
		Idn        bnd.Bnd
		Txt        string
		ActScpPrnt *Scp
		XprScp     *xpr.Scp
		Xprs       []xpr.Xpr
	}
	HstStgysAsn struct {
		hst.StgysScp
		X HstStgysAct
	}
	HstStgysAcs struct {
		hst.StgysScp
	}
	HstStgysEach struct {
		hst.StgyScp
		X    HstStgysAct
		Acts []Act
	}
	HstStgysPllEach struct {
		X          HstStgysAct
		Idn        bnd.Bnd
		Txt        string
		ActScpPrnt *Scp
		XprScp     *xpr.Scp
		Xprs       []xpr.Xpr
	}
	HstStgysPllEachSeg struct {
		Val        hst.Stgy
		Idn        bnd.Bnd
		Txt        string
		ActScpPrnt *Scp
		XprScp     *xpr.Scp
		Xprs       []xpr.Xpr
	}
	RltPrvAsn struct {
		rlt.PrvScp
		X RltPrvAct
	}
	RltPrvAcs struct {
		rlt.PrvScp
	}
	RltInstrAsn struct {
		rlt.InstrScp
		X RltInstrAct
	}
	RltInstrAcs struct {
		rlt.InstrScp
	}
	RltInrvlAsn struct {
		rlt.InrvlScp
		X RltInrvlAct
	}
	RltInrvlAcs struct {
		rlt.InrvlScp
	}
	RltSideAsn struct {
		rlt.SideScp
		X RltSideAct
	}
	RltSideAcs struct {
		rlt.SideScp
	}
	RltStmAsn struct {
		rlt.StmScp
		X RltStmAct
	}
	RltStmAcs struct {
		rlt.StmScp
	}
	RltCndAsn struct {
		rlt.CndScp
		X RltCndAct
	}
	RltCndAcs struct {
		rlt.CndScp
	}
	RltStgyAsn struct {
		rlt.StgyScp
		X RltStgyAct
	}
	RltStgyAcs struct {
		rlt.StgyScp
	}
	RltPrvsAsn struct {
		rlt.PrvsScp
		X RltPrvsAct
	}
	RltPrvsAcs struct {
		rlt.PrvsScp
	}
	RltPrvsEach struct {
		rlt.PrvScp
		X    RltPrvsAct
		Acts []Act
	}
	RltPrvsPllEach struct {
		X          RltPrvsAct
		Idn        bnd.Bnd
		Txt        string
		ActScpPrnt *Scp
		XprScp     *xpr.Scp
		Xprs       []xpr.Xpr
	}
	RltPrvsPllEachSeg struct {
		Val        rlt.Prv
		Idn        bnd.Bnd
		Txt        string
		ActScpPrnt *Scp
		XprScp     *xpr.Scp
		Xprs       []xpr.Xpr
	}
	RltInstrsAsn struct {
		rlt.InstrsScp
		X RltInstrsAct
	}
	RltInstrsAcs struct {
		rlt.InstrsScp
	}
	RltInstrsEach struct {
		rlt.InstrScp
		X    RltInstrsAct
		Acts []Act
	}
	RltInstrsPllEach struct {
		X          RltInstrsAct
		Idn        bnd.Bnd
		Txt        string
		ActScpPrnt *Scp
		XprScp     *xpr.Scp
		Xprs       []xpr.Xpr
	}
	RltInstrsPllEachSeg struct {
		Val        rlt.Instr
		Idn        bnd.Bnd
		Txt        string
		ActScpPrnt *Scp
		XprScp     *xpr.Scp
		Xprs       []xpr.Xpr
	}
	RltInrvlsAsn struct {
		rlt.InrvlsScp
		X RltInrvlsAct
	}
	RltInrvlsAcs struct {
		rlt.InrvlsScp
	}
	RltInrvlsEach struct {
		rlt.InrvlScp
		X    RltInrvlsAct
		Acts []Act
	}
	RltInrvlsPllEach struct {
		X          RltInrvlsAct
		Idn        bnd.Bnd
		Txt        string
		ActScpPrnt *Scp
		XprScp     *xpr.Scp
		Xprs       []xpr.Xpr
	}
	RltInrvlsPllEachSeg struct {
		Val        rlt.Inrvl
		Idn        bnd.Bnd
		Txt        string
		ActScpPrnt *Scp
		XprScp     *xpr.Scp
		Xprs       []xpr.Xpr
	}
	RltSidesAsn struct {
		rlt.SidesScp
		X RltSidesAct
	}
	RltSidesAcs struct {
		rlt.SidesScp
	}
	RltSidesEach struct {
		rlt.SideScp
		X    RltSidesAct
		Acts []Act
	}
	RltSidesPllEach struct {
		X          RltSidesAct
		Idn        bnd.Bnd
		Txt        string
		ActScpPrnt *Scp
		XprScp     *xpr.Scp
		Xprs       []xpr.Xpr
	}
	RltSidesPllEachSeg struct {
		Val        rlt.Side
		Idn        bnd.Bnd
		Txt        string
		ActScpPrnt *Scp
		XprScp     *xpr.Scp
		Xprs       []xpr.Xpr
	}
	RltStmsAsn struct {
		rlt.StmsScp
		X RltStmsAct
	}
	RltStmsAcs struct {
		rlt.StmsScp
	}
	RltStmsEach struct {
		rlt.StmScp
		X    RltStmsAct
		Acts []Act
	}
	RltStmsPllEach struct {
		X          RltStmsAct
		Idn        bnd.Bnd
		Txt        string
		ActScpPrnt *Scp
		XprScp     *xpr.Scp
		Xprs       []xpr.Xpr
	}
	RltStmsPllEachSeg struct {
		Val        rlt.Stm
		Idn        bnd.Bnd
		Txt        string
		ActScpPrnt *Scp
		XprScp     *xpr.Scp
		Xprs       []xpr.Xpr
	}
	RltCndsAsn struct {
		rlt.CndsScp
		X RltCndsAct
	}
	RltCndsAcs struct {
		rlt.CndsScp
	}
	RltCndsEach struct {
		rlt.CndScp
		X    RltCndsAct
		Acts []Act
	}
	RltCndsPllEach struct {
		X          RltCndsAct
		Idn        bnd.Bnd
		Txt        string
		ActScpPrnt *Scp
		XprScp     *xpr.Scp
		Xprs       []xpr.Xpr
	}
	RltCndsPllEachSeg struct {
		Val        rlt.Cnd
		Idn        bnd.Bnd
		Txt        string
		ActScpPrnt *Scp
		XprScp     *xpr.Scp
		Xprs       []xpr.Xpr
	}
	RltStgysAsn struct {
		rlt.StgysScp
		X RltStgysAct
	}
	RltStgysAcs struct {
		rlt.StgysScp
	}
	RltStgysEach struct {
		rlt.StgyScp
		X    RltStgysAct
		Acts []Act
	}
	RltStgysPllEach struct {
		X          RltStgysAct
		Idn        bnd.Bnd
		Txt        string
		ActScpPrnt *Scp
		XprScp     *xpr.Scp
		Xprs       []xpr.Xpr
	}
	RltStgysPllEachSeg struct {
		Val        rlt.Stgy
		Idn        bnd.Bnd
		Txt        string
		ActScpPrnt *Scp
		XprScp     *xpr.Scp
		Xprs       []xpr.Xpr
	}
	FntFntAsn struct {
		fnt.FntScp
		X FntFntAct
	}
	FntFntAcs struct {
		fnt.FntScp
	}
	ClrClrAsn struct {
		clr.ClrScp
		X ClrClrAct
	}
	ClrClrAcs struct {
		clr.ClrScp
	}
	PenPenAsn struct {
		pen.PenScp
		X PenPenAct
	}
	PenPenAcs struct {
		pen.PenScp
	}
	PenPensAsn struct {
		pen.PensScp
		X PenPensAct
	}
	PenPensAcs struct {
		pen.PensScp
	}
	PenPensEach struct {
		pen.PenScp
		X    PenPensAct
		Acts []Act
	}
	PenPensPllEach struct {
		X          PenPensAct
		Idn        bnd.Bnd
		Txt        string
		ActScpPrnt *Scp
		XprScp     *xpr.Scp
		Xprs       []xpr.Xpr
	}
	PenPensPllEachSeg struct {
		Val        pen.Pen
		Idn        bnd.Bnd
		Txt        string
		ActScpPrnt *Scp
		XprScp     *xpr.Scp
		Xprs       []xpr.Xpr
	}
	PltPltAsn struct {
		plt.PltScp
		X PltPltAct
	}
	PltPltAcs struct {
		plt.PltScp
	}
	PltPltsAsn struct {
		plt.PltsScp
		X PltPltsAct
	}
	PltPltsAcs struct {
		plt.PltsScp
	}
	PltPltsEach struct {
		plt.PltScp
		X    PltPltsAct
		Acts []Act
	}
	PltPltsPllEach struct {
		X          PltPltsAct
		Idn        bnd.Bnd
		Txt        string
		ActScpPrnt *Scp
		XprScp     *xpr.Scp
		Xprs       []xpr.Xpr
	}
	PltPltsPllEachSeg struct {
		Val        plt.Plt
		Idn        bnd.Bnd
		Txt        string
		ActScpPrnt *Scp
		XprScp     *xpr.Scp
		Xprs       []xpr.Xpr
	}
	PltStmAsn struct {
		plt.StmScp
		X PltStmAct
	}
	PltStmAcs struct {
		plt.StmScp
	}
	PltFltsSctrAsn struct {
		plt.FltsSctrScp
		X PltFltsSctrAct
	}
	PltFltsSctrAcs struct {
		plt.FltsSctrScp
	}
	PltFltsSctrDistAsn struct {
		plt.FltsSctrDistScp
		X PltFltsSctrDistAct
	}
	PltFltsSctrDistAcs struct {
		plt.FltsSctrDistScp
	}
	PltHrzAsn struct {
		plt.HrzScp
		X PltHrzAct
	}
	PltHrzAcs struct {
		plt.HrzScp
	}
	PltVrtAsn struct {
		plt.VrtScp
		X PltVrtAct
	}
	PltVrtAcs struct {
		plt.VrtScp
	}
	PltDpthAsn struct {
		plt.DpthScp
		X PltDpthAct
	}
	PltDpthAcs struct {
		plt.DpthScp
	}
	SysMuAsn struct {
		sys.MuScp
		X SysMuAct
	}
	SysMuAcs struct {
		sys.MuScp
	}
	BolBolThen struct {
		X    BolBolAct
		Acts []Act
	}
	BolBolElse struct {
		X    BolBolAct
		Acts []Act
	}
	AnaPrfmPnlPctGet struct {
		X AnaPrfmAct
	}
	AnaPrfmScsPctGet struct {
		X AnaPrfmAct
	}
	AnaPrfmPipPerDayGet struct {
		X AnaPrfmAct
	}
	AnaPrfmUsdPerDayGet struct {
		X AnaPrfmAct
	}
	AnaPrfmScsPerDayGet struct {
		X AnaPrfmAct
	}
	AnaPrfmOpnPerDayGet struct {
		X AnaPrfmAct
	}
	AnaPrfmPnlUsdGet struct {
		X AnaPrfmAct
	}
	AnaPrfmPipAvgGet struct {
		X AnaPrfmAct
	}
	AnaPrfmPipMdnGet struct {
		X AnaPrfmAct
	}
	AnaPrfmPipMinGet struct {
		X AnaPrfmAct
	}
	AnaPrfmPipMaxGet struct {
		X AnaPrfmAct
	}
	AnaPrfmPipSumGet struct {
		X AnaPrfmAct
	}
	AnaPrfmDurAvgGet struct {
		X AnaPrfmAct
	}
	AnaPrfmDurMdnGet struct {
		X AnaPrfmAct
	}
	AnaPrfmDurMinGet struct {
		X AnaPrfmAct
	}
	AnaPrfmDurMaxGet struct {
		X AnaPrfmAct
	}
	AnaPrfmLosLimMaxGet struct {
		X AnaPrfmAct
	}
	AnaPrfmDurLimMaxGet struct {
		X AnaPrfmAct
	}
	AnaPrfmDayCntGet struct {
		X AnaPrfmAct
	}
	AnaPrfmTrdCntGet struct {
		X AnaPrfmAct
	}
	AnaPrfmTrdPctGet struct {
		X AnaPrfmAct
	}
	AnaPrfmCstTotUsdGet struct {
		X AnaPrfmAct
	}
	AnaPrfmCstSpdUsdGet struct {
		X AnaPrfmAct
	}
	AnaPrfmCstComUsdGet struct {
		X AnaPrfmAct
	}
	AnaPrfmDltPnlPctAGet struct {
		X AnaPrfmDltAct
	}
	AnaPrfmDltPnlPctBGet struct {
		X AnaPrfmDltAct
	}
	AnaPrfmDltPnlPctDltGet struct {
		X AnaPrfmDltAct
	}
	AnaPrfmDltScsPctAGet struct {
		X AnaPrfmDltAct
	}
	AnaPrfmDltScsPctBGet struct {
		X AnaPrfmDltAct
	}
	AnaPrfmDltScsPctDltGet struct {
		X AnaPrfmDltAct
	}
	AnaPrfmDltPipPerDayAGet struct {
		X AnaPrfmDltAct
	}
	AnaPrfmDltPipPerDayBGet struct {
		X AnaPrfmDltAct
	}
	AnaPrfmDltPipPerDayDltGet struct {
		X AnaPrfmDltAct
	}
	AnaPrfmDltUsdPerDayAGet struct {
		X AnaPrfmDltAct
	}
	AnaPrfmDltUsdPerDayBGet struct {
		X AnaPrfmDltAct
	}
	AnaPrfmDltUsdPerDayDltGet struct {
		X AnaPrfmDltAct
	}
	AnaPrfmDltScsPerDayAGet struct {
		X AnaPrfmDltAct
	}
	AnaPrfmDltScsPerDayBGet struct {
		X AnaPrfmDltAct
	}
	AnaPrfmDltScsPerDayDltGet struct {
		X AnaPrfmDltAct
	}
	AnaPrfmDltOpnPerDayAGet struct {
		X AnaPrfmDltAct
	}
	AnaPrfmDltOpnPerDayBGet struct {
		X AnaPrfmDltAct
	}
	AnaPrfmDltOpnPerDayDltGet struct {
		X AnaPrfmDltAct
	}
	AnaPrfmDltPnlUsdAGet struct {
		X AnaPrfmDltAct
	}
	AnaPrfmDltPnlUsdBGet struct {
		X AnaPrfmDltAct
	}
	AnaPrfmDltPnlUsdDltGet struct {
		X AnaPrfmDltAct
	}
	AnaPrfmDltPipAvgAGet struct {
		X AnaPrfmDltAct
	}
	AnaPrfmDltPipAvgBGet struct {
		X AnaPrfmDltAct
	}
	AnaPrfmDltPipAvgDltGet struct {
		X AnaPrfmDltAct
	}
	AnaPrfmDltPipMdnAGet struct {
		X AnaPrfmDltAct
	}
	AnaPrfmDltPipMdnBGet struct {
		X AnaPrfmDltAct
	}
	AnaPrfmDltPipMdnDltGet struct {
		X AnaPrfmDltAct
	}
	AnaPrfmDltPipMinAGet struct {
		X AnaPrfmDltAct
	}
	AnaPrfmDltPipMinBGet struct {
		X AnaPrfmDltAct
	}
	AnaPrfmDltPipMinDltGet struct {
		X AnaPrfmDltAct
	}
	AnaPrfmDltPipMaxAGet struct {
		X AnaPrfmDltAct
	}
	AnaPrfmDltPipMaxBGet struct {
		X AnaPrfmDltAct
	}
	AnaPrfmDltPipMaxDltGet struct {
		X AnaPrfmDltAct
	}
	AnaPrfmDltPipSumAGet struct {
		X AnaPrfmDltAct
	}
	AnaPrfmDltPipSumBGet struct {
		X AnaPrfmDltAct
	}
	AnaPrfmDltPipSumDltGet struct {
		X AnaPrfmDltAct
	}
	AnaPrfmDltDurAvgAGet struct {
		X AnaPrfmDltAct
	}
	AnaPrfmDltDurAvgBGet struct {
		X AnaPrfmDltAct
	}
	AnaPrfmDltDurAvgDltGet struct {
		X AnaPrfmDltAct
	}
	AnaPrfmDltDurMdnAGet struct {
		X AnaPrfmDltAct
	}
	AnaPrfmDltDurMdnBGet struct {
		X AnaPrfmDltAct
	}
	AnaPrfmDltDurMdnDltGet struct {
		X AnaPrfmDltAct
	}
	AnaPrfmDltDurMinAGet struct {
		X AnaPrfmDltAct
	}
	AnaPrfmDltDurMinBGet struct {
		X AnaPrfmDltAct
	}
	AnaPrfmDltDurMinDltGet struct {
		X AnaPrfmDltAct
	}
	AnaPrfmDltDurMaxAGet struct {
		X AnaPrfmDltAct
	}
	AnaPrfmDltDurMaxBGet struct {
		X AnaPrfmDltAct
	}
	AnaPrfmDltDurMaxDltGet struct {
		X AnaPrfmDltAct
	}
	AnaPrfmDltTrdCntAGet struct {
		X AnaPrfmDltAct
	}
	AnaPrfmDltTrdCntBGet struct {
		X AnaPrfmDltAct
	}
	AnaPrfmDltTrdCntDltGet struct {
		X AnaPrfmDltAct
	}
	AnaPrfmDltTrdPctAGet struct {
		X AnaPrfmDltAct
	}
	AnaPrfmDltTrdPctBGet struct {
		X AnaPrfmDltAct
	}
	AnaPrfmDltTrdPctDltGet struct {
		X AnaPrfmDltAct
	}
	AnaPrfmDltPthBGet struct {
		X AnaPrfmDltAct
	}
	PenPenClrSetGet struct {
		X  PenPenAct
		I0 ClrClrAct
	}
	PenPenWidSetGet struct {
		X  PenPenAct
		I0 UntUntAct
	}
	PltFltAxisYMinSetGet struct {
		X  PltFltAxisYAct
		I0 FltFltAct
	}
	PltFltAxisYMaxSetGet struct {
		X  PltFltAxisYAct
		I0 FltFltAct
	}
	PltFltAxisYEqiDstSetGet struct {
		X  PltFltAxisYAct
		I0 FltFltAct
	}
	PltStmTitleSetGet struct {
		X  PltStmAct
		I0 StrStrAct
	}
	PltFltsSctrYGet struct {
		X PltFltsSctrAct
	}
	PltFltsSctrTitleSetGet struct {
		X  PltFltsSctrAct
		I0 StrStrAct
	}
	PltFltsSctrOutlierSetGet struct {
		X  PltFltsSctrAct
		I0 BolBolAct
	}
	PltHrzPltsGet struct {
		X PltHrzAct
	}
	PltVrtPltsGet struct {
		X PltVrtAct
	}
	PltDpthPltsGet struct {
		X PltDpthAct
	}
	StrZero           struct{}
	StrEmpty          struct{}
	BolZero           struct{}
	BolFls            struct{}
	BolTru            struct{}
	FltZero           struct{}
	FltOne            struct{}
	FltNegOne         struct{}
	FltHndrd          struct{}
	FltMin            struct{}
	FltMax            struct{}
	FltTiny           struct{}
	UntZero           struct{}
	UntOne            struct{}
	UntMin            struct{}
	UntMax            struct{}
	IntZero           struct{}
	IntOne            struct{}
	IntNegOne         struct{}
	IntMin            struct{}
	IntMax            struct{}
	TmeZero           struct{}
	TmeOne            struct{}
	TmeNegOne         struct{}
	TmeMin            struct{}
	TmeMax            struct{}
	TmeSecond         struct{}
	TmeMinute         struct{}
	TmeHour           struct{}
	TmeDay            struct{}
	TmeWeek           struct{}
	TmeS1             struct{}
	TmeS5             struct{}
	TmeS10            struct{}
	TmeS15            struct{}
	TmeS20            struct{}
	TmeS30            struct{}
	TmeS40            struct{}
	TmeS50            struct{}
	TmeM1             struct{}
	TmeM5             struct{}
	TmeM10            struct{}
	TmeM15            struct{}
	TmeM20            struct{}
	TmeM30            struct{}
	TmeM40            struct{}
	TmeM50            struct{}
	TmeH1             struct{}
	TmeD1             struct{}
	TmeResolution     struct{}
	ClrBlack          struct{}
	ClrWhite          struct{}
	ClrRed50          struct{}
	ClrRed100         struct{}
	ClrRed200         struct{}
	ClrRed300         struct{}
	ClrRed400         struct{}
	ClrRed500         struct{}
	ClrRed600         struct{}
	ClrRed700         struct{}
	ClrRed800         struct{}
	ClrRed900         struct{}
	ClrRedA100        struct{}
	ClrRedA200        struct{}
	ClrRedA400        struct{}
	ClrRedA700        struct{}
	ClrPink50         struct{}
	ClrPink100        struct{}
	ClrPink200        struct{}
	ClrPink300        struct{}
	ClrPink400        struct{}
	ClrPink500        struct{}
	ClrPink600        struct{}
	ClrPink700        struct{}
	ClrPink800        struct{}
	ClrPink900        struct{}
	ClrPinkA100       struct{}
	ClrPinkA200       struct{}
	ClrPinkA400       struct{}
	ClrPinkA700       struct{}
	ClrPurple50       struct{}
	ClrPurple100      struct{}
	ClrPurple200      struct{}
	ClrPurple300      struct{}
	ClrPurple400      struct{}
	ClrPurple500      struct{}
	ClrPurple600      struct{}
	ClrPurple700      struct{}
	ClrPurple800      struct{}
	ClrPurple900      struct{}
	ClrPurpleA100     struct{}
	ClrPurpleA200     struct{}
	ClrPurpleA400     struct{}
	ClrPurpleA700     struct{}
	ClrDeepPurple50   struct{}
	ClrDeepPurple100  struct{}
	ClrDeepPurple200  struct{}
	ClrDeepPurple300  struct{}
	ClrDeepPurple400  struct{}
	ClrDeepPurple500  struct{}
	ClrDeepPurple600  struct{}
	ClrDeepPurple700  struct{}
	ClrDeepPurple800  struct{}
	ClrDeepPurple900  struct{}
	ClrDeepPurpleA100 struct{}
	ClrDeepPurpleA200 struct{}
	ClrDeepPurpleA400 struct{}
	ClrDeepPurpleA700 struct{}
	ClrIndigo50       struct{}
	ClrIndigo100      struct{}
	ClrIndigo200      struct{}
	ClrIndigo300      struct{}
	ClrIndigo400      struct{}
	ClrIndigo500      struct{}
	ClrIndigo600      struct{}
	ClrIndigo700      struct{}
	ClrIndigo800      struct{}
	ClrIndigo900      struct{}
	ClrIndigoA100     struct{}
	ClrIndigoA200     struct{}
	ClrIndigoA400     struct{}
	ClrIndigoA700     struct{}
	ClrBlue50         struct{}
	ClrBlue100        struct{}
	ClrBlue200        struct{}
	ClrBlue300        struct{}
	ClrBlue400        struct{}
	ClrBlue500        struct{}
	ClrBlue600        struct{}
	ClrBlue700        struct{}
	ClrBlue800        struct{}
	ClrBlue900        struct{}
	ClrBlueA100       struct{}
	ClrBlueA200       struct{}
	ClrBlueA400       struct{}
	ClrBlueA700       struct{}
	ClrLightBlue50    struct{}
	ClrLightBlue100   struct{}
	ClrLightBlue200   struct{}
	ClrLightBlue300   struct{}
	ClrLightBlue400   struct{}
	ClrLightBlue500   struct{}
	ClrLightBlue600   struct{}
	ClrLightBlue700   struct{}
	ClrLightBlue800   struct{}
	ClrLightBlue900   struct{}
	ClrLightBlueA100  struct{}
	ClrLightBlueA200  struct{}
	ClrLightBlueA400  struct{}
	ClrLightBlueA700  struct{}
	ClrCyan50         struct{}
	ClrCyan100        struct{}
	ClrCyan200        struct{}
	ClrCyan300        struct{}
	ClrCyan400        struct{}
	ClrCyan500        struct{}
	ClrCyan600        struct{}
	ClrCyan700        struct{}
	ClrCyan800        struct{}
	ClrCyan900        struct{}
	ClrCyanA100       struct{}
	ClrCyanA200       struct{}
	ClrCyanA400       struct{}
	ClrCyanA700       struct{}
	ClrTeal50         struct{}
	ClrTeal100        struct{}
	ClrTeal200        struct{}
	ClrTeal300        struct{}
	ClrTeal400        struct{}
	ClrTeal500        struct{}
	ClrTeal600        struct{}
	ClrTeal700        struct{}
	ClrTeal800        struct{}
	ClrTeal900        struct{}
	ClrTealA100       struct{}
	ClrTealA200       struct{}
	ClrTealA400       struct{}
	ClrTealA700       struct{}
	ClrGreen50        struct{}
	ClrGreen100       struct{}
	ClrGreen200       struct{}
	ClrGreen300       struct{}
	ClrGreen400       struct{}
	ClrGreen500       struct{}
	ClrGreen600       struct{}
	ClrGreen700       struct{}
	ClrGreen800       struct{}
	ClrGreen900       struct{}
	ClrGreenA100      struct{}
	ClrGreenA200      struct{}
	ClrGreenA400      struct{}
	ClrGreenA700      struct{}
	ClrLightGreen50   struct{}
	ClrLightGreen100  struct{}
	ClrLightGreen200  struct{}
	ClrLightGreen300  struct{}
	ClrLightGreen400  struct{}
	ClrLightGreen500  struct{}
	ClrLightGreen600  struct{}
	ClrLightGreen700  struct{}
	ClrLightGreen800  struct{}
	ClrLightGreen900  struct{}
	ClrLightGreenA100 struct{}
	ClrLightGreenA200 struct{}
	ClrLightGreenA400 struct{}
	ClrLightGreenA700 struct{}
	ClrLime50         struct{}
	ClrLime100        struct{}
	ClrLime200        struct{}
	ClrLime300        struct{}
	ClrLime400        struct{}
	ClrLime500        struct{}
	ClrLime600        struct{}
	ClrLime700        struct{}
	ClrLime800        struct{}
	ClrLime900        struct{}
	ClrLimeA100       struct{}
	ClrLimeA200       struct{}
	ClrLimeA400       struct{}
	ClrLimeA700       struct{}
	ClrYellow50       struct{}
	ClrYellow100      struct{}
	ClrYellow200      struct{}
	ClrYellow300      struct{}
	ClrYellow400      struct{}
	ClrYellow500      struct{}
	ClrYellow600      struct{}
	ClrYellow700      struct{}
	ClrYellow800      struct{}
	ClrYellow900      struct{}
	ClrYellowA100     struct{}
	ClrYellowA200     struct{}
	ClrYellowA400     struct{}
	ClrYellowA700     struct{}
	ClrAmber50        struct{}
	ClrAmber100       struct{}
	ClrAmber200       struct{}
	ClrAmber300       struct{}
	ClrAmber400       struct{}
	ClrAmber500       struct{}
	ClrAmber600       struct{}
	ClrAmber700       struct{}
	ClrAmber800       struct{}
	ClrAmber900       struct{}
	ClrAmberA100      struct{}
	ClrAmberA200      struct{}
	ClrAmberA400      struct{}
	ClrAmberA700      struct{}
	ClrOrange50       struct{}
	ClrOrange100      struct{}
	ClrOrange200      struct{}
	ClrOrange300      struct{}
	ClrOrange400      struct{}
	ClrOrange500      struct{}
	ClrOrange600      struct{}
	ClrOrange700      struct{}
	ClrOrange800      struct{}
	ClrOrange900      struct{}
	ClrOrangeA100     struct{}
	ClrOrangeA200     struct{}
	ClrOrangeA400     struct{}
	ClrOrangeA700     struct{}
	ClrDeepOrange50   struct{}
	ClrDeepOrange100  struct{}
	ClrDeepOrange200  struct{}
	ClrDeepOrange300  struct{}
	ClrDeepOrange400  struct{}
	ClrDeepOrange500  struct{}
	ClrDeepOrange600  struct{}
	ClrDeepOrange700  struct{}
	ClrDeepOrange800  struct{}
	ClrDeepOrange900  struct{}
	ClrDeepOrangeA100 struct{}
	ClrDeepOrangeA200 struct{}
	ClrDeepOrangeA400 struct{}
	ClrDeepOrangeA700 struct{}
	ClrBrown50        struct{}
	ClrBrown100       struct{}
	ClrBrown200       struct{}
	ClrBrown300       struct{}
	ClrBrown400       struct{}
	ClrBrown500       struct{}
	ClrBrown600       struct{}
	ClrBrown700       struct{}
	ClrBrown800       struct{}
	ClrBrown900       struct{}
	ClrGrey50         struct{}
	ClrGrey100        struct{}
	ClrGrey200        struct{}
	ClrGrey300        struct{}
	ClrGrey400        struct{}
	ClrGrey500        struct{}
	ClrGrey600        struct{}
	ClrGrey700        struct{}
	ClrGrey800        struct{}
	ClrGrey900        struct{}
	ClrBlueGrey50     struct{}
	ClrBlueGrey100    struct{}
	ClrBlueGrey200    struct{}
	ClrBlueGrey300    struct{}
	ClrBlueGrey400    struct{}
	ClrBlueGrey500    struct{}
	ClrBlueGrey600    struct{}
	ClrBlueGrey700    struct{}
	ClrBlueGrey800    struct{}
	ClrBlueGrey900    struct{}
	PenBlack          struct{}
	PenWhite          struct{}
	PenRed50          struct{}
	PenRed100         struct{}
	PenRed200         struct{}
	PenRed300         struct{}
	PenRed400         struct{}
	PenRed500         struct{}
	PenRed600         struct{}
	PenRed700         struct{}
	PenRed800         struct{}
	PenRed900         struct{}
	PenRedA100        struct{}
	PenRedA200        struct{}
	PenRedA400        struct{}
	PenRedA700        struct{}
	PenPink50         struct{}
	PenPink100        struct{}
	PenPink200        struct{}
	PenPink300        struct{}
	PenPink400        struct{}
	PenPink500        struct{}
	PenPink600        struct{}
	PenPink700        struct{}
	PenPink800        struct{}
	PenPink900        struct{}
	PenPinkA100       struct{}
	PenPinkA200       struct{}
	PenPinkA400       struct{}
	PenPinkA700       struct{}
	PenPurple50       struct{}
	PenPurple100      struct{}
	PenPurple200      struct{}
	PenPurple300      struct{}
	PenPurple400      struct{}
	PenPurple500      struct{}
	PenPurple600      struct{}
	PenPurple700      struct{}
	PenPurple800      struct{}
	PenPurple900      struct{}
	PenPurpleA100     struct{}
	PenPurpleA200     struct{}
	PenPurpleA400     struct{}
	PenPurpleA700     struct{}
	PenDeepPurple50   struct{}
	PenDeepPurple100  struct{}
	PenDeepPurple200  struct{}
	PenDeepPurple300  struct{}
	PenDeepPurple400  struct{}
	PenDeepPurple500  struct{}
	PenDeepPurple600  struct{}
	PenDeepPurple700  struct{}
	PenDeepPurple800  struct{}
	PenDeepPurple900  struct{}
	PenDeepPurpleA100 struct{}
	PenDeepPurpleA200 struct{}
	PenDeepPurpleA400 struct{}
	PenDeepPurpleA700 struct{}
	PenIndigo50       struct{}
	PenIndigo100      struct{}
	PenIndigo200      struct{}
	PenIndigo300      struct{}
	PenIndigo400      struct{}
	PenIndigo500      struct{}
	PenIndigo600      struct{}
	PenIndigo700      struct{}
	PenIndigo800      struct{}
	PenIndigo900      struct{}
	PenIndigoA100     struct{}
	PenIndigoA200     struct{}
	PenIndigoA400     struct{}
	PenIndigoA700     struct{}
	PenBlue50         struct{}
	PenBlue100        struct{}
	PenBlue200        struct{}
	PenBlue300        struct{}
	PenBlue400        struct{}
	PenBlue500        struct{}
	PenBlue600        struct{}
	PenBlue700        struct{}
	PenBlue800        struct{}
	PenBlue900        struct{}
	PenBlueA100       struct{}
	PenBlueA200       struct{}
	PenBlueA400       struct{}
	PenBlueA700       struct{}
	PenLightBlue50    struct{}
	PenLightBlue100   struct{}
	PenLightBlue200   struct{}
	PenLightBlue300   struct{}
	PenLightBlue400   struct{}
	PenLightBlue500   struct{}
	PenLightBlue600   struct{}
	PenLightBlue700   struct{}
	PenLightBlue800   struct{}
	PenLightBlue900   struct{}
	PenLightBlueA100  struct{}
	PenLightBlueA200  struct{}
	PenLightBlueA400  struct{}
	PenLightBlueA700  struct{}
	PenCyan50         struct{}
	PenCyan100        struct{}
	PenCyan200        struct{}
	PenCyan300        struct{}
	PenCyan400        struct{}
	PenCyan500        struct{}
	PenCyan600        struct{}
	PenCyan700        struct{}
	PenCyan800        struct{}
	PenCyan900        struct{}
	PenCyanA100       struct{}
	PenCyanA200       struct{}
	PenCyanA400       struct{}
	PenCyanA700       struct{}
	PenTeal50         struct{}
	PenTeal100        struct{}
	PenTeal200        struct{}
	PenTeal300        struct{}
	PenTeal400        struct{}
	PenTeal500        struct{}
	PenTeal600        struct{}
	PenTeal700        struct{}
	PenTeal800        struct{}
	PenTeal900        struct{}
	PenTealA100       struct{}
	PenTealA200       struct{}
	PenTealA400       struct{}
	PenTealA700       struct{}
	PenGreen50        struct{}
	PenGreen100       struct{}
	PenGreen200       struct{}
	PenGreen300       struct{}
	PenGreen400       struct{}
	PenGreen500       struct{}
	PenGreen600       struct{}
	PenGreen700       struct{}
	PenGreen800       struct{}
	PenGreen900       struct{}
	PenGreenA100      struct{}
	PenGreenA200      struct{}
	PenGreenA400      struct{}
	PenGreenA700      struct{}
	PenLightGreen50   struct{}
	PenLightGreen100  struct{}
	PenLightGreen200  struct{}
	PenLightGreen300  struct{}
	PenLightGreen400  struct{}
	PenLightGreen500  struct{}
	PenLightGreen600  struct{}
	PenLightGreen700  struct{}
	PenLightGreen800  struct{}
	PenLightGreen900  struct{}
	PenLightGreenA100 struct{}
	PenLightGreenA200 struct{}
	PenLightGreenA400 struct{}
	PenLightGreenA700 struct{}
	PenLime50         struct{}
	PenLime100        struct{}
	PenLime200        struct{}
	PenLime300        struct{}
	PenLime400        struct{}
	PenLime500        struct{}
	PenLime600        struct{}
	PenLime700        struct{}
	PenLime800        struct{}
	PenLime900        struct{}
	PenLimeA100       struct{}
	PenLimeA200       struct{}
	PenLimeA400       struct{}
	PenLimeA700       struct{}
	PenYellow50       struct{}
	PenYellow100      struct{}
	PenYellow200      struct{}
	PenYellow300      struct{}
	PenYellow400      struct{}
	PenYellow500      struct{}
	PenYellow600      struct{}
	PenYellow700      struct{}
	PenYellow800      struct{}
	PenYellow900      struct{}
	PenYellowA100     struct{}
	PenYellowA200     struct{}
	PenYellowA400     struct{}
	PenYellowA700     struct{}
	PenAmber50        struct{}
	PenAmber100       struct{}
	PenAmber200       struct{}
	PenAmber300       struct{}
	PenAmber400       struct{}
	PenAmber500       struct{}
	PenAmber600       struct{}
	PenAmber700       struct{}
	PenAmber800       struct{}
	PenAmber900       struct{}
	PenAmberA100      struct{}
	PenAmberA200      struct{}
	PenAmberA400      struct{}
	PenAmberA700      struct{}
	PenOrange50       struct{}
	PenOrange100      struct{}
	PenOrange200      struct{}
	PenOrange300      struct{}
	PenOrange400      struct{}
	PenOrange500      struct{}
	PenOrange600      struct{}
	PenOrange700      struct{}
	PenOrange800      struct{}
	PenOrange900      struct{}
	PenOrangeA100     struct{}
	PenOrangeA200     struct{}
	PenOrangeA400     struct{}
	PenOrangeA700     struct{}
	PenDeepOrange50   struct{}
	PenDeepOrange100  struct{}
	PenDeepOrange200  struct{}
	PenDeepOrange300  struct{}
	PenDeepOrange400  struct{}
	PenDeepOrange500  struct{}
	PenDeepOrange600  struct{}
	PenDeepOrange700  struct{}
	PenDeepOrange800  struct{}
	PenDeepOrange900  struct{}
	PenDeepOrangeA100 struct{}
	PenDeepOrangeA200 struct{}
	PenDeepOrangeA400 struct{}
	PenDeepOrangeA700 struct{}
	PenBrown50        struct{}
	PenBrown100       struct{}
	PenBrown200       struct{}
	PenBrown300       struct{}
	PenBrown400       struct{}
	PenBrown500       struct{}
	PenBrown600       struct{}
	PenBrown700       struct{}
	PenBrown800       struct{}
	PenBrown900       struct{}
	PenBlueGrey50     struct{}
	PenBlueGrey100    struct{}
	PenBlueGrey200    struct{}
	PenBlueGrey300    struct{}
	PenBlueGrey400    struct{}
	PenBlueGrey500    struct{}
	PenBlueGrey600    struct{}
	PenBlueGrey700    struct{}
	PenBlueGrey800    struct{}
	PenBlueGrey900    struct{}
	PenGrey50         struct{}
	PenGrey100        struct{}
	PenGrey200        struct{}
	PenGrey300        struct{}
	PenGrey400        struct{}
	PenGrey500        struct{}
	PenGrey600        struct{}
	PenGrey700        struct{}
	PenGrey800        struct{}
	PenGrey900        struct{}
	FltScl            struct{}
	UntStkWidth       struct{}
	UntShpRadius      struct{}
	UntAxisPad        struct{}
	UntBarPad         struct{}
	UntLen            struct{}
	UntPad            struct{}
	ClrBakClr         struct{}
	ClrBrdrClr        struct{}
	UntBrdrLen        struct{}
	UntInrvlTxtLen    struct{}
	ClrInrvlTxtClrX   struct{}
	ClrInrvlTxtClrY   struct{}
	ClrMsgClr         struct{}
	ClrTitleClr       struct{}
	ClrPrfClr         struct{}
	ClrLosClr         struct{}
	PenPrfPen         struct{}
	PenLosPen         struct{}
	FltOutlierLim     struct{}
	StrIfo            struct {
		I0 []Act
	}
	StrIfof struct {
		I0 StrStrAct
		I1 []Act
	}
	StrFmt struct {
		I0 StrStrAct
		I1 []Act
	}
	TmeNow    struct{}
	FltNewRng struct {
		I0 FltFltAct
		I1 FltFltAct
	}
	FltNewRngArnd struct {
		I0 FltFltAct
		I1 FltFltAct
	}
	FltNewRngFul struct{}
	TmeNewRng    struct {
		I0 TmeTmeAct
		I1 TmeTmeAct
	}
	TmeNewRngArnd struct {
		I0 TmeTmeAct
		I1 TmeTmeAct
	}
	TmeNewRngFul struct{}
	StrsNew      struct {
		I0 []StrStrAct
	}
	StrsMake struct {
		I0 UntUntAct
	}
	StrsMakeEmp struct {
		I0 UntUntAct
	}
	BolsNew struct {
		I0 []BolBolAct
	}
	BolsMake struct {
		I0 UntUntAct
	}
	BolsMakeEmp struct {
		I0 UntUntAct
	}
	FltsNew struct {
		I0 []FltFltAct
	}
	FltsMake struct {
		I0 UntUntAct
	}
	FltsMakeEmp struct {
		I0 UntUntAct
	}
	FltsAddsLss struct {
		I0 FltFltAct
		I1 FltFltAct
		I2 FltFltAct
	}
	FltsAddsLeq struct {
		I0 FltFltAct
		I1 FltFltAct
		I2 FltFltAct
	}
	FltsSubsGtr struct {
		I0 FltFltAct
		I1 FltFltAct
		I2 FltFltAct
	}
	FltsSubsGeq struct {
		I0 FltFltAct
		I1 FltFltAct
		I2 FltFltAct
	}
	FltsMulsLss struct {
		I0 FltFltAct
		I1 FltFltAct
		I2 FltFltAct
	}
	FltsMulsLeq struct {
		I0 FltFltAct
		I1 FltFltAct
		I2 FltFltAct
	}
	FltsDivsGtr struct {
		I0 FltFltAct
		I1 FltFltAct
		I2 FltFltAct
	}
	FltsDivsGeq struct {
		I0 FltFltAct
		I1 FltFltAct
		I2 FltFltAct
	}
	FltsFibsLeq struct {
		I0 FltFltAct
	}
	UntsNew struct {
		I0 []UntUntAct
	}
	UntsMake struct {
		I0 UntUntAct
	}
	UntsMakeEmp struct {
		I0 UntUntAct
	}
	UntsAddsLss struct {
		I0 UntUntAct
		I1 UntUntAct
		I2 UntUntAct
	}
	UntsAddsLeq struct {
		I0 UntUntAct
		I1 UntUntAct
		I2 UntUntAct
	}
	UntsSubsGtr struct {
		I0 UntUntAct
		I1 UntUntAct
		I2 UntUntAct
	}
	UntsSubsGeq struct {
		I0 UntUntAct
		I1 UntUntAct
		I2 UntUntAct
	}
	UntsMulsLss struct {
		I0 UntUntAct
		I1 UntUntAct
		I2 UntUntAct
	}
	UntsMulsLeq struct {
		I0 UntUntAct
		I1 UntUntAct
		I2 UntUntAct
	}
	UntsDivsGtr struct {
		I0 UntUntAct
		I1 UntUntAct
		I2 UntUntAct
	}
	UntsDivsGeq struct {
		I0 UntUntAct
		I1 UntUntAct
		I2 UntUntAct
	}
	UntsFibsLeq struct {
		I0 UntUntAct
	}
	IntsNew struct {
		I0 []IntIntAct
	}
	IntsMake struct {
		I0 UntUntAct
	}
	IntsMakeEmp struct {
		I0 UntUntAct
	}
	TmesNew struct {
		I0 []TmeTmeAct
	}
	TmesMake struct {
		I0 UntUntAct
	}
	TmesMakeEmp struct {
		I0 UntUntAct
	}
	TmesAddsLss struct {
		I0 TmeTmeAct
		I1 TmeTmeAct
		I2 TmeTmeAct
	}
	TmesAddsLeq struct {
		I0 TmeTmeAct
		I1 TmeTmeAct
		I2 TmeTmeAct
	}
	TmesSubsGtr struct {
		I0 TmeTmeAct
		I1 TmeTmeAct
		I2 TmeTmeAct
	}
	TmesSubsGeq struct {
		I0 TmeTmeAct
		I1 TmeTmeAct
		I2 TmeTmeAct
	}
	TmesMulsLss struct {
		I0 TmeTmeAct
		I1 TmeTmeAct
		I2 TmeTmeAct
	}
	TmesMulsLeq struct {
		I0 TmeTmeAct
		I1 TmeTmeAct
		I2 TmeTmeAct
	}
	TmesDivsGtr struct {
		I0 TmeTmeAct
		I1 TmeTmeAct
		I2 TmeTmeAct
	}
	TmesDivsGeq struct {
		I0 TmeTmeAct
		I1 TmeTmeAct
		I2 TmeTmeAct
	}
	TmesFibsLeq struct {
		I0 TmeTmeAct
	}
	BndsNew struct {
		I0 []BndBndAct
	}
	BndsMake struct {
		I0 UntUntAct
	}
	BndsMakeEmp struct {
		I0 UntUntAct
	}
	TmeNewRngs struct {
		I0 []TmeRngAct
	}
	TmeMakeRngs struct {
		I0 UntUntAct
	}
	TmeMakeEmpRngs struct {
		I0 UntUntAct
	}
	AnaNewTrds struct {
		I0 []AnaTrdAct
	}
	AnaMakeTrds struct {
		I0 UntUntAct
	}
	AnaMakeEmpTrds struct {
		I0 UntUntAct
	}
	AnaNewPrfms struct {
		I0 []AnaPrfmAct
	}
	AnaMakePrfms struct {
		I0 UntUntAct
	}
	AnaMakeEmpPrfms struct {
		I0 UntUntAct
	}
	HstOan     struct{}
	HstNewPrvs struct {
		I0 []HstPrvAct
	}
	HstMakePrvs struct {
		I0 UntUntAct
	}
	HstMakeEmpPrvs struct {
		I0 UntUntAct
	}
	HstNewInstrs struct {
		I0 []HstInstrAct
	}
	HstMakeInstrs struct {
		I0 UntUntAct
	}
	HstMakeEmpInstrs struct {
		I0 UntUntAct
	}
	HstNewInrvls struct {
		I0 []HstInrvlAct
	}
	HstMakeInrvls struct {
		I0 UntUntAct
	}
	HstMakeEmpInrvls struct {
		I0 UntUntAct
	}
	HstNewSides struct {
		I0 []HstSideAct
	}
	HstMakeSides struct {
		I0 UntUntAct
	}
	HstMakeEmpSides struct {
		I0 UntUntAct
	}
	HstNewStms struct {
		I0 []HstStmAct
	}
	HstMakeStms struct {
		I0 UntUntAct
	}
	HstMakeEmpStms struct {
		I0 UntUntAct
	}
	HstNewCnds struct {
		I0 []HstCndAct
	}
	HstMakeCnds struct {
		I0 UntUntAct
	}
	HstMakeEmpCnds struct {
		I0 UntUntAct
	}
	HstNewStgys struct {
		I0 []HstStgyAct
	}
	HstMakeStgys struct {
		I0 UntUntAct
	}
	HstMakeEmpStgys struct {
		I0 UntUntAct
	}
	RltOan     struct{}
	RltNewPrvs struct {
		I0 []RltPrvAct
	}
	RltMakePrvs struct {
		I0 UntUntAct
	}
	RltMakeEmpPrvs struct {
		I0 UntUntAct
	}
	RltNewInstrs struct {
		I0 []RltInstrAct
	}
	RltMakeInstrs struct {
		I0 UntUntAct
	}
	RltMakeEmpInstrs struct {
		I0 UntUntAct
	}
	RltNewInrvls struct {
		I0 []RltInrvlAct
	}
	RltMakeInrvls struct {
		I0 UntUntAct
	}
	RltMakeEmpInrvls struct {
		I0 UntUntAct
	}
	RltNewSides struct {
		I0 []RltSideAct
	}
	RltMakeSides struct {
		I0 UntUntAct
	}
	RltMakeEmpSides struct {
		I0 UntUntAct
	}
	RltNewStms struct {
		I0 []RltStmAct
	}
	RltMakeStms struct {
		I0 UntUntAct
	}
	RltMakeEmpStms struct {
		I0 UntUntAct
	}
	RltNewCnds struct {
		I0 []RltCndAct
	}
	RltMakeCnds struct {
		I0 UntUntAct
	}
	RltMakeEmpCnds struct {
		I0 UntUntAct
	}
	RltNewStgys struct {
		I0 []RltStgyAct
	}
	RltMakeStgys struct {
		I0 UntUntAct
	}
	RltMakeEmpStgys struct {
		I0 UntUntAct
	}
	ClrRgba struct {
		I0 FltFltAct
		I1 FltFltAct
		I2 FltFltAct
		I3 FltFltAct
	}
	ClrRgb struct {
		I0 FltFltAct
		I1 FltFltAct
		I2 FltFltAct
	}
	ClrHex struct {
		I0 StrStrAct
	}
	PenNew struct {
		I0 ClrClrAct
		I1 []UntUntAct
	}
	PenRgba struct {
		I0 FltFltAct
		I1 FltFltAct
		I2 FltFltAct
		I3 FltFltAct
		I4 []UntUntAct
	}
	PenRgb struct {
		I0 FltFltAct
		I1 FltFltAct
		I2 FltFltAct
		I3 []UntUntAct
	}
	PenHex struct {
		I0 StrStrAct
		I1 []UntUntAct
	}
	PenNewPens struct {
		I0 []PenPenAct
	}
	PenMakePens struct {
		I0 UntUntAct
	}
	PenMakeEmpPens struct {
		I0 UntUntAct
	}
	PltNewPlts struct {
		I0 []PltPltAct
	}
	PltMakePlts struct {
		I0 UntUntAct
	}
	PltMakeEmpPlts struct {
		I0 UntUntAct
	}
	PltNewStm          struct{}
	PltNewFltsSctr     struct{}
	PltNewFltsSctrDist struct{}
	PltNewHrz          struct {
		I0 []PltPltAct
	}
	PltNewVrt struct {
		I0 []PltPltAct
	}
	PltNewDpth struct {
		I0 []PltPltAct
	}
	SysNewMu    struct{}
	StrStrLower struct {
		X StrStrAct
	}
	StrStrUpper struct {
		X StrStrAct
	}
	StrStrEql struct {
		X  StrStrAct
		I0 StrStrAct
	}
	StrStrNeq struct {
		X  StrStrAct
		I0 StrStrAct
	}
	StrStrLss struct {
		X  StrStrAct
		I0 StrStrAct
	}
	StrStrGtr struct {
		X  StrStrAct
		I0 StrStrAct
	}
	StrStrLeq struct {
		X  StrStrAct
		I0 StrStrAct
	}
	StrStrGeq struct {
		X  StrStrAct
		I0 StrStrAct
	}
	BolBolNot struct {
		X BolBolAct
	}
	BolBolEql struct {
		X  BolBolAct
		I0 BolBolAct
	}
	BolBolNeq struct {
		X  BolBolAct
		I0 BolBolAct
	}
	FltFltEql struct {
		X  FltFltAct
		I0 FltFltAct
	}
	FltFltNeq struct {
		X  FltFltAct
		I0 FltFltAct
	}
	FltFltTrnc struct {
		X  FltFltAct
		I0 UntUntAct
	}
	FltFltIsNaN struct {
		X FltFltAct
	}
	FltFltIsInfPos struct {
		X FltFltAct
	}
	FltFltIsInfNeg struct {
		X FltFltAct
	}
	FltFltIsValid struct {
		X FltFltAct
	}
	FltFltPct struct {
		X  FltFltAct
		I0 FltFltAct
	}
	FltFltLss struct {
		X  FltFltAct
		I0 FltFltAct
	}
	FltFltGtr struct {
		X  FltFltAct
		I0 FltFltAct
	}
	FltFltLeq struct {
		X  FltFltAct
		I0 FltFltAct
	}
	FltFltGeq struct {
		X  FltFltAct
		I0 FltFltAct
	}
	FltFltPos struct {
		X FltFltAct
	}
	FltFltNeg struct {
		X FltFltAct
	}
	FltFltInv struct {
		X FltFltAct
	}
	FltFltAdd struct {
		X  FltFltAct
		I0 FltFltAct
	}
	FltFltSub struct {
		X  FltFltAct
		I0 FltFltAct
	}
	FltFltMul struct {
		X  FltFltAct
		I0 FltFltAct
	}
	FltFltDiv struct {
		X  FltFltAct
		I0 FltFltAct
	}
	FltFltRem struct {
		X  FltFltAct
		I0 FltFltAct
	}
	FltFltPow struct {
		X  FltFltAct
		I0 FltFltAct
	}
	FltFltSqr struct {
		X FltFltAct
	}
	FltFltSqrt struct {
		X FltFltAct
	}
	FltFltMin struct {
		X  FltFltAct
		I0 FltFltAct
	}
	FltFltMax struct {
		X  FltFltAct
		I0 FltFltAct
	}
	FltFltMid struct {
		X  FltFltAct
		I0 FltFltAct
	}
	FltFltAvg struct {
		X  FltFltAct
		I0 FltFltAct
	}
	FltFltAvgGeo struct {
		X  FltFltAct
		I0 FltFltAct
	}
	FltFltSelEql struct {
		X  FltFltAct
		I0 FltFltAct
	}
	FltFltSelNeq struct {
		X  FltFltAct
		I0 FltFltAct
	}
	FltFltSelLss struct {
		X  FltFltAct
		I0 FltFltAct
	}
	FltFltSelGtr struct {
		X  FltFltAct
		I0 FltFltAct
	}
	FltFltSelLeq struct {
		X  FltFltAct
		I0 FltFltAct
	}
	FltFltSelGeq struct {
		X  FltFltAct
		I0 FltFltAct
	}
	UntUntEql struct {
		X  UntUntAct
		I0 UntUntAct
	}
	UntUntNeq struct {
		X  UntUntAct
		I0 UntUntAct
	}
	UntUntLss struct {
		X  UntUntAct
		I0 UntUntAct
	}
	UntUntGtr struct {
		X  UntUntAct
		I0 UntUntAct
	}
	UntUntLeq struct {
		X  UntUntAct
		I0 UntUntAct
	}
	UntUntGeq struct {
		X  UntUntAct
		I0 UntUntAct
	}
	UntUntAdd struct {
		X  UntUntAct
		I0 UntUntAct
	}
	UntUntSub struct {
		X  UntUntAct
		I0 UntUntAct
	}
	UntUntMul struct {
		X  UntUntAct
		I0 UntUntAct
	}
	UntUntDiv struct {
		X  UntUntAct
		I0 UntUntAct
	}
	UntUntRem struct {
		X  UntUntAct
		I0 UntUntAct
	}
	UntUntPow struct {
		X  UntUntAct
		I0 UntUntAct
	}
	UntUntSqr struct {
		X UntUntAct
	}
	UntUntSqrt struct {
		X UntUntAct
	}
	UntUntMin struct {
		X  UntUntAct
		I0 UntUntAct
	}
	UntUntMax struct {
		X  UntUntAct
		I0 UntUntAct
	}
	UntUntMid struct {
		X  UntUntAct
		I0 UntUntAct
	}
	UntUntAvg struct {
		X  UntUntAct
		I0 UntUntAct
	}
	UntUntAvgGeo struct {
		X  UntUntAct
		I0 UntUntAct
	}
	IntIntEql struct {
		X  IntIntAct
		I0 IntIntAct
	}
	IntIntNeq struct {
		X  IntIntAct
		I0 IntIntAct
	}
	IntIntLss struct {
		X  IntIntAct
		I0 IntIntAct
	}
	IntIntGtr struct {
		X  IntIntAct
		I0 IntIntAct
	}
	IntIntLeq struct {
		X  IntIntAct
		I0 IntIntAct
	}
	IntIntGeq struct {
		X  IntIntAct
		I0 IntIntAct
	}
	IntIntPos struct {
		X IntIntAct
	}
	IntIntNeg struct {
		X IntIntAct
	}
	IntIntInv struct {
		X IntIntAct
	}
	IntIntAdd struct {
		X  IntIntAct
		I0 IntIntAct
	}
	IntIntSub struct {
		X  IntIntAct
		I0 IntIntAct
	}
	IntIntMul struct {
		X  IntIntAct
		I0 IntIntAct
	}
	IntIntDiv struct {
		X  IntIntAct
		I0 IntIntAct
	}
	IntIntRem struct {
		X  IntIntAct
		I0 IntIntAct
	}
	IntIntPow struct {
		X  IntIntAct
		I0 IntIntAct
	}
	IntIntSqr struct {
		X IntIntAct
	}
	IntIntSqrt struct {
		X IntIntAct
	}
	IntIntMin struct {
		X  IntIntAct
		I0 IntIntAct
	}
	IntIntMax struct {
		X  IntIntAct
		I0 IntIntAct
	}
	IntIntMid struct {
		X  IntIntAct
		I0 IntIntAct
	}
	IntIntAvg struct {
		X  IntIntAct
		I0 IntIntAct
	}
	IntIntAvgGeo struct {
		X  IntIntAct
		I0 IntIntAct
	}
	TmeTmeWeekdayCnt struct {
		X  TmeTmeAct
		I0 TmeTmeAct
	}
	TmeTmeDte struct {
		X TmeTmeAct
	}
	TmeTmeToSunday struct {
		X TmeTmeAct
	}
	TmeTmeToMonday struct {
		X TmeTmeAct
	}
	TmeTmeToTuesday struct {
		X TmeTmeAct
	}
	TmeTmeToWednesday struct {
		X TmeTmeAct
	}
	TmeTmeToThursday struct {
		X TmeTmeAct
	}
	TmeTmeToFriday struct {
		X TmeTmeAct
	}
	TmeTmeToSaturday struct {
		X TmeTmeAct
	}
	TmeTmeIsSunday struct {
		X TmeTmeAct
	}
	TmeTmeIsMonday struct {
		X TmeTmeAct
	}
	TmeTmeIsTuesday struct {
		X TmeTmeAct
	}
	TmeTmeIsWednesday struct {
		X TmeTmeAct
	}
	TmeTmeIsThursday struct {
		X TmeTmeAct
	}
	TmeTmeIsFriday struct {
		X TmeTmeAct
	}
	TmeTmeIsSaturday struct {
		X TmeTmeAct
	}
	TmeTmeEql struct {
		X  TmeTmeAct
		I0 TmeTmeAct
	}
	TmeTmeNeq struct {
		X  TmeTmeAct
		I0 TmeTmeAct
	}
	TmeTmeLss struct {
		X  TmeTmeAct
		I0 TmeTmeAct
	}
	TmeTmeGtr struct {
		X  TmeTmeAct
		I0 TmeTmeAct
	}
	TmeTmeLeq struct {
		X  TmeTmeAct
		I0 TmeTmeAct
	}
	TmeTmeGeq struct {
		X  TmeTmeAct
		I0 TmeTmeAct
	}
	TmeTmePos struct {
		X TmeTmeAct
	}
	TmeTmeNeg struct {
		X TmeTmeAct
	}
	TmeTmeInv struct {
		X TmeTmeAct
	}
	TmeTmeAdd struct {
		X  TmeTmeAct
		I0 TmeTmeAct
	}
	TmeTmeSub struct {
		X  TmeTmeAct
		I0 TmeTmeAct
	}
	TmeTmeMul struct {
		X  TmeTmeAct
		I0 TmeTmeAct
	}
	TmeTmeDiv struct {
		X  TmeTmeAct
		I0 TmeTmeAct
	}
	TmeTmeRem struct {
		X  TmeTmeAct
		I0 TmeTmeAct
	}
	TmeTmePow struct {
		X  TmeTmeAct
		I0 TmeTmeAct
	}
	TmeTmeSqr struct {
		X TmeTmeAct
	}
	TmeTmeSqrt struct {
		X TmeTmeAct
	}
	TmeTmeMin struct {
		X  TmeTmeAct
		I0 TmeTmeAct
	}
	TmeTmeMax struct {
		X  TmeTmeAct
		I0 TmeTmeAct
	}
	TmeTmeMid struct {
		X  TmeTmeAct
		I0 TmeTmeAct
	}
	TmeTmeAvg struct {
		X  TmeTmeAct
		I0 TmeTmeAct
	}
	TmeTmeAvgGeo struct {
		X  TmeTmeAct
		I0 TmeTmeAct
	}
	BndBndCnt struct {
		X BndBndAct
	}
	BndBndLen struct {
		X BndBndAct
	}
	BndBndLstIdx struct {
		X BndBndAct
	}
	BndBndIsValid struct {
		X BndBndAct
	}
	FltRngLen struct {
		X FltRngAct
	}
	FltRngIsValid struct {
		X FltRngAct
	}
	FltRngEnsure struct {
		X FltRngAct
	}
	FltRngMinSub struct {
		X  FltRngAct
		I0 FltFltAct
	}
	FltRngMaxAdd struct {
		X  FltRngAct
		I0 FltFltAct
	}
	FltRngMrg struct {
		X  FltRngAct
		I0 FltRngAct
	}
	TmeRngLen struct {
		X TmeRngAct
	}
	TmeRngIsValid struct {
		X TmeRngAct
	}
	TmeRngEnsure struct {
		X TmeRngAct
	}
	TmeRngMinSub struct {
		X  TmeRngAct
		I0 TmeTmeAct
	}
	TmeRngMaxAdd struct {
		X  TmeRngAct
		I0 TmeTmeAct
	}
	TmeRngMrg struct {
		X  TmeRngAct
		I0 TmeRngAct
	}
	StrsStrsCnt struct {
		X StrsStrsAct
	}
	StrsStrsCpy struct {
		X StrsStrsAct
	}
	StrsStrsClr struct {
		X StrsStrsAct
	}
	StrsStrsRand struct {
		X StrsStrsAct
	}
	StrsStrsMrg struct {
		X  StrsStrsAct
		I0 []StrsStrsAct
	}
	StrsStrsPush struct {
		X  StrsStrsAct
		I0 []StrStrAct
	}
	StrsStrsPop struct {
		X StrsStrsAct
	}
	StrsStrsQue struct {
		X  StrsStrsAct
		I0 []StrStrAct
	}
	StrsStrsDque struct {
		X StrsStrsAct
	}
	StrsStrsIns struct {
		X  StrsStrsAct
		I0 UntUntAct
		I1 StrStrAct
	}
	StrsStrsUpd struct {
		X  StrsStrsAct
		I0 UntUntAct
		I1 StrStrAct
	}
	StrsStrsDel struct {
		X  StrsStrsAct
		I0 UntUntAct
	}
	StrsStrsAt struct {
		X  StrsStrsAct
		I0 UntUntAct
	}
	StrsStrsIn struct {
		X  StrsStrsAct
		I0 UntUntAct
		I1 UntUntAct
	}
	StrsStrsInBnd struct {
		X  StrsStrsAct
		I0 BndBndAct
	}
	StrsStrsFrom struct {
		X  StrsStrsAct
		I0 UntUntAct
	}
	StrsStrsTo struct {
		X  StrsStrsAct
		I0 UntUntAct
	}
	StrsStrsFst struct {
		X StrsStrsAct
	}
	StrsStrsMdl struct {
		X StrsStrsAct
	}
	StrsStrsLst struct {
		X StrsStrsAct
	}
	StrsStrsFstIdx struct {
		X StrsStrsAct
	}
	StrsStrsMdlIdx struct {
		X StrsStrsAct
	}
	StrsStrsLstIdx struct {
		X StrsStrsAct
	}
	StrsStrsRev struct {
		X StrsStrsAct
	}
	StrsStrsSrchIdxEql struct {
		X  StrsStrsAct
		I0 StrStrAct
	}
	StrsStrsSrchIdx struct {
		X  StrsStrsAct
		I0 StrStrAct
		I1 []BolBolAct
	}
	StrsStrsHas struct {
		X  StrsStrsAct
		I0 StrStrAct
	}
	StrsStrsSrtAsc struct {
		X StrsStrsAct
	}
	StrsStrsSrtDsc struct {
		X StrsStrsAct
	}
	BolsBolsCnt struct {
		X BolsBolsAct
	}
	BolsBolsCpy struct {
		X BolsBolsAct
	}
	BolsBolsClr struct {
		X BolsBolsAct
	}
	BolsBolsRand struct {
		X BolsBolsAct
	}
	BolsBolsMrg struct {
		X  BolsBolsAct
		I0 []BolsBolsAct
	}
	BolsBolsPush struct {
		X  BolsBolsAct
		I0 []BolBolAct
	}
	BolsBolsPop struct {
		X BolsBolsAct
	}
	BolsBolsQue struct {
		X  BolsBolsAct
		I0 []BolBolAct
	}
	BolsBolsDque struct {
		X BolsBolsAct
	}
	BolsBolsIns struct {
		X  BolsBolsAct
		I0 UntUntAct
		I1 BolBolAct
	}
	BolsBolsUpd struct {
		X  BolsBolsAct
		I0 UntUntAct
		I1 BolBolAct
	}
	BolsBolsDel struct {
		X  BolsBolsAct
		I0 UntUntAct
	}
	BolsBolsAt struct {
		X  BolsBolsAct
		I0 UntUntAct
	}
	BolsBolsIn struct {
		X  BolsBolsAct
		I0 UntUntAct
		I1 UntUntAct
	}
	BolsBolsInBnd struct {
		X  BolsBolsAct
		I0 BndBndAct
	}
	BolsBolsFrom struct {
		X  BolsBolsAct
		I0 UntUntAct
	}
	BolsBolsTo struct {
		X  BolsBolsAct
		I0 UntUntAct
	}
	BolsBolsFst struct {
		X BolsBolsAct
	}
	BolsBolsMdl struct {
		X BolsBolsAct
	}
	BolsBolsLst struct {
		X BolsBolsAct
	}
	BolsBolsFstIdx struct {
		X BolsBolsAct
	}
	BolsBolsMdlIdx struct {
		X BolsBolsAct
	}
	BolsBolsLstIdx struct {
		X BolsBolsAct
	}
	BolsBolsRev struct {
		X BolsBolsAct
	}
	FltsFltsCnt struct {
		X FltsFltsAct
	}
	FltsFltsCpy struct {
		X FltsFltsAct
	}
	FltsFltsClr struct {
		X FltsFltsAct
	}
	FltsFltsRand struct {
		X FltsFltsAct
	}
	FltsFltsMrg struct {
		X  FltsFltsAct
		I0 []FltsFltsAct
	}
	FltsFltsPush struct {
		X  FltsFltsAct
		I0 []FltFltAct
	}
	FltsFltsPop struct {
		X FltsFltsAct
	}
	FltsFltsQue struct {
		X  FltsFltsAct
		I0 []FltFltAct
	}
	FltsFltsDque struct {
		X FltsFltsAct
	}
	FltsFltsIns struct {
		X  FltsFltsAct
		I0 UntUntAct
		I1 FltFltAct
	}
	FltsFltsUpd struct {
		X  FltsFltsAct
		I0 UntUntAct
		I1 FltFltAct
	}
	FltsFltsDel struct {
		X  FltsFltsAct
		I0 UntUntAct
	}
	FltsFltsAt struct {
		X  FltsFltsAct
		I0 UntUntAct
	}
	FltsFltsIn struct {
		X  FltsFltsAct
		I0 UntUntAct
		I1 UntUntAct
	}
	FltsFltsInBnd struct {
		X  FltsFltsAct
		I0 BndBndAct
	}
	FltsFltsFrom struct {
		X  FltsFltsAct
		I0 UntUntAct
	}
	FltsFltsTo struct {
		X  FltsFltsAct
		I0 UntUntAct
	}
	FltsFltsFst struct {
		X FltsFltsAct
	}
	FltsFltsMdl struct {
		X FltsFltsAct
	}
	FltsFltsLst struct {
		X FltsFltsAct
	}
	FltsFltsFstIdx struct {
		X FltsFltsAct
	}
	FltsFltsMdlIdx struct {
		X FltsFltsAct
	}
	FltsFltsLstIdx struct {
		X FltsFltsAct
	}
	FltsFltsRev struct {
		X FltsFltsAct
	}
	FltsFltsSrchIdxEql struct {
		X  FltsFltsAct
		I0 FltFltAct
	}
	FltsFltsSrchIdx struct {
		X  FltsFltsAct
		I0 FltFltAct
		I1 []BolBolAct
	}
	FltsFltsHas struct {
		X  FltsFltsAct
		I0 FltFltAct
	}
	FltsFltsSrtAsc struct {
		X FltsFltsAct
	}
	FltsFltsSrtDsc struct {
		X FltsFltsAct
	}
	FltsFltsUnaPos struct {
		X FltsFltsAct
	}
	FltsFltsUnaNeg struct {
		X FltsFltsAct
	}
	FltsFltsUnaInv struct {
		X FltsFltsAct
	}
	FltsFltsUnaSqr struct {
		X FltsFltsAct
	}
	FltsFltsUnaSqrt struct {
		X FltsFltsAct
	}
	FltsFltsSclAdd struct {
		X  FltsFltsAct
		I0 FltFltAct
	}
	FltsFltsSclSub struct {
		X  FltsFltsAct
		I0 FltFltAct
	}
	FltsFltsSclMul struct {
		X  FltsFltsAct
		I0 FltFltAct
	}
	FltsFltsSclDiv struct {
		X  FltsFltsAct
		I0 FltFltAct
	}
	FltsFltsSclRem struct {
		X  FltsFltsAct
		I0 FltFltAct
	}
	FltsFltsSclPow struct {
		X  FltsFltsAct
		I0 FltFltAct
	}
	FltsFltsSclMin struct {
		X  FltsFltsAct
		I0 FltFltAct
	}
	FltsFltsSclMax struct {
		X  FltsFltsAct
		I0 FltFltAct
	}
	FltsFltsSelEql struct {
		X  FltsFltsAct
		I0 FltFltAct
	}
	FltsFltsSelNeq struct {
		X  FltsFltsAct
		I0 FltFltAct
	}
	FltsFltsSelLss struct {
		X  FltsFltsAct
		I0 FltFltAct
	}
	FltsFltsSelGtr struct {
		X  FltsFltsAct
		I0 FltFltAct
	}
	FltsFltsSelLeq struct {
		X  FltsFltsAct
		I0 FltFltAct
	}
	FltsFltsSelGeq struct {
		X  FltsFltsAct
		I0 FltFltAct
	}
	FltsFltsCntEql struct {
		X  FltsFltsAct
		I0 FltFltAct
	}
	FltsFltsCntNeq struct {
		X  FltsFltsAct
		I0 FltFltAct
	}
	FltsFltsCntLss struct {
		X  FltsFltsAct
		I0 FltFltAct
	}
	FltsFltsCntGtr struct {
		X  FltsFltsAct
		I0 FltFltAct
	}
	FltsFltsCntLeq struct {
		X  FltsFltsAct
		I0 FltFltAct
	}
	FltsFltsCntGeq struct {
		X  FltsFltsAct
		I0 FltFltAct
	}
	FltsFltsInrAdd struct {
		X  FltsFltsAct
		I0 UntUntAct
	}
	FltsFltsInrSub struct {
		X  FltsFltsAct
		I0 UntUntAct
	}
	FltsFltsInrMul struct {
		X  FltsFltsAct
		I0 UntUntAct
	}
	FltsFltsInrDiv struct {
		X  FltsFltsAct
		I0 UntUntAct
	}
	FltsFltsInrRem struct {
		X  FltsFltsAct
		I0 UntUntAct
	}
	FltsFltsInrPow struct {
		X  FltsFltsAct
		I0 UntUntAct
	}
	FltsFltsInrMin struct {
		X  FltsFltsAct
		I0 UntUntAct
	}
	FltsFltsInrMax struct {
		X  FltsFltsAct
		I0 UntUntAct
	}
	FltsFltsSum struct {
		X FltsFltsAct
	}
	FltsFltsPrd struct {
		X FltsFltsAct
	}
	FltsFltsMin struct {
		X FltsFltsAct
	}
	FltsFltsMax struct {
		X FltsFltsAct
	}
	FltsFltsMid struct {
		X FltsFltsAct
	}
	FltsFltsMdn struct {
		X FltsFltsAct
	}
	FltsFltsSma struct {
		X FltsFltsAct
	}
	FltsFltsGma struct {
		X FltsFltsAct
	}
	FltsFltsWma struct {
		X FltsFltsAct
	}
	FltsFltsVrnc struct {
		X FltsFltsAct
	}
	FltsFltsStd struct {
		X FltsFltsAct
	}
	FltsFltsZscr struct {
		X FltsFltsAct
	}
	FltsFltsZscrInplace struct {
		X FltsFltsAct
	}
	FltsFltsRngFul struct {
		X FltsFltsAct
	}
	FltsFltsRngLst struct {
		X FltsFltsAct
	}
	FltsFltsProLst struct {
		X FltsFltsAct
	}
	FltsFltsProSma struct {
		X FltsFltsAct
	}
	FltsFltsSubSumPos struct {
		X FltsFltsAct
	}
	FltsFltsSubSumNeg struct {
		X FltsFltsAct
	}
	FltsFltsRsi struct {
		X FltsFltsAct
	}
	FltsFltsWrsi struct {
		X FltsFltsAct
	}
	FltsFltsPro struct {
		X FltsFltsAct
	}
	FltsFltsAlma struct {
		X FltsFltsAct
	}
	FltsFltsProAlma struct {
		X FltsFltsAct
	}
	FltsFltsCntrDist struct {
		X  FltsFltsAct
		I0 []BolBolAct
	}
	UntsUntsCnt struct {
		X UntsUntsAct
	}
	UntsUntsCpy struct {
		X UntsUntsAct
	}
	UntsUntsClr struct {
		X UntsUntsAct
	}
	UntsUntsRand struct {
		X UntsUntsAct
	}
	UntsUntsMrg struct {
		X  UntsUntsAct
		I0 []UntsUntsAct
	}
	UntsUntsPush struct {
		X  UntsUntsAct
		I0 []UntUntAct
	}
	UntsUntsPop struct {
		X UntsUntsAct
	}
	UntsUntsQue struct {
		X  UntsUntsAct
		I0 []UntUntAct
	}
	UntsUntsDque struct {
		X UntsUntsAct
	}
	UntsUntsIns struct {
		X  UntsUntsAct
		I0 UntUntAct
		I1 UntUntAct
	}
	UntsUntsUpd struct {
		X  UntsUntsAct
		I0 UntUntAct
		I1 UntUntAct
	}
	UntsUntsDel struct {
		X  UntsUntsAct
		I0 UntUntAct
	}
	UntsUntsAt struct {
		X  UntsUntsAct
		I0 UntUntAct
	}
	UntsUntsIn struct {
		X  UntsUntsAct
		I0 UntUntAct
		I1 UntUntAct
	}
	UntsUntsInBnd struct {
		X  UntsUntsAct
		I0 BndBndAct
	}
	UntsUntsFrom struct {
		X  UntsUntsAct
		I0 UntUntAct
	}
	UntsUntsTo struct {
		X  UntsUntsAct
		I0 UntUntAct
	}
	UntsUntsFst struct {
		X UntsUntsAct
	}
	UntsUntsMdl struct {
		X UntsUntsAct
	}
	UntsUntsLst struct {
		X UntsUntsAct
	}
	UntsUntsFstIdx struct {
		X UntsUntsAct
	}
	UntsUntsMdlIdx struct {
		X UntsUntsAct
	}
	UntsUntsLstIdx struct {
		X UntsUntsAct
	}
	UntsUntsRev struct {
		X UntsUntsAct
	}
	UntsUntsSrchIdxEql struct {
		X  UntsUntsAct
		I0 UntUntAct
	}
	UntsUntsSrchIdx struct {
		X  UntsUntsAct
		I0 UntUntAct
		I1 []BolBolAct
	}
	UntsUntsHas struct {
		X  UntsUntsAct
		I0 UntUntAct
	}
	UntsUntsSrtAsc struct {
		X UntsUntsAct
	}
	UntsUntsSrtDsc struct {
		X UntsUntsAct
	}
	UntsUntsInrAdd struct {
		X  UntsUntsAct
		I0 UntUntAct
	}
	UntsUntsInrSub struct {
		X  UntsUntsAct
		I0 UntUntAct
	}
	UntsUntsInrMul struct {
		X  UntsUntsAct
		I0 UntUntAct
	}
	UntsUntsInrDiv struct {
		X  UntsUntsAct
		I0 UntUntAct
	}
	UntsUntsInrRem struct {
		X  UntsUntsAct
		I0 UntUntAct
	}
	UntsUntsInrPow struct {
		X  UntsUntsAct
		I0 UntUntAct
	}
	UntsUntsInrMin struct {
		X  UntsUntsAct
		I0 UntUntAct
	}
	UntsUntsInrMax struct {
		X  UntsUntsAct
		I0 UntUntAct
	}
	UntsUntsSum struct {
		X UntsUntsAct
	}
	UntsUntsPrd struct {
		X UntsUntsAct
	}
	UntsUntsMin struct {
		X UntsUntsAct
	}
	UntsUntsMax struct {
		X UntsUntsAct
	}
	UntsUntsMid struct {
		X UntsUntsAct
	}
	UntsUntsMdn struct {
		X UntsUntsAct
	}
	UntsUntsSma struct {
		X UntsUntsAct
	}
	UntsUntsGma struct {
		X UntsUntsAct
	}
	UntsUntsWma struct {
		X UntsUntsAct
	}
	UntsUntsVrnc struct {
		X UntsUntsAct
	}
	UntsUntsStd struct {
		X UntsUntsAct
	}
	UntsUntsZscr struct {
		X UntsUntsAct
	}
	UntsUntsZscrInplace struct {
		X UntsUntsAct
	}
	UntsUntsRngFul struct {
		X UntsUntsAct
	}
	UntsUntsRngLst struct {
		X UntsUntsAct
	}
	UntsUntsProLst struct {
		X UntsUntsAct
	}
	UntsUntsProSma struct {
		X UntsUntsAct
	}
	IntsIntsCnt struct {
		X IntsIntsAct
	}
	IntsIntsCpy struct {
		X IntsIntsAct
	}
	IntsIntsClr struct {
		X IntsIntsAct
	}
	IntsIntsRand struct {
		X IntsIntsAct
	}
	IntsIntsMrg struct {
		X  IntsIntsAct
		I0 []IntsIntsAct
	}
	IntsIntsPush struct {
		X  IntsIntsAct
		I0 []IntIntAct
	}
	IntsIntsPop struct {
		X IntsIntsAct
	}
	IntsIntsQue struct {
		X  IntsIntsAct
		I0 []IntIntAct
	}
	IntsIntsDque struct {
		X IntsIntsAct
	}
	IntsIntsIns struct {
		X  IntsIntsAct
		I0 UntUntAct
		I1 IntIntAct
	}
	IntsIntsUpd struct {
		X  IntsIntsAct
		I0 UntUntAct
		I1 IntIntAct
	}
	IntsIntsDel struct {
		X  IntsIntsAct
		I0 UntUntAct
	}
	IntsIntsAt struct {
		X  IntsIntsAct
		I0 UntUntAct
	}
	IntsIntsIn struct {
		X  IntsIntsAct
		I0 UntUntAct
		I1 UntUntAct
	}
	IntsIntsInBnd struct {
		X  IntsIntsAct
		I0 BndBndAct
	}
	IntsIntsFrom struct {
		X  IntsIntsAct
		I0 UntUntAct
	}
	IntsIntsTo struct {
		X  IntsIntsAct
		I0 UntUntAct
	}
	IntsIntsFst struct {
		X IntsIntsAct
	}
	IntsIntsMdl struct {
		X IntsIntsAct
	}
	IntsIntsLst struct {
		X IntsIntsAct
	}
	IntsIntsFstIdx struct {
		X IntsIntsAct
	}
	IntsIntsMdlIdx struct {
		X IntsIntsAct
	}
	IntsIntsLstIdx struct {
		X IntsIntsAct
	}
	IntsIntsRev struct {
		X IntsIntsAct
	}
	IntsIntsSrchIdxEql struct {
		X  IntsIntsAct
		I0 IntIntAct
	}
	IntsIntsSrchIdx struct {
		X  IntsIntsAct
		I0 IntIntAct
		I1 []BolBolAct
	}
	IntsIntsHas struct {
		X  IntsIntsAct
		I0 IntIntAct
	}
	IntsIntsSrtAsc struct {
		X IntsIntsAct
	}
	IntsIntsSrtDsc struct {
		X IntsIntsAct
	}
	TmesTmesBnd struct {
		X  TmesTmesAct
		I0 TmeRngAct
	}
	TmesTmesWeekdayCnt struct {
		X TmesTmesAct
	}
	TmesTmesCnt struct {
		X TmesTmesAct
	}
	TmesTmesCpy struct {
		X TmesTmesAct
	}
	TmesTmesClr struct {
		X TmesTmesAct
	}
	TmesTmesRand struct {
		X TmesTmesAct
	}
	TmesTmesMrg struct {
		X  TmesTmesAct
		I0 []TmesTmesAct
	}
	TmesTmesPush struct {
		X  TmesTmesAct
		I0 []TmeTmeAct
	}
	TmesTmesPop struct {
		X TmesTmesAct
	}
	TmesTmesQue struct {
		X  TmesTmesAct
		I0 []TmeTmeAct
	}
	TmesTmesDque struct {
		X TmesTmesAct
	}
	TmesTmesIns struct {
		X  TmesTmesAct
		I0 UntUntAct
		I1 TmeTmeAct
	}
	TmesTmesUpd struct {
		X  TmesTmesAct
		I0 UntUntAct
		I1 TmeTmeAct
	}
	TmesTmesDel struct {
		X  TmesTmesAct
		I0 UntUntAct
	}
	TmesTmesAt struct {
		X  TmesTmesAct
		I0 UntUntAct
	}
	TmesTmesIn struct {
		X  TmesTmesAct
		I0 UntUntAct
		I1 UntUntAct
	}
	TmesTmesInBnd struct {
		X  TmesTmesAct
		I0 BndBndAct
	}
	TmesTmesFrom struct {
		X  TmesTmesAct
		I0 UntUntAct
	}
	TmesTmesTo struct {
		X  TmesTmesAct
		I0 UntUntAct
	}
	TmesTmesFst struct {
		X TmesTmesAct
	}
	TmesTmesMdl struct {
		X TmesTmesAct
	}
	TmesTmesLst struct {
		X TmesTmesAct
	}
	TmesTmesFstIdx struct {
		X TmesTmesAct
	}
	TmesTmesMdlIdx struct {
		X TmesTmesAct
	}
	TmesTmesLstIdx struct {
		X TmesTmesAct
	}
	TmesTmesRev struct {
		X TmesTmesAct
	}
	TmesTmesSrchIdxEql struct {
		X  TmesTmesAct
		I0 TmeTmeAct
	}
	TmesTmesSrchIdx struct {
		X  TmesTmesAct
		I0 TmeTmeAct
		I1 []BolBolAct
	}
	TmesTmesHas struct {
		X  TmesTmesAct
		I0 TmeTmeAct
	}
	TmesTmesSrtAsc struct {
		X TmesTmesAct
	}
	TmesTmesSrtDsc struct {
		X TmesTmesAct
	}
	TmesTmesInrAdd struct {
		X  TmesTmesAct
		I0 UntUntAct
	}
	TmesTmesInrSub struct {
		X  TmesTmesAct
		I0 UntUntAct
	}
	TmesTmesInrMul struct {
		X  TmesTmesAct
		I0 UntUntAct
	}
	TmesTmesInrDiv struct {
		X  TmesTmesAct
		I0 UntUntAct
	}
	TmesTmesInrRem struct {
		X  TmesTmesAct
		I0 UntUntAct
	}
	TmesTmesInrPow struct {
		X  TmesTmesAct
		I0 UntUntAct
	}
	TmesTmesInrMin struct {
		X  TmesTmesAct
		I0 UntUntAct
	}
	TmesTmesInrMax struct {
		X  TmesTmesAct
		I0 UntUntAct
	}
	TmesTmesSum struct {
		X TmesTmesAct
	}
	TmesTmesPrd struct {
		X TmesTmesAct
	}
	TmesTmesMin struct {
		X TmesTmesAct
	}
	TmesTmesMax struct {
		X TmesTmesAct
	}
	TmesTmesMid struct {
		X TmesTmesAct
	}
	TmesTmesMdn struct {
		X TmesTmesAct
	}
	TmesTmesSma struct {
		X TmesTmesAct
	}
	TmesTmesGma struct {
		X TmesTmesAct
	}
	TmesTmesWma struct {
		X TmesTmesAct
	}
	TmesTmesVrnc struct {
		X TmesTmesAct
	}
	TmesTmesStd struct {
		X TmesTmesAct
	}
	TmesTmesZscr struct {
		X TmesTmesAct
	}
	TmesTmesZscrInplace struct {
		X TmesTmesAct
	}
	TmesTmesRngFul struct {
		X TmesTmesAct
	}
	TmesTmesRngLst struct {
		X TmesTmesAct
	}
	TmesTmesProLst struct {
		X TmesTmesAct
	}
	TmesTmesProSma struct {
		X TmesTmesAct
	}
	BndsBndsCnt struct {
		X BndsBndsAct
	}
	BndsBndsCpy struct {
		X BndsBndsAct
	}
	BndsBndsClr struct {
		X BndsBndsAct
	}
	BndsBndsRand struct {
		X BndsBndsAct
	}
	BndsBndsMrg struct {
		X  BndsBndsAct
		I0 []BndsBndsAct
	}
	BndsBndsPush struct {
		X  BndsBndsAct
		I0 []BndBndAct
	}
	BndsBndsPop struct {
		X BndsBndsAct
	}
	BndsBndsQue struct {
		X  BndsBndsAct
		I0 []BndBndAct
	}
	BndsBndsDque struct {
		X BndsBndsAct
	}
	BndsBndsIns struct {
		X  BndsBndsAct
		I0 UntUntAct
		I1 BndBndAct
	}
	BndsBndsUpd struct {
		X  BndsBndsAct
		I0 UntUntAct
		I1 BndBndAct
	}
	BndsBndsDel struct {
		X  BndsBndsAct
		I0 UntUntAct
	}
	BndsBndsAt struct {
		X  BndsBndsAct
		I0 UntUntAct
	}
	BndsBndsIn struct {
		X  BndsBndsAct
		I0 UntUntAct
		I1 UntUntAct
	}
	BndsBndsInBnd struct {
		X  BndsBndsAct
		I0 BndBndAct
	}
	BndsBndsFrom struct {
		X  BndsBndsAct
		I0 UntUntAct
	}
	BndsBndsTo struct {
		X  BndsBndsAct
		I0 UntUntAct
	}
	BndsBndsFst struct {
		X BndsBndsAct
	}
	BndsBndsMdl struct {
		X BndsBndsAct
	}
	BndsBndsLst struct {
		X BndsBndsAct
	}
	BndsBndsFstIdx struct {
		X BndsBndsAct
	}
	BndsBndsMdlIdx struct {
		X BndsBndsAct
	}
	BndsBndsLstIdx struct {
		X BndsBndsAct
	}
	BndsBndsRev struct {
		X BndsBndsAct
	}
	TmeRngsCnt struct {
		X TmeRngsAct
	}
	TmeRngsCpy struct {
		X TmeRngsAct
	}
	TmeRngsClr struct {
		X TmeRngsAct
	}
	TmeRngsRand struct {
		X TmeRngsAct
	}
	TmeRngsMrg struct {
		X  TmeRngsAct
		I0 []TmeRngsAct
	}
	TmeRngsPush struct {
		X  TmeRngsAct
		I0 []TmeRngAct
	}
	TmeRngsPop struct {
		X TmeRngsAct
	}
	TmeRngsQue struct {
		X  TmeRngsAct
		I0 []TmeRngAct
	}
	TmeRngsDque struct {
		X TmeRngsAct
	}
	TmeRngsIns struct {
		X  TmeRngsAct
		I0 UntUntAct
		I1 TmeRngAct
	}
	TmeRngsUpd struct {
		X  TmeRngsAct
		I0 UntUntAct
		I1 TmeRngAct
	}
	TmeRngsDel struct {
		X  TmeRngsAct
		I0 UntUntAct
	}
	TmeRngsAt struct {
		X  TmeRngsAct
		I0 UntUntAct
	}
	TmeRngsIn struct {
		X  TmeRngsAct
		I0 UntUntAct
		I1 UntUntAct
	}
	TmeRngsInBnd struct {
		X  TmeRngsAct
		I0 BndBndAct
	}
	TmeRngsFrom struct {
		X  TmeRngsAct
		I0 UntUntAct
	}
	TmeRngsTo struct {
		X  TmeRngsAct
		I0 UntUntAct
	}
	TmeRngsFst struct {
		X TmeRngsAct
	}
	TmeRngsMdl struct {
		X TmeRngsAct
	}
	TmeRngsLst struct {
		X TmeRngsAct
	}
	TmeRngsFstIdx struct {
		X TmeRngsAct
	}
	TmeRngsMdlIdx struct {
		X TmeRngsAct
	}
	TmeRngsLstIdx struct {
		X TmeRngsAct
	}
	TmeRngsRev struct {
		X TmeRngsAct
	}
	TmeRngsSrchIdx struct {
		X  TmeRngsAct
		I0 TmeTmeAct
	}
	TmeRngsRngMrg struct {
		X  TmeRngsAct
		I0 UntUntAct
		I1 UntUntAct
	}
	AnaTrdOpnMid struct {
		X AnaTrdAct
	}
	AnaTrdClsMid struct {
		X AnaTrdAct
	}
	AnaTrdsCnt struct {
		X AnaTrdsAct
	}
	AnaTrdsCpy struct {
		X AnaTrdsAct
	}
	AnaTrdsClr struct {
		X AnaTrdsAct
	}
	AnaTrdsRand struct {
		X AnaTrdsAct
	}
	AnaTrdsMrg struct {
		X  AnaTrdsAct
		I0 []AnaTrdsAct
	}
	AnaTrdsPush struct {
		X  AnaTrdsAct
		I0 []AnaTrdAct
	}
	AnaTrdsPop struct {
		X AnaTrdsAct
	}
	AnaTrdsQue struct {
		X  AnaTrdsAct
		I0 []AnaTrdAct
	}
	AnaTrdsDque struct {
		X AnaTrdsAct
	}
	AnaTrdsIns struct {
		X  AnaTrdsAct
		I0 UntUntAct
		I1 AnaTrdAct
	}
	AnaTrdsUpd struct {
		X  AnaTrdsAct
		I0 UntUntAct
		I1 AnaTrdAct
	}
	AnaTrdsDel struct {
		X  AnaTrdsAct
		I0 UntUntAct
	}
	AnaTrdsAt struct {
		X  AnaTrdsAct
		I0 UntUntAct
	}
	AnaTrdsIn struct {
		X  AnaTrdsAct
		I0 UntUntAct
		I1 UntUntAct
	}
	AnaTrdsInBnd struct {
		X  AnaTrdsAct
		I0 BndBndAct
	}
	AnaTrdsFrom struct {
		X  AnaTrdsAct
		I0 UntUntAct
	}
	AnaTrdsTo struct {
		X  AnaTrdsAct
		I0 UntUntAct
	}
	AnaTrdsFst struct {
		X AnaTrdsAct
	}
	AnaTrdsMdl struct {
		X AnaTrdsAct
	}
	AnaTrdsLst struct {
		X AnaTrdsAct
	}
	AnaTrdsFstIdx struct {
		X AnaTrdsAct
	}
	AnaTrdsMdlIdx struct {
		X AnaTrdsAct
	}
	AnaTrdsLstIdx struct {
		X AnaTrdsAct
	}
	AnaTrdsRev struct {
		X AnaTrdsAct
	}
	AnaTrdsSelClsResEql struct {
		X  AnaTrdsAct
		I0 StrStrAct
	}
	AnaTrdsSelClsResNeq struct {
		X  AnaTrdsAct
		I0 StrStrAct
	}
	AnaTrdsSelClsResLss struct {
		X  AnaTrdsAct
		I0 StrStrAct
	}
	AnaTrdsSelClsResGtr struct {
		X  AnaTrdsAct
		I0 StrStrAct
	}
	AnaTrdsSelClsResLeq struct {
		X  AnaTrdsAct
		I0 StrStrAct
	}
	AnaTrdsSelClsResGeq struct {
		X  AnaTrdsAct
		I0 StrStrAct
	}
	AnaTrdsSelClsReqEql struct {
		X  AnaTrdsAct
		I0 StrStrAct
	}
	AnaTrdsSelClsReqNeq struct {
		X  AnaTrdsAct
		I0 StrStrAct
	}
	AnaTrdsSelClsReqLss struct {
		X  AnaTrdsAct
		I0 StrStrAct
	}
	AnaTrdsSelClsReqGtr struct {
		X  AnaTrdsAct
		I0 StrStrAct
	}
	AnaTrdsSelClsReqLeq struct {
		X  AnaTrdsAct
		I0 StrStrAct
	}
	AnaTrdsSelClsReqGeq struct {
		X  AnaTrdsAct
		I0 StrStrAct
	}
	AnaTrdsSelOpnResEql struct {
		X  AnaTrdsAct
		I0 StrStrAct
	}
	AnaTrdsSelOpnResNeq struct {
		X  AnaTrdsAct
		I0 StrStrAct
	}
	AnaTrdsSelOpnResLss struct {
		X  AnaTrdsAct
		I0 StrStrAct
	}
	AnaTrdsSelOpnResGtr struct {
		X  AnaTrdsAct
		I0 StrStrAct
	}
	AnaTrdsSelOpnResLeq struct {
		X  AnaTrdsAct
		I0 StrStrAct
	}
	AnaTrdsSelOpnResGeq struct {
		X  AnaTrdsAct
		I0 StrStrAct
	}
	AnaTrdsSelOpnReqEql struct {
		X  AnaTrdsAct
		I0 StrStrAct
	}
	AnaTrdsSelOpnReqNeq struct {
		X  AnaTrdsAct
		I0 StrStrAct
	}
	AnaTrdsSelOpnReqLss struct {
		X  AnaTrdsAct
		I0 StrStrAct
	}
	AnaTrdsSelOpnReqGtr struct {
		X  AnaTrdsAct
		I0 StrStrAct
	}
	AnaTrdsSelOpnReqLeq struct {
		X  AnaTrdsAct
		I0 StrStrAct
	}
	AnaTrdsSelOpnReqGeq struct {
		X  AnaTrdsAct
		I0 StrStrAct
	}
	AnaTrdsSelInstrEql struct {
		X  AnaTrdsAct
		I0 StrStrAct
	}
	AnaTrdsSelInstrNeq struct {
		X  AnaTrdsAct
		I0 StrStrAct
	}
	AnaTrdsSelInstrLss struct {
		X  AnaTrdsAct
		I0 StrStrAct
	}
	AnaTrdsSelInstrGtr struct {
		X  AnaTrdsAct
		I0 StrStrAct
	}
	AnaTrdsSelInstrLeq struct {
		X  AnaTrdsAct
		I0 StrStrAct
	}
	AnaTrdsSelInstrGeq struct {
		X  AnaTrdsAct
		I0 StrStrAct
	}
	AnaTrdsSelUnitsEql struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelUnitsNeq struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelUnitsLss struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelUnitsGtr struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelUnitsLeq struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelUnitsGeq struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelMrgnRtioEql struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelMrgnRtioNeq struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelMrgnRtioLss struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelMrgnRtioGtr struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelMrgnRtioLeq struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelMrgnRtioGeq struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelTrdPctEql struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelTrdPctNeq struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelTrdPctLss struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelTrdPctGtr struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelTrdPctLeq struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelTrdPctGeq struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelClsBalUsdActEql struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelClsBalUsdActNeq struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelClsBalUsdActLss struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelClsBalUsdActGtr struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelClsBalUsdActLeq struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelClsBalUsdActGeq struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelClsBalUsdEql struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelClsBalUsdNeq struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelClsBalUsdLss struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelClsBalUsdGtr struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelClsBalUsdLeq struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelClsBalUsdGeq struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelOpnBalUsdEql struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelOpnBalUsdNeq struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelOpnBalUsdLss struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelOpnBalUsdGtr struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelOpnBalUsdLeq struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelOpnBalUsdGeq struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelCstOpnSpdUsdEql struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelCstOpnSpdUsdNeq struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelCstOpnSpdUsdLss struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelCstOpnSpdUsdGtr struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelCstOpnSpdUsdLeq struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelCstOpnSpdUsdGeq struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelCstClsSpdUsdEql struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelCstClsSpdUsdNeq struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelCstClsSpdUsdLss struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelCstClsSpdUsdGtr struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelCstClsSpdUsdLeq struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelCstClsSpdUsdGeq struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelCstComUsdEql struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelCstComUsdNeq struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelCstComUsdLss struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelCstComUsdGtr struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelCstComUsdLeq struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelCstComUsdGeq struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelPnlGrsUsdEql struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelPnlGrsUsdNeq struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelPnlGrsUsdLss struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelPnlGrsUsdGtr struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelPnlGrsUsdLeq struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelPnlGrsUsdGeq struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelPnlUsdEql struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelPnlUsdNeq struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelPnlUsdLss struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelPnlUsdGtr struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelPnlUsdLeq struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelPnlUsdGeq struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelPnlPctPredictEql struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelPnlPctPredictNeq struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelPnlPctPredictLss struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelPnlPctPredictGtr struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelPnlPctPredictLeq struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelPnlPctPredictGeq struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelPnlPctEql struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelPnlPctNeq struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelPnlPctLss struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelPnlPctGtr struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelPnlPctLeq struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelPnlPctGeq struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelIsLongEql struct {
		X  AnaTrdsAct
		I0 BolBolAct
	}
	AnaTrdsSelIsLongNeq struct {
		X  AnaTrdsAct
		I0 BolBolAct
	}
	AnaTrdsSelDurEql struct {
		X  AnaTrdsAct
		I0 TmeTmeAct
	}
	AnaTrdsSelDurNeq struct {
		X  AnaTrdsAct
		I0 TmeTmeAct
	}
	AnaTrdsSelDurLss struct {
		X  AnaTrdsAct
		I0 TmeTmeAct
	}
	AnaTrdsSelDurGtr struct {
		X  AnaTrdsAct
		I0 TmeTmeAct
	}
	AnaTrdsSelDurLeq struct {
		X  AnaTrdsAct
		I0 TmeTmeAct
	}
	AnaTrdsSelDurGeq struct {
		X  AnaTrdsAct
		I0 TmeTmeAct
	}
	AnaTrdsSelPipEql struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelPipNeq struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelPipLss struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelPipGtr struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelPipLeq struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelPipGeq struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelClsRsnEql struct {
		X  AnaTrdsAct
		I0 StrStrAct
	}
	AnaTrdsSelClsRsnNeq struct {
		X  AnaTrdsAct
		I0 StrStrAct
	}
	AnaTrdsSelClsRsnLss struct {
		X  AnaTrdsAct
		I0 StrStrAct
	}
	AnaTrdsSelClsRsnGtr struct {
		X  AnaTrdsAct
		I0 StrStrAct
	}
	AnaTrdsSelClsRsnLeq struct {
		X  AnaTrdsAct
		I0 StrStrAct
	}
	AnaTrdsSelClsRsnGeq struct {
		X  AnaTrdsAct
		I0 StrStrAct
	}
	AnaTrdsSelClsSpdEql struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelClsSpdNeq struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelClsSpdLss struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelClsSpdGtr struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelClsSpdLeq struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelClsSpdGeq struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelOpnSpdEql struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelOpnSpdNeq struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelOpnSpdLss struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelOpnSpdGtr struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelOpnSpdLeq struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelOpnSpdGeq struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelClsAskEql struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelClsAskNeq struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelClsAskLss struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelClsAskGtr struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelClsAskLeq struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelClsAskGeq struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelOpnAskEql struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelOpnAskNeq struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelOpnAskLss struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelOpnAskGtr struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelOpnAskLeq struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelOpnAskGeq struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelClsBidEql struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelClsBidNeq struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelClsBidLss struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelClsBidGtr struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelClsBidLeq struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelClsBidGeq struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelOpnBidEql struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelOpnBidNeq struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelOpnBidLss struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelOpnBidGtr struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelOpnBidLeq struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelOpnBidGeq struct {
		X  AnaTrdsAct
		I0 FltFltAct
	}
	AnaTrdsSelClsTmeEql struct {
		X  AnaTrdsAct
		I0 TmeTmeAct
	}
	AnaTrdsSelClsTmeNeq struct {
		X  AnaTrdsAct
		I0 TmeTmeAct
	}
	AnaTrdsSelClsTmeLss struct {
		X  AnaTrdsAct
		I0 TmeTmeAct
	}
	AnaTrdsSelClsTmeGtr struct {
		X  AnaTrdsAct
		I0 TmeTmeAct
	}
	AnaTrdsSelClsTmeLeq struct {
		X  AnaTrdsAct
		I0 TmeTmeAct
	}
	AnaTrdsSelClsTmeGeq struct {
		X  AnaTrdsAct
		I0 TmeTmeAct
	}
	AnaTrdsSelOpnTmeEql struct {
		X  AnaTrdsAct
		I0 TmeTmeAct
	}
	AnaTrdsSelOpnTmeNeq struct {
		X  AnaTrdsAct
		I0 TmeTmeAct
	}
	AnaTrdsSelOpnTmeLss struct {
		X  AnaTrdsAct
		I0 TmeTmeAct
	}
	AnaTrdsSelOpnTmeGtr struct {
		X  AnaTrdsAct
		I0 TmeTmeAct
	}
	AnaTrdsSelOpnTmeLeq struct {
		X  AnaTrdsAct
		I0 TmeTmeAct
	}
	AnaTrdsSelOpnTmeGeq struct {
		X  AnaTrdsAct
		I0 TmeTmeAct
	}
	AnaTrdsOpnTmes struct {
		X AnaTrdsAct
	}
	AnaTrdsClsTmes struct {
		X AnaTrdsAct
	}
	AnaTrdsOpnBids struct {
		X AnaTrdsAct
	}
	AnaTrdsClsBids struct {
		X AnaTrdsAct
	}
	AnaTrdsOpnAsks struct {
		X AnaTrdsAct
	}
	AnaTrdsClsAsks struct {
		X AnaTrdsAct
	}
	AnaTrdsOpnSpds struct {
		X AnaTrdsAct
	}
	AnaTrdsClsSpds struct {
		X AnaTrdsAct
	}
	AnaTrdsClsRsns struct {
		X AnaTrdsAct
	}
	AnaTrdsPips struct {
		X AnaTrdsAct
	}
	AnaTrdsDurs struct {
		X AnaTrdsAct
	}
	AnaTrdsIsLongs struct {
		X AnaTrdsAct
	}
	AnaTrdsPnlPcts struct {
		X AnaTrdsAct
	}
	AnaTrdsPnlPctPredicts struct {
		X AnaTrdsAct
	}
	AnaTrdsPnlUsds struct {
		X AnaTrdsAct
	}
	AnaTrdsPnlGrsUsds struct {
		X AnaTrdsAct
	}
	AnaTrdsCstComUsds struct {
		X AnaTrdsAct
	}
	AnaTrdsCstClsSpdUsds struct {
		X AnaTrdsAct
	}
	AnaTrdsCstOpnSpdUsds struct {
		X AnaTrdsAct
	}
	AnaTrdsOpnBalUsds struct {
		X AnaTrdsAct
	}
	AnaTrdsClsBalUsds struct {
		X AnaTrdsAct
	}
	AnaTrdsClsBalUsdActs struct {
		X AnaTrdsAct
	}
	AnaTrdsTrdPcts struct {
		X AnaTrdsAct
	}
	AnaTrdsMrgnRtios struct {
		X AnaTrdsAct
	}
	AnaTrdsUnitss struct {
		X AnaTrdsAct
	}
	AnaTrdsInstrs struct {
		X AnaTrdsAct
	}
	AnaTrdsOpnReqs struct {
		X AnaTrdsAct
	}
	AnaTrdsOpnRess struct {
		X AnaTrdsAct
	}
	AnaTrdsClsReqs struct {
		X AnaTrdsAct
	}
	AnaTrdsClsRess struct {
		X AnaTrdsAct
	}
	AnaPrfmDlt struct {
		X  AnaPrfmAct
		I0 AnaPrfmAct
	}
	AnaPrfmsCnt struct {
		X AnaPrfmsAct
	}
	AnaPrfmsCpy struct {
		X AnaPrfmsAct
	}
	AnaPrfmsClr struct {
		X AnaPrfmsAct
	}
	AnaPrfmsRand struct {
		X AnaPrfmsAct
	}
	AnaPrfmsMrg struct {
		X  AnaPrfmsAct
		I0 []AnaPrfmsAct
	}
	AnaPrfmsPush struct {
		X  AnaPrfmsAct
		I0 []AnaPrfmAct
	}
	AnaPrfmsPop struct {
		X AnaPrfmsAct
	}
	AnaPrfmsQue struct {
		X  AnaPrfmsAct
		I0 []AnaPrfmAct
	}
	AnaPrfmsDque struct {
		X AnaPrfmsAct
	}
	AnaPrfmsIns struct {
		X  AnaPrfmsAct
		I0 UntUntAct
		I1 AnaPrfmAct
	}
	AnaPrfmsUpd struct {
		X  AnaPrfmsAct
		I0 UntUntAct
		I1 AnaPrfmAct
	}
	AnaPrfmsDel struct {
		X  AnaPrfmsAct
		I0 UntUntAct
	}
	AnaPrfmsAt struct {
		X  AnaPrfmsAct
		I0 UntUntAct
	}
	AnaPrfmsIn struct {
		X  AnaPrfmsAct
		I0 UntUntAct
		I1 UntUntAct
	}
	AnaPrfmsInBnd struct {
		X  AnaPrfmsAct
		I0 BndBndAct
	}
	AnaPrfmsFrom struct {
		X  AnaPrfmsAct
		I0 UntUntAct
	}
	AnaPrfmsTo struct {
		X  AnaPrfmsAct
		I0 UntUntAct
	}
	AnaPrfmsFst struct {
		X AnaPrfmsAct
	}
	AnaPrfmsMdl struct {
		X AnaPrfmsAct
	}
	AnaPrfmsLst struct {
		X AnaPrfmsAct
	}
	AnaPrfmsFstIdx struct {
		X AnaPrfmsAct
	}
	AnaPrfmsMdlIdx struct {
		X AnaPrfmsAct
	}
	AnaPrfmsLstIdx struct {
		X AnaPrfmsAct
	}
	AnaPrfmsRev struct {
		X AnaPrfmsAct
	}
	AnaPrfmsPnlPcts struct {
		X AnaPrfmsAct
	}
	AnaPrfmsScsPcts struct {
		X AnaPrfmsAct
	}
	AnaPrfmsPipPerDays struct {
		X AnaPrfmsAct
	}
	AnaPrfmsUsdPerDays struct {
		X AnaPrfmsAct
	}
	AnaPrfmsScsPerDays struct {
		X AnaPrfmsAct
	}
	AnaPrfmsOpnPerDays struct {
		X AnaPrfmsAct
	}
	AnaPrfmsPnlUsds struct {
		X AnaPrfmsAct
	}
	AnaPrfmsPipAvgs struct {
		X AnaPrfmsAct
	}
	AnaPrfmsPipMdns struct {
		X AnaPrfmsAct
	}
	AnaPrfmsPipMins struct {
		X AnaPrfmsAct
	}
	AnaPrfmsPipMaxs struct {
		X AnaPrfmsAct
	}
	AnaPrfmsPipSums struct {
		X AnaPrfmsAct
	}
	AnaPrfmsDurAvgs struct {
		X AnaPrfmsAct
	}
	AnaPrfmsDurMdns struct {
		X AnaPrfmsAct
	}
	AnaPrfmsDurMins struct {
		X AnaPrfmsAct
	}
	AnaPrfmsDurMaxs struct {
		X AnaPrfmsAct
	}
	AnaPrfmsLosLimMaxs struct {
		X AnaPrfmsAct
	}
	AnaPrfmsDurLimMaxs struct {
		X AnaPrfmsAct
	}
	AnaPrfmsDayCnts struct {
		X AnaPrfmsAct
	}
	AnaPrfmsTrdCnts struct {
		X AnaPrfmsAct
	}
	AnaPrfmsTrdPcts struct {
		X AnaPrfmsAct
	}
	AnaPrfmsCstTotUsds struct {
		X AnaPrfmsAct
	}
	AnaPrfmsCstSpdUsds struct {
		X AnaPrfmsAct
	}
	AnaPrfmsCstComUsds struct {
		X AnaPrfmsAct
	}
	AnaPrfmsPths struct {
		X AnaPrfmsAct
	}
	HstPrvsCnt struct {
		X HstPrvsAct
	}
	HstPrvsCpy struct {
		X HstPrvsAct
	}
	HstPrvsClr struct {
		X HstPrvsAct
	}
	HstPrvsRand struct {
		X HstPrvsAct
	}
	HstPrvsMrg struct {
		X  HstPrvsAct
		I0 []HstPrvsAct
	}
	HstPrvsPush struct {
		X  HstPrvsAct
		I0 []HstPrvAct
	}
	HstPrvsPop struct {
		X HstPrvsAct
	}
	HstPrvsQue struct {
		X  HstPrvsAct
		I0 []HstPrvAct
	}
	HstPrvsDque struct {
		X HstPrvsAct
	}
	HstPrvsIns struct {
		X  HstPrvsAct
		I0 UntUntAct
		I1 HstPrvAct
	}
	HstPrvsUpd struct {
		X  HstPrvsAct
		I0 UntUntAct
		I1 HstPrvAct
	}
	HstPrvsDel struct {
		X  HstPrvsAct
		I0 UntUntAct
	}
	HstPrvsAt struct {
		X  HstPrvsAct
		I0 UntUntAct
	}
	HstPrvsIn struct {
		X  HstPrvsAct
		I0 UntUntAct
		I1 UntUntAct
	}
	HstPrvsInBnd struct {
		X  HstPrvsAct
		I0 BndBndAct
	}
	HstPrvsFrom struct {
		X  HstPrvsAct
		I0 UntUntAct
	}
	HstPrvsTo struct {
		X  HstPrvsAct
		I0 UntUntAct
	}
	HstPrvsFst struct {
		X HstPrvsAct
	}
	HstPrvsMdl struct {
		X HstPrvsAct
	}
	HstPrvsLst struct {
		X HstPrvsAct
	}
	HstPrvsFstIdx struct {
		X HstPrvsAct
	}
	HstPrvsMdlIdx struct {
		X HstPrvsAct
	}
	HstPrvsLstIdx struct {
		X HstPrvsAct
	}
	HstPrvsRev struct {
		X HstPrvsAct
	}
	HstInstrsCnt struct {
		X HstInstrsAct
	}
	HstInstrsCpy struct {
		X HstInstrsAct
	}
	HstInstrsClr struct {
		X HstInstrsAct
	}
	HstInstrsRand struct {
		X HstInstrsAct
	}
	HstInstrsMrg struct {
		X  HstInstrsAct
		I0 []HstInstrsAct
	}
	HstInstrsPush struct {
		X  HstInstrsAct
		I0 []HstInstrAct
	}
	HstInstrsPop struct {
		X HstInstrsAct
	}
	HstInstrsQue struct {
		X  HstInstrsAct
		I0 []HstInstrAct
	}
	HstInstrsDque struct {
		X HstInstrsAct
	}
	HstInstrsIns struct {
		X  HstInstrsAct
		I0 UntUntAct
		I1 HstInstrAct
	}
	HstInstrsUpd struct {
		X  HstInstrsAct
		I0 UntUntAct
		I1 HstInstrAct
	}
	HstInstrsDel struct {
		X  HstInstrsAct
		I0 UntUntAct
	}
	HstInstrsAt struct {
		X  HstInstrsAct
		I0 UntUntAct
	}
	HstInstrsIn struct {
		X  HstInstrsAct
		I0 UntUntAct
		I1 UntUntAct
	}
	HstInstrsInBnd struct {
		X  HstInstrsAct
		I0 BndBndAct
	}
	HstInstrsFrom struct {
		X  HstInstrsAct
		I0 UntUntAct
	}
	HstInstrsTo struct {
		X  HstInstrsAct
		I0 UntUntAct
	}
	HstInstrsFst struct {
		X HstInstrsAct
	}
	HstInstrsMdl struct {
		X HstInstrsAct
	}
	HstInstrsLst struct {
		X HstInstrsAct
	}
	HstInstrsFstIdx struct {
		X HstInstrsAct
	}
	HstInstrsMdlIdx struct {
		X HstInstrsAct
	}
	HstInstrsLstIdx struct {
		X HstInstrsAct
	}
	HstInstrsRev struct {
		X HstInstrsAct
	}
	HstInrvlsCnt struct {
		X HstInrvlsAct
	}
	HstInrvlsCpy struct {
		X HstInrvlsAct
	}
	HstInrvlsClr struct {
		X HstInrvlsAct
	}
	HstInrvlsRand struct {
		X HstInrvlsAct
	}
	HstInrvlsMrg struct {
		X  HstInrvlsAct
		I0 []HstInrvlsAct
	}
	HstInrvlsPush struct {
		X  HstInrvlsAct
		I0 []HstInrvlAct
	}
	HstInrvlsPop struct {
		X HstInrvlsAct
	}
	HstInrvlsQue struct {
		X  HstInrvlsAct
		I0 []HstInrvlAct
	}
	HstInrvlsDque struct {
		X HstInrvlsAct
	}
	HstInrvlsIns struct {
		X  HstInrvlsAct
		I0 UntUntAct
		I1 HstInrvlAct
	}
	HstInrvlsUpd struct {
		X  HstInrvlsAct
		I0 UntUntAct
		I1 HstInrvlAct
	}
	HstInrvlsDel struct {
		X  HstInrvlsAct
		I0 UntUntAct
	}
	HstInrvlsAt struct {
		X  HstInrvlsAct
		I0 UntUntAct
	}
	HstInrvlsIn struct {
		X  HstInrvlsAct
		I0 UntUntAct
		I1 UntUntAct
	}
	HstInrvlsInBnd struct {
		X  HstInrvlsAct
		I0 BndBndAct
	}
	HstInrvlsFrom struct {
		X  HstInrvlsAct
		I0 UntUntAct
	}
	HstInrvlsTo struct {
		X  HstInrvlsAct
		I0 UntUntAct
	}
	HstInrvlsFst struct {
		X HstInrvlsAct
	}
	HstInrvlsMdl struct {
		X HstInrvlsAct
	}
	HstInrvlsLst struct {
		X HstInrvlsAct
	}
	HstInrvlsFstIdx struct {
		X HstInrvlsAct
	}
	HstInrvlsMdlIdx struct {
		X HstInrvlsAct
	}
	HstInrvlsLstIdx struct {
		X HstInrvlsAct
	}
	HstInrvlsRev struct {
		X HstInrvlsAct
	}
	HstSidesCnt struct {
		X HstSidesAct
	}
	HstSidesCpy struct {
		X HstSidesAct
	}
	HstSidesClr struct {
		X HstSidesAct
	}
	HstSidesRand struct {
		X HstSidesAct
	}
	HstSidesMrg struct {
		X  HstSidesAct
		I0 []HstSidesAct
	}
	HstSidesPush struct {
		X  HstSidesAct
		I0 []HstSideAct
	}
	HstSidesPop struct {
		X HstSidesAct
	}
	HstSidesQue struct {
		X  HstSidesAct
		I0 []HstSideAct
	}
	HstSidesDque struct {
		X HstSidesAct
	}
	HstSidesIns struct {
		X  HstSidesAct
		I0 UntUntAct
		I1 HstSideAct
	}
	HstSidesUpd struct {
		X  HstSidesAct
		I0 UntUntAct
		I1 HstSideAct
	}
	HstSidesDel struct {
		X  HstSidesAct
		I0 UntUntAct
	}
	HstSidesAt struct {
		X  HstSidesAct
		I0 UntUntAct
	}
	HstSidesIn struct {
		X  HstSidesAct
		I0 UntUntAct
		I1 UntUntAct
	}
	HstSidesInBnd struct {
		X  HstSidesAct
		I0 BndBndAct
	}
	HstSidesFrom struct {
		X  HstSidesAct
		I0 UntUntAct
	}
	HstSidesTo struct {
		X  HstSidesAct
		I0 UntUntAct
	}
	HstSidesFst struct {
		X HstSidesAct
	}
	HstSidesMdl struct {
		X HstSidesAct
	}
	HstSidesLst struct {
		X HstSidesAct
	}
	HstSidesFstIdx struct {
		X HstSidesAct
	}
	HstSidesMdlIdx struct {
		X HstSidesAct
	}
	HstSidesLstIdx struct {
		X HstSidesAct
	}
	HstSidesRev struct {
		X HstSidesAct
	}
	HstStmsCnt struct {
		X HstStmsAct
	}
	HstStmsCpy struct {
		X HstStmsAct
	}
	HstStmsClr struct {
		X HstStmsAct
	}
	HstStmsRand struct {
		X HstStmsAct
	}
	HstStmsMrg struct {
		X  HstStmsAct
		I0 []HstStmsAct
	}
	HstStmsPush struct {
		X  HstStmsAct
		I0 []HstStmAct
	}
	HstStmsPop struct {
		X HstStmsAct
	}
	HstStmsQue struct {
		X  HstStmsAct
		I0 []HstStmAct
	}
	HstStmsDque struct {
		X HstStmsAct
	}
	HstStmsIns struct {
		X  HstStmsAct
		I0 UntUntAct
		I1 HstStmAct
	}
	HstStmsUpd struct {
		X  HstStmsAct
		I0 UntUntAct
		I1 HstStmAct
	}
	HstStmsDel struct {
		X  HstStmsAct
		I0 UntUntAct
	}
	HstStmsAt struct {
		X  HstStmsAct
		I0 UntUntAct
	}
	HstStmsIn struct {
		X  HstStmsAct
		I0 UntUntAct
		I1 UntUntAct
	}
	HstStmsInBnd struct {
		X  HstStmsAct
		I0 BndBndAct
	}
	HstStmsFrom struct {
		X  HstStmsAct
		I0 UntUntAct
	}
	HstStmsTo struct {
		X  HstStmsAct
		I0 UntUntAct
	}
	HstStmsFst struct {
		X HstStmsAct
	}
	HstStmsMdl struct {
		X HstStmsAct
	}
	HstStmsLst struct {
		X HstStmsAct
	}
	HstStmsFstIdx struct {
		X HstStmsAct
	}
	HstStmsMdlIdx struct {
		X HstStmsAct
	}
	HstStmsLstIdx struct {
		X HstStmsAct
	}
	HstStmsRev struct {
		X HstStmsAct
	}
	HstCndsCnt struct {
		X HstCndsAct
	}
	HstCndsCpy struct {
		X HstCndsAct
	}
	HstCndsClr struct {
		X HstCndsAct
	}
	HstCndsRand struct {
		X HstCndsAct
	}
	HstCndsMrg struct {
		X  HstCndsAct
		I0 []HstCndsAct
	}
	HstCndsPush struct {
		X  HstCndsAct
		I0 []HstCndAct
	}
	HstCndsPop struct {
		X HstCndsAct
	}
	HstCndsQue struct {
		X  HstCndsAct
		I0 []HstCndAct
	}
	HstCndsDque struct {
		X HstCndsAct
	}
	HstCndsIns struct {
		X  HstCndsAct
		I0 UntUntAct
		I1 HstCndAct
	}
	HstCndsUpd struct {
		X  HstCndsAct
		I0 UntUntAct
		I1 HstCndAct
	}
	HstCndsDel struct {
		X  HstCndsAct
		I0 UntUntAct
	}
	HstCndsAt struct {
		X  HstCndsAct
		I0 UntUntAct
	}
	HstCndsIn struct {
		X  HstCndsAct
		I0 UntUntAct
		I1 UntUntAct
	}
	HstCndsInBnd struct {
		X  HstCndsAct
		I0 BndBndAct
	}
	HstCndsFrom struct {
		X  HstCndsAct
		I0 UntUntAct
	}
	HstCndsTo struct {
		X  HstCndsAct
		I0 UntUntAct
	}
	HstCndsFst struct {
		X HstCndsAct
	}
	HstCndsMdl struct {
		X HstCndsAct
	}
	HstCndsLst struct {
		X HstCndsAct
	}
	HstCndsFstIdx struct {
		X HstCndsAct
	}
	HstCndsMdlIdx struct {
		X HstCndsAct
	}
	HstCndsLstIdx struct {
		X HstCndsAct
	}
	HstCndsRev struct {
		X HstCndsAct
	}
	HstStgysCnt struct {
		X HstStgysAct
	}
	HstStgysCpy struct {
		X HstStgysAct
	}
	HstStgysClr struct {
		X HstStgysAct
	}
	HstStgysRand struct {
		X HstStgysAct
	}
	HstStgysMrg struct {
		X  HstStgysAct
		I0 []HstStgysAct
	}
	HstStgysPush struct {
		X  HstStgysAct
		I0 []HstStgyAct
	}
	HstStgysPop struct {
		X HstStgysAct
	}
	HstStgysQue struct {
		X  HstStgysAct
		I0 []HstStgyAct
	}
	HstStgysDque struct {
		X HstStgysAct
	}
	HstStgysIns struct {
		X  HstStgysAct
		I0 UntUntAct
		I1 HstStgyAct
	}
	HstStgysUpd struct {
		X  HstStgysAct
		I0 UntUntAct
		I1 HstStgyAct
	}
	HstStgysDel struct {
		X  HstStgysAct
		I0 UntUntAct
	}
	HstStgysAt struct {
		X  HstStgysAct
		I0 UntUntAct
	}
	HstStgysIn struct {
		X  HstStgysAct
		I0 UntUntAct
		I1 UntUntAct
	}
	HstStgysInBnd struct {
		X  HstStgysAct
		I0 BndBndAct
	}
	HstStgysFrom struct {
		X  HstStgysAct
		I0 UntUntAct
	}
	HstStgysTo struct {
		X  HstStgysAct
		I0 UntUntAct
	}
	HstStgysFst struct {
		X HstStgysAct
	}
	HstStgysMdl struct {
		X HstStgysAct
	}
	HstStgysLst struct {
		X HstStgysAct
	}
	HstStgysFstIdx struct {
		X HstStgysAct
	}
	HstStgysMdlIdx struct {
		X HstStgysAct
	}
	HstStgysLstIdx struct {
		X HstStgysAct
	}
	HstStgysRev struct {
		X HstStgysAct
	}
	RltPrvsCnt struct {
		X RltPrvsAct
	}
	RltPrvsCpy struct {
		X RltPrvsAct
	}
	RltPrvsClr struct {
		X RltPrvsAct
	}
	RltPrvsRand struct {
		X RltPrvsAct
	}
	RltPrvsMrg struct {
		X  RltPrvsAct
		I0 []RltPrvsAct
	}
	RltPrvsPush struct {
		X  RltPrvsAct
		I0 []RltPrvAct
	}
	RltPrvsPop struct {
		X RltPrvsAct
	}
	RltPrvsQue struct {
		X  RltPrvsAct
		I0 []RltPrvAct
	}
	RltPrvsDque struct {
		X RltPrvsAct
	}
	RltPrvsIns struct {
		X  RltPrvsAct
		I0 UntUntAct
		I1 RltPrvAct
	}
	RltPrvsUpd struct {
		X  RltPrvsAct
		I0 UntUntAct
		I1 RltPrvAct
	}
	RltPrvsDel struct {
		X  RltPrvsAct
		I0 UntUntAct
	}
	RltPrvsAt struct {
		X  RltPrvsAct
		I0 UntUntAct
	}
	RltPrvsIn struct {
		X  RltPrvsAct
		I0 UntUntAct
		I1 UntUntAct
	}
	RltPrvsInBnd struct {
		X  RltPrvsAct
		I0 BndBndAct
	}
	RltPrvsFrom struct {
		X  RltPrvsAct
		I0 UntUntAct
	}
	RltPrvsTo struct {
		X  RltPrvsAct
		I0 UntUntAct
	}
	RltPrvsFst struct {
		X RltPrvsAct
	}
	RltPrvsMdl struct {
		X RltPrvsAct
	}
	RltPrvsLst struct {
		X RltPrvsAct
	}
	RltPrvsFstIdx struct {
		X RltPrvsAct
	}
	RltPrvsMdlIdx struct {
		X RltPrvsAct
	}
	RltPrvsLstIdx struct {
		X RltPrvsAct
	}
	RltPrvsRev struct {
		X RltPrvsAct
	}
	RltInstrsCnt struct {
		X RltInstrsAct
	}
	RltInstrsCpy struct {
		X RltInstrsAct
	}
	RltInstrsClr struct {
		X RltInstrsAct
	}
	RltInstrsRand struct {
		X RltInstrsAct
	}
	RltInstrsMrg struct {
		X  RltInstrsAct
		I0 []RltInstrsAct
	}
	RltInstrsPush struct {
		X  RltInstrsAct
		I0 []RltInstrAct
	}
	RltInstrsPop struct {
		X RltInstrsAct
	}
	RltInstrsQue struct {
		X  RltInstrsAct
		I0 []RltInstrAct
	}
	RltInstrsDque struct {
		X RltInstrsAct
	}
	RltInstrsIns struct {
		X  RltInstrsAct
		I0 UntUntAct
		I1 RltInstrAct
	}
	RltInstrsUpd struct {
		X  RltInstrsAct
		I0 UntUntAct
		I1 RltInstrAct
	}
	RltInstrsDel struct {
		X  RltInstrsAct
		I0 UntUntAct
	}
	RltInstrsAt struct {
		X  RltInstrsAct
		I0 UntUntAct
	}
	RltInstrsIn struct {
		X  RltInstrsAct
		I0 UntUntAct
		I1 UntUntAct
	}
	RltInstrsInBnd struct {
		X  RltInstrsAct
		I0 BndBndAct
	}
	RltInstrsFrom struct {
		X  RltInstrsAct
		I0 UntUntAct
	}
	RltInstrsTo struct {
		X  RltInstrsAct
		I0 UntUntAct
	}
	RltInstrsFst struct {
		X RltInstrsAct
	}
	RltInstrsMdl struct {
		X RltInstrsAct
	}
	RltInstrsLst struct {
		X RltInstrsAct
	}
	RltInstrsFstIdx struct {
		X RltInstrsAct
	}
	RltInstrsMdlIdx struct {
		X RltInstrsAct
	}
	RltInstrsLstIdx struct {
		X RltInstrsAct
	}
	RltInstrsRev struct {
		X RltInstrsAct
	}
	RltInrvlsCnt struct {
		X RltInrvlsAct
	}
	RltInrvlsCpy struct {
		X RltInrvlsAct
	}
	RltInrvlsClr struct {
		X RltInrvlsAct
	}
	RltInrvlsRand struct {
		X RltInrvlsAct
	}
	RltInrvlsMrg struct {
		X  RltInrvlsAct
		I0 []RltInrvlsAct
	}
	RltInrvlsPush struct {
		X  RltInrvlsAct
		I0 []RltInrvlAct
	}
	RltInrvlsPop struct {
		X RltInrvlsAct
	}
	RltInrvlsQue struct {
		X  RltInrvlsAct
		I0 []RltInrvlAct
	}
	RltInrvlsDque struct {
		X RltInrvlsAct
	}
	RltInrvlsIns struct {
		X  RltInrvlsAct
		I0 UntUntAct
		I1 RltInrvlAct
	}
	RltInrvlsUpd struct {
		X  RltInrvlsAct
		I0 UntUntAct
		I1 RltInrvlAct
	}
	RltInrvlsDel struct {
		X  RltInrvlsAct
		I0 UntUntAct
	}
	RltInrvlsAt struct {
		X  RltInrvlsAct
		I0 UntUntAct
	}
	RltInrvlsIn struct {
		X  RltInrvlsAct
		I0 UntUntAct
		I1 UntUntAct
	}
	RltInrvlsInBnd struct {
		X  RltInrvlsAct
		I0 BndBndAct
	}
	RltInrvlsFrom struct {
		X  RltInrvlsAct
		I0 UntUntAct
	}
	RltInrvlsTo struct {
		X  RltInrvlsAct
		I0 UntUntAct
	}
	RltInrvlsFst struct {
		X RltInrvlsAct
	}
	RltInrvlsMdl struct {
		X RltInrvlsAct
	}
	RltInrvlsLst struct {
		X RltInrvlsAct
	}
	RltInrvlsFstIdx struct {
		X RltInrvlsAct
	}
	RltInrvlsMdlIdx struct {
		X RltInrvlsAct
	}
	RltInrvlsLstIdx struct {
		X RltInrvlsAct
	}
	RltInrvlsRev struct {
		X RltInrvlsAct
	}
	RltSidesCnt struct {
		X RltSidesAct
	}
	RltSidesCpy struct {
		X RltSidesAct
	}
	RltSidesClr struct {
		X RltSidesAct
	}
	RltSidesRand struct {
		X RltSidesAct
	}
	RltSidesMrg struct {
		X  RltSidesAct
		I0 []RltSidesAct
	}
	RltSidesPush struct {
		X  RltSidesAct
		I0 []RltSideAct
	}
	RltSidesPop struct {
		X RltSidesAct
	}
	RltSidesQue struct {
		X  RltSidesAct
		I0 []RltSideAct
	}
	RltSidesDque struct {
		X RltSidesAct
	}
	RltSidesIns struct {
		X  RltSidesAct
		I0 UntUntAct
		I1 RltSideAct
	}
	RltSidesUpd struct {
		X  RltSidesAct
		I0 UntUntAct
		I1 RltSideAct
	}
	RltSidesDel struct {
		X  RltSidesAct
		I0 UntUntAct
	}
	RltSidesAt struct {
		X  RltSidesAct
		I0 UntUntAct
	}
	RltSidesIn struct {
		X  RltSidesAct
		I0 UntUntAct
		I1 UntUntAct
	}
	RltSidesInBnd struct {
		X  RltSidesAct
		I0 BndBndAct
	}
	RltSidesFrom struct {
		X  RltSidesAct
		I0 UntUntAct
	}
	RltSidesTo struct {
		X  RltSidesAct
		I0 UntUntAct
	}
	RltSidesFst struct {
		X RltSidesAct
	}
	RltSidesMdl struct {
		X RltSidesAct
	}
	RltSidesLst struct {
		X RltSidesAct
	}
	RltSidesFstIdx struct {
		X RltSidesAct
	}
	RltSidesMdlIdx struct {
		X RltSidesAct
	}
	RltSidesLstIdx struct {
		X RltSidesAct
	}
	RltSidesRev struct {
		X RltSidesAct
	}
	RltStmsCnt struct {
		X RltStmsAct
	}
	RltStmsCpy struct {
		X RltStmsAct
	}
	RltStmsClr struct {
		X RltStmsAct
	}
	RltStmsRand struct {
		X RltStmsAct
	}
	RltStmsMrg struct {
		X  RltStmsAct
		I0 []RltStmsAct
	}
	RltStmsPush struct {
		X  RltStmsAct
		I0 []RltStmAct
	}
	RltStmsPop struct {
		X RltStmsAct
	}
	RltStmsQue struct {
		X  RltStmsAct
		I0 []RltStmAct
	}
	RltStmsDque struct {
		X RltStmsAct
	}
	RltStmsIns struct {
		X  RltStmsAct
		I0 UntUntAct
		I1 RltStmAct
	}
	RltStmsUpd struct {
		X  RltStmsAct
		I0 UntUntAct
		I1 RltStmAct
	}
	RltStmsDel struct {
		X  RltStmsAct
		I0 UntUntAct
	}
	RltStmsAt struct {
		X  RltStmsAct
		I0 UntUntAct
	}
	RltStmsIn struct {
		X  RltStmsAct
		I0 UntUntAct
		I1 UntUntAct
	}
	RltStmsInBnd struct {
		X  RltStmsAct
		I0 BndBndAct
	}
	RltStmsFrom struct {
		X  RltStmsAct
		I0 UntUntAct
	}
	RltStmsTo struct {
		X  RltStmsAct
		I0 UntUntAct
	}
	RltStmsFst struct {
		X RltStmsAct
	}
	RltStmsMdl struct {
		X RltStmsAct
	}
	RltStmsLst struct {
		X RltStmsAct
	}
	RltStmsFstIdx struct {
		X RltStmsAct
	}
	RltStmsMdlIdx struct {
		X RltStmsAct
	}
	RltStmsLstIdx struct {
		X RltStmsAct
	}
	RltStmsRev struct {
		X RltStmsAct
	}
	RltCndsCnt struct {
		X RltCndsAct
	}
	RltCndsCpy struct {
		X RltCndsAct
	}
	RltCndsClr struct {
		X RltCndsAct
	}
	RltCndsRand struct {
		X RltCndsAct
	}
	RltCndsMrg struct {
		X  RltCndsAct
		I0 []RltCndsAct
	}
	RltCndsPush struct {
		X  RltCndsAct
		I0 []RltCndAct
	}
	RltCndsPop struct {
		X RltCndsAct
	}
	RltCndsQue struct {
		X  RltCndsAct
		I0 []RltCndAct
	}
	RltCndsDque struct {
		X RltCndsAct
	}
	RltCndsIns struct {
		X  RltCndsAct
		I0 UntUntAct
		I1 RltCndAct
	}
	RltCndsUpd struct {
		X  RltCndsAct
		I0 UntUntAct
		I1 RltCndAct
	}
	RltCndsDel struct {
		X  RltCndsAct
		I0 UntUntAct
	}
	RltCndsAt struct {
		X  RltCndsAct
		I0 UntUntAct
	}
	RltCndsIn struct {
		X  RltCndsAct
		I0 UntUntAct
		I1 UntUntAct
	}
	RltCndsInBnd struct {
		X  RltCndsAct
		I0 BndBndAct
	}
	RltCndsFrom struct {
		X  RltCndsAct
		I0 UntUntAct
	}
	RltCndsTo struct {
		X  RltCndsAct
		I0 UntUntAct
	}
	RltCndsFst struct {
		X RltCndsAct
	}
	RltCndsMdl struct {
		X RltCndsAct
	}
	RltCndsLst struct {
		X RltCndsAct
	}
	RltCndsFstIdx struct {
		X RltCndsAct
	}
	RltCndsMdlIdx struct {
		X RltCndsAct
	}
	RltCndsLstIdx struct {
		X RltCndsAct
	}
	RltCndsRev struct {
		X RltCndsAct
	}
	RltStgysCnt struct {
		X RltStgysAct
	}
	RltStgysCpy struct {
		X RltStgysAct
	}
	RltStgysClr struct {
		X RltStgysAct
	}
	RltStgysRand struct {
		X RltStgysAct
	}
	RltStgysMrg struct {
		X  RltStgysAct
		I0 []RltStgysAct
	}
	RltStgysPush struct {
		X  RltStgysAct
		I0 []RltStgyAct
	}
	RltStgysPop struct {
		X RltStgysAct
	}
	RltStgysQue struct {
		X  RltStgysAct
		I0 []RltStgyAct
	}
	RltStgysDque struct {
		X RltStgysAct
	}
	RltStgysIns struct {
		X  RltStgysAct
		I0 UntUntAct
		I1 RltStgyAct
	}
	RltStgysUpd struct {
		X  RltStgysAct
		I0 UntUntAct
		I1 RltStgyAct
	}
	RltStgysDel struct {
		X  RltStgysAct
		I0 UntUntAct
	}
	RltStgysAt struct {
		X  RltStgysAct
		I0 UntUntAct
	}
	RltStgysIn struct {
		X  RltStgysAct
		I0 UntUntAct
		I1 UntUntAct
	}
	RltStgysInBnd struct {
		X  RltStgysAct
		I0 BndBndAct
	}
	RltStgysFrom struct {
		X  RltStgysAct
		I0 UntUntAct
	}
	RltStgysTo struct {
		X  RltStgysAct
		I0 UntUntAct
	}
	RltStgysFst struct {
		X RltStgysAct
	}
	RltStgysMdl struct {
		X RltStgysAct
	}
	RltStgysLst struct {
		X RltStgysAct
	}
	RltStgysFstIdx struct {
		X RltStgysAct
	}
	RltStgysMdlIdx struct {
		X RltStgysAct
	}
	RltStgysLstIdx struct {
		X RltStgysAct
	}
	RltStgysRev struct {
		X RltStgysAct
	}
	ClrClrOpa struct {
		X  ClrClrAct
		I0 FltFltAct
	}
	ClrClrInv struct {
		X ClrClrAct
	}
	PenPenOpa struct {
		X  PenPenAct
		I0 FltFltAct
	}
	PenPenInv struct {
		X PenPenAct
	}
	PenPensCnt struct {
		X PenPensAct
	}
	PenPensCpy struct {
		X PenPensAct
	}
	PenPensClr struct {
		X PenPensAct
	}
	PenPensRand struct {
		X PenPensAct
	}
	PenPensMrg struct {
		X  PenPensAct
		I0 []PenPensAct
	}
	PenPensPush struct {
		X  PenPensAct
		I0 []PenPenAct
	}
	PenPensPop struct {
		X PenPensAct
	}
	PenPensQue struct {
		X  PenPensAct
		I0 []PenPenAct
	}
	PenPensDque struct {
		X PenPensAct
	}
	PenPensIns struct {
		X  PenPensAct
		I0 UntUntAct
		I1 PenPenAct
	}
	PenPensUpd struct {
		X  PenPensAct
		I0 UntUntAct
		I1 PenPenAct
	}
	PenPensDel struct {
		X  PenPensAct
		I0 UntUntAct
	}
	PenPensAt struct {
		X  PenPensAct
		I0 UntUntAct
	}
	PenPensIn struct {
		X  PenPensAct
		I0 UntUntAct
		I1 UntUntAct
	}
	PenPensInBnd struct {
		X  PenPensAct
		I0 BndBndAct
	}
	PenPensFrom struct {
		X  PenPensAct
		I0 UntUntAct
	}
	PenPensTo struct {
		X  PenPensAct
		I0 UntUntAct
	}
	PenPensFst struct {
		X PenPensAct
	}
	PenPensMdl struct {
		X PenPensAct
	}
	PenPensLst struct {
		X PenPensAct
	}
	PenPensFstIdx struct {
		X PenPensAct
	}
	PenPensMdlIdx struct {
		X PenPensAct
	}
	PenPensLstIdx struct {
		X PenPensAct
	}
	PenPensRev struct {
		X PenPensAct
	}
	PltPltsCnt struct {
		X PltPltsAct
	}
	PltPltsCpy struct {
		X PltPltsAct
	}
	PltPltsClr struct {
		X PltPltsAct
	}
	PltPltsRand struct {
		X PltPltsAct
	}
	PltPltsMrg struct {
		X  PltPltsAct
		I0 []PltPltsAct
	}
	PltPltsPush struct {
		X  PltPltsAct
		I0 []PltPltAct
	}
	PltPltsPop struct {
		X PltPltsAct
	}
	PltPltsQue struct {
		X  PltPltsAct
		I0 []PltPltAct
	}
	PltPltsDque struct {
		X PltPltsAct
	}
	PltPltsIns struct {
		X  PltPltsAct
		I0 UntUntAct
		I1 PltPltAct
	}
	PltPltsUpd struct {
		X  PltPltsAct
		I0 UntUntAct
		I1 PltPltAct
	}
	PltPltsDel struct {
		X  PltPltsAct
		I0 UntUntAct
	}
	PltPltsAt struct {
		X  PltPltsAct
		I0 UntUntAct
	}
	PltPltsIn struct {
		X  PltPltsAct
		I0 UntUntAct
		I1 UntUntAct
	}
	PltPltsInBnd struct {
		X  PltPltsAct
		I0 BndBndAct
	}
	PltPltsFrom struct {
		X  PltPltsAct
		I0 UntUntAct
	}
	PltPltsTo struct {
		X  PltPltsAct
		I0 UntUntAct
	}
	PltPltsFst struct {
		X PltPltsAct
	}
	PltPltsMdl struct {
		X PltPltsAct
	}
	PltPltsLst struct {
		X PltPltsAct
	}
	PltPltsFstIdx struct {
		X PltPltsAct
	}
	PltPltsMdlIdx struct {
		X PltPltsAct
	}
	PltPltsLstIdx struct {
		X PltPltsAct
	}
	PltPltsRev struct {
		X PltPltsAct
	}
	PltTmeAxisXVis struct {
		X  PltTmeAxisXAct
		I0 BolBolAct
	}
	PltFltAxisYVis struct {
		X  PltFltAxisYAct
		I0 BolBolAct
	}
	PltStmX struct {
		X PltStmAct
	}
	PltStmY struct {
		X PltStmAct
	}
	PltStmStm struct {
		X  PltStmAct
		I0 PenPenAct
		I1 []HstStmAct
	}
	PltStmStmBnd struct {
		X  PltStmAct
		I0 ClrClrAct
		I1 PenPenAct
		I2 HstStmAct
		I3 HstStmAct
	}
	PltStmCnd struct {
		X  PltStmAct
		I0 PenPenAct
		I1 []HstCndAct
	}
	PltStmHrzLn struct {
		X  PltStmAct
		I0 PenPenAct
		I1 []FltFltAct
	}
	PltStmVrtLn struct {
		X  PltStmAct
		I0 PenPenAct
		I1 []TmeTmeAct
	}
	PltStmHrzBnd struct {
		X  PltStmAct
		I0 ClrClrAct
		I1 PenPenAct
		I2 FltFltAct
		I3 FltFltAct
	}
	PltStmVrtBnd struct {
		X  PltStmAct
		I0 ClrClrAct
		I1 PenPenAct
		I2 TmeTmeAct
		I3 TmeTmeAct
	}
	PltStmHrzSclVal struct {
		X  PltStmAct
		I0 TmeTmeAct
	}
	PltStmVrtSclVal struct {
		X  PltStmAct
		I0 FltFltAct
	}
	PltStmSho struct {
		X PltStmAct
	}
	PltStmSiz struct {
		X  PltStmAct
		I0 UntUntAct
		I1 UntUntAct
	}
	PltStmScl struct {
		X  PltStmAct
		I0 FltFltAct
	}
	PltStmHrzScl struct {
		X  PltStmAct
		I0 FltFltAct
	}
	PltStmVrtScl struct {
		X  PltStmAct
		I0 FltFltAct
	}
	PltFltsSctrFlts struct {
		X  PltFltsSctrAct
		I0 ClrClrAct
		I1 []FltsFltsAct
	}
	PltFltsSctrPrfLos struct {
		X  PltFltsSctrAct
		I0 TmesTmesAct
		I1 TmesTmesAct
		I2 []HstStmAct
	}
	PltFltsSctrSho struct {
		X PltFltsSctrAct
	}
	PltFltsSctrSiz struct {
		X  PltFltsSctrAct
		I0 UntUntAct
		I1 UntUntAct
	}
	PltFltsSctrScl struct {
		X  PltFltsSctrAct
		I0 FltFltAct
	}
	PltFltsSctrHrzScl struct {
		X  PltFltsSctrAct
		I0 FltFltAct
	}
	PltFltsSctrVrtScl struct {
		X  PltFltsSctrAct
		I0 FltFltAct
	}
	PltFltsSctrDistFlts struct {
		X  PltFltsSctrDistAct
		I0 ClrClrAct
		I1 UntUntAct
		I2 []FltsFltsAct
	}
	PltFltsSctrDistSho struct {
		X PltFltsSctrDistAct
	}
	PltFltsSctrDistSiz struct {
		X  PltFltsSctrDistAct
		I0 UntUntAct
		I1 UntUntAct
	}
	PltFltsSctrDistScl struct {
		X  PltFltsSctrDistAct
		I0 FltFltAct
	}
	PltFltsSctrDistHrzScl struct {
		X  PltFltsSctrDistAct
		I0 FltFltAct
	}
	PltFltsSctrDistVrtScl struct {
		X  PltFltsSctrDistAct
		I0 FltFltAct
	}
	PltHrzPlt struct {
		X  PltHrzAct
		I0 []PltPltAct
	}
	PltHrzSho struct {
		X PltHrzAct
	}
	PltHrzSiz struct {
		X  PltHrzAct
		I0 UntUntAct
		I1 UntUntAct
	}
	PltHrzScl struct {
		X  PltHrzAct
		I0 FltFltAct
	}
	PltHrzHrzScl struct {
		X  PltHrzAct
		I0 FltFltAct
	}
	PltHrzVrtScl struct {
		X  PltHrzAct
		I0 FltFltAct
	}
	PltVrtPlt struct {
		X  PltVrtAct
		I0 []PltPltAct
	}
	PltVrtSho struct {
		X PltVrtAct
	}
	PltVrtSiz struct {
		X  PltVrtAct
		I0 UntUntAct
		I1 UntUntAct
	}
	PltVrtScl struct {
		X  PltVrtAct
		I0 FltFltAct
	}
	PltVrtHrzScl struct {
		X  PltVrtAct
		I0 FltFltAct
	}
	PltVrtVrtScl struct {
		X  PltVrtAct
		I0 FltFltAct
	}
	PltDpthPlt struct {
		X  PltDpthAct
		I0 []PltPltAct
	}
	PltDpthSho struct {
		X PltDpthAct
	}
	PltDpthSiz struct {
		X  PltDpthAct
		I0 UntUntAct
		I1 UntUntAct
	}
	PltDpthScl struct {
		X  PltDpthAct
		I0 FltFltAct
	}
	PltDpthHrzScl struct {
		X  PltDpthAct
		I0 FltFltAct
	}
	PltDpthVrtScl struct {
		X  PltDpthAct
		I0 FltFltAct
	}
	SysMuLck struct {
		X SysMuAct
	}
	SysMuUlck struct {
		X SysMuAct
	}
	HstPrvName struct {
		X HstPrvAct
	}
	HstPrvEurUsd struct {
		X  HstPrvAct
		I0 []TmeRngAct
	}
	HstPrvAudUsd struct {
		X  HstPrvAct
		I0 []TmeRngAct
	}
	HstPrvNzdUsd struct {
		X  HstPrvAct
		I0 []TmeRngAct
	}
	HstPrvGbpUsd struct {
		X  HstPrvAct
		I0 []TmeRngAct
	}
	HstInstrName struct {
		X HstInstrAct
	}
	HstInstrI struct {
		X  HstInstrAct
		I0 TmeTmeAct
	}
	HstInrvlName struct {
		X HstInrvlAct
	}
	HstInrvlBid struct {
		X HstInrvlAct
	}
	HstInrvlAsk struct {
		X HstInrvlAct
	}
	HstSideName struct {
		X HstSideAct
	}
	HstSideFst struct {
		X HstSideAct
	}
	HstSideLst struct {
		X HstSideAct
	}
	HstSideSum struct {
		X HstSideAct
	}
	HstSidePrd struct {
		X HstSideAct
	}
	HstSideMin struct {
		X HstSideAct
	}
	HstSideMax struct {
		X HstSideAct
	}
	HstSideMid struct {
		X HstSideAct
	}
	HstSideMdn struct {
		X HstSideAct
	}
	HstSideSma struct {
		X HstSideAct
	}
	HstSideGma struct {
		X HstSideAct
	}
	HstSideWma struct {
		X HstSideAct
	}
	HstSideRsi struct {
		X HstSideAct
	}
	HstSideWrsi struct {
		X HstSideAct
	}
	HstSideAlma struct {
		X HstSideAct
	}
	HstSideVrnc struct {
		X HstSideAct
	}
	HstSideStd struct {
		X HstSideAct
	}
	HstSideRngFul struct {
		X HstSideAct
	}
	HstSideRngLst struct {
		X HstSideAct
	}
	HstSideProLst struct {
		X HstSideAct
	}
	HstSideProSma struct {
		X HstSideAct
	}
	HstSideProAlma struct {
		X HstSideAct
	}
	HstSideSar struct {
		X  HstSideAct
		I0 FltFltAct
		I1 FltFltAct
	}
	HstSideEma struct {
		X HstSideAct
	}
	HstStmName struct {
		X HstStmAct
	}
	HstStmAt struct {
		X  HstStmAct
		I0 TmesTmesAct
	}
	HstStmUnaPos struct {
		X HstStmAct
	}
	HstStmUnaNeg struct {
		X HstStmAct
	}
	HstStmUnaInv struct {
		X HstStmAct
	}
	HstStmUnaSqr struct {
		X HstStmAct
	}
	HstStmUnaSqrt struct {
		X HstStmAct
	}
	HstStmSclAdd struct {
		X  HstStmAct
		I0 FltFltAct
	}
	HstStmSclSub struct {
		X  HstStmAct
		I0 FltFltAct
	}
	HstStmSclMul struct {
		X  HstStmAct
		I0 FltFltAct
	}
	HstStmSclDiv struct {
		X  HstStmAct
		I0 FltFltAct
	}
	HstStmSclRem struct {
		X  HstStmAct
		I0 FltFltAct
	}
	HstStmSclPow struct {
		X  HstStmAct
		I0 FltFltAct
	}
	HstStmSclMin struct {
		X  HstStmAct
		I0 FltFltAct
	}
	HstStmSclMax struct {
		X  HstStmAct
		I0 FltFltAct
	}
	HstStmSelEql struct {
		X  HstStmAct
		I0 FltFltAct
	}
	HstStmSelNeq struct {
		X  HstStmAct
		I0 FltFltAct
	}
	HstStmSelLss struct {
		X  HstStmAct
		I0 FltFltAct
	}
	HstStmSelGtr struct {
		X  HstStmAct
		I0 FltFltAct
	}
	HstStmSelLeq struct {
		X  HstStmAct
		I0 FltFltAct
	}
	HstStmSelGeq struct {
		X  HstStmAct
		I0 FltFltAct
	}
	HstStmAggFst struct {
		X  HstStmAct
		I0 UntUntAct
	}
	HstStmAggLst struct {
		X  HstStmAct
		I0 UntUntAct
	}
	HstStmAggSum struct {
		X  HstStmAct
		I0 UntUntAct
	}
	HstStmAggPrd struct {
		X  HstStmAct
		I0 UntUntAct
	}
	HstStmAggMin struct {
		X  HstStmAct
		I0 UntUntAct
	}
	HstStmAggMax struct {
		X  HstStmAct
		I0 UntUntAct
	}
	HstStmAggMid struct {
		X  HstStmAct
		I0 UntUntAct
	}
	HstStmAggMdn struct {
		X  HstStmAct
		I0 UntUntAct
	}
	HstStmAggSma struct {
		X  HstStmAct
		I0 UntUntAct
	}
	HstStmAggGma struct {
		X  HstStmAct
		I0 UntUntAct
	}
	HstStmAggWma struct {
		X  HstStmAct
		I0 UntUntAct
	}
	HstStmAggRsi struct {
		X  HstStmAct
		I0 UntUntAct
	}
	HstStmAggWrsi struct {
		X  HstStmAct
		I0 UntUntAct
	}
	HstStmAggAlma struct {
		X  HstStmAct
		I0 UntUntAct
	}
	HstStmAggVrnc struct {
		X  HstStmAct
		I0 UntUntAct
	}
	HstStmAggStd struct {
		X  HstStmAct
		I0 UntUntAct
	}
	HstStmAggRngFul struct {
		X  HstStmAct
		I0 UntUntAct
	}
	HstStmAggRngLst struct {
		X  HstStmAct
		I0 UntUntAct
	}
	HstStmAggProLst struct {
		X  HstStmAct
		I0 UntUntAct
	}
	HstStmAggProSma struct {
		X  HstStmAct
		I0 UntUntAct
	}
	HstStmAggProAlma struct {
		X  HstStmAct
		I0 UntUntAct
	}
	HstStmAggEma struct {
		X  HstStmAct
		I0 UntUntAct
	}
	HstStmInrAdd struct {
		X  HstStmAct
		I0 UntUntAct
	}
	HstStmInrSub struct {
		X  HstStmAct
		I0 UntUntAct
	}
	HstStmInrMul struct {
		X  HstStmAct
		I0 UntUntAct
	}
	HstStmInrDiv struct {
		X  HstStmAct
		I0 UntUntAct
	}
	HstStmInrRem struct {
		X  HstStmAct
		I0 UntUntAct
	}
	HstStmInrPow struct {
		X  HstStmAct
		I0 UntUntAct
	}
	HstStmInrMin struct {
		X  HstStmAct
		I0 UntUntAct
	}
	HstStmInrMax struct {
		X  HstStmAct
		I0 UntUntAct
	}
	HstStmInrSlp struct {
		X  HstStmAct
		I0 UntUntAct
	}
	HstStmOtrAdd struct {
		X  HstStmAct
		I0 UntUntAct
		I1 HstStmAct
	}
	HstStmOtrSub struct {
		X  HstStmAct
		I0 UntUntAct
		I1 HstStmAct
	}
	HstStmOtrMul struct {
		X  HstStmAct
		I0 UntUntAct
		I1 HstStmAct
	}
	HstStmOtrDiv struct {
		X  HstStmAct
		I0 UntUntAct
		I1 HstStmAct
	}
	HstStmOtrRem struct {
		X  HstStmAct
		I0 UntUntAct
		I1 HstStmAct
	}
	HstStmOtrPow struct {
		X  HstStmAct
		I0 UntUntAct
		I1 HstStmAct
	}
	HstStmOtrMin struct {
		X  HstStmAct
		I0 UntUntAct
		I1 HstStmAct
	}
	HstStmOtrMax struct {
		X  HstStmAct
		I0 UntUntAct
		I1 HstStmAct
	}
	HstStmSclEql struct {
		X  HstStmAct
		I0 FltFltAct
	}
	HstStmSclNeq struct {
		X  HstStmAct
		I0 FltFltAct
	}
	HstStmSclLss struct {
		X  HstStmAct
		I0 FltFltAct
	}
	HstStmSclGtr struct {
		X  HstStmAct
		I0 FltFltAct
	}
	HstStmSclLeq struct {
		X  HstStmAct
		I0 FltFltAct
	}
	HstStmSclGeq struct {
		X  HstStmAct
		I0 FltFltAct
	}
	HstStmInrEql struct {
		X  HstStmAct
		I0 UntUntAct
	}
	HstStmInrNeq struct {
		X  HstStmAct
		I0 UntUntAct
	}
	HstStmInrLss struct {
		X  HstStmAct
		I0 UntUntAct
	}
	HstStmInrGtr struct {
		X  HstStmAct
		I0 UntUntAct
	}
	HstStmInrLeq struct {
		X  HstStmAct
		I0 UntUntAct
	}
	HstStmInrGeq struct {
		X  HstStmAct
		I0 UntUntAct
	}
	HstStmOtrEql struct {
		X  HstStmAct
		I0 UntUntAct
		I1 HstStmAct
	}
	HstStmOtrNeq struct {
		X  HstStmAct
		I0 UntUntAct
		I1 HstStmAct
	}
	HstStmOtrLss struct {
		X  HstStmAct
		I0 UntUntAct
		I1 HstStmAct
	}
	HstStmOtrGtr struct {
		X  HstStmAct
		I0 UntUntAct
		I1 HstStmAct
	}
	HstStmOtrLeq struct {
		X  HstStmAct
		I0 UntUntAct
		I1 HstStmAct
	}
	HstStmOtrGeq struct {
		X  HstStmAct
		I0 UntUntAct
		I1 HstStmAct
	}
	HstCndName struct {
		X HstCndAct
	}
	HstCndAnd struct {
		X  HstCndAct
		I0 HstCndAct
	}
	HstCndSeq struct {
		X  HstCndAct
		I0 TmeTmeAct
		I1 HstCndAct
	}
	HstCndStgy struct {
		X  HstCndAct
		I0 BolBolAct
		I1 FltFltAct
		I2 FltFltAct
		I3 TmeTmeAct
		I4 FltFltAct
		I5 HstInstrAct
		I6 HstStmsAct
		I7 []HstCndAct
	}
	HstStgyName struct {
		X HstStgyAct
	}
	RltPrvMayTrd struct {
		X RltPrvAct
	}
	RltPrvEurUsd struct {
		X  RltPrvAct
		I0 []TmeRngAct
	}
	RltPrvAudUsd struct {
		X  RltPrvAct
		I0 []TmeRngAct
	}
	RltPrvNzdUsd struct {
		X  RltPrvAct
		I0 []TmeRngAct
	}
	RltPrvGbpUsd struct {
		X  RltPrvAct
		I0 []TmeRngAct
	}
	RltInstrI struct {
		X  RltInstrAct
		I0 TmeTmeAct
	}
	RltInrvlBid struct {
		X RltInrvlAct
	}
	RltInrvlAsk struct {
		X RltInrvlAct
	}
	RltSideFst struct {
		X RltSideAct
	}
	RltSideLst struct {
		X RltSideAct
	}
	RltSideSum struct {
		X RltSideAct
	}
	RltSidePrd struct {
		X RltSideAct
	}
	RltSideMin struct {
		X RltSideAct
	}
	RltSideMax struct {
		X RltSideAct
	}
	RltSideMid struct {
		X RltSideAct
	}
	RltSideMdn struct {
		X RltSideAct
	}
	RltSideSma struct {
		X RltSideAct
	}
	RltSideGma struct {
		X RltSideAct
	}
	RltSideWma struct {
		X RltSideAct
	}
	RltSideRsi struct {
		X RltSideAct
	}
	RltSideWrsi struct {
		X RltSideAct
	}
	RltSideAlma struct {
		X RltSideAct
	}
	RltSideVrnc struct {
		X RltSideAct
	}
	RltSideStd struct {
		X RltSideAct
	}
	RltSideRngFul struct {
		X RltSideAct
	}
	RltSideRngLst struct {
		X RltSideAct
	}
	RltSideProLst struct {
		X RltSideAct
	}
	RltSideProSma struct {
		X RltSideAct
	}
	RltSideProAlma struct {
		X RltSideAct
	}
	RltSideSar struct {
		X  RltSideAct
		I0 FltFltAct
		I1 FltFltAct
	}
	RltSideEma struct {
		X RltSideAct
	}
	RltStmUnaPos struct {
		X RltStmAct
	}
	RltStmUnaNeg struct {
		X RltStmAct
	}
	RltStmUnaInv struct {
		X RltStmAct
	}
	RltStmUnaSqr struct {
		X RltStmAct
	}
	RltStmUnaSqrt struct {
		X RltStmAct
	}
	RltStmSclAdd struct {
		X  RltStmAct
		I0 FltFltAct
	}
	RltStmSclSub struct {
		X  RltStmAct
		I0 FltFltAct
	}
	RltStmSclMul struct {
		X  RltStmAct
		I0 FltFltAct
	}
	RltStmSclDiv struct {
		X  RltStmAct
		I0 FltFltAct
	}
	RltStmSclRem struct {
		X  RltStmAct
		I0 FltFltAct
	}
	RltStmSclPow struct {
		X  RltStmAct
		I0 FltFltAct
	}
	RltStmSclMin struct {
		X  RltStmAct
		I0 FltFltAct
	}
	RltStmSclMax struct {
		X  RltStmAct
		I0 FltFltAct
	}
	RltStmSelEql struct {
		X  RltStmAct
		I0 FltFltAct
	}
	RltStmSelNeq struct {
		X  RltStmAct
		I0 FltFltAct
	}
	RltStmSelLss struct {
		X  RltStmAct
		I0 FltFltAct
	}
	RltStmSelGtr struct {
		X  RltStmAct
		I0 FltFltAct
	}
	RltStmSelLeq struct {
		X  RltStmAct
		I0 FltFltAct
	}
	RltStmSelGeq struct {
		X  RltStmAct
		I0 FltFltAct
	}
	RltStmAggFst struct {
		X  RltStmAct
		I0 UntUntAct
	}
	RltStmAggLst struct {
		X  RltStmAct
		I0 UntUntAct
	}
	RltStmAggSum struct {
		X  RltStmAct
		I0 UntUntAct
	}
	RltStmAggPrd struct {
		X  RltStmAct
		I0 UntUntAct
	}
	RltStmAggMin struct {
		X  RltStmAct
		I0 UntUntAct
	}
	RltStmAggMax struct {
		X  RltStmAct
		I0 UntUntAct
	}
	RltStmAggMid struct {
		X  RltStmAct
		I0 UntUntAct
	}
	RltStmAggMdn struct {
		X  RltStmAct
		I0 UntUntAct
	}
	RltStmAggSma struct {
		X  RltStmAct
		I0 UntUntAct
	}
	RltStmAggGma struct {
		X  RltStmAct
		I0 UntUntAct
	}
	RltStmAggWma struct {
		X  RltStmAct
		I0 UntUntAct
	}
	RltStmAggRsi struct {
		X  RltStmAct
		I0 UntUntAct
	}
	RltStmAggWrsi struct {
		X  RltStmAct
		I0 UntUntAct
	}
	RltStmAggAlma struct {
		X  RltStmAct
		I0 UntUntAct
	}
	RltStmAggVrnc struct {
		X  RltStmAct
		I0 UntUntAct
	}
	RltStmAggStd struct {
		X  RltStmAct
		I0 UntUntAct
	}
	RltStmAggRngFul struct {
		X  RltStmAct
		I0 UntUntAct
	}
	RltStmAggRngLst struct {
		X  RltStmAct
		I0 UntUntAct
	}
	RltStmAggProLst struct {
		X  RltStmAct
		I0 UntUntAct
	}
	RltStmAggProSma struct {
		X  RltStmAct
		I0 UntUntAct
	}
	RltStmAggProAlma struct {
		X  RltStmAct
		I0 UntUntAct
	}
	RltStmAggEma struct {
		X  RltStmAct
		I0 UntUntAct
	}
	RltStmInrAdd struct {
		X  RltStmAct
		I0 UntUntAct
	}
	RltStmInrSub struct {
		X  RltStmAct
		I0 UntUntAct
	}
	RltStmInrMul struct {
		X  RltStmAct
		I0 UntUntAct
	}
	RltStmInrDiv struct {
		X  RltStmAct
		I0 UntUntAct
	}
	RltStmInrRem struct {
		X  RltStmAct
		I0 UntUntAct
	}
	RltStmInrPow struct {
		X  RltStmAct
		I0 UntUntAct
	}
	RltStmInrMin struct {
		X  RltStmAct
		I0 UntUntAct
	}
	RltStmInrMax struct {
		X  RltStmAct
		I0 UntUntAct
	}
	RltStmInrSlp struct {
		X  RltStmAct
		I0 UntUntAct
	}
	RltStmOtrAdd struct {
		X  RltStmAct
		I0 UntUntAct
		I1 RltStmAct
	}
	RltStmOtrSub struct {
		X  RltStmAct
		I0 UntUntAct
		I1 RltStmAct
	}
	RltStmOtrMul struct {
		X  RltStmAct
		I0 UntUntAct
		I1 RltStmAct
	}
	RltStmOtrDiv struct {
		X  RltStmAct
		I0 UntUntAct
		I1 RltStmAct
	}
	RltStmOtrRem struct {
		X  RltStmAct
		I0 UntUntAct
		I1 RltStmAct
	}
	RltStmOtrPow struct {
		X  RltStmAct
		I0 UntUntAct
		I1 RltStmAct
	}
	RltStmOtrMin struct {
		X  RltStmAct
		I0 UntUntAct
		I1 RltStmAct
	}
	RltStmOtrMax struct {
		X  RltStmAct
		I0 UntUntAct
		I1 RltStmAct
	}
	RltStmSclEql struct {
		X  RltStmAct
		I0 FltFltAct
	}
	RltStmSclNeq struct {
		X  RltStmAct
		I0 FltFltAct
	}
	RltStmSclLss struct {
		X  RltStmAct
		I0 FltFltAct
	}
	RltStmSclGtr struct {
		X  RltStmAct
		I0 FltFltAct
	}
	RltStmSclLeq struct {
		X  RltStmAct
		I0 FltFltAct
	}
	RltStmSclGeq struct {
		X  RltStmAct
		I0 FltFltAct
	}
	RltStmInrEql struct {
		X  RltStmAct
		I0 UntUntAct
	}
	RltStmInrNeq struct {
		X  RltStmAct
		I0 UntUntAct
	}
	RltStmInrLss struct {
		X  RltStmAct
		I0 UntUntAct
	}
	RltStmInrGtr struct {
		X  RltStmAct
		I0 UntUntAct
	}
	RltStmInrLeq struct {
		X  RltStmAct
		I0 UntUntAct
	}
	RltStmInrGeq struct {
		X  RltStmAct
		I0 UntUntAct
	}
	RltStmOtrEql struct {
		X  RltStmAct
		I0 UntUntAct
		I1 RltStmAct
	}
	RltStmOtrNeq struct {
		X  RltStmAct
		I0 UntUntAct
		I1 RltStmAct
	}
	RltStmOtrLss struct {
		X  RltStmAct
		I0 UntUntAct
		I1 RltStmAct
	}
	RltStmOtrGtr struct {
		X  RltStmAct
		I0 UntUntAct
		I1 RltStmAct
	}
	RltStmOtrLeq struct {
		X  RltStmAct
		I0 UntUntAct
		I1 RltStmAct
	}
	RltStmOtrGeq struct {
		X  RltStmAct
		I0 UntUntAct
		I1 RltStmAct
	}
	RltCndAnd struct {
		X  RltCndAct
		I0 RltCndAct
	}
	RltCndSeq struct {
		X  RltCndAct
		I0 TmeTmeAct
		I1 RltCndAct
	}
	RltCndStgy struct {
		X  RltCndAct
		I0 BolBolAct
		I1 FltFltAct
		I2 FltFltAct
		I3 TmeTmeAct
		I4 FltFltAct
		I5 RltInstrAct
		I6 RltStmsAct
		I7 []RltCndAct
	}
	PltPltSho struct {
		X PltPltAct
	}
	PltPltSiz struct {
		X  PltPltAct
		I0 UntUntAct
		I1 UntUntAct
	}
	PltPltScl struct {
		X  PltPltAct
		I0 FltFltAct
	}
	PltPltHrzScl struct {
		X  PltPltAct
		I0 FltFltAct
	}
	PltPltVrtScl struct {
		X  PltPltAct
		I0 FltFltAct
	}
	IfcIfc struct {
		X Act
	}
	PllWait struct {
		Txt        string
		ActScpPrnt *Scp
		XprScp     *xpr.Scp
		Xprs       []xpr.Xpr
		Wg         *sync.WaitGroup
	}
	PllWaitSeg struct {
		Txt        string
		ActScpPrnt *Scp
		XprScp     *xpr.Scp
		Xpr        xpr.Xpr
		Wg         *sync.WaitGroup
	}
)

func Run(txt string) {
	var actr Actr
	actr.Run(txt)
}
func (x StrStrLit) Act()                   { x.StrStr() }
func (x StrStrLit) Ifc() interface{}       { return x.StrStr() }
func (x StrStrLit) StrStr() str.Str        { return prs.StrTrm(x.Trm, x.Txt) }
func (x BolBolLit) Act()                   { x.BolBol() }
func (x BolBolLit) Ifc() interface{}       { return x.BolBol() }
func (x BolBolLit) BolBol() bol.Bol        { return prs.BolTrm(x.Trm, x.Txt) }
func (x FltFltLit) Act()                   { x.FltFlt() }
func (x FltFltLit) Ifc() interface{}       { return x.FltFlt() }
func (x FltFltLit) FltFlt() flt.Flt        { return prs.FltTrm(x.Trm, x.Txt) }
func (x UntUntLit) Act()                   { x.UntUnt() }
func (x UntUntLit) Ifc() interface{}       { return x.UntUnt() }
func (x UntUntLit) UntUnt() unt.Unt        { return prs.UntTrm(x.Trm, x.Txt) }
func (x IntIntLit) Act()                   { x.IntInt() }
func (x IntIntLit) Ifc() interface{}       { return x.IntInt() }
func (x IntIntLit) IntInt() int.Int        { return prs.IntTrm(x.Trm, x.Txt) }
func (x TmeTmeLit) Act()                   { x.TmeTme() }
func (x TmeTmeLit) Ifc() interface{}       { return x.TmeTme() }
func (x TmeTmeLit) TmeTme() tme.Tme        { return prs.TmeTrm(x.Trm, x.Txt) }
func (x BndBndLit) Act()                   { x.BndBnd() }
func (x BndBndLit) Ifc() interface{}       { return x.BndBnd() }
func (x BndBndLit) BndBnd() bnd.Bnd        { return prs.BndTrm(x.Trm, x.Txt) }
func (x FltRngLit) Act()                   { x.FltRng() }
func (x FltRngLit) Ifc() interface{}       { return x.FltRng() }
func (x FltRngLit) FltRng() flt.Rng        { return prs.FltRngTrm(x.Trm, x.Txt) }
func (x TmeRngLit) Act()                   { x.TmeRng() }
func (x TmeRngLit) Ifc() interface{}       { return x.TmeRng() }
func (x TmeRngLit) TmeRng() tme.Rng        { return prs.TmeRngTrm(x.Trm, x.Txt) }
func (x StrsStrsLit) Act()                 { x.StrsStrs() }
func (x StrsStrsLit) Ifc() interface{}     { return x.StrsStrs() }
func (x StrsStrsLit) StrsStrs() *strs.Strs { return prs.StrsTrm(x.Trm, x.Txt) }
func (x BolsBolsLit) Act()                 { x.BolsBols() }
func (x BolsBolsLit) Ifc() interface{}     { return x.BolsBols() }
func (x BolsBolsLit) BolsBols() *bols.Bols { return prs.BolsTrm(x.Trm, x.Txt) }
func (x FltsFltsLit) Act()                 { x.FltsFlts() }
func (x FltsFltsLit) Ifc() interface{}     { return x.FltsFlts() }
func (x FltsFltsLit) FltsFlts() *flts.Flts { return prs.FltsTrm(x.Trm, x.Txt) }
func (x UntsUntsLit) Act()                 { x.UntsUnts() }
func (x UntsUntsLit) Ifc() interface{}     { return x.UntsUnts() }
func (x UntsUntsLit) UntsUnts() *unts.Unts { return prs.UntsTrm(x.Trm, x.Txt) }
func (x IntsIntsLit) Act()                 { x.IntsInts() }
func (x IntsIntsLit) Ifc() interface{}     { return x.IntsInts() }
func (x IntsIntsLit) IntsInts() *ints.Ints { return prs.IntsTrm(x.Trm, x.Txt) }
func (x TmesTmesLit) Act()                 { x.TmesTmes() }
func (x TmesTmesLit) Ifc() interface{}     { return x.TmesTmes() }
func (x TmesTmesLit) TmesTmes() *tmes.Tmes { return prs.TmesTrm(x.Trm, x.Txt) }
func (x BndsBndsLit) Act()                 { x.BndsBnds() }
func (x BndsBndsLit) Ifc() interface{}     { return x.BndsBnds() }
func (x BndsBndsLit) BndsBnds() *bnds.Bnds { return prs.BndsTrm(x.Trm, x.Txt) }
func (x TmeRngsLit) Act()                  { x.TmeRngs() }
func (x TmeRngsLit) Ifc() interface{}      { return x.TmeRngs() }
func (x TmeRngsLit) TmeRngs() *tme.Rngs    { return prs.TmeRngsTrm(x.Trm, x.Txt) }
func (x StrStrAsn) Act()                   { x.StrStr() }
func (x StrStrAsn) Ifc() interface{}       { return x.StrStr() }
func (x StrStrAsn) StrStr() str.Str {
	x.Arr[x.Idx] = x.X.StrStr()
	return x.Arr[x.Idx]
}
func (x StrStrAcs) Act()             { x.StrStr() }
func (x StrStrAcs) Ifc() interface{} { return x.StrStr() }
func (x StrStrAcs) StrStr() str.Str  { return x.Arr[x.Idx] }
func (x BolBolAsn) Act()             { x.BolBol() }
func (x BolBolAsn) Ifc() interface{} { return x.BolBol() }
func (x BolBolAsn) BolBol() bol.Bol {
	x.Arr[x.Idx] = x.X.BolBol()
	return x.Arr[x.Idx]
}
func (x BolBolAcs) Act()             { x.BolBol() }
func (x BolBolAcs) Ifc() interface{} { return x.BolBol() }
func (x BolBolAcs) BolBol() bol.Bol  { return x.Arr[x.Idx] }
func (x FltFltAsn) Act()             { x.FltFlt() }
func (x FltFltAsn) Ifc() interface{} { return x.FltFlt() }
func (x FltFltAsn) FltFlt() flt.Flt {
	x.Arr[x.Idx] = x.X.FltFlt()
	return x.Arr[x.Idx]
}
func (x FltFltAcs) Act()             { x.FltFlt() }
func (x FltFltAcs) Ifc() interface{} { return x.FltFlt() }
func (x FltFltAcs) FltFlt() flt.Flt  { return x.Arr[x.Idx] }
func (x UntUntAsn) Act()             { x.UntUnt() }
func (x UntUntAsn) Ifc() interface{} { return x.UntUnt() }
func (x UntUntAsn) UntUnt() unt.Unt {
	x.Arr[x.Idx] = x.X.UntUnt()
	return x.Arr[x.Idx]
}
func (x UntUntAcs) Act()             { x.UntUnt() }
func (x UntUntAcs) Ifc() interface{} { return x.UntUnt() }
func (x UntUntAcs) UntUnt() unt.Unt  { return x.Arr[x.Idx] }
func (x IntIntAsn) Act()             { x.IntInt() }
func (x IntIntAsn) Ifc() interface{} { return x.IntInt() }
func (x IntIntAsn) IntInt() int.Int {
	x.Arr[x.Idx] = x.X.IntInt()
	return x.Arr[x.Idx]
}
func (x IntIntAcs) Act()             { x.IntInt() }
func (x IntIntAcs) Ifc() interface{} { return x.IntInt() }
func (x IntIntAcs) IntInt() int.Int  { return x.Arr[x.Idx] }
func (x TmeTmeAsn) Act()             { x.TmeTme() }
func (x TmeTmeAsn) Ifc() interface{} { return x.TmeTme() }
func (x TmeTmeAsn) TmeTme() tme.Tme {
	x.Arr[x.Idx] = x.X.TmeTme()
	return x.Arr[x.Idx]
}
func (x TmeTmeAcs) Act()             { x.TmeTme() }
func (x TmeTmeAcs) Ifc() interface{} { return x.TmeTme() }
func (x TmeTmeAcs) TmeTme() tme.Tme  { return x.Arr[x.Idx] }
func (x BndBndAsn) Act()             { x.BndBnd() }
func (x BndBndAsn) Ifc() interface{} { return x.BndBnd() }
func (x BndBndAsn) BndBnd() bnd.Bnd {
	x.Arr[x.Idx] = x.X.BndBnd()
	return x.Arr[x.Idx]
}
func (x BndBndAcs) Act()             { x.BndBnd() }
func (x BndBndAcs) Ifc() interface{} { return x.BndBnd() }
func (x BndBndAcs) BndBnd() bnd.Bnd  { return x.Arr[x.Idx] }
func (x FltRngAsn) Act()             { x.FltRng() }
func (x FltRngAsn) Ifc() interface{} { return x.FltRng() }
func (x FltRngAsn) FltRng() flt.Rng {
	x.Arr[x.Idx] = x.X.FltRng()
	return x.Arr[x.Idx]
}
func (x FltRngAcs) Act()             { x.FltRng() }
func (x FltRngAcs) Ifc() interface{} { return x.FltRng() }
func (x FltRngAcs) FltRng() flt.Rng  { return x.Arr[x.Idx] }
func (x TmeRngAsn) Act()             { x.TmeRng() }
func (x TmeRngAsn) Ifc() interface{} { return x.TmeRng() }
func (x TmeRngAsn) TmeRng() tme.Rng {
	x.Arr[x.Idx] = x.X.TmeRng()
	return x.Arr[x.Idx]
}
func (x TmeRngAcs) Act()               { x.TmeRng() }
func (x TmeRngAcs) Ifc() interface{}   { return x.TmeRng() }
func (x TmeRngAcs) TmeRng() tme.Rng    { return x.Arr[x.Idx] }
func (x StrsStrsAsn) Act()             { x.StrsStrs() }
func (x StrsStrsAsn) Ifc() interface{} { return x.StrsStrs() }
func (x StrsStrsAsn) StrsStrs() *strs.Strs {
	x.Arr[x.Idx] = x.X.StrsStrs()
	return x.Arr[x.Idx]
}
func (x StrsStrsAcs) Act()                 { x.StrsStrs() }
func (x StrsStrsAcs) Ifc() interface{}     { return x.StrsStrs() }
func (x StrsStrsAcs) StrsStrs() *strs.Strs { return x.Arr[x.Idx] }
func (x StrsStrsEach) Act()                { x.StrsStrs() }
func (x StrsStrsEach) Ifc() interface{}    { return x.StrsStrs() }
func (x StrsStrsEach) StrsStrs() *strs.Strs {
	vs := x.X.StrsStrs()
	if vs != nil {
		for _, v := range *vs {
			x.Arr[x.Idx] = v // set cur elm to scp
			for _, a := range x.Acts {
				a.Act()
			}
		}
	}
	return vs
}
func (x StrsStrsPllEach) Act()             { x.StrsStrs() }
func (x StrsStrsPllEach) Ifc() interface{} { return x.StrsStrs() }
func (x *StrsStrsPllEachSeg) Act() {
	var actr Actr
	actr.Reset(x.Txt)
	scp := NewScp(x.XprScp, x.ActScpPrnt)
	elmScp := scp.StrStr(x.Txt[x.Idn.Idx:x.Idn.Lim])
	acts := actr.Acts(scp, x.Xprs...)
	elmScp.Arr[elmScp.Idx] = x.Val // set cur elm to scp
	for _, a := range acts {
		a.Act()
	}
}
func (x StrsStrsPllEach) StrsStrs() *strs.Strs {
	vs := x.X.StrsStrs()
	segs := make([]sys.Act, len(*vs))
	for n, v := range *vs {
		segs[n] = &StrsStrsPllEachSeg{
			Val:        v,
			Idn:        x.Idn,
			Txt:        x.Txt,
			ActScpPrnt: x.ActScpPrnt,
			XprScp:     x.XprScp,
			Xprs:       x.Xprs,
		}
	}
	sys.Run().Pll(segs...) // run segs in pll
	return vs
}
func (x BolsBolsAsn) Act()             { x.BolsBols() }
func (x BolsBolsAsn) Ifc() interface{} { return x.BolsBols() }
func (x BolsBolsAsn) BolsBols() *bols.Bols {
	x.Arr[x.Idx] = x.X.BolsBols()
	return x.Arr[x.Idx]
}
func (x BolsBolsAcs) Act()                 { x.BolsBols() }
func (x BolsBolsAcs) Ifc() interface{}     { return x.BolsBols() }
func (x BolsBolsAcs) BolsBols() *bols.Bols { return x.Arr[x.Idx] }
func (x BolsBolsEach) Act()                { x.BolsBols() }
func (x BolsBolsEach) Ifc() interface{}    { return x.BolsBols() }
func (x BolsBolsEach) BolsBols() *bols.Bols {
	vs := x.X.BolsBols()
	if vs != nil {
		for _, v := range *vs {
			x.Arr[x.Idx] = v // set cur elm to scp
			for _, a := range x.Acts {
				a.Act()
			}
		}
	}
	return vs
}
func (x BolsBolsPllEach) Act()             { x.BolsBols() }
func (x BolsBolsPllEach) Ifc() interface{} { return x.BolsBols() }
func (x *BolsBolsPllEachSeg) Act() {
	var actr Actr
	actr.Reset(x.Txt)
	scp := NewScp(x.XprScp, x.ActScpPrnt)
	elmScp := scp.BolBol(x.Txt[x.Idn.Idx:x.Idn.Lim])
	acts := actr.Acts(scp, x.Xprs...)
	elmScp.Arr[elmScp.Idx] = x.Val // set cur elm to scp
	for _, a := range acts {
		a.Act()
	}
}
func (x BolsBolsPllEach) BolsBols() *bols.Bols {
	vs := x.X.BolsBols()
	segs := make([]sys.Act, len(*vs))
	for n, v := range *vs {
		segs[n] = &BolsBolsPllEachSeg{
			Val:        v,
			Idn:        x.Idn,
			Txt:        x.Txt,
			ActScpPrnt: x.ActScpPrnt,
			XprScp:     x.XprScp,
			Xprs:       x.Xprs,
		}
	}
	sys.Run().Pll(segs...) // run segs in pll
	return vs
}
func (x FltsFltsAsn) Act()             { x.FltsFlts() }
func (x FltsFltsAsn) Ifc() interface{} { return x.FltsFlts() }
func (x FltsFltsAsn) FltsFlts() *flts.Flts {
	x.Arr[x.Idx] = x.X.FltsFlts()
	return x.Arr[x.Idx]
}
func (x FltsFltsAcs) Act()                 { x.FltsFlts() }
func (x FltsFltsAcs) Ifc() interface{}     { return x.FltsFlts() }
func (x FltsFltsAcs) FltsFlts() *flts.Flts { return x.Arr[x.Idx] }
func (x FltsFltsEach) Act()                { x.FltsFlts() }
func (x FltsFltsEach) Ifc() interface{}    { return x.FltsFlts() }
func (x FltsFltsEach) FltsFlts() *flts.Flts {
	vs := x.X.FltsFlts()
	if vs != nil {
		for _, v := range *vs {
			x.Arr[x.Idx] = v // set cur elm to scp
			for _, a := range x.Acts {
				a.Act()
			}
		}
	}
	return vs
}
func (x FltsFltsPllEach) Act()             { x.FltsFlts() }
func (x FltsFltsPllEach) Ifc() interface{} { return x.FltsFlts() }
func (x *FltsFltsPllEachSeg) Act() {
	var actr Actr
	actr.Reset(x.Txt)
	scp := NewScp(x.XprScp, x.ActScpPrnt)
	elmScp := scp.FltFlt(x.Txt[x.Idn.Idx:x.Idn.Lim])
	acts := actr.Acts(scp, x.Xprs...)
	elmScp.Arr[elmScp.Idx] = x.Val // set cur elm to scp
	for _, a := range acts {
		a.Act()
	}
}
func (x FltsFltsPllEach) FltsFlts() *flts.Flts {
	vs := x.X.FltsFlts()
	segs := make([]sys.Act, len(*vs))
	for n, v := range *vs {
		segs[n] = &FltsFltsPllEachSeg{
			Val:        v,
			Idn:        x.Idn,
			Txt:        x.Txt,
			ActScpPrnt: x.ActScpPrnt,
			XprScp:     x.XprScp,
			Xprs:       x.Xprs,
		}
	}
	sys.Run().Pll(segs...) // run segs in pll
	return vs
}
func (x UntsUntsAsn) Act()             { x.UntsUnts() }
func (x UntsUntsAsn) Ifc() interface{} { return x.UntsUnts() }
func (x UntsUntsAsn) UntsUnts() *unts.Unts {
	x.Arr[x.Idx] = x.X.UntsUnts()
	return x.Arr[x.Idx]
}
func (x UntsUntsAcs) Act()                 { x.UntsUnts() }
func (x UntsUntsAcs) Ifc() interface{}     { return x.UntsUnts() }
func (x UntsUntsAcs) UntsUnts() *unts.Unts { return x.Arr[x.Idx] }
func (x UntsUntsEach) Act()                { x.UntsUnts() }
func (x UntsUntsEach) Ifc() interface{}    { return x.UntsUnts() }
func (x UntsUntsEach) UntsUnts() *unts.Unts {
	vs := x.X.UntsUnts()
	if vs != nil {
		for _, v := range *vs {
			x.Arr[x.Idx] = v // set cur elm to scp
			for _, a := range x.Acts {
				a.Act()
			}
		}
	}
	return vs
}
func (x UntsUntsPllEach) Act()             { x.UntsUnts() }
func (x UntsUntsPllEach) Ifc() interface{} { return x.UntsUnts() }
func (x *UntsUntsPllEachSeg) Act() {
	var actr Actr
	actr.Reset(x.Txt)
	scp := NewScp(x.XprScp, x.ActScpPrnt)
	elmScp := scp.UntUnt(x.Txt[x.Idn.Idx:x.Idn.Lim])
	acts := actr.Acts(scp, x.Xprs...)
	elmScp.Arr[elmScp.Idx] = x.Val // set cur elm to scp
	for _, a := range acts {
		a.Act()
	}
}
func (x UntsUntsPllEach) UntsUnts() *unts.Unts {
	vs := x.X.UntsUnts()
	segs := make([]sys.Act, len(*vs))
	for n, v := range *vs {
		segs[n] = &UntsUntsPllEachSeg{
			Val:        v,
			Idn:        x.Idn,
			Txt:        x.Txt,
			ActScpPrnt: x.ActScpPrnt,
			XprScp:     x.XprScp,
			Xprs:       x.Xprs,
		}
	}
	sys.Run().Pll(segs...) // run segs in pll
	return vs
}
func (x IntsIntsAsn) Act()             { x.IntsInts() }
func (x IntsIntsAsn) Ifc() interface{} { return x.IntsInts() }
func (x IntsIntsAsn) IntsInts() *ints.Ints {
	x.Arr[x.Idx] = x.X.IntsInts()
	return x.Arr[x.Idx]
}
func (x IntsIntsAcs) Act()                 { x.IntsInts() }
func (x IntsIntsAcs) Ifc() interface{}     { return x.IntsInts() }
func (x IntsIntsAcs) IntsInts() *ints.Ints { return x.Arr[x.Idx] }
func (x IntsIntsEach) Act()                { x.IntsInts() }
func (x IntsIntsEach) Ifc() interface{}    { return x.IntsInts() }
func (x IntsIntsEach) IntsInts() *ints.Ints {
	vs := x.X.IntsInts()
	if vs != nil {
		for _, v := range *vs {
			x.Arr[x.Idx] = v // set cur elm to scp
			for _, a := range x.Acts {
				a.Act()
			}
		}
	}
	return vs
}
func (x IntsIntsPllEach) Act()             { x.IntsInts() }
func (x IntsIntsPllEach) Ifc() interface{} { return x.IntsInts() }
func (x *IntsIntsPllEachSeg) Act() {
	var actr Actr
	actr.Reset(x.Txt)
	scp := NewScp(x.XprScp, x.ActScpPrnt)
	elmScp := scp.IntInt(x.Txt[x.Idn.Idx:x.Idn.Lim])
	acts := actr.Acts(scp, x.Xprs...)
	elmScp.Arr[elmScp.Idx] = x.Val // set cur elm to scp
	for _, a := range acts {
		a.Act()
	}
}
func (x IntsIntsPllEach) IntsInts() *ints.Ints {
	vs := x.X.IntsInts()
	segs := make([]sys.Act, len(*vs))
	for n, v := range *vs {
		segs[n] = &IntsIntsPllEachSeg{
			Val:        v,
			Idn:        x.Idn,
			Txt:        x.Txt,
			ActScpPrnt: x.ActScpPrnt,
			XprScp:     x.XprScp,
			Xprs:       x.Xprs,
		}
	}
	sys.Run().Pll(segs...) // run segs in pll
	return vs
}
func (x TmesTmesAsn) Act()             { x.TmesTmes() }
func (x TmesTmesAsn) Ifc() interface{} { return x.TmesTmes() }
func (x TmesTmesAsn) TmesTmes() *tmes.Tmes {
	x.Arr[x.Idx] = x.X.TmesTmes()
	return x.Arr[x.Idx]
}
func (x TmesTmesAcs) Act()                 { x.TmesTmes() }
func (x TmesTmesAcs) Ifc() interface{}     { return x.TmesTmes() }
func (x TmesTmesAcs) TmesTmes() *tmes.Tmes { return x.Arr[x.Idx] }
func (x TmesTmesEach) Act()                { x.TmesTmes() }
func (x TmesTmesEach) Ifc() interface{}    { return x.TmesTmes() }
func (x TmesTmesEach) TmesTmes() *tmes.Tmes {
	vs := x.X.TmesTmes()
	if vs != nil {
		for _, v := range *vs {
			x.Arr[x.Idx] = v // set cur elm to scp
			for _, a := range x.Acts {
				a.Act()
			}
		}
	}
	return vs
}
func (x TmesTmesPllEach) Act()             { x.TmesTmes() }
func (x TmesTmesPllEach) Ifc() interface{} { return x.TmesTmes() }
func (x *TmesTmesPllEachSeg) Act() {
	var actr Actr
	actr.Reset(x.Txt)
	scp := NewScp(x.XprScp, x.ActScpPrnt)
	elmScp := scp.TmeTme(x.Txt[x.Idn.Idx:x.Idn.Lim])
	acts := actr.Acts(scp, x.Xprs...)
	elmScp.Arr[elmScp.Idx] = x.Val // set cur elm to scp
	for _, a := range acts {
		a.Act()
	}
}
func (x TmesTmesPllEach) TmesTmes() *tmes.Tmes {
	vs := x.X.TmesTmes()
	segs := make([]sys.Act, len(*vs))
	for n, v := range *vs {
		segs[n] = &TmesTmesPllEachSeg{
			Val:        v,
			Idn:        x.Idn,
			Txt:        x.Txt,
			ActScpPrnt: x.ActScpPrnt,
			XprScp:     x.XprScp,
			Xprs:       x.Xprs,
		}
	}
	sys.Run().Pll(segs...) // run segs in pll
	return vs
}
func (x BndsBndsAsn) Act()             { x.BndsBnds() }
func (x BndsBndsAsn) Ifc() interface{} { return x.BndsBnds() }
func (x BndsBndsAsn) BndsBnds() *bnds.Bnds {
	x.Arr[x.Idx] = x.X.BndsBnds()
	return x.Arr[x.Idx]
}
func (x BndsBndsAcs) Act()                 { x.BndsBnds() }
func (x BndsBndsAcs) Ifc() interface{}     { return x.BndsBnds() }
func (x BndsBndsAcs) BndsBnds() *bnds.Bnds { return x.Arr[x.Idx] }
func (x BndsBndsEach) Act()                { x.BndsBnds() }
func (x BndsBndsEach) Ifc() interface{}    { return x.BndsBnds() }
func (x BndsBndsEach) BndsBnds() *bnds.Bnds {
	vs := x.X.BndsBnds()
	if vs != nil {
		for _, v := range *vs {
			x.Arr[x.Idx] = v // set cur elm to scp
			for _, a := range x.Acts {
				a.Act()
			}
		}
	}
	return vs
}
func (x BndsBndsPllEach) Act()             { x.BndsBnds() }
func (x BndsBndsPllEach) Ifc() interface{} { return x.BndsBnds() }
func (x *BndsBndsPllEachSeg) Act() {
	var actr Actr
	actr.Reset(x.Txt)
	scp := NewScp(x.XprScp, x.ActScpPrnt)
	elmScp := scp.BndBnd(x.Txt[x.Idn.Idx:x.Idn.Lim])
	acts := actr.Acts(scp, x.Xprs...)
	elmScp.Arr[elmScp.Idx] = x.Val // set cur elm to scp
	for _, a := range acts {
		a.Act()
	}
}
func (x BndsBndsPllEach) BndsBnds() *bnds.Bnds {
	vs := x.X.BndsBnds()
	segs := make([]sys.Act, len(*vs))
	for n, v := range *vs {
		segs[n] = &BndsBndsPllEachSeg{
			Val:        v,
			Idn:        x.Idn,
			Txt:        x.Txt,
			ActScpPrnt: x.ActScpPrnt,
			XprScp:     x.XprScp,
			Xprs:       x.Xprs,
		}
	}
	sys.Run().Pll(segs...) // run segs in pll
	return vs
}
func (x TmeRngsAsn) Act()             { x.TmeRngs() }
func (x TmeRngsAsn) Ifc() interface{} { return x.TmeRngs() }
func (x TmeRngsAsn) TmeRngs() *tme.Rngs {
	x.Arr[x.Idx] = x.X.TmeRngs()
	return x.Arr[x.Idx]
}
func (x TmeRngsAcs) Act()               { x.TmeRngs() }
func (x TmeRngsAcs) Ifc() interface{}   { return x.TmeRngs() }
func (x TmeRngsAcs) TmeRngs() *tme.Rngs { return x.Arr[x.Idx] }
func (x TmeRngsEach) Act()              { x.TmeRngs() }
func (x TmeRngsEach) Ifc() interface{}  { return x.TmeRngs() }
func (x TmeRngsEach) TmeRngs() *tme.Rngs {
	vs := x.X.TmeRngs()
	if vs != nil {
		for _, v := range *vs {
			x.Arr[x.Idx] = v // set cur elm to scp
			for _, a := range x.Acts {
				a.Act()
			}
		}
	}
	return vs
}
func (x TmeRngsPllEach) Act()             { x.TmeRngs() }
func (x TmeRngsPllEach) Ifc() interface{} { return x.TmeRngs() }
func (x *TmeRngsPllEachSeg) Act() {
	var actr Actr
	actr.Reset(x.Txt)
	scp := NewScp(x.XprScp, x.ActScpPrnt)
	elmScp := scp.TmeRng(x.Txt[x.Idn.Idx:x.Idn.Lim])
	acts := actr.Acts(scp, x.Xprs...)
	elmScp.Arr[elmScp.Idx] = x.Val // set cur elm to scp
	for _, a := range acts {
		a.Act()
	}
}
func (x TmeRngsPllEach) TmeRngs() *tme.Rngs {
	vs := x.X.TmeRngs()
	segs := make([]sys.Act, len(*vs))
	for n, v := range *vs {
		segs[n] = &TmeRngsPllEachSeg{
			Val:        v,
			Idn:        x.Idn,
			Txt:        x.Txt,
			ActScpPrnt: x.ActScpPrnt,
			XprScp:     x.XprScp,
			Xprs:       x.Xprs,
		}
	}
	sys.Run().Pll(segs...) // run segs in pll
	return vs
}
func (x AnaTrdAsn) Act()             { x.AnaTrd() }
func (x AnaTrdAsn) Ifc() interface{} { return x.AnaTrd() }
func (x AnaTrdAsn) AnaTrd() *ana.Trd {
	x.Arr[x.Idx] = x.X.AnaTrd()
	return x.Arr[x.Idx]
}
func (x AnaTrdAcs) Act()              { x.AnaTrd() }
func (x AnaTrdAcs) Ifc() interface{}  { return x.AnaTrd() }
func (x AnaTrdAcs) AnaTrd() *ana.Trd  { return x.Arr[x.Idx] }
func (x AnaTrdsAsn) Act()             { x.AnaTrds() }
func (x AnaTrdsAsn) Ifc() interface{} { return x.AnaTrds() }
func (x AnaTrdsAsn) AnaTrds() *ana.Trds {
	x.Arr[x.Idx] = x.X.AnaTrds()
	return x.Arr[x.Idx]
}
func (x AnaTrdsAcs) Act()               { x.AnaTrds() }
func (x AnaTrdsAcs) Ifc() interface{}   { return x.AnaTrds() }
func (x AnaTrdsAcs) AnaTrds() *ana.Trds { return x.Arr[x.Idx] }
func (x AnaTrdsEach) Act()              { x.AnaTrds() }
func (x AnaTrdsEach) Ifc() interface{}  { return x.AnaTrds() }
func (x AnaTrdsEach) AnaTrds() *ana.Trds {
	vs := x.X.AnaTrds()
	if vs != nil {
		for _, v := range *vs {
			x.Arr[x.Idx] = v // set cur elm to scp
			for _, a := range x.Acts {
				a.Act()
			}
		}
	}
	return vs
}
func (x AnaTrdsPllEach) Act()             { x.AnaTrds() }
func (x AnaTrdsPllEach) Ifc() interface{} { return x.AnaTrds() }
func (x *AnaTrdsPllEachSeg) Act() {
	var actr Actr
	actr.Reset(x.Txt)
	scp := NewScp(x.XprScp, x.ActScpPrnt)
	elmScp := scp.AnaTrd(x.Txt[x.Idn.Idx:x.Idn.Lim])
	acts := actr.Acts(scp, x.Xprs...)
	elmScp.Arr[elmScp.Idx] = x.Val // set cur elm to scp
	for _, a := range acts {
		a.Act()
	}
}
func (x AnaTrdsPllEach) AnaTrds() *ana.Trds {
	vs := x.X.AnaTrds()
	segs := make([]sys.Act, len(*vs))
	for n, v := range *vs {
		segs[n] = &AnaTrdsPllEachSeg{
			Val:        v,
			Idn:        x.Idn,
			Txt:        x.Txt,
			ActScpPrnt: x.ActScpPrnt,
			XprScp:     x.XprScp,
			Xprs:       x.Xprs,
		}
	}
	sys.Run().Pll(segs...) // run segs in pll
	return vs
}
func (x AnaPrfmAsn) Act()             { x.AnaPrfm() }
func (x AnaPrfmAsn) Ifc() interface{} { return x.AnaPrfm() }
func (x AnaPrfmAsn) AnaPrfm() *ana.Prfm {
	x.Arr[x.Idx] = x.X.AnaPrfm()
	return x.Arr[x.Idx]
}
func (x AnaPrfmAcs) Act()               { x.AnaPrfm() }
func (x AnaPrfmAcs) Ifc() interface{}   { return x.AnaPrfm() }
func (x AnaPrfmAcs) AnaPrfm() *ana.Prfm { return x.Arr[x.Idx] }
func (x AnaPrfmsAsn) Act()              { x.AnaPrfms() }
func (x AnaPrfmsAsn) Ifc() interface{}  { return x.AnaPrfms() }
func (x AnaPrfmsAsn) AnaPrfms() *ana.Prfms {
	x.Arr[x.Idx] = x.X.AnaPrfms()
	return x.Arr[x.Idx]
}
func (x AnaPrfmsAcs) Act()                 { x.AnaPrfms() }
func (x AnaPrfmsAcs) Ifc() interface{}     { return x.AnaPrfms() }
func (x AnaPrfmsAcs) AnaPrfms() *ana.Prfms { return x.Arr[x.Idx] }
func (x AnaPrfmsEach) Act()                { x.AnaPrfms() }
func (x AnaPrfmsEach) Ifc() interface{}    { return x.AnaPrfms() }
func (x AnaPrfmsEach) AnaPrfms() *ana.Prfms {
	vs := x.X.AnaPrfms()
	if vs != nil {
		for _, v := range *vs {
			x.Arr[x.Idx] = v // set cur elm to scp
			for _, a := range x.Acts {
				a.Act()
			}
		}
	}
	return vs
}
func (x AnaPrfmsPllEach) Act()             { x.AnaPrfms() }
func (x AnaPrfmsPllEach) Ifc() interface{} { return x.AnaPrfms() }
func (x *AnaPrfmsPllEachSeg) Act() {
	var actr Actr
	actr.Reset(x.Txt)
	scp := NewScp(x.XprScp, x.ActScpPrnt)
	elmScp := scp.AnaPrfm(x.Txt[x.Idn.Idx:x.Idn.Lim])
	acts := actr.Acts(scp, x.Xprs...)
	elmScp.Arr[elmScp.Idx] = x.Val // set cur elm to scp
	for _, a := range acts {
		a.Act()
	}
}
func (x AnaPrfmsPllEach) AnaPrfms() *ana.Prfms {
	vs := x.X.AnaPrfms()
	segs := make([]sys.Act, len(*vs))
	for n, v := range *vs {
		segs[n] = &AnaPrfmsPllEachSeg{
			Val:        v,
			Idn:        x.Idn,
			Txt:        x.Txt,
			ActScpPrnt: x.ActScpPrnt,
			XprScp:     x.XprScp,
			Xprs:       x.Xprs,
		}
	}
	sys.Run().Pll(segs...) // run segs in pll
	return vs
}
func (x AnaPrfmDltAsn) Act()             { x.AnaPrfmDlt() }
func (x AnaPrfmDltAsn) Ifc() interface{} { return x.AnaPrfmDlt() }
func (x AnaPrfmDltAsn) AnaPrfmDlt() *ana.PrfmDlt {
	x.Arr[x.Idx] = x.X.AnaPrfmDlt()
	return x.Arr[x.Idx]
}
func (x AnaPrfmDltAcs) Act()                     { x.AnaPrfmDlt() }
func (x AnaPrfmDltAcs) Ifc() interface{}         { return x.AnaPrfmDlt() }
func (x AnaPrfmDltAcs) AnaPrfmDlt() *ana.PrfmDlt { return x.Arr[x.Idx] }
func (x AnaPortAsn) Act()                        { x.AnaPort() }
func (x AnaPortAsn) Ifc() interface{}            { return x.AnaPort() }
func (x AnaPortAsn) AnaPort() *ana.Port {
	x.Arr[x.Idx] = x.X.AnaPort()
	return x.Arr[x.Idx]
}
func (x AnaPortAcs) Act()               { x.AnaPort() }
func (x AnaPortAcs) Ifc() interface{}   { return x.AnaPort() }
func (x AnaPortAcs) AnaPort() *ana.Port { return x.Arr[x.Idx] }
func (x HstPrvAsn) Act()                { x.HstPrv() }
func (x HstPrvAsn) Ifc() interface{}    { return x.HstPrv() }
func (x HstPrvAsn) HstPrv() hst.Prv {
	x.Arr[x.Idx] = x.X.HstPrv()
	return x.Arr[x.Idx]
}
func (x HstPrvAcs) Act()               { x.HstPrv() }
func (x HstPrvAcs) Ifc() interface{}   { return x.HstPrv() }
func (x HstPrvAcs) HstPrv() hst.Prv    { return x.Arr[x.Idx] }
func (x HstInstrAsn) Act()             { x.HstInstr() }
func (x HstInstrAsn) Ifc() interface{} { return x.HstInstr() }
func (x HstInstrAsn) HstInstr() hst.Instr {
	x.Arr[x.Idx] = x.X.HstInstr()
	return x.Arr[x.Idx]
}
func (x HstInstrAcs) Act()                { x.HstInstr() }
func (x HstInstrAcs) Ifc() interface{}    { return x.HstInstr() }
func (x HstInstrAcs) HstInstr() hst.Instr { return x.Arr[x.Idx] }
func (x HstInrvlAsn) Act()                { x.HstInrvl() }
func (x HstInrvlAsn) Ifc() interface{}    { return x.HstInrvl() }
func (x HstInrvlAsn) HstInrvl() hst.Inrvl {
	x.Arr[x.Idx] = x.X.HstInrvl()
	return x.Arr[x.Idx]
}
func (x HstInrvlAcs) Act()                { x.HstInrvl() }
func (x HstInrvlAcs) Ifc() interface{}    { return x.HstInrvl() }
func (x HstInrvlAcs) HstInrvl() hst.Inrvl { return x.Arr[x.Idx] }
func (x HstSideAsn) Act()                 { x.HstSide() }
func (x HstSideAsn) Ifc() interface{}     { return x.HstSide() }
func (x HstSideAsn) HstSide() hst.Side {
	x.Arr[x.Idx] = x.X.HstSide()
	return x.Arr[x.Idx]
}
func (x HstSideAcs) Act()              { x.HstSide() }
func (x HstSideAcs) Ifc() interface{}  { return x.HstSide() }
func (x HstSideAcs) HstSide() hst.Side { return x.Arr[x.Idx] }
func (x HstStmAsn) Act()               { x.HstStm() }
func (x HstStmAsn) Ifc() interface{}   { return x.HstStm() }
func (x HstStmAsn) HstStm() hst.Stm {
	x.Arr[x.Idx] = x.X.HstStm()
	return x.Arr[x.Idx]
}
func (x HstStmAcs) Act()             { x.HstStm() }
func (x HstStmAcs) Ifc() interface{} { return x.HstStm() }
func (x HstStmAcs) HstStm() hst.Stm  { return x.Arr[x.Idx] }
func (x HstCndAsn) Act()             { x.HstCnd() }
func (x HstCndAsn) Ifc() interface{} { return x.HstCnd() }
func (x HstCndAsn) HstCnd() hst.Cnd {
	x.Arr[x.Idx] = x.X.HstCnd()
	return x.Arr[x.Idx]
}
func (x HstCndAcs) Act()              { x.HstCnd() }
func (x HstCndAcs) Ifc() interface{}  { return x.HstCnd() }
func (x HstCndAcs) HstCnd() hst.Cnd   { return x.Arr[x.Idx] }
func (x HstStgyAsn) Act()             { x.HstStgy() }
func (x HstStgyAsn) Ifc() interface{} { return x.HstStgy() }
func (x HstStgyAsn) HstStgy() hst.Stgy {
	x.Arr[x.Idx] = x.X.HstStgy()
	return x.Arr[x.Idx]
}
func (x HstStgyAcs) Act()              { x.HstStgy() }
func (x HstStgyAcs) Ifc() interface{}  { return x.HstStgy() }
func (x HstStgyAcs) HstStgy() hst.Stgy { return x.Arr[x.Idx] }
func (x HstPrvsAsn) Act()              { x.HstPrvs() }
func (x HstPrvsAsn) Ifc() interface{}  { return x.HstPrvs() }
func (x HstPrvsAsn) HstPrvs() *hst.Prvs {
	x.Arr[x.Idx] = x.X.HstPrvs()
	return x.Arr[x.Idx]
}
func (x HstPrvsAcs) Act()               { x.HstPrvs() }
func (x HstPrvsAcs) Ifc() interface{}   { return x.HstPrvs() }
func (x HstPrvsAcs) HstPrvs() *hst.Prvs { return x.Arr[x.Idx] }
func (x HstPrvsEach) Act()              { x.HstPrvs() }
func (x HstPrvsEach) Ifc() interface{}  { return x.HstPrvs() }
func (x HstPrvsEach) HstPrvs() *hst.Prvs {
	vs := x.X.HstPrvs()
	if vs != nil {
		for _, v := range *vs {
			x.Arr[x.Idx] = v // set cur elm to scp
			for _, a := range x.Acts {
				a.Act()
			}
		}
	}
	return vs
}
func (x HstPrvsPllEach) Act()             { x.HstPrvs() }
func (x HstPrvsPllEach) Ifc() interface{} { return x.HstPrvs() }
func (x *HstPrvsPllEachSeg) Act() {
	var actr Actr
	actr.Reset(x.Txt)
	scp := NewScp(x.XprScp, x.ActScpPrnt)
	elmScp := scp.HstPrv(x.Txt[x.Idn.Idx:x.Idn.Lim])
	acts := actr.Acts(scp, x.Xprs...)
	elmScp.Arr[elmScp.Idx] = x.Val // set cur elm to scp
	for _, a := range acts {
		a.Act()
	}
}
func (x HstPrvsPllEach) HstPrvs() *hst.Prvs {
	vs := x.X.HstPrvs()
	segs := make([]sys.Act, len(*vs))
	for n, v := range *vs {
		segs[n] = &HstPrvsPllEachSeg{
			Val:        v,
			Idn:        x.Idn,
			Txt:        x.Txt,
			ActScpPrnt: x.ActScpPrnt,
			XprScp:     x.XprScp,
			Xprs:       x.Xprs,
		}
	}
	sys.Run().Pll(segs...) // run segs in pll
	return vs
}
func (x HstInstrsAsn) Act()             { x.HstInstrs() }
func (x HstInstrsAsn) Ifc() interface{} { return x.HstInstrs() }
func (x HstInstrsAsn) HstInstrs() *hst.Instrs {
	x.Arr[x.Idx] = x.X.HstInstrs()
	return x.Arr[x.Idx]
}
func (x HstInstrsAcs) Act()                   { x.HstInstrs() }
func (x HstInstrsAcs) Ifc() interface{}       { return x.HstInstrs() }
func (x HstInstrsAcs) HstInstrs() *hst.Instrs { return x.Arr[x.Idx] }
func (x HstInstrsEach) Act()                  { x.HstInstrs() }
func (x HstInstrsEach) Ifc() interface{}      { return x.HstInstrs() }
func (x HstInstrsEach) HstInstrs() *hst.Instrs {
	vs := x.X.HstInstrs()
	if vs != nil {
		for _, v := range *vs {
			x.Arr[x.Idx] = v // set cur elm to scp
			for _, a := range x.Acts {
				a.Act()
			}
		}
	}
	return vs
}
func (x HstInstrsPllEach) Act()             { x.HstInstrs() }
func (x HstInstrsPllEach) Ifc() interface{} { return x.HstInstrs() }
func (x *HstInstrsPllEachSeg) Act() {
	var actr Actr
	actr.Reset(x.Txt)
	scp := NewScp(x.XprScp, x.ActScpPrnt)
	elmScp := scp.HstInstr(x.Txt[x.Idn.Idx:x.Idn.Lim])
	acts := actr.Acts(scp, x.Xprs...)
	elmScp.Arr[elmScp.Idx] = x.Val // set cur elm to scp
	for _, a := range acts {
		a.Act()
	}
}
func (x HstInstrsPllEach) HstInstrs() *hst.Instrs {
	vs := x.X.HstInstrs()
	segs := make([]sys.Act, len(*vs))
	for n, v := range *vs {
		segs[n] = &HstInstrsPllEachSeg{
			Val:        v,
			Idn:        x.Idn,
			Txt:        x.Txt,
			ActScpPrnt: x.ActScpPrnt,
			XprScp:     x.XprScp,
			Xprs:       x.Xprs,
		}
	}
	sys.Run().Pll(segs...) // run segs in pll
	return vs
}
func (x HstInrvlsAsn) Act()             { x.HstInrvls() }
func (x HstInrvlsAsn) Ifc() interface{} { return x.HstInrvls() }
func (x HstInrvlsAsn) HstInrvls() *hst.Inrvls {
	x.Arr[x.Idx] = x.X.HstInrvls()
	return x.Arr[x.Idx]
}
func (x HstInrvlsAcs) Act()                   { x.HstInrvls() }
func (x HstInrvlsAcs) Ifc() interface{}       { return x.HstInrvls() }
func (x HstInrvlsAcs) HstInrvls() *hst.Inrvls { return x.Arr[x.Idx] }
func (x HstInrvlsEach) Act()                  { x.HstInrvls() }
func (x HstInrvlsEach) Ifc() interface{}      { return x.HstInrvls() }
func (x HstInrvlsEach) HstInrvls() *hst.Inrvls {
	vs := x.X.HstInrvls()
	if vs != nil {
		for _, v := range *vs {
			x.Arr[x.Idx] = v // set cur elm to scp
			for _, a := range x.Acts {
				a.Act()
			}
		}
	}
	return vs
}
func (x HstInrvlsPllEach) Act()             { x.HstInrvls() }
func (x HstInrvlsPllEach) Ifc() interface{} { return x.HstInrvls() }
func (x *HstInrvlsPllEachSeg) Act() {
	var actr Actr
	actr.Reset(x.Txt)
	scp := NewScp(x.XprScp, x.ActScpPrnt)
	elmScp := scp.HstInrvl(x.Txt[x.Idn.Idx:x.Idn.Lim])
	acts := actr.Acts(scp, x.Xprs...)
	elmScp.Arr[elmScp.Idx] = x.Val // set cur elm to scp
	for _, a := range acts {
		a.Act()
	}
}
func (x HstInrvlsPllEach) HstInrvls() *hst.Inrvls {
	vs := x.X.HstInrvls()
	segs := make([]sys.Act, len(*vs))
	for n, v := range *vs {
		segs[n] = &HstInrvlsPllEachSeg{
			Val:        v,
			Idn:        x.Idn,
			Txt:        x.Txt,
			ActScpPrnt: x.ActScpPrnt,
			XprScp:     x.XprScp,
			Xprs:       x.Xprs,
		}
	}
	sys.Run().Pll(segs...) // run segs in pll
	return vs
}
func (x HstSidesAsn) Act()             { x.HstSides() }
func (x HstSidesAsn) Ifc() interface{} { return x.HstSides() }
func (x HstSidesAsn) HstSides() *hst.Sides {
	x.Arr[x.Idx] = x.X.HstSides()
	return x.Arr[x.Idx]
}
func (x HstSidesAcs) Act()                 { x.HstSides() }
func (x HstSidesAcs) Ifc() interface{}     { return x.HstSides() }
func (x HstSidesAcs) HstSides() *hst.Sides { return x.Arr[x.Idx] }
func (x HstSidesEach) Act()                { x.HstSides() }
func (x HstSidesEach) Ifc() interface{}    { return x.HstSides() }
func (x HstSidesEach) HstSides() *hst.Sides {
	vs := x.X.HstSides()
	if vs != nil {
		for _, v := range *vs {
			x.Arr[x.Idx] = v // set cur elm to scp
			for _, a := range x.Acts {
				a.Act()
			}
		}
	}
	return vs
}
func (x HstSidesPllEach) Act()             { x.HstSides() }
func (x HstSidesPllEach) Ifc() interface{} { return x.HstSides() }
func (x *HstSidesPllEachSeg) Act() {
	var actr Actr
	actr.Reset(x.Txt)
	scp := NewScp(x.XprScp, x.ActScpPrnt)
	elmScp := scp.HstSide(x.Txt[x.Idn.Idx:x.Idn.Lim])
	acts := actr.Acts(scp, x.Xprs...)
	elmScp.Arr[elmScp.Idx] = x.Val // set cur elm to scp
	for _, a := range acts {
		a.Act()
	}
}
func (x HstSidesPllEach) HstSides() *hst.Sides {
	vs := x.X.HstSides()
	segs := make([]sys.Act, len(*vs))
	for n, v := range *vs {
		segs[n] = &HstSidesPllEachSeg{
			Val:        v,
			Idn:        x.Idn,
			Txt:        x.Txt,
			ActScpPrnt: x.ActScpPrnt,
			XprScp:     x.XprScp,
			Xprs:       x.Xprs,
		}
	}
	sys.Run().Pll(segs...) // run segs in pll
	return vs
}
func (x HstStmsAsn) Act()             { x.HstStms() }
func (x HstStmsAsn) Ifc() interface{} { return x.HstStms() }
func (x HstStmsAsn) HstStms() *hst.Stms {
	x.Arr[x.Idx] = x.X.HstStms()
	return x.Arr[x.Idx]
}
func (x HstStmsAcs) Act()               { x.HstStms() }
func (x HstStmsAcs) Ifc() interface{}   { return x.HstStms() }
func (x HstStmsAcs) HstStms() *hst.Stms { return x.Arr[x.Idx] }
func (x HstStmsEach) Act()              { x.HstStms() }
func (x HstStmsEach) Ifc() interface{}  { return x.HstStms() }
func (x HstStmsEach) HstStms() *hst.Stms {
	vs := x.X.HstStms()
	if vs != nil {
		for _, v := range *vs {
			x.Arr[x.Idx] = v // set cur elm to scp
			for _, a := range x.Acts {
				a.Act()
			}
		}
	}
	return vs
}
func (x HstStmsPllEach) Act()             { x.HstStms() }
func (x HstStmsPllEach) Ifc() interface{} { return x.HstStms() }
func (x *HstStmsPllEachSeg) Act() {
	var actr Actr
	actr.Reset(x.Txt)
	scp := NewScp(x.XprScp, x.ActScpPrnt)
	elmScp := scp.HstStm(x.Txt[x.Idn.Idx:x.Idn.Lim])
	acts := actr.Acts(scp, x.Xprs...)
	elmScp.Arr[elmScp.Idx] = x.Val // set cur elm to scp
	for _, a := range acts {
		a.Act()
	}
}
func (x HstStmsPllEach) HstStms() *hst.Stms {
	vs := x.X.HstStms()
	segs := make([]sys.Act, len(*vs))
	for n, v := range *vs {
		segs[n] = &HstStmsPllEachSeg{
			Val:        v,
			Idn:        x.Idn,
			Txt:        x.Txt,
			ActScpPrnt: x.ActScpPrnt,
			XprScp:     x.XprScp,
			Xprs:       x.Xprs,
		}
	}
	sys.Run().Pll(segs...) // run segs in pll
	return vs
}
func (x HstCndsAsn) Act()             { x.HstCnds() }
func (x HstCndsAsn) Ifc() interface{} { return x.HstCnds() }
func (x HstCndsAsn) HstCnds() *hst.Cnds {
	x.Arr[x.Idx] = x.X.HstCnds()
	return x.Arr[x.Idx]
}
func (x HstCndsAcs) Act()               { x.HstCnds() }
func (x HstCndsAcs) Ifc() interface{}   { return x.HstCnds() }
func (x HstCndsAcs) HstCnds() *hst.Cnds { return x.Arr[x.Idx] }
func (x HstCndsEach) Act()              { x.HstCnds() }
func (x HstCndsEach) Ifc() interface{}  { return x.HstCnds() }
func (x HstCndsEach) HstCnds() *hst.Cnds {
	vs := x.X.HstCnds()
	if vs != nil {
		for _, v := range *vs {
			x.Arr[x.Idx] = v // set cur elm to scp
			for _, a := range x.Acts {
				a.Act()
			}
		}
	}
	return vs
}
func (x HstCndsPllEach) Act()             { x.HstCnds() }
func (x HstCndsPllEach) Ifc() interface{} { return x.HstCnds() }
func (x *HstCndsPllEachSeg) Act() {
	var actr Actr
	actr.Reset(x.Txt)
	scp := NewScp(x.XprScp, x.ActScpPrnt)
	elmScp := scp.HstCnd(x.Txt[x.Idn.Idx:x.Idn.Lim])
	acts := actr.Acts(scp, x.Xprs...)
	elmScp.Arr[elmScp.Idx] = x.Val // set cur elm to scp
	for _, a := range acts {
		a.Act()
	}
}
func (x HstCndsPllEach) HstCnds() *hst.Cnds {
	vs := x.X.HstCnds()
	segs := make([]sys.Act, len(*vs))
	for n, v := range *vs {
		segs[n] = &HstCndsPllEachSeg{
			Val:        v,
			Idn:        x.Idn,
			Txt:        x.Txt,
			ActScpPrnt: x.ActScpPrnt,
			XprScp:     x.XprScp,
			Xprs:       x.Xprs,
		}
	}
	sys.Run().Pll(segs...) // run segs in pll
	return vs
}
func (x HstStgysAsn) Act()             { x.HstStgys() }
func (x HstStgysAsn) Ifc() interface{} { return x.HstStgys() }
func (x HstStgysAsn) HstStgys() *hst.Stgys {
	x.Arr[x.Idx] = x.X.HstStgys()
	return x.Arr[x.Idx]
}
func (x HstStgysAcs) Act()                 { x.HstStgys() }
func (x HstStgysAcs) Ifc() interface{}     { return x.HstStgys() }
func (x HstStgysAcs) HstStgys() *hst.Stgys { return x.Arr[x.Idx] }
func (x HstStgysEach) Act()                { x.HstStgys() }
func (x HstStgysEach) Ifc() interface{}    { return x.HstStgys() }
func (x HstStgysEach) HstStgys() *hst.Stgys {
	vs := x.X.HstStgys()
	if vs != nil {
		for _, v := range *vs {
			x.Arr[x.Idx] = v // set cur elm to scp
			for _, a := range x.Acts {
				a.Act()
			}
		}
	}
	return vs
}
func (x HstStgysPllEach) Act()             { x.HstStgys() }
func (x HstStgysPllEach) Ifc() interface{} { return x.HstStgys() }
func (x *HstStgysPllEachSeg) Act() {
	var actr Actr
	actr.Reset(x.Txt)
	scp := NewScp(x.XprScp, x.ActScpPrnt)
	elmScp := scp.HstStgy(x.Txt[x.Idn.Idx:x.Idn.Lim])
	acts := actr.Acts(scp, x.Xprs...)
	elmScp.Arr[elmScp.Idx] = x.Val // set cur elm to scp
	for _, a := range acts {
		a.Act()
	}
}
func (x HstStgysPllEach) HstStgys() *hst.Stgys {
	vs := x.X.HstStgys()
	segs := make([]sys.Act, len(*vs))
	for n, v := range *vs {
		segs[n] = &HstStgysPllEachSeg{
			Val:        v,
			Idn:        x.Idn,
			Txt:        x.Txt,
			ActScpPrnt: x.ActScpPrnt,
			XprScp:     x.XprScp,
			Xprs:       x.Xprs,
		}
	}
	sys.Run().Pll(segs...) // run segs in pll
	return vs
}
func (x RltPrvAsn) Act()             { x.RltPrv() }
func (x RltPrvAsn) Ifc() interface{} { return x.RltPrv() }
func (x RltPrvAsn) RltPrv() rlt.Prv {
	x.Arr[x.Idx] = x.X.RltPrv()
	return x.Arr[x.Idx]
}
func (x RltPrvAcs) Act()               { x.RltPrv() }
func (x RltPrvAcs) Ifc() interface{}   { return x.RltPrv() }
func (x RltPrvAcs) RltPrv() rlt.Prv    { return x.Arr[x.Idx] }
func (x RltInstrAsn) Act()             { x.RltInstr() }
func (x RltInstrAsn) Ifc() interface{} { return x.RltInstr() }
func (x RltInstrAsn) RltInstr() rlt.Instr {
	x.Arr[x.Idx] = x.X.RltInstr()
	return x.Arr[x.Idx]
}
func (x RltInstrAcs) Act()                { x.RltInstr() }
func (x RltInstrAcs) Ifc() interface{}    { return x.RltInstr() }
func (x RltInstrAcs) RltInstr() rlt.Instr { return x.Arr[x.Idx] }
func (x RltInrvlAsn) Act()                { x.RltInrvl() }
func (x RltInrvlAsn) Ifc() interface{}    { return x.RltInrvl() }
func (x RltInrvlAsn) RltInrvl() rlt.Inrvl {
	x.Arr[x.Idx] = x.X.RltInrvl()
	return x.Arr[x.Idx]
}
func (x RltInrvlAcs) Act()                { x.RltInrvl() }
func (x RltInrvlAcs) Ifc() interface{}    { return x.RltInrvl() }
func (x RltInrvlAcs) RltInrvl() rlt.Inrvl { return x.Arr[x.Idx] }
func (x RltSideAsn) Act()                 { x.RltSide() }
func (x RltSideAsn) Ifc() interface{}     { return x.RltSide() }
func (x RltSideAsn) RltSide() rlt.Side {
	x.Arr[x.Idx] = x.X.RltSide()
	return x.Arr[x.Idx]
}
func (x RltSideAcs) Act()              { x.RltSide() }
func (x RltSideAcs) Ifc() interface{}  { return x.RltSide() }
func (x RltSideAcs) RltSide() rlt.Side { return x.Arr[x.Idx] }
func (x RltStmAsn) Act()               { x.RltStm() }
func (x RltStmAsn) Ifc() interface{}   { return x.RltStm() }
func (x RltStmAsn) RltStm() rlt.Stm {
	x.Arr[x.Idx] = x.X.RltStm()
	return x.Arr[x.Idx]
}
func (x RltStmAcs) Act()             { x.RltStm() }
func (x RltStmAcs) Ifc() interface{} { return x.RltStm() }
func (x RltStmAcs) RltStm() rlt.Stm  { return x.Arr[x.Idx] }
func (x RltCndAsn) Act()             { x.RltCnd() }
func (x RltCndAsn) Ifc() interface{} { return x.RltCnd() }
func (x RltCndAsn) RltCnd() rlt.Cnd {
	x.Arr[x.Idx] = x.X.RltCnd()
	return x.Arr[x.Idx]
}
func (x RltCndAcs) Act()              { x.RltCnd() }
func (x RltCndAcs) Ifc() interface{}  { return x.RltCnd() }
func (x RltCndAcs) RltCnd() rlt.Cnd   { return x.Arr[x.Idx] }
func (x RltStgyAsn) Act()             { x.RltStgy() }
func (x RltStgyAsn) Ifc() interface{} { return x.RltStgy() }
func (x RltStgyAsn) RltStgy() rlt.Stgy {
	x.Arr[x.Idx] = x.X.RltStgy()
	return x.Arr[x.Idx]
}
func (x RltStgyAcs) Act()              { x.RltStgy() }
func (x RltStgyAcs) Ifc() interface{}  { return x.RltStgy() }
func (x RltStgyAcs) RltStgy() rlt.Stgy { return x.Arr[x.Idx] }
func (x RltPrvsAsn) Act()              { x.RltPrvs() }
func (x RltPrvsAsn) Ifc() interface{}  { return x.RltPrvs() }
func (x RltPrvsAsn) RltPrvs() *rlt.Prvs {
	x.Arr[x.Idx] = x.X.RltPrvs()
	return x.Arr[x.Idx]
}
func (x RltPrvsAcs) Act()               { x.RltPrvs() }
func (x RltPrvsAcs) Ifc() interface{}   { return x.RltPrvs() }
func (x RltPrvsAcs) RltPrvs() *rlt.Prvs { return x.Arr[x.Idx] }
func (x RltPrvsEach) Act()              { x.RltPrvs() }
func (x RltPrvsEach) Ifc() interface{}  { return x.RltPrvs() }
func (x RltPrvsEach) RltPrvs() *rlt.Prvs {
	vs := x.X.RltPrvs()
	if vs != nil {
		for _, v := range *vs {
			x.Arr[x.Idx] = v // set cur elm to scp
			for _, a := range x.Acts {
				a.Act()
			}
		}
	}
	return vs
}
func (x RltPrvsPllEach) Act()             { x.RltPrvs() }
func (x RltPrvsPllEach) Ifc() interface{} { return x.RltPrvs() }
func (x *RltPrvsPllEachSeg) Act() {
	var actr Actr
	actr.Reset(x.Txt)
	scp := NewScp(x.XprScp, x.ActScpPrnt)
	elmScp := scp.RltPrv(x.Txt[x.Idn.Idx:x.Idn.Lim])
	acts := actr.Acts(scp, x.Xprs...)
	elmScp.Arr[elmScp.Idx] = x.Val // set cur elm to scp
	for _, a := range acts {
		a.Act()
	}
}
func (x RltPrvsPllEach) RltPrvs() *rlt.Prvs {
	vs := x.X.RltPrvs()
	segs := make([]sys.Act, len(*vs))
	for n, v := range *vs {
		segs[n] = &RltPrvsPllEachSeg{
			Val:        v,
			Idn:        x.Idn,
			Txt:        x.Txt,
			ActScpPrnt: x.ActScpPrnt,
			XprScp:     x.XprScp,
			Xprs:       x.Xprs,
		}
	}
	sys.Run().Pll(segs...) // run segs in pll
	return vs
}
func (x RltInstrsAsn) Act()             { x.RltInstrs() }
func (x RltInstrsAsn) Ifc() interface{} { return x.RltInstrs() }
func (x RltInstrsAsn) RltInstrs() *rlt.Instrs {
	x.Arr[x.Idx] = x.X.RltInstrs()
	return x.Arr[x.Idx]
}
func (x RltInstrsAcs) Act()                   { x.RltInstrs() }
func (x RltInstrsAcs) Ifc() interface{}       { return x.RltInstrs() }
func (x RltInstrsAcs) RltInstrs() *rlt.Instrs { return x.Arr[x.Idx] }
func (x RltInstrsEach) Act()                  { x.RltInstrs() }
func (x RltInstrsEach) Ifc() interface{}      { return x.RltInstrs() }
func (x RltInstrsEach) RltInstrs() *rlt.Instrs {
	vs := x.X.RltInstrs()
	if vs != nil {
		for _, v := range *vs {
			x.Arr[x.Idx] = v // set cur elm to scp
			for _, a := range x.Acts {
				a.Act()
			}
		}
	}
	return vs
}
func (x RltInstrsPllEach) Act()             { x.RltInstrs() }
func (x RltInstrsPllEach) Ifc() interface{} { return x.RltInstrs() }
func (x *RltInstrsPllEachSeg) Act() {
	var actr Actr
	actr.Reset(x.Txt)
	scp := NewScp(x.XprScp, x.ActScpPrnt)
	elmScp := scp.RltInstr(x.Txt[x.Idn.Idx:x.Idn.Lim])
	acts := actr.Acts(scp, x.Xprs...)
	elmScp.Arr[elmScp.Idx] = x.Val // set cur elm to scp
	for _, a := range acts {
		a.Act()
	}
}
func (x RltInstrsPllEach) RltInstrs() *rlt.Instrs {
	vs := x.X.RltInstrs()
	segs := make([]sys.Act, len(*vs))
	for n, v := range *vs {
		segs[n] = &RltInstrsPllEachSeg{
			Val:        v,
			Idn:        x.Idn,
			Txt:        x.Txt,
			ActScpPrnt: x.ActScpPrnt,
			XprScp:     x.XprScp,
			Xprs:       x.Xprs,
		}
	}
	sys.Run().Pll(segs...) // run segs in pll
	return vs
}
func (x RltInrvlsAsn) Act()             { x.RltInrvls() }
func (x RltInrvlsAsn) Ifc() interface{} { return x.RltInrvls() }
func (x RltInrvlsAsn) RltInrvls() *rlt.Inrvls {
	x.Arr[x.Idx] = x.X.RltInrvls()
	return x.Arr[x.Idx]
}
func (x RltInrvlsAcs) Act()                   { x.RltInrvls() }
func (x RltInrvlsAcs) Ifc() interface{}       { return x.RltInrvls() }
func (x RltInrvlsAcs) RltInrvls() *rlt.Inrvls { return x.Arr[x.Idx] }
func (x RltInrvlsEach) Act()                  { x.RltInrvls() }
func (x RltInrvlsEach) Ifc() interface{}      { return x.RltInrvls() }
func (x RltInrvlsEach) RltInrvls() *rlt.Inrvls {
	vs := x.X.RltInrvls()
	if vs != nil {
		for _, v := range *vs {
			x.Arr[x.Idx] = v // set cur elm to scp
			for _, a := range x.Acts {
				a.Act()
			}
		}
	}
	return vs
}
func (x RltInrvlsPllEach) Act()             { x.RltInrvls() }
func (x RltInrvlsPllEach) Ifc() interface{} { return x.RltInrvls() }
func (x *RltInrvlsPllEachSeg) Act() {
	var actr Actr
	actr.Reset(x.Txt)
	scp := NewScp(x.XprScp, x.ActScpPrnt)
	elmScp := scp.RltInrvl(x.Txt[x.Idn.Idx:x.Idn.Lim])
	acts := actr.Acts(scp, x.Xprs...)
	elmScp.Arr[elmScp.Idx] = x.Val // set cur elm to scp
	for _, a := range acts {
		a.Act()
	}
}
func (x RltInrvlsPllEach) RltInrvls() *rlt.Inrvls {
	vs := x.X.RltInrvls()
	segs := make([]sys.Act, len(*vs))
	for n, v := range *vs {
		segs[n] = &RltInrvlsPllEachSeg{
			Val:        v,
			Idn:        x.Idn,
			Txt:        x.Txt,
			ActScpPrnt: x.ActScpPrnt,
			XprScp:     x.XprScp,
			Xprs:       x.Xprs,
		}
	}
	sys.Run().Pll(segs...) // run segs in pll
	return vs
}
func (x RltSidesAsn) Act()             { x.RltSides() }
func (x RltSidesAsn) Ifc() interface{} { return x.RltSides() }
func (x RltSidesAsn) RltSides() *rlt.Sides {
	x.Arr[x.Idx] = x.X.RltSides()
	return x.Arr[x.Idx]
}
func (x RltSidesAcs) Act()                 { x.RltSides() }
func (x RltSidesAcs) Ifc() interface{}     { return x.RltSides() }
func (x RltSidesAcs) RltSides() *rlt.Sides { return x.Arr[x.Idx] }
func (x RltSidesEach) Act()                { x.RltSides() }
func (x RltSidesEach) Ifc() interface{}    { return x.RltSides() }
func (x RltSidesEach) RltSides() *rlt.Sides {
	vs := x.X.RltSides()
	if vs != nil {
		for _, v := range *vs {
			x.Arr[x.Idx] = v // set cur elm to scp
			for _, a := range x.Acts {
				a.Act()
			}
		}
	}
	return vs
}
func (x RltSidesPllEach) Act()             { x.RltSides() }
func (x RltSidesPllEach) Ifc() interface{} { return x.RltSides() }
func (x *RltSidesPllEachSeg) Act() {
	var actr Actr
	actr.Reset(x.Txt)
	scp := NewScp(x.XprScp, x.ActScpPrnt)
	elmScp := scp.RltSide(x.Txt[x.Idn.Idx:x.Idn.Lim])
	acts := actr.Acts(scp, x.Xprs...)
	elmScp.Arr[elmScp.Idx] = x.Val // set cur elm to scp
	for _, a := range acts {
		a.Act()
	}
}
func (x RltSidesPllEach) RltSides() *rlt.Sides {
	vs := x.X.RltSides()
	segs := make([]sys.Act, len(*vs))
	for n, v := range *vs {
		segs[n] = &RltSidesPllEachSeg{
			Val:        v,
			Idn:        x.Idn,
			Txt:        x.Txt,
			ActScpPrnt: x.ActScpPrnt,
			XprScp:     x.XprScp,
			Xprs:       x.Xprs,
		}
	}
	sys.Run().Pll(segs...) // run segs in pll
	return vs
}
func (x RltStmsAsn) Act()             { x.RltStms() }
func (x RltStmsAsn) Ifc() interface{} { return x.RltStms() }
func (x RltStmsAsn) RltStms() *rlt.Stms {
	x.Arr[x.Idx] = x.X.RltStms()
	return x.Arr[x.Idx]
}
func (x RltStmsAcs) Act()               { x.RltStms() }
func (x RltStmsAcs) Ifc() interface{}   { return x.RltStms() }
func (x RltStmsAcs) RltStms() *rlt.Stms { return x.Arr[x.Idx] }
func (x RltStmsEach) Act()              { x.RltStms() }
func (x RltStmsEach) Ifc() interface{}  { return x.RltStms() }
func (x RltStmsEach) RltStms() *rlt.Stms {
	vs := x.X.RltStms()
	if vs != nil {
		for _, v := range *vs {
			x.Arr[x.Idx] = v // set cur elm to scp
			for _, a := range x.Acts {
				a.Act()
			}
		}
	}
	return vs
}
func (x RltStmsPllEach) Act()             { x.RltStms() }
func (x RltStmsPllEach) Ifc() interface{} { return x.RltStms() }
func (x *RltStmsPllEachSeg) Act() {
	var actr Actr
	actr.Reset(x.Txt)
	scp := NewScp(x.XprScp, x.ActScpPrnt)
	elmScp := scp.RltStm(x.Txt[x.Idn.Idx:x.Idn.Lim])
	acts := actr.Acts(scp, x.Xprs...)
	elmScp.Arr[elmScp.Idx] = x.Val // set cur elm to scp
	for _, a := range acts {
		a.Act()
	}
}
func (x RltStmsPllEach) RltStms() *rlt.Stms {
	vs := x.X.RltStms()
	segs := make([]sys.Act, len(*vs))
	for n, v := range *vs {
		segs[n] = &RltStmsPllEachSeg{
			Val:        v,
			Idn:        x.Idn,
			Txt:        x.Txt,
			ActScpPrnt: x.ActScpPrnt,
			XprScp:     x.XprScp,
			Xprs:       x.Xprs,
		}
	}
	sys.Run().Pll(segs...) // run segs in pll
	return vs
}
func (x RltCndsAsn) Act()             { x.RltCnds() }
func (x RltCndsAsn) Ifc() interface{} { return x.RltCnds() }
func (x RltCndsAsn) RltCnds() *rlt.Cnds {
	x.Arr[x.Idx] = x.X.RltCnds()
	return x.Arr[x.Idx]
}
func (x RltCndsAcs) Act()               { x.RltCnds() }
func (x RltCndsAcs) Ifc() interface{}   { return x.RltCnds() }
func (x RltCndsAcs) RltCnds() *rlt.Cnds { return x.Arr[x.Idx] }
func (x RltCndsEach) Act()              { x.RltCnds() }
func (x RltCndsEach) Ifc() interface{}  { return x.RltCnds() }
func (x RltCndsEach) RltCnds() *rlt.Cnds {
	vs := x.X.RltCnds()
	if vs != nil {
		for _, v := range *vs {
			x.Arr[x.Idx] = v // set cur elm to scp
			for _, a := range x.Acts {
				a.Act()
			}
		}
	}
	return vs
}
func (x RltCndsPllEach) Act()             { x.RltCnds() }
func (x RltCndsPllEach) Ifc() interface{} { return x.RltCnds() }
func (x *RltCndsPllEachSeg) Act() {
	var actr Actr
	actr.Reset(x.Txt)
	scp := NewScp(x.XprScp, x.ActScpPrnt)
	elmScp := scp.RltCnd(x.Txt[x.Idn.Idx:x.Idn.Lim])
	acts := actr.Acts(scp, x.Xprs...)
	elmScp.Arr[elmScp.Idx] = x.Val // set cur elm to scp
	for _, a := range acts {
		a.Act()
	}
}
func (x RltCndsPllEach) RltCnds() *rlt.Cnds {
	vs := x.X.RltCnds()
	segs := make([]sys.Act, len(*vs))
	for n, v := range *vs {
		segs[n] = &RltCndsPllEachSeg{
			Val:        v,
			Idn:        x.Idn,
			Txt:        x.Txt,
			ActScpPrnt: x.ActScpPrnt,
			XprScp:     x.XprScp,
			Xprs:       x.Xprs,
		}
	}
	sys.Run().Pll(segs...) // run segs in pll
	return vs
}
func (x RltStgysAsn) Act()             { x.RltStgys() }
func (x RltStgysAsn) Ifc() interface{} { return x.RltStgys() }
func (x RltStgysAsn) RltStgys() *rlt.Stgys {
	x.Arr[x.Idx] = x.X.RltStgys()
	return x.Arr[x.Idx]
}
func (x RltStgysAcs) Act()                 { x.RltStgys() }
func (x RltStgysAcs) Ifc() interface{}     { return x.RltStgys() }
func (x RltStgysAcs) RltStgys() *rlt.Stgys { return x.Arr[x.Idx] }
func (x RltStgysEach) Act()                { x.RltStgys() }
func (x RltStgysEach) Ifc() interface{}    { return x.RltStgys() }
func (x RltStgysEach) RltStgys() *rlt.Stgys {
	vs := x.X.RltStgys()
	if vs != nil {
		for _, v := range *vs {
			x.Arr[x.Idx] = v // set cur elm to scp
			for _, a := range x.Acts {
				a.Act()
			}
		}
	}
	return vs
}
func (x RltStgysPllEach) Act()             { x.RltStgys() }
func (x RltStgysPllEach) Ifc() interface{} { return x.RltStgys() }
func (x *RltStgysPllEachSeg) Act() {
	var actr Actr
	actr.Reset(x.Txt)
	scp := NewScp(x.XprScp, x.ActScpPrnt)
	elmScp := scp.RltStgy(x.Txt[x.Idn.Idx:x.Idn.Lim])
	acts := actr.Acts(scp, x.Xprs...)
	elmScp.Arr[elmScp.Idx] = x.Val // set cur elm to scp
	for _, a := range acts {
		a.Act()
	}
}
func (x RltStgysPllEach) RltStgys() *rlt.Stgys {
	vs := x.X.RltStgys()
	segs := make([]sys.Act, len(*vs))
	for n, v := range *vs {
		segs[n] = &RltStgysPllEachSeg{
			Val:        v,
			Idn:        x.Idn,
			Txt:        x.Txt,
			ActScpPrnt: x.ActScpPrnt,
			XprScp:     x.XprScp,
			Xprs:       x.Xprs,
		}
	}
	sys.Run().Pll(segs...) // run segs in pll
	return vs
}
func (x FntFntAsn) Act()             { x.FntFnt() }
func (x FntFntAsn) Ifc() interface{} { return x.FntFnt() }
func (x FntFntAsn) FntFnt() *fnt.Fnt {
	x.Arr[x.Idx] = x.X.FntFnt()
	return x.Arr[x.Idx]
}
func (x FntFntAcs) Act()             { x.FntFnt() }
func (x FntFntAcs) Ifc() interface{} { return x.FntFnt() }
func (x FntFntAcs) FntFnt() *fnt.Fnt { return x.Arr[x.Idx] }
func (x ClrClrAsn) Act()             { x.ClrClr() }
func (x ClrClrAsn) Ifc() interface{} { return x.ClrClr() }
func (x ClrClrAsn) ClrClr() clr.Clr {
	x.Arr[x.Idx] = x.X.ClrClr()
	return x.Arr[x.Idx]
}
func (x ClrClrAcs) Act()             { x.ClrClr() }
func (x ClrClrAcs) Ifc() interface{} { return x.ClrClr() }
func (x ClrClrAcs) ClrClr() clr.Clr  { return x.Arr[x.Idx] }
func (x PenPenAsn) Act()             { x.PenPen() }
func (x PenPenAsn) Ifc() interface{} { return x.PenPen() }
func (x PenPenAsn) PenPen() pen.Pen {
	x.Arr[x.Idx] = x.X.PenPen()
	return x.Arr[x.Idx]
}
func (x PenPenAcs) Act()              { x.PenPen() }
func (x PenPenAcs) Ifc() interface{}  { return x.PenPen() }
func (x PenPenAcs) PenPen() pen.Pen   { return x.Arr[x.Idx] }
func (x PenPensAsn) Act()             { x.PenPens() }
func (x PenPensAsn) Ifc() interface{} { return x.PenPens() }
func (x PenPensAsn) PenPens() *pen.Pens {
	x.Arr[x.Idx] = x.X.PenPens()
	return x.Arr[x.Idx]
}
func (x PenPensAcs) Act()               { x.PenPens() }
func (x PenPensAcs) Ifc() interface{}   { return x.PenPens() }
func (x PenPensAcs) PenPens() *pen.Pens { return x.Arr[x.Idx] }
func (x PenPensEach) Act()              { x.PenPens() }
func (x PenPensEach) Ifc() interface{}  { return x.PenPens() }
func (x PenPensEach) PenPens() *pen.Pens {
	vs := x.X.PenPens()
	if vs != nil {
		for _, v := range *vs {
			x.Arr[x.Idx] = v // set cur elm to scp
			for _, a := range x.Acts {
				a.Act()
			}
		}
	}
	return vs
}
func (x PenPensPllEach) Act()             { x.PenPens() }
func (x PenPensPllEach) Ifc() interface{} { return x.PenPens() }
func (x *PenPensPllEachSeg) Act() {
	var actr Actr
	actr.Reset(x.Txt)
	scp := NewScp(x.XprScp, x.ActScpPrnt)
	elmScp := scp.PenPen(x.Txt[x.Idn.Idx:x.Idn.Lim])
	acts := actr.Acts(scp, x.Xprs...)
	elmScp.Arr[elmScp.Idx] = x.Val // set cur elm to scp
	for _, a := range acts {
		a.Act()
	}
}
func (x PenPensPllEach) PenPens() *pen.Pens {
	vs := x.X.PenPens()
	segs := make([]sys.Act, len(*vs))
	for n, v := range *vs {
		segs[n] = &PenPensPllEachSeg{
			Val:        v,
			Idn:        x.Idn,
			Txt:        x.Txt,
			ActScpPrnt: x.ActScpPrnt,
			XprScp:     x.XprScp,
			Xprs:       x.Xprs,
		}
	}
	sys.Run().Pll(segs...) // run segs in pll
	return vs
}
func (x PltPltAsn) Act()             { x.PltPlt() }
func (x PltPltAsn) Ifc() interface{} { return x.PltPlt() }
func (x PltPltAsn) PltPlt() plt.Plt {
	x.Arr[x.Idx] = x.X.PltPlt()
	return x.Arr[x.Idx]
}
func (x PltPltAcs) Act()              { x.PltPlt() }
func (x PltPltAcs) Ifc() interface{}  { return x.PltPlt() }
func (x PltPltAcs) PltPlt() plt.Plt   { return x.Arr[x.Idx] }
func (x PltPltsAsn) Act()             { x.PltPlts() }
func (x PltPltsAsn) Ifc() interface{} { return x.PltPlts() }
func (x PltPltsAsn) PltPlts() *plt.Plts {
	x.Arr[x.Idx] = x.X.PltPlts()
	return x.Arr[x.Idx]
}
func (x PltPltsAcs) Act()               { x.PltPlts() }
func (x PltPltsAcs) Ifc() interface{}   { return x.PltPlts() }
func (x PltPltsAcs) PltPlts() *plt.Plts { return x.Arr[x.Idx] }
func (x PltPltsEach) Act()              { x.PltPlts() }
func (x PltPltsEach) Ifc() interface{}  { return x.PltPlts() }
func (x PltPltsEach) PltPlts() *plt.Plts {
	vs := x.X.PltPlts()
	if vs != nil {
		for _, v := range *vs {
			x.Arr[x.Idx] = v // set cur elm to scp
			for _, a := range x.Acts {
				a.Act()
			}
		}
	}
	return vs
}
func (x PltPltsPllEach) Act()             { x.PltPlts() }
func (x PltPltsPllEach) Ifc() interface{} { return x.PltPlts() }
func (x *PltPltsPllEachSeg) Act() {
	var actr Actr
	actr.Reset(x.Txt)
	scp := NewScp(x.XprScp, x.ActScpPrnt)
	elmScp := scp.PltPlt(x.Txt[x.Idn.Idx:x.Idn.Lim])
	acts := actr.Acts(scp, x.Xprs...)
	elmScp.Arr[elmScp.Idx] = x.Val // set cur elm to scp
	for _, a := range acts {
		a.Act()
	}
}
func (x PltPltsPllEach) PltPlts() *plt.Plts {
	vs := x.X.PltPlts()
	segs := make([]sys.Act, len(*vs))
	for n, v := range *vs {
		segs[n] = &PltPltsPllEachSeg{
			Val:        v,
			Idn:        x.Idn,
			Txt:        x.Txt,
			ActScpPrnt: x.ActScpPrnt,
			XprScp:     x.XprScp,
			Xprs:       x.Xprs,
		}
	}
	sys.Run().Pll(segs...) // run segs in pll
	return vs
}
func (x PltStmAsn) Act()             { x.PltStm() }
func (x PltStmAsn) Ifc() interface{} { return x.PltStm() }
func (x PltStmAsn) PltPlt() plt.Plt  { return x.PltStm() }
func (x PltStmAsn) PltStm() *plt.Stm {
	x.Arr[x.Idx] = x.X.PltStm()
	return x.Arr[x.Idx]
}
func (x PltStmAcs) Act()                  { x.PltStm() }
func (x PltStmAcs) Ifc() interface{}      { return x.PltStm() }
func (x PltStmAcs) PltPlt() plt.Plt       { return x.PltStm() }
func (x PltStmAcs) PltStm() *plt.Stm      { return x.Arr[x.Idx] }
func (x PltFltsSctrAsn) Act()             { x.PltFltsSctr() }
func (x PltFltsSctrAsn) Ifc() interface{} { return x.PltFltsSctr() }
func (x PltFltsSctrAsn) PltPlt() plt.Plt  { return x.PltFltsSctr() }
func (x PltFltsSctrAsn) PltFltsSctr() *plt.FltsSctr {
	x.Arr[x.Idx] = x.X.PltFltsSctr()
	return x.Arr[x.Idx]
}
func (x PltFltsSctrAcs) Act()                       { x.PltFltsSctr() }
func (x PltFltsSctrAcs) Ifc() interface{}           { return x.PltFltsSctr() }
func (x PltFltsSctrAcs) PltPlt() plt.Plt            { return x.PltFltsSctr() }
func (x PltFltsSctrAcs) PltFltsSctr() *plt.FltsSctr { return x.Arr[x.Idx] }
func (x PltFltsSctrDistAsn) Act()                   { x.PltFltsSctrDist() }
func (x PltFltsSctrDistAsn) Ifc() interface{}       { return x.PltFltsSctrDist() }
func (x PltFltsSctrDistAsn) PltPlt() plt.Plt        { return x.PltFltsSctrDist() }
func (x PltFltsSctrDistAsn) PltFltsSctrDist() *plt.FltsSctrDist {
	x.Arr[x.Idx] = x.X.PltFltsSctrDist()
	return x.Arr[x.Idx]
}
func (x PltFltsSctrDistAcs) Act()                               { x.PltFltsSctrDist() }
func (x PltFltsSctrDistAcs) Ifc() interface{}                   { return x.PltFltsSctrDist() }
func (x PltFltsSctrDistAcs) PltPlt() plt.Plt                    { return x.PltFltsSctrDist() }
func (x PltFltsSctrDistAcs) PltFltsSctrDist() *plt.FltsSctrDist { return x.Arr[x.Idx] }
func (x PltHrzAsn) Act()                                        { x.PltHrz() }
func (x PltHrzAsn) Ifc() interface{}                            { return x.PltHrz() }
func (x PltHrzAsn) PltPlt() plt.Plt                             { return x.PltHrz() }
func (x PltHrzAsn) PltHrz() *plt.Hrz {
	x.Arr[x.Idx] = x.X.PltHrz()
	return x.Arr[x.Idx]
}
func (x PltHrzAcs) Act()             { x.PltHrz() }
func (x PltHrzAcs) Ifc() interface{} { return x.PltHrz() }
func (x PltHrzAcs) PltPlt() plt.Plt  { return x.PltHrz() }
func (x PltHrzAcs) PltHrz() *plt.Hrz { return x.Arr[x.Idx] }
func (x PltVrtAsn) Act()             { x.PltVrt() }
func (x PltVrtAsn) Ifc() interface{} { return x.PltVrt() }
func (x PltVrtAsn) PltPlt() plt.Plt  { return x.PltVrt() }
func (x PltVrtAsn) PltVrt() *plt.Vrt {
	x.Arr[x.Idx] = x.X.PltVrt()
	return x.Arr[x.Idx]
}
func (x PltVrtAcs) Act()              { x.PltVrt() }
func (x PltVrtAcs) Ifc() interface{}  { return x.PltVrt() }
func (x PltVrtAcs) PltPlt() plt.Plt   { return x.PltVrt() }
func (x PltVrtAcs) PltVrt() *plt.Vrt  { return x.Arr[x.Idx] }
func (x PltDpthAsn) Act()             { x.PltDpth() }
func (x PltDpthAsn) Ifc() interface{} { return x.PltDpth() }
func (x PltDpthAsn) PltPlt() plt.Plt  { return x.PltDpth() }
func (x PltDpthAsn) PltDpth() *plt.Dpth {
	x.Arr[x.Idx] = x.X.PltDpth()
	return x.Arr[x.Idx]
}
func (x PltDpthAcs) Act()               { x.PltDpth() }
func (x PltDpthAcs) Ifc() interface{}   { return x.PltDpth() }
func (x PltDpthAcs) PltPlt() plt.Plt    { return x.PltDpth() }
func (x PltDpthAcs) PltDpth() *plt.Dpth { return x.Arr[x.Idx] }
func (x SysMuAsn) Act()                 { x.SysMu() }
func (x SysMuAsn) Ifc() interface{}     { return x.SysMu() }
func (x SysMuAsn) SysMu() *sys.Mu {
	x.Arr[x.Idx] = x.X.SysMu()
	return x.Arr[x.Idx]
}
func (x SysMuAcs) Act()               { x.SysMu() }
func (x SysMuAcs) Ifc() interface{}   { return x.SysMu() }
func (x SysMuAcs) SysMu() *sys.Mu     { return x.Arr[x.Idx] }
func (x BolBolThen) Act()             { x.BolBol() }
func (x BolBolThen) Ifc() interface{} { return x.BolBol() }
func (x BolBolThen) BolBol() bol.Bol {
	v := x.X.BolBol()
	if v {
		for _, a := range x.Acts {
			a.Act()
		}
	}
	return v
}
func (x BolBolElse) Act()             { x.BolBol() }
func (x BolBolElse) Ifc() interface{} { return x.BolBol() }
func (x BolBolElse) BolBol() bol.Bol {
	v := x.X.BolBol()
	if !v {
		for _, a := range x.Acts {
			a.Act()
		}
	}
	return v
}
func (x AnaPrfmPnlPctGet) Act()                      { x.FltFlt() }
func (x AnaPrfmPnlPctGet) Ifc() interface{}          { return x.FltFlt() }
func (x AnaPrfmPnlPctGet) FltFlt() flt.Flt           { return x.X.AnaPrfm().PnlPct }
func (x AnaPrfmScsPctGet) Act()                      { x.FltFlt() }
func (x AnaPrfmScsPctGet) Ifc() interface{}          { return x.FltFlt() }
func (x AnaPrfmScsPctGet) FltFlt() flt.Flt           { return x.X.AnaPrfm().ScsPct }
func (x AnaPrfmPipPerDayGet) Act()                   { x.FltFlt() }
func (x AnaPrfmPipPerDayGet) Ifc() interface{}       { return x.FltFlt() }
func (x AnaPrfmPipPerDayGet) FltFlt() flt.Flt        { return x.X.AnaPrfm().PipPerDay }
func (x AnaPrfmUsdPerDayGet) Act()                   { x.FltFlt() }
func (x AnaPrfmUsdPerDayGet) Ifc() interface{}       { return x.FltFlt() }
func (x AnaPrfmUsdPerDayGet) FltFlt() flt.Flt        { return x.X.AnaPrfm().UsdPerDay }
func (x AnaPrfmScsPerDayGet) Act()                   { x.FltFlt() }
func (x AnaPrfmScsPerDayGet) Ifc() interface{}       { return x.FltFlt() }
func (x AnaPrfmScsPerDayGet) FltFlt() flt.Flt        { return x.X.AnaPrfm().ScsPerDay }
func (x AnaPrfmOpnPerDayGet) Act()                   { x.FltFlt() }
func (x AnaPrfmOpnPerDayGet) Ifc() interface{}       { return x.FltFlt() }
func (x AnaPrfmOpnPerDayGet) FltFlt() flt.Flt        { return x.X.AnaPrfm().OpnPerDay }
func (x AnaPrfmPnlUsdGet) Act()                      { x.FltFlt() }
func (x AnaPrfmPnlUsdGet) Ifc() interface{}          { return x.FltFlt() }
func (x AnaPrfmPnlUsdGet) FltFlt() flt.Flt           { return x.X.AnaPrfm().PnlUsd }
func (x AnaPrfmPipAvgGet) Act()                      { x.FltFlt() }
func (x AnaPrfmPipAvgGet) Ifc() interface{}          { return x.FltFlt() }
func (x AnaPrfmPipAvgGet) FltFlt() flt.Flt           { return x.X.AnaPrfm().PipAvg }
func (x AnaPrfmPipMdnGet) Act()                      { x.FltFlt() }
func (x AnaPrfmPipMdnGet) Ifc() interface{}          { return x.FltFlt() }
func (x AnaPrfmPipMdnGet) FltFlt() flt.Flt           { return x.X.AnaPrfm().PipMdn }
func (x AnaPrfmPipMinGet) Act()                      { x.FltFlt() }
func (x AnaPrfmPipMinGet) Ifc() interface{}          { return x.FltFlt() }
func (x AnaPrfmPipMinGet) FltFlt() flt.Flt           { return x.X.AnaPrfm().PipMin }
func (x AnaPrfmPipMaxGet) Act()                      { x.FltFlt() }
func (x AnaPrfmPipMaxGet) Ifc() interface{}          { return x.FltFlt() }
func (x AnaPrfmPipMaxGet) FltFlt() flt.Flt           { return x.X.AnaPrfm().PipMax }
func (x AnaPrfmPipSumGet) Act()                      { x.FltFlt() }
func (x AnaPrfmPipSumGet) Ifc() interface{}          { return x.FltFlt() }
func (x AnaPrfmPipSumGet) FltFlt() flt.Flt           { return x.X.AnaPrfm().PipSum }
func (x AnaPrfmDurAvgGet) Act()                      { x.TmeTme() }
func (x AnaPrfmDurAvgGet) Ifc() interface{}          { return x.TmeTme() }
func (x AnaPrfmDurAvgGet) TmeTme() tme.Tme           { return x.X.AnaPrfm().DurAvg }
func (x AnaPrfmDurMdnGet) Act()                      { x.TmeTme() }
func (x AnaPrfmDurMdnGet) Ifc() interface{}          { return x.TmeTme() }
func (x AnaPrfmDurMdnGet) TmeTme() tme.Tme           { return x.X.AnaPrfm().DurMdn }
func (x AnaPrfmDurMinGet) Act()                      { x.TmeTme() }
func (x AnaPrfmDurMinGet) Ifc() interface{}          { return x.TmeTme() }
func (x AnaPrfmDurMinGet) TmeTme() tme.Tme           { return x.X.AnaPrfm().DurMin }
func (x AnaPrfmDurMaxGet) Act()                      { x.TmeTme() }
func (x AnaPrfmDurMaxGet) Ifc() interface{}          { return x.TmeTme() }
func (x AnaPrfmDurMaxGet) TmeTme() tme.Tme           { return x.X.AnaPrfm().DurMax }
func (x AnaPrfmLosLimMaxGet) Act()                   { x.FltFlt() }
func (x AnaPrfmLosLimMaxGet) Ifc() interface{}       { return x.FltFlt() }
func (x AnaPrfmLosLimMaxGet) FltFlt() flt.Flt        { return x.X.AnaPrfm().LosLimMax }
func (x AnaPrfmDurLimMaxGet) Act()                   { x.TmeTme() }
func (x AnaPrfmDurLimMaxGet) Ifc() interface{}       { return x.TmeTme() }
func (x AnaPrfmDurLimMaxGet) TmeTme() tme.Tme        { return x.X.AnaPrfm().DurLimMax }
func (x AnaPrfmDayCntGet) Act()                      { x.UntUnt() }
func (x AnaPrfmDayCntGet) Ifc() interface{}          { return x.UntUnt() }
func (x AnaPrfmDayCntGet) UntUnt() unt.Unt           { return x.X.AnaPrfm().DayCnt }
func (x AnaPrfmTrdCntGet) Act()                      { x.UntUnt() }
func (x AnaPrfmTrdCntGet) Ifc() interface{}          { return x.UntUnt() }
func (x AnaPrfmTrdCntGet) UntUnt() unt.Unt           { return x.X.AnaPrfm().TrdCnt }
func (x AnaPrfmTrdPctGet) Act()                      { x.FltFlt() }
func (x AnaPrfmTrdPctGet) Ifc() interface{}          { return x.FltFlt() }
func (x AnaPrfmTrdPctGet) FltFlt() flt.Flt           { return x.X.AnaPrfm().TrdPct }
func (x AnaPrfmCstTotUsdGet) Act()                   { x.FltFlt() }
func (x AnaPrfmCstTotUsdGet) Ifc() interface{}       { return x.FltFlt() }
func (x AnaPrfmCstTotUsdGet) FltFlt() flt.Flt        { return x.X.AnaPrfm().CstTotUsd }
func (x AnaPrfmCstSpdUsdGet) Act()                   { x.FltFlt() }
func (x AnaPrfmCstSpdUsdGet) Ifc() interface{}       { return x.FltFlt() }
func (x AnaPrfmCstSpdUsdGet) FltFlt() flt.Flt        { return x.X.AnaPrfm().CstSpdUsd }
func (x AnaPrfmCstComUsdGet) Act()                   { x.FltFlt() }
func (x AnaPrfmCstComUsdGet) Ifc() interface{}       { return x.FltFlt() }
func (x AnaPrfmCstComUsdGet) FltFlt() flt.Flt        { return x.X.AnaPrfm().CstComUsd }
func (x AnaPrfmDltPnlPctAGet) Act()                  { x.FltFlt() }
func (x AnaPrfmDltPnlPctAGet) Ifc() interface{}      { return x.FltFlt() }
func (x AnaPrfmDltPnlPctAGet) FltFlt() flt.Flt       { return x.X.AnaPrfmDlt().PnlPctA }
func (x AnaPrfmDltPnlPctBGet) Act()                  { x.FltFlt() }
func (x AnaPrfmDltPnlPctBGet) Ifc() interface{}      { return x.FltFlt() }
func (x AnaPrfmDltPnlPctBGet) FltFlt() flt.Flt       { return x.X.AnaPrfmDlt().PnlPctB }
func (x AnaPrfmDltPnlPctDltGet) Act()                { x.FltFlt() }
func (x AnaPrfmDltPnlPctDltGet) Ifc() interface{}    { return x.FltFlt() }
func (x AnaPrfmDltPnlPctDltGet) FltFlt() flt.Flt     { return x.X.AnaPrfmDlt().PnlPctDlt }
func (x AnaPrfmDltScsPctAGet) Act()                  { x.FltFlt() }
func (x AnaPrfmDltScsPctAGet) Ifc() interface{}      { return x.FltFlt() }
func (x AnaPrfmDltScsPctAGet) FltFlt() flt.Flt       { return x.X.AnaPrfmDlt().ScsPctA }
func (x AnaPrfmDltScsPctBGet) Act()                  { x.FltFlt() }
func (x AnaPrfmDltScsPctBGet) Ifc() interface{}      { return x.FltFlt() }
func (x AnaPrfmDltScsPctBGet) FltFlt() flt.Flt       { return x.X.AnaPrfmDlt().ScsPctB }
func (x AnaPrfmDltScsPctDltGet) Act()                { x.FltFlt() }
func (x AnaPrfmDltScsPctDltGet) Ifc() interface{}    { return x.FltFlt() }
func (x AnaPrfmDltScsPctDltGet) FltFlt() flt.Flt     { return x.X.AnaPrfmDlt().ScsPctDlt }
func (x AnaPrfmDltPipPerDayAGet) Act()               { x.FltFlt() }
func (x AnaPrfmDltPipPerDayAGet) Ifc() interface{}   { return x.FltFlt() }
func (x AnaPrfmDltPipPerDayAGet) FltFlt() flt.Flt    { return x.X.AnaPrfmDlt().PipPerDayA }
func (x AnaPrfmDltPipPerDayBGet) Act()               { x.FltFlt() }
func (x AnaPrfmDltPipPerDayBGet) Ifc() interface{}   { return x.FltFlt() }
func (x AnaPrfmDltPipPerDayBGet) FltFlt() flt.Flt    { return x.X.AnaPrfmDlt().PipPerDayB }
func (x AnaPrfmDltPipPerDayDltGet) Act()             { x.FltFlt() }
func (x AnaPrfmDltPipPerDayDltGet) Ifc() interface{} { return x.FltFlt() }
func (x AnaPrfmDltPipPerDayDltGet) FltFlt() flt.Flt  { return x.X.AnaPrfmDlt().PipPerDayDlt }
func (x AnaPrfmDltUsdPerDayAGet) Act()               { x.FltFlt() }
func (x AnaPrfmDltUsdPerDayAGet) Ifc() interface{}   { return x.FltFlt() }
func (x AnaPrfmDltUsdPerDayAGet) FltFlt() flt.Flt    { return x.X.AnaPrfmDlt().UsdPerDayA }
func (x AnaPrfmDltUsdPerDayBGet) Act()               { x.FltFlt() }
func (x AnaPrfmDltUsdPerDayBGet) Ifc() interface{}   { return x.FltFlt() }
func (x AnaPrfmDltUsdPerDayBGet) FltFlt() flt.Flt    { return x.X.AnaPrfmDlt().UsdPerDayB }
func (x AnaPrfmDltUsdPerDayDltGet) Act()             { x.FltFlt() }
func (x AnaPrfmDltUsdPerDayDltGet) Ifc() interface{} { return x.FltFlt() }
func (x AnaPrfmDltUsdPerDayDltGet) FltFlt() flt.Flt  { return x.X.AnaPrfmDlt().UsdPerDayDlt }
func (x AnaPrfmDltScsPerDayAGet) Act()               { x.FltFlt() }
func (x AnaPrfmDltScsPerDayAGet) Ifc() interface{}   { return x.FltFlt() }
func (x AnaPrfmDltScsPerDayAGet) FltFlt() flt.Flt    { return x.X.AnaPrfmDlt().ScsPerDayA }
func (x AnaPrfmDltScsPerDayBGet) Act()               { x.FltFlt() }
func (x AnaPrfmDltScsPerDayBGet) Ifc() interface{}   { return x.FltFlt() }
func (x AnaPrfmDltScsPerDayBGet) FltFlt() flt.Flt    { return x.X.AnaPrfmDlt().ScsPerDayB }
func (x AnaPrfmDltScsPerDayDltGet) Act()             { x.FltFlt() }
func (x AnaPrfmDltScsPerDayDltGet) Ifc() interface{} { return x.FltFlt() }
func (x AnaPrfmDltScsPerDayDltGet) FltFlt() flt.Flt  { return x.X.AnaPrfmDlt().ScsPerDayDlt }
func (x AnaPrfmDltOpnPerDayAGet) Act()               { x.FltFlt() }
func (x AnaPrfmDltOpnPerDayAGet) Ifc() interface{}   { return x.FltFlt() }
func (x AnaPrfmDltOpnPerDayAGet) FltFlt() flt.Flt    { return x.X.AnaPrfmDlt().OpnPerDayA }
func (x AnaPrfmDltOpnPerDayBGet) Act()               { x.FltFlt() }
func (x AnaPrfmDltOpnPerDayBGet) Ifc() interface{}   { return x.FltFlt() }
func (x AnaPrfmDltOpnPerDayBGet) FltFlt() flt.Flt    { return x.X.AnaPrfmDlt().OpnPerDayB }
func (x AnaPrfmDltOpnPerDayDltGet) Act()             { x.FltFlt() }
func (x AnaPrfmDltOpnPerDayDltGet) Ifc() interface{} { return x.FltFlt() }
func (x AnaPrfmDltOpnPerDayDltGet) FltFlt() flt.Flt  { return x.X.AnaPrfmDlt().OpnPerDayDlt }
func (x AnaPrfmDltPnlUsdAGet) Act()                  { x.FltFlt() }
func (x AnaPrfmDltPnlUsdAGet) Ifc() interface{}      { return x.FltFlt() }
func (x AnaPrfmDltPnlUsdAGet) FltFlt() flt.Flt       { return x.X.AnaPrfmDlt().PnlUsdA }
func (x AnaPrfmDltPnlUsdBGet) Act()                  { x.FltFlt() }
func (x AnaPrfmDltPnlUsdBGet) Ifc() interface{}      { return x.FltFlt() }
func (x AnaPrfmDltPnlUsdBGet) FltFlt() flt.Flt       { return x.X.AnaPrfmDlt().PnlUsdB }
func (x AnaPrfmDltPnlUsdDltGet) Act()                { x.FltFlt() }
func (x AnaPrfmDltPnlUsdDltGet) Ifc() interface{}    { return x.FltFlt() }
func (x AnaPrfmDltPnlUsdDltGet) FltFlt() flt.Flt     { return x.X.AnaPrfmDlt().PnlUsdDlt }
func (x AnaPrfmDltPipAvgAGet) Act()                  { x.FltFlt() }
func (x AnaPrfmDltPipAvgAGet) Ifc() interface{}      { return x.FltFlt() }
func (x AnaPrfmDltPipAvgAGet) FltFlt() flt.Flt       { return x.X.AnaPrfmDlt().PipAvgA }
func (x AnaPrfmDltPipAvgBGet) Act()                  { x.FltFlt() }
func (x AnaPrfmDltPipAvgBGet) Ifc() interface{}      { return x.FltFlt() }
func (x AnaPrfmDltPipAvgBGet) FltFlt() flt.Flt       { return x.X.AnaPrfmDlt().PipAvgB }
func (x AnaPrfmDltPipAvgDltGet) Act()                { x.FltFlt() }
func (x AnaPrfmDltPipAvgDltGet) Ifc() interface{}    { return x.FltFlt() }
func (x AnaPrfmDltPipAvgDltGet) FltFlt() flt.Flt     { return x.X.AnaPrfmDlt().PipAvgDlt }
func (x AnaPrfmDltPipMdnAGet) Act()                  { x.FltFlt() }
func (x AnaPrfmDltPipMdnAGet) Ifc() interface{}      { return x.FltFlt() }
func (x AnaPrfmDltPipMdnAGet) FltFlt() flt.Flt       { return x.X.AnaPrfmDlt().PipMdnA }
func (x AnaPrfmDltPipMdnBGet) Act()                  { x.FltFlt() }
func (x AnaPrfmDltPipMdnBGet) Ifc() interface{}      { return x.FltFlt() }
func (x AnaPrfmDltPipMdnBGet) FltFlt() flt.Flt       { return x.X.AnaPrfmDlt().PipMdnB }
func (x AnaPrfmDltPipMdnDltGet) Act()                { x.FltFlt() }
func (x AnaPrfmDltPipMdnDltGet) Ifc() interface{}    { return x.FltFlt() }
func (x AnaPrfmDltPipMdnDltGet) FltFlt() flt.Flt     { return x.X.AnaPrfmDlt().PipMdnDlt }
func (x AnaPrfmDltPipMinAGet) Act()                  { x.FltFlt() }
func (x AnaPrfmDltPipMinAGet) Ifc() interface{}      { return x.FltFlt() }
func (x AnaPrfmDltPipMinAGet) FltFlt() flt.Flt       { return x.X.AnaPrfmDlt().PipMinA }
func (x AnaPrfmDltPipMinBGet) Act()                  { x.FltFlt() }
func (x AnaPrfmDltPipMinBGet) Ifc() interface{}      { return x.FltFlt() }
func (x AnaPrfmDltPipMinBGet) FltFlt() flt.Flt       { return x.X.AnaPrfmDlt().PipMinB }
func (x AnaPrfmDltPipMinDltGet) Act()                { x.FltFlt() }
func (x AnaPrfmDltPipMinDltGet) Ifc() interface{}    { return x.FltFlt() }
func (x AnaPrfmDltPipMinDltGet) FltFlt() flt.Flt     { return x.X.AnaPrfmDlt().PipMinDlt }
func (x AnaPrfmDltPipMaxAGet) Act()                  { x.FltFlt() }
func (x AnaPrfmDltPipMaxAGet) Ifc() interface{}      { return x.FltFlt() }
func (x AnaPrfmDltPipMaxAGet) FltFlt() flt.Flt       { return x.X.AnaPrfmDlt().PipMaxA }
func (x AnaPrfmDltPipMaxBGet) Act()                  { x.FltFlt() }
func (x AnaPrfmDltPipMaxBGet) Ifc() interface{}      { return x.FltFlt() }
func (x AnaPrfmDltPipMaxBGet) FltFlt() flt.Flt       { return x.X.AnaPrfmDlt().PipMaxB }
func (x AnaPrfmDltPipMaxDltGet) Act()                { x.FltFlt() }
func (x AnaPrfmDltPipMaxDltGet) Ifc() interface{}    { return x.FltFlt() }
func (x AnaPrfmDltPipMaxDltGet) FltFlt() flt.Flt     { return x.X.AnaPrfmDlt().PipMaxDlt }
func (x AnaPrfmDltPipSumAGet) Act()                  { x.FltFlt() }
func (x AnaPrfmDltPipSumAGet) Ifc() interface{}      { return x.FltFlt() }
func (x AnaPrfmDltPipSumAGet) FltFlt() flt.Flt       { return x.X.AnaPrfmDlt().PipSumA }
func (x AnaPrfmDltPipSumBGet) Act()                  { x.FltFlt() }
func (x AnaPrfmDltPipSumBGet) Ifc() interface{}      { return x.FltFlt() }
func (x AnaPrfmDltPipSumBGet) FltFlt() flt.Flt       { return x.X.AnaPrfmDlt().PipSumB }
func (x AnaPrfmDltPipSumDltGet) Act()                { x.FltFlt() }
func (x AnaPrfmDltPipSumDltGet) Ifc() interface{}    { return x.FltFlt() }
func (x AnaPrfmDltPipSumDltGet) FltFlt() flt.Flt     { return x.X.AnaPrfmDlt().PipSumDlt }
func (x AnaPrfmDltDurAvgAGet) Act()                  { x.TmeTme() }
func (x AnaPrfmDltDurAvgAGet) Ifc() interface{}      { return x.TmeTme() }
func (x AnaPrfmDltDurAvgAGet) TmeTme() tme.Tme       { return x.X.AnaPrfmDlt().DurAvgA }
func (x AnaPrfmDltDurAvgBGet) Act()                  { x.TmeTme() }
func (x AnaPrfmDltDurAvgBGet) Ifc() interface{}      { return x.TmeTme() }
func (x AnaPrfmDltDurAvgBGet) TmeTme() tme.Tme       { return x.X.AnaPrfmDlt().DurAvgB }
func (x AnaPrfmDltDurAvgDltGet) Act()                { x.FltFlt() }
func (x AnaPrfmDltDurAvgDltGet) Ifc() interface{}    { return x.FltFlt() }
func (x AnaPrfmDltDurAvgDltGet) FltFlt() flt.Flt     { return x.X.AnaPrfmDlt().DurAvgDlt }
func (x AnaPrfmDltDurMdnAGet) Act()                  { x.TmeTme() }
func (x AnaPrfmDltDurMdnAGet) Ifc() interface{}      { return x.TmeTme() }
func (x AnaPrfmDltDurMdnAGet) TmeTme() tme.Tme       { return x.X.AnaPrfmDlt().DurMdnA }
func (x AnaPrfmDltDurMdnBGet) Act()                  { x.TmeTme() }
func (x AnaPrfmDltDurMdnBGet) Ifc() interface{}      { return x.TmeTme() }
func (x AnaPrfmDltDurMdnBGet) TmeTme() tme.Tme       { return x.X.AnaPrfmDlt().DurMdnB }
func (x AnaPrfmDltDurMdnDltGet) Act()                { x.FltFlt() }
func (x AnaPrfmDltDurMdnDltGet) Ifc() interface{}    { return x.FltFlt() }
func (x AnaPrfmDltDurMdnDltGet) FltFlt() flt.Flt     { return x.X.AnaPrfmDlt().DurMdnDlt }
func (x AnaPrfmDltDurMinAGet) Act()                  { x.TmeTme() }
func (x AnaPrfmDltDurMinAGet) Ifc() interface{}      { return x.TmeTme() }
func (x AnaPrfmDltDurMinAGet) TmeTme() tme.Tme       { return x.X.AnaPrfmDlt().DurMinA }
func (x AnaPrfmDltDurMinBGet) Act()                  { x.TmeTme() }
func (x AnaPrfmDltDurMinBGet) Ifc() interface{}      { return x.TmeTme() }
func (x AnaPrfmDltDurMinBGet) TmeTme() tme.Tme       { return x.X.AnaPrfmDlt().DurMinB }
func (x AnaPrfmDltDurMinDltGet) Act()                { x.FltFlt() }
func (x AnaPrfmDltDurMinDltGet) Ifc() interface{}    { return x.FltFlt() }
func (x AnaPrfmDltDurMinDltGet) FltFlt() flt.Flt     { return x.X.AnaPrfmDlt().DurMinDlt }
func (x AnaPrfmDltDurMaxAGet) Act()                  { x.TmeTme() }
func (x AnaPrfmDltDurMaxAGet) Ifc() interface{}      { return x.TmeTme() }
func (x AnaPrfmDltDurMaxAGet) TmeTme() tme.Tme       { return x.X.AnaPrfmDlt().DurMaxA }
func (x AnaPrfmDltDurMaxBGet) Act()                  { x.TmeTme() }
func (x AnaPrfmDltDurMaxBGet) Ifc() interface{}      { return x.TmeTme() }
func (x AnaPrfmDltDurMaxBGet) TmeTme() tme.Tme       { return x.X.AnaPrfmDlt().DurMaxB }
func (x AnaPrfmDltDurMaxDltGet) Act()                { x.FltFlt() }
func (x AnaPrfmDltDurMaxDltGet) Ifc() interface{}    { return x.FltFlt() }
func (x AnaPrfmDltDurMaxDltGet) FltFlt() flt.Flt     { return x.X.AnaPrfmDlt().DurMaxDlt }
func (x AnaPrfmDltTrdCntAGet) Act()                  { x.UntUnt() }
func (x AnaPrfmDltTrdCntAGet) Ifc() interface{}      { return x.UntUnt() }
func (x AnaPrfmDltTrdCntAGet) UntUnt() unt.Unt       { return x.X.AnaPrfmDlt().TrdCntA }
func (x AnaPrfmDltTrdCntBGet) Act()                  { x.UntUnt() }
func (x AnaPrfmDltTrdCntBGet) Ifc() interface{}      { return x.UntUnt() }
func (x AnaPrfmDltTrdCntBGet) UntUnt() unt.Unt       { return x.X.AnaPrfmDlt().TrdCntB }
func (x AnaPrfmDltTrdCntDltGet) Act()                { x.FltFlt() }
func (x AnaPrfmDltTrdCntDltGet) Ifc() interface{}    { return x.FltFlt() }
func (x AnaPrfmDltTrdCntDltGet) FltFlt() flt.Flt     { return x.X.AnaPrfmDlt().TrdCntDlt }
func (x AnaPrfmDltTrdPctAGet) Act()                  { x.FltFlt() }
func (x AnaPrfmDltTrdPctAGet) Ifc() interface{}      { return x.FltFlt() }
func (x AnaPrfmDltTrdPctAGet) FltFlt() flt.Flt       { return x.X.AnaPrfmDlt().TrdPctA }
func (x AnaPrfmDltTrdPctBGet) Act()                  { x.FltFlt() }
func (x AnaPrfmDltTrdPctBGet) Ifc() interface{}      { return x.FltFlt() }
func (x AnaPrfmDltTrdPctBGet) FltFlt() flt.Flt       { return x.X.AnaPrfmDlt().TrdPctB }
func (x AnaPrfmDltTrdPctDltGet) Act()                { x.FltFlt() }
func (x AnaPrfmDltTrdPctDltGet) Ifc() interface{}    { return x.FltFlt() }
func (x AnaPrfmDltTrdPctDltGet) FltFlt() flt.Flt     { return x.X.AnaPrfmDlt().TrdPctDlt }
func (x AnaPrfmDltPthBGet) Act()                     { x.StrStr() }
func (x AnaPrfmDltPthBGet) Ifc() interface{}         { return x.StrStr() }
func (x AnaPrfmDltPthBGet) StrStr() str.Str          { return x.X.AnaPrfmDlt().PthB }
func (x PenPenClrSetGet) Act()                       { x.ClrClr() }
func (x PenPenClrSetGet) Ifc() interface{}           { return x.ClrClr() }
func (x PenPenClrSetGet) ClrClr() clr.Clr {
	v := x.X.PenPen()
	if x.I0 != nil {
		v.Clr = x.I0.ClrClr()
	}
	return v.Clr
}
func (x PenPenWidSetGet) Act()             { x.UntUnt() }
func (x PenPenWidSetGet) Ifc() interface{} { return x.UntUnt() }
func (x PenPenWidSetGet) UntUnt() unt.Unt {
	v := x.X.PenPen()
	if x.I0 != nil {
		v.Wid = x.I0.UntUnt()
	}
	return v.Wid
}
func (x PltFltAxisYMinSetGet) Act()             { x.FltFlt() }
func (x PltFltAxisYMinSetGet) Ifc() interface{} { return x.FltFlt() }
func (x PltFltAxisYMinSetGet) FltFlt() flt.Flt {
	v := x.X.PltFltAxisY()
	if x.I0 != nil {
		v.Min = x.I0.FltFlt()
	}
	return v.Min
}
func (x PltFltAxisYMaxSetGet) Act()             { x.FltFlt() }
func (x PltFltAxisYMaxSetGet) Ifc() interface{} { return x.FltFlt() }
func (x PltFltAxisYMaxSetGet) FltFlt() flt.Flt {
	v := x.X.PltFltAxisY()
	if x.I0 != nil {
		v.Max = x.I0.FltFlt()
	}
	return v.Max
}
func (x PltFltAxisYEqiDstSetGet) Act()             { x.FltFlt() }
func (x PltFltAxisYEqiDstSetGet) Ifc() interface{} { return x.FltFlt() }
func (x PltFltAxisYEqiDstSetGet) FltFlt() flt.Flt {
	v := x.X.PltFltAxisY()
	if x.I0 != nil {
		v.EqiDst = x.I0.FltFlt()
	}
	return v.EqiDst
}
func (x PltStmTitleSetGet) Act()             { x.StrStr() }
func (x PltStmTitleSetGet) Ifc() interface{} { return x.StrStr() }
func (x PltStmTitleSetGet) StrStr() str.Str {
	v := x.X.PltStm()
	if x.I0 != nil {
		v.Title = x.I0.StrStr()
	}
	return v.Title
}
func (x PltFltsSctrYGet) Act()                       { x.PltFltAxisY() }
func (x PltFltsSctrYGet) Ifc() interface{}           { return x.PltFltAxisY() }
func (x PltFltsSctrYGet) PltFltAxisY() *plt.FltAxisY { return x.X.PltFltsSctr().Y }
func (x PltFltsSctrTitleSetGet) Act()                { x.StrStr() }
func (x PltFltsSctrTitleSetGet) Ifc() interface{}    { return x.StrStr() }
func (x PltFltsSctrTitleSetGet) StrStr() str.Str {
	v := x.X.PltFltsSctr()
	if x.I0 != nil {
		v.Title = x.I0.StrStr()
	}
	return v.Title
}
func (x PltFltsSctrOutlierSetGet) Act()             { x.BolBol() }
func (x PltFltsSctrOutlierSetGet) Ifc() interface{} { return x.BolBol() }
func (x PltFltsSctrOutlierSetGet) BolBol() bol.Bol {
	v := x.X.PltFltsSctr()
	if x.I0 != nil {
		v.Outlier = x.I0.BolBol()
	}
	return v.Outlier
}
func (x PltHrzPltsGet) Act()                 { x.PltPlts() }
func (x PltHrzPltsGet) Ifc() interface{}     { return x.PltPlts() }
func (x PltHrzPltsGet) PltPlts() *plt.Plts   { return x.X.PltHrz().Plts }
func (x PltVrtPltsGet) Act()                 { x.PltPlts() }
func (x PltVrtPltsGet) Ifc() interface{}     { return x.PltPlts() }
func (x PltVrtPltsGet) PltPlts() *plt.Plts   { return x.X.PltVrt().Plts }
func (x PltDpthPltsGet) Act()                { x.PltPlts() }
func (x PltDpthPltsGet) Ifc() interface{}    { return x.PltPlts() }
func (x PltDpthPltsGet) PltPlts() *plt.Plts  { return x.X.PltDpth().Plts }
func (x StrZero) Act()                       { x.StrStr() }
func (x StrZero) Ifc() interface{}           { return x.StrStr() }
func (x StrZero) StrStr() str.Str            { return str.Zero }
func (x StrEmpty) Act()                      { x.StrStr() }
func (x StrEmpty) Ifc() interface{}          { return x.StrStr() }
func (x StrEmpty) StrStr() str.Str           { return str.Empty }
func (x BolZero) Act()                       { x.BolBol() }
func (x BolZero) Ifc() interface{}           { return x.BolBol() }
func (x BolZero) BolBol() bol.Bol            { return bol.Zero }
func (x BolFls) Act()                        { x.BolBol() }
func (x BolFls) Ifc() interface{}            { return x.BolBol() }
func (x BolFls) BolBol() bol.Bol             { return bol.Fls }
func (x BolTru) Act()                        { x.BolBol() }
func (x BolTru) Ifc() interface{}            { return x.BolBol() }
func (x BolTru) BolBol() bol.Bol             { return bol.Tru }
func (x FltZero) Act()                       { x.FltFlt() }
func (x FltZero) Ifc() interface{}           { return x.FltFlt() }
func (x FltZero) FltFlt() flt.Flt            { return flt.Zero }
func (x FltOne) Act()                        { x.FltFlt() }
func (x FltOne) Ifc() interface{}            { return x.FltFlt() }
func (x FltOne) FltFlt() flt.Flt             { return flt.One }
func (x FltNegOne) Act()                     { x.FltFlt() }
func (x FltNegOne) Ifc() interface{}         { return x.FltFlt() }
func (x FltNegOne) FltFlt() flt.Flt          { return flt.NegOne }
func (x FltHndrd) Act()                      { x.FltFlt() }
func (x FltHndrd) Ifc() interface{}          { return x.FltFlt() }
func (x FltHndrd) FltFlt() flt.Flt           { return flt.Hndrd }
func (x FltMin) Act()                        { x.FltFlt() }
func (x FltMin) Ifc() interface{}            { return x.FltFlt() }
func (x FltMin) FltFlt() flt.Flt             { return flt.Min }
func (x FltMax) Act()                        { x.FltFlt() }
func (x FltMax) Ifc() interface{}            { return x.FltFlt() }
func (x FltMax) FltFlt() flt.Flt             { return flt.Max }
func (x FltTiny) Act()                       { x.FltFlt() }
func (x FltTiny) Ifc() interface{}           { return x.FltFlt() }
func (x FltTiny) FltFlt() flt.Flt            { return flt.Tiny }
func (x UntZero) Act()                       { x.UntUnt() }
func (x UntZero) Ifc() interface{}           { return x.UntUnt() }
func (x UntZero) UntUnt() unt.Unt            { return unt.Zero }
func (x UntOne) Act()                        { x.UntUnt() }
func (x UntOne) Ifc() interface{}            { return x.UntUnt() }
func (x UntOne) UntUnt() unt.Unt             { return unt.One }
func (x UntMin) Act()                        { x.UntUnt() }
func (x UntMin) Ifc() interface{}            { return x.UntUnt() }
func (x UntMin) UntUnt() unt.Unt             { return unt.Min }
func (x UntMax) Act()                        { x.UntUnt() }
func (x UntMax) Ifc() interface{}            { return x.UntUnt() }
func (x UntMax) UntUnt() unt.Unt             { return unt.Max }
func (x IntZero) Act()                       { x.IntInt() }
func (x IntZero) Ifc() interface{}           { return x.IntInt() }
func (x IntZero) IntInt() int.Int            { return int.Zero }
func (x IntOne) Act()                        { x.IntInt() }
func (x IntOne) Ifc() interface{}            { return x.IntInt() }
func (x IntOne) IntInt() int.Int             { return int.One }
func (x IntNegOne) Act()                     { x.IntInt() }
func (x IntNegOne) Ifc() interface{}         { return x.IntInt() }
func (x IntNegOne) IntInt() int.Int          { return int.NegOne }
func (x IntMin) Act()                        { x.IntInt() }
func (x IntMin) Ifc() interface{}            { return x.IntInt() }
func (x IntMin) IntInt() int.Int             { return int.Min }
func (x IntMax) Act()                        { x.IntInt() }
func (x IntMax) Ifc() interface{}            { return x.IntInt() }
func (x IntMax) IntInt() int.Int             { return int.Max }
func (x TmeZero) Act()                       { x.TmeTme() }
func (x TmeZero) Ifc() interface{}           { return x.TmeTme() }
func (x TmeZero) TmeTme() tme.Tme            { return tme.Zero }
func (x TmeOne) Act()                        { x.TmeTme() }
func (x TmeOne) Ifc() interface{}            { return x.TmeTme() }
func (x TmeOne) TmeTme() tme.Tme             { return tme.One }
func (x TmeNegOne) Act()                     { x.TmeTme() }
func (x TmeNegOne) Ifc() interface{}         { return x.TmeTme() }
func (x TmeNegOne) TmeTme() tme.Tme          { return tme.NegOne }
func (x TmeMin) Act()                        { x.TmeTme() }
func (x TmeMin) Ifc() interface{}            { return x.TmeTme() }
func (x TmeMin) TmeTme() tme.Tme             { return tme.Min }
func (x TmeMax) Act()                        { x.TmeTme() }
func (x TmeMax) Ifc() interface{}            { return x.TmeTme() }
func (x TmeMax) TmeTme() tme.Tme             { return tme.Max }
func (x TmeSecond) Act()                     { x.TmeTme() }
func (x TmeSecond) Ifc() interface{}         { return x.TmeTme() }
func (x TmeSecond) TmeTme() tme.Tme          { return tme.Second }
func (x TmeMinute) Act()                     { x.TmeTme() }
func (x TmeMinute) Ifc() interface{}         { return x.TmeTme() }
func (x TmeMinute) TmeTme() tme.Tme          { return tme.Minute }
func (x TmeHour) Act()                       { x.TmeTme() }
func (x TmeHour) Ifc() interface{}           { return x.TmeTme() }
func (x TmeHour) TmeTme() tme.Tme            { return tme.Hour }
func (x TmeDay) Act()                        { x.TmeTme() }
func (x TmeDay) Ifc() interface{}            { return x.TmeTme() }
func (x TmeDay) TmeTme() tme.Tme             { return tme.Day }
func (x TmeWeek) Act()                       { x.TmeTme() }
func (x TmeWeek) Ifc() interface{}           { return x.TmeTme() }
func (x TmeWeek) TmeTme() tme.Tme            { return tme.Week }
func (x TmeS1) Act()                         { x.TmeTme() }
func (x TmeS1) Ifc() interface{}             { return x.TmeTme() }
func (x TmeS1) TmeTme() tme.Tme              { return tme.S1 }
func (x TmeS5) Act()                         { x.TmeTme() }
func (x TmeS5) Ifc() interface{}             { return x.TmeTme() }
func (x TmeS5) TmeTme() tme.Tme              { return tme.S5 }
func (x TmeS10) Act()                        { x.TmeTme() }
func (x TmeS10) Ifc() interface{}            { return x.TmeTme() }
func (x TmeS10) TmeTme() tme.Tme             { return tme.S10 }
func (x TmeS15) Act()                        { x.TmeTme() }
func (x TmeS15) Ifc() interface{}            { return x.TmeTme() }
func (x TmeS15) TmeTme() tme.Tme             { return tme.S15 }
func (x TmeS20) Act()                        { x.TmeTme() }
func (x TmeS20) Ifc() interface{}            { return x.TmeTme() }
func (x TmeS20) TmeTme() tme.Tme             { return tme.S20 }
func (x TmeS30) Act()                        { x.TmeTme() }
func (x TmeS30) Ifc() interface{}            { return x.TmeTme() }
func (x TmeS30) TmeTme() tme.Tme             { return tme.S30 }
func (x TmeS40) Act()                        { x.TmeTme() }
func (x TmeS40) Ifc() interface{}            { return x.TmeTme() }
func (x TmeS40) TmeTme() tme.Tme             { return tme.S40 }
func (x TmeS50) Act()                        { x.TmeTme() }
func (x TmeS50) Ifc() interface{}            { return x.TmeTme() }
func (x TmeS50) TmeTme() tme.Tme             { return tme.S50 }
func (x TmeM1) Act()                         { x.TmeTme() }
func (x TmeM1) Ifc() interface{}             { return x.TmeTme() }
func (x TmeM1) TmeTme() tme.Tme              { return tme.M1 }
func (x TmeM5) Act()                         { x.TmeTme() }
func (x TmeM5) Ifc() interface{}             { return x.TmeTme() }
func (x TmeM5) TmeTme() tme.Tme              { return tme.M5 }
func (x TmeM10) Act()                        { x.TmeTme() }
func (x TmeM10) Ifc() interface{}            { return x.TmeTme() }
func (x TmeM10) TmeTme() tme.Tme             { return tme.M10 }
func (x TmeM15) Act()                        { x.TmeTme() }
func (x TmeM15) Ifc() interface{}            { return x.TmeTme() }
func (x TmeM15) TmeTme() tme.Tme             { return tme.M15 }
func (x TmeM20) Act()                        { x.TmeTme() }
func (x TmeM20) Ifc() interface{}            { return x.TmeTme() }
func (x TmeM20) TmeTme() tme.Tme             { return tme.M20 }
func (x TmeM30) Act()                        { x.TmeTme() }
func (x TmeM30) Ifc() interface{}            { return x.TmeTme() }
func (x TmeM30) TmeTme() tme.Tme             { return tme.M30 }
func (x TmeM40) Act()                        { x.TmeTme() }
func (x TmeM40) Ifc() interface{}            { return x.TmeTme() }
func (x TmeM40) TmeTme() tme.Tme             { return tme.M40 }
func (x TmeM50) Act()                        { x.TmeTme() }
func (x TmeM50) Ifc() interface{}            { return x.TmeTme() }
func (x TmeM50) TmeTme() tme.Tme             { return tme.M50 }
func (x TmeH1) Act()                         { x.TmeTme() }
func (x TmeH1) Ifc() interface{}             { return x.TmeTme() }
func (x TmeH1) TmeTme() tme.Tme              { return tme.H1 }
func (x TmeD1) Act()                         { x.TmeTme() }
func (x TmeD1) Ifc() interface{}             { return x.TmeTme() }
func (x TmeD1) TmeTme() tme.Tme              { return tme.D1 }
func (x TmeResolution) Act()                 { x.TmeTme() }
func (x TmeResolution) Ifc() interface{}     { return x.TmeTme() }
func (x TmeResolution) TmeTme() tme.Tme      { return tme.Resolution }
func (x ClrBlack) Act()                      { x.ClrClr() }
func (x ClrBlack) Ifc() interface{}          { return x.ClrClr() }
func (x ClrBlack) ClrClr() clr.Clr           { return clr.Black }
func (x ClrWhite) Act()                      { x.ClrClr() }
func (x ClrWhite) Ifc() interface{}          { return x.ClrClr() }
func (x ClrWhite) ClrClr() clr.Clr           { return clr.White }
func (x ClrRed50) Act()                      { x.ClrClr() }
func (x ClrRed50) Ifc() interface{}          { return x.ClrClr() }
func (x ClrRed50) ClrClr() clr.Clr           { return clr.Red50 }
func (x ClrRed100) Act()                     { x.ClrClr() }
func (x ClrRed100) Ifc() interface{}         { return x.ClrClr() }
func (x ClrRed100) ClrClr() clr.Clr          { return clr.Red100 }
func (x ClrRed200) Act()                     { x.ClrClr() }
func (x ClrRed200) Ifc() interface{}         { return x.ClrClr() }
func (x ClrRed200) ClrClr() clr.Clr          { return clr.Red200 }
func (x ClrRed300) Act()                     { x.ClrClr() }
func (x ClrRed300) Ifc() interface{}         { return x.ClrClr() }
func (x ClrRed300) ClrClr() clr.Clr          { return clr.Red300 }
func (x ClrRed400) Act()                     { x.ClrClr() }
func (x ClrRed400) Ifc() interface{}         { return x.ClrClr() }
func (x ClrRed400) ClrClr() clr.Clr          { return clr.Red400 }
func (x ClrRed500) Act()                     { x.ClrClr() }
func (x ClrRed500) Ifc() interface{}         { return x.ClrClr() }
func (x ClrRed500) ClrClr() clr.Clr          { return clr.Red500 }
func (x ClrRed600) Act()                     { x.ClrClr() }
func (x ClrRed600) Ifc() interface{}         { return x.ClrClr() }
func (x ClrRed600) ClrClr() clr.Clr          { return clr.Red600 }
func (x ClrRed700) Act()                     { x.ClrClr() }
func (x ClrRed700) Ifc() interface{}         { return x.ClrClr() }
func (x ClrRed700) ClrClr() clr.Clr          { return clr.Red700 }
func (x ClrRed800) Act()                     { x.ClrClr() }
func (x ClrRed800) Ifc() interface{}         { return x.ClrClr() }
func (x ClrRed800) ClrClr() clr.Clr          { return clr.Red800 }
func (x ClrRed900) Act()                     { x.ClrClr() }
func (x ClrRed900) Ifc() interface{}         { return x.ClrClr() }
func (x ClrRed900) ClrClr() clr.Clr          { return clr.Red900 }
func (x ClrRedA100) Act()                    { x.ClrClr() }
func (x ClrRedA100) Ifc() interface{}        { return x.ClrClr() }
func (x ClrRedA100) ClrClr() clr.Clr         { return clr.RedA100 }
func (x ClrRedA200) Act()                    { x.ClrClr() }
func (x ClrRedA200) Ifc() interface{}        { return x.ClrClr() }
func (x ClrRedA200) ClrClr() clr.Clr         { return clr.RedA200 }
func (x ClrRedA400) Act()                    { x.ClrClr() }
func (x ClrRedA400) Ifc() interface{}        { return x.ClrClr() }
func (x ClrRedA400) ClrClr() clr.Clr         { return clr.RedA400 }
func (x ClrRedA700) Act()                    { x.ClrClr() }
func (x ClrRedA700) Ifc() interface{}        { return x.ClrClr() }
func (x ClrRedA700) ClrClr() clr.Clr         { return clr.RedA700 }
func (x ClrPink50) Act()                     { x.ClrClr() }
func (x ClrPink50) Ifc() interface{}         { return x.ClrClr() }
func (x ClrPink50) ClrClr() clr.Clr          { return clr.Pink50 }
func (x ClrPink100) Act()                    { x.ClrClr() }
func (x ClrPink100) Ifc() interface{}        { return x.ClrClr() }
func (x ClrPink100) ClrClr() clr.Clr         { return clr.Pink100 }
func (x ClrPink200) Act()                    { x.ClrClr() }
func (x ClrPink200) Ifc() interface{}        { return x.ClrClr() }
func (x ClrPink200) ClrClr() clr.Clr         { return clr.Pink200 }
func (x ClrPink300) Act()                    { x.ClrClr() }
func (x ClrPink300) Ifc() interface{}        { return x.ClrClr() }
func (x ClrPink300) ClrClr() clr.Clr         { return clr.Pink300 }
func (x ClrPink400) Act()                    { x.ClrClr() }
func (x ClrPink400) Ifc() interface{}        { return x.ClrClr() }
func (x ClrPink400) ClrClr() clr.Clr         { return clr.Pink400 }
func (x ClrPink500) Act()                    { x.ClrClr() }
func (x ClrPink500) Ifc() interface{}        { return x.ClrClr() }
func (x ClrPink500) ClrClr() clr.Clr         { return clr.Pink500 }
func (x ClrPink600) Act()                    { x.ClrClr() }
func (x ClrPink600) Ifc() interface{}        { return x.ClrClr() }
func (x ClrPink600) ClrClr() clr.Clr         { return clr.Pink600 }
func (x ClrPink700) Act()                    { x.ClrClr() }
func (x ClrPink700) Ifc() interface{}        { return x.ClrClr() }
func (x ClrPink700) ClrClr() clr.Clr         { return clr.Pink700 }
func (x ClrPink800) Act()                    { x.ClrClr() }
func (x ClrPink800) Ifc() interface{}        { return x.ClrClr() }
func (x ClrPink800) ClrClr() clr.Clr         { return clr.Pink800 }
func (x ClrPink900) Act()                    { x.ClrClr() }
func (x ClrPink900) Ifc() interface{}        { return x.ClrClr() }
func (x ClrPink900) ClrClr() clr.Clr         { return clr.Pink900 }
func (x ClrPinkA100) Act()                   { x.ClrClr() }
func (x ClrPinkA100) Ifc() interface{}       { return x.ClrClr() }
func (x ClrPinkA100) ClrClr() clr.Clr        { return clr.PinkA100 }
func (x ClrPinkA200) Act()                   { x.ClrClr() }
func (x ClrPinkA200) Ifc() interface{}       { return x.ClrClr() }
func (x ClrPinkA200) ClrClr() clr.Clr        { return clr.PinkA200 }
func (x ClrPinkA400) Act()                   { x.ClrClr() }
func (x ClrPinkA400) Ifc() interface{}       { return x.ClrClr() }
func (x ClrPinkA400) ClrClr() clr.Clr        { return clr.PinkA400 }
func (x ClrPinkA700) Act()                   { x.ClrClr() }
func (x ClrPinkA700) Ifc() interface{}       { return x.ClrClr() }
func (x ClrPinkA700) ClrClr() clr.Clr        { return clr.PinkA700 }
func (x ClrPurple50) Act()                   { x.ClrClr() }
func (x ClrPurple50) Ifc() interface{}       { return x.ClrClr() }
func (x ClrPurple50) ClrClr() clr.Clr        { return clr.Purple50 }
func (x ClrPurple100) Act()                  { x.ClrClr() }
func (x ClrPurple100) Ifc() interface{}      { return x.ClrClr() }
func (x ClrPurple100) ClrClr() clr.Clr       { return clr.Purple100 }
func (x ClrPurple200) Act()                  { x.ClrClr() }
func (x ClrPurple200) Ifc() interface{}      { return x.ClrClr() }
func (x ClrPurple200) ClrClr() clr.Clr       { return clr.Purple200 }
func (x ClrPurple300) Act()                  { x.ClrClr() }
func (x ClrPurple300) Ifc() interface{}      { return x.ClrClr() }
func (x ClrPurple300) ClrClr() clr.Clr       { return clr.Purple300 }
func (x ClrPurple400) Act()                  { x.ClrClr() }
func (x ClrPurple400) Ifc() interface{}      { return x.ClrClr() }
func (x ClrPurple400) ClrClr() clr.Clr       { return clr.Purple400 }
func (x ClrPurple500) Act()                  { x.ClrClr() }
func (x ClrPurple500) Ifc() interface{}      { return x.ClrClr() }
func (x ClrPurple500) ClrClr() clr.Clr       { return clr.Purple500 }
func (x ClrPurple600) Act()                  { x.ClrClr() }
func (x ClrPurple600) Ifc() interface{}      { return x.ClrClr() }
func (x ClrPurple600) ClrClr() clr.Clr       { return clr.Purple600 }
func (x ClrPurple700) Act()                  { x.ClrClr() }
func (x ClrPurple700) Ifc() interface{}      { return x.ClrClr() }
func (x ClrPurple700) ClrClr() clr.Clr       { return clr.Purple700 }
func (x ClrPurple800) Act()                  { x.ClrClr() }
func (x ClrPurple800) Ifc() interface{}      { return x.ClrClr() }
func (x ClrPurple800) ClrClr() clr.Clr       { return clr.Purple800 }
func (x ClrPurple900) Act()                  { x.ClrClr() }
func (x ClrPurple900) Ifc() interface{}      { return x.ClrClr() }
func (x ClrPurple900) ClrClr() clr.Clr       { return clr.Purple900 }
func (x ClrPurpleA100) Act()                 { x.ClrClr() }
func (x ClrPurpleA100) Ifc() interface{}     { return x.ClrClr() }
func (x ClrPurpleA100) ClrClr() clr.Clr      { return clr.PurpleA100 }
func (x ClrPurpleA200) Act()                 { x.ClrClr() }
func (x ClrPurpleA200) Ifc() interface{}     { return x.ClrClr() }
func (x ClrPurpleA200) ClrClr() clr.Clr      { return clr.PurpleA200 }
func (x ClrPurpleA400) Act()                 { x.ClrClr() }
func (x ClrPurpleA400) Ifc() interface{}     { return x.ClrClr() }
func (x ClrPurpleA400) ClrClr() clr.Clr      { return clr.PurpleA400 }
func (x ClrPurpleA700) Act()                 { x.ClrClr() }
func (x ClrPurpleA700) Ifc() interface{}     { return x.ClrClr() }
func (x ClrPurpleA700) ClrClr() clr.Clr      { return clr.PurpleA700 }
func (x ClrDeepPurple50) Act()               { x.ClrClr() }
func (x ClrDeepPurple50) Ifc() interface{}   { return x.ClrClr() }
func (x ClrDeepPurple50) ClrClr() clr.Clr    { return clr.DeepPurple50 }
func (x ClrDeepPurple100) Act()              { x.ClrClr() }
func (x ClrDeepPurple100) Ifc() interface{}  { return x.ClrClr() }
func (x ClrDeepPurple100) ClrClr() clr.Clr   { return clr.DeepPurple100 }
func (x ClrDeepPurple200) Act()              { x.ClrClr() }
func (x ClrDeepPurple200) Ifc() interface{}  { return x.ClrClr() }
func (x ClrDeepPurple200) ClrClr() clr.Clr   { return clr.DeepPurple200 }
func (x ClrDeepPurple300) Act()              { x.ClrClr() }
func (x ClrDeepPurple300) Ifc() interface{}  { return x.ClrClr() }
func (x ClrDeepPurple300) ClrClr() clr.Clr   { return clr.DeepPurple300 }
func (x ClrDeepPurple400) Act()              { x.ClrClr() }
func (x ClrDeepPurple400) Ifc() interface{}  { return x.ClrClr() }
func (x ClrDeepPurple400) ClrClr() clr.Clr   { return clr.DeepPurple400 }
func (x ClrDeepPurple500) Act()              { x.ClrClr() }
func (x ClrDeepPurple500) Ifc() interface{}  { return x.ClrClr() }
func (x ClrDeepPurple500) ClrClr() clr.Clr   { return clr.DeepPurple500 }
func (x ClrDeepPurple600) Act()              { x.ClrClr() }
func (x ClrDeepPurple600) Ifc() interface{}  { return x.ClrClr() }
func (x ClrDeepPurple600) ClrClr() clr.Clr   { return clr.DeepPurple600 }
func (x ClrDeepPurple700) Act()              { x.ClrClr() }
func (x ClrDeepPurple700) Ifc() interface{}  { return x.ClrClr() }
func (x ClrDeepPurple700) ClrClr() clr.Clr   { return clr.DeepPurple700 }
func (x ClrDeepPurple800) Act()              { x.ClrClr() }
func (x ClrDeepPurple800) Ifc() interface{}  { return x.ClrClr() }
func (x ClrDeepPurple800) ClrClr() clr.Clr   { return clr.DeepPurple800 }
func (x ClrDeepPurple900) Act()              { x.ClrClr() }
func (x ClrDeepPurple900) Ifc() interface{}  { return x.ClrClr() }
func (x ClrDeepPurple900) ClrClr() clr.Clr   { return clr.DeepPurple900 }
func (x ClrDeepPurpleA100) Act()             { x.ClrClr() }
func (x ClrDeepPurpleA100) Ifc() interface{} { return x.ClrClr() }
func (x ClrDeepPurpleA100) ClrClr() clr.Clr  { return clr.DeepPurpleA100 }
func (x ClrDeepPurpleA200) Act()             { x.ClrClr() }
func (x ClrDeepPurpleA200) Ifc() interface{} { return x.ClrClr() }
func (x ClrDeepPurpleA200) ClrClr() clr.Clr  { return clr.DeepPurpleA200 }
func (x ClrDeepPurpleA400) Act()             { x.ClrClr() }
func (x ClrDeepPurpleA400) Ifc() interface{} { return x.ClrClr() }
func (x ClrDeepPurpleA400) ClrClr() clr.Clr  { return clr.DeepPurpleA400 }
func (x ClrDeepPurpleA700) Act()             { x.ClrClr() }
func (x ClrDeepPurpleA700) Ifc() interface{} { return x.ClrClr() }
func (x ClrDeepPurpleA700) ClrClr() clr.Clr  { return clr.DeepPurpleA700 }
func (x ClrIndigo50) Act()                   { x.ClrClr() }
func (x ClrIndigo50) Ifc() interface{}       { return x.ClrClr() }
func (x ClrIndigo50) ClrClr() clr.Clr        { return clr.Indigo50 }
func (x ClrIndigo100) Act()                  { x.ClrClr() }
func (x ClrIndigo100) Ifc() interface{}      { return x.ClrClr() }
func (x ClrIndigo100) ClrClr() clr.Clr       { return clr.Indigo100 }
func (x ClrIndigo200) Act()                  { x.ClrClr() }
func (x ClrIndigo200) Ifc() interface{}      { return x.ClrClr() }
func (x ClrIndigo200) ClrClr() clr.Clr       { return clr.Indigo200 }
func (x ClrIndigo300) Act()                  { x.ClrClr() }
func (x ClrIndigo300) Ifc() interface{}      { return x.ClrClr() }
func (x ClrIndigo300) ClrClr() clr.Clr       { return clr.Indigo300 }
func (x ClrIndigo400) Act()                  { x.ClrClr() }
func (x ClrIndigo400) Ifc() interface{}      { return x.ClrClr() }
func (x ClrIndigo400) ClrClr() clr.Clr       { return clr.Indigo400 }
func (x ClrIndigo500) Act()                  { x.ClrClr() }
func (x ClrIndigo500) Ifc() interface{}      { return x.ClrClr() }
func (x ClrIndigo500) ClrClr() clr.Clr       { return clr.Indigo500 }
func (x ClrIndigo600) Act()                  { x.ClrClr() }
func (x ClrIndigo600) Ifc() interface{}      { return x.ClrClr() }
func (x ClrIndigo600) ClrClr() clr.Clr       { return clr.Indigo600 }
func (x ClrIndigo700) Act()                  { x.ClrClr() }
func (x ClrIndigo700) Ifc() interface{}      { return x.ClrClr() }
func (x ClrIndigo700) ClrClr() clr.Clr       { return clr.Indigo700 }
func (x ClrIndigo800) Act()                  { x.ClrClr() }
func (x ClrIndigo800) Ifc() interface{}      { return x.ClrClr() }
func (x ClrIndigo800) ClrClr() clr.Clr       { return clr.Indigo800 }
func (x ClrIndigo900) Act()                  { x.ClrClr() }
func (x ClrIndigo900) Ifc() interface{}      { return x.ClrClr() }
func (x ClrIndigo900) ClrClr() clr.Clr       { return clr.Indigo900 }
func (x ClrIndigoA100) Act()                 { x.ClrClr() }
func (x ClrIndigoA100) Ifc() interface{}     { return x.ClrClr() }
func (x ClrIndigoA100) ClrClr() clr.Clr      { return clr.IndigoA100 }
func (x ClrIndigoA200) Act()                 { x.ClrClr() }
func (x ClrIndigoA200) Ifc() interface{}     { return x.ClrClr() }
func (x ClrIndigoA200) ClrClr() clr.Clr      { return clr.IndigoA200 }
func (x ClrIndigoA400) Act()                 { x.ClrClr() }
func (x ClrIndigoA400) Ifc() interface{}     { return x.ClrClr() }
func (x ClrIndigoA400) ClrClr() clr.Clr      { return clr.IndigoA400 }
func (x ClrIndigoA700) Act()                 { x.ClrClr() }
func (x ClrIndigoA700) Ifc() interface{}     { return x.ClrClr() }
func (x ClrIndigoA700) ClrClr() clr.Clr      { return clr.IndigoA700 }
func (x ClrBlue50) Act()                     { x.ClrClr() }
func (x ClrBlue50) Ifc() interface{}         { return x.ClrClr() }
func (x ClrBlue50) ClrClr() clr.Clr          { return clr.Blue50 }
func (x ClrBlue100) Act()                    { x.ClrClr() }
func (x ClrBlue100) Ifc() interface{}        { return x.ClrClr() }
func (x ClrBlue100) ClrClr() clr.Clr         { return clr.Blue100 }
func (x ClrBlue200) Act()                    { x.ClrClr() }
func (x ClrBlue200) Ifc() interface{}        { return x.ClrClr() }
func (x ClrBlue200) ClrClr() clr.Clr         { return clr.Blue200 }
func (x ClrBlue300) Act()                    { x.ClrClr() }
func (x ClrBlue300) Ifc() interface{}        { return x.ClrClr() }
func (x ClrBlue300) ClrClr() clr.Clr         { return clr.Blue300 }
func (x ClrBlue400) Act()                    { x.ClrClr() }
func (x ClrBlue400) Ifc() interface{}        { return x.ClrClr() }
func (x ClrBlue400) ClrClr() clr.Clr         { return clr.Blue400 }
func (x ClrBlue500) Act()                    { x.ClrClr() }
func (x ClrBlue500) Ifc() interface{}        { return x.ClrClr() }
func (x ClrBlue500) ClrClr() clr.Clr         { return clr.Blue500 }
func (x ClrBlue600) Act()                    { x.ClrClr() }
func (x ClrBlue600) Ifc() interface{}        { return x.ClrClr() }
func (x ClrBlue600) ClrClr() clr.Clr         { return clr.Blue600 }
func (x ClrBlue700) Act()                    { x.ClrClr() }
func (x ClrBlue700) Ifc() interface{}        { return x.ClrClr() }
func (x ClrBlue700) ClrClr() clr.Clr         { return clr.Blue700 }
func (x ClrBlue800) Act()                    { x.ClrClr() }
func (x ClrBlue800) Ifc() interface{}        { return x.ClrClr() }
func (x ClrBlue800) ClrClr() clr.Clr         { return clr.Blue800 }
func (x ClrBlue900) Act()                    { x.ClrClr() }
func (x ClrBlue900) Ifc() interface{}        { return x.ClrClr() }
func (x ClrBlue900) ClrClr() clr.Clr         { return clr.Blue900 }
func (x ClrBlueA100) Act()                   { x.ClrClr() }
func (x ClrBlueA100) Ifc() interface{}       { return x.ClrClr() }
func (x ClrBlueA100) ClrClr() clr.Clr        { return clr.BlueA100 }
func (x ClrBlueA200) Act()                   { x.ClrClr() }
func (x ClrBlueA200) Ifc() interface{}       { return x.ClrClr() }
func (x ClrBlueA200) ClrClr() clr.Clr        { return clr.BlueA200 }
func (x ClrBlueA400) Act()                   { x.ClrClr() }
func (x ClrBlueA400) Ifc() interface{}       { return x.ClrClr() }
func (x ClrBlueA400) ClrClr() clr.Clr        { return clr.BlueA400 }
func (x ClrBlueA700) Act()                   { x.ClrClr() }
func (x ClrBlueA700) Ifc() interface{}       { return x.ClrClr() }
func (x ClrBlueA700) ClrClr() clr.Clr        { return clr.BlueA700 }
func (x ClrLightBlue50) Act()                { x.ClrClr() }
func (x ClrLightBlue50) Ifc() interface{}    { return x.ClrClr() }
func (x ClrLightBlue50) ClrClr() clr.Clr     { return clr.LightBlue50 }
func (x ClrLightBlue100) Act()               { x.ClrClr() }
func (x ClrLightBlue100) Ifc() interface{}   { return x.ClrClr() }
func (x ClrLightBlue100) ClrClr() clr.Clr    { return clr.LightBlue100 }
func (x ClrLightBlue200) Act()               { x.ClrClr() }
func (x ClrLightBlue200) Ifc() interface{}   { return x.ClrClr() }
func (x ClrLightBlue200) ClrClr() clr.Clr    { return clr.LightBlue200 }
func (x ClrLightBlue300) Act()               { x.ClrClr() }
func (x ClrLightBlue300) Ifc() interface{}   { return x.ClrClr() }
func (x ClrLightBlue300) ClrClr() clr.Clr    { return clr.LightBlue300 }
func (x ClrLightBlue400) Act()               { x.ClrClr() }
func (x ClrLightBlue400) Ifc() interface{}   { return x.ClrClr() }
func (x ClrLightBlue400) ClrClr() clr.Clr    { return clr.LightBlue400 }
func (x ClrLightBlue500) Act()               { x.ClrClr() }
func (x ClrLightBlue500) Ifc() interface{}   { return x.ClrClr() }
func (x ClrLightBlue500) ClrClr() clr.Clr    { return clr.LightBlue500 }
func (x ClrLightBlue600) Act()               { x.ClrClr() }
func (x ClrLightBlue600) Ifc() interface{}   { return x.ClrClr() }
func (x ClrLightBlue600) ClrClr() clr.Clr    { return clr.LightBlue600 }
func (x ClrLightBlue700) Act()               { x.ClrClr() }
func (x ClrLightBlue700) Ifc() interface{}   { return x.ClrClr() }
func (x ClrLightBlue700) ClrClr() clr.Clr    { return clr.LightBlue700 }
func (x ClrLightBlue800) Act()               { x.ClrClr() }
func (x ClrLightBlue800) Ifc() interface{}   { return x.ClrClr() }
func (x ClrLightBlue800) ClrClr() clr.Clr    { return clr.LightBlue800 }
func (x ClrLightBlue900) Act()               { x.ClrClr() }
func (x ClrLightBlue900) Ifc() interface{}   { return x.ClrClr() }
func (x ClrLightBlue900) ClrClr() clr.Clr    { return clr.LightBlue900 }
func (x ClrLightBlueA100) Act()              { x.ClrClr() }
func (x ClrLightBlueA100) Ifc() interface{}  { return x.ClrClr() }
func (x ClrLightBlueA100) ClrClr() clr.Clr   { return clr.LightBlueA100 }
func (x ClrLightBlueA200) Act()              { x.ClrClr() }
func (x ClrLightBlueA200) Ifc() interface{}  { return x.ClrClr() }
func (x ClrLightBlueA200) ClrClr() clr.Clr   { return clr.LightBlueA200 }
func (x ClrLightBlueA400) Act()              { x.ClrClr() }
func (x ClrLightBlueA400) Ifc() interface{}  { return x.ClrClr() }
func (x ClrLightBlueA400) ClrClr() clr.Clr   { return clr.LightBlueA400 }
func (x ClrLightBlueA700) Act()              { x.ClrClr() }
func (x ClrLightBlueA700) Ifc() interface{}  { return x.ClrClr() }
func (x ClrLightBlueA700) ClrClr() clr.Clr   { return clr.LightBlueA700 }
func (x ClrCyan50) Act()                     { x.ClrClr() }
func (x ClrCyan50) Ifc() interface{}         { return x.ClrClr() }
func (x ClrCyan50) ClrClr() clr.Clr          { return clr.Cyan50 }
func (x ClrCyan100) Act()                    { x.ClrClr() }
func (x ClrCyan100) Ifc() interface{}        { return x.ClrClr() }
func (x ClrCyan100) ClrClr() clr.Clr         { return clr.Cyan100 }
func (x ClrCyan200) Act()                    { x.ClrClr() }
func (x ClrCyan200) Ifc() interface{}        { return x.ClrClr() }
func (x ClrCyan200) ClrClr() clr.Clr         { return clr.Cyan200 }
func (x ClrCyan300) Act()                    { x.ClrClr() }
func (x ClrCyan300) Ifc() interface{}        { return x.ClrClr() }
func (x ClrCyan300) ClrClr() clr.Clr         { return clr.Cyan300 }
func (x ClrCyan400) Act()                    { x.ClrClr() }
func (x ClrCyan400) Ifc() interface{}        { return x.ClrClr() }
func (x ClrCyan400) ClrClr() clr.Clr         { return clr.Cyan400 }
func (x ClrCyan500) Act()                    { x.ClrClr() }
func (x ClrCyan500) Ifc() interface{}        { return x.ClrClr() }
func (x ClrCyan500) ClrClr() clr.Clr         { return clr.Cyan500 }
func (x ClrCyan600) Act()                    { x.ClrClr() }
func (x ClrCyan600) Ifc() interface{}        { return x.ClrClr() }
func (x ClrCyan600) ClrClr() clr.Clr         { return clr.Cyan600 }
func (x ClrCyan700) Act()                    { x.ClrClr() }
func (x ClrCyan700) Ifc() interface{}        { return x.ClrClr() }
func (x ClrCyan700) ClrClr() clr.Clr         { return clr.Cyan700 }
func (x ClrCyan800) Act()                    { x.ClrClr() }
func (x ClrCyan800) Ifc() interface{}        { return x.ClrClr() }
func (x ClrCyan800) ClrClr() clr.Clr         { return clr.Cyan800 }
func (x ClrCyan900) Act()                    { x.ClrClr() }
func (x ClrCyan900) Ifc() interface{}        { return x.ClrClr() }
func (x ClrCyan900) ClrClr() clr.Clr         { return clr.Cyan900 }
func (x ClrCyanA100) Act()                   { x.ClrClr() }
func (x ClrCyanA100) Ifc() interface{}       { return x.ClrClr() }
func (x ClrCyanA100) ClrClr() clr.Clr        { return clr.CyanA100 }
func (x ClrCyanA200) Act()                   { x.ClrClr() }
func (x ClrCyanA200) Ifc() interface{}       { return x.ClrClr() }
func (x ClrCyanA200) ClrClr() clr.Clr        { return clr.CyanA200 }
func (x ClrCyanA400) Act()                   { x.ClrClr() }
func (x ClrCyanA400) Ifc() interface{}       { return x.ClrClr() }
func (x ClrCyanA400) ClrClr() clr.Clr        { return clr.CyanA400 }
func (x ClrCyanA700) Act()                   { x.ClrClr() }
func (x ClrCyanA700) Ifc() interface{}       { return x.ClrClr() }
func (x ClrCyanA700) ClrClr() clr.Clr        { return clr.CyanA700 }
func (x ClrTeal50) Act()                     { x.ClrClr() }
func (x ClrTeal50) Ifc() interface{}         { return x.ClrClr() }
func (x ClrTeal50) ClrClr() clr.Clr          { return clr.Teal50 }
func (x ClrTeal100) Act()                    { x.ClrClr() }
func (x ClrTeal100) Ifc() interface{}        { return x.ClrClr() }
func (x ClrTeal100) ClrClr() clr.Clr         { return clr.Teal100 }
func (x ClrTeal200) Act()                    { x.ClrClr() }
func (x ClrTeal200) Ifc() interface{}        { return x.ClrClr() }
func (x ClrTeal200) ClrClr() clr.Clr         { return clr.Teal200 }
func (x ClrTeal300) Act()                    { x.ClrClr() }
func (x ClrTeal300) Ifc() interface{}        { return x.ClrClr() }
func (x ClrTeal300) ClrClr() clr.Clr         { return clr.Teal300 }
func (x ClrTeal400) Act()                    { x.ClrClr() }
func (x ClrTeal400) Ifc() interface{}        { return x.ClrClr() }
func (x ClrTeal400) ClrClr() clr.Clr         { return clr.Teal400 }
func (x ClrTeal500) Act()                    { x.ClrClr() }
func (x ClrTeal500) Ifc() interface{}        { return x.ClrClr() }
func (x ClrTeal500) ClrClr() clr.Clr         { return clr.Teal500 }
func (x ClrTeal600) Act()                    { x.ClrClr() }
func (x ClrTeal600) Ifc() interface{}        { return x.ClrClr() }
func (x ClrTeal600) ClrClr() clr.Clr         { return clr.Teal600 }
func (x ClrTeal700) Act()                    { x.ClrClr() }
func (x ClrTeal700) Ifc() interface{}        { return x.ClrClr() }
func (x ClrTeal700) ClrClr() clr.Clr         { return clr.Teal700 }
func (x ClrTeal800) Act()                    { x.ClrClr() }
func (x ClrTeal800) Ifc() interface{}        { return x.ClrClr() }
func (x ClrTeal800) ClrClr() clr.Clr         { return clr.Teal800 }
func (x ClrTeal900) Act()                    { x.ClrClr() }
func (x ClrTeal900) Ifc() interface{}        { return x.ClrClr() }
func (x ClrTeal900) ClrClr() clr.Clr         { return clr.Teal900 }
func (x ClrTealA100) Act()                   { x.ClrClr() }
func (x ClrTealA100) Ifc() interface{}       { return x.ClrClr() }
func (x ClrTealA100) ClrClr() clr.Clr        { return clr.TealA100 }
func (x ClrTealA200) Act()                   { x.ClrClr() }
func (x ClrTealA200) Ifc() interface{}       { return x.ClrClr() }
func (x ClrTealA200) ClrClr() clr.Clr        { return clr.TealA200 }
func (x ClrTealA400) Act()                   { x.ClrClr() }
func (x ClrTealA400) Ifc() interface{}       { return x.ClrClr() }
func (x ClrTealA400) ClrClr() clr.Clr        { return clr.TealA400 }
func (x ClrTealA700) Act()                   { x.ClrClr() }
func (x ClrTealA700) Ifc() interface{}       { return x.ClrClr() }
func (x ClrTealA700) ClrClr() clr.Clr        { return clr.TealA700 }
func (x ClrGreen50) Act()                    { x.ClrClr() }
func (x ClrGreen50) Ifc() interface{}        { return x.ClrClr() }
func (x ClrGreen50) ClrClr() clr.Clr         { return clr.Green50 }
func (x ClrGreen100) Act()                   { x.ClrClr() }
func (x ClrGreen100) Ifc() interface{}       { return x.ClrClr() }
func (x ClrGreen100) ClrClr() clr.Clr        { return clr.Green100 }
func (x ClrGreen200) Act()                   { x.ClrClr() }
func (x ClrGreen200) Ifc() interface{}       { return x.ClrClr() }
func (x ClrGreen200) ClrClr() clr.Clr        { return clr.Green200 }
func (x ClrGreen300) Act()                   { x.ClrClr() }
func (x ClrGreen300) Ifc() interface{}       { return x.ClrClr() }
func (x ClrGreen300) ClrClr() clr.Clr        { return clr.Green300 }
func (x ClrGreen400) Act()                   { x.ClrClr() }
func (x ClrGreen400) Ifc() interface{}       { return x.ClrClr() }
func (x ClrGreen400) ClrClr() clr.Clr        { return clr.Green400 }
func (x ClrGreen500) Act()                   { x.ClrClr() }
func (x ClrGreen500) Ifc() interface{}       { return x.ClrClr() }
func (x ClrGreen500) ClrClr() clr.Clr        { return clr.Green500 }
func (x ClrGreen600) Act()                   { x.ClrClr() }
func (x ClrGreen600) Ifc() interface{}       { return x.ClrClr() }
func (x ClrGreen600) ClrClr() clr.Clr        { return clr.Green600 }
func (x ClrGreen700) Act()                   { x.ClrClr() }
func (x ClrGreen700) Ifc() interface{}       { return x.ClrClr() }
func (x ClrGreen700) ClrClr() clr.Clr        { return clr.Green700 }
func (x ClrGreen800) Act()                   { x.ClrClr() }
func (x ClrGreen800) Ifc() interface{}       { return x.ClrClr() }
func (x ClrGreen800) ClrClr() clr.Clr        { return clr.Green800 }
func (x ClrGreen900) Act()                   { x.ClrClr() }
func (x ClrGreen900) Ifc() interface{}       { return x.ClrClr() }
func (x ClrGreen900) ClrClr() clr.Clr        { return clr.Green900 }
func (x ClrGreenA100) Act()                  { x.ClrClr() }
func (x ClrGreenA100) Ifc() interface{}      { return x.ClrClr() }
func (x ClrGreenA100) ClrClr() clr.Clr       { return clr.GreenA100 }
func (x ClrGreenA200) Act()                  { x.ClrClr() }
func (x ClrGreenA200) Ifc() interface{}      { return x.ClrClr() }
func (x ClrGreenA200) ClrClr() clr.Clr       { return clr.GreenA200 }
func (x ClrGreenA400) Act()                  { x.ClrClr() }
func (x ClrGreenA400) Ifc() interface{}      { return x.ClrClr() }
func (x ClrGreenA400) ClrClr() clr.Clr       { return clr.GreenA400 }
func (x ClrGreenA700) Act()                  { x.ClrClr() }
func (x ClrGreenA700) Ifc() interface{}      { return x.ClrClr() }
func (x ClrGreenA700) ClrClr() clr.Clr       { return clr.GreenA700 }
func (x ClrLightGreen50) Act()               { x.ClrClr() }
func (x ClrLightGreen50) Ifc() interface{}   { return x.ClrClr() }
func (x ClrLightGreen50) ClrClr() clr.Clr    { return clr.LightGreen50 }
func (x ClrLightGreen100) Act()              { x.ClrClr() }
func (x ClrLightGreen100) Ifc() interface{}  { return x.ClrClr() }
func (x ClrLightGreen100) ClrClr() clr.Clr   { return clr.LightGreen100 }
func (x ClrLightGreen200) Act()              { x.ClrClr() }
func (x ClrLightGreen200) Ifc() interface{}  { return x.ClrClr() }
func (x ClrLightGreen200) ClrClr() clr.Clr   { return clr.LightGreen200 }
func (x ClrLightGreen300) Act()              { x.ClrClr() }
func (x ClrLightGreen300) Ifc() interface{}  { return x.ClrClr() }
func (x ClrLightGreen300) ClrClr() clr.Clr   { return clr.LightGreen300 }
func (x ClrLightGreen400) Act()              { x.ClrClr() }
func (x ClrLightGreen400) Ifc() interface{}  { return x.ClrClr() }
func (x ClrLightGreen400) ClrClr() clr.Clr   { return clr.LightGreen400 }
func (x ClrLightGreen500) Act()              { x.ClrClr() }
func (x ClrLightGreen500) Ifc() interface{}  { return x.ClrClr() }
func (x ClrLightGreen500) ClrClr() clr.Clr   { return clr.LightGreen500 }
func (x ClrLightGreen600) Act()              { x.ClrClr() }
func (x ClrLightGreen600) Ifc() interface{}  { return x.ClrClr() }
func (x ClrLightGreen600) ClrClr() clr.Clr   { return clr.LightGreen600 }
func (x ClrLightGreen700) Act()              { x.ClrClr() }
func (x ClrLightGreen700) Ifc() interface{}  { return x.ClrClr() }
func (x ClrLightGreen700) ClrClr() clr.Clr   { return clr.LightGreen700 }
func (x ClrLightGreen800) Act()              { x.ClrClr() }
func (x ClrLightGreen800) Ifc() interface{}  { return x.ClrClr() }
func (x ClrLightGreen800) ClrClr() clr.Clr   { return clr.LightGreen800 }
func (x ClrLightGreen900) Act()              { x.ClrClr() }
func (x ClrLightGreen900) Ifc() interface{}  { return x.ClrClr() }
func (x ClrLightGreen900) ClrClr() clr.Clr   { return clr.LightGreen900 }
func (x ClrLightGreenA100) Act()             { x.ClrClr() }
func (x ClrLightGreenA100) Ifc() interface{} { return x.ClrClr() }
func (x ClrLightGreenA100) ClrClr() clr.Clr  { return clr.LightGreenA100 }
func (x ClrLightGreenA200) Act()             { x.ClrClr() }
func (x ClrLightGreenA200) Ifc() interface{} { return x.ClrClr() }
func (x ClrLightGreenA200) ClrClr() clr.Clr  { return clr.LightGreenA200 }
func (x ClrLightGreenA400) Act()             { x.ClrClr() }
func (x ClrLightGreenA400) Ifc() interface{} { return x.ClrClr() }
func (x ClrLightGreenA400) ClrClr() clr.Clr  { return clr.LightGreenA400 }
func (x ClrLightGreenA700) Act()             { x.ClrClr() }
func (x ClrLightGreenA700) Ifc() interface{} { return x.ClrClr() }
func (x ClrLightGreenA700) ClrClr() clr.Clr  { return clr.LightGreenA700 }
func (x ClrLime50) Act()                     { x.ClrClr() }
func (x ClrLime50) Ifc() interface{}         { return x.ClrClr() }
func (x ClrLime50) ClrClr() clr.Clr          { return clr.Lime50 }
func (x ClrLime100) Act()                    { x.ClrClr() }
func (x ClrLime100) Ifc() interface{}        { return x.ClrClr() }
func (x ClrLime100) ClrClr() clr.Clr         { return clr.Lime100 }
func (x ClrLime200) Act()                    { x.ClrClr() }
func (x ClrLime200) Ifc() interface{}        { return x.ClrClr() }
func (x ClrLime200) ClrClr() clr.Clr         { return clr.Lime200 }
func (x ClrLime300) Act()                    { x.ClrClr() }
func (x ClrLime300) Ifc() interface{}        { return x.ClrClr() }
func (x ClrLime300) ClrClr() clr.Clr         { return clr.Lime300 }
func (x ClrLime400) Act()                    { x.ClrClr() }
func (x ClrLime400) Ifc() interface{}        { return x.ClrClr() }
func (x ClrLime400) ClrClr() clr.Clr         { return clr.Lime400 }
func (x ClrLime500) Act()                    { x.ClrClr() }
func (x ClrLime500) Ifc() interface{}        { return x.ClrClr() }
func (x ClrLime500) ClrClr() clr.Clr         { return clr.Lime500 }
func (x ClrLime600) Act()                    { x.ClrClr() }
func (x ClrLime600) Ifc() interface{}        { return x.ClrClr() }
func (x ClrLime600) ClrClr() clr.Clr         { return clr.Lime600 }
func (x ClrLime700) Act()                    { x.ClrClr() }
func (x ClrLime700) Ifc() interface{}        { return x.ClrClr() }
func (x ClrLime700) ClrClr() clr.Clr         { return clr.Lime700 }
func (x ClrLime800) Act()                    { x.ClrClr() }
func (x ClrLime800) Ifc() interface{}        { return x.ClrClr() }
func (x ClrLime800) ClrClr() clr.Clr         { return clr.Lime800 }
func (x ClrLime900) Act()                    { x.ClrClr() }
func (x ClrLime900) Ifc() interface{}        { return x.ClrClr() }
func (x ClrLime900) ClrClr() clr.Clr         { return clr.Lime900 }
func (x ClrLimeA100) Act()                   { x.ClrClr() }
func (x ClrLimeA100) Ifc() interface{}       { return x.ClrClr() }
func (x ClrLimeA100) ClrClr() clr.Clr        { return clr.LimeA100 }
func (x ClrLimeA200) Act()                   { x.ClrClr() }
func (x ClrLimeA200) Ifc() interface{}       { return x.ClrClr() }
func (x ClrLimeA200) ClrClr() clr.Clr        { return clr.LimeA200 }
func (x ClrLimeA400) Act()                   { x.ClrClr() }
func (x ClrLimeA400) Ifc() interface{}       { return x.ClrClr() }
func (x ClrLimeA400) ClrClr() clr.Clr        { return clr.LimeA400 }
func (x ClrLimeA700) Act()                   { x.ClrClr() }
func (x ClrLimeA700) Ifc() interface{}       { return x.ClrClr() }
func (x ClrLimeA700) ClrClr() clr.Clr        { return clr.LimeA700 }
func (x ClrYellow50) Act()                   { x.ClrClr() }
func (x ClrYellow50) Ifc() interface{}       { return x.ClrClr() }
func (x ClrYellow50) ClrClr() clr.Clr        { return clr.Yellow50 }
func (x ClrYellow100) Act()                  { x.ClrClr() }
func (x ClrYellow100) Ifc() interface{}      { return x.ClrClr() }
func (x ClrYellow100) ClrClr() clr.Clr       { return clr.Yellow100 }
func (x ClrYellow200) Act()                  { x.ClrClr() }
func (x ClrYellow200) Ifc() interface{}      { return x.ClrClr() }
func (x ClrYellow200) ClrClr() clr.Clr       { return clr.Yellow200 }
func (x ClrYellow300) Act()                  { x.ClrClr() }
func (x ClrYellow300) Ifc() interface{}      { return x.ClrClr() }
func (x ClrYellow300) ClrClr() clr.Clr       { return clr.Yellow300 }
func (x ClrYellow400) Act()                  { x.ClrClr() }
func (x ClrYellow400) Ifc() interface{}      { return x.ClrClr() }
func (x ClrYellow400) ClrClr() clr.Clr       { return clr.Yellow400 }
func (x ClrYellow500) Act()                  { x.ClrClr() }
func (x ClrYellow500) Ifc() interface{}      { return x.ClrClr() }
func (x ClrYellow500) ClrClr() clr.Clr       { return clr.Yellow500 }
func (x ClrYellow600) Act()                  { x.ClrClr() }
func (x ClrYellow600) Ifc() interface{}      { return x.ClrClr() }
func (x ClrYellow600) ClrClr() clr.Clr       { return clr.Yellow600 }
func (x ClrYellow700) Act()                  { x.ClrClr() }
func (x ClrYellow700) Ifc() interface{}      { return x.ClrClr() }
func (x ClrYellow700) ClrClr() clr.Clr       { return clr.Yellow700 }
func (x ClrYellow800) Act()                  { x.ClrClr() }
func (x ClrYellow800) Ifc() interface{}      { return x.ClrClr() }
func (x ClrYellow800) ClrClr() clr.Clr       { return clr.Yellow800 }
func (x ClrYellow900) Act()                  { x.ClrClr() }
func (x ClrYellow900) Ifc() interface{}      { return x.ClrClr() }
func (x ClrYellow900) ClrClr() clr.Clr       { return clr.Yellow900 }
func (x ClrYellowA100) Act()                 { x.ClrClr() }
func (x ClrYellowA100) Ifc() interface{}     { return x.ClrClr() }
func (x ClrYellowA100) ClrClr() clr.Clr      { return clr.YellowA100 }
func (x ClrYellowA200) Act()                 { x.ClrClr() }
func (x ClrYellowA200) Ifc() interface{}     { return x.ClrClr() }
func (x ClrYellowA200) ClrClr() clr.Clr      { return clr.YellowA200 }
func (x ClrYellowA400) Act()                 { x.ClrClr() }
func (x ClrYellowA400) Ifc() interface{}     { return x.ClrClr() }
func (x ClrYellowA400) ClrClr() clr.Clr      { return clr.YellowA400 }
func (x ClrYellowA700) Act()                 { x.ClrClr() }
func (x ClrYellowA700) Ifc() interface{}     { return x.ClrClr() }
func (x ClrYellowA700) ClrClr() clr.Clr      { return clr.YellowA700 }
func (x ClrAmber50) Act()                    { x.ClrClr() }
func (x ClrAmber50) Ifc() interface{}        { return x.ClrClr() }
func (x ClrAmber50) ClrClr() clr.Clr         { return clr.Amber50 }
func (x ClrAmber100) Act()                   { x.ClrClr() }
func (x ClrAmber100) Ifc() interface{}       { return x.ClrClr() }
func (x ClrAmber100) ClrClr() clr.Clr        { return clr.Amber100 }
func (x ClrAmber200) Act()                   { x.ClrClr() }
func (x ClrAmber200) Ifc() interface{}       { return x.ClrClr() }
func (x ClrAmber200) ClrClr() clr.Clr        { return clr.Amber200 }
func (x ClrAmber300) Act()                   { x.ClrClr() }
func (x ClrAmber300) Ifc() interface{}       { return x.ClrClr() }
func (x ClrAmber300) ClrClr() clr.Clr        { return clr.Amber300 }
func (x ClrAmber400) Act()                   { x.ClrClr() }
func (x ClrAmber400) Ifc() interface{}       { return x.ClrClr() }
func (x ClrAmber400) ClrClr() clr.Clr        { return clr.Amber400 }
func (x ClrAmber500) Act()                   { x.ClrClr() }
func (x ClrAmber500) Ifc() interface{}       { return x.ClrClr() }
func (x ClrAmber500) ClrClr() clr.Clr        { return clr.Amber500 }
func (x ClrAmber600) Act()                   { x.ClrClr() }
func (x ClrAmber600) Ifc() interface{}       { return x.ClrClr() }
func (x ClrAmber600) ClrClr() clr.Clr        { return clr.Amber600 }
func (x ClrAmber700) Act()                   { x.ClrClr() }
func (x ClrAmber700) Ifc() interface{}       { return x.ClrClr() }
func (x ClrAmber700) ClrClr() clr.Clr        { return clr.Amber700 }
func (x ClrAmber800) Act()                   { x.ClrClr() }
func (x ClrAmber800) Ifc() interface{}       { return x.ClrClr() }
func (x ClrAmber800) ClrClr() clr.Clr        { return clr.Amber800 }
func (x ClrAmber900) Act()                   { x.ClrClr() }
func (x ClrAmber900) Ifc() interface{}       { return x.ClrClr() }
func (x ClrAmber900) ClrClr() clr.Clr        { return clr.Amber900 }
func (x ClrAmberA100) Act()                  { x.ClrClr() }
func (x ClrAmberA100) Ifc() interface{}      { return x.ClrClr() }
func (x ClrAmberA100) ClrClr() clr.Clr       { return clr.AmberA100 }
func (x ClrAmberA200) Act()                  { x.ClrClr() }
func (x ClrAmberA200) Ifc() interface{}      { return x.ClrClr() }
func (x ClrAmberA200) ClrClr() clr.Clr       { return clr.AmberA200 }
func (x ClrAmberA400) Act()                  { x.ClrClr() }
func (x ClrAmberA400) Ifc() interface{}      { return x.ClrClr() }
func (x ClrAmberA400) ClrClr() clr.Clr       { return clr.AmberA400 }
func (x ClrAmberA700) Act()                  { x.ClrClr() }
func (x ClrAmberA700) Ifc() interface{}      { return x.ClrClr() }
func (x ClrAmberA700) ClrClr() clr.Clr       { return clr.AmberA700 }
func (x ClrOrange50) Act()                   { x.ClrClr() }
func (x ClrOrange50) Ifc() interface{}       { return x.ClrClr() }
func (x ClrOrange50) ClrClr() clr.Clr        { return clr.Orange50 }
func (x ClrOrange100) Act()                  { x.ClrClr() }
func (x ClrOrange100) Ifc() interface{}      { return x.ClrClr() }
func (x ClrOrange100) ClrClr() clr.Clr       { return clr.Orange100 }
func (x ClrOrange200) Act()                  { x.ClrClr() }
func (x ClrOrange200) Ifc() interface{}      { return x.ClrClr() }
func (x ClrOrange200) ClrClr() clr.Clr       { return clr.Orange200 }
func (x ClrOrange300) Act()                  { x.ClrClr() }
func (x ClrOrange300) Ifc() interface{}      { return x.ClrClr() }
func (x ClrOrange300) ClrClr() clr.Clr       { return clr.Orange300 }
func (x ClrOrange400) Act()                  { x.ClrClr() }
func (x ClrOrange400) Ifc() interface{}      { return x.ClrClr() }
func (x ClrOrange400) ClrClr() clr.Clr       { return clr.Orange400 }
func (x ClrOrange500) Act()                  { x.ClrClr() }
func (x ClrOrange500) Ifc() interface{}      { return x.ClrClr() }
func (x ClrOrange500) ClrClr() clr.Clr       { return clr.Orange500 }
func (x ClrOrange600) Act()                  { x.ClrClr() }
func (x ClrOrange600) Ifc() interface{}      { return x.ClrClr() }
func (x ClrOrange600) ClrClr() clr.Clr       { return clr.Orange600 }
func (x ClrOrange700) Act()                  { x.ClrClr() }
func (x ClrOrange700) Ifc() interface{}      { return x.ClrClr() }
func (x ClrOrange700) ClrClr() clr.Clr       { return clr.Orange700 }
func (x ClrOrange800) Act()                  { x.ClrClr() }
func (x ClrOrange800) Ifc() interface{}      { return x.ClrClr() }
func (x ClrOrange800) ClrClr() clr.Clr       { return clr.Orange800 }
func (x ClrOrange900) Act()                  { x.ClrClr() }
func (x ClrOrange900) Ifc() interface{}      { return x.ClrClr() }
func (x ClrOrange900) ClrClr() clr.Clr       { return clr.Orange900 }
func (x ClrOrangeA100) Act()                 { x.ClrClr() }
func (x ClrOrangeA100) Ifc() interface{}     { return x.ClrClr() }
func (x ClrOrangeA100) ClrClr() clr.Clr      { return clr.OrangeA100 }
func (x ClrOrangeA200) Act()                 { x.ClrClr() }
func (x ClrOrangeA200) Ifc() interface{}     { return x.ClrClr() }
func (x ClrOrangeA200) ClrClr() clr.Clr      { return clr.OrangeA200 }
func (x ClrOrangeA400) Act()                 { x.ClrClr() }
func (x ClrOrangeA400) Ifc() interface{}     { return x.ClrClr() }
func (x ClrOrangeA400) ClrClr() clr.Clr      { return clr.OrangeA400 }
func (x ClrOrangeA700) Act()                 { x.ClrClr() }
func (x ClrOrangeA700) Ifc() interface{}     { return x.ClrClr() }
func (x ClrOrangeA700) ClrClr() clr.Clr      { return clr.OrangeA700 }
func (x ClrDeepOrange50) Act()               { x.ClrClr() }
func (x ClrDeepOrange50) Ifc() interface{}   { return x.ClrClr() }
func (x ClrDeepOrange50) ClrClr() clr.Clr    { return clr.DeepOrange50 }
func (x ClrDeepOrange100) Act()              { x.ClrClr() }
func (x ClrDeepOrange100) Ifc() interface{}  { return x.ClrClr() }
func (x ClrDeepOrange100) ClrClr() clr.Clr   { return clr.DeepOrange100 }
func (x ClrDeepOrange200) Act()              { x.ClrClr() }
func (x ClrDeepOrange200) Ifc() interface{}  { return x.ClrClr() }
func (x ClrDeepOrange200) ClrClr() clr.Clr   { return clr.DeepOrange200 }
func (x ClrDeepOrange300) Act()              { x.ClrClr() }
func (x ClrDeepOrange300) Ifc() interface{}  { return x.ClrClr() }
func (x ClrDeepOrange300) ClrClr() clr.Clr   { return clr.DeepOrange300 }
func (x ClrDeepOrange400) Act()              { x.ClrClr() }
func (x ClrDeepOrange400) Ifc() interface{}  { return x.ClrClr() }
func (x ClrDeepOrange400) ClrClr() clr.Clr   { return clr.DeepOrange400 }
func (x ClrDeepOrange500) Act()              { x.ClrClr() }
func (x ClrDeepOrange500) Ifc() interface{}  { return x.ClrClr() }
func (x ClrDeepOrange500) ClrClr() clr.Clr   { return clr.DeepOrange500 }
func (x ClrDeepOrange600) Act()              { x.ClrClr() }
func (x ClrDeepOrange600) Ifc() interface{}  { return x.ClrClr() }
func (x ClrDeepOrange600) ClrClr() clr.Clr   { return clr.DeepOrange600 }
func (x ClrDeepOrange700) Act()              { x.ClrClr() }
func (x ClrDeepOrange700) Ifc() interface{}  { return x.ClrClr() }
func (x ClrDeepOrange700) ClrClr() clr.Clr   { return clr.DeepOrange700 }
func (x ClrDeepOrange800) Act()              { x.ClrClr() }
func (x ClrDeepOrange800) Ifc() interface{}  { return x.ClrClr() }
func (x ClrDeepOrange800) ClrClr() clr.Clr   { return clr.DeepOrange800 }
func (x ClrDeepOrange900) Act()              { x.ClrClr() }
func (x ClrDeepOrange900) Ifc() interface{}  { return x.ClrClr() }
func (x ClrDeepOrange900) ClrClr() clr.Clr   { return clr.DeepOrange900 }
func (x ClrDeepOrangeA100) Act()             { x.ClrClr() }
func (x ClrDeepOrangeA100) Ifc() interface{} { return x.ClrClr() }
func (x ClrDeepOrangeA100) ClrClr() clr.Clr  { return clr.DeepOrangeA100 }
func (x ClrDeepOrangeA200) Act()             { x.ClrClr() }
func (x ClrDeepOrangeA200) Ifc() interface{} { return x.ClrClr() }
func (x ClrDeepOrangeA200) ClrClr() clr.Clr  { return clr.DeepOrangeA200 }
func (x ClrDeepOrangeA400) Act()             { x.ClrClr() }
func (x ClrDeepOrangeA400) Ifc() interface{} { return x.ClrClr() }
func (x ClrDeepOrangeA400) ClrClr() clr.Clr  { return clr.DeepOrangeA400 }
func (x ClrDeepOrangeA700) Act()             { x.ClrClr() }
func (x ClrDeepOrangeA700) Ifc() interface{} { return x.ClrClr() }
func (x ClrDeepOrangeA700) ClrClr() clr.Clr  { return clr.DeepOrangeA700 }
func (x ClrBrown50) Act()                    { x.ClrClr() }
func (x ClrBrown50) Ifc() interface{}        { return x.ClrClr() }
func (x ClrBrown50) ClrClr() clr.Clr         { return clr.Brown50 }
func (x ClrBrown100) Act()                   { x.ClrClr() }
func (x ClrBrown100) Ifc() interface{}       { return x.ClrClr() }
func (x ClrBrown100) ClrClr() clr.Clr        { return clr.Brown100 }
func (x ClrBrown200) Act()                   { x.ClrClr() }
func (x ClrBrown200) Ifc() interface{}       { return x.ClrClr() }
func (x ClrBrown200) ClrClr() clr.Clr        { return clr.Brown200 }
func (x ClrBrown300) Act()                   { x.ClrClr() }
func (x ClrBrown300) Ifc() interface{}       { return x.ClrClr() }
func (x ClrBrown300) ClrClr() clr.Clr        { return clr.Brown300 }
func (x ClrBrown400) Act()                   { x.ClrClr() }
func (x ClrBrown400) Ifc() interface{}       { return x.ClrClr() }
func (x ClrBrown400) ClrClr() clr.Clr        { return clr.Brown400 }
func (x ClrBrown500) Act()                   { x.ClrClr() }
func (x ClrBrown500) Ifc() interface{}       { return x.ClrClr() }
func (x ClrBrown500) ClrClr() clr.Clr        { return clr.Brown500 }
func (x ClrBrown600) Act()                   { x.ClrClr() }
func (x ClrBrown600) Ifc() interface{}       { return x.ClrClr() }
func (x ClrBrown600) ClrClr() clr.Clr        { return clr.Brown600 }
func (x ClrBrown700) Act()                   { x.ClrClr() }
func (x ClrBrown700) Ifc() interface{}       { return x.ClrClr() }
func (x ClrBrown700) ClrClr() clr.Clr        { return clr.Brown700 }
func (x ClrBrown800) Act()                   { x.ClrClr() }
func (x ClrBrown800) Ifc() interface{}       { return x.ClrClr() }
func (x ClrBrown800) ClrClr() clr.Clr        { return clr.Brown800 }
func (x ClrBrown900) Act()                   { x.ClrClr() }
func (x ClrBrown900) Ifc() interface{}       { return x.ClrClr() }
func (x ClrBrown900) ClrClr() clr.Clr        { return clr.Brown900 }
func (x ClrGrey50) Act()                     { x.ClrClr() }
func (x ClrGrey50) Ifc() interface{}         { return x.ClrClr() }
func (x ClrGrey50) ClrClr() clr.Clr          { return clr.Grey50 }
func (x ClrGrey100) Act()                    { x.ClrClr() }
func (x ClrGrey100) Ifc() interface{}        { return x.ClrClr() }
func (x ClrGrey100) ClrClr() clr.Clr         { return clr.Grey100 }
func (x ClrGrey200) Act()                    { x.ClrClr() }
func (x ClrGrey200) Ifc() interface{}        { return x.ClrClr() }
func (x ClrGrey200) ClrClr() clr.Clr         { return clr.Grey200 }
func (x ClrGrey300) Act()                    { x.ClrClr() }
func (x ClrGrey300) Ifc() interface{}        { return x.ClrClr() }
func (x ClrGrey300) ClrClr() clr.Clr         { return clr.Grey300 }
func (x ClrGrey400) Act()                    { x.ClrClr() }
func (x ClrGrey400) Ifc() interface{}        { return x.ClrClr() }
func (x ClrGrey400) ClrClr() clr.Clr         { return clr.Grey400 }
func (x ClrGrey500) Act()                    { x.ClrClr() }
func (x ClrGrey500) Ifc() interface{}        { return x.ClrClr() }
func (x ClrGrey500) ClrClr() clr.Clr         { return clr.Grey500 }
func (x ClrGrey600) Act()                    { x.ClrClr() }
func (x ClrGrey600) Ifc() interface{}        { return x.ClrClr() }
func (x ClrGrey600) ClrClr() clr.Clr         { return clr.Grey600 }
func (x ClrGrey700) Act()                    { x.ClrClr() }
func (x ClrGrey700) Ifc() interface{}        { return x.ClrClr() }
func (x ClrGrey700) ClrClr() clr.Clr         { return clr.Grey700 }
func (x ClrGrey800) Act()                    { x.ClrClr() }
func (x ClrGrey800) Ifc() interface{}        { return x.ClrClr() }
func (x ClrGrey800) ClrClr() clr.Clr         { return clr.Grey800 }
func (x ClrGrey900) Act()                    { x.ClrClr() }
func (x ClrGrey900) Ifc() interface{}        { return x.ClrClr() }
func (x ClrGrey900) ClrClr() clr.Clr         { return clr.Grey900 }
func (x ClrBlueGrey50) Act()                 { x.ClrClr() }
func (x ClrBlueGrey50) Ifc() interface{}     { return x.ClrClr() }
func (x ClrBlueGrey50) ClrClr() clr.Clr      { return clr.BlueGrey50 }
func (x ClrBlueGrey100) Act()                { x.ClrClr() }
func (x ClrBlueGrey100) Ifc() interface{}    { return x.ClrClr() }
func (x ClrBlueGrey100) ClrClr() clr.Clr     { return clr.BlueGrey100 }
func (x ClrBlueGrey200) Act()                { x.ClrClr() }
func (x ClrBlueGrey200) Ifc() interface{}    { return x.ClrClr() }
func (x ClrBlueGrey200) ClrClr() clr.Clr     { return clr.BlueGrey200 }
func (x ClrBlueGrey300) Act()                { x.ClrClr() }
func (x ClrBlueGrey300) Ifc() interface{}    { return x.ClrClr() }
func (x ClrBlueGrey300) ClrClr() clr.Clr     { return clr.BlueGrey300 }
func (x ClrBlueGrey400) Act()                { x.ClrClr() }
func (x ClrBlueGrey400) Ifc() interface{}    { return x.ClrClr() }
func (x ClrBlueGrey400) ClrClr() clr.Clr     { return clr.BlueGrey400 }
func (x ClrBlueGrey500) Act()                { x.ClrClr() }
func (x ClrBlueGrey500) Ifc() interface{}    { return x.ClrClr() }
func (x ClrBlueGrey500) ClrClr() clr.Clr     { return clr.BlueGrey500 }
func (x ClrBlueGrey600) Act()                { x.ClrClr() }
func (x ClrBlueGrey600) Ifc() interface{}    { return x.ClrClr() }
func (x ClrBlueGrey600) ClrClr() clr.Clr     { return clr.BlueGrey600 }
func (x ClrBlueGrey700) Act()                { x.ClrClr() }
func (x ClrBlueGrey700) Ifc() interface{}    { return x.ClrClr() }
func (x ClrBlueGrey700) ClrClr() clr.Clr     { return clr.BlueGrey700 }
func (x ClrBlueGrey800) Act()                { x.ClrClr() }
func (x ClrBlueGrey800) Ifc() interface{}    { return x.ClrClr() }
func (x ClrBlueGrey800) ClrClr() clr.Clr     { return clr.BlueGrey800 }
func (x ClrBlueGrey900) Act()                { x.ClrClr() }
func (x ClrBlueGrey900) Ifc() interface{}    { return x.ClrClr() }
func (x ClrBlueGrey900) ClrClr() clr.Clr     { return clr.BlueGrey900 }
func (x PenBlack) Act()                      { x.PenPen() }
func (x PenBlack) Ifc() interface{}          { return x.PenPen() }
func (x PenBlack) PenPen() pen.Pen           { return pen.Black }
func (x PenWhite) Act()                      { x.PenPen() }
func (x PenWhite) Ifc() interface{}          { return x.PenPen() }
func (x PenWhite) PenPen() pen.Pen           { return pen.White }
func (x PenRed50) Act()                      { x.PenPen() }
func (x PenRed50) Ifc() interface{}          { return x.PenPen() }
func (x PenRed50) PenPen() pen.Pen           { return pen.Red50 }
func (x PenRed100) Act()                     { x.PenPen() }
func (x PenRed100) Ifc() interface{}         { return x.PenPen() }
func (x PenRed100) PenPen() pen.Pen          { return pen.Red100 }
func (x PenRed200) Act()                     { x.PenPen() }
func (x PenRed200) Ifc() interface{}         { return x.PenPen() }
func (x PenRed200) PenPen() pen.Pen          { return pen.Red200 }
func (x PenRed300) Act()                     { x.PenPen() }
func (x PenRed300) Ifc() interface{}         { return x.PenPen() }
func (x PenRed300) PenPen() pen.Pen          { return pen.Red300 }
func (x PenRed400) Act()                     { x.PenPen() }
func (x PenRed400) Ifc() interface{}         { return x.PenPen() }
func (x PenRed400) PenPen() pen.Pen          { return pen.Red400 }
func (x PenRed500) Act()                     { x.PenPen() }
func (x PenRed500) Ifc() interface{}         { return x.PenPen() }
func (x PenRed500) PenPen() pen.Pen          { return pen.Red500 }
func (x PenRed600) Act()                     { x.PenPen() }
func (x PenRed600) Ifc() interface{}         { return x.PenPen() }
func (x PenRed600) PenPen() pen.Pen          { return pen.Red600 }
func (x PenRed700) Act()                     { x.PenPen() }
func (x PenRed700) Ifc() interface{}         { return x.PenPen() }
func (x PenRed700) PenPen() pen.Pen          { return pen.Red700 }
func (x PenRed800) Act()                     { x.PenPen() }
func (x PenRed800) Ifc() interface{}         { return x.PenPen() }
func (x PenRed800) PenPen() pen.Pen          { return pen.Red800 }
func (x PenRed900) Act()                     { x.PenPen() }
func (x PenRed900) Ifc() interface{}         { return x.PenPen() }
func (x PenRed900) PenPen() pen.Pen          { return pen.Red900 }
func (x PenRedA100) Act()                    { x.PenPen() }
func (x PenRedA100) Ifc() interface{}        { return x.PenPen() }
func (x PenRedA100) PenPen() pen.Pen         { return pen.RedA100 }
func (x PenRedA200) Act()                    { x.PenPen() }
func (x PenRedA200) Ifc() interface{}        { return x.PenPen() }
func (x PenRedA200) PenPen() pen.Pen         { return pen.RedA200 }
func (x PenRedA400) Act()                    { x.PenPen() }
func (x PenRedA400) Ifc() interface{}        { return x.PenPen() }
func (x PenRedA400) PenPen() pen.Pen         { return pen.RedA400 }
func (x PenRedA700) Act()                    { x.PenPen() }
func (x PenRedA700) Ifc() interface{}        { return x.PenPen() }
func (x PenRedA700) PenPen() pen.Pen         { return pen.RedA700 }
func (x PenPink50) Act()                     { x.PenPen() }
func (x PenPink50) Ifc() interface{}         { return x.PenPen() }
func (x PenPink50) PenPen() pen.Pen          { return pen.Pink50 }
func (x PenPink100) Act()                    { x.PenPen() }
func (x PenPink100) Ifc() interface{}        { return x.PenPen() }
func (x PenPink100) PenPen() pen.Pen         { return pen.Pink100 }
func (x PenPink200) Act()                    { x.PenPen() }
func (x PenPink200) Ifc() interface{}        { return x.PenPen() }
func (x PenPink200) PenPen() pen.Pen         { return pen.Pink200 }
func (x PenPink300) Act()                    { x.PenPen() }
func (x PenPink300) Ifc() interface{}        { return x.PenPen() }
func (x PenPink300) PenPen() pen.Pen         { return pen.Pink300 }
func (x PenPink400) Act()                    { x.PenPen() }
func (x PenPink400) Ifc() interface{}        { return x.PenPen() }
func (x PenPink400) PenPen() pen.Pen         { return pen.Pink400 }
func (x PenPink500) Act()                    { x.PenPen() }
func (x PenPink500) Ifc() interface{}        { return x.PenPen() }
func (x PenPink500) PenPen() pen.Pen         { return pen.Pink500 }
func (x PenPink600) Act()                    { x.PenPen() }
func (x PenPink600) Ifc() interface{}        { return x.PenPen() }
func (x PenPink600) PenPen() pen.Pen         { return pen.Pink600 }
func (x PenPink700) Act()                    { x.PenPen() }
func (x PenPink700) Ifc() interface{}        { return x.PenPen() }
func (x PenPink700) PenPen() pen.Pen         { return pen.Pink700 }
func (x PenPink800) Act()                    { x.PenPen() }
func (x PenPink800) Ifc() interface{}        { return x.PenPen() }
func (x PenPink800) PenPen() pen.Pen         { return pen.Pink800 }
func (x PenPink900) Act()                    { x.PenPen() }
func (x PenPink900) Ifc() interface{}        { return x.PenPen() }
func (x PenPink900) PenPen() pen.Pen         { return pen.Pink900 }
func (x PenPinkA100) Act()                   { x.PenPen() }
func (x PenPinkA100) Ifc() interface{}       { return x.PenPen() }
func (x PenPinkA100) PenPen() pen.Pen        { return pen.PinkA100 }
func (x PenPinkA200) Act()                   { x.PenPen() }
func (x PenPinkA200) Ifc() interface{}       { return x.PenPen() }
func (x PenPinkA200) PenPen() pen.Pen        { return pen.PinkA200 }
func (x PenPinkA400) Act()                   { x.PenPen() }
func (x PenPinkA400) Ifc() interface{}       { return x.PenPen() }
func (x PenPinkA400) PenPen() pen.Pen        { return pen.PinkA400 }
func (x PenPinkA700) Act()                   { x.PenPen() }
func (x PenPinkA700) Ifc() interface{}       { return x.PenPen() }
func (x PenPinkA700) PenPen() pen.Pen        { return pen.PinkA700 }
func (x PenPurple50) Act()                   { x.PenPen() }
func (x PenPurple50) Ifc() interface{}       { return x.PenPen() }
func (x PenPurple50) PenPen() pen.Pen        { return pen.Purple50 }
func (x PenPurple100) Act()                  { x.PenPen() }
func (x PenPurple100) Ifc() interface{}      { return x.PenPen() }
func (x PenPurple100) PenPen() pen.Pen       { return pen.Purple100 }
func (x PenPurple200) Act()                  { x.PenPen() }
func (x PenPurple200) Ifc() interface{}      { return x.PenPen() }
func (x PenPurple200) PenPen() pen.Pen       { return pen.Purple200 }
func (x PenPurple300) Act()                  { x.PenPen() }
func (x PenPurple300) Ifc() interface{}      { return x.PenPen() }
func (x PenPurple300) PenPen() pen.Pen       { return pen.Purple300 }
func (x PenPurple400) Act()                  { x.PenPen() }
func (x PenPurple400) Ifc() interface{}      { return x.PenPen() }
func (x PenPurple400) PenPen() pen.Pen       { return pen.Purple400 }
func (x PenPurple500) Act()                  { x.PenPen() }
func (x PenPurple500) Ifc() interface{}      { return x.PenPen() }
func (x PenPurple500) PenPen() pen.Pen       { return pen.Purple500 }
func (x PenPurple600) Act()                  { x.PenPen() }
func (x PenPurple600) Ifc() interface{}      { return x.PenPen() }
func (x PenPurple600) PenPen() pen.Pen       { return pen.Purple600 }
func (x PenPurple700) Act()                  { x.PenPen() }
func (x PenPurple700) Ifc() interface{}      { return x.PenPen() }
func (x PenPurple700) PenPen() pen.Pen       { return pen.Purple700 }
func (x PenPurple800) Act()                  { x.PenPen() }
func (x PenPurple800) Ifc() interface{}      { return x.PenPen() }
func (x PenPurple800) PenPen() pen.Pen       { return pen.Purple800 }
func (x PenPurple900) Act()                  { x.PenPen() }
func (x PenPurple900) Ifc() interface{}      { return x.PenPen() }
func (x PenPurple900) PenPen() pen.Pen       { return pen.Purple900 }
func (x PenPurpleA100) Act()                 { x.PenPen() }
func (x PenPurpleA100) Ifc() interface{}     { return x.PenPen() }
func (x PenPurpleA100) PenPen() pen.Pen      { return pen.PurpleA100 }
func (x PenPurpleA200) Act()                 { x.PenPen() }
func (x PenPurpleA200) Ifc() interface{}     { return x.PenPen() }
func (x PenPurpleA200) PenPen() pen.Pen      { return pen.PurpleA200 }
func (x PenPurpleA400) Act()                 { x.PenPen() }
func (x PenPurpleA400) Ifc() interface{}     { return x.PenPen() }
func (x PenPurpleA400) PenPen() pen.Pen      { return pen.PurpleA400 }
func (x PenPurpleA700) Act()                 { x.PenPen() }
func (x PenPurpleA700) Ifc() interface{}     { return x.PenPen() }
func (x PenPurpleA700) PenPen() pen.Pen      { return pen.PurpleA700 }
func (x PenDeepPurple50) Act()               { x.PenPen() }
func (x PenDeepPurple50) Ifc() interface{}   { return x.PenPen() }
func (x PenDeepPurple50) PenPen() pen.Pen    { return pen.DeepPurple50 }
func (x PenDeepPurple100) Act()              { x.PenPen() }
func (x PenDeepPurple100) Ifc() interface{}  { return x.PenPen() }
func (x PenDeepPurple100) PenPen() pen.Pen   { return pen.DeepPurple100 }
func (x PenDeepPurple200) Act()              { x.PenPen() }
func (x PenDeepPurple200) Ifc() interface{}  { return x.PenPen() }
func (x PenDeepPurple200) PenPen() pen.Pen   { return pen.DeepPurple200 }
func (x PenDeepPurple300) Act()              { x.PenPen() }
func (x PenDeepPurple300) Ifc() interface{}  { return x.PenPen() }
func (x PenDeepPurple300) PenPen() pen.Pen   { return pen.DeepPurple300 }
func (x PenDeepPurple400) Act()              { x.PenPen() }
func (x PenDeepPurple400) Ifc() interface{}  { return x.PenPen() }
func (x PenDeepPurple400) PenPen() pen.Pen   { return pen.DeepPurple400 }
func (x PenDeepPurple500) Act()              { x.PenPen() }
func (x PenDeepPurple500) Ifc() interface{}  { return x.PenPen() }
func (x PenDeepPurple500) PenPen() pen.Pen   { return pen.DeepPurple500 }
func (x PenDeepPurple600) Act()              { x.PenPen() }
func (x PenDeepPurple600) Ifc() interface{}  { return x.PenPen() }
func (x PenDeepPurple600) PenPen() pen.Pen   { return pen.DeepPurple600 }
func (x PenDeepPurple700) Act()              { x.PenPen() }
func (x PenDeepPurple700) Ifc() interface{}  { return x.PenPen() }
func (x PenDeepPurple700) PenPen() pen.Pen   { return pen.DeepPurple700 }
func (x PenDeepPurple800) Act()              { x.PenPen() }
func (x PenDeepPurple800) Ifc() interface{}  { return x.PenPen() }
func (x PenDeepPurple800) PenPen() pen.Pen   { return pen.DeepPurple800 }
func (x PenDeepPurple900) Act()              { x.PenPen() }
func (x PenDeepPurple900) Ifc() interface{}  { return x.PenPen() }
func (x PenDeepPurple900) PenPen() pen.Pen   { return pen.DeepPurple900 }
func (x PenDeepPurpleA100) Act()             { x.PenPen() }
func (x PenDeepPurpleA100) Ifc() interface{} { return x.PenPen() }
func (x PenDeepPurpleA100) PenPen() pen.Pen  { return pen.DeepPurpleA100 }
func (x PenDeepPurpleA200) Act()             { x.PenPen() }
func (x PenDeepPurpleA200) Ifc() interface{} { return x.PenPen() }
func (x PenDeepPurpleA200) PenPen() pen.Pen  { return pen.DeepPurpleA200 }
func (x PenDeepPurpleA400) Act()             { x.PenPen() }
func (x PenDeepPurpleA400) Ifc() interface{} { return x.PenPen() }
func (x PenDeepPurpleA400) PenPen() pen.Pen  { return pen.DeepPurpleA400 }
func (x PenDeepPurpleA700) Act()             { x.PenPen() }
func (x PenDeepPurpleA700) Ifc() interface{} { return x.PenPen() }
func (x PenDeepPurpleA700) PenPen() pen.Pen  { return pen.DeepPurpleA700 }
func (x PenIndigo50) Act()                   { x.PenPen() }
func (x PenIndigo50) Ifc() interface{}       { return x.PenPen() }
func (x PenIndigo50) PenPen() pen.Pen        { return pen.Indigo50 }
func (x PenIndigo100) Act()                  { x.PenPen() }
func (x PenIndigo100) Ifc() interface{}      { return x.PenPen() }
func (x PenIndigo100) PenPen() pen.Pen       { return pen.Indigo100 }
func (x PenIndigo200) Act()                  { x.PenPen() }
func (x PenIndigo200) Ifc() interface{}      { return x.PenPen() }
func (x PenIndigo200) PenPen() pen.Pen       { return pen.Indigo200 }
func (x PenIndigo300) Act()                  { x.PenPen() }
func (x PenIndigo300) Ifc() interface{}      { return x.PenPen() }
func (x PenIndigo300) PenPen() pen.Pen       { return pen.Indigo300 }
func (x PenIndigo400) Act()                  { x.PenPen() }
func (x PenIndigo400) Ifc() interface{}      { return x.PenPen() }
func (x PenIndigo400) PenPen() pen.Pen       { return pen.Indigo400 }
func (x PenIndigo500) Act()                  { x.PenPen() }
func (x PenIndigo500) Ifc() interface{}      { return x.PenPen() }
func (x PenIndigo500) PenPen() pen.Pen       { return pen.Indigo500 }
func (x PenIndigo600) Act()                  { x.PenPen() }
func (x PenIndigo600) Ifc() interface{}      { return x.PenPen() }
func (x PenIndigo600) PenPen() pen.Pen       { return pen.Indigo600 }
func (x PenIndigo700) Act()                  { x.PenPen() }
func (x PenIndigo700) Ifc() interface{}      { return x.PenPen() }
func (x PenIndigo700) PenPen() pen.Pen       { return pen.Indigo700 }
func (x PenIndigo800) Act()                  { x.PenPen() }
func (x PenIndigo800) Ifc() interface{}      { return x.PenPen() }
func (x PenIndigo800) PenPen() pen.Pen       { return pen.Indigo800 }
func (x PenIndigo900) Act()                  { x.PenPen() }
func (x PenIndigo900) Ifc() interface{}      { return x.PenPen() }
func (x PenIndigo900) PenPen() pen.Pen       { return pen.Indigo900 }
func (x PenIndigoA100) Act()                 { x.PenPen() }
func (x PenIndigoA100) Ifc() interface{}     { return x.PenPen() }
func (x PenIndigoA100) PenPen() pen.Pen      { return pen.IndigoA100 }
func (x PenIndigoA200) Act()                 { x.PenPen() }
func (x PenIndigoA200) Ifc() interface{}     { return x.PenPen() }
func (x PenIndigoA200) PenPen() pen.Pen      { return pen.IndigoA200 }
func (x PenIndigoA400) Act()                 { x.PenPen() }
func (x PenIndigoA400) Ifc() interface{}     { return x.PenPen() }
func (x PenIndigoA400) PenPen() pen.Pen      { return pen.IndigoA400 }
func (x PenIndigoA700) Act()                 { x.PenPen() }
func (x PenIndigoA700) Ifc() interface{}     { return x.PenPen() }
func (x PenIndigoA700) PenPen() pen.Pen      { return pen.IndigoA700 }
func (x PenBlue50) Act()                     { x.PenPen() }
func (x PenBlue50) Ifc() interface{}         { return x.PenPen() }
func (x PenBlue50) PenPen() pen.Pen          { return pen.Blue50 }
func (x PenBlue100) Act()                    { x.PenPen() }
func (x PenBlue100) Ifc() interface{}        { return x.PenPen() }
func (x PenBlue100) PenPen() pen.Pen         { return pen.Blue100 }
func (x PenBlue200) Act()                    { x.PenPen() }
func (x PenBlue200) Ifc() interface{}        { return x.PenPen() }
func (x PenBlue200) PenPen() pen.Pen         { return pen.Blue200 }
func (x PenBlue300) Act()                    { x.PenPen() }
func (x PenBlue300) Ifc() interface{}        { return x.PenPen() }
func (x PenBlue300) PenPen() pen.Pen         { return pen.Blue300 }
func (x PenBlue400) Act()                    { x.PenPen() }
func (x PenBlue400) Ifc() interface{}        { return x.PenPen() }
func (x PenBlue400) PenPen() pen.Pen         { return pen.Blue400 }
func (x PenBlue500) Act()                    { x.PenPen() }
func (x PenBlue500) Ifc() interface{}        { return x.PenPen() }
func (x PenBlue500) PenPen() pen.Pen         { return pen.Blue500 }
func (x PenBlue600) Act()                    { x.PenPen() }
func (x PenBlue600) Ifc() interface{}        { return x.PenPen() }
func (x PenBlue600) PenPen() pen.Pen         { return pen.Blue600 }
func (x PenBlue700) Act()                    { x.PenPen() }
func (x PenBlue700) Ifc() interface{}        { return x.PenPen() }
func (x PenBlue700) PenPen() pen.Pen         { return pen.Blue700 }
func (x PenBlue800) Act()                    { x.PenPen() }
func (x PenBlue800) Ifc() interface{}        { return x.PenPen() }
func (x PenBlue800) PenPen() pen.Pen         { return pen.Blue800 }
func (x PenBlue900) Act()                    { x.PenPen() }
func (x PenBlue900) Ifc() interface{}        { return x.PenPen() }
func (x PenBlue900) PenPen() pen.Pen         { return pen.Blue900 }
func (x PenBlueA100) Act()                   { x.PenPen() }
func (x PenBlueA100) Ifc() interface{}       { return x.PenPen() }
func (x PenBlueA100) PenPen() pen.Pen        { return pen.BlueA100 }
func (x PenBlueA200) Act()                   { x.PenPen() }
func (x PenBlueA200) Ifc() interface{}       { return x.PenPen() }
func (x PenBlueA200) PenPen() pen.Pen        { return pen.BlueA200 }
func (x PenBlueA400) Act()                   { x.PenPen() }
func (x PenBlueA400) Ifc() interface{}       { return x.PenPen() }
func (x PenBlueA400) PenPen() pen.Pen        { return pen.BlueA400 }
func (x PenBlueA700) Act()                   { x.PenPen() }
func (x PenBlueA700) Ifc() interface{}       { return x.PenPen() }
func (x PenBlueA700) PenPen() pen.Pen        { return pen.BlueA700 }
func (x PenLightBlue50) Act()                { x.PenPen() }
func (x PenLightBlue50) Ifc() interface{}    { return x.PenPen() }
func (x PenLightBlue50) PenPen() pen.Pen     { return pen.LightBlue50 }
func (x PenLightBlue100) Act()               { x.PenPen() }
func (x PenLightBlue100) Ifc() interface{}   { return x.PenPen() }
func (x PenLightBlue100) PenPen() pen.Pen    { return pen.LightBlue100 }
func (x PenLightBlue200) Act()               { x.PenPen() }
func (x PenLightBlue200) Ifc() interface{}   { return x.PenPen() }
func (x PenLightBlue200) PenPen() pen.Pen    { return pen.LightBlue200 }
func (x PenLightBlue300) Act()               { x.PenPen() }
func (x PenLightBlue300) Ifc() interface{}   { return x.PenPen() }
func (x PenLightBlue300) PenPen() pen.Pen    { return pen.LightBlue300 }
func (x PenLightBlue400) Act()               { x.PenPen() }
func (x PenLightBlue400) Ifc() interface{}   { return x.PenPen() }
func (x PenLightBlue400) PenPen() pen.Pen    { return pen.LightBlue400 }
func (x PenLightBlue500) Act()               { x.PenPen() }
func (x PenLightBlue500) Ifc() interface{}   { return x.PenPen() }
func (x PenLightBlue500) PenPen() pen.Pen    { return pen.LightBlue500 }
func (x PenLightBlue600) Act()               { x.PenPen() }
func (x PenLightBlue600) Ifc() interface{}   { return x.PenPen() }
func (x PenLightBlue600) PenPen() pen.Pen    { return pen.LightBlue600 }
func (x PenLightBlue700) Act()               { x.PenPen() }
func (x PenLightBlue700) Ifc() interface{}   { return x.PenPen() }
func (x PenLightBlue700) PenPen() pen.Pen    { return pen.LightBlue700 }
func (x PenLightBlue800) Act()               { x.PenPen() }
func (x PenLightBlue800) Ifc() interface{}   { return x.PenPen() }
func (x PenLightBlue800) PenPen() pen.Pen    { return pen.LightBlue800 }
func (x PenLightBlue900) Act()               { x.PenPen() }
func (x PenLightBlue900) Ifc() interface{}   { return x.PenPen() }
func (x PenLightBlue900) PenPen() pen.Pen    { return pen.LightBlue900 }
func (x PenLightBlueA100) Act()              { x.PenPen() }
func (x PenLightBlueA100) Ifc() interface{}  { return x.PenPen() }
func (x PenLightBlueA100) PenPen() pen.Pen   { return pen.LightBlueA100 }
func (x PenLightBlueA200) Act()              { x.PenPen() }
func (x PenLightBlueA200) Ifc() interface{}  { return x.PenPen() }
func (x PenLightBlueA200) PenPen() pen.Pen   { return pen.LightBlueA200 }
func (x PenLightBlueA400) Act()              { x.PenPen() }
func (x PenLightBlueA400) Ifc() interface{}  { return x.PenPen() }
func (x PenLightBlueA400) PenPen() pen.Pen   { return pen.LightBlueA400 }
func (x PenLightBlueA700) Act()              { x.PenPen() }
func (x PenLightBlueA700) Ifc() interface{}  { return x.PenPen() }
func (x PenLightBlueA700) PenPen() pen.Pen   { return pen.LightBlueA700 }
func (x PenCyan50) Act()                     { x.PenPen() }
func (x PenCyan50) Ifc() interface{}         { return x.PenPen() }
func (x PenCyan50) PenPen() pen.Pen          { return pen.Cyan50 }
func (x PenCyan100) Act()                    { x.PenPen() }
func (x PenCyan100) Ifc() interface{}        { return x.PenPen() }
func (x PenCyan100) PenPen() pen.Pen         { return pen.Cyan100 }
func (x PenCyan200) Act()                    { x.PenPen() }
func (x PenCyan200) Ifc() interface{}        { return x.PenPen() }
func (x PenCyan200) PenPen() pen.Pen         { return pen.Cyan200 }
func (x PenCyan300) Act()                    { x.PenPen() }
func (x PenCyan300) Ifc() interface{}        { return x.PenPen() }
func (x PenCyan300) PenPen() pen.Pen         { return pen.Cyan300 }
func (x PenCyan400) Act()                    { x.PenPen() }
func (x PenCyan400) Ifc() interface{}        { return x.PenPen() }
func (x PenCyan400) PenPen() pen.Pen         { return pen.Cyan400 }
func (x PenCyan500) Act()                    { x.PenPen() }
func (x PenCyan500) Ifc() interface{}        { return x.PenPen() }
func (x PenCyan500) PenPen() pen.Pen         { return pen.Cyan500 }
func (x PenCyan600) Act()                    { x.PenPen() }
func (x PenCyan600) Ifc() interface{}        { return x.PenPen() }
func (x PenCyan600) PenPen() pen.Pen         { return pen.Cyan600 }
func (x PenCyan700) Act()                    { x.PenPen() }
func (x PenCyan700) Ifc() interface{}        { return x.PenPen() }
func (x PenCyan700) PenPen() pen.Pen         { return pen.Cyan700 }
func (x PenCyan800) Act()                    { x.PenPen() }
func (x PenCyan800) Ifc() interface{}        { return x.PenPen() }
func (x PenCyan800) PenPen() pen.Pen         { return pen.Cyan800 }
func (x PenCyan900) Act()                    { x.PenPen() }
func (x PenCyan900) Ifc() interface{}        { return x.PenPen() }
func (x PenCyan900) PenPen() pen.Pen         { return pen.Cyan900 }
func (x PenCyanA100) Act()                   { x.PenPen() }
func (x PenCyanA100) Ifc() interface{}       { return x.PenPen() }
func (x PenCyanA100) PenPen() pen.Pen        { return pen.CyanA100 }
func (x PenCyanA200) Act()                   { x.PenPen() }
func (x PenCyanA200) Ifc() interface{}       { return x.PenPen() }
func (x PenCyanA200) PenPen() pen.Pen        { return pen.CyanA200 }
func (x PenCyanA400) Act()                   { x.PenPen() }
func (x PenCyanA400) Ifc() interface{}       { return x.PenPen() }
func (x PenCyanA400) PenPen() pen.Pen        { return pen.CyanA400 }
func (x PenCyanA700) Act()                   { x.PenPen() }
func (x PenCyanA700) Ifc() interface{}       { return x.PenPen() }
func (x PenCyanA700) PenPen() pen.Pen        { return pen.CyanA700 }
func (x PenTeal50) Act()                     { x.PenPen() }
func (x PenTeal50) Ifc() interface{}         { return x.PenPen() }
func (x PenTeal50) PenPen() pen.Pen          { return pen.Teal50 }
func (x PenTeal100) Act()                    { x.PenPen() }
func (x PenTeal100) Ifc() interface{}        { return x.PenPen() }
func (x PenTeal100) PenPen() pen.Pen         { return pen.Teal100 }
func (x PenTeal200) Act()                    { x.PenPen() }
func (x PenTeal200) Ifc() interface{}        { return x.PenPen() }
func (x PenTeal200) PenPen() pen.Pen         { return pen.Teal200 }
func (x PenTeal300) Act()                    { x.PenPen() }
func (x PenTeal300) Ifc() interface{}        { return x.PenPen() }
func (x PenTeal300) PenPen() pen.Pen         { return pen.Teal300 }
func (x PenTeal400) Act()                    { x.PenPen() }
func (x PenTeal400) Ifc() interface{}        { return x.PenPen() }
func (x PenTeal400) PenPen() pen.Pen         { return pen.Teal400 }
func (x PenTeal500) Act()                    { x.PenPen() }
func (x PenTeal500) Ifc() interface{}        { return x.PenPen() }
func (x PenTeal500) PenPen() pen.Pen         { return pen.Teal500 }
func (x PenTeal600) Act()                    { x.PenPen() }
func (x PenTeal600) Ifc() interface{}        { return x.PenPen() }
func (x PenTeal600) PenPen() pen.Pen         { return pen.Teal600 }
func (x PenTeal700) Act()                    { x.PenPen() }
func (x PenTeal700) Ifc() interface{}        { return x.PenPen() }
func (x PenTeal700) PenPen() pen.Pen         { return pen.Teal700 }
func (x PenTeal800) Act()                    { x.PenPen() }
func (x PenTeal800) Ifc() interface{}        { return x.PenPen() }
func (x PenTeal800) PenPen() pen.Pen         { return pen.Teal800 }
func (x PenTeal900) Act()                    { x.PenPen() }
func (x PenTeal900) Ifc() interface{}        { return x.PenPen() }
func (x PenTeal900) PenPen() pen.Pen         { return pen.Teal900 }
func (x PenTealA100) Act()                   { x.PenPen() }
func (x PenTealA100) Ifc() interface{}       { return x.PenPen() }
func (x PenTealA100) PenPen() pen.Pen        { return pen.TealA100 }
func (x PenTealA200) Act()                   { x.PenPen() }
func (x PenTealA200) Ifc() interface{}       { return x.PenPen() }
func (x PenTealA200) PenPen() pen.Pen        { return pen.TealA200 }
func (x PenTealA400) Act()                   { x.PenPen() }
func (x PenTealA400) Ifc() interface{}       { return x.PenPen() }
func (x PenTealA400) PenPen() pen.Pen        { return pen.TealA400 }
func (x PenTealA700) Act()                   { x.PenPen() }
func (x PenTealA700) Ifc() interface{}       { return x.PenPen() }
func (x PenTealA700) PenPen() pen.Pen        { return pen.TealA700 }
func (x PenGreen50) Act()                    { x.PenPen() }
func (x PenGreen50) Ifc() interface{}        { return x.PenPen() }
func (x PenGreen50) PenPen() pen.Pen         { return pen.Green50 }
func (x PenGreen100) Act()                   { x.PenPen() }
func (x PenGreen100) Ifc() interface{}       { return x.PenPen() }
func (x PenGreen100) PenPen() pen.Pen        { return pen.Green100 }
func (x PenGreen200) Act()                   { x.PenPen() }
func (x PenGreen200) Ifc() interface{}       { return x.PenPen() }
func (x PenGreen200) PenPen() pen.Pen        { return pen.Green200 }
func (x PenGreen300) Act()                   { x.PenPen() }
func (x PenGreen300) Ifc() interface{}       { return x.PenPen() }
func (x PenGreen300) PenPen() pen.Pen        { return pen.Green300 }
func (x PenGreen400) Act()                   { x.PenPen() }
func (x PenGreen400) Ifc() interface{}       { return x.PenPen() }
func (x PenGreen400) PenPen() pen.Pen        { return pen.Green400 }
func (x PenGreen500) Act()                   { x.PenPen() }
func (x PenGreen500) Ifc() interface{}       { return x.PenPen() }
func (x PenGreen500) PenPen() pen.Pen        { return pen.Green500 }
func (x PenGreen600) Act()                   { x.PenPen() }
func (x PenGreen600) Ifc() interface{}       { return x.PenPen() }
func (x PenGreen600) PenPen() pen.Pen        { return pen.Green600 }
func (x PenGreen700) Act()                   { x.PenPen() }
func (x PenGreen700) Ifc() interface{}       { return x.PenPen() }
func (x PenGreen700) PenPen() pen.Pen        { return pen.Green700 }
func (x PenGreen800) Act()                   { x.PenPen() }
func (x PenGreen800) Ifc() interface{}       { return x.PenPen() }
func (x PenGreen800) PenPen() pen.Pen        { return pen.Green800 }
func (x PenGreen900) Act()                   { x.PenPen() }
func (x PenGreen900) Ifc() interface{}       { return x.PenPen() }
func (x PenGreen900) PenPen() pen.Pen        { return pen.Green900 }
func (x PenGreenA100) Act()                  { x.PenPen() }
func (x PenGreenA100) Ifc() interface{}      { return x.PenPen() }
func (x PenGreenA100) PenPen() pen.Pen       { return pen.GreenA100 }
func (x PenGreenA200) Act()                  { x.PenPen() }
func (x PenGreenA200) Ifc() interface{}      { return x.PenPen() }
func (x PenGreenA200) PenPen() pen.Pen       { return pen.GreenA200 }
func (x PenGreenA400) Act()                  { x.PenPen() }
func (x PenGreenA400) Ifc() interface{}      { return x.PenPen() }
func (x PenGreenA400) PenPen() pen.Pen       { return pen.GreenA400 }
func (x PenGreenA700) Act()                  { x.PenPen() }
func (x PenGreenA700) Ifc() interface{}      { return x.PenPen() }
func (x PenGreenA700) PenPen() pen.Pen       { return pen.GreenA700 }
func (x PenLightGreen50) Act()               { x.PenPen() }
func (x PenLightGreen50) Ifc() interface{}   { return x.PenPen() }
func (x PenLightGreen50) PenPen() pen.Pen    { return pen.LightGreen50 }
func (x PenLightGreen100) Act()              { x.PenPen() }
func (x PenLightGreen100) Ifc() interface{}  { return x.PenPen() }
func (x PenLightGreen100) PenPen() pen.Pen   { return pen.LightGreen100 }
func (x PenLightGreen200) Act()              { x.PenPen() }
func (x PenLightGreen200) Ifc() interface{}  { return x.PenPen() }
func (x PenLightGreen200) PenPen() pen.Pen   { return pen.LightGreen200 }
func (x PenLightGreen300) Act()              { x.PenPen() }
func (x PenLightGreen300) Ifc() interface{}  { return x.PenPen() }
func (x PenLightGreen300) PenPen() pen.Pen   { return pen.LightGreen300 }
func (x PenLightGreen400) Act()              { x.PenPen() }
func (x PenLightGreen400) Ifc() interface{}  { return x.PenPen() }
func (x PenLightGreen400) PenPen() pen.Pen   { return pen.LightGreen400 }
func (x PenLightGreen500) Act()              { x.PenPen() }
func (x PenLightGreen500) Ifc() interface{}  { return x.PenPen() }
func (x PenLightGreen500) PenPen() pen.Pen   { return pen.LightGreen500 }
func (x PenLightGreen600) Act()              { x.PenPen() }
func (x PenLightGreen600) Ifc() interface{}  { return x.PenPen() }
func (x PenLightGreen600) PenPen() pen.Pen   { return pen.LightGreen600 }
func (x PenLightGreen700) Act()              { x.PenPen() }
func (x PenLightGreen700) Ifc() interface{}  { return x.PenPen() }
func (x PenLightGreen700) PenPen() pen.Pen   { return pen.LightGreen700 }
func (x PenLightGreen800) Act()              { x.PenPen() }
func (x PenLightGreen800) Ifc() interface{}  { return x.PenPen() }
func (x PenLightGreen800) PenPen() pen.Pen   { return pen.LightGreen800 }
func (x PenLightGreen900) Act()              { x.PenPen() }
func (x PenLightGreen900) Ifc() interface{}  { return x.PenPen() }
func (x PenLightGreen900) PenPen() pen.Pen   { return pen.LightGreen900 }
func (x PenLightGreenA100) Act()             { x.PenPen() }
func (x PenLightGreenA100) Ifc() interface{} { return x.PenPen() }
func (x PenLightGreenA100) PenPen() pen.Pen  { return pen.LightGreenA100 }
func (x PenLightGreenA200) Act()             { x.PenPen() }
func (x PenLightGreenA200) Ifc() interface{} { return x.PenPen() }
func (x PenLightGreenA200) PenPen() pen.Pen  { return pen.LightGreenA200 }
func (x PenLightGreenA400) Act()             { x.PenPen() }
func (x PenLightGreenA400) Ifc() interface{} { return x.PenPen() }
func (x PenLightGreenA400) PenPen() pen.Pen  { return pen.LightGreenA400 }
func (x PenLightGreenA700) Act()             { x.PenPen() }
func (x PenLightGreenA700) Ifc() interface{} { return x.PenPen() }
func (x PenLightGreenA700) PenPen() pen.Pen  { return pen.LightGreenA700 }
func (x PenLime50) Act()                     { x.PenPen() }
func (x PenLime50) Ifc() interface{}         { return x.PenPen() }
func (x PenLime50) PenPen() pen.Pen          { return pen.Lime50 }
func (x PenLime100) Act()                    { x.PenPen() }
func (x PenLime100) Ifc() interface{}        { return x.PenPen() }
func (x PenLime100) PenPen() pen.Pen         { return pen.Lime100 }
func (x PenLime200) Act()                    { x.PenPen() }
func (x PenLime200) Ifc() interface{}        { return x.PenPen() }
func (x PenLime200) PenPen() pen.Pen         { return pen.Lime200 }
func (x PenLime300) Act()                    { x.PenPen() }
func (x PenLime300) Ifc() interface{}        { return x.PenPen() }
func (x PenLime300) PenPen() pen.Pen         { return pen.Lime300 }
func (x PenLime400) Act()                    { x.PenPen() }
func (x PenLime400) Ifc() interface{}        { return x.PenPen() }
func (x PenLime400) PenPen() pen.Pen         { return pen.Lime400 }
func (x PenLime500) Act()                    { x.PenPen() }
func (x PenLime500) Ifc() interface{}        { return x.PenPen() }
func (x PenLime500) PenPen() pen.Pen         { return pen.Lime500 }
func (x PenLime600) Act()                    { x.PenPen() }
func (x PenLime600) Ifc() interface{}        { return x.PenPen() }
func (x PenLime600) PenPen() pen.Pen         { return pen.Lime600 }
func (x PenLime700) Act()                    { x.PenPen() }
func (x PenLime700) Ifc() interface{}        { return x.PenPen() }
func (x PenLime700) PenPen() pen.Pen         { return pen.Lime700 }
func (x PenLime800) Act()                    { x.PenPen() }
func (x PenLime800) Ifc() interface{}        { return x.PenPen() }
func (x PenLime800) PenPen() pen.Pen         { return pen.Lime800 }
func (x PenLime900) Act()                    { x.PenPen() }
func (x PenLime900) Ifc() interface{}        { return x.PenPen() }
func (x PenLime900) PenPen() pen.Pen         { return pen.Lime900 }
func (x PenLimeA100) Act()                   { x.PenPen() }
func (x PenLimeA100) Ifc() interface{}       { return x.PenPen() }
func (x PenLimeA100) PenPen() pen.Pen        { return pen.LimeA100 }
func (x PenLimeA200) Act()                   { x.PenPen() }
func (x PenLimeA200) Ifc() interface{}       { return x.PenPen() }
func (x PenLimeA200) PenPen() pen.Pen        { return pen.LimeA200 }
func (x PenLimeA400) Act()                   { x.PenPen() }
func (x PenLimeA400) Ifc() interface{}       { return x.PenPen() }
func (x PenLimeA400) PenPen() pen.Pen        { return pen.LimeA400 }
func (x PenLimeA700) Act()                   { x.PenPen() }
func (x PenLimeA700) Ifc() interface{}       { return x.PenPen() }
func (x PenLimeA700) PenPen() pen.Pen        { return pen.LimeA700 }
func (x PenYellow50) Act()                   { x.PenPen() }
func (x PenYellow50) Ifc() interface{}       { return x.PenPen() }
func (x PenYellow50) PenPen() pen.Pen        { return pen.Yellow50 }
func (x PenYellow100) Act()                  { x.PenPen() }
func (x PenYellow100) Ifc() interface{}      { return x.PenPen() }
func (x PenYellow100) PenPen() pen.Pen       { return pen.Yellow100 }
func (x PenYellow200) Act()                  { x.PenPen() }
func (x PenYellow200) Ifc() interface{}      { return x.PenPen() }
func (x PenYellow200) PenPen() pen.Pen       { return pen.Yellow200 }
func (x PenYellow300) Act()                  { x.PenPen() }
func (x PenYellow300) Ifc() interface{}      { return x.PenPen() }
func (x PenYellow300) PenPen() pen.Pen       { return pen.Yellow300 }
func (x PenYellow400) Act()                  { x.PenPen() }
func (x PenYellow400) Ifc() interface{}      { return x.PenPen() }
func (x PenYellow400) PenPen() pen.Pen       { return pen.Yellow400 }
func (x PenYellow500) Act()                  { x.PenPen() }
func (x PenYellow500) Ifc() interface{}      { return x.PenPen() }
func (x PenYellow500) PenPen() pen.Pen       { return pen.Yellow500 }
func (x PenYellow600) Act()                  { x.PenPen() }
func (x PenYellow600) Ifc() interface{}      { return x.PenPen() }
func (x PenYellow600) PenPen() pen.Pen       { return pen.Yellow600 }
func (x PenYellow700) Act()                  { x.PenPen() }
func (x PenYellow700) Ifc() interface{}      { return x.PenPen() }
func (x PenYellow700) PenPen() pen.Pen       { return pen.Yellow700 }
func (x PenYellow800) Act()                  { x.PenPen() }
func (x PenYellow800) Ifc() interface{}      { return x.PenPen() }
func (x PenYellow800) PenPen() pen.Pen       { return pen.Yellow800 }
func (x PenYellow900) Act()                  { x.PenPen() }
func (x PenYellow900) Ifc() interface{}      { return x.PenPen() }
func (x PenYellow900) PenPen() pen.Pen       { return pen.Yellow900 }
func (x PenYellowA100) Act()                 { x.PenPen() }
func (x PenYellowA100) Ifc() interface{}     { return x.PenPen() }
func (x PenYellowA100) PenPen() pen.Pen      { return pen.YellowA100 }
func (x PenYellowA200) Act()                 { x.PenPen() }
func (x PenYellowA200) Ifc() interface{}     { return x.PenPen() }
func (x PenYellowA200) PenPen() pen.Pen      { return pen.YellowA200 }
func (x PenYellowA400) Act()                 { x.PenPen() }
func (x PenYellowA400) Ifc() interface{}     { return x.PenPen() }
func (x PenYellowA400) PenPen() pen.Pen      { return pen.YellowA400 }
func (x PenYellowA700) Act()                 { x.PenPen() }
func (x PenYellowA700) Ifc() interface{}     { return x.PenPen() }
func (x PenYellowA700) PenPen() pen.Pen      { return pen.YellowA700 }
func (x PenAmber50) Act()                    { x.PenPen() }
func (x PenAmber50) Ifc() interface{}        { return x.PenPen() }
func (x PenAmber50) PenPen() pen.Pen         { return pen.Amber50 }
func (x PenAmber100) Act()                   { x.PenPen() }
func (x PenAmber100) Ifc() interface{}       { return x.PenPen() }
func (x PenAmber100) PenPen() pen.Pen        { return pen.Amber100 }
func (x PenAmber200) Act()                   { x.PenPen() }
func (x PenAmber200) Ifc() interface{}       { return x.PenPen() }
func (x PenAmber200) PenPen() pen.Pen        { return pen.Amber200 }
func (x PenAmber300) Act()                   { x.PenPen() }
func (x PenAmber300) Ifc() interface{}       { return x.PenPen() }
func (x PenAmber300) PenPen() pen.Pen        { return pen.Amber300 }
func (x PenAmber400) Act()                   { x.PenPen() }
func (x PenAmber400) Ifc() interface{}       { return x.PenPen() }
func (x PenAmber400) PenPen() pen.Pen        { return pen.Amber400 }
func (x PenAmber500) Act()                   { x.PenPen() }
func (x PenAmber500) Ifc() interface{}       { return x.PenPen() }
func (x PenAmber500) PenPen() pen.Pen        { return pen.Amber500 }
func (x PenAmber600) Act()                   { x.PenPen() }
func (x PenAmber600) Ifc() interface{}       { return x.PenPen() }
func (x PenAmber600) PenPen() pen.Pen        { return pen.Amber600 }
func (x PenAmber700) Act()                   { x.PenPen() }
func (x PenAmber700) Ifc() interface{}       { return x.PenPen() }
func (x PenAmber700) PenPen() pen.Pen        { return pen.Amber700 }
func (x PenAmber800) Act()                   { x.PenPen() }
func (x PenAmber800) Ifc() interface{}       { return x.PenPen() }
func (x PenAmber800) PenPen() pen.Pen        { return pen.Amber800 }
func (x PenAmber900) Act()                   { x.PenPen() }
func (x PenAmber900) Ifc() interface{}       { return x.PenPen() }
func (x PenAmber900) PenPen() pen.Pen        { return pen.Amber900 }
func (x PenAmberA100) Act()                  { x.PenPen() }
func (x PenAmberA100) Ifc() interface{}      { return x.PenPen() }
func (x PenAmberA100) PenPen() pen.Pen       { return pen.AmberA100 }
func (x PenAmberA200) Act()                  { x.PenPen() }
func (x PenAmberA200) Ifc() interface{}      { return x.PenPen() }
func (x PenAmberA200) PenPen() pen.Pen       { return pen.AmberA200 }
func (x PenAmberA400) Act()                  { x.PenPen() }
func (x PenAmberA400) Ifc() interface{}      { return x.PenPen() }
func (x PenAmberA400) PenPen() pen.Pen       { return pen.AmberA400 }
func (x PenAmberA700) Act()                  { x.PenPen() }
func (x PenAmberA700) Ifc() interface{}      { return x.PenPen() }
func (x PenAmberA700) PenPen() pen.Pen       { return pen.AmberA700 }
func (x PenOrange50) Act()                   { x.PenPen() }
func (x PenOrange50) Ifc() interface{}       { return x.PenPen() }
func (x PenOrange50) PenPen() pen.Pen        { return pen.Orange50 }
func (x PenOrange100) Act()                  { x.PenPen() }
func (x PenOrange100) Ifc() interface{}      { return x.PenPen() }
func (x PenOrange100) PenPen() pen.Pen       { return pen.Orange100 }
func (x PenOrange200) Act()                  { x.PenPen() }
func (x PenOrange200) Ifc() interface{}      { return x.PenPen() }
func (x PenOrange200) PenPen() pen.Pen       { return pen.Orange200 }
func (x PenOrange300) Act()                  { x.PenPen() }
func (x PenOrange300) Ifc() interface{}      { return x.PenPen() }
func (x PenOrange300) PenPen() pen.Pen       { return pen.Orange300 }
func (x PenOrange400) Act()                  { x.PenPen() }
func (x PenOrange400) Ifc() interface{}      { return x.PenPen() }
func (x PenOrange400) PenPen() pen.Pen       { return pen.Orange400 }
func (x PenOrange500) Act()                  { x.PenPen() }
func (x PenOrange500) Ifc() interface{}      { return x.PenPen() }
func (x PenOrange500) PenPen() pen.Pen       { return pen.Orange500 }
func (x PenOrange600) Act()                  { x.PenPen() }
func (x PenOrange600) Ifc() interface{}      { return x.PenPen() }
func (x PenOrange600) PenPen() pen.Pen       { return pen.Orange600 }
func (x PenOrange700) Act()                  { x.PenPen() }
func (x PenOrange700) Ifc() interface{}      { return x.PenPen() }
func (x PenOrange700) PenPen() pen.Pen       { return pen.Orange700 }
func (x PenOrange800) Act()                  { x.PenPen() }
func (x PenOrange800) Ifc() interface{}      { return x.PenPen() }
func (x PenOrange800) PenPen() pen.Pen       { return pen.Orange800 }
func (x PenOrange900) Act()                  { x.PenPen() }
func (x PenOrange900) Ifc() interface{}      { return x.PenPen() }
func (x PenOrange900) PenPen() pen.Pen       { return pen.Orange900 }
func (x PenOrangeA100) Act()                 { x.PenPen() }
func (x PenOrangeA100) Ifc() interface{}     { return x.PenPen() }
func (x PenOrangeA100) PenPen() pen.Pen      { return pen.OrangeA100 }
func (x PenOrangeA200) Act()                 { x.PenPen() }
func (x PenOrangeA200) Ifc() interface{}     { return x.PenPen() }
func (x PenOrangeA200) PenPen() pen.Pen      { return pen.OrangeA200 }
func (x PenOrangeA400) Act()                 { x.PenPen() }
func (x PenOrangeA400) Ifc() interface{}     { return x.PenPen() }
func (x PenOrangeA400) PenPen() pen.Pen      { return pen.OrangeA400 }
func (x PenOrangeA700) Act()                 { x.PenPen() }
func (x PenOrangeA700) Ifc() interface{}     { return x.PenPen() }
func (x PenOrangeA700) PenPen() pen.Pen      { return pen.OrangeA700 }
func (x PenDeepOrange50) Act()               { x.PenPen() }
func (x PenDeepOrange50) Ifc() interface{}   { return x.PenPen() }
func (x PenDeepOrange50) PenPen() pen.Pen    { return pen.DeepOrange50 }
func (x PenDeepOrange100) Act()              { x.PenPen() }
func (x PenDeepOrange100) Ifc() interface{}  { return x.PenPen() }
func (x PenDeepOrange100) PenPen() pen.Pen   { return pen.DeepOrange100 }
func (x PenDeepOrange200) Act()              { x.PenPen() }
func (x PenDeepOrange200) Ifc() interface{}  { return x.PenPen() }
func (x PenDeepOrange200) PenPen() pen.Pen   { return pen.DeepOrange200 }
func (x PenDeepOrange300) Act()              { x.PenPen() }
func (x PenDeepOrange300) Ifc() interface{}  { return x.PenPen() }
func (x PenDeepOrange300) PenPen() pen.Pen   { return pen.DeepOrange300 }
func (x PenDeepOrange400) Act()              { x.PenPen() }
func (x PenDeepOrange400) Ifc() interface{}  { return x.PenPen() }
func (x PenDeepOrange400) PenPen() pen.Pen   { return pen.DeepOrange400 }
func (x PenDeepOrange500) Act()              { x.PenPen() }
func (x PenDeepOrange500) Ifc() interface{}  { return x.PenPen() }
func (x PenDeepOrange500) PenPen() pen.Pen   { return pen.DeepOrange500 }
func (x PenDeepOrange600) Act()              { x.PenPen() }
func (x PenDeepOrange600) Ifc() interface{}  { return x.PenPen() }
func (x PenDeepOrange600) PenPen() pen.Pen   { return pen.DeepOrange600 }
func (x PenDeepOrange700) Act()              { x.PenPen() }
func (x PenDeepOrange700) Ifc() interface{}  { return x.PenPen() }
func (x PenDeepOrange700) PenPen() pen.Pen   { return pen.DeepOrange700 }
func (x PenDeepOrange800) Act()              { x.PenPen() }
func (x PenDeepOrange800) Ifc() interface{}  { return x.PenPen() }
func (x PenDeepOrange800) PenPen() pen.Pen   { return pen.DeepOrange800 }
func (x PenDeepOrange900) Act()              { x.PenPen() }
func (x PenDeepOrange900) Ifc() interface{}  { return x.PenPen() }
func (x PenDeepOrange900) PenPen() pen.Pen   { return pen.DeepOrange900 }
func (x PenDeepOrangeA100) Act()             { x.PenPen() }
func (x PenDeepOrangeA100) Ifc() interface{} { return x.PenPen() }
func (x PenDeepOrangeA100) PenPen() pen.Pen  { return pen.DeepOrangeA100 }
func (x PenDeepOrangeA200) Act()             { x.PenPen() }
func (x PenDeepOrangeA200) Ifc() interface{} { return x.PenPen() }
func (x PenDeepOrangeA200) PenPen() pen.Pen  { return pen.DeepOrangeA200 }
func (x PenDeepOrangeA400) Act()             { x.PenPen() }
func (x PenDeepOrangeA400) Ifc() interface{} { return x.PenPen() }
func (x PenDeepOrangeA400) PenPen() pen.Pen  { return pen.DeepOrangeA400 }
func (x PenDeepOrangeA700) Act()             { x.PenPen() }
func (x PenDeepOrangeA700) Ifc() interface{} { return x.PenPen() }
func (x PenDeepOrangeA700) PenPen() pen.Pen  { return pen.DeepOrangeA700 }
func (x PenBrown50) Act()                    { x.PenPen() }
func (x PenBrown50) Ifc() interface{}        { return x.PenPen() }
func (x PenBrown50) PenPen() pen.Pen         { return pen.Brown50 }
func (x PenBrown100) Act()                   { x.PenPen() }
func (x PenBrown100) Ifc() interface{}       { return x.PenPen() }
func (x PenBrown100) PenPen() pen.Pen        { return pen.Brown100 }
func (x PenBrown200) Act()                   { x.PenPen() }
func (x PenBrown200) Ifc() interface{}       { return x.PenPen() }
func (x PenBrown200) PenPen() pen.Pen        { return pen.Brown200 }
func (x PenBrown300) Act()                   { x.PenPen() }
func (x PenBrown300) Ifc() interface{}       { return x.PenPen() }
func (x PenBrown300) PenPen() pen.Pen        { return pen.Brown300 }
func (x PenBrown400) Act()                   { x.PenPen() }
func (x PenBrown400) Ifc() interface{}       { return x.PenPen() }
func (x PenBrown400) PenPen() pen.Pen        { return pen.Brown400 }
func (x PenBrown500) Act()                   { x.PenPen() }
func (x PenBrown500) Ifc() interface{}       { return x.PenPen() }
func (x PenBrown500) PenPen() pen.Pen        { return pen.Brown500 }
func (x PenBrown600) Act()                   { x.PenPen() }
func (x PenBrown600) Ifc() interface{}       { return x.PenPen() }
func (x PenBrown600) PenPen() pen.Pen        { return pen.Brown600 }
func (x PenBrown700) Act()                   { x.PenPen() }
func (x PenBrown700) Ifc() interface{}       { return x.PenPen() }
func (x PenBrown700) PenPen() pen.Pen        { return pen.Brown700 }
func (x PenBrown800) Act()                   { x.PenPen() }
func (x PenBrown800) Ifc() interface{}       { return x.PenPen() }
func (x PenBrown800) PenPen() pen.Pen        { return pen.Brown800 }
func (x PenBrown900) Act()                   { x.PenPen() }
func (x PenBrown900) Ifc() interface{}       { return x.PenPen() }
func (x PenBrown900) PenPen() pen.Pen        { return pen.Brown900 }
func (x PenBlueGrey50) Act()                 { x.PenPen() }
func (x PenBlueGrey50) Ifc() interface{}     { return x.PenPen() }
func (x PenBlueGrey50) PenPen() pen.Pen      { return pen.BlueGrey50 }
func (x PenBlueGrey100) Act()                { x.PenPen() }
func (x PenBlueGrey100) Ifc() interface{}    { return x.PenPen() }
func (x PenBlueGrey100) PenPen() pen.Pen     { return pen.BlueGrey100 }
func (x PenBlueGrey200) Act()                { x.PenPen() }
func (x PenBlueGrey200) Ifc() interface{}    { return x.PenPen() }
func (x PenBlueGrey200) PenPen() pen.Pen     { return pen.BlueGrey200 }
func (x PenBlueGrey300) Act()                { x.PenPen() }
func (x PenBlueGrey300) Ifc() interface{}    { return x.PenPen() }
func (x PenBlueGrey300) PenPen() pen.Pen     { return pen.BlueGrey300 }
func (x PenBlueGrey400) Act()                { x.PenPen() }
func (x PenBlueGrey400) Ifc() interface{}    { return x.PenPen() }
func (x PenBlueGrey400) PenPen() pen.Pen     { return pen.BlueGrey400 }
func (x PenBlueGrey500) Act()                { x.PenPen() }
func (x PenBlueGrey500) Ifc() interface{}    { return x.PenPen() }
func (x PenBlueGrey500) PenPen() pen.Pen     { return pen.BlueGrey500 }
func (x PenBlueGrey600) Act()                { x.PenPen() }
func (x PenBlueGrey600) Ifc() interface{}    { return x.PenPen() }
func (x PenBlueGrey600) PenPen() pen.Pen     { return pen.BlueGrey600 }
func (x PenBlueGrey700) Act()                { x.PenPen() }
func (x PenBlueGrey700) Ifc() interface{}    { return x.PenPen() }
func (x PenBlueGrey700) PenPen() pen.Pen     { return pen.BlueGrey700 }
func (x PenBlueGrey800) Act()                { x.PenPen() }
func (x PenBlueGrey800) Ifc() interface{}    { return x.PenPen() }
func (x PenBlueGrey800) PenPen() pen.Pen     { return pen.BlueGrey800 }
func (x PenBlueGrey900) Act()                { x.PenPen() }
func (x PenBlueGrey900) Ifc() interface{}    { return x.PenPen() }
func (x PenBlueGrey900) PenPen() pen.Pen     { return pen.BlueGrey900 }
func (x PenGrey50) Act()                     { x.PenPen() }
func (x PenGrey50) Ifc() interface{}         { return x.PenPen() }
func (x PenGrey50) PenPen() pen.Pen          { return pen.Grey50 }
func (x PenGrey100) Act()                    { x.PenPen() }
func (x PenGrey100) Ifc() interface{}        { return x.PenPen() }
func (x PenGrey100) PenPen() pen.Pen         { return pen.Grey100 }
func (x PenGrey200) Act()                    { x.PenPen() }
func (x PenGrey200) Ifc() interface{}        { return x.PenPen() }
func (x PenGrey200) PenPen() pen.Pen         { return pen.Grey200 }
func (x PenGrey300) Act()                    { x.PenPen() }
func (x PenGrey300) Ifc() interface{}        { return x.PenPen() }
func (x PenGrey300) PenPen() pen.Pen         { return pen.Grey300 }
func (x PenGrey400) Act()                    { x.PenPen() }
func (x PenGrey400) Ifc() interface{}        { return x.PenPen() }
func (x PenGrey400) PenPen() pen.Pen         { return pen.Grey400 }
func (x PenGrey500) Act()                    { x.PenPen() }
func (x PenGrey500) Ifc() interface{}        { return x.PenPen() }
func (x PenGrey500) PenPen() pen.Pen         { return pen.Grey500 }
func (x PenGrey600) Act()                    { x.PenPen() }
func (x PenGrey600) Ifc() interface{}        { return x.PenPen() }
func (x PenGrey600) PenPen() pen.Pen         { return pen.Grey600 }
func (x PenGrey700) Act()                    { x.PenPen() }
func (x PenGrey700) Ifc() interface{}        { return x.PenPen() }
func (x PenGrey700) PenPen() pen.Pen         { return pen.Grey700 }
func (x PenGrey800) Act()                    { x.PenPen() }
func (x PenGrey800) Ifc() interface{}        { return x.PenPen() }
func (x PenGrey800) PenPen() pen.Pen         { return pen.Grey800 }
func (x PenGrey900) Act()                    { x.PenPen() }
func (x PenGrey900) Ifc() interface{}        { return x.PenPen() }
func (x PenGrey900) PenPen() pen.Pen         { return pen.Grey900 }
func (x FltScl) Act()                        { x.FltFlt() }
func (x FltScl) Ifc() interface{}            { return x.FltFlt() }
func (x FltScl) FltFlt() flt.Flt             { return plt.Scl }
func (x UntStkWidth) Act()                   { x.UntUnt() }
func (x UntStkWidth) Ifc() interface{}       { return x.UntUnt() }
func (x UntStkWidth) UntUnt() unt.Unt        { return plt.StkWidth }
func (x UntShpRadius) Act()                  { x.UntUnt() }
func (x UntShpRadius) Ifc() interface{}      { return x.UntUnt() }
func (x UntShpRadius) UntUnt() unt.Unt       { return plt.ShpRadius }
func (x UntAxisPad) Act()                    { x.UntUnt() }
func (x UntAxisPad) Ifc() interface{}        { return x.UntUnt() }
func (x UntAxisPad) UntUnt() unt.Unt         { return plt.AxisPad }
func (x UntBarPad) Act()                     { x.UntUnt() }
func (x UntBarPad) Ifc() interface{}         { return x.UntUnt() }
func (x UntBarPad) UntUnt() unt.Unt          { return plt.BarPad }
func (x UntLen) Act()                        { x.UntUnt() }
func (x UntLen) Ifc() interface{}            { return x.UntUnt() }
func (x UntLen) UntUnt() unt.Unt             { return plt.Len }
func (x UntPad) Act()                        { x.UntUnt() }
func (x UntPad) Ifc() interface{}            { return x.UntUnt() }
func (x UntPad) UntUnt() unt.Unt             { return plt.Pad }
func (x ClrBakClr) Act()                     { x.ClrClr() }
func (x ClrBakClr) Ifc() interface{}         { return x.ClrClr() }
func (x ClrBakClr) ClrClr() clr.Clr          { return plt.BakClr }
func (x ClrBrdrClr) Act()                    { x.ClrClr() }
func (x ClrBrdrClr) Ifc() interface{}        { return x.ClrClr() }
func (x ClrBrdrClr) ClrClr() clr.Clr         { return plt.BrdrClr }
func (x UntBrdrLen) Act()                    { x.UntUnt() }
func (x UntBrdrLen) Ifc() interface{}        { return x.UntUnt() }
func (x UntBrdrLen) UntUnt() unt.Unt         { return plt.BrdrLen }
func (x UntInrvlTxtLen) Act()                { x.UntUnt() }
func (x UntInrvlTxtLen) Ifc() interface{}    { return x.UntUnt() }
func (x UntInrvlTxtLen) UntUnt() unt.Unt     { return plt.InrvlTxtLen }
func (x ClrInrvlTxtClrX) Act()               { x.ClrClr() }
func (x ClrInrvlTxtClrX) Ifc() interface{}   { return x.ClrClr() }
func (x ClrInrvlTxtClrX) ClrClr() clr.Clr    { return plt.InrvlTxtClrX }
func (x ClrInrvlTxtClrY) Act()               { x.ClrClr() }
func (x ClrInrvlTxtClrY) Ifc() interface{}   { return x.ClrClr() }
func (x ClrInrvlTxtClrY) ClrClr() clr.Clr    { return plt.InrvlTxtClrY }
func (x ClrMsgClr) Act()                     { x.ClrClr() }
func (x ClrMsgClr) Ifc() interface{}         { return x.ClrClr() }
func (x ClrMsgClr) ClrClr() clr.Clr          { return plt.MsgClr }
func (x ClrTitleClr) Act()                   { x.ClrClr() }
func (x ClrTitleClr) Ifc() interface{}       { return x.ClrClr() }
func (x ClrTitleClr) ClrClr() clr.Clr        { return plt.TitleClr }
func (x ClrPrfClr) Act()                     { x.ClrClr() }
func (x ClrPrfClr) Ifc() interface{}         { return x.ClrClr() }
func (x ClrPrfClr) ClrClr() clr.Clr          { return plt.PrfClr }
func (x ClrLosClr) Act()                     { x.ClrClr() }
func (x ClrLosClr) Ifc() interface{}         { return x.ClrClr() }
func (x ClrLosClr) ClrClr() clr.Clr          { return plt.LosClr }
func (x PenPrfPen) Act()                     { x.PenPen() }
func (x PenPrfPen) Ifc() interface{}         { return x.PenPen() }
func (x PenPrfPen) PenPen() pen.Pen          { return plt.PrfPen }
func (x PenLosPen) Act()                     { x.PenPen() }
func (x PenLosPen) Ifc() interface{}         { return x.PenPen() }
func (x PenLosPen) PenPen() pen.Pen          { return plt.LosPen }
func (x FltOutlierLim) Act()                 { x.FltFlt() }
func (x FltOutlierLim) Ifc() interface{}     { return x.FltFlt() }
func (x FltOutlierLim) FltFlt() flt.Flt      { return plt.OutlierLim }
func (x StrIfo) Act()                        { x.StrStr() }
func (x StrIfo) Ifc() interface{}            { return x.StrStr() }
func (x StrIfo) StrStr() str.Str {
	var i0 []interface{}
	for _, cur := range x.I0 {
		i0 = append(i0, cur.Ifc())
	}
	return log.Ifo(i0...)
}
func (x StrIfof) Act()             { x.StrStr() }
func (x StrIfof) Ifc() interface{} { return x.StrStr() }
func (x StrIfof) StrStr() str.Str {
	var i1 []interface{}
	for _, cur := range x.I1 {
		i1 = append(i1, cur.Ifc())
	}
	return log.Ifof(x.I0.StrStr(), i1...)
}
func (x StrFmt) Act()             { x.StrStr() }
func (x StrFmt) Ifc() interface{} { return x.StrStr() }
func (x StrFmt) StrStr() str.Str {
	var i1 []interface{}
	for _, cur := range x.I1 {
		i1 = append(i1, cur.Ifc())
	}
	return str.Fmt(x.I0.StrStr(), i1...)
}
func (x TmeNow) Act()                    { x.TmeTme() }
func (x TmeNow) Ifc() interface{}        { return x.TmeTme() }
func (x TmeNow) TmeTme() tme.Tme         { return tme.Now() }
func (x FltNewRng) Act()                 { x.FltRng() }
func (x FltNewRng) Ifc() interface{}     { return x.FltRng() }
func (x FltNewRng) FltRng() flt.Rng      { return flt.NewRng(x.I0.FltFlt(), x.I1.FltFlt()) }
func (x FltNewRngArnd) Act()             { x.FltRng() }
func (x FltNewRngArnd) Ifc() interface{} { return x.FltRng() }
func (x FltNewRngArnd) FltRng() flt.Rng  { return flt.NewRngArnd(x.I0.FltFlt(), x.I1.FltFlt()) }
func (x FltNewRngFul) Act()              { x.FltRng() }
func (x FltNewRngFul) Ifc() interface{}  { return x.FltRng() }
func (x FltNewRngFul) FltRng() flt.Rng   { return flt.NewRngFul() }
func (x TmeNewRng) Act()                 { x.TmeRng() }
func (x TmeNewRng) Ifc() interface{}     { return x.TmeRng() }
func (x TmeNewRng) TmeRng() tme.Rng      { return tme.NewRng(x.I0.TmeTme(), x.I1.TmeTme()) }
func (x TmeNewRngArnd) Act()             { x.TmeRng() }
func (x TmeNewRngArnd) Ifc() interface{} { return x.TmeRng() }
func (x TmeNewRngArnd) TmeRng() tme.Rng  { return tme.NewRngArnd(x.I0.TmeTme(), x.I1.TmeTme()) }
func (x TmeNewRngFul) Act()              { x.TmeRng() }
func (x TmeNewRngFul) Ifc() interface{}  { return x.TmeRng() }
func (x TmeNewRngFul) TmeRng() tme.Rng   { return tme.NewRngFul() }
func (x StrsNew) Act()                   { x.StrsStrs() }
func (x StrsNew) Ifc() interface{}       { return x.StrsStrs() }
func (x StrsNew) StrsStrs() *strs.Strs {
	var i0 []str.Str
	for _, cur := range x.I0 {
		i0 = append(i0, cur.StrStr())
	}
	return strs.New(i0...)
}
func (x StrsMake) Act()                    { x.StrsStrs() }
func (x StrsMake) Ifc() interface{}        { return x.StrsStrs() }
func (x StrsMake) StrsStrs() *strs.Strs    { return strs.Make(x.I0.UntUnt()) }
func (x StrsMakeEmp) Act()                 { x.StrsStrs() }
func (x StrsMakeEmp) Ifc() interface{}     { return x.StrsStrs() }
func (x StrsMakeEmp) StrsStrs() *strs.Strs { return strs.MakeEmp(x.I0.UntUnt()) }
func (x BolsNew) Act()                     { x.BolsBols() }
func (x BolsNew) Ifc() interface{}         { return x.BolsBols() }
func (x BolsNew) BolsBols() *bols.Bols {
	var i0 []bol.Bol
	for _, cur := range x.I0 {
		i0 = append(i0, cur.BolBol())
	}
	return bols.New(i0...)
}
func (x BolsMake) Act()                    { x.BolsBols() }
func (x BolsMake) Ifc() interface{}        { return x.BolsBols() }
func (x BolsMake) BolsBols() *bols.Bols    { return bols.Make(x.I0.UntUnt()) }
func (x BolsMakeEmp) Act()                 { x.BolsBols() }
func (x BolsMakeEmp) Ifc() interface{}     { return x.BolsBols() }
func (x BolsMakeEmp) BolsBols() *bols.Bols { return bols.MakeEmp(x.I0.UntUnt()) }
func (x FltsNew) Act()                     { x.FltsFlts() }
func (x FltsNew) Ifc() interface{}         { return x.FltsFlts() }
func (x FltsNew) FltsFlts() *flts.Flts {
	var i0 []flt.Flt
	for _, cur := range x.I0 {
		i0 = append(i0, cur.FltFlt())
	}
	return flts.New(i0...)
}
func (x FltsMake) Act()                    { x.FltsFlts() }
func (x FltsMake) Ifc() interface{}        { return x.FltsFlts() }
func (x FltsMake) FltsFlts() *flts.Flts    { return flts.Make(x.I0.UntUnt()) }
func (x FltsMakeEmp) Act()                 { x.FltsFlts() }
func (x FltsMakeEmp) Ifc() interface{}     { return x.FltsFlts() }
func (x FltsMakeEmp) FltsFlts() *flts.Flts { return flts.MakeEmp(x.I0.UntUnt()) }
func (x FltsAddsLss) Act()                 { x.FltsFlts() }
func (x FltsAddsLss) Ifc() interface{}     { return x.FltsFlts() }
func (x FltsAddsLss) FltsFlts() *flts.Flts {
	return flts.AddsLss(x.I0.FltFlt(), x.I1.FltFlt(), x.I2.FltFlt())
}
func (x FltsAddsLeq) Act()             { x.FltsFlts() }
func (x FltsAddsLeq) Ifc() interface{} { return x.FltsFlts() }
func (x FltsAddsLeq) FltsFlts() *flts.Flts {
	return flts.AddsLeq(x.I0.FltFlt(), x.I1.FltFlt(), x.I2.FltFlt())
}
func (x FltsSubsGtr) Act()             { x.FltsFlts() }
func (x FltsSubsGtr) Ifc() interface{} { return x.FltsFlts() }
func (x FltsSubsGtr) FltsFlts() *flts.Flts {
	return flts.SubsGtr(x.I0.FltFlt(), x.I1.FltFlt(), x.I2.FltFlt())
}
func (x FltsSubsGeq) Act()             { x.FltsFlts() }
func (x FltsSubsGeq) Ifc() interface{} { return x.FltsFlts() }
func (x FltsSubsGeq) FltsFlts() *flts.Flts {
	return flts.SubsGeq(x.I0.FltFlt(), x.I1.FltFlt(), x.I2.FltFlt())
}
func (x FltsMulsLss) Act()             { x.FltsFlts() }
func (x FltsMulsLss) Ifc() interface{} { return x.FltsFlts() }
func (x FltsMulsLss) FltsFlts() *flts.Flts {
	return flts.MulsLss(x.I0.FltFlt(), x.I1.FltFlt(), x.I2.FltFlt())
}
func (x FltsMulsLeq) Act()             { x.FltsFlts() }
func (x FltsMulsLeq) Ifc() interface{} { return x.FltsFlts() }
func (x FltsMulsLeq) FltsFlts() *flts.Flts {
	return flts.MulsLeq(x.I0.FltFlt(), x.I1.FltFlt(), x.I2.FltFlt())
}
func (x FltsDivsGtr) Act()             { x.FltsFlts() }
func (x FltsDivsGtr) Ifc() interface{} { return x.FltsFlts() }
func (x FltsDivsGtr) FltsFlts() *flts.Flts {
	return flts.DivsGtr(x.I0.FltFlt(), x.I1.FltFlt(), x.I2.FltFlt())
}
func (x FltsDivsGeq) Act()             { x.FltsFlts() }
func (x FltsDivsGeq) Ifc() interface{} { return x.FltsFlts() }
func (x FltsDivsGeq) FltsFlts() *flts.Flts {
	return flts.DivsGeq(x.I0.FltFlt(), x.I1.FltFlt(), x.I2.FltFlt())
}
func (x FltsFibsLeq) Act()                 { x.FltsFlts() }
func (x FltsFibsLeq) Ifc() interface{}     { return x.FltsFlts() }
func (x FltsFibsLeq) FltsFlts() *flts.Flts { return flts.FibsLeq(x.I0.FltFlt()) }
func (x UntsNew) Act()                     { x.UntsUnts() }
func (x UntsNew) Ifc() interface{}         { return x.UntsUnts() }
func (x UntsNew) UntsUnts() *unts.Unts {
	var i0 []unt.Unt
	for _, cur := range x.I0 {
		i0 = append(i0, cur.UntUnt())
	}
	return unts.New(i0...)
}
func (x UntsMake) Act()                    { x.UntsUnts() }
func (x UntsMake) Ifc() interface{}        { return x.UntsUnts() }
func (x UntsMake) UntsUnts() *unts.Unts    { return unts.Make(x.I0.UntUnt()) }
func (x UntsMakeEmp) Act()                 { x.UntsUnts() }
func (x UntsMakeEmp) Ifc() interface{}     { return x.UntsUnts() }
func (x UntsMakeEmp) UntsUnts() *unts.Unts { return unts.MakeEmp(x.I0.UntUnt()) }
func (x UntsAddsLss) Act()                 { x.UntsUnts() }
func (x UntsAddsLss) Ifc() interface{}     { return x.UntsUnts() }
func (x UntsAddsLss) UntsUnts() *unts.Unts {
	return unts.AddsLss(x.I0.UntUnt(), x.I1.UntUnt(), x.I2.UntUnt())
}
func (x UntsAddsLeq) Act()             { x.UntsUnts() }
func (x UntsAddsLeq) Ifc() interface{} { return x.UntsUnts() }
func (x UntsAddsLeq) UntsUnts() *unts.Unts {
	return unts.AddsLeq(x.I0.UntUnt(), x.I1.UntUnt(), x.I2.UntUnt())
}
func (x UntsSubsGtr) Act()             { x.UntsUnts() }
func (x UntsSubsGtr) Ifc() interface{} { return x.UntsUnts() }
func (x UntsSubsGtr) UntsUnts() *unts.Unts {
	return unts.SubsGtr(x.I0.UntUnt(), x.I1.UntUnt(), x.I2.UntUnt())
}
func (x UntsSubsGeq) Act()             { x.UntsUnts() }
func (x UntsSubsGeq) Ifc() interface{} { return x.UntsUnts() }
func (x UntsSubsGeq) UntsUnts() *unts.Unts {
	return unts.SubsGeq(x.I0.UntUnt(), x.I1.UntUnt(), x.I2.UntUnt())
}
func (x UntsMulsLss) Act()             { x.UntsUnts() }
func (x UntsMulsLss) Ifc() interface{} { return x.UntsUnts() }
func (x UntsMulsLss) UntsUnts() *unts.Unts {
	return unts.MulsLss(x.I0.UntUnt(), x.I1.UntUnt(), x.I2.UntUnt())
}
func (x UntsMulsLeq) Act()             { x.UntsUnts() }
func (x UntsMulsLeq) Ifc() interface{} { return x.UntsUnts() }
func (x UntsMulsLeq) UntsUnts() *unts.Unts {
	return unts.MulsLeq(x.I0.UntUnt(), x.I1.UntUnt(), x.I2.UntUnt())
}
func (x UntsDivsGtr) Act()             { x.UntsUnts() }
func (x UntsDivsGtr) Ifc() interface{} { return x.UntsUnts() }
func (x UntsDivsGtr) UntsUnts() *unts.Unts {
	return unts.DivsGtr(x.I0.UntUnt(), x.I1.UntUnt(), x.I2.UntUnt())
}
func (x UntsDivsGeq) Act()             { x.UntsUnts() }
func (x UntsDivsGeq) Ifc() interface{} { return x.UntsUnts() }
func (x UntsDivsGeq) UntsUnts() *unts.Unts {
	return unts.DivsGeq(x.I0.UntUnt(), x.I1.UntUnt(), x.I2.UntUnt())
}
func (x UntsFibsLeq) Act()                 { x.UntsUnts() }
func (x UntsFibsLeq) Ifc() interface{}     { return x.UntsUnts() }
func (x UntsFibsLeq) UntsUnts() *unts.Unts { return unts.FibsLeq(x.I0.UntUnt()) }
func (x IntsNew) Act()                     { x.IntsInts() }
func (x IntsNew) Ifc() interface{}         { return x.IntsInts() }
func (x IntsNew) IntsInts() *ints.Ints {
	var i0 []int.Int
	for _, cur := range x.I0 {
		i0 = append(i0, cur.IntInt())
	}
	return ints.New(i0...)
}
func (x IntsMake) Act()                    { x.IntsInts() }
func (x IntsMake) Ifc() interface{}        { return x.IntsInts() }
func (x IntsMake) IntsInts() *ints.Ints    { return ints.Make(x.I0.UntUnt()) }
func (x IntsMakeEmp) Act()                 { x.IntsInts() }
func (x IntsMakeEmp) Ifc() interface{}     { return x.IntsInts() }
func (x IntsMakeEmp) IntsInts() *ints.Ints { return ints.MakeEmp(x.I0.UntUnt()) }
func (x TmesNew) Act()                     { x.TmesTmes() }
func (x TmesNew) Ifc() interface{}         { return x.TmesTmes() }
func (x TmesNew) TmesTmes() *tmes.Tmes {
	var i0 []tme.Tme
	for _, cur := range x.I0 {
		i0 = append(i0, cur.TmeTme())
	}
	return tmes.New(i0...)
}
func (x TmesMake) Act()                    { x.TmesTmes() }
func (x TmesMake) Ifc() interface{}        { return x.TmesTmes() }
func (x TmesMake) TmesTmes() *tmes.Tmes    { return tmes.Make(x.I0.UntUnt()) }
func (x TmesMakeEmp) Act()                 { x.TmesTmes() }
func (x TmesMakeEmp) Ifc() interface{}     { return x.TmesTmes() }
func (x TmesMakeEmp) TmesTmes() *tmes.Tmes { return tmes.MakeEmp(x.I0.UntUnt()) }
func (x TmesAddsLss) Act()                 { x.TmesTmes() }
func (x TmesAddsLss) Ifc() interface{}     { return x.TmesTmes() }
func (x TmesAddsLss) TmesTmes() *tmes.Tmes {
	return tmes.AddsLss(x.I0.TmeTme(), x.I1.TmeTme(), x.I2.TmeTme())
}
func (x TmesAddsLeq) Act()             { x.TmesTmes() }
func (x TmesAddsLeq) Ifc() interface{} { return x.TmesTmes() }
func (x TmesAddsLeq) TmesTmes() *tmes.Tmes {
	return tmes.AddsLeq(x.I0.TmeTme(), x.I1.TmeTme(), x.I2.TmeTme())
}
func (x TmesSubsGtr) Act()             { x.TmesTmes() }
func (x TmesSubsGtr) Ifc() interface{} { return x.TmesTmes() }
func (x TmesSubsGtr) TmesTmes() *tmes.Tmes {
	return tmes.SubsGtr(x.I0.TmeTme(), x.I1.TmeTme(), x.I2.TmeTme())
}
func (x TmesSubsGeq) Act()             { x.TmesTmes() }
func (x TmesSubsGeq) Ifc() interface{} { return x.TmesTmes() }
func (x TmesSubsGeq) TmesTmes() *tmes.Tmes {
	return tmes.SubsGeq(x.I0.TmeTme(), x.I1.TmeTme(), x.I2.TmeTme())
}
func (x TmesMulsLss) Act()             { x.TmesTmes() }
func (x TmesMulsLss) Ifc() interface{} { return x.TmesTmes() }
func (x TmesMulsLss) TmesTmes() *tmes.Tmes {
	return tmes.MulsLss(x.I0.TmeTme(), x.I1.TmeTme(), x.I2.TmeTme())
}
func (x TmesMulsLeq) Act()             { x.TmesTmes() }
func (x TmesMulsLeq) Ifc() interface{} { return x.TmesTmes() }
func (x TmesMulsLeq) TmesTmes() *tmes.Tmes {
	return tmes.MulsLeq(x.I0.TmeTme(), x.I1.TmeTme(), x.I2.TmeTme())
}
func (x TmesDivsGtr) Act()             { x.TmesTmes() }
func (x TmesDivsGtr) Ifc() interface{} { return x.TmesTmes() }
func (x TmesDivsGtr) TmesTmes() *tmes.Tmes {
	return tmes.DivsGtr(x.I0.TmeTme(), x.I1.TmeTme(), x.I2.TmeTme())
}
func (x TmesDivsGeq) Act()             { x.TmesTmes() }
func (x TmesDivsGeq) Ifc() interface{} { return x.TmesTmes() }
func (x TmesDivsGeq) TmesTmes() *tmes.Tmes {
	return tmes.DivsGeq(x.I0.TmeTme(), x.I1.TmeTme(), x.I2.TmeTme())
}
func (x TmesFibsLeq) Act()                 { x.TmesTmes() }
func (x TmesFibsLeq) Ifc() interface{}     { return x.TmesTmes() }
func (x TmesFibsLeq) TmesTmes() *tmes.Tmes { return tmes.FibsLeq(x.I0.TmeTme()) }
func (x BndsNew) Act()                     { x.BndsBnds() }
func (x BndsNew) Ifc() interface{}         { return x.BndsBnds() }
func (x BndsNew) BndsBnds() *bnds.Bnds {
	var i0 []bnd.Bnd
	for _, cur := range x.I0 {
		i0 = append(i0, cur.BndBnd())
	}
	return bnds.New(i0...)
}
func (x BndsMake) Act()                    { x.BndsBnds() }
func (x BndsMake) Ifc() interface{}        { return x.BndsBnds() }
func (x BndsMake) BndsBnds() *bnds.Bnds    { return bnds.Make(x.I0.UntUnt()) }
func (x BndsMakeEmp) Act()                 { x.BndsBnds() }
func (x BndsMakeEmp) Ifc() interface{}     { return x.BndsBnds() }
func (x BndsMakeEmp) BndsBnds() *bnds.Bnds { return bnds.MakeEmp(x.I0.UntUnt()) }
func (x TmeNewRngs) Act()                  { x.TmeRngs() }
func (x TmeNewRngs) Ifc() interface{}      { return x.TmeRngs() }
func (x TmeNewRngs) TmeRngs() *tme.Rngs {
	var i0 []tme.Rng
	for _, cur := range x.I0 {
		i0 = append(i0, cur.TmeRng())
	}
	return tme.NewRngs(i0...)
}
func (x TmeMakeRngs) Act()                  { x.TmeRngs() }
func (x TmeMakeRngs) Ifc() interface{}      { return x.TmeRngs() }
func (x TmeMakeRngs) TmeRngs() *tme.Rngs    { return tme.MakeRngs(x.I0.UntUnt()) }
func (x TmeMakeEmpRngs) Act()               { x.TmeRngs() }
func (x TmeMakeEmpRngs) Ifc() interface{}   { return x.TmeRngs() }
func (x TmeMakeEmpRngs) TmeRngs() *tme.Rngs { return tme.MakeEmpRngs(x.I0.UntUnt()) }
func (x AnaNewTrds) Act()                   { x.AnaTrds() }
func (x AnaNewTrds) Ifc() interface{}       { return x.AnaTrds() }
func (x AnaNewTrds) AnaTrds() *ana.Trds {
	var i0 []*ana.Trd
	for _, cur := range x.I0 {
		i0 = append(i0, cur.AnaTrd())
	}
	return ana.NewTrds(i0...)
}
func (x AnaMakeTrds) Act()                  { x.AnaTrds() }
func (x AnaMakeTrds) Ifc() interface{}      { return x.AnaTrds() }
func (x AnaMakeTrds) AnaTrds() *ana.Trds    { return ana.MakeTrds(x.I0.UntUnt()) }
func (x AnaMakeEmpTrds) Act()               { x.AnaTrds() }
func (x AnaMakeEmpTrds) Ifc() interface{}   { return x.AnaTrds() }
func (x AnaMakeEmpTrds) AnaTrds() *ana.Trds { return ana.MakeEmpTrds(x.I0.UntUnt()) }
func (x AnaNewPrfms) Act()                  { x.AnaPrfms() }
func (x AnaNewPrfms) Ifc() interface{}      { return x.AnaPrfms() }
func (x AnaNewPrfms) AnaPrfms() *ana.Prfms {
	var i0 []*ana.Prfm
	for _, cur := range x.I0 {
		i0 = append(i0, cur.AnaPrfm())
	}
	return ana.NewPrfms(i0...)
}
func (x AnaMakePrfms) Act()                    { x.AnaPrfms() }
func (x AnaMakePrfms) Ifc() interface{}        { return x.AnaPrfms() }
func (x AnaMakePrfms) AnaPrfms() *ana.Prfms    { return ana.MakePrfms(x.I0.UntUnt()) }
func (x AnaMakeEmpPrfms) Act()                 { x.AnaPrfms() }
func (x AnaMakeEmpPrfms) Ifc() interface{}     { return x.AnaPrfms() }
func (x AnaMakeEmpPrfms) AnaPrfms() *ana.Prfms { return ana.MakeEmpPrfms(x.I0.UntUnt()) }
func (x HstOan) Act()                          { x.HstPrv() }
func (x HstOan) Ifc() interface{}              { return x.HstPrv() }
func (x HstOan) HstPrv() hst.Prv               { return hst.Oan() }
func (x HstNewPrvs) Act()                      { x.HstPrvs() }
func (x HstNewPrvs) Ifc() interface{}          { return x.HstPrvs() }
func (x HstNewPrvs) HstPrvs() *hst.Prvs {
	var i0 []hst.Prv
	for _, cur := range x.I0 {
		i0 = append(i0, cur.HstPrv())
	}
	return hst.NewPrvs(i0...)
}
func (x HstMakePrvs) Act()                  { x.HstPrvs() }
func (x HstMakePrvs) Ifc() interface{}      { return x.HstPrvs() }
func (x HstMakePrvs) HstPrvs() *hst.Prvs    { return hst.MakePrvs(x.I0.UntUnt()) }
func (x HstMakeEmpPrvs) Act()               { x.HstPrvs() }
func (x HstMakeEmpPrvs) Ifc() interface{}   { return x.HstPrvs() }
func (x HstMakeEmpPrvs) HstPrvs() *hst.Prvs { return hst.MakeEmpPrvs(x.I0.UntUnt()) }
func (x HstNewInstrs) Act()                 { x.HstInstrs() }
func (x HstNewInstrs) Ifc() interface{}     { return x.HstInstrs() }
func (x HstNewInstrs) HstInstrs() *hst.Instrs {
	var i0 []hst.Instr
	for _, cur := range x.I0 {
		i0 = append(i0, cur.HstInstr())
	}
	return hst.NewInstrs(i0...)
}
func (x HstMakeInstrs) Act()                      { x.HstInstrs() }
func (x HstMakeInstrs) Ifc() interface{}          { return x.HstInstrs() }
func (x HstMakeInstrs) HstInstrs() *hst.Instrs    { return hst.MakeInstrs(x.I0.UntUnt()) }
func (x HstMakeEmpInstrs) Act()                   { x.HstInstrs() }
func (x HstMakeEmpInstrs) Ifc() interface{}       { return x.HstInstrs() }
func (x HstMakeEmpInstrs) HstInstrs() *hst.Instrs { return hst.MakeEmpInstrs(x.I0.UntUnt()) }
func (x HstNewInrvls) Act()                       { x.HstInrvls() }
func (x HstNewInrvls) Ifc() interface{}           { return x.HstInrvls() }
func (x HstNewInrvls) HstInrvls() *hst.Inrvls {
	var i0 []hst.Inrvl
	for _, cur := range x.I0 {
		i0 = append(i0, cur.HstInrvl())
	}
	return hst.NewInrvls(i0...)
}
func (x HstMakeInrvls) Act()                      { x.HstInrvls() }
func (x HstMakeInrvls) Ifc() interface{}          { return x.HstInrvls() }
func (x HstMakeInrvls) HstInrvls() *hst.Inrvls    { return hst.MakeInrvls(x.I0.UntUnt()) }
func (x HstMakeEmpInrvls) Act()                   { x.HstInrvls() }
func (x HstMakeEmpInrvls) Ifc() interface{}       { return x.HstInrvls() }
func (x HstMakeEmpInrvls) HstInrvls() *hst.Inrvls { return hst.MakeEmpInrvls(x.I0.UntUnt()) }
func (x HstNewSides) Act()                        { x.HstSides() }
func (x HstNewSides) Ifc() interface{}            { return x.HstSides() }
func (x HstNewSides) HstSides() *hst.Sides {
	var i0 []hst.Side
	for _, cur := range x.I0 {
		i0 = append(i0, cur.HstSide())
	}
	return hst.NewSides(i0...)
}
func (x HstMakeSides) Act()                    { x.HstSides() }
func (x HstMakeSides) Ifc() interface{}        { return x.HstSides() }
func (x HstMakeSides) HstSides() *hst.Sides    { return hst.MakeSides(x.I0.UntUnt()) }
func (x HstMakeEmpSides) Act()                 { x.HstSides() }
func (x HstMakeEmpSides) Ifc() interface{}     { return x.HstSides() }
func (x HstMakeEmpSides) HstSides() *hst.Sides { return hst.MakeEmpSides(x.I0.UntUnt()) }
func (x HstNewStms) Act()                      { x.HstStms() }
func (x HstNewStms) Ifc() interface{}          { return x.HstStms() }
func (x HstNewStms) HstStms() *hst.Stms {
	var i0 []hst.Stm
	for _, cur := range x.I0 {
		i0 = append(i0, cur.HstStm())
	}
	return hst.NewStms(i0...)
}
func (x HstMakeStms) Act()                  { x.HstStms() }
func (x HstMakeStms) Ifc() interface{}      { return x.HstStms() }
func (x HstMakeStms) HstStms() *hst.Stms    { return hst.MakeStms(x.I0.UntUnt()) }
func (x HstMakeEmpStms) Act()               { x.HstStms() }
func (x HstMakeEmpStms) Ifc() interface{}   { return x.HstStms() }
func (x HstMakeEmpStms) HstStms() *hst.Stms { return hst.MakeEmpStms(x.I0.UntUnt()) }
func (x HstNewCnds) Act()                   { x.HstCnds() }
func (x HstNewCnds) Ifc() interface{}       { return x.HstCnds() }
func (x HstNewCnds) HstCnds() *hst.Cnds {
	var i0 []hst.Cnd
	for _, cur := range x.I0 {
		i0 = append(i0, cur.HstCnd())
	}
	return hst.NewCnds(i0...)
}
func (x HstMakeCnds) Act()                  { x.HstCnds() }
func (x HstMakeCnds) Ifc() interface{}      { return x.HstCnds() }
func (x HstMakeCnds) HstCnds() *hst.Cnds    { return hst.MakeCnds(x.I0.UntUnt()) }
func (x HstMakeEmpCnds) Act()               { x.HstCnds() }
func (x HstMakeEmpCnds) Ifc() interface{}   { return x.HstCnds() }
func (x HstMakeEmpCnds) HstCnds() *hst.Cnds { return hst.MakeEmpCnds(x.I0.UntUnt()) }
func (x HstNewStgys) Act()                  { x.HstStgys() }
func (x HstNewStgys) Ifc() interface{}      { return x.HstStgys() }
func (x HstNewStgys) HstStgys() *hst.Stgys {
	var i0 []hst.Stgy
	for _, cur := range x.I0 {
		i0 = append(i0, cur.HstStgy())
	}
	return hst.NewStgys(i0...)
}
func (x HstMakeStgys) Act()                    { x.HstStgys() }
func (x HstMakeStgys) Ifc() interface{}        { return x.HstStgys() }
func (x HstMakeStgys) HstStgys() *hst.Stgys    { return hst.MakeStgys(x.I0.UntUnt()) }
func (x HstMakeEmpStgys) Act()                 { x.HstStgys() }
func (x HstMakeEmpStgys) Ifc() interface{}     { return x.HstStgys() }
func (x HstMakeEmpStgys) HstStgys() *hst.Stgys { return hst.MakeEmpStgys(x.I0.UntUnt()) }
func (x RltOan) Act()                          { x.RltPrv() }
func (x RltOan) Ifc() interface{}              { return x.RltPrv() }
func (x RltOan) RltPrv() rlt.Prv               { return rlt.Oan() }
func (x RltNewPrvs) Act()                      { x.RltPrvs() }
func (x RltNewPrvs) Ifc() interface{}          { return x.RltPrvs() }
func (x RltNewPrvs) RltPrvs() *rlt.Prvs {
	var i0 []rlt.Prv
	for _, cur := range x.I0 {
		i0 = append(i0, cur.RltPrv())
	}
	return rlt.NewPrvs(i0...)
}
func (x RltMakePrvs) Act()                  { x.RltPrvs() }
func (x RltMakePrvs) Ifc() interface{}      { return x.RltPrvs() }
func (x RltMakePrvs) RltPrvs() *rlt.Prvs    { return rlt.MakePrvs(x.I0.UntUnt()) }
func (x RltMakeEmpPrvs) Act()               { x.RltPrvs() }
func (x RltMakeEmpPrvs) Ifc() interface{}   { return x.RltPrvs() }
func (x RltMakeEmpPrvs) RltPrvs() *rlt.Prvs { return rlt.MakeEmpPrvs(x.I0.UntUnt()) }
func (x RltNewInstrs) Act()                 { x.RltInstrs() }
func (x RltNewInstrs) Ifc() interface{}     { return x.RltInstrs() }
func (x RltNewInstrs) RltInstrs() *rlt.Instrs {
	var i0 []rlt.Instr
	for _, cur := range x.I0 {
		i0 = append(i0, cur.RltInstr())
	}
	return rlt.NewInstrs(i0...)
}
func (x RltMakeInstrs) Act()                      { x.RltInstrs() }
func (x RltMakeInstrs) Ifc() interface{}          { return x.RltInstrs() }
func (x RltMakeInstrs) RltInstrs() *rlt.Instrs    { return rlt.MakeInstrs(x.I0.UntUnt()) }
func (x RltMakeEmpInstrs) Act()                   { x.RltInstrs() }
func (x RltMakeEmpInstrs) Ifc() interface{}       { return x.RltInstrs() }
func (x RltMakeEmpInstrs) RltInstrs() *rlt.Instrs { return rlt.MakeEmpInstrs(x.I0.UntUnt()) }
func (x RltNewInrvls) Act()                       { x.RltInrvls() }
func (x RltNewInrvls) Ifc() interface{}           { return x.RltInrvls() }
func (x RltNewInrvls) RltInrvls() *rlt.Inrvls {
	var i0 []rlt.Inrvl
	for _, cur := range x.I0 {
		i0 = append(i0, cur.RltInrvl())
	}
	return rlt.NewInrvls(i0...)
}
func (x RltMakeInrvls) Act()                      { x.RltInrvls() }
func (x RltMakeInrvls) Ifc() interface{}          { return x.RltInrvls() }
func (x RltMakeInrvls) RltInrvls() *rlt.Inrvls    { return rlt.MakeInrvls(x.I0.UntUnt()) }
func (x RltMakeEmpInrvls) Act()                   { x.RltInrvls() }
func (x RltMakeEmpInrvls) Ifc() interface{}       { return x.RltInrvls() }
func (x RltMakeEmpInrvls) RltInrvls() *rlt.Inrvls { return rlt.MakeEmpInrvls(x.I0.UntUnt()) }
func (x RltNewSides) Act()                        { x.RltSides() }
func (x RltNewSides) Ifc() interface{}            { return x.RltSides() }
func (x RltNewSides) RltSides() *rlt.Sides {
	var i0 []rlt.Side
	for _, cur := range x.I0 {
		i0 = append(i0, cur.RltSide())
	}
	return rlt.NewSides(i0...)
}
func (x RltMakeSides) Act()                    { x.RltSides() }
func (x RltMakeSides) Ifc() interface{}        { return x.RltSides() }
func (x RltMakeSides) RltSides() *rlt.Sides    { return rlt.MakeSides(x.I0.UntUnt()) }
func (x RltMakeEmpSides) Act()                 { x.RltSides() }
func (x RltMakeEmpSides) Ifc() interface{}     { return x.RltSides() }
func (x RltMakeEmpSides) RltSides() *rlt.Sides { return rlt.MakeEmpSides(x.I0.UntUnt()) }
func (x RltNewStms) Act()                      { x.RltStms() }
func (x RltNewStms) Ifc() interface{}          { return x.RltStms() }
func (x RltNewStms) RltStms() *rlt.Stms {
	var i0 []rlt.Stm
	for _, cur := range x.I0 {
		i0 = append(i0, cur.RltStm())
	}
	return rlt.NewStms(i0...)
}
func (x RltMakeStms) Act()                  { x.RltStms() }
func (x RltMakeStms) Ifc() interface{}      { return x.RltStms() }
func (x RltMakeStms) RltStms() *rlt.Stms    { return rlt.MakeStms(x.I0.UntUnt()) }
func (x RltMakeEmpStms) Act()               { x.RltStms() }
func (x RltMakeEmpStms) Ifc() interface{}   { return x.RltStms() }
func (x RltMakeEmpStms) RltStms() *rlt.Stms { return rlt.MakeEmpStms(x.I0.UntUnt()) }
func (x RltNewCnds) Act()                   { x.RltCnds() }
func (x RltNewCnds) Ifc() interface{}       { return x.RltCnds() }
func (x RltNewCnds) RltCnds() *rlt.Cnds {
	var i0 []rlt.Cnd
	for _, cur := range x.I0 {
		i0 = append(i0, cur.RltCnd())
	}
	return rlt.NewCnds(i0...)
}
func (x RltMakeCnds) Act()                  { x.RltCnds() }
func (x RltMakeCnds) Ifc() interface{}      { return x.RltCnds() }
func (x RltMakeCnds) RltCnds() *rlt.Cnds    { return rlt.MakeCnds(x.I0.UntUnt()) }
func (x RltMakeEmpCnds) Act()               { x.RltCnds() }
func (x RltMakeEmpCnds) Ifc() interface{}   { return x.RltCnds() }
func (x RltMakeEmpCnds) RltCnds() *rlt.Cnds { return rlt.MakeEmpCnds(x.I0.UntUnt()) }
func (x RltNewStgys) Act()                  { x.RltStgys() }
func (x RltNewStgys) Ifc() interface{}      { return x.RltStgys() }
func (x RltNewStgys) RltStgys() *rlt.Stgys {
	var i0 []rlt.Stgy
	for _, cur := range x.I0 {
		i0 = append(i0, cur.RltStgy())
	}
	return rlt.NewStgys(i0...)
}
func (x RltMakeStgys) Act()                    { x.RltStgys() }
func (x RltMakeStgys) Ifc() interface{}        { return x.RltStgys() }
func (x RltMakeStgys) RltStgys() *rlt.Stgys    { return rlt.MakeStgys(x.I0.UntUnt()) }
func (x RltMakeEmpStgys) Act()                 { x.RltStgys() }
func (x RltMakeEmpStgys) Ifc() interface{}     { return x.RltStgys() }
func (x RltMakeEmpStgys) RltStgys() *rlt.Stgys { return rlt.MakeEmpStgys(x.I0.UntUnt()) }
func (x ClrRgba) Act()                         { x.ClrClr() }
func (x ClrRgba) Ifc() interface{}             { return x.ClrClr() }
func (x ClrRgba) ClrClr() clr.Clr {
	return clr.Rgba(x.I0.FltFlt(), x.I1.FltFlt(), x.I2.FltFlt(), x.I3.FltFlt())
}
func (x ClrRgb) Act()             { x.ClrClr() }
func (x ClrRgb) Ifc() interface{} { return x.ClrClr() }
func (x ClrRgb) ClrClr() clr.Clr  { return clr.Rgb(x.I0.FltFlt(), x.I1.FltFlt(), x.I2.FltFlt()) }
func (x ClrHex) Act()             { x.ClrClr() }
func (x ClrHex) Ifc() interface{} { return x.ClrClr() }
func (x ClrHex) ClrClr() clr.Clr  { return clr.Hex(x.I0.StrStr()) }
func (x PenNew) Act()             { x.PenPen() }
func (x PenNew) Ifc() interface{} { return x.PenPen() }
func (x PenNew) PenPen() pen.Pen {
	var i1 []unt.Unt
	for _, cur := range x.I1 {
		i1 = append(i1, cur.UntUnt())
	}
	return pen.New(x.I0.ClrClr(), i1...)
}
func (x PenRgba) Act()             { x.PenPen() }
func (x PenRgba) Ifc() interface{} { return x.PenPen() }
func (x PenRgba) PenPen() pen.Pen {
	var i4 []unt.Unt
	for _, cur := range x.I4 {
		i4 = append(i4, cur.UntUnt())
	}
	return pen.Rgba(x.I0.FltFlt(), x.I1.FltFlt(), x.I2.FltFlt(), x.I3.FltFlt(), i4...)
}
func (x PenRgb) Act()             { x.PenPen() }
func (x PenRgb) Ifc() interface{} { return x.PenPen() }
func (x PenRgb) PenPen() pen.Pen {
	var i3 []unt.Unt
	for _, cur := range x.I3 {
		i3 = append(i3, cur.UntUnt())
	}
	return pen.Rgb(x.I0.FltFlt(), x.I1.FltFlt(), x.I2.FltFlt(), i3...)
}
func (x PenHex) Act()             { x.PenPen() }
func (x PenHex) Ifc() interface{} { return x.PenPen() }
func (x PenHex) PenPen() pen.Pen {
	var i1 []unt.Unt
	for _, cur := range x.I1 {
		i1 = append(i1, cur.UntUnt())
	}
	return pen.Hex(x.I0.StrStr(), i1...)
}
func (x PenNewPens) Act()             { x.PenPens() }
func (x PenNewPens) Ifc() interface{} { return x.PenPens() }
func (x PenNewPens) PenPens() *pen.Pens {
	var i0 []pen.Pen
	for _, cur := range x.I0 {
		i0 = append(i0, cur.PenPen())
	}
	return pen.NewPens(i0...)
}
func (x PenMakePens) Act()                  { x.PenPens() }
func (x PenMakePens) Ifc() interface{}      { return x.PenPens() }
func (x PenMakePens) PenPens() *pen.Pens    { return pen.MakePens(x.I0.UntUnt()) }
func (x PenMakeEmpPens) Act()               { x.PenPens() }
func (x PenMakeEmpPens) Ifc() interface{}   { return x.PenPens() }
func (x PenMakeEmpPens) PenPens() *pen.Pens { return pen.MakeEmpPens(x.I0.UntUnt()) }
func (x PltNewPlts) Act()                   { x.PltPlts() }
func (x PltNewPlts) Ifc() interface{}       { return x.PltPlts() }
func (x PltNewPlts) PltPlts() *plt.Plts {
	var i0 []plt.Plt
	for _, cur := range x.I0 {
		i0 = append(i0, cur.PltPlt())
	}
	return plt.NewPlts(i0...)
}
func (x PltMakePlts) Act()                                      { x.PltPlts() }
func (x PltMakePlts) Ifc() interface{}                          { return x.PltPlts() }
func (x PltMakePlts) PltPlts() *plt.Plts                        { return plt.MakePlts(x.I0.UntUnt()) }
func (x PltMakeEmpPlts) Act()                                   { x.PltPlts() }
func (x PltMakeEmpPlts) Ifc() interface{}                       { return x.PltPlts() }
func (x PltMakeEmpPlts) PltPlts() *plt.Plts                     { return plt.MakeEmpPlts(x.I0.UntUnt()) }
func (x PltNewStm) Act()                                        { x.PltStm() }
func (x PltNewStm) Ifc() interface{}                            { return x.PltStm() }
func (x PltNewStm) PltPlt() plt.Plt                             { return x.PltStm() }
func (x PltNewStm) PltStm() *plt.Stm                            { return plt.NewStm() }
func (x PltNewFltsSctr) Act()                                   { x.PltFltsSctr() }
func (x PltNewFltsSctr) Ifc() interface{}                       { return x.PltFltsSctr() }
func (x PltNewFltsSctr) PltPlt() plt.Plt                        { return x.PltFltsSctr() }
func (x PltNewFltsSctr) PltFltsSctr() *plt.FltsSctr             { return plt.NewFltsSctr() }
func (x PltNewFltsSctrDist) Act()                               { x.PltFltsSctrDist() }
func (x PltNewFltsSctrDist) Ifc() interface{}                   { return x.PltFltsSctrDist() }
func (x PltNewFltsSctrDist) PltPlt() plt.Plt                    { return x.PltFltsSctrDist() }
func (x PltNewFltsSctrDist) PltFltsSctrDist() *plt.FltsSctrDist { return plt.NewFltsSctrDist() }
func (x PltNewHrz) Act()                                        { x.PltHrz() }
func (x PltNewHrz) Ifc() interface{}                            { return x.PltHrz() }
func (x PltNewHrz) PltPlt() plt.Plt                             { return x.PltHrz() }
func (x PltNewHrz) PltHrz() *plt.Hrz {
	var i0 []plt.Plt
	for _, cur := range x.I0 {
		i0 = append(i0, cur.PltPlt())
	}
	return plt.NewHrz(i0...)
}
func (x PltNewVrt) Act()             { x.PltVrt() }
func (x PltNewVrt) Ifc() interface{} { return x.PltVrt() }
func (x PltNewVrt) PltPlt() plt.Plt  { return x.PltVrt() }
func (x PltNewVrt) PltVrt() *plt.Vrt {
	var i0 []plt.Plt
	for _, cur := range x.I0 {
		i0 = append(i0, cur.PltPlt())
	}
	return plt.NewVrt(i0...)
}
func (x PltNewDpth) Act()             { x.PltDpth() }
func (x PltNewDpth) Ifc() interface{} { return x.PltDpth() }
func (x PltNewDpth) PltPlt() plt.Plt  { return x.PltDpth() }
func (x PltNewDpth) PltDpth() *plt.Dpth {
	var i0 []plt.Plt
	for _, cur := range x.I0 {
		i0 = append(i0, cur.PltPlt())
	}
	return plt.NewDpth(i0...)
}
func (x SysNewMu) Act()                      { x.SysMu() }
func (x SysNewMu) Ifc() interface{}          { return x.SysMu() }
func (x SysNewMu) SysMu() *sys.Mu            { return sys.NewMu() }
func (x StrStrLower) Act()                   { x.StrStr() }
func (x StrStrLower) Ifc() interface{}       { return x.StrStr() }
func (x StrStrLower) StrStr() str.Str        { return x.X.StrStr().Lower() }
func (x StrStrUpper) Act()                   { x.StrStr() }
func (x StrStrUpper) Ifc() interface{}       { return x.StrStr() }
func (x StrStrUpper) StrStr() str.Str        { return x.X.StrStr().Upper() }
func (x StrStrEql) Act()                     { x.BolBol() }
func (x StrStrEql) Ifc() interface{}         { return x.BolBol() }
func (x StrStrEql) BolBol() bol.Bol          { return x.X.StrStr().Eql(x.I0.StrStr()) }
func (x StrStrNeq) Act()                     { x.BolBol() }
func (x StrStrNeq) Ifc() interface{}         { return x.BolBol() }
func (x StrStrNeq) BolBol() bol.Bol          { return x.X.StrStr().Neq(x.I0.StrStr()) }
func (x StrStrLss) Act()                     { x.BolBol() }
func (x StrStrLss) Ifc() interface{}         { return x.BolBol() }
func (x StrStrLss) BolBol() bol.Bol          { return x.X.StrStr().Lss(x.I0.StrStr()) }
func (x StrStrGtr) Act()                     { x.BolBol() }
func (x StrStrGtr) Ifc() interface{}         { return x.BolBol() }
func (x StrStrGtr) BolBol() bol.Bol          { return x.X.StrStr().Gtr(x.I0.StrStr()) }
func (x StrStrLeq) Act()                     { x.BolBol() }
func (x StrStrLeq) Ifc() interface{}         { return x.BolBol() }
func (x StrStrLeq) BolBol() bol.Bol          { return x.X.StrStr().Leq(x.I0.StrStr()) }
func (x StrStrGeq) Act()                     { x.BolBol() }
func (x StrStrGeq) Ifc() interface{}         { return x.BolBol() }
func (x StrStrGeq) BolBol() bol.Bol          { return x.X.StrStr().Geq(x.I0.StrStr()) }
func (x BolBolNot) Act()                     { x.BolBol() }
func (x BolBolNot) Ifc() interface{}         { return x.BolBol() }
func (x BolBolNot) BolBol() bol.Bol          { return x.X.BolBol().Not() }
func (x BolBolEql) Act()                     { x.BolBol() }
func (x BolBolEql) Ifc() interface{}         { return x.BolBol() }
func (x BolBolEql) BolBol() bol.Bol          { return x.X.BolBol().Eql(x.I0.BolBol()) }
func (x BolBolNeq) Act()                     { x.BolBol() }
func (x BolBolNeq) Ifc() interface{}         { return x.BolBol() }
func (x BolBolNeq) BolBol() bol.Bol          { return x.X.BolBol().Neq(x.I0.BolBol()) }
func (x FltFltEql) Act()                     { x.BolBol() }
func (x FltFltEql) Ifc() interface{}         { return x.BolBol() }
func (x FltFltEql) BolBol() bol.Bol          { return x.X.FltFlt().Eql(x.I0.FltFlt()) }
func (x FltFltNeq) Act()                     { x.BolBol() }
func (x FltFltNeq) Ifc() interface{}         { return x.BolBol() }
func (x FltFltNeq) BolBol() bol.Bol          { return x.X.FltFlt().Neq(x.I0.FltFlt()) }
func (x FltFltTrnc) Act()                    { x.FltFlt() }
func (x FltFltTrnc) Ifc() interface{}        { return x.FltFlt() }
func (x FltFltTrnc) FltFlt() flt.Flt         { return x.X.FltFlt().Trnc(x.I0.UntUnt()) }
func (x FltFltIsNaN) Act()                   { x.BolBol() }
func (x FltFltIsNaN) Ifc() interface{}       { return x.BolBol() }
func (x FltFltIsNaN) BolBol() bol.Bol        { return x.X.FltFlt().IsNaN() }
func (x FltFltIsInfPos) Act()                { x.BolBol() }
func (x FltFltIsInfPos) Ifc() interface{}    { return x.BolBol() }
func (x FltFltIsInfPos) BolBol() bol.Bol     { return x.X.FltFlt().IsInfPos() }
func (x FltFltIsInfNeg) Act()                { x.BolBol() }
func (x FltFltIsInfNeg) Ifc() interface{}    { return x.BolBol() }
func (x FltFltIsInfNeg) BolBol() bol.Bol     { return x.X.FltFlt().IsInfNeg() }
func (x FltFltIsValid) Act()                 { x.BolBol() }
func (x FltFltIsValid) Ifc() interface{}     { return x.BolBol() }
func (x FltFltIsValid) BolBol() bol.Bol      { return x.X.FltFlt().IsValid() }
func (x FltFltPct) Act()                     { x.FltFlt() }
func (x FltFltPct) Ifc() interface{}         { return x.FltFlt() }
func (x FltFltPct) FltFlt() flt.Flt          { return x.X.FltFlt().Pct(x.I0.FltFlt()) }
func (x FltFltLss) Act()                     { x.BolBol() }
func (x FltFltLss) Ifc() interface{}         { return x.BolBol() }
func (x FltFltLss) BolBol() bol.Bol          { return x.X.FltFlt().Lss(x.I0.FltFlt()) }
func (x FltFltGtr) Act()                     { x.BolBol() }
func (x FltFltGtr) Ifc() interface{}         { return x.BolBol() }
func (x FltFltGtr) BolBol() bol.Bol          { return x.X.FltFlt().Gtr(x.I0.FltFlt()) }
func (x FltFltLeq) Act()                     { x.BolBol() }
func (x FltFltLeq) Ifc() interface{}         { return x.BolBol() }
func (x FltFltLeq) BolBol() bol.Bol          { return x.X.FltFlt().Leq(x.I0.FltFlt()) }
func (x FltFltGeq) Act()                     { x.BolBol() }
func (x FltFltGeq) Ifc() interface{}         { return x.BolBol() }
func (x FltFltGeq) BolBol() bol.Bol          { return x.X.FltFlt().Geq(x.I0.FltFlt()) }
func (x FltFltPos) Act()                     { x.FltFlt() }
func (x FltFltPos) Ifc() interface{}         { return x.FltFlt() }
func (x FltFltPos) FltFlt() flt.Flt          { return x.X.FltFlt().Pos() }
func (x FltFltNeg) Act()                     { x.FltFlt() }
func (x FltFltNeg) Ifc() interface{}         { return x.FltFlt() }
func (x FltFltNeg) FltFlt() flt.Flt          { return x.X.FltFlt().Neg() }
func (x FltFltInv) Act()                     { x.FltFlt() }
func (x FltFltInv) Ifc() interface{}         { return x.FltFlt() }
func (x FltFltInv) FltFlt() flt.Flt          { return x.X.FltFlt().Inv() }
func (x FltFltAdd) Act()                     { x.FltFlt() }
func (x FltFltAdd) Ifc() interface{}         { return x.FltFlt() }
func (x FltFltAdd) FltFlt() flt.Flt          { return x.X.FltFlt().Add(x.I0.FltFlt()) }
func (x FltFltSub) Act()                     { x.FltFlt() }
func (x FltFltSub) Ifc() interface{}         { return x.FltFlt() }
func (x FltFltSub) FltFlt() flt.Flt          { return x.X.FltFlt().Sub(x.I0.FltFlt()) }
func (x FltFltMul) Act()                     { x.FltFlt() }
func (x FltFltMul) Ifc() interface{}         { return x.FltFlt() }
func (x FltFltMul) FltFlt() flt.Flt          { return x.X.FltFlt().Mul(x.I0.FltFlt()) }
func (x FltFltDiv) Act()                     { x.FltFlt() }
func (x FltFltDiv) Ifc() interface{}         { return x.FltFlt() }
func (x FltFltDiv) FltFlt() flt.Flt          { return x.X.FltFlt().Div(x.I0.FltFlt()) }
func (x FltFltRem) Act()                     { x.FltFlt() }
func (x FltFltRem) Ifc() interface{}         { return x.FltFlt() }
func (x FltFltRem) FltFlt() flt.Flt          { return x.X.FltFlt().Rem(x.I0.FltFlt()) }
func (x FltFltPow) Act()                     { x.FltFlt() }
func (x FltFltPow) Ifc() interface{}         { return x.FltFlt() }
func (x FltFltPow) FltFlt() flt.Flt          { return x.X.FltFlt().Pow(x.I0.FltFlt()) }
func (x FltFltSqr) Act()                     { x.FltFlt() }
func (x FltFltSqr) Ifc() interface{}         { return x.FltFlt() }
func (x FltFltSqr) FltFlt() flt.Flt          { return x.X.FltFlt().Sqr() }
func (x FltFltSqrt) Act()                    { x.FltFlt() }
func (x FltFltSqrt) Ifc() interface{}        { return x.FltFlt() }
func (x FltFltSqrt) FltFlt() flt.Flt         { return x.X.FltFlt().Sqrt() }
func (x FltFltMin) Act()                     { x.FltFlt() }
func (x FltFltMin) Ifc() interface{}         { return x.FltFlt() }
func (x FltFltMin) FltFlt() flt.Flt          { return x.X.FltFlt().Min(x.I0.FltFlt()) }
func (x FltFltMax) Act()                     { x.FltFlt() }
func (x FltFltMax) Ifc() interface{}         { return x.FltFlt() }
func (x FltFltMax) FltFlt() flt.Flt          { return x.X.FltFlt().Max(x.I0.FltFlt()) }
func (x FltFltMid) Act()                     { x.FltFlt() }
func (x FltFltMid) Ifc() interface{}         { return x.FltFlt() }
func (x FltFltMid) FltFlt() flt.Flt          { return x.X.FltFlt().Mid(x.I0.FltFlt()) }
func (x FltFltAvg) Act()                     { x.FltFlt() }
func (x FltFltAvg) Ifc() interface{}         { return x.FltFlt() }
func (x FltFltAvg) FltFlt() flt.Flt          { return x.X.FltFlt().Avg(x.I0.FltFlt()) }
func (x FltFltAvgGeo) Act()                  { x.FltFlt() }
func (x FltFltAvgGeo) Ifc() interface{}      { return x.FltFlt() }
func (x FltFltAvgGeo) FltFlt() flt.Flt       { return x.X.FltFlt().AvgGeo(x.I0.FltFlt()) }
func (x FltFltSelEql) Act()                  { x.FltFlt() }
func (x FltFltSelEql) Ifc() interface{}      { return x.FltFlt() }
func (x FltFltSelEql) FltFlt() flt.Flt       { return x.X.FltFlt().SelEql(x.I0.FltFlt()) }
func (x FltFltSelNeq) Act()                  { x.FltFlt() }
func (x FltFltSelNeq) Ifc() interface{}      { return x.FltFlt() }
func (x FltFltSelNeq) FltFlt() flt.Flt       { return x.X.FltFlt().SelNeq(x.I0.FltFlt()) }
func (x FltFltSelLss) Act()                  { x.FltFlt() }
func (x FltFltSelLss) Ifc() interface{}      { return x.FltFlt() }
func (x FltFltSelLss) FltFlt() flt.Flt       { return x.X.FltFlt().SelLss(x.I0.FltFlt()) }
func (x FltFltSelGtr) Act()                  { x.FltFlt() }
func (x FltFltSelGtr) Ifc() interface{}      { return x.FltFlt() }
func (x FltFltSelGtr) FltFlt() flt.Flt       { return x.X.FltFlt().SelGtr(x.I0.FltFlt()) }
func (x FltFltSelLeq) Act()                  { x.FltFlt() }
func (x FltFltSelLeq) Ifc() interface{}      { return x.FltFlt() }
func (x FltFltSelLeq) FltFlt() flt.Flt       { return x.X.FltFlt().SelLeq(x.I0.FltFlt()) }
func (x FltFltSelGeq) Act()                  { x.FltFlt() }
func (x FltFltSelGeq) Ifc() interface{}      { return x.FltFlt() }
func (x FltFltSelGeq) FltFlt() flt.Flt       { return x.X.FltFlt().SelGeq(x.I0.FltFlt()) }
func (x UntUntEql) Act()                     { x.BolBol() }
func (x UntUntEql) Ifc() interface{}         { return x.BolBol() }
func (x UntUntEql) BolBol() bol.Bol          { return x.X.UntUnt().Eql(x.I0.UntUnt()) }
func (x UntUntNeq) Act()                     { x.BolBol() }
func (x UntUntNeq) Ifc() interface{}         { return x.BolBol() }
func (x UntUntNeq) BolBol() bol.Bol          { return x.X.UntUnt().Neq(x.I0.UntUnt()) }
func (x UntUntLss) Act()                     { x.BolBol() }
func (x UntUntLss) Ifc() interface{}         { return x.BolBol() }
func (x UntUntLss) BolBol() bol.Bol          { return x.X.UntUnt().Lss(x.I0.UntUnt()) }
func (x UntUntGtr) Act()                     { x.BolBol() }
func (x UntUntGtr) Ifc() interface{}         { return x.BolBol() }
func (x UntUntGtr) BolBol() bol.Bol          { return x.X.UntUnt().Gtr(x.I0.UntUnt()) }
func (x UntUntLeq) Act()                     { x.BolBol() }
func (x UntUntLeq) Ifc() interface{}         { return x.BolBol() }
func (x UntUntLeq) BolBol() bol.Bol          { return x.X.UntUnt().Leq(x.I0.UntUnt()) }
func (x UntUntGeq) Act()                     { x.BolBol() }
func (x UntUntGeq) Ifc() interface{}         { return x.BolBol() }
func (x UntUntGeq) BolBol() bol.Bol          { return x.X.UntUnt().Geq(x.I0.UntUnt()) }
func (x UntUntAdd) Act()                     { x.UntUnt() }
func (x UntUntAdd) Ifc() interface{}         { return x.UntUnt() }
func (x UntUntAdd) UntUnt() unt.Unt          { return x.X.UntUnt().Add(x.I0.UntUnt()) }
func (x UntUntSub) Act()                     { x.UntUnt() }
func (x UntUntSub) Ifc() interface{}         { return x.UntUnt() }
func (x UntUntSub) UntUnt() unt.Unt          { return x.X.UntUnt().Sub(x.I0.UntUnt()) }
func (x UntUntMul) Act()                     { x.UntUnt() }
func (x UntUntMul) Ifc() interface{}         { return x.UntUnt() }
func (x UntUntMul) UntUnt() unt.Unt          { return x.X.UntUnt().Mul(x.I0.UntUnt()) }
func (x UntUntDiv) Act()                     { x.UntUnt() }
func (x UntUntDiv) Ifc() interface{}         { return x.UntUnt() }
func (x UntUntDiv) UntUnt() unt.Unt          { return x.X.UntUnt().Div(x.I0.UntUnt()) }
func (x UntUntRem) Act()                     { x.UntUnt() }
func (x UntUntRem) Ifc() interface{}         { return x.UntUnt() }
func (x UntUntRem) UntUnt() unt.Unt          { return x.X.UntUnt().Rem(x.I0.UntUnt()) }
func (x UntUntPow) Act()                     { x.UntUnt() }
func (x UntUntPow) Ifc() interface{}         { return x.UntUnt() }
func (x UntUntPow) UntUnt() unt.Unt          { return x.X.UntUnt().Pow(x.I0.UntUnt()) }
func (x UntUntSqr) Act()                     { x.UntUnt() }
func (x UntUntSqr) Ifc() interface{}         { return x.UntUnt() }
func (x UntUntSqr) UntUnt() unt.Unt          { return x.X.UntUnt().Sqr() }
func (x UntUntSqrt) Act()                    { x.UntUnt() }
func (x UntUntSqrt) Ifc() interface{}        { return x.UntUnt() }
func (x UntUntSqrt) UntUnt() unt.Unt         { return x.X.UntUnt().Sqrt() }
func (x UntUntMin) Act()                     { x.UntUnt() }
func (x UntUntMin) Ifc() interface{}         { return x.UntUnt() }
func (x UntUntMin) UntUnt() unt.Unt          { return x.X.UntUnt().Min(x.I0.UntUnt()) }
func (x UntUntMax) Act()                     { x.UntUnt() }
func (x UntUntMax) Ifc() interface{}         { return x.UntUnt() }
func (x UntUntMax) UntUnt() unt.Unt          { return x.X.UntUnt().Max(x.I0.UntUnt()) }
func (x UntUntMid) Act()                     { x.UntUnt() }
func (x UntUntMid) Ifc() interface{}         { return x.UntUnt() }
func (x UntUntMid) UntUnt() unt.Unt          { return x.X.UntUnt().Mid(x.I0.UntUnt()) }
func (x UntUntAvg) Act()                     { x.UntUnt() }
func (x UntUntAvg) Ifc() interface{}         { return x.UntUnt() }
func (x UntUntAvg) UntUnt() unt.Unt          { return x.X.UntUnt().Avg(x.I0.UntUnt()) }
func (x UntUntAvgGeo) Act()                  { x.UntUnt() }
func (x UntUntAvgGeo) Ifc() interface{}      { return x.UntUnt() }
func (x UntUntAvgGeo) UntUnt() unt.Unt       { return x.X.UntUnt().AvgGeo(x.I0.UntUnt()) }
func (x IntIntEql) Act()                     { x.BolBol() }
func (x IntIntEql) Ifc() interface{}         { return x.BolBol() }
func (x IntIntEql) BolBol() bol.Bol          { return x.X.IntInt().Eql(x.I0.IntInt()) }
func (x IntIntNeq) Act()                     { x.BolBol() }
func (x IntIntNeq) Ifc() interface{}         { return x.BolBol() }
func (x IntIntNeq) BolBol() bol.Bol          { return x.X.IntInt().Neq(x.I0.IntInt()) }
func (x IntIntLss) Act()                     { x.BolBol() }
func (x IntIntLss) Ifc() interface{}         { return x.BolBol() }
func (x IntIntLss) BolBol() bol.Bol          { return x.X.IntInt().Lss(x.I0.IntInt()) }
func (x IntIntGtr) Act()                     { x.BolBol() }
func (x IntIntGtr) Ifc() interface{}         { return x.BolBol() }
func (x IntIntGtr) BolBol() bol.Bol          { return x.X.IntInt().Gtr(x.I0.IntInt()) }
func (x IntIntLeq) Act()                     { x.BolBol() }
func (x IntIntLeq) Ifc() interface{}         { return x.BolBol() }
func (x IntIntLeq) BolBol() bol.Bol          { return x.X.IntInt().Leq(x.I0.IntInt()) }
func (x IntIntGeq) Act()                     { x.BolBol() }
func (x IntIntGeq) Ifc() interface{}         { return x.BolBol() }
func (x IntIntGeq) BolBol() bol.Bol          { return x.X.IntInt().Geq(x.I0.IntInt()) }
func (x IntIntPos) Act()                     { x.IntInt() }
func (x IntIntPos) Ifc() interface{}         { return x.IntInt() }
func (x IntIntPos) IntInt() int.Int          { return x.X.IntInt().Pos() }
func (x IntIntNeg) Act()                     { x.IntInt() }
func (x IntIntNeg) Ifc() interface{}         { return x.IntInt() }
func (x IntIntNeg) IntInt() int.Int          { return x.X.IntInt().Neg() }
func (x IntIntInv) Act()                     { x.IntInt() }
func (x IntIntInv) Ifc() interface{}         { return x.IntInt() }
func (x IntIntInv) IntInt() int.Int          { return x.X.IntInt().Inv() }
func (x IntIntAdd) Act()                     { x.IntInt() }
func (x IntIntAdd) Ifc() interface{}         { return x.IntInt() }
func (x IntIntAdd) IntInt() int.Int          { return x.X.IntInt().Add(x.I0.IntInt()) }
func (x IntIntSub) Act()                     { x.IntInt() }
func (x IntIntSub) Ifc() interface{}         { return x.IntInt() }
func (x IntIntSub) IntInt() int.Int          { return x.X.IntInt().Sub(x.I0.IntInt()) }
func (x IntIntMul) Act()                     { x.IntInt() }
func (x IntIntMul) Ifc() interface{}         { return x.IntInt() }
func (x IntIntMul) IntInt() int.Int          { return x.X.IntInt().Mul(x.I0.IntInt()) }
func (x IntIntDiv) Act()                     { x.IntInt() }
func (x IntIntDiv) Ifc() interface{}         { return x.IntInt() }
func (x IntIntDiv) IntInt() int.Int          { return x.X.IntInt().Div(x.I0.IntInt()) }
func (x IntIntRem) Act()                     { x.IntInt() }
func (x IntIntRem) Ifc() interface{}         { return x.IntInt() }
func (x IntIntRem) IntInt() int.Int          { return x.X.IntInt().Rem(x.I0.IntInt()) }
func (x IntIntPow) Act()                     { x.IntInt() }
func (x IntIntPow) Ifc() interface{}         { return x.IntInt() }
func (x IntIntPow) IntInt() int.Int          { return x.X.IntInt().Pow(x.I0.IntInt()) }
func (x IntIntSqr) Act()                     { x.IntInt() }
func (x IntIntSqr) Ifc() interface{}         { return x.IntInt() }
func (x IntIntSqr) IntInt() int.Int          { return x.X.IntInt().Sqr() }
func (x IntIntSqrt) Act()                    { x.IntInt() }
func (x IntIntSqrt) Ifc() interface{}        { return x.IntInt() }
func (x IntIntSqrt) IntInt() int.Int         { return x.X.IntInt().Sqrt() }
func (x IntIntMin) Act()                     { x.IntInt() }
func (x IntIntMin) Ifc() interface{}         { return x.IntInt() }
func (x IntIntMin) IntInt() int.Int          { return x.X.IntInt().Min(x.I0.IntInt()) }
func (x IntIntMax) Act()                     { x.IntInt() }
func (x IntIntMax) Ifc() interface{}         { return x.IntInt() }
func (x IntIntMax) IntInt() int.Int          { return x.X.IntInt().Max(x.I0.IntInt()) }
func (x IntIntMid) Act()                     { x.IntInt() }
func (x IntIntMid) Ifc() interface{}         { return x.IntInt() }
func (x IntIntMid) IntInt() int.Int          { return x.X.IntInt().Mid(x.I0.IntInt()) }
func (x IntIntAvg) Act()                     { x.IntInt() }
func (x IntIntAvg) Ifc() interface{}         { return x.IntInt() }
func (x IntIntAvg) IntInt() int.Int          { return x.X.IntInt().Avg(x.I0.IntInt()) }
func (x IntIntAvgGeo) Act()                  { x.IntInt() }
func (x IntIntAvgGeo) Ifc() interface{}      { return x.IntInt() }
func (x IntIntAvgGeo) IntInt() int.Int       { return x.X.IntInt().AvgGeo(x.I0.IntInt()) }
func (x TmeTmeWeekdayCnt) Act()              { x.UntUnt() }
func (x TmeTmeWeekdayCnt) Ifc() interface{}  { return x.UntUnt() }
func (x TmeTmeWeekdayCnt) UntUnt() unt.Unt   { return x.X.TmeTme().WeekdayCnt(x.I0.TmeTme()) }
func (x TmeTmeDte) Act()                     { x.TmeTme() }
func (x TmeTmeDte) Ifc() interface{}         { return x.TmeTme() }
func (x TmeTmeDte) TmeTme() tme.Tme          { return x.X.TmeTme().Dte() }
func (x TmeTmeToSunday) Act()                { x.TmeTme() }
func (x TmeTmeToSunday) Ifc() interface{}    { return x.TmeTme() }
func (x TmeTmeToSunday) TmeTme() tme.Tme     { return x.X.TmeTme().ToSunday() }
func (x TmeTmeToMonday) Act()                { x.TmeTme() }
func (x TmeTmeToMonday) Ifc() interface{}    { return x.TmeTme() }
func (x TmeTmeToMonday) TmeTme() tme.Tme     { return x.X.TmeTme().ToMonday() }
func (x TmeTmeToTuesday) Act()               { x.TmeTme() }
func (x TmeTmeToTuesday) Ifc() interface{}   { return x.TmeTme() }
func (x TmeTmeToTuesday) TmeTme() tme.Tme    { return x.X.TmeTme().ToTuesday() }
func (x TmeTmeToWednesday) Act()             { x.TmeTme() }
func (x TmeTmeToWednesday) Ifc() interface{} { return x.TmeTme() }
func (x TmeTmeToWednesday) TmeTme() tme.Tme  { return x.X.TmeTme().ToWednesday() }
func (x TmeTmeToThursday) Act()              { x.TmeTme() }
func (x TmeTmeToThursday) Ifc() interface{}  { return x.TmeTme() }
func (x TmeTmeToThursday) TmeTme() tme.Tme   { return x.X.TmeTme().ToThursday() }
func (x TmeTmeToFriday) Act()                { x.TmeTme() }
func (x TmeTmeToFriday) Ifc() interface{}    { return x.TmeTme() }
func (x TmeTmeToFriday) TmeTme() tme.Tme     { return x.X.TmeTme().ToFriday() }
func (x TmeTmeToSaturday) Act()              { x.TmeTme() }
func (x TmeTmeToSaturday) Ifc() interface{}  { return x.TmeTme() }
func (x TmeTmeToSaturday) TmeTme() tme.Tme   { return x.X.TmeTme().ToSaturday() }
func (x TmeTmeIsSunday) Act()                { x.BolBol() }
func (x TmeTmeIsSunday) Ifc() interface{}    { return x.BolBol() }
func (x TmeTmeIsSunday) BolBol() bol.Bol     { return x.X.TmeTme().IsSunday() }
func (x TmeTmeIsMonday) Act()                { x.BolBol() }
func (x TmeTmeIsMonday) Ifc() interface{}    { return x.BolBol() }
func (x TmeTmeIsMonday) BolBol() bol.Bol     { return x.X.TmeTme().IsMonday() }
func (x TmeTmeIsTuesday) Act()               { x.BolBol() }
func (x TmeTmeIsTuesday) Ifc() interface{}   { return x.BolBol() }
func (x TmeTmeIsTuesday) BolBol() bol.Bol    { return x.X.TmeTme().IsTuesday() }
func (x TmeTmeIsWednesday) Act()             { x.BolBol() }
func (x TmeTmeIsWednesday) Ifc() interface{} { return x.BolBol() }
func (x TmeTmeIsWednesday) BolBol() bol.Bol  { return x.X.TmeTme().IsWednesday() }
func (x TmeTmeIsThursday) Act()              { x.BolBol() }
func (x TmeTmeIsThursday) Ifc() interface{}  { return x.BolBol() }
func (x TmeTmeIsThursday) BolBol() bol.Bol   { return x.X.TmeTme().IsThursday() }
func (x TmeTmeIsFriday) Act()                { x.BolBol() }
func (x TmeTmeIsFriday) Ifc() interface{}    { return x.BolBol() }
func (x TmeTmeIsFriday) BolBol() bol.Bol     { return x.X.TmeTme().IsFriday() }
func (x TmeTmeIsSaturday) Act()              { x.BolBol() }
func (x TmeTmeIsSaturday) Ifc() interface{}  { return x.BolBol() }
func (x TmeTmeIsSaturday) BolBol() bol.Bol   { return x.X.TmeTme().IsSaturday() }
func (x TmeTmeEql) Act()                     { x.BolBol() }
func (x TmeTmeEql) Ifc() interface{}         { return x.BolBol() }
func (x TmeTmeEql) BolBol() bol.Bol          { return x.X.TmeTme().Eql(x.I0.TmeTme()) }
func (x TmeTmeNeq) Act()                     { x.BolBol() }
func (x TmeTmeNeq) Ifc() interface{}         { return x.BolBol() }
func (x TmeTmeNeq) BolBol() bol.Bol          { return x.X.TmeTme().Neq(x.I0.TmeTme()) }
func (x TmeTmeLss) Act()                     { x.BolBol() }
func (x TmeTmeLss) Ifc() interface{}         { return x.BolBol() }
func (x TmeTmeLss) BolBol() bol.Bol          { return x.X.TmeTme().Lss(x.I0.TmeTme()) }
func (x TmeTmeGtr) Act()                     { x.BolBol() }
func (x TmeTmeGtr) Ifc() interface{}         { return x.BolBol() }
func (x TmeTmeGtr) BolBol() bol.Bol          { return x.X.TmeTme().Gtr(x.I0.TmeTme()) }
func (x TmeTmeLeq) Act()                     { x.BolBol() }
func (x TmeTmeLeq) Ifc() interface{}         { return x.BolBol() }
func (x TmeTmeLeq) BolBol() bol.Bol          { return x.X.TmeTme().Leq(x.I0.TmeTme()) }
func (x TmeTmeGeq) Act()                     { x.BolBol() }
func (x TmeTmeGeq) Ifc() interface{}         { return x.BolBol() }
func (x TmeTmeGeq) BolBol() bol.Bol          { return x.X.TmeTme().Geq(x.I0.TmeTme()) }
func (x TmeTmePos) Act()                     { x.TmeTme() }
func (x TmeTmePos) Ifc() interface{}         { return x.TmeTme() }
func (x TmeTmePos) TmeTme() tme.Tme          { return x.X.TmeTme().Pos() }
func (x TmeTmeNeg) Act()                     { x.TmeTme() }
func (x TmeTmeNeg) Ifc() interface{}         { return x.TmeTme() }
func (x TmeTmeNeg) TmeTme() tme.Tme          { return x.X.TmeTme().Neg() }
func (x TmeTmeInv) Act()                     { x.TmeTme() }
func (x TmeTmeInv) Ifc() interface{}         { return x.TmeTme() }
func (x TmeTmeInv) TmeTme() tme.Tme          { return x.X.TmeTme().Inv() }
func (x TmeTmeAdd) Act()                     { x.TmeTme() }
func (x TmeTmeAdd) Ifc() interface{}         { return x.TmeTme() }
func (x TmeTmeAdd) TmeTme() tme.Tme          { return x.X.TmeTme().Add(x.I0.TmeTme()) }
func (x TmeTmeSub) Act()                     { x.TmeTme() }
func (x TmeTmeSub) Ifc() interface{}         { return x.TmeTme() }
func (x TmeTmeSub) TmeTme() tme.Tme          { return x.X.TmeTme().Sub(x.I0.TmeTme()) }
func (x TmeTmeMul) Act()                     { x.TmeTme() }
func (x TmeTmeMul) Ifc() interface{}         { return x.TmeTme() }
func (x TmeTmeMul) TmeTme() tme.Tme          { return x.X.TmeTme().Mul(x.I0.TmeTme()) }
func (x TmeTmeDiv) Act()                     { x.TmeTme() }
func (x TmeTmeDiv) Ifc() interface{}         { return x.TmeTme() }
func (x TmeTmeDiv) TmeTme() tme.Tme          { return x.X.TmeTme().Div(x.I0.TmeTme()) }
func (x TmeTmeRem) Act()                     { x.TmeTme() }
func (x TmeTmeRem) Ifc() interface{}         { return x.TmeTme() }
func (x TmeTmeRem) TmeTme() tme.Tme          { return x.X.TmeTme().Rem(x.I0.TmeTme()) }
func (x TmeTmePow) Act()                     { x.TmeTme() }
func (x TmeTmePow) Ifc() interface{}         { return x.TmeTme() }
func (x TmeTmePow) TmeTme() tme.Tme          { return x.X.TmeTme().Pow(x.I0.TmeTme()) }
func (x TmeTmeSqr) Act()                     { x.TmeTme() }
func (x TmeTmeSqr) Ifc() interface{}         { return x.TmeTme() }
func (x TmeTmeSqr) TmeTme() tme.Tme          { return x.X.TmeTme().Sqr() }
func (x TmeTmeSqrt) Act()                    { x.TmeTme() }
func (x TmeTmeSqrt) Ifc() interface{}        { return x.TmeTme() }
func (x TmeTmeSqrt) TmeTme() tme.Tme         { return x.X.TmeTme().Sqrt() }
func (x TmeTmeMin) Act()                     { x.TmeTme() }
func (x TmeTmeMin) Ifc() interface{}         { return x.TmeTme() }
func (x TmeTmeMin) TmeTme() tme.Tme          { return x.X.TmeTme().Min(x.I0.TmeTme()) }
func (x TmeTmeMax) Act()                     { x.TmeTme() }
func (x TmeTmeMax) Ifc() interface{}         { return x.TmeTme() }
func (x TmeTmeMax) TmeTme() tme.Tme          { return x.X.TmeTme().Max(x.I0.TmeTme()) }
func (x TmeTmeMid) Act()                     { x.TmeTme() }
func (x TmeTmeMid) Ifc() interface{}         { return x.TmeTme() }
func (x TmeTmeMid) TmeTme() tme.Tme          { return x.X.TmeTme().Mid(x.I0.TmeTme()) }
func (x TmeTmeAvg) Act()                     { x.TmeTme() }
func (x TmeTmeAvg) Ifc() interface{}         { return x.TmeTme() }
func (x TmeTmeAvg) TmeTme() tme.Tme          { return x.X.TmeTme().Avg(x.I0.TmeTme()) }
func (x TmeTmeAvgGeo) Act()                  { x.TmeTme() }
func (x TmeTmeAvgGeo) Ifc() interface{}      { return x.TmeTme() }
func (x TmeTmeAvgGeo) TmeTme() tme.Tme       { return x.X.TmeTme().AvgGeo(x.I0.TmeTme()) }
func (x BndBndCnt) Act()                     { x.UntUnt() }
func (x BndBndCnt) Ifc() interface{}         { return x.UntUnt() }
func (x BndBndCnt) UntUnt() unt.Unt          { return x.X.BndBnd().Cnt() }
func (x BndBndLen) Act()                     { x.UntUnt() }
func (x BndBndLen) Ifc() interface{}         { return x.UntUnt() }
func (x BndBndLen) UntUnt() unt.Unt          { return x.X.BndBnd().Len() }
func (x BndBndLstIdx) Act()                  { x.UntUnt() }
func (x BndBndLstIdx) Ifc() interface{}      { return x.UntUnt() }
func (x BndBndLstIdx) UntUnt() unt.Unt       { return x.X.BndBnd().LstIdx() }
func (x BndBndIsValid) Act()                 { x.BolBol() }
func (x BndBndIsValid) Ifc() interface{}     { return x.BolBol() }
func (x BndBndIsValid) BolBol() bol.Bol      { return x.X.BndBnd().IsValid() }
func (x FltRngLen) Act()                     { x.FltFlt() }
func (x FltRngLen) Ifc() interface{}         { return x.FltFlt() }
func (x FltRngLen) FltFlt() flt.Flt          { return x.X.FltRng().Len() }
func (x FltRngIsValid) Act()                 { x.BolBol() }
func (x FltRngIsValid) Ifc() interface{}     { return x.BolBol() }
func (x FltRngIsValid) BolBol() bol.Bol      { return x.X.FltRng().IsValid() }
func (x FltRngEnsure) Act()                  { x.FltRng() }
func (x FltRngEnsure) Ifc() interface{}      { return x.FltRng() }
func (x FltRngEnsure) FltRng() flt.Rng       { return x.X.FltRng().Ensure() }
func (x FltRngMinSub) Act()                  { x.FltRng() }
func (x FltRngMinSub) Ifc() interface{}      { return x.FltRng() }
func (x FltRngMinSub) FltRng() flt.Rng       { return x.X.FltRng().MinSub(x.I0.FltFlt()) }
func (x FltRngMaxAdd) Act()                  { x.FltRng() }
func (x FltRngMaxAdd) Ifc() interface{}      { return x.FltRng() }
func (x FltRngMaxAdd) FltRng() flt.Rng       { return x.X.FltRng().MaxAdd(x.I0.FltFlt()) }
func (x FltRngMrg) Act()                     { x.FltRng() }
func (x FltRngMrg) Ifc() interface{}         { return x.FltRng() }
func (x FltRngMrg) FltRng() flt.Rng          { return x.X.FltRng().Mrg(x.I0.FltRng()) }
func (x TmeRngLen) Act()                     { x.TmeTme() }
func (x TmeRngLen) Ifc() interface{}         { return x.TmeTme() }
func (x TmeRngLen) TmeTme() tme.Tme          { return x.X.TmeRng().Len() }
func (x TmeRngIsValid) Act()                 { x.BolBol() }
func (x TmeRngIsValid) Ifc() interface{}     { return x.BolBol() }
func (x TmeRngIsValid) BolBol() bol.Bol      { return x.X.TmeRng().IsValid() }
func (x TmeRngEnsure) Act()                  { x.TmeRng() }
func (x TmeRngEnsure) Ifc() interface{}      { return x.TmeRng() }
func (x TmeRngEnsure) TmeRng() tme.Rng       { return x.X.TmeRng().Ensure() }
func (x TmeRngMinSub) Act()                  { x.TmeRng() }
func (x TmeRngMinSub) Ifc() interface{}      { return x.TmeRng() }
func (x TmeRngMinSub) TmeRng() tme.Rng       { return x.X.TmeRng().MinSub(x.I0.TmeTme()) }
func (x TmeRngMaxAdd) Act()                  { x.TmeRng() }
func (x TmeRngMaxAdd) Ifc() interface{}      { return x.TmeRng() }
func (x TmeRngMaxAdd) TmeRng() tme.Rng       { return x.X.TmeRng().MaxAdd(x.I0.TmeTme()) }
func (x TmeRngMrg) Act()                     { x.TmeRng() }
func (x TmeRngMrg) Ifc() interface{}         { return x.TmeRng() }
func (x TmeRngMrg) TmeRng() tme.Rng          { return x.X.TmeRng().Mrg(x.I0.TmeRng()) }
func (x StrsStrsCnt) Act()                   { x.UntUnt() }
func (x StrsStrsCnt) Ifc() interface{}       { return x.UntUnt() }
func (x StrsStrsCnt) UntUnt() unt.Unt        { return x.X.StrsStrs().Cnt() }
func (x StrsStrsCpy) Act()                   { x.StrsStrs() }
func (x StrsStrsCpy) Ifc() interface{}       { return x.StrsStrs() }
func (x StrsStrsCpy) StrsStrs() *strs.Strs   { return x.X.StrsStrs().Cpy() }
func (x StrsStrsClr) Act()                   { x.StrsStrs() }
func (x StrsStrsClr) Ifc() interface{}       { return x.StrsStrs() }
func (x StrsStrsClr) StrsStrs() *strs.Strs   { return x.X.StrsStrs().Clr() }
func (x StrsStrsRand) Act()                  { x.StrsStrs() }
func (x StrsStrsRand) Ifc() interface{}      { return x.StrsStrs() }
func (x StrsStrsRand) StrsStrs() *strs.Strs  { return x.X.StrsStrs().Rand() }
func (x StrsStrsMrg) Act()                   { x.StrsStrs() }
func (x StrsStrsMrg) Ifc() interface{}       { return x.StrsStrs() }
func (x StrsStrsMrg) StrsStrs() *strs.Strs {
	var i0 []*strs.Strs
	for _, cur := range x.I0 {
		i0 = append(i0, cur.StrsStrs())
	}
	return x.X.StrsStrs().Mrg(i0...)
}
func (x StrsStrsPush) Act()             { x.StrsStrs() }
func (x StrsStrsPush) Ifc() interface{} { return x.StrsStrs() }
func (x StrsStrsPush) StrsStrs() *strs.Strs {
	var i0 []str.Str
	for _, cur := range x.I0 {
		i0 = append(i0, cur.StrStr())
	}
	return x.X.StrsStrs().Push(i0...)
}
func (x StrsStrsPop) Act()             { x.StrStr() }
func (x StrsStrsPop) Ifc() interface{} { return x.StrStr() }
func (x StrsStrsPop) StrStr() str.Str  { return x.X.StrsStrs().Pop() }
func (x StrsStrsQue) Act()             { x.StrsStrs() }
func (x StrsStrsQue) Ifc() interface{} { return x.StrsStrs() }
func (x StrsStrsQue) StrsStrs() *strs.Strs {
	var i0 []str.Str
	for _, cur := range x.I0 {
		i0 = append(i0, cur.StrStr())
	}
	return x.X.StrsStrs().Que(i0...)
}
func (x StrsStrsDque) Act()                   { x.StrStr() }
func (x StrsStrsDque) Ifc() interface{}       { return x.StrStr() }
func (x StrsStrsDque) StrStr() str.Str        { return x.X.StrsStrs().Dque() }
func (x StrsStrsIns) Act()                    { x.StrsStrs() }
func (x StrsStrsIns) Ifc() interface{}        { return x.StrsStrs() }
func (x StrsStrsIns) StrsStrs() *strs.Strs    { return x.X.StrsStrs().Ins(x.I0.UntUnt(), x.I1.StrStr()) }
func (x StrsStrsUpd) Act()                    { x.StrsStrs() }
func (x StrsStrsUpd) Ifc() interface{}        { return x.StrsStrs() }
func (x StrsStrsUpd) StrsStrs() *strs.Strs    { return x.X.StrsStrs().Upd(x.I0.UntUnt(), x.I1.StrStr()) }
func (x StrsStrsDel) Act()                    { x.StrStr() }
func (x StrsStrsDel) Ifc() interface{}        { return x.StrStr() }
func (x StrsStrsDel) StrStr() str.Str         { return x.X.StrsStrs().Del(x.I0.UntUnt()) }
func (x StrsStrsAt) Act()                     { x.StrStr() }
func (x StrsStrsAt) Ifc() interface{}         { return x.StrStr() }
func (x StrsStrsAt) StrStr() str.Str          { return x.X.StrsStrs().At(x.I0.UntUnt()) }
func (x StrsStrsIn) Act()                     { x.StrsStrs() }
func (x StrsStrsIn) Ifc() interface{}         { return x.StrsStrs() }
func (x StrsStrsIn) StrsStrs() *strs.Strs     { return x.X.StrsStrs().In(x.I0.UntUnt(), x.I1.UntUnt()) }
func (x StrsStrsInBnd) Act()                  { x.StrsStrs() }
func (x StrsStrsInBnd) Ifc() interface{}      { return x.StrsStrs() }
func (x StrsStrsInBnd) StrsStrs() *strs.Strs  { return x.X.StrsStrs().InBnd(x.I0.BndBnd()) }
func (x StrsStrsFrom) Act()                   { x.StrsStrs() }
func (x StrsStrsFrom) Ifc() interface{}       { return x.StrsStrs() }
func (x StrsStrsFrom) StrsStrs() *strs.Strs   { return x.X.StrsStrs().From(x.I0.UntUnt()) }
func (x StrsStrsTo) Act()                     { x.StrsStrs() }
func (x StrsStrsTo) Ifc() interface{}         { return x.StrsStrs() }
func (x StrsStrsTo) StrsStrs() *strs.Strs     { return x.X.StrsStrs().To(x.I0.UntUnt()) }
func (x StrsStrsFst) Act()                    { x.StrStr() }
func (x StrsStrsFst) Ifc() interface{}        { return x.StrStr() }
func (x StrsStrsFst) StrStr() str.Str         { return x.X.StrsStrs().Fst() }
func (x StrsStrsMdl) Act()                    { x.StrStr() }
func (x StrsStrsMdl) Ifc() interface{}        { return x.StrStr() }
func (x StrsStrsMdl) StrStr() str.Str         { return x.X.StrsStrs().Mdl() }
func (x StrsStrsLst) Act()                    { x.StrStr() }
func (x StrsStrsLst) Ifc() interface{}        { return x.StrStr() }
func (x StrsStrsLst) StrStr() str.Str         { return x.X.StrsStrs().Lst() }
func (x StrsStrsFstIdx) Act()                 { x.UntUnt() }
func (x StrsStrsFstIdx) Ifc() interface{}     { return x.UntUnt() }
func (x StrsStrsFstIdx) UntUnt() unt.Unt      { return x.X.StrsStrs().FstIdx() }
func (x StrsStrsMdlIdx) Act()                 { x.UntUnt() }
func (x StrsStrsMdlIdx) Ifc() interface{}     { return x.UntUnt() }
func (x StrsStrsMdlIdx) UntUnt() unt.Unt      { return x.X.StrsStrs().MdlIdx() }
func (x StrsStrsLstIdx) Act()                 { x.UntUnt() }
func (x StrsStrsLstIdx) Ifc() interface{}     { return x.UntUnt() }
func (x StrsStrsLstIdx) UntUnt() unt.Unt      { return x.X.StrsStrs().LstIdx() }
func (x StrsStrsRev) Act()                    { x.StrsStrs() }
func (x StrsStrsRev) Ifc() interface{}        { return x.StrsStrs() }
func (x StrsStrsRev) StrsStrs() *strs.Strs    { return x.X.StrsStrs().Rev() }
func (x StrsStrsSrchIdxEql) Act()             { x.UntUnt() }
func (x StrsStrsSrchIdxEql) Ifc() interface{} { return x.UntUnt() }
func (x StrsStrsSrchIdxEql) UntUnt() unt.Unt  { return x.X.StrsStrs().SrchIdxEql(x.I0.StrStr()) }
func (x StrsStrsSrchIdx) Act()                { x.UntUnt() }
func (x StrsStrsSrchIdx) Ifc() interface{}    { return x.UntUnt() }
func (x StrsStrsSrchIdx) UntUnt() unt.Unt {
	var i1 []bol.Bol
	for _, cur := range x.I1 {
		i1 = append(i1, cur.BolBol())
	}
	return x.X.StrsStrs().SrchIdx(x.I0.StrStr(), i1...)
}
func (x StrsStrsHas) Act()                    { x.BolBol() }
func (x StrsStrsHas) Ifc() interface{}        { return x.BolBol() }
func (x StrsStrsHas) BolBol() bol.Bol         { return x.X.StrsStrs().Has(x.I0.StrStr()) }
func (x StrsStrsSrtAsc) Act()                 { x.StrsStrs() }
func (x StrsStrsSrtAsc) Ifc() interface{}     { return x.StrsStrs() }
func (x StrsStrsSrtAsc) StrsStrs() *strs.Strs { return x.X.StrsStrs().SrtAsc() }
func (x StrsStrsSrtDsc) Act()                 { x.StrsStrs() }
func (x StrsStrsSrtDsc) Ifc() interface{}     { return x.StrsStrs() }
func (x StrsStrsSrtDsc) StrsStrs() *strs.Strs { return x.X.StrsStrs().SrtDsc() }
func (x BolsBolsCnt) Act()                    { x.UntUnt() }
func (x BolsBolsCnt) Ifc() interface{}        { return x.UntUnt() }
func (x BolsBolsCnt) UntUnt() unt.Unt         { return x.X.BolsBols().Cnt() }
func (x BolsBolsCpy) Act()                    { x.BolsBols() }
func (x BolsBolsCpy) Ifc() interface{}        { return x.BolsBols() }
func (x BolsBolsCpy) BolsBols() *bols.Bols    { return x.X.BolsBols().Cpy() }
func (x BolsBolsClr) Act()                    { x.BolsBols() }
func (x BolsBolsClr) Ifc() interface{}        { return x.BolsBols() }
func (x BolsBolsClr) BolsBols() *bols.Bols    { return x.X.BolsBols().Clr() }
func (x BolsBolsRand) Act()                   { x.BolsBols() }
func (x BolsBolsRand) Ifc() interface{}       { return x.BolsBols() }
func (x BolsBolsRand) BolsBols() *bols.Bols   { return x.X.BolsBols().Rand() }
func (x BolsBolsMrg) Act()                    { x.BolsBols() }
func (x BolsBolsMrg) Ifc() interface{}        { return x.BolsBols() }
func (x BolsBolsMrg) BolsBols() *bols.Bols {
	var i0 []*bols.Bols
	for _, cur := range x.I0 {
		i0 = append(i0, cur.BolsBols())
	}
	return x.X.BolsBols().Mrg(i0...)
}
func (x BolsBolsPush) Act()             { x.BolsBols() }
func (x BolsBolsPush) Ifc() interface{} { return x.BolsBols() }
func (x BolsBolsPush) BolsBols() *bols.Bols {
	var i0 []bol.Bol
	for _, cur := range x.I0 {
		i0 = append(i0, cur.BolBol())
	}
	return x.X.BolsBols().Push(i0...)
}
func (x BolsBolsPop) Act()             { x.BolBol() }
func (x BolsBolsPop) Ifc() interface{} { return x.BolBol() }
func (x BolsBolsPop) BolBol() bol.Bol  { return x.X.BolsBols().Pop() }
func (x BolsBolsQue) Act()             { x.BolsBols() }
func (x BolsBolsQue) Ifc() interface{} { return x.BolsBols() }
func (x BolsBolsQue) BolsBols() *bols.Bols {
	var i0 []bol.Bol
	for _, cur := range x.I0 {
		i0 = append(i0, cur.BolBol())
	}
	return x.X.BolsBols().Que(i0...)
}
func (x BolsBolsDque) Act()                  { x.BolBol() }
func (x BolsBolsDque) Ifc() interface{}      { return x.BolBol() }
func (x BolsBolsDque) BolBol() bol.Bol       { return x.X.BolsBols().Dque() }
func (x BolsBolsIns) Act()                   { x.BolsBols() }
func (x BolsBolsIns) Ifc() interface{}       { return x.BolsBols() }
func (x BolsBolsIns) BolsBols() *bols.Bols   { return x.X.BolsBols().Ins(x.I0.UntUnt(), x.I1.BolBol()) }
func (x BolsBolsUpd) Act()                   { x.BolsBols() }
func (x BolsBolsUpd) Ifc() interface{}       { return x.BolsBols() }
func (x BolsBolsUpd) BolsBols() *bols.Bols   { return x.X.BolsBols().Upd(x.I0.UntUnt(), x.I1.BolBol()) }
func (x BolsBolsDel) Act()                   { x.BolBol() }
func (x BolsBolsDel) Ifc() interface{}       { return x.BolBol() }
func (x BolsBolsDel) BolBol() bol.Bol        { return x.X.BolsBols().Del(x.I0.UntUnt()) }
func (x BolsBolsAt) Act()                    { x.BolBol() }
func (x BolsBolsAt) Ifc() interface{}        { return x.BolBol() }
func (x BolsBolsAt) BolBol() bol.Bol         { return x.X.BolsBols().At(x.I0.UntUnt()) }
func (x BolsBolsIn) Act()                    { x.BolsBols() }
func (x BolsBolsIn) Ifc() interface{}        { return x.BolsBols() }
func (x BolsBolsIn) BolsBols() *bols.Bols    { return x.X.BolsBols().In(x.I0.UntUnt(), x.I1.UntUnt()) }
func (x BolsBolsInBnd) Act()                 { x.BolsBols() }
func (x BolsBolsInBnd) Ifc() interface{}     { return x.BolsBols() }
func (x BolsBolsInBnd) BolsBols() *bols.Bols { return x.X.BolsBols().InBnd(x.I0.BndBnd()) }
func (x BolsBolsFrom) Act()                  { x.BolsBols() }
func (x BolsBolsFrom) Ifc() interface{}      { return x.BolsBols() }
func (x BolsBolsFrom) BolsBols() *bols.Bols  { return x.X.BolsBols().From(x.I0.UntUnt()) }
func (x BolsBolsTo) Act()                    { x.BolsBols() }
func (x BolsBolsTo) Ifc() interface{}        { return x.BolsBols() }
func (x BolsBolsTo) BolsBols() *bols.Bols    { return x.X.BolsBols().To(x.I0.UntUnt()) }
func (x BolsBolsFst) Act()                   { x.BolBol() }
func (x BolsBolsFst) Ifc() interface{}       { return x.BolBol() }
func (x BolsBolsFst) BolBol() bol.Bol        { return x.X.BolsBols().Fst() }
func (x BolsBolsMdl) Act()                   { x.BolBol() }
func (x BolsBolsMdl) Ifc() interface{}       { return x.BolBol() }
func (x BolsBolsMdl) BolBol() bol.Bol        { return x.X.BolsBols().Mdl() }
func (x BolsBolsLst) Act()                   { x.BolBol() }
func (x BolsBolsLst) Ifc() interface{}       { return x.BolBol() }
func (x BolsBolsLst) BolBol() bol.Bol        { return x.X.BolsBols().Lst() }
func (x BolsBolsFstIdx) Act()                { x.UntUnt() }
func (x BolsBolsFstIdx) Ifc() interface{}    { return x.UntUnt() }
func (x BolsBolsFstIdx) UntUnt() unt.Unt     { return x.X.BolsBols().FstIdx() }
func (x BolsBolsMdlIdx) Act()                { x.UntUnt() }
func (x BolsBolsMdlIdx) Ifc() interface{}    { return x.UntUnt() }
func (x BolsBolsMdlIdx) UntUnt() unt.Unt     { return x.X.BolsBols().MdlIdx() }
func (x BolsBolsLstIdx) Act()                { x.UntUnt() }
func (x BolsBolsLstIdx) Ifc() interface{}    { return x.UntUnt() }
func (x BolsBolsLstIdx) UntUnt() unt.Unt     { return x.X.BolsBols().LstIdx() }
func (x BolsBolsRev) Act()                   { x.BolsBols() }
func (x BolsBolsRev) Ifc() interface{}       { return x.BolsBols() }
func (x BolsBolsRev) BolsBols() *bols.Bols   { return x.X.BolsBols().Rev() }
func (x FltsFltsCnt) Act()                   { x.UntUnt() }
func (x FltsFltsCnt) Ifc() interface{}       { return x.UntUnt() }
func (x FltsFltsCnt) UntUnt() unt.Unt        { return x.X.FltsFlts().Cnt() }
func (x FltsFltsCpy) Act()                   { x.FltsFlts() }
func (x FltsFltsCpy) Ifc() interface{}       { return x.FltsFlts() }
func (x FltsFltsCpy) FltsFlts() *flts.Flts   { return x.X.FltsFlts().Cpy() }
func (x FltsFltsClr) Act()                   { x.FltsFlts() }
func (x FltsFltsClr) Ifc() interface{}       { return x.FltsFlts() }
func (x FltsFltsClr) FltsFlts() *flts.Flts   { return x.X.FltsFlts().Clr() }
func (x FltsFltsRand) Act()                  { x.FltsFlts() }
func (x FltsFltsRand) Ifc() interface{}      { return x.FltsFlts() }
func (x FltsFltsRand) FltsFlts() *flts.Flts  { return x.X.FltsFlts().Rand() }
func (x FltsFltsMrg) Act()                   { x.FltsFlts() }
func (x FltsFltsMrg) Ifc() interface{}       { return x.FltsFlts() }
func (x FltsFltsMrg) FltsFlts() *flts.Flts {
	var i0 []*flts.Flts
	for _, cur := range x.I0 {
		i0 = append(i0, cur.FltsFlts())
	}
	return x.X.FltsFlts().Mrg(i0...)
}
func (x FltsFltsPush) Act()             { x.FltsFlts() }
func (x FltsFltsPush) Ifc() interface{} { return x.FltsFlts() }
func (x FltsFltsPush) FltsFlts() *flts.Flts {
	var i0 []flt.Flt
	for _, cur := range x.I0 {
		i0 = append(i0, cur.FltFlt())
	}
	return x.X.FltsFlts().Push(i0...)
}
func (x FltsFltsPop) Act()             { x.FltFlt() }
func (x FltsFltsPop) Ifc() interface{} { return x.FltFlt() }
func (x FltsFltsPop) FltFlt() flt.Flt  { return x.X.FltsFlts().Pop() }
func (x FltsFltsQue) Act()             { x.FltsFlts() }
func (x FltsFltsQue) Ifc() interface{} { return x.FltsFlts() }
func (x FltsFltsQue) FltsFlts() *flts.Flts {
	var i0 []flt.Flt
	for _, cur := range x.I0 {
		i0 = append(i0, cur.FltFlt())
	}
	return x.X.FltsFlts().Que(i0...)
}
func (x FltsFltsDque) Act()                   { x.FltFlt() }
func (x FltsFltsDque) Ifc() interface{}       { return x.FltFlt() }
func (x FltsFltsDque) FltFlt() flt.Flt        { return x.X.FltsFlts().Dque() }
func (x FltsFltsIns) Act()                    { x.FltsFlts() }
func (x FltsFltsIns) Ifc() interface{}        { return x.FltsFlts() }
func (x FltsFltsIns) FltsFlts() *flts.Flts    { return x.X.FltsFlts().Ins(x.I0.UntUnt(), x.I1.FltFlt()) }
func (x FltsFltsUpd) Act()                    { x.FltsFlts() }
func (x FltsFltsUpd) Ifc() interface{}        { return x.FltsFlts() }
func (x FltsFltsUpd) FltsFlts() *flts.Flts    { return x.X.FltsFlts().Upd(x.I0.UntUnt(), x.I1.FltFlt()) }
func (x FltsFltsDel) Act()                    { x.FltFlt() }
func (x FltsFltsDel) Ifc() interface{}        { return x.FltFlt() }
func (x FltsFltsDel) FltFlt() flt.Flt         { return x.X.FltsFlts().Del(x.I0.UntUnt()) }
func (x FltsFltsAt) Act()                     { x.FltFlt() }
func (x FltsFltsAt) Ifc() interface{}         { return x.FltFlt() }
func (x FltsFltsAt) FltFlt() flt.Flt          { return x.X.FltsFlts().At(x.I0.UntUnt()) }
func (x FltsFltsIn) Act()                     { x.FltsFlts() }
func (x FltsFltsIn) Ifc() interface{}         { return x.FltsFlts() }
func (x FltsFltsIn) FltsFlts() *flts.Flts     { return x.X.FltsFlts().In(x.I0.UntUnt(), x.I1.UntUnt()) }
func (x FltsFltsInBnd) Act()                  { x.FltsFlts() }
func (x FltsFltsInBnd) Ifc() interface{}      { return x.FltsFlts() }
func (x FltsFltsInBnd) FltsFlts() *flts.Flts  { return x.X.FltsFlts().InBnd(x.I0.BndBnd()) }
func (x FltsFltsFrom) Act()                   { x.FltsFlts() }
func (x FltsFltsFrom) Ifc() interface{}       { return x.FltsFlts() }
func (x FltsFltsFrom) FltsFlts() *flts.Flts   { return x.X.FltsFlts().From(x.I0.UntUnt()) }
func (x FltsFltsTo) Act()                     { x.FltsFlts() }
func (x FltsFltsTo) Ifc() interface{}         { return x.FltsFlts() }
func (x FltsFltsTo) FltsFlts() *flts.Flts     { return x.X.FltsFlts().To(x.I0.UntUnt()) }
func (x FltsFltsFst) Act()                    { x.FltFlt() }
func (x FltsFltsFst) Ifc() interface{}        { return x.FltFlt() }
func (x FltsFltsFst) FltFlt() flt.Flt         { return x.X.FltsFlts().Fst() }
func (x FltsFltsMdl) Act()                    { x.FltFlt() }
func (x FltsFltsMdl) Ifc() interface{}        { return x.FltFlt() }
func (x FltsFltsMdl) FltFlt() flt.Flt         { return x.X.FltsFlts().Mdl() }
func (x FltsFltsLst) Act()                    { x.FltFlt() }
func (x FltsFltsLst) Ifc() interface{}        { return x.FltFlt() }
func (x FltsFltsLst) FltFlt() flt.Flt         { return x.X.FltsFlts().Lst() }
func (x FltsFltsFstIdx) Act()                 { x.UntUnt() }
func (x FltsFltsFstIdx) Ifc() interface{}     { return x.UntUnt() }
func (x FltsFltsFstIdx) UntUnt() unt.Unt      { return x.X.FltsFlts().FstIdx() }
func (x FltsFltsMdlIdx) Act()                 { x.UntUnt() }
func (x FltsFltsMdlIdx) Ifc() interface{}     { return x.UntUnt() }
func (x FltsFltsMdlIdx) UntUnt() unt.Unt      { return x.X.FltsFlts().MdlIdx() }
func (x FltsFltsLstIdx) Act()                 { x.UntUnt() }
func (x FltsFltsLstIdx) Ifc() interface{}     { return x.UntUnt() }
func (x FltsFltsLstIdx) UntUnt() unt.Unt      { return x.X.FltsFlts().LstIdx() }
func (x FltsFltsRev) Act()                    { x.FltsFlts() }
func (x FltsFltsRev) Ifc() interface{}        { return x.FltsFlts() }
func (x FltsFltsRev) FltsFlts() *flts.Flts    { return x.X.FltsFlts().Rev() }
func (x FltsFltsSrchIdxEql) Act()             { x.UntUnt() }
func (x FltsFltsSrchIdxEql) Ifc() interface{} { return x.UntUnt() }
func (x FltsFltsSrchIdxEql) UntUnt() unt.Unt  { return x.X.FltsFlts().SrchIdxEql(x.I0.FltFlt()) }
func (x FltsFltsSrchIdx) Act()                { x.UntUnt() }
func (x FltsFltsSrchIdx) Ifc() interface{}    { return x.UntUnt() }
func (x FltsFltsSrchIdx) UntUnt() unt.Unt {
	var i1 []bol.Bol
	for _, cur := range x.I1 {
		i1 = append(i1, cur.BolBol())
	}
	return x.X.FltsFlts().SrchIdx(x.I0.FltFlt(), i1...)
}
func (x FltsFltsHas) Act()                         { x.BolBol() }
func (x FltsFltsHas) Ifc() interface{}             { return x.BolBol() }
func (x FltsFltsHas) BolBol() bol.Bol              { return x.X.FltsFlts().Has(x.I0.FltFlt()) }
func (x FltsFltsSrtAsc) Act()                      { x.FltsFlts() }
func (x FltsFltsSrtAsc) Ifc() interface{}          { return x.FltsFlts() }
func (x FltsFltsSrtAsc) FltsFlts() *flts.Flts      { return x.X.FltsFlts().SrtAsc() }
func (x FltsFltsSrtDsc) Act()                      { x.FltsFlts() }
func (x FltsFltsSrtDsc) Ifc() interface{}          { return x.FltsFlts() }
func (x FltsFltsSrtDsc) FltsFlts() *flts.Flts      { return x.X.FltsFlts().SrtDsc() }
func (x FltsFltsUnaPos) Act()                      { x.FltsFlts() }
func (x FltsFltsUnaPos) Ifc() interface{}          { return x.FltsFlts() }
func (x FltsFltsUnaPos) FltsFlts() *flts.Flts      { return x.X.FltsFlts().UnaPos() }
func (x FltsFltsUnaNeg) Act()                      { x.FltsFlts() }
func (x FltsFltsUnaNeg) Ifc() interface{}          { return x.FltsFlts() }
func (x FltsFltsUnaNeg) FltsFlts() *flts.Flts      { return x.X.FltsFlts().UnaNeg() }
func (x FltsFltsUnaInv) Act()                      { x.FltsFlts() }
func (x FltsFltsUnaInv) Ifc() interface{}          { return x.FltsFlts() }
func (x FltsFltsUnaInv) FltsFlts() *flts.Flts      { return x.X.FltsFlts().UnaInv() }
func (x FltsFltsUnaSqr) Act()                      { x.FltsFlts() }
func (x FltsFltsUnaSqr) Ifc() interface{}          { return x.FltsFlts() }
func (x FltsFltsUnaSqr) FltsFlts() *flts.Flts      { return x.X.FltsFlts().UnaSqr() }
func (x FltsFltsUnaSqrt) Act()                     { x.FltsFlts() }
func (x FltsFltsUnaSqrt) Ifc() interface{}         { return x.FltsFlts() }
func (x FltsFltsUnaSqrt) FltsFlts() *flts.Flts     { return x.X.FltsFlts().UnaSqrt() }
func (x FltsFltsSclAdd) Act()                      { x.FltsFlts() }
func (x FltsFltsSclAdd) Ifc() interface{}          { return x.FltsFlts() }
func (x FltsFltsSclAdd) FltsFlts() *flts.Flts      { return x.X.FltsFlts().SclAdd(x.I0.FltFlt()) }
func (x FltsFltsSclSub) Act()                      { x.FltsFlts() }
func (x FltsFltsSclSub) Ifc() interface{}          { return x.FltsFlts() }
func (x FltsFltsSclSub) FltsFlts() *flts.Flts      { return x.X.FltsFlts().SclSub(x.I0.FltFlt()) }
func (x FltsFltsSclMul) Act()                      { x.FltsFlts() }
func (x FltsFltsSclMul) Ifc() interface{}          { return x.FltsFlts() }
func (x FltsFltsSclMul) FltsFlts() *flts.Flts      { return x.X.FltsFlts().SclMul(x.I0.FltFlt()) }
func (x FltsFltsSclDiv) Act()                      { x.FltsFlts() }
func (x FltsFltsSclDiv) Ifc() interface{}          { return x.FltsFlts() }
func (x FltsFltsSclDiv) FltsFlts() *flts.Flts      { return x.X.FltsFlts().SclDiv(x.I0.FltFlt()) }
func (x FltsFltsSclRem) Act()                      { x.FltsFlts() }
func (x FltsFltsSclRem) Ifc() interface{}          { return x.FltsFlts() }
func (x FltsFltsSclRem) FltsFlts() *flts.Flts      { return x.X.FltsFlts().SclRem(x.I0.FltFlt()) }
func (x FltsFltsSclPow) Act()                      { x.FltsFlts() }
func (x FltsFltsSclPow) Ifc() interface{}          { return x.FltsFlts() }
func (x FltsFltsSclPow) FltsFlts() *flts.Flts      { return x.X.FltsFlts().SclPow(x.I0.FltFlt()) }
func (x FltsFltsSclMin) Act()                      { x.FltsFlts() }
func (x FltsFltsSclMin) Ifc() interface{}          { return x.FltsFlts() }
func (x FltsFltsSclMin) FltsFlts() *flts.Flts      { return x.X.FltsFlts().SclMin(x.I0.FltFlt()) }
func (x FltsFltsSclMax) Act()                      { x.FltsFlts() }
func (x FltsFltsSclMax) Ifc() interface{}          { return x.FltsFlts() }
func (x FltsFltsSclMax) FltsFlts() *flts.Flts      { return x.X.FltsFlts().SclMax(x.I0.FltFlt()) }
func (x FltsFltsSelEql) Act()                      { x.FltsFlts() }
func (x FltsFltsSelEql) Ifc() interface{}          { return x.FltsFlts() }
func (x FltsFltsSelEql) FltsFlts() *flts.Flts      { return x.X.FltsFlts().SelEql(x.I0.FltFlt()) }
func (x FltsFltsSelNeq) Act()                      { x.FltsFlts() }
func (x FltsFltsSelNeq) Ifc() interface{}          { return x.FltsFlts() }
func (x FltsFltsSelNeq) FltsFlts() *flts.Flts      { return x.X.FltsFlts().SelNeq(x.I0.FltFlt()) }
func (x FltsFltsSelLss) Act()                      { x.FltsFlts() }
func (x FltsFltsSelLss) Ifc() interface{}          { return x.FltsFlts() }
func (x FltsFltsSelLss) FltsFlts() *flts.Flts      { return x.X.FltsFlts().SelLss(x.I0.FltFlt()) }
func (x FltsFltsSelGtr) Act()                      { x.FltsFlts() }
func (x FltsFltsSelGtr) Ifc() interface{}          { return x.FltsFlts() }
func (x FltsFltsSelGtr) FltsFlts() *flts.Flts      { return x.X.FltsFlts().SelGtr(x.I0.FltFlt()) }
func (x FltsFltsSelLeq) Act()                      { x.FltsFlts() }
func (x FltsFltsSelLeq) Ifc() interface{}          { return x.FltsFlts() }
func (x FltsFltsSelLeq) FltsFlts() *flts.Flts      { return x.X.FltsFlts().SelLeq(x.I0.FltFlt()) }
func (x FltsFltsSelGeq) Act()                      { x.FltsFlts() }
func (x FltsFltsSelGeq) Ifc() interface{}          { return x.FltsFlts() }
func (x FltsFltsSelGeq) FltsFlts() *flts.Flts      { return x.X.FltsFlts().SelGeq(x.I0.FltFlt()) }
func (x FltsFltsCntEql) Act()                      { x.FltFlt() }
func (x FltsFltsCntEql) Ifc() interface{}          { return x.FltFlt() }
func (x FltsFltsCntEql) FltFlt() flt.Flt           { return x.X.FltsFlts().CntEql(x.I0.FltFlt()) }
func (x FltsFltsCntNeq) Act()                      { x.FltFlt() }
func (x FltsFltsCntNeq) Ifc() interface{}          { return x.FltFlt() }
func (x FltsFltsCntNeq) FltFlt() flt.Flt           { return x.X.FltsFlts().CntNeq(x.I0.FltFlt()) }
func (x FltsFltsCntLss) Act()                      { x.FltFlt() }
func (x FltsFltsCntLss) Ifc() interface{}          { return x.FltFlt() }
func (x FltsFltsCntLss) FltFlt() flt.Flt           { return x.X.FltsFlts().CntLss(x.I0.FltFlt()) }
func (x FltsFltsCntGtr) Act()                      { x.FltFlt() }
func (x FltsFltsCntGtr) Ifc() interface{}          { return x.FltFlt() }
func (x FltsFltsCntGtr) FltFlt() flt.Flt           { return x.X.FltsFlts().CntGtr(x.I0.FltFlt()) }
func (x FltsFltsCntLeq) Act()                      { x.FltFlt() }
func (x FltsFltsCntLeq) Ifc() interface{}          { return x.FltFlt() }
func (x FltsFltsCntLeq) FltFlt() flt.Flt           { return x.X.FltsFlts().CntLeq(x.I0.FltFlt()) }
func (x FltsFltsCntGeq) Act()                      { x.FltFlt() }
func (x FltsFltsCntGeq) Ifc() interface{}          { return x.FltFlt() }
func (x FltsFltsCntGeq) FltFlt() flt.Flt           { return x.X.FltsFlts().CntGeq(x.I0.FltFlt()) }
func (x FltsFltsInrAdd) Act()                      { x.FltsFlts() }
func (x FltsFltsInrAdd) Ifc() interface{}          { return x.FltsFlts() }
func (x FltsFltsInrAdd) FltsFlts() *flts.Flts      { return x.X.FltsFlts().InrAdd(x.I0.UntUnt()) }
func (x FltsFltsInrSub) Act()                      { x.FltsFlts() }
func (x FltsFltsInrSub) Ifc() interface{}          { return x.FltsFlts() }
func (x FltsFltsInrSub) FltsFlts() *flts.Flts      { return x.X.FltsFlts().InrSub(x.I0.UntUnt()) }
func (x FltsFltsInrMul) Act()                      { x.FltsFlts() }
func (x FltsFltsInrMul) Ifc() interface{}          { return x.FltsFlts() }
func (x FltsFltsInrMul) FltsFlts() *flts.Flts      { return x.X.FltsFlts().InrMul(x.I0.UntUnt()) }
func (x FltsFltsInrDiv) Act()                      { x.FltsFlts() }
func (x FltsFltsInrDiv) Ifc() interface{}          { return x.FltsFlts() }
func (x FltsFltsInrDiv) FltsFlts() *flts.Flts      { return x.X.FltsFlts().InrDiv(x.I0.UntUnt()) }
func (x FltsFltsInrRem) Act()                      { x.FltsFlts() }
func (x FltsFltsInrRem) Ifc() interface{}          { return x.FltsFlts() }
func (x FltsFltsInrRem) FltsFlts() *flts.Flts      { return x.X.FltsFlts().InrRem(x.I0.UntUnt()) }
func (x FltsFltsInrPow) Act()                      { x.FltsFlts() }
func (x FltsFltsInrPow) Ifc() interface{}          { return x.FltsFlts() }
func (x FltsFltsInrPow) FltsFlts() *flts.Flts      { return x.X.FltsFlts().InrPow(x.I0.UntUnt()) }
func (x FltsFltsInrMin) Act()                      { x.FltsFlts() }
func (x FltsFltsInrMin) Ifc() interface{}          { return x.FltsFlts() }
func (x FltsFltsInrMin) FltsFlts() *flts.Flts      { return x.X.FltsFlts().InrMin(x.I0.UntUnt()) }
func (x FltsFltsInrMax) Act()                      { x.FltsFlts() }
func (x FltsFltsInrMax) Ifc() interface{}          { return x.FltsFlts() }
func (x FltsFltsInrMax) FltsFlts() *flts.Flts      { return x.X.FltsFlts().InrMax(x.I0.UntUnt()) }
func (x FltsFltsSum) Act()                         { x.FltFlt() }
func (x FltsFltsSum) Ifc() interface{}             { return x.FltFlt() }
func (x FltsFltsSum) FltFlt() flt.Flt              { return x.X.FltsFlts().Sum() }
func (x FltsFltsPrd) Act()                         { x.FltFlt() }
func (x FltsFltsPrd) Ifc() interface{}             { return x.FltFlt() }
func (x FltsFltsPrd) FltFlt() flt.Flt              { return x.X.FltsFlts().Prd() }
func (x FltsFltsMin) Act()                         { x.FltFlt() }
func (x FltsFltsMin) Ifc() interface{}             { return x.FltFlt() }
func (x FltsFltsMin) FltFlt() flt.Flt              { return x.X.FltsFlts().Min() }
func (x FltsFltsMax) Act()                         { x.FltFlt() }
func (x FltsFltsMax) Ifc() interface{}             { return x.FltFlt() }
func (x FltsFltsMax) FltFlt() flt.Flt              { return x.X.FltsFlts().Max() }
func (x FltsFltsMid) Act()                         { x.FltFlt() }
func (x FltsFltsMid) Ifc() interface{}             { return x.FltFlt() }
func (x FltsFltsMid) FltFlt() flt.Flt              { return x.X.FltsFlts().Mid() }
func (x FltsFltsMdn) Act()                         { x.FltFlt() }
func (x FltsFltsMdn) Ifc() interface{}             { return x.FltFlt() }
func (x FltsFltsMdn) FltFlt() flt.Flt              { return x.X.FltsFlts().Mdn() }
func (x FltsFltsSma) Act()                         { x.FltFlt() }
func (x FltsFltsSma) Ifc() interface{}             { return x.FltFlt() }
func (x FltsFltsSma) FltFlt() flt.Flt              { return x.X.FltsFlts().Sma() }
func (x FltsFltsGma) Act()                         { x.FltFlt() }
func (x FltsFltsGma) Ifc() interface{}             { return x.FltFlt() }
func (x FltsFltsGma) FltFlt() flt.Flt              { return x.X.FltsFlts().Gma() }
func (x FltsFltsWma) Act()                         { x.FltFlt() }
func (x FltsFltsWma) Ifc() interface{}             { return x.FltFlt() }
func (x FltsFltsWma) FltFlt() flt.Flt              { return x.X.FltsFlts().Wma() }
func (x FltsFltsVrnc) Act()                        { x.FltFlt() }
func (x FltsFltsVrnc) Ifc() interface{}            { return x.FltFlt() }
func (x FltsFltsVrnc) FltFlt() flt.Flt             { return x.X.FltsFlts().Vrnc() }
func (x FltsFltsStd) Act()                         { x.FltFlt() }
func (x FltsFltsStd) Ifc() interface{}             { return x.FltFlt() }
func (x FltsFltsStd) FltFlt() flt.Flt              { return x.X.FltsFlts().Std() }
func (x FltsFltsZscr) Act()                        { x.FltsFlts() }
func (x FltsFltsZscr) Ifc() interface{}            { return x.FltsFlts() }
func (x FltsFltsZscr) FltsFlts() *flts.Flts        { return x.X.FltsFlts().Zscr() }
func (x FltsFltsZscrInplace) Act()                 { x.FltsFlts() }
func (x FltsFltsZscrInplace) Ifc() interface{}     { return x.FltsFlts() }
func (x FltsFltsZscrInplace) FltsFlts() *flts.Flts { return x.X.FltsFlts().ZscrInplace() }
func (x FltsFltsRngFul) Act()                      { x.FltFlt() }
func (x FltsFltsRngFul) Ifc() interface{}          { return x.FltFlt() }
func (x FltsFltsRngFul) FltFlt() flt.Flt           { return x.X.FltsFlts().RngFul() }
func (x FltsFltsRngLst) Act()                      { x.FltFlt() }
func (x FltsFltsRngLst) Ifc() interface{}          { return x.FltFlt() }
func (x FltsFltsRngLst) FltFlt() flt.Flt           { return x.X.FltsFlts().RngLst() }
func (x FltsFltsProLst) Act()                      { x.FltFlt() }
func (x FltsFltsProLst) Ifc() interface{}          { return x.FltFlt() }
func (x FltsFltsProLst) FltFlt() flt.Flt           { return x.X.FltsFlts().ProLst() }
func (x FltsFltsProSma) Act()                      { x.FltFlt() }
func (x FltsFltsProSma) Ifc() interface{}          { return x.FltFlt() }
func (x FltsFltsProSma) FltFlt() flt.Flt           { return x.X.FltsFlts().ProSma() }
func (x FltsFltsSubSumPos) Act()                   { x.FltFlt() }
func (x FltsFltsSubSumPos) Ifc() interface{}       { return x.FltFlt() }
func (x FltsFltsSubSumPos) FltFlt() flt.Flt        { return x.X.FltsFlts().SubSumPos() }
func (x FltsFltsSubSumNeg) Act()                   { x.FltFlt() }
func (x FltsFltsSubSumNeg) Ifc() interface{}       { return x.FltFlt() }
func (x FltsFltsSubSumNeg) FltFlt() flt.Flt        { return x.X.FltsFlts().SubSumNeg() }
func (x FltsFltsRsi) Act()                         { x.FltFlt() }
func (x FltsFltsRsi) Ifc() interface{}             { return x.FltFlt() }
func (x FltsFltsRsi) FltFlt() flt.Flt              { return x.X.FltsFlts().Rsi() }
func (x FltsFltsWrsi) Act()                        { x.FltFlt() }
func (x FltsFltsWrsi) Ifc() interface{}            { return x.FltFlt() }
func (x FltsFltsWrsi) FltFlt() flt.Flt             { return x.X.FltsFlts().Wrsi() }
func (x FltsFltsPro) Act()                         { x.FltsFlts() }
func (x FltsFltsPro) Ifc() interface{}             { return x.FltsFlts() }
func (x FltsFltsPro) FltsFlts() *flts.Flts         { return x.X.FltsFlts().Pro() }
func (x FltsFltsAlma) Act()                        { x.FltFlt() }
func (x FltsFltsAlma) Ifc() interface{}            { return x.FltFlt() }
func (x FltsFltsAlma) FltFlt() flt.Flt             { return x.X.FltsFlts().Alma() }
func (x FltsFltsProAlma) Act()                     { x.FltFlt() }
func (x FltsFltsProAlma) Ifc() interface{}         { return x.FltFlt() }
func (x FltsFltsProAlma) FltFlt() flt.Flt          { return x.X.FltsFlts().ProAlma() }
func (x FltsFltsCntrDist) Act()                    { x.FltsFlts() }
func (x FltsFltsCntrDist) Ifc() interface{}        { return x.FltsFlts() }
func (x FltsFltsCntrDist) FltsFlts() *flts.Flts {
	var i0 []bol.Bol
	for _, cur := range x.I0 {
		i0 = append(i0, cur.BolBol())
	}
	return x.X.FltsFlts().CntrDist(i0...)
}
func (x UntsUntsCnt) Act()                  { x.UntUnt() }
func (x UntsUntsCnt) Ifc() interface{}      { return x.UntUnt() }
func (x UntsUntsCnt) UntUnt() unt.Unt       { return x.X.UntsUnts().Cnt() }
func (x UntsUntsCpy) Act()                  { x.UntsUnts() }
func (x UntsUntsCpy) Ifc() interface{}      { return x.UntsUnts() }
func (x UntsUntsCpy) UntsUnts() *unts.Unts  { return x.X.UntsUnts().Cpy() }
func (x UntsUntsClr) Act()                  { x.UntsUnts() }
func (x UntsUntsClr) Ifc() interface{}      { return x.UntsUnts() }
func (x UntsUntsClr) UntsUnts() *unts.Unts  { return x.X.UntsUnts().Clr() }
func (x UntsUntsRand) Act()                 { x.UntsUnts() }
func (x UntsUntsRand) Ifc() interface{}     { return x.UntsUnts() }
func (x UntsUntsRand) UntsUnts() *unts.Unts { return x.X.UntsUnts().Rand() }
func (x UntsUntsMrg) Act()                  { x.UntsUnts() }
func (x UntsUntsMrg) Ifc() interface{}      { return x.UntsUnts() }
func (x UntsUntsMrg) UntsUnts() *unts.Unts {
	var i0 []*unts.Unts
	for _, cur := range x.I0 {
		i0 = append(i0, cur.UntsUnts())
	}
	return x.X.UntsUnts().Mrg(i0...)
}
func (x UntsUntsPush) Act()             { x.UntsUnts() }
func (x UntsUntsPush) Ifc() interface{} { return x.UntsUnts() }
func (x UntsUntsPush) UntsUnts() *unts.Unts {
	var i0 []unt.Unt
	for _, cur := range x.I0 {
		i0 = append(i0, cur.UntUnt())
	}
	return x.X.UntsUnts().Push(i0...)
}
func (x UntsUntsPop) Act()             { x.UntUnt() }
func (x UntsUntsPop) Ifc() interface{} { return x.UntUnt() }
func (x UntsUntsPop) UntUnt() unt.Unt  { return x.X.UntsUnts().Pop() }
func (x UntsUntsQue) Act()             { x.UntsUnts() }
func (x UntsUntsQue) Ifc() interface{} { return x.UntsUnts() }
func (x UntsUntsQue) UntsUnts() *unts.Unts {
	var i0 []unt.Unt
	for _, cur := range x.I0 {
		i0 = append(i0, cur.UntUnt())
	}
	return x.X.UntsUnts().Que(i0...)
}
func (x UntsUntsDque) Act()                   { x.UntUnt() }
func (x UntsUntsDque) Ifc() interface{}       { return x.UntUnt() }
func (x UntsUntsDque) UntUnt() unt.Unt        { return x.X.UntsUnts().Dque() }
func (x UntsUntsIns) Act()                    { x.UntsUnts() }
func (x UntsUntsIns) Ifc() interface{}        { return x.UntsUnts() }
func (x UntsUntsIns) UntsUnts() *unts.Unts    { return x.X.UntsUnts().Ins(x.I0.UntUnt(), x.I1.UntUnt()) }
func (x UntsUntsUpd) Act()                    { x.UntsUnts() }
func (x UntsUntsUpd) Ifc() interface{}        { return x.UntsUnts() }
func (x UntsUntsUpd) UntsUnts() *unts.Unts    { return x.X.UntsUnts().Upd(x.I0.UntUnt(), x.I1.UntUnt()) }
func (x UntsUntsDel) Act()                    { x.UntUnt() }
func (x UntsUntsDel) Ifc() interface{}        { return x.UntUnt() }
func (x UntsUntsDel) UntUnt() unt.Unt         { return x.X.UntsUnts().Del(x.I0.UntUnt()) }
func (x UntsUntsAt) Act()                     { x.UntUnt() }
func (x UntsUntsAt) Ifc() interface{}         { return x.UntUnt() }
func (x UntsUntsAt) UntUnt() unt.Unt          { return x.X.UntsUnts().At(x.I0.UntUnt()) }
func (x UntsUntsIn) Act()                     { x.UntsUnts() }
func (x UntsUntsIn) Ifc() interface{}         { return x.UntsUnts() }
func (x UntsUntsIn) UntsUnts() *unts.Unts     { return x.X.UntsUnts().In(x.I0.UntUnt(), x.I1.UntUnt()) }
func (x UntsUntsInBnd) Act()                  { x.UntsUnts() }
func (x UntsUntsInBnd) Ifc() interface{}      { return x.UntsUnts() }
func (x UntsUntsInBnd) UntsUnts() *unts.Unts  { return x.X.UntsUnts().InBnd(x.I0.BndBnd()) }
func (x UntsUntsFrom) Act()                   { x.UntsUnts() }
func (x UntsUntsFrom) Ifc() interface{}       { return x.UntsUnts() }
func (x UntsUntsFrom) UntsUnts() *unts.Unts   { return x.X.UntsUnts().From(x.I0.UntUnt()) }
func (x UntsUntsTo) Act()                     { x.UntsUnts() }
func (x UntsUntsTo) Ifc() interface{}         { return x.UntsUnts() }
func (x UntsUntsTo) UntsUnts() *unts.Unts     { return x.X.UntsUnts().To(x.I0.UntUnt()) }
func (x UntsUntsFst) Act()                    { x.UntUnt() }
func (x UntsUntsFst) Ifc() interface{}        { return x.UntUnt() }
func (x UntsUntsFst) UntUnt() unt.Unt         { return x.X.UntsUnts().Fst() }
func (x UntsUntsMdl) Act()                    { x.UntUnt() }
func (x UntsUntsMdl) Ifc() interface{}        { return x.UntUnt() }
func (x UntsUntsMdl) UntUnt() unt.Unt         { return x.X.UntsUnts().Mdl() }
func (x UntsUntsLst) Act()                    { x.UntUnt() }
func (x UntsUntsLst) Ifc() interface{}        { return x.UntUnt() }
func (x UntsUntsLst) UntUnt() unt.Unt         { return x.X.UntsUnts().Lst() }
func (x UntsUntsFstIdx) Act()                 { x.UntUnt() }
func (x UntsUntsFstIdx) Ifc() interface{}     { return x.UntUnt() }
func (x UntsUntsFstIdx) UntUnt() unt.Unt      { return x.X.UntsUnts().FstIdx() }
func (x UntsUntsMdlIdx) Act()                 { x.UntUnt() }
func (x UntsUntsMdlIdx) Ifc() interface{}     { return x.UntUnt() }
func (x UntsUntsMdlIdx) UntUnt() unt.Unt      { return x.X.UntsUnts().MdlIdx() }
func (x UntsUntsLstIdx) Act()                 { x.UntUnt() }
func (x UntsUntsLstIdx) Ifc() interface{}     { return x.UntUnt() }
func (x UntsUntsLstIdx) UntUnt() unt.Unt      { return x.X.UntsUnts().LstIdx() }
func (x UntsUntsRev) Act()                    { x.UntsUnts() }
func (x UntsUntsRev) Ifc() interface{}        { return x.UntsUnts() }
func (x UntsUntsRev) UntsUnts() *unts.Unts    { return x.X.UntsUnts().Rev() }
func (x UntsUntsSrchIdxEql) Act()             { x.UntUnt() }
func (x UntsUntsSrchIdxEql) Ifc() interface{} { return x.UntUnt() }
func (x UntsUntsSrchIdxEql) UntUnt() unt.Unt  { return x.X.UntsUnts().SrchIdxEql(x.I0.UntUnt()) }
func (x UntsUntsSrchIdx) Act()                { x.UntUnt() }
func (x UntsUntsSrchIdx) Ifc() interface{}    { return x.UntUnt() }
func (x UntsUntsSrchIdx) UntUnt() unt.Unt {
	var i1 []bol.Bol
	for _, cur := range x.I1 {
		i1 = append(i1, cur.BolBol())
	}
	return x.X.UntsUnts().SrchIdx(x.I0.UntUnt(), i1...)
}
func (x UntsUntsHas) Act()                         { x.BolBol() }
func (x UntsUntsHas) Ifc() interface{}             { return x.BolBol() }
func (x UntsUntsHas) BolBol() bol.Bol              { return x.X.UntsUnts().Has(x.I0.UntUnt()) }
func (x UntsUntsSrtAsc) Act()                      { x.UntsUnts() }
func (x UntsUntsSrtAsc) Ifc() interface{}          { return x.UntsUnts() }
func (x UntsUntsSrtAsc) UntsUnts() *unts.Unts      { return x.X.UntsUnts().SrtAsc() }
func (x UntsUntsSrtDsc) Act()                      { x.UntsUnts() }
func (x UntsUntsSrtDsc) Ifc() interface{}          { return x.UntsUnts() }
func (x UntsUntsSrtDsc) UntsUnts() *unts.Unts      { return x.X.UntsUnts().SrtDsc() }
func (x UntsUntsInrAdd) Act()                      { x.UntsUnts() }
func (x UntsUntsInrAdd) Ifc() interface{}          { return x.UntsUnts() }
func (x UntsUntsInrAdd) UntsUnts() *unts.Unts      { return x.X.UntsUnts().InrAdd(x.I0.UntUnt()) }
func (x UntsUntsInrSub) Act()                      { x.UntsUnts() }
func (x UntsUntsInrSub) Ifc() interface{}          { return x.UntsUnts() }
func (x UntsUntsInrSub) UntsUnts() *unts.Unts      { return x.X.UntsUnts().InrSub(x.I0.UntUnt()) }
func (x UntsUntsInrMul) Act()                      { x.UntsUnts() }
func (x UntsUntsInrMul) Ifc() interface{}          { return x.UntsUnts() }
func (x UntsUntsInrMul) UntsUnts() *unts.Unts      { return x.X.UntsUnts().InrMul(x.I0.UntUnt()) }
func (x UntsUntsInrDiv) Act()                      { x.UntsUnts() }
func (x UntsUntsInrDiv) Ifc() interface{}          { return x.UntsUnts() }
func (x UntsUntsInrDiv) UntsUnts() *unts.Unts      { return x.X.UntsUnts().InrDiv(x.I0.UntUnt()) }
func (x UntsUntsInrRem) Act()                      { x.UntsUnts() }
func (x UntsUntsInrRem) Ifc() interface{}          { return x.UntsUnts() }
func (x UntsUntsInrRem) UntsUnts() *unts.Unts      { return x.X.UntsUnts().InrRem(x.I0.UntUnt()) }
func (x UntsUntsInrPow) Act()                      { x.UntsUnts() }
func (x UntsUntsInrPow) Ifc() interface{}          { return x.UntsUnts() }
func (x UntsUntsInrPow) UntsUnts() *unts.Unts      { return x.X.UntsUnts().InrPow(x.I0.UntUnt()) }
func (x UntsUntsInrMin) Act()                      { x.UntsUnts() }
func (x UntsUntsInrMin) Ifc() interface{}          { return x.UntsUnts() }
func (x UntsUntsInrMin) UntsUnts() *unts.Unts      { return x.X.UntsUnts().InrMin(x.I0.UntUnt()) }
func (x UntsUntsInrMax) Act()                      { x.UntsUnts() }
func (x UntsUntsInrMax) Ifc() interface{}          { return x.UntsUnts() }
func (x UntsUntsInrMax) UntsUnts() *unts.Unts      { return x.X.UntsUnts().InrMax(x.I0.UntUnt()) }
func (x UntsUntsSum) Act()                         { x.UntUnt() }
func (x UntsUntsSum) Ifc() interface{}             { return x.UntUnt() }
func (x UntsUntsSum) UntUnt() unt.Unt              { return x.X.UntsUnts().Sum() }
func (x UntsUntsPrd) Act()                         { x.UntUnt() }
func (x UntsUntsPrd) Ifc() interface{}             { return x.UntUnt() }
func (x UntsUntsPrd) UntUnt() unt.Unt              { return x.X.UntsUnts().Prd() }
func (x UntsUntsMin) Act()                         { x.UntUnt() }
func (x UntsUntsMin) Ifc() interface{}             { return x.UntUnt() }
func (x UntsUntsMin) UntUnt() unt.Unt              { return x.X.UntsUnts().Min() }
func (x UntsUntsMax) Act()                         { x.UntUnt() }
func (x UntsUntsMax) Ifc() interface{}             { return x.UntUnt() }
func (x UntsUntsMax) UntUnt() unt.Unt              { return x.X.UntsUnts().Max() }
func (x UntsUntsMid) Act()                         { x.UntUnt() }
func (x UntsUntsMid) Ifc() interface{}             { return x.UntUnt() }
func (x UntsUntsMid) UntUnt() unt.Unt              { return x.X.UntsUnts().Mid() }
func (x UntsUntsMdn) Act()                         { x.UntUnt() }
func (x UntsUntsMdn) Ifc() interface{}             { return x.UntUnt() }
func (x UntsUntsMdn) UntUnt() unt.Unt              { return x.X.UntsUnts().Mdn() }
func (x UntsUntsSma) Act()                         { x.UntUnt() }
func (x UntsUntsSma) Ifc() interface{}             { return x.UntUnt() }
func (x UntsUntsSma) UntUnt() unt.Unt              { return x.X.UntsUnts().Sma() }
func (x UntsUntsGma) Act()                         { x.UntUnt() }
func (x UntsUntsGma) Ifc() interface{}             { return x.UntUnt() }
func (x UntsUntsGma) UntUnt() unt.Unt              { return x.X.UntsUnts().Gma() }
func (x UntsUntsWma) Act()                         { x.UntUnt() }
func (x UntsUntsWma) Ifc() interface{}             { return x.UntUnt() }
func (x UntsUntsWma) UntUnt() unt.Unt              { return x.X.UntsUnts().Wma() }
func (x UntsUntsVrnc) Act()                        { x.UntUnt() }
func (x UntsUntsVrnc) Ifc() interface{}            { return x.UntUnt() }
func (x UntsUntsVrnc) UntUnt() unt.Unt             { return x.X.UntsUnts().Vrnc() }
func (x UntsUntsStd) Act()                         { x.UntUnt() }
func (x UntsUntsStd) Ifc() interface{}             { return x.UntUnt() }
func (x UntsUntsStd) UntUnt() unt.Unt              { return x.X.UntsUnts().Std() }
func (x UntsUntsZscr) Act()                        { x.UntsUnts() }
func (x UntsUntsZscr) Ifc() interface{}            { return x.UntsUnts() }
func (x UntsUntsZscr) UntsUnts() *unts.Unts        { return x.X.UntsUnts().Zscr() }
func (x UntsUntsZscrInplace) Act()                 { x.UntsUnts() }
func (x UntsUntsZscrInplace) Ifc() interface{}     { return x.UntsUnts() }
func (x UntsUntsZscrInplace) UntsUnts() *unts.Unts { return x.X.UntsUnts().ZscrInplace() }
func (x UntsUntsRngFul) Act()                      { x.UntUnt() }
func (x UntsUntsRngFul) Ifc() interface{}          { return x.UntUnt() }
func (x UntsUntsRngFul) UntUnt() unt.Unt           { return x.X.UntsUnts().RngFul() }
func (x UntsUntsRngLst) Act()                      { x.UntUnt() }
func (x UntsUntsRngLst) Ifc() interface{}          { return x.UntUnt() }
func (x UntsUntsRngLst) UntUnt() unt.Unt           { return x.X.UntsUnts().RngLst() }
func (x UntsUntsProLst) Act()                      { x.UntUnt() }
func (x UntsUntsProLst) Ifc() interface{}          { return x.UntUnt() }
func (x UntsUntsProLst) UntUnt() unt.Unt           { return x.X.UntsUnts().ProLst() }
func (x UntsUntsProSma) Act()                      { x.UntUnt() }
func (x UntsUntsProSma) Ifc() interface{}          { return x.UntUnt() }
func (x UntsUntsProSma) UntUnt() unt.Unt           { return x.X.UntsUnts().ProSma() }
func (x IntsIntsCnt) Act()                         { x.UntUnt() }
func (x IntsIntsCnt) Ifc() interface{}             { return x.UntUnt() }
func (x IntsIntsCnt) UntUnt() unt.Unt              { return x.X.IntsInts().Cnt() }
func (x IntsIntsCpy) Act()                         { x.IntsInts() }
func (x IntsIntsCpy) Ifc() interface{}             { return x.IntsInts() }
func (x IntsIntsCpy) IntsInts() *ints.Ints         { return x.X.IntsInts().Cpy() }
func (x IntsIntsClr) Act()                         { x.IntsInts() }
func (x IntsIntsClr) Ifc() interface{}             { return x.IntsInts() }
func (x IntsIntsClr) IntsInts() *ints.Ints         { return x.X.IntsInts().Clr() }
func (x IntsIntsRand) Act()                        { x.IntsInts() }
func (x IntsIntsRand) Ifc() interface{}            { return x.IntsInts() }
func (x IntsIntsRand) IntsInts() *ints.Ints        { return x.X.IntsInts().Rand() }
func (x IntsIntsMrg) Act()                         { x.IntsInts() }
func (x IntsIntsMrg) Ifc() interface{}             { return x.IntsInts() }
func (x IntsIntsMrg) IntsInts() *ints.Ints {
	var i0 []*ints.Ints
	for _, cur := range x.I0 {
		i0 = append(i0, cur.IntsInts())
	}
	return x.X.IntsInts().Mrg(i0...)
}
func (x IntsIntsPush) Act()             { x.IntsInts() }
func (x IntsIntsPush) Ifc() interface{} { return x.IntsInts() }
func (x IntsIntsPush) IntsInts() *ints.Ints {
	var i0 []int.Int
	for _, cur := range x.I0 {
		i0 = append(i0, cur.IntInt())
	}
	return x.X.IntsInts().Push(i0...)
}
func (x IntsIntsPop) Act()             { x.IntInt() }
func (x IntsIntsPop) Ifc() interface{} { return x.IntInt() }
func (x IntsIntsPop) IntInt() int.Int  { return x.X.IntsInts().Pop() }
func (x IntsIntsQue) Act()             { x.IntsInts() }
func (x IntsIntsQue) Ifc() interface{} { return x.IntsInts() }
func (x IntsIntsQue) IntsInts() *ints.Ints {
	var i0 []int.Int
	for _, cur := range x.I0 {
		i0 = append(i0, cur.IntInt())
	}
	return x.X.IntsInts().Que(i0...)
}
func (x IntsIntsDque) Act()                   { x.IntInt() }
func (x IntsIntsDque) Ifc() interface{}       { return x.IntInt() }
func (x IntsIntsDque) IntInt() int.Int        { return x.X.IntsInts().Dque() }
func (x IntsIntsIns) Act()                    { x.IntsInts() }
func (x IntsIntsIns) Ifc() interface{}        { return x.IntsInts() }
func (x IntsIntsIns) IntsInts() *ints.Ints    { return x.X.IntsInts().Ins(x.I0.UntUnt(), x.I1.IntInt()) }
func (x IntsIntsUpd) Act()                    { x.IntsInts() }
func (x IntsIntsUpd) Ifc() interface{}        { return x.IntsInts() }
func (x IntsIntsUpd) IntsInts() *ints.Ints    { return x.X.IntsInts().Upd(x.I0.UntUnt(), x.I1.IntInt()) }
func (x IntsIntsDel) Act()                    { x.IntInt() }
func (x IntsIntsDel) Ifc() interface{}        { return x.IntInt() }
func (x IntsIntsDel) IntInt() int.Int         { return x.X.IntsInts().Del(x.I0.UntUnt()) }
func (x IntsIntsAt) Act()                     { x.IntInt() }
func (x IntsIntsAt) Ifc() interface{}         { return x.IntInt() }
func (x IntsIntsAt) IntInt() int.Int          { return x.X.IntsInts().At(x.I0.UntUnt()) }
func (x IntsIntsIn) Act()                     { x.IntsInts() }
func (x IntsIntsIn) Ifc() interface{}         { return x.IntsInts() }
func (x IntsIntsIn) IntsInts() *ints.Ints     { return x.X.IntsInts().In(x.I0.UntUnt(), x.I1.UntUnt()) }
func (x IntsIntsInBnd) Act()                  { x.IntsInts() }
func (x IntsIntsInBnd) Ifc() interface{}      { return x.IntsInts() }
func (x IntsIntsInBnd) IntsInts() *ints.Ints  { return x.X.IntsInts().InBnd(x.I0.BndBnd()) }
func (x IntsIntsFrom) Act()                   { x.IntsInts() }
func (x IntsIntsFrom) Ifc() interface{}       { return x.IntsInts() }
func (x IntsIntsFrom) IntsInts() *ints.Ints   { return x.X.IntsInts().From(x.I0.UntUnt()) }
func (x IntsIntsTo) Act()                     { x.IntsInts() }
func (x IntsIntsTo) Ifc() interface{}         { return x.IntsInts() }
func (x IntsIntsTo) IntsInts() *ints.Ints     { return x.X.IntsInts().To(x.I0.UntUnt()) }
func (x IntsIntsFst) Act()                    { x.IntInt() }
func (x IntsIntsFst) Ifc() interface{}        { return x.IntInt() }
func (x IntsIntsFst) IntInt() int.Int         { return x.X.IntsInts().Fst() }
func (x IntsIntsMdl) Act()                    { x.IntInt() }
func (x IntsIntsMdl) Ifc() interface{}        { return x.IntInt() }
func (x IntsIntsMdl) IntInt() int.Int         { return x.X.IntsInts().Mdl() }
func (x IntsIntsLst) Act()                    { x.IntInt() }
func (x IntsIntsLst) Ifc() interface{}        { return x.IntInt() }
func (x IntsIntsLst) IntInt() int.Int         { return x.X.IntsInts().Lst() }
func (x IntsIntsFstIdx) Act()                 { x.UntUnt() }
func (x IntsIntsFstIdx) Ifc() interface{}     { return x.UntUnt() }
func (x IntsIntsFstIdx) UntUnt() unt.Unt      { return x.X.IntsInts().FstIdx() }
func (x IntsIntsMdlIdx) Act()                 { x.UntUnt() }
func (x IntsIntsMdlIdx) Ifc() interface{}     { return x.UntUnt() }
func (x IntsIntsMdlIdx) UntUnt() unt.Unt      { return x.X.IntsInts().MdlIdx() }
func (x IntsIntsLstIdx) Act()                 { x.UntUnt() }
func (x IntsIntsLstIdx) Ifc() interface{}     { return x.UntUnt() }
func (x IntsIntsLstIdx) UntUnt() unt.Unt      { return x.X.IntsInts().LstIdx() }
func (x IntsIntsRev) Act()                    { x.IntsInts() }
func (x IntsIntsRev) Ifc() interface{}        { return x.IntsInts() }
func (x IntsIntsRev) IntsInts() *ints.Ints    { return x.X.IntsInts().Rev() }
func (x IntsIntsSrchIdxEql) Act()             { x.UntUnt() }
func (x IntsIntsSrchIdxEql) Ifc() interface{} { return x.UntUnt() }
func (x IntsIntsSrchIdxEql) UntUnt() unt.Unt  { return x.X.IntsInts().SrchIdxEql(x.I0.IntInt()) }
func (x IntsIntsSrchIdx) Act()                { x.UntUnt() }
func (x IntsIntsSrchIdx) Ifc() interface{}    { return x.UntUnt() }
func (x IntsIntsSrchIdx) UntUnt() unt.Unt {
	var i1 []bol.Bol
	for _, cur := range x.I1 {
		i1 = append(i1, cur.BolBol())
	}
	return x.X.IntsInts().SrchIdx(x.I0.IntInt(), i1...)
}
func (x IntsIntsHas) Act()                    { x.BolBol() }
func (x IntsIntsHas) Ifc() interface{}        { return x.BolBol() }
func (x IntsIntsHas) BolBol() bol.Bol         { return x.X.IntsInts().Has(x.I0.IntInt()) }
func (x IntsIntsSrtAsc) Act()                 { x.IntsInts() }
func (x IntsIntsSrtAsc) Ifc() interface{}     { return x.IntsInts() }
func (x IntsIntsSrtAsc) IntsInts() *ints.Ints { return x.X.IntsInts().SrtAsc() }
func (x IntsIntsSrtDsc) Act()                 { x.IntsInts() }
func (x IntsIntsSrtDsc) Ifc() interface{}     { return x.IntsInts() }
func (x IntsIntsSrtDsc) IntsInts() *ints.Ints { return x.X.IntsInts().SrtDsc() }
func (x TmesTmesBnd) Act()                    { x.BndBnd() }
func (x TmesTmesBnd) Ifc() interface{}        { return x.BndBnd() }
func (x TmesTmesBnd) BndBnd() bnd.Bnd         { return x.X.TmesTmes().Bnd(x.I0.TmeRng()) }
func (x TmesTmesWeekdayCnt) Act()             { x.UntUnt() }
func (x TmesTmesWeekdayCnt) Ifc() interface{} { return x.UntUnt() }
func (x TmesTmesWeekdayCnt) UntUnt() unt.Unt  { return x.X.TmesTmes().WeekdayCnt() }
func (x TmesTmesCnt) Act()                    { x.UntUnt() }
func (x TmesTmesCnt) Ifc() interface{}        { return x.UntUnt() }
func (x TmesTmesCnt) UntUnt() unt.Unt         { return x.X.TmesTmes().Cnt() }
func (x TmesTmesCpy) Act()                    { x.TmesTmes() }
func (x TmesTmesCpy) Ifc() interface{}        { return x.TmesTmes() }
func (x TmesTmesCpy) TmesTmes() *tmes.Tmes    { return x.X.TmesTmes().Cpy() }
func (x TmesTmesClr) Act()                    { x.TmesTmes() }
func (x TmesTmesClr) Ifc() interface{}        { return x.TmesTmes() }
func (x TmesTmesClr) TmesTmes() *tmes.Tmes    { return x.X.TmesTmes().Clr() }
func (x TmesTmesRand) Act()                   { x.TmesTmes() }
func (x TmesTmesRand) Ifc() interface{}       { return x.TmesTmes() }
func (x TmesTmesRand) TmesTmes() *tmes.Tmes   { return x.X.TmesTmes().Rand() }
func (x TmesTmesMrg) Act()                    { x.TmesTmes() }
func (x TmesTmesMrg) Ifc() interface{}        { return x.TmesTmes() }
func (x TmesTmesMrg) TmesTmes() *tmes.Tmes {
	var i0 []*tmes.Tmes
	for _, cur := range x.I0 {
		i0 = append(i0, cur.TmesTmes())
	}
	return x.X.TmesTmes().Mrg(i0...)
}
func (x TmesTmesPush) Act()             { x.TmesTmes() }
func (x TmesTmesPush) Ifc() interface{} { return x.TmesTmes() }
func (x TmesTmesPush) TmesTmes() *tmes.Tmes {
	var i0 []tme.Tme
	for _, cur := range x.I0 {
		i0 = append(i0, cur.TmeTme())
	}
	return x.X.TmesTmes().Push(i0...)
}
func (x TmesTmesPop) Act()             { x.TmeTme() }
func (x TmesTmesPop) Ifc() interface{} { return x.TmeTme() }
func (x TmesTmesPop) TmeTme() tme.Tme  { return x.X.TmesTmes().Pop() }
func (x TmesTmesQue) Act()             { x.TmesTmes() }
func (x TmesTmesQue) Ifc() interface{} { return x.TmesTmes() }
func (x TmesTmesQue) TmesTmes() *tmes.Tmes {
	var i0 []tme.Tme
	for _, cur := range x.I0 {
		i0 = append(i0, cur.TmeTme())
	}
	return x.X.TmesTmes().Que(i0...)
}
func (x TmesTmesDque) Act()                   { x.TmeTme() }
func (x TmesTmesDque) Ifc() interface{}       { return x.TmeTme() }
func (x TmesTmesDque) TmeTme() tme.Tme        { return x.X.TmesTmes().Dque() }
func (x TmesTmesIns) Act()                    { x.TmesTmes() }
func (x TmesTmesIns) Ifc() interface{}        { return x.TmesTmes() }
func (x TmesTmesIns) TmesTmes() *tmes.Tmes    { return x.X.TmesTmes().Ins(x.I0.UntUnt(), x.I1.TmeTme()) }
func (x TmesTmesUpd) Act()                    { x.TmesTmes() }
func (x TmesTmesUpd) Ifc() interface{}        { return x.TmesTmes() }
func (x TmesTmesUpd) TmesTmes() *tmes.Tmes    { return x.X.TmesTmes().Upd(x.I0.UntUnt(), x.I1.TmeTme()) }
func (x TmesTmesDel) Act()                    { x.TmeTme() }
func (x TmesTmesDel) Ifc() interface{}        { return x.TmeTme() }
func (x TmesTmesDel) TmeTme() tme.Tme         { return x.X.TmesTmes().Del(x.I0.UntUnt()) }
func (x TmesTmesAt) Act()                     { x.TmeTme() }
func (x TmesTmesAt) Ifc() interface{}         { return x.TmeTme() }
func (x TmesTmesAt) TmeTme() tme.Tme          { return x.X.TmesTmes().At(x.I0.UntUnt()) }
func (x TmesTmesIn) Act()                     { x.TmesTmes() }
func (x TmesTmesIn) Ifc() interface{}         { return x.TmesTmes() }
func (x TmesTmesIn) TmesTmes() *tmes.Tmes     { return x.X.TmesTmes().In(x.I0.UntUnt(), x.I1.UntUnt()) }
func (x TmesTmesInBnd) Act()                  { x.TmesTmes() }
func (x TmesTmesInBnd) Ifc() interface{}      { return x.TmesTmes() }
func (x TmesTmesInBnd) TmesTmes() *tmes.Tmes  { return x.X.TmesTmes().InBnd(x.I0.BndBnd()) }
func (x TmesTmesFrom) Act()                   { x.TmesTmes() }
func (x TmesTmesFrom) Ifc() interface{}       { return x.TmesTmes() }
func (x TmesTmesFrom) TmesTmes() *tmes.Tmes   { return x.X.TmesTmes().From(x.I0.UntUnt()) }
func (x TmesTmesTo) Act()                     { x.TmesTmes() }
func (x TmesTmesTo) Ifc() interface{}         { return x.TmesTmes() }
func (x TmesTmesTo) TmesTmes() *tmes.Tmes     { return x.X.TmesTmes().To(x.I0.UntUnt()) }
func (x TmesTmesFst) Act()                    { x.TmeTme() }
func (x TmesTmesFst) Ifc() interface{}        { return x.TmeTme() }
func (x TmesTmesFst) TmeTme() tme.Tme         { return x.X.TmesTmes().Fst() }
func (x TmesTmesMdl) Act()                    { x.TmeTme() }
func (x TmesTmesMdl) Ifc() interface{}        { return x.TmeTme() }
func (x TmesTmesMdl) TmeTme() tme.Tme         { return x.X.TmesTmes().Mdl() }
func (x TmesTmesLst) Act()                    { x.TmeTme() }
func (x TmesTmesLst) Ifc() interface{}        { return x.TmeTme() }
func (x TmesTmesLst) TmeTme() tme.Tme         { return x.X.TmesTmes().Lst() }
func (x TmesTmesFstIdx) Act()                 { x.UntUnt() }
func (x TmesTmesFstIdx) Ifc() interface{}     { return x.UntUnt() }
func (x TmesTmesFstIdx) UntUnt() unt.Unt      { return x.X.TmesTmes().FstIdx() }
func (x TmesTmesMdlIdx) Act()                 { x.UntUnt() }
func (x TmesTmesMdlIdx) Ifc() interface{}     { return x.UntUnt() }
func (x TmesTmesMdlIdx) UntUnt() unt.Unt      { return x.X.TmesTmes().MdlIdx() }
func (x TmesTmesLstIdx) Act()                 { x.UntUnt() }
func (x TmesTmesLstIdx) Ifc() interface{}     { return x.UntUnt() }
func (x TmesTmesLstIdx) UntUnt() unt.Unt      { return x.X.TmesTmes().LstIdx() }
func (x TmesTmesRev) Act()                    { x.TmesTmes() }
func (x TmesTmesRev) Ifc() interface{}        { return x.TmesTmes() }
func (x TmesTmesRev) TmesTmes() *tmes.Tmes    { return x.X.TmesTmes().Rev() }
func (x TmesTmesSrchIdxEql) Act()             { x.UntUnt() }
func (x TmesTmesSrchIdxEql) Ifc() interface{} { return x.UntUnt() }
func (x TmesTmesSrchIdxEql) UntUnt() unt.Unt  { return x.X.TmesTmes().SrchIdxEql(x.I0.TmeTme()) }
func (x TmesTmesSrchIdx) Act()                { x.UntUnt() }
func (x TmesTmesSrchIdx) Ifc() interface{}    { return x.UntUnt() }
func (x TmesTmesSrchIdx) UntUnt() unt.Unt {
	var i1 []bol.Bol
	for _, cur := range x.I1 {
		i1 = append(i1, cur.BolBol())
	}
	return x.X.TmesTmes().SrchIdx(x.I0.TmeTme(), i1...)
}
func (x TmesTmesHas) Act()                         { x.BolBol() }
func (x TmesTmesHas) Ifc() interface{}             { return x.BolBol() }
func (x TmesTmesHas) BolBol() bol.Bol              { return x.X.TmesTmes().Has(x.I0.TmeTme()) }
func (x TmesTmesSrtAsc) Act()                      { x.TmesTmes() }
func (x TmesTmesSrtAsc) Ifc() interface{}          { return x.TmesTmes() }
func (x TmesTmesSrtAsc) TmesTmes() *tmes.Tmes      { return x.X.TmesTmes().SrtAsc() }
func (x TmesTmesSrtDsc) Act()                      { x.TmesTmes() }
func (x TmesTmesSrtDsc) Ifc() interface{}          { return x.TmesTmes() }
func (x TmesTmesSrtDsc) TmesTmes() *tmes.Tmes      { return x.X.TmesTmes().SrtDsc() }
func (x TmesTmesInrAdd) Act()                      { x.TmesTmes() }
func (x TmesTmesInrAdd) Ifc() interface{}          { return x.TmesTmes() }
func (x TmesTmesInrAdd) TmesTmes() *tmes.Tmes      { return x.X.TmesTmes().InrAdd(x.I0.UntUnt()) }
func (x TmesTmesInrSub) Act()                      { x.TmesTmes() }
func (x TmesTmesInrSub) Ifc() interface{}          { return x.TmesTmes() }
func (x TmesTmesInrSub) TmesTmes() *tmes.Tmes      { return x.X.TmesTmes().InrSub(x.I0.UntUnt()) }
func (x TmesTmesInrMul) Act()                      { x.TmesTmes() }
func (x TmesTmesInrMul) Ifc() interface{}          { return x.TmesTmes() }
func (x TmesTmesInrMul) TmesTmes() *tmes.Tmes      { return x.X.TmesTmes().InrMul(x.I0.UntUnt()) }
func (x TmesTmesInrDiv) Act()                      { x.TmesTmes() }
func (x TmesTmesInrDiv) Ifc() interface{}          { return x.TmesTmes() }
func (x TmesTmesInrDiv) TmesTmes() *tmes.Tmes      { return x.X.TmesTmes().InrDiv(x.I0.UntUnt()) }
func (x TmesTmesInrRem) Act()                      { x.TmesTmes() }
func (x TmesTmesInrRem) Ifc() interface{}          { return x.TmesTmes() }
func (x TmesTmesInrRem) TmesTmes() *tmes.Tmes      { return x.X.TmesTmes().InrRem(x.I0.UntUnt()) }
func (x TmesTmesInrPow) Act()                      { x.TmesTmes() }
func (x TmesTmesInrPow) Ifc() interface{}          { return x.TmesTmes() }
func (x TmesTmesInrPow) TmesTmes() *tmes.Tmes      { return x.X.TmesTmes().InrPow(x.I0.UntUnt()) }
func (x TmesTmesInrMin) Act()                      { x.TmesTmes() }
func (x TmesTmesInrMin) Ifc() interface{}          { return x.TmesTmes() }
func (x TmesTmesInrMin) TmesTmes() *tmes.Tmes      { return x.X.TmesTmes().InrMin(x.I0.UntUnt()) }
func (x TmesTmesInrMax) Act()                      { x.TmesTmes() }
func (x TmesTmesInrMax) Ifc() interface{}          { return x.TmesTmes() }
func (x TmesTmesInrMax) TmesTmes() *tmes.Tmes      { return x.X.TmesTmes().InrMax(x.I0.UntUnt()) }
func (x TmesTmesSum) Act()                         { x.TmeTme() }
func (x TmesTmesSum) Ifc() interface{}             { return x.TmeTme() }
func (x TmesTmesSum) TmeTme() tme.Tme              { return x.X.TmesTmes().Sum() }
func (x TmesTmesPrd) Act()                         { x.TmeTme() }
func (x TmesTmesPrd) Ifc() interface{}             { return x.TmeTme() }
func (x TmesTmesPrd) TmeTme() tme.Tme              { return x.X.TmesTmes().Prd() }
func (x TmesTmesMin) Act()                         { x.TmeTme() }
func (x TmesTmesMin) Ifc() interface{}             { return x.TmeTme() }
func (x TmesTmesMin) TmeTme() tme.Tme              { return x.X.TmesTmes().Min() }
func (x TmesTmesMax) Act()                         { x.TmeTme() }
func (x TmesTmesMax) Ifc() interface{}             { return x.TmeTme() }
func (x TmesTmesMax) TmeTme() tme.Tme              { return x.X.TmesTmes().Max() }
func (x TmesTmesMid) Act()                         { x.TmeTme() }
func (x TmesTmesMid) Ifc() interface{}             { return x.TmeTme() }
func (x TmesTmesMid) TmeTme() tme.Tme              { return x.X.TmesTmes().Mid() }
func (x TmesTmesMdn) Act()                         { x.TmeTme() }
func (x TmesTmesMdn) Ifc() interface{}             { return x.TmeTme() }
func (x TmesTmesMdn) TmeTme() tme.Tme              { return x.X.TmesTmes().Mdn() }
func (x TmesTmesSma) Act()                         { x.TmeTme() }
func (x TmesTmesSma) Ifc() interface{}             { return x.TmeTme() }
func (x TmesTmesSma) TmeTme() tme.Tme              { return x.X.TmesTmes().Sma() }
func (x TmesTmesGma) Act()                         { x.TmeTme() }
func (x TmesTmesGma) Ifc() interface{}             { return x.TmeTme() }
func (x TmesTmesGma) TmeTme() tme.Tme              { return x.X.TmesTmes().Gma() }
func (x TmesTmesWma) Act()                         { x.TmeTme() }
func (x TmesTmesWma) Ifc() interface{}             { return x.TmeTme() }
func (x TmesTmesWma) TmeTme() tme.Tme              { return x.X.TmesTmes().Wma() }
func (x TmesTmesVrnc) Act()                        { x.TmeTme() }
func (x TmesTmesVrnc) Ifc() interface{}            { return x.TmeTme() }
func (x TmesTmesVrnc) TmeTme() tme.Tme             { return x.X.TmesTmes().Vrnc() }
func (x TmesTmesStd) Act()                         { x.TmeTme() }
func (x TmesTmesStd) Ifc() interface{}             { return x.TmeTme() }
func (x TmesTmesStd) TmeTme() tme.Tme              { return x.X.TmesTmes().Std() }
func (x TmesTmesZscr) Act()                        { x.TmesTmes() }
func (x TmesTmesZscr) Ifc() interface{}            { return x.TmesTmes() }
func (x TmesTmesZscr) TmesTmes() *tmes.Tmes        { return x.X.TmesTmes().Zscr() }
func (x TmesTmesZscrInplace) Act()                 { x.TmesTmes() }
func (x TmesTmesZscrInplace) Ifc() interface{}     { return x.TmesTmes() }
func (x TmesTmesZscrInplace) TmesTmes() *tmes.Tmes { return x.X.TmesTmes().ZscrInplace() }
func (x TmesTmesRngFul) Act()                      { x.TmeTme() }
func (x TmesTmesRngFul) Ifc() interface{}          { return x.TmeTme() }
func (x TmesTmesRngFul) TmeTme() tme.Tme           { return x.X.TmesTmes().RngFul() }
func (x TmesTmesRngLst) Act()                      { x.TmeTme() }
func (x TmesTmesRngLst) Ifc() interface{}          { return x.TmeTme() }
func (x TmesTmesRngLst) TmeTme() tme.Tme           { return x.X.TmesTmes().RngLst() }
func (x TmesTmesProLst) Act()                      { x.TmeTme() }
func (x TmesTmesProLst) Ifc() interface{}          { return x.TmeTme() }
func (x TmesTmesProLst) TmeTme() tme.Tme           { return x.X.TmesTmes().ProLst() }
func (x TmesTmesProSma) Act()                      { x.TmeTme() }
func (x TmesTmesProSma) Ifc() interface{}          { return x.TmeTme() }
func (x TmesTmesProSma) TmeTme() tme.Tme           { return x.X.TmesTmes().ProSma() }
func (x BndsBndsCnt) Act()                         { x.UntUnt() }
func (x BndsBndsCnt) Ifc() interface{}             { return x.UntUnt() }
func (x BndsBndsCnt) UntUnt() unt.Unt              { return x.X.BndsBnds().Cnt() }
func (x BndsBndsCpy) Act()                         { x.BndsBnds() }
func (x BndsBndsCpy) Ifc() interface{}             { return x.BndsBnds() }
func (x BndsBndsCpy) BndsBnds() *bnds.Bnds         { return x.X.BndsBnds().Cpy() }
func (x BndsBndsClr) Act()                         { x.BndsBnds() }
func (x BndsBndsClr) Ifc() interface{}             { return x.BndsBnds() }
func (x BndsBndsClr) BndsBnds() *bnds.Bnds         { return x.X.BndsBnds().Clr() }
func (x BndsBndsRand) Act()                        { x.BndsBnds() }
func (x BndsBndsRand) Ifc() interface{}            { return x.BndsBnds() }
func (x BndsBndsRand) BndsBnds() *bnds.Bnds        { return x.X.BndsBnds().Rand() }
func (x BndsBndsMrg) Act()                         { x.BndsBnds() }
func (x BndsBndsMrg) Ifc() interface{}             { return x.BndsBnds() }
func (x BndsBndsMrg) BndsBnds() *bnds.Bnds {
	var i0 []*bnds.Bnds
	for _, cur := range x.I0 {
		i0 = append(i0, cur.BndsBnds())
	}
	return x.X.BndsBnds().Mrg(i0...)
}
func (x BndsBndsPush) Act()             { x.BndsBnds() }
func (x BndsBndsPush) Ifc() interface{} { return x.BndsBnds() }
func (x BndsBndsPush) BndsBnds() *bnds.Bnds {
	var i0 []bnd.Bnd
	for _, cur := range x.I0 {
		i0 = append(i0, cur.BndBnd())
	}
	return x.X.BndsBnds().Push(i0...)
}
func (x BndsBndsPop) Act()             { x.BndBnd() }
func (x BndsBndsPop) Ifc() interface{} { return x.BndBnd() }
func (x BndsBndsPop) BndBnd() bnd.Bnd  { return x.X.BndsBnds().Pop() }
func (x BndsBndsQue) Act()             { x.BndsBnds() }
func (x BndsBndsQue) Ifc() interface{} { return x.BndsBnds() }
func (x BndsBndsQue) BndsBnds() *bnds.Bnds {
	var i0 []bnd.Bnd
	for _, cur := range x.I0 {
		i0 = append(i0, cur.BndBnd())
	}
	return x.X.BndsBnds().Que(i0...)
}
func (x BndsBndsDque) Act()                  { x.BndBnd() }
func (x BndsBndsDque) Ifc() interface{}      { return x.BndBnd() }
func (x BndsBndsDque) BndBnd() bnd.Bnd       { return x.X.BndsBnds().Dque() }
func (x BndsBndsIns) Act()                   { x.BndsBnds() }
func (x BndsBndsIns) Ifc() interface{}       { return x.BndsBnds() }
func (x BndsBndsIns) BndsBnds() *bnds.Bnds   { return x.X.BndsBnds().Ins(x.I0.UntUnt(), x.I1.BndBnd()) }
func (x BndsBndsUpd) Act()                   { x.BndsBnds() }
func (x BndsBndsUpd) Ifc() interface{}       { return x.BndsBnds() }
func (x BndsBndsUpd) BndsBnds() *bnds.Bnds   { return x.X.BndsBnds().Upd(x.I0.UntUnt(), x.I1.BndBnd()) }
func (x BndsBndsDel) Act()                   { x.BndBnd() }
func (x BndsBndsDel) Ifc() interface{}       { return x.BndBnd() }
func (x BndsBndsDel) BndBnd() bnd.Bnd        { return x.X.BndsBnds().Del(x.I0.UntUnt()) }
func (x BndsBndsAt) Act()                    { x.BndBnd() }
func (x BndsBndsAt) Ifc() interface{}        { return x.BndBnd() }
func (x BndsBndsAt) BndBnd() bnd.Bnd         { return x.X.BndsBnds().At(x.I0.UntUnt()) }
func (x BndsBndsIn) Act()                    { x.BndsBnds() }
func (x BndsBndsIn) Ifc() interface{}        { return x.BndsBnds() }
func (x BndsBndsIn) BndsBnds() *bnds.Bnds    { return x.X.BndsBnds().In(x.I0.UntUnt(), x.I1.UntUnt()) }
func (x BndsBndsInBnd) Act()                 { x.BndsBnds() }
func (x BndsBndsInBnd) Ifc() interface{}     { return x.BndsBnds() }
func (x BndsBndsInBnd) BndsBnds() *bnds.Bnds { return x.X.BndsBnds().InBnd(x.I0.BndBnd()) }
func (x BndsBndsFrom) Act()                  { x.BndsBnds() }
func (x BndsBndsFrom) Ifc() interface{}      { return x.BndsBnds() }
func (x BndsBndsFrom) BndsBnds() *bnds.Bnds  { return x.X.BndsBnds().From(x.I0.UntUnt()) }
func (x BndsBndsTo) Act()                    { x.BndsBnds() }
func (x BndsBndsTo) Ifc() interface{}        { return x.BndsBnds() }
func (x BndsBndsTo) BndsBnds() *bnds.Bnds    { return x.X.BndsBnds().To(x.I0.UntUnt()) }
func (x BndsBndsFst) Act()                   { x.BndBnd() }
func (x BndsBndsFst) Ifc() interface{}       { return x.BndBnd() }
func (x BndsBndsFst) BndBnd() bnd.Bnd        { return x.X.BndsBnds().Fst() }
func (x BndsBndsMdl) Act()                   { x.BndBnd() }
func (x BndsBndsMdl) Ifc() interface{}       { return x.BndBnd() }
func (x BndsBndsMdl) BndBnd() bnd.Bnd        { return x.X.BndsBnds().Mdl() }
func (x BndsBndsLst) Act()                   { x.BndBnd() }
func (x BndsBndsLst) Ifc() interface{}       { return x.BndBnd() }
func (x BndsBndsLst) BndBnd() bnd.Bnd        { return x.X.BndsBnds().Lst() }
func (x BndsBndsFstIdx) Act()                { x.UntUnt() }
func (x BndsBndsFstIdx) Ifc() interface{}    { return x.UntUnt() }
func (x BndsBndsFstIdx) UntUnt() unt.Unt     { return x.X.BndsBnds().FstIdx() }
func (x BndsBndsMdlIdx) Act()                { x.UntUnt() }
func (x BndsBndsMdlIdx) Ifc() interface{}    { return x.UntUnt() }
func (x BndsBndsMdlIdx) UntUnt() unt.Unt     { return x.X.BndsBnds().MdlIdx() }
func (x BndsBndsLstIdx) Act()                { x.UntUnt() }
func (x BndsBndsLstIdx) Ifc() interface{}    { return x.UntUnt() }
func (x BndsBndsLstIdx) UntUnt() unt.Unt     { return x.X.BndsBnds().LstIdx() }
func (x BndsBndsRev) Act()                   { x.BndsBnds() }
func (x BndsBndsRev) Ifc() interface{}       { return x.BndsBnds() }
func (x BndsBndsRev) BndsBnds() *bnds.Bnds   { return x.X.BndsBnds().Rev() }
func (x TmeRngsCnt) Act()                    { x.UntUnt() }
func (x TmeRngsCnt) Ifc() interface{}        { return x.UntUnt() }
func (x TmeRngsCnt) UntUnt() unt.Unt         { return x.X.TmeRngs().Cnt() }
func (x TmeRngsCpy) Act()                    { x.TmeRngs() }
func (x TmeRngsCpy) Ifc() interface{}        { return x.TmeRngs() }
func (x TmeRngsCpy) TmeRngs() *tme.Rngs      { return x.X.TmeRngs().Cpy() }
func (x TmeRngsClr) Act()                    { x.TmeRngs() }
func (x TmeRngsClr) Ifc() interface{}        { return x.TmeRngs() }
func (x TmeRngsClr) TmeRngs() *tme.Rngs      { return x.X.TmeRngs().Clr() }
func (x TmeRngsRand) Act()                   { x.TmeRngs() }
func (x TmeRngsRand) Ifc() interface{}       { return x.TmeRngs() }
func (x TmeRngsRand) TmeRngs() *tme.Rngs     { return x.X.TmeRngs().Rand() }
func (x TmeRngsMrg) Act()                    { x.TmeRngs() }
func (x TmeRngsMrg) Ifc() interface{}        { return x.TmeRngs() }
func (x TmeRngsMrg) TmeRngs() *tme.Rngs {
	var i0 []*tme.Rngs
	for _, cur := range x.I0 {
		i0 = append(i0, cur.TmeRngs())
	}
	return x.X.TmeRngs().Mrg(i0...)
}
func (x TmeRngsPush) Act()             { x.TmeRngs() }
func (x TmeRngsPush) Ifc() interface{} { return x.TmeRngs() }
func (x TmeRngsPush) TmeRngs() *tme.Rngs {
	var i0 []tme.Rng
	for _, cur := range x.I0 {
		i0 = append(i0, cur.TmeRng())
	}
	return x.X.TmeRngs().Push(i0...)
}
func (x TmeRngsPop) Act()             { x.TmeRng() }
func (x TmeRngsPop) Ifc() interface{} { return x.TmeRng() }
func (x TmeRngsPop) TmeRng() tme.Rng  { return x.X.TmeRngs().Pop() }
func (x TmeRngsQue) Act()             { x.TmeRngs() }
func (x TmeRngsQue) Ifc() interface{} { return x.TmeRngs() }
func (x TmeRngsQue) TmeRngs() *tme.Rngs {
	var i0 []tme.Rng
	for _, cur := range x.I0 {
		i0 = append(i0, cur.TmeRng())
	}
	return x.X.TmeRngs().Que(i0...)
}
func (x TmeRngsDque) Act()                { x.TmeRng() }
func (x TmeRngsDque) Ifc() interface{}    { return x.TmeRng() }
func (x TmeRngsDque) TmeRng() tme.Rng     { return x.X.TmeRngs().Dque() }
func (x TmeRngsIns) Act()                 { x.TmeRngs() }
func (x TmeRngsIns) Ifc() interface{}     { return x.TmeRngs() }
func (x TmeRngsIns) TmeRngs() *tme.Rngs   { return x.X.TmeRngs().Ins(x.I0.UntUnt(), x.I1.TmeRng()) }
func (x TmeRngsUpd) Act()                 { x.TmeRngs() }
func (x TmeRngsUpd) Ifc() interface{}     { return x.TmeRngs() }
func (x TmeRngsUpd) TmeRngs() *tme.Rngs   { return x.X.TmeRngs().Upd(x.I0.UntUnt(), x.I1.TmeRng()) }
func (x TmeRngsDel) Act()                 { x.TmeRng() }
func (x TmeRngsDel) Ifc() interface{}     { return x.TmeRng() }
func (x TmeRngsDel) TmeRng() tme.Rng      { return x.X.TmeRngs().Del(x.I0.UntUnt()) }
func (x TmeRngsAt) Act()                  { x.TmeRng() }
func (x TmeRngsAt) Ifc() interface{}      { return x.TmeRng() }
func (x TmeRngsAt) TmeRng() tme.Rng       { return x.X.TmeRngs().At(x.I0.UntUnt()) }
func (x TmeRngsIn) Act()                  { x.TmeRngs() }
func (x TmeRngsIn) Ifc() interface{}      { return x.TmeRngs() }
func (x TmeRngsIn) TmeRngs() *tme.Rngs    { return x.X.TmeRngs().In(x.I0.UntUnt(), x.I1.UntUnt()) }
func (x TmeRngsInBnd) Act()               { x.TmeRngs() }
func (x TmeRngsInBnd) Ifc() interface{}   { return x.TmeRngs() }
func (x TmeRngsInBnd) TmeRngs() *tme.Rngs { return x.X.TmeRngs().InBnd(x.I0.BndBnd()) }
func (x TmeRngsFrom) Act()                { x.TmeRngs() }
func (x TmeRngsFrom) Ifc() interface{}    { return x.TmeRngs() }
func (x TmeRngsFrom) TmeRngs() *tme.Rngs  { return x.X.TmeRngs().From(x.I0.UntUnt()) }
func (x TmeRngsTo) Act()                  { x.TmeRngs() }
func (x TmeRngsTo) Ifc() interface{}      { return x.TmeRngs() }
func (x TmeRngsTo) TmeRngs() *tme.Rngs    { return x.X.TmeRngs().To(x.I0.UntUnt()) }
func (x TmeRngsFst) Act()                 { x.TmeRng() }
func (x TmeRngsFst) Ifc() interface{}     { return x.TmeRng() }
func (x TmeRngsFst) TmeRng() tme.Rng      { return x.X.TmeRngs().Fst() }
func (x TmeRngsMdl) Act()                 { x.TmeRng() }
func (x TmeRngsMdl) Ifc() interface{}     { return x.TmeRng() }
func (x TmeRngsMdl) TmeRng() tme.Rng      { return x.X.TmeRngs().Mdl() }
func (x TmeRngsLst) Act()                 { x.TmeRng() }
func (x TmeRngsLst) Ifc() interface{}     { return x.TmeRng() }
func (x TmeRngsLst) TmeRng() tme.Rng      { return x.X.TmeRngs().Lst() }
func (x TmeRngsFstIdx) Act()              { x.UntUnt() }
func (x TmeRngsFstIdx) Ifc() interface{}  { return x.UntUnt() }
func (x TmeRngsFstIdx) UntUnt() unt.Unt   { return x.X.TmeRngs().FstIdx() }
func (x TmeRngsMdlIdx) Act()              { x.UntUnt() }
func (x TmeRngsMdlIdx) Ifc() interface{}  { return x.UntUnt() }
func (x TmeRngsMdlIdx) UntUnt() unt.Unt   { return x.X.TmeRngs().MdlIdx() }
func (x TmeRngsLstIdx) Act()              { x.UntUnt() }
func (x TmeRngsLstIdx) Ifc() interface{}  { return x.UntUnt() }
func (x TmeRngsLstIdx) UntUnt() unt.Unt   { return x.X.TmeRngs().LstIdx() }
func (x TmeRngsRev) Act()                 { x.TmeRngs() }
func (x TmeRngsRev) Ifc() interface{}     { return x.TmeRngs() }
func (x TmeRngsRev) TmeRngs() *tme.Rngs   { return x.X.TmeRngs().Rev() }
func (x TmeRngsSrchIdx) Act()             { x.UntUnt() }
func (x TmeRngsSrchIdx) Ifc() interface{} { return x.UntUnt() }
func (x TmeRngsSrchIdx) UntUnt() unt.Unt  { return x.X.TmeRngs().SrchIdx(x.I0.TmeTme()) }
func (x TmeRngsRngMrg) Act()              { x.TmeRng() }
func (x TmeRngsRngMrg) Ifc() interface{}  { return x.TmeRng() }
func (x TmeRngsRngMrg) TmeRng() tme.Rng   { return x.X.TmeRngs().RngMrg(x.I0.UntUnt(), x.I1.UntUnt()) }
func (x AnaTrdOpnMid) Act()               { x.FltFlt() }
func (x AnaTrdOpnMid) Ifc() interface{}   { return x.FltFlt() }
func (x AnaTrdOpnMid) FltFlt() flt.Flt    { return x.X.AnaTrd().OpnMid() }
func (x AnaTrdClsMid) Act()               { x.FltFlt() }
func (x AnaTrdClsMid) Ifc() interface{}   { return x.FltFlt() }
func (x AnaTrdClsMid) FltFlt() flt.Flt    { return x.X.AnaTrd().ClsMid() }
func (x AnaTrdsCnt) Act()                 { x.UntUnt() }
func (x AnaTrdsCnt) Ifc() interface{}     { return x.UntUnt() }
func (x AnaTrdsCnt) UntUnt() unt.Unt      { return x.X.AnaTrds().Cnt() }
func (x AnaTrdsCpy) Act()                 { x.AnaTrds() }
func (x AnaTrdsCpy) Ifc() interface{}     { return x.AnaTrds() }
func (x AnaTrdsCpy) AnaTrds() *ana.Trds   { return x.X.AnaTrds().Cpy() }
func (x AnaTrdsClr) Act()                 { x.AnaTrds() }
func (x AnaTrdsClr) Ifc() interface{}     { return x.AnaTrds() }
func (x AnaTrdsClr) AnaTrds() *ana.Trds   { return x.X.AnaTrds().Clr() }
func (x AnaTrdsRand) Act()                { x.AnaTrds() }
func (x AnaTrdsRand) Ifc() interface{}    { return x.AnaTrds() }
func (x AnaTrdsRand) AnaTrds() *ana.Trds  { return x.X.AnaTrds().Rand() }
func (x AnaTrdsMrg) Act()                 { x.AnaTrds() }
func (x AnaTrdsMrg) Ifc() interface{}     { return x.AnaTrds() }
func (x AnaTrdsMrg) AnaTrds() *ana.Trds {
	var i0 []*ana.Trds
	for _, cur := range x.I0 {
		i0 = append(i0, cur.AnaTrds())
	}
	return x.X.AnaTrds().Mrg(i0...)
}
func (x AnaTrdsPush) Act()             { x.AnaTrds() }
func (x AnaTrdsPush) Ifc() interface{} { return x.AnaTrds() }
func (x AnaTrdsPush) AnaTrds() *ana.Trds {
	var i0 []*ana.Trd
	for _, cur := range x.I0 {
		i0 = append(i0, cur.AnaTrd())
	}
	return x.X.AnaTrds().Push(i0...)
}
func (x AnaTrdsPop) Act()             { x.AnaTrd() }
func (x AnaTrdsPop) Ifc() interface{} { return x.AnaTrd() }
func (x AnaTrdsPop) AnaTrd() *ana.Trd { return x.X.AnaTrds().Pop() }
func (x AnaTrdsQue) Act()             { x.AnaTrds() }
func (x AnaTrdsQue) Ifc() interface{} { return x.AnaTrds() }
func (x AnaTrdsQue) AnaTrds() *ana.Trds {
	var i0 []*ana.Trd
	for _, cur := range x.I0 {
		i0 = append(i0, cur.AnaTrd())
	}
	return x.X.AnaTrds().Que(i0...)
}
func (x AnaTrdsDque) Act()                       { x.AnaTrd() }
func (x AnaTrdsDque) Ifc() interface{}           { return x.AnaTrd() }
func (x AnaTrdsDque) AnaTrd() *ana.Trd           { return x.X.AnaTrds().Dque() }
func (x AnaTrdsIns) Act()                        { x.AnaTrds() }
func (x AnaTrdsIns) Ifc() interface{}            { return x.AnaTrds() }
func (x AnaTrdsIns) AnaTrds() *ana.Trds          { return x.X.AnaTrds().Ins(x.I0.UntUnt(), x.I1.AnaTrd()) }
func (x AnaTrdsUpd) Act()                        { x.AnaTrds() }
func (x AnaTrdsUpd) Ifc() interface{}            { return x.AnaTrds() }
func (x AnaTrdsUpd) AnaTrds() *ana.Trds          { return x.X.AnaTrds().Upd(x.I0.UntUnt(), x.I1.AnaTrd()) }
func (x AnaTrdsDel) Act()                        { x.AnaTrd() }
func (x AnaTrdsDel) Ifc() interface{}            { return x.AnaTrd() }
func (x AnaTrdsDel) AnaTrd() *ana.Trd            { return x.X.AnaTrds().Del(x.I0.UntUnt()) }
func (x AnaTrdsAt) Act()                         { x.AnaTrd() }
func (x AnaTrdsAt) Ifc() interface{}             { return x.AnaTrd() }
func (x AnaTrdsAt) AnaTrd() *ana.Trd             { return x.X.AnaTrds().At(x.I0.UntUnt()) }
func (x AnaTrdsIn) Act()                         { x.AnaTrds() }
func (x AnaTrdsIn) Ifc() interface{}             { return x.AnaTrds() }
func (x AnaTrdsIn) AnaTrds() *ana.Trds           { return x.X.AnaTrds().In(x.I0.UntUnt(), x.I1.UntUnt()) }
func (x AnaTrdsInBnd) Act()                      { x.AnaTrds() }
func (x AnaTrdsInBnd) Ifc() interface{}          { return x.AnaTrds() }
func (x AnaTrdsInBnd) AnaTrds() *ana.Trds        { return x.X.AnaTrds().InBnd(x.I0.BndBnd()) }
func (x AnaTrdsFrom) Act()                       { x.AnaTrds() }
func (x AnaTrdsFrom) Ifc() interface{}           { return x.AnaTrds() }
func (x AnaTrdsFrom) AnaTrds() *ana.Trds         { return x.X.AnaTrds().From(x.I0.UntUnt()) }
func (x AnaTrdsTo) Act()                         { x.AnaTrds() }
func (x AnaTrdsTo) Ifc() interface{}             { return x.AnaTrds() }
func (x AnaTrdsTo) AnaTrds() *ana.Trds           { return x.X.AnaTrds().To(x.I0.UntUnt()) }
func (x AnaTrdsFst) Act()                        { x.AnaTrd() }
func (x AnaTrdsFst) Ifc() interface{}            { return x.AnaTrd() }
func (x AnaTrdsFst) AnaTrd() *ana.Trd            { return x.X.AnaTrds().Fst() }
func (x AnaTrdsMdl) Act()                        { x.AnaTrd() }
func (x AnaTrdsMdl) Ifc() interface{}            { return x.AnaTrd() }
func (x AnaTrdsMdl) AnaTrd() *ana.Trd            { return x.X.AnaTrds().Mdl() }
func (x AnaTrdsLst) Act()                        { x.AnaTrd() }
func (x AnaTrdsLst) Ifc() interface{}            { return x.AnaTrd() }
func (x AnaTrdsLst) AnaTrd() *ana.Trd            { return x.X.AnaTrds().Lst() }
func (x AnaTrdsFstIdx) Act()                     { x.UntUnt() }
func (x AnaTrdsFstIdx) Ifc() interface{}         { return x.UntUnt() }
func (x AnaTrdsFstIdx) UntUnt() unt.Unt          { return x.X.AnaTrds().FstIdx() }
func (x AnaTrdsMdlIdx) Act()                     { x.UntUnt() }
func (x AnaTrdsMdlIdx) Ifc() interface{}         { return x.UntUnt() }
func (x AnaTrdsMdlIdx) UntUnt() unt.Unt          { return x.X.AnaTrds().MdlIdx() }
func (x AnaTrdsLstIdx) Act()                     { x.UntUnt() }
func (x AnaTrdsLstIdx) Ifc() interface{}         { return x.UntUnt() }
func (x AnaTrdsLstIdx) UntUnt() unt.Unt          { return x.X.AnaTrds().LstIdx() }
func (x AnaTrdsRev) Act()                        { x.AnaTrds() }
func (x AnaTrdsRev) Ifc() interface{}            { return x.AnaTrds() }
func (x AnaTrdsRev) AnaTrds() *ana.Trds          { return x.X.AnaTrds().Rev() }
func (x AnaTrdsSelClsResEql) Act()               { x.AnaTrds() }
func (x AnaTrdsSelClsResEql) Ifc() interface{}   { return x.AnaTrds() }
func (x AnaTrdsSelClsResEql) AnaTrds() *ana.Trds { return x.X.AnaTrds().SelClsResEql(x.I0.StrStr()) }
func (x AnaTrdsSelClsResNeq) Act()               { x.AnaTrds() }
func (x AnaTrdsSelClsResNeq) Ifc() interface{}   { return x.AnaTrds() }
func (x AnaTrdsSelClsResNeq) AnaTrds() *ana.Trds { return x.X.AnaTrds().SelClsResNeq(x.I0.StrStr()) }
func (x AnaTrdsSelClsResLss) Act()               { x.AnaTrds() }
func (x AnaTrdsSelClsResLss) Ifc() interface{}   { return x.AnaTrds() }
func (x AnaTrdsSelClsResLss) AnaTrds() *ana.Trds { return x.X.AnaTrds().SelClsResLss(x.I0.StrStr()) }
func (x AnaTrdsSelClsResGtr) Act()               { x.AnaTrds() }
func (x AnaTrdsSelClsResGtr) Ifc() interface{}   { return x.AnaTrds() }
func (x AnaTrdsSelClsResGtr) AnaTrds() *ana.Trds { return x.X.AnaTrds().SelClsResGtr(x.I0.StrStr()) }
func (x AnaTrdsSelClsResLeq) Act()               { x.AnaTrds() }
func (x AnaTrdsSelClsResLeq) Ifc() interface{}   { return x.AnaTrds() }
func (x AnaTrdsSelClsResLeq) AnaTrds() *ana.Trds { return x.X.AnaTrds().SelClsResLeq(x.I0.StrStr()) }
func (x AnaTrdsSelClsResGeq) Act()               { x.AnaTrds() }
func (x AnaTrdsSelClsResGeq) Ifc() interface{}   { return x.AnaTrds() }
func (x AnaTrdsSelClsResGeq) AnaTrds() *ana.Trds { return x.X.AnaTrds().SelClsResGeq(x.I0.StrStr()) }
func (x AnaTrdsSelClsReqEql) Act()               { x.AnaTrds() }
func (x AnaTrdsSelClsReqEql) Ifc() interface{}   { return x.AnaTrds() }
func (x AnaTrdsSelClsReqEql) AnaTrds() *ana.Trds { return x.X.AnaTrds().SelClsReqEql(x.I0.StrStr()) }
func (x AnaTrdsSelClsReqNeq) Act()               { x.AnaTrds() }
func (x AnaTrdsSelClsReqNeq) Ifc() interface{}   { return x.AnaTrds() }
func (x AnaTrdsSelClsReqNeq) AnaTrds() *ana.Trds { return x.X.AnaTrds().SelClsReqNeq(x.I0.StrStr()) }
func (x AnaTrdsSelClsReqLss) Act()               { x.AnaTrds() }
func (x AnaTrdsSelClsReqLss) Ifc() interface{}   { return x.AnaTrds() }
func (x AnaTrdsSelClsReqLss) AnaTrds() *ana.Trds { return x.X.AnaTrds().SelClsReqLss(x.I0.StrStr()) }
func (x AnaTrdsSelClsReqGtr) Act()               { x.AnaTrds() }
func (x AnaTrdsSelClsReqGtr) Ifc() interface{}   { return x.AnaTrds() }
func (x AnaTrdsSelClsReqGtr) AnaTrds() *ana.Trds { return x.X.AnaTrds().SelClsReqGtr(x.I0.StrStr()) }
func (x AnaTrdsSelClsReqLeq) Act()               { x.AnaTrds() }
func (x AnaTrdsSelClsReqLeq) Ifc() interface{}   { return x.AnaTrds() }
func (x AnaTrdsSelClsReqLeq) AnaTrds() *ana.Trds { return x.X.AnaTrds().SelClsReqLeq(x.I0.StrStr()) }
func (x AnaTrdsSelClsReqGeq) Act()               { x.AnaTrds() }
func (x AnaTrdsSelClsReqGeq) Ifc() interface{}   { return x.AnaTrds() }
func (x AnaTrdsSelClsReqGeq) AnaTrds() *ana.Trds { return x.X.AnaTrds().SelClsReqGeq(x.I0.StrStr()) }
func (x AnaTrdsSelOpnResEql) Act()               { x.AnaTrds() }
func (x AnaTrdsSelOpnResEql) Ifc() interface{}   { return x.AnaTrds() }
func (x AnaTrdsSelOpnResEql) AnaTrds() *ana.Trds { return x.X.AnaTrds().SelOpnResEql(x.I0.StrStr()) }
func (x AnaTrdsSelOpnResNeq) Act()               { x.AnaTrds() }
func (x AnaTrdsSelOpnResNeq) Ifc() interface{}   { return x.AnaTrds() }
func (x AnaTrdsSelOpnResNeq) AnaTrds() *ana.Trds { return x.X.AnaTrds().SelOpnResNeq(x.I0.StrStr()) }
func (x AnaTrdsSelOpnResLss) Act()               { x.AnaTrds() }
func (x AnaTrdsSelOpnResLss) Ifc() interface{}   { return x.AnaTrds() }
func (x AnaTrdsSelOpnResLss) AnaTrds() *ana.Trds { return x.X.AnaTrds().SelOpnResLss(x.I0.StrStr()) }
func (x AnaTrdsSelOpnResGtr) Act()               { x.AnaTrds() }
func (x AnaTrdsSelOpnResGtr) Ifc() interface{}   { return x.AnaTrds() }
func (x AnaTrdsSelOpnResGtr) AnaTrds() *ana.Trds { return x.X.AnaTrds().SelOpnResGtr(x.I0.StrStr()) }
func (x AnaTrdsSelOpnResLeq) Act()               { x.AnaTrds() }
func (x AnaTrdsSelOpnResLeq) Ifc() interface{}   { return x.AnaTrds() }
func (x AnaTrdsSelOpnResLeq) AnaTrds() *ana.Trds { return x.X.AnaTrds().SelOpnResLeq(x.I0.StrStr()) }
func (x AnaTrdsSelOpnResGeq) Act()               { x.AnaTrds() }
func (x AnaTrdsSelOpnResGeq) Ifc() interface{}   { return x.AnaTrds() }
func (x AnaTrdsSelOpnResGeq) AnaTrds() *ana.Trds { return x.X.AnaTrds().SelOpnResGeq(x.I0.StrStr()) }
func (x AnaTrdsSelOpnReqEql) Act()               { x.AnaTrds() }
func (x AnaTrdsSelOpnReqEql) Ifc() interface{}   { return x.AnaTrds() }
func (x AnaTrdsSelOpnReqEql) AnaTrds() *ana.Trds { return x.X.AnaTrds().SelOpnReqEql(x.I0.StrStr()) }
func (x AnaTrdsSelOpnReqNeq) Act()               { x.AnaTrds() }
func (x AnaTrdsSelOpnReqNeq) Ifc() interface{}   { return x.AnaTrds() }
func (x AnaTrdsSelOpnReqNeq) AnaTrds() *ana.Trds { return x.X.AnaTrds().SelOpnReqNeq(x.I0.StrStr()) }
func (x AnaTrdsSelOpnReqLss) Act()               { x.AnaTrds() }
func (x AnaTrdsSelOpnReqLss) Ifc() interface{}   { return x.AnaTrds() }
func (x AnaTrdsSelOpnReqLss) AnaTrds() *ana.Trds { return x.X.AnaTrds().SelOpnReqLss(x.I0.StrStr()) }
func (x AnaTrdsSelOpnReqGtr) Act()               { x.AnaTrds() }
func (x AnaTrdsSelOpnReqGtr) Ifc() interface{}   { return x.AnaTrds() }
func (x AnaTrdsSelOpnReqGtr) AnaTrds() *ana.Trds { return x.X.AnaTrds().SelOpnReqGtr(x.I0.StrStr()) }
func (x AnaTrdsSelOpnReqLeq) Act()               { x.AnaTrds() }
func (x AnaTrdsSelOpnReqLeq) Ifc() interface{}   { return x.AnaTrds() }
func (x AnaTrdsSelOpnReqLeq) AnaTrds() *ana.Trds { return x.X.AnaTrds().SelOpnReqLeq(x.I0.StrStr()) }
func (x AnaTrdsSelOpnReqGeq) Act()               { x.AnaTrds() }
func (x AnaTrdsSelOpnReqGeq) Ifc() interface{}   { return x.AnaTrds() }
func (x AnaTrdsSelOpnReqGeq) AnaTrds() *ana.Trds { return x.X.AnaTrds().SelOpnReqGeq(x.I0.StrStr()) }
func (x AnaTrdsSelInstrEql) Act()                { x.AnaTrds() }
func (x AnaTrdsSelInstrEql) Ifc() interface{}    { return x.AnaTrds() }
func (x AnaTrdsSelInstrEql) AnaTrds() *ana.Trds  { return x.X.AnaTrds().SelInstrEql(x.I0.StrStr()) }
func (x AnaTrdsSelInstrNeq) Act()                { x.AnaTrds() }
func (x AnaTrdsSelInstrNeq) Ifc() interface{}    { return x.AnaTrds() }
func (x AnaTrdsSelInstrNeq) AnaTrds() *ana.Trds  { return x.X.AnaTrds().SelInstrNeq(x.I0.StrStr()) }
func (x AnaTrdsSelInstrLss) Act()                { x.AnaTrds() }
func (x AnaTrdsSelInstrLss) Ifc() interface{}    { return x.AnaTrds() }
func (x AnaTrdsSelInstrLss) AnaTrds() *ana.Trds  { return x.X.AnaTrds().SelInstrLss(x.I0.StrStr()) }
func (x AnaTrdsSelInstrGtr) Act()                { x.AnaTrds() }
func (x AnaTrdsSelInstrGtr) Ifc() interface{}    { return x.AnaTrds() }
func (x AnaTrdsSelInstrGtr) AnaTrds() *ana.Trds  { return x.X.AnaTrds().SelInstrGtr(x.I0.StrStr()) }
func (x AnaTrdsSelInstrLeq) Act()                { x.AnaTrds() }
func (x AnaTrdsSelInstrLeq) Ifc() interface{}    { return x.AnaTrds() }
func (x AnaTrdsSelInstrLeq) AnaTrds() *ana.Trds  { return x.X.AnaTrds().SelInstrLeq(x.I0.StrStr()) }
func (x AnaTrdsSelInstrGeq) Act()                { x.AnaTrds() }
func (x AnaTrdsSelInstrGeq) Ifc() interface{}    { return x.AnaTrds() }
func (x AnaTrdsSelInstrGeq) AnaTrds() *ana.Trds  { return x.X.AnaTrds().SelInstrGeq(x.I0.StrStr()) }
func (x AnaTrdsSelUnitsEql) Act()                { x.AnaTrds() }
func (x AnaTrdsSelUnitsEql) Ifc() interface{}    { return x.AnaTrds() }
func (x AnaTrdsSelUnitsEql) AnaTrds() *ana.Trds  { return x.X.AnaTrds().SelUnitsEql(x.I0.FltFlt()) }
func (x AnaTrdsSelUnitsNeq) Act()                { x.AnaTrds() }
func (x AnaTrdsSelUnitsNeq) Ifc() interface{}    { return x.AnaTrds() }
func (x AnaTrdsSelUnitsNeq) AnaTrds() *ana.Trds  { return x.X.AnaTrds().SelUnitsNeq(x.I0.FltFlt()) }
func (x AnaTrdsSelUnitsLss) Act()                { x.AnaTrds() }
func (x AnaTrdsSelUnitsLss) Ifc() interface{}    { return x.AnaTrds() }
func (x AnaTrdsSelUnitsLss) AnaTrds() *ana.Trds  { return x.X.AnaTrds().SelUnitsLss(x.I0.FltFlt()) }
func (x AnaTrdsSelUnitsGtr) Act()                { x.AnaTrds() }
func (x AnaTrdsSelUnitsGtr) Ifc() interface{}    { return x.AnaTrds() }
func (x AnaTrdsSelUnitsGtr) AnaTrds() *ana.Trds  { return x.X.AnaTrds().SelUnitsGtr(x.I0.FltFlt()) }
func (x AnaTrdsSelUnitsLeq) Act()                { x.AnaTrds() }
func (x AnaTrdsSelUnitsLeq) Ifc() interface{}    { return x.AnaTrds() }
func (x AnaTrdsSelUnitsLeq) AnaTrds() *ana.Trds  { return x.X.AnaTrds().SelUnitsLeq(x.I0.FltFlt()) }
func (x AnaTrdsSelUnitsGeq) Act()                { x.AnaTrds() }
func (x AnaTrdsSelUnitsGeq) Ifc() interface{}    { return x.AnaTrds() }
func (x AnaTrdsSelUnitsGeq) AnaTrds() *ana.Trds  { return x.X.AnaTrds().SelUnitsGeq(x.I0.FltFlt()) }
func (x AnaTrdsSelMrgnRtioEql) Act()             { x.AnaTrds() }
func (x AnaTrdsSelMrgnRtioEql) Ifc() interface{} { return x.AnaTrds() }
func (x AnaTrdsSelMrgnRtioEql) AnaTrds() *ana.Trds {
	return x.X.AnaTrds().SelMrgnRtioEql(x.I0.FltFlt())
}
func (x AnaTrdsSelMrgnRtioNeq) Act()             { x.AnaTrds() }
func (x AnaTrdsSelMrgnRtioNeq) Ifc() interface{} { return x.AnaTrds() }
func (x AnaTrdsSelMrgnRtioNeq) AnaTrds() *ana.Trds {
	return x.X.AnaTrds().SelMrgnRtioNeq(x.I0.FltFlt())
}
func (x AnaTrdsSelMrgnRtioLss) Act()             { x.AnaTrds() }
func (x AnaTrdsSelMrgnRtioLss) Ifc() interface{} { return x.AnaTrds() }
func (x AnaTrdsSelMrgnRtioLss) AnaTrds() *ana.Trds {
	return x.X.AnaTrds().SelMrgnRtioLss(x.I0.FltFlt())
}
func (x AnaTrdsSelMrgnRtioGtr) Act()             { x.AnaTrds() }
func (x AnaTrdsSelMrgnRtioGtr) Ifc() interface{} { return x.AnaTrds() }
func (x AnaTrdsSelMrgnRtioGtr) AnaTrds() *ana.Trds {
	return x.X.AnaTrds().SelMrgnRtioGtr(x.I0.FltFlt())
}
func (x AnaTrdsSelMrgnRtioLeq) Act()             { x.AnaTrds() }
func (x AnaTrdsSelMrgnRtioLeq) Ifc() interface{} { return x.AnaTrds() }
func (x AnaTrdsSelMrgnRtioLeq) AnaTrds() *ana.Trds {
	return x.X.AnaTrds().SelMrgnRtioLeq(x.I0.FltFlt())
}
func (x AnaTrdsSelMrgnRtioGeq) Act()             { x.AnaTrds() }
func (x AnaTrdsSelMrgnRtioGeq) Ifc() interface{} { return x.AnaTrds() }
func (x AnaTrdsSelMrgnRtioGeq) AnaTrds() *ana.Trds {
	return x.X.AnaTrds().SelMrgnRtioGeq(x.I0.FltFlt())
}
func (x AnaTrdsSelTrdPctEql) Act()                   { x.AnaTrds() }
func (x AnaTrdsSelTrdPctEql) Ifc() interface{}       { return x.AnaTrds() }
func (x AnaTrdsSelTrdPctEql) AnaTrds() *ana.Trds     { return x.X.AnaTrds().SelTrdPctEql(x.I0.FltFlt()) }
func (x AnaTrdsSelTrdPctNeq) Act()                   { x.AnaTrds() }
func (x AnaTrdsSelTrdPctNeq) Ifc() interface{}       { return x.AnaTrds() }
func (x AnaTrdsSelTrdPctNeq) AnaTrds() *ana.Trds     { return x.X.AnaTrds().SelTrdPctNeq(x.I0.FltFlt()) }
func (x AnaTrdsSelTrdPctLss) Act()                   { x.AnaTrds() }
func (x AnaTrdsSelTrdPctLss) Ifc() interface{}       { return x.AnaTrds() }
func (x AnaTrdsSelTrdPctLss) AnaTrds() *ana.Trds     { return x.X.AnaTrds().SelTrdPctLss(x.I0.FltFlt()) }
func (x AnaTrdsSelTrdPctGtr) Act()                   { x.AnaTrds() }
func (x AnaTrdsSelTrdPctGtr) Ifc() interface{}       { return x.AnaTrds() }
func (x AnaTrdsSelTrdPctGtr) AnaTrds() *ana.Trds     { return x.X.AnaTrds().SelTrdPctGtr(x.I0.FltFlt()) }
func (x AnaTrdsSelTrdPctLeq) Act()                   { x.AnaTrds() }
func (x AnaTrdsSelTrdPctLeq) Ifc() interface{}       { return x.AnaTrds() }
func (x AnaTrdsSelTrdPctLeq) AnaTrds() *ana.Trds     { return x.X.AnaTrds().SelTrdPctLeq(x.I0.FltFlt()) }
func (x AnaTrdsSelTrdPctGeq) Act()                   { x.AnaTrds() }
func (x AnaTrdsSelTrdPctGeq) Ifc() interface{}       { return x.AnaTrds() }
func (x AnaTrdsSelTrdPctGeq) AnaTrds() *ana.Trds     { return x.X.AnaTrds().SelTrdPctGeq(x.I0.FltFlt()) }
func (x AnaTrdsSelClsBalUsdActEql) Act()             { x.AnaTrds() }
func (x AnaTrdsSelClsBalUsdActEql) Ifc() interface{} { return x.AnaTrds() }
func (x AnaTrdsSelClsBalUsdActEql) AnaTrds() *ana.Trds {
	return x.X.AnaTrds().SelClsBalUsdActEql(x.I0.FltFlt())
}
func (x AnaTrdsSelClsBalUsdActNeq) Act()             { x.AnaTrds() }
func (x AnaTrdsSelClsBalUsdActNeq) Ifc() interface{} { return x.AnaTrds() }
func (x AnaTrdsSelClsBalUsdActNeq) AnaTrds() *ana.Trds {
	return x.X.AnaTrds().SelClsBalUsdActNeq(x.I0.FltFlt())
}
func (x AnaTrdsSelClsBalUsdActLss) Act()             { x.AnaTrds() }
func (x AnaTrdsSelClsBalUsdActLss) Ifc() interface{} { return x.AnaTrds() }
func (x AnaTrdsSelClsBalUsdActLss) AnaTrds() *ana.Trds {
	return x.X.AnaTrds().SelClsBalUsdActLss(x.I0.FltFlt())
}
func (x AnaTrdsSelClsBalUsdActGtr) Act()             { x.AnaTrds() }
func (x AnaTrdsSelClsBalUsdActGtr) Ifc() interface{} { return x.AnaTrds() }
func (x AnaTrdsSelClsBalUsdActGtr) AnaTrds() *ana.Trds {
	return x.X.AnaTrds().SelClsBalUsdActGtr(x.I0.FltFlt())
}
func (x AnaTrdsSelClsBalUsdActLeq) Act()             { x.AnaTrds() }
func (x AnaTrdsSelClsBalUsdActLeq) Ifc() interface{} { return x.AnaTrds() }
func (x AnaTrdsSelClsBalUsdActLeq) AnaTrds() *ana.Trds {
	return x.X.AnaTrds().SelClsBalUsdActLeq(x.I0.FltFlt())
}
func (x AnaTrdsSelClsBalUsdActGeq) Act()             { x.AnaTrds() }
func (x AnaTrdsSelClsBalUsdActGeq) Ifc() interface{} { return x.AnaTrds() }
func (x AnaTrdsSelClsBalUsdActGeq) AnaTrds() *ana.Trds {
	return x.X.AnaTrds().SelClsBalUsdActGeq(x.I0.FltFlt())
}
func (x AnaTrdsSelClsBalUsdEql) Act()             { x.AnaTrds() }
func (x AnaTrdsSelClsBalUsdEql) Ifc() interface{} { return x.AnaTrds() }
func (x AnaTrdsSelClsBalUsdEql) AnaTrds() *ana.Trds {
	return x.X.AnaTrds().SelClsBalUsdEql(x.I0.FltFlt())
}
func (x AnaTrdsSelClsBalUsdNeq) Act()             { x.AnaTrds() }
func (x AnaTrdsSelClsBalUsdNeq) Ifc() interface{} { return x.AnaTrds() }
func (x AnaTrdsSelClsBalUsdNeq) AnaTrds() *ana.Trds {
	return x.X.AnaTrds().SelClsBalUsdNeq(x.I0.FltFlt())
}
func (x AnaTrdsSelClsBalUsdLss) Act()             { x.AnaTrds() }
func (x AnaTrdsSelClsBalUsdLss) Ifc() interface{} { return x.AnaTrds() }
func (x AnaTrdsSelClsBalUsdLss) AnaTrds() *ana.Trds {
	return x.X.AnaTrds().SelClsBalUsdLss(x.I0.FltFlt())
}
func (x AnaTrdsSelClsBalUsdGtr) Act()             { x.AnaTrds() }
func (x AnaTrdsSelClsBalUsdGtr) Ifc() interface{} { return x.AnaTrds() }
func (x AnaTrdsSelClsBalUsdGtr) AnaTrds() *ana.Trds {
	return x.X.AnaTrds().SelClsBalUsdGtr(x.I0.FltFlt())
}
func (x AnaTrdsSelClsBalUsdLeq) Act()             { x.AnaTrds() }
func (x AnaTrdsSelClsBalUsdLeq) Ifc() interface{} { return x.AnaTrds() }
func (x AnaTrdsSelClsBalUsdLeq) AnaTrds() *ana.Trds {
	return x.X.AnaTrds().SelClsBalUsdLeq(x.I0.FltFlt())
}
func (x AnaTrdsSelClsBalUsdGeq) Act()             { x.AnaTrds() }
func (x AnaTrdsSelClsBalUsdGeq) Ifc() interface{} { return x.AnaTrds() }
func (x AnaTrdsSelClsBalUsdGeq) AnaTrds() *ana.Trds {
	return x.X.AnaTrds().SelClsBalUsdGeq(x.I0.FltFlt())
}
func (x AnaTrdsSelOpnBalUsdEql) Act()             { x.AnaTrds() }
func (x AnaTrdsSelOpnBalUsdEql) Ifc() interface{} { return x.AnaTrds() }
func (x AnaTrdsSelOpnBalUsdEql) AnaTrds() *ana.Trds {
	return x.X.AnaTrds().SelOpnBalUsdEql(x.I0.FltFlt())
}
func (x AnaTrdsSelOpnBalUsdNeq) Act()             { x.AnaTrds() }
func (x AnaTrdsSelOpnBalUsdNeq) Ifc() interface{} { return x.AnaTrds() }
func (x AnaTrdsSelOpnBalUsdNeq) AnaTrds() *ana.Trds {
	return x.X.AnaTrds().SelOpnBalUsdNeq(x.I0.FltFlt())
}
func (x AnaTrdsSelOpnBalUsdLss) Act()             { x.AnaTrds() }
func (x AnaTrdsSelOpnBalUsdLss) Ifc() interface{} { return x.AnaTrds() }
func (x AnaTrdsSelOpnBalUsdLss) AnaTrds() *ana.Trds {
	return x.X.AnaTrds().SelOpnBalUsdLss(x.I0.FltFlt())
}
func (x AnaTrdsSelOpnBalUsdGtr) Act()             { x.AnaTrds() }
func (x AnaTrdsSelOpnBalUsdGtr) Ifc() interface{} { return x.AnaTrds() }
func (x AnaTrdsSelOpnBalUsdGtr) AnaTrds() *ana.Trds {
	return x.X.AnaTrds().SelOpnBalUsdGtr(x.I0.FltFlt())
}
func (x AnaTrdsSelOpnBalUsdLeq) Act()             { x.AnaTrds() }
func (x AnaTrdsSelOpnBalUsdLeq) Ifc() interface{} { return x.AnaTrds() }
func (x AnaTrdsSelOpnBalUsdLeq) AnaTrds() *ana.Trds {
	return x.X.AnaTrds().SelOpnBalUsdLeq(x.I0.FltFlt())
}
func (x AnaTrdsSelOpnBalUsdGeq) Act()             { x.AnaTrds() }
func (x AnaTrdsSelOpnBalUsdGeq) Ifc() interface{} { return x.AnaTrds() }
func (x AnaTrdsSelOpnBalUsdGeq) AnaTrds() *ana.Trds {
	return x.X.AnaTrds().SelOpnBalUsdGeq(x.I0.FltFlt())
}
func (x AnaTrdsSelCstOpnSpdUsdEql) Act()             { x.AnaTrds() }
func (x AnaTrdsSelCstOpnSpdUsdEql) Ifc() interface{} { return x.AnaTrds() }
func (x AnaTrdsSelCstOpnSpdUsdEql) AnaTrds() *ana.Trds {
	return x.X.AnaTrds().SelCstOpnSpdUsdEql(x.I0.FltFlt())
}
func (x AnaTrdsSelCstOpnSpdUsdNeq) Act()             { x.AnaTrds() }
func (x AnaTrdsSelCstOpnSpdUsdNeq) Ifc() interface{} { return x.AnaTrds() }
func (x AnaTrdsSelCstOpnSpdUsdNeq) AnaTrds() *ana.Trds {
	return x.X.AnaTrds().SelCstOpnSpdUsdNeq(x.I0.FltFlt())
}
func (x AnaTrdsSelCstOpnSpdUsdLss) Act()             { x.AnaTrds() }
func (x AnaTrdsSelCstOpnSpdUsdLss) Ifc() interface{} { return x.AnaTrds() }
func (x AnaTrdsSelCstOpnSpdUsdLss) AnaTrds() *ana.Trds {
	return x.X.AnaTrds().SelCstOpnSpdUsdLss(x.I0.FltFlt())
}
func (x AnaTrdsSelCstOpnSpdUsdGtr) Act()             { x.AnaTrds() }
func (x AnaTrdsSelCstOpnSpdUsdGtr) Ifc() interface{} { return x.AnaTrds() }
func (x AnaTrdsSelCstOpnSpdUsdGtr) AnaTrds() *ana.Trds {
	return x.X.AnaTrds().SelCstOpnSpdUsdGtr(x.I0.FltFlt())
}
func (x AnaTrdsSelCstOpnSpdUsdLeq) Act()             { x.AnaTrds() }
func (x AnaTrdsSelCstOpnSpdUsdLeq) Ifc() interface{} { return x.AnaTrds() }
func (x AnaTrdsSelCstOpnSpdUsdLeq) AnaTrds() *ana.Trds {
	return x.X.AnaTrds().SelCstOpnSpdUsdLeq(x.I0.FltFlt())
}
func (x AnaTrdsSelCstOpnSpdUsdGeq) Act()             { x.AnaTrds() }
func (x AnaTrdsSelCstOpnSpdUsdGeq) Ifc() interface{} { return x.AnaTrds() }
func (x AnaTrdsSelCstOpnSpdUsdGeq) AnaTrds() *ana.Trds {
	return x.X.AnaTrds().SelCstOpnSpdUsdGeq(x.I0.FltFlt())
}
func (x AnaTrdsSelCstClsSpdUsdEql) Act()             { x.AnaTrds() }
func (x AnaTrdsSelCstClsSpdUsdEql) Ifc() interface{} { return x.AnaTrds() }
func (x AnaTrdsSelCstClsSpdUsdEql) AnaTrds() *ana.Trds {
	return x.X.AnaTrds().SelCstClsSpdUsdEql(x.I0.FltFlt())
}
func (x AnaTrdsSelCstClsSpdUsdNeq) Act()             { x.AnaTrds() }
func (x AnaTrdsSelCstClsSpdUsdNeq) Ifc() interface{} { return x.AnaTrds() }
func (x AnaTrdsSelCstClsSpdUsdNeq) AnaTrds() *ana.Trds {
	return x.X.AnaTrds().SelCstClsSpdUsdNeq(x.I0.FltFlt())
}
func (x AnaTrdsSelCstClsSpdUsdLss) Act()             { x.AnaTrds() }
func (x AnaTrdsSelCstClsSpdUsdLss) Ifc() interface{} { return x.AnaTrds() }
func (x AnaTrdsSelCstClsSpdUsdLss) AnaTrds() *ana.Trds {
	return x.X.AnaTrds().SelCstClsSpdUsdLss(x.I0.FltFlt())
}
func (x AnaTrdsSelCstClsSpdUsdGtr) Act()             { x.AnaTrds() }
func (x AnaTrdsSelCstClsSpdUsdGtr) Ifc() interface{} { return x.AnaTrds() }
func (x AnaTrdsSelCstClsSpdUsdGtr) AnaTrds() *ana.Trds {
	return x.X.AnaTrds().SelCstClsSpdUsdGtr(x.I0.FltFlt())
}
func (x AnaTrdsSelCstClsSpdUsdLeq) Act()             { x.AnaTrds() }
func (x AnaTrdsSelCstClsSpdUsdLeq) Ifc() interface{} { return x.AnaTrds() }
func (x AnaTrdsSelCstClsSpdUsdLeq) AnaTrds() *ana.Trds {
	return x.X.AnaTrds().SelCstClsSpdUsdLeq(x.I0.FltFlt())
}
func (x AnaTrdsSelCstClsSpdUsdGeq) Act()             { x.AnaTrds() }
func (x AnaTrdsSelCstClsSpdUsdGeq) Ifc() interface{} { return x.AnaTrds() }
func (x AnaTrdsSelCstClsSpdUsdGeq) AnaTrds() *ana.Trds {
	return x.X.AnaTrds().SelCstClsSpdUsdGeq(x.I0.FltFlt())
}
func (x AnaTrdsSelCstComUsdEql) Act()             { x.AnaTrds() }
func (x AnaTrdsSelCstComUsdEql) Ifc() interface{} { return x.AnaTrds() }
func (x AnaTrdsSelCstComUsdEql) AnaTrds() *ana.Trds {
	return x.X.AnaTrds().SelCstComUsdEql(x.I0.FltFlt())
}
func (x AnaTrdsSelCstComUsdNeq) Act()             { x.AnaTrds() }
func (x AnaTrdsSelCstComUsdNeq) Ifc() interface{} { return x.AnaTrds() }
func (x AnaTrdsSelCstComUsdNeq) AnaTrds() *ana.Trds {
	return x.X.AnaTrds().SelCstComUsdNeq(x.I0.FltFlt())
}
func (x AnaTrdsSelCstComUsdLss) Act()             { x.AnaTrds() }
func (x AnaTrdsSelCstComUsdLss) Ifc() interface{} { return x.AnaTrds() }
func (x AnaTrdsSelCstComUsdLss) AnaTrds() *ana.Trds {
	return x.X.AnaTrds().SelCstComUsdLss(x.I0.FltFlt())
}
func (x AnaTrdsSelCstComUsdGtr) Act()             { x.AnaTrds() }
func (x AnaTrdsSelCstComUsdGtr) Ifc() interface{} { return x.AnaTrds() }
func (x AnaTrdsSelCstComUsdGtr) AnaTrds() *ana.Trds {
	return x.X.AnaTrds().SelCstComUsdGtr(x.I0.FltFlt())
}
func (x AnaTrdsSelCstComUsdLeq) Act()             { x.AnaTrds() }
func (x AnaTrdsSelCstComUsdLeq) Ifc() interface{} { return x.AnaTrds() }
func (x AnaTrdsSelCstComUsdLeq) AnaTrds() *ana.Trds {
	return x.X.AnaTrds().SelCstComUsdLeq(x.I0.FltFlt())
}
func (x AnaTrdsSelCstComUsdGeq) Act()             { x.AnaTrds() }
func (x AnaTrdsSelCstComUsdGeq) Ifc() interface{} { return x.AnaTrds() }
func (x AnaTrdsSelCstComUsdGeq) AnaTrds() *ana.Trds {
	return x.X.AnaTrds().SelCstComUsdGeq(x.I0.FltFlt())
}
func (x AnaTrdsSelPnlGrsUsdEql) Act()             { x.AnaTrds() }
func (x AnaTrdsSelPnlGrsUsdEql) Ifc() interface{} { return x.AnaTrds() }
func (x AnaTrdsSelPnlGrsUsdEql) AnaTrds() *ana.Trds {
	return x.X.AnaTrds().SelPnlGrsUsdEql(x.I0.FltFlt())
}
func (x AnaTrdsSelPnlGrsUsdNeq) Act()             { x.AnaTrds() }
func (x AnaTrdsSelPnlGrsUsdNeq) Ifc() interface{} { return x.AnaTrds() }
func (x AnaTrdsSelPnlGrsUsdNeq) AnaTrds() *ana.Trds {
	return x.X.AnaTrds().SelPnlGrsUsdNeq(x.I0.FltFlt())
}
func (x AnaTrdsSelPnlGrsUsdLss) Act()             { x.AnaTrds() }
func (x AnaTrdsSelPnlGrsUsdLss) Ifc() interface{} { return x.AnaTrds() }
func (x AnaTrdsSelPnlGrsUsdLss) AnaTrds() *ana.Trds {
	return x.X.AnaTrds().SelPnlGrsUsdLss(x.I0.FltFlt())
}
func (x AnaTrdsSelPnlGrsUsdGtr) Act()             { x.AnaTrds() }
func (x AnaTrdsSelPnlGrsUsdGtr) Ifc() interface{} { return x.AnaTrds() }
func (x AnaTrdsSelPnlGrsUsdGtr) AnaTrds() *ana.Trds {
	return x.X.AnaTrds().SelPnlGrsUsdGtr(x.I0.FltFlt())
}
func (x AnaTrdsSelPnlGrsUsdLeq) Act()             { x.AnaTrds() }
func (x AnaTrdsSelPnlGrsUsdLeq) Ifc() interface{} { return x.AnaTrds() }
func (x AnaTrdsSelPnlGrsUsdLeq) AnaTrds() *ana.Trds {
	return x.X.AnaTrds().SelPnlGrsUsdLeq(x.I0.FltFlt())
}
func (x AnaTrdsSelPnlGrsUsdGeq) Act()             { x.AnaTrds() }
func (x AnaTrdsSelPnlGrsUsdGeq) Ifc() interface{} { return x.AnaTrds() }
func (x AnaTrdsSelPnlGrsUsdGeq) AnaTrds() *ana.Trds {
	return x.X.AnaTrds().SelPnlGrsUsdGeq(x.I0.FltFlt())
}
func (x AnaTrdsSelPnlUsdEql) Act()                    { x.AnaTrds() }
func (x AnaTrdsSelPnlUsdEql) Ifc() interface{}        { return x.AnaTrds() }
func (x AnaTrdsSelPnlUsdEql) AnaTrds() *ana.Trds      { return x.X.AnaTrds().SelPnlUsdEql(x.I0.FltFlt()) }
func (x AnaTrdsSelPnlUsdNeq) Act()                    { x.AnaTrds() }
func (x AnaTrdsSelPnlUsdNeq) Ifc() interface{}        { return x.AnaTrds() }
func (x AnaTrdsSelPnlUsdNeq) AnaTrds() *ana.Trds      { return x.X.AnaTrds().SelPnlUsdNeq(x.I0.FltFlt()) }
func (x AnaTrdsSelPnlUsdLss) Act()                    { x.AnaTrds() }
func (x AnaTrdsSelPnlUsdLss) Ifc() interface{}        { return x.AnaTrds() }
func (x AnaTrdsSelPnlUsdLss) AnaTrds() *ana.Trds      { return x.X.AnaTrds().SelPnlUsdLss(x.I0.FltFlt()) }
func (x AnaTrdsSelPnlUsdGtr) Act()                    { x.AnaTrds() }
func (x AnaTrdsSelPnlUsdGtr) Ifc() interface{}        { return x.AnaTrds() }
func (x AnaTrdsSelPnlUsdGtr) AnaTrds() *ana.Trds      { return x.X.AnaTrds().SelPnlUsdGtr(x.I0.FltFlt()) }
func (x AnaTrdsSelPnlUsdLeq) Act()                    { x.AnaTrds() }
func (x AnaTrdsSelPnlUsdLeq) Ifc() interface{}        { return x.AnaTrds() }
func (x AnaTrdsSelPnlUsdLeq) AnaTrds() *ana.Trds      { return x.X.AnaTrds().SelPnlUsdLeq(x.I0.FltFlt()) }
func (x AnaTrdsSelPnlUsdGeq) Act()                    { x.AnaTrds() }
func (x AnaTrdsSelPnlUsdGeq) Ifc() interface{}        { return x.AnaTrds() }
func (x AnaTrdsSelPnlUsdGeq) AnaTrds() *ana.Trds      { return x.X.AnaTrds().SelPnlUsdGeq(x.I0.FltFlt()) }
func (x AnaTrdsSelPnlPctPredictEql) Act()             { x.AnaTrds() }
func (x AnaTrdsSelPnlPctPredictEql) Ifc() interface{} { return x.AnaTrds() }
func (x AnaTrdsSelPnlPctPredictEql) AnaTrds() *ana.Trds {
	return x.X.AnaTrds().SelPnlPctPredictEql(x.I0.FltFlt())
}
func (x AnaTrdsSelPnlPctPredictNeq) Act()             { x.AnaTrds() }
func (x AnaTrdsSelPnlPctPredictNeq) Ifc() interface{} { return x.AnaTrds() }
func (x AnaTrdsSelPnlPctPredictNeq) AnaTrds() *ana.Trds {
	return x.X.AnaTrds().SelPnlPctPredictNeq(x.I0.FltFlt())
}
func (x AnaTrdsSelPnlPctPredictLss) Act()             { x.AnaTrds() }
func (x AnaTrdsSelPnlPctPredictLss) Ifc() interface{} { return x.AnaTrds() }
func (x AnaTrdsSelPnlPctPredictLss) AnaTrds() *ana.Trds {
	return x.X.AnaTrds().SelPnlPctPredictLss(x.I0.FltFlt())
}
func (x AnaTrdsSelPnlPctPredictGtr) Act()             { x.AnaTrds() }
func (x AnaTrdsSelPnlPctPredictGtr) Ifc() interface{} { return x.AnaTrds() }
func (x AnaTrdsSelPnlPctPredictGtr) AnaTrds() *ana.Trds {
	return x.X.AnaTrds().SelPnlPctPredictGtr(x.I0.FltFlt())
}
func (x AnaTrdsSelPnlPctPredictLeq) Act()             { x.AnaTrds() }
func (x AnaTrdsSelPnlPctPredictLeq) Ifc() interface{} { return x.AnaTrds() }
func (x AnaTrdsSelPnlPctPredictLeq) AnaTrds() *ana.Trds {
	return x.X.AnaTrds().SelPnlPctPredictLeq(x.I0.FltFlt())
}
func (x AnaTrdsSelPnlPctPredictGeq) Act()             { x.AnaTrds() }
func (x AnaTrdsSelPnlPctPredictGeq) Ifc() interface{} { return x.AnaTrds() }
func (x AnaTrdsSelPnlPctPredictGeq) AnaTrds() *ana.Trds {
	return x.X.AnaTrds().SelPnlPctPredictGeq(x.I0.FltFlt())
}
func (x AnaTrdsSelPnlPctEql) Act()                   { x.AnaTrds() }
func (x AnaTrdsSelPnlPctEql) Ifc() interface{}       { return x.AnaTrds() }
func (x AnaTrdsSelPnlPctEql) AnaTrds() *ana.Trds     { return x.X.AnaTrds().SelPnlPctEql(x.I0.FltFlt()) }
func (x AnaTrdsSelPnlPctNeq) Act()                   { x.AnaTrds() }
func (x AnaTrdsSelPnlPctNeq) Ifc() interface{}       { return x.AnaTrds() }
func (x AnaTrdsSelPnlPctNeq) AnaTrds() *ana.Trds     { return x.X.AnaTrds().SelPnlPctNeq(x.I0.FltFlt()) }
func (x AnaTrdsSelPnlPctLss) Act()                   { x.AnaTrds() }
func (x AnaTrdsSelPnlPctLss) Ifc() interface{}       { return x.AnaTrds() }
func (x AnaTrdsSelPnlPctLss) AnaTrds() *ana.Trds     { return x.X.AnaTrds().SelPnlPctLss(x.I0.FltFlt()) }
func (x AnaTrdsSelPnlPctGtr) Act()                   { x.AnaTrds() }
func (x AnaTrdsSelPnlPctGtr) Ifc() interface{}       { return x.AnaTrds() }
func (x AnaTrdsSelPnlPctGtr) AnaTrds() *ana.Trds     { return x.X.AnaTrds().SelPnlPctGtr(x.I0.FltFlt()) }
func (x AnaTrdsSelPnlPctLeq) Act()                   { x.AnaTrds() }
func (x AnaTrdsSelPnlPctLeq) Ifc() interface{}       { return x.AnaTrds() }
func (x AnaTrdsSelPnlPctLeq) AnaTrds() *ana.Trds     { return x.X.AnaTrds().SelPnlPctLeq(x.I0.FltFlt()) }
func (x AnaTrdsSelPnlPctGeq) Act()                   { x.AnaTrds() }
func (x AnaTrdsSelPnlPctGeq) Ifc() interface{}       { return x.AnaTrds() }
func (x AnaTrdsSelPnlPctGeq) AnaTrds() *ana.Trds     { return x.X.AnaTrds().SelPnlPctGeq(x.I0.FltFlt()) }
func (x AnaTrdsSelIsLongEql) Act()                   { x.AnaTrds() }
func (x AnaTrdsSelIsLongEql) Ifc() interface{}       { return x.AnaTrds() }
func (x AnaTrdsSelIsLongEql) AnaTrds() *ana.Trds     { return x.X.AnaTrds().SelIsLongEql(x.I0.BolBol()) }
func (x AnaTrdsSelIsLongNeq) Act()                   { x.AnaTrds() }
func (x AnaTrdsSelIsLongNeq) Ifc() interface{}       { return x.AnaTrds() }
func (x AnaTrdsSelIsLongNeq) AnaTrds() *ana.Trds     { return x.X.AnaTrds().SelIsLongNeq(x.I0.BolBol()) }
func (x AnaTrdsSelDurEql) Act()                      { x.AnaTrds() }
func (x AnaTrdsSelDurEql) Ifc() interface{}          { return x.AnaTrds() }
func (x AnaTrdsSelDurEql) AnaTrds() *ana.Trds        { return x.X.AnaTrds().SelDurEql(x.I0.TmeTme()) }
func (x AnaTrdsSelDurNeq) Act()                      { x.AnaTrds() }
func (x AnaTrdsSelDurNeq) Ifc() interface{}          { return x.AnaTrds() }
func (x AnaTrdsSelDurNeq) AnaTrds() *ana.Trds        { return x.X.AnaTrds().SelDurNeq(x.I0.TmeTme()) }
func (x AnaTrdsSelDurLss) Act()                      { x.AnaTrds() }
func (x AnaTrdsSelDurLss) Ifc() interface{}          { return x.AnaTrds() }
func (x AnaTrdsSelDurLss) AnaTrds() *ana.Trds        { return x.X.AnaTrds().SelDurLss(x.I0.TmeTme()) }
func (x AnaTrdsSelDurGtr) Act()                      { x.AnaTrds() }
func (x AnaTrdsSelDurGtr) Ifc() interface{}          { return x.AnaTrds() }
func (x AnaTrdsSelDurGtr) AnaTrds() *ana.Trds        { return x.X.AnaTrds().SelDurGtr(x.I0.TmeTme()) }
func (x AnaTrdsSelDurLeq) Act()                      { x.AnaTrds() }
func (x AnaTrdsSelDurLeq) Ifc() interface{}          { return x.AnaTrds() }
func (x AnaTrdsSelDurLeq) AnaTrds() *ana.Trds        { return x.X.AnaTrds().SelDurLeq(x.I0.TmeTme()) }
func (x AnaTrdsSelDurGeq) Act()                      { x.AnaTrds() }
func (x AnaTrdsSelDurGeq) Ifc() interface{}          { return x.AnaTrds() }
func (x AnaTrdsSelDurGeq) AnaTrds() *ana.Trds        { return x.X.AnaTrds().SelDurGeq(x.I0.TmeTme()) }
func (x AnaTrdsSelPipEql) Act()                      { x.AnaTrds() }
func (x AnaTrdsSelPipEql) Ifc() interface{}          { return x.AnaTrds() }
func (x AnaTrdsSelPipEql) AnaTrds() *ana.Trds        { return x.X.AnaTrds().SelPipEql(x.I0.FltFlt()) }
func (x AnaTrdsSelPipNeq) Act()                      { x.AnaTrds() }
func (x AnaTrdsSelPipNeq) Ifc() interface{}          { return x.AnaTrds() }
func (x AnaTrdsSelPipNeq) AnaTrds() *ana.Trds        { return x.X.AnaTrds().SelPipNeq(x.I0.FltFlt()) }
func (x AnaTrdsSelPipLss) Act()                      { x.AnaTrds() }
func (x AnaTrdsSelPipLss) Ifc() interface{}          { return x.AnaTrds() }
func (x AnaTrdsSelPipLss) AnaTrds() *ana.Trds        { return x.X.AnaTrds().SelPipLss(x.I0.FltFlt()) }
func (x AnaTrdsSelPipGtr) Act()                      { x.AnaTrds() }
func (x AnaTrdsSelPipGtr) Ifc() interface{}          { return x.AnaTrds() }
func (x AnaTrdsSelPipGtr) AnaTrds() *ana.Trds        { return x.X.AnaTrds().SelPipGtr(x.I0.FltFlt()) }
func (x AnaTrdsSelPipLeq) Act()                      { x.AnaTrds() }
func (x AnaTrdsSelPipLeq) Ifc() interface{}          { return x.AnaTrds() }
func (x AnaTrdsSelPipLeq) AnaTrds() *ana.Trds        { return x.X.AnaTrds().SelPipLeq(x.I0.FltFlt()) }
func (x AnaTrdsSelPipGeq) Act()                      { x.AnaTrds() }
func (x AnaTrdsSelPipGeq) Ifc() interface{}          { return x.AnaTrds() }
func (x AnaTrdsSelPipGeq) AnaTrds() *ana.Trds        { return x.X.AnaTrds().SelPipGeq(x.I0.FltFlt()) }
func (x AnaTrdsSelClsRsnEql) Act()                   { x.AnaTrds() }
func (x AnaTrdsSelClsRsnEql) Ifc() interface{}       { return x.AnaTrds() }
func (x AnaTrdsSelClsRsnEql) AnaTrds() *ana.Trds     { return x.X.AnaTrds().SelClsRsnEql(x.I0.StrStr()) }
func (x AnaTrdsSelClsRsnNeq) Act()                   { x.AnaTrds() }
func (x AnaTrdsSelClsRsnNeq) Ifc() interface{}       { return x.AnaTrds() }
func (x AnaTrdsSelClsRsnNeq) AnaTrds() *ana.Trds     { return x.X.AnaTrds().SelClsRsnNeq(x.I0.StrStr()) }
func (x AnaTrdsSelClsRsnLss) Act()                   { x.AnaTrds() }
func (x AnaTrdsSelClsRsnLss) Ifc() interface{}       { return x.AnaTrds() }
func (x AnaTrdsSelClsRsnLss) AnaTrds() *ana.Trds     { return x.X.AnaTrds().SelClsRsnLss(x.I0.StrStr()) }
func (x AnaTrdsSelClsRsnGtr) Act()                   { x.AnaTrds() }
func (x AnaTrdsSelClsRsnGtr) Ifc() interface{}       { return x.AnaTrds() }
func (x AnaTrdsSelClsRsnGtr) AnaTrds() *ana.Trds     { return x.X.AnaTrds().SelClsRsnGtr(x.I0.StrStr()) }
func (x AnaTrdsSelClsRsnLeq) Act()                   { x.AnaTrds() }
func (x AnaTrdsSelClsRsnLeq) Ifc() interface{}       { return x.AnaTrds() }
func (x AnaTrdsSelClsRsnLeq) AnaTrds() *ana.Trds     { return x.X.AnaTrds().SelClsRsnLeq(x.I0.StrStr()) }
func (x AnaTrdsSelClsRsnGeq) Act()                   { x.AnaTrds() }
func (x AnaTrdsSelClsRsnGeq) Ifc() interface{}       { return x.AnaTrds() }
func (x AnaTrdsSelClsRsnGeq) AnaTrds() *ana.Trds     { return x.X.AnaTrds().SelClsRsnGeq(x.I0.StrStr()) }
func (x AnaTrdsSelClsSpdEql) Act()                   { x.AnaTrds() }
func (x AnaTrdsSelClsSpdEql) Ifc() interface{}       { return x.AnaTrds() }
func (x AnaTrdsSelClsSpdEql) AnaTrds() *ana.Trds     { return x.X.AnaTrds().SelClsSpdEql(x.I0.FltFlt()) }
func (x AnaTrdsSelClsSpdNeq) Act()                   { x.AnaTrds() }
func (x AnaTrdsSelClsSpdNeq) Ifc() interface{}       { return x.AnaTrds() }
func (x AnaTrdsSelClsSpdNeq) AnaTrds() *ana.Trds     { return x.X.AnaTrds().SelClsSpdNeq(x.I0.FltFlt()) }
func (x AnaTrdsSelClsSpdLss) Act()                   { x.AnaTrds() }
func (x AnaTrdsSelClsSpdLss) Ifc() interface{}       { return x.AnaTrds() }
func (x AnaTrdsSelClsSpdLss) AnaTrds() *ana.Trds     { return x.X.AnaTrds().SelClsSpdLss(x.I0.FltFlt()) }
func (x AnaTrdsSelClsSpdGtr) Act()                   { x.AnaTrds() }
func (x AnaTrdsSelClsSpdGtr) Ifc() interface{}       { return x.AnaTrds() }
func (x AnaTrdsSelClsSpdGtr) AnaTrds() *ana.Trds     { return x.X.AnaTrds().SelClsSpdGtr(x.I0.FltFlt()) }
func (x AnaTrdsSelClsSpdLeq) Act()                   { x.AnaTrds() }
func (x AnaTrdsSelClsSpdLeq) Ifc() interface{}       { return x.AnaTrds() }
func (x AnaTrdsSelClsSpdLeq) AnaTrds() *ana.Trds     { return x.X.AnaTrds().SelClsSpdLeq(x.I0.FltFlt()) }
func (x AnaTrdsSelClsSpdGeq) Act()                   { x.AnaTrds() }
func (x AnaTrdsSelClsSpdGeq) Ifc() interface{}       { return x.AnaTrds() }
func (x AnaTrdsSelClsSpdGeq) AnaTrds() *ana.Trds     { return x.X.AnaTrds().SelClsSpdGeq(x.I0.FltFlt()) }
func (x AnaTrdsSelOpnSpdEql) Act()                   { x.AnaTrds() }
func (x AnaTrdsSelOpnSpdEql) Ifc() interface{}       { return x.AnaTrds() }
func (x AnaTrdsSelOpnSpdEql) AnaTrds() *ana.Trds     { return x.X.AnaTrds().SelOpnSpdEql(x.I0.FltFlt()) }
func (x AnaTrdsSelOpnSpdNeq) Act()                   { x.AnaTrds() }
func (x AnaTrdsSelOpnSpdNeq) Ifc() interface{}       { return x.AnaTrds() }
func (x AnaTrdsSelOpnSpdNeq) AnaTrds() *ana.Trds     { return x.X.AnaTrds().SelOpnSpdNeq(x.I0.FltFlt()) }
func (x AnaTrdsSelOpnSpdLss) Act()                   { x.AnaTrds() }
func (x AnaTrdsSelOpnSpdLss) Ifc() interface{}       { return x.AnaTrds() }
func (x AnaTrdsSelOpnSpdLss) AnaTrds() *ana.Trds     { return x.X.AnaTrds().SelOpnSpdLss(x.I0.FltFlt()) }
func (x AnaTrdsSelOpnSpdGtr) Act()                   { x.AnaTrds() }
func (x AnaTrdsSelOpnSpdGtr) Ifc() interface{}       { return x.AnaTrds() }
func (x AnaTrdsSelOpnSpdGtr) AnaTrds() *ana.Trds     { return x.X.AnaTrds().SelOpnSpdGtr(x.I0.FltFlt()) }
func (x AnaTrdsSelOpnSpdLeq) Act()                   { x.AnaTrds() }
func (x AnaTrdsSelOpnSpdLeq) Ifc() interface{}       { return x.AnaTrds() }
func (x AnaTrdsSelOpnSpdLeq) AnaTrds() *ana.Trds     { return x.X.AnaTrds().SelOpnSpdLeq(x.I0.FltFlt()) }
func (x AnaTrdsSelOpnSpdGeq) Act()                   { x.AnaTrds() }
func (x AnaTrdsSelOpnSpdGeq) Ifc() interface{}       { return x.AnaTrds() }
func (x AnaTrdsSelOpnSpdGeq) AnaTrds() *ana.Trds     { return x.X.AnaTrds().SelOpnSpdGeq(x.I0.FltFlt()) }
func (x AnaTrdsSelClsAskEql) Act()                   { x.AnaTrds() }
func (x AnaTrdsSelClsAskEql) Ifc() interface{}       { return x.AnaTrds() }
func (x AnaTrdsSelClsAskEql) AnaTrds() *ana.Trds     { return x.X.AnaTrds().SelClsAskEql(x.I0.FltFlt()) }
func (x AnaTrdsSelClsAskNeq) Act()                   { x.AnaTrds() }
func (x AnaTrdsSelClsAskNeq) Ifc() interface{}       { return x.AnaTrds() }
func (x AnaTrdsSelClsAskNeq) AnaTrds() *ana.Trds     { return x.X.AnaTrds().SelClsAskNeq(x.I0.FltFlt()) }
func (x AnaTrdsSelClsAskLss) Act()                   { x.AnaTrds() }
func (x AnaTrdsSelClsAskLss) Ifc() interface{}       { return x.AnaTrds() }
func (x AnaTrdsSelClsAskLss) AnaTrds() *ana.Trds     { return x.X.AnaTrds().SelClsAskLss(x.I0.FltFlt()) }
func (x AnaTrdsSelClsAskGtr) Act()                   { x.AnaTrds() }
func (x AnaTrdsSelClsAskGtr) Ifc() interface{}       { return x.AnaTrds() }
func (x AnaTrdsSelClsAskGtr) AnaTrds() *ana.Trds     { return x.X.AnaTrds().SelClsAskGtr(x.I0.FltFlt()) }
func (x AnaTrdsSelClsAskLeq) Act()                   { x.AnaTrds() }
func (x AnaTrdsSelClsAskLeq) Ifc() interface{}       { return x.AnaTrds() }
func (x AnaTrdsSelClsAskLeq) AnaTrds() *ana.Trds     { return x.X.AnaTrds().SelClsAskLeq(x.I0.FltFlt()) }
func (x AnaTrdsSelClsAskGeq) Act()                   { x.AnaTrds() }
func (x AnaTrdsSelClsAskGeq) Ifc() interface{}       { return x.AnaTrds() }
func (x AnaTrdsSelClsAskGeq) AnaTrds() *ana.Trds     { return x.X.AnaTrds().SelClsAskGeq(x.I0.FltFlt()) }
func (x AnaTrdsSelOpnAskEql) Act()                   { x.AnaTrds() }
func (x AnaTrdsSelOpnAskEql) Ifc() interface{}       { return x.AnaTrds() }
func (x AnaTrdsSelOpnAskEql) AnaTrds() *ana.Trds     { return x.X.AnaTrds().SelOpnAskEql(x.I0.FltFlt()) }
func (x AnaTrdsSelOpnAskNeq) Act()                   { x.AnaTrds() }
func (x AnaTrdsSelOpnAskNeq) Ifc() interface{}       { return x.AnaTrds() }
func (x AnaTrdsSelOpnAskNeq) AnaTrds() *ana.Trds     { return x.X.AnaTrds().SelOpnAskNeq(x.I0.FltFlt()) }
func (x AnaTrdsSelOpnAskLss) Act()                   { x.AnaTrds() }
func (x AnaTrdsSelOpnAskLss) Ifc() interface{}       { return x.AnaTrds() }
func (x AnaTrdsSelOpnAskLss) AnaTrds() *ana.Trds     { return x.X.AnaTrds().SelOpnAskLss(x.I0.FltFlt()) }
func (x AnaTrdsSelOpnAskGtr) Act()                   { x.AnaTrds() }
func (x AnaTrdsSelOpnAskGtr) Ifc() interface{}       { return x.AnaTrds() }
func (x AnaTrdsSelOpnAskGtr) AnaTrds() *ana.Trds     { return x.X.AnaTrds().SelOpnAskGtr(x.I0.FltFlt()) }
func (x AnaTrdsSelOpnAskLeq) Act()                   { x.AnaTrds() }
func (x AnaTrdsSelOpnAskLeq) Ifc() interface{}       { return x.AnaTrds() }
func (x AnaTrdsSelOpnAskLeq) AnaTrds() *ana.Trds     { return x.X.AnaTrds().SelOpnAskLeq(x.I0.FltFlt()) }
func (x AnaTrdsSelOpnAskGeq) Act()                   { x.AnaTrds() }
func (x AnaTrdsSelOpnAskGeq) Ifc() interface{}       { return x.AnaTrds() }
func (x AnaTrdsSelOpnAskGeq) AnaTrds() *ana.Trds     { return x.X.AnaTrds().SelOpnAskGeq(x.I0.FltFlt()) }
func (x AnaTrdsSelClsBidEql) Act()                   { x.AnaTrds() }
func (x AnaTrdsSelClsBidEql) Ifc() interface{}       { return x.AnaTrds() }
func (x AnaTrdsSelClsBidEql) AnaTrds() *ana.Trds     { return x.X.AnaTrds().SelClsBidEql(x.I0.FltFlt()) }
func (x AnaTrdsSelClsBidNeq) Act()                   { x.AnaTrds() }
func (x AnaTrdsSelClsBidNeq) Ifc() interface{}       { return x.AnaTrds() }
func (x AnaTrdsSelClsBidNeq) AnaTrds() *ana.Trds     { return x.X.AnaTrds().SelClsBidNeq(x.I0.FltFlt()) }
func (x AnaTrdsSelClsBidLss) Act()                   { x.AnaTrds() }
func (x AnaTrdsSelClsBidLss) Ifc() interface{}       { return x.AnaTrds() }
func (x AnaTrdsSelClsBidLss) AnaTrds() *ana.Trds     { return x.X.AnaTrds().SelClsBidLss(x.I0.FltFlt()) }
func (x AnaTrdsSelClsBidGtr) Act()                   { x.AnaTrds() }
func (x AnaTrdsSelClsBidGtr) Ifc() interface{}       { return x.AnaTrds() }
func (x AnaTrdsSelClsBidGtr) AnaTrds() *ana.Trds     { return x.X.AnaTrds().SelClsBidGtr(x.I0.FltFlt()) }
func (x AnaTrdsSelClsBidLeq) Act()                   { x.AnaTrds() }
func (x AnaTrdsSelClsBidLeq) Ifc() interface{}       { return x.AnaTrds() }
func (x AnaTrdsSelClsBidLeq) AnaTrds() *ana.Trds     { return x.X.AnaTrds().SelClsBidLeq(x.I0.FltFlt()) }
func (x AnaTrdsSelClsBidGeq) Act()                   { x.AnaTrds() }
func (x AnaTrdsSelClsBidGeq) Ifc() interface{}       { return x.AnaTrds() }
func (x AnaTrdsSelClsBidGeq) AnaTrds() *ana.Trds     { return x.X.AnaTrds().SelClsBidGeq(x.I0.FltFlt()) }
func (x AnaTrdsSelOpnBidEql) Act()                   { x.AnaTrds() }
func (x AnaTrdsSelOpnBidEql) Ifc() interface{}       { return x.AnaTrds() }
func (x AnaTrdsSelOpnBidEql) AnaTrds() *ana.Trds     { return x.X.AnaTrds().SelOpnBidEql(x.I0.FltFlt()) }
func (x AnaTrdsSelOpnBidNeq) Act()                   { x.AnaTrds() }
func (x AnaTrdsSelOpnBidNeq) Ifc() interface{}       { return x.AnaTrds() }
func (x AnaTrdsSelOpnBidNeq) AnaTrds() *ana.Trds     { return x.X.AnaTrds().SelOpnBidNeq(x.I0.FltFlt()) }
func (x AnaTrdsSelOpnBidLss) Act()                   { x.AnaTrds() }
func (x AnaTrdsSelOpnBidLss) Ifc() interface{}       { return x.AnaTrds() }
func (x AnaTrdsSelOpnBidLss) AnaTrds() *ana.Trds     { return x.X.AnaTrds().SelOpnBidLss(x.I0.FltFlt()) }
func (x AnaTrdsSelOpnBidGtr) Act()                   { x.AnaTrds() }
func (x AnaTrdsSelOpnBidGtr) Ifc() interface{}       { return x.AnaTrds() }
func (x AnaTrdsSelOpnBidGtr) AnaTrds() *ana.Trds     { return x.X.AnaTrds().SelOpnBidGtr(x.I0.FltFlt()) }
func (x AnaTrdsSelOpnBidLeq) Act()                   { x.AnaTrds() }
func (x AnaTrdsSelOpnBidLeq) Ifc() interface{}       { return x.AnaTrds() }
func (x AnaTrdsSelOpnBidLeq) AnaTrds() *ana.Trds     { return x.X.AnaTrds().SelOpnBidLeq(x.I0.FltFlt()) }
func (x AnaTrdsSelOpnBidGeq) Act()                   { x.AnaTrds() }
func (x AnaTrdsSelOpnBidGeq) Ifc() interface{}       { return x.AnaTrds() }
func (x AnaTrdsSelOpnBidGeq) AnaTrds() *ana.Trds     { return x.X.AnaTrds().SelOpnBidGeq(x.I0.FltFlt()) }
func (x AnaTrdsSelClsTmeEql) Act()                   { x.AnaTrds() }
func (x AnaTrdsSelClsTmeEql) Ifc() interface{}       { return x.AnaTrds() }
func (x AnaTrdsSelClsTmeEql) AnaTrds() *ana.Trds     { return x.X.AnaTrds().SelClsTmeEql(x.I0.TmeTme()) }
func (x AnaTrdsSelClsTmeNeq) Act()                   { x.AnaTrds() }
func (x AnaTrdsSelClsTmeNeq) Ifc() interface{}       { return x.AnaTrds() }
func (x AnaTrdsSelClsTmeNeq) AnaTrds() *ana.Trds     { return x.X.AnaTrds().SelClsTmeNeq(x.I0.TmeTme()) }
func (x AnaTrdsSelClsTmeLss) Act()                   { x.AnaTrds() }
func (x AnaTrdsSelClsTmeLss) Ifc() interface{}       { return x.AnaTrds() }
func (x AnaTrdsSelClsTmeLss) AnaTrds() *ana.Trds     { return x.X.AnaTrds().SelClsTmeLss(x.I0.TmeTme()) }
func (x AnaTrdsSelClsTmeGtr) Act()                   { x.AnaTrds() }
func (x AnaTrdsSelClsTmeGtr) Ifc() interface{}       { return x.AnaTrds() }
func (x AnaTrdsSelClsTmeGtr) AnaTrds() *ana.Trds     { return x.X.AnaTrds().SelClsTmeGtr(x.I0.TmeTme()) }
func (x AnaTrdsSelClsTmeLeq) Act()                   { x.AnaTrds() }
func (x AnaTrdsSelClsTmeLeq) Ifc() interface{}       { return x.AnaTrds() }
func (x AnaTrdsSelClsTmeLeq) AnaTrds() *ana.Trds     { return x.X.AnaTrds().SelClsTmeLeq(x.I0.TmeTme()) }
func (x AnaTrdsSelClsTmeGeq) Act()                   { x.AnaTrds() }
func (x AnaTrdsSelClsTmeGeq) Ifc() interface{}       { return x.AnaTrds() }
func (x AnaTrdsSelClsTmeGeq) AnaTrds() *ana.Trds     { return x.X.AnaTrds().SelClsTmeGeq(x.I0.TmeTme()) }
func (x AnaTrdsSelOpnTmeEql) Act()                   { x.AnaTrds() }
func (x AnaTrdsSelOpnTmeEql) Ifc() interface{}       { return x.AnaTrds() }
func (x AnaTrdsSelOpnTmeEql) AnaTrds() *ana.Trds     { return x.X.AnaTrds().SelOpnTmeEql(x.I0.TmeTme()) }
func (x AnaTrdsSelOpnTmeNeq) Act()                   { x.AnaTrds() }
func (x AnaTrdsSelOpnTmeNeq) Ifc() interface{}       { return x.AnaTrds() }
func (x AnaTrdsSelOpnTmeNeq) AnaTrds() *ana.Trds     { return x.X.AnaTrds().SelOpnTmeNeq(x.I0.TmeTme()) }
func (x AnaTrdsSelOpnTmeLss) Act()                   { x.AnaTrds() }
func (x AnaTrdsSelOpnTmeLss) Ifc() interface{}       { return x.AnaTrds() }
func (x AnaTrdsSelOpnTmeLss) AnaTrds() *ana.Trds     { return x.X.AnaTrds().SelOpnTmeLss(x.I0.TmeTme()) }
func (x AnaTrdsSelOpnTmeGtr) Act()                   { x.AnaTrds() }
func (x AnaTrdsSelOpnTmeGtr) Ifc() interface{}       { return x.AnaTrds() }
func (x AnaTrdsSelOpnTmeGtr) AnaTrds() *ana.Trds     { return x.X.AnaTrds().SelOpnTmeGtr(x.I0.TmeTme()) }
func (x AnaTrdsSelOpnTmeLeq) Act()                   { x.AnaTrds() }
func (x AnaTrdsSelOpnTmeLeq) Ifc() interface{}       { return x.AnaTrds() }
func (x AnaTrdsSelOpnTmeLeq) AnaTrds() *ana.Trds     { return x.X.AnaTrds().SelOpnTmeLeq(x.I0.TmeTme()) }
func (x AnaTrdsSelOpnTmeGeq) Act()                   { x.AnaTrds() }
func (x AnaTrdsSelOpnTmeGeq) Ifc() interface{}       { return x.AnaTrds() }
func (x AnaTrdsSelOpnTmeGeq) AnaTrds() *ana.Trds     { return x.X.AnaTrds().SelOpnTmeGeq(x.I0.TmeTme()) }
func (x AnaTrdsOpnTmes) Act()                        { x.TmesTmes() }
func (x AnaTrdsOpnTmes) Ifc() interface{}            { return x.TmesTmes() }
func (x AnaTrdsOpnTmes) TmesTmes() *tmes.Tmes        { return x.X.AnaTrds().OpnTmes() }
func (x AnaTrdsClsTmes) Act()                        { x.TmesTmes() }
func (x AnaTrdsClsTmes) Ifc() interface{}            { return x.TmesTmes() }
func (x AnaTrdsClsTmes) TmesTmes() *tmes.Tmes        { return x.X.AnaTrds().ClsTmes() }
func (x AnaTrdsOpnBids) Act()                        { x.FltsFlts() }
func (x AnaTrdsOpnBids) Ifc() interface{}            { return x.FltsFlts() }
func (x AnaTrdsOpnBids) FltsFlts() *flts.Flts        { return x.X.AnaTrds().OpnBids() }
func (x AnaTrdsClsBids) Act()                        { x.FltsFlts() }
func (x AnaTrdsClsBids) Ifc() interface{}            { return x.FltsFlts() }
func (x AnaTrdsClsBids) FltsFlts() *flts.Flts        { return x.X.AnaTrds().ClsBids() }
func (x AnaTrdsOpnAsks) Act()                        { x.FltsFlts() }
func (x AnaTrdsOpnAsks) Ifc() interface{}            { return x.FltsFlts() }
func (x AnaTrdsOpnAsks) FltsFlts() *flts.Flts        { return x.X.AnaTrds().OpnAsks() }
func (x AnaTrdsClsAsks) Act()                        { x.FltsFlts() }
func (x AnaTrdsClsAsks) Ifc() interface{}            { return x.FltsFlts() }
func (x AnaTrdsClsAsks) FltsFlts() *flts.Flts        { return x.X.AnaTrds().ClsAsks() }
func (x AnaTrdsOpnSpds) Act()                        { x.FltsFlts() }
func (x AnaTrdsOpnSpds) Ifc() interface{}            { return x.FltsFlts() }
func (x AnaTrdsOpnSpds) FltsFlts() *flts.Flts        { return x.X.AnaTrds().OpnSpds() }
func (x AnaTrdsClsSpds) Act()                        { x.FltsFlts() }
func (x AnaTrdsClsSpds) Ifc() interface{}            { return x.FltsFlts() }
func (x AnaTrdsClsSpds) FltsFlts() *flts.Flts        { return x.X.AnaTrds().ClsSpds() }
func (x AnaTrdsClsRsns) Act()                        { x.StrsStrs() }
func (x AnaTrdsClsRsns) Ifc() interface{}            { return x.StrsStrs() }
func (x AnaTrdsClsRsns) StrsStrs() *strs.Strs        { return x.X.AnaTrds().ClsRsns() }
func (x AnaTrdsPips) Act()                           { x.FltsFlts() }
func (x AnaTrdsPips) Ifc() interface{}               { return x.FltsFlts() }
func (x AnaTrdsPips) FltsFlts() *flts.Flts           { return x.X.AnaTrds().Pips() }
func (x AnaTrdsDurs) Act()                           { x.TmesTmes() }
func (x AnaTrdsDurs) Ifc() interface{}               { return x.TmesTmes() }
func (x AnaTrdsDurs) TmesTmes() *tmes.Tmes           { return x.X.AnaTrds().Durs() }
func (x AnaTrdsIsLongs) Act()                        { x.BolsBols() }
func (x AnaTrdsIsLongs) Ifc() interface{}            { return x.BolsBols() }
func (x AnaTrdsIsLongs) BolsBols() *bols.Bols        { return x.X.AnaTrds().IsLongs() }
func (x AnaTrdsPnlPcts) Act()                        { x.FltsFlts() }
func (x AnaTrdsPnlPcts) Ifc() interface{}            { return x.FltsFlts() }
func (x AnaTrdsPnlPcts) FltsFlts() *flts.Flts        { return x.X.AnaTrds().PnlPcts() }
func (x AnaTrdsPnlPctPredicts) Act()                 { x.FltsFlts() }
func (x AnaTrdsPnlPctPredicts) Ifc() interface{}     { return x.FltsFlts() }
func (x AnaTrdsPnlPctPredicts) FltsFlts() *flts.Flts { return x.X.AnaTrds().PnlPctPredicts() }
func (x AnaTrdsPnlUsds) Act()                        { x.FltsFlts() }
func (x AnaTrdsPnlUsds) Ifc() interface{}            { return x.FltsFlts() }
func (x AnaTrdsPnlUsds) FltsFlts() *flts.Flts        { return x.X.AnaTrds().PnlUsds() }
func (x AnaTrdsPnlGrsUsds) Act()                     { x.FltsFlts() }
func (x AnaTrdsPnlGrsUsds) Ifc() interface{}         { return x.FltsFlts() }
func (x AnaTrdsPnlGrsUsds) FltsFlts() *flts.Flts     { return x.X.AnaTrds().PnlGrsUsds() }
func (x AnaTrdsCstComUsds) Act()                     { x.FltsFlts() }
func (x AnaTrdsCstComUsds) Ifc() interface{}         { return x.FltsFlts() }
func (x AnaTrdsCstComUsds) FltsFlts() *flts.Flts     { return x.X.AnaTrds().CstComUsds() }
func (x AnaTrdsCstClsSpdUsds) Act()                  { x.FltsFlts() }
func (x AnaTrdsCstClsSpdUsds) Ifc() interface{}      { return x.FltsFlts() }
func (x AnaTrdsCstClsSpdUsds) FltsFlts() *flts.Flts  { return x.X.AnaTrds().CstClsSpdUsds() }
func (x AnaTrdsCstOpnSpdUsds) Act()                  { x.FltsFlts() }
func (x AnaTrdsCstOpnSpdUsds) Ifc() interface{}      { return x.FltsFlts() }
func (x AnaTrdsCstOpnSpdUsds) FltsFlts() *flts.Flts  { return x.X.AnaTrds().CstOpnSpdUsds() }
func (x AnaTrdsOpnBalUsds) Act()                     { x.FltsFlts() }
func (x AnaTrdsOpnBalUsds) Ifc() interface{}         { return x.FltsFlts() }
func (x AnaTrdsOpnBalUsds) FltsFlts() *flts.Flts     { return x.X.AnaTrds().OpnBalUsds() }
func (x AnaTrdsClsBalUsds) Act()                     { x.FltsFlts() }
func (x AnaTrdsClsBalUsds) Ifc() interface{}         { return x.FltsFlts() }
func (x AnaTrdsClsBalUsds) FltsFlts() *flts.Flts     { return x.X.AnaTrds().ClsBalUsds() }
func (x AnaTrdsClsBalUsdActs) Act()                  { x.FltsFlts() }
func (x AnaTrdsClsBalUsdActs) Ifc() interface{}      { return x.FltsFlts() }
func (x AnaTrdsClsBalUsdActs) FltsFlts() *flts.Flts  { return x.X.AnaTrds().ClsBalUsdActs() }
func (x AnaTrdsTrdPcts) Act()                        { x.FltsFlts() }
func (x AnaTrdsTrdPcts) Ifc() interface{}            { return x.FltsFlts() }
func (x AnaTrdsTrdPcts) FltsFlts() *flts.Flts        { return x.X.AnaTrds().TrdPcts() }
func (x AnaTrdsMrgnRtios) Act()                      { x.FltsFlts() }
func (x AnaTrdsMrgnRtios) Ifc() interface{}          { return x.FltsFlts() }
func (x AnaTrdsMrgnRtios) FltsFlts() *flts.Flts      { return x.X.AnaTrds().MrgnRtios() }
func (x AnaTrdsUnitss) Act()                         { x.FltsFlts() }
func (x AnaTrdsUnitss) Ifc() interface{}             { return x.FltsFlts() }
func (x AnaTrdsUnitss) FltsFlts() *flts.Flts         { return x.X.AnaTrds().Unitss() }
func (x AnaTrdsInstrs) Act()                         { x.StrsStrs() }
func (x AnaTrdsInstrs) Ifc() interface{}             { return x.StrsStrs() }
func (x AnaTrdsInstrs) StrsStrs() *strs.Strs         { return x.X.AnaTrds().Instrs() }
func (x AnaTrdsOpnReqs) Act()                        { x.StrsStrs() }
func (x AnaTrdsOpnReqs) Ifc() interface{}            { return x.StrsStrs() }
func (x AnaTrdsOpnReqs) StrsStrs() *strs.Strs        { return x.X.AnaTrds().OpnReqs() }
func (x AnaTrdsOpnRess) Act()                        { x.StrsStrs() }
func (x AnaTrdsOpnRess) Ifc() interface{}            { return x.StrsStrs() }
func (x AnaTrdsOpnRess) StrsStrs() *strs.Strs        { return x.X.AnaTrds().OpnRess() }
func (x AnaTrdsClsReqs) Act()                        { x.StrsStrs() }
func (x AnaTrdsClsReqs) Ifc() interface{}            { return x.StrsStrs() }
func (x AnaTrdsClsReqs) StrsStrs() *strs.Strs        { return x.X.AnaTrds().ClsReqs() }
func (x AnaTrdsClsRess) Act()                        { x.StrsStrs() }
func (x AnaTrdsClsRess) Ifc() interface{}            { return x.StrsStrs() }
func (x AnaTrdsClsRess) StrsStrs() *strs.Strs        { return x.X.AnaTrds().ClsRess() }
func (x AnaPrfmDlt) Act()                            { x.AnaPrfmDlt() }
func (x AnaPrfmDlt) Ifc() interface{}                { return x.AnaPrfmDlt() }
func (x AnaPrfmDlt) AnaPrfmDlt() *ana.PrfmDlt        { return x.X.AnaPrfm().Dlt(x.I0.AnaPrfm()) }
func (x AnaPrfmsCnt) Act()                           { x.UntUnt() }
func (x AnaPrfmsCnt) Ifc() interface{}               { return x.UntUnt() }
func (x AnaPrfmsCnt) UntUnt() unt.Unt                { return x.X.AnaPrfms().Cnt() }
func (x AnaPrfmsCpy) Act()                           { x.AnaPrfms() }
func (x AnaPrfmsCpy) Ifc() interface{}               { return x.AnaPrfms() }
func (x AnaPrfmsCpy) AnaPrfms() *ana.Prfms           { return x.X.AnaPrfms().Cpy() }
func (x AnaPrfmsClr) Act()                           { x.AnaPrfms() }
func (x AnaPrfmsClr) Ifc() interface{}               { return x.AnaPrfms() }
func (x AnaPrfmsClr) AnaPrfms() *ana.Prfms           { return x.X.AnaPrfms().Clr() }
func (x AnaPrfmsRand) Act()                          { x.AnaPrfms() }
func (x AnaPrfmsRand) Ifc() interface{}              { return x.AnaPrfms() }
func (x AnaPrfmsRand) AnaPrfms() *ana.Prfms          { return x.X.AnaPrfms().Rand() }
func (x AnaPrfmsMrg) Act()                           { x.AnaPrfms() }
func (x AnaPrfmsMrg) Ifc() interface{}               { return x.AnaPrfms() }
func (x AnaPrfmsMrg) AnaPrfms() *ana.Prfms {
	var i0 []*ana.Prfms
	for _, cur := range x.I0 {
		i0 = append(i0, cur.AnaPrfms())
	}
	return x.X.AnaPrfms().Mrg(i0...)
}
func (x AnaPrfmsPush) Act()             { x.AnaPrfms() }
func (x AnaPrfmsPush) Ifc() interface{} { return x.AnaPrfms() }
func (x AnaPrfmsPush) AnaPrfms() *ana.Prfms {
	var i0 []*ana.Prfm
	for _, cur := range x.I0 {
		i0 = append(i0, cur.AnaPrfm())
	}
	return x.X.AnaPrfms().Push(i0...)
}
func (x AnaPrfmsPop) Act()               { x.AnaPrfm() }
func (x AnaPrfmsPop) Ifc() interface{}   { return x.AnaPrfm() }
func (x AnaPrfmsPop) AnaPrfm() *ana.Prfm { return x.X.AnaPrfms().Pop() }
func (x AnaPrfmsQue) Act()               { x.AnaPrfms() }
func (x AnaPrfmsQue) Ifc() interface{}   { return x.AnaPrfms() }
func (x AnaPrfmsQue) AnaPrfms() *ana.Prfms {
	var i0 []*ana.Prfm
	for _, cur := range x.I0 {
		i0 = append(i0, cur.AnaPrfm())
	}
	return x.X.AnaPrfms().Que(i0...)
}
func (x AnaPrfmsDque) Act()                       { x.AnaPrfm() }
func (x AnaPrfmsDque) Ifc() interface{}           { return x.AnaPrfm() }
func (x AnaPrfmsDque) AnaPrfm() *ana.Prfm         { return x.X.AnaPrfms().Dque() }
func (x AnaPrfmsIns) Act()                        { x.AnaPrfms() }
func (x AnaPrfmsIns) Ifc() interface{}            { return x.AnaPrfms() }
func (x AnaPrfmsIns) AnaPrfms() *ana.Prfms        { return x.X.AnaPrfms().Ins(x.I0.UntUnt(), x.I1.AnaPrfm()) }
func (x AnaPrfmsUpd) Act()                        { x.AnaPrfms() }
func (x AnaPrfmsUpd) Ifc() interface{}            { return x.AnaPrfms() }
func (x AnaPrfmsUpd) AnaPrfms() *ana.Prfms        { return x.X.AnaPrfms().Upd(x.I0.UntUnt(), x.I1.AnaPrfm()) }
func (x AnaPrfmsDel) Act()                        { x.AnaPrfm() }
func (x AnaPrfmsDel) Ifc() interface{}            { return x.AnaPrfm() }
func (x AnaPrfmsDel) AnaPrfm() *ana.Prfm          { return x.X.AnaPrfms().Del(x.I0.UntUnt()) }
func (x AnaPrfmsAt) Act()                         { x.AnaPrfm() }
func (x AnaPrfmsAt) Ifc() interface{}             { return x.AnaPrfm() }
func (x AnaPrfmsAt) AnaPrfm() *ana.Prfm           { return x.X.AnaPrfms().At(x.I0.UntUnt()) }
func (x AnaPrfmsIn) Act()                         { x.AnaPrfms() }
func (x AnaPrfmsIn) Ifc() interface{}             { return x.AnaPrfms() }
func (x AnaPrfmsIn) AnaPrfms() *ana.Prfms         { return x.X.AnaPrfms().In(x.I0.UntUnt(), x.I1.UntUnt()) }
func (x AnaPrfmsInBnd) Act()                      { x.AnaPrfms() }
func (x AnaPrfmsInBnd) Ifc() interface{}          { return x.AnaPrfms() }
func (x AnaPrfmsInBnd) AnaPrfms() *ana.Prfms      { return x.X.AnaPrfms().InBnd(x.I0.BndBnd()) }
func (x AnaPrfmsFrom) Act()                       { x.AnaPrfms() }
func (x AnaPrfmsFrom) Ifc() interface{}           { return x.AnaPrfms() }
func (x AnaPrfmsFrom) AnaPrfms() *ana.Prfms       { return x.X.AnaPrfms().From(x.I0.UntUnt()) }
func (x AnaPrfmsTo) Act()                         { x.AnaPrfms() }
func (x AnaPrfmsTo) Ifc() interface{}             { return x.AnaPrfms() }
func (x AnaPrfmsTo) AnaPrfms() *ana.Prfms         { return x.X.AnaPrfms().To(x.I0.UntUnt()) }
func (x AnaPrfmsFst) Act()                        { x.AnaPrfm() }
func (x AnaPrfmsFst) Ifc() interface{}            { return x.AnaPrfm() }
func (x AnaPrfmsFst) AnaPrfm() *ana.Prfm          { return x.X.AnaPrfms().Fst() }
func (x AnaPrfmsMdl) Act()                        { x.AnaPrfm() }
func (x AnaPrfmsMdl) Ifc() interface{}            { return x.AnaPrfm() }
func (x AnaPrfmsMdl) AnaPrfm() *ana.Prfm          { return x.X.AnaPrfms().Mdl() }
func (x AnaPrfmsLst) Act()                        { x.AnaPrfm() }
func (x AnaPrfmsLst) Ifc() interface{}            { return x.AnaPrfm() }
func (x AnaPrfmsLst) AnaPrfm() *ana.Prfm          { return x.X.AnaPrfms().Lst() }
func (x AnaPrfmsFstIdx) Act()                     { x.UntUnt() }
func (x AnaPrfmsFstIdx) Ifc() interface{}         { return x.UntUnt() }
func (x AnaPrfmsFstIdx) UntUnt() unt.Unt          { return x.X.AnaPrfms().FstIdx() }
func (x AnaPrfmsMdlIdx) Act()                     { x.UntUnt() }
func (x AnaPrfmsMdlIdx) Ifc() interface{}         { return x.UntUnt() }
func (x AnaPrfmsMdlIdx) UntUnt() unt.Unt          { return x.X.AnaPrfms().MdlIdx() }
func (x AnaPrfmsLstIdx) Act()                     { x.UntUnt() }
func (x AnaPrfmsLstIdx) Ifc() interface{}         { return x.UntUnt() }
func (x AnaPrfmsLstIdx) UntUnt() unt.Unt          { return x.X.AnaPrfms().LstIdx() }
func (x AnaPrfmsRev) Act()                        { x.AnaPrfms() }
func (x AnaPrfmsRev) Ifc() interface{}            { return x.AnaPrfms() }
func (x AnaPrfmsRev) AnaPrfms() *ana.Prfms        { return x.X.AnaPrfms().Rev() }
func (x AnaPrfmsPnlPcts) Act()                    { x.FltsFlts() }
func (x AnaPrfmsPnlPcts) Ifc() interface{}        { return x.FltsFlts() }
func (x AnaPrfmsPnlPcts) FltsFlts() *flts.Flts    { return x.X.AnaPrfms().PnlPcts() }
func (x AnaPrfmsScsPcts) Act()                    { x.FltsFlts() }
func (x AnaPrfmsScsPcts) Ifc() interface{}        { return x.FltsFlts() }
func (x AnaPrfmsScsPcts) FltsFlts() *flts.Flts    { return x.X.AnaPrfms().ScsPcts() }
func (x AnaPrfmsPipPerDays) Act()                 { x.FltsFlts() }
func (x AnaPrfmsPipPerDays) Ifc() interface{}     { return x.FltsFlts() }
func (x AnaPrfmsPipPerDays) FltsFlts() *flts.Flts { return x.X.AnaPrfms().PipPerDays() }
func (x AnaPrfmsUsdPerDays) Act()                 { x.FltsFlts() }
func (x AnaPrfmsUsdPerDays) Ifc() interface{}     { return x.FltsFlts() }
func (x AnaPrfmsUsdPerDays) FltsFlts() *flts.Flts { return x.X.AnaPrfms().UsdPerDays() }
func (x AnaPrfmsScsPerDays) Act()                 { x.FltsFlts() }
func (x AnaPrfmsScsPerDays) Ifc() interface{}     { return x.FltsFlts() }
func (x AnaPrfmsScsPerDays) FltsFlts() *flts.Flts { return x.X.AnaPrfms().ScsPerDays() }
func (x AnaPrfmsOpnPerDays) Act()                 { x.FltsFlts() }
func (x AnaPrfmsOpnPerDays) Ifc() interface{}     { return x.FltsFlts() }
func (x AnaPrfmsOpnPerDays) FltsFlts() *flts.Flts { return x.X.AnaPrfms().OpnPerDays() }
func (x AnaPrfmsPnlUsds) Act()                    { x.FltsFlts() }
func (x AnaPrfmsPnlUsds) Ifc() interface{}        { return x.FltsFlts() }
func (x AnaPrfmsPnlUsds) FltsFlts() *flts.Flts    { return x.X.AnaPrfms().PnlUsds() }
func (x AnaPrfmsPipAvgs) Act()                    { x.FltsFlts() }
func (x AnaPrfmsPipAvgs) Ifc() interface{}        { return x.FltsFlts() }
func (x AnaPrfmsPipAvgs) FltsFlts() *flts.Flts    { return x.X.AnaPrfms().PipAvgs() }
func (x AnaPrfmsPipMdns) Act()                    { x.FltsFlts() }
func (x AnaPrfmsPipMdns) Ifc() interface{}        { return x.FltsFlts() }
func (x AnaPrfmsPipMdns) FltsFlts() *flts.Flts    { return x.X.AnaPrfms().PipMdns() }
func (x AnaPrfmsPipMins) Act()                    { x.FltsFlts() }
func (x AnaPrfmsPipMins) Ifc() interface{}        { return x.FltsFlts() }
func (x AnaPrfmsPipMins) FltsFlts() *flts.Flts    { return x.X.AnaPrfms().PipMins() }
func (x AnaPrfmsPipMaxs) Act()                    { x.FltsFlts() }
func (x AnaPrfmsPipMaxs) Ifc() interface{}        { return x.FltsFlts() }
func (x AnaPrfmsPipMaxs) FltsFlts() *flts.Flts    { return x.X.AnaPrfms().PipMaxs() }
func (x AnaPrfmsPipSums) Act()                    { x.FltsFlts() }
func (x AnaPrfmsPipSums) Ifc() interface{}        { return x.FltsFlts() }
func (x AnaPrfmsPipSums) FltsFlts() *flts.Flts    { return x.X.AnaPrfms().PipSums() }
func (x AnaPrfmsDurAvgs) Act()                    { x.TmesTmes() }
func (x AnaPrfmsDurAvgs) Ifc() interface{}        { return x.TmesTmes() }
func (x AnaPrfmsDurAvgs) TmesTmes() *tmes.Tmes    { return x.X.AnaPrfms().DurAvgs() }
func (x AnaPrfmsDurMdns) Act()                    { x.TmesTmes() }
func (x AnaPrfmsDurMdns) Ifc() interface{}        { return x.TmesTmes() }
func (x AnaPrfmsDurMdns) TmesTmes() *tmes.Tmes    { return x.X.AnaPrfms().DurMdns() }
func (x AnaPrfmsDurMins) Act()                    { x.TmesTmes() }
func (x AnaPrfmsDurMins) Ifc() interface{}        { return x.TmesTmes() }
func (x AnaPrfmsDurMins) TmesTmes() *tmes.Tmes    { return x.X.AnaPrfms().DurMins() }
func (x AnaPrfmsDurMaxs) Act()                    { x.TmesTmes() }
func (x AnaPrfmsDurMaxs) Ifc() interface{}        { return x.TmesTmes() }
func (x AnaPrfmsDurMaxs) TmesTmes() *tmes.Tmes    { return x.X.AnaPrfms().DurMaxs() }
func (x AnaPrfmsLosLimMaxs) Act()                 { x.FltsFlts() }
func (x AnaPrfmsLosLimMaxs) Ifc() interface{}     { return x.FltsFlts() }
func (x AnaPrfmsLosLimMaxs) FltsFlts() *flts.Flts { return x.X.AnaPrfms().LosLimMaxs() }
func (x AnaPrfmsDurLimMaxs) Act()                 { x.TmesTmes() }
func (x AnaPrfmsDurLimMaxs) Ifc() interface{}     { return x.TmesTmes() }
func (x AnaPrfmsDurLimMaxs) TmesTmes() *tmes.Tmes { return x.X.AnaPrfms().DurLimMaxs() }
func (x AnaPrfmsDayCnts) Act()                    { x.UntsUnts() }
func (x AnaPrfmsDayCnts) Ifc() interface{}        { return x.UntsUnts() }
func (x AnaPrfmsDayCnts) UntsUnts() *unts.Unts    { return x.X.AnaPrfms().DayCnts() }
func (x AnaPrfmsTrdCnts) Act()                    { x.UntsUnts() }
func (x AnaPrfmsTrdCnts) Ifc() interface{}        { return x.UntsUnts() }
func (x AnaPrfmsTrdCnts) UntsUnts() *unts.Unts    { return x.X.AnaPrfms().TrdCnts() }
func (x AnaPrfmsTrdPcts) Act()                    { x.FltsFlts() }
func (x AnaPrfmsTrdPcts) Ifc() interface{}        { return x.FltsFlts() }
func (x AnaPrfmsTrdPcts) FltsFlts() *flts.Flts    { return x.X.AnaPrfms().TrdPcts() }
func (x AnaPrfmsCstTotUsds) Act()                 { x.FltsFlts() }
func (x AnaPrfmsCstTotUsds) Ifc() interface{}     { return x.FltsFlts() }
func (x AnaPrfmsCstTotUsds) FltsFlts() *flts.Flts { return x.X.AnaPrfms().CstTotUsds() }
func (x AnaPrfmsCstSpdUsds) Act()                 { x.FltsFlts() }
func (x AnaPrfmsCstSpdUsds) Ifc() interface{}     { return x.FltsFlts() }
func (x AnaPrfmsCstSpdUsds) FltsFlts() *flts.Flts { return x.X.AnaPrfms().CstSpdUsds() }
func (x AnaPrfmsCstComUsds) Act()                 { x.FltsFlts() }
func (x AnaPrfmsCstComUsds) Ifc() interface{}     { return x.FltsFlts() }
func (x AnaPrfmsCstComUsds) FltsFlts() *flts.Flts { return x.X.AnaPrfms().CstComUsds() }
func (x AnaPrfmsPths) Act()                       { x.StrsStrs() }
func (x AnaPrfmsPths) Ifc() interface{}           { return x.StrsStrs() }
func (x AnaPrfmsPths) StrsStrs() *strs.Strs       { return x.X.AnaPrfms().Pths() }
func (x HstPrvsCnt) Act()                         { x.UntUnt() }
func (x HstPrvsCnt) Ifc() interface{}             { return x.UntUnt() }
func (x HstPrvsCnt) UntUnt() unt.Unt              { return x.X.HstPrvs().Cnt() }
func (x HstPrvsCpy) Act()                         { x.HstPrvs() }
func (x HstPrvsCpy) Ifc() interface{}             { return x.HstPrvs() }
func (x HstPrvsCpy) HstPrvs() *hst.Prvs           { return x.X.HstPrvs().Cpy() }
func (x HstPrvsClr) Act()                         { x.HstPrvs() }
func (x HstPrvsClr) Ifc() interface{}             { return x.HstPrvs() }
func (x HstPrvsClr) HstPrvs() *hst.Prvs           { return x.X.HstPrvs().Clr() }
func (x HstPrvsRand) Act()                        { x.HstPrvs() }
func (x HstPrvsRand) Ifc() interface{}            { return x.HstPrvs() }
func (x HstPrvsRand) HstPrvs() *hst.Prvs          { return x.X.HstPrvs().Rand() }
func (x HstPrvsMrg) Act()                         { x.HstPrvs() }
func (x HstPrvsMrg) Ifc() interface{}             { return x.HstPrvs() }
func (x HstPrvsMrg) HstPrvs() *hst.Prvs {
	var i0 []*hst.Prvs
	for _, cur := range x.I0 {
		i0 = append(i0, cur.HstPrvs())
	}
	return x.X.HstPrvs().Mrg(i0...)
}
func (x HstPrvsPush) Act()             { x.HstPrvs() }
func (x HstPrvsPush) Ifc() interface{} { return x.HstPrvs() }
func (x HstPrvsPush) HstPrvs() *hst.Prvs {
	var i0 []hst.Prv
	for _, cur := range x.I0 {
		i0 = append(i0, cur.HstPrv())
	}
	return x.X.HstPrvs().Push(i0...)
}
func (x HstPrvsPop) Act()             { x.HstPrv() }
func (x HstPrvsPop) Ifc() interface{} { return x.HstPrv() }
func (x HstPrvsPop) HstPrv() hst.Prv  { return x.X.HstPrvs().Pop() }
func (x HstPrvsQue) Act()             { x.HstPrvs() }
func (x HstPrvsQue) Ifc() interface{} { return x.HstPrvs() }
func (x HstPrvsQue) HstPrvs() *hst.Prvs {
	var i0 []hst.Prv
	for _, cur := range x.I0 {
		i0 = append(i0, cur.HstPrv())
	}
	return x.X.HstPrvs().Que(i0...)
}
func (x HstPrvsDque) Act()                     { x.HstPrv() }
func (x HstPrvsDque) Ifc() interface{}         { return x.HstPrv() }
func (x HstPrvsDque) HstPrv() hst.Prv          { return x.X.HstPrvs().Dque() }
func (x HstPrvsIns) Act()                      { x.HstPrvs() }
func (x HstPrvsIns) Ifc() interface{}          { return x.HstPrvs() }
func (x HstPrvsIns) HstPrvs() *hst.Prvs        { return x.X.HstPrvs().Ins(x.I0.UntUnt(), x.I1.HstPrv()) }
func (x HstPrvsUpd) Act()                      { x.HstPrvs() }
func (x HstPrvsUpd) Ifc() interface{}          { return x.HstPrvs() }
func (x HstPrvsUpd) HstPrvs() *hst.Prvs        { return x.X.HstPrvs().Upd(x.I0.UntUnt(), x.I1.HstPrv()) }
func (x HstPrvsDel) Act()                      { x.HstPrv() }
func (x HstPrvsDel) Ifc() interface{}          { return x.HstPrv() }
func (x HstPrvsDel) HstPrv() hst.Prv           { return x.X.HstPrvs().Del(x.I0.UntUnt()) }
func (x HstPrvsAt) Act()                       { x.HstPrv() }
func (x HstPrvsAt) Ifc() interface{}           { return x.HstPrv() }
func (x HstPrvsAt) HstPrv() hst.Prv            { return x.X.HstPrvs().At(x.I0.UntUnt()) }
func (x HstPrvsIn) Act()                       { x.HstPrvs() }
func (x HstPrvsIn) Ifc() interface{}           { return x.HstPrvs() }
func (x HstPrvsIn) HstPrvs() *hst.Prvs         { return x.X.HstPrvs().In(x.I0.UntUnt(), x.I1.UntUnt()) }
func (x HstPrvsInBnd) Act()                    { x.HstPrvs() }
func (x HstPrvsInBnd) Ifc() interface{}        { return x.HstPrvs() }
func (x HstPrvsInBnd) HstPrvs() *hst.Prvs      { return x.X.HstPrvs().InBnd(x.I0.BndBnd()) }
func (x HstPrvsFrom) Act()                     { x.HstPrvs() }
func (x HstPrvsFrom) Ifc() interface{}         { return x.HstPrvs() }
func (x HstPrvsFrom) HstPrvs() *hst.Prvs       { return x.X.HstPrvs().From(x.I0.UntUnt()) }
func (x HstPrvsTo) Act()                       { x.HstPrvs() }
func (x HstPrvsTo) Ifc() interface{}           { return x.HstPrvs() }
func (x HstPrvsTo) HstPrvs() *hst.Prvs         { return x.X.HstPrvs().To(x.I0.UntUnt()) }
func (x HstPrvsFst) Act()                      { x.HstPrv() }
func (x HstPrvsFst) Ifc() interface{}          { return x.HstPrv() }
func (x HstPrvsFst) HstPrv() hst.Prv           { return x.X.HstPrvs().Fst() }
func (x HstPrvsMdl) Act()                      { x.HstPrv() }
func (x HstPrvsMdl) Ifc() interface{}          { return x.HstPrv() }
func (x HstPrvsMdl) HstPrv() hst.Prv           { return x.X.HstPrvs().Mdl() }
func (x HstPrvsLst) Act()                      { x.HstPrv() }
func (x HstPrvsLst) Ifc() interface{}          { return x.HstPrv() }
func (x HstPrvsLst) HstPrv() hst.Prv           { return x.X.HstPrvs().Lst() }
func (x HstPrvsFstIdx) Act()                   { x.UntUnt() }
func (x HstPrvsFstIdx) Ifc() interface{}       { return x.UntUnt() }
func (x HstPrvsFstIdx) UntUnt() unt.Unt        { return x.X.HstPrvs().FstIdx() }
func (x HstPrvsMdlIdx) Act()                   { x.UntUnt() }
func (x HstPrvsMdlIdx) Ifc() interface{}       { return x.UntUnt() }
func (x HstPrvsMdlIdx) UntUnt() unt.Unt        { return x.X.HstPrvs().MdlIdx() }
func (x HstPrvsLstIdx) Act()                   { x.UntUnt() }
func (x HstPrvsLstIdx) Ifc() interface{}       { return x.UntUnt() }
func (x HstPrvsLstIdx) UntUnt() unt.Unt        { return x.X.HstPrvs().LstIdx() }
func (x HstPrvsRev) Act()                      { x.HstPrvs() }
func (x HstPrvsRev) Ifc() interface{}          { return x.HstPrvs() }
func (x HstPrvsRev) HstPrvs() *hst.Prvs        { return x.X.HstPrvs().Rev() }
func (x HstInstrsCnt) Act()                    { x.UntUnt() }
func (x HstInstrsCnt) Ifc() interface{}        { return x.UntUnt() }
func (x HstInstrsCnt) UntUnt() unt.Unt         { return x.X.HstInstrs().Cnt() }
func (x HstInstrsCpy) Act()                    { x.HstInstrs() }
func (x HstInstrsCpy) Ifc() interface{}        { return x.HstInstrs() }
func (x HstInstrsCpy) HstInstrs() *hst.Instrs  { return x.X.HstInstrs().Cpy() }
func (x HstInstrsClr) Act()                    { x.HstInstrs() }
func (x HstInstrsClr) Ifc() interface{}        { return x.HstInstrs() }
func (x HstInstrsClr) HstInstrs() *hst.Instrs  { return x.X.HstInstrs().Clr() }
func (x HstInstrsRand) Act()                   { x.HstInstrs() }
func (x HstInstrsRand) Ifc() interface{}       { return x.HstInstrs() }
func (x HstInstrsRand) HstInstrs() *hst.Instrs { return x.X.HstInstrs().Rand() }
func (x HstInstrsMrg) Act()                    { x.HstInstrs() }
func (x HstInstrsMrg) Ifc() interface{}        { return x.HstInstrs() }
func (x HstInstrsMrg) HstInstrs() *hst.Instrs {
	var i0 []*hst.Instrs
	for _, cur := range x.I0 {
		i0 = append(i0, cur.HstInstrs())
	}
	return x.X.HstInstrs().Mrg(i0...)
}
func (x HstInstrsPush) Act()             { x.HstInstrs() }
func (x HstInstrsPush) Ifc() interface{} { return x.HstInstrs() }
func (x HstInstrsPush) HstInstrs() *hst.Instrs {
	var i0 []hst.Instr
	for _, cur := range x.I0 {
		i0 = append(i0, cur.HstInstr())
	}
	return x.X.HstInstrs().Push(i0...)
}
func (x HstInstrsPop) Act()                { x.HstInstr() }
func (x HstInstrsPop) Ifc() interface{}    { return x.HstInstr() }
func (x HstInstrsPop) HstInstr() hst.Instr { return x.X.HstInstrs().Pop() }
func (x HstInstrsQue) Act()                { x.HstInstrs() }
func (x HstInstrsQue) Ifc() interface{}    { return x.HstInstrs() }
func (x HstInstrsQue) HstInstrs() *hst.Instrs {
	var i0 []hst.Instr
	for _, cur := range x.I0 {
		i0 = append(i0, cur.HstInstr())
	}
	return x.X.HstInstrs().Que(i0...)
}
func (x HstInstrsDque) Act()                { x.HstInstr() }
func (x HstInstrsDque) Ifc() interface{}    { return x.HstInstr() }
func (x HstInstrsDque) HstInstr() hst.Instr { return x.X.HstInstrs().Dque() }
func (x HstInstrsIns) Act()                 { x.HstInstrs() }
func (x HstInstrsIns) Ifc() interface{}     { return x.HstInstrs() }
func (x HstInstrsIns) HstInstrs() *hst.Instrs {
	return x.X.HstInstrs().Ins(x.I0.UntUnt(), x.I1.HstInstr())
}
func (x HstInstrsUpd) Act()             { x.HstInstrs() }
func (x HstInstrsUpd) Ifc() interface{} { return x.HstInstrs() }
func (x HstInstrsUpd) HstInstrs() *hst.Instrs {
	return x.X.HstInstrs().Upd(x.I0.UntUnt(), x.I1.HstInstr())
}
func (x HstInstrsDel) Act()                     { x.HstInstr() }
func (x HstInstrsDel) Ifc() interface{}         { return x.HstInstr() }
func (x HstInstrsDel) HstInstr() hst.Instr      { return x.X.HstInstrs().Del(x.I0.UntUnt()) }
func (x HstInstrsAt) Act()                      { x.HstInstr() }
func (x HstInstrsAt) Ifc() interface{}          { return x.HstInstr() }
func (x HstInstrsAt) HstInstr() hst.Instr       { return x.X.HstInstrs().At(x.I0.UntUnt()) }
func (x HstInstrsIn) Act()                      { x.HstInstrs() }
func (x HstInstrsIn) Ifc() interface{}          { return x.HstInstrs() }
func (x HstInstrsIn) HstInstrs() *hst.Instrs    { return x.X.HstInstrs().In(x.I0.UntUnt(), x.I1.UntUnt()) }
func (x HstInstrsInBnd) Act()                   { x.HstInstrs() }
func (x HstInstrsInBnd) Ifc() interface{}       { return x.HstInstrs() }
func (x HstInstrsInBnd) HstInstrs() *hst.Instrs { return x.X.HstInstrs().InBnd(x.I0.BndBnd()) }
func (x HstInstrsFrom) Act()                    { x.HstInstrs() }
func (x HstInstrsFrom) Ifc() interface{}        { return x.HstInstrs() }
func (x HstInstrsFrom) HstInstrs() *hst.Instrs  { return x.X.HstInstrs().From(x.I0.UntUnt()) }
func (x HstInstrsTo) Act()                      { x.HstInstrs() }
func (x HstInstrsTo) Ifc() interface{}          { return x.HstInstrs() }
func (x HstInstrsTo) HstInstrs() *hst.Instrs    { return x.X.HstInstrs().To(x.I0.UntUnt()) }
func (x HstInstrsFst) Act()                     { x.HstInstr() }
func (x HstInstrsFst) Ifc() interface{}         { return x.HstInstr() }
func (x HstInstrsFst) HstInstr() hst.Instr      { return x.X.HstInstrs().Fst() }
func (x HstInstrsMdl) Act()                     { x.HstInstr() }
func (x HstInstrsMdl) Ifc() interface{}         { return x.HstInstr() }
func (x HstInstrsMdl) HstInstr() hst.Instr      { return x.X.HstInstrs().Mdl() }
func (x HstInstrsLst) Act()                     { x.HstInstr() }
func (x HstInstrsLst) Ifc() interface{}         { return x.HstInstr() }
func (x HstInstrsLst) HstInstr() hst.Instr      { return x.X.HstInstrs().Lst() }
func (x HstInstrsFstIdx) Act()                  { x.UntUnt() }
func (x HstInstrsFstIdx) Ifc() interface{}      { return x.UntUnt() }
func (x HstInstrsFstIdx) UntUnt() unt.Unt       { return x.X.HstInstrs().FstIdx() }
func (x HstInstrsMdlIdx) Act()                  { x.UntUnt() }
func (x HstInstrsMdlIdx) Ifc() interface{}      { return x.UntUnt() }
func (x HstInstrsMdlIdx) UntUnt() unt.Unt       { return x.X.HstInstrs().MdlIdx() }
func (x HstInstrsLstIdx) Act()                  { x.UntUnt() }
func (x HstInstrsLstIdx) Ifc() interface{}      { return x.UntUnt() }
func (x HstInstrsLstIdx) UntUnt() unt.Unt       { return x.X.HstInstrs().LstIdx() }
func (x HstInstrsRev) Act()                     { x.HstInstrs() }
func (x HstInstrsRev) Ifc() interface{}         { return x.HstInstrs() }
func (x HstInstrsRev) HstInstrs() *hst.Instrs   { return x.X.HstInstrs().Rev() }
func (x HstInrvlsCnt) Act()                     { x.UntUnt() }
func (x HstInrvlsCnt) Ifc() interface{}         { return x.UntUnt() }
func (x HstInrvlsCnt) UntUnt() unt.Unt          { return x.X.HstInrvls().Cnt() }
func (x HstInrvlsCpy) Act()                     { x.HstInrvls() }
func (x HstInrvlsCpy) Ifc() interface{}         { return x.HstInrvls() }
func (x HstInrvlsCpy) HstInrvls() *hst.Inrvls   { return x.X.HstInrvls().Cpy() }
func (x HstInrvlsClr) Act()                     { x.HstInrvls() }
func (x HstInrvlsClr) Ifc() interface{}         { return x.HstInrvls() }
func (x HstInrvlsClr) HstInrvls() *hst.Inrvls   { return x.X.HstInrvls().Clr() }
func (x HstInrvlsRand) Act()                    { x.HstInrvls() }
func (x HstInrvlsRand) Ifc() interface{}        { return x.HstInrvls() }
func (x HstInrvlsRand) HstInrvls() *hst.Inrvls  { return x.X.HstInrvls().Rand() }
func (x HstInrvlsMrg) Act()                     { x.HstInrvls() }
func (x HstInrvlsMrg) Ifc() interface{}         { return x.HstInrvls() }
func (x HstInrvlsMrg) HstInrvls() *hst.Inrvls {
	var i0 []*hst.Inrvls
	for _, cur := range x.I0 {
		i0 = append(i0, cur.HstInrvls())
	}
	return x.X.HstInrvls().Mrg(i0...)
}
func (x HstInrvlsPush) Act()             { x.HstInrvls() }
func (x HstInrvlsPush) Ifc() interface{} { return x.HstInrvls() }
func (x HstInrvlsPush) HstInrvls() *hst.Inrvls {
	var i0 []hst.Inrvl
	for _, cur := range x.I0 {
		i0 = append(i0, cur.HstInrvl())
	}
	return x.X.HstInrvls().Push(i0...)
}
func (x HstInrvlsPop) Act()                { x.HstInrvl() }
func (x HstInrvlsPop) Ifc() interface{}    { return x.HstInrvl() }
func (x HstInrvlsPop) HstInrvl() hst.Inrvl { return x.X.HstInrvls().Pop() }
func (x HstInrvlsQue) Act()                { x.HstInrvls() }
func (x HstInrvlsQue) Ifc() interface{}    { return x.HstInrvls() }
func (x HstInrvlsQue) HstInrvls() *hst.Inrvls {
	var i0 []hst.Inrvl
	for _, cur := range x.I0 {
		i0 = append(i0, cur.HstInrvl())
	}
	return x.X.HstInrvls().Que(i0...)
}
func (x HstInrvlsDque) Act()                { x.HstInrvl() }
func (x HstInrvlsDque) Ifc() interface{}    { return x.HstInrvl() }
func (x HstInrvlsDque) HstInrvl() hst.Inrvl { return x.X.HstInrvls().Dque() }
func (x HstInrvlsIns) Act()                 { x.HstInrvls() }
func (x HstInrvlsIns) Ifc() interface{}     { return x.HstInrvls() }
func (x HstInrvlsIns) HstInrvls() *hst.Inrvls {
	return x.X.HstInrvls().Ins(x.I0.UntUnt(), x.I1.HstInrvl())
}
func (x HstInrvlsUpd) Act()             { x.HstInrvls() }
func (x HstInrvlsUpd) Ifc() interface{} { return x.HstInrvls() }
func (x HstInrvlsUpd) HstInrvls() *hst.Inrvls {
	return x.X.HstInrvls().Upd(x.I0.UntUnt(), x.I1.HstInrvl())
}
func (x HstInrvlsDel) Act()                     { x.HstInrvl() }
func (x HstInrvlsDel) Ifc() interface{}         { return x.HstInrvl() }
func (x HstInrvlsDel) HstInrvl() hst.Inrvl      { return x.X.HstInrvls().Del(x.I0.UntUnt()) }
func (x HstInrvlsAt) Act()                      { x.HstInrvl() }
func (x HstInrvlsAt) Ifc() interface{}          { return x.HstInrvl() }
func (x HstInrvlsAt) HstInrvl() hst.Inrvl       { return x.X.HstInrvls().At(x.I0.UntUnt()) }
func (x HstInrvlsIn) Act()                      { x.HstInrvls() }
func (x HstInrvlsIn) Ifc() interface{}          { return x.HstInrvls() }
func (x HstInrvlsIn) HstInrvls() *hst.Inrvls    { return x.X.HstInrvls().In(x.I0.UntUnt(), x.I1.UntUnt()) }
func (x HstInrvlsInBnd) Act()                   { x.HstInrvls() }
func (x HstInrvlsInBnd) Ifc() interface{}       { return x.HstInrvls() }
func (x HstInrvlsInBnd) HstInrvls() *hst.Inrvls { return x.X.HstInrvls().InBnd(x.I0.BndBnd()) }
func (x HstInrvlsFrom) Act()                    { x.HstInrvls() }
func (x HstInrvlsFrom) Ifc() interface{}        { return x.HstInrvls() }
func (x HstInrvlsFrom) HstInrvls() *hst.Inrvls  { return x.X.HstInrvls().From(x.I0.UntUnt()) }
func (x HstInrvlsTo) Act()                      { x.HstInrvls() }
func (x HstInrvlsTo) Ifc() interface{}          { return x.HstInrvls() }
func (x HstInrvlsTo) HstInrvls() *hst.Inrvls    { return x.X.HstInrvls().To(x.I0.UntUnt()) }
func (x HstInrvlsFst) Act()                     { x.HstInrvl() }
func (x HstInrvlsFst) Ifc() interface{}         { return x.HstInrvl() }
func (x HstInrvlsFst) HstInrvl() hst.Inrvl      { return x.X.HstInrvls().Fst() }
func (x HstInrvlsMdl) Act()                     { x.HstInrvl() }
func (x HstInrvlsMdl) Ifc() interface{}         { return x.HstInrvl() }
func (x HstInrvlsMdl) HstInrvl() hst.Inrvl      { return x.X.HstInrvls().Mdl() }
func (x HstInrvlsLst) Act()                     { x.HstInrvl() }
func (x HstInrvlsLst) Ifc() interface{}         { return x.HstInrvl() }
func (x HstInrvlsLst) HstInrvl() hst.Inrvl      { return x.X.HstInrvls().Lst() }
func (x HstInrvlsFstIdx) Act()                  { x.UntUnt() }
func (x HstInrvlsFstIdx) Ifc() interface{}      { return x.UntUnt() }
func (x HstInrvlsFstIdx) UntUnt() unt.Unt       { return x.X.HstInrvls().FstIdx() }
func (x HstInrvlsMdlIdx) Act()                  { x.UntUnt() }
func (x HstInrvlsMdlIdx) Ifc() interface{}      { return x.UntUnt() }
func (x HstInrvlsMdlIdx) UntUnt() unt.Unt       { return x.X.HstInrvls().MdlIdx() }
func (x HstInrvlsLstIdx) Act()                  { x.UntUnt() }
func (x HstInrvlsLstIdx) Ifc() interface{}      { return x.UntUnt() }
func (x HstInrvlsLstIdx) UntUnt() unt.Unt       { return x.X.HstInrvls().LstIdx() }
func (x HstInrvlsRev) Act()                     { x.HstInrvls() }
func (x HstInrvlsRev) Ifc() interface{}         { return x.HstInrvls() }
func (x HstInrvlsRev) HstInrvls() *hst.Inrvls   { return x.X.HstInrvls().Rev() }
func (x HstSidesCnt) Act()                      { x.UntUnt() }
func (x HstSidesCnt) Ifc() interface{}          { return x.UntUnt() }
func (x HstSidesCnt) UntUnt() unt.Unt           { return x.X.HstSides().Cnt() }
func (x HstSidesCpy) Act()                      { x.HstSides() }
func (x HstSidesCpy) Ifc() interface{}          { return x.HstSides() }
func (x HstSidesCpy) HstSides() *hst.Sides      { return x.X.HstSides().Cpy() }
func (x HstSidesClr) Act()                      { x.HstSides() }
func (x HstSidesClr) Ifc() interface{}          { return x.HstSides() }
func (x HstSidesClr) HstSides() *hst.Sides      { return x.X.HstSides().Clr() }
func (x HstSidesRand) Act()                     { x.HstSides() }
func (x HstSidesRand) Ifc() interface{}         { return x.HstSides() }
func (x HstSidesRand) HstSides() *hst.Sides     { return x.X.HstSides().Rand() }
func (x HstSidesMrg) Act()                      { x.HstSides() }
func (x HstSidesMrg) Ifc() interface{}          { return x.HstSides() }
func (x HstSidesMrg) HstSides() *hst.Sides {
	var i0 []*hst.Sides
	for _, cur := range x.I0 {
		i0 = append(i0, cur.HstSides())
	}
	return x.X.HstSides().Mrg(i0...)
}
func (x HstSidesPush) Act()             { x.HstSides() }
func (x HstSidesPush) Ifc() interface{} { return x.HstSides() }
func (x HstSidesPush) HstSides() *hst.Sides {
	var i0 []hst.Side
	for _, cur := range x.I0 {
		i0 = append(i0, cur.HstSide())
	}
	return x.X.HstSides().Push(i0...)
}
func (x HstSidesPop) Act()              { x.HstSide() }
func (x HstSidesPop) Ifc() interface{}  { return x.HstSide() }
func (x HstSidesPop) HstSide() hst.Side { return x.X.HstSides().Pop() }
func (x HstSidesQue) Act()              { x.HstSides() }
func (x HstSidesQue) Ifc() interface{}  { return x.HstSides() }
func (x HstSidesQue) HstSides() *hst.Sides {
	var i0 []hst.Side
	for _, cur := range x.I0 {
		i0 = append(i0, cur.HstSide())
	}
	return x.X.HstSides().Que(i0...)
}
func (x HstSidesDque) Act()                  { x.HstSide() }
func (x HstSidesDque) Ifc() interface{}      { return x.HstSide() }
func (x HstSidesDque) HstSide() hst.Side     { return x.X.HstSides().Dque() }
func (x HstSidesIns) Act()                   { x.HstSides() }
func (x HstSidesIns) Ifc() interface{}       { return x.HstSides() }
func (x HstSidesIns) HstSides() *hst.Sides   { return x.X.HstSides().Ins(x.I0.UntUnt(), x.I1.HstSide()) }
func (x HstSidesUpd) Act()                   { x.HstSides() }
func (x HstSidesUpd) Ifc() interface{}       { return x.HstSides() }
func (x HstSidesUpd) HstSides() *hst.Sides   { return x.X.HstSides().Upd(x.I0.UntUnt(), x.I1.HstSide()) }
func (x HstSidesDel) Act()                   { x.HstSide() }
func (x HstSidesDel) Ifc() interface{}       { return x.HstSide() }
func (x HstSidesDel) HstSide() hst.Side      { return x.X.HstSides().Del(x.I0.UntUnt()) }
func (x HstSidesAt) Act()                    { x.HstSide() }
func (x HstSidesAt) Ifc() interface{}        { return x.HstSide() }
func (x HstSidesAt) HstSide() hst.Side       { return x.X.HstSides().At(x.I0.UntUnt()) }
func (x HstSidesIn) Act()                    { x.HstSides() }
func (x HstSidesIn) Ifc() interface{}        { return x.HstSides() }
func (x HstSidesIn) HstSides() *hst.Sides    { return x.X.HstSides().In(x.I0.UntUnt(), x.I1.UntUnt()) }
func (x HstSidesInBnd) Act()                 { x.HstSides() }
func (x HstSidesInBnd) Ifc() interface{}     { return x.HstSides() }
func (x HstSidesInBnd) HstSides() *hst.Sides { return x.X.HstSides().InBnd(x.I0.BndBnd()) }
func (x HstSidesFrom) Act()                  { x.HstSides() }
func (x HstSidesFrom) Ifc() interface{}      { return x.HstSides() }
func (x HstSidesFrom) HstSides() *hst.Sides  { return x.X.HstSides().From(x.I0.UntUnt()) }
func (x HstSidesTo) Act()                    { x.HstSides() }
func (x HstSidesTo) Ifc() interface{}        { return x.HstSides() }
func (x HstSidesTo) HstSides() *hst.Sides    { return x.X.HstSides().To(x.I0.UntUnt()) }
func (x HstSidesFst) Act()                   { x.HstSide() }
func (x HstSidesFst) Ifc() interface{}       { return x.HstSide() }
func (x HstSidesFst) HstSide() hst.Side      { return x.X.HstSides().Fst() }
func (x HstSidesMdl) Act()                   { x.HstSide() }
func (x HstSidesMdl) Ifc() interface{}       { return x.HstSide() }
func (x HstSidesMdl) HstSide() hst.Side      { return x.X.HstSides().Mdl() }
func (x HstSidesLst) Act()                   { x.HstSide() }
func (x HstSidesLst) Ifc() interface{}       { return x.HstSide() }
func (x HstSidesLst) HstSide() hst.Side      { return x.X.HstSides().Lst() }
func (x HstSidesFstIdx) Act()                { x.UntUnt() }
func (x HstSidesFstIdx) Ifc() interface{}    { return x.UntUnt() }
func (x HstSidesFstIdx) UntUnt() unt.Unt     { return x.X.HstSides().FstIdx() }
func (x HstSidesMdlIdx) Act()                { x.UntUnt() }
func (x HstSidesMdlIdx) Ifc() interface{}    { return x.UntUnt() }
func (x HstSidesMdlIdx) UntUnt() unt.Unt     { return x.X.HstSides().MdlIdx() }
func (x HstSidesLstIdx) Act()                { x.UntUnt() }
func (x HstSidesLstIdx) Ifc() interface{}    { return x.UntUnt() }
func (x HstSidesLstIdx) UntUnt() unt.Unt     { return x.X.HstSides().LstIdx() }
func (x HstSidesRev) Act()                   { x.HstSides() }
func (x HstSidesRev) Ifc() interface{}       { return x.HstSides() }
func (x HstSidesRev) HstSides() *hst.Sides   { return x.X.HstSides().Rev() }
func (x HstStmsCnt) Act()                    { x.UntUnt() }
func (x HstStmsCnt) Ifc() interface{}        { return x.UntUnt() }
func (x HstStmsCnt) UntUnt() unt.Unt         { return x.X.HstStms().Cnt() }
func (x HstStmsCpy) Act()                    { x.HstStms() }
func (x HstStmsCpy) Ifc() interface{}        { return x.HstStms() }
func (x HstStmsCpy) HstStms() *hst.Stms      { return x.X.HstStms().Cpy() }
func (x HstStmsClr) Act()                    { x.HstStms() }
func (x HstStmsClr) Ifc() interface{}        { return x.HstStms() }
func (x HstStmsClr) HstStms() *hst.Stms      { return x.X.HstStms().Clr() }
func (x HstStmsRand) Act()                   { x.HstStms() }
func (x HstStmsRand) Ifc() interface{}       { return x.HstStms() }
func (x HstStmsRand) HstStms() *hst.Stms     { return x.X.HstStms().Rand() }
func (x HstStmsMrg) Act()                    { x.HstStms() }
func (x HstStmsMrg) Ifc() interface{}        { return x.HstStms() }
func (x HstStmsMrg) HstStms() *hst.Stms {
	var i0 []*hst.Stms
	for _, cur := range x.I0 {
		i0 = append(i0, cur.HstStms())
	}
	return x.X.HstStms().Mrg(i0...)
}
func (x HstStmsPush) Act()             { x.HstStms() }
func (x HstStmsPush) Ifc() interface{} { return x.HstStms() }
func (x HstStmsPush) HstStms() *hst.Stms {
	var i0 []hst.Stm
	for _, cur := range x.I0 {
		i0 = append(i0, cur.HstStm())
	}
	return x.X.HstStms().Push(i0...)
}
func (x HstStmsPop) Act()             { x.HstStm() }
func (x HstStmsPop) Ifc() interface{} { return x.HstStm() }
func (x HstStmsPop) HstStm() hst.Stm  { return x.X.HstStms().Pop() }
func (x HstStmsQue) Act()             { x.HstStms() }
func (x HstStmsQue) Ifc() interface{} { return x.HstStms() }
func (x HstStmsQue) HstStms() *hst.Stms {
	var i0 []hst.Stm
	for _, cur := range x.I0 {
		i0 = append(i0, cur.HstStm())
	}
	return x.X.HstStms().Que(i0...)
}
func (x HstStmsDque) Act()                { x.HstStm() }
func (x HstStmsDque) Ifc() interface{}    { return x.HstStm() }
func (x HstStmsDque) HstStm() hst.Stm     { return x.X.HstStms().Dque() }
func (x HstStmsIns) Act()                 { x.HstStms() }
func (x HstStmsIns) Ifc() interface{}     { return x.HstStms() }
func (x HstStmsIns) HstStms() *hst.Stms   { return x.X.HstStms().Ins(x.I0.UntUnt(), x.I1.HstStm()) }
func (x HstStmsUpd) Act()                 { x.HstStms() }
func (x HstStmsUpd) Ifc() interface{}     { return x.HstStms() }
func (x HstStmsUpd) HstStms() *hst.Stms   { return x.X.HstStms().Upd(x.I0.UntUnt(), x.I1.HstStm()) }
func (x HstStmsDel) Act()                 { x.HstStm() }
func (x HstStmsDel) Ifc() interface{}     { return x.HstStm() }
func (x HstStmsDel) HstStm() hst.Stm      { return x.X.HstStms().Del(x.I0.UntUnt()) }
func (x HstStmsAt) Act()                  { x.HstStm() }
func (x HstStmsAt) Ifc() interface{}      { return x.HstStm() }
func (x HstStmsAt) HstStm() hst.Stm       { return x.X.HstStms().At(x.I0.UntUnt()) }
func (x HstStmsIn) Act()                  { x.HstStms() }
func (x HstStmsIn) Ifc() interface{}      { return x.HstStms() }
func (x HstStmsIn) HstStms() *hst.Stms    { return x.X.HstStms().In(x.I0.UntUnt(), x.I1.UntUnt()) }
func (x HstStmsInBnd) Act()               { x.HstStms() }
func (x HstStmsInBnd) Ifc() interface{}   { return x.HstStms() }
func (x HstStmsInBnd) HstStms() *hst.Stms { return x.X.HstStms().InBnd(x.I0.BndBnd()) }
func (x HstStmsFrom) Act()                { x.HstStms() }
func (x HstStmsFrom) Ifc() interface{}    { return x.HstStms() }
func (x HstStmsFrom) HstStms() *hst.Stms  { return x.X.HstStms().From(x.I0.UntUnt()) }
func (x HstStmsTo) Act()                  { x.HstStms() }
func (x HstStmsTo) Ifc() interface{}      { return x.HstStms() }
func (x HstStmsTo) HstStms() *hst.Stms    { return x.X.HstStms().To(x.I0.UntUnt()) }
func (x HstStmsFst) Act()                 { x.HstStm() }
func (x HstStmsFst) Ifc() interface{}     { return x.HstStm() }
func (x HstStmsFst) HstStm() hst.Stm      { return x.X.HstStms().Fst() }
func (x HstStmsMdl) Act()                 { x.HstStm() }
func (x HstStmsMdl) Ifc() interface{}     { return x.HstStm() }
func (x HstStmsMdl) HstStm() hst.Stm      { return x.X.HstStms().Mdl() }
func (x HstStmsLst) Act()                 { x.HstStm() }
func (x HstStmsLst) Ifc() interface{}     { return x.HstStm() }
func (x HstStmsLst) HstStm() hst.Stm      { return x.X.HstStms().Lst() }
func (x HstStmsFstIdx) Act()              { x.UntUnt() }
func (x HstStmsFstIdx) Ifc() interface{}  { return x.UntUnt() }
func (x HstStmsFstIdx) UntUnt() unt.Unt   { return x.X.HstStms().FstIdx() }
func (x HstStmsMdlIdx) Act()              { x.UntUnt() }
func (x HstStmsMdlIdx) Ifc() interface{}  { return x.UntUnt() }
func (x HstStmsMdlIdx) UntUnt() unt.Unt   { return x.X.HstStms().MdlIdx() }
func (x HstStmsLstIdx) Act()              { x.UntUnt() }
func (x HstStmsLstIdx) Ifc() interface{}  { return x.UntUnt() }
func (x HstStmsLstIdx) UntUnt() unt.Unt   { return x.X.HstStms().LstIdx() }
func (x HstStmsRev) Act()                 { x.HstStms() }
func (x HstStmsRev) Ifc() interface{}     { return x.HstStms() }
func (x HstStmsRev) HstStms() *hst.Stms   { return x.X.HstStms().Rev() }
func (x HstCndsCnt) Act()                 { x.UntUnt() }
func (x HstCndsCnt) Ifc() interface{}     { return x.UntUnt() }
func (x HstCndsCnt) UntUnt() unt.Unt      { return x.X.HstCnds().Cnt() }
func (x HstCndsCpy) Act()                 { x.HstCnds() }
func (x HstCndsCpy) Ifc() interface{}     { return x.HstCnds() }
func (x HstCndsCpy) HstCnds() *hst.Cnds   { return x.X.HstCnds().Cpy() }
func (x HstCndsClr) Act()                 { x.HstCnds() }
func (x HstCndsClr) Ifc() interface{}     { return x.HstCnds() }
func (x HstCndsClr) HstCnds() *hst.Cnds   { return x.X.HstCnds().Clr() }
func (x HstCndsRand) Act()                { x.HstCnds() }
func (x HstCndsRand) Ifc() interface{}    { return x.HstCnds() }
func (x HstCndsRand) HstCnds() *hst.Cnds  { return x.X.HstCnds().Rand() }
func (x HstCndsMrg) Act()                 { x.HstCnds() }
func (x HstCndsMrg) Ifc() interface{}     { return x.HstCnds() }
func (x HstCndsMrg) HstCnds() *hst.Cnds {
	var i0 []*hst.Cnds
	for _, cur := range x.I0 {
		i0 = append(i0, cur.HstCnds())
	}
	return x.X.HstCnds().Mrg(i0...)
}
func (x HstCndsPush) Act()             { x.HstCnds() }
func (x HstCndsPush) Ifc() interface{} { return x.HstCnds() }
func (x HstCndsPush) HstCnds() *hst.Cnds {
	var i0 []hst.Cnd
	for _, cur := range x.I0 {
		i0 = append(i0, cur.HstCnd())
	}
	return x.X.HstCnds().Push(i0...)
}
func (x HstCndsPop) Act()             { x.HstCnd() }
func (x HstCndsPop) Ifc() interface{} { return x.HstCnd() }
func (x HstCndsPop) HstCnd() hst.Cnd  { return x.X.HstCnds().Pop() }
func (x HstCndsQue) Act()             { x.HstCnds() }
func (x HstCndsQue) Ifc() interface{} { return x.HstCnds() }
func (x HstCndsQue) HstCnds() *hst.Cnds {
	var i0 []hst.Cnd
	for _, cur := range x.I0 {
		i0 = append(i0, cur.HstCnd())
	}
	return x.X.HstCnds().Que(i0...)
}
func (x HstCndsDque) Act()                  { x.HstCnd() }
func (x HstCndsDque) Ifc() interface{}      { return x.HstCnd() }
func (x HstCndsDque) HstCnd() hst.Cnd       { return x.X.HstCnds().Dque() }
func (x HstCndsIns) Act()                   { x.HstCnds() }
func (x HstCndsIns) Ifc() interface{}       { return x.HstCnds() }
func (x HstCndsIns) HstCnds() *hst.Cnds     { return x.X.HstCnds().Ins(x.I0.UntUnt(), x.I1.HstCnd()) }
func (x HstCndsUpd) Act()                   { x.HstCnds() }
func (x HstCndsUpd) Ifc() interface{}       { return x.HstCnds() }
func (x HstCndsUpd) HstCnds() *hst.Cnds     { return x.X.HstCnds().Upd(x.I0.UntUnt(), x.I1.HstCnd()) }
func (x HstCndsDel) Act()                   { x.HstCnd() }
func (x HstCndsDel) Ifc() interface{}       { return x.HstCnd() }
func (x HstCndsDel) HstCnd() hst.Cnd        { return x.X.HstCnds().Del(x.I0.UntUnt()) }
func (x HstCndsAt) Act()                    { x.HstCnd() }
func (x HstCndsAt) Ifc() interface{}        { return x.HstCnd() }
func (x HstCndsAt) HstCnd() hst.Cnd         { return x.X.HstCnds().At(x.I0.UntUnt()) }
func (x HstCndsIn) Act()                    { x.HstCnds() }
func (x HstCndsIn) Ifc() interface{}        { return x.HstCnds() }
func (x HstCndsIn) HstCnds() *hst.Cnds      { return x.X.HstCnds().In(x.I0.UntUnt(), x.I1.UntUnt()) }
func (x HstCndsInBnd) Act()                 { x.HstCnds() }
func (x HstCndsInBnd) Ifc() interface{}     { return x.HstCnds() }
func (x HstCndsInBnd) HstCnds() *hst.Cnds   { return x.X.HstCnds().InBnd(x.I0.BndBnd()) }
func (x HstCndsFrom) Act()                  { x.HstCnds() }
func (x HstCndsFrom) Ifc() interface{}      { return x.HstCnds() }
func (x HstCndsFrom) HstCnds() *hst.Cnds    { return x.X.HstCnds().From(x.I0.UntUnt()) }
func (x HstCndsTo) Act()                    { x.HstCnds() }
func (x HstCndsTo) Ifc() interface{}        { return x.HstCnds() }
func (x HstCndsTo) HstCnds() *hst.Cnds      { return x.X.HstCnds().To(x.I0.UntUnt()) }
func (x HstCndsFst) Act()                   { x.HstCnd() }
func (x HstCndsFst) Ifc() interface{}       { return x.HstCnd() }
func (x HstCndsFst) HstCnd() hst.Cnd        { return x.X.HstCnds().Fst() }
func (x HstCndsMdl) Act()                   { x.HstCnd() }
func (x HstCndsMdl) Ifc() interface{}       { return x.HstCnd() }
func (x HstCndsMdl) HstCnd() hst.Cnd        { return x.X.HstCnds().Mdl() }
func (x HstCndsLst) Act()                   { x.HstCnd() }
func (x HstCndsLst) Ifc() interface{}       { return x.HstCnd() }
func (x HstCndsLst) HstCnd() hst.Cnd        { return x.X.HstCnds().Lst() }
func (x HstCndsFstIdx) Act()                { x.UntUnt() }
func (x HstCndsFstIdx) Ifc() interface{}    { return x.UntUnt() }
func (x HstCndsFstIdx) UntUnt() unt.Unt     { return x.X.HstCnds().FstIdx() }
func (x HstCndsMdlIdx) Act()                { x.UntUnt() }
func (x HstCndsMdlIdx) Ifc() interface{}    { return x.UntUnt() }
func (x HstCndsMdlIdx) UntUnt() unt.Unt     { return x.X.HstCnds().MdlIdx() }
func (x HstCndsLstIdx) Act()                { x.UntUnt() }
func (x HstCndsLstIdx) Ifc() interface{}    { return x.UntUnt() }
func (x HstCndsLstIdx) UntUnt() unt.Unt     { return x.X.HstCnds().LstIdx() }
func (x HstCndsRev) Act()                   { x.HstCnds() }
func (x HstCndsRev) Ifc() interface{}       { return x.HstCnds() }
func (x HstCndsRev) HstCnds() *hst.Cnds     { return x.X.HstCnds().Rev() }
func (x HstStgysCnt) Act()                  { x.UntUnt() }
func (x HstStgysCnt) Ifc() interface{}      { return x.UntUnt() }
func (x HstStgysCnt) UntUnt() unt.Unt       { return x.X.HstStgys().Cnt() }
func (x HstStgysCpy) Act()                  { x.HstStgys() }
func (x HstStgysCpy) Ifc() interface{}      { return x.HstStgys() }
func (x HstStgysCpy) HstStgys() *hst.Stgys  { return x.X.HstStgys().Cpy() }
func (x HstStgysClr) Act()                  { x.HstStgys() }
func (x HstStgysClr) Ifc() interface{}      { return x.HstStgys() }
func (x HstStgysClr) HstStgys() *hst.Stgys  { return x.X.HstStgys().Clr() }
func (x HstStgysRand) Act()                 { x.HstStgys() }
func (x HstStgysRand) Ifc() interface{}     { return x.HstStgys() }
func (x HstStgysRand) HstStgys() *hst.Stgys { return x.X.HstStgys().Rand() }
func (x HstStgysMrg) Act()                  { x.HstStgys() }
func (x HstStgysMrg) Ifc() interface{}      { return x.HstStgys() }
func (x HstStgysMrg) HstStgys() *hst.Stgys {
	var i0 []*hst.Stgys
	for _, cur := range x.I0 {
		i0 = append(i0, cur.HstStgys())
	}
	return x.X.HstStgys().Mrg(i0...)
}
func (x HstStgysPush) Act()             { x.HstStgys() }
func (x HstStgysPush) Ifc() interface{} { return x.HstStgys() }
func (x HstStgysPush) HstStgys() *hst.Stgys {
	var i0 []hst.Stgy
	for _, cur := range x.I0 {
		i0 = append(i0, cur.HstStgy())
	}
	return x.X.HstStgys().Push(i0...)
}
func (x HstStgysPop) Act()              { x.HstStgy() }
func (x HstStgysPop) Ifc() interface{}  { return x.HstStgy() }
func (x HstStgysPop) HstStgy() hst.Stgy { return x.X.HstStgys().Pop() }
func (x HstStgysQue) Act()              { x.HstStgys() }
func (x HstStgysQue) Ifc() interface{}  { return x.HstStgys() }
func (x HstStgysQue) HstStgys() *hst.Stgys {
	var i0 []hst.Stgy
	for _, cur := range x.I0 {
		i0 = append(i0, cur.HstStgy())
	}
	return x.X.HstStgys().Que(i0...)
}
func (x HstStgysDque) Act()                  { x.HstStgy() }
func (x HstStgysDque) Ifc() interface{}      { return x.HstStgy() }
func (x HstStgysDque) HstStgy() hst.Stgy     { return x.X.HstStgys().Dque() }
func (x HstStgysIns) Act()                   { x.HstStgys() }
func (x HstStgysIns) Ifc() interface{}       { return x.HstStgys() }
func (x HstStgysIns) HstStgys() *hst.Stgys   { return x.X.HstStgys().Ins(x.I0.UntUnt(), x.I1.HstStgy()) }
func (x HstStgysUpd) Act()                   { x.HstStgys() }
func (x HstStgysUpd) Ifc() interface{}       { return x.HstStgys() }
func (x HstStgysUpd) HstStgys() *hst.Stgys   { return x.X.HstStgys().Upd(x.I0.UntUnt(), x.I1.HstStgy()) }
func (x HstStgysDel) Act()                   { x.HstStgy() }
func (x HstStgysDel) Ifc() interface{}       { return x.HstStgy() }
func (x HstStgysDel) HstStgy() hst.Stgy      { return x.X.HstStgys().Del(x.I0.UntUnt()) }
func (x HstStgysAt) Act()                    { x.HstStgy() }
func (x HstStgysAt) Ifc() interface{}        { return x.HstStgy() }
func (x HstStgysAt) HstStgy() hst.Stgy       { return x.X.HstStgys().At(x.I0.UntUnt()) }
func (x HstStgysIn) Act()                    { x.HstStgys() }
func (x HstStgysIn) Ifc() interface{}        { return x.HstStgys() }
func (x HstStgysIn) HstStgys() *hst.Stgys    { return x.X.HstStgys().In(x.I0.UntUnt(), x.I1.UntUnt()) }
func (x HstStgysInBnd) Act()                 { x.HstStgys() }
func (x HstStgysInBnd) Ifc() interface{}     { return x.HstStgys() }
func (x HstStgysInBnd) HstStgys() *hst.Stgys { return x.X.HstStgys().InBnd(x.I0.BndBnd()) }
func (x HstStgysFrom) Act()                  { x.HstStgys() }
func (x HstStgysFrom) Ifc() interface{}      { return x.HstStgys() }
func (x HstStgysFrom) HstStgys() *hst.Stgys  { return x.X.HstStgys().From(x.I0.UntUnt()) }
func (x HstStgysTo) Act()                    { x.HstStgys() }
func (x HstStgysTo) Ifc() interface{}        { return x.HstStgys() }
func (x HstStgysTo) HstStgys() *hst.Stgys    { return x.X.HstStgys().To(x.I0.UntUnt()) }
func (x HstStgysFst) Act()                   { x.HstStgy() }
func (x HstStgysFst) Ifc() interface{}       { return x.HstStgy() }
func (x HstStgysFst) HstStgy() hst.Stgy      { return x.X.HstStgys().Fst() }
func (x HstStgysMdl) Act()                   { x.HstStgy() }
func (x HstStgysMdl) Ifc() interface{}       { return x.HstStgy() }
func (x HstStgysMdl) HstStgy() hst.Stgy      { return x.X.HstStgys().Mdl() }
func (x HstStgysLst) Act()                   { x.HstStgy() }
func (x HstStgysLst) Ifc() interface{}       { return x.HstStgy() }
func (x HstStgysLst) HstStgy() hst.Stgy      { return x.X.HstStgys().Lst() }
func (x HstStgysFstIdx) Act()                { x.UntUnt() }
func (x HstStgysFstIdx) Ifc() interface{}    { return x.UntUnt() }
func (x HstStgysFstIdx) UntUnt() unt.Unt     { return x.X.HstStgys().FstIdx() }
func (x HstStgysMdlIdx) Act()                { x.UntUnt() }
func (x HstStgysMdlIdx) Ifc() interface{}    { return x.UntUnt() }
func (x HstStgysMdlIdx) UntUnt() unt.Unt     { return x.X.HstStgys().MdlIdx() }
func (x HstStgysLstIdx) Act()                { x.UntUnt() }
func (x HstStgysLstIdx) Ifc() interface{}    { return x.UntUnt() }
func (x HstStgysLstIdx) UntUnt() unt.Unt     { return x.X.HstStgys().LstIdx() }
func (x HstStgysRev) Act()                   { x.HstStgys() }
func (x HstStgysRev) Ifc() interface{}       { return x.HstStgys() }
func (x HstStgysRev) HstStgys() *hst.Stgys   { return x.X.HstStgys().Rev() }
func (x RltPrvsCnt) Act()                    { x.UntUnt() }
func (x RltPrvsCnt) Ifc() interface{}        { return x.UntUnt() }
func (x RltPrvsCnt) UntUnt() unt.Unt         { return x.X.RltPrvs().Cnt() }
func (x RltPrvsCpy) Act()                    { x.RltPrvs() }
func (x RltPrvsCpy) Ifc() interface{}        { return x.RltPrvs() }
func (x RltPrvsCpy) RltPrvs() *rlt.Prvs      { return x.X.RltPrvs().Cpy() }
func (x RltPrvsClr) Act()                    { x.RltPrvs() }
func (x RltPrvsClr) Ifc() interface{}        { return x.RltPrvs() }
func (x RltPrvsClr) RltPrvs() *rlt.Prvs      { return x.X.RltPrvs().Clr() }
func (x RltPrvsRand) Act()                   { x.RltPrvs() }
func (x RltPrvsRand) Ifc() interface{}       { return x.RltPrvs() }
func (x RltPrvsRand) RltPrvs() *rlt.Prvs     { return x.X.RltPrvs().Rand() }
func (x RltPrvsMrg) Act()                    { x.RltPrvs() }
func (x RltPrvsMrg) Ifc() interface{}        { return x.RltPrvs() }
func (x RltPrvsMrg) RltPrvs() *rlt.Prvs {
	var i0 []*rlt.Prvs
	for _, cur := range x.I0 {
		i0 = append(i0, cur.RltPrvs())
	}
	return x.X.RltPrvs().Mrg(i0...)
}
func (x RltPrvsPush) Act()             { x.RltPrvs() }
func (x RltPrvsPush) Ifc() interface{} { return x.RltPrvs() }
func (x RltPrvsPush) RltPrvs() *rlt.Prvs {
	var i0 []rlt.Prv
	for _, cur := range x.I0 {
		i0 = append(i0, cur.RltPrv())
	}
	return x.X.RltPrvs().Push(i0...)
}
func (x RltPrvsPop) Act()             { x.RltPrv() }
func (x RltPrvsPop) Ifc() interface{} { return x.RltPrv() }
func (x RltPrvsPop) RltPrv() rlt.Prv  { return x.X.RltPrvs().Pop() }
func (x RltPrvsQue) Act()             { x.RltPrvs() }
func (x RltPrvsQue) Ifc() interface{} { return x.RltPrvs() }
func (x RltPrvsQue) RltPrvs() *rlt.Prvs {
	var i0 []rlt.Prv
	for _, cur := range x.I0 {
		i0 = append(i0, cur.RltPrv())
	}
	return x.X.RltPrvs().Que(i0...)
}
func (x RltPrvsDque) Act()                     { x.RltPrv() }
func (x RltPrvsDque) Ifc() interface{}         { return x.RltPrv() }
func (x RltPrvsDque) RltPrv() rlt.Prv          { return x.X.RltPrvs().Dque() }
func (x RltPrvsIns) Act()                      { x.RltPrvs() }
func (x RltPrvsIns) Ifc() interface{}          { return x.RltPrvs() }
func (x RltPrvsIns) RltPrvs() *rlt.Prvs        { return x.X.RltPrvs().Ins(x.I0.UntUnt(), x.I1.RltPrv()) }
func (x RltPrvsUpd) Act()                      { x.RltPrvs() }
func (x RltPrvsUpd) Ifc() interface{}          { return x.RltPrvs() }
func (x RltPrvsUpd) RltPrvs() *rlt.Prvs        { return x.X.RltPrvs().Upd(x.I0.UntUnt(), x.I1.RltPrv()) }
func (x RltPrvsDel) Act()                      { x.RltPrv() }
func (x RltPrvsDel) Ifc() interface{}          { return x.RltPrv() }
func (x RltPrvsDel) RltPrv() rlt.Prv           { return x.X.RltPrvs().Del(x.I0.UntUnt()) }
func (x RltPrvsAt) Act()                       { x.RltPrv() }
func (x RltPrvsAt) Ifc() interface{}           { return x.RltPrv() }
func (x RltPrvsAt) RltPrv() rlt.Prv            { return x.X.RltPrvs().At(x.I0.UntUnt()) }
func (x RltPrvsIn) Act()                       { x.RltPrvs() }
func (x RltPrvsIn) Ifc() interface{}           { return x.RltPrvs() }
func (x RltPrvsIn) RltPrvs() *rlt.Prvs         { return x.X.RltPrvs().In(x.I0.UntUnt(), x.I1.UntUnt()) }
func (x RltPrvsInBnd) Act()                    { x.RltPrvs() }
func (x RltPrvsInBnd) Ifc() interface{}        { return x.RltPrvs() }
func (x RltPrvsInBnd) RltPrvs() *rlt.Prvs      { return x.X.RltPrvs().InBnd(x.I0.BndBnd()) }
func (x RltPrvsFrom) Act()                     { x.RltPrvs() }
func (x RltPrvsFrom) Ifc() interface{}         { return x.RltPrvs() }
func (x RltPrvsFrom) RltPrvs() *rlt.Prvs       { return x.X.RltPrvs().From(x.I0.UntUnt()) }
func (x RltPrvsTo) Act()                       { x.RltPrvs() }
func (x RltPrvsTo) Ifc() interface{}           { return x.RltPrvs() }
func (x RltPrvsTo) RltPrvs() *rlt.Prvs         { return x.X.RltPrvs().To(x.I0.UntUnt()) }
func (x RltPrvsFst) Act()                      { x.RltPrv() }
func (x RltPrvsFst) Ifc() interface{}          { return x.RltPrv() }
func (x RltPrvsFst) RltPrv() rlt.Prv           { return x.X.RltPrvs().Fst() }
func (x RltPrvsMdl) Act()                      { x.RltPrv() }
func (x RltPrvsMdl) Ifc() interface{}          { return x.RltPrv() }
func (x RltPrvsMdl) RltPrv() rlt.Prv           { return x.X.RltPrvs().Mdl() }
func (x RltPrvsLst) Act()                      { x.RltPrv() }
func (x RltPrvsLst) Ifc() interface{}          { return x.RltPrv() }
func (x RltPrvsLst) RltPrv() rlt.Prv           { return x.X.RltPrvs().Lst() }
func (x RltPrvsFstIdx) Act()                   { x.UntUnt() }
func (x RltPrvsFstIdx) Ifc() interface{}       { return x.UntUnt() }
func (x RltPrvsFstIdx) UntUnt() unt.Unt        { return x.X.RltPrvs().FstIdx() }
func (x RltPrvsMdlIdx) Act()                   { x.UntUnt() }
func (x RltPrvsMdlIdx) Ifc() interface{}       { return x.UntUnt() }
func (x RltPrvsMdlIdx) UntUnt() unt.Unt        { return x.X.RltPrvs().MdlIdx() }
func (x RltPrvsLstIdx) Act()                   { x.UntUnt() }
func (x RltPrvsLstIdx) Ifc() interface{}       { return x.UntUnt() }
func (x RltPrvsLstIdx) UntUnt() unt.Unt        { return x.X.RltPrvs().LstIdx() }
func (x RltPrvsRev) Act()                      { x.RltPrvs() }
func (x RltPrvsRev) Ifc() interface{}          { return x.RltPrvs() }
func (x RltPrvsRev) RltPrvs() *rlt.Prvs        { return x.X.RltPrvs().Rev() }
func (x RltInstrsCnt) Act()                    { x.UntUnt() }
func (x RltInstrsCnt) Ifc() interface{}        { return x.UntUnt() }
func (x RltInstrsCnt) UntUnt() unt.Unt         { return x.X.RltInstrs().Cnt() }
func (x RltInstrsCpy) Act()                    { x.RltInstrs() }
func (x RltInstrsCpy) Ifc() interface{}        { return x.RltInstrs() }
func (x RltInstrsCpy) RltInstrs() *rlt.Instrs  { return x.X.RltInstrs().Cpy() }
func (x RltInstrsClr) Act()                    { x.RltInstrs() }
func (x RltInstrsClr) Ifc() interface{}        { return x.RltInstrs() }
func (x RltInstrsClr) RltInstrs() *rlt.Instrs  { return x.X.RltInstrs().Clr() }
func (x RltInstrsRand) Act()                   { x.RltInstrs() }
func (x RltInstrsRand) Ifc() interface{}       { return x.RltInstrs() }
func (x RltInstrsRand) RltInstrs() *rlt.Instrs { return x.X.RltInstrs().Rand() }
func (x RltInstrsMrg) Act()                    { x.RltInstrs() }
func (x RltInstrsMrg) Ifc() interface{}        { return x.RltInstrs() }
func (x RltInstrsMrg) RltInstrs() *rlt.Instrs {
	var i0 []*rlt.Instrs
	for _, cur := range x.I0 {
		i0 = append(i0, cur.RltInstrs())
	}
	return x.X.RltInstrs().Mrg(i0...)
}
func (x RltInstrsPush) Act()             { x.RltInstrs() }
func (x RltInstrsPush) Ifc() interface{} { return x.RltInstrs() }
func (x RltInstrsPush) RltInstrs() *rlt.Instrs {
	var i0 []rlt.Instr
	for _, cur := range x.I0 {
		i0 = append(i0, cur.RltInstr())
	}
	return x.X.RltInstrs().Push(i0...)
}
func (x RltInstrsPop) Act()                { x.RltInstr() }
func (x RltInstrsPop) Ifc() interface{}    { return x.RltInstr() }
func (x RltInstrsPop) RltInstr() rlt.Instr { return x.X.RltInstrs().Pop() }
func (x RltInstrsQue) Act()                { x.RltInstrs() }
func (x RltInstrsQue) Ifc() interface{}    { return x.RltInstrs() }
func (x RltInstrsQue) RltInstrs() *rlt.Instrs {
	var i0 []rlt.Instr
	for _, cur := range x.I0 {
		i0 = append(i0, cur.RltInstr())
	}
	return x.X.RltInstrs().Que(i0...)
}
func (x RltInstrsDque) Act()                { x.RltInstr() }
func (x RltInstrsDque) Ifc() interface{}    { return x.RltInstr() }
func (x RltInstrsDque) RltInstr() rlt.Instr { return x.X.RltInstrs().Dque() }
func (x RltInstrsIns) Act()                 { x.RltInstrs() }
func (x RltInstrsIns) Ifc() interface{}     { return x.RltInstrs() }
func (x RltInstrsIns) RltInstrs() *rlt.Instrs {
	return x.X.RltInstrs().Ins(x.I0.UntUnt(), x.I1.RltInstr())
}
func (x RltInstrsUpd) Act()             { x.RltInstrs() }
func (x RltInstrsUpd) Ifc() interface{} { return x.RltInstrs() }
func (x RltInstrsUpd) RltInstrs() *rlt.Instrs {
	return x.X.RltInstrs().Upd(x.I0.UntUnt(), x.I1.RltInstr())
}
func (x RltInstrsDel) Act()                     { x.RltInstr() }
func (x RltInstrsDel) Ifc() interface{}         { return x.RltInstr() }
func (x RltInstrsDel) RltInstr() rlt.Instr      { return x.X.RltInstrs().Del(x.I0.UntUnt()) }
func (x RltInstrsAt) Act()                      { x.RltInstr() }
func (x RltInstrsAt) Ifc() interface{}          { return x.RltInstr() }
func (x RltInstrsAt) RltInstr() rlt.Instr       { return x.X.RltInstrs().At(x.I0.UntUnt()) }
func (x RltInstrsIn) Act()                      { x.RltInstrs() }
func (x RltInstrsIn) Ifc() interface{}          { return x.RltInstrs() }
func (x RltInstrsIn) RltInstrs() *rlt.Instrs    { return x.X.RltInstrs().In(x.I0.UntUnt(), x.I1.UntUnt()) }
func (x RltInstrsInBnd) Act()                   { x.RltInstrs() }
func (x RltInstrsInBnd) Ifc() interface{}       { return x.RltInstrs() }
func (x RltInstrsInBnd) RltInstrs() *rlt.Instrs { return x.X.RltInstrs().InBnd(x.I0.BndBnd()) }
func (x RltInstrsFrom) Act()                    { x.RltInstrs() }
func (x RltInstrsFrom) Ifc() interface{}        { return x.RltInstrs() }
func (x RltInstrsFrom) RltInstrs() *rlt.Instrs  { return x.X.RltInstrs().From(x.I0.UntUnt()) }
func (x RltInstrsTo) Act()                      { x.RltInstrs() }
func (x RltInstrsTo) Ifc() interface{}          { return x.RltInstrs() }
func (x RltInstrsTo) RltInstrs() *rlt.Instrs    { return x.X.RltInstrs().To(x.I0.UntUnt()) }
func (x RltInstrsFst) Act()                     { x.RltInstr() }
func (x RltInstrsFst) Ifc() interface{}         { return x.RltInstr() }
func (x RltInstrsFst) RltInstr() rlt.Instr      { return x.X.RltInstrs().Fst() }
func (x RltInstrsMdl) Act()                     { x.RltInstr() }
func (x RltInstrsMdl) Ifc() interface{}         { return x.RltInstr() }
func (x RltInstrsMdl) RltInstr() rlt.Instr      { return x.X.RltInstrs().Mdl() }
func (x RltInstrsLst) Act()                     { x.RltInstr() }
func (x RltInstrsLst) Ifc() interface{}         { return x.RltInstr() }
func (x RltInstrsLst) RltInstr() rlt.Instr      { return x.X.RltInstrs().Lst() }
func (x RltInstrsFstIdx) Act()                  { x.UntUnt() }
func (x RltInstrsFstIdx) Ifc() interface{}      { return x.UntUnt() }
func (x RltInstrsFstIdx) UntUnt() unt.Unt       { return x.X.RltInstrs().FstIdx() }
func (x RltInstrsMdlIdx) Act()                  { x.UntUnt() }
func (x RltInstrsMdlIdx) Ifc() interface{}      { return x.UntUnt() }
func (x RltInstrsMdlIdx) UntUnt() unt.Unt       { return x.X.RltInstrs().MdlIdx() }
func (x RltInstrsLstIdx) Act()                  { x.UntUnt() }
func (x RltInstrsLstIdx) Ifc() interface{}      { return x.UntUnt() }
func (x RltInstrsLstIdx) UntUnt() unt.Unt       { return x.X.RltInstrs().LstIdx() }
func (x RltInstrsRev) Act()                     { x.RltInstrs() }
func (x RltInstrsRev) Ifc() interface{}         { return x.RltInstrs() }
func (x RltInstrsRev) RltInstrs() *rlt.Instrs   { return x.X.RltInstrs().Rev() }
func (x RltInrvlsCnt) Act()                     { x.UntUnt() }
func (x RltInrvlsCnt) Ifc() interface{}         { return x.UntUnt() }
func (x RltInrvlsCnt) UntUnt() unt.Unt          { return x.X.RltInrvls().Cnt() }
func (x RltInrvlsCpy) Act()                     { x.RltInrvls() }
func (x RltInrvlsCpy) Ifc() interface{}         { return x.RltInrvls() }
func (x RltInrvlsCpy) RltInrvls() *rlt.Inrvls   { return x.X.RltInrvls().Cpy() }
func (x RltInrvlsClr) Act()                     { x.RltInrvls() }
func (x RltInrvlsClr) Ifc() interface{}         { return x.RltInrvls() }
func (x RltInrvlsClr) RltInrvls() *rlt.Inrvls   { return x.X.RltInrvls().Clr() }
func (x RltInrvlsRand) Act()                    { x.RltInrvls() }
func (x RltInrvlsRand) Ifc() interface{}        { return x.RltInrvls() }
func (x RltInrvlsRand) RltInrvls() *rlt.Inrvls  { return x.X.RltInrvls().Rand() }
func (x RltInrvlsMrg) Act()                     { x.RltInrvls() }
func (x RltInrvlsMrg) Ifc() interface{}         { return x.RltInrvls() }
func (x RltInrvlsMrg) RltInrvls() *rlt.Inrvls {
	var i0 []*rlt.Inrvls
	for _, cur := range x.I0 {
		i0 = append(i0, cur.RltInrvls())
	}
	return x.X.RltInrvls().Mrg(i0...)
}
func (x RltInrvlsPush) Act()             { x.RltInrvls() }
func (x RltInrvlsPush) Ifc() interface{} { return x.RltInrvls() }
func (x RltInrvlsPush) RltInrvls() *rlt.Inrvls {
	var i0 []rlt.Inrvl
	for _, cur := range x.I0 {
		i0 = append(i0, cur.RltInrvl())
	}
	return x.X.RltInrvls().Push(i0...)
}
func (x RltInrvlsPop) Act()                { x.RltInrvl() }
func (x RltInrvlsPop) Ifc() interface{}    { return x.RltInrvl() }
func (x RltInrvlsPop) RltInrvl() rlt.Inrvl { return x.X.RltInrvls().Pop() }
func (x RltInrvlsQue) Act()                { x.RltInrvls() }
func (x RltInrvlsQue) Ifc() interface{}    { return x.RltInrvls() }
func (x RltInrvlsQue) RltInrvls() *rlt.Inrvls {
	var i0 []rlt.Inrvl
	for _, cur := range x.I0 {
		i0 = append(i0, cur.RltInrvl())
	}
	return x.X.RltInrvls().Que(i0...)
}
func (x RltInrvlsDque) Act()                { x.RltInrvl() }
func (x RltInrvlsDque) Ifc() interface{}    { return x.RltInrvl() }
func (x RltInrvlsDque) RltInrvl() rlt.Inrvl { return x.X.RltInrvls().Dque() }
func (x RltInrvlsIns) Act()                 { x.RltInrvls() }
func (x RltInrvlsIns) Ifc() interface{}     { return x.RltInrvls() }
func (x RltInrvlsIns) RltInrvls() *rlt.Inrvls {
	return x.X.RltInrvls().Ins(x.I0.UntUnt(), x.I1.RltInrvl())
}
func (x RltInrvlsUpd) Act()             { x.RltInrvls() }
func (x RltInrvlsUpd) Ifc() interface{} { return x.RltInrvls() }
func (x RltInrvlsUpd) RltInrvls() *rlt.Inrvls {
	return x.X.RltInrvls().Upd(x.I0.UntUnt(), x.I1.RltInrvl())
}
func (x RltInrvlsDel) Act()                     { x.RltInrvl() }
func (x RltInrvlsDel) Ifc() interface{}         { return x.RltInrvl() }
func (x RltInrvlsDel) RltInrvl() rlt.Inrvl      { return x.X.RltInrvls().Del(x.I0.UntUnt()) }
func (x RltInrvlsAt) Act()                      { x.RltInrvl() }
func (x RltInrvlsAt) Ifc() interface{}          { return x.RltInrvl() }
func (x RltInrvlsAt) RltInrvl() rlt.Inrvl       { return x.X.RltInrvls().At(x.I0.UntUnt()) }
func (x RltInrvlsIn) Act()                      { x.RltInrvls() }
func (x RltInrvlsIn) Ifc() interface{}          { return x.RltInrvls() }
func (x RltInrvlsIn) RltInrvls() *rlt.Inrvls    { return x.X.RltInrvls().In(x.I0.UntUnt(), x.I1.UntUnt()) }
func (x RltInrvlsInBnd) Act()                   { x.RltInrvls() }
func (x RltInrvlsInBnd) Ifc() interface{}       { return x.RltInrvls() }
func (x RltInrvlsInBnd) RltInrvls() *rlt.Inrvls { return x.X.RltInrvls().InBnd(x.I0.BndBnd()) }
func (x RltInrvlsFrom) Act()                    { x.RltInrvls() }
func (x RltInrvlsFrom) Ifc() interface{}        { return x.RltInrvls() }
func (x RltInrvlsFrom) RltInrvls() *rlt.Inrvls  { return x.X.RltInrvls().From(x.I0.UntUnt()) }
func (x RltInrvlsTo) Act()                      { x.RltInrvls() }
func (x RltInrvlsTo) Ifc() interface{}          { return x.RltInrvls() }
func (x RltInrvlsTo) RltInrvls() *rlt.Inrvls    { return x.X.RltInrvls().To(x.I0.UntUnt()) }
func (x RltInrvlsFst) Act()                     { x.RltInrvl() }
func (x RltInrvlsFst) Ifc() interface{}         { return x.RltInrvl() }
func (x RltInrvlsFst) RltInrvl() rlt.Inrvl      { return x.X.RltInrvls().Fst() }
func (x RltInrvlsMdl) Act()                     { x.RltInrvl() }
func (x RltInrvlsMdl) Ifc() interface{}         { return x.RltInrvl() }
func (x RltInrvlsMdl) RltInrvl() rlt.Inrvl      { return x.X.RltInrvls().Mdl() }
func (x RltInrvlsLst) Act()                     { x.RltInrvl() }
func (x RltInrvlsLst) Ifc() interface{}         { return x.RltInrvl() }
func (x RltInrvlsLst) RltInrvl() rlt.Inrvl      { return x.X.RltInrvls().Lst() }
func (x RltInrvlsFstIdx) Act()                  { x.UntUnt() }
func (x RltInrvlsFstIdx) Ifc() interface{}      { return x.UntUnt() }
func (x RltInrvlsFstIdx) UntUnt() unt.Unt       { return x.X.RltInrvls().FstIdx() }
func (x RltInrvlsMdlIdx) Act()                  { x.UntUnt() }
func (x RltInrvlsMdlIdx) Ifc() interface{}      { return x.UntUnt() }
func (x RltInrvlsMdlIdx) UntUnt() unt.Unt       { return x.X.RltInrvls().MdlIdx() }
func (x RltInrvlsLstIdx) Act()                  { x.UntUnt() }
func (x RltInrvlsLstIdx) Ifc() interface{}      { return x.UntUnt() }
func (x RltInrvlsLstIdx) UntUnt() unt.Unt       { return x.X.RltInrvls().LstIdx() }
func (x RltInrvlsRev) Act()                     { x.RltInrvls() }
func (x RltInrvlsRev) Ifc() interface{}         { return x.RltInrvls() }
func (x RltInrvlsRev) RltInrvls() *rlt.Inrvls   { return x.X.RltInrvls().Rev() }
func (x RltSidesCnt) Act()                      { x.UntUnt() }
func (x RltSidesCnt) Ifc() interface{}          { return x.UntUnt() }
func (x RltSidesCnt) UntUnt() unt.Unt           { return x.X.RltSides().Cnt() }
func (x RltSidesCpy) Act()                      { x.RltSides() }
func (x RltSidesCpy) Ifc() interface{}          { return x.RltSides() }
func (x RltSidesCpy) RltSides() *rlt.Sides      { return x.X.RltSides().Cpy() }
func (x RltSidesClr) Act()                      { x.RltSides() }
func (x RltSidesClr) Ifc() interface{}          { return x.RltSides() }
func (x RltSidesClr) RltSides() *rlt.Sides      { return x.X.RltSides().Clr() }
func (x RltSidesRand) Act()                     { x.RltSides() }
func (x RltSidesRand) Ifc() interface{}         { return x.RltSides() }
func (x RltSidesRand) RltSides() *rlt.Sides     { return x.X.RltSides().Rand() }
func (x RltSidesMrg) Act()                      { x.RltSides() }
func (x RltSidesMrg) Ifc() interface{}          { return x.RltSides() }
func (x RltSidesMrg) RltSides() *rlt.Sides {
	var i0 []*rlt.Sides
	for _, cur := range x.I0 {
		i0 = append(i0, cur.RltSides())
	}
	return x.X.RltSides().Mrg(i0...)
}
func (x RltSidesPush) Act()             { x.RltSides() }
func (x RltSidesPush) Ifc() interface{} { return x.RltSides() }
func (x RltSidesPush) RltSides() *rlt.Sides {
	var i0 []rlt.Side
	for _, cur := range x.I0 {
		i0 = append(i0, cur.RltSide())
	}
	return x.X.RltSides().Push(i0...)
}
func (x RltSidesPop) Act()              { x.RltSide() }
func (x RltSidesPop) Ifc() interface{}  { return x.RltSide() }
func (x RltSidesPop) RltSide() rlt.Side { return x.X.RltSides().Pop() }
func (x RltSidesQue) Act()              { x.RltSides() }
func (x RltSidesQue) Ifc() interface{}  { return x.RltSides() }
func (x RltSidesQue) RltSides() *rlt.Sides {
	var i0 []rlt.Side
	for _, cur := range x.I0 {
		i0 = append(i0, cur.RltSide())
	}
	return x.X.RltSides().Que(i0...)
}
func (x RltSidesDque) Act()                  { x.RltSide() }
func (x RltSidesDque) Ifc() interface{}      { return x.RltSide() }
func (x RltSidesDque) RltSide() rlt.Side     { return x.X.RltSides().Dque() }
func (x RltSidesIns) Act()                   { x.RltSides() }
func (x RltSidesIns) Ifc() interface{}       { return x.RltSides() }
func (x RltSidesIns) RltSides() *rlt.Sides   { return x.X.RltSides().Ins(x.I0.UntUnt(), x.I1.RltSide()) }
func (x RltSidesUpd) Act()                   { x.RltSides() }
func (x RltSidesUpd) Ifc() interface{}       { return x.RltSides() }
func (x RltSidesUpd) RltSides() *rlt.Sides   { return x.X.RltSides().Upd(x.I0.UntUnt(), x.I1.RltSide()) }
func (x RltSidesDel) Act()                   { x.RltSide() }
func (x RltSidesDel) Ifc() interface{}       { return x.RltSide() }
func (x RltSidesDel) RltSide() rlt.Side      { return x.X.RltSides().Del(x.I0.UntUnt()) }
func (x RltSidesAt) Act()                    { x.RltSide() }
func (x RltSidesAt) Ifc() interface{}        { return x.RltSide() }
func (x RltSidesAt) RltSide() rlt.Side       { return x.X.RltSides().At(x.I0.UntUnt()) }
func (x RltSidesIn) Act()                    { x.RltSides() }
func (x RltSidesIn) Ifc() interface{}        { return x.RltSides() }
func (x RltSidesIn) RltSides() *rlt.Sides    { return x.X.RltSides().In(x.I0.UntUnt(), x.I1.UntUnt()) }
func (x RltSidesInBnd) Act()                 { x.RltSides() }
func (x RltSidesInBnd) Ifc() interface{}     { return x.RltSides() }
func (x RltSidesInBnd) RltSides() *rlt.Sides { return x.X.RltSides().InBnd(x.I0.BndBnd()) }
func (x RltSidesFrom) Act()                  { x.RltSides() }
func (x RltSidesFrom) Ifc() interface{}      { return x.RltSides() }
func (x RltSidesFrom) RltSides() *rlt.Sides  { return x.X.RltSides().From(x.I0.UntUnt()) }
func (x RltSidesTo) Act()                    { x.RltSides() }
func (x RltSidesTo) Ifc() interface{}        { return x.RltSides() }
func (x RltSidesTo) RltSides() *rlt.Sides    { return x.X.RltSides().To(x.I0.UntUnt()) }
func (x RltSidesFst) Act()                   { x.RltSide() }
func (x RltSidesFst) Ifc() interface{}       { return x.RltSide() }
func (x RltSidesFst) RltSide() rlt.Side      { return x.X.RltSides().Fst() }
func (x RltSidesMdl) Act()                   { x.RltSide() }
func (x RltSidesMdl) Ifc() interface{}       { return x.RltSide() }
func (x RltSidesMdl) RltSide() rlt.Side      { return x.X.RltSides().Mdl() }
func (x RltSidesLst) Act()                   { x.RltSide() }
func (x RltSidesLst) Ifc() interface{}       { return x.RltSide() }
func (x RltSidesLst) RltSide() rlt.Side      { return x.X.RltSides().Lst() }
func (x RltSidesFstIdx) Act()                { x.UntUnt() }
func (x RltSidesFstIdx) Ifc() interface{}    { return x.UntUnt() }
func (x RltSidesFstIdx) UntUnt() unt.Unt     { return x.X.RltSides().FstIdx() }
func (x RltSidesMdlIdx) Act()                { x.UntUnt() }
func (x RltSidesMdlIdx) Ifc() interface{}    { return x.UntUnt() }
func (x RltSidesMdlIdx) UntUnt() unt.Unt     { return x.X.RltSides().MdlIdx() }
func (x RltSidesLstIdx) Act()                { x.UntUnt() }
func (x RltSidesLstIdx) Ifc() interface{}    { return x.UntUnt() }
func (x RltSidesLstIdx) UntUnt() unt.Unt     { return x.X.RltSides().LstIdx() }
func (x RltSidesRev) Act()                   { x.RltSides() }
func (x RltSidesRev) Ifc() interface{}       { return x.RltSides() }
func (x RltSidesRev) RltSides() *rlt.Sides   { return x.X.RltSides().Rev() }
func (x RltStmsCnt) Act()                    { x.UntUnt() }
func (x RltStmsCnt) Ifc() interface{}        { return x.UntUnt() }
func (x RltStmsCnt) UntUnt() unt.Unt         { return x.X.RltStms().Cnt() }
func (x RltStmsCpy) Act()                    { x.RltStms() }
func (x RltStmsCpy) Ifc() interface{}        { return x.RltStms() }
func (x RltStmsCpy) RltStms() *rlt.Stms      { return x.X.RltStms().Cpy() }
func (x RltStmsClr) Act()                    { x.RltStms() }
func (x RltStmsClr) Ifc() interface{}        { return x.RltStms() }
func (x RltStmsClr) RltStms() *rlt.Stms      { return x.X.RltStms().Clr() }
func (x RltStmsRand) Act()                   { x.RltStms() }
func (x RltStmsRand) Ifc() interface{}       { return x.RltStms() }
func (x RltStmsRand) RltStms() *rlt.Stms     { return x.X.RltStms().Rand() }
func (x RltStmsMrg) Act()                    { x.RltStms() }
func (x RltStmsMrg) Ifc() interface{}        { return x.RltStms() }
func (x RltStmsMrg) RltStms() *rlt.Stms {
	var i0 []*rlt.Stms
	for _, cur := range x.I0 {
		i0 = append(i0, cur.RltStms())
	}
	return x.X.RltStms().Mrg(i0...)
}
func (x RltStmsPush) Act()             { x.RltStms() }
func (x RltStmsPush) Ifc() interface{} { return x.RltStms() }
func (x RltStmsPush) RltStms() *rlt.Stms {
	var i0 []rlt.Stm
	for _, cur := range x.I0 {
		i0 = append(i0, cur.RltStm())
	}
	return x.X.RltStms().Push(i0...)
}
func (x RltStmsPop) Act()             { x.RltStm() }
func (x RltStmsPop) Ifc() interface{} { return x.RltStm() }
func (x RltStmsPop) RltStm() rlt.Stm  { return x.X.RltStms().Pop() }
func (x RltStmsQue) Act()             { x.RltStms() }
func (x RltStmsQue) Ifc() interface{} { return x.RltStms() }
func (x RltStmsQue) RltStms() *rlt.Stms {
	var i0 []rlt.Stm
	for _, cur := range x.I0 {
		i0 = append(i0, cur.RltStm())
	}
	return x.X.RltStms().Que(i0...)
}
func (x RltStmsDque) Act()                { x.RltStm() }
func (x RltStmsDque) Ifc() interface{}    { return x.RltStm() }
func (x RltStmsDque) RltStm() rlt.Stm     { return x.X.RltStms().Dque() }
func (x RltStmsIns) Act()                 { x.RltStms() }
func (x RltStmsIns) Ifc() interface{}     { return x.RltStms() }
func (x RltStmsIns) RltStms() *rlt.Stms   { return x.X.RltStms().Ins(x.I0.UntUnt(), x.I1.RltStm()) }
func (x RltStmsUpd) Act()                 { x.RltStms() }
func (x RltStmsUpd) Ifc() interface{}     { return x.RltStms() }
func (x RltStmsUpd) RltStms() *rlt.Stms   { return x.X.RltStms().Upd(x.I0.UntUnt(), x.I1.RltStm()) }
func (x RltStmsDel) Act()                 { x.RltStm() }
func (x RltStmsDel) Ifc() interface{}     { return x.RltStm() }
func (x RltStmsDel) RltStm() rlt.Stm      { return x.X.RltStms().Del(x.I0.UntUnt()) }
func (x RltStmsAt) Act()                  { x.RltStm() }
func (x RltStmsAt) Ifc() interface{}      { return x.RltStm() }
func (x RltStmsAt) RltStm() rlt.Stm       { return x.X.RltStms().At(x.I0.UntUnt()) }
func (x RltStmsIn) Act()                  { x.RltStms() }
func (x RltStmsIn) Ifc() interface{}      { return x.RltStms() }
func (x RltStmsIn) RltStms() *rlt.Stms    { return x.X.RltStms().In(x.I0.UntUnt(), x.I1.UntUnt()) }
func (x RltStmsInBnd) Act()               { x.RltStms() }
func (x RltStmsInBnd) Ifc() interface{}   { return x.RltStms() }
func (x RltStmsInBnd) RltStms() *rlt.Stms { return x.X.RltStms().InBnd(x.I0.BndBnd()) }
func (x RltStmsFrom) Act()                { x.RltStms() }
func (x RltStmsFrom) Ifc() interface{}    { return x.RltStms() }
func (x RltStmsFrom) RltStms() *rlt.Stms  { return x.X.RltStms().From(x.I0.UntUnt()) }
func (x RltStmsTo) Act()                  { x.RltStms() }
func (x RltStmsTo) Ifc() interface{}      { return x.RltStms() }
func (x RltStmsTo) RltStms() *rlt.Stms    { return x.X.RltStms().To(x.I0.UntUnt()) }
func (x RltStmsFst) Act()                 { x.RltStm() }
func (x RltStmsFst) Ifc() interface{}     { return x.RltStm() }
func (x RltStmsFst) RltStm() rlt.Stm      { return x.X.RltStms().Fst() }
func (x RltStmsMdl) Act()                 { x.RltStm() }
func (x RltStmsMdl) Ifc() interface{}     { return x.RltStm() }
func (x RltStmsMdl) RltStm() rlt.Stm      { return x.X.RltStms().Mdl() }
func (x RltStmsLst) Act()                 { x.RltStm() }
func (x RltStmsLst) Ifc() interface{}     { return x.RltStm() }
func (x RltStmsLst) RltStm() rlt.Stm      { return x.X.RltStms().Lst() }
func (x RltStmsFstIdx) Act()              { x.UntUnt() }
func (x RltStmsFstIdx) Ifc() interface{}  { return x.UntUnt() }
func (x RltStmsFstIdx) UntUnt() unt.Unt   { return x.X.RltStms().FstIdx() }
func (x RltStmsMdlIdx) Act()              { x.UntUnt() }
func (x RltStmsMdlIdx) Ifc() interface{}  { return x.UntUnt() }
func (x RltStmsMdlIdx) UntUnt() unt.Unt   { return x.X.RltStms().MdlIdx() }
func (x RltStmsLstIdx) Act()              { x.UntUnt() }
func (x RltStmsLstIdx) Ifc() interface{}  { return x.UntUnt() }
func (x RltStmsLstIdx) UntUnt() unt.Unt   { return x.X.RltStms().LstIdx() }
func (x RltStmsRev) Act()                 { x.RltStms() }
func (x RltStmsRev) Ifc() interface{}     { return x.RltStms() }
func (x RltStmsRev) RltStms() *rlt.Stms   { return x.X.RltStms().Rev() }
func (x RltCndsCnt) Act()                 { x.UntUnt() }
func (x RltCndsCnt) Ifc() interface{}     { return x.UntUnt() }
func (x RltCndsCnt) UntUnt() unt.Unt      { return x.X.RltCnds().Cnt() }
func (x RltCndsCpy) Act()                 { x.RltCnds() }
func (x RltCndsCpy) Ifc() interface{}     { return x.RltCnds() }
func (x RltCndsCpy) RltCnds() *rlt.Cnds   { return x.X.RltCnds().Cpy() }
func (x RltCndsClr) Act()                 { x.RltCnds() }
func (x RltCndsClr) Ifc() interface{}     { return x.RltCnds() }
func (x RltCndsClr) RltCnds() *rlt.Cnds   { return x.X.RltCnds().Clr() }
func (x RltCndsRand) Act()                { x.RltCnds() }
func (x RltCndsRand) Ifc() interface{}    { return x.RltCnds() }
func (x RltCndsRand) RltCnds() *rlt.Cnds  { return x.X.RltCnds().Rand() }
func (x RltCndsMrg) Act()                 { x.RltCnds() }
func (x RltCndsMrg) Ifc() interface{}     { return x.RltCnds() }
func (x RltCndsMrg) RltCnds() *rlt.Cnds {
	var i0 []*rlt.Cnds
	for _, cur := range x.I0 {
		i0 = append(i0, cur.RltCnds())
	}
	return x.X.RltCnds().Mrg(i0...)
}
func (x RltCndsPush) Act()             { x.RltCnds() }
func (x RltCndsPush) Ifc() interface{} { return x.RltCnds() }
func (x RltCndsPush) RltCnds() *rlt.Cnds {
	var i0 []rlt.Cnd
	for _, cur := range x.I0 {
		i0 = append(i0, cur.RltCnd())
	}
	return x.X.RltCnds().Push(i0...)
}
func (x RltCndsPop) Act()             { x.RltCnd() }
func (x RltCndsPop) Ifc() interface{} { return x.RltCnd() }
func (x RltCndsPop) RltCnd() rlt.Cnd  { return x.X.RltCnds().Pop() }
func (x RltCndsQue) Act()             { x.RltCnds() }
func (x RltCndsQue) Ifc() interface{} { return x.RltCnds() }
func (x RltCndsQue) RltCnds() *rlt.Cnds {
	var i0 []rlt.Cnd
	for _, cur := range x.I0 {
		i0 = append(i0, cur.RltCnd())
	}
	return x.X.RltCnds().Que(i0...)
}
func (x RltCndsDque) Act()                  { x.RltCnd() }
func (x RltCndsDque) Ifc() interface{}      { return x.RltCnd() }
func (x RltCndsDque) RltCnd() rlt.Cnd       { return x.X.RltCnds().Dque() }
func (x RltCndsIns) Act()                   { x.RltCnds() }
func (x RltCndsIns) Ifc() interface{}       { return x.RltCnds() }
func (x RltCndsIns) RltCnds() *rlt.Cnds     { return x.X.RltCnds().Ins(x.I0.UntUnt(), x.I1.RltCnd()) }
func (x RltCndsUpd) Act()                   { x.RltCnds() }
func (x RltCndsUpd) Ifc() interface{}       { return x.RltCnds() }
func (x RltCndsUpd) RltCnds() *rlt.Cnds     { return x.X.RltCnds().Upd(x.I0.UntUnt(), x.I1.RltCnd()) }
func (x RltCndsDel) Act()                   { x.RltCnd() }
func (x RltCndsDel) Ifc() interface{}       { return x.RltCnd() }
func (x RltCndsDel) RltCnd() rlt.Cnd        { return x.X.RltCnds().Del(x.I0.UntUnt()) }
func (x RltCndsAt) Act()                    { x.RltCnd() }
func (x RltCndsAt) Ifc() interface{}        { return x.RltCnd() }
func (x RltCndsAt) RltCnd() rlt.Cnd         { return x.X.RltCnds().At(x.I0.UntUnt()) }
func (x RltCndsIn) Act()                    { x.RltCnds() }
func (x RltCndsIn) Ifc() interface{}        { return x.RltCnds() }
func (x RltCndsIn) RltCnds() *rlt.Cnds      { return x.X.RltCnds().In(x.I0.UntUnt(), x.I1.UntUnt()) }
func (x RltCndsInBnd) Act()                 { x.RltCnds() }
func (x RltCndsInBnd) Ifc() interface{}     { return x.RltCnds() }
func (x RltCndsInBnd) RltCnds() *rlt.Cnds   { return x.X.RltCnds().InBnd(x.I0.BndBnd()) }
func (x RltCndsFrom) Act()                  { x.RltCnds() }
func (x RltCndsFrom) Ifc() interface{}      { return x.RltCnds() }
func (x RltCndsFrom) RltCnds() *rlt.Cnds    { return x.X.RltCnds().From(x.I0.UntUnt()) }
func (x RltCndsTo) Act()                    { x.RltCnds() }
func (x RltCndsTo) Ifc() interface{}        { return x.RltCnds() }
func (x RltCndsTo) RltCnds() *rlt.Cnds      { return x.X.RltCnds().To(x.I0.UntUnt()) }
func (x RltCndsFst) Act()                   { x.RltCnd() }
func (x RltCndsFst) Ifc() interface{}       { return x.RltCnd() }
func (x RltCndsFst) RltCnd() rlt.Cnd        { return x.X.RltCnds().Fst() }
func (x RltCndsMdl) Act()                   { x.RltCnd() }
func (x RltCndsMdl) Ifc() interface{}       { return x.RltCnd() }
func (x RltCndsMdl) RltCnd() rlt.Cnd        { return x.X.RltCnds().Mdl() }
func (x RltCndsLst) Act()                   { x.RltCnd() }
func (x RltCndsLst) Ifc() interface{}       { return x.RltCnd() }
func (x RltCndsLst) RltCnd() rlt.Cnd        { return x.X.RltCnds().Lst() }
func (x RltCndsFstIdx) Act()                { x.UntUnt() }
func (x RltCndsFstIdx) Ifc() interface{}    { return x.UntUnt() }
func (x RltCndsFstIdx) UntUnt() unt.Unt     { return x.X.RltCnds().FstIdx() }
func (x RltCndsMdlIdx) Act()                { x.UntUnt() }
func (x RltCndsMdlIdx) Ifc() interface{}    { return x.UntUnt() }
func (x RltCndsMdlIdx) UntUnt() unt.Unt     { return x.X.RltCnds().MdlIdx() }
func (x RltCndsLstIdx) Act()                { x.UntUnt() }
func (x RltCndsLstIdx) Ifc() interface{}    { return x.UntUnt() }
func (x RltCndsLstIdx) UntUnt() unt.Unt     { return x.X.RltCnds().LstIdx() }
func (x RltCndsRev) Act()                   { x.RltCnds() }
func (x RltCndsRev) Ifc() interface{}       { return x.RltCnds() }
func (x RltCndsRev) RltCnds() *rlt.Cnds     { return x.X.RltCnds().Rev() }
func (x RltStgysCnt) Act()                  { x.UntUnt() }
func (x RltStgysCnt) Ifc() interface{}      { return x.UntUnt() }
func (x RltStgysCnt) UntUnt() unt.Unt       { return x.X.RltStgys().Cnt() }
func (x RltStgysCpy) Act()                  { x.RltStgys() }
func (x RltStgysCpy) Ifc() interface{}      { return x.RltStgys() }
func (x RltStgysCpy) RltStgys() *rlt.Stgys  { return x.X.RltStgys().Cpy() }
func (x RltStgysClr) Act()                  { x.RltStgys() }
func (x RltStgysClr) Ifc() interface{}      { return x.RltStgys() }
func (x RltStgysClr) RltStgys() *rlt.Stgys  { return x.X.RltStgys().Clr() }
func (x RltStgysRand) Act()                 { x.RltStgys() }
func (x RltStgysRand) Ifc() interface{}     { return x.RltStgys() }
func (x RltStgysRand) RltStgys() *rlt.Stgys { return x.X.RltStgys().Rand() }
func (x RltStgysMrg) Act()                  { x.RltStgys() }
func (x RltStgysMrg) Ifc() interface{}      { return x.RltStgys() }
func (x RltStgysMrg) RltStgys() *rlt.Stgys {
	var i0 []*rlt.Stgys
	for _, cur := range x.I0 {
		i0 = append(i0, cur.RltStgys())
	}
	return x.X.RltStgys().Mrg(i0...)
}
func (x RltStgysPush) Act()             { x.RltStgys() }
func (x RltStgysPush) Ifc() interface{} { return x.RltStgys() }
func (x RltStgysPush) RltStgys() *rlt.Stgys {
	var i0 []rlt.Stgy
	for _, cur := range x.I0 {
		i0 = append(i0, cur.RltStgy())
	}
	return x.X.RltStgys().Push(i0...)
}
func (x RltStgysPop) Act()              { x.RltStgy() }
func (x RltStgysPop) Ifc() interface{}  { return x.RltStgy() }
func (x RltStgysPop) RltStgy() rlt.Stgy { return x.X.RltStgys().Pop() }
func (x RltStgysQue) Act()              { x.RltStgys() }
func (x RltStgysQue) Ifc() interface{}  { return x.RltStgys() }
func (x RltStgysQue) RltStgys() *rlt.Stgys {
	var i0 []rlt.Stgy
	for _, cur := range x.I0 {
		i0 = append(i0, cur.RltStgy())
	}
	return x.X.RltStgys().Que(i0...)
}
func (x RltStgysDque) Act()                  { x.RltStgy() }
func (x RltStgysDque) Ifc() interface{}      { return x.RltStgy() }
func (x RltStgysDque) RltStgy() rlt.Stgy     { return x.X.RltStgys().Dque() }
func (x RltStgysIns) Act()                   { x.RltStgys() }
func (x RltStgysIns) Ifc() interface{}       { return x.RltStgys() }
func (x RltStgysIns) RltStgys() *rlt.Stgys   { return x.X.RltStgys().Ins(x.I0.UntUnt(), x.I1.RltStgy()) }
func (x RltStgysUpd) Act()                   { x.RltStgys() }
func (x RltStgysUpd) Ifc() interface{}       { return x.RltStgys() }
func (x RltStgysUpd) RltStgys() *rlt.Stgys   { return x.X.RltStgys().Upd(x.I0.UntUnt(), x.I1.RltStgy()) }
func (x RltStgysDel) Act()                   { x.RltStgy() }
func (x RltStgysDel) Ifc() interface{}       { return x.RltStgy() }
func (x RltStgysDel) RltStgy() rlt.Stgy      { return x.X.RltStgys().Del(x.I0.UntUnt()) }
func (x RltStgysAt) Act()                    { x.RltStgy() }
func (x RltStgysAt) Ifc() interface{}        { return x.RltStgy() }
func (x RltStgysAt) RltStgy() rlt.Stgy       { return x.X.RltStgys().At(x.I0.UntUnt()) }
func (x RltStgysIn) Act()                    { x.RltStgys() }
func (x RltStgysIn) Ifc() interface{}        { return x.RltStgys() }
func (x RltStgysIn) RltStgys() *rlt.Stgys    { return x.X.RltStgys().In(x.I0.UntUnt(), x.I1.UntUnt()) }
func (x RltStgysInBnd) Act()                 { x.RltStgys() }
func (x RltStgysInBnd) Ifc() interface{}     { return x.RltStgys() }
func (x RltStgysInBnd) RltStgys() *rlt.Stgys { return x.X.RltStgys().InBnd(x.I0.BndBnd()) }
func (x RltStgysFrom) Act()                  { x.RltStgys() }
func (x RltStgysFrom) Ifc() interface{}      { return x.RltStgys() }
func (x RltStgysFrom) RltStgys() *rlt.Stgys  { return x.X.RltStgys().From(x.I0.UntUnt()) }
func (x RltStgysTo) Act()                    { x.RltStgys() }
func (x RltStgysTo) Ifc() interface{}        { return x.RltStgys() }
func (x RltStgysTo) RltStgys() *rlt.Stgys    { return x.X.RltStgys().To(x.I0.UntUnt()) }
func (x RltStgysFst) Act()                   { x.RltStgy() }
func (x RltStgysFst) Ifc() interface{}       { return x.RltStgy() }
func (x RltStgysFst) RltStgy() rlt.Stgy      { return x.X.RltStgys().Fst() }
func (x RltStgysMdl) Act()                   { x.RltStgy() }
func (x RltStgysMdl) Ifc() interface{}       { return x.RltStgy() }
func (x RltStgysMdl) RltStgy() rlt.Stgy      { return x.X.RltStgys().Mdl() }
func (x RltStgysLst) Act()                   { x.RltStgy() }
func (x RltStgysLst) Ifc() interface{}       { return x.RltStgy() }
func (x RltStgysLst) RltStgy() rlt.Stgy      { return x.X.RltStgys().Lst() }
func (x RltStgysFstIdx) Act()                { x.UntUnt() }
func (x RltStgysFstIdx) Ifc() interface{}    { return x.UntUnt() }
func (x RltStgysFstIdx) UntUnt() unt.Unt     { return x.X.RltStgys().FstIdx() }
func (x RltStgysMdlIdx) Act()                { x.UntUnt() }
func (x RltStgysMdlIdx) Ifc() interface{}    { return x.UntUnt() }
func (x RltStgysMdlIdx) UntUnt() unt.Unt     { return x.X.RltStgys().MdlIdx() }
func (x RltStgysLstIdx) Act()                { x.UntUnt() }
func (x RltStgysLstIdx) Ifc() interface{}    { return x.UntUnt() }
func (x RltStgysLstIdx) UntUnt() unt.Unt     { return x.X.RltStgys().LstIdx() }
func (x RltStgysRev) Act()                   { x.RltStgys() }
func (x RltStgysRev) Ifc() interface{}       { return x.RltStgys() }
func (x RltStgysRev) RltStgys() *rlt.Stgys   { return x.X.RltStgys().Rev() }
func (x ClrClrOpa) Act()                     { x.ClrClr() }
func (x ClrClrOpa) Ifc() interface{}         { return x.ClrClr() }
func (x ClrClrOpa) ClrClr() clr.Clr          { return x.X.ClrClr().Opa(x.I0.FltFlt()) }
func (x ClrClrInv) Act()                     { x.ClrClr() }
func (x ClrClrInv) Ifc() interface{}         { return x.ClrClr() }
func (x ClrClrInv) ClrClr() clr.Clr          { return x.X.ClrClr().Inv() }
func (x PenPenOpa) Act()                     { x.PenPen() }
func (x PenPenOpa) Ifc() interface{}         { return x.PenPen() }
func (x PenPenOpa) PenPen() pen.Pen          { return x.X.PenPen().Opa(x.I0.FltFlt()) }
func (x PenPenInv) Act()                     { x.PenPen() }
func (x PenPenInv) Ifc() interface{}         { return x.PenPen() }
func (x PenPenInv) PenPen() pen.Pen          { return x.X.PenPen().Inv() }
func (x PenPensCnt) Act()                    { x.UntUnt() }
func (x PenPensCnt) Ifc() interface{}        { return x.UntUnt() }
func (x PenPensCnt) UntUnt() unt.Unt         { return x.X.PenPens().Cnt() }
func (x PenPensCpy) Act()                    { x.PenPens() }
func (x PenPensCpy) Ifc() interface{}        { return x.PenPens() }
func (x PenPensCpy) PenPens() *pen.Pens      { return x.X.PenPens().Cpy() }
func (x PenPensClr) Act()                    { x.PenPens() }
func (x PenPensClr) Ifc() interface{}        { return x.PenPens() }
func (x PenPensClr) PenPens() *pen.Pens      { return x.X.PenPens().Clr() }
func (x PenPensRand) Act()                   { x.PenPens() }
func (x PenPensRand) Ifc() interface{}       { return x.PenPens() }
func (x PenPensRand) PenPens() *pen.Pens     { return x.X.PenPens().Rand() }
func (x PenPensMrg) Act()                    { x.PenPens() }
func (x PenPensMrg) Ifc() interface{}        { return x.PenPens() }
func (x PenPensMrg) PenPens() *pen.Pens {
	var i0 []*pen.Pens
	for _, cur := range x.I0 {
		i0 = append(i0, cur.PenPens())
	}
	return x.X.PenPens().Mrg(i0...)
}
func (x PenPensPush) Act()             { x.PenPens() }
func (x PenPensPush) Ifc() interface{} { return x.PenPens() }
func (x PenPensPush) PenPens() *pen.Pens {
	var i0 []pen.Pen
	for _, cur := range x.I0 {
		i0 = append(i0, cur.PenPen())
	}
	return x.X.PenPens().Push(i0...)
}
func (x PenPensPop) Act()             { x.PenPen() }
func (x PenPensPop) Ifc() interface{} { return x.PenPen() }
func (x PenPensPop) PenPen() pen.Pen  { return x.X.PenPens().Pop() }
func (x PenPensQue) Act()             { x.PenPens() }
func (x PenPensQue) Ifc() interface{} { return x.PenPens() }
func (x PenPensQue) PenPens() *pen.Pens {
	var i0 []pen.Pen
	for _, cur := range x.I0 {
		i0 = append(i0, cur.PenPen())
	}
	return x.X.PenPens().Que(i0...)
}
func (x PenPensDque) Act()                { x.PenPen() }
func (x PenPensDque) Ifc() interface{}    { return x.PenPen() }
func (x PenPensDque) PenPen() pen.Pen     { return x.X.PenPens().Dque() }
func (x PenPensIns) Act()                 { x.PenPens() }
func (x PenPensIns) Ifc() interface{}     { return x.PenPens() }
func (x PenPensIns) PenPens() *pen.Pens   { return x.X.PenPens().Ins(x.I0.UntUnt(), x.I1.PenPen()) }
func (x PenPensUpd) Act()                 { x.PenPens() }
func (x PenPensUpd) Ifc() interface{}     { return x.PenPens() }
func (x PenPensUpd) PenPens() *pen.Pens   { return x.X.PenPens().Upd(x.I0.UntUnt(), x.I1.PenPen()) }
func (x PenPensDel) Act()                 { x.PenPen() }
func (x PenPensDel) Ifc() interface{}     { return x.PenPen() }
func (x PenPensDel) PenPen() pen.Pen      { return x.X.PenPens().Del(x.I0.UntUnt()) }
func (x PenPensAt) Act()                  { x.PenPen() }
func (x PenPensAt) Ifc() interface{}      { return x.PenPen() }
func (x PenPensAt) PenPen() pen.Pen       { return x.X.PenPens().At(x.I0.UntUnt()) }
func (x PenPensIn) Act()                  { x.PenPens() }
func (x PenPensIn) Ifc() interface{}      { return x.PenPens() }
func (x PenPensIn) PenPens() *pen.Pens    { return x.X.PenPens().In(x.I0.UntUnt(), x.I1.UntUnt()) }
func (x PenPensInBnd) Act()               { x.PenPens() }
func (x PenPensInBnd) Ifc() interface{}   { return x.PenPens() }
func (x PenPensInBnd) PenPens() *pen.Pens { return x.X.PenPens().InBnd(x.I0.BndBnd()) }
func (x PenPensFrom) Act()                { x.PenPens() }
func (x PenPensFrom) Ifc() interface{}    { return x.PenPens() }
func (x PenPensFrom) PenPens() *pen.Pens  { return x.X.PenPens().From(x.I0.UntUnt()) }
func (x PenPensTo) Act()                  { x.PenPens() }
func (x PenPensTo) Ifc() interface{}      { return x.PenPens() }
func (x PenPensTo) PenPens() *pen.Pens    { return x.X.PenPens().To(x.I0.UntUnt()) }
func (x PenPensFst) Act()                 { x.PenPen() }
func (x PenPensFst) Ifc() interface{}     { return x.PenPen() }
func (x PenPensFst) PenPen() pen.Pen      { return x.X.PenPens().Fst() }
func (x PenPensMdl) Act()                 { x.PenPen() }
func (x PenPensMdl) Ifc() interface{}     { return x.PenPen() }
func (x PenPensMdl) PenPen() pen.Pen      { return x.X.PenPens().Mdl() }
func (x PenPensLst) Act()                 { x.PenPen() }
func (x PenPensLst) Ifc() interface{}     { return x.PenPen() }
func (x PenPensLst) PenPen() pen.Pen      { return x.X.PenPens().Lst() }
func (x PenPensFstIdx) Act()              { x.UntUnt() }
func (x PenPensFstIdx) Ifc() interface{}  { return x.UntUnt() }
func (x PenPensFstIdx) UntUnt() unt.Unt   { return x.X.PenPens().FstIdx() }
func (x PenPensMdlIdx) Act()              { x.UntUnt() }
func (x PenPensMdlIdx) Ifc() interface{}  { return x.UntUnt() }
func (x PenPensMdlIdx) UntUnt() unt.Unt   { return x.X.PenPens().MdlIdx() }
func (x PenPensLstIdx) Act()              { x.UntUnt() }
func (x PenPensLstIdx) Ifc() interface{}  { return x.UntUnt() }
func (x PenPensLstIdx) UntUnt() unt.Unt   { return x.X.PenPens().LstIdx() }
func (x PenPensRev) Act()                 { x.PenPens() }
func (x PenPensRev) Ifc() interface{}     { return x.PenPens() }
func (x PenPensRev) PenPens() *pen.Pens   { return x.X.PenPens().Rev() }
func (x PltPltsCnt) Act()                 { x.UntUnt() }
func (x PltPltsCnt) Ifc() interface{}     { return x.UntUnt() }
func (x PltPltsCnt) UntUnt() unt.Unt      { return x.X.PltPlts().Cnt() }
func (x PltPltsCpy) Act()                 { x.PltPlts() }
func (x PltPltsCpy) Ifc() interface{}     { return x.PltPlts() }
func (x PltPltsCpy) PltPlts() *plt.Plts   { return x.X.PltPlts().Cpy() }
func (x PltPltsClr) Act()                 { x.PltPlts() }
func (x PltPltsClr) Ifc() interface{}     { return x.PltPlts() }
func (x PltPltsClr) PltPlts() *plt.Plts   { return x.X.PltPlts().Clr() }
func (x PltPltsRand) Act()                { x.PltPlts() }
func (x PltPltsRand) Ifc() interface{}    { return x.PltPlts() }
func (x PltPltsRand) PltPlts() *plt.Plts  { return x.X.PltPlts().Rand() }
func (x PltPltsMrg) Act()                 { x.PltPlts() }
func (x PltPltsMrg) Ifc() interface{}     { return x.PltPlts() }
func (x PltPltsMrg) PltPlts() *plt.Plts {
	var i0 []*plt.Plts
	for _, cur := range x.I0 {
		i0 = append(i0, cur.PltPlts())
	}
	return x.X.PltPlts().Mrg(i0...)
}
func (x PltPltsPush) Act()             { x.PltPlts() }
func (x PltPltsPush) Ifc() interface{} { return x.PltPlts() }
func (x PltPltsPush) PltPlts() *plt.Plts {
	var i0 []plt.Plt
	for _, cur := range x.I0 {
		i0 = append(i0, cur.PltPlt())
	}
	return x.X.PltPlts().Push(i0...)
}
func (x PltPltsPop) Act()             { x.PltPlt() }
func (x PltPltsPop) Ifc() interface{} { return x.PltPlt() }
func (x PltPltsPop) PltPlt() plt.Plt  { return x.X.PltPlts().Pop() }
func (x PltPltsQue) Act()             { x.PltPlts() }
func (x PltPltsQue) Ifc() interface{} { return x.PltPlts() }
func (x PltPltsQue) PltPlts() *plt.Plts {
	var i0 []plt.Plt
	for _, cur := range x.I0 {
		i0 = append(i0, cur.PltPlt())
	}
	return x.X.PltPlts().Que(i0...)
}
func (x PltPltsDque) Act()                          { x.PltPlt() }
func (x PltPltsDque) Ifc() interface{}              { return x.PltPlt() }
func (x PltPltsDque) PltPlt() plt.Plt               { return x.X.PltPlts().Dque() }
func (x PltPltsIns) Act()                           { x.PltPlts() }
func (x PltPltsIns) Ifc() interface{}               { return x.PltPlts() }
func (x PltPltsIns) PltPlts() *plt.Plts             { return x.X.PltPlts().Ins(x.I0.UntUnt(), x.I1.PltPlt()) }
func (x PltPltsUpd) Act()                           { x.PltPlts() }
func (x PltPltsUpd) Ifc() interface{}               { return x.PltPlts() }
func (x PltPltsUpd) PltPlts() *plt.Plts             { return x.X.PltPlts().Upd(x.I0.UntUnt(), x.I1.PltPlt()) }
func (x PltPltsDel) Act()                           { x.PltPlt() }
func (x PltPltsDel) Ifc() interface{}               { return x.PltPlt() }
func (x PltPltsDel) PltPlt() plt.Plt                { return x.X.PltPlts().Del(x.I0.UntUnt()) }
func (x PltPltsAt) Act()                            { x.PltPlt() }
func (x PltPltsAt) Ifc() interface{}                { return x.PltPlt() }
func (x PltPltsAt) PltPlt() plt.Plt                 { return x.X.PltPlts().At(x.I0.UntUnt()) }
func (x PltPltsIn) Act()                            { x.PltPlts() }
func (x PltPltsIn) Ifc() interface{}                { return x.PltPlts() }
func (x PltPltsIn) PltPlts() *plt.Plts              { return x.X.PltPlts().In(x.I0.UntUnt(), x.I1.UntUnt()) }
func (x PltPltsInBnd) Act()                         { x.PltPlts() }
func (x PltPltsInBnd) Ifc() interface{}             { return x.PltPlts() }
func (x PltPltsInBnd) PltPlts() *plt.Plts           { return x.X.PltPlts().InBnd(x.I0.BndBnd()) }
func (x PltPltsFrom) Act()                          { x.PltPlts() }
func (x PltPltsFrom) Ifc() interface{}              { return x.PltPlts() }
func (x PltPltsFrom) PltPlts() *plt.Plts            { return x.X.PltPlts().From(x.I0.UntUnt()) }
func (x PltPltsTo) Act()                            { x.PltPlts() }
func (x PltPltsTo) Ifc() interface{}                { return x.PltPlts() }
func (x PltPltsTo) PltPlts() *plt.Plts              { return x.X.PltPlts().To(x.I0.UntUnt()) }
func (x PltPltsFst) Act()                           { x.PltPlt() }
func (x PltPltsFst) Ifc() interface{}               { return x.PltPlt() }
func (x PltPltsFst) PltPlt() plt.Plt                { return x.X.PltPlts().Fst() }
func (x PltPltsMdl) Act()                           { x.PltPlt() }
func (x PltPltsMdl) Ifc() interface{}               { return x.PltPlt() }
func (x PltPltsMdl) PltPlt() plt.Plt                { return x.X.PltPlts().Mdl() }
func (x PltPltsLst) Act()                           { x.PltPlt() }
func (x PltPltsLst) Ifc() interface{}               { return x.PltPlt() }
func (x PltPltsLst) PltPlt() plt.Plt                { return x.X.PltPlts().Lst() }
func (x PltPltsFstIdx) Act()                        { x.UntUnt() }
func (x PltPltsFstIdx) Ifc() interface{}            { return x.UntUnt() }
func (x PltPltsFstIdx) UntUnt() unt.Unt             { return x.X.PltPlts().FstIdx() }
func (x PltPltsMdlIdx) Act()                        { x.UntUnt() }
func (x PltPltsMdlIdx) Ifc() interface{}            { return x.UntUnt() }
func (x PltPltsMdlIdx) UntUnt() unt.Unt             { return x.X.PltPlts().MdlIdx() }
func (x PltPltsLstIdx) Act()                        { x.UntUnt() }
func (x PltPltsLstIdx) Ifc() interface{}            { return x.UntUnt() }
func (x PltPltsLstIdx) UntUnt() unt.Unt             { return x.X.PltPlts().LstIdx() }
func (x PltPltsRev) Act()                           { x.PltPlts() }
func (x PltPltsRev) Ifc() interface{}               { return x.PltPlts() }
func (x PltPltsRev) PltPlts() *plt.Plts             { return x.X.PltPlts().Rev() }
func (x PltTmeAxisXVis) Act()                       { x.PltTmeAxisX() }
func (x PltTmeAxisXVis) Ifc() interface{}           { return x.PltTmeAxisX() }
func (x PltTmeAxisXVis) PltTmeAxisX() *plt.TmeAxisX { return x.X.PltTmeAxisX().Vis(x.I0.BolBol()) }
func (x PltFltAxisYVis) Act()                       { x.PltFltAxisY() }
func (x PltFltAxisYVis) Ifc() interface{}           { return x.PltFltAxisY() }
func (x PltFltAxisYVis) PltFltAxisY() *plt.FltAxisY { return x.X.PltFltAxisY().Vis(x.I0.BolBol()) }
func (x PltStmX) Act()                              { x.PltTmeAxisX() }
func (x PltStmX) Ifc() interface{}                  { return x.PltTmeAxisX() }
func (x PltStmX) PltTmeAxisX() *plt.TmeAxisX        { return x.X.PltStm().X() }
func (x PltStmY) Act()                              { x.PltFltAxisY() }
func (x PltStmY) Ifc() interface{}                  { return x.PltFltAxisY() }
func (x PltStmY) PltFltAxisY() *plt.FltAxisY        { return x.X.PltStm().Y() }
func (x PltStmStm) Act()                            { x.PltStm() }
func (x PltStmStm) Ifc() interface{}                { return x.PltStm() }
func (x PltStmStm) PltPlt() plt.Plt                 { return x.PltStm() }
func (x PltStmStm) PltStm() *plt.Stm {
	var i1 []hst.Stm
	for _, cur := range x.I1 {
		i1 = append(i1, cur.HstStm())
	}
	return x.X.PltStm().Stm(x.I0.PenPen(), i1...)
}
func (x PltStmStmBnd) Act()             { x.PltStm() }
func (x PltStmStmBnd) Ifc() interface{} { return x.PltStm() }
func (x PltStmStmBnd) PltPlt() plt.Plt  { return x.PltStm() }
func (x PltStmStmBnd) PltStm() *plt.Stm {
	return x.X.PltStm().StmBnd(x.I0.ClrClr(), x.I1.PenPen(), x.I2.HstStm(), x.I3.HstStm())
}
func (x PltStmCnd) Act()             { x.PltStm() }
func (x PltStmCnd) Ifc() interface{} { return x.PltStm() }
func (x PltStmCnd) PltPlt() plt.Plt  { return x.PltStm() }
func (x PltStmCnd) PltStm() *plt.Stm {
	var i1 []hst.Cnd
	for _, cur := range x.I1 {
		i1 = append(i1, cur.HstCnd())
	}
	return x.X.PltStm().Cnd(x.I0.PenPen(), i1...)
}
func (x PltStmHrzLn) Act()             { x.PltStm() }
func (x PltStmHrzLn) Ifc() interface{} { return x.PltStm() }
func (x PltStmHrzLn) PltPlt() plt.Plt  { return x.PltStm() }
func (x PltStmHrzLn) PltStm() *plt.Stm {
	var i1 []flt.Flt
	for _, cur := range x.I1 {
		i1 = append(i1, cur.FltFlt())
	}
	return x.X.PltStm().HrzLn(x.I0.PenPen(), i1...)
}
func (x PltStmVrtLn) Act()             { x.PltStm() }
func (x PltStmVrtLn) Ifc() interface{} { return x.PltStm() }
func (x PltStmVrtLn) PltPlt() plt.Plt  { return x.PltStm() }
func (x PltStmVrtLn) PltStm() *plt.Stm {
	var i1 []tme.Tme
	for _, cur := range x.I1 {
		i1 = append(i1, cur.TmeTme())
	}
	return x.X.PltStm().VrtLn(x.I0.PenPen(), i1...)
}
func (x PltStmHrzBnd) Act()             { x.PltStm() }
func (x PltStmHrzBnd) Ifc() interface{} { return x.PltStm() }
func (x PltStmHrzBnd) PltPlt() plt.Plt  { return x.PltStm() }
func (x PltStmHrzBnd) PltStm() *plt.Stm {
	return x.X.PltStm().HrzBnd(x.I0.ClrClr(), x.I1.PenPen(), x.I2.FltFlt(), x.I3.FltFlt())
}
func (x PltStmVrtBnd) Act()             { x.PltStm() }
func (x PltStmVrtBnd) Ifc() interface{} { return x.PltStm() }
func (x PltStmVrtBnd) PltPlt() plt.Plt  { return x.PltStm() }
func (x PltStmVrtBnd) PltStm() *plt.Stm {
	return x.X.PltStm().VrtBnd(x.I0.ClrClr(), x.I1.PenPen(), x.I2.TmeTme(), x.I3.TmeTme())
}
func (x PltStmHrzSclVal) Act()             { x.PltStm() }
func (x PltStmHrzSclVal) Ifc() interface{} { return x.PltStm() }
func (x PltStmHrzSclVal) PltPlt() plt.Plt  { return x.PltStm() }
func (x PltStmHrzSclVal) PltStm() *plt.Stm { return x.X.PltStm().HrzSclVal(x.I0.TmeTme()) }
func (x PltStmVrtSclVal) Act()             { x.PltStm() }
func (x PltStmVrtSclVal) Ifc() interface{} { return x.PltStm() }
func (x PltStmVrtSclVal) PltPlt() plt.Plt  { return x.PltStm() }
func (x PltStmVrtSclVal) PltStm() *plt.Stm { return x.X.PltStm().VrtSclVal(x.I0.FltFlt()) }
func (x PltStmSho) Act()                   { x.PltPlt() }
func (x PltStmSho) Ifc() interface{}       { return x.PltPlt() }
func (x PltStmSho) PltPlt() plt.Plt        { return x.X.PltStm().Sho() }
func (x PltStmSiz) Act()                   { x.PltPlt() }
func (x PltStmSiz) Ifc() interface{}       { return x.PltPlt() }
func (x PltStmSiz) PltPlt() plt.Plt        { return x.X.PltStm().Siz(x.I0.UntUnt(), x.I1.UntUnt()) }
func (x PltStmScl) Act()                   { x.PltPlt() }
func (x PltStmScl) Ifc() interface{}       { return x.PltPlt() }
func (x PltStmScl) PltPlt() plt.Plt        { return x.X.PltStm().Scl(x.I0.FltFlt()) }
func (x PltStmHrzScl) Act()                { x.PltPlt() }
func (x PltStmHrzScl) Ifc() interface{}    { return x.PltPlt() }
func (x PltStmHrzScl) PltPlt() plt.Plt     { return x.X.PltStm().HrzScl(x.I0.FltFlt()) }
func (x PltStmVrtScl) Act()                { x.PltPlt() }
func (x PltStmVrtScl) Ifc() interface{}    { return x.PltPlt() }
func (x PltStmVrtScl) PltPlt() plt.Plt     { return x.X.PltStm().VrtScl(x.I0.FltFlt()) }
func (x PltFltsSctrFlts) Act()             { x.PltFltsSctr() }
func (x PltFltsSctrFlts) Ifc() interface{} { return x.PltFltsSctr() }
func (x PltFltsSctrFlts) PltPlt() plt.Plt  { return x.PltFltsSctr() }
func (x PltFltsSctrFlts) PltFltsSctr() *plt.FltsSctr {
	var i1 []*flts.Flts
	for _, cur := range x.I1 {
		i1 = append(i1, cur.FltsFlts())
	}
	return x.X.PltFltsSctr().Flts(x.I0.ClrClr(), i1...)
}
func (x PltFltsSctrPrfLos) Act()             { x.PltFltsSctr() }
func (x PltFltsSctrPrfLos) Ifc() interface{} { return x.PltFltsSctr() }
func (x PltFltsSctrPrfLos) PltPlt() plt.Plt  { return x.PltFltsSctr() }
func (x PltFltsSctrPrfLos) PltFltsSctr() *plt.FltsSctr {
	var i2 []hst.Stm
	for _, cur := range x.I2 {
		i2 = append(i2, cur.HstStm())
	}
	return x.X.PltFltsSctr().PrfLos(x.I0.TmesTmes(), x.I1.TmesTmes(), i2...)
}
func (x PltFltsSctrSho) Act()                  { x.PltPlt() }
func (x PltFltsSctrSho) Ifc() interface{}      { return x.PltPlt() }
func (x PltFltsSctrSho) PltPlt() plt.Plt       { return x.X.PltFltsSctr().Sho() }
func (x PltFltsSctrSiz) Act()                  { x.PltPlt() }
func (x PltFltsSctrSiz) Ifc() interface{}      { return x.PltPlt() }
func (x PltFltsSctrSiz) PltPlt() plt.Plt       { return x.X.PltFltsSctr().Siz(x.I0.UntUnt(), x.I1.UntUnt()) }
func (x PltFltsSctrScl) Act()                  { x.PltPlt() }
func (x PltFltsSctrScl) Ifc() interface{}      { return x.PltPlt() }
func (x PltFltsSctrScl) PltPlt() plt.Plt       { return x.X.PltFltsSctr().Scl(x.I0.FltFlt()) }
func (x PltFltsSctrHrzScl) Act()               { x.PltPlt() }
func (x PltFltsSctrHrzScl) Ifc() interface{}   { return x.PltPlt() }
func (x PltFltsSctrHrzScl) PltPlt() plt.Plt    { return x.X.PltFltsSctr().HrzScl(x.I0.FltFlt()) }
func (x PltFltsSctrVrtScl) Act()               { x.PltPlt() }
func (x PltFltsSctrVrtScl) Ifc() interface{}   { return x.PltPlt() }
func (x PltFltsSctrVrtScl) PltPlt() plt.Plt    { return x.X.PltFltsSctr().VrtScl(x.I0.FltFlt()) }
func (x PltFltsSctrDistFlts) Act()             { x.PltFltsSctrDist() }
func (x PltFltsSctrDistFlts) Ifc() interface{} { return x.PltFltsSctrDist() }
func (x PltFltsSctrDistFlts) PltPlt() plt.Plt  { return x.PltFltsSctrDist() }
func (x PltFltsSctrDistFlts) PltFltsSctrDist() *plt.FltsSctrDist {
	var i2 []*flts.Flts
	for _, cur := range x.I2 {
		i2 = append(i2, cur.FltsFlts())
	}
	return x.X.PltFltsSctrDist().Flts(x.I0.ClrClr(), x.I1.UntUnt(), i2...)
}
func (x PltFltsSctrDistSho) Act()             { x.PltPlt() }
func (x PltFltsSctrDistSho) Ifc() interface{} { return x.PltPlt() }
func (x PltFltsSctrDistSho) PltPlt() plt.Plt  { return x.X.PltFltsSctrDist().Sho() }
func (x PltFltsSctrDistSiz) Act()             { x.PltPlt() }
func (x PltFltsSctrDistSiz) Ifc() interface{} { return x.PltPlt() }
func (x PltFltsSctrDistSiz) PltPlt() plt.Plt {
	return x.X.PltFltsSctrDist().Siz(x.I0.UntUnt(), x.I1.UntUnt())
}
func (x PltFltsSctrDistScl) Act()                { x.PltPlt() }
func (x PltFltsSctrDistScl) Ifc() interface{}    { return x.PltPlt() }
func (x PltFltsSctrDistScl) PltPlt() plt.Plt     { return x.X.PltFltsSctrDist().Scl(x.I0.FltFlt()) }
func (x PltFltsSctrDistHrzScl) Act()             { x.PltPlt() }
func (x PltFltsSctrDistHrzScl) Ifc() interface{} { return x.PltPlt() }
func (x PltFltsSctrDistHrzScl) PltPlt() plt.Plt  { return x.X.PltFltsSctrDist().HrzScl(x.I0.FltFlt()) }
func (x PltFltsSctrDistVrtScl) Act()             { x.PltPlt() }
func (x PltFltsSctrDistVrtScl) Ifc() interface{} { return x.PltPlt() }
func (x PltFltsSctrDistVrtScl) PltPlt() plt.Plt  { return x.X.PltFltsSctrDist().VrtScl(x.I0.FltFlt()) }
func (x PltHrzPlt) Act()                         { x.PltHrz() }
func (x PltHrzPlt) Ifc() interface{}             { return x.PltHrz() }
func (x PltHrzPlt) PltPlt() plt.Plt              { return x.PltHrz() }
func (x PltHrzPlt) PltHrz() *plt.Hrz {
	var i0 []plt.Plt
	for _, cur := range x.I0 {
		i0 = append(i0, cur.PltPlt())
	}
	return x.X.PltHrz().Plt(i0...)
}
func (x PltHrzSho) Act()                { x.PltPlt() }
func (x PltHrzSho) Ifc() interface{}    { return x.PltPlt() }
func (x PltHrzSho) PltPlt() plt.Plt     { return x.X.PltHrz().Sho() }
func (x PltHrzSiz) Act()                { x.PltPlt() }
func (x PltHrzSiz) Ifc() interface{}    { return x.PltPlt() }
func (x PltHrzSiz) PltPlt() plt.Plt     { return x.X.PltHrz().Siz(x.I0.UntUnt(), x.I1.UntUnt()) }
func (x PltHrzScl) Act()                { x.PltPlt() }
func (x PltHrzScl) Ifc() interface{}    { return x.PltPlt() }
func (x PltHrzScl) PltPlt() plt.Plt     { return x.X.PltHrz().Scl(x.I0.FltFlt()) }
func (x PltHrzHrzScl) Act()             { x.PltPlt() }
func (x PltHrzHrzScl) Ifc() interface{} { return x.PltPlt() }
func (x PltHrzHrzScl) PltPlt() plt.Plt  { return x.X.PltHrz().HrzScl(x.I0.FltFlt()) }
func (x PltHrzVrtScl) Act()             { x.PltPlt() }
func (x PltHrzVrtScl) Ifc() interface{} { return x.PltPlt() }
func (x PltHrzVrtScl) PltPlt() plt.Plt  { return x.X.PltHrz().VrtScl(x.I0.FltFlt()) }
func (x PltVrtPlt) Act()                { x.PltVrt() }
func (x PltVrtPlt) Ifc() interface{}    { return x.PltVrt() }
func (x PltVrtPlt) PltPlt() plt.Plt     { return x.PltVrt() }
func (x PltVrtPlt) PltVrt() *plt.Vrt {
	var i0 []plt.Plt
	for _, cur := range x.I0 {
		i0 = append(i0, cur.PltPlt())
	}
	return x.X.PltVrt().Plt(i0...)
}
func (x PltVrtSho) Act()                { x.PltPlt() }
func (x PltVrtSho) Ifc() interface{}    { return x.PltPlt() }
func (x PltVrtSho) PltPlt() plt.Plt     { return x.X.PltVrt().Sho() }
func (x PltVrtSiz) Act()                { x.PltPlt() }
func (x PltVrtSiz) Ifc() interface{}    { return x.PltPlt() }
func (x PltVrtSiz) PltPlt() plt.Plt     { return x.X.PltVrt().Siz(x.I0.UntUnt(), x.I1.UntUnt()) }
func (x PltVrtScl) Act()                { x.PltPlt() }
func (x PltVrtScl) Ifc() interface{}    { return x.PltPlt() }
func (x PltVrtScl) PltPlt() plt.Plt     { return x.X.PltVrt().Scl(x.I0.FltFlt()) }
func (x PltVrtHrzScl) Act()             { x.PltPlt() }
func (x PltVrtHrzScl) Ifc() interface{} { return x.PltPlt() }
func (x PltVrtHrzScl) PltPlt() plt.Plt  { return x.X.PltVrt().HrzScl(x.I0.FltFlt()) }
func (x PltVrtVrtScl) Act()             { x.PltPlt() }
func (x PltVrtVrtScl) Ifc() interface{} { return x.PltPlt() }
func (x PltVrtVrtScl) PltPlt() plt.Plt  { return x.X.PltVrt().VrtScl(x.I0.FltFlt()) }
func (x PltDpthPlt) Act()               { x.PltDpth() }
func (x PltDpthPlt) Ifc() interface{}   { return x.PltDpth() }
func (x PltDpthPlt) PltPlt() plt.Plt    { return x.PltDpth() }
func (x PltDpthPlt) PltDpth() *plt.Dpth {
	var i0 []plt.Plt
	for _, cur := range x.I0 {
		i0 = append(i0, cur.PltPlt())
	}
	return x.X.PltDpth().Plt(i0...)
}
func (x PltDpthSho) Act()                { x.PltPlt() }
func (x PltDpthSho) Ifc() interface{}    { return x.PltPlt() }
func (x PltDpthSho) PltPlt() plt.Plt     { return x.X.PltDpth().Sho() }
func (x PltDpthSiz) Act()                { x.PltPlt() }
func (x PltDpthSiz) Ifc() interface{}    { return x.PltPlt() }
func (x PltDpthSiz) PltPlt() plt.Plt     { return x.X.PltDpth().Siz(x.I0.UntUnt(), x.I1.UntUnt()) }
func (x PltDpthScl) Act()                { x.PltPlt() }
func (x PltDpthScl) Ifc() interface{}    { return x.PltPlt() }
func (x PltDpthScl) PltPlt() plt.Plt     { return x.X.PltDpth().Scl(x.I0.FltFlt()) }
func (x PltDpthHrzScl) Act()             { x.PltPlt() }
func (x PltDpthHrzScl) Ifc() interface{} { return x.PltPlt() }
func (x PltDpthHrzScl) PltPlt() plt.Plt  { return x.X.PltDpth().HrzScl(x.I0.FltFlt()) }
func (x PltDpthVrtScl) Act()             { x.PltPlt() }
func (x PltDpthVrtScl) Ifc() interface{} { return x.PltPlt() }
func (x PltDpthVrtScl) PltPlt() plt.Plt  { return x.X.PltDpth().VrtScl(x.I0.FltFlt()) }
func (x SysMuLck) Act()                  { x.SysMu() }
func (x SysMuLck) Ifc() interface{}      { return x.SysMu() }
func (x SysMuLck) SysMu() *sys.Mu        { return x.X.SysMu().Lck() }
func (x SysMuUlck) Act()                 { x.SysMu() }
func (x SysMuUlck) Ifc() interface{}     { return x.SysMu() }
func (x SysMuUlck) SysMu() *sys.Mu       { return x.X.SysMu().Ulck() }
func (x HstPrvName) Act()                { x.StrStr() }
func (x HstPrvName) Ifc() interface{}    { return x.StrStr() }
func (x HstPrvName) StrStr() str.Str     { return x.X.HstPrv().Name() }
func (x HstPrvEurUsd) Act()              { x.HstInstr() }
func (x HstPrvEurUsd) Ifc() interface{}  { return x.HstInstr() }
func (x HstPrvEurUsd) HstInstr() hst.Instr {
	var i0 []tme.Rng
	for _, cur := range x.I0 {
		i0 = append(i0, cur.TmeRng())
	}
	return x.X.HstPrv().EurUsd(i0...)
}
func (x HstPrvAudUsd) Act()             { x.HstInstr() }
func (x HstPrvAudUsd) Ifc() interface{} { return x.HstInstr() }
func (x HstPrvAudUsd) HstInstr() hst.Instr {
	var i0 []tme.Rng
	for _, cur := range x.I0 {
		i0 = append(i0, cur.TmeRng())
	}
	return x.X.HstPrv().AudUsd(i0...)
}
func (x HstPrvNzdUsd) Act()             { x.HstInstr() }
func (x HstPrvNzdUsd) Ifc() interface{} { return x.HstInstr() }
func (x HstPrvNzdUsd) HstInstr() hst.Instr {
	var i0 []tme.Rng
	for _, cur := range x.I0 {
		i0 = append(i0, cur.TmeRng())
	}
	return x.X.HstPrv().NzdUsd(i0...)
}
func (x HstPrvGbpUsd) Act()             { x.HstInstr() }
func (x HstPrvGbpUsd) Ifc() interface{} { return x.HstInstr() }
func (x HstPrvGbpUsd) HstInstr() hst.Instr {
	var i0 []tme.Rng
	for _, cur := range x.I0 {
		i0 = append(i0, cur.TmeRng())
	}
	return x.X.HstPrv().GbpUsd(i0...)
}
func (x HstInstrName) Act()                 { x.StrStr() }
func (x HstInstrName) Ifc() interface{}     { return x.StrStr() }
func (x HstInstrName) StrStr() str.Str      { return x.X.HstInstr().Name() }
func (x HstInstrI) Act()                    { x.HstInrvl() }
func (x HstInstrI) Ifc() interface{}        { return x.HstInrvl() }
func (x HstInstrI) HstInrvl() hst.Inrvl     { return x.X.HstInstr().I(x.I0.TmeTme()) }
func (x HstInrvlName) Act()                 { x.StrStr() }
func (x HstInrvlName) Ifc() interface{}     { return x.StrStr() }
func (x HstInrvlName) StrStr() str.Str      { return x.X.HstInrvl().Name() }
func (x HstInrvlBid) Act()                  { x.HstSide() }
func (x HstInrvlBid) Ifc() interface{}      { return x.HstSide() }
func (x HstInrvlBid) HstSide() hst.Side     { return x.X.HstInrvl().Bid() }
func (x HstInrvlAsk) Act()                  { x.HstSide() }
func (x HstInrvlAsk) Ifc() interface{}      { return x.HstSide() }
func (x HstInrvlAsk) HstSide() hst.Side     { return x.X.HstInrvl().Ask() }
func (x HstSideName) Act()                  { x.StrStr() }
func (x HstSideName) Ifc() interface{}      { return x.StrStr() }
func (x HstSideName) StrStr() str.Str       { return x.X.HstSide().Name() }
func (x HstSideFst) Act()                   { x.HstStm() }
func (x HstSideFst) Ifc() interface{}       { return x.HstStm() }
func (x HstSideFst) HstStm() hst.Stm        { return x.X.HstSide().Fst() }
func (x HstSideLst) Act()                   { x.HstStm() }
func (x HstSideLst) Ifc() interface{}       { return x.HstStm() }
func (x HstSideLst) HstStm() hst.Stm        { return x.X.HstSide().Lst() }
func (x HstSideSum) Act()                   { x.HstStm() }
func (x HstSideSum) Ifc() interface{}       { return x.HstStm() }
func (x HstSideSum) HstStm() hst.Stm        { return x.X.HstSide().Sum() }
func (x HstSidePrd) Act()                   { x.HstStm() }
func (x HstSidePrd) Ifc() interface{}       { return x.HstStm() }
func (x HstSidePrd) HstStm() hst.Stm        { return x.X.HstSide().Prd() }
func (x HstSideMin) Act()                   { x.HstStm() }
func (x HstSideMin) Ifc() interface{}       { return x.HstStm() }
func (x HstSideMin) HstStm() hst.Stm        { return x.X.HstSide().Min() }
func (x HstSideMax) Act()                   { x.HstStm() }
func (x HstSideMax) Ifc() interface{}       { return x.HstStm() }
func (x HstSideMax) HstStm() hst.Stm        { return x.X.HstSide().Max() }
func (x HstSideMid) Act()                   { x.HstStm() }
func (x HstSideMid) Ifc() interface{}       { return x.HstStm() }
func (x HstSideMid) HstStm() hst.Stm        { return x.X.HstSide().Mid() }
func (x HstSideMdn) Act()                   { x.HstStm() }
func (x HstSideMdn) Ifc() interface{}       { return x.HstStm() }
func (x HstSideMdn) HstStm() hst.Stm        { return x.X.HstSide().Mdn() }
func (x HstSideSma) Act()                   { x.HstStm() }
func (x HstSideSma) Ifc() interface{}       { return x.HstStm() }
func (x HstSideSma) HstStm() hst.Stm        { return x.X.HstSide().Sma() }
func (x HstSideGma) Act()                   { x.HstStm() }
func (x HstSideGma) Ifc() interface{}       { return x.HstStm() }
func (x HstSideGma) HstStm() hst.Stm        { return x.X.HstSide().Gma() }
func (x HstSideWma) Act()                   { x.HstStm() }
func (x HstSideWma) Ifc() interface{}       { return x.HstStm() }
func (x HstSideWma) HstStm() hst.Stm        { return x.X.HstSide().Wma() }
func (x HstSideRsi) Act()                   { x.HstStm() }
func (x HstSideRsi) Ifc() interface{}       { return x.HstStm() }
func (x HstSideRsi) HstStm() hst.Stm        { return x.X.HstSide().Rsi() }
func (x HstSideWrsi) Act()                  { x.HstStm() }
func (x HstSideWrsi) Ifc() interface{}      { return x.HstStm() }
func (x HstSideWrsi) HstStm() hst.Stm       { return x.X.HstSide().Wrsi() }
func (x HstSideAlma) Act()                  { x.HstStm() }
func (x HstSideAlma) Ifc() interface{}      { return x.HstStm() }
func (x HstSideAlma) HstStm() hst.Stm       { return x.X.HstSide().Alma() }
func (x HstSideVrnc) Act()                  { x.HstStm() }
func (x HstSideVrnc) Ifc() interface{}      { return x.HstStm() }
func (x HstSideVrnc) HstStm() hst.Stm       { return x.X.HstSide().Vrnc() }
func (x HstSideStd) Act()                   { x.HstStm() }
func (x HstSideStd) Ifc() interface{}       { return x.HstStm() }
func (x HstSideStd) HstStm() hst.Stm        { return x.X.HstSide().Std() }
func (x HstSideRngFul) Act()                { x.HstStm() }
func (x HstSideRngFul) Ifc() interface{}    { return x.HstStm() }
func (x HstSideRngFul) HstStm() hst.Stm     { return x.X.HstSide().RngFul() }
func (x HstSideRngLst) Act()                { x.HstStm() }
func (x HstSideRngLst) Ifc() interface{}    { return x.HstStm() }
func (x HstSideRngLst) HstStm() hst.Stm     { return x.X.HstSide().RngLst() }
func (x HstSideProLst) Act()                { x.HstStm() }
func (x HstSideProLst) Ifc() interface{}    { return x.HstStm() }
func (x HstSideProLst) HstStm() hst.Stm     { return x.X.HstSide().ProLst() }
func (x HstSideProSma) Act()                { x.HstStm() }
func (x HstSideProSma) Ifc() interface{}    { return x.HstStm() }
func (x HstSideProSma) HstStm() hst.Stm     { return x.X.HstSide().ProSma() }
func (x HstSideProAlma) Act()               { x.HstStm() }
func (x HstSideProAlma) Ifc() interface{}   { return x.HstStm() }
func (x HstSideProAlma) HstStm() hst.Stm    { return x.X.HstSide().ProAlma() }
func (x HstSideSar) Act()                   { x.HstStm() }
func (x HstSideSar) Ifc() interface{}       { return x.HstStm() }
func (x HstSideSar) HstStm() hst.Stm        { return x.X.HstSide().Sar(x.I0.FltFlt(), x.I1.FltFlt()) }
func (x HstSideEma) Act()                   { x.HstStm() }
func (x HstSideEma) Ifc() interface{}       { return x.HstStm() }
func (x HstSideEma) HstStm() hst.Stm        { return x.X.HstSide().Ema() }
func (x HstStmName) Act()                   { x.StrStr() }
func (x HstStmName) Ifc() interface{}       { return x.StrStr() }
func (x HstStmName) StrStr() str.Str        { return x.X.HstStm().Name() }
func (x HstStmAt) Act()                     { x.FltsFlts() }
func (x HstStmAt) Ifc() interface{}         { return x.FltsFlts() }
func (x HstStmAt) FltsFlts() *flts.Flts     { return x.X.HstStm().At(x.I0.TmesTmes()) }
func (x HstStmUnaPos) Act()                 { x.HstStm() }
func (x HstStmUnaPos) Ifc() interface{}     { return x.HstStm() }
func (x HstStmUnaPos) HstStm() hst.Stm      { return x.X.HstStm().UnaPos() }
func (x HstStmUnaNeg) Act()                 { x.HstStm() }
func (x HstStmUnaNeg) Ifc() interface{}     { return x.HstStm() }
func (x HstStmUnaNeg) HstStm() hst.Stm      { return x.X.HstStm().UnaNeg() }
func (x HstStmUnaInv) Act()                 { x.HstStm() }
func (x HstStmUnaInv) Ifc() interface{}     { return x.HstStm() }
func (x HstStmUnaInv) HstStm() hst.Stm      { return x.X.HstStm().UnaInv() }
func (x HstStmUnaSqr) Act()                 { x.HstStm() }
func (x HstStmUnaSqr) Ifc() interface{}     { return x.HstStm() }
func (x HstStmUnaSqr) HstStm() hst.Stm      { return x.X.HstStm().UnaSqr() }
func (x HstStmUnaSqrt) Act()                { x.HstStm() }
func (x HstStmUnaSqrt) Ifc() interface{}    { return x.HstStm() }
func (x HstStmUnaSqrt) HstStm() hst.Stm     { return x.X.HstStm().UnaSqrt() }
func (x HstStmSclAdd) Act()                 { x.HstStm() }
func (x HstStmSclAdd) Ifc() interface{}     { return x.HstStm() }
func (x HstStmSclAdd) HstStm() hst.Stm      { return x.X.HstStm().SclAdd(x.I0.FltFlt()) }
func (x HstStmSclSub) Act()                 { x.HstStm() }
func (x HstStmSclSub) Ifc() interface{}     { return x.HstStm() }
func (x HstStmSclSub) HstStm() hst.Stm      { return x.X.HstStm().SclSub(x.I0.FltFlt()) }
func (x HstStmSclMul) Act()                 { x.HstStm() }
func (x HstStmSclMul) Ifc() interface{}     { return x.HstStm() }
func (x HstStmSclMul) HstStm() hst.Stm      { return x.X.HstStm().SclMul(x.I0.FltFlt()) }
func (x HstStmSclDiv) Act()                 { x.HstStm() }
func (x HstStmSclDiv) Ifc() interface{}     { return x.HstStm() }
func (x HstStmSclDiv) HstStm() hst.Stm      { return x.X.HstStm().SclDiv(x.I0.FltFlt()) }
func (x HstStmSclRem) Act()                 { x.HstStm() }
func (x HstStmSclRem) Ifc() interface{}     { return x.HstStm() }
func (x HstStmSclRem) HstStm() hst.Stm      { return x.X.HstStm().SclRem(x.I0.FltFlt()) }
func (x HstStmSclPow) Act()                 { x.HstStm() }
func (x HstStmSclPow) Ifc() interface{}     { return x.HstStm() }
func (x HstStmSclPow) HstStm() hst.Stm      { return x.X.HstStm().SclPow(x.I0.FltFlt()) }
func (x HstStmSclMin) Act()                 { x.HstStm() }
func (x HstStmSclMin) Ifc() interface{}     { return x.HstStm() }
func (x HstStmSclMin) HstStm() hst.Stm      { return x.X.HstStm().SclMin(x.I0.FltFlt()) }
func (x HstStmSclMax) Act()                 { x.HstStm() }
func (x HstStmSclMax) Ifc() interface{}     { return x.HstStm() }
func (x HstStmSclMax) HstStm() hst.Stm      { return x.X.HstStm().SclMax(x.I0.FltFlt()) }
func (x HstStmSelEql) Act()                 { x.HstStm() }
func (x HstStmSelEql) Ifc() interface{}     { return x.HstStm() }
func (x HstStmSelEql) HstStm() hst.Stm      { return x.X.HstStm().SelEql(x.I0.FltFlt()) }
func (x HstStmSelNeq) Act()                 { x.HstStm() }
func (x HstStmSelNeq) Ifc() interface{}     { return x.HstStm() }
func (x HstStmSelNeq) HstStm() hst.Stm      { return x.X.HstStm().SelNeq(x.I0.FltFlt()) }
func (x HstStmSelLss) Act()                 { x.HstStm() }
func (x HstStmSelLss) Ifc() interface{}     { return x.HstStm() }
func (x HstStmSelLss) HstStm() hst.Stm      { return x.X.HstStm().SelLss(x.I0.FltFlt()) }
func (x HstStmSelGtr) Act()                 { x.HstStm() }
func (x HstStmSelGtr) Ifc() interface{}     { return x.HstStm() }
func (x HstStmSelGtr) HstStm() hst.Stm      { return x.X.HstStm().SelGtr(x.I0.FltFlt()) }
func (x HstStmSelLeq) Act()                 { x.HstStm() }
func (x HstStmSelLeq) Ifc() interface{}     { return x.HstStm() }
func (x HstStmSelLeq) HstStm() hst.Stm      { return x.X.HstStm().SelLeq(x.I0.FltFlt()) }
func (x HstStmSelGeq) Act()                 { x.HstStm() }
func (x HstStmSelGeq) Ifc() interface{}     { return x.HstStm() }
func (x HstStmSelGeq) HstStm() hst.Stm      { return x.X.HstStm().SelGeq(x.I0.FltFlt()) }
func (x HstStmAggFst) Act()                 { x.HstStm() }
func (x HstStmAggFst) Ifc() interface{}     { return x.HstStm() }
func (x HstStmAggFst) HstStm() hst.Stm      { return x.X.HstStm().AggFst(x.I0.UntUnt()) }
func (x HstStmAggLst) Act()                 { x.HstStm() }
func (x HstStmAggLst) Ifc() interface{}     { return x.HstStm() }
func (x HstStmAggLst) HstStm() hst.Stm      { return x.X.HstStm().AggLst(x.I0.UntUnt()) }
func (x HstStmAggSum) Act()                 { x.HstStm() }
func (x HstStmAggSum) Ifc() interface{}     { return x.HstStm() }
func (x HstStmAggSum) HstStm() hst.Stm      { return x.X.HstStm().AggSum(x.I0.UntUnt()) }
func (x HstStmAggPrd) Act()                 { x.HstStm() }
func (x HstStmAggPrd) Ifc() interface{}     { return x.HstStm() }
func (x HstStmAggPrd) HstStm() hst.Stm      { return x.X.HstStm().AggPrd(x.I0.UntUnt()) }
func (x HstStmAggMin) Act()                 { x.HstStm() }
func (x HstStmAggMin) Ifc() interface{}     { return x.HstStm() }
func (x HstStmAggMin) HstStm() hst.Stm      { return x.X.HstStm().AggMin(x.I0.UntUnt()) }
func (x HstStmAggMax) Act()                 { x.HstStm() }
func (x HstStmAggMax) Ifc() interface{}     { return x.HstStm() }
func (x HstStmAggMax) HstStm() hst.Stm      { return x.X.HstStm().AggMax(x.I0.UntUnt()) }
func (x HstStmAggMid) Act()                 { x.HstStm() }
func (x HstStmAggMid) Ifc() interface{}     { return x.HstStm() }
func (x HstStmAggMid) HstStm() hst.Stm      { return x.X.HstStm().AggMid(x.I0.UntUnt()) }
func (x HstStmAggMdn) Act()                 { x.HstStm() }
func (x HstStmAggMdn) Ifc() interface{}     { return x.HstStm() }
func (x HstStmAggMdn) HstStm() hst.Stm      { return x.X.HstStm().AggMdn(x.I0.UntUnt()) }
func (x HstStmAggSma) Act()                 { x.HstStm() }
func (x HstStmAggSma) Ifc() interface{}     { return x.HstStm() }
func (x HstStmAggSma) HstStm() hst.Stm      { return x.X.HstStm().AggSma(x.I0.UntUnt()) }
func (x HstStmAggGma) Act()                 { x.HstStm() }
func (x HstStmAggGma) Ifc() interface{}     { return x.HstStm() }
func (x HstStmAggGma) HstStm() hst.Stm      { return x.X.HstStm().AggGma(x.I0.UntUnt()) }
func (x HstStmAggWma) Act()                 { x.HstStm() }
func (x HstStmAggWma) Ifc() interface{}     { return x.HstStm() }
func (x HstStmAggWma) HstStm() hst.Stm      { return x.X.HstStm().AggWma(x.I0.UntUnt()) }
func (x HstStmAggRsi) Act()                 { x.HstStm() }
func (x HstStmAggRsi) Ifc() interface{}     { return x.HstStm() }
func (x HstStmAggRsi) HstStm() hst.Stm      { return x.X.HstStm().AggRsi(x.I0.UntUnt()) }
func (x HstStmAggWrsi) Act()                { x.HstStm() }
func (x HstStmAggWrsi) Ifc() interface{}    { return x.HstStm() }
func (x HstStmAggWrsi) HstStm() hst.Stm     { return x.X.HstStm().AggWrsi(x.I0.UntUnt()) }
func (x HstStmAggAlma) Act()                { x.HstStm() }
func (x HstStmAggAlma) Ifc() interface{}    { return x.HstStm() }
func (x HstStmAggAlma) HstStm() hst.Stm     { return x.X.HstStm().AggAlma(x.I0.UntUnt()) }
func (x HstStmAggVrnc) Act()                { x.HstStm() }
func (x HstStmAggVrnc) Ifc() interface{}    { return x.HstStm() }
func (x HstStmAggVrnc) HstStm() hst.Stm     { return x.X.HstStm().AggVrnc(x.I0.UntUnt()) }
func (x HstStmAggStd) Act()                 { x.HstStm() }
func (x HstStmAggStd) Ifc() interface{}     { return x.HstStm() }
func (x HstStmAggStd) HstStm() hst.Stm      { return x.X.HstStm().AggStd(x.I0.UntUnt()) }
func (x HstStmAggRngFul) Act()              { x.HstStm() }
func (x HstStmAggRngFul) Ifc() interface{}  { return x.HstStm() }
func (x HstStmAggRngFul) HstStm() hst.Stm   { return x.X.HstStm().AggRngFul(x.I0.UntUnt()) }
func (x HstStmAggRngLst) Act()              { x.HstStm() }
func (x HstStmAggRngLst) Ifc() interface{}  { return x.HstStm() }
func (x HstStmAggRngLst) HstStm() hst.Stm   { return x.X.HstStm().AggRngLst(x.I0.UntUnt()) }
func (x HstStmAggProLst) Act()              { x.HstStm() }
func (x HstStmAggProLst) Ifc() interface{}  { return x.HstStm() }
func (x HstStmAggProLst) HstStm() hst.Stm   { return x.X.HstStm().AggProLst(x.I0.UntUnt()) }
func (x HstStmAggProSma) Act()              { x.HstStm() }
func (x HstStmAggProSma) Ifc() interface{}  { return x.HstStm() }
func (x HstStmAggProSma) HstStm() hst.Stm   { return x.X.HstStm().AggProSma(x.I0.UntUnt()) }
func (x HstStmAggProAlma) Act()             { x.HstStm() }
func (x HstStmAggProAlma) Ifc() interface{} { return x.HstStm() }
func (x HstStmAggProAlma) HstStm() hst.Stm  { return x.X.HstStm().AggProAlma(x.I0.UntUnt()) }
func (x HstStmAggEma) Act()                 { x.HstStm() }
func (x HstStmAggEma) Ifc() interface{}     { return x.HstStm() }
func (x HstStmAggEma) HstStm() hst.Stm      { return x.X.HstStm().AggEma(x.I0.UntUnt()) }
func (x HstStmInrAdd) Act()                 { x.HstStm() }
func (x HstStmInrAdd) Ifc() interface{}     { return x.HstStm() }
func (x HstStmInrAdd) HstStm() hst.Stm      { return x.X.HstStm().InrAdd(x.I0.UntUnt()) }
func (x HstStmInrSub) Act()                 { x.HstStm() }
func (x HstStmInrSub) Ifc() interface{}     { return x.HstStm() }
func (x HstStmInrSub) HstStm() hst.Stm      { return x.X.HstStm().InrSub(x.I0.UntUnt()) }
func (x HstStmInrMul) Act()                 { x.HstStm() }
func (x HstStmInrMul) Ifc() interface{}     { return x.HstStm() }
func (x HstStmInrMul) HstStm() hst.Stm      { return x.X.HstStm().InrMul(x.I0.UntUnt()) }
func (x HstStmInrDiv) Act()                 { x.HstStm() }
func (x HstStmInrDiv) Ifc() interface{}     { return x.HstStm() }
func (x HstStmInrDiv) HstStm() hst.Stm      { return x.X.HstStm().InrDiv(x.I0.UntUnt()) }
func (x HstStmInrRem) Act()                 { x.HstStm() }
func (x HstStmInrRem) Ifc() interface{}     { return x.HstStm() }
func (x HstStmInrRem) HstStm() hst.Stm      { return x.X.HstStm().InrRem(x.I0.UntUnt()) }
func (x HstStmInrPow) Act()                 { x.HstStm() }
func (x HstStmInrPow) Ifc() interface{}     { return x.HstStm() }
func (x HstStmInrPow) HstStm() hst.Stm      { return x.X.HstStm().InrPow(x.I0.UntUnt()) }
func (x HstStmInrMin) Act()                 { x.HstStm() }
func (x HstStmInrMin) Ifc() interface{}     { return x.HstStm() }
func (x HstStmInrMin) HstStm() hst.Stm      { return x.X.HstStm().InrMin(x.I0.UntUnt()) }
func (x HstStmInrMax) Act()                 { x.HstStm() }
func (x HstStmInrMax) Ifc() interface{}     { return x.HstStm() }
func (x HstStmInrMax) HstStm() hst.Stm      { return x.X.HstStm().InrMax(x.I0.UntUnt()) }
func (x HstStmInrSlp) Act()                 { x.HstStm() }
func (x HstStmInrSlp) Ifc() interface{}     { return x.HstStm() }
func (x HstStmInrSlp) HstStm() hst.Stm      { return x.X.HstStm().InrSlp(x.I0.UntUnt()) }
func (x HstStmOtrAdd) Act()                 { x.HstStm() }
func (x HstStmOtrAdd) Ifc() interface{}     { return x.HstStm() }
func (x HstStmOtrAdd) HstStm() hst.Stm      { return x.X.HstStm().OtrAdd(x.I0.UntUnt(), x.I1.HstStm()) }
func (x HstStmOtrSub) Act()                 { x.HstStm() }
func (x HstStmOtrSub) Ifc() interface{}     { return x.HstStm() }
func (x HstStmOtrSub) HstStm() hst.Stm      { return x.X.HstStm().OtrSub(x.I0.UntUnt(), x.I1.HstStm()) }
func (x HstStmOtrMul) Act()                 { x.HstStm() }
func (x HstStmOtrMul) Ifc() interface{}     { return x.HstStm() }
func (x HstStmOtrMul) HstStm() hst.Stm      { return x.X.HstStm().OtrMul(x.I0.UntUnt(), x.I1.HstStm()) }
func (x HstStmOtrDiv) Act()                 { x.HstStm() }
func (x HstStmOtrDiv) Ifc() interface{}     { return x.HstStm() }
func (x HstStmOtrDiv) HstStm() hst.Stm      { return x.X.HstStm().OtrDiv(x.I0.UntUnt(), x.I1.HstStm()) }
func (x HstStmOtrRem) Act()                 { x.HstStm() }
func (x HstStmOtrRem) Ifc() interface{}     { return x.HstStm() }
func (x HstStmOtrRem) HstStm() hst.Stm      { return x.X.HstStm().OtrRem(x.I0.UntUnt(), x.I1.HstStm()) }
func (x HstStmOtrPow) Act()                 { x.HstStm() }
func (x HstStmOtrPow) Ifc() interface{}     { return x.HstStm() }
func (x HstStmOtrPow) HstStm() hst.Stm      { return x.X.HstStm().OtrPow(x.I0.UntUnt(), x.I1.HstStm()) }
func (x HstStmOtrMin) Act()                 { x.HstStm() }
func (x HstStmOtrMin) Ifc() interface{}     { return x.HstStm() }
func (x HstStmOtrMin) HstStm() hst.Stm      { return x.X.HstStm().OtrMin(x.I0.UntUnt(), x.I1.HstStm()) }
func (x HstStmOtrMax) Act()                 { x.HstStm() }
func (x HstStmOtrMax) Ifc() interface{}     { return x.HstStm() }
func (x HstStmOtrMax) HstStm() hst.Stm      { return x.X.HstStm().OtrMax(x.I0.UntUnt(), x.I1.HstStm()) }
func (x HstStmSclEql) Act()                 { x.HstCnd() }
func (x HstStmSclEql) Ifc() interface{}     { return x.HstCnd() }
func (x HstStmSclEql) HstCnd() hst.Cnd      { return x.X.HstStm().SclEql(x.I0.FltFlt()) }
func (x HstStmSclNeq) Act()                 { x.HstCnd() }
func (x HstStmSclNeq) Ifc() interface{}     { return x.HstCnd() }
func (x HstStmSclNeq) HstCnd() hst.Cnd      { return x.X.HstStm().SclNeq(x.I0.FltFlt()) }
func (x HstStmSclLss) Act()                 { x.HstCnd() }
func (x HstStmSclLss) Ifc() interface{}     { return x.HstCnd() }
func (x HstStmSclLss) HstCnd() hst.Cnd      { return x.X.HstStm().SclLss(x.I0.FltFlt()) }
func (x HstStmSclGtr) Act()                 { x.HstCnd() }
func (x HstStmSclGtr) Ifc() interface{}     { return x.HstCnd() }
func (x HstStmSclGtr) HstCnd() hst.Cnd      { return x.X.HstStm().SclGtr(x.I0.FltFlt()) }
func (x HstStmSclLeq) Act()                 { x.HstCnd() }
func (x HstStmSclLeq) Ifc() interface{}     { return x.HstCnd() }
func (x HstStmSclLeq) HstCnd() hst.Cnd      { return x.X.HstStm().SclLeq(x.I0.FltFlt()) }
func (x HstStmSclGeq) Act()                 { x.HstCnd() }
func (x HstStmSclGeq) Ifc() interface{}     { return x.HstCnd() }
func (x HstStmSclGeq) HstCnd() hst.Cnd      { return x.X.HstStm().SclGeq(x.I0.FltFlt()) }
func (x HstStmInrEql) Act()                 { x.HstCnd() }
func (x HstStmInrEql) Ifc() interface{}     { return x.HstCnd() }
func (x HstStmInrEql) HstCnd() hst.Cnd      { return x.X.HstStm().InrEql(x.I0.UntUnt()) }
func (x HstStmInrNeq) Act()                 { x.HstCnd() }
func (x HstStmInrNeq) Ifc() interface{}     { return x.HstCnd() }
func (x HstStmInrNeq) HstCnd() hst.Cnd      { return x.X.HstStm().InrNeq(x.I0.UntUnt()) }
func (x HstStmInrLss) Act()                 { x.HstCnd() }
func (x HstStmInrLss) Ifc() interface{}     { return x.HstCnd() }
func (x HstStmInrLss) HstCnd() hst.Cnd      { return x.X.HstStm().InrLss(x.I0.UntUnt()) }
func (x HstStmInrGtr) Act()                 { x.HstCnd() }
func (x HstStmInrGtr) Ifc() interface{}     { return x.HstCnd() }
func (x HstStmInrGtr) HstCnd() hst.Cnd      { return x.X.HstStm().InrGtr(x.I0.UntUnt()) }
func (x HstStmInrLeq) Act()                 { x.HstCnd() }
func (x HstStmInrLeq) Ifc() interface{}     { return x.HstCnd() }
func (x HstStmInrLeq) HstCnd() hst.Cnd      { return x.X.HstStm().InrLeq(x.I0.UntUnt()) }
func (x HstStmInrGeq) Act()                 { x.HstCnd() }
func (x HstStmInrGeq) Ifc() interface{}     { return x.HstCnd() }
func (x HstStmInrGeq) HstCnd() hst.Cnd      { return x.X.HstStm().InrGeq(x.I0.UntUnt()) }
func (x HstStmOtrEql) Act()                 { x.HstCnd() }
func (x HstStmOtrEql) Ifc() interface{}     { return x.HstCnd() }
func (x HstStmOtrEql) HstCnd() hst.Cnd      { return x.X.HstStm().OtrEql(x.I0.UntUnt(), x.I1.HstStm()) }
func (x HstStmOtrNeq) Act()                 { x.HstCnd() }
func (x HstStmOtrNeq) Ifc() interface{}     { return x.HstCnd() }
func (x HstStmOtrNeq) HstCnd() hst.Cnd      { return x.X.HstStm().OtrNeq(x.I0.UntUnt(), x.I1.HstStm()) }
func (x HstStmOtrLss) Act()                 { x.HstCnd() }
func (x HstStmOtrLss) Ifc() interface{}     { return x.HstCnd() }
func (x HstStmOtrLss) HstCnd() hst.Cnd      { return x.X.HstStm().OtrLss(x.I0.UntUnt(), x.I1.HstStm()) }
func (x HstStmOtrGtr) Act()                 { x.HstCnd() }
func (x HstStmOtrGtr) Ifc() interface{}     { return x.HstCnd() }
func (x HstStmOtrGtr) HstCnd() hst.Cnd      { return x.X.HstStm().OtrGtr(x.I0.UntUnt(), x.I1.HstStm()) }
func (x HstStmOtrLeq) Act()                 { x.HstCnd() }
func (x HstStmOtrLeq) Ifc() interface{}     { return x.HstCnd() }
func (x HstStmOtrLeq) HstCnd() hst.Cnd      { return x.X.HstStm().OtrLeq(x.I0.UntUnt(), x.I1.HstStm()) }
func (x HstStmOtrGeq) Act()                 { x.HstCnd() }
func (x HstStmOtrGeq) Ifc() interface{}     { return x.HstCnd() }
func (x HstStmOtrGeq) HstCnd() hst.Cnd      { return x.X.HstStm().OtrGeq(x.I0.UntUnt(), x.I1.HstStm()) }
func (x HstCndName) Act()                   { x.StrStr() }
func (x HstCndName) Ifc() interface{}       { return x.StrStr() }
func (x HstCndName) StrStr() str.Str        { return x.X.HstCnd().Name() }
func (x HstCndAnd) Act()                    { x.HstCnd() }
func (x HstCndAnd) Ifc() interface{}        { return x.HstCnd() }
func (x HstCndAnd) HstCnd() hst.Cnd         { return x.X.HstCnd().And(x.I0.HstCnd()) }
func (x HstCndSeq) Act()                    { x.HstCnd() }
func (x HstCndSeq) Ifc() interface{}        { return x.HstCnd() }
func (x HstCndSeq) HstCnd() hst.Cnd         { return x.X.HstCnd().Seq(x.I0.TmeTme(), x.I1.HstCnd()) }
func (x HstCndStgy) Act()                   { x.HstStgy() }
func (x HstCndStgy) Ifc() interface{}       { return x.HstStgy() }
func (x HstCndStgy) HstStgy() hst.Stgy {
	var i7 []hst.Cnd
	for _, cur := range x.I7 {
		i7 = append(i7, cur.HstCnd())
	}
	return x.X.HstCnd().Stgy(x.I0.BolBol(), x.I1.FltFlt(), x.I2.FltFlt(), x.I3.TmeTme(), x.I4.FltFlt(), x.I5.HstInstr(), x.I6.HstStms(), i7...)
}
func (x HstStgyName) Act()              { x.StrStr() }
func (x HstStgyName) Ifc() interface{}  { return x.StrStr() }
func (x HstStgyName) StrStr() str.Str   { return x.X.HstStgy().Name() }
func (x RltPrvMayTrd) Act()             { x.BolBol() }
func (x RltPrvMayTrd) Ifc() interface{} { return x.BolBol() }
func (x RltPrvMayTrd) BolBol() bol.Bol  { return x.X.RltPrv().MayTrd() }
func (x RltPrvEurUsd) Act()             { x.RltInstr() }
func (x RltPrvEurUsd) Ifc() interface{} { return x.RltInstr() }
func (x RltPrvEurUsd) RltInstr() rlt.Instr {
	var i0 []tme.Rng
	for _, cur := range x.I0 {
		i0 = append(i0, cur.TmeRng())
	}
	return x.X.RltPrv().EurUsd(i0...)
}
func (x RltPrvAudUsd) Act()             { x.RltInstr() }
func (x RltPrvAudUsd) Ifc() interface{} { return x.RltInstr() }
func (x RltPrvAudUsd) RltInstr() rlt.Instr {
	var i0 []tme.Rng
	for _, cur := range x.I0 {
		i0 = append(i0, cur.TmeRng())
	}
	return x.X.RltPrv().AudUsd(i0...)
}
func (x RltPrvNzdUsd) Act()             { x.RltInstr() }
func (x RltPrvNzdUsd) Ifc() interface{} { return x.RltInstr() }
func (x RltPrvNzdUsd) RltInstr() rlt.Instr {
	var i0 []tme.Rng
	for _, cur := range x.I0 {
		i0 = append(i0, cur.TmeRng())
	}
	return x.X.RltPrv().NzdUsd(i0...)
}
func (x RltPrvGbpUsd) Act()             { x.RltInstr() }
func (x RltPrvGbpUsd) Ifc() interface{} { return x.RltInstr() }
func (x RltPrvGbpUsd) RltInstr() rlt.Instr {
	var i0 []tme.Rng
	for _, cur := range x.I0 {
		i0 = append(i0, cur.TmeRng())
	}
	return x.X.RltPrv().GbpUsd(i0...)
}
func (x RltInstrI) Act()                    { x.RltInrvl() }
func (x RltInstrI) Ifc() interface{}        { return x.RltInrvl() }
func (x RltInstrI) RltInrvl() rlt.Inrvl     { return x.X.RltInstr().I(x.I0.TmeTme()) }
func (x RltInrvlBid) Act()                  { x.RltSide() }
func (x RltInrvlBid) Ifc() interface{}      { return x.RltSide() }
func (x RltInrvlBid) RltSide() rlt.Side     { return x.X.RltInrvl().Bid() }
func (x RltInrvlAsk) Act()                  { x.RltSide() }
func (x RltInrvlAsk) Ifc() interface{}      { return x.RltSide() }
func (x RltInrvlAsk) RltSide() rlt.Side     { return x.X.RltInrvl().Ask() }
func (x RltSideFst) Act()                   { x.RltStm() }
func (x RltSideFst) Ifc() interface{}       { return x.RltStm() }
func (x RltSideFst) RltStm() rlt.Stm        { return x.X.RltSide().Fst() }
func (x RltSideLst) Act()                   { x.RltStm() }
func (x RltSideLst) Ifc() interface{}       { return x.RltStm() }
func (x RltSideLst) RltStm() rlt.Stm        { return x.X.RltSide().Lst() }
func (x RltSideSum) Act()                   { x.RltStm() }
func (x RltSideSum) Ifc() interface{}       { return x.RltStm() }
func (x RltSideSum) RltStm() rlt.Stm        { return x.X.RltSide().Sum() }
func (x RltSidePrd) Act()                   { x.RltStm() }
func (x RltSidePrd) Ifc() interface{}       { return x.RltStm() }
func (x RltSidePrd) RltStm() rlt.Stm        { return x.X.RltSide().Prd() }
func (x RltSideMin) Act()                   { x.RltStm() }
func (x RltSideMin) Ifc() interface{}       { return x.RltStm() }
func (x RltSideMin) RltStm() rlt.Stm        { return x.X.RltSide().Min() }
func (x RltSideMax) Act()                   { x.RltStm() }
func (x RltSideMax) Ifc() interface{}       { return x.RltStm() }
func (x RltSideMax) RltStm() rlt.Stm        { return x.X.RltSide().Max() }
func (x RltSideMid) Act()                   { x.RltStm() }
func (x RltSideMid) Ifc() interface{}       { return x.RltStm() }
func (x RltSideMid) RltStm() rlt.Stm        { return x.X.RltSide().Mid() }
func (x RltSideMdn) Act()                   { x.RltStm() }
func (x RltSideMdn) Ifc() interface{}       { return x.RltStm() }
func (x RltSideMdn) RltStm() rlt.Stm        { return x.X.RltSide().Mdn() }
func (x RltSideSma) Act()                   { x.RltStm() }
func (x RltSideSma) Ifc() interface{}       { return x.RltStm() }
func (x RltSideSma) RltStm() rlt.Stm        { return x.X.RltSide().Sma() }
func (x RltSideGma) Act()                   { x.RltStm() }
func (x RltSideGma) Ifc() interface{}       { return x.RltStm() }
func (x RltSideGma) RltStm() rlt.Stm        { return x.X.RltSide().Gma() }
func (x RltSideWma) Act()                   { x.RltStm() }
func (x RltSideWma) Ifc() interface{}       { return x.RltStm() }
func (x RltSideWma) RltStm() rlt.Stm        { return x.X.RltSide().Wma() }
func (x RltSideRsi) Act()                   { x.RltStm() }
func (x RltSideRsi) Ifc() interface{}       { return x.RltStm() }
func (x RltSideRsi) RltStm() rlt.Stm        { return x.X.RltSide().Rsi() }
func (x RltSideWrsi) Act()                  { x.RltStm() }
func (x RltSideWrsi) Ifc() interface{}      { return x.RltStm() }
func (x RltSideWrsi) RltStm() rlt.Stm       { return x.X.RltSide().Wrsi() }
func (x RltSideAlma) Act()                  { x.RltStm() }
func (x RltSideAlma) Ifc() interface{}      { return x.RltStm() }
func (x RltSideAlma) RltStm() rlt.Stm       { return x.X.RltSide().Alma() }
func (x RltSideVrnc) Act()                  { x.RltStm() }
func (x RltSideVrnc) Ifc() interface{}      { return x.RltStm() }
func (x RltSideVrnc) RltStm() rlt.Stm       { return x.X.RltSide().Vrnc() }
func (x RltSideStd) Act()                   { x.RltStm() }
func (x RltSideStd) Ifc() interface{}       { return x.RltStm() }
func (x RltSideStd) RltStm() rlt.Stm        { return x.X.RltSide().Std() }
func (x RltSideRngFul) Act()                { x.RltStm() }
func (x RltSideRngFul) Ifc() interface{}    { return x.RltStm() }
func (x RltSideRngFul) RltStm() rlt.Stm     { return x.X.RltSide().RngFul() }
func (x RltSideRngLst) Act()                { x.RltStm() }
func (x RltSideRngLst) Ifc() interface{}    { return x.RltStm() }
func (x RltSideRngLst) RltStm() rlt.Stm     { return x.X.RltSide().RngLst() }
func (x RltSideProLst) Act()                { x.RltStm() }
func (x RltSideProLst) Ifc() interface{}    { return x.RltStm() }
func (x RltSideProLst) RltStm() rlt.Stm     { return x.X.RltSide().ProLst() }
func (x RltSideProSma) Act()                { x.RltStm() }
func (x RltSideProSma) Ifc() interface{}    { return x.RltStm() }
func (x RltSideProSma) RltStm() rlt.Stm     { return x.X.RltSide().ProSma() }
func (x RltSideProAlma) Act()               { x.RltStm() }
func (x RltSideProAlma) Ifc() interface{}   { return x.RltStm() }
func (x RltSideProAlma) RltStm() rlt.Stm    { return x.X.RltSide().ProAlma() }
func (x RltSideSar) Act()                   { x.RltStm() }
func (x RltSideSar) Ifc() interface{}       { return x.RltStm() }
func (x RltSideSar) RltStm() rlt.Stm        { return x.X.RltSide().Sar(x.I0.FltFlt(), x.I1.FltFlt()) }
func (x RltSideEma) Act()                   { x.RltStm() }
func (x RltSideEma) Ifc() interface{}       { return x.RltStm() }
func (x RltSideEma) RltStm() rlt.Stm        { return x.X.RltSide().Ema() }
func (x RltStmUnaPos) Act()                 { x.RltStm() }
func (x RltStmUnaPos) Ifc() interface{}     { return x.RltStm() }
func (x RltStmUnaPos) RltStm() rlt.Stm      { return x.X.RltStm().UnaPos() }
func (x RltStmUnaNeg) Act()                 { x.RltStm() }
func (x RltStmUnaNeg) Ifc() interface{}     { return x.RltStm() }
func (x RltStmUnaNeg) RltStm() rlt.Stm      { return x.X.RltStm().UnaNeg() }
func (x RltStmUnaInv) Act()                 { x.RltStm() }
func (x RltStmUnaInv) Ifc() interface{}     { return x.RltStm() }
func (x RltStmUnaInv) RltStm() rlt.Stm      { return x.X.RltStm().UnaInv() }
func (x RltStmUnaSqr) Act()                 { x.RltStm() }
func (x RltStmUnaSqr) Ifc() interface{}     { return x.RltStm() }
func (x RltStmUnaSqr) RltStm() rlt.Stm      { return x.X.RltStm().UnaSqr() }
func (x RltStmUnaSqrt) Act()                { x.RltStm() }
func (x RltStmUnaSqrt) Ifc() interface{}    { return x.RltStm() }
func (x RltStmUnaSqrt) RltStm() rlt.Stm     { return x.X.RltStm().UnaSqrt() }
func (x RltStmSclAdd) Act()                 { x.RltStm() }
func (x RltStmSclAdd) Ifc() interface{}     { return x.RltStm() }
func (x RltStmSclAdd) RltStm() rlt.Stm      { return x.X.RltStm().SclAdd(x.I0.FltFlt()) }
func (x RltStmSclSub) Act()                 { x.RltStm() }
func (x RltStmSclSub) Ifc() interface{}     { return x.RltStm() }
func (x RltStmSclSub) RltStm() rlt.Stm      { return x.X.RltStm().SclSub(x.I0.FltFlt()) }
func (x RltStmSclMul) Act()                 { x.RltStm() }
func (x RltStmSclMul) Ifc() interface{}     { return x.RltStm() }
func (x RltStmSclMul) RltStm() rlt.Stm      { return x.X.RltStm().SclMul(x.I0.FltFlt()) }
func (x RltStmSclDiv) Act()                 { x.RltStm() }
func (x RltStmSclDiv) Ifc() interface{}     { return x.RltStm() }
func (x RltStmSclDiv) RltStm() rlt.Stm      { return x.X.RltStm().SclDiv(x.I0.FltFlt()) }
func (x RltStmSclRem) Act()                 { x.RltStm() }
func (x RltStmSclRem) Ifc() interface{}     { return x.RltStm() }
func (x RltStmSclRem) RltStm() rlt.Stm      { return x.X.RltStm().SclRem(x.I0.FltFlt()) }
func (x RltStmSclPow) Act()                 { x.RltStm() }
func (x RltStmSclPow) Ifc() interface{}     { return x.RltStm() }
func (x RltStmSclPow) RltStm() rlt.Stm      { return x.X.RltStm().SclPow(x.I0.FltFlt()) }
func (x RltStmSclMin) Act()                 { x.RltStm() }
func (x RltStmSclMin) Ifc() interface{}     { return x.RltStm() }
func (x RltStmSclMin) RltStm() rlt.Stm      { return x.X.RltStm().SclMin(x.I0.FltFlt()) }
func (x RltStmSclMax) Act()                 { x.RltStm() }
func (x RltStmSclMax) Ifc() interface{}     { return x.RltStm() }
func (x RltStmSclMax) RltStm() rlt.Stm      { return x.X.RltStm().SclMax(x.I0.FltFlt()) }
func (x RltStmSelEql) Act()                 { x.RltStm() }
func (x RltStmSelEql) Ifc() interface{}     { return x.RltStm() }
func (x RltStmSelEql) RltStm() rlt.Stm      { return x.X.RltStm().SelEql(x.I0.FltFlt()) }
func (x RltStmSelNeq) Act()                 { x.RltStm() }
func (x RltStmSelNeq) Ifc() interface{}     { return x.RltStm() }
func (x RltStmSelNeq) RltStm() rlt.Stm      { return x.X.RltStm().SelNeq(x.I0.FltFlt()) }
func (x RltStmSelLss) Act()                 { x.RltStm() }
func (x RltStmSelLss) Ifc() interface{}     { return x.RltStm() }
func (x RltStmSelLss) RltStm() rlt.Stm      { return x.X.RltStm().SelLss(x.I0.FltFlt()) }
func (x RltStmSelGtr) Act()                 { x.RltStm() }
func (x RltStmSelGtr) Ifc() interface{}     { return x.RltStm() }
func (x RltStmSelGtr) RltStm() rlt.Stm      { return x.X.RltStm().SelGtr(x.I0.FltFlt()) }
func (x RltStmSelLeq) Act()                 { x.RltStm() }
func (x RltStmSelLeq) Ifc() interface{}     { return x.RltStm() }
func (x RltStmSelLeq) RltStm() rlt.Stm      { return x.X.RltStm().SelLeq(x.I0.FltFlt()) }
func (x RltStmSelGeq) Act()                 { x.RltStm() }
func (x RltStmSelGeq) Ifc() interface{}     { return x.RltStm() }
func (x RltStmSelGeq) RltStm() rlt.Stm      { return x.X.RltStm().SelGeq(x.I0.FltFlt()) }
func (x RltStmAggFst) Act()                 { x.RltStm() }
func (x RltStmAggFst) Ifc() interface{}     { return x.RltStm() }
func (x RltStmAggFst) RltStm() rlt.Stm      { return x.X.RltStm().AggFst(x.I0.UntUnt()) }
func (x RltStmAggLst) Act()                 { x.RltStm() }
func (x RltStmAggLst) Ifc() interface{}     { return x.RltStm() }
func (x RltStmAggLst) RltStm() rlt.Stm      { return x.X.RltStm().AggLst(x.I0.UntUnt()) }
func (x RltStmAggSum) Act()                 { x.RltStm() }
func (x RltStmAggSum) Ifc() interface{}     { return x.RltStm() }
func (x RltStmAggSum) RltStm() rlt.Stm      { return x.X.RltStm().AggSum(x.I0.UntUnt()) }
func (x RltStmAggPrd) Act()                 { x.RltStm() }
func (x RltStmAggPrd) Ifc() interface{}     { return x.RltStm() }
func (x RltStmAggPrd) RltStm() rlt.Stm      { return x.X.RltStm().AggPrd(x.I0.UntUnt()) }
func (x RltStmAggMin) Act()                 { x.RltStm() }
func (x RltStmAggMin) Ifc() interface{}     { return x.RltStm() }
func (x RltStmAggMin) RltStm() rlt.Stm      { return x.X.RltStm().AggMin(x.I0.UntUnt()) }
func (x RltStmAggMax) Act()                 { x.RltStm() }
func (x RltStmAggMax) Ifc() interface{}     { return x.RltStm() }
func (x RltStmAggMax) RltStm() rlt.Stm      { return x.X.RltStm().AggMax(x.I0.UntUnt()) }
func (x RltStmAggMid) Act()                 { x.RltStm() }
func (x RltStmAggMid) Ifc() interface{}     { return x.RltStm() }
func (x RltStmAggMid) RltStm() rlt.Stm      { return x.X.RltStm().AggMid(x.I0.UntUnt()) }
func (x RltStmAggMdn) Act()                 { x.RltStm() }
func (x RltStmAggMdn) Ifc() interface{}     { return x.RltStm() }
func (x RltStmAggMdn) RltStm() rlt.Stm      { return x.X.RltStm().AggMdn(x.I0.UntUnt()) }
func (x RltStmAggSma) Act()                 { x.RltStm() }
func (x RltStmAggSma) Ifc() interface{}     { return x.RltStm() }
func (x RltStmAggSma) RltStm() rlt.Stm      { return x.X.RltStm().AggSma(x.I0.UntUnt()) }
func (x RltStmAggGma) Act()                 { x.RltStm() }
func (x RltStmAggGma) Ifc() interface{}     { return x.RltStm() }
func (x RltStmAggGma) RltStm() rlt.Stm      { return x.X.RltStm().AggGma(x.I0.UntUnt()) }
func (x RltStmAggWma) Act()                 { x.RltStm() }
func (x RltStmAggWma) Ifc() interface{}     { return x.RltStm() }
func (x RltStmAggWma) RltStm() rlt.Stm      { return x.X.RltStm().AggWma(x.I0.UntUnt()) }
func (x RltStmAggRsi) Act()                 { x.RltStm() }
func (x RltStmAggRsi) Ifc() interface{}     { return x.RltStm() }
func (x RltStmAggRsi) RltStm() rlt.Stm      { return x.X.RltStm().AggRsi(x.I0.UntUnt()) }
func (x RltStmAggWrsi) Act()                { x.RltStm() }
func (x RltStmAggWrsi) Ifc() interface{}    { return x.RltStm() }
func (x RltStmAggWrsi) RltStm() rlt.Stm     { return x.X.RltStm().AggWrsi(x.I0.UntUnt()) }
func (x RltStmAggAlma) Act()                { x.RltStm() }
func (x RltStmAggAlma) Ifc() interface{}    { return x.RltStm() }
func (x RltStmAggAlma) RltStm() rlt.Stm     { return x.X.RltStm().AggAlma(x.I0.UntUnt()) }
func (x RltStmAggVrnc) Act()                { x.RltStm() }
func (x RltStmAggVrnc) Ifc() interface{}    { return x.RltStm() }
func (x RltStmAggVrnc) RltStm() rlt.Stm     { return x.X.RltStm().AggVrnc(x.I0.UntUnt()) }
func (x RltStmAggStd) Act()                 { x.RltStm() }
func (x RltStmAggStd) Ifc() interface{}     { return x.RltStm() }
func (x RltStmAggStd) RltStm() rlt.Stm      { return x.X.RltStm().AggStd(x.I0.UntUnt()) }
func (x RltStmAggRngFul) Act()              { x.RltStm() }
func (x RltStmAggRngFul) Ifc() interface{}  { return x.RltStm() }
func (x RltStmAggRngFul) RltStm() rlt.Stm   { return x.X.RltStm().AggRngFul(x.I0.UntUnt()) }
func (x RltStmAggRngLst) Act()              { x.RltStm() }
func (x RltStmAggRngLst) Ifc() interface{}  { return x.RltStm() }
func (x RltStmAggRngLst) RltStm() rlt.Stm   { return x.X.RltStm().AggRngLst(x.I0.UntUnt()) }
func (x RltStmAggProLst) Act()              { x.RltStm() }
func (x RltStmAggProLst) Ifc() interface{}  { return x.RltStm() }
func (x RltStmAggProLst) RltStm() rlt.Stm   { return x.X.RltStm().AggProLst(x.I0.UntUnt()) }
func (x RltStmAggProSma) Act()              { x.RltStm() }
func (x RltStmAggProSma) Ifc() interface{}  { return x.RltStm() }
func (x RltStmAggProSma) RltStm() rlt.Stm   { return x.X.RltStm().AggProSma(x.I0.UntUnt()) }
func (x RltStmAggProAlma) Act()             { x.RltStm() }
func (x RltStmAggProAlma) Ifc() interface{} { return x.RltStm() }
func (x RltStmAggProAlma) RltStm() rlt.Stm  { return x.X.RltStm().AggProAlma(x.I0.UntUnt()) }
func (x RltStmAggEma) Act()                 { x.RltStm() }
func (x RltStmAggEma) Ifc() interface{}     { return x.RltStm() }
func (x RltStmAggEma) RltStm() rlt.Stm      { return x.X.RltStm().AggEma(x.I0.UntUnt()) }
func (x RltStmInrAdd) Act()                 { x.RltStm() }
func (x RltStmInrAdd) Ifc() interface{}     { return x.RltStm() }
func (x RltStmInrAdd) RltStm() rlt.Stm      { return x.X.RltStm().InrAdd(x.I0.UntUnt()) }
func (x RltStmInrSub) Act()                 { x.RltStm() }
func (x RltStmInrSub) Ifc() interface{}     { return x.RltStm() }
func (x RltStmInrSub) RltStm() rlt.Stm      { return x.X.RltStm().InrSub(x.I0.UntUnt()) }
func (x RltStmInrMul) Act()                 { x.RltStm() }
func (x RltStmInrMul) Ifc() interface{}     { return x.RltStm() }
func (x RltStmInrMul) RltStm() rlt.Stm      { return x.X.RltStm().InrMul(x.I0.UntUnt()) }
func (x RltStmInrDiv) Act()                 { x.RltStm() }
func (x RltStmInrDiv) Ifc() interface{}     { return x.RltStm() }
func (x RltStmInrDiv) RltStm() rlt.Stm      { return x.X.RltStm().InrDiv(x.I0.UntUnt()) }
func (x RltStmInrRem) Act()                 { x.RltStm() }
func (x RltStmInrRem) Ifc() interface{}     { return x.RltStm() }
func (x RltStmInrRem) RltStm() rlt.Stm      { return x.X.RltStm().InrRem(x.I0.UntUnt()) }
func (x RltStmInrPow) Act()                 { x.RltStm() }
func (x RltStmInrPow) Ifc() interface{}     { return x.RltStm() }
func (x RltStmInrPow) RltStm() rlt.Stm      { return x.X.RltStm().InrPow(x.I0.UntUnt()) }
func (x RltStmInrMin) Act()                 { x.RltStm() }
func (x RltStmInrMin) Ifc() interface{}     { return x.RltStm() }
func (x RltStmInrMin) RltStm() rlt.Stm      { return x.X.RltStm().InrMin(x.I0.UntUnt()) }
func (x RltStmInrMax) Act()                 { x.RltStm() }
func (x RltStmInrMax) Ifc() interface{}     { return x.RltStm() }
func (x RltStmInrMax) RltStm() rlt.Stm      { return x.X.RltStm().InrMax(x.I0.UntUnt()) }
func (x RltStmInrSlp) Act()                 { x.RltStm() }
func (x RltStmInrSlp) Ifc() interface{}     { return x.RltStm() }
func (x RltStmInrSlp) RltStm() rlt.Stm      { return x.X.RltStm().InrSlp(x.I0.UntUnt()) }
func (x RltStmOtrAdd) Act()                 { x.RltStm() }
func (x RltStmOtrAdd) Ifc() interface{}     { return x.RltStm() }
func (x RltStmOtrAdd) RltStm() rlt.Stm      { return x.X.RltStm().OtrAdd(x.I0.UntUnt(), x.I1.RltStm()) }
func (x RltStmOtrSub) Act()                 { x.RltStm() }
func (x RltStmOtrSub) Ifc() interface{}     { return x.RltStm() }
func (x RltStmOtrSub) RltStm() rlt.Stm      { return x.X.RltStm().OtrSub(x.I0.UntUnt(), x.I1.RltStm()) }
func (x RltStmOtrMul) Act()                 { x.RltStm() }
func (x RltStmOtrMul) Ifc() interface{}     { return x.RltStm() }
func (x RltStmOtrMul) RltStm() rlt.Stm      { return x.X.RltStm().OtrMul(x.I0.UntUnt(), x.I1.RltStm()) }
func (x RltStmOtrDiv) Act()                 { x.RltStm() }
func (x RltStmOtrDiv) Ifc() interface{}     { return x.RltStm() }
func (x RltStmOtrDiv) RltStm() rlt.Stm      { return x.X.RltStm().OtrDiv(x.I0.UntUnt(), x.I1.RltStm()) }
func (x RltStmOtrRem) Act()                 { x.RltStm() }
func (x RltStmOtrRem) Ifc() interface{}     { return x.RltStm() }
func (x RltStmOtrRem) RltStm() rlt.Stm      { return x.X.RltStm().OtrRem(x.I0.UntUnt(), x.I1.RltStm()) }
func (x RltStmOtrPow) Act()                 { x.RltStm() }
func (x RltStmOtrPow) Ifc() interface{}     { return x.RltStm() }
func (x RltStmOtrPow) RltStm() rlt.Stm      { return x.X.RltStm().OtrPow(x.I0.UntUnt(), x.I1.RltStm()) }
func (x RltStmOtrMin) Act()                 { x.RltStm() }
func (x RltStmOtrMin) Ifc() interface{}     { return x.RltStm() }
func (x RltStmOtrMin) RltStm() rlt.Stm      { return x.X.RltStm().OtrMin(x.I0.UntUnt(), x.I1.RltStm()) }
func (x RltStmOtrMax) Act()                 { x.RltStm() }
func (x RltStmOtrMax) Ifc() interface{}     { return x.RltStm() }
func (x RltStmOtrMax) RltStm() rlt.Stm      { return x.X.RltStm().OtrMax(x.I0.UntUnt(), x.I1.RltStm()) }
func (x RltStmSclEql) Act()                 { x.RltCnd() }
func (x RltStmSclEql) Ifc() interface{}     { return x.RltCnd() }
func (x RltStmSclEql) RltCnd() rlt.Cnd      { return x.X.RltStm().SclEql(x.I0.FltFlt()) }
func (x RltStmSclNeq) Act()                 { x.RltCnd() }
func (x RltStmSclNeq) Ifc() interface{}     { return x.RltCnd() }
func (x RltStmSclNeq) RltCnd() rlt.Cnd      { return x.X.RltStm().SclNeq(x.I0.FltFlt()) }
func (x RltStmSclLss) Act()                 { x.RltCnd() }
func (x RltStmSclLss) Ifc() interface{}     { return x.RltCnd() }
func (x RltStmSclLss) RltCnd() rlt.Cnd      { return x.X.RltStm().SclLss(x.I0.FltFlt()) }
func (x RltStmSclGtr) Act()                 { x.RltCnd() }
func (x RltStmSclGtr) Ifc() interface{}     { return x.RltCnd() }
func (x RltStmSclGtr) RltCnd() rlt.Cnd      { return x.X.RltStm().SclGtr(x.I0.FltFlt()) }
func (x RltStmSclLeq) Act()                 { x.RltCnd() }
func (x RltStmSclLeq) Ifc() interface{}     { return x.RltCnd() }
func (x RltStmSclLeq) RltCnd() rlt.Cnd      { return x.X.RltStm().SclLeq(x.I0.FltFlt()) }
func (x RltStmSclGeq) Act()                 { x.RltCnd() }
func (x RltStmSclGeq) Ifc() interface{}     { return x.RltCnd() }
func (x RltStmSclGeq) RltCnd() rlt.Cnd      { return x.X.RltStm().SclGeq(x.I0.FltFlt()) }
func (x RltStmInrEql) Act()                 { x.RltCnd() }
func (x RltStmInrEql) Ifc() interface{}     { return x.RltCnd() }
func (x RltStmInrEql) RltCnd() rlt.Cnd      { return x.X.RltStm().InrEql(x.I0.UntUnt()) }
func (x RltStmInrNeq) Act()                 { x.RltCnd() }
func (x RltStmInrNeq) Ifc() interface{}     { return x.RltCnd() }
func (x RltStmInrNeq) RltCnd() rlt.Cnd      { return x.X.RltStm().InrNeq(x.I0.UntUnt()) }
func (x RltStmInrLss) Act()                 { x.RltCnd() }
func (x RltStmInrLss) Ifc() interface{}     { return x.RltCnd() }
func (x RltStmInrLss) RltCnd() rlt.Cnd      { return x.X.RltStm().InrLss(x.I0.UntUnt()) }
func (x RltStmInrGtr) Act()                 { x.RltCnd() }
func (x RltStmInrGtr) Ifc() interface{}     { return x.RltCnd() }
func (x RltStmInrGtr) RltCnd() rlt.Cnd      { return x.X.RltStm().InrGtr(x.I0.UntUnt()) }
func (x RltStmInrLeq) Act()                 { x.RltCnd() }
func (x RltStmInrLeq) Ifc() interface{}     { return x.RltCnd() }
func (x RltStmInrLeq) RltCnd() rlt.Cnd      { return x.X.RltStm().InrLeq(x.I0.UntUnt()) }
func (x RltStmInrGeq) Act()                 { x.RltCnd() }
func (x RltStmInrGeq) Ifc() interface{}     { return x.RltCnd() }
func (x RltStmInrGeq) RltCnd() rlt.Cnd      { return x.X.RltStm().InrGeq(x.I0.UntUnt()) }
func (x RltStmOtrEql) Act()                 { x.RltCnd() }
func (x RltStmOtrEql) Ifc() interface{}     { return x.RltCnd() }
func (x RltStmOtrEql) RltCnd() rlt.Cnd      { return x.X.RltStm().OtrEql(x.I0.UntUnt(), x.I1.RltStm()) }
func (x RltStmOtrNeq) Act()                 { x.RltCnd() }
func (x RltStmOtrNeq) Ifc() interface{}     { return x.RltCnd() }
func (x RltStmOtrNeq) RltCnd() rlt.Cnd      { return x.X.RltStm().OtrNeq(x.I0.UntUnt(), x.I1.RltStm()) }
func (x RltStmOtrLss) Act()                 { x.RltCnd() }
func (x RltStmOtrLss) Ifc() interface{}     { return x.RltCnd() }
func (x RltStmOtrLss) RltCnd() rlt.Cnd      { return x.X.RltStm().OtrLss(x.I0.UntUnt(), x.I1.RltStm()) }
func (x RltStmOtrGtr) Act()                 { x.RltCnd() }
func (x RltStmOtrGtr) Ifc() interface{}     { return x.RltCnd() }
func (x RltStmOtrGtr) RltCnd() rlt.Cnd      { return x.X.RltStm().OtrGtr(x.I0.UntUnt(), x.I1.RltStm()) }
func (x RltStmOtrLeq) Act()                 { x.RltCnd() }
func (x RltStmOtrLeq) Ifc() interface{}     { return x.RltCnd() }
func (x RltStmOtrLeq) RltCnd() rlt.Cnd      { return x.X.RltStm().OtrLeq(x.I0.UntUnt(), x.I1.RltStm()) }
func (x RltStmOtrGeq) Act()                 { x.RltCnd() }
func (x RltStmOtrGeq) Ifc() interface{}     { return x.RltCnd() }
func (x RltStmOtrGeq) RltCnd() rlt.Cnd      { return x.X.RltStm().OtrGeq(x.I0.UntUnt(), x.I1.RltStm()) }
func (x RltCndAnd) Act()                    { x.RltCnd() }
func (x RltCndAnd) Ifc() interface{}        { return x.RltCnd() }
func (x RltCndAnd) RltCnd() rlt.Cnd         { return x.X.RltCnd().And(x.I0.RltCnd()) }
func (x RltCndSeq) Act()                    { x.RltCnd() }
func (x RltCndSeq) Ifc() interface{}        { return x.RltCnd() }
func (x RltCndSeq) RltCnd() rlt.Cnd         { return x.X.RltCnd().Seq(x.I0.TmeTme(), x.I1.RltCnd()) }
func (x RltCndStgy) Act()                   { x.RltStgy() }
func (x RltCndStgy) Ifc() interface{}       { return x.RltStgy() }
func (x RltCndStgy) RltStgy() rlt.Stgy {
	var i7 []rlt.Cnd
	for _, cur := range x.I7 {
		i7 = append(i7, cur.RltCnd())
	}
	return x.X.RltCnd().Stgy(x.I0.BolBol(), x.I1.FltFlt(), x.I2.FltFlt(), x.I3.TmeTme(), x.I4.FltFlt(), x.I5.RltInstr(), x.I6.RltStms(), i7...)
}
func (x PltPltSho) Act()                { x.PltPlt() }
func (x PltPltSho) Ifc() interface{}    { return x.PltPlt() }
func (x PltPltSho) PltPlt() plt.Plt     { return x.X.PltPlt().Sho() }
func (x PltPltSiz) Act()                { x.PltPlt() }
func (x PltPltSiz) Ifc() interface{}    { return x.PltPlt() }
func (x PltPltSiz) PltPlt() plt.Plt     { return x.X.PltPlt().Siz(x.I0.UntUnt(), x.I1.UntUnt()) }
func (x PltPltScl) Act()                { x.PltPlt() }
func (x PltPltScl) Ifc() interface{}    { return x.PltPlt() }
func (x PltPltScl) PltPlt() plt.Plt     { return x.X.PltPlt().Scl(x.I0.FltFlt()) }
func (x PltPltHrzScl) Act()             { x.PltPlt() }
func (x PltPltHrzScl) Ifc() interface{} { return x.PltPlt() }
func (x PltPltHrzScl) PltPlt() plt.Plt  { return x.X.PltPlt().HrzScl(x.I0.FltFlt()) }
func (x PltPltVrtScl) Act()             { x.PltPlt() }
func (x PltPltVrtScl) Ifc() interface{} { return x.PltPlt() }
func (x PltPltVrtScl) PltPlt() plt.Plt  { return x.X.PltPlt().VrtScl(x.I0.FltFlt()) }
func (x IfcIfc) Act()                   { x.Ifc() }
func (x IfcIfc) Ifc() interface{}       { return x.X.Ifc() }
func (x *PllWaitSeg) Act() {
	defer x.Wg.Done()
	var actr Actr
	actr.Reset(x.Txt)
	scp := NewScp(x.XprScp, x.ActScpPrnt)
	act := actr.Act(scp, x.Xpr)
	act.Act()
}
func (x PllWait) Act() {
	wg := &sync.WaitGroup{}
	segs := make([]sys.Act, len(x.Xprs))
	for n, xpr := range x.Xprs {
		segs[n] = &PllWaitSeg{
			Txt:        x.Txt,
			ActScpPrnt: x.ActScpPrnt,
			XprScp:     x.XprScp,
			Xpr:        xpr,
			Wg:         wg,
		}
	}
	wg.Add(len(x.Xprs))
	sys.Run().Pll(segs...) // run segs in pll
	wg.Wait()
}
func (x PllWait) Ifc() interface{} { return nil }
func (x *Actr) Cmpl(txt string) []Act {
	xprScp, xprs := x.Prs(txt)
	return x.Acts(NewScp(xprScp), xprs...)
}
func (x *Actr) Cmplf(format string, args ...interface{}) []Act {
	xprScp, xprs := x.Prsf(format, args...)
	return x.Acts(NewScp(xprScp), xprs...)
}
func (x *Actr) Run(txt string) {
	for _, a := range x.Cmpl(txt) {
		a.Act()
	}
}
func (x *Actr) Runf(format string, args ...interface{}) {
	for _, a := range x.Cmplf(format, args...) {
		a.Act()
	}
}
func (x *Actr) RunIfc(txt string) (r []interface{}) {
	txt = x.Reduce(txt)
	var actr Actr // new instance for each cmpl assures lock-free access
	for _, a := range actr.Cmpl(txt) {
		r = append(r, a.(sys.Ifc).Ifc())
	}
	return r
}
func (x *Actr) RunIfcf(format string, args ...interface{}) (r []interface{}) {
	var actr Actr // new instance for each cmpl assures lock-free access
	return actr.RunIfc(fmt.Sprintf(format, args...))
}
func (x *Actr) RunRlt(txt string) []interface{} {
	var actr Actr // new instance for each cmpl assures lock-free access. and no over-write of on-going cmpl for rlt
	return actr.RunIfc(strings.Replace(txt, "hst.", "rlt.", -1))
}
func (x *Actr) RunHst(txt string) []interface{} {
	var actr Actr // new instance for each cmpl assures lock-free access. and no over-write of on-going cmpl for rlt
	return actr.RunIfc(strings.Replace(txt, "rlt.", "hst.", -1))
}
func (x *Actr) Reduce(txt string) string {
	type Node struct {
		Xpr string
		Idn string
	}
	// FIND ALL COMMON SUB-EXPRESSIONS
	// ASSIGN EACH SUB-EXPRESSION TO A SINGLE VARIABLE
	// TO AVOID DUPLICATE HST OR RLT CALCULATIONS
	var trmr trm.Trmr
	var sb strings.Builder
	var idnCnt uint32
	cur := &Node{Idn: "hst"}
	stck := append([]*Node{}, cur)
	for len(stck) != 0 {
		cur := stck[len(stck)-1]
		stck = stck[:len(stck)-1]
		trmr.Reset(txt)
		curNodes := make(map[string]*Node)
		xprBnds := trmr.Prefixs(cur.Idn)
		xprs := make([]string, len(xprBnds))
		for n, xprBnd := range xprBnds {
			xprs[n] = txt[xprBnd.Idx:xprBnd.Lim]
		}
		for _, xpr := range xprs {
			if _, ok := curNodes[xpr]; !ok {
				idnCnt++
				node := &Node{Xpr: xpr, Idn: fmt.Sprintf("v%v", idnCnt)}
				sb.WriteString(fmt.Sprintf("%v.asn(%v)\n", node.Xpr, node.Idn))
				reNode := regexp.MustCompile(regexp.QuoteMeta(node.Xpr))
				txt = reNode.ReplaceAllString(txt, node.Idn)
				stck = append(stck, node)
				curNodes[xpr] = node
			}
		}
	}
	sb.WriteString(txt)
	return sb.String()
}
func (x *Actr) Acts(scp *Scp, vs ...xpr.Xpr) (r []Act) {
	for _, v := range vs {
		r = append(r, x.Act(scp, v))
	}
	return r
}
func (x *Actr) Act(scp *Scp, v xpr.Xpr) Act {
	switch X := v.(type) {
	case *xpr.PllWait:
		return PllWait{Txt: x.Txt, ActScpPrnt: scp, XprScp: X.Scp, Xprs: X.Xprs}
	case xpr.SysMuXpr:
		return x.SysMuAct(scp, X)
	case xpr.LogLogrXpr:
		return x.LogLogrAct(scp, X)
	case xpr.StrStrXpr:
		return x.StrStrAct(scp, X)
	case xpr.BolBolXpr:
		return x.BolBolAct(scp, X)
	case xpr.FltFltXpr:
		return x.FltFltAct(scp, X)
	case xpr.FltRngXpr:
		return x.FltRngAct(scp, X)
	case xpr.UntUntXpr:
		return x.UntUntAct(scp, X)
	case xpr.IntIntXpr:
		return x.IntIntAct(scp, X)
	case xpr.TmeTmeXpr:
		return x.TmeTmeAct(scp, X)
	case xpr.TmeRngXpr:
		return x.TmeRngAct(scp, X)
	case xpr.TmeRngsXpr:
		return x.TmeRngsAct(scp, X)
	case xpr.BndBndXpr:
		return x.BndBndAct(scp, X)
	case xpr.StrsStrsXpr:
		return x.StrsStrsAct(scp, X)
	case xpr.BolsBolsXpr:
		return x.BolsBolsAct(scp, X)
	case xpr.FltsFltsXpr:
		return x.FltsFltsAct(scp, X)
	case xpr.UntsUntsXpr:
		return x.UntsUntsAct(scp, X)
	case xpr.IntsIntsXpr:
		return x.IntsIntsAct(scp, X)
	case xpr.TmesTmesXpr:
		return x.TmesTmesAct(scp, X)
	case xpr.BndsBndsXpr:
		return x.BndsBndsAct(scp, X)
	case xpr.AnaTrdXpr:
		return x.AnaTrdAct(scp, X)
	case xpr.AnaTrdsXpr:
		return x.AnaTrdsAct(scp, X)
	case xpr.AnaPrfmXpr:
		return x.AnaPrfmAct(scp, X)
	case xpr.AnaPrfmsXpr:
		return x.AnaPrfmsAct(scp, X)
	case xpr.AnaPrfmDltXpr:
		return x.AnaPrfmDltAct(scp, X)
	case xpr.AnaPortXpr:
		return x.AnaPortAct(scp, X)
	case xpr.HstPrvXpr:
		return x.HstPrvAct(scp, X)
	case xpr.HstInstrXpr:
		return x.HstInstrAct(scp, X)
	case xpr.HstInrvlXpr:
		return x.HstInrvlAct(scp, X)
	case xpr.HstSideXpr:
		return x.HstSideAct(scp, X)
	case xpr.HstStmXpr:
		return x.HstStmAct(scp, X)
	case xpr.HstCndXpr:
		return x.HstCndAct(scp, X)
	case xpr.HstStgyXpr:
		return x.HstStgyAct(scp, X)
	case xpr.HstPrvsXpr:
		return x.HstPrvsAct(scp, X)
	case xpr.HstInstrsXpr:
		return x.HstInstrsAct(scp, X)
	case xpr.HstInrvlsXpr:
		return x.HstInrvlsAct(scp, X)
	case xpr.HstSidesXpr:
		return x.HstSidesAct(scp, X)
	case xpr.HstStmsXpr:
		return x.HstStmsAct(scp, X)
	case xpr.HstCndsXpr:
		return x.HstCndsAct(scp, X)
	case xpr.HstStgysXpr:
		return x.HstStgysAct(scp, X)
	case xpr.RltPrvXpr:
		return x.RltPrvAct(scp, X)
	case xpr.RltInstrXpr:
		return x.RltInstrAct(scp, X)
	case xpr.RltInrvlXpr:
		return x.RltInrvlAct(scp, X)
	case xpr.RltSideXpr:
		return x.RltSideAct(scp, X)
	case xpr.RltStmXpr:
		return x.RltStmAct(scp, X)
	case xpr.RltCndXpr:
		return x.RltCndAct(scp, X)
	case xpr.RltStgyXpr:
		return x.RltStgyAct(scp, X)
	case xpr.RltPrvsXpr:
		return x.RltPrvsAct(scp, X)
	case xpr.RltInstrsXpr:
		return x.RltInstrsAct(scp, X)
	case xpr.RltInrvlsXpr:
		return x.RltInrvlsAct(scp, X)
	case xpr.RltSidesXpr:
		return x.RltSidesAct(scp, X)
	case xpr.RltStmsXpr:
		return x.RltStmsAct(scp, X)
	case xpr.RltCndsXpr:
		return x.RltCndsAct(scp, X)
	case xpr.RltStgysXpr:
		return x.RltStgysAct(scp, X)
	case xpr.FntFntXpr:
		return x.FntFntAct(scp, X)
	case xpr.ClrClrXpr:
		return x.ClrClrAct(scp, X)
	case xpr.PenPenXpr:
		return x.PenPenAct(scp, X)
	case xpr.PenPensXpr:
		return x.PenPensAct(scp, X)
	case xpr.PltPltXpr:
		return x.PltPltAct(scp, X)
	case xpr.PltPltsXpr:
		return x.PltPltsAct(scp, X)
	case xpr.PltTmeAxisXXpr:
		return x.PltTmeAxisXAct(scp, X)
	case xpr.PltFltAxisYXpr:
		return x.PltFltAxisYAct(scp, X)
	case xpr.PltStmXpr:
		return x.PltStmAct(scp, X)
	case xpr.PltFltsSctrXpr:
		return x.PltFltsSctrAct(scp, X)
	case xpr.PltFltsSctrDistXpr:
		return x.PltFltsSctrDistAct(scp, X)
	case xpr.PltHrzXpr:
		return x.PltHrzAct(scp, X)
	case xpr.PltVrtXpr:
		return x.PltVrtAct(scp, X)
	case xpr.PltDpthXpr:
		return x.PltDpthAct(scp, X)
	}
	panic(x.Erf("Act: no action found %v", reflect.ValueOf(v).Elem().Type().Name()))
}
func (x *Actr) LogLogrAct(scp *Scp, v xpr.LogLogrXpr) LogLogrAct {
	switch v.(type) {
	}
	panic(x.Erf("LogLogrAct: no action found %v", reflect.ValueOf(v).Elem().Type().Name()))
}
func (x *Actr) StrStrAct(scp *Scp, v xpr.StrStrXpr) StrStrAct {
	switch X := v.(type) {
	case *xpr.StrStrLit:
		return StrStrLit{Trm: X.Trm, Txt: x.Txt}
	case *xpr.StrStrAsn:
		asnScp := scp
		if X.Depth != 0 {
			for d := X.Depth; asnScp != nil && d != 0; d-- { // recurse up scp by depth
				asnScp = asnScp.Prnt
			}
		}
		return StrStrAsn{StrScp: asnScp.StrStr(x.Txt[X.Idn.Idx:X.Idn.Lim]), X: x.StrStrAct(scp, X.X)}
	case *xpr.StrStrAcs:
		return StrStrAcs{StrScp: scp.StrStr(x.Txt[X.Trm.Idx:X.Trm.Lim])}
	case *xpr.AnaPrfmDltPthBGet:
		return AnaPrfmDltPthBGet{X: x.AnaPrfmDltAct(scp, X.X)}
	case *xpr.PltStmTitleSetGet:
		if X.I0 == nil {
			return PltStmTitleSetGet{X: x.PltStmAct(scp, X.X)}
		} else {
			return PltStmTitleSetGet{X: x.PltStmAct(scp, X.X), I0: x.StrStrAct(scp, X.I0)}
		}
	case *xpr.PltFltsSctrTitleSetGet:
		if X.I0 == nil {
			return PltFltsSctrTitleSetGet{X: x.PltFltsSctrAct(scp, X.X)}
		} else {
			return PltFltsSctrTitleSetGet{X: x.PltFltsSctrAct(scp, X.X), I0: x.StrStrAct(scp, X.I0)}
		}
	case *xpr.StrZero:
		return StrZero{}
	case *xpr.StrEmpty:
		return StrEmpty{}
	case *xpr.LogIfo:
		var i0 []Act
		for _, cur := range X.I0 {
			i0 = append(i0, x.Act(scp, cur))
		}
		return StrIfo{I0: i0}
	case *xpr.LogIfof:
		var i1 []Act
		for _, cur := range X.I1 {
			i1 = append(i1, x.Act(scp, cur))
		}
		return StrIfof{I0: x.StrStrAct(scp, X.I0), I1: i1}
	case *xpr.StrFmt:
		var i1 []Act
		for _, cur := range X.I1 {
			i1 = append(i1, x.Act(scp, cur))
		}
		return StrFmt{I0: x.StrStrAct(scp, X.I0), I1: i1}
	case *xpr.StrStrLower:
		return StrStrLower{X: x.StrStrAct(scp, X.X)}
	case *xpr.StrStrUpper:
		return StrStrUpper{X: x.StrStrAct(scp, X.X)}
	case *xpr.StrsStrsPop:
		return StrsStrsPop{X: x.StrsStrsAct(scp, X.X)}
	case *xpr.StrsStrsDque:
		return StrsStrsDque{X: x.StrsStrsAct(scp, X.X)}
	case *xpr.StrsStrsDel:
		return StrsStrsDel{X: x.StrsStrsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.StrsStrsAt:
		return StrsStrsAt{X: x.StrsStrsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.StrsStrsFst:
		return StrsStrsFst{X: x.StrsStrsAct(scp, X.X)}
	case *xpr.StrsStrsMdl:
		return StrsStrsMdl{X: x.StrsStrsAct(scp, X.X)}
	case *xpr.StrsStrsLst:
		return StrsStrsLst{X: x.StrsStrsAct(scp, X.X)}
	case *xpr.HstPrvName:
		return HstPrvName{X: x.HstPrvAct(scp, X.X)}
	case *xpr.HstInstrName:
		return HstInstrName{X: x.HstInstrAct(scp, X.X)}
	case *xpr.HstInrvlName:
		return HstInrvlName{X: x.HstInrvlAct(scp, X.X)}
	case *xpr.HstSideName:
		return HstSideName{X: x.HstSideAct(scp, X.X)}
	case *xpr.HstStmName:
		return HstStmName{X: x.HstStmAct(scp, X.X)}
	case *xpr.HstCndName:
		return HstCndName{X: x.HstCndAct(scp, X.X)}
	case *xpr.HstStgyName:
		return HstStgyName{X: x.HstStgyAct(scp, X.X)}
	}
	panic(x.Erf("StrStrAct: no action found %v", reflect.ValueOf(v).Elem().Type().Name()))
}
func (x *Actr) BolBolAct(scp *Scp, v xpr.BolBolXpr) BolBolAct {
	switch X := v.(type) {
	case *xpr.BolBolLit:
		return BolBolLit{Trm: X.Trm, Txt: x.Txt}
	case *xpr.BolBolAsn:
		asnScp := scp
		if X.Depth != 0 {
			for d := X.Depth; asnScp != nil && d != 0; d-- { // recurse up scp by depth
				asnScp = asnScp.Prnt
			}
		}
		return BolBolAsn{BolScp: asnScp.BolBol(x.Txt[X.Idn.Idx:X.Idn.Lim]), X: x.BolBolAct(scp, X.X)}
	case *xpr.BolBolAcs:
		return BolBolAcs{BolScp: scp.BolBol(x.Txt[X.Trm.Idx:X.Trm.Lim])}
	case *xpr.BolBolThen:
		return BolBolThen{X: x.BolBolAct(scp, X.X), Acts: x.Acts(NewScp(X.Scp, scp), X.Xprs...)}
	case *xpr.BolBolElse:
		return BolBolElse{X: x.BolBolAct(scp, X.X), Acts: x.Acts(NewScp(X.Scp, scp), X.Xprs...)}
	case *xpr.PltFltsSctrOutlierSetGet:
		if X.I0 == nil {
			return PltFltsSctrOutlierSetGet{X: x.PltFltsSctrAct(scp, X.X)}
		} else {
			return PltFltsSctrOutlierSetGet{X: x.PltFltsSctrAct(scp, X.X), I0: x.BolBolAct(scp, X.I0)}
		}
	case *xpr.BolZero:
		return BolZero{}
	case *xpr.BolFls:
		return BolFls{}
	case *xpr.BolTru:
		return BolTru{}
	case *xpr.StrStrEql:
		return StrStrEql{X: x.StrStrAct(scp, X.X), I0: x.StrStrAct(scp, X.I0)}
	case *xpr.StrStrNeq:
		return StrStrNeq{X: x.StrStrAct(scp, X.X), I0: x.StrStrAct(scp, X.I0)}
	case *xpr.StrStrLss:
		return StrStrLss{X: x.StrStrAct(scp, X.X), I0: x.StrStrAct(scp, X.I0)}
	case *xpr.StrStrGtr:
		return StrStrGtr{X: x.StrStrAct(scp, X.X), I0: x.StrStrAct(scp, X.I0)}
	case *xpr.StrStrLeq:
		return StrStrLeq{X: x.StrStrAct(scp, X.X), I0: x.StrStrAct(scp, X.I0)}
	case *xpr.StrStrGeq:
		return StrStrGeq{X: x.StrStrAct(scp, X.X), I0: x.StrStrAct(scp, X.I0)}
	case *xpr.BolBolNot:
		return BolBolNot{X: x.BolBolAct(scp, X.X)}
	case *xpr.BolBolEql:
		return BolBolEql{X: x.BolBolAct(scp, X.X), I0: x.BolBolAct(scp, X.I0)}
	case *xpr.BolBolNeq:
		return BolBolNeq{X: x.BolBolAct(scp, X.X), I0: x.BolBolAct(scp, X.I0)}
	case *xpr.FltFltEql:
		return FltFltEql{X: x.FltFltAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.FltFltNeq:
		return FltFltNeq{X: x.FltFltAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.FltFltIsNaN:
		return FltFltIsNaN{X: x.FltFltAct(scp, X.X)}
	case *xpr.FltFltIsInfPos:
		return FltFltIsInfPos{X: x.FltFltAct(scp, X.X)}
	case *xpr.FltFltIsInfNeg:
		return FltFltIsInfNeg{X: x.FltFltAct(scp, X.X)}
	case *xpr.FltFltIsValid:
		return FltFltIsValid{X: x.FltFltAct(scp, X.X)}
	case *xpr.FltFltLss:
		return FltFltLss{X: x.FltFltAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.FltFltGtr:
		return FltFltGtr{X: x.FltFltAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.FltFltLeq:
		return FltFltLeq{X: x.FltFltAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.FltFltGeq:
		return FltFltGeq{X: x.FltFltAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.UntUntEql:
		return UntUntEql{X: x.UntUntAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.UntUntNeq:
		return UntUntNeq{X: x.UntUntAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.UntUntLss:
		return UntUntLss{X: x.UntUntAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.UntUntGtr:
		return UntUntGtr{X: x.UntUntAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.UntUntLeq:
		return UntUntLeq{X: x.UntUntAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.UntUntGeq:
		return UntUntGeq{X: x.UntUntAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.IntIntEql:
		return IntIntEql{X: x.IntIntAct(scp, X.X), I0: x.IntIntAct(scp, X.I0)}
	case *xpr.IntIntNeq:
		return IntIntNeq{X: x.IntIntAct(scp, X.X), I0: x.IntIntAct(scp, X.I0)}
	case *xpr.IntIntLss:
		return IntIntLss{X: x.IntIntAct(scp, X.X), I0: x.IntIntAct(scp, X.I0)}
	case *xpr.IntIntGtr:
		return IntIntGtr{X: x.IntIntAct(scp, X.X), I0: x.IntIntAct(scp, X.I0)}
	case *xpr.IntIntLeq:
		return IntIntLeq{X: x.IntIntAct(scp, X.X), I0: x.IntIntAct(scp, X.I0)}
	case *xpr.IntIntGeq:
		return IntIntGeq{X: x.IntIntAct(scp, X.X), I0: x.IntIntAct(scp, X.I0)}
	case *xpr.TmeTmeIsSunday:
		return TmeTmeIsSunday{X: x.TmeTmeAct(scp, X.X)}
	case *xpr.TmeTmeIsMonday:
		return TmeTmeIsMonday{X: x.TmeTmeAct(scp, X.X)}
	case *xpr.TmeTmeIsTuesday:
		return TmeTmeIsTuesday{X: x.TmeTmeAct(scp, X.X)}
	case *xpr.TmeTmeIsWednesday:
		return TmeTmeIsWednesday{X: x.TmeTmeAct(scp, X.X)}
	case *xpr.TmeTmeIsThursday:
		return TmeTmeIsThursday{X: x.TmeTmeAct(scp, X.X)}
	case *xpr.TmeTmeIsFriday:
		return TmeTmeIsFriday{X: x.TmeTmeAct(scp, X.X)}
	case *xpr.TmeTmeIsSaturday:
		return TmeTmeIsSaturday{X: x.TmeTmeAct(scp, X.X)}
	case *xpr.TmeTmeEql:
		return TmeTmeEql{X: x.TmeTmeAct(scp, X.X), I0: x.TmeTmeAct(scp, X.I0)}
	case *xpr.TmeTmeNeq:
		return TmeTmeNeq{X: x.TmeTmeAct(scp, X.X), I0: x.TmeTmeAct(scp, X.I0)}
	case *xpr.TmeTmeLss:
		return TmeTmeLss{X: x.TmeTmeAct(scp, X.X), I0: x.TmeTmeAct(scp, X.I0)}
	case *xpr.TmeTmeGtr:
		return TmeTmeGtr{X: x.TmeTmeAct(scp, X.X), I0: x.TmeTmeAct(scp, X.I0)}
	case *xpr.TmeTmeLeq:
		return TmeTmeLeq{X: x.TmeTmeAct(scp, X.X), I0: x.TmeTmeAct(scp, X.I0)}
	case *xpr.TmeTmeGeq:
		return TmeTmeGeq{X: x.TmeTmeAct(scp, X.X), I0: x.TmeTmeAct(scp, X.I0)}
	case *xpr.BndBndIsValid:
		return BndBndIsValid{X: x.BndBndAct(scp, X.X)}
	case *xpr.FltRngIsValid:
		return FltRngIsValid{X: x.FltRngAct(scp, X.X)}
	case *xpr.TmeRngIsValid:
		return TmeRngIsValid{X: x.TmeRngAct(scp, X.X)}
	case *xpr.StrsStrsHas:
		return StrsStrsHas{X: x.StrsStrsAct(scp, X.X), I0: x.StrStrAct(scp, X.I0)}
	case *xpr.BolsBolsPop:
		return BolsBolsPop{X: x.BolsBolsAct(scp, X.X)}
	case *xpr.BolsBolsDque:
		return BolsBolsDque{X: x.BolsBolsAct(scp, X.X)}
	case *xpr.BolsBolsDel:
		return BolsBolsDel{X: x.BolsBolsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.BolsBolsAt:
		return BolsBolsAt{X: x.BolsBolsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.BolsBolsFst:
		return BolsBolsFst{X: x.BolsBolsAct(scp, X.X)}
	case *xpr.BolsBolsMdl:
		return BolsBolsMdl{X: x.BolsBolsAct(scp, X.X)}
	case *xpr.BolsBolsLst:
		return BolsBolsLst{X: x.BolsBolsAct(scp, X.X)}
	case *xpr.FltsFltsHas:
		return FltsFltsHas{X: x.FltsFltsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.UntsUntsHas:
		return UntsUntsHas{X: x.UntsUntsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.IntsIntsHas:
		return IntsIntsHas{X: x.IntsIntsAct(scp, X.X), I0: x.IntIntAct(scp, X.I0)}
	case *xpr.TmesTmesHas:
		return TmesTmesHas{X: x.TmesTmesAct(scp, X.X), I0: x.TmeTmeAct(scp, X.I0)}
	case *xpr.RltPrvMayTrd:
		return RltPrvMayTrd{X: x.RltPrvAct(scp, X.X)}
	}
	panic(x.Erf("BolBolAct: no action found %v", reflect.ValueOf(v).Elem().Type().Name()))
}
func (x *Actr) FltFltAct(scp *Scp, v xpr.FltFltXpr) FltFltAct {
	switch X := v.(type) {
	case *xpr.FltFltLit:
		return FltFltLit{Trm: X.Trm, Txt: x.Txt}
	case *xpr.FltFltAsn:
		asnScp := scp
		if X.Depth != 0 {
			for d := X.Depth; asnScp != nil && d != 0; d-- { // recurse up scp by depth
				asnScp = asnScp.Prnt
			}
		}
		return FltFltAsn{FltScp: asnScp.FltFlt(x.Txt[X.Idn.Idx:X.Idn.Lim]), X: x.FltFltAct(scp, X.X)}
	case *xpr.FltFltAcs:
		return FltFltAcs{FltScp: scp.FltFlt(x.Txt[X.Trm.Idx:X.Trm.Lim])}
	case *xpr.AnaPrfmPnlPctGet:
		return AnaPrfmPnlPctGet{X: x.AnaPrfmAct(scp, X.X)}
	case *xpr.AnaPrfmScsPctGet:
		return AnaPrfmScsPctGet{X: x.AnaPrfmAct(scp, X.X)}
	case *xpr.AnaPrfmPipPerDayGet:
		return AnaPrfmPipPerDayGet{X: x.AnaPrfmAct(scp, X.X)}
	case *xpr.AnaPrfmUsdPerDayGet:
		return AnaPrfmUsdPerDayGet{X: x.AnaPrfmAct(scp, X.X)}
	case *xpr.AnaPrfmScsPerDayGet:
		return AnaPrfmScsPerDayGet{X: x.AnaPrfmAct(scp, X.X)}
	case *xpr.AnaPrfmOpnPerDayGet:
		return AnaPrfmOpnPerDayGet{X: x.AnaPrfmAct(scp, X.X)}
	case *xpr.AnaPrfmPnlUsdGet:
		return AnaPrfmPnlUsdGet{X: x.AnaPrfmAct(scp, X.X)}
	case *xpr.AnaPrfmPipAvgGet:
		return AnaPrfmPipAvgGet{X: x.AnaPrfmAct(scp, X.X)}
	case *xpr.AnaPrfmPipMdnGet:
		return AnaPrfmPipMdnGet{X: x.AnaPrfmAct(scp, X.X)}
	case *xpr.AnaPrfmPipMinGet:
		return AnaPrfmPipMinGet{X: x.AnaPrfmAct(scp, X.X)}
	case *xpr.AnaPrfmPipMaxGet:
		return AnaPrfmPipMaxGet{X: x.AnaPrfmAct(scp, X.X)}
	case *xpr.AnaPrfmPipSumGet:
		return AnaPrfmPipSumGet{X: x.AnaPrfmAct(scp, X.X)}
	case *xpr.AnaPrfmLosLimMaxGet:
		return AnaPrfmLosLimMaxGet{X: x.AnaPrfmAct(scp, X.X)}
	case *xpr.AnaPrfmTrdPctGet:
		return AnaPrfmTrdPctGet{X: x.AnaPrfmAct(scp, X.X)}
	case *xpr.AnaPrfmCstTotUsdGet:
		return AnaPrfmCstTotUsdGet{X: x.AnaPrfmAct(scp, X.X)}
	case *xpr.AnaPrfmCstSpdUsdGet:
		return AnaPrfmCstSpdUsdGet{X: x.AnaPrfmAct(scp, X.X)}
	case *xpr.AnaPrfmCstComUsdGet:
		return AnaPrfmCstComUsdGet{X: x.AnaPrfmAct(scp, X.X)}
	case *xpr.AnaPrfmDltPnlPctAGet:
		return AnaPrfmDltPnlPctAGet{X: x.AnaPrfmDltAct(scp, X.X)}
	case *xpr.AnaPrfmDltPnlPctBGet:
		return AnaPrfmDltPnlPctBGet{X: x.AnaPrfmDltAct(scp, X.X)}
	case *xpr.AnaPrfmDltPnlPctDltGet:
		return AnaPrfmDltPnlPctDltGet{X: x.AnaPrfmDltAct(scp, X.X)}
	case *xpr.AnaPrfmDltScsPctAGet:
		return AnaPrfmDltScsPctAGet{X: x.AnaPrfmDltAct(scp, X.X)}
	case *xpr.AnaPrfmDltScsPctBGet:
		return AnaPrfmDltScsPctBGet{X: x.AnaPrfmDltAct(scp, X.X)}
	case *xpr.AnaPrfmDltScsPctDltGet:
		return AnaPrfmDltScsPctDltGet{X: x.AnaPrfmDltAct(scp, X.X)}
	case *xpr.AnaPrfmDltPipPerDayAGet:
		return AnaPrfmDltPipPerDayAGet{X: x.AnaPrfmDltAct(scp, X.X)}
	case *xpr.AnaPrfmDltPipPerDayBGet:
		return AnaPrfmDltPipPerDayBGet{X: x.AnaPrfmDltAct(scp, X.X)}
	case *xpr.AnaPrfmDltPipPerDayDltGet:
		return AnaPrfmDltPipPerDayDltGet{X: x.AnaPrfmDltAct(scp, X.X)}
	case *xpr.AnaPrfmDltUsdPerDayAGet:
		return AnaPrfmDltUsdPerDayAGet{X: x.AnaPrfmDltAct(scp, X.X)}
	case *xpr.AnaPrfmDltUsdPerDayBGet:
		return AnaPrfmDltUsdPerDayBGet{X: x.AnaPrfmDltAct(scp, X.X)}
	case *xpr.AnaPrfmDltUsdPerDayDltGet:
		return AnaPrfmDltUsdPerDayDltGet{X: x.AnaPrfmDltAct(scp, X.X)}
	case *xpr.AnaPrfmDltScsPerDayAGet:
		return AnaPrfmDltScsPerDayAGet{X: x.AnaPrfmDltAct(scp, X.X)}
	case *xpr.AnaPrfmDltScsPerDayBGet:
		return AnaPrfmDltScsPerDayBGet{X: x.AnaPrfmDltAct(scp, X.X)}
	case *xpr.AnaPrfmDltScsPerDayDltGet:
		return AnaPrfmDltScsPerDayDltGet{X: x.AnaPrfmDltAct(scp, X.X)}
	case *xpr.AnaPrfmDltOpnPerDayAGet:
		return AnaPrfmDltOpnPerDayAGet{X: x.AnaPrfmDltAct(scp, X.X)}
	case *xpr.AnaPrfmDltOpnPerDayBGet:
		return AnaPrfmDltOpnPerDayBGet{X: x.AnaPrfmDltAct(scp, X.X)}
	case *xpr.AnaPrfmDltOpnPerDayDltGet:
		return AnaPrfmDltOpnPerDayDltGet{X: x.AnaPrfmDltAct(scp, X.X)}
	case *xpr.AnaPrfmDltPnlUsdAGet:
		return AnaPrfmDltPnlUsdAGet{X: x.AnaPrfmDltAct(scp, X.X)}
	case *xpr.AnaPrfmDltPnlUsdBGet:
		return AnaPrfmDltPnlUsdBGet{X: x.AnaPrfmDltAct(scp, X.X)}
	case *xpr.AnaPrfmDltPnlUsdDltGet:
		return AnaPrfmDltPnlUsdDltGet{X: x.AnaPrfmDltAct(scp, X.X)}
	case *xpr.AnaPrfmDltPipAvgAGet:
		return AnaPrfmDltPipAvgAGet{X: x.AnaPrfmDltAct(scp, X.X)}
	case *xpr.AnaPrfmDltPipAvgBGet:
		return AnaPrfmDltPipAvgBGet{X: x.AnaPrfmDltAct(scp, X.X)}
	case *xpr.AnaPrfmDltPipAvgDltGet:
		return AnaPrfmDltPipAvgDltGet{X: x.AnaPrfmDltAct(scp, X.X)}
	case *xpr.AnaPrfmDltPipMdnAGet:
		return AnaPrfmDltPipMdnAGet{X: x.AnaPrfmDltAct(scp, X.X)}
	case *xpr.AnaPrfmDltPipMdnBGet:
		return AnaPrfmDltPipMdnBGet{X: x.AnaPrfmDltAct(scp, X.X)}
	case *xpr.AnaPrfmDltPipMdnDltGet:
		return AnaPrfmDltPipMdnDltGet{X: x.AnaPrfmDltAct(scp, X.X)}
	case *xpr.AnaPrfmDltPipMinAGet:
		return AnaPrfmDltPipMinAGet{X: x.AnaPrfmDltAct(scp, X.X)}
	case *xpr.AnaPrfmDltPipMinBGet:
		return AnaPrfmDltPipMinBGet{X: x.AnaPrfmDltAct(scp, X.X)}
	case *xpr.AnaPrfmDltPipMinDltGet:
		return AnaPrfmDltPipMinDltGet{X: x.AnaPrfmDltAct(scp, X.X)}
	case *xpr.AnaPrfmDltPipMaxAGet:
		return AnaPrfmDltPipMaxAGet{X: x.AnaPrfmDltAct(scp, X.X)}
	case *xpr.AnaPrfmDltPipMaxBGet:
		return AnaPrfmDltPipMaxBGet{X: x.AnaPrfmDltAct(scp, X.X)}
	case *xpr.AnaPrfmDltPipMaxDltGet:
		return AnaPrfmDltPipMaxDltGet{X: x.AnaPrfmDltAct(scp, X.X)}
	case *xpr.AnaPrfmDltPipSumAGet:
		return AnaPrfmDltPipSumAGet{X: x.AnaPrfmDltAct(scp, X.X)}
	case *xpr.AnaPrfmDltPipSumBGet:
		return AnaPrfmDltPipSumBGet{X: x.AnaPrfmDltAct(scp, X.X)}
	case *xpr.AnaPrfmDltPipSumDltGet:
		return AnaPrfmDltPipSumDltGet{X: x.AnaPrfmDltAct(scp, X.X)}
	case *xpr.AnaPrfmDltDurAvgDltGet:
		return AnaPrfmDltDurAvgDltGet{X: x.AnaPrfmDltAct(scp, X.X)}
	case *xpr.AnaPrfmDltDurMdnDltGet:
		return AnaPrfmDltDurMdnDltGet{X: x.AnaPrfmDltAct(scp, X.X)}
	case *xpr.AnaPrfmDltDurMinDltGet:
		return AnaPrfmDltDurMinDltGet{X: x.AnaPrfmDltAct(scp, X.X)}
	case *xpr.AnaPrfmDltDurMaxDltGet:
		return AnaPrfmDltDurMaxDltGet{X: x.AnaPrfmDltAct(scp, X.X)}
	case *xpr.AnaPrfmDltTrdCntDltGet:
		return AnaPrfmDltTrdCntDltGet{X: x.AnaPrfmDltAct(scp, X.X)}
	case *xpr.AnaPrfmDltTrdPctAGet:
		return AnaPrfmDltTrdPctAGet{X: x.AnaPrfmDltAct(scp, X.X)}
	case *xpr.AnaPrfmDltTrdPctBGet:
		return AnaPrfmDltTrdPctBGet{X: x.AnaPrfmDltAct(scp, X.X)}
	case *xpr.AnaPrfmDltTrdPctDltGet:
		return AnaPrfmDltTrdPctDltGet{X: x.AnaPrfmDltAct(scp, X.X)}
	case *xpr.PltFltAxisYMinSetGet:
		if X.I0 == nil {
			return PltFltAxisYMinSetGet{X: x.PltFltAxisYAct(scp, X.X)}
		} else {
			return PltFltAxisYMinSetGet{X: x.PltFltAxisYAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
		}
	case *xpr.PltFltAxisYMaxSetGet:
		if X.I0 == nil {
			return PltFltAxisYMaxSetGet{X: x.PltFltAxisYAct(scp, X.X)}
		} else {
			return PltFltAxisYMaxSetGet{X: x.PltFltAxisYAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
		}
	case *xpr.PltFltAxisYEqiDstSetGet:
		if X.I0 == nil {
			return PltFltAxisYEqiDstSetGet{X: x.PltFltAxisYAct(scp, X.X)}
		} else {
			return PltFltAxisYEqiDstSetGet{X: x.PltFltAxisYAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
		}
	case *xpr.FltZero:
		return FltZero{}
	case *xpr.FltOne:
		return FltOne{}
	case *xpr.FltNegOne:
		return FltNegOne{}
	case *xpr.FltHndrd:
		return FltHndrd{}
	case *xpr.FltMin:
		return FltMin{}
	case *xpr.FltMax:
		return FltMax{}
	case *xpr.FltTiny:
		return FltTiny{}
	case *xpr.FltScl:
		return FltScl{}
	case *xpr.FltOutlierLim:
		return FltOutlierLim{}
	case *xpr.FltFltTrnc:
		return FltFltTrnc{X: x.FltFltAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.FltFltPct:
		return FltFltPct{X: x.FltFltAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.FltFltPos:
		return FltFltPos{X: x.FltFltAct(scp, X.X)}
	case *xpr.FltFltNeg:
		return FltFltNeg{X: x.FltFltAct(scp, X.X)}
	case *xpr.FltFltInv:
		return FltFltInv{X: x.FltFltAct(scp, X.X)}
	case *xpr.FltFltAdd:
		return FltFltAdd{X: x.FltFltAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.FltFltSub:
		return FltFltSub{X: x.FltFltAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.FltFltMul:
		return FltFltMul{X: x.FltFltAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.FltFltDiv:
		return FltFltDiv{X: x.FltFltAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.FltFltRem:
		return FltFltRem{X: x.FltFltAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.FltFltPow:
		return FltFltPow{X: x.FltFltAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.FltFltSqr:
		return FltFltSqr{X: x.FltFltAct(scp, X.X)}
	case *xpr.FltFltSqrt:
		return FltFltSqrt{X: x.FltFltAct(scp, X.X)}
	case *xpr.FltFltMin:
		return FltFltMin{X: x.FltFltAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.FltFltMax:
		return FltFltMax{X: x.FltFltAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.FltFltMid:
		return FltFltMid{X: x.FltFltAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.FltFltAvg:
		return FltFltAvg{X: x.FltFltAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.FltFltAvgGeo:
		return FltFltAvgGeo{X: x.FltFltAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.FltFltSelEql:
		return FltFltSelEql{X: x.FltFltAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.FltFltSelNeq:
		return FltFltSelNeq{X: x.FltFltAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.FltFltSelLss:
		return FltFltSelLss{X: x.FltFltAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.FltFltSelGtr:
		return FltFltSelGtr{X: x.FltFltAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.FltFltSelLeq:
		return FltFltSelLeq{X: x.FltFltAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.FltFltSelGeq:
		return FltFltSelGeq{X: x.FltFltAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.FltRngLen:
		return FltRngLen{X: x.FltRngAct(scp, X.X)}
	case *xpr.FltsFltsPop:
		return FltsFltsPop{X: x.FltsFltsAct(scp, X.X)}
	case *xpr.FltsFltsDque:
		return FltsFltsDque{X: x.FltsFltsAct(scp, X.X)}
	case *xpr.FltsFltsDel:
		return FltsFltsDel{X: x.FltsFltsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.FltsFltsAt:
		return FltsFltsAt{X: x.FltsFltsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.FltsFltsFst:
		return FltsFltsFst{X: x.FltsFltsAct(scp, X.X)}
	case *xpr.FltsFltsMdl:
		return FltsFltsMdl{X: x.FltsFltsAct(scp, X.X)}
	case *xpr.FltsFltsLst:
		return FltsFltsLst{X: x.FltsFltsAct(scp, X.X)}
	case *xpr.FltsFltsCntEql:
		return FltsFltsCntEql{X: x.FltsFltsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.FltsFltsCntNeq:
		return FltsFltsCntNeq{X: x.FltsFltsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.FltsFltsCntLss:
		return FltsFltsCntLss{X: x.FltsFltsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.FltsFltsCntGtr:
		return FltsFltsCntGtr{X: x.FltsFltsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.FltsFltsCntLeq:
		return FltsFltsCntLeq{X: x.FltsFltsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.FltsFltsCntGeq:
		return FltsFltsCntGeq{X: x.FltsFltsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.FltsFltsSum:
		return FltsFltsSum{X: x.FltsFltsAct(scp, X.X)}
	case *xpr.FltsFltsPrd:
		return FltsFltsPrd{X: x.FltsFltsAct(scp, X.X)}
	case *xpr.FltsFltsMin:
		return FltsFltsMin{X: x.FltsFltsAct(scp, X.X)}
	case *xpr.FltsFltsMax:
		return FltsFltsMax{X: x.FltsFltsAct(scp, X.X)}
	case *xpr.FltsFltsMid:
		return FltsFltsMid{X: x.FltsFltsAct(scp, X.X)}
	case *xpr.FltsFltsMdn:
		return FltsFltsMdn{X: x.FltsFltsAct(scp, X.X)}
	case *xpr.FltsFltsSma:
		return FltsFltsSma{X: x.FltsFltsAct(scp, X.X)}
	case *xpr.FltsFltsGma:
		return FltsFltsGma{X: x.FltsFltsAct(scp, X.X)}
	case *xpr.FltsFltsWma:
		return FltsFltsWma{X: x.FltsFltsAct(scp, X.X)}
	case *xpr.FltsFltsVrnc:
		return FltsFltsVrnc{X: x.FltsFltsAct(scp, X.X)}
	case *xpr.FltsFltsStd:
		return FltsFltsStd{X: x.FltsFltsAct(scp, X.X)}
	case *xpr.FltsFltsRngFul:
		return FltsFltsRngFul{X: x.FltsFltsAct(scp, X.X)}
	case *xpr.FltsFltsRngLst:
		return FltsFltsRngLst{X: x.FltsFltsAct(scp, X.X)}
	case *xpr.FltsFltsProLst:
		return FltsFltsProLst{X: x.FltsFltsAct(scp, X.X)}
	case *xpr.FltsFltsProSma:
		return FltsFltsProSma{X: x.FltsFltsAct(scp, X.X)}
	case *xpr.FltsFltsSubSumPos:
		return FltsFltsSubSumPos{X: x.FltsFltsAct(scp, X.X)}
	case *xpr.FltsFltsSubSumNeg:
		return FltsFltsSubSumNeg{X: x.FltsFltsAct(scp, X.X)}
	case *xpr.FltsFltsRsi:
		return FltsFltsRsi{X: x.FltsFltsAct(scp, X.X)}
	case *xpr.FltsFltsWrsi:
		return FltsFltsWrsi{X: x.FltsFltsAct(scp, X.X)}
	case *xpr.FltsFltsAlma:
		return FltsFltsAlma{X: x.FltsFltsAct(scp, X.X)}
	case *xpr.FltsFltsProAlma:
		return FltsFltsProAlma{X: x.FltsFltsAct(scp, X.X)}
	case *xpr.AnaTrdOpnMid:
		return AnaTrdOpnMid{X: x.AnaTrdAct(scp, X.X)}
	case *xpr.AnaTrdClsMid:
		return AnaTrdClsMid{X: x.AnaTrdAct(scp, X.X)}
	}
	panic(x.Erf("FltFltAct: no action found %v", reflect.ValueOf(v).Elem().Type().Name()))
}
func (x *Actr) UntUntAct(scp *Scp, v xpr.UntUntXpr) UntUntAct {
	switch X := v.(type) {
	case *xpr.UntUntLit:
		return UntUntLit{Trm: X.Trm, Txt: x.Txt}
	case *xpr.UntUntAsn:
		asnScp := scp
		if X.Depth != 0 {
			for d := X.Depth; asnScp != nil && d != 0; d-- { // recurse up scp by depth
				asnScp = asnScp.Prnt
			}
		}
		return UntUntAsn{UntScp: asnScp.UntUnt(x.Txt[X.Idn.Idx:X.Idn.Lim]), X: x.UntUntAct(scp, X.X)}
	case *xpr.UntUntAcs:
		return UntUntAcs{UntScp: scp.UntUnt(x.Txt[X.Trm.Idx:X.Trm.Lim])}
	case *xpr.AnaPrfmDayCntGet:
		return AnaPrfmDayCntGet{X: x.AnaPrfmAct(scp, X.X)}
	case *xpr.AnaPrfmTrdCntGet:
		return AnaPrfmTrdCntGet{X: x.AnaPrfmAct(scp, X.X)}
	case *xpr.AnaPrfmDltTrdCntAGet:
		return AnaPrfmDltTrdCntAGet{X: x.AnaPrfmDltAct(scp, X.X)}
	case *xpr.AnaPrfmDltTrdCntBGet:
		return AnaPrfmDltTrdCntBGet{X: x.AnaPrfmDltAct(scp, X.X)}
	case *xpr.PenPenWidSetGet:
		if X.I0 == nil {
			return PenPenWidSetGet{X: x.PenPenAct(scp, X.X)}
		} else {
			return PenPenWidSetGet{X: x.PenPenAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
		}
	case *xpr.UntZero:
		return UntZero{}
	case *xpr.UntOne:
		return UntOne{}
	case *xpr.UntMin:
		return UntMin{}
	case *xpr.UntMax:
		return UntMax{}
	case *xpr.UntStkWidth:
		return UntStkWidth{}
	case *xpr.UntShpRadius:
		return UntShpRadius{}
	case *xpr.UntAxisPad:
		return UntAxisPad{}
	case *xpr.UntBarPad:
		return UntBarPad{}
	case *xpr.UntLen:
		return UntLen{}
	case *xpr.UntPad:
		return UntPad{}
	case *xpr.UntBrdrLen:
		return UntBrdrLen{}
	case *xpr.UntInrvlTxtLen:
		return UntInrvlTxtLen{}
	case *xpr.UntUntAdd:
		return UntUntAdd{X: x.UntUntAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.UntUntSub:
		return UntUntSub{X: x.UntUntAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.UntUntMul:
		return UntUntMul{X: x.UntUntAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.UntUntDiv:
		return UntUntDiv{X: x.UntUntAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.UntUntRem:
		return UntUntRem{X: x.UntUntAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.UntUntPow:
		return UntUntPow{X: x.UntUntAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.UntUntSqr:
		return UntUntSqr{X: x.UntUntAct(scp, X.X)}
	case *xpr.UntUntSqrt:
		return UntUntSqrt{X: x.UntUntAct(scp, X.X)}
	case *xpr.UntUntMin:
		return UntUntMin{X: x.UntUntAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.UntUntMax:
		return UntUntMax{X: x.UntUntAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.UntUntMid:
		return UntUntMid{X: x.UntUntAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.UntUntAvg:
		return UntUntAvg{X: x.UntUntAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.UntUntAvgGeo:
		return UntUntAvgGeo{X: x.UntUntAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.TmeTmeWeekdayCnt:
		return TmeTmeWeekdayCnt{X: x.TmeTmeAct(scp, X.X), I0: x.TmeTmeAct(scp, X.I0)}
	case *xpr.BndBndCnt:
		return BndBndCnt{X: x.BndBndAct(scp, X.X)}
	case *xpr.BndBndLen:
		return BndBndLen{X: x.BndBndAct(scp, X.X)}
	case *xpr.BndBndLstIdx:
		return BndBndLstIdx{X: x.BndBndAct(scp, X.X)}
	case *xpr.StrsStrsCnt:
		return StrsStrsCnt{X: x.StrsStrsAct(scp, X.X)}
	case *xpr.StrsStrsFstIdx:
		return StrsStrsFstIdx{X: x.StrsStrsAct(scp, X.X)}
	case *xpr.StrsStrsMdlIdx:
		return StrsStrsMdlIdx{X: x.StrsStrsAct(scp, X.X)}
	case *xpr.StrsStrsLstIdx:
		return StrsStrsLstIdx{X: x.StrsStrsAct(scp, X.X)}
	case *xpr.StrsStrsSrchIdxEql:
		return StrsStrsSrchIdxEql{X: x.StrsStrsAct(scp, X.X), I0: x.StrStrAct(scp, X.I0)}
	case *xpr.StrsStrsSrchIdx:
		var i1 []BolBolAct
		for _, cur := range X.I1 {
			i1 = append(i1, x.BolBolAct(scp, cur))
		}
		return StrsStrsSrchIdx{X: x.StrsStrsAct(scp, X.X), I0: x.StrStrAct(scp, X.I0), I1: i1}
	case *xpr.BolsBolsCnt:
		return BolsBolsCnt{X: x.BolsBolsAct(scp, X.X)}
	case *xpr.BolsBolsFstIdx:
		return BolsBolsFstIdx{X: x.BolsBolsAct(scp, X.X)}
	case *xpr.BolsBolsMdlIdx:
		return BolsBolsMdlIdx{X: x.BolsBolsAct(scp, X.X)}
	case *xpr.BolsBolsLstIdx:
		return BolsBolsLstIdx{X: x.BolsBolsAct(scp, X.X)}
	case *xpr.FltsFltsCnt:
		return FltsFltsCnt{X: x.FltsFltsAct(scp, X.X)}
	case *xpr.FltsFltsFstIdx:
		return FltsFltsFstIdx{X: x.FltsFltsAct(scp, X.X)}
	case *xpr.FltsFltsMdlIdx:
		return FltsFltsMdlIdx{X: x.FltsFltsAct(scp, X.X)}
	case *xpr.FltsFltsLstIdx:
		return FltsFltsLstIdx{X: x.FltsFltsAct(scp, X.X)}
	case *xpr.FltsFltsSrchIdxEql:
		return FltsFltsSrchIdxEql{X: x.FltsFltsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.FltsFltsSrchIdx:
		var i1 []BolBolAct
		for _, cur := range X.I1 {
			i1 = append(i1, x.BolBolAct(scp, cur))
		}
		return FltsFltsSrchIdx{X: x.FltsFltsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0), I1: i1}
	case *xpr.UntsUntsCnt:
		return UntsUntsCnt{X: x.UntsUntsAct(scp, X.X)}
	case *xpr.UntsUntsPop:
		return UntsUntsPop{X: x.UntsUntsAct(scp, X.X)}
	case *xpr.UntsUntsDque:
		return UntsUntsDque{X: x.UntsUntsAct(scp, X.X)}
	case *xpr.UntsUntsDel:
		return UntsUntsDel{X: x.UntsUntsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.UntsUntsAt:
		return UntsUntsAt{X: x.UntsUntsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.UntsUntsFst:
		return UntsUntsFst{X: x.UntsUntsAct(scp, X.X)}
	case *xpr.UntsUntsMdl:
		return UntsUntsMdl{X: x.UntsUntsAct(scp, X.X)}
	case *xpr.UntsUntsLst:
		return UntsUntsLst{X: x.UntsUntsAct(scp, X.X)}
	case *xpr.UntsUntsFstIdx:
		return UntsUntsFstIdx{X: x.UntsUntsAct(scp, X.X)}
	case *xpr.UntsUntsMdlIdx:
		return UntsUntsMdlIdx{X: x.UntsUntsAct(scp, X.X)}
	case *xpr.UntsUntsLstIdx:
		return UntsUntsLstIdx{X: x.UntsUntsAct(scp, X.X)}
	case *xpr.UntsUntsSrchIdxEql:
		return UntsUntsSrchIdxEql{X: x.UntsUntsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.UntsUntsSrchIdx:
		var i1 []BolBolAct
		for _, cur := range X.I1 {
			i1 = append(i1, x.BolBolAct(scp, cur))
		}
		return UntsUntsSrchIdx{X: x.UntsUntsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: i1}
	case *xpr.UntsUntsSum:
		return UntsUntsSum{X: x.UntsUntsAct(scp, X.X)}
	case *xpr.UntsUntsPrd:
		return UntsUntsPrd{X: x.UntsUntsAct(scp, X.X)}
	case *xpr.UntsUntsMin:
		return UntsUntsMin{X: x.UntsUntsAct(scp, X.X)}
	case *xpr.UntsUntsMax:
		return UntsUntsMax{X: x.UntsUntsAct(scp, X.X)}
	case *xpr.UntsUntsMid:
		return UntsUntsMid{X: x.UntsUntsAct(scp, X.X)}
	case *xpr.UntsUntsMdn:
		return UntsUntsMdn{X: x.UntsUntsAct(scp, X.X)}
	case *xpr.UntsUntsSma:
		return UntsUntsSma{X: x.UntsUntsAct(scp, X.X)}
	case *xpr.UntsUntsGma:
		return UntsUntsGma{X: x.UntsUntsAct(scp, X.X)}
	case *xpr.UntsUntsWma:
		return UntsUntsWma{X: x.UntsUntsAct(scp, X.X)}
	case *xpr.UntsUntsVrnc:
		return UntsUntsVrnc{X: x.UntsUntsAct(scp, X.X)}
	case *xpr.UntsUntsStd:
		return UntsUntsStd{X: x.UntsUntsAct(scp, X.X)}
	case *xpr.UntsUntsRngFul:
		return UntsUntsRngFul{X: x.UntsUntsAct(scp, X.X)}
	case *xpr.UntsUntsRngLst:
		return UntsUntsRngLst{X: x.UntsUntsAct(scp, X.X)}
	case *xpr.UntsUntsProLst:
		return UntsUntsProLst{X: x.UntsUntsAct(scp, X.X)}
	case *xpr.UntsUntsProSma:
		return UntsUntsProSma{X: x.UntsUntsAct(scp, X.X)}
	case *xpr.IntsIntsCnt:
		return IntsIntsCnt{X: x.IntsIntsAct(scp, X.X)}
	case *xpr.IntsIntsFstIdx:
		return IntsIntsFstIdx{X: x.IntsIntsAct(scp, X.X)}
	case *xpr.IntsIntsMdlIdx:
		return IntsIntsMdlIdx{X: x.IntsIntsAct(scp, X.X)}
	case *xpr.IntsIntsLstIdx:
		return IntsIntsLstIdx{X: x.IntsIntsAct(scp, X.X)}
	case *xpr.IntsIntsSrchIdxEql:
		return IntsIntsSrchIdxEql{X: x.IntsIntsAct(scp, X.X), I0: x.IntIntAct(scp, X.I0)}
	case *xpr.IntsIntsSrchIdx:
		var i1 []BolBolAct
		for _, cur := range X.I1 {
			i1 = append(i1, x.BolBolAct(scp, cur))
		}
		return IntsIntsSrchIdx{X: x.IntsIntsAct(scp, X.X), I0: x.IntIntAct(scp, X.I0), I1: i1}
	case *xpr.TmesTmesWeekdayCnt:
		return TmesTmesWeekdayCnt{X: x.TmesTmesAct(scp, X.X)}
	case *xpr.TmesTmesCnt:
		return TmesTmesCnt{X: x.TmesTmesAct(scp, X.X)}
	case *xpr.TmesTmesFstIdx:
		return TmesTmesFstIdx{X: x.TmesTmesAct(scp, X.X)}
	case *xpr.TmesTmesMdlIdx:
		return TmesTmesMdlIdx{X: x.TmesTmesAct(scp, X.X)}
	case *xpr.TmesTmesLstIdx:
		return TmesTmesLstIdx{X: x.TmesTmesAct(scp, X.X)}
	case *xpr.TmesTmesSrchIdxEql:
		return TmesTmesSrchIdxEql{X: x.TmesTmesAct(scp, X.X), I0: x.TmeTmeAct(scp, X.I0)}
	case *xpr.TmesTmesSrchIdx:
		var i1 []BolBolAct
		for _, cur := range X.I1 {
			i1 = append(i1, x.BolBolAct(scp, cur))
		}
		return TmesTmesSrchIdx{X: x.TmesTmesAct(scp, X.X), I0: x.TmeTmeAct(scp, X.I0), I1: i1}
	case *xpr.BndsBndsCnt:
		return BndsBndsCnt{X: x.BndsBndsAct(scp, X.X)}
	case *xpr.BndsBndsFstIdx:
		return BndsBndsFstIdx{X: x.BndsBndsAct(scp, X.X)}
	case *xpr.BndsBndsMdlIdx:
		return BndsBndsMdlIdx{X: x.BndsBndsAct(scp, X.X)}
	case *xpr.BndsBndsLstIdx:
		return BndsBndsLstIdx{X: x.BndsBndsAct(scp, X.X)}
	case *xpr.TmeRngsCnt:
		return TmeRngsCnt{X: x.TmeRngsAct(scp, X.X)}
	case *xpr.TmeRngsFstIdx:
		return TmeRngsFstIdx{X: x.TmeRngsAct(scp, X.X)}
	case *xpr.TmeRngsMdlIdx:
		return TmeRngsMdlIdx{X: x.TmeRngsAct(scp, X.X)}
	case *xpr.TmeRngsLstIdx:
		return TmeRngsLstIdx{X: x.TmeRngsAct(scp, X.X)}
	case *xpr.TmeRngsSrchIdx:
		return TmeRngsSrchIdx{X: x.TmeRngsAct(scp, X.X), I0: x.TmeTmeAct(scp, X.I0)}
	case *xpr.AnaTrdsCnt:
		return AnaTrdsCnt{X: x.AnaTrdsAct(scp, X.X)}
	case *xpr.AnaTrdsFstIdx:
		return AnaTrdsFstIdx{X: x.AnaTrdsAct(scp, X.X)}
	case *xpr.AnaTrdsMdlIdx:
		return AnaTrdsMdlIdx{X: x.AnaTrdsAct(scp, X.X)}
	case *xpr.AnaTrdsLstIdx:
		return AnaTrdsLstIdx{X: x.AnaTrdsAct(scp, X.X)}
	case *xpr.AnaPrfmsCnt:
		return AnaPrfmsCnt{X: x.AnaPrfmsAct(scp, X.X)}
	case *xpr.AnaPrfmsFstIdx:
		return AnaPrfmsFstIdx{X: x.AnaPrfmsAct(scp, X.X)}
	case *xpr.AnaPrfmsMdlIdx:
		return AnaPrfmsMdlIdx{X: x.AnaPrfmsAct(scp, X.X)}
	case *xpr.AnaPrfmsLstIdx:
		return AnaPrfmsLstIdx{X: x.AnaPrfmsAct(scp, X.X)}
	case *xpr.HstPrvsCnt:
		return HstPrvsCnt{X: x.HstPrvsAct(scp, X.X)}
	case *xpr.HstPrvsFstIdx:
		return HstPrvsFstIdx{X: x.HstPrvsAct(scp, X.X)}
	case *xpr.HstPrvsMdlIdx:
		return HstPrvsMdlIdx{X: x.HstPrvsAct(scp, X.X)}
	case *xpr.HstPrvsLstIdx:
		return HstPrvsLstIdx{X: x.HstPrvsAct(scp, X.X)}
	case *xpr.HstInstrsCnt:
		return HstInstrsCnt{X: x.HstInstrsAct(scp, X.X)}
	case *xpr.HstInstrsFstIdx:
		return HstInstrsFstIdx{X: x.HstInstrsAct(scp, X.X)}
	case *xpr.HstInstrsMdlIdx:
		return HstInstrsMdlIdx{X: x.HstInstrsAct(scp, X.X)}
	case *xpr.HstInstrsLstIdx:
		return HstInstrsLstIdx{X: x.HstInstrsAct(scp, X.X)}
	case *xpr.HstInrvlsCnt:
		return HstInrvlsCnt{X: x.HstInrvlsAct(scp, X.X)}
	case *xpr.HstInrvlsFstIdx:
		return HstInrvlsFstIdx{X: x.HstInrvlsAct(scp, X.X)}
	case *xpr.HstInrvlsMdlIdx:
		return HstInrvlsMdlIdx{X: x.HstInrvlsAct(scp, X.X)}
	case *xpr.HstInrvlsLstIdx:
		return HstInrvlsLstIdx{X: x.HstInrvlsAct(scp, X.X)}
	case *xpr.HstSidesCnt:
		return HstSidesCnt{X: x.HstSidesAct(scp, X.X)}
	case *xpr.HstSidesFstIdx:
		return HstSidesFstIdx{X: x.HstSidesAct(scp, X.X)}
	case *xpr.HstSidesMdlIdx:
		return HstSidesMdlIdx{X: x.HstSidesAct(scp, X.X)}
	case *xpr.HstSidesLstIdx:
		return HstSidesLstIdx{X: x.HstSidesAct(scp, X.X)}
	case *xpr.HstStmsCnt:
		return HstStmsCnt{X: x.HstStmsAct(scp, X.X)}
	case *xpr.HstStmsFstIdx:
		return HstStmsFstIdx{X: x.HstStmsAct(scp, X.X)}
	case *xpr.HstStmsMdlIdx:
		return HstStmsMdlIdx{X: x.HstStmsAct(scp, X.X)}
	case *xpr.HstStmsLstIdx:
		return HstStmsLstIdx{X: x.HstStmsAct(scp, X.X)}
	case *xpr.HstCndsCnt:
		return HstCndsCnt{X: x.HstCndsAct(scp, X.X)}
	case *xpr.HstCndsFstIdx:
		return HstCndsFstIdx{X: x.HstCndsAct(scp, X.X)}
	case *xpr.HstCndsMdlIdx:
		return HstCndsMdlIdx{X: x.HstCndsAct(scp, X.X)}
	case *xpr.HstCndsLstIdx:
		return HstCndsLstIdx{X: x.HstCndsAct(scp, X.X)}
	case *xpr.HstStgysCnt:
		return HstStgysCnt{X: x.HstStgysAct(scp, X.X)}
	case *xpr.HstStgysFstIdx:
		return HstStgysFstIdx{X: x.HstStgysAct(scp, X.X)}
	case *xpr.HstStgysMdlIdx:
		return HstStgysMdlIdx{X: x.HstStgysAct(scp, X.X)}
	case *xpr.HstStgysLstIdx:
		return HstStgysLstIdx{X: x.HstStgysAct(scp, X.X)}
	case *xpr.RltPrvsCnt:
		return RltPrvsCnt{X: x.RltPrvsAct(scp, X.X)}
	case *xpr.RltPrvsFstIdx:
		return RltPrvsFstIdx{X: x.RltPrvsAct(scp, X.X)}
	case *xpr.RltPrvsMdlIdx:
		return RltPrvsMdlIdx{X: x.RltPrvsAct(scp, X.X)}
	case *xpr.RltPrvsLstIdx:
		return RltPrvsLstIdx{X: x.RltPrvsAct(scp, X.X)}
	case *xpr.RltInstrsCnt:
		return RltInstrsCnt{X: x.RltInstrsAct(scp, X.X)}
	case *xpr.RltInstrsFstIdx:
		return RltInstrsFstIdx{X: x.RltInstrsAct(scp, X.X)}
	case *xpr.RltInstrsMdlIdx:
		return RltInstrsMdlIdx{X: x.RltInstrsAct(scp, X.X)}
	case *xpr.RltInstrsLstIdx:
		return RltInstrsLstIdx{X: x.RltInstrsAct(scp, X.X)}
	case *xpr.RltInrvlsCnt:
		return RltInrvlsCnt{X: x.RltInrvlsAct(scp, X.X)}
	case *xpr.RltInrvlsFstIdx:
		return RltInrvlsFstIdx{X: x.RltInrvlsAct(scp, X.X)}
	case *xpr.RltInrvlsMdlIdx:
		return RltInrvlsMdlIdx{X: x.RltInrvlsAct(scp, X.X)}
	case *xpr.RltInrvlsLstIdx:
		return RltInrvlsLstIdx{X: x.RltInrvlsAct(scp, X.X)}
	case *xpr.RltSidesCnt:
		return RltSidesCnt{X: x.RltSidesAct(scp, X.X)}
	case *xpr.RltSidesFstIdx:
		return RltSidesFstIdx{X: x.RltSidesAct(scp, X.X)}
	case *xpr.RltSidesMdlIdx:
		return RltSidesMdlIdx{X: x.RltSidesAct(scp, X.X)}
	case *xpr.RltSidesLstIdx:
		return RltSidesLstIdx{X: x.RltSidesAct(scp, X.X)}
	case *xpr.RltStmsCnt:
		return RltStmsCnt{X: x.RltStmsAct(scp, X.X)}
	case *xpr.RltStmsFstIdx:
		return RltStmsFstIdx{X: x.RltStmsAct(scp, X.X)}
	case *xpr.RltStmsMdlIdx:
		return RltStmsMdlIdx{X: x.RltStmsAct(scp, X.X)}
	case *xpr.RltStmsLstIdx:
		return RltStmsLstIdx{X: x.RltStmsAct(scp, X.X)}
	case *xpr.RltCndsCnt:
		return RltCndsCnt{X: x.RltCndsAct(scp, X.X)}
	case *xpr.RltCndsFstIdx:
		return RltCndsFstIdx{X: x.RltCndsAct(scp, X.X)}
	case *xpr.RltCndsMdlIdx:
		return RltCndsMdlIdx{X: x.RltCndsAct(scp, X.X)}
	case *xpr.RltCndsLstIdx:
		return RltCndsLstIdx{X: x.RltCndsAct(scp, X.X)}
	case *xpr.RltStgysCnt:
		return RltStgysCnt{X: x.RltStgysAct(scp, X.X)}
	case *xpr.RltStgysFstIdx:
		return RltStgysFstIdx{X: x.RltStgysAct(scp, X.X)}
	case *xpr.RltStgysMdlIdx:
		return RltStgysMdlIdx{X: x.RltStgysAct(scp, X.X)}
	case *xpr.RltStgysLstIdx:
		return RltStgysLstIdx{X: x.RltStgysAct(scp, X.X)}
	case *xpr.PenPensCnt:
		return PenPensCnt{X: x.PenPensAct(scp, X.X)}
	case *xpr.PenPensFstIdx:
		return PenPensFstIdx{X: x.PenPensAct(scp, X.X)}
	case *xpr.PenPensMdlIdx:
		return PenPensMdlIdx{X: x.PenPensAct(scp, X.X)}
	case *xpr.PenPensLstIdx:
		return PenPensLstIdx{X: x.PenPensAct(scp, X.X)}
	case *xpr.PltPltsCnt:
		return PltPltsCnt{X: x.PltPltsAct(scp, X.X)}
	case *xpr.PltPltsFstIdx:
		return PltPltsFstIdx{X: x.PltPltsAct(scp, X.X)}
	case *xpr.PltPltsMdlIdx:
		return PltPltsMdlIdx{X: x.PltPltsAct(scp, X.X)}
	case *xpr.PltPltsLstIdx:
		return PltPltsLstIdx{X: x.PltPltsAct(scp, X.X)}
	}
	panic(x.Erf("UntUntAct: no action found %v", reflect.ValueOf(v).Elem().Type().Name()))
}
func (x *Actr) IntIntAct(scp *Scp, v xpr.IntIntXpr) IntIntAct {
	switch X := v.(type) {
	case *xpr.IntIntLit:
		return IntIntLit{Trm: X.Trm, Txt: x.Txt}
	case *xpr.IntIntAsn:
		asnScp := scp
		if X.Depth != 0 {
			for d := X.Depth; asnScp != nil && d != 0; d-- { // recurse up scp by depth
				asnScp = asnScp.Prnt
			}
		}
		return IntIntAsn{IntScp: asnScp.IntInt(x.Txt[X.Idn.Idx:X.Idn.Lim]), X: x.IntIntAct(scp, X.X)}
	case *xpr.IntIntAcs:
		return IntIntAcs{IntScp: scp.IntInt(x.Txt[X.Trm.Idx:X.Trm.Lim])}
	case *xpr.IntZero:
		return IntZero{}
	case *xpr.IntOne:
		return IntOne{}
	case *xpr.IntNegOne:
		return IntNegOne{}
	case *xpr.IntMin:
		return IntMin{}
	case *xpr.IntMax:
		return IntMax{}
	case *xpr.IntIntPos:
		return IntIntPos{X: x.IntIntAct(scp, X.X)}
	case *xpr.IntIntNeg:
		return IntIntNeg{X: x.IntIntAct(scp, X.X)}
	case *xpr.IntIntInv:
		return IntIntInv{X: x.IntIntAct(scp, X.X)}
	case *xpr.IntIntAdd:
		return IntIntAdd{X: x.IntIntAct(scp, X.X), I0: x.IntIntAct(scp, X.I0)}
	case *xpr.IntIntSub:
		return IntIntSub{X: x.IntIntAct(scp, X.X), I0: x.IntIntAct(scp, X.I0)}
	case *xpr.IntIntMul:
		return IntIntMul{X: x.IntIntAct(scp, X.X), I0: x.IntIntAct(scp, X.I0)}
	case *xpr.IntIntDiv:
		return IntIntDiv{X: x.IntIntAct(scp, X.X), I0: x.IntIntAct(scp, X.I0)}
	case *xpr.IntIntRem:
		return IntIntRem{X: x.IntIntAct(scp, X.X), I0: x.IntIntAct(scp, X.I0)}
	case *xpr.IntIntPow:
		return IntIntPow{X: x.IntIntAct(scp, X.X), I0: x.IntIntAct(scp, X.I0)}
	case *xpr.IntIntSqr:
		return IntIntSqr{X: x.IntIntAct(scp, X.X)}
	case *xpr.IntIntSqrt:
		return IntIntSqrt{X: x.IntIntAct(scp, X.X)}
	case *xpr.IntIntMin:
		return IntIntMin{X: x.IntIntAct(scp, X.X), I0: x.IntIntAct(scp, X.I0)}
	case *xpr.IntIntMax:
		return IntIntMax{X: x.IntIntAct(scp, X.X), I0: x.IntIntAct(scp, X.I0)}
	case *xpr.IntIntMid:
		return IntIntMid{X: x.IntIntAct(scp, X.X), I0: x.IntIntAct(scp, X.I0)}
	case *xpr.IntIntAvg:
		return IntIntAvg{X: x.IntIntAct(scp, X.X), I0: x.IntIntAct(scp, X.I0)}
	case *xpr.IntIntAvgGeo:
		return IntIntAvgGeo{X: x.IntIntAct(scp, X.X), I0: x.IntIntAct(scp, X.I0)}
	case *xpr.IntsIntsPop:
		return IntsIntsPop{X: x.IntsIntsAct(scp, X.X)}
	case *xpr.IntsIntsDque:
		return IntsIntsDque{X: x.IntsIntsAct(scp, X.X)}
	case *xpr.IntsIntsDel:
		return IntsIntsDel{X: x.IntsIntsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.IntsIntsAt:
		return IntsIntsAt{X: x.IntsIntsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.IntsIntsFst:
		return IntsIntsFst{X: x.IntsIntsAct(scp, X.X)}
	case *xpr.IntsIntsMdl:
		return IntsIntsMdl{X: x.IntsIntsAct(scp, X.X)}
	case *xpr.IntsIntsLst:
		return IntsIntsLst{X: x.IntsIntsAct(scp, X.X)}
	}
	panic(x.Erf("IntIntAct: no action found %v", reflect.ValueOf(v).Elem().Type().Name()))
}
func (x *Actr) TmeTmeAct(scp *Scp, v xpr.TmeTmeXpr) TmeTmeAct {
	switch X := v.(type) {
	case *xpr.TmeTmeLit:
		return TmeTmeLit{Trm: X.Trm, Txt: x.Txt}
	case *xpr.TmeTmeAsn:
		asnScp := scp
		if X.Depth != 0 {
			for d := X.Depth; asnScp != nil && d != 0; d-- { // recurse up scp by depth
				asnScp = asnScp.Prnt
			}
		}
		return TmeTmeAsn{TmeScp: asnScp.TmeTme(x.Txt[X.Idn.Idx:X.Idn.Lim]), X: x.TmeTmeAct(scp, X.X)}
	case *xpr.TmeTmeAcs:
		return TmeTmeAcs{TmeScp: scp.TmeTme(x.Txt[X.Trm.Idx:X.Trm.Lim])}
	case *xpr.AnaPrfmDurAvgGet:
		return AnaPrfmDurAvgGet{X: x.AnaPrfmAct(scp, X.X)}
	case *xpr.AnaPrfmDurMdnGet:
		return AnaPrfmDurMdnGet{X: x.AnaPrfmAct(scp, X.X)}
	case *xpr.AnaPrfmDurMinGet:
		return AnaPrfmDurMinGet{X: x.AnaPrfmAct(scp, X.X)}
	case *xpr.AnaPrfmDurMaxGet:
		return AnaPrfmDurMaxGet{X: x.AnaPrfmAct(scp, X.X)}
	case *xpr.AnaPrfmDurLimMaxGet:
		return AnaPrfmDurLimMaxGet{X: x.AnaPrfmAct(scp, X.X)}
	case *xpr.AnaPrfmDltDurAvgAGet:
		return AnaPrfmDltDurAvgAGet{X: x.AnaPrfmDltAct(scp, X.X)}
	case *xpr.AnaPrfmDltDurAvgBGet:
		return AnaPrfmDltDurAvgBGet{X: x.AnaPrfmDltAct(scp, X.X)}
	case *xpr.AnaPrfmDltDurMdnAGet:
		return AnaPrfmDltDurMdnAGet{X: x.AnaPrfmDltAct(scp, X.X)}
	case *xpr.AnaPrfmDltDurMdnBGet:
		return AnaPrfmDltDurMdnBGet{X: x.AnaPrfmDltAct(scp, X.X)}
	case *xpr.AnaPrfmDltDurMinAGet:
		return AnaPrfmDltDurMinAGet{X: x.AnaPrfmDltAct(scp, X.X)}
	case *xpr.AnaPrfmDltDurMinBGet:
		return AnaPrfmDltDurMinBGet{X: x.AnaPrfmDltAct(scp, X.X)}
	case *xpr.AnaPrfmDltDurMaxAGet:
		return AnaPrfmDltDurMaxAGet{X: x.AnaPrfmDltAct(scp, X.X)}
	case *xpr.AnaPrfmDltDurMaxBGet:
		return AnaPrfmDltDurMaxBGet{X: x.AnaPrfmDltAct(scp, X.X)}
	case *xpr.TmeZero:
		return TmeZero{}
	case *xpr.TmeOne:
		return TmeOne{}
	case *xpr.TmeNegOne:
		return TmeNegOne{}
	case *xpr.TmeMin:
		return TmeMin{}
	case *xpr.TmeMax:
		return TmeMax{}
	case *xpr.TmeSecond:
		return TmeSecond{}
	case *xpr.TmeMinute:
		return TmeMinute{}
	case *xpr.TmeHour:
		return TmeHour{}
	case *xpr.TmeDay:
		return TmeDay{}
	case *xpr.TmeWeek:
		return TmeWeek{}
	case *xpr.TmeS1:
		return TmeS1{}
	case *xpr.TmeS5:
		return TmeS5{}
	case *xpr.TmeS10:
		return TmeS10{}
	case *xpr.TmeS15:
		return TmeS15{}
	case *xpr.TmeS20:
		return TmeS20{}
	case *xpr.TmeS30:
		return TmeS30{}
	case *xpr.TmeS40:
		return TmeS40{}
	case *xpr.TmeS50:
		return TmeS50{}
	case *xpr.TmeM1:
		return TmeM1{}
	case *xpr.TmeM5:
		return TmeM5{}
	case *xpr.TmeM10:
		return TmeM10{}
	case *xpr.TmeM15:
		return TmeM15{}
	case *xpr.TmeM20:
		return TmeM20{}
	case *xpr.TmeM30:
		return TmeM30{}
	case *xpr.TmeM40:
		return TmeM40{}
	case *xpr.TmeM50:
		return TmeM50{}
	case *xpr.TmeH1:
		return TmeH1{}
	case *xpr.TmeD1:
		return TmeD1{}
	case *xpr.TmeResolution:
		return TmeResolution{}
	case *xpr.TmeNow:
		return TmeNow{}
	case *xpr.TmeTmeDte:
		return TmeTmeDte{X: x.TmeTmeAct(scp, X.X)}
	case *xpr.TmeTmeToSunday:
		return TmeTmeToSunday{X: x.TmeTmeAct(scp, X.X)}
	case *xpr.TmeTmeToMonday:
		return TmeTmeToMonday{X: x.TmeTmeAct(scp, X.X)}
	case *xpr.TmeTmeToTuesday:
		return TmeTmeToTuesday{X: x.TmeTmeAct(scp, X.X)}
	case *xpr.TmeTmeToWednesday:
		return TmeTmeToWednesday{X: x.TmeTmeAct(scp, X.X)}
	case *xpr.TmeTmeToThursday:
		return TmeTmeToThursday{X: x.TmeTmeAct(scp, X.X)}
	case *xpr.TmeTmeToFriday:
		return TmeTmeToFriday{X: x.TmeTmeAct(scp, X.X)}
	case *xpr.TmeTmeToSaturday:
		return TmeTmeToSaturday{X: x.TmeTmeAct(scp, X.X)}
	case *xpr.TmeTmePos:
		return TmeTmePos{X: x.TmeTmeAct(scp, X.X)}
	case *xpr.TmeTmeNeg:
		return TmeTmeNeg{X: x.TmeTmeAct(scp, X.X)}
	case *xpr.TmeTmeInv:
		return TmeTmeInv{X: x.TmeTmeAct(scp, X.X)}
	case *xpr.TmeTmeAdd:
		return TmeTmeAdd{X: x.TmeTmeAct(scp, X.X), I0: x.TmeTmeAct(scp, X.I0)}
	case *xpr.TmeTmeSub:
		return TmeTmeSub{X: x.TmeTmeAct(scp, X.X), I0: x.TmeTmeAct(scp, X.I0)}
	case *xpr.TmeTmeMul:
		return TmeTmeMul{X: x.TmeTmeAct(scp, X.X), I0: x.TmeTmeAct(scp, X.I0)}
	case *xpr.TmeTmeDiv:
		return TmeTmeDiv{X: x.TmeTmeAct(scp, X.X), I0: x.TmeTmeAct(scp, X.I0)}
	case *xpr.TmeTmeRem:
		return TmeTmeRem{X: x.TmeTmeAct(scp, X.X), I0: x.TmeTmeAct(scp, X.I0)}
	case *xpr.TmeTmePow:
		return TmeTmePow{X: x.TmeTmeAct(scp, X.X), I0: x.TmeTmeAct(scp, X.I0)}
	case *xpr.TmeTmeSqr:
		return TmeTmeSqr{X: x.TmeTmeAct(scp, X.X)}
	case *xpr.TmeTmeSqrt:
		return TmeTmeSqrt{X: x.TmeTmeAct(scp, X.X)}
	case *xpr.TmeTmeMin:
		return TmeTmeMin{X: x.TmeTmeAct(scp, X.X), I0: x.TmeTmeAct(scp, X.I0)}
	case *xpr.TmeTmeMax:
		return TmeTmeMax{X: x.TmeTmeAct(scp, X.X), I0: x.TmeTmeAct(scp, X.I0)}
	case *xpr.TmeTmeMid:
		return TmeTmeMid{X: x.TmeTmeAct(scp, X.X), I0: x.TmeTmeAct(scp, X.I0)}
	case *xpr.TmeTmeAvg:
		return TmeTmeAvg{X: x.TmeTmeAct(scp, X.X), I0: x.TmeTmeAct(scp, X.I0)}
	case *xpr.TmeTmeAvgGeo:
		return TmeTmeAvgGeo{X: x.TmeTmeAct(scp, X.X), I0: x.TmeTmeAct(scp, X.I0)}
	case *xpr.TmeRngLen:
		return TmeRngLen{X: x.TmeRngAct(scp, X.X)}
	case *xpr.TmesTmesPop:
		return TmesTmesPop{X: x.TmesTmesAct(scp, X.X)}
	case *xpr.TmesTmesDque:
		return TmesTmesDque{X: x.TmesTmesAct(scp, X.X)}
	case *xpr.TmesTmesDel:
		return TmesTmesDel{X: x.TmesTmesAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.TmesTmesAt:
		return TmesTmesAt{X: x.TmesTmesAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.TmesTmesFst:
		return TmesTmesFst{X: x.TmesTmesAct(scp, X.X)}
	case *xpr.TmesTmesMdl:
		return TmesTmesMdl{X: x.TmesTmesAct(scp, X.X)}
	case *xpr.TmesTmesLst:
		return TmesTmesLst{X: x.TmesTmesAct(scp, X.X)}
	case *xpr.TmesTmesSum:
		return TmesTmesSum{X: x.TmesTmesAct(scp, X.X)}
	case *xpr.TmesTmesPrd:
		return TmesTmesPrd{X: x.TmesTmesAct(scp, X.X)}
	case *xpr.TmesTmesMin:
		return TmesTmesMin{X: x.TmesTmesAct(scp, X.X)}
	case *xpr.TmesTmesMax:
		return TmesTmesMax{X: x.TmesTmesAct(scp, X.X)}
	case *xpr.TmesTmesMid:
		return TmesTmesMid{X: x.TmesTmesAct(scp, X.X)}
	case *xpr.TmesTmesMdn:
		return TmesTmesMdn{X: x.TmesTmesAct(scp, X.X)}
	case *xpr.TmesTmesSma:
		return TmesTmesSma{X: x.TmesTmesAct(scp, X.X)}
	case *xpr.TmesTmesGma:
		return TmesTmesGma{X: x.TmesTmesAct(scp, X.X)}
	case *xpr.TmesTmesWma:
		return TmesTmesWma{X: x.TmesTmesAct(scp, X.X)}
	case *xpr.TmesTmesVrnc:
		return TmesTmesVrnc{X: x.TmesTmesAct(scp, X.X)}
	case *xpr.TmesTmesStd:
		return TmesTmesStd{X: x.TmesTmesAct(scp, X.X)}
	case *xpr.TmesTmesRngFul:
		return TmesTmesRngFul{X: x.TmesTmesAct(scp, X.X)}
	case *xpr.TmesTmesRngLst:
		return TmesTmesRngLst{X: x.TmesTmesAct(scp, X.X)}
	case *xpr.TmesTmesProLst:
		return TmesTmesProLst{X: x.TmesTmesAct(scp, X.X)}
	case *xpr.TmesTmesProSma:
		return TmesTmesProSma{X: x.TmesTmesAct(scp, X.X)}
	}
	panic(x.Erf("TmeTmeAct: no action found %v", reflect.ValueOf(v).Elem().Type().Name()))
}
func (x *Actr) BndBndAct(scp *Scp, v xpr.BndBndXpr) BndBndAct {
	switch X := v.(type) {
	case *xpr.BndBndLit:
		return BndBndLit{Trm: X.Trm, Txt: x.Txt}
	case *xpr.BndBndAsn:
		asnScp := scp
		if X.Depth != 0 {
			for d := X.Depth; asnScp != nil && d != 0; d-- { // recurse up scp by depth
				asnScp = asnScp.Prnt
			}
		}
		return BndBndAsn{BndScp: asnScp.BndBnd(x.Txt[X.Idn.Idx:X.Idn.Lim]), X: x.BndBndAct(scp, X.X)}
	case *xpr.BndBndAcs:
		return BndBndAcs{BndScp: scp.BndBnd(x.Txt[X.Trm.Idx:X.Trm.Lim])}
	case *xpr.TmesTmesBnd:
		return TmesTmesBnd{X: x.TmesTmesAct(scp, X.X), I0: x.TmeRngAct(scp, X.I0)}
	case *xpr.BndsBndsPop:
		return BndsBndsPop{X: x.BndsBndsAct(scp, X.X)}
	case *xpr.BndsBndsDque:
		return BndsBndsDque{X: x.BndsBndsAct(scp, X.X)}
	case *xpr.BndsBndsDel:
		return BndsBndsDel{X: x.BndsBndsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.BndsBndsAt:
		return BndsBndsAt{X: x.BndsBndsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.BndsBndsFst:
		return BndsBndsFst{X: x.BndsBndsAct(scp, X.X)}
	case *xpr.BndsBndsMdl:
		return BndsBndsMdl{X: x.BndsBndsAct(scp, X.X)}
	case *xpr.BndsBndsLst:
		return BndsBndsLst{X: x.BndsBndsAct(scp, X.X)}
	}
	panic(x.Erf("BndBndAct: no action found %v", reflect.ValueOf(v).Elem().Type().Name()))
}
func (x *Actr) FltRngAct(scp *Scp, v xpr.FltRngXpr) FltRngAct {
	switch X := v.(type) {
	case *xpr.FltRngLit:
		return FltRngLit{Trm: X.Trm, Txt: x.Txt}
	case *xpr.FltRngAsn:
		asnScp := scp
		if X.Depth != 0 {
			for d := X.Depth; asnScp != nil && d != 0; d-- { // recurse up scp by depth
				asnScp = asnScp.Prnt
			}
		}
		return FltRngAsn{RngScp: asnScp.FltRng(x.Txt[X.Idn.Idx:X.Idn.Lim]), X: x.FltRngAct(scp, X.X)}
	case *xpr.FltRngAcs:
		return FltRngAcs{RngScp: scp.FltRng(x.Txt[X.Trm.Idx:X.Trm.Lim])}
	case *xpr.FltNewRng:
		return FltNewRng{I0: x.FltFltAct(scp, X.I0), I1: x.FltFltAct(scp, X.I1)}
	case *xpr.FltNewRngArnd:
		return FltNewRngArnd{I0: x.FltFltAct(scp, X.I0), I1: x.FltFltAct(scp, X.I1)}
	case *xpr.FltNewRngFul:
		return FltNewRngFul{}
	case *xpr.FltRngEnsure:
		return FltRngEnsure{X: x.FltRngAct(scp, X.X)}
	case *xpr.FltRngMinSub:
		return FltRngMinSub{X: x.FltRngAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.FltRngMaxAdd:
		return FltRngMaxAdd{X: x.FltRngAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.FltRngMrg:
		return FltRngMrg{X: x.FltRngAct(scp, X.X), I0: x.FltRngAct(scp, X.I0)}
	}
	panic(x.Erf("FltRngAct: no action found %v", reflect.ValueOf(v).Elem().Type().Name()))
}
func (x *Actr) TmeRngAct(scp *Scp, v xpr.TmeRngXpr) TmeRngAct {
	switch X := v.(type) {
	case *xpr.TmeRngLit:
		return TmeRngLit{Trm: X.Trm, Txt: x.Txt}
	case *xpr.TmeRngAsn:
		asnScp := scp
		if X.Depth != 0 {
			for d := X.Depth; asnScp != nil && d != 0; d-- { // recurse up scp by depth
				asnScp = asnScp.Prnt
			}
		}
		return TmeRngAsn{RngScp: asnScp.TmeRng(x.Txt[X.Idn.Idx:X.Idn.Lim]), X: x.TmeRngAct(scp, X.X)}
	case *xpr.TmeRngAcs:
		return TmeRngAcs{RngScp: scp.TmeRng(x.Txt[X.Trm.Idx:X.Trm.Lim])}
	case *xpr.TmeNewRng:
		return TmeNewRng{I0: x.TmeTmeAct(scp, X.I0), I1: x.TmeTmeAct(scp, X.I1)}
	case *xpr.TmeNewRngArnd:
		return TmeNewRngArnd{I0: x.TmeTmeAct(scp, X.I0), I1: x.TmeTmeAct(scp, X.I1)}
	case *xpr.TmeNewRngFul:
		return TmeNewRngFul{}
	case *xpr.TmeRngEnsure:
		return TmeRngEnsure{X: x.TmeRngAct(scp, X.X)}
	case *xpr.TmeRngMinSub:
		return TmeRngMinSub{X: x.TmeRngAct(scp, X.X), I0: x.TmeTmeAct(scp, X.I0)}
	case *xpr.TmeRngMaxAdd:
		return TmeRngMaxAdd{X: x.TmeRngAct(scp, X.X), I0: x.TmeTmeAct(scp, X.I0)}
	case *xpr.TmeRngMrg:
		return TmeRngMrg{X: x.TmeRngAct(scp, X.X), I0: x.TmeRngAct(scp, X.I0)}
	case *xpr.TmeRngsPop:
		return TmeRngsPop{X: x.TmeRngsAct(scp, X.X)}
	case *xpr.TmeRngsDque:
		return TmeRngsDque{X: x.TmeRngsAct(scp, X.X)}
	case *xpr.TmeRngsDel:
		return TmeRngsDel{X: x.TmeRngsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.TmeRngsAt:
		return TmeRngsAt{X: x.TmeRngsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.TmeRngsFst:
		return TmeRngsFst{X: x.TmeRngsAct(scp, X.X)}
	case *xpr.TmeRngsMdl:
		return TmeRngsMdl{X: x.TmeRngsAct(scp, X.X)}
	case *xpr.TmeRngsLst:
		return TmeRngsLst{X: x.TmeRngsAct(scp, X.X)}
	case *xpr.TmeRngsRngMrg:
		return TmeRngsRngMrg{X: x.TmeRngsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.UntUntAct(scp, X.I1)}
	}
	panic(x.Erf("TmeRngAct: no action found %v", reflect.ValueOf(v).Elem().Type().Name()))
}
func (x *Actr) StrsStrsAct(scp *Scp, v xpr.StrsStrsXpr) StrsStrsAct {
	switch X := v.(type) {
	case *xpr.StrsStrsLit:
		return StrsStrsLit{Trm: X.Trm, Txt: x.Txt}
	case *xpr.StrsStrsAsn:
		asnScp := scp
		if X.Depth != 0 {
			for d := X.Depth; asnScp != nil && d != 0; d-- { // recurse up scp by depth
				asnScp = asnScp.Prnt
			}
		}
		return StrsStrsAsn{StrsScp: asnScp.StrsStrs(x.Txt[X.Idn.Idx:X.Idn.Lim]), X: x.StrsStrsAct(scp, X.X)}
	case *xpr.StrsStrsAcs:
		return StrsStrsAcs{StrsScp: scp.StrsStrs(x.Txt[X.Trm.Idx:X.Trm.Lim])}
	case *xpr.StrsStrsEach:
		eachScp := NewScp(X.Scp, scp)
		return StrsStrsEach{X: x.StrsStrsAct(scp, X.X), StrScp: eachScp.StrStr(x.Txt[X.Idn.Idx:X.Idn.Lim]), Acts: x.Acts(eachScp, X.Xprs...)}
	case *xpr.StrsStrsPllEach:
		return StrsStrsPllEach{X: x.StrsStrsAct(scp, X.X), Txt: x.Txt, Idn: X.Idn, ActScpPrnt: scp, XprScp: X.Scp, Xprs: X.Xprs}
	case *xpr.StrsNew:
		var i0 []StrStrAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.StrStrAct(scp, cur))
		}
		return StrsNew{I0: i0}
	case *xpr.StrsMake:
		return StrsMake{I0: x.UntUntAct(scp, X.I0)}
	case *xpr.StrsMakeEmp:
		return StrsMakeEmp{I0: x.UntUntAct(scp, X.I0)}
	case *xpr.StrsStrsCpy:
		return StrsStrsCpy{X: x.StrsStrsAct(scp, X.X)}
	case *xpr.StrsStrsClr:
		return StrsStrsClr{X: x.StrsStrsAct(scp, X.X)}
	case *xpr.StrsStrsRand:
		return StrsStrsRand{X: x.StrsStrsAct(scp, X.X)}
	case *xpr.StrsStrsMrg:
		var i0 []StrsStrsAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.StrsStrsAct(scp, cur))
		}
		return StrsStrsMrg{X: x.StrsStrsAct(scp, X.X), I0: i0}
	case *xpr.StrsStrsPush:
		var i0 []StrStrAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.StrStrAct(scp, cur))
		}
		return StrsStrsPush{X: x.StrsStrsAct(scp, X.X), I0: i0}
	case *xpr.StrsStrsQue:
		var i0 []StrStrAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.StrStrAct(scp, cur))
		}
		return StrsStrsQue{X: x.StrsStrsAct(scp, X.X), I0: i0}
	case *xpr.StrsStrsIns:
		return StrsStrsIns{X: x.StrsStrsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.StrStrAct(scp, X.I1)}
	case *xpr.StrsStrsUpd:
		return StrsStrsUpd{X: x.StrsStrsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.StrStrAct(scp, X.I1)}
	case *xpr.StrsStrsIn:
		return StrsStrsIn{X: x.StrsStrsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.UntUntAct(scp, X.I1)}
	case *xpr.StrsStrsInBnd:
		return StrsStrsInBnd{X: x.StrsStrsAct(scp, X.X), I0: x.BndBndAct(scp, X.I0)}
	case *xpr.StrsStrsFrom:
		return StrsStrsFrom{X: x.StrsStrsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.StrsStrsTo:
		return StrsStrsTo{X: x.StrsStrsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.StrsStrsRev:
		return StrsStrsRev{X: x.StrsStrsAct(scp, X.X)}
	case *xpr.StrsStrsSrtAsc:
		return StrsStrsSrtAsc{X: x.StrsStrsAct(scp, X.X)}
	case *xpr.StrsStrsSrtDsc:
		return StrsStrsSrtDsc{X: x.StrsStrsAct(scp, X.X)}
	case *xpr.AnaTrdsClsRsns:
		return AnaTrdsClsRsns{X: x.AnaTrdsAct(scp, X.X)}
	case *xpr.AnaTrdsInstrs:
		return AnaTrdsInstrs{X: x.AnaTrdsAct(scp, X.X)}
	case *xpr.AnaTrdsOpnReqs:
		return AnaTrdsOpnReqs{X: x.AnaTrdsAct(scp, X.X)}
	case *xpr.AnaTrdsOpnRess:
		return AnaTrdsOpnRess{X: x.AnaTrdsAct(scp, X.X)}
	case *xpr.AnaTrdsClsReqs:
		return AnaTrdsClsReqs{X: x.AnaTrdsAct(scp, X.X)}
	case *xpr.AnaTrdsClsRess:
		return AnaTrdsClsRess{X: x.AnaTrdsAct(scp, X.X)}
	case *xpr.AnaPrfmsPths:
		return AnaPrfmsPths{X: x.AnaPrfmsAct(scp, X.X)}
	}
	panic(x.Erf("StrsStrsAct: no action found %v", reflect.ValueOf(v).Elem().Type().Name()))
}
func (x *Actr) BolsBolsAct(scp *Scp, v xpr.BolsBolsXpr) BolsBolsAct {
	switch X := v.(type) {
	case *xpr.BolsBolsLit:
		return BolsBolsLit{Trm: X.Trm, Txt: x.Txt}
	case *xpr.BolsBolsAsn:
		asnScp := scp
		if X.Depth != 0 {
			for d := X.Depth; asnScp != nil && d != 0; d-- { // recurse up scp by depth
				asnScp = asnScp.Prnt
			}
		}
		return BolsBolsAsn{BolsScp: asnScp.BolsBols(x.Txt[X.Idn.Idx:X.Idn.Lim]), X: x.BolsBolsAct(scp, X.X)}
	case *xpr.BolsBolsAcs:
		return BolsBolsAcs{BolsScp: scp.BolsBols(x.Txt[X.Trm.Idx:X.Trm.Lim])}
	case *xpr.BolsBolsEach:
		eachScp := NewScp(X.Scp, scp)
		return BolsBolsEach{X: x.BolsBolsAct(scp, X.X), BolScp: eachScp.BolBol(x.Txt[X.Idn.Idx:X.Idn.Lim]), Acts: x.Acts(eachScp, X.Xprs...)}
	case *xpr.BolsBolsPllEach:
		return BolsBolsPllEach{X: x.BolsBolsAct(scp, X.X), Txt: x.Txt, Idn: X.Idn, ActScpPrnt: scp, XprScp: X.Scp, Xprs: X.Xprs}
	case *xpr.BolsNew:
		var i0 []BolBolAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.BolBolAct(scp, cur))
		}
		return BolsNew{I0: i0}
	case *xpr.BolsMake:
		return BolsMake{I0: x.UntUntAct(scp, X.I0)}
	case *xpr.BolsMakeEmp:
		return BolsMakeEmp{I0: x.UntUntAct(scp, X.I0)}
	case *xpr.BolsBolsCpy:
		return BolsBolsCpy{X: x.BolsBolsAct(scp, X.X)}
	case *xpr.BolsBolsClr:
		return BolsBolsClr{X: x.BolsBolsAct(scp, X.X)}
	case *xpr.BolsBolsRand:
		return BolsBolsRand{X: x.BolsBolsAct(scp, X.X)}
	case *xpr.BolsBolsMrg:
		var i0 []BolsBolsAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.BolsBolsAct(scp, cur))
		}
		return BolsBolsMrg{X: x.BolsBolsAct(scp, X.X), I0: i0}
	case *xpr.BolsBolsPush:
		var i0 []BolBolAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.BolBolAct(scp, cur))
		}
		return BolsBolsPush{X: x.BolsBolsAct(scp, X.X), I0: i0}
	case *xpr.BolsBolsQue:
		var i0 []BolBolAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.BolBolAct(scp, cur))
		}
		return BolsBolsQue{X: x.BolsBolsAct(scp, X.X), I0: i0}
	case *xpr.BolsBolsIns:
		return BolsBolsIns{X: x.BolsBolsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.BolBolAct(scp, X.I1)}
	case *xpr.BolsBolsUpd:
		return BolsBolsUpd{X: x.BolsBolsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.BolBolAct(scp, X.I1)}
	case *xpr.BolsBolsIn:
		return BolsBolsIn{X: x.BolsBolsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.UntUntAct(scp, X.I1)}
	case *xpr.BolsBolsInBnd:
		return BolsBolsInBnd{X: x.BolsBolsAct(scp, X.X), I0: x.BndBndAct(scp, X.I0)}
	case *xpr.BolsBolsFrom:
		return BolsBolsFrom{X: x.BolsBolsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.BolsBolsTo:
		return BolsBolsTo{X: x.BolsBolsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.BolsBolsRev:
		return BolsBolsRev{X: x.BolsBolsAct(scp, X.X)}
	case *xpr.AnaTrdsIsLongs:
		return AnaTrdsIsLongs{X: x.AnaTrdsAct(scp, X.X)}
	}
	panic(x.Erf("BolsBolsAct: no action found %v", reflect.ValueOf(v).Elem().Type().Name()))
}
func (x *Actr) FltsFltsAct(scp *Scp, v xpr.FltsFltsXpr) FltsFltsAct {
	switch X := v.(type) {
	case *xpr.FltsFltsLit:
		return FltsFltsLit{Trm: X.Trm, Txt: x.Txt}
	case *xpr.FltsFltsAsn:
		asnScp := scp
		if X.Depth != 0 {
			for d := X.Depth; asnScp != nil && d != 0; d-- { // recurse up scp by depth
				asnScp = asnScp.Prnt
			}
		}
		return FltsFltsAsn{FltsScp: asnScp.FltsFlts(x.Txt[X.Idn.Idx:X.Idn.Lim]), X: x.FltsFltsAct(scp, X.X)}
	case *xpr.FltsFltsAcs:
		return FltsFltsAcs{FltsScp: scp.FltsFlts(x.Txt[X.Trm.Idx:X.Trm.Lim])}
	case *xpr.FltsFltsEach:
		eachScp := NewScp(X.Scp, scp)
		return FltsFltsEach{X: x.FltsFltsAct(scp, X.X), FltScp: eachScp.FltFlt(x.Txt[X.Idn.Idx:X.Idn.Lim]), Acts: x.Acts(eachScp, X.Xprs...)}
	case *xpr.FltsFltsPllEach:
		return FltsFltsPllEach{X: x.FltsFltsAct(scp, X.X), Txt: x.Txt, Idn: X.Idn, ActScpPrnt: scp, XprScp: X.Scp, Xprs: X.Xprs}
	case *xpr.FltsNew:
		var i0 []FltFltAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.FltFltAct(scp, cur))
		}
		return FltsNew{I0: i0}
	case *xpr.FltsMake:
		return FltsMake{I0: x.UntUntAct(scp, X.I0)}
	case *xpr.FltsMakeEmp:
		return FltsMakeEmp{I0: x.UntUntAct(scp, X.I0)}
	case *xpr.FltsAddsLss:
		return FltsAddsLss{I0: x.FltFltAct(scp, X.I0), I1: x.FltFltAct(scp, X.I1), I2: x.FltFltAct(scp, X.I2)}
	case *xpr.FltsAddsLeq:
		return FltsAddsLeq{I0: x.FltFltAct(scp, X.I0), I1: x.FltFltAct(scp, X.I1), I2: x.FltFltAct(scp, X.I2)}
	case *xpr.FltsSubsGtr:
		return FltsSubsGtr{I0: x.FltFltAct(scp, X.I0), I1: x.FltFltAct(scp, X.I1), I2: x.FltFltAct(scp, X.I2)}
	case *xpr.FltsSubsGeq:
		return FltsSubsGeq{I0: x.FltFltAct(scp, X.I0), I1: x.FltFltAct(scp, X.I1), I2: x.FltFltAct(scp, X.I2)}
	case *xpr.FltsMulsLss:
		return FltsMulsLss{I0: x.FltFltAct(scp, X.I0), I1: x.FltFltAct(scp, X.I1), I2: x.FltFltAct(scp, X.I2)}
	case *xpr.FltsMulsLeq:
		return FltsMulsLeq{I0: x.FltFltAct(scp, X.I0), I1: x.FltFltAct(scp, X.I1), I2: x.FltFltAct(scp, X.I2)}
	case *xpr.FltsDivsGtr:
		return FltsDivsGtr{I0: x.FltFltAct(scp, X.I0), I1: x.FltFltAct(scp, X.I1), I2: x.FltFltAct(scp, X.I2)}
	case *xpr.FltsDivsGeq:
		return FltsDivsGeq{I0: x.FltFltAct(scp, X.I0), I1: x.FltFltAct(scp, X.I1), I2: x.FltFltAct(scp, X.I2)}
	case *xpr.FltsFibsLeq:
		return FltsFibsLeq{I0: x.FltFltAct(scp, X.I0)}
	case *xpr.FltsFltsCpy:
		return FltsFltsCpy{X: x.FltsFltsAct(scp, X.X)}
	case *xpr.FltsFltsClr:
		return FltsFltsClr{X: x.FltsFltsAct(scp, X.X)}
	case *xpr.FltsFltsRand:
		return FltsFltsRand{X: x.FltsFltsAct(scp, X.X)}
	case *xpr.FltsFltsMrg:
		var i0 []FltsFltsAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.FltsFltsAct(scp, cur))
		}
		return FltsFltsMrg{X: x.FltsFltsAct(scp, X.X), I0: i0}
	case *xpr.FltsFltsPush:
		var i0 []FltFltAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.FltFltAct(scp, cur))
		}
		return FltsFltsPush{X: x.FltsFltsAct(scp, X.X), I0: i0}
	case *xpr.FltsFltsQue:
		var i0 []FltFltAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.FltFltAct(scp, cur))
		}
		return FltsFltsQue{X: x.FltsFltsAct(scp, X.X), I0: i0}
	case *xpr.FltsFltsIns:
		return FltsFltsIns{X: x.FltsFltsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.FltFltAct(scp, X.I1)}
	case *xpr.FltsFltsUpd:
		return FltsFltsUpd{X: x.FltsFltsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.FltFltAct(scp, X.I1)}
	case *xpr.FltsFltsIn:
		return FltsFltsIn{X: x.FltsFltsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.UntUntAct(scp, X.I1)}
	case *xpr.FltsFltsInBnd:
		return FltsFltsInBnd{X: x.FltsFltsAct(scp, X.X), I0: x.BndBndAct(scp, X.I0)}
	case *xpr.FltsFltsFrom:
		return FltsFltsFrom{X: x.FltsFltsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.FltsFltsTo:
		return FltsFltsTo{X: x.FltsFltsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.FltsFltsRev:
		return FltsFltsRev{X: x.FltsFltsAct(scp, X.X)}
	case *xpr.FltsFltsSrtAsc:
		return FltsFltsSrtAsc{X: x.FltsFltsAct(scp, X.X)}
	case *xpr.FltsFltsSrtDsc:
		return FltsFltsSrtDsc{X: x.FltsFltsAct(scp, X.X)}
	case *xpr.FltsFltsUnaPos:
		return FltsFltsUnaPos{X: x.FltsFltsAct(scp, X.X)}
	case *xpr.FltsFltsUnaNeg:
		return FltsFltsUnaNeg{X: x.FltsFltsAct(scp, X.X)}
	case *xpr.FltsFltsUnaInv:
		return FltsFltsUnaInv{X: x.FltsFltsAct(scp, X.X)}
	case *xpr.FltsFltsUnaSqr:
		return FltsFltsUnaSqr{X: x.FltsFltsAct(scp, X.X)}
	case *xpr.FltsFltsUnaSqrt:
		return FltsFltsUnaSqrt{X: x.FltsFltsAct(scp, X.X)}
	case *xpr.FltsFltsSclAdd:
		return FltsFltsSclAdd{X: x.FltsFltsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.FltsFltsSclSub:
		return FltsFltsSclSub{X: x.FltsFltsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.FltsFltsSclMul:
		return FltsFltsSclMul{X: x.FltsFltsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.FltsFltsSclDiv:
		return FltsFltsSclDiv{X: x.FltsFltsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.FltsFltsSclRem:
		return FltsFltsSclRem{X: x.FltsFltsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.FltsFltsSclPow:
		return FltsFltsSclPow{X: x.FltsFltsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.FltsFltsSclMin:
		return FltsFltsSclMin{X: x.FltsFltsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.FltsFltsSclMax:
		return FltsFltsSclMax{X: x.FltsFltsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.FltsFltsSelEql:
		return FltsFltsSelEql{X: x.FltsFltsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.FltsFltsSelNeq:
		return FltsFltsSelNeq{X: x.FltsFltsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.FltsFltsSelLss:
		return FltsFltsSelLss{X: x.FltsFltsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.FltsFltsSelGtr:
		return FltsFltsSelGtr{X: x.FltsFltsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.FltsFltsSelLeq:
		return FltsFltsSelLeq{X: x.FltsFltsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.FltsFltsSelGeq:
		return FltsFltsSelGeq{X: x.FltsFltsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.FltsFltsInrAdd:
		return FltsFltsInrAdd{X: x.FltsFltsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.FltsFltsInrSub:
		return FltsFltsInrSub{X: x.FltsFltsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.FltsFltsInrMul:
		return FltsFltsInrMul{X: x.FltsFltsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.FltsFltsInrDiv:
		return FltsFltsInrDiv{X: x.FltsFltsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.FltsFltsInrRem:
		return FltsFltsInrRem{X: x.FltsFltsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.FltsFltsInrPow:
		return FltsFltsInrPow{X: x.FltsFltsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.FltsFltsInrMin:
		return FltsFltsInrMin{X: x.FltsFltsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.FltsFltsInrMax:
		return FltsFltsInrMax{X: x.FltsFltsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.FltsFltsZscr:
		return FltsFltsZscr{X: x.FltsFltsAct(scp, X.X)}
	case *xpr.FltsFltsZscrInplace:
		return FltsFltsZscrInplace{X: x.FltsFltsAct(scp, X.X)}
	case *xpr.FltsFltsPro:
		return FltsFltsPro{X: x.FltsFltsAct(scp, X.X)}
	case *xpr.FltsFltsCntrDist:
		var i0 []BolBolAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.BolBolAct(scp, cur))
		}
		return FltsFltsCntrDist{X: x.FltsFltsAct(scp, X.X), I0: i0}
	case *xpr.AnaTrdsOpnBids:
		return AnaTrdsOpnBids{X: x.AnaTrdsAct(scp, X.X)}
	case *xpr.AnaTrdsClsBids:
		return AnaTrdsClsBids{X: x.AnaTrdsAct(scp, X.X)}
	case *xpr.AnaTrdsOpnAsks:
		return AnaTrdsOpnAsks{X: x.AnaTrdsAct(scp, X.X)}
	case *xpr.AnaTrdsClsAsks:
		return AnaTrdsClsAsks{X: x.AnaTrdsAct(scp, X.X)}
	case *xpr.AnaTrdsOpnSpds:
		return AnaTrdsOpnSpds{X: x.AnaTrdsAct(scp, X.X)}
	case *xpr.AnaTrdsClsSpds:
		return AnaTrdsClsSpds{X: x.AnaTrdsAct(scp, X.X)}
	case *xpr.AnaTrdsPips:
		return AnaTrdsPips{X: x.AnaTrdsAct(scp, X.X)}
	case *xpr.AnaTrdsPnlPcts:
		return AnaTrdsPnlPcts{X: x.AnaTrdsAct(scp, X.X)}
	case *xpr.AnaTrdsPnlPctPredicts:
		return AnaTrdsPnlPctPredicts{X: x.AnaTrdsAct(scp, X.X)}
	case *xpr.AnaTrdsPnlUsds:
		return AnaTrdsPnlUsds{X: x.AnaTrdsAct(scp, X.X)}
	case *xpr.AnaTrdsPnlGrsUsds:
		return AnaTrdsPnlGrsUsds{X: x.AnaTrdsAct(scp, X.X)}
	case *xpr.AnaTrdsCstComUsds:
		return AnaTrdsCstComUsds{X: x.AnaTrdsAct(scp, X.X)}
	case *xpr.AnaTrdsCstClsSpdUsds:
		return AnaTrdsCstClsSpdUsds{X: x.AnaTrdsAct(scp, X.X)}
	case *xpr.AnaTrdsCstOpnSpdUsds:
		return AnaTrdsCstOpnSpdUsds{X: x.AnaTrdsAct(scp, X.X)}
	case *xpr.AnaTrdsOpnBalUsds:
		return AnaTrdsOpnBalUsds{X: x.AnaTrdsAct(scp, X.X)}
	case *xpr.AnaTrdsClsBalUsds:
		return AnaTrdsClsBalUsds{X: x.AnaTrdsAct(scp, X.X)}
	case *xpr.AnaTrdsClsBalUsdActs:
		return AnaTrdsClsBalUsdActs{X: x.AnaTrdsAct(scp, X.X)}
	case *xpr.AnaTrdsTrdPcts:
		return AnaTrdsTrdPcts{X: x.AnaTrdsAct(scp, X.X)}
	case *xpr.AnaTrdsMrgnRtios:
		return AnaTrdsMrgnRtios{X: x.AnaTrdsAct(scp, X.X)}
	case *xpr.AnaTrdsUnitss:
		return AnaTrdsUnitss{X: x.AnaTrdsAct(scp, X.X)}
	case *xpr.AnaPrfmsPnlPcts:
		return AnaPrfmsPnlPcts{X: x.AnaPrfmsAct(scp, X.X)}
	case *xpr.AnaPrfmsScsPcts:
		return AnaPrfmsScsPcts{X: x.AnaPrfmsAct(scp, X.X)}
	case *xpr.AnaPrfmsPipPerDays:
		return AnaPrfmsPipPerDays{X: x.AnaPrfmsAct(scp, X.X)}
	case *xpr.AnaPrfmsUsdPerDays:
		return AnaPrfmsUsdPerDays{X: x.AnaPrfmsAct(scp, X.X)}
	case *xpr.AnaPrfmsScsPerDays:
		return AnaPrfmsScsPerDays{X: x.AnaPrfmsAct(scp, X.X)}
	case *xpr.AnaPrfmsOpnPerDays:
		return AnaPrfmsOpnPerDays{X: x.AnaPrfmsAct(scp, X.X)}
	case *xpr.AnaPrfmsPnlUsds:
		return AnaPrfmsPnlUsds{X: x.AnaPrfmsAct(scp, X.X)}
	case *xpr.AnaPrfmsPipAvgs:
		return AnaPrfmsPipAvgs{X: x.AnaPrfmsAct(scp, X.X)}
	case *xpr.AnaPrfmsPipMdns:
		return AnaPrfmsPipMdns{X: x.AnaPrfmsAct(scp, X.X)}
	case *xpr.AnaPrfmsPipMins:
		return AnaPrfmsPipMins{X: x.AnaPrfmsAct(scp, X.X)}
	case *xpr.AnaPrfmsPipMaxs:
		return AnaPrfmsPipMaxs{X: x.AnaPrfmsAct(scp, X.X)}
	case *xpr.AnaPrfmsPipSums:
		return AnaPrfmsPipSums{X: x.AnaPrfmsAct(scp, X.X)}
	case *xpr.AnaPrfmsLosLimMaxs:
		return AnaPrfmsLosLimMaxs{X: x.AnaPrfmsAct(scp, X.X)}
	case *xpr.AnaPrfmsTrdPcts:
		return AnaPrfmsTrdPcts{X: x.AnaPrfmsAct(scp, X.X)}
	case *xpr.AnaPrfmsCstTotUsds:
		return AnaPrfmsCstTotUsds{X: x.AnaPrfmsAct(scp, X.X)}
	case *xpr.AnaPrfmsCstSpdUsds:
		return AnaPrfmsCstSpdUsds{X: x.AnaPrfmsAct(scp, X.X)}
	case *xpr.AnaPrfmsCstComUsds:
		return AnaPrfmsCstComUsds{X: x.AnaPrfmsAct(scp, X.X)}
	case *xpr.HstStmAt:
		return HstStmAt{X: x.HstStmAct(scp, X.X), I0: x.TmesTmesAct(scp, X.I0)}
	}
	panic(x.Erf("FltsFltsAct: no action found %v", reflect.ValueOf(v).Elem().Type().Name()))
}
func (x *Actr) UntsUntsAct(scp *Scp, v xpr.UntsUntsXpr) UntsUntsAct {
	switch X := v.(type) {
	case *xpr.UntsUntsLit:
		return UntsUntsLit{Trm: X.Trm, Txt: x.Txt}
	case *xpr.UntsUntsAsn:
		asnScp := scp
		if X.Depth != 0 {
			for d := X.Depth; asnScp != nil && d != 0; d-- { // recurse up scp by depth
				asnScp = asnScp.Prnt
			}
		}
		return UntsUntsAsn{UntsScp: asnScp.UntsUnts(x.Txt[X.Idn.Idx:X.Idn.Lim]), X: x.UntsUntsAct(scp, X.X)}
	case *xpr.UntsUntsAcs:
		return UntsUntsAcs{UntsScp: scp.UntsUnts(x.Txt[X.Trm.Idx:X.Trm.Lim])}
	case *xpr.UntsUntsEach:
		eachScp := NewScp(X.Scp, scp)
		return UntsUntsEach{X: x.UntsUntsAct(scp, X.X), UntScp: eachScp.UntUnt(x.Txt[X.Idn.Idx:X.Idn.Lim]), Acts: x.Acts(eachScp, X.Xprs...)}
	case *xpr.UntsUntsPllEach:
		return UntsUntsPllEach{X: x.UntsUntsAct(scp, X.X), Txt: x.Txt, Idn: X.Idn, ActScpPrnt: scp, XprScp: X.Scp, Xprs: X.Xprs}
	case *xpr.UntsNew:
		var i0 []UntUntAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.UntUntAct(scp, cur))
		}
		return UntsNew{I0: i0}
	case *xpr.UntsMake:
		return UntsMake{I0: x.UntUntAct(scp, X.I0)}
	case *xpr.UntsMakeEmp:
		return UntsMakeEmp{I0: x.UntUntAct(scp, X.I0)}
	case *xpr.UntsAddsLss:
		return UntsAddsLss{I0: x.UntUntAct(scp, X.I0), I1: x.UntUntAct(scp, X.I1), I2: x.UntUntAct(scp, X.I2)}
	case *xpr.UntsAddsLeq:
		return UntsAddsLeq{I0: x.UntUntAct(scp, X.I0), I1: x.UntUntAct(scp, X.I1), I2: x.UntUntAct(scp, X.I2)}
	case *xpr.UntsSubsGtr:
		return UntsSubsGtr{I0: x.UntUntAct(scp, X.I0), I1: x.UntUntAct(scp, X.I1), I2: x.UntUntAct(scp, X.I2)}
	case *xpr.UntsSubsGeq:
		return UntsSubsGeq{I0: x.UntUntAct(scp, X.I0), I1: x.UntUntAct(scp, X.I1), I2: x.UntUntAct(scp, X.I2)}
	case *xpr.UntsMulsLss:
		return UntsMulsLss{I0: x.UntUntAct(scp, X.I0), I1: x.UntUntAct(scp, X.I1), I2: x.UntUntAct(scp, X.I2)}
	case *xpr.UntsMulsLeq:
		return UntsMulsLeq{I0: x.UntUntAct(scp, X.I0), I1: x.UntUntAct(scp, X.I1), I2: x.UntUntAct(scp, X.I2)}
	case *xpr.UntsDivsGtr:
		return UntsDivsGtr{I0: x.UntUntAct(scp, X.I0), I1: x.UntUntAct(scp, X.I1), I2: x.UntUntAct(scp, X.I2)}
	case *xpr.UntsDivsGeq:
		return UntsDivsGeq{I0: x.UntUntAct(scp, X.I0), I1: x.UntUntAct(scp, X.I1), I2: x.UntUntAct(scp, X.I2)}
	case *xpr.UntsFibsLeq:
		return UntsFibsLeq{I0: x.UntUntAct(scp, X.I0)}
	case *xpr.UntsUntsCpy:
		return UntsUntsCpy{X: x.UntsUntsAct(scp, X.X)}
	case *xpr.UntsUntsClr:
		return UntsUntsClr{X: x.UntsUntsAct(scp, X.X)}
	case *xpr.UntsUntsRand:
		return UntsUntsRand{X: x.UntsUntsAct(scp, X.X)}
	case *xpr.UntsUntsMrg:
		var i0 []UntsUntsAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.UntsUntsAct(scp, cur))
		}
		return UntsUntsMrg{X: x.UntsUntsAct(scp, X.X), I0: i0}
	case *xpr.UntsUntsPush:
		var i0 []UntUntAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.UntUntAct(scp, cur))
		}
		return UntsUntsPush{X: x.UntsUntsAct(scp, X.X), I0: i0}
	case *xpr.UntsUntsQue:
		var i0 []UntUntAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.UntUntAct(scp, cur))
		}
		return UntsUntsQue{X: x.UntsUntsAct(scp, X.X), I0: i0}
	case *xpr.UntsUntsIns:
		return UntsUntsIns{X: x.UntsUntsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.UntUntAct(scp, X.I1)}
	case *xpr.UntsUntsUpd:
		return UntsUntsUpd{X: x.UntsUntsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.UntUntAct(scp, X.I1)}
	case *xpr.UntsUntsIn:
		return UntsUntsIn{X: x.UntsUntsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.UntUntAct(scp, X.I1)}
	case *xpr.UntsUntsInBnd:
		return UntsUntsInBnd{X: x.UntsUntsAct(scp, X.X), I0: x.BndBndAct(scp, X.I0)}
	case *xpr.UntsUntsFrom:
		return UntsUntsFrom{X: x.UntsUntsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.UntsUntsTo:
		return UntsUntsTo{X: x.UntsUntsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.UntsUntsRev:
		return UntsUntsRev{X: x.UntsUntsAct(scp, X.X)}
	case *xpr.UntsUntsSrtAsc:
		return UntsUntsSrtAsc{X: x.UntsUntsAct(scp, X.X)}
	case *xpr.UntsUntsSrtDsc:
		return UntsUntsSrtDsc{X: x.UntsUntsAct(scp, X.X)}
	case *xpr.UntsUntsInrAdd:
		return UntsUntsInrAdd{X: x.UntsUntsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.UntsUntsInrSub:
		return UntsUntsInrSub{X: x.UntsUntsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.UntsUntsInrMul:
		return UntsUntsInrMul{X: x.UntsUntsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.UntsUntsInrDiv:
		return UntsUntsInrDiv{X: x.UntsUntsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.UntsUntsInrRem:
		return UntsUntsInrRem{X: x.UntsUntsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.UntsUntsInrPow:
		return UntsUntsInrPow{X: x.UntsUntsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.UntsUntsInrMin:
		return UntsUntsInrMin{X: x.UntsUntsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.UntsUntsInrMax:
		return UntsUntsInrMax{X: x.UntsUntsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.UntsUntsZscr:
		return UntsUntsZscr{X: x.UntsUntsAct(scp, X.X)}
	case *xpr.UntsUntsZscrInplace:
		return UntsUntsZscrInplace{X: x.UntsUntsAct(scp, X.X)}
	case *xpr.AnaPrfmsDayCnts:
		return AnaPrfmsDayCnts{X: x.AnaPrfmsAct(scp, X.X)}
	case *xpr.AnaPrfmsTrdCnts:
		return AnaPrfmsTrdCnts{X: x.AnaPrfmsAct(scp, X.X)}
	}
	panic(x.Erf("UntsUntsAct: no action found %v", reflect.ValueOf(v).Elem().Type().Name()))
}
func (x *Actr) IntsIntsAct(scp *Scp, v xpr.IntsIntsXpr) IntsIntsAct {
	switch X := v.(type) {
	case *xpr.IntsIntsLit:
		return IntsIntsLit{Trm: X.Trm, Txt: x.Txt}
	case *xpr.IntsIntsAsn:
		asnScp := scp
		if X.Depth != 0 {
			for d := X.Depth; asnScp != nil && d != 0; d-- { // recurse up scp by depth
				asnScp = asnScp.Prnt
			}
		}
		return IntsIntsAsn{IntsScp: asnScp.IntsInts(x.Txt[X.Idn.Idx:X.Idn.Lim]), X: x.IntsIntsAct(scp, X.X)}
	case *xpr.IntsIntsAcs:
		return IntsIntsAcs{IntsScp: scp.IntsInts(x.Txt[X.Trm.Idx:X.Trm.Lim])}
	case *xpr.IntsIntsEach:
		eachScp := NewScp(X.Scp, scp)
		return IntsIntsEach{X: x.IntsIntsAct(scp, X.X), IntScp: eachScp.IntInt(x.Txt[X.Idn.Idx:X.Idn.Lim]), Acts: x.Acts(eachScp, X.Xprs...)}
	case *xpr.IntsIntsPllEach:
		return IntsIntsPllEach{X: x.IntsIntsAct(scp, X.X), Txt: x.Txt, Idn: X.Idn, ActScpPrnt: scp, XprScp: X.Scp, Xprs: X.Xprs}
	case *xpr.IntsNew:
		var i0 []IntIntAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.IntIntAct(scp, cur))
		}
		return IntsNew{I0: i0}
	case *xpr.IntsMake:
		return IntsMake{I0: x.UntUntAct(scp, X.I0)}
	case *xpr.IntsMakeEmp:
		return IntsMakeEmp{I0: x.UntUntAct(scp, X.I0)}
	case *xpr.IntsIntsCpy:
		return IntsIntsCpy{X: x.IntsIntsAct(scp, X.X)}
	case *xpr.IntsIntsClr:
		return IntsIntsClr{X: x.IntsIntsAct(scp, X.X)}
	case *xpr.IntsIntsRand:
		return IntsIntsRand{X: x.IntsIntsAct(scp, X.X)}
	case *xpr.IntsIntsMrg:
		var i0 []IntsIntsAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.IntsIntsAct(scp, cur))
		}
		return IntsIntsMrg{X: x.IntsIntsAct(scp, X.X), I0: i0}
	case *xpr.IntsIntsPush:
		var i0 []IntIntAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.IntIntAct(scp, cur))
		}
		return IntsIntsPush{X: x.IntsIntsAct(scp, X.X), I0: i0}
	case *xpr.IntsIntsQue:
		var i0 []IntIntAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.IntIntAct(scp, cur))
		}
		return IntsIntsQue{X: x.IntsIntsAct(scp, X.X), I0: i0}
	case *xpr.IntsIntsIns:
		return IntsIntsIns{X: x.IntsIntsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.IntIntAct(scp, X.I1)}
	case *xpr.IntsIntsUpd:
		return IntsIntsUpd{X: x.IntsIntsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.IntIntAct(scp, X.I1)}
	case *xpr.IntsIntsIn:
		return IntsIntsIn{X: x.IntsIntsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.UntUntAct(scp, X.I1)}
	case *xpr.IntsIntsInBnd:
		return IntsIntsInBnd{X: x.IntsIntsAct(scp, X.X), I0: x.BndBndAct(scp, X.I0)}
	case *xpr.IntsIntsFrom:
		return IntsIntsFrom{X: x.IntsIntsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.IntsIntsTo:
		return IntsIntsTo{X: x.IntsIntsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.IntsIntsRev:
		return IntsIntsRev{X: x.IntsIntsAct(scp, X.X)}
	case *xpr.IntsIntsSrtAsc:
		return IntsIntsSrtAsc{X: x.IntsIntsAct(scp, X.X)}
	case *xpr.IntsIntsSrtDsc:
		return IntsIntsSrtDsc{X: x.IntsIntsAct(scp, X.X)}
	}
	panic(x.Erf("IntsIntsAct: no action found %v", reflect.ValueOf(v).Elem().Type().Name()))
}
func (x *Actr) TmesTmesAct(scp *Scp, v xpr.TmesTmesXpr) TmesTmesAct {
	switch X := v.(type) {
	case *xpr.TmesTmesLit:
		return TmesTmesLit{Trm: X.Trm, Txt: x.Txt}
	case *xpr.TmesTmesAsn:
		asnScp := scp
		if X.Depth != 0 {
			for d := X.Depth; asnScp != nil && d != 0; d-- { // recurse up scp by depth
				asnScp = asnScp.Prnt
			}
		}
		return TmesTmesAsn{TmesScp: asnScp.TmesTmes(x.Txt[X.Idn.Idx:X.Idn.Lim]), X: x.TmesTmesAct(scp, X.X)}
	case *xpr.TmesTmesAcs:
		return TmesTmesAcs{TmesScp: scp.TmesTmes(x.Txt[X.Trm.Idx:X.Trm.Lim])}
	case *xpr.TmesTmesEach:
		eachScp := NewScp(X.Scp, scp)
		return TmesTmesEach{X: x.TmesTmesAct(scp, X.X), TmeScp: eachScp.TmeTme(x.Txt[X.Idn.Idx:X.Idn.Lim]), Acts: x.Acts(eachScp, X.Xprs...)}
	case *xpr.TmesTmesPllEach:
		return TmesTmesPllEach{X: x.TmesTmesAct(scp, X.X), Txt: x.Txt, Idn: X.Idn, ActScpPrnt: scp, XprScp: X.Scp, Xprs: X.Xprs}
	case *xpr.TmesNew:
		var i0 []TmeTmeAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.TmeTmeAct(scp, cur))
		}
		return TmesNew{I0: i0}
	case *xpr.TmesMake:
		return TmesMake{I0: x.UntUntAct(scp, X.I0)}
	case *xpr.TmesMakeEmp:
		return TmesMakeEmp{I0: x.UntUntAct(scp, X.I0)}
	case *xpr.TmesAddsLss:
		return TmesAddsLss{I0: x.TmeTmeAct(scp, X.I0), I1: x.TmeTmeAct(scp, X.I1), I2: x.TmeTmeAct(scp, X.I2)}
	case *xpr.TmesAddsLeq:
		return TmesAddsLeq{I0: x.TmeTmeAct(scp, X.I0), I1: x.TmeTmeAct(scp, X.I1), I2: x.TmeTmeAct(scp, X.I2)}
	case *xpr.TmesSubsGtr:
		return TmesSubsGtr{I0: x.TmeTmeAct(scp, X.I0), I1: x.TmeTmeAct(scp, X.I1), I2: x.TmeTmeAct(scp, X.I2)}
	case *xpr.TmesSubsGeq:
		return TmesSubsGeq{I0: x.TmeTmeAct(scp, X.I0), I1: x.TmeTmeAct(scp, X.I1), I2: x.TmeTmeAct(scp, X.I2)}
	case *xpr.TmesMulsLss:
		return TmesMulsLss{I0: x.TmeTmeAct(scp, X.I0), I1: x.TmeTmeAct(scp, X.I1), I2: x.TmeTmeAct(scp, X.I2)}
	case *xpr.TmesMulsLeq:
		return TmesMulsLeq{I0: x.TmeTmeAct(scp, X.I0), I1: x.TmeTmeAct(scp, X.I1), I2: x.TmeTmeAct(scp, X.I2)}
	case *xpr.TmesDivsGtr:
		return TmesDivsGtr{I0: x.TmeTmeAct(scp, X.I0), I1: x.TmeTmeAct(scp, X.I1), I2: x.TmeTmeAct(scp, X.I2)}
	case *xpr.TmesDivsGeq:
		return TmesDivsGeq{I0: x.TmeTmeAct(scp, X.I0), I1: x.TmeTmeAct(scp, X.I1), I2: x.TmeTmeAct(scp, X.I2)}
	case *xpr.TmesFibsLeq:
		return TmesFibsLeq{I0: x.TmeTmeAct(scp, X.I0)}
	case *xpr.TmesTmesCpy:
		return TmesTmesCpy{X: x.TmesTmesAct(scp, X.X)}
	case *xpr.TmesTmesClr:
		return TmesTmesClr{X: x.TmesTmesAct(scp, X.X)}
	case *xpr.TmesTmesRand:
		return TmesTmesRand{X: x.TmesTmesAct(scp, X.X)}
	case *xpr.TmesTmesMrg:
		var i0 []TmesTmesAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.TmesTmesAct(scp, cur))
		}
		return TmesTmesMrg{X: x.TmesTmesAct(scp, X.X), I0: i0}
	case *xpr.TmesTmesPush:
		var i0 []TmeTmeAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.TmeTmeAct(scp, cur))
		}
		return TmesTmesPush{X: x.TmesTmesAct(scp, X.X), I0: i0}
	case *xpr.TmesTmesQue:
		var i0 []TmeTmeAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.TmeTmeAct(scp, cur))
		}
		return TmesTmesQue{X: x.TmesTmesAct(scp, X.X), I0: i0}
	case *xpr.TmesTmesIns:
		return TmesTmesIns{X: x.TmesTmesAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.TmeTmeAct(scp, X.I1)}
	case *xpr.TmesTmesUpd:
		return TmesTmesUpd{X: x.TmesTmesAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.TmeTmeAct(scp, X.I1)}
	case *xpr.TmesTmesIn:
		return TmesTmesIn{X: x.TmesTmesAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.UntUntAct(scp, X.I1)}
	case *xpr.TmesTmesInBnd:
		return TmesTmesInBnd{X: x.TmesTmesAct(scp, X.X), I0: x.BndBndAct(scp, X.I0)}
	case *xpr.TmesTmesFrom:
		return TmesTmesFrom{X: x.TmesTmesAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.TmesTmesTo:
		return TmesTmesTo{X: x.TmesTmesAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.TmesTmesRev:
		return TmesTmesRev{X: x.TmesTmesAct(scp, X.X)}
	case *xpr.TmesTmesSrtAsc:
		return TmesTmesSrtAsc{X: x.TmesTmesAct(scp, X.X)}
	case *xpr.TmesTmesSrtDsc:
		return TmesTmesSrtDsc{X: x.TmesTmesAct(scp, X.X)}
	case *xpr.TmesTmesInrAdd:
		return TmesTmesInrAdd{X: x.TmesTmesAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.TmesTmesInrSub:
		return TmesTmesInrSub{X: x.TmesTmesAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.TmesTmesInrMul:
		return TmesTmesInrMul{X: x.TmesTmesAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.TmesTmesInrDiv:
		return TmesTmesInrDiv{X: x.TmesTmesAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.TmesTmesInrRem:
		return TmesTmesInrRem{X: x.TmesTmesAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.TmesTmesInrPow:
		return TmesTmesInrPow{X: x.TmesTmesAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.TmesTmesInrMin:
		return TmesTmesInrMin{X: x.TmesTmesAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.TmesTmesInrMax:
		return TmesTmesInrMax{X: x.TmesTmesAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.TmesTmesZscr:
		return TmesTmesZscr{X: x.TmesTmesAct(scp, X.X)}
	case *xpr.TmesTmesZscrInplace:
		return TmesTmesZscrInplace{X: x.TmesTmesAct(scp, X.X)}
	case *xpr.AnaTrdsOpnTmes:
		return AnaTrdsOpnTmes{X: x.AnaTrdsAct(scp, X.X)}
	case *xpr.AnaTrdsClsTmes:
		return AnaTrdsClsTmes{X: x.AnaTrdsAct(scp, X.X)}
	case *xpr.AnaTrdsDurs:
		return AnaTrdsDurs{X: x.AnaTrdsAct(scp, X.X)}
	case *xpr.AnaPrfmsDurAvgs:
		return AnaPrfmsDurAvgs{X: x.AnaPrfmsAct(scp, X.X)}
	case *xpr.AnaPrfmsDurMdns:
		return AnaPrfmsDurMdns{X: x.AnaPrfmsAct(scp, X.X)}
	case *xpr.AnaPrfmsDurMins:
		return AnaPrfmsDurMins{X: x.AnaPrfmsAct(scp, X.X)}
	case *xpr.AnaPrfmsDurMaxs:
		return AnaPrfmsDurMaxs{X: x.AnaPrfmsAct(scp, X.X)}
	case *xpr.AnaPrfmsDurLimMaxs:
		return AnaPrfmsDurLimMaxs{X: x.AnaPrfmsAct(scp, X.X)}
	}
	panic(x.Erf("TmesTmesAct: no action found %v", reflect.ValueOf(v).Elem().Type().Name()))
}
func (x *Actr) BndsBndsAct(scp *Scp, v xpr.BndsBndsXpr) BndsBndsAct {
	switch X := v.(type) {
	case *xpr.BndsBndsLit:
		return BndsBndsLit{Trm: X.Trm, Txt: x.Txt}
	case *xpr.BndsBndsAsn:
		asnScp := scp
		if X.Depth != 0 {
			for d := X.Depth; asnScp != nil && d != 0; d-- { // recurse up scp by depth
				asnScp = asnScp.Prnt
			}
		}
		return BndsBndsAsn{BndsScp: asnScp.BndsBnds(x.Txt[X.Idn.Idx:X.Idn.Lim]), X: x.BndsBndsAct(scp, X.X)}
	case *xpr.BndsBndsAcs:
		return BndsBndsAcs{BndsScp: scp.BndsBnds(x.Txt[X.Trm.Idx:X.Trm.Lim])}
	case *xpr.BndsBndsEach:
		eachScp := NewScp(X.Scp, scp)
		return BndsBndsEach{X: x.BndsBndsAct(scp, X.X), BndScp: eachScp.BndBnd(x.Txt[X.Idn.Idx:X.Idn.Lim]), Acts: x.Acts(eachScp, X.Xprs...)}
	case *xpr.BndsBndsPllEach:
		return BndsBndsPllEach{X: x.BndsBndsAct(scp, X.X), Txt: x.Txt, Idn: X.Idn, ActScpPrnt: scp, XprScp: X.Scp, Xprs: X.Xprs}
	case *xpr.BndsNew:
		var i0 []BndBndAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.BndBndAct(scp, cur))
		}
		return BndsNew{I0: i0}
	case *xpr.BndsMake:
		return BndsMake{I0: x.UntUntAct(scp, X.I0)}
	case *xpr.BndsMakeEmp:
		return BndsMakeEmp{I0: x.UntUntAct(scp, X.I0)}
	case *xpr.BndsBndsCpy:
		return BndsBndsCpy{X: x.BndsBndsAct(scp, X.X)}
	case *xpr.BndsBndsClr:
		return BndsBndsClr{X: x.BndsBndsAct(scp, X.X)}
	case *xpr.BndsBndsRand:
		return BndsBndsRand{X: x.BndsBndsAct(scp, X.X)}
	case *xpr.BndsBndsMrg:
		var i0 []BndsBndsAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.BndsBndsAct(scp, cur))
		}
		return BndsBndsMrg{X: x.BndsBndsAct(scp, X.X), I0: i0}
	case *xpr.BndsBndsPush:
		var i0 []BndBndAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.BndBndAct(scp, cur))
		}
		return BndsBndsPush{X: x.BndsBndsAct(scp, X.X), I0: i0}
	case *xpr.BndsBndsQue:
		var i0 []BndBndAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.BndBndAct(scp, cur))
		}
		return BndsBndsQue{X: x.BndsBndsAct(scp, X.X), I0: i0}
	case *xpr.BndsBndsIns:
		return BndsBndsIns{X: x.BndsBndsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.BndBndAct(scp, X.I1)}
	case *xpr.BndsBndsUpd:
		return BndsBndsUpd{X: x.BndsBndsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.BndBndAct(scp, X.I1)}
	case *xpr.BndsBndsIn:
		return BndsBndsIn{X: x.BndsBndsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.UntUntAct(scp, X.I1)}
	case *xpr.BndsBndsInBnd:
		return BndsBndsInBnd{X: x.BndsBndsAct(scp, X.X), I0: x.BndBndAct(scp, X.I0)}
	case *xpr.BndsBndsFrom:
		return BndsBndsFrom{X: x.BndsBndsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.BndsBndsTo:
		return BndsBndsTo{X: x.BndsBndsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.BndsBndsRev:
		return BndsBndsRev{X: x.BndsBndsAct(scp, X.X)}
	}
	panic(x.Erf("BndsBndsAct: no action found %v", reflect.ValueOf(v).Elem().Type().Name()))
}
func (x *Actr) TmeRngsAct(scp *Scp, v xpr.TmeRngsXpr) TmeRngsAct {
	switch X := v.(type) {
	case *xpr.TmeRngsLit:
		return TmeRngsLit{Trm: X.Trm, Txt: x.Txt}
	case *xpr.TmeRngsAsn:
		asnScp := scp
		if X.Depth != 0 {
			for d := X.Depth; asnScp != nil && d != 0; d-- { // recurse up scp by depth
				asnScp = asnScp.Prnt
			}
		}
		return TmeRngsAsn{RngsScp: asnScp.TmeRngs(x.Txt[X.Idn.Idx:X.Idn.Lim]), X: x.TmeRngsAct(scp, X.X)}
	case *xpr.TmeRngsAcs:
		return TmeRngsAcs{RngsScp: scp.TmeRngs(x.Txt[X.Trm.Idx:X.Trm.Lim])}
	case *xpr.TmeRngsEach:
		eachScp := NewScp(X.Scp, scp)
		return TmeRngsEach{X: x.TmeRngsAct(scp, X.X), RngScp: eachScp.TmeRng(x.Txt[X.Idn.Idx:X.Idn.Lim]), Acts: x.Acts(eachScp, X.Xprs...)}
	case *xpr.TmeRngsPllEach:
		return TmeRngsPllEach{X: x.TmeRngsAct(scp, X.X), Txt: x.Txt, Idn: X.Idn, ActScpPrnt: scp, XprScp: X.Scp, Xprs: X.Xprs}
	case *xpr.TmeNewRngs:
		var i0 []TmeRngAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.TmeRngAct(scp, cur))
		}
		return TmeNewRngs{I0: i0}
	case *xpr.TmeMakeRngs:
		return TmeMakeRngs{I0: x.UntUntAct(scp, X.I0)}
	case *xpr.TmeMakeEmpRngs:
		return TmeMakeEmpRngs{I0: x.UntUntAct(scp, X.I0)}
	case *xpr.TmeRngsCpy:
		return TmeRngsCpy{X: x.TmeRngsAct(scp, X.X)}
	case *xpr.TmeRngsClr:
		return TmeRngsClr{X: x.TmeRngsAct(scp, X.X)}
	case *xpr.TmeRngsRand:
		return TmeRngsRand{X: x.TmeRngsAct(scp, X.X)}
	case *xpr.TmeRngsMrg:
		var i0 []TmeRngsAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.TmeRngsAct(scp, cur))
		}
		return TmeRngsMrg{X: x.TmeRngsAct(scp, X.X), I0: i0}
	case *xpr.TmeRngsPush:
		var i0 []TmeRngAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.TmeRngAct(scp, cur))
		}
		return TmeRngsPush{X: x.TmeRngsAct(scp, X.X), I0: i0}
	case *xpr.TmeRngsQue:
		var i0 []TmeRngAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.TmeRngAct(scp, cur))
		}
		return TmeRngsQue{X: x.TmeRngsAct(scp, X.X), I0: i0}
	case *xpr.TmeRngsIns:
		return TmeRngsIns{X: x.TmeRngsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.TmeRngAct(scp, X.I1)}
	case *xpr.TmeRngsUpd:
		return TmeRngsUpd{X: x.TmeRngsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.TmeRngAct(scp, X.I1)}
	case *xpr.TmeRngsIn:
		return TmeRngsIn{X: x.TmeRngsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.UntUntAct(scp, X.I1)}
	case *xpr.TmeRngsInBnd:
		return TmeRngsInBnd{X: x.TmeRngsAct(scp, X.X), I0: x.BndBndAct(scp, X.I0)}
	case *xpr.TmeRngsFrom:
		return TmeRngsFrom{X: x.TmeRngsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.TmeRngsTo:
		return TmeRngsTo{X: x.TmeRngsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.TmeRngsRev:
		return TmeRngsRev{X: x.TmeRngsAct(scp, X.X)}
	}
	panic(x.Erf("TmeRngsAct: no action found %v", reflect.ValueOf(v).Elem().Type().Name()))
}
func (x *Actr) AnaTrdAct(scp *Scp, v xpr.AnaTrdXpr) AnaTrdAct {
	switch X := v.(type) {
	case *xpr.AnaTrdAsn:
		asnScp := scp
		if X.Depth != 0 {
			for d := X.Depth; asnScp != nil && d != 0; d-- { // recurse up scp by depth
				asnScp = asnScp.Prnt
			}
		}
		return AnaTrdAsn{TrdScp: asnScp.AnaTrd(x.Txt[X.Idn.Idx:X.Idn.Lim]), X: x.AnaTrdAct(scp, X.X)}
	case *xpr.AnaTrdAcs:
		return AnaTrdAcs{TrdScp: scp.AnaTrd(x.Txt[X.Trm.Idx:X.Trm.Lim])}
	case *xpr.AnaTrdsPop:
		return AnaTrdsPop{X: x.AnaTrdsAct(scp, X.X)}
	case *xpr.AnaTrdsDque:
		return AnaTrdsDque{X: x.AnaTrdsAct(scp, X.X)}
	case *xpr.AnaTrdsDel:
		return AnaTrdsDel{X: x.AnaTrdsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.AnaTrdsAt:
		return AnaTrdsAt{X: x.AnaTrdsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.AnaTrdsFst:
		return AnaTrdsFst{X: x.AnaTrdsAct(scp, X.X)}
	case *xpr.AnaTrdsMdl:
		return AnaTrdsMdl{X: x.AnaTrdsAct(scp, X.X)}
	case *xpr.AnaTrdsLst:
		return AnaTrdsLst{X: x.AnaTrdsAct(scp, X.X)}
	}
	panic(x.Erf("AnaTrdAct: no action found %v", reflect.ValueOf(v).Elem().Type().Name()))
}
func (x *Actr) AnaTrdsAct(scp *Scp, v xpr.AnaTrdsXpr) AnaTrdsAct {
	switch X := v.(type) {
	case *xpr.AnaTrdsAsn:
		asnScp := scp
		if X.Depth != 0 {
			for d := X.Depth; asnScp != nil && d != 0; d-- { // recurse up scp by depth
				asnScp = asnScp.Prnt
			}
		}
		return AnaTrdsAsn{TrdsScp: asnScp.AnaTrds(x.Txt[X.Idn.Idx:X.Idn.Lim]), X: x.AnaTrdsAct(scp, X.X)}
	case *xpr.AnaTrdsAcs:
		return AnaTrdsAcs{TrdsScp: scp.AnaTrds(x.Txt[X.Trm.Idx:X.Trm.Lim])}
	case *xpr.AnaTrdsEach:
		eachScp := NewScp(X.Scp, scp)
		return AnaTrdsEach{X: x.AnaTrdsAct(scp, X.X), TrdScp: eachScp.AnaTrd(x.Txt[X.Idn.Idx:X.Idn.Lim]), Acts: x.Acts(eachScp, X.Xprs...)}
	case *xpr.AnaTrdsPllEach:
		return AnaTrdsPllEach{X: x.AnaTrdsAct(scp, X.X), Txt: x.Txt, Idn: X.Idn, ActScpPrnt: scp, XprScp: X.Scp, Xprs: X.Xprs}
	case *xpr.AnaNewTrds:
		var i0 []AnaTrdAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.AnaTrdAct(scp, cur))
		}
		return AnaNewTrds{I0: i0}
	case *xpr.AnaMakeTrds:
		return AnaMakeTrds{I0: x.UntUntAct(scp, X.I0)}
	case *xpr.AnaMakeEmpTrds:
		return AnaMakeEmpTrds{I0: x.UntUntAct(scp, X.I0)}
	case *xpr.AnaTrdsCpy:
		return AnaTrdsCpy{X: x.AnaTrdsAct(scp, X.X)}
	case *xpr.AnaTrdsClr:
		return AnaTrdsClr{X: x.AnaTrdsAct(scp, X.X)}
	case *xpr.AnaTrdsRand:
		return AnaTrdsRand{X: x.AnaTrdsAct(scp, X.X)}
	case *xpr.AnaTrdsMrg:
		var i0 []AnaTrdsAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.AnaTrdsAct(scp, cur))
		}
		return AnaTrdsMrg{X: x.AnaTrdsAct(scp, X.X), I0: i0}
	case *xpr.AnaTrdsPush:
		var i0 []AnaTrdAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.AnaTrdAct(scp, cur))
		}
		return AnaTrdsPush{X: x.AnaTrdsAct(scp, X.X), I0: i0}
	case *xpr.AnaTrdsQue:
		var i0 []AnaTrdAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.AnaTrdAct(scp, cur))
		}
		return AnaTrdsQue{X: x.AnaTrdsAct(scp, X.X), I0: i0}
	case *xpr.AnaTrdsIns:
		return AnaTrdsIns{X: x.AnaTrdsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.AnaTrdAct(scp, X.I1)}
	case *xpr.AnaTrdsUpd:
		return AnaTrdsUpd{X: x.AnaTrdsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.AnaTrdAct(scp, X.I1)}
	case *xpr.AnaTrdsIn:
		return AnaTrdsIn{X: x.AnaTrdsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.UntUntAct(scp, X.I1)}
	case *xpr.AnaTrdsInBnd:
		return AnaTrdsInBnd{X: x.AnaTrdsAct(scp, X.X), I0: x.BndBndAct(scp, X.I0)}
	case *xpr.AnaTrdsFrom:
		return AnaTrdsFrom{X: x.AnaTrdsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.AnaTrdsTo:
		return AnaTrdsTo{X: x.AnaTrdsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.AnaTrdsRev:
		return AnaTrdsRev{X: x.AnaTrdsAct(scp, X.X)}
	case *xpr.AnaTrdsSelClsResEql:
		return AnaTrdsSelClsResEql{X: x.AnaTrdsAct(scp, X.X), I0: x.StrStrAct(scp, X.I0)}
	case *xpr.AnaTrdsSelClsResNeq:
		return AnaTrdsSelClsResNeq{X: x.AnaTrdsAct(scp, X.X), I0: x.StrStrAct(scp, X.I0)}
	case *xpr.AnaTrdsSelClsResLss:
		return AnaTrdsSelClsResLss{X: x.AnaTrdsAct(scp, X.X), I0: x.StrStrAct(scp, X.I0)}
	case *xpr.AnaTrdsSelClsResGtr:
		return AnaTrdsSelClsResGtr{X: x.AnaTrdsAct(scp, X.X), I0: x.StrStrAct(scp, X.I0)}
	case *xpr.AnaTrdsSelClsResLeq:
		return AnaTrdsSelClsResLeq{X: x.AnaTrdsAct(scp, X.X), I0: x.StrStrAct(scp, X.I0)}
	case *xpr.AnaTrdsSelClsResGeq:
		return AnaTrdsSelClsResGeq{X: x.AnaTrdsAct(scp, X.X), I0: x.StrStrAct(scp, X.I0)}
	case *xpr.AnaTrdsSelClsReqEql:
		return AnaTrdsSelClsReqEql{X: x.AnaTrdsAct(scp, X.X), I0: x.StrStrAct(scp, X.I0)}
	case *xpr.AnaTrdsSelClsReqNeq:
		return AnaTrdsSelClsReqNeq{X: x.AnaTrdsAct(scp, X.X), I0: x.StrStrAct(scp, X.I0)}
	case *xpr.AnaTrdsSelClsReqLss:
		return AnaTrdsSelClsReqLss{X: x.AnaTrdsAct(scp, X.X), I0: x.StrStrAct(scp, X.I0)}
	case *xpr.AnaTrdsSelClsReqGtr:
		return AnaTrdsSelClsReqGtr{X: x.AnaTrdsAct(scp, X.X), I0: x.StrStrAct(scp, X.I0)}
	case *xpr.AnaTrdsSelClsReqLeq:
		return AnaTrdsSelClsReqLeq{X: x.AnaTrdsAct(scp, X.X), I0: x.StrStrAct(scp, X.I0)}
	case *xpr.AnaTrdsSelClsReqGeq:
		return AnaTrdsSelClsReqGeq{X: x.AnaTrdsAct(scp, X.X), I0: x.StrStrAct(scp, X.I0)}
	case *xpr.AnaTrdsSelOpnResEql:
		return AnaTrdsSelOpnResEql{X: x.AnaTrdsAct(scp, X.X), I0: x.StrStrAct(scp, X.I0)}
	case *xpr.AnaTrdsSelOpnResNeq:
		return AnaTrdsSelOpnResNeq{X: x.AnaTrdsAct(scp, X.X), I0: x.StrStrAct(scp, X.I0)}
	case *xpr.AnaTrdsSelOpnResLss:
		return AnaTrdsSelOpnResLss{X: x.AnaTrdsAct(scp, X.X), I0: x.StrStrAct(scp, X.I0)}
	case *xpr.AnaTrdsSelOpnResGtr:
		return AnaTrdsSelOpnResGtr{X: x.AnaTrdsAct(scp, X.X), I0: x.StrStrAct(scp, X.I0)}
	case *xpr.AnaTrdsSelOpnResLeq:
		return AnaTrdsSelOpnResLeq{X: x.AnaTrdsAct(scp, X.X), I0: x.StrStrAct(scp, X.I0)}
	case *xpr.AnaTrdsSelOpnResGeq:
		return AnaTrdsSelOpnResGeq{X: x.AnaTrdsAct(scp, X.X), I0: x.StrStrAct(scp, X.I0)}
	case *xpr.AnaTrdsSelOpnReqEql:
		return AnaTrdsSelOpnReqEql{X: x.AnaTrdsAct(scp, X.X), I0: x.StrStrAct(scp, X.I0)}
	case *xpr.AnaTrdsSelOpnReqNeq:
		return AnaTrdsSelOpnReqNeq{X: x.AnaTrdsAct(scp, X.X), I0: x.StrStrAct(scp, X.I0)}
	case *xpr.AnaTrdsSelOpnReqLss:
		return AnaTrdsSelOpnReqLss{X: x.AnaTrdsAct(scp, X.X), I0: x.StrStrAct(scp, X.I0)}
	case *xpr.AnaTrdsSelOpnReqGtr:
		return AnaTrdsSelOpnReqGtr{X: x.AnaTrdsAct(scp, X.X), I0: x.StrStrAct(scp, X.I0)}
	case *xpr.AnaTrdsSelOpnReqLeq:
		return AnaTrdsSelOpnReqLeq{X: x.AnaTrdsAct(scp, X.X), I0: x.StrStrAct(scp, X.I0)}
	case *xpr.AnaTrdsSelOpnReqGeq:
		return AnaTrdsSelOpnReqGeq{X: x.AnaTrdsAct(scp, X.X), I0: x.StrStrAct(scp, X.I0)}
	case *xpr.AnaTrdsSelInstrEql:
		return AnaTrdsSelInstrEql{X: x.AnaTrdsAct(scp, X.X), I0: x.StrStrAct(scp, X.I0)}
	case *xpr.AnaTrdsSelInstrNeq:
		return AnaTrdsSelInstrNeq{X: x.AnaTrdsAct(scp, X.X), I0: x.StrStrAct(scp, X.I0)}
	case *xpr.AnaTrdsSelInstrLss:
		return AnaTrdsSelInstrLss{X: x.AnaTrdsAct(scp, X.X), I0: x.StrStrAct(scp, X.I0)}
	case *xpr.AnaTrdsSelInstrGtr:
		return AnaTrdsSelInstrGtr{X: x.AnaTrdsAct(scp, X.X), I0: x.StrStrAct(scp, X.I0)}
	case *xpr.AnaTrdsSelInstrLeq:
		return AnaTrdsSelInstrLeq{X: x.AnaTrdsAct(scp, X.X), I0: x.StrStrAct(scp, X.I0)}
	case *xpr.AnaTrdsSelInstrGeq:
		return AnaTrdsSelInstrGeq{X: x.AnaTrdsAct(scp, X.X), I0: x.StrStrAct(scp, X.I0)}
	case *xpr.AnaTrdsSelUnitsEql:
		return AnaTrdsSelUnitsEql{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelUnitsNeq:
		return AnaTrdsSelUnitsNeq{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelUnitsLss:
		return AnaTrdsSelUnitsLss{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelUnitsGtr:
		return AnaTrdsSelUnitsGtr{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelUnitsLeq:
		return AnaTrdsSelUnitsLeq{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelUnitsGeq:
		return AnaTrdsSelUnitsGeq{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelMrgnRtioEql:
		return AnaTrdsSelMrgnRtioEql{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelMrgnRtioNeq:
		return AnaTrdsSelMrgnRtioNeq{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelMrgnRtioLss:
		return AnaTrdsSelMrgnRtioLss{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelMrgnRtioGtr:
		return AnaTrdsSelMrgnRtioGtr{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelMrgnRtioLeq:
		return AnaTrdsSelMrgnRtioLeq{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelMrgnRtioGeq:
		return AnaTrdsSelMrgnRtioGeq{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelTrdPctEql:
		return AnaTrdsSelTrdPctEql{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelTrdPctNeq:
		return AnaTrdsSelTrdPctNeq{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelTrdPctLss:
		return AnaTrdsSelTrdPctLss{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelTrdPctGtr:
		return AnaTrdsSelTrdPctGtr{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelTrdPctLeq:
		return AnaTrdsSelTrdPctLeq{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelTrdPctGeq:
		return AnaTrdsSelTrdPctGeq{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelClsBalUsdActEql:
		return AnaTrdsSelClsBalUsdActEql{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelClsBalUsdActNeq:
		return AnaTrdsSelClsBalUsdActNeq{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelClsBalUsdActLss:
		return AnaTrdsSelClsBalUsdActLss{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelClsBalUsdActGtr:
		return AnaTrdsSelClsBalUsdActGtr{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelClsBalUsdActLeq:
		return AnaTrdsSelClsBalUsdActLeq{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelClsBalUsdActGeq:
		return AnaTrdsSelClsBalUsdActGeq{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelClsBalUsdEql:
		return AnaTrdsSelClsBalUsdEql{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelClsBalUsdNeq:
		return AnaTrdsSelClsBalUsdNeq{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelClsBalUsdLss:
		return AnaTrdsSelClsBalUsdLss{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelClsBalUsdGtr:
		return AnaTrdsSelClsBalUsdGtr{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelClsBalUsdLeq:
		return AnaTrdsSelClsBalUsdLeq{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelClsBalUsdGeq:
		return AnaTrdsSelClsBalUsdGeq{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelOpnBalUsdEql:
		return AnaTrdsSelOpnBalUsdEql{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelOpnBalUsdNeq:
		return AnaTrdsSelOpnBalUsdNeq{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelOpnBalUsdLss:
		return AnaTrdsSelOpnBalUsdLss{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelOpnBalUsdGtr:
		return AnaTrdsSelOpnBalUsdGtr{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelOpnBalUsdLeq:
		return AnaTrdsSelOpnBalUsdLeq{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelOpnBalUsdGeq:
		return AnaTrdsSelOpnBalUsdGeq{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelCstOpnSpdUsdEql:
		return AnaTrdsSelCstOpnSpdUsdEql{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelCstOpnSpdUsdNeq:
		return AnaTrdsSelCstOpnSpdUsdNeq{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelCstOpnSpdUsdLss:
		return AnaTrdsSelCstOpnSpdUsdLss{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelCstOpnSpdUsdGtr:
		return AnaTrdsSelCstOpnSpdUsdGtr{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelCstOpnSpdUsdLeq:
		return AnaTrdsSelCstOpnSpdUsdLeq{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelCstOpnSpdUsdGeq:
		return AnaTrdsSelCstOpnSpdUsdGeq{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelCstClsSpdUsdEql:
		return AnaTrdsSelCstClsSpdUsdEql{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelCstClsSpdUsdNeq:
		return AnaTrdsSelCstClsSpdUsdNeq{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelCstClsSpdUsdLss:
		return AnaTrdsSelCstClsSpdUsdLss{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelCstClsSpdUsdGtr:
		return AnaTrdsSelCstClsSpdUsdGtr{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelCstClsSpdUsdLeq:
		return AnaTrdsSelCstClsSpdUsdLeq{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelCstClsSpdUsdGeq:
		return AnaTrdsSelCstClsSpdUsdGeq{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelCstComUsdEql:
		return AnaTrdsSelCstComUsdEql{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelCstComUsdNeq:
		return AnaTrdsSelCstComUsdNeq{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelCstComUsdLss:
		return AnaTrdsSelCstComUsdLss{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelCstComUsdGtr:
		return AnaTrdsSelCstComUsdGtr{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelCstComUsdLeq:
		return AnaTrdsSelCstComUsdLeq{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelCstComUsdGeq:
		return AnaTrdsSelCstComUsdGeq{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelPnlGrsUsdEql:
		return AnaTrdsSelPnlGrsUsdEql{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelPnlGrsUsdNeq:
		return AnaTrdsSelPnlGrsUsdNeq{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelPnlGrsUsdLss:
		return AnaTrdsSelPnlGrsUsdLss{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelPnlGrsUsdGtr:
		return AnaTrdsSelPnlGrsUsdGtr{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelPnlGrsUsdLeq:
		return AnaTrdsSelPnlGrsUsdLeq{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelPnlGrsUsdGeq:
		return AnaTrdsSelPnlGrsUsdGeq{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelPnlUsdEql:
		return AnaTrdsSelPnlUsdEql{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelPnlUsdNeq:
		return AnaTrdsSelPnlUsdNeq{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelPnlUsdLss:
		return AnaTrdsSelPnlUsdLss{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelPnlUsdGtr:
		return AnaTrdsSelPnlUsdGtr{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelPnlUsdLeq:
		return AnaTrdsSelPnlUsdLeq{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelPnlUsdGeq:
		return AnaTrdsSelPnlUsdGeq{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelPnlPctPredictEql:
		return AnaTrdsSelPnlPctPredictEql{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelPnlPctPredictNeq:
		return AnaTrdsSelPnlPctPredictNeq{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelPnlPctPredictLss:
		return AnaTrdsSelPnlPctPredictLss{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelPnlPctPredictGtr:
		return AnaTrdsSelPnlPctPredictGtr{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelPnlPctPredictLeq:
		return AnaTrdsSelPnlPctPredictLeq{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelPnlPctPredictGeq:
		return AnaTrdsSelPnlPctPredictGeq{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelPnlPctEql:
		return AnaTrdsSelPnlPctEql{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelPnlPctNeq:
		return AnaTrdsSelPnlPctNeq{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelPnlPctLss:
		return AnaTrdsSelPnlPctLss{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelPnlPctGtr:
		return AnaTrdsSelPnlPctGtr{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelPnlPctLeq:
		return AnaTrdsSelPnlPctLeq{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelPnlPctGeq:
		return AnaTrdsSelPnlPctGeq{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelIsLongEql:
		return AnaTrdsSelIsLongEql{X: x.AnaTrdsAct(scp, X.X), I0: x.BolBolAct(scp, X.I0)}
	case *xpr.AnaTrdsSelIsLongNeq:
		return AnaTrdsSelIsLongNeq{X: x.AnaTrdsAct(scp, X.X), I0: x.BolBolAct(scp, X.I0)}
	case *xpr.AnaTrdsSelDurEql:
		return AnaTrdsSelDurEql{X: x.AnaTrdsAct(scp, X.X), I0: x.TmeTmeAct(scp, X.I0)}
	case *xpr.AnaTrdsSelDurNeq:
		return AnaTrdsSelDurNeq{X: x.AnaTrdsAct(scp, X.X), I0: x.TmeTmeAct(scp, X.I0)}
	case *xpr.AnaTrdsSelDurLss:
		return AnaTrdsSelDurLss{X: x.AnaTrdsAct(scp, X.X), I0: x.TmeTmeAct(scp, X.I0)}
	case *xpr.AnaTrdsSelDurGtr:
		return AnaTrdsSelDurGtr{X: x.AnaTrdsAct(scp, X.X), I0: x.TmeTmeAct(scp, X.I0)}
	case *xpr.AnaTrdsSelDurLeq:
		return AnaTrdsSelDurLeq{X: x.AnaTrdsAct(scp, X.X), I0: x.TmeTmeAct(scp, X.I0)}
	case *xpr.AnaTrdsSelDurGeq:
		return AnaTrdsSelDurGeq{X: x.AnaTrdsAct(scp, X.X), I0: x.TmeTmeAct(scp, X.I0)}
	case *xpr.AnaTrdsSelPipEql:
		return AnaTrdsSelPipEql{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelPipNeq:
		return AnaTrdsSelPipNeq{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelPipLss:
		return AnaTrdsSelPipLss{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelPipGtr:
		return AnaTrdsSelPipGtr{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelPipLeq:
		return AnaTrdsSelPipLeq{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelPipGeq:
		return AnaTrdsSelPipGeq{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelClsRsnEql:
		return AnaTrdsSelClsRsnEql{X: x.AnaTrdsAct(scp, X.X), I0: x.StrStrAct(scp, X.I0)}
	case *xpr.AnaTrdsSelClsRsnNeq:
		return AnaTrdsSelClsRsnNeq{X: x.AnaTrdsAct(scp, X.X), I0: x.StrStrAct(scp, X.I0)}
	case *xpr.AnaTrdsSelClsRsnLss:
		return AnaTrdsSelClsRsnLss{X: x.AnaTrdsAct(scp, X.X), I0: x.StrStrAct(scp, X.I0)}
	case *xpr.AnaTrdsSelClsRsnGtr:
		return AnaTrdsSelClsRsnGtr{X: x.AnaTrdsAct(scp, X.X), I0: x.StrStrAct(scp, X.I0)}
	case *xpr.AnaTrdsSelClsRsnLeq:
		return AnaTrdsSelClsRsnLeq{X: x.AnaTrdsAct(scp, X.X), I0: x.StrStrAct(scp, X.I0)}
	case *xpr.AnaTrdsSelClsRsnGeq:
		return AnaTrdsSelClsRsnGeq{X: x.AnaTrdsAct(scp, X.X), I0: x.StrStrAct(scp, X.I0)}
	case *xpr.AnaTrdsSelClsSpdEql:
		return AnaTrdsSelClsSpdEql{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelClsSpdNeq:
		return AnaTrdsSelClsSpdNeq{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelClsSpdLss:
		return AnaTrdsSelClsSpdLss{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelClsSpdGtr:
		return AnaTrdsSelClsSpdGtr{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelClsSpdLeq:
		return AnaTrdsSelClsSpdLeq{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelClsSpdGeq:
		return AnaTrdsSelClsSpdGeq{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelOpnSpdEql:
		return AnaTrdsSelOpnSpdEql{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelOpnSpdNeq:
		return AnaTrdsSelOpnSpdNeq{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelOpnSpdLss:
		return AnaTrdsSelOpnSpdLss{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelOpnSpdGtr:
		return AnaTrdsSelOpnSpdGtr{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelOpnSpdLeq:
		return AnaTrdsSelOpnSpdLeq{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelOpnSpdGeq:
		return AnaTrdsSelOpnSpdGeq{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelClsAskEql:
		return AnaTrdsSelClsAskEql{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelClsAskNeq:
		return AnaTrdsSelClsAskNeq{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelClsAskLss:
		return AnaTrdsSelClsAskLss{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelClsAskGtr:
		return AnaTrdsSelClsAskGtr{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelClsAskLeq:
		return AnaTrdsSelClsAskLeq{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelClsAskGeq:
		return AnaTrdsSelClsAskGeq{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelOpnAskEql:
		return AnaTrdsSelOpnAskEql{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelOpnAskNeq:
		return AnaTrdsSelOpnAskNeq{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelOpnAskLss:
		return AnaTrdsSelOpnAskLss{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelOpnAskGtr:
		return AnaTrdsSelOpnAskGtr{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelOpnAskLeq:
		return AnaTrdsSelOpnAskLeq{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelOpnAskGeq:
		return AnaTrdsSelOpnAskGeq{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelClsBidEql:
		return AnaTrdsSelClsBidEql{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelClsBidNeq:
		return AnaTrdsSelClsBidNeq{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelClsBidLss:
		return AnaTrdsSelClsBidLss{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelClsBidGtr:
		return AnaTrdsSelClsBidGtr{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelClsBidLeq:
		return AnaTrdsSelClsBidLeq{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelClsBidGeq:
		return AnaTrdsSelClsBidGeq{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelOpnBidEql:
		return AnaTrdsSelOpnBidEql{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelOpnBidNeq:
		return AnaTrdsSelOpnBidNeq{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelOpnBidLss:
		return AnaTrdsSelOpnBidLss{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelOpnBidGtr:
		return AnaTrdsSelOpnBidGtr{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelOpnBidLeq:
		return AnaTrdsSelOpnBidLeq{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelOpnBidGeq:
		return AnaTrdsSelOpnBidGeq{X: x.AnaTrdsAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.AnaTrdsSelClsTmeEql:
		return AnaTrdsSelClsTmeEql{X: x.AnaTrdsAct(scp, X.X), I0: x.TmeTmeAct(scp, X.I0)}
	case *xpr.AnaTrdsSelClsTmeNeq:
		return AnaTrdsSelClsTmeNeq{X: x.AnaTrdsAct(scp, X.X), I0: x.TmeTmeAct(scp, X.I0)}
	case *xpr.AnaTrdsSelClsTmeLss:
		return AnaTrdsSelClsTmeLss{X: x.AnaTrdsAct(scp, X.X), I0: x.TmeTmeAct(scp, X.I0)}
	case *xpr.AnaTrdsSelClsTmeGtr:
		return AnaTrdsSelClsTmeGtr{X: x.AnaTrdsAct(scp, X.X), I0: x.TmeTmeAct(scp, X.I0)}
	case *xpr.AnaTrdsSelClsTmeLeq:
		return AnaTrdsSelClsTmeLeq{X: x.AnaTrdsAct(scp, X.X), I0: x.TmeTmeAct(scp, X.I0)}
	case *xpr.AnaTrdsSelClsTmeGeq:
		return AnaTrdsSelClsTmeGeq{X: x.AnaTrdsAct(scp, X.X), I0: x.TmeTmeAct(scp, X.I0)}
	case *xpr.AnaTrdsSelOpnTmeEql:
		return AnaTrdsSelOpnTmeEql{X: x.AnaTrdsAct(scp, X.X), I0: x.TmeTmeAct(scp, X.I0)}
	case *xpr.AnaTrdsSelOpnTmeNeq:
		return AnaTrdsSelOpnTmeNeq{X: x.AnaTrdsAct(scp, X.X), I0: x.TmeTmeAct(scp, X.I0)}
	case *xpr.AnaTrdsSelOpnTmeLss:
		return AnaTrdsSelOpnTmeLss{X: x.AnaTrdsAct(scp, X.X), I0: x.TmeTmeAct(scp, X.I0)}
	case *xpr.AnaTrdsSelOpnTmeGtr:
		return AnaTrdsSelOpnTmeGtr{X: x.AnaTrdsAct(scp, X.X), I0: x.TmeTmeAct(scp, X.I0)}
	case *xpr.AnaTrdsSelOpnTmeLeq:
		return AnaTrdsSelOpnTmeLeq{X: x.AnaTrdsAct(scp, X.X), I0: x.TmeTmeAct(scp, X.I0)}
	case *xpr.AnaTrdsSelOpnTmeGeq:
		return AnaTrdsSelOpnTmeGeq{X: x.AnaTrdsAct(scp, X.X), I0: x.TmeTmeAct(scp, X.I0)}
	}
	panic(x.Erf("AnaTrdsAct: no action found %v", reflect.ValueOf(v).Elem().Type().Name()))
}
func (x *Actr) AnaPrfmAct(scp *Scp, v xpr.AnaPrfmXpr) AnaPrfmAct {
	switch X := v.(type) {
	case *xpr.AnaPrfmAsn:
		asnScp := scp
		if X.Depth != 0 {
			for d := X.Depth; asnScp != nil && d != 0; d-- { // recurse up scp by depth
				asnScp = asnScp.Prnt
			}
		}
		return AnaPrfmAsn{PrfmScp: asnScp.AnaPrfm(x.Txt[X.Idn.Idx:X.Idn.Lim]), X: x.AnaPrfmAct(scp, X.X)}
	case *xpr.AnaPrfmAcs:
		return AnaPrfmAcs{PrfmScp: scp.AnaPrfm(x.Txt[X.Trm.Idx:X.Trm.Lim])}
	case *xpr.AnaPrfmsPop:
		return AnaPrfmsPop{X: x.AnaPrfmsAct(scp, X.X)}
	case *xpr.AnaPrfmsDque:
		return AnaPrfmsDque{X: x.AnaPrfmsAct(scp, X.X)}
	case *xpr.AnaPrfmsDel:
		return AnaPrfmsDel{X: x.AnaPrfmsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.AnaPrfmsAt:
		return AnaPrfmsAt{X: x.AnaPrfmsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.AnaPrfmsFst:
		return AnaPrfmsFst{X: x.AnaPrfmsAct(scp, X.X)}
	case *xpr.AnaPrfmsMdl:
		return AnaPrfmsMdl{X: x.AnaPrfmsAct(scp, X.X)}
	case *xpr.AnaPrfmsLst:
		return AnaPrfmsLst{X: x.AnaPrfmsAct(scp, X.X)}
	}
	panic(x.Erf("AnaPrfmAct: no action found %v", reflect.ValueOf(v).Elem().Type().Name()))
}
func (x *Actr) AnaPrfmsAct(scp *Scp, v xpr.AnaPrfmsXpr) AnaPrfmsAct {
	switch X := v.(type) {
	case *xpr.AnaPrfmsAsn:
		asnScp := scp
		if X.Depth != 0 {
			for d := X.Depth; asnScp != nil && d != 0; d-- { // recurse up scp by depth
				asnScp = asnScp.Prnt
			}
		}
		return AnaPrfmsAsn{PrfmsScp: asnScp.AnaPrfms(x.Txt[X.Idn.Idx:X.Idn.Lim]), X: x.AnaPrfmsAct(scp, X.X)}
	case *xpr.AnaPrfmsAcs:
		return AnaPrfmsAcs{PrfmsScp: scp.AnaPrfms(x.Txt[X.Trm.Idx:X.Trm.Lim])}
	case *xpr.AnaPrfmsEach:
		eachScp := NewScp(X.Scp, scp)
		return AnaPrfmsEach{X: x.AnaPrfmsAct(scp, X.X), PrfmScp: eachScp.AnaPrfm(x.Txt[X.Idn.Idx:X.Idn.Lim]), Acts: x.Acts(eachScp, X.Xprs...)}
	case *xpr.AnaPrfmsPllEach:
		return AnaPrfmsPllEach{X: x.AnaPrfmsAct(scp, X.X), Txt: x.Txt, Idn: X.Idn, ActScpPrnt: scp, XprScp: X.Scp, Xprs: X.Xprs}
	case *xpr.AnaNewPrfms:
		var i0 []AnaPrfmAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.AnaPrfmAct(scp, cur))
		}
		return AnaNewPrfms{I0: i0}
	case *xpr.AnaMakePrfms:
		return AnaMakePrfms{I0: x.UntUntAct(scp, X.I0)}
	case *xpr.AnaMakeEmpPrfms:
		return AnaMakeEmpPrfms{I0: x.UntUntAct(scp, X.I0)}
	case *xpr.AnaPrfmsCpy:
		return AnaPrfmsCpy{X: x.AnaPrfmsAct(scp, X.X)}
	case *xpr.AnaPrfmsClr:
		return AnaPrfmsClr{X: x.AnaPrfmsAct(scp, X.X)}
	case *xpr.AnaPrfmsRand:
		return AnaPrfmsRand{X: x.AnaPrfmsAct(scp, X.X)}
	case *xpr.AnaPrfmsMrg:
		var i0 []AnaPrfmsAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.AnaPrfmsAct(scp, cur))
		}
		return AnaPrfmsMrg{X: x.AnaPrfmsAct(scp, X.X), I0: i0}
	case *xpr.AnaPrfmsPush:
		var i0 []AnaPrfmAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.AnaPrfmAct(scp, cur))
		}
		return AnaPrfmsPush{X: x.AnaPrfmsAct(scp, X.X), I0: i0}
	case *xpr.AnaPrfmsQue:
		var i0 []AnaPrfmAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.AnaPrfmAct(scp, cur))
		}
		return AnaPrfmsQue{X: x.AnaPrfmsAct(scp, X.X), I0: i0}
	case *xpr.AnaPrfmsIns:
		return AnaPrfmsIns{X: x.AnaPrfmsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.AnaPrfmAct(scp, X.I1)}
	case *xpr.AnaPrfmsUpd:
		return AnaPrfmsUpd{X: x.AnaPrfmsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.AnaPrfmAct(scp, X.I1)}
	case *xpr.AnaPrfmsIn:
		return AnaPrfmsIn{X: x.AnaPrfmsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.UntUntAct(scp, X.I1)}
	case *xpr.AnaPrfmsInBnd:
		return AnaPrfmsInBnd{X: x.AnaPrfmsAct(scp, X.X), I0: x.BndBndAct(scp, X.I0)}
	case *xpr.AnaPrfmsFrom:
		return AnaPrfmsFrom{X: x.AnaPrfmsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.AnaPrfmsTo:
		return AnaPrfmsTo{X: x.AnaPrfmsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.AnaPrfmsRev:
		return AnaPrfmsRev{X: x.AnaPrfmsAct(scp, X.X)}
	}
	panic(x.Erf("AnaPrfmsAct: no action found %v", reflect.ValueOf(v).Elem().Type().Name()))
}
func (x *Actr) AnaPrfmDltAct(scp *Scp, v xpr.AnaPrfmDltXpr) AnaPrfmDltAct {
	switch X := v.(type) {
	case *xpr.AnaPrfmDltAsn:
		asnScp := scp
		if X.Depth != 0 {
			for d := X.Depth; asnScp != nil && d != 0; d-- { // recurse up scp by depth
				asnScp = asnScp.Prnt
			}
		}
		return AnaPrfmDltAsn{PrfmDltScp: asnScp.AnaPrfmDlt(x.Txt[X.Idn.Idx:X.Idn.Lim]), X: x.AnaPrfmDltAct(scp, X.X)}
	case *xpr.AnaPrfmDltAcs:
		return AnaPrfmDltAcs{PrfmDltScp: scp.AnaPrfmDlt(x.Txt[X.Trm.Idx:X.Trm.Lim])}
	case *xpr.AnaPrfmDlt:
		return AnaPrfmDlt{X: x.AnaPrfmAct(scp, X.X), I0: x.AnaPrfmAct(scp, X.I0)}
	}
	panic(x.Erf("AnaPrfmDltAct: no action found %v", reflect.ValueOf(v).Elem().Type().Name()))
}
func (x *Actr) AnaPortAct(scp *Scp, v xpr.AnaPortXpr) AnaPortAct {
	switch X := v.(type) {
	case *xpr.AnaPortAsn:
		asnScp := scp
		if X.Depth != 0 {
			for d := X.Depth; asnScp != nil && d != 0; d-- { // recurse up scp by depth
				asnScp = asnScp.Prnt
			}
		}
		return AnaPortAsn{PortScp: asnScp.AnaPort(x.Txt[X.Idn.Idx:X.Idn.Lim]), X: x.AnaPortAct(scp, X.X)}
	case *xpr.AnaPortAcs:
		return AnaPortAcs{PortScp: scp.AnaPort(x.Txt[X.Trm.Idx:X.Trm.Lim])}
	}
	panic(x.Erf("AnaPortAct: no action found %v", reflect.ValueOf(v).Elem().Type().Name()))
}
func (x *Actr) HstPrvAct(scp *Scp, v xpr.HstPrvXpr) HstPrvAct {
	switch X := v.(type) {
	case *xpr.HstPrvAsn:
		asnScp := scp
		if X.Depth != 0 {
			for d := X.Depth; asnScp != nil && d != 0; d-- { // recurse up scp by depth
				asnScp = asnScp.Prnt
			}
		}
		return HstPrvAsn{PrvScp: asnScp.HstPrv(x.Txt[X.Idn.Idx:X.Idn.Lim]), X: x.HstPrvAct(scp, X.X)}
	case *xpr.HstPrvAcs:
		return HstPrvAcs{PrvScp: scp.HstPrv(x.Txt[X.Trm.Idx:X.Trm.Lim])}
	case *xpr.HstOan:
		return HstOan{}
	case *xpr.HstPrvsPop:
		return HstPrvsPop{X: x.HstPrvsAct(scp, X.X)}
	case *xpr.HstPrvsDque:
		return HstPrvsDque{X: x.HstPrvsAct(scp, X.X)}
	case *xpr.HstPrvsDel:
		return HstPrvsDel{X: x.HstPrvsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.HstPrvsAt:
		return HstPrvsAt{X: x.HstPrvsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.HstPrvsFst:
		return HstPrvsFst{X: x.HstPrvsAct(scp, X.X)}
	case *xpr.HstPrvsMdl:
		return HstPrvsMdl{X: x.HstPrvsAct(scp, X.X)}
	case *xpr.HstPrvsLst:
		return HstPrvsLst{X: x.HstPrvsAct(scp, X.X)}
	}
	panic(x.Erf("HstPrvAct: no action found %v", reflect.ValueOf(v).Elem().Type().Name()))
}
func (x *Actr) HstInstrAct(scp *Scp, v xpr.HstInstrXpr) HstInstrAct {
	switch X := v.(type) {
	case *xpr.HstInstrAsn:
		asnScp := scp
		if X.Depth != 0 {
			for d := X.Depth; asnScp != nil && d != 0; d-- { // recurse up scp by depth
				asnScp = asnScp.Prnt
			}
		}
		return HstInstrAsn{InstrScp: asnScp.HstInstr(x.Txt[X.Idn.Idx:X.Idn.Lim]), X: x.HstInstrAct(scp, X.X)}
	case *xpr.HstInstrAcs:
		return HstInstrAcs{InstrScp: scp.HstInstr(x.Txt[X.Trm.Idx:X.Trm.Lim])}
	case *xpr.HstInstrsPop:
		return HstInstrsPop{X: x.HstInstrsAct(scp, X.X)}
	case *xpr.HstInstrsDque:
		return HstInstrsDque{X: x.HstInstrsAct(scp, X.X)}
	case *xpr.HstInstrsDel:
		return HstInstrsDel{X: x.HstInstrsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.HstInstrsAt:
		return HstInstrsAt{X: x.HstInstrsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.HstInstrsFst:
		return HstInstrsFst{X: x.HstInstrsAct(scp, X.X)}
	case *xpr.HstInstrsMdl:
		return HstInstrsMdl{X: x.HstInstrsAct(scp, X.X)}
	case *xpr.HstInstrsLst:
		return HstInstrsLst{X: x.HstInstrsAct(scp, X.X)}
	case *xpr.HstPrvEurUsd:
		var i0 []TmeRngAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.TmeRngAct(scp, cur))
		}
		return HstPrvEurUsd{X: x.HstPrvAct(scp, X.X), I0: i0}
	case *xpr.HstPrvAudUsd:
		var i0 []TmeRngAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.TmeRngAct(scp, cur))
		}
		return HstPrvAudUsd{X: x.HstPrvAct(scp, X.X), I0: i0}
	case *xpr.HstPrvNzdUsd:
		var i0 []TmeRngAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.TmeRngAct(scp, cur))
		}
		return HstPrvNzdUsd{X: x.HstPrvAct(scp, X.X), I0: i0}
	case *xpr.HstPrvGbpUsd:
		var i0 []TmeRngAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.TmeRngAct(scp, cur))
		}
		return HstPrvGbpUsd{X: x.HstPrvAct(scp, X.X), I0: i0}
	}
	panic(x.Erf("HstInstrAct: no action found %v", reflect.ValueOf(v).Elem().Type().Name()))
}
func (x *Actr) HstInrvlAct(scp *Scp, v xpr.HstInrvlXpr) HstInrvlAct {
	switch X := v.(type) {
	case *xpr.HstInrvlAsn:
		asnScp := scp
		if X.Depth != 0 {
			for d := X.Depth; asnScp != nil && d != 0; d-- { // recurse up scp by depth
				asnScp = asnScp.Prnt
			}
		}
		return HstInrvlAsn{InrvlScp: asnScp.HstInrvl(x.Txt[X.Idn.Idx:X.Idn.Lim]), X: x.HstInrvlAct(scp, X.X)}
	case *xpr.HstInrvlAcs:
		return HstInrvlAcs{InrvlScp: scp.HstInrvl(x.Txt[X.Trm.Idx:X.Trm.Lim])}
	case *xpr.HstInrvlsPop:
		return HstInrvlsPop{X: x.HstInrvlsAct(scp, X.X)}
	case *xpr.HstInrvlsDque:
		return HstInrvlsDque{X: x.HstInrvlsAct(scp, X.X)}
	case *xpr.HstInrvlsDel:
		return HstInrvlsDel{X: x.HstInrvlsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.HstInrvlsAt:
		return HstInrvlsAt{X: x.HstInrvlsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.HstInrvlsFst:
		return HstInrvlsFst{X: x.HstInrvlsAct(scp, X.X)}
	case *xpr.HstInrvlsMdl:
		return HstInrvlsMdl{X: x.HstInrvlsAct(scp, X.X)}
	case *xpr.HstInrvlsLst:
		return HstInrvlsLst{X: x.HstInrvlsAct(scp, X.X)}
	case *xpr.HstInstrI:
		return HstInstrI{X: x.HstInstrAct(scp, X.X), I0: x.TmeTmeAct(scp, X.I0)}
	}
	panic(x.Erf("HstInrvlAct: no action found %v", reflect.ValueOf(v).Elem().Type().Name()))
}
func (x *Actr) HstSideAct(scp *Scp, v xpr.HstSideXpr) HstSideAct {
	switch X := v.(type) {
	case *xpr.HstSideAsn:
		asnScp := scp
		if X.Depth != 0 {
			for d := X.Depth; asnScp != nil && d != 0; d-- { // recurse up scp by depth
				asnScp = asnScp.Prnt
			}
		}
		return HstSideAsn{SideScp: asnScp.HstSide(x.Txt[X.Idn.Idx:X.Idn.Lim]), X: x.HstSideAct(scp, X.X)}
	case *xpr.HstSideAcs:
		return HstSideAcs{SideScp: scp.HstSide(x.Txt[X.Trm.Idx:X.Trm.Lim])}
	case *xpr.HstSidesPop:
		return HstSidesPop{X: x.HstSidesAct(scp, X.X)}
	case *xpr.HstSidesDque:
		return HstSidesDque{X: x.HstSidesAct(scp, X.X)}
	case *xpr.HstSidesDel:
		return HstSidesDel{X: x.HstSidesAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.HstSidesAt:
		return HstSidesAt{X: x.HstSidesAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.HstSidesFst:
		return HstSidesFst{X: x.HstSidesAct(scp, X.X)}
	case *xpr.HstSidesMdl:
		return HstSidesMdl{X: x.HstSidesAct(scp, X.X)}
	case *xpr.HstSidesLst:
		return HstSidesLst{X: x.HstSidesAct(scp, X.X)}
	case *xpr.HstInrvlBid:
		return HstInrvlBid{X: x.HstInrvlAct(scp, X.X)}
	case *xpr.HstInrvlAsk:
		return HstInrvlAsk{X: x.HstInrvlAct(scp, X.X)}
	}
	panic(x.Erf("HstSideAct: no action found %v", reflect.ValueOf(v).Elem().Type().Name()))
}
func (x *Actr) HstStmAct(scp *Scp, v xpr.HstStmXpr) HstStmAct {
	switch X := v.(type) {
	case *xpr.HstStmAsn:
		asnScp := scp
		if X.Depth != 0 {
			for d := X.Depth; asnScp != nil && d != 0; d-- { // recurse up scp by depth
				asnScp = asnScp.Prnt
			}
		}
		return HstStmAsn{StmScp: asnScp.HstStm(x.Txt[X.Idn.Idx:X.Idn.Lim]), X: x.HstStmAct(scp, X.X)}
	case *xpr.HstStmAcs:
		return HstStmAcs{StmScp: scp.HstStm(x.Txt[X.Trm.Idx:X.Trm.Lim])}
	case *xpr.HstStmsPop:
		return HstStmsPop{X: x.HstStmsAct(scp, X.X)}
	case *xpr.HstStmsDque:
		return HstStmsDque{X: x.HstStmsAct(scp, X.X)}
	case *xpr.HstStmsDel:
		return HstStmsDel{X: x.HstStmsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.HstStmsAt:
		return HstStmsAt{X: x.HstStmsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.HstStmsFst:
		return HstStmsFst{X: x.HstStmsAct(scp, X.X)}
	case *xpr.HstStmsMdl:
		return HstStmsMdl{X: x.HstStmsAct(scp, X.X)}
	case *xpr.HstStmsLst:
		return HstStmsLst{X: x.HstStmsAct(scp, X.X)}
	case *xpr.HstSideFst:
		return HstSideFst{X: x.HstSideAct(scp, X.X)}
	case *xpr.HstSideLst:
		return HstSideLst{X: x.HstSideAct(scp, X.X)}
	case *xpr.HstSideSum:
		return HstSideSum{X: x.HstSideAct(scp, X.X)}
	case *xpr.HstSidePrd:
		return HstSidePrd{X: x.HstSideAct(scp, X.X)}
	case *xpr.HstSideMin:
		return HstSideMin{X: x.HstSideAct(scp, X.X)}
	case *xpr.HstSideMax:
		return HstSideMax{X: x.HstSideAct(scp, X.X)}
	case *xpr.HstSideMid:
		return HstSideMid{X: x.HstSideAct(scp, X.X)}
	case *xpr.HstSideMdn:
		return HstSideMdn{X: x.HstSideAct(scp, X.X)}
	case *xpr.HstSideSma:
		return HstSideSma{X: x.HstSideAct(scp, X.X)}
	case *xpr.HstSideGma:
		return HstSideGma{X: x.HstSideAct(scp, X.X)}
	case *xpr.HstSideWma:
		return HstSideWma{X: x.HstSideAct(scp, X.X)}
	case *xpr.HstSideRsi:
		return HstSideRsi{X: x.HstSideAct(scp, X.X)}
	case *xpr.HstSideWrsi:
		return HstSideWrsi{X: x.HstSideAct(scp, X.X)}
	case *xpr.HstSideAlma:
		return HstSideAlma{X: x.HstSideAct(scp, X.X)}
	case *xpr.HstSideVrnc:
		return HstSideVrnc{X: x.HstSideAct(scp, X.X)}
	case *xpr.HstSideStd:
		return HstSideStd{X: x.HstSideAct(scp, X.X)}
	case *xpr.HstSideRngFul:
		return HstSideRngFul{X: x.HstSideAct(scp, X.X)}
	case *xpr.HstSideRngLst:
		return HstSideRngLst{X: x.HstSideAct(scp, X.X)}
	case *xpr.HstSideProLst:
		return HstSideProLst{X: x.HstSideAct(scp, X.X)}
	case *xpr.HstSideProSma:
		return HstSideProSma{X: x.HstSideAct(scp, X.X)}
	case *xpr.HstSideProAlma:
		return HstSideProAlma{X: x.HstSideAct(scp, X.X)}
	case *xpr.HstSideSar:
		return HstSideSar{X: x.HstSideAct(scp, X.X), I0: x.FltFltAct(scp, X.I0), I1: x.FltFltAct(scp, X.I1)}
	case *xpr.HstSideEma:
		return HstSideEma{X: x.HstSideAct(scp, X.X)}
	case *xpr.HstStmUnaPos:
		return HstStmUnaPos{X: x.HstStmAct(scp, X.X)}
	case *xpr.HstStmUnaNeg:
		return HstStmUnaNeg{X: x.HstStmAct(scp, X.X)}
	case *xpr.HstStmUnaInv:
		return HstStmUnaInv{X: x.HstStmAct(scp, X.X)}
	case *xpr.HstStmUnaSqr:
		return HstStmUnaSqr{X: x.HstStmAct(scp, X.X)}
	case *xpr.HstStmUnaSqrt:
		return HstStmUnaSqrt{X: x.HstStmAct(scp, X.X)}
	case *xpr.HstStmSclAdd:
		return HstStmSclAdd{X: x.HstStmAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.HstStmSclSub:
		return HstStmSclSub{X: x.HstStmAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.HstStmSclMul:
		return HstStmSclMul{X: x.HstStmAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.HstStmSclDiv:
		return HstStmSclDiv{X: x.HstStmAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.HstStmSclRem:
		return HstStmSclRem{X: x.HstStmAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.HstStmSclPow:
		return HstStmSclPow{X: x.HstStmAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.HstStmSclMin:
		return HstStmSclMin{X: x.HstStmAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.HstStmSclMax:
		return HstStmSclMax{X: x.HstStmAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.HstStmSelEql:
		return HstStmSelEql{X: x.HstStmAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.HstStmSelNeq:
		return HstStmSelNeq{X: x.HstStmAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.HstStmSelLss:
		return HstStmSelLss{X: x.HstStmAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.HstStmSelGtr:
		return HstStmSelGtr{X: x.HstStmAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.HstStmSelLeq:
		return HstStmSelLeq{X: x.HstStmAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.HstStmSelGeq:
		return HstStmSelGeq{X: x.HstStmAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.HstStmAggFst:
		return HstStmAggFst{X: x.HstStmAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.HstStmAggLst:
		return HstStmAggLst{X: x.HstStmAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.HstStmAggSum:
		return HstStmAggSum{X: x.HstStmAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.HstStmAggPrd:
		return HstStmAggPrd{X: x.HstStmAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.HstStmAggMin:
		return HstStmAggMin{X: x.HstStmAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.HstStmAggMax:
		return HstStmAggMax{X: x.HstStmAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.HstStmAggMid:
		return HstStmAggMid{X: x.HstStmAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.HstStmAggMdn:
		return HstStmAggMdn{X: x.HstStmAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.HstStmAggSma:
		return HstStmAggSma{X: x.HstStmAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.HstStmAggGma:
		return HstStmAggGma{X: x.HstStmAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.HstStmAggWma:
		return HstStmAggWma{X: x.HstStmAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.HstStmAggRsi:
		return HstStmAggRsi{X: x.HstStmAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.HstStmAggWrsi:
		return HstStmAggWrsi{X: x.HstStmAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.HstStmAggAlma:
		return HstStmAggAlma{X: x.HstStmAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.HstStmAggVrnc:
		return HstStmAggVrnc{X: x.HstStmAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.HstStmAggStd:
		return HstStmAggStd{X: x.HstStmAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.HstStmAggRngFul:
		return HstStmAggRngFul{X: x.HstStmAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.HstStmAggRngLst:
		return HstStmAggRngLst{X: x.HstStmAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.HstStmAggProLst:
		return HstStmAggProLst{X: x.HstStmAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.HstStmAggProSma:
		return HstStmAggProSma{X: x.HstStmAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.HstStmAggProAlma:
		return HstStmAggProAlma{X: x.HstStmAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.HstStmAggEma:
		return HstStmAggEma{X: x.HstStmAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.HstStmInrAdd:
		return HstStmInrAdd{X: x.HstStmAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.HstStmInrSub:
		return HstStmInrSub{X: x.HstStmAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.HstStmInrMul:
		return HstStmInrMul{X: x.HstStmAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.HstStmInrDiv:
		return HstStmInrDiv{X: x.HstStmAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.HstStmInrRem:
		return HstStmInrRem{X: x.HstStmAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.HstStmInrPow:
		return HstStmInrPow{X: x.HstStmAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.HstStmInrMin:
		return HstStmInrMin{X: x.HstStmAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.HstStmInrMax:
		return HstStmInrMax{X: x.HstStmAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.HstStmInrSlp:
		return HstStmInrSlp{X: x.HstStmAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.HstStmOtrAdd:
		return HstStmOtrAdd{X: x.HstStmAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.HstStmAct(scp, X.I1)}
	case *xpr.HstStmOtrSub:
		return HstStmOtrSub{X: x.HstStmAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.HstStmAct(scp, X.I1)}
	case *xpr.HstStmOtrMul:
		return HstStmOtrMul{X: x.HstStmAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.HstStmAct(scp, X.I1)}
	case *xpr.HstStmOtrDiv:
		return HstStmOtrDiv{X: x.HstStmAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.HstStmAct(scp, X.I1)}
	case *xpr.HstStmOtrRem:
		return HstStmOtrRem{X: x.HstStmAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.HstStmAct(scp, X.I1)}
	case *xpr.HstStmOtrPow:
		return HstStmOtrPow{X: x.HstStmAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.HstStmAct(scp, X.I1)}
	case *xpr.HstStmOtrMin:
		return HstStmOtrMin{X: x.HstStmAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.HstStmAct(scp, X.I1)}
	case *xpr.HstStmOtrMax:
		return HstStmOtrMax{X: x.HstStmAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.HstStmAct(scp, X.I1)}
	}
	panic(x.Erf("HstStmAct: no action found %v", reflect.ValueOf(v).Elem().Type().Name()))
}
func (x *Actr) HstCndAct(scp *Scp, v xpr.HstCndXpr) HstCndAct {
	switch X := v.(type) {
	case *xpr.HstCndAsn:
		asnScp := scp
		if X.Depth != 0 {
			for d := X.Depth; asnScp != nil && d != 0; d-- { // recurse up scp by depth
				asnScp = asnScp.Prnt
			}
		}
		return HstCndAsn{CndScp: asnScp.HstCnd(x.Txt[X.Idn.Idx:X.Idn.Lim]), X: x.HstCndAct(scp, X.X)}
	case *xpr.HstCndAcs:
		return HstCndAcs{CndScp: scp.HstCnd(x.Txt[X.Trm.Idx:X.Trm.Lim])}
	case *xpr.HstCndsPop:
		return HstCndsPop{X: x.HstCndsAct(scp, X.X)}
	case *xpr.HstCndsDque:
		return HstCndsDque{X: x.HstCndsAct(scp, X.X)}
	case *xpr.HstCndsDel:
		return HstCndsDel{X: x.HstCndsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.HstCndsAt:
		return HstCndsAt{X: x.HstCndsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.HstCndsFst:
		return HstCndsFst{X: x.HstCndsAct(scp, X.X)}
	case *xpr.HstCndsMdl:
		return HstCndsMdl{X: x.HstCndsAct(scp, X.X)}
	case *xpr.HstCndsLst:
		return HstCndsLst{X: x.HstCndsAct(scp, X.X)}
	case *xpr.HstStmSclEql:
		return HstStmSclEql{X: x.HstStmAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.HstStmSclNeq:
		return HstStmSclNeq{X: x.HstStmAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.HstStmSclLss:
		return HstStmSclLss{X: x.HstStmAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.HstStmSclGtr:
		return HstStmSclGtr{X: x.HstStmAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.HstStmSclLeq:
		return HstStmSclLeq{X: x.HstStmAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.HstStmSclGeq:
		return HstStmSclGeq{X: x.HstStmAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.HstStmInrEql:
		return HstStmInrEql{X: x.HstStmAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.HstStmInrNeq:
		return HstStmInrNeq{X: x.HstStmAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.HstStmInrLss:
		return HstStmInrLss{X: x.HstStmAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.HstStmInrGtr:
		return HstStmInrGtr{X: x.HstStmAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.HstStmInrLeq:
		return HstStmInrLeq{X: x.HstStmAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.HstStmInrGeq:
		return HstStmInrGeq{X: x.HstStmAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.HstStmOtrEql:
		return HstStmOtrEql{X: x.HstStmAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.HstStmAct(scp, X.I1)}
	case *xpr.HstStmOtrNeq:
		return HstStmOtrNeq{X: x.HstStmAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.HstStmAct(scp, X.I1)}
	case *xpr.HstStmOtrLss:
		return HstStmOtrLss{X: x.HstStmAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.HstStmAct(scp, X.I1)}
	case *xpr.HstStmOtrGtr:
		return HstStmOtrGtr{X: x.HstStmAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.HstStmAct(scp, X.I1)}
	case *xpr.HstStmOtrLeq:
		return HstStmOtrLeq{X: x.HstStmAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.HstStmAct(scp, X.I1)}
	case *xpr.HstStmOtrGeq:
		return HstStmOtrGeq{X: x.HstStmAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.HstStmAct(scp, X.I1)}
	case *xpr.HstCndAnd:
		return HstCndAnd{X: x.HstCndAct(scp, X.X), I0: x.HstCndAct(scp, X.I0)}
	case *xpr.HstCndSeq:
		return HstCndSeq{X: x.HstCndAct(scp, X.X), I0: x.TmeTmeAct(scp, X.I0), I1: x.HstCndAct(scp, X.I1)}
	}
	panic(x.Erf("HstCndAct: no action found %v", reflect.ValueOf(v).Elem().Type().Name()))
}
func (x *Actr) HstStgyAct(scp *Scp, v xpr.HstStgyXpr) HstStgyAct {
	switch X := v.(type) {
	case *xpr.HstStgyAsn:
		asnScp := scp
		if X.Depth != 0 {
			for d := X.Depth; asnScp != nil && d != 0; d-- { // recurse up scp by depth
				asnScp = asnScp.Prnt
			}
		}
		return HstStgyAsn{StgyScp: asnScp.HstStgy(x.Txt[X.Idn.Idx:X.Idn.Lim]), X: x.HstStgyAct(scp, X.X)}
	case *xpr.HstStgyAcs:
		return HstStgyAcs{StgyScp: scp.HstStgy(x.Txt[X.Trm.Idx:X.Trm.Lim])}
	case *xpr.HstStgysPop:
		return HstStgysPop{X: x.HstStgysAct(scp, X.X)}
	case *xpr.HstStgysDque:
		return HstStgysDque{X: x.HstStgysAct(scp, X.X)}
	case *xpr.HstStgysDel:
		return HstStgysDel{X: x.HstStgysAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.HstStgysAt:
		return HstStgysAt{X: x.HstStgysAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.HstStgysFst:
		return HstStgysFst{X: x.HstStgysAct(scp, X.X)}
	case *xpr.HstStgysMdl:
		return HstStgysMdl{X: x.HstStgysAct(scp, X.X)}
	case *xpr.HstStgysLst:
		return HstStgysLst{X: x.HstStgysAct(scp, X.X)}
	case *xpr.HstCndStgy:
		var i7 []HstCndAct
		for _, cur := range X.I7 {
			i7 = append(i7, x.HstCndAct(scp, cur))
		}
		return HstCndStgy{X: x.HstCndAct(scp, X.X), I0: x.BolBolAct(scp, X.I0), I1: x.FltFltAct(scp, X.I1), I2: x.FltFltAct(scp, X.I2), I3: x.TmeTmeAct(scp, X.I3), I4: x.FltFltAct(scp, X.I4), I5: x.HstInstrAct(scp, X.I5), I6: x.HstStmsAct(scp, X.I6), I7: i7}
	}
	panic(x.Erf("HstStgyAct: no action found %v", reflect.ValueOf(v).Elem().Type().Name()))
}
func (x *Actr) HstPrvsAct(scp *Scp, v xpr.HstPrvsXpr) HstPrvsAct {
	switch X := v.(type) {
	case *xpr.HstPrvsAsn:
		asnScp := scp
		if X.Depth != 0 {
			for d := X.Depth; asnScp != nil && d != 0; d-- { // recurse up scp by depth
				asnScp = asnScp.Prnt
			}
		}
		return HstPrvsAsn{PrvsScp: asnScp.HstPrvs(x.Txt[X.Idn.Idx:X.Idn.Lim]), X: x.HstPrvsAct(scp, X.X)}
	case *xpr.HstPrvsAcs:
		return HstPrvsAcs{PrvsScp: scp.HstPrvs(x.Txt[X.Trm.Idx:X.Trm.Lim])}
	case *xpr.HstPrvsEach:
		eachScp := NewScp(X.Scp, scp)
		return HstPrvsEach{X: x.HstPrvsAct(scp, X.X), PrvScp: eachScp.HstPrv(x.Txt[X.Idn.Idx:X.Idn.Lim]), Acts: x.Acts(eachScp, X.Xprs...)}
	case *xpr.HstPrvsPllEach:
		return HstPrvsPllEach{X: x.HstPrvsAct(scp, X.X), Txt: x.Txt, Idn: X.Idn, ActScpPrnt: scp, XprScp: X.Scp, Xprs: X.Xprs}
	case *xpr.HstNewPrvs:
		var i0 []HstPrvAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.HstPrvAct(scp, cur))
		}
		return HstNewPrvs{I0: i0}
	case *xpr.HstMakePrvs:
		return HstMakePrvs{I0: x.UntUntAct(scp, X.I0)}
	case *xpr.HstMakeEmpPrvs:
		return HstMakeEmpPrvs{I0: x.UntUntAct(scp, X.I0)}
	case *xpr.HstPrvsCpy:
		return HstPrvsCpy{X: x.HstPrvsAct(scp, X.X)}
	case *xpr.HstPrvsClr:
		return HstPrvsClr{X: x.HstPrvsAct(scp, X.X)}
	case *xpr.HstPrvsRand:
		return HstPrvsRand{X: x.HstPrvsAct(scp, X.X)}
	case *xpr.HstPrvsMrg:
		var i0 []HstPrvsAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.HstPrvsAct(scp, cur))
		}
		return HstPrvsMrg{X: x.HstPrvsAct(scp, X.X), I0: i0}
	case *xpr.HstPrvsPush:
		var i0 []HstPrvAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.HstPrvAct(scp, cur))
		}
		return HstPrvsPush{X: x.HstPrvsAct(scp, X.X), I0: i0}
	case *xpr.HstPrvsQue:
		var i0 []HstPrvAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.HstPrvAct(scp, cur))
		}
		return HstPrvsQue{X: x.HstPrvsAct(scp, X.X), I0: i0}
	case *xpr.HstPrvsIns:
		return HstPrvsIns{X: x.HstPrvsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.HstPrvAct(scp, X.I1)}
	case *xpr.HstPrvsUpd:
		return HstPrvsUpd{X: x.HstPrvsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.HstPrvAct(scp, X.I1)}
	case *xpr.HstPrvsIn:
		return HstPrvsIn{X: x.HstPrvsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.UntUntAct(scp, X.I1)}
	case *xpr.HstPrvsInBnd:
		return HstPrvsInBnd{X: x.HstPrvsAct(scp, X.X), I0: x.BndBndAct(scp, X.I0)}
	case *xpr.HstPrvsFrom:
		return HstPrvsFrom{X: x.HstPrvsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.HstPrvsTo:
		return HstPrvsTo{X: x.HstPrvsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.HstPrvsRev:
		return HstPrvsRev{X: x.HstPrvsAct(scp, X.X)}
	}
	panic(x.Erf("HstPrvsAct: no action found %v", reflect.ValueOf(v).Elem().Type().Name()))
}
func (x *Actr) HstInstrsAct(scp *Scp, v xpr.HstInstrsXpr) HstInstrsAct {
	switch X := v.(type) {
	case *xpr.HstInstrsAsn:
		asnScp := scp
		if X.Depth != 0 {
			for d := X.Depth; asnScp != nil && d != 0; d-- { // recurse up scp by depth
				asnScp = asnScp.Prnt
			}
		}
		return HstInstrsAsn{InstrsScp: asnScp.HstInstrs(x.Txt[X.Idn.Idx:X.Idn.Lim]), X: x.HstInstrsAct(scp, X.X)}
	case *xpr.HstInstrsAcs:
		return HstInstrsAcs{InstrsScp: scp.HstInstrs(x.Txt[X.Trm.Idx:X.Trm.Lim])}
	case *xpr.HstInstrsEach:
		eachScp := NewScp(X.Scp, scp)
		return HstInstrsEach{X: x.HstInstrsAct(scp, X.X), InstrScp: eachScp.HstInstr(x.Txt[X.Idn.Idx:X.Idn.Lim]), Acts: x.Acts(eachScp, X.Xprs...)}
	case *xpr.HstInstrsPllEach:
		return HstInstrsPllEach{X: x.HstInstrsAct(scp, X.X), Txt: x.Txt, Idn: X.Idn, ActScpPrnt: scp, XprScp: X.Scp, Xprs: X.Xprs}
	case *xpr.HstNewInstrs:
		var i0 []HstInstrAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.HstInstrAct(scp, cur))
		}
		return HstNewInstrs{I0: i0}
	case *xpr.HstMakeInstrs:
		return HstMakeInstrs{I0: x.UntUntAct(scp, X.I0)}
	case *xpr.HstMakeEmpInstrs:
		return HstMakeEmpInstrs{I0: x.UntUntAct(scp, X.I0)}
	case *xpr.HstInstrsCpy:
		return HstInstrsCpy{X: x.HstInstrsAct(scp, X.X)}
	case *xpr.HstInstrsClr:
		return HstInstrsClr{X: x.HstInstrsAct(scp, X.X)}
	case *xpr.HstInstrsRand:
		return HstInstrsRand{X: x.HstInstrsAct(scp, X.X)}
	case *xpr.HstInstrsMrg:
		var i0 []HstInstrsAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.HstInstrsAct(scp, cur))
		}
		return HstInstrsMrg{X: x.HstInstrsAct(scp, X.X), I0: i0}
	case *xpr.HstInstrsPush:
		var i0 []HstInstrAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.HstInstrAct(scp, cur))
		}
		return HstInstrsPush{X: x.HstInstrsAct(scp, X.X), I0: i0}
	case *xpr.HstInstrsQue:
		var i0 []HstInstrAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.HstInstrAct(scp, cur))
		}
		return HstInstrsQue{X: x.HstInstrsAct(scp, X.X), I0: i0}
	case *xpr.HstInstrsIns:
		return HstInstrsIns{X: x.HstInstrsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.HstInstrAct(scp, X.I1)}
	case *xpr.HstInstrsUpd:
		return HstInstrsUpd{X: x.HstInstrsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.HstInstrAct(scp, X.I1)}
	case *xpr.HstInstrsIn:
		return HstInstrsIn{X: x.HstInstrsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.UntUntAct(scp, X.I1)}
	case *xpr.HstInstrsInBnd:
		return HstInstrsInBnd{X: x.HstInstrsAct(scp, X.X), I0: x.BndBndAct(scp, X.I0)}
	case *xpr.HstInstrsFrom:
		return HstInstrsFrom{X: x.HstInstrsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.HstInstrsTo:
		return HstInstrsTo{X: x.HstInstrsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.HstInstrsRev:
		return HstInstrsRev{X: x.HstInstrsAct(scp, X.X)}
	}
	panic(x.Erf("HstInstrsAct: no action found %v", reflect.ValueOf(v).Elem().Type().Name()))
}
func (x *Actr) HstInrvlsAct(scp *Scp, v xpr.HstInrvlsXpr) HstInrvlsAct {
	switch X := v.(type) {
	case *xpr.HstInrvlsAsn:
		asnScp := scp
		if X.Depth != 0 {
			for d := X.Depth; asnScp != nil && d != 0; d-- { // recurse up scp by depth
				asnScp = asnScp.Prnt
			}
		}
		return HstInrvlsAsn{InrvlsScp: asnScp.HstInrvls(x.Txt[X.Idn.Idx:X.Idn.Lim]), X: x.HstInrvlsAct(scp, X.X)}
	case *xpr.HstInrvlsAcs:
		return HstInrvlsAcs{InrvlsScp: scp.HstInrvls(x.Txt[X.Trm.Idx:X.Trm.Lim])}
	case *xpr.HstInrvlsEach:
		eachScp := NewScp(X.Scp, scp)
		return HstInrvlsEach{X: x.HstInrvlsAct(scp, X.X), InrvlScp: eachScp.HstInrvl(x.Txt[X.Idn.Idx:X.Idn.Lim]), Acts: x.Acts(eachScp, X.Xprs...)}
	case *xpr.HstInrvlsPllEach:
		return HstInrvlsPllEach{X: x.HstInrvlsAct(scp, X.X), Txt: x.Txt, Idn: X.Idn, ActScpPrnt: scp, XprScp: X.Scp, Xprs: X.Xprs}
	case *xpr.HstNewInrvls:
		var i0 []HstInrvlAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.HstInrvlAct(scp, cur))
		}
		return HstNewInrvls{I0: i0}
	case *xpr.HstMakeInrvls:
		return HstMakeInrvls{I0: x.UntUntAct(scp, X.I0)}
	case *xpr.HstMakeEmpInrvls:
		return HstMakeEmpInrvls{I0: x.UntUntAct(scp, X.I0)}
	case *xpr.HstInrvlsCpy:
		return HstInrvlsCpy{X: x.HstInrvlsAct(scp, X.X)}
	case *xpr.HstInrvlsClr:
		return HstInrvlsClr{X: x.HstInrvlsAct(scp, X.X)}
	case *xpr.HstInrvlsRand:
		return HstInrvlsRand{X: x.HstInrvlsAct(scp, X.X)}
	case *xpr.HstInrvlsMrg:
		var i0 []HstInrvlsAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.HstInrvlsAct(scp, cur))
		}
		return HstInrvlsMrg{X: x.HstInrvlsAct(scp, X.X), I0: i0}
	case *xpr.HstInrvlsPush:
		var i0 []HstInrvlAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.HstInrvlAct(scp, cur))
		}
		return HstInrvlsPush{X: x.HstInrvlsAct(scp, X.X), I0: i0}
	case *xpr.HstInrvlsQue:
		var i0 []HstInrvlAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.HstInrvlAct(scp, cur))
		}
		return HstInrvlsQue{X: x.HstInrvlsAct(scp, X.X), I0: i0}
	case *xpr.HstInrvlsIns:
		return HstInrvlsIns{X: x.HstInrvlsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.HstInrvlAct(scp, X.I1)}
	case *xpr.HstInrvlsUpd:
		return HstInrvlsUpd{X: x.HstInrvlsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.HstInrvlAct(scp, X.I1)}
	case *xpr.HstInrvlsIn:
		return HstInrvlsIn{X: x.HstInrvlsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.UntUntAct(scp, X.I1)}
	case *xpr.HstInrvlsInBnd:
		return HstInrvlsInBnd{X: x.HstInrvlsAct(scp, X.X), I0: x.BndBndAct(scp, X.I0)}
	case *xpr.HstInrvlsFrom:
		return HstInrvlsFrom{X: x.HstInrvlsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.HstInrvlsTo:
		return HstInrvlsTo{X: x.HstInrvlsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.HstInrvlsRev:
		return HstInrvlsRev{X: x.HstInrvlsAct(scp, X.X)}
	}
	panic(x.Erf("HstInrvlsAct: no action found %v", reflect.ValueOf(v).Elem().Type().Name()))
}
func (x *Actr) HstSidesAct(scp *Scp, v xpr.HstSidesXpr) HstSidesAct {
	switch X := v.(type) {
	case *xpr.HstSidesAsn:
		asnScp := scp
		if X.Depth != 0 {
			for d := X.Depth; asnScp != nil && d != 0; d-- { // recurse up scp by depth
				asnScp = asnScp.Prnt
			}
		}
		return HstSidesAsn{SidesScp: asnScp.HstSides(x.Txt[X.Idn.Idx:X.Idn.Lim]), X: x.HstSidesAct(scp, X.X)}
	case *xpr.HstSidesAcs:
		return HstSidesAcs{SidesScp: scp.HstSides(x.Txt[X.Trm.Idx:X.Trm.Lim])}
	case *xpr.HstSidesEach:
		eachScp := NewScp(X.Scp, scp)
		return HstSidesEach{X: x.HstSidesAct(scp, X.X), SideScp: eachScp.HstSide(x.Txt[X.Idn.Idx:X.Idn.Lim]), Acts: x.Acts(eachScp, X.Xprs...)}
	case *xpr.HstSidesPllEach:
		return HstSidesPllEach{X: x.HstSidesAct(scp, X.X), Txt: x.Txt, Idn: X.Idn, ActScpPrnt: scp, XprScp: X.Scp, Xprs: X.Xprs}
	case *xpr.HstNewSides:
		var i0 []HstSideAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.HstSideAct(scp, cur))
		}
		return HstNewSides{I0: i0}
	case *xpr.HstMakeSides:
		return HstMakeSides{I0: x.UntUntAct(scp, X.I0)}
	case *xpr.HstMakeEmpSides:
		return HstMakeEmpSides{I0: x.UntUntAct(scp, X.I0)}
	case *xpr.HstSidesCpy:
		return HstSidesCpy{X: x.HstSidesAct(scp, X.X)}
	case *xpr.HstSidesClr:
		return HstSidesClr{X: x.HstSidesAct(scp, X.X)}
	case *xpr.HstSidesRand:
		return HstSidesRand{X: x.HstSidesAct(scp, X.X)}
	case *xpr.HstSidesMrg:
		var i0 []HstSidesAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.HstSidesAct(scp, cur))
		}
		return HstSidesMrg{X: x.HstSidesAct(scp, X.X), I0: i0}
	case *xpr.HstSidesPush:
		var i0 []HstSideAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.HstSideAct(scp, cur))
		}
		return HstSidesPush{X: x.HstSidesAct(scp, X.X), I0: i0}
	case *xpr.HstSidesQue:
		var i0 []HstSideAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.HstSideAct(scp, cur))
		}
		return HstSidesQue{X: x.HstSidesAct(scp, X.X), I0: i0}
	case *xpr.HstSidesIns:
		return HstSidesIns{X: x.HstSidesAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.HstSideAct(scp, X.I1)}
	case *xpr.HstSidesUpd:
		return HstSidesUpd{X: x.HstSidesAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.HstSideAct(scp, X.I1)}
	case *xpr.HstSidesIn:
		return HstSidesIn{X: x.HstSidesAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.UntUntAct(scp, X.I1)}
	case *xpr.HstSidesInBnd:
		return HstSidesInBnd{X: x.HstSidesAct(scp, X.X), I0: x.BndBndAct(scp, X.I0)}
	case *xpr.HstSidesFrom:
		return HstSidesFrom{X: x.HstSidesAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.HstSidesTo:
		return HstSidesTo{X: x.HstSidesAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.HstSidesRev:
		return HstSidesRev{X: x.HstSidesAct(scp, X.X)}
	}
	panic(x.Erf("HstSidesAct: no action found %v", reflect.ValueOf(v).Elem().Type().Name()))
}
func (x *Actr) HstStmsAct(scp *Scp, v xpr.HstStmsXpr) HstStmsAct {
	switch X := v.(type) {
	case *xpr.HstStmsAsn:
		asnScp := scp
		if X.Depth != 0 {
			for d := X.Depth; asnScp != nil && d != 0; d-- { // recurse up scp by depth
				asnScp = asnScp.Prnt
			}
		}
		return HstStmsAsn{StmsScp: asnScp.HstStms(x.Txt[X.Idn.Idx:X.Idn.Lim]), X: x.HstStmsAct(scp, X.X)}
	case *xpr.HstStmsAcs:
		return HstStmsAcs{StmsScp: scp.HstStms(x.Txt[X.Trm.Idx:X.Trm.Lim])}
	case *xpr.HstStmsEach:
		eachScp := NewScp(X.Scp, scp)
		return HstStmsEach{X: x.HstStmsAct(scp, X.X), StmScp: eachScp.HstStm(x.Txt[X.Idn.Idx:X.Idn.Lim]), Acts: x.Acts(eachScp, X.Xprs...)}
	case *xpr.HstStmsPllEach:
		return HstStmsPllEach{X: x.HstStmsAct(scp, X.X), Txt: x.Txt, Idn: X.Idn, ActScpPrnt: scp, XprScp: X.Scp, Xprs: X.Xprs}
	case *xpr.HstNewStms:
		var i0 []HstStmAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.HstStmAct(scp, cur))
		}
		return HstNewStms{I0: i0}
	case *xpr.HstMakeStms:
		return HstMakeStms{I0: x.UntUntAct(scp, X.I0)}
	case *xpr.HstMakeEmpStms:
		return HstMakeEmpStms{I0: x.UntUntAct(scp, X.I0)}
	case *xpr.HstStmsCpy:
		return HstStmsCpy{X: x.HstStmsAct(scp, X.X)}
	case *xpr.HstStmsClr:
		return HstStmsClr{X: x.HstStmsAct(scp, X.X)}
	case *xpr.HstStmsRand:
		return HstStmsRand{X: x.HstStmsAct(scp, X.X)}
	case *xpr.HstStmsMrg:
		var i0 []HstStmsAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.HstStmsAct(scp, cur))
		}
		return HstStmsMrg{X: x.HstStmsAct(scp, X.X), I0: i0}
	case *xpr.HstStmsPush:
		var i0 []HstStmAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.HstStmAct(scp, cur))
		}
		return HstStmsPush{X: x.HstStmsAct(scp, X.X), I0: i0}
	case *xpr.HstStmsQue:
		var i0 []HstStmAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.HstStmAct(scp, cur))
		}
		return HstStmsQue{X: x.HstStmsAct(scp, X.X), I0: i0}
	case *xpr.HstStmsIns:
		return HstStmsIns{X: x.HstStmsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.HstStmAct(scp, X.I1)}
	case *xpr.HstStmsUpd:
		return HstStmsUpd{X: x.HstStmsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.HstStmAct(scp, X.I1)}
	case *xpr.HstStmsIn:
		return HstStmsIn{X: x.HstStmsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.UntUntAct(scp, X.I1)}
	case *xpr.HstStmsInBnd:
		return HstStmsInBnd{X: x.HstStmsAct(scp, X.X), I0: x.BndBndAct(scp, X.I0)}
	case *xpr.HstStmsFrom:
		return HstStmsFrom{X: x.HstStmsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.HstStmsTo:
		return HstStmsTo{X: x.HstStmsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.HstStmsRev:
		return HstStmsRev{X: x.HstStmsAct(scp, X.X)}
	}
	panic(x.Erf("HstStmsAct: no action found %v", reflect.ValueOf(v).Elem().Type().Name()))
}
func (x *Actr) HstCndsAct(scp *Scp, v xpr.HstCndsXpr) HstCndsAct {
	switch X := v.(type) {
	case *xpr.HstCndsAsn:
		asnScp := scp
		if X.Depth != 0 {
			for d := X.Depth; asnScp != nil && d != 0; d-- { // recurse up scp by depth
				asnScp = asnScp.Prnt
			}
		}
		return HstCndsAsn{CndsScp: asnScp.HstCnds(x.Txt[X.Idn.Idx:X.Idn.Lim]), X: x.HstCndsAct(scp, X.X)}
	case *xpr.HstCndsAcs:
		return HstCndsAcs{CndsScp: scp.HstCnds(x.Txt[X.Trm.Idx:X.Trm.Lim])}
	case *xpr.HstCndsEach:
		eachScp := NewScp(X.Scp, scp)
		return HstCndsEach{X: x.HstCndsAct(scp, X.X), CndScp: eachScp.HstCnd(x.Txt[X.Idn.Idx:X.Idn.Lim]), Acts: x.Acts(eachScp, X.Xprs...)}
	case *xpr.HstCndsPllEach:
		return HstCndsPllEach{X: x.HstCndsAct(scp, X.X), Txt: x.Txt, Idn: X.Idn, ActScpPrnt: scp, XprScp: X.Scp, Xprs: X.Xprs}
	case *xpr.HstNewCnds:
		var i0 []HstCndAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.HstCndAct(scp, cur))
		}
		return HstNewCnds{I0: i0}
	case *xpr.HstMakeCnds:
		return HstMakeCnds{I0: x.UntUntAct(scp, X.I0)}
	case *xpr.HstMakeEmpCnds:
		return HstMakeEmpCnds{I0: x.UntUntAct(scp, X.I0)}
	case *xpr.HstCndsCpy:
		return HstCndsCpy{X: x.HstCndsAct(scp, X.X)}
	case *xpr.HstCndsClr:
		return HstCndsClr{X: x.HstCndsAct(scp, X.X)}
	case *xpr.HstCndsRand:
		return HstCndsRand{X: x.HstCndsAct(scp, X.X)}
	case *xpr.HstCndsMrg:
		var i0 []HstCndsAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.HstCndsAct(scp, cur))
		}
		return HstCndsMrg{X: x.HstCndsAct(scp, X.X), I0: i0}
	case *xpr.HstCndsPush:
		var i0 []HstCndAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.HstCndAct(scp, cur))
		}
		return HstCndsPush{X: x.HstCndsAct(scp, X.X), I0: i0}
	case *xpr.HstCndsQue:
		var i0 []HstCndAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.HstCndAct(scp, cur))
		}
		return HstCndsQue{X: x.HstCndsAct(scp, X.X), I0: i0}
	case *xpr.HstCndsIns:
		return HstCndsIns{X: x.HstCndsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.HstCndAct(scp, X.I1)}
	case *xpr.HstCndsUpd:
		return HstCndsUpd{X: x.HstCndsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.HstCndAct(scp, X.I1)}
	case *xpr.HstCndsIn:
		return HstCndsIn{X: x.HstCndsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.UntUntAct(scp, X.I1)}
	case *xpr.HstCndsInBnd:
		return HstCndsInBnd{X: x.HstCndsAct(scp, X.X), I0: x.BndBndAct(scp, X.I0)}
	case *xpr.HstCndsFrom:
		return HstCndsFrom{X: x.HstCndsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.HstCndsTo:
		return HstCndsTo{X: x.HstCndsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.HstCndsRev:
		return HstCndsRev{X: x.HstCndsAct(scp, X.X)}
	}
	panic(x.Erf("HstCndsAct: no action found %v", reflect.ValueOf(v).Elem().Type().Name()))
}
func (x *Actr) HstStgysAct(scp *Scp, v xpr.HstStgysXpr) HstStgysAct {
	switch X := v.(type) {
	case *xpr.HstStgysAsn:
		asnScp := scp
		if X.Depth != 0 {
			for d := X.Depth; asnScp != nil && d != 0; d-- { // recurse up scp by depth
				asnScp = asnScp.Prnt
			}
		}
		return HstStgysAsn{StgysScp: asnScp.HstStgys(x.Txt[X.Idn.Idx:X.Idn.Lim]), X: x.HstStgysAct(scp, X.X)}
	case *xpr.HstStgysAcs:
		return HstStgysAcs{StgysScp: scp.HstStgys(x.Txt[X.Trm.Idx:X.Trm.Lim])}
	case *xpr.HstStgysEach:
		eachScp := NewScp(X.Scp, scp)
		return HstStgysEach{X: x.HstStgysAct(scp, X.X), StgyScp: eachScp.HstStgy(x.Txt[X.Idn.Idx:X.Idn.Lim]), Acts: x.Acts(eachScp, X.Xprs...)}
	case *xpr.HstStgysPllEach:
		return HstStgysPllEach{X: x.HstStgysAct(scp, X.X), Txt: x.Txt, Idn: X.Idn, ActScpPrnt: scp, XprScp: X.Scp, Xprs: X.Xprs}
	case *xpr.HstNewStgys:
		var i0 []HstStgyAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.HstStgyAct(scp, cur))
		}
		return HstNewStgys{I0: i0}
	case *xpr.HstMakeStgys:
		return HstMakeStgys{I0: x.UntUntAct(scp, X.I0)}
	case *xpr.HstMakeEmpStgys:
		return HstMakeEmpStgys{I0: x.UntUntAct(scp, X.I0)}
	case *xpr.HstStgysCpy:
		return HstStgysCpy{X: x.HstStgysAct(scp, X.X)}
	case *xpr.HstStgysClr:
		return HstStgysClr{X: x.HstStgysAct(scp, X.X)}
	case *xpr.HstStgysRand:
		return HstStgysRand{X: x.HstStgysAct(scp, X.X)}
	case *xpr.HstStgysMrg:
		var i0 []HstStgysAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.HstStgysAct(scp, cur))
		}
		return HstStgysMrg{X: x.HstStgysAct(scp, X.X), I0: i0}
	case *xpr.HstStgysPush:
		var i0 []HstStgyAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.HstStgyAct(scp, cur))
		}
		return HstStgysPush{X: x.HstStgysAct(scp, X.X), I0: i0}
	case *xpr.HstStgysQue:
		var i0 []HstStgyAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.HstStgyAct(scp, cur))
		}
		return HstStgysQue{X: x.HstStgysAct(scp, X.X), I0: i0}
	case *xpr.HstStgysIns:
		return HstStgysIns{X: x.HstStgysAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.HstStgyAct(scp, X.I1)}
	case *xpr.HstStgysUpd:
		return HstStgysUpd{X: x.HstStgysAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.HstStgyAct(scp, X.I1)}
	case *xpr.HstStgysIn:
		return HstStgysIn{X: x.HstStgysAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.UntUntAct(scp, X.I1)}
	case *xpr.HstStgysInBnd:
		return HstStgysInBnd{X: x.HstStgysAct(scp, X.X), I0: x.BndBndAct(scp, X.I0)}
	case *xpr.HstStgysFrom:
		return HstStgysFrom{X: x.HstStgysAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.HstStgysTo:
		return HstStgysTo{X: x.HstStgysAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.HstStgysRev:
		return HstStgysRev{X: x.HstStgysAct(scp, X.X)}
	}
	panic(x.Erf("HstStgysAct: no action found %v", reflect.ValueOf(v).Elem().Type().Name()))
}
func (x *Actr) RltPrvAct(scp *Scp, v xpr.RltPrvXpr) RltPrvAct {
	switch X := v.(type) {
	case *xpr.RltPrvAsn:
		asnScp := scp
		if X.Depth != 0 {
			for d := X.Depth; asnScp != nil && d != 0; d-- { // recurse up scp by depth
				asnScp = asnScp.Prnt
			}
		}
		return RltPrvAsn{PrvScp: asnScp.RltPrv(x.Txt[X.Idn.Idx:X.Idn.Lim]), X: x.RltPrvAct(scp, X.X)}
	case *xpr.RltPrvAcs:
		return RltPrvAcs{PrvScp: scp.RltPrv(x.Txt[X.Trm.Idx:X.Trm.Lim])}
	case *xpr.RltOan:
		return RltOan{}
	case *xpr.RltPrvsPop:
		return RltPrvsPop{X: x.RltPrvsAct(scp, X.X)}
	case *xpr.RltPrvsDque:
		return RltPrvsDque{X: x.RltPrvsAct(scp, X.X)}
	case *xpr.RltPrvsDel:
		return RltPrvsDel{X: x.RltPrvsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.RltPrvsAt:
		return RltPrvsAt{X: x.RltPrvsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.RltPrvsFst:
		return RltPrvsFst{X: x.RltPrvsAct(scp, X.X)}
	case *xpr.RltPrvsMdl:
		return RltPrvsMdl{X: x.RltPrvsAct(scp, X.X)}
	case *xpr.RltPrvsLst:
		return RltPrvsLst{X: x.RltPrvsAct(scp, X.X)}
	}
	panic(x.Erf("RltPrvAct: no action found %v", reflect.ValueOf(v).Elem().Type().Name()))
}
func (x *Actr) RltInstrAct(scp *Scp, v xpr.RltInstrXpr) RltInstrAct {
	switch X := v.(type) {
	case *xpr.RltInstrAsn:
		asnScp := scp
		if X.Depth != 0 {
			for d := X.Depth; asnScp != nil && d != 0; d-- { // recurse up scp by depth
				asnScp = asnScp.Prnt
			}
		}
		return RltInstrAsn{InstrScp: asnScp.RltInstr(x.Txt[X.Idn.Idx:X.Idn.Lim]), X: x.RltInstrAct(scp, X.X)}
	case *xpr.RltInstrAcs:
		return RltInstrAcs{InstrScp: scp.RltInstr(x.Txt[X.Trm.Idx:X.Trm.Lim])}
	case *xpr.RltInstrsPop:
		return RltInstrsPop{X: x.RltInstrsAct(scp, X.X)}
	case *xpr.RltInstrsDque:
		return RltInstrsDque{X: x.RltInstrsAct(scp, X.X)}
	case *xpr.RltInstrsDel:
		return RltInstrsDel{X: x.RltInstrsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.RltInstrsAt:
		return RltInstrsAt{X: x.RltInstrsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.RltInstrsFst:
		return RltInstrsFst{X: x.RltInstrsAct(scp, X.X)}
	case *xpr.RltInstrsMdl:
		return RltInstrsMdl{X: x.RltInstrsAct(scp, X.X)}
	case *xpr.RltInstrsLst:
		return RltInstrsLst{X: x.RltInstrsAct(scp, X.X)}
	case *xpr.RltPrvEurUsd:
		var i0 []TmeRngAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.TmeRngAct(scp, cur))
		}
		return RltPrvEurUsd{X: x.RltPrvAct(scp, X.X), I0: i0}
	case *xpr.RltPrvAudUsd:
		var i0 []TmeRngAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.TmeRngAct(scp, cur))
		}
		return RltPrvAudUsd{X: x.RltPrvAct(scp, X.X), I0: i0}
	case *xpr.RltPrvNzdUsd:
		var i0 []TmeRngAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.TmeRngAct(scp, cur))
		}
		return RltPrvNzdUsd{X: x.RltPrvAct(scp, X.X), I0: i0}
	case *xpr.RltPrvGbpUsd:
		var i0 []TmeRngAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.TmeRngAct(scp, cur))
		}
		return RltPrvGbpUsd{X: x.RltPrvAct(scp, X.X), I0: i0}
	}
	panic(x.Erf("RltInstrAct: no action found %v", reflect.ValueOf(v).Elem().Type().Name()))
}
func (x *Actr) RltInrvlAct(scp *Scp, v xpr.RltInrvlXpr) RltInrvlAct {
	switch X := v.(type) {
	case *xpr.RltInrvlAsn:
		asnScp := scp
		if X.Depth != 0 {
			for d := X.Depth; asnScp != nil && d != 0; d-- { // recurse up scp by depth
				asnScp = asnScp.Prnt
			}
		}
		return RltInrvlAsn{InrvlScp: asnScp.RltInrvl(x.Txt[X.Idn.Idx:X.Idn.Lim]), X: x.RltInrvlAct(scp, X.X)}
	case *xpr.RltInrvlAcs:
		return RltInrvlAcs{InrvlScp: scp.RltInrvl(x.Txt[X.Trm.Idx:X.Trm.Lim])}
	case *xpr.RltInrvlsPop:
		return RltInrvlsPop{X: x.RltInrvlsAct(scp, X.X)}
	case *xpr.RltInrvlsDque:
		return RltInrvlsDque{X: x.RltInrvlsAct(scp, X.X)}
	case *xpr.RltInrvlsDel:
		return RltInrvlsDel{X: x.RltInrvlsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.RltInrvlsAt:
		return RltInrvlsAt{X: x.RltInrvlsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.RltInrvlsFst:
		return RltInrvlsFst{X: x.RltInrvlsAct(scp, X.X)}
	case *xpr.RltInrvlsMdl:
		return RltInrvlsMdl{X: x.RltInrvlsAct(scp, X.X)}
	case *xpr.RltInrvlsLst:
		return RltInrvlsLst{X: x.RltInrvlsAct(scp, X.X)}
	case *xpr.RltInstrI:
		return RltInstrI{X: x.RltInstrAct(scp, X.X), I0: x.TmeTmeAct(scp, X.I0)}
	}
	panic(x.Erf("RltInrvlAct: no action found %v", reflect.ValueOf(v).Elem().Type().Name()))
}
func (x *Actr) RltSideAct(scp *Scp, v xpr.RltSideXpr) RltSideAct {
	switch X := v.(type) {
	case *xpr.RltSideAsn:
		asnScp := scp
		if X.Depth != 0 {
			for d := X.Depth; asnScp != nil && d != 0; d-- { // recurse up scp by depth
				asnScp = asnScp.Prnt
			}
		}
		return RltSideAsn{SideScp: asnScp.RltSide(x.Txt[X.Idn.Idx:X.Idn.Lim]), X: x.RltSideAct(scp, X.X)}
	case *xpr.RltSideAcs:
		return RltSideAcs{SideScp: scp.RltSide(x.Txt[X.Trm.Idx:X.Trm.Lim])}
	case *xpr.RltSidesPop:
		return RltSidesPop{X: x.RltSidesAct(scp, X.X)}
	case *xpr.RltSidesDque:
		return RltSidesDque{X: x.RltSidesAct(scp, X.X)}
	case *xpr.RltSidesDel:
		return RltSidesDel{X: x.RltSidesAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.RltSidesAt:
		return RltSidesAt{X: x.RltSidesAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.RltSidesFst:
		return RltSidesFst{X: x.RltSidesAct(scp, X.X)}
	case *xpr.RltSidesMdl:
		return RltSidesMdl{X: x.RltSidesAct(scp, X.X)}
	case *xpr.RltSidesLst:
		return RltSidesLst{X: x.RltSidesAct(scp, X.X)}
	case *xpr.RltInrvlBid:
		return RltInrvlBid{X: x.RltInrvlAct(scp, X.X)}
	case *xpr.RltInrvlAsk:
		return RltInrvlAsk{X: x.RltInrvlAct(scp, X.X)}
	}
	panic(x.Erf("RltSideAct: no action found %v", reflect.ValueOf(v).Elem().Type().Name()))
}
func (x *Actr) RltStmAct(scp *Scp, v xpr.RltStmXpr) RltStmAct {
	switch X := v.(type) {
	case *xpr.RltStmAsn:
		asnScp := scp
		if X.Depth != 0 {
			for d := X.Depth; asnScp != nil && d != 0; d-- { // recurse up scp by depth
				asnScp = asnScp.Prnt
			}
		}
		return RltStmAsn{StmScp: asnScp.RltStm(x.Txt[X.Idn.Idx:X.Idn.Lim]), X: x.RltStmAct(scp, X.X)}
	case *xpr.RltStmAcs:
		return RltStmAcs{StmScp: scp.RltStm(x.Txt[X.Trm.Idx:X.Trm.Lim])}
	case *xpr.RltStmsPop:
		return RltStmsPop{X: x.RltStmsAct(scp, X.X)}
	case *xpr.RltStmsDque:
		return RltStmsDque{X: x.RltStmsAct(scp, X.X)}
	case *xpr.RltStmsDel:
		return RltStmsDel{X: x.RltStmsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.RltStmsAt:
		return RltStmsAt{X: x.RltStmsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.RltStmsFst:
		return RltStmsFst{X: x.RltStmsAct(scp, X.X)}
	case *xpr.RltStmsMdl:
		return RltStmsMdl{X: x.RltStmsAct(scp, X.X)}
	case *xpr.RltStmsLst:
		return RltStmsLst{X: x.RltStmsAct(scp, X.X)}
	case *xpr.RltSideFst:
		return RltSideFst{X: x.RltSideAct(scp, X.X)}
	case *xpr.RltSideLst:
		return RltSideLst{X: x.RltSideAct(scp, X.X)}
	case *xpr.RltSideSum:
		return RltSideSum{X: x.RltSideAct(scp, X.X)}
	case *xpr.RltSidePrd:
		return RltSidePrd{X: x.RltSideAct(scp, X.X)}
	case *xpr.RltSideMin:
		return RltSideMin{X: x.RltSideAct(scp, X.X)}
	case *xpr.RltSideMax:
		return RltSideMax{X: x.RltSideAct(scp, X.X)}
	case *xpr.RltSideMid:
		return RltSideMid{X: x.RltSideAct(scp, X.X)}
	case *xpr.RltSideMdn:
		return RltSideMdn{X: x.RltSideAct(scp, X.X)}
	case *xpr.RltSideSma:
		return RltSideSma{X: x.RltSideAct(scp, X.X)}
	case *xpr.RltSideGma:
		return RltSideGma{X: x.RltSideAct(scp, X.X)}
	case *xpr.RltSideWma:
		return RltSideWma{X: x.RltSideAct(scp, X.X)}
	case *xpr.RltSideRsi:
		return RltSideRsi{X: x.RltSideAct(scp, X.X)}
	case *xpr.RltSideWrsi:
		return RltSideWrsi{X: x.RltSideAct(scp, X.X)}
	case *xpr.RltSideAlma:
		return RltSideAlma{X: x.RltSideAct(scp, X.X)}
	case *xpr.RltSideVrnc:
		return RltSideVrnc{X: x.RltSideAct(scp, X.X)}
	case *xpr.RltSideStd:
		return RltSideStd{X: x.RltSideAct(scp, X.X)}
	case *xpr.RltSideRngFul:
		return RltSideRngFul{X: x.RltSideAct(scp, X.X)}
	case *xpr.RltSideRngLst:
		return RltSideRngLst{X: x.RltSideAct(scp, X.X)}
	case *xpr.RltSideProLst:
		return RltSideProLst{X: x.RltSideAct(scp, X.X)}
	case *xpr.RltSideProSma:
		return RltSideProSma{X: x.RltSideAct(scp, X.X)}
	case *xpr.RltSideProAlma:
		return RltSideProAlma{X: x.RltSideAct(scp, X.X)}
	case *xpr.RltSideSar:
		return RltSideSar{X: x.RltSideAct(scp, X.X), I0: x.FltFltAct(scp, X.I0), I1: x.FltFltAct(scp, X.I1)}
	case *xpr.RltSideEma:
		return RltSideEma{X: x.RltSideAct(scp, X.X)}
	case *xpr.RltStmUnaPos:
		return RltStmUnaPos{X: x.RltStmAct(scp, X.X)}
	case *xpr.RltStmUnaNeg:
		return RltStmUnaNeg{X: x.RltStmAct(scp, X.X)}
	case *xpr.RltStmUnaInv:
		return RltStmUnaInv{X: x.RltStmAct(scp, X.X)}
	case *xpr.RltStmUnaSqr:
		return RltStmUnaSqr{X: x.RltStmAct(scp, X.X)}
	case *xpr.RltStmUnaSqrt:
		return RltStmUnaSqrt{X: x.RltStmAct(scp, X.X)}
	case *xpr.RltStmSclAdd:
		return RltStmSclAdd{X: x.RltStmAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.RltStmSclSub:
		return RltStmSclSub{X: x.RltStmAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.RltStmSclMul:
		return RltStmSclMul{X: x.RltStmAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.RltStmSclDiv:
		return RltStmSclDiv{X: x.RltStmAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.RltStmSclRem:
		return RltStmSclRem{X: x.RltStmAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.RltStmSclPow:
		return RltStmSclPow{X: x.RltStmAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.RltStmSclMin:
		return RltStmSclMin{X: x.RltStmAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.RltStmSclMax:
		return RltStmSclMax{X: x.RltStmAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.RltStmSelEql:
		return RltStmSelEql{X: x.RltStmAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.RltStmSelNeq:
		return RltStmSelNeq{X: x.RltStmAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.RltStmSelLss:
		return RltStmSelLss{X: x.RltStmAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.RltStmSelGtr:
		return RltStmSelGtr{X: x.RltStmAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.RltStmSelLeq:
		return RltStmSelLeq{X: x.RltStmAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.RltStmSelGeq:
		return RltStmSelGeq{X: x.RltStmAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.RltStmAggFst:
		return RltStmAggFst{X: x.RltStmAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.RltStmAggLst:
		return RltStmAggLst{X: x.RltStmAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.RltStmAggSum:
		return RltStmAggSum{X: x.RltStmAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.RltStmAggPrd:
		return RltStmAggPrd{X: x.RltStmAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.RltStmAggMin:
		return RltStmAggMin{X: x.RltStmAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.RltStmAggMax:
		return RltStmAggMax{X: x.RltStmAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.RltStmAggMid:
		return RltStmAggMid{X: x.RltStmAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.RltStmAggMdn:
		return RltStmAggMdn{X: x.RltStmAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.RltStmAggSma:
		return RltStmAggSma{X: x.RltStmAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.RltStmAggGma:
		return RltStmAggGma{X: x.RltStmAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.RltStmAggWma:
		return RltStmAggWma{X: x.RltStmAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.RltStmAggRsi:
		return RltStmAggRsi{X: x.RltStmAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.RltStmAggWrsi:
		return RltStmAggWrsi{X: x.RltStmAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.RltStmAggAlma:
		return RltStmAggAlma{X: x.RltStmAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.RltStmAggVrnc:
		return RltStmAggVrnc{X: x.RltStmAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.RltStmAggStd:
		return RltStmAggStd{X: x.RltStmAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.RltStmAggRngFul:
		return RltStmAggRngFul{X: x.RltStmAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.RltStmAggRngLst:
		return RltStmAggRngLst{X: x.RltStmAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.RltStmAggProLst:
		return RltStmAggProLst{X: x.RltStmAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.RltStmAggProSma:
		return RltStmAggProSma{X: x.RltStmAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.RltStmAggProAlma:
		return RltStmAggProAlma{X: x.RltStmAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.RltStmAggEma:
		return RltStmAggEma{X: x.RltStmAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.RltStmInrAdd:
		return RltStmInrAdd{X: x.RltStmAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.RltStmInrSub:
		return RltStmInrSub{X: x.RltStmAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.RltStmInrMul:
		return RltStmInrMul{X: x.RltStmAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.RltStmInrDiv:
		return RltStmInrDiv{X: x.RltStmAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.RltStmInrRem:
		return RltStmInrRem{X: x.RltStmAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.RltStmInrPow:
		return RltStmInrPow{X: x.RltStmAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.RltStmInrMin:
		return RltStmInrMin{X: x.RltStmAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.RltStmInrMax:
		return RltStmInrMax{X: x.RltStmAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.RltStmInrSlp:
		return RltStmInrSlp{X: x.RltStmAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.RltStmOtrAdd:
		return RltStmOtrAdd{X: x.RltStmAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.RltStmAct(scp, X.I1)}
	case *xpr.RltStmOtrSub:
		return RltStmOtrSub{X: x.RltStmAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.RltStmAct(scp, X.I1)}
	case *xpr.RltStmOtrMul:
		return RltStmOtrMul{X: x.RltStmAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.RltStmAct(scp, X.I1)}
	case *xpr.RltStmOtrDiv:
		return RltStmOtrDiv{X: x.RltStmAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.RltStmAct(scp, X.I1)}
	case *xpr.RltStmOtrRem:
		return RltStmOtrRem{X: x.RltStmAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.RltStmAct(scp, X.I1)}
	case *xpr.RltStmOtrPow:
		return RltStmOtrPow{X: x.RltStmAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.RltStmAct(scp, X.I1)}
	case *xpr.RltStmOtrMin:
		return RltStmOtrMin{X: x.RltStmAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.RltStmAct(scp, X.I1)}
	case *xpr.RltStmOtrMax:
		return RltStmOtrMax{X: x.RltStmAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.RltStmAct(scp, X.I1)}
	}
	panic(x.Erf("RltStmAct: no action found %v", reflect.ValueOf(v).Elem().Type().Name()))
}
func (x *Actr) RltCndAct(scp *Scp, v xpr.RltCndXpr) RltCndAct {
	switch X := v.(type) {
	case *xpr.RltCndAsn:
		asnScp := scp
		if X.Depth != 0 {
			for d := X.Depth; asnScp != nil && d != 0; d-- { // recurse up scp by depth
				asnScp = asnScp.Prnt
			}
		}
		return RltCndAsn{CndScp: asnScp.RltCnd(x.Txt[X.Idn.Idx:X.Idn.Lim]), X: x.RltCndAct(scp, X.X)}
	case *xpr.RltCndAcs:
		return RltCndAcs{CndScp: scp.RltCnd(x.Txt[X.Trm.Idx:X.Trm.Lim])}
	case *xpr.RltCndsPop:
		return RltCndsPop{X: x.RltCndsAct(scp, X.X)}
	case *xpr.RltCndsDque:
		return RltCndsDque{X: x.RltCndsAct(scp, X.X)}
	case *xpr.RltCndsDel:
		return RltCndsDel{X: x.RltCndsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.RltCndsAt:
		return RltCndsAt{X: x.RltCndsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.RltCndsFst:
		return RltCndsFst{X: x.RltCndsAct(scp, X.X)}
	case *xpr.RltCndsMdl:
		return RltCndsMdl{X: x.RltCndsAct(scp, X.X)}
	case *xpr.RltCndsLst:
		return RltCndsLst{X: x.RltCndsAct(scp, X.X)}
	case *xpr.RltStmSclEql:
		return RltStmSclEql{X: x.RltStmAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.RltStmSclNeq:
		return RltStmSclNeq{X: x.RltStmAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.RltStmSclLss:
		return RltStmSclLss{X: x.RltStmAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.RltStmSclGtr:
		return RltStmSclGtr{X: x.RltStmAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.RltStmSclLeq:
		return RltStmSclLeq{X: x.RltStmAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.RltStmSclGeq:
		return RltStmSclGeq{X: x.RltStmAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.RltStmInrEql:
		return RltStmInrEql{X: x.RltStmAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.RltStmInrNeq:
		return RltStmInrNeq{X: x.RltStmAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.RltStmInrLss:
		return RltStmInrLss{X: x.RltStmAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.RltStmInrGtr:
		return RltStmInrGtr{X: x.RltStmAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.RltStmInrLeq:
		return RltStmInrLeq{X: x.RltStmAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.RltStmInrGeq:
		return RltStmInrGeq{X: x.RltStmAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.RltStmOtrEql:
		return RltStmOtrEql{X: x.RltStmAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.RltStmAct(scp, X.I1)}
	case *xpr.RltStmOtrNeq:
		return RltStmOtrNeq{X: x.RltStmAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.RltStmAct(scp, X.I1)}
	case *xpr.RltStmOtrLss:
		return RltStmOtrLss{X: x.RltStmAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.RltStmAct(scp, X.I1)}
	case *xpr.RltStmOtrGtr:
		return RltStmOtrGtr{X: x.RltStmAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.RltStmAct(scp, X.I1)}
	case *xpr.RltStmOtrLeq:
		return RltStmOtrLeq{X: x.RltStmAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.RltStmAct(scp, X.I1)}
	case *xpr.RltStmOtrGeq:
		return RltStmOtrGeq{X: x.RltStmAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.RltStmAct(scp, X.I1)}
	case *xpr.RltCndAnd:
		return RltCndAnd{X: x.RltCndAct(scp, X.X), I0: x.RltCndAct(scp, X.I0)}
	case *xpr.RltCndSeq:
		return RltCndSeq{X: x.RltCndAct(scp, X.X), I0: x.TmeTmeAct(scp, X.I0), I1: x.RltCndAct(scp, X.I1)}
	}
	panic(x.Erf("RltCndAct: no action found %v", reflect.ValueOf(v).Elem().Type().Name()))
}
func (x *Actr) RltStgyAct(scp *Scp, v xpr.RltStgyXpr) RltStgyAct {
	switch X := v.(type) {
	case *xpr.RltStgyAsn:
		asnScp := scp
		if X.Depth != 0 {
			for d := X.Depth; asnScp != nil && d != 0; d-- { // recurse up scp by depth
				asnScp = asnScp.Prnt
			}
		}
		return RltStgyAsn{StgyScp: asnScp.RltStgy(x.Txt[X.Idn.Idx:X.Idn.Lim]), X: x.RltStgyAct(scp, X.X)}
	case *xpr.RltStgyAcs:
		return RltStgyAcs{StgyScp: scp.RltStgy(x.Txt[X.Trm.Idx:X.Trm.Lim])}
	case *xpr.RltStgysPop:
		return RltStgysPop{X: x.RltStgysAct(scp, X.X)}
	case *xpr.RltStgysDque:
		return RltStgysDque{X: x.RltStgysAct(scp, X.X)}
	case *xpr.RltStgysDel:
		return RltStgysDel{X: x.RltStgysAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.RltStgysAt:
		return RltStgysAt{X: x.RltStgysAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.RltStgysFst:
		return RltStgysFst{X: x.RltStgysAct(scp, X.X)}
	case *xpr.RltStgysMdl:
		return RltStgysMdl{X: x.RltStgysAct(scp, X.X)}
	case *xpr.RltStgysLst:
		return RltStgysLst{X: x.RltStgysAct(scp, X.X)}
	case *xpr.RltCndStgy:
		var i7 []RltCndAct
		for _, cur := range X.I7 {
			i7 = append(i7, x.RltCndAct(scp, cur))
		}
		return RltCndStgy{X: x.RltCndAct(scp, X.X), I0: x.BolBolAct(scp, X.I0), I1: x.FltFltAct(scp, X.I1), I2: x.FltFltAct(scp, X.I2), I3: x.TmeTmeAct(scp, X.I3), I4: x.FltFltAct(scp, X.I4), I5: x.RltInstrAct(scp, X.I5), I6: x.RltStmsAct(scp, X.I6), I7: i7}
	}
	panic(x.Erf("RltStgyAct: no action found %v", reflect.ValueOf(v).Elem().Type().Name()))
}
func (x *Actr) RltPrvsAct(scp *Scp, v xpr.RltPrvsXpr) RltPrvsAct {
	switch X := v.(type) {
	case *xpr.RltPrvsAsn:
		asnScp := scp
		if X.Depth != 0 {
			for d := X.Depth; asnScp != nil && d != 0; d-- { // recurse up scp by depth
				asnScp = asnScp.Prnt
			}
		}
		return RltPrvsAsn{PrvsScp: asnScp.RltPrvs(x.Txt[X.Idn.Idx:X.Idn.Lim]), X: x.RltPrvsAct(scp, X.X)}
	case *xpr.RltPrvsAcs:
		return RltPrvsAcs{PrvsScp: scp.RltPrvs(x.Txt[X.Trm.Idx:X.Trm.Lim])}
	case *xpr.RltPrvsEach:
		eachScp := NewScp(X.Scp, scp)
		return RltPrvsEach{X: x.RltPrvsAct(scp, X.X), PrvScp: eachScp.RltPrv(x.Txt[X.Idn.Idx:X.Idn.Lim]), Acts: x.Acts(eachScp, X.Xprs...)}
	case *xpr.RltPrvsPllEach:
		return RltPrvsPllEach{X: x.RltPrvsAct(scp, X.X), Txt: x.Txt, Idn: X.Idn, ActScpPrnt: scp, XprScp: X.Scp, Xprs: X.Xprs}
	case *xpr.RltNewPrvs:
		var i0 []RltPrvAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.RltPrvAct(scp, cur))
		}
		return RltNewPrvs{I0: i0}
	case *xpr.RltMakePrvs:
		return RltMakePrvs{I0: x.UntUntAct(scp, X.I0)}
	case *xpr.RltMakeEmpPrvs:
		return RltMakeEmpPrvs{I0: x.UntUntAct(scp, X.I0)}
	case *xpr.RltPrvsCpy:
		return RltPrvsCpy{X: x.RltPrvsAct(scp, X.X)}
	case *xpr.RltPrvsClr:
		return RltPrvsClr{X: x.RltPrvsAct(scp, X.X)}
	case *xpr.RltPrvsRand:
		return RltPrvsRand{X: x.RltPrvsAct(scp, X.X)}
	case *xpr.RltPrvsMrg:
		var i0 []RltPrvsAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.RltPrvsAct(scp, cur))
		}
		return RltPrvsMrg{X: x.RltPrvsAct(scp, X.X), I0: i0}
	case *xpr.RltPrvsPush:
		var i0 []RltPrvAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.RltPrvAct(scp, cur))
		}
		return RltPrvsPush{X: x.RltPrvsAct(scp, X.X), I0: i0}
	case *xpr.RltPrvsQue:
		var i0 []RltPrvAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.RltPrvAct(scp, cur))
		}
		return RltPrvsQue{X: x.RltPrvsAct(scp, X.X), I0: i0}
	case *xpr.RltPrvsIns:
		return RltPrvsIns{X: x.RltPrvsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.RltPrvAct(scp, X.I1)}
	case *xpr.RltPrvsUpd:
		return RltPrvsUpd{X: x.RltPrvsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.RltPrvAct(scp, X.I1)}
	case *xpr.RltPrvsIn:
		return RltPrvsIn{X: x.RltPrvsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.UntUntAct(scp, X.I1)}
	case *xpr.RltPrvsInBnd:
		return RltPrvsInBnd{X: x.RltPrvsAct(scp, X.X), I0: x.BndBndAct(scp, X.I0)}
	case *xpr.RltPrvsFrom:
		return RltPrvsFrom{X: x.RltPrvsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.RltPrvsTo:
		return RltPrvsTo{X: x.RltPrvsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.RltPrvsRev:
		return RltPrvsRev{X: x.RltPrvsAct(scp, X.X)}
	}
	panic(x.Erf("RltPrvsAct: no action found %v", reflect.ValueOf(v).Elem().Type().Name()))
}
func (x *Actr) RltInstrsAct(scp *Scp, v xpr.RltInstrsXpr) RltInstrsAct {
	switch X := v.(type) {
	case *xpr.RltInstrsAsn:
		asnScp := scp
		if X.Depth != 0 {
			for d := X.Depth; asnScp != nil && d != 0; d-- { // recurse up scp by depth
				asnScp = asnScp.Prnt
			}
		}
		return RltInstrsAsn{InstrsScp: asnScp.RltInstrs(x.Txt[X.Idn.Idx:X.Idn.Lim]), X: x.RltInstrsAct(scp, X.X)}
	case *xpr.RltInstrsAcs:
		return RltInstrsAcs{InstrsScp: scp.RltInstrs(x.Txt[X.Trm.Idx:X.Trm.Lim])}
	case *xpr.RltInstrsEach:
		eachScp := NewScp(X.Scp, scp)
		return RltInstrsEach{X: x.RltInstrsAct(scp, X.X), InstrScp: eachScp.RltInstr(x.Txt[X.Idn.Idx:X.Idn.Lim]), Acts: x.Acts(eachScp, X.Xprs...)}
	case *xpr.RltInstrsPllEach:
		return RltInstrsPllEach{X: x.RltInstrsAct(scp, X.X), Txt: x.Txt, Idn: X.Idn, ActScpPrnt: scp, XprScp: X.Scp, Xprs: X.Xprs}
	case *xpr.RltNewInstrs:
		var i0 []RltInstrAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.RltInstrAct(scp, cur))
		}
		return RltNewInstrs{I0: i0}
	case *xpr.RltMakeInstrs:
		return RltMakeInstrs{I0: x.UntUntAct(scp, X.I0)}
	case *xpr.RltMakeEmpInstrs:
		return RltMakeEmpInstrs{I0: x.UntUntAct(scp, X.I0)}
	case *xpr.RltInstrsCpy:
		return RltInstrsCpy{X: x.RltInstrsAct(scp, X.X)}
	case *xpr.RltInstrsClr:
		return RltInstrsClr{X: x.RltInstrsAct(scp, X.X)}
	case *xpr.RltInstrsRand:
		return RltInstrsRand{X: x.RltInstrsAct(scp, X.X)}
	case *xpr.RltInstrsMrg:
		var i0 []RltInstrsAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.RltInstrsAct(scp, cur))
		}
		return RltInstrsMrg{X: x.RltInstrsAct(scp, X.X), I0: i0}
	case *xpr.RltInstrsPush:
		var i0 []RltInstrAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.RltInstrAct(scp, cur))
		}
		return RltInstrsPush{X: x.RltInstrsAct(scp, X.X), I0: i0}
	case *xpr.RltInstrsQue:
		var i0 []RltInstrAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.RltInstrAct(scp, cur))
		}
		return RltInstrsQue{X: x.RltInstrsAct(scp, X.X), I0: i0}
	case *xpr.RltInstrsIns:
		return RltInstrsIns{X: x.RltInstrsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.RltInstrAct(scp, X.I1)}
	case *xpr.RltInstrsUpd:
		return RltInstrsUpd{X: x.RltInstrsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.RltInstrAct(scp, X.I1)}
	case *xpr.RltInstrsIn:
		return RltInstrsIn{X: x.RltInstrsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.UntUntAct(scp, X.I1)}
	case *xpr.RltInstrsInBnd:
		return RltInstrsInBnd{X: x.RltInstrsAct(scp, X.X), I0: x.BndBndAct(scp, X.I0)}
	case *xpr.RltInstrsFrom:
		return RltInstrsFrom{X: x.RltInstrsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.RltInstrsTo:
		return RltInstrsTo{X: x.RltInstrsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.RltInstrsRev:
		return RltInstrsRev{X: x.RltInstrsAct(scp, X.X)}
	}
	panic(x.Erf("RltInstrsAct: no action found %v", reflect.ValueOf(v).Elem().Type().Name()))
}
func (x *Actr) RltInrvlsAct(scp *Scp, v xpr.RltInrvlsXpr) RltInrvlsAct {
	switch X := v.(type) {
	case *xpr.RltInrvlsAsn:
		asnScp := scp
		if X.Depth != 0 {
			for d := X.Depth; asnScp != nil && d != 0; d-- { // recurse up scp by depth
				asnScp = asnScp.Prnt
			}
		}
		return RltInrvlsAsn{InrvlsScp: asnScp.RltInrvls(x.Txt[X.Idn.Idx:X.Idn.Lim]), X: x.RltInrvlsAct(scp, X.X)}
	case *xpr.RltInrvlsAcs:
		return RltInrvlsAcs{InrvlsScp: scp.RltInrvls(x.Txt[X.Trm.Idx:X.Trm.Lim])}
	case *xpr.RltInrvlsEach:
		eachScp := NewScp(X.Scp, scp)
		return RltInrvlsEach{X: x.RltInrvlsAct(scp, X.X), InrvlScp: eachScp.RltInrvl(x.Txt[X.Idn.Idx:X.Idn.Lim]), Acts: x.Acts(eachScp, X.Xprs...)}
	case *xpr.RltInrvlsPllEach:
		return RltInrvlsPllEach{X: x.RltInrvlsAct(scp, X.X), Txt: x.Txt, Idn: X.Idn, ActScpPrnt: scp, XprScp: X.Scp, Xprs: X.Xprs}
	case *xpr.RltNewInrvls:
		var i0 []RltInrvlAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.RltInrvlAct(scp, cur))
		}
		return RltNewInrvls{I0: i0}
	case *xpr.RltMakeInrvls:
		return RltMakeInrvls{I0: x.UntUntAct(scp, X.I0)}
	case *xpr.RltMakeEmpInrvls:
		return RltMakeEmpInrvls{I0: x.UntUntAct(scp, X.I0)}
	case *xpr.RltInrvlsCpy:
		return RltInrvlsCpy{X: x.RltInrvlsAct(scp, X.X)}
	case *xpr.RltInrvlsClr:
		return RltInrvlsClr{X: x.RltInrvlsAct(scp, X.X)}
	case *xpr.RltInrvlsRand:
		return RltInrvlsRand{X: x.RltInrvlsAct(scp, X.X)}
	case *xpr.RltInrvlsMrg:
		var i0 []RltInrvlsAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.RltInrvlsAct(scp, cur))
		}
		return RltInrvlsMrg{X: x.RltInrvlsAct(scp, X.X), I0: i0}
	case *xpr.RltInrvlsPush:
		var i0 []RltInrvlAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.RltInrvlAct(scp, cur))
		}
		return RltInrvlsPush{X: x.RltInrvlsAct(scp, X.X), I0: i0}
	case *xpr.RltInrvlsQue:
		var i0 []RltInrvlAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.RltInrvlAct(scp, cur))
		}
		return RltInrvlsQue{X: x.RltInrvlsAct(scp, X.X), I0: i0}
	case *xpr.RltInrvlsIns:
		return RltInrvlsIns{X: x.RltInrvlsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.RltInrvlAct(scp, X.I1)}
	case *xpr.RltInrvlsUpd:
		return RltInrvlsUpd{X: x.RltInrvlsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.RltInrvlAct(scp, X.I1)}
	case *xpr.RltInrvlsIn:
		return RltInrvlsIn{X: x.RltInrvlsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.UntUntAct(scp, X.I1)}
	case *xpr.RltInrvlsInBnd:
		return RltInrvlsInBnd{X: x.RltInrvlsAct(scp, X.X), I0: x.BndBndAct(scp, X.I0)}
	case *xpr.RltInrvlsFrom:
		return RltInrvlsFrom{X: x.RltInrvlsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.RltInrvlsTo:
		return RltInrvlsTo{X: x.RltInrvlsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.RltInrvlsRev:
		return RltInrvlsRev{X: x.RltInrvlsAct(scp, X.X)}
	}
	panic(x.Erf("RltInrvlsAct: no action found %v", reflect.ValueOf(v).Elem().Type().Name()))
}
func (x *Actr) RltSidesAct(scp *Scp, v xpr.RltSidesXpr) RltSidesAct {
	switch X := v.(type) {
	case *xpr.RltSidesAsn:
		asnScp := scp
		if X.Depth != 0 {
			for d := X.Depth; asnScp != nil && d != 0; d-- { // recurse up scp by depth
				asnScp = asnScp.Prnt
			}
		}
		return RltSidesAsn{SidesScp: asnScp.RltSides(x.Txt[X.Idn.Idx:X.Idn.Lim]), X: x.RltSidesAct(scp, X.X)}
	case *xpr.RltSidesAcs:
		return RltSidesAcs{SidesScp: scp.RltSides(x.Txt[X.Trm.Idx:X.Trm.Lim])}
	case *xpr.RltSidesEach:
		eachScp := NewScp(X.Scp, scp)
		return RltSidesEach{X: x.RltSidesAct(scp, X.X), SideScp: eachScp.RltSide(x.Txt[X.Idn.Idx:X.Idn.Lim]), Acts: x.Acts(eachScp, X.Xprs...)}
	case *xpr.RltSidesPllEach:
		return RltSidesPllEach{X: x.RltSidesAct(scp, X.X), Txt: x.Txt, Idn: X.Idn, ActScpPrnt: scp, XprScp: X.Scp, Xprs: X.Xprs}
	case *xpr.RltNewSides:
		var i0 []RltSideAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.RltSideAct(scp, cur))
		}
		return RltNewSides{I0: i0}
	case *xpr.RltMakeSides:
		return RltMakeSides{I0: x.UntUntAct(scp, X.I0)}
	case *xpr.RltMakeEmpSides:
		return RltMakeEmpSides{I0: x.UntUntAct(scp, X.I0)}
	case *xpr.RltSidesCpy:
		return RltSidesCpy{X: x.RltSidesAct(scp, X.X)}
	case *xpr.RltSidesClr:
		return RltSidesClr{X: x.RltSidesAct(scp, X.X)}
	case *xpr.RltSidesRand:
		return RltSidesRand{X: x.RltSidesAct(scp, X.X)}
	case *xpr.RltSidesMrg:
		var i0 []RltSidesAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.RltSidesAct(scp, cur))
		}
		return RltSidesMrg{X: x.RltSidesAct(scp, X.X), I0: i0}
	case *xpr.RltSidesPush:
		var i0 []RltSideAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.RltSideAct(scp, cur))
		}
		return RltSidesPush{X: x.RltSidesAct(scp, X.X), I0: i0}
	case *xpr.RltSidesQue:
		var i0 []RltSideAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.RltSideAct(scp, cur))
		}
		return RltSidesQue{X: x.RltSidesAct(scp, X.X), I0: i0}
	case *xpr.RltSidesIns:
		return RltSidesIns{X: x.RltSidesAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.RltSideAct(scp, X.I1)}
	case *xpr.RltSidesUpd:
		return RltSidesUpd{X: x.RltSidesAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.RltSideAct(scp, X.I1)}
	case *xpr.RltSidesIn:
		return RltSidesIn{X: x.RltSidesAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.UntUntAct(scp, X.I1)}
	case *xpr.RltSidesInBnd:
		return RltSidesInBnd{X: x.RltSidesAct(scp, X.X), I0: x.BndBndAct(scp, X.I0)}
	case *xpr.RltSidesFrom:
		return RltSidesFrom{X: x.RltSidesAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.RltSidesTo:
		return RltSidesTo{X: x.RltSidesAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.RltSidesRev:
		return RltSidesRev{X: x.RltSidesAct(scp, X.X)}
	}
	panic(x.Erf("RltSidesAct: no action found %v", reflect.ValueOf(v).Elem().Type().Name()))
}
func (x *Actr) RltStmsAct(scp *Scp, v xpr.RltStmsXpr) RltStmsAct {
	switch X := v.(type) {
	case *xpr.RltStmsAsn:
		asnScp := scp
		if X.Depth != 0 {
			for d := X.Depth; asnScp != nil && d != 0; d-- { // recurse up scp by depth
				asnScp = asnScp.Prnt
			}
		}
		return RltStmsAsn{StmsScp: asnScp.RltStms(x.Txt[X.Idn.Idx:X.Idn.Lim]), X: x.RltStmsAct(scp, X.X)}
	case *xpr.RltStmsAcs:
		return RltStmsAcs{StmsScp: scp.RltStms(x.Txt[X.Trm.Idx:X.Trm.Lim])}
	case *xpr.RltStmsEach:
		eachScp := NewScp(X.Scp, scp)
		return RltStmsEach{X: x.RltStmsAct(scp, X.X), StmScp: eachScp.RltStm(x.Txt[X.Idn.Idx:X.Idn.Lim]), Acts: x.Acts(eachScp, X.Xprs...)}
	case *xpr.RltStmsPllEach:
		return RltStmsPllEach{X: x.RltStmsAct(scp, X.X), Txt: x.Txt, Idn: X.Idn, ActScpPrnt: scp, XprScp: X.Scp, Xprs: X.Xprs}
	case *xpr.RltNewStms:
		var i0 []RltStmAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.RltStmAct(scp, cur))
		}
		return RltNewStms{I0: i0}
	case *xpr.RltMakeStms:
		return RltMakeStms{I0: x.UntUntAct(scp, X.I0)}
	case *xpr.RltMakeEmpStms:
		return RltMakeEmpStms{I0: x.UntUntAct(scp, X.I0)}
	case *xpr.RltStmsCpy:
		return RltStmsCpy{X: x.RltStmsAct(scp, X.X)}
	case *xpr.RltStmsClr:
		return RltStmsClr{X: x.RltStmsAct(scp, X.X)}
	case *xpr.RltStmsRand:
		return RltStmsRand{X: x.RltStmsAct(scp, X.X)}
	case *xpr.RltStmsMrg:
		var i0 []RltStmsAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.RltStmsAct(scp, cur))
		}
		return RltStmsMrg{X: x.RltStmsAct(scp, X.X), I0: i0}
	case *xpr.RltStmsPush:
		var i0 []RltStmAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.RltStmAct(scp, cur))
		}
		return RltStmsPush{X: x.RltStmsAct(scp, X.X), I0: i0}
	case *xpr.RltStmsQue:
		var i0 []RltStmAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.RltStmAct(scp, cur))
		}
		return RltStmsQue{X: x.RltStmsAct(scp, X.X), I0: i0}
	case *xpr.RltStmsIns:
		return RltStmsIns{X: x.RltStmsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.RltStmAct(scp, X.I1)}
	case *xpr.RltStmsUpd:
		return RltStmsUpd{X: x.RltStmsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.RltStmAct(scp, X.I1)}
	case *xpr.RltStmsIn:
		return RltStmsIn{X: x.RltStmsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.UntUntAct(scp, X.I1)}
	case *xpr.RltStmsInBnd:
		return RltStmsInBnd{X: x.RltStmsAct(scp, X.X), I0: x.BndBndAct(scp, X.I0)}
	case *xpr.RltStmsFrom:
		return RltStmsFrom{X: x.RltStmsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.RltStmsTo:
		return RltStmsTo{X: x.RltStmsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.RltStmsRev:
		return RltStmsRev{X: x.RltStmsAct(scp, X.X)}
	}
	panic(x.Erf("RltStmsAct: no action found %v", reflect.ValueOf(v).Elem().Type().Name()))
}
func (x *Actr) RltCndsAct(scp *Scp, v xpr.RltCndsXpr) RltCndsAct {
	switch X := v.(type) {
	case *xpr.RltCndsAsn:
		asnScp := scp
		if X.Depth != 0 {
			for d := X.Depth; asnScp != nil && d != 0; d-- { // recurse up scp by depth
				asnScp = asnScp.Prnt
			}
		}
		return RltCndsAsn{CndsScp: asnScp.RltCnds(x.Txt[X.Idn.Idx:X.Idn.Lim]), X: x.RltCndsAct(scp, X.X)}
	case *xpr.RltCndsAcs:
		return RltCndsAcs{CndsScp: scp.RltCnds(x.Txt[X.Trm.Idx:X.Trm.Lim])}
	case *xpr.RltCndsEach:
		eachScp := NewScp(X.Scp, scp)
		return RltCndsEach{X: x.RltCndsAct(scp, X.X), CndScp: eachScp.RltCnd(x.Txt[X.Idn.Idx:X.Idn.Lim]), Acts: x.Acts(eachScp, X.Xprs...)}
	case *xpr.RltCndsPllEach:
		return RltCndsPllEach{X: x.RltCndsAct(scp, X.X), Txt: x.Txt, Idn: X.Idn, ActScpPrnt: scp, XprScp: X.Scp, Xprs: X.Xprs}
	case *xpr.RltNewCnds:
		var i0 []RltCndAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.RltCndAct(scp, cur))
		}
		return RltNewCnds{I0: i0}
	case *xpr.RltMakeCnds:
		return RltMakeCnds{I0: x.UntUntAct(scp, X.I0)}
	case *xpr.RltMakeEmpCnds:
		return RltMakeEmpCnds{I0: x.UntUntAct(scp, X.I0)}
	case *xpr.RltCndsCpy:
		return RltCndsCpy{X: x.RltCndsAct(scp, X.X)}
	case *xpr.RltCndsClr:
		return RltCndsClr{X: x.RltCndsAct(scp, X.X)}
	case *xpr.RltCndsRand:
		return RltCndsRand{X: x.RltCndsAct(scp, X.X)}
	case *xpr.RltCndsMrg:
		var i0 []RltCndsAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.RltCndsAct(scp, cur))
		}
		return RltCndsMrg{X: x.RltCndsAct(scp, X.X), I0: i0}
	case *xpr.RltCndsPush:
		var i0 []RltCndAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.RltCndAct(scp, cur))
		}
		return RltCndsPush{X: x.RltCndsAct(scp, X.X), I0: i0}
	case *xpr.RltCndsQue:
		var i0 []RltCndAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.RltCndAct(scp, cur))
		}
		return RltCndsQue{X: x.RltCndsAct(scp, X.X), I0: i0}
	case *xpr.RltCndsIns:
		return RltCndsIns{X: x.RltCndsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.RltCndAct(scp, X.I1)}
	case *xpr.RltCndsUpd:
		return RltCndsUpd{X: x.RltCndsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.RltCndAct(scp, X.I1)}
	case *xpr.RltCndsIn:
		return RltCndsIn{X: x.RltCndsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.UntUntAct(scp, X.I1)}
	case *xpr.RltCndsInBnd:
		return RltCndsInBnd{X: x.RltCndsAct(scp, X.X), I0: x.BndBndAct(scp, X.I0)}
	case *xpr.RltCndsFrom:
		return RltCndsFrom{X: x.RltCndsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.RltCndsTo:
		return RltCndsTo{X: x.RltCndsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.RltCndsRev:
		return RltCndsRev{X: x.RltCndsAct(scp, X.X)}
	}
	panic(x.Erf("RltCndsAct: no action found %v", reflect.ValueOf(v).Elem().Type().Name()))
}
func (x *Actr) RltStgysAct(scp *Scp, v xpr.RltStgysXpr) RltStgysAct {
	switch X := v.(type) {
	case *xpr.RltStgysAsn:
		asnScp := scp
		if X.Depth != 0 {
			for d := X.Depth; asnScp != nil && d != 0; d-- { // recurse up scp by depth
				asnScp = asnScp.Prnt
			}
		}
		return RltStgysAsn{StgysScp: asnScp.RltStgys(x.Txt[X.Idn.Idx:X.Idn.Lim]), X: x.RltStgysAct(scp, X.X)}
	case *xpr.RltStgysAcs:
		return RltStgysAcs{StgysScp: scp.RltStgys(x.Txt[X.Trm.Idx:X.Trm.Lim])}
	case *xpr.RltStgysEach:
		eachScp := NewScp(X.Scp, scp)
		return RltStgysEach{X: x.RltStgysAct(scp, X.X), StgyScp: eachScp.RltStgy(x.Txt[X.Idn.Idx:X.Idn.Lim]), Acts: x.Acts(eachScp, X.Xprs...)}
	case *xpr.RltStgysPllEach:
		return RltStgysPllEach{X: x.RltStgysAct(scp, X.X), Txt: x.Txt, Idn: X.Idn, ActScpPrnt: scp, XprScp: X.Scp, Xprs: X.Xprs}
	case *xpr.RltNewStgys:
		var i0 []RltStgyAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.RltStgyAct(scp, cur))
		}
		return RltNewStgys{I0: i0}
	case *xpr.RltMakeStgys:
		return RltMakeStgys{I0: x.UntUntAct(scp, X.I0)}
	case *xpr.RltMakeEmpStgys:
		return RltMakeEmpStgys{I0: x.UntUntAct(scp, X.I0)}
	case *xpr.RltStgysCpy:
		return RltStgysCpy{X: x.RltStgysAct(scp, X.X)}
	case *xpr.RltStgysClr:
		return RltStgysClr{X: x.RltStgysAct(scp, X.X)}
	case *xpr.RltStgysRand:
		return RltStgysRand{X: x.RltStgysAct(scp, X.X)}
	case *xpr.RltStgysMrg:
		var i0 []RltStgysAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.RltStgysAct(scp, cur))
		}
		return RltStgysMrg{X: x.RltStgysAct(scp, X.X), I0: i0}
	case *xpr.RltStgysPush:
		var i0 []RltStgyAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.RltStgyAct(scp, cur))
		}
		return RltStgysPush{X: x.RltStgysAct(scp, X.X), I0: i0}
	case *xpr.RltStgysQue:
		var i0 []RltStgyAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.RltStgyAct(scp, cur))
		}
		return RltStgysQue{X: x.RltStgysAct(scp, X.X), I0: i0}
	case *xpr.RltStgysIns:
		return RltStgysIns{X: x.RltStgysAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.RltStgyAct(scp, X.I1)}
	case *xpr.RltStgysUpd:
		return RltStgysUpd{X: x.RltStgysAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.RltStgyAct(scp, X.I1)}
	case *xpr.RltStgysIn:
		return RltStgysIn{X: x.RltStgysAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.UntUntAct(scp, X.I1)}
	case *xpr.RltStgysInBnd:
		return RltStgysInBnd{X: x.RltStgysAct(scp, X.X), I0: x.BndBndAct(scp, X.I0)}
	case *xpr.RltStgysFrom:
		return RltStgysFrom{X: x.RltStgysAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.RltStgysTo:
		return RltStgysTo{X: x.RltStgysAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.RltStgysRev:
		return RltStgysRev{X: x.RltStgysAct(scp, X.X)}
	}
	panic(x.Erf("RltStgysAct: no action found %v", reflect.ValueOf(v).Elem().Type().Name()))
}
func (x *Actr) FntFntAct(scp *Scp, v xpr.FntFntXpr) FntFntAct {
	switch X := v.(type) {
	case *xpr.FntFntAsn:
		asnScp := scp
		if X.Depth != 0 {
			for d := X.Depth; asnScp != nil && d != 0; d-- { // recurse up scp by depth
				asnScp = asnScp.Prnt
			}
		}
		return FntFntAsn{FntScp: asnScp.FntFnt(x.Txt[X.Idn.Idx:X.Idn.Lim]), X: x.FntFntAct(scp, X.X)}
	case *xpr.FntFntAcs:
		return FntFntAcs{FntScp: scp.FntFnt(x.Txt[X.Trm.Idx:X.Trm.Lim])}
	}
	panic(x.Erf("FntFntAct: no action found %v", reflect.ValueOf(v).Elem().Type().Name()))
}
func (x *Actr) ClrClrAct(scp *Scp, v xpr.ClrClrXpr) ClrClrAct {
	switch X := v.(type) {
	case *xpr.ClrClrAsn:
		asnScp := scp
		if X.Depth != 0 {
			for d := X.Depth; asnScp != nil && d != 0; d-- { // recurse up scp by depth
				asnScp = asnScp.Prnt
			}
		}
		return ClrClrAsn{ClrScp: asnScp.ClrClr(x.Txt[X.Idn.Idx:X.Idn.Lim]), X: x.ClrClrAct(scp, X.X)}
	case *xpr.ClrClrAcs:
		return ClrClrAcs{ClrScp: scp.ClrClr(x.Txt[X.Trm.Idx:X.Trm.Lim])}
	case *xpr.PenPenClrSetGet:
		if X.I0 == nil {
			return PenPenClrSetGet{X: x.PenPenAct(scp, X.X)}
		} else {
			return PenPenClrSetGet{X: x.PenPenAct(scp, X.X), I0: x.ClrClrAct(scp, X.I0)}
		}
	case *xpr.ClrBlack:
		return ClrBlack{}
	case *xpr.ClrWhite:
		return ClrWhite{}
	case *xpr.ClrRed50:
		return ClrRed50{}
	case *xpr.ClrRed100:
		return ClrRed100{}
	case *xpr.ClrRed200:
		return ClrRed200{}
	case *xpr.ClrRed300:
		return ClrRed300{}
	case *xpr.ClrRed400:
		return ClrRed400{}
	case *xpr.ClrRed500:
		return ClrRed500{}
	case *xpr.ClrRed600:
		return ClrRed600{}
	case *xpr.ClrRed700:
		return ClrRed700{}
	case *xpr.ClrRed800:
		return ClrRed800{}
	case *xpr.ClrRed900:
		return ClrRed900{}
	case *xpr.ClrRedA100:
		return ClrRedA100{}
	case *xpr.ClrRedA200:
		return ClrRedA200{}
	case *xpr.ClrRedA400:
		return ClrRedA400{}
	case *xpr.ClrRedA700:
		return ClrRedA700{}
	case *xpr.ClrPink50:
		return ClrPink50{}
	case *xpr.ClrPink100:
		return ClrPink100{}
	case *xpr.ClrPink200:
		return ClrPink200{}
	case *xpr.ClrPink300:
		return ClrPink300{}
	case *xpr.ClrPink400:
		return ClrPink400{}
	case *xpr.ClrPink500:
		return ClrPink500{}
	case *xpr.ClrPink600:
		return ClrPink600{}
	case *xpr.ClrPink700:
		return ClrPink700{}
	case *xpr.ClrPink800:
		return ClrPink800{}
	case *xpr.ClrPink900:
		return ClrPink900{}
	case *xpr.ClrPinkA100:
		return ClrPinkA100{}
	case *xpr.ClrPinkA200:
		return ClrPinkA200{}
	case *xpr.ClrPinkA400:
		return ClrPinkA400{}
	case *xpr.ClrPinkA700:
		return ClrPinkA700{}
	case *xpr.ClrPurple50:
		return ClrPurple50{}
	case *xpr.ClrPurple100:
		return ClrPurple100{}
	case *xpr.ClrPurple200:
		return ClrPurple200{}
	case *xpr.ClrPurple300:
		return ClrPurple300{}
	case *xpr.ClrPurple400:
		return ClrPurple400{}
	case *xpr.ClrPurple500:
		return ClrPurple500{}
	case *xpr.ClrPurple600:
		return ClrPurple600{}
	case *xpr.ClrPurple700:
		return ClrPurple700{}
	case *xpr.ClrPurple800:
		return ClrPurple800{}
	case *xpr.ClrPurple900:
		return ClrPurple900{}
	case *xpr.ClrPurpleA100:
		return ClrPurpleA100{}
	case *xpr.ClrPurpleA200:
		return ClrPurpleA200{}
	case *xpr.ClrPurpleA400:
		return ClrPurpleA400{}
	case *xpr.ClrPurpleA700:
		return ClrPurpleA700{}
	case *xpr.ClrDeepPurple50:
		return ClrDeepPurple50{}
	case *xpr.ClrDeepPurple100:
		return ClrDeepPurple100{}
	case *xpr.ClrDeepPurple200:
		return ClrDeepPurple200{}
	case *xpr.ClrDeepPurple300:
		return ClrDeepPurple300{}
	case *xpr.ClrDeepPurple400:
		return ClrDeepPurple400{}
	case *xpr.ClrDeepPurple500:
		return ClrDeepPurple500{}
	case *xpr.ClrDeepPurple600:
		return ClrDeepPurple600{}
	case *xpr.ClrDeepPurple700:
		return ClrDeepPurple700{}
	case *xpr.ClrDeepPurple800:
		return ClrDeepPurple800{}
	case *xpr.ClrDeepPurple900:
		return ClrDeepPurple900{}
	case *xpr.ClrDeepPurpleA100:
		return ClrDeepPurpleA100{}
	case *xpr.ClrDeepPurpleA200:
		return ClrDeepPurpleA200{}
	case *xpr.ClrDeepPurpleA400:
		return ClrDeepPurpleA400{}
	case *xpr.ClrDeepPurpleA700:
		return ClrDeepPurpleA700{}
	case *xpr.ClrIndigo50:
		return ClrIndigo50{}
	case *xpr.ClrIndigo100:
		return ClrIndigo100{}
	case *xpr.ClrIndigo200:
		return ClrIndigo200{}
	case *xpr.ClrIndigo300:
		return ClrIndigo300{}
	case *xpr.ClrIndigo400:
		return ClrIndigo400{}
	case *xpr.ClrIndigo500:
		return ClrIndigo500{}
	case *xpr.ClrIndigo600:
		return ClrIndigo600{}
	case *xpr.ClrIndigo700:
		return ClrIndigo700{}
	case *xpr.ClrIndigo800:
		return ClrIndigo800{}
	case *xpr.ClrIndigo900:
		return ClrIndigo900{}
	case *xpr.ClrIndigoA100:
		return ClrIndigoA100{}
	case *xpr.ClrIndigoA200:
		return ClrIndigoA200{}
	case *xpr.ClrIndigoA400:
		return ClrIndigoA400{}
	case *xpr.ClrIndigoA700:
		return ClrIndigoA700{}
	case *xpr.ClrBlue50:
		return ClrBlue50{}
	case *xpr.ClrBlue100:
		return ClrBlue100{}
	case *xpr.ClrBlue200:
		return ClrBlue200{}
	case *xpr.ClrBlue300:
		return ClrBlue300{}
	case *xpr.ClrBlue400:
		return ClrBlue400{}
	case *xpr.ClrBlue500:
		return ClrBlue500{}
	case *xpr.ClrBlue600:
		return ClrBlue600{}
	case *xpr.ClrBlue700:
		return ClrBlue700{}
	case *xpr.ClrBlue800:
		return ClrBlue800{}
	case *xpr.ClrBlue900:
		return ClrBlue900{}
	case *xpr.ClrBlueA100:
		return ClrBlueA100{}
	case *xpr.ClrBlueA200:
		return ClrBlueA200{}
	case *xpr.ClrBlueA400:
		return ClrBlueA400{}
	case *xpr.ClrBlueA700:
		return ClrBlueA700{}
	case *xpr.ClrLightBlue50:
		return ClrLightBlue50{}
	case *xpr.ClrLightBlue100:
		return ClrLightBlue100{}
	case *xpr.ClrLightBlue200:
		return ClrLightBlue200{}
	case *xpr.ClrLightBlue300:
		return ClrLightBlue300{}
	case *xpr.ClrLightBlue400:
		return ClrLightBlue400{}
	case *xpr.ClrLightBlue500:
		return ClrLightBlue500{}
	case *xpr.ClrLightBlue600:
		return ClrLightBlue600{}
	case *xpr.ClrLightBlue700:
		return ClrLightBlue700{}
	case *xpr.ClrLightBlue800:
		return ClrLightBlue800{}
	case *xpr.ClrLightBlue900:
		return ClrLightBlue900{}
	case *xpr.ClrLightBlueA100:
		return ClrLightBlueA100{}
	case *xpr.ClrLightBlueA200:
		return ClrLightBlueA200{}
	case *xpr.ClrLightBlueA400:
		return ClrLightBlueA400{}
	case *xpr.ClrLightBlueA700:
		return ClrLightBlueA700{}
	case *xpr.ClrCyan50:
		return ClrCyan50{}
	case *xpr.ClrCyan100:
		return ClrCyan100{}
	case *xpr.ClrCyan200:
		return ClrCyan200{}
	case *xpr.ClrCyan300:
		return ClrCyan300{}
	case *xpr.ClrCyan400:
		return ClrCyan400{}
	case *xpr.ClrCyan500:
		return ClrCyan500{}
	case *xpr.ClrCyan600:
		return ClrCyan600{}
	case *xpr.ClrCyan700:
		return ClrCyan700{}
	case *xpr.ClrCyan800:
		return ClrCyan800{}
	case *xpr.ClrCyan900:
		return ClrCyan900{}
	case *xpr.ClrCyanA100:
		return ClrCyanA100{}
	case *xpr.ClrCyanA200:
		return ClrCyanA200{}
	case *xpr.ClrCyanA400:
		return ClrCyanA400{}
	case *xpr.ClrCyanA700:
		return ClrCyanA700{}
	case *xpr.ClrTeal50:
		return ClrTeal50{}
	case *xpr.ClrTeal100:
		return ClrTeal100{}
	case *xpr.ClrTeal200:
		return ClrTeal200{}
	case *xpr.ClrTeal300:
		return ClrTeal300{}
	case *xpr.ClrTeal400:
		return ClrTeal400{}
	case *xpr.ClrTeal500:
		return ClrTeal500{}
	case *xpr.ClrTeal600:
		return ClrTeal600{}
	case *xpr.ClrTeal700:
		return ClrTeal700{}
	case *xpr.ClrTeal800:
		return ClrTeal800{}
	case *xpr.ClrTeal900:
		return ClrTeal900{}
	case *xpr.ClrTealA100:
		return ClrTealA100{}
	case *xpr.ClrTealA200:
		return ClrTealA200{}
	case *xpr.ClrTealA400:
		return ClrTealA400{}
	case *xpr.ClrTealA700:
		return ClrTealA700{}
	case *xpr.ClrGreen50:
		return ClrGreen50{}
	case *xpr.ClrGreen100:
		return ClrGreen100{}
	case *xpr.ClrGreen200:
		return ClrGreen200{}
	case *xpr.ClrGreen300:
		return ClrGreen300{}
	case *xpr.ClrGreen400:
		return ClrGreen400{}
	case *xpr.ClrGreen500:
		return ClrGreen500{}
	case *xpr.ClrGreen600:
		return ClrGreen600{}
	case *xpr.ClrGreen700:
		return ClrGreen700{}
	case *xpr.ClrGreen800:
		return ClrGreen800{}
	case *xpr.ClrGreen900:
		return ClrGreen900{}
	case *xpr.ClrGreenA100:
		return ClrGreenA100{}
	case *xpr.ClrGreenA200:
		return ClrGreenA200{}
	case *xpr.ClrGreenA400:
		return ClrGreenA400{}
	case *xpr.ClrGreenA700:
		return ClrGreenA700{}
	case *xpr.ClrLightGreen50:
		return ClrLightGreen50{}
	case *xpr.ClrLightGreen100:
		return ClrLightGreen100{}
	case *xpr.ClrLightGreen200:
		return ClrLightGreen200{}
	case *xpr.ClrLightGreen300:
		return ClrLightGreen300{}
	case *xpr.ClrLightGreen400:
		return ClrLightGreen400{}
	case *xpr.ClrLightGreen500:
		return ClrLightGreen500{}
	case *xpr.ClrLightGreen600:
		return ClrLightGreen600{}
	case *xpr.ClrLightGreen700:
		return ClrLightGreen700{}
	case *xpr.ClrLightGreen800:
		return ClrLightGreen800{}
	case *xpr.ClrLightGreen900:
		return ClrLightGreen900{}
	case *xpr.ClrLightGreenA100:
		return ClrLightGreenA100{}
	case *xpr.ClrLightGreenA200:
		return ClrLightGreenA200{}
	case *xpr.ClrLightGreenA400:
		return ClrLightGreenA400{}
	case *xpr.ClrLightGreenA700:
		return ClrLightGreenA700{}
	case *xpr.ClrLime50:
		return ClrLime50{}
	case *xpr.ClrLime100:
		return ClrLime100{}
	case *xpr.ClrLime200:
		return ClrLime200{}
	case *xpr.ClrLime300:
		return ClrLime300{}
	case *xpr.ClrLime400:
		return ClrLime400{}
	case *xpr.ClrLime500:
		return ClrLime500{}
	case *xpr.ClrLime600:
		return ClrLime600{}
	case *xpr.ClrLime700:
		return ClrLime700{}
	case *xpr.ClrLime800:
		return ClrLime800{}
	case *xpr.ClrLime900:
		return ClrLime900{}
	case *xpr.ClrLimeA100:
		return ClrLimeA100{}
	case *xpr.ClrLimeA200:
		return ClrLimeA200{}
	case *xpr.ClrLimeA400:
		return ClrLimeA400{}
	case *xpr.ClrLimeA700:
		return ClrLimeA700{}
	case *xpr.ClrYellow50:
		return ClrYellow50{}
	case *xpr.ClrYellow100:
		return ClrYellow100{}
	case *xpr.ClrYellow200:
		return ClrYellow200{}
	case *xpr.ClrYellow300:
		return ClrYellow300{}
	case *xpr.ClrYellow400:
		return ClrYellow400{}
	case *xpr.ClrYellow500:
		return ClrYellow500{}
	case *xpr.ClrYellow600:
		return ClrYellow600{}
	case *xpr.ClrYellow700:
		return ClrYellow700{}
	case *xpr.ClrYellow800:
		return ClrYellow800{}
	case *xpr.ClrYellow900:
		return ClrYellow900{}
	case *xpr.ClrYellowA100:
		return ClrYellowA100{}
	case *xpr.ClrYellowA200:
		return ClrYellowA200{}
	case *xpr.ClrYellowA400:
		return ClrYellowA400{}
	case *xpr.ClrYellowA700:
		return ClrYellowA700{}
	case *xpr.ClrAmber50:
		return ClrAmber50{}
	case *xpr.ClrAmber100:
		return ClrAmber100{}
	case *xpr.ClrAmber200:
		return ClrAmber200{}
	case *xpr.ClrAmber300:
		return ClrAmber300{}
	case *xpr.ClrAmber400:
		return ClrAmber400{}
	case *xpr.ClrAmber500:
		return ClrAmber500{}
	case *xpr.ClrAmber600:
		return ClrAmber600{}
	case *xpr.ClrAmber700:
		return ClrAmber700{}
	case *xpr.ClrAmber800:
		return ClrAmber800{}
	case *xpr.ClrAmber900:
		return ClrAmber900{}
	case *xpr.ClrAmberA100:
		return ClrAmberA100{}
	case *xpr.ClrAmberA200:
		return ClrAmberA200{}
	case *xpr.ClrAmberA400:
		return ClrAmberA400{}
	case *xpr.ClrAmberA700:
		return ClrAmberA700{}
	case *xpr.ClrOrange50:
		return ClrOrange50{}
	case *xpr.ClrOrange100:
		return ClrOrange100{}
	case *xpr.ClrOrange200:
		return ClrOrange200{}
	case *xpr.ClrOrange300:
		return ClrOrange300{}
	case *xpr.ClrOrange400:
		return ClrOrange400{}
	case *xpr.ClrOrange500:
		return ClrOrange500{}
	case *xpr.ClrOrange600:
		return ClrOrange600{}
	case *xpr.ClrOrange700:
		return ClrOrange700{}
	case *xpr.ClrOrange800:
		return ClrOrange800{}
	case *xpr.ClrOrange900:
		return ClrOrange900{}
	case *xpr.ClrOrangeA100:
		return ClrOrangeA100{}
	case *xpr.ClrOrangeA200:
		return ClrOrangeA200{}
	case *xpr.ClrOrangeA400:
		return ClrOrangeA400{}
	case *xpr.ClrOrangeA700:
		return ClrOrangeA700{}
	case *xpr.ClrDeepOrange50:
		return ClrDeepOrange50{}
	case *xpr.ClrDeepOrange100:
		return ClrDeepOrange100{}
	case *xpr.ClrDeepOrange200:
		return ClrDeepOrange200{}
	case *xpr.ClrDeepOrange300:
		return ClrDeepOrange300{}
	case *xpr.ClrDeepOrange400:
		return ClrDeepOrange400{}
	case *xpr.ClrDeepOrange500:
		return ClrDeepOrange500{}
	case *xpr.ClrDeepOrange600:
		return ClrDeepOrange600{}
	case *xpr.ClrDeepOrange700:
		return ClrDeepOrange700{}
	case *xpr.ClrDeepOrange800:
		return ClrDeepOrange800{}
	case *xpr.ClrDeepOrange900:
		return ClrDeepOrange900{}
	case *xpr.ClrDeepOrangeA100:
		return ClrDeepOrangeA100{}
	case *xpr.ClrDeepOrangeA200:
		return ClrDeepOrangeA200{}
	case *xpr.ClrDeepOrangeA400:
		return ClrDeepOrangeA400{}
	case *xpr.ClrDeepOrangeA700:
		return ClrDeepOrangeA700{}
	case *xpr.ClrBrown50:
		return ClrBrown50{}
	case *xpr.ClrBrown100:
		return ClrBrown100{}
	case *xpr.ClrBrown200:
		return ClrBrown200{}
	case *xpr.ClrBrown300:
		return ClrBrown300{}
	case *xpr.ClrBrown400:
		return ClrBrown400{}
	case *xpr.ClrBrown500:
		return ClrBrown500{}
	case *xpr.ClrBrown600:
		return ClrBrown600{}
	case *xpr.ClrBrown700:
		return ClrBrown700{}
	case *xpr.ClrBrown800:
		return ClrBrown800{}
	case *xpr.ClrBrown900:
		return ClrBrown900{}
	case *xpr.ClrGrey50:
		return ClrGrey50{}
	case *xpr.ClrGrey100:
		return ClrGrey100{}
	case *xpr.ClrGrey200:
		return ClrGrey200{}
	case *xpr.ClrGrey300:
		return ClrGrey300{}
	case *xpr.ClrGrey400:
		return ClrGrey400{}
	case *xpr.ClrGrey500:
		return ClrGrey500{}
	case *xpr.ClrGrey600:
		return ClrGrey600{}
	case *xpr.ClrGrey700:
		return ClrGrey700{}
	case *xpr.ClrGrey800:
		return ClrGrey800{}
	case *xpr.ClrGrey900:
		return ClrGrey900{}
	case *xpr.ClrBlueGrey50:
		return ClrBlueGrey50{}
	case *xpr.ClrBlueGrey100:
		return ClrBlueGrey100{}
	case *xpr.ClrBlueGrey200:
		return ClrBlueGrey200{}
	case *xpr.ClrBlueGrey300:
		return ClrBlueGrey300{}
	case *xpr.ClrBlueGrey400:
		return ClrBlueGrey400{}
	case *xpr.ClrBlueGrey500:
		return ClrBlueGrey500{}
	case *xpr.ClrBlueGrey600:
		return ClrBlueGrey600{}
	case *xpr.ClrBlueGrey700:
		return ClrBlueGrey700{}
	case *xpr.ClrBlueGrey800:
		return ClrBlueGrey800{}
	case *xpr.ClrBlueGrey900:
		return ClrBlueGrey900{}
	case *xpr.ClrBakClr:
		return ClrBakClr{}
	case *xpr.ClrBrdrClr:
		return ClrBrdrClr{}
	case *xpr.ClrInrvlTxtClrX:
		return ClrInrvlTxtClrX{}
	case *xpr.ClrInrvlTxtClrY:
		return ClrInrvlTxtClrY{}
	case *xpr.ClrMsgClr:
		return ClrMsgClr{}
	case *xpr.ClrTitleClr:
		return ClrTitleClr{}
	case *xpr.ClrPrfClr:
		return ClrPrfClr{}
	case *xpr.ClrLosClr:
		return ClrLosClr{}
	case *xpr.ClrRgba:
		return ClrRgba{I0: x.FltFltAct(scp, X.I0), I1: x.FltFltAct(scp, X.I1), I2: x.FltFltAct(scp, X.I2), I3: x.FltFltAct(scp, X.I3)}
	case *xpr.ClrRgb:
		return ClrRgb{I0: x.FltFltAct(scp, X.I0), I1: x.FltFltAct(scp, X.I1), I2: x.FltFltAct(scp, X.I2)}
	case *xpr.ClrHex:
		return ClrHex{I0: x.StrStrAct(scp, X.I0)}
	case *xpr.ClrClrOpa:
		return ClrClrOpa{X: x.ClrClrAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.ClrClrInv:
		return ClrClrInv{X: x.ClrClrAct(scp, X.X)}
	}
	panic(x.Erf("ClrClrAct: no action found %v", reflect.ValueOf(v).Elem().Type().Name()))
}
func (x *Actr) PenPenAct(scp *Scp, v xpr.PenPenXpr) PenPenAct {
	switch X := v.(type) {
	case *xpr.PenPenAsn:
		asnScp := scp
		if X.Depth != 0 {
			for d := X.Depth; asnScp != nil && d != 0; d-- { // recurse up scp by depth
				asnScp = asnScp.Prnt
			}
		}
		return PenPenAsn{PenScp: asnScp.PenPen(x.Txt[X.Idn.Idx:X.Idn.Lim]), X: x.PenPenAct(scp, X.X)}
	case *xpr.PenPenAcs:
		return PenPenAcs{PenScp: scp.PenPen(x.Txt[X.Trm.Idx:X.Trm.Lim])}
	case *xpr.PenBlack:
		return PenBlack{}
	case *xpr.PenWhite:
		return PenWhite{}
	case *xpr.PenRed50:
		return PenRed50{}
	case *xpr.PenRed100:
		return PenRed100{}
	case *xpr.PenRed200:
		return PenRed200{}
	case *xpr.PenRed300:
		return PenRed300{}
	case *xpr.PenRed400:
		return PenRed400{}
	case *xpr.PenRed500:
		return PenRed500{}
	case *xpr.PenRed600:
		return PenRed600{}
	case *xpr.PenRed700:
		return PenRed700{}
	case *xpr.PenRed800:
		return PenRed800{}
	case *xpr.PenRed900:
		return PenRed900{}
	case *xpr.PenRedA100:
		return PenRedA100{}
	case *xpr.PenRedA200:
		return PenRedA200{}
	case *xpr.PenRedA400:
		return PenRedA400{}
	case *xpr.PenRedA700:
		return PenRedA700{}
	case *xpr.PenPink50:
		return PenPink50{}
	case *xpr.PenPink100:
		return PenPink100{}
	case *xpr.PenPink200:
		return PenPink200{}
	case *xpr.PenPink300:
		return PenPink300{}
	case *xpr.PenPink400:
		return PenPink400{}
	case *xpr.PenPink500:
		return PenPink500{}
	case *xpr.PenPink600:
		return PenPink600{}
	case *xpr.PenPink700:
		return PenPink700{}
	case *xpr.PenPink800:
		return PenPink800{}
	case *xpr.PenPink900:
		return PenPink900{}
	case *xpr.PenPinkA100:
		return PenPinkA100{}
	case *xpr.PenPinkA200:
		return PenPinkA200{}
	case *xpr.PenPinkA400:
		return PenPinkA400{}
	case *xpr.PenPinkA700:
		return PenPinkA700{}
	case *xpr.PenPurple50:
		return PenPurple50{}
	case *xpr.PenPurple100:
		return PenPurple100{}
	case *xpr.PenPurple200:
		return PenPurple200{}
	case *xpr.PenPurple300:
		return PenPurple300{}
	case *xpr.PenPurple400:
		return PenPurple400{}
	case *xpr.PenPurple500:
		return PenPurple500{}
	case *xpr.PenPurple600:
		return PenPurple600{}
	case *xpr.PenPurple700:
		return PenPurple700{}
	case *xpr.PenPurple800:
		return PenPurple800{}
	case *xpr.PenPurple900:
		return PenPurple900{}
	case *xpr.PenPurpleA100:
		return PenPurpleA100{}
	case *xpr.PenPurpleA200:
		return PenPurpleA200{}
	case *xpr.PenPurpleA400:
		return PenPurpleA400{}
	case *xpr.PenPurpleA700:
		return PenPurpleA700{}
	case *xpr.PenDeepPurple50:
		return PenDeepPurple50{}
	case *xpr.PenDeepPurple100:
		return PenDeepPurple100{}
	case *xpr.PenDeepPurple200:
		return PenDeepPurple200{}
	case *xpr.PenDeepPurple300:
		return PenDeepPurple300{}
	case *xpr.PenDeepPurple400:
		return PenDeepPurple400{}
	case *xpr.PenDeepPurple500:
		return PenDeepPurple500{}
	case *xpr.PenDeepPurple600:
		return PenDeepPurple600{}
	case *xpr.PenDeepPurple700:
		return PenDeepPurple700{}
	case *xpr.PenDeepPurple800:
		return PenDeepPurple800{}
	case *xpr.PenDeepPurple900:
		return PenDeepPurple900{}
	case *xpr.PenDeepPurpleA100:
		return PenDeepPurpleA100{}
	case *xpr.PenDeepPurpleA200:
		return PenDeepPurpleA200{}
	case *xpr.PenDeepPurpleA400:
		return PenDeepPurpleA400{}
	case *xpr.PenDeepPurpleA700:
		return PenDeepPurpleA700{}
	case *xpr.PenIndigo50:
		return PenIndigo50{}
	case *xpr.PenIndigo100:
		return PenIndigo100{}
	case *xpr.PenIndigo200:
		return PenIndigo200{}
	case *xpr.PenIndigo300:
		return PenIndigo300{}
	case *xpr.PenIndigo400:
		return PenIndigo400{}
	case *xpr.PenIndigo500:
		return PenIndigo500{}
	case *xpr.PenIndigo600:
		return PenIndigo600{}
	case *xpr.PenIndigo700:
		return PenIndigo700{}
	case *xpr.PenIndigo800:
		return PenIndigo800{}
	case *xpr.PenIndigo900:
		return PenIndigo900{}
	case *xpr.PenIndigoA100:
		return PenIndigoA100{}
	case *xpr.PenIndigoA200:
		return PenIndigoA200{}
	case *xpr.PenIndigoA400:
		return PenIndigoA400{}
	case *xpr.PenIndigoA700:
		return PenIndigoA700{}
	case *xpr.PenBlue50:
		return PenBlue50{}
	case *xpr.PenBlue100:
		return PenBlue100{}
	case *xpr.PenBlue200:
		return PenBlue200{}
	case *xpr.PenBlue300:
		return PenBlue300{}
	case *xpr.PenBlue400:
		return PenBlue400{}
	case *xpr.PenBlue500:
		return PenBlue500{}
	case *xpr.PenBlue600:
		return PenBlue600{}
	case *xpr.PenBlue700:
		return PenBlue700{}
	case *xpr.PenBlue800:
		return PenBlue800{}
	case *xpr.PenBlue900:
		return PenBlue900{}
	case *xpr.PenBlueA100:
		return PenBlueA100{}
	case *xpr.PenBlueA200:
		return PenBlueA200{}
	case *xpr.PenBlueA400:
		return PenBlueA400{}
	case *xpr.PenBlueA700:
		return PenBlueA700{}
	case *xpr.PenLightBlue50:
		return PenLightBlue50{}
	case *xpr.PenLightBlue100:
		return PenLightBlue100{}
	case *xpr.PenLightBlue200:
		return PenLightBlue200{}
	case *xpr.PenLightBlue300:
		return PenLightBlue300{}
	case *xpr.PenLightBlue400:
		return PenLightBlue400{}
	case *xpr.PenLightBlue500:
		return PenLightBlue500{}
	case *xpr.PenLightBlue600:
		return PenLightBlue600{}
	case *xpr.PenLightBlue700:
		return PenLightBlue700{}
	case *xpr.PenLightBlue800:
		return PenLightBlue800{}
	case *xpr.PenLightBlue900:
		return PenLightBlue900{}
	case *xpr.PenLightBlueA100:
		return PenLightBlueA100{}
	case *xpr.PenLightBlueA200:
		return PenLightBlueA200{}
	case *xpr.PenLightBlueA400:
		return PenLightBlueA400{}
	case *xpr.PenLightBlueA700:
		return PenLightBlueA700{}
	case *xpr.PenCyan50:
		return PenCyan50{}
	case *xpr.PenCyan100:
		return PenCyan100{}
	case *xpr.PenCyan200:
		return PenCyan200{}
	case *xpr.PenCyan300:
		return PenCyan300{}
	case *xpr.PenCyan400:
		return PenCyan400{}
	case *xpr.PenCyan500:
		return PenCyan500{}
	case *xpr.PenCyan600:
		return PenCyan600{}
	case *xpr.PenCyan700:
		return PenCyan700{}
	case *xpr.PenCyan800:
		return PenCyan800{}
	case *xpr.PenCyan900:
		return PenCyan900{}
	case *xpr.PenCyanA100:
		return PenCyanA100{}
	case *xpr.PenCyanA200:
		return PenCyanA200{}
	case *xpr.PenCyanA400:
		return PenCyanA400{}
	case *xpr.PenCyanA700:
		return PenCyanA700{}
	case *xpr.PenTeal50:
		return PenTeal50{}
	case *xpr.PenTeal100:
		return PenTeal100{}
	case *xpr.PenTeal200:
		return PenTeal200{}
	case *xpr.PenTeal300:
		return PenTeal300{}
	case *xpr.PenTeal400:
		return PenTeal400{}
	case *xpr.PenTeal500:
		return PenTeal500{}
	case *xpr.PenTeal600:
		return PenTeal600{}
	case *xpr.PenTeal700:
		return PenTeal700{}
	case *xpr.PenTeal800:
		return PenTeal800{}
	case *xpr.PenTeal900:
		return PenTeal900{}
	case *xpr.PenTealA100:
		return PenTealA100{}
	case *xpr.PenTealA200:
		return PenTealA200{}
	case *xpr.PenTealA400:
		return PenTealA400{}
	case *xpr.PenTealA700:
		return PenTealA700{}
	case *xpr.PenGreen50:
		return PenGreen50{}
	case *xpr.PenGreen100:
		return PenGreen100{}
	case *xpr.PenGreen200:
		return PenGreen200{}
	case *xpr.PenGreen300:
		return PenGreen300{}
	case *xpr.PenGreen400:
		return PenGreen400{}
	case *xpr.PenGreen500:
		return PenGreen500{}
	case *xpr.PenGreen600:
		return PenGreen600{}
	case *xpr.PenGreen700:
		return PenGreen700{}
	case *xpr.PenGreen800:
		return PenGreen800{}
	case *xpr.PenGreen900:
		return PenGreen900{}
	case *xpr.PenGreenA100:
		return PenGreenA100{}
	case *xpr.PenGreenA200:
		return PenGreenA200{}
	case *xpr.PenGreenA400:
		return PenGreenA400{}
	case *xpr.PenGreenA700:
		return PenGreenA700{}
	case *xpr.PenLightGreen50:
		return PenLightGreen50{}
	case *xpr.PenLightGreen100:
		return PenLightGreen100{}
	case *xpr.PenLightGreen200:
		return PenLightGreen200{}
	case *xpr.PenLightGreen300:
		return PenLightGreen300{}
	case *xpr.PenLightGreen400:
		return PenLightGreen400{}
	case *xpr.PenLightGreen500:
		return PenLightGreen500{}
	case *xpr.PenLightGreen600:
		return PenLightGreen600{}
	case *xpr.PenLightGreen700:
		return PenLightGreen700{}
	case *xpr.PenLightGreen800:
		return PenLightGreen800{}
	case *xpr.PenLightGreen900:
		return PenLightGreen900{}
	case *xpr.PenLightGreenA100:
		return PenLightGreenA100{}
	case *xpr.PenLightGreenA200:
		return PenLightGreenA200{}
	case *xpr.PenLightGreenA400:
		return PenLightGreenA400{}
	case *xpr.PenLightGreenA700:
		return PenLightGreenA700{}
	case *xpr.PenLime50:
		return PenLime50{}
	case *xpr.PenLime100:
		return PenLime100{}
	case *xpr.PenLime200:
		return PenLime200{}
	case *xpr.PenLime300:
		return PenLime300{}
	case *xpr.PenLime400:
		return PenLime400{}
	case *xpr.PenLime500:
		return PenLime500{}
	case *xpr.PenLime600:
		return PenLime600{}
	case *xpr.PenLime700:
		return PenLime700{}
	case *xpr.PenLime800:
		return PenLime800{}
	case *xpr.PenLime900:
		return PenLime900{}
	case *xpr.PenLimeA100:
		return PenLimeA100{}
	case *xpr.PenLimeA200:
		return PenLimeA200{}
	case *xpr.PenLimeA400:
		return PenLimeA400{}
	case *xpr.PenLimeA700:
		return PenLimeA700{}
	case *xpr.PenYellow50:
		return PenYellow50{}
	case *xpr.PenYellow100:
		return PenYellow100{}
	case *xpr.PenYellow200:
		return PenYellow200{}
	case *xpr.PenYellow300:
		return PenYellow300{}
	case *xpr.PenYellow400:
		return PenYellow400{}
	case *xpr.PenYellow500:
		return PenYellow500{}
	case *xpr.PenYellow600:
		return PenYellow600{}
	case *xpr.PenYellow700:
		return PenYellow700{}
	case *xpr.PenYellow800:
		return PenYellow800{}
	case *xpr.PenYellow900:
		return PenYellow900{}
	case *xpr.PenYellowA100:
		return PenYellowA100{}
	case *xpr.PenYellowA200:
		return PenYellowA200{}
	case *xpr.PenYellowA400:
		return PenYellowA400{}
	case *xpr.PenYellowA700:
		return PenYellowA700{}
	case *xpr.PenAmber50:
		return PenAmber50{}
	case *xpr.PenAmber100:
		return PenAmber100{}
	case *xpr.PenAmber200:
		return PenAmber200{}
	case *xpr.PenAmber300:
		return PenAmber300{}
	case *xpr.PenAmber400:
		return PenAmber400{}
	case *xpr.PenAmber500:
		return PenAmber500{}
	case *xpr.PenAmber600:
		return PenAmber600{}
	case *xpr.PenAmber700:
		return PenAmber700{}
	case *xpr.PenAmber800:
		return PenAmber800{}
	case *xpr.PenAmber900:
		return PenAmber900{}
	case *xpr.PenAmberA100:
		return PenAmberA100{}
	case *xpr.PenAmberA200:
		return PenAmberA200{}
	case *xpr.PenAmberA400:
		return PenAmberA400{}
	case *xpr.PenAmberA700:
		return PenAmberA700{}
	case *xpr.PenOrange50:
		return PenOrange50{}
	case *xpr.PenOrange100:
		return PenOrange100{}
	case *xpr.PenOrange200:
		return PenOrange200{}
	case *xpr.PenOrange300:
		return PenOrange300{}
	case *xpr.PenOrange400:
		return PenOrange400{}
	case *xpr.PenOrange500:
		return PenOrange500{}
	case *xpr.PenOrange600:
		return PenOrange600{}
	case *xpr.PenOrange700:
		return PenOrange700{}
	case *xpr.PenOrange800:
		return PenOrange800{}
	case *xpr.PenOrange900:
		return PenOrange900{}
	case *xpr.PenOrangeA100:
		return PenOrangeA100{}
	case *xpr.PenOrangeA200:
		return PenOrangeA200{}
	case *xpr.PenOrangeA400:
		return PenOrangeA400{}
	case *xpr.PenOrangeA700:
		return PenOrangeA700{}
	case *xpr.PenDeepOrange50:
		return PenDeepOrange50{}
	case *xpr.PenDeepOrange100:
		return PenDeepOrange100{}
	case *xpr.PenDeepOrange200:
		return PenDeepOrange200{}
	case *xpr.PenDeepOrange300:
		return PenDeepOrange300{}
	case *xpr.PenDeepOrange400:
		return PenDeepOrange400{}
	case *xpr.PenDeepOrange500:
		return PenDeepOrange500{}
	case *xpr.PenDeepOrange600:
		return PenDeepOrange600{}
	case *xpr.PenDeepOrange700:
		return PenDeepOrange700{}
	case *xpr.PenDeepOrange800:
		return PenDeepOrange800{}
	case *xpr.PenDeepOrange900:
		return PenDeepOrange900{}
	case *xpr.PenDeepOrangeA100:
		return PenDeepOrangeA100{}
	case *xpr.PenDeepOrangeA200:
		return PenDeepOrangeA200{}
	case *xpr.PenDeepOrangeA400:
		return PenDeepOrangeA400{}
	case *xpr.PenDeepOrangeA700:
		return PenDeepOrangeA700{}
	case *xpr.PenBrown50:
		return PenBrown50{}
	case *xpr.PenBrown100:
		return PenBrown100{}
	case *xpr.PenBrown200:
		return PenBrown200{}
	case *xpr.PenBrown300:
		return PenBrown300{}
	case *xpr.PenBrown400:
		return PenBrown400{}
	case *xpr.PenBrown500:
		return PenBrown500{}
	case *xpr.PenBrown600:
		return PenBrown600{}
	case *xpr.PenBrown700:
		return PenBrown700{}
	case *xpr.PenBrown800:
		return PenBrown800{}
	case *xpr.PenBrown900:
		return PenBrown900{}
	case *xpr.PenBlueGrey50:
		return PenBlueGrey50{}
	case *xpr.PenBlueGrey100:
		return PenBlueGrey100{}
	case *xpr.PenBlueGrey200:
		return PenBlueGrey200{}
	case *xpr.PenBlueGrey300:
		return PenBlueGrey300{}
	case *xpr.PenBlueGrey400:
		return PenBlueGrey400{}
	case *xpr.PenBlueGrey500:
		return PenBlueGrey500{}
	case *xpr.PenBlueGrey600:
		return PenBlueGrey600{}
	case *xpr.PenBlueGrey700:
		return PenBlueGrey700{}
	case *xpr.PenBlueGrey800:
		return PenBlueGrey800{}
	case *xpr.PenBlueGrey900:
		return PenBlueGrey900{}
	case *xpr.PenGrey50:
		return PenGrey50{}
	case *xpr.PenGrey100:
		return PenGrey100{}
	case *xpr.PenGrey200:
		return PenGrey200{}
	case *xpr.PenGrey300:
		return PenGrey300{}
	case *xpr.PenGrey400:
		return PenGrey400{}
	case *xpr.PenGrey500:
		return PenGrey500{}
	case *xpr.PenGrey600:
		return PenGrey600{}
	case *xpr.PenGrey700:
		return PenGrey700{}
	case *xpr.PenGrey800:
		return PenGrey800{}
	case *xpr.PenGrey900:
		return PenGrey900{}
	case *xpr.PenPrfPen:
		return PenPrfPen{}
	case *xpr.PenLosPen:
		return PenLosPen{}
	case *xpr.PenNew:
		var i1 []UntUntAct
		for _, cur := range X.I1 {
			i1 = append(i1, x.UntUntAct(scp, cur))
		}
		return PenNew{I0: x.ClrClrAct(scp, X.I0), I1: i1}
	case *xpr.PenRgba:
		var i4 []UntUntAct
		for _, cur := range X.I4 {
			i4 = append(i4, x.UntUntAct(scp, cur))
		}
		return PenRgba{I0: x.FltFltAct(scp, X.I0), I1: x.FltFltAct(scp, X.I1), I2: x.FltFltAct(scp, X.I2), I3: x.FltFltAct(scp, X.I3), I4: i4}
	case *xpr.PenRgb:
		var i3 []UntUntAct
		for _, cur := range X.I3 {
			i3 = append(i3, x.UntUntAct(scp, cur))
		}
		return PenRgb{I0: x.FltFltAct(scp, X.I0), I1: x.FltFltAct(scp, X.I1), I2: x.FltFltAct(scp, X.I2), I3: i3}
	case *xpr.PenHex:
		var i1 []UntUntAct
		for _, cur := range X.I1 {
			i1 = append(i1, x.UntUntAct(scp, cur))
		}
		return PenHex{I0: x.StrStrAct(scp, X.I0), I1: i1}
	case *xpr.PenPenOpa:
		return PenPenOpa{X: x.PenPenAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.PenPenInv:
		return PenPenInv{X: x.PenPenAct(scp, X.X)}
	case *xpr.PenPensPop:
		return PenPensPop{X: x.PenPensAct(scp, X.X)}
	case *xpr.PenPensDque:
		return PenPensDque{X: x.PenPensAct(scp, X.X)}
	case *xpr.PenPensDel:
		return PenPensDel{X: x.PenPensAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.PenPensAt:
		return PenPensAt{X: x.PenPensAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.PenPensFst:
		return PenPensFst{X: x.PenPensAct(scp, X.X)}
	case *xpr.PenPensMdl:
		return PenPensMdl{X: x.PenPensAct(scp, X.X)}
	case *xpr.PenPensLst:
		return PenPensLst{X: x.PenPensAct(scp, X.X)}
	}
	panic(x.Erf("PenPenAct: no action found %v", reflect.ValueOf(v).Elem().Type().Name()))
}
func (x *Actr) PenPensAct(scp *Scp, v xpr.PenPensXpr) PenPensAct {
	switch X := v.(type) {
	case *xpr.PenPensAsn:
		asnScp := scp
		if X.Depth != 0 {
			for d := X.Depth; asnScp != nil && d != 0; d-- { // recurse up scp by depth
				asnScp = asnScp.Prnt
			}
		}
		return PenPensAsn{PensScp: asnScp.PenPens(x.Txt[X.Idn.Idx:X.Idn.Lim]), X: x.PenPensAct(scp, X.X)}
	case *xpr.PenPensAcs:
		return PenPensAcs{PensScp: scp.PenPens(x.Txt[X.Trm.Idx:X.Trm.Lim])}
	case *xpr.PenPensEach:
		eachScp := NewScp(X.Scp, scp)
		return PenPensEach{X: x.PenPensAct(scp, X.X), PenScp: eachScp.PenPen(x.Txt[X.Idn.Idx:X.Idn.Lim]), Acts: x.Acts(eachScp, X.Xprs...)}
	case *xpr.PenPensPllEach:
		return PenPensPllEach{X: x.PenPensAct(scp, X.X), Txt: x.Txt, Idn: X.Idn, ActScpPrnt: scp, XprScp: X.Scp, Xprs: X.Xprs}
	case *xpr.PenNewPens:
		var i0 []PenPenAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.PenPenAct(scp, cur))
		}
		return PenNewPens{I0: i0}
	case *xpr.PenMakePens:
		return PenMakePens{I0: x.UntUntAct(scp, X.I0)}
	case *xpr.PenMakeEmpPens:
		return PenMakeEmpPens{I0: x.UntUntAct(scp, X.I0)}
	case *xpr.PenPensCpy:
		return PenPensCpy{X: x.PenPensAct(scp, X.X)}
	case *xpr.PenPensClr:
		return PenPensClr{X: x.PenPensAct(scp, X.X)}
	case *xpr.PenPensRand:
		return PenPensRand{X: x.PenPensAct(scp, X.X)}
	case *xpr.PenPensMrg:
		var i0 []PenPensAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.PenPensAct(scp, cur))
		}
		return PenPensMrg{X: x.PenPensAct(scp, X.X), I0: i0}
	case *xpr.PenPensPush:
		var i0 []PenPenAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.PenPenAct(scp, cur))
		}
		return PenPensPush{X: x.PenPensAct(scp, X.X), I0: i0}
	case *xpr.PenPensQue:
		var i0 []PenPenAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.PenPenAct(scp, cur))
		}
		return PenPensQue{X: x.PenPensAct(scp, X.X), I0: i0}
	case *xpr.PenPensIns:
		return PenPensIns{X: x.PenPensAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.PenPenAct(scp, X.I1)}
	case *xpr.PenPensUpd:
		return PenPensUpd{X: x.PenPensAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.PenPenAct(scp, X.I1)}
	case *xpr.PenPensIn:
		return PenPensIn{X: x.PenPensAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.UntUntAct(scp, X.I1)}
	case *xpr.PenPensInBnd:
		return PenPensInBnd{X: x.PenPensAct(scp, X.X), I0: x.BndBndAct(scp, X.I0)}
	case *xpr.PenPensFrom:
		return PenPensFrom{X: x.PenPensAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.PenPensTo:
		return PenPensTo{X: x.PenPensAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.PenPensRev:
		return PenPensRev{X: x.PenPensAct(scp, X.X)}
	}
	panic(x.Erf("PenPensAct: no action found %v", reflect.ValueOf(v).Elem().Type().Name()))
}
func (x *Actr) PltPltAct(scp *Scp, v xpr.PltPltXpr) PltPltAct {
	switch X := v.(type) {
	case *xpr.PltPltAsn:
		asnScp := scp
		if X.Depth != 0 {
			for d := X.Depth; asnScp != nil && d != 0; d-- { // recurse up scp by depth
				asnScp = asnScp.Prnt
			}
		}
		return PltPltAsn{PltScp: asnScp.PltPlt(x.Txt[X.Idn.Idx:X.Idn.Lim]), X: x.PltPltAct(scp, X.X)}
	case *xpr.PltPltAcs:
		return PltPltAcs{PltScp: scp.PltPlt(x.Txt[X.Trm.Idx:X.Trm.Lim])}
	case *xpr.PltPltsPop:
		return PltPltsPop{X: x.PltPltsAct(scp, X.X)}
	case *xpr.PltPltsDque:
		return PltPltsDque{X: x.PltPltsAct(scp, X.X)}
	case *xpr.PltPltsDel:
		return PltPltsDel{X: x.PltPltsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.PltPltsAt:
		return PltPltsAt{X: x.PltPltsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.PltPltsFst:
		return PltPltsFst{X: x.PltPltsAct(scp, X.X)}
	case *xpr.PltPltsMdl:
		return PltPltsMdl{X: x.PltPltsAct(scp, X.X)}
	case *xpr.PltPltsLst:
		return PltPltsLst{X: x.PltPltsAct(scp, X.X)}
	case *xpr.PltStmSho:
		return PltStmSho{X: x.PltStmAct(scp, X.X)}
	case *xpr.PltStmSiz:
		return PltStmSiz{X: x.PltStmAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.UntUntAct(scp, X.I1)}
	case *xpr.PltStmScl:
		return PltStmScl{X: x.PltStmAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.PltStmHrzScl:
		return PltStmHrzScl{X: x.PltStmAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.PltStmVrtScl:
		return PltStmVrtScl{X: x.PltStmAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.PltFltsSctrSho:
		return PltFltsSctrSho{X: x.PltFltsSctrAct(scp, X.X)}
	case *xpr.PltFltsSctrSiz:
		return PltFltsSctrSiz{X: x.PltFltsSctrAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.UntUntAct(scp, X.I1)}
	case *xpr.PltFltsSctrScl:
		return PltFltsSctrScl{X: x.PltFltsSctrAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.PltFltsSctrHrzScl:
		return PltFltsSctrHrzScl{X: x.PltFltsSctrAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.PltFltsSctrVrtScl:
		return PltFltsSctrVrtScl{X: x.PltFltsSctrAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.PltFltsSctrDistSho:
		return PltFltsSctrDistSho{X: x.PltFltsSctrDistAct(scp, X.X)}
	case *xpr.PltFltsSctrDistSiz:
		return PltFltsSctrDistSiz{X: x.PltFltsSctrDistAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.UntUntAct(scp, X.I1)}
	case *xpr.PltFltsSctrDistScl:
		return PltFltsSctrDistScl{X: x.PltFltsSctrDistAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.PltFltsSctrDistHrzScl:
		return PltFltsSctrDistHrzScl{X: x.PltFltsSctrDistAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.PltFltsSctrDistVrtScl:
		return PltFltsSctrDistVrtScl{X: x.PltFltsSctrDistAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.PltHrzSho:
		return PltHrzSho{X: x.PltHrzAct(scp, X.X)}
	case *xpr.PltHrzSiz:
		return PltHrzSiz{X: x.PltHrzAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.UntUntAct(scp, X.I1)}
	case *xpr.PltHrzScl:
		return PltHrzScl{X: x.PltHrzAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.PltHrzHrzScl:
		return PltHrzHrzScl{X: x.PltHrzAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.PltHrzVrtScl:
		return PltHrzVrtScl{X: x.PltHrzAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.PltVrtSho:
		return PltVrtSho{X: x.PltVrtAct(scp, X.X)}
	case *xpr.PltVrtSiz:
		return PltVrtSiz{X: x.PltVrtAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.UntUntAct(scp, X.I1)}
	case *xpr.PltVrtScl:
		return PltVrtScl{X: x.PltVrtAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.PltVrtHrzScl:
		return PltVrtHrzScl{X: x.PltVrtAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.PltVrtVrtScl:
		return PltVrtVrtScl{X: x.PltVrtAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.PltDpthSho:
		return PltDpthSho{X: x.PltDpthAct(scp, X.X)}
	case *xpr.PltDpthSiz:
		return PltDpthSiz{X: x.PltDpthAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.UntUntAct(scp, X.I1)}
	case *xpr.PltDpthScl:
		return PltDpthScl{X: x.PltDpthAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.PltDpthHrzScl:
		return PltDpthHrzScl{X: x.PltDpthAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.PltDpthVrtScl:
		return PltDpthVrtScl{X: x.PltDpthAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.PltPltSho:
		return PltPltSho{X: x.PltPltAct(scp, X.X)}
	case *xpr.PltPltSiz:
		return PltPltSiz{X: x.PltPltAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.UntUntAct(scp, X.I1)}
	case *xpr.PltPltScl:
		return PltPltScl{X: x.PltPltAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.PltPltHrzScl:
		return PltPltHrzScl{X: x.PltPltAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.PltPltVrtScl:
		return PltPltVrtScl{X: x.PltPltAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.PltStmAsn:
		asnScp := scp
		if X.Depth != 0 {
			for d := X.Depth; asnScp != nil && d != 0; d-- { // recurse up scp by depth
				asnScp = asnScp.Prnt
			}
		}
		return PltStmAsn{StmScp: asnScp.PltStm(x.Txt[X.Idn.Idx:X.Idn.Lim]), X: x.PltStmAct(scp, X.X)}
	case *xpr.PltStmAcs:
		return PltStmAcs{StmScp: scp.PltStm(x.Txt[X.Trm.Idx:X.Trm.Lim])}
	case *xpr.PltNewStm:
		return PltNewStm{}
	case *xpr.PltStmStm:
		var i1 []HstStmAct
		for _, cur := range X.I1 {
			i1 = append(i1, x.HstStmAct(scp, cur))
		}
		return PltStmStm{X: x.PltStmAct(scp, X.X), I0: x.PenPenAct(scp, X.I0), I1: i1}
	case *xpr.PltStmStmBnd:
		return PltStmStmBnd{X: x.PltStmAct(scp, X.X), I0: x.ClrClrAct(scp, X.I0), I1: x.PenPenAct(scp, X.I1), I2: x.HstStmAct(scp, X.I2), I3: x.HstStmAct(scp, X.I3)}
	case *xpr.PltStmCnd:
		var i1 []HstCndAct
		for _, cur := range X.I1 {
			i1 = append(i1, x.HstCndAct(scp, cur))
		}
		return PltStmCnd{X: x.PltStmAct(scp, X.X), I0: x.PenPenAct(scp, X.I0), I1: i1}
	case *xpr.PltStmHrzLn:
		var i1 []FltFltAct
		for _, cur := range X.I1 {
			i1 = append(i1, x.FltFltAct(scp, cur))
		}
		return PltStmHrzLn{X: x.PltStmAct(scp, X.X), I0: x.PenPenAct(scp, X.I0), I1: i1}
	case *xpr.PltStmVrtLn:
		var i1 []TmeTmeAct
		for _, cur := range X.I1 {
			i1 = append(i1, x.TmeTmeAct(scp, cur))
		}
		return PltStmVrtLn{X: x.PltStmAct(scp, X.X), I0: x.PenPenAct(scp, X.I0), I1: i1}
	case *xpr.PltStmHrzBnd:
		return PltStmHrzBnd{X: x.PltStmAct(scp, X.X), I0: x.ClrClrAct(scp, X.I0), I1: x.PenPenAct(scp, X.I1), I2: x.FltFltAct(scp, X.I2), I3: x.FltFltAct(scp, X.I3)}
	case *xpr.PltStmVrtBnd:
		return PltStmVrtBnd{X: x.PltStmAct(scp, X.X), I0: x.ClrClrAct(scp, X.I0), I1: x.PenPenAct(scp, X.I1), I2: x.TmeTmeAct(scp, X.I2), I3: x.TmeTmeAct(scp, X.I3)}
	case *xpr.PltStmHrzSclVal:
		return PltStmHrzSclVal{X: x.PltStmAct(scp, X.X), I0: x.TmeTmeAct(scp, X.I0)}
	case *xpr.PltStmVrtSclVal:
		return PltStmVrtSclVal{X: x.PltStmAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	case *xpr.PltFltsSctrAsn:
		asnScp := scp
		if X.Depth != 0 {
			for d := X.Depth; asnScp != nil && d != 0; d-- { // recurse up scp by depth
				asnScp = asnScp.Prnt
			}
		}
		return PltFltsSctrAsn{FltsSctrScp: asnScp.PltFltsSctr(x.Txt[X.Idn.Idx:X.Idn.Lim]), X: x.PltFltsSctrAct(scp, X.X)}
	case *xpr.PltFltsSctrAcs:
		return PltFltsSctrAcs{FltsSctrScp: scp.PltFltsSctr(x.Txt[X.Trm.Idx:X.Trm.Lim])}
	case *xpr.PltNewFltsSctr:
		return PltNewFltsSctr{}
	case *xpr.PltFltsSctrFlts:
		var i1 []FltsFltsAct
		for _, cur := range X.I1 {
			i1 = append(i1, x.FltsFltsAct(scp, cur))
		}
		return PltFltsSctrFlts{X: x.PltFltsSctrAct(scp, X.X), I0: x.ClrClrAct(scp, X.I0), I1: i1}
	case *xpr.PltFltsSctrPrfLos:
		var i2 []HstStmAct
		for _, cur := range X.I2 {
			i2 = append(i2, x.HstStmAct(scp, cur))
		}
		return PltFltsSctrPrfLos{X: x.PltFltsSctrAct(scp, X.X), I0: x.TmesTmesAct(scp, X.I0), I1: x.TmesTmesAct(scp, X.I1), I2: i2}
	case *xpr.PltFltsSctrDistAsn:
		asnScp := scp
		if X.Depth != 0 {
			for d := X.Depth; asnScp != nil && d != 0; d-- { // recurse up scp by depth
				asnScp = asnScp.Prnt
			}
		}
		return PltFltsSctrDistAsn{FltsSctrDistScp: asnScp.PltFltsSctrDist(x.Txt[X.Idn.Idx:X.Idn.Lim]), X: x.PltFltsSctrDistAct(scp, X.X)}
	case *xpr.PltFltsSctrDistAcs:
		return PltFltsSctrDistAcs{FltsSctrDistScp: scp.PltFltsSctrDist(x.Txt[X.Trm.Idx:X.Trm.Lim])}
	case *xpr.PltNewFltsSctrDist:
		return PltNewFltsSctrDist{}
	case *xpr.PltFltsSctrDistFlts:
		var i2 []FltsFltsAct
		for _, cur := range X.I2 {
			i2 = append(i2, x.FltsFltsAct(scp, cur))
		}
		return PltFltsSctrDistFlts{X: x.PltFltsSctrDistAct(scp, X.X), I0: x.ClrClrAct(scp, X.I0), I1: x.UntUntAct(scp, X.I1), I2: i2}
	case *xpr.PltHrzAsn:
		asnScp := scp
		if X.Depth != 0 {
			for d := X.Depth; asnScp != nil && d != 0; d-- { // recurse up scp by depth
				asnScp = asnScp.Prnt
			}
		}
		return PltHrzAsn{HrzScp: asnScp.PltHrz(x.Txt[X.Idn.Idx:X.Idn.Lim]), X: x.PltHrzAct(scp, X.X)}
	case *xpr.PltHrzAcs:
		return PltHrzAcs{HrzScp: scp.PltHrz(x.Txt[X.Trm.Idx:X.Trm.Lim])}
	case *xpr.PltNewHrz:
		var i0 []PltPltAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.PltPltAct(scp, cur))
		}
		return PltNewHrz{I0: i0}
	case *xpr.PltHrzPlt:
		var i0 []PltPltAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.PltPltAct(scp, cur))
		}
		return PltHrzPlt{X: x.PltHrzAct(scp, X.X), I0: i0}
	case *xpr.PltVrtAsn:
		asnScp := scp
		if X.Depth != 0 {
			for d := X.Depth; asnScp != nil && d != 0; d-- { // recurse up scp by depth
				asnScp = asnScp.Prnt
			}
		}
		return PltVrtAsn{VrtScp: asnScp.PltVrt(x.Txt[X.Idn.Idx:X.Idn.Lim]), X: x.PltVrtAct(scp, X.X)}
	case *xpr.PltVrtAcs:
		return PltVrtAcs{VrtScp: scp.PltVrt(x.Txt[X.Trm.Idx:X.Trm.Lim])}
	case *xpr.PltNewVrt:
		var i0 []PltPltAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.PltPltAct(scp, cur))
		}
		return PltNewVrt{I0: i0}
	case *xpr.PltVrtPlt:
		var i0 []PltPltAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.PltPltAct(scp, cur))
		}
		return PltVrtPlt{X: x.PltVrtAct(scp, X.X), I0: i0}
	case *xpr.PltDpthAsn:
		asnScp := scp
		if X.Depth != 0 {
			for d := X.Depth; asnScp != nil && d != 0; d-- { // recurse up scp by depth
				asnScp = asnScp.Prnt
			}
		}
		return PltDpthAsn{DpthScp: asnScp.PltDpth(x.Txt[X.Idn.Idx:X.Idn.Lim]), X: x.PltDpthAct(scp, X.X)}
	case *xpr.PltDpthAcs:
		return PltDpthAcs{DpthScp: scp.PltDpth(x.Txt[X.Trm.Idx:X.Trm.Lim])}
	case *xpr.PltNewDpth:
		var i0 []PltPltAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.PltPltAct(scp, cur))
		}
		return PltNewDpth{I0: i0}
	case *xpr.PltDpthPlt:
		var i0 []PltPltAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.PltPltAct(scp, cur))
		}
		return PltDpthPlt{X: x.PltDpthAct(scp, X.X), I0: i0}
	}
	panic(x.Erf("PltPltAct: no action found %v", reflect.ValueOf(v).Elem().Type().Name()))
}
func (x *Actr) PltPltsAct(scp *Scp, v xpr.PltPltsXpr) PltPltsAct {
	switch X := v.(type) {
	case *xpr.PltPltsAsn:
		asnScp := scp
		if X.Depth != 0 {
			for d := X.Depth; asnScp != nil && d != 0; d-- { // recurse up scp by depth
				asnScp = asnScp.Prnt
			}
		}
		return PltPltsAsn{PltsScp: asnScp.PltPlts(x.Txt[X.Idn.Idx:X.Idn.Lim]), X: x.PltPltsAct(scp, X.X)}
	case *xpr.PltPltsAcs:
		return PltPltsAcs{PltsScp: scp.PltPlts(x.Txt[X.Trm.Idx:X.Trm.Lim])}
	case *xpr.PltPltsEach:
		eachScp := NewScp(X.Scp, scp)
		return PltPltsEach{X: x.PltPltsAct(scp, X.X), PltScp: eachScp.PltPlt(x.Txt[X.Idn.Idx:X.Idn.Lim]), Acts: x.Acts(eachScp, X.Xprs...)}
	case *xpr.PltPltsPllEach:
		return PltPltsPllEach{X: x.PltPltsAct(scp, X.X), Txt: x.Txt, Idn: X.Idn, ActScpPrnt: scp, XprScp: X.Scp, Xprs: X.Xprs}
	case *xpr.PltHrzPltsGet:
		return PltHrzPltsGet{X: x.PltHrzAct(scp, X.X)}
	case *xpr.PltVrtPltsGet:
		return PltVrtPltsGet{X: x.PltVrtAct(scp, X.X)}
	case *xpr.PltDpthPltsGet:
		return PltDpthPltsGet{X: x.PltDpthAct(scp, X.X)}
	case *xpr.PltNewPlts:
		var i0 []PltPltAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.PltPltAct(scp, cur))
		}
		return PltNewPlts{I0: i0}
	case *xpr.PltMakePlts:
		return PltMakePlts{I0: x.UntUntAct(scp, X.I0)}
	case *xpr.PltMakeEmpPlts:
		return PltMakeEmpPlts{I0: x.UntUntAct(scp, X.I0)}
	case *xpr.PltPltsCpy:
		return PltPltsCpy{X: x.PltPltsAct(scp, X.X)}
	case *xpr.PltPltsClr:
		return PltPltsClr{X: x.PltPltsAct(scp, X.X)}
	case *xpr.PltPltsRand:
		return PltPltsRand{X: x.PltPltsAct(scp, X.X)}
	case *xpr.PltPltsMrg:
		var i0 []PltPltsAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.PltPltsAct(scp, cur))
		}
		return PltPltsMrg{X: x.PltPltsAct(scp, X.X), I0: i0}
	case *xpr.PltPltsPush:
		var i0 []PltPltAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.PltPltAct(scp, cur))
		}
		return PltPltsPush{X: x.PltPltsAct(scp, X.X), I0: i0}
	case *xpr.PltPltsQue:
		var i0 []PltPltAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.PltPltAct(scp, cur))
		}
		return PltPltsQue{X: x.PltPltsAct(scp, X.X), I0: i0}
	case *xpr.PltPltsIns:
		return PltPltsIns{X: x.PltPltsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.PltPltAct(scp, X.I1)}
	case *xpr.PltPltsUpd:
		return PltPltsUpd{X: x.PltPltsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.PltPltAct(scp, X.I1)}
	case *xpr.PltPltsIn:
		return PltPltsIn{X: x.PltPltsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0), I1: x.UntUntAct(scp, X.I1)}
	case *xpr.PltPltsInBnd:
		return PltPltsInBnd{X: x.PltPltsAct(scp, X.X), I0: x.BndBndAct(scp, X.I0)}
	case *xpr.PltPltsFrom:
		return PltPltsFrom{X: x.PltPltsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.PltPltsTo:
		return PltPltsTo{X: x.PltPltsAct(scp, X.X), I0: x.UntUntAct(scp, X.I0)}
	case *xpr.PltPltsRev:
		return PltPltsRev{X: x.PltPltsAct(scp, X.X)}
	}
	panic(x.Erf("PltPltsAct: no action found %v", reflect.ValueOf(v).Elem().Type().Name()))
}
func (x *Actr) PltTmeAxisXAct(scp *Scp, v xpr.PltTmeAxisXXpr) PltTmeAxisXAct {
	switch X := v.(type) {
	case *xpr.PltTmeAxisXVis:
		return PltTmeAxisXVis{X: x.PltTmeAxisXAct(scp, X.X), I0: x.BolBolAct(scp, X.I0)}
	case *xpr.PltStmX:
		return PltStmX{X: x.PltStmAct(scp, X.X)}
	}
	panic(x.Erf("PltTmeAxisXAct: no action found %v", reflect.ValueOf(v).Elem().Type().Name()))
}
func (x *Actr) PltFltAxisYAct(scp *Scp, v xpr.PltFltAxisYXpr) PltFltAxisYAct {
	switch X := v.(type) {
	case *xpr.PltFltsSctrYGet:
		return PltFltsSctrYGet{X: x.PltFltsSctrAct(scp, X.X)}
	case *xpr.PltFltAxisYVis:
		return PltFltAxisYVis{X: x.PltFltAxisYAct(scp, X.X), I0: x.BolBolAct(scp, X.I0)}
	case *xpr.PltStmY:
		return PltStmY{X: x.PltStmAct(scp, X.X)}
	}
	panic(x.Erf("PltFltAxisYAct: no action found %v", reflect.ValueOf(v).Elem().Type().Name()))
}
func (x *Actr) PltStmAct(scp *Scp, v xpr.PltStmXpr) PltStmAct {
	switch X := v.(type) {
	case *xpr.PltStmAsn:
		asnScp := scp
		if X.Depth != 0 {
			for d := X.Depth; asnScp != nil && d != 0; d-- { // recurse up scp by depth
				asnScp = asnScp.Prnt
			}
		}
		return PltStmAsn{StmScp: asnScp.PltStm(x.Txt[X.Idn.Idx:X.Idn.Lim]), X: x.PltStmAct(scp, X.X)}
	case *xpr.PltStmAcs:
		return PltStmAcs{StmScp: scp.PltStm(x.Txt[X.Trm.Idx:X.Trm.Lim])}
	case *xpr.PltNewStm:
		return PltNewStm{}
	case *xpr.PltStmStm:
		var i1 []HstStmAct
		for _, cur := range X.I1 {
			i1 = append(i1, x.HstStmAct(scp, cur))
		}
		return PltStmStm{X: x.PltStmAct(scp, X.X), I0: x.PenPenAct(scp, X.I0), I1: i1}
	case *xpr.PltStmStmBnd:
		return PltStmStmBnd{X: x.PltStmAct(scp, X.X), I0: x.ClrClrAct(scp, X.I0), I1: x.PenPenAct(scp, X.I1), I2: x.HstStmAct(scp, X.I2), I3: x.HstStmAct(scp, X.I3)}
	case *xpr.PltStmCnd:
		var i1 []HstCndAct
		for _, cur := range X.I1 {
			i1 = append(i1, x.HstCndAct(scp, cur))
		}
		return PltStmCnd{X: x.PltStmAct(scp, X.X), I0: x.PenPenAct(scp, X.I0), I1: i1}
	case *xpr.PltStmHrzLn:
		var i1 []FltFltAct
		for _, cur := range X.I1 {
			i1 = append(i1, x.FltFltAct(scp, cur))
		}
		return PltStmHrzLn{X: x.PltStmAct(scp, X.X), I0: x.PenPenAct(scp, X.I0), I1: i1}
	case *xpr.PltStmVrtLn:
		var i1 []TmeTmeAct
		for _, cur := range X.I1 {
			i1 = append(i1, x.TmeTmeAct(scp, cur))
		}
		return PltStmVrtLn{X: x.PltStmAct(scp, X.X), I0: x.PenPenAct(scp, X.I0), I1: i1}
	case *xpr.PltStmHrzBnd:
		return PltStmHrzBnd{X: x.PltStmAct(scp, X.X), I0: x.ClrClrAct(scp, X.I0), I1: x.PenPenAct(scp, X.I1), I2: x.FltFltAct(scp, X.I2), I3: x.FltFltAct(scp, X.I3)}
	case *xpr.PltStmVrtBnd:
		return PltStmVrtBnd{X: x.PltStmAct(scp, X.X), I0: x.ClrClrAct(scp, X.I0), I1: x.PenPenAct(scp, X.I1), I2: x.TmeTmeAct(scp, X.I2), I3: x.TmeTmeAct(scp, X.I3)}
	case *xpr.PltStmHrzSclVal:
		return PltStmHrzSclVal{X: x.PltStmAct(scp, X.X), I0: x.TmeTmeAct(scp, X.I0)}
	case *xpr.PltStmVrtSclVal:
		return PltStmVrtSclVal{X: x.PltStmAct(scp, X.X), I0: x.FltFltAct(scp, X.I0)}
	}
	panic(x.Erf("PltStmAct: no action found %v", reflect.ValueOf(v).Elem().Type().Name()))
}
func (x *Actr) PltFltsSctrAct(scp *Scp, v xpr.PltFltsSctrXpr) PltFltsSctrAct {
	switch X := v.(type) {
	case *xpr.PltFltsSctrAsn:
		asnScp := scp
		if X.Depth != 0 {
			for d := X.Depth; asnScp != nil && d != 0; d-- { // recurse up scp by depth
				asnScp = asnScp.Prnt
			}
		}
		return PltFltsSctrAsn{FltsSctrScp: asnScp.PltFltsSctr(x.Txt[X.Idn.Idx:X.Idn.Lim]), X: x.PltFltsSctrAct(scp, X.X)}
	case *xpr.PltFltsSctrAcs:
		return PltFltsSctrAcs{FltsSctrScp: scp.PltFltsSctr(x.Txt[X.Trm.Idx:X.Trm.Lim])}
	case *xpr.PltNewFltsSctr:
		return PltNewFltsSctr{}
	case *xpr.PltFltsSctrFlts:
		var i1 []FltsFltsAct
		for _, cur := range X.I1 {
			i1 = append(i1, x.FltsFltsAct(scp, cur))
		}
		return PltFltsSctrFlts{X: x.PltFltsSctrAct(scp, X.X), I0: x.ClrClrAct(scp, X.I0), I1: i1}
	case *xpr.PltFltsSctrPrfLos:
		var i2 []HstStmAct
		for _, cur := range X.I2 {
			i2 = append(i2, x.HstStmAct(scp, cur))
		}
		return PltFltsSctrPrfLos{X: x.PltFltsSctrAct(scp, X.X), I0: x.TmesTmesAct(scp, X.I0), I1: x.TmesTmesAct(scp, X.I1), I2: i2}
	}
	panic(x.Erf("PltFltsSctrAct: no action found %v", reflect.ValueOf(v).Elem().Type().Name()))
}
func (x *Actr) PltFltsSctrDistAct(scp *Scp, v xpr.PltFltsSctrDistXpr) PltFltsSctrDistAct {
	switch X := v.(type) {
	case *xpr.PltFltsSctrDistAsn:
		asnScp := scp
		if X.Depth != 0 {
			for d := X.Depth; asnScp != nil && d != 0; d-- { // recurse up scp by depth
				asnScp = asnScp.Prnt
			}
		}
		return PltFltsSctrDistAsn{FltsSctrDistScp: asnScp.PltFltsSctrDist(x.Txt[X.Idn.Idx:X.Idn.Lim]), X: x.PltFltsSctrDistAct(scp, X.X)}
	case *xpr.PltFltsSctrDistAcs:
		return PltFltsSctrDistAcs{FltsSctrDistScp: scp.PltFltsSctrDist(x.Txt[X.Trm.Idx:X.Trm.Lim])}
	case *xpr.PltNewFltsSctrDist:
		return PltNewFltsSctrDist{}
	case *xpr.PltFltsSctrDistFlts:
		var i2 []FltsFltsAct
		for _, cur := range X.I2 {
			i2 = append(i2, x.FltsFltsAct(scp, cur))
		}
		return PltFltsSctrDistFlts{X: x.PltFltsSctrDistAct(scp, X.X), I0: x.ClrClrAct(scp, X.I0), I1: x.UntUntAct(scp, X.I1), I2: i2}
	}
	panic(x.Erf("PltFltsSctrDistAct: no action found %v", reflect.ValueOf(v).Elem().Type().Name()))
}
func (x *Actr) PltHrzAct(scp *Scp, v xpr.PltHrzXpr) PltHrzAct {
	switch X := v.(type) {
	case *xpr.PltHrzAsn:
		asnScp := scp
		if X.Depth != 0 {
			for d := X.Depth; asnScp != nil && d != 0; d-- { // recurse up scp by depth
				asnScp = asnScp.Prnt
			}
		}
		return PltHrzAsn{HrzScp: asnScp.PltHrz(x.Txt[X.Idn.Idx:X.Idn.Lim]), X: x.PltHrzAct(scp, X.X)}
	case *xpr.PltHrzAcs:
		return PltHrzAcs{HrzScp: scp.PltHrz(x.Txt[X.Trm.Idx:X.Trm.Lim])}
	case *xpr.PltNewHrz:
		var i0 []PltPltAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.PltPltAct(scp, cur))
		}
		return PltNewHrz{I0: i0}
	case *xpr.PltHrzPlt:
		var i0 []PltPltAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.PltPltAct(scp, cur))
		}
		return PltHrzPlt{X: x.PltHrzAct(scp, X.X), I0: i0}
	}
	panic(x.Erf("PltHrzAct: no action found %v", reflect.ValueOf(v).Elem().Type().Name()))
}
func (x *Actr) PltVrtAct(scp *Scp, v xpr.PltVrtXpr) PltVrtAct {
	switch X := v.(type) {
	case *xpr.PltVrtAsn:
		asnScp := scp
		if X.Depth != 0 {
			for d := X.Depth; asnScp != nil && d != 0; d-- { // recurse up scp by depth
				asnScp = asnScp.Prnt
			}
		}
		return PltVrtAsn{VrtScp: asnScp.PltVrt(x.Txt[X.Idn.Idx:X.Idn.Lim]), X: x.PltVrtAct(scp, X.X)}
	case *xpr.PltVrtAcs:
		return PltVrtAcs{VrtScp: scp.PltVrt(x.Txt[X.Trm.Idx:X.Trm.Lim])}
	case *xpr.PltNewVrt:
		var i0 []PltPltAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.PltPltAct(scp, cur))
		}
		return PltNewVrt{I0: i0}
	case *xpr.PltVrtPlt:
		var i0 []PltPltAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.PltPltAct(scp, cur))
		}
		return PltVrtPlt{X: x.PltVrtAct(scp, X.X), I0: i0}
	}
	panic(x.Erf("PltVrtAct: no action found %v", reflect.ValueOf(v).Elem().Type().Name()))
}
func (x *Actr) PltDpthAct(scp *Scp, v xpr.PltDpthXpr) PltDpthAct {
	switch X := v.(type) {
	case *xpr.PltDpthAsn:
		asnScp := scp
		if X.Depth != 0 {
			for d := X.Depth; asnScp != nil && d != 0; d-- { // recurse up scp by depth
				asnScp = asnScp.Prnt
			}
		}
		return PltDpthAsn{DpthScp: asnScp.PltDpth(x.Txt[X.Idn.Idx:X.Idn.Lim]), X: x.PltDpthAct(scp, X.X)}
	case *xpr.PltDpthAcs:
		return PltDpthAcs{DpthScp: scp.PltDpth(x.Txt[X.Trm.Idx:X.Trm.Lim])}
	case *xpr.PltNewDpth:
		var i0 []PltPltAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.PltPltAct(scp, cur))
		}
		return PltNewDpth{I0: i0}
	case *xpr.PltDpthPlt:
		var i0 []PltPltAct
		for _, cur := range X.I0 {
			i0 = append(i0, x.PltPltAct(scp, cur))
		}
		return PltDpthPlt{X: x.PltDpthAct(scp, X.X), I0: i0}
	}
	panic(x.Erf("PltDpthAct: no action found %v", reflect.ValueOf(v).Elem().Type().Name()))
}
func (x *Actr) SysMuAct(scp *Scp, v xpr.SysMuXpr) SysMuAct {
	switch X := v.(type) {
	case *xpr.SysMuAsn:
		asnScp := scp
		if X.Depth != 0 {
			for d := X.Depth; asnScp != nil && d != 0; d-- { // recurse up scp by depth
				asnScp = asnScp.Prnt
			}
		}
		return SysMuAsn{MuScp: asnScp.SysMu(x.Txt[X.Idn.Idx:X.Idn.Lim]), X: x.SysMuAct(scp, X.X)}
	case *xpr.SysMuAcs:
		return SysMuAcs{MuScp: scp.SysMu(x.Txt[X.Trm.Idx:X.Trm.Lim])}
	case *xpr.SysNewMu:
		return SysNewMu{}
	case *xpr.SysMuLck:
		return SysMuLck{X: x.SysMuAct(scp, X.X)}
	case *xpr.SysMuUlck:
		return SysMuUlck{X: x.SysMuAct(scp, X.X)}
	}
	panic(x.Erf("SysMuAct: no action found %v", reflect.ValueOf(v).Elem().Type().Name()))
}
