package vis_test

import (
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"os"
	"sys"
	"sys/ana/vis"
	"sys/ana/vis/clr"
	"sys/err"
	"sys/tst"
	"testing"
)

var (
	stk    = float32(16)
	stkHlf = stk * .5
	stkClr = clr.White

	length    = float32(500)
	siz       = vis.NewSiz(uint32(length), uint32(length))
	pad       = float32(20)
	lnrLength = length - (2 * pad)

	hlf  = lnrLength / 2
	thrd = lnrLength / 3
	qtr  = lnrLength / 4

	stp    = float32(pad) * 1.5
	lft    = pad
	rht    = length - pad
	top    = pad
	btm    = length - pad
	mid    = lft + hlf
	midLft = mid - pad
	midRht = mid + pad
	midTop = mid - pad
	midBtm = mid + pad
	lftQtr = lft + qtr
	rhtQtr = rht - qtr
	topQtr = top + qtr
	btmQtr = btm - qtr
)

func TestVisStkPth2SegLinePerp(t *testing.T) {
	v, img, pxlFns := newStk()

	// HRZ LFT-RHT
	v.Pth(midStp1Rht, mid)
	v.Pth(rhtQtr, mid)
	v.Pth(rht, mid)
	v.RndrDraw(stk, stkClr, img)
	// HRZ RHT-LFT
	v.Pth(midStp1Lft, mid)
	v.Pth(lftQtr, mid)
	v.Pth(lft, mid)
	v.RndrDraw(stk, stkClr, img)
	// VRT BTM-TOP
	v.Pth(mid, midStp1Top)
	v.Pth(mid, topQtr)
	v.Pth(mid, top)
	v.RndrDraw(stk, stkClr, img)
	// VRT TOP-BTM
	v.Pth(mid, midStp1Btm)
	v.Pth(mid, btmQtr)
	v.Pth(mid, btm)
	v.RndrDraw(stk, stkClr, img)

	// QUAD I
	v.Pth(midStp1Rht, midStp1Top)
	v.Pth(rht, midStp1Top)
	v.Pth(rht, top)
	v.RndrDraw(stk, stkClr, img)
	v.Pth(midStp1Rht, midStp2Top)
	v.Pth(midStp1Rht, top)
	v.Pth(rhtStp1, top)
	v.RndrDraw(stk, stkClr, img)

	// QUAD II
	v.Pth(midStp1Lft, midStp1Top)
	v.Pth(lft, midStp1Top)
	v.Pth(lft, top)
	v.RndrDraw(stk, stkClr, img)
	v.Pth(midStp1Lft, midStp2Top)
	v.Pth(midStp1Lft, top)
	v.Pth(lftStp1, top)
	v.RndrDraw(stk, stkClr, img)

	// QUAD III
	v.Pth(midStp1Lft, midStp1Btm)
	v.Pth(lft, midStp1Btm)
	v.Pth(lft, btm)
	v.RndrDraw(stk, stkClr, img)
	v.Pth(midStp1Lft, midStp2Btm)
	v.Pth(midStp1Lft, btm)
	v.Pth(lftStp1, btm)
	v.RndrDraw(stk, stkClr, img)

	// QUAD IIII
	v.Pth(midStp1Rht, midStp1Btm)
	v.Pth(rht, midStp1Btm)
	v.Pth(rht, btm)
	v.RndrDraw(stk, stkClr, img)
	v.Pth(midStp1Rht, midStp2Btm)
	v.Pth(midStp1Rht, btm)
	v.Pth(rhtStp1, btm)
	v.RndrDraw(stk, stkClr, img)

	drawStk(v, img, pxlFns)
}

