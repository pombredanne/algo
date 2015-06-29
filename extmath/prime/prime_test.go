package prime

import "testing"

func TestSieve(t *testing.T) {
	var s Sieve32
	primes100k := 0
	for i := 0; i < 10; i++ {
		primes := s.Next(nil)
		for _, p := range primes {
			if div := divisor(p); div != p {
				t.Errorf("%d %% %d == 0", p, div)
			}
			if p < 100000 {
				primes100k++
			}
		}
	}

	if primes100k != 9592 {
		t.Errorf("expected 9592 primes below 100,000, got %d", primes100k)
	}
}

func TestBigPrimes(t *testing.T) {
	for _, p := range segment(1<<32 - 1<<16, nil) {
		if divisor(p) != p {
			t.Errorf("not a prime: %d", p)
		}
	}
}

func divisor(p uint32) uint32 {
	if p&1 == 0 {
		return 2
	}
	for div := uint64(3); div*div <= uint64(p); div += 2 {
		if uint64(p)%div == 0 {
			return uint32(div)
		}
	}
	return p
}

func BenchmarkSieve(b *testing.B) {
	var nprimes int
	for i := 0; i < b.N; i++ {
		var s Sieve32
		var primes []uint32 = nil
		nprimes = 0
		for j := 0; j < 200; j++ {
			primes = s.Next(primes)
			nprimes += len(primes)
		}
	}
	b.Logf("got %d primes", nprimes)
}
