package atr

type (
	Atr uint64
)

const (
	Lit Atr = 1 << iota
	Arr
	Rng
	Fbr
	Idn
	Rel
	Sgn
	Ari
	Sel
	Srt
	Grp

	Cfg
	Jsn

	Pkt // packet: rlt rx/tx pkt
	Ana // system analytics: ana, oan, dom, instr, inrvl, side, stm, cnd, stgy...

	Trm // lvls
	Xpr
	Act
	Tst
	Test
	TestXpr
	TestAct

	Scp
	Get    // surface struct field get in language
	SetGet // surface struct field set and get in language
	Bsc
	Bse
	Struct
	Stm
	Cnd
	Stgy
	Ui
	Srch
	Key
	Flg
	Iota
	Prnt
	Slf
	Empty
	Rxs
	Dlt

	StrWrt // StrWrt/String
	BytWrt // BytWrt/Bytes
	TstLrg // large test values (flts for pll testing)
	TstSkp
	BytSkp // For Perf.Pth
	LitSkp // for Perf.Pth
	StrSkp
	EqlSkp
	BqSkp
	PrntSkp // for Hst
	SelSkp  // for PrtArrFldSel
	TstZeroSkp
	FldSkp

	PrtFldSet
	PrtActFn

	None Atr = 0

	BytLitSkp            = BytSkp | LitSkp
	BytStrSkp            = BytSkp | StrSkp
	BytStrEqlSkp         = BytSkp | StrSkp | EqlSkp
	BytLitEqlSelSkp      = BytSkp | LitSkp | EqlSkp | SelSkp
	BytLitStrSkp         = BytSkp | LitSkp | StrSkp
	BytLitStrEqlSkp      = BytSkp | LitSkp | StrSkp | EqlSkp
	BytLitStrEqlSelSkp   = BytSkp | LitSkp | StrSkp | EqlSkp | SelSkp
	BytLitStrEqlBqSkp    = BytSkp | LitSkp | StrSkp | EqlSkp | BqSkp
	BytLitStrEqlBqTstSkp = BytSkp | LitSkp | StrSkp | EqlSkp | BqSkp | TstSkp
	BytLitStrEqlBqSelSkp = BytSkp | LitSkp | StrSkp | EqlSkp | BqSkp | SelSkp

	Ext    = Xpr | Act
	Lng    = Trm | Xpr | Act
	LngScp = Trm | Xpr | Act | Scp

	LngTest      = Trm | Xpr | Act | TestXpr | TestAct
	Typ          = Trm | Xpr | Act | Scp | Tst | Test
	TypLng       = Trm | Xpr | Act | Scp | Tst
	TypLngStruct = TypLng | Struct

	TypRoot   = Trm | Xpr | Act | Tst
	TypBsc    = Typ | Bsc | Idn | Lit | StrWrt | BytWrt | Cfg
	TypRng    = Typ | Rng | Lit | Struct | Cfg | StrWrt
	Num       = Idn | Rel | Ari
	NumSgn    = Num | Sgn
	TypArr    = Typ | Arr
	TypBscArr = TypArr | Lit | StrWrt | BytWrt
	TypIdn    = Trm | Xpr

	TypStr = TypBsc | Rel
	TypBol = TypBsc
	TypFlt = TypBsc | NumSgn | TstLrg // TstLrg for flts pll testing
	TypUnt = TypBsc | Num
	TypInt = TypBsc | NumSgn
	TypTme = TypBsc | NumSgn | Pkt
	TypBnd = TypBsc | Struct | Pkt

	AnaInstr     = Struct | Tst | Test | Idn
	AnaStm       = Struct | Tst | Idn
	AnaTrd       = LngScp | Struct | Idn | Tst | Pkt
	AnaStgyPrfm  = LngScp | Struct | Idn | Tst
	AnaPortPrfm  = LngScp | Struct | Idn | Tst
	AnaTmeIdx    = Struct | Tst | Idn | Pkt
	AnaTmeIdxs   = Arr | Tst | Test
	AnaTmeFlt    = Struct | Tst | Idn | Pkt
	AnaTmeFlts   = Struct | Tst | Idn | Pkt
	HstTrdStmSeg = Typ&^Test | Struct | Idn
	HstTrdsStm   = Typ&^Test | Struct | Idn | PrntSkp

	TypAnaBse    = Struct | Ana | Bse
	TypAnaStruct = Trm | Xpr | Act | Scp | Ana | Struct | Tst
	TypAnaIfc    = Trm | Xpr | Act | Scp | Ana | Idn | Tst | Test // Ifc defined in AnaPth embedded

	TypAnaStm  = Struct | Ana | Stm | Idn | Tst
	TypAnaCnd  = Struct | Ana | Cnd | Idn | Tst
	TypAnaStgy = Struct | Ana | Stgy | Idn | Tst
	TypAnaPort = LngScp | Struct | Idn | Tst

	TypUi       = Trm | Xpr | Act | Scp | Ui | BytStrEqlSkp // BytStrEqlSkp FOR PrtArr
	TypUiStruct = TypUi | Struct
	// TypUiIfc           = TypUi
	TypUiStructSupport = TypUiStruct &^ Scp
)

