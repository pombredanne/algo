package traverse

import (
	"errors"
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
		{7}, // would cause out-of-bounds panic if visited, but it never is
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
	for _, traversal := range []typ{bfs, dfs, IterativeDeepening} {
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

	// Test BFS order.
	expected = []int{1, 4, 2, 5, 3}
	for _, traversal := range []typ{bfs, IterativeDeepening} {
		vs := make([]int, 0)
		err := traversal(g, func(from, to int) error {
			vs = append(vs, to)
			if to == 3 {
				return errors.New("found it!")
			}
			return nil
		}, 0)

		if err == nil || err.Error() != "found it!" {
			t.Errorf("unexpected error value %v", err)
		}
	}
}
