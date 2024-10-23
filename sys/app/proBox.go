package app

import (
	"fmt"
	"sys"
	"sys/bsc/tme"
	"sys/err"
	// "sys/lng/pro/act"
	// "sys/lng/pro/xpr"
	"sys/run"
	"time"
)

type (
	ProBox struct {
		Box
		app         *App
		script      *ReadBox
		bar         *Box
		barMsg      *Text
		barEllapsed *Text
		barState    *Text
		fle         *Fle
	}
)

func NewProBox(app *App) (r *ProBox) {
	r = &ProBox{}
	r.app = app
	r.script = NewReadBox()

	r.bar = NewHBox()
	r.barMsg = NewText()
	r.barMsg.Padding.Set(1)
	r.barEllapsed = NewText()
	r.barEllapsed.Padding.Left = 1
	r.barEllapsed.Style = Style{Fg: ColorGray}
	r.barState = NewText()
	r.barState.Padding.Set(1)
	r.barState.Style = Style{Fg: ColorRed}
	r.bar.Add(r.barMsg, NewSpacer(), r.barEllapsed, r.barState)

	r.Box.SetBorder(true)
	r.Box.Alignment = Vertical
	r.Box.Add(r.script)
	r.Box.Add(r.bar)

	return r
}
func (x *ProBox) SetFle(fle *Fle) {
	x.fle = fle
	if fle == nil {
		x.Box.SetTitle("")
		x.script.SetText("")
	} else {
		x.Box.SetTitle(x.fle.Name)
		x.script.Clear()
		x.script.SetText(x.fle.Txt)
		x.ParseFle()
	}
}
func (x *ProBox) SetSizePolicy(h, v SizePolicy) {
	x.Box.SetSizePolicy(h, v)
	x.script.SetSizePolicy(h, v)
	x.bar.SetSizePolicyX(h)
}
func (x *ProBox) OnKeyEvent(e KeyEvent) {
	switch e.Key {
	case KeyCtrlR:
		x.ReloadFle()
	case KeyCtrlSpace:
		x.ReloadFle()
		x.RunFle()
	}
}
func (x *ProBox) ReloadFle() {
	if x.fle != nil {
		x.fle.Load()
		x.SetFle(x.fle)
		x.ParseFle(true)
	}
}
func (x *ProBox) ParseFle(reloaded ...bool) { // running on UI thread
	if x.fle != nil {
		x.barState.SetTxt("Parsing")
		x.barState.Style = Style{Fg: ColorWhite}
		start := time.Now()
		go func() { // run on background thread; don't block UI
			defer func() {
				var er error
				if v := recover(); v != nil {
					sys.Log(err.New(v).Full())
					switch t := v.(type) {
					case error:
						er = t
					default:
						er = err.New(v)
					}
				}
				_ui.Update(func() { x.ParseCompleted(start, er, len(reloaded) != 0) })
			}()
			// var xprr xpr.Xprr
			// xprr.Prs(x.fle.Txt)
		}()
	}
}
func (x *ProBox) ParseCompleted(start time.Time, er error, reloaded bool) { // run on UI thread
	var prefix string
	if reloaded {
		prefix = "Reloaded & "
	}
	if er == nil {
		x.barState.SetTxt(prefix + "Parsed")
		x.barState.Style = Style{Fg: ColorGreen}
	} else {
		switch t := er.(type) {
		case *err.XprErr:
			x.barState.SetTxtf("%vParse error: Ln:%v Col:%v Ch:%v", prefix, t.Ln, t.Col, string(t.Ch))
			x.app.UI.Console.SetTextf("%v\n", t.Err.Full())
		case *err.Err:
			x.barState.SetTxtf("%vParsed: Error", prefix)
			x.app.UI.Console.SetTextf("%v\n", t.Full())
		case error:
			x.barState.SetTxt(t.Error())
			x.app.UI.Console.SetTextf("%v\n", er)
		}
		x.barState.Style = Style{Fg: ColorRed}
	}
	x.SetEllapsedTxt(time.Now().Sub(start))
}

func (x *ProBox) RunFle() { // running on UI thread
	if x.fle != nil {
		x.barState.SetTxt("Running...")
		x.barState.Style = Style{Fg: ColorWhite}
		start := time.Now()
		tkr := run.NewTkr(time.Second, x.UpdateEllapsed)
		go func() { // run on background thread; don't block UI
			defer func() {
				var er error
				if v := recover(); v != nil {
					e := err.New(v)
					er = e
					sys.Log(e.Full())
				}
				tkr.Stop()
				_ui.Update(func() { x.RunCompleted(start, er) })
			}()
			// actr := act.Actr{}
			// actr.Run(x.fle.Txt)
		}()
	}
}
func (x *ProBox) RunCompleted(start time.Time, er error) { // run on UI thread
	if er == nil {
		x.barState.SetTxt("Run: Completed")
		x.barState.Style = Style{Fg: ColorGreen}
	} else {
		x.barState.SetTxt(fmt.Sprintf("Run: Error"))
		x.barState.Style = Style{Fg: ColorRed}
		switch t := er.(type) {
		case *err.XprErr:
			x.barState.SetTxtf("Run error: Ln:%v Col:%v Ch:%v", t.Ln, t.Col, string(t.Ch))
			x.app.UI.Console.SetTextf("%v\n", t.Err.Full())
		case *err.Err:
			x.barState.SetTxtf("Run: Error")
			x.app.UI.Console.SetTextf("%v\n", t.Full())
		case error:
			x.barState.SetTxt(t.Error())
			x.app.UI.Console.SetTextf("%v\n", er)
		}
	}
	x.SetEllapsedTxt(time.Now().Sub(start))
}

func (x *ProBox) UpdateEllapsed(ellapsed time.Duration) { // run on UI thread
	// Log.Println("ProBox.UpdateEllapsed", ellapsed)
	_ui.Update(func() { x.SetEllapsedTxt(ellapsed) })
}
func (x *ProBox) SetEllapsedTxt(ellapsed time.Duration) {
	x.barEllapsed.SetTxt(tme.Tme(ellapsed.Seconds()).DurString())
}

func (x *ProBox) String() string {
	if x.fle != nil {
		return fmt.Sprintf("%v:%p", x.fle.Name, x)
	}
	return fmt.Sprintf("ProBox:%p", x)
}
