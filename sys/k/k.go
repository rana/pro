package k

const ( // golang keywords
	Func = "func"
	Map  = "map"
	Make = "make"
)

const (
	Cor    = "cor" // lvls
	Trm    = "trm"
	Xpr    = "xpr"
	Act    = "act"
	Size   = "size"
	Zero   = "zero" // bol cnsts
	Fls    = "fls"
	Tru    = "tru"
	False  = "false"
	True   = "true"
	One    = "one" // Num cnsts
	Hndrd  = "hndrd"
	NegOne = "negOne"
	Min    = "min"
	Max    = "max"
	MinMax = "minMax"
	Tiny   = "tiny"
	Second = "second" // tme cnsts
	Minute = "minute"
	Hour   = "hour"
	Day    = "day"
	Week   = "week"
	Month  = "month"
	Year   = "year"
	Date   = "date"
	Time   = "time"
	Dte    = "dte"
	Zone   = "zone"
	Empty  = "empty" // str cnst
	Ext    = "ext"
	Str    = "str" // basics
	Bol    = "bol"
	Unt    = "unt"
	Int    = "int"
	Flt    = "flt"
	Tme    = "tme"
	Bnd    = "bnd"
	Strs   = "strs" // basic arrs
	Bols   = "bols"
	Unts   = "unts"
	Ints   = "ints"
	Flts   = "flts"
	Tmes   = "tmes"
	Bnds   = "bnds"
	Jsn    = "jsn"
	Idx    = "idx"
	Lim    = "lim"
	Prs    = "prs"
	Xprs   = "xprs"
	Acts   = "acts"

	Test      = "test"
	Bsc       = "bsc"
	Sys       = "sys" // dir/pkg
	Spc       = "spc"
	Pnt       = "pnt"
	Pnts      = "pnts"
	Arr       = "arr"
	Fbr       = "fbr"
	Wve       = "wve"
	Scn       = "scn"
	Scnr      = "scnr"
	Trmr      = "trmr"
	Xprr      = "xprr"
	Actr      = "actr"
	Cfgr      = "cfgr"
	Jsnr      = "jsnr"
	Knd       = "knd"
	Scp       = "scp"
	Txt       = "txt"
	TrdState  = "trdState"
	TrdRsn    = "trdRsn"
	TrdRsnCls = "trdRsnCls"
	TrdRsnOpn = "trdRsnOpn"
	Bse       = "bse"
	PthBse    = "pthBse"
	VisBse    = "visBse"
	AnaBse    = "anaBse"
	AnaVisBse = "anaVisBse"
	FbrBse    = "fbrBse"
	Tst       = "tst"
	Trc       = "trc"
	TrcOpt    = "trcOpt"
	Name      = "name"
	Title     = "title"
	Lbl       = "lbl"

	Txr          = "txr"
	Hst          = "hst"
	HstBse       = "hstBse"
	Rlt          = "rlt"
	RltBse       = "rltBse"
	Dur          = "dur"
	Rng          = "rng"
	Rngs         = "rngs"
	Tic          = "tic"
	Tics         = "tics"
	TmeFlt       = "tmeFlt"
	TmeFlts      = "tmeFlts"
	TmeIdx       = "tmeIdx"
	TmeIdxs      = "tmeIdxs"
	TmeRng       = "tmeRng"
	TmeRngs      = "tmeRngs"
	Ana          = "ana"
	Anar         = "anar"
	PrvBse       = "prvBse"
	Prv          = "prv"
	Prvs         = "prvs"
	Instr        = "instr"
	Instrs       = "instrs"
	InstrBse     = "instrBse"
	InstrFbr     = "instrFbr"
	InstrFbrs    = "instrFbrs"
	InstrFbrBse  = "instrFbrBse"
	Inrvl        = "inrvl"
	Inrvls       = "inrvls"
	InrvlBse     = "inrvlBse"
	InrvlFbr     = "inrvlFbr"
	Side         = "side"
	Sides        = "sides"
	SideFbr      = "sideFbr"
	Stm          = "stm"
	Stms         = "stms"
	StmFbr       = "stmFbr"
	StmWve       = "stmWve"
	Cnd          = "cnd"
	Cnds         = "cnds"
	CndFbr       = "cndFbr"
	Stgy         = "stgy"
	Stgys        = "stgys"
	Ftr          = "ftr"
	Ftrs         = "ftrs"
	Splt         = "splt"
	Splts        = "splts"
	Prcp         = "prcp"
	PrcpStm      = "prcpStm"
	PrcpSplt     = "prcpSplt"
	StmSplt      = "stmSplt"
	StmSpltSide  = "stmSpltSide"
	StmFbrSplt   = "stmFbrSplt"
	StmWveSplt   = "stmWveSplt"
	SpltStgy     = "spltStgy"
	Tune         = "tune"
	TuneStm      = "tuneStm"
	TuneSacf     = "tuneSacf"
	TuneLosLim   = "tuneLosLim"
	TuneSacfTil  = "tuneSacfTil"
	Pnl          = "pnl"
	StgyPnl      = "stgyPnl"
	StgyPnls     = "stgyPnls"
	StgyFbr      = "stgyFbr"
	SrchTrdStm   = "srchTrdStm"
	Chan         = "chan"
	FltsSctr     = "fltsSctr"
	FltsSctrDist = "fltsSctrDist"
	TmeAxisX     = "tmeAxisX"
	FltAxisY     = "fltAxisY"
	GenPth       = "genPth"
	Dlt          = "dlt"
)

