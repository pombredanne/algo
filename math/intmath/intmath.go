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
func Popcount32(v uint32) int {
	// https://graphics.stanford.edu/~seander/bithacks.html#CountBitsSet64
	// (Assumes fast 64-bit instructions.)
	v64 := uint64(v)
	c := ((v64 & 0xfff) * 0x1001001001001 & 0x84210842108421) % 0x1f
	c += (((v64 & 0xfff000) >> 12) * 0x1001001001001 & 0x84210842108421) % 0x1f
	c += ((v64 >> 24) * 0x1001001001001 & 0x84210842108421) % 0x1f
	return int(c)
}
