// Copyright 2015 Lars Buitinck
//
// MIT-licensed. See the file LICENSE for details.

package graph

// Strongly connected components algorithm.
//
// Returns components labels for the vertices in g.
func StrongComponents(g FiniteDirected) []int {
	// Tarjan's algorithm.
	n := g.NVertices()

	index := make([]int, n) // 0 means undefined
	label := make([]int, n)
	lowlink := make([]int, n)
	onstack := make([]bool, n)
	stack := make([]int, 0)

	curlabel, nextindex := 0, 1
	var visit func(int)
	visit = func(u int) {
		index[u] = nextindex
		lowlink[u] = nextindex
		nextindex++
		stack = append(stack, u)
		onstack[u] = true

		for _, v := range g.Neighbors(u) {
			if index[v] == 0 {
				visit(v)
				lowlink[u] = min(lowlink[u], lowlink[v])
			} else if onstack[v] {
				lowlink[u] = min(lowlink[u], index[v])
			}
		}

		if lowlink[u] == index[u] {
			for {
				v := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				onstack[v] = false
				label[v] = curlabel
				if u == v {
					break
				}
			}
			curlabel++
		}
	}

	for u := 0; u < n; u++ {
		if index[u] == 0 {
			visit(u)
		}
	}
	return label
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
