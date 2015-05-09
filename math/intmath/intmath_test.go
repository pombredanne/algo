package intmath

import "testing"

func TestLog2(t *testing.T) {
	for _, c := range []struct{ n, log int }{
		{0, -1}, {1, 0}, {1 << 16, 16}, {1<<4 - 1, 3}, {1<<5 + 3, 5},
	} {
		if got := Log2(c.n); got != c.log {
			t.Errorf("expected %d, got %d", c.log, got)
		}
	}
}
