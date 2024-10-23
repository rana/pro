package tpl

import (
	"fmt"
	"strings"
	"sys"
	"unicode"
	"unicode/utf8"
)

type (
	// Lbl ...
	Lbl struct {
		Name string
	}
	// Lblr ...
	Lblr interface {
		Title() string
		Camel() string
		Lower() string
		Upper() string
		Plural() string
		Singular() string
		CamelSingular() string
		CamelPlural() string
		IsFstUpr() bool
	}
)

func (x *Lbl) Title() string         { return strings.Title(x.Name) }
func (x *Lbl) Camel() string         { return sys.Camel(x.Name) }
func (x *Lbl) Lower() string         { return strings.ToLower(x.Name) }
func (x *Lbl) Upper() string         { return strings.ToUpper(x.Name) }
func (x *Lbl) Plural() string        { return fmt.Sprintf("%vs", x.Name) }
func (x *Lbl) Singular() string      { return sys.Singular(x.Name) }
func (x *Lbl) CamelSingular() string { return sys.Camel(sys.Singular(x.Name)) }
func (x *Lbl) CamelPlural() string   { return sys.Camel(sys.Plural(x.Name)) }
func (x *Lbl) IsFstUpr() bool {
	if x.Name == "" {
		return false
	}
	ch, _ := utf8.DecodeRuneInString(x.Name)
	return unicode.IsUpper(ch)
}
