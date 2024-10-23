package main

import (
	"sys"
	"sys/ana/cfg"
	"sys/ana/hst/srch"
	"sys/app"
	"sys/err"
)

func main() {
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
	// ap := app.New(nil)
	// ap.Run()

	cfg := cfg.Load("/home/rana/go/src/sys/cmd/sys.cfg")
	cfg.Ui = false
	cfg.Test = true // for acount balance set

	app.New(cfg)
	srch.SrchMl0()
}
