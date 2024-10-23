package tst

// func TicsEql(t *testing.T, e []*oan.Tic, a []*prv.Tic, msgs ...interface{}) {
// 	if e == nil && a == nil || len(e) == 0 && len(a) == 0 {
// 		return
// 	}
// 	if e == nil {
// 		t.Helper()
// 		t.Fatal(append(msgs, "Tmes expected is nil")...)
// 	}
// 	if a == nil {
// 		t.Helper()
// 		t.Fatal(append(msgs, fmt.Errorf("Tmes actual is nil (expected:%v)", len(e)))...)
// 	}
// 	if len(e) != len(a) {
// 		t.Helper()
// 		t.Fatal(append(msgs, fmt.Errorf("Tmes length unequal (expected:%v actual:%v)", len(e), len(a)))...)
// 	}
// 	for n := 0; n < len(e); n++ {
// 		eCur := prv.Tic{
// 			MayTrd: e[n].MayTrd(),
// 			Tme:    e[n].Tme,
// 			Bids:   e[n].Bids,
// 			Asks:   e[n].Asks,
// 		}
// 		if eCur.Neq(a[n]) {
// 			t.Helper()
// 			t.Fatal(append(msgs, fmt.Errorf("Tics element unequal \n      idx:%v \n expected:%v \n   actual:%v", n, e[n], a[n]))...)
// 		}
// 	}
// }
