// Copyright 2015 Lars Buitinck
//
// MIT-licensed. See the file LICENSE for details.

// Integer math, i.e., bit fiddling.
package intmath

// Integer base-2 logarithm: floor(log2(v)).
//
// Returns -1 for negative n.
func Log2(v int) int {
	r := -1
	for v > 0 {
		v >>= 1
		r++
	}
	return r
}

// Number of bits set in v.
func Popcount32(v uint32) (c int) {
	// https://graphics.stanford.edu/~seander/bithacks.html#CountBitsSetKernighan
	for c = 0; v != 0; c++ {
		v &= v - 1
	}
	return
}
