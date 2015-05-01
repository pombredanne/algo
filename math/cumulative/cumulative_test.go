package cumulative

import (
	"math"
	"testing"
)

func assertClose(t *testing.T, exp, obs, epsilon float64) {
	if math.Abs(exp-obs) > epsilon {
		t.Errorf("expected %g, got %g (difference > %g)", exp, obs, epsilon)
	}
}

func TestAccumulators(t *testing.T) {
	var avg Mean
	var sum Sum
	var Var Variance

	for i := 0; i < 5000; i++ {
		avg.Add(.1)
		sum.Add(.1)
		Var.Add(.1)
	}

	assertClose(t, .1, avg.Value(), 1e-15)
	assertClose(t, 500, sum.Sum, 1e-15)
	if Var.Value() != 0 {
		t.Errorf("expected zero variance, got %f", Var.Value())
	}

	if sum.Value() != sum.Sum {
		t.Errorf("Sum.Sum not equal to Sum.Value: %g != %g",
			sum.Value(), sum.Sum)
	}
}