func (x Atr) IsNone() bool { return x == None }

func (x Atr) IsStr() bool { return x == TypStr }
func (x Atr) IsBol() bool { return x == TypBol }
func (x Atr) IsFlt() bool { return x == TypFlt }
func (x Atr) IsUnt() bool { return x == TypUnt }
func (x Atr) IsInt() bool { return x == TypInt }
func (x Atr) IsTme() bool { return x == TypTme }
func (x Atr) IsBnd() bool { return x == TypBnd }

func (x Atr) IsNum() bool {
	return x == TypFlt || x == TypUnt || x == TypInt || x == TypTme || x == TypBnd
}
func (x Atr) IsCmp() bool { return x&Idn == Idn || x&Rel == Rel }

func (x Atr) IsLit() bool { return x&Lit == Lit }
func (x Atr) IsArr() bool { return x&Arr == Arr }
func (x Atr) IsRng() bool { return x&Rng == Rng }
func (x Atr) IsFbr() bool { return x&Fbr == Fbr }
func (x Atr) IsIdn() bool { return x&Idn == Idn }
func (x Atr) IsRel() bool { return x&Rel == Rel }
func (x Atr) IsSgn() bool { return x&Sgn == Sgn }
func (x Atr) IsAri() bool { return x&Ari == Ari }
func (x Atr) IsSel() bool { return x&Sel == Sel }
func (x Atr) IsSrt() bool { return x&Srt == Srt }
func (x Atr) IsGrp() bool { return x&Grp == Grp }

func (x Atr) IsCfg() bool { return x&Cfg == Cfg }
func (x Atr) IsJsn() bool { return x&Jsn == Jsn }
func (x Atr) IsPkt() bool { return x&Pkt == Pkt }
func (x Atr) IsAna() bool { return x&Ana == Ana }

func (x Atr) IsTrm() bool     { return x&Trm == Trm }
func (x Atr) IsXpr() bool     { return x&Xpr == Xpr }
func (x Atr) IsAct() bool     { return x&Act == Act }
func (x Atr) IsTst() bool     { return x&Tst == Tst }
func (x Atr) IsTest() bool    { return x&Test == Test }
func (x Atr) IsTestXpr() bool { return x&TestXpr == TestXpr }
func (x Atr) IsTestAct() bool { return x&TestAct == TestAct }

func (x Atr) IsScp() bool    { return x&Scp == Scp }
func (x Atr) IsGet() bool    { return x&Get == Get }
func (x Atr) IsSetGet() bool { return x&SetGet == SetGet }
func (x Atr) IsBsc() bool    { return x&Bsc == Bsc }
func (x Atr) IsBse() bool    { return x&Bse == Bse }
func (x Atr) IsStm() bool    { return x&Stm == Stm }
func (x Atr) IsCnd() bool    { return x&Cnd == Cnd }
func (x Atr) IsStgy() bool   { return x&Stgy == Stgy }
func (x Atr) IsUi() bool     { return x&Ui == Ui }
func (x Atr) IsSrch() bool   { return x&Srch == Srch }
func (x Atr) IsKey() bool    { return x&Key == Key }
func (x Atr) IsFlg() bool    { return x&Flg == Flg }
func (x Atr) IsIota() bool   { return x&Iota == Iota }
func (x Atr) IsPrnt() bool   { return x&Prnt == Prnt }
func (x Atr) IsSlf() bool    { return x&Slf == Slf }
func (x Atr) IsEmpty() bool  { return x&Empty == Empty }
func (x Atr) IsRxs() bool    { return x&Rxs == Rxs }
func (x Atr) IsDlt() bool    { return x&Dlt == Dlt }

func (x Atr) IsTstLrg() bool     { return x&TstLrg == TstLrg }
func (x Atr) IsStruct() bool     { return x&Struct == Struct }
func (x Atr) IsStrWrt() bool     { return x&StrWrt == StrWrt }
func (x Atr) IsBytWrt() bool     { return x&BytWrt == BytWrt }
func (x Atr) IsTstSkp() bool     { return x&TstSkp == TstSkp }
func (x Atr) IsBytSkp() bool     { return x&BytSkp == BytSkp }
func (x Atr) IsLitSkp() bool     { return x&LitSkp == LitSkp }
func (x Atr) IsStrSkp() bool     { return x&StrSkp == StrSkp }
func (x Atr) IsEqlSkp() bool     { return x&EqlSkp == EqlSkp }
func (x Atr) IsBqSkp() bool      { return x&BqSkp == BqSkp }
func (x Atr) IsPrntSkp() bool    { return x&PrntSkp == PrntSkp }
func (x Atr) IsSelSkp() bool     { return x&SelSkp == SelSkp }
func (x Atr) IsTstZeroSkp() bool { return x&TstZeroSkp == TstZeroSkp }
func (x Atr) IsFldSkp() bool     { return x&FldSkp == FldSkp }
func (x Atr) IsPrtFldSet() bool  { return x&PrtFldSet == PrtFldSet }
func (x Atr) IsPrtActFn() bool   { return x&PrtActFn == PrtActFn }
