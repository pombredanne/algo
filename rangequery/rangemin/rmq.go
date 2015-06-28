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
//
// An Index allows finding the (index of the) minimum for subranges of Data,
// which must stay static for the lifetime of the Index.
type Index struct {
	Data Interface
	n    int // Caches data.Len(); we don't know how expensive that call is.
	table
}

type table struct {
	entry []int
	ncols int
}

func newTable(n, k int) table {
	return table{entry: make([]int, n*k), ncols: k}
}

func (t table) at(i int, j uint) int {
	return t.entry[i*t.ncols+int(j)]
}

func (t table) set(i int, j uint, v int) {
	t.entry[i*t.ncols+int(j)] = v
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
	n := data.Len()
	if n == 0 {
		panic("data.Len() must be > 0")
	}
	logn := intmath.Log2(n)
	t := newTable(n, logn)

	// Base case: unit length ranges.
	for i := 0; i < n; i++ {
		t.set(i, 0, i)
	}
	for j := uint(1); (1 << j) <= n; j++ {
		for i := 0; i+int(1<<j)-1 < n; i++ {
			l := t.at(i, j-1)
			r := t.at(i+int(1<<(j-1)), j-1)

			if data.Less(l, r) {
				t.set(i, j, l)
			} else {
				t.set(i, j, r)
			}
		}
	}

	return &Index{Data: data, n: n, table: t}
}

// Returns the index of the minimum of r.Data[i:j]. Takes constant time.
//
// The index is relative to the start of r.Data. Ties are broken in an
// arbitrary way.
//
// i >= j is a runtime error.
func (r *Index) Min(i, j int) int {
	switch {
	case i >= j:
		panic("got i >= j in Index.Min")
	case j > r.n:
		panic("j > data.Len() in Index.Min")
	case i+1 == j:
		return i
	}

	k := uint(intmath.Log2(j - i))
	a := r.at(i, k)
	b := r.at(j-int(1<<k), k)
	if r.Data.Less(a, b) {
		return a
	}
	return b
}
