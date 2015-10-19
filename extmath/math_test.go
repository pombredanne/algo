package extmath

import (
	"math"
	"testing"
)

func TestDigamma(t *testing.T) {
	// y values based on the SciPy implementation's output.
	for _, c := range []struct{ x, y float64 }{
		{1e-60, -9.9999999999999995e+59},
		{1e-10, -10000000000.577215},
		{.2, -5.2890398965921879},
		{1, -0.57721566490153287},
		{3.14159, 0.9772123148520715},
		{12.15, 2.4556127846565494},
		{916, 6.8194704138277933},
	} {
		if math.Abs(Digamma(c.x)-c.y) > 3e-9 {
			t.Errorf("expected %f, got %f", c.y, Digamma(c.x))
		}
	}
}

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
