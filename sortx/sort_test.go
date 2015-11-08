package sortx

import (
	"math/rand"
	"sort"
	"strconv"
	"testing"
)

func TestPartial(t *testing.T) {
	r := rand.New(rand.NewSource(42))
	for k := 0; k < 10; k++ {
		a := sort.IntSlice(r.Perm(10))
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
		b.Fatalf("not sorted: %v", data[:100])
	}
}

func TestSelect(t *testing.T) {
	r := rand.New(rand.NewSource(42))
	for k := 0; k < 16; k++ {
		a := sort.IntSlice(r.Perm(16))
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
	data := randomStrings(100000)

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

func shuffle(a []float64, r *rand.Rand) {
	for i := len(a)-1; i > 0; i-- {
		j := r.Intn(i+1)
		a[i], a[j] = a[j], a[i]
	}
}
