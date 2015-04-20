package doubleheap

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

	perm = makePerm(13)
	for i := 0; i < 13; i++ {
		if max := PopMax(&perm); max != 13-i-1 {
			t.Errorf("expected %d, got %d", 13-i-1, max)
		}
	}
}

func TestPush(t *testing.T) {
	h := ints{0}
	min, max := 0, 0

	for _, x := range []int{7, 14, 15, -1, 2, 2, -2, 100} {
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
