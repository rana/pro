package tpl

import (
	"strings"
	"sys/tpl/atr"
	"sys/tpl/mod"
)

type (
	Fld struct {
		Lbl
		mod.Mod
		atr.Atr           // for BytSkp
		Struct    *Struct // struct which contains the fld
		Typ       Typ
		Trm       *TypFn
		Lex       *TypFn
		GetXpr    *Struct
		SetGetXpr *Struct
		GetAct    *Struct
		SetGetAct *Struct
	}
	Flds   []*Fld
	FldMap map[string]*Fld
)

func NewFlds(vs ...*Fld) *Flds {
	r := Flds(vs)
	return &r
}
func (x *Flds) Ok() bool          { return len(*x) != 0 }
func (x *Flds) Cnt() int          { return len(*x) }
func (x *Flds) AddFld(vs ...*Fld) { *x = append(*x, vs...) }
func (x *Flds) Ins(idx int, vs ...*Fld) *Flds {
	*x = append((*x)[:idx], append(vs, (*x)[idx:]...)...)
	return x
}
func (x *Flds) Push(vs ...*Fld) *Flds {
	*x = append(*x, vs...)
	return x
}
func (x *Flds) Pop() (r *Fld) {
	if len(*x) == 0 {
		return r
	}
	r = (*x)[len(*x)-1]
	*x = (*x)[:len(*x)-1]
	return r
}
func (x *Flds) Dque() (r *Fld) {
	r = (*x)[0]
	if len(*x) == 1 {
		*x = (*x)[:0]
	} else {
		copy(*x, (*x)[1:])
		*x = (*x)[:len(*x)-1]
	}
	return r
}
func (x *Flds) OneFld(prd func(*Fld) bool) *Fld {
	var flds Flds
	flds = append(flds, *x...)
	for len(flds) != 0 {
		cur := flds.Dque()
		if prd(cur) {
			return cur
		}
		if cur.Name == "" {
			fStrct, ok := cur.Typ.(*Struct)
			if ok {
				flds.Ins(0, fStrct.Flds...)
			}
		}
	}
	return nil
}
func (x *Flds) OneFldInv(prd func(*Fld) bool) *Fld {
	var flds Flds
	flds = append(flds, *x...)
	for len(flds) != 0 {
		cur := flds.Dque()
		if prd(cur) {
			return cur
		}
		if cur.Name == "" {
			fStrct, ok := cur.Typ.(*Struct)
			if ok {
				flds.Push(fStrct.Flds...)
			}
		}
	}
	return nil
}
func (x *Flds) SelFlds(prd func(*Fld) bool) *Flds {
	var flds, r Flds
	flds = append(flds, *x...)
	for len(flds) != 0 {
		cur := flds.Dque()
		if prd(cur) {
			r = append(r, cur)
		}
		if cur.Name == "" {
			fStrct, ok := cur.Typ.(*Struct)
			if ok {
				flds.Ins(0, fStrct.Flds...)
			}
		}
	}
	return &r
}

func (x *Flds) TypRefs() (r Typs) {
	for _, f := range *x {
		if f.Typ != nil {
			r = append(r, f.Typ)
		}
	}
	return r
}
func (x *Flds) WriteFlds(b *strings.Builder, f *FleBse) {
	for _, v := range *x {
		v.WriteFld(b, f)
	}
}

func (x *Fld) IsTrm() bool {
	return x.IsFstUpr() &&
		x.Typ.Bse().IsTrm() &&
		!x.IsLitSkp() &&
		x.Typ.Bse().LitTrm != nil
}
func (x *Fld) IsEmbeddedBytWrt() bool {
	return x.Name == "" && (x.Typ.Bse().PrtStructBytWrt() != nil || x.Typ.Bse().PrtArrBytWrt() != nil)
}
func (x *Fld) WriteFld(b *strings.Builder, f *FleBse) {
	b.WriteString(x.Name)
	b.WriteRune(' ')
	if x.Mod.IsSlice() {
		b.WriteString("[]")
	}
	if x.Mod.IsPtr() {
		b.WriteRune('*')
	}
	// sys.Logf(" -- %v.%v.%v %v %v", x.Struct.Pkg.Name, x.Struct.Title(), x.Name, x.Typ == nil, reflect.ValueOf(x.Typ).IsNil())
	typ := x.Typ.Ref(f)
	if x.Name == "" && x.Typ.Bse().IsPtr() && len(typ) != 0 {
		typ = typ[1:] // allow typ embed without ptr
	}
	b.WriteString(typ)
	b.WriteRune('\n')
}
