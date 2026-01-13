package vectors

import "math"

// Normalize returns a new vector with the same direction but unit length (magnitude of 1.0).
// If the vector is a zero vector, it returns the zero vector unchanged.
//
// This method works with both float32 and float64 vectors,
// maintaining the same type as the input vector.
//
// Example:
//
//	v64 := vectors.Vec[float64]{3.0, 4.0}
//	norm64 := v64.Normalize() // Returns Vec[float64]{0.6, 0.8}
//
//	v32 := vectors.Vec[float32]{3.0, 4.0}
//	norm32 := v32.Normalize() // Returns Vec[float32]{0.6, 0.8}
func (v Vec[T]) Normalize() Vec[T] {
	sumOfSquares := T(0)
	for _, component := range v {
		sumOfSquares += component * component
	}
	length := T(math.Sqrt(float64(sumOfSquares)))

	// Avoid division by zero if the vector is a zero vector
	if length == 0 {
		return v
	}

	normalizedVector := make(Vec[T], len(v))
	for i, component := range v {
		normalizedVector[i] = component / length
	}
	return normalizedVector
}
