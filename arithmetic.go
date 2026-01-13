package vectors

import "errors"

// Add returns a new vector that is the element-wise sum of v and other.
// It returns an error if the vectors have different lengths.
// The result is: result[i] = v[i] + other[i]
//
// This method works with both float32 and float64 vectors,
// maintaining the same type as the input vectors.
//
// Example:
//
//	v1 := vectors.Vec[float64]{1.0, 2.0, 3.0}
//	v2 := vectors.Vec[float64]{4.0, 5.0, 6.0}
//	sum, err := v1.Add(v2) // Returns Vec[float64]{5.0, 7.0, 9.0}
func (v Vec[T]) Add(other Vec[T]) (Vec[T], error) {
	if len(v) != len(other) {
		return nil, errors.New("vectors must have the same length")
	}

	result := make(Vec[T], len(v))
	for i := range len(v) {
		result[i] = v[i] + other[i]
	}

	return result, nil
}

// Subtract returns a new vector that is the element-wise difference of v and other.
// It returns an error if the vectors have different lengths.
// The result is: result[i] = v[i] - other[i]
//
// This method works with both float32 and float64 vectors,
// maintaining the same type as the input vectors.
//
// Example:
//
//	v1 := vectors.Vec[float64]{5.0, 7.0, 9.0}
//	v2 := vectors.Vec[float64]{1.0, 2.0, 3.0}
//	diff, err := v1.Subtract(v2) // Returns Vec[float64]{4.0, 5.0, 6.0}
func (v Vec[T]) Subtract(other Vec[T]) (Vec[T], error) {
	if len(v) != len(other) {
		return nil, errors.New("vectors must have the same length")
	}

	result := make(Vec[T], len(v))
	for i := range len(v) {
		result[i] = v[i] - other[i]
	}

	return result, nil
}
