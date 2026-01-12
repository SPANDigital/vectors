package vectors

import "errors"

// NegativeInnerProduct calculates the negative inner product (negative dot product) between two vectors.
// It returns an error if the vectors have different lengths.
// The negative inner product is: -(Σ(ai * bi))
// This is commonly used as a distance metric in machine learning.
func NegativeInnerProduct(a Vector, b Vector) (float64, error) {
	if len(a) != len(b) {
		return 0.0, errors.New("vectors must have the same length")
	}

	sum := 0.0
	for i := range len(a) {
		sum += a[i] * b[i]
	}

	return -sum, nil
}
