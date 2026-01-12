package vectors

import (
	"errors"
	"math"
)

// EuclideanDistance calculates the Euclidean distance between two vectors.
// It returns an error if the vectors have different lengths.
// The Euclidean distance is the straight-line distance between two points: √(Σ(ai - bi)²)
func EuclideanDistance(a Vector, b Vector) (float64, error) {
	if len(a) != len(b) {
		return 0.0, errors.New("vectors must have the same length")
	}

	sumOfSquares := 0.0
	for i := range len(a) {
		diff := a[i] - b[i]
		sumOfSquares += diff * diff
	}

	return math.Sqrt(sumOfSquares), nil
}
