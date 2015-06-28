package prime

import "testing"

func TestSieve(t *testing.T) {
	var s Sieve32
	for i := 0; i < 10; i++ {
		primes := s.Next(nil)
		for _, p := range primes {
			if div := divider(p); div != p {
				t.Errorf("%d %% %d == 0", p, div)
			}
		}
	}
}

func divider(p uint32) uint32 {
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
	b.Log(nprimes)
}
