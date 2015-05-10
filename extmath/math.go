// Mathematical utilities.
package math

import "math"

// log(exp(a) + exp(b)), evaluated in a numerically stable way.
func LogAddExp(a, b float64) float64 {
	if b > a {
		a, b = b, a
	}
	return a + math.Log1p(math.Exp(b-a))
}
