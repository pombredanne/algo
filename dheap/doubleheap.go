// Package dheap implements double-ended heaps (min-max heaps).
package dheap

import "container/heap"

// Establishes heap order in h.
//
// Every non-empty heap must be initialized using this function before any of
// the other functions of this package is used on it.
func Init(h heap.Interface) {
	n := h.Len()
	for i := parent(n); i >= 0; i-- {
		siftDown(h, i, n)
	}
}

// Index of maximum element in h.
//
// There's no MinInd because the minimum of a heap is always at index zero.
func MaxInd(h heap.Interface) (maxind int) {
	switch h.Len() {
	case 1:
		maxind = 0
	case 2:
		maxind = 1
	default:
		maxind = 1
		if h.Less(1, 2) {
			maxind = 2
		}
	}
	return
}

// Removes and returns the maximum element of h.
func PopMax(h heap.Interface) interface{} {
	n := h.Len()
	if n <= 2 {
		return h.Pop()
	}

	i := 1
	if h.Less(1, 2) {
		i = 2
	}
	h.Swap(i, n-1)
	x := h.Pop()
	siftDownMax(h, i, n-1)
	return x
}

// Removes and returns the minimum element of h.
func PopMin(h heap.Interface) interface{} {
	n := h.Len() - 1
	h.Swap(0, n)
	x := h.Pop()
	siftDownMin(h, 0, n)
	return x
}

// Adds x to the heap x.
func Push(h heap.Interface, x interface{}) {
	h.Push(x)
	siftUp(h, h.Len()-1)
}

// Implementation follows Atkinson, Sack, Santoro and Strothotte (1986),
// Min-max heaps and generalized priority queues. CACM 29(10):996-1000.

func hasChild(i, n int) bool {
	return i*2+1 < n
}

func minLevel(i int) bool {
	lg := -1
	for i++; i > 0; {
		i >>= 1
		lg++
	}
	return (lg & 1) == 0
}

func parent(i int) int {
	return (i - 1) >> 1
}

func siftDown(h heap.Interface, i, n int) {
	if minLevel(i) {
		siftDownMin(h, i, n)
	} else {
		siftDownMax(h, i, n)
	}
}

// TODO implement iterative versions, described at:
// http://www.diku.dk/forskning/performance-engineering/Jesper/heaplab/heapsurvey_html/node11.html
func siftDownMax(h heap.Interface, i, n int) {
	if hasChild(i, n) {
		m := maxTwoGen(h, i, n)
		if m > i*2+2 { // must be a grandchild
			if h.Less(i, m) {
				h.Swap(i, m)
				if h.Less(m, parent(m)) {
					h.Swap(m, parent(m))
				}
				siftDownMax(h, m, n)
			}
		} else if h.Less(i, m) {
			h.Swap(i, m)
		}
	}
}

func siftDownMin(h heap.Interface, i, n int) {
	if hasChild(i, n) {
		m := minTwoGen(h, i, n)
		if m > i*2+2 { // must be a grandchild
			if h.Less(m, i) {
				h.Swap(i, m)
				if h.Less(parent(m), m) {
					h.Swap(m, parent(m))
				}
				siftDownMin(h, m, n)
			}
		} else if h.Less(m, i) {
			h.Swap(i, m)
		}
	}
}

// Index of maximum of children and grandchildren (if any) of i.
// Precondition: i has at least one child.
func maxTwoGen(h heap.Interface, i, n int) int {
	c1, c2 := i*2+1, i*2+2
	m := c1
	for _, k := range []int{c2, c1*2 + 1, c1*2 + 2, c2*2 + 1, c2*2 + 2} {
		if k >= n {
			break
		}
		if h.Less(m, k) {
			m = k
		}
	}
	return m
}

// Index of minimum of children and grandchildren (if any) of i.
// Precondition: i has at least one child.
func minTwoGen(h heap.Interface, i, n int) int {
	c1, c2 := i*2+1, i*2+2
	m := c1
	for _, k := range []int{c2, c1*2 + 1, c1*2 + 2, c2*2 + 1, c2*2 + 2} {
		if k >= n {
			break
		}
		if h.Less(k, m) {
			m = k
		}
	}
	return m
}

func siftUp(h heap.Interface, i int) {
	if minLevel(i) {
		if i > 0 && h.Less(parent(i), i) {
			h.Swap(i, parent(i))
			siftUpMax(h, parent(i))
		} else {
			siftUpMin(h, i)
		}
	} else {
		if i > 0 && h.Less(i, parent(i)) {
			h.Swap(i, parent(i))
			siftUpMin(h, parent(i))
		} else {
			siftUpMax(h, i)
		}
	}
}

func siftUpMax(h heap.Interface, i int) {
	for gp := parent(parent(i)); gp >= 0; i, gp = gp, parent(parent(gp)) {
		if h.Less(i, gp) {
			break
		}
		h.Swap(i, gp)
	}
}

func siftUpMin(h heap.Interface, i int) {
	for gp := parent(parent(i)); gp >= 0; i, gp = gp, parent(parent(gp)) {
		if h.Less(gp, i) {
			break
		}
		h.Swap(i, gp)
	}
}
