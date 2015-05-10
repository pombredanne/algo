package graph

import "testing"

func TestSCC(t *testing.T) {
	// Graph from https://commons.wikimedia.org/wiki/File:Scc.png
	g := AdjacencyList{{1}, {2, 4, 5}, {3, 6}, {2, 7}, {0, 5}, {6}, {5}, {6, 3}}
	labels := StrongComponents(g)

	for _, component := range [][]int{{0, 1, 4}, {2, 3, 7}, {5, 6}} {
		label := labels[component[0]]
		for _, u := range component[1:] {
			if labels[u] != label {
				t.Error("expected labels[%d] = %d, got %d", u, labels[u], label)
			}
		}
	}
}
