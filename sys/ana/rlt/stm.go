package rlt

import (
	"sys"
	"sys/ana"
)

func (x *StmRte1Sar) Rx(inPkt ana.TmeFlts) (r []sys.Act) {
	x.mu.Lock()
	if ana.Cfg.Trc.IsRltStm() {
		sys.Logf("rlt.StmRte1Sar(%v).Rx %p inPkt %v", x.Prm(), x, inPkt)
	}
	outPkt := ana.TmeFlt{Tme: inPkt.Tme}
	newLo, newHi := inPkt.Flts.Min(), inPkt.Flts.Max()
	if x.Sar == 0 { // fst x.Sar
		x.IsLong = inPkt.Flts.Lst().Sub(inPkt.Flts.Fst()).Geq(0)
		if x.IsLong {
			x.Ep = newHi
			x.Sar = newLo
		} else {
			x.Ep = newLo
			x.Sar = newHi
		}
		outPkt.Flt = x.Sar
	} else { // non-fst
		if x.IsLong { // LONG
			if newLo < x.Sar { // LONG-TO-SHRT: Switch to short if the low penetrates the SAR value
				x.IsLong = false
				x.Sar = x.Ep
				if x.Sar < x.PrvHi { // Make sure the overide SAR is within prv range
					x.Sar = x.PrvHi
				}
				if x.Sar < newHi { // Make sure the overide SAR is within new range
					x.Sar = newHi
				}
				outPkt.Flt = x.Sar                // Output SAR
				x.Af = x.AfInc                    // reset AF to init value
				x.Ep = newLo                      // reset EP to init shrt value
				x.Sar = x.Sar + x.Af*(x.Ep-x.Sar) // calculate new long SAR
				if x.Sar < x.PrvHi {              // Make sure the overide SAR is within prv range
					x.Sar = x.PrvHi
				}
				if x.Sar < newHi { // Make sure the overide SAR is within new range
					x.Sar = newHi
				}
			} else { // LONG: CONTINUE
				outPkt.Flt = x.Sar // Output SAR (which was calculated in previous iteration)
				if newHi > x.Ep {  // adjust AF and EP
					x.Ep = newHi
					if x.Af < x.AfMax {
						x.Af += x.AfInc
					}
				}
				x.Sar = x.Sar + x.Af*(x.Ep-x.Sar) // calculate new SAR
				if x.Sar > x.PrvLo {              // Make sure the overide SAR is within prv range
					x.Sar = x.PrvLo
				}
				if x.Sar > newLo { // Make sure the overide SAR is within new range
					x.Sar = newLo
				}
			}
		} else { // SHRT
			if newHi > x.Sar { // SHRT-TO-LONG: Switch to long if the high penetrates the SAR value
				x.IsLong = true
				x.Sar = x.Ep
				if x.Sar > x.PrvLo { // Make sure the overide SAR is within prv range
					x.Sar = x.PrvLo
				}
				if x.Sar > newLo { // Make sure the overide SAR is within new range
					x.Sar = newLo
				}
				outPkt.Flt = x.Sar                // Output SAR
				x.Af = x.AfInc                    // reset AF to init value
				x.Ep = newHi                      // reset EP to init long value
				x.Sar = x.Sar + x.Af*(x.Ep-x.Sar) // calculate the new SAR value
				if x.Sar > x.PrvLo {              // Make sure the overide SAR is within prv range
					x.Sar = x.PrvLo
				}
				if x.Sar > newLo { // Make sure the overide SAR is within new range
					x.Sar = newLo
				}
			} else { // SHRT: CONTINUE
				outPkt.Flt = x.Sar // Output SAR (which was calculated in previous iteration)
				if newLo < x.Ep {
					x.Ep = newLo
					if x.Af < x.AfMax {
						x.Af += x.AfInc
					}
				}
				x.Sar = x.Sar + x.Af*(x.Ep-x.Sar) // calculate the new SAR value
				if x.Sar < x.PrvHi {              // Make sure the overide SAR is within prv range
					x.Sar = x.PrvHi
				}
				if x.Sar < newHi { // Make sure the overide SAR is within new range
					x.Sar = newHi
				}
			}
		}
	}
	x.PrvLo, x.PrvHi = newLo, newHi
	for _, rx := range x.Rxs {
		r = append(r, ana.NewTmeFltTx(outPkt, rx))
	}
	x.mu.Unlock()
	return r
}
