package tpl

import (
	"fmt"
	"sys/k"
)

type (
	PrtArrBytWrt struct {
		PrtBse
		Arr *Arr
	}
)

func (x *PrtArrBytWrt) InitPrtTyp() {
	x.Arr = x.f.GetPrt((*PrtArr)(nil)).(*PrtArr).Arr
}
func (x *PrtArrBytWrt) InitPrtTypFn() {
	if !x.t.IsBytSkp() {
		x.bytWrt()
		x.bytRed()
	}
}
func (x *PrtArrBytWrt) bytWrt() (r *TypFn) {
	r = x.f.TypFn(k.BytWrt)
	r.InPrm(BufferPtr, "b")
	r.Addf("bLen := make([]byte, 4) // array length")
	r.Addf("binary.LittleEndian.PutUint32(bLen, uint32(len(*x)))")
	r.Addf("b.Write(bLen)")
	r.Add("for _, v := range *x {")
	r.Add("v.BytWrt(b)")
	r.Add("}")
	return r
}
func (x *PrtArrBytWrt) bytRed() (r *TypFn) {
	x.f.Import("encoding/binary")
	r = x.f.TypFn(k.BytRed)
	r.InPrmSlice(Byte, "b")
	r.OutPrm(Int, "idx")
	r.Add("if len(b) >= 4 {")
	r.Addf("*x = make(%v, binary.LittleEndian.Uint32(b[:4])) // overwrite any previous existing", x.t.Title())
	r.Add("idx = 4")
	r.Add("for n := 0; n < len(*x); n++ {")
	// size := fmt.Sprintf("%v.Size", x.Arr.Elm.Pkg.Ref(x.f))
	size := fmt.Sprintf("%v", x.Arr.Elm.Size.Ref(x.f))
	if x.Arr.Elm.IsStruct() && x.Arr.Elm.IsPtr() {
		r.Addf("(*x)[n] = %v{}", x.Arr.Elm.Adr(x.f))
	}
	r.Addf("(*x)[n].BytRed(b[idx:idx+%v])", size)
	r.Addf("idx += %v", size)
	r.Add("}")
	r.Add("}")
	r.Add("return idx")
	return r
}
