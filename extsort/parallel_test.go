package extsort

import (
	"math/rand"
	"sort"
	"testing"
)

func TestParallel(t *testing.T) {
	a := sort.IntSlice(rand.Perm(10000))
	Parallel(a, 1)
	if !sort.IsSorted(a) {
		t.Error("not sorted")
	}
}

func BenchmarkParallel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		a := sort.IntSlice(rand.Perm(100000))
		b.StartTimer()
		Parallel(a, 10)
		//sort.Sort(a)
	}
}
