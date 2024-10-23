package tpl

import (
	"sys"
	"sys/k"
	"sys/tpl/atr"
	"sys/tpl/mod"
	"time"
)

const (
	TmeSize = 4
)

type (
	FleTme struct {
		FleBse
		PrtIdn
		PrtRel
		PrtSgn
		PrtAri
		PrtString
		PrtBytes
		// PrtLog
		// PrtIfc
		PrtPkt
		Zero      *Cnst
		One       *Cnst
		NegOne    *Cnst
		Min       *Cnst
		Max       *Cnst
		Size      *Cnst
		DurStrLim *Cnst

		S1 *Cnst
		M1 *Cnst
		H1 *Cnst
		D1 *Cnst
	}
	FleTmes struct {
		FleBse
		PrtArr
		// PrtArrIdn
		PrtArrRel
		PrtArrSrt
		PrtArrSer
		PrtArrInr
		PrtArrAgg
		PrtArrStrWrt
		PrtArrBytWrt
		// PrtLog
		// PrtIfc
	}
)

func (x *DirBsc) NewTme() (r *FleTme) {
	r = &FleTme{}
	x.Tme = r
	r.Name = k.Tme
	r.Pkg = x.Pkg.New(r.Name)
	r.Alias(r.Name, Int32, atr.TypTme)
	r.AddFle(r)
	return r
}
func (x *FleTme) NewArr() (r *FleTmes) {
	r = &FleTmes{}
	r.FleBse = *NewArr(x, &r.PrtArr)
	r.AddFle(r)
	return r
}
func (x *FleTme) InitVals(bse *TypBse) {
	bse.Lits = sys.Vs("0s", "1s", "10s", "-1s", "-10m", "1w2d3h4m5s", "-1w2d3h4m5s", "2000y1n2d3h4m5s", "2000y2d3h4m5s", "2000y3h4m5s")
	bse.Vals = sys.Vs("0", "1", "10", "-1", "-10*60", "788645", "-788645", "946782245", "946782245", "946695845")
	// 788645 = 1w2d3h4m5s = (7*24*60*60)+(2*24*60*60)+(3*60*60)+(4*60)+(5) =
}
func (x *FleTme) InitCnst() {
	x.Zero = x.Cnst(k.Zero, "0")
	x.One = x.Cnst(k.One, "1")
	x.NegOne = x.Cnst(k.NegOne, "-1")
	x.Min = x.Cnst(k.Min, "-1 << 31")
	x.Max = x.Cnst(k.Max, "1<<31-1")
	x.Cnst(k.Second, "1")
	x.Cnst(k.Minute, "60")
	x.Cnst(k.Hour, "60*60")
	x.Cnst(k.Day, "24*60*60")
	x.Cnst(k.Week, "7*24*60*60")
	x.S1 = x.Cnst(k.S1, "1")
	x.Cnst(k.S5, "5")
	x.Cnst(k.S10, "10")
	x.Cnst(k.S15, "15")
	x.Cnst(k.S20, "20")
	x.Cnst(k.S30, "30")
	x.Cnst(k.S40, "40")
	x.Cnst(k.S50, "50")
	x.M1 = x.Cnst(k.M1, "1*60")
	x.Cnst(k.M5, "5*60")
	x.Cnst(k.M10, "10*60")
	x.Cnst(k.M15, "15*60")
	x.Cnst(k.M20, "20*60")
	x.Cnst(k.M30, "30*60")
	x.Cnst(k.M40, "40*60")
	x.Cnst(k.M50, "50*60")
	x.H1 = x.Cnst(k.H1, "1*60*60")
	x.D1 = x.Cnst(k.D1, "1*60*60*24")
	x.Size = x.CnstSize(TmeSize)
	x.DurStrLim = x.Cnst("DurStrLim", "1*60*60*24*365*10") // 10 year monotonic
	x.DurStrLim.Atr = atr.None
	x.Cnst(k.Resolution, "1") // minimum resolution (1 second); perhaps it is changes in the future
}
func (x *FleTme) InitPkgFn() {

}
func (x *FleTme) InitTypFn() {
	x.Now()
	x.WeekdayCnt()
	x.strWrt()
	x.bytWrt()
	x.BytRed()
	x.NewDte()
	x.NewTme()
	x.NewDteTme()
	x.PkgTime()
	x.MemTime()
	x.Dte()
	weekdays := []time.Weekday{time.Sunday, time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday, time.Saturday}
	for _, w := range weekdays {
		x.ToDayOfWeek(w)
	}
	for _, w := range weekdays {
		x.IsDayOfWeek(w)
	}
	x.PkgDuration()
	x.MemDuration()
	x.DurWrt()
	x.DurString()
}
func (x *FleTme) InitTrm(bse *TypBse, trmr *FleTrmr) {
	x.InitLexLit(func(r *TypFn, t *FleTrmr) {
		r.InPrmVariadic(Bool, "skpDsh")
		r.Add("r.Idx = x.Idx")
		r.Add("scn := x.Scn")

		r.Addf("r.Year, ok = x.TmePrtLit('y')")
		r.Add("if ok {")
		r.Add("scn = x.Scn")
		r.Add("} else {")
		r.Add("x.Scn = scn")
		r.Add("}")

		r.Addf("r.Month, ok = x.TmePrtLit('n')")
		r.Add("if ok {")
		r.Add("scn = x.Scn")
		r.Add("} else {")
		r.Add("x.Scn = scn")
		r.Add("}")

		r.Addf("r.Week, ok = x.TmePrtLit('w')")
		r.Add("if ok {")
		r.Add("scn = x.Scn")
		r.Add("} else {")
		r.Add("x.Scn = scn")
		r.Add("}")

		r.Addf("r.Day, ok = x.TmePrtLit('d')")
		r.Add("if ok {")
		r.Add("scn = x.Scn")
		r.Add("} else {")
		r.Add("x.Scn = scn")
		r.Add("}")

		r.Addf("r.Hour, ok = x.TmePrtLit('h')")
		r.Add("if ok {")
		r.Add("scn = x.Scn")
		r.Add("} else {")
		r.Add("x.Scn = scn")
		r.Add("}")

		r.Addf("r.Minute, ok = x.TmePrtLit('m')")
		r.Add("if ok {")
		r.Add("scn = x.Scn")
		r.Add("} else {")
		r.Add("x.Scn = scn")
		r.Add("}")

		r.Addf("r.Second, ok = x.TmePrtLit('s')")
		r.Add("if ok {")
		r.Add("scn = x.Scn")
		r.Add("} else {")
		r.Add("x.Scn = scn")
		r.Add("}")

		r.Add("if r.Second.Lim == 0 && r.Minute.Lim == 0 && r.Hour.Lim == 0 && r.Day.Lim == 0 && r.Week.Lim == 0 && r.Month.Lim == 0 && r.Year.Lim == 0 {")
		r.Add("return r, false")
		r.Add("}")
		r.Add("if unicode.IsLetter(x.Ch) || unicode.IsDigit(x.Ch) || x.Ch == '_' || (len(skpDsh) == 0 && x.Ch == '-') {")
		r.Add("return r, false")
		r.Add("}")

		r.Add("r.Lim = x.Idx")
		r.Add("return r, true")
	}, trmr.StructTme())
	x.InitPrsLit(func(r *PkgFn, f *FlePrs) {
		f.Import("time")
		r.Add("year := TmePrtTrm(trm.Year, txt)")
		r.Add("if year != 0 {")
		r.Add("month := TmePrtTrm(trm.Month, txt)")
		r.Add("if month == 0 {")
		r.Add("month = 1")
		r.Add("}")
		r.Add("day := TmePrtTrm(trm.Day, txt)")
		r.Add("if day == 0 {")
		r.Add("day = 1")
		r.Add("}")
		r.Add("return tme.Time(time.Date(int(year), time.Month(month), int(day), 0, 0, 0, 0, time.UTC)) +")
		r.Add("TmePrtTrm(trm.Hour, txt)*tme.Hour +")
		r.Add("TmePrtTrm(trm.Minute, txt)*tme.Minute +")
		r.Add("TmePrtTrm(trm.Second, txt)*tme.Second")
		r.Add("}")

		r.Add("week := TmePrtTrm(trm.Week, txt) * tme.Week")
		r.Add("day := TmePrtTrm(trm.Day, txt) * tme.Day")
		r.Add("hour := TmePrtTrm(trm.Hour, txt) * tme.Hour")
		r.Add("minute := TmePrtTrm(trm.Minute, txt) * tme.Minute")
		r.Add("second := TmePrtTrm(trm.Second, txt) * tme.Second")
		r.Add("switch {")
		r.Add("case week < 0:")
		r.Add("day = -day")
		r.Add("hour = -hour")
		r.Add("minute = -minute")
		r.Add("second = -second")
		r.Add("case day < 0:")
		r.Add("hour = -hour")
		r.Add("minute = -minute")
		r.Add("second = -second")
		r.Add("case hour < 0:")
		r.Add("minute = -minute")
		r.Add("second = -second")
		r.Add("case minute < 0:")
		r.Add("second = -second")
		r.Add("}")
		r.Add("return week + day + hour + minute + second")
	})
	x.InitPrsCfg()
}
func (x *FleTme) strWrt() (r *TypFn) {
	x.Import("strconv")
	r = x.TypFn("StrWrt")
	r.InPrm(BuilderPtr, "b")
	r.Add("t := x.Time()")
	r.Addf("if x >= %v { // date/time", x.DurStrLim.Name)
	r.Add("y, n, d := t.Date()")
	r.Add("b.WriteString(strconv.FormatInt(int64(y), 10))")
	r.Add("b.WriteRune('y')")
	r.Add("b.WriteString(strconv.FormatInt(int64(n), 10))")
	r.Add("b.WriteRune('n')")
	r.Add("b.WriteString(strconv.FormatInt(int64(d), 10))")
	r.Add("b.WriteRune('d')")
	r.Add("h, m, s := t.Hour(), t.Minute(), t.Second()")
	r.Add("if h != 0 {")
	r.Add("b.WriteString(strconv.FormatInt(int64(h), 10))")
	r.Add("b.WriteRune('h')")
	r.Add("}")
	r.Add("if m != 0 {")
	r.Add("b.WriteString(strconv.FormatInt(int64(m), 10))")
	r.Add("b.WriteRune('m')")
	r.Add("}")
	r.Add("if s != 0 {")
	r.Add("b.WriteString(strconv.FormatInt(int64(s), 10))")
	r.Add("b.WriteRune('s')")
	r.Add("}")
	r.Add("} else { // dur")
	r.Add("x.DurWrt(b)")
	r.Add("}")

	return r
}
func (x *FleTme) bytWrt() (r *TypFn) {
	x.Import("encoding/binary")
	x.Import("unsafe")
	r = x.TypFn("BytWrt")
	r.InPrm(BufferPtr, "b")
	r.Add("v := make([]byte, Size)")
	r.Add("binary.LittleEndian.PutUint32(v, *(*uint32)(unsafe.Pointer(&x)))")
	r.Add("b.Write(v)")
	return r
}
func (x *FleTme) BytRed() (r *TypFn) {
	x.Import("unsafe")
	r = x.TypFn("BytRed")
	r.Rxr.Mod = mod.Ptr
	r.InPrmSlice(Byte, "b")
	r.OutPrm(Int)
	r.Add("bits := binary.LittleEndian.Uint32(b[:Size])")
	r.Add("*x = *(*Tme)(unsafe.Pointer(&bits))")
	r.Add("return Size")
	return r
}
func (x *FleTme) NewDte() (r *PkgFn) {
	r = x.PkgFn("NewDte")
	r.InPrm(Int, "y")
	r.InPrm(Int, "n")
	r.InPrm(Int, "d")
	r.OutPrm(x)
	r.Add("return Time(time.Date(y, time.Month(n), d, 0, 0, 0, 0, time.UTC))")
	return r
}
func (x *FleTme) NewTme() (r *PkgFn) {
	r = x.PkgFn("NewTme")
	r.InPrm(Int, "h")
	r.InPrm(Int, "m")
	r.InPrm(Int, "s")
	r.OutPrm(x)
	r.Add("return (Tme(h)*Hour) + (Tme(m)*Minute) + Tme(s)")
	return r
}
func (x *FleTme) NewDteTme() (r *PkgFn) {
	r = x.PkgFn("NewDteTme")
	r.InPrm(Int, "y")
	r.InPrm(Int, "n")
	r.InPrm(Int, "d")
	r.InPrm(Int, "h")
	r.InPrm(Int, "m")
	r.InPrm(Int, "s")
	r.OutPrm(x)
	r.Add("return Time(time.Date(y, time.Month(n), d, h, m, s, 0, time.UTC))")
	return r
}
func (x *FleTme) PkgTime() (r *PkgFn) {
	r = x.PkgFn(k.Time)
	r.InPrm(Time, "v")
	r.OutPrm(x)
	r.Add("return Tme(v.UTC().Unix())")
	return r
}
func (x *FleTme) MemTime() (r *TypFn) {
	r = x.TypFn(k.Time)
	r.OutPrm(Time)
	r.Addf("return time.Unix(int64(x), 0).UTC()")
	return r
}
func (x *FleTme) Dte() (r *TypFn) {
	r = x.TypFn(k.Dte)
	r.OutPrm(x)
	r.Add("y, n, d := x.Time().Date()")
	r.Add("return NewDte(y, int(n), d)")
	return r
}
func (x *FleTme) ToDayOfWeek(w time.Weekday) (r *TypFn) {
	r = x.TypFnf("To%v", w)
	r.OutPrm(x)
	r.Add("t := x.Time()")
	r.Add("y, n, d := t.Date()")
	r.Add("t = time.Date(y, n, d, 0, 0, 0, 0, t.Location()) // to start of day")
	r.Addf("t = t.Add(time.Duration(-t.Weekday()+time.%v) * time.Hour * 24) // to day of weeek", w)
	r.Add("return Time(t)")
	return r
}
func (x *FleTme) IsDayOfWeek(w time.Weekday) (r *TypFn) {
	r = x.TypFnf("Is%v", w)
	r.OutPrm(_sys.Bsc.Bol)
	r.Addf("return x.Time().Weekday() == time.%v", w)
	return r
}
func (x *FleTme) Now() (r *PkgFn) {
	r = x.PkgFn("Now")
	r.OutPrm(x)
	r.Add("return Time(time.Now())")
	return r
}
func (x *FleTme) WeekdayCnt() (r *TypFn) {
	r = x.TypFn("WeekdayCnt")
	r.InPrm(x, "a")
	r.OutPrm(_sys.Bsc.Unt, "r")
	r.Add("min, max := x.MinMax(a)")
	r.Add("for cur := min; cur.Lss(max); cur = cur.Add(Day) {")
	r.Add("weekday := cur.Time().Weekday()")
	r.Add("if weekday != 6 && weekday != 7 {")
	r.Add("r++")
	r.Add("}")
	r.Add("}")
	r.Add("return r")
	return r
}

