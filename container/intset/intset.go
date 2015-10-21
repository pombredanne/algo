// Copyright 2015 Lars Buitinck
//
// MIT-licensed. See the file LICENSE for details.

// This implements the dynamic integer set data structure described by Briggs
// and Torczon (1993), "An Efficient Representation for Sparse Sets", ACM
// Letters on Prog. Lang. and Sys., and by Russ Cox at
// http://research.swtch.com/sparse. Various implementations of this structure
// exist in the Go source code and popular libraries, but none seem to have a
// public API.

// Package intset implements an integer set data structure.
package intset

type Set struct {
	p []pair
	n int
}

type pair struct{ dense, sparse int }

// Construct an empty set of non-negative integers.
//
// All future members must be < universe.
func New(universe int) *Set {
	return &Set{p: make([]pair, universe), n: 0}
}

// Add k to s.
func (s *Set) Add(k int) {
	a := s.p[k].sparse
	if a >= s.n || s.p[a].dense != k {
		s.p[k].sparse = s.n
		s.p[s.n].dense = k
		s.n++
	}
}

// Reports whether k is in s.
func (s *Set) Contains(k int) bool {
	a := s.p[k].sparse
	return a < s.n && s.p[a].dense == k
}

// Iterate over members of s, calling f for each of them.
//
// The order of iteration is arbitrary.
func (s *Set) Do(f func(int)) {
	for i := 0; i < s.n; i++ {
		f(s.p[i].dense)
	}
}

// Number of items in s.
func (s *Set) Len() int {
	return s.n
}

// Remove k from s. Succeeds even k is not contained in s.
func (s *Set) Remove(k int) {
	a := s.p[k].sparse
	n := s.n - 1
	if a <= n && s.p[a].dense == k {
		e := s.p[n].dense
		s.n = n
		s.p[a].dense = e
		s.p[e].sparse = a
	}
}
