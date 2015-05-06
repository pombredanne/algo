// Copyright 2015 Lars Buitinck
//
// MIT-licensed. See the file LICENSE for details.

// Graph search and traversal.
// The API of this package is likely to change in the near future.
package traverse

import "github.com/larsmans/algo/graph"

// Traverses g in breadth-first order, starting at the given start nodes and
// calling callback for each vertex.
//
// The arguments for the callback function are the "from" and the "to"
// vertices. The callback is not called for the start nodes, and is called
// exactly once for each "to" vertex. Traversal stops when the callback
// returns a non-nil error; the error is propagated.
func BreadthFirst(g graph.Directed, callback func(from, to int) error,
	start ...int) error {

	closed := make(map[int]bool)
	for _, u := range start {
		closed[u] = true
	}
	queue := start

	for len(queue) != 0 {
		u := queue[0]
		queue = queue[1:]

		for _, v := range g.Neighbors(u) {
			if !closed[v] {
				if err := callback(u, v); err != nil {
					return err
				}
				queue = append(queue, v)
				closed[v] = true
			}
		}
	}
	return nil
}

// Traverses g in depth-first order, starting at the given start nodes and
// calling callback for each vertex.
//
// The arguments for the callback function are the "from" and the "to"
// vertices. The callback is not called for the start nodes, and is called
// exactly once for each "to" vertex. Traversal stops when the callback
// returns a non-nil error; the error is propagated.
func DepthFirst(g graph.Directed, callback func(from, to int) error,
	start ...int) error {

	closed := make(map[int]bool)
	for _, i := range start {
		closed[i] = true
	}

	stack := start
	// Reverse the stack, so that we search the start nodes in the given order
	for i, j := 0, len(stack)-1; i < j; {
		stack[i], stack[j] = stack[j], stack[i]
		i++
		j--
	}

	for len(stack) != 0 {
		u := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		for _, v := range g.Neighbors(u) {
			if !closed[v] {
				if err := callback(u, v); err != nil {
					return err
				}
				closed[v] = true
				stack = append(stack, v)
			}
		}
	}
	return nil
}

// Traverses g in breadth-first order, starting at the given start node and
// calling callback for each vertex.
//
// Compared to BreadthFirst, this function uses less memory but may take more
// time. It also supports only a single start node.
func IterativeDeepening(g graph.Directed, callback func(from, to int) error,
	start int) (err error) {

	closed := map[int]bool{start: true}

	// Reports whether the callback got called during traversal (so we've seen
	// a new part of the graph), and if so, what it returned.
	var depthLimDF func(int, int) (bool, error)
	depthLimDF = func(u, depth int) (called bool, err error) {
		for _, v := range g.Neighbors(u) {
			if depth == 0 && !closed[v] {
				if err = callback(u, v); err != nil {
					return true, err
				}
				called = true
				closed[v] = true
			}
			if depth > 0 {
				var calledV bool
				calledV, err = depthLimDF(v, depth-1)
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
		morenodes, err = depthLimDF(start, depth)
		if !morenodes || err != nil {
			break
		}
	}
	return
}
