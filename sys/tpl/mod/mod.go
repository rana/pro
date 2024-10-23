package mod

type (
	Mod uint32
)

const (
	Ptr Mod = 1 << iota
	Slice
	Variadic
	Lowercase
	CurlyBracket
	None Mod = 0
)

var (
	modStrs = map[Mod]string{
		Ptr:          "Ptr",
		Slice:        "Slice",
		Variadic:     "Variadic",
		Lowercase:    "Lowercase",
		CurlyBracket: "CurlyBracket",
		None:         "None",
	}
)

func (x Mod) IsPtr() bool          { return x&Ptr == Ptr }
func (x Mod) IsSlice() bool        { return x&Slice == Slice }
func (x Mod) IsVariadic() bool     { return x&Variadic == Variadic }
func (x Mod) IsLowercase() bool    { return x&Lowercase == Lowercase }
func (x Mod) IsCurlyBracket() bool { return x&CurlyBracket == CurlyBracket }
func (x Mod) String() string       { return modStrs[x] }
