package app

import (
	"path/filepath"
	"sys"
	"sys/ana"
	"sys/ana/cfg"
	"sys/ana/ml"
	"sys/cld"
	"sys/dsk"
	"sys/fs"
	"sys/lng/pro/act"
	"sys/run"
)

type (
	App struct {
		Cfg  *cfg.Cfg
		Runr *run.Runr
		Cldr *cld.Cldr
		Dskr *dsk.Dskr
		Actr *act.Actr
		Ticr *ana.Ticr
		Oan  *ana.Oan
		UI   *UI
		Lrnr *ml.Lrnr
	}
)

func New(c *cfg.Cfg, dir ...string) (r *App) {
	r = &App{}
	if len(dir) == 0 {
		dir = append(dir, fs.WorkingDir())
	}
	if c == nil {
		c = cfg.Load(filepath.Join(dir[0], cfg.Filename))
	}
	r.Cfg = c
	ana.Cfg = c
	c.Wd = dir[0]
	r.Runr = run.NewRunr()
	r.Cldr = cld.New(c.BqProject.Unquo(), c.BqDataset.Unquo())
	if c.DskPth != "" { // allow some tests not to use Dsk
		r.Dskr = dsk.New(c.DskPth)
	}
	r.Actr = &act.Actr{}
	r.Lrnr = ml.NewLrnr(filepath.Join(c.Wd, ml.Filename))
	sys.NewSys(r.Runr, r.Cldr, r.Dskr, r.Actr, r.Lrnr)
	r.Ticr = ana.NewTicr(c)
	r.Oan = ana.NewOan(c, r.Ticr)
	if c.Ui {
		r.UI = NewUI(r)
	}
	return r
}
func (x *App) Cls() {
	if x.Oan != nil {
		x.Oan.Cls()
		x.Oan = nil
	}
	if x.Ticr != nil {
		x.Ticr.Cls()
		x.Ticr = nil
	}
	if x.UI != nil {
		x.UI.Quit()
	}
	if x.Dskr != nil {
		x.Dskr.Cls()
		x.Dskr = nil
	}
	if x.Runr != nil {
		x.Runr.Cls()
		x.Runr = nil
	}
	if x.Lrnr != nil {
		x.Lrnr.Cls()
		x.Lrnr = nil
	}
	// x.Actr = nil
}

// func (x *App) Run() {
// 	if x.Cfg.Rlt.OpnScript != "" {
// 		x.Oan.OpnStart()
// 		x.Actr.RunRlt(x.Cfg.Rlt.OpnScript.Unquo())
// 		x.Oan.OpnEnd()
// 	}
// 	if x.UI != nil {
// 		if er := x.UI.Run(); er != nil { // blocks until quit
// 			sys.Log(er)
// 			err.Panic(er)
// 		}
// 	}
// }
