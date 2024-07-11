package vectors

import "math"

func (v Vector) IsNormalized() bool {
	sumOfSquares := 0.0
	for _, component := range v {
		sumOfSquares += component * component
	}
	length := math.Sqrt(sumOfSquares)
	return math.Abs(length-1.0) < 1e-9 // Use a small epsilon for floating point comparison
}
