package tpl

import (
	"fmt"
	"strings"
)

type (
	// Cmnts is a comments template.
	Cmnts []string
)

// NewCmnts creates a comments template.
func NewCmnts(vs ...string) Cmnts { return vs }

// NewCmntf creates a comments template.
func NewCmntf(format string, args ...interface{}) Cmnts { return []string{fmt.Sprintf(format, args...)} }

// WriteCmnts writes comments to the buffer.
func (x Cmnts) WriteCmnts(b *strings.Builder) {
	if len(x) > 0 {
		for _, cmnt := range x {
			b.WriteString("// ")
			b.WriteString(cmnt)
			b.WriteString("\n")
		}
	}
}

// Add adds a comment to the comments template.
func (x *Cmnts) Add(cmnts ...string) {
	for _, cmnt := range cmnts {
		*x = append(*x, cmnt)
	}
}

// Addf adds a formatted comment to the comments template.
func (x *Cmnts) Addf(format string, args ...interface{}) {
	*x = append(*x, fmt.Sprintf(format, args...))
}
