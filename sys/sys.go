package sys

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"os"
	"os/exec"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"sys/k"
	"time"
	"unicode"
	"unicode/utf8"

	"cloud.google.com/go/bigquery"
)

var (
	_sys *Sys
	_log *log.Logger
)

type (
	// Act interface {
	// 	Act()
	// }
	Sys struct {
		Runr IRunr
		Cldr ICldr
		Dskr IDskr
		Actr IActr
		Lrnr ILrnr
		id   uint32
	}
	IRunr interface {
		Seq(acts ...Act)
		Pll(acts ...Act)
	}
	ICldr interface {
		Query(txt string) *bigquery.RowIterator
		GetTable(v interface{}) (r *bigquery.Table)
		Cfg() (projectName, datasetName string)
		Mu() *sync.Mutex
	}
	IDskr interface {
		SavInstrStm(key []byte, val []byte)
		LoadInstrStm(key []byte) []byte
		DelInstrStm(key []byte)

		SavInstrDetail(key []byte, val []byte)
		LoadInstrDetail(key []byte) []byte
		DelInstrDetail(key []byte)

		SavMl(key []byte, val []byte)
		LoadMl(key []byte) []byte
		DelMl(key []byte)
	}
	IActr interface {
		RunIfcf(format string, args ...interface{}) []interface{}
		RunIfc(txt string) []interface{}
		RunRlt(txt string) []interface{}
		RunHst(txt string) []interface{}
	}
	ILrnr interface {
		Fit(key string, ftrNames []string, ftrs [][]float32, lbls []float32) (bits []byte)
		Predict(key string, x []float32) (y float32)
		SaveNetToDsk(key string, net []byte)
		LoadNetFromDsk(key string)
	}
)

func NewSys(runr IRunr, cldr ICldr, dskr IDskr, actr IActr, lrnr ILrnr) {
	_sys = &Sys{
		Runr: runr,
		Cldr: cldr,
		Dskr: dskr,
		Actr: actr,
		Lrnr: lrnr,
	}
}
func Run() IRunr     { return _sys.Runr }
func Cld() ICldr     { return _sys.Cldr }
func Dsk() IDskr     { return _sys.Dskr }
func Actr() IActr    { return _sys.Actr }
func Lrnr() ILrnr    { return _sys.Lrnr }
func HasDsk() bool   { return _sys.Dskr != nil && !reflect.ValueOf(_sys.Dskr).IsNil() }
func NextID() uint32 { return atomic.AddUint32(&_sys.id, 1) }

// TODO: UI?

func Log(vs ...interface{}) {
	if _log != nil {
		_log.Println(vs...)
	}
}
func Logf(tmpl string, args ...interface{}) {
	if _log != nil {
		_log.Printf(tmpl, args...)
	}
}

// func LogIfc(vs ...Ifc) {
// 	if _log != nil {
// 		is := make([]interface{}, len(vs))
// 		for n, ifc := range vs {
// 			is[n] = ifc //.Ifc()
// 		}
// 		_log.Println(is...)
// 	}
// }
// func LogIfcf(tmpl string, vs ...Ifc) {
// 	if _log != nil {
// 		is := make([]interface{}, len(vs))
// 		for n, ifc := range vs {
// 			is[n] = ifc //.Ifc()
// 		}
// 		_log.Printf(tmpl, is...)
// 	}
// }

func init() {
	// var logpath = "/home/rana/sys.log"
	// curDir, er := os.Getwd()
	// if er != nil {
	// 	panic(er)
	// }
	// logpath := filepath.Join(curDir, "sys.log")

	var logpath = "/home/rana/log/sys.log"
	var file, err1 = os.Create(logpath)
	if err1 != nil {
		panic(err1)
	}
	// _log = log.New(file, "", log.LstdFlags|log.Lshortfile)
	_log = log.New(file, "", 0)
	_log.Println(logpath)
	_log.Println(time.Now())
}

func Vs(vs ...string) []string           { return vs }
func Is(vs ...interface{}) []interface{} { return vs }

// func Is(vs ...string) (is []interface{}) {
// 	is = make([]interface{}, len(vs))
// 	for n, v := range vs {
// 		is[n] = v
// 	}
// 	return is
// }
func VsStruct(name string, vs ...string) (r []string) {
	for _, v := range vs {
		r = append(r, fmt.Sprintf("%v{%v}", name, v))
	}
	return r
}
func VsCtor(name string, vs ...string) (r []string) {
	for _, v := range vs {
		r = append(r, fmt.Sprintf("%v(%v)", name, v))
	}
	return r
}

