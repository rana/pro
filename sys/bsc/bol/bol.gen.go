package bol

import (
	"bytes"
	"strings"
	"sys/k"
)

const (
	Zero = Bol(false)
	Fls  = Bol(false)
	Tru  = Bol(true)
	Size = int(1)
)

type (
	Bol    bool
	BolScp struct {
		Idx uint32
		Arr []Bol
	}
)

func (x Bol) Not() Bol { return !x }
func (x Bol) StrWrt(b *strings.Builder) {
	if x {
		b.WriteString(k.Tru)
	} else {
		b.WriteString(k.Fls)
	}
}
func (x Bol) BytWrt(b *bytes.Buffer) {
	if x {
		b.WriteByte(1)
	} else {
		b.WriteByte(0)
	}
}
func (x *Bol) BytRed(b []byte) int {
	*x = Bol(b[0] == 1)
	return Size
}
func (x Bol) Eql(a Bol) Bol { return x == a }
func (x Bol) Neq(a Bol) Bol { return x != a }
func (x Bol) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x Bol) Bytes() []byte {
	b := &bytes.Buffer{}
	x.BytWrt(b)
	return b.Bytes()
}
