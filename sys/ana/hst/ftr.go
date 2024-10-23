package hst

// /*
// // Start with the basic example from https://docs.julialang.org/en/release-0.6/manual/embedding/
// //
// // Obviously the paths below may need to be modified to match your julia install location and version number.
// //
// // #cgo CFLAGS: -fPIC -DJULIA_INIT_DIR="/Applications/Julia-0.6.app/Contents/Resources/julia/lib" -I/Applications/Julia-0.6.app/Contents/Resources/julia/include/julia -I.
// // #cgo LDFLAGS: -L/Applications/Julia-0.6.app/Contents/Resources/julia/lib/julia  -L/Applications/Julia-0.6.app/Contents/Resources/julia/lib -Wl,-rpath,/Applications/Julia-0.6.app/Contents/Resources/julia/lib -ljulia
// //
// // MUST ADD TO julia.h
// // 		#define JULIA_ENABLE_THREADING
// //
// //
// #cgo CFLAGS: -fPIC -DJULIA_INIT_DIR="/usr/local/julia-1.0.0/lib" -I/usr/local/julia-1.0.0/include/julia -I.
// #cgo LDFLAGS: -L/usr/local/julia-1.0.0/lib/julia  -L/usr/local/julia-1.0.0/lib -Wl,-rpath,/usr/local/julia-1.0.0/lib -ljulia
// #include <julia.h>
// */
// import "C"
// import (
// 	"sys/bsc/flts"
// 	"sys/bsc/tmes"
// 	// "sys/err"
// 	// "io/ioutil"
// 	// "sys/ana/ml"
// )

// func (x *FtrMlFtr) Calc() {
// 	// PRODUCE MODEL FEATURES
// 	// NORMALIZE ANY STMS AS NEEDED
// 	// COPY STMS TO ARRAY FOR USE BY NEURAL NET
// 	x.Stgy.Port() // call port to generate stgy trds
// 	trds := x.Stgy.Bse().Trds
// 	if len(*trds) == 0 {
// 		return
// 	}
// 	trdTmes := tmes.Make(trds.Cnt())
// 	x.Lbls = make([]float32, trds.Cnt())
// 	for n, trd := range *trds {
// 		(*trdTmes)[n] = trd.OpnTme      // gather trade times
// 		x.Lbls[n] = float32(trd.PnlPct) // gather ml feature labels
// 	}
// 	x.Ftrs = make([]*flts.Flts, len(*x.Stms)+len(*x.StmsToNorm))
// 	x.FtrNames = make([]string, len(*x.Stms)+len(*x.StmsToNorm))
// 	for n := 0; n < len(*x.Stms); n++ { // gather stm vals (no normalization)
// 		x.FtrNames[n] = (*x.Stms)[n].String()
// 		x.Ftrs[n] = (*x.Stms)[n].At(trdTmes)
// 	}
// 	for n := 0; n < len(*x.StmsToNorm); n++ { // gather stm vals and normalize
// 		x.FtrNames[len(*x.Stms)+n] = (*x.StmsToNorm)[n].String()
// 		x.Ftrs[len(*x.Stms)+n] = (*x.StmsToNorm)[n].At(trdTmes)
// 		x.Ftrs[len(*x.Stms)+n].ZscrInplace() // normalize with z-score
// 	}
// }
