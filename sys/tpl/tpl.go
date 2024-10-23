package tpl

import (
	"flag"
)

var (
	Wrt   = flag.Bool("wrt", true, "writes files")
	Clean = flag.Bool("clean", false, "delete previously generated files and empty folders")
	Long  = flag.Bool("long", false, "generate longer test cases")
	Trc   = *flag.Int("trc", 0, "generate trace statements")

	//Opt = Test | Arr | Hst | Rlt | Oan | Instr | Inrvl | Side | Stm | Cnd | Stgy | Port | Vis
	Opt = TestOpt | ArrOpt | HstOpt | RltOpt | CndOpt | VisOpt
)

func init() {
	flag.Parse() // flag.Parse() Must be called after all flags are defined and before flags are accessed by the program.
}

const (
	TestOpt TplOpt = 1 << iota
	TestLongOpt
	BscOpt
	HstOpt
	RltOpt
	VisOpt
	OanOpt
	InstrOpt
	InrvlOpt
	SideOpt
	StmOpt
	StmsOpt
	CndOpt
	CndsOpt
	StgyOpt
	PortOpt
	ArrOpt
	XprOpt
	ActOpt
	NoOpt TplOpt = 0
)

type (
	TplOpt uint32
)

func (x TplOpt) Is(vs ...TplOpt) (r bool) {
	for _, v := range vs {
		if x&v != v {
			return false
		}
	}
	return true
}

func (x TplOpt) IsTest() bool     { return x&TestOpt == TestOpt }
func (x TplOpt) IsTestLong() bool { return x&TestLongOpt == TestLongOpt }
func (x TplOpt) IsHst() bool      { return x&HstOpt == HstOpt }
func (x TplOpt) IsRlt() bool      { return x&RltOpt == RltOpt }
func (x TplOpt) IsVis() bool      { return x&VisOpt == VisOpt }

func (x TplOpt) IsOan() bool   { return x&OanOpt == OanOpt }
func (x TplOpt) IsInstr() bool { return x&InstrOpt == InstrOpt }
func (x TplOpt) IsInrvl() bool { return x&InrvlOpt == InrvlOpt }
func (x TplOpt) IsSide() bool  { return x&SideOpt == SideOpt }

func (x TplOpt) IsStm() bool  { return x&StmOpt == StmOpt }
func (x TplOpt) IsCnd() bool  { return x&CndOpt == CndOpt }
func (x TplOpt) IsStgy() bool { return x&StgyOpt == StgyOpt }
func (x TplOpt) IsPort() bool { return x&PortOpt == PortOpt }
func (x TplOpt) IsArr() bool  { return x&ArrOpt == ArrOpt }

func (x TplOpt) IsXpr() bool { return x&XprOpt == XprOpt }
func (x TplOpt) IsAct() bool { return x&ActOpt == ActOpt }