const ( // TrcOpt
	Runr     = "runr"
	TicrRx   = "ticrRx"
	TicrTx   = "ticrTx"
	RltInstr = "rltInstr"
	RltInrvl = "rltInrvl"
	RltSide  = "rltSide"
	RltStm   = "rltStm"
	RltCnd   = "rltCnd"
	RltStgy  = "rltStgy"
	RltPort  = "rltPort"
	RltPrfm  = "rltPrfm"
	HstInstr = "hstInstr"
	HstInrvl = "hstInrvl"
	HstSide  = "hstSide"
	HstStm   = "hstStm"
	HstCnd   = "hstCnd"
	HstStgy  = "hstStgy"
	HstPort  = "hstPort"
	HstPrfm  = "hstPrfm"

	HstInstrFbr = "hstInstrFbr"
	HstInrvlFbr = "hstInrvlFbr"
	HstSideFbr  = "hstSideFbr"
	HstStmFbr   = "hstStmFbr"
	HstCndFbr   = "hstCndFbr"
	HstStgyFbr  = "hstStgyFbr"
	HstPortFbr  = "hstPortFbr"
	HstPrfmFbr  = "hstPrfmFbr"

	HstInstrWve = "hstInstrWve"
	HstInrvlWve = "hstInrvlWve"
	HstSideWve  = "hstSideWve"
	HstStmWve   = "hstStmWve"
	HstCndWve   = "hstCndWve"
	HstStgyWve  = "hstStgyWve"
	HstPortWve  = "hstPortWve"
	HstPrfmWve  = "hstPrfmWve"
)
const ( // trms
	Spce         = "spce"
	Cmnt         = "cmnt"
	Idn          = "idn"
	Asn          = "asn"
	Each         = "each"
	PllEach      = "pllEach"
	PllWait      = "pllWait"
	Then         = "then"
	Else         = "else"
	Lng          = "lng"
	Opt          = "opt"
	Cfg          = "cfg"
	Prfm         = "prfm"
	PrfmDlt      = "prfmDlt"
	StgyPrfm     = "stgyPrfm"
	StgyPrfmChan = "stgyPrfmChan"
	PortPrfm     = "portPrfm"
	Trd          = "trd"
	Trds         = "trds"
	TrdStmSeg    = "trdStmSeg"
	TrdsStm      = "trdsStm"
	OpnTrd       = "opnTrd"
	ClsTrd       = "clsTrd"
	MayTrd       = "mayTrd"
	CalcOpn      = "calcOpn"
	CalcCls      = "calcCls"
	CalcPrfm     = "calcPrfm"
	LoadHst      = "loadHst"
	AcntRefresh  = "acntRefresh"
)
const (
	Acs       = "acs"
	Lit       = "lit"
	Cnst      = "cnst"
	Ifc       = "ifc"
	Interface = "interface"
	Mu        = "mu"
)

const ( // instr names (used by Oanda)
	EurUsdName = "eur_usd"
	AudUsdName = "aud_usd"
	NzdUsdName = "nzd_usd"
	GbpUsdName = "gbp_usd"
)

