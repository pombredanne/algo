// Copyright 2015 Lars Buitinck
//
// MIT-licensed. See the file LICENSE for details.

package sortx

import "sort"

// Partially sort data, so that the first k elements are the k smallest,
// in sorted order.
//
// Average time complexity O(n + k log k).
func Partial(data sort.Interface, k int) {
	partialSort(data, k, 0, data.Len())
}

// Partition data so that the k'th smallest element is at position k,
// the first k are smaller than the k'th smallest and the rest are larger
// than it.
//
// Average time complexity O(n).
func Select(data sort.Interface, k int) {
	quickselect(data, k, 0, data.Len())
}

// TODO: implement introselect.
func quickselect(data sort.Interface, k, i, j int) {
	for j-i > 2 {
		p := medianOfThree(data, i, j)
		p = partition(data, i, j, p)
		if k == p {
			return
		} else if k < p {
			j = p
		} else {
			i = p + 1
		}
	}
	if j-i == 2 && data.Less(i+1, i) {
		data.Swap(i, i+1)
	}
}

// Partial quicksort algorithm due to Martínez (2004),
// http://www.cs.upc.edu/~conrado/research/reports/ALCOMFT-TR-03-50.pdf
func partialSort(data sort.Interface, k, lo, hi int) {
	for hi-lo > 5 {
		p := medianOfThree(data, lo, hi)
		p = partition(data, lo, hi, p)
		if p < k-1 {
			partialSort(data, k, p+1, hi)
		}
		hi = p
	}

	// Finish off with a selection sort.
	if hi-lo-1 < k {
		k = hi - lo - 1
	}
	for ; k > 0; k-- {
		min := lo
		for i := lo + 1; i < hi; i++ {
			if data.Less(i, min) {
				min = i
			}
		}
		data.Swap(lo, min)
		lo++
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
