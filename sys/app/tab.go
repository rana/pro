package app

type (
	Taber struct {
		*Box
		topbar   *HList
		tabs     []*Tab
		selected int
	}
	Tab struct {
		Name string
		View Ctrl
	}
)

func NewTaber() (r *Taber) {
	r = &Taber{}
	// r.MayFocus = true
	r.topbar = NewHList()
	r.topbar.SetSizePolicy(Minimum, Maximum)
	r.topbar.OnSelectionChanged(r.TabChanged)
	r.Box = NewVBox(r.topbar)
	r.Box.SetSizePolicy(Maximum, Preferred)
	return r
}
func (x *Taber) Add(name string, w Ctrl) {
	for n, tab := range x.tabs {
		if tab.Name == name {
			x.Navigate(n) // navigate to duplicate
			return
		}
	}
	// TODO:
	// w.Bse().Prnt = x
	x.tabs = append(x.tabs, &Tab{
		Name: name,
		View: w,
	})
	x.topbar.Add(x.tabs[len(x.tabs)-1])
	x.Next()
}
func (x *Taber) RemoveAt(i int) {
	if len(x.tabs) != 0 {
		if i < 0 {
			i = 0
		} else if i >= len(x.tabs) {
			i = len(x.tabs) - 1
		}

	}
}
func (x *Taber) Prev() { x.Navigate(x.selected - 1) }
func (x *Taber) Next() { x.Navigate(x.selected + 1) }
func (x *Taber) Navigate(i int) {
	if len(x.tabs) != 0 {
		if i < 0 {
			i = 0
		} else if i >= len(x.tabs) {
			i = len(x.tabs) - 1
		}
		x.selected = i
		x.Box.Remove(1)
		x.Box.Add(x.tabs[x.selected].View)
	}
}
func (x *Taber) TabChanged(l *List) {
	x.Navigate(l.Selected())
}
func (x *Taber) Selected() int { return x.selected }
func (x *Taber) SelectedTab() *Tab {
	if len(x.tabs) == 0 {
		return nil
	}
	return x.tabs[x.selected]
}
func (x *Taber) OnKeyEvent(ev KeyEvent) {
	switch ev.Key {
	case KeyBacktab, KeyLeft:
		x.Prev()
	case KeyTab, KeyRight:
		x.Next()
	}
	// x.Box.OnKeyEvent(ev)
}

func (x *Tab) GetName() string { return x.Name } // for Itm interface
