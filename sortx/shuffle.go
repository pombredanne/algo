// Copyright 2015 Lars Buitinck
//
// MIT-licensed. See the file LICENSE for details.

package sortx

import (
	"math/rand"
	"sort"
)

// Randomly shuffle data.
//
// If r == nil, uses the default rand.Source of the math/rand package.
func Shuffle(data sort.Interface, r *rand.Rand) {
	if r == nil {
		r = rand.New(randpkg{})
	}
	n := data.Len()
	for i := 0; i < n-1; i++ {
		j := i + r.Intn(n-i)
		if i != j {
			data.Swap(i, j)
		}
	}
}

type randpkg struct{}

func (r randpkg) Int63() int64    { return rand.Int63() }
func (r randpkg) Seed(seed int64) { rand.Seed(seed) }
