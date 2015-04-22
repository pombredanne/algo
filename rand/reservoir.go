// Copyright 2013–2015 Lars Buitinck
//
// MIT-licensed. See the file LICENSE for details.

package rand

import "math/rand"

// Randomly sample k items from the channel ch (reservoir sampling).
//
// Consumes all items arriving on ch. If less than k items arrive, they are
// all returned.
func SampleStream(r *rand.Rand, ch <-chan interface{}, k int) []interface{} {
	// Vitter's algorithm R; also explained at DADS,
	// http://xlinux.nist.gov/dads/HTML/reservoirSampling.html.
	sample := make([]interface{}, 0, k)
	j := 0
	for ; j < k; j++ {
		x, ok := <-ch
		if !ok {
			return sample
		}
		sample = append(sample, x)
	}
	for x := range ch {
		j++
		i := r.Intn(j)
		if i < k {
			sample[i] = x
		}
	}
	return sample
}
