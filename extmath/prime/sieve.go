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

// TODO publish this
// Return the primes between lo*s.Blocksize() and (lo+1)*s.Blocksize().
//
// Multiple goroutines can call Block concurrently, to make a parallel
// prime sieve.
//
// buf is a buffer of memory that may be reused to store the result. It may be
// nil.
func segment(lo uint32, buf []uint32) (primes []uint32) {
	// Segmented sieve of Eratosthenes, after Sorenson (1990),
	// https://research.cs.wisc.edu/techreports/1990/TR909.pdf
	if lo == 0 {
		if cap(buf) >= len(primes16) {
			primes = buf
		} else {
			primes = make([]uint32, len(primes16))
		}

		for i, p := range primes16 {
			primes[i] = uint32(p)
		}
		return
	}

	// TODO Only store even numbers.
	composite := [block]bool{}
	for _, p16 := range primes16[1:] {
		p, lo := uint64(p16), uint64(lo)
		// Formula for first multiple from Sorenson.
		for mult := p + lo - (lo % p); mult < lo+block; mult += p {
			composite[mult-lo] = true
		}
	}

	primes = buf[:0]
	for i := 1; i < len(composite); i += 2 {
		if !composite[i] {
			primes = append(primes, lo+uint32(i))
		}
	}
	return
}

// Generates the next block of primes.
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
	primes = segment(s.lo, buf)
	s.lo += block
	if s.lo == 0 {
		s.lo = 1<<32 - 1
	}
	return
}
