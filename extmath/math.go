// Mathematical utilities.
package extmath

import "math"

// log(exp(a) + exp(b)), evaluated in a numerically stable way.
func LogAddExp(a, b float64) float64 {
	if b > a {
		a, b = b, a
	}
	return a + math.Log1p(math.Exp(b-a))
}

// Logistic function: 1/(1+exp(-x)).
func Logistic(x float64) float64 {
	return .5 * (1 + math.Tanh(.5*x))
}

// Logit, the inverse of the logistic function: log(p/(1-p)).
//
// Special cases are:
//	Logit(1) = Inf
//	Logit(0) = -Inf
//	Logit(p < 0) = NaN
//	Logit(p > 1) = NaN
func Logit(p float64) float64 {
	return math.Log(p/(1-p))
}