func (x *FleTme) PkgDuration() (r *PkgFn) {
	r = x.PkgFn("Duration")
	r.InPrm(Duration, "v")
	r.OutPrm(x)
	r.Add("return Tme(v.Seconds())")
	return r
}
func (x *FleTme) MemDuration() (r *TypFn) {
	r = x.TypFn("Duration")
	r.OutPrm(Duration)
	r.Add("return time.Duration(x) * time.Second")
	return r
}
func (x *FleTme) DurWrt() (r *TypFn) {
	x.Import("fmt")
	r = x.TypFn("DurWrt")
	r.InPrm(BuilderPtr, "b")
	r.Add("if x == 0 {")
	r.Add("b.WriteString(\"0s\")")
	r.Add("return")
	r.Add("}")
	r.Add("if x < 0 {")
	r.Add("b.WriteRune('-')")
	r.Add("x = -x")
	r.Add("}")
	r.Add("if x >= Week {")
	r.Add("v := x / Week")
	r.Add("b.WriteString(fmt.Sprintf(\"%vw\", int32(v)))")
	r.Add("x -= v * Week")
	r.Add("if x == 0 {")
	r.Add("return")
	r.Add("}")
	r.Add("}")
	r.Add("if x >= Day {")
	r.Add("v := x / Day")
	r.Add("b.WriteString(fmt.Sprintf(\"%vd\", int32(v)))")
	r.Add("x -= v * Day")
	r.Add("if x == 0 {")
	r.Add("return")
	r.Add("}")
	r.Add("}")
	r.Add("if x >= Hour {")
	r.Add("v := x / Hour")
	r.Add("b.WriteString(fmt.Sprintf(\"%vh\", int32(v)))")
	r.Add("x -= v * Hour")
	r.Add("if x == 0 {")
	r.Add("return")
	r.Add("}")
	r.Add("}")
	r.Add("if x >= Minute {")
	r.Add("v := x / Minute")
	r.Add("b.WriteString(fmt.Sprintf(\"%vm\", int32(v)))")
	r.Add("x -= v * Minute")
	r.Add("if x == 0 {")
	r.Add("return")
	r.Add("}")
	r.Add("}")
	r.Add("if x != 0 {")
	r.Add("b.WriteString(fmt.Sprintf(\"%vs\", int32(x)))")
	r.Add("}")
	return r
}
func (x *FleTme) DurString() (r *TypFn) {
	r = x.TypFn("DurString")
	r.OutPrm(String)
	r.Addf("b := %v{}", BuilderPtr.Adr(x))
	r.Add("x.DurWrt(b)")
	r.Add("return b.String()")
	return r
}

