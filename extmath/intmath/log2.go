// Copyright 2015 Lars Buitinck
//
// MIT-licensed. See the file LICENSE for details.

// +build !amd64

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
