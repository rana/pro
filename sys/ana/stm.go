package ana

import (
	"sys/bsc/bnd"
	"sys/bsc/flt"
	"sys/bsc/flts"
	"sys/bsc/tme"
	"sys/bsc/tmes"
	"sys/bsc/unt"
	"sys/bsc/unts"
	"sys/err"
	"sys/lng/jsn"
)

const (
	MktSessionGap = tme.D1
)

func NewStm() (r *Stm) {
	r = &Stm{}
	r.Tmes = tmes.New()
	r.Bids = flts.New()
	r.Asks = flts.New()
	r.BidLims = unts.New()
	r.AskLims = unts.New()
	return r
}
func (x *Stm) Cnt() unt.Unt { return x.Tmes.Cnt() }
func (x *Stm) Rng() (r tme.Rng) {
	r.Min = x.Tmes.Fst()
	r.Max = x.Tmes.Lst()
	return r
}
func (x *Stm) Spds(i *Instr) (r *flts.Flts) {
	r = flts.Make(x.Tmes.Cnt())
	for n := unt.Zero; n < x.Tmes.Cnt(); n++ {
		r.Upd(n, i.Spd(x.BidAt(n), x.AskAt(n)))
	}
	return r
}
func (x *Stm) BidsAsks(t tme.Tme) (bids, asks *flts.Flts) {
	idx := x.Tmes.SrchIdxEql(t)
	if idx == unt.Unt(len(*x.Tmes)) {
		return bids, asks
	}
	return x.BidsAsksAt(idx)
}
func (x *Stm) BidAsk(t tme.Tme) (bid, ask flt.Flt) {
	idx := x.Tmes.SrchIdxEql(t)
	if idx == unt.Unt(len(*x.Tmes)) {
		return bid, ask
	}
	return x.BidAskAt(idx)
}
func (x *Stm) BidsAt(idx unt.Unt) *flts.Flts {
	// Bid: the price for me to sell at
	var b bnd.Bnd
	if idx > 0 {
		b.Idx = x.BidLims.At(idx - 1)
	}
	b.Lim = x.BidLims.At(idx)
	return x.Bids.InBnd(b)
}
func (x *Stm) BidAt(idx unt.Unt) flt.Flt {
	// Bid: the price for me to sell at
	return x.BidsAt(idx).Lst()
}
func (x *Stm) AsksAt(idx unt.Unt) *flts.Flts {
	// Ask: the price for me to buy at
	var b bnd.Bnd
	if idx > 0 {
		b.Idx = x.AskLims.At(idx - 1)
	}
	b.Lim = x.AskLims.At(idx)
	return x.Asks.InBnd(b)
}
func (x *Stm) AskAt(idx unt.Unt) flt.Flt {
	// Ask: the price for me to buy at
	return x.AsksAt(idx).Lst()
}
func (x *Stm) BidsAsksAt(idx unt.Unt) (bids, asks *flts.Flts) {
	// Bid: the price for me to sell at
	// Ask: the price for me to buy at
	return x.BidsAt(idx), x.AsksAt(idx)
}
func (x *Stm) BidAskAt(idx unt.Unt) (bid, ask flt.Flt) {
	// Bid: the price for me to sell at
	// Ask: the price for me to buy at
	return x.BidsAt(idx).Lst(), x.AsksAt(idx).Lst()
}
func (x *Stm) In(idx, lim unt.Unt) (r *Stm) {
	r = &Stm{}
	r.Tmes = x.Tmes.In(idx, lim)
	tmeBnd := bnd.Bnd{Idx: idx, Lim: lim}
	bidBnd := x.BidBndByTmeBnd(tmeBnd)
	askBnd := x.AskBndByTmeBnd(tmeBnd)
	r.Bids = x.Bids.InBnd(bidBnd)
	r.Asks = x.Asks.InBnd(askBnd)
	r.BidLims = x.BidLims.In(idx, lim) //.Cpy() // cpy to adjust
	r.AskLims = x.AskLims.In(idx, lim) //.Cpy() // cpy to adjust
	return r
}
func (x *Stm) InRng(fromTo ...tme.Tme) (r *Stm) {
	if x.Tmes.Cnt() == 0 {
		return x
	}
	var idx, lim unt.Unt
	if len(fromTo) == 0 {
		return x
	} else if len(fromTo) == 1 {
		idx = x.Tmes.SrchIdx(fromTo[0], true)
		lim = x.Tmes.Cnt()
	} else if len(fromTo) == 2 {
		if fromTo[0] > fromTo[1] {
			fromTo[0], fromTo[1] = fromTo[1], fromTo[0]
		}
		// sys.Log("Ana.Stm.InRng", "fromTo[0]", fromTo[0])
		// sys.Log("Ana.Stm.InRng", "fromTo[1]", fromTo[1])

		idx = x.Tmes.SrchIdx(fromTo[0], true)
		lim = x.Tmes.SrchIdx(fromTo[1], true)

	}
	if idx >= x.Tmes.Cnt() {
		return NewStm()
	}
	if lim > x.Tmes.Cnt() {
		lim = x.Tmes.Cnt()
	}
	return x.In(idx, lim)
}
func (x *Stm) From(idx unt.Unt) (r *Stm) { return x.In(idx, x.Tmes.Cnt()) }
func (x *Stm) To(lim unt.Unt) (r *Stm)   { return x.In(0, lim) }
func (x *Stm) Cpy() (r *Stm) {
	r = &Stm{}
	r.Tmes = x.Tmes.Cpy()
	r.Bids = x.Bids.Cpy()
	r.Asks = x.Asks.Cpy()
	r.BidLims = x.BidLims.Cpy()
	r.AskLims = x.AskLims.Cpy()
	return r
}
func (x *Stm) Clr() (r *Stm) {
	r = &Stm{}
	r.Tmes = x.Tmes.Clr()
	r.Bids = x.Bids.Clr()
	r.Asks = x.Asks.Clr()
	r.BidLims = x.BidLims.Clr()
	r.AskLims = x.AskLims.Clr()
	return r
}
func (x *Stm) Tic(idx unt.Unt) (r *Tic) {
	r = &Tic{}
	r.Tme = x.Tmes.At(idx)
	r.Bids = x.BidsAt(idx)
	r.Asks = x.AsksAt(idx)
	return r
}
func (x *Stm) Push(v *Stm) *Stm {
	x.Tmes.Push(*v.Tmes...)
	x.Bids.Push(*v.Bids...)
	x.Asks.Push(*v.Asks...)
	x.BidLims.Push(*v.BidLims...)
	x.AskLims.Push(*v.AskLims...)
	return x
}
func (x *Stm) PushTic(v *Tic) {
	x.Tmes.Push(v.Tme)
	x.Bids.Push(*v.Bids...)
	x.Asks.Push(*v.Asks...)
	x.BidLims.Push(x.Bids.Cnt())
	x.AskLims.Push(x.Asks.Cnt())
}
func (x *Stm) PushGap(gapTme tme.Tme, v *Tic) {
	x.Tmes.Push(gapTme)
	x.Bids.Push(*v.Bids...)
	x.Asks.Push(*v.Asks...)
	x.BidLims.Push(x.Bids.Cnt())
	x.AskLims.Push(x.Asks.Cnt())
}
func (x *Stm) BidBndByTmeBnd(tmeBnd bnd.Bnd) (r bnd.Bnd) {
	if tmeBnd.Idx != 0 {
		r.Idx = (*x.BidLims)[tmeBnd.Idx-1] // bid idx is prv lim
	}
	r.Lim = (*x.BidLims)[tmeBnd.Lim-1]
	return r
}
func (x *Stm) AskBndByTmeBnd(tmeBnd bnd.Bnd) (r bnd.Bnd) {
	if tmeBnd.Idx != 0 {
		r.Idx = (*x.AskLims)[tmeBnd.Idx-1] // bid idx is prv lim
	}
	r.Lim = (*x.AskLims)[tmeBnd.Lim-1]
	return r
}
func (x *Stm) BidsByTmeBnd(tmeBnd bnd.Bnd) *flts.Flts {
	return x.Bids.InBnd(x.BidBndByTmeBnd(tmeBnd))
}
func (x *Stm) AsksByTmeBnd(tmeBnd bnd.Bnd) *flts.Flts {
	return x.Asks.InBnd(x.AskBndByTmeBnd(tmeBnd))
}

