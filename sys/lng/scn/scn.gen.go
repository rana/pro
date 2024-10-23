package scn

import (
	"fmt"
	"sys/bsc/unt"
)

type (
	Scn struct {
		Ch   rune
		Size int
		Idx  unt.Unt
		Ln   unt.Unt
		Col  unt.Unt
		End  bool
	}
)

func (x *Scn) String() string {
	return fmt.Sprintf("Ch:%v:'%v' Size:%v Idx:%v Ln:%v Col:%v End:%v \n", x.Ch, string(x.Ch), x.Size, x.Idx, x.Ln, x.Col, x.End)
}
