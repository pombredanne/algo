package disjsets

import (
	"math/rand"
	"testing"
)

func TestDisjsets(t *testing.T) {
	n := 10
	f := New(n)
	nsets := n

	for i := 0; i < n; i++ {
		if f.Find(i) != i {
			t.Fatalf("Find(%d) != %d", i, i)
		}
	}

	union := func(a, b int, expect bool) {
		if f.Union(a, b) != expect {
			t.Errorf("expected %v, got %v", expect, !expect)
		}
		if f.Find(a) != f.Find(b) {
			t.Errorf("%d and %d not in same set after union", a, b)
		}
		if expect {
			nsets--
			if f.NSets() != nsets {
				t.Errorf("expected NSets = %d, got %d", nsets, f.NSets())
			}
		}
		if f.Len() != n {
			t.Errorf("error in Len: expected %d, got %d", n, f.Len())
		}
	}

	union(1, 5, true)
	union(1, 5, false)
	union(4, 3, true)
	union(4, 2, true)
	union(2, 3, false)
	union(0, 5, true)
}

func BenchmarkDisjsets(b *testing.B) {
	b.StopTimer()
	n := 100000
	r := rand.New(rand.NewSource(37))

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		f := New(n)
		for j := 0; j < 100000; j++ {
			x, y := pair(r, n)
			f.Union(x, y)
		}
	}
}

func pair(r *rand.Rand, n int) (int, int) {
	return r.Intn(n), r.Intn(n)
}