const (
	Oan    = "oan" // prvs
	PrvOan = "prvOan"
	EurUsd = "eurUsd" // instrs (used by prolang)
	AudUsd = "audUsd"
	NzdUsd = "nzdUsd"
	GbpUsd = "gbpUsd"
	I      = "i"
	S      = "s"
	M      = "m"
	H      = "h"
	D      = "d"
	S1     = "s1" // inrvls (and tme cnsts)
	S5     = "s5"
	S10    = "s10"
	S15    = "s15"
	S20    = "s20"
	S30    = "s30"
	S40    = "s40"
	S50    = "s50"
	M1     = "m1"
	M5     = "m5"
	M10    = "m10"
	M15    = "m15"
	M20    = "m20"
	M30    = "m30"
	M40    = "m40"
	M50    = "m50"
	H1     = "h1"
	D1     = "d1"
	Bid    = "bid"
	Ask    = "ask"
	Bids   = "bids" // sides
	Asks   = "asks"
	Fst    = "fst" // agg (flts/sideInrvls)
	Lst    = "lst"
	Sum    = "sum"
	Prd    = "prd"
	// Min
	// Max
	Mid        = "mid"
	Avg        = "avg"
	AvgGeo     = "avgGeo"
	Mdn        = "mdn"
	Sma        = "sma"
	Gma        = "gma"
	Wma        = "wma"
	SmaMas     = "smaMas"
	GmaMas     = "gmaMas"
	WmaMas     = "wmaMas"
	SubSumPos  = "subSumPos"
	SubSumNeg  = "subSumNeg"
	Rsi        = "rsi"
	Wrsi       = "wrsi"
	Vrnc       = "vrnc"
	Std        = "std"
	Zscr       = "zscr"
	RngFul     = "rngFul"
	RngLst     = "rngLst"
	Pro        = "pro"
	ProLst     = "proLst"
	ProSma     = "proSma"
	ProAlma    = "proAlma"
	Sar        = "sar"
	Alma       = "alma"
	Ema        = "ema"
	Pos        = "pos" // una
	Neg        = "neg"
	Inv        = "inv"
	Sgn        = "sgn"
	Pct        = "pct"
	Sqr        = "sqr"
	Sqrt       = "sqrt"
	Add        = "add" // flt
	Sub        = "sub"
	Mul        = "mul"
	Div        = "div"
	Rem        = "rem"
	Pow        = "pow"
	Slp        = "slp"
	SlpScl     = "slpScl"
	PipScl     = "pipScl"
	Pipette    = "pipette"
	PipetteScl = "pipetteScl"
	MktWeeks   = "mktWeeks"
	MktDays    = "mktDays"
	MktHrs     = "mktHrs"
	// Min
	// Max
	CrsInvSum = "crsInvSum"
	CntrDist  = "cntrDist"
	Segs      = "segs"
	SegActs   = "segActs"
	At        = "at"
	Ats       = "ats"
	IsValid   = "isValid"
	Mrg       = "mrg"
	Union     = "union"
	Ensure    = "ensure"
	Len       = "len"
	PrfLos    = "prfLos"
	Vals      = "vals"
	Rev       = "rev"
)

const ( // stm op categories
	Rte        = "rte" // root
	Rtes       = "rtes"
	Una        = "una" // unary
	Unas       = "unas"
	Scl        = "scl" // scalar
	Scls       = "scls"
	Sel        = "sel"
	Sels       = "sels"
	Agg        = "agg" // aggregate
	Aggs       = "aggs"
	Inr        = "inr" // inner
	Inrs       = "inrs"
	Otr        = "otr" // outer
	Otrs       = "otrs"
	Rel        = "rel" // relational
	Rels       = "rels"
	Ort        = "ort"
	Orts       = "orts"
	StmBse     = "stmBse"
	StmRte     = "stmRte"
	StmRteSar  = "stmRteSar"
	StmRteAlma = "stmRteAlma"
	StmRtes    = "stmRtes"
	StmUna     = "stmUna"
	StmUnas    = "stmUnas"
	StmScl     = "stmScl"
	StmScls    = "stmScls"
	StmSel     = "stmSel"
	StmSels    = "stmSels"
	StmAgg     = "stmAgg"
	StmAggAlma = "stmAggAlma"
	StmAggs    = "stmAggs"
	StmInr     = "stmInr"
	StmInrSlp  = "stmInrSlp"
	StmInrs    = "stmInrs"
	StmOtr     = "stmOtr"
	StmOtrs    = "stmOtrs"
	CndBse     = "cndBse"
	CndScl     = "cndScl"
	CndInr     = "cndInr"
	CndInrs    = "cndInrs"
	CndOtr     = "cndOtr"
	CndOtrs    = "cndOtrs"
	CndAnd     = "cndAnd"
	CndSeq     = "cndSeq"
	StgyBse    = "stgyBse"
	StgyLong   = "stgyLong"
	StgyShrt   = "stgyShrt"
	Port       = "port"
)

