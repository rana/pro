package app

import (
	"fmt"
	"log"
	"path/filepath"
	"sys"
	"sys/err"
	"sys/fs"
	// tui "github.com/marcusolsson/tui-go"
)

var (
	_ui *UI
)

type (
	UI struct {
		*UIBse
		App         *App
		Main        *Box
		MainLeft    *VList
		MainRight   *ProBox
		Btm         *Box
		ScriptTaber *Taber
		Console     *ReadBox
		HasLeft     bool
		HasBtm      bool
	}
)

func NewUI(app *App) (r *UI) {
	var er error
	r = &UI{}
	_ui = r // global access; for Theme access
	r.App = app
	r.UIBse, er = NewUIBse() //r.Root)
	if er != nil {
		err.Panic(er)
	}
	r.HasLeft = true
	r.HasBtm = true

	r.Console = NewReadBox()
	r.Console.SetTitle("Console")
	r.Console.SetBorder(true)
	r.Console.SetSizePolicy(Expanding, Preferred)

	r.MainLeft = NewVList()
	r.MainLeft.SetTitle("Scripts")
	r.MainLeft.SetBorder(true)
	r.MainLeft.SetSizePolicy(Preferred, Maximum)
	r.MainLeft.OnItemActivated(r.FleOpn)

	r.MainRight = NewProBox(app)
	r.MainRight.SetSizePolicy(Expanding, Expanding)

	r.Main = NewHBox(r.MainLeft, r.MainRight)
	r.Main.Title = "Main"
	r.Main.SetSizePolicy(Preferred, Expanding)

	r.LoadScripts()
	r.ScriptTaber = NewTaber()

	// r.Console.SetText(`BEGIN
	// 	kjhsdf
	// 	SDFHJKSFD SDIJLKSA asdJKHADShj asdHJKASD AS ASDKJhasD ASDuhaSDYU
	// 	KSA asdJKHADS KJhasD ASDuh
	// 	abc
	// 	def
	// 	hij
	// 	klm
	// 	123
	// 	rty
	// 	lkjfg
	// 	kjhsdf SDFHJKSFD SDIJLKSA asdJKHADShj asdHJKASD AS ASDK
	// 	END
	// 	`)

	r.Btm = NewHBox(r.Console)
	r.Btm.Title = "Btm"
	r.Btm.SetSizePolicy(Expanding, Expanding)

	r.Root = NewVBox(r.Main, r.Btm)
	r.Root.Title = "Root"

	r.SetKeybinding("Ctrl+q", app.Cls)
	r.SetKeybinding("Ctrl+b", r.ToggleLeft)
	r.SetKeybinding("`", r.ToggleBtm)
	SetLogger(log.New(r, "", 0))
	// SetLogger(log.New(os.Stderr, "", log.LstdFlags))
	return r
}

func (x *UI) Init() {
}

func (x *UI) LoadScripts() {
	defer func() { // may crash in prd with no scripts dir
		if v := recover(); v != nil {
			switch t := v.(type) {
			case *err.Err:
				sys.Log(t.Full())
			default:
				sys.Log(err.New(v).Full())
			}
		}
	}()
	dir := filepath.Join(string(x.App.Cfg.Wd), "scripts")
	pths := fs.Fles(dir)
	for _, pth := range pths {
		x.MainLeft.Add(&Fle{
			Name: fs.Name(pth),
			Pth:  pth,
		})
	}
}
func (x *UI) ToggleLeft() {
	if x.HasLeft {
		x.Main.Remove(0)
	} else {
		x.Main.Insert(0, x.MainLeft)
	}
	x.HasLeft = !x.HasLeft
}
func (x *UI) ToggleBtm() {
	if x.HasBtm {
		x.Root.Remove(1)
	} else {
		x.Root.Insert(1, x.Btm)
	}
	x.HasBtm = !x.HasBtm
}
func (x *UI) FleOpn(l *List) {
	// // x.Logf("l.Selected() : %v", l.Selected())
	// txt := fs.LoadText(fle.Pth)
	// x.ScriptTaber.Add(fle.Name, NewReadBox(txt))

	fle := l.SelectedItem().(*Fle)
	fle.Load()
	x.MainRight.SetFle(fle)
	x.Main.FocusNext()
}
func (x *UI) Logf(format string, args ...interface{}) { x.Write([]byte(fmt.Sprintf(format, args...))) }
func (x *UI) Log(s string)                            { x.Write([]byte(s)) }
func (x *UI) Write(p []byte) (n int, err error) { // io.Writer for logging
	x.Console.SetText(string(p))
	return len(p), nil
}
