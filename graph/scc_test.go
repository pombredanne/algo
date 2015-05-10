package graph

import (
	"math/rand"
	"testing"
)

func TestSCC(t *testing.T) {
	// Graph from https://commons.wikimedia.org/wiki/File:Scc.png
	g := AdjacencyList{{1}, {2, 4, 5}, {3, 6}, {2, 7}, {0, 5}, {6}, {5}, {6, 3}}
	labels, ncomp := StrongComponents(g)

	if ncomp != 3 {
		t.Errorf("expected three components, got %d", ncomp)
	}
	for _, component := range [][]int{{0, 1, 4}, {2, 3, 7}, {5, 6}} {
		label := labels[component[0]]
		for _, u := range component[1:] {
			if labels[u] != label {
				t.Errorf("expected labels[%d] = %d, got %d",
					u, labels[u], label)
			}
		}
	}
}

func BenchmarkSCC(b *testing.B) {
	b.StopTimer()

	g := make(AdjacencyList, 10000)
	r := rand.New(rand.NewSource(125))
	for u := range g {
		nedges := r.Intn(len(g)) / 10
		for i := 0; i < nedges; i++ {
			g[u] = append(g[u], r.Intn(len(g)))
		}
	}

	_, ncomp := StrongComponents(g)
	b.Logf("number of strongly connected components: %d", ncomp)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		StrongComponents(g)
	}
}
