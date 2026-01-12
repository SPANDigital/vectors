package vectors

import "math"

func (a Vector) Equals(b Vector) bool {
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
func (v Vector) EqualsWithEpsilon(other Vector, epsilon float64) bool {
	if len(v) != len(other) {
		return false
	}
	for i, component := range v {
		if math.Abs(component-other[i]) > epsilon {
			return false
		}
	}
	return true
}
