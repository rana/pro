package tpl

import (
	"sys/k"
)

type (
	DirSys struct {
		DirBse
		Fles
		Ext  *ExtFle
		Ifc  *FleSysIfc
		Trc  *DirTrc
		Log  *DirLog
		Tst  *FleTst
		Test *DirTest
		Bsc  *DirBsc
		Ana  *DirAna
		Lng  *DirLng
		Idn  *FleSysIdn
		Act  *FleSysAct
		Mu   *FleSysMu
	}
	DirBse struct {
		Pkg *Pkg
	}
	Dir interface {
		DirPkg() *Pkg
	}
)

var (
	_sys *DirSys
)

func NewSys() *DirSys {
	_sys = &DirSys{}
	_sys.Pkg = NewPkg(k.Sys)
	_sys.NewExt()
	_sys.NewIfc()
	_sys.NewOpt()
	_sys.NewLog()
	_sys.NewTst()
	_sys.NewTest()
	_sys.NewBsc()
	_sys.NewAna()
	_sys.NewLng()
	_sys.NewIdn()
	_sys.NewAct()
	_sys.NewMu()
	return _sys
}

func (x *DirBse) DirPkg() *Pkg { return x.Pkg }
