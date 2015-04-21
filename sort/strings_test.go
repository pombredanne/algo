package sort

import (
	"math/rand"
	"strconv"
	"testing"
)

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
