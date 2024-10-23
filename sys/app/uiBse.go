package app

import (
	"sys"
	"sys/err"

	"github.com/gdamore/tcell"
)

type UIBse struct {
	painter     *Painter
	Root        *Box
	screen      tcell.Screen
	keybindings []*keybinding
	eventQueue  chan event
	quit        chan struct{}
}

func NewUIBse() (*UIBse, error) {
	screen, err := tcell.NewScreen()
	if err != nil {
		return nil, err
	}
	s := &tcellSurface{
		screen: screen,
	}
	p := NewPainter(s, NewTheme())
	return &UIBse{
		painter:     p,
		keybindings: make([]*keybinding, 0),
		quit:        make(chan struct{}, 1),
		screen:      screen,
		eventQueue:  make(chan event),
	}, nil
}

func (x *UIBse) Theme() *Theme     { return x.painter.theme }
func (x *UIBse) SetTheme(t *Theme) { x.painter.theme = t }

func (x *UIBse) SetKeybinding(seq string, fn func()) {
	x.keybindings = append(x.keybindings, &keybinding{
		sequence: seq,
		handler:  fn,
	})
}

// ClearKeybindings reinitialises ui.keybindings so as to revert to a
// clear/original state
func (x *UIBse) ClearKeybindings() {
	x.keybindings = make([]*keybinding, 0)
}

func (x *UIBse) Run() error {
	if err := x.screen.Init(); err != nil {
		return err
	}
	x.Root.FocusNext()
	x.screen.SetStyle(tcell.StyleDefault)
	x.screen.EnableMouse()
	x.screen.Clear()

	go x.loop()

	for {
		select {
		case <-x.quit:
			return nil
		case ev := <-x.eventQueue:
			x.handleEvent(ev)
		}
	}
}
func (x *UIBse) loop() {
	defer func() {
		if v := recover(); v != nil {
			switch t := v.(type) {
			case *err.Err:
				sys.Log(t.Full())
			default:
				sys.Log(err.New(v).Full())
			}
		}
	}()
	for {
		switch ev := x.screen.PollEvent().(type) {
		case *tcell.EventKey:
			x.handleKeyEvent(ev)
		case *tcell.EventMouse:
			x.handleMouseEvent(ev)
		case *tcell.EventResize:
			x.handleResizeEvent(ev)
		}
	}
}
func (x *UIBse) handleEvent(ev event) {
	switch e := ev.(type) {
	case KeyEvent:
		// logger.Printf("Received key event: %s", e.Name())
		for _, b := range x.keybindings {
			if b.match(e) {
				b.handler()
			}
		}
		switch e.Key {
		case KeyTab:
			// logger.Printf("UIBse.handleEvent: KeyTab")
			x.Root.FocusNext()
		case KeyBacktab:
			// logger.Printf("UIBse.handleEvent: KeyBacktab")
			x.Root.FocusPrev()
		default:
			x.Root.OnKeyEvent(e)
		}
		x.painter.Repaint(x.Root)
	case callbackEvent:
		// Gets stuck in a print loop when the logger is a widget.
		//logger.Printf("Received callback event")
		e.cbFn()
		x.painter.Repaint(x.Root)
	case paintEvent:
		// logger.Printf("Received paint event")
		x.painter.Repaint(x.Root)
	}
}

func (x *UIBse) handleKeyEvent(tev *tcell.EventKey) {
	x.eventQueue <- KeyEvent{
		Key:       Key(tev.Key()),
		Rune:      tev.Rune(),
		Modifiers: ModMask(tev.Modifiers()),
	}
}

func (x *UIBse) handleMouseEvent(ev *tcell.EventMouse) {
	e := &MouseEvent{}
	e.X, e.Y = ev.Position()
	x.eventQueue <- e
}

func (x *UIBse) handleResizeEvent(ev *tcell.EventResize) {
	x.eventQueue <- paintEvent{}
}

// Quit signals to the UIBse to start shutting down.
func (x *UIBse) Quit() {
	logger.Printf("Quitting")
	x.screen.Fini()
	x.quit <- struct{}{}
}

// Schedule an update of the UIBse, running the given
// function in the UIBse goroutine.
//
// Use this to update the UIBse in response to external events,
// like a timer tick.
// This method should be used any time you call methods
// to change UIBse objects after the first call to `UIBse.Run()`.
//
// Changes invoked outside of either this callback or the
// other event handler callbacks may appear to work, but
// is likely a race condition.  (Run your program with
// `go run -race` or `go install -race` to detect this!)
//
// Calling Update from within an event handler, or from within an Update call,
// is an error, and will deadlock.
func (x *UIBse) Update(fn func()) {
	blk := make(chan struct{})
	x.eventQueue <- callbackEvent{func() {
		fn()
		close(blk)
	}}
	<-blk
}

func convertColor(col Color, fg bool) tcell.Color {
	switch col {
	case ColorDefault:
		if fg {
			return tcell.ColorWhite
		}
		return tcell.ColorDefault
	case ColorBlack:
		return tcell.ColorBlack
	case ColorWhite:
		return tcell.ColorWhite
	case ColorRed:
		return tcell.ColorRed
	case ColorGreen:
		return tcell.ColorGreen
	case ColorBlue:
		return tcell.ColorBlue
	case ColorCyan:
		return tcell.ColorDarkCyan
	case ColorMagenta:
		return tcell.ColorDarkMagenta
	case ColorYellow:
		return tcell.ColorYellow
	case ColorNavy:
		return tcell.ColorNavy
	case ColorGray:
		return tcell.ColorGray
	default:
		if col > 0 {
			return tcell.Color(col)
		}
		return tcell.ColorDefault
	}
}
