// Mathematical utilities.
package extmath

import "math"

// OEIS A001620
const euler = 0.57721566490153286060651209008240243104215933593992

// Digamma (psi) function: the logarithmic derivative of the gamma function.
//
// Currently only implemented for positive x.
func Digamma(x float64) float64 {
	if x <= 0 {
		panic("Digamma not implemented for non-positive x")
	}

	// J. Bernardo (1976). Algorithm AS 103: Psi (Digamma) Function.
	// Applied Statistics, http://www.uv.es/~bernardo/1976AppStatist.pdf.
	if x <= 1e-6 {
		return -euler - 1/x
	}

	result := 0.
	for x < 6 {
		result -= 1 / x
		x += 1
	}

	r := 1 / x
	result += math.Log(x) - .5*r
	r = r * r
	result -= r * ((1. / 12.) - r*((1./120.)-r*(1./252.)))
	return result
}

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