func TestVisStkPth2SegHrzInOut(t *testing.T) {
	v, img, pxlFns := newStk()

	// QUAD I
	v.Pth(midRht, midStp1Top)
	v.Pth(rhtQtr, midStp1Top)
	v.Pth(rht, midStp2Top)
	v.RndrDraw(stk, stkClr, img)
	v.Pth(midRht, midStp2Top)
	v.Pth(rhtQtr, midStp2Top)
	v.Pth(rht, topQtr)
	v.RndrDraw(stk, stkClr, img)
	v.Pth(midRht, midStp3Top)
	v.Pth(rhtQtr, midStp3Top)
	v.Pth(rht, top)
	v.RndrDraw(stk, stkClr, img)
	v.Pth(midRht, midStp4Top)
	v.Pth(rhtQtr, midStp4Top)
	v.Pth(rhtQtrStp1Rht, top)
	v.RndrDraw(stk, stkClr, img)

	// QUAD II
	v.Pth(midLft, midStp1Top)
	v.Pth(lftQtr, midStp1Top)
	v.Pth(lft, midStp2Top)
	v.RndrDraw(stk, stkClr, img)
	v.Pth(midLft, midStp2Top)
	v.Pth(lftQtr, midStp2Top)
	v.Pth(lft, topQtr)
	v.RndrDraw(stk, stkClr, img)
	v.Pth(midLft, midStp3Top)
	v.Pth(lftQtr, midStp3Top)
	v.Pth(lft, top)
	v.RndrDraw(stk, stkClr, img)
	v.Pth(midLft, midStp4Top)
	v.Pth(lftQtr, midStp4Top)
	v.Pth(lftQtrStp1Lft, top)
	v.RndrDraw(stk, stkClr, img)

	// QUAD III
	v.Pth(midLft, midStp1Btm)
	v.Pth(lftQtr, midStp1Btm)
	v.Pth(lft, midStp2Btm)
	v.RndrDraw(stk, stkClr, img)
	v.Pth(midLft, midStp2Btm)
	v.Pth(lftQtr, midStp2Btm)
	v.Pth(lft, btmQtr)
	v.RndrDraw(stk, stkClr, img)
	v.Pth(midLft, midStp3Btm)
	v.Pth(lftQtr, midStp3Btm)
	v.Pth(lft, btm)
	v.RndrDraw(stk, stkClr, img)
	v.Pth(midLft, midStp4Btm)
	v.Pth(lftQtr, midStp4Btm)
	v.Pth(lftQtrStp1Lft, btm)
	v.RndrDraw(stk, stkClr, img)

	// QUAD IIII
	v.Pth(midRht, midStp1Btm)
	v.Pth(rhtQtr, midStp1Btm)
	v.Pth(rht, midStp2Btm)
	v.RndrDraw(stk, stkClr, img)
	v.Pth(midRht, midStp2Btm)
	v.Pth(rhtQtr, midStp2Btm)
	v.Pth(rht, btmQtr)
	v.RndrDraw(stk, stkClr, img)
	v.Pth(midRht, midStp3Btm)
	v.Pth(rhtQtr, midStp3Btm)
	v.Pth(rht, btm)
	v.RndrDraw(stk, stkClr, img)
	v.Pth(midRht, midStp4Btm)
	v.Pth(rhtQtr, midStp4Btm)
	v.Pth(rhtQtrStp1Rht, btm)
	v.RndrDraw(stk, stkClr, img)

	drawStk(v, img, pxlFns)
}

func TestVisStkPth2SegHrzOutIn(t *testing.T) {
	v, img, pxlFns := newStk()

	// LFT
	v.Pth(lft, midStp1Top)
	v.Pth(midStp1Lft, mid)
	v.Pth(lft, midStp1Btm)
	v.RndrDraw(stk, stkClr, img)
	// RHT
	v.Pth(rht, midStp1Top)
	v.Pth(midStp1Rht, mid)
	v.Pth(rht, midStp1Btm)
	v.RndrDraw(stk, stkClr, img)

	// QUAD I
	v.Pth(rhtQtr, midStp2Top)
	v.Pth(midRht, midStp2Top)
	v.Pth(rht, topQtr)
	v.RndrDraw(stk, stkClr, img)
	v.Pth(rhtQtr, midStp3Top)
	v.Pth(midRht, midStp3Top)
	v.Pth(rht, top)
	v.RndrDraw(stk, stkClr, img)

	// QUAD II
	v.Pth(lftQtr, midStp2Top)
	v.Pth(midLft, midStp2Top)
	v.Pth(lft, topQtr)
	v.RndrDraw(stk, stkClr, img)
	v.Pth(lftQtr, midStp3Top)
	v.Pth(midLft, midStp3Top)
	v.Pth(lft, top)
	v.RndrDraw(stk, stkClr, img)

	// QUAD III
	v.Pth(lftQtr, midStp2Btm)
	v.Pth(midLft, midStp2Btm)
	v.Pth(lft, btmQtr)
	v.RndrDraw(stk, stkClr, img)
	v.Pth(lftQtr, midStp3Btm)
	v.Pth(midLft, midStp3Btm)
	v.Pth(lft, btm)
	v.RndrDraw(stk, stkClr, img)

	// // QUAD IIII
	v.Pth(rhtQtr, midStp2Btm)
	v.Pth(midRht, midStp2Btm)
	v.Pth(rht, btmQtr)
	v.RndrDraw(stk, stkClr, img)
	v.Pth(rhtQtr, midStp3Btm)
	v.Pth(midRht, midStp3Btm)
	v.Pth(rht, btm)
	v.RndrDraw(stk, stkClr, img)

	drawStk(v, img, pxlFns)
}

