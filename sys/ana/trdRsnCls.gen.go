package ana

var (
	trdRsnClsNames = map[TrdRsnCls]string{
		NoTrdRsnCls: "noTrdRsnCls",
		Prf:         "prf",
		Los:         "los",
		Dur:         "dur",
		Cnd:         "cnd",
	}
)

const (
	NoTrdRsnCls TrdRsnCls = iota
	Prf
	Los
	Dur
	Cnd
)

type (
	TrdRsnCls byte
)

func (x TrdRsnCls) String() string { return trdRsnClsNames[x] }
