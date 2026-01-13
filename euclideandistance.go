package vectors

import (
	"errors"
	"math"
)

// EuclideanDistance calculates the Euclidean distance between two vectors.
// It returns an error if the vectors have different lengths.
// The Euclidean distance is the straight-line distance between two points: √(Σ(ai - bi)²)
//
// This function works with both float32 and float64 vectors,
// maintaining the same type as the input vectors. Results are returned
// in the same type as the inputs.
//
// Example:
//
//	v1 := vectors.Vec[float64]{1.0, 2.0}
//	v2 := vectors.Vec[float64]{4.0, 6.0}
//	distance, err := vectors.EuclideanDistance(v1, v2) // Returns float64: 5.0
func EuclideanDistance[T Float](a Vec[T], b Vec[T]) (T, error) {
	if len(a) != len(b) {
		return T(0), errors.New("vectors must have the same length")
	}

	sumOfSquares := T(0)
	for i := range len(a) {
		diff := a[i] - b[i]
		sumOfSquares += diff * diff
	}

	return T(math.Sqrt(float64(sumOfSquares))), nil
}
