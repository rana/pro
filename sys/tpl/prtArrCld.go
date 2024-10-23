package tpl

import (
	"sys/k"
	"sys/tpl/atr"
)

type (
	PrtArrCld struct {
		PrtBse
	}
)

func (x *PrtArrCld) InitPrtTypFn() {
	x.cldSav()
	x.cldQry()
	x.cldQryf()
}
func (x *PrtArrCld) cldSav() (r *TypFn) {
	x.f.Import("context")
	x.f.Import(_sys)
	x.f.Import("sys/err")
	// x.f.Import("sys/cld")
	r = x.f.TypFna(k.CldSav, atr.Lng)
	r.OutPrm(x.t)
	r.Add("if len(*x) != 0 {")
	r.Add("sys.Cld().Mu().Lock()")
	r.Add("defer sys.Cld().Mu().Unlock()")
	// r.Add("cldr := cld.New(sys.Cld().Cfg()) // new instance to avoid locking")
	r.Add("tbl := sys.Cld().GetTable((*x)[0])")
	r.Add("u := tbl.Uploader()")
	r.Add("if er := u.Put(context.Background(), *x); er != nil {")
	r.Add("err.Panicf(\"cld sav: %v\", er)")
	r.Add("}")
	r.Add("}")
	r.Add("return x")
	return r
}
func (x *PrtArrCld) cldQry() (r *TypFn) {
	x.f.Import("google.golang.org/api/iterator")
	r = x.f.TypFna(k.CldQry, atr.Lng)
	r.InPrm(String, "txt")
	r.OutPrm(x.t)
	r.Add("sys.Cld().Mu().Lock()")
	r.Add("defer sys.Cld().Mu().Unlock()")
	r.Add("itr := sys.Cld().Query(txt)")
	r.Add("for {")
	r.Add("var v StgyPrfm")
	r.Add("er := itr.Next(&v)")
	r.Add("if er == iterator.Done {")
	r.Add("break")
	r.Add("}")
	r.Add("if er != nil {")
	r.Add("err.Panicf(\"cld query: %v\", er)")
	r.Add("}")
	r.Add("x.Push(&v)")
	r.Add("}")
	r.Add("return x")
	return r
}
func (x *PrtArrCld) cldQryf() (r *TypFn) {
	x.f.Import("fmt")
	r = x.f.TypFna(k.CldQry+"f", atr.Lng)
	r.InPrm(String, "format")
	r.InPrmVariadic(Interface, "args")
	r.OutPrm(x.t)
	r.Add("x.CldQry(fmt.Sprintf(format, args...))")
	r.Add("return x")
	return r
}
