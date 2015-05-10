package sort

import (
	"math/rand"
	"sort"
	"strconv"
	"testing"
)

func TestPartial(t *testing.T) {
	r := rand.New(rand.NewSource(42))
	a := sort.IntSlice{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for k := range a {
		shuffle(a, r)
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
	r := rand.New(rand.NewSource(42))
	a := sort.IntSlice{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	for k := range a {
		shuffle(a, r)
		Select(a, k)
		if a[k] != k {
			t.Errorf("expected %d, got %d", k, a[k])
		}

		// Test partitioning around the k'th smallest element.
		sort.Sort(a[:k])
		sort.Sort(a[k+1:])
		if !sort.IsSorted(a) {
			t.Error("data not correctly partitioned")
		}
	}
}

func BenchmarkSelect(b *testing.B) {
	b.StopTimer()
	a := make(sort.Float64Slice, 100000)
	for i := range a {
		a[i] = float64(i)
	}
	r := rand.New(rand.NewSource(42))

	for i := 0; i < b.N; i++ {
		shuffle(a, r)
		b.StartTimer()
		for _, k := range []int{5, 166, 900, 126, 0} {
			Select(a, k)
			//Partial(a, k+1)
			if a[k] != float64(k) {
				b.Fatalf("expected %d, got %f", k, a[k])
			}
		}
		b.StopTimer()
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
