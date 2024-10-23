package app

// SizePolicy determines the space occupied by a widget.
type SizePolicy int

const (
	// Preferred interprets the size hint as the preferred size.
	Preferred SizePolicy = iota
	// Minimum allows the widget to shrink down to the size hint.
	Minimum
	// Maximum allows the widget to grow up to the size hint.
	Maximum
	// Expanding makes the widget expand to the available space.
	Expanding
)
