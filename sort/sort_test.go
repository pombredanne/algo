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
	data := sort.StringSlice(randomStrings(10000))

	for i := 0; i < b.N; i++ {
		a := make(sort.StringSlice, len(data))
		copy(a, data)

		b.StartTimer()
		Partial(a, 100)
		//sort.Strings(a)
		b.StopTimer()
	}

	Partial(data, 100)
	if !sort.IsSorted(data[:100]) {
		b.Fatal("not sorted: %v", data[:100])
	}
}

func TestSelect(t *testing.T) {
	a := sort.IntSlice{0, 5, 4, 1, 2, 9, 3, 8, 6, 7}
	for _, k := range []int{6, 9, 1, 0} {
		Select(a, k)
		if a[k] != k {
			t.Errorf("expected %d, got %d", k, a[k])
		}
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
