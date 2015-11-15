package bdd

import (
	"testing"

	"github.com/larsmans/algo/logic/boolean"
)

func TestBDD(t *testing.T) {
	b := NewBuilder()

	// Construct the formula (0 and 1) or (2 and (1 or 0)).
	var v []*Robdd
	for i := 0; i < 4; i++ {
		v = append(v, b.NewLiteral(boolean.Var(i)))
	}
	f := b.Apply(OpAnd, v[0], v[1])
	g := b.Apply(OpAnd, v[2], b.Apply(OpOr, v[1], v[0]))
	h := b.Apply(OpOr, f, g)
	_ = h

	for _, c := range []struct {
		bdd    *Robdd
		assign map[boolean.Var]bool
	}{
		{f, map[boolean.Var]bool{0: false, 1: false}},
		{g, map[boolean.Var]bool{0: true, 1: true, 2: false}},
		{h, map[boolean.Var]bool{0: false, 1: false, 2: true}},
	} {
		if checkSat(c.bdd.root, c.assign) {
			t.Error("construction error")
		}
	}

	for _, bdd := range []*Robdd{f, g, h} {
		if assign := bdd.AnySat(); !checkSat(bdd.root, assign) {
			t.Errorf("not a SAT solution: %v", assign)
		}
	}
}

func TestSAT(t *testing.T) {
	b := NewBuilder()
	unsat := b.Constant(false)

	check := func() {
		if assign := unsat.AnySat(); assign != nil {
			t.Errorf("AnySat returned %v for unsatisfiable formula",
				assign)
		}
		if n := unsat.Size(); n != 1 {
			t.Errorf("unsatisfiable formula has %d nodes, expected one", n)
		}
	}
	check()

	unsat = b.Apply(OpAnd, unsat, b.Constant(true))
	unsat = b.Apply(OpAnd, unsat, b.NewLiteral(boolean.Var(1)))
	unsat = b.Apply(OpAnd, unsat, unsat)
}

func checkSat(u *node, assign map[boolean.Var]bool) bool {
	switch {
	case u == &truenode:
		return true
	case u == &falsenode:
		return false
	}
	if val, ok := assign[u.v]; ok {
		if val {
			return checkSat(u.hi, assign)
		} else {
			return checkSat(u.lo, assign)
		}
	}
	return checkSat(u.hi, assign) || checkSat(u.lo, assign)
}

func BenchmarkBuilder(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		b := NewBuilder()
		lit := make([]*Robdd, 10)
		for j := range lit {
			lit[j] = b.NewLiteral(boolean.Var(j))
		}
		d1 := b.Apply(OpXor, lit[0], lit[1])
		d2 := b.Apply(OpOr, lit[2], lit[3])
		d3 := b.Apply(OpAnd, lit[4], lit[5])
		d4 := b.Apply(OpIff, lit[6], lit[7])
		d5 := b.Apply(OpImplies, lit[8], lit[9])

		d1 = b.Apply(OpXor, d1, d3)
		d2 = b.Apply(OpAnd, d2, d5)
		d3 = b.Apply(OpOr, d1, d4)
		d4 = b.Apply(OpImplies, d2, d3)
		d5 = b.Apply(OpIff, b.NewLiteral(boolean.Var(10)), d4)
	}
}
