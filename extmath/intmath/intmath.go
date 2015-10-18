// Copyright 2015 Lars Buitinck
//
// MIT-licensed. See the file LICENSE for details.

// Integer math, i.e., bit fiddling.
package intmath

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

// Number of bits set in v.
func Popcount64(v uint64) (c int) {
	// https://graphics.stanford.edu/~seander/bithacks.html#CountBitsSetKern
	for c = 0; v != 0; c++ {
		v &= v - 1
	}
	return
}
