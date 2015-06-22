// Copyright 2013–2015 Lars Buitinck
//
// MIT-licensed. See the file LICENSE for details.

// Deprecated. Use the Sampler from github.com/larsmans/algo/stream instead.
package rand

import (
	"github.com/larsmans/algo/stream"
	"math/rand"
)

// Randomly sample k items from the channel ch (reservoir sampling).
//
// Consumes all items arriving on ch. If less than k items arrive, they are
// all returned.
func SampleStream(r *rand.Rand, ch <-chan interface{}, k int) []interface{} {
	s := stream.NewSampler(r, k)
	for x := range ch {
		s.Add(x)
	}
	return s.Sample()
}
