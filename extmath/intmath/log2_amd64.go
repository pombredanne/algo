// Copyright 2015 Lars Buitinck
//
// MIT-licensed. See the file LICENSE for details.

package intmath

// Integer base-2 logarithm: floor(log2(v)).
//
// Returns -1 for negative n.
func Log2(v int) int {
	// https://graphics.stanford.edu/~seander/bithacks.html#IntegerLog
	if v <= 0 {
		return -1
	}
	r := 0
	if v & 0x7FFFFFFF00000000 != 0 {
		r |= 32
		v >>= 32
	}
	if v & 0xFFFF0000 != 0 {
		r |= 16
		v >>= 16
	}
	if v & 0xFF00 != 0 {
		r |= 8
		v >>= 8
	}
	if v & 0xF0 != 0 {
		r |= 4
		v >>= 4
	}
	if v & 0xC != 0 {
		r |= 2
		v >>= 2
	}
	if v & 0x2 != 0 {
		r |= 1
		v >>= 1
	}
	return r
}
