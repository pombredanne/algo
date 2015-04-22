package dheap

import (
	"math/rand"
	"testing"
)

type ints []int

func (a ints) Len() int            { return len(a) }
func (a ints) Less(i, j int) bool  { return a[i] < a[j] }
func (a *ints) Push(x interface{}) { *a = append(*a, x.(int)) }
func (a ints) Swap(i, j int)       { a[i], a[j] = a[j], a[i] }

func (a *ints) Pop() interface{} {
	b := *a
	x := b[len(b)-1]
	*a = b[:len(b)-1]
	return x
}

func makePerm(n int) ints {
	p := ints(rand.New(rand.NewSource(42)).Perm(n))
	Init(&p)
	return p
}

func TestPop(t *testing.T) {
	perm := makePerm(10)
	for i := 0; i < 10; i++ {
		if min := PopMin(&perm); min != i {
			t.Errorf("expected %d, got %d", i, min)
		}
	}

	perm = makePerm(15)
	for i := 0; i < 15; i++ {
		atMaxind := perm[MaxInd(&perm)]
		if max := PopMax(&perm); max != 15-i-1 {
			t.Errorf("expected %d, got %d", 15-i-1, max)
			if max != atMaxind {
				t.Errorf("expected %d at MaxInd, found %d", 15-i-1, atMaxind)
			}
		}
	}
}

func TestPush(t *testing.T) {
	h := ints{0}
	min, max := 0, 0

	for _, x := range []int{7, 14, 15, -1, 2, 2, -2, 100, 11, 11} {
		Push(&h, x)
		if x > max {
			max = x
		} else if x < min {
			min = x
		}

		if h[0] != min {
			t.Errorf("expected %d, got %d", min, h[0])
		}
		if got := h[MaxInd(&h)]; got != max {
			println(h[0], h[1], h[2])
			println(MaxInd(&h), h.Less(1, 2))
			t.Errorf("expected %d, got %d", max, got)
		}
	}
}

func BenchmarkHeap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var h ints
		for j := 0; j < 2000; j++ {
			Push(&h, j%99)
		}
		for j := 0; j < 10000; j++ {
			switch j % 5 {
			case 2:
				PopMax(&h)
			case 4:
				PopMin(&h)
			default:
				Push(&h, j%98)
			}
		}
	}
}
