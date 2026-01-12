package vectors

import "math"

// Magnitude returns the length (magnitude) of the vector.
func (v Vector) Magnitude() float64 {
	sumOfSquares := 0.0
	for _, component := range v {
		sumOfSquares += component * component
	}
	return math.Sqrt(sumOfSquares)
}