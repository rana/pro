package hst

// import (
// 	"strings"
// 	"sys"
// 	"sys/bsc/flt"
// 	"sys/lng/pro/trm/prs"
// )

// type (
// 	ReduceAndCnd struct {
// 		Txt    string
// 		LstIdx int
// 		Prefix string
// 		AndCnd Cnd
// 		Val    flt.Flt
// 	}
// )

// func (x *Prcp) ReduceAndCnds() (finalCnds []Cnd) {
// 	if len(x.andCnds) < 2 {
// 		return x.andCnds
// 	}
// 	similarCnds := make(map[string][]*ReduceAndCnd)
// 	for m := 0; m < len(x.andCnds); m++ { // group by prefix
// 		andCnd := x.andCnds[m]
// 		v := &ReduceAndCnd{}
// 		v.Txt = andCnd.String()
// 		v.LstIdx = strings.LastIndex(v.Txt, "(")
// 		v.AndCnd = andCnd
// 		v.Prefix = v.Txt[:v.LstIdx]
// 		if arr, ok := similarCnds[v.Prefix]; ok {
// 			similarCnds[v.Prefix] = append(arr, v)
// 		} else {
// 			similarCnds[v.Prefix] = []*ReduceAndCnd{v}
// 		}
// 		sys.Log("ReduceAndCnds v.Txt", v.Txt)
// 	}
// 	for prefix, arr := range similarCnds {
// 		var idxFinal int
// 		if len(arr) > 1 {
// 			for _, itm := range arr { // prs vals
// 				itm.Val = prs.Flt(itm.Txt[itm.LstIdx+1 : len(itm.Txt)-1])
// 			}
// 			switch { // determine extreme val
// 			case strings.HasSuffix(prefix, "Leq"), strings.HasSuffix(prefix, "Lss"):
// 				max := flt.Min
// 				for n, itm := range arr {
// 					if itm.Val > max { // leg,lss than largest
// 						max = itm.Val
// 						idxFinal = n
// 					}
// 				}
// 			case strings.HasSuffix(prefix, "Geq"), strings.HasSuffix(prefix, "Gtr"):
// 				min := flt.Max
// 				for n, itm := range arr {
// 					if itm.Val < min { // geq,gtr than smallest
// 						min = itm.Val
// 						idxFinal = n
// 					}
// 				}
// 			}
// 		}
// 		finalCnds = append(finalCnds, arr[idxFinal].AndCnd)
// 	}
// 	for n, finalCnd := range finalCnds {
// 		sys.Log("finalCnd", n, finalCnd)
// 	}
// 	return finalCnds
// }

// func (x *Prcp) Name() string {
// 	var sb strings.Builder
// 	for n, stm := range *x.Stms {
// 		if n != 0 {
// 			sb.WriteRune('-')
// 		}
// 		sb.WriteString(stm.Name().Unquo())
// 	}
// 	return sb.String()
// }
