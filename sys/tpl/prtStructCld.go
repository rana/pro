package tpl

import (
	"strings"
	"sys/err"
	"sys/k"
	"sys/tpl/atr"
)

type (
	PrtStructCld struct {
		PrtBse
		Schema  *TypFn
		CldName *TypFn
		CldSav  *TypFn
		Save    *TypFn // bigquery ValueSaver interface
	}
)

func (x *PrtStructCld) InitPrtTypFn() {
	x.Schema = x.cldSchema()
	x.CldName = x.cldName()
	x.Save = x.save()
	x.CldSav = x.cldSav()
}
func (x *PrtStructCld) cldName() (r *TypFn) {
	r = x.f.TypFna("CldName", atr.None)
	r.OutPrm(String)
	r.Addf("return %q", x.t.Title())
	return r
}
func (x *PrtStructCld) cldSchema() (r *TypFn) {
	x.f.Import("context")
	x.f.Import("sys/err")
	x.f.Import("cloud.google.com/go/bigquery")
	rxr := x.f.Typ().(*Struct)
	r = x.f.TypFn(k.CldSchema)
	r.OutPrm(NewExt("bigquery.Schema"), "r")
	if len(rxr.Flds) == 0 {
		r.Add("panic(\"no flds\")")
	} else {
		var b strings.Builder
		var cnt int
		b.WriteString("// ")
		for _, fld := range rxr.Flds {
			if !fld.IsBqSkp() {
				r.Addf("r = append(r, &bigquery.FieldSchema{Name:%q, Required:true, Type:bigquery.%v})", fld.Name, BqFieldType(fld.Typ))
				if cnt != 0 {
					b.WriteString(", ")
				}
				b.WriteString(fld.Name)
				cnt++
			}
		}
		r.Add(b.String()) // for bq copy-paste sql
	}
	r.Add("return r")
	return r
}
func (x *PrtStructCld) save() (r *TypFn) { // bq ValueSaver interface
	rxr := x.f.Typ().(*Struct)
	r = x.f.TypFn("Save")
	r.OutPrm(NewExt("map[string]bigquery.Value"), "r")
	r.OutPrm(String, "insertID")
	r.OutPrm(Error, "er")
	if len(rxr.Flds) == 0 {
		r.Add("panic(\"no flds\")")
	} else {
		r.Add("r = make(map[string]bigquery.Value)")
		for _, fld := range rxr.Flds {
			if !fld.IsBqSkp() {
				r.Addf("r[%q] = x.%v", fld.Name, fld.Name)
			}
		}
	}
	r.Add("return r, insertID, er")
	return r
}
func (x *PrtStructCld) cldSav() (r *TypFn) {
	x.f.Import(_sys)
	r = x.f.TypFna(k.CldSav, atr.Lng)
	r.OutPrm(x.t)
	r.Add("tbl := sys.Cld().GetTable(x)")
	r.Add("u := tbl.Uploader()")
	r.Add("if er := u.Put(context.Background(), x); er != nil {")
	r.Add("err.Panicf(\"cld sav: %v\", er)")
	r.Add("}")
	r.Add("return x")
	return r
}
func BqFieldType(t Typ) string {
	switch {
	case t == _sys.Bsc.Str.Typ():
		return "StringFieldType"
	case t == _sys.Bsc.Bol.Typ():
		return "BooleanFieldType"
	case t == _sys.Bsc.Flt.Typ():
		return "FloatFieldType"
	case t == _sys.Bsc.Unt.Typ():
		return "IntegerFieldType"
	case t == _sys.Bsc.Int.Typ():
		return "IntegerFieldType"
	case t == _sys.Bsc.Tme.Typ():
		return "IntegerFieldType"
	}
	err.Panicf("Bq field type: '%v' unsupported", t.Full())
	return ""
}
