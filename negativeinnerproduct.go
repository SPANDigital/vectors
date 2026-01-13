package vectors

import "errors"

// NegativeInnerProduct calculates the negative inner product (negative dot product) between two vectors.
// It returns an error if the vectors have different lengths.
// The negative inner product is: -(Σ(ai * bi))
// This is commonly used as a distance metric in machine learning.
//
// This function works with both float32 and float64 vectors,
// maintaining the same type as the input vectors. Results are returned
// in the same type as the inputs.
//
// Example:
//
//	v1 := vectors.Vec[float64]{1.0, 2.0, 3.0}
//	v2 := vectors.Vec[float64]{4.0, 5.0, 6.0}
//	product, err := vectors.NegativeInnerProduct(v1, v2) // Returns float64: -32.0
func NegativeInnerProduct[T Float](a Vec[T], b Vec[T]) (T, error) {
	if len(a) != len(b) {
		return T(0), errors.New("vectors must have the same length")
	}

	sum := T(0)
	for i := range len(a) {
		sum += a[i] * b[i]
	}

	return -sum, nil
}