func TestVisStkPth2SegVrtInOut(t *testing.T) {
	v, img, pxlFns := newStk()

	// QUAD I
	v.Pth(midStp1Rht, midStp1Top)
	v.Pth(midStp1Rht, topQtr)
	v.Pth(midStp2Rht, top)
	v.RndrDraw(stk, stkClr, img)
	v.Pth(midStp2Rht, midStp1Top)
	v.Pth(midStp2Rht, topQtr)
	v.Pth(rhtQtr, top)
	v.RndrDraw(stk, stkClr, img)
	v.Pth(midStp3Rht, midStp1Top)
	v.Pth(midStp3Rht, topQtr)
	v.Pth(rht, top)
	v.RndrDraw(stk, stkClr, img)
	v.Pth(midStp4Rht, midStp1Top)
	v.Pth(midStp4Rht, topQtr)
	v.Pth(rht, topQtrStp1Top)
	v.RndrDraw(stk, stkClr, img)

	// QUAD II
	v.Pth(midStp1Lft, midStp1Top)
	v.Pth(midStp1Lft, topQtr)
	v.Pth(midStp2Lft, top)
	v.RndrDraw(stk, stkClr, img)
	v.Pth(midStp2Lft, midStp1Top)
	v.Pth(midStp2Lft, topQtr)
	v.Pth(lftQtr, top)
	v.RndrDraw(stk, stkClr, img)
	v.Pth(midStp3Lft, midStp1Top)
	v.Pth(midStp3Lft, topQtr)
	v.Pth(lft, top)
	v.RndrDraw(stk, stkClr, img)
	v.Pth(midStp4Lft, midStp1Top)
	v.Pth(midStp4Lft, topQtr)
	v.Pth(lft, topQtrStp1Top)
	v.RndrDraw(stk, stkClr, img)

	// QUAD III
	v.Pth(midStp1Lft, midStp1Btm)
	v.Pth(midStp1Lft, btmQtr)
	v.Pth(midStp2Lft, btm)
	v.RndrDraw(stk, stkClr, img)
	v.Pth(midStp2Lft, midStp1Btm)
	v.Pth(midStp2Lft, btmQtr)
	v.Pth(lftQtr, btm)
	v.RndrDraw(stk, stkClr, img)
	v.Pth(midStp3Lft, midStp1Btm)
	v.Pth(midStp3Lft, btmQtr)
	v.Pth(lft, btm)
	v.RndrDraw(stk, stkClr, img)
	v.Pth(midStp4Lft, midStp1Btm)
	v.Pth(midStp4Lft, btmQtr)
	v.Pth(lft, btmQtrStp1Btm)
	v.RndrDraw(stk, stkClr, img)

	// QUAD IIII
	v.Pth(midStp1Rht, midStp1Btm)
	v.Pth(midStp1Rht, btmQtr)
	v.Pth(midStp2Rht, btm)
	v.RndrDraw(stk, stkClr, img)
	v.Pth(midStp2Rht, midStp1Btm)
	v.Pth(midStp2Rht, btmQtr)
	v.Pth(rhtQtr, btm)
	v.RndrDraw(stk, stkClr, img)
	v.Pth(midStp3Rht, midStp1Btm)
	v.Pth(midStp3Rht, btmQtr)
	v.Pth(rht, btm)
	v.RndrDraw(stk, stkClr, img)
	v.Pth(midStp4Rht, midStp1Btm)
	v.Pth(midStp4Rht, btmQtr)
	v.Pth(rht, btmQtrStp1Btm)
	v.RndrDraw(stk, stkClr, img)

	drawStk(v, img, pxlFns)
}

