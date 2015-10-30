package sortx

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
	for PrevPermutation(a) {
		nperm--
	}
	if nperm != 1 {
		t.Error("PrevPermutation stopped after %d steps", 120-nperm)
	}
	if !sort.IsSorted(a) {
		t.Error("expected data to be sorted, got %v", a)
	}

	// Multiset permutations.
	a = sort.IntSlice{1, 2, 2}
	perms := [][]int{{1, 2, 2}, {2, 1, 2}, {2, 2, 1}}
	for _, p := range perms[1:] {
		ok := NextPermutation(a)
		if !ok {
			t.Errorf("NextPermutation returned false at %v", a)
		}
		for i := range a {
			if a[i] != p[i] {
				t.Errorf("wanted %v, got %v", p, a)
			}
		}
	}
	if NextPermutation(a) {
		t.Errorf(`array was reverse-sorted, but got "next permutation" %v`, a)
	}
	for j := 0; j < len(perms)-1; j++ {
		p := perms[len(perms)-j-1]
		for i := range a {
			if a[i] != p[i] {
				t.Errorf("wanted %v, got %v", p, a)
			}
		}
		ok := PrevPermutation(a)
		if !ok {
			t.Errorf("PrevPermutation returned false at %v", a)
		}
	}

	// Corner cases.
	for _, a = range []sort.IntSlice{{}, {6}} {
		if NextPermutation(a) {
			t.Errorf(`got "next permutation" for length-%d array`, len(a))
		}
	}
}

func BenchmarkPrev(b *testing.B) {
	a := make(sort.IntSlice, 1000)
	for i := range a {
		a[i] = -i
	}

	for i := 0; i < b.N; i++ {
		PrevPermutation(a)
	}
}
