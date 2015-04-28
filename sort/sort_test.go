package sort

import (
	"math/rand"
	"sort"
	"strconv"
	"testing"
)

func TestPartial(t *testing.T) {
	a := sort.IntSlice{0, 5, 4, 1, 2, 9, 3, 8, 6, 7}
	for _, k := range []int{1, 3, 5, 10} {
		Partial(a, k)
		for i := 0; i < k; i++ {
			if a[i] != i {
				t.Errorf("expected %d, got %d", i, a[i])
			}
		}
	}
}

func BenchmarkPartial(b *testing.B) {
	b.StopTimer()
	a := sort.StringSlice(randomStrings(10000))

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		Partial(a, 100)
		//sort.Strings(a)
	}
	b.StopTimer()

	if !sort.IsSorted(a[:100]) {
		b.Fatal("not sorted: %v", a[:100])
	}
}

func TestStrings(t *testing.T) {
	data := randomStrings(114)
	Strings(data)
	for i := 1; i < len(data); i++ {
		if data[i-1] > data[i] {
			t.Errorf("%q > %q", data[i-1], data[i])
		}
	}
}

func BenchmarkStrings(b *testing.B) {
	b.StopTimer()
	data := randomStrings(10000)

	for i := 0; i < b.N; i++ {
		a := make([]string, len(data))
		copy(a, data)

		b.StartTimer()
		Strings(a)
		//sort.Strings(data)
		b.StopTimer()
	}
}

func randomStrings(n int) []string {
	rng := rand.New(rand.NewSource(42))
	strs := make([]string, n)
	for i := 0; i < n; i++ {
		// Rather cheap solution
		strs[i] = strconv.Itoa(int(rng.Int31()))
	}
	return strs
}
