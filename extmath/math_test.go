package extmath

import (
	"math"
	"testing"
)

func TestLogAddExp(t *testing.T) {
	x := LogAddExp(1e-10, math.Inf(-1))
	if x != 1e-10 {
		t.Errorf("expected exactly 1e-10, got %g", x)
	}
	t.Logf("naïve version returns %g",
		math.Log(math.Exp(1e-10)+math.Exp(math.Inf(-1))))
}
