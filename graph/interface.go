// Package graph implements generic graph algorithms.
package graph

// Simple adjacency list representation for graphs.
//
// The set of vertices represented by an AdjacencyList a is {0, ..., len(a)}.
// a[u] is the set of neighbors of vertex u. Ensuring uniqueness is left to
// the user.
type AdjacencyList [][]int

func (g AdjacencyList) Neighbors(u int) []int {
	return g[u]
}

func (g AdjacencyList) NVertices() int {
	return len(g)
}

func (g AdjacencyList) OutDegree(u int) int {
	return len(g[u])
}

var _ FiniteDirected = AdjacencyList{}
var _ FiniteDirected = &AdjacencyList{}

// Interface for general graphs.
//
// Vertices (nodes) are represented as non-negative integers. The set of
// vertices in the graph is implicit; all that matters is that, given a vertex,
// we can find its neighbors.
//
// This interface is agnostic wrt. directedness of the graph. Derived
// interfaces may place additional constraints on the semantics of Neighbors.
type Graph interface {
	// The neighbors of u. Order doesn't matter; consumer algorithms will not
	// modify the returned slice.
	Neighbors(u int) []int
}

// Interface for directed graphs.
type Directed interface {
	Graph

	// Out-degree of u; must match the length of Neighbors(u).
	OutDegree(u int) int
}

// Interface for directed graphs of finite/known size.
type FiniteDirected interface {
	Directed

	// The number of vertices in the graph.
	NVertices() int
}
