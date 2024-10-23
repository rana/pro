package hst

// import (
// 	"regexp"
// 	"strings"
// 	"sys"
// 	"sys/ana"
// )

// func (x *StgyRlngLongRlng) Calc() Stgy {
// 	// construct prcp string used by each bckRng tuning
// 	// in which the instr rng is replaced/inserted
// 	b := &strings.Builder{}
// 	x.Prcp.StrWrt(b)
// 	opnCndTxt := x.Cnd.String()
// 	re := regexp.MustCompile("0s-0s")
// 	var stgy Stgy
// 	var prfm Prfm
// 	for n, bckRng := range *x.Rngs {
// 		if ana.Cfg.Trc.IsHstStgy() {
// 			// sys.Logf("hst.StgyRlngLongRlng(%v) %p %v of %v %v", x.Prm(), x, n, len(*x.Rngs), bckRng)
// 			sys.Logf("%p %v of %v %v", x, n, len(*x.Rngs), bckRng)
// 		}
// 		// prcp: curPrcp is for current rng
// 		curPrcpTxt := re.ReplaceAllString(b.String(), bckRng.String())
// 		curPrcp := sys.Actr().RunIfc(curPrcpTxt)[0].(*Prcp)
// 		if n == 0 { // tune fst
// 			// opnCnd: curCnd is for current rng
// 			curOpnCndTxt := re.ReplaceAllString(opnCndTxt, bckRng.String())
// 			sys.Log("cnd txt", curOpnCndTxt)
// 			vs := sys.Actr().RunIfc(curOpnCndTxt)
// 			curOpnCnd := vs[len(vs)-1].(Cnd)
// 			stgy = curOpnCnd.Long(x.PrfLim, x.LosLim, x.DurLim, x.Instr, x.Clss...)
// 			splt := stgy.Port().Splt(x.SpltPnt)
// 			prfm = curPrcp.Splt(splt).TuneSacfTil(
// 				x.SpltPnt, x.StmCntLim,
// 				x.TrimItrLim, x.TrimMin, x.TrimForgiveLim,
// 			)
// 			if ana.Cfg.Trc.IsHstStgy() {
// 				sys.Logf("%p %v", x, prfm.Ana())
// 			}
// 		} else { // tune consecutive, later, non-fst rngs
// 			splt := prfm.Port().Splt(x.SpltPnt)
// 			prfm = curPrcp.Splt(splt).TuneSacfTil(
// 				x.SpltPnt, x.StmCntLim,
// 				x.TrimItrLim, x.TrimMin, x.TrimForgiveLim,
// 			)
// 			if ana.Cfg.Trc.IsHstStgy() {
// 				sys.Logf("%p %v", x, prfm)
// 			}
// 		}
// 	}
// 	if prfm != nil {
// 		return prfm.Port().Stgys().At(0)
// 	}
// 	// return default long stgy in case error in prms ie, no rngs or something else
// 	return x.Cnd.Long(x.PrfLim, x.LosLim, x.DurLim, x.Instr, x.Clss...)
// }
