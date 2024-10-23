package hst

import (
	"sys/bsc/bnd"
	"sys/bsc/flt"
	"sys/bsc/flts"
	"sys/bsc/tmes"
	"sys/bsc/unt"
)

type (
	StmAtSeg struct {
		bnd.Bnd
		Stm    *StmBse
		AtTmes *tmes.Tmes
		Out    *flts.Flts
	}
	StmAtfSeg struct {
		bnd.Bnd
		Stm    *StmBse
		AtTmes *tmes.Tmes
		Out    []float32
	}
)

func (x *StmAtSeg) Act() {
	for n := x.Idx; n < x.Lim; n++ { // expects all 'AtTmes' exist in Stm.Tmes by GapFil
		idx := x.Stm.Tmes.SrchIdxEql((*x.AtTmes)[n])
		if int(idx) < len(*x.Stm.Tmes) && (*x.Stm.Tmes)[idx] == (*x.AtTmes)[n] {
			*x.Out = append(*x.Out, (*x.Stm.Vals)[idx])
		}
	}
}
func (x *StmAtfSeg) Act() {
	for n := x.Idx; n < x.Lim; n++ { // expects all 'AtTmes' exist in Stm.Tmes by GapFil
		idx := x.Stm.Tmes.SrchIdxEql((*x.AtTmes)[n])
		if int(idx) < len(*x.Stm.Tmes) && (*x.Stm.Tmes)[idx] == (*x.AtTmes)[n] {
			x.Out = append(x.Out, float32((*x.Stm.Vals)[idx]))
		}
	}
}

func (x *StmRte1Sar) Calc() {
	// PLL UNAVAILABLE DUE TO UNKNOWN POS OF UP-DWN TRANSTIONS
	// http://stockcharts.com/school/doku.php?id=chart_school:technical_indicators:parabolic_sar
	// https://www.fxpro.co.uk/help-section/articles/fxpro-quant/parabolic-stop-reverse-system-(parabolic-sar)
	side := x.Side.Bse()
	inrvl := side.Inrvl.Bse()
	x.Tmes = inrvl.Tmes
	x.Vals = flts.Make(inrvl.TmeBnds.Cnt())
	var sar, ep, af flt.Flt

	// fst sar
	vals := side.Vals.InBnd(side.ValBnds.At(0))
	isLong := vals.Lst().Sub(vals.Fst()).Geq(0)
	newHi := vals.Max()
	newLo := vals.Min()
	if isLong {
		ep = newHi
		sar = newLo
	} else {
		ep = newLo
		sar = newHi
	}
	x.Vals.Upd(0, sar)
	for n := unt.One; n < side.ValBnds.Cnt(); n++ {
		vals = side.Vals.InBnd(side.ValBnds.At(n))
		prvLo, prvHi := newLo, newHi
		newLo, newHi = vals.Min(), vals.Max()
		if isLong { // LONG
			// USE newLo < sar WITH GAP FILLED DATA; ORIGINAL ALGO HAS newLo <= sar FOR NON-GAP FILLED
			if newLo < sar { // LONG-TO-SHRT: Switch to short if the low penetrates the SAR value
				isLong = false
				sar = ep
				if sar < prvHi { // Make sure the overide SAR is within prv range
					sar = prvHi
				}
				if sar < newHi { // Make sure the overide SAR is within new range
					sar = newHi
				}
				x.Vals.Upd(n, sar)      // Output sar
				af = x.AfInc            // reset AF to init value
				ep = newLo              // reset EP to init shrt value
				sar = sar + af*(ep-sar) // calculate new long sar
				if sar < prvHi {        // Make sure the overide SAR is within prv range
					sar = prvHi
				}
				if sar < newHi { // Make sure the overide SAR is within new range
					sar = newHi
				}
			} else { // LONG: CONTINUE
				x.Vals.Upd(n, sar) // Output SAR (which was calculated in previous iteration)
				if newHi > ep {    // adjust AF and EP
					ep = newHi
					if af < x.AfMax {
						af += x.AfInc
					}
				}
				sar = sar + af*(ep-sar) // calculate new sar
				if sar > prvLo {        // Make sure the overide SAR is within prv range
					sar = prvLo
				}
				if sar > newLo { // Make sure the overide SAR is within new range
					sar = newLo
				}
			}
		} else { // SHRT
			// USE newHi > sar WITH GAP FILLED DATA; ORIGINAL ALGO HAS newHi >= sar FOR NON-GAP FILLED
			if newHi > sar { // SHRT-TO-LONG: Switch to long if the high penetrates the SAR value
				isLong = true
				sar = ep
				if sar > prvLo { // Make sure the overide SAR is within prv range
					sar = prvLo
				}
				if sar > newLo { // Make sure the overide SAR is within new range
					sar = newLo
				}
				x.Vals.Upd(n, sar)      // Output SAR
				af = x.AfInc            // reset AF to init value
				ep = newHi              // reset EP to init long value
				sar = sar + af*(ep-sar) // calculate the new SAR value
				if sar > prvLo {        // Make sure the overide SAR is within prv range
					sar = prvLo
				}
				if sar > newLo { // Make sure the overide SAR is within new range
					sar = newLo
				}
			} else { // SHRT: CONTINUE
				x.Vals.Upd(n, sar) // Output SAR (which was calculated in previous iteration)
				if newLo < ep {
					ep = newLo
					if af < x.AfMax {
						af += x.AfInc
					}
				}
				sar = sar + af*(ep-sar) // calculate the new SAR value
				if sar < prvHi {        // Make sure the overide SAR is within prv range
					sar = prvHi
				}
				if sar < newHi { // Make sure the overide SAR is within new range
					sar = newHi
				}
			}
		} // long/shrt
	} // for
}

