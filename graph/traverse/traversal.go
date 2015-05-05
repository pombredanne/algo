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