// Trim removes whitespace from the front and back.
func Trim(s string) string {
	f := 0 // front
	for ; f < len(s); f++ {
		ch, _ := utf8.DecodeRuneInString(s[f:])
		if !unicode.IsSpace(ch) {
			break
		}
	}
	b := len(s) // back
	for ; b > -1; b-- {
		ch, _ := utf8.DecodeLastRuneInString(s[:b])
		if !unicode.IsSpace(ch) {
			break
		}
	}
	return s[f:b]
}

// TrimFront removes whitespace from the front.
func TrimFront(s string) string {
	f := 0
	for ; f < len(s); f++ {
		ch, _ := utf8.DecodeRuneInString(s[f:])
		if !unicode.IsSpace(ch) {
			break
		}
	}
	return s[f:]
}

// TrimBack removes whitespace from the back.
func TrimBack(s string) string {
	b := len(s)
	for ; b > -1; b-- {
		ch, _ := utf8.DecodeLastRuneInString(s[:b])
		if !unicode.IsSpace(ch) {
			break
		}
	}
	return s[:b]
}

// Camel returns a strign in camel case e.g., 'camelCase'.
func Camel(s string) string {
	if len(s) == 0 {
		return s
	}
	if s[0] > 'Z' {
		return s
	}
	b := make([]byte, len(s))
	b[0] = s[0] + 32
	b2 := b[1:]
	copy(b2, s[1:])
	return string(b)
}

// StackTrace returns the calling go routine stack trace.
func StackTrace() string {
	buf := make([]byte, 1<<16)
	n := runtime.Stack(buf, false)
	return string(buf[:n])
}

func DayOfWeek(t time.Time, w time.Weekday) time.Time {
	year, month, day := t.Date()
	t = time.Date(year, month, day, 0, 0, 0, 0, t.Location())
	return t.Add(time.Duration(-t.Weekday()+w) * time.Hour * 24)
}

func Uint64(a, b uint32) uint64 {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, a)
	binary.Write(buf, binary.LittleEndian, b)
	return binary.LittleEndian.Uint64(buf.Bytes())
}

func GoroutineID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}
func InstrTitle(name string) string {
	strs := strings.Split(name, "_")
	return fmt.Sprintf("%v%v", strings.Title(strs[0]), strings.Title(strs[1]))
}
func InstrCamel(name string) string { return Camel(InstrTitle(name)) }
func Test(s string) string          { return fmt.Sprintf("%v.test", s) }
func Plural(s string) string        { return fmt.Sprintf("%vs", s) }
func Plurals(vs ...string) (r []string) {
	r = make([]string, len(vs))
	for n, v := range vs {
		r[n] = Plural(v)
	}
	return r
}
func Singular(s string) string      { return s[:len(s)-1] }
func SingularTitle(s string) string { return strings.Title(s[:len(s)-1]) }
func NumTitle(num string) (r string) {
	r = strings.Replace(num, ".", "", 1)
	r = strings.Replace(r, "-", "Neg", 1)
	r = strings.Replace(r, "+", "Pos", 1)
	return r
}
func NumTitles(nums ...string) (r []string) {
	r = make([]string, len(nums))
	for n, num := range nums {
		r[n] = NumTitle(num)
	}
	return r
}
func Rx(name string) string  { return fmt.Sprintf("Rx%v", strings.Title(name)) }
func RxA(name string) string { return fmt.Sprintf("Rx%vA", strings.Title(name)) }
func RxB(name string) string { return fmt.Sprintf("Rx%vB", strings.Title(name)) }
func Tx(name string) string  { return fmt.Sprintf("Tx%v", strings.Title(name)) }
func CnjCamel(prefix, name string) string {
	return fmt.Sprintf("%v%v", Camel(prefix), strings.Title(name))
}
func CnjTitle(prefix, name string) string {
	return fmt.Sprintf("%v%v", strings.Title(prefix), strings.Title(name))
}
func Cnj(prefix, name string) string {
	return fmt.Sprintf("%v%v", prefix, strings.Title(name))
}
func Cnjss(prefix, name string) string {
	return fmt.Sprintf("%v%vs", prefix, strings.Title(name))
}
func Decnj(prefix, name string) string { return name[len(prefix):] }

