package srch_test

import (
	"sys/ana/cfg"
	"sys/ana/hst/srch"
	"sys/app"
	"testing"
)

func TestSrch(t *testing.T) {
	cfg := cfg.Load("/home/rana/go/src/sys/cmd/sys.cfg")
	cfg.Ui = false
	cfg.Test = true // for acount balance set
	ap := app.New(cfg, "/home/rana/go/src/sys/cmd")
	defer ap.Cls()

	srch.SrchMl0()
}
