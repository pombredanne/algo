// Copyright 2015 Lars Buitinck
//
// MIT-licensed. See the file LICENSE for details.

// Graph search and traversal.
package traverse

import "github.com/larsmans/algo/graph"

// Traverses g in breadth-first order, starting at the given start vertices and
// calling callback for each vertex.
//
// The arguments for the callback function are the "from" and the "to"
// vertices; the former is -1 for the start nodes. The callback is called
// exactly once for each "to" vertex in the graph. Traversal stops when the
// callback returns a non-nil error; the error is propagated.
func BreadthFirst(g graph.Directed, callback func(from, to int) error,
	start ...int) error {

	closed := make(map[int]bool)

	queue, prev := start, -1
	for len(queue) != 0 {
		u := queue[0]
		queue = queue[1:]

		if closed[u] {
			continue
		}
		if err := callback(prev, u); err != nil {
			return err
		}
		closed[u] = true

		for _, v := range g.Neighbors(u) {
			if !closed[v] {
				queue = append(queue, v)
			}
		}
		prev = u
	}
	return nil
}

// Traverses g in breadth-first order, starting at the given start vertex and
// calling callback for each vertex.
//
// Compared to BreadthFirst, this function uses less memory but may take more
// time. It also supports only a single start node.
func IterativeDeepening(g graph.Directed, callback func(from, to int) error,
	start int) (err error) {

	closed := make(map[int]bool)

	// Reports whether the callback got called during traversal (so we've seen
	// a new part of the graph), and if so, what it returned.
	var depthLimDF func(int, int, int) (bool, error)
	depthLimDF = func(from, u, depth int) (called bool, err error) {
		if depth == 0 && !closed[u] {
			if err = callback(from, u); err != nil {
				return true, err
			}
			called = true
			closed[u] = true
		} else if depth > 0 {
			for _, v := range g.Neighbors(u) {
				var calledV bool
				calledV, err = depthLimDF(u, v, depth-1)
				if calledV {
					called = true
				}
				if err != nil {
					break
				}
			}
		}
		return
	}

	for depth := 0; ; depth++ {
		var morenodes bool
		morenodes, err = depthLimDF(-1, start, depth)
		if !morenodes || err != nil {
			break
		}
	}
	return
}
