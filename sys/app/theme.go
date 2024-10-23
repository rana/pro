package app

// Color represents a color.
type Color int32

// Common colors.
const (
	ColorDefault Color = iota
	ColorBlack
	ColorWhite
	ColorRed
	ColorGreen
	ColorBlue
	ColorCyan
	ColorMagenta
	ColorYellow
	ColorNavy
	ColorGray
)

// Decoration represents a bold/underline/etc. state
type Decoration int

// Decoration modes: Inherit from parent widget, explicitly on, or explicitly off.
const (
	DecorationInherit Decoration = iota
	DecorationOn
	DecorationOff
)

// Style determines how a cell should be painted.
// The zero value uses default from
type Style struct {
	Fg Color
	Bg Color

	Reverse   Decoration
	Bold      Decoration
	Underline Decoration
}

// mergeIn returns the receiver Style, with any changes in delta applied.
func (s Style) Merge(delta Style) Style {
	result := s
	if delta.Fg != ColorDefault {
		result.Fg = delta.Fg
	}
	if delta.Bg != ColorDefault {
		result.Bg = delta.Bg
	}
	if delta.Reverse != DecorationInherit {
		result.Reverse = delta.Reverse
	}
	if delta.Bold != DecorationInherit {
		result.Bold = delta.Bold
	}
	if delta.Underline != DecorationInherit {
		result.Underline = delta.Underline
	}
	return result
}

// Theme defines the styles for a set of identifiers.
type Theme struct {
	styles map[string]Style
}

func (p *Theme) SetStyle(key string, i Style) {
	p.styles[key] = i
}

// Style returns the style associated with an identifier.
// If there is no Style associated with the name, it returns a default Style.
func (p *Theme) Style(key string) Style {
	return p.styles[key]
}

func (p *Theme) HasStyle(key string) bool {
	_, ok := p.styles[key]
	return ok
}

func NewTheme() (r *Theme) {
	r = &Theme{}
	r.styles = make(map[string]Style)
	r.styles["box.border.focused"] = Style{Fg: ColorGreen} // TODO: REMOVE
	r.styles[PrtBorderKey] = Style{Fg: ColorGray}
	r.styles[PrtTitleKey] = Style{Fg: ColorGray}
	r.styles[PrtBorderFocusedKey] = Style{Fg: ColorGreen}
	r.styles[ListItemSelectedKey] = Style{Bg: ColorGray}
	r.styles[ListItemSelectedFocusedKey] = Style{Bg: ColorGreen}
	r.styles["table.cell.selected"] = Style{Reverse: DecorationOn}
	r.styles["button.focused"] = Style{Reverse: DecorationOn}
	r.styles["label.tab"] = Style{Reverse: DecorationOff}
	r.styles["label.tab.selected"] = Style{Reverse: DecorationOn, Fg: ColorGreen, Bg: ColorWhite}
	return r
}

const (
	PrtBorderKey               = "prt.border"
	PrtBorderFocusedKey        = "prt.border.focused"
	PrtTitleKey                = "prt.title"
	PrtTitleFocusedKey         = "prt.title.focused"
	TextKey                    = "text"
	TextFocusedKey             = "text.focused"
	ListItemKey                = "list.item"
	ListItemSelectedKey        = "list.item.selected"
	ListItemSelectedFocusedKey = "list.item.selected.focused"
)