const ( // methods
	Prnt       = "prnt"
	Set        = "set"
	Get        = "get"
	Slf        = "slf"
	Cnt        = "cnt"
	Calc       = "calc"
	Trnc       = "trnc"
	Rand       = "rand"
	Ifo        = "ifo"
	New        = "new"
	Now        = "now" // tme fn
	Not        = "not"
	Eql        = "eql" // idn
	Neq        = "neq"
	Lss        = "lss" // rel
	Gtr        = "gtr"
	Leq        = "leq"
	Geq        = "geq"
	Asc        = "asc"
	Dsc        = "dsc"
	And        = "and"
	Seq        = "seq"
	Ml         = "ml"
	Long       = "long"
	LongPll    = "longPll"
	LongMl     = "longMl"
	LongSrch   = "longSrch"
	LongRlng   = "longRlng"
	Shrt       = "shrt"
	ShrtPll    = "shrtPll"
	ShrtSrch   = "shrtSrch"
	ShrtRlng   = "shrtRlng"
	Unsub      = "unsub"
	Opn        = "opn"
	Cls        = "cls"
	TryOpn     = "tryOpn"
	TryCls     = "tryCls"
	PrfLosDur  = "prfLosTme"
	Cpy        = "cpy"
	Cmpl       = "cmpl"
	Cmp        = "cmp"
	SrtAsc     = "srtAsc"
	SrtAscEql  = "srtAscEql"
	SrtDsc     = "srtDsc"
	SrtDscEql  = "srtDscEql"
	SrtQuick   = "srtQuick"
	SrtIns     = "srtIns"
	SrtMdnOf3  = "srtMdnOf3"
	Swp        = "swp"
	Eval       = "eval"
	DstToInstr = "dstToInstr"
	String     = "string"
	Bytes      = "bytes"
	BytWrt     = "bytWrt"
	BytRed     = "bytRed"
	KeyWrt     = "keyWrt"
	StrWrt     = "strWrt"
	StrRed     = "strRed"
	PthWrt     = "pthWrt"
	Pth        = "pth"
	PrmWrt     = "prmWrt"
	Prm        = "prm"
	Ttl        = "ttl"
	Is         = "is"
	IsFlg      = "isFlg"
	Key        = "key"
	Pll        = "pll"
	Fn         = "fn"
	FnAct      = "fnAct"
	Has        = "has"
	HstStmBse  = "hstStmBse"
	RltStmBse  = "rltStmBse"
	HstCndBse  = "hstCndBse"
	RltCndBse  = "rltCndBse"
	HstStgyBse = "hstStgyBse"
	RltStgyBse = "rltStgyBse"
	CldSav     = "cldSav"
	CldQry     = "cldQry"
	CldSchema  = "cldSchema"
	SrchLong   = "srchLong"
	Rndr       = "rndr"
	Sho        = "sho"
	Siz        = "siz"
	Loop       = "loop"
	Pip        = "pip"
	PipPerDay  = "pipPerDay"
	SrchIdx    = "srchIdx"
	Rx         = "rx"
	RxA        = "rxA"
	Tx         = "tx"
	Align      = "align"

	Measure    = "measure"
	Draw       = "draw"
	Sampl      = "sampl"
	Rgba       = "rgba"
	Rgb        = "rgb"
	Hex        = "hex"
	Opa        = "opa"
	StmBnd     = "stmBnd"
	HrzLn      = "hrzLn"
	HrzBnd     = "hrzBnd"
	HrzBndAt   = "hrzBndAt"
	HrzScl     = "hrzScl"
	HrzSclVal  = "hrzSclVal"
	HrzRng     = "HrzRng"
	VrtLn      = "vrtLn"
	VrtBnd     = "vrtBnd"
	VrtBndAt   = "vrtBndAt"
	VrtScl     = "vrtScl"
	VrtSclVal  = "vrtSclVal"
	VrtRng     = "vrtRng"
	ToPlts     = "toPlts"
	Run        = "run"
	Resolution = "resolution"
)

const (
	Vis  = "vis"
	Fnt  = "fnt"
	Clr  = "clr"
	Pen  = "pen"
	Plt  = "plt"
	Hrz  = "hrz"
	Vrt  = "vrt"
	Dpth = "dpth"
	Log  = "log"
)

const ( // print decorations
	Ty = " (-)" // star wars Ty fighter
)

const ( // material design color names
	Red        = "red"
	Pink       = "pink"
	Purple     = "purple"
	DeepPurple = "deepPurple"
	Indigo     = "indigo"
	Blue       = "blue"
	LightBlue  = "lightBlue"
	Cyan       = "cyan"
	Teal       = "teal"
	Green      = "green"
	LightGreen = "lightGreen"
	Lime       = "lime"
	Yellow     = "yellow"
	Amber      = "amber"
	Orange     = "orange"
	DeepOrange = "deepOrange"
	Brown      = "brown"
	BlueGrey   = "blueGrey"
	Grey       = "grey"
)

const (
	Clr50  = "50"
	Clr100 = "100"
	Clr200 = "200"
	Clr300 = "300"
	Clr400 = "400"
	Clr500 = "500"
	Clr600 = "600"
	Clr700 = "700"
	Clr800 = "800"
	Clr900 = "900"
)
