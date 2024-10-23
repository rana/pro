package tpl

import (
	"fmt"
	"strings"
)

type (
	Block struct {
		Lines
	}
)

func (x *Block) WriteBlock(b *strings.Builder, skpLstNewln ...bool) {
	if len(x.Lines) == 0 {
		b.WriteString(" {}")
	} else if len(x.Lines) == 1 {
		b.WriteString(fmt.Sprintf(" { %v }", x.Lines[0]))
	} else {
		b.WriteString(" {\n")
		x.WriteLines(b)
		b.WriteString("}")
	}
	if len(skpLstNewln) == 0 {
		b.WriteRune('\n')
	}
}