func TestVisStkPth2SegVrtOutIn(t *testing.T) {
	v, img, pxlFns := newStk()

	// TOP
	v.Pth(midStp1Lft, top)
	v.Pth(mid, midStp1Top)
	v.Pth(midStp1Rht, top)
	v.RndrDraw(stk, stkClr, img)
	// BTM
	v.Pth(midStp1Lft, btm)
	v.Pth(mid, midStp1Btm)
	v.Pth(midStp1Rht, btm)
	v.RndrDraw(stk, stkClr, img)

	// QUAD I
	v.Pth(midStp2Rht, topQtr)
	v.Pth(midStp2Rht, midStp1Top)
	v.Pth(rhtQtr, top)
	v.RndrDraw(stk, stkClr, img)
	v.Pth(midStp4Rht, topQtr)
	v.Pth(midStp4Rht, midStp1Top)
	v.Pth(midStp5Rht, top)
	v.RndrDraw(stk, stkClr, img)

	// QUAD II
	v.Pth(midStp2Lft, topQtr)
	v.Pth(midStp2Lft, midStp1Top)
	v.Pth(lftQtr, top)
	v.RndrDraw(stk, stkClr, img)
	v.Pth(midStp4Lft, topQtr)
	v.Pth(midStp4Lft, midStp1Top)
	v.Pth(midStp5Lft, top)
	v.RndrDraw(stk, stkClr, img)

	// QUAD III
	v.Pth(midStp2Lft, btmQtr)
	v.Pth(midStp2Lft, midStp1Btm)
	v.Pth(lftQtr, btm)
	v.RndrDraw(stk, stkClr, img)
	v.Pth(midStp3Lft, btmQtr)
	v.Pth(midStp3Lft, midStp1Btm)
	v.Pth(lft, btm)
	v.RndrDraw(stk, stkClr, img)

	// QUAD IIII
	v.Pth(midStp2Rht, btmQtr)
	v.Pth(midStp2Rht, midStp1Btm)
	v.Pth(rhtQtr, btm)
	v.RndrDraw(stk, stkClr, img)
	v.Pth(midStp3Rht, btmQtr)
	v.Pth(midStp3Rht, midStp1Btm)
	v.Pth(rht, btm)
	v.RndrDraw(stk, stkClr, img)

	drawStk(v, img, pxlFns)
}

func TestVisStkPth2SegDiagClose(t *testing.T) {
	v, img, pxlFns := newStk()

	// LFT
	v.Pth(lft, midStp1Top)
	v.Pth(midStp1Lft, mid)
	v.Pth(lft, midStp1Btm)
	v.RndrDraw(stk, stkClr, img)
	// RHT
	v.Pth(rht, midStp1Top)
	v.Pth(midStp1Rht, mid)
	v.Pth(rht, midStp1Btm)
	v.RndrDraw(stk, stkClr, img)

	// TOP
	v.Pth(midStp1Lft, top)
	v.Pth(mid, midStp1Top)
	v.Pth(midStp1Rht, top)
	v.RndrDraw(stk, stkClr, img)
	// BTM
	v.Pth(midStp1Lft, btm)
	v.Pth(mid, midStp1Btm)
	v.Pth(midStp1Rht, btm)
	v.RndrDraw(stk, stkClr, img)

	drawStk(v, img, pxlFns)
}

