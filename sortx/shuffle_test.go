package sortx

import (
	"math/rand"
	"sort"
	"testing"
)

func TestShuffle(t *testing.T) {
	dist := make([]int, 5*5)
	nrounds := 1000
	seeds := []int64{51, 512, 6612, 1, 6, 62623, 0x5216, 0xAA}

	for _, seed := range seeds {
		r := rand.New(rand.NewSource(seed))
		for i := 0; i < nrounds; i++ {
			a := sort.IntSlice{0, 1, 2, 3, 4}
			Shuffle(a, r)
			for i, j := range a {
				dist[i*5+j]++
			}
		}
	}

	// Check whether the distribution of permutations behaves as expected.
	expected := float64(len(seeds)*nrounds) / 5
	for i, n := range dist {
		if float64(n) < .95*expected || float64(n) > 1.05*expected {
			t.Logf("(%d, %d): %d too far form expectation %f\n",
				i/5, i%5, n, expected)
		}
	}
}

func TestShuffleNil(t *testing.T) {
	a := sort.StringSlice{"foo", "bar", "baz", "quux"}
	b := make(sort.StringSlice, len(a))
	copy(b, a)

	rand.Seed(42)
	Shuffle(a, nil)
	rand.Seed(42)
	Shuffle(b, nil)
	for i := range a {
		if a[i] != b[i] {
			t.Errorf("mismatch at %d: %q != %q", i, a[i], b[i])
		}
	}
}
