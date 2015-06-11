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

func TestLogistic(t *testing.T) {
	for _, x := range []float64{
		1e-300, 1e-15, 1e-7, .2, .3, .5,
	} {
		for _, x := range []float64{x, -x} {
			p := Logistic(x)
			logit := Logit(p)
			if err := math.Abs(logit - x); err > 3e-16 {
				t.Errorf("mismatch: %g differs from %g by %g", Logit(p), x, err)
			}
		}
	}

	for _, c := range []struct{p, logit float64} {
		{1, math.Inf(1)},
		{0, math.Inf(-1)},
	} {
		if logit := Logit(c.p); logit != c.logit {
			t.Errorf("Logit error: got %g, wanted %g", logit, c.logit)
		}
	}
	for _, p := range []float64{-1, -1e300, 1+1e300, 2} {
		if !math.IsNaN(Logit(p)) {
			t.Errorf("expected NaN for Logit(%g), got %g", p, Logit(p))
		}
	}
}
