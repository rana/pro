package app

import (
	"image"
)

type (
	List struct {
		BorderBse
		Padding            Length
		items              []*ListItem
		selected           int
		onItemActivated    func(*List)
		onSelectionChanged func(*List)
		measure            func()
		sizeHint           image.Point
		minSizeHint        image.Point
		title              *PrtTitle
		border             *PrtBorder
	}
	ListItem struct {
		Itm   Itm
		Style Style
		X     int
		Width int
	}
)

func (x *List) SetTitle(text string) {
	x.BorderBse.SetTitle(text)
	x.ClearSizeHints()
}
func (x *List) SetBorder(has bool) {
	x.BorderBse.SetBorder(has)
	x.ClearSizeHints()
}
func (x *List) ClearSizeHints() {
	x.sizeHint = image.ZP
	x.minSizeHint = image.ZP
}
func (x *List) SizeHint() image.Point {
	if x.sizeHint.Eq(image.ZP) {
		x.measure()
	}
	return x.sizeHint
}
func (x *List) MinSizeHint() image.Point {
	if x.minSizeHint.Eq(image.ZP) {
		x.measure()
	}
	return x.minSizeHint
}
func (x *List) Prev() { x.Navigate(x.selected - 1) }
func (x *List) Next() { x.Navigate(x.selected + 1) }
func (x *List) Navigate(i int, isFst ...bool) {
	if len(x.items) != 0 {
		if i < 0 {
			i = -i
			i = (len(x.items) - i) % len(x.items)
		} else if i >= len(x.items) {
			i = i % len(x.items)
		}
		x.selected = i
		x.ApplyStyle(len(isFst) == 0) // assume any call to Navigate occurs when focused
		x.RaiseSelectionChanged()
	}
}
func (x *List) RaiseActivate() {
	if x.onItemActivated != nil {
		x.onItemActivated(x)
	}
}
func (x *List) RaiseSelectionChanged() {
	if x.onSelectionChanged != nil {
		x.onSelectionChanged(x)
	}
}

func (x *List) Add(vs ...Itm) {
	if len(vs) != 0 {
		isFst := len(x.items) == 0
		for _, v := range vs {
			x.items = append(x.items, &ListItem{
				Itm:   v,
				Style: _ui.Theme().Style(ListItemKey),
			})
		}
		x.ClearSizeHints()
		if isFst {
			x.Navigate(0, true)
		}
	}
}
func (x *List) RemoveAt(i int) {
	if len(x.items) != 0 {
		if i >= 0 && i < len(x.items) {
			x.items = append(x.items[i:], x.items[i+1:]...)
			if x.selected >= len(x.items) {
				x.selected = len(x.items) - 1
			}
			x.ClearSizeHints()
			x.RaiseSelectionChanged()
		}
	}
}
func (x *List) ClearItems() {
	x.items = nil
	x.selected = 0
	x.ClearSizeHints()
	x.RaiseSelectionChanged()
}

func (x *List) Selected() int {
	return x.selected
}
func (x *List) SelectedItem() Itm {
	if x.selected > -1 && x.selected < len(x.items) {
		return x.items[x.selected].Itm
	}
	return nil
}

// OnItemActivated gets called when activated (through pressing KeyEnter).
func (x *List) OnItemActivated(fn func(*List)) {
	x.onItemActivated = fn
}

// OnSelectionChanged gets called whenever a new item is selected.
func (x *List) OnSelectionChanged(fn func(*List)) {
	x.onSelectionChanged = fn
}
func (x *List) Focus(prev ...bool) bool {
	x.ApplyStyle(true)
	return true
}
func (x *List) Unfocus() { x.ApplyStyle(false) }
func (x *List) ApplyStyle(isFocused bool) {
	x.BorderBse.ApplyStyle(isFocused)
	for n, item := range x.items {
		if n == x.selected {
			if isFocused {
				item.Style = _ui.Theme().Style(ListItemSelectedFocusedKey)
			} else {
				item.Style = _ui.Theme().Style(ListItemSelectedKey)
			}
		} else {
			item.Style = _ui.Theme().Style(ListItemKey)
		}
	}
}
