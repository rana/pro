package vis

type (
	LenX struct {
		Lft uint32
		Rht uint32
	}
	LenY struct {
		Btm uint32
		Top uint32
	}
	LenXY struct {
		LenX
		LenY
	}
)

func NewLenX(lft, rht uint32) (r LenX) {
	r.Lft = lft
	r.Rht = rht
	return r
}
func NewLenY(btm, top uint32) (r LenY) {
	r.Btm = btm
	r.Top = top
	return r
}
func NewLenXY(lft, rht, btm, top uint32) (r LenXY) {
	r.Lft = lft
	r.Rht = rht
	r.Btm = btm
	r.Top = top
	return r
}

func (x *LenX) Width() uint32  { return x.Lft + x.Rht }
func (x *LenY) Height() uint32 { return x.Btm + x.Top }
