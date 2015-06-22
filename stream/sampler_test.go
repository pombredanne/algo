package stream

import (
	"math/rand"
	"testing"
)

func TestSampler(t *testing.T) {
	test(t, 3, 4, 37)

	s1 := test(t, 10000, 101, 42)
	s2 := test(t, 10000, 101, 42)
	for i := range s1 {
		if x, y := s1[i], s2[i]; x != y {
			t.Errorf("%d != %d, despite equal random seeds", x, y)
		}
	}
}

func test(t *testing.T, n, k int, seed int64) []interface{} {
	r := rand.New(rand.NewSource(seed))
	sam := NewSampler(r, k)

	for i := 0; i < n; i++ {
		if sam.Nitems() != i {
			t.Errorf("expected Nitems() = %d, got %d", n, sam.Nitems())
		}
		sam.Add(i)
	}

	if sam.Size() != k {
		t.Errorf("expected Size() = %d, got %d", k, sam.Size())
	}
	if n < k {
		k = n
	}
	sample := sam.Sample()
	if len(sample) != k {
		t.Errorf("expected a sample of length %d, got %d", k, len(sample))
	}
	unique := make(map[int]bool)
	for _, x := range sample {
		unique[x.(int)] = true
	}
	if len(unique) != k {
		t.Error("duplicate items in sample")
	}
	return sample
}
