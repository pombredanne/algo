package extsort

import (
	"sort"
	"testing"
)

func TestPermute(t *testing.T) {
	a := sort.IntSlice{1, 2, 3, 4, 5}
	nperm := 1
	for NextPermutation(a) {
		nperm++
	}
	if nperm != 120 {
		t.Errorf("%v has 120 permutations, got %d", a, nperm)
	}

	// Multiset permutations.
	a = sort.IntSlice{1, 2, 2}
	for _, p := range [][]int{{2, 1, 2}, {2, 2, 1}} {
		NextPermutation(a)
		for i := range a {
			if a[i] != p[i] {
				t.Errorf("wanted %v, got %v", p, a)
			}
		}
	}
	if NextPermutation(a) {
		t.Errorf(`array was reverse-sorted, but got "next permutation" %v`,
			a)
	}

	// Corner cases.
	for _, a = range []sort.IntSlice{{}, {6}} {
		if NextPermutation(a) {
			t.Errorf(`got "next permutation" for length-%d array`, len(a))
		}
	}
}
