package vectors

import "errors"

// DotProduct calculates the dot product (inner product) between two vectors.
// It returns an error if the vectors have different lengths.
// The dot product is: Σ(ai * bi)
func DotProduct(a Vector, b Vector) (float64, error) {
	if len(a) != len(b) {
		return 0.0, errors.New("vectors must have the same length")
	}

	sum := 0.0
	for i := range len(a) {
		sum += a[i] * b[i]
	}

	return sum, nil
}
