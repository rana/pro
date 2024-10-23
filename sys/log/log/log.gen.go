package log

import (
	"sys"
	"sys/bsc/str"
)

type (
	Logr string
)

func Ifo(vs ...interface{}) (r str.Str) {
	sys.Log(vs...)
	return r
}
func Ifof(tmpl str.Str, vs ...interface{}) (r str.Str) {
	sys.Logf(string(tmpl), vs...)
	return r
}
