package prime

import "testing"

var first = [...]uint32{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47}

func TestSieve(t *testing.T) {
	var s Sieve32

	for _, k := range []int{1, 2, 5, 8, len(first)} {
		s = Sieve32{} // reset
		got := make([]uint32, k)
		s.Next(got)
		for i := range got {
			if got[i] != first[i] {
				t.Errorf("%d'th prime is %d, got %d", i, first[i], got[i])
			}
		}
	}
}

func BenchmarkSieve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var s Sieve32
		pr := make([]uint32, 10)
		for j := 0; j < 2000; j++ {
			s.Next(pr)
		}
	}
}
