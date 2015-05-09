// Copyright 2015 Lars Buitinck
//
// MIT-licensed. See the file LICENSE for details.

// Package bloom implements Bloom filters (approximate sets).
package bloom

import (
	"fmt"
	"github.com/larsmans/algo/math/intmath"
	"math"
	"math/rand"
)

// Bloom filter with 32-bit hash function.
type Filter32 struct {
	bits []uint32
	seed []uint32
}

// Construct new Bloom filter with given number of buckets and hashes.
//
// The random seed r is used to construct the hash functions.
// The number of buckets will be rounded up to a multiple of 32.
//
// Returns nil and an error if the number of buckets exceeds (1<<32)-1.
func New32(nbuckets, nhashes int, r *rand.Rand) (f *Filter32, err error) {
	// Round to next multiple of 32.
	nmod32 := nbuckets & ((1 << 5) - 1)
	if nmod32 != 0 {
		nbuckets += 32 - nmod32
	}

	if nbuckets > math.MaxUint32 {
		err = fmt.Errorf("nbuckets = %d (after rounding) too large", nbuckets)
		return
	}

	seeds := make([]uint32, nhashes)
	for i := range seeds {
		seeds[i] = r.Uint32()
	}

	f = &Filter32{make([]uint32, nbuckets>>5), seeds}
	return
}

// Add key with hash value h to the filter.
func (f *Filter32) Add(h uint32) {
	for _, s := range f.seed {
		hs := (h ^ s) % uint32(f.Capacity())

		idx, mask := indexMask(hs)
		f.bits[idx] |= mask
	}
}

// Reports whether a key with hash value h was added to the filter.
//
// May returns false positives.
func (f *Filter32) Get(h uint32) bool {
	for _, s := range f.seed {
		hs := (h ^ s) % uint32(f.Capacity())

		idx, mask := indexMask(hs)
		if f.bits[idx]&mask == 0 {
			return false
		}
	}
	return true
}

func indexMask(h uint32) (index, mask uint32) {
	// hi, lo := h/32, h%32
	hi, lo := h>>5, h&((1<<5)-1)
	return hi, 1 << lo
}

// The capacity of the filter (number of buckets).
func (f *Filter32) Capacity() int {
	return len(f.bits) << 5
}

// Clears (empties) the filter.
func (f *Filter32) Clear() {
	for i := range f.bits {
		f.bits[i] = 0
	}
}

// Returns an approximation of the number of distinct keys stored in f.
func (f *Filter32) NKeys() float64 {
	nset := 0
	for _, b := range f.bits {
		nset += intmath.Popcount32(b)
	}
	length := float64(f.Capacity())
	return -(length * math.Log(1-float64(nset)/length)) / float64(len(f.seed))
}
