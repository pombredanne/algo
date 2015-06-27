// Package rangemin implement a range-min query (RMQ) index.
package rangemin

import "github.com/larsmans/algo/extmath/intmath"

// Interface that represents static arrays of well-ordered items.
//
// Note that this is a subset of sort.Interface.
type Interface interface {
	Len() int
	Less(int, int) bool
}

// Range-min query index for Data.
type Index struct {
	Data Interface
	table
}

type table struct {
	entry []int
	ncols uint
}

func newTable(n, k uint) table {
	return table{entry: make([]int, int(n*k)), ncols: uint(k)}
}

func (t table) at(i, j uint) int {
	return t.entry[int(i*t.ncols+j)]
}

func (t table) set(i, j uint, v int) {
	t.entry[int(i*t.ncols+j)] = v
}

// Construct new range-min query index.
//
// The data argument is stored on the index. Modifying its contents may
// invalidate the range-min index's results.
//
// Takes O(n log n) time where n = data.Len().
func New(data Interface) *Index {
	// Sparse table DP algorithm from Bender et al.:
	// https://www3.cs.stonybrook.edu/~bender/pub/JALG05-daglca.pdf (§2.3)
	n := uint(data.Len())
	if n == 0 {
		panic("data.Len() must be > 0")
	}
	logn := uint(intmath.Log2(int(n)))
	t := newTable(n, logn)

	// Base case: unit length ranges.
	for i := uint(0); i < n; i++ {
		t.set(i, 0, int(i))
	}
	for j := uint(1); (1 << j) <= n; j++ {
		for i := uint(0); i+(1<<j)-1 < n; i++ {
			l := t.at(i, j-1)
			r := t.at(i+(1<<(j-1)), j-1)

			if data.Less(l, r) {
				t.set(i, j, l)
			} else {
				t.set(i, j, r)
			}
		}
	}

	return &Index{Data: data, table: t}
}

// Returns the index of the minimum of r.Data[i:j]. Takes constant time.
//
// The index is relative to the start of r.Data.
//
// i > j is a runtime error.
func (r *Index) Min(i, j int) int {
	switch {
	case i == j:
		return i
	case i > j:
		panic("got i > j in Index.Min")
	}

	k := uint(intmath.Log2(j - i + 1))
	a := r.at(uint(i), k)
	b := r.at(uint(j-(1<<k)+1), k)
	if r.Data.Less(a, b) {
		return a
	}
	return b
}
