package pen

import (
	"sys/ana/vis/clr"
	"sys/bsc/flt"
	"sys/bsc/str"
	"sys/bsc/unt"
)

var (
	Black          = Pen{Clr: clr.Black, Wid: 1}
	White          = Pen{Clr: clr.White, Wid: 1}
	Red50          = Pen{Clr: clr.Red50, Wid: 1}
	Red100         = Pen{Clr: clr.Red100, Wid: 1}
	Red200         = Pen{Clr: clr.Red200, Wid: 1}
	Red300         = Pen{Clr: clr.Red300, Wid: 1}
	Red400         = Pen{Clr: clr.Red400, Wid: 1}
	Red500         = Pen{Clr: clr.Red500, Wid: 1}
	Red600         = Pen{Clr: clr.Red600, Wid: 1}
	Red700         = Pen{Clr: clr.Red700, Wid: 1}
	Red800         = Pen{Clr: clr.Red800, Wid: 1}
	Red900         = Pen{Clr: clr.Red900, Wid: 1}
	RedA100        = Pen{Clr: clr.RedA100, Wid: 1}
	RedA200        = Pen{Clr: clr.RedA200, Wid: 1}
	RedA400        = Pen{Clr: clr.RedA400, Wid: 1}
	RedA700        = Pen{Clr: clr.RedA700, Wid: 1}
	Pink50         = Pen{Clr: clr.Pink50, Wid: 1}
	Pink100        = Pen{Clr: clr.Pink100, Wid: 1}
	Pink200        = Pen{Clr: clr.Pink200, Wid: 1}
	Pink300        = Pen{Clr: clr.Pink300, Wid: 1}
	Pink400        = Pen{Clr: clr.Pink400, Wid: 1}
	Pink500        = Pen{Clr: clr.Pink500, Wid: 1}
	Pink600        = Pen{Clr: clr.Pink600, Wid: 1}
	Pink700        = Pen{Clr: clr.Pink700, Wid: 1}
	Pink800        = Pen{Clr: clr.Pink800, Wid: 1}
	Pink900        = Pen{Clr: clr.Pink900, Wid: 1}
	PinkA100       = Pen{Clr: clr.PinkA100, Wid: 1}
	PinkA200       = Pen{Clr: clr.PinkA200, Wid: 1}
	PinkA400       = Pen{Clr: clr.PinkA400, Wid: 1}
	PinkA700       = Pen{Clr: clr.PinkA700, Wid: 1}
	Purple50       = Pen{Clr: clr.Purple50, Wid: 1}
	Purple100      = Pen{Clr: clr.Purple100, Wid: 1}
	Purple200      = Pen{Clr: clr.Purple200, Wid: 1}
	Purple300      = Pen{Clr: clr.Purple300, Wid: 1}
	Purple400      = Pen{Clr: clr.Purple400, Wid: 1}
	Purple500      = Pen{Clr: clr.Purple500, Wid: 1}
	Purple600      = Pen{Clr: clr.Purple600, Wid: 1}
	Purple700      = Pen{Clr: clr.Purple700, Wid: 1}
	Purple800      = Pen{Clr: clr.Purple800, Wid: 1}
	Purple900      = Pen{Clr: clr.Purple900, Wid: 1}
	PurpleA100     = Pen{Clr: clr.PurpleA100, Wid: 1}
	PurpleA200     = Pen{Clr: clr.PurpleA200, Wid: 1}
	PurpleA400     = Pen{Clr: clr.PurpleA400, Wid: 1}
	PurpleA700     = Pen{Clr: clr.PurpleA700, Wid: 1}
	DeepPurple50   = Pen{Clr: clr.DeepPurple50, Wid: 1}
	DeepPurple100  = Pen{Clr: clr.DeepPurple100, Wid: 1}
	DeepPurple200  = Pen{Clr: clr.DeepPurple200, Wid: 1}
	DeepPurple300  = Pen{Clr: clr.DeepPurple300, Wid: 1}
	DeepPurple400  = Pen{Clr: clr.DeepPurple400, Wid: 1}
	DeepPurple500  = Pen{Clr: clr.DeepPurple500, Wid: 1}
	DeepPurple600  = Pen{Clr: clr.DeepPurple600, Wid: 1}
	DeepPurple700  = Pen{Clr: clr.DeepPurple700, Wid: 1}
	DeepPurple800  = Pen{Clr: clr.DeepPurple800, Wid: 1}
	DeepPurple900  = Pen{Clr: clr.DeepPurple900, Wid: 1}
	DeepPurpleA100 = Pen{Clr: clr.DeepPurpleA100, Wid: 1}
	DeepPurpleA200 = Pen{Clr: clr.DeepPurpleA200, Wid: 1}
	DeepPurpleA400 = Pen{Clr: clr.DeepPurpleA400, Wid: 1}
	DeepPurpleA700 = Pen{Clr: clr.DeepPurpleA700, Wid: 1}
	Indigo50       = Pen{Clr: clr.Indigo50, Wid: 1}
	Indigo100      = Pen{Clr: clr.Indigo100, Wid: 1}
	Indigo200      = Pen{Clr: clr.Indigo200, Wid: 1}
	Indigo300      = Pen{Clr: clr.Indigo300, Wid: 1}
	Indigo400      = Pen{Clr: clr.Indigo400, Wid: 1}
	Indigo500      = Pen{Clr: clr.Indigo500, Wid: 1}
	Indigo600      = Pen{Clr: clr.Indigo600, Wid: 1}
	Indigo700      = Pen{Clr: clr.Indigo700, Wid: 1}
	Indigo800      = Pen{Clr: clr.Indigo800, Wid: 1}
	Indigo900      = Pen{Clr: clr.Indigo900, Wid: 1}
	IndigoA100     = Pen{Clr: clr.IndigoA100, Wid: 1}
	IndigoA200     = Pen{Clr: clr.IndigoA200, Wid: 1}
	IndigoA400     = Pen{Clr: clr.IndigoA400, Wid: 1}
	IndigoA700     = Pen{Clr: clr.IndigoA700, Wid: 1}
	Blue50         = Pen{Clr: clr.Blue50, Wid: 1}
	Blue100        = Pen{Clr: clr.Blue100, Wid: 1}
	Blue200        = Pen{Clr: clr.Blue200, Wid: 1}
	Blue300        = Pen{Clr: clr.Blue300, Wid: 1}
	Blue400        = Pen{Clr: clr.Blue400, Wid: 1}
	Blue500        = Pen{Clr: clr.Blue500, Wid: 1}
	Blue600        = Pen{Clr: clr.Blue600, Wid: 1}
	Blue700        = Pen{Clr: clr.Blue700, Wid: 1}
	Blue800        = Pen{Clr: clr.Blue800, Wid: 1}
	Blue900        = Pen{Clr: clr.Blue900, Wid: 1}
	BlueA100       = Pen{Clr: clr.BlueA100, Wid: 1}
	BlueA200       = Pen{Clr: clr.BlueA200, Wid: 1}
	BlueA400       = Pen{Clr: clr.BlueA400, Wid: 1}
	BlueA700       = Pen{Clr: clr.BlueA700, Wid: 1}
	LightBlue50    = Pen{Clr: clr.LightBlue50, Wid: 1}
	LightBlue100   = Pen{Clr: clr.LightBlue100, Wid: 1}
	LightBlue200   = Pen{Clr: clr.LightBlue200, Wid: 1}
	LightBlue300   = Pen{Clr: clr.LightBlue300, Wid: 1}
	LightBlue400   = Pen{Clr: clr.LightBlue400, Wid: 1}
	LightBlue500   = Pen{Clr: clr.LightBlue500, Wid: 1}
	LightBlue600   = Pen{Clr: clr.LightBlue600, Wid: 1}
	LightBlue700   = Pen{Clr: clr.LightBlue700, Wid: 1}
	LightBlue800   = Pen{Clr: clr.LightBlue800, Wid: 1}
	LightBlue900   = Pen{Clr: clr.LightBlue900, Wid: 1}
	LightBlueA100  = Pen{Clr: clr.LightBlueA100, Wid: 1}
	LightBlueA200  = Pen{Clr: clr.LightBlueA200, Wid: 1}
	LightBlueA400  = Pen{Clr: clr.LightBlueA400, Wid: 1}
	LightBlueA700  = Pen{Clr: clr.LightBlueA700, Wid: 1}
	Cyan50         = Pen{Clr: clr.Cyan50, Wid: 1}
	Cyan100        = Pen{Clr: clr.Cyan100, Wid: 1}
	Cyan200        = Pen{Clr: clr.Cyan200, Wid: 1}
	Cyan300        = Pen{Clr: clr.Cyan300, Wid: 1}
	Cyan400        = Pen{Clr: clr.Cyan400, Wid: 1}
	Cyan500        = Pen{Clr: clr.Cyan500, Wid: 1}
	Cyan600        = Pen{Clr: clr.Cyan600, Wid: 1}
	Cyan700        = Pen{Clr: clr.Cyan700, Wid: 1}
	Cyan800        = Pen{Clr: clr.Cyan800, Wid: 1}
	Cyan900        = Pen{Clr: clr.Cyan900, Wid: 1}
	CyanA100       = Pen{Clr: clr.CyanA100, Wid: 1}
	CyanA200       = Pen{Clr: clr.CyanA200, Wid: 1}
	CyanA400       = Pen{Clr: clr.CyanA400, Wid: 1}
	CyanA700       = Pen{Clr: clr.CyanA700, Wid: 1}
	Teal50         = Pen{Clr: clr.Teal50, Wid: 1}
	Teal100        = Pen{Clr: clr.Teal100, Wid: 1}
	Teal200        = Pen{Clr: clr.Teal200, Wid: 1}
	Teal300        = Pen{Clr: clr.Teal300, Wid: 1}
	Teal400        = Pen{Clr: clr.Teal400, Wid: 1}
	Teal500        = Pen{Clr: clr.Teal500, Wid: 1}
	Teal600        = Pen{Clr: clr.Teal600, Wid: 1}
	Teal700        = Pen{Clr: clr.Teal700, Wid: 1}
	Teal800        = Pen{Clr: clr.Teal800, Wid: 1}
	Teal900        = Pen{Clr: clr.Teal900, Wid: 1}
	TealA100       = Pen{Clr: clr.TealA100, Wid: 1}
	TealA200       = Pen{Clr: clr.TealA200, Wid: 1}
	TealA400       = Pen{Clr: clr.TealA400, Wid: 1}
	TealA700       = Pen{Clr: clr.TealA700, Wid: 1}
	Green50        = Pen{Clr: clr.Green50, Wid: 1}
	Green100       = Pen{Clr: clr.Green100, Wid: 1}
	Green200       = Pen{Clr: clr.Green200, Wid: 1}
	Green300       = Pen{Clr: clr.Green300, Wid: 1}
	Green400       = Pen{Clr: clr.Green400, Wid: 1}
	Green500       = Pen{Clr: clr.Green500, Wid: 1}
	Green600       = Pen{Clr: clr.Green600, Wid: 1}
	Green700       = Pen{Clr: clr.Green700, Wid: 1}
	Green800       = Pen{Clr: clr.Green800, Wid: 1}
	Green900       = Pen{Clr: clr.Green900, Wid: 1}
	GreenA100      = Pen{Clr: clr.GreenA100, Wid: 1}
	GreenA200      = Pen{Clr: clr.GreenA200, Wid: 1}
	GreenA400      = Pen{Clr: clr.GreenA400, Wid: 1}
	GreenA700      = Pen{Clr: clr.GreenA700, Wid: 1}
	LightGreen50   = Pen{Clr: clr.LightGreen50, Wid: 1}
	LightGreen100  = Pen{Clr: clr.LightGreen100, Wid: 1}
	LightGreen200  = Pen{Clr: clr.LightGreen200, Wid: 1}
	LightGreen300  = Pen{Clr: clr.LightGreen300, Wid: 1}
	LightGreen400  = Pen{Clr: clr.LightGreen400, Wid: 1}
	LightGreen500  = Pen{Clr: clr.LightGreen500, Wid: 1}
	LightGreen600  = Pen{Clr: clr.LightGreen600, Wid: 1}
	LightGreen700  = Pen{Clr: clr.LightGreen700, Wid: 1}
	LightGreen800  = Pen{Clr: clr.LightGreen800, Wid: 1}
	LightGreen900  = Pen{Clr: clr.LightGreen900, Wid: 1}
	LightGreenA100 = Pen{Clr: clr.LightGreenA100, Wid: 1}
	LightGreenA200 = Pen{Clr: clr.LightGreenA200, Wid: 1}
	LightGreenA400 = Pen{Clr: clr.LightGreenA400, Wid: 1}
	LightGreenA700 = Pen{Clr: clr.LightGreenA700, Wid: 1}
	Lime50         = Pen{Clr: clr.Lime50, Wid: 1}
	Lime100        = Pen{Clr: clr.Lime100, Wid: 1}
	Lime200        = Pen{Clr: clr.Lime200, Wid: 1}
	Lime300        = Pen{Clr: clr.Lime300, Wid: 1}
	Lime400        = Pen{Clr: clr.Lime400, Wid: 1}
	Lime500        = Pen{Clr: clr.Lime500, Wid: 1}
	Lime600        = Pen{Clr: clr.Lime600, Wid: 1}
	Lime700        = Pen{Clr: clr.Lime700, Wid: 1}
	Lime800        = Pen{Clr: clr.Lime800, Wid: 1}
	Lime900        = Pen{Clr: clr.Lime900, Wid: 1}
	LimeA100       = Pen{Clr: clr.LimeA100, Wid: 1}
	LimeA200       = Pen{Clr: clr.LimeA200, Wid: 1}
	LimeA400       = Pen{Clr: clr.LimeA400, Wid: 1}
	LimeA700       = Pen{Clr: clr.LimeA700, Wid: 1}
	Yellow50       = Pen{Clr: clr.Yellow50, Wid: 1}
	Yellow100      = Pen{Clr: clr.Yellow100, Wid: 1}
	Yellow200      = Pen{Clr: clr.Yellow200, Wid: 1}
	Yellow300      = Pen{Clr: clr.Yellow300, Wid: 1}
	Yellow400      = Pen{Clr: clr.Yellow400, Wid: 1}
	Yellow500      = Pen{Clr: clr.Yellow500, Wid: 1}
	Yellow600      = Pen{Clr: clr.Yellow600, Wid: 1}
	Yellow700      = Pen{Clr: clr.Yellow700, Wid: 1}
	Yellow800      = Pen{Clr: clr.Yellow800, Wid: 1}
	Yellow900      = Pen{Clr: clr.Yellow900, Wid: 1}
	YellowA100     = Pen{Clr: clr.YellowA100, Wid: 1}
	YellowA200     = Pen{Clr: clr.YellowA200, Wid: 1}
	YellowA400     = Pen{Clr: clr.YellowA400, Wid: 1}
	YellowA700     = Pen{Clr: clr.YellowA700, Wid: 1}
	Amber50        = Pen{Clr: clr.Amber50, Wid: 1}
	Amber100       = Pen{Clr: clr.Amber100, Wid: 1}
	Amber200       = Pen{Clr: clr.Amber200, Wid: 1}
	Amber300       = Pen{Clr: clr.Amber300, Wid: 1}
	Amber400       = Pen{Clr: clr.Amber400, Wid: 1}
	Amber500       = Pen{Clr: clr.Amber500, Wid: 1}
	Amber600       = Pen{Clr: clr.Amber600, Wid: 1}
	Amber700       = Pen{Clr: clr.Amber700, Wid: 1}
	Amber800       = Pen{Clr: clr.Amber800, Wid: 1}
	Amber900       = Pen{Clr: clr.Amber900, Wid: 1}
	AmberA100      = Pen{Clr: clr.AmberA100, Wid: 1}
	AmberA200      = Pen{Clr: clr.AmberA200, Wid: 1}
	AmberA400      = Pen{Clr: clr.AmberA400, Wid: 1}
	AmberA700      = Pen{Clr: clr.AmberA700, Wid: 1}
	Orange50       = Pen{Clr: clr.Orange50, Wid: 1}
	Orange100      = Pen{Clr: clr.Orange100, Wid: 1}
	Orange200      = Pen{Clr: clr.Orange200, Wid: 1}
	Orange300      = Pen{Clr: clr.Orange300, Wid: 1}
	Orange400      = Pen{Clr: clr.Orange400, Wid: 1}
	Orange500      = Pen{Clr: clr.Orange500, Wid: 1}
	Orange600      = Pen{Clr: clr.Orange600, Wid: 1}
	Orange700      = Pen{Clr: clr.Orange700, Wid: 1}
	Orange800      = Pen{Clr: clr.Orange800, Wid: 1}
	Orange900      = Pen{Clr: clr.Orange900, Wid: 1}
	OrangeA100     = Pen{Clr: clr.OrangeA100, Wid: 1}
	OrangeA200     = Pen{Clr: clr.OrangeA200, Wid: 1}
	OrangeA400     = Pen{Clr: clr.OrangeA400, Wid: 1}
	OrangeA700     = Pen{Clr: clr.OrangeA700, Wid: 1}
	DeepOrange50   = Pen{Clr: clr.DeepOrange50, Wid: 1}
	DeepOrange100  = Pen{Clr: clr.DeepOrange100, Wid: 1}
	DeepOrange200  = Pen{Clr: clr.DeepOrange200, Wid: 1}
	DeepOrange300  = Pen{Clr: clr.DeepOrange300, Wid: 1}
	DeepOrange400  = Pen{Clr: clr.DeepOrange400, Wid: 1}
	DeepOrange500  = Pen{Clr: clr.DeepOrange500, Wid: 1}
	DeepOrange600  = Pen{Clr: clr.DeepOrange600, Wid: 1}
	DeepOrange700  = Pen{Clr: clr.DeepOrange700, Wid: 1}
	DeepOrange800  = Pen{Clr: clr.DeepOrange800, Wid: 1}
	DeepOrange900  = Pen{Clr: clr.DeepOrange900, Wid: 1}
	DeepOrangeA100 = Pen{Clr: clr.DeepOrangeA100, Wid: 1}
	DeepOrangeA200 = Pen{Clr: clr.DeepOrangeA200, Wid: 1}
	DeepOrangeA400 = Pen{Clr: clr.DeepOrangeA400, Wid: 1}
	DeepOrangeA700 = Pen{Clr: clr.DeepOrangeA700, Wid: 1}
	Brown50        = Pen{Clr: clr.Brown50, Wid: 1}
	Brown100       = Pen{Clr: clr.Brown100, Wid: 1}
	Brown200       = Pen{Clr: clr.Brown200, Wid: 1}
	Brown300       = Pen{Clr: clr.Brown300, Wid: 1}
	Brown400       = Pen{Clr: clr.Brown400, Wid: 1}
	Brown500       = Pen{Clr: clr.Brown500, Wid: 1}
	Brown600       = Pen{Clr: clr.Brown600, Wid: 1}
	Brown700       = Pen{Clr: clr.Brown700, Wid: 1}
	Brown800       = Pen{Clr: clr.Brown800, Wid: 1}
	Brown900       = Pen{Clr: clr.Brown900, Wid: 1}
	BlueGrey50     = Pen{Clr: clr.BlueGrey50, Wid: 1}
	BlueGrey100    = Pen{Clr: clr.BlueGrey100, Wid: 1}
	BlueGrey200    = Pen{Clr: clr.BlueGrey200, Wid: 1}
	BlueGrey300    = Pen{Clr: clr.BlueGrey300, Wid: 1}
	BlueGrey400    = Pen{Clr: clr.BlueGrey400, Wid: 1}
	BlueGrey500    = Pen{Clr: clr.BlueGrey500, Wid: 1}
	BlueGrey600    = Pen{Clr: clr.BlueGrey600, Wid: 1}
	BlueGrey700    = Pen{Clr: clr.BlueGrey700, Wid: 1}
	BlueGrey800    = Pen{Clr: clr.BlueGrey800, Wid: 1}
	BlueGrey900    = Pen{Clr: clr.BlueGrey900, Wid: 1}
	Grey50         = Pen{Clr: clr.Grey50, Wid: 1}
	Grey100        = Pen{Clr: clr.Grey100, Wid: 1}
	Grey200        = Pen{Clr: clr.Grey200, Wid: 1}
	Grey300        = Pen{Clr: clr.Grey300, Wid: 1}
	Grey400        = Pen{Clr: clr.Grey400, Wid: 1}
	Grey500        = Pen{Clr: clr.Grey500, Wid: 1}
	Grey600        = Pen{Clr: clr.Grey600, Wid: 1}
	Grey700        = Pen{Clr: clr.Grey700, Wid: 1}
	Grey800        = Pen{Clr: clr.Grey800, Wid: 1}
	Grey900        = Pen{Clr: clr.Grey900, Wid: 1}
)