func CnjRte(name string) string { return Cnj(k.Rte, name) }
func CnjUna(name string) string { return Cnj(k.Una, name) }
func CnjScl(name string) string { return Cnj(k.Scl, name) }
func CnjSel(name string) string { return Cnj(k.Sel, name) }
func CnjAgg(name string) string { return Cnj(k.Agg, name) }
func CnjInr(name string) string { return Cnj(k.Inr, name) }
func CnjOtr(name string) string { return Cnj(k.Otr, name) }
func CnjRel(name string) string { return Cnj(k.Rel, name) }
func CnjCnt(name string) string { return Cnj(k.Cnt, name) }
func CnjCnd(name string) string { return Cnj(k.Cnd, name) }

func DecnjRte(name string) string { return Decnj(k.Rte, name) }
func DecnjUna(name string) string { return Decnj(k.Una, name) }
func DecnjScl(name string) string { return Decnj(k.Scl, name) }
func DecnjSel(name string) string { return Decnj(k.Sel, name) }
func DecnjAgg(name string) string { return Decnj(k.Agg, name) }
func DecnjInr(name string) string { return Decnj(k.Inr, name) }
func DecnjOtr(name string) string { return Decnj(k.Otr, name) }
func DecnjRel(name string) string { return Decnj(k.Rel, name) }
func DecnjCnt(name string) string { return Decnj(k.Cnt, name) }
func DecnjCnd(name string) string { return Decnj(k.Cnd, name) }

func CnjRtess(name string) string { return Cnjss(k.Rte, name) }
func CnjUnass(name string) string { return Cnjss(k.Una, name) }
func CnjSclss(name string) string { return Cnjss(k.Scl, name) }
func CnjSelss(name string) string { return Cnjss(k.Sel, name) }
func CnjAggss(name string) string { return Cnjss(k.Agg, name) }
func CnjInrss(name string) string { return Cnjss(k.Inr, name) }
func CnjOtrss(name string) string { return Cnjss(k.Otr, name) }
func CnjRelss(name string) string { return Cnjss(k.Rel, name) }
func CnjCntss(name string) string { return Cnjss(k.Cnt, name) }
func CnjCndss(name string) string { return Cnjss(k.Cnd, name) }

func Cnjs(cnj func(string) string, names ...string) (r []string) {
	for _, name := range names {
		r = append(r, cnj(name))
	}
	return r
}
func CnjUnas(names ...string) (r []string) { return Cnjs(CnjUna, names...) }
func CnjScls(names ...string) (r []string) { return Cnjs(CnjScl, names...) }
func CnjSels(names ...string) (r []string) { return Cnjs(CnjSel, names...) }
func CnjAggs(names ...string) (r []string) { return Cnjs(CnjAgg, names...) }
func CnjInrs(names ...string) (r []string) { return Cnjs(CnjInr, names...) }
func CnjOtrs(names ...string) (r []string) { return Cnjs(CnjOtr, names...) }
func CnjRels(names ...string) (r []string) { return Cnjs(CnjRel, names...) }
func CnjCnts(names ...string) (r []string) { return Cnjs(CnjCnt, names...) }
func CnjCnds(names ...string) (r []string) { return Cnjs(CnjCnd, names...) }

