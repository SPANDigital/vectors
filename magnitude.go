package vectors

import "math"

// Magnitude returns the length (magnitude) of the vector.
//
// This method works with both float32 and float64 vectors,
// maintaining the same type as the input vector. Results are returned
// in the same type as the input.
//
// Example:
//
//	v64 := vectors.Vec[float64]{3.0, 4.0}
//	mag64 := v64.Magnitude() // Returns float64: 5.0
//
//	v32 := vectors.Vec[float32]{3.0, 4.0}
//	mag32 := v32.Magnitude() // Returns float32: 5.0
func (v Vec[T]) Magnitude() T {
	sumOfSquares := T(0)
	for _, component := range v {
		sumOfSquares += component * component
	}
	return T(math.Sqrt(float64(sumOfSquares)))
}