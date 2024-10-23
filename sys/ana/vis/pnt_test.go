package vis_test

import (
	"fmt"
	"sys/ana/vis"
	"testing"
)

func TestAnglTau(t *testing.T) {
	fmt.Println("TauQtr", vis.TauQtr, " TauHlf", vis.TauHlf, " Tau3Qtr", vis.Tau3Qtr, " Tau", vis.Tau)
	p1 := vis.Pnt{X: 0, Y: 0}

	fmt.Println("HRZ: RHT")
	p2 := vis.Pnt{X: 5, Y: 0}
	angl, dir := p2.AnglTau(p1)
	fmt.Println("p1", p1, "p2", p2, "dir", dir, "angl", angl)

	fmt.Println("HRZ: LFT")
	p2 = vis.Pnt{X: -5, Y: 0}
	angl, dir = p2.AnglTau(p1)
	fmt.Println("p1", p1, "p2", p2, "dir", dir, "angl", angl)

	fmt.Println("VRT: TOP")
	p2 = vis.Pnt{X: 0, Y: -5}
	angl, dir = p2.AnglTau(p1)
	fmt.Println("p1", p1, "p2", p2, "dir", dir, "angl", angl)

	fmt.Println("VRT: BTM")
	p2 = vis.Pnt{X: 0, Y: 5}
	angl, dir = p2.AnglTau(p1)
	fmt.Println("p1", p1, "p2", p2, "dir", dir, "angl", angl)

	fmt.Println("QUAD I")
	p2 = vis.Pnt{X: 5, Y: 1}
	angl, dir = p2.AnglTau(p1)
	fmt.Println("p1", p1, "p2", p2, "dir", dir, "angl", angl)
	p2 = vis.Pnt{X: 1, Y: 5}
	angl, dir = p2.AnglTau(p1)
	fmt.Println("p1", p1, "p2", p2, "dir", dir, "angl", angl)

	fmt.Println("QUAD II")
	p2 = vis.Pnt{X: -5, Y: 1}
	angl, dir = p2.AnglTau(p1)
	fmt.Println("p1", p1, "p2", p2, "dir", dir, "angl", angl)
	p2 = vis.Pnt{X: -1, Y: 5}
	angl, dir = p2.AnglTau(p1)
	fmt.Println("p1", p1, "p2", p2, "dir", dir, "angl", angl)

	fmt.Println("QUAD III")
	p2 = vis.Pnt{X: -5, Y: -1}
	angl, dir = p2.AnglTau(p1)
	fmt.Println("p1", p1, "p2", p2, "dir", dir, "angl", angl)
	p2 = vis.Pnt{X: -1, Y: -5}
	angl, dir = p2.AnglTau(p1)
	fmt.Println("p1", p1, "p2", p2, "dir", dir, "angl", angl)

	fmt.Println("QUAD IIII")
	p2 = vis.Pnt{X: 5, Y: -1}
	angl, dir = p2.AnglTau(p1)
	fmt.Println("p1", p1, "p2", p2, "dir", dir, "angl", angl)
	p2 = vis.Pnt{X: 1, Y: -5}
	angl, dir = p2.AnglTau(p1)
	fmt.Println("p1", p1, "p2", p2, "dir", dir, "angl", angl)
}
