package vectors

import "math"

// IsNormalized checks whether the vector has unit length (magnitude of 1.0).
// Uses type-appropriate epsilon comparison for floating-point precision
// (1e-9 for float64, 1e-7 for float32).
//
// This method works with both float32 and float64 vectors.
//
// Example:
//
//	v64 := vectors.Vec[float64]{0.6, 0.8}
//	isNorm64 := v64.IsNormalized() // Returns true
//
//	v32 := vectors.Vec[float32]{0.6, 0.8}
//	isNorm32 := v32.IsNormalized() // Returns true
func (v Vec[T]) IsNormalized() bool {
	sumOfSquares := T(0)
	for _, component := range v {
		sumOfSquares += component * component
	}
	length := T(math.Sqrt(float64(sumOfSquares)))
	eps := epsilon[T]()
	return T(math.Abs(float64(length-1.0))) <= eps
}
