// Copyright 2013–2015 Lars Buitinck
//
// MIT-licensed. See the file LICENSE for details.

package stream

import "math/rand"

// Reservoir sampler. Can be used to take a random sample from a stream.
//
// A Sampler keeps track of a sample (without replacement) of the items that
// are presented to its Add method. At any time, it may be queried for a
// sample of these items.
type Sampler struct {
	n      int
	r      *rand.Rand
	sample []interface{}
}

// Construct a Sampler that using the PRNG r and produces a sample of k items.
func NewSampler(r *rand.Rand, k int) *Sampler {
	return &Sampler{
		r:      r,
		sample: make([]interface{}, 0, k),
	}
}

// Present a new item x to the sampler.
func (s *Sampler) Add(x interface{}) {
	// Vitter's algorithm R; also explained at DADS,
	// http://xlinux.nist.gov/dads/HTML/reservoirSampling.html.
	s.n++
	if len(s.sample) < cap(s.sample) {
		s.sample = append(s.sample, x)
		return
	} else if i := s.r.Intn(s.n); i < len(s.sample) {
		s.sample[i] = x
	}
}

// Returns the number of items seen so far.
func (s *Sampler) Nitems() int {
	return s.n
}

// Returns the current sample in the sampler.
//
// The return value is a view into an internal buffer of the Sampler and should
// not be modified. It has length at most s.Size(); if less, the Sampler has
// not yet seen s.Size() elements.
func (s *Sampler) Sample() []interface{} {
	return s.sample
}

// Sample size of s.
func (s *Sampler) Size() int {
	return cap(s.sample)
}
