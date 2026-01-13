package vectors

import "errors"

// DotProduct calculates the dot product (inner product) between two vectors.
// It returns an error if the vectors have different lengths.
// The dot product is: Σ(ai * bi)
//
// This function works with both float32 and float64 vectors,
// maintaining the same type as the input vectors. Results are returned
// in the same type as the inputs.
//
// Example:
//
//	v1 := vectors.Vec[float64]{1.0, 2.0, 3.0}
//	v2 := vectors.Vec[float64]{4.0, 5.0, 6.0}
//	product, err := vectors.DotProduct(v1, v2) // Returns float64: 32.0
//
//	v3 := vectors.Vec[float32]{1.0, 2.0}
//	v4 := vectors.Vec[float32]{3.0, 4.0}
//	product2, err := vectors.DotProduct(v3, v4) // Returns float32: 11.0
func DotProduct[T Float](a Vec[T], b Vec[T]) (T, error) {
	if len(a) != len(b) {
		return T(0), errors.New("vectors must have the same length")
	}

	sum := T(0)
	for i := range len(a) {
		sum += a[i] * b[i]
	}

	return sum, nil
}
