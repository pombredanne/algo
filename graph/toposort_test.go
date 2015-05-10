package graph

import (
	//"math"
	"math/rand"
	"testing"
)

func TestTopoSort(t *testing.T) {
	// Example from CLRS, 3rd ed., p. 613.
	const (
		undershorts = iota
		socks
		pants
		shoes
		watch // unconnected node
		shirt
		belt
		tie
		jacket
		nvertices
	)
	g := make(AdjacencyList, nvertices)
	g[undershorts] = []int{pants, shoes}
	g[pants] = []int{shoes, belt}
	g[shirt] = []int{belt, tie}
	g[belt] = []int{jacket}
	g[shirt] = []int{tie}
	g[tie] = []int{jacket}

	label, err := TopoSort(g)

	seen := make([]bool, nvertices)
	if err != nil {
		t.Errorf("got err = %v from TopoSort, wanted nil", err)
	}
	for _, u := range label {
		seen[u] = true
	}
	for u := range seen {
		if !seen[u] {
			t.Errorf("vertex %d not labeled", u)
		}
	}

	for _, c := range []struct{ before, after int }{
		{socks, shoes},
		{undershorts, pants},
		{undershorts, shoes},
		{pants, belt},
		{shirt, belt},
		{shirt, tie},
		{belt, jacket},
		{tie, jacket},
	} {
		if label[c.before] > label[c.after] {
			t.Errorf("expected %d before %d, but %d > %d",
				c.before, c.after, label[c.before], label[c.after])
		}
	}

	// Introduce a cycle.
	g[jacket] = []int{pants}
	label, err = TopoSort(g)
	if label != nil || err == nil {
		t.Errorf("expected nil return and non-nil error, got %v and %v",
			label, err)
	}
}

func BenchmarkTopoSort(b *testing.B) {
	b.StopTimer()

	// Make a random, sparsish graph.
	g := make(AdjacencyList, 1000)
	r := rand.New(rand.NewSource(126))
	for u := range g {
		max := len(g) - u - 1
		nedges := int(((r.NormFloat64()) * .25) * float64(max))
		switch {
		case nedges < 0:
			nedges = 0
		case nedges > max:
			nedges = max
		}
		for v := 0; v < nedges; v++ {
			g[u] = append(g[u], u + 1 + r.Intn(len(g) - u - 1))
		}
	}

	_, err := TopoSort(g)
	if err != nil {
		b.Fatal(err)
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		TopoSort(g)
	}
}
