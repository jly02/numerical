// Package diff implements multiple numerical
// differentiation algorithms.
package diff

import "math"

// fn type allows functions to be passed as parameters
type fn func(float64) float64

// DiffQuot returns a slice containing the derivative of the given
// function, computed using a difference quotient.
func DiffQuot(f fn, x_0, x_1, dx float64) []float64 {
	// number of points to be sampled
	sample_num := int(math.Abs(x_1-x_0) / dx)
	df_dx := make([]float64, sample_num)

	// start at x_0
	x_i := x_0
	for i := 0; i < sample_num; i++ {
		// Newton's difference quotient
		df_dx[i] = (f(x_i+dx) - f(x_i)) / dx
		x_i += dx
	}

	return df_dx
}

// SymDiffQuot returns a slice containing the derivative of the
// given function, computed with a symmetric difference quotient.
func SymDiffQuot(f fn, x_0, x_1, dx float64) []float64 {
	// number of points to be sampled
	sample_num := int(math.Abs(x_1-x_0) / dx)
	df_dx := make([]float64, sample_num)

	// start at x_0
	x_i := x_0
	for i := 0; i < sample_num; i++ {
		// Symmetric difference quotient (secant)
		df_dx[i] = (f(x_i+dx) - f(x_i-dx)) / (2 * dx)
		x_i += dx
	}

	return df_dx
}
