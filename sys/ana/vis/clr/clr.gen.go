package clr

import (
	"fmt"
	"image/color"
	"strings"
	"sys/bsc/flt"
	"sys/bsc/str"
)

var (
	Black          = Clr{0x00, 0x00, 0x00, 0xff}
	White          = Clr{0xff, 0xff, 0xff, 0xff}
	Red50          = Clr{0xff, 0xeb, 0xee, 0xff}
	Red100         = Clr{0xff, 0xcd, 0xd2, 0xff}
	Red200         = Clr{0xef, 0x9a, 0x9a, 0xff}
	Red300         = Clr{0xe5, 0x73, 0x73, 0xff}
	Red400         = Clr{0xef, 0x53, 0x50, 0xff}
	Red500         = Clr{0xf4, 0x43, 0x36, 0xff}
	Red600         = Clr{0xe5, 0x39, 0x35, 0xff}
	Red700         = Clr{0xd3, 0x2f, 0x2f, 0xff}
	Red800         = Clr{0xc6, 0x28, 0x28, 0xff}
	Red900         = Clr{0xb7, 0x1c, 0x1c, 0xff}
	RedA100        = Clr{0xff, 0x8a, 0x80, 0xff}
	RedA200        = Clr{0xff, 0x52, 0x52, 0xff}
	RedA400        = Clr{0xff, 0x17, 0x44, 0xff}
	RedA700        = Clr{0xd5, 0x00, 0x00, 0xff}
	Pink50         = Clr{0xfc, 0xe4, 0xec, 0xff}
	Pink100        = Clr{0xf8, 0xbb, 0xd0, 0xff}
	Pink200        = Clr{0xf4, 0x8f, 0xb1, 0xff}
	Pink300        = Clr{0xf0, 0x62, 0x92, 0xff}
	Pink400        = Clr{0xec, 0x40, 0x7a, 0xff}
	Pink500        = Clr{0xe9, 0x1e, 0x63, 0xff}
	Pink600        = Clr{0xd8, 0x1b, 0x60, 0xff}
	Pink700        = Clr{0xc2, 0x18, 0x5b, 0xff}
	Pink800        = Clr{0xad, 0x14, 0x57, 0xff}
	Pink900        = Clr{0x88, 0x0e, 0x4f, 0xff}
	PinkA100       = Clr{0xff, 0x80, 0xab, 0xff}
	PinkA200       = Clr{0xff, 0x40, 0x81, 0xff}
	PinkA400       = Clr{0xf5, 0x00, 0x57, 0xff}
	PinkA700       = Clr{0xc5, 0x11, 0x62, 0xff}
	Purple50       = Clr{0xf3, 0xe5, 0xf5, 0xff}
	Purple100      = Clr{0xe1, 0xbe, 0xe7, 0xff}
	Purple200      = Clr{0xce, 0x93, 0xd8, 0xff}
	Purple300      = Clr{0xba, 0x68, 0xc8, 0xff}
	Purple400      = Clr{0xab, 0x47, 0xbc, 0xff}
	Purple500      = Clr{0x9c, 0x27, 0xb0, 0xff}
	Purple600      = Clr{0x8e, 0x24, 0xaa, 0xff}
	Purple700      = Clr{0x7b, 0x1f, 0xa2, 0xff}
	Purple800      = Clr{0x6a, 0x1b, 0x9a, 0xff}
	Purple900      = Clr{0x4a, 0x14, 0x8c, 0xff}
	PurpleA100     = Clr{0xea, 0x80, 0xfc, 0xff}
	PurpleA200     = Clr{0xe0, 0x40, 0xfb, 0xff}
	PurpleA400     = Clr{0xd5, 0x00, 0xf9, 0xff}
	PurpleA700     = Clr{0xaa, 0x00, 0xff, 0xff}
	DeepPurple50   = Clr{0xed, 0xe7, 0xf6, 0xff}
	DeepPurple100  = Clr{0xd1, 0xc4, 0xe9, 0xff}
	DeepPurple200  = Clr{0xb3, 0x9d, 0xdb, 0xff}
	DeepPurple300  = Clr{0x95, 0x75, 0xcd, 0xff}
	DeepPurple400  = Clr{0x7e, 0x57, 0xc2, 0xff}
	DeepPurple500  = Clr{0x67, 0x3a, 0xb7, 0xff}
	DeepPurple600  = Clr{0x5e, 0x35, 0xb1, 0xff}
	DeepPurple700  = Clr{0x51, 0x2d, 0xa8, 0xff}
	DeepPurple800  = Clr{0x45, 0x27, 0xa0, 0xff}
	DeepPurple900  = Clr{0x31, 0x1b, 0x92, 0xff}
	DeepPurpleA100 = Clr{0xb3, 0x88, 0xff, 0xff}
	DeepPurpleA200 = Clr{0x7c, 0x4d, 0xff, 0xff}
	DeepPurpleA400 = Clr{0x65, 0x1f, 0xff, 0xff}
	DeepPurpleA700 = Clr{0x62, 0x00, 0xea, 0xff}
	Indigo50       = Clr{0xe8, 0xea, 0xf6, 0xff}
	Indigo100      = Clr{0xc5, 0xca, 0xe9, 0xff}
	Indigo200      = Clr{0x9f, 0xa8, 0xda, 0xff}
	Indigo300      = Clr{0x79, 0x86, 0xcb, 0xff}
	Indigo400      = Clr{0x5c, 0x6b, 0xc0, 0xff}
	Indigo500      = Clr{0x3f, 0x51, 0xb5, 0xff}
	Indigo600      = Clr{0x39, 0x49, 0xab, 0xff}
	Indigo700      = Clr{0x30, 0x3f, 0x9f, 0xff}
	Indigo800      = Clr{0x28, 0x35, 0x93, 0xff}
	Indigo900      = Clr{0x1a, 0x23, 0x7e, 0xff}
	IndigoA100     = Clr{0x8c, 0x9e, 0xff, 0xff}
	IndigoA200     = Clr{0x53, 0x6d, 0xfe, 0xff}
	IndigoA400     = Clr{0x3d, 0x5a, 0xfe, 0xff}
	IndigoA700     = Clr{0x30, 0x4f, 0xfe, 0xff}
	Blue50         = Clr{0xe3, 0xf2, 0xfd, 0xff}
	Blue100        = Clr{0xbb, 0xde, 0xfb, 0xff}
	Blue200        = Clr{0x90, 0xca, 0xf9, 0xff}
	Blue300        = Clr{0x64, 0xb5, 0xf6, 0xff}
	Blue400        = Clr{0x42, 0xa5, 0xf5, 0xff}
	Blue500        = Clr{0x21, 0x96, 0xf3, 0xff}
	Blue600        = Clr{0x1e, 0x88, 0xe5, 0xff}
	Blue700        = Clr{0x19, 0x76, 0xd2, 0xff}
	Blue800        = Clr{0x15, 0x65, 0xc0, 0xff}
	Blue900        = Clr{0x0d, 0x47, 0xa1, 0xff}
	BlueA100       = Clr{0x82, 0xb1, 0xff, 0xff}
	BlueA200       = Clr{0x44, 0x8a, 0xff, 0xff}
	BlueA400       = Clr{0x29, 0x79, 0xff, 0xff}
	BlueA700       = Clr{0x29, 0x62, 0xff, 0xff}
	LightBlue50    = Clr{0xe1, 0xf5, 0xfe, 0xff}
	LightBlue100   = Clr{0xb3, 0xe5, 0xfc, 0xff}
	LightBlue200   = Clr{0x81, 0xd4, 0xfa, 0xff}
	LightBlue300   = Clr{0x4f, 0xc3, 0xf7, 0xff}
	LightBlue400   = Clr{0x29, 0xb6, 0xf6, 0xff}
	LightBlue500   = Clr{0x03, 0xa9, 0xf4, 0xff}
	LightBlue600   = Clr{0x03, 0x9b, 0xe5, 0xff}
	LightBlue700   = Clr{0x02, 0x88, 0xd1, 0xff}
	LightBlue800   = Clr{0x02, 0x77, 0xbd, 0xff}
	LightBlue900   = Clr{0x01, 0x57, 0x9b, 0xff}
	LightBlueA100  = Clr{0x80, 0xd8, 0xff, 0xff}
	LightBlueA200  = Clr{0x40, 0xc4, 0xff, 0xff}
	LightBlueA400  = Clr{0x00, 0xb0, 0xff, 0xff}
	LightBlueA700  = Clr{0x00, 0x91, 0xea, 0xff}
	Cyan50         = Clr{0xe0, 0xf7, 0xfa, 0xff}
	Cyan100        = Clr{0xb2, 0xeb, 0xf2, 0xff}
	Cyan200        = Clr{0x80, 0xde, 0xea, 0xff}
	Cyan300        = Clr{0x4d, 0xd0, 0xe1, 0xff}
	Cyan400        = Clr{0x26, 0xc6, 0xda, 0xff}
	Cyan500        = Clr{0x00, 0xbc, 0xd4, 0xff}
	Cyan600        = Clr{0x00, 0xac, 0xc1, 0xff}
	Cyan700        = Clr{0x00, 0x97, 0xa7, 0xff}
	Cyan800        = Clr{0x00, 0x83, 0x8f, 0xff}
	Cyan900        = Clr{0x00, 0x60, 0x64, 0xff}
	CyanA100       = Clr{0x84, 0xff, 0xff, 0xff}
	CyanA200       = Clr{0x18, 0xff, 0xff, 0xff}
	CyanA400       = Clr{0x00, 0xe5, 0xff, 0xff}
	CyanA700       = Clr{0x00, 0xb8, 0xd4, 0xff}
	Teal50         = Clr{0xe0, 0xf2, 0xf1, 0xff}
	Teal100        = Clr{0xb2, 0xdf, 0xdb, 0xff}
	Teal200        = Clr{0x80, 0xcb, 0xc4, 0xff}
	Teal300        = Clr{0x4d, 0xb6, 0xac, 0xff}
	Teal400        = Clr{0x26, 0xa6, 0x9a, 0xff}
	Teal500        = Clr{0x00, 0x96, 0x88, 0xff}
	Teal600        = Clr{0x00, 0x89, 0x7b, 0xff}
	Teal700        = Clr{0x00, 0x79, 0x6b, 0xff}
	Teal800        = Clr{0x00, 0x69, 0x5c, 0xff}
	Teal900        = Clr{0x00, 0x4d, 0x40, 0xff}
	TealA100       = Clr{0xa7, 0xff, 0xeb, 0xff}
	TealA200       = Clr{0x64, 0xff, 0xda, 0xff}
	TealA400       = Clr{0x1d, 0xe9, 0xb6, 0xff}
	TealA700       = Clr{0x00, 0xbf, 0xa5, 0xff}
	Green50        = Clr{0xe8, 0xf5, 0xe9, 0xff}
	Green100       = Clr{0xc8, 0xe6, 0xc9, 0xff}
	Green200       = Clr{0xa5, 0xd6, 0xa7, 0xff}
	Green300       = Clr{0x81, 0xc7, 0x84, 0xff}
	Green400       = Clr{0x66, 0xbb, 0x6a, 0xff}
	Green500       = Clr{0x4c, 0xaf, 0x50, 0xff}
	Green600       = Clr{0x43, 0xa0, 0x47, 0xff}
	Green700       = Clr{0x38, 0x8e, 0x3c, 0xff}
	Green800       = Clr{0x2e, 0x7d, 0x32, 0xff}
	Green900       = Clr{0x1b, 0x5e, 0x20, 0xff}
	GreenA100      = Clr{0xb9, 0xf6, 0xca, 0xff}
	GreenA200      = Clr{0x69, 0xf0, 0xae, 0xff}
	GreenA400      = Clr{0x00, 0xe6, 0x76, 0xff}
	GreenA700      = Clr{0x00, 0xc8, 0x53, 0xff}
	LightGreen50   = Clr{0xf1, 0xf8, 0xe9, 0xff}
	LightGreen100  = Clr{0xdc, 0xed, 0xc8, 0xff}
	LightGreen200  = Clr{0xc5, 0xe1, 0xa5, 0xff}
	LightGreen300  = Clr{0xae, 0xd5, 0x81, 0xff}
	LightGreen400  = Clr{0x9c, 0xcc, 0x65, 0xff}
	LightGreen500  = Clr{0x8b, 0xc3, 0x4a, 0xff}
	LightGreen600  = Clr{0x7c, 0xb3, 0x42, 0xff}
	LightGreen700  = Clr{0x68, 0x9f, 0x38, 0xff}
	LightGreen800  = Clr{0x55, 0x8b, 0x2f, 0xff}
	LightGreen900  = Clr{0x33, 0x69, 0x1e, 0xff}
	LightGreenA100 = Clr{0xcc, 0xff, 0x90, 0xff}
	LightGreenA200 = Clr{0xb2, 0xff, 0x59, 0xff}
	LightGreenA400 = Clr{0x76, 0xff, 0x03, 0xff}
	LightGreenA700 = Clr{0x64, 0xdd, 0x17, 0xff}
	Lime50         = Clr{0xf9, 0xfb, 0xe7, 0xff}
	Lime100        = Clr{0xf0, 0xf4, 0xc3, 0xff}
	Lime200        = Clr{0xe6, 0xee, 0x9c, 0xff}
	Lime300        = Clr{0xdc, 0xe7, 0x75, 0xff}
	Lime400        = Clr{0xd4, 0xe1, 0x57, 0xff}
	Lime500        = Clr{0xcd, 0xdc, 0x39, 0xff}
	Lime600        = Clr{0xc0, 0xca, 0x33, 0xff}
	Lime700        = Clr{0xaf, 0xb4, 0x2b, 0xff}
	Lime800        = Clr{0x9e, 0x9d, 0x24, 0xff}
	Lime900        = Clr{0x82, 0x77, 0x17, 0xff}
	LimeA100       = Clr{0xf4, 0xff, 0x81, 0xff}
	LimeA200       = Clr{0xee, 0xff, 0x41, 0xff}
	LimeA400       = Clr{0xc6, 0xff, 0x00, 0xff}
	LimeA700       = Clr{0xae, 0xea, 0x00, 0xff}
	Yellow50       = Clr{0xff, 0xfd, 0xe7, 0xff}
	Yellow100      = Clr{0xff, 0xf9, 0xc4, 0xff}
	Yellow200      = Clr{0xff, 0xf5, 0x9d, 0xff}
	Yellow300      = Clr{0xff, 0xf1, 0x76, 0xff}
	Yellow400      = Clr{0xff, 0xee, 0x58, 0xff}
	Yellow500      = Clr{0xff, 0xeb, 0x3b, 0xff}
	Yellow600      = Clr{0xfd, 0xd8, 0x35, 0xff}
	Yellow700      = Clr{0xfb, 0xc0, 0x2d, 0xff}
	Yellow800      = Clr{0xf9, 0xa8, 0x25, 0xff}
	Yellow900      = Clr{0xf5, 0x7f, 0x17, 0xff}
	YellowA100     = Clr{0xff, 0xff, 0x8d, 0xff}
	YellowA200     = Clr{0xff, 0xff, 0x00, 0xff}
	YellowA400     = Clr{0xff, 0xea, 0x00, 0xff}
	YellowA700     = Clr{0xff, 0xd6, 0x00, 0xff}
	Amber50        = Clr{0xff, 0xf8, 0xe1, 0xff}
	Amber100       = Clr{0xff, 0xec, 0xb3, 0xff}
	Amber200       = Clr{0xff, 0xe0, 0x82, 0xff}
	Amber300       = Clr{0xff, 0xd5, 0x4f, 0xff}
	Amber400       = Clr{0xff, 0xca, 0x28, 0xff}
	Amber500       = Clr{0xff, 0xc1, 0x07, 0xff}
	Amber600       = Clr{0xff, 0xb3, 0x00, 0xff}
	Amber700       = Clr{0xff, 0xa0, 0x00, 0xff}
	Amber800       = Clr{0xff, 0x8f, 0x00, 0xff}
	Amber900       = Clr{0xff, 0x6f, 0x00, 0xff}
	AmberA100      = Clr{0xff, 0xe5, 0x7f, 0xff}
	AmberA200      = Clr{0xff, 0xd7, 0x40, 0xff}
	AmberA400      = Clr{0xff, 0xc4, 0x00, 0xff}
	AmberA700      = Clr{0xff, 0xab, 0x00, 0xff}
	Orange50       = Clr{0xff, 0xf3, 0xe0, 0xff}
	Orange100      = Clr{0xff, 0xe0, 0xb2, 0xff}
	Orange200      = Clr{0xff, 0xcc, 0x80, 0xff}
	Orange300      = Clr{0xff, 0xb7, 0x4d, 0xff}
	Orange400      = Clr{0xff, 0xa7, 0x26, 0xff}
	Orange500      = Clr{0xff, 0x98, 0x00, 0xff}
	Orange600      = Clr{0xfb, 0x8c, 0x00, 0xff}
	Orange700      = Clr{0xf5, 0x7c, 0x00, 0xff}
	Orange800      = Clr{0xef, 0x6c, 0x00, 0xff}
	Orange900      = Clr{0xe6, 0x51, 0x00, 0xff}
	OrangeA100     = Clr{0xff, 0xd1, 0x80, 0xff}
	OrangeA200     = Clr{0xff, 0xab, 0x40, 0xff}
	OrangeA400     = Clr{0xff, 0x91, 0x00, 0xff}
	OrangeA700     = Clr{0xff, 0x6d, 0x00, 0xff}
	DeepOrange50   = Clr{0xfb, 0xe9, 0xe7, 0xff}
	DeepOrange100  = Clr{0xff, 0xcc, 0xbc, 0xff}
	DeepOrange200  = Clr{0xff, 0xab, 0x91, 0xff}
	DeepOrange300  = Clr{0xff, 0x8a, 0x65, 0xff}
	DeepOrange400  = Clr{0xff, 0x70, 0x43, 0xff}
	DeepOrange500  = Clr{0xff, 0x57, 0x22, 0xff}
	DeepOrange600  = Clr{0xf4, 0x51, 0x1e, 0xff}
	DeepOrange700  = Clr{0xe6, 0x4a, 0x19, 0xff}
	DeepOrange800  = Clr{0xd8, 0x43, 0x15, 0xff}
	DeepOrange900  = Clr{0xbf, 0x36, 0x0c, 0xff}
	DeepOrangeA100 = Clr{0xff, 0x9e, 0x80, 0xff}
	DeepOrangeA200 = Clr{0xff, 0x6e, 0x40, 0xff}
	DeepOrangeA400 = Clr{0xff, 0x3d, 0x00, 0xff}
	DeepOrangeA700 = Clr{0xdd, 0x2c, 0x00, 0xff}
	Brown50        = Clr{0xef, 0xeb, 0xe9, 0xff}
	Brown100       = Clr{0xd7, 0xcc, 0xc8, 0xff}
	Brown200       = Clr{0xbc, 0xaa, 0xa4, 0xff}
	Brown300       = Clr{0xa1, 0x88, 0x7f, 0xff}
	Brown400       = Clr{0x8d, 0x6e, 0x63, 0xff}
	Brown500       = Clr{0x79, 0x55, 0x48, 0xff}
	Brown600       = Clr{0x6d, 0x4c, 0x41, 0xff}
	Brown700       = Clr{0x5d, 0x40, 0x37, 0xff}
	Brown800       = Clr{0x4e, 0x34, 0x2e, 0xff}
	Brown900       = Clr{0x3e, 0x27, 0x23, 0xff}
	Grey50         = Clr{0xfa, 0xfa, 0xfa, 0xff}
	Grey100        = Clr{0xf5, 0xf5, 0xf5, 0xff}
	Grey200        = Clr{0xee, 0xee, 0xee, 0xff}
	Grey300        = Clr{0xe0, 0xe0, 0xe0, 0xff}
	Grey400        = Clr{0xbd, 0xbd, 0xbd, 0xff}
	Grey500        = Clr{0x9e, 0x9e, 0x9e, 0xff}
	Grey600        = Clr{0x75, 0x75, 0x75, 0xff}
	Grey700        = Clr{0x61, 0x61, 0x61, 0xff}
	Grey800        = Clr{0x42, 0x42, 0x42, 0xff}
	Grey900        = Clr{0x21, 0x21, 0x21, 0xff}
	BlueGrey50     = Clr{0xec, 0xef, 0xf1, 0xff}
	BlueGrey100    = Clr{0xcf, 0xd8, 0xdc, 0xff}
	BlueGrey200    = Clr{0xb0, 0xbe, 0xc5, 0xff}
	BlueGrey300    = Clr{0x90, 0xa4, 0xae, 0xff}
	BlueGrey400    = Clr{0x78, 0x90, 0x9c, 0xff}
	BlueGrey500    = Clr{0x60, 0x7d, 0x8b, 0xff}
	BlueGrey600    = Clr{0x54, 0x6e, 0x7a, 0xff}
	BlueGrey700    = Clr{0x45, 0x5a, 0x64, 0xff}
	BlueGrey800    = Clr{0x37, 0x47, 0x4f, 0xff}
	BlueGrey900    = Clr{0x26, 0x32, 0x38, 0xff}
)

