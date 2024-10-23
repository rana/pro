package ana_test

// Buy Market AUD/USD
// 166,666 Units @ 0.75441
// 5/13/2018, 3:01:29 PM
// Half Spread Cost -9.1666
// Commission -8.33
// Balance of 9991.6

// Close Trade AUD/USD
// 166,666 Units @ 0.75469
// 5/13/2018, 3:58:05 PM
// Half Spread Cost -4.1666
// Commission -8.33
// Balance of 10029.88

// PnlGrsUsd of 46.67 (Reported by Oanda UI)
// func TestAnaTrdLong(t *testing.T) {
// 	// validate my pricing calculations against Oanda pricing calculations
// 	// Oan.Instr ana.instr(name:"aud_usd" typ:"CURRENCY" pip:0.0001 marginRate:0.03 displayPrecision:5 tradeUnitsPrecision:0 minTrdSize:1 maxTrailingStopDistance:1.0 minTrailingStopDistance:0.0005 maxPositionSize:0 maxOrderUnits:100000000 spdMin:0.19 spdMax:9.9 spdMdn:0.39 spdAvg:0.54 spdStd:0.62 spdOpnLim:0.68 dayCnt:46.0 fst:2018-4-1/21h0m0s lst:2018-5-24/3h58m0s)

// 	ap := app.New(tst.Cfg) // for Cfg.Oan.ComPerUnitUsd in trd.CalcPnl
// 	defer ap.Cls()
// 	balUsd := flt.Flt(10000)
// 	trdPct := flt.Flt(.377205)
// 	marginRate := flt.Flt(0.03)
// 	mrgnRtio := flt.Flt(1 / marginRate)

// 	var trd ana.Trd
// 	trd.IsLong = true
// 	trd.OpnAsk = 0.75441
// 	trd.ClsBid = 0.75469
// 	trd.CalcPnl(balUsd, trdPct, mrgnRtio)
// 	tst.FltEql(t, 10030.874, trd.ClsBalUsd)
// }

// Sell Market AUD/USD
// 167,164 Units @ 0.75463
// 5/13/2018, 3:58:25 PM
// Half Spread Cost -5.8507
// Commission -8.36
// Balance of 10021.53

// Close Trade AUD/USD
// 167,164 Units @ 0.75634
// 5/13/2018, 5:59:48 PM
// Half Spread Cost -3.3433
// Commission -8.36
// Balance of 9727.14

// PnlGrsUsd of -285.85 (Reported by Oanda UI)
// func TestAnaTrdShrt(t *testing.T) {
// 	// validate my pricing calculations against Oanda pricing calculations
// 	// Oan.Instr ana.instr(name:"aud_usd" typ:"CURRENCY" pip:0.0001 marginRate:0.03 displayPrecision:5 tradeUnitsPrecision:0 minTrdSize:1 maxTrailingStopDistance:1.0 minTrailingStopDistance:0.0005 maxPositionSize:0 maxOrderUnits:100000000 spdMin:0.19 spdMax:9.9 spdMdn:0.39 spdAvg:0.54 spdStd:0.62 spdOpnLim:0.68 dayCnt:46.0 fst:2018-4-1/21h0m0s lst:2018-5-24/3h58m0s)

// 	ap := app.New(tst.Cfg) // for Cfg.Oan.ComPerUnitUsd in trd.CalcPnl
// 	defer ap.Cls()
// 	balUsd := flt.Flt(10021.53)
// 	trdPct := flt.Flt(.378485) // ?
// 	marginRate := flt.Flt(0.03)
// 	mrgnRtio := flt.Flt(1 / marginRate)

// 	var trd ana.Trd
// 	trd.IsLong = false
// 	trd.OpnBid = 0.75463
// 	trd.ClsAsk = 0.75634

// 	trd.CalcPnl(balUsd, trdPct, mrgnRtio)
// 	tst.FltEql(t, 9727.189, trd.ClsBalUsd)
// }