func Titles(ss ...string) string {
	b := &bytes.Buffer{}
	for _, s := range ss {
		if len(s) > 0 {
			ch, _ := utf8.DecodeRuneInString(s)
			if unicode.IsUpper(ch) {
				b.WriteString(s)
			} else {
				b.WriteRune(unicode.ToUpper(ch))
				if len(s) > 1 {
					b.WriteString(s[1:])
				}
			}
		}
	}
	return b.String()
}
func Prms(prms ...interface{}) string {
	if len(prms) == 0 {
		return ""
	}
	b := &bytes.Buffer{}
	for n, prm := range prms {
		if n != 0 {
			b.WriteRune(',')
		}
		b.WriteString(fmt.Sprintf("%v", prm))
	}
	return b.String()
}
func Join(ss ...string) string {
	b := &bytes.Buffer{}
	for _, s := range ss {
		b.WriteString(s)
	}
	return b.String()
}
func JoinPth(ss ...string) string {
	b := &bytes.Buffer{}
	for n, s := range ss {
		if n != 0 {
			b.WriteRune('.')
		}
		b.WriteString(s)
	}
	return b.String()
}
func GetInstr(pth string) (idx, lim int) {
	// TODO: A BETTER WAY?
	idx = strings.Index(pth, ".Oan()")
	if idx > -1 {
		idx = idx + 7
		lim := strings.Index(pth[idx:], ")")
		if lim < 0 {
			return -1, -1
		}
		return idx, idx + lim
	}
	return -1, -1
}
func CpyInstr(src, dst string) string {
	srcIdx, srcLim := GetInstr(src)
	if srcLim <= 0 {
		return ""
	}
	dstIdx, dstLim := GetInstr(dst)
	if dstLim <= 0 {
		return ""
	}
	b := &strings.Builder{}
	b.WriteString(dst[:dstIdx])
	b.WriteString(src[srcIdx:srcLim])
	b.WriteString(dst[dstLim:])
	return b.String()
}
func Idx(s, sub string, cnt int) (cur int) {
	if len(s) == 0 || len(sub) == 0 || cnt < 1 {
		return -1
	}
	cur, found := 0, 0
	for cur > -1 && found < cnt {
		cur = cur + strings.Index(s[cur:], sub)
		if cur > -1 {
			found++
			cur++
		}
	}
	if cur > 0 {
		cur--
	}
	return cur
}

func StrLens(vs ...[]string) (r []int) {
	r = make([]int, len(vs))
	for n, v := range vs {
		r[n] = len(v)
	}
	return r
}
func PermIdxsStrs(vs ...[]string) [][]int { return PermIdxs(StrLens(vs...)...) }
func PermIdxs(lims ...int) [][]int {
	idxss := make([][]int, len(lims))
	for n, lim := range lims {
		idxs := make([]int, lim)
		idxss[n] = idxs
		for i := 0; i < lim; i++ {
			idxs[i] = i
		}
	}
	return PermInts(idxss)
}
func PermInts(idxs [][]int) [][]int {
	toRet := [][]int{}
	if len(idxs) == 1 {
		for _, vid := range idxs[0] {
			toRet = append(toRet, []int{vid})
		}
		return toRet
	}
	t := PermInts(idxs[1:])
	for _, vid := range idxs[0] {
		for _, perm := range t {
			toRetAdd := append([]int{vid}, perm...)
			toRet = append(toRet, toRetAdd)
		}
	}
	return toRet
}
func Abcs() (r []rune) {
	r = make([]rune, 26)
	ch := 'a'
	for n := 0; n < len(r); n++ {
		r[n] = ch
		ch++
	}
	return r
}
func NavUrl(url string) {
	exec.Command("xdg-open", url).Start()
}
func NavUrlf(format string, args ...interface{}) {
	NavUrl(fmt.Sprintf(format, args...))
}

func OpnImg(filename string) {
	// exec.Command("eog", "--fullscreen", filename).Start()
	exec.Command("eog", "--fullscreen", "--new-instance", filename).Start()
}

func FnNameFull(depth ...int) string {
	skp := 1
	if len(depth) > 0 {
		skp += depth[0]
	}
	pc, _, _, _ := runtime.Caller(skp)
	fn := runtime.FuncForPC(pc)
	return fn.Name()
}
func FnName(depth ...int) string {
	skp := 1
	if len(depth) > 0 {
		skp += depth[0]
	}
	pc, _, _, _ := runtime.Caller(skp)
	fn := runtime.FuncForPC(pc)
	full := fn.Name()
	return full[strings.Index(full, ".")+1:]
}

func minMax(x, y float32) (float32, float32) {
	if x < y {
		return x, y
	}
	return y, x
}
func min(x, y float32) float32 {
	if x < y {
		return x
	}
	return y
}
func Max(x, y float32) float32 {
	if x > y {
		return x
	}
	return y
}
func abs(v float32) float32 {
	if v < 0 {
		return -v
	}
	return v
}
func ceil(v float32) float32 {
	return float32(int(v) + 1)
}
func dist(x1, y1, x2, y2 float32) float32 {
	return float32(math.Hypot(float64(x1)-float64(x2), float64(y1)-float64(y2)))
}
func slope(x1, y1, x2, y2 float32) float32 {
	return (y2 - y1) / (x2 - x1)
}

func MaxI(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func MinI(x, y int) int {
	if x < y {
		return x
	}
	return y
}
