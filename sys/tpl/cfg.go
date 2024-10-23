package tpl

import "sys/k"

type (
	DirCfg struct {
		DirBse
		Cfgr *FleCfgr
	}
)

func (x *DirPro) NewCfg() (r *DirCfg) {
	r = &DirCfg{}
	x.Cfg = r
	r.Pkg = x.Pkg.New(k.Cfg)
	r.Cfgr = r.NewCfgr()
	return r
}
