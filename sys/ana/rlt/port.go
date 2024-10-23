package rlt

// func (x *Port) CalcPrfm() { // for testing
// 	// MAINLY, USED FOR VALIDATING HST
// 	// NOT USED TO CALCULATE ACTUAL RLT PERFORMANCE
// 	if x.Port.Trds.Cnt() == 0 || x.Stgys.Cnt() == 0 {
// 		return
// 	}
// 	losLimMax, durLimMax := flt.Min, tme.Min
// 	for _, stgy := range *x.Stgys {
// 		switch v := stgy.(type) {
// 		case *StgyLong:
// 			losLimMax = losLimMax.Max(v.LosLim)
// 			durLimMax = durLimMax.Max(v.DurLim)
// 		case *StgyShrt:
// 			losLimMax = losLimMax.Max(v.LosLim)
// 			durLimMax = durLimMax.Max(v.DurLim)
// 		default:
// 			panic("Stgy no implemented")
// 		}
// 	}
// 	x.Port.CalcPrfm(x.TrdDayCnt(), x.Stgys.Cnt(), losLimMax, durLimMax, str.Str(x.GenPth()))
// }

// func (x *Port) GenPth() string {
// 	var b strings.Builder
// 	b.WriteString("rlt.newPort(")
// 	for n, stgy := range *x.Stgys {
// 		if n != 0 {
// 			b.WriteString(" ")
// 		}
// 		b.WriteString(stgy.String())
// 	}
// 	b.WriteString(")")
// 	return b.String()
// }
