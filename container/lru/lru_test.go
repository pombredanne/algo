package lru_test

import (
	"github.com/larsmans/algo/container/lru"
	"math/rand"
	"strconv"
	"testing"
)

func TestLRU(t *testing.T) {
	cache := lru.New(nil, 4)
	cache.Add("foo", 0)
	cache.Add("bar", 1)
	cache.Add("baz", 2)
	if i, ok := cache.Check("foo"); i != 0 || !ok {
		t.Errorf("expected (0, true), got (%d, %v)", i, ok)
	}

	// Overwrite key.
	cache.Add("baz", 3)
	if i, ok := cache.Check("baz"); i != 3 || !ok {
		t.Errorf("expected (3, true), got (%d, %v)", i, ok)
	}
	if n := cache.Len(); n != 3 {
		t.Errorf("expected Len() == 3, got %d", n)
	}

	cache.Add("quux", 4)
	cache.Add("quuux", 5)
	// "bar" will have been evicted, because we checked for "foo" and "baz".
	all := map[string]int{"foo": 0, "baz": 3, "quux": 4, "quuux": 5}
	cache.Do(func(k, v interface{}) {
		expect, ok := all[k.(string)]
		if !ok {
			t.Errorf("unexpected key in cache: %v", k)
		}
		if v.(int) != expect {
			t.Errorf("wrong value for %s: wanted %d, got %v", k, expect, v)
		}
		delete(all, k.(string))
	})
	if len(all) != 0 {
		t.Errorf("not found in cache: %v", all)
	}
}

// Test cornercase: LRU with capacity 1.
func TestLRU1(t *testing.T) {
	cache := lru.New(nil, 1)
	cache.Add(1, 2)
	cache.Add(1, 3)
	cache.Add("t", "r")
	if k, ok := cache.Check("t"); k.(string) != "r" || !ok {
		t.Error(`expected "r" for "t", got %v, %v`, k, ok)
	}
}

func TestLRUFunc(t *testing.T) {
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
