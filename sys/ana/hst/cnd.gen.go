package hst

import (
	"strings"
	"sys"
	"sys/ana"
	"sys/bsc/bnd"
	"sys/bsc/bnds"
	"sys/bsc/bol"
	"sys/bsc/flt"
	"sys/bsc/flts"
	"sys/bsc/str"
	"sys/bsc/tme"
	"sys/bsc/tmes"
	"sys/bsc/unt"
	"sys/err"
)

type (
	Cnd interface {
		Name() str.Str
		PrmWrt(b *strings.Builder)
		Prm() string
		StrWrt(b *strings.Builder)
		String() string
		Bse() *CndBse
		And(a Cnd) Cnd
		Seq(dur tme.Tme, a Cnd) Cnd
		Stgy(isLong bol.Bol, prfLim, losLim flt.Flt, durLim tme.Tme, minPnlPct flt.Flt, instr Instr, ftrStms *Stms, clss ...Cnd) Stgy
	}
	CndBse struct {
		Slf  Cnd
		Tmes *tmes.Tmes
	}
	CndSeg struct {
		bnd.Bnd
		Tmes *tmes.Tmes
		Out  *tmes.Tmes
	}
	CndScp struct {
		Idx uint32
		Arr []Cnd
	}
	CndSclEql struct {
		CndBse
		Stm Stm
		Scl flt.Flt
	}
	CndSclEqlSeg struct {
		CndSeg
		Vals *flts.Flts
		Scl  flt.Flt
	}
	CndSclNeq struct {
		CndBse
		Stm Stm
		Scl flt.Flt
	}
	CndSclNeqSeg struct {
		CndSeg
		Vals *flts.Flts
		Scl  flt.Flt
	}
	CndSclLss struct {
		CndBse
		Stm Stm
		Scl flt.Flt
	}
	CndSclLssSeg struct {
		CndSeg
		Vals *flts.Flts
		Scl  flt.Flt
	}
	CndSclGtr struct {
		CndBse
		Stm Stm
		Scl flt.Flt
	}
	CndSclGtrSeg struct {
		CndSeg
		Vals *flts.Flts
		Scl  flt.Flt
	}
	CndSclLeq struct {
		CndBse
		Stm Stm
		Scl flt.Flt
	}
	CndSclLeqSeg struct {
		CndSeg
		Vals *flts.Flts
		Scl  flt.Flt
	}
	CndSclGeq struct {
		CndBse
		Stm Stm
		Scl flt.Flt
	}
	CndSclGeqSeg struct {
		CndSeg
		Vals *flts.Flts
		Scl  flt.Flt
	}
	CndInrEql struct {
		CndBse
		Stm Stm
		Off unt.Unt
	}
	CndInrEqlSeg struct {
		CndSeg
		Vals *flts.Flts
		Off  unt.Unt
	}
	CndInrNeq struct {
		CndBse
		Stm Stm
		Off unt.Unt
	}
	CndInrNeqSeg struct {
		CndSeg
		Vals *flts.Flts
		Off  unt.Unt
	}
	CndInrLss struct {
		CndBse
		Stm Stm
		Off unt.Unt
	}
	CndInrLssSeg struct {
		CndSeg
		Vals *flts.Flts
		Off  unt.Unt
	}
	CndInrGtr struct {
		CndBse
		Stm Stm
		Off unt.Unt
	}
	CndInrGtrSeg struct {
		CndSeg
		Vals *flts.Flts
		Off  unt.Unt
	}
	CndInrLeq struct {
		CndBse
		Stm Stm
		Off unt.Unt
	}
	CndInrLeqSeg struct {
		CndSeg
		Vals *flts.Flts
		Off  unt.Unt
	}
	CndInrGeq struct {
		CndBse
		Stm Stm
		Off unt.Unt
	}
	CndInrGeqSeg struct {
		CndSeg
		Vals *flts.Flts
		Off  unt.Unt
	}
	CndOtrEql struct {
		CndBse
		Stm Stm
		Off unt.Unt
		A   Stm
	}
	CndOtrEqlSeg struct {
		CndSeg
		Vals  *flts.Flts
		Off   unt.Unt
		ValsA *flts.Flts
	}
	CndOtrNeq struct {
		CndBse
		Stm Stm
		Off unt.Unt
		A   Stm
	}
	CndOtrNeqSeg struct {
		CndSeg
		Vals  *flts.Flts
		Off   unt.Unt
		ValsA *flts.Flts
	}
	CndOtrLss struct {
		CndBse
		Stm Stm
		Off unt.Unt
		A   Stm
	}
	CndOtrLssSeg struct {
		CndSeg
		Vals  *flts.Flts
		Off   unt.Unt
		ValsA *flts.Flts
	}
	CndOtrGtr struct {
		CndBse
		Stm Stm
		Off unt.Unt
		A   Stm
	}
	CndOtrGtrSeg struct {
		CndSeg
		Vals  *flts.Flts
		Off   unt.Unt
		ValsA *flts.Flts
	}
	CndOtrLeq struct {
		CndBse
		Stm Stm
		Off unt.Unt
		A   Stm
	}
	CndOtrLeqSeg struct {
		CndSeg
		Vals  *flts.Flts
		Off   unt.Unt
		ValsA *flts.Flts
	}
	CndOtrGeq struct {
		CndBse
		Stm Stm
		Off unt.Unt
		A   Stm
	}
	CndOtrGeqSeg struct {
		CndSeg
		Vals  *flts.Flts
		Off   unt.Unt
		ValsA *flts.Flts
	}
	CndCnd1And struct {
		CndBse
		Cnd Cnd
		A   Cnd
	}
	CndCnd1AndSeg struct {
		CndSeg
		TmesA *tmes.Tmes
	}
	CndCnd2Seq struct {
		CndBse
		Cnd Cnd
		Dur tme.Tme
		A   Cnd
	}
	CndCnd2SeqSeg struct {
		CndSeg
		Dur   tme.Tme
		TmesA *tmes.Tmes
	}
)