type (
	Pen struct {
		Clr clr.Clr
		Wid unt.Unt
	}
	PenScp struct {
		Idx uint32
		Arr []Pen
	}
)

func New(clr clr.Clr, wid ...unt.Unt) (p Pen) {
	p.Clr = clr
	if len(wid) != 0 {
		p.Wid = wid[0]
	} else {
		p.Wid = 1
	}
	return p
}
func Rgba(r, g, b, a flt.Flt, wid ...unt.Unt) (p Pen) {
	p.Clr = clr.Rgba(r, g, b, a)
	if len(wid) != 0 {
		p.Wid = wid[0]
	} else {
		p.Wid = 1
	}
	return p
}
func Rgb(r, g, b flt.Flt, wid ...unt.Unt) (p Pen) {
	p.Clr = clr.Rgb(r, g, b)
	if len(wid) != 0 {
		p.Wid = wid[0]
	} else {
		p.Wid = 1
	}
	return p
}
func Hex(txt str.Str, wid ...unt.Unt) (p Pen) {
	p.Clr = clr.Hex(txt)
	if len(wid) != 0 {
		p.Wid = wid[0]
	} else {
		p.Wid = 1
	}
	return p
}
func (x Pen) Opa(pct flt.Flt) Pen {
	x.Clr = x.Clr.Opa(pct)
	return x
}
func (x Pen) Inv() Pen {
	x.Clr = x.Clr.Inv()
	return x
}
