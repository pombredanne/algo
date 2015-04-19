package lru

import "testing"

func TestLRU(t *testing.T) {
	var ncalls int
	makeobj := func(key interface{}) interface{} {
		ncalls++
		return key
	}

	cache := New(makeobj, 3)
	for i, k := range []int{1, 2, 3, 1, 2, 4, 5, 4, 4, 1} {
		if cache.Len() != min(i, 3) {
			t.Errorf("expected cache Len() of %d, got %d",
				min(i, 3), cache.Len())
		}
		if x := cache.Get(k); x != k {
			t.Errorf("expected %d, got %v", k, x)
		}
	}
	if ncalls != 6 {
		t.Errorf("expected six calls, got %d", ncalls)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
