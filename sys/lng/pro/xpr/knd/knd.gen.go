package knd

var (
	kndNames = map[Knd]string{
		NoKnd:           "noKnd",
		SysInterface:    "sysInterface",
		LogLogr:         "logLogr",
		StrStr:          "strStr",
		BolBol:          "bolBol",
		FltFlt:          "fltFlt",
		UntUnt:          "untUnt",
		IntInt:          "intInt",
		TmeTme:          "tmeTme",
		BndBnd:          "bndBnd",
		FltRng:          "fltRng",
		TmeRng:          "tmeRng",
		StrsStrs:        "strsStrs",
		BolsBols:        "bolsBols",
		FltsFlts:        "fltsFlts",
		UntsUnts:        "untsUnts",
		IntsInts:        "intsInts",
		TmesTmes:        "tmesTmes",
		BndsBnds:        "bndsBnds",
		TmeRngs:         "tmeRngs",
		AnaTrd:          "anaTrd",
		AnaTrds:         "anaTrds",
		AnaPrfm:         "anaPrfm",
		AnaPrfms:        "anaPrfms",
		AnaPrfmDlt:      "anaPrfmDlt",
		AnaPort:         "anaPort",
		HstPrv:          "hstPrv",
		HstInstr:        "hstInstr",
		HstInrvl:        "hstInrvl",
		HstSide:         "hstSide",
		HstStm:          "hstStm",
		HstCnd:          "hstCnd",
		HstStgy:         "hstStgy",
		HstPrvs:         "hstPrvs",
		HstInstrs:       "hstInstrs",
		HstInrvls:       "hstInrvls",
		HstSides:        "hstSides",
		HstStms:         "hstStms",
		HstCnds:         "hstCnds",
		HstStgys:        "hstStgys",
		RltPrv:          "rltPrv",
		RltInstr:        "rltInstr",
		RltInrvl:        "rltInrvl",
		RltSide:         "rltSide",
		RltStm:          "rltStm",
		RltCnd:          "rltCnd",
		RltStgy:         "rltStgy",
		RltPrvs:         "rltPrvs",
		RltInstrs:       "rltInstrs",
		RltInrvls:       "rltInrvls",
		RltSides:        "rltSides",
		RltStms:         "rltStms",
		RltCnds:         "rltCnds",
		RltStgys:        "rltStgys",
		FntFnt:          "fntFnt",
		ClrClr:          "clrClr",
		PenPen:          "penPen",
		PenPens:         "penPens",
		PltPlt:          "pltPlt",
		PltPlts:         "pltPlts",
		PltTmeAxisX:     "pltTmeAxisX",
		PltFltAxisY:     "pltFltAxisY",
		PltStm:          "pltStm",
		PltFltsSctr:     "pltFltsSctr",
		PltFltsSctrDist: "pltFltsSctrDist",
		PltHrz:          "pltHrz",
		PltVrt:          "pltVrt",
		PltDpth:         "pltDpth",
		SysIdn:          "sysIdn",
		SysMu:           "sysMu",
	}
)

const (
	NoKnd Knd = iota
	SysInterface
	LogLogr
	StrStr
	BolBol
	FltFlt
	UntUnt
	IntInt
	TmeTme
	BndBnd
	FltRng
	TmeRng
	StrsStrs
	BolsBols
	FltsFlts
	UntsUnts
	IntsInts
	TmesTmes
	BndsBnds
	TmeRngs
	AnaTrd
	AnaTrds
	AnaPrfm
	AnaPrfms
	AnaPrfmDlt
	AnaPort
	HstPrv
	HstInstr
	HstInrvl
	HstSide
	HstStm
	HstCnd
	HstStgy
	HstPrvs
	HstInstrs
	HstInrvls
	HstSides
	HstStms
	HstCnds
	HstStgys
	RltPrv
	RltInstr
	RltInrvl
	RltSide
	RltStm
	RltCnd
	RltStgy
	RltPrvs
	RltInstrs
	RltInrvls
	RltSides
	RltStms
	RltCnds
	RltStgys
	FntFnt
	ClrClr
	PenPen
	PenPens
	PltPlt
	PltPlts
	PltTmeAxisX
	PltFltAxisY
	PltStm
	PltFltsSctr
	PltFltsSctrDist
	PltHrz
	PltVrt
	PltDpth
	SysIdn
	SysMu
)

type (
	Knd byte
)

func (x Knd) String() string { return kndNames[x] }
