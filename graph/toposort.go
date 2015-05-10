// Copyright 2015 Lars Buitinck
//
// MIT-licensed. See the file LICENSE for details.

package graph

import "errors"

// Topological sorting of vertices in a directed acyclic graph.
//
// Returns labels (indices) for the vertices of g in a topological order,
// or nil and an error if g is not a directed acyclic graph.
//
// To actually get the vertices topologically ordered in an array v, do
// for i, l := range labels { v[l] = i }
func TopoSort(g FiniteDirected) (labels []int, err error) {
	n := g.NVertices()
	labels = make([]int, n)

	ordnum := n
	flags := make([]struct{ onstack, visited bool }, n)

	var visit func(int)
	visit = func(u int) {
		if flags[u].onstack {
			err = errors.New("directed cycle in input to TopoSort")
			return
		}
		if flags[u].visited {
			return
		}
		flags[u].onstack = true
		for _, v := range g.Neighbors(u) {
			if visit(v); err != nil {
				return
			}
		}
		flags[u].visited = true
		flags[u].onstack = false
		ordnum--
		labels[u] = ordnum
	}

	for u := 0; u < n; u++ {
		if visit(u); err != nil {
			return nil, err
		}
	}
	return
}
