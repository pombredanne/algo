package intmath

import (
	"math/rand"
	"testing"
)

func TestLog2(t *testing.T) {
	for _, c := range []struct{ n, log int }{
		{0, -1}, {1, 0}, {1 << 16, 16}, {1<<4 - 1, 3}, {1<<5 + 3, 5},
	} {
		if got := Log2(c.n); got != c.log {
			t.Errorf("expected %d, got %d", c.log, got)
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
