// Copyright 2015 Lars Buitinck
//
// MIT-licensed. See the file LICENSE for details.

package sort

import "sort"

// Partially sort data, so that the first k elements are the k smallest,
// in sorted order.
//
// Average time complexity O(n + k log k).
func Partial(data sort.Interface, k int) {
	partialSort(data, k, 0, data.Len())
}

// Partial quicksort algorithm due to Martínez (2004),
// http://www.cs.upc.edu/~conrado/research/reports/ALCOMFT-TR-03-50.pdf
func partialSort(data sort.Interface, k, i, j int) {
	for j-i > 2 {
		p := medianOfThree(data, i, j)
		p = partition(data, i, j, p)
		if p < k-1 {
			partialSort(data, k, p+1, j)
		}
		j = p
	}
	if j-i == 2 && data.Less(i+1, i) {
		data.Swap(i, i+1)
	}
}

func medianOfThree(data sort.Interface, i, j int) int {
	mid := i + (j-i)/2
	j--
	if data.Less(j, i) {
		i, j = j, i
	}
	if data.Less(mid, i) {
		return i
	}
	if data.Less(j, mid) {
		return j
	}
	return mid
}

// Based on Bentley's qsort3 (Programming Pearls, 2000, p. 120).
func partition(data sort.Interface, i, j, p int) int {
	data.Swap(i, p)
	p = i

	for {
		i++
		for i <= j && data.Less(i, p) {
			i++
		}
		j--
		for data.Less(p, j) {
			j--
		}

		if i > j {
			break
		}
		data.Swap(i, j)
	}
	data.Swap(p, j)
	return j
}
