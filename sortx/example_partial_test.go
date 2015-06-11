package sortx_test

import (
	"fmt"
	"github.com/larsmans/algo/sortx"
	"sort"
)

func ExamplePartial() {
	data := sort.IntSlice{9, 5, 1, 7, 4, 3, 11, 21, 2, 42, 37, 2, 8, 6}
	topK := 5 // We want the topK largest elements.
	sortx.Partial(sort.Reverse(data), topK)
	for _, x := range data[:topK] {
		fmt.Printf(" %d", x)
	}
	// Output:
	// 42 37 21 11 9
}
