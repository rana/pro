package ana

import (
	"sys"
	"sys/bsc/flt"
	"sys/bsc/tme"
	"sys/lng/jsn"
)

type (
	InstrTic struct {
		Pkt TmeIdx
		I   *Instr
	}
)

func (x *Instr) CalcStats() {
	if x.HstStm == nil {
		return
	}
	spds := x.HstStm.Spds(x)
	x.SpdMin = spds.Min()
	x.SpdMax = spds.Max()
	x.SpdMdn = spds.Mdn()
	x.SpdAvg = spds.Sma().Trnc(2)
	x.SpdStd = spds.Std().Trnc(2)
	x.SpdOpnLim = x.SpdMdn.Mul(x.SpdAvg).Sqrt().Mul(1.5).Trnc(2)
	x.Fst = x.HstStm.Tmes.Fst()
	x.Lst = x.HstStm.Tmes.Lst()
	x.TmeCnt = x.HstStm.Tmes.Cnt()
	x.DayCnt = x.Fst.WeekdayCnt(x.Lst)
}
func (x *Instr) CalcMktWeeks(mktHr tme.Tme) { // pass var due to tst bug
	x.MktWeeks = tme.NewRngs()
	x.MktDays = tme.NewRngs()
	x.MktHrs = tme.NewRngs()
	fst := x.Fst.ToSunday().Add(mktHr)
	lst := x.Lst.ToFriday().Add(mktHr)
	for cur := fst; cur < lst; cur = cur.Add(tme.Week).ToSunday().Add(mktHr) {
		x.MktWeeks.Push(tme.NewRng(cur, cur.ToFriday().Add(mktHr)))
		sun := cur
		mon := cur.ToMonday().Add(mktHr)
		tue := cur.ToTuesday().Add(mktHr)
		wed := cur.ToWednesday().Add(mktHr)
		thr := cur.ToThursday().Add(mktHr)
		fri := cur.ToFriday().Add(mktHr)
		x.MktDays.Push(tme.NewRng(sun, mon))
		x.MktDays.Push(tme.NewRng(mon, tue))
		x.MktDays.Push(tme.NewRng(tue, wed))
		x.MktDays.Push(tme.NewRng(wed, thr))
		x.MktDays.Push(tme.NewRng(thr, fri))
		for curHr := sun; curHr < lst; curHr += tme.Hour {
			x.MktHrs.Push(tme.NewRng(curHr, curHr+tme.Hour))
		}
	}
	// sys.Logf("Instr.CalcMktWeeks fst:%v lst:%v rngs:%v", x.Fst, x.Lst, x.MktWeeks)
}

func MktWeek(v tme.Tme) (r tme.Rng) {
	r.Min = v.ToSunday().Add(Cfg.MktHr)
	r.Max = v.ToFriday().Add(Cfg.MktHr)
	sys.Logf("ana.MktWeek v:%v r:%v", v, r)
	return r
}
func MktWeekMin(v tme.Tme) (r tme.Tme) {
	return v.ToSunday().Add(Cfg.MktHr)
}
func MktWeekMax(v tme.Tme) (r tme.Tme) {
	return v.ToFriday().Add(Cfg.MktHr)
}
func MktNextWeekMin(v tme.Tme) (r tme.Tme) {
	return v.Add(tme.Week).ToSunday().Add(Cfg.MktHr)
}
func MktPrvWeekMax(v tme.Tme) (r tme.Tme) {
	if v.IsSaturday() {
		return v.ToFriday().Add(Cfg.MktHr)
	}
	return v.Sub(tme.Week).ToFriday().Add(Cfg.MktHr)
}

func (x *Instr) Sub(rx TmeIdxRx, id uint32) {
	x.RltSubsMu.Lock()
	if len(x.RltSubs) == 0 {
		x.Prv.Sub(x) // Sub will init RltSubs
	}
	x.RltSubs[sys.Uint64(id, 0)] = rx
	x.RltSubsMu.Unlock()
}
func (x *Instr) Unsub(id uint32) {
	x.RltSubsMu.Lock()
	delete(x.RltSubs, sys.Uint64(id, 0))
	if len(x.RltSubs) == 0 {
		x.Prv.Unsub(x)
	}
	x.RltSubsMu.Unlock()
}
func (x *Instr) OanJsnRed(txt string) {
	var j jsn.Jsnr
	j.Reset(txt)
	arrBnds := j.ArrObjs("instruments")
	if len(arrBnds) != 0 {
		j.Reset(txt[arrBnds[0].Idx:arrBnds[0].Lim])
		x.Name = j.Str("name").Lower()
		x.Typ = j.Str("type")
		pip := j.Int("pipLocation") // "pipLocation": -4,
		mag := 1
		for n := pip; n < 0; n++ { // convert pipLocation to single floating point pip i.e., .0001
			mag *= 10
		}
		x.Pip = flt.One.Div(flt.Flt(mag))
		x.MrgnRtio = flt.One.Div(j.StrFlt("marginRate")).Trnc(0) // 1 / .02 = 50   (50:1 eur_usd)
		x.DisplayPrecision = j.Unt("displayPrecision")
		x.TradeUnitsPrecision = j.Unt("tradeUnitsPrecision")
		x.MinTrdSize = j.StrUnt("minimumTradeSize")
		x.MaxTrailingStopDistance = j.StrFlt("maximumTrailingStopDistance")
		x.MinTrailingStopDistance = j.StrFlt("minimumTrailingStopDistance")
		x.MaxPositionSize = j.StrUnt("maximumPositionSize")
		x.MaxOrderUnits = j.StrUnt("maximumOrderUnits")
	}
}
func (x *Instr) Spd(bid, ask flt.Flt) flt.Flt { return ask.Sub(bid).Div(x.Pip).Trnc(2) }
func (x *Instr) Pipette() flt.Flt             { return 0.1 * x.Pip }
func (x *Instr) PipetteScl() flt.Flt          { return 10.0 / x.Pip }
func (x *Instr) PipScl() flt.Flt              { return 1.0 / x.Pip }
