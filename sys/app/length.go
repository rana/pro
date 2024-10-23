package app

type (
	Length struct {
		LengthX
		LengthY
	}
	LengthX struct {
		Left  int
		Right int
	}
	LengthY struct {
		Btm int
		Top int
	}
)

func NewLength(left, right, bottom, top int) (r Length) {
	r.Left = left
	r.Right = right
	r.Btm = bottom
	r.Top = top
	return r
}
func NewLengthX(left, right int) (r LengthX) {
	r.Left = left
	r.Right = right
	return r
}
func NewLengthY(bottom, top int) (r LengthY) {
	r.Btm = bottom
	r.Top = top
	return r
}
func (x *LengthX) Set(leftAndRight int) {
	x.Left = leftAndRight
	x.Right = leftAndRight
}
func (x *LengthX) Width() int { return x.Left + x.Right }
func (x *LengthY) Set(bottomAndTop int) {
	x.Btm = bottomAndTop
	x.Top = bottomAndTop
}
func (x *LengthY) Height() int { return x.Btm + x.Top }
