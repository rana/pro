package srch

import (
	"fmt"
	"math"
	"math/big"
	"sys/bsc/flt"
)

// FROM https://alistairisrael.wordpress.com/2009/09/22/simple-efficient-pnk-algorithm/
// FROM https://github.com/porkchop/kperms

// const MaxUint64 = uint32(1<<32 - 1)

var one = big.NewInt(1)

type Perm struct {
	N   uint32
	K   uint32
	Idx uint32
	Cnt uint32
	buf []uint32
}

func NewPerm(n, k int) (x *Perm) {
	if n < 1 {
		panic(fmt.Sprintf("n must be >= 1 (n:%v)", n))
	}
	if k > n {
		panic(fmt.Sprintf("k must be <= n (n:%v k:%v)", n, k))
	}
	if k < 1 {
		panic(fmt.Sprintf("k must be >= 1 (k:%v)", k))
	}
	if n > math.MaxUint32 {
		panic(fmt.Sprintf("n must be <= math.MaxUint32 (n:%v)", n))
	}

	x = &Perm{}
	x.N = uint32(n)
	x.K = uint32(k)
	num := factorial(big.NewInt(int64(n)))
	denom := factorial(big.NewInt(int64(n - k)))
	num.Div(num, denom)
	if !num.IsUint64() {
		panic(fmt.Sprintf("num permutation greater than Uint64.Max (actual:%v)", num))
	}
	x.Cnt = uint32(num.Uint64())
	x.Reset()
	return x
}
func AllPermIdxs(n, k int) [][]uint32 {
	p := NewPerm(n, k)
	return p.AllPerms()
}

func (x *Perm) Reset() {
	x.Idx = uint32(math.MaxUint32)
	x.buf = make([]uint32, x.N)
	// for i := uint32(0); i < x.N; i++ {
	// 	x.buf[i] = i
	// }
}

func (x *Perm) Perm() []uint32 { return x.buf[0:x.K] }
func (x *Perm) AllPerms() (r [][]uint32) {
	r = make([][]uint32, x.Cnt)
	for x.Next() {
		r[x.Idx] = make([]uint32, x.K)
		copy(r[x.Idx], x.buf[0:x.K])
	}
	return r
}

func (x *Perm) String() string {
	return fmt.Sprintf("n:%v k:%v cnt:%v idx:%v", x.N, x.K, x.Cnt, x.Idx)
}

func (x *Perm) Prgrs() string {
	return fmt.Sprintf("%v%%  %v of %v", flt.Flt(float32(x.Idx)/float32(x.Cnt)).Mul(100).Trnc(2), x.Idx+1, x.Cnt)
}

func (x *Perm) Next() bool {
	x.Idx++
	if x.Idx >= x.Cnt {
		return false
	}

	if x.Idx == 0 {
		for i := uint32(0); i < x.N; i++ {
			x.buf[i] = i
		}
		return true
	}

	edge := x.K - 1

	// find j in (k…n-1) where aj > aedge
	j := x.K
	for j < x.N && x.buf[edge] >= x.buf[j] {
		j++
	}

	if j < x.N {
		// swap(x.buf, edge, j)
		x.buf[edge], x.buf[j] = x.buf[j], x.buf[edge]
	} else {
		reverse(x.buf, x.K, x.N-1)

		// find rightmost ascent to left of edge
		i := edge - 1
		for i >= 0 && x.buf[i] >= x.buf[i+1] {
			i--
		}

		if i < 0 {
			// no more permutations
			return false
		}

		// find j in (n-1…i+1) where aj > ai
		j = x.N - 1
		for j > i && x.buf[i] >= x.buf[j] {
			j--
		}

		// swap(x.buf, i, j)
		x.buf[i], x.buf[j] = x.buf[j], x.buf[i]
		reverse(x.buf, i+1, x.N-1)
	}

	return x.Idx != x.Cnt
}

func swap(a []uint32, i uint32, j uint32) {
	a[i], a[j] = a[j], a[i]
}

func reverse(a []uint32, i uint32, j uint32) {
	r := a[i : j+1]
	for i, j = 0, j-i; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
}

func factorial(n *big.Int) (r *big.Int) {
	if n.Cmp(one) == 1 { // n > 1
		return big.NewInt(0).Mul(n, factorial(big.NewInt(0).Sub(n, one)))
	}
	return one
}
