package vectors

import "math"

// Equals compares two vectors for exact equality.
// Returns false if vectors have different lengths or any components differ.
//
// This method works with both float32 and float64 vectors.
//
// Example:
//
//	v1 := vectors.Vec[float64]{1.0, 2.0, 3.0}
//	v2 := vectors.Vec[float64]{1.0, 2.0, 3.0}
//	equal := v1.Equals(v2) // Returns true
func (a Vec[T]) Equals(b Vec[T]) bool {
	if len(a) != len(b) {
		return false
	}
	for i, component := range a {
		if component != b[i] {
			return false
		}
	}
	return true
}

// EqualsWithEpsilon compares two vectors for equality within a specified tolerance.
// Returns true if the absolute difference between each corresponding component is
// less than or equal to epsilon. Returns false if vectors have different lengths.
//
// This method works with both float32 and float64 vectors,
// maintaining the same type as the input vectors.
//
// Example:
//
//	v1 := vectors.Vec[float64]{1.0, 2.0}
//	v2 := vectors.Vec[float64]{1.0001, 2.0001}
//	equal := v1.EqualsWithEpsilon(v2, 0.001) // Returns true
//
//	v3 := vectors.Vec[float32]{0.6, 0.8}
//	v4 := vectors.Vec[float32]{0.6000001, 0.8000001}
//	equal2 := v3.EqualsWithEpsilon(v4, vectors.DefaultEpsilon32) // Returns true
func (v Vec[T]) EqualsWithEpsilon(other Vec[T], epsilon T) bool {
	if len(v) != len(other) {
		return false
	}
	for i, component := range v {
		if T(math.Abs(float64(component-other[i]))) > epsilon {
			return false
		}
	}
	return true
}
