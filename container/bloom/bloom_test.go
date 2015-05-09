package bloom_test

import (
	"github.com/larsmans/algo/container/bloom"
	"math"
	"math/rand"
	"testing"
)

func TestBloom(t *testing.T) {
	r := rand.New(rand.NewSource(42))

	f, err := bloom.New32(1<<32-1, 10, r)
	if f != nil || err == nil {
		t.Error("expected an error")
	}

	distinct := make(map[uint32]bool)
	keys := make([]uint32, 1000)
	for i := range keys {
		k := r.Uint32()
		keys[i] = k
		distinct[k] = true
	}

	f1, _ := bloom.New32(1<<11-3, 15, r)
	f2, _ := bloom.New32(1<<11-29, 15, r)
	for _, f := range []*bloom.Filter32{f1, f2} {
		if c := f.Capacity(); c != 1<<11 {
			t.Errorf("expected capacity %d, got %d", 1<<16, c)
		}

		for _, k := range keys {
			f.Add(k)
			if !f.Get(k) {
				t.Fatalf("inserted key %d missing from filter", k)
			}
		}
	}

	n1, n2, actual := f1.NKeys(), f2.NKeys(), float64(len(distinct))
	if n1 != n2 {
		t.Errorf("NKeys() not the same: %f != %f", n1, n2)
	}
	if diff := math.Abs(n1 - actual); diff > .05*actual {
		t.Errorf("NKeys() not accurate: got %f, actual %.0f", n1, actual)
	}

	f1.Clear()
	if n := f1.NKeys(); n != 0 {
		t.Errorf("expected zero keys after Clear(), got %d", n)
	}
}

func BenchmarkBloom(b *testing.B) {
	b.StopTimer()
	r := rand.New(rand.NewSource(18))
	f, _ := bloom.New32(10000, 14, r)

	for i := 0; i < b.N; i++ {
		b.StartTimer()
		for j := uint32(0); j < 5000; j += 2 {
			f.Add(j)
		}
		for j := uint32(0); j < 2500; j++ {
			f.Get(j)
		}
		b.StopTimer()
		f.Clear()
	}
}
