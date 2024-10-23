package tpl

import (
	"sys/k"
	"sys/tpl/atr"
)

type (
	DirClr struct {
		DirBse
		Clr *FleClr
	}
	FleClr struct {
		FleBse
	}
)

func (x *DirVis) NewClr() (r *DirClr) { // dir
	r = &DirClr{}
	x.Clr = r
	r.Pkg = x.Pkg.New(k.Clr)
	r.NewClr()
	return r
}
func (x *DirClr) NewClr() (r *FleClr) { // fle
	r = &FleClr{}
	x.Clr = r
	r.Name = k.Clr
	r.Pkg = x.Pkg
	r.Alias(r.Name, NewExt("color.RGBA"), atr.TypUiStruct) // Struct for Var declaration (even though alias to struct)
	r.AddFle(r)
	r.Import("image/color")
	return r
}
func (x *FleClr) InitPkgFn() {
	x.Rgba()
	x.Rgb()
	x.Hex()
}
func (x *FleClr) Rgba() (r *PkgFn) {
	r = x.PkgFna(k.Rgba, atr.Lng)
	r.InPrm(_sys.Bsc.Flt, "r")
	r.InPrm(_sys.Bsc.Flt, "g")
	r.InPrm(_sys.Bsc.Flt, "b")
	r.InPrm(_sys.Bsc.Flt, "a")
	r.OutPrm(x, "c")
	r.Add("c.R = uint8(r * 255)")
	r.Add("c.G = uint8(g * 255)")
	r.Add("c.B = uint8(b * 255)")
	r.Add("c.A = uint8(a * 255)")
	r.Add("return c")
	return r
}
func (x *FleClr) Rgb() (r *PkgFn) {
	r = x.PkgFna(k.Rgb, atr.Lng)
	r.InPrm(_sys.Bsc.Flt, "r")
	r.InPrm(_sys.Bsc.Flt, "g")
	r.InPrm(_sys.Bsc.Flt, "b")
	r.OutPrm(x, "c")
	r.Add("c.R = uint8(r * 255)")
	r.Add("c.G = uint8(g * 255)")
	r.Add("c.B = uint8(b * 255)")
	r.Add("c.A = 255")
	r.Add("return c")
	return r
}
func (x *FleClr) Hex() (r *PkgFn) {
	x.Import("fmt")
	x.Import("strings")
	r = x.PkgFna(k.Hex, atr.Lng)
	r.InPrm(_sys.Bsc.Str, "txt")
	r.OutPrm(x, "c")
	r.Add("s := strings.TrimSpace(string(txt))")
	r.Add("s = strings.TrimPrefix(s, \"#\")")
	r.Add("switch {")
	r.Add("case len(s) == 3:")
	r.Add("fmt.Sscanf(s, \"%1x%1x%1x\", &c.R, &c.G, &c.B)")
	r.Add("c.A = 255")
	r.Add("case len(s) == 6:")
	r.Add("fmt.Sscanf(s, \"%02x%02x%02x\", &c.R, &c.G, &c.B)")
	r.Add("c.A = 255")
	r.Add("case len(s) == 8:")
	r.Add("fmt.Sscanf(s, \"%02x%02x%02x%02x\", &c.R, &c.G, &c.B, &c.A)")
	r.Add("}")
	r.Add("return c")
	return r
}
func (x *FleClr) InitTypFn() {
	x.Opa()
	x.Inv()
}
func (x *FleClr) Opa() (r *TypFn) {
	r = x.TypFn(k.Opa) // set opacity
	r.InPrm(_sys.Bsc.Flt, "pct")
	r.OutPrm(x)
	r.Add("x.A = uint8(pct * 255)")
	r.Add("return x")
	return r
}
func (x *FleClr) Inv() (r *TypFn) {
	r = x.TypFn(k.Inv) // inverse
	r.OutPrm(x)
	r.Add("x.R = 255 - x.R")
	r.Add("x.G = 255 - x.G")
	r.Add("x.B = 255 - x.B")
	r.Add("return x")
	return r
}
func (x *FleClr) InitVar() {
	// https://godoc.org/golang.org/x/exp/shiny/materialdesign/colornames
	x.Var("Black", "0x00, 0x00, 0x00, 0xff")
	x.Var("White", "0xff, 0xff, 0xff, 0xff")
	x.Var("Red50", "0xff, 0xeb, 0xee, 0xff")
	x.Var("Red100", "0xff, 0xcd, 0xd2, 0xff")
	x.Var("Red200", "0xef, 0x9a, 0x9a, 0xff")
	x.Var("Red300", "0xe5, 0x73, 0x73, 0xff")
	x.Var("Red400", "0xef, 0x53, 0x50, 0xff")
	x.Var("Red500", "0xf4, 0x43, 0x36, 0xff")
	x.Var("Red600", "0xe5, 0x39, 0x35, 0xff")
	x.Var("Red700", "0xd3, 0x2f, 0x2f, 0xff")
	x.Var("Red800", "0xc6, 0x28, 0x28, 0xff")
	x.Var("Red900", "0xb7, 0x1c, 0x1c, 0xff")
	x.Var("RedA100", "0xff, 0x8a, 0x80, 0xff")
	x.Var("RedA200", "0xff, 0x52, 0x52, 0xff")
	x.Var("RedA400", "0xff, 0x17, 0x44, 0xff")
	x.Var("RedA700", "0xd5, 0x00, 0x00, 0xff")
	x.Var("Pink50", "0xfc, 0xe4, 0xec, 0xff")
	x.Var("Pink100", "0xf8, 0xbb, 0xd0, 0xff")
	x.Var("Pink200", "0xf4, 0x8f, 0xb1, 0xff")
	x.Var("Pink300", "0xf0, 0x62, 0x92, 0xff")
	x.Var("Pink400", "0xec, 0x40, 0x7a, 0xff")
	x.Var("Pink500", "0xe9, 0x1e, 0x63, 0xff")
	x.Var("Pink600", "0xd8, 0x1b, 0x60, 0xff")
	x.Var("Pink700", "0xc2, 0x18, 0x5b, 0xff")
	x.Var("Pink800", "0xad, 0x14, 0x57, 0xff")
	x.Var("Pink900", "0x88, 0x0e, 0x4f, 0xff")
	x.Var("PinkA100", "0xff, 0x80, 0xab, 0xff")
	x.Var("PinkA200", "0xff, 0x40, 0x81, 0xff")
	x.Var("PinkA400", "0xf5, 0x00, 0x57, 0xff")
	x.Var("PinkA700", "0xc5, 0x11, 0x62, 0xff")
	x.Var("Purple50", "0xf3, 0xe5, 0xf5, 0xff")
	x.Var("Purple100", "0xe1, 0xbe, 0xe7, 0xff")
	x.Var("Purple200", "0xce, 0x93, 0xd8, 0xff")
	x.Var("Purple300", "0xba, 0x68, 0xc8, 0xff")
	x.Var("Purple400", "0xab, 0x47, 0xbc, 0xff")
	x.Var("Purple500", "0x9c, 0x27, 0xb0, 0xff")
	x.Var("Purple600", "0x8e, 0x24, 0xaa, 0xff")
	x.Var("Purple700", "0x7b, 0x1f, 0xa2, 0xff")
	x.Var("Purple800", "0x6a, 0x1b, 0x9a, 0xff")
	x.Var("Purple900", "0x4a, 0x14, 0x8c, 0xff")
	x.Var("PurpleA100", "0xea, 0x80, 0xfc, 0xff")
	x.Var("PurpleA200", "0xe0, 0x40, 0xfb, 0xff")
	x.Var("PurpleA400", "0xd5, 0x00, 0xf9, 0xff")
	x.Var("PurpleA700", "0xaa, 0x00, 0xff, 0xff")
	x.Var("DeepPurple50", "0xed, 0xe7, 0xf6, 0xff")
	x.Var("DeepPurple100", "0xd1, 0xc4, 0xe9, 0xff")
	x.Var("DeepPurple200", "0xb3, 0x9d, 0xdb, 0xff")
	x.Var("DeepPurple300", "0x95, 0x75, 0xcd, 0xff")
	x.Var("DeepPurple400", "0x7e, 0x57, 0xc2, 0xff")
	x.Var("DeepPurple500", "0x67, 0x3a, 0xb7, 0xff")
	x.Var("DeepPurple600", "0x5e, 0x35, 0xb1, 0xff")
	x.Var("DeepPurple700", "0x51, 0x2d, 0xa8, 0xff")
	x.Var("DeepPurple800", "0x45, 0x27, 0xa0, 0xff")
	x.Var("DeepPurple900", "0x31, 0x1b, 0x92, 0xff")
	x.Var("DeepPurpleA100", "0xb3, 0x88, 0xff, 0xff")
	x.Var("DeepPurpleA200", "0x7c, 0x4d, 0xff, 0xff")
	x.Var("DeepPurpleA400", "0x65, 0x1f, 0xff, 0xff")
	x.Var("DeepPurpleA700", "0x62, 0x00, 0xea, 0xff")
	x.Var("Indigo50", "0xe8, 0xea, 0xf6, 0xff")
	x.Var("Indigo100", "0xc5, 0xca, 0xe9, 0xff")
	x.Var("Indigo200", "0x9f, 0xa8, 0xda, 0xff")
	x.Var("Indigo300", "0x79, 0x86, 0xcb, 0xff")
	x.Var("Indigo400", "0x5c, 0x6b, 0xc0, 0xff")
	x.Var("Indigo500", "0x3f, 0x51, 0xb5, 0xff")
	x.Var("Indigo600", "0x39, 0x49, 0xab, 0xff")
	x.Var("Indigo700", "0x30, 0x3f, 0x9f, 0xff")
	x.Var("Indigo800", "0x28, 0x35, 0x93, 0xff")
	x.Var("Indigo900", "0x1a, 0x23, 0x7e, 0xff")
	x.Var("IndigoA100", "0x8c, 0x9e, 0xff, 0xff")
	x.Var("IndigoA200", "0x53, 0x6d, 0xfe, 0xff")
	x.Var("IndigoA400", "0x3d, 0x5a, 0xfe, 0xff")
	x.Var("IndigoA700", "0x30, 0x4f, 0xfe, 0xff")
	x.Var("Blue50", "0xe3, 0xf2, 0xfd, 0xff")
	x.Var("Blue100", "0xbb, 0xde, 0xfb, 0xff")
	x.Var("Blue200", "0x90, 0xca, 0xf9, 0xff")
	x.Var("Blue300", "0x64, 0xb5, 0xf6, 0xff")
	x.Var("Blue400", "0x42, 0xa5, 0xf5, 0xff")
	x.Var("Blue500", "0x21, 0x96, 0xf3, 0xff")
	x.Var("Blue600", "0x1e, 0x88, 0xe5, 0xff")
	x.Var("Blue700", "0x19, 0x76, 0xd2, 0xff")
	x.Var("Blue800", "0x15, 0x65, 0xc0, 0xff")
	x.Var("Blue900", "0x0d, 0x47, 0xa1, 0xff")
	x.Var("BlueA100", "0x82, 0xb1, 0xff, 0xff")
	x.Var("BlueA200", "0x44, 0x8a, 0xff, 0xff")
	x.Var("BlueA400", "0x29, 0x79, 0xff, 0xff")
	x.Var("BlueA700", "0x29, 0x62, 0xff, 0xff")
	x.Var("LightBlue50", "0xe1, 0xf5, 0xfe, 0xff")
	x.Var("LightBlue100", "0xb3, 0xe5, 0xfc, 0xff")
	x.Var("LightBlue200", "0x81, 0xd4, 0xfa, 0xff")
	x.Var("LightBlue300", "0x4f, 0xc3, 0xf7, 0xff")
	x.Var("LightBlue400", "0x29, 0xb6, 0xf6, 0xff")
	x.Var("LightBlue500", "0x03, 0xa9, 0xf4, 0xff")
	x.Var("LightBlue600", "0x03, 0x9b, 0xe5, 0xff")
	x.Var("LightBlue700", "0x02, 0x88, 0xd1, 0xff")
	x.Var("LightBlue800", "0x02, 0x77, 0xbd, 0xff")
	x.Var("LightBlue900", "0x01, 0x57, 0x9b, 0xff")
	x.Var("LightBlueA100", "0x80, 0xd8, 0xff, 0xff")
	x.Var("LightBlueA200", "0x40, 0xc4, 0xff, 0xff")
	x.Var("LightBlueA400", "0x00, 0xb0, 0xff, 0xff")
	x.Var("LightBlueA700", "0x00, 0x91, 0xea, 0xff")
	x.Var("Cyan50", "0xe0, 0xf7, 0xfa, 0xff")
	x.Var("Cyan100", "0xb2, 0xeb, 0xf2, 0xff")
	x.Var("Cyan200", "0x80, 0xde, 0xea, 0xff")
	x.Var("Cyan300", "0x4d, 0xd0, 0xe1, 0xff")
	x.Var("Cyan400", "0x26, 0xc6, 0xda, 0xff")
	x.Var("Cyan500", "0x00, 0xbc, 0xd4, 0xff")
	x.Var("Cyan600", "0x00, 0xac, 0xc1, 0xff")
	x.Var("Cyan700", "0x00, 0x97, 0xa7, 0xff")
	x.Var("Cyan800", "0x00, 0x83, 0x8f, 0xff")
	x.Var("Cyan900", "0x00, 0x60, 0x64, 0xff")
	x.Var("CyanA100", "0x84, 0xff, 0xff, 0xff")
	x.Var("CyanA200", "0x18, 0xff, 0xff, 0xff")
	x.Var("CyanA400", "0x00, 0xe5, 0xff, 0xff")
	x.Var("CyanA700", "0x00, 0xb8, 0xd4, 0xff")
	x.Var("Teal50", "0xe0, 0xf2, 0xf1, 0xff")
	x.Var("Teal100", "0xb2, 0xdf, 0xdb, 0xff")
	x.Var("Teal200", "0x80, 0xcb, 0xc4, 0xff")
	x.Var("Teal300", "0x4d, 0xb6, 0xac, 0xff")
	x.Var("Teal400", "0x26, 0xa6, 0x9a, 0xff")
	x.Var("Teal500", "0x00, 0x96, 0x88, 0xff")
	x.Var("Teal600", "0x00, 0x89, 0x7b, 0xff")
	x.Var("Teal700", "0x00, 0x79, 0x6b, 0xff")
	x.Var("Teal800", "0x00, 0x69, 0x5c, 0xff")
	x.Var("Teal900", "0x00, 0x4d, 0x40, 0xff")
	x.Var("TealA100", "0xa7, 0xff, 0xeb, 0xff")
	x.Var("TealA200", "0x64, 0xff, 0xda, 0xff")
	x.Var("TealA400", "0x1d, 0xe9, 0xb6, 0xff")
	x.Var("TealA700", "0x00, 0xbf, 0xa5, 0xff")
	x.Var("Green50", "0xe8, 0xf5, 0xe9, 0xff")
	x.Var("Green100", "0xc8, 0xe6, 0xc9, 0xff")
	x.Var("Green200", "0xa5, 0xd6, 0xa7, 0xff")
	x.Var("Green300", "0x81, 0xc7, 0x84, 0xff")
	x.Var("Green400", "0x66, 0xbb, 0x6a, 0xff")
	x.Var("Green500", "0x4c, 0xaf, 0x50, 0xff")
	x.Var("Green600", "0x43, 0xa0, 0x47, 0xff")
	x.Var("Green700", "0x38, 0x8e, 0x3c, 0xff")
	x.Var("Green800", "0x2e, 0x7d, 0x32, 0xff")
	x.Var("Green900", "0x1b, 0x5e, 0x20, 0xff")
	x.Var("GreenA100", "0xb9, 0xf6, 0xca, 0xff")
	x.Var("GreenA200", "0x69, 0xf0, 0xae, 0xff")
	x.Var("GreenA400", "0x00, 0xe6, 0x76, 0xff")
	x.Var("GreenA700", "0x00, 0xc8, 0x53, 0xff")
	x.Var("LightGreen50", "0xf1, 0xf8, 0xe9, 0xff")
	x.Var("LightGreen100", "0xdc, 0xed, 0xc8, 0xff")
	x.Var("LightGreen200", "0xc5, 0xe1, 0xa5, 0xff")
	x.Var("LightGreen300", "0xae, 0xd5, 0x81, 0xff")
	x.Var("LightGreen400", "0x9c, 0xcc, 0x65, 0xff")
	x.Var("LightGreen500", "0x8b, 0xc3, 0x4a, 0xff")
	x.Var("LightGreen600", "0x7c, 0xb3, 0x42, 0xff")
	x.Var("LightGreen700", "0x68, 0x9f, 0x38, 0xff")
	x.Var("LightGreen800", "0x55, 0x8b, 0x2f, 0xff")
	x.Var("LightGreen900", "0x33, 0x69, 0x1e, 0xff")
	x.Var("LightGreenA100", "0xcc, 0xff, 0x90, 0xff")
	x.Var("LightGreenA200", "0xb2, 0xff, 0x59, 0xff")
	x.Var("LightGreenA400", "0x76, 0xff, 0x03, 0xff")
	x.Var("LightGreenA700", "0x64, 0xdd, 0x17, 0xff")
	x.Var("Lime50", "0xf9, 0xfb, 0xe7, 0xff")
	x.Var("Lime100", "0xf0, 0xf4, 0xc3, 0xff")
	x.Var("Lime200", "0xe6, 0xee, 0x9c, 0xff")
	x.Var("Lime300", "0xdc, 0xe7, 0x75, 0xff")
	x.Var("Lime400", "0xd4, 0xe1, 0x57, 0xff")
	x.Var("Lime500", "0xcd, 0xdc, 0x39, 0xff")
	x.Var("Lime600", "0xc0, 0xca, 0x33, 0xff")
	x.Var("Lime700", "0xaf, 0xb4, 0x2b, 0xff")
	x.Var("Lime800", "0x9e, 0x9d, 0x24, 0xff")
	x.Var("Lime900", "0x82, 0x77, 0x17, 0xff")
	x.Var("LimeA100", "0xf4, 0xff, 0x81, 0xff")
	x.Var("LimeA200", "0xee, 0xff, 0x41, 0xff")
	x.Var("LimeA400", "0xc6, 0xff, 0x00, 0xff")
	x.Var("LimeA700", "0xae, 0xea, 0x00, 0xff")
	x.Var("Yellow50", "0xff, 0xfd, 0xe7, 0xff")
	x.Var("Yellow100", "0xff, 0xf9, 0xc4, 0xff")
	x.Var("Yellow200", "0xff, 0xf5, 0x9d, 0xff")
	x.Var("Yellow300", "0xff, 0xf1, 0x76, 0xff")
	x.Var("Yellow400", "0xff, 0xee, 0x58, 0xff")
	x.Var("Yellow500", "0xff, 0xeb, 0x3b, 0xff")
	x.Var("Yellow600", "0xfd, 0xd8, 0x35, 0xff")
	x.Var("Yellow700", "0xfb, 0xc0, 0x2d, 0xff")
	x.Var("Yellow800", "0xf9, 0xa8, 0x25, 0xff")
	x.Var("Yellow900", "0xf5, 0x7f, 0x17, 0xff")
	x.Var("YellowA100", "0xff, 0xff, 0x8d, 0xff")
	x.Var("YellowA200", "0xff, 0xff, 0x00, 0xff")
	x.Var("YellowA400", "0xff, 0xea, 0x00, 0xff")
	x.Var("YellowA700", "0xff, 0xd6, 0x00, 0xff")
	x.Var("Amber50", "0xff, 0xf8, 0xe1, 0xff")
	x.Var("Amber100", "0xff, 0xec, 0xb3, 0xff")
	x.Var("Amber200", "0xff, 0xe0, 0x82, 0xff")
	x.Var("Amber300", "0xff, 0xd5, 0x4f, 0xff")
	x.Var("Amber400", "0xff, 0xca, 0x28, 0xff")
	x.Var("Amber500", "0xff, 0xc1, 0x07, 0xff")
	x.Var("Amber600", "0xff, 0xb3, 0x00, 0xff")
	x.Var("Amber700", "0xff, 0xa0, 0x00, 0xff")
	x.Var("Amber800", "0xff, 0x8f, 0x00, 0xff")
	x.Var("Amber900", "0xff, 0x6f, 0x00, 0xff")
	x.Var("AmberA100", "0xff, 0xe5, 0x7f, 0xff")
	x.Var("AmberA200", "0xff, 0xd7, 0x40, 0xff")
	x.Var("AmberA400", "0xff, 0xc4, 0x00, 0xff")
	x.Var("AmberA700", "0xff, 0xab, 0x00, 0xff")
	x.Var("Orange50", "0xff, 0xf3, 0xe0, 0xff")
	x.Var("Orange100", "0xff, 0xe0, 0xb2, 0xff")
	x.Var("Orange200", "0xff, 0xcc, 0x80, 0xff")
	x.Var("Orange300", "0xff, 0xb7, 0x4d, 0xff")
	x.Var("Orange400", "0xff, 0xa7, 0x26, 0xff")
	x.Var("Orange500", "0xff, 0x98, 0x00, 0xff")
	x.Var("Orange600", "0xfb, 0x8c, 0x00, 0xff")
	x.Var("Orange700", "0xf5, 0x7c, 0x00, 0xff")
	x.Var("Orange800", "0xef, 0x6c, 0x00, 0xff")
	x.Var("Orange900", "0xe6, 0x51, 0x00, 0xff")
	x.Var("OrangeA100", "0xff, 0xd1, 0x80, 0xff")
	x.Var("OrangeA200", "0xff, 0xab, 0x40, 0xff")
	x.Var("OrangeA400", "0xff, 0x91, 0x00, 0xff")
	x.Var("OrangeA700", "0xff, 0x6d, 0x00, 0xff")
	x.Var("DeepOrange50", "0xfb, 0xe9, 0xe7, 0xff")
	x.Var("DeepOrange100", "0xff, 0xcc, 0xbc, 0xff")
	x.Var("DeepOrange200", "0xff, 0xab, 0x91, 0xff")
	x.Var("DeepOrange300", "0xff, 0x8a, 0x65, 0xff")
	x.Var("DeepOrange400", "0xff, 0x70, 0x43, 0xff")
	x.Var("DeepOrange500", "0xff, 0x57, 0x22, 0xff")
	x.Var("DeepOrange600", "0xf4, 0x51, 0x1e, 0xff")
	x.Var("DeepOrange700", "0xe6, 0x4a, 0x19, 0xff")
	x.Var("DeepOrange800", "0xd8, 0x43, 0x15, 0xff")
	x.Var("DeepOrange900", "0xbf, 0x36, 0x0c, 0xff")
	x.Var("DeepOrangeA100", "0xff, 0x9e, 0x80, 0xff")
	x.Var("DeepOrangeA200", "0xff, 0x6e, 0x40, 0xff")
	x.Var("DeepOrangeA400", "0xff, 0x3d, 0x00, 0xff")
	x.Var("DeepOrangeA700", "0xdd, 0x2c, 0x00, 0xff")
	x.Var("Brown50", "0xef, 0xeb, 0xe9, 0xff")
	x.Var("Brown100", "0xd7, 0xcc, 0xc8, 0xff")
	x.Var("Brown200", "0xbc, 0xaa, 0xa4, 0xff")
	x.Var("Brown300", "0xa1, 0x88, 0x7f, 0xff")
	x.Var("Brown400", "0x8d, 0x6e, 0x63, 0xff")
	x.Var("Brown500", "0x79, 0x55, 0x48, 0xff")
	x.Var("Brown600", "0x6d, 0x4c, 0x41, 0xff")
	x.Var("Brown700", "0x5d, 0x40, 0x37, 0xff")
	x.Var("Brown800", "0x4e, 0x34, 0x2e, 0xff")
	x.Var("Brown900", "0x3e, 0x27, 0x23, 0xff")
	x.Var("Grey50", "0xfa, 0xfa, 0xfa, 0xff")
	x.Var("Grey100", "0xf5, 0xf5, 0xf5, 0xff")
	x.Var("Grey200", "0xee, 0xee, 0xee, 0xff")
	x.Var("Grey300", "0xe0, 0xe0, 0xe0, 0xff")
	x.Var("Grey400", "0xbd, 0xbd, 0xbd, 0xff")
	x.Var("Grey500", "0x9e, 0x9e, 0x9e, 0xff")
	x.Var("Grey600", "0x75, 0x75, 0x75, 0xff")
	x.Var("Grey700", "0x61, 0x61, 0x61, 0xff")
	x.Var("Grey800", "0x42, 0x42, 0x42, 0xff")
	x.Var("Grey900", "0x21, 0x21, 0x21, 0xff")
	x.Var("BlueGrey50", "0xec, 0xef, 0xf1, 0xff")
	x.Var("BlueGrey100", "0xcf, 0xd8, 0xdc, 0xff")
	x.Var("BlueGrey200", "0xb0, 0xbe, 0xc5, 0xff")
	x.Var("BlueGrey300", "0x90, 0xa4, 0xae, 0xff")
	x.Var("BlueGrey400", "0x78, 0x90, 0x9c, 0xff")
	x.Var("BlueGrey500", "0x60, 0x7d, 0x8b, 0xff")
	x.Var("BlueGrey600", "0x54, 0x6e, 0x7a, 0xff")
	x.Var("BlueGrey700", "0x45, 0x5a, 0x64, 0xff")
	x.Var("BlueGrey800", "0x37, 0x47, 0x4f, 0xff")
	x.Var("BlueGrey900", "0x26, 0x32, 0x38, 0xff")

}