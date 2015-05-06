// Package graph implements generic graph algorithms.
//
// This package does not offer graph data structures, relying instead on the
// client to implement its interfaces.
package graph

// Interface for directed graphs.
type Directed interface {
	// The neighbors of u. Order doesn't matter; consumer algorithms will not
	// modify the returned slice.
	Neighbors(u int) []int
}

// Interface for directed graphs of finite/known size.
type FiniteDirected interface {
	Directed

	// The number of vertices in the graph.
	NVertices() int
}
