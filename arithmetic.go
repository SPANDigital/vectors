package vectors

import "errors"

// Add returns a new vector that is the element-wise sum of v and other.
// It returns an error if the vectors have different lengths.
// The result is: result[i] = v[i] + other[i]
func (v Vector) Add(other Vector) (Vector, error) {
	if len(v) != len(other) {
		return nil, errors.New("vectors must have the same length")
	}

	result := make(Vector, len(v))
	for i := range len(v) {
		result[i] = v[i] + other[i]
	}

	return result, nil
}

// Subtract returns a new vector that is the element-wise difference of v and other.
// It returns an error if the vectors have different lengths.
// The result is: result[i] = v[i] - other[i]
func (v Vector) Subtract(other Vector) (Vector, error) {
	if len(v) != len(other) {
		return nil, errors.New("vectors must have the same length")
	}

	result := make(Vector, len(v))
	for i := range len(v) {
		result[i] = v[i] - other[i]
	}

	return result, nil
}