// TODO: SIMPLIFY?
func AlignStmOtr(x, a Stm, off unt.Unt) (xBnd, aBnd bnd.Bnd) {
	if x.Bse().Tmes.Cnt() == 0 || a.Bse().Tmes.Cnt() <= off {
		return bnd.Bnd{}, bnd.Bnd{}
	}
	xBse, aBse := x.Bse(), a.Bse()
	if xBse.Tmes.Fst() != aBse.Tmes.Fst() { // align front
		found := false
		if xBse.Tmes.Fst() > aBse.Tmes.Fst() { // x-tme gtr
			for xN := unt.Zero; xN < xBse.Tmes.Cnt(); xN++ {
				aN := aBse.Tmes.SrchIdxEql(xBse.Tmes.At(xN))
				if aN != aBse.Tmes.Cnt() {
					xBnd.Idx = xN
					aBnd.Idx = aN
					found = true
					break
				}
			}
		} else { // a-tme gtr
			for aN := unt.Zero; aN < aBse.Tmes.Cnt(); aN++ {
				xN := xBse.Tmes.SrchIdxEql(aBse.Tmes.At(aN))
				if xN != xBse.Tmes.Cnt() {
					xBnd.Idx = xN
					aBnd.Idx = aN
					found = true
					break
				}
			}
		}
		if !found { // no front alignment
			return bnd.Bnd{}, bnd.Bnd{}
		}
	}
	// align back
	found := false
	for n := unt.Zero; n < xBse.Tmes.Cnt(); n++ {
		xN := xBse.Tmes.Cnt() - n - 1 // x-idx end
		aN := aBse.Tmes.SrchIdxEql(xBse.Tmes.At(xN))
		if aN != aBse.Tmes.Cnt() && aN+off < aBse.Tmes.Cnt() {
			xBnd.Lim = xN + 1
			aBnd.Lim = aN + off + 1
			found = true
			break
		}
	}
	if !found { // no back alignment
		return bnd.Bnd{}, bnd.Bnd{}
	}
	if xBnd.Idx >= xBnd.Lim || aBnd.Idx >= aBnd.Lim { // invalid range
		return bnd.Bnd{}, bnd.Bnd{}
	}
	return xBnd, aBnd
}
