package intmath

import (
	"math/rand"
	"testing"
)

func TestLog2(t *testing.T) {
	for _, c := range []struct{ n, log int }{
		{-5, -1}, {0, -1}, {1, 0}, {1 << 16, 16}, {1<<4 - 1, 3},
		{1<<5 + 3, 5}, {1<<33 + 100, 33},
	} {
		if got := Log2(c.n); got != c.log {
			t.Errorf("expected %d, got %d", c.log, got)
		}
	}
}

func BenchmarkLog2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// Big strides because time of naive algo is O(log n).
		for j := 0; j < 100000000; j += 10000 {
			Log2(j)
		}
	}
}

func TestPopcount(t *testing.T) {
	for _, c := range []struct {
		n   uint32
		pop int
	}{
		{0, 0}, {1 << 5, 1}, {1<<32 - 1, 32}, {1<<32 - 2, 31}, {69161, 7},
	} {
		if got := Popcount32(c.n); got != c.pop {
			t.Errorf("expected %d, got %d", c.pop, got)
		}
		if got := Popcount64(uint64(c.n)); got != c.pop {
			t.Errorf("expected %d, got %d", c.pop, got)
		}
	}
}

func BenchmarkPopcount(b *testing.B) {
	b.StopTimer()
	r := rand.New(rand.NewSource(747))
	input := make([]uint32, 100000)
	for i := range input {
		input[i] = r.Uint32()
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for _, a := range input {
			Popcount32(a)
		}
	}
}
