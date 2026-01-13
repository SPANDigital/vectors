package vectors

import (
	"errors"
	"math"
)

// TaxicabDistance calculates the taxicab distance (Manhattan distance, L1 distance) between two vectors.
// It returns an error if the vectors have different lengths.
// The taxicab distance is the sum of absolute differences: Σ|ai - bi|
//
// This function works with both float32 and float64 vectors,
// maintaining the same type as the input vectors. Results are returned
// in the same type as the inputs.
//
// Example:
//
//	v1 := vectors.Vec[float64]{1.0, 2.0}
//	v2 := vectors.Vec[float64]{4.0, 6.0}
//	distance, err := vectors.TaxicabDistance(v1, v2) // Returns float64: 7.0
func TaxicabDistance[T Float](a Vec[T], b Vec[T]) (T, error) {
	if len(a) != len(b) {
		return T(0), errors.New("vectors must have the same length")
	}

	sum := T(0)
	for i := range len(a) {
		sum += T(math.Abs(float64(a[i] - b[i])))
	}

	return sum, nil
}