func TestVisStkPth2SegDiagInOut(t *testing.T) {
	v, img, pxlFns := newStk()

	// QUAD I
	v.Pth(midRht, topStp2)
	v.Pth(rhtQtr, topStp1)
	v.Pth(rhtQtr, top)
	v.RndrDraw(stk, stkClr, img)
	v.Pth(midRht, topStp3)
	v.Pth(rhtQtr, topStp2)
	v.Pth(midRht, topStp4)
	v.RndrDraw(stk, stkClr, img)
	v.Pth(midRht, topStp5)
	v.Pth(rhtQtr, topStp3)
	v.Pth(rht, topStp3)
	v.RndrDraw(stk, stkClr, img)
	v.Pth(midRht, midStp2Top)
	v.Pth(rhtQtr, midStp3Top)
	v.Pth(rht, midStp2Top)
	v.RndrDraw(stk, stkClr, img)
	v.Pth(midRht, midStp1Top)
	v.Pth(rhtQtr, midStp2Top)
	v.Pth(rhtQtr, midStp1Top)
	v.RndrDraw(stk, stkClr, img)

	// QUAD II
	v.Pth(midLft, topStp2)
	v.Pth(lftQtr, topStp1)
	v.Pth(lftQtr, top)
	v.RndrDraw(stk, stkClr, img)
	v.Pth(midLft, topStp3)
	v.Pth(lftQtr, topStp2)
	v.Pth(midLft, topStp4)
	v.RndrDraw(stk, stkClr, img)
	v.Pth(midLft, topStp5)
	v.Pth(lftQtr, topStp3)
	v.Pth(lft, topStp3)
	v.RndrDraw(stk, stkClr, img)
	v.Pth(midLft, midStp2Top)
	v.Pth(lftQtr, midStp3Top)
	v.Pth(lft, midStp2Top)
	v.RndrDraw(stk, stkClr, img)
	v.Pth(midLft, midStp1Top)
	v.Pth(lftQtr, midStp2Top)
	v.Pth(lftQtr, midStp1Top)
	v.RndrDraw(stk, stkClr, img)

	// QUAD III
	v.Pth(midLft, btmStp2)
	v.Pth(lftQtr, btmStp1)
	v.Pth(lftQtr, btm)
	v.RndrDraw(stk, stkClr, img)
	v.Pth(midLft, btmStp3)
	v.Pth(lftQtr, btmStp2)
	v.Pth(midLft, btmStp4)
	v.RndrDraw(stk, stkClr, img)
	v.Pth(midLft, btmStp5)
	v.Pth(lftQtr, btmStp3)
	v.Pth(lft, btmStp3)
	v.RndrDraw(stk, stkClr, img)
	v.Pth(midLft, midStp2Btm)
	v.Pth(lftQtr, midStp3Btm)
	v.Pth(lft, midStp2Btm)
	v.RndrDraw(stk, stkClr, img)
	v.Pth(midLft, midStp1Btm)
	v.Pth(lftQtr, midStp2Btm)
	v.Pth(lftQtr, midStp1Btm)
	v.RndrDraw(stk, stkClr, img)

	// QUAD IIII
	v.Pth(midRht, btmStp2)
	v.Pth(rhtQtr, btmStp1)
	v.Pth(rhtQtr, btm)
	v.RndrDraw(stk, stkClr, img)
	v.Pth(midRht, btmStp3)
	v.Pth(rhtQtr, btmStp2)
	v.Pth(midRht, btmStp4)
	v.RndrDraw(stk, stkClr, img)
	v.Pth(midRht, btmStp5)
	v.Pth(rhtQtr, btmStp3)
	v.Pth(rht, btmStp3)
	v.RndrDraw(stk, stkClr, img)
	v.Pth(midRht, midStp2Btm)
	v.Pth(rhtQtr, midStp3Btm)
	v.Pth(rht, midStp2Btm)
	v.RndrDraw(stk, stkClr, img)
	v.Pth(midRht, midStp1Btm)
	v.Pth(rhtQtr, midStp2Btm)
	v.Pth(rhtQtr, midStp1Btm)
	v.RndrDraw(stk, stkClr, img)

	drawStk(v, img, pxlFns)
}

