package vectors

import "math"

func (v Vector) Normalize() Vector {
	sumOfSquares := 0.0
	for _, component := range v {
		sumOfSquares += component * component
	}
	length := math.Sqrt(sumOfSquares)

	// Avoid division by zero if the vector is a zero vector
	if length == 0 {
		return v
	}

	normalizedVector := make(Vector, len(v))
	for i, component := range v {
		normalizedVector[i] = component / length
	}
	return normalizedVector
}
