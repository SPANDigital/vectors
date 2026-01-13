package vectors

import (
	"errors"
	"math"
)

// CosineSimilarity calculates the cosine similarity between two vectors.
// Returns a value between -1.0 and 1.0, where:
// - 1.0 indicates identical direction
// - 0.0 indicates orthogonal vectors
// - -1.0 indicates opposite direction
//
// Returns an error if either vector is a zero vector.
//
// Note: This implementation handles vectors of different lengths by treating
// missing components as zero.
//
// This function works with both float32 and float64 vectors,
// maintaining the same type as the input vectors. Results are returned
// in the same type as the inputs.
//
// Example:
//
//	v1 := vectors.Vec[float64]{1.0, 2.0, 3.0}
//	v2 := vectors.Vec[float64]{4.0, 5.0, 6.0}
//	similarity, err := vectors.CosineSimilarity(v1, v2) // Returns float64
func CosineSimilarity[T Float](a Vec[T], b Vec[T]) (cosine T, err error) {
	count := 0
	length_a := len(a)
	length_b := len(b)
	if length_a > length_b {
		count = length_a
	} else {
		count = length_b
	}
	sumA := T(0)
	s1 := T(0)
	s2 := T(0)
	for k := 0; k < count; k++ {
		if k >= length_a {
			s2 += T(math.Pow(float64(b[k]), 2))
			continue
		}
		if k >= length_b {
			s1 += T(math.Pow(float64(a[k]), 2))
			continue
		}
		sumA += a[k] * b[k]
		s1 += T(math.Pow(float64(a[k]), 2))
		s2 += T(math.Pow(float64(b[k]), 2))
	}
	if s1 == 0 || s2 == 0 {
		return T(0), errors.New("vectors should not be null (all zeros)")
	}
	return sumA / (T(math.Sqrt(float64(s1))) * T(math.Sqrt(float64(s2)))), nil
}
