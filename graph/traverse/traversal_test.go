package traverse

import (
	"github.com/larsmans/algo/graph"
	"reflect"
	"sort"
	"testing"
)

type adjacencyList [][]int

func (g adjacencyList) Neighbors(u int) []int {
	return g[u]
}

func TestUnweighted(t *testing.T) {
	g := adjacencyList([][]int{
		{1, 4}, {2, 5, 0}, {3}, {1, 3}, {4}, {2, 5},
		{7}, // would cause out-of-bounds if visited
	})

	var vset map[int]bool
	callback := func(from, to int) error {
		if vset[to] {
			t.Errorf("%d already visited", to)
		}
		vset[to] = true
		return nil
	}

	type typ func(g graph.Directed, cb func(int, int) error, start int) error
	bfs := func(g graph.Directed, cb func(int, int) error, start int) error {
		return BreadthFirst(g, cb, start)
	}
	dfs := func(g graph.Directed, cb func(int, int) error, start int) error {
		return DepthFirst(g, cb, start)
	}

	expected := []int{1, 2, 3, 4, 5}
	for _, traversal := range []typ{bfs, dfs} {
		vset = make(map[int]bool)
		traversal(g, callback, 0)

		var vs []int
		for i := range vset {
			vs = append(vs, i)
		}
		sort.Ints(vs)
		if !reflect.DeepEqual(vs, expected) {
			t.Errorf("set of visited nodes = %v, expected %v", vs, expected)
		}
	}
}
