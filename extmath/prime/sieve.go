// Copyright 2015 Lars Buitinck
//
// MIT-licensed. See the file LICENSE for details.

// Prime number generation.
package prime

const block = 1 << 16

// 32-bit prime sieve: generates all primes less than (1<<32).
type Sieve32 struct {
	lo uint32
}

/*
func New32(lo uint32) *Sieve32 {
	return &Sieve32{lo: lo}
}
*/

// Generates a new block of primes.
//
// The size of the block is not guaranteed.
//
// buf is a buffer of memory that may be reused to store the result. It may be
// nil.
//
// Example usage:
//
//	var primes []uint32 = nil
//	for {
//		primes = sieve.Next(primes)
//		for _, p := range primes {
//			if p > maxwanted {
//				break
//			}
//			// process p
//		}
//	}
func (s *Sieve32) Next(buf []uint32) (primes []uint32) {
	// Segmented sieve of Eratosthenes, after Sorenson (1990),
	// https://research.cs.wisc.edu/techreports/1990/TR909.pdf
	if s.lo == 0 {
		if cap(buf) >= len(primes16) {
			primes = buf
		} else {
			primes = make([]uint32, len(primes16))
		}

		for i, p := range primes16 {
			primes[i] = uint32(p)
		}
		s.lo = block
		return
	} else if s.lo == 1<<32-1 {
		return
	}

	lo := s.lo
	// TODO Only process even numbers.
	composite := [block]bool{}
	composite[0] = true
	for _, p16 := range primes16 {
		p := uint32(p16)
		// Formula for first multiple from Sorenson.
		for mult := p + lo - (lo % p); mult < lo+block; mult += p {
			composite[mult-lo] = true
		}
	}

	primes = buf[:0]
	for i, c := range composite {
		if !c {
			primes = append(primes, lo+uint32(i))
		}
	}
	s.lo += block
	if s.lo == 0 {
		s.lo = 1<<32 - 1
	}
	return
}