func (x *CndSclEql) Name() str.Str             { return str.Str("SclEql") }
func (x *CndSclEql) PrmWrt(b *strings.Builder) { x.Scl.StrWrt(b) }
func (x *CndSclEql) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *CndSclEql) StrWrt(b *strings.Builder) {
	x.Stm.StrWrt(b)
	b.WriteString(".sclEql(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *CndSclEql) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *CndSclEqlSeg) Act() {
	for n := x.Idx; n < x.Lim; n++ {
		if x.Vals.At(n).Eql(x.Scl) {
			*x.Out = append(*x.Out, (*x.Tmes)[n])
		}
	}
}
func (x *CndSclNeq) Name() str.Str             { return str.Str("SclNeq") }
func (x *CndSclNeq) PrmWrt(b *strings.Builder) { x.Scl.StrWrt(b) }
func (x *CndSclNeq) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *CndSclNeq) StrWrt(b *strings.Builder) {
	x.Stm.StrWrt(b)
	b.WriteString(".sclNeq(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *CndSclNeq) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *CndSclNeqSeg) Act() {
	for n := x.Idx; n < x.Lim; n++ {
		if x.Vals.At(n).Neq(x.Scl) {
			*x.Out = append(*x.Out, (*x.Tmes)[n])
		}
	}
}
func (x *CndSclLss) Name() str.Str             { return str.Str("SclLss") }
func (x *CndSclLss) PrmWrt(b *strings.Builder) { x.Scl.StrWrt(b) }
func (x *CndSclLss) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *CndSclLss) StrWrt(b *strings.Builder) {
	x.Stm.StrWrt(b)
	b.WriteString(".sclLss(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *CndSclLss) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *CndSclLssSeg) Act() {
	for n := x.Idx; n < x.Lim; n++ {
		if x.Vals.At(n).Lss(x.Scl) {
			*x.Out = append(*x.Out, (*x.Tmes)[n])
		}
	}
}
func (x *CndSclGtr) Name() str.Str             { return str.Str("SclGtr") }
func (x *CndSclGtr) PrmWrt(b *strings.Builder) { x.Scl.StrWrt(b) }
func (x *CndSclGtr) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *CndSclGtr) StrWrt(b *strings.Builder) {
	x.Stm.StrWrt(b)
	b.WriteString(".sclGtr(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *CndSclGtr) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *CndSclGtrSeg) Act() {
	for n := x.Idx; n < x.Lim; n++ {
		if x.Vals.At(n).Gtr(x.Scl) {
			*x.Out = append(*x.Out, (*x.Tmes)[n])
		}
	}
}
func (x *CndSclLeq) Name() str.Str             { return str.Str("SclLeq") }
func (x *CndSclLeq) PrmWrt(b *strings.Builder) { x.Scl.StrWrt(b) }
func (x *CndSclLeq) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *CndSclLeq) StrWrt(b *strings.Builder) {
	x.Stm.StrWrt(b)
	b.WriteString(".sclLeq(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *CndSclLeq) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *CndSclLeqSeg) Act() {
	for n := x.Idx; n < x.Lim; n++ {
		if x.Vals.At(n).Leq(x.Scl) {
			*x.Out = append(*x.Out, (*x.Tmes)[n])
		}
	}
}
func (x *CndSclGeq) Name() str.Str             { return str.Str("SclGeq") }
func (x *CndSclGeq) PrmWrt(b *strings.Builder) { x.Scl.StrWrt(b) }
func (x *CndSclGeq) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *CndSclGeq) StrWrt(b *strings.Builder) {
	x.Stm.StrWrt(b)
	b.WriteString(".sclGeq(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *CndSclGeq) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *CndSclGeqSeg) Act() {
	for n := x.Idx; n < x.Lim; n++ {
		if x.Vals.At(n).Geq(x.Scl) {
			*x.Out = append(*x.Out, (*x.Tmes)[n])
		}
	}
}
func (x *CndInrEql) Name() str.Str             { return str.Str("InrEql") }
func (x *CndInrEql) PrmWrt(b *strings.Builder) { x.Off.StrWrt(b) }
func (x *CndInrEql) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *CndInrEql) StrWrt(b *strings.Builder) {
	x.Stm.StrWrt(b)
	b.WriteString(".inrEql(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *CndInrEql) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *CndInrEqlSeg) Act() {
	for n := x.Idx; n < x.Lim; n++ {
		if x.Vals.At(n + x.Off).Eql((*x.Vals)[n]) {
			*x.Out = append(*x.Out, (*x.Tmes)[n+x.Off])
		}
	}
}
func (x *CndInrNeq) Name() str.Str             { return str.Str("InrNeq") }
func (x *CndInrNeq) PrmWrt(b *strings.Builder) { x.Off.StrWrt(b) }
func (x *CndInrNeq) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *CndInrNeq) StrWrt(b *strings.Builder) {
	x.Stm.StrWrt(b)
	b.WriteString(".inrNeq(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *CndInrNeq) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *CndInrNeqSeg) Act() {
	for n := x.Idx; n < x.Lim; n++ {
		if x.Vals.At(n + x.Off).Neq((*x.Vals)[n]) {
			*x.Out = append(*x.Out, (*x.Tmes)[n+x.Off])
		}
	}
}
func (x *CndInrLss) Name() str.Str             { return str.Str("InrLss") }
func (x *CndInrLss) PrmWrt(b *strings.Builder) { x.Off.StrWrt(b) }
func (x *CndInrLss) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *CndInrLss) StrWrt(b *strings.Builder) {
	x.Stm.StrWrt(b)
	b.WriteString(".inrLss(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *CndInrLss) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *CndInrLssSeg) Act() {
	for n := x.Idx; n < x.Lim; n++ {
		if x.Vals.At(n + x.Off).Lss((*x.Vals)[n]) {
			*x.Out = append(*x.Out, (*x.Tmes)[n+x.Off])
		}
	}
}
func (x *CndInrGtr) Name() str.Str             { return str.Str("InrGtr") }
func (x *CndInrGtr) PrmWrt(b *strings.Builder) { x.Off.StrWrt(b) }
func (x *CndInrGtr) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *CndInrGtr) StrWrt(b *strings.Builder) {
	x.Stm.StrWrt(b)
	b.WriteString(".inrGtr(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *CndInrGtr) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *CndInrGtrSeg) Act() {
	for n := x.Idx; n < x.Lim; n++ {
		if x.Vals.At(n + x.Off).Gtr((*x.Vals)[n]) {
			*x.Out = append(*x.Out, (*x.Tmes)[n+x.Off])
		}
	}
}
func (x *CndInrLeq) Name() str.Str             { return str.Str("InrLeq") }
func (x *CndInrLeq) PrmWrt(b *strings.Builder) { x.Off.StrWrt(b) }
func (x *CndInrLeq) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *CndInrLeq) StrWrt(b *strings.Builder) {
	x.Stm.StrWrt(b)
	b.WriteString(".inrLeq(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *CndInrLeq) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *CndInrLeqSeg) Act() {
	for n := x.Idx; n < x.Lim; n++ {
		if x.Vals.At(n + x.Off).Leq((*x.Vals)[n]) {
			*x.Out = append(*x.Out, (*x.Tmes)[n+x.Off])
		}
	}
}
func (x *CndInrGeq) Name() str.Str             { return str.Str("InrGeq") }
func (x *CndInrGeq) PrmWrt(b *strings.Builder) { x.Off.StrWrt(b) }
func (x *CndInrGeq) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *CndInrGeq) StrWrt(b *strings.Builder) {
	x.Stm.StrWrt(b)
	b.WriteString(".inrGeq(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *CndInrGeq) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *CndInrGeqSeg) Act() {
	for n := x.Idx; n < x.Lim; n++ {
		if x.Vals.At(n + x.Off).Geq((*x.Vals)[n]) {
			*x.Out = append(*x.Out, (*x.Tmes)[n+x.Off])
		}
	}
}
func (x *CndOtrEql) Name() str.Str { return str.Str("OtrEql") }
func (x *CndOtrEql) PrmWrt(b *strings.Builder) {
	x.Off.StrWrt(b)
	b.WriteRune(' ')
	x.A.StrWrt(b)
}
func (x *CndOtrEql) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *CndOtrEql) StrWrt(b *strings.Builder) {
	x.Stm.StrWrt(b)
	b.WriteString(".otrEql(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *CndOtrEql) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *CndOtrEqlSeg) Act() {
	for n := x.Idx; n < x.Lim; n++ {
		if x.Vals.At(n).Eql((*x.ValsA)[n+x.Off]) {
			*x.Out = append(*x.Out, (*x.Tmes)[n+x.Off])
		}
	}
}
func (x *CndOtrNeq) Name() str.Str { return str.Str("OtrNeq") }
func (x *CndOtrNeq) PrmWrt(b *strings.Builder) {
	x.Off.StrWrt(b)
	b.WriteRune(' ')
	x.A.StrWrt(b)
}
func (x *CndOtrNeq) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *CndOtrNeq) StrWrt(b *strings.Builder) {
	x.Stm.StrWrt(b)
	b.WriteString(".otrNeq(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *CndOtrNeq) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *CndOtrNeqSeg) Act() {
	for n := x.Idx; n < x.Lim; n++ {
		if x.Vals.At(n).Neq((*x.ValsA)[n+x.Off]) {
			*x.Out = append(*x.Out, (*x.Tmes)[n+x.Off])
		}
	}
}
func (x *CndOtrLss) Name() str.Str { return str.Str("OtrLss") }
func (x *CndOtrLss) PrmWrt(b *strings.Builder) {
	x.Off.StrWrt(b)
	b.WriteRune(' ')
	x.A.StrWrt(b)
}
func (x *CndOtrLss) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *CndOtrLss) StrWrt(b *strings.Builder) {
	x.Stm.StrWrt(b)
	b.WriteString(".otrLss(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *CndOtrLss) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *CndOtrLssSeg) Act() {
	for n := x.Idx; n < x.Lim; n++ {
		if x.Vals.At(n).Lss((*x.ValsA)[n+x.Off]) {
			*x.Out = append(*x.Out, (*x.Tmes)[n+x.Off])
		}
	}
}
func (x *CndOtrGtr) Name() str.Str { return str.Str("OtrGtr") }
func (x *CndOtrGtr) PrmWrt(b *strings.Builder) {
	x.Off.StrWrt(b)
	b.WriteRune(' ')
	x.A.StrWrt(b)
}
func (x *CndOtrGtr) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *CndOtrGtr) StrWrt(b *strings.Builder) {
	x.Stm.StrWrt(b)
	b.WriteString(".otrGtr(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *CndOtrGtr) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *CndOtrGtrSeg) Act() {
	for n := x.Idx; n < x.Lim; n++ {
		if x.Vals.At(n).Gtr((*x.ValsA)[n+x.Off]) {
			*x.Out = append(*x.Out, (*x.Tmes)[n+x.Off])
		}
	}
}
func (x *CndOtrLeq) Name() str.Str { return str.Str("OtrLeq") }
func (x *CndOtrLeq) PrmWrt(b *strings.Builder) {
	x.Off.StrWrt(b)
	b.WriteRune(' ')
	x.A.StrWrt(b)
}
func (x *CndOtrLeq) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *CndOtrLeq) StrWrt(b *strings.Builder) {
	x.Stm.StrWrt(b)
	b.WriteString(".otrLeq(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *CndOtrLeq) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *CndOtrLeqSeg) Act() {
	for n := x.Idx; n < x.Lim; n++ {
		if x.Vals.At(n).Leq((*x.ValsA)[n+x.Off]) {
			*x.Out = append(*x.Out, (*x.Tmes)[n+x.Off])
		}
	}
}
func (x *CndOtrGeq) Name() str.Str { return str.Str("OtrGeq") }
func (x *CndOtrGeq) PrmWrt(b *strings.Builder) {
	x.Off.StrWrt(b)
	b.WriteRune(' ')
	x.A.StrWrt(b)
}
func (x *CndOtrGeq) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *CndOtrGeq) StrWrt(b *strings.Builder) {
	x.Stm.StrWrt(b)
	b.WriteString(".otrGeq(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *CndOtrGeq) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *CndOtrGeqSeg) Act() {
	for n := x.Idx; n < x.Lim; n++ {
		if x.Vals.At(n).Geq((*x.ValsA)[n+x.Off]) {
			*x.Out = append(*x.Out, (*x.Tmes)[n+x.Off])
		}
	}
}
func (x *CndBse) Bse() *CndBse { return x }
func (x *CndBse) And(a Cnd) Cnd {
	r := &CndCnd1And{}
	r.Slf = r
	r.Cnd = x.Slf
	r.A = a
	if ana.Cfg.Trc.IsHstCnd() {
		sys.Logf("%p hst.CndCnd1And(%v)", r, r.Prm())
	}
	if x.Tmes.Cnt() == 0 || r.A.Bse().Tmes.Cnt() == 0 {
		return r
	}
	r.Tmes = tmes.New()
	segBnds, acts := bnds.Segs(x.Tmes.Cnt())
	for n, segBnd := range *segBnds {
		seg := &CndCnd1AndSeg{}
		seg.Bnd = segBnd
		seg.Tmes = x.Tmes
		seg.TmesA = r.A.Bse().Tmes
		seg.Out = tmes.New()
		acts[n] = seg
	}
	sys.Run().Pll(acts...)
	for _, act := range acts {
		r.Tmes.Mrg(act.(*CndCnd1AndSeg).Out)
	}
	return r
}
func (x *CndCnd1And) Name() str.Str             { return str.Str("And") }
func (x *CndCnd1And) PrmWrt(b *strings.Builder) { x.A.StrWrt(b) }
func (x *CndCnd1And) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *CndCnd1And) StrWrt(b *strings.Builder) {
	x.Cnd.StrWrt(b)
	b.WriteString(".and(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *CndCnd1And) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *CndCnd1AndSeg) Act() {
	for n := x.Idx; n < x.Lim; n++ {
		if x.TmesA.Has((*x.Tmes)[n]) {
			*x.Out = append(*x.Out, (*x.Tmes)[n])
		}
	}
}
func (x *CndBse) Seq(dur tme.Tme, a Cnd) Cnd {
	r := &CndCnd2Seq{}
	r.Slf = r
	r.Cnd = x.Slf
	r.Dur = dur
	r.A = a
	if ana.Cfg.Trc.IsHstCnd() {
		sys.Logf("%p hst.CndCnd2Seq(%v)", r, r.Prm())
	}
	if x.Tmes.Cnt() == 0 || r.A.Bse().Tmes.Cnt() == 0 {
		return r
	}
	r.Tmes = tmes.New()
	segBnds, acts := bnds.Segs(x.Tmes.Cnt())
	for n, segBnd := range *segBnds {
		seg := &CndCnd2SeqSeg{}
		seg.Bnd = segBnd
		seg.Tmes = x.Tmes
		seg.TmesA = r.A.Bse().Tmes
		seg.Dur = r.Dur
		seg.Out = tmes.New()
		acts[n] = seg
	}
	sys.Run().Pll(acts...)
	for _, act := range acts {
		r.Tmes.Mrg(act.(*CndCnd2SeqSeg).Out)
	}
	return r
}
func (x *CndCnd2Seq) Name() str.Str { return str.Str("Seq") }
func (x *CndCnd2Seq) PrmWrt(b *strings.Builder) {
	x.Dur.StrWrt(b)
	b.WriteRune(' ')
	x.A.StrWrt(b)
}
func (x *CndCnd2Seq) Prm() string {
	b := &strings.Builder{}
	x.PrmWrt(b)
	return b.String()
}
func (x *CndCnd2Seq) StrWrt(b *strings.Builder) {
	x.Cnd.StrWrt(b)
	b.WriteString(".seq(")
	x.PrmWrt(b)
	b.WriteRune(')')
}
func (x *CndCnd2Seq) String() string {
	b := &strings.Builder{}
	x.StrWrt(b)
	return b.String()
}
func (x *CndCnd2SeqSeg) Act() {
	for n := x.Idx; n < x.Lim; n++ {
		if x.TmesA.Has((*x.Tmes)[n] + x.Dur) {
			*x.Out = append(*x.Out, (*x.Tmes)[n]+x.Dur)
		}
	}
}
func (x *CndBse) Stgy(isLong bol.Bol, prfLim, losLim flt.Flt, durLim tme.Tme, minPnlPct flt.Flt, instr Instr, ftrStms *Stms, clss ...Cnd) Stgy {
	r := &StgyStgy{}
	r.Slf = r
	r.Cnd = x.Slf
	r.IsLong = isLong
	r.PrfLim = prfLim
	r.LosLim = losLim
	r.DurLim = durLim
	r.MinPnlPct = minPnlPct
	r.Instr = instr
	r.FtrStms = ftrStms
	r.Clss = clss
	if ana.Cfg.Trc.IsHstStgy() {
		sys.Logf("%p hst.StgyStgy(%v)", r, r.Prm())
	}
	if ftrStms.Cnt() == 0 {
		err.Panic("ftrStms is empty")
	}
	if x.Tmes.Cnt() == 0 {
		return r
	}
	return r
}