func (x *Stm) OanJsnRed(j *jsn.Jsnr) {
	// assume Stm and instr are correctly aligned: eur/usd -> eur/usd
	// assume current value is tradeable (tradeable check outside method)
	// AGGREGATE WITHIN 1 SECOND TO SIMPLIFY STGY PKT ORDERING AND OVERALL CALCULATION
	// AS LONG AS tme.Tme is 32-bit
	// NEED TO COORDINATE WITH HART BEAT TO ENSURE AGGREGATE
	// VALUES ARE SENT PROPERLY
	txt := j.Txt
	curTme := j.StrTme("time")
	if x.Tmes.Cnt() > 0 && x.Tmes.Lst() == curTme {
		// aggregate ticks with eql tme; 1s resolution
		// tme already exists; updating existing Lims, pushing prices
		bidObjs := j.ArrObjs("bids")
		if len(bidObjs) != 0 {
			for _, o := range bidObjs {
				j.Reset(txt[o.Idx:o.Lim])
				x.Bids.Push(j.StrFlt("price"))
			}
			x.BidLims.Upd(x.BidLims.LstIdx(), x.Bids.Cnt()) // extend existing bid lim
		}
		j.Reset(txt)
		askObjs := j.ArrObjs("asks")
		if len(askObjs) != 0 {
			for _, o := range askObjs {
				j.Reset(txt[o.Idx:o.Lim])
				x.Asks.Push(j.StrFlt("price"))
			}
			x.AskLims.Upd(x.AskLims.LstIdx(), x.Asks.Cnt()) // extend existing ask lim
		}
	} else { // non-aggregate tick
		x.Tmes.Push(curTme)
		bidObjs := j.ArrObjs("bids")
		if len(bidObjs) != 0 {
			for _, o := range bidObjs {
				j.Reset(txt[o.Idx:o.Lim])
				x.Bids.Push(j.StrFlt("price"))
			}
			x.BidLims.Push(x.Bids.Cnt()) // push new bid lim
		}
		j.Reset(txt)
		askObjs := j.ArrObjs("asks")
		if len(askObjs) != 0 {
			for _, o := range askObjs {
				j.Reset(txt[o.Idx:o.Lim])
				x.Asks.Push(j.StrFlt("price"))
			}
			x.AskLims.Push(x.Asks.Cnt()) // push new ask lim
		}
	}
}