func (x *FleTmes) InitTypFn() {
	x.Times()
	x.Bnd()
	x.WeekdayCnt()
}

// func (x *FleTmes) InitVals(bse *TypBse) {
// 	bse.PrmLit = "[2s 4s 8s]" // for InrvlFbr
// }
func (x *FleTmes) Times() (r *TypFn) {
	x.Import("time")
	r = x.TypFn("Times") // for plotting
	r.OutPrmSlice(Time, "r")
	r.Add("r = make([]time.Time, len(*x))")
	r.Add("for n, t := range *x {")
	r.Add("r[n] = t.Time()")
	r.Add("}")
	r.Add("return r")
	return r
}
func (x *FleTmes) Bnd() (r *TypFn) {
	r = x.TypFn(k.Bnd)
	r.InPrm(_sys.Bsc.TmeRng, "rng")
	r.OutPrm(_sys.Bsc.Bnd, "r")
	r.Add("if len(*x) == 0 {")
	r.Add("return r")
	r.Add("}")
	r.Add("rng = rng.Ensure()")
	r.Add("r.Idx = x.SrchIdx(rng.Min, true)")
	r.Add("r.Lim = x.SrchIdx(rng.Max, true)")
	r.Add("if r.Idx >= x.Cnt() {")
	r.Addf("return %v{}", _sys.Bsc.Bnd.Ref(x))
	r.Add("}")
	r.Add("if r.Lim > x.Cnt() {")
	r.Add("r.Lim = x.Cnt()")
	r.Add("}")
	r.Add("return r")
	return r
}
func (x *FleTmes) WeekdayCnt() (r *TypFn) {
	r = x.TypFn("WeekdayCnt")
	r.OutPrm(_sys.Bsc.Unt, "r")
	r.Add("if len(*x) == 0 {")
	r.Add("return 0")
	r.Add("}")
	r.Add("return x.Fst().WeekdayCnt((*x)[len(*x)-1])")
	return r
}