func TestVisStkPth2SegDiagOutIn(t *testing.T) {
	v, img, pxlFns := newStk()

	// QUAD I
	v.Pth(rht, topStp1)
	v.Pth(rhtQtr, topStp2)
	v.Pth(rhtQtr, topStp1)
	v.RndrDraw(stk, stkClr, img)
	v.Pth(rht, topStp2)
	v.Pth(rhtQtr, topStp3)
	v.Pth(midRht, topStp2)
	v.RndrDraw(stk, stkClr, img)
	v.Pth(rht, topStp3)
	v.Pth(rhtQtr, topStp4)
	v.Pth(midRht, topStp4)
	v.RndrDraw(stk, stkClr, img)
	v.Pth(rht, topStp4)
	v.Pth(rhtQtr, topStp5)
	v.Pth(rht, topStp5)
	v.RndrDraw(stk, stkClr, img)
	v.Pth(rht, topStp6)
	v.Pth(rhtQtr, topStp6)
	v.Pth(rhtQtr, midStp1Top)
	v.RndrDraw(stk, stkClr, img)

	// QUAD II
	v.Pth(lft, topStp1)
	v.Pth(lftQtr, topStp2)
	v.Pth(lftQtr, topStp1)
	v.RndrDraw(stk, stkClr, img)
	v.Pth(lft, topStp2)
	v.Pth(lftQtr, topStp3)
	v.Pth(midLft, topStp2)
	v.RndrDraw(stk, stkClr, img)
	v.Pth(lft, topStp3)
	v.Pth(lftQtr, topStp4)
	v.Pth(midLft, topStp4)
	v.RndrDraw(stk, stkClr, img)
	v.Pth(lft, topStp4)
	v.Pth(lftQtr, topStp5)
	v.Pth(lft, topStp5)
	v.RndrDraw(stk, stkClr, img)
	v.Pth(lft, topStp6)
	v.Pth(lftQtr, topStp6)
	v.Pth(lftQtr, midStp1Top)
	v.RndrDraw(stk, stkClr, img)

	// QUAD III
	v.Pth(lft, btmStp1)
	v.Pth(lftQtr, btmStp2)
	v.Pth(lftQtr, btmStp1)
	v.RndrDraw(stk, stkClr, img)
	v.Pth(lft, btmStp2)
	v.Pth(lftQtr, btmStp3)
	v.Pth(midLft, btmStp2)
	v.RndrDraw(stk, stkClr, img)
	v.Pth(lft, btmStp3)
	v.Pth(lftQtr, btmStp4)
	v.Pth(midLft, btmStp4)
	v.RndrDraw(stk, stkClr, img)
	v.Pth(lft, btmStp4)
	v.Pth(lftQtr, btmStp5)
	v.Pth(lft, btmStp5)
	v.RndrDraw(stk, stkClr, img)
	v.Pth(lft, btmStp6)
	v.Pth(lftQtr, btmStp6)
	v.Pth(lftQtr, midStp1Btm)
	v.RndrDraw(stk, stkClr, img)

	// QUAD IIII
	v.Pth(rht, btmStp1)
	v.Pth(rhtQtr, btmStp2)
	v.Pth(rhtQtr, btmStp1)
	v.RndrDraw(stk, stkClr, img)
	v.Pth(rht, btmStp2)
	v.Pth(rhtQtr, btmStp3)
	v.Pth(midRht, btmStp2)
	v.RndrDraw(stk, stkClr, img)
	v.Pth(rht, btmStp3)
	v.Pth(rhtQtr, btmStp4)
	v.Pth(midRht, btmStp4)
	v.RndrDraw(stk, stkClr, img)
	v.Pth(rht, btmStp4)
	v.Pth(rhtQtr, btmStp5)
	v.Pth(rht, btmStp5)
	v.RndrDraw(stk, stkClr, img)
	v.Pth(rht, btmStp6)
	v.Pth(rhtQtr, btmStp6)
	v.Pth(rhtQtr, midStp1Btm)
	v.RndrDraw(stk, stkClr, img)

	drawStk(v, img, pxlFns)
}

func TestVisStkPthMany(t *testing.T) {
	v, img, pxlFns := newStk()

	v.Pth(lft, top)
	v.Pth(lftQtr, topQtr)
	v.Pth(mid, topStp1)
	v.Pth(rhtQtr, mid)
	v.Pth(rhtStp2, topStp1)
	v.Pth(rht, top)
	v.RndrDraw(stk, stkClr, img)

	v.Pth(lft, topQtr)
	v.Pth(lftQtr, top)
	v.Pth(topStp1, mid)
	v.Pth(mid, rhtQtr)
	v.Pth(rhtStp2, btmStp1)
	v.Pth(rhtStp2, btm)
	v.RndrDraw(stk, stkClr, img)

	drawStk(v, img, pxlFns)
}

