package ana

var (
	trdRsnOpnNames = map[TrdRsnOpn]string{
		NoTrdRsnOpn: "noTrdRsnOpn",
		ErrMktWeek:  "errMktWeek",
		ErrOpnCnd:   "errOpnCnd",
		NearMktOpn:  "nearMktOpn",
		NearMktCls:  "nearMktCls",
		SpdLrg:      "spdLrg",
		NoCls:       "noCls",
		PrvReject:   "prvReject",
		InTrd:       "inTrd",
		FilHstGap:   "filHstGap",
		NoCapital:   "noCapital",
		PrvErr:      "prvErr",
	}
)

const (
	NoTrdRsnOpn TrdRsnOpn = iota
	ErrMktWeek
	ErrOpnCnd
	NearMktOpn
	NearMktCls
	SpdLrg
	NoCls
	PrvReject
	InTrd
	FilHstGap
	NoCapital
	PrvErr
)

type (
	TrdRsnOpn byte
)

func (x TrdRsnOpn) String() string { return trdRsnOpnNames[x] }
