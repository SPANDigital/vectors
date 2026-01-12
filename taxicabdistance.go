package vectors

import (
	"errors"
	"math"
)

// TaxicabDistance calculates the taxicab distance (Manhattan distance, L1 distance) between two vectors.
// It returns an error if the vectors have different lengths.
// The taxicab distance is the sum of absolute differences: Σ|ai - bi|
func TaxicabDistance(a Vector, b Vector) (float64, error) {
	if len(a) != len(b) {
		return 0.0, errors.New("vectors must have the same length")
	}

	sum := 0.0
	for i := range len(a) {
		sum += math.Abs(a[i] - b[i])
	}

	return sum, nil
}