func (x *Stm) GapFilTo(heartTme tme.Tme) {
	if x.Tmes.Lst() != heartTme {
		lstTic := x.Tic(x.Tmes.LstIdx())
		for gapTme := x.Tmes.Lst() + tme.Resolution; gapTme <= heartTme; gapTme += tme.Resolution {
			gapTic := &Tic{}
			gapTic.Tme = gapTme
			gapTic.Bids = lstTic.Bids
			gapTic.Asks = lstTic.Asks
			x.PushTic(gapTic)
		}
	}
}
func (x *Stm) GapFilLst() (r TmeIdx) {
	if len(*x.Tmes) != 0 {
		// lst: gap fill from last point to mkt session close
		lst := x.Tmes.Lst()
		lstMktWeekMax := MktWeekMax(lst)
		if lst+tme.Resolution < lstMktWeekMax {
			tic := x.Tic(x.Tmes.LstIdx())
			for gapTme := lst + tme.Resolution; gapTme < lstMktWeekMax; gapTme += tme.Resolution { // less than max
				x.PushGap(gapTme, tic)
			}
		}
	}
	return r
}
func (x *Stm) GapFil(skpMktWeek ...bool) {
	// fill empty data gap between PRV and LST
	// using 1s inrvl resolution with 32-bit tme.Tme
	if len(*x.Tmes) == 1 {
		if len(skpMktWeek) == 0 {
			// fst: gap fill from mkt session open to first point
			fst := (*x.Tmes)[0]
			fstMktWeekMin := MktWeekMin(fst)
			if fstMktWeekMin < fst {
				tic := x.Tic(0)
				x.Clr() // clear all data in stm (only fst point)
				for gapTme := fstMktWeekMin; gapTme <= fst; gapTme += tme.Resolution {
					x.PushGap(gapTme, tic)
				}
			}
		}
	} else if len(*x.Tmes) > 1 {
		prvTme := x.Tmes.At(x.Tmes.LstIdx() - 1)
		lstTme := x.Tmes.Lst()
		if lstTme-prvTme > tme.Resolution {
			prvTic := x.Tic(x.Tmes.LstIdx() - 1)
			lstTic := x.Tic(x.Tmes.LstIdx())
			*x = *x.To(x.Tmes.LstIdx())        // trim lst which will be over written with gap value(s)
			if lstTme-prvTme < MktSessionGap { // IN MKT SESSION
				for gapTme := prvTme + tme.Resolution; gapTme < lstTme; gapTme += tme.Resolution {
					x.PushGap(gapTme, prvTic) // push gap
				}
			} else { // BETWEEN MKT SESSION
				lst := prvTic.Tme // lst: gap fill from last point to mkt session close
				lstMktWeekMax := MktWeekMax(lst)
				if lst+tme.Resolution < lstMktWeekMax {
					for gapTme := lst + tme.Resolution; gapTme < lstMktWeekMax; gapTme += tme.Resolution { // less than max
						x.PushGap(gapTme, prvTic)
					}
				}
				fst := lstTic.Tme // fst: gap fill from mkt session open to first point
				fstMktWeekMin := MktWeekMin(fst)
				if fstMktWeekMin < fst {
					for gapTme := fstMktWeekMin; gapTme < fst; gapTme += tme.Resolution { // less than fst; PushTic(lstTic) will push fst
						x.PushGap(gapTme, prvTic)
					}
				}
				// sys.Log("prvTme", prvTme)
				// sys.Log("lstTme", lstTme)
				// for n := idx - 1; n < idx+cnt+2; n++ {
				// 	sys.Log(x.Tmes.At(n), x.Bids.In(x.BidLims.At(n-1), x.BidLims.At(n)), x.BidLims.At(n))
				// }
				// sys.Log("-")
			}
			x.PushTic(lstTic) // push lst

		}
	}
}
func (x *Stm) Validate() {
	x.ValidateTmeGap()
	x.ValidateLen()
	x.ValidateLims()
}
func (x *Stm) ValidateTmeGap() {
	for n := unt.One; n < x.Tmes.Cnt(); n++ {
		delta := x.Tmes.At(n).Sub(x.Tmes.At(n - 1))
		if delta < MktSessionGap && delta > tme.Resolution {
			err.Panicf("INVALID GAP (n-1:%v:%v n:%v:%v)", n-1, x.Tmes.At(n-1), n, x.Tmes.At(n))
		}
	}
}
func (x *Stm) ValidateLen() {
	if x.Tmes.Cnt() != x.BidLims.Cnt() || x.Tmes.Cnt() != x.AskLims.Cnt() {
		err.Panicf("INVALID LEN (Tmes:%v BidLims:%v AskLims:%v)",
			x.Tmes.Cnt(), x.BidLims.Cnt(), x.AskLims.Cnt())
	}
}
func (x *Stm) ValidateLims() {
	x.ValidateOneLims("BidLims", x.BidLims)
	x.ValidateOneLims("AskLims", x.AskLims)
}
func (x *Stm) ValidateOneLims(name string, vs *unts.Unts) {
	for n := unt.One; n < vs.Cnt(); n++ {
		if vs.At(n-1) > vs.At(n) {
			err.Panicf("INVALID %v (n:%v vs.At(n-1):%v vs.At(n):%v)", name, n, vs.At(n-1), vs.At(n))
		}
	}
}
