package scn

import (
	"fmt"
	"sys/bsc/unt"
	"sys/err"
	"unicode/utf8"
)

const (
	BOM   = rune(0xFEFF)
	EndCh = rune(-1)
)

type (
	Scnr struct {
		Scn
		Txt string
	}
)

func (x *Scnr) Reset(txt string) {
	x.Txt = txt
	x.Ch = ' '
	x.Size = 0
	x.Idx = 0
	x.Ln = 1
	x.Col = 0
	x.End = false
	x.NextRune()
	if x.Ch == BOM {
		x.NextRune()
	}
}
func (x *Scnr) Resetf(format string, args ...interface{}) { x.Reset(fmt.Sprintf(format, args...)) }
func (x *Scnr) NextRune() bool {
	x.Idx += unt.Unt(x.Size)
	x.End = x.Idx >= unt.Unt(len(x.Txt))
	if x.End {
		x.Ch = EndCh // necessary for char check beyond lim to fail
	} else {
		x.Ch, x.Size = utf8.DecodeRuneInString(x.Txt[x.Idx:])
		if x.Ch == '\n' {
			x.Ln++
			x.Col = 0
		} else {
			x.Col++
		}
	}
	return x.End
}
func (x *Scnr) PeekRune() (r rune) {
	scn := x.Scn
	x.NextRune()
	r = x.Ch
	x.Scn = scn
	return r
}
func (x *Scnr) SkpSet(opn, cls rune) {
	if x.Ch != opn { // must start with opn rune
		err.Panicf("Scnr: missing open rune '%v'", string(opn))
	}
	x.NextRune() // skp opn
	if x.Ch != cls {
		depth := 0
		for !x.End {
			x.NextRune()
			if x.Ch == opn {
				depth++
			} else if x.Ch == cls {
				if depth == 0 {
					break
				} else {
					depth--
				}
			}
		}
		if x.Ch != cls { // must end with cls rune
			err.Panicf("Scnr: missing close rune '%v'", string(cls))
		}
	}
	x.NextRune() // skp cls
}
func (x *Scnr) PrintScn() {
	fmt.Printf("Ch:'%v' %v Size:%v Idx:%v Ln:%v Col:%v End:%v \n", string(x.Ch), x.Ch, x.Size, x.Idx, x.Ln, x.Col, x.End)
}
