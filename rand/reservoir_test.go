package rand

import (
	"math/rand"
	"testing"
)

func TestSampleStream(t *testing.T) {
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
	ch := make(chan interface{})
	go func() {
		for i := 0; i < n; i++ {
			ch <- i
		}
		close(ch)
	}()
	sample := SampleStream(r, ch, k)

	if n < k {
		k = n
	}
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