func TestVisStkPthNaNCheck(t *testing.T) {
	v, img, pxlFns := newStk()
	curWidth := float32(1)
	img = image.NewRGBA(image.Rect(0, 0, 2560, 1440))
	draw.Draw(img, img.Bounds(), clr.Black.Uniform(), image.Point{}, draw.Src) // draw background

	// NAN CHECK
	v.Pth(69.562225, 1230.6343)
	v.Pth(70.96111, 1236.2117)
	v.Pth(75.157776, 1252.9438)

	v.RndrDraw(curWidth, stkClr, img)

	drawStk(v, img, pxlFns)
}

// // func TestVisTxt(t *testing.T) {
// // 	v, img, pxlFns := newStk()
// // 	v.PxlLnHrz(lft, rht, mid, gridClr)
// // 	v.PxlLnVrt(mid, top, btm, gridClr)
// // 	v.PxlRct(vis.Rct{Lft: lft, Rht: rht, Top: top, Btm: btm}, gridClr)

// // 	txt := "Aum"
// // 	face := roboto.Medium(12)

// // 	// CENTER
// // 	v.Txt(mid, mid, .5, .3, face, txt)
// // 	// LFT-TOP
// // 	v.Txt(lft, top, 1, 0, face, txt)
// // 	// RHT-TOP
// // 	v.Txt(rht, top, 0, 0, face, txt)
// // 	// LFT-BTM
// // 	v.Txt(lft, btm, 1, 1, face, txt)
// // 	// LFT-BTM
// // 	v.Txt(rht, btm, 0, 1, face, txt)

// // 	drawStk(v, img, pxlFns)
// // }

func newStk() (s *vis.Stk, img *image.RGBA, pxlFns []func()) {
	fmt.Println("-----", sys.FnName(1))
	s = vis.NewStk()
	img = image.NewRGBA(image.Rect(0, 0, int(length), int(length)))
	draw.Draw(img, img.Bounds(), clr.Black.Uniform(), image.Point{}, draw.Src) // draw background
	// v = vis.NewVis(vis.NewSiz(uint32(length), uint32(length)))
	// v.DrawBak(clr.Black)
	// v.SetStk(stk)
	// vis.PxlRct(vis.Rct{Rht: uint32(length), Btm: uint32(length)}, clr.Grey700.Color())
	return s, img, pxlFns
}
func drawStk(s *vis.Stk, img *image.RGBA, pxlFns []func()) {
	for _, pxlFn := range pxlFns {
		pxlFn()
	}
	path := fmt.Sprintf("%v%v.png", *tst.ImgDir, sys.FnName(1))
	file, er := os.Create(path)
	if er != nil {
		err.Panic(er)
	}
	defer file.Close()
	er = png.Encode(file, img)
	if er != nil {
		err.Panic(er)
	}
	if *tst.ImgOpn {
		sys.OpnImg(path)
	}
}

func encodePNG(fnName string, src image.Image) error {
	dstFilename := fmt.Sprintf("%v%v.png", *tst.ImgDir, fnName)
	f, err := os.Create(dstFilename)
	if err != nil {
		return err
	}
	encErr := png.Encode(f, src)
	closeErr := f.Close()
	if encErr != nil {
		return encErr
	}
	return closeErr
}

