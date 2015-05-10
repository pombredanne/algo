package sort

import "sort"

// Generates the lexicographically next multiset permutation, in-place.
//
// Reports next when there is no next permutation, i.e., data is reverse-sorted.
func NextPermutation(data sort.Interface) bool {
	n := data.Len()
	var k, l int
	for k = n-2; k >= 0 && !data.Less(k, k+1); k-- { }
	if k < 0 {
		return false
	}
	for l = k + 1; l < n-1; l++ {
		if data.Less(l+1, k) {
			break
		}
	}
	data.Swap(k, l)
	reverse(data, k+1, n)
	return true
}

func reverse(data sort.Interface, i, j int) {
	for j--; i < j; i, j = i+1, j-1 {
		data.Swap(i, j)
	}
}
