// Package bdd implements binary decision diagrams.
//
// BDDs are compact representations of boolean functions that allow finding
// satisfying assignments of truth values to their variables.
package bdd

import (
	"fmt"

	"github.com/larsmans/algo/logic/boolean"
)

type BoolOp int

// Binary boolean operators.
const (
	OpAnd BoolOp = iota
	OpIff
	OpImplies
	OpOr
	OpXor
)

// Reduced, ordered, binary decision diagram (ROBDD) builder.
//
// A Builder can be used to construct ROBDDs; multiple of these can be combined
// into larger BDDs only if they were constructed using the same Builder.
type Builder struct {
	nvars   int
	inverse map[node]*node

	nodecache *[24]node
	ncached   int
}

func NewBuilder() *Builder {
	return &Builder{
		inverse: make(map[node]*node),
	}
}

// Reduced, ordered, binary decision diagram.
type Robdd struct {
	root    *node
	builder *Builder
}

type node struct {
	v      boolean.Var
	lo, hi *node
}

// Boolean constants.
var (
	truenode  = node{v: -1}
	falsenode = node{v: -1}
)

func init() {
	truenode.lo, falsenode.lo = &falsenode, &truenode
	truenode.hi, falsenode.hi = &truenode, &falsenode
}

func isterminal(u *node) bool { return u == &truenode || u == &falsenode }

// Constant constructs a new BDD that represents the boolean constant v.
func (b *Builder) Constant(v bool) *Robdd {
	node := &falsenode
	if v {
		node = &truenode
	}
	return &Robdd{root: node, builder: b}
}

// NewLiteral constructs a new BDD that represents the variable v.
func (b *Builder) NewLiteral(v boolean.Var) *Robdd {
	u := b.mknode(v, &falsenode, &truenode)
	if int(v) > b.nvars {
		b.nvars = int(v) + 1
	}
	return &Robdd{root: u, builder: b}
}

// Implements the variable ordering between nodes.
// At least one of u, v must be a non-terminal.
func varbefore(u, v *node) bool {
	switch {
	case isterminal(u):
		return false
	case isterminal(v):
		return true
	default:
		return u.v < v.v
	}
}

// Apply the boolean operator op to d1 and d2 to construct a new BDD.
//
// d1 and d2 must have been constructed using the same Builder.
func (b *Builder) Apply(op BoolOp, d1, d2 *Robdd) *Robdd {
	if d1.builder != d2.builder {
		panic("cannot Apply operator to BDD from distinct Builders")
	}
	memoized := make(map[[2]*node]*node)

	var apply func(u, v *node) *node
	apply = func(u, v *node) (r *node) {
		if isterminal(u) && isterminal(v) {
			return applyTerminals(op, u, v)
		}
		if r, ok := memoized[[2]*node{u, v}]; ok {
			return r
		}
		if u.v == v.v {
			r = b.mknode(u.v, apply(u.lo, v.lo), apply(u.hi, v.hi))
		} else if varbefore(u, v) {
			r = b.mknode(u.v, apply(u.lo, v), apply(u.hi, v))
		} else {
			r = b.mknode(v.v, apply(u, v.lo), apply(u, v.hi))
		}
		memoized[[2]*node{u, v}] = r
		return r
	}
	return &Robdd{root: apply(d1.root, d2.root), builder: b}
}

func applyTerminals(op BoolOp, u, v *node) *node {
	var a, b, c bool
	a, b = (u == &truenode), (v == &truenode)

	switch op {
	case OpAnd:
		c = a && b
	case OpIff:
		c = a == b
	case OpImplies:
		c = !a || b
	case OpOr:
		c = a || b
	case OpXor:
		c = a != b
	default:
		panic(fmt.Errorf("unknown boolean operator %d", op))
	}
	if c {
		return &truenode
	}
	return &falsenode
}

func (b *Builder) mknode(v boolean.Var, l, h *node) *node {
	if l == h {
		return l
	}
	if u, ok := b.inverse[node{v, l, h}]; ok {
		return u
	}
	return b.newnode(v, l, h)
}

func (b *Builder) newnode(v boolean.Var, l, h *node) *node {
	if b.ncached == 0 {
		b.nodecache = &[24]node{}
		b.ncached = len(*b.nodecache)
	}
	b.ncached--
	n := &b.nodecache[b.ncached]

	*n = node{v, l, h}
	b.inverse[*n] = n
	return n
}

// Returns an assignment of truth values to variables that satisfies d, or
// nil if d is not satisfiable.
func (d *Robdd) AnySat() map[boolean.Var]bool {
	assign := make(map[boolean.Var]bool)
	for u := d.root; ; {
		switch {
		case u == &falsenode:
			return nil
		case u == &truenode:
			return assign
		case u.lo == &falsenode:
			assign[u.v] = true
			u = u.hi
		default:
			assign[u.v] = false
			u = u.lo
		}
	}
}

// Size reports the number of nodes in d.
func (d *Robdd) Size() int {
	seen := make(map[*node]bool)
	var size func(*node) int
	size = func(n *node) int {
		if seen[n] {
			return 0
		}
		seen[n] = true
		if isterminal(n) {
			return 1
		}
		return size(n.lo) + size(n.hi)
	}
	return size(d.root)
}
