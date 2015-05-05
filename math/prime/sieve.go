// Copyright 2015 Lars Buitinck
//
// MIT-licensed. See the file LICENSE for details.

package prime

// 32-bit prime sieve: generates all primes less than (1<<32).
type Sieve32 struct {
	// Maps composite numbers to one of their factors.
	factors map[uint64]uint64
	next    uint64
}

// Generates the next len(buf) primes.
// Returns a prefix of buf, filled with primes.
func (s *Sieve32) Next(buf []uint32) (primes []uint32) {
	// Sieve of Eratosthenes, inspired by David Eppstein's Python version at
	// https://code.activestate.com/recipes/117119-sieve-of-eratosthenes/
	// with the optimization by Alex Martelli. Further optimized by simple
	// wheel factorization: only check 6n-1 and 6n+1.

	i, n := 0, len(buf)
	primes = buf
	q, factors := s.next, s.factors

	if factors == nil && n > 0 {
		factors = make(map[uint64]uint64)
		s.factors = factors
		primes[0] = 2
		i++
	}
	if q == 0 && i < n {
		primes[i] = 3
		i++
		q = 5
	}

	for q < (1<<32-3) && i < n {
		factor, composite := factors[q]
		if composite {
			delete(factors, q)
			x := factor + q
			for {
				if _, composite = factors[x]; !composite {
					break
				}
				x += factor
			}
			factors[x] = factor
		} else {
			factors[q*q] = 2 * q
			primes[i] = uint32(q)
			i++
		}

		if (q+1)%6 == 0 {
			q += 2
		} else {
			q += 4
		}
	}
	s.next = q
	return primes[:i]
}