type (
	Clr    color.RGBA
	ClrScp struct {
		Idx uint32
		Arr []Clr
	}
)

func Rgba(r, g, b, a flt.Flt) (c Clr) {
	c.R = uint8(r * 255)
	c.G = uint8(g * 255)
	c.B = uint8(b * 255)
	c.A = uint8(a * 255)
	return c
}
func Rgb(r, g, b flt.Flt) (c Clr) {
	c.R = uint8(r * 255)
	c.G = uint8(g * 255)
	c.B = uint8(b * 255)
	c.A = 255
	return c
}
func Hex(txt str.Str) (c Clr) {
	s := strings.TrimSpace(string(txt))
	s = strings.TrimPrefix(s, "#")
	switch {
	case len(s) == 3:
		fmt.Sscanf(s, "%1x%1x%1x", &c.R, &c.G, &c.B)
		c.A = 255
	case len(s) == 6:
		fmt.Sscanf(s, "%02x%02x%02x", &c.R, &c.G, &c.B)
		c.A = 255
	case len(s) == 8:
		fmt.Sscanf(s, "%02x%02x%02x%02x", &c.R, &c.G, &c.B, &c.A)
	}
	return c
}
func (x Clr) Opa(pct flt.Flt) Clr {
	x.A = uint8(pct * 255)
	return x
}
func (x Clr) Inv() Clr {
	x.R = 255 - x.R
	x.G = 255 - x.G
	x.B = 255 - x.B
	return x
}
