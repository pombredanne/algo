// Copyright 2015 Lars Buitinck
//
// MIT-licensed. See the file LICENSE for details.

// The types in this package implement accumulators for various statistics.
//
// The zero value of each type represents an "empty" accumulator.
// The Value method can be called at any time to obtain the accumulated value.
package cumulative

// Cumulative moving average.
type Mean struct {
	sum Sum
	n   uint64
}

func (avg *Mean) Add(x float64) {
	avg.n++
	avg.sum.Add((x - avg.sum.Sum) / float64(avg.n))
}

// Returns the mean of the values seen by Add.
//
// The mean of zero values is undefined; requesting it may cause a panic.
func (avg *Mean) Value() float64 {
	return avg.sum.Sum
}

// Stable cumulative sum (Kahan's algorithm).
//
// The Sum member is the value returned by the Value method.
type Sum struct {
	Sum, c float64
}

func (s *Sum) Add(x float64) {
	y := x - s.c
	t := s.Sum + y
	s.c = (t - s.Sum) - y
	s.Sum = t
}

func (s *Sum) Value() float64 {
	return s.Sum
}

// Cumulative (weighted) variance.
type Variance struct {
	m2, mean, wsum float64
	nobs           int64
}

func (v *Variance) Add(x float64) {
	v.AddW(x, 1)
}

// Adds value x with weight w to the accumulator v.
func (v *Variance) AddW(x, w float64) {
	wsum := v.wsum + w
	delta := x - v.mean
	r := delta * w / wsum
	v.mean += r
	v.m2 += v.wsum * delta * r
	v.wsum = wsum
}

// Reports the variance of the values seen by Add and AddW.
func (v *Variance) Value() float64 {
	if v.nobs == 1 {
		return 0
	}
	nobs := float64(v.nobs)
	return v.m2 / v.wsum * nobs / (nobs - 1)
}
