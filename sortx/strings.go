// Copyright 2015 Lars Buitinck
//
// MIT-licensed. See the file LICENSE for details.

// Specialized sorting algorithms.
package sortx

import "github.com/larsmans/algo/extmath/intmath"

// Strings sorts a slice of strings in increasing order (byte-wise
// lexicographically).
//
// This function is equivalent to sort.Strings, but faster than the
// implementation in Go 1.4.
func Strings(a []string) {
	n := len(a)
	radixQuicksort(a, 0, 0, n, 2*intmath.Log2(n+1))
}

// XXX The following could be generalized by using an interface like
// type StringKeys interface {
//	Key(int) string
// }

// Three-way radix quicksort (aka. multi-key quicksort), after Bentley and
// Sedgewick (1997). Fast Algorithms for Sorting and Searching Strings.
// Proc. SODA, http://www.cs.princeton.edu/~rs/strings/paper.pdf. Also
// http://www.drdobbs.com/database/sorting-strings-with-three-way-radix-qui/184410724.
func radixQuicksort(data []string, index, a, b, depth int) {
	for b-a > 1 && depth > 0 {
		pivot := medianOfThreeBytes(
			char(data[a], index),
			char(data[a+(b-a)/2], index),
			char(data[b-1], index))

		lo, hi := a, b
		for i := lo; i < hi; {
			t := char(data[i], index)
			if t < pivot {
				data[lo], data[i] = data[i], data[lo]
				lo++
				i++
			} else if t > pivot {
				hi--
				data[i], data[hi] = data[hi], data[i]
			} else {
				i++
			}
		}
		depth--
		radixQuicksort(data, index, a, lo, depth)
		radixQuicksort(data, index, hi, b, depth)
		if pivot == -1 {
			return
		}
		a, b = lo, hi
		index++
	}
}

func medianOfThreeBytes(a, b, c int) int {
	if a > c {
		a, c = c, a
	}
	if a > b {
		return a
	}
	if b > c {
		return c
	}
	return b
}

func char(s string, i int) int {
	if i == len(s) { // i > len(s) shouldn't happen, so let it panic.
		return -1
	}
	return int(s[i])
}
