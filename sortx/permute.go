// Copyright 2015 Lars Buitinck
//
// MIT-licensed. See the file LICENSE for details.

package sortx

import "sort"

// Generates the lexicographically next multiset permutation, in-place.
//
// Returns false when there is no next permutation, i.e., data is
// reverse-sorted.
func NextPermutation(data sort.Interface) bool {
	n := data.Len()
	var k, l int
	for k = n - 2; k >= 0 && !data.Less(k, k+1); k-- {
	}
	if k < 0 {
		return false
	}
	for l = k + 1; l < n-1; l++ {
		if data.Less(l+1, k) {
			break
		}
	}
	data.Swap(k, l)
	reverse(data, k+1, n)
	return true
}

// Generates the lexicographically previous multiset permutation, in-place.
//
// Returns false when there is no previous permutation, i.e., data is sorted.
func PrevPermutation(data sort.Interface) bool {
	n := data.Len()
	var k, l int
	for k = n - 2; k >= 0 && !data.Less(k+1, k); k-- {
	}
	if k < 0 {
		return false
	}
	for l = k + 1; l < n-1; l++ {
		if data.Less(k, l+1) {
			break
		}
	}
	data.Swap(k, l)
	reverse(data, k+1, n)
	return true
}

func reverse(data sort.Interface, i, j int) {
	for j--; i < j; i, j = i+1, j-1 {
		data.Swap(i, j)
	}
}
