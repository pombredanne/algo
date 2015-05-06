package lru_test

import (
	"github.com/larsmans/algo/container/lru"
	"math/rand"
	"strconv"
	"testing"
)

func TestLRU(t *testing.T) {
	var ncalls int
	makeobj := func(key interface{}) interface{} {
		ncalls++
		return key
	}

	cache := lru.New(makeobj, 3)
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

func BenchmarkStrconv1000(b *testing.B) {
	b.StopTimer()
	cache := lru.New(func(x interface{}) interface{} {
		i := x.(int)
		return strconv.Itoa(i)
	}, 1000)

	rng := rand.New(rand.NewSource(42))

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		cache.Get(rng.Intn(0xFFFFFFFF))
	}
}
