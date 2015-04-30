// Copyright 2013–2015 Lars Buitinck
//
// MIT-licensed. See the file LICENSE for details.

// Package disjsets implements disjoint-set forests (union-find structures).
package disjsets

type set struct {
	parent int
	rank   uint8 // Max. rank is log(N); assume N < 2**256
}

// A disjoint-set forest. Sets are represented as integer indices in the
// range [0, n) where n in the number of elements in the set.
type Forest struct {
	sets  []set
	nsets int
}

// Construct a new disjoint-set forest of n elements in n singleton sets.
func New(n int) *Forest {
	sets := make([]set, n)
	for i := range sets {
		sets[i].parent = i
		sets[i].rank = 0
	}
	return &Forest{sets, n}
}

// Find the representative of the set that x belongs to.
//
// Note: this function may modify the forest.
func (forest *Forest) Find(x int) int {
	n := &forest.sets[x]
	if n.parent != x {
		n.parent = forest.Find(n.parent)
	}
	return n.parent
}

// Reports the number of elements in the forest.
func (forest *Forest) Len() int {
	return len(forest.sets)
}

// Reports the number of disjoint sets in the forest.
func (forest *Forest) NSets() int {
	return forest.nsets
}

// Merge the sets that x and y belong to. Returns true if a merger occurred,
// false if x and y were already in the same set.
func (forest *Forest) Union(x, y int) bool {
	xrootidx := forest.Find(x)
	yrootidx := forest.Find(y)

	if xrootidx == yrootidx {
		return false
	}

	xroot := &forest.sets[xrootidx]
	yroot := &forest.sets[yrootidx]

	if xroot.rank < yroot.rank {
		xroot.parent = yrootidx
	} else if xroot.rank > yroot.rank {
		yroot.parent = xrootidx
	} else {
		yroot.parent = xrootidx
		xroot.rank++
	}
	forest.nsets--
	return true
}