var (
	midStp1Top = mid - stp
	midStp2Top = mid - (2 * stp)
	midStp3Top = mid - (3 * stp)
	midStp4Top = mid - (4 * stp)
	midStp5Top = mid - (5 * stp)
	midStp6Top = mid - (6 * stp)
	midStp1Btm = mid + stp
	midStp2Btm = mid + (2 * stp)
	midStp3Btm = mid + (3 * stp)
	midStp4Btm = mid + (4 * stp)
	midStp5Btm = mid + (5 * stp)
	midStp6Btm = mid + (6 * stp)

	midStp1Lft = midStp1Top
	midStp2Lft = midStp2Top
	midStp3Lft = midStp3Top
	midStp4Lft = midStp4Top
	midStp5Lft = midStp5Top
	midStp6Lft = midStp6Top
	midStp1Rht = midStp1Btm
	midStp2Rht = midStp2Btm
	midStp3Rht = midStp3Btm
	midStp4Rht = midStp4Btm
	midStp5Rht = midStp5Btm
	midStp6Rht = midStp6Btm

	topStp1 = top + stp
	topStp2 = top + (2 * stp)
	topStp3 = top + (3 * stp)
	topStp4 = top + (4 * stp)
	topStp5 = top + (5 * stp)
	topStp6 = top + (6 * stp)
	btmStp1 = btm - stp
	btmStp2 = btm - (2 * stp)
	btmStp3 = btm - (3 * stp)
	btmStp4 = btm - (4 * stp)
	btmStp5 = btm - (5 * stp)
	btmStp6 = btm - (6 * stp)
	lftStp1 = lft + stp
	lftStp2 = lft + (2 * stp)
	lftStp3 = lft + (3 * stp)
	lftStp4 = lft + (4 * stp)
	lftStp5 = lft + (5 * stp)
	lftStp6 = lft + (6 * stp)
	rhtStp1 = rht - stp
	rhtStp2 = rht - (2 * stp)
	rhtStp3 = rht - (3 * stp)
	rhtStp4 = rht - (4 * stp)
	rhtStp5 = rht - (5 * stp)
	rhtStp6 = rht - (6 * stp)

	rhtQtrStp1Rht = rhtQtr + stp
	rhtQtrStp2Rht = rhtQtr + (2 * stp)
	rhtQtrStp3Rht = rhtQtr + (3 * stp)
	rhtQtrStp4Rht = rhtQtr + (4 * stp)
	rhtQtrStp5Rht = rhtQtr + (5 * stp)
	rhtQtrStp6Rht = rhtQtr + (6 * stp)
	rhtQtrStp1Lft = rhtQtr - stp
	rhtQtrStp2Lft = rhtQtr - (2 * stp)
	rhtQtrStp3Lft = rhtQtr - (3 * stp)
	rhtQtrStp4Lft = rhtQtr - (4 * stp)
	rhtQtrStp5Lft = rhtQtr - (5 * stp)
	rhtQtrStp6Lft = rhtQtr - (6 * stp)

	lftQtrStp1Rht = lftQtr + stp
	lftQtrStp2Rht = lftQtr + (2 * stp)
	lftQtrStp3Rht = lftQtr + (3 * stp)
	lftQtrStp4Rht = lftQtr + (4 * stp)
	lftQtrStp5Rht = lftQtr + (5 * stp)
	lftQtrStp6Rht = lftQtr + (6 * stp)
	lftQtrStp1Lft = lftQtr - stp
	lftQtrStp2Lft = lftQtr - (2 * stp)
	lftQtrStp3Lft = lftQtr - (3 * stp)
	lftQtrStp4Lft = lftQtr - (4 * stp)
	lftQtrStp5Lft = lftQtr - (5 * stp)
	lftQtrStp6Lft = lftQtr - (6 * stp)

	topQtrStp1Top = topQtr - stp
	topQtrStp2Top = topQtr - (2 * stp)
	topQtrStp3Top = topQtr - (3 * stp)
	topQtrStp4Top = topQtr - (4 * stp)
	topQtrStp5Top = topQtr - (5 * stp)
	topQtrStp6Top = topQtr - (6 * stp)
	topQtrStp1Btm = topQtr + stp
	topQtrStp2Btm = topQtr + (2 * stp)
	topQtrStp3Btm = topQtr + (3 * stp)
	topQtrStp4Btm = topQtr + (4 * stp)
	topQtrStp5Btm = topQtr + (5 * stp)
	topQtrStp6Btm = topQtr + (6 * stp)

	btmQtrStp1Top = btmQtr - stp
	btmQtrStp2Top = btmQtr - (2 * stp)
	btmQtrStp3Top = btmQtr - (3 * stp)
	btmQtrStp4Top = btmQtr - (4 * stp)
	btmQtrStp5Top = btmQtr - (5 * stp)
	btmQtrStp6Top = btmQtr - (6 * stp)
	btmQtrStp1Btm = btmQtr + stp
	btmQtrStp2Btm = btmQtr + (2 * stp)
	btmQtrStp3Btm = btmQtr + (3 * stp)
	btmQtrStp4Btm = btmQtr + (4 * stp)
	btmQtrStp5Btm = btmQtr + (5 * stp)
	btmQtrStp6Btm = btmQtr + (6 * stp)
)
