package str

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strings"
	"sys/bsc/bol"
)

const (
	Zero  = Str("")
	Empty = Str("")
)

type (
	Str    string
	Cmp    func(a, b Str) bol.Bol
	StrScp struct {
		Idx uint32
		Arr []Str
	}
)

func Eql(a, b Str) bol.Bol                { return a == b }
func Lss(a, b Str) bol.Bol                { return a < b }
func Gtr(a, b Str) bol.Bol                { return a > b }
func Fmt(tmpl Str, vs ...interface{}) Str { return Str(fmt.Sprintf(string(tmpl), vs...)) }
func (x Str) Lower() Str                  { return Str(strings.ToLower(string(x))) }
func (x Str) Upper() Str                  { return Str(strings.ToUpper(string(x))) }
func (x Str) Unquo() string               { return string(x) }
func (x Str) StrWrt(b *strings.Builder) {
	b.WriteRune('"')
	b.WriteString(string(x))
	b.WriteRune('"')
}
func (x Str) BytWrt(b *bytes.Buffer) {
	bLen := make([]byte, 4)
	binary.LittleEndian.PutUint32(bLen, uint32(len(x)))
	b.Write(bLen)      // string length
	b.Write([]byte(x)) // string content
}
func (x *Str) BytRed(b []byte) int {
	sLen := binary.LittleEndian.Uint32(b[:4])
	if sLen > 0 {
		*x = Str(string(b[4 : 4+sLen]))
	}
	return 4 + int(sLen)
}
func (x Str) Eql(a Str) bol.Bol { return x == a }
func (x Str) Neq(a Str) bol.Bol { return x != a }
func (x Str) Lss(a Str) bol.Bol { return x < a }
func (x Str) Gtr(a Str) bol.Bol { return x > a }
func (x Str) Leq(a Str) bol.Bol { return x <= a }
func (x Str) Geq(a Str) bol.Bol { return x >= a }
func (x Str) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x Str) Bytes() []byte {
	b := &bytes.Buffer{}
	x.BytWrt(b)
	return b.Bytes()
}
