package tpl

import (
	"fmt"
	"strings"
)

type (
	Lines []string
)

func (x *Lines) Add(lns ...string) { *x = append(*x, lns...) }
func (x *Lines) Addf(format string, args ...interface{}) {
	*x = append(*x, fmt.Sprintf(format, args...))
}
func (x *Lines) Ins(idx int, ln string) {
	*x = append((*x)[:idx], append([]string{ln}, (*x)[idx:]...)...)
}
func (x *Lines) Cnt() int        { return len(*x) }
func (x *Lines) Exist() bool     { return len(*x) != 0 }
func (x *Lines) RemFrom(idx int) { *x = (*x)[:idx] }
func (x *Lines) Cpy() Lines {
	r := make(Lines, len(*x))
	copy(r, *x)
	return r
}

func (x Lines) WriteLines(b *strings.Builder) {
	if len(x) > 0 {
		for _, line := range x {
			b.WriteString(line)
			b.WriteRune('\n')
		}
	}
}
